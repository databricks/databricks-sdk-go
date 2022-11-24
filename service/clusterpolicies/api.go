// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
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

// Create a new policy
//
// Creates a new policy with prescribed settings.
func (a *ClusterPoliciesAPI) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete a cluster policy
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) Delete(ctx context.Context, request DeletePolicy) error {
	return a.impl.Delete(ctx, request)
}

// Delete a cluster policy
//
// Delete a policy for a cluster. Clusters governed by this policy can still
// run, but cannot be edited.
func (a *ClusterPoliciesAPI) DeleteByPolicyId(ctx context.Context, policyId string) error {
	return a.impl.Delete(ctx, DeletePolicy{
		PolicyId: policyId,
	})
}

// Update a cluster policy
//
// Update an existing policy for cluster. This operation may make some clusters
// governed by the previous policy invalid.
func (a *ClusterPoliciesAPI) Edit(ctx context.Context, request EditPolicy) error {
	return a.impl.Edit(ctx, request)
}

// Get entity
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) Get(ctx context.Context, request GetRequest) (*Policy, error) {
	return a.impl.Get(ctx, request)
}

// Get entity
//
// Get a cluster policy entity. Creation and editing is available to admins
// only.
func (a *ClusterPoliciesAPI) GetByPolicyId(ctx context.Context, policyId string) (*Policy, error) {
	return a.impl.Get(ctx, GetRequest{
		PolicyId: policyId,
	})
}

// Get a cluster policy
//
// Returns a list of policies accessible by the requesting user.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) ListAll(ctx context.Context) ([]Policy, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.Policies, nil
}

// PolicyNameToPolicyIdMap calls [ClusterPoliciesAPI.ListAll] and creates a map of results with [Policy].Name as key and [Policy].PolicyId as value.
//
// Note: All [Policy] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) PolicyNameToPolicyIdMap(ctx context.Context) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.Name] = v.PolicyId
	}
	return mapping, nil
}

// GetPolicyByName calls [ClusterPoliciesAPI.PolicyNameToPolicyIdMap] and returns a single [Policy].
//
// Note: All [Policy] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ClusterPoliciesAPI) GetPolicyByName(ctx context.Context, name string) (*Policy, error) {
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.Name != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Policy named '%s' does not exist", name)
}
