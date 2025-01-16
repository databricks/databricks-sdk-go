// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just account_metastore_assignments API methods
type accountMetastoreAssignmentsImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoreAssignmentsImpl) Create(ctx context.Context, request AccountsCreateMetastoreAssignment) error {
	var createResponse CreateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createResponse)
	return err
}

func (a *accountMetastoreAssignmentsImpl) Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *accountMetastoreAssignmentsImpl) Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*AccountsMetastoreAssignment, error) {
	var accountsMetastoreAssignment AccountsMetastoreAssignment
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastore", a.client.ConfiguredAccountID(), request.WorkspaceId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &accountsMetastoreAssignment)
	return &accountsMetastoreAssignment, err
}

func (a *accountMetastoreAssignmentsImpl) List(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error) {
	var listAccountMetastoreAssignmentsResponse ListAccountMetastoreAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/workspaces", a.client.ConfiguredAccountID(), request.MetastoreId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAccountMetastoreAssignmentsResponse)
	return &listAccountMetastoreAssignmentsResponse, err
}

func (a *accountMetastoreAssignmentsImpl) Update(ctx context.Context, request AccountsUpdateMetastoreAssignment) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/metastores/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.MetastoreId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just account_metastores API methods
type accountMetastoresImpl struct {
	client *client.DatabricksClient
}

func (a *accountMetastoresImpl) Create(ctx context.Context, request AccountsCreateMetastore) (*AccountsMetastoreInfo, error) {
	var accountsMetastoreInfo AccountsMetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &accountsMetastoreInfo)
	return &accountsMetastoreInfo, err
}

func (a *accountMetastoresImpl) Delete(ctx context.Context, request DeleteAccountMetastoreRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *accountMetastoresImpl) Get(ctx context.Context, request GetAccountMetastoreRequest) (*AccountsMetastoreInfo, error) {
	var accountsMetastoreInfo AccountsMetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &accountsMetastoreInfo)
	return &accountsMetastoreInfo, err
}

func (a *accountMetastoresImpl) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *accountMetastoresImpl) Update(ctx context.Context, request AccountsUpdateMetastore) (*AccountsMetastoreInfo, error) {
	var accountsMetastoreInfo AccountsMetastoreInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v", a.client.ConfiguredAccountID(), request.MetastoreId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &accountsMetastoreInfo)
	return &accountsMetastoreInfo, err
}

// unexported type that holds implementations of just account_storage_credentials API methods
type accountStorageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *accountStorageCredentialsImpl) Create(ctx context.Context, request AccountsCreateStorageCredential) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

func (a *accountStorageCredentialsImpl) Delete(ctx context.Context, request DeleteAccountStorageCredentialRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *accountStorageCredentialsImpl) Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

func (a *accountStorageCredentialsImpl) List(ctx context.Context, request ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error) {
	var listAccountStorageCredentialsResponse ListAccountStorageCredentialsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials", a.client.ConfiguredAccountID(), request.MetastoreId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listAccountStorageCredentialsResponse)
	return &listAccountStorageCredentialsResponse, err
}

func (a *accountStorageCredentialsImpl) Update(ctx context.Context, request AccountsUpdateStorageCredential) (*AccountsStorageCredentialInfo, error) {
	var accountsStorageCredentialInfo AccountsStorageCredentialInfo
	path := fmt.Sprintf("/api/2.0/accounts/%v/metastores/%v/storage-credentials/%v", a.client.ConfiguredAccountID(), request.MetastoreId, request.StorageCredentialName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &accountsStorageCredentialInfo)
	return &accountsStorageCredentialInfo, err
}

// unexported type that holds implementations of just artifact_allowlists API methods
type artifactAllowlistsImpl struct {
	client *client.DatabricksClient
}

func (a *artifactAllowlistsImpl) Get(ctx context.Context, request GetArtifactAllowlistRequest) (*ArtifactAllowlistInfo, error) {
	var artifactAllowlistInfo ArtifactAllowlistInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &artifactAllowlistInfo)
	return &artifactAllowlistInfo, err
}

