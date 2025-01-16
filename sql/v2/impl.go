// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just alerts API methods
type alertsImpl struct {
	client *client.DatabricksClient
}

func (a *alertsImpl) Create(ctx context.Context, request CreateAlertRequest) (*Alert, error) {
	var alert Alert
	path := "/api/2.0/sql/alerts"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &alert)
	return &alert, err
}

func (a *alertsImpl) Delete(ctx context.Context, request TrashAlertRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &empty)
	return err
}

func (a *alertsImpl) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &alert)
	return &alert, err
}

func (a *alertsImpl) List(ctx context.Context, request ListAlertsRequest) (*ListAlertsResponse, error) {
	var listAlertsResponse ListAlertsResponse
	path := "/api/2.0/sql/alerts"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAlertsResponse)
	return &listAlertsResponse, err
}

func (a *alertsImpl) Update(ctx context.Context, request UpdateAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/sql/alerts/%v", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &alert)
	return &alert, err
}

// unexported type that holds implementations of just alerts_legacy API methods
type alertsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *alertsLegacyImpl) Create(ctx context.Context, request CreateAlert) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := "/api/2.0/preview/sql/alerts"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyImpl) Delete(ctx context.Context, request DeleteAlertsLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *alertsLegacyImpl) Get(ctx context.Context, request GetAlertsLegacyRequest) (*LegacyAlert, error) {
	var legacyAlert LegacyAlert
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &legacyAlert)
	return &legacyAlert, err
}

func (a *alertsLegacyImpl) List(ctx context.Context) ([]LegacyAlert, error) {
	var legacyAlertList []LegacyAlert
	path := "/api/2.0/preview/sql/alerts"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &legacyAlertList)
	return legacyAlertList, err
}

func (a *alertsLegacyImpl) Update(ctx context.Context, request EditAlert) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just dashboard_widgets API methods
type dashboardWidgetsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardWidgetsImpl) Create(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := "/api/2.0/preview/sql/widgets"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &widget)
	return &widget, err
}

func (a *dashboardWidgetsImpl) Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *dashboardWidgetsImpl) Update(ctx context.Context, request CreateWidget) (*Widget, error) {
	var widget Widget
	path := fmt.Sprintf("/api/2.0/preview/sql/widgets/%v", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &widget)
	return &widget, err
}

// unexported type that holds implementations of just dashboards API methods
type dashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *dashboardsImpl) Create(ctx context.Context, request DashboardPostContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/preview/sql/dashboards"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsImpl) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *dashboardsImpl) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsImpl) List(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error) {
	var listResponse ListResponse
	path := "/api/2.0/preview/sql/dashboards"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listResponse)
	return &listResponse, err
}

func (a *dashboardsImpl) Restore(ctx context.Context, request RestoreDashboardRequest) error {
	var restoreResponse RestoreResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &restoreResponse)
	return err
}

func (a *dashboardsImpl) Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &dashboard)
	return &dashboard, err
}

// unexported type that holds implementations of just data_sources API methods
type dataSourcesImpl struct {
	client *client.DatabricksClient
}

func (a *dataSourcesImpl) List(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0/preview/sql/data_sources"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &dataSourceList)
	return dataSourceList, err
}

// unexported type that holds implementations of just dbsql_permissions API methods
type dbsqlPermissionsImpl struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsImpl) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {
	var getResponse GetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getResponse)
	return &getResponse, err
}

func (a *dbsqlPermissionsImpl) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	var setResponse SetResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &setResponse)
	return &setResponse, err
}

func (a *dbsqlPermissionsImpl) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &success)
	return &success, err
}

// unexported type that holds implementations of just queries API methods
type queriesImpl struct {
	client *client.DatabricksClient
}

func (a *queriesImpl) Create(ctx context.Context, request CreateQueryRequest) (*Query, error) {
	var query Query
	path := "/api/2.0/sql/queries"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &query)
	return &query, err
}

func (a *queriesImpl) Delete(ctx context.Context, request TrashQueryRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &empty)
	return err
}

func (a *queriesImpl) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &query)
	return &query, err
}

func (a *queriesImpl) List(ctx context.Context, request ListQueriesRequest) (*ListQueryObjectsResponse, error) {
	var listQueryObjectsResponse ListQueryObjectsResponse
	path := "/api/2.0/sql/queries"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listQueryObjectsResponse)
	return &listQueryObjectsResponse, err
}

func (a *queriesImpl) ListVisualizations(ctx context.Context, request ListVisualizationsForQueryRequest) (*ListVisualizationsForQueryResponse, error) {
	var listVisualizationsForQueryResponse ListVisualizationsForQueryResponse
	path := fmt.Sprintf("/api/2.0/sql/queries/%v/visualizations", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listVisualizationsForQueryResponse)
	return &listVisualizationsForQueryResponse, err
}

func (a *queriesImpl) Update(ctx context.Context, request UpdateQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/sql/queries/%v", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &query)
	return &query, err
}

// unexported type that holds implementations of just queries_legacy API methods
type queriesLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queriesLegacyImpl) Create(ctx context.Context, request QueryPostContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := "/api/2.0/preview/sql/queries"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &legacyQuery)
	return &legacyQuery, err
}

func (a *queriesLegacyImpl) Delete(ctx context.Context, request DeleteQueriesLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *queriesLegacyImpl) Get(ctx context.Context, request GetQueriesLegacyRequest) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &legacyQuery)
	return &legacyQuery, err
}

func (a *queriesLegacyImpl) List(ctx context.Context, request ListQueriesLegacyRequest) (*QueryList, error) {
	var queryList QueryList
	path := "/api/2.0/preview/sql/queries"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &queryList)
	return &queryList, err
}

