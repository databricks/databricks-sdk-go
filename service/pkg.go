// Databricks SDK for Go APIs
//
// - [iam.AccountAccessControlAPI]: These APIs manage access rules on resources in an account.
//
// - [iam.AccountAccessControlProxyAPI]: These APIs manage access rules on resources in an account.
//
// - [sql.AlertsAPI]: The alerts API can be used to perform CRUD operations on alerts.
//
// - [billing.BillableUsageAPI]: This API allows you to download billable usage logs for the specified account and date range.
//
// - [billing.BudgetsAPI]: These APIs manage budget configuration including notifications for exceeding a budget for a period.
//
// - [catalog.CatalogsAPI]: A catalog is the first layer of Unity Catalog’s three-level namespace.
//
// - [compute.ClusterPoliciesAPI]: Cluster policy limits the ability to configure clusters based on a set of rules.
//
// - [compute.ClustersAPI]: The Clusters API allows you to create, start, edit, list, terminate, and delete clusters.
//
// - [compute.CommandExecutionAPI]: This API allows execution of Python, Scala, SQL, or R commands on running Databricks Clusters.
//
// - [catalog.ConnectionsAPI]: Connections allow for creating a connection to an external data source.
//
// - [provisioning.CredentialsAPI]: These APIs manage credential configurations for this workspace.
//
// - [iam.CurrentUserAPI]: This API allows retrieving information about currently authenticated user or service principal.
//
// - [oauth2.CustomAppIntegrationAPI]: These APIs enable administrators to manage custom oauth app integrations, which is required for adding/using Custom OAuth App Integration like Tableau Cloud for Databricks in AWS cloud.
//
// - [sql.DashboardsAPI]: In general, there is little need to modify dashboards using the API.
//
// - [sql.DataSourcesAPI]: This API is provided to assist you in making new query objects.
//
// - [files.DbfsAPI]: DBFS API makes it simple to interact with various data sources without having to include a users credentials every time to read a file.
//
// - [sql.DbsqlPermissionsAPI]: The SQL Permissions API is similar to the endpoints of the :method:permissions/set.
//
// - [provisioning.EncryptionKeysAPI]: These APIs manage encryption key configurations for this workspace (optional).
//
// - [ml.ExperimentsAPI]: Experiments are the primary unit of organization in MLflow; all MLflow runs belong to an experiment.
//
// - [catalog.ExternalLocationsAPI]: An external location is an object that combines a cloud storage path with a storage credential that authorizes access to the cloud storage path.
//
// - [catalog.FunctionsAPI]: Functions implement User-Defined Functions (UDFs) in Unity Catalog.
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
// - [ml.ModelRegistryAPI]: MLflow Model Registry is a centralized model repository and a UI and set of APIs that enable you to manage the full lifecycle of MLflow Models.
//
// - [provisioning.NetworksAPI]: These APIs manage network configurations for customer-managed VPCs (optional).
//
// - [oauth2.OAuthEnrollmentAPI]: These APIs enable administrators to enroll OAuth for their accounts, which is required for adding/using any OAuth published/custom application integration.
//
// - [iam.PermissionsAPI]: Permissions API are used to create read, write, edit, update and manage access for various users on different objects and endpoints.
//
// - [pipelines.PipelinesAPI]: The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.
//
// - [compute.PolicyFamiliesAPI]: View available policy families.
//
// - [provisioning.PrivateAccessAPI]: These APIs manage private access settings for this account.
//
// - [sharing.ProvidersAPI]: Databricks Providers REST API.
//
// - [oauth2.PublishedAppIntegrationAPI]: These APIs enable administrators to manage published oauth app integrations, which is required for adding/using Published OAuth App Integration like Tableau Cloud for Databricks in AWS cloud.
//
// - [sql.QueriesAPI]: These endpoints are used for CRUD operations on query definitions.
//
// - [sql.QueryHistoryAPI]: Access the history of queries through SQL warehouses.
//
// - [sharing.RecipientActivationAPI]: Databricks Recipient Activation REST API.
//
// - [sharing.RecipientsAPI]: Databricks Recipients REST API.
//
// - [workspace.ReposAPI]: The Repos API allows users to manage their git repos.
//
// - [catalog.SchemasAPI]: A schema (also called a database) is the second layer of Unity Catalog’s three-level namespace.
//
// - [workspace.SecretsAPI]: The Secrets API allows you to manage secrets, secret scopes, and access permissions.
//
// - [oauth2.ServicePrincipalSecretsAPI]: These APIs enable administrators to manage service principal secrets.
//
// - [iam.ServicePrincipalsAPI]: Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.
//
// - [iam.AccountServicePrincipalsAPI]: Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.
//
// - [serving.ServingEndpointsAPI]: The Serving Endpoints API allows you to create, update, and delete model serving endpoints.
//
// - [settings.AccountSettingsAPI]: TBD.
//
// - [sharing.SharesAPI]: Databricks Shares REST API.
//
// - [sql.StatementExecutionAPI]: The SQL Statement Execution API manages the execution of arbitrary SQL statements and the fetching of result data.
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
// - [settings.TokenManagementAPI]: Enables administrators to get all tokens and delete tokens for other users.
//
// - [settings.TokensAPI]: The Token API allows you to create, list, and revoke tokens that can be used to authenticate and access Databricks REST APIs.
//
// - [iam.UsersAPI]: User identities recognized by Databricks and represented by email addresses.
//
// - [iam.AccountUsersAPI]: User identities recognized by Databricks and represented by email addresses.
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
// - [catalog.WorkspaceBindingsAPI]: A catalog in Databricks can be configured as __OPEN__ or __ISOLATED__.
//
// - [settings.WorkspaceConfAPI]: This API allows updating known workspace settings for advanced users.
//
// - [provisioning.WorkspacesAPI]: These APIs manage workspaces for this account.
package service

