// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"context"
)

// These APIs manage metastore assignments to a workspace.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountMetastoreAssignmentsService interface {

	// Creates an assignment to a metastore for a workspace
	Create(ctx context.Context, request AccountsCreateMetastoreAssignment) error

	// Deletes a metastore assignment to a workspace, leaving the workspace with
	// no metastore.
	Delete(ctx context.Context, request DeleteAccountMetastoreAssignmentRequest) error

	// Gets the metastore assignment, if any, for the workspace specified by ID.
	// If the workspace is assigned a metastore, the mappig will be returned. If
	// no metastore is assigned to the workspace, the assignment will not be
	// found and a 404 returned.
	Get(ctx context.Context, request GetAccountMetastoreAssignmentRequest) (*AccountsMetastoreAssignment, error)

	// Gets a list of all Databricks workspace IDs that have been assigned to
	// given metastore.
	List(ctx context.Context, request ListAccountMetastoreAssignmentsRequest) (*ListAccountMetastoreAssignmentsResponse, error)

	// Updates an assignment to a metastore for a workspace. Currently, only the
	// default catalog may be updated.
	Update(ctx context.Context, request AccountsUpdateMetastoreAssignment) error
}

// These APIs manage Unity Catalog metastores for an account. A metastore
// contains catalogs that can be associated with workspaces
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountMetastoresService interface {

	// Creates a Unity Catalog metastore.
	Create(ctx context.Context, request AccountsCreateMetastore) (*AccountsMetastoreInfo, error)

	// Deletes a Unity Catalog metastore for an account, both specified by ID.
	Delete(ctx context.Context, request DeleteAccountMetastoreRequest) error

	// Gets a Unity Catalog metastore from an account, both specified by ID.
	Get(ctx context.Context, request GetAccountMetastoreRequest) (*AccountsMetastoreInfo, error)

	// Gets all Unity Catalog metastores associated with an account specified by
	// ID.
	List(ctx context.Context) (*ListMetastoresResponse, error)

	// Updates an existing Unity Catalog metastore.
	Update(ctx context.Context, request AccountsUpdateMetastore) (*AccountsMetastoreInfo, error)
}

// These APIs manage storage credentials for a particular metastore.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type AccountStorageCredentialsService interface {

	// Creates a new storage credential. The request object is specific to the
	// cloud:
	//
	// * **AwsIamRole** for AWS credentials * **AzureServicePrincipal** for
	// Azure credentials * **GcpServiceAcountKey** for GCP credentials.
	//
	// The caller must be a metastore admin and have the
	// **CREATE_STORAGE_CREDENTIAL** privilege on the metastore.
	Create(ctx context.Context, request AccountsCreateStorageCredential) (*AccountsStorageCredentialInfo, error)

	// Deletes a storage credential from the metastore. The caller must be an
	// owner of the storage credential.
	Delete(ctx context.Context, request DeleteAccountStorageCredentialRequest) error

	// Gets a storage credential from the metastore. The caller must be a
	// metastore admin, the owner of the storage credential, or have a level of
	// privilege on the storage credential.
	Get(ctx context.Context, request GetAccountStorageCredentialRequest) (*AccountsStorageCredentialInfo, error)

	// Gets a list of all storage credentials that have been assigned to given
	// metastore.
	List(ctx context.Context, request ListAccountStorageCredentialsRequest) (*ListAccountStorageCredentialsResponse, error)

	// Updates a storage credential on the metastore. The caller must be the
	// owner of the storage credential. If the caller is a metastore admin, only
	// the __owner__ credential can be changed.
	Update(ctx context.Context, request AccountsUpdateStorageCredential) (*AccountsStorageCredentialInfo, error)
}

// In Databricks Runtime 13.3 and above, you can add libraries and init scripts
// to the `allowlist` in UC so that users can leverage these artifacts on
// compute configured with shared access mode.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ArtifactAllowlistsService interface {

	// Get the artifact allowlist of a certain artifact type. The caller must be
	// a metastore admin or have the **MANAGE ALLOWLIST** privilege on the
	// metastore.
	Get(ctx context.Context, request GetArtifactAllowlistRequest) (*ArtifactAllowlistInfo, error)

	// Set the artifact allowlist of a certain artifact type. The whole artifact
	// allowlist is replaced with the new allowlist. The caller must be a
	// metastore admin or have the **MANAGE ALLOWLIST** privilege on the
	// metastore.
	Update(ctx context.Context, request SetArtifactAllowlist) (*ArtifactAllowlistInfo, error)
}

// A catalog is the first layer of Unity Catalog’s three-level namespace.
// It’s used to organize your data assets. Users can see all catalogs on which
// they have been assigned the USE_CATALOG data permission.
//
// In Unity Catalog, admins and data stewards manage users and their access to
// data centrally across all of the workspaces in a Databricks account. Users in
// different workspaces can share access to the same data, depending on
// privileges granted centrally in Unity Catalog.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CatalogsService interface {

	// Creates a new catalog instance in the parent metastore if the caller is a
	// metastore admin or has the **CREATE_CATALOG** privilege.
	Create(ctx context.Context, request CreateCatalog) (*CatalogInfo, error)

	// Deletes the catalog that matches the supplied name. The caller must be a
	// metastore admin or the owner of the catalog.
	Delete(ctx context.Context, request DeleteCatalogRequest) error

	// Gets the specified catalog in a metastore. The caller must be a metastore
	// admin, the owner of the catalog, or a user that has the **USE_CATALOG**
	// privilege set for their account.
	Get(ctx context.Context, request GetCatalogRequest) (*CatalogInfo, error)

	// Gets an array of catalogs in the metastore. If the caller is the
	// metastore admin, all catalogs will be retrieved. Otherwise, only catalogs
	// owned by the caller (or for which the caller has the **USE_CATALOG**
	// privilege) will be retrieved. There is no guarantee of a specific
	// ordering of the elements in the array.
	List(ctx context.Context, request ListCatalogsRequest) (*ListCatalogsResponse, error)

	// Updates the catalog that matches the supplied name. The caller must be
	// either the owner of the catalog, or a metastore admin (when changing the
	// owner field of the catalog).
	Update(ctx context.Context, request UpdateCatalog) (*CatalogInfo, error)
}

