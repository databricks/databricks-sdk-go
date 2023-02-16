// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package unitycatalog

import (
	"context"
)

// A catalog is the first layer of Unity Catalog’s three-level namespace.
// It’s used to organize your data assets. Users can see all catalogs on which
// they have been assigned the USE_CATALOG data permission.
//
// In Unity Catalog, admins and data stewards manage users and their access to
// data centrally across all of the workspaces in a Databricks account. Users in
// different workspaces can share access to the same data, depending on
// privileges granted centrally in Unity Catalog.
type CatalogsService interface {

	// Create a catalog.
	//
	// Creates a new catalog instance in the parent metastore if the caller is a
	// metastore admin or has the **CREATE_CATALOG** privilege.
	Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error)

	// Delete a catalog.
	//
	// Deletes the catalog that matches the supplied name. The caller must be a
	// metastore admin or the owner of the catalog.
	Delete(ctx context.Context, request DeleteCatalogRequest) error

	// Get a catalog.
	//
	// Gets the specified catalog in a metastore. The caller must be a metastore
	// admin, the owner of the catalog, or a user that has the **USE_CATALOG**
	// privilege set for their account.
	Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error)

	// List catalogs.
	//
	// Gets an array of catalogs in the metastore. If the caller is the
	// metastore admin, all catalogs will be retrieved. Otherwise, only catalogs
	// owned by the caller (or for which the caller has the **USE_CATALOG**
	// privilege) will be retrieved. There is no guarantee of a specific
	// ordering of the elements in the array.
	//
	// Use ListAll() to get all CatalogInfo instances
	List(ctx context.Context) (*ListCatalogsResponse, error)

	// Update a catalog.
	//
	// Updates the catalog that matches the supplied name. The caller must be
	// either the owner of the catalog, or a metastore admin (when changing the
	// owner field of the catalog).
	Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error)
}

// An external location is an object that combines a cloud storage path with a
// storage credential that authorizes access to the cloud storage path. Each
// external location is subject to Unity Catalog access-control policies that
// control which users and groups can access the credential. If a user does not
// have access to an external location in Unity Catalog, the request fails and
// Unity Catalog does not attempt to authenticate to your cloud tenant on the
// user’s behalf.
//
// Databricks recommends using external locations rather than using storage
// credentials directly.
//
// To create external locations, you must be a metastore admin or a user with
// the **CREATE_EXTERNAL_LOCATION** privilege.
type ExternalLocationsService interface {

	// Create an external location.
	//
	// Creates a new external location entry in the metastore. The caller must
	// be a metastore admin or have the **CREATE_EXTERNAL_LOCATION** privilege
	// on both the metastore and the associated storage credential.
	Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error)

	// Delete an external location.
	//
	// Deletes the specified external location from the metastore. The caller
	// must be the owner of the external location.
	Delete(ctx context.Context, request DeleteExternalLocationRequest) error

	// Get an external location.
	//
	// Gets an external location from the metastore. The caller must be either a
	// metastore admin, the owner of the external location, or a user that has
	// some privilege on the external location.
	Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error)

	// List external locations.
	//
	// Gets an array of external locations (__ExternalLocationInfo__ objects)
	// from the metastore. The caller must be a metastore admin, the owner of
	// the external location, or a user that has some privilege on the external
	// location. There is no guarantee of a specific ordering of the elements in
	// the array.
	//
	// Use ListAll() to get all ExternalLocationInfo instances
	List(ctx context.Context) (*ListExternalLocationsResponse, error)

	// Update an external location.
	//
	// Updates an external location in the metastore. The caller must be the
	// owner of the external location, or be a metastore admin. In the second
	// case, the admin can only update the name of the external location.
	Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error)
}

// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
//
// The function implementation can be any SQL expression or Query, and it can be
// invoked wherever a table reference is allowed in a query. In Unity Catalog, a
// function resides at the same level as a table, so it can be referenced with
// the form __catalog_name__.__schema_name__.__function_name__.
type FunctionsService interface {

	// Create a function.
	//
	// Creates a new function
	//
	// The user must have the following permissions in order for the function to
	// be created: - **USE_CATALOG** on the function's parent catalog -
	// **USE_SCHEMA** and **CREATE_FUNCTION** on the function's parent schema
	Create(ctx context.Context, request CreateFunction) (*FunctionInfo, error)

	// Delete a function.
	//
	// Deletes the function that matches the supplied name. For the deletion to
	// succeed, the user must satisfy one of the following conditions: - Is the
	// owner of the function's parent catalog - Is the owner of the function's
	// parent schema and have the **USE_CATALOG** privilege on its parent
	// catalog - Is the owner of the function itself and have both the
	// **USE_CATALOG** privilege on its parent catalog and the **USE_SCHEMA**
	// privilege on its parent schema
	Delete(ctx context.Context, request DeleteFunctionRequest) error

	// Get a function.
	//
	// Gets a function from within a parent catalog and schema. For the fetch to
	// succeed, the user must satisfy one of the following requirements: - Is a
	// metastore admin - Is an owner of the function's parent catalog - Have the
	// **USE_CATALOG** privilege on the function's parent catalog and be the
	// owner of the function - Have the **USE_CATALOG** privilege on the
	// function's parent catalog, the **USE_SCHEMA** privilege on the function's
	// parent schema, and the **EXECUTE** privilege on the function itself
	Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error)

	// List functions.
	//
	// List functions within the specified parent catalog and schema. If the
	// user is a metastore admin, all functions are returned in the output list.
	// Otherwise, the user must have the **USE_CATALOG** privilege on the
	// catalog and the **USE_SCHEMA** privilege on the schema, and the output
	// list contains only functions for which either the user has the
	// **EXECUTE** privilege or the user is the owner. There is no guarantee of
	// a specific ordering of the elements in the array.
	List(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error)

	// Update a function.
	//
	// Updates the function that matches the supplied name. Only the owner of
	// the function can be updated. If the user is not a metastore admin, the
	// user must be a member of the group that is the new function owner. - Is a
	// metastore admin - Is the owner of the function's parent catalog - Is the
	// owner of the function's parent schema and has the **USE_CATALOG**
	// privilege on its parent catalog - Is the owner of the function itself and
	// has the **USE_CATALOG** privilege on its parent catalog as well as the
	// **USE_SCHEMA** privilege on the function's parent schema.
	Update(ctx context.Context, request UpdateFunction) (*FunctionInfo, error)
}

// In Unity Catalog, data is secure by default. Initially, users have no access
// to data in a metastore. Access can be granted by either a metastore admin,
// the owner of an object, or the owner of the catalog or schema that contains
// the object. Securable objects in Unity Catalog are hierarchical and
// privileges are inherited downward.
//
// Securable objects in Unity Catalog are hierarchical and privileges are
// inherited downward. This means that granting a privilege on the catalog
// automatically grants the privilege to all current and future objects within
// the catalog. Similarly, privileges granted on a schema are inherited by all
// current and future objects within that schema.
type GrantsService interface {

	// Get permissions.
	//
	// Gets the permissions for a securable.
	Get(ctx context.Context, request GetGrantRequest) (*PermissionsList, error)

	// Get effective permissions.
	//
	// Gets the effective permissions for a securable.
	GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error)

	// Update permissions.
	//
	// Updates the permissions for a securable.
	Update(ctx context.Context, request UpdatePermissions) (*PermissionsList, error)
}

