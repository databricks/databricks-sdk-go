// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just ClusterPolicies API methods
type clusterPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *clusterPoliciesImpl) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {
	var createPolicyResponse CreatePolicyResponse
	path := "/api/2.0/policies/clusters/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

func (a *clusterPoliciesImpl) Delete(ctx context.Context, request DeletePolicy) error {
	var deletePolicyResponse DeletePolicyResponse
	path := "/api/2.0/policies/clusters/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deletePolicyResponse)
	return err
}

func (a *clusterPoliciesImpl) Edit(ctx context.Context, request EditPolicy) error {
	var editPolicyResponse EditPolicyResponse
	path := "/api/2.0/policies/clusters/edit"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &editPolicyResponse)
	return err
}

func (a *clusterPoliciesImpl) Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error) {
	var policy Policy
	path := "/api/2.0/policies/clusters/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &policy)
	return &policy, err
}

func (a *clusterPoliciesImpl) GetPermissionLevels(ctx context.Context, request GetClusterPolicyPermissionLevelsRequest) (*GetClusterPolicyPermissionLevelsResponse, error) {
	var getClusterPolicyPermissionLevelsResponse GetClusterPolicyPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v/permissionLevels", request.ClusterPolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getClusterPolicyPermissionLevelsResponse)
	return &getClusterPolicyPermissionLevelsResponse, err
}

func (a *clusterPoliciesImpl) GetPermissions(ctx context.Context, request GetClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	var clusterPolicyPermissions ClusterPolicyPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &clusterPolicyPermissions)
	return &clusterPolicyPermissions, err
}

func (a *clusterPoliciesImpl) List(ctx context.Context, request ListClusterPoliciesRequest) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listPoliciesResponse)
	return &listPoliciesResponse, err
}

func (a *clusterPoliciesImpl) SetPermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	var clusterPolicyPermissions ClusterPolicyPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &clusterPolicyPermissions)
	return &clusterPolicyPermissions, err
}

func (a *clusterPoliciesImpl) UpdatePermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error) {
	var clusterPolicyPermissions ClusterPolicyPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &clusterPolicyPermissions)
	return &clusterPolicyPermissions, err
}

// unexported type that holds implementations of just Clusters API methods
type clustersImpl struct {
	client *client.DatabricksClient
}

func (a *clustersImpl) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	var changeClusterOwnerResponse ChangeClusterOwnerResponse
	path := "/api/2.1/clusters/change-owner"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &changeClusterOwnerResponse)
	return err
}

func (a *clustersImpl) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	var createClusterResponse CreateClusterResponse
	path := "/api/2.1/clusters/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createClusterResponse)
	return &createClusterResponse, err
}

func (a *clustersImpl) Delete(ctx context.Context, request DeleteCluster) error {
	var deleteClusterResponse DeleteClusterResponse
	path := "/api/2.1/clusters/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteClusterResponse)
	return err
}

func (a *clustersImpl) Edit(ctx context.Context, request EditCluster) error {
	var editClusterResponse EditClusterResponse
	path := "/api/2.1/clusters/edit"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &editClusterResponse)
	return err
}

func (a *clustersImpl) Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {
	var getEventsResponse GetEventsResponse
	path := "/api/2.1/clusters/events"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &getEventsResponse)
	return &getEventsResponse, err
}

func (a *clustersImpl) Get(ctx context.Context, request GetClusterRequest) (*ClusterDetails, error) {
	var clusterDetails ClusterDetails
	path := "/api/2.1/clusters/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &clusterDetails)
	return &clusterDetails, err
}

func (a *clustersImpl) GetPermissionLevels(ctx context.Context, request GetClusterPermissionLevelsRequest) (*GetClusterPermissionLevelsResponse, error) {
	var getClusterPermissionLevelsResponse GetClusterPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v/permissionLevels", request.ClusterId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getClusterPermissionLevelsResponse)
	return &getClusterPermissionLevelsResponse, err
}

func (a *clustersImpl) GetPermissions(ctx context.Context, request GetClusterPermissionsRequest) (*ClusterPermissions, error) {
	var clusterPermissions ClusterPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &clusterPermissions)
	return &clusterPermissions, err
}

func (a *clustersImpl) List(ctx context.Context, request ListClustersRequest) (*ListClustersResponse, error) {
	var listClustersResponse ListClustersResponse
	path := "/api/2.1/clusters/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listClustersResponse)
	return &listClustersResponse, err
}

func (a *clustersImpl) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	var listNodeTypesResponse ListNodeTypesResponse
	path := "/api/2.1/clusters/list-node-types"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listNodeTypesResponse)
	return &listNodeTypesResponse, err
}

func (a *clustersImpl) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	var listAvailableZonesResponse ListAvailableZonesResponse
	path := "/api/2.1/clusters/list-zones"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listAvailableZonesResponse)
	return &listAvailableZonesResponse, err
}

func (a *clustersImpl) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	var permanentDeleteClusterResponse PermanentDeleteClusterResponse
	path := "/api/2.1/clusters/permanent-delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &permanentDeleteClusterResponse)
	return err
}

