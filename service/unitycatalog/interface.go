// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type CatalogsService interface {
	Create(ctx context.Context, request CreateCatalog) (*CreateCatalogResponse, error)

	DeleteCatalog(ctx context.Context, request DeleteCatalogRequest) error

	// DeleteCatalogByName calls DeleteCatalog, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteCatalogByName(ctx context.Context, name string) error

	GetCatalog(ctx context.Context, request GetCatalogRequest) (*GetCatalogResponse, error)

	// GetCatalogByName calls GetCatalog, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetCatalogByName(ctx context.Context, name string) (*GetCatalogResponse, error)

	List(ctx context.Context) (*ListCatalogsResponse, error)

	Update(ctx context.Context, request UpdateCatalog) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ExternalLocationsService interface {
	Create(ctx context.Context, request CreateExternalLocation) (*CreateExternalLocationResponse, error)

	DeleteExternalLocation(ctx context.Context, request DeleteExternalLocationRequest) error

	// DeleteExternalLocationByName calls DeleteExternalLocation, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteExternalLocationByName(ctx context.Context, name string) error

	GetExternalLocation(ctx context.Context, request GetExternalLocationRequest) (*GetExternalLocationResponse, error)

	// GetExternalLocationByName calls GetExternalLocation, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetExternalLocationByName(ctx context.Context, name string) (*GetExternalLocationResponse, error)

	List(ctx context.Context) (*ListExternalLocationsResponse, error)

	Update(ctx context.Context, request UpdateExternalLocation) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type GrantsService interface {
	UpdatePermissions(ctx context.Context, request UpdatePermissions) error

	// <needs contents>
	GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error)

	// GetPermissionsBySecurableTypeAndFullName calls GetPermissions, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetPermissionsBySecurableTypeAndFullName(ctx context.Context, securableType string, fullName string) (*GetPermissionsResponse, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type MetastoresService interface {
	Create(ctx context.Context, request CreateMetastore) (*CreateMetastoreResponse, error)

	CreateMetastoreAssignment(ctx context.Context, request CreateMetastoreAssignment) error

	DeleteMetastore(ctx context.Context, request DeleteMetastoreRequest) error

	// DeleteMetastoreById calls DeleteMetastore, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteMetastoreById(ctx context.Context, id string) error

	DeleteMetastoreAssignment(ctx context.Context, request DeleteMetastoreAssignmentRequest) error

	// DeleteMetastoreAssignmentByWorkspaceId calls DeleteMetastoreAssignment, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteMetastoreAssignmentByWorkspaceId(ctx context.Context, workspaceId int) error

	GetMetastore(ctx context.Context, request GetMetastoreRequest) (*GetMetastoreResponse, error)

	// GetMetastoreById calls GetMetastore, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetMetastoreById(ctx context.Context, id string) (*GetMetastoreResponse, error)

	GetMetastoreSummary(ctx context.Context) (*GetMetastoreSummaryResponse, error)

	List(ctx context.Context) (*ListMetastoresResponse, error)

	Update(ctx context.Context, request UpdateMetastore) error

	UpdateMetastoreAssignment(ctx context.Context, request UpdateMetastoreAssignment) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ProvidersService interface {

	// <needs contents>
	Create(ctx context.Context, request CreateProvider) (*CreateProviderResponse, error)

	// <needs content>
	DeleteProvider(ctx context.Context, request DeleteProviderRequest) error

	// DeleteProviderByName calls DeleteProvider, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteProviderByName(ctx context.Context, name string) error

	// <needs contents>
	GetProvider(ctx context.Context, request GetProviderRequest) (*GetProviderResponse, error)

	// GetProviderByName calls GetProvider, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetProviderByName(ctx context.Context, name string) (*GetProviderResponse, error)

	// <needs contents>
	List(ctx context.Context) (*ListProvidersResponse, error)

	ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error)

	// ListSharesByName calls ListShares, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSharesByName(ctx context.Context, name string) (*ListProviderSharesResponse, error)

	// <needs content>
	Update(ctx context.Context, request UpdateProvider) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type RecipientActivationService interface {

	// <needs content>
	GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error

	// GetActivationUrlInfoByActivationUrl calls GetActivationUrlInfo, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetActivationUrlInfoByActivationUrl(ctx context.Context, activationUrl string) error

	// Rpc to retrieve access token with an activation token. This is an public
	// API without any authentication.
	RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error)

	// RetrieveTokenByActivationUrl calls RetrieveToken, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	RetrieveTokenByActivationUrl(ctx context.Context, activationUrl string) (*RetrieveTokenResponse, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type RecipientsService interface {

	// <needs content>
	Create(ctx context.Context, request CreateRecipient) (*CreateRecipientResponse, error)

	// <needs content>
	DeleteRecipient(ctx context.Context, request DeleteRecipientRequest) error

	// DeleteRecipientByName calls DeleteRecipient, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteRecipientByName(ctx context.Context, name string) error

	GetRecipient(ctx context.Context, request GetRecipientRequest) (*GetRecipientResponse, error)

	// GetRecipientByName calls GetRecipient, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetRecipientByName(ctx context.Context, name string) (*GetRecipientResponse, error)

	// <needs content>
	GetRecipientSharePermissions(ctx context.Context, request GetRecipientSharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error)

	// GetRecipientSharePermissionsByName calls GetRecipientSharePermissions, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetRecipientSharePermissionsByName(ctx context.Context, name string) (*GetRecipientSharePermissionsResponse, error)

	List(ctx context.Context) (*ListRecipientsResponse, error)

	// <needs content>
	RotateRecipientToken(ctx context.Context, request RotateRecipientToken) (*RotateRecipientTokenResponse, error)

	// <needs content>
	Update(ctx context.Context, request UpdateRecipient) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type SchemasService interface {

	// <needs content>
	Create(ctx context.Context, request CreateSchema) (*CreateSchemaResponse, error)

	// <needs content>
	DeleteSchema(ctx context.Context, request DeleteSchemaRequest) error

	// DeleteSchemaByFullName calls DeleteSchema, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteSchemaByFullName(ctx context.Context, fullName string) error

	// <needs content>
	GetSchema(ctx context.Context, request GetSchemaRequest) (*GetSchemaResponse, error)

	// GetSchemaByFullName calls GetSchema, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetSchemaByFullName(ctx context.Context, fullName string) (*GetSchemaResponse, error)

	// <needs content>
	List(ctx context.Context, request ListRequest) (*ListSchemasResponse, error)

	// <needs content>
	Update(ctx context.Context, request UpdateSchema) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type SharesService interface {

	// <needs content>
	Create(ctx context.Context, request CreateShare) (*CreateShareResponse, error)

	// <needs content>
	DeleteShare(ctx context.Context, request DeleteShareRequest) error

	// DeleteShareByName calls DeleteShare, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteShareByName(ctx context.Context, name string) error

	GetShare(ctx context.Context, request GetShareRequest) (*GetShareResponse, error)

	// GetShareByName calls GetShare, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetShareByName(ctx context.Context, name string) (*GetShareResponse, error)

	// Permissions API for data sharing.
	GetSharePermissions(ctx context.Context, request GetSharePermissionsRequest) (*GetSharePermissionsResponse, error)

	// GetSharePermissionsByName calls GetSharePermissions, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetSharePermissionsByName(ctx context.Context, name string) (*GetSharePermissionsResponse, error)

	// <needs content>
	List(ctx context.Context) (*ListSharesResponse, error)

	// <needs content>
	Update(ctx context.Context, request UpdateShare) error

	// <needs content>
	UpdateSharePermissions(ctx context.Context, request UpdateSharePermissions) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type StorageCredentialsService interface {

	// <needs contents>
	Create(ctx context.Context, request CreateStorageCredential) (*CreateStorageCredentialResponse, error)

	// <needs content>
	DeleteStorageCredential(ctx context.Context, request DeleteStorageCredentialRequest) error

	// DeleteStorageCredentialByName calls DeleteStorageCredential, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteStorageCredentialByName(ctx context.Context, name string) error

	// <needs content>
	GetStorageCredentials(ctx context.Context, request GetStorageCredentialsRequest) (*GetStorageCredentialResponse, error)

	// GetStorageCredentialsByName calls GetStorageCredentials, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetStorageCredentialsByName(ctx context.Context, name string) (*GetStorageCredentialResponse, error)

	// <needs content>
	List(ctx context.Context) (*ListStorageCredentialsResponse, error)

	Update(ctx context.Context, request UpdateStorageCredential) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type TablesService interface {

	// <needs content>
	Create(ctx context.Context, request CreateTable) (*CreateTableResponse, error)

	// <needs contents>
	CreateStagingTable(ctx context.Context, request CreateStagingTable) (*CreateStagingTableResponse, error)

	// <needs content>
	DeleteTable(ctx context.Context, request DeleteTableRequest) error

	// DeleteTableByFullName calls DeleteTable, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteTableByFullName(ctx context.Context, fullName string) error

	// <needs content>
	GetTable(ctx context.Context, request GetTableRequest) (*GetTableResponse, error)

	// GetTableByFullName calls GetTable, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetTableByFullName(ctx context.Context, fullName string) (*GetTableResponse, error)

	// <needs content>
	List(ctx context.Context, request ListRequest) (*ListTablesResponse, error)

	ListTableSummaries(ctx context.Context, request ListTableSummariesRequest) (*ListTableSummariesResponse, error)

	// <needs content>
	Update(ctx context.Context, request UpdateTable) error
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type UnityFilesService interface {
	ListFiles(ctx context.Context, request ListFilesRequest) (*ListFilesResponse, error)
}
