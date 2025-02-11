// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just ClusterPolicies API methods
type clusterPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *clusterPoliciesImpl) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {
	var createPolicyResponse CreatePolicyResponse
	path := "/api/2.0/policies/clusters/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

func (a *clusterPoliciesImpl) Delete(ctx context.Context, request DeletePolicy) error {
	var deletePolicyResponse DeletePolicyResponse
	path := "/api/2.0/policies/clusters/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deletePolicyResponse)
	return err
}

func (a *clusterPoliciesImpl) Edit(ctx context.Context, request EditPolicy) error {
	var editPolicyResponse EditPolicyResponse
	path := "/api/2.0/policies/clusters/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &editPolicyResponse)
	return err
}

func (a *clusterPoliciesImpl) Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error) {
	var policy Policy
	path := "/api/2.0/policies/clusters/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &policy)
	return &policy, err
}

func (a *clusterPoliciesImpl) GetPermissionLevels(ctx context.Context, request GetClusterPolicyPermissionLevelsRequest) (*GetClusterPolicyPermissionLevelsResponse, error) {
	var getClusterPolicyPermissionLevelsResponse GetClusterPolicyPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v/permissionLevels", request.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getClusterPolicyPermissionLevelsResponse)
	return &getClusterPolicyPermissionLevelsResponse, err
}

func (a *clusterPoliciesImpl) GetPermissions(ctx context.Context, request GetClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	var clusterPolicyPermissions ClusterPolicyPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &clusterPolicyPermissions)
	return &clusterPolicyPermissions, err
}

// List cluster policies.
//
// Returns a list of policies accessible by the requesting user.
func (a *clusterPoliciesImpl) List(ctx context.Context, request ListClusterPoliciesRequest) listing.Iterator[Policy] {

	getNextPage := func(ctx context.Context, req ListClusterPoliciesRequest) (*ListPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListPoliciesResponse) []Policy {
		return resp.Policies
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List cluster policies.
//
// Returns a list of policies accessible by the requesting user.
func (a *clusterPoliciesImpl) ListAll(ctx context.Context, request ListClusterPoliciesRequest) ([]Policy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[Policy](ctx, iterator)
}

func (a *clusterPoliciesImpl) internalList(ctx context.Context, request ListClusterPoliciesRequest) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listPoliciesResponse)
	return &listPoliciesResponse, err
}

func (a *clusterPoliciesImpl) SetPermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	var clusterPolicyPermissions ClusterPolicyPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &clusterPolicyPermissions)
	return &clusterPolicyPermissions, err
}

func (a *clusterPoliciesImpl) UpdatePermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	var clusterPolicyPermissions ClusterPolicyPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &clusterPolicyPermissions)
	return &clusterPolicyPermissions, err
}

// unexported type that holds implementations of just Clusters API methods
type clustersImpl struct {
	client *client.DatabricksClient
}

func (a *clustersImpl) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	var changeClusterOwnerResponse ChangeClusterOwnerResponse
	path := "/api/2.1/clusters/change-owner"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &changeClusterOwnerResponse)
	return err
}

func (a *clustersImpl) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	var createClusterResponse CreateClusterResponse
	path := "/api/2.1/clusters/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createClusterResponse)
	return &createClusterResponse, err
}

func (a *clustersImpl) Delete(ctx context.Context, request DeleteCluster) error {
	var deleteClusterResponse DeleteClusterResponse
	path := "/api/2.1/clusters/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteClusterResponse)
	return err
}

func (a *clustersImpl) Edit(ctx context.Context, request EditCluster) error {
	var editClusterResponse EditClusterResponse
	path := "/api/2.1/clusters/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &editClusterResponse)
	return err
}

