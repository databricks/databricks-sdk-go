// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"context"
)

// Cluster policy limits the ability to configure clusters based on a set of
// rules. The policy rules limit the attributes or attribute values available
// for cluster creation. Cluster policies have ACLs that limit their use to
// specific users and groups.
//
// Cluster policies let you limit users to create clusters with prescribed
// settings, simplify the user interface and enable more users to create their
// own clusters (by fixing and hiding some values), control cost by limiting per
// cluster maximum cost (by setting limits on attributes whose values contribute
// to hourly price).
//
// Cluster policy permissions limit which policies a user can select in the
// Policy drop-down when the user creates a cluster: - A user who has cluster
// create permission can select the Unrestricted policy and create
// fully-configurable clusters. - A user who has both cluster create permission
// and access to cluster policies can select the Unrestricted policy and
// policies they have access to. - A user that has access to only cluster
// policies, can select the policies they have access to.
//
// If no policies have been created in the workspace, the Policy drop-down does
// not display.
//
// Only admin users can create, edit, and delete policies. Admin users also have
// access to all policies.
type ClusterPoliciesService interface {

	// Create a new policy.
	//
	// Creates a new policy with prescribed settings.
	Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error)

	// Delete a cluster policy.
	//
	// Delete a policy for a cluster. Clusters governed by this policy can still
	// run, but cannot be edited.
	Delete(ctx context.Context, request DeletePolicy) error

	// Update a cluster policy.
	//
	// Update an existing policy for cluster. This operation may make some
	// clusters governed by the previous policy invalid.
	Edit(ctx context.Context, request EditPolicy) error

	// Get entity.
	//
	// Get a cluster policy entity. Creation and editing is available to admins
	// only.
	Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error)

	// Get a cluster policy.
	//
	// Returns a list of policies accessible by the requesting user.
	//
	// Use ListAll() to get all Policy instances
	List(ctx context.Context, request ListClusterPoliciesRequest) (*ListPoliciesResponse, error)
}

