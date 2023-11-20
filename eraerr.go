package eraerr

import (
	"fmt"
	"strings"
)

// WrapErr return wrapped error with msg, the format is
// msg: err.
// please note that we are using go1.17, which does not support wrap multiple
// error (multiple %w).
func WrapErr(err error, msg string) error {
	return fmt.Errorf("%s:: %w", msg, err)
}

// AddErrCause return wrapped error with msg, the format is
// err: msg.
// please note that we are using go1.17, which does not support wrap multiple
// error (multiple %w).
func AddErrCause(err error, msg string) error {
	return fmt.Errorf("%w:: %s", err, msg)
}

// GetErrCause return string from the deepest error.
func GetErrCause(err error) string {
	errTrace := strings.Split(err.Error(), ":: ")
	return errTrace[len(errTrace)-1]
}
