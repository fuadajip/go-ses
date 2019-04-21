package error

import (
	"github.com/pkg/errors"
)

type (
	mywishesError struct {
		Code      int
		ErrorCode string
		Message   string
		Status    string
	}
)

// Error returns error type as a string
func (m *mywishesError) Error() string {
	return m.Message
}

// New returns new error message in standard pkg errors new
func New(msg string) error {
	return errors.New(msg)

}

// Wrap returns a new error that adds context to the original
func Wrap(code int, errorCode string, err error, msg string, status string) error {
	return errors.Wrap(&mywishesError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   msg,
		Status:    status,
	}, err.Error())
}
