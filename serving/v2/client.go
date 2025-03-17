// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ServingEndpointsClient struct {
	ServingEndpointsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServingEndpointsClient{
		apiClient:                 databricksClient.ApiClient(),
		ServingEndpointsInterface: NewServingEndpoints(databricksClient),
	}, nil
}

type ServingEndpointsDataPlaneClient struct {
	ServingEndpointsDataPlaneInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServingEndpointsDataPlaneClient{
		apiClient:                          databricksClient.ApiClient(),
		ServingEndpointsDataPlaneInterface: NewServingEndpointsDataPlane(databricksClient),
	}, nil
}
