package nits

import "time"

// ptrUtility is an empty structure that is prepared only for creating methods.
type ptrUtility struct{}

// Ptr is an entity that allows the methods of PtrUtility to be executed from outside the package without initializing PtrUtility.
// nolint: gochecknoglobals
var Ptr ptrUtility

// Bool returns a pointer to the argument.
func (ptrUtility) Bool(value bool) *bool {
	return &value
}

// Int returns a pointer to the argument.
func (ptrUtility) Int(value int) *int {
	return &value
}

// Int8 returns a pointer to the argument.
func (ptrUtility) Int8(value int8) *int8 {
	return &value
}

// Int16 returns a pointer to the argument.
func (ptrUtility) Int16(value int16) *int16 {
	return &value
}

// Int32 returns a pointer to the argument.
func (ptrUtility) Int32(value int32) *int32 {
	return &value
}

// Int64 returns a pointer to the argument.
func (ptrUtility) Int64(value int64) *int64 {
	return &value
}

// Uint returns a pointer to the argument.
func (ptrUtility) Uint(value uint) *uint {
	return &value
}

// Uint8 returns a pointer to the argument.
func (ptrUtility) Uint8(value uint8) *uint8 {
	return &value
}

// Uint16 returns a pointer to the argument.
func (ptrUtility) Uint16(value uint16) *uint16 {
	return &value
}

// Uint32 returns a pointer to the argument.
func (ptrUtility) Uint32(value uint32) *uint32 {
	return &value
}

// Uint64 returns a pointer to the argument.
func (ptrUtility) Uint64(value uint64) *uint64 {
	return &value
}

// Complex64 returns a pointer to the argument.
func (ptrUtility) Complex64(value complex64) *complex64 {
	return &value
}

// Complex128 returns a pointer to the argument.
func (ptrUtility) Complex128(value complex128) *complex128 {
	return &value
}

// Float32 returns a pointer to the argument.
func (ptrUtility) Float32(value float32) *float32 {
	return &value
}

// Float64 returns a pointer to the argument.
func (ptrUtility) Float64(value float64) *float64 {
	return &value
}

// Time returns a pointer to the argument.
func (ptrUtility) Time(value time.Time) *time.Time {
	return &value
}

// Duration returns a pointer to the argument.
func (ptrUtility) Duration(value time.Duration) *time.Duration {
	return &value
}

// String returns a pointer to the argument.
func (ptrUtility) String(value string) *string {
	return &value
}
