package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func HTTPClientConfigFromConfig(cfg *Config) (httpclient.ClientConfig, error) {
	if skippable, ok := cfg.HTTPTransport.(interface {
		SkipRetryOnIO() bool
	}); ok && skippable.SkipRetryOnIO() {
		cfg.Loaders = []Loader{noopLoader{}}
		cfg.Credentials = noopAuth{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return httpclient.ClientConfig{}, err
	}

	visitors := []httpclient.RequestVisitor{
		func(r *http.Request) error {
			if r.URL == nil {
				return fmt.Errorf("no URL found in request")
			}
			url, err := url.Parse(cfg.Host)
			if err != nil {
				return err
			}
			r.URL.Host = url.Host
			r.URL.Scheme = url.Scheme
			return nil
		},
		authInUserAgentVisitor(cfg),
		func(r *http.Request) error {
			// Detect if we are running in a CI/CD environment
			provider := useragent.CiCdProvider()
			if provider == "" {
				return nil
			}
			// Add the detected CI/CD provider to the user agent
			ctx := useragent.InContext(r.Context(), useragent.CicdKey, provider)
			*r = *r.WithContext(ctx) // replace request
			return nil
		},
		func(r *http.Request) error {
			// Detect if the SDK is being run in a Databricks Runtime.
			v := useragent.Runtime()
			if v == "" {
				return nil
			}
			// Add the detected Databricks Runtime version to the user agent
			ctx := useragent.InContext(r.Context(), useragent.RuntimeKey, v)
			*r = *r.WithContext(ctx) // replace request
			return nil
		},
	}

	// Workspace IDs must be provided explicitly in request headers when using unified hosts.
	// This visitor assumes that cfg.WorkspaceId is only set for workspace clients.
	if cfg.HostType() == UnifiedHost && cfg.WorkspaceId != "" {
		visitors = append(visitors, func(r *http.Request) error {
			r.Header.Set("X-Databricks-Org-Id", cfg.WorkspaceId)
			return nil
		})
	}

	return httpclient.ClientConfig{
		AccountID:          cfg.AccountID,
		Host:               cfg.Host,
		RetryTimeout:       time.Duration(cfg.RetryTimeoutSeconds) * time.Second,
		HTTPTimeout:        time.Duration(cfg.HTTPTimeoutSeconds) * time.Second,
		RateLimitPerSecond: cfg.RateLimitPerSecond,
		DebugHeaders:       cfg.DebugHeaders,
		DebugTruncateBytes: cfg.DebugTruncateBytes,
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		Transport:          cfg.HTTPTransport,
		AuthVisitor:        cfg.Authenticate,
		Visitors:           visitors,
		TransientErrors: []string{
			// This is temporary workaround for SCIM API returning 500.
			// TODO: Remove when it's fixed.
			"REQUEST_LIMIT_EXCEEDED",
		},
		ErrorMapper: apierr.GetAPIError,
		ErrorRetriable: func(ctx context.Context, err error) bool {
			var apiErr *apierr.APIError
			if errors.As(err, &apiErr) {
				return apiErr.IsRetriable(ctx)
			}
			return false
		},
	}, nil
}

// noopLoader skips configuration loading
type noopLoader struct{}

func (noopLoader) Name() string                { return "noop" }
func (noopLoader) Configure(cfg *Config) error { return nil }

// noopAuth skips authentication
type noopAuth struct{}

func (noopAuth) Name() string { return "noop" }
func (noopAuth) Configure(context.Context, *Config) (credentials.CredentialsProvider, error) {
	visitor := func(r *http.Request) error { return nil }
	return credentials.CredentialsProviderFn(visitor), nil
}

// Deprecated: use [HTTPClientConfigFromConfig] with [httpclient.NewApiClient].
func (c *Config) NewApiClient() (*httpclient.ApiClient, error) {
	cfg, err := HTTPClientConfigFromConfig(c)
	if err != nil {
		return nil, err
	}
	return httpclient.NewApiClient(cfg), nil
}
