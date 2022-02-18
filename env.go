package nits

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// envUtility is an empty structure that is prepared only for creating methods.
type envUtility struct{}

// Env is an entity that allows the methods of EnvUtility to be executed from outside the package without initializing EnvUtility.
// nolint: gochecknoglobals
var Env envUtility

// ErrEnvironmentVariableIsNotSetOrEmpty environment variable is not set or empty.
var ErrEnvironmentVariableIsNotSetOrEmpty = errors.New("environment variable is not set or empty")

// GetOrDefaultString returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (envUtility) GetOrDefaultString(env, defaultValue string) (value string) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return defaultValue
	}

	return valueString
}

// GetOrDefaultBool returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (envUtility) GetOrDefaultBool(env string, defaultValue bool) (value bool) {
	valueString := os.Getenv(env)

	v, err := strconv.ParseBool(valueString)
	if err != nil {
		return defaultValue
	}

	return v
}

// GetOrDefaultInt64 returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (envUtility) GetOrDefaultInt64(env string, defaultValue int64) (value int64) {
	valueString := os.Getenv(env)

	v, err := Strconv.Atoi64(valueString)
	if err != nil {
		return defaultValue
	}

	return v
}

// GetOrDefaultSecond returns the value of the environment variable `env` if it is set, or `defaultValue` if it is not set.
func (envUtility) GetOrDefaultSecond(env string, defaultValue time.Duration) (value time.Duration) {
	valueString := Env.GetOrDefaultInt64(env, -1)

	if valueString < 0 {
		return defaultValue
	}

	return time.Duration(valueString) * time.Second
}

// GetString returns the value of the environment variable `env` if it is set, or the error if it is not set.
func (envUtility) GetString(env string) (value string, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return "", fmt.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	return valueString, nil
}

// GetBool returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (envUtility) GetBool(env string) (value bool, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return false, fmt.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	v, err := strconv.ParseBool(valueString)
	if err != nil {
		return false, fmt.Errorf("%s: strconv.ParseBool: %w", env, err)
	}

	return v, nil
}

// GetInt64 returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (envUtility) GetInt64(env string) (value int64, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return 0, fmt.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	v, err := Strconv.Atoi64(valueString)
	if err != nil {
		return 0, fmt.Errorf("%s: strconv.Atoi: %w", env, err)
	}

	return v, nil
}

// GetSecond returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (envUtility) GetSecond(env string) (value time.Duration, err error) {
	valueString := os.Getenv(env)

	if valueString == "" {
		return 0, fmt.Errorf("%s: %w", env, ErrEnvironmentVariableIsNotSetOrEmpty)
	}

	v, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, fmt.Errorf("%s: strconv.Atoi: %w", env, err)
	}

	return time.Duration(v) * time.Second, nil
}

// GetCSV returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (envUtility) GetCSV(env string) (values []string, err error) {
	csvString, err := Env.GetString(env)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(csvString))

	values, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("reader.Read: %w", err)
	}

	return values, nil
}

// GetCSVExcludeEmptyString returns the value of the environment variable `env` if it is set, or the error if it is not set or invalid.
func (envUtility) GetCSVExcludeEmptyString(env string) (values []string, err error) {
	csv, err := Env.GetCSV(env)
	if err != nil {
		return nil, fmt.Errorf("GetCSV: %w", err)
	}

	return Slice.ExcludeString(csv, ""), nil
}
