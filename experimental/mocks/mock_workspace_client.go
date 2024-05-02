// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/compute"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/dashboards"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/files"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/jobs"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/marketplace"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/ml"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/pipelines"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/serving"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settings"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sharing"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/sql"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/vectorsearch"
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
	cli := &MockWorkspaceClient{
		WorkspaceClient: &databricks.WorkspaceClient{
			Config: nil,

			AccountAccessControlProxy:           iam.NewMockAccountAccessControlProxyInterface(t),
			Alerts:                              sql.NewMockAlertsInterface(t),
			Apps:                                serving.NewMockAppsInterface(t),
			ArtifactAllowlists:                  catalog.NewMockArtifactAllowlistsInterface(t),
			Catalogs:                            catalog.NewMockCatalogsInterface(t),
			CleanRooms:                          sharing.NewMockCleanRoomsInterface(t),
			ClusterPolicies:                     compute.NewMockClusterPoliciesInterface(t),
			Clusters:                            compute.NewMockClustersInterface(t),
			CommandExecution:                    compute.NewMockCommandExecutionInterface(t),
			Connections:                         catalog.NewMockConnectionsInterface(t),
			ConsumerFulfillments:                marketplace.NewMockConsumerFulfillmentsInterface(t),
			ConsumerInstallations:               marketplace.NewMockConsumerInstallationsInterface(t),
			ConsumerListings:                    marketplace.NewMockConsumerListingsInterface(t),
			ConsumerPersonalizationRequests:     marketplace.NewMockConsumerPersonalizationRequestsInterface(t),
			ConsumerProviders:                   marketplace.NewMockConsumerProvidersInterface(t),
			CredentialsManager:                  settings.NewMockCredentialsManagerInterface(t),
			CurrentUser:                         iam.NewMockCurrentUserInterface(t),
			DashboardWidgets:                    sql.NewMockDashboardWidgetsInterface(t),
			Dashboards:                          sql.NewMockDashboardsInterface(t),
			DataSources:                         sql.NewMockDataSourcesInterface(t),
			Dbfs:                                files.NewMockDbfsInterface(t),
			DbsqlPermissions:                    sql.NewMockDbsqlPermissionsInterface(t),
			Experiments:                         ml.NewMockExperimentsInterface(t),
			ExternalLocations:                   catalog.NewMockExternalLocationsInterface(t),
			Files:                               files.NewMockFilesInterface(t),
			Functions:                           catalog.NewMockFunctionsInterface(t),
			GitCredentials:                      workspace.NewMockGitCredentialsInterface(t),
			GlobalInitScripts:                   compute.NewMockGlobalInitScriptsInterface(t),
			Grants:                              catalog.NewMockGrantsInterface(t),
			Groups:                              iam.NewMockGroupsInterface(t),
			InstancePools:                       compute.NewMockInstancePoolsInterface(t),
			InstanceProfiles:                    compute.NewMockInstanceProfilesInterface(t),
			IpAccessLists:                       settings.NewMockIpAccessListsInterface(t),
			Jobs:                                jobs.NewMockJobsInterface(t),
			LakehouseMonitors:                   catalog.NewMockLakehouseMonitorsInterface(t),
			Lakeview:                            dashboards.NewMockLakeviewInterface(t),
			Libraries:                           compute.NewMockLibrariesInterface(t),
			Metastores:                          catalog.NewMockMetastoresInterface(t),
			ModelRegistry:                       ml.NewMockModelRegistryInterface(t),
			ModelVersions:                       catalog.NewMockModelVersionsInterface(t),
			OnlineTables:                        catalog.NewMockOnlineTablesInterface(t),
			PermissionMigration:                 iam.NewMockPermissionMigrationInterface(t),
			Permissions:                         iam.NewMockPermissionsInterface(t),
			Pipelines:                           pipelines.NewMockPipelinesInterface(t),
			PolicyFamilies:                      compute.NewMockPolicyFamiliesInterface(t),
			ProviderExchangeFilters:             marketplace.NewMockProviderExchangeFiltersInterface(t),
			ProviderExchanges:                   marketplace.NewMockProviderExchangesInterface(t),
			ProviderFiles:                       marketplace.NewMockProviderFilesInterface(t),
			ProviderListings:                    marketplace.NewMockProviderListingsInterface(t),
			ProviderPersonalizationRequests:     marketplace.NewMockProviderPersonalizationRequestsInterface(t),
			ProviderProviderAnalyticsDashboards: marketplace.NewMockProviderProviderAnalyticsDashboardsInterface(t),
			ProviderProviders:                   marketplace.NewMockProviderProvidersInterface(t),
			Providers:                           sharing.NewMockProvidersInterface(t),
			Queries:                             sql.NewMockQueriesInterface(t),
			QueryHistory:                        sql.NewMockQueryHistoryInterface(t),
			QueryVisualizations:                 sql.NewMockQueryVisualizationsInterface(t),
			RecipientActivation:                 sharing.NewMockRecipientActivationInterface(t),
			Recipients:                          sharing.NewMockRecipientsInterface(t),
			RegisteredModels:                    catalog.NewMockRegisteredModelsInterface(t),
			Repos:                               workspace.NewMockReposInterface(t),
			Schemas:                             catalog.NewMockSchemasInterface(t),
			Secrets:                             workspace.NewMockSecretsInterface(t),
			ServicePrincipals:                   iam.NewMockServicePrincipalsInterface(t),
			ServingEndpoints:                    serving.NewMockServingEndpointsInterface(t),
			Settings:                            settings.NewMockSettingsInterface(t),
			Shares:                              sharing.NewMockSharesInterface(t),
			StatementExecution:                  sql.NewMockStatementExecutionInterface(t),
			StorageCredentials:                  catalog.NewMockStorageCredentialsInterface(t),
			SystemSchemas:                       catalog.NewMockSystemSchemasInterface(t),
			TableConstraints:                    catalog.NewMockTableConstraintsInterface(t),
			Tables:                              catalog.NewMockTablesInterface(t),
			TokenManagement:                     settings.NewMockTokenManagementInterface(t),
			Tokens:                              settings.NewMockTokensInterface(t),
			Users:                               iam.NewMockUsersInterface(t),
			VectorSearchEndpoints:               vectorsearch.NewMockVectorSearchEndpointsInterface(t),
			VectorSearchIndexes:                 vectorsearch.NewMockVectorSearchIndexesInterface(t),
			Volumes:                             catalog.NewMockVolumesInterface(t),
			Warehouses:                          sql.NewMockWarehousesInterface(t),
			Workspace:                           workspace.NewMockWorkspaceInterface(t),
			WorkspaceBindings:                   catalog.NewMockWorkspaceBindingsInterface(t),
			WorkspaceConf:                       settings.NewMockWorkspaceConfInterface(t),
		},
	}

	mocksettingsAPI := cli.GetMockSettingsAPI()

	mockautomaticClusterUpdate := settings.NewMockAutomaticClusterUpdateInterface(t)
	mocksettingsAPI.On("AutomaticClusterUpdate").Return(mockautomaticClusterUpdate).Maybe()

	mockcomplianceSecurityProfile := settings.NewMockComplianceSecurityProfileInterface(t)
	mocksettingsAPI.On("ComplianceSecurityProfile").Return(mockcomplianceSecurityProfile).Maybe()

	mockdefaultNamespace := settings.NewMockDefaultNamespaceInterface(t)
	mocksettingsAPI.On("DefaultNamespace").Return(mockdefaultNamespace).Maybe()

	mockenhancedSecurityMonitoring := settings.NewMockEnhancedSecurityMonitoringInterface(t)
	mocksettingsAPI.On("EnhancedSecurityMonitoring").Return(mockenhancedSecurityMonitoring).Maybe()

	mockrestrictWorkspaceAdmins := settings.NewMockRestrictWorkspaceAdminsInterface(t)
	mocksettingsAPI.On("RestrictWorkspaceAdmins").Return(mockrestrictWorkspaceAdmins).Maybe()

	return cli
}

