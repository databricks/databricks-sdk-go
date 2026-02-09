package u2m

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/httpclient"
)

var ErrOAuthNotSupported = errors.New("databricks OAuth is not supported for this host")

// OAuthEndpointSupplier provides the http functionality needed for interacting with the
// Databricks OAuth APIs.
type OAuthEndpointSupplier interface {
	// GetWorkspaceOAuthEndpoints returns the OAuth2 endpoints for the workspace.
	GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error)

	// GetAccountOAuthEndpoints returns the OAuth2 endpoints for the account.
	GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error)

	// GetUnifiedOAuthEndpoints returns the OAuth2 endpoints for the unified host.
	GetUnifiedOAuthEndpoints(ctx context.Context, host string, accountId string) (*OAuthAuthorizationServer, error)
}

// BasicOAuthEndpointSupplier is an implementation of the OAuthEndpointSupplier interface.
type BasicOAuthEndpointSupplier struct {
	// Client is the ApiClient to use for making HTTP requests.
	Client *httpclient.ApiClient
}

// getOAuthEndpointsByDiscoveryUrl queries the OIDC discovery endpoint to get the OAuth endpoints.
func (c *BasicOAuthEndpointSupplier) getOAuthEndpointsByDiscoveryUrl(ctx context.Context, discoveryUrl string) (*OAuthAuthorizationServer, error) {
	var oauthEndpoints OAuthAuthorizationServer
	if err := c.Client.Do(ctx, "GET", discoveryUrl, httpclient.WithResponseUnmarshal(&oauthEndpoints)); err != nil {
		if errors.Is(err, apierr.ErrNotFound) {
			return nil, ErrOAuthNotSupported
		}
		return nil, fmt.Errorf("failed to get OAuth endpoints: %w", err)
	}
	return &oauthEndpoints, nil
}

// GetWorkspaceOAuthEndpoints returns the OAuth endpoints for the given workspace.
func (c *BasicOAuthEndpointSupplier) GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error) {
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", workspaceHost)
	return c.getOAuthEndpointsByDiscoveryUrl(ctx, oidc)
}

// GetAccountOAuthEndpoints returns the OAuth2 endpoints for the account. The
// account-level OAuth endpoints are fixed based on the account ID and host.
func (c *BasicOAuthEndpointSupplier) GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error) {
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", accountHost)
	return c.getOAuthEndpointsByDiscoveryUrl(ctx, oidc)
}

// GetUnifiedOAuthEndpoints returns the OAuth2 endpoints for the unified host
func (c *BasicOAuthEndpointSupplier) GetUnifiedOAuthEndpoints(ctx context.Context, host string, accountId string) (*OAuthAuthorizationServer, error) {
	oidc := fmt.Sprintf("%s/oidc/accounts/%s/.well-known/oauth-authorization-server", host, accountId)
	return c.getOAuthEndpointsByDiscoveryUrl(ctx, oidc)
}

// OAuthAuthorizationServer contains the OAuth endpoints for a Databricks account
// or workspace.
type OAuthAuthorizationServer struct {
	// AuthorizationEndpoint is the URL to redirect users to for authorization.
	// It typically ends with /v1/authroize.
	AuthorizationEndpoint string `json:"authorization_endpoint"`

	// TokenEndpoint is the URL to exchange an authorization code for an access token.
	// It typically ends with /v1/token.
	TokenEndpoint string `json:"token_endpoint"`
}
