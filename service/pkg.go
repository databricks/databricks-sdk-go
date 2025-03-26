// Databricks SDK for Go APIs
//
// - [iam.AccessControlAPI]: Rule based Access Control for Databricks Resources.
//
// - [iam.AccountAccessControlAPI]: These APIs manage access rules on resources in an account.
//
// - [iam.AccountAccessControlProxyAPI]: These APIs manage access rules on resources in an account.
//
// - [settings.AibiDashboardEmbeddingAccessPolicyAPI]: Controls whether AI/BI published dashboard embedding is enabled, conditionally enabled, or disabled at the workspace level.
//
// - [settings.AibiDashboardEmbeddingApprovedDomainsAPI]: Controls the list of domains approved to host the embedded AI/BI dashboards.
//
// - [sql.AlertsAPI]: The alerts API can be used to perform CRUD operations on alerts.
//
// - [sql.AlertsLegacyAPI]: The alerts API can be used to perform CRUD operations on alerts.
//
// - [apps.AppsAPI]: Apps run directly on a customer’s Databricks instance, integrate with their data, use and extend Databricks services, and enable users to interact through single sign-on.
//
// - [catalog.ArtifactAllowlistsAPI]: In Databricks Runtime 13.3 and above, you can add libraries and init scripts to the `allowlist` in UC so that users can leverage these artifacts on compute configured with shared access mode.
//
// - [settings.AutomaticClusterUpdateAPI]: Controls whether automatic cluster update is enabled for the current workspace.
//
// - [billing.BillableUsageAPI]: This API allows you to download billable usage logs for the specified account and date range.
//
// - [billing.BudgetPolicyAPI]: A service serves REST API about Budget policies.
//
// - [catalog.CatalogsAPI]: A catalog is the first layer of Unity Catalog’s three-level namespace.
//
// - [cleanrooms.CleanRoomAssetsAPI]: Clean room assets are data and code objects — Tables, volumes, and notebooks that are shared with the clean room.
//
// - [cleanrooms.CleanRoomTaskRunsAPI]: Clean room task runs are the executions of notebooks in a clean room.
//
// - [cleanrooms.CleanRoomsAPI]: A clean room uses Delta Sharing and serverless compute to provide a secure and privacy-protecting environment where multiple parties can work together on sensitive enterprise data without direct access to each other’s data.
//
// - [compute.ClusterPoliciesAPI]: You can use cluster policies to control users' ability to configure clusters based on a set of rules.
//
// - [compute.ClustersAPI]: The Clusters API allows you to create, start, edit, list, terminate, and delete clusters.
//
// - [compute.CommandExecutionAPI]: This API allows execution of Python, Scala, SQL, or R commands on running Databricks Clusters.
//
// - [settings.ComplianceSecurityProfileAPI]: Controls whether to enable the compliance security profile for the current workspace.
//
// - [catalog.ConnectionsAPI]: Connections allow for creating a connection to an external data source.
//
// - [marketplace.ConsumerFulfillmentsAPI]: Fulfillments are entities that allow consumers to preview installations.
//
// - [marketplace.ConsumerInstallationsAPI]: Installations are entities that allow consumers to interact with Databricks Marketplace listings.
//
// - [marketplace.ConsumerListingsAPI]: Listings are the core entities in the Marketplace.
//
// - [marketplace.ConsumerPersonalizationRequestsAPI]: Personalization Requests allow customers to interact with the individualized Marketplace listing flow.
//
// - [marketplace.ConsumerProvidersAPI]: Providers are the entities that publish listings to the Marketplace.
//
// - [catalog.CredentialsAPI]: A credential represents an authentication and authorization mechanism for accessing services on your cloud tenant.
//
// - [provisioning.CredentialsAPI]: These APIs manage credential configurations for this workspace.
//
// - [settings.CredentialsManagerAPI]: Credentials manager interacts with with Identity Providers to to perform token exchanges using stored credentials and refresh tokens.
//
// - [settings.CspEnablementAccountAPI]: The compliance security profile settings at the account level control whether to enable it for new workspaces.
//
// - [iam.CurrentUserAPI]: This API allows retrieving information about currently authenticated user or service principal.
//
// - [oauth2.CustomAppIntegrationAPI]: These APIs enable administrators to manage custom OAuth app integrations, which is required for adding/using Custom OAuth App Integration like Tableau Cloud for Databricks in AWS cloud.
//
// - [sql.DashboardWidgetsAPI]: This is an evolving API that facilitates the addition and removal of widgets from existing dashboards within the Databricks Workspace.
//
// - [sql.DashboardsAPI]: In general, there is little need to modify dashboards using the API.
//
// - [sql.DataSourcesAPI]: This API is provided to assist you in making new query objects.
//
// - [files.DbfsAPI]: DBFS API makes it simple to interact with various data sources without having to include a users credentials every time to read a file.
//
// - [sql.DbsqlPermissionsAPI]: The SQL Permissions API is similar to the endpoints of the :method:permissions/set.
//
// - [settings.DefaultNamespaceAPI]: The default namespace setting API allows users to configure the default namespace for a Databricks workspace.
//
// - [settings.DisableLegacyAccessAPI]: 'Disabling legacy access' has the following impacts: 1.
//
// - [settings.DisableLegacyDbfsAPI]: When this setting is on, access to DBFS root and DBFS mounts is disallowed (as well as creation of new mounts).
//
// - [settings.DisableLegacyFeaturesAPI]: Disable legacy features for new Databricks workspaces.
//
// - [settings.EnableIpAccessListsAPI]: Controls the enforcement of IP access lists for accessing the account console.
//
// - [provisioning.EncryptionKeysAPI]: These APIs manage encryption key configurations for this workspace (optional).
//
// - [settings.EnhancedSecurityMonitoringAPI]: Controls whether enhanced security monitoring is enabled for the current workspace.
//
// - [settings.EsmEnablementAccountAPI]: The enhanced security monitoring setting at the account level controls whether to enable the feature on new workspaces.
//
// - [ml.ExperimentsAPI]: Experiments are the primary unit of organization in MLflow; all MLflow runs belong to an experiment.
//
// - [catalog.ExternalLocationsAPI]: An external location is an object that combines a cloud storage path with a storage credential that authorizes access to the cloud storage path.
//
// - [oauth2.AccountFederationPolicyAPI]: These APIs manage account federation policies.
//
// - [files.FilesAPI]: The Files API is a standard HTTP API that allows you to read, write, list, and delete files and directories by referring to their URI.
//
// - [catalog.FunctionsAPI]: Functions implement User-Defined Functions (UDFs) in Unity Catalog.
//
// - [dashboards.GenieAPI]: Genie provides a no-code experience for business users, powered by AI/BI.
//
// - [workspace.GitCredentialsAPI]: Registers personal access token for Databricks to do operations on behalf of the user.
//
// - [compute.GlobalInitScriptsAPI]: The Global Init Scripts API enables Workspace administrators to configure global initialization scripts for their workspace.
//
// - [catalog.GrantsAPI]: In Unity Catalog, data is secure by default.
//
// - [iam.GroupsAPI]: Groups simplify identity management, making it easier to assign access to Databricks workspace, data, and other securable objects.
//
// - [iam.AccountGroupsAPI]: Groups simplify identity management, making it easier to assign access to Databricks account, data, and other securable objects.
//
// - [compute.InstancePoolsAPI]: Instance Pools API are used to create, edit, delete and list instance pools by using ready-to-use cloud instances which reduces a cluster start and auto-scaling times.
//
// - [compute.InstanceProfilesAPI]: The Instance Profiles API allows admins to add, list, and remove instance profiles that users can launch clusters with.
//
// - [settings.IpAccessListsAPI]: IP Access List enables admins to configure IP access lists.
//
// - [settings.AccountIpAccessListsAPI]: The Accounts IP Access List API enables account admins to configure IP access lists for access to the account console.
//
// - [jobs.JobsAPI]: The Jobs API allows you to create, edit, and delete jobs.
//
// - [dashboards.LakeviewAPI]: These APIs provide specific management operations for Lakeview dashboards.
//
// - [dashboards.LakeviewEmbeddedAPI]: Token-based Lakeview APIs for embedding dashboards in external applications.
//
// - [compute.LibrariesAPI]: The Libraries API allows you to install and uninstall libraries and get the status of libraries on a cluster.
//
// - [billing.LogDeliveryAPI]: These APIs manage log delivery configurations for this account.
//
// - [catalog.AccountMetastoreAssignmentsAPI]: These APIs manage metastore assignments to a workspace.
//
// - [catalog.MetastoresAPI]: A metastore is the top-level container of objects in Unity Catalog.
//
// - [catalog.AccountMetastoresAPI]: These APIs manage Unity Catalog metastores for an account.
//
// - [ml.ModelRegistryAPI]: Note: This API reference documents APIs for the Workspace Model Registry.
//
// - [catalog.ModelVersionsAPI]: Databricks provides a hosted version of MLflow Model Registry in Unity Catalog.
//
// - [settings.NetworkConnectivityAPI]: These APIs provide configurations for the network connectivity of your workspaces for serverless compute resources.
//
// - [provisioning.NetworksAPI]: These APIs manage network configurations for customer-managed VPCs (optional).
//
// - [settings.NotificationDestinationsAPI]: The notification destinations API lets you programmatically manage a workspace's notification destinations.
//
// - [oauth2.OAuthPublishedAppsAPI]: These APIs enable administrators to view all the available published OAuth applications in Databricks.
//
// - [catalog.OnlineTablesAPI]: Online tables provide lower latency and higher QPS access to data from Delta tables.
//
// - [iam.PermissionMigrationAPI]: APIs for migrating acl permissions, used only by the ucx tool: https://github.com/databrickslabs/ucx.
//
// - [iam.PermissionsAPI]: Permissions API are used to create read, write, edit, update and manage access for various users on different objects and endpoints.
//
// - [settings.PersonalComputeAPI]: The Personal Compute enablement setting lets you control which users can use the Personal Compute default policy to create compute resources.
//
// - [pipelines.PipelinesAPI]: The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.
//
// - [compute.PolicyComplianceForClustersAPI]: The policy compliance APIs allow you to view and manage the policy compliance status of clusters in your workspace.
//
// - [jobs.PolicyComplianceForJobsAPI]: The compliance APIs allow you to view and manage the policy compliance status of jobs in your workspace.
//
// - [compute.PolicyFamiliesAPI]: View available policy families.
//
// - [provisioning.PrivateAccessAPI]: These APIs manage private access settings for this account.
//
// - [marketplace.ProviderExchangeFiltersAPI]: Marketplace exchanges filters curate which groups can access an exchange.
//
// - [marketplace.ProviderExchangesAPI]: Marketplace exchanges allow providers to share their listings with a curated set of customers.
//
// - [marketplace.ProviderFilesAPI]: Marketplace offers a set of file APIs for various purposes such as preview notebooks and provider icons.
//
// - [marketplace.ProviderListingsAPI]: Listings are the core entities in the Marketplace.
//
// - [marketplace.ProviderPersonalizationRequestsAPI]: Personalization requests are an alternate to instantly available listings.
//
// - [marketplace.ProviderProviderAnalyticsDashboardsAPI]: Manage templated analytics solution for providers.
//
// - [marketplace.ProviderProvidersAPI]: Providers are entities that manage assets in Marketplace.
//
// - [sharing.ProvidersAPI]: A data provider is an object representing the organization in the real world who shares the data.
//
// - [oauth2.PublishedAppIntegrationAPI]: These APIs enable administrators to manage published OAuth app integrations, which is required for adding/using Published OAuth App Integration like Tableau Desktop for Databricks in AWS cloud.
//
// - [catalog.QualityMonitorsAPI]: A monitor computes and monitors data or model quality metrics for a table over time.
//
// - [sql.QueriesAPI]: The queries API can be used to perform CRUD operations on queries.
//
// - [sql.QueriesLegacyAPI]: These endpoints are used for CRUD operations on query definitions.
//
// - [dashboards.QueryExecutionAPI]: Query execution APIs for AI / BI Dashboards.
//
// - [sql.QueryHistoryAPI]: A service responsible for storing and retrieving the list of queries run against SQL endpoints and serverless compute.
//
// - [sql.QueryVisualizationsAPI]: This is an evolving API that facilitates the addition and removal of visualizations from existing queries in the Databricks Workspace.
//
// - [sql.QueryVisualizationsLegacyAPI]: This is an evolving API that facilitates the addition and removal of vizualisations from existing queries within the Databricks Workspace.
//
// - [sharing.RecipientActivationAPI]: The Recipient Activation API is only applicable in the open sharing model where the recipient object has the authentication type of `TOKEN`.
//
// - [sharing.RecipientsAPI]: A recipient is an object you create using :method:recipients/create to represent an organization which you want to allow access shares.
//
// - [sql.RedashConfigAPI]: Redash V2 service for workspace configurations (internal).
//
// - [catalog.RegisteredModelsAPI]: Databricks provides a hosted version of MLflow Model Registry in Unity Catalog.
//
// - [workspace.ReposAPI]: The Repos API allows users to manage their git repos.
//
// - [catalog.ResourceQuotasAPI]: Unity Catalog enforces resource quotas on all securable objects, which limits the number of resources that can be created.
//
// - [settings.RestrictWorkspaceAdminsAPI]: The Restrict Workspace Admins setting lets you control the capabilities of workspace admins.
//
// - [catalog.SchemasAPI]: A schema (also called a database) is the second layer of Unity Catalog’s three-level namespace.
//
// - [workspace.SecretsAPI]: The Secrets API allows you to manage secrets, secret scopes, and access permissions.
//
// - [oauth2.ServicePrincipalFederationPolicyAPI]: These APIs manage service principal federation policies.
//
// - [oauth2.ServicePrincipalSecretsAPI]: These APIs enable administrators to manage service principal secrets.
//
// - [iam.ServicePrincipalsAPI]: Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.
//
// - [iam.AccountServicePrincipalsAPI]: Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.
//
// - [serving.ServingEndpointsAPI]: The Serving Endpoints API allows you to create, update, and delete model serving endpoints.
//
// - [serving.ServingEndpointsDataPlaneAPI]: Serving endpoints DataPlane provides a set of operations to interact with data plane endpoints for Serving endpoints service.
//
// - [settings.SettingsAPI]: Workspace Settings API allows users to manage settings at the workspace level.
//
// - [settings.AccountSettingsAPI]: Accounts Settings API allows users to manage settings at the account level.
//
// - [sharing.SharesAPI]: A share is a container instantiated with :method:shares/create.
//
// - [sql.StatementExecutionAPI]: The Databricks SQL Statement Execution API can be used to execute SQL statements on a SQL warehouse and fetch the result.
//
// - [provisioning.StorageAPI]: These APIs manage storage configurations for this workspace.
//
// - [catalog.StorageCredentialsAPI]: A storage credential represents an authentication and authorization mechanism for accessing data stored on your cloud tenant.
//
// - [catalog.AccountStorageCredentialsAPI]: These APIs manage storage credentials for a particular metastore.
//
// - [catalog.SystemSchemasAPI]: A system schema is a schema that lives within the system catalog.
//
// - [catalog.TableConstraintsAPI]: Primary key and foreign key constraints encode relationships between fields in tables.
//
// - [catalog.TablesAPI]: A table resides in the third layer of Unity Catalog’s three-level namespace.
//
// - [catalog.TemporaryTableCredentialsAPI]: Temporary Table Credentials refer to short-lived, downscoped credentials used to access cloud storage locationswhere table data is stored in Databricks.
//
// - [settings.TokenManagementAPI]: Enables administrators to get all tokens and delete tokens for other users.
//
// - [settings.TokensAPI]: The Token API allows you to create, list, and revoke tokens that can be used to authenticate and access Databricks REST APIs.
//
// - [billing.UsageDashboardsAPI]: These APIs manage usage dashboards for this account.
//
// - [iam.UsersAPI]: User identities recognized by Databricks and represented by email addresses.
//
// - [iam.AccountUsersAPI]: User identities recognized by Databricks and represented by email addresses.
//
// - [vectorsearch.VectorSearchEndpointsAPI]: **Endpoint**: Represents the compute resources to host vector search indexes.
//
// - [vectorsearch.VectorSearchIndexesAPI]: **Index**: An efficient representation of your embedding vectors that supports real-time and efficient approximate nearest neighbor (ANN) search queries.
//
// - [catalog.VolumesAPI]: Volumes are a Unity Catalog (UC) capability for accessing, storing, governing, organizing and processing files.
//
// - [provisioning.VpcEndpointsAPI]: These APIs manage VPC endpoint configurations for this account.
//
// - [sql.WarehousesAPI]: A SQL warehouse is a compute resource that lets you run SQL commands on data objects within Databricks SQL.
//
// - [workspace.WorkspaceAPI]: The Workspace API allows you to list, import, export, and delete notebooks and folders.
//
// - [iam.WorkspaceAssignmentAPI]: The Workspace Permission Assignment API allows you to manage workspace permissions for principals in your account.
//
// - [catalog.WorkspaceBindingsAPI]: A securable in Databricks can be configured as __OPEN__ or __ISOLATED__.
//
// - [settings.WorkspaceConfAPI]: This API allows updating known workspace settings for advanced users.
//
// - [provisioning.WorkspacesAPI]: These APIs manage workspaces for this account.
//
// - [billing.BudgetsAPI]: These APIs manage budget configurations for this account.
//
// - [ml.ForecastingAPI]: The Forecasting API allows you to create and get serverless forecasting experiments.
package service

