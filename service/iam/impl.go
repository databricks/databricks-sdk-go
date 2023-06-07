// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just AccountAccessControl API methods
type accountAccessControlImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/assignable-roles", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/access-control/rule-sets", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPut, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountAccessControlShardProxy API methods
type accountAccessControlShardProxyImpl struct {
	client *client.DatabricksClient
}

func (a *accountAccessControlShardProxyImpl) GetAssignableRolesForResource(ctx context.Context, request GetAssignableRolesForResourceRequest) (*GetAssignableRolesForResourceResponse, error) {
	var getAssignableRolesForResourceResponse GetAssignableRolesForResourceResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/access-control/assignable-roles")
	err := a.client.Do(ctx, http.MethodGet, path, request, &getAssignableRolesForResourceResponse)
	return &getAssignableRolesForResourceResponse, err
}

func (a *accountAccessControlShardProxyImpl) GetRuleSet(ctx context.Context, request GetRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets")
	err := a.client.Do(ctx, http.MethodGet, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

func (a *accountAccessControlShardProxyImpl) UpdateRuleSet(ctx context.Context, request UpdateRuleSetRequest) (*RuleSetResponse, error) {
	var ruleSetResponse RuleSetResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/access-control/rule-sets")
	err := a.client.Do(ctx, http.MethodPut, path, request, &ruleSetResponse)
	return &ruleSetResponse, err
}

// unexported type that holds implementations of just AccountGroups API methods
type accountGroupsImpl struct {
	client *client.DatabricksClient
}

func (a *accountGroupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &group)
	return &group, err
}

func (a *accountGroupsImpl) Delete(ctx context.Context, request DeleteAccountGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountGroupsImpl) Get(ctx context.Context, request GetAccountGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &group)
	return &group, err
}

func (a *accountGroupsImpl) List(ctx context.Context, request ListAccountGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *accountGroupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *accountGroupsImpl) Update(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Groups/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just AccountServicePrincipals API methods
type accountServicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *accountServicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsImpl) Delete(ctx context.Context, request DeleteAccountServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountServicePrincipalsImpl) Get(ctx context.Context, request GetAccountServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *accountServicePrincipalsImpl) List(ctx context.Context, request ListAccountServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *accountServicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *accountServicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/ServicePrincipals/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just AccountUsers API methods
type accountUsersImpl struct {
	client *client.DatabricksClient
}

func (a *accountUsersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &user)
	return &user, err
}

func (a *accountUsersImpl) Delete(ctx context.Context, request DeleteAccountUserRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *accountUsersImpl) Get(ctx context.Context, request GetAccountUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &user)
	return &user, err
}

func (a *accountUsersImpl) List(ctx context.Context, request ListAccountUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *accountUsersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *accountUsersImpl) Update(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/scim/v2/Users/%v", a.client.ConfiguredAccountID(), request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just CurrentUser API methods
type currentUserImpl struct {
	client *client.DatabricksClient
}

func (a *currentUserImpl) Me(ctx context.Context) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Me"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &user)
	return &user, err
}

// unexported type that holds implementations of just Groups API methods
type groupsImpl struct {
	client *client.DatabricksClient
}

func (a *groupsImpl) Create(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Do(ctx, http.MethodPost, path, request, &group)
	return &group, err
}

func (a *groupsImpl) Delete(ctx context.Context, request DeleteGroupRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *groupsImpl) Get(ctx context.Context, request GetGroupRequest) (*Group, error) {
	var group Group
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &group)
	return &group, err
}

func (a *groupsImpl) List(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

func (a *groupsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *groupsImpl) Update(ctx context.Context, request Group) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Groups/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just PermissionsClusterPolicies API methods
type permissionsClusterPoliciesImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsClusterPoliciesImpl) Get(ctx context.Context, request GetPermissionsClusterPolicyRequest) (*ClusterPolicyObjectPermissions, error) {
	var clusterPolicyObjectPermissions ClusterPolicyObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterPolicyObjectPermissions)
	return &clusterPolicyObjectPermissions, err
}

func (a *permissionsClusterPoliciesImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*ClusterPolicyGetPermissionLevelsResponse, error) {
	var clusterPolicyGetPermissionLevelsResponse ClusterPolicyGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v/permissionLevels", request.ClusterPolicyId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterPolicyGetPermissionLevelsResponse)
	return &clusterPolicyGetPermissionLevelsResponse, err
}

func (a *permissionsClusterPoliciesImpl) Set(ctx context.Context, request ClusterPolicyPermissionsRequest) error {
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *permissionsClusterPoliciesImpl) Update(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyObjectPermissions, error) {
	var clusterPolicyObjectPermissions ClusterPolicyObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/cluster-policies/%v", request.ClusterPolicyId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &clusterPolicyObjectPermissions)
	return &clusterPolicyObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsClusters API methods
type permissionsClustersImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsClustersImpl) Get(ctx context.Context, request GetPermissionsClusterRequest) (*ClusterObjectPermissions, error) {
	var clusterObjectPermissions ClusterObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterObjectPermissions)
	return &clusterObjectPermissions, err
}

func (a *permissionsClustersImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*ClusterGetPermissionLevelsResponse, error) {
	var clusterGetPermissionLevelsResponse ClusterGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v/permissionLevels", request.ClusterId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &clusterGetPermissionLevelsResponse)
	return &clusterGetPermissionLevelsResponse, err
}

func (a *permissionsClustersImpl) Set(ctx context.Context, request ClusterPermissionsRequest) (*ClusterObjectPermissions, error) {
	var clusterObjectPermissions ClusterObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &clusterObjectPermissions)
	return &clusterObjectPermissions, err
}

func (a *permissionsClustersImpl) Update(ctx context.Context, request ClusterPermissionsRequest) (*ClusterObjectPermissions, error) {
	var clusterObjectPermissions ClusterObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/clusters/%v", request.ClusterId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &clusterObjectPermissions)
	return &clusterObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsDeltaLiveTables API methods
type permissionsDeltaLiveTablesImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsDeltaLiveTablesImpl) Get(ctx context.Context, request GetPermissionsDeltaLiveTableRequest) (*DeltaLiveTableObjectPermissions, error) {
	var deltaLiveTableObjectPermissions DeltaLiveTableObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &deltaLiveTableObjectPermissions)
	return &deltaLiveTableObjectPermissions, err
}

func (a *permissionsDeltaLiveTablesImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*DeltaLiveTableGetPermissionLevelsResponse, error) {
	var deltaLiveTableGetPermissionLevelsResponse DeltaLiveTableGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v/permissionLevels", request.PipelineId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &deltaLiveTableGetPermissionLevelsResponse)
	return &deltaLiveTableGetPermissionLevelsResponse, err
}

func (a *permissionsDeltaLiveTablesImpl) Set(ctx context.Context, request DeltaLiveTablePermissionsRequest) (*DeltaLiveTableObjectPermissions, error) {
	var deltaLiveTableObjectPermissions DeltaLiveTableObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &deltaLiveTableObjectPermissions)
	return &deltaLiveTableObjectPermissions, err
}

func (a *permissionsDeltaLiveTablesImpl) Update(ctx context.Context, request DeltaLiveTablePermissionsRequest) (*DeltaLiveTableObjectPermissions, error) {
	var deltaLiveTableObjectPermissions DeltaLiveTableObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/pipelines/%v", request.PipelineId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &deltaLiveTableObjectPermissions)
	return &deltaLiveTableObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsDirectories API methods
type permissionsDirectoriesImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsDirectoriesImpl) Get(ctx context.Context, request GetPermissionsDirectoryRequest) (*DirectoryObjectPermissions, error) {
	var directoryObjectPermissions DirectoryObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/directories/%v", request.DirectoryId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &directoryObjectPermissions)
	return &directoryObjectPermissions, err
}

