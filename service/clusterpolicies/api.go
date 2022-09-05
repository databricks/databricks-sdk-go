// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type ClusterpoliciesService interface {
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

func New(client *client.DatabricksClient) ClusterpoliciesService {
	return &ClusterpoliciesAPI{
		client: client,
	}
}

type ClusterpoliciesAPI struct {
	client *client.DatabricksClient
}

// Creates a new Policy
func (a *ClusterpoliciesAPI) CreatePolicy(ctx context.Context, request CreatePolicyRequest) (*CreatePolicyResponse, error) {
	var createPolicyResponse CreatePolicyResponse
	path := "/api/2.0/policies/clusters/create"
	err := a.client.Post(ctx, path, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

// Delete a policy
func (a *ClusterpoliciesAPI) DeletePolicy(ctx context.Context, request DeletePolicyRequest) error {
	path := "/api/2.0/policies/clusters/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Update an existing policy
func (a *ClusterpoliciesAPI) EditPolicy(ctx context.Context, request EditPolicyRequest) error {
	path := "/api/2.0/policies/clusters/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns a Policy
func (a *ClusterpoliciesAPI) GetPolicy(ctx context.Context, request GetPolicyRequest) (*GetPolicyResponse, error) {
	var getPolicyResponse GetPolicyResponse
	path := "/api/2.0/policies/clusters/get"
	err := a.client.Get(ctx, path, request, &getPolicyResponse)
	return &getPolicyResponse, err
}

// Returns list of policies
func (a *ClusterpoliciesAPI) ListPolicies(ctx context.Context) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	err := a.client.Get(ctx, path, nil, &listPoliciesResponse)
	return &listPoliciesResponse, err
}
