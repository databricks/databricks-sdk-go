// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type VectorSearchEndpointsClient struct {
	VectorSearchEndpointsInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewVectorSearchEndpointsClient(cfg *config.Config) (*VectorSearchEndpointsClient, error) {
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

	return &VectorSearchEndpointsClient{
		Config:                         cfg,
		apiClient:                      apiClient,
		VectorSearchEndpointsInterface: NewVectorSearchEndpoints(databricksClient),
	}, nil
}

type VectorSearchIndexesClient struct {
	VectorSearchIndexesInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewVectorSearchIndexesClient(cfg *config.Config) (*VectorSearchIndexesClient, error) {
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

	return &VectorSearchIndexesClient{
		Config:                       cfg,
		apiClient:                    apiClient,
		VectorSearchIndexesInterface: NewVectorSearchIndexes(databricksClient),
	}, nil
}