func (a *permissionsDirectoriesImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*DirectoryGetPermissionLevelsResponse, error) {
	var directoryGetPermissionLevelsResponse DirectoryGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/directory/%v/permissionLevels", request.DirectoryId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &directoryGetPermissionLevelsResponse)
	return &directoryGetPermissionLevelsResponse, err
}

func (a *permissionsDirectoriesImpl) Set(ctx context.Context, request DirectoryPermissionsRequest) (*DirectoryObjectPermissions, error) {
	var directoryObjectPermissions DirectoryObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/directories/")
	err := a.client.Do(ctx, http.MethodPut, path, request, &directoryObjectPermissions)
	return &directoryObjectPermissions, err
}

func (a *permissionsDirectoriesImpl) Update(ctx context.Context, request DirectoryPermissionsRequest) (*DirectoryObjectPermissions, error) {
	var directoryObjectPermissions DirectoryObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/directories/")
	err := a.client.Do(ctx, http.MethodPatch, path, request, &directoryObjectPermissions)
	return &directoryObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsJobs API methods
type permissionsJobsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsJobsImpl) Get(ctx context.Context, request GetPermissionsJobRequest) (*JobObjectPermissions, error) {
	var jobObjectPermissions JobObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &jobObjectPermissions)
	return &jobObjectPermissions, err
}

func (a *permissionsJobsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*JobGetPermissionLevelsResponse, error) {
	var jobGetPermissionLevelsResponse JobGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v/permissionLevels", request.JobId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &jobGetPermissionLevelsResponse)
	return &jobGetPermissionLevelsResponse, err
}

func (a *permissionsJobsImpl) Set(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	var jobObjectPermissions JobObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &jobObjectPermissions)
	return &jobObjectPermissions, err
}

func (a *permissionsJobsImpl) Update(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	var jobObjectPermissions JobObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/jobs/%v", request.JobId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &jobObjectPermissions)
	return &jobObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsMLflowExperiments API methods
type permissionsMLflowExperimentsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsMLflowExperimentsImpl) Get(ctx context.Context, request GetPermissionsMLflowExperimentRequest) (*MLflowExperimentObjectPermissions, error) {
	var mLflowExperimentObjectPermissions MLflowExperimentObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &mLflowExperimentObjectPermissions)
	return &mLflowExperimentObjectPermissions, err
}

func (a *permissionsMLflowExperimentsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*MLflowExperimentGetPermissionLevelsResponse, error) {
	var mLflowExperimentGetPermissionLevelsResponse MLflowExperimentGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v/permissionLevels", request.ExperimentId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &mLflowExperimentGetPermissionLevelsResponse)
	return &mLflowExperimentGetPermissionLevelsResponse, err
}

func (a *permissionsMLflowExperimentsImpl) Set(ctx context.Context, request MLflowExperimentPermissionsRequest) (*MLflowExperimentObjectPermissions, error) {
	var mLflowExperimentObjectPermissions MLflowExperimentObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &mLflowExperimentObjectPermissions)
	return &mLflowExperimentObjectPermissions, err
}

func (a *permissionsMLflowExperimentsImpl) Update(ctx context.Context, request MLflowExperimentPermissionsRequest) (*MLflowExperimentObjectPermissions, error) {
	var mLflowExperimentObjectPermissions MLflowExperimentObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/experiments/%v", request.ExperimentId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &mLflowExperimentObjectPermissions)
	return &mLflowExperimentObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsMLFlowRegisteredModels API methods
type permissionsMlFlowRegisteredModelsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsMlFlowRegisteredModelsImpl) Get(ctx context.Context, request GetPermissionsMlFlowRegisteredModelRequest) (*MLflowRegisteredModelObjectPermissions, error) {
	var mLflowRegisteredModelObjectPermissions MLflowRegisteredModelObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &mLflowRegisteredModelObjectPermissions)
	return &mLflowRegisteredModelObjectPermissions, err
}

func (a *permissionsMlFlowRegisteredModelsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*MLflowRegisteredModelGetPermissionLevelsResponse, error) {
	var mLflowRegisteredModelGetPermissionLevelsResponse MLflowRegisteredModelGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v/permissionLevels", request.RegisteredModelId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &mLflowRegisteredModelGetPermissionLevelsResponse)
	return &mLflowRegisteredModelGetPermissionLevelsResponse, err
}

func (a *permissionsMlFlowRegisteredModelsImpl) Set(ctx context.Context, request MLflowRegisteredModelPermissionsRequest) (*MLflowRegisteredModelObjectPermissions, error) {
	var mLflowRegisteredModelObjectPermissions MLflowRegisteredModelObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &mLflowRegisteredModelObjectPermissions)
	return &mLflowRegisteredModelObjectPermissions, err
}

func (a *permissionsMlFlowRegisteredModelsImpl) Update(ctx context.Context, request MLflowRegisteredModelPermissionsRequest) (*MLflowRegisteredModelObjectPermissions, error) {
	var mLflowRegisteredModelObjectPermissions MLflowRegisteredModelObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/registered-models/%v", request.RegisteredModelId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &mLflowRegisteredModelObjectPermissions)
	return &mLflowRegisteredModelObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsNotebooks API methods
type permissionsNotebooksImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsNotebooksImpl) Get(ctx context.Context, request GetPermissionsNotebookRequest) (*JobObjectPermissions, error) {
	var jobObjectPermissions JobObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/notebooks/%v", request.NotebookId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &jobObjectPermissions)
	return &jobObjectPermissions, err
}

func (a *permissionsNotebooksImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*JobGetPermissionLevelsResponse, error) {
	var jobGetPermissionLevelsResponse JobGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/notebooks/%v/permissionLevels", request.NotebookId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &jobGetPermissionLevelsResponse)
	return &jobGetPermissionLevelsResponse, err
}

func (a *permissionsNotebooksImpl) Set(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	var jobObjectPermissions JobObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/notebooks/%v", request.NotebookId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &jobObjectPermissions)
	return &jobObjectPermissions, err
}

func (a *permissionsNotebooksImpl) Update(ctx context.Context, request JobPermissionsRequest) (*JobObjectPermissions, error) {
	var jobObjectPermissions JobObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/notebooks/%v", request.NotebookId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &jobObjectPermissions)
	return &jobObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsPasswords API methods
type permissionsPasswordsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsPasswordsImpl) Get(ctx context.Context) (*PasswordObjectPermissions, error) {
	var passwordObjectPermissions PasswordObjectPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &passwordObjectPermissions)
	return &passwordObjectPermissions, err
}

func (a *permissionsPasswordsImpl) GetPermissionLevels(ctx context.Context) (*PasswordGetPermissionLevelsResponse, error) {
	var passwordGetPermissionLevelsResponse PasswordGetPermissionLevelsResponse
	path := "/api/2.0/permissions/authorization/passwords/permissionLevels"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &passwordGetPermissionLevelsResponse)
	return &passwordGetPermissionLevelsResponse, err
}

func (a *permissionsPasswordsImpl) Set(ctx context.Context, request PasswordPermissionsRequest) (*PasswordObjectPermissions, error) {
	var passwordObjectPermissions PasswordObjectPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	err := a.client.Do(ctx, http.MethodPut, path, request, &passwordObjectPermissions)
	return &passwordObjectPermissions, err
}

func (a *permissionsPasswordsImpl) Update(ctx context.Context, request PasswordPermissionsRequest) (*PasswordObjectPermissions, error) {
	var passwordObjectPermissions PasswordObjectPermissions
	path := "/api/2.0/permissions/authorization/passwords"
	err := a.client.Do(ctx, http.MethodPatch, path, request, &passwordObjectPermissions)
	return &passwordObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsPools API methods
type permissionsPoolsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsPoolsImpl) Get(ctx context.Context, request GetPermissionsPoolRequest) (*PoolObjectPermissions, error) {
	var poolObjectPermissions PoolObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v", request.InstancePoolId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &poolObjectPermissions)
	return &poolObjectPermissions, err
}

