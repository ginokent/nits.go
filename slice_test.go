package nits_test

import (
	"testing"

	"github.com/nitpickers/nits.go"
)

func TestSliceContainsInt(t *testing.T) {
	t.Parallel()

	testSlice := []int{0, 1, 2}

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		if !nits.Slice.ContainsInt(testSlice, 0) {
			t.Fail()
		}
	})
	t.Run("false", func(t *testing.T) {
		t.Parallel()

		if nits.Slice.ContainsInt(testSlice, 3) {
			t.Fail()
		}
	})
}

func TestSliceContainsString(t *testing.T) {
	t.Parallel()

	testSlice := []string{"0", "1", "2"}

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		if !nits.Slice.ContainsString(testSlice, "0") {
			t.Fail()
		}
	})
	t.Run("false", func(t *testing.T) {
		t.Parallel()

		if nits.Slice.ContainsString(testSlice, "3") {
			t.Fail()
		}
	})
}

func TestSliceEqualInt(t *testing.T) {
	t.Parallel()

	testSlice := []int{0, 1, 2}
	testSliceEqual := []int{0, 1, 2}
	testSliceNotEqual1 := []int{1, 1, 1}
	testSliceNotEqual2 := []int{1, 1}

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		if !nits.Slice.EqualInt(testSlice, testSliceEqual) {
			t.Fail()
		}
	})
	t.Run("false1", func(t *testing.T) {
		t.Parallel()

		if nits.Slice.EqualInt(testSlice, testSliceNotEqual1) {
			t.Fail()
		}
	})
	t.Run("false2", func(t *testing.T) {
		t.Parallel()

		if nits.Slice.EqualInt(testSlice, testSliceNotEqual2) {
			t.Fail()
		}
	})
}

func TestSliceEqualString(t *testing.T) {
	t.Parallel()

	testSlice := []string{"0", "1", "2"}
	testSliceEqual := []string{"0", "1", "2"}
	testSliceNotEqual1 := []string{"1", "1", "1"}
	testSliceNotEqual2 := []string{"1", "1"}

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		if !nits.Slice.EqualString(testSlice, testSliceEqual) {
			t.Fail()
		}
	})
	t.Run("false1", func(t *testing.T) {
		t.Parallel()

		if nits.Slice.EqualString(testSlice, testSliceNotEqual1) {
			t.Fail()
		}
	})
	t.Run("false2", func(t *testing.T) {
		t.Parallel()

		if nits.Slice.EqualString(testSlice, testSliceNotEqual2) {
			t.Fail()
		}
	})
}

func TestSliceExcludeString(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		testSlice := []string{"0", "1", ""}
		expect := []string{"0", "1"}
		actual := nits.Slice.ExcludeString(testSlice, "")
		if !nits.Slice.EqualString(expect, actual) {
			t.Error("!util.Slice.EqualString(expect, actual)", expect, actual)
		}
	})
}
