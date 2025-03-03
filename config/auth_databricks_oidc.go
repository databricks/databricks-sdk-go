package config

import (
	"context"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type DatabricksOIDCCredentials struct{}

// Configure implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" || cfg.ClientID == "" {
		return nil, nil
	}

	// Get the OIDC token from the environment.
	audience := strings.TrimPrefix(cfg.CanonicalHostName(), "https://")
	if cfg.IsAccountClient() {
		audience = cfg.AccountID
	}
	supplier := GithubOIDCTokenSupplier{
		cfg: cfg,
	}
	idToken, err := supplier.GetOIDCToken(ctx, audience)
	if err != nil {
		return nil, err
	}
	if idToken == "" {
		logger.Debugf(ctx, "No OIDC token found")
		return nil, nil
	}

	endpoints, err := oidcEndpoints(ctx, cfg)
	if err != nil {
		return nil, err
	}

	tsConfig := clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: "",
		AuthStyle:    oauth2.AuthStyleInParams,
		TokenURL:     endpoints.TokenEndpoint,
		Scopes:       []string{"all-apis"},
		EndpointParams: url.Values{
			"grant_type": {httpclient.JWTGrantType},
			"assertion":  {idToken},
		},
	}
	ts := tsConfig.TokenSource(ctx)
	visitor := refreshableVisitor(ts)
	return credentials.NewOAuthCredentialsProvider(visitor, ts.Token), nil
}

// Name implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Name() string {
	return "inhouse-oidc"
}
