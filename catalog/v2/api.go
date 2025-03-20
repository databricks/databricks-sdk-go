// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Metastore Assignments, Account Metastores, Account Storage Credentials, Artifact Allowlists, Catalogs, Connections, Credentials, External Locations, Functions, Grants, Metastores, Model Versions, Online Tables, Quality Monitors, Registered Models, Resource Quotas, Schemas, Storage Credentials, System Schemas, Table Constraints, Tables, Temporary Table Credentials, Volumes, Workspace Bindings, etc.
package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// These APIs manage metastore assignments to a workspace.
type AccountMetastoreAssignmentsAPI struct {
	accountMetastoreAssignmentsImpl
}

// Delete a metastore assignment.
//
// Deletes a metastore assignment to a workspace, leaving the workspace with no
// metastore.
func (a *AccountMetastoreAssignmentsAPI) DeleteByWorkspaceIdAndMetastoreId(ctx context.Context, workspaceId int64, metastoreId string) (*DeleteResponse, error) {
	return a.accountMetastoreAssignmentsImpl.Delete(ctx, DeleteAccountMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
		MetastoreId: metastoreId,
	})
}

// Gets the metastore assignment for a workspace.
//
// Gets the metastore assignment, if any, for the workspace specified by ID. If
// the workspace is assigned a metastore, the mappig will be returned. If no
// metastore is assigned to the workspace, the assignment will not be found and
// a 404 returned.
func (a *AccountMetastoreAssignmentsAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*AccountsMetastoreAssignment, error) {
	return a.accountMetastoreAssignmentsImpl.Get(ctx, GetAccountMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get all workspaces assigned to a metastore.
//
// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *AccountMetastoreAssignmentsAPI) ListByMetastoreId(ctx context.Context, metastoreId string) (*ListAccountMetastoreAssignmentsResponse, error) {
	return a.accountMetastoreAssignmentsImpl.internalList(ctx, ListAccountMetastoreAssignmentsRequest{
		MetastoreId: metastoreId,
	})
}

// These APIs manage Unity Catalog metastores for an account. A metastore
// contains catalogs that can be associated with workspaces
type AccountMetastoresAPI struct {
	accountMetastoresImpl
}

// Delete a metastore.
//
// Deletes a Unity Catalog metastore for an account, both specified by ID.
func (a *AccountMetastoresAPI) DeleteByMetastoreId(ctx context.Context, metastoreId string) (*DeleteResponse, error) {
	return a.accountMetastoresImpl.Delete(ctx, DeleteAccountMetastoreRequest{
		MetastoreId: metastoreId,
	})
}

// Get a metastore.
//
// Gets a Unity Catalog metastore from an account, both specified by ID.
func (a *AccountMetastoresAPI) GetByMetastoreId(ctx context.Context, metastoreId string) (*AccountsMetastoreInfo, error) {
	return a.accountMetastoresImpl.Get(ctx, GetAccountMetastoreRequest{
		MetastoreId: metastoreId,
	})
}

// These APIs manage storage credentials for a particular metastore.
type AccountStorageCredentialsAPI struct {
	accountStorageCredentialsImpl
}

// Delete a storage credential.
//
// Deletes a storage credential from the metastore. The caller must be an owner
// of the storage credential.
func (a *AccountStorageCredentialsAPI) DeleteByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) (*DeleteResponse, error) {
	return a.accountStorageCredentialsImpl.Delete(ctx, DeleteAccountStorageCredentialRequest{
		MetastoreId:           metastoreId,
		StorageCredentialName: storageCredentialName,
	})
}

// Gets the named storage credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have a level of privilege on
// the storage credential.
func (a *AccountStorageCredentialsAPI) GetByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) (*AccountsStorageCredentialInfo, error) {
	return a.accountStorageCredentialsImpl.Get(ctx, GetAccountStorageCredentialRequest{
		MetastoreId:           metastoreId,
		StorageCredentialName: storageCredentialName,
	})
}

// Get all storage credentials assigned to a metastore.
//
// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *AccountStorageCredentialsAPI) ListByMetastoreId(ctx context.Context, metastoreId string) (*ListAccountStorageCredentialsResponse, error) {
	return a.accountStorageCredentialsImpl.internalList(ctx, ListAccountStorageCredentialsRequest{
		MetastoreId: metastoreId,
	})
}

// In Databricks Runtime 13.3 and above, you can add libraries and init scripts
// to the `allowlist` in UC so that users can leverage these artifacts on
// compute configured with shared access mode.
type ArtifactAllowlistsAPI struct {
	artifactAllowlistsImpl
}