// List cluster activity events.
//
// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// nparameters necessary to request the next page of events.
func (a *clustersImpl) Events(ctx context.Context, request GetEvents) listing.Iterator[ClusterEvent] {

	getNextPage := func(ctx context.Context, req GetEvents) (*GetEventsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalEvents(ctx, req)
	}
	getItems := func(resp *GetEventsResponse) []ClusterEvent {
		return resp.Events
	}
	getNextReq := func(resp *GetEventsResponse) *GetEvents {
		if resp.NextPage == nil {
			return nil
		}
		request = *resp.NextPage

		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List cluster activity events.
//
// Retrieves a list of events about the activity of a cluster. This API is
// paginated. If there are more events to read, the response includes all the
// nparameters necessary to request the next page of events.
func (a *clustersImpl) EventsAll(ctx context.Context, request GetEvents) ([]ClusterEvent, error) {
	iterator := a.Events(ctx, request)
	return listing.ToSliceN[ClusterEvent, int64](ctx, iterator, request.Limit)

}

func (a *clustersImpl) internalEvents(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {
	var getEventsResponse GetEventsResponse
	path := "/api/2.1/clusters/events"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &getEventsResponse)
	return &getEventsResponse, err
}

func (a *clustersImpl) Get(ctx context.Context, request GetClusterRequest) (*ClusterDetails, error) {
	var clusterDetails ClusterDetails
	path := "/api/2.1/clusters/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &clusterDetails)
	return &clusterDetails, err
}

func (a *clustersImpl) GetPermissionLevels(ctx context.Context, request GetClusterPermissionLevelsRequest) (*GetClusterPermissionLevelsResponse, error) {
	var getClusterPermissionLevelsResponse GetClusterPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v/permissionLevels", request.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getClusterPermissionLevelsResponse)
	return &getClusterPermissionLevelsResponse, err
}

func (a *clustersImpl) GetPermissions(ctx context.Context, request GetClusterPermissionsRequest) (*ClusterPermissions, error) {
	var clusterPermissions ClusterPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &clusterPermissions)
	return &clusterPermissions, err
}

// List clusters.
//
// Return information about all pinned and active clusters, and all clusters
// terminated within the last 30 days. Clusters terminated prior to this period
// are not included.
func (a *clustersImpl) List(ctx context.Context, request ListClustersRequest) listing.Iterator[ClusterDetails] {

	getNextPage := func(ctx context.Context, req ListClustersRequest) (*ListClustersResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListClustersResponse) []ClusterDetails {
		return resp.Clusters
	}
	getNextReq := func(resp *ListClustersResponse) *ListClustersRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List clusters.
//
// Return information about all pinned and active clusters, and all clusters
// terminated within the last 30 days. Clusters terminated prior to this period
// are not included.
func (a *clustersImpl) ListAll(ctx context.Context, request ListClustersRequest) ([]ClusterDetails, error) {
	iterator := a.List(ctx, request)
	return listing.ToSliceN[ClusterDetails, int](ctx, iterator, request.PageSize)

}

func (a *clustersImpl) internalList(ctx context.Context, request ListClustersRequest) (*ListClustersResponse, error) {
	var listClustersResponse ListClustersResponse
	path := "/api/2.1/clusters/list"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listClustersResponse)
	return &listClustersResponse, err
}

func (a *clustersImpl) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	var listNodeTypesResponse ListNodeTypesResponse
	path := "/api/2.1/clusters/list-node-types"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listNodeTypesResponse)
	return &listNodeTypesResponse, err
}

func (a *clustersImpl) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	var listAvailableZonesResponse ListAvailableZonesResponse
	path := "/api/2.1/clusters/list-zones"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listAvailableZonesResponse)
	return &listAvailableZonesResponse, err
}

func (a *clustersImpl) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	var permanentDeleteClusterResponse PermanentDeleteClusterResponse
	path := "/api/2.1/clusters/permanent-delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &permanentDeleteClusterResponse)
	return err
}

func (a *clustersImpl) Pin(ctx context.Context, request PinCluster) error {
	var pinClusterResponse PinClusterResponse
	path := "/api/2.1/clusters/pin"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &pinClusterResponse)
	return err
}

