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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ExperimentsClient{
		apiClient:            databricksClient.ApiClient(),
		ExperimentsInterface: NewExperiments(databricksClient),
	}, nil
}

type ForecastingClient struct {
	ForecastingInterface
	apiClient *httpclient.ApiClient
}

func NewForecastingClient(cfg *config.Config) (*ForecastingClient, error) {
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

	return &ForecastingClient{
		apiClient:            databricksClient.ApiClient(),
		ForecastingInterface: NewForecasting(databricksClient),
	}, nil
}

type ModelRegistryClient struct {
	ModelRegistryInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ModelRegistryClient{
		apiClient:              databricksClient.ApiClient(),
		ModelRegistryInterface: NewModelRegistry(databricksClient),
	}, nil
}
