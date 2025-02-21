package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"golang.org/x/oauth2"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// OAuthClient is an interface for Databricks OAuth client.
type OAuthClient struct {
	client httpClient
}

const JWTGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"

// Returns a new OAuth token using the provided token. The token must be a JWT token.
// The resulting token is scoped to the authorization details provided.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (oac *OAuthClient) GetOAuthToken(ctx context.Context, authDetails string, token *oauth2.Token) (*oauth2.Token, error) {
	path := "/oidc/v1/token"
	data := getOAuthTokenRequest{
		GrantType:            JWTGrantType,
		AuthorizationDetails: authDetails,
		Assertion:            token.AccessToken,
	}
	var response credentials.OAuthToken
	opts := []DoOption{
		WithUrlEncodedData(data),
		WithResponseUnmarshal(&response),
	}
	err := c.Do(ctx, http.MethodPost, path, opts...)
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: response.AccessToken,
		TokenType:   response.TokenType,
		Expiry:      time.Now().Add(time.Duration(response.ExpiresIn) * time.Second),
	}, nil
}

// GetOAuthTokenRequest is the request to get an OAuth token. It follows the
// OAuth 2.0 Rich Authorization Requests specification.
//
// https://datatracker.ietf.org/doc/html/rfc9396
type getOAuthTokenRequest struct {
	// Defines the method used to get the token.
	GrantType string `url:"grant_type"`
	// An array of authorization details that the token should be scoped to. This needs to be passed in string format.
	AuthorizationDetails string `url:"authorization_details"`
	// The token that will be exchanged for an OAuth token.
	Assertion string `url:"assertion"`
}
