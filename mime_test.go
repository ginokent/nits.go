package nits_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/newtstat/nits.go"
)

type TestErrorReader struct {
	e string
}

func (t *TestErrorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("%s: p=%#v: %w", t.e, p, ErrTestError)
}

func TestDetectContentType(t *testing.T) {
	t.Parallel()

	t.Run("DetectContentType/success", func(t *testing.T) {
		t.Parallel()

		const expect = "text/html; charset=utf-8"
		actual, err := nits.MIME.DetectContentType(strings.NewReader("<!DOCTYPE html>"))
		if err != nil {
			t.Error(err)
		}
		if expect != actual {
			t.Error()
		}
	})

	t.Run("DetectContentType/error", func(t *testing.T) {
		t.Parallel()

		r := &TestErrorReader{e: "error"}
		if _, err := nits.MIME.DetectContentType(r); err == nil {
			t.Error()
		}
	})
}
