package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/logger"
)

var errOAuthNotSupported = errors.New("databricks OAuth is not supported for this host")

type M2mCredentials struct {
}

func (c M2mCredentials) Name() string {
	return "oauth-m2m"
}

func (c M2mCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
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
	return refreshableVisitor(ts), nil
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
	oidcResponse, err := http.Get(oidc)
	if err != nil {
		return nil, fmt.Errorf("fetch .well-known: %w", err)
	}
	if oidcResponse.StatusCode != 200 {
		logger.Warnf(ctx, "Failure getting OIDC response. Response status: %s", oidcResponse.Status)
		return nil, errOAuthNotSupported
	}
	if oidcResponse.Body == nil {
		return nil, fmt.Errorf("fetch .well-known: empty body")
	}
	defer oidcResponse.Body.Close()
	raw, err := io.ReadAll(oidcResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read .well-known: %w", err)
	}
	var oauthEndpoints oauthAuthorizationServer
	err = json.Unmarshal(raw, &oauthEndpoints)
	if err != nil {
		return nil, fmt.Errorf("parse .well-known: %w", err)
	}
	return &oauthEndpoints, nil
}

type oauthAuthorizationServer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"` // ../v1/authorize
	TokenEndpoint         string `json:"token_endpoint"`         // ../v1/token
}