// Connections allow for creating a connection to an external data source.
//
// A connection is an abstraction of an external data source that can be
// connected from Databricks Compute. Creating a connection object is the first
// step to managing external data sources within Unity Catalog, with the second
// step being creating a data object (catalog, schema, or table) using the
// connection. Data objects derived from a connection can be written to or read
// from similar to other Unity Catalog data objects based on cloud storage.
// Users may create different types of connections with each connection having a
// unique set of configuration options to support credential management and
// other settings.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ConnectionsService interface {

	// Creates a new connection
	//
	// Creates a new connection to an external data source. It allows users to
	// specify connection details and configurations for interaction with the
	// external server.
	Create(ctx context.Context, request CreateConnection) (*ConnectionInfo, error)

	// Deletes the connection that matches the supplied name.
	Delete(ctx context.Context, request DeleteConnectionRequest) error

	// Gets a connection from it's name.
	Get(ctx context.Context, request GetConnectionRequest) (*ConnectionInfo, error)

	// List all connections.
	List(ctx context.Context, request ListConnectionsRequest) (*ListConnectionsResponse, error)

	// Updates the connection that matches the supplied name.
	Update(ctx context.Context, request UpdateConnection) (*ConnectionInfo, error)
}

// A credential represents an authentication and authorization mechanism for
// accessing services on your cloud tenant. Each credential is subject to Unity
// Catalog access-control policies that control which users and groups can
// access the credential.
//
// To create credentials, you must be a Databricks account admin or have the
// `CREATE SERVICE CREDENTIAL` privilege. The user who creates the credential
// can delegate ownership to another user or group to manage permissions on it.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CredentialsService interface {

	// Creates a new credential. The type of credential to be created is
	// determined by the **purpose** field, which should be either **SERVICE**
	// or **STORAGE**.
	//
	// The caller must be a metastore admin or have the metastore privilege
	// **CREATE_STORAGE_CREDENTIAL** for storage credentials, or
	// **CREATE_SERVICE_CREDENTIAL** for service credentials.
	CreateCredential(ctx context.Context, request CreateCredentialRequest) (*CredentialInfo, error)

	// Deletes a service or storage credential from the metastore. The caller
	// must be an owner of the credential.
	DeleteCredential(ctx context.Context, request DeleteCredentialRequest) error

	// Returns a set of temporary credentials generated using the specified
	// service credential. The caller must be a metastore admin or have the
	// metastore privilege **ACCESS** on the service credential.
	GenerateTemporaryServiceCredential(ctx context.Context, request GenerateTemporaryServiceCredentialRequest) (*TemporaryCredentials, error)

	// Gets a service or storage credential from the metastore. The caller must
	// be a metastore admin, the owner of the credential, or have any permission
	// on the credential.
	GetCredential(ctx context.Context, request GetCredentialRequest) (*CredentialInfo, error)

	// Gets an array of credentials (as __CredentialInfo__ objects).
	//
	// The array is limited to only the credentials that the caller has
	// permission to access. If the caller is a metastore admin, retrieval of
	// credentials is unrestricted. There is no guarantee of a specific ordering
	// of the elements in the array.
	ListCredentials(ctx context.Context, request ListCredentialsRequest) (*ListCredentialsResponse, error)

	// Updates a service or storage credential on the metastore.
	//
	// The caller must be the owner of the credential or a metastore admin or
	// have the `MANAGE` permission. If the caller is a metastore admin, only
	// the __owner__ field can be changed.
	UpdateCredential(ctx context.Context, request UpdateCredentialRequest) (*CredentialInfo, error)

	// Validates a credential.
	//
	// For service credentials (purpose is **SERVICE**), either the
	// __credential_name__ or the cloud-specific credential must be provided.
	//
	// For storage credentials (purpose is **STORAGE**), at least one of
	// __external_location_name__ and __url__ need to be provided. If only one
	// of them is provided, it will be used for validation. And if both are
	// provided, the __url__ will be used for validation, and
	// __external_location_name__ will be ignored when checking overlapping
	// urls. Either the __credential_name__ or the cloud-specific credential
	// must be provided.
	//
	// The caller must be a metastore admin or the credential owner or have the
	// required permission on the metastore and the credential (e.g.,
	// **CREATE_EXTERNAL_LOCATION** when purpose is **STORAGE**).
	ValidateCredential(ctx context.Context, request ValidateCredentialRequest) (*ValidateCredentialResponse, error)
}

