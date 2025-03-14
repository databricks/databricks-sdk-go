// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package apps

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AppsClient struct {
	AppsInterface
	apiClient *httpclient.ApiClient
}

func NewAppsClient(cfg *config.Config) (*AppsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}

	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AppsClient{
		apiClient:     databricksClient.ApiClient(),
		AppsInterface: NewApps(databricksClient),
	}, nil
}
