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

type ErrorRetryer func(context.Context, *http.Request, *common.ResponseWrapper, error) bool

func DefaultErrorRetriable(ctx context.Context, req *http.Request, resp *common.ResponseWrapper, err error) bool {
	return CombineRetriers(
		RetryOnTooManyRequests,
		RetryOnGatewayTimeout,
		RetryUrlErrors,
	)(ctx, req, resp, err)
}

func RetryOnTooManyRequests(ctx context.Context, _ *http.Request, resp *common.ResponseWrapper, err error) bool {
	if resp.Response == nil {
		return false
	}
	return resp.Response.StatusCode == http.StatusTooManyRequests
}

func RetryOnGatewayTimeout(ctx context.Context, _ *http.Request, resp *common.ResponseWrapper, err error) bool {
	if resp.Response == nil {
		return false
	}
	return resp.Response.StatusCode == http.StatusGatewayTimeout
}

func CombineRetriers(retriers ...ErrorRetryer) ErrorRetryer {
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

func RetryTransientErrors(errors []string) ErrorRetryer {
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

type RestApiMatcher struct {
	Method string
	Path   regexp.Regexp
}

func (m *RestApiMatcher) Matches(req *http.Request) bool {
	return req.Method == m.Method && m.Path.MatchString(req.URL.Path)
}

func shouldRetryIdempotentRequest(ctx context.Context, req *http.Request, resp *common.ResponseWrapper, err error) bool {
	return RetryOnGatewayTimeout(ctx, req, resp, err)
}

func RetryIdempotentRequests(methods []RestApiMatcher) ErrorRetryer {
	return func(ctx context.Context, r *http.Request, rw *common.ResponseWrapper, err error) bool {
		for _, m := range methods {
			if m.Matches(r) && shouldRetryIdempotentRequest(ctx, r, rw, err) {
				logger.Debugf(ctx, "Attempting retry because of gateway timeout")
				return true
			}
		}
		return false
	}
}
