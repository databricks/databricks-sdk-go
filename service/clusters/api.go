// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/retries"
)

func NewClusters(client *client.DatabricksClient) ClustersService {
	return &ClustersAPI{
		client: client,
	}
}

type ClustersAPI struct {
	client *client.DatabricksClient
}

// Public version of editClusterOwner, allowing admins to change cluster owner
func (a *ClustersAPI) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	path := "/api/2.0/clusters/change-owner"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Creates a new Spark cluster. This method will acquire new instances from the
// cloud provider if necessary. This method is asynchronous; the returned
// “cluster_id“ can be used to poll the cluster status. When this method
// returns, the cluster will be in a “PENDING“ state. The cluster will be
// usable once it enters a “RUNNING“ state. Note: Databricks may not be able
// to acquire some of the requested nodes, due to cloud provider limitations
// (account limits, spot price, ...) or transient network issues. If Databricks
// acquires at least 85% of the requested on-demand nodes, cluster creation will
// succeed. Otherwise the cluster will terminate with an informative error
// message. An example request: .. code:: { "cluster_name": "my-cluster",
// "spark_version": "2.0.x-scala2.10", "node_type_id": "r3.xlarge",
// "spark_conf": { "spark.speculation": true }, "aws_attributes": {
// "availability": "SPOT", "zone_id": "us-west-2a" }, "num_workers": 25 } See
// below as an example for an autoscaling cluster. Note that this cluster will
// start with `2` nodes, the minimum. .. code:: { "cluster_name":
// "autoscaling-cluster", "spark_version": "2.0.x-scala2.10", "node_type_id":
// "r3.xlarge", "autoscale" : { "min_workers": 2, "max_workers": 50 } }
func (a *ClustersAPI) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	var createClusterResponse CreateClusterResponse
	path := "/api/2.0/clusters/create"
	err := a.client.Post(ctx, path, request, &createClusterResponse)
	return &createClusterResponse, err
}

