package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
)

type HttpError struct {
	*http.Response
	Message string
	err     error
}

func (r *HttpError) Unwrap() error {
	return r.err
}

func (r *HttpError) Error() string {
	return fmt.Sprintf("http %d: %s", r.StatusCode, r.Message)
}

// DefaultErrorMapper returns *HttpError
func DefaultErrorMapper(ctx context.Context, resp common.ResponseWrapper) error {
	if resp.Response.StatusCode < 400 {
		return nil
	}
	raw, err := io.ReadAll(resp.ReadCloser)
	if err != nil {
		return &HttpError{
			Response: resp.Response,
			Message:  "failed to read response",
			err:      err,
		}
	}
	return &HttpError{
		Response: resp.Response,
		Message:  string(raw),
	}
}

// ErrorRetrier determines whether a request should be retried. The request should be retried if
// and only if the function returns true.
type ErrorRetrier func(context.Context, *http.Request, *common.ResponseWrapper, error) bool

// DefaultErrorRetriable is the ErrorRetrier used if none is specified. It retries on 429 and 504 errors.
func DefaultErrorRetriable(ctx context.Context, req *http.Request, resp *common.ResponseWrapper, err error) bool {
	return CombineRetriers(
		RetryOnTooManyRequests,
		RetryOnGatewayTimeout,
		RetryUrlErrors,
	)(ctx, req, resp, err)
}

// RetryOnTooManyRequests retries when the response status code is 429.
func RetryOnTooManyRequests(ctx context.Context, _ *http.Request, resp *common.ResponseWrapper, err error) bool {
	if resp.Response == nil {
		return false
	}
	return resp.Response.StatusCode == http.StatusTooManyRequests
}

// RetryOnGatewayTimeout retries when the response status code is 504.
func RetryOnGatewayTimeout(ctx context.Context, _ *http.Request, resp *common.ResponseWrapper, err error) bool {
	if resp.Response == nil {
		return false
	}
	return resp.Response.StatusCode == http.StatusGatewayTimeout
}

// CombineRetriers combines multiple ErrorRetriers into a single ErrorRetrier. The combined ErrorRetrier
// will return true if any of the input ErrorRetriers return true.
func CombineRetriers(retriers ...ErrorRetrier) ErrorRetrier {
	return func(ctx context.Context, req *http.Request, resp *common.ResponseWrapper, err error) bool {
		for _, retrier := range retriers {
			if retrier(ctx, req, resp, err) {
				return true
			}
		}
		return false
	}
}

var urlErrorTransientErrorMessages = []string{
	"connection reset by peer",
	"TLS handshake timeout",
	"connection refused",
	"Unexpected error",
	"i/o timeout",
}

// RetryUrlErrors retries when the error is a *url.Error with a transient error message.
func RetryUrlErrors(ctx context.Context, _ *http.Request, _ *common.ResponseWrapper, err error) bool {
	var urlError *url.Error
	if !errors.As(err, &urlError) {
		return false
	}
	for _, msg := range urlErrorTransientErrorMessages {
		if strings.Contains(err.Error(), msg) {
			logger.Debugf(ctx, "Attempting retry because of IO error: %s", err)
			return true
		}
	}
	return false
}

// RetryTransientErrors retries when the error message contains any of the provided substrings.
func RetryTransientErrors(errors []string) ErrorRetrier {
	return func(ctx context.Context, _ *http.Request, _ *common.ResponseWrapper, err error) bool {
		message := err.Error()
		// Handle transient errors for retries
		for _, substring := range errors {
			if strings.Contains(message, substring) {
				logger.Debugf(ctx, "Attempting retry because of %#v", substring)
				return true
			}
		}
		return false
	}
}

// RestApiMatcher matches a request based on the HTTP method and path.
type RestApiMatcher struct {
	// Method is the HTTP method to match.
	Method string
	// Path is the regular expression to match the path.
	Path regexp.Regexp
}

// Matches returns true if the request matches the method and path.
func (m *RestApiMatcher) Matches(req *http.Request) bool {
	return req.Method == m.Method && m.Path.MatchString(req.URL.Path)
}

// RetryMatchedRequests applies a retrier that only applies to requests matching one of the provided matchers.
func RetryMatchedRequests(matchers []RestApiMatcher, retryer ErrorRetrier) ErrorRetrier {
	return func(ctx context.Context, r *http.Request, rw *common.ResponseWrapper, err error) bool {
		for _, m := range matchers {
			if m.Matches(r) && retryer(ctx, r, rw, err) {
				logger.Debugf(ctx, "Attempting retry because of gateway timeout")
				return true
			}
		}
		return false
	}
}
