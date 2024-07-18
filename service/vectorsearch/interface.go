// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
)

// **Endpoint**: Represents the compute resources to host vector search indexes.
type VectorSearchEndpointsService interface {

	// Create an endpoint.
	//
	// Create a new endpoint.
	CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error)

	// Delete an endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Get an endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error)

	// List all endpoints.
	//
	// Use ListEndpointsAll() to get all EndpointInfo instances, which will iterate over every result page.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointResponse, error)
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
type VectorSearchIndexesService interface {

	// Create an index.
	//
	// Create a new index.
	CreateIndex(ctx context.Context, request CreateVectorIndexRequest) (*CreateVectorIndexResponse, error)

	// Delete data from index.
	//
	// Handles the deletion of data from a specified vector index.
	DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error)

	// Delete an index.
	//
	// Delete an index.
	DeleteIndex(ctx context.Context, request DeleteIndexRequest) error

	// Get an index.
	//
	// Get an index.
	GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error)

	// List indexes.
	//
	// List all indexes in the given endpoint.
	//
	// Use ListIndexesAll() to get all MiniVectorIndex instances, which will iterate over every result page.
	ListIndexes(ctx context.Context, request ListIndexesRequest) (*ListVectorIndexesResponse, error)

	// Query an index.
	//
	// Query the specified vector index.
	QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error)

	// Query next page.
	//
	// Use `next_page_token` returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` request to fetch next page of results.
	QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error)

	// Scan an index.
	//
	// Scan the specified vector index and return the first `num_results`
	// entries after the exclusive `primary_key`.
	ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error)

	// Synchronize an index.
	//
	// Triggers a synchronization process for a specified vector index.
	SyncIndex(ctx context.Context, request SyncIndexRequest) error

	// Upsert data into an index.
	//
	// Handles the upserting of data into a specified vector index.
	UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error)
}
