// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Alerts API methods
type alertsImpl struct {
	client *client.DatabricksClient
}

func (a *alertsImpl) Create(ctx context.Context, request CreateAlert) (*Alert, error) {
	var alert Alert
	path := "/api/2.0/preview/sql/alerts"
	err := a.client.Do(ctx, http.MethodPost, path, request, &alert)
	return &alert, err
}

func (a *alertsImpl) Delete(ctx context.Context, request DeleteAlertRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *alertsImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &alert)
	return &alert, err
}

func (a *alertsImpl) List(ctx context.Context) ([]Alert, error) {
	var alertList []Alert
	path := "/api/2.0/preview/sql/alerts"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &alertList)
	return alertList, err
}

func (a *alertsImpl) Update(ctx context.Context, request EditAlert) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just Dashboards API methods
type dashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsImpl) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/preview/sql/dashboards"
	err := a.client.Do(ctx, http.MethodPost, path, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *dashboardsImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsImpl) List(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/preview/sql/dashboards"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listResponse)
	return &listResponse, err
}

func (a *dashboardsImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just DataSources API methods
type dataSourcesImpl struct {
	client *client.DatabricksClient
}

func (a *dataSourcesImpl) List(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0/preview/sql/data_sources"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &dataSourceList)
	return dataSourceList, err
}

// unexported type that holds implementations of just DbsqlPermissions API methods
type dbsqlPermissionsImpl struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsImpl) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getResponse)
	return &getResponse, err
}

func (a *dbsqlPermissionsImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	var setResponse SetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &setResponse)
	return &setResponse, err
}

func (a *dbsqlPermissionsImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &success)
	return &success, err
}

// unexported type that holds implementations of just Queries API methods
type queriesImpl struct {
	client *client.DatabricksClient
}

func (a *queriesImpl) Create(ctx context.Context, request QueryPostContent) (*Query, error) {
	var query Query
	path := "/api/2.0/preview/sql/queries"
	err := a.client.Do(ctx, http.MethodPost, path, request, &query)
	return &query, err
}

func (a *queriesImpl) Delete(ctx context.Context, request DeleteQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *queriesImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &query)
	return &query, err
}

func (a *queriesImpl) List(ctx context.Context, request ListQueriesRequest) (*QueryList, error) {
	var queryList QueryList
	path := "/api/2.0/preview/sql/queries"
	err := a.client.Do(ctx, http.MethodGet, path, request, &queryList)
	return &queryList, err
}

func (a *queriesImpl) Restore(ctx context.Context, request RestoreQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *queriesImpl) Update(ctx context.Context, request QueryEditContent) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &query)
	return &query, err
}

// unexported type that holds implementations of just QueryHistory API methods
type queryHistoryImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryImpl) List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/sql/history/queries"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

// unexported type that holds implementations of just StatementExecution API methods
type statementExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionImpl) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/cancel", request.StatementId)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *statementExecutionImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*ExecuteStatementResponse, error) {
	var executeStatementResponse ExecuteStatementResponse
	path := "/api/2.0/sql/statements/"
	err := a.client.Do(ctx, http.MethodPost, path, request, &executeStatementResponse)
	return &executeStatementResponse, err
}

func (a *statementExecutionImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*GetStatementResponse, error) {
	var getStatementResponse GetStatementResponse
	path := fmt.Sprintf("/api/2.0/sql/statements/%v", request.StatementId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getStatementResponse)
	return &getStatementResponse, err
}

func (a *statementExecutionImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	var resultData ResultData
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/result/chunks/%v", request.StatementId, request.ChunkIndex)
	err := a.client.Do(ctx, http.MethodGet, path, request, &resultData)
	return &resultData, err
}

// unexported type that holds implementations of just Warehouses API methods
type warehousesImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesImpl) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

func (a *warehousesImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *warehousesImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *warehousesImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
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

func (a *warehousesImpl) List(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
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

func (a *warehousesImpl) Start(ctx context.Context, request StartRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *warehousesImpl) Stop(ctx context.Context, request StopRequest) error {
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}
