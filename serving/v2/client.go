// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ServingEndpointsClient struct {
	ServingEndpointsAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServingEndpointsClient{
		ServingEndpointsAPI: ServingEndpointsAPI{
			servingEndpointsImpl: servingEndpointsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ServingEndpointsDataPlaneClient struct {
	ServingEndpointsDataPlaneAPI
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ServingEndpointsDataPlaneClient{
		ServingEndpointsDataPlaneAPI: ServingEndpointsDataPlaneAPI{
			servingEndpointsDataPlaneImpl: servingEndpointsDataPlaneImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
