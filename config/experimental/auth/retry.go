package auth

import (
	"context"
	"errors"
	"net"
	"net/http"
	"slices"
	"syscall"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/api"
	"golang.org/x/oauth2"
)

// retryingTokenSource wraps a TokenSource with retry logic for transient
// failures during token acquisition. Each retry calls the underlying Token()
// method, which creates a fresh HTTP request -- avoiding the body-rewinding
// problems that occur when retrying at the transport level.
type retryingTokenSource struct {
	inner TokenSource
	opts  []api.Option
}

var defaultOptions = []api.Option{
	api.WithTimeout(1 * time.Minute),
	api.WithRetrier(func() api.Retrier {
		return api.RetryOn(api.BackoffPolicy{}, isRetriableTokenError)
	}),
}

// NewRetryingTokenSource wraps inner with retry logic for transient failures.
// By default it uses exponential backoff with a 1-minute timeout and a 30-second maximum delay.
// Additional api.Option values can be provided to override the defaults.
func NewRetryingTokenSource(inner TokenSource, opts ...api.Option) TokenSource {
	return &retryingTokenSource{
		inner: inner,
		opts:  append(defaultOptions, opts...),
	}
}

// Token returns a token from the underlying source, retrying on transient
// errors.
func (r *retryingTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	return api.ExecuteWithResult(ctx, r.inner.Token, r.opts...)
}

// httpStatusCoder is implemented by errors that carry an HTTP status code.
// This interface avoids importing httpclient (which would create a cycle)
// while still allowing to classify httpclient.HttpError by status code.
type httpStatusCoder interface {
	HTTPStatusCode() int
}

var retriableCodes = []int{
	http.StatusTooManyRequests,    // 429
	http.StatusBadGateway,         // 502
	http.StatusServiceUnavailable, // 503
	http.StatusGatewayTimeout,     // 504
}

// isRetriableTokenError returns true if the error is a transient failure that
// should be retried. This covers HTTP errors from the SDK's transport layer,
// OAuth2 token endpoint errors, and transient network errors.
func isRetriableTokenError(err error) bool {
	if code := httpStatusCode(err); code != 0 {
		return slices.Contains(retriableCodes, code)
	}
	return isTransientNetworkError(err)
}

// httpStatusCode extracts the HTTP status code from an error, if available.
func httpStatusCode(err error) int {
	// Check oauth2.RetrieveError (has Response field with StatusCode).
	var retrieveErr *oauth2.RetrieveError
	if errors.As(err, &retrieveErr) && retrieveErr.Response != nil {
		return retrieveErr.Response.StatusCode
	}
	// Check for any error that exposes a StatusCode() method.
	var sc httpStatusCoder
	if errors.As(err, &sc) {
		return sc.HTTPStatusCode()
	}
	return 0
}

// isTransientNetworkError reports whether err represents a transient network
// condition that is likely to resolve on retry.
func isTransientNetworkError(err error) bool {
	if errors.Is(err, syscall.ECONNRESET) {
		return true
	}
	if errors.Is(err, syscall.ECONNREFUSED) {
		return true
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}
	return false
}