// A metastore is the top-level container of objects in Unity Catalog. It stores
// data assets (tables and views) and the permissions that govern access to
// them. Databricks account admins can create metastores and assign them to
// Databricks workspaces to control which workloads use each metastore. For a
// workspace to use Unity Catalog, it must have a Unity Catalog metastore
// attached.
//
// Each metastore is configured with a root storage location in a cloud storage
// account. This storage location is used for metadata and managed tables data.
//
// NOTE: This metastore is distinct from the metastore included in Databricks
// workspaces created before Unity Catalog was released. If your workspace
// includes a legacy Hive metastore, the data in that metastore is available in
// a catalog named hive_metastore.
type MetastoresService interface {

	// Create an assignment.
	//
	// Creates a new metastore assignment. If an assignment for the same
	// __workspace_id__ exists, it will be overwritten by the new
	// __metastore_id__ and __default_catalog_name__. The caller must be an
	// account admin.
	Assign(ctx context.Context, request CreateMetastoreAssignment) error

	// Create a metastore.
	//
	// Creates a new metastore based on a provided name and storage root path.
	Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error)

	// Get metastore assignment for workspace.
	//
	// Gets the metastore assignment for the workspace being accessed.
	Current(ctx context.Context) (*MetastoreAssignment, error)

	// Delete a metastore.
	//
	// Deletes a metastore. The caller must be a metastore admin.
	Delete(ctx context.Context, request DeleteMetastoreRequest) error

	// Get a metastore.
	//
	// Gets a metastore that matches the supplied ID. The caller must be a
	// metastore admin to retrieve this info.
	Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error)

	// List metastores.
	//
	// Gets an array of the available metastores (as __MetastoreInfo__ objects).
	// The caller must be an admin to retrieve this info. There is no guarantee
	// of a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all MetastoreInfo instances
	List(ctx context.Context) (*ListMetastoresResponse, error)

	// Get a metastore summary.
	//
	// Gets information about a metastore. This summary includes the storage
	// credential, the cloud vendor, the cloud region, and the global metastore
	// ID.
	Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error)

	// Delete an assignment.
	//
	// Deletes a metastore assignment. The caller must be an account
	// administrator.
	Unassign(ctx context.Context, request UnassignRequest) error

	// Update a metastore.
	//
	// Updates information for a specific metastore. The caller must be a
	// metastore admin.
	Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error)

	// Update an assignment.
	//
	// Updates a metastore assignment. This operation can be used to update
	// __metastore_id__ or __default_catalog_name__ for a specified Workspace,
	// if the Workspace is already assigned a metastore. The caller must be an
	// account admin to update __metastore_id__; otherwise, the caller can be a
	// Workspace admin.
	UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error
}

// Databricks Delta Sharing: Providers REST API
type ProvidersService interface {

	// Create an auth provider.
	//
	// Creates a new authentication provider minimally based on a name and
	// authentication type. The caller must be an admin on the metastore.
	Create(ctx context.Context, request CreateProvider) (*ProviderInfo, error)

	// Delete a provider.
	//
	// Deletes an authentication provider, if the caller is a metastore admin or
	// is the owner of the provider.
	Delete(ctx context.Context, request DeleteProviderRequest) error

	// Get a provider.
	//
	// Gets a specific authentication provider. The caller must supply the name
	// of the provider, and must either be a metastore admin or the owner of the
	// provider.
	Get(ctx context.Context, request GetProviderRequest) (*ProviderInfo, error)

	// List providers.
	//
	// Gets an array of available authentication providers. The caller must
	// either be a metastore admin or the owner of the providers. Providers not
	// owned by the caller are not included in the response. There is no
	// guarantee of a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all ProviderInfo instances
	List(ctx context.Context, request ListProvidersRequest) (*ListProvidersResponse, error)

	// List shares by Provider.
	//
	// Gets an array of a specified provider's shares within the metastore
	// where:
	//
	// * the caller is a metastore admin, or * the caller is the owner.
	ListShares(ctx context.Context, request ListSharesRequest) (*ListProviderSharesResponse, error)

	// Update a provider.
	//
	// Updates the information for an authentication provider, if the caller is
	// a metastore admin or is the owner of the provider. If the update changes
	// the provider name, the caller must be both a metastore admin and the
	// owner of the provider.
	Update(ctx context.Context, request UpdateProvider) (*ProviderInfo, error)
}

// Databricks Delta Sharing: Recipient Activation REST API
type RecipientActivationService interface {

	// Get a share activation URL.
	//
	// Gets an activation URL for a share.
	GetActivationUrlInfo(ctx context.Context, request GetActivationUrlInfoRequest) error

	// Get an access token.
	//
	// Retrieve access token with an activation url. This is a public API
	// without any authentication.
	RetrieveToken(ctx context.Context, request RetrieveTokenRequest) (*RetrieveTokenResponse, error)
}

