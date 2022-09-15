// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewCatalogs(client *client.DatabricksClient) CatalogsService {
	return &CatalogsAPI{
		client: client,
	}
}

type CatalogsAPI struct {
	client *client.DatabricksClient
}

func (a *CatalogsAPI) Create(ctx context.Context, request CreateCatalog) (*CreateCatalogResponse, error) {
	var createCatalogResponse CreateCatalogResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Post(ctx, path, request, &createCatalogResponse)
	return &createCatalogResponse, err
}

func (a *CatalogsAPI) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *CatalogsAPI) DeleteCatalogByName(ctx context.Context, name string) error {
	return a.DeleteCatalog(ctx, DeleteCatalogRequest{
		Name: name,
	})
}

func (a *CatalogsAPI) GetCatalog(ctx context.Context, request GetCatalogRequest) (*GetCatalogResponse, error) {
	var getCatalogResponse GetCatalogResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getCatalogResponse)
	return &getCatalogResponse, err
}

func (a *CatalogsAPI) GetCatalogByName(ctx context.Context, name string) (*GetCatalogResponse, error) {
	return a.GetCatalog(ctx, GetCatalogRequest{
		Name: name,
	})
}

func (a *CatalogsAPI) List(ctx context.Context) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Get(ctx, path, nil, &listCatalogsResponse)
	return &listCatalogsResponse, err
}

func (a *CatalogsAPI) Update(ctx context.Context, request UpdateCatalog) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewExternalLocations(client *client.DatabricksClient) ExternalLocationsService {
	return &ExternalLocationsAPI{
		client: client,
	}
}

type ExternalLocationsAPI struct {
	client *client.DatabricksClient
}

func (a *ExternalLocationsAPI) Create(ctx context.Context, request CreateExternalLocation) (*CreateExternalLocationResponse, error) {
	var createExternalLocationResponse CreateExternalLocationResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Post(ctx, path, request, &createExternalLocationResponse)
	return &createExternalLocationResponse, err
}

func (a *ExternalLocationsAPI) DeleteExternalLocation(ctx context.Context, request DeleteExternalLocation) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *ExternalLocationsAPI) DeleteExternalLocationByName(ctx context.Context, name string) error {
	return a.DeleteExternalLocation(ctx, DeleteExternalLocation{
		Name: name,
	})
}

func (a *ExternalLocationsAPI) GetExternalLocation(ctx context.Context, request GetExternalLocationRequest) (*GetExternalLocationResponse, error) {
	var getExternalLocationResponse GetExternalLocationResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getExternalLocationResponse)
	return &getExternalLocationResponse, err
}

func (a *ExternalLocationsAPI) GetExternalLocationByName(ctx context.Context, name string) (*GetExternalLocationResponse, error) {
	return a.GetExternalLocation(ctx, GetExternalLocationRequest{
		Name: name,
	})
}

func (a *ExternalLocationsAPI) List(ctx context.Context) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Get(ctx, path, nil, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}

func (a *ExternalLocationsAPI) Update(ctx context.Context, request UpdateExternalLocation) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewGrants(client *client.DatabricksClient) GrantsService {
	return &GrantsAPI{
		client: client,
	}
}

type GrantsAPI struct {
	client *client.DatabricksClient
}

func (a *GrantsAPI) UpdatePermissions(ctx context.Context, request UpdatePermissions) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

