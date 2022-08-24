// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type InstancepoolsService interface {
    // Creates a new Instance Pool 
    CreateInstancePool(ctx context.Context, createInstancePoolRequest CreateInstancePoolRequest) (*CreateInstancePoolResponse, error)
    // Deletes an Instance Pool 
    DeleteInstancePool(ctx context.Context, deleteInstancePoolRequest DeleteInstancePoolRequest) error
    // Edits an existing Instance Pool 
    EditInstancePool(ctx context.Context, editInstancePoolRequest EditInstancePoolRequest) error
    // Returns an Instance Pool 
    GetInstancePool(ctx context.Context, getInstancePoolRequest GetInstancePoolRequest) (*GetInstancePoolResponse, error)
    // Returns list of Instance Pools 
    ListInstancePools(ctx context.Context) (*ListInstancePoolsResponse, error)
}

func New(client *client.DatabricksClient) InstancepoolsService {
	return &InstancepoolsAPI{
		client: client,
	}
}

type InstancepoolsAPI struct {
	client *client.DatabricksClient
}

// Creates a new Instance Pool 
func (a *InstancepoolsAPI) CreateInstancePool(ctx context.Context, request CreateInstancePoolRequest) (*CreateInstancePoolResponse, error) {
	var createInstancePoolResponse CreateInstancePoolResponse
	path := "/api/2.0/instance-pools/create"
	err := a.client.Post(ctx, path, request, &createInstancePoolResponse)
	return &createInstancePoolResponse, err
}

// Deletes an Instance Pool 
func (a *InstancepoolsAPI) DeleteInstancePool(ctx context.Context, request DeleteInstancePoolRequest) error {
	path := "/api/2.0/instance-pools/delete"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Edits an existing Instance Pool 
func (a *InstancepoolsAPI) EditInstancePool(ctx context.Context, request EditInstancePoolRequest) error {
	path := "/api/2.0/instance-pools/edit"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Returns an Instance Pool 
func (a *InstancepoolsAPI) GetInstancePool(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePoolResponse, error) {
	var getInstancePoolResponse GetInstancePoolResponse
	path := "/api/2.0/instance-pools/get"
	err := a.client.Get(ctx, path, request, &getInstancePoolResponse)
	return &getInstancePoolResponse, err
}

// Returns list of Instance Pools 
func (a *InstancepoolsAPI) ListInstancePools(ctx context.Context) (*ListInstancePoolsResponse, error) {
	var listInstancePoolsResponse ListInstancePoolsResponse
	path := "/api/2.0/instance-pools/list"
	err := a.client.Get(ctx, path, nil, &listInstancePoolsResponse)
	return &listInstancePoolsResponse, err
}