func (a *artifactAllowlistsImpl) Update(ctx context.Context, request SetArtifactAllowlist) (*ArtifactAllowlistInfo, error) {
	var artifactAllowlistInfo ArtifactAllowlistInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/artifact-allowlists/%v", request.ArtifactType)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &artifactAllowlistInfo)
	return &artifactAllowlistInfo, err
}

// unexported type that holds implementations of just catalogs API methods
type catalogsImpl struct {
	client *client.DatabricksClient
}

func (a *catalogsImpl) Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := "/api/2.1/unity-catalog/catalogs"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &catalogInfo)
	return &catalogInfo, err
}

func (a *catalogsImpl) Delete(ctx context.Context, request DeleteCatalogRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *catalogsImpl) Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &catalogInfo)
	return &catalogInfo, err
}

func (a *catalogsImpl) List(ctx context.Context, request ListCatalogsRequest) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *catalogsImpl) Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error) {
	var catalogInfo CatalogInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &catalogInfo)
	return &catalogInfo, err
}

// unexported type that holds implementations of just connections API methods
type connectionsImpl struct {
	client *client.DatabricksClient
}

func (a *connectionsImpl) Create(ctx context.Context, request CreateConnection) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := "/api/2.1/unity-catalog/connections"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &connectionInfo)
	return &connectionInfo, err
}

func (a *connectionsImpl) Delete(ctx context.Context, request DeleteConnectionRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *connectionsImpl) Get(ctx context.Context, request GetConnectionRequest) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &connectionInfo)
	return &connectionInfo, err
}

func (a *connectionsImpl) List(ctx context.Context, request ListConnectionsRequest) (*ListConnectionsResponse, error) {
	var listConnectionsResponse ListConnectionsResponse
	path := "/api/2.1/unity-catalog/connections"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listConnectionsResponse)
	return &listConnectionsResponse, err
}

func (a *connectionsImpl) Update(ctx context.Context, request UpdateConnection) (*ConnectionInfo, error) {
	var connectionInfo ConnectionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/connections/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &connectionInfo)
	return &connectionInfo, err
}

// unexported type that holds implementations of just credentials API methods
type credentialsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsImpl) CreateCredential(ctx context.Context, request CreateCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := "/api/2.1/unity-catalog/credentials"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error {
	var deleteCredentialResponse DeleteCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteCredentialResponse)
	return err
}

func (a *credentialsImpl) GenerateTemporaryServiceCredential(ctx context.Context, request GenerateTemporaryServiceCredentialRequest) (*TemporaryCredentials, error) {
	var temporaryCredentials TemporaryCredentials
	path := "/api/2.1/unity-catalog/temporary-service-credentials"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &temporaryCredentials)
	return &temporaryCredentials, err
}

func (a *credentialsImpl) GetCredential(ctx context.Context, request GetCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) ListCredentials(ctx context.Context, request ListCredentialsRequest) (*ListCredentialsResponse, error) {
	var listCredentialsResponse ListCredentialsResponse
	path := "/api/2.1/unity-catalog/credentials"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listCredentialsResponse)
	return &listCredentialsResponse, err
}

func (a *credentialsImpl) UpdateCredential(ctx context.Context, request UpdateCredentialRequest) (*CredentialInfo, error) {
	var credentialInfo CredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/credentials/%v", request.NameArg)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &credentialInfo)
	return &credentialInfo, err
}

func (a *credentialsImpl) ValidateCredential(ctx context.Context, request ValidateCredentialRequest) (*ValidateCredentialResponse, error) {
	var validateCredentialResponse ValidateCredentialResponse
	path := "/api/2.1/unity-catalog/validate-credentials"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &validateCredentialResponse)
	return &validateCredentialResponse, err
}

// unexported type that holds implementations of just external_locations API methods
type externalLocationsImpl struct {
	client *client.DatabricksClient
}

func (a *externalLocationsImpl) Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := "/api/2.1/unity-catalog/external-locations"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

func (a *externalLocationsImpl) Delete(ctx context.Context, request DeleteExternalLocationRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *externalLocationsImpl) Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

