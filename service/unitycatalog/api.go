// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type UnitycatalogService interface {
	CreateCatalog(ctx context.Context, createCatalogRequest CreateCatalogRequest) (*CreateCatalogResponse, error)

	CreateExternalLocation(ctx context.Context, createExternalLocationRequest CreateExternalLocationRequest) (*CreateExternalLocationResponse, error)

	CreateMetastore(ctx context.Context, createMetastoreRequest CreateMetastoreRequest) (*CreateMetastoreResponse, error)

	CreateMetastoreAssignment(ctx context.Context, createMetastoreAssignmentRequest CreateMetastoreAssignmentRequest) error

	CreateSchema(ctx context.Context, createSchemaRequest CreateSchemaRequest) (*CreateSchemaResponse, error)

	CreateStagingTable(ctx context.Context, createStagingTableRequest CreateStagingTableRequest) (*CreateStagingTableResponse, error)

	CreateStorageCredential(ctx context.Context, createStorageCredentialRequest CreateStorageCredentialRequest) (*CreateStorageCredentialResponse, error)

	CreateTable(ctx context.Context, createTableRequest CreateTableRequest) (*CreateTableResponse, error)

	DeleteCatalog(ctx context.Context, deleteCatalogRequest DeleteCatalogRequest) error

	DeleteExternalLocation(ctx context.Context, deleteExternalLocationRequest DeleteExternalLocationRequest) error

	DeleteMetastore(ctx context.Context, deleteMetastoreRequest DeleteMetastoreRequest) error

	DeleteMetastoreAssignment(ctx context.Context, deleteMetastoreAssignmentRequest DeleteMetastoreAssignmentRequest) error

	DeleteSchema(ctx context.Context, deleteSchemaRequest DeleteSchemaRequest) error

	DeleteStorageCredential(ctx context.Context, deleteStorageCredentialRequest DeleteStorageCredentialRequest) error

	DeleteTable(ctx context.Context, deleteTableRequest DeleteTableRequest) error

	GetCatalog(ctx context.Context, getCatalogRequest GetCatalogRequest) (*GetCatalogResponse, error)

	GetExternalLocation(ctx context.Context, getExternalLocationRequest GetExternalLocationRequest) (*GetExternalLocationResponse, error)

	GetMetastore(ctx context.Context, getMetastoreRequest GetMetastoreRequest) (*GetMetastoreResponse, error)

	GetMetastoreSummary(ctx context.Context) (*GetMetastoreSummaryResponse, error)

	GetPermissions(ctx context.Context, getPermissionsRequest GetPermissionsRequest) (*GetPermissionsResponse, error)

	GetSchema(ctx context.Context, getSchemaRequest GetSchemaRequest) (*GetSchemaResponse, error)

	GetStorageCredential(ctx context.Context, getStorageCredentialRequest GetStorageCredentialRequest) (*GetStorageCredentialResponse, error)

	GetTable(ctx context.Context, getTableRequest GetTableRequest) (*GetTableResponse, error)

	ListCatalogs(ctx context.Context) (*ListCatalogsResponse, error)

	ListExternalLocations(ctx context.Context) (*ListExternalLocationsResponse, error)

	ListMetastores(ctx context.Context) (*ListMetastoresResponse, error)

	ListSchemas(ctx context.Context, listSchemasRequest ListSchemasRequest) (*ListSchemasResponse, error)

	ListStorageCredentials(ctx context.Context) (*ListStorageCredentialsResponse, error)

	ListTables(ctx context.Context, listTablesRequest ListTablesRequest) (*ListTablesResponse, error)

	UpdateCatalog(ctx context.Context, updateCatalogRequest UpdateCatalogRequest) error

	UpdateExternalLocation(ctx context.Context, updateExternalLocationRequest UpdateExternalLocationRequest) error

	UpdateMetastore(ctx context.Context, updateMetastoreRequest UpdateMetastoreRequest) error

	UpdateMetastoreAssignment(ctx context.Context, updateMetastoreAssignmentRequest UpdateMetastoreAssignmentRequest) error

	UpdatePermissions(ctx context.Context, updatePermissionsRequest UpdatePermissionsRequest) error

	UpdateSchema(ctx context.Context, updateSchemaRequest UpdateSchemaRequest) error

	UpdateStorageCredential(ctx context.Context, updateStorageCredentialRequest UpdateStorageCredentialRequest) error

	UpdateTable(ctx context.Context, updateTableRequest UpdateTableRequest) error
	GetPermissionsBySecurableTypeAndSecurableFullName(ctx context.Context, securableType string, securableFullName string) (*GetPermissionsResponse, error)
	GetExternalLocationByNameArg(ctx context.Context, nameArg string) (*GetExternalLocationResponse, error)
	DeleteExternalLocationByNameArg(ctx context.Context, nameArg string) error
	DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error
	GetStorageCredentialByNameArg(ctx context.Context, nameArg string) (*GetStorageCredentialResponse, error)
	DeleteStorageCredentialByNameArg(ctx context.Context, nameArg string) error
	DeleteCatalogByNameArg(ctx context.Context, nameArg string) error
	GetCatalogByNameArg(ctx context.Context, nameArg string) (*GetCatalogResponse, error)
	GetTableByFullNameArg(ctx context.Context, fullNameArg string) (*GetTableResponse, error)
	DeleteTableByFullNameArg(ctx context.Context, fullNameArg string) error
	GetMetastoreById(ctx context.Context, id string) (*GetMetastoreResponse, error)
	DeleteMetastoreById(ctx context.Context, id string) error
	GetSchemaByFullNameArg(ctx context.Context, fullNameArg string) (*GetSchemaResponse, error)
	DeleteSchemaByFullNameArg(ctx context.Context, fullNameArg string) error
}