func (a *permissionsPoolsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*PoolGetPermissionLevelsResponse, error) {
	var poolGetPermissionLevelsResponse PoolGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/%v/permissionLevels", request.InstancePoolId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &poolGetPermissionLevelsResponse)
	return &poolGetPermissionLevelsResponse, err
}

func (a *permissionsPoolsImpl) Set(ctx context.Context, request PoolPermissionsRequest) (*PoolObjectPermissions, error) {
	var poolObjectPermissions PoolObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/")
	err := a.client.Do(ctx, http.MethodPut, path, request, &poolObjectPermissions)
	return &poolObjectPermissions, err
}

func (a *permissionsPoolsImpl) Update(ctx context.Context, request PoolPermissionsRequest) (*PoolObjectPermissions, error) {
	var poolObjectPermissions PoolObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/instance-pools/")
	err := a.client.Do(ctx, http.MethodPatch, path, request, &poolObjectPermissions)
	return &poolObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsRepos API methods
type permissionsReposImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsReposImpl) Get(ctx context.Context, request GetPermissionsRepoRequest) (*RepoObjectPermissions, error) {
	var repoObjectPermissions RepoObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &repoObjectPermissions)
	return &repoObjectPermissions, err
}

func (a *permissionsReposImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*RepoGetPermissionLevelsResponse, error) {
	var repoGetPermissionLevelsResponse RepoGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v/permissionLevels", request.RepoId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &repoGetPermissionLevelsResponse)
	return &repoGetPermissionLevelsResponse, err
}

func (a *permissionsReposImpl) Set(ctx context.Context, request RepoPermissionsRequest) (*RepoObjectPermissions, error) {
	var repoObjectPermissions RepoObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &repoObjectPermissions)
	return &repoObjectPermissions, err
}

func (a *permissionsReposImpl) Update(ctx context.Context, request RepoPermissionsRequest) (*RepoObjectPermissions, error) {
	var repoObjectPermissions RepoObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/repos/%v", request.RepoId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &repoObjectPermissions)
	return &repoObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsSQLWarehouses API methods
type permissionsSqlWarehousesImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsSqlWarehousesImpl) Get(ctx context.Context, request GetPermissionsSqlWarehousRequest) (*SqlWarehouseObjectPermissions, error) {
	var sqlWarehouseObjectPermissions SqlWarehouseObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/sql/warehouses/%v", request.WarehouseId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &sqlWarehouseObjectPermissions)
	return &sqlWarehouseObjectPermissions, err
}

func (a *permissionsSqlWarehousesImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*SqlWarehouseGetPermissionLevelsResponse, error) {
	var sqlWarehouseGetPermissionLevelsResponse SqlWarehouseGetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/sql/warehouses/%v/permissionLevels", request.WarehouseId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &sqlWarehouseGetPermissionLevelsResponse)
	return &sqlWarehouseGetPermissionLevelsResponse, err
}

func (a *permissionsSqlWarehousesImpl) Set(ctx context.Context, request SqlWarehousePermissionsRequest) (*SqlWarehouseObjectPermissions, error) {
	var sqlWarehouseObjectPermissions SqlWarehouseObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/sql/warehouses/%v", request.WarehouseId)
	err := a.client.Do(ctx, http.MethodPut, path, request, &sqlWarehouseObjectPermissions)
	return &sqlWarehouseObjectPermissions, err
}

func (a *permissionsSqlWarehousesImpl) Update(ctx context.Context, request SqlWarehousePermissionsRequest) (*SqlWarehouseObjectPermissions, error) {
	var sqlWarehouseObjectPermissions SqlWarehouseObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/sql/warehouses/%v", request.WarehouseId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, &sqlWarehouseObjectPermissions)
	return &sqlWarehouseObjectPermissions, err
}

