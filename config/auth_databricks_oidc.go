package config

import (
	"context"
	"errors"
	"net/url"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func NewOIDCTokenExchange(
	cfg *OIDCTokenExchangeConfig,
	idTokenProvider IDTokenSource,
) auth.TokenSource {
	return &oidcTokenExchange{
		clientID:      cfg.ClientID,
		accountID:     cfg.AccountID,
		host:          cfg.Host,
		tokenEndpoint: cfg.TokenEndpoint,
		audience:      cfg.Audience,
		tokenProvider: idTokenProvider,
	}
}

type OIDCTokenExchangeConfig struct {
	ClientID        string
	AccountID       string
	Host            string
	TokenEndpoint   string
	Audience        string
	IdTokenProvider IDTokenSource
}

// oidcTokenExchange is a auth.TokenSource which exchanges a token using
// Workload Identity Federation.
type oidcTokenExchange struct {
	clientID      string
	accountID     string
	host          string
	tokenEndpoint string
	audience      string
	tokenProvider IDTokenSource
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
	audience := w.determineAudience()
	idToken, err := w.tokenProvider.IDToken(ctx, audience)
	if err != nil {
		return nil, err
	}
	c := &clientcredentials.Config{
		ClientID:  w.clientID,
		AuthStyle: oauth2.AuthStyleInParams,
		TokenURL:  w.tokenEndpoint,
		Scopes:    []string{"all-apis"},
		EndpointParams: url.Values{
			"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
			"subject_token":      {idToken.Value},
			"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
		},
	}
	return c.Token(ctx)
}

func (w *oidcTokenExchange) determineAudience() string {
	if w.audience != "" {
		return w.audience
	}
	// For Databricks Accounts, the account id is the default audience.
	if w.accountID != "" {
		return w.accountID
	}
	// For Databricks Workspaces, the auth endpoint is the default audience.
	return w.tokenEndpoint
}
