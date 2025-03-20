// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// **Endpoint**: Represents the compute resources to host vector search indexes.
type VectorSearchEndpointsClient struct {
	vectorSearchEndpointsBaseClient
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VectorSearchEndpointsClient{
		vectorSearchEndpointsBaseClient: vectorSearchEndpointsBaseClient{
			vectorSearchEndpointsImpl: vectorSearchEndpointsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// **Index**: An efficient representation of your embedding vectors that
// supports real-time and efficient approximate nearest neighbor (ANN) search
// queries.
//
// There are 2 types of Vector Search indexes: * **Delta Sync Index**: An index
// that automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. * **Direct Vector Access Index**: An index that supports direct read
// and write of vectors and metadata through our REST and SDK APIs. With this
// model, the user manages index updates.
type VectorSearchIndexesClient struct {
	vectorSearchIndexesBaseClient
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VectorSearchIndexesClient{
		vectorSearchIndexesBaseClient: vectorSearchIndexesBaseClient{
			vectorSearchIndexesImpl: vectorSearchIndexesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
