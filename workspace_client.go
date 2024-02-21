// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databricks

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/dashboards"
	"github.com/databricks/databricks-sdk-go/service/files"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/databricks-sdk-go/service/serving"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/databricks-sdk-go/service/sql"
	"github.com/databricks/databricks-sdk-go/service/vectorsearch"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

type WorkspaceClient struct {
	Config    *config.Config
	apiClient *httpclient.ApiClient

	// These APIs manage access rules on resources in an account. Currently,
	// only grant rules are supported. A grant rule specifies a role assigned to
	// a set of principals. A list of rules attached to a resource is called a
	// rule set. A workspace must belong to an account for these APIs to work.
	AccountAccessControlProxy iam.AccountAccessControlProxyInterface

	// The alerts API can be used to perform CRUD operations on alerts. An alert
	// is a Databricks SQL object that periodically runs a query, evaluates a
	// condition of its result, and notifies one or more users and/or
	// notification destinations if the condition was met. Alerts can be
	// scheduled using the `sql_task` type of the Jobs API, e.g.
	// :method:jobs/create.
	Alerts sql.AlertsInterface

	// Lakehouse Apps run directly on a customer’s Databricks instance,
	// integrate with their data, use and extend Databricks services, and enable
	// users to interact through single sign-on.
	Apps serving.AppsInterface

	// In Databricks Runtime 13.3 and above, you can add libraries and init
	// scripts to the `allowlist` in UC so that users can leverage these
	// artifacts on compute configured with shared access mode.
	ArtifactAllowlists catalog.ArtifactAllowlistsInterface

	// A catalog is the first layer of Unity Catalog’s three-level namespace.
	// It’s used to organize your data assets. Users can see all catalogs on
	// which they have been assigned the USE_CATALOG data permission.
	//
	// In Unity Catalog, admins and data stewards manage users and their access
	// to data centrally across all of the workspaces in a Databricks account.
	// Users in different workspaces can share access to the same data,
	// depending on privileges granted centrally in Unity Catalog.
	Catalogs catalog.CatalogsInterface

	// A clean room is a secure, privacy-protecting environment where two or
	// more parties can share sensitive enterprise data, including customer
	// data, for measurements, insights, activation and other use cases.
	//
	// To create clean rooms, you must be a metastore admin or a user with the
	// **CREATE_CLEAN_ROOM** privilege.
	CleanRooms sharing.CleanRoomsInterface

	// You can use cluster policies to control users' ability to configure
	// clusters based on a set of rules. These rules specify which attributes or
	// attribute values can be used during cluster creation. Cluster policies
	// have ACLs that limit their use to specific users and groups.
	//
	// With cluster policies, you can: - Auto-install cluster libraries on the
	// next restart by listing them in the policy's "libraries" field (Public
	// Preview). - Limit users to creating clusters with the prescribed
	// settings. - Simplify the user interface, enabling more users to create
	// clusters, by fixing and hiding some fields. - Manage costs by setting
	// limits on attributes that impact the hourly rate.
	//
	// Cluster policy permissions limit which policies a user can select in the
	// Policy drop-down when the user creates a cluster: - A user who has
	// unrestricted cluster create permission can select the Unrestricted policy
	// and create fully-configurable clusters. - A user who has both
	// unrestricted cluster create permission and access to cluster policies can
	// select the Unrestricted policy and policies they have access to. - A user
	// that has access to only cluster policies, can select the policies they
	// have access to.
	//
	// If no policies exist in the workspace, the Policy drop-down doesn't
	// appear. Only admin users can create, edit, and delete policies. Admin
	// users also have access to all policies.
	ClusterPolicies compute.ClusterPoliciesInterface

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
	Clusters compute.ClustersInterface

	// This API allows execution of Python, Scala, SQL, or R commands on running
	// Databricks Clusters.
	CommandExecution compute.CommandExecutionInterface

	// Connections allow for creating a connection to an external data source.
	//
	// A connection is an abstraction of an external data source that can be
	// connected from Databricks Compute. Creating a connection object is the
	// first step to managing external data sources within Unity Catalog, with
	// the second step being creating a data object (catalog, schema, or table)
	// using the connection. Data objects derived from a connection can be
	// written to or read from similar to other Unity Catalog data objects based
	// on cloud storage. Users may create different types of connections with
	// each connection having a unique set of configuration options to support
	// credential management and other settings.
	Connections catalog.ConnectionsInterface

	// Credentials manager interacts with with Identity Providers to to perform
	// token exchanges using stored credentials and refresh tokens.
	CredentialsManager settings.CredentialsManagerInterface

	// This API allows retrieving information about currently authenticated user
	// or service principal.
	CurrentUser iam.CurrentUserInterface

	// This is an evolving API that facilitates the addition and removal of
	// widgets from existing dashboards within the Databricks Workspace. Data
	// structures may change over time.
	DashboardWidgets sql.DashboardWidgetsInterface

	// In general, there is little need to modify dashboards using the API.
	// However, it can be useful to use dashboard objects to look-up a
	// collection of related query IDs. The API can also be used to duplicate
	// multiple dashboards at once since you can get a dashboard definition with
	// a GET request and then POST it to create a new one. Dashboards can be
	// scheduled using the `sql_task` type of the Jobs API, e.g.
	// :method:jobs/create.
	Dashboards sql.DashboardsInterface

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
	DataSources sql.DataSourcesInterface

	// DBFS API makes it simple to interact with various data sources without
	// having to include a users credentials every time to read a file.
	Dbfs files.DbfsInterface

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
	DbsqlPermissions sql.DbsqlPermissionsInterface

	// Experiments are the primary unit of organization in MLflow; all MLflow
	// runs belong to an experiment. Each experiment lets you visualize, search,
	// and compare runs, as well as download run artifacts or metadata for
	// analysis in other tools. Experiments are maintained in a Databricks
	// hosted MLflow tracking server.
	//
	// Experiments are located in the workspace file tree. You manage
	// experiments using the same tools you use to manage other workspace
	// objects such as folders, notebooks, and libraries.
	Experiments ml.ExperimentsInterface

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
	ExternalLocations catalog.ExternalLocationsInterface

	// The Files API allows you to read, write, list, and delete files and
	// directories. We support Unity Catalog volumes with paths starting with
	// "/Volumes/<catalog>/<schema>/<volume>".
	//
	// The Files API is designed like a standard HTTP API, rather than as a JSON
	// RPC API. This is intended to make it easier and more efficient to work
	// with file contents as raw bytes.
	//
	// Because the Files API is a standard HTTP API, the URI path is used to
	// specify the file or directory to operate on. The path is always absolute.
	//
	// The Files API has separate endpoints for working with files, `/fs/files`,
	// and working with directories, `/fs/directories`. The standard HTTP
	// methods `GET`, `HEAD`, `PUT`, and `DELETE` work as expected on these
	// endpoints.
	Files files.FilesInterface

	// Functions implement User-Defined Functions (UDFs) in Unity Catalog.
	//
	// The function implementation can be any SQL expression or Query, and it
	// can be invoked wherever a table reference is allowed in a query. In Unity
	// Catalog, a function resides at the same level as a table, so it can be
	// referenced with the form
	// __catalog_name__.__schema_name__.__function_name__.
	Functions catalog.FunctionsInterface

	// Registers personal access token for Databricks to do operations on behalf
	// of the user.
	//
	// See [more info].
	//
	// [more info]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	GitCredentials workspace.GitCredentialsInterface

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
	GlobalInitScripts compute.GlobalInitScriptsInterface

	// In Unity Catalog, data is secure by default. Initially, users have no
	// access to data in a metastore. Access can be granted by either a
	// metastore admin, the owner of an object, or the owner of the catalog or
	// schema that contains the object. Securable objects in Unity Catalog are
	// hierarchical and privileges are inherited downward.
	//
	// Securable objects in Unity Catalog are hierarchical and privileges are
	// inherited downward. This means that granting a privilege on the catalog
	// automatically grants the privilege to all current and future objects
	// within the catalog. Similarly, privileges granted on a schema are
	// inherited by all current and future objects within that schema.
	Grants catalog.GrantsInterface

	// Groups simplify identity management, making it easier to assign access to
	// Databricks workspace, data, and other securable objects.
	//
	// It is best practice to assign access to workspaces and access-control
	// policies in Unity Catalog to groups, instead of to users individually.
	// All Databricks workspace identities can be assigned as members of groups,
	// and members inherit permissions that are assigned to their group.
	Groups iam.GroupsInterface

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
	InstancePools compute.InstancePoolsInterface

	// The Instance Profiles API allows admins to add, list, and remove instance
	// profiles that users can launch clusters with. Regular users can list the
	// instance profiles available to them. See [Secure access to S3 buckets]
	// using instance profiles for more information.
	//
	// [Secure access to S3 buckets]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html
	InstanceProfiles compute.InstanceProfilesInterface

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
	IpAccessLists settings.IpAccessListsInterface

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
	// [Secrets CLI] to manage secrets in the [Databricks CLI]. Use the [Secrets
	// utility] to reference secrets in notebooks and jobs.
	//
	// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
	// [Secrets CLI]: https://docs.databricks.com/dev-tools/cli/secrets-cli.html
	// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
	Jobs jobs.JobsInterface

	// A monitor computes and monitors data or model quality metrics for a table
	// over time. It generates metrics tables and a dashboard that you can use
	// to monitor table health and set alerts.
	//
	// Most write operations require the user to be the owner of the table (or
	// its parent schema or parent catalog). Viewing the dashboard, computed
	// metrics, or monitor configuration only requires the user to have
	// **SELECT** privileges on the table (along with **USE_SCHEMA** and
	// **USE_CATALOG**).
	LakehouseMonitors catalog.LakehouseMonitorsInterface

	// These APIs provide specific management operations for Lakeview
	// dashboards. Generic resource management can be done with Workspace API
	// (import, export, get-status, list, delete).
	Lakeview dashboards.LakeviewInterface

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
	Libraries compute.LibrariesInterface

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
	// available in a catalog named hive_metastore.
	Metastores catalog.MetastoresInterface

	// Note: This API reference documents APIs for the Workspace Model Registry.
	// Databricks recommends using [Models in Unity
	// Catalog](/api/workspace/registeredmodels) instead. Models in Unity
	// Catalog provides centralized model governance, cross-workspace access,
	// lineage, and deployment. Workspace Model Registry will be deprecated in
	// the future.
	//
	// The Workspace Model Registry is a centralized model repository and a UI
	// and set of APIs that enable you to manage the full lifecycle of MLflow
	// Models.
	ModelRegistry ml.ModelRegistryInterface

	// Databricks provides a hosted version of MLflow Model Registry in Unity
	// Catalog. Models in Unity Catalog provide centralized access control,
	// auditing, lineage, and discovery of ML models across Databricks
	// workspaces.
	//
	// This API reference documents the REST endpoints for managing model
	// versions in Unity Catalog. For more details, see the [registered models
	// API docs](/api/workspace/registeredmodels).
	ModelVersions catalog.ModelVersionsInterface

	// Online tables provide lower latency and higher QPS access to data from
	// Delta tables.
	OnlineTables catalog.OnlineTablesInterface

	// Permissions API are used to create read, write, edit, update and manage
	// access for various users on different objects and endpoints.
	//
	// * **[Cluster permissions](:service:clusters)** — Manage which users can
	// manage, restart, or attach to clusters.
	//
	// * **[Cluster policy permissions](:service:clusterpolicies)** — Manage
	// which users can use cluster policies.
	//
	// * **[Delta Live Tables pipeline permissions](:service:pipelines)** —
	// Manage which users can view, manage, run, cancel, or own a Delta Live
	// Tables pipeline.
	//
	// * **[Job permissions](:service:jobs)** — Manage which users can view,
	// manage, trigger, cancel, or own a job.
	//
	// * **[MLflow experiment permissions](:service:experiments)** — Manage
	// which users can read, edit, or manage MLflow experiments.
	//
	// * **[MLflow registered model permissions](:service:modelregistry)** —
	// Manage which users can read, edit, or manage MLflow registered models.
	//
	// * **[Password permissions](:service:users)** — Manage which users can
	// use password login when SSO is enabled.
	//
	// * **[Instance Pool permissions](:service:instancepools)** — Manage
	// which users can manage or attach to pools.
	//
	// * **[Repo permissions](repos)** — Manage which users can read, run,
	// edit, or manage a repo.
	//
	// * **[Serving endpoint permissions](:service:servingendpoints)** —
	// Manage which users can view, query, or manage a serving endpoint.
	//
	// * **[SQL warehouse permissions](:service:warehouses)** — Manage which
	// users can use or manage SQL warehouses.
	//
	// * **[Token permissions](:service:tokenmanagement)** — Manage which
	// users can create or use tokens.
	//
	// * **[Workspace object permissions](:service:workspace)** — Manage which
	// users can read, run, edit, or manage directories, files, and notebooks.
	//
	// For the mapping of the required permissions for specific actions or
	// abilities and other important information, see [Access Control].
	//
	// [Access Control]: https://docs.databricks.com/security/auth-authz/access-control/index.html
	Permissions iam.PermissionsInterface

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
	Pipelines pipelines.PipelinesInterface

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
	PolicyFamilies compute.PolicyFamiliesInterface

	// A data provider is an object representing the organization in the real
	// world who shares the data. A provider contains shares which further
	// contain the shared data.
	Providers sharing.ProvidersInterface

	// These endpoints are used for CRUD operations on query definitions. Query
	// definitions include the target SQL warehouse, query text, name,
	// description, tags, parameters, and visualizations. Queries can be
	// scheduled using the `sql_task` type of the Jobs API, e.g.
	// :method:jobs/create.
	Queries sql.QueriesInterface

	// Access the history of queries through SQL warehouses.
	QueryHistory sql.QueryHistoryInterface

	// This is an evolving API that facilitates the addition and removal of
	// vizualisations from existing queries within the Databricks Workspace.
	// Data structures may change over time.
	QueryVisualizations sql.QueryVisualizationsInterface

	// The Recipient Activation API is only applicable in the open sharing model
	// where the recipient object has the authentication type of `TOKEN`. The
	// data recipient follows the activation link shared by the data provider to
	// download the credential file that includes the access token. The
	// recipient will then use the credential file to establish a secure
	// connection with the provider to receive the shared data.
	//
	// Note that you can download the credential file only once. Recipients
	// should treat the downloaded credential as a secret and must not share it
	// outside of their organization.
	RecipientActivation sharing.RecipientActivationInterface

	// A recipient is an object you create using :method:recipients/create to
	// represent an organization which you want to allow access shares. The way
	// how sharing works differs depending on whether or not your recipient has
	// access to a Databricks workspace that is enabled for Unity Catalog:
	//
	// - For recipients with access to a Databricks workspace that is enabled
	// for Unity Catalog, you can create a recipient object along with a unique
	// sharing identifier you get from the recipient. The sharing identifier is
	// the key identifier that enables the secure connection. This sharing mode
	// is called **Databricks-to-Databricks sharing**.
	//
	// - For recipients without access to a Databricks workspace that is enabled
	// for Unity Catalog, when you create a recipient object, Databricks
	// generates an activation link you can send to the recipient. The recipient
	// follows the activation link to download the credential file, and then
	// uses the credential file to establish a secure connection to receive the
	// shared data. This sharing mode is called **open sharing**.
	Recipients sharing.RecipientsInterface

	// Databricks provides a hosted version of MLflow Model Registry in Unity
	// Catalog. Models in Unity Catalog provide centralized access control,
	// auditing, lineage, and discovery of ML models across Databricks
	// workspaces.
	//
	// An MLflow registered model resides in the third layer of Unity
	// Catalog’s three-level namespace. Registered models contain model
	// versions, which correspond to actual ML models (MLflow models). Creating
	// new model versions currently requires use of the MLflow Python client.
	// Once model versions are created, you can load them for batch inference
	// using MLflow Python client APIs, or deploy them for real-time serving
	// using Databricks Model Serving.
	//
	// All operations on registered models and model versions require
	// USE_CATALOG permissions on the enclosing catalog and USE_SCHEMA
	// permissions on the enclosing schema. In addition, the following
	// additional privileges are required for various operations:
	//
	// * To create a registered model, users must additionally have the
	// CREATE_MODEL permission on the target schema. * To view registered model
	// or model version metadata, model version data files, or invoke a model
	// version, users must additionally have the EXECUTE permission on the
	// registered model * To update registered model or model version tags,
	// users must additionally have APPLY TAG permissions on the registered
	// model * To update other registered model or model version metadata
	// (comments, aliases) create a new model version, or update permissions on
	// the registered model, users must be owners of the registered model.
	//
	// Note: The securable type for models is "FUNCTION". When using REST APIs
	// (e.g. tagging, grants) that specify a securable type, use "FUNCTION" as
	// the securable type.
	RegisteredModels catalog.RegisteredModelsInterface

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
	Repos workspace.ReposInterface

	// A schema (also called a database) is the second layer of Unity
	// Catalog’s three-level namespace. A schema organizes tables, views and
	// functions. To access (or list) a table or view in a schema, users must
	// have the USE_SCHEMA data permission on the schema and its parent catalog,
	// and they must have the SELECT permission on the table or view.
	Schemas catalog.SchemasInterface

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
	Secrets workspace.SecretsInterface

	// Identities for use with jobs, automated tools, and systems such as
	// scripts, apps, and CI/CD platforms. Databricks recommends creating
	// service principals to run production jobs or modify production data. If
	// all processes that act on production data run with service principals,
	// interactive users do not need any write, delete, or modify privileges in
	// production. This eliminates the risk of a user overwriting production
	// data by accident.
	ServicePrincipals iam.ServicePrincipalsInterface

	// The Serving Endpoints API allows you to create, update, and delete model
	// serving endpoints.
	//
	// You can use a serving endpoint to serve models from the Databricks Model
	// Registry or from Unity Catalog. Endpoints expose the underlying models as
	// scalable REST API endpoints using serverless compute. This means the
	// endpoints and associated compute resources are fully managed by
	// Databricks and will not appear in your cloud account. A serving endpoint
	// can consist of one or more MLflow models from the Databricks Model
	// Registry, called served entities. A serving endpoint can have at most ten
	// served entities. You can configure traffic settings to define how
	// requests should be routed to your served entities behind an endpoint.
	// Additionally, you can configure the scale of resources that should be
	// applied to each served entity.
	ServingEndpoints serving.ServingEndpointsInterface

	// The default namespace setting API allows users to configure the default
	// namespace for a Databricks workspace.
	//
	// Through this API, users can retrieve, set, or modify the default
	// namespace used when queries do not reference a fully qualified
	// three-level name. For example, if you use the API to set 'retail_prod' as
	// the default catalog, then a query 'SELECT * FROM myTable' would reference
	// the object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed).
	//
	// This setting requires a restart of clusters and SQL warehouses to take
	// effect. Additionally, the default namespace only applies when using Unity
	// Catalog-enabled compute.
	Settings settings.SettingsInterface

	// A share is a container instantiated with :method:shares/create. Once
	// created you can iteratively register a collection of existing data assets
	// defined within the metastore using :method:shares/update. You can
	// register data assets under their original name, qualified by their
	// original schema, or provide alternate exposed names.
	Shares sharing.SharesInterface

	// The Databricks SQL Statement Execution API can be used to execute SQL
	// statements on a SQL warehouse and fetch the result.
	//
	// **Getting started**
	//
	// We suggest beginning with the [Databricks SQL Statement Execution API
	// tutorial].
	//
	// **Overview of statement execution and result fetching**
	//
	// Statement execution begins by issuing a
	// :method:statementexecution/executeStatement request with a valid SQL
	// statement and warehouse ID, along with optional parameters such as the
	// data catalog and output format. If no other parameters are specified, the
	// server will wait for up to 10s before returning a response. If the
	// statement has completed within this timespan, the response will include
	// the result data as a JSON array and metadata. Otherwise, if no result is
	// available after the 10s timeout expired, the response will provide the
	// statement ID that can be used to poll for results by using a
	// :method:statementexecution/getStatement request.
	//
	// You can specify whether the call should behave synchronously,
	// asynchronously or start synchronously with a fallback to asynchronous
	// execution. This is controlled with the `wait_timeout` and
	// `on_wait_timeout` settings. If `wait_timeout` is set between 5-50 seconds
	// (default: 10s), the call waits for results up to the specified timeout;
	// when set to `0s`, the call is asynchronous and responds immediately with
	// a statement ID. The `on_wait_timeout` setting specifies what should
	// happen when the timeout is reached while the statement execution has not
	// yet finished. This can be set to either `CONTINUE`, to fallback to
	// asynchronous mode, or it can be set to `CANCEL`, which cancels the
	// statement.
	//
	// In summary: - Synchronous mode - `wait_timeout=30s` and
	// `on_wait_timeout=CANCEL` - The call waits up to 30 seconds; if the
	// statement execution finishes within this time, the result data is
	// returned directly in the response. If the execution takes longer than 30
	// seconds, the execution is canceled and the call returns with a `CANCELED`
	// state. - Asynchronous mode - `wait_timeout=0s` (`on_wait_timeout` is
	// ignored) - The call doesn't wait for the statement to finish but returns
	// directly with a statement ID. The status of the statement execution can
	// be polled by issuing :method:statementexecution/getStatement with the
	// statement ID. Once the execution has succeeded, this call also returns
	// the result and metadata in the response. - Hybrid mode (default) -
	// `wait_timeout=10s` and `on_wait_timeout=CONTINUE` - The call waits for up
	// to 10 seconds; if the statement execution finishes within this time, the
	// result data is returned directly in the response. If the execution takes
	// longer than 10 seconds, a statement ID is returned. The statement ID can
	// be used to fetch status and results in the same way as in the
	// asynchronous mode.
	//
	// Depending on the size, the result can be split into multiple chunks. If
	// the statement execution is successful, the statement response contains a
	// manifest and the first chunk of the result. The manifest contains schema
	// information and provides metadata for each chunk in the result. Result
	// chunks can be retrieved by index with
	// :method:statementexecution/getStatementResultChunkN which may be called
	// in any order and in parallel. For sequential fetching, each chunk, apart
	// from the last, also contains a `next_chunk_index` and
	// `next_chunk_internal_link` that point to the next chunk.
	//
	// A statement can be canceled with
	// :method:statementexecution/cancelExecution.
	//
	// **Fetching result data: format and disposition**
	//
	// To specify the format of the result data, use the `format` field, which
	// can be set to one of the following options: `JSON_ARRAY` (JSON),
	// `ARROW_STREAM` ([Apache Arrow Columnar]), or `CSV`.
	//
	// There are two ways to receive statement results, controlled by the
	// `disposition` setting, which can be either `INLINE` or `EXTERNAL_LINKS`:
	//
	// - `INLINE`: In this mode, the result data is directly included in the
	// response. It's best suited for smaller results. This mode can only be
	// used with the `JSON_ARRAY` format.
	//
	// - `EXTERNAL_LINKS`: In this mode, the response provides links that can be
	// used to download the result data in chunks separately. This approach is
	// ideal for larger results and offers higher throughput. This mode can be
	// used with all the formats: `JSON_ARRAY`, `ARROW_STREAM`, and `CSV`.
	//
	// By default, the API uses `format=JSON_ARRAY` and `disposition=INLINE`.
	//
	// **Limits and limitations**
	//
	// Note: The byte limit for INLINE disposition is based on internal storage
	// metrics and will not exactly match the byte count of the actual payload.
	//
	// - Statements with `disposition=INLINE` are limited to 25 MiB and will
	// fail when this limit is exceeded. - Statements with
	// `disposition=EXTERNAL_LINKS` are limited to 100 GiB. Result sets larger
	// than this limit will be truncated. Truncation is indicated by the
	// `truncated` field in the result manifest. - The maximum query text size
	// is 16 MiB. - Cancelation might silently fail. A successful response from
	// a cancel request indicates that the cancel request was successfully
	// received and sent to the processing engine. However, an outstanding
	// statement might have already completed execution when the cancel request
	// arrives. Polling for status until a terminal state is reached is a
	// reliable way to determine the final state. - Wait timeouts are
	// approximate, occur server-side, and cannot account for things such as
	// caller delays and network latency from caller to service. - The system
	// will auto-close a statement after one hour if the client stops polling
	// and thus you must poll at least once an hour. - The results are only
	// available for one hour after success; polling does not extend this.
	//
	// [Apache Arrow Columnar]: https://arrow.apache.org/overview/
	// [Databricks SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
	StatementExecution sql.StatementExecutionInterface

	// A storage credential represents an authentication and authorization
	// mechanism for accessing data stored on your cloud tenant. Each storage
	// credential is subject to Unity Catalog access-control policies that
	// control which users and groups can access the credential. If a user does
	// not have access to a storage credential in Unity Catalog, the request
	// fails and Unity Catalog does not attempt to authenticate to your cloud
	// tenant on the user’s behalf.
	//
	// Databricks recommends using external locations rather than using storage
	// credentials directly.
	//
	// To create storage credentials, you must be a Databricks account admin.
	// The account admin who creates the storage credential can delegate
	// ownership to another user or group to manage permissions on it.
	StorageCredentials catalog.StorageCredentialsInterface

	// A system schema is a schema that lives within the system catalog. A
	// system schema may contain information about customer usage of Unity
	// Catalog such as audit-logs, billing-logs, lineage information, etc.
	SystemSchemas catalog.SystemSchemasInterface

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
	TableConstraints catalog.TableConstraintsInterface

	// A table resides in the third layer of Unity Catalog’s three-level
	// namespace. It contains rows of data. To create a table, users must have
	// CREATE_TABLE and USE_SCHEMA permissions on the schema, and they must have
	// the USE_CATALOG permission on its parent catalog. To query a table, users
	// must have the SELECT permission on the table, and they must have the
	// USE_CATALOG permission on its parent catalog and the USE_SCHEMA
	// permission on its parent schema.
	//
	// A table can be managed or external. From an API perspective, a __VIEW__
	// is a particular kind of table (rather than a managed or external table).
	Tables catalog.TablesInterface

	// Enables administrators to get all tokens and delete tokens for other
	// users. Admins can either get every token, get a specific token by ID, or
	// get all tokens for a particular user.
	TokenManagement settings.TokenManagementInterface

	// The Token API allows you to create, list, and revoke tokens that can be
	// used to authenticate and access Databricks REST APIs.
	Tokens settings.TokensInterface

	// User identities recognized by Databricks and represented by email
	// addresses.
	//
	// Databricks recommends using SCIM provisioning to sync users and groups
	// automatically from your identity provider to your Databricks workspace.
	// SCIM streamlines onboarding a new employee or team by using your identity
	// provider to create users and groups in Databricks workspace and give them
	// the proper level of access. When a user leaves your organization or no
	// longer needs access to Databricks workspace, admins can terminate the
	// user in your identity provider and that user’s account will also be
	// removed from Databricks workspace. This ensures a consistent offboarding
	// process and prevents unauthorized users from accessing sensitive data.
	Users iam.UsersInterface

	// **Endpoint**: Represents the compute resources to host vector search
	// indexes.
	VectorSearchEndpoints vectorsearch.VectorSearchEndpointsInterface

	// **Index**: An efficient representation of your embedding vectors that
	// supports real-time and efficient approximate nearest neighbor (ANN)
	// search queries.
	//
	// There are 2 types of Vector Search indexes: * **Delta Sync Index**: An
	// index that automatically syncs with a source Delta Table, automatically
	// and incrementally updating the index as the underlying data in the Delta
	// Table changes. * **Direct Vector Access Index**: An index that supports
	// direct read and write of vectors and metadata through our REST and SDK
	// APIs. With this model, the user manages index updates.
	VectorSearchIndexes vectorsearch.VectorSearchIndexesInterface

	// Volumes are a Unity Catalog (UC) capability for accessing, storing,
	// governing, organizing and processing files. Use cases include running
	// machine learning on unstructured data such as image, audio, video, or PDF
	// files, organizing data sets during the data exploration stages in data
	// science, working with libraries that require access to the local file
	// system on cluster machines, storing library and config files of arbitrary
	// formats such as .whl or .txt centrally and providing secure access across
	// workspaces to it, or transforming and querying non-tabular data files in
	// ETL.
	Volumes catalog.VolumesInterface

	// A SQL warehouse is a compute resource that lets you run SQL commands on
	// data objects within Databricks SQL. Compute resources are infrastructure
	// resources that provide processing capabilities in the cloud.
	Warehouses sql.WarehousesInterface

	// The Workspace API allows you to list, import, export, and delete
	// notebooks and folders.
	//
	// A notebook is a web-based interface to a document that contains runnable
	// code, visualizations, and explanatory text.
	Workspace workspace.WorkspaceInterface

	// A securable in Databricks can be configured as __OPEN__ or __ISOLATED__.
	// An __OPEN__ securable can be accessed from any workspace, while an
	// __ISOLATED__ securable can only be accessed from a configured list of
	// workspaces. This API allows you to configure (bind) securables to
	// workspaces.
	//
	// NOTE: The __isolation_mode__ is configured for the securable itself
	// (using its Update method) and the workspace bindings are only consulted
	// when the securable's __isolation_mode__ is set to __ISOLATED__.
	//
	// A securable's workspace bindings can be configured by a metastore admin
	// or the owner of the securable.
	//
	// The original path
	// (/api/2.1/unity-catalog/workspace-bindings/catalogs/{name}) is
	// deprecated. Please use the new path
	// (/api/2.1/unity-catalog/bindings/{securable_type}/{securable_name}) which
	// introduces the ability to bind a securable in READ_ONLY mode (catalogs
	// only).
	//
	// Securables that support binding: - catalog
	WorkspaceBindings catalog.WorkspaceBindingsInterface

	// This API allows updating known workspace settings for advanced users.
	WorkspaceConf settings.WorkspaceConfInterface
}

