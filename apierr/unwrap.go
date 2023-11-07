package apierr

import (
	"context"
	"errors"
	"os"
)

var (
	ErrInternal                = errors.New("some invariants expected by the underlying system have been broken")
	ErrDataLoss                = inheritErr(ErrInternal, "unrecoverable data loss or corruption")
	ErrUnknown                 = inheritErr(ErrInternal, "this error is used as a fallback if the platform-side mapping is missing some reason")
	ErrTemporarilyUnavailable  = errors.New("the service is currently unavailable")
	ErrBadRequest              = inheritErr(os.ErrInvalid, "the request is invalid")
	ErrInvalidParameterValue   = inheritErr(ErrBadRequest, "supplied value for a parameter was invalid")
	ErrDeadlineExceeded        = inheritErr(os.ErrDeadlineExceeded, "the deadline expired before the operation could complete")
	ErrCancelled               = inheritErr(context.Canceled, "the operation was explicitly canceled by the caller")
	ErrNotFound                = inheritErr(os.ErrNotExist, "the operation was performed on a resource that does not exist")
	ErrUnauthenticated         = inheritErr(os.ErrPermission, "the request does not have valid authentication (AuthN) credentials for the operation")
	ErrPermissionDenied        = inheritErr(os.ErrPermission, "the caller does not have permission to execute the specified operation")
	ErrTooManyRequests         = errors.New("maps to HTTP code: 429 Too Many Requests")
	ErrResourceExhausted       = inheritErr(ErrTooManyRequests, "operation is rejected due to per-user rate limiting")
	ErrRequestLimitExceeded    = inheritErr(ErrTooManyRequests, "cluster request was rejected because it would exceed a resource limit")
	ErrResourceConflict        = inheritErr(os.ErrExist, "maps to all HTTP 409 (Conflict) responses")
	ErrAlreadyExists           = inheritErr(ErrResourceConflict, "operation was rejected due a conflict with an existing resource")
	ErrAborted                 = inheritErr(ErrResourceConflict, "the operation was aborted, typically due to a concurrency issue such as a sequencer check failure")
	ErrOperationNotImplemented = errors.New("the operation is not implemented or is not supported/enabled in this service")

	statusCodeMapping = map[int]error{
		400: ErrBadRequest,
		401: ErrUnauthenticated,
		403: ErrPermissionDenied,
		404: ErrNotFound,
		409: ErrResourceConflict,
		429: ErrTooManyRequests,
		499: ErrCancelled,
		500: ErrInternal,
		501: ErrOperationNotImplemented,
		503: ErrTemporarilyUnavailable,
		504: ErrDeadlineExceeded,
	}

	errorCodeMapping = map[string]error{
		// HTTP 400 variants
		"INVALID_PARAMETER_VALUE": ErrInvalidParameterValue,

		// HTTP 409 variants
		"ABORTED":                 ErrAborted,
		"ALREADY_EXISTS":          ErrAlreadyExists,
		"RESOURCE_ALREADY_EXISTS": ErrAlreadyExists,

		// HTTP 429 variants
		"RESOURCE_EXHAUSTED":     ErrResourceExhausted,
		"REQUEST_LIMIT_EXCEEDED": ErrRequestLimitExceeded,

		// HTTP 500 variants
		"UNKNOWN":   ErrUnknown,
		"DATA_LOSS": ErrDataLoss,
	}
)

func inheritErr(err error, message string) *wrapError {
	return &wrapError{
		wrap:    err,
		message: message,
	}
}

type wrapError struct {
	wrap    error
	message string
}

func (e *wrapError) Error() string {
	return e.message
}

func (e *wrapError) Unwrap() error {
	return e.wrap
}

// Unwrap error for easier client code checking
//
// See https://pkg.go.dev/errors#example-Unwrap
func (apiError *APIError) Unwrap() error {
	byErrorCode, ok := errorCodeMapping[apiError.ErrorCode]
	if ok {
		return byErrorCode
	}
	byStatusCode, ok := statusCodeMapping[apiError.StatusCode]
	if ok {
		return byStatusCode
	}
	// A nil error returned from e.Unwrap() indicates that e does not wrap
	// any error.
	return nil
}
