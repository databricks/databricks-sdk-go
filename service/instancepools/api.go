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

// Create a new instance pool
//
// Creates a new instance pool using idle and ready-to-use cloud instances.
func (a *InstancePoolsAPI) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Post(ctx, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

// Delete an instance pool
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) Delete(ctx context.Context, request DeleteInstancePool) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Delete an instance pool
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error {
	return a.Delete(ctx, DeleteInstancePool{
		InstancePoolId: instancePoolId,
	})
}

// Edit an existing instance pool
//
// Modifies the configuration of an existing instance pool.
func (a *InstancePoolsAPI) Edit(ctx context.Context, request EditInstancePool) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Get instance pool information
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) Get(ctx context.Context, request GetRequest) (*GetInstancePool, error) {
	var getInstancePool GetInstancePool
	path := "/api/2.0/instance-pools/get"
	err := a.client.Get(ctx, path, request, &getInstancePool)
	return &getInstancePool, err
}

// Get instance pool information
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePool, error) {
	return a.Get(ctx, GetRequest{
		InstancePoolId: instancePoolId,
	})
}

// List instance pool info
//
// Gets a list of instance pools with their statistics.
func (a *InstancePoolsAPI) List(ctx context.Context) (*ListInstancePools, error) {
	var listInstancePools ListInstancePools
	path := "/api/2.0/instance-pools/list"
	err := a.client.Get(ctx, path, nil, &listInstancePools)
	return &listInstancePools, err
}

func (a *InstancePoolsAPI) ListAll(ctx context.Context) ([]InstancePoolAndStats, error) {
	response, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.InstancePools, nil
}

func (a *InstancePoolsAPI) InstancePoolAndStatsInstancePoolNameToInstancePoolIdMap(ctx context.Context) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.InstancePoolName] = v.InstancePoolId
	}
	return mapping, nil
}

func (a *InstancePoolsAPI) GetInstancePoolAndStatsByInstancePoolName(ctx context.Context, name string) (*InstancePoolAndStats, error) {
	result, err := a.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.InstancePoolName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("InstancePoolAndStats named '%s' does not exist", name)
}