// <needs contents>
func (a *GrantsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.FullName)
	err := a.client.Get(ctx, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

// <needs contents>
func (a *GrantsAPI) GetPermissionsBySecurableTypeAndFullName(ctx context.Context, securableType string, fullName string) (*GetPermissionsResponse, error) {
	return a.GetPermissions(ctx, GetPermissionsRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

func NewMetastores(client *client.DatabricksClient) MetastoresService {
	return &MetastoresAPI{
		client: client,
	}
}

type MetastoresAPI struct {
	client *client.DatabricksClient
}

func (a *MetastoresAPI) Create(ctx context.Context, request CreateMetastore) (*CreateMetastoreResponse, error) {
	var createMetastoreResponse CreateMetastoreResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Post(ctx, path, request, &createMetastoreResponse)
	return &createMetastoreResponse, err
}

func (a *MetastoresAPI) CreateMetastoreAssignment(ctx context.Context, request CreateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *MetastoresAPI) DeleteMetastore(ctx context.Context, request DeleteMetastore) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *MetastoresAPI) DeleteMetastoreById(ctx context.Context, id string) error {
	return a.DeleteMetastore(ctx, DeleteMetastore{
		Id: id,
	})
}

func (a *MetastoresAPI) DeleteMetastoreAssignment(ctx context.Context, request DeleteMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *MetastoresAPI) DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error {
	return a.DeleteMetastoreAssignment(ctx, DeleteMetastoreAssignment{
		WorkspaceId: workspaceId,
	})
}

func (a *MetastoresAPI) GetMetastore(ctx context.Context, request GetMetastoreRequest) (*GetMetastoreResponse, error) {
	var getMetastoreResponse GetMetastoreResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Get(ctx, path, request, &getMetastoreResponse)
	return &getMetastoreResponse, err
}

func (a *MetastoresAPI) GetMetastoreById(ctx context.Context, id string) (*GetMetastoreResponse, error) {
	return a.GetMetastore(ctx, GetMetastoreRequest{
		Id: id,
	})
}

func (a *MetastoresAPI) GetMetastoreSummary(ctx context.Context) (*GetMetastoreSummaryResponse, error) {
	var getMetastoreSummaryResponse GetMetastoreSummaryResponse
	path := "/api/2.1/unity-catalog/metastore_summary"
	err := a.client.Get(ctx, path, nil, &getMetastoreSummaryResponse)
	return &getMetastoreSummaryResponse, err
}

func (a *MetastoresAPI) List(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Get(ctx, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}

func (a *MetastoresAPI) Update(ctx context.Context, request UpdateMetastore) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}

func (a *MetastoresAPI) UpdateMetastoreAssignment(ctx context.Context, request UpdateMetastoreAssignment) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewProviders(client *client.DatabricksClient) ProvidersService {
	return &ProvidersAPI{
		client: client,
	}
}

type ProvidersAPI struct {
	client *client.DatabricksClient
}

// <needs contents>
func (a *ProvidersAPI) Create(ctx context.Context, request CreateProvider) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Post(ctx, path, request, &createProviderResponse)
	return &createProviderResponse, err
}

// <needs content>
func (a *ProvidersAPI) DeleteProvider(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// <needs content>
func (a *ProvidersAPI) DeleteProviderByName(ctx context.Context, name string) error {
	return a.DeleteProvider(ctx, DeleteProviderRequest{
		Name: name,
	})
}

// <needs contents>
func (a *ProvidersAPI) GetProvider(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getProviderResponse)
	return &getProviderResponse, err
}

// <needs contents>
func (a *ProvidersAPI) GetProviderByName(ctx context.Context, name string) (*GetProviderResponse, error) {
	return a.GetProvider(ctx, GetProviderRequest{
		Name: name,
	})
}

// <needs contents>
func (a *ProvidersAPI) List(ctx context.Context) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Get(ctx, path, nil, &listProvidersResponse)
	return &listProvidersResponse, err
}

func (a *ProvidersAPI) ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.Name)
	err := a.client.Get(ctx, path, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}

func (a *ProvidersAPI) ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error) {
	return a.ListShares(ctx, ListSharesRequest{
		Name: name,
	})
}

// <needs content>
func (a *ProvidersAPI) Update(ctx context.Context, request UpdateProvider) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewRecipientActivation(client *client.DatabricksClient) RecipientActivationService {
	return &RecipientActivationAPI{
		client: client,
	}
}

type RecipientActivationAPI struct {
	client *client.DatabricksClient
}

// <needs content>
func (a *RecipientActivationAPI) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

// <needs content>
func (a *RecipientActivationAPI) GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error {
	return a.GetActivationUrlInfo(ctx, GetActivationUrlInfoRequest{
		ActivationUrl: activationUrl,
	})
}

// Rpc to retrieve access token with an activation token. This is an public API
// without any authentication.
func (a *RecipientActivationAPI) RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error) {
	var retrieveTokenResponse RetrieveTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, &retrieveTokenResponse)
	return &retrieveTokenResponse, err
}

// Rpc to retrieve access token with an activation token. This is an public API
// without any authentication.
func (a *RecipientActivationAPI) RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error) {
	return a.RetrieveToken(ctx, RetrieveTokenRequest{
		ActivationUrl: activationUrl,
	})
}

func NewRecipients(client *client.DatabricksClient) RecipientsService {
	return &RecipientsAPI{
		client: client,
	}
}

type RecipientsAPI struct {
	client *client.DatabricksClient
}

// <needs content>
func (a *RecipientsAPI) Create(ctx context.Context, request CreateRecipient) (*CreateRecipientResponse, error) {
	var createRecipientResponse CreateRecipientResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Post(ctx, path, request, &createRecipientResponse)
	return &createRecipientResponse, err
}

// <needs content>
func (a *RecipientsAPI) DeleteRecipient(ctx context.Context, request DeleteRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// <needs content>
func (a *RecipientsAPI) DeleteRecipientByName(ctx context.Context, name string) error {
	return a.DeleteRecipient(ctx, DeleteRecipientRequest{
		Name: name,
	})
}

func (a *RecipientsAPI) GetRecipient(ctx context.Context, request GetRecipientRequest) (*GetRecipientResponse, error) {
	var getRecipientResponse GetRecipientResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientResponse)
	return &getRecipientResponse, err
}

