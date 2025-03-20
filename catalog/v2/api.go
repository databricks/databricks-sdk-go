// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Account Metastore Assignments, Account Metastores, Account Storage Credentials, Artifact Allowlists, Catalogs, Connections, Credentials, External Locations, Functions, Grants, Metastores, Model Versions, Online Tables, Quality Monitors, Registered Models, Resource Quotas, Schemas, Storage Credentials, System Schemas, Table Constraints, Tables, Temporary Table Credentials, Volumes, Workspace Bindings, etc.
package catalog

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type accountMetastoreAssignmentsBaseClient struct {
	accountMetastoreAssignmentsImpl
}

// Delete a metastore assignment.
//
// Deletes a metastore assignment to a workspace, leaving the workspace with no
// metastore.
func (a *accountMetastoreAssignmentsBaseClient) DeleteByWorkspaceIdAndMetastoreId(ctx context.Context, workspaceId int64, metastoreId string) (*DeleteResponse, error) {
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
func (a *accountMetastoreAssignmentsBaseClient) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*AccountsMetastoreAssignment, error) {
	return a.accountMetastoreAssignmentsImpl.Get(ctx, GetAccountMetastoreAssignmentRequest{
		WorkspaceId: workspaceId,
	})
}

// Get all workspaces assigned to a metastore.
//
// Gets a list of all Databricks workspace IDs that have been assigned to given
// metastore.
func (a *accountMetastoreAssignmentsBaseClient) ListByMetastoreId(ctx context.Context, metastoreId string) (*ListAccountMetastoreAssignmentsResponse, error) {
	return a.accountMetastoreAssignmentsImpl.internalList(ctx, ListAccountMetastoreAssignmentsRequest{
		MetastoreId: metastoreId,
	})
}

type accountMetastoresBaseClient struct {
	accountMetastoresImpl
}

// Delete a metastore.
//
// Deletes a Unity Catalog metastore for an account, both specified by ID.
func (a *accountMetastoresBaseClient) DeleteByMetastoreId(ctx context.Context, metastoreId string) (*DeleteResponse, error) {
	return a.accountMetastoresImpl.Delete(ctx, DeleteAccountMetastoreRequest{
		MetastoreId: metastoreId,
	})
}

// Get a metastore.
//
// Gets a Unity Catalog metastore from an account, both specified by ID.
func (a *accountMetastoresBaseClient) GetByMetastoreId(ctx context.Context, metastoreId string) (*AccountsMetastoreInfo, error) {
	return a.accountMetastoresImpl.Get(ctx, GetAccountMetastoreRequest{
		MetastoreId: metastoreId,
	})
}

type accountStorageCredentialsBaseClient struct {
	accountStorageCredentialsImpl
}

