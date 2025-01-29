// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just VectorSearchEndpoints API methods
type vectorSearchEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *vectorSearchEndpointsImpl) CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error) {
	var endpointInfo EndpointInfo
	path := "/api/2.0/vector-search/endpoints"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &endpointInfo)
	return &endpointInfo, err
}

func (a *vectorSearchEndpointsImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error {
	var deleteEndpointResponse DeleteEndpointResponse
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", request.EndpointName)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteEndpointResponse)
	return err
}

func (a *vectorSearchEndpointsImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error) {
	var endpointInfo EndpointInfo
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", request.EndpointName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &endpointInfo)
	return &endpointInfo, err
}

func (a *vectorSearchEndpointsImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointResponse, error) {
	var listEndpointResponse ListEndpointResponse
	path := "/api/2.0/vector-search/endpoints"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listEndpointResponse)
	return &listEndpointResponse, err
}

// unexported type that holds implementations of just VectorSearchIndexes API methods
type vectorSearchIndexesImpl struct {
	client *client.DatabricksClient
}

func (a *vectorSearchIndexesImpl) CreateIndex(ctx context.Context, request CreateVectorIndexRequest) (*CreateVectorIndexResponse, error) {
	var createVectorIndexResponse CreateVectorIndexResponse
	path := "/api/2.0/vector-search/indexes"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createVectorIndexResponse)
	return &createVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error) {
	var deleteDataVectorIndexResponse DeleteDataVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/delete-data", request.IndexName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteDataVectorIndexResponse)
	return &deleteDataVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) DeleteIndex(ctx context.Context, request DeleteIndexRequest) error {
	var deleteIndexResponse DeleteIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", request.IndexName)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteIndexResponse)
	return err
}

func (a *vectorSearchIndexesImpl) GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error) {
	var vectorIndex VectorIndex
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", request.IndexName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &vectorIndex)
	return &vectorIndex, err
}

func (a *vectorSearchIndexesImpl) ListIndexes(ctx context.Context, request ListIndexesRequest) (*ListVectorIndexesResponse, error) {
	var listVectorIndexesResponse ListVectorIndexesResponse
	path := "/api/2.0/vector-search/indexes"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listVectorIndexesResponse)
	return &listVectorIndexesResponse, err
}

func (a *vectorSearchIndexesImpl) QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error) {
	var queryVectorIndexResponse QueryVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query", request.IndexName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &queryVectorIndexResponse)
	return &queryVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error) {
	var queryVectorIndexResponse QueryVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query-next-page", request.IndexName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &queryVectorIndexResponse)
	return &queryVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error) {
	var scanVectorIndexResponse ScanVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/scan", request.IndexName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &scanVectorIndexResponse)
	return &scanVectorIndexResponse, err
}

func (a *vectorSearchIndexesImpl) SyncIndex(ctx context.Context, request SyncIndexRequest) error {
	var syncIndexResponse SyncIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/sync", request.IndexName)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &syncIndexResponse)
	return err
}

func (a *vectorSearchIndexesImpl) UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error) {
	var upsertDataVectorIndexResponse UpsertDataVectorIndexResponse
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/upsert-data", request.IndexName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &upsertDataVectorIndexResponse)
	return &upsertDataVectorIndexResponse, err
}
