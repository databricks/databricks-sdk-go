package oauth

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

var ErrOAuthNotSupported = errors.New("databricks OAuth is not supported for this host")

// GetAccountOAuthEndpoints returns the OAuth endpoints for the given account.
func GetAccountOAuthEndpoints(ctx context.Context, accountsHost, accountId string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", accountsHost, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", accountsHost, accountId),
	}, nil
}

// GetWorkspaceOAuthEndpoints returns the OAuth endpoints for the given workspace,
// It queries the OIDC discovery endpoint to get the OAuth endpoints using the
// provided ApiClient.
func GetWorkspaceOAuthEndpoints(ctx context.Context, c *httpclient.ApiClient, host string) (*OAuthAuthorizationServer, error) {
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", host)
	var oauthEndpoints OAuthAuthorizationServer
	if err := c.Do(ctx, "GET", oidc, httpclient.WithResponseUnmarshal(&oauthEndpoints)); err != nil {
		return nil, ErrOAuthNotSupported
	}
	return &oauthEndpoints, nil
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
