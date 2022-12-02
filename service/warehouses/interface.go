// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
)

// Access the history of queries through SQL warehouses.
type QueryHistoryService interface {

	// List
	//
	// List the history of queries through SQL warehouses.
	//
	// You can filter by user ID, warehouse ID, status, and time range.
	//
	// Use ListQueriesAll() to get all QueryInfo instances, which will iterate over every result page.
	ListQueries(ctx context.Context, request ListQueriesRequest) (*ListQueriesResponse, error)
}

// A SQL warehouse is a compute resource that lets you run SQL commands on data
// objects within Databricks SQL. Compute resources are infrastructure resources
// that provide processing capabilities in the cloud.
type WarehousesService interface {

	// Create a warehouse
	//
	// Creates a new SQL warehouse.
	CreateWarehouse(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error)

	// Delete a warehouse
	//
	// Deletes a SQL warehouse.
	DeleteWarehouse(ctx context.Context, request DeleteWarehouse) error

	// Update a warehouse
	//
	// Updates the configuration for a SQL warehouse.
	EditWarehouse(ctx context.Context, request EditWarehouseRequest) error

	// Get warehouse info
	//
	// Gets the information for a single SQL warehouse.
	GetWarehouse(ctx context.Context, request GetWarehouse) (*GetWarehouseResponse, error)

	// Get the workspace configuration
	//
	// Gets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error)

	// List warehouses
	//
	// Lists all SQL warehouses that a user has manager permissions on.
	//
	// Use ListWarehousesAll() to get all EndpointInfo instances
	ListWarehouses(ctx context.Context, request ListWarehouses) (*ListWarehousesResponse, error)

	// Set the workspace configuration
	//
	// Sets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error

	// Start a warehouse
	//
	// Starts a SQL warehouse.
	StartWarehouse(ctx context.Context, request StartWarehouse) error

	// Stop a warehouse
	//
	// Stops a SQL warehouse.
	StopWarehouse(ctx context.Context, request StopWarehouse) error
}