var ErrNotWorkspaceClient = errors.New("invalid Databricks Workspace configuration")

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
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	if cfg.IsAccountClient() {
		return nil, ErrNotWorkspaceClient
	}
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}
	return &WorkspaceClient{
		Config:    cfg,
		apiClient: apiClient,

		AccountAccessControlProxy: iam.NewAccountAccessControlProxy(databricksClient),
		Alerts:                    sql.NewAlerts(databricksClient),
		Apps:                      serving.NewApps(databricksClient),
		ArtifactAllowlists:        catalog.NewArtifactAllowlists(databricksClient),
		Catalogs:                  catalog.NewCatalogs(databricksClient),
		CleanRooms:                sharing.NewCleanRooms(databricksClient),
		ClusterPolicies:           compute.NewClusterPolicies(databricksClient),
		Clusters:                  compute.NewClusters(databricksClient),
		CommandExecution:          compute.NewCommandExecution(databricksClient),
		Connections:               catalog.NewConnections(databricksClient),
		CredentialsManager:        settings.NewCredentialsManager(databricksClient),
		CurrentUser:               iam.NewCurrentUser(databricksClient),
		DashboardWidgets:          sql.NewDashboardWidgets(databricksClient),
		Dashboards:                sql.NewDashboards(databricksClient),
		DataSources:               sql.NewDataSources(databricksClient),
		Dbfs:                      files.NewDbfs(databricksClient),
		DbsqlPermissions:          sql.NewDbsqlPermissions(databricksClient),
		Experiments:               ml.NewExperiments(databricksClient),
		ExternalLocations:         catalog.NewExternalLocations(databricksClient),
		Files:                     files.NewFiles(databricksClient),
		Functions:                 catalog.NewFunctions(databricksClient),
		GitCredentials:            workspace.NewGitCredentials(databricksClient),
		GlobalInitScripts:         compute.NewGlobalInitScripts(databricksClient),
		Grants:                    catalog.NewGrants(databricksClient),
		Groups:                    iam.NewGroups(databricksClient),
		InstancePools:             compute.NewInstancePools(databricksClient),
		InstanceProfiles:          compute.NewInstanceProfiles(databricksClient),
		IpAccessLists:             settings.NewIpAccessLists(databricksClient),
		Jobs:                      jobs.NewJobs(databricksClient),
		LakehouseMonitors:         catalog.NewLakehouseMonitors(databricksClient),
		Lakeview:                  dashboards.NewLakeview(databricksClient),
		Libraries:                 compute.NewLibraries(databricksClient),
		Metastores:                catalog.NewMetastores(databricksClient),
		ModelRegistry:             ml.NewModelRegistry(databricksClient),
		ModelVersions:             catalog.NewModelVersions(databricksClient),
		OnlineTables:              catalog.NewOnlineTables(databricksClient),
		Permissions:               iam.NewPermissions(databricksClient),
		Pipelines:                 pipelines.NewPipelines(databricksClient),
		PolicyFamilies:            compute.NewPolicyFamilies(databricksClient),
		Providers:                 sharing.NewProviders(databricksClient),
		Queries:                   sql.NewQueries(databricksClient),
		QueryHistory:              sql.NewQueryHistory(databricksClient),
		QueryVisualizations:       sql.NewQueryVisualizations(databricksClient),
		RecipientActivation:       sharing.NewRecipientActivation(databricksClient),
		Recipients:                sharing.NewRecipients(databricksClient),
		RegisteredModels:          catalog.NewRegisteredModels(databricksClient),
		Repos:                     workspace.NewRepos(databricksClient),
		Schemas:                   catalog.NewSchemas(databricksClient),
		Secrets:                   workspace.NewSecrets(databricksClient),
		ServicePrincipals:         iam.NewServicePrincipals(databricksClient),
		ServingEndpoints:          serving.NewServingEndpoints(databricksClient),
		Settings:                  settings.NewSettings(databricksClient),
		Shares:                    sharing.NewShares(databricksClient),
		StatementExecution:        sql.NewStatementExecution(databricksClient),
		StorageCredentials:        catalog.NewStorageCredentials(databricksClient),
		SystemSchemas:             catalog.NewSystemSchemas(databricksClient),
		TableConstraints:          catalog.NewTableConstraints(databricksClient),
		Tables:                    catalog.NewTables(databricksClient),
		TokenManagement:           settings.NewTokenManagement(databricksClient),
		Tokens:                    settings.NewTokens(databricksClient),
		Users:                     iam.NewUsers(databricksClient),
		VectorSearchEndpoints:     vectorsearch.NewVectorSearchEndpoints(databricksClient),
		VectorSearchIndexes:       vectorsearch.NewVectorSearchIndexes(databricksClient),
		Volumes:                   catalog.NewVolumes(databricksClient),
		Warehouses:                sql.NewWarehouses(databricksClient),
		Workspace:                 workspace.NewWorkspace(databricksClient),
		WorkspaceBindings:         catalog.NewWorkspaceBindings(databricksClient),
		WorkspaceConf:             settings.NewWorkspaceConf(databricksClient),
	}, nil
}
