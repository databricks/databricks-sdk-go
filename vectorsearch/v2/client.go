// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type VectorSearchEndpointsClient struct {
	VectorSearchEndpointsInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VectorSearchEndpointsClient{
		VectorSearchEndpointsInterface: NewVectorSearchEndpoints(databricksClient),
	}, nil
}

type VectorSearchIndexesClient struct {
	VectorSearchIndexesInterface
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VectorSearchIndexesClient{
		VectorSearchIndexesInterface: NewVectorSearchIndexes(databricksClient),
	}, nil
}
