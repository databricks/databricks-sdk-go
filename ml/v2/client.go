// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ExperimentsClient struct {
	ExperimentsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ExperimentsClient{
		ExperimentsAPI: ExperimentsAPI{
			experimentsImpl: experimentsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ForecastingClient struct {
	ForecastingAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ForecastingClient{
		ForecastingAPI: ForecastingAPI{
			forecastingImpl: forecastingImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ModelRegistryClient struct {
	ModelRegistryAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ModelRegistryClient{
		ModelRegistryAPI: ModelRegistryAPI{
			modelRegistryImpl: modelRegistryImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
