// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"
)

func NewInstancePools(client databricksClient) *InstancePoolsAPI {
	return &InstancePoolsAPI{
		impl: &instancePoolsImpl{
			client: client,
		},
	}
}

// Instance Pools API are used to create, edit, delete and list instance pools
// by using ready-to-use cloud instances which reduces a cluster start and
// auto-scaling times.
//
// Databricks pools reduce cluster start and auto-scaling times by maintaining a
// set of idle, ready-to-use instances. When a cluster is attached to a pool,
// cluster nodes are created using the pool’s idle instances. If the pool has
// no idle instances, the pool expands by allocating a new instance from the
// instance provider in order to accommodate the cluster’s request. When a
// cluster releases an instance, it returns to the pool and is free for another
// cluster to use. Only clusters attached to a pool can use that pool’s idle
// instances.
//
// You can specify a different pool for the driver node and worker nodes, or use
// the same pool for both.
//
// Databricks does not charge DBUs while instances are idle in the pool.
// Instance provider billing does apply. See pricing.
type InstancePoolsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(InstancePoolsService)
	impl InstancePoolsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *InstancePoolsAPI) WithImpl(impl InstancePoolsService) *InstancePoolsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level InstancePools API implementation
func (a *InstancePoolsAPI) Impl() InstancePoolsService {
	return a.impl
}

// Create a new instance pool
//
// Creates a new instance pool using idle and ready-to-use cloud instances.
func (a *InstancePoolsAPI) Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete an instance pool
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) Delete(ctx context.Context, request DeleteInstancePool) error {
	return a.impl.Delete(ctx, request)
}

// Delete an instance pool
//
// Deletes the instance pool permanently. The idle instances in the pool are
// terminated asynchronously.
func (a *InstancePoolsAPI) DeleteByInstancePoolId(ctx context.Context, instancePoolId string) error {
	return a.impl.Delete(ctx, DeleteInstancePool{
		InstancePoolId: instancePoolId,
	})
}

// Edit an existing instance pool
//
// Modifies the configuration of an existing instance pool.
func (a *InstancePoolsAPI) Edit(ctx context.Context, request EditInstancePool) error {
	return a.impl.Edit(ctx, request)
}

// Get instance pool information
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) Get(ctx context.Context, request GetRequest) (*GetInstancePool, error) {
	return a.impl.Get(ctx, request)
}

// Get instance pool information
//
// Retrieve the information for an instance pool based on its identifier.
func (a *InstancePoolsAPI) GetByInstancePoolId(ctx context.Context, instancePoolId string) (*GetInstancePool, error) {
	return a.impl.Get(ctx, GetRequest{
		InstancePoolId: instancePoolId,
	})
}

// List instance pool info
//
// Gets a list of instance pools with their statistics.
//
// This method is generated by Databricks SDK Code Generator.
func (a *InstancePoolsAPI) ListAll(ctx context.Context) ([]InstancePoolAndStats, error) {
	response, err := a.impl.List(ctx)
	if err != nil {
		return nil, err
	}
	return response.InstancePools, nil
}
