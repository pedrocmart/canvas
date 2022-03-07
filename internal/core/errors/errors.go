package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// Error internal app errors.
type Error struct {
	Code   int
	Status int
	Err    error
}

func (t Error) Error() string {
	return t.Err.Error()
}

// GetStatus return error status.
func (t Error) GetStatus() int {
	return t.Status
}

// GetCode return error status.
func (t Error) GetCode() int {
	return t.Code
}

// Wrap error.
func Wrap(format string, args ...interface{}) error {
	return Error{
		Status: http.StatusInternalServerError,
		Err:    errors.Errorf(format, args...),
	}
}

// Wrapf error.
func Wrapf(err error, format string, args ...interface{}) error {
	status := http.StatusInternalServerError
	code := 0

	e, ok := err.(Error)
	if ok {
		status = e.GetStatus()
		code = e.GetCode()
	}

	return Error{
		Code:   code,
		Status: status,
		Err:    errors.Wrapf(err, format, args...),
	}
}

// Internal error.
func Internal(format string, args ...interface{}) error {
	return Error{
		Status: http.StatusInternalServerError,
		Err:    errors.Errorf(format, args...),
	}
}

// InternalWrap error wrap.
func InternalWrap(err error, format string, args ...interface{}) error {
	return Error{
		Status: http.StatusInternalServerError,
		Err:    errors.Wrapf(err, format, args...),
	}
}

// BadRequest error.
func BadRequest(format string, args ...interface{}) error {
	return Error{
		Status: http.StatusBadRequest,
		Err:    errors.Errorf(format, args...),
	}
}

// BadRequestWrap error wrap.
func BadRequestWrap(err error, format string, args ...interface{}) error {
	return Error{
		Status: http.StatusBadRequest,
		Err:    errors.Wrapf(err, format, args...),
	}
}

// NotFound error.
func NotFound(format string, args ...interface{}) error {
	return Error{
		Status: http.StatusNotFound,
		Err:    errors.Errorf(format, args...),
	}
}

// NotFoundWrap error wrap.
func NotFoundWrap(err error, format string, args ...interface{}) error {
	return Error{
		Status: http.StatusNotFound,
		Err:    errors.Wrapf(err, format, args...),
	}
}

// Unauthorized error.
func Unauthorized(format string, args ...interface{}) error {
	return Error{
		Status: http.StatusUnauthorized,
		Err:    errors.Errorf(format, args...),
	}
}

// UnauthorizedWrap error wrap.
func UnauthorizedWrap(err error, format string, args ...interface{}) error {
	return Error{
		Status: http.StatusUnauthorized,
		Err:    errors.Wrapf(err, format, args...),
	}
}

// MethodNotAllowed error.
func MethodNotAllowed(format string, args ...interface{}) error {
	return Error{
		Status: http.StatusMethodNotAllowed,
		Err:    errors.Errorf(format, args...),
	}
}

// MethodNotAllowedWrap error wrap.
func MethodNotAllowedWrap(err error, format string, args ...interface{}) error {
	return Error{
		Status: http.StatusMethodNotAllowed,
		Err:    errors.Wrapf(err, format, args...),
	}
}

func WrapCode(err error, code int) error {
	e, ok := err.(Error)
	if ok {
		e.Code = code

		return e
	}

	return Error{
		Code:   code,
		Status: http.StatusInternalServerError,
		Err:    err,
	}
}
