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
	err := a.client.Do(ctx, http.MethodPost, path, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

func (a *clusterPoliciesImpl) Delete(ctx context.Context, request DeletePolicy) error {
	path := "/api/2.0/policies/clusters/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clusterPoliciesImpl) Edit(ctx context.Context, request EditPolicy) error {
	path := "/api/2.0/policies/clusters/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clusterPoliciesImpl) Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error) {
	var policy Policy
	path := "/api/2.0/policies/clusters/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &policy)
	return &policy, err
}

func (a *clusterPoliciesImpl) List(ctx context.Context, request ListClusterPoliciesRequest) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listPoliciesResponse)
	return &listPoliciesResponse, err
}

// unexported type that holds implementations of just Clusters API methods
type clustersImpl struct {
	client *client.DatabricksClient
}

func (a *clustersImpl) ChangeOwner(ctx context.Context, request ChangeClusterOwner) error {
	path := "/api/2.0/clusters/change-owner"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error) {
	var createClusterResponse CreateClusterResponse
	path := "/api/2.0/clusters/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createClusterResponse)
	return &createClusterResponse, err
}

func (a *clustersImpl) Delete(ctx context.Context, request DeleteCluster) error {
	path := "/api/2.0/clusters/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Edit(ctx context.Context, request EditCluster) error {
	path := "/api/2.0/clusters/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error) {
	var getEventsResponse GetEventsResponse
	path := "/api/2.0/clusters/events"
	err := a.client.Do(ctx, http.MethodPost, path, request, &getEventsResponse)
	return &getEventsResponse, err
}

func (a *clustersImpl) Get(ctx context.Context, request GetClusterRequest) (*ClusterInfo, error) {
	var clusterInfo ClusterInfo
	path := "/api/2.0/clusters/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterInfo)
	return &clusterInfo, err
}

func (a *clustersImpl) List(ctx context.Context, request ListClustersRequest) (*ListClustersResponse, error) {
	var listClustersResponse ListClustersResponse
	path := "/api/2.0/clusters/list"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listClustersResponse)
	return &listClustersResponse, err
}

func (a *clustersImpl) ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error) {
	var listNodeTypesResponse ListNodeTypesResponse
	path := "/api/2.0/clusters/list-node-types"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listNodeTypesResponse)
	return &listNodeTypesResponse, err
}

func (a *clustersImpl) ListZones(ctx context.Context) (*ListAvailableZonesResponse, error) {
	var listAvailableZonesResponse ListAvailableZonesResponse
	path := "/api/2.0/clusters/list-zones"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listAvailableZonesResponse)
	return &listAvailableZonesResponse, err
}

func (a *clustersImpl) PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error {
	path := "/api/2.0/clusters/permanent-delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Pin(ctx context.Context, request PinCluster) error {
	path := "/api/2.0/clusters/pin"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Resize(ctx context.Context, request ResizeCluster) error {
	path := "/api/2.0/clusters/resize"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Restart(ctx context.Context, request RestartCluster) error {
	path := "/api/2.0/clusters/restart"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error) {
	var getSparkVersionsResponse GetSparkVersionsResponse
	path := "/api/2.0/clusters/spark-versions"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getSparkVersionsResponse)
	return &getSparkVersionsResponse, err
}

func (a *clustersImpl) Start(ctx context.Context, request StartCluster) error {
	path := "/api/2.0/clusters/start"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *clustersImpl) Unpin(ctx context.Context, request UnpinCluster) error {
	path := "/api/2.0/clusters/unpin"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just CommandExecution API methods
type commandExecutionImpl struct {
	client *client.DatabricksClient
}

func (a *commandExecutionImpl) Cancel(ctx context.Context, request CancelCommand) error {
	path := "/api/1.2/commands/cancel"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *commandExecutionImpl) CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error) {
	var commandStatusResponse CommandStatusResponse
	path := "/api/1.2/commands/status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &commandStatusResponse)
	return &commandStatusResponse, err
}

func (a *commandExecutionImpl) ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error) {
	var contextStatusResponse ContextStatusResponse
	path := "/api/1.2/contexts/status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &contextStatusResponse)
	return &contextStatusResponse, err
}

func (a *commandExecutionImpl) Create(ctx context.Context, request CreateContext) (*Created, error) {
	var created Created
	path := "/api/1.2/contexts/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &created)
	return &created, err
}