// Get an artifact allowlist.
//
// Get the artifact allowlist of a certain artifact type. The caller must be a
// metastore admin or have the **MANAGE ALLOWLIST** privilege on the metastore.
func (a *ArtifactAllowlistsAPI) GetByArtifactType(ctx context.Context, artifactType ArtifactType) (*ArtifactAllowlistInfo, error) {
	return a.artifactAllowlistsImpl.Get(ctx, GetArtifactAllowlistRequest{
		ArtifactType: artifactType,
	})
}

// A catalog is the first layer of Unity Catalog’s three-level namespace.
// It’s used to organize your data assets. Users can see all catalogs on which
// they have been assigned the USE_CATALOG data permission.
//
// In Unity Catalog, admins and data stewards manage users and their access to
// data centrally across all of the workspaces in a Databricks account. Users in
// different workspaces can share access to the same data, depending on
// privileges granted centrally in Unity Catalog.
type CatalogsAPI struct {
	catalogsImpl
}

// Delete a catalog.
//
// Deletes the catalog that matches the supplied name. The caller must be a
// metastore admin or the owner of the catalog.
func (a *CatalogsAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.catalogsImpl.Delete(ctx, DeleteCatalogRequest{
		Name: name,
	})
}

// Get a catalog.
//
// Gets the specified catalog in a metastore. The caller must be a metastore
// admin, the owner of the catalog, or a user that has the **USE_CATALOG**
// privilege set for their account.
func (a *CatalogsAPI) GetByName(ctx context.Context, name string) (*CatalogInfo, error) {
	return a.catalogsImpl.Get(ctx, GetCatalogRequest{
		Name: name,
	})
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
type ConnectionsAPI struct {
	connectionsImpl
}

// Delete a connection.
//
// Deletes the connection that matches the supplied name.
func (a *ConnectionsAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.connectionsImpl.Delete(ctx, DeleteConnectionRequest{
		Name: name,
	})
}

// Get a connection.
//
// Gets a connection from it's name.
func (a *ConnectionsAPI) GetByName(ctx context.Context, name string) (*ConnectionInfo, error) {
	return a.connectionsImpl.Get(ctx, GetConnectionRequest{
		Name: name,
	})
}

// ConnectionInfoNameToFullNameMap calls [ConnectionsAPI.ListAll] and creates a map of results with [ConnectionInfo].Name as key and [ConnectionInfo].FullName as value.
//
// Returns an error if there's more than one [ConnectionInfo] with the same .Name.
//
// Note: All [ConnectionInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ConnectionsAPI) ConnectionInfoNameToFullNameMap(ctx context.Context, request ListConnectionsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.FullName
	}
	return mapping, nil
}

// A credential represents an authentication and authorization mechanism for
// accessing services on your cloud tenant. Each credential is subject to Unity
// Catalog access-control policies that control which users and groups can
// access the credential.
//
// To create credentials, you must be a Databricks account admin or have the
// `CREATE SERVICE CREDENTIAL` privilege. The user who creates the credential
// can delegate ownership to another user or group to manage permissions on it.
type CredentialsAPI struct {
	credentialsImpl
}

// Delete a credential.
//
// Deletes a service or storage credential from the metastore. The caller must
// be an owner of the credential.
func (a *CredentialsAPI) DeleteCredentialByNameArg(ctx context.Context, nameArg string) (*DeleteCredentialResponse, error) {
	return a.credentialsImpl.DeleteCredential(ctx, DeleteCredentialRequest{
		NameArg: nameArg,
	})
}

