// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

import (
	"context"
	"time"
)

// <needs content added>
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type ClustersService interface {

	// Change cluster owner
	//
	// Change the owner of the cluster. You must be an admin to perform this
	// operation.
	ChangeOwner(ctx context.Context, request ChangeClusterOwner) error

	// Create new cluster
	//
	// Creates a new Spark cluster. This method will acquire new instances from
	// the cloud provider if necessary. This method is asynchronous; the
	// returned ``cluster_id`` can be used to poll the cluster status. When this
	// method returns, the cluster will be in\na ``PENDING`` state. The cluster
	// will be usable once it enters a ``RUNNING`` state.
	//
	// Note: Databricks may not be able to acquire some of the requested nodes,
	// due to cloud provider limitations (account limits, spot price, etc.) or
	// transient network issues.
	//
	// If Databricks acquires at least 85% of the requested on-demand nodes,
	// cluster creation will succeed. Otherwise the cluster will terminate with
	// an informative error message.
	Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error)

	// CreateAndWait calls Create() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	CreateAndWait(ctx context.Context, request CreateCluster, timeout ...time.Duration) (*ClusterInfo, error)

	// Terminate cluster
	//
	// Terminates the Spark cluster with the specified ID. The cluster is
	// removed asynchronously. Once the termination has completed, the cluster
	// will be in a ``TERMINATED`` state. If the cluster is already in a
	// ``TERMINATING`` or ``TERMINATED`` state, nothing will happen.
	Delete(ctx context.Context, request DeleteCluster) error

	// DeleteAndWait calls Delete() and waits to reach TERMINATED state
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteAndWait(ctx context.Context, request DeleteCluster, timeout ...time.Duration) (*ClusterInfo, error)
	// DeleteByClusterId calls Delete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByClusterId(ctx context.Context, clusterId string) error

	// DeleteByClusterIdAndWait calls DeleteByClusterId and waits until ClusterInfo is in desired state.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteByClusterIdAndWait(ctx context.Context, clusterId string, timeout ...time.Duration) (*ClusterInfo, error)

	// Update cluster configuration
	//
	// Updates the configuration of a cluster to match the provided attributes
	// and size. A cluster can be updated if it is in a ``RUNNING`` or
	// ``TERMINATED`` state.
	//
	// If a cluster is updated while in a ``RUNNING`` state, it will be
	// restarted so that the new attributes can take effect.
	//
	// If a cluster is updated while in a ``TERMINATED`` state, it will remain
	// ``TERMINATED``. The next time it is started using the ``clusters/start``
	// API, the new attributes will take effect. Any attempt to update a cluster
	// in any other state will be rejected with an ``INVALID_STATE`` error code.
	//
	// Clusters created by the Databricks Jobs service cannot be edited.
	Edit(ctx context.Context, request EditCluster) error

	// EditAndWait calls Edit() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	EditAndWait(ctx context.Context, request EditCluster, timeout ...time.Duration) (*ClusterInfo, error)

	// List cluster activity events
	//
	// Retrieves a list of events about the activity of a cluster. This API is
	// paginated. If there are more events to read, the response includes all
	// the nparameters necessary to request the next page of events.
	Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error)

<<<<<<< HEAD
	// Get cluster info
	//
	// "Retrieves the information for a cluster given its identifier. Clusters
=======
	// EventsAll calls Events() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	EventsAll(ctx context.Context, request GetEvents) ([]ClusterEvent, error)

	// Retrieves the information for a cluster given its identifier. Clusters