// The Clusters API allows you to create, start, edit, list, terminate, and
// delete clusters.
//
// Databricks maps cluster node instance types to compute units known as DBUs.
// See the instance type pricing page for a list of the supported instance types
// and their corresponding DBUs.
//
// A Databricks cluster is a set of computation resources and configurations on
// which you run data engineering, data science, and data analytics workloads,
// such as production ETL pipelines, streaming analytics, ad-hoc analytics, and
// machine learning.
//
// You run these workloads as a set of commands in a notebook or as an automated
// job. Databricks makes a distinction between all-purpose clusters and job
// clusters. You use all-purpose clusters to analyze data collaboratively using
// interactive notebooks. You use job clusters to run fast and robust automated
// jobs.
//
// You can create an all-purpose cluster using the UI, CLI, or REST API. You can
// manually terminate and restart an all-purpose cluster. Multiple users can
// share such clusters to do collaborative interactive analysis.
//
// IMPORTANT: Databricks retains cluster configuration information for up to 200
// all-purpose clusters terminated in the last 30 days and up to 30 job clusters
// recently terminated by the job scheduler. To keep an all-purpose cluster
// configuration even after it has been terminated for more than 30 days, an
// administrator can pin a cluster to the cluster list.
type ClustersService interface {

	// Change cluster owner.
	//
	// Change the owner of the cluster. You must be an admin to perform this
	// operation.
	ChangeOwner(ctx context.Context, request ChangeClusterOwner) error

	// Create new cluster.
	//
	// Creates a new Spark cluster. This method will acquire new instances from
	// the cloud provider if necessary. This method is asynchronous; the
	// returned `cluster_id` can be used to poll the cluster status. When this
	// method returns, the cluster will be in a `PENDING` state. The cluster
	// will be usable once it enters a `RUNNING` state.
	//
	// Note: Databricks may not be able to acquire some of the requested nodes,
	// due to cloud provider limitations (account limits, spot price, etc.) or
	// transient network issues.
	//
	// If Databricks acquires at least 85% of the requested on-demand nodes,
	// cluster creation will succeed. Otherwise the cluster will terminate with
	// an informative error message.
	Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error)

	// Terminate cluster.
	//
	// Terminates the Spark cluster with the specified ID. The cluster is
	// removed asynchronously. Once the termination has completed, the cluster
	// will be in a `TERMINATED` state. If the cluster is already in a
	// `TERMINATING` or `TERMINATED` state, nothing will happen.
	Delete(ctx context.Context, request DeleteCluster) error

	// Update cluster configuration.
	//
	// Updates the configuration of a cluster to match the provided attributes
	// and size. A cluster can be updated if it is in a `RUNNING` or
	// `TERMINATED` state.
	//
	// If a cluster is updated while in a `RUNNING` state, it will be restarted
	// so that the new attributes can take effect.
	//
	// If a cluster is updated while in a `TERMINATED` state, it will remain
	// `TERMINATED`. The next time it is started using the `clusters/start` API,
	// the new attributes will take effect. Any attempt to update a cluster in
	// any other state will be rejected with an `INVALID_STATE` error code.
	//
	// Clusters created by the Databricks Jobs service cannot be edited.
	Edit(ctx context.Context, request EditCluster) error

	// List cluster activity events.
	//
	// Retrieves a list of events about the activity of a cluster. This API is
	// paginated. If there are more events to read, the response includes all
	// the nparameters necessary to request the next page of events.
	//
	// Use EventsAll() to get all ClusterEvent instances, which will iterate over every result page.
	Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error)

	// Get cluster info.
	//
	// "Retrieves the information for a cluster given its identifier. Clusters
	// can be described while they are running, or up to 60 days after they are
	// terminated.
	Get(ctx context.Context, request GetClusterRequest) (*ClusterInfo, error)

	// List all clusters.
	//
	// Return information about all pinned clusters, active clusters, up to 200
	// of the most recently terminated all-purpose clusters in the past 30 days,
	// and up to 30 of the most recently terminated job clusters in the past 30
	// days.
	//
	// For example, if there is 1 pinned cluster, 4 active clusters, 45
	// terminated all-purpose clusters in the past 30 days, and 50 terminated
	// job clusters in the past 30 days, then this API returns the 1 pinned
	// cluster, 4 active clusters, all 45 terminated all-purpose clusters, and
	// the 30 most recently terminated job clusters.
	//
	// Use ListAll() to get all ClusterInfo instances
	List(ctx context.Context, request ListClustersRequest) (*ListClustersResponse, error)

	// List node types.
	//
	// Returns a list of supported Spark node types. These node types can be
	// used to launch a cluster.
	ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error)

	// List availability zones.
	//
	// Returns a list of availability zones where clusters can be created in
	// (For example, us-west-2a). These zones can be used to launch a cluster.
	ListZones(ctx context.Context) (*ListAvailableZonesResponse, error)

	// Permanently delete cluster.
	//
	// Permanently deletes a Spark cluster. This cluster is terminated and
	// resources are asynchronously removed.
	//
	// In addition, users will no longer see permanently deleted clusters in the
	// cluster list, and API users can no longer perform any action on
	// permanently deleted clusters.
	PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error

	// Pin cluster.
	//
	// Pinning a cluster ensures that the cluster will always be returned by the
	// ListClusters API. Pinning a cluster that is already pinned will have no
	// effect. This API can only be called by workspace admins.
	Pin(ctx context.Context, request PinCluster) error

	// Resize cluster.
	//
	// Resizes a cluster to have a desired number of workers. This will fail
	// unless the cluster is in a `RUNNING` state.
	Resize(ctx context.Context, request ResizeCluster) error

	// Restart cluster.
	//
	// Restarts a Spark cluster with the supplied ID. If the cluster is not
	// currently in a `RUNNING` state, nothing will happen.
	Restart(ctx context.Context, request RestartCluster) error

	// List available Spark versions.
	//
	// Returns the list of available Spark versions. These versions can be used
	// to launch a cluster.
	SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error)

	// Start terminated cluster.
	//
	// Starts a terminated Spark cluster with the supplied ID. This works
	// similar to `createCluster` except:
	//
	// * The previous cluster id and attributes are preserved. * The cluster
	// starts with the last specified cluster size. * If the previous cluster
	// was an autoscaling cluster, the current cluster starts with the minimum
	// number of nodes. * If the cluster is not currently in a `TERMINATED`
	// state, nothing will happen. * Clusters launched to run a job cannot be
	// started.
	Start(ctx context.Context, request StartCluster) error

	// Unpin cluster.
	//
	// Unpinning a cluster will allow the cluster to eventually be removed from
	// the ListClusters API. Unpinning a cluster that is not pinned will have no
	// effect. This API can only be called by workspace admins.
	Unpin(ctx context.Context, request UnpinCluster) error
}