// External Lineage APIs enable defining and managing lineage relationships
// between Databricks objects and external systems. These APIs allow users to
// capture data flows connecting Databricks tables, models, and file paths with
// external metadata objects.
//
// With these APIs, users can create, update, delete, and list lineage
// relationships with support for column-level mappings and custom properties.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ExternalLineageService interface {

	// Creates an external lineage relationship between a Databricks or external
	// metadata object and another external metadata object.
	CreateExternalLineageRelationship(ctx context.Context, request CreateExternalLineageRelationshipRequest) (*ExternalLineageRelationship, error)

	// Deletes an external lineage relationship between a Databricks or external
	// metadata object and another external metadata object.
	DeleteExternalLineageRelationship(ctx context.Context, request DeleteExternalLineageRelationshipRequest) error

	// Lists external lineage relationships of a Databricks object or external
	// metadata given a supplied direction.
	ListExternalLineageRelationships(ctx context.Context, request ListExternalLineageRelationshipsRequest) (*ListExternalLineageRelationshipsResponse, error)

	// Updates an external lineage relationship between a Databricks or external
	// metadata object and another external metadata object.
	UpdateExternalLineageRelationship(ctx context.Context, request UpdateExternalLineageRelationshipRequest) (*ExternalLineageRelationship, error)
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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ExternalLocationsService interface {

	// Creates a new external location entry in the metastore. The caller must
	// be a metastore admin or have the **CREATE_EXTERNAL_LOCATION** privilege
	// on both the metastore and the associated storage credential.
	Create(ctx context.Context, request CreateExternalLocation) (*ExternalLocationInfo, error)

	// Deletes the specified external location from the metastore. The caller
	// must be the owner of the external location.
	Delete(ctx context.Context, request DeleteExternalLocationRequest) error

	// Gets an external location from the metastore. The caller must be either a
	// metastore admin, the owner of the external location, or a user that has
	// some privilege on the external location.
	Get(ctx context.Context, request GetExternalLocationRequest) (*ExternalLocationInfo, error)

	// Gets an array of external locations (__ExternalLocationInfo__ objects)
	// from the metastore. The caller must be a metastore admin, the owner of
	// the external location, or a user that has some privilege on the external
	// location. There is no guarantee of a specific ordering of the elements in
	// the array.
	List(ctx context.Context, request ListExternalLocationsRequest) (*ListExternalLocationsResponse, error)

	// Updates an external location in the metastore. The caller must be the
	// owner of the external location, or be a metastore admin. In the second
	// case, the admin can only update the name of the external location.
	Update(ctx context.Context, request UpdateExternalLocation) (*ExternalLocationInfo, error)
}

// External Metadata objects enable customers to register and manage metadata
// about external systems within Unity Catalog.
//
// These APIs provide a standardized way to create, update, retrieve, list, and
// delete external metadata objects. Fine-grained authorization ensures that
// only users with appropriate permissions can view and manage external metadata
// objects.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ExternalMetadataService interface {

	// Creates a new external metadata object in the parent metastore if the
	// caller is a metastore admin or has the **CREATE_EXTERNAL_METADATA**
	// privilege. Grants **BROWSE** to all account users upon creation by
	// default.
	CreateExternalMetadata(ctx context.Context, request CreateExternalMetadataRequest) (*ExternalMetadata, error)

	// Deletes the external metadata object that matches the supplied name. The
	// caller must be a metastore admin, the owner of the external metadata
	// object, or a user that has the **MANAGE** privilege.
	DeleteExternalMetadata(ctx context.Context, request DeleteExternalMetadataRequest) error

	// Gets the specified external metadata object in a metastore. The caller
	// must be a metastore admin, the owner of the external metadata object, or
	// a user that has the **BROWSE** privilege.
	GetExternalMetadata(ctx context.Context, request GetExternalMetadataRequest) (*ExternalMetadata, error)

	// Gets an array of external metadata objects in the metastore. If the
	// caller is the metastore admin, all external metadata objects will be
	// retrieved. Otherwise, only external metadata objects that the caller has
	// **BROWSE** on will be retrieved. There is no guarantee of a specific
	// ordering of the elements in the array.
	ListExternalMetadata(ctx context.Context, request ListExternalMetadataRequest) (*ListExternalMetadataResponse, error)

	// Updates the external metadata object that matches the supplied name. The
	// caller can only update either the owner or other metadata fields in one
	// request. The caller must be a metastore admin, the owner of the external
	// metadata object, or a user that has the **MODIFY** privilege. If the
	// caller is updating the owner, they must also have the **MANAGE**
	// privilege.
	UpdateExternalMetadata(ctx context.Context, request UpdateExternalMetadataRequest) (*ExternalMetadata, error)
}

// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
//
// The function implementation can be any SQL expression or Query, and it can be
// invoked wherever a table reference is allowed in a query. In Unity Catalog, a
// function resides at the same level as a table, so it can be referenced with
// the form __catalog_name__.__schema_name__.__function_name__.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type FunctionsService interface {

	// **WARNING: This API is experimental and will change in future versions**
	//
	// Creates a new function
	//
	// The user must have the following permissions in order for the function to
	// be created: - **USE_CATALOG** on the function's parent catalog -
	// **USE_SCHEMA** and **CREATE_FUNCTION** on the function's parent schema
	Create(ctx context.Context, request CreateFunctionRequest) (*FunctionInfo, error)

	// Deletes the function that matches the supplied name. For the deletion to
	// succeed, the user must satisfy one of the following conditions: - Is the
	// owner of the function's parent catalog - Is the owner of the function's
	// parent schema and have the **USE_CATALOG** privilege on its parent
	// catalog - Is the owner of the function itself and have both the
	// **USE_CATALOG** privilege on its parent catalog and the **USE_SCHEMA**
	// privilege on its parent schema
	Delete(ctx context.Context, request DeleteFunctionRequest) error

	// Gets a function from within a parent catalog and schema. For the fetch to
	// succeed, the user must satisfy one of the following requirements: - Is a
	// metastore admin - Is an owner of the function's parent catalog - Have the
	// **USE_CATALOG** privilege on the function's parent catalog and be the
	// owner of the function - Have the **USE_CATALOG** privilege on the
	// function's parent catalog, the **USE_SCHEMA** privilege on the function's
	// parent schema, and the **EXECUTE** privilege on the function itself
	Get(ctx context.Context, request GetFunctionRequest) (*FunctionInfo, error)

	// List functions within the specified parent catalog and schema. If the
	// user is a metastore admin, all functions are returned in the output list.
	// Otherwise, the user must have the **USE_CATALOG** privilege on the
	// catalog and the **USE_SCHEMA** privilege on the schema, and the output
	// list contains only functions for which either the user has the
	// **EXECUTE** privilege or the user is the owner. There is no guarantee of
	// a specific ordering of the elements in the array.
	List(ctx context.Context, request ListFunctionsRequest) (*ListFunctionsResponse, error)

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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GrantsService interface {

	// Gets the permissions for a securable. Does not include inherited
	// permissions.
	Get(ctx context.Context, request GetGrantRequest) (*GetPermissionsResponse, error)

	// Gets the effective permissions for a securable. Includes inherited
	// permissions from any parent securables.
	GetEffective(ctx context.Context, request GetEffectiveRequest) (*EffectivePermissionsList, error)

	// Updates the permissions for a securable.
	Update(ctx context.Context, request UpdatePermissions) (*UpdatePermissionsResponse, error)
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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type MetastoresService interface {

	// Creates a new metastore assignment. If an assignment for the same
	// __workspace_id__ exists, it will be overwritten by the new
	// __metastore_id__ and __default_catalog_name__. The caller must be an
	// account admin.
	Assign(ctx context.Context, request CreateMetastoreAssignment) error

	// Creates a new metastore based on a provided name and optional storage
	// root path. By default (if the __owner__ field is not set), the owner of
	// the new metastore is the user calling the __createMetastore__ API. If the
	// __owner__ field is set to the empty string (**""**), the ownership is
	// assigned to the System User instead.
	Create(ctx context.Context, request CreateMetastore) (*MetastoreInfo, error)

	// Gets the metastore assignment for the workspace being accessed.
	Current(ctx context.Context) (*MetastoreAssignment, error)

	// Deletes a metastore. The caller must be a metastore admin.
	Delete(ctx context.Context, request DeleteMetastoreRequest) error

	// Gets a metastore that matches the supplied ID. The caller must be a
	// metastore admin to retrieve this info.
	Get(ctx context.Context, request GetMetastoreRequest) (*MetastoreInfo, error)

	// Gets an array of the available metastores (as __MetastoreInfo__ objects).
	// The caller must be an admin to retrieve this info. There is no guarantee
	// of a specific ordering of the elements in the array.
	List(ctx context.Context, request ListMetastoresRequest) (*ListMetastoresResponse, error)

	// Gets information about a metastore. This summary includes the storage
	// credential, the cloud vendor, the cloud region, and the global metastore
	// ID.
	Summary(ctx context.Context) (*GetMetastoreSummaryResponse, error)

	// Deletes a metastore assignment. The caller must be an account
	// administrator.
	Unassign(ctx context.Context, request UnassignRequest) error

	// Updates information for a specific metastore. The caller must be a
	// metastore admin. If the __owner__ field is set to the empty string
	// (**""**), the ownership is updated to the System User.
	Update(ctx context.Context, request UpdateMetastore) (*MetastoreInfo, error)

	// Updates a metastore assignment. This operation can be used to update
	// __metastore_id__ or __default_catalog_name__ for a specified Workspace,
	// if the Workspace is already assigned a metastore. The caller must be an
	// account admin to update __metastore_id__; otherwise, the caller can be a
	// Workspace admin.
	UpdateAssignment(ctx context.Context, request UpdateMetastoreAssignment) error
}

// Databricks provides a hosted version of MLflow Model Registry in Unity
// Catalog. Models in Unity Catalog provide centralized access control,
// auditing, lineage, and discovery of ML models across Databricks workspaces.
//
// This API reference documents the REST endpoints for managing model versions
// in Unity Catalog. For more details, see the [registered models API
// docs](/api/workspace/registeredmodels).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ModelVersionsService interface {

	// Deletes a model version from the specified registered model. Any aliases
	// assigned to the model version will also be deleted.
	//
	// The caller must be a metastore admin or an owner of the parent registered
	// model. For the latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	Delete(ctx context.Context, request DeleteModelVersionRequest) error

	// Get a model version.
	//
	// The caller must be a metastore admin or an owner of (or have the
	// **EXECUTE** privilege on) the parent registered model. For the latter
	// case, the caller must also be the owner or have the **USE_CATALOG**
	// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
	// parent schema.
	Get(ctx context.Context, request GetModelVersionRequest) (*ModelVersionInfo, error)

	// Get a model version by alias.
	//
	// The caller must be a metastore admin or an owner of (or have the
	// **EXECUTE** privilege on) the registered model. For the latter case, the
	// caller must also be the owner or have the **USE_CATALOG** privilege on
	// the parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	GetByAlias(ctx context.Context, request GetByAliasRequest) (*ModelVersionInfo, error)

	// List model versions. You can list model versions under a particular
	// schema, or list all model versions in the current metastore.
	//
	// The returned models are filtered based on the privileges of the calling
	// user. For example, the metastore admin is able to list all the model
	// versions. A regular user needs to be the owner or have the **EXECUTE**
	// privilege on the parent registered model to recieve the model versions in
	// the response. For the latter case, the caller must also be the owner or
	// have the **USE_CATALOG** privilege on the parent catalog and the
	// **USE_SCHEMA** privilege on the parent schema.
	//
	// There is no guarantee of a specific ordering of the elements in the
	// response. The elements in the response will not contain any aliases or
	// tags.
	List(ctx context.Context, request ListModelVersionsRequest) (*ListModelVersionsResponse, error)

	// Updates the specified model version.
	//
	// The caller must be a metastore admin or an owner of the parent registered
	// model. For the latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	//
	// Currently only the comment of the model version can be updated.
	Update(ctx context.Context, request UpdateModelVersionRequest) (*ModelVersionInfo, error)
}

