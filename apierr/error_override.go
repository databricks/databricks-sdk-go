package apierr

import (
	"context"
	"net/http"
	"regexp"
	"strconv"

	"github.com/databricks/databricks-sdk-go/logger"
)

// errorOverride is a struct that allows for customizing error responses from
// the API. This is necessary to ensure consistency between different APIs that
// currently have different error behavior.
//
// The actual overrides themselves are autogenerated in error_overrides.go.
type errorOverride struct {
	// The name of the override. This is used for logging purposes.
	debugName string

	// pathRegex is a substring that must be present in the request path for
	// this override to be applied.
	pathRegex *regexp.Regexp

	// verb is the HTTP verb that must be used for this override to be applied.
	verb string

	// statusCodeMatcher is a regular expression that must match the status code
	// for this override to be applied. If nil, this field is ignored.
	statusCodeMatcher *regexp.Regexp

	// errorCodeMatcher is a regular expression that must match the error code
	// for this override to be applied. If nil, this field is ignored.
	errorCodeMatcher *regexp.Regexp

	// messageMatcher is a regular expression that must match the error message
	// for this override to be applied. If nil, this field is ignored.
	messageMatcher *regexp.Regexp

	// customError is the error that will be returned by APIError.Unwrap() if
	// this override is applied.
	customError error
}

func (e *errorOverride) matches(err *APIError, resp *http.Response) bool {
	if e.verb != resp.Request.Method {
		return false
	}
	if e.pathRegex != nil && !e.pathRegex.MatchString(resp.Request.URL.Path) {
		return false
	}
	statusCode := strconv.Itoa(resp.StatusCode)
	if e.statusCodeMatcher != nil && !e.statusCodeMatcher.MatchString(statusCode) {
		return false
	}
	if e.errorCodeMatcher != nil && !e.errorCodeMatcher.MatchString(err.ErrorCode) {
		return false
	}
	if e.messageMatcher != nil && !e.messageMatcher.MatchString(err.Message) {
		return false
	}
	return true
}

func applyOverrides(ctx context.Context, err *APIError, resp *http.Response) {
	for _, override := range allOverrides {
		if override.matches(err, resp) {
			logger.Debugf(ctx, "Applying error override: %s", override.debugName)
			err.unwrap = override.customError
			return
		}
	}
}
