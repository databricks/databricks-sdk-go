package config

import (
	"context"
	"errors"
	"net/url"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Constructs all Workload Identity Federation Credentials Strategies
func OidcTokenCredentialStrategies(cfg *Config) []CredentialsStrategy {
	providers := map[string]TokenProvider{
		"github-oidc": &GithubProvider{
			actionsIDTokenRequestURL:   cfg.ActionsIDTokenRequestURL,
			actionsIDTokenRequestToken: cfg.ActionsIDTokenRequestToken,
			refreshClient:              cfg.refreshClient,
		},
		// Add new providers at the end of the list
	}
	strategies := []CredentialsStrategy{}
	for name, provider := range providers {
		strategies = append(strategies, newOidcTokenStrategy(name, cfg, provider))
	}
	return strategies
}

func newOidcTokenStrategy(
	name string,
	cfg *Config,
	tokenProvider TokenProvider,
) CredentialsStrategy {
	accountId := ""
	// NOTE: cfg.AccountID can be present even if the client is not an account client.
	if cfg.IsAccountClient() {
		accountId = cfg.AccountID
	}
	oidcTokenExchange := &oidcTokenExchange{
		clientID:              cfg.ClientID,
		accountID:             accountId,
		host:                  cfg.Host,
		tokenEndpointProvider: cfg.getOidcEndpoints,
		audience:              cfg.TokenAudience,
		tokenProvider:         tokenProvider,
	}
	return NewTokenSourceStrategy(name, oidcTokenExchange)
}

// oidcTokenExchange is a auth.TokenSource which exchanges a token using
// Workload Identity Federation.
type oidcTokenExchange struct {
	clientID              string
	accountID             string
	host                  string
	tokenEndpointProvider func(ctx context.Context) (*u2m.OAuthAuthorizationServer, error)
	audience              string
	tokenProvider         TokenProvider
}

// Token implements [TokenSource.Token]
func (w *oidcTokenExchange) Token(ctx context.Context) (*oauth2.Token, error) {
	if w.clientID == "" {
		logger.Debugf(ctx, "Missing ClientID")
		return nil, errors.New("missing ClientID")
	}
	if w.host == "" {
		logger.Debugf(ctx, "Missing Host")
		return nil, errors.New("missing Host")
	}
	audience, err := w.determineAudience(ctx)
	if err != nil {
		return nil, err
	}
	idToken, err := w.tokenProvider.IDToken(ctx, audience)
	if err != nil {
		return nil, err
	}
	endpoints, err := w.tokenEndpointProvider(ctx)
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

func (w *oidcTokenExchange) determineAudience(ctx context.Context) (string, error) {
	if w.audience != "" {
		return w.audience, nil
	}
	// For Databricks Accounts, the account id is the default audience.
	if w.accountID != "" {
		return w.accountID, nil
	}
	endpoints, err := w.tokenEndpointProvider(ctx)
	if err != nil {
		return "", err
	}
	// For Databricks Workspaces, the auth endpoint is the default audience.
	return endpoints.TokenEndpoint, nil
}
