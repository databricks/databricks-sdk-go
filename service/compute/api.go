// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Cluster Policies, Clusters, Command Execution, Global Init Scripts, Instance Pools, Instance Profiles, Libraries, Policy Families, etc.
package compute

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewClusterPolicies(client *client.DatabricksClient) *ClusterPoliciesAPI {
	return &ClusterPoliciesAPI{
		impl: &clusterPoliciesImpl{
			client: client,
		},
	}
}

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
type ClusterPoliciesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ClusterPoliciesService)
	impl ClusterPoliciesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ClusterPoliciesAPI) WithImpl(impl ClusterPoliciesService) *ClusterPoliciesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level ClusterPolicies API implementation
func (a *ClusterPoliciesAPI) Impl() ClusterPoliciesService {
	return a.impl
}

// Create a new policy.
//
// Creates a new policy with prescribed settings.
func (a *ClusterPoliciesAPI) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a cluster policy.
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) Delete(ctx context.Context, request DeletePolicy) error {
	return a.impl.Delete(ctx, request)
}

// Delete a cluster policy.
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) DeleteByPolicyId(ctx context.Context, policyId string) error {
	return a.impl.Delete(ctx, DeletePolicy{
		PolicyId: policyId,
	})
}

// Update a cluster policy.
//
// Update an existing policy for cluster. This operation may make some clusters
// governed by the previous policy invalid.
func (a *ClusterPoliciesAPI) Edit(ctx context.Context, request EditPolicy) error {
	return a.impl.Edit(ctx, request)
}

// Get entity.
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error) {
	return a.impl.Get(ctx, request)
}

// Get entity.
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) GetByPolicyId(ctx context.Context, policyId string) (*Policy, error) {
	return a.impl.Get(ctx, GetClusterPolicyRequest{
		PolicyId: policyId,
	})
}

// Get cluster policy permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ClusterPoliciesAPI) GetClusterPolicyPermissionLevels(ctx context.Context, request GetClusterPolicyPermissionLevelsRequest) (*GetClusterPolicyPermissionLevelsResponse, error) {
	return a.impl.GetClusterPolicyPermissionLevels(ctx, request)
}

// Get cluster policy permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ClusterPoliciesAPI) GetClusterPolicyPermissionLevelsByClusterPolicyId(ctx context.Context, clusterPolicyId string) (*GetClusterPolicyPermissionLevelsResponse, error) {
	return a.impl.GetClusterPolicyPermissionLevels(ctx, GetClusterPolicyPermissionLevelsRequest{
		ClusterPolicyId: clusterPolicyId,
	})
}

// Get cluster policy permissions.
//
// Gets the permissions of a cluster policy. Cluster policies can inherit
// permissions from their root object.
func (a *ClusterPoliciesAPI) GetClusterPolicyPermissions(ctx context.Context, request GetClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	return a.impl.GetClusterPolicyPermissions(ctx, request)
}

// Get cluster policy permissions.
//
// Gets the permissions of a cluster policy. Cluster policies can inherit
// permissions from their root object.
func (a *ClusterPoliciesAPI) GetClusterPolicyPermissionsByClusterPolicyId(ctx context.Context, clusterPolicyId string) (*ClusterPolicyPermissions, error) {
	return a.impl.GetClusterPolicyPermissions(ctx, GetClusterPolicyPermissionsRequest{
		ClusterPolicyId: clusterPolicyId,
	})
}

// Get a cluster policy.
//
// Returns a list of policies accessible by the requesting user.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) ListAll(ctx context.Context, listing ListClusterPoliciesRequest) ([]Policy, error) {
	response, err := a.impl.List(ctx, listing)
	if err != nil {
		return nil, err
	}
	return response.Policies, nil
}

// PolicyNameToPolicyIdMap calls [ClusterPoliciesAPI.ListAll] and creates a map of results with [Policy].Name as key and [Policy].PolicyId as value.
//
// Returns an error if there's more than one [Policy] with the same .Name.
//
// Note: All [Policy] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) PolicyNameToPolicyIdMap(ctx context.Context, listing ListClusterPoliciesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, listing)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.PolicyId
	}
	return mapping, nil
}

// GetByName calls [ClusterPoliciesAPI.PolicyNameToPolicyIdMap] and returns a single [Policy].
//
// Returns an error if there's more than one [Policy] with the same .Name.
//
// Note: All [Policy] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) GetByName(ctx context.Context, name string) (*Policy, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListClusterPoliciesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]Policy{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("Policy named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of Policy named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Set cluster policy permissions.
//
// Sets permissions on a cluster policy. Cluster policies can inherit
// permissions from their root object.
func (a *ClusterPoliciesAPI) SetClusterPolicyPermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	return a.impl.SetClusterPolicyPermissions(ctx, request)
}

// Update cluster policy permissions.
//
// Updates the permissions on a cluster policy. Cluster policies can inherit
// permissions from their root object.
func (a *ClusterPoliciesAPI) UpdateClusterPolicyPermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	return a.impl.UpdateClusterPolicyPermissions(ctx, request)
}

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

