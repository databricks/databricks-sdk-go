// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just CleanRoomAssets API methods
type cleanRoomAssetsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomAssetsImpl) Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.Asset, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

func (a *cleanRoomAssetsImpl) Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) error {
	var deleteCleanRoomAssetResponse DeleteCleanRoomAssetResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.AssetFullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteCleanRoomAssetResponse)
	return err
}

func (a *cleanRoomAssetsImpl) Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.AssetFullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

func (a *cleanRoomAssetsImpl) List(ctx context.Context, request ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error) {
	var listCleanRoomAssetsResponse ListCleanRoomAssetsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets", request.CleanRoomName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listCleanRoomAssetsResponse)
	return &listCleanRoomAssetsResponse, err
}

func (a *cleanRoomAssetsImpl) Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error) {
	var cleanRoomAsset CleanRoomAsset
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/assets/%v/%v", request.CleanRoomName, request.AssetType, request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request.Asset, &cleanRoomAsset)
	return &cleanRoomAsset, err
}

// unexported type that holds implementations of just CleanRoomTaskRuns API methods
type cleanRoomTaskRunsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomTaskRunsImpl) List(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	var listCleanRoomNotebookTaskRunsResponse ListCleanRoomNotebookTaskRunsResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/runs", request.CleanRoomName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listCleanRoomNotebookTaskRunsResponse)
	return &listCleanRoomNotebookTaskRunsResponse, err
}

// unexported type that holds implementations of just CleanRooms API methods
type cleanRoomsImpl struct {
	client *client.DatabricksClient
}

func (a *cleanRoomsImpl) Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := "/api/2.0/clean-rooms"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.CleanRoom, &cleanRoom)
	return &cleanRoom, err
}

func (a *cleanRoomsImpl) CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error) {
	var createCleanRoomOutputCatalogResponse CreateCleanRoomOutputCatalogResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v/output-catalogs", request.CleanRoomName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.OutputCatalog, &createCleanRoomOutputCatalogResponse)
	return &createCleanRoomOutputCatalogResponse, err
}

func (a *cleanRoomsImpl) Delete(ctx context.Context, request DeleteCleanRoomRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *cleanRoomsImpl) Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &cleanRoom)
	return &cleanRoom, err
}

func (a *cleanRoomsImpl) List(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error) {
	var listCleanRoomsResponse ListCleanRoomsResponse
	path := "/api/2.0/clean-rooms"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listCleanRoomsResponse)
	return &listCleanRoomsResponse, err
}

func (a *cleanRoomsImpl) Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error) {
	var cleanRoom CleanRoom
	path := fmt.Sprintf("/api/2.0/clean-rooms/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &cleanRoom)
	return &cleanRoom, err
}