// Databricks Delta Sharing: Recipients REST API
type RecipientsService interface {

	// Create a share recipient.
	//
	// Creates a new recipient with the delta sharing authentication type in the
	// metastore. The caller must be a metastore admin or has the
	// **CREATE_RECIPIENT** privilege on the metastore.
	Create(ctx context.Context, request CreateRecipient) (*RecipientInfo, error)

	// Delete a share recipient.
	//
	// Deletes the specified recipient from the metastore. The caller must be
	// the owner of the recipient.
	Delete(ctx context.Context, request DeleteRecipientRequest) error

	// Get a share recipient.
	//
	// Gets a share recipient from the metastore if:
	//
	// * the caller is the owner of the share recipient, or: * is a metastore
	// admin
	Get(ctx context.Context, request GetRecipientRequest) (*RecipientInfo, error)

	// List share recipients.
	//
	// Gets an array of all share recipients within the current metastore where:
	//
	// * the caller is a metastore admin, or * the caller is the owner. There is
	// no guarantee of a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all RecipientInfo instances
	List(ctx context.Context, request ListRecipientsRequest) (*ListRecipientsResponse, error)

	// Rotate a token.
	//
	// Refreshes the specified recipient's delta sharing authentication token
	// with the provided token info. The caller must be the owner of the
	// recipient.
	RotateToken(ctx context.Context, request RotateRecipientToken) (*RecipientInfo, error)

	// Get recipient share permissions.
	//
	// Gets the share permissions for the specified Recipient. The caller must
	// be a metastore admin or the owner of the Recipient.
	SharePermissions(ctx context.Context, request SharePermissionsRequest) (*GetRecipientSharePermissionsResponse, error)

	// Update a share recipient.
	//
	// Updates an existing recipient in the metastore. The caller must be a
	// metastore admin or the owner of the recipient. If the recipient name will
	// be updated, the user must be both a metastore admin and the owner of the
	// recipient.
	Update(ctx context.Context, request UpdateRecipient) error
}

// A schema (also called a database) is the second layer of Unity Catalog’s
// three-level namespace. A schema organizes tables, views and functions. To
// access (or list) a table or view in a schema, users must have the USE_SCHEMA
// data permission on the schema and its parent catalog, and they must have the
// SELECT permission on the table or view.
type SchemasService interface {

	// Create a schema.
	//
	// Creates a new schema for catalog in the Metatastore. The caller must be a
	// metastore admin, or have the **CREATE_SCHEMA** privilege in the parent
	// catalog.
	Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error)

	// Delete a schema.
	//
	// Deletes the specified schema from the parent catalog. The caller must be
	// the owner of the schema or an owner of the parent catalog.
	Delete(ctx context.Context, request DeleteSchemaRequest) error

	// Get a schema.
	//
	// Gets the specified schema within the metastore. The caller must be a
	// metastore admin, the owner of the schema, or a user that has the
	// **USE_SCHEMA** privilege on the schema.
	Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error)

	// List schemas.
	//
	// Gets an array of schemas for a catalog in the metastore. If the caller is
	// the metastore admin or the owner of the parent catalog, all schemas for
	// the catalog will be retrieved. Otherwise, only schemas owned by the
	// caller (or for which the caller has the **USE_SCHEMA** privilege) will be
	// retrieved. There is no guarantee of a specific ordering of the elements
	// in the array.
	//
	// Use ListAll() to get all SchemaInfo instances
	List(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error)

	// Update a schema.
	//
	// Updates a schema for a catalog. The caller must be the owner of the
	// schema or a metastore admin. If the caller is a metastore admin, only the
	// __owner__ field can be changed in the update. If the __name__ field must
	// be updated, the caller must be a metastore admin or have the
	// **CREATE_SCHEMA** privilege on the parent catalog.
	Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error)
}

