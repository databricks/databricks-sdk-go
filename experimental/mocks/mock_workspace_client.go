// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/compute"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/files"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/jobs"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/ml"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/pipelines"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/serving"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settings"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sharing"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sql"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/workspace"
)

// NewWorkspaceClient creates new mocked version of Databricks SDK client for Workspaces
// which can be used for testing.
func NewWorkspaceClient(t interface {
	mock.TestingT
	Cleanup(func())
}, cfg *config.Config) *databricks.WorkspaceClient {
	return &databricks.WorkspaceClient{
		Config: cfg,

		AccountAccessControlProxy: iam.NewMockAccountAccessControlProxyInterface(t),
		Alerts:                    sql.NewMockAlertsInterface(t),
		Apps:                      serving.NewMockAppsInterface(t),
		ArtifactAllowlists:        catalog.NewMockArtifactAllowlistsInterface(t),
		Catalogs:                  catalog.NewMockCatalogsInterface(t),
		CleanRooms:                sharing.NewMockCleanRoomsInterface(t),
		ClusterPolicies:           compute.NewMockClusterPoliciesInterface(t),
		Clusters:                  compute.NewMockClustersInterface(t),
		CommandExecution:          compute.NewMockCommandExecutionInterface(t),
		Connections:               catalog.NewMockConnectionsInterface(t),
		CredentialsManager:        settings.NewMockCredentialsManagerInterface(t),
		CurrentUser:               iam.NewMockCurrentUserInterface(t),
		DashboardWidgets:          sql.NewMockDashboardWidgetsInterface(t),
		Dashboards:                sql.NewMockDashboardsInterface(t),
		DataSources:               sql.NewMockDataSourcesInterface(t),
		Dbfs:                      files.NewMockDbfsInterface(t),
		DbsqlPermissions:          sql.NewMockDbsqlPermissionsInterface(t),
		Experiments:               ml.NewMockExperimentsInterface(t),
		ExternalLocations:         catalog.NewMockExternalLocationsInterface(t),
		Files:                     files.NewMockFilesInterface(t),
		Functions:                 catalog.NewMockFunctionsInterface(t),
		GitCredentials:            workspace.NewMockGitCredentialsInterface(t),
		GlobalInitScripts:         compute.NewMockGlobalInitScriptsInterface(t),
		Grants:                    catalog.NewMockGrantsInterface(t),
		Groups:                    iam.NewMockGroupsInterface(t),
		InstancePools:             compute.NewMockInstancePoolsInterface(t),
		InstanceProfiles:          compute.NewMockInstanceProfilesInterface(t),
		IpAccessLists:             settings.NewMockIpAccessListsInterface(t),
		Jobs:                      jobs.NewMockJobsInterface(t),
		Libraries:                 compute.NewMockLibrariesInterface(t),
		Metastores:                catalog.NewMockMetastoresInterface(t),
		ModelRegistry:             ml.NewMockModelRegistryInterface(t),
		ModelVersions:             catalog.NewMockModelVersionsInterface(t),
		Permissions:               iam.NewMockPermissionsInterface(t),
		Pipelines:                 pipelines.NewMockPipelinesInterface(t),
		PolicyFamilies:            compute.NewMockPolicyFamiliesInterface(t),
		Providers:                 sharing.NewMockProvidersInterface(t),
		Queries:                   sql.NewMockQueriesInterface(t),
		QueryHistory:              sql.NewMockQueryHistoryInterface(t),
		QueryVisualizations:       sql.NewMockQueryVisualizationsInterface(t),
		RecipientActivation:       sharing.NewMockRecipientActivationInterface(t),
		Recipients:                sharing.NewMockRecipientsInterface(t),
		RegisteredModels:          catalog.NewMockRegisteredModelsInterface(t),
		Repos:                     workspace.NewMockReposInterface(t),
		Schemas:                   catalog.NewMockSchemasInterface(t),
		Secrets:                   workspace.NewMockSecretsInterface(t),
		ServicePrincipals:         iam.NewMockServicePrincipalsInterface(t),
		ServingEndpoints:          serving.NewMockServingEndpointsInterface(t),
		Settings:                  settings.NewMockSettingsInterface(t),
		Shares:                    sharing.NewMockSharesInterface(t),
		StatementExecution:        sql.NewMockStatementExecutionInterface(t),
		StorageCredentials:        catalog.NewMockStorageCredentialsInterface(t),
		SystemSchemas:             catalog.NewMockSystemSchemasInterface(t),
		TableConstraints:          catalog.NewMockTableConstraintsInterface(t),
		Tables:                    catalog.NewMockTablesInterface(t),
		TokenManagement:           settings.NewMockTokenManagementInterface(t),
		Tokens:                    settings.NewMockTokensInterface(t),
		Users:                     iam.NewMockUsersInterface(t),
		Volumes:                   catalog.NewMockVolumesInterface(t),
		Warehouses:                sql.NewMockWarehousesInterface(t),
		Workspace:                 workspace.NewMockWorkspaceInterface(t),
		WorkspaceBindings:         catalog.NewMockWorkspaceBindingsInterface(t),
		WorkspaceConf:             settings.NewMockWorkspaceConfInterface(t),
	}
}

