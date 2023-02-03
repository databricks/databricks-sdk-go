// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databricks

import (
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"

	"github.com/databricks/databricks-sdk-go/service/clusterpolicies"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/service/endpoints"
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
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/tokenmanagement"
	"github.com/databricks/databricks-sdk-go/service/tokens"
	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/databricks-sdk-go/service/workspaceconf"
)

type WorkspaceClient struct {
	Config *config.Config

	// The alerts API can be used to perform CRUD operations on alerts. An alert
	// is a Databricks SQL object that periodically runs a query, evaluates a
	// condition of its result, and notifies one or more users and/or alert
	// destinations if the condition was met.
	//
	// **Note**: Programmatic operations on refresh schedules via the Databricks
	// SQL API are deprecated. Alert refresh schedules can be created, updated,
	// fetched and deleted using Jobs API, e.g. :method:jobs/create.
	Alerts *sql.AlertsAPI

	// A catalog is the first layer of Unity Catalog’s three-level namespace.
	// It’s used to organize your data assets. Users can see all catalogs on
	// which they have been assigned the USE_CATALOG data permission.
	//
	// In Unity Catalog, admins and data stewards manage users and their access
	// to data centrally across all of the workspaces in a Databricks account.
	// Users in different workspaces can share access to the same data,
	// depending on privileges granted centrally in Unity Catalog.
	Catalogs *unitycatalog.CatalogsAPI

	// Cluster policy limits the ability to configure clusters based on a set of
	// rules. The policy rules limit the attributes or attribute values
	// available for cluster creation. Cluster policies have ACLs that limit
	// their use to specific users and groups.
	//
	// Cluster policies let you limit users to create clusters with prescribed
	// settings, simplify the user interface and enable more users to create
	// their own clusters (by fixing and hiding some values), control cost by
	// limiting per cluster maximum cost (by setting limits on attributes whose
	// values contribute to hourly price).
	//
	// Cluster policy permissions limit which policies a user can select in the
	// Policy drop-down when the user creates a cluster: - A user who has
	// cluster create permission can select the Unrestricted policy and create
	// fully-configurable clusters. - A user who has both cluster create
	// permission and access to cluster policies can select the Unrestricted
	// policy and policies they have access to. - A user that has access to only
	// cluster policies, can select the policies they have access to.
	//
	// If no policies have been created in the workspace, the Policy drop-down
	// does not display.
	//
	// Only admin users can create, edit, and delete policies. Admin users also
	// have access to all policies.
	ClusterPolicies *clusterpolicies.ClusterPoliciesAPI

	// The Clusters API allows you to create, start, edit, list, terminate, and
	// delete clusters.
	//
	// Databricks maps cluster node instance types to compute units known as
	// DBUs. See the instance type pricing page for a list of the supported
	// instance types and their corresponding DBUs.
	//
	// A Databricks cluster is a set of computation resources and configurations
	// on which you run data engineering, data science, and data analytics
	// workloads, such as production ETL pipelines, streaming analytics, ad-hoc
	// analytics, and machine learning.
	//
	// You run these workloads as a set of commands in a notebook or as an
	// automated job. Databricks makes a distinction between all-purpose
	// clusters and job clusters. You use all-purpose clusters to analyze data
	// collaboratively using interactive notebooks. You use job clusters to run
	// fast and robust automated jobs.
	//
	// You can create an all-purpose cluster using the UI, CLI, or REST API. You
	// can manually terminate and restart an all-purpose cluster. Multiple users
	// can share such clusters to do collaborative interactive analysis.
	//
	// IMPORTANT: Databricks retains cluster configuration information for up to
	// 200 all-purpose clusters terminated in the last 30 days and up to 30 job
	// clusters recently terminated by the job scheduler. To keep an all-purpose
	// cluster configuration even after it has been terminated for more than 30
	// days, an administrator can pin a cluster to the cluster list.
	Clusters *clusters.ClustersAPI

	// This API allows executing commands on running clusters.
	CommandExecutor commands.CommandExecutor

	// This API allows retrieving information about currently authenticated user
	// or service principal.
	CurrentUser *scim.CurrentUserAPI

	// In general, there is little need to modify dashboards using the API.
	// However, it can be useful to use dashboard objects to look-up a
	// collection of related query IDs. The API can also be used to duplicate
	// multiple dashboards at once since you can get a dashboard definition with
	// a GET request and then POST it to create a new one.
	//
	// **Note**: Programmatic operations on refresh schedules via the Databricks
	// SQL API are deprecated. Dashboard refresh schedules can be created,
	// updated, fetched and deleted using Jobs API, e.g. :method:jobs/create.
	Dashboards *sql.DashboardsAPI

	// This API is provided to assist you in making new query objects. When
	// creating a query object, you may optionally specify a `data_source_id`
	// for the SQL warehouse against which it will run. If you don't already
	// know the `data_source_id` for your desired SQL warehouse, this API will
	// help you find it.
	//
	// This API does not support searches. It returns the full list of SQL
	// warehouses in your workspace. We advise you to use any text editor, REST
	// client, or `grep` to search the response from this API for the name of
	// your SQL warehouse as it appears in Databricks SQL.
	DataSources *sql.DataSourcesAPI

	// DBFS API makes it simple to interact with various data sources without
	// having to include a users credentials every time to read a file.
	Dbfs *dbfs.DbfsAPI

	// The SQL Permissions API is similar to the endpoints of the
	// :method:permissions/set. However, this exposes only one endpoint, which
	// gets the Access Control List for a given object. You cannot modify any
	// permissions using this API.
	//
	// There are three levels of permission:
	//
	// - `CAN_VIEW`: Allows read-only access
	//
	// - `CAN_RUN`: Allows read access and run access (superset of `CAN_VIEW`)
	//
	// - `CAN_MANAGE`: Allows all actions: read, run, edit, delete, modify
	// permissions (superset of `CAN_RUN`)
	DbsqlPermissions *sql.DbsqlPermissionsAPI

	Experiments *mlflow.ExperimentsAPI

	// An external location is an object that combines a cloud storage path with
	// a storage credential that authorizes access to the cloud storage path.
	// Each external location is subject to Unity Catalog access-control
	// policies that control which users and groups can access the credential.
	// If a user does not have access to an external location in Unity Catalog,
	// the request fails and Unity Catalog does not attempt to authenticate to
	// your cloud tenant on the user’s behalf.
	//
	// Databricks recommends using external locations rather than using storage
	// credentials directly.
	//
	// To create external locations, you must be a metastore admin or a user
	// with the **CREATE_EXTERNAL_LOCATION** privilege.
	ExternalLocations *unitycatalog.ExternalLocationsAPI

	// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
	//
	// The function implementation can be any SQL expression or Query, and it
	// can be invoked wherever a table reference is allowed in a query. In Unity
	// Catalog, a function resides at the same level as a table, so it can be
	// referenced with the form
	// __catalog_name__.__schema_name__.__function_name__.
	Functions *unitycatalog.FunctionsAPI

	// Registers personal access token for Databricks to do operations on behalf
	// of the user.
	//
	// See [more info].
	//
	// [more info]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	GitCredentials *gitcredentials.GitCredentialsAPI

	// The Global Init Scripts API enables Workspace administrators to configure
	// global initialization scripts for their workspace. These scripts run on
	// every node in every cluster in the workspace.
	//
	// **Important:** Existing clusters must be restarted to pick up any changes
	// made to global init scripts. Global init scripts are run in order. If the
	// init script returns with a bad exit code, the Apache Spark container
	// fails to launch and init scripts with later position are skipped. If
	// enough containers fail, the entire cluster fails with a
	// `GLOBAL_INIT_SCRIPT_FAILURE` error code.
	GlobalInitScripts *globalinitscripts.GlobalInitScriptsAPI

	// In Unity Catalog, data is secure by default. Initially, users have no
	// access to data in a metastore. Access can be granted by either a
	// metastore admin, the owner of an object, or the owner of the catalog or
	// schema that contains the object. Securable objects in Unity Catalog are
	// hierarchical and privileges are inherited downward.
	//
	// Initially, users have no access to data in a metastore. Access can be
	// granted by either a metastore admin, the owner of an object, or the owner
	// of the catalog or schema that contains the object.
	//
	// Securable objects in Unity Catalog are hierarchical and privileges are
	// inherited downward. This means that granting a privilege on the catalog
	// automatically grants the privilege to all current and future objects
	// within the catalog. Similarly, privileges granted on a schema are
	// inherited by all current and future objects within that schema.
	Grants *unitycatalog.GrantsAPI

	// Groups simplify identity management, making it easier to assign access to
	// Databricks Workspace, data, and other securable objects.
	//
	// It is best practice to assign access to workspaces and access-control
	// policies in Unity Catalog to groups, instead of to users individually.
	// All Databricks Workspace identities can be assigned as members of groups,
	// and members inherit permissions that are assigned to their group.
	Groups *scim.GroupsAPI

	// Instance Pools API are used to create, edit, delete and list instance
	// pools by using ready-to-use cloud instances which reduces a cluster start
	// and auto-scaling times.
	//
	// Databricks pools reduce cluster start and auto-scaling times by
	// maintaining a set of idle, ready-to-use instances. When a cluster is
	// attached to a pool, cluster nodes are created using the pool’s idle
	// instances. If the pool has no idle instances, the pool expands by
	// allocating a new instance from the instance provider in order to
	// accommodate the cluster’s request. When a cluster releases an instance,
	// it returns to the pool and is free for another cluster to use. Only
	// clusters attached to a pool can use that pool’s idle instances.
	//
	// You can specify a different pool for the driver node and worker nodes, or
	// use the same pool for both.
	//
	// Databricks does not charge DBUs while instances are idle in the pool.
	// Instance provider billing does apply. See pricing.
	InstancePools *instancepools.InstancePoolsAPI

	// The Instance Profiles API allows admins to add, list, and remove instance
	// profiles that users can launch clusters with. Regular users can list the
	// instance profiles available to them. See [Secure access to S3 buckets]
	// using instance profiles for more information.
	//
	// [Secure access to S3 buckets]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html
	InstanceProfiles *clusters.InstanceProfilesAPI

	// IP Access List enables admins to configure IP access lists.
	//
	// IP access lists affect web application access and REST API access to this
	// workspace only. If the feature is disabled for a workspace, all access is
	// allowed for this workspace. There is support for allow lists (inclusion)
	// and block lists (exclusion).
	//
	// When a connection is attempted: 1. **First, all block lists are
	// checked.** If the connection IP address matches any block list, the
	// connection is rejected. 2. **If the connection was not rejected by block
	// lists**, the IP address is compared with the allow lists.
	//
	// If there is at least one allow list for the workspace, the connection is
	// allowed only if the IP address matches an allow list. If there are no
	// allow lists for the workspace, all IP addresses are allowed.
	//
	// For all allow lists and block lists combined, the workspace supports a
	// maximum of 1000 IP/CIDR values, where one CIDR counts as a single value.
	//
	// After changes to the IP access list feature, it can take a few minutes
	// for changes to take effect.
	IpAccessLists *ipaccesslists.IpAccessListsAPI

	// The Jobs API allows you to create, edit, and delete jobs.
	//
	// You can use a Databricks job to run a data processing or data analysis
	// task in a Databricks cluster with scalable resources. Your job can
	// consist of a single task or can be a large, multi-task workflow with
	// complex dependencies. Databricks manages the task orchestration, cluster
	// management, monitoring, and error reporting for all of your jobs. You can
	// run your jobs immediately or periodically through an easy-to-use
	// scheduling system. You can implement job tasks using notebooks, JARS,
	// Delta Live Tables pipelines, or Python, Scala, Spark submit, and Java
	// applications.
	//
	// You should never hard code secrets or store them in plain text. Use the
	// :service:secrets to manage secrets in the [Databricks CLI]. Use the
	// [Secrets utility] to reference secrets in notebooks and jobs.
	//
	// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
	// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
	Jobs *jobs.JobsAPI

	// The Libraries API allows you to install and uninstall libraries and get
	// the status of libraries on a cluster.
	//
	// To make third-party or custom code available to notebooks and jobs
	// running on your clusters, you can install a library. Libraries can be
	// written in Python, Java, Scala, and R. You can upload Java, Scala, and
	// Python libraries and point to external packages in PyPI, Maven, and CRAN
	// repositories.
	//
	// Cluster libraries can be used by all notebooks running on a cluster. You
	// can install a cluster library directly from a public repository such as
	// PyPI or Maven, using a previously installed workspace library, or using
	// an init script.
	//
	// When you install a library on a cluster, a notebook already attached to
	// that cluster will not immediately see the new library. You must first
	// detach and then reattach the notebook to the cluster.
	//
	// When you uninstall a library from a cluster, the library is removed only
	// when you restart the cluster. Until you restart the cluster, the status
	// of the uninstalled library appears as Uninstall pending restart.
	Libraries *libraries.LibrariesAPI

	MLflowArtifacts *mlflow.MLflowArtifactsAPI

	// These endpoints are modified versions of the MLflow API that accept
	// additional input parameters or return additional information.
	MLflowDatabricks *mlflow.MLflowDatabricksAPI

	MLflowMetrics *mlflow.MLflowMetricsAPI

	MLflowRuns *mlflow.MLflowRunsAPI

	// A metastore is the top-level container of objects in Unity Catalog. It
	// stores data assets (tables and views) and the permissions that govern
	// access to them. Databricks account admins can create metastores and
	// assign them to Databricks workspaces to control which workloads use each
	// metastore. For a workspace to use Unity Catalog, it must have a Unity
	// Catalog metastore attached.
	//
	// Each metastore is configured with a root storage location in a cloud
	// storage account. This storage location is used for metadata and managed
	// tables data.
	//
	// NOTE: This metastore is distinct from the metastore included in
	// Databricks workspaces created before Unity Catalog was released. If your
	// workspace includes a legacy Hive metastore, the data in that metastore is
	// available in Unity Catalog in a catalog named hive_metastore.
	Metastores *unitycatalog.MetastoresAPI

	ModelVersionComments *mlflow.ModelVersionCommentsAPI

	ModelVersions *mlflow.ModelVersionsAPI

	// Permissions API are used to create read, write, edit, update and manage
	// access for various users on different objects and endpoints.
	Permissions *permissions.PermissionsAPI

	// The Delta Live Tables API allows you to create, edit, delete, start, and
	// view details about pipelines.
	//
	// Delta Live Tables is a framework for building reliable, maintainable, and
	// testable data processing pipelines. You define the transformations to
	// perform on your data, and Delta Live Tables manages task orchestration,
	// cluster management, monitoring, data quality, and error handling.
	//
	// Instead of defining your data pipelines using a series of separate Apache
	// Spark tasks, Delta Live Tables manages how your data is transformed based
	// on a target schema you define for each processing step. You can also
	// enforce data quality with Delta Live Tables expectations. Expectations
	// allow you to define expected data quality and specify how to handle
	// records that fail those expectations.
	Pipelines *pipelines.PipelinesAPI

	// View available policy families. A policy family contains a policy
	// definition providing best practices for configuring clusters for a
	// particular use case.
	//
	// Databricks manages and provides policy families for several common
	// cluster use cases. You cannot create, edit, or delete policy families.
	//
	// Policy families cannot be used directly to create clusters. Instead, you
	// create cluster policies using a policy family. Cluster policies created
	// using a policy family inherit the policy family's policy definition.
	PolicyFamilies *clusterpolicies.PolicyFamiliesAPI

	// Databricks Delta Sharing: Providers REST API
	Providers *unitycatalog.ProvidersAPI

	// These endpoints are used for CRUD operations on query definitions. Query
	// definitions include the target SQL warehouse, query text, name,
	// description, tags, execution schedule, parameters, and visualizations.
	//
	// **Note**: Programmatic operations on refresh schedules via the Databricks
	// SQL API are deprecated. Query refresh schedules can be created, updated,
	// fetched and deleted using Jobs API, e.g. :method:jobs/create.
	Queries *sql.QueriesAPI

	// Access the history of queries through SQL warehouses.
	QueryHistory *sql.QueryHistoryAPI

	// Databricks Delta Sharing: Recipient Activation REST API
	RecipientActivation *unitycatalog.RecipientActivationAPI

	// Databricks Delta Sharing: Recipients REST API
	Recipients *unitycatalog.RecipientsAPI

	RegisteredModels *mlflow.RegisteredModelsAPI

	RegistryWebhooks *mlflow.RegistryWebhooksAPI

	// The Repos API allows users to manage their git repos. Users can use the
	// API to access all repos that they have manage permissions on.
	//
	// Databricks Repos is a visual Git client in Databricks. It supports common
	// Git operations such a cloning a repository, committing and pushing,
	// pulling, branch management, and visual comparison of diffs when
	// committing.
	//
	// Within Repos you can develop code in notebooks or other files and follow
	// data science and engineering code development best practices using Git
	// for version control, collaboration, and CI/CD.
	Repos *repos.ReposAPI

	// A schema (also called a database) is the second layer of Unity
	// Catalog’s three-level namespace. A schema organizes tables and views.
	// To access (or list) a table or view in a schema, users must have the
	// USE_SCHEMA data permission on the schema and its parent catalog, and they
	// must have the SELECT permission on the table or view. There is no
	// guarantee of a specific ordering of the elements in the array.
	Schemas *unitycatalog.SchemasAPI

	// The Secrets API allows you to manage secrets, secret scopes, and access
	// permissions.
	//
	// Sometimes accessing data requires that you authenticate to external data
	// sources through JDBC. Instead of directly entering your credentials into
	// a notebook, use Databricks secrets to store your credentials and
	// reference them in notebooks and jobs.
	//
	// Administrators, secret creators, and users granted permission can read
	// Databricks secrets. While Databricks makes an effort to redact secret
	// values that might be displayed in notebooks, it is not possible to
	// prevent such users from reading secrets.
	Secrets *secrets.SecretsAPI

	// Identities for use with jobs, automated tools, and systems such as
	// scripts, apps, and CI/CD platforms. Databricks recommends creating
	// service principals to run production jobs or modify production data. If
	// all processes that act on production data run with service principals,
	// interactive users do not need any write, delete, or modify privileges in
	// production. This eliminates the risk of a user overwriting production
	// data by accident.
	ServicePrincipals *scim.ServicePrincipalsAPI

	// The Serverless Real-Time Inference Serving Endpoints API allows you to
	// create, update, and delete model serving endpoints.
	//
	// You can use a serving endpoint to serve models from the Databricks Model
	// Registry. Endpoints expose the underlying models as scalable REST API
	// endpoints using serverless compute. This means the endpoints and
	// associated compute resources are fully managed by Databricks and will not
	// appear in your cloud account. A serving endpoint can consist of one or
	// more MLflow models from the Databricks Model Registry, called served
	// models. A serving endpoint can have at most ten served models. You can
	// configure traffic settings to define how requests should be routed to
	// your served models behind an endpoint. Additionally, you can configure
	// the scale of resources that should be applied to each served model.
	ServingEndpoints *endpoints.ServingEndpointsAPI

	// Databricks Delta Sharing: Shares REST API
	Shares *unitycatalog.SharesAPI

	// The SQL Statement Execution API manages the execution of arbitrary SQL
	// statements and the fetching of result data.
	//
	// # Release Status
	//
	// This feature is in [Private Preview]. To try it, reach out to your
	// Databricks contact.
	//
	// # Getting started
	//
	// We suggest beginning with the [SQL Statement Execution API tutorial].
	//
	// # Overview of statement execution and result fetching
	//
	// Statement execution begins by calling
	// :method:StatementExecution/executeStatement with a valid SQL statement
	// and warehouse ID, along with optional parameters such as the data catalog
	// and output format.
	//
	// When submitting the statement, the call can behave synchronously or
	// asynchronously, based on the `wait_timeout` setting. When set between
	// 5-50 seconds (default: 10) the call behaves synchronously; when set to
	// `0s`, the call is asynchronous and responds immediately if accepted.
	//
	// ----
	//
	// ### **Warning: drop authorization header when fetching data through
	// external links**
	//
	// External link URLs do not require an Authorization header or token, and
	// thus all calls to fetch external links must remove the Authorization
	// header.
	//
	// ----
	//
	// Similar to INLINE mode, callers can iterate through the result set, by
	// using the field `next_chunk_internal_link`. Each internal link response
	// will contain an external link to the raw chunk data, and additionally
	// contain the next_chunk_internal_link if there are more chunks.
	//
	// Unlike INLINE mode, when using EXTERNAL_LINKS, chunks may be fetched out
	// of order, and in parallel to achieve higher throughput.
	//
	// # Limits and limitations
	//
	// - All byte limits are calculated based on internal storage metrics, and
	// will not match byte counts of actual payloads. - INLINE mode statements
	// limited to 16 MiB, and will abort when this limit is exceeded. -
	// Cancelation may silently fail: A successful response from a cancel
	// request indicates that the cancel request was successfully received and
	// sent to the processing engine. However, for example, an outstanding
	// statement may complete execution during signal delivery, with cancel
	// signal arriving too late to be meaningful. Polling for status until a
	// terminal state is reached a reliable way to see final state. - Wait
	// timeouts are approximate, occur server-side, and cannot account for
	// caller delays, network latency from caller to service, and similarly. -
	// After a statement has been submitted and a statement_id produced, that
	// statement's status and result will automatically close after either of 2
	// conditions: - The last result chunk is fetched (or resolved to an
	// external link). - Ten (10) minutes pass with no calls to get status or
	// fetch result data. Best practice: in asynchronous clients, poll for
	// status regularly (and with backoff) to keep the statement open and alive.
	//
	// # Private Preview limitations
	//
	// - `EXTERNAL_LINKS` mode will fail for result sets < 5MB. - After any
	// cancel or close operation, the statement will no longer be visible from
	// the API, specifically - After fetching last result chunk (including
	// `chunk_index=0`), the statement is closed; a short time after closure,
	// the statement will no longer be visible to the API, and further calls may
	// return 404. Thus calling GET .../{statement_id} will return a 404 NOT
	// FOUND error. - In practice, this means that a CANCEL and subsequent poll
	// will often return a NOT FOUND. This will be addressed in a future update.
	//
	// [Private Preview]: https://docs.databricks.com/release-notes/release-types.html
	// [SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
	StatementExecution *sql.StatementExecutionAPI

	// A storage credential represents an authentication and authorization
	// mechanism for accessing data stored on your cloud tenant, using an IAM
	// role. Each storage credential is subject to Unity Catalog access-control
	// policies that control which users and groups can access the credential.
	// If a user does not have access to a storage credential in Unity Catalog,
	// the request fails and Unity Catalog does not attempt to authenticate to
	// your cloud tenant on the user’s behalf.
	//
	// Databricks recommends using external locations rather than using storage
	// credentials directly.
	//
	// To create storage credentials, you must be a Databricks account admin.
	// The account admin who creates the storage credential can delegate
	// ownership to another user or group to manage permissions on it.
	StorageCredentials *unitycatalog.StorageCredentialsAPI

	// Primary key and foreign key constraints encode relationships between
	// fields in tables.
	//
	// Primary and foreign keys are informational only and are not enforced.
	// Foreign keys must reference a primary key in another table. This primary
	// key is the parent constraint of the foreign key and the table this
	// primary key is on is the parent table of the foreign key. Similarly, the
	// foreign key is the child constraint of its referenced primary key; the
	// table of the foreign key is the child table of the primary key.
	//
	// You can declare primary keys and foreign keys as part of the table
	// specification during table creation. You can also add or drop constraints
	// on existing tables.
	TableConstraints *unitycatalog.TableConstraintsAPI

	// A table resides in the third layer of Unity Catalog’s three-level
	// namespace. It contains rows of data. To create a table, users must have
	// CREATE_TABLE and USE_SCHEMA permissions on the schema, and they must have
	// the USE_CATALOG permission on its parent catalog. To query a table, users
	// must have the SELECT permission on the table, and they must have the
	// USE_CATALOG permission on its parent catalog and the USE_SCHEMA
	// permission on its parent schema.
	//
	// A table can be managed or external.
	Tables *unitycatalog.TablesAPI

	// Enables administrators to get all tokens and delete tokens for other
	// users. Admins can either get every token, get a specific token by ID, or
	// get all tokens for a particular user.
	TokenManagement *tokenmanagement.TokenManagementAPI

	// The Token API allows you to create, list, and revoke tokens that can be
	// used to authenticate and access Databricks REST APIs.
	Tokens *tokens.TokensAPI

	TransitionRequests *mlflow.TransitionRequestsAPI

	// User identities recognized by Databricks and represented by email
	// addresses.
	//
	// Databricks recommends using SCIM provisioning to sync users and groups
	// automatically from your identity provider to your Databricks Workspace.
	// SCIM streamlines onboarding a new employee or team by using your identity
	// provider to create users and groups in Databricks Workspace and give them
	// the proper level of access. When a user leaves your organization or no
	// longer needs access to Databricks Workspace, admins can terminate the
	// user in your identity provider and that user’s account will also be
	// removed from Databricks Workspace. This ensures a consistent offboarding
	// process and prevents unauthorized users from accessing sensitive data.
	Users *scim.UsersAPI

	// A SQL warehouse is a compute resource that lets you run SQL commands on
	// data objects within Databricks SQL. Compute resources are infrastructure
	// resources that provide processing capabilities in the cloud.
	Warehouses *sql.WarehousesAPI

	// The Workspace API allows you to list, import, export, and delete
	// notebooks and folders.
	//
	// A notebook is a web-based interface to a document that contains runnable
	// code, visualizations, and explanatory text.
	Workspace *workspace.WorkspaceAPI

	// This API allows updating known workspace settings for advanced users.
	WorkspaceConf *workspaceconf.WorkspaceConfAPI
}

