package config

import (
	"context"
	"net/url"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type DatabricksOIDCCredentials struct{}

// Configure implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" || cfg.ClientID == "" || cfg.TokenAudience == "" {
		return nil, nil
	}

	// Get the OIDC token from the environment.
	supplier := GithubOIDCTokenSupplier{
		cfg: cfg,
	}

	ts := &databricksOIDCTokenSource{
		ctx:               ctx,
		jwtTokenSupplicer: &supplier,
		audience:          cfg.TokenAudience,
		clientId:          cfg.ClientID,
		cfg:               cfg,
	}

	visitor := refreshableVisitor(ts)
	return credentials.NewOAuthCredentialsProvider(visitor, ts.Token), nil
}

// Name implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Name() string {
	return "databricks-oidc"
}

type databricksOIDCTokenSource struct {
	ctx               context.Context
	jwtTokenSupplicer *GithubOIDCTokenSupplier
	audience          string
	clientId          string
	cfg               *Config
}

func (d *databricksOIDCTokenSource) Token() (*oauth2.Token, error) {
	token, err := d.jwtTokenSupplicer.GetOIDCToken(d.ctx, d.audience)
	if err != nil {
		return nil, err
	}
	if token == "" {
		logger.Debugf(d.ctx, "No OIDC token found")
		return nil, nil
	}

	endpoints, err := oidcEndpoints(d.ctx, d.cfg)
	if err != nil {
		return nil, err
	}
	logger.Debugf(d.ctx, "Getting tokken for client %s", d.clientId)

	tsConfig := clientcredentials.Config{
		ClientID:  d.clientId,
		AuthStyle: oauth2.AuthStyleInParams,
		TokenURL:  endpoints.TokenEndpoint,
		Scopes:    []string{"all-apis"},
		EndpointParams: url.Values{
			"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
			"subject_token":      {token},
			"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
		},
	}

	// Request the token
	return tsConfig.Token(d.ctx)
}