// This API allows execution of Python, Scala, SQL, or R commands on running
// Databricks Clusters.
type CommandExecutionService interface {

	// Cancel a command.
	//
	// Cancels a currently running command within an execution context.
	//
	// The command ID is obtained from a prior successful call to __execute__.
	Cancel(ctx context.Context, request CancelCommand) error

	// Get command info.
	//
	// Gets the status of and, if available, the results from a currently
	// executing command.
	//
	// The command ID is obtained from a prior successful call to __execute__.
	CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error)

	// Get status.
	//
	// Gets the status for an execution context.
	ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error)

	// Create an execution context.
	//
	// Creates an execution context for running cluster commands.
	//
	// If successful, this method returns the ID of the new execution context.
	Create(ctx context.Context, request CreateContext) (*Created, error)

	// Delete an execution context.
	//
	// Deletes an execution context.
	Destroy(ctx context.Context, request DestroyContext) error

	// Run a command.
	//
	// Runs a cluster command in the given execution context, using the provided
	// language.
	//
	// If successful, it returns an ID for tracking the status of the command's
	// execution.
	Execute(ctx context.Context, request Command) (*Created, error)
}

// The Global Init Scripts API enables Workspace administrators to configure
// global initialization scripts for their workspace. These scripts run on every
// node in every cluster in the workspace.
//
// **Important:** Existing clusters must be restarted to pick up any changes
// made to global init scripts. Global init scripts are run in order. If the
// init script returns with a bad exit code, the Apache Spark container fails to
// launch and init scripts with later position are skipped. If enough containers
// fail, the entire cluster fails with a `GLOBAL_INIT_SCRIPT_FAILURE` error
// code.
type GlobalInitScriptsService interface {

	// Create init script.
	//
	// Creates a new global init script in this workspace.
	Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error)

	// Delete init script.
	//
	// Deletes a global init script.
	Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error

	// Get an init script.
	//
	// Gets all the details of a script, including its Base64-encoded contents.
	Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error)

	// Get init scripts.
	//
	// "Get a list of all global init scripts for this workspace. This returns
	// all properties for each script but **not** the script contents. To
	// retrieve the contents of a script, use the [get a global init
	// script](#operation/get-script) operation.
	//
	// Use ListAll() to get all GlobalInitScriptDetails instances
	List(ctx context.Context) (*ListGlobalInitScriptsResponse, error)

	// Update init script.
	//
	// Updates a global init script, specifying only the fields to change. All
	// fields are optional. Unspecified fields retain their current value.
	Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error
}

// Instance Pools API are used to create, edit, delete and list instance pools
// by using ready-to-use cloud instances which reduces a cluster start and
// auto-scaling times.
//
// Databricks pools reduce cluster start and auto-scaling times by maintaining a
// set of idle, ready-to-use instances. When a cluster is attached to a pool,
// cluster nodes are created using the pool’s idle instances. If the pool has
// no idle instances, the pool expands by allocating a new instance from the
// instance provider in order to accommodate the cluster’s request. When a
// cluster releases an instance, it returns to the pool and is free for another
// cluster to use. Only clusters attached to a pool can use that pool’s idle
// instances.
//
// You can specify a different pool for the driver node and worker nodes, or use
// the same pool for both.
//
// Databricks does not charge DBUs while instances are idle in the pool.
// Instance provider billing does apply. See pricing.
type InstancePoolsService interface {

	// Create a new instance pool.
	//
	// Creates a new instance pool using idle and ready-to-use cloud instances.
	Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error)

	// Delete an instance pool.
	//
	// Deletes the instance pool permanently. The idle instances in the pool are
	// terminated asynchronously.
	Delete(ctx context.Context, request DeleteInstancePool) error

	// Edit an existing instance pool.
	//
	// Modifies the configuration of an existing instance pool.
	Edit(ctx context.Context, request EditInstancePool) error

	// Get instance pool information.
	//
	// Retrieve the information for an instance pool based on its identifier.
	Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error)

	// List instance pool info.
	//
	// Gets a list of instance pools with their statistics.
	//
	// Use ListAll() to get all InstancePoolAndStats instances
	List(ctx context.Context) (*ListInstancePools, error)
}