func (a *clustersImpl) Resize(ctx context.Context, request ResizeCluster) error {
	var resizeClusterResponse ResizeClusterResponse
	path := "/api/2.1/clusters/resize"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &resizeClusterResponse)
	return err
}

func (a *clustersImpl) Restart(ctx context.Context, request RestartCluster) error {
	var restartClusterResponse RestartClusterResponse
	path := "/api/2.1/clusters/restart"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &restartClusterResponse)
	return err
}

func (a *clustersImpl) SetPermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {
	var clusterPermissions ClusterPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &clusterPermissions)
	return &clusterPermissions, err
}

func (a *clustersImpl) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	var getSparkVersionsResponse GetSparkVersionsResponse
	path := "/api/2.1/clusters/spark-versions"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &getSparkVersionsResponse)
	return &getSparkVersionsResponse, err
}

func (a *clustersImpl) Start(ctx context.Context, request StartCluster) error {
	var startClusterResponse StartClusterResponse
	path := "/api/2.1/clusters/start"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &startClusterResponse)
	return err
}

func (a *clustersImpl) Unpin(ctx context.Context, request UnpinCluster) error {
	var unpinClusterResponse UnpinClusterResponse
	path := "/api/2.1/clusters/unpin"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &unpinClusterResponse)
	return err
}

func (a *clustersImpl) Update(ctx context.Context, request UpdateCluster) error {
	var updateClusterResponse UpdateClusterResponse
	path := "/api/2.1/clusters/update"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &updateClusterResponse)
	return err
}

func (a *clustersImpl) UpdatePermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {
	var clusterPermissions ClusterPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &clusterPermissions)
	return &clusterPermissions, err
}

// unexported type that holds implementations of just CommandExecution API methods
type commandExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *commandExecutionImpl) Cancel(ctx context.Context, request CancelCommand) error {
	var cancelResponse CancelResponse
	path := "/api/1.2/commands/cancel"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &cancelResponse)
	return err
}

func (a *commandExecutionImpl) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	var commandStatusResponse CommandStatusResponse
	path := "/api/1.2/commands/status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &commandStatusResponse)
	return &commandStatusResponse, err
}

func (a *commandExecutionImpl) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	var contextStatusResponse ContextStatusResponse
	path := "/api/1.2/contexts/status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &contextStatusResponse)
	return &contextStatusResponse, err
}

func (a *commandExecutionImpl) Create(ctx context.Context, request CreateContext) (*Created, error) {
	var created Created
	path := "/api/1.2/contexts/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &created)
	return &created, err
}

func (a *commandExecutionImpl) Destroy(ctx context.Context, request DestroyContext) error {
	var destroyResponse DestroyResponse
	path := "/api/1.2/contexts/destroy"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &destroyResponse)
	return err
}

func (a *commandExecutionImpl) Execute(ctx context.Context, request Command) (*Created, error) {
	var created Created
	path := "/api/1.2/commands/execute"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &created)
	return &created, err
}

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsImpl) Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/global-init-scripts"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createResponse)
	return &createResponse, err
}

func (a *globalInitScriptsImpl) Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *globalInitScriptsImpl) Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

// Get init scripts.
//
// Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](:method:globalinitscripts/get) operation.
func (a *globalInitScriptsImpl) List(ctx context.Context) listing.Iterator[GlobalInitScriptDetails] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListGlobalInitScriptsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListGlobalInitScriptsResponse) []GlobalInitScriptDetails {
		return resp.Scripts
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get init scripts.
//
// Get a list of all global init scripts for this workspace. This returns all
// properties for each script but **not** the script contents. To retrieve the
// contents of a script, use the [get a global init
// script](:method:globalinitscripts/get) operation.
func (a *globalInitScriptsImpl) ListAll(ctx context.Context) ([]GlobalInitScriptDetails, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[GlobalInitScriptDetails](ctx, iterator)
}

func (a *globalInitScriptsImpl) internalList(ctx context.Context) (*ListGlobalInitScriptsResponse, error) {
	var listGlobalInitScriptsResponse ListGlobalInitScriptsResponse
	path := "/api/2.0/global-init-scripts"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listGlobalInitScriptsResponse)
	return &listGlobalInitScriptsResponse, err
}

