// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"
	
	
	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewInstancePools(client *client.DatabricksClient) InstancePoolsService {
	return &InstancePoolsAPI{
		client: client,
	}
}

type InstancePoolsAPI struct {
	client *client.DatabricksClient
}

// Creates a new Instance Pool
func (a *InstancePoolsAPI) Create(ctx context.Context, request CreateInstancePoolRequest) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Post(ctx, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

// Deletes an Instance Pool
func (a *InstancePoolsAPI) Delete(ctx context.Context, request DeleteInstancePoolRequest) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes an Instance Pool
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error {
	return a.Delete(ctx, DeleteInstancePoolRequest{
		InstancePoolId: instancePoolId,
	})
}

// Edits an existing Instance Pool
func (a *InstancePoolsAPI) Edit(ctx context.Context, request EditInstancePoolRequest) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns an Instance Pool
func (a *InstancePoolsAPI) Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePoolResponse, error) {
	var getInstancePoolResponse GetInstancePoolResponse
	path := "/api/2.0/instance-pools/get"
	err := a.client.Get(ctx, path, request, &getInstancePoolResponse)
	return &getInstancePoolResponse, err
}

// Returns an Instance Pool
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePoolResponse, error) {
	return a.Get(ctx, GetInstancePoolRequest{
		InstancePoolId: instancePoolId,
	})
}

// Returns list of Instance Pools
func (a *InstancePoolsAPI) List(ctx context.Context) (*ListInstancePoolsResponse, error) {
	var listInstancePoolsResponse ListInstancePoolsResponse
	path := "/api/2.0/instance-pools/list"
	err := a.client.Get(ctx, path, nil, &listInstancePoolsResponse)
	return &listInstancePoolsResponse, err
}

