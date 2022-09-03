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


func (a *CatalogsAPI) CreateCatalog(ctx context.Context, request CreateCatalogRequest) (*CreateCatalogResponse, error) {
	var createCatalogResponse CreateCatalogResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Post(ctx, path, request, &createCatalogResponse)
	return &createCatalogResponse, err
}


func (a *CatalogsAPI) DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *CatalogsAPI) DeleteCatalogByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteCatalog(ctx, DeleteCatalogRequest{
		NameArg: nameArg,
	})
}


func (a *CatalogsAPI) GetCatalog(ctx context.Context, request GetCatalogRequest) (*GetCatalogResponse, error) {
	var getCatalogResponse GetCatalogResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getCatalogResponse)
	return &getCatalogResponse, err
}


func (a *CatalogsAPI) GetCatalogByNameArg(ctx context.Context, nameArg string) (*GetCatalogResponse, error) {
	return a.GetCatalog(ctx, GetCatalogRequest{
		NameArg: nameArg,
	})
}


func (a *CatalogsAPI) ListCatalogs(ctx context.Context) (*ListCatalogsResponse, error) {
	var listCatalogsResponse ListCatalogsResponse
	path := "/api/2.1/unity-catalog/catalogs"
	err := a.client.Get(ctx, path, nil, &listCatalogsResponse)
	return &listCatalogsResponse, err
}


func (a *CatalogsAPI) UpdateCatalog(ctx context.Context, request UpdateCatalogRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/catalogs/%v", request.NameArg)
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


func (a *ExternalLocationsAPI) CreateExternalLocation(ctx context.Context, request CreateExternalLocationRequest) (*CreateExternalLocationResponse, error) {
	var createExternalLocationResponse CreateExternalLocationResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Post(ctx, path, request, &createExternalLocationResponse)
	return &createExternalLocationResponse, err
}


func (a *ExternalLocationsAPI) DeleteExternalLocation(ctx context.Context, request DeleteExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *ExternalLocationsAPI) DeleteExternalLocationByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteExternalLocation(ctx, DeleteExternalLocationRequest{
		NameArg: nameArg,
	})
}


func (a *ExternalLocationsAPI) GetExternalLocation(ctx context.Context, request GetExternalLocationRequest) (*GetExternalLocationResponse, error) {
	var getExternalLocationResponse GetExternalLocationResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getExternalLocationResponse)
	return &getExternalLocationResponse, err
}


func (a *ExternalLocationsAPI) GetExternalLocationByNameArg(ctx context.Context, nameArg string) (*GetExternalLocationResponse, error) {
	return a.GetExternalLocation(ctx, GetExternalLocationRequest{
		NameArg: nameArg,
	})
}


func (a *ExternalLocationsAPI) ListExternalLocations(ctx context.Context) (*ListExternalLocationsResponse, error) {
	var listExternalLocationsResponse ListExternalLocationsResponse
	path := "/api/2.1/unity-catalog/external-locations"
	err := a.client.Get(ctx, path, nil, &listExternalLocationsResponse)
	return &listExternalLocationsResponse, err
}


