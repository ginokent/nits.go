package nits

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// httpUtility is an empty structure that is prepared only for creating methods.
type httpUtility struct{}

// HTTP is an entity that allows the methods of HTTPUtility to be executed from outside the package without initializing HTTPUtility.
// nolint: gochecknoglobals
var HTTP httpUtility

// TerminateHandlerFunc returns http.HandlerFunc that sends SIGTERM to the shutdown signal.
func (httpUtility) TerminateHandlerFunc(preprocess http.HandlerFunc, shutdownChan chan<- os.Signal, signal os.Signal) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preprocess(w, r)
		shutdownChan <- signal
	}
}

// ListenAndServe will start *http.Server and ignore http.ErrServerClosed on shutdown.
// ListenAndServe expects to be used in conjunction with Shutdown:.
//
//	func startServer(ctx context.Context, server *http.Server) error {
//		shutdownChan := make(chan os.Signal, 1)
//		signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
//
//		// wait signal for shutdown
//		go func() {
//			sig, err := nits.HTTP.Shutdown(ctx, server, 5*time.Second, shutdownChan)
//			if err != nil {
//				log.Printf("Shutdown: signal=%s: %v", sig, err)
//			}
//		}()
//
//		// start server
//		if err := nits.HTTP.ListenAndServe(server); err != nil {
//			return fmt.Errorf("ListenAndServe: %w", err)
//		}
//
//		return nil
//	}
//
func (httpUtility) ListenAndServe(server *http.Server) error {
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
//				log.Printf("Shutdown: signal=%s: %v", sig, err)
//			}
//		}()
//
//		// start server
//		if err := nits.HTTP.ListenAndServe(server); err != nil {
//			return fmt.Errorf("ListenAndServe: %w", err)
//		}
//
//		return nil
//	}
//
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

func (httpUtility) AddMiddlewares(filoMiddlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		for i := range filoMiddlewares {
			handler = filoMiddlewares[i](handler)
		}

		return handler
	}
}

type (
	Method  = string
	Methods map[Method]http.Handler
)

func (httpUtility) NewMethodsHandler(methodNotAllowed http.Handler) (methods func(methods Methods) http.Handler, register func(Method, http.Handler) Methods) {
	return func(handlers Methods) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				handlerFunc, ok := handlers[strings.ToUpper(r.Method)]
				if !ok {
					methodNotAllowed.ServeHTTP(rw, r)

					return
				}

				handlerFunc.ServeHTTP(rw, r)
			})
		}, func(method Method, handler http.Handler) Methods {
			return Methods{method: handler}
		}
}

func (m Methods) Add(method Method, handler http.Handler) Methods {
	m[method] = handler

	return m
}

type (
	BasicAuthUsername = string
	BasicAuthPassword = string
)

func (httpUtility) BasicAuth(basicAuthUsers map[BasicAuthUsername]BasicAuthPassword) func(r *http.Request) bool {
	return func(r *http.Request) bool {
		requestUsername, requestPassword, ok := r.BasicAuth()
		if !ok {
			return false
		}

		for username, password := range basicAuthUsers {
			if requestUsername == username && requestPassword == password {
				return true
			}
		}

		return false
	}
}
