// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ServingEndpointsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ServingEndpoints ServingEndpointsInterface
}

func NewServingEndpointsClient(cfg *config.Config) (*ServingEndpointsClient, error) {
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

	return &ServingEndpointsClient{
		cfg:              cfg,
		apiClient:        apiClient,
		ServingEndpoints: NewServingEndpoints(databricksClient),
	}, nil
}

type ServingEndpointsDataPlaneClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ServingEndpointsDataPlane ServingEndpointsDataPlaneInterface
}

func NewServingEndpointsDataPlaneClient(cfg *config.Config) (*ServingEndpointsDataPlaneClient, error) {
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

	return &ServingEndpointsDataPlaneClient{
		cfg:                       cfg,
		apiClient:                 apiClient,
		ServingEndpointsDataPlane: NewServingEndpointsDataPlane(databricksClient),
	}, nil
}
