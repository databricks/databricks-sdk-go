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
		impl: &clustersImpl{
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
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ClustersService)
	impl ClustersService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ClustersAPI) WithImpl(impl ClustersService) *ClustersAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Clusters API implementation
func (a *ClustersAPI) Impl() ClustersService {
	return a.impl
}

// Change cluster owner
//
// Change the owner of the cluster. You must be an admin to perform this
// operation.
func (a *ClustersAPI) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	return a.impl.ChangeOwner(ctx, request)
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
	return a.impl.Create(ctx, request)
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
	return a.impl.Delete(ctx, request)
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
	return a.impl.Delete(ctx, DeleteCluster{
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
	return a.impl.Edit(ctx, request)
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
		response, err := a.impl.Events(ctx, request)
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
	return a.impl.Get(ctx, request)
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
	return a.impl.Get(ctx, GetRequest{
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
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Clusters, nil
}

// ClusterInfoClusterNameToClusterIdMap calls [ClustersAPI.ListAll] and creates a map of results with [ClusterInfo].ClusterName as key and [ClusterInfo].ClusterId as value.
//
// Note: All [ClusterInfo] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) ClusterInfoClusterNameToClusterIdMap(ctx context.Context, request ListRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.ClusterName] = v.ClusterId
	}
	return mapping, nil
}

// GetClusterInfoByClusterName calls [ClustersAPI.ClusterInfoClusterNameToClusterIdMap] and returns a single [ClusterInfo].
//
// Note: All [ClusterInfo] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) GetClusterInfoByClusterName(ctx context.Context, name string) (*ClusterInfo, error) {
	result, err := a.ListAll(ctx, ListRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.ClusterName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("ClusterInfo named '%s' does not exist", name)
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
	return a.impl.List(ctx, ListRequest{
		CanUseClient: canUseClient,
	})
}

// List node types
//
// Returns a list of supported Spark node types. These node types can be used to
// launch a cluster.
func (a *ClustersAPI) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	return a.impl.ListNodeTypes(ctx)
}

// List availability zones
//
// Returns a list of availability zones where clusters can be created in (For
// example, us-west-2a). These zones can be used to launch a cluster.
func (a *ClustersAPI) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	return a.impl.ListZones(ctx)
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
	return a.impl.PermanentDelete(ctx, request)
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
	return a.impl.PermanentDelete(ctx, PermanentDeleteCluster{
		ClusterId: clusterId,
	})
}

// Pin cluster
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) Pin(ctx context.Context, request PinCluster) error {
	return a.impl.Pin(ctx, request)
}

// Pin cluster
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) PinByClusterId(ctx context.Context, clusterId string) error {
	return a.impl.Pin(ctx, PinCluster{
		ClusterId: clusterId,
	})
}

// Resize cluster
//
// Resizes a cluster to have a desired number of workers. This will fail unless
// the cluster is in a `RUNNING` state.
func (a *ClustersAPI) Resize(ctx context.Context, request ResizeCluster) error {
	return a.impl.Resize(ctx, request)
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
	return a.impl.Restart(ctx, request)
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
	return a.impl.SparkVersions(ctx)
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
	return a.impl.Start(ctx, request)
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
	return a.impl.Start(ctx, StartCluster{
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
	return a.impl.Unpin(ctx, request)
}

// Unpin cluster
//
// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins.
func (a *ClustersAPI) UnpinByClusterId(ctx context.Context, clusterId string) error {
	return a.impl.Unpin(ctx, UnpinCluster{
		ClusterId: clusterId,
	})
}

func NewInstanceProfiles(client *client.DatabricksClient) *InstanceProfilesAPI {
	return &InstanceProfilesAPI{
		impl: &instanceProfilesImpl{
			client: client,
		},
	}
}

// The Instance Profiles API allows admins to add, list, and remove instance
// profiles that users can launch clusters with. Regular users can list the
// instance profiles available to them. See [Secure access to S3
// buckets](https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html)
// using instance profiles for more information.
type InstanceProfilesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(InstanceProfilesService)
	impl InstanceProfilesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *InstanceProfilesAPI) WithImpl(impl InstanceProfilesService) *InstanceProfilesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level InstanceProfiles API implementation
func (a *InstanceProfilesAPI) Impl() InstanceProfilesService {
	return a.impl
}

// Register an instance profile
//
// In the UI, you can select the instance profile when launching clusters. This
// API is only available to admin users.
func (a *InstanceProfilesAPI) Add(ctx context.Context, request AddInstanceProfile) error {
	return a.impl.Add(ctx, request)
}

// Edit an instance profile
//
// The only supported field to change is the optional IAM role ARN associated
// with the instance profile. It is required to specify the IAM role ARN if both
// of the following are true:
//
// * Your role name and instance profile name do not match. The name is the part
// after the last slash in each ARN. * You want to use the instance profile with
// [Databricks SQL
// Serverless](https://docs.databricks.com/sql/admin/serverless.html).
//
// To understand where these fields are in the AWS console, see [Enable
// serverless SQL
// warehouses](https://docs.databricks.com/sql/admin/serverless.html).
//
// This API is only available to admin users.
func (a *InstanceProfilesAPI) Edit(ctx context.Context, request InstanceProfile) error {
	return a.impl.Edit(ctx, request)
}

// List available instance profiles
//
// List the instance profiles that the calling user can use to launch a cluster.
//
// This API is available to all users.
//
// This method is generated by Databricks SDK Code Generator.
func (a *InstanceProfilesAPI) ListAll(ctx context.Context) ([]InstanceProfile, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.InstanceProfiles, nil
}

// Remove the instance profile
//
// Remove the instance profile with the provided ARN. Existing clusters with
// this instance profile will continue to function.
//
// This API is only accessible to admin users.
func (a *InstanceProfilesAPI) Remove(ctx context.Context, request RemoveInstanceProfile) error {
	return a.impl.Remove(ctx, request)
}

// Remove the instance profile
//
// Remove the instance profile with the provided ARN. Existing clusters with
// this instance profile will continue to function.
//
// This API is only accessible to admin users.
func (a *InstanceProfilesAPI) RemoveByInstanceProfileArn(ctx context.Context, instanceProfileArn string) error {
	return a.impl.Remove(ctx, RemoveInstanceProfile{
		InstanceProfileArn: instanceProfileArn,
	})
}
