package oauth

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestOidcEndpointsForAccounts(t *testing.T) {
	s, err := GetAccountOAuthEndpoints(context.Background(), "https://abc", "xyz")
	assert.NoError(t, err)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/authorize", s.AuthorizationEndpoint)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/token", s.TokenEndpoint)
}

func TestOidcForWorkspace(t *testing.T) {
	p := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.MappingTransport{
			"GET /oidc/.well-known/oauth-authorization-server": {
				Status: 200,
				Response: map[string]string{
					"authorization_endpoint": "a",
					"token_endpoint":         "b",
				},
			},
		},
	})
	endpoints, err := GetWorkspaceOAuthEndpoints(context.Background(), p, "https://abc")
	assert.NoError(t, err)
	assert.Equal(t, "a", endpoints.AuthorizationEndpoint)
	assert.Equal(t, "b", endpoints.TokenEndpoint)
}
