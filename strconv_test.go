package nits_test

import (
	"math"
	"testing"

	"github.com/nitpickers/nits.go"
)

func TestAtoi64(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MaxInt64
		actual, err := nits.Strconv.Atoi64("9223372036854775807")
		if err != nil {
			t.Error(err)
		}

		if expect != actual {
			t.Error("expect != actual", expect, actual)
		}
	})

	t.Run("error()", func(t *testing.T) {
		t.Parallel()

		_, err := nits.Strconv.Atoi64("9223372036854775808")
		if err == nil {
			t.Error("err == nil")
		}
	})
}

func TestI64toa(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = "9223372036854775807"

		actual := nits.Strconv.I64toa(math.MaxInt64)
		if expect != actual {
			t.Error("expect != actual", expect, actual)
		}
	})
}

func Test_strconvUtility_ParseBool(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args string
		want bool
	}{
		{"success()", "true", true},
		{"error(fail)", "fail", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := nits.Strconv.ParseBool(tt.args); got != tt.want {
				t.Errorf("strconvUtility.ParseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
