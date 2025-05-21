package config

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/config/credentials"
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
	ts := (&clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AuthStyle:    oauth2.AuthStyleInHeader,
		TokenURL:     endpoints.TokenEndpoint,
		Scopes:       []string{"all-apis"},
	}).TokenSource(ctx)

	visitor := refreshableVisitor(ts)
	return credentials.NewOAuthCredentialsProvider(visitor, ts.Token), nil
}
