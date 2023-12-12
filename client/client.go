package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func New(cfg *config.Config) (*DatabricksClient, error) {
	if skippable, ok := cfg.HTTPTransport.(interface {
		SkipRetryOnIO() bool
	}); ok && skippable.SkipRetryOnIO() {
		cfg.Loaders = []config.Loader{noopLoader{}}
		cfg.Credentials = noopAuth{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	retryTimeout := time.Duration(orDefault(cfg.RetryTimeoutSeconds, 300)) * time.Second
	httpTimeout := time.Duration(orDefault(cfg.HTTPTimeoutSeconds, 60)) * time.Second
	return &DatabricksClient{
		Config: cfg,
		client: httpclient.NewApiClient(httpclient.ClientConfig{
			RetryTimeout:       retryTimeout,
			HTTPTimeout:        httpTimeout,
			RateLimitPerSecond: orDefault(cfg.RateLimitPerSecond, 15),
			DebugHeaders:       cfg.DebugHeaders,
			DebugTruncateBytes: cfg.DebugTruncateBytes,
			InsecureSkipVerify: cfg.InsecureSkipVerify,
			Transport:          cfg.HTTPTransport,
			Visitors: []httpclient.RequestVisitor{
				cfg.Authenticate,
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
				func(r *http.Request) error {
					ctx := useragent.InContext(r.Context(), "auth", cfg.AuthType)
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
		}),
	}, nil
}

type DatabricksClient struct {
	Config *config.Config
	client *httpclient.ApiClient
}

// ConfiguredAccountID returns Databricks Account ID if it's provided in config,
// empty string otherwise
func (c *DatabricksClient) ConfiguredAccountID() string {
	return c.Config.AccountID
}

// Do sends an HTTP request against path.
func (c *DatabricksClient) Do(ctx context.Context, method, path string,
	headers map[string]string, request, response any,
	visitors ...func(*http.Request) error) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))
	// Remove extra `/` from path for files API.
	// Once the OpenAPI spec doesn't include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}
	return c.client.Do(ctx, method, path, opts...)
}

func orDefault(configured, _default int) int {
	if configured == 0 {
		return _default
	}
	return configured
}

// noopLoader skips configuration loading
type noopLoader struct{}

func (noopLoader) Name() string                       { return "noop" }
func (noopLoader) Configure(cfg *config.Config) error { return nil }

// noopAuth skips authentication
type noopAuth struct{}

func (noopAuth) Name() string { return "noop" }
func (noopAuth) Configure(context.Context, *config.Config) (func(*http.Request) error, error) {
	return nil, nil
}