// Online tables provide lower latency and higher QPS access to data from Delta
// tables.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type OnlineTablesService interface {

	// Create a new Online Table.
	Create(ctx context.Context, request CreateOnlineTableRequest) (*OnlineTable, error)

	// Delete an online table. Warning: This will delete all the data in the
	// online table. If the source Delta table was deleted or modified since
	// this Online Table was created, this will lose the data forever!
	Delete(ctx context.Context, request DeleteOnlineTableRequest) error

	// Get information about an existing online table and its status.
	Get(ctx context.Context, request GetOnlineTableRequest) (*OnlineTable, error)
}

// A monitor computes and monitors data or model quality metrics for a table
// over time. It generates metrics tables and a dashboard that you can use to
// monitor table health and set alerts. Most write operations require the user
// to be the owner of the table (or its parent schema or parent catalog).
// Viewing the dashboard, computed metrics, or monitor configuration only
// requires the user to have **SELECT** privileges on the table (along with
// **USE_SCHEMA** and **USE_CATALOG**).
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type QualityMonitorsService interface {

	// Cancels an already-initiated refresh job.
	CancelRefresh(ctx context.Context, request CancelRefreshRequest) error

	// Creates a new monitor for the specified table.
	//
	// The caller must either: 1. be an owner of the table's parent catalog,
	// have **USE_SCHEMA** on the table's parent schema, and have **SELECT**
	// access on the table 2. have **USE_CATALOG** on the table's parent
	// catalog, be an owner of the table's parent schema, and have **SELECT**
	// access on the table. 3. have the following permissions: - **USE_CATALOG**
	// on the table's parent catalog - **USE_SCHEMA** on the table's parent
	// schema - be an owner of the table.
	//
	// Workspace assets, such as the dashboard, will be created in the workspace
	// where this call was made.
	Create(ctx context.Context, request CreateMonitor) (*MonitorInfo, error)

	// Deletes a monitor for the specified table.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - be an owner of the table.
	//
	// Additionally, the call must be made from the workspace where the monitor
	// was created.
	//
	// Note that the metric tables and dashboard will not be deleted as part of
	// this call; those assets must be manually cleaned up (if desired).
	Delete(ctx context.Context, request DeleteQualityMonitorRequest) (*DeleteMonitorResponse, error)

	// Gets a monitor for the specified table.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema. 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - **SELECT** privilege on the table.
	//
	// The returned information includes configuration values, as well as
	// information on assets created by the monitor. Some information (e.g.,
	// dashboard) may be filtered out if the caller is in a different workspace
	// than where the monitor was created.
	Get(ctx context.Context, request GetQualityMonitorRequest) (*MonitorInfo, error)

	// Gets info about a specific monitor refresh using the given refresh ID.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - **SELECT** privilege on the table.
	//
	// Additionally, the call must be made from the workspace where the monitor
	// was created.
	GetRefresh(ctx context.Context, request GetRefreshRequest) (*MonitorRefreshInfo, error)

	// Gets an array containing the history of the most recent refreshes (up to
	// 25) for this table.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - **SELECT** privilege on the table.
	//
	// Additionally, the call must be made from the workspace where the monitor
	// was created.
	ListRefreshes(ctx context.Context, request ListRefreshesRequest) (*MonitorRefreshListResponse, error)

	// Regenerates the monitoring dashboard for the specified table.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - be an owner of the table
	//
	// The call must be made from the workspace where the monitor was created.
	// The dashboard will be regenerated in the assets directory that was
	// specified when the monitor was created.
	RegenerateDashboard(ctx context.Context, request RegenerateDashboardRequest) (*RegenerateDashboardResponse, error)

	// Queues a metric refresh on the monitor for the specified table. The
	// refresh will execute in the background.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - be an owner of the table
	//
	// Additionally, the call must be made from the workspace where the monitor
	// was created.
	RunRefresh(ctx context.Context, request RunRefreshRequest) (*MonitorRefreshInfo, error)

	// Updates a monitor for the specified table.
	//
	// The caller must either: 1. be an owner of the table's parent catalog 2.
	// have **USE_CATALOG** on the table's parent catalog and be an owner of the
	// table's parent schema 3. have the following permissions: -
	// **USE_CATALOG** on the table's parent catalog - **USE_SCHEMA** on the
	// table's parent schema - be an owner of the table.
	//
	// Additionally, the call must be made from the workspace where the monitor
	// was created, and the caller must be the original creator of the monitor.
	//
	// Certain configuration fields, such as output asset identifiers, cannot be
	// updated.
	Update(ctx context.Context, request UpdateMonitor) (*MonitorInfo, error)
}

