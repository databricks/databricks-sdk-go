// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
)

// **Endpoint**: Represents the compute resources to host vector search indexes.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type VectorSearchEndpointsService interface {

	// Create a new endpoint.
	CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error)

	// Delete a vector search endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Get details for a single vector search endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error)

	// List all vector search endpoints in the workspace.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointResponse, error)

	// Update the budget policy of an endpoint
	UpdateEndpointBudgetPolicy(ctx context.Context, request PatchEndpointBudgetPolicyRequest) (*PatchEndpointBudgetPolicyResponse, error)

	// Update the custom tags of an endpoint.
	UpdateEndpointCustomTags(ctx context.Context, request UpdateEndpointCustomTagsRequest) (*UpdateEndpointCustomTagsResponse, error)
}

// **Index**: An efficient representation of your embedding vectors that
// supports real-time and efficient approximate nearest neighbor (ANN) search
// queries.
//
// There are 2 types of Vector Search indexes: - **Delta Sync Index**: An index
// that automatically syncs with a source Delta Table, automatically and
// incrementally updating the index as the underlying data in the Delta Table
// changes. - **Direct Vector Access Index**: An index that supports direct read
// and write of vectors and metadata through our REST and SDK APIs. With this
// model, the user manages index updates.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type VectorSearchIndexesService interface {

	// Create a new index.
	CreateIndex(ctx context.Context, request CreateVectorIndexRequest) (*VectorIndex, error)

	// Handles the deletion of data from a specified vector index.
	DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error)

	// Delete an index.
	DeleteIndex(ctx context.Context, request DeleteIndexRequest) error

	// Get an index.
	GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error)

	// List all indexes in the given endpoint.
	ListIndexes(ctx context.Context, request ListIndexesRequest) (*ListVectorIndexesResponse, error)

	// Query the specified vector index.
	QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error)

	// Use `next_page_token` returned from previous `QueryVectorIndex` or
	// `QueryVectorIndexNextPage` request to fetch next page of results.
	QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error)

	// Scan the specified vector index and return the first `num_results`
	// entries after the exclusive `primary_key`.
	ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error)

	// Triggers a synchronization process for a specified vector index.
	SyncIndex(ctx context.Context, request SyncIndexRequest) error

	// Handles the upserting of data into a specified vector index.
	UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error)
}