func (m *MockWorkspaceClient) GetMockAutomaticClusterUpdateAPI() *settings.MockAutomaticClusterUpdateInterface {
	api, ok := m.GetMockSettingsAPI().AutomaticClusterUpdate().(*settings.MockAutomaticClusterUpdateInterface)
	if !ok {
		panic(fmt.Sprintf("expected AutomaticClusterUpdate to be *settings.MockAutomaticClusterUpdateInterface, actual was %T", m.GetMockSettingsAPI().AutomaticClusterUpdate()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockComplianceSecurityProfileAPI() *settings.MockComplianceSecurityProfileInterface {
	api, ok := m.GetMockSettingsAPI().ComplianceSecurityProfile().(*settings.MockComplianceSecurityProfileInterface)
	if !ok {
		panic(fmt.Sprintf("expected ComplianceSecurityProfile to be *settings.MockComplianceSecurityProfileInterface, actual was %T", m.GetMockSettingsAPI().ComplianceSecurityProfile()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDefaultNamespaceAPI() *settings.MockDefaultNamespaceInterface {
	api, ok := m.GetMockSettingsAPI().DefaultNamespace().(*settings.MockDefaultNamespaceInterface)
	if !ok {
		panic(fmt.Sprintf("expected DefaultNamespace to be *settings.MockDefaultNamespaceInterface, actual was %T", m.GetMockSettingsAPI().DefaultNamespace()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockEnhancedSecurityMonitoringAPI() *settings.MockEnhancedSecurityMonitoringInterface {
	api, ok := m.GetMockSettingsAPI().EnhancedSecurityMonitoring().(*settings.MockEnhancedSecurityMonitoringInterface)
	if !ok {
		panic(fmt.Sprintf("expected EnhancedSecurityMonitoring to be *settings.MockEnhancedSecurityMonitoringInterface, actual was %T", m.GetMockSettingsAPI().EnhancedSecurityMonitoring()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRestrictWorkspaceAdminsAPI() *settings.MockRestrictWorkspaceAdminsInterface {
	api, ok := m.GetMockSettingsAPI().RestrictWorkspaceAdmins().(*settings.MockRestrictWorkspaceAdminsInterface)
	if !ok {
		panic(fmt.Sprintf("expected RestrictWorkspaceAdmins to be *settings.MockRestrictWorkspaceAdminsInterface, actual was %T", m.GetMockSettingsAPI().RestrictWorkspaceAdmins()))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAccountAccessControlProxyAPI() *iam.MockAccountAccessControlProxyInterface {
	api, ok := m.WorkspaceClient.AccountAccessControlProxy.(*iam.MockAccountAccessControlProxyInterface)
	if !ok {
		panic(fmt.Sprintf("expected AccountAccessControlProxy to be *iam.MockAccountAccessControlProxyInterface, actual was %T", m.WorkspaceClient.AccountAccessControlProxy))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAlertsAPI() *sql.MockAlertsInterface {
	api, ok := m.WorkspaceClient.Alerts.(*sql.MockAlertsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Alerts to be *sql.MockAlertsInterface, actual was %T", m.WorkspaceClient.Alerts))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockAppsAPI() *serving.MockAppsInterface {
	api, ok := m.WorkspaceClient.Apps.(*serving.MockAppsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Apps to be *serving.MockAppsInterface, actual was %T", m.WorkspaceClient.Apps))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockArtifactAllowlistsAPI() *catalog.MockArtifactAllowlistsInterface {
	api, ok := m.WorkspaceClient.ArtifactAllowlists.(*catalog.MockArtifactAllowlistsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ArtifactAllowlists to be *catalog.MockArtifactAllowlistsInterface, actual was %T", m.WorkspaceClient.ArtifactAllowlists))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCatalogsAPI() *catalog.MockCatalogsInterface {
	api, ok := m.WorkspaceClient.Catalogs.(*catalog.MockCatalogsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Catalogs to be *catalog.MockCatalogsInterface, actual was %T", m.WorkspaceClient.Catalogs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCleanRoomsAPI() *sharing.MockCleanRoomsInterface {
	api, ok := m.WorkspaceClient.CleanRooms.(*sharing.MockCleanRoomsInterface)
	if !ok {
		panic(fmt.Sprintf("expected CleanRooms to be *sharing.MockCleanRoomsInterface, actual was %T", m.WorkspaceClient.CleanRooms))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockClusterPoliciesAPI() *compute.MockClusterPoliciesInterface {
	api, ok := m.WorkspaceClient.ClusterPolicies.(*compute.MockClusterPoliciesInterface)
	if !ok {
		panic(fmt.Sprintf("expected ClusterPolicies to be *compute.MockClusterPoliciesInterface, actual was %T", m.WorkspaceClient.ClusterPolicies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockClustersAPI() *compute.MockClustersInterface {
	api, ok := m.WorkspaceClient.Clusters.(*compute.MockClustersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Clusters to be *compute.MockClustersInterface, actual was %T", m.WorkspaceClient.Clusters))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCommandExecutionAPI() *compute.MockCommandExecutionInterface {
	api, ok := m.WorkspaceClient.CommandExecution.(*compute.MockCommandExecutionInterface)
	if !ok {
		panic(fmt.Sprintf("expected CommandExecution to be *compute.MockCommandExecutionInterface, actual was %T", m.WorkspaceClient.CommandExecution))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConnectionsAPI() *catalog.MockConnectionsInterface {
	api, ok := m.WorkspaceClient.Connections.(*catalog.MockConnectionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Connections to be *catalog.MockConnectionsInterface, actual was %T", m.WorkspaceClient.Connections))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerFulfillmentsAPI() *marketplace.MockConsumerFulfillmentsInterface {
	api, ok := m.WorkspaceClient.ConsumerFulfillments.(*marketplace.MockConsumerFulfillmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerFulfillments to be *marketplace.MockConsumerFulfillmentsInterface, actual was %T", m.WorkspaceClient.ConsumerFulfillments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerInstallationsAPI() *marketplace.MockConsumerInstallationsInterface {
	api, ok := m.WorkspaceClient.ConsumerInstallations.(*marketplace.MockConsumerInstallationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerInstallations to be *marketplace.MockConsumerInstallationsInterface, actual was %T", m.WorkspaceClient.ConsumerInstallations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerListingsAPI() *marketplace.MockConsumerListingsInterface {
	api, ok := m.WorkspaceClient.ConsumerListings.(*marketplace.MockConsumerListingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerListings to be *marketplace.MockConsumerListingsInterface, actual was %T", m.WorkspaceClient.ConsumerListings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerPersonalizationRequestsAPI() *marketplace.MockConsumerPersonalizationRequestsInterface {
	api, ok := m.WorkspaceClient.ConsumerPersonalizationRequests.(*marketplace.MockConsumerPersonalizationRequestsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerPersonalizationRequests to be *marketplace.MockConsumerPersonalizationRequestsInterface, actual was %T", m.WorkspaceClient.ConsumerPersonalizationRequests))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockConsumerProvidersAPI() *marketplace.MockConsumerProvidersInterface {
	api, ok := m.WorkspaceClient.ConsumerProviders.(*marketplace.MockConsumerProvidersInterface)
	if !ok {
		panic(fmt.Sprintf("expected ConsumerProviders to be *marketplace.MockConsumerProvidersInterface, actual was %T", m.WorkspaceClient.ConsumerProviders))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCredentialsManagerAPI() *settings.MockCredentialsManagerInterface {
	api, ok := m.WorkspaceClient.CredentialsManager.(*settings.MockCredentialsManagerInterface)
	if !ok {
		panic(fmt.Sprintf("expected CredentialsManager to be *settings.MockCredentialsManagerInterface, actual was %T", m.WorkspaceClient.CredentialsManager))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockCurrentUserAPI() *iam.MockCurrentUserInterface {
	api, ok := m.WorkspaceClient.CurrentUser.(*iam.MockCurrentUserInterface)
	if !ok {
		panic(fmt.Sprintf("expected CurrentUser to be *iam.MockCurrentUserInterface, actual was %T", m.WorkspaceClient.CurrentUser))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardWidgetsAPI() *sql.MockDashboardWidgetsInterface {
	api, ok := m.WorkspaceClient.DashboardWidgets.(*sql.MockDashboardWidgetsInterface)
	if !ok {
		panic(fmt.Sprintf("expected DashboardWidgets to be *sql.MockDashboardWidgetsInterface, actual was %T", m.WorkspaceClient.DashboardWidgets))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDashboardsAPI() *sql.MockDashboardsInterface {
	api, ok := m.WorkspaceClient.Dashboards.(*sql.MockDashboardsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Dashboards to be *sql.MockDashboardsInterface, actual was %T", m.WorkspaceClient.Dashboards))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDataSourcesAPI() *sql.MockDataSourcesInterface {
	api, ok := m.WorkspaceClient.DataSources.(*sql.MockDataSourcesInterface)
	if !ok {
		panic(fmt.Sprintf("expected DataSources to be *sql.MockDataSourcesInterface, actual was %T", m.WorkspaceClient.DataSources))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDbfsAPI() *files.MockDbfsInterface {
	api, ok := m.WorkspaceClient.Dbfs.(*files.MockDbfsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Dbfs to be *files.MockDbfsInterface, actual was %T", m.WorkspaceClient.Dbfs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockDbsqlPermissionsAPI() *sql.MockDbsqlPermissionsInterface {
	api, ok := m.WorkspaceClient.DbsqlPermissions.(*sql.MockDbsqlPermissionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected DbsqlPermissions to be *sql.MockDbsqlPermissionsInterface, actual was %T", m.WorkspaceClient.DbsqlPermissions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExperimentsAPI() *ml.MockExperimentsInterface {
	api, ok := m.WorkspaceClient.Experiments.(*ml.MockExperimentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Experiments to be *ml.MockExperimentsInterface, actual was %T", m.WorkspaceClient.Experiments))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockExternalLocationsAPI() *catalog.MockExternalLocationsInterface {
	api, ok := m.WorkspaceClient.ExternalLocations.(*catalog.MockExternalLocationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ExternalLocations to be *catalog.MockExternalLocationsInterface, actual was %T", m.WorkspaceClient.ExternalLocations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFilesAPI() *files.MockFilesInterface {
	api, ok := m.WorkspaceClient.Files.(*files.MockFilesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Files to be *files.MockFilesInterface, actual was %T", m.WorkspaceClient.Files))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockFunctionsAPI() *catalog.MockFunctionsInterface {
	api, ok := m.WorkspaceClient.Functions.(*catalog.MockFunctionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Functions to be *catalog.MockFunctionsInterface, actual was %T", m.WorkspaceClient.Functions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGitCredentialsAPI() *workspace.MockGitCredentialsInterface {
	api, ok := m.WorkspaceClient.GitCredentials.(*workspace.MockGitCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected GitCredentials to be *workspace.MockGitCredentialsInterface, actual was %T", m.WorkspaceClient.GitCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGlobalInitScriptsAPI() *compute.MockGlobalInitScriptsInterface {
	api, ok := m.WorkspaceClient.GlobalInitScripts.(*compute.MockGlobalInitScriptsInterface)
	if !ok {
		panic(fmt.Sprintf("expected GlobalInitScripts to be *compute.MockGlobalInitScriptsInterface, actual was %T", m.WorkspaceClient.GlobalInitScripts))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGrantsAPI() *catalog.MockGrantsInterface {
	api, ok := m.WorkspaceClient.Grants.(*catalog.MockGrantsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Grants to be *catalog.MockGrantsInterface, actual was %T", m.WorkspaceClient.Grants))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockGroupsAPI() *iam.MockGroupsInterface {
	api, ok := m.WorkspaceClient.Groups.(*iam.MockGroupsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Groups to be *iam.MockGroupsInterface, actual was %T", m.WorkspaceClient.Groups))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockInstancePoolsAPI() *compute.MockInstancePoolsInterface {
	api, ok := m.WorkspaceClient.InstancePools.(*compute.MockInstancePoolsInterface)
	if !ok {
		panic(fmt.Sprintf("expected InstancePools to be *compute.MockInstancePoolsInterface, actual was %T", m.WorkspaceClient.InstancePools))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockInstanceProfilesAPI() *compute.MockInstanceProfilesInterface {
	api, ok := m.WorkspaceClient.InstanceProfiles.(*compute.MockInstanceProfilesInterface)
	if !ok {
		panic(fmt.Sprintf("expected InstanceProfiles to be *compute.MockInstanceProfilesInterface, actual was %T", m.WorkspaceClient.InstanceProfiles))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockIpAccessListsAPI() *settings.MockIpAccessListsInterface {
	api, ok := m.WorkspaceClient.IpAccessLists.(*settings.MockIpAccessListsInterface)
	if !ok {
		panic(fmt.Sprintf("expected IpAccessLists to be *settings.MockIpAccessListsInterface, actual was %T", m.WorkspaceClient.IpAccessLists))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockJobsAPI() *jobs.MockJobsInterface {
	api, ok := m.WorkspaceClient.Jobs.(*jobs.MockJobsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Jobs to be *jobs.MockJobsInterface, actual was %T", m.WorkspaceClient.Jobs))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLakehouseMonitorsAPI() *catalog.MockLakehouseMonitorsInterface {
	api, ok := m.WorkspaceClient.LakehouseMonitors.(*catalog.MockLakehouseMonitorsInterface)
	if !ok {
		panic(fmt.Sprintf("expected LakehouseMonitors to be *catalog.MockLakehouseMonitorsInterface, actual was %T", m.WorkspaceClient.LakehouseMonitors))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLakeviewAPI() *dashboards.MockLakeviewInterface {
	api, ok := m.WorkspaceClient.Lakeview.(*dashboards.MockLakeviewInterface)
	if !ok {
		panic(fmt.Sprintf("expected Lakeview to be *dashboards.MockLakeviewInterface, actual was %T", m.WorkspaceClient.Lakeview))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockLibrariesAPI() *compute.MockLibrariesInterface {
	api, ok := m.WorkspaceClient.Libraries.(*compute.MockLibrariesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Libraries to be *compute.MockLibrariesInterface, actual was %T", m.WorkspaceClient.Libraries))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockMetastoresAPI() *catalog.MockMetastoresInterface {
	api, ok := m.WorkspaceClient.Metastores.(*catalog.MockMetastoresInterface)
	if !ok {
		panic(fmt.Sprintf("expected Metastores to be *catalog.MockMetastoresInterface, actual was %T", m.WorkspaceClient.Metastores))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockModelRegistryAPI() *ml.MockModelRegistryInterface {
	api, ok := m.WorkspaceClient.ModelRegistry.(*ml.MockModelRegistryInterface)
	if !ok {
		panic(fmt.Sprintf("expected ModelRegistry to be *ml.MockModelRegistryInterface, actual was %T", m.WorkspaceClient.ModelRegistry))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockModelVersionsAPI() *catalog.MockModelVersionsInterface {
	api, ok := m.WorkspaceClient.ModelVersions.(*catalog.MockModelVersionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ModelVersions to be *catalog.MockModelVersionsInterface, actual was %T", m.WorkspaceClient.ModelVersions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockOnlineTablesAPI() *catalog.MockOnlineTablesInterface {
	api, ok := m.WorkspaceClient.OnlineTables.(*catalog.MockOnlineTablesInterface)
	if !ok {
		panic(fmt.Sprintf("expected OnlineTables to be *catalog.MockOnlineTablesInterface, actual was %T", m.WorkspaceClient.OnlineTables))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPermissionMigrationAPI() *iam.MockPermissionMigrationInterface {
	api, ok := m.WorkspaceClient.PermissionMigration.(*iam.MockPermissionMigrationInterface)
	if !ok {
		panic(fmt.Sprintf("expected PermissionMigration to be *iam.MockPermissionMigrationInterface, actual was %T", m.WorkspaceClient.PermissionMigration))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPermissionsAPI() *iam.MockPermissionsInterface {
	api, ok := m.WorkspaceClient.Permissions.(*iam.MockPermissionsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Permissions to be *iam.MockPermissionsInterface, actual was %T", m.WorkspaceClient.Permissions))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPipelinesAPI() *pipelines.MockPipelinesInterface {
	api, ok := m.WorkspaceClient.Pipelines.(*pipelines.MockPipelinesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Pipelines to be *pipelines.MockPipelinesInterface, actual was %T", m.WorkspaceClient.Pipelines))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockPolicyFamiliesAPI() *compute.MockPolicyFamiliesInterface {
	api, ok := m.WorkspaceClient.PolicyFamilies.(*compute.MockPolicyFamiliesInterface)
	if !ok {
		panic(fmt.Sprintf("expected PolicyFamilies to be *compute.MockPolicyFamiliesInterface, actual was %T", m.WorkspaceClient.PolicyFamilies))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderExchangeFiltersAPI() *marketplace.MockProviderExchangeFiltersInterface {
	api, ok := m.WorkspaceClient.ProviderExchangeFilters.(*marketplace.MockProviderExchangeFiltersInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderExchangeFilters to be *marketplace.MockProviderExchangeFiltersInterface, actual was %T", m.WorkspaceClient.ProviderExchangeFilters))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderExchangesAPI() *marketplace.MockProviderExchangesInterface {
	api, ok := m.WorkspaceClient.ProviderExchanges.(*marketplace.MockProviderExchangesInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderExchanges to be *marketplace.MockProviderExchangesInterface, actual was %T", m.WorkspaceClient.ProviderExchanges))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderFilesAPI() *marketplace.MockProviderFilesInterface {
	api, ok := m.WorkspaceClient.ProviderFiles.(*marketplace.MockProviderFilesInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderFiles to be *marketplace.MockProviderFilesInterface, actual was %T", m.WorkspaceClient.ProviderFiles))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderListingsAPI() *marketplace.MockProviderListingsInterface {
	api, ok := m.WorkspaceClient.ProviderListings.(*marketplace.MockProviderListingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderListings to be *marketplace.MockProviderListingsInterface, actual was %T", m.WorkspaceClient.ProviderListings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderPersonalizationRequestsAPI() *marketplace.MockProviderPersonalizationRequestsInterface {
	api, ok := m.WorkspaceClient.ProviderPersonalizationRequests.(*marketplace.MockProviderPersonalizationRequestsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderPersonalizationRequests to be *marketplace.MockProviderPersonalizationRequestsInterface, actual was %T", m.WorkspaceClient.ProviderPersonalizationRequests))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderProviderAnalyticsDashboardsAPI() *marketplace.MockProviderProviderAnalyticsDashboardsInterface {
	api, ok := m.WorkspaceClient.ProviderProviderAnalyticsDashboards.(*marketplace.MockProviderProviderAnalyticsDashboardsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderProviderAnalyticsDashboards to be *marketplace.MockProviderProviderAnalyticsDashboardsInterface, actual was %T", m.WorkspaceClient.ProviderProviderAnalyticsDashboards))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProviderProvidersAPI() *marketplace.MockProviderProvidersInterface {
	api, ok := m.WorkspaceClient.ProviderProviders.(*marketplace.MockProviderProvidersInterface)
	if !ok {
		panic(fmt.Sprintf("expected ProviderProviders to be *marketplace.MockProviderProvidersInterface, actual was %T", m.WorkspaceClient.ProviderProviders))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockProvidersAPI() *sharing.MockProvidersInterface {
	api, ok := m.WorkspaceClient.Providers.(*sharing.MockProvidersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Providers to be *sharing.MockProvidersInterface, actual was %T", m.WorkspaceClient.Providers))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueriesAPI() *sql.MockQueriesInterface {
	api, ok := m.WorkspaceClient.Queries.(*sql.MockQueriesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Queries to be *sql.MockQueriesInterface, actual was %T", m.WorkspaceClient.Queries))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryHistoryAPI() *sql.MockQueryHistoryInterface {
	api, ok := m.WorkspaceClient.QueryHistory.(*sql.MockQueryHistoryInterface)
	if !ok {
		panic(fmt.Sprintf("expected QueryHistory to be *sql.MockQueryHistoryInterface, actual was %T", m.WorkspaceClient.QueryHistory))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockQueryVisualizationsAPI() *sql.MockQueryVisualizationsInterface {
	api, ok := m.WorkspaceClient.QueryVisualizations.(*sql.MockQueryVisualizationsInterface)
	if !ok {
		panic(fmt.Sprintf("expected QueryVisualizations to be *sql.MockQueryVisualizationsInterface, actual was %T", m.WorkspaceClient.QueryVisualizations))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientActivationAPI() *sharing.MockRecipientActivationInterface {
	api, ok := m.WorkspaceClient.RecipientActivation.(*sharing.MockRecipientActivationInterface)
	if !ok {
		panic(fmt.Sprintf("expected RecipientActivation to be *sharing.MockRecipientActivationInterface, actual was %T", m.WorkspaceClient.RecipientActivation))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRecipientsAPI() *sharing.MockRecipientsInterface {
	api, ok := m.WorkspaceClient.Recipients.(*sharing.MockRecipientsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Recipients to be *sharing.MockRecipientsInterface, actual was %T", m.WorkspaceClient.Recipients))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockRegisteredModelsAPI() *catalog.MockRegisteredModelsInterface {
	api, ok := m.WorkspaceClient.RegisteredModels.(*catalog.MockRegisteredModelsInterface)
	if !ok {
		panic(fmt.Sprintf("expected RegisteredModels to be *catalog.MockRegisteredModelsInterface, actual was %T", m.WorkspaceClient.RegisteredModels))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockReposAPI() *workspace.MockReposInterface {
	api, ok := m.WorkspaceClient.Repos.(*workspace.MockReposInterface)
	if !ok {
		panic(fmt.Sprintf("expected Repos to be *workspace.MockReposInterface, actual was %T", m.WorkspaceClient.Repos))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSchemasAPI() *catalog.MockSchemasInterface {
	api, ok := m.WorkspaceClient.Schemas.(*catalog.MockSchemasInterface)
	if !ok {
		panic(fmt.Sprintf("expected Schemas to be *catalog.MockSchemasInterface, actual was %T", m.WorkspaceClient.Schemas))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSecretsAPI() *workspace.MockSecretsInterface {
	api, ok := m.WorkspaceClient.Secrets.(*workspace.MockSecretsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Secrets to be *workspace.MockSecretsInterface, actual was %T", m.WorkspaceClient.Secrets))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServicePrincipalsAPI() *iam.MockServicePrincipalsInterface {
	api, ok := m.WorkspaceClient.ServicePrincipals.(*iam.MockServicePrincipalsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipals to be *iam.MockServicePrincipalsInterface, actual was %T", m.WorkspaceClient.ServicePrincipals))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockServingEndpointsAPI() *serving.MockServingEndpointsInterface {
	api, ok := m.WorkspaceClient.ServingEndpoints.(*serving.MockServingEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServingEndpoints to be *serving.MockServingEndpointsInterface, actual was %T", m.WorkspaceClient.ServingEndpoints))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSettingsAPI() *settings.MockSettingsInterface {
	api, ok := m.WorkspaceClient.Settings.(*settings.MockSettingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Settings to be *settings.MockSettingsInterface, actual was %T", m.WorkspaceClient.Settings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSharesAPI() *sharing.MockSharesInterface {
	api, ok := m.WorkspaceClient.Shares.(*sharing.MockSharesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Shares to be *sharing.MockSharesInterface, actual was %T", m.WorkspaceClient.Shares))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStatementExecutionAPI() *sql.MockStatementExecutionInterface {
	api, ok := m.WorkspaceClient.StatementExecution.(*sql.MockStatementExecutionInterface)
	if !ok {
		panic(fmt.Sprintf("expected StatementExecution to be *sql.MockStatementExecutionInterface, actual was %T", m.WorkspaceClient.StatementExecution))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockStorageCredentialsAPI() *catalog.MockStorageCredentialsInterface {
	api, ok := m.WorkspaceClient.StorageCredentials.(*catalog.MockStorageCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected StorageCredentials to be *catalog.MockStorageCredentialsInterface, actual was %T", m.WorkspaceClient.StorageCredentials))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockSystemSchemasAPI() *catalog.MockSystemSchemasInterface {
	api, ok := m.WorkspaceClient.SystemSchemas.(*catalog.MockSystemSchemasInterface)
	if !ok {
		panic(fmt.Sprintf("expected SystemSchemas to be *catalog.MockSystemSchemasInterface, actual was %T", m.WorkspaceClient.SystemSchemas))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTableConstraintsAPI() *catalog.MockTableConstraintsInterface {
	api, ok := m.WorkspaceClient.TableConstraints.(*catalog.MockTableConstraintsInterface)
	if !ok {
		panic(fmt.Sprintf("expected TableConstraints to be *catalog.MockTableConstraintsInterface, actual was %T", m.WorkspaceClient.TableConstraints))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTablesAPI() *catalog.MockTablesInterface {
	api, ok := m.WorkspaceClient.Tables.(*catalog.MockTablesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Tables to be *catalog.MockTablesInterface, actual was %T", m.WorkspaceClient.Tables))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTokenManagementAPI() *settings.MockTokenManagementInterface {
	api, ok := m.WorkspaceClient.TokenManagement.(*settings.MockTokenManagementInterface)
	if !ok {
		panic(fmt.Sprintf("expected TokenManagement to be *settings.MockTokenManagementInterface, actual was %T", m.WorkspaceClient.TokenManagement))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockTokensAPI() *settings.MockTokensInterface {
	api, ok := m.WorkspaceClient.Tokens.(*settings.MockTokensInterface)
	if !ok {
		panic(fmt.Sprintf("expected Tokens to be *settings.MockTokensInterface, actual was %T", m.WorkspaceClient.Tokens))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockUsersAPI() *iam.MockUsersInterface {
	api, ok := m.WorkspaceClient.Users.(*iam.MockUsersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Users to be *iam.MockUsersInterface, actual was %T", m.WorkspaceClient.Users))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVectorSearchEndpointsAPI() *vectorsearch.MockVectorSearchEndpointsInterface {
	api, ok := m.WorkspaceClient.VectorSearchEndpoints.(*vectorsearch.MockVectorSearchEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected VectorSearchEndpoints to be *vectorsearch.MockVectorSearchEndpointsInterface, actual was %T", m.WorkspaceClient.VectorSearchEndpoints))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVectorSearchIndexesAPI() *vectorsearch.MockVectorSearchIndexesInterface {
	api, ok := m.WorkspaceClient.VectorSearchIndexes.(*vectorsearch.MockVectorSearchIndexesInterface)
	if !ok {
		panic(fmt.Sprintf("expected VectorSearchIndexes to be *vectorsearch.MockVectorSearchIndexesInterface, actual was %T", m.WorkspaceClient.VectorSearchIndexes))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockVolumesAPI() *catalog.MockVolumesInterface {
	api, ok := m.WorkspaceClient.Volumes.(*catalog.MockVolumesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Volumes to be *catalog.MockVolumesInterface, actual was %T", m.WorkspaceClient.Volumes))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWarehousesAPI() *sql.MockWarehousesInterface {
	api, ok := m.WorkspaceClient.Warehouses.(*sql.MockWarehousesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Warehouses to be *sql.MockWarehousesInterface, actual was %T", m.WorkspaceClient.Warehouses))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceAPI() *workspace.MockWorkspaceInterface {
	api, ok := m.WorkspaceClient.Workspace.(*workspace.MockWorkspaceInterface)
	if !ok {
		panic(fmt.Sprintf("expected Workspace to be *workspace.MockWorkspaceInterface, actual was %T", m.WorkspaceClient.Workspace))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceBindingsAPI() *catalog.MockWorkspaceBindingsInterface {
	api, ok := m.WorkspaceClient.WorkspaceBindings.(*catalog.MockWorkspaceBindingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceBindings to be *catalog.MockWorkspaceBindingsInterface, actual was %T", m.WorkspaceClient.WorkspaceBindings))
	}
	return api
}

func (m *MockWorkspaceClient) GetMockWorkspaceConfAPI() *settings.MockWorkspaceConfInterface {
	api, ok := m.WorkspaceClient.WorkspaceConf.(*settings.MockWorkspaceConfInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceConf to be *settings.MockWorkspaceConfInterface, actual was %T", m.WorkspaceClient.WorkspaceConf))
	}
	return api
}
