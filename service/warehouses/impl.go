// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just QueryHistory API methods
type queryHistoryImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryImpl) ListQueries(ctx context.Context, request ListQueriesRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/sql/history/queries"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

// unexported type that holds implementations of just Warehouses API methods
type warehousesImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesImpl) CreateWarehouse(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

func (a *warehousesImpl) DeleteWarehouse(ctx context.Context, request DeleteWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *warehousesImpl) EditWarehouse(ctx context.Context, request EditWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *warehousesImpl) GetWarehouse(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

func (a *warehousesImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

func (a *warehousesImpl) ListWarehouses(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

func (a *warehousesImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	path := "/api/2.0/sql/config/warehouses"
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *warehousesImpl) StartWarehouse(ctx context.Context, request StartWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *warehousesImpl) StopWarehouse(ctx context.Context, request StopWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
