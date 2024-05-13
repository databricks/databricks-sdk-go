package httpclient

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/credentials"
	"golang.org/x/oauth2"
)

type GetOAuthTokenRequest struct {
	GrantType            string `url:"grant_type"`
	AuthorizationDetails string `url:"authorization_details"`
	Assertion            string `url:"assertion"`
}

// Returns a new OAuth token using the provided token. The token must be a JWT token.
// The resulting token is scoped to the authorization details provided.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (c *ApiClient) GetOAuthToken(authDetails string, token *oauth2.Token) (*credentials.OAuthToken, error) {
	path := "/oidc/v1/token"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := GetOAuthTokenRequest{
		GrantType:            "urn:ietf:params:oauth:grant-type:jwt-bearer",
		AuthorizationDetails: authDetails,
		Assertion:            token.AccessToken,
	}
	var response credentials.OAuthToken
	opts := []DoOption{
		WithRequestHeaders(headers),
		WithUrlEncodedData(data),
		WithResponseUnmarshal(&response),
	}
	err := c.Do(c.InContextForOAuth2(context.Background()), http.MethodPost, path, opts...)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
