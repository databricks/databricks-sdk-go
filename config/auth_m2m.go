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

	// This already implements the auth.TokenSource interface. There is no need
	// to wrap it in auth.TokenSourceFn.
	//
	// IMPORTANT: the code should not rely on Config.TokenSource which returns
	// a oauth2.TokenSource that is already wrapped in a cache. This leads to
	// double-caching issues that defeats the proactive async refresh provided
	// by NewCachedTokenSource (see [issue1549] for context).
	//
	// Using clientcredentials.Config.Token directly avoids that issue at the
	// expense of allocating (and then discarding) a reusable token source at
	// each call.
	//
	// [issue1549]: https://github.com/databricks/databricks-sdk-go/issues/1549
	ts := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		TokenURL:     endpoints.TokenEndpoint,
		Scopes:       cfg.GetScopes(),
	}

	return credentials.NewOAuthCredentialsProviderFromTokenSource(
		auth.NewCachedTokenSource(auth.NewRetryingTokenSource(ts), cacheOptions(cfg)...),
	), nil
}
