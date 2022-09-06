// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
	"time"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type WarehousesService interface {
	// Creates a new SQL warehouse.
	CreateWarehouse(ctx context.Context, createWarehouseRequest CreateWarehouseRequest) (*CreateWarehouseResponse, error)
	// CreateWarehouse and wait to reach RUNNING state
	CreateWarehouseAndWait(ctx context.Context, request CreateWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error)
	// Deletes a SQL warehouse.
	DeleteWarehouse(ctx context.Context, deleteWarehouseRequest DeleteWarehouseRequest) error
	// DeleteWarehouse and wait to reach DELETED state
	DeleteWarehouseAndWait(ctx context.Context, request DeleteWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error)
	DeleteWarehouseById(ctx context.Context, id string) error
	DeleteWarehouseByIdAndWait(ctx context.Context, id string, timeout ...time.Duration) (*GetWarehouseResponse, error)
	// Edits a SQL warehouse.
	EditWarehouse(ctx context.Context, editWarehouseRequest EditWarehouseRequest) error
	// EditWarehouse and wait to reach RUNNING state
	EditWarehouseAndWait(ctx context.Context, request EditWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error)
	// Gets the information for a single SQL warehouse.
	GetWarehouse(ctx context.Context, getWarehouseRequest GetWarehouseRequest) (*GetWarehouseResponse, error)
	// GetWarehouse and wait to reach RUNNING state
	GetWarehouseAndWait(ctx context.Context, request GetWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error)
	GetWarehouseById(ctx context.Context, id string) (*GetWarehouseResponse, error)
	GetWarehouseByIdAndWait(ctx context.Context, id string, timeout ...time.Duration) (*GetWarehouseResponse, error)
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
	// StartWarehouse and wait to reach RUNNING state
	StartWarehouseAndWait(ctx context.Context, request StartWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error)
	// Stops a SQL warehouse.
	StopWarehouse(ctx context.Context, stopWarehouseRequest StopWarehouseRequest) error
	// StopWarehouse and wait to reach STOPPED state
	StopWarehouseAndWait(ctx context.Context, request StopWarehouseRequest, timeout ...time.Duration) (*GetWarehouseResponse, error)
}
