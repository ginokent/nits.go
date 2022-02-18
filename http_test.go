// nolint: testpackage
package nits

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/newtstat/nits.go/nitstest"
)

const testAddr = "127.0.0.1:56789"

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

			fn := HTTP.TerminateHandlerFunc(testcase.preprocess, testcase.shutdownSignal, testcase.want)
			fn(&testResponseWriter{err: nil}, &http.Request{})
			if got := <-testcase.shutdownSignal; got != testcase.want {
				t.Errorf("got != tt.want: %v != %v", got, testcase.want)
			}
		})
	}
}

// nolint: paralleltest
func Test_httpUtility_Listen(t *testing.T) {
	t.Run("success("+testAddr+")", func(t *testing.T) {
		server := &http.Server{Addr: testAddr}
		errChan := make(chan error, 1)
		go func() {
			errChan <- HTTP.ListenAndServe(server)
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
		if err := HTTP.ListenAndServe(server); err == nil {
			t.Error("err == nil")
		}
	})
}

// nolint: paralleltest
func Test_httpUtility_Shotdown(t *testing.T) {
	t.Run("success("+testAddr+")", func(t *testing.T) {
		server := &http.Server{Addr: testAddr}
		errChan := make(chan error, 1)
		go func() {
			errChan <- HTTP.ListenAndServe(server)
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
		server := &http.Server{Addr: testAddr}
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

// nolint: paralleltest
func Test_httpUtility_AddMiddlewares(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		router := http.NewServeMux()
		router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			if _, err := rw.Write([]byte("handler")); err != nil {
				t.Errorf("rw.Write: %v", err)
			}
		})

		middleware := func(num int) func(http.Handler) http.Handler {
			return func(original http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					if _, err := rw.Write([]byte(fmt.Sprintf("middleware %d preprocess", num))); err != nil {
						t.Errorf("rw.Write: %v", err)
					}

					original.ServeHTTP(rw, r)

					if _, err := rw.Write([]byte(fmt.Sprintf("middleware %d postprocess", num))); err != nil {
						t.Errorf("rw.Write: %v", err)
					}
				})
			}
		}

		r := HTTP.AddMiddlewares(middleware(1), middleware(2))(router)

		server := &http.Server{
			Handler: r,
			Addr:    testAddr,
		}

		go func() {
			if err := HTTP.ListenAndServe(server); err != nil {
				t.Errorf("listen: %v", err)
			}
		}()

		resp, err := http.Get(fmt.Sprintf("http://%s", testAddr))
		if err != nil {
			t.Errorf("err != nil: %v", err)
			t.FailNow()
		}
		defer func() { _ = resp.Body.Close() }()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, resp.Body); err != nil {
			t.Errorf("err != nil: %v", err)
		}

		shutdownChan := make(chan os.Signal, 1)
		shutdownChan <- syscall.SIGTERM
		if _, err := HTTP.Shutdown(context.Background(), server, 1*time.Second, shutdownChan); err != nil {
			t.Errorf("err != nil: %v", err)
		}

		const expect = "middleware 2 preprocess" +
			"middleware 1 preprocess" +
			"handler" +
			"middleware 1 postprocess" +
			"middleware 2 postprocess"
		nitstest.FailIfNotEqual(t, expect, buf.String())
	})

	t.Run("test", func(t *testing.T) {
		router := http.NewServeMux()
		router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			if _, err := rw.Write([]byte("handler")); err != nil {
				t.Errorf("rw.Write: %v", err)
			}
		})

		middleware := func(num int) func(http.Handler) http.Handler {
			return func(original http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					if _, err := rw.Write([]byte(fmt.Sprintf("middleware %d preprocess", num))); err != nil {
						t.Errorf("rw.Write: %v", err)
					}

					original.ServeHTTP(rw, r)

					if _, err := rw.Write([]byte(fmt.Sprintf("middleware %d postprocess", num))); err != nil {
						t.Errorf("rw.Write: %v", err)
					}
				})
			}
		}

		r := HTTP.AddMiddlewares(middleware(1))(router)
		r = HTTP.AddMiddlewares(middleware(2))(r)

		server := &http.Server{
			Handler: r,
			Addr:    testAddr,
		}

		go func() {
			if err := HTTP.ListenAndServe(server); err != nil {
				t.Errorf("listen: %v", err)
			}
		}()

		resp, err := http.Get(fmt.Sprintf("http://%s", testAddr))
		if err != nil {
			t.Errorf("err != nil: %v", err)
		}
		defer func() { _ = resp.Body.Close() }()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, resp.Body); err != nil {
			t.Errorf("err != nil: %v", err)
		}

		shutdownChan := make(chan os.Signal, 1)
		shutdownChan <- syscall.SIGTERM
		if _, err := HTTP.Shutdown(context.Background(), server, 1*time.Second, shutdownChan); err != nil {
			t.Errorf("err != nil: %v", err)
		}

		const expect = "middleware 2 preprocess" +
			"middleware 1 preprocess" +
			"handler" +
			"middleware 1 postprocess" +
			"middleware 2 postprocess"
		nitstest.FailIfNotEqual(t, expect, buf.String())
	})
}

