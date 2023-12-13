// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/mocks/service/compute"
	"github.com/databricks/databricks-sdk-go/mocks/service/files"
	"github.com/databricks/databricks-sdk-go/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/mocks/service/jobs"
	"github.com/databricks/databricks-sdk-go/mocks/service/ml"
	"github.com/databricks/databricks-sdk-go/mocks/service/pipelines"
	"github.com/databricks/databricks-sdk-go/mocks/service/serving"
	"github.com/databricks/databricks-sdk-go/mocks/service/settings"
	"github.com/databricks/databricks-sdk-go/mocks/service/sharing"
	"github.com/databricks/databricks-sdk-go/mocks/service/sql"
	"github.com/databricks/databricks-sdk-go/mocks/service/workspace"
)

// NewWorkspaceClient creates new mocked version of Databricks SDK client for Workspaces
// which can be used for testing.
func NewWorkspaceClient(t interface {
	mock.TestingT
	Cleanup(func())
}, cfg *config.Config) *databricks.WorkspaceClient {
	return &databricks.WorkspaceClient{
		Config: cfg,

		AccountAccessControlProxy: iam.NewMockAccountAccessControlProxyAPIInterface(t),
		Alerts:                    sql.NewMockAlertsAPIInterface(t),
		Apps:                      serving.NewMockAppsAPIInterface(t),
		ArtifactAllowlists:        catalog.NewMockArtifactAllowlistsAPIInterface(t),
		Catalogs:                  catalog.NewMockCatalogsAPIInterface(t),
		CleanRooms:                sharing.NewMockCleanRoomsAPIInterface(t),
		ClusterPolicies:           compute.NewMockClusterPoliciesAPIInterface(t),
		Clusters:                  compute.NewMockClustersAPIInterface(t),
		CommandExecution:          compute.NewMockCommandExecutionAPIInterface(t),
		Connections:               catalog.NewMockConnectionsAPIInterface(t),
		CredentialsManager:        settings.NewMockCredentialsManagerAPIInterface(t),
		CurrentUser:               iam.NewMockCurrentUserAPIInterface(t),
		DashboardWidgets:          sql.NewMockDashboardWidgetsAPIInterface(t),
		Dashboards:                sql.NewMockDashboardsAPIInterface(t),
		DataSources:               sql.NewMockDataSourcesAPIInterface(t),
		Dbfs:                      files.NewMockDbfsAPIInterface(t),
		DbsqlPermissions:          sql.NewMockDbsqlPermissionsAPIInterface(t),
		Experiments:               ml.NewMockExperimentsAPIInterface(t),
		ExternalLocations:         catalog.NewMockExternalLocationsAPIInterface(t),
		Files:                     files.NewMockFilesAPIInterface(t),
		Functions:                 catalog.NewMockFunctionsAPIInterface(t),
		GitCredentials:            workspace.NewMockGitCredentialsAPIInterface(t),
		GlobalInitScripts:         compute.NewMockGlobalInitScriptsAPIInterface(t),
		Grants:                    catalog.NewMockGrantsAPIInterface(t),
		Groups:                    iam.NewMockGroupsAPIInterface(t),
		InstancePools:             compute.NewMockInstancePoolsAPIInterface(t),
		InstanceProfiles:          compute.NewMockInstanceProfilesAPIInterface(t),
		IpAccessLists:             settings.NewMockIpAccessListsAPIInterface(t),
		Jobs:                      jobs.NewMockJobsAPIInterface(t),
		Libraries:                 compute.NewMockLibrariesAPIInterface(t),
		Metastores:                catalog.NewMockMetastoresAPIInterface(t),
		ModelRegistry:             ml.NewMockModelRegistryAPIInterface(t),
		ModelVersions:             catalog.NewMockModelVersionsAPIInterface(t),
		Permissions:               iam.NewMockPermissionsAPIInterface(t),
		Pipelines:                 pipelines.NewMockPipelinesAPIInterface(t),
		PolicyFamilies:            compute.NewMockPolicyFamiliesAPIInterface(t),
		Providers:                 sharing.NewMockProvidersAPIInterface(t),
		Queries:                   sql.NewMockQueriesAPIInterface(t),
		QueryHistory:              sql.NewMockQueryHistoryAPIInterface(t),
		QueryVisualizations:       sql.NewMockQueryVisualizationsAPIInterface(t),
		RecipientActivation:       sharing.NewMockRecipientActivationAPIInterface(t),
		Recipients:                sharing.NewMockRecipientsAPIInterface(t),
		RegisteredModels:          catalog.NewMockRegisteredModelsAPIInterface(t),
		Repos:                     workspace.NewMockReposAPIInterface(t),
		Schemas:                   catalog.NewMockSchemasAPIInterface(t),
		Secrets:                   workspace.NewMockSecretsAPIInterface(t),
		ServicePrincipals:         iam.NewMockServicePrincipalsAPIInterface(t),
		ServingEndpoints:          serving.NewMockServingEndpointsAPIInterface(t),
		Settings:                  settings.NewMockSettingsAPIInterface(t),
		Shares:                    sharing.NewMockSharesAPIInterface(t),
		StatementExecution:        sql.NewMockStatementExecutionAPIInterface(t),
		StorageCredentials:        catalog.NewMockStorageCredentialsAPIInterface(t),
		SystemSchemas:             catalog.NewMockSystemSchemasAPIInterface(t),
		TableConstraints:          catalog.NewMockTableConstraintsAPIInterface(t),
		Tables:                    catalog.NewMockTablesAPIInterface(t),
		TokenManagement:           settings.NewMockTokenManagementAPIInterface(t),
		Tokens:                    settings.NewMockTokensAPIInterface(t),
		Users:                     iam.NewMockUsersAPIInterface(t),
		Volumes:                   catalog.NewMockVolumesAPIInterface(t),
		Warehouses:                sql.NewMockWarehousesAPIInterface(t),
		Workspace:                 workspace.NewMockWorkspaceAPIInterface(t),
		WorkspaceBindings:         catalog.NewMockWorkspaceBindingsAPIInterface(t),
		WorkspaceConf:             settings.NewMockWorkspaceConfAPIInterface(t),
	}
}

