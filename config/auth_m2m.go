package config

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/logger"
)

type M2mCredentials struct{}

func (c M2mCredentials) Name() string {
	return "oauth-m2m"
}

func (c M2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, nil
	}
	endpoints, err := cfg.getOidcEndpoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
	}
	logger.Debugf(ctx, "Generating Databricks OAuth token for Service Principal (%s)", cfg.ClientID)
	ccfg := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		TokenURL:     endpoints.TokenEndpoint,
		Scopes:       cfg.GetScopes(),
	}

	// Use a direct (non-caching) token source so that cachedTokenSource is the
	// single cache layer. clientcredentials.Config.TokenSource returns an
	// oauth2.ReuseTokenSource which adds a second cache with a 10 s expiryDelta.
	// With double-caching the async refresh in cachedTokenSource calls through to
	// ReuseTokenSource, which returns its own cached token without making an HTTP
	// request until only ~10 s remain — defeating the proactive 20-min refresh
	// window and causing a burst of 401s at token expiry.
	// See https://github.com/databricks/databricks-sdk-go/issues/1549.
	//
	// ctx is captured from Configure; the oauth2 HTTP client is bound to it via
	// InContextForOAuth2 and must be used for all token requests.
	directTS := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return ccfg.Token(ctx)
	})
	return credentials.NewOAuthCredentialsProviderFromTokenSource(
		auth.NewCachedTokenSource(directTS, cacheOptions(cfg)...),
	), nil
}