func New(client *client.DatabricksClient) UnitycatalogService {
	return &UnitycatalogAPI{
		client: client,
	}
}

type UnitycatalogAPI struct {
	client *client.DatabricksClient
}

func (a *UnitycatalogAPI) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (*CreateCatalogResponse, error) {
	var createCatalogResponse CreateCatalogResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Post(ctx, path, request, &createCatalogResponse)
	return &createCatalogResponse, err
}

func (a *UnitycatalogAPI) CreateExternalLocation(ctx context.Context, request CreateExternalLocationRequest) (*CreateExternalLocationResponse, error) {
	var createExternalLocationResponse CreateExternalLocationResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Post(ctx, path, request, &createExternalLocationResponse)
	return &createExternalLocationResponse, err
}

func (a *UnitycatalogAPI) CreateMetastore(ctx context.Context, request CreateMetastoreRequest) (*CreateMetastoreResponse, error) {
	var createMetastoreResponse CreateMetastoreResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Post(ctx, path, request, &createMetastoreResponse)
	return &createMetastoreResponse, err
}

func (a *UnitycatalogAPI) CreateMetastoreAssignment(ctx context.Context, request CreateMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) CreateSchema(ctx context.Context, request CreateSchemaRequest) (*CreateSchemaResponse, error) {
	var createSchemaResponse CreateSchemaResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Post(ctx, path, request, &createSchemaResponse)
	return &createSchemaResponse, err
}

func (a *UnitycatalogAPI) CreateStagingTable(ctx context.Context, request CreateStagingTableRequest) (*CreateStagingTableResponse, error) {
	var createStagingTableResponse CreateStagingTableResponse
	path := "/api/2.1/unity-catalog/staging-tables"
	err := a.client.Post(ctx, path, request, &createStagingTableResponse)
	return &createStagingTableResponse, err
}

func (a *UnitycatalogAPI) CreateStorageCredential(ctx context.Context, request CreateStorageCredentialRequest) (*CreateStorageCredentialResponse, error) {
	var createStorageCredentialResponse CreateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Post(ctx, path, request, &createStorageCredentialResponse)
	return &createStorageCredentialResponse, err
}

func (a *UnitycatalogAPI) CreateTable(ctx context.Context, request CreateTableRequest) (*CreateTableResponse, error) {
	var createTableResponse CreateTableResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Post(ctx, path, request, &createTableResponse)
	return &createTableResponse, err
}

func (a *UnitycatalogAPI) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) DeleteExternalLocation(ctx context.Context, request DeleteExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) DeleteMetastore(ctx context.Context, request DeleteMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) DeleteMetastoreAssignment(ctx context.Context, request DeleteMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) DeleteSchema(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullNameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) DeleteStorageCredential(ctx context.Context, request DeleteStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) DeleteTable(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullNameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) GetCatalog(ctx context.Context, request GetCatalogRequest) (*GetCatalogResponse, error) {
	var getCatalogResponse GetCatalogResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getCatalogResponse)
	return &getCatalogResponse, err
}

func (a *UnitycatalogAPI) GetExternalLocation(ctx context.Context, request GetExternalLocationRequest) (*GetExternalLocationResponse, error) {
	var getExternalLocationResponse GetExternalLocationResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getExternalLocationResponse)
	return &getExternalLocationResponse, err
}

func (a *UnitycatalogAPI) GetMetastore(ctx context.Context, request GetMetastoreRequest) (*GetMetastoreResponse, error) {
	var getMetastoreResponse GetMetastoreResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Get(ctx, path, request, &getMetastoreResponse)
	return &getMetastoreResponse, err
}

func (a *UnitycatalogAPI) GetMetastoreSummary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"
	err := a.client.Get(ctx, path, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *UnitycatalogAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.SecurableFullName)
	err := a.client.Get(ctx, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

func (a *UnitycatalogAPI) GetSchema(ctx context.Context, request GetSchemaRequest) (*GetSchemaResponse, error) {
	var getSchemaResponse GetSchemaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullNameArg)
	err := a.client.Get(ctx, path, request, &getSchemaResponse)
	return &getSchemaResponse, err
}

