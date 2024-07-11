package config

import (
	"context"
	"net/url"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const jwtBearerGrantTypeURN = "urn:ietf:params:oauth:grant-type:jwt-bearer"

type DatabricksOIDCCredentials struct{}

// Configure implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" || cfg.ClientID == "" {
		return nil, nil
	}
	if cfg.IsAccountClient() {
		logger.Debugf(ctx, "In-house OIDC is not yet supported for account clients")
	}

	// Get the OIDC token from the environment.
	// TODO: align audience with auth service expected audience
	idToken, err := cfg.getAllOIDCSuppliers().GetOIDCToken(ctx, "")
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
			"grant_type": {jwtBearerGrantTypeURN},
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

var _ CredentialsStrategy = DatabricksOIDCCredentials{}
