package nits_test

import (
	"testing"

	"github.com/newtstat/nits.go"
)

type testJSONStruct struct {
	Test string `json:"test"`
}

const (
	testString     = "test"
	testDataString = `{"test":"test"}`
)

func TestMustMarshal(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		data := nits.JSON.MustMarshal(&testJSONStruct{testString})

		if string(data) != testDataString {
			t.Error("string(data) != TestString", string(data))
		}
	})

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		func() { // FOR panic()
			defer func() { _ = recover() }()
			_ = nits.JSON.MustMarshal(func() {})
		}()
	})
}

func TestMustUnmarshal(t *testing.T) {
	t.Parallel()

	testData := []byte(testDataString)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		var actual testJSONStruct
		nits.JSON.MustUnmarshal(testData, &actual)

		if actual.Test != testString {
			t.Error("actual.Test != TestString", actual.Test)
		}
	})

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		func() { // FOR panic()
			defer func() { _ = recover() }()
			nits.JSON.MustUnmarshal(testData, func() {})
		}()
	})
}
