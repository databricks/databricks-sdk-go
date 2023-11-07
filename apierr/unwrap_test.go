package apierr

import (
	"errors"
	"os"
	"testing"
)

var testTable = []mapping{
	statusMapping(400, ErrBadRequest),
	codeMapping(400, "INVALID_PARAMETER_VALUE", ErrBadRequest),
	codeMapping(400, "INVALID_PARAMETER_VALUE", ErrInvalidParameterValue),
	codeMapping(400, "REQUEST_LIMIT_EXCEEDED", ErrTooManyRequests),
	statusMapping(401, ErrUnauthenticated),
	statusMapping(401, os.ErrPermission),
	statusMapping(403, ErrPermissionDenied),
	statusMapping(403, os.ErrPermission),
	statusMapping(404, ErrNotFound),
	statusMapping(404, os.ErrNotExist),
	statusMapping(409, ErrResourceConflict),
	codeMapping(409, "ABORTED", ErrAborted),
	codeMapping(409, "ABORTED", ErrResourceConflict),
	codeMapping(409, "ALREADY_EXISTS", ErrAlreadyExists),
	codeMapping(409, "ALREADY_EXISTS", ErrResourceConflict),
	statusMapping(429, ErrTooManyRequests),
	codeMapping(429, "REQUEST_LIMIT_EXCEEDED", ErrTooManyRequests),
	codeMapping(429, "REQUEST_LIMIT_EXCEEDED", ErrRequestLimitExceeded),
	codeMapping(429, "RESOURCE_EXHAUSTED", ErrTooManyRequests),
	codeMapping(429, "RESOURCE_EXHAUSTED", ErrResourceExhausted),
	statusMapping(499, ErrCancelled),
	statusMapping(500, ErrInternal),
	codeMapping(500, "UNKNOWN", ErrInternal),
	codeMapping(500, "UNKNOWN", ErrUnknown),
	codeMapping(500, "DATA_LOSS", ErrInternal),
	codeMapping(500, "DATA_LOSS", ErrDataLoss),
	statusMapping(501, ErrOperationNotImplemented),
	statusMapping(503, ErrTemporarilyUnavailable),
	statusMapping(504, ErrDeadlineExceeded),
	statusMapping(504, os.ErrDeadlineExceeded),
}

func TestErrorMapping(t *testing.T) {
	for _, m := range testTable {
		if !errors.Is(m.APIError, m.mapTo) {
			t.Errorf("error_code='%s', http_status='%d', expected to unwrap to `%s`",
				m.ErrorCode, m.StatusCode, m.mapTo)
		}
	}
}

type mapping struct {
	*APIError
	mapTo error
}

func statusMapping(statusCode int, err error) mapping {
	return mapping{
		APIError: &APIError{
			StatusCode: statusCode,
		},
		mapTo: err,
	}
}

func codeMapping(statusCode int, errorCode string, err error) mapping {
	return mapping{
		APIError: &APIError{
			StatusCode: statusCode,
			ErrorCode:  errorCode,
		},
		mapTo: err,
	}
}
