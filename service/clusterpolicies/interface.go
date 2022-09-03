// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"
)

type ClusterPoliciesService interface {
	// Creates a new Policy
	Create(ctx context.Context, createPolicyRequest CreatePolicyRequest) (*CreatePolicyResponse, error)
	// Delete a policy
	Delete(ctx context.Context, deletePolicyRequest DeletePolicyRequest) error
	DeleteByPolicyId(ctx context.Context, policyId string) error
	// Update an existing policy
	Edit(ctx context.Context, editPolicyRequest EditPolicyRequest) error
	// Returns a Policy
	Get(ctx context.Context, getPolicyRequest GetPolicyRequest) (*GetPolicyResponse, error)
	GetByPolicyId(ctx context.Context, policyId string) (*GetPolicyResponse, error)
	// Returns list of policies
	List(ctx context.Context) (*ListPoliciesResponse, error)
}