// The Instance Profiles API allows admins to add, list, and remove instance
// profiles that users can launch clusters with. Regular users can list the
// instance profiles available to them. See [Secure access to S3 buckets] using
// instance profiles for more information.
//
// [Secure access to S3 buckets]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html
type InstanceProfilesService interface {

	// Register an instance profile.
	//
	// In the UI, you can select the instance profile when launching clusters.
	// This API is only available to admin users.
	Add(ctx context.Context, request AddInstanceProfile) error

	// Edit an instance profile.
	//
	// The only supported field to change is the optional IAM role ARN
	// associated with the instance profile. It is required to specify the IAM
	// role ARN if both of the following are true:
	//
	// * Your role name and instance profile name do not match. The name is the
	// part after the last slash in each ARN. * You want to use the instance
	// profile with [Databricks SQL Serverless].
	//
	// To understand where these fields are in the AWS console, see [Enable
	// serverless SQL warehouses].
	//
	// This API is only available to admin users.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	// [Enable serverless SQL warehouses]: https://docs.databricks.com/sql/admin/serverless.html
	Edit(ctx context.Context, request InstanceProfile) error

	// List available instance profiles.
	//
	// List the instance profiles that the calling user can use to launch a
	// cluster.
	//
	// This API is available to all users.
	//
	// Use ListAll() to get all InstanceProfile instances
	List(ctx context.Context) (*ListInstanceProfilesResponse, error)

	// Remove the instance profile.
	//
	// Remove the instance profile with the provided ARN. Existing clusters with
	// this instance profile will continue to function.
	//
	// This API is only accessible to admin users.
	Remove(ctx context.Context, request RemoveInstanceProfile) error
}

// The Libraries API allows you to install and uninstall libraries and get the
// status of libraries on a cluster.
//
// To make third-party or custom code available to notebooks and jobs running on
// your clusters, you can install a library. Libraries can be written in Python,
// Java, Scala, and R. You can upload Java, Scala, and Python libraries and
// point to external packages in PyPI, Maven, and CRAN repositories.
//
// Cluster libraries can be used by all notebooks running on a cluster. You can
// install a cluster library directly from a public repository such as PyPI or
// Maven, using a previously installed workspace library, or using an init
// script.
//
// When you install a library on a cluster, a notebook already attached to that
// cluster will not immediately see the new library. You must first detach and
// then reattach the notebook to the cluster.
//
// When you uninstall a library from a cluster, the library is removed only when
// you restart the cluster. Until you restart the cluster, the status of the
// uninstalled library appears as Uninstall pending restart.
type LibrariesService interface {

	// Get all statuses.
	//
	// Get the status of all libraries on all clusters. A status will be
	// available for all libraries installed on this cluster via the API or the
	// libraries UI as well as libraries set to be installed on all clusters via
	// the libraries UI.
	AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error)

	// Get status.
	//
	// Get the status of libraries on a cluster. A status will be available for
	// all libraries installed on this cluster via the API or the libraries UI
	// as well as libraries set to be installed on all clusters via the
	// libraries UI. The order of returned libraries will be as follows.
	//
	// 1. Libraries set to be installed on this cluster will be returned first.
	// Within this group, the final order will be order in which the libraries
	// were added to the cluster.
	//
	// 2. Libraries set to be installed on all clusters are returned next.
	// Within this group there is no order guarantee.
	//
	// 3. Libraries that were previously requested on this cluster or on all
	// clusters, but now marked for removal. Within this group there is no order
	// guarantee.
	ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterLibraryStatuses, error)

	// Add a library.
	//
	// Add libraries to be installed on a cluster. The installation is
	// asynchronous; it happens in the background after the completion of this
	// request.
	//
	// **Note**: The actual set of libraries to be installed on a cluster is the
	// union of the libraries specified via this method and the libraries set to
	// be installed on all clusters via the libraries UI.
	Install(ctx context.Context, request InstallLibraries) error

	// Uninstall libraries.
	//
	// Set libraries to be uninstalled on a cluster. The libraries won't be
	// uninstalled until the cluster is restarted. Uninstalling libraries that
	// are not installed on the cluster will have no impact but is not an error.
	Uninstall(ctx context.Context, request UninstallLibraries) error
}

// View available policy families. A policy family contains a policy definition
// providing best practices for configuring clusters for a particular use case.
//
// Databricks manages and provides policy families for several common cluster
// use cases. You cannot create, edit, or delete policy families.
//
// Policy families cannot be used directly to create clusters. Instead, you
// create cluster policies using a policy family. Cluster policies created using
// a policy family inherit the policy family's policy definition.
type PolicyFamiliesService interface {
	Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error)

	//
	// Use ListAll() to get all PolicyFamily instances, which will iterate over every result page.
	List(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error)
}
