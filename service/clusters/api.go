// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

func NewClusters(client *client.DatabricksClient) *ClustersAPI {
	return &ClustersAPI{
		ClustersService: &clustersAPI{
			client: client,
		},
	}
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
type ClustersAPI struct {
	// ClustersService contains low-level REST API interface.
	ClustersService
}

// Change cluster owner
//
// Change the owner of the cluster. You must be an admin to perform this
// operation.
func (a *ClustersAPI) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	return a.ClustersService.ChangeOwner(ctx, request)
}

// Create new cluster
//
// Creates a new Spark cluster. This method will acquire new instances from the
// cloud provider if necessary. This method is asynchronous; the returned
// “cluster_id“ can be used to poll the cluster status. When this method
// returns, the cluster will be in\na “PENDING“ state. The cluster will be
// usable once it enters a “RUNNING“ state.
//
// Note: Databricks may not be able to acquire some of the requested nodes, due
// to cloud provider limitations (account limits, spot price, etc.) or transient
// network issues.
//
// If Databricks acquires at least 85% of the requested on-demand nodes, cluster
// creation will succeed. Otherwise the cluster will terminate with an
// informative error message.
func (a *ClustersAPI) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	return a.ClustersService.Create(ctx, request)
}

// Calls [ClustersAPI.Create] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) CreateAndWait(ctx context.Context, createCluster CreateCluster, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	createClusterResponse, err := a.Create(ctx, createCluster)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: createClusterResponse.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError, ClusterInfoStateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Terminate cluster
//
// Terminates the Spark cluster with the specified ID. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// “TERMINATED“ state. If the cluster is already in a “TERMINATING“ or
// “TERMINATED“ state, nothing will happen.
func (a *ClustersAPI) Delete(ctx context.Context, request DeleteCluster) error {
	return a.ClustersService.Delete(ctx, request)
}

// Calls [ClustersAPI.Delete] and waits to reach TERMINATED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) DeleteAndWait(ctx context.Context, deleteCluster DeleteCluster, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Delete(ctx, deleteCluster)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: deleteCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
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

// Terminate cluster
//
// Terminates the Spark cluster with the specified ID. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// “TERMINATED“ state. If the cluster is already in a “TERMINATING“ or
// “TERMINATED“ state, nothing will happen.
func (a *ClustersAPI) DeleteByClusterId(ctx context.Context, clusterId string) error {
	return a.Delete(ctx, DeleteCluster{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) DeleteByClusterIdAndWait(ctx context.Context, clusterId string, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	return a.DeleteAndWait(ctx, DeleteCluster{
		ClusterId: clusterId,
	}, options...)
}

// Update cluster configuration
//
// Updates the configuration of a cluster to match the provided attributes and
// size. A cluster can be updated if it is in a “RUNNING“ or “TERMINATED“
// state.
//
// If a cluster is updated while in a “RUNNING“ state, it will be restarted so
// that the new attributes can take effect.
//
// If a cluster is updated while in a “TERMINATED“ state, it will remain
// “TERMINATED“. The next time it is started using the “clusters/start“ API,
// the new attributes will take effect. Any attempt to update a cluster in any
// other state will be rejected with an “INVALID_STATE“ error code.
//
// Clusters created by the Databricks Jobs service cannot be edited.
func (a *ClustersAPI) Edit(ctx context.Context, request EditCluster) error {
	return a.ClustersService.Edit(ctx, request)
}

// Calls [ClustersAPI.Edit] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) EditAndWait(ctx context.Context, editCluster EditCluster, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Edit(ctx, editCluster)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: editCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError, ClusterInfoStateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// List cluster activity events
//
// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// nparameters necessary to request the next page of events.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) EventsAll(ctx context.Context, request GetEvents) ([]ClusterEvent, error) {
	var results []ClusterEvent
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.Events(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Events) == 0 {
			break
		}
		for _, v := range response.Events {
			results = append(results, v)
		}
		request.Offset += int64(len(response.Events))
	}
	return results, nil
}

// Get cluster info
//
// "Retrieves the information for a cluster given its identifier. Clusters can
// be described while they are running, or up to 60 days after they are
// terminated.
func (a *ClustersAPI) Get(ctx context.Context, request GetRequest) (*ClusterInfo, error) {
	return a.ClustersService.Get(ctx, request)
}