func (a *ExternalLocationsAPI) UpdateExternalLocation(ctx context.Context, request UpdateExternalLocationRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/external-locations/%v", request.NameArg)
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


func (a *GrantsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.SecurableFullName)
	err := a.client.Get(ctx, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}


func (a *GrantsAPI) GetPermissionsBySecurableTypeAndSecurableFullName(ctx context.Context, securableType string, securableFullName string) (*GetPermissionsResponse, error) {
	return a.GetPermissions(ctx, GetPermissionsRequest{
		SecurableType: securableType,
		SecurableFullName: securableFullName,
	})
}


func (a *GrantsAPI) UpdatePermissions(ctx context.Context, request UpdatePermissionsRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/permissions/%v/%v", request.SecurableType, request.SecurableFullName)
	err := a.client.Patch(ctx, path, request)
	return err
}


func NewMetastores(client *client.DatabricksClient) MetastoresService {
	return &MetastoresAPI{
		client: client,
	}
}

type MetastoresAPI struct {
	client *client.DatabricksClient
}


func (a *MetastoresAPI) CreateMetastore(ctx context.Context, request CreateMetastoreRequest) (*CreateMetastoreResponse, error) {
	var createMetastoreResponse CreateMetastoreResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Post(ctx, path, request, &createMetastoreResponse)
	return &createMetastoreResponse, err
}


func (a *MetastoresAPI) CreateMetastoreAssignment(ctx context.Context, request CreateMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *MetastoresAPI) DeleteMetastore(ctx context.Context, request DeleteMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *MetastoresAPI) DeleteMetastoreById(ctx context.Context, id string) error {
	return a.DeleteMetastore(ctx, DeleteMetastoreRequest{
		Id: id,
	})
}


func (a *MetastoresAPI) DeleteMetastoreAssignment(ctx context.Context, request DeleteMetastoreAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/workspaces/%v/metastore", request.WorkspaceId)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *MetastoresAPI) DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error {
	return a.DeleteMetastoreAssignment(ctx, DeleteMetastoreAssignmentRequest{
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


func (a *MetastoresAPI) ListMetastores(ctx context.Context) (*ListMetastoresResponse, error) {
	var listMetastoresResponse ListMetastoresResponse
	path := "/api/2.1/unity-catalog/metastores"
	err := a.client.Get(ctx, path, nil, &listMetastoresResponse)
	return &listMetastoresResponse, err
}


func (a *MetastoresAPI) UpdateMetastore(ctx context.Context, request UpdateMetastoreRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/metastores/%v", request.Id)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *MetastoresAPI) UpdateMetastoreAssignment(ctx context.Context, request UpdateMetastoreAssignmentRequest) error {
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


func (a *ProvidersAPI) CreateProvider(ctx context.Context, request CreateProviderRequest) (*CreateProviderResponse, error) {
	var createProviderResponse CreateProviderResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Post(ctx, path, request, &createProviderResponse)
	return &createProviderResponse, err
}


func (a *ProvidersAPI) DeleteProvider(ctx context.Context, request DeleteProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *ProvidersAPI) DeleteProviderByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteProvider(ctx, DeleteProviderRequest{
		NameArg: nameArg,
	})
}


func (a *ProvidersAPI) GetProvider(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error) {
	var getProviderResponse GetProviderResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getProviderResponse)
	return &getProviderResponse, err
}


func (a *ProvidersAPI) GetProviderByNameArg(ctx context.Context, nameArg string) (*GetProviderResponse, error) {
	return a.GetProvider(ctx, GetProviderRequest{
		NameArg: nameArg,
	})
}


func (a *ProvidersAPI) ListProviderShares(ctx context.Context, request ListProviderSharesRequest) (*ListProviderSharesResponse, error) {
	var listProviderSharesResponse ListProviderSharesResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v/shares", request.ProviderNameArg)
	err := a.client.Get(ctx, path, request, &listProviderSharesResponse)
	return &listProviderSharesResponse, err
}


func (a *ProvidersAPI) ListProviderSharesByProviderNameArg(ctx context.Context, providerNameArg string) (*ListProviderSharesResponse, error) {
	return a.ListProviderShares(ctx, ListProviderSharesRequest{
		ProviderNameArg: providerNameArg,
	})
}


func (a *ProvidersAPI) ListProviders(ctx context.Context) (*ListProvidersResponse, error) {
	var listProvidersResponse ListProvidersResponse
	path := "/api/2.1/unity-catalog/providers"
	err := a.client.Get(ctx, path, nil, &listProvidersResponse)
	return &listProvidersResponse, err
}


func (a *ProvidersAPI) UpdateProvider(ctx context.Context, request UpdateProviderRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/providers/%v", request.NameArg)
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


func (a *RecipientActivationAPI) GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/public/data_sharing_activation_info/%v", request.ActivationUrl)
	err := a.client.Get(ctx, path, request, nil)
	return err
}


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


func (a *RecipientsAPI) CreateRecipient(ctx context.Context, request CreateRecipientRequest) (*CreateRecipientResponse, error) {
	var createRecipientResponse CreateRecipientResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Post(ctx, path, request, &createRecipientResponse)
	return &createRecipientResponse, err
}


func (a *RecipientsAPI) DeleteRecipient(ctx context.Context, request DeleteRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}


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


func (a *RecipientsAPI) GetRecipientSharePermissions(ctx context.Context, request GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error) {
	var getRecipientSharePermissionsResponse GetRecipientSharePermissionsResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/share-permissions", request.Name)
	err := a.client.Get(ctx, path, request, &getRecipientSharePermissionsResponse)
	return &getRecipientSharePermissionsResponse, err
}


func (a *RecipientsAPI) GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error) {
	return a.GetRecipientSharePermissions(ctx, GetRecipientSharePermissionsRequest{
		Name: name,
	})
}


func (a *RecipientsAPI) ListRecipients(ctx context.Context) (*ListRecipientsResponse, error) {
	var listRecipientsResponse ListRecipientsResponse
	path := "/api/2.1/unity-catalog/recipients"
	err := a.client.Get(ctx, path, nil, &listRecipientsResponse)
	return &listRecipientsResponse, err
}


func (a *RecipientsAPI) RotateRecipientToken(ctx context.Context, request RotateRecipientTokenRequest) (*RotateRecipientTokenResponse, error) {
	var rotateRecipientTokenResponse RotateRecipientTokenResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v/rotate-token", request.Name)
	err := a.client.Post(ctx, path, request, &rotateRecipientTokenResponse)
	return &rotateRecipientTokenResponse, err
}


func (a *RecipientsAPI) UpdateRecipient(ctx context.Context, request UpdateRecipientRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/recipients/%v", request.NameArg)
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


func (a *SchemasAPI) CreateSchema(ctx context.Context, request CreateSchemaRequest) (*CreateSchemaResponse, error) {
	var createSchemaResponse CreateSchemaResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Post(ctx, path, request, &createSchemaResponse)
	return &createSchemaResponse, err
}


func (a *SchemasAPI) DeleteSchema(ctx context.Context, request DeleteSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullNameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *SchemasAPI) DeleteSchemaByFullNameArg(ctx context.Context, fullNameArg string) error {
	return a.DeleteSchema(ctx, DeleteSchemaRequest{
		FullNameArg: fullNameArg,
	})
}


func (a *SchemasAPI) GetSchema(ctx context.Context, request GetSchemaRequest) (*GetSchemaResponse, error) {
	var getSchemaResponse GetSchemaResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullNameArg)
	err := a.client.Get(ctx, path, request, &getSchemaResponse)
	return &getSchemaResponse, err
}


func (a *SchemasAPI) GetSchemaByFullNameArg(ctx context.Context, fullNameArg string) (*GetSchemaResponse, error) {
	return a.GetSchema(ctx, GetSchemaRequest{
		FullNameArg: fullNameArg,
	})
}


func (a *SchemasAPI) ListSchemas(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error) {
	var listSchemasResponse ListSchemasResponse
	path := "/api/2.1/unity-catalog/schemas"
	err := a.client.Get(ctx, path, request, &listSchemasResponse)
	return &listSchemasResponse, err
}


func (a *SchemasAPI) UpdateSchema(ctx context.Context, request UpdateSchemaRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/schemas/%v", request.FullNameArg)
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


func (a *SharesAPI) CreateShare(ctx context.Context, request CreateShareRequest) (*CreateShareResponse, error) {
	var createShareResponse CreateShareResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Post(ctx, path, request, &createShareResponse)
	return &createShareResponse, err
}


func (a *SharesAPI) DeleteShare(ctx context.Context, request DeleteShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Delete(ctx, path, request)
	return err
}


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


func (a *SharesAPI) ListShares(ctx context.Context) (*ListSharesResponse, error) {
	var listSharesResponse ListSharesResponse
	path := "/api/2.1/unity-catalog/shares"
	err := a.client.Get(ctx, path, nil, &listSharesResponse)
	return &listSharesResponse, err
}


func (a *SharesAPI) UpdateShare(ctx context.Context, request UpdateShareRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/shares/%v", request.Name)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *SharesAPI) UpdateSharePermissions(ctx context.Context, request UpdateSharePermissionsRequest) error {
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


func (a *StorageCredentialsAPI) CreateStorageCredential(ctx context.Context, request CreateStorageCredentialRequest) (*CreateStorageCredentialResponse, error) {
	var createStorageCredentialResponse CreateStorageCredentialResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Post(ctx, path, request, &createStorageCredentialResponse)
	return &createStorageCredentialResponse, err
}


func (a *StorageCredentialsAPI) DeleteStorageCredential(ctx context.Context, request DeleteStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.NameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *StorageCredentialsAPI) DeleteStorageCredentialByNameArg(ctx context.Context, nameArg string) error {
	return a.DeleteStorageCredential(ctx, DeleteStorageCredentialRequest{
		NameArg: nameArg,
	})
}


func (a *StorageCredentialsAPI) GetStorageCredential(ctx context.Context, request GetStorageCredentialRequest) (*GetStorageCredentialResponse, error) {
	var getStorageCredentialResponse GetStorageCredentialResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.NameArg)
	err := a.client.Get(ctx, path, request, &getStorageCredentialResponse)
	return &getStorageCredentialResponse, err
}


func (a *StorageCredentialsAPI) GetStorageCredentialByNameArg(ctx context.Context, nameArg string) (*GetStorageCredentialResponse, error) {
	return a.GetStorageCredential(ctx, GetStorageCredentialRequest{
		NameArg: nameArg,
	})
}


func (a *StorageCredentialsAPI) ListStorageCredentials(ctx context.Context) (*ListStorageCredentialsResponse, error) {
	var listStorageCredentialsResponse ListStorageCredentialsResponse
	path := "/api/2.1/unity-catalog/storage-credentials"
	err := a.client.Get(ctx, path, nil, &listStorageCredentialsResponse)
	return &listStorageCredentialsResponse, err
}


func (a *StorageCredentialsAPI) UpdateStorageCredential(ctx context.Context, request UpdateStorageCredentialRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/storage-credentials/%v", request.NameArg)
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


func (a *TablesAPI) CreateStagingTable(ctx context.Context, request CreateStagingTableRequest) (*CreateStagingTableResponse, error) {
	var createStagingTableResponse CreateStagingTableResponse
	path := "/api/2.1/unity-catalog/staging-tables"
	err := a.client.Post(ctx, path, request, &createStagingTableResponse)
	return &createStagingTableResponse, err
}


func (a *TablesAPI) CreateTable(ctx context.Context, request CreateTableRequest) (*CreateTableResponse, error) {
	var createTableResponse CreateTableResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Post(ctx, path, request, &createTableResponse)
	return &createTableResponse, err
}


func (a *TablesAPI) DeleteTable(ctx context.Context, request DeleteTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullNameArg)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *TablesAPI) DeleteTableByFullNameArg(ctx context.Context, fullNameArg string) error {
	return a.DeleteTable(ctx, DeleteTableRequest{
		FullNameArg: fullNameArg,
	})
}


func (a *TablesAPI) GetTable(ctx context.Context, request GetTableRequest) (*GetTableResponse, error) {
	var getTableResponse GetTableResponse
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullNameArg)
	err := a.client.Get(ctx, path, request, &getTableResponse)
	return &getTableResponse, err
}


func (a *TablesAPI) GetTableByFullNameArg(ctx context.Context, fullNameArg string) (*GetTableResponse, error) {
	return a.GetTable(ctx, GetTableRequest{
		FullNameArg: fullNameArg,
	})
}


func (a *TablesAPI) ListTableSummaries(ctx context.Context, request ListTableSummariesRequest) (*ListTableSummariesResponse, error) {
	var listTableSummariesResponse ListTableSummariesResponse
	path := "/api/2.1/unity-catalog/table-summaries"
	err := a.client.Get(ctx, path, request, &listTableSummariesResponse)
	return &listTableSummariesResponse, err
}


func (a *TablesAPI) ListTables(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error) {
	var listTablesResponse ListTablesResponse
	path := "/api/2.1/unity-catalog/tables"
	err := a.client.Get(ctx, path, request, &listTablesResponse)
	return &listTablesResponse, err
}


func (a *TablesAPI) UpdateTable(ctx context.Context, request UpdateTableRequest) error {
	path := fmt.Sprintf("/api/2.1/unity-catalog/tables/%v", request.FullNameArg)
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

