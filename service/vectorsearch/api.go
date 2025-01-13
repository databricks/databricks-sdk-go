// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Vector Search Endpoints, Vector Search Indexes, etc.
package vectorsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type VectorSearchEndpointsInterface interface {

	// WaitGetEndpointVectorSearchEndpointOnline repeatedly calls [VectorSearchEndpointsAPI.GetEndpoint] and waits to reach ONLINE state
	WaitGetEndpointVectorSearchEndpointOnline(ctx context.Context, endpointName string,
		timeout time.Duration, callback func(*EndpointInfo)) (*EndpointInfo, error)

	// Create an endpoint.
	//
	// Create a new endpoint.
	CreateEndpoint(ctx context.Context, createEndpoint CreateEndpoint) (*WaitGetEndpointVectorSearchEndpointOnline[EndpointInfo], error)

	// Calls [VectorSearchEndpointsAPIInterface.CreateEndpoint] and waits to reach ONLINE state
	//
	// You can override the default timeout of 20 minutes by calling adding
	// retries.Timeout[EndpointInfo](60*time.Minute) functional option.
	//
	// Deprecated: use [VectorSearchEndpointsAPIInterface.CreateEndpoint].Get() or [VectorSearchEndpointsAPIInterface.WaitGetEndpointVectorSearchEndpointOnline]
	CreateEndpointAndWait(ctx context.Context, createEndpoint CreateEndpoint, options ...retries.Option[EndpointInfo]) (*EndpointInfo, error)

	// Delete an endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Delete an endpoint.
	DeleteEndpointByEndpointName(ctx context.Context, endpointName string) error

	// Get an endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error)

	// Get an endpoint.
	GetEndpointByEndpointName(ctx context.Context, endpointName string) (*EndpointInfo, error)

	// List all endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[EndpointInfo]

	// List all endpoints.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]EndpointInfo, error)
}

func NewVectorSearchEndpoints(client *client.DatabricksClient) *VectorSearchEndpointsAPI {
	return &VectorSearchEndpointsAPI{
		vectorSearchEndpointsImpl: vectorSearchEndpointsImpl{
			client: client,
		},
	}
}

// **Endpoint**: Represents the compute resources to host vector search indexes.
type VectorSearchEndpointsAPI struct {
	vectorSearchEndpointsImpl
}