func (a *globalInitScriptsImpl) Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just InstancePools API methods
type instancePoolsImpl struct {
	client *client.DatabricksClient
}

func (a *instancePoolsImpl) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

func (a *instancePoolsImpl) Delete(ctx context.Context, request DeleteInstancePool) error {
	var deleteInstancePoolResponse DeleteInstancePoolResponse
	path := "/api/2.0/instance-pools/delete"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &deleteInstancePoolResponse)
	return err
}

func (a *instancePoolsImpl) Edit(ctx context.Context, request EditInstancePool) error {
	var editInstancePoolResponse EditInstancePoolResponse
	path := "/api/2.0/instance-pools/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &editInstancePoolResponse)
	return err
}

func (a *instancePoolsImpl) Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getInstancePool)
	return &getInstancePool, err
}

func (a *instancePoolsImpl) GetPermissionLevels(ctx context.Context, request GetInstancePoolPermissionLevelsRequest) (*GetInstancePoolPermissionLevelsResponse, error) {
	var getInstancePoolPermissionLevelsResponse GetInstancePoolPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v/permissionLevels", request.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getInstancePoolPermissionLevelsResponse)
	return &getInstancePoolPermissionLevelsResponse, err
}

func (a *instancePoolsImpl) GetPermissions(ctx context.Context, request GetInstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	var instancePoolPermissions InstancePoolPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &instancePoolPermissions)
	return &instancePoolPermissions, err
}

// List instance pool info.
//
// Gets a list of instance pools with their statistics.
func (a *instancePoolsImpl) List(ctx context.Context) listing.Iterator[InstancePoolAndStats] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListInstancePools, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListInstancePools) []InstancePoolAndStats {
		return resp.InstancePools
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List instance pool info.
//
// Gets a list of instance pools with their statistics.
func (a *instancePoolsImpl) ListAll(ctx context.Context) ([]InstancePoolAndStats, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[InstancePoolAndStats](ctx, iterator)
}

func (a *instancePoolsImpl) internalList(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listInstancePools)
	return &listInstancePools, err
}

func (a *instancePoolsImpl) SetPermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	var instancePoolPermissions InstancePoolPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &instancePoolPermissions)
	return &instancePoolPermissions, err
}

func (a *instancePoolsImpl) UpdatePermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	var instancePoolPermissions InstancePoolPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &instancePoolPermissions)
	return &instancePoolPermissions, err
}

// unexported type that holds implementations of just InstanceProfiles API methods
type instanceProfilesImpl struct {
	client *client.DatabricksClient
}

func (a *instanceProfilesImpl) Add(ctx context.Context, request AddInstanceProfile) error {
	var addResponse AddResponse
	path := "/api/2.0/instance-profiles/add"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &addResponse)
	return err
}

func (a *instanceProfilesImpl) Edit(ctx context.Context, request InstanceProfile) error {
	var editResponse EditResponse
	path := "/api/2.0/instance-profiles/edit"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &editResponse)
	return err
}

// List available instance profiles.
//
// List the instance profiles that the calling user can use to launch a cluster.
//
// This API is available to all users.
func (a *instanceProfilesImpl) List(ctx context.Context) listing.Iterator[InstanceProfile] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListInstanceProfilesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx)
	}
	getItems := func(resp *ListInstanceProfilesResponse) []InstanceProfile {
		return resp.InstanceProfiles
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// List available instance profiles.
//
// List the instance profiles that the calling user can use to launch a cluster.
//
// This API is available to all users.
func (a *instanceProfilesImpl) ListAll(ctx context.Context) ([]InstanceProfile, error) {
	iterator := a.List(ctx)
	return listing.ToSlice[InstanceProfile](ctx, iterator)
}

func (a *instanceProfilesImpl) internalList(ctx context.Context) (*ListInstanceProfilesResponse, error) {
	var listInstanceProfilesResponse ListInstanceProfilesResponse
	path := "/api/2.0/instance-profiles/list"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listInstanceProfilesResponse)
	return &listInstanceProfilesResponse, err
}