// Get a credential.
//
// Gets a service or storage credential from the metastore. The caller must be a
// metastore admin, the owner of the credential, or have any permission on the
// credential.
func (a *CredentialsAPI) GetCredentialByNameArg(ctx context.Context, nameArg string) (*CredentialInfo, error) {
	return a.credentialsImpl.GetCredential(ctx, GetCredentialRequest{
		NameArg: nameArg,
	})
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
type ExternalLocationsAPI struct {
	externalLocationsImpl
}

// Delete an external location.
//
// Deletes the specified external location from the metastore. The caller must
// be the owner of the external location.
func (a *ExternalLocationsAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.externalLocationsImpl.Delete(ctx, DeleteExternalLocationRequest{
		Name: name,
	})
}

// Get an external location.
//
// Gets an external location from the metastore. The caller must be either a
// metastore admin, the owner of the external location, or a user that has some
// privilege on the external location.
func (a *ExternalLocationsAPI) GetByName(ctx context.Context, name string) (*ExternalLocationInfo, error) {
	return a.externalLocationsImpl.Get(ctx, GetExternalLocationRequest{
		Name: name,
	})
}

// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
//
// The function implementation can be any SQL expression or Query, and it can be
// invoked wherever a table reference is allowed in a query. In Unity Catalog, a
// function resides at the same level as a table, so it can be referenced with
// the form __catalog_name__.__schema_name__.__function_name__.
type FunctionsAPI struct {
	functionsImpl
}

// Delete a function.
//
// Deletes the function that matches the supplied name. For the deletion to
// succeed, the user must satisfy one of the following conditions: - Is the
// owner of the function's parent catalog - Is the owner of the function's
// parent schema and have the **USE_CATALOG** privilege on its parent catalog -
// Is the owner of the function itself and have both the **USE_CATALOG**
// privilege on its parent catalog and the **USE_SCHEMA** privilege on its
// parent schema
func (a *FunctionsAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.functionsImpl.Delete(ctx, DeleteFunctionRequest{
		Name: name,
	})
}

// Get a function.
//
// Gets a function from within a parent catalog and schema. For the fetch to
// succeed, the user must satisfy one of the following requirements: - Is a
// metastore admin - Is an owner of the function's parent catalog - Have the
// **USE_CATALOG** privilege on the function's parent catalog and be the owner
// of the function - Have the **USE_CATALOG** privilege on the function's parent
// catalog, the **USE_SCHEMA** privilege on the function's parent schema, and
// the **EXECUTE** privilege on the function itself
func (a *FunctionsAPI) GetByName(ctx context.Context, name string) (*FunctionInfo, error) {
	return a.functionsImpl.Get(ctx, GetFunctionRequest{
		Name: name,
	})
}

// FunctionInfoNameToFullNameMap calls [FunctionsAPI.ListAll] and creates a map of results with [FunctionInfo].Name as key and [FunctionInfo].FullName as value.
//
// Returns an error if there's more than one [FunctionInfo] with the same .Name.
//
// Note: All [FunctionInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *FunctionsAPI) FunctionInfoNameToFullNameMap(ctx context.Context, request ListFunctionsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.FullName
	}
	return mapping, nil
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
type GrantsAPI struct {
	grantsImpl
}