func (a *UnitycatalogAPI) GetStorageCredential(ctx context.Context, request GetStorageCredentialRequest) (*GetStorageCredentialResponse, error) {
	var getStorageCredentialResponse GetStorageCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getStorageCredentialResponse)
	return &getStorageCredentialResponse, err
}

func (a *UnitycatalogAPI) GetTable(ctx context.Context, request GetTableRequest) (*GetTableResponse, error) {
	var getTableResponse GetTableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullNameArg)
	err := a.client.Get(ctx, path, request, &getTableResponse)
	return &getTableResponse, err
}

func (a *UnitycatalogAPI) ListCatalogs(ctx context.Context) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Get(ctx, path, nil, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *UnitycatalogAPI) ListExternalLocations(ctx context.Context) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Get(ctx, path, nil, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *UnitycatalogAPI) ListMetastores(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Get(ctx, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *UnitycatalogAPI) ListSchemas(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Get(ctx, path, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *UnitycatalogAPI) ListStorageCredentials(ctx context.Context) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Get(ctx, path, nil, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *UnitycatalogAPI) ListTables(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Get(ctx, path, request, &listTablesResponse)
	return &listTablesResponse, err
}

func (a *UnitycatalogAPI) UpdateCatalog(ctx context.Context, request UpdateCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.NameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdateExternalLocation(ctx context.Context, request UpdateExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.NameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdateMetastore(ctx context.Context, request UpdateMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdateMetastoreAssignment(ctx context.Context, request UpdateMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdatePermissions(ctx context.Context, request UpdatePermissionsRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.SecurableFullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdateSchema(ctx context.Context, request UpdateSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullNameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdateStorageCredential(ctx context.Context, request UpdateStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.NameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) UpdateTable(ctx context.Context, request UpdateTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullNameArg)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *UnitycatalogAPI) GetPermissionsBySecurableTypeAndSecurableFullName(ctx context.Context, securableType string, securableFullName string) (*GetPermissionsResponse, error) {
	return a.GetPermissions(ctx, GetPermissionsRequest{
		SecurableType:     securableType,
		SecurableFullName: securableFullName,
	})
}

func (a *UnitycatalogAPI) GetExternalLocationByNameArg(ctx context.Context, nameArg string) (*GetExternalLocationResponse, error) {
	return a.GetExternalLocation(ctx, GetExternalLocationRequest{
		NameArg: nameArg,
	})
}

func (a *UnitycatalogAPI) DeleteExternalLocationByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteExternalLocation(ctx, DeleteExternalLocationRequest{
		NameArg: nameArg,
	})
}

func (a *UnitycatalogAPI) DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error {
	return a.DeleteMetastoreAssignment(ctx, DeleteMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

func (a *UnitycatalogAPI) GetStorageCredentialByNameArg(ctx context.Context, nameArg string) (*GetStorageCredentialResponse, error) {
	return a.GetStorageCredential(ctx, GetStorageCredentialRequest{
		NameArg: nameArg,
	})
}

func (a *UnitycatalogAPI) DeleteStorageCredentialByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteStorageCredential(ctx, DeleteStorageCredentialRequest{
		NameArg: nameArg,
	})
}

func (a *UnitycatalogAPI) DeleteCatalogByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteCatalog(ctx, DeleteCatalogRequest{
		NameArg: nameArg,
	})
}

func (a *UnitycatalogAPI) GetCatalogByNameArg(ctx context.Context, nameArg string) (*GetCatalogResponse, error) {
	return a.GetCatalog(ctx, GetCatalogRequest{
		NameArg: nameArg,
	})
}

func (a *UnitycatalogAPI) GetTableByFullNameArg(ctx context.Context, fullNameArg string) (*GetTableResponse, error) {
	return a.GetTable(ctx, GetTableRequest{
		FullNameArg: fullNameArg,
	})
}

func (a *UnitycatalogAPI) DeleteTableByFullNameArg(ctx context.Context, fullNameArg string) error {
	return a.DeleteTable(ctx, DeleteTableRequest{
		FullNameArg: fullNameArg,
	})
}

func (a *UnitycatalogAPI) GetMetastoreById(ctx context.Context, id string) (*GetMetastoreResponse, error) {
	return a.GetMetastore(ctx, GetMetastoreRequest{
		Id: id,
	})
}

func (a *UnitycatalogAPI) DeleteMetastoreById(ctx context.Context, id string) error {
	return a.DeleteMetastore(ctx, DeleteMetastoreRequest{
		Id: id,
	})
}

func (a *UnitycatalogAPI) GetSchemaByFullNameArg(ctx context.Context, fullNameArg string) (*GetSchemaResponse, error) {
	return a.GetSchema(ctx, GetSchemaRequest{
		FullNameArg: fullNameArg,
	})
}

func (a *UnitycatalogAPI) DeleteSchemaByFullNameArg(ctx context.Context, fullNameArg string) error {
	return a.DeleteSchema(ctx, DeleteSchemaRequest{
		FullNameArg: fullNameArg,
	})
}
