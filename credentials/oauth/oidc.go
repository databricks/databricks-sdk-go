package oauth

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

var ErrOAuthNotSupported = errors.New("databricks OAuth is not supported for this host")

func GetAccountOAuthEndpoints(ctx context.Context, accountsHost, accountId string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", accountsHost, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", accountsHost, accountId),
	}, nil
}

func GetWorkspaceOAuthEndpoints(ctx context.Context, c *httpclient.ApiClient, host string) (*OAuthAuthorizationServer, error) {
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", host)
	var oauthEndpoints OAuthAuthorizationServer
	if err := c.Do(ctx, "GET", oidc, httpclient.WithResponseUnmarshal(&oauthEndpoints)); err != nil {
		return nil, ErrOAuthNotSupported
	}
	return &oauthEndpoints, nil
}

type OAuthAuthorizationServer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"` // ../v1/authorize
	TokenEndpoint         string `json:"token_endpoint"`         // ../v1/token
}