// unexported type that holds implementations of just PermissionsTokens API methods
type permissionsTokensImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsTokensImpl) Get(ctx context.Context) (*TokenObjectPermissions, error) {
	var tokenObjectPermissions TokenObjectPermissions
	path := "/api/2.0/permissions/authorization/tokens"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &tokenObjectPermissions)
	return &tokenObjectPermissions, err
}

func (a *permissionsTokensImpl) GetPermissionLevels(ctx context.Context) (*TokenGetPermissionLevelsResponse, error) {
	var tokenGetPermissionLevelsResponse TokenGetPermissionLevelsResponse
	path := "/api/2.0/permissions/authorization/tokens/permissionLevels"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &tokenGetPermissionLevelsResponse)
	return &tokenGetPermissionLevelsResponse, err
}

func (a *permissionsTokensImpl) Set(ctx context.Context, request TokenPermissionsRequest) (*TokenObjectPermissions, error) {
	var tokenObjectPermissions TokenObjectPermissions
	path := "/api/2.0/permissions/authorization/tokens"
	err := a.client.Do(ctx, http.MethodPut, path, request, &tokenObjectPermissions)
	return &tokenObjectPermissions, err
}

func (a *permissionsTokensImpl) Update(ctx context.Context, request TokenPermissionsRequest) (*TokenObjectPermissions, error) {
	var tokenObjectPermissions TokenObjectPermissions
	path := "/api/2.0/permissions/authorization/tokens"
	err := a.client.Do(ctx, http.MethodPatch, path, request, &tokenObjectPermissions)
	return &tokenObjectPermissions, err
}

// unexported type that holds implementations of just ServicePrincipals API methods
type servicePrincipalsImpl struct {
	client *client.DatabricksClient
}

func (a *servicePrincipalsImpl) Create(ctx context.Context, request ServicePrincipal) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Do(ctx, http.MethodPost, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsImpl) Delete(ctx context.Context, request DeleteServicePrincipalRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *servicePrincipalsImpl) Get(ctx context.Context, request GetServicePrincipalRequest) (*ServicePrincipal, error) {
	var servicePrincipal ServicePrincipal
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &servicePrincipal)
	return &servicePrincipal, err
}

func (a *servicePrincipalsImpl) List(ctx context.Context, request ListServicePrincipalsRequest) (*ListServicePrincipalResponse, error) {
	var listServicePrincipalResponse ListServicePrincipalResponse
	path := "/api/2.0/preview/scim/v2/ServicePrincipals"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listServicePrincipalResponse)
	return &listServicePrincipalResponse, err
}

func (a *servicePrincipalsImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *servicePrincipalsImpl) Update(ctx context.Context, request ServicePrincipal) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/ServicePrincipals/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just Users API methods
type usersImpl struct {
	client *client.DatabricksClient
}

func (a *usersImpl) Create(ctx context.Context, request User) (*User, error) {
	var user User
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Do(ctx, http.MethodPost, path, request, &user)
	return &user, err
}

func (a *usersImpl) Delete(ctx context.Context, request DeleteUserRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *usersImpl) Get(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodGet, path, request, &user)
	return &user, err
}

func (a *usersImpl) List(ctx context.Context, request ListUsersRequest) (*ListUsersResponse, error) {
	var listUsersResponse ListUsersResponse
	path := "/api/2.0/preview/scim/v2/Users"
	err := a.client.Do(ctx, http.MethodGet, path, request, &listUsersResponse)
	return &listUsersResponse, err
}

func (a *usersImpl) Patch(ctx context.Context, request PartialUpdate) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

func (a *usersImpl) Update(ctx context.Context, request User) error {
	path := fmt.Sprintf("/api/2.0/preview/scim/v2/Users/%v", request.Id)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentImpl) Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *workspaceAssignmentImpl) Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &workspacePermissions)
	return &workspacePermissions, err
}

func (a *workspaceAssignmentImpl) List(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}