// nolint: paralleltest
func Test_httpUtility_HandleMethods(t *testing.T) {
	tests := []struct {
		name             string
		requestMethod    string
		methodNotAllowed http.HandlerFunc
		method           string
		handler          http.Handler
		expect           string
	}{
		{"success(GET,200)", "GET", http.NotFound, "GET", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if _, err := rw.Write([]byte("OK")); err != nil {
				t.Errorf("rw.Write: %v", err)
			}
		}), "OK"},
		{"success(POST,404)", "POST", http.NotFound, "GET", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if _, err := rw.Write([]byte("OK")); err != nil {
				t.Errorf("rw.Write: %v", err)
			}
		}), "404 page not found\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			methods, add := HTTP.NewMethodsHandler(tt.methodNotAllowed)

			router := http.NewServeMux()
			router.Handle("/", methods(add(tt.method, tt.handler).Register(tt.method, tt.handler)))

			server := &http.Server{
				Handler: router,
				Addr:    testAddr,
			}

			go func() {
				if err := HTTP.ListenAndServe(server); err != nil {
					t.Errorf("listen: %v", err)
				}
			}()

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, tt.requestMethod, fmt.Sprintf("http://%s", testAddr), http.NoBody)
			if err != nil {
				t.Errorf("err != nil: %v", err)
			}
			client := new(http.Client)
			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("err != nil: %v", err)
			}
			defer func() { _ = resp.Body.Close() }()

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				t.Errorf("err != nil: %v", err)
			}

			shutdownChan := make(chan os.Signal, 1)
			shutdownChan <- syscall.SIGTERM
			if _, err := HTTP.Shutdown(context.Background(), server, 1*time.Second, shutdownChan); err != nil {
				t.Errorf("err != nil: %v", err)
			}

			nitstest.FailIfNotEqual(t, tt.expect, buf.String())
		})
	}
}

func Test_httpUtility_BasicAuth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		basicAuthUsers map[BasicAuthUsername]BasicAuthPassword
		request        *http.Request
		expect         bool
	}{
		{"success()", map[BasicAuthUsername]BasicAuthPassword{"user": "pass"}, &http.Request{Header: map[string][]string{"Authorization": {"Basic dXNlcjpwYXNz"}}}, true},
		{"failure(Authorization:null)", map[BasicAuthUsername]BasicAuthPassword{"user": "pass"}, &http.Request{Header: map[string][]string{"Authorization": {""}}}, false},
		{"failure(user:failure)", map[BasicAuthUsername]BasicAuthPassword{"user": "failure"}, &http.Request{Header: map[string][]string{"Authorization": {"Basic dXNlcjpwYXNz"}}}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := HTTP.BasicAuth(tt.basicAuthUsers)(tt.request)
			if actual != tt.expect {
				nitstest.FailIfNotEqual(t, tt.expect, actual)
			}
		})
	}
}
