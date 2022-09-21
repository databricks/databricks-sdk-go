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

// Creates a new Instance Pool using idle and ready-to-use cloud instances
func (a *InstancePoolsAPI) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Post(ctx, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) Delete(ctx context.Context, request DeleteInstancePool) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error {
	return a.Delete(ctx, DeleteInstancePool{
		InstancePoolId: instancePoolId,
	})
}

// Modifies the configuration of an existing instance pool.
func (a *InstancePoolsAPI) Edit(ctx context.Context, request EditInstancePool) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Retrieve the information for an instance pool given its identifier.
func (a *InstancePoolsAPI) Get(ctx context.Context, request GetRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	err := a.client.Get(ctx, path, request, &getInstancePool)
	return &getInstancePool, err
}

// Retrieve the information for an instance pool given its identifier.
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePool, error) {
	return a.Get(ctx, GetRequest{
		InstancePoolId: instancePoolId,
	})
}

// Returns list of Instance Pools with their statistics included
func (a *InstancePoolsAPI) List(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"
	err := a.client.Get(ctx, path, nil, &listInstancePools)
	return &listInstancePools, err
}