func (a *externalLocationsImpl) List(ctx context.Context, request ListExternalLocationsRequest) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *externalLocationsImpl) Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error) {
	var externalLocationInfo ExternalLocationInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &externalLocationInfo)
	return &externalLocationInfo, err
}

// unexported type that holds implementations of just functions API methods
type functionsImpl struct {
	client *client.DatabricksClient
}

func (a *functionsImpl) Create(ctx context.Context, request CreateFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := "/api/2.1/unity-catalog/functions"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &functionInfo)
	return &functionInfo, err
}

func (a *functionsImpl) Delete(ctx context.Context, request DeleteFunctionRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *functionsImpl) Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &functionInfo)
	return &functionInfo, err
}

func (a *functionsImpl) List(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error) {
	var listFunctionsResponse ListFunctionsResponse
	path := "/api/2.1/unity-catalog/functions"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listFunctionsResponse)
	return &listFunctionsResponse, err
}

func (a *functionsImpl) Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error) {
	var functionInfo FunctionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/functions/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &functionInfo)
	return &functionInfo, err
}

// unexported type that holds implementations of just grants API methods
type grantsImpl struct {
	client *client.DatabricksClient
}

func (a *grantsImpl) Get(ctx context.Context, request GetGrantRequest) (*PermissionsList, error) {
	var permissionsList PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &permissionsList)
	return &permissionsList, err
}

func (a *grantsImpl) GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error) {
	var effectivePermissionsList EffectivePermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/effective-permissions/%v/%v", request.SecurableType, request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &effectivePermissionsList)
	return &effectivePermissionsList, err
}

func (a *grantsImpl) Update(ctx context.Context, request UpdatePermissions) (*PermissionsList, error) {
	var permissionsList PermissionsList
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &permissionsList)
	return &permissionsList, err
}

// unexported type that holds implementations of just metastores API methods
type metastoresImpl struct {
	client *client.DatabricksClient
}

func (a *metastoresImpl) Assign(ctx context.Context, request CreateMetastoreAssignment) error {
	var assignResponse AssignResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &assignResponse)
	return err
}

func (a *metastoresImpl) Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := "/api/2.1/unity-catalog/metastores"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) Current(ctx context.Context) (*MetastoreAssignment, error) {
	var metastoreAssignment MetastoreAssignment
	path := "/api/2.1/unity-catalog/current-metastore-assignment"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &metastoreAssignment)
	return &metastoreAssignment, err
}

func (a *metastoresImpl) Delete(ctx context.Context, request DeleteMetastoreRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *metastoresImpl) Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *metastoresImpl) Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *metastoresImpl) Unassign(ctx context.Context, request UnassignRequest) error {
	var unassignResponse UnassignResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &unassignResponse)
	return err
}

func (a *metastoresImpl) Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error) {
	var metastoreInfo MetastoreInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &metastoreInfo)
	return &metastoreInfo, err
}

func (a *metastoresImpl) UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	var updateAssignmentResponse UpdateAssignmentResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateAssignmentResponse)
	return err
}

// unexported type that holds implementations of just model_versions API methods
type modelVersionsImpl struct {
	client *client.DatabricksClient
}

func (a *modelVersionsImpl) Delete(ctx context.Context, request DeleteModelVersionRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *modelVersionsImpl) Get(ctx context.Context, request GetModelVersionRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

func (a *modelVersionsImpl) GetByAlias(ctx context.Context, request GetByAliasRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

func (a *modelVersionsImpl) List(ctx context.Context, request ListModelVersionsRequest) (*ListModelVersionsResponse, error) {
	var listModelVersionsResponse ListModelVersionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listModelVersionsResponse)
	return &listModelVersionsResponse, err
}

func (a *modelVersionsImpl) Update(ctx context.Context, request UpdateModelVersionRequest) (*ModelVersionInfo, error) {
	var modelVersionInfo ModelVersionInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/versions/%v", request.FullName, request.Version)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &modelVersionInfo)
	return &modelVersionInfo, err
}

// unexported type that holds implementations of just online_tables API methods
type onlineTablesImpl struct {
	client *client.DatabricksClient
}

func (a *onlineTablesImpl) Create(ctx context.Context, request CreateOnlineTableRequest) (*OnlineTable, error) {
	var onlineTable OnlineTable
	path := "/api/2.0/online-tables"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request.Table, &onlineTable)
	return &onlineTable, err
}