// Get permissions.
//
// Gets the permissions for a securable.
func (a *GrantsAPI) GetBySecurableTypeAndFullName(ctx context.Context, securableType SecurableType, fullName string) (*PermissionsList, error) {
	return a.grantsImpl.Get(ctx, GetGrantRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

// Get effective permissions.
//
// Gets the effective permissions for a securable.
func (a *GrantsAPI) GetEffectiveBySecurableTypeAndFullName(ctx context.Context, securableType SecurableType, fullName string) (*EffectivePermissionsList, error) {
	return a.grantsImpl.GetEffective(ctx, GetEffectiveRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
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
type MetastoresAPI struct {
	metastoresImpl
}

// Delete a metastore.
//
// Deletes a metastore. The caller must be a metastore admin.
func (a *MetastoresAPI) DeleteById(ctx context.Context, id string) (*DeleteResponse, error) {
	return a.metastoresImpl.Delete(ctx, DeleteMetastoreRequest{
		Id: id,
	})
}

// Get a metastore.
//
// Gets a metastore that matches the supplied ID. The caller must be a metastore
// admin to retrieve this info.
func (a *MetastoresAPI) GetById(ctx context.Context, id string) (*MetastoreInfo, error) {
	return a.metastoresImpl.Get(ctx, GetMetastoreRequest{
		Id: id,
	})
}

// MetastoreInfoNameToMetastoreIdMap calls [MetastoresAPI.ListAll] and creates a map of results with [MetastoreInfo].Name as key and [MetastoreInfo].MetastoreId as value.
//
// Returns an error if there's more than one [MetastoreInfo] with the same .Name.
//
// Note: All [MetastoreInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *MetastoresAPI) MetastoreInfoNameToMetastoreIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.MetastoreId
	}
	return mapping, nil
}

// GetByName calls [MetastoresAPI.MetastoreInfoNameToMetastoreIdMap] and returns a single [MetastoreInfo].
//
// Returns an error if there's more than one [MetastoreInfo] with the same .Name.
//
// Note: All [MetastoreInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *MetastoresAPI) GetByName(ctx context.Context, name string) (*MetastoreInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]MetastoreInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("MetastoreInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of MetastoreInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Delete an assignment.
//
// Deletes a metastore assignment. The caller must be an account administrator.
func (a *MetastoresAPI) UnassignByWorkspaceId(ctx context.Context, workspaceId int64) (*UnassignResponse, error) {
	return a.metastoresImpl.Unassign(ctx, UnassignRequest{
		WorkspaceId: workspaceId,
	})
}

// Databricks provides a hosted version of MLflow Model Registry in Unity
// Catalog. Models in Unity Catalog provide centralized access control,
// auditing, lineage, and discovery of ML models across Databricks workspaces.
//
// This API reference documents the REST endpoints for managing model versions
// in Unity Catalog. For more details, see the [registered models API
// docs](/api/workspace/registeredmodels).
type ModelVersionsAPI struct {
	modelVersionsImpl
}

// Delete a Model Version.
//
// Deletes a model version from the specified registered model. Any aliases
// assigned to the model version will also be deleted.
//
// The caller must be a metastore admin or an owner of the parent registered
// model. For the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
func (a *ModelVersionsAPI) DeleteByFullNameAndVersion(ctx context.Context, fullName string, version int) (*DeleteResponse, error) {
	return a.modelVersionsImpl.Delete(ctx, DeleteModelVersionRequest{
		FullName: fullName,
		Version:  version,
	})
}

// Get a Model Version.
//
// Get a model version.
//
// The caller must be a metastore admin or an owner of (or have the **EXECUTE**
// privilege on) the parent registered model. For the latter case, the caller
// must also be the owner or have the **USE_CATALOG** privilege on the parent
// catalog and the **USE_SCHEMA** privilege on the parent schema.
func (a *ModelVersionsAPI) GetByFullNameAndVersion(ctx context.Context, fullName string, version int) (*ModelVersionInfo, error) {
	return a.modelVersionsImpl.Get(ctx, GetModelVersionRequest{
		FullName: fullName,
		Version:  version,
	})
}

// Get Model Version By Alias.
//
// Get a model version by alias.
//
// The caller must be a metastore admin or an owner of (or have the **EXECUTE**
// privilege on) the registered model. For the latter case, the caller must also
// be the owner or have the **USE_CATALOG** privilege on the parent catalog and
// the **USE_SCHEMA** privilege on the parent schema.
func (a *ModelVersionsAPI) GetByAliasByFullNameAndAlias(ctx context.Context, fullName string, alias string) (*ModelVersionInfo, error) {
	return a.modelVersionsImpl.GetByAlias(ctx, GetByAliasRequest{
		FullName: fullName,
		Alias:    alias,
	})
}

// List Model Versions.
//
// List model versions. You can list model versions under a particular schema,
// or list all model versions in the current metastore.
//
// The returned models are filtered based on the privileges of the calling user.
// For example, the metastore admin is able to list all the model versions. A
// regular user needs to be the owner or have the **EXECUTE** privilege on the
// parent registered model to recieve the model versions in the response. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
//
// There is no guarantee of a specific ordering of the elements in the response.
// The elements in the response will not contain any aliases or tags.
func (a *ModelVersionsAPI) ListByFullName(ctx context.Context, fullName string) (*ListModelVersionsResponse, error) {
	return a.modelVersionsImpl.internalList(ctx, ListModelVersionsRequest{
		FullName: fullName,
	})
}

// Online tables provide lower latency and higher QPS access to data from Delta
// tables.
type OnlineTablesAPI struct {
	onlineTablesImpl
}

// Delete an Online Table.
//
// Delete an online table. Warning: This will delete all the data in the online
// table. If the source Delta table was deleted or modified since this Online
// Table was created, this will lose the data forever!
func (a *OnlineTablesAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.onlineTablesImpl.Delete(ctx, DeleteOnlineTableRequest{
		Name: name,
	})
}

// Get an Online Table.
//
// Get information about an existing online table and its status.
func (a *OnlineTablesAPI) GetByName(ctx context.Context, name string) (*OnlineTable, error) {
	return a.onlineTablesImpl.Get(ctx, GetOnlineTableRequest{
		Name: name,
	})
}

// A monitor computes and monitors data or model quality metrics for a table
// over time. It generates metrics tables and a dashboard that you can use to
// monitor table health and set alerts.
//
// Most write operations require the user to be the owner of the table (or its
// parent schema or parent catalog). Viewing the dashboard, computed metrics, or
// monitor configuration only requires the user to have **SELECT** privileges on
// the table (along with **USE_SCHEMA** and **USE_CATALOG**).
type QualityMonitorsAPI struct {
	qualityMonitorsImpl
}

// Delete a table monitor.
//
// Deletes a monitor for the specified table.
//
// The caller must either: 1. be an owner of the table's parent catalog 2. have
// **USE_CATALOG** on the table's parent catalog and be an owner of the table's
// parent schema 3. have the following permissions: - **USE_CATALOG** on the
// table's parent catalog - **USE_SCHEMA** on the table's parent schema - be an
// owner of the table.
//
// Additionally, the call must be made from the workspace where the monitor was
// created.
//
// Note that the metric tables and dashboard will not be deleted as part of this
// call; those assets must be manually cleaned up (if desired).
func (a *QualityMonitorsAPI) DeleteByTableName(ctx context.Context, tableName string) (*DeleteResponse, error) {
	return a.qualityMonitorsImpl.Delete(ctx, DeleteQualityMonitorRequest{
		TableName: tableName,
	})
}

// Get a table monitor.
//
// Gets a monitor for the specified table.
//
// The caller must either: 1. be an owner of the table's parent catalog 2. have
// **USE_CATALOG** on the table's parent catalog and be an owner of the table's
// parent schema. 3. have the following permissions: - **USE_CATALOG** on the
// table's parent catalog - **USE_SCHEMA** on the table's parent schema -
// **SELECT** privilege on the table.
//
// The returned information includes configuration values, as well as
// information on assets created by the monitor. Some information (e.g.,
// dashboard) may be filtered out if the caller is in a different workspace than
// where the monitor was created.
func (a *QualityMonitorsAPI) GetByTableName(ctx context.Context, tableName string) (*MonitorInfo, error) {
	return a.qualityMonitorsImpl.Get(ctx, GetQualityMonitorRequest{
		TableName: tableName,
	})
}

// Get refresh.
//
// Gets info about a specific monitor refresh using the given refresh ID.
//
// The caller must either: 1. be an owner of the table's parent catalog 2. have
// **USE_CATALOG** on the table's parent catalog and be an owner of the table's
// parent schema 3. have the following permissions: - **USE_CATALOG** on the
// table's parent catalog - **USE_SCHEMA** on the table's parent schema -
// **SELECT** privilege on the table.
//
// Additionally, the call must be made from the workspace where the monitor was
// created.
func (a *QualityMonitorsAPI) GetRefreshByTableNameAndRefreshId(ctx context.Context, tableName string, refreshId string) (*MonitorRefreshInfo, error) {
	return a.qualityMonitorsImpl.GetRefresh(ctx, GetRefreshRequest{
		TableName: tableName,
		RefreshId: refreshId,
	})
}

// List refreshes.
//
// Gets an array containing the history of the most recent refreshes (up to 25)
// for this table.
//
// The caller must either: 1. be an owner of the table's parent catalog 2. have
// **USE_CATALOG** on the table's parent catalog and be an owner of the table's
// parent schema 3. have the following permissions: - **USE_CATALOG** on the
// table's parent catalog - **USE_SCHEMA** on the table's parent schema -
// **SELECT** privilege on the table.
//
// Additionally, the call must be made from the workspace where the monitor was
// created.
func (a *QualityMonitorsAPI) ListRefreshesByTableName(ctx context.Context, tableName string) (*MonitorRefreshListResponse, error) {
	return a.qualityMonitorsImpl.ListRefreshes(ctx, ListRefreshesRequest{
		TableName: tableName,
	})
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
type RegisteredModelsAPI struct {
	registeredModelsImpl
}

// Delete a Registered Model.
//
// Deletes a registered model and all its model versions from the specified
// parent catalog and schema.
//
// The caller must be a metastore admin or an owner of the registered model. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
func (a *RegisteredModelsAPI) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
	return a.registeredModelsImpl.Delete(ctx, DeleteRegisteredModelRequest{
		FullName: fullName,
	})
}

// Delete a Registered Model Alias.
//
// Deletes a registered model alias.
//
// The caller must be a metastore admin or an owner of the registered model. For
// the latter case, the caller must also be the owner or have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema.
func (a *RegisteredModelsAPI) DeleteAliasByFullNameAndAlias(ctx context.Context, fullName string, alias string) (*DeleteAliasResponse, error) {
	return a.registeredModelsImpl.DeleteAlias(ctx, DeleteAliasRequest{
		FullName: fullName,
		Alias:    alias,
	})
}

// Get a Registered Model.
//
// Get a registered model.
//
// The caller must be a metastore admin or an owner of (or have the **EXECUTE**
// privilege on) the registered model. For the latter case, the caller must also
// be the owner or have the **USE_CATALOG** privilege on the parent catalog and
// the **USE_SCHEMA** privilege on the parent schema.
func (a *RegisteredModelsAPI) GetByFullName(ctx context.Context, fullName string) (*RegisteredModelInfo, error) {
	return a.registeredModelsImpl.Get(ctx, GetRegisteredModelRequest{
		FullName: fullName,
	})
}

// RegisteredModelInfoNameToFullNameMap calls [RegisteredModelsAPI.ListAll] and creates a map of results with [RegisteredModelInfo].Name as key and [RegisteredModelInfo].FullName as value.
//
// Returns an error if there's more than one [RegisteredModelInfo] with the same .Name.
//
// Note: All [RegisteredModelInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RegisteredModelsAPI) RegisteredModelInfoNameToFullNameMap(ctx context.Context, request ListRegisteredModelsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.FullName
	}
	return mapping, nil
}

