package config

import (
	"net/url"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
)

func TestM2mHappyFlow(t *testing.T) {
	assertHeaders(t, &Config{
		Host:         "a",
		ClientID:     "b",
		ClientSecret: "c",
		HTTPTransport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Response: oauthAuthorizationServer{
					AuthorizationEndpoint: "https://localhost:1234/dummy/auth",
					TokenEndpoint:         "https://localhost:1234/dummy/token",
				},
			},
			"POST /dummy/token": {
				ExpectedRequest: url.Values{
					"grant_type": []string{"client_credentials"},
					"scope":      []string{"all-apis"},
				},
				Response: `...`,
			},
		},
	}, map[string]string{
		"Authorization":                            "Bearer cde",
		"X-Databricks-Azure-Sp-Management-Token":   "def",
		"X-Databricks-Azure-Workspace-Resource-Id": "/a/b/c",
	})
}
