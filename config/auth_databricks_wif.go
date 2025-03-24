package config

import (
	"context"
	"errors"
	"net/url"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// DatabricksWIFCredentials uses a Token Supplier to get a JWT Token and exchanges
// it for a Databricks Token.
//
// Supported suppliers:
// - GitHub OIDC
type DatabricksWIFCredentials struct{}

// Configure implements CredentialsStrategy.
func (d DatabricksWIFCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" || cfg.ClientID == "" {
		return nil, nil
	}

	supplier := GithubOIDCTokenSupplier{
		cfg: cfg,
	}

	// If no supplier can get an IdToken, skip this CredentialsStrategy
	idToken, err := supplier.GetOIDCToken(ctx, cfg.TokenAudience)
	if err != nil {
		return nil, err
	}
	if idToken == "" {
		return nil, nil
	}

	endpoints, err := oidcEndpoints(ctx, cfg)
	if err != nil {
		return nil, err
	}

	ts := &databricksWIFTokenSource{
		ctx:             ctx,
		idTokenSupplier: &supplier,
		audience:        cfg.TokenAudience,
		clientId:        cfg.ClientID,
		cfg:             cfg,
		tokenEndpoint:   endpoints.TokenEndpoint,
	}

	visitor := refreshableVisitor(ts)
	return credentials.NewOAuthCredentialsProvider(visitor, ts.Token), nil
}

// Name implements CredentialsStrategy.
func (d DatabricksWIFCredentials) Name() string {
	return "databricks-wif"
}

type databricksWIFTokenSource struct {
	ctx             context.Context
	idTokenSupplier *GithubOIDCTokenSupplier
	tokenEndpoint   string
	audience        string
	clientId        string
	cfg             *Config
}

func (d *databricksWIFTokenSource) Token() (*oauth2.Token, error) {
	token, err := d.idTokenSupplier.GetOIDCToken(d.ctx, d.audience)
	if err != nil {
		return nil, err
	}
	if token == "" {
		// It should not happen, since we check before constructing the token source.
		logger.Debugf(d.ctx, "No OIDC token found")
		return nil, errors.New("cannot get OIDC token")
	}

	tsConfig := clientcredentials.Config{
		ClientID:  d.clientId,
		AuthStyle: oauth2.AuthStyleInParams,
		TokenURL:  d.tokenEndpoint,
		Scopes:    []string{"all-apis"},
		EndpointParams: url.Values{
			"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
			"subject_token":      {token},
			"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
		},
	}

	logger.Debugf(d.ctx, "Getting token for client %s", d.clientId)

	return tsConfig.Token(d.ctx)
}