func (a *queriesLegacyImpl) Restore(ctx context.Context, request RestoreQueriesLegacyRequest) error {
	var restoreResponse RestoreResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &restoreResponse)
	return err
}

func (a *queriesLegacyImpl) Update(ctx context.Context, request QueryEditContent) (*LegacyQuery, error) {
	var legacyQuery LegacyQuery
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &legacyQuery)
	return &legacyQuery, err
}

// unexported type that holds implementations of just query_history API methods
type queryHistoryImpl struct {
	client *client.DatabricksClient
}

func (a *queryHistoryImpl) List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/sql/history/queries"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

// unexported type that holds implementations of just query_visualizations API methods
type queryVisualizationsImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsImpl) Create(ctx context.Context, request CreateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := "/api/2.0/sql/visualizations"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &visualization)
	return &visualization, err
}

func (a *queryVisualizationsImpl) Delete(ctx context.Context, request DeleteVisualizationRequest) error {
	var empty Empty
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &empty)
	return err
}

func (a *queryVisualizationsImpl) Update(ctx context.Context, request UpdateVisualizationRequest) (*Visualization, error) {
	var visualization Visualization
	path := fmt.Sprintf("/api/2.0/sql/visualizations/%v", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &visualization)
	return &visualization, err
}

// unexported type that holds implementations of just query_visualizations_legacy API methods
type queryVisualizationsLegacyImpl struct {
	client *client.DatabricksClient
}

func (a *queryVisualizationsLegacyImpl) Create(ctx context.Context, request CreateQueryVisualizationsLegacyRequest) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := "/api/2.0/preview/sql/visualizations"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &legacyVisualization)
	return &legacyVisualization, err
}

func (a *queryVisualizationsLegacyImpl) Delete(ctx context.Context, request DeleteQueryVisualizationsLegacyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *queryVisualizationsLegacyImpl) Update(ctx context.Context, request LegacyVisualization) (*LegacyVisualization, error) {
	var legacyVisualization LegacyVisualization
	path := fmt.Sprintf("/api/2.0/preview/sql/visualizations/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &legacyVisualization)
	return &legacyVisualization, err
}

// unexported type that holds implementations of just statement_execution API methods
type statementExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *statementExecutionImpl) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {
	var cancelExecutionResponse CancelExecutionResponse
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/cancel", request.StatementId)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &cancelExecutionResponse)
	return err
}

func (a *statementExecutionImpl) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := "/api/2.0/sql/statements/"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionImpl) GetStatement(ctx context.Context, request GetStatementRequest) (*StatementResponse, error) {
	var statementResponse StatementResponse
	path := fmt.Sprintf("/api/2.0/sql/statements/%v", request.StatementId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &statementResponse)
	return &statementResponse, err
}

func (a *statementExecutionImpl) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	var resultData ResultData
	path := fmt.Sprintf("/api/2.0/sql/statements/%v/result/chunks/%v", request.StatementId, request.ChunkIndex)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &resultData)
	return &resultData, err
}

// unexported type that holds implementations of just warehouses API methods
type warehousesImpl struct {
	client *client.DatabricksClient
}

func (a *warehousesImpl) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	var createWarehouseResponse CreateWarehouseResponse
	path := "/api/2.0/sql/warehouses"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createWarehouseResponse)
	return &createWarehouseResponse, err
}

func (a *warehousesImpl) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	var deleteWarehouseResponse DeleteWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteWarehouseResponse)
	return err
}

func (a *warehousesImpl) Edit(ctx context.Context, request EditWarehouseRequest) error {
	var editWarehouseResponse EditWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/edit", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &editWarehouseResponse)
	return err
}

func (a *warehousesImpl) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	var getWarehouseResponse GetWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getWarehouseResponse)
	return &getWarehouseResponse, err
}

func (a *warehousesImpl) GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error) {
	var getWarehousePermissionLevelsResponse GetWarehousePermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v/permissionLevels", request.WarehouseId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getWarehousePermissionLevelsResponse)
	return &getWarehousePermissionLevelsResponse, err
}

func (a *warehousesImpl) GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesImpl) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	var getWorkspaceWarehouseConfigResponse GetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &getWorkspaceWarehouseConfigResponse)
	return &getWorkspaceWarehouseConfigResponse, err
}

func (a *warehousesImpl) List(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error) {
	var listWarehousesResponse ListWarehousesResponse
	path := "/api/2.0/sql/warehouses"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listWarehousesResponse)
	return &listWarehousesResponse, err
}

func (a *warehousesImpl) SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &warehousePermissions)
	return &warehousePermissions, err
}

func (a *warehousesImpl) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	var setWorkspaceWarehouseConfigResponse SetWorkspaceWarehouseConfigResponse
	path := "/api/2.0/sql/config/warehouses"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &setWorkspaceWarehouseConfigResponse)
	return err
}

func (a *warehousesImpl) Start(ctx context.Context, request StartRequest) error {
	var startWarehouseResponse StartWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/start", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &startWarehouseResponse)
	return err
}

func (a *warehousesImpl) Stop(ctx context.Context, request StopRequest) error {
	var stopWarehouseResponse StopWarehouseResponse
	path := fmt.Sprintf("/api/2.0/sql/warehouses/%v/stop", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &stopWarehouseResponse)
	return err
}

func (a *warehousesImpl) UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error) {
	var warehousePermissions WarehousePermissions
	path := fmt.Sprintf("/api/2.0/permissions/warehouses/%v", request.WarehouseId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &warehousePermissions)
	return &warehousePermissions, err
}
