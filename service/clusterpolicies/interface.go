// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"
)



type ClusterPoliciesService interface {
    // Creates a new Policy
    CreatePolicy(ctx context.Context, createPolicyRequest CreatePolicyRequest) (*CreatePolicyResponse, error)
    // Delete a policy
    DeletePolicy(ctx context.Context, deletePolicyRequest DeletePolicyRequest) error
    // Update an existing policy
    EditPolicy(ctx context.Context, editPolicyRequest EditPolicyRequest) error
    // Returns a Policy
    GetPolicy(ctx context.Context, getPolicyRequest GetPolicyRequest) (*GetPolicyResponse, error)
    // Returns list of policies
    ListPolicies(ctx context.Context) (*ListPoliciesResponse, error)
}
