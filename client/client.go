package client

import (
	"context"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
)

func New(cfg *config.Config) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	client, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	return &DatabricksClient{
		Config: cfg,
		client: client,
	}, nil
}

func NewWithClient(cfg *config.Config, client *httpclient.ApiClient) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	return &DatabricksClient{
		Config: cfg,
		client: client,
	}, nil
}

type DatabricksClient struct {
	Config *config.Config
	client *httpclient.ApiClient
}

// ConfiguredAccountID returns Databricks Account ID if it's provided in config,
// empty string otherwise
func (c *DatabricksClient) ConfiguredAccountID() string {
	return c.Config.AccountID
}

// Do sends an HTTP request against path.
func (c *DatabricksClient) Do(ctx context.Context, method, path string,
	headers map[string]string, request, response any,
	visitors ...func(*http.Request) error) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))
	// Remove extra `/` from path for files API.
	// Once the OpenAPI spec doesn't include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}
	return c.client.Do(ctx, method, path, opts...)
}