import (
	"github.com/databricks/databricks-sdk-go/service/apps"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/cleanrooms"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/marketplace"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

// adding this trick for godoc to use it as relative import, so that we have
// a clear index of all services in this package at Go package docs:
// https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service
// See: https://pkg.go.dev/golang.org/x/tools/internal/imports#ImportPathToAssumedName
var (
	_ *iam.AccessControlAPI                               = nil
	_ *iam.AccountAccessControlAPI                        = nil
	_ *iam.AccountAccessControlProxyAPI                   = nil
	_ *settings.AibiDashboardEmbeddingAccessPolicyAPI     = nil
	_ *settings.AibiDashboardEmbeddingApprovedDomainsAPI  = nil
	_ *sql.AlertsAPI                                      = nil
	_ *sql.AlertsLegacyAPI                                = nil
	_ *apps.AppsAPI                                       = nil
	_ *catalog.ArtifactAllowlistsAPI                      = nil
	_ *settings.AutomaticClusterUpdateAPI                 = nil
	_ *billing.BillableUsageAPI                           = nil
	_ *billing.BudgetPolicyAPI                            = nil
	_ *catalog.CatalogsAPI                                = nil
	_ *cleanrooms.CleanRoomAssetsAPI                      = nil
	_ *cleanrooms.CleanRoomTaskRunsAPI                    = nil
	_ *cleanrooms.CleanRoomsAPI                           = nil
	_ *compute.ClusterPoliciesAPI                         = nil
	_ *compute.ClustersAPI                                = nil
	_ *compute.CommandExecutionAPI                        = nil
	_ *settings.ComplianceSecurityProfileAPI              = nil
	_ *catalog.ConnectionsAPI                             = nil
	_ *marketplace.ConsumerFulfillmentsAPI                = nil
	_ *marketplace.ConsumerInstallationsAPI               = nil
	_ *marketplace.ConsumerListingsAPI                    = nil
	_ *marketplace.ConsumerPersonalizationRequestsAPI     = nil
	_ *marketplace.ConsumerProvidersAPI                   = nil
	_ *catalog.CredentialsAPI                             = nil
	_ *provisioning.CredentialsAPI                        = nil
	_ *settings.CredentialsManagerAPI                     = nil
	_ *settings.CspEnablementAccountAPI                   = nil
	_ *iam.CurrentUserAPI                                 = nil
	_ *oauth2.CustomAppIntegrationAPI                     = nil
	_ *sql.DashboardWidgetsAPI                            = nil
	_ *sql.DashboardsAPI                                  = nil
	_ *sql.DataSourcesAPI                                 = nil
	_ *files.DbfsAPI                                      = nil
	_ *sql.DbsqlPermissionsAPI                            = nil
	_ *settings.DefaultNamespaceAPI                       = nil
	_ *settings.DisableLegacyAccessAPI                    = nil
	_ *settings.DisableLegacyDbfsAPI                      = nil
	_ *settings.DisableLegacyFeaturesAPI                  = nil
	_ *settings.EnableIpAccessListsAPI                    = nil
	_ *provisioning.EncryptionKeysAPI                     = nil
	_ *settings.EnhancedSecurityMonitoringAPI             = nil
	_ *settings.EsmEnablementAccountAPI                   = nil
	_ *ml.ExperimentsAPI                                  = nil
	_ *catalog.ExternalLocationsAPI                       = nil
	_ *oauth2.AccountFederationPolicyAPI                  = nil
	_ *files.FilesAPI                                     = nil
	_ *catalog.FunctionsAPI                               = nil
	_ *dashboards.GenieAPI                                = nil
	_ *workspace.GitCredentialsAPI                        = nil
	_ *compute.GlobalInitScriptsAPI                       = nil
	_ *catalog.GrantsAPI                                  = nil
	_ *iam.GroupsAPI                                      = nil
	_ *iam.AccountGroupsAPI                               = nil
	_ *compute.InstancePoolsAPI                           = nil
	_ *compute.InstanceProfilesAPI                        = nil
	_ *settings.IpAccessListsAPI                          = nil
	_ *settings.AccountIpAccessListsAPI                   = nil
	_ *jobs.JobsAPI                                       = nil
	_ *dashboards.LakeviewAPI                             = nil
	_ *dashboards.LakeviewEmbeddedAPI                     = nil
	_ *compute.LibrariesAPI                               = nil
	_ *billing.LogDeliveryAPI                             = nil
	_ *catalog.AccountMetastoreAssignmentsAPI             = nil
	_ *catalog.MetastoresAPI                              = nil
	_ *catalog.AccountMetastoresAPI                       = nil
	_ *ml.ModelRegistryAPI                                = nil
	_ *catalog.ModelVersionsAPI                           = nil
	_ *settings.NetworkConnectivityAPI                    = nil
	_ *provisioning.NetworksAPI                           = nil
	_ *settings.NotificationDestinationsAPI               = nil
	_ *oauth2.OAuthPublishedAppsAPI                       = nil
	_ *catalog.OnlineTablesAPI                            = nil
	_ *iam.PermissionMigrationAPI                         = nil
	_ *iam.PermissionsAPI                                 = nil
	_ *settings.PersonalComputeAPI                        = nil
	_ *pipelines.PipelinesAPI                             = nil
	_ *compute.PolicyComplianceForClustersAPI             = nil
	_ *jobs.PolicyComplianceForJobsAPI                    = nil
	_ *compute.PolicyFamiliesAPI                          = nil
	_ *provisioning.PrivateAccessAPI                      = nil
	_ *marketplace.ProviderExchangeFiltersAPI             = nil
	_ *marketplace.ProviderExchangesAPI                   = nil
	_ *marketplace.ProviderFilesAPI                       = nil
	_ *marketplace.ProviderListingsAPI                    = nil
	_ *marketplace.ProviderPersonalizationRequestsAPI     = nil
	_ *marketplace.ProviderProviderAnalyticsDashboardsAPI = nil
	_ *marketplace.ProviderProvidersAPI                   = nil
	_ *sharing.ProvidersAPI                               = nil
	_ *oauth2.PublishedAppIntegrationAPI                  = nil
	_ *catalog.QualityMonitorsAPI                         = nil
	_ *sql.QueriesAPI                                     = nil
	_ *sql.QueriesLegacyAPI                               = nil
	_ *dashboards.QueryExecutionAPI                       = nil
	_ *sql.QueryHistoryAPI                                = nil
	_ *sql.QueryVisualizationsAPI                         = nil
	_ *sql.QueryVisualizationsLegacyAPI                   = nil
	_ *sharing.RecipientActivationAPI                     = nil
	_ *sharing.RecipientsAPI                              = nil
	_ *sql.RedashConfigAPI                                = nil
	_ *catalog.RegisteredModelsAPI                        = nil
	_ *workspace.ReposAPI                                 = nil
	_ *catalog.ResourceQuotasAPI                          = nil
	_ *settings.RestrictWorkspaceAdminsAPI                = nil
	_ *catalog.SchemasAPI                                 = nil
	_ *workspace.SecretsAPI                               = nil
	_ *oauth2.ServicePrincipalFederationPolicyAPI         = nil
	_ *oauth2.ServicePrincipalSecretsAPI                  = nil
	_ *iam.ServicePrincipalsAPI                           = nil
	_ *iam.AccountServicePrincipalsAPI                    = nil
	_ *serving.ServingEndpointsAPI                        = nil
	_ *serving.ServingEndpointsDataPlaneAPI               = nil
	_ *settings.SettingsAPI                               = nil
	_ *settings.AccountSettingsAPI                        = nil
	_ *sharing.SharesAPI                                  = nil
	_ *sql.StatementExecutionAPI                          = nil
	_ *provisioning.StorageAPI                            = nil
	_ *catalog.StorageCredentialsAPI                      = nil
	_ *catalog.AccountStorageCredentialsAPI               = nil
	_ *catalog.SystemSchemasAPI                           = nil
	_ *catalog.TableConstraintsAPI                        = nil
	_ *catalog.TablesAPI                                  = nil
	_ *catalog.TemporaryTableCredentialsAPI               = nil
	_ *settings.TokenManagementAPI                        = nil
	_ *settings.TokensAPI                                 = nil
	_ *billing.UsageDashboardsAPI                         = nil
	_ *iam.UsersAPI                                       = nil
	_ *iam.AccountUsersAPI                                = nil
	_ *vectorsearch.VectorSearchEndpointsAPI              = nil
	_ *vectorsearch.VectorSearchIndexesAPI                = nil
	_ *catalog.VolumesAPI                                 = nil
	_ *provisioning.VpcEndpointsAPI                       = nil
	_ *sql.WarehousesAPI                                  = nil
	_ *workspace.WorkspaceAPI                             = nil
	_ *iam.WorkspaceAssignmentAPI                         = nil
	_ *catalog.WorkspaceBindingsAPI                       = nil
	_ *settings.WorkspaceConfAPI                          = nil
	_ *provisioning.WorkspacesAPI                         = nil
	_ *billing.BudgetsAPI                                 = nil
	_ *ml.ForecastingAPI                                  = nil
)
