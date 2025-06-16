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

	requestPb, pbErr := createDatabaseCatalogRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseCatalogPb databaseCatalogPb
	path := "/api/2.0/database/catalogs"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).Catalog,
		&databaseCatalogPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseCatalogFromPb(&databaseCatalogPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) CreateDatabaseInstance(ctx context.Context, request CreateDatabaseInstanceRequest) (*DatabaseInstance, error) {

	requestPb, pbErr := createDatabaseInstanceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstancePb databaseInstancePb
	path := "/api/2.0/database/instances"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).DatabaseInstance,
		&databaseInstancePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceFromPb(&databaseInstancePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) CreateDatabaseInstanceRole(ctx context.Context, request CreateDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error) {

	requestPb, pbErr := createDatabaseInstanceRoleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstanceRolePb databaseInstanceRolePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles", requestPb.InstanceName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).DatabaseInstanceRole,
		&databaseInstanceRolePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceRoleFromPb(&databaseInstanceRolePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) CreateDatabaseTable(ctx context.Context, request CreateDatabaseTableRequest) (*DatabaseTable, error) {

	requestPb, pbErr := createDatabaseTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseTablePb databaseTablePb
	path := "/api/2.0/database/tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).Table,
		&databaseTablePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseTableFromPb(&databaseTablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) CreateSyncedDatabaseTable(ctx context.Context, request CreateSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error) {

	requestPb, pbErr := createSyncedDatabaseTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var syncedDatabaseTablePb syncedDatabaseTablePb
	path := "/api/2.0/database/synced_tables"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb).SyncedTable,
		&syncedDatabaseTablePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := syncedDatabaseTableFromPb(&syncedDatabaseTablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) DeleteDatabaseCatalog(ctx context.Context, request DeleteDatabaseCatalogRequest) error {

	requestPb, pbErr := deleteDatabaseCatalogRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteDatabaseCatalogResponsePb deleteDatabaseCatalogResponsePb
	path := fmt.Sprintf("/api/2.0/database/catalogs/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteDatabaseCatalogResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *databaseImpl) DeleteDatabaseInstance(ctx context.Context, request DeleteDatabaseInstanceRequest) error {

	requestPb, pbErr := deleteDatabaseInstanceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteDatabaseInstanceResponsePb deleteDatabaseInstanceResponsePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteDatabaseInstanceResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *databaseImpl) DeleteDatabaseInstanceRole(ctx context.Context, request DeleteDatabaseInstanceRoleRequest) error {

	requestPb, pbErr := deleteDatabaseInstanceRoleRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteDatabaseInstanceRoleResponsePb deleteDatabaseInstanceRoleResponsePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles/%v", requestPb.InstanceName, requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteDatabaseInstanceRoleResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *databaseImpl) DeleteDatabaseTable(ctx context.Context, request DeleteDatabaseTableRequest) error {

	requestPb, pbErr := deleteDatabaseTableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteDatabaseTableResponsePb deleteDatabaseTableResponsePb
	path := fmt.Sprintf("/api/2.0/database/tables/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteDatabaseTableResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *databaseImpl) DeleteSyncedDatabaseTable(ctx context.Context, request DeleteSyncedDatabaseTableRequest) error {

	requestPb, pbErr := deleteSyncedDatabaseTableRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteSyncedDatabaseTableResponsePb deleteSyncedDatabaseTableResponsePb
	path := fmt.Sprintf("/api/2.0/database/synced_tables/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteSyncedDatabaseTableResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *databaseImpl) FailoverDatabaseInstance(ctx context.Context, request FailoverDatabaseInstanceRequest) (*DatabaseInstance, error) {

	requestPb, pbErr := failoverDatabaseInstanceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstancePb databaseInstancePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v/failover", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseInstancePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceFromPb(&databaseInstancePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) FindDatabaseInstanceByUid(ctx context.Context, request FindDatabaseInstanceByUidRequest) (*DatabaseInstance, error) {

	requestPb, pbErr := findDatabaseInstanceByUidRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstancePb databaseInstancePb
	path := "/api/2.0/database/instances:findByUid"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseInstancePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceFromPb(&databaseInstancePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error) {

	requestPb, pbErr := generateDatabaseCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseCredentialPb databaseCredentialPb
	path := "/api/2.0/database/credentials"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseCredentialPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseCredentialFromPb(&databaseCredentialPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) GetDatabaseCatalog(ctx context.Context, request GetDatabaseCatalogRequest) (*DatabaseCatalog, error) {

	requestPb, pbErr := getDatabaseCatalogRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseCatalogPb databaseCatalogPb
	path := fmt.Sprintf("/api/2.0/database/catalogs/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseCatalogPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseCatalogFromPb(&databaseCatalogPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) GetDatabaseInstance(ctx context.Context, request GetDatabaseInstanceRequest) (*DatabaseInstance, error) {

	requestPb, pbErr := getDatabaseInstanceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstancePb databaseInstancePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseInstancePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceFromPb(&databaseInstancePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) GetDatabaseInstanceRole(ctx context.Context, request GetDatabaseInstanceRoleRequest) (*DatabaseInstanceRole, error) {

	requestPb, pbErr := getDatabaseInstanceRoleRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstanceRolePb databaseInstanceRolePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles/%v", requestPb.InstanceName, requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseInstanceRolePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceRoleFromPb(&databaseInstanceRolePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) GetDatabaseTable(ctx context.Context, request GetDatabaseTableRequest) (*DatabaseTable, error) {

	requestPb, pbErr := getDatabaseTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseTablePb databaseTablePb
	path := fmt.Sprintf("/api/2.0/database/tables/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&databaseTablePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseTableFromPb(&databaseTablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) GetSyncedDatabaseTable(ctx context.Context, request GetSyncedDatabaseTableRequest) (*SyncedDatabaseTable, error) {

	requestPb, pbErr := getSyncedDatabaseTableRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var syncedDatabaseTablePb syncedDatabaseTablePb
	path := fmt.Sprintf("/api/2.0/database/synced_tables/%v", requestPb.Name)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&syncedDatabaseTablePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := syncedDatabaseTableFromPb(&syncedDatabaseTablePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listDatabaseInstanceRolesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listDatabaseInstanceRolesResponsePb listDatabaseInstanceRolesResponsePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v/roles", requestPb.InstanceName)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listDatabaseInstanceRolesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listDatabaseInstanceRolesResponseFromPb(&listDatabaseInstanceRolesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
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

	requestPb, pbErr := listDatabaseInstancesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listDatabaseInstancesResponsePb listDatabaseInstancesResponsePb
	path := "/api/2.0/database/instances"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listDatabaseInstancesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listDatabaseInstancesResponseFromPb(&listDatabaseInstancesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *databaseImpl) UpdateDatabaseInstance(ctx context.Context, request UpdateDatabaseInstanceRequest) (*DatabaseInstance, error) {

	requestPb, pbErr := updateDatabaseInstanceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var databaseInstancePb databaseInstancePb
	path := fmt.Sprintf("/api/2.0/database/instances/%v", requestPb.Name)
	queryParams := make(map[string]any)
	if requestPb.UpdateMask != "" {
		queryParams["update_mask"] = requestPb.UpdateMask
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb).DatabaseInstance,
		&databaseInstancePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := databaseInstanceFromPb(&databaseInstancePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
