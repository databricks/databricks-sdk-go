// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just CleanRoomAssets API methods
type cleanRoomAssetsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAssetsImpl) Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error) {

	requestPb, pbErr := createCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", requestPb.CleanRoomName)
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
		(*requestPb).Asset,
		&cleanRoomAssetPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAssetsImpl) Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) error {

	requestPb, pbErr := deleteCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", requestPb.CleanRoomName, requestPb.AssetType, requestPb.Name)
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

func (a *cleanRoomAssetsImpl) Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error) {

	requestPb, pbErr := getCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", requestPb.CleanRoomName, requestPb.AssetType, requestPb.Name)
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
		&cleanRoomAssetPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listCleanRoomAssetsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomAssetsResponsePb listCleanRoomAssetsResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", requestPb.CleanRoomName)
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
		&listCleanRoomAssetsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listCleanRoomAssetsResponseFromPb(&listCleanRoomAssetsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomAssetsImpl) Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error) {

	requestPb, pbErr := updateCleanRoomAssetRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomAssetPb cleanRoomAssetPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", requestPb.CleanRoomName, requestPb.AssetType, requestPb.Name)
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
		(*requestPb).Asset,
		&cleanRoomAssetPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cleanRoomAssetFromPb(&cleanRoomAssetPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CleanRoomTaskRuns API methods
type cleanRoomTaskRunsImpl struct {
	client *client.DatabricksClient
}

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

// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsImpl) ListAll(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) ([]CleanRoomNotebookTaskRun, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoomNotebookTaskRun](ctx, iterator)
}

func (a *cleanRoomTaskRunsImpl) internalList(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {

	requestPb, pbErr := listCleanRoomNotebookTaskRunsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomNotebookTaskRunsResponsePb listCleanRoomNotebookTaskRunsResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/runs", requestPb.CleanRoomName)
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
		&listCleanRoomNotebookTaskRunsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listCleanRoomNotebookTaskRunsResponseFromPb(&listCleanRoomNotebookTaskRunsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just CleanRooms API methods
type cleanRoomsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomsImpl) Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error) {

	requestPb, pbErr := createCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomPb cleanRoomPb
	path := "/api/2.0/clean-rooms"
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
		(*requestPb).CleanRoom,
		&cleanRoomPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cleanRoomFromPb(&cleanRoomPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomsImpl) CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error) {

	requestPb, pbErr := createCleanRoomOutputCatalogRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createCleanRoomOutputCatalogResponsePb createCleanRoomOutputCatalogResponsePb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/output-catalogs", requestPb.CleanRoomName)
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
		(*requestPb).OutputCatalog,
		&createCleanRoomOutputCatalogResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createCleanRoomOutputCatalogResponseFromPb(&createCleanRoomOutputCatalogResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomsImpl) Delete(ctx context.Context, request DeleteCleanRoomRequest) error {

	requestPb, pbErr := deleteCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", requestPb.Name)
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

func (a *cleanRoomsImpl) Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error) {

	requestPb, pbErr := getCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomPb cleanRoomPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", requestPb.Name)
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
		&cleanRoomPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cleanRoomFromPb(&cleanRoomPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

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

// Get a list of all clean rooms of the metastore. Only clean rooms the caller
// has access to are returned.
func (a *cleanRoomsImpl) ListAll(ctx context.Context, request ListCleanRoomsRequest) ([]CleanRoom, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[CleanRoom](ctx, iterator)
}

func (a *cleanRoomsImpl) internalList(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {

	requestPb, pbErr := listCleanRoomsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listCleanRoomsResponsePb listCleanRoomsResponsePb
	path := "/api/2.0/clean-rooms"
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
		&listCleanRoomsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listCleanRoomsResponseFromPb(&listCleanRoomsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *cleanRoomsImpl) Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error) {

	requestPb, pbErr := updateCleanRoomRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var cleanRoomPb cleanRoomPb
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", requestPb.Name)
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
		&cleanRoomPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := cleanRoomFromPb(&cleanRoomPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
