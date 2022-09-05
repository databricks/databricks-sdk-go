// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type WarehousesService interface {
	// Creates a new SQL warehouse.
	CreateWarehouse(ctx context.Context, createWarehouseRequest CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	// Deletes a SQL warehouse.
	DeleteWarehouse(ctx context.Context, deleteWarehouseRequest DeleteWarehouseRequest) error
	// Edits a SQL warehouse.
	EditWarehouse(ctx context.Context, editWarehouseRequest EditWarehouseRequest) error
	// Gets the information for a single SQL warehouse.
	GetWarehouse(ctx context.Context, getWarehouseRequest GetWarehouseRequest) (*GetWarehouseResponse, error)
	// Gets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error)
	// Lists all SQL warehouse a user has manager permissions for.
	ListWarehouses(ctx context.Context, listWarehousesRequest ListWarehousesRequest) (*ListWarehousesResponse, error)
	// Sets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	SetWorkspaceWarehouseConfig(ctx context.Context, setWorkspaceWarehouseConfigRequest SetWorkspaceWarehouseConfigRequest) error
	// Starts a SQL warehouse.
	StartWarehouse(ctx context.Context, startWarehouseRequest StartWarehouseRequest) error
	// Stops a SQL warehouse.
	StopWarehouse(ctx context.Context, stopWarehouseRequest StopWarehouseRequest) error
	GetWarehouseById(ctx context.Context, id string) (*GetWarehouseResponse, error)
	DeleteWarehouseById(ctx context.Context, id string) error
}

func New(client *client.DatabricksClient) WarehousesService {
	return &WarehousesAPI{
		client: client,
	}
}

type WarehousesAPI struct {
	client *client.DatabricksClient
}

// Creates a new SQL warehouse.
func (a *WarehousesAPI) CreateWarehouse(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Post(ctx, path, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

// Deletes a SQL warehouse.
func (a *WarehousesAPI) DeleteWarehouse(ctx context.Context, request DeleteWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Edits a SQL warehouse.
func (a *WarehousesAPI) EditWarehouse(ctx context.Context, request EditWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) GetWarehouse(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Get(ctx, path, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

// Gets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"
	err := a.client.Get(ctx, path, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

// Lists all SQL warehouse a user has manager permissions for.
func (a *WarehousesAPI) ListWarehouses(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Get(ctx, path, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

// Sets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	path := "/api/2.0/sql/config/warehouses"
	err := a.client.Put(ctx, path, request)
	return err
}

// Starts a SQL warehouse.
func (a *WarehousesAPI) StartWarehouse(ctx context.Context, request StartWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Stops a SQL warehouse.
func (a *WarehousesAPI) StopWarehouse(ctx context.Context, request StopWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *WarehousesAPI) GetWarehouseById(ctx context.Context, id string) (*GetWarehouseResponse, error) {
	return a.GetWarehouse(ctx, GetWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) DeleteWarehouseById(ctx context.Context, id string) error {
	return a.DeleteWarehouse(ctx, DeleteWarehouseRequest{
		Id: id,
	})
}
