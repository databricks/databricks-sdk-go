package httpclient

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/credentials"
	"golang.org/x/oauth2"
)

type GetOAuthTokenRequest struct {
	GrantType            string                           `url:"grant_type"`
	AuthorizationDetails credentials.AuthorizationDetails `url:"authorization_details"`
	Assertion            string                           `url:"assertion"`
	ExpiresIn            int                              `url:"expires_in,omitempty"`
	RefreshToken         bool                             `url:"refresh_token,omitempty"`
}

// Returns a new OAuth token using the provided token. The token must be a JWT token.
// The resulting token is scoped to the authorization details provided.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (c *ApiClient) GetOAuthToken(authorizationDetails string, token *oauth2.Token) (*credentials.OAuthToken, error) {
	path := "/oidc/v1/token"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := GetOAuthTokenRequest{
		GrantType:            "urn:ietf:params:oauth:grant-type:jwt-bearer",
		AuthorizationDetails: credentials.AuthorizationDetails(authorizationDetails),
		Assertion:            token.AccessToken,
		// By default, the tokens API does not refresh the token, so you can get a token that is about to expire.
		// We keep a cache of the Tokens, and only request when a new one is needed
		ExpiresIn:    3600,
		RefreshToken: true,
	}
	var response credentials.OAuthToken
	opts := []DoOption{
		WithRequestHeaders(headers),
		WithUrlEncodedData(data),
		WithResponseUnmarshal(&response),
	}
	err := c.Do(context.Background(), http.MethodPost, path, opts...)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