// Databricks provides a hosted version of MLflow Model Registry in Unity
// Catalog. Models in Unity Catalog provide centralized access control,
// auditing, lineage, and discovery of ML models across Databricks workspaces.
//
// An MLflow registered model resides in the third layer of Unity Catalog’s
// three-level namespace. Registered models contain model versions, which
// correspond to actual ML models (MLflow models). Creating new model versions
// currently requires use of the MLflow Python client. Once model versions are
// created, you can load them for batch inference using MLflow Python client
// APIs, or deploy them for real-time serving using Databricks Model Serving.
//
// All operations on registered models and model versions require USE_CATALOG
// permissions on the enclosing catalog and USE_SCHEMA permissions on the
// enclosing schema. In addition, the following additional privileges are
// required for various operations:
//
// * To create a registered model, users must additionally have the CREATE_MODEL
// permission on the target schema. * To view registered model or model version
// metadata, model version data files, or invoke a model version, users must
// additionally have the EXECUTE permission on the registered model * To update
// registered model or model version tags, users must additionally have APPLY
// TAG permissions on the registered model * To update other registered model or
// model version metadata (comments, aliases) create a new model version, or
// update permissions on the registered model, users must be owners of the
// registered model.
//
// Note: The securable type for models is "FUNCTION". When using REST APIs (e.g.
// tagging, grants) that specify a securable type, use "FUNCTION" as the
// securable type.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type RegisteredModelsService interface {

	// Creates a new registered model in Unity Catalog.
	//
	// File storage for model versions in the registered model will be located
	// in the default location which is specified by the parent schema, or the
	// parent catalog, or the Metastore.
	//
	// For registered model creation to succeed, the user must satisfy the
	// following conditions: - The caller must be a metastore admin, or be the
	// owner of the parent catalog and schema, or have the **USE_CATALOG**
	// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
	// parent schema. - The caller must have the **CREATE MODEL** or **CREATE
	// FUNCTION** privilege on the parent schema.
	Create(ctx context.Context, request CreateRegisteredModelRequest) (*RegisteredModelInfo, error)

	// Deletes a registered model and all its model versions from the specified
	// parent catalog and schema.
	//
	// The caller must be a metastore admin or an owner of the registered model.
	// For the latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	Delete(ctx context.Context, request DeleteRegisteredModelRequest) error

	// Deletes a registered model alias.
	//
	// The caller must be a metastore admin or an owner of the registered model.
	// For the latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	DeleteAlias(ctx context.Context, request DeleteAliasRequest) error

	// Get a registered model.
	//
	// The caller must be a metastore admin or an owner of (or have the
	// **EXECUTE** privilege on) the registered model. For the latter case, the
	// caller must also be the owner or have the **USE_CATALOG** privilege on
	// the parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	Get(ctx context.Context, request GetRegisteredModelRequest) (*RegisteredModelInfo, error)

	// List registered models. You can list registered models under a particular
	// schema, or list all registered models in the current metastore.
	//
	// The returned models are filtered based on the privileges of the calling
	// user. For example, the metastore admin is able to list all the registered
	// models. A regular user needs to be the owner or have the **EXECUTE**
	// privilege on the registered model to recieve the registered models in the
	// response. For the latter case, the caller must also be the owner or have
	// the **USE_CATALOG** privilege on the parent catalog and the
	// **USE_SCHEMA** privilege on the parent schema.
	//
	// There is no guarantee of a specific ordering of the elements in the
	// response.
	List(ctx context.Context, request ListRegisteredModelsRequest) (*ListRegisteredModelsResponse, error)

	// Set an alias on the specified registered model.
	//
	// The caller must be a metastore admin or an owner of the registered model.
	// For the latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	SetAlias(ctx context.Context, request SetRegisteredModelAliasRequest) (*RegisteredModelAlias, error)

	// Updates the specified registered model.
	//
	// The caller must be a metastore admin or an owner of the registered model.
	// For the latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	//
	// Currently only the name, the owner or the comment of the registered model
	// can be updated.
	Update(ctx context.Context, request UpdateRegisteredModelRequest) (*RegisteredModelInfo, error)
}

// Unity Catalog enforces resource quotas on all securable objects, which limits
// the number of resources that can be created. Quotas are expressed in terms of
// a resource type and a parent (for example, tables per metastore or schemas
// per catalog). The resource quota APIs enable you to monitor your current
// usage and limits. For more information on resource quotas see the [Unity
// Catalog documentation].
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Unity Catalog documentation]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#resource-quotas
type ResourceQuotasService interface {

	// The GetQuota API returns usage information for a single resource quota,
	// defined as a child-parent pair. This API also refreshes the quota count
	// if it is out of date. Refreshes are triggered asynchronously. The updated
	// count might not be returned in the first call.
	GetQuota(ctx context.Context, request GetQuotaRequest) (*GetQuotaResponse, error)

	// ListQuotas returns all quota values under the metastore. There are no
	// SLAs on the freshness of the counts returned. This API does not trigger a
	// refresh of quota counts.
	ListQuotas(ctx context.Context, request ListQuotasRequest) (*ListQuotasResponse, error)
}

