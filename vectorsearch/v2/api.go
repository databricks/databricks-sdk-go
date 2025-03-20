// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Vector Search Endpoints, Vector Search Indexes, etc.
package vectorsearch

import (
	"context"
)

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