// Databricks Delta Sharing: Shares REST API
type SharesService interface {

	// Create a share.
	//
	// Creates a new share for data objects. Data objects can be added at this
	// time or after creation with **update**. The caller must be a metastore
	// admin or have the **CREATE_SHARE** privilege on the metastore.
	Create(ctx context.Context, request CreateShare) (*ShareInfo, error)

	// Delete a share.
	//
	// Deletes a data object share from the metastore. The caller must be an
	// owner of the share.
	Delete(ctx context.Context, request DeleteShareRequest) error

	// Get a share.
	//
	// Gets a data object share from the metastore. The caller must be a
	// metastore admin or the owner of the share.
	Get(ctx context.Context, request GetShareRequest) (*ShareInfo, error)

	// List shares.
	//
	// Gets an array of data object shares from the metastore. The caller must
	// be a metastore admin or the owner of the share. There is no guarantee of
	// a specific ordering of the elements in the array.
	//
	// Use ListAll() to get all ShareInfo instances
	List(ctx context.Context) (*ListSharesResponse, error)

	// Get permissions.
	//
	// Gets the permissions for a data share from the metastore. The caller must
	// be a metastore admin or the owner of the share.
	SharePermissions(ctx context.Context, request SharePermissionsRequest) (*PermissionsList, error)

	// Update a share.
	//
	// Updates the share with the changes and data objects in the request. The
	// caller must be the owner of the share or a metastore admin.
	//
	// When the caller is a metastore admin, only the __owner__ field can be
	// updated.
	//
	// In the case that the share name is changed, **updateShare** requires that
	// the caller is both the share owner and a metastore admin.
	//
	// For each table that is added through this method, the share owner must
	// also have **SELECT** privilege on the table. This privilege must be
	// maintained indefinitely for recipients to be able to access the table.
	// Typically, you should use a group as the share owner.
	//
	// Table removals through **update** do not require additional privileges.
	Update(ctx context.Context, request UpdateShare) (*ShareInfo, error)

	// Update permissions.
	//
	// Updates the permissions for a data share in the metastore. The caller
	// must be a metastore admin or an owner of the share.
	//
	// For new recipient grants, the user must also be the owner of the
	// recipients. recipient revocations do not require additional privileges.
	UpdatePermissions(ctx context.Context, request UpdateSharePermissions) error
}

// A storage credential represents an authentication and authorization mechanism
// for accessing data stored on your cloud tenant. Each storage credential is
// subject to Unity Catalog access-control policies that control which users and
// groups can access the credential. If a user does not have access to a storage
// credential in Unity Catalog, the request fails and Unity Catalog does not
// attempt to authenticate to your cloud tenant on the user’s behalf.
//
// Databricks recommends using external locations rather than using storage
// credentials directly.
//
// To create storage credentials, you must be a Databricks account admin. The
// account admin who creates the storage credential can delegate ownership to
// another user or group to manage permissions on it.
type StorageCredentialsService interface {

	// Create a storage credential.
	//
	// Creates a new storage credential. The request object is specific to the
	// cloud:
	//
	// * **AwsIamRole** for AWS credentials * **AzureServicePrincipal** for
	// Azure credentials * **GcpServiceAcountKey** for GCP credentials.
	//
	// The caller must be a metastore admin and have the
	// **CREATE_STORAGE_CREDENTIAL** privilege on the metastore.
	Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error)

	// Delete a credential.
	//
	// Deletes a storage credential from the metastore. The caller must be an
	// owner of the storage credential.
	Delete(ctx context.Context, request DeleteStorageCredentialRequest) error

	// Get a credential.
	//
	// Gets a storage credential from the metastore. The caller must be a
	// metastore admin, the owner of the storage credential, or have some
	// permission on the storage credential.
	Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error)

	// List credentials.
	//
	// Gets an array of storage credentials (as __StorageCredentialInfo__
	// objects). The array is limited to only those storage credentials the
	// caller has permission to access. If the caller is a metastore admin, all
	// storage credentials will be retrieved. There is no guarantee of a
	// specific ordering of the elements in the array.
	//
	// Use ListAll() to get all StorageCredentialInfo instances
	List(ctx context.Context) (*ListStorageCredentialsResponse, error)

	// Update a credential.
	//
	// Updates a storage credential on the metastore. The caller must be the
	// owner of the storage credential or a metastore admin. If the caller is a
	// metastore admin, only the __owner__ credential can be changed.
	Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error)

	// Validate a storage credential.
	//
	// Validates a storage credential. At least one of
	// __external_location_name__ and __url__ need to be provided. If only one
	// of them is provided, it will be used for validation. And if both are
	// provided, the __url__ will be used for validation, and
	// __external_location_name__ will be ignored when checking overlapping
	// urls.
	//
	// Either the __storage_credential_name__ or the cloud-specific credential
	// must be provided.
	//
	// The caller must be a metastore admin or the storage credential owner or
	// have the **CREATE_EXTERNAL_LOCATION** privilege on the metastore and the
	// storage credential.
	Validate(ctx context.Context, request ValidateStorageCredential) (*ValidateStorageCredentialResponse, error)
}