// WaitGetClusterRunning repeatedly calls [ClustersAPI.Get] and waits to reach RUNNING state
func (a *ClustersAPI) WaitGetClusterRunning(ctx context.Context, clusterId string,
	timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[ClusterDetails](ctx, timeout, func() (*ClusterDetails, *retries.Err) {
		clusterDetails, err := a.Get(ctx, GetClusterRequest{
			ClusterId: clusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(clusterDetails)
		}
		status := clusterDetails.State
		statusMessage := clusterDetails.StateMessage
		switch status {
		case StateRunning: // target state
			return clusterDetails, nil
		case StateError, StateTerminated:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetClusterRunning is a wrapper that calls [ClustersAPI.WaitGetClusterRunning] and waits to reach RUNNING state.
type WaitGetClusterRunning[R any] struct {
	Response  *R
	ClusterId string `json:"cluster_id"`
	poll      func(time.Duration, func(*ClusterDetails)) (*ClusterDetails, error)
	callback  func(*ClusterDetails)
	timeout   time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetClusterRunning[R]) OnProgress(callback func(*ClusterDetails)) *WaitGetClusterRunning[R] {
	w.callback = callback
	return w
}

// Get the ClusterDetails with the default timeout of 20 minutes.
func (w *WaitGetClusterRunning[R]) Get() (*ClusterDetails, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the ClusterDetails with custom timeout.
func (w *WaitGetClusterRunning[R]) GetWithTimeout(timeout time.Duration) (*ClusterDetails, error) {
	return w.poll(timeout, w.callback)
}

// WaitGetClusterTerminated repeatedly calls [ClustersAPI.Get] and waits to reach TERMINATED state
func (a *ClustersAPI) WaitGetClusterTerminated(ctx context.Context, clusterId string,
	timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[ClusterDetails](ctx, timeout, func() (*ClusterDetails, *retries.Err) {
		clusterDetails, err := a.Get(ctx, GetClusterRequest{
			ClusterId: clusterId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(clusterDetails)
		}
		status := clusterDetails.State
		statusMessage := clusterDetails.StateMessage
		switch status {
		case StateTerminated: // target state
			return clusterDetails, nil
		case StateError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateTerminated, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetClusterTerminated is a wrapper that calls [ClustersAPI.WaitGetClusterTerminated] and waits to reach TERMINATED state.
type WaitGetClusterTerminated[R any] struct {
	Response  *R
	ClusterId string `json:"cluster_id"`
	poll      func(time.Duration, func(*ClusterDetails)) (*ClusterDetails, error)
	callback  func(*ClusterDetails)
	timeout   time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetClusterTerminated[R]) OnProgress(callback func(*ClusterDetails)) *WaitGetClusterTerminated[R] {
	w.callback = callback
	return w
}

// Get the ClusterDetails with the default timeout of 20 minutes.
func (w *WaitGetClusterTerminated[R]) Get() (*ClusterDetails, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the ClusterDetails with custom timeout.
func (w *WaitGetClusterTerminated[R]) GetWithTimeout(timeout time.Duration) (*ClusterDetails, error) {
	return w.poll(timeout, w.callback)
}

// Change cluster owner.
//
// Change the owner of the cluster. You must be an admin to perform this
// operation.
func (a *ClustersAPI) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	return a.impl.ChangeOwner(ctx, request)
}

// Create new cluster.
//
// Creates a new Spark cluster. This method will acquire new instances from the
// cloud provider if necessary. Note: Databricks may not be able to acquire some
// of the requested nodes, due to cloud provider limitations (account limits,
// spot price, etc.) or transient network issues.
//
// If Databricks acquires at least 85% of the requested on-demand nodes, cluster
// creation will succeed. Otherwise the cluster will terminate with an
// informative error message.
func (a *ClustersAPI) Create(ctx context.Context, createCluster CreateCluster) (*WaitGetClusterRunning[CreateClusterResponse], error) {
	createClusterResponse, err := a.impl.Create(ctx, createCluster)
	if err != nil {
		return nil, err
	}
	return &WaitGetClusterRunning[CreateClusterResponse]{
		Response:  createClusterResponse,
		ClusterId: createClusterResponse.ClusterId,
		poll: func(timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
			return a.WaitGetClusterRunning(ctx, createClusterResponse.ClusterId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ClustersAPI.Create] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterDetails](60*time.Minute) functional option.
//
// Deprecated: use [ClustersAPI.Create].Get() or [ClustersAPI.WaitGetClusterRunning]
func (a *ClustersAPI) CreateAndWait(ctx context.Context, createCluster CreateCluster, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	wait, err := a.Create(ctx, createCluster)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ClusterDetails]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ClusterDetails) {
		for _, o := range options {
			o(&retries.Info[ClusterDetails]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Terminate cluster.
//
// Terminates the Spark cluster with the specified ID. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// `TERMINATED` state. If the cluster is already in a `TERMINATING` or
// `TERMINATED` state, nothing will happen.
func (a *ClustersAPI) Delete(ctx context.Context, deleteCluster DeleteCluster) (*WaitGetClusterTerminated[any], error) {
	err := a.impl.Delete(ctx, deleteCluster)
	if err != nil {
		return nil, err
	}
	return &WaitGetClusterTerminated[any]{

		ClusterId: deleteCluster.ClusterId,
		poll: func(timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
			return a.WaitGetClusterTerminated(ctx, deleteCluster.ClusterId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ClustersAPI.Delete] and waits to reach TERMINATED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterDetails](60*time.Minute) functional option.
//
// Deprecated: use [ClustersAPI.Delete].Get() or [ClustersAPI.WaitGetClusterTerminated]
func (a *ClustersAPI) DeleteAndWait(ctx context.Context, deleteCluster DeleteCluster, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	wait, err := a.Delete(ctx, deleteCluster)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ClusterDetails]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ClusterDetails) {
		for _, o := range options {
			o(&retries.Info[ClusterDetails]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Terminate cluster.
//
// Terminates the Spark cluster with the specified ID. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// `TERMINATED` state. If the cluster is already in a `TERMINATING` or
// `TERMINATED` state, nothing will happen.
func (a *ClustersAPI) DeleteByClusterId(ctx context.Context, clusterId string) error {
	return a.impl.Delete(ctx, DeleteCluster{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) DeleteByClusterIdAndWait(ctx context.Context, clusterId string, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	return a.DeleteAndWait(ctx, DeleteCluster{
		ClusterId: clusterId,
	}, options...)
}

// Update cluster configuration.
//
// Updates the configuration of a cluster to match the provided attributes and
// size. A cluster can be updated if it is in a `RUNNING` or `TERMINATED` state.
//
// If a cluster is updated while in a `RUNNING` state, it will be restarted so
// that the new attributes can take effect.
//
// If a cluster is updated while in a `TERMINATED` state, it will remain
// `TERMINATED`. The next time it is started using the `clusters/start` API, the
// new attributes will take effect. Any attempt to update a cluster in any other
// state will be rejected with an `INVALID_STATE` error code.
//
// Clusters created by the Databricks Jobs service cannot be edited.
func (a *ClustersAPI) Edit(ctx context.Context, editCluster EditCluster) (*WaitGetClusterRunning[any], error) {
	err := a.impl.Edit(ctx, editCluster)
	if err != nil {
		return nil, err
	}
	return &WaitGetClusterRunning[any]{

		ClusterId: editCluster.ClusterId,
		poll: func(timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
			return a.WaitGetClusterRunning(ctx, editCluster.ClusterId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ClustersAPI.Edit] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterDetails](60*time.Minute) functional option.
//
// Deprecated: use [ClustersAPI.Edit].Get() or [ClustersAPI.WaitGetClusterRunning]
func (a *ClustersAPI) EditAndWait(ctx context.Context, editCluster EditCluster, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	wait, err := a.Edit(ctx, editCluster)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ClusterDetails]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ClusterDetails) {
		for _, o := range options {
			o(&retries.Info[ClusterDetails]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// List cluster activity events.
//
// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// nparameters necessary to request the next page of events.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) EventsAll(ctx context.Context, listing GetEvents) ([]ClusterEvent, error) {
	var results []ClusterEvent
	request := GetEvents{
		ClusterId:  listing.ClusterId,
		EndTime:    listing.EndTime,
		EventTypes: listing.EventTypes,
		Limit:      listing.Limit,
		Offset:     listing.Offset,
		Order:      listing.Order,
		StartTime:  listing.StartTime,
	}
	var totalCount int64 = 0
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
		count := int64(len(response.Events))
		totalCount += count
		if response.NextPage == nil {
			break
		}
		// cluster events API returns next page information as part of the result
		request = *response.NextPage
		limit := request.Limit
		if limit > 0 && totalCount >= limit {
			break
		}
	}
	return results, nil
}

// Get cluster info.
//
// Retrieves the information for a cluster given its identifier. Clusters can be
// described while they are running, or up to 60 days after they are terminated.
func (a *ClustersAPI) Get(ctx context.Context, request GetClusterRequest) (*ClusterDetails, error) {
	return a.impl.Get(ctx, request)
}

// Get cluster info.
//
// Retrieves the information for a cluster given its identifier. Clusters can be
// described while they are running, or up to 60 days after they are terminated.
func (a *ClustersAPI) GetByClusterId(ctx context.Context, clusterId string) (*ClusterDetails, error) {
	return a.impl.Get(ctx, GetClusterRequest{
		ClusterId: clusterId,
	})
}

// Get cluster permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ClustersAPI) GetClusterPermissionLevels(ctx context.Context, request GetClusterPermissionLevelsRequest) (*GetClusterPermissionLevelsResponse, error) {
	return a.impl.GetClusterPermissionLevels(ctx, request)
}

// Get cluster permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ClustersAPI) GetClusterPermissionLevelsByClusterId(ctx context.Context, clusterId string) (*GetClusterPermissionLevelsResponse, error) {
	return a.impl.GetClusterPermissionLevels(ctx, GetClusterPermissionLevelsRequest{
		ClusterId: clusterId,
	})
}

// Get cluster permissions.
//
// Gets the permissions of a cluster. Clusters can inherit permissions from
// their root object.
func (a *ClustersAPI) GetClusterPermissions(ctx context.Context, request GetClusterPermissionsRequest) (*ClusterPermissions, error) {
	return a.impl.GetClusterPermissions(ctx, request)
}

// Get cluster permissions.
//
// Gets the permissions of a cluster. Clusters can inherit permissions from
// their root object.
func (a *ClustersAPI) GetClusterPermissionsByClusterId(ctx context.Context, clusterId string) (*ClusterPermissions, error) {
	return a.impl.GetClusterPermissions(ctx, GetClusterPermissionsRequest{
		ClusterId: clusterId,
	})
}

// List all clusters.
//
// Return information about all pinned clusters, active clusters, up to 200 of
// the most recently terminated all-purpose clusters in the past 30 days, and up
// to 30 of the most recently terminated job clusters in the past 30 days.
//
// For example, if there is 1 pinned cluster, 4 active clusters, 45 terminated
// all-purpose clusters in the past 30 days, and 50 terminated job clusters in
// the past 30 days, then this API returns the 1 pinned cluster, 4 active
// clusters, all 45 terminated all-purpose clusters, and the 30 most recently
// terminated job clusters.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) ListAll(ctx context.Context, listing ListClustersRequest) ([]ClusterDetails, error) {
	response, err := a.impl.List(ctx, listing)
	if err != nil {
		return nil, err
	}
	return response.Clusters, nil
}

// ClusterDetailsClusterNameToClusterIdMap calls [ClustersAPI.ListAll] and creates a map of results with [ClusterDetails].ClusterName as key and [ClusterDetails].ClusterId as value.
//
// Returns an error if there's more than one [ClusterDetails] with the same .ClusterName.
//
// Note: All [ClusterDetails] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) ClusterDetailsClusterNameToClusterIdMap(ctx context.Context, listing ListClustersRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, listing)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.ClusterName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .ClusterName: %s", key)
		}
		mapping[key] = v.ClusterId
	}
	return mapping, nil
}

// GetByClusterName calls [ClustersAPI.ClusterDetailsClusterNameToClusterIdMap] and returns a single [ClusterDetails].
//
// Returns an error if there's more than one [ClusterDetails] with the same .ClusterName.
//
// Note: All [ClusterDetails] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) GetByClusterName(ctx context.Context, name string) (*ClusterDetails, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListClustersRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]ClusterDetails{}
	for _, v := range result {
		key := v.ClusterName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("ClusterDetails named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of ClusterDetails named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// List all clusters.
//
// Return information about all pinned clusters, active clusters, up to 200 of
// the most recently terminated all-purpose clusters in the past 30 days, and up
// to 30 of the most recently terminated job clusters in the past 30 days.
//
// For example, if there is 1 pinned cluster, 4 active clusters, 45 terminated
// all-purpose clusters in the past 30 days, and 50 terminated job clusters in
// the past 30 days, then this API returns the 1 pinned cluster, 4 active
// clusters, all 45 terminated all-purpose clusters, and the 30 most recently
// terminated job clusters.
func (a *ClustersAPI) ListByCanUseClient(ctx context.Context, canUseClient string) (*ListClustersResponse, error) {
	return a.impl.List(ctx, ListClustersRequest{
		CanUseClient: canUseClient,
	})
}

// List node types.
//
// Returns a list of supported Spark node types. These node types can be used to
// launch a cluster.
func (a *ClustersAPI) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	return a.impl.ListNodeTypes(ctx)
}

// List availability zones.
//
// Returns a list of availability zones where clusters can be created in (For
// example, us-west-2a). These zones can be used to launch a cluster.
func (a *ClustersAPI) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	return a.impl.ListZones(ctx)
}

// Permanently delete cluster.
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

// Permanently delete cluster.
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

// Pin cluster.
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) Pin(ctx context.Context, request PinCluster) error {
	return a.impl.Pin(ctx, request)
}

// Pin cluster.
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) PinByClusterId(ctx context.Context, clusterId string) error {
	return a.impl.Pin(ctx, PinCluster{
		ClusterId: clusterId,
	})
}

// Resize cluster.
//
// Resizes a cluster to have a desired number of workers. This will fail unless
// the cluster is in a `RUNNING` state.
func (a *ClustersAPI) Resize(ctx context.Context, resizeCluster ResizeCluster) (*WaitGetClusterRunning[any], error) {
	err := a.impl.Resize(ctx, resizeCluster)
	if err != nil {
		return nil, err
	}
	return &WaitGetClusterRunning[any]{

		ClusterId: resizeCluster.ClusterId,
		poll: func(timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
			return a.WaitGetClusterRunning(ctx, resizeCluster.ClusterId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ClustersAPI.Resize] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterDetails](60*time.Minute) functional option.
//
// Deprecated: use [ClustersAPI.Resize].Get() or [ClustersAPI.WaitGetClusterRunning]
func (a *ClustersAPI) ResizeAndWait(ctx context.Context, resizeCluster ResizeCluster, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	wait, err := a.Resize(ctx, resizeCluster)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ClusterDetails]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ClusterDetails) {
		for _, o := range options {
			o(&retries.Info[ClusterDetails]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Restart cluster.
//
// Restarts a Spark cluster with the supplied ID. If the cluster is not
// currently in a `RUNNING` state, nothing will happen.
func (a *ClustersAPI) Restart(ctx context.Context, restartCluster RestartCluster) (*WaitGetClusterRunning[any], error) {
	err := a.impl.Restart(ctx, restartCluster)
	if err != nil {
		return nil, err
	}
	return &WaitGetClusterRunning[any]{

		ClusterId: restartCluster.ClusterId,
		poll: func(timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
			return a.WaitGetClusterRunning(ctx, restartCluster.ClusterId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ClustersAPI.Restart] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterDetails](60*time.Minute) functional option.
//
// Deprecated: use [ClustersAPI.Restart].Get() or [ClustersAPI.WaitGetClusterRunning]
func (a *ClustersAPI) RestartAndWait(ctx context.Context, restartCluster RestartCluster, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	wait, err := a.Restart(ctx, restartCluster)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ClusterDetails]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ClusterDetails) {
		for _, o := range options {
			o(&retries.Info[ClusterDetails]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Set cluster permissions.
//
// Sets permissions on a cluster. Clusters can inherit permissions from their
// root object.
func (a *ClustersAPI) SetClusterPermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {
	return a.impl.SetClusterPermissions(ctx, request)
}

// List available Spark versions.
//
// Returns the list of available Spark versions. These versions can be used to
// launch a cluster.
func (a *ClustersAPI) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	return a.impl.SparkVersions(ctx)
}

// Start terminated cluster.
//
// Starts a terminated Spark cluster with the supplied ID. This works similar to
// `createCluster` except:
//
// * The previous cluster id and attributes are preserved. * The cluster starts
// with the last specified cluster size. * If the previous cluster was an
// autoscaling cluster, the current cluster starts with the minimum number of
// nodes. * If the cluster is not currently in a `TERMINATED` state, nothing
// will happen. * Clusters launched to run a job cannot be started.
func (a *ClustersAPI) Start(ctx context.Context, startCluster StartCluster) (*WaitGetClusterRunning[any], error) {
	err := a.impl.Start(ctx, startCluster)
	if err != nil {
		return nil, err
	}
	return &WaitGetClusterRunning[any]{

		ClusterId: startCluster.ClusterId,
		poll: func(timeout time.Duration, callback func(*ClusterDetails)) (*ClusterDetails, error) {
			return a.WaitGetClusterRunning(ctx, startCluster.ClusterId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [ClustersAPI.Start] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ClusterDetails](60*time.Minute) functional option.
//
// Deprecated: use [ClustersAPI.Start].Get() or [ClustersAPI.WaitGetClusterRunning]
func (a *ClustersAPI) StartAndWait(ctx context.Context, startCluster StartCluster, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	wait, err := a.Start(ctx, startCluster)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ClusterDetails]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ClusterDetails) {
		for _, o := range options {
			o(&retries.Info[ClusterDetails]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Start terminated cluster.
//
// Starts a terminated Spark cluster with the supplied ID. This works similar to
// `createCluster` except:
//
// * The previous cluster id and attributes are preserved. * The cluster starts
// with the last specified cluster size. * If the previous cluster was an
// autoscaling cluster, the current cluster starts with the minimum number of
// nodes. * If the cluster is not currently in a `TERMINATED` state, nothing
// will happen. * Clusters launched to run a job cannot be started.
func (a *ClustersAPI) StartByClusterId(ctx context.Context, clusterId string) error {
	return a.impl.Start(ctx, StartCluster{
		ClusterId: clusterId,
	})
}

func (a *ClustersAPI) StartByClusterIdAndWait(ctx context.Context, clusterId string, options ...retries.Option[ClusterDetails]) (*ClusterDetails, error) {
	return a.StartAndWait(ctx, StartCluster{
		ClusterId: clusterId,
	}, options...)
}

// Unpin cluster.
//
// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins.
func (a *ClustersAPI) Unpin(ctx context.Context, request UnpinCluster) error {
	return a.impl.Unpin(ctx, request)
}

// Unpin cluster.
//
// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins.
func (a *ClustersAPI) UnpinByClusterId(ctx context.Context, clusterId string) error {
	return a.impl.Unpin(ctx, UnpinCluster{
		ClusterId: clusterId,
	})
}

// Update cluster permissions.
//
// Updates the permissions on a cluster. Clusters can inherit permissions from
// their root object.
func (a *ClustersAPI) UpdateClusterPermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {
	return a.impl.UpdateClusterPermissions(ctx, request)
}

func NewCommandExecution(client *client.DatabricksClient) *CommandExecutionAPI {
	return &CommandExecutionAPI{
		impl: &commandExecutionImpl{
			client: client,
		},
	}
}

// This API allows execution of Python, Scala, SQL, or R commands on running
// Databricks Clusters.
type CommandExecutionAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CommandExecutionService)
	impl CommandExecutionService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CommandExecutionAPI) WithImpl(impl CommandExecutionService) *CommandExecutionAPI {
	a.impl = impl
	return a
}

// Impl returns low-level CommandExecution API implementation
func (a *CommandExecutionAPI) Impl() CommandExecutionService {
	return a.impl
}

// WaitCommandStatusCommandExecutionCancelled repeatedly calls [CommandExecutionAPI.CommandStatus] and waits to reach Cancelled state
func (a *CommandExecutionAPI) WaitCommandStatusCommandExecutionCancelled(ctx context.Context, clusterId string, commandId string, contextId string,
	timeout time.Duration, callback func(*CommandStatusResponse)) (*CommandStatusResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[CommandStatusResponse](ctx, timeout, func() (*CommandStatusResponse, *retries.Err) {
		commandStatusResponse, err := a.CommandStatus(ctx, CommandStatusRequest{
			ClusterId: clusterId,
			CommandId: commandId,
			ContextId: contextId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(commandStatusResponse)
		}
		status := commandStatusResponse.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		if commandStatusResponse.Results != nil {
			statusMessage = commandStatusResponse.Results.Cause
		}
		switch status {
		case CommandStatusCancelled: // target state
			return commandStatusResponse, nil
		case CommandStatusError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				CommandStatusCancelled, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitCommandStatusCommandExecutionCancelled is a wrapper that calls [CommandExecutionAPI.WaitCommandStatusCommandExecutionCancelled] and waits to reach Cancelled state.
type WaitCommandStatusCommandExecutionCancelled[R any] struct {
	Response  *R
	ClusterId string `json:"clusterId"`
	CommandId string `json:"commandId"`
	ContextId string `json:"contextId"`
	poll      func(time.Duration, func(*CommandStatusResponse)) (*CommandStatusResponse, error)
	callback  func(*CommandStatusResponse)
	timeout   time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitCommandStatusCommandExecutionCancelled[R]) OnProgress(callback func(*CommandStatusResponse)) *WaitCommandStatusCommandExecutionCancelled[R] {
	w.callback = callback
	return w
}

// Get the CommandStatusResponse with the default timeout of 20 minutes.
func (w *WaitCommandStatusCommandExecutionCancelled[R]) Get() (*CommandStatusResponse, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the CommandStatusResponse with custom timeout.
func (w *WaitCommandStatusCommandExecutionCancelled[R]) GetWithTimeout(timeout time.Duration) (*CommandStatusResponse, error) {
	return w.poll(timeout, w.callback)
}

// WaitCommandStatusCommandExecutionFinishedOrError repeatedly calls [CommandExecutionAPI.CommandStatus] and waits to reach Finished or Error state
func (a *CommandExecutionAPI) WaitCommandStatusCommandExecutionFinishedOrError(ctx context.Context, clusterId string, commandId string, contextId string,
	timeout time.Duration, callback func(*CommandStatusResponse)) (*CommandStatusResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[CommandStatusResponse](ctx, timeout, func() (*CommandStatusResponse, *retries.Err) {
		commandStatusResponse, err := a.CommandStatus(ctx, CommandStatusRequest{
			ClusterId: clusterId,
			CommandId: commandId,
			ContextId: contextId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(commandStatusResponse)
		}
		status := commandStatusResponse.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case CommandStatusFinished, CommandStatusError: // target state
			return commandStatusResponse, nil
		case CommandStatusCancelled, CommandStatusCancelling:
			err := fmt.Errorf("failed to reach %s or %s, got %s: %s",
				CommandStatusFinished, CommandStatusError, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitCommandStatusCommandExecutionFinishedOrError is a wrapper that calls [CommandExecutionAPI.WaitCommandStatusCommandExecutionFinishedOrError] and waits to reach Finished or Error state.
type WaitCommandStatusCommandExecutionFinishedOrError[R any] struct {
	Response  *R
	ClusterId string `json:"clusterId"`
	CommandId string `json:"commandId"`
	ContextId string `json:"contextId"`
	poll      func(time.Duration, func(*CommandStatusResponse)) (*CommandStatusResponse, error)
	callback  func(*CommandStatusResponse)
	timeout   time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitCommandStatusCommandExecutionFinishedOrError[R]) OnProgress(callback func(*CommandStatusResponse)) *WaitCommandStatusCommandExecutionFinishedOrError[R] {
	w.callback = callback
	return w
}

// Get the CommandStatusResponse with the default timeout of 20 minutes.
func (w *WaitCommandStatusCommandExecutionFinishedOrError[R]) Get() (*CommandStatusResponse, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the CommandStatusResponse with custom timeout.
func (w *WaitCommandStatusCommandExecutionFinishedOrError[R]) GetWithTimeout(timeout time.Duration) (*CommandStatusResponse, error) {
	return w.poll(timeout, w.callback)
}

// WaitContextStatusCommandExecutionRunning repeatedly calls [CommandExecutionAPI.ContextStatus] and waits to reach Running state
func (a *CommandExecutionAPI) WaitContextStatusCommandExecutionRunning(ctx context.Context, clusterId string, contextId string,
	timeout time.Duration, callback func(*ContextStatusResponse)) (*ContextStatusResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[ContextStatusResponse](ctx, timeout, func() (*ContextStatusResponse, *retries.Err) {
		contextStatusResponse, err := a.ContextStatus(ctx, ContextStatusRequest{
			ClusterId: clusterId,
			ContextId: contextId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(contextStatusResponse)
		}
		status := contextStatusResponse.Status
		statusMessage := fmt.Sprintf("current status: %s", status)
		switch status {
		case ContextStatusRunning: // target state
			return contextStatusResponse, nil
		case ContextStatusError:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				ContextStatusRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitContextStatusCommandExecutionRunning is a wrapper that calls [CommandExecutionAPI.WaitContextStatusCommandExecutionRunning] and waits to reach Running state.
type WaitContextStatusCommandExecutionRunning[R any] struct {
	Response  *R
	ClusterId string `json:"clusterId"`
	ContextId string `json:"contextId"`
	poll      func(time.Duration, func(*ContextStatusResponse)) (*ContextStatusResponse, error)
	callback  func(*ContextStatusResponse)
	timeout   time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitContextStatusCommandExecutionRunning[R]) OnProgress(callback func(*ContextStatusResponse)) *WaitContextStatusCommandExecutionRunning[R] {
	w.callback = callback
	return w
}

// Get the ContextStatusResponse with the default timeout of 20 minutes.
func (w *WaitContextStatusCommandExecutionRunning[R]) Get() (*ContextStatusResponse, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the ContextStatusResponse with custom timeout.
func (w *WaitContextStatusCommandExecutionRunning[R]) GetWithTimeout(timeout time.Duration) (*ContextStatusResponse, error) {
	return w.poll(timeout, w.callback)
}

// Cancel a command.
//
// Cancels a currently running command within an execution context.
//
// The command ID is obtained from a prior successful call to __execute__.
func (a *CommandExecutionAPI) Cancel(ctx context.Context, cancelCommand CancelCommand) (*WaitCommandStatusCommandExecutionCancelled[any], error) {
	err := a.impl.Cancel(ctx, cancelCommand)
	if err != nil {
		return nil, err
	}
	return &WaitCommandStatusCommandExecutionCancelled[any]{

		ClusterId: cancelCommand.ClusterId,
		CommandId: cancelCommand.CommandId,
		ContextId: cancelCommand.ContextId,
		poll: func(timeout time.Duration, callback func(*CommandStatusResponse)) (*CommandStatusResponse, error) {
			return a.WaitCommandStatusCommandExecutionCancelled(ctx, cancelCommand.ClusterId, cancelCommand.CommandId, cancelCommand.ContextId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [CommandExecutionAPI.Cancel] and waits to reach Cancelled state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[CommandStatusResponse](60*time.Minute) functional option.
//
// Deprecated: use [CommandExecutionAPI.Cancel].Get() or [CommandExecutionAPI.WaitCommandStatusCommandExecutionCancelled]
func (a *CommandExecutionAPI) CancelAndWait(ctx context.Context, cancelCommand CancelCommand, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error) {
	wait, err := a.Cancel(ctx, cancelCommand)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[CommandStatusResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *CommandStatusResponse) {
		for _, o := range options {
			o(&retries.Info[CommandStatusResponse]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Get command info.
//
// Gets the status of and, if available, the results from a currently executing
// command.
//
// The command ID is obtained from a prior successful call to __execute__.
func (a *CommandExecutionAPI) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	return a.impl.CommandStatus(ctx, request)
}

// Get status.
//
// Gets the status for an execution context.
func (a *CommandExecutionAPI) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	return a.impl.ContextStatus(ctx, request)
}

// Create an execution context.
//
// Creates an execution context for running cluster commands.
//
// If successful, this method returns the ID of the new execution context.
func (a *CommandExecutionAPI) Create(ctx context.Context, createContext CreateContext) (*WaitContextStatusCommandExecutionRunning[Created], error) {
	created, err := a.impl.Create(ctx, createContext)
	if err != nil {
		return nil, err
	}
	return &WaitContextStatusCommandExecutionRunning[Created]{
		Response:  created,
		ClusterId: createContext.ClusterId,
		ContextId: created.Id,
		poll: func(timeout time.Duration, callback func(*ContextStatusResponse)) (*ContextStatusResponse, error) {
			return a.WaitContextStatusCommandExecutionRunning(ctx, createContext.ClusterId, created.Id, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [CommandExecutionAPI.Create] and waits to reach Running state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[ContextStatusResponse](60*time.Minute) functional option.
//
// Deprecated: use [CommandExecutionAPI.Create].Get() or [CommandExecutionAPI.WaitContextStatusCommandExecutionRunning]
func (a *CommandExecutionAPI) CreateAndWait(ctx context.Context, createContext CreateContext, options ...retries.Option[ContextStatusResponse]) (*ContextStatusResponse, error) {
	wait, err := a.Create(ctx, createContext)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[ContextStatusResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *ContextStatusResponse) {
		for _, o := range options {
			o(&retries.Info[ContextStatusResponse]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete an execution context.
//
// Deletes an execution context.
func (a *CommandExecutionAPI) Destroy(ctx context.Context, request DestroyContext) error {
	return a.impl.Destroy(ctx, request)
}

// Run a command.
//
// Runs a cluster command in the given execution context, using the provided
// language.
//
// If successful, it returns an ID for tracking the status of the command's
// execution.
func (a *CommandExecutionAPI) Execute(ctx context.Context, command Command) (*WaitCommandStatusCommandExecutionFinishedOrError[Created], error) {
	created, err := a.impl.Execute(ctx, command)
	if err != nil {
		return nil, err
	}
	return &WaitCommandStatusCommandExecutionFinishedOrError[Created]{
		Response:  created,
		ClusterId: command.ClusterId,
		CommandId: created.Id,
		ContextId: command.ContextId,
		poll: func(timeout time.Duration, callback func(*CommandStatusResponse)) (*CommandStatusResponse, error) {
			return a.WaitCommandStatusCommandExecutionFinishedOrError(ctx, command.ClusterId, created.Id, command.ContextId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [CommandExecutionAPI.Execute] and waits to reach Finished or Error state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[CommandStatusResponse](60*time.Minute) functional option.
//
// Deprecated: use [CommandExecutionAPI.Execute].Get() or [CommandExecutionAPI.WaitCommandStatusCommandExecutionFinishedOrError]
func (a *CommandExecutionAPI) ExecuteAndWait(ctx context.Context, command Command, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error) {
	wait, err := a.Execute(ctx, command)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[CommandStatusResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *CommandStatusResponse) {
		for _, o := range options {
			o(&retries.Info[CommandStatusResponse]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

func NewGlobalInitScripts(client *client.DatabricksClient) *GlobalInitScriptsAPI {
	return &GlobalInitScriptsAPI{
		impl: &globalInitScriptsImpl{
			client: client,
		},
	}
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
type GlobalInitScriptsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(GlobalInitScriptsService)
	impl GlobalInitScriptsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *GlobalInitScriptsAPI) WithImpl(impl GlobalInitScriptsService) *GlobalInitScriptsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level GlobalInitScripts API implementation
func (a *GlobalInitScriptsAPI) Impl() GlobalInitScriptsService {
	return a.impl
}

// Create init script.
//
// Creates a new global init script in this workspace.
func (a *GlobalInitScriptsAPI) Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete init script.
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete init script.
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) DeleteByScriptId(ctx context.Context, scriptId string) error {
	return a.impl.Delete(ctx, DeleteGlobalInitScriptRequest{
		ScriptId: scriptId,
	})
}

// Get an init script.
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	return a.impl.Get(ctx, request)
}

// Get an init script.
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetByScriptId(ctx context.Context, scriptId string) (*GlobalInitScriptDetailsWithContent, error) {
	return a.impl.Get(ctx, GetGlobalInitScriptRequest{
		ScriptId: scriptId,
	})
}

// Get init scripts.
//
// Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](#operation/get-script) operation.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GlobalInitScriptsAPI) ListAll(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Scripts, nil
}

// GlobalInitScriptDetailsNameToScriptIdMap calls [GlobalInitScriptsAPI.ListAll] and creates a map of results with [GlobalInitScriptDetails].Name as key and [GlobalInitScriptDetails].ScriptId as value.
//
// Returns an error if there's more than one [GlobalInitScriptDetails] with the same .Name.
//
// Note: All [GlobalInitScriptDetails] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GlobalInitScriptsAPI) GlobalInitScriptDetailsNameToScriptIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.ScriptId
	}
	return mapping, nil
}

// GetByName calls [GlobalInitScriptsAPI.GlobalInitScriptDetailsNameToScriptIdMap] and returns a single [GlobalInitScriptDetails].
//
// Returns an error if there's more than one [GlobalInitScriptDetails] with the same .Name.
//
// Note: All [GlobalInitScriptDetails] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *GlobalInitScriptsAPI) GetByName(ctx context.Context, name string) (*GlobalInitScriptDetails, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]GlobalInitScriptDetails{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("GlobalInitScriptDetails named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of GlobalInitScriptDetails named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Update init script.
//
// Updates a global init script, specifying only the fields to change. All
// fields are optional. Unspecified fields retain their current value.
func (a *GlobalInitScriptsAPI) Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	return a.impl.Update(ctx, request)
}

func NewInstancePools(client *client.DatabricksClient) *InstancePoolsAPI {
	return &InstancePoolsAPI{
		impl: &instancePoolsImpl{
			client: client,
		},
	}
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
type InstancePoolsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(InstancePoolsService)
	impl InstancePoolsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *InstancePoolsAPI) WithImpl(impl InstancePoolsService) *InstancePoolsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level InstancePools API implementation
func (a *InstancePoolsAPI) Impl() InstancePoolsService {
	return a.impl
}

// Create a new instance pool.
//
// Creates a new instance pool using idle and ready-to-use cloud instances.
func (a *InstancePoolsAPI) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete an instance pool.
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) Delete(ctx context.Context, request DeleteInstancePool) error {
	return a.impl.Delete(ctx, request)
}

// Delete an instance pool.
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error {
	return a.impl.Delete(ctx, DeleteInstancePool{
		InstancePoolId: instancePoolId,
	})
}

// Edit an existing instance pool.
//
// Modifies the configuration of an existing instance pool.
func (a *InstancePoolsAPI) Edit(ctx context.Context, request EditInstancePool) error {
	return a.impl.Edit(ctx, request)
}

// Get instance pool information.
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error) {
	return a.impl.Get(ctx, request)
}

// Get instance pool information.
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePool, error) {
	return a.impl.Get(ctx, GetInstancePoolRequest{
		InstancePoolId: instancePoolId,
	})
}

// Get instance pool permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *InstancePoolsAPI) GetInstancePoolPermissionLevels(ctx context.Context, request GetInstancePoolPermissionLevelsRequest) (*GetInstancePoolPermissionLevelsResponse, error) {
	return a.impl.GetInstancePoolPermissionLevels(ctx, request)
}

// Get instance pool permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *InstancePoolsAPI) GetInstancePoolPermissionLevelsByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePoolPermissionLevelsResponse, error) {
	return a.impl.GetInstancePoolPermissionLevels(ctx, GetInstancePoolPermissionLevelsRequest{
		InstancePoolId: instancePoolId,
	})
}

// Get instance pool permissions.
//
// Gets the permissions of an instance pool. Instance pools can inherit
// permissions from their root object.
func (a *InstancePoolsAPI) GetInstancePoolPermissions(ctx context.Context, request GetInstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	return a.impl.GetInstancePoolPermissions(ctx, request)
}

// Get instance pool permissions.
//
// Gets the permissions of an instance pool. Instance pools can inherit
// permissions from their root object.
func (a *InstancePoolsAPI) GetInstancePoolPermissionsByInstancePoolId(ctx context.Context, instancePoolId string) (*InstancePoolPermissions, error) {
	return a.impl.GetInstancePoolPermissions(ctx, GetInstancePoolPermissionsRequest{
		InstancePoolId: instancePoolId,
	})
}

// List instance pool info.
//
// Gets a list of instance pools with their statistics.
//
// This method is generated by Databricks SDK Code Generator.
func (a *InstancePoolsAPI) ListAll(ctx context.Context) ([]InstancePoolAndStats, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.InstancePools, nil
}

// InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap calls [InstancePoolsAPI.ListAll] and creates a map of results with [InstancePoolAndStats].InstancePoolName as key and [InstancePoolAndStats].InstancePoolId as value.
//
// Returns an error if there's more than one [InstancePoolAndStats] with the same .InstancePoolName.
//
// Note: All [InstancePoolAndStats] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *InstancePoolsAPI) InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.InstancePoolName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .InstancePoolName: %s", key)
		}
		mapping[key] = v.InstancePoolId
	}
	return mapping, nil
}

// GetByInstancePoolName calls [InstancePoolsAPI.InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap] and returns a single [InstancePoolAndStats].
//
// Returns an error if there's more than one [InstancePoolAndStats] with the same .InstancePoolName.
//
// Note: All [InstancePoolAndStats] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *InstancePoolsAPI) GetByInstancePoolName(ctx context.Context, name string) (*InstancePoolAndStats, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]InstancePoolAndStats{}
	for _, v := range result {
		key := v.InstancePoolName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("InstancePoolAndStats named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of InstancePoolAndStats named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Set instance pool permissions.
//
// Sets permissions on an instance pool. Instance pools can inherit permissions
// from their root object.
func (a *InstancePoolsAPI) SetInstancePoolPermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	return a.impl.SetInstancePoolPermissions(ctx, request)
}

// Update instance pool permissions.
//
// Updates the permissions on an instance pool. Instance pools can inherit
// permissions from their root object.
func (a *InstancePoolsAPI) UpdateInstancePoolPermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	return a.impl.UpdateInstancePoolPermissions(ctx, request)
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
// instance profiles available to them. See [Secure access to S3 buckets] using
// instance profiles for more information.
//
// [Secure access to S3 buckets]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html
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

// Register an instance profile.
//
// In the UI, you can select the instance profile when launching clusters. This
// API is only available to admin users.
func (a *InstanceProfilesAPI) Add(ctx context.Context, request AddInstanceProfile) error {
	return a.impl.Add(ctx, request)
}

// Edit an instance profile.
//
// The only supported field to change is the optional IAM role ARN associated
// with the instance profile. It is required to specify the IAM role ARN if both
// of the following are true:
//
// * Your role name and instance profile name do not match. The name is the part
// after the last slash in each ARN. * You want to use the instance profile with
// [Databricks SQL Serverless].
//
// To understand where these fields are in the AWS console, see [Enable
// serverless SQL warehouses].
//
// This API is only available to admin users.
//
// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
// [Enable serverless SQL warehouses]: https://docs.databricks.com/sql/admin/serverless.html
func (a *InstanceProfilesAPI) Edit(ctx context.Context, request InstanceProfile) error {
	return a.impl.Edit(ctx, request)
}

// List available instance profiles.
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

// Remove the instance profile.
//
// Remove the instance profile with the provided ARN. Existing clusters with
// this instance profile will continue to function.
//
// This API is only accessible to admin users.
func (a *InstanceProfilesAPI) Remove(ctx context.Context, request RemoveInstanceProfile) error {
	return a.impl.Remove(ctx, request)
}

// Remove the instance profile.
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

func NewLibraries(client *client.DatabricksClient) *LibrariesAPI {
	return &LibrariesAPI{
		impl: &librariesImpl{
			client: client,
		},
	}
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
type LibrariesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(LibrariesService)
	impl LibrariesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *LibrariesAPI) WithImpl(impl LibrariesService) *LibrariesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Libraries API implementation
func (a *LibrariesAPI) Impl() LibrariesService {
	return a.impl
}

// Get all statuses.
//
// Get the status of all libraries on all clusters. A status will be available
// for all libraries installed on this cluster via the API or the libraries UI
// as well as libraries set to be installed on all clusters via the libraries
// UI.
func (a *LibrariesAPI) AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	return a.impl.AllClusterStatuses(ctx)
}

// Get status.
//
// Get the status of libraries on a cluster. A status will be available for all
// libraries installed on this cluster via the API or the libraries UI as well
// as libraries set to be installed on all clusters via the libraries UI. The
// order of returned libraries will be as follows.
//
// 1. Libraries set to be installed on this cluster will be returned first.
// Within this group, the final order will be order in which the libraries were
// added to the cluster.
//
// 2. Libraries set to be installed on all clusters are returned next. Within
// this group there is no order guarantee.
//
// 3. Libraries that were previously requested on this cluster or on all
// clusters, but now marked for removal. Within this group there is no order
// guarantee.
func (a *LibrariesAPI) ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterLibraryStatuses, error) {
	return a.impl.ClusterStatus(ctx, request)
}

// Get status.
//
// Get the status of libraries on a cluster. A status will be available for all
// libraries installed on this cluster via the API or the libraries UI as well
// as libraries set to be installed on all clusters via the libraries UI. The
// order of returned libraries will be as follows.
//
// 1. Libraries set to be installed on this cluster will be returned first.
// Within this group, the final order will be order in which the libraries were
// added to the cluster.
//
// 2. Libraries set to be installed on all clusters are returned next. Within
// this group there is no order guarantee.
//
// 3. Libraries that were previously requested on this cluster or on all
// clusters, but now marked for removal. Within this group there is no order
// guarantee.
func (a *LibrariesAPI) ClusterStatusByClusterId(ctx context.Context, clusterId string) (*ClusterLibraryStatuses, error) {
	return a.impl.ClusterStatus(ctx, ClusterStatusRequest{
		ClusterId: clusterId,
	})
}

// Add a library.
//
// Add libraries to be installed on a cluster. The installation is asynchronous;
// it happens in the background after the completion of this request.
//
// **Note**: The actual set of libraries to be installed on a cluster is the
// union of the libraries specified via this method and the libraries set to be
// installed on all clusters via the libraries UI.
func (a *LibrariesAPI) Install(ctx context.Context, request InstallLibraries) error {
	return a.impl.Install(ctx, request)
}

// Uninstall libraries.
//
// Set libraries to be uninstalled on a cluster. The libraries won't be
// uninstalled until the cluster is restarted. Uninstalling libraries that are
// not installed on the cluster will have no impact but is not an error.
func (a *LibrariesAPI) Uninstall(ctx context.Context, request UninstallLibraries) error {
	return a.impl.Uninstall(ctx, request)
}

func NewPolicyFamilies(client *client.DatabricksClient) *PolicyFamiliesAPI {
	return &PolicyFamiliesAPI{
		impl: &policyFamiliesImpl{
			client: client,
		},
	}
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
type PolicyFamiliesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PolicyFamiliesService)
	impl PolicyFamiliesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PolicyFamiliesAPI) WithImpl(impl PolicyFamiliesService) *PolicyFamiliesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PolicyFamilies API implementation
func (a *PolicyFamiliesAPI) Impl() PolicyFamiliesService {
	return a.impl
}

// Get policy family information.
//
// Retrieve the information for an policy family based on its identifier.
func (a *PolicyFamiliesAPI) Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error) {
	return a.impl.Get(ctx, request)
}

// Get policy family information.
//
// Retrieve the information for an policy family based on its identifier.
func (a *PolicyFamiliesAPI) GetByPolicyFamilyId(ctx context.Context, policyFamilyId string) (*PolicyFamily, error) {
	return a.impl.Get(ctx, GetPolicyFamilyRequest{
		PolicyFamilyId: policyFamilyId,
	})
}

// List policy families.
//
// Retrieve a list of policy families. This API is paginated.
//
// This method is generated by Databricks SDK Code Generator.
func (a *PolicyFamiliesAPI) ListAll(ctx context.Context, listing ListPolicyFamiliesRequest) ([]PolicyFamily, error) {
	var results []PolicyFamily
	request := ListPolicyFamiliesInternal{
		MaxResults: listing.MaxResults,
	}
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.PolicyFamilies) == 0 {
			break
		}
		for _, v := range response.PolicyFamilies {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}