func (a *commandExecutionImpl) Destroy(ctx context.Context, request DestroyContext) error {
	path := "/api/1.2/contexts/destroy"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *commandExecutionImpl) Execute(ctx context.Context, request Command) (*Created, error) {
	var created Created
	path := "/api/1.2/commands/execute"
	err := a.client.Do(ctx, http.MethodPost, path, request, &created)
	return &created, err
}

// unexported type that holds implementations of just GlobalInitScripts API methods
type globalInitScriptsImpl struct {
	client *client.DatabricksClient
}

func (a *globalInitScriptsImpl) Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error) {
	var createResponse CreateResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createResponse)
	return &createResponse, err
}

func (a *globalInitScriptsImpl) Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *globalInitScriptsImpl) Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error) {
	var globalInitScriptDetailsWithContent GlobalInitScriptDetailsWithContent
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &globalInitScriptDetailsWithContent)
	return &globalInitScriptDetailsWithContent, err
}

func (a *globalInitScriptsImpl) List(ctx context.Context) (*ListGlobalInitScriptsResponse, error) {
	var listGlobalInitScriptsResponse ListGlobalInitScriptsResponse
	path := "/api/2.0/global-init-scripts"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listGlobalInitScriptsResponse)
	return &listGlobalInitScriptsResponse, err
}

func (a *globalInitScriptsImpl) Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error {
	path := fmt.Sprintf("/api/2.0/global-init-scripts/%v", request.ScriptId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just InstancePools API methods
type instancePoolsImpl struct {
	client *client.DatabricksClient
}

func (a *instancePoolsImpl) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Do(ctx, http.MethodPost, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

func (a *instancePoolsImpl) Delete(ctx context.Context, request DeleteInstancePool) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instancePoolsImpl) Edit(ctx context.Context, request EditInstancePool) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instancePoolsImpl) Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &getInstancePool)
	return &getInstancePool, err
}

func (a *instancePoolsImpl) List(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listInstancePools)
	return &listInstancePools, err
}

// unexported type that holds implementations of just InstanceProfiles API methods
type instanceProfilesImpl struct {
	client *client.DatabricksClient
}

func (a *instanceProfilesImpl) Add(ctx context.Context, request AddInstanceProfile) error {
	path := "/api/2.0/instance-profiles/add"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instanceProfilesImpl) Edit(ctx context.Context, request InstanceProfile) error {
	path := "/api/2.0/instance-profiles/edit"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *instanceProfilesImpl) List(ctx context.Context) (*ListInstanceProfilesResponse, error) {
	var listInstanceProfilesResponse ListInstanceProfilesResponse
	path := "/api/2.0/instance-profiles/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listInstanceProfilesResponse)
	return &listInstanceProfilesResponse, err
}

func (a *instanceProfilesImpl) Remove(ctx context.Context, request RemoveInstanceProfile) error {
	path := "/api/2.0/instance-profiles/remove"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just Libraries API methods
type librariesImpl struct {
	client *client.DatabricksClient
}

func (a *librariesImpl) AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error) {
	var listAllClusterLibraryStatusesResponse ListAllClusterLibraryStatusesResponse
	path := "/api/2.0/libraries/all-cluster-statuses"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listAllClusterLibraryStatusesResponse)
	return &listAllClusterLibraryStatusesResponse, err
}

func (a *librariesImpl) ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterLibraryStatuses, error) {
	var clusterLibraryStatuses ClusterLibraryStatuses
	path := "/api/2.0/libraries/cluster-status"
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterLibraryStatuses)
	return &clusterLibraryStatuses, err
}

func (a *librariesImpl) Install(ctx context.Context, request InstallLibraries) error {
	path := "/api/2.0/libraries/install"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *librariesImpl) Uninstall(ctx context.Context, request UninstallLibraries) error {
	path := "/api/2.0/libraries/uninstall"
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

// unexported type that holds implementations of just PolicyFamilies API methods
type policyFamiliesImpl struct {
	client *client.DatabricksClient
}

func (a *policyFamiliesImpl) Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error) {
	var policyFamily PolicyFamily
	path := fmt.Sprintf("/api/2.0/policy-families/%v", request.PolicyFamilyId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &policyFamily)
	return &policyFamily, err
}

func (a *policyFamiliesImpl) List(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error) {
	var listPolicyFamiliesResponse ListPolicyFamiliesResponse
	path := "/api/2.0/policy-families"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listPolicyFamiliesResponse)
	return &listPolicyFamiliesResponse, err
}