// Delete a storage credential.
//
// Deletes a storage credential from the metastore. The caller must be an owner
// of the storage credential.
func (a *accountStorageCredentialsBaseClient) DeleteByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) (*DeleteResponse, error) {
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
func (a *accountStorageCredentialsBaseClient) GetByMetastoreIdAndStorageCredentialName(ctx context.Context, metastoreId string, storageCredentialName string) (*AccountsStorageCredentialInfo, error) {
	return a.accountStorageCredentialsImpl.Get(ctx, GetAccountStorageCredentialRequest{
		MetastoreId:           metastoreId,
		StorageCredentialName: storageCredentialName,
	})
}

// Get all storage credentials assigned to a metastore.
//
// Gets a list of all storage credentials that have been assigned to given
// metastore.
func (a *accountStorageCredentialsBaseClient) ListByMetastoreId(ctx context.Context, metastoreId string) (*ListAccountStorageCredentialsResponse, error) {
	return a.accountStorageCredentialsImpl.internalList(ctx, ListAccountStorageCredentialsRequest{
		MetastoreId: metastoreId,
	})
}

type artifactAllowlistsBaseClient struct {
	artifactAllowlistsImpl
}

// Get an artifact allowlist.
//
// Get the artifact allowlist of a certain artifact type. The caller must be a
// metastore admin or have the **MANAGE ALLOWLIST** privilege on the metastore.
func (a *artifactAllowlistsBaseClient) GetByArtifactType(ctx context.Context, artifactType ArtifactType) (*ArtifactAllowlistInfo, error) {
	return a.artifactAllowlistsImpl.Get(ctx, GetArtifactAllowlistRequest{
		ArtifactType: artifactType,
	})
}

type catalogsBaseClient struct {
	catalogsImpl
}

// Delete a catalog.
//
// Deletes the catalog that matches the supplied name. The caller must be a
// metastore admin or the owner of the catalog.
func (a *catalogsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.catalogsImpl.Delete(ctx, DeleteCatalogRequest{
		Name: name,
	})
}

// Get a catalog.
//
// Gets the specified catalog in a metastore. The caller must be a metastore
// admin, the owner of the catalog, or a user that has the **USE_CATALOG**
// privilege set for their account.
func (a *catalogsBaseClient) GetByName(ctx context.Context, name string) (*CatalogInfo, error) {
	return a.catalogsImpl.Get(ctx, GetCatalogRequest{
		Name: name,
	})
}

type connectionsBaseClient struct {
	connectionsImpl
}

// Delete a connection.
//
// Deletes the connection that matches the supplied name.
func (a *connectionsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.connectionsImpl.Delete(ctx, DeleteConnectionRequest{
		Name: name,
	})
}

// Get a connection.
//
// Gets a connection from it's name.
func (a *connectionsBaseClient) GetByName(ctx context.Context, name string) (*ConnectionInfo, error) {
	return a.connectionsImpl.Get(ctx, GetConnectionRequest{
		Name: name,
	})
}

// ConnectionInfoNameToFullNameMap calls [connectionsBaseClient.ListAll] and creates a map of results with [ConnectionInfo].Name as key and [ConnectionInfo].FullName as value.
//
// Returns an error if there's more than one [ConnectionInfo] with the same .Name.
//
// Note: All [ConnectionInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *connectionsBaseClient) ConnectionInfoNameToFullNameMap(ctx context.Context, request ListConnectionsRequest) (map[string]string, error) {
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

type credentialsBaseClient struct {
	credentialsImpl
}

// Delete a credential.
//
// Deletes a service or storage credential from the metastore. The caller must
// be an owner of the credential.
func (a *credentialsBaseClient) DeleteCredentialByNameArg(ctx context.Context, nameArg string) (*DeleteCredentialResponse, error) {
	return a.credentialsImpl.DeleteCredential(ctx, DeleteCredentialRequest{
		NameArg: nameArg,
	})
}

// Get a credential.
//
// Gets a service or storage credential from the metastore. The caller must be a
// metastore admin, the owner of the credential, or have any permission on the
// credential.
func (a *credentialsBaseClient) GetCredentialByNameArg(ctx context.Context, nameArg string) (*CredentialInfo, error) {
	return a.credentialsImpl.GetCredential(ctx, GetCredentialRequest{
		NameArg: nameArg,
	})
}

type externalLocationsBaseClient struct {
	externalLocationsImpl
}

// Delete an external location.
//
// Deletes the specified external location from the metastore. The caller must
// be the owner of the external location.
func (a *externalLocationsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.externalLocationsImpl.Delete(ctx, DeleteExternalLocationRequest{
		Name: name,
	})
}

// Get an external location.
//
// Gets an external location from the metastore. The caller must be either a
// metastore admin, the owner of the external location, or a user that has some
// privilege on the external location.
func (a *externalLocationsBaseClient) GetByName(ctx context.Context, name string) (*ExternalLocationInfo, error) {
	return a.externalLocationsImpl.Get(ctx, GetExternalLocationRequest{
		Name: name,
	})
}

type functionsBaseClient struct {
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
func (a *functionsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
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
func (a *functionsBaseClient) GetByName(ctx context.Context, name string) (*FunctionInfo, error) {
	return a.functionsImpl.Get(ctx, GetFunctionRequest{
		Name: name,
	})
}

// FunctionInfoNameToFullNameMap calls [functionsBaseClient.ListAll] and creates a map of results with [FunctionInfo].Name as key and [FunctionInfo].FullName as value.
//
// Returns an error if there's more than one [FunctionInfo] with the same .Name.
//
// Note: All [FunctionInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *functionsBaseClient) FunctionInfoNameToFullNameMap(ctx context.Context, request ListFunctionsRequest) (map[string]string, error) {
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

type grantsBaseClient struct {
	grantsImpl
}

// Get permissions.
//
// Gets the permissions for a securable.
func (a *grantsBaseClient) GetBySecurableTypeAndFullName(ctx context.Context, securableType SecurableType, fullName string) (*PermissionsList, error) {
	return a.grantsImpl.Get(ctx, GetGrantRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

// Get effective permissions.
//
// Gets the effective permissions for a securable.
func (a *grantsBaseClient) GetEffectiveBySecurableTypeAndFullName(ctx context.Context, securableType SecurableType, fullName string) (*EffectivePermissionsList, error) {
	return a.grantsImpl.GetEffective(ctx, GetEffectiveRequest{
		SecurableType: securableType,
		FullName:      fullName,
	})
}

type metastoresBaseClient struct {
	metastoresImpl
}

// Delete a metastore.
//
// Deletes a metastore. The caller must be a metastore admin.
func (a *metastoresBaseClient) DeleteById(ctx context.Context, id string) (*DeleteResponse, error) {
	return a.metastoresImpl.Delete(ctx, DeleteMetastoreRequest{
		Id: id,
	})
}

// Get a metastore.
//
// Gets a metastore that matches the supplied ID. The caller must be a metastore
// admin to retrieve this info.
func (a *metastoresBaseClient) GetById(ctx context.Context, id string) (*MetastoreInfo, error) {
	return a.metastoresImpl.Get(ctx, GetMetastoreRequest{
		Id: id,
	})
}

// MetastoreInfoNameToMetastoreIdMap calls [metastoresBaseClient.ListAll] and creates a map of results with [MetastoreInfo].Name as key and [MetastoreInfo].MetastoreId as value.
//
// Returns an error if there's more than one [MetastoreInfo] with the same .Name.
//
// Note: All [MetastoreInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *metastoresBaseClient) MetastoreInfoNameToMetastoreIdMap(ctx context.Context) (map[string]string, error) {
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

// GetByName calls [metastoresBaseClient.MetastoreInfoNameToMetastoreIdMap] and returns a single [MetastoreInfo].
//
// Returns an error if there's more than one [MetastoreInfo] with the same .Name.
//
// Note: All [MetastoreInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *metastoresBaseClient) GetByName(ctx context.Context, name string) (*MetastoreInfo, error) {
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
func (a *metastoresBaseClient) UnassignByWorkspaceId(ctx context.Context, workspaceId int64) (*UnassignResponse, error) {
	return a.metastoresImpl.Unassign(ctx, UnassignRequest{
		WorkspaceId: workspaceId,
	})
}

type modelVersionsBaseClient struct {
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
func (a *modelVersionsBaseClient) DeleteByFullNameAndVersion(ctx context.Context, fullName string, version int) (*DeleteResponse, error) {
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
func (a *modelVersionsBaseClient) GetByFullNameAndVersion(ctx context.Context, fullName string, version int) (*ModelVersionInfo, error) {
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
func (a *modelVersionsBaseClient) GetByAliasByFullNameAndAlias(ctx context.Context, fullName string, alias string) (*ModelVersionInfo, error) {
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
func (a *modelVersionsBaseClient) ListByFullName(ctx context.Context, fullName string) (*ListModelVersionsResponse, error) {
	return a.modelVersionsImpl.internalList(ctx, ListModelVersionsRequest{
		FullName: fullName,
	})
}

type onlineTablesBaseClient struct {
	onlineTablesImpl
}

// Delete an Online Table.
//
// Delete an online table. Warning: This will delete all the data in the online
// table. If the source Delta table was deleted or modified since this Online
// Table was created, this will lose the data forever!
func (a *onlineTablesBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.onlineTablesImpl.Delete(ctx, DeleteOnlineTableRequest{
		Name: name,
	})
}

// Get an Online Table.
//
// Get information about an existing online table and its status.
func (a *onlineTablesBaseClient) GetByName(ctx context.Context, name string) (*OnlineTable, error) {
	return a.onlineTablesImpl.Get(ctx, GetOnlineTableRequest{
		Name: name,
	})
}

type qualityMonitorsBaseClient struct {
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
func (a *qualityMonitorsBaseClient) DeleteByTableName(ctx context.Context, tableName string) (*DeleteResponse, error) {
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
func (a *qualityMonitorsBaseClient) GetByTableName(ctx context.Context, tableName string) (*MonitorInfo, error) {
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
func (a *qualityMonitorsBaseClient) GetRefreshByTableNameAndRefreshId(ctx context.Context, tableName string, refreshId string) (*MonitorRefreshInfo, error) {
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
func (a *qualityMonitorsBaseClient) ListRefreshesByTableName(ctx context.Context, tableName string) (*MonitorRefreshListResponse, error) {
	return a.qualityMonitorsImpl.ListRefreshes(ctx, ListRefreshesRequest{
		TableName: tableName,
	})
}

type registeredModelsBaseClient struct {
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
func (a *registeredModelsBaseClient) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
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
func (a *registeredModelsBaseClient) DeleteAliasByFullNameAndAlias(ctx context.Context, fullName string, alias string) (*DeleteAliasResponse, error) {
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
func (a *registeredModelsBaseClient) GetByFullName(ctx context.Context, fullName string) (*RegisteredModelInfo, error) {
	return a.registeredModelsImpl.Get(ctx, GetRegisteredModelRequest{
		FullName: fullName,
	})
}

// RegisteredModelInfoNameToFullNameMap calls [registeredModelsBaseClient.ListAll] and creates a map of results with [RegisteredModelInfo].Name as key and [RegisteredModelInfo].FullName as value.
//
// Returns an error if there's more than one [RegisteredModelInfo] with the same .Name.
//
// Note: All [RegisteredModelInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *registeredModelsBaseClient) RegisteredModelInfoNameToFullNameMap(ctx context.Context, request ListRegisteredModelsRequest) (map[string]string, error) {
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

// GetByName calls [registeredModelsBaseClient.RegisteredModelInfoNameToFullNameMap] and returns a single [RegisteredModelInfo].
//
// Returns an error if there's more than one [RegisteredModelInfo] with the same .Name.
//
// Note: All [RegisteredModelInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *registeredModelsBaseClient) GetByName(ctx context.Context, name string) (*RegisteredModelInfo, error) {
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

type resourceQuotasBaseClient struct {
	resourceQuotasImpl
}

// Get information for a single resource quota.
//
// The GetQuota API returns usage information for a single resource quota,
// defined as a child-parent pair. This API also refreshes the quota count if it
// is out of date. Refreshes are triggered asynchronously. The updated count
// might not be returned in the first call.
func (a *resourceQuotasBaseClient) GetQuotaByParentSecurableTypeAndParentFullNameAndQuotaName(ctx context.Context, parentSecurableType string, parentFullName string, quotaName string) (*GetQuotaResponse, error) {
	return a.resourceQuotasImpl.GetQuota(ctx, GetQuotaRequest{
		ParentSecurableType: parentSecurableType,
		ParentFullName:      parentFullName,
		QuotaName:           quotaName,
	})
}

type schemasBaseClient struct {
	schemasImpl
}

// Delete a schema.
//
// Deletes the specified schema from the parent catalog. The caller must be the
// owner of the schema or an owner of the parent catalog.
func (a *schemasBaseClient) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
	return a.schemasImpl.Delete(ctx, DeleteSchemaRequest{
		FullName: fullName,
	})
}

// Get a schema.
//
// Gets the specified schema within the metastore. The caller must be a
// metastore admin, the owner of the schema, or a user that has the
// **USE_SCHEMA** privilege on the schema.
func (a *schemasBaseClient) GetByFullName(ctx context.Context, fullName string) (*SchemaInfo, error) {
	return a.schemasImpl.Get(ctx, GetSchemaRequest{
		FullName: fullName,
	})
}

// SchemaInfoNameToFullNameMap calls [schemasBaseClient.ListAll] and creates a map of results with [SchemaInfo].Name as key and [SchemaInfo].FullName as value.
//
// Returns an error if there's more than one [SchemaInfo] with the same .Name.
//
// Note: All [SchemaInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *schemasBaseClient) SchemaInfoNameToFullNameMap(ctx context.Context, request ListSchemasRequest) (map[string]string, error) {
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

// GetByName calls [schemasBaseClient.SchemaInfoNameToFullNameMap] and returns a single [SchemaInfo].
//
// Returns an error if there's more than one [SchemaInfo] with the same .Name.
//
// Note: All [SchemaInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *schemasBaseClient) GetByName(ctx context.Context, name string) (*SchemaInfo, error) {
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

type storageCredentialsBaseClient struct {
	storageCredentialsImpl
}

// Delete a credential.
//
// Deletes a storage credential from the metastore. The caller must be an owner
// of the storage credential.
func (a *storageCredentialsBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.storageCredentialsImpl.Delete(ctx, DeleteStorageCredentialRequest{
		Name: name,
	})
}

// Get a credential.
//
// Gets a storage credential from the metastore. The caller must be a metastore
// admin, the owner of the storage credential, or have some permission on the
// storage credential.
func (a *storageCredentialsBaseClient) GetByName(ctx context.Context, name string) (*StorageCredentialInfo, error) {
	return a.storageCredentialsImpl.Get(ctx, GetStorageCredentialRequest{
		Name: name,
	})
}

// StorageCredentialInfoNameToIdMap calls [storageCredentialsBaseClient.ListAll] and creates a map of results with [StorageCredentialInfo].Name as key and [StorageCredentialInfo].Id as value.
//
// Returns an error if there's more than one [StorageCredentialInfo] with the same .Name.
//
// Note: All [StorageCredentialInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *storageCredentialsBaseClient) StorageCredentialInfoNameToIdMap(ctx context.Context, request ListStorageCredentialsRequest) (map[string]string, error) {
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

type systemSchemasBaseClient struct {
	systemSchemasImpl
}

// Disable a system schema.
//
// Disables the system schema and removes it from the system catalog. The caller
// must be an account admin or a metastore admin.
func (a *systemSchemasBaseClient) DisableByMetastoreIdAndSchemaName(ctx context.Context, metastoreId string, schemaName string) (*DisableResponse, error) {
	return a.systemSchemasImpl.Disable(ctx, DisableRequest{
		MetastoreId: metastoreId,
		SchemaName:  schemaName,
	})
}

// List system schemas.
//
// Gets an array of system schemas for a metastore. The caller must be an
// account admin or a metastore admin.
func (a *systemSchemasBaseClient) ListByMetastoreId(ctx context.Context, metastoreId string) (*ListSystemSchemasResponse, error) {
	return a.systemSchemasImpl.internalList(ctx, ListSystemSchemasRequest{
		MetastoreId: metastoreId,
	})
}

type tableConstraintsBaseClient struct {
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
func (a *tableConstraintsBaseClient) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
	return a.tableConstraintsImpl.Delete(ctx, DeleteTableConstraintRequest{
		FullName: fullName,
	})
}

type tablesBaseClient struct {
	tablesImpl
}

// Delete a table.
//
// Deletes a table from the specified parent catalog and schema. The caller must
// be the owner of the parent catalog, have the **USE_CATALOG** privilege on the
// parent catalog and be the owner of the parent schema, or be the owner of the
// table and have the **USE_CATALOG** privilege on the parent catalog and the
// **USE_SCHEMA** privilege on the parent schema.
func (a *tablesBaseClient) DeleteByFullName(ctx context.Context, fullName string) (*DeleteResponse, error) {
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
func (a *tablesBaseClient) ExistsByFullName(ctx context.Context, fullName string) (*TableExistsResponse, error) {
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
func (a *tablesBaseClient) GetByFullName(ctx context.Context, fullName string) (*TableInfo, error) {
	return a.tablesImpl.Get(ctx, GetTableRequest{
		FullName: fullName,
	})
}

// TableInfoNameToTableIdMap calls [tablesBaseClient.ListAll] and creates a map of results with [TableInfo].Name as key and [TableInfo].TableId as value.
//
// Returns an error if there's more than one [TableInfo] with the same .Name.
//
// Note: All [TableInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *tablesBaseClient) TableInfoNameToTableIdMap(ctx context.Context, request ListTablesRequest) (map[string]string, error) {
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

// GetByName calls [tablesBaseClient.TableInfoNameToTableIdMap] and returns a single [TableInfo].
//
// Returns an error if there's more than one [TableInfo] with the same .Name.
//
// Note: All [TableInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *tablesBaseClient) GetByName(ctx context.Context, name string) (*TableInfo, error) {
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

type temporaryTableCredentialsBaseClient struct {
	temporaryTableCredentialsImpl
}

type volumesBaseClient struct {
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
func (a *volumesBaseClient) DeleteByName(ctx context.Context, name string) (*DeleteResponse, error) {
	return a.volumesImpl.Delete(ctx, DeleteVolumeRequest{
		Name: name,
	})
}

// VolumeInfoNameToVolumeIdMap calls [volumesBaseClient.ListAll] and creates a map of results with [VolumeInfo].Name as key and [VolumeInfo].VolumeId as value.
//
// Returns an error if there's more than one [VolumeInfo] with the same .Name.
//
// Note: All [VolumeInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *volumesBaseClient) VolumeInfoNameToVolumeIdMap(ctx context.Context, request ListVolumesRequest) (map[string]string, error) {
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

// GetByName calls [volumesBaseClient.VolumeInfoNameToVolumeIdMap] and returns a single [VolumeInfo].
//
// Returns an error if there's more than one [VolumeInfo] with the same .Name.
//
// Note: All [VolumeInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *volumesBaseClient) GetByName(ctx context.Context, name string) (*VolumeInfo, error) {
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
func (a *volumesBaseClient) ReadByName(ctx context.Context, name string) (*VolumeInfo, error) {
	return a.volumesImpl.Read(ctx, ReadVolumeRequest{
		Name: name,
	})
}

type workspaceBindingsBaseClient struct {
	workspaceBindingsImpl
}

// Get catalog workspace bindings.
//
// Gets workspace bindings of the catalog. The caller must be a metastore admin
// or an owner of the catalog.
func (a *workspaceBindingsBaseClient) GetByName(ctx context.Context, name string) (*CurrentWorkspaceBindings, error) {
	return a.workspaceBindingsImpl.Get(ctx, GetWorkspaceBindingRequest{
		Name: name,
	})
}

// Get securable workspace bindings.
//
// Gets workspace bindings of the securable. The caller must be a metastore
// admin or an owner of the securable.
func (a *workspaceBindingsBaseClient) GetBindingsBySecurableTypeAndSecurableName(ctx context.Context, securableType GetBindingsSecurableType, securableName string) (*WorkspaceBindingsResponse, error) {
	return a.workspaceBindingsImpl.internalGetBindings(ctx, GetBindingsRequest{
		SecurableType: securableType,
		SecurableName: securableName,
	})
}
