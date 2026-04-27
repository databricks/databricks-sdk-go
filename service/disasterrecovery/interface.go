// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disasterrecovery

import (
	"context"
)

// Manage disaster recovery configurations and execute failover operations.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DisasterRecoveryService interface {

	// Create a new failover group.
	CreateFailoverGroup(ctx context.Context, request CreateFailoverGroupRequest) (*FailoverGroup, error)

	// Create a new stable URL.
	CreateStableUrl(ctx context.Context, request CreateStableUrlRequest) (*StableUrl, error)

	// Delete a failover group.
	DeleteFailoverGroup(ctx context.Context, request DeleteFailoverGroupRequest) error

	// Delete a stable URL.
	DeleteStableUrl(ctx context.Context, request DeleteStableUrlRequest) error

	// Initiate a failover to a new primary region.
	FailoverFailoverGroup(ctx context.Context, request FailoverFailoverGroupRequest) (*FailoverGroup, error)

	// Get a failover group.
	GetFailoverGroup(ctx context.Context, request GetFailoverGroupRequest) (*FailoverGroup, error)

	// Get a stable URL.
	GetStableUrl(ctx context.Context, request GetStableUrlRequest) (*StableUrl, error)

	// List failover groups.
	ListFailoverGroups(ctx context.Context, request ListFailoverGroupsRequest) (*ListFailoverGroupsResponse, error)

	// List stable URLs for an account.
	ListStableUrls(ctx context.Context, request ListStableUrlsRequest) (*ListStableUrlsResponse, error)

	// Update a failover group.
	UpdateFailoverGroup(ctx context.Context, request UpdateFailoverGroupRequest) (*FailoverGroup, error)
}
