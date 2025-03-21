// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Clean Room Assets, Clean Room Task Runs, Clean Rooms, etc.
package cleanrooms

import (
	"context"
)

type cleanRoomAssetsBaseClient struct {
	cleanRoomAssetsImpl
}

// Delete an asset.
//
// Delete a clean room asset - unshare/remove the asset from the clean room
func (a *cleanRoomAssetsBaseClient) DeleteByCleanRoomNameAndAssetTypeAndAssetFullName(ctx context.Context, cleanRoomName string, assetType CleanRoomAssetAssetType, assetFullName string) (*DeleteCleanRoomAssetResponse, error) {
	return a.cleanRoomAssetsImpl.Delete(ctx, DeleteCleanRoomAssetRequest{
		CleanRoomName: cleanRoomName,
		AssetType:     assetType,
		AssetFullName: assetFullName,
	})
}

// Get an asset.
//
// Get the details of a clean room asset by its type and full name.
func (a *cleanRoomAssetsBaseClient) GetByCleanRoomNameAndAssetTypeAndAssetFullName(ctx context.Context, cleanRoomName string, assetType CleanRoomAssetAssetType, assetFullName string) (*CleanRoomAsset, error) {
	return a.cleanRoomAssetsImpl.Get(ctx, GetCleanRoomAssetRequest{
		CleanRoomName: cleanRoomName,
		AssetType:     assetType,
		AssetFullName: assetFullName,
	})
}

// List assets.
func (a *cleanRoomAssetsBaseClient) ListByCleanRoomName(ctx context.Context, cleanRoomName string) (*ListCleanRoomAssetsResponse, error) {
	return a.cleanRoomAssetsImpl.internalList(ctx, ListCleanRoomAssetsRequest{
		CleanRoomName: cleanRoomName,
	})
}

type cleanRoomTaskRunsBaseClient struct {
	cleanRoomTaskRunsImpl
}

// List notebook task runs.
//
// List all the historical notebook task runs in a clean room.
func (a *cleanRoomTaskRunsBaseClient) ListByCleanRoomName(ctx context.Context, cleanRoomName string) (*ListCleanRoomNotebookTaskRunsResponse, error) {
	return a.cleanRoomTaskRunsImpl.internalList(ctx, ListCleanRoomNotebookTaskRunsRequest{
		CleanRoomName: cleanRoomName,
	})
}

type cleanRoomsBaseClient struct {
	cleanRoomsImpl
}

// Delete a clean room.
//
// Delete a clean room. After deletion, the clean room will be removed from the
// metastore. If the other collaborators have not deleted the clean room, they
// will still have the clean room in their metastore, but it will be in a
// DELETED state and no operations other than deletion can be performed on it.
func (a *cleanRoomsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.cleanRoomsImpl.Delete(ctx, DeleteCleanRoomRequest{
		Name: name,
	})
}

// Get a clean room.
//
// Get the details of a clean room given its name.
func (a *cleanRoomsBaseClient) GetByName(ctx context.Context, name string) (*CleanRoom, error) {
	return a.cleanRoomsImpl.Get(ctx, GetCleanRoomRequest{
		Name: name,
	})
}
