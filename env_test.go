package nits_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/newtstat/nits.go"
)

const (
	testEnvKey = "UTGO_ENV_TEST"
)

// nolint: paralleltest
func TestGetOrDefaultString(t *testing.T) {
	testEnvDefaultValue := "defaultValue"
	testEnvValue := "value"

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, testEnvValue)
		actual := nits.Env.GetOrDefaultString(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Error()
		}
	})

	t.Run("success(default)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual := nits.Env.GetOrDefaultString(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetOrDefaultBool(t *testing.T) {
	testEnvDefaultValue := false
	testEnvValue := true

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, strconv.FormatBool(testEnvValue))
		actual := nits.Env.GetOrDefaultBool(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Error()
		}
	})

	t.Run("success(default)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual := nits.Env.GetOrDefaultBool(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetOrDefaultInt64(t *testing.T) {
	const (
		testEnvDefaultValue = 1
		testEnvValue        = 2
	)

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, strconv.Itoa(testEnvValue))
		actual := nits.Env.GetOrDefaultInt64(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Error()
		}
	})

	t.Run("success(default)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual := nits.Env.GetOrDefaultInt64(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetOrDefaultSecond(t *testing.T) {
	testEnvDefaultValue := 1 * time.Second
	testEnvValue := 2 * time.Second

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, "2")
		actual := nits.Env.GetOrDefaultSecond(testEnvKey, testEnvDefaultValue)
		if actual != testEnvValue {
			t.Error()
		}
	})

	t.Run("success(default)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual := nits.Env.GetOrDefaultSecond(testEnvKey, testEnvDefaultValue)
		if actual != testEnvDefaultValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetString(t *testing.T) {
	testEnvSuccessValue := "value"

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, testEnvSuccessValue)
		actual, err := nits.Env.GetString(testEnvKey)
		if err != nil {
			t.Error(err, "!=", nil)
		}
		if actual != testEnvSuccessValue {
			t.Error()
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual, err := nits.Env.GetString(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != "" {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetBool(t *testing.T) {
	testEnvSuccessValue := true
	testEnvErrorValue := false

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, strconv.FormatBool(testEnvSuccessValue))
		actual, err := nits.Env.GetBool(testEnvKey)
		if err != nil {
			t.Error(err, "!=", nil)
		}
		if actual != testEnvSuccessValue {
			t.Error()
		}
	})

	t.Run("error(empty)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual, err := nits.Env.GetBool(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != testEnvErrorValue {
			t.Error()
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Setenv(testEnvKey, "error")
		actual, err := nits.Env.GetBool(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != testEnvErrorValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetInt64(t *testing.T) {
	const (
		testEnvSuccessValue = 1
		testEnvErrorValue   = 0
	)

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, strconv.Itoa(testEnvSuccessValue))
		actual, err := nits.Env.GetInt64(testEnvKey)
		if err != nil {
			t.Error(err, "!=", nil)
		}
		if actual != testEnvSuccessValue {
			t.Error()
		}
	})

	t.Run("error(empty)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual, err := nits.Env.GetInt64(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != testEnvErrorValue {
			t.Error()
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Setenv(testEnvKey, "error")
		actual, err := nits.Env.GetInt64(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != testEnvErrorValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetSecond(t *testing.T) {
	testEnvSuccessValue := 1 * time.Second
	testEnvErrorValue := 0 * time.Second

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, "1")
		actual, err := nits.Env.GetSecond(testEnvKey)
		if err != nil {
			t.Error(err, "!=", nil)
		}
		if actual != testEnvSuccessValue {
			t.Error()
		}
	})

	t.Run("error(empty)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual, err := nits.Env.GetSecond(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != testEnvErrorValue {
			t.Error()
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Setenv(testEnvKey, "error")
		actual, err := nits.Env.GetSecond(testEnvKey)
		if err == nil {
			t.Error()
		}
		if actual != testEnvErrorValue {
			t.Error()
		}
	})
}

// nolint: paralleltest
func TestGetCSV(t *testing.T) {
	testEnvSuccessValue := "a,,c"
	testEnvSuccessCSV := []string{"a", "", "c"}
	testEnvErrorValue := "\r\n"

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, testEnvSuccessValue)
		actual, err := nits.Env.GetCSV(testEnvKey)
		if err != nil {
			t.Error(err)
		}
		if !nits.Slice.EqualString(testEnvSuccessCSV, actual) {
			t.Error("!util.Slice.EqualString(testEnvSuccessCSV, actual)", testEnvSuccessCSV, "!=", actual)
		}
	})

	t.Run("error(empty)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual, err := nits.Env.GetCSV(testEnvKey)
		if err == nil {
			t.Error("err == nil")
		}
		if actual != nil {
			t.Error("actual != nil", actual, "!=", nil)
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Setenv(testEnvKey, testEnvErrorValue)
		actual, err := nits.Env.GetCSV(testEnvKey)
		if err == nil {
			t.Error("err == nil")
		}
		if actual != nil {
			t.Error("actual != nil", actual, "!=", nil)
		}
	})
}

// nolint: paralleltest
func TestGetCSVExcludeEmpty(t *testing.T) {
	testEnvSuccessValue := "a,,c"
	testEnvSuccessCSV := []string{"a", "c"}
	testEnvErrorValue := "\r\n"

	t.Run("success()", func(t *testing.T) {
		t.Setenv(testEnvKey, testEnvSuccessValue)
		actual, err := nits.Env.GetCSVExcludeEmptyString(testEnvKey)
		if err != nil {
			t.Error(err)
		}
		if !nits.Slice.EqualString(testEnvSuccessCSV, actual) {
			t.Error("!util.Slice.EqualString(testEnvSuccessCSV, actual)", testEnvSuccessCSV, actual)
		}
	})

	t.Run("error(empty)", func(t *testing.T) {
		t.Setenv(testEnvKey, "")
		actual, err := nits.Env.GetCSVExcludeEmptyString(testEnvKey)
		if err == nil {
			t.Error("err == nil")
		}
		if actual != nil {
			t.Error("actual != nil", actual, "!=", nil)
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Setenv(testEnvKey, testEnvErrorValue)
		actual, err := nits.Env.GetCSVExcludeEmptyString(testEnvKey)
		if err == nil {
			t.Error("err == nil")
		}
		if actual != nil {
			t.Error("actual != nil", actual, "!=", nil)
		}
	})
}
