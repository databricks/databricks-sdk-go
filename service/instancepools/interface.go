// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"
	
)



type InstancePoolsService interface {
    // Creates a new Instance Pool
    Create(ctx context.Context, createInstancePoolRequest CreateInstancePoolRequest) (*CreateInstancePoolResponse, error)
    // Deletes an Instance Pool
    Delete(ctx context.Context, deleteInstancePoolRequest DeleteInstancePoolRequest) error
	DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error
    // Edits an existing Instance Pool
    Edit(ctx context.Context, editInstancePoolRequest EditInstancePoolRequest) error
    // Returns an Instance Pool
    Get(ctx context.Context, getInstancePoolRequest GetInstancePoolRequest) (*GetInstancePoolResponse, error)
	GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePoolResponse, error)
    // Returns list of Instance Pools
    List(ctx context.Context) (*ListInstancePoolsResponse, error)
}