func (a *onlineTablesImpl) Delete(ctx context.Context, request DeleteOnlineTableRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *onlineTablesImpl) Get(ctx context.Context, request GetOnlineTableRequest) (*OnlineTable, error) {
	var onlineTable OnlineTable
	path := fmt.Sprintf("/api/2.0/online-tables/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &onlineTable)
	return &onlineTable, err
}

// unexported type that holds implementations of just quality_monitors API methods
type qualityMonitorsImpl struct {
	client *client.DatabricksClient
}

func (a *qualityMonitorsImpl) CancelRefresh(ctx context.Context, request CancelRefreshRequest) error {
	var cancelRefreshResponse CancelRefreshResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v/cancel", request.TableName, request.RefreshId)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &cancelRefreshResponse)
	return err
}

func (a *qualityMonitorsImpl) Create(ctx context.Context, request CreateMonitor) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &monitorInfo)
	return &monitorInfo, err
}

func (a *qualityMonitorsImpl) Delete(ctx context.Context, request DeleteQualityMonitorRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *qualityMonitorsImpl) Get(ctx context.Context, request GetQualityMonitorRequest) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &monitorInfo)
	return &monitorInfo, err
}

func (a *qualityMonitorsImpl) GetRefresh(ctx context.Context, request GetRefreshRequest) (*MonitorRefreshInfo, error) {
	var monitorRefreshInfo MonitorRefreshInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes/%v", request.TableName, request.RefreshId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &monitorRefreshInfo)
	return &monitorRefreshInfo, err
}

func (a *qualityMonitorsImpl) ListRefreshes(ctx context.Context, request ListRefreshesRequest) (*MonitorRefreshListResponse, error) {
	var monitorRefreshListResponse MonitorRefreshListResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &monitorRefreshListResponse)
	return &monitorRefreshListResponse, err
}

func (a *qualityMonitorsImpl) RegenerateDashboard(ctx context.Context, request RegenerateDashboardRequest) (*RegenerateDashboardResponse, error) {
	var regenerateDashboardResponse RegenerateDashboardResponse
	path := fmt.Sprintf("/api/2.1/quality-monitoring/tables/%v/monitor/dashboard", request.TableName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &regenerateDashboardResponse)
	return &regenerateDashboardResponse, err
}

func (a *qualityMonitorsImpl) RunRefresh(ctx context.Context, request RunRefreshRequest) (*MonitorRefreshInfo, error) {
	var monitorRefreshInfo MonitorRefreshInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor/refreshes", request.TableName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, nil, &monitorRefreshInfo)
	return &monitorRefreshInfo, err
}

func (a *qualityMonitorsImpl) Update(ctx context.Context, request UpdateMonitor) (*MonitorInfo, error) {
	var monitorInfo MonitorInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/monitor", request.TableName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &monitorInfo)
	return &monitorInfo, err
}

// unexported type that holds implementations of just registered_models API methods
type registeredModelsImpl struct {
	client *client.DatabricksClient
}

func (a *registeredModelsImpl) Create(ctx context.Context, request CreateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := "/api/2.1/unity-catalog/models"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

func (a *registeredModelsImpl) Delete(ctx context.Context, request DeleteRegisteredModelRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *registeredModelsImpl) DeleteAlias(ctx context.Context, request DeleteAliasRequest) error {
	var deleteAliasResponse DeleteAliasResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteAliasResponse)
	return err
}

func (a *registeredModelsImpl) Get(ctx context.Context, request GetRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

func (a *registeredModelsImpl) List(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error) {
	var listRegisteredModelsResponse ListRegisteredModelsResponse
	path := "/api/2.1/unity-catalog/models"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listRegisteredModelsResponse)
	return &listRegisteredModelsResponse, err
}

