// Databricks SDK for Go APIs
//
// - [dbsql.AlertsAPI]: The alerts API can be used to perform CRUD operations on alerts.
//
// - [billing.BillableUsageAPI]: This API allows you to download billable usage logs for the specified account and date range.
//
// - [billing.BudgetsAPI]: These APIs manage budget configuration including notifications for exceeding a budget for a period.
//
// - [unitycatalog.CatalogsAPI]: A catalog is the first layer of Unity Catalog’s three-level namespace.
//
// - [clusterpolicies.ClusterPoliciesAPI]: Cluster policy limits the ability to configure clusters based on a set of rules.
//
// - [clusters.ClustersAPI]: The Clusters API allows you to create, start, edit, list, terminate, and delete clusters.
//
// - [commands.CommandExecutionAPI]: This API allows execution of Python, Scala, SQL, or R commands on running Databricks Clusters.
//
// - [deployment.CredentialsAPI]: These APIs manage credential configurations for this workspace.
//
// - [scim.CurrentUserAPI]: This API allows retrieving information about currently authenticated user or service principal.
//
// - [dbsql.DashboardsAPI]: In general, there is little need to modify dashboards using the API.
//
// - [dbsql.DataSourcesAPI]: This API is provided to assist you in making new query objects.
//
// - [dbfs.DbfsAPI]: DBFS API makes it simple to interact with various data sources without having to include a users credentials every time to read a file.
//
// - [dbsql.DbsqlPermissionsAPI]: The SQL Permissions API is similar to the endpoints of the :method:permissions/setobjectpermissions.
//
// - [deployment.EncryptionKeysAPI]: These APIs manage encryption key configurations for this workspace (optional).
//
// - [unitycatalog.ExternalLocationsAPI]: An external location is an object that combines a cloud storage path with a storage credential that authorizes access to the cloud storage path.
//
// - [gitcredentials.GitCredentialsAPI]: Registers personal access token for Databricks to do operations on behalf of the user.
//
// - [globalinitscripts.GlobalInitScriptsAPI]: The Global Init Scripts API enables Workspace administrators to configure global initialization scripts for their workspace.
//
// - [unitycatalog.GrantsAPI]: In Unity Catalog, data is secure by default.
//
// - [scim.GroupsAPI]: Groups simplify identity management, making it easier to assign access to Databricks Workspace, data, and other securable objects.
//
// - [scim.AccountGroupsAPI]: Groups simplify identity management, making it easier to assign access to Databricks Account, data, and other securable objects.
//
// - [instancepools.InstancePoolsAPI]: Instance Pools API are used to create, edit, delete and list instance pools by using ready-to-use cloud instances which reduces a cluster start and auto-scaling times.
//
// - [clusters.InstanceProfilesAPI]: The Instance Profiles API allows admins to add, list, and remove instance profiles that users can launch clusters with.
//
// - [ipaccesslists.IpAccessListsAPI]: The IP Access List API enables Databricks admins to configure IP access lists for a workspace.
//
// - [jobs.JobsAPI]: The Jobs API allows you to create, edit, and delete jobs.
//
// - [libraries.LibrariesAPI]: The Libraries API allows you to install and uninstall libraries and get the status of libraries on a cluster.
//
// - [billing.LogDeliveryAPI]: These APIs manage log delivery configurations for this account.
//
// - [mlflow.MLflowDatabricksAPI]: These endpoints are modified versions of the MLflow API that accept additional input parameters or return additional information.
//
// - [unitycatalog.MetastoresAPI]: A metastore is the top-level container of objects in Unity Catalog.
//
// - [deployment.NetworksAPI]: These APIs manage network configurations for customer-managed VPCs (optional).
//
// - [permissions.PermissionsAPI]: Permissions API are used to create read, write, edit, update and manage access for various users on different objects and endpoints.
//
// - [pipelines.PipelinesAPI]: The Delta Live Tables API allows you to create, edit, delete, start, and view details about pipelines.
//
// - [deployment.PrivateAccessAPI]: These APIs manage private access settings for this account.
//
// - [unitycatalog.ProvidersAPI]: Databricks Delta Sharing: Providers REST API.
//
// - [dbsql.QueriesAPI]: These endpoints are used for CRUD operations on query definitions.
//
// - [warehouses.QueryHistoryAPI]: Access the history of queries through SQL warehouses.
//
// - [unitycatalog.RecipientActivationAPI]: Databricks Delta Sharing: Recipient Activation REST API.
//
// - [unitycatalog.RecipientsAPI]: Databricks Delta Sharing: Recipients REST API.
//
// - [repos.ReposAPI]: The Repos API allows users to manage their git repos.
//
// - [unitycatalog.SchemasAPI]: A schema (also called a database) is the second layer of Unity Catalog’s three-level namespace.
//
// - [secrets.SecretsAPI]: The Secrets API allows you to manage secrets, secret scopes, and access permissions.
//
// - [scim.ServicePrincipalsAPI]: Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.
//
// - [scim.AccountServicePrincipalsAPI]: Identities for use with jobs, automated tools, and systems such as scripts, apps, and CI/CD platforms.
//
// - [unitycatalog.SharesAPI]: Databricks Delta Sharing: Shares REST API.
//
// - [deployment.StorageAPI]: These APIs manage storage configurations for this workspace.
//
// - [unitycatalog.StorageCredentialsAPI]: A storage credential represents an authentication and authorization mechanism for accessing data stored on your cloud tenant, using an IAM role.
//
// - [unitycatalog.TablesAPI]: A table resides in the third layer of Unity Catalog’s three-level namespace.
//
// - [tokenmanagement.TokenManagementAPI]: Enables administrators to get all tokens and delete tokens for other users.
//
// - [tokens.TokensAPI]: The Token API allows you to create, list, and revoke tokens that can be used to authenticate and access Databricks REST APIs.
//
// - [scim.UsersAPI]: User identities recognized by Databricks and represented by email addresses.
//
// - [scim.AccountUsersAPI]: User identities recognized by Databricks and represented by email addresses.
//
// - [deployment.VpcEndpointsAPI]: These APIs manage VPC endpoint configurations for this account.
//
// - [warehouses.WarehousesAPI]: A SQL warehouse is a compute resource that lets you run SQL commands on data objects within Databricks SQL.
//
// - [workspace.WorkspaceAPI]: The Workspace API allows you to list, import, export, and delete notebooks and folders.
//
// - [permissions.WorkspaceAssignmentAPI]: Databricks Workspace Assignment REST API.
//
// - [workspaceconf.WorkspaceConfAPI]: This API allows updating known workspace settings for advanced users.
//
// - [deployment.WorkspacesAPI]: These APIs manage workspaces for this account.
package service