// Create and wait to reach RUNNING state
func (a *ClustersAPI) CreateAndWait(ctx context.Context, createCluster CreateCluster, timeout ...time.Duration) (*ClusterInfo, error) {
	createClusterResponse, err := a.Create(ctx, createCluster)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: createClusterResponse.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Terminates a Spark cluster given its id. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// “TERMINATED“ state. If the cluster is already in a “TERMINATING“ or
// “TERMINATED“ state, nothing will happen. An example request: .. code:: {
// "cluster_id": "1202-211320-brick1" }
func (a *ClustersAPI) Delete(ctx context.Context, request DeleteCluster) error {
	path := "/api/2.0/clusters/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete and wait to reach TERMINATED state
func (a *ClustersAPI) DeleteAndWait(ctx context.Context, deleteCluster DeleteCluster, timeout ...time.Duration) (*ClusterInfo, error) {
	err := a.Delete(ctx, deleteCluster)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: deleteCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateTerminated: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateTerminated, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Terminates a Spark cluster given its id. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// “TERMINATED“ state. If the cluster is already in a “TERMINATING“ or
// “TERMINATED“ state, nothing will happen. An example request: .. code:: {
// "cluster_id": "1202-211320-brick1" }
func (a *ClustersAPI) DeleteByClusterId(ctx context.Context, clusterId string) error {
	return a.Delete(ctx, DeleteCluster{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) DeleteByClusterIdAndWait(ctx context.Context, clusterId string, timeout ...time.Duration) (*ClusterInfo, error) {
	return a.DeleteAndWait(ctx, DeleteCluster{
		ClusterId: clusterId,
	}, timeout...)
}

// Edits the configuration of a cluster to match the provided attributes and
// size. A cluster can be edited if it is in a “RUNNING“ or “TERMINATED“
// state. If a cluster is edited while in a “RUNNING“ state, it will be
// restarted so that the new attributes can take effect. If a cluster is edited
// while in a “TERMINATED“ state, it will remain “TERMINATED“. The next time
// it is started using the “clusters/start“ API, the new attributes will take
// effect. An attempt to edit a cluster in any other state will be rejected with
// an “INVALID_STATE“ error code. Clusters created by the Databricks Jobs
// service cannot be edited. An example request: .. code:: { "cluster_id":
// "1202-211320-brick1", "num_workers": 10, "spark_version": "3.3.x-scala2.11",
// "node_type_id": "i3.2xlarge" }
func (a *ClustersAPI) Edit(ctx context.Context, request EditCluster) error {
	path := "/api/2.0/clusters/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Edit and wait to reach RUNNING state
func (a *ClustersAPI) EditAndWait(ctx context.Context, editCluster EditCluster, timeout ...time.Duration) (*ClusterInfo, error) {
	err := a.Edit(ctx, editCluster)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: editCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// parameters necessary to request the next page of events. An example request:
// “/clusters/events?cluster_id=1202-211320-brick1“ An example response: {
// "events": [ { "cluster_id": "1202-211320-brick1", "timestamp": 1509572145487,
// "event_type": "RESTARTING", "event_details": { "username": "admin" } }, ... {
// "cluster_id": "1202-211320-brick1", "timestamp": 1509505807923, "event_type":
// "TERMINATING", "event_details": { "termination_reason": { "code":
// "USER_REQUEST", "parameters": [ "username": "admin" ] } } ], "next_page": {
// "cluster_id": "1202-211320-brick1", "end_time": 1509572145487, "order":
// "DESC", "offset": 50 }, "total_count": 303 } Example request to retrieve the
// next page of events
// “/clusters/events?cluster_id=1202-211320-brick1&end_time=1509572145487&order=DESC&offset=50“
func (a *ClustersAPI) Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {
	var getEventsResponse GetEventsResponse
	path := "/api/2.0/clusters/events"
	err := a.client.Post(ctx, path, request, &getEventsResponse)
	return &getEventsResponse, err
}

// Retrieves the information for a cluster given its identifier. Clusters can be
// described while they are running, or up to 60 days after they are terminated.
// An example request: “/clusters/get?cluster_id=1202-211320-brick1“
func (a *ClustersAPI) Get(ctx context.Context, request GetRequest) (*ClusterInfo, error) {
	var clusterInfo ClusterInfo
	path := "/api/2.0/clusters/get"
	err := a.client.Get(ctx, path, request, &clusterInfo)
	return &clusterInfo, err
}

// Get and wait to reach RUNNING state
func (a *ClustersAPI) GetAndWait(ctx context.Context, getRequest GetRequest, timeout ...time.Duration) (*ClusterInfo, error) {
	clusterInfo, err := a.Get(ctx, getRequest)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: clusterInfo.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Retrieves the information for a cluster given its identifier. Clusters can be
// described while they are running, or up to 60 days after they are terminated.
// An example request: “/clusters/get?cluster_id=1202-211320-brick1“
func (a *ClustersAPI) GetByClusterId(ctx context.Context, clusterId string) (*ClusterInfo, error) {
	return a.Get(ctx, GetRequest{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) GetByClusterIdAndWait(ctx context.Context, clusterId string, timeout ...time.Duration) (*ClusterInfo, error) {
	return a.GetAndWait(ctx, GetRequest{
		ClusterId: clusterId,
	}, timeout...)
}

// Returns information about all pinned clusters, currently active clusters, up
// to 70 of the most recently terminated interactive clusters in the past 7
// days, and up to 30 of the most recently terminated job clusters in the past 7
// days. For example, if there is 1 pinned cluster, 4 active clusters, 45
// terminated interactive clusters in the past 7 days, and 50 terminated job
// clusters in the past 7 days, then this API returns the 1 pinned cluster, 4
// active clusters, all 45 terminated interactive clusters, and the 30 most
// recently terminated job clusters.
func (a *ClustersAPI) List(ctx context.Context, request ListRequest) (*ListClustersResponse, error) {
	var listClustersResponse ListClustersResponse
	path := "/api/2.0/clusters/list"
	err := a.client.Get(ctx, path, request, &listClustersResponse)
	return &listClustersResponse, err
}

// Returns information about all pinned clusters, currently active clusters, up
// to 70 of the most recently terminated interactive clusters in the past 7
// days, and up to 30 of the most recently terminated job clusters in the past 7
// days. For example, if there is 1 pinned cluster, 4 active clusters, 45
// terminated interactive clusters in the past 7 days, and 50 terminated job
// clusters in the past 7 days, then this API returns the 1 pinned cluster, 4
// active clusters, all 45 terminated interactive clusters, and the 30 most
// recently terminated job clusters.
func (a *ClustersAPI) ListByCanUseClient(ctx context.Context, canUseClient string) (*ListClustersResponse, error) {
	return a.List(ctx, ListRequest{
		CanUseClient: canUseClient,
	})
}

// Returns a list of supported Spark node types. These node types can be used to
// launch a cluster.
func (a *ClustersAPI) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	var listNodeTypesResponse ListNodeTypesResponse
	path := "/api/2.0/clusters/list-node-types"
	err := a.client.Get(ctx, path, nil, &listNodeTypesResponse)
	return &listNodeTypesResponse, err
}

// Returns a list of availability zones where clusters can be created in (ex:
// us-west-2a). These zones can be used to launch a cluster.
func (a *ClustersAPI) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	var listAvailableZonesResponse ListAvailableZonesResponse
	path := "/api/2.0/clusters/list-zones"
	err := a.client.Get(ctx, path, nil, &listAvailableZonesResponse)
	return &listAvailableZonesResponse, err
}

// Permanently deletes a Spark cluster. This cluster is terminated and resources
// are asynchronously removed. In addition, users will no longer see permanently
// deleted clusters in the cluster list, and API users can no longer perform any
// action on permanently deleted clusters. An example request: .. code:: {
// "cluster_id": "1202-211320-brick1" }
func (a *ClustersAPI) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	path := "/api/2.0/clusters/permanent-delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Permanently deletes a Spark cluster. This cluster is terminated and resources
// are asynchronously removed. In addition, users will no longer see permanently
// deleted clusters in the cluster list, and API users can no longer perform any
// action on permanently deleted clusters. An example request: .. code:: {
// "cluster_id": "1202-211320-brick1" }
func (a *ClustersAPI) PermanentDeleteByClusterId(ctx context.Context, clusterId string) error {
	return a.PermanentDelete(ctx, PermanentDeleteCluster{
		ClusterId: clusterId,
	})
}

// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins. An example request:
// “/clusters/pin?cluster_id=1202-211320-brick1“
func (a *ClustersAPI) Pin(ctx context.Context, request PinCluster) error {
	path := "/api/2.0/clusters/pin"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins. An example request:
// “/clusters/pin?cluster_id=1202-211320-brick1“
func (a *ClustersAPI) PinByClusterId(ctx context.Context, clusterId string) error {
	return a.Pin(ctx, PinCluster{
		ClusterId: clusterId,
	})
}

// Resizes a cluster to have a desired number of workers. This will fail unless
// the cluster is in a “RUNNING“ state. An example request: .. code:: {
// "cluster_id": "1202-211320-brick1", "num_workers": 30 }
func (a *ClustersAPI) Resize(ctx context.Context, request ResizeCluster) error {
	path := "/api/2.0/clusters/resize"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Resize and wait to reach RUNNING state
func (a *ClustersAPI) ResizeAndWait(ctx context.Context, resizeCluster ResizeCluster, timeout ...time.Duration) (*ClusterInfo, error) {
	err := a.Resize(ctx, resizeCluster)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: resizeCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Restarts a Spark cluster given its id. If the cluster is not currently in a
// “RUNNING“ state, nothing will happen. An example request: .. code:: {
// "cluster_id": "1202-211320-brick1" }
func (a *ClustersAPI) Restart(ctx context.Context, request RestartCluster) error {
	path := "/api/2.0/clusters/restart"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Restart and wait to reach RUNNING state
func (a *ClustersAPI) RestartAndWait(ctx context.Context, restartCluster RestartCluster, timeout ...time.Duration) (*ClusterInfo, error) {
	err := a.Restart(ctx, restartCluster)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: restartCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Returns the list of available Spark versions. These versions can be used to
// launch a cluster.
func (a *ClustersAPI) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	var getSparkVersionsResponse GetSparkVersionsResponse
	path := "/api/2.0/clusters/spark-versions"
	err := a.client.Get(ctx, path, nil, &getSparkVersionsResponse)
	return &getSparkVersionsResponse, err
}

// Starts a terminated Spark cluster given its id. This works similar to
// `createCluster` except: - The previous cluster id and attributes are
// preserved. - The cluster starts with the last specified cluster size. - If
// the previous cluster was an autoscaling cluster, the current cluster starts
// with the minimum number of nodes. - If the cluster is not currently in a
// “TERMINATED“ state, nothing will happen. - Clusters launched to run a job
// cannot be started. An example request: .. code:: { "cluster_id":
// "1202-211320-brick1" }
func (a *ClustersAPI) Start(ctx context.Context, request StartCluster) error {
	path := "/api/2.0/clusters/start"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Start and wait to reach RUNNING state
func (a *ClustersAPI) StartAndWait(ctx context.Context, startCluster StartCluster, timeout ...time.Duration) (*ClusterInfo, error) {
	err := a.Start(ctx, startCluster)
	if err != nil {
		return nil, err
	}
	if len(timeout) == 0 {
		timeout = []time.Duration{20 * time.Minute}
	}
	return retries.Poll[ClusterInfo](ctx, timeout[0], func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: startCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Starts a terminated Spark cluster given its id. This works similar to
// `createCluster` except: - The previous cluster id and attributes are
// preserved. - The cluster starts with the last specified cluster size. - If
// the previous cluster was an autoscaling cluster, the current cluster starts
// with the minimum number of nodes. - If the cluster is not currently in a
// “TERMINATED“ state, nothing will happen. - Clusters launched to run a job
// cannot be started. An example request: .. code:: { "cluster_id":
// "1202-211320-brick1" }
func (a *ClustersAPI) StartByClusterId(ctx context.Context, clusterId string) error {
	return a.Start(ctx, StartCluster{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) StartByClusterIdAndWait(ctx context.Context, clusterId string, timeout ...time.Duration) (*ClusterInfo, error) {
	return a.StartAndWait(ctx, StartCluster{
		ClusterId: clusterId,
	}, timeout...)
}

// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins. An example request:
// “/clusters/unpin?cluster_id=1202-211320-brick1“
func (a *ClustersAPI) Unpin(ctx context.Context, request UnpinCluster) error {
	path := "/api/2.0/clusters/unpin"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins. An example request:
// “/clusters/unpin?cluster_id=1202-211320-brick1“
func (a *ClustersAPI) UnpinByClusterId(ctx context.Context, clusterId string) error {
	return a.Unpin(ctx, UnpinCluster{
		ClusterId: clusterId,
	})
}