func (a *registeredModelsImpl) SetAlias(ctx context.Context, request SetRegisteredModelAliasRequest) (*RegisteredModelAlias, error) {
	var registeredModelAlias RegisteredModelAlias
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v/aliases/%v", request.FullName, request.Alias)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &registeredModelAlias)
	return &registeredModelAlias, err
}

func (a *registeredModelsImpl) Update(ctx context.Context, request UpdateRegisteredModelRequest) (*RegisteredModelInfo, error) {
	var registeredModelInfo RegisteredModelInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/models/%v", request.FullName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &registeredModelInfo)
	return &registeredModelInfo, err
}

// unexported type that holds implementations of just resource_quotas API methods
type resourceQuotasImpl struct {
	client *client.DatabricksClient
}

func (a *resourceQuotasImpl) GetQuota(ctx context.Context, request GetQuotaRequest) (*GetQuotaResponse, error) {
	var getQuotaResponse GetQuotaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/resource-quotas/%v/%v/%v", request.ParentSecurableType, request.ParentFullName, request.QuotaName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getQuotaResponse)
	return &getQuotaResponse, err
}

func (a *resourceQuotasImpl) ListQuotas(ctx context.Context, request ListQuotasRequest) (*ListQuotasResponse, error) {
	var listQuotasResponse ListQuotasResponse
	path := "/api/2.1/unity-catalog/resource-quotas/all-resource-quotas"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listQuotasResponse)
	return &listQuotasResponse, err
}

// unexported type that holds implementations of just schemas API methods
type schemasImpl struct {
	client *client.DatabricksClient
}

func (a *schemasImpl) Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := "/api/2.1/unity-catalog/schemas"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &schemaInfo)
	return &schemaInfo, err
}

func (a *schemasImpl) Delete(ctx context.Context, request DeleteSchemaRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *schemasImpl) Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &schemaInfo)
	return &schemaInfo, err
}

func (a *schemasImpl) List(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

func (a *schemasImpl) Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error) {
	var schemaInfo SchemaInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &schemaInfo)
	return &schemaInfo, err
}

// unexported type that holds implementations of just storage_credentials API methods
type storageCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *storageCredentialsImpl) Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := "/api/2.1/unity-catalog/storage-credentials"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Delete(ctx context.Context, request DeleteStorageCredentialRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *storageCredentialsImpl) Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) List(ctx context.Context, request ListStorageCredentialsRequest) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *storageCredentialsImpl) Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error) {
	var storageCredentialInfo StorageCredentialInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &storageCredentialInfo)
	return &storageCredentialInfo, err
}

func (a *storageCredentialsImpl) Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error) {
	var validateStorageCredentialResponse ValidateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/validate-storage-credentials"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &validateStorageCredentialResponse)
	return &validateStorageCredentialResponse, err
}

// unexported type that holds implementations of just system_schemas API methods
type systemSchemasImpl struct {
	client *client.DatabricksClient
}

func (a *systemSchemasImpl) Disable(ctx context.Context, request DisableRequest) error {
	var disableResponse DisableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &disableResponse)
	return err
}

func (a *systemSchemasImpl) Enable(ctx context.Context, request EnableRequest) error {
	var enableResponse EnableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas/%v", request.MetastoreId, request.SchemaName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, nil, &enableResponse)
	return err
}

func (a *systemSchemasImpl) List(ctx context.Context, request ListSystemSchemasRequest) (*ListSystemSchemasResponse, error) {
	var listSystemSchemasResponse ListSystemSchemasResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v/systemschemas", request.MetastoreId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listSystemSchemasResponse)
	return &listSystemSchemasResponse, err
}

// unexported type that holds implementations of just table_constraints API methods
type tableConstraintsImpl struct {
	client *client.DatabricksClient
}

func (a *tableConstraintsImpl) Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error) {
	var tableConstraint TableConstraint
	path := "/api/2.1/unity-catalog/constraints"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &tableConstraint)
	return &tableConstraint, err
}