// GetByName calls [RegisteredModelsAPI.RegisteredModelInfoNameToFullNameMap] and returns a single [RegisteredModelInfo].
//
// Returns an error if there's more than one [RegisteredModelInfo] with the same .Name.
//
// Note: All [RegisteredModelInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *RegisteredModelsAPI) GetByName(ctx context.Context, name string) (*RegisteredModelInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListRegisteredModelsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]RegisteredModelInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("RegisteredModelInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of RegisteredModelInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Unity Catalog enforces resource quotas on all securable objects, which limits
// the number of resources that can be created. Quotas are expressed in terms of
// a resource type and a parent (for example, tables per metastore or schemas
// per catalog). The resource quota APIs enable you to monitor your current
// usage and limits. For more information on resource quotas see the [Unity
// Catalog documentation].
//
// [Unity Catalog documentation]: https://docs.databricks.com/en/data-governance/unity-catalog/index.html#resource-quotas
type ResourceQuotasAPI struct {
	resourceQuotasImpl
}

// Get information for a single resource quota.
//
// The GetQuota API returns usage information for a single resource quota,
// defined as a child-parent pair. This API also refreshes the quota count if it
// is out of date. Refreshes are triggered asynchronously. The updated count
// might not be returned in the first call.
func (a *ResourceQuotasAPI) GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName(ctx context.Context, parentSecurableType string, parentFullName string, quotaName string) (*GetQuotaResponse, error) {
	return a.resourceQuotasImpl.GetQuota(ctx, GetQuotaRequest{
		ParentSecurableType: parentSecurableType,
		ParentFullName:      parentFullName,
		QuotaName:           quotaName,
	})
}