func GetMockAccountAccessControlProxyAPI(c *databricks.WorkspaceClient) *iam.MockAccountAccessControlProxyInterface {
	api, ok := c.AccountAccessControlProxy.(*iam.MockAccountAccessControlProxyInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAlertsAPI(c *databricks.WorkspaceClient) *sql.MockAlertsInterface {
	api, ok := c.Alerts.(*sql.MockAlertsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAppsAPI(c *databricks.WorkspaceClient) *serving.MockAppsInterface {
	api, ok := c.Apps.(*serving.MockAppsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockArtifactAllowlistsAPI(c *databricks.WorkspaceClient) *catalog.MockArtifactAllowlistsInterface {
	api, ok := c.ArtifactAllowlists.(*catalog.MockArtifactAllowlistsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCatalogsAPI(c *databricks.WorkspaceClient) *catalog.MockCatalogsInterface {
	api, ok := c.Catalogs.(*catalog.MockCatalogsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCleanRoomsAPI(c *databricks.WorkspaceClient) *sharing.MockCleanRoomsInterface {
	api, ok := c.CleanRooms.(*sharing.MockCleanRoomsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockClusterPoliciesAPI(c *databricks.WorkspaceClient) *compute.MockClusterPoliciesInterface {
	api, ok := c.ClusterPolicies.(*compute.MockClusterPoliciesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockClustersAPI(c *databricks.WorkspaceClient) *compute.MockClustersInterface {
	api, ok := c.Clusters.(*compute.MockClustersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCommandExecutionAPI(c *databricks.WorkspaceClient) *compute.MockCommandExecutionInterface {
	api, ok := c.CommandExecution.(*compute.MockCommandExecutionInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockConnectionsAPI(c *databricks.WorkspaceClient) *catalog.MockConnectionsInterface {
	api, ok := c.Connections.(*catalog.MockConnectionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCredentialsManagerAPI(c *databricks.WorkspaceClient) *settings.MockCredentialsManagerInterface {
	api, ok := c.CredentialsManager.(*settings.MockCredentialsManagerInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCurrentUserAPI(c *databricks.WorkspaceClient) *iam.MockCurrentUserInterface {
	api, ok := c.CurrentUser.(*iam.MockCurrentUserInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockDashboardWidgetsAPI(c *databricks.WorkspaceClient) *sql.MockDashboardWidgetsInterface {
	api, ok := c.DashboardWidgets.(*sql.MockDashboardWidgetsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockDashboardsAPI(c *databricks.WorkspaceClient) *sql.MockDashboardsInterface {
	api, ok := c.Dashboards.(*sql.MockDashboardsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockDataSourcesAPI(c *databricks.WorkspaceClient) *sql.MockDataSourcesInterface {
	api, ok := c.DataSources.(*sql.MockDataSourcesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockDbfsAPI(c *databricks.WorkspaceClient) *files.MockDbfsInterface {
	api, ok := c.Dbfs.(*files.MockDbfsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockDbsqlPermissionsAPI(c *databricks.WorkspaceClient) *sql.MockDbsqlPermissionsInterface {
	api, ok := c.DbsqlPermissions.(*sql.MockDbsqlPermissionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockExperimentsAPI(c *databricks.WorkspaceClient) *ml.MockExperimentsInterface {
	api, ok := c.Experiments.(*ml.MockExperimentsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockExternalLocationsAPI(c *databricks.WorkspaceClient) *catalog.MockExternalLocationsInterface {
	api, ok := c.ExternalLocations.(*catalog.MockExternalLocationsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockFilesAPI(c *databricks.WorkspaceClient) *files.MockFilesInterface {
	api, ok := c.Files.(*files.MockFilesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockFunctionsAPI(c *databricks.WorkspaceClient) *catalog.MockFunctionsInterface {
	api, ok := c.Functions.(*catalog.MockFunctionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockGitCredentialsAPI(c *databricks.WorkspaceClient) *workspace.MockGitCredentialsInterface {
	api, ok := c.GitCredentials.(*workspace.MockGitCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockGlobalInitScriptsAPI(c *databricks.WorkspaceClient) *compute.MockGlobalInitScriptsInterface {
	api, ok := c.GlobalInitScripts.(*compute.MockGlobalInitScriptsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockGrantsAPI(c *databricks.WorkspaceClient) *catalog.MockGrantsInterface {
	api, ok := c.Grants.(*catalog.MockGrantsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockGroupsAPI(c *databricks.WorkspaceClient) *iam.MockGroupsInterface {
	api, ok := c.Groups.(*iam.MockGroupsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockInstancePoolsAPI(c *databricks.WorkspaceClient) *compute.MockInstancePoolsInterface {
	api, ok := c.InstancePools.(*compute.MockInstancePoolsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockInstanceProfilesAPI(c *databricks.WorkspaceClient) *compute.MockInstanceProfilesInterface {
	api, ok := c.InstanceProfiles.(*compute.MockInstanceProfilesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockIpAccessListsAPI(c *databricks.WorkspaceClient) *settings.MockIpAccessListsInterface {
	api, ok := c.IpAccessLists.(*settings.MockIpAccessListsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockJobsAPI(c *databricks.WorkspaceClient) *jobs.MockJobsInterface {
	api, ok := c.Jobs.(*jobs.MockJobsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockLibrariesAPI(c *databricks.WorkspaceClient) *compute.MockLibrariesInterface {
	api, ok := c.Libraries.(*compute.MockLibrariesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockMetastoresAPI(c *databricks.WorkspaceClient) *catalog.MockMetastoresInterface {
	api, ok := c.Metastores.(*catalog.MockMetastoresInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockModelRegistryAPI(c *databricks.WorkspaceClient) *ml.MockModelRegistryInterface {
	api, ok := c.ModelRegistry.(*ml.MockModelRegistryInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockModelVersionsAPI(c *databricks.WorkspaceClient) *catalog.MockModelVersionsInterface {
	api, ok := c.ModelVersions.(*catalog.MockModelVersionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockPermissionsAPI(c *databricks.WorkspaceClient) *iam.MockPermissionsInterface {
	api, ok := c.Permissions.(*iam.MockPermissionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockPipelinesAPI(c *databricks.WorkspaceClient) *pipelines.MockPipelinesInterface {
	api, ok := c.Pipelines.(*pipelines.MockPipelinesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockPolicyFamiliesAPI(c *databricks.WorkspaceClient) *compute.MockPolicyFamiliesInterface {
	api, ok := c.PolicyFamilies.(*compute.MockPolicyFamiliesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockProvidersAPI(c *databricks.WorkspaceClient) *sharing.MockProvidersInterface {
	api, ok := c.Providers.(*sharing.MockProvidersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockQueriesAPI(c *databricks.WorkspaceClient) *sql.MockQueriesInterface {
	api, ok := c.Queries.(*sql.MockQueriesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockQueryHistoryAPI(c *databricks.WorkspaceClient) *sql.MockQueryHistoryInterface {
	api, ok := c.QueryHistory.(*sql.MockQueryHistoryInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockQueryVisualizationsAPI(c *databricks.WorkspaceClient) *sql.MockQueryVisualizationsInterface {
	api, ok := c.QueryVisualizations.(*sql.MockQueryVisualizationsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockRecipientActivationAPI(c *databricks.WorkspaceClient) *sharing.MockRecipientActivationInterface {
	api, ok := c.RecipientActivation.(*sharing.MockRecipientActivationInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockRecipientsAPI(c *databricks.WorkspaceClient) *sharing.MockRecipientsInterface {
	api, ok := c.Recipients.(*sharing.MockRecipientsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockRegisteredModelsAPI(c *databricks.WorkspaceClient) *catalog.MockRegisteredModelsInterface {
	api, ok := c.RegisteredModels.(*catalog.MockRegisteredModelsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockReposAPI(c *databricks.WorkspaceClient) *workspace.MockReposInterface {
	api, ok := c.Repos.(*workspace.MockReposInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockSchemasAPI(c *databricks.WorkspaceClient) *catalog.MockSchemasInterface {
	api, ok := c.Schemas.(*catalog.MockSchemasInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockSecretsAPI(c *databricks.WorkspaceClient) *workspace.MockSecretsInterface {
	api, ok := c.Secrets.(*workspace.MockSecretsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockServicePrincipalsAPI(c *databricks.WorkspaceClient) *iam.MockServicePrincipalsInterface {
	api, ok := c.ServicePrincipals.(*iam.MockServicePrincipalsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockServingEndpointsAPI(c *databricks.WorkspaceClient) *serving.MockServingEndpointsInterface {
	api, ok := c.ServingEndpoints.(*serving.MockServingEndpointsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockSettingsAPI(c *databricks.WorkspaceClient) *settings.MockSettingsInterface {
	api, ok := c.Settings.(*settings.MockSettingsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockSharesAPI(c *databricks.WorkspaceClient) *sharing.MockSharesInterface {
	api, ok := c.Shares.(*sharing.MockSharesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockStatementExecutionAPI(c *databricks.WorkspaceClient) *sql.MockStatementExecutionInterface {
	api, ok := c.StatementExecution.(*sql.MockStatementExecutionInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockStorageCredentialsAPI(c *databricks.WorkspaceClient) *catalog.MockStorageCredentialsInterface {
	api, ok := c.StorageCredentials.(*catalog.MockStorageCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockSystemSchemasAPI(c *databricks.WorkspaceClient) *catalog.MockSystemSchemasInterface {
	api, ok := c.SystemSchemas.(*catalog.MockSystemSchemasInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockTableConstraintsAPI(c *databricks.WorkspaceClient) *catalog.MockTableConstraintsInterface {
	api, ok := c.TableConstraints.(*catalog.MockTableConstraintsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockTablesAPI(c *databricks.WorkspaceClient) *catalog.MockTablesInterface {
	api, ok := c.Tables.(*catalog.MockTablesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockTokenManagementAPI(c *databricks.WorkspaceClient) *settings.MockTokenManagementInterface {
	api, ok := c.TokenManagement.(*settings.MockTokenManagementInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockTokensAPI(c *databricks.WorkspaceClient) *settings.MockTokensInterface {
	api, ok := c.Tokens.(*settings.MockTokensInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockUsersAPI(c *databricks.WorkspaceClient) *iam.MockUsersInterface {
	api, ok := c.Users.(*iam.MockUsersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockVolumesAPI(c *databricks.WorkspaceClient) *catalog.MockVolumesInterface {
	api, ok := c.Volumes.(*catalog.MockVolumesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockWarehousesAPI(c *databricks.WorkspaceClient) *sql.MockWarehousesInterface {
	api, ok := c.Warehouses.(*sql.MockWarehousesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockWorkspaceAPI(c *databricks.WorkspaceClient) *workspace.MockWorkspaceInterface {
	api, ok := c.Workspace.(*workspace.MockWorkspaceInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockWorkspaceBindingsAPI(c *databricks.WorkspaceClient) *catalog.MockWorkspaceBindingsInterface {
	api, ok := c.WorkspaceBindings.(*catalog.MockWorkspaceBindingsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockWorkspaceConfAPI(c *databricks.WorkspaceClient) *settings.MockWorkspaceConfInterface {
	api, ok := c.WorkspaceConf.(*settings.MockWorkspaceConfInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
