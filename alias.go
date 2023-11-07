package databricks

import (
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/databricks-sdk-go/version"
)

type Config config.Config

var (
	ErrInternal                = apierr.ErrInternal
	ErrDataLoss                = apierr.ErrDataLoss
	ErrUnknown                 = apierr.ErrUnknown
	ErrTemporarilyUnavailable  = apierr.ErrTemporarilyUnavailable
	ErrBadRequest              = apierr.ErrBadRequest
	ErrInvalidParameterValue   = apierr.ErrInvalidParameterValue
	ErrDeadlineExceeded        = apierr.ErrDeadlineExceeded
	ErrCancelled               = apierr.ErrCancelled
	ErrNotFound                = apierr.ErrNotFound
	ErrUnauthenticated         = apierr.ErrUnauthenticated
	ErrPermissionDenied        = apierr.ErrPermissionDenied
	ErrTooManyRequests         = apierr.ErrTooManyRequests
	ErrResourceExhausted       = apierr.ErrResourceExhausted
	ErrRequestLimitExceeded    = apierr.ErrRequestLimitExceeded
	ErrResourceConflict        = apierr.ErrResourceConflict
	ErrAlreadyExists           = apierr.ErrAlreadyExists
	ErrAborted                 = apierr.ErrAborted
	ErrOperationNotImplemented = apierr.ErrOperationNotImplemented
)

// Must panics if error is not nil. It's intended to be used with
// [databricks.NewWorkspaceClient] and [databricks.NewAccountClient].
func Must[T any](c T, err error) T {
	if err != nil {
		panic(err)
	}
	return c
}

// Version of this SDK
func Version() string {
	return version.Version
}

// WithProduct is expected to be set by developers to differentiate their app from others.
//
// Example setting is:
//
//	func init() {
//		databricks.WithProduct("your-product", "0.0.1")
//	}
func WithProduct(name, version string) {
	useragent.WithProduct(name, version)
}