// Calls [ClustersAPI.Get] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) GetAndWait(ctx context.Context, getRequest GetRequest, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	clusterInfo, err := a.Get(ctx, getRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: clusterInfo.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError, ClusterInfoStateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get cluster info
//
// "Retrieves the information for a cluster given its identifier. Clusters can
// be described while they are running, or up to 60 days after they are
// terminated.
func (a *ClustersAPI) GetByClusterId(ctx context.Context, clusterId string) (*ClusterInfo, error) {
	return a.Get(ctx, GetRequest{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) GetByClusterIdAndWait(ctx context.Context, clusterId string, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	return a.GetAndWait(ctx, GetRequest{
		ClusterId: clusterId,
	}, options...)
}

// List all clusters
//
// Returns information about all pinned clusters, currently active clusters, up
// to 70 of the most recently terminated interactive clusters in the past 7
// days, and up to 30 of the most recently terminated job clusters in the past 7
// days.
//
// For example, if there is 1 pinned cluster, 4 active clusters, 45 terminated
// interactive clusters in the past 7 days, and 50 terminated job clusters\nin
// the past 7 days, then this API returns the 1 pinned cluster, 4 active
// clusters, all 45 terminated interactive clusters, and the 30 most recently
// terminated job clusters.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) ListAll(ctx context.Context, request ListRequest) ([]ClusterInfo, error) {
	response, err := a.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Clusters, nil
}

// List all clusters
//
// Returns information about all pinned clusters, currently active clusters, up
// to 70 of the most recently terminated interactive clusters in the past 7
// days, and up to 30 of the most recently terminated job clusters in the past 7
// days.
//
// For example, if there is 1 pinned cluster, 4 active clusters, 45 terminated
// interactive clusters in the past 7 days, and 50 terminated job clusters\nin
// the past 7 days, then this API returns the 1 pinned cluster, 4 active
// clusters, all 45 terminated interactive clusters, and the 30 most recently
// terminated job clusters.
func (a *ClustersAPI) ListByCanUseClient(ctx context.Context, canUseClient string) (*ListClustersResponse, error) {
	return a.List(ctx, ListRequest{
		CanUseClient: canUseClient,
	})
}

// List node types
//
// Returns a list of supported Spark node types. These node types can be used to
// launch a cluster.
func (a *ClustersAPI) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	return a.ClustersService.ListNodeTypes(ctx)
}

// List availability zones
//
// Returns a list of availability zones where clusters can be created in (For
// example, us-west-2a). These zones can be used to launch a cluster.
func (a *ClustersAPI) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	return a.ClustersService.ListZones(ctx)
}

// Permanently delete cluster
//
// Permanently deletes a Spark cluster. This cluster is terminated and resources
// are asynchronously removed.
//
// In addition, users will no longer see permanently deleted clusters in the
// cluster list, and API users can no longer perform any action on permanently
// deleted clusters.
func (a *ClustersAPI) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	return a.ClustersService.PermanentDelete(ctx, request)
}

// Permanently delete cluster
//
// Permanently deletes a Spark cluster. This cluster is terminated and resources
// are asynchronously removed.
//
// In addition, users will no longer see permanently deleted clusters in the
// cluster list, and API users can no longer perform any action on permanently
// deleted clusters.
func (a *ClustersAPI) PermanentDeleteByClusterId(ctx context.Context, clusterId string) error {
	return a.PermanentDelete(ctx, PermanentDeleteCluster{
		ClusterId: clusterId,
	})
}

// Pin cluster
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) Pin(ctx context.Context, request PinCluster) error {
	return a.ClustersService.Pin(ctx, request)
}

// Pin cluster
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) PinByClusterId(ctx context.Context, clusterId string) error {
	return a.Pin(ctx, PinCluster{
		ClusterId: clusterId,
	})
}

// Resize cluster
//
// Resizes a cluster to have a desired number of workers. This will fail unless
// the cluster is in a `RUNNING` state.
func (a *ClustersAPI) Resize(ctx context.Context, request ResizeCluster) error {
	return a.ClustersService.Resize(ctx, request)
}

// Calls [ClustersAPI.Resize] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) ResizeAndWait(ctx context.Context, resizeCluster ResizeCluster, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Resize(ctx, resizeCluster)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: resizeCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError, ClusterInfoStateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Restart cluster
//
// Restarts a Spark cluster with the supplied ID. If the cluster is not
// currently in a `RUNNING` state, nothing will happen.
func (a *ClustersAPI) Restart(ctx context.Context, request RestartCluster) error {
	return a.ClustersService.Restart(ctx, request)
}

// Calls [ClustersAPI.Restart] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) RestartAndWait(ctx context.Context, restartCluster RestartCluster, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Restart(ctx, restartCluster)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: restartCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError, ClusterInfoStateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// List available Spark versions
//
// Returns the list of available Spark versions. These versions can be used to
// launch a cluster.
func (a *ClustersAPI) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	return a.ClustersService.SparkVersions(ctx)
}

// Start terminated cluster
//
// Starts a terminated Spark cluster with the supplied ID. This works similar to
// `createCluster` except:
//
// * The previous cluster id and attributes are preserved. * The cluster starts
// with the last specified cluster size. * If the previous cluster was an
// autoscaling cluster, the current cluster starts with the minimum number of
// nodes. * If the cluster is not currently in a “TERMINATED“ state, nothing
// will happen. * Clusters launched to run a job cannot be started.
func (a *ClustersAPI) Start(ctx context.Context, request StartCluster) error {
	return a.ClustersService.Start(ctx, request)
}

