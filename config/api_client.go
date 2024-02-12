package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func (c *Config) NewApiClient() (*httpclient.ApiClient, error) {
	if skippable, ok := c.HTTPTransport.(interface {
		SkipRetryOnIO() bool
	}); ok && skippable.SkipRetryOnIO() {
		c.Loaders = []Loader{noopLoader{}}
		c.Credentials = noopAuth{}
	}
	err := c.EnsureResolved()
	if err != nil {
		return nil, err
	}
	retryTimeout := time.Duration(orDefault(c.RetryTimeoutSeconds, 300)) * time.Second
	httpTimeout := time.Duration(orDefault(c.HTTPTimeoutSeconds, 60)) * time.Second
	return httpclient.NewApiClient(httpclient.ClientConfig{
		RetryTimeout:       retryTimeout,
		HTTPTimeout:        httpTimeout,
		RateLimitPerSecond: orDefault(c.RateLimitPerSecond, 15),
		DebugHeaders:       c.DebugHeaders,
		DebugTruncateBytes: c.DebugTruncateBytes,
		InsecureSkipVerify: c.InsecureSkipVerify,
		Transport:          c.HTTPTransport,
		Visitors: []httpclient.RequestVisitor{
			c.Authenticate,
			func(r *http.Request) error {
				if r.URL == nil {
					return fmt.Errorf("no URL found in request")
				}
				url, err := url.Parse(c.Host)
				if err != nil {
					return err
				}
				r.URL.Host = url.Host
				r.URL.Scheme = url.Scheme
				return nil
			},
			func(r *http.Request) error {
				ctx := useragent.InContext(r.Context(), "auth", c.AuthType)
				*r = *r.WithContext(ctx) // replace request
				return nil
			},
			func(r *http.Request) error {
				// Detect if we are running in a CI/CD environment
				provider := useragent.CiCdProvider()
				if provider == "" {
					return nil
				}
				// Add the detected CI/CD provider to the user agent
				ctx := useragent.InContext(r.Context(), "cicd", provider)
				*r = *r.WithContext(ctx) // replace request
				return nil
			},
		},
		TransientErrors: []string{
			"REQUEST_LIMIT_EXCEEDED", // This is temporary workaround for SCIM API returning 500.  Remove when it's fixed
		},
		ErrorMapper: apierr.GetAPIError,
		ErrorRetriable: func(ctx context.Context, err error) bool {
			var apiErr *apierr.APIError
			if errors.As(err, &apiErr) {
				return apiErr.IsRetriable(ctx)
			}
			return false
		},
	}), nil
}

func orDefault(configured, _default int) int {
	if configured == 0 {
		return _default
	}
	return configured
}

// noopLoader skips configuration loading
type noopLoader struct{}

func (noopLoader) Name() string                { return "noop" }
func (noopLoader) Configure(cfg *Config) error { return nil }

// noopAuth skips authentication
type noopAuth struct{}

func (noopAuth) Name() string { return "noop" }
func (noopAuth) Configure(context.Context, *Config) (func(*http.Request) error, error) {
	return func(r *http.Request) error { return nil }, nil
}
