package u2m

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestBasicOAuthClient_GetAccountOAuthEndpoints(t *testing.T) {
	c := &BasicOAuthEndpointSupplier{}
	s, err := c.GetAccountOAuthEndpoints(context.Background(), "https://abc", "xyz")
	assert.NoError(t, err)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/authorize", s.AuthorizationEndpoint)
	assert.Equal(t, "https://abc/oidc/accounts/xyz/v1/token", s.TokenEndpoint)
}

func TestGetWorkspaceOAuthEndpoints(t *testing.T) {
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
	c := &BasicOAuthEndpointSupplier{Client: p}
	endpoints, err := c.GetWorkspaceOAuthEndpoints(context.Background(), "https://abc")
	assert.NoError(t, err)
	assert.Equal(t, "a", endpoints.AuthorizationEndpoint)
	assert.Equal(t, "b", endpoints.TokenEndpoint)
}