func (a *clustersImpl) Pin(ctx context.Context, request PinCluster) error {
	var pinClusterResponse PinClusterResponse
	path := "/api/2.1/clusters/pin"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &pinClusterResponse)
	return err
}

func (a *clustersImpl) Resize(ctx context.Context, request ResizeCluster) error {
	var resizeClusterResponse ResizeClusterResponse
	path := "/api/2.1/clusters/resize"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &resizeClusterResponse)
	return err
}

func (a *clustersImpl) Restart(ctx context.Context, request RestartCluster) error {
	var restartClusterResponse RestartClusterResponse
	path := "/api/2.1/clusters/restart"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &restartClusterResponse)
	return err
}

func (a *clustersImpl) SetPermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {
	var clusterPermissions ClusterPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &clusterPermissions)
	return &clusterPermissions, err
}

func (a *clustersImpl) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	var getSparkVersionsResponse GetSparkVersionsResponse
	path := "/api/2.1/clusters/spark-versions"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &getSparkVersionsResponse)
	return &getSparkVersionsResponse, err
}

func (a *clustersImpl) Start(ctx context.Context, request StartCluster) error {
	var startClusterResponse StartClusterResponse
	path := "/api/2.1/clusters/start"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &startClusterResponse)
	return err
}

func (a *clustersImpl) Unpin(ctx context.Context, request UnpinCluster) error {
	var unpinClusterResponse UnpinClusterResponse
	path := "/api/2.1/clusters/unpin"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &unpinClusterResponse)
	return err
}

func (a *clustersImpl) Update(ctx context.Context, request UpdateCluster) error {
	var updateClusterResponse UpdateClusterResponse
	path := "/api/2.1/clusters/update"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &updateClusterResponse)
	return err
}

func (a *clustersImpl) UpdatePermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error) {
	var clusterPermissions ClusterPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &clusterPermissions)
	return &clusterPermissions, err
}

// unexported type that holds implementations of just CommandExecution API methods
type commandExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *commandExecutionImpl) Cancel(ctx context.Context, request CancelCommand) error {
	var cancelResponse CancelResponse
	path := "/api/1.2/commands/cancel"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &cancelResponse)
	return err
}

func (a *commandExecutionImpl) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	var commandStatusResponse CommandStatusResponse
	path := "/api/1.2/commands/status"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &commandStatusResponse)
	return &commandStatusResponse, err
}

func (a *commandExecutionImpl) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	var contextStatusResponse ContextStatusResponse
	path := "/api/1.2/contexts/status"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &contextStatusResponse)
	return &contextStatusResponse, err
}

func (a *commandExecutionImpl) Create(ctx context.Context, request CreateContext) (*Created, error) {
	var created Created
	path := "/api/1.2/contexts/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &created)
	return &created, err
}

func (a *commandExecutionImpl) Destroy(ctx context.Context, request DestroyContext) error {
	var destroyResponse DestroyResponse
	path := "/api/1.2/contexts/destroy"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &destroyResponse)
	return err
}

func (a *commandExecutionImpl) Execute(ctx context.Context, request Command) (*Created, error) {
	var created Created
	path := "/api/1.2/commands/execute"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &created)
	return &created, err
}

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsImpl) Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/global-init-scripts"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createResponse)
	return &createResponse, err
}

func (a *globalInitScriptsImpl) Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	headers := make(map[string]string)
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteResponse)
	return err
}

func (a *globalInitScriptsImpl) Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

func (a *globalInitScriptsImpl) List(ctx context.Context) (*ListGlobalInitScriptsResponse, error) {
	var listGlobalInitScriptsResponse ListGlobalInitScriptsResponse
	path := "/api/2.0/global-init-scripts"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listGlobalInitScriptsResponse)
	return &listGlobalInitScriptsResponse, err
}

func (a *globalInitScriptsImpl) Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &updateResponse)
	return err
}

// unexported type that holds implementations of just InstancePools API methods
type instancePoolsImpl struct {
	client *client.DatabricksClient
}

func (a *instancePoolsImpl) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

func (a *instancePoolsImpl) Delete(ctx context.Context, request DeleteInstancePool) error {
	var deleteInstancePoolResponse DeleteInstancePoolResponse
	path := "/api/2.0/instance-pools/delete"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &deleteInstancePoolResponse)
	return err
}

func (a *instancePoolsImpl) Edit(ctx context.Context, request EditInstancePool) error {
	var editInstancePoolResponse EditInstancePoolResponse
	path := "/api/2.0/instance-pools/edit"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &editInstancePoolResponse)
	return err
}

func (a *instancePoolsImpl) Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getInstancePool)
	return &getInstancePool, err
}

func (a *instancePoolsImpl) GetPermissionLevels(ctx context.Context, request GetInstancePoolPermissionLevelsRequest) (*GetInstancePoolPermissionLevelsResponse, error) {
	var getInstancePoolPermissionLevelsResponse GetInstancePoolPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v/permissionLevels", request.InstancePoolId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getInstancePoolPermissionLevelsResponse)
	return &getInstancePoolPermissionLevelsResponse, err
}

