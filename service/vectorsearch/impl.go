// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/service/vectorsearch/vectorsearchpb"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just VectorSearchEndpoints API methods
type vectorSearchEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *vectorSearchEndpointsImpl) CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error) {
	requestPb, pbErr := CreateEndpointToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var endpointInfoPb vectorsearchpb.EndpointInfoPb
	path := "/api/2.0/vector-search/endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&endpointInfoPb,
	)
	resp, err := EndpointInfoFromPb(&endpointInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchEndpointsImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error {
	requestPb, pbErr := DeleteEndpointRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", request.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *vectorSearchEndpointsImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error) {
	requestPb, pbErr := GetEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var endpointInfoPb vectorsearchpb.EndpointInfoPb
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", request.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&endpointInfoPb,
	)
	resp, err := EndpointInfoFromPb(&endpointInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// List all vector search endpoints in the workspace.
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

// List all vector search endpoints in the workspace.
func (a *vectorSearchEndpointsImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]EndpointInfo, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[EndpointInfo](ctx, iterator)
}

func (a *vectorSearchEndpointsImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointResponse, error) {
	requestPb, pbErr := ListEndpointsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listEndpointResponsePb vectorsearchpb.ListEndpointResponsePb
	path := "/api/2.0/vector-search/endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listEndpointResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListEndpointResponseFromPb(&listEndpointResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchEndpointsImpl) UpdateEndpointBudgetPolicy(ctx context.Context, request PatchEndpointBudgetPolicyRequest) (*PatchEndpointBudgetPolicyResponse, error) {
	requestPb, pbErr := PatchEndpointBudgetPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var patchEndpointBudgetPolicyResponsePb vectorsearchpb.PatchEndpointBudgetPolicyResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v/budget-policy", request.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&patchEndpointBudgetPolicyResponsePb,
	)
	resp, err := PatchEndpointBudgetPolicyResponseFromPb(&patchEndpointBudgetPolicyResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchEndpointsImpl) UpdateEndpointCustomTags(ctx context.Context, request UpdateEndpointCustomTagsRequest) (*UpdateEndpointCustomTagsResponse, error) {
	requestPb, pbErr := UpdateEndpointCustomTagsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateEndpointCustomTagsResponsePb vectorsearchpb.UpdateEndpointCustomTagsResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v/tags", request.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		&updateEndpointCustomTagsResponsePb,
	)
	resp, err := UpdateEndpointCustomTagsResponseFromPb(&updateEndpointCustomTagsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just VectorSearchIndexes API methods
type vectorSearchIndexesImpl struct {
	client *client.DatabricksClient
}

func (a *vectorSearchIndexesImpl) CreateIndex(ctx context.Context, request CreateVectorIndexRequest) (*VectorIndex, error) {
	requestPb, pbErr := CreateVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vectorIndexPb vectorsearchpb.VectorIndexPb
	path := "/api/2.0/vector-search/indexes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&vectorIndexPb,
	)
	resp, err := VectorIndexFromPb(&vectorIndexPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error) {
	requestPb, pbErr := DeleteDataVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDataVectorIndexResponsePb vectorsearchpb.DeleteDataVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/delete-data", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		&deleteDataVectorIndexResponsePb,
	)
	resp, err := DeleteDataVectorIndexResponseFromPb(&deleteDataVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) DeleteIndex(ctx context.Context, request DeleteIndexRequest) error {
	requestPb, pbErr := DeleteIndexRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *vectorSearchIndexesImpl) GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error) {
	requestPb, pbErr := GetIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vectorIndexPb vectorsearchpb.VectorIndexPb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&vectorIndexPb,
	)
	resp, err := VectorIndexFromPb(&vectorIndexPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// List all indexes in the given endpoint.
func (a *vectorSearchIndexesImpl) ListIndexesAll(ctx context.Context, request ListIndexesRequest) ([]MiniVectorIndex, error) {
	iterator := a.ListIndexes(ctx, request)
	return listing.ToSlice[MiniVectorIndex](ctx, iterator)
}

func (a *vectorSearchIndexesImpl) internalListIndexes(ctx context.Context, request ListIndexesRequest) (*ListVectorIndexesResponse, error) {
	requestPb, pbErr := ListIndexesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listVectorIndexesResponsePb vectorsearchpb.ListVectorIndexesResponsePb
	path := "/api/2.0/vector-search/indexes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&listVectorIndexesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := ListVectorIndexesResponseFromPb(&listVectorIndexesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error) {
	requestPb, pbErr := QueryVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryVectorIndexResponsePb vectorsearchpb.QueryVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&queryVectorIndexResponsePb,
	)
	resp, err := QueryVectorIndexResponseFromPb(&queryVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error) {
	requestPb, pbErr := QueryVectorIndexNextPageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryVectorIndexResponsePb vectorsearchpb.QueryVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query-next-page", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&queryVectorIndexResponsePb,
	)
	resp, err := QueryVectorIndexResponseFromPb(&queryVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error) {
	requestPb, pbErr := ScanVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var scanVectorIndexResponsePb vectorsearchpb.ScanVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/scan", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&scanVectorIndexResponsePb,
	)
	resp, err := ScanVectorIndexResponseFromPb(&scanVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) SyncIndex(ctx context.Context, request SyncIndexRequest) error {

	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/sync", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)

	return err
}

func (a *vectorSearchIndexesImpl) UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error) {
	requestPb, pbErr := UpsertDataVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var upsertDataVectorIndexResponsePb vectorsearchpb.UpsertDataVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/upsert-data", request.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&upsertDataVectorIndexResponsePb,
	)
	resp, err := UpsertDataVectorIndexResponseFromPb(&upsertDataVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
