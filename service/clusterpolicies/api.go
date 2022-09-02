// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"
	

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewClusterPolicies(client *client.DatabricksClient) ClusterPoliciesService {
	return &ClusterPoliciesAPI{
		client: client,
	}
}

type ClusterPoliciesAPI struct {
	client *client.DatabricksClient
}

// Creates a new Policy
func (a *ClusterPoliciesAPI) CreatePolicy(ctx context.Context, request CreatePolicyRequest) (*CreatePolicyResponse, error) {
	var createPolicyResponse CreatePolicyResponse
	path := "/api/2.0/policies/clusters/create"
	err := a.client.Post(ctx, path, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

// Delete a policy
func (a *ClusterPoliciesAPI) DeletePolicy(ctx context.Context, request DeletePolicyRequest) error {
	path := "/api/2.0/policies/clusters/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Update an existing policy
func (a *ClusterPoliciesAPI) EditPolicy(ctx context.Context, request EditPolicyRequest) error {
	path := "/api/2.0/policies/clusters/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns a Policy
func (a *ClusterPoliciesAPI) GetPolicy(ctx context.Context, request GetPolicyRequest) (*GetPolicyResponse, error) {
	var getPolicyResponse GetPolicyResponse
	path := "/api/2.0/policies/clusters/get"
	err := a.client.Get(ctx, path, request, &getPolicyResponse)
	return &getPolicyResponse, err
}

// Returns list of policies
func (a *ClusterPoliciesAPI) ListPolicies(ctx context.Context) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	err := a.client.Get(ctx, path, nil, &listPoliciesResponse)
	return &listPoliciesResponse, err
}