>>>>>>> 499b182 (regenerate)
	// can be described while they are running, or up to 60 days after they are
	// terminated.
	Get(ctx context.Context, request GetRequest) (*ClusterInfo, error)

	// GetAndWait calls Get() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	GetAndWait(ctx context.Context, request GetRequest, timeout ...time.Duration) (*ClusterInfo, error)
	// GetByClusterId calls Get, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByClusterId(ctx context.Context, clusterId string) (*ClusterInfo, error)

	// GetByClusterIdAndWait calls GetByClusterId and waits until ClusterInfo is in desired state.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByClusterIdAndWait(ctx context.Context, clusterId string, timeout ...time.Duration) (*ClusterInfo, error)

	// List all clusters
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
	List(ctx context.Context, request ListRequest) (*ListClustersResponse, error)

	// ListByCanUseClient calls List, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListByCanUseClient(ctx context.Context, canUseClient string) (*ListClustersResponse, error)
	// ListAll calls List() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListRequest) ([]ClusterInfo, error)

	// ClusterInfoClusterNameToClusterIdMap retrieves a mapping to access ID by name
	//
	// This method is generated by Databricks SDK Code Generator.
	ClusterInfoClusterNameToClusterIdMap(ctx context.Context, request ListRequest) (map[string]string, error)

	// GetClusterInfoByClusterName retrieves ClusterInfo by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetClusterInfoByClusterName(ctx context.Context, name string) (*ClusterInfo, error)

	// List node types
	//
	// Returns a list of supported Spark node types. These node types can be
	// used to launch a cluster.
	ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error)

	// List availability zones
	//
	// Returns a list of availability zones where clusters can be created in
	// (For example, us-west-2a). These zones can be used to launch a cluster.
	ListZones(ctx context.Context) (*ListAvailableZonesResponse, error)

	// Permanently delete cluster
	//
	// Permanently deletes a Spark cluster. This cluster is terminated and
	// resources are asynchronously removed.
	//
	// In addition, users will no longer see permanently deleted clusters in the
	// cluster list, and API users can no longer perform any action on
	// permanently deleted clusters.
	PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error

	// PermanentDeleteByClusterId calls PermanentDelete, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	PermanentDeleteByClusterId(ctx context.Context, clusterId string) error

	// Pin cluster
	//
	// Pinning a cluster ensures that the cluster will always be returned by the
	// ListClusters API. Pinning a cluster that is already pinned will have no
	// effect. This API can only be called by workspace admins.
	Pin(ctx context.Context, request PinCluster) error

	// PinByClusterId calls Pin, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	PinByClusterId(ctx context.Context, clusterId string) error

	// Resize cluster
	//
	// Resizes a cluster to have a desired number of workers. This will fail
	// unless the cluster is in a `RUNNING` state.
	Resize(ctx context.Context, request ResizeCluster) error

	// ResizeAndWait calls Resize() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	ResizeAndWait(ctx context.Context, request ResizeCluster, timeout ...time.Duration) (*ClusterInfo, error)

	// Restart cluster
	//
	// Restarts a Spark cluster with the supplied ID. If the cluster is not
	// currently in a `RUNNING` state, nothing will happen.
	Restart(ctx context.Context, request RestartCluster) error

	// RestartAndWait calls Restart() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	RestartAndWait(ctx context.Context, request RestartCluster, timeout ...time.Duration) (*ClusterInfo, error)

	// List available Spark versions
	//
	// Returns the list of available Spark versions. These versions can be used
	// to launch a cluster.
	SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error)

	// Start terminated cluster
	//
	// Starts a terminated Spark cluster with the supplied ID. This works
	// similar to `createCluster` except:
	//
	// * The previous cluster id and attributes are preserved. * The cluster
	// starts with the last specified cluster size. * If the previous cluster
	// was an autoscaling cluster, the current cluster starts with the minimum
	// number of nodes. * If the cluster is not currently in a ``TERMINATED``
	// state, nothing will happen. * Clusters launched to run a job cannot be
	// started.
	Start(ctx context.Context, request StartCluster) error

	// StartAndWait calls Start() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	StartAndWait(ctx context.Context, request StartCluster, timeout ...time.Duration) (*ClusterInfo, error)
	// StartByClusterId calls Start, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	StartByClusterId(ctx context.Context, clusterId string) error

	// StartByClusterIdAndWait calls StartByClusterId and waits until ClusterInfo is in desired state.
	//
	// This method is generated by Databricks SDK Code Generator.
	StartByClusterIdAndWait(ctx context.Context, clusterId string, timeout ...time.Duration) (*ClusterInfo, error)

	// Unpin cluster
	//
	// Unpinning a cluster will allow the cluster to eventually be removed from
	// the ListClusters API. Unpinning a cluster that is not pinned will have no
	// effect. This API can only be called by workspace admins.
	Unpin(ctx context.Context, request UnpinCluster) error

	// UnpinByClusterId calls Unpin, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	UnpinByClusterId(ctx context.Context, clusterId string) error

	// GetOrCreateRunningCluster creates an autoterminating cluster if it doesn't exist
	GetOrCreateRunningCluster(ctx context.Context, name string, custom ...CreateCluster) (c *ClusterInfo, err error)
}