// A schema (also called a database) is the second layer of Unity Catalog’s
// three-level namespace. A schema organizes tables, views and functions. To
// access (or list) a table or view in a schema, users must have the USE_SCHEMA
// data permission on the schema and its parent catalog, and they must have the
// SELECT permission on the table or view.
type SchemasAPI struct {
	schemasImpl
}

// Delete a schema.
//
// Deletes the specified schema from the parent catalog. The caller must be the
// owner of the schema or an owner of the parent catalog.
func (a *SchemasAPI) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
	return a.schemasImpl.Delete(ctx, DeleteSchemaRequest{
		FullName: fullName,
	})
}

// Get a schema.
//
// Gets the specified schema within the metastore. The caller must be a
// metastore admin, the owner of the schema, or a user that has the
// **USE_SCHEMA** privilege on the schema.
func (a *SchemasAPI) GetByFullName(ctx context.Context, fullName string) (*SchemaInfo, error) {
	return a.schemasImpl.Get(ctx, GetSchemaRequest{
		FullName: fullName,
	})
}

// SchemaInfoNameToFullNameMap calls [SchemasAPI.ListAll] and creates a map of results with [SchemaInfo].Name as key and [SchemaInfo].FullName as value.
//
// Returns an error if there's more than one [SchemaInfo] with the same .Name.
//
// Note: All [SchemaInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SchemasAPI) SchemaInfoNameToFullNameMap(ctx context.Context, request ListSchemasRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.FullName
	}
	return mapping, nil
}

// GetByName calls [SchemasAPI.SchemaInfoNameToFullNameMap] and returns a single [SchemaInfo].
//
// Returns an error if there's more than one [SchemaInfo] with the same .Name.
//
// Note: All [SchemaInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *SchemasAPI) GetByName(ctx context.Context, name string) (*SchemaInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListSchemasRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]SchemaInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("SchemaInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of SchemaInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
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
type StorageCredentialsAPI struct {
	storageCredentialsImpl
}

