// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mlpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ExperimentsPreviewClient struct {
	ExperimentsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewExperimentsPreviewClient(cfg *config.Config) (*ExperimentsPreviewClient, error) {
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

	return &ExperimentsPreviewClient{
		Config:                      cfg,
		apiClient:                   apiClient,
		ExperimentsPreviewInterface: NewExperimentsPreview(databricksClient),
	}, nil
}

type ModelRegistryPreviewClient struct {
	ModelRegistryPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewModelRegistryPreviewClient(cfg *config.Config) (*ModelRegistryPreviewClient, error) {
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

	return &ModelRegistryPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		ModelRegistryPreviewInterface: NewModelRegistryPreview(databricksClient),
	}, nil
}
