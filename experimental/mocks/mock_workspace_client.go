// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

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

type MockWorkspaceClient struct {
	WorkspaceClient *databricks.WorkspaceClient
}

// NewMockWorkspaceClient creates new mocked version of Databricks SDK client for Workspaces
// which can be used for testing.
func NewMockWorkspaceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWorkspaceClient {
	return &MockWorkspaceClient{
		WorkspaceClient: &databricks.WorkspaceClient{
			Config: nil,

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
		},
	}
}

func (m *MockWorkspaceClient) GetMockAccountAccessControlProxyAPI() *iam.MockAccountAccessControlProxyInterface {
	api, ok := m.WorkspaceClient.AccountAccessControlProxy.(*iam.MockAccountAccessControlProxyInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAlertsAPI() *sql.MockAlertsInterface {
	api, ok := m.WorkspaceClient.Alerts.(*sql.MockAlertsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAppsAPI() *serving.MockAppsInterface {
	api, ok := m.WorkspaceClient.Apps.(*serving.MockAppsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockArtifactAllowlistsAPI() *catalog.MockArtifactAllowlistsInterface {
	api, ok := m.WorkspaceClient.ArtifactAllowlists.(*catalog.MockArtifactAllowlistsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCatalogsAPI() *catalog.MockCatalogsInterface {
	api, ok := m.WorkspaceClient.Catalogs.(*catalog.MockCatalogsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomsAPI() *sharing.MockCleanRoomsInterface {
	api, ok := m.WorkspaceClient.CleanRooms.(*sharing.MockCleanRoomsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockClusterPoliciesAPI() *compute.MockClusterPoliciesInterface {
	api, ok := m.WorkspaceClient.ClusterPolicies.(*compute.MockClusterPoliciesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockClustersAPI() *compute.MockClustersInterface {
	api, ok := m.WorkspaceClient.Clusters.(*compute.MockClustersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCommandExecutionAPI() *compute.MockCommandExecutionInterface {
	api, ok := m.WorkspaceClient.CommandExecution.(*compute.MockCommandExecutionInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConnectionsAPI() *catalog.MockConnectionsInterface {
	api, ok := m.WorkspaceClient.Connections.(*catalog.MockConnectionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCredentialsManagerAPI() *settings.MockCredentialsManagerInterface {
	api, ok := m.WorkspaceClient.CredentialsManager.(*settings.MockCredentialsManagerInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCurrentUserAPI() *iam.MockCurrentUserInterface {
	api, ok := m.WorkspaceClient.CurrentUser.(*iam.MockCurrentUserInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardWidgetsAPI() *sql.MockDashboardWidgetsInterface {
	api, ok := m.WorkspaceClient.DashboardWidgets.(*sql.MockDashboardWidgetsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardsAPI() *sql.MockDashboardsInterface {
	api, ok := m.WorkspaceClient.Dashboards.(*sql.MockDashboardsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDataSourcesAPI() *sql.MockDataSourcesInterface {
	api, ok := m.WorkspaceClient.DataSources.(*sql.MockDataSourcesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDbfsAPI() *files.MockDbfsInterface {
	api, ok := m.WorkspaceClient.Dbfs.(*files.MockDbfsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDbsqlPermissionsAPI() *sql.MockDbsqlPermissionsInterface {
	api, ok := m.WorkspaceClient.DbsqlPermissions.(*sql.MockDbsqlPermissionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExperimentsAPI() *ml.MockExperimentsInterface {
	api, ok := m.WorkspaceClient.Experiments.(*ml.MockExperimentsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExternalLocationsAPI() *catalog.MockExternalLocationsInterface {
	api, ok := m.WorkspaceClient.ExternalLocations.(*catalog.MockExternalLocationsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFilesAPI() *files.MockFilesInterface {
	api, ok := m.WorkspaceClient.Files.(*files.MockFilesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFunctionsAPI() *catalog.MockFunctionsInterface {
	api, ok := m.WorkspaceClient.Functions.(*catalog.MockFunctionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGitCredentialsAPI() *workspace.MockGitCredentialsInterface {
	api, ok := m.WorkspaceClient.GitCredentials.(*workspace.MockGitCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGlobalInitScriptsAPI() *compute.MockGlobalInitScriptsInterface {
	api, ok := m.WorkspaceClient.GlobalInitScripts.(*compute.MockGlobalInitScriptsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGrantsAPI() *catalog.MockGrantsInterface {
	api, ok := m.WorkspaceClient.Grants.(*catalog.MockGrantsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGroupsAPI() *iam.MockGroupsInterface {
	api, ok := m.WorkspaceClient.Groups.(*iam.MockGroupsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockInstancePoolsAPI() *compute.MockInstancePoolsInterface {
	api, ok := m.WorkspaceClient.InstancePools.(*compute.MockInstancePoolsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockInstanceProfilesAPI() *compute.MockInstanceProfilesInterface {
	api, ok := m.WorkspaceClient.InstanceProfiles.(*compute.MockInstanceProfilesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockIpAccessListsAPI() *settings.MockIpAccessListsInterface {
	api, ok := m.WorkspaceClient.IpAccessLists.(*settings.MockIpAccessListsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockJobsAPI() *jobs.MockJobsInterface {
	api, ok := m.WorkspaceClient.Jobs.(*jobs.MockJobsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLibrariesAPI() *compute.MockLibrariesInterface {
	api, ok := m.WorkspaceClient.Libraries.(*compute.MockLibrariesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockMetastoresAPI() *catalog.MockMetastoresInterface {
	api, ok := m.WorkspaceClient.Metastores.(*catalog.MockMetastoresInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockModelRegistryAPI() *ml.MockModelRegistryInterface {
	api, ok := m.WorkspaceClient.ModelRegistry.(*ml.MockModelRegistryInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockModelVersionsAPI() *catalog.MockModelVersionsInterface {
	api, ok := m.WorkspaceClient.ModelVersions.(*catalog.MockModelVersionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPermissionsAPI() *iam.MockPermissionsInterface {
	api, ok := m.WorkspaceClient.Permissions.(*iam.MockPermissionsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPipelinesAPI() *pipelines.MockPipelinesInterface {
	api, ok := m.WorkspaceClient.Pipelines.(*pipelines.MockPipelinesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPolicyFamiliesAPI() *compute.MockPolicyFamiliesInterface {
	api, ok := m.WorkspaceClient.PolicyFamilies.(*compute.MockPolicyFamiliesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProvidersAPI() *sharing.MockProvidersInterface {
	api, ok := m.WorkspaceClient.Providers.(*sharing.MockProvidersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueriesAPI() *sql.MockQueriesInterface {
	api, ok := m.WorkspaceClient.Queries.(*sql.MockQueriesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryHistoryAPI() *sql.MockQueryHistoryInterface {
	api, ok := m.WorkspaceClient.QueryHistory.(*sql.MockQueryHistoryInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryVisualizationsAPI() *sql.MockQueryVisualizationsInterface {
	api, ok := m.WorkspaceClient.QueryVisualizations.(*sql.MockQueryVisualizationsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientActivationAPI() *sharing.MockRecipientActivationInterface {
	api, ok := m.WorkspaceClient.RecipientActivation.(*sharing.MockRecipientActivationInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientsAPI() *sharing.MockRecipientsInterface {
	api, ok := m.WorkspaceClient.Recipients.(*sharing.MockRecipientsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRegisteredModelsAPI() *catalog.MockRegisteredModelsInterface {
	api, ok := m.WorkspaceClient.RegisteredModels.(*catalog.MockRegisteredModelsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockReposAPI() *workspace.MockReposInterface {
	api, ok := m.WorkspaceClient.Repos.(*workspace.MockReposInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSchemasAPI() *catalog.MockSchemasInterface {
	api, ok := m.WorkspaceClient.Schemas.(*catalog.MockSchemasInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSecretsAPI() *workspace.MockSecretsInterface {
	api, ok := m.WorkspaceClient.Secrets.(*workspace.MockSecretsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServicePrincipalsAPI() *iam.MockServicePrincipalsInterface {
	api, ok := m.WorkspaceClient.ServicePrincipals.(*iam.MockServicePrincipalsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServingEndpointsAPI() *serving.MockServingEndpointsInterface {
	api, ok := m.WorkspaceClient.ServingEndpoints.(*serving.MockServingEndpointsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSettingsAPI() *settings.MockSettingsInterface {
	api, ok := m.WorkspaceClient.Settings.(*settings.MockSettingsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSharesAPI() *sharing.MockSharesInterface {
	api, ok := m.WorkspaceClient.Shares.(*sharing.MockSharesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStatementExecutionAPI() *sql.MockStatementExecutionInterface {
	api, ok := m.WorkspaceClient.StatementExecution.(*sql.MockStatementExecutionInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStorageCredentialsAPI() *catalog.MockStorageCredentialsInterface {
	api, ok := m.WorkspaceClient.StorageCredentials.(*catalog.MockStorageCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSystemSchemasAPI() *catalog.MockSystemSchemasInterface {
	api, ok := m.WorkspaceClient.SystemSchemas.(*catalog.MockSystemSchemasInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTableConstraintsAPI() *catalog.MockTableConstraintsInterface {
	api, ok := m.WorkspaceClient.TableConstraints.(*catalog.MockTableConstraintsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTablesAPI() *catalog.MockTablesInterface {
	api, ok := m.WorkspaceClient.Tables.(*catalog.MockTablesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTokenManagementAPI() *settings.MockTokenManagementInterface {
	api, ok := m.WorkspaceClient.TokenManagement.(*settings.MockTokenManagementInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTokensAPI() *settings.MockTokensInterface {
	api, ok := m.WorkspaceClient.Tokens.(*settings.MockTokensInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockUsersAPI() *iam.MockUsersInterface {
	api, ok := m.WorkspaceClient.Users.(*iam.MockUsersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVolumesAPI() *catalog.MockVolumesInterface {
	api, ok := m.WorkspaceClient.Volumes.(*catalog.MockVolumesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWarehousesAPI() *sql.MockWarehousesInterface {
	api, ok := m.WorkspaceClient.Warehouses.(*sql.MockWarehousesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceAPI() *workspace.MockWorkspaceInterface {
	api, ok := m.WorkspaceClient.Workspace.(*workspace.MockWorkspaceInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceBindingsAPI() *catalog.MockWorkspaceBindingsInterface {
	api, ok := m.WorkspaceClient.WorkspaceBindings.(*catalog.MockWorkspaceBindingsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceConfAPI() *settings.MockWorkspaceConfInterface {
	api, ok := m.WorkspaceClient.WorkspaceConf.(*settings.MockWorkspaceConfInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
