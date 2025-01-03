package httpclient

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestOidcEndpointsForAccounts(t *testing.T) {
	p := NewApiClient(ClientConfig{})
	s, err := p.GetOidcEndpoints(context.Background(), "https://abc", "xyz")
	assert.NoError(t, err)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/authorize", s.AuthorizationEndpoint)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/token", s.TokenEndpoint)
}

func TestOidcForWorkspace(t *testing.T) {
	p := NewApiClient(ClientConfig{
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
	endpoints, err := p.GetOidcEndpoints(context.Background(), "https://abc", "")
	assert.NoError(t, err)
	assert.Equal(t, "a", endpoints.AuthorizationEndpoint)
	assert.Equal(t, "b", endpoints.TokenEndpoint)
}
