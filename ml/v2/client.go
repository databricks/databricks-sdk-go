// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// Experiments are the primary unit of organization in MLflow; all MLflow runs
// belong to an experiment. Each experiment lets you visualize, search, and
// compare runs, as well as download run artifacts or metadata for analysis in
// other tools. Experiments are maintained in a Databricks hosted MLflow
// tracking server.
//
// Experiments are located in the workspace file tree. You manage experiments
// using the same tools you use to manage other workspace objects such as
// folders, notebooks, and libraries.
type ExperimentsClient struct {
	experimentsBaseClient
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
		experimentsBaseClient: experimentsBaseClient{
			experimentsImpl: experimentsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The Forecasting API allows you to create and get serverless forecasting
// experiments
type ForecastingClient struct {
	forecastingBaseClient
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
		forecastingBaseClient: forecastingBaseClient{
			forecastingImpl: forecastingImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Note: This API reference documents APIs for the Workspace Model Registry.
// Databricks recommends using [Models in Unity
// Catalog](/api/workspace/registeredmodels) instead. Models in Unity Catalog
// provides centralized model governance, cross-workspace access, lineage, and
// deployment. Workspace Model Registry will be deprecated in the future.
//
// The Workspace Model Registry is a centralized model repository and a UI and
// set of APIs that enable you to manage the full lifecycle of MLflow Models.
type ModelRegistryClient struct {
	modelRegistryBaseClient
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
		modelRegistryBaseClient: modelRegistryBaseClient{
			modelRegistryImpl: modelRegistryImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