func (a *RecipientsAPI) GetRecipientByName(ctx context.Context, name string) (*GetRecipientResponse, error) {
	return a.GetRecipient(ctx, GetRecipientRequest{
		Name: name,
	})
}

// <needs content>
func (a *RecipientsAPI) GetRecipientSharePermissions(ctx context.Context, request GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}

// <needs content>
func (a *RecipientsAPI) GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.GetRecipientSharePermissions(ctx, GetRecipientSharePermissionsRequest{
		Name: name,
	})
}

func (a *RecipientsAPI) List(ctx context.Context) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Get(ctx, path, nil, &listRecipientsResponse)
	return &listRecipientsResponse, err
}

// <needs content>
func (a *RecipientsAPI) RotateRecipientToken(ctx context.Context, request RotateRecipientToken) (*RotateRecipientTokenResponse, error) {
	var rotateRecipientTokenResponse RotateRecipientTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	err := a.client.Post(ctx, path, request, &rotateRecipientTokenResponse)
	return &rotateRecipientTokenResponse, err
}

// <needs content>
func (a *RecipientsAPI) Update(ctx context.Context, request UpdateRecipient) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewSchemas(client *client.DatabricksClient) SchemasService {
	return &SchemasAPI{
		client: client,
	}
}

type SchemasAPI struct {
	client *client.DatabricksClient
}

// <needs content>
func (a *SchemasAPI) Create(ctx context.Context, request CreateSchema) (*CreateSchemaResponse, error) {
	var createSchemaResponse CreateSchemaResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Post(ctx, path, request, &createSchemaResponse)
	return &createSchemaResponse, err
}

// <needs content>
func (a *SchemasAPI) DeleteSchema(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Delete(ctx, path, request)
	return err
}

// <needs content>
func (a *SchemasAPI) DeleteSchemaByFullName(ctx context.Context, fullName string) error {
	return a.DeleteSchema(ctx, DeleteSchemaRequest{
		FullName: fullName,
	})
}

// <needs content>
func (a *SchemasAPI) GetSchema(ctx context.Context, request GetSchemaRequest) (*GetSchemaResponse, error) {
	var getSchemaResponse GetSchemaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Get(ctx, path, request, &getSchemaResponse)
	return &getSchemaResponse, err
}

// <needs content>
func (a *SchemasAPI) GetSchemaByFullName(ctx context.Context, fullName string) (*GetSchemaResponse, error) {
	return a.GetSchema(ctx, GetSchemaRequest{
		FullName: fullName,
	})
}

// <needs content>
func (a *SchemasAPI) List(ctx context.Context, request ListRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Get(ctx, path, request, &listSchemasResponse)
	return &listSchemasResponse, err
}

// <needs content>
func (a *SchemasAPI) Update(ctx context.Context, request UpdateSchema) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewShares(client *client.DatabricksClient) SharesService {
	return &SharesAPI{
		client: client,
	}
}

type SharesAPI struct {
	client *client.DatabricksClient
}

// <needs content>
func (a *SharesAPI) Create(ctx context.Context, request CreateShare) (*CreateShareResponse, error) {
	var createShareResponse CreateShareResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Post(ctx, path, request, &createShareResponse)
	return &createShareResponse, err
}

// <needs content>
func (a *SharesAPI) DeleteShare(ctx context.Context, request DeleteShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// <needs content>
func (a *SharesAPI) DeleteShareByName(ctx context.Context, name string) error {
	return a.DeleteShare(ctx, DeleteShareRequest{
		Name: name,
	})
}

func (a *SharesAPI) GetShare(ctx context.Context, request GetShareRequest) (*GetShareResponse, error) {
	var getShareResponse GetShareResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getShareResponse)
	return &getShareResponse, err
}

func (a *SharesAPI) GetShareByName(ctx context.Context, name string) (*GetShareResponse, error) {
	return a.GetShare(ctx, GetShareRequest{
		Name: name,
	})
}

// Permissions API for data sharing.
func (a *SharesAPI) GetSharePermissions(ctx context.Context, request GetSharePermissionsRequest) (*GetSharePermissionsResponse, error) {
	var getSharePermissionsResponse GetSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getSharePermissionsResponse)
	return &getSharePermissionsResponse, err
}

// Permissions API for data sharing.
func (a *SharesAPI) GetSharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error) {
	return a.GetSharePermissions(ctx, GetSharePermissionsRequest{
		Name: name,
	})
}

