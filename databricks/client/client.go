package client

import (
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type DatabricksClient struct {
	client *httpclient.ApiClient
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
		client: httpclient.NewApiClient(clientCfg),
	}, nil
}

// ApiClient returns the inner Api Client.
func (c *DatabricksClient) ApiClient() *httpclient.ApiClient {
	return c.client
}