// WaitGetEndpointVectorSearchEndpointOnline repeatedly calls [VectorSearchEndpointsAPI.GetEndpoint] and waits to reach ONLINE state
func (a *VectorSearchEndpointsAPI) WaitGetEndpointVectorSearchEndpointOnline(ctx context.Context, endpointName string,
	timeout time.Duration, callback func(*EndpointInfo)) (*EndpointInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[EndpointInfo](ctx, timeout, func() (*EndpointInfo, *retries.Err) {
		endpointInfo, err := a.GetEndpoint(ctx, GetEndpointRequest{
			EndpointName: endpointName,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(endpointInfo)
		}
		status := endpointInfo.EndpointStatus.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if endpointInfo.EndpointStatus != nil {
			statusMessage = endpointInfo.EndpointStatus.Message
		}
		switch status {
		case EndpointStatusStateOnline: // target state
			return endpointInfo, nil
		case EndpointStatusStateOffline:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				EndpointStatusStateOnline, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetEndpointVectorSearchEndpointOnline is a wrapper that calls [VectorSearchEndpointsAPI.WaitGetEndpointVectorSearchEndpointOnline] and waits to reach ONLINE state.
type WaitGetEndpointVectorSearchEndpointOnline[R any] struct {
	Response     *R
	EndpointName string `json:"endpoint_name"`
	Poll         func(time.Duration, func(*EndpointInfo)) (*EndpointInfo, error)
	callback     func(*EndpointInfo)
	timeout      time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetEndpointVectorSearchEndpointOnline[R]) OnProgress(callback func(*EndpointInfo)) *WaitGetEndpointVectorSearchEndpointOnline[R] {
	w.callback = callback
	return w
}

// Get the EndpointInfo with the default timeout of 20 minutes.
func (w *WaitGetEndpointVectorSearchEndpointOnline[R]) Get() (*EndpointInfo, error) {
	return w.Poll(w.timeout, w.callback)
}

// Get the EndpointInfo with custom timeout.
func (w *WaitGetEndpointVectorSearchEndpointOnline[R]) GetWithTimeout(timeout time.Duration) (*EndpointInfo, error) {
	return w.Poll(timeout, w.callback)
}

// Create an endpoint.
//
// Create a new endpoint.
func (a *VectorSearchEndpointsAPI) CreateEndpoint(ctx context.Context, createEndpoint CreateEndpoint) (*WaitGetEndpointVectorSearchEndpointOnline[EndpointInfo], error) {
	endpointInfo, err := a.vectorSearchEndpointsImpl.CreateEndpoint(ctx, createEndpoint)
	if err != nil {
		return nil, err
	}
	return &WaitGetEndpointVectorSearchEndpointOnline[EndpointInfo]{
		Response:     endpointInfo,
		EndpointName: endpointInfo.Name,
		Poll: func(timeout time.Duration, callback func(*EndpointInfo)) (*EndpointInfo, error) {
			return a.WaitGetEndpointVectorSearchEndpointOnline(ctx, endpointInfo.Name, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [VectorSearchEndpointsAPI.CreateEndpoint] and waits to reach ONLINE state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[EndpointInfo](60*time.Minute) functional option.
//
// Deprecated: use [VectorSearchEndpointsAPI.CreateEndpoint].Get() or [VectorSearchEndpointsAPI.WaitGetEndpointVectorSearchEndpointOnline]
func (a *VectorSearchEndpointsAPI) CreateEndpointAndWait(ctx context.Context, createEndpoint CreateEndpoint, options ...retries.Option[EndpointInfo]) (*EndpointInfo, error) {
	wait, err := a.CreateEndpoint(ctx, createEndpoint)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[EndpointInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *EndpointInfo) {
		for _, o := range options {
			o(&retries.Info[EndpointInfo]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete an endpoint.
func (a *VectorSearchEndpointsAPI) DeleteEndpointByEndpointName(ctx context.Context, endpointName string) error {
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

// List all endpoints.
//
// This method is generated by Databricks SDK Code Generator.
func (a *VectorSearchEndpointsAPI) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[EndpointInfo] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.vectorSearchEndpointsImpl.ListEndpoints(ctx, req)
	}
	getItems := func(resp *ListEndpointResponse) []EndpointInfo {
		return resp.Endpoints
	}
	getNextReq := func(resp *ListEndpointResponse) *ListEndpointsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List all endpoints.
//
// This method is generated by Databricks SDK Code Generator.
func (a *VectorSearchEndpointsAPI) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]EndpointInfo, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}

type VectorSearchIndexesInterface interface {

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

	// Delete an index.
	//
	// Delete an index.
	DeleteIndexByIndexName(ctx context.Context, indexName string) error

	// Get an index.
	//
	// Get an index.
	GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error)

	// Get an index.
	//
	// Get an index.
	GetIndexByIndexName(ctx context.Context, indexName string) (*VectorIndex, error)

	// List indexes.
	//
	// List all indexes in the given endpoint.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListIndexes(ctx context.Context, request ListIndexesRequest) listing.Iterator[MiniVectorIndex]

	// List indexes.
	//
	// List all indexes in the given endpoint.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListIndexesAll(ctx context.Context, request ListIndexesRequest) ([]MiniVectorIndex, error)

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
	// Scan the specified vector index and return the first `num_results` entries
	// after the exclusive `primary_key`.
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

func NewVectorSearchIndexes(client *client.DatabricksClient) *VectorSearchIndexesAPI {
	return &VectorSearchIndexesAPI{
		vectorSearchIndexesImpl: vectorSearchIndexesImpl{
			client: client,
		},
	}
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
func (a *VectorSearchIndexesAPI) DeleteIndexByIndexName(ctx context.Context, indexName string) error {
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

// List indexes.
//
// List all indexes in the given endpoint.
//
// This method is generated by Databricks SDK Code Generator.
func (a *VectorSearchIndexesAPI) ListIndexes(ctx context.Context, request ListIndexesRequest) listing.Iterator[MiniVectorIndex] {

	getNextPage := func(ctx context.Context, req ListIndexesRequest) (*ListVectorIndexesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.vectorSearchIndexesImpl.ListIndexes(ctx, req)
	}
	getItems := func(resp *ListVectorIndexesResponse) []MiniVectorIndex {
		return resp.VectorIndexes
	}
	getNextReq := func(resp *ListVectorIndexesResponse) *ListIndexesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List indexes.
//
// List all indexes in the given endpoint.
//
// This method is generated by Databricks SDK Code Generator.
func (a *VectorSearchIndexesAPI) ListIndexesAll(ctx context.Context, request ListIndexesRequest) ([]MiniVectorIndex, error) {
	iterator := a.ListIndexes(ctx, request)
	return listing.ToSlice[MiniVectorIndex](ctx, iterator)
}
