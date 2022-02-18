package nits

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

// httpUtility is an empty structure that is prepared only for creating methods.
type httpUtility struct{}

// HTTP is an entity that allows the methods of HTTPUtility to be executed from outside the package without initializing HTTPUtility.
// nolint: gochecknoglobals
var HTTP httpUtility

// Terminate returns http.HandlerFunc that sends SIGTERM to the shutdown signal.
func (httpUtility) Terminate(preprocess http.HandlerFunc, shutdownChan chan<- os.Signal, signal os.Signal) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preprocess(w, r)
		shutdownChan <- signal
	}
}

// Listen will start *http.Server and ignore http.ErrServerClosed on shutdown.
// Listen expects to be used in conjunction with Shutdown:.
//
//	func startServer(ctx context.Context, server *http.Server) error {
//		shutdownChan := make(chan os.Signal, 1)
//		signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
//
//		// wait signal for shutdown
//		go func() {
//			sig, err := nits.HTTP.Shutdown(ctx, server, 5*time.Second, shutdownChan)
//			if err != nil {
//				log.Printf("Shutdown: %v", err)
//			}
//		}()
//
//		// start server
//		if err := nits.HTTP.Listen(server); err != nil {
//			return fmt.Errorf("Listen: %w", err)
//		}
//
//		return nil
//	}
func (httpUtility) Listen(server *http.Server) error {
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("ListenAndServe: %w", err)
	}

	return nil
}

// Shutdown waits for a shutdown signal and terminates *http_Server when catches the signal.
// Shutdown expects to be executed in goroutine:
//
//	func startServer(ctx context.Context, server *http.Server) error {
//		shutdownChan := make(chan os.Signal, 1)
//		signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
//
//		// wait signal for shutdown
//		go func() {
//			sig, err := nits.HTTP.Shutdown(ctx, server, 5*time.Second, shutdownChan)
//			if err != nil {
//				log.Printf("Shutdown: %v", err)
//			}
//		}()
//
//		// start server
//		if err := nits.HTTP.Listen(server); err != nil {
//			return fmt.Errorf("Listen: %w", err)
//		}
//
//		return nil
//	}
func (httpUtility) Shutdown(ctx context.Context, server *http.Server, shutdownTimeout time.Duration, shutdownChan <-chan os.Signal) (caught os.Signal, err error) {
	return HTTP.shutdown(ctx, server, shutdownTimeout, shutdownChan)
}

type shutdowner interface {
	Shutdown(ctx context.Context) error
}

func (httpUtility) shutdown(ctx context.Context, server shutdowner, shutdownTimeout time.Duration, shutdownChan <-chan os.Signal) (caught os.Signal, err error) {
	select {
	case sig := <-shutdownChan:
		caught = sig
	case <-ctx.Done():
		if cErr := ctx.Err(); cErr != nil {
			err = cErr
		}
	}

	shutdownCtx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return caught, fmt.Errorf("Shutdown: %w", err)
	}

	return caught, err
}
