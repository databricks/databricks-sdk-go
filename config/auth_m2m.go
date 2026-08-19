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
	ts, err := databricksOAuthTokenSource(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return credentials.NewOAuthCredentialsProviderFromTokenSource(
		auth.NewCachedTokenSource(ts, cacheOptions(cfg)...),
	), nil
}

// databricksOAuthTokenSource returns a token source that mints Databricks OAuth
// access tokens for the configured service principal (client_id/client_secret)
// via the client-credentials grant.
//
// The returned source is uncached: it performs a network call on each Token()
// and does not itself dedupe. Callers must wrap it in auth.NewCachedTokenSource
// (as M2mCredentials does), or pass it to serviceToServiceVisitor, which caches
// internally.
func databricksOAuthTokenSource(ctx context.Context, cfg *Config) (auth.TokenSource, error) {
	endpoints, err := cfg.getOidcEndpoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
	}
	logger.Debugf(ctx, "Generating Databricks OAuth token for Service Principal (%s)", cfg.ClientID)

	// IMPORTANT: do not use Config.TokenSource, which returns an
	// oauth2.TokenSource already wrapped in a cache. This leads to
	// double-caching that defeats the proactive async refresh provided by
	// NewCachedTokenSource (see [issue1549] for context).
	//
	// [issue1549]: https://github.com/databricks/databricks-sdk-go/issues/1549
	ccfg := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		TokenURL:     endpoints.TokenEndpoint,
		Scopes:       cfg.GetScopes(),
	}
	ts := auth.TokenSourceFn(func(ctx context.Context) (*oauth2.Token, error) {
		// Callers like CredentialsProvider.SetHeaders pass context.Background()
		// rather than the actual context from the HTTP client. Because of this,
		// the request would bypass the configured transport.
		//
		// The following is a workaround to ensure the refresh client's
		// transport is always used.
		ctx = cfg.refreshClient.InContextForOAuth2(ctx)
		return ccfg.Token(ctx)
	})

	return auth.NewRetryingTokenSource(ts), nil
}