// A schema (also called a database) is the second layer of Unity Catalog’s
// three-level namespace. A schema organizes tables, views and functions. To
// access (or list) a table or view in a schema, users must have the USE_SCHEMA
// data permission on the schema and its parent catalog, and they must have the
// SELECT permission on the table or view.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SchemasService interface {

	// Creates a new schema for catalog in the Metatastore. The caller must be a
	// metastore admin, or have the **CREATE_SCHEMA** privilege in the parent
	// catalog.
	Create(ctx context.Context, request CreateSchema) (*SchemaInfo, error)

	// Deletes the specified schema from the parent catalog. The caller must be
	// the owner of the schema or an owner of the parent catalog.
	Delete(ctx context.Context, request DeleteSchemaRequest) error

	// Gets the specified schema within the metastore. The caller must be a
	// metastore admin, the owner of the schema, or a user that has the
	// **USE_SCHEMA** privilege on the schema.
	Get(ctx context.Context, request GetSchemaRequest) (*SchemaInfo, error)

	// Gets an array of schemas for a catalog in the metastore. If the caller is
	// the metastore admin or the owner of the parent catalog, all schemas for
	// the catalog will be retrieved. Otherwise, only schemas owned by the
	// caller (or for which the caller has the **USE_SCHEMA** privilege) will be
	// retrieved. There is no guarantee of a specific ordering of the elements
	// in the array.
	List(ctx context.Context, request ListSchemasRequest) (*ListSchemasResponse, error)

	// Updates a schema for a catalog. The caller must be the owner of the
	// schema or a metastore admin. If the caller is a metastore admin, only the
	// __owner__ field can be changed in the update. If the __name__ field must
	// be updated, the caller must be a metastore admin or have the
	// **CREATE_SCHEMA** privilege on the parent catalog.
	Update(ctx context.Context, request UpdateSchema) (*SchemaInfo, error)
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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type StorageCredentialsService interface {

	// Creates a new storage credential.
	//
	// The caller must be a metastore admin or have the
	// **CREATE_STORAGE_CREDENTIAL** privilege on the metastore.
	Create(ctx context.Context, request CreateStorageCredential) (*StorageCredentialInfo, error)

	// Deletes a storage credential from the metastore. The caller must be an
	// owner of the storage credential.
	Delete(ctx context.Context, request DeleteStorageCredentialRequest) error

	// Gets a storage credential from the metastore. The caller must be a
	// metastore admin, the owner of the storage credential, or have some
	// permission on the storage credential.
	Get(ctx context.Context, request GetStorageCredentialRequest) (*StorageCredentialInfo, error)

	// Gets an array of storage credentials (as __StorageCredentialInfo__
	// objects). The array is limited to only those storage credentials the
	// caller has permission to access. If the caller is a metastore admin,
	// retrieval of credentials is unrestricted. There is no guarantee of a
	// specific ordering of the elements in the array.
	List(ctx context.Context, request ListStorageCredentialsRequest) (*ListStorageCredentialsResponse, error)

	// Updates a storage credential on the metastore.
	//
	// The caller must be the owner of the storage credential or a metastore
	// admin. If the caller is a metastore admin, only the **owner** field can
	// be changed.
	Update(ctx context.Context, request UpdateStorageCredential) (*StorageCredentialInfo, error)

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

// A system schema is a schema that lives within the system catalog. A system
// schema may contain information about customer usage of Unity Catalog such as
// audit-logs, billing-logs, lineage information, etc.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SystemSchemasService interface {

	// Disables the system schema and removes it from the system catalog. The
	// caller must be an account admin or a metastore admin.
	Disable(ctx context.Context, request DisableRequest) error

	// Enables the system schema and adds it to the system catalog. The caller
	// must be an account admin or a metastore admin.
	Enable(ctx context.Context, request EnableRequest) error

	// Gets an array of system schemas for a metastore. The caller must be an
	// account admin or a metastore admin.
	List(ctx context.Context, request ListSystemSchemasRequest) (*ListSystemSchemasResponse, error)
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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TableConstraintsService interface {

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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TablesService interface {

	// Deletes a table from the specified parent catalog and schema. The caller
	// must be the owner of the parent catalog, have the **USE_CATALOG**
	// privilege on the parent catalog and be the owner of the parent schema, or
	// be the owner of the table and have the **USE_CATALOG** privilege on the
	// parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	Delete(ctx context.Context, request DeleteTableRequest) error

	// Gets if a table exists in the metastore for a specific catalog and
	// schema. The caller must satisfy one of the following requirements: * Be a
	// metastore admin * Be the owner of the parent catalog * Be the owner of
	// the parent schema and have the USE_CATALOG privilege on the parent
	// catalog * Have the **USE_CATALOG** privilege on the parent catalog and
	// the **USE_SCHEMA** privilege on the parent schema, and either be the
	// table owner or have the SELECT privilege on the table. * Have BROWSE
	// privilege on the parent catalog * Have BROWSE privilege on the parent
	// schema.
	Exists(ctx context.Context, request ExistsRequest) (*TableExistsResponse, error)

	// Gets a table from the metastore for a specific catalog and schema. The
	// caller must satisfy one of the following requirements: * Be a metastore
	// admin * Be the owner of the parent catalog * Be the owner of the parent
	// schema and have the USE_CATALOG privilege on the parent catalog * Have
	// the **USE_CATALOG** privilege on the parent catalog and the
	// **USE_SCHEMA** privilege on the parent schema, and either be the table
	// owner or have the SELECT privilege on the table.
	Get(ctx context.Context, request GetTableRequest) (*TableInfo, error)

	// Gets an array of all tables for the current metastore under the parent
	// catalog and schema. The caller must be a metastore admin or an owner of
	// (or have the **SELECT** privilege on) the table. For the latter case, the
	// caller must also be the owner or have the **USE_CATALOG** privilege on
	// the parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	// There is no guarantee of a specific ordering of the elements in the
	// array.
	List(ctx context.Context, request ListTablesRequest) (*ListTablesResponse, error)

	// Gets an array of summaries for tables for a schema and catalog within the
	// metastore. The table summaries returned are either:
	//
	// * summaries for tables (within the current metastore and parent catalog
	// and schema), when the user is a metastore admin, or: * summaries for
	// tables and schemas (within the current metastore and parent catalog) for
	// which the user has ownership or the **SELECT** privilege on the table and
	// ownership or **USE_SCHEMA** privilege on the schema, provided that the
	// user also has ownership or the **USE_CATALOG** privilege on the parent
	// catalog.
	//
	// There is no guarantee of a specific ordering of the elements in the
	// array.
	ListSummaries(ctx context.Context, request ListSummariesRequest) (*ListTableSummariesResponse, error)

	// Change the owner of the table. The caller must be the owner of the parent
	// catalog, have the **USE_CATALOG** privilege on the parent catalog and be
	// the owner of the parent schema, or be the owner of the table and have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	Update(ctx context.Context, request UpdateTableRequest) error
}

// Temporary Table Credentials refer to short-lived, downscoped credentials used
// to access cloud storage locationswhere table data is stored in Databricks.
// These credentials are employed to provide secure and time-limitedaccess to
// data in cloud environments such as AWS, Azure, and Google Cloud. Each cloud
// provider has its own typeof credentials: AWS uses temporary session tokens
// via AWS Security Token Service (STS), Azure utilizesShared Access Signatures
// (SAS) for its data storage services, and Google Cloud supports temporary
// credentialsthrough OAuth 2.0.Temporary table credentials ensure that data
// access is limited in scope and duration, reducing the risk ofunauthorized
// access or misuse. To use the temporary table credentials API, a metastore
// admin needs to enable the external_access_enabled flag (off by default) at
// the metastore level, and user needs to be granted the EXTERNAL USE SCHEMA
// permission at the schema level by catalog admin. Note that EXTERNAL USE
// SCHEMA is a schema level permission that can only be granted by catalog admin
// explicitly and is not included in schema ownership or ALL PRIVILEGES on the
// schema for security reason.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type TemporaryTableCredentialsService interface {

	// Get a short-lived credential for directly accessing the table data on
	// cloud storage. The metastore must have external_access_enabled flag set
	// to true (default false). The caller must have EXTERNAL_USE_SCHEMA
	// privilege on the parent schema and this privilege can only be granted by
	// catalog owners.
	GenerateTemporaryTableCredentials(ctx context.Context, request GenerateTemporaryTableCredentialRequest) (*GenerateTemporaryTableCredentialResponse, error)
}

// Volumes are a Unity Catalog (UC) capability for accessing, storing,
// governing, organizing and processing files. Use cases include running machine
// learning on unstructured data such as image, audio, video, or PDF files,
// organizing data sets during the data exploration stages in data science,
// working with libraries that require access to the local file system on
// cluster machines, storing library and config files of arbitrary formats such
// as .whl or .txt centrally and providing secure access across workspaces to
// it, or transforming and querying non-tabular data files in ETL.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type VolumesService interface {

	// Creates a new volume.
	//
	// The user could create either an external volume or a managed volume. An
	// external volume will be created in the specified external location, while
	// a managed volume will be located in the default location which is
	// specified by the parent schema, or the parent catalog, or the Metastore.
	//
	// For the volume creation to succeed, the user must satisfy following
	// conditions: - The caller must be a metastore admin, or be the owner of
	// the parent catalog and schema, or have the **USE_CATALOG** privilege on
	// the parent catalog and the **USE_SCHEMA** privilege on the parent schema.
	// - The caller must have **CREATE VOLUME** privilege on the parent schema.
	//
	// For an external volume, following conditions also need to satisfy - The
	// caller must have **CREATE EXTERNAL VOLUME** privilege on the external
	// location. - There are no other tables, nor volumes existing in the
	// specified storage location. - The specified storage location is not under
	// the location of other tables, nor volumes, or catalogs or schemas.
	Create(ctx context.Context, request CreateVolumeRequestContent) (*VolumeInfo, error)

	// Deletes a volume from the specified parent catalog and schema.
	//
	// The caller must be a metastore admin or an owner of the volume. For the
	// latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	Delete(ctx context.Context, request DeleteVolumeRequest) error

	// Gets an array of volumes for the current metastore under the parent
	// catalog and schema.
	//
	// The returned volumes are filtered based on the privileges of the calling
	// user. For example, the metastore admin is able to list all the volumes. A
	// regular user needs to be the owner or have the **READ VOLUME** privilege
	// on the volume to recieve the volumes in the response. For the latter
	// case, the caller must also be the owner or have the **USE_CATALOG**
	// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
	// parent schema.
	//
	// There is no guarantee of a specific ordering of the elements in the
	// array.
	List(ctx context.Context, request ListVolumesRequest) (*ListVolumesResponseContent, error)

	// Gets a volume from the metastore for a specific catalog and schema.
	//
	// The caller must be a metastore admin or an owner of (or have the **READ
	// VOLUME** privilege on) the volume. For the latter case, the caller must
	// also be the owner or have the **USE_CATALOG** privilege on the parent
	// catalog and the **USE_SCHEMA** privilege on the parent schema.
	Read(ctx context.Context, request ReadVolumeRequest) (*VolumeInfo, error)

	// Updates the specified volume under the specified parent catalog and
	// schema.
	//
	// The caller must be a metastore admin or an owner of the volume. For the
	// latter case, the caller must also be the owner or have the
	// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
	// privilege on the parent schema.
	//
	// Currently only the name, the owner or the comment of the volume could be
	// updated.
	Update(ctx context.Context, request UpdateVolumeRequestContent) (*VolumeInfo, error)
}

