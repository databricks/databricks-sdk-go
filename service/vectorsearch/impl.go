// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vectorsearch

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just VectorSearchEndpoints API methods
type vectorSearchEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *vectorSearchEndpointsImpl) CreateEndpoint(ctx context.Context, request CreateEndpoint) (*EndpointInfo, error) {

	requestPb, pbErr := createEndpointToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var endpointInfoPb endpointInfoPb
	path := "/api/2.0/vector-search/endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&endpointInfoPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := endpointInfoFromPb(&endpointInfoPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchEndpointsImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error {

	requestPb, pbErr := deleteEndpointRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", requestPb.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *vectorSearchEndpointsImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*EndpointInfo, error) {

	requestPb, pbErr := getEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var endpointInfoPb endpointInfoPb
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v", requestPb.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&endpointInfoPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := endpointInfoFromPb(&endpointInfoPb)
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

	requestPb, pbErr := listEndpointsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listEndpointResponsePb listEndpointResponsePb
	path := "/api/2.0/vector-search/endpoints"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listEndpointResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listEndpointResponseFromPb(&listEndpointResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchEndpointsImpl) UpdateEndpointBudgetPolicy(ctx context.Context, request PatchEndpointBudgetPolicyRequest) (*PatchEndpointBudgetPolicyResponse, error) {

	requestPb, pbErr := patchEndpointBudgetPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var patchEndpointBudgetPolicyResponsePb patchEndpointBudgetPolicyResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v/budget-policy", requestPb.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&patchEndpointBudgetPolicyResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := patchEndpointBudgetPolicyResponseFromPb(&patchEndpointBudgetPolicyResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchEndpointsImpl) UpdateEndpointCustomTags(ctx context.Context, request UpdateEndpointCustomTagsRequest) (*UpdateEndpointCustomTagsResponse, error) {

	requestPb, pbErr := updateEndpointCustomTagsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateEndpointCustomTagsResponsePb updateEndpointCustomTagsResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/endpoints/%v/tags", requestPb.EndpointName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateEndpointCustomTagsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := updateEndpointCustomTagsResponseFromPb(&updateEndpointCustomTagsResponsePb)
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

	requestPb, pbErr := createVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vectorIndexPb vectorIndexPb
	path := "/api/2.0/vector-search/indexes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&vectorIndexPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := vectorIndexFromPb(&vectorIndexPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) DeleteDataVectorIndex(ctx context.Context, request DeleteDataVectorIndexRequest) (*DeleteDataVectorIndexResponse, error) {

	requestPb, pbErr := deleteDataVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var deleteDataVectorIndexResponsePb deleteDataVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/delete-data", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteDataVectorIndexResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := deleteDataVectorIndexResponseFromPb(&deleteDataVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) DeleteIndex(ctx context.Context, request DeleteIndexRequest) error {

	requestPb, pbErr := deleteIndexRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *vectorSearchIndexesImpl) GetIndex(ctx context.Context, request GetIndexRequest) (*VectorIndex, error) {

	requestPb, pbErr := getIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vectorIndexPb vectorIndexPb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&vectorIndexPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := vectorIndexFromPb(&vectorIndexPb)
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

	requestPb, pbErr := listIndexesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listVectorIndexesResponsePb listVectorIndexesResponsePb
	path := "/api/2.0/vector-search/indexes"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listVectorIndexesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listVectorIndexesResponseFromPb(&listVectorIndexesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) QueryIndex(ctx context.Context, request QueryVectorIndexRequest) (*QueryVectorIndexResponse, error) {

	requestPb, pbErr := queryVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryVectorIndexResponsePb queryVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryVectorIndexResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryVectorIndexResponseFromPb(&queryVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) QueryNextPage(ctx context.Context, request QueryVectorIndexNextPageRequest) (*QueryVectorIndexResponse, error) {

	requestPb, pbErr := queryVectorIndexNextPageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var queryVectorIndexResponsePb queryVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/query-next-page", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&queryVectorIndexResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := queryVectorIndexResponseFromPb(&queryVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) ScanIndex(ctx context.Context, request ScanVectorIndexRequest) (*ScanVectorIndexResponse, error) {

	requestPb, pbErr := scanVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var scanVectorIndexResponsePb scanVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/scan", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&scanVectorIndexResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := scanVectorIndexResponseFromPb(&scanVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vectorSearchIndexesImpl) SyncIndex(ctx context.Context, request SyncIndexRequest) error {

	requestPb, pbErr := syncIndexRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/sync", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *vectorSearchIndexesImpl) UpsertDataVectorIndex(ctx context.Context, request UpsertDataVectorIndexRequest) (*UpsertDataVectorIndexResponse, error) {

	requestPb, pbErr := upsertDataVectorIndexRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var upsertDataVectorIndexResponsePb upsertDataVectorIndexResponsePb
	path := fmt.Sprintf("/api/2.0/vector-search/indexes/%v/upsert-data", requestPb.IndexName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&upsertDataVectorIndexResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := upsertDataVectorIndexResponseFromPb(&upsertDataVectorIndexResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
