package client

import (
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
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
