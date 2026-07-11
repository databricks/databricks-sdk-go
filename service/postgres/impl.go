// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
	"golang.org/x/exp/slices"
)

// unexported type that holds implementations of just Postgres API methods
type postgresImpl struct {
	client *client.DatabricksClient
}

func (a *postgresImpl) CreateBranch(ctx context.Context, request CreateBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/branches", request.Parent)
	queryParams := make(map[string]any)

	if request.BranchId != "" || slices.Contains(request.ForceSendFields, "BranchId") {
		queryParams["branch_id"] = request.BranchId
	}

	if request.ReplaceExisting != false || slices.Contains(request.ForceSendFields, "ReplaceExisting") {
		queryParams["replace_existing"] = request.ReplaceExisting
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Branch, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/postgres/catalogs"
	queryParams := make(map[string]any)

	if request.CatalogId != "" {
		queryParams["catalog_id"] = request.CatalogId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Catalog, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateCdfConfig(ctx context.Context, request CreateCdfConfigRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/cdf-configs", request.Parent)
	queryParams := make(map[string]any)

	if request.CdfConfigId != "" || slices.Contains(request.ForceSendFields, "CdfConfigId") {
		queryParams["cdf_config_id"] = request.CdfConfigId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.CdfConfig, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateDataApi(ctx context.Context, request CreateDataApiRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/data-api", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DataApi, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateDatabase(ctx context.Context, request CreateDatabaseRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/databases", request.Parent)
	queryParams := make(map[string]any)

	if request.DatabaseId != "" || slices.Contains(request.ForceSendFields, "DatabaseId") {
		queryParams["database_id"] = request.DatabaseId
	}

	if request.ReplaceExisting != false || slices.Contains(request.ForceSendFields, "ReplaceExisting") {
		queryParams["replace_existing"] = request.ReplaceExisting
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Database, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)

	if request.EndpointId != "" || slices.Contains(request.ForceSendFields, "EndpointId") {
		queryParams["endpoint_id"] = request.EndpointId
	}

	if request.ReplaceExisting != false || slices.Contains(request.ForceSendFields, "ReplaceExisting") {
		queryParams["replace_existing"] = request.ReplaceExisting
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Endpoint, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateProject(ctx context.Context, request CreateProjectRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/postgres/projects"
	queryParams := make(map[string]any)

	if request.ProjectId != "" {
		queryParams["project_id"] = request.ProjectId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Project, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateRole(ctx context.Context, request CreateRoleRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/roles", request.Parent)
	queryParams := make(map[string]any)

	if request.ReplaceExisting != false || slices.Contains(request.ForceSendFields, "ReplaceExisting") {
		queryParams["replace_existing"] = request.ReplaceExisting
	}

	if request.RoleId != "" || slices.Contains(request.ForceSendFields, "RoleId") {
		queryParams["role_id"] = request.RoleId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Role, &operation)
	return &operation, err
}

func (a *postgresImpl) CreateSyncedTable(ctx context.Context, request CreateSyncedTableRequest) (*Operation, error) {
	var operation Operation
	path := "/api/2.0/postgres/synced_tables"
	queryParams := make(map[string]any)

	if request.SyncedTableId != "" {
		queryParams["synced_table_id"] = request.SyncedTableId
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.SyncedTable, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteBranch(ctx context.Context, request DeleteBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteCdfConfig(ctx context.Context, request DeleteCdfConfigRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteDataApi(ctx context.Context, request DeleteDataApiRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteDatabase(ctx context.Context, request DeleteDatabaseRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteProject(ctx context.Context, request DeleteProjectRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteRole(ctx context.Context, request DeleteRoleRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) DeleteSyncedTable(ctx context.Context, request DeleteSyncedTableRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error) {
	var databaseCredential DatabaseCredential
	path := "/api/2.0/postgres/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &databaseCredential)
	return &databaseCredential, err
}

func (a *postgresImpl) GetBranch(ctx context.Context, request GetBranchRequest) (*Branch, error) {
	var branch Branch
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &branch)
	return &branch, err
}

func (a *postgresImpl) GetCatalog(ctx context.Context, request GetCatalogRequest) (*Catalog, error) {
	var catalog Catalog
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &catalog)
	return &catalog, err
}

func (a *postgresImpl) GetCdfConfig(ctx context.Context, request GetCdfConfigRequest) (*CdfConfig, error) {
	var cdfConfig CdfConfig
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &cdfConfig)
	return &cdfConfig, err
}

func (a *postgresImpl) GetCdfStatus(ctx context.Context, request GetCdfStatusRequest) (*CdfStatus, error) {
	var cdfStatus CdfStatus
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &cdfStatus)
	return &cdfStatus, err
}

func (a *postgresImpl) GetDataApi(ctx context.Context, request GetDataApiRequest) (*DataApi, error) {
	var dataApi DataApi
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &dataApi)
	return &dataApi, err
}

func (a *postgresImpl) GetDatabase(ctx context.Context, request GetDatabaseRequest) (*Database, error) {
	var database Database
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &database)
	return &database, err
}

func (a *postgresImpl) GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error) {
	var endpoint Endpoint
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &endpoint)
	return &endpoint, err
}

func (a *postgresImpl) GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) GetProject(ctx context.Context, request GetProjectRequest) (*Project, error) {
	var project Project
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &project)
	return &project, err
}

func (a *postgresImpl) GetRole(ctx context.Context, request GetRoleRequest) (*Role, error) {
	var role Role
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &role)
	return &role, err
}

func (a *postgresImpl) GetSyncedTable(ctx context.Context, request GetSyncedTableRequest) (*SyncedTable, error) {
	var syncedTable SyncedTable
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &syncedTable)
	return &syncedTable, err
}

// Returns a paginated list of database branches in the project.
func (a *postgresImpl) ListBranches(ctx context.Context, request ListBranchesRequest) listing.Iterator[Branch] {

	getNextPage := func(ctx context.Context, req ListBranchesRequest) (*ListBranchesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListBranches(ctx, req)
	}
	getItems := func(resp *ListBranchesResponse) []Branch {
		return resp.Branches
	}
	getNextReq := func(resp *ListBranchesResponse) *ListBranchesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Returns a paginated list of database branches in the project.
func (a *postgresImpl) ListBranchesAll(ctx context.Context, request ListBranchesRequest) ([]Branch, error) {
	iterator := a.ListBranches(ctx, request)
	return listing.ToSlice[Branch](ctx, iterator)
}

func (a *postgresImpl) internalListBranches(ctx context.Context, request ListBranchesRequest) (*ListBranchesResponse, error) {
	var listBranchesResponse ListBranchesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/branches", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBranchesResponse)
	return &listBranchesResponse, err
}

// List the Lakebase CDF configurations (CdfConfigs) under a database.
func (a *postgresImpl) ListCdfConfigs(ctx context.Context, request ListCdfConfigsRequest) listing.Iterator[CdfConfig] {

	getNextPage := func(ctx context.Context, req ListCdfConfigsRequest) (*ListCdfConfigsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCdfConfigs(ctx, req)
	}
	getItems := func(resp *ListCdfConfigsResponse) []CdfConfig {
		return resp.CdfConfigs
	}
	getNextReq := func(resp *ListCdfConfigsResponse) *ListCdfConfigsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List the Lakebase CDF configurations (CdfConfigs) under a database.
func (a *postgresImpl) ListCdfConfigsAll(ctx context.Context, request ListCdfConfigsRequest) ([]CdfConfig, error) {
	iterator := a.ListCdfConfigs(ctx, request)
	return listing.ToSlice[CdfConfig](ctx, iterator)
}

func (a *postgresImpl) internalListCdfConfigs(ctx context.Context, request ListCdfConfigsRequest) (*ListCdfConfigsResponse, error) {
	var listCdfConfigsResponse ListCdfConfigsResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/cdf-configs", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCdfConfigsResponse)
	return &listCdfConfigsResponse, err
}

// List the replication statuses of all tables replicated under a Lakebase CDF
// configuration.
func (a *postgresImpl) ListCdfStatuses(ctx context.Context, request ListCdfStatusesRequest) listing.Iterator[CdfStatus] {

	getNextPage := func(ctx context.Context, req ListCdfStatusesRequest) (*ListCdfStatusesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCdfStatuses(ctx, req)
	}
	getItems := func(resp *ListCdfStatusesResponse) []CdfStatus {
		return resp.CdfStatuses
	}
	getNextReq := func(resp *ListCdfStatusesResponse) *ListCdfStatusesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List the replication statuses of all tables replicated under a Lakebase CDF
// configuration.
func (a *postgresImpl) ListCdfStatusesAll(ctx context.Context, request ListCdfStatusesRequest) ([]CdfStatus, error) {
	iterator := a.ListCdfStatuses(ctx, request)
	return listing.ToSlice[CdfStatus](ctx, iterator)
}

func (a *postgresImpl) internalListCdfStatuses(ctx context.Context, request ListCdfStatusesRequest) (*ListCdfStatusesResponse, error) {
	var listCdfStatusesResponse ListCdfStatusesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/cdf-statuses", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listCdfStatusesResponse)
	return &listCdfStatusesResponse, err
}

// List Databases.
func (a *postgresImpl) ListDatabases(ctx context.Context, request ListDatabasesRequest) listing.Iterator[Database] {

	getNextPage := func(ctx context.Context, req ListDatabasesRequest) (*ListDatabasesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatabases(ctx, req)
	}
	getItems := func(resp *ListDatabasesResponse) []Database {
		return resp.Databases
	}
	getNextReq := func(resp *ListDatabasesResponse) *ListDatabasesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List Databases.
func (a *postgresImpl) ListDatabasesAll(ctx context.Context, request ListDatabasesRequest) ([]Database, error) {
	iterator := a.ListDatabases(ctx, request)
	return listing.ToSlice[Database](ctx, iterator)
}

func (a *postgresImpl) internalListDatabases(ctx context.Context, request ListDatabasesRequest) (*ListDatabasesResponse, error) {
	var listDatabasesResponse ListDatabasesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/databases", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatabasesResponse)
	return &listDatabasesResponse, err
}

// Returns a paginated list of compute endpoints in the branch.
func (a *postgresImpl) ListEndpoints(ctx context.Context, request ListEndpointsRequest) listing.Iterator[Endpoint] {

	getNextPage := func(ctx context.Context, req ListEndpointsRequest) (*ListEndpointsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListEndpoints(ctx, req)
	}
	getItems := func(resp *ListEndpointsResponse) []Endpoint {
		return resp.Endpoints
	}
	getNextReq := func(resp *ListEndpointsResponse) *ListEndpointsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Returns a paginated list of compute endpoints in the branch.
func (a *postgresImpl) ListEndpointsAll(ctx context.Context, request ListEndpointsRequest) ([]Endpoint, error) {
	iterator := a.ListEndpoints(ctx, request)
	return listing.ToSlice[Endpoint](ctx, iterator)
}

func (a *postgresImpl) internalListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error) {
	var listEndpointsResponse ListEndpointsResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/endpoints", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listEndpointsResponse)
	return &listEndpointsResponse, err
}

// Returns a paginated list of database projects in the workspace that the user
// has permission to access.
func (a *postgresImpl) ListProjects(ctx context.Context, request ListProjectsRequest) listing.Iterator[Project] {

	getNextPage := func(ctx context.Context, req ListProjectsRequest) (*ListProjectsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListProjects(ctx, req)
	}
	getItems := func(resp *ListProjectsResponse) []Project {
		return resp.Projects
	}
	getNextReq := func(resp *ListProjectsResponse) *ListProjectsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Returns a paginated list of database projects in the workspace that the user
// has permission to access.
func (a *postgresImpl) ListProjectsAll(ctx context.Context, request ListProjectsRequest) ([]Project, error) {
	iterator := a.ListProjects(ctx, request)
	return listing.ToSlice[Project](ctx, iterator)
}

func (a *postgresImpl) internalListProjects(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error) {
	var listProjectsResponse ListProjectsResponse
	path := "/api/2.0/postgres/projects"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listProjectsResponse)
	return &listProjectsResponse, err
}

// Returns a paginated list of Postgres roles in the branch.
func (a *postgresImpl) ListRoles(ctx context.Context, request ListRolesRequest) listing.Iterator[Role] {

	getNextPage := func(ctx context.Context, req ListRolesRequest) (*ListRolesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListRoles(ctx, req)
	}
	getItems := func(resp *ListRolesResponse) []Role {
		return resp.Roles
	}
	getNextReq := func(resp *ListRolesResponse) *ListRolesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Returns a paginated list of Postgres roles in the branch.
func (a *postgresImpl) ListRolesAll(ctx context.Context, request ListRolesRequest) ([]Role, error) {
	iterator := a.ListRoles(ctx, request)
	return listing.ToSlice[Role](ctx, iterator)
}

func (a *postgresImpl) internalListRoles(ctx context.Context, request ListRolesRequest) (*ListRolesResponse, error) {
	var listRolesResponse ListRolesResponse
	path := fmt.Sprintf("/api/2.0/postgres/%v/roles", request.Parent)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listRolesResponse)
	return &listRolesResponse, err
}

func (a *postgresImpl) UndeleteBranch(ctx context.Context, request UndeleteBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/undelete", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) UndeleteProject(ctx context.Context, request UndeleteProjectRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v/undelete", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateBranch(ctx context.Context, request UpdateBranchRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Branch, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateDataApi(ctx context.Context, request UpdateDataApiRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DataApi, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateDatabase(ctx context.Context, request UpdateDatabaseRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Database, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Endpoint, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateProject(ctx context.Context, request UpdateProjectRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Project, &operation)
	return &operation, err
}

func (a *postgresImpl) UpdateRole(ctx context.Context, request UpdateRoleRequest) (*Operation, error) {
	var operation Operation
	path := fmt.Sprintf("/api/2.0/postgres/%v", request.Name)
	queryParams := make(map[string]any)

	updateMaskJson, updateMaskMarshallError := json.Marshal(request.UpdateMask)
	if updateMaskMarshallError != nil {
		return nil, updateMaskMarshallError
	}

	queryParams["update_mask"] = strings.Trim(string(updateMaskJson), `"`)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	cfg := a.client.Config
	if cfg.WorkspaceID != "" {
		headers["X-Databricks-Workspace-Id"] = cfg.WorkspaceID
	}
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Role, &operation)
	return &operation, err
}
