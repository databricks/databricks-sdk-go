package config

import (
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/internal/credentials/oauth"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func TestM2mHappyFlow(t *testing.T) {
	assertHeaders(t, &Config{
		Host:         "a",
		ClientID:     "b",
		ClientSecret: "c",
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: oauth.OAuthAuthorizationServer{
					AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
					TokenEndpoint:         "https://localhost:1234/dummy/token",
				},
			},
			"POST /dummy/token": {
				ExpectedHeaders: map[string]string{
					"Authorization": "Basic Yjpj",
					"Content-Type":  "application/x-www-form-urlencoded",
				},
				ExpectedRequest: url.Values{
					"grant_type": {"client_credentials"},
					"scope":      {"all-apis"},
				},
				Response: oauth2.Token{
					TokenType:   "Some",
					AccessToken: "cde",
				},
			},
		},
	}, map[string]string{
		"Authorization": "Some cde",
	})
}

func TestM2mHappyFlowForAccount(t *testing.T) {
	assertHeaders(t, &Config{
		Host:         "accounts.cloud.databricks.com",
		AccountID:    "a",
		ClientID:     "b",
		ClientSecret: "c",
		HTTPTransport: fixtures.MappingTransport{
			"POST /oidc/accounts/a/v1/token": {
				ExpectedHeaders: map[string]string{
					"Authorization": "Basic Yjpj",
					"Content-Type":  "application/x-www-form-urlencoded",
				},
				ExpectedRequest: url.Values{
					"grant_type": {"client_credentials"},
					"scope":      {"all-apis"},
				},
				Response: oauth2.Token{
					TokenType:   "Some",
					AccessToken: "cde",
				},
			},
		},
	}, map[string]string{
		"Authorization": "Some cde",
	})
}

func TestM2mNotSupported(t *testing.T) {
	_, err := authenticateRequest(&Config{
		Host:         "a",
		ClientID:     "b",
		ClientSecret: "c",
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Status:   404,
				Response: `...`,
			},
		},
	})
	require.ErrorIs(t, err, oauth.ErrOAuthNotSupported)
}
