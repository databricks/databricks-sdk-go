// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

import (
	"context"
)

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
type InstancePoolsService interface {

	// Create a new instance pool
	//
	// Creates a new instance pool using idle and ready-to-use cloud instances.
	Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error)

	// Delete an instance pool
	//
	// Deletes the instance pool permanently. The idle instances in the pool are
	// terminated asynchronously.
	Delete(ctx context.Context, request DeleteInstancePool) error

	// Edit an existing instance pool
	//
	// Modifies the configuration of an existing instance pool.
	Edit(ctx context.Context, request EditInstancePool) error

	// Get instance pool information
	//
	// Retrieve the information for an instance pool based on its identifier.
	Get(ctx context.Context, request GetRequest) (*GetInstancePool, error)

	// List instance pool info
	//
	// Gets a list of instance pools with their statistics.
	//
	// Use ListAll() to get all InstancePoolAndStats instances
	List(ctx context.Context) (*ListInstancePools, error)
}
