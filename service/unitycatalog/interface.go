// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
	
)



type CatalogsService interface {
    
    CreateCatalog(ctx context.Context, createCatalogRequest CreateCatalogRequest) (*CreateCatalogResponse, error)
    
    DeleteCatalog(ctx context.Context, deleteCatalogRequest DeleteCatalogRequest) error
	DeleteCatalogByNameArg(ctx context.Context, nameArg string) error
    
    GetCatalog(ctx context.Context, getCatalogRequest GetCatalogRequest) (*GetCatalogResponse, error)
	GetCatalogByNameArg(ctx context.Context, nameArg string) (*GetCatalogResponse, error)
    
    ListCatalogs(ctx context.Context) (*ListCatalogsResponse, error)
    
    UpdateCatalog(ctx context.Context, updateCatalogRequest UpdateCatalogRequest) error
}


type ExternalLocationsService interface {
    
    CreateExternalLocation(ctx context.Context, createExternalLocationRequest CreateExternalLocationRequest) (*CreateExternalLocationResponse, error)
    
    DeleteExternalLocation(ctx context.Context, deleteExternalLocationRequest DeleteExternalLocationRequest) error
	DeleteExternalLocationByNameArg(ctx context.Context, nameArg string) error
    
    GetExternalLocation(ctx context.Context, getExternalLocationRequest GetExternalLocationRequest) (*GetExternalLocationResponse, error)
	GetExternalLocationByNameArg(ctx context.Context, nameArg string) (*GetExternalLocationResponse, error)
    
    ListExternalLocations(ctx context.Context) (*ListExternalLocationsResponse, error)
    
    UpdateExternalLocation(ctx context.Context, updateExternalLocationRequest UpdateExternalLocationRequest) error
}


type GrantsService interface {
    
    GetPermissions(ctx context.Context, getPermissionsRequest GetPermissionsRequest) (*GetPermissionsResponse, error)
	GetPermissionsBySecurableTypeAndSecurableFullName(ctx context.Context, securableType string, securableFullName string) (*GetPermissionsResponse, error)
    
    UpdatePermissions(ctx context.Context, updatePermissionsRequest UpdatePermissionsRequest) error
}


type MetastoresService interface {
    
    CreateMetastore(ctx context.Context, createMetastoreRequest CreateMetastoreRequest) (*CreateMetastoreResponse, error)
    
    CreateMetastoreAssignment(ctx context.Context, createMetastoreAssignmentRequest CreateMetastoreAssignmentRequest) error
    
    DeleteMetastore(ctx context.Context, deleteMetastoreRequest DeleteMetastoreRequest) error
	DeleteMetastoreById(ctx context.Context, id string) error
    
    DeleteMetastoreAssignment(ctx context.Context, deleteMetastoreAssignmentRequest DeleteMetastoreAssignmentRequest) error
	DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error
    
    GetMetastore(ctx context.Context, getMetastoreRequest GetMetastoreRequest) (*GetMetastoreResponse, error)
	GetMetastoreById(ctx context.Context, id string) (*GetMetastoreResponse, error)
    
    GetMetastoreSummary(ctx context.Context) (*GetMetastoreSummaryResponse, error)
    
    ListMetastores(ctx context.Context) (*ListMetastoresResponse, error)
    
    UpdateMetastore(ctx context.Context, updateMetastoreRequest UpdateMetastoreRequest) error
    
    UpdateMetastoreAssignment(ctx context.Context, updateMetastoreAssignmentRequest UpdateMetastoreAssignmentRequest) error
}


type ProvidersService interface {
    
    CreateProvider(ctx context.Context, createProviderRequest CreateProviderRequest) (*CreateProviderResponse, error)
    
    DeleteProvider(ctx context.Context, deleteProviderRequest DeleteProviderRequest) error
	DeleteProviderByNameArg(ctx context.Context, nameArg string) error
    
    GetProvider(ctx context.Context, getProviderRequest GetProviderRequest) (*GetProviderResponse, error)
	GetProviderByNameArg(ctx context.Context, nameArg string) (*GetProviderResponse, error)
    
    ListProviderShares(ctx context.Context, listProviderSharesRequest ListProviderSharesRequest) (*ListProviderSharesResponse, error)
	ListProviderSharesByProviderNameArg(ctx context.Context, providerNameArg string) (*ListProviderSharesResponse, error)
    
    ListProviders(ctx context.Context) (*ListProvidersResponse, error)
    
    UpdateProvider(ctx context.Context, updateProviderRequest UpdateProviderRequest) error
}


type RecipientActivationService interface {
    
    GetActivationUrlInfo(ctx context.Context, getActivationUrlInfoRequest GetActivationUrlInfoRequest) error
	GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error
    // Rpc to retrieve access token with an activation token. This is an public
    // API without any authentication.
    RetrieveToken(ctx context.Context, retrieveTokenRequest RetrieveTokenRequest) (*RetrieveTokenResponse, error)
	RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error)
}


