// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just Database API methods
type databaseImpl struct {
	client *client.DatabricksClient
}

func (a *databaseImpl) CreateDatabaseCatalog(ctx context.Context, request CreateDatabaseCatalogRequest) (*DatabaseCatalog, error) {
	var databaseCatalog DatabaseCatalog
	path := "/api/2.0/database/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Catalog, &databaseCatalog)
	return &databaseCatalog, err
}

func (a *databaseImpl) CreateDatabaseInstance(ctx context.Context, request CreateDatabaseInstanceRequest) (*DatabaseInstance, error) {
	var databaseInstance DatabaseInstance
	path := "/api/2.0/database/instances"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DatabaseInstance, &databaseInstance)
	return &databaseInstance, err
}

func (a *databaseImpl) CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error) {
	var databaseInstanceRole DatabaseInstanceRole
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles", request.InstanceName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.DatabaseInstanceRole, &databaseInstanceRole)
	return &databaseInstanceRole, err
}

func (a *databaseImpl) CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error) {
	var databaseTable DatabaseTable
	path := "/api/2.0/database/tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.Table, &databaseTable)
	return &databaseTable, err
}

func (a *databaseImpl) CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error) {
	var syncedDatabaseTable SyncedDatabaseTable
	path := "/api/2.0/database/synced_tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request.SyncedTable, &syncedDatabaseTable)
	return &syncedDatabaseTable, err
}

func (a *databaseImpl) DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error {
	path := fmt.Sprintf("/api/2.0/database/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *databaseImpl) DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error {
	path := fmt.Sprintf("/api/2.0/database/instances/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *databaseImpl) DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error {
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles/%v", request.InstanceName, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *databaseImpl) DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error {
	path := fmt.Sprintf("/api/2.0/database/tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *databaseImpl) DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error {
	path := fmt.Sprintf("/api/2.0/database/synced_tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, nil)
	return err
}

func (a *databaseImpl) FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error) {
	var databaseInstance DatabaseInstance
	path := "/api/2.0/database/instances:findByUid"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseInstance)
	return &databaseInstance, err
}

func (a *databaseImpl) GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error) {
	var databaseCredential DatabaseCredential
	path := "/api/2.0/database/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &databaseCredential)
	return &databaseCredential, err
}

func (a *databaseImpl) GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error) {
	var databaseCatalog DatabaseCatalog
	path := fmt.Sprintf("/api/2.0/database/catalogs/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseCatalog)
	return &databaseCatalog, err
}

func (a *databaseImpl) GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error) {
	var databaseInstance DatabaseInstance
	path := fmt.Sprintf("/api/2.0/database/instances/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseInstance)
	return &databaseInstance, err
}

func (a *databaseImpl) GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error) {
	var databaseInstanceRole DatabaseInstanceRole
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles/%v", request.InstanceName, request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseInstanceRole)
	return &databaseInstanceRole, err
}

func (a *databaseImpl) GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error) {
	var databaseTable DatabaseTable
	path := fmt.Sprintf("/api/2.0/database/tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &databaseTable)
	return &databaseTable, err
}

func (a *databaseImpl) GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error) {
	var syncedDatabaseTable SyncedDatabaseTable
	path := fmt.Sprintf("/api/2.0/database/synced_tables/%v", request.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &syncedDatabaseTable)
	return &syncedDatabaseTable, err
}

// START OF PG ROLE APIs Section
func (a *databaseImpl) ListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) listing.Iterator[DatabaseInstanceRole] {

	getNextPage := func(ctx context.Context, req ListDatabaseInstanceRolesRequest) (*ListDatabaseInstanceRolesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatabaseInstanceRoles(ctx, req)
	}
	getItems := func(resp *ListDatabaseInstanceRolesResponse) []DatabaseInstanceRole {
		return resp.DatabaseInstanceRoles
	}
	getNextReq := func(resp *ListDatabaseInstanceRolesResponse) *ListDatabaseInstanceRolesRequest {
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

// START OF PG ROLE APIs Section
func (a *databaseImpl) ListDatabaseInstanceRolesAll(ctx context.Context, request ListDatabaseInstanceRolesRequest) ([]DatabaseInstanceRole, error) {
	iterator := a.ListDatabaseInstanceRoles(ctx, request)
	return listing.ToSlice[DatabaseInstanceRole](ctx, iterator)
}

func (a *databaseImpl) internalListDatabaseInstanceRoles(ctx context.Context, request ListDatabaseInstanceRolesRequest) (*ListDatabaseInstanceRolesResponse, error) {
	var listDatabaseInstanceRolesResponse ListDatabaseInstanceRolesResponse
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles", request.InstanceName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatabaseInstanceRolesResponse)
	return &listDatabaseInstanceRolesResponse, err
}

// List Database Instances.
func (a *databaseImpl) ListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) listing.Iterator[DatabaseInstance] {

	getNextPage := func(ctx context.Context, req ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListDatabaseInstances(ctx, req)
	}
	getItems := func(resp *ListDatabaseInstancesResponse) []DatabaseInstance {
		return resp.DatabaseInstances
	}
	getNextReq := func(resp *ListDatabaseInstancesResponse) *ListDatabaseInstancesRequest {
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

// List Database Instances.
func (a *databaseImpl) ListDatabaseInstancesAll(ctx context.Context, request ListDatabaseInstancesRequest) ([]DatabaseInstance, error) {
	iterator := a.ListDatabaseInstances(ctx, request)
	return listing.ToSlice[DatabaseInstance](ctx, iterator)
}

func (a *databaseImpl) internalListDatabaseInstances(ctx context.Context, request ListDatabaseInstancesRequest) (*ListDatabaseInstancesResponse, error) {
	var listDatabaseInstancesResponse ListDatabaseInstancesResponse
	path := "/api/2.0/database/instances"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listDatabaseInstancesResponse)
	return &listDatabaseInstancesResponse, err
}

func (a *databaseImpl) UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error) {
	var databaseInstance DatabaseInstance
	path := fmt.Sprintf("/api/2.0/database/instances/%v", request.Name)
	queryParams := make(map[string]any)
	if request.UpdateMask != "" {
		queryParams["update_mask"] = request.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.DatabaseInstance, &databaseInstance)
	return &databaseInstance, err
}
