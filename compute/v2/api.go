// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Cluster Policies, Clusters, Command Execution, Global Init Scripts, Instance Pools, Instance Profiles, Libraries, Policy Compliance For Clusters, Policy Families, etc.
package compute

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type ClusterPoliciesAPI struct {
	clusterPoliciesImpl
}

// Delete a cluster policy.
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) DeleteByPolicyId(ctx context.Context, policyId string) (*DeletePolicyResponse, error) {
	return a.clusterPoliciesImpl.Delete(ctx, DeletePolicy{
		PolicyId: policyId,
	})
}

// Get a cluster policy.
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) GetByPolicyId(ctx context.Context, policyId string) (*Policy, error) {
	return a.clusterPoliciesImpl.Get(ctx, GetClusterPolicyRequest{
		PolicyId: policyId,
	})
}

// Get cluster policy permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ClusterPoliciesAPI) GetPermissionLevelsByClusterPolicyId(ctx context.Context, clusterPolicyId string) (*GetClusterPolicyPermissionLevelsResponse, error) {
	return a.clusterPoliciesImpl.GetPermissionLevels(ctx, GetClusterPolicyPermissionLevelsRequest{
		ClusterPolicyId: clusterPolicyId,
	})
}

// Get cluster policy permissions.
//
// Gets the permissions of a cluster policy. Cluster policies can inherit
// permissions from their root object.
func (a *ClusterPoliciesAPI) GetPermissionsByClusterPolicyId(ctx context.Context, clusterPolicyId string) (*ClusterPolicyPermissions, error) {
	return a.clusterPoliciesImpl.GetPermissions(ctx, GetClusterPolicyPermissionsRequest{
		ClusterPolicyId: clusterPolicyId,
	})
}

// PolicyNameToPolicyIdMap calls [ClusterPoliciesAPI.ListAll] and creates a map of results with [Policy].Name as key and [Policy].PolicyId as value.
//
// Returns an error if there's more than one [Policy] with the same .Name.
//
// Note: All [Policy] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) PolicyNameToPolicyIdMap(ctx context.Context, request ListClusterPoliciesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
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

type ClustersAPI struct {
	clustersImpl
}

// Terminate cluster.
//
// Terminates the Spark cluster with the specified ID. The cluster is removed
// asynchronously. Once the termination has completed, the cluster will be in a
// `TERMINATED` state. If the cluster is already in a `TERMINATING` or
// `TERMINATED` state, nothing will happen.
func (a *ClustersAPI) DeleteByClusterId(ctx context.Context, clusterId string) (*DeleteClusterResponse, error) {
	return a.clustersImpl.Delete(ctx, DeleteCluster{
		ClusterId: clusterId,
	})
}

// Get cluster info.
//
// Retrieves the information for a cluster given its identifier. Clusters can be
// described while they are running, or up to 60 days after they are terminated.
func (a *ClustersAPI) GetByClusterId(ctx context.Context, clusterId string) (*ClusterDetails, error) {
	return a.clustersImpl.Get(ctx, GetClusterRequest{
		ClusterId: clusterId,
	})
}

// Get cluster permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *ClustersAPI) GetPermissionLevelsByClusterId(ctx context.Context, clusterId string) (*GetClusterPermissionLevelsResponse, error) {
	return a.clustersImpl.GetPermissionLevels(ctx, GetClusterPermissionLevelsRequest{
		ClusterId: clusterId,
	})
}

// Get cluster permissions.
//
// Gets the permissions of a cluster. Clusters can inherit permissions from
// their root object.
func (a *ClustersAPI) GetPermissionsByClusterId(ctx context.Context, clusterId string) (*ClusterPermissions, error) {
	return a.clustersImpl.GetPermissions(ctx, GetClusterPermissionsRequest{
		ClusterId: clusterId,
	})
}

// ClusterDetailsClusterNameToClusterIdMap calls [ClustersAPI.ListAll] and creates a map of results with [ClusterDetails].ClusterName as key and [ClusterDetails].ClusterId as value.
//
// Returns an error if there's more than one [ClusterDetails] with the same .ClusterName.
//
// Note: All [ClusterDetails] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClustersAPI) ClusterDetailsClusterNameToClusterIdMap(ctx context.Context, request ListClustersRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
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

// Permanently delete cluster.
//
// Permanently deletes a Spark cluster. This cluster is terminated and resources
// are asynchronously removed.
//
// In addition, users will no longer see permanently deleted clusters in the
// cluster list, and API users can no longer perform any action on permanently
// deleted clusters.
func (a *ClustersAPI) PermanentDeleteByClusterId(ctx context.Context, clusterId string) (*PermanentDeleteClusterResponse, error) {
	return a.clustersImpl.PermanentDelete(ctx, PermanentDeleteCluster{
		ClusterId: clusterId,
	})
}