// A securable in Databricks can be configured as __OPEN__ or __ISOLATED__. An
// __OPEN__ securable can be accessed from any workspace, while an __ISOLATED__
// securable can only be accessed from a configured list of workspaces. This API
// allows you to configure (bind) securables to workspaces.
//
// NOTE: The __isolation_mode__ is configured for the securable itself (using
// its Update method) and the workspace bindings are only consulted when the
// securable's __isolation_mode__ is set to __ISOLATED__.
//
// A securable's workspace bindings can be configured by a metastore admin or
// the owner of the securable.
//
// The original path (/api/2.1/unity-catalog/workspace-bindings/catalogs/{name})
// is deprecated. Please use the new path
// (/api/2.1/unity-catalog/bindings/{securable_type}/{securable_name}) which
// introduces the ability to bind a securable in READ_ONLY mode (catalogs only).
//
// Securable types that support binding: - catalog - storage_credential -
// credential - external_location
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type WorkspaceBindingsService interface {

	// Gets workspace bindings of the catalog. The caller must be a metastore
	// admin or an owner of the catalog.
	Get(ctx context.Context, request GetWorkspaceBindingRequest) (*GetCatalogWorkspaceBindingsResponse, error)

	// Gets workspace bindings of the securable. The caller must be a metastore
	// admin or an owner of the securable.
	GetBindings(ctx context.Context, request GetBindingsRequest) (*GetWorkspaceBindingsResponse, error)

	// Updates workspace bindings of the catalog. The caller must be a metastore
	// admin or an owner of the catalog.
	Update(ctx context.Context, request UpdateWorkspaceBindings) (*UpdateCatalogWorkspaceBindingsResponse, error)

	// Updates workspace bindings of the securable. The caller must be a
	// metastore admin or an owner of the securable.
	UpdateBindings(ctx context.Context, request UpdateWorkspaceBindingsParameters) (*UpdateWorkspaceBindingsResponse, error)
}
