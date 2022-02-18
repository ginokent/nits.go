// nolint: testpackage
package nits

import (
	"context"
	"io"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"
)

type testResponseWriter struct {
	err error
}

func (w *testResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *testResponseWriter) Write([]byte) (int, error) {
	return 0, w.err
}

func (w *testResponseWriter) WriteHeader(statusCode int) {}

func Test_httpUtility_Terminate(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name           string
		preprocess     func(w http.ResponseWriter, r *http.Request)
		shutdownSignal chan os.Signal
		want           os.Signal
	}{
		{"success(SIGTERM)", func(w http.ResponseWriter, r *http.Request) {}, make(chan os.Signal, 1), syscall.SIGTERM},
		{"success(SIGINT)", func(w http.ResponseWriter, r *http.Request) {}, make(chan os.Signal, 1), syscall.SIGINT},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			fn := HTTP.Terminate(testcase.preprocess, testcase.shutdownSignal, testcase.want)
			fn(&testResponseWriter{err: nil}, &http.Request{})
			if got := <-testcase.shutdownSignal; got != testcase.want {
				t.Errorf("got != tt.want: %v != %v", got, testcase.want)
			}
		})
	}
}

// nolint: paralleltest
func Test_httpUtility_Listen(t *testing.T) {
	t.Run("success(127.0.0.1:56789)", func(t *testing.T) {
		server := &http.Server{Addr: "127.0.0.1:56789"}
		errChan := make(chan error, 1)
		go func() {
			errChan <- HTTP.Listen(server)
		}()
		shutdownChan := make(chan os.Signal, 1)
		shutdownChan <- syscall.SIGTERM
		if _, err := HTTP.Shutdown(context.Background(), server, 5*time.Second, shutdownChan); err != nil {
			t.Errorf("err != nil: %v", err)
		}
		if err := <-errChan; err != nil {
			t.Errorf("err != nil: %v", err)
		}
	})

	t.Run("error(127.0.0.1:66666)", func(t *testing.T) {
		server := &http.Server{Addr: "127.0.0.1:66666"}
		if err := HTTP.Listen(server); err == nil {
			t.Error("err == nil")
		}
	})
}

// nolint: paralleltest
func Test_httpUtility_Shotdown(t *testing.T) {
	t.Run("success(127.0.0.1:56789)", func(t *testing.T) {
		server := &http.Server{Addr: "127.0.0.1:56789"}
		errChan := make(chan error, 1)
		go func() {
			errChan <- HTTP.Listen(server)
		}()
		shutdownChan := make(chan os.Signal, 1)
		shutdownChan <- syscall.SIGTERM
		if _, err := HTTP.Shutdown(context.Background(), server, 5*time.Second, shutdownChan); err != nil {
			t.Errorf("err != nil: %v", err)
		}
		if err := <-errChan; err != nil {
			t.Errorf("err != nil: %v", err)
		}
	})

	t.Run("error(cancel)", func(t *testing.T) {
		server := &http.Server{Addr: "127.0.0.1:56789"}
		shutdownChan := make(chan os.Signal, 1)
		shutdownCtx, cancel := context.WithCancel(context.Background())
		cancel()
		if _, err := HTTP.Shutdown(shutdownCtx, server, 0, shutdownChan); err == nil {
			t.Error("err == nil")
		}
	})

	t.Run("error(Shutdown)", func(t *testing.T) {
		server := &errShutdowner{err: io.EOF}
		shutdownChan := make(chan os.Signal, 1)
		shutdownChan <- syscall.SIGTERM
		if _, err := HTTP.shutdown(context.Background(), server, time.Nanosecond, shutdownChan); err == nil {
			t.Error("err == nil")
		}
	})
}

type errShutdowner struct {
	err error
}

func (e *errShutdowner) Shutdown(ctx context.Context) error {
	return e.err
}