// Pin cluster.
//
// Pinning a cluster ensures that the cluster will always be returned by the
// ListClusters API. Pinning a cluster that is already pinned will have no
// effect. This API can only be called by workspace admins.
func (a *ClustersAPI) PinByClusterId(ctx context.Context, clusterId string) (*PinClusterResponse, error) {
	return a.clustersImpl.Pin(ctx, PinCluster{
		ClusterId: clusterId,
	})
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
func (a *ClustersAPI) StartByClusterId(ctx context.Context, clusterId string) (*StartClusterResponse, error) {
	return a.clustersImpl.Start(ctx, StartCluster{
		ClusterId: clusterId,
	})
}

// Unpin cluster.
//
// Unpinning a cluster will allow the cluster to eventually be removed from the
// ListClusters API. Unpinning a cluster that is not pinned will have no effect.
// This API can only be called by workspace admins.
func (a *ClustersAPI) UnpinByClusterId(ctx context.Context, clusterId string) (*UnpinClusterResponse, error) {
	return a.clustersImpl.Unpin(ctx, UnpinCluster{
		ClusterId: clusterId,
	})
}

type CommandExecutionAPI struct {
	commandExecutionImpl
}

type GlobalInitScriptsAPI struct {
	globalInitScriptsImpl
}

// Delete init script.
//
// Deletes a global init script.
func (a *GlobalInitScriptsAPI) DeleteByScriptId(ctx context.Context, scriptId string) (*DeleteResponse, error) {
	return a.globalInitScriptsImpl.Delete(ctx, DeleteGlobalInitScriptRequest{
		ScriptId: scriptId,
	})
}

// Get an init script.
//
// Gets all the details of a script, including its Base64-encoded contents.
func (a *GlobalInitScriptsAPI) GetByScriptId(ctx context.Context, scriptId string) (*GlobalInitScriptDetailsWithContent, error) {
	return a.globalInitScriptsImpl.Get(ctx, GetGlobalInitScriptRequest{
		ScriptId: scriptId,
	})
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

type InstancePoolsAPI struct {
	instancePoolsImpl
}

// Delete an instance pool.
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) (*DeleteInstancePoolResponse, error) {
	return a.instancePoolsImpl.Delete(ctx, DeleteInstancePool{
		InstancePoolId: instancePoolId,
	})
}

// Get instance pool information.
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePool, error) {
	return a.instancePoolsImpl.Get(ctx, GetInstancePoolRequest{
		InstancePoolId: instancePoolId,
	})
}

// Get instance pool permission levels.
//
// Gets the permission levels that a user can have on an object.
func (a *InstancePoolsAPI) GetPermissionLevelsByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePoolPermissionLevelsResponse, error) {
	return a.instancePoolsImpl.GetPermissionLevels(ctx, GetInstancePoolPermissionLevelsRequest{
		InstancePoolId: instancePoolId,
	})
}

// Get instance pool permissions.
//
// Gets the permissions of an instance pool. Instance pools can inherit
// permissions from their root object.
func (a *InstancePoolsAPI) GetPermissionsByInstancePoolId(ctx context.Context, instancePoolId string) (*InstancePoolPermissions, error) {
	return a.instancePoolsImpl.GetPermissions(ctx, GetInstancePoolPermissionsRequest{
		InstancePoolId: instancePoolId,
	})
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

type InstanceProfilesAPI struct {
	instanceProfilesImpl
}

// Remove the instance profile.
//
// Remove the instance profile with the provided ARN. Existing clusters with
// this instance profile will continue to function.
//
// This API is only accessible to admin users.
func (a *InstanceProfilesAPI) RemoveByInstanceProfileArn(ctx context.Context, instanceProfileArn string) (*RemoveResponse, error) {
	return a.instanceProfilesImpl.Remove(ctx, RemoveInstanceProfile{
		InstanceProfileArn: instanceProfileArn,
	})
}

type LibrariesAPI struct {
	librariesImpl
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
func (a *LibrariesAPI) ClusterStatusByClusterId(ctx context.Context, clusterId string) (*ClusterLibraryStatuses, error) {
	return a.librariesImpl.internalClusterStatus(ctx, ClusterStatus{
		ClusterId: clusterId,
	})
}

type PolicyComplianceForClustersAPI struct {
	policyComplianceForClustersImpl
}

// Get cluster policy compliance.
//
// Returns the policy compliance status of a cluster. Clusters could be out of
// compliance if their policy was updated after the cluster was last edited.
func (a *PolicyComplianceForClustersAPI) GetComplianceByClusterId(ctx context.Context, clusterId string) (*GetClusterComplianceResponse, error) {
	return a.policyComplianceForClustersImpl.GetCompliance(ctx, GetClusterComplianceRequest{
		ClusterId: clusterId,
	})
}

type PolicyFamiliesAPI struct {
	policyFamiliesImpl
}

// Get policy family information.
//
// Retrieve the information for an policy family based on its identifier and
// version
func (a *PolicyFamiliesAPI) GetByPolicyFamilyId(ctx context.Context, policyFamilyId string) (*PolicyFamily, error) {
	return a.policyFamiliesImpl.Get(ctx, GetPolicyFamilyRequest{
		PolicyFamilyId: policyFamilyId,
	})
}
