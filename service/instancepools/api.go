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
func (a *InstancePoolsAPI) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Post(ctx, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

// Deletes an Instance Pool
func (a *InstancePoolsAPI) Delete(ctx context.Context, request DeleteInstancePool) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes an Instance Pool
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error {
	return a.Delete(ctx, DeleteInstancePool{
		InstancePoolId: instancePoolId,
	})
}

// Edits an existing Instance Pool
func (a *InstancePoolsAPI) Edit(ctx context.Context, request EditInstancePool) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns an Instance Pool
func (a *InstancePoolsAPI) Get(ctx context.Context, request GetRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	err := a.client.Get(ctx, path, request, &getInstancePool)
	return &getInstancePool, err
}

// Returns an Instance Pool
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePool, error) {
	return a.Get(ctx, GetRequest{
		InstancePoolId: instancePoolId,
	})
}

// Returns list of Instance Pools
func (a *InstancePoolsAPI) List(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"
	err := a.client.Get(ctx, path, nil, &listInstancePools)
	return &listInstancePools, err
}
