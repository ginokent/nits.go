package nitstest

import (
	"bytes"
	"errors"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"testing"
)

// nolint: gochecknoglobals
var formatEqual = func() string {
	if useColor, _ := strconv.ParseBool(os.Getenv("GO_TEST_COLOR")); useColor {
		return "\n\033[31m--- noteql\033[0m\n\033[32m+++ actual\033[0m\n\033[31m-%v\033[0m\n\033[32m+%v\033[0m\n"
	}

	return "\n--- noteql\n+++ actual\n-%v\n+%v\n"
}()

// nolint: gochecknoglobals
var formatNotEqual = func() string {
	if useColor, _ := strconv.ParseBool(os.Getenv("GO_TEST_COLOR")); useColor {
		return "\n\033[31m--- expect\033[0m\n\033[32m+++ actual\033[0m\n\033[31m-%v\033[0m\n\033[32m+%v\033[0m\n"
	}

	return "\n--- expect\n+++ actual\n-%v\n+%v\n"
}()

func FailIfEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()

	if expect == actual {
		t.Errorf(formatEqual, expect, actual)
	}
}

func FailIfNotEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()

	if expect != actual {
		t.Errorf(formatNotEqual, expect, actual)
	}
}

func FailIfBytesEqual(t *testing.T, expect, actual []byte) {
	t.Helper()

	if bytes.Equal(expect, actual) {
		t.Errorf(formatEqual, string(expect), string(actual))
	}
}

func FailIfNotBytesEqual(t *testing.T, expect, actual []byte) {
	t.Helper()

	if !bytes.Equal(expect, actual) {
		t.Errorf(formatNotEqual, string(expect), string(actual))
	}
}

func FailIfDeepEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()

	if reflect.DeepEqual(expect, actual) {
		t.Errorf(formatEqual, expect, actual)
	}
}

func FailIfNotDeepEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf(formatNotEqual, expect, actual)
	}
}

func FailIfErrorIs(t *testing.T, expect, actual error) {
	t.Helper()

	if errors.Is(actual, expect) {
		t.Errorf(formatEqual, expect, actual)
	}
}

func FailIfNotErrorIs(t *testing.T, expect, actual error) {
	t.Helper()

	if !errors.Is(actual, expect) {
		t.Errorf(formatNotEqual, expect, actual)
	}
}

func FailIfRegexpMatchString(t *testing.T, expect *regexp.Regexp, actual string) {
	t.Helper()

	if expect.MatchString(actual) {
		t.Errorf(formatEqual, expect.String(), actual)
	}
}

func FailIfNotRegexpMatchString(t *testing.T, expect *regexp.Regexp, actual string) {
	t.Helper()

	if !expect.MatchString(actual) {
		t.Errorf(formatNotEqual, expect.String(), actual)
	}
}
