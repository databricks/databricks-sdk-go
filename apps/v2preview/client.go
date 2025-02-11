// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package appspreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type AppsPreviewClient struct {
	AppsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewAppsPreviewClient(cfg *config.Config) (*AppsPreviewClient, error) {
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &AppsPreviewClient{
		Config:               cfg,
		apiClient:            apiClient,
		AppsPreviewInterface: NewAppsPreview(databricksClient),
	}, nil
}