func (a *tableConstraintsImpl) Delete(ctx context.Context, request DeleteTableConstraintRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/constraints/%v", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

// unexported type that holds implementations of just tables API methods
type tablesImpl struct {
	client *client.DatabricksClient
}

func (a *tablesImpl) Delete(ctx context.Context, request DeleteTableRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *tablesImpl) Exists(ctx context.Context, request ExistsRequest) (*TableExistsResponse, error) {
	var tableExistsResponse TableExistsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v/exists", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &tableExistsResponse)
	return &tableExistsResponse, err
}

func (a *tablesImpl) Get(ctx context.Context, request GetTableRequest) (*TableInfo, error) {
	var tableInfo TableInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &tableInfo)
	return &tableInfo, err
}

func (a *tablesImpl) List(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listTablesResponse)
	return &listTablesResponse, err
}

func (a *tablesImpl) ListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

func (a *tablesImpl) Update(ctx context.Context, request UpdateTableRequest) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just temporary_table_credentials API methods
type temporaryTableCredentialsImpl struct {
	client *client.DatabricksClient
}

func (a *temporaryTableCredentialsImpl) GenerateTemporaryTableCredentials(ctx context.Context, request GenerateTemporaryTableCredentialRequest) (*GenerateTemporaryTableCredentialResponse, error) {
	var generateTemporaryTableCredentialResponse GenerateTemporaryTableCredentialResponse
	path := "/api/2.0/unity-catalog/temporary-table-credentials"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &generateTemporaryTableCredentialResponse)
	return &generateTemporaryTableCredentialResponse, err
}

// unexported type that holds implementations of just volumes API methods
type volumesImpl struct {
	client *client.DatabricksClient
}

func (a *volumesImpl) Create(ctx context.Context, request CreateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := "/api/2.1/unity-catalog/volumes"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Delete(ctx context.Context, request DeleteVolumeRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *volumesImpl) List(ctx context.Context, request ListVolumesRequest) (*ListVolumesResponseContent, error) {
	var listVolumesResponseContent ListVolumesResponseContent
	path := "/api/2.1/unity-catalog/volumes"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listVolumesResponseContent)
	return &listVolumesResponseContent, err
}

func (a *volumesImpl) Read(ctx context.Context, request ReadVolumeRequest) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &volumeInfo)
	return &volumeInfo, err
}

func (a *volumesImpl) Update(ctx context.Context, request UpdateVolumeRequestContent) (*VolumeInfo, error) {
	var volumeInfo VolumeInfo
	path := fmt.Sprintf("/api/2.1/unity-catalog/volumes/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &volumeInfo)
	return &volumeInfo, err
}

// unexported type that holds implementations of just workspace_bindings API methods
type workspaceBindingsImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceBindingsImpl) Get(ctx context.Context, request GetWorkspaceBindingRequest) (*CurrentWorkspaceBindings, error) {
	var currentWorkspaceBindings CurrentWorkspaceBindings
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &currentWorkspaceBindings)
	return &currentWorkspaceBindings, err
}

func (a *workspaceBindingsImpl) GetBindings(ctx context.Context, request GetBindingsRequest) (*WorkspaceBindingsResponse, error) {
	var workspaceBindingsResponse WorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &workspaceBindingsResponse)
	return &workspaceBindingsResponse, err
}

func (a *workspaceBindingsImpl) Update(ctx context.Context, request UpdateWorkspaceBindings) (*CurrentWorkspaceBindings, error) {
	var currentWorkspaceBindings CurrentWorkspaceBindings
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspace-bindings/catalogs/%v", request.Name)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &currentWorkspaceBindings)
	return &currentWorkspaceBindings, err
}

func (a *workspaceBindingsImpl) UpdateBindings(ctx context.Context, request UpdateWorkspaceBindingsParameters) (*WorkspaceBindingsResponse, error) {
	var workspaceBindingsResponse WorkspaceBindingsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/bindings/%v/%v", request.SecurableType, request.SecurableName)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &workspaceBindingsResponse)
	return &workspaceBindingsResponse, err
}
