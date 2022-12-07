// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

import (
	"context"
)

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
	// method returns, the cluster will be in\na `PENDING` state. The cluster
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
	Get(ctx context.Context, request Get) (*ClusterInfo, error)

	// List all clusters.
	//
	// Returns information about all pinned clusters, currently active clusters,
	// up to 70 of the most recently terminated interactive clusters in the past
	// 7 days, and up to 30 of the most recently terminated job clusters in the
	// past 7 days.
	//
	// For example, if there is 1 pinned cluster, 4 active clusters, 45
	// terminated interactive clusters in the past 7 days, and 50 terminated job
	// clusters\nin the past 7 days, then this API returns the 1 pinned cluster,
	// 4 active clusters, all 45 terminated interactive clusters, and the 30
	// most recently terminated job clusters.
	//
	// Use ListAll() to get all ClusterInfo instances
	List(ctx context.Context, request List) (*ListClustersResponse, error)

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
