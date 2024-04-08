package httpclient

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func (c *ApiClient) GetOAuthToken(authorizationDetails interface{}, token *oauth2.Token) (string, error) {
	path := "/oidc/v1/token"
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	data := map[string]any{
		"grant_type":            "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"authorization_details": authorizationDetails,
		"assertion":             token.AccessToken,
	}
	response := ""
	opts := []DoOption{
		WithRequestHeaders(headers),
		WithCreateTokenRequest(data),
		WithResponseUnmarshal(&response),
	}
	err := c.Do(context.Background(), http.MethodPost, path, opts...)
	if err != nil {
		return "", err
	}
	return response, nil
}
