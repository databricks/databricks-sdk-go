// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package servingpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ServingEndpointsClient struct {
	ServingEndpointsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ServingEndpointsClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		ServingEndpointsInterface: NewServingEndpoints(databricksClient),
	}, nil
}

type ServingEndpointsDataPlanePreviewClient struct {
	ServingEndpointsDataPlanePreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewServingEndpointsDataPlanePreviewClient(cfg *config.Config) (*ServingEndpointsDataPlanePreviewClient, error) {
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

	return &ServingEndpointsDataPlanePreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		ServingEndpointsDataPlanePreviewInterface: NewServingEndpointsDataPlanePreview(databricksClient),
	}, nil
}

type ServingEndpointsPreviewClient struct {
	ServingEndpointsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewServingEndpointsPreviewClient(cfg *config.Config) (*ServingEndpointsPreviewClient, error) {
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

	return &ServingEndpointsPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		ServingEndpointsPreviewInterface: NewServingEndpointsPreview(databricksClient),
	}, nil
}
