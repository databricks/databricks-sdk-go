package client

import (
	"context"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
	"golang.org/x/oauth2"
)

type DatabricksClient struct {
	accountID string
	client    *httpclient.ApiClient
}

func New(cfg *config.Config) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	clientCfg, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &DatabricksClient{
		accountID: cfg.AccountID,
		client:    httpclient.NewApiClient(clientCfg),
	}, nil
}

// ConfiguredAccountID returns Databricks Account ID if it's provided in config,
// empty string otherwise
func (c *DatabricksClient) ConfiguredAccountID() string {
	return c.accountID
}

// ApiClient returns the inner Api Client.
func (c *DatabricksClient) ApiClient() *httpclient.ApiClient {
	return c.client
}

// Do sends an HTTP request against path.
func (c *DatabricksClient) Do(
	ctx context.Context,
	method string,
	path string,
	headers map[string]string,
	queryParams map[string]any,
	request any,
	response any,
	visitors ...func(*http.Request) error,
) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithQueryParameters(queryParams))
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))

	// Remove extra `/` from path for files API. Once the OpenAPI spec doesn't
	// include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}

	return c.client.Do(ctx, method, path, opts...)
}

// GetOAuthToken returns a new OAuth token using the provided token. The token
// must be a JWT token. The resulting token is scoped to the authorization
// details provided.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (c *DatabricksClient) GetOAuthToken(ctx context.Context, authDetails string, token *oauth2.Token) (*oauth2.Token, error) {
	return c.client.GetOAuthToken(ctx, authDetails, token)
}