// <needs content>
func (a *SharesAPI) List(ctx context.Context) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Get(ctx, path, nil, &listSharesResponse)
	return &listSharesResponse, err
}

// <needs content>
func (a *SharesAPI) Update(ctx context.Context, request UpdateShare) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

// <needs content>
func (a *SharesAPI) UpdateSharePermissions(ctx context.Context, request UpdateSharePermissions) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v/permissions", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewStorageCredentials(client *client.DatabricksClient) StorageCredentialsService {
	return &StorageCredentialsAPI{
		client: client,
	}
}

type StorageCredentialsAPI struct {
	client *client.DatabricksClient
}

// <needs contents>
func (a *StorageCredentialsAPI) Create(ctx context.Context, request CreateStorageCredential) (*CreateStorageCredentialResponse, error) {
	var createStorageCredentialResponse CreateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Post(ctx, path, request, &createStorageCredentialResponse)
	return &createStorageCredentialResponse, err
}

// <needs content>
func (a *StorageCredentialsAPI) DeleteStorageCredential(ctx context.Context, request DeleteStorageCredential) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}

// <needs content>
func (a *StorageCredentialsAPI) DeleteStorageCredentialByName(ctx context.Context, name string) error {
	return a.DeleteStorageCredential(ctx, DeleteStorageCredential{
		Name: name,
	})
}

// <needs content>
func (a *StorageCredentialsAPI) GetStorageCredentials(ctx context.Context, request GetStorageCredentialsRequest) (*GetStorageCredentialResponse, error) {
	var getStorageCredentialResponse GetStorageCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Get(ctx, path, request, &getStorageCredentialResponse)
	return &getStorageCredentialResponse, err
}

// <needs content>
func (a *StorageCredentialsAPI) GetStorageCredentialsByName(ctx context.Context, name string) (*GetStorageCredentialResponse, error) {
	return a.GetStorageCredentials(ctx, GetStorageCredentialsRequest{
		Name: name,
	})
}

// <needs content>
func (a *StorageCredentialsAPI) List(ctx context.Context) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Get(ctx, path, nil, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}

func (a *StorageCredentialsAPI) Update(ctx context.Context, request UpdateStorageCredential) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewTables(client *client.DatabricksClient) TablesService {
	return &TablesAPI{
		client: client,
	}
}

type TablesAPI struct {
	client *client.DatabricksClient
}

// <needs content>
func (a *TablesAPI) Create(ctx context.Context, request CreateTable) (*CreateTableResponse, error) {
	var createTableResponse CreateTableResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Post(ctx, path, request, &createTableResponse)
	return &createTableResponse, err
}

// <needs contents>
func (a *TablesAPI) CreateStagingTable(ctx context.Context, request CreateStagingTable) (*CreateStagingTableResponse, error) {
	var createStagingTableResponse CreateStagingTableResponse
	path := "/api/2.1/unity-catalog/staging-tables"
	err := a.client.Post(ctx, path, request, &createStagingTableResponse)
	return &createStagingTableResponse, err
}

// <needs content>
func (a *TablesAPI) DeleteTable(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Delete(ctx, path, request)
	return err
}

// <needs content>
func (a *TablesAPI) DeleteTableByFullName(ctx context.Context, fullName string) error {
	return a.DeleteTable(ctx, DeleteTableRequest{
		FullName: fullName,
	})
}

// <needs content>
func (a *TablesAPI) GetTable(ctx context.Context, request GetTableRequest) (*GetTableResponse, error) {
	var getTableResponse GetTableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Get(ctx, path, request, &getTableResponse)
	return &getTableResponse, err
}

// <needs content>
func (a *TablesAPI) GetTableByFullName(ctx context.Context, fullName string) (*GetTableResponse, error) {
	return a.GetTable(ctx, GetTableRequest{
		FullName: fullName,
	})
}

// <needs content>
func (a *TablesAPI) List(ctx context.Context, request ListRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Get(ctx, path, request, &listTablesResponse)
	return &listTablesResponse, err
}

func (a *TablesAPI) ListTableSummaries(ctx context.Context, request ListTableSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	err := a.client.Get(ctx, path, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}

// <needs content>
func (a *TablesAPI) Update(ctx context.Context, request UpdateTable) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullName)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewUnityFiles(client *client.DatabricksClient) UnityFilesService {
	return &UnityFilesAPI{
		client: client,
	}
}

type UnityFilesAPI struct {
	client *client.DatabricksClient
}

func (a *UnityFilesAPI) ListFiles(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error) {
	var listFilesResponse ListFilesResponse
	path := "/api/2.1/unity-catalog/files"
	err := a.client.Get(ctx, path, request, &listFilesResponse)
	return &listFilesResponse, err
}
