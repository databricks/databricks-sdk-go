package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func New(cfg *config.Config) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	return newWithTransport(cfg, cfg.HTTPTransport), nil
}

func newWithTransport(cfg *config.Config, transport http.RoundTripper) *DatabricksClient {
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
			Transport:          transport,
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
				"com.databricks.backend.manager.util.UnknownWorkerEnvironmentException",
				"does not have any associated worker environments",
				"There is no worker environment with id",
				"Unknown worker environment",
				"ClusterNotReadyException",
				"connection reset by peer",
				"TLS handshake timeout",
				"connection refused",
				"Unexpected error",
				"i/o timeout",
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
	}
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
		opts = append(opts, httpclient.WithVisitor(v))
	}
	opts = append(opts, httpclient.WithHeaders(headers))
	opts = append(opts, httpclient.WithData(request))
	opts = append(opts, httpclient.WithUnmarshal(response))
	return c.client.Do(ctx, method, path, opts...)
}

func orDefault(configured, _default int) int {
	if configured == 0 {
		return _default
	}
	return configured
}
