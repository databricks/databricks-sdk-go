// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ExperimentsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Experiments ExperimentsInterface
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
		cfg:         cfg,
		apiClient:   apiClient,
		Experiments: NewExperiments(databricksClient),
	}, nil
}

type ModelRegistryClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ModelRegistry ModelRegistryInterface
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
		cfg:           cfg,
		apiClient:     apiClient,
		ModelRegistry: NewModelRegistry(databricksClient),
	}, nil
}