func (a *instanceProfilesImpl) Remove(ctx context.Context, request RemoveInstanceProfile) error {
	var removeResponse RemoveResponse
	path := "/api/2.0/instance-profiles/remove"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &removeResponse)
	return err
}

// unexported type that holds implementations of just Libraries API methods
type librariesImpl struct {
	client *client.DatabricksClient
}

// Get all statuses.
//
// Get the status of all libraries on all clusters. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI.
func (a *librariesImpl) AllClusterStatuses(ctx context.Context) listing.Iterator[ClusterLibraryStatuses] {
	request := struct{}{}

	getNextPage := func(ctx context.Context, req struct{}) (*ListAllClusterLibraryStatusesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalAllClusterStatuses(ctx)
	}
	getItems := func(resp *ListAllClusterLibraryStatusesResponse) []ClusterLibraryStatuses {
		return resp.Statuses
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all statuses.
//
// Get the status of all libraries on all clusters. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI.
func (a *librariesImpl) AllClusterStatusesAll(ctx context.Context) ([]ClusterLibraryStatuses, error) {
	iterator := a.AllClusterStatuses(ctx)
	return listing.ToSlice[ClusterLibraryStatuses](ctx, iterator)
}

func (a *librariesImpl) internalAllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	var listAllClusterLibraryStatusesResponse ListAllClusterLibraryStatusesResponse
	path := "/api/2.0/libraries/all-cluster-statuses"

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &listAllClusterLibraryStatusesResponse)
	return &listAllClusterLibraryStatusesResponse, err
}

// Get status.
//
// Get the status of libraries on a cluster. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI. The
// order of returned libraries is as follows: 1. Libraries set to be installed
// on this cluster, in the order that the libraries were added to the cluster,
// are returned first. 2. Libraries that were previously requested to be
// installed on this cluster or, but are now marked for removal, in no
// particular order, are returned last.
func (a *librariesImpl) ClusterStatus(ctx context.Context, request ClusterStatus) listing.Iterator[LibraryFullStatus] {

	getNextPage := func(ctx context.Context, req ClusterStatus) (*ClusterLibraryStatuses, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalClusterStatus(ctx, req)
	}
	getItems := func(resp *ClusterLibraryStatuses) []LibraryFullStatus {
		return resp.LibraryStatuses
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get status.
//
// Get the status of libraries on a cluster. A status is returned for all
// libraries installed on this cluster via the API or the libraries UI. The
// order of returned libraries is as follows: 1. Libraries set to be installed
// on this cluster, in the order that the libraries were added to the cluster,
// are returned first. 2. Libraries that were previously requested to be
// installed on this cluster or, but are now marked for removal, in no
// particular order, are returned last.
func (a *librariesImpl) ClusterStatusAll(ctx context.Context, request ClusterStatus) ([]LibraryFullStatus, error) {
	iterator := a.ClusterStatus(ctx, request)
	return listing.ToSlice[LibraryFullStatus](ctx, iterator)
}

func (a *librariesImpl) internalClusterStatus(ctx context.Context, request ClusterStatus) (*ClusterLibraryStatuses, error) {
	var clusterLibraryStatuses ClusterLibraryStatuses
	path := "/api/2.0/libraries/cluster-status"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &clusterLibraryStatuses)
	return &clusterLibraryStatuses, err
}

func (a *librariesImpl) Install(ctx context.Context, request InstallLibraries) error {
	var installLibrariesResponse InstallLibrariesResponse
	path := "/api/2.0/libraries/install"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &installLibrariesResponse)
	return err
}

func (a *librariesImpl) Uninstall(ctx context.Context, request UninstallLibraries) error {
	var uninstallLibrariesResponse UninstallLibrariesResponse
	path := "/api/2.0/libraries/uninstall"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &uninstallLibrariesResponse)
	return err
}

