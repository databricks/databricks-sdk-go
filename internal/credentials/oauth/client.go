package oauth

import (
	"context"
	"net/http"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

// OAuthClient provides the http functionality needed for interacting with the
// Databricks OAuth APIs.
type OAuthClient interface {
	// GetHttpClient returns an HTTP client for OAuth2 requests.
	GetHttpClient(context.Context) *http.Client

	// GetWorkspaceOAuthEndpoints returns the OAuth2 endpoints for the workspace.
	GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error)

	// GetAccountOAuthEndpoints returns the OAuth2 endpoints for the account.
	GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error)
}

type BasicOAuthClient struct {
	client *httpclient.ApiClient
}

func (c *BasicOAuthClient) GetHttpClient(_ context.Context) *http.Client {
	return c.client.ToHttpClient()
}

func (c *BasicOAuthClient) GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error) {
	return GetWorkspaceOAuthEndpoints(ctx, c.client, workspaceHost)
}

func (c *BasicOAuthClient) GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error) {
	return GetAccountOAuthEndpoints(ctx, accountHost, accountId)
}