func GetMockAccountAccessControlProxyAPI(c *databricks.WorkspaceClient) *iam.MockAccountAccessControlProxyAPIInterface {
	api, ok := c.AccountAccessControlProxy.(*iam.MockAccountAccessControlProxyAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockAlertsAPI(c *databricks.WorkspaceClient) *sql.MockAlertsAPIInterface {
	api, ok := c.Alerts.(*sql.MockAlertsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockAppsAPI(c *databricks.WorkspaceClient) *serving.MockAppsAPIInterface {
	api, ok := c.Apps.(*serving.MockAppsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockArtifactAllowlistsAPI(c *databricks.WorkspaceClient) *catalog.MockArtifactAllowlistsAPIInterface {
	api, ok := c.ArtifactAllowlists.(*catalog.MockArtifactAllowlistsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockCatalogsAPI(c *databricks.WorkspaceClient) *catalog.MockCatalogsAPIInterface {
	api, ok := c.Catalogs.(*catalog.MockCatalogsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockCleanRoomsAPI(c *databricks.WorkspaceClient) *sharing.MockCleanRoomsAPIInterface {
	api, ok := c.CleanRooms.(*sharing.MockCleanRoomsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockClusterPoliciesAPI(c *databricks.WorkspaceClient) *compute.MockClusterPoliciesAPIInterface {
	api, ok := c.ClusterPolicies.(*compute.MockClusterPoliciesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockClustersAPI(c *databricks.WorkspaceClient) *compute.MockClustersAPIInterface {
	api, ok := c.Clusters.(*compute.MockClustersAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockCommandExecutionAPI(c *databricks.WorkspaceClient) *compute.MockCommandExecutionAPIInterface {
	api, ok := c.CommandExecution.(*compute.MockCommandExecutionAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockConnectionsAPI(c *databricks.WorkspaceClient) *catalog.MockConnectionsAPIInterface {
	api, ok := c.Connections.(*catalog.MockConnectionsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockCredentialsManagerAPI(c *databricks.WorkspaceClient) *settings.MockCredentialsManagerAPIInterface {
	api, ok := c.CredentialsManager.(*settings.MockCredentialsManagerAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockCurrentUserAPI(c *databricks.WorkspaceClient) *iam.MockCurrentUserAPIInterface {
	api, ok := c.CurrentUser.(*iam.MockCurrentUserAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockDashboardWidgetsAPI(c *databricks.WorkspaceClient) *sql.MockDashboardWidgetsAPIInterface {
	api, ok := c.DashboardWidgets.(*sql.MockDashboardWidgetsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockDashboardsAPI(c *databricks.WorkspaceClient) *sql.MockDashboardsAPIInterface {
	api, ok := c.Dashboards.(*sql.MockDashboardsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockDataSourcesAPI(c *databricks.WorkspaceClient) *sql.MockDataSourcesAPIInterface {
	api, ok := c.DataSources.(*sql.MockDataSourcesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockDbfsAPI(c *databricks.WorkspaceClient) *files.MockDbfsAPIInterface {
	api, ok := c.Dbfs.(*files.MockDbfsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockDbsqlPermissionsAPI(c *databricks.WorkspaceClient) *sql.MockDbsqlPermissionsAPIInterface {
	api, ok := c.DbsqlPermissions.(*sql.MockDbsqlPermissionsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockExperimentsAPI(c *databricks.WorkspaceClient) *ml.MockExperimentsAPIInterface {
	api, ok := c.Experiments.(*ml.MockExperimentsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockExternalLocationsAPI(c *databricks.WorkspaceClient) *catalog.MockExternalLocationsAPIInterface {
	api, ok := c.ExternalLocations.(*catalog.MockExternalLocationsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockFilesAPI(c *databricks.WorkspaceClient) *files.MockFilesAPIInterface {
	api, ok := c.Files.(*files.MockFilesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockFunctionsAPI(c *databricks.WorkspaceClient) *catalog.MockFunctionsAPIInterface {
	api, ok := c.Functions.(*catalog.MockFunctionsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockGitCredentialsAPI(c *databricks.WorkspaceClient) *workspace.MockGitCredentialsAPIInterface {
	api, ok := c.GitCredentials.(*workspace.MockGitCredentialsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockGlobalInitScriptsAPI(c *databricks.WorkspaceClient) *compute.MockGlobalInitScriptsAPIInterface {
	api, ok := c.GlobalInitScripts.(*compute.MockGlobalInitScriptsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockGrantsAPI(c *databricks.WorkspaceClient) *catalog.MockGrantsAPIInterface {
	api, ok := c.Grants.(*catalog.MockGrantsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockGroupsAPI(c *databricks.WorkspaceClient) *iam.MockGroupsAPIInterface {
	api, ok := c.Groups.(*iam.MockGroupsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockInstancePoolsAPI(c *databricks.WorkspaceClient) *compute.MockInstancePoolsAPIInterface {
	api, ok := c.InstancePools.(*compute.MockInstancePoolsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockInstanceProfilesAPI(c *databricks.WorkspaceClient) *compute.MockInstanceProfilesAPIInterface {
	api, ok := c.InstanceProfiles.(*compute.MockInstanceProfilesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockIpAccessListsAPI(c *databricks.WorkspaceClient) *settings.MockIpAccessListsAPIInterface {
	api, ok := c.IpAccessLists.(*settings.MockIpAccessListsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockJobsAPI(c *databricks.WorkspaceClient) *jobs.MockJobsAPIInterface {
	api, ok := c.Jobs.(*jobs.MockJobsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockLibrariesAPI(c *databricks.WorkspaceClient) *compute.MockLibrariesAPIInterface {
	api, ok := c.Libraries.(*compute.MockLibrariesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockMetastoresAPI(c *databricks.WorkspaceClient) *catalog.MockMetastoresAPIInterface {
	api, ok := c.Metastores.(*catalog.MockMetastoresAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockModelRegistryAPI(c *databricks.WorkspaceClient) *ml.MockModelRegistryAPIInterface {
	api, ok := c.ModelRegistry.(*ml.MockModelRegistryAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockModelVersionsAPI(c *databricks.WorkspaceClient) *catalog.MockModelVersionsAPIInterface {
	api, ok := c.ModelVersions.(*catalog.MockModelVersionsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockPermissionsAPI(c *databricks.WorkspaceClient) *iam.MockPermissionsAPIInterface {
	api, ok := c.Permissions.(*iam.MockPermissionsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockPipelinesAPI(c *databricks.WorkspaceClient) *pipelines.MockPipelinesAPIInterface {
	api, ok := c.Pipelines.(*pipelines.MockPipelinesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockPolicyFamiliesAPI(c *databricks.WorkspaceClient) *compute.MockPolicyFamiliesAPIInterface {
	api, ok := c.PolicyFamilies.(*compute.MockPolicyFamiliesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockProvidersAPI(c *databricks.WorkspaceClient) *sharing.MockProvidersAPIInterface {
	api, ok := c.Providers.(*sharing.MockProvidersAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockQueriesAPI(c *databricks.WorkspaceClient) *sql.MockQueriesAPIInterface {
	api, ok := c.Queries.(*sql.MockQueriesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockQueryHistoryAPI(c *databricks.WorkspaceClient) *sql.MockQueryHistoryAPIInterface {
	api, ok := c.QueryHistory.(*sql.MockQueryHistoryAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockQueryVisualizationsAPI(c *databricks.WorkspaceClient) *sql.MockQueryVisualizationsAPIInterface {
	api, ok := c.QueryVisualizations.(*sql.MockQueryVisualizationsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockRecipientActivationAPI(c *databricks.WorkspaceClient) *sharing.MockRecipientActivationAPIInterface {
	api, ok := c.RecipientActivation.(*sharing.MockRecipientActivationAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockRecipientsAPI(c *databricks.WorkspaceClient) *sharing.MockRecipientsAPIInterface {
	api, ok := c.Recipients.(*sharing.MockRecipientsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockRegisteredModelsAPI(c *databricks.WorkspaceClient) *catalog.MockRegisteredModelsAPIInterface {
	api, ok := c.RegisteredModels.(*catalog.MockRegisteredModelsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockReposAPI(c *databricks.WorkspaceClient) *workspace.MockReposAPIInterface {
	api, ok := c.Repos.(*workspace.MockReposAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockSchemasAPI(c *databricks.WorkspaceClient) *catalog.MockSchemasAPIInterface {
	api, ok := c.Schemas.(*catalog.MockSchemasAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockSecretsAPI(c *databricks.WorkspaceClient) *workspace.MockSecretsAPIInterface {
	api, ok := c.Secrets.(*workspace.MockSecretsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockServicePrincipalsAPI(c *databricks.WorkspaceClient) *iam.MockServicePrincipalsAPIInterface {
	api, ok := c.ServicePrincipals.(*iam.MockServicePrincipalsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockServingEndpointsAPI(c *databricks.WorkspaceClient) *serving.MockServingEndpointsAPIInterface {
	api, ok := c.ServingEndpoints.(*serving.MockServingEndpointsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockSettingsAPI(c *databricks.WorkspaceClient) *settings.MockSettingsAPIInterface {
	api, ok := c.Settings.(*settings.MockSettingsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockSharesAPI(c *databricks.WorkspaceClient) *sharing.MockSharesAPIInterface {
	api, ok := c.Shares.(*sharing.MockSharesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockStatementExecutionAPI(c *databricks.WorkspaceClient) *sql.MockStatementExecutionAPIInterface {
	api, ok := c.StatementExecution.(*sql.MockStatementExecutionAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockStorageCredentialsAPI(c *databricks.WorkspaceClient) *catalog.MockStorageCredentialsAPIInterface {
	api, ok := c.StorageCredentials.(*catalog.MockStorageCredentialsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockSystemSchemasAPI(c *databricks.WorkspaceClient) *catalog.MockSystemSchemasAPIInterface {
	api, ok := c.SystemSchemas.(*catalog.MockSystemSchemasAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockTableConstraintsAPI(c *databricks.WorkspaceClient) *catalog.MockTableConstraintsAPIInterface {
	api, ok := c.TableConstraints.(*catalog.MockTableConstraintsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockTablesAPI(c *databricks.WorkspaceClient) *catalog.MockTablesAPIInterface {
	api, ok := c.Tables.(*catalog.MockTablesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockTokenManagementAPI(c *databricks.WorkspaceClient) *settings.MockTokenManagementAPIInterface {
	api, ok := c.TokenManagement.(*settings.MockTokenManagementAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockTokensAPI(c *databricks.WorkspaceClient) *settings.MockTokensAPIInterface {
	api, ok := c.Tokens.(*settings.MockTokensAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockUsersAPI(c *databricks.WorkspaceClient) *iam.MockUsersAPIInterface {
	api, ok := c.Users.(*iam.MockUsersAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockVolumesAPI(c *databricks.WorkspaceClient) *catalog.MockVolumesAPIInterface {
	api, ok := c.Volumes.(*catalog.MockVolumesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockWarehousesAPI(c *databricks.WorkspaceClient) *sql.MockWarehousesAPIInterface {
	api, ok := c.Warehouses.(*sql.MockWarehousesAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockWorkspaceAPI(c *databricks.WorkspaceClient) *workspace.MockWorkspaceAPIInterface {
	api, ok := c.Workspace.(*workspace.MockWorkspaceAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockWorkspaceBindingsAPI(c *databricks.WorkspaceClient) *catalog.MockWorkspaceBindingsAPIInterface {
	api, ok := c.WorkspaceBindings.(*catalog.MockWorkspaceBindingsAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
func GetMockWorkspaceConfAPI(c *databricks.WorkspaceClient) *settings.MockWorkspaceConfAPIInterface {
	api, ok := c.WorkspaceConf.(*settings.MockWorkspaceConfAPIInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