import (
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/clusterpolicies"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/service/dbsql"
	"github.com/databricks/databricks-sdk-go/service/deployment"
	"github.com/databricks/databricks-sdk-go/service/gitcredentials"
	"github.com/databricks/databricks-sdk-go/service/globalinitscripts"
	"github.com/databricks/databricks-sdk-go/service/instancepools"
	"github.com/databricks/databricks-sdk-go/service/ipaccesslists"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/libraries"
	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/repos"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/databricks/databricks-sdk-go/service/secrets"
	"github.com/databricks/databricks-sdk-go/service/tokenmanagement"
	"github.com/databricks/databricks-sdk-go/service/tokens"
	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/databricks-sdk-go/service/warehouses"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/databricks-sdk-go/service/workspaceconf"
)

// adding this trick for godoc to use it as relative import, so that we have
// a clear index of all services in this package at Go package docs:
// https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service
// See: https://pkg.go.dev/golang.org/x/tools/internal/imports#ImportPathToAssumedName
var (
	_ *dbsql.AlertsAPI                        = nil
	_ *billing.BillableUsageAPI               = nil
	_ *billing.BudgetsAPI                     = nil
	_ *unitycatalog.CatalogsAPI               = nil
	_ *clusterpolicies.ClusterPoliciesAPI     = nil
	_ *clusters.ClustersAPI                   = nil
	_ *commands.CommandExecutionAPI           = nil
	_ *deployment.CredentialsAPI              = nil
	_ *scim.CurrentUserAPI                    = nil
	_ *dbsql.DashboardsAPI                    = nil
	_ *dbsql.DataSourcesAPI                   = nil
	_ *dbfs.DbfsAPI                           = nil
	_ *dbsql.DbsqlPermissionsAPI              = nil
	_ *deployment.EncryptionKeysAPI           = nil
	_ *mlflow.ExperimentsAPI                  = nil
	_ *unitycatalog.ExternalLocationsAPI      = nil
	_ *gitcredentials.GitCredentialsAPI       = nil
	_ *globalinitscripts.GlobalInitScriptsAPI = nil
	_ *unitycatalog.GrantsAPI                 = nil
	_ *scim.GroupsAPI                         = nil
	_ *scim.AccountGroupsAPI                  = nil
	_ *instancepools.InstancePoolsAPI         = nil
	_ *clusters.InstanceProfilesAPI           = nil
	_ *ipaccesslists.IpAccessListsAPI         = nil
	_ *jobs.JobsAPI                           = nil
	_ *libraries.LibrariesAPI                 = nil
	_ *billing.LogDeliveryAPI                 = nil
	_ *mlflow.MLflowArtifactsAPI              = nil
	_ *mlflow.MLflowDatabricksAPI             = nil
	_ *mlflow.MLflowMetricsAPI                = nil
	_ *mlflow.MLflowRunsAPI                   = nil
	_ *unitycatalog.MetastoresAPI             = nil
	_ *mlflow.ModelVersionCommentsAPI         = nil
	_ *mlflow.ModelVersionsAPI                = nil
	_ *deployment.NetworksAPI                 = nil
	_ *permissions.PermissionsAPI             = nil
	_ *pipelines.PipelinesAPI                 = nil
	_ *deployment.PrivateAccessAPI            = nil
	_ *unitycatalog.ProvidersAPI              = nil
	_ *dbsql.QueriesAPI                       = nil
	_ *warehouses.QueryHistoryAPI             = nil
	_ *unitycatalog.RecipientActivationAPI    = nil
	_ *unitycatalog.RecipientsAPI             = nil
	_ *mlflow.RegisteredModelsAPI             = nil
	_ *mlflow.RegistryWebhooksAPI             = nil
	_ *repos.ReposAPI                         = nil
	_ *unitycatalog.SchemasAPI                = nil
	_ *secrets.SecretsAPI                     = nil
	_ *scim.ServicePrincipalsAPI              = nil
	_ *scim.AccountServicePrincipalsAPI       = nil
	_ *unitycatalog.SharesAPI                 = nil
	_ *deployment.StorageAPI                  = nil
	_ *unitycatalog.StorageCredentialsAPI     = nil
	_ *unitycatalog.TablesAPI                 = nil
	_ *tokenmanagement.TokenManagementAPI     = nil
	_ *tokens.TokensAPI                       = nil
	_ *mlflow.TransitionRequestsAPI           = nil
	_ *scim.UsersAPI                          = nil
	_ *scim.AccountUsersAPI                   = nil
	_ *deployment.VpcEndpointsAPI             = nil
	_ *warehouses.WarehousesAPI               = nil
	_ *workspace.WorkspaceAPI                 = nil
	_ *permissions.WorkspaceAssignmentAPI     = nil
	_ *workspaceconf.WorkspaceConfAPI         = nil
	_ *deployment.WorkspacesAPI               = nil
)