func (a *instancePoolsImpl) GetPermissions(ctx context.Context, request GetInstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	var instancePoolPermissions InstancePoolPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &instancePoolPermissions)
	return &instancePoolPermissions, err
}

func (a *instancePoolsImpl) List(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listInstancePools)
	return &listInstancePools, err
}

func (a *instancePoolsImpl) SetPermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	var instancePoolPermissions InstancePoolPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &instancePoolPermissions)
	return &instancePoolPermissions, err
}

func (a *instancePoolsImpl) UpdatePermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error) {
	var instancePoolPermissions InstancePoolPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &instancePoolPermissions)
	return &instancePoolPermissions, err
}

// unexported type that holds implementations of just InstanceProfiles API methods
type instanceProfilesImpl struct {
	client *client.DatabricksClient
}

func (a *instanceProfilesImpl) Add(ctx context.Context, request AddInstanceProfile) error {
	var addResponse AddResponse
	path := "/api/2.0/instance-profiles/add"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &addResponse)
	return err
}

func (a *instanceProfilesImpl) Edit(ctx context.Context, request InstanceProfile) error {
	var editResponse EditResponse
	path := "/api/2.0/instance-profiles/edit"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &editResponse)
	return err
}

func (a *instanceProfilesImpl) List(ctx context.Context) (*ListInstanceProfilesResponse, error) {
	var listInstanceProfilesResponse ListInstanceProfilesResponse
	path := "/api/2.0/instance-profiles/list"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listInstanceProfilesResponse)
	return &listInstanceProfilesResponse, err
}

func (a *instanceProfilesImpl) Remove(ctx context.Context, request RemoveInstanceProfile) error {
	var removeResponse RemoveResponse
	path := "/api/2.0/instance-profiles/remove"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &removeResponse)
	return err
}

// unexported type that holds implementations of just Libraries API methods
type librariesImpl struct {
	client *client.DatabricksClient
}

func (a *librariesImpl) AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	var listAllClusterLibraryStatusesResponse ListAllClusterLibraryStatusesResponse
	path := "/api/2.0/libraries/all-cluster-statuses"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &listAllClusterLibraryStatusesResponse)
	return &listAllClusterLibraryStatusesResponse, err
}

func (a *librariesImpl) ClusterStatus(ctx context.Context, request ClusterStatus) (*ClusterLibraryStatuses, error) {
	var clusterLibraryStatuses ClusterLibraryStatuses
	path := "/api/2.0/libraries/cluster-status"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &clusterLibraryStatuses)
	return &clusterLibraryStatuses, err
}

func (a *librariesImpl) Install(ctx context.Context, request InstallLibraries) error {
	var installLibrariesResponse InstallLibrariesResponse
	path := "/api/2.0/libraries/install"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &installLibrariesResponse)
	return err
}

func (a *librariesImpl) Uninstall(ctx context.Context, request UninstallLibraries) error {
	var uninstallLibrariesResponse UninstallLibrariesResponse
	path := "/api/2.0/libraries/uninstall"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &uninstallLibrariesResponse)
	return err
}

// unexported type that holds implementations of just PolicyComplianceForClusters API methods
type policyComplianceForClustersImpl struct {
	client *client.DatabricksClient
}

func (a *policyComplianceForClustersImpl) EnforceCompliance(ctx context.Context, request EnforceClusterComplianceRequest) (*EnforceClusterComplianceResponse, error) {
	var enforceClusterComplianceResponse EnforceClusterComplianceResponse
	path := "/api/2.0/policies/clusters/enforce-compliance"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &enforceClusterComplianceResponse)
	return &enforceClusterComplianceResponse, err
}

func (a *policyComplianceForClustersImpl) GetCompliance(ctx context.Context, request GetClusterComplianceRequest) (*GetClusterComplianceResponse, error) {
	var getClusterComplianceResponse GetClusterComplianceResponse
	path := "/api/2.0/policies/clusters/get-compliance"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getClusterComplianceResponse)
	return &getClusterComplianceResponse, err
}

func (a *policyComplianceForClustersImpl) ListCompliance(ctx context.Context, request ListClusterCompliancesRequest) (*ListClusterCompliancesResponse, error) {
	var listClusterCompliancesResponse ListClusterCompliancesResponse
	path := "/api/2.0/policies/clusters/list-compliance"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listClusterCompliancesResponse)
	return &listClusterCompliancesResponse, err
}

// unexported type that holds implementations of just PolicyFamilies API methods
type policyFamiliesImpl struct {
	client *client.DatabricksClient
}

func (a *policyFamiliesImpl) Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error) {
	var policyFamily PolicyFamily
	path := fmt.Sprintf("/api/2.0/policy-families/%v", request.PolicyFamilyId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &policyFamily)
	return &policyFamily, err
}

func (a *policyFamiliesImpl) List(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error) {
	var listPolicyFamiliesResponse ListPolicyFamiliesResponse
	path := "/api/2.0/policy-families"
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listPolicyFamiliesResponse)
	return &listPolicyFamiliesResponse, err
}
