// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Vector Search Endpoints, Vector Search Indexes, etc.
package vectorsearch

import (
	"context"
)

// **Endpoint**: Represents the compute resources to host vector search indexes.
type VectorSearchEndpointsAPI struct {
	vectorSearchEndpointsImpl
}

// Delete an endpoint.
func (a *VectorSearchEndpointsAPI) DeleteEndpointByEndpointName(ctx context.Context, endpointName string) (*DeleteEndpointResponse, error) {
	return a.vectorSearchEndpointsImpl.DeleteEndpoint(ctx, DeleteEndpointRequest{
		EndpointName: endpointName,
	})
}

// Get an endpoint.
func (a *VectorSearchEndpointsAPI) GetEndpointByEndpointName(ctx context.Context, endpointName string) (*EndpointInfo, error) {
	return a.vectorSearchEndpointsImpl.GetEndpoint(ctx, GetEndpointRequest{
		EndpointName: endpointName,
	})
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
type VectorSearchIndexesAPI struct {
	vectorSearchIndexesImpl
}

// Delete an index.
//
// Delete an index.
func (a *VectorSearchIndexesAPI) DeleteIndexByIndexName(ctx context.Context, indexName string) (*DeleteIndexResponse, error) {
	return a.vectorSearchIndexesImpl.DeleteIndex(ctx, DeleteIndexRequest{
		IndexName: indexName,
	})
}

// Get an index.
//
// Get an index.
func (a *VectorSearchIndexesAPI) GetIndexByIndexName(ctx context.Context, indexName string) (*VectorIndex, error) {
	return a.vectorSearchIndexesImpl.GetIndex(ctx, GetIndexRequest{
		IndexName: indexName,
	})
}
