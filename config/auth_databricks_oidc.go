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
func NewDatabricksOIDCTokenSource(
	cfg *DatabricksOIDCTokenSourceConfig,
) auth.TokenSource {
	return &databricksOIDCTokenSource{
		clientID:              cfg.ClientID,
		accountID:             cfg.AccountID,
		host:                  cfg.Host,
		tokenEndpointProvider: cfg.TokenEndpointProvider,
		audience:              cfg.Audience,
		idTokenProvider:       cfg.IdTokenSource,
	}
}

// Config for Databricks OIDC TokenSource.
type DatabricksOIDCTokenSourceConfig struct {
	ClientID              string
	AccountID             string
	Host                  string
	TokenEndpointProvider func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error)
	Audience              string
	IdTokenSource         IDTokenSource
}

// databricksOIDCTokenSource is a auth.TokenSource which exchanges a token using
// Workload Identity Federation.
type databricksOIDCTokenSource struct {
	clientID              string
	accountID             string
	host                  string
	tokenEndpointProvider func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error)
	audience              string
	idTokenProvider       IDTokenSource
}

// Token implements [TokenSource.Token]
func (w *databricksOIDCTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	if w.clientID == "" {
		logger.Debugf(ctx, "Missing ClientID")
		return nil, errors.New("missing ClientID")
	}
	if w.host == "" {
		logger.Debugf(ctx, "Missing Host")
		return nil, errors.New("missing Host")
	}
	endpoints, err := w.tokenEndpointProvider(ctx)
	if err != nil {
		return nil, err
	}
	audience := w.determineAudience(endpoints)
	idToken, err := w.idTokenProvider.IDToken(ctx, audience)
	if err != nil {
		return nil, err
	}

	c := &clientcredentials.Config{
		ClientID:  w.clientID,
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
	if w.audience != "" {
		return w.audience
	}
	// For Databricks Accounts, the account id is the default audience.
	if w.accountID != "" {
		return w.accountID
	}
	// For Databricks Workspaces, the auth endpoint is the default audience.
	return endpoints.TokenEndpoint
}