// Primary key and foreign key constraints encode relationships between fields
// in tables.
//
// Primary and foreign keys are informational only and are not enforced. Foreign
// keys must reference a primary key in another table. This primary key is the
// parent constraint of the foreign key and the table this primary key is on is
// the parent table of the foreign key. Similarly, the foreign key is the child
// constraint of its referenced primary key; the table of the foreign key is the
// child table of the primary key.
//
// You can declare primary keys and foreign keys as part of the table
// specification during table creation. You can also add or drop constraints on
// existing tables.
type TableConstraintsService interface {

	// Create a table constraint.
	//
	// Creates a new table constraint.
	//
	// For the table constraint creation to succeed, the user must satisfy both
	// of these conditions: - the user must have the **USE_CATALOG** privilege
	// on the table's parent catalog, the **USE_SCHEMA** privilege on the
	// table's parent schema, and be the owner of the table. - if the new
	// constraint is a __ForeignKeyConstraint__, the user must have the
	// **USE_CATALOG** privilege on the referenced parent table's catalog, the
	// **USE_SCHEMA** privilege on the referenced parent table's schema, and be
	// the owner of the referenced parent table.
	Create(ctx context.Context, request CreateTableConstraint) (*TableConstraint, error)

	// Delete a table constraint.
	//
	// Deletes a table constraint.
	//
	// For the table constraint deletion to succeed, the user must satisfy both
	// of these conditions: - the user must have the **USE_CATALOG** privilege
	// on the table's parent catalog, the **USE_SCHEMA** privilege on the
	// table's parent schema, and be the owner of the table. - if __cascade__
	// argument is **true**, the user must have the following permissions on all
	// of the child tables: the **USE_CATALOG** privilege on the table's
	// catalog, the **USE_SCHEMA** privilege on the table's schema, and be the
	// owner of the table.
	Delete(ctx context.Context, request DeleteTableConstraintRequest) error
}

// A table resides in the third layer of Unity Catalog’s three-level
// namespace. It contains rows of data. To create a table, users must have
// CREATE_TABLE and USE_SCHEMA permissions on the schema, and they must have the
// USE_CATALOG permission on its parent catalog. To query a table, users must
// have the SELECT permission on the table, and they must have the USE_CATALOG
// permission on its parent catalog and the USE_SCHEMA permission on its parent
// schema.
//
// A table can be managed or external. From an API perspective, a __VIEW__ is a
// particular kind of table (rather than a managed or external table).
type TablesService interface {

	// Delete a table.
	//
	// Deletes a table from the specified parent catalog and schema. The caller
	// must be the owner of the parent catalog, have the **USE_CATALOG**
	// privilege on the parent catalog and be the owner of the parent schema, or
	// be the owner of the table and have the **USE_CATALOG** privilege on the
	// parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	Delete(ctx context.Context, request DeleteTableRequest) error

	// Get a table.
	//
	// Gets a table from the metastore for a specific catalog and schema. The
	// caller must be a metastore admin, be the owner of the table and have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema, or be the owner of the table and have the
	// **SELECT** privilege on it as well.
	Get(ctx context.Context, request GetTableRequest) (*TableInfo, error)

	// List tables.
	//
	// Gets an array of all tables for the current metastore under the parent
	// catalog and schema. The caller must be a metastore admin or an owner of
	// (or have the **SELECT** privilege on) the table. For the latter case, the
	// caller must also be the owner or have the **USE_CATALOG** privilege on
	// the parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	// There is no guarantee of a specific ordering of the elements in the
	// array.
	//
	// Use ListAll() to get all TableInfo instances
	List(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error)

	// List table summaries.
	//
	// Gets an array of summaries for tables for a schema and catalog within the
	// metastore. The table summaries returned are either:
	//
	// * summaries for all tables (within the current metastore and parent
	// catalog and schema), when the user is a metastore admin, or: * summaries
	// for all tables and schemas (within the current metastore and parent
	// catalog) for which the user has ownership or the **SELECT** privilege on
	// the table and ownership or **USE_SCHEMA** privilege on the schema,
	// provided that the user also has ownership or the **USE_CATALOG**
	// privilege on the parent catalog.
	//
	// There is no guarantee of a specific ordering of the elements in the
	// array.
	ListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error)
}