type RecipientsService interface {
    
    CreateRecipient(ctx context.Context, createRecipientRequest CreateRecipientRequest) (*CreateRecipientResponse, error)
    
    DeleteRecipient(ctx context.Context, deleteRecipientRequest DeleteRecipientRequest) error
	DeleteRecipientByName(ctx context.Context, name string) error
    
    GetRecipient(ctx context.Context, getRecipientRequest GetRecipientRequest) (*GetRecipientResponse, error)
	GetRecipientByName(ctx context.Context, name string) (*GetRecipientResponse, error)
    
    GetRecipientSharePermissions(ctx context.Context, getRecipientSharePermissionsRequest GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error)
	GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error)
    
    ListRecipients(ctx context.Context) (*ListRecipientsResponse, error)
    
    RotateRecipientToken(ctx context.Context, rotateRecipientTokenRequest RotateRecipientTokenRequest) (*RotateRecipientTokenResponse, error)
    
    UpdateRecipient(ctx context.Context, updateRecipientRequest UpdateRecipientRequest) error
}


type SchemasService interface {
    
    CreateSchema(ctx context.Context, createSchemaRequest CreateSchemaRequest) (*CreateSchemaResponse, error)
    
    DeleteSchema(ctx context.Context, deleteSchemaRequest DeleteSchemaRequest) error
	DeleteSchemaByFullNameArg(ctx context.Context, fullNameArg string) error
    
    GetSchema(ctx context.Context, getSchemaRequest GetSchemaRequest) (*GetSchemaResponse, error)
	GetSchemaByFullNameArg(ctx context.Context, fullNameArg string) (*GetSchemaResponse, error)
    
    ListSchemas(ctx context.Context, listSchemasRequest ListSchemasRequest) (*ListSchemasResponse, error)
    
    UpdateSchema(ctx context.Context, updateSchemaRequest UpdateSchemaRequest) error
}


type SharesService interface {
    
    CreateShare(ctx context.Context, createShareRequest CreateShareRequest) (*CreateShareResponse, error)
    
    DeleteShare(ctx context.Context, deleteShareRequest DeleteShareRequest) error
	DeleteShareByName(ctx context.Context, name string) error
    
    GetShare(ctx context.Context, getShareRequest GetShareRequest) (*GetShareResponse, error)
	GetShareByName(ctx context.Context, name string) (*GetShareResponse, error)
    // Permissions API for data sharing.
    GetSharePermissions(ctx context.Context, getSharePermissionsRequest GetSharePermissionsRequest) (*GetSharePermissionsResponse, error)
	GetSharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error)
    
    ListShares(ctx context.Context) (*ListSharesResponse, error)
    
    UpdateShare(ctx context.Context, updateShareRequest UpdateShareRequest) error
    
    UpdateSharePermissions(ctx context.Context, updateSharePermissionsRequest UpdateSharePermissionsRequest) error
}


type StorageCredentialsService interface {
    
    CreateStorageCredential(ctx context.Context, createStorageCredentialRequest CreateStorageCredentialRequest) (*CreateStorageCredentialResponse, error)
    
    DeleteStorageCredential(ctx context.Context, deleteStorageCredentialRequest DeleteStorageCredentialRequest) error
	DeleteStorageCredentialByNameArg(ctx context.Context, nameArg string) error
    
    GetStorageCredential(ctx context.Context, getStorageCredentialRequest GetStorageCredentialRequest) (*GetStorageCredentialResponse, error)
	GetStorageCredentialByNameArg(ctx context.Context, nameArg string) (*GetStorageCredentialResponse, error)
    
    ListStorageCredentials(ctx context.Context) (*ListStorageCredentialsResponse, error)
    
    UpdateStorageCredential(ctx context.Context, updateStorageCredentialRequest UpdateStorageCredentialRequest) error
}


type TablesService interface {
    
    CreateStagingTable(ctx context.Context, createStagingTableRequest CreateStagingTableRequest) (*CreateStagingTableResponse, error)
    
    CreateTable(ctx context.Context, createTableRequest CreateTableRequest) (*CreateTableResponse, error)
    
    DeleteTable(ctx context.Context, deleteTableRequest DeleteTableRequest) error
	DeleteTableByFullNameArg(ctx context.Context, fullNameArg string) error
    
    GetTable(ctx context.Context, getTableRequest GetTableRequest) (*GetTableResponse, error)
	GetTableByFullNameArg(ctx context.Context, fullNameArg string) (*GetTableResponse, error)
    
    ListTableSummaries(ctx context.Context, listTableSummariesRequest ListTableSummariesRequest) (*ListTableSummariesResponse, error)
    
    ListTables(ctx context.Context, listTablesRequest ListTablesRequest) (*ListTablesResponse, error)
    
    UpdateTable(ctx context.Context, updateTableRequest UpdateTableRequest) error
}


type UnityFilesService interface {
    
    ListFiles(ctx context.Context, listFilesRequest ListFilesRequest) (*ListFilesResponse, error)
}