// Calls [ClustersAPI.Start] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterInfo](60*time.Minute) functional option.
func (a *ClustersAPI) StartAndWait(ctx context.Context, startCluster StartCluster, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Start(ctx, startCluster)
	if err != nil {
		return nil, err
	}
	i := retries.Info[ClusterInfo]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[ClusterInfo](ctx, i.Timeout, func() (*ClusterInfo, *retries.Err) {
		clusterInfo, err := a.Get(ctx, GetRequest{
			ClusterId: startCluster.ClusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[ClusterInfo]{
				Info:    *clusterInfo,
				Timeout: i.Timeout,
			})
		}
		status := clusterInfo.State
		statusMessage := clusterInfo.StateMessage
		switch status {
		case ClusterInfoStateRunning: // target state
			return clusterInfo, nil
		case ClusterInfoStateError, ClusterInfoStateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ClusterInfoStateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Start terminated cluster
//
// Starts a terminated Spark cluster with the supplied ID. This works similar to
// `createCluster` except:
//
// * The previous cluster id and attributes are preserved. * The cluster starts
// with the last specified cluster size. * If the previous cluster was an
// autoscaling cluster, the current cluster starts with the minimum number of
// nodes. * If the cluster is not currently in a “TERMINATED“ state, nothing
// will happen. * Clusters launched to run a job cannot be started.
func (a *ClustersAPI) StartByClusterId(ctx context.Context, clusterId string) error {
	return a.Start(ctx, StartCluster{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) StartByClusterIdAndWait(ctx context.Context, clusterId string, options ...retries.Option[ClusterInfo]) (*ClusterInfo, error) {
	return a.StartAndWait(ctx, StartCluster{
		ClusterId: clusterId,
	}, options...)
}

// Unpin cluster
//
// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins.
func (a *ClustersAPI) Unpin(ctx context.Context, request UnpinCluster) error {
	return a.ClustersService.Unpin(ctx, request)
}

// Unpin cluster
//
// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins.
func (a *ClustersAPI) UnpinByClusterId(ctx context.Context, clusterId string) error {
	return a.Unpin(ctx, UnpinCluster{
		ClusterId: clusterId,
	})
}

// unexported type that holds implementations of just Clusters API methods
type clustersAPI struct {
	client *client.DatabricksClient
}

func (a *clustersAPI) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	path := "/api/2.0/clusters/change-owner"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	var createClusterResponse CreateClusterResponse
	path := "/api/2.0/clusters/create"
	err := a.client.Post(ctx, path, request, &createClusterResponse)
	return &createClusterResponse, err
}

func (a *clustersAPI) Delete(ctx context.Context, request DeleteCluster) error {
	path := "/api/2.0/clusters/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Edit(ctx context.Context, request EditCluster) error {
	path := "/api/2.0/clusters/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {
	var getEventsResponse GetEventsResponse
	path := "/api/2.0/clusters/events"
	err := a.client.Post(ctx, path, request, &getEventsResponse)
	return &getEventsResponse, err
}

func (a *clustersAPI) Get(ctx context.Context, request GetRequest) (*ClusterInfo, error) {
	var clusterInfo ClusterInfo
	path := "/api/2.0/clusters/get"
	err := a.client.Get(ctx, path, request, &clusterInfo)
	return &clusterInfo, err
}

func (a *clustersAPI) List(ctx context.Context, request ListRequest) (*ListClustersResponse, error) {
	var listClustersResponse ListClustersResponse
	path := "/api/2.0/clusters/list"
	err := a.client.Get(ctx, path, request, &listClustersResponse)
	return &listClustersResponse, err
}

func (a *clustersAPI) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	var listNodeTypesResponse ListNodeTypesResponse
	path := "/api/2.0/clusters/list-node-types"
	err := a.client.Get(ctx, path, nil, &listNodeTypesResponse)
	return &listNodeTypesResponse, err
}

func (a *clustersAPI) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	var listAvailableZonesResponse ListAvailableZonesResponse
	path := "/api/2.0/clusters/list-zones"
	err := a.client.Get(ctx, path, nil, &listAvailableZonesResponse)
	return &listAvailableZonesResponse, err
}

func (a *clustersAPI) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	path := "/api/2.0/clusters/permanent-delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Pin(ctx context.Context, request PinCluster) error {
	path := "/api/2.0/clusters/pin"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Resize(ctx context.Context, request ResizeCluster) error {
	path := "/api/2.0/clusters/resize"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Restart(ctx context.Context, request RestartCluster) error {
	path := "/api/2.0/clusters/restart"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	var getSparkVersionsResponse GetSparkVersionsResponse
	path := "/api/2.0/clusters/spark-versions"
	err := a.client.Get(ctx, path, nil, &getSparkVersionsResponse)
	return &getSparkVersionsResponse, err
}

func (a *clustersAPI) Start(ctx context.Context, request StartCluster) error {
	path := "/api/2.0/clusters/start"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *clustersAPI) Unpin(ctx context.Context, request UnpinCluster) error {
	path := "/api/2.0/clusters/unpin"
	err := a.client.Post(ctx, path, request, nil)
	return err
}
