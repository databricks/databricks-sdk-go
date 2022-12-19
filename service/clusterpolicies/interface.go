// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"
)

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
type ClusterPoliciesService interface {

	// Create a new policy.
	//
	// Creates a new policy with prescribed settings.
	Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error)

	// Delete a cluster policy.
	//
	// Delete a policy for a cluster. Clusters governed by this policy can still
	// run, but cannot be edited.
	Delete(ctx context.Context, request DeletePolicy) error

	// Update a cluster policy.
	//
	// Update an existing policy for cluster. This operation may make some
	// clusters governed by the previous policy invalid.
	Edit(ctx context.Context, request EditPolicy) error

	// Get entity.
	//
	// Get a cluster policy entity. Creation and editing is available to admins
	// only.
	Get(ctx context.Context, request Get) (*Policy, error)

	// Get a cluster policy.
	//
	// Returns a list of policies accessible by the requesting user.
	//
	// Use ListAll() to get all Policy instances
	List(ctx context.Context, request List) (*ListPoliciesResponse, error)
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
type PolicyFamiliesService interface {
	Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error)

	//
	// Use ListAll() to get all PolicyFamily instances, which will iterate over every result page.
	List(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error)
}
