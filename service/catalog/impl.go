// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountMetastoreAssignments API methods
type accountMetastoreAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoreAssignmentsImpl) Create(ctx context.Context, request CreateMetastoreAssignment) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &metastoreAssignment)
	return &metastoreAssignment, err
}

func (a *accountMetastoreAssignmentsImpl) Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountMetastoreAssignmentsImpl) Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastore", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &metastoreAssignment)
	return &metastoreAssignment, err
}

func (a *accountMetastoreAssignmentsImpl) List(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) ([]MetastoreAssignment, error) {
	var metastoreAssignmentList []MetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/workspaces", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &metastoreAssignmentList)
	return metastoreAssignmentList, err
}

func (a *accountMetastoreAssignmentsImpl) Update(ctx context.Context, request UpdateMetastoreAssignment) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &metastoreAssignment)
	return &metastoreAssignment, err
}

// unexported type that holds implementations of just AccountMetastores API methods
type accountMetastoresImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoresImpl) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *accountMetastoresImpl) Delete(ctx context.Context, request DeleteAccountMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountMetastoresImpl) Get(ctx context.Context, request GetAccountMetastoreRequest) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *accountMetastoresImpl) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *accountMetastoresImpl) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &metastoreInfo)
	return &metastoreInfo, err
}

// unexported type that holds implementations of just AccountStorageCredentials API methods
type accountStorageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *accountStorageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *accountStorageCredentialsImpl) Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *accountStorageCredentialsImpl) List(ctx context.Context, request ListAccountStorageCredentialsRequest) ([]StorageCredentialInfo, error) {
	var storageCredentialInfoList []StorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &storageCredentialInfoList)
	return storageCredentialInfoList, err
}

// unexported type that holds implementations of just Catalogs API methods
type catalogsImpl struct {
	client *client.DatabricksClient
}

func (a *catalogsImpl) Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Do(ctx, http.MethodPost, path, request, &catalogInfo)
	return &catalogInfo, err
}

func (a *catalogsImpl) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *catalogsImpl) Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &catalogInfo)
	return &catalogInfo, err
}

func (a *catalogsImpl) List(ctx context.Context) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *catalogsImpl) Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &catalogInfo)
	return &catalogInfo, err
}

// unexported type that holds implementations of just ExternalLocations API methods
type externalLocationsImpl struct {
	client *client.DatabricksClient
}

func (a *externalLocationsImpl) Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Do(ctx, http.MethodPost, path, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

func (a *externalLocationsImpl) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *externalLocationsImpl) Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

func (a *externalLocationsImpl) List(ctx context.Context) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *externalLocationsImpl) Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

// unexported type that holds implementations of just Functions API methods
type functionsImpl struct {
	client *client.DatabricksClient
}

func (a *functionsImpl) Create(ctx context.Context, request CreateFunction) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := "/api/2.1/unity-catalog/functions"
	err := a.client.Do(ctx, http.MethodPost, path, request, &functionInfo)
	return &functionInfo, err
}

func (a *functionsImpl) Delete(ctx context.Context, request DeleteFunctionRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *functionsImpl) Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &functionInfo)
	return &functionInfo, err
}

func (a *functionsImpl) List(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error) {
	var listFunctionsResponse ListFunctionsResponse
	path := "/api/2.1/unity-catalog/functions"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listFunctionsResponse)
	return &listFunctionsResponse, err
}

func (a *functionsImpl) Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &functionInfo)
	return &functionInfo, err
}

// unexported type that holds implementations of just Grants API methods
type grantsImpl struct {
	client *client.DatabricksClient
}

func (a *grantsImpl) Get(ctx context.Context, request GetGrantRequest) (*PermissionsList, error) {
	var permissionsList PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &permissionsList)
	return &permissionsList, err
}

func (a *grantsImpl) GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error) {
	var effectivePermissionsList EffectivePermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/effective-permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &effectivePermissionsList)
	return &effectivePermissionsList, err
}

func (a *grantsImpl) Update(ctx context.Context, request UpdatePermissions) (*PermissionsList, error) {
	var permissionsList PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &permissionsList)
	return &permissionsList, err
}

