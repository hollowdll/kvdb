package kvdb

import (
	"regexp"
	"strings"

	kvdberrors "github.com/hollowdll/kvdb/errors"
)

const (
	// DbNameMaxSize is the maximum length of database name in bytes.
	DbNameMaxSize int = 32
	// DbKeyMaxSize is the maximum length of database key in bytes.
	DbKeyMaxSize int = 1024
)

// isBlank returns true if input is blank.
func isBlank(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

// isTooLong returns true if input is longer than target.
// Target should be in bytes.
func isTooLong(input string, targetBytes int) bool {
	return len(input) > targetBytes
}

// dbNamecontainsValidCharacters checks if database name
// contains valid characters by matching it against a regexp.
func dbNamecontainsValidCharacters(name string) bool {
	pattern := regexp.MustCompile("^[A-Za-z0-9-_]+$")
	return pattern.MatchString(name)
}

// validateDatabaseName validates database name.
// Returns error if validation error is matched.
func validateDatabaseName(name string) error {
	if isBlank(name) {
		return kvdberrors.ErrDatabaseNameRequired
	}
	if isTooLong(name, DbNameMaxSize) {
		return kvdberrors.ErrDatabaseNameTooLong
	}
	if !dbNamecontainsValidCharacters(name) {
		return kvdberrors.ErrDatabaseNameInvalid
	}

	return nil
}

// validateDatabaseKey validates database key.
// Returns error if validation error is matched.
func validateDatabaseKey(key DatabaseKey) error {
	if isBlank(string(key)) {
		return kvdberrors.ErrDatabaseKeyRequired
	}
	if isTooLong(string(key), DbKeyMaxSize) {
		return kvdberrors.ErrDatabaseKeyTooLong
	}

	return nil
}
