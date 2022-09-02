// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"
)



type InstancePoolsService interface {
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