import (
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/oauth2"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

// adding this trick for godoc to use it as relative import, so that we have
// a clear index of all services in this package at Go package docs:
// https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service
// See: https://pkg.go.dev/golang.org/x/tools/internal/imports#ImportPathToAssumedName
var (
	_ *iam.AccountAccessControlAPI            = nil
	_ *iam.AccountAccessControlProxyAPI       = nil
	_ *sql.AlertsAPI                          = nil
	_ *billing.BillableUsageAPI               = nil
	_ *billing.BudgetsAPI                     = nil
	_ *catalog.CatalogsAPI                    = nil
	_ *compute.ClusterPoliciesAPI             = nil
	_ *compute.ClustersAPI                    = nil
	_ *compute.CommandExecutionAPI            = nil
	_ *catalog.ConnectionsAPI                 = nil
	_ *provisioning.CredentialsAPI            = nil
	_ *iam.CurrentUserAPI                     = nil
	_ *oauth2.CustomAppIntegrationAPI         = nil
	_ *sql.DashboardsAPI                      = nil
	_ *sql.DataSourcesAPI                     = nil
	_ *files.DbfsAPI                          = nil
	_ *sql.DbsqlPermissionsAPI                = nil
	_ *provisioning.EncryptionKeysAPI         = nil
	_ *ml.ExperimentsAPI                      = nil
	_ *catalog.ExternalLocationsAPI           = nil
	_ *catalog.FunctionsAPI                   = nil
	_ *workspace.GitCredentialsAPI            = nil
	_ *compute.GlobalInitScriptsAPI           = nil
	_ *catalog.GrantsAPI                      = nil
	_ *iam.GroupsAPI                          = nil
	_ *iam.AccountGroupsAPI                   = nil
	_ *compute.InstancePoolsAPI               = nil
	_ *compute.InstanceProfilesAPI            = nil
	_ *settings.IpAccessListsAPI              = nil
	_ *settings.AccountIpAccessListsAPI       = nil
	_ *jobs.JobsAPI                           = nil
	_ *compute.LibrariesAPI                   = nil
	_ *billing.LogDeliveryAPI                 = nil
	_ *catalog.AccountMetastoreAssignmentsAPI = nil
	_ *catalog.MetastoresAPI                  = nil
	_ *catalog.AccountMetastoresAPI           = nil
	_ *ml.ModelRegistryAPI                    = nil
	_ *provisioning.NetworksAPI               = nil
	_ *oauth2.OAuthEnrollmentAPI              = nil
	_ *iam.PermissionsAPI                     = nil
	_ *pipelines.PipelinesAPI                 = nil
	_ *compute.PolicyFamiliesAPI              = nil
	_ *provisioning.PrivateAccessAPI          = nil
	_ *sharing.ProvidersAPI                   = nil
	_ *oauth2.PublishedAppIntegrationAPI      = nil
	_ *sql.QueriesAPI                         = nil
	_ *sql.QueryHistoryAPI                    = nil
	_ *sharing.RecipientActivationAPI         = nil
	_ *sharing.RecipientsAPI                  = nil
	_ *workspace.ReposAPI                     = nil
	_ *catalog.SchemasAPI                     = nil
	_ *workspace.SecretsAPI                   = nil
	_ *oauth2.ServicePrincipalSecretsAPI      = nil
	_ *iam.ServicePrincipalsAPI               = nil
	_ *iam.AccountServicePrincipalsAPI        = nil
	_ *serving.ServingEndpointsAPI            = nil
	_ *settings.AccountSettingsAPI            = nil
	_ *sharing.SharesAPI                      = nil
	_ *sql.StatementExecutionAPI              = nil
	_ *provisioning.StorageAPI                = nil
	_ *catalog.StorageCredentialsAPI          = nil
	_ *catalog.AccountStorageCredentialsAPI   = nil
	_ *catalog.SystemSchemasAPI               = nil
	_ *catalog.TableConstraintsAPI            = nil
	_ *catalog.TablesAPI                      = nil
	_ *settings.TokenManagementAPI            = nil
	_ *settings.TokensAPI                     = nil
	_ *iam.UsersAPI                           = nil
	_ *iam.AccountUsersAPI                    = nil
	_ *catalog.VolumesAPI                     = nil
	_ *provisioning.VpcEndpointsAPI           = nil
	_ *sql.WarehousesAPI                      = nil
	_ *workspace.WorkspaceAPI                 = nil
	_ *iam.WorkspaceAssignmentAPI             = nil
	_ *catalog.WorkspaceBindingsAPI           = nil
	_ *settings.WorkspaceConfAPI              = nil
	_ *provisioning.WorkspacesAPI             = nil
)
