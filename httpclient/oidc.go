package httpclient

import (
	"context"
	"errors"
	"fmt"
)

var ErrOAuthNotSupported = errors.New("databricks OAuth is not supported for this host")

func (c *ApiClient) GetOidcEndpoints(ctx context.Context, host, accountId string) (*OAuthAuthorizationServer, error) {
	prefix := host
	if /* cfg.IsAccountClient() && */ accountId != "" {
		// TODO: technically, we could use the same config profile for both workspace
		// and account, but we have to add logic for determining accounts host from
		// workspace host.
		prefix := fmt.Sprintf("%s/oidc/accounts/%s", host, accountId)
		return &OAuthAuthorizationServer{
			AuthorizationEndpoint: fmt.Sprintf("%s/v1/authorize", prefix),
			TokenEndpoint:         fmt.Sprintf("%s/v1/token", prefix),
		}, nil
	}
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", prefix)
	var oauthEndpoints OAuthAuthorizationServer
	err := c.Do(ctx, "GET", oidc, WithResponseUnmarshal(&oauthEndpoints))
	if err != nil {
		return nil, ErrOAuthNotSupported
	}
	return &oauthEndpoints, nil
}

type OAuthAuthorizationServer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"` // ../v1/authorize
	TokenEndpoint         string `json:"token_endpoint"`         // ../v1/token
}
