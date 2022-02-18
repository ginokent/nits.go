package nits_test

import (
	"math"
	"testing"
	"time"

	"github.com/nitpickers/nits.go"
)

func TestPtrUtility_Bool(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = true
		actual := nits.Ptr.Bool(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Int(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MinInt
		actual := nits.Ptr.Int(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Int8(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MinInt8
		actual := nits.Ptr.Int8(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Int16(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MinInt16
		actual := nits.Ptr.Int16(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Int32(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MinInt32
		actual := nits.Ptr.Int32(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Int64(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MinInt64
		actual := nits.Ptr.Int64(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Uint(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MaxUint
		actual := nits.Ptr.Uint(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Uint8(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MaxUint8
		actual := nits.Ptr.Uint8(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Uint16(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MaxUint16
		actual := nits.Ptr.Uint16(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Uint32(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MaxUint32
		actual := nits.Ptr.Uint32(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Uint64(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = math.MaxUint64
		actual := nits.Ptr.Uint64(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Complex64(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect complex64 = math.Pi
		actual := nits.Ptr.Complex64(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Complex128(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect complex128 = math.Pi
		actual := nits.Ptr.Complex128(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Float32(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect float32 = math.Pi
		actual := nits.Ptr.Float32(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Float64(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect float64 = math.Pi
		actual := nits.Ptr.Float64(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_Time(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		expect := time.Unix(0, 0)
		actual := nits.Ptr.Time(expect)
		if !expect.Equal(*actual) {
			t.Error()
		}
	})
}

func TestPtrUtility_Duration(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		expect := 1 * time.Second
		actual := nits.Ptr.Duration(expect)
		if expect != *actual {
			t.Error()
		}
	})
}

func TestPtrUtility_String(t *testing.T) {
	t.Parallel()

	t.Run("success()", func(t *testing.T) {
		t.Parallel()

		const expect = "test string\n"
		actual := nits.Ptr.String(expect)
		if expect != *actual {
			t.Error()
		}
	})
}
