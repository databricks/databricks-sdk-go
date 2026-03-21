package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

type GoogleDefaultCredentials struct {
	// options used to enable unit testing mode for OIDC
	opts []option.ClientOption
}

func (c GoogleDefaultCredentials) Name() string {
	return "google-id"
}

func (c GoogleDefaultCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.GoogleServiceAccount == "" || !cfg.IsGcp() {
		return nil, nil
	}
	inner, err := c.idTokenSource(ctx, cfg.Host, cfg.GoogleServiceAccount, c.opts...)
	if err != nil {
		return nil, err
	}

	// Disable async token refresh for GCP credential providers.
	//
	// Google's impersonate.IDTokenSource and impersonate.CredentialsTokenSource
	// internally wrap their token sources in oauth2.ReuseTokenSource (with a
	// 10-second expiryDelta). These inner sources are unexported types, so we
	// cannot bypass the caching the way we did for M2M OAuth (#1550) and Azure
	// Client Secret (#1573).
	//
	// With async refresh enabled, the SDK's cachedTokenSource fires a proactive
	// refresh ~20 minutes before expiry, but the call is swallowed by the inner
	// ReuseTokenSource which returns its own cached token -- making the async
	// refresh entirely wasted work (see #1549).
	//
	// Recreating the token source on each refresh is not viable either because
	// creation is expensive (credential discovery, GCE metadata probes, HTTP
	// client setup with TLS).
	//
	// The Google library's own ReuseTokenSource still provides some token
	// renewal before expiry (10-second window), so tokens are refreshed without
	// the SDK's async mechanism.
	//
	// Do NOT re-enable async refresh here unless the Google libraries expose an
	// uncached token source or a way to control their internal caching.
	opts := append(cacheOptions(cfg), auth.WithAsyncRefresh(false))

	// Always attempt to create SA token source for the secondary header.
	// If it fails, fall back to refreshableVisitor with a warning.
	platform, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: cfg.GoogleServiceAccount,
		Scopes: []string{
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/compute",
		},
	}, c.opts...)
	if err != nil {
		logger.Warnf(ctx, "Failed to create GCP SA access token source: %v. Proceeding without SA token.", err)
		visitor := refreshableVisitor(inner, opts...)
		return credentials.CredentialsProviderFn(visitor), nil
	}
	logger.Infof(ctx, "Using Google Default Application Credentials")
	visitor := serviceToServiceVisitor(inner, platform, "X-Databricks-GCP-SA-Access-Token", true, opts...)
	return credentials.NewOAuthCredentialsProvider(visitor, inner.Token), nil
}

func (c GoogleDefaultCredentials) idTokenSource(ctx context.Context, host, serviceAccount string,
	opts ...option.ClientOption) (oauth2.TokenSource, error) {
	ts, err := impersonate.IDTokenSource(ctx, impersonate.IDTokenConfig{
		Audience:        host,
		TargetPrincipal: serviceAccount,
		IncludeEmail:    true,
	}, opts...)
	if err != nil {
		err = fmt.Errorf("could not obtain OIDC token. %w Running 'gcloud auth application-default login' may help", err)
		return nil, err
	}
	return ts, err
}
