// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just VectorSearchEndpoints API methods
type vectorSearchEndpointsImpl struct {
	client *httpclient.ApiClient
}

func (a *vectorSearchEndpointsImpl) CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error) {
	var endpointInfo EndpointInfo
	path := "/api/2.0/vector-search/endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &endpointInfo)
	return &endpointInfo, err
}

func (a *vectorSearchEndpointsImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (*DeleteEndpointResponse, error) {
	var deleteEndpointResponse DeleteEndpointResponse
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", request.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deleteEndpointResponse)
	return &deleteEndpointResponse, err
}

func (a *vectorSearchEndpointsImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error) {
	var endpointInfo EndpointInfo
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", request.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &endpointInfo)
	return &endpointInfo, err
}

// List all endpoints.
func (a *vectorSearchEndpointsImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[EndpointInfo] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEndpoints(ctx, req)
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
func (a *vectorSearchEndpointsImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]EndpointInfo, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}
func (a *vectorSearchEndpointsImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointResponse, error) {
	var listEndpointResponse ListEndpointResponse
	path := "/api/2.0/vector-search/endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listEndpointResponse)
	return &listEndpointResponse, err
}

// unexported type that holds implementations of just VectorSearchIndexes API methods
type vectorSearchIndexesImpl struct {
	client *httpclient.ApiClient
}

func (a *vectorSearchIndexesImpl) CreateIndex(ctx context.Context, request CreateVectorIndexRequest) (*CreateVectorIndexResponse, error) {
	var createVectorIndexResponse CreateVectorIndexResponse
	path := "/api/2.0/vector-search/indexes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &createVectorIndexResponse)
	return &createVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error) {
	var deleteDataVectorIndexResponse DeleteDataVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/delete-data", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &deleteDataVectorIndexResponse)
	return &deleteDataVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) DeleteIndex(ctx context.Context, request DeleteIndexRequest) (*DeleteIndexResponse, error) {
	var deleteIndexResponse DeleteIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deleteIndexResponse)
	return &deleteIndexResponse, err
}

func (a *vectorSearchIndexesImpl) GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error) {
	var vectorIndex VectorIndex
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &vectorIndex)
	return &vectorIndex, err
}

// List indexes.
//
// List all indexes in the given endpoint.
func (a *vectorSearchIndexesImpl) ListIndexes(ctx context.Context, request ListIndexesRequest) listing.Iterator[MiniVectorIndex] {

	getNextPage := func(ctx context.Context, req ListIndexesRequest) (*ListVectorIndexesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListIndexes(ctx, req)
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
func (a *vectorSearchIndexesImpl) ListIndexesAll(ctx context.Context, request ListIndexesRequest) ([]MiniVectorIndex, error) {
	iterator := a.ListIndexes(ctx, request)
	return listing.ToSlice[MiniVectorIndex](ctx, iterator)
}
func (a *vectorSearchIndexesImpl) internalListIndexes(ctx context.Context, request ListIndexesRequest) (*ListVectorIndexesResponse, error) {
	var listVectorIndexesResponse ListVectorIndexesResponse
	path := "/api/2.0/vector-search/indexes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listVectorIndexesResponse)
	return &listVectorIndexesResponse, err
}

func (a *vectorSearchIndexesImpl) QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error) {
	var queryVectorIndexResponse QueryVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &queryVectorIndexResponse)
	return &queryVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error) {
	var queryVectorIndexResponse QueryVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query-next-page", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &queryVectorIndexResponse)
	return &queryVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error) {
	var scanVectorIndexResponse ScanVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/scan", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &scanVectorIndexResponse)
	return &scanVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) SyncIndex(ctx context.Context, request SyncIndexRequest) (*SyncIndexResponse, error) {
	var syncIndexResponse SyncIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/sync", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, nil, &syncIndexResponse)
	return &syncIndexResponse, err
}

func (a *vectorSearchIndexesImpl) UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error) {
	var upsertDataVectorIndexResponse UpsertDataVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/upsert-data", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request, &upsertDataVectorIndexResponse)
	return &upsertDataVectorIndexResponse, err
}

func do(
	client *httpclient.ApiClient,
	ctx context.Context,
	method string,
	path string,
	headers map[string]string,
	queryParams map[string]any,
	request any,
	response any,
	visitors ...func(*http.Request) error,
) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithQueryParameters(queryParams))
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))

	// Remove extra `/` from path for files API. Once the OpenAPI spec doesn't
	// include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}

	return client.Do(ctx, method, path, opts...)
}