// Delete a credential.
//
// Deletes a storage credential from the metastore. The caller must be an owner
// of the storage credential.
func (a *StorageCredentialsAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.storageCredentialsImpl.Delete(ctx, DeleteStorageCredentialRequest{
		Name: name,
	})
}

// Get a credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have some permission on the
// storage credential.
func (a *StorageCredentialsAPI) GetByName(ctx context.Context, name string) (*StorageCredentialInfo, error) {
	return a.storageCredentialsImpl.Get(ctx, GetStorageCredentialRequest{
		Name: name,
	})
}

// StorageCredentialInfoNameToIdMap calls [StorageCredentialsAPI.ListAll] and creates a map of results with [StorageCredentialInfo].Name as key and [StorageCredentialInfo].Id as value.
//
// Returns an error if there's more than one [StorageCredentialInfo] with the same .Name.
//
// Note: All [StorageCredentialInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *StorageCredentialsAPI) StorageCredentialInfoNameToIdMap(ctx context.Context, request ListStorageCredentialsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// A system schema is a schema that lives within the system catalog. A system
// schema may contain information about customer usage of Unity Catalog such as
// audit-logs, billing-logs, lineage information, etc.
type SystemSchemasAPI struct {
	systemSchemasImpl
}

// Disable a system schema.
//
// Disables the system schema and removes it from the system catalog. The caller
// must be an account admin or a metastore admin.
func (a *SystemSchemasAPI) DisableByMetastoreIdAndSchemaName(ctx context.Context, metastoreId string, schemaName string) (*DisableResponse, error) {
	return a.systemSchemasImpl.Disable(ctx, DisableRequest{
		MetastoreId: metastoreId,
		SchemaName:  schemaName,
	})
}

// List system schemas.
//
// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
func (a *SystemSchemasAPI) ListByMetastoreId(ctx context.Context, metastoreId string) (*ListSystemSchemasResponse, error) {
	return a.systemSchemasImpl.internalList(ctx, ListSystemSchemasRequest{
		MetastoreId: metastoreId,
	})
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
type TableConstraintsAPI struct {
	tableConstraintsImpl
}

// Delete a table constraint.
//
// Deletes a table constraint.
//
// For the table constraint deletion to succeed, the user must satisfy both of
// these conditions: - the user must have the **USE_CATALOG** privilege on the
// table's parent catalog, the **USE_SCHEMA** privilege on the table's parent
// schema, and be the owner of the table. - if __cascade__ argument is **true**,
// the user must have the following permissions on all of the child tables: the
// **USE_CATALOG** privilege on the table's catalog, the **USE_SCHEMA**
// privilege on the table's schema, and be the owner of the table.
func (a *TableConstraintsAPI) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
	return a.tableConstraintsImpl.Delete(ctx, DeleteTableConstraintRequest{
		FullName: fullName,
	})
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
type TablesAPI struct {
	tablesImpl
}

// Delete a table.
//
// Deletes a table from the specified parent catalog and schema. The caller must
// be the owner of the parent catalog, have the **USE_CATALOG** privilege on the
// parent catalog and be the owner of the parent schema, or be the owner of the
// table and have the **USE_CATALOG** privilege on the parent catalog and the
// **USE_SCHEMA** privilege on the parent schema.
func (a *TablesAPI) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
	return a.tablesImpl.Delete(ctx, DeleteTableRequest{
		FullName: fullName,
	})
}

// Get boolean reflecting if table exists.
//
// Gets if a table exists in the metastore for a specific catalog and schema.
// The caller must satisfy one of the following requirements: * Be a metastore
// admin * Be the owner of the parent catalog * Be the owner of the parent
// schema and have the USE_CATALOG privilege on the parent catalog * Have the
// **USE_CATALOG** privilege on the parent catalog and the **USE_SCHEMA**
// privilege on the parent schema, and either be the table owner or have the
// SELECT privilege on the table. * Have BROWSE privilege on the parent catalog
// * Have BROWSE privilege on the parent schema.
func (a *TablesAPI) ExistsByFullName(ctx context.Context, fullName string) (*TableExistsResponse, error) {
	return a.tablesImpl.Exists(ctx, ExistsRequest{
		FullName: fullName,
	})
}

