// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusterpolicies

import (
	"context"
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

func (a *clusterPoliciesImpl) Get(ctx context.Context, request GetRequest) (*Policy, error) {
	var policy Policy
	path := "/api/2.0/policies/clusters/get"
	err := a.client.Do(ctx, http.MethodGet, path, request, &policy)
	return &policy, err
}

func (a *clusterPoliciesImpl) List(ctx context.Context) (*ListPoliciesResponse, error) {
	var listPoliciesResponse ListPoliciesResponse
	path := "/api/2.0/policies/clusters/list"
	err := a.client.Do(ctx, http.MethodGet, path, nil, &listPoliciesResponse)
	return &listPoliciesResponse, err
}
