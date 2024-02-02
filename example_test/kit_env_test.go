package example_test

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

// fetchOsEnvBool
//
//	fetch os env by key.
//	if not found will return devValue.
//	return env not same as true (will be lowercase, so TRUE is same)
func fetchOsEnvBool(key string, devValue bool) bool {
	if os.Getenv(key) == "" {
		return devValue
	}
	return strings.ToLower(os.Getenv(key)) == "true"
}

// fetchOsEnvInt
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to int, return devValue
func fetchOsEnvInt(key string, devValue int) int {
	if os.Getenv(key) == "" {
		return devValue
	}
	outNum, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return devValue
	}

	return outNum
}

// fetchOsEnvStr
//
//	fetch os env by key.
//	return not found will return devValue.
func fetchOsEnvStr(key, devValue string) string {
	if os.Getenv(key) == "" {
		return devValue
	}
	return os.Getenv(key)
}

// fetchOsEnvInt
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
func fetchOsEnvArray(key string) []string {
	var devValueStr []string
	if os.Getenv(key) == "" {
		return devValueStr
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	if len(splitVal) == 0 {
		return devValueStr
	}
	for _, item := range splitVal {
		devValueStr = append(devValueStr, strings.TrimSpace(item))
	}

	return devValueStr
}

// setEnvStr
//
//	set env by key and val
func setEnvStr(t *testing.T, key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		t.Fatalf("set env key [%v] string err: %v", key, err)
	}
}

// setEnvBool
//
//	set env by key and val
//
//nolint:golint,unused
func setEnvBool(t *testing.T, key string, val bool) {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		t.Fatalf("set env key [%v] bool err: %v", key, err)
	}
}

// setEnvU64
//
//	set env by key and val
//
//nolint:golint,unused
func setEnvU64(t *testing.T, key string, val uint64) {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		t.Fatalf("set env key [%v] uint64 err: %v", key, err)
	}
}

// setEnvInt64
//
//	set env by key and val
//
//nolint:golint,unused
func setEnvInt64(t *testing.T, key string, val int64) {
	err := os.Setenv(key, strconv.FormatInt(val, 10))
	if err != nil {
		t.Fatalf("set env key [%v] int64 err: %v", key, err)
	}
}

// printEnvPrefix
//
//	print env by prefix
//
//nolint:golint,unused
func printEnvPrefix(t *testing.T, prefix string) {
	for _, e := range os.Environ() {
		if strings.Index(e, prefix) == 0 {
			t.Logf("printEnvPrefix [ %s ]: %s\n", prefix, e)
		}
	}
}
