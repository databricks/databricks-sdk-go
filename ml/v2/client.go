// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ExperimentsClient struct {
	ExperimentsInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewExperimentsClient(cfg *config.Config) (*ExperimentsClient, error) {
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

	return &ExperimentsClient{
		Config:               cfg,
		apiClient:            apiClient,
		ExperimentsInterface: NewExperiments(databricksClient),
	}, nil
}

type ModelRegistryClient struct {
	ModelRegistryInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewModelRegistryClient(cfg *config.Config) (*ModelRegistryClient, error) {
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

	return &ModelRegistryClient{
		Config:                 cfg,
		apiClient:              apiClient,
		ModelRegistryInterface: NewModelRegistry(databricksClient),
	}, nil
}
