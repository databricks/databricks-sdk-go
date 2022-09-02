// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
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