// unexported type that holds implementations of just PolicyComplianceForClusters API methods
type policyComplianceForClustersImpl struct {
	client *client.DatabricksClient
}

func (a *policyComplianceForClustersImpl) EnforceCompliance(ctx context.Context, request EnforceClusterComplianceRequest) (*EnforceClusterComplianceResponse, error) {
	var enforceClusterComplianceResponse EnforceClusterComplianceResponse
	path := "/api/2.0/policies/clusters/enforce-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &enforceClusterComplianceResponse)
	return &enforceClusterComplianceResponse, err
}

func (a *policyComplianceForClustersImpl) GetCompliance(ctx context.Context, request GetClusterComplianceRequest) (*GetClusterComplianceResponse, error) {
	var getClusterComplianceResponse GetClusterComplianceResponse
	path := "/api/2.0/policies/clusters/get-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getClusterComplianceResponse)
	return &getClusterComplianceResponse, err
}

// List cluster policy compliance.
//
// Returns the policy compliance status of all clusters that use a given policy.
// Clusters could be out of compliance if their policy was updated after the
// cluster was last edited.
func (a *policyComplianceForClustersImpl) ListCompliance(ctx context.Context, request ListClusterCompliancesRequest) listing.Iterator[ClusterCompliance] {

	getNextPage := func(ctx context.Context, req ListClusterCompliancesRequest) (*ListClusterCompliancesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalListCompliance(ctx, req)
	}
	getItems := func(resp *ListClusterCompliancesResponse) []ClusterCompliance {
		return resp.Clusters
	}
	getNextReq := func(resp *ListClusterCompliancesResponse) *ListClusterCompliancesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List cluster policy compliance.
//
// Returns the policy compliance status of all clusters that use a given policy.
// Clusters could be out of compliance if their policy was updated after the
// cluster was last edited.
func (a *policyComplianceForClustersImpl) ListComplianceAll(ctx context.Context, request ListClusterCompliancesRequest) ([]ClusterCompliance, error) {
	iterator := a.ListCompliance(ctx, request)
	return listing.ToSlice[ClusterCompliance](ctx, iterator)
}

func (a *policyComplianceForClustersImpl) internalListCompliance(ctx context.Context, request ListClusterCompliancesRequest) (*ListClusterCompliancesResponse, error) {
	var listClusterCompliancesResponse ListClusterCompliancesResponse
	path := "/api/2.0/policies/clusters/list-compliance"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listClusterCompliancesResponse)
	return &listClusterCompliancesResponse, err
}

// unexported type that holds implementations of just PolicyFamilies API methods
type policyFamiliesImpl struct {
	client *client.DatabricksClient
}

func (a *policyFamiliesImpl) Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error) {
	var policyFamily PolicyFamily
	path := fmt.Sprintf("/api/2.0/policy-families/%v", request.PolicyFamilyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &policyFamily)
	return &policyFamily, err
}

// List policy families.
//
// Returns the list of policy definition types available to use at their latest
// version. This API is paginated.
func (a *policyFamiliesImpl) List(ctx context.Context, request ListPolicyFamiliesRequest) listing.Iterator[PolicyFamily] {

	getNextPage := func(ctx context.Context, req ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListPolicyFamiliesResponse) []PolicyFamily {
		return resp.PolicyFamilies
	}
	getNextReq := func(resp *ListPolicyFamiliesResponse) *ListPolicyFamiliesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List policy families.
//
// Returns the list of policy definition types available to use at their latest
// version. This API is paginated.
func (a *policyFamiliesImpl) ListAll(ctx context.Context, request ListPolicyFamiliesRequest) ([]PolicyFamily, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[PolicyFamily](ctx, iterator)
}

func (a *policyFamiliesImpl) internalList(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error) {
	var listPolicyFamiliesResponse ListPolicyFamiliesResponse
	path := "/api/2.0/policy-families"
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listPolicyFamiliesResponse)
	return &listPolicyFamiliesResponse, err
}
