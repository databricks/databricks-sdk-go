package config

import (
	"context"
	"errors"
	"net/url"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Creates a new Databricks OIDC TokenSource.
func NewDatabricksOIDCTokenSource(cfg DatabricksOIDCTokenSourceConfig) auth.TokenSource {
	return &databricksOIDCTokenSource{
		cfg: cfg,
	}
}

// Config for Databricks OIDC TokenSource.
type DatabricksOIDCTokenSourceConfig struct {
	// ClientID is the client ID of the Databricks OIDC application. For
	// Databricks Service Principal, this is the Application ID of the Service Principal.
	ClientID string
	// [Optional] AccountID is the account ID of the Databricks Account.
	// This is only used for Account level tokens.
	AccountID string
	// Host is the host of the Databricks account or workspace.
	Host string
	// TokenEndpointProvider returns the token endpoint for the Databricks OIDC application.
	TokenEndpointProvider func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error)
	// Audience is the audience of the Databricks OIDC application.
	// This is only used for Workspace level tokens.
	Audience string
	// IdTokenSource returns the IDToken to be used for the token exchange.
	IdTokenSource IDTokenSource
}

// databricksOIDCTokenSource is a auth.TokenSource which exchanges a token using
// Workload Identity Federation.
type databricksOIDCTokenSource struct {
	cfg DatabricksOIDCTokenSourceConfig
}

// Token implements [TokenSource.Token]
func (w *databricksOIDCTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	if w.cfg.ClientID == "" {
		logger.Debugf(ctx, "Missing ClientID")
		return nil, errors.New("missing ClientID")
	}
	if w.cfg.Host == "" {
		logger.Debugf(ctx, "Missing Host")
		return nil, errors.New("missing Host")
	}
	endpoints, err := w.cfg.TokenEndpointProvider(ctx)
	if err != nil {
		return nil, err
	}
	audience := w.determineAudience(endpoints)
	idToken, err := w.cfg.IdTokenSource.IDToken(ctx, audience)
	if err != nil {
		return nil, err
	}

	c := &clientcredentials.Config{
		ClientID:  w.cfg.ClientID,
		AuthStyle: oauth2.AuthStyleInParams,
		TokenURL:  endpoints.TokenEndpoint,
		Scopes:    []string{"all-apis"},
		EndpointParams: url.Values{
			"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
			"subject_token":      {idToken.Value},
			"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
		},
	}
	return c.Token(ctx)
}

func (w *databricksOIDCTokenSource) determineAudience(endpoints *u2m.OAuthAuthorizationServer) string {
	if w.cfg.Audience != "" {
		return w.cfg.Audience
	}
	// For Databricks Accounts, the account id is the default audience.
	if w.cfg.AccountID != "" {
		return w.cfg.AccountID
	}
	// For Databricks Workspaces, the auth endpoint is the default audience.
	return endpoints.TokenEndpoint
}
