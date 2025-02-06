// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
)

// Clean room assets are data and code objects — Tables, volumes, and
// notebooks that are shared with the clean room.
type CleanRoomAssetsService interface {

	// Create an asset.
	//
	// Create a clean room asset —share an asset like a notebook or table into
	// the clean room. For each UC asset that is added through this method, the
	// clean room owner must also have enough privilege on the asset to consume
	// it. The privilege must be maintained indefinitely for the clean room to
	// be able to access the asset. Typically, you should use a group as the
	// clean room owner.
	Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error)

	// Delete an asset.
	//
	// Delete a clean room asset - unshare/remove the asset from the clean room
	Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) error

	// Get an asset.
	//
	// Get the details of a clean room asset by its type and full name.
	Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error)

	// List assets.
	//
	// Use ListAll() to get all CleanRoomAsset instances, which will iterate over every result page.
	List(ctx context.Context, request ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error)

	// Update an asset.
	//
	// Update a clean room asset. For example, updating the content of a
	// notebook; changing the shared partitions of a table; etc.
	Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error)
}

// Clean room task runs are the executions of notebooks in a clean room.
type CleanRoomTaskRunsService interface {

	// List notebook task runs.
	//
	// List all the historical notebook task runs in a clean room.
	//
	// Use ListAll() to get all CleanRoomNotebookTaskRun instances, which will iterate over every result page.
	List(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error)
}

// A clean room uses Delta Sharing and serverless compute to provide a secure
// and privacy-protecting environment where multiple parties can work together
// on sensitive enterprise data without direct access to each other’s data.
type CleanRoomsService interface {

	// Create a clean room.
	//
	// Create a new clean room with the specified collaborators. This method is
	// asynchronous; the returned name field inside the clean_room field can be
	// used to poll the clean room status, using the :method:cleanrooms/get
	// method. When this method returns, the clean room will be in a
	// PROVISIONING state, with only name, owner, comment, created_at and status
	// populated. The clean room will be usable once it enters an ACTIVE state.
	//
	// The caller must be a metastore admin or have the **CREATE_CLEAN_ROOM**
	// privilege on the metastore.
	Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error)

	// Create an output catalog.
	//
	// Create the output catalog of the clean room.
	CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error)

	// Delete a clean room.
	//
	// Delete a clean room. After deletion, the clean room will be removed from
	// the metastore. If the other collaborators have not deleted the clean
	// room, they will still have the clean room in their metastore, but it will
	// be in a DELETED state and no operations other than deletion can be
	// performed on it.
	Delete(ctx context.Context, request DeleteCleanRoomRequest) error

	// Get a clean room.
	//
	// Get the details of a clean room given its name.
	Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error)

	// List clean rooms.
	//
	// Get a list of all clean rooms of the metastore. Only clean rooms the
	// caller has access to are returned.
	//
	// Use ListAll() to get all CleanRoom instances, which will iterate over every result page.
	List(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error)

	// Update a clean room.
	//
	// Update a clean room. The caller must be the owner of the clean room, have
	// **MODIFY_CLEAN_ROOM** privilege, or be metastore admin.
	//
	// When the caller is a metastore admin, only the __owner__ field can be
	// updated.
	Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error)
}
