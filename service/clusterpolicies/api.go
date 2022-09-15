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
func (a *ClusterPoliciesAPI) Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error) {
	var createPolicyResponse CreatePolicyResponse
	path := "/api/2.0/policies/clusters/create"
	err := a.client.Post(ctx, path, request, &createPolicyResponse)
	return &createPolicyResponse, err
}

// Delete a policy
func (a *ClusterPoliciesAPI) Delete(ctx context.Context, request DeletePolicy) error {
	path := "/api/2.0/policies/clusters/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete a policy
func (a *ClusterPoliciesAPI) DeleteByPolicyId(ctx context.Context, policyId string) error {
	return a.Delete(ctx, DeletePolicy{
		PolicyId: policyId,
	})
}

// Update an existing policy
func (a *ClusterPoliciesAPI) Edit(ctx context.Context, request EditPolicy) error {
	path := "/api/2.0/policies/clusters/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns a Policy
func (a *ClusterPoliciesAPI) Get(ctx context.Context, request GetRequest) (*Policy, error) {
	var policy Policy
	path := "/api/2.0/policies/clusters/get"
	err := a.client.Get(ctx, path, request, &policy)
	return &policy, err
}

// Returns a Policy
func (a *ClusterPoliciesAPI) GetByPolicyId(ctx context.Context, policyId string) (*Policy, error) {
	return a.Get(ctx, GetRequest{
		PolicyId: policyId,
	})
}

// Returns list of policies
func (a *ClusterPoliciesAPI) List(ctx context.Context) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	err := a.client.Get(ctx, path, nil, &listPoliciesResponse)
	return &listPoliciesResponse, err
}
