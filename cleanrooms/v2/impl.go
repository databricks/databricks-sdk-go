// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just CleanRoomAssets API methods
type cleanRoomAssetsImpl struct {
	client *httpclient.ApiClient
}

func (a *cleanRoomAssetsImpl) Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request.Asset, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

func (a *cleanRoomAssetsImpl) Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) (*DeleteCleanRoomAssetResponse, error) {
	var deleteCleanRoomAssetResponse DeleteCleanRoomAssetResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.AssetFullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deleteCleanRoomAssetResponse)
	return &deleteCleanRoomAssetResponse, err
}

func (a *cleanRoomAssetsImpl) Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.AssetFullName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

// List assets.
func (a *cleanRoomAssetsImpl) List(ctx context.Context, request ListCleanRoomAssetsRequest) listing.Iterator[CleanRoomAsset] {

	getNextPage := func(ctx context.Context, req ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomAssetsResponse) []CleanRoomAsset {
		return resp.Assets
	}
	getNextReq := func(resp *ListCleanRoomAssetsResponse) *ListCleanRoomAssetsRequest {
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

// List assets.
func (a *cleanRoomAssetsImpl) ListAll(ctx context.Context, request ListCleanRoomAssetsRequest) ([]CleanRoomAsset, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomAsset](ctx, iterator)
}
func (a *cleanRoomAssetsImpl) internalList(ctx context.Context, request ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
	var listCleanRoomAssetsResponse ListCleanRoomAssetsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomAssetsResponse)
	return &listCleanRoomAssetsResponse, err
}

func (a *cleanRoomAssetsImpl) Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPatch, path, headers, queryParams, request.Asset, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

// unexported type that holds implementations of just CleanRoomTaskRuns API methods
type cleanRoomTaskRunsImpl struct {
	client *httpclient.ApiClient
}

// List notebook task runs.
//
// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) List(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) listing.Iterator[CleanRoomNotebookTaskRun] {

	getNextPage := func(ctx context.Context, req ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomNotebookTaskRunsResponse) []CleanRoomNotebookTaskRun {
		return resp.Runs
	}
	getNextReq := func(resp *ListCleanRoomNotebookTaskRunsResponse) *ListCleanRoomNotebookTaskRunsRequest {
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

// List notebook task runs.
//
// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) ListAll(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) ([]CleanRoomNotebookTaskRun, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomNotebookTaskRun](ctx, iterator)
}
func (a *cleanRoomTaskRunsImpl) internalList(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	var listCleanRoomNotebookTaskRunsResponse ListCleanRoomNotebookTaskRunsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/runs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomNotebookTaskRunsResponse)
	return &listCleanRoomNotebookTaskRunsResponse, err
}

// unexported type that holds implementations of just CleanRooms API methods
type cleanRoomsImpl struct {
	client *httpclient.ApiClient
}

func (a *cleanRoomsImpl) Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := "/api/2.0/clean-rooms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request.CleanRoom, &cleanRoom)
	return &cleanRoom, err
}

func (a *cleanRoomsImpl) CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error) {
	var createCleanRoomOutputCatalogResponse CreateCleanRoomOutputCatalogResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/output-catalogs", request.CleanRoomName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPost, path, headers, queryParams, request.OutputCatalog, &createCleanRoomOutputCatalogResponse)
	return &createCleanRoomOutputCatalogResponse, err
}

func (a *cleanRoomsImpl) Delete(ctx context.Context, request DeleteCleanRoomRequest) (*DeleteResponse, error) {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return &deleteResponse, err
}

func (a *cleanRoomsImpl) Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &cleanRoom)
	return &cleanRoom, err
}

// List clean rooms.
//
// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) List(ctx context.Context, request ListCleanRoomsRequest) listing.Iterator[CleanRoom] {

	getNextPage := func(ctx context.Context, req ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListCleanRoomsResponse) []CleanRoom {
		return resp.CleanRooms
	}
	getNextReq := func(resp *ListCleanRoomsResponse) *ListCleanRoomsRequest {
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

// List clean rooms.
//
// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) ListAll(ctx context.Context, request ListCleanRoomsRequest) ([]CleanRoom, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoom](ctx, iterator)
}
func (a *cleanRoomsImpl) internalList(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
	var listCleanRoomsResponse ListCleanRoomsResponse
	path := "/api/2.0/clean-rooms"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := do(a.client, ctx, http.MethodGet, path, headers, queryParams, request, &listCleanRoomsResponse)
	return &listCleanRoomsResponse, err
}

func (a *cleanRoomsImpl) Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := do(a.client, ctx, http.MethodPatch, path, headers, queryParams, request, &cleanRoom)
	return &cleanRoom, err
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
