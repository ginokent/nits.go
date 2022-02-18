package nits

import (
	"fmt"
	"strconv"
)

// strconvUtility is an empty structure that is prepared only for creating methods.
type strconvUtility struct{}

// Strconv is an entity that allows the methods of StrconvUtility to be executed from outside the package without initializing StrconvUtility.
// nolint: gochecknoglobals
var Strconv strconvUtility

// Atoi64 is equivalent to ParseInt(s, 10, 64).
func (strconvUtility) Atoi64(value string) (int64, error) {
	const (
		base    = 10
		bitSize = 64
	)

	v, err := strconv.ParseInt(value, base, bitSize)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt: %w", err)
	}

	return v, nil
}

// I64toa is equivalent to FormatInt(i, 10).
func (strconvUtility) I64toa(i int64) string {
	const base = 10

	return strconv.FormatInt(i, base)
}
