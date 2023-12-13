// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/compute"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/files"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/jobs"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/ml"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/pipelines"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/serving"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/settings"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/sharing"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/sql"
	"github.com/databricks/databricks-sdk-go/internal/mocks/service/workspace"
)

// NewWorkspaceClient creates new mocked version of Databricks SDK client for Workspaces
// which can be used for testing.
func NewWorkspaceClient(t interface {
	mock.TestingT
	Cleanup(func())
}, cfg *config.Config) (*databricks.WorkspaceClient, error) {
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
	}, nil
}