// NewWorkspaceClient creates new Databricks SDK client for Workspaces or
// returns error in case configuration is wrong
func NewWorkspaceClient(c ...*Config) (*WorkspaceClient, error) {
	var cfg *config.Config
	if len(c) == 1 {
		// first config
		cfg = (*config.Config)(c[0])
	} else {
		// default config
		cfg = &config.Config{}
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	return &WorkspaceClient{
		Config:               cfg,
		Alerts:               sql.NewAlerts(apiClient),
		Catalogs:             unitycatalog.NewCatalogs(apiClient),
		ClusterPolicies:      clusterpolicies.NewClusterPolicies(apiClient),
		Clusters:             clusters.NewClusters(apiClient),
		CommandExecutor:      commands.NewCommandExecutor(apiClient),
		CurrentUser:          scim.NewCurrentUser(apiClient),
		Dashboards:           sql.NewDashboards(apiClient),
		DataSources:          sql.NewDataSources(apiClient),
		Dbfs:                 dbfs.NewDbfs(apiClient),
		DbsqlPermissions:     sql.NewDbsqlPermissions(apiClient),
		Experiments:          mlflow.NewExperiments(apiClient),
		ExternalLocations:    unitycatalog.NewExternalLocations(apiClient),
		Functions:            unitycatalog.NewFunctions(apiClient),
		GitCredentials:       gitcredentials.NewGitCredentials(apiClient),
		GlobalInitScripts:    globalinitscripts.NewGlobalInitScripts(apiClient),
		Grants:               unitycatalog.NewGrants(apiClient),
		Groups:               scim.NewGroups(apiClient),
		InstancePools:        instancepools.NewInstancePools(apiClient),
		InstanceProfiles:     clusters.NewInstanceProfiles(apiClient),
		IpAccessLists:        ipaccesslists.NewIpAccessLists(apiClient),
		Jobs:                 jobs.NewJobs(apiClient),
		Libraries:            libraries.NewLibraries(apiClient),
		MLflowArtifacts:      mlflow.NewMLflowArtifacts(apiClient),
		MLflowDatabricks:     mlflow.NewMLflowDatabricks(apiClient),
		MLflowMetrics:        mlflow.NewMLflowMetrics(apiClient),
		MLflowRuns:           mlflow.NewMLflowRuns(apiClient),
		Metastores:           unitycatalog.NewMetastores(apiClient),
		ModelVersionComments: mlflow.NewModelVersionComments(apiClient),
		ModelVersions:        mlflow.NewModelVersions(apiClient),
		Permissions:          permissions.NewPermissions(apiClient),
		Pipelines:            pipelines.NewPipelines(apiClient),
		PolicyFamilies:       clusterpolicies.NewPolicyFamilies(apiClient),
		Providers:            unitycatalog.NewProviders(apiClient),
		Queries:              sql.NewQueries(apiClient),
		QueryHistory:         sql.NewQueryHistory(apiClient),
		RecipientActivation:  unitycatalog.NewRecipientActivation(apiClient),
		Recipients:           unitycatalog.NewRecipients(apiClient),
		RegisteredModels:     mlflow.NewRegisteredModels(apiClient),
		RegistryWebhooks:     mlflow.NewRegistryWebhooks(apiClient),
		Repos:                repos.NewRepos(apiClient),
		Schemas:              unitycatalog.NewSchemas(apiClient),
		Secrets:              secrets.NewSecrets(apiClient),
		ServicePrincipals:    scim.NewServicePrincipals(apiClient),
		ServingEndpoints:     endpoints.NewServingEndpoints(apiClient),
		Shares:               unitycatalog.NewShares(apiClient),
		StatementExecution:   sql.NewStatementExecution(apiClient),
		StorageCredentials:   unitycatalog.NewStorageCredentials(apiClient),
		TableConstraints:     unitycatalog.NewTableConstraints(apiClient),
		Tables:               unitycatalog.NewTables(apiClient),
		TokenManagement:      tokenmanagement.NewTokenManagement(apiClient),
		Tokens:               tokens.NewTokens(apiClient),
		TransitionRequests:   mlflow.NewTransitionRequests(apiClient),
		Users:                scim.NewUsers(apiClient),
		Warehouses:           sql.NewWarehouses(apiClient),
		Workspace:            workspace.NewWorkspace(apiClient),
		WorkspaceConf:        workspaceconf.NewWorkspaceConf(apiClient),
	}, nil
}