// unexported type that holds implementations of just Metastores API methods
type metastoresImpl struct {
	client *client.DatabricksClient
}

func (a *metastoresImpl) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *metastoresImpl) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Do(ctx, http.MethodPost, path, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) Current(ctx context.Context) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := "/api/2.1/unity-catalog/current-metastore-assignment"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &metastoreAssignment)
	return &metastoreAssignment, err
}

func (a *metastoresImpl) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *metastoresImpl) Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *metastoresImpl) Maintenance(ctx context.Context, request UpdateAutoMaintenance) (*UpdateAutoMaintenanceResponse, error) {
	var updateAutoMaintenanceResponse UpdateAutoMaintenanceResponse
	path := "/api/2.0/auto-maintenance/service"
	err := a.client.Do(ctx, http.MethodPatch, path, request, &updateAutoMaintenanceResponse)
	return &updateAutoMaintenanceResponse, err
}

func (a *metastoresImpl) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *metastoresImpl) Unassign(ctx context.Context, request UnassignRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *metastoresImpl) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just Schemas API methods
type schemasImpl struct {
	client *client.DatabricksClient
}

func (a *schemasImpl) Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Do(ctx, http.MethodPost, path, request, &schemaInfo)
	return &schemaInfo, err
}

func (a *schemasImpl) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *schemasImpl) Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &schemaInfo)
	return &schemaInfo, err
}

func (a *schemasImpl) List(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *schemasImpl) Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &schemaInfo)
	return &schemaInfo, err
}

// unexported type that holds implementations of just StorageCredentials API methods
type storageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *storageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Do(ctx, http.MethodPost, path, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *storageCredentialsImpl) Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Do(ctx, http.MethodGet, path, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) List(ctx context.Context) ([]StorageCredentialInfo, error) {
	var storageCredentialInfoList []StorageCredentialInfo
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &storageCredentialInfoList)
	return storageCredentialInfoList, err
}

func (a *storageCredentialsImpl) Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error) {
	var validateStorageCredentialResponse ValidateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/validate-storage-credentials"
	err := a.client.Do(ctx, http.MethodPost, path, request, &validateStorageCredentialResponse)
	return &validateStorageCredentialResponse, err
}

// unexported type that holds implementations of just TableConstraints API methods
type tableConstraintsImpl struct {
	client *client.DatabricksClient
}

func (a *tableConstraintsImpl) Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error) {
	var tableConstraint TableConstraint
	path := "/api/2.1/unity-catalog/constraints"
	err := a.client.Do(ctx, http.MethodPost, path, request, &tableConstraint)
	return &tableConstraint, err
}

func (a *tableConstraintsImpl) Delete(ctx context.Context, request DeleteTableConstraintRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/constraints/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

// unexported type that holds implementations of just Tables API methods
type tablesImpl struct {
	client *client.DatabricksClient
}

func (a *tablesImpl) Delete(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *tablesImpl) Get(ctx context.Context, request GetTableRequest) (*TableInfo, error) {
	var tableInfo TableInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Do(ctx, http.MethodGet, path, request, &tableInfo)
	return &tableInfo, err
}

func (a *tablesImpl) List(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listTablesResponse)
	return &listTablesResponse, err
}

func (a *tablesImpl) ListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

// unexported type that holds implementations of just Volumes API methods
type volumesImpl struct {
	client *client.DatabricksClient
}

func (a *volumesImpl) Create(ctx context.Context, request CreateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := "/api/2.1/unity-catalog/volumes"
	err := a.client.Do(ctx, http.MethodPost, path, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Delete(ctx context.Context, request DeleteVolumeRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.FullNameArg)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *volumesImpl) List(ctx context.Context, request ListVolumesRequest) (*ListVolumesResponseContent, error) {
	var listVolumesResponseContent ListVolumesResponseContent
	path := "/api/2.1/unity-catalog/volumes"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listVolumesResponseContent)
	return &listVolumesResponseContent, err
}

func (a *volumesImpl) Read(ctx context.Context, request ReadVolumeRequest) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.FullNameArg)
	err := a.client.Do(ctx, http.MethodGet, path, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Update(ctx context.Context, request UpdateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.FullNameArg)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &volumeInfo)
	return &volumeInfo, err
}
