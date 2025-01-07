package config

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

var errOAuthNotSupported = errors.New("databricks OAuth is not supported for this host")

type M2mCredentials struct {
}

func (c M2mCredentials) Name() string {
	return "oauth-m2m"
}

func (c M2mCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" {
		return nil, nil
	}
	endpoints, err := oidcEndpoints(ctx, cfg)
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

func oidcEndpoints(ctx context.Context, cfg *Config) (*oauthAuthorizationServer, error) {
	prefix := cfg.Host
	if cfg.IsAccountClient() && cfg.AccountID != "" {
		// TODO: technically, we could use the same config profile for both workspace
		// and account, but we have to add logic for determining accounts host from
		// workspace host.
		prefix := fmt.Sprintf("%s/oidc/accounts/%s", cfg.Host, cfg.AccountID)
		return &oauthAuthorizationServer{
			AuthorizationEndpoint: fmt.Sprintf("%s/v1/authorize", prefix),
			TokenEndpoint:         fmt.Sprintf("%s/v1/token", prefix),
		}, nil
	}
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", prefix)
	var oauthEndpoints oauthAuthorizationServer
	err := cfg.refreshClient.Do(ctx, "GET", oidc,
		httpclient.WithResponseUnmarshal(&oauthEndpoints))
	if err != nil {
		return nil, errOAuthNotSupported
	}
	return &oauthEndpoints, nil
}

type oauthAuthorizationServer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"` // ../v1/authorize
	TokenEndpoint         string `json:"token_endpoint"`         // ../v1/token
}
