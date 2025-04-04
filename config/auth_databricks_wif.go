package config

import (
	"context"
	"errors"
	"net/url"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Constructs all Workload Identity Federation Credentials Strategies
func WifTokenCredentialStrategies(cfg *Config) []CredentialsStrategy {
	providers := map[string]TokenProvider{
		"github-oidc-databricks-wif": &GithubProvider{
			cfg: cfg,
		},
		// Add new providers at the end of the list
	}
	strategies := []CredentialsStrategy{}
	for name, provider := range providers {
		strategies = append(strategies, newWifTokenStrategy(name, cfg, provider))
	}
	return strategies
}

func newWifTokenStrategy(
	name string,
	cfg *Config,
	tokenProvider TokenProvider,
) CredentialsStrategy {
	wifTokenExchange := &wifTokenExchange{
		cfg:           cfg,
		tokenProvider: tokenProvider,
	}
	return NewTokenSourceStrategy(name, wifTokenExchange)
}

// wifTokenExchange is a auth.TokenSource which exchanges a token using
// Workload Identity Federation.
type wifTokenExchange struct {
	cfg           *Config
	tokenProvider TokenProvider
}

func (w *wifTokenExchange) Token(ctx context.Context) (*oauth2.Token, error) {
	if w.cfg.ClientID == "" {
		logger.Debugf(ctx, "Missing cfg.ClientID")
		return nil, errors.New("missing cfg.ClientID")
	}
	if w.cfg.Host == "" {
		logger.Debugf(ctx, "Missing cfg.Host")
		return nil, errors.New("missing cfg.Host")
	}
	audience, err := w.getAudience(ctx)
	if err != nil {
		return nil, err
	}
	idToken, err := w.tokenProvider.IDToken(ctx, audience)
	if err != nil {
		return nil, err
	}
	endpoints, err := w.cfg.getOidcEndpoints(ctx)
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

func (w *wifTokenExchange) getAudience(ctx context.Context) (string, error) {
	if w.cfg.TokenAudience != "" {
		return w.cfg.TokenAudience, nil
	}
	// For Databricks Accounts, the account id is the default audience.
	if w.cfg.IsAccountClient() {
		return w.cfg.AccountID, nil
	}
	// For Databricks Workspaces, the auth endpoint is the default audience.
	endpoints, err := w.cfg.getOidcEndpoints(ctx)
	if err != nil {
		return "", err
	}
	return endpoints.TokenEndpoint, nil
}