// Get a table.
//
// Gets a table from the metastore for a specific catalog and schema. The caller
// must satisfy one of the following requirements: * Be a metastore admin * Be
// the owner of the parent catalog * Be the owner of the parent schema and have
// the USE_CATALOG privilege on the parent catalog * Have the **USE_CATALOG**
// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
// parent schema, and either be the table owner or have the SELECT privilege on
// the table.
func (a *TablesAPI) GetByFullName(ctx context.Context, fullName string) (*TableInfo, error) {
	return a.tablesImpl.Get(ctx, GetTableRequest{
		FullName: fullName,
	})
}

// TableInfoNameToTableIdMap calls [TablesAPI.ListAll] and creates a map of results with [TableInfo].Name as key and [TableInfo].TableId as value.
//
// Returns an error if there's more than one [TableInfo] with the same .Name.
//
// Note: All [TableInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TablesAPI) TableInfoNameToTableIdMap(ctx context.Context, request ListTablesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.TableId
	}
	return mapping, nil
}

// GetByName calls [TablesAPI.TableInfoNameToTableIdMap] and returns a single [TableInfo].
//
// Returns an error if there's more than one [TableInfo] with the same .Name.
//
// Note: All [TableInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *TablesAPI) GetByName(ctx context.Context, name string) (*TableInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListTablesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]TableInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("TableInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of TableInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
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
type TemporaryTableCredentialsAPI struct {
	temporaryTableCredentialsImpl
}

// Volumes are a Unity Catalog (UC) capability for accessing, storing,
// governing, organizing and processing files. Use cases include running machine
// learning on unstructured data such as image, audio, video, or PDF files,
// organizing data sets during the data exploration stages in data science,
// working with libraries that require access to the local file system on
// cluster machines, storing library and config files of arbitrary formats such
// as .whl or .txt centrally and providing secure access across workspaces to
// it, or transforming and querying non-tabular data files in ETL.
type VolumesAPI struct {
	volumesImpl
}

// Delete a Volume.
//
// Deletes a volume from the specified parent catalog and schema.
//
// The caller must be a metastore admin or an owner of the volume. For the
// latter case, the caller must also be the owner or have the **USE_CATALOG**
// privilege on the parent catalog and the **USE_SCHEMA** privilege on the
// parent schema.
func (a *VolumesAPI) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.volumesImpl.Delete(ctx, DeleteVolumeRequest{
		Name: name,
	})
}

// VolumeInfoNameToVolumeIdMap calls [VolumesAPI.ListAll] and creates a map of results with [VolumeInfo].Name as key and [VolumeInfo].VolumeId as value.
//
// Returns an error if there's more than one [VolumeInfo] with the same .Name.
//
// Note: All [VolumeInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *VolumesAPI) VolumeInfoNameToVolumeIdMap(ctx context.Context, request ListVolumesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.VolumeId
	}
	return mapping, nil
}

// GetByName calls [VolumesAPI.VolumeInfoNameToVolumeIdMap] and returns a single [VolumeInfo].
//
// Returns an error if there's more than one [VolumeInfo] with the same .Name.
//
// Note: All [VolumeInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *VolumesAPI) GetByName(ctx context.Context, name string) (*VolumeInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListVolumesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]VolumeInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("VolumeInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of VolumeInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Get a Volume.
//
// Gets a volume from the metastore for a specific catalog and schema.
//
// The caller must be a metastore admin or an owner of (or have the **READ
// VOLUME** privilege on) the volume. For the latter case, the caller must also
// be the owner or have the **USE_CATALOG** privilege on the parent catalog and
// the **USE_SCHEMA** privilege on the parent schema.
func (a *VolumesAPI) ReadByName(ctx context.Context, name string) (*VolumeInfo, error) {
	return a.volumesImpl.Read(ctx, ReadVolumeRequest{
		Name: name,
	})
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
// external_location
type WorkspaceBindingsAPI struct {
	workspaceBindingsImpl
}

// Get catalog workspace bindings.
//
// Gets workspace bindings of the catalog. The caller must be a metastore admin
// or an owner of the catalog.
func (a *WorkspaceBindingsAPI) GetByName(ctx context.Context, name string) (*CurrentWorkspaceBindings, error) {
	return a.workspaceBindingsImpl.Get(ctx, GetWorkspaceBindingRequest{
		Name: name,
	})
}

// Get securable workspace bindings.
//
// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
func (a *WorkspaceBindingsAPI) GetBindingsBySecurableTypeAndSecurableName(ctx context.Context, securableType GetBindingsSecurableType, securableName string) (*WorkspaceBindingsResponse, error) {
	return a.workspaceBindingsImpl.internalGetBindings(ctx, GetBindingsRequest{
		SecurableType: securableType,
		SecurableName: securableName,
	})
}
