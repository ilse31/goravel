package response

import (
	ierr "goravel/app/helpers/constants"
	"net/http"

	ctx "github.com/goravel/framework/contracts/http"

	"github.com/pkg/errors"
)

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	HTTPCode  int    `json:"-"`
	Success   bool   `json:"success" example:"false"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code,omitempty"`
	Internal  error  `json:"-"`
}

// Render implements http.Response.
func (e ErrorResponse) Render() error {
	panic("unimplemented")
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by CustomHTTPErrorHandler
func (e ErrorResponse) StatusCode() int {
	return e.HTTPCode
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrInternalServerError creates a new error response representing an internal server error (HTTP 500)
func ErrInternalServerError(err error, net ctx.Context) ctx.Response {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(ierr.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = ierr.ErrInternal.Code
		errorMessage = ierr.ErrInternal.Message
	}

	return net.Response().Json(http.StatusInternalServerError, ErrorResponse{
		HTTPCode:  http.StatusInternalServerError,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	})
}

func ErrUnauthorized(err error, net ctx.Context) ctx.Response {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(ierr.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = ierr.ErrUnauthorized.Code
		errorMessage = ierr.ErrUnauthorized.Message
	}

	return net.Response().Json(http.StatusUnauthorized, ErrorResponse{
		HTTPCode:  http.StatusUnauthorized,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	})
}

// ErrForbidden creates a new error response representing an authorization failure (HTTP 403)
func ErrForbidden(err error, net ctx.Context) ctx.Response {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(ierr.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = ierr.ErrForbidden.Code
		errorMessage = ierr.ErrForbidden.Message
	}

	return net.Response().Json(http.StatusForbidden, ErrorResponse{
		HTTPCode:  http.StatusForbidden,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	})
}

// ErrNotFound creates a new error response representing a resource not found (HTTP 404)
func ErrNotFound(err error, net ctx.Context) ctx.Response {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(ierr.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = ierr.ErrResourceNotFound.Code
		errorMessage = ierr.ErrResourceNotFound.Message
	}

	return net.Response().Json(http.StatusNotFound, ErrorResponse{
		HTTPCode:  http.StatusNotFound,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	})
}

// ErrBadRequest creates a new error response representing a bad request (HTTP 400)
func ErrBadRequest(err error, net ctx.Context) ctx.Response {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	originalErr := errors.Cause(err)
	var errorCode string
	var errorMessage string

	if val, ok := originalErr.(ierr.Error); ok {
		errorCode = val.Code
		errorMessage = val.Message
	} else {
		errorCode = ierr.ErrBadRequest.Code
		errorMessage = ierr.ErrBadRequest.Message
	}

	return net.Response().Json(http.StatusBadRequest, ErrorResponse{
		HTTPCode:  http.StatusBadRequest,
		Message:   errorMessage,
		ErrorCode: errorCode,
		Internal:  err,
	})
}

func HTTPError(err error, statusCode int, errorCode string, message string) ErrorResponse {

	if _, ok := err.(stackTracer); !ok {
		err = errors.WithStack(err)
	}

	return ErrorResponse{
		HTTPCode:  statusCode,
		Message:   message,
		ErrorCode: errorCode,
		Internal:  err,
	}
}
