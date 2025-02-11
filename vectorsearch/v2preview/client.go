// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearchpreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type VectorSearchEndpointsPreviewClient struct {
	VectorSearchEndpointsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewVectorSearchEndpointsPreviewClient(cfg *config.Config) (*VectorSearchEndpointsPreviewClient, error) {
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

	return &VectorSearchEndpointsPreviewClient{
		Config:                                cfg,
		apiClient:                             apiClient,
		VectorSearchEndpointsPreviewInterface: NewVectorSearchEndpointsPreview(databricksClient),
	}, nil
}

type VectorSearchIndexesPreviewClient struct {
	VectorSearchIndexesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewVectorSearchIndexesPreviewClient(cfg *config.Config) (*VectorSearchIndexesPreviewClient, error) {
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

	return &VectorSearchIndexesPreviewClient{
		Config:                              cfg,
		apiClient:                           apiClient,
		VectorSearchIndexesPreviewInterface: NewVectorSearchIndexesPreview(databricksClient),
	}, nil
}
