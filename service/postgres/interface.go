// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"context"
)

// The Postgres API provides access to a Postgres database via REST API or
// direct SQL.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PostgresService interface {

	// Create a Database Branch.
	CreateDatabaseBranch(ctx context.Context, request CreateDatabaseBranchRequest) (*Operation, error)

	// Create a Database Endpoint.
	CreateDatabaseEndpoint(ctx context.Context, request CreateDatabaseEndpointRequest) (*Operation, error)

	// Create a Database Project.
	CreateDatabaseProject(ctx context.Context, request CreateDatabaseProjectRequest) (*Operation, error)

	// Delete a Database Branch.
	DeleteDatabaseBranch(ctx context.Context, request DeleteDatabaseBranchRequest) error

	// Delete a Database Endpoint.
	DeleteDatabaseEndpoint(ctx context.Context, request DeleteDatabaseEndpointRequest) error

	// Delete a Database Project.
	DeleteDatabaseProject(ctx context.Context, request DeleteDatabaseProjectRequest) error

	// Get a Database Branch.
	GetDatabaseBranch(ctx context.Context, request GetDatabaseBranchRequest) (*DatabaseBranch, error)

	// Get a Database Endpoint.
	GetDatabaseEndpoint(ctx context.Context, request GetDatabaseEndpointRequest) (*DatabaseEndpoint, error)

	// Get a Database Operation.
	GetDatabaseOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Get a Database Project.
	GetDatabaseProject(ctx context.Context, request GetDatabaseProjectRequest) (*DatabaseProject, error)

	// List Database Branches.
	ListDatabaseBranches(ctx context.Context, request ListDatabaseBranchesRequest) (*ListDatabaseBranchesResponse, error)

	// List Database Endpoints.
	ListDatabaseEndpoints(ctx context.Context, request ListDatabaseEndpointsRequest) (*ListDatabaseEndpointsResponse, error)

	// List Database Projects.
	ListDatabaseProjects(ctx context.Context, request ListDatabaseProjectsRequest) (*ListDatabaseProjectsResponse, error)

	// Restart a Database Endpoint.
	RestartDatabaseEndpoint(ctx context.Context, request RestartDatabaseEndpointRequest) (*Operation, error)

	// Update a Database Branch.
	UpdateDatabaseBranch(ctx context.Context, request UpdateDatabaseBranchRequest) (*Operation, error)

	// Update a Database Endpoint.
	UpdateDatabaseEndpoint(ctx context.Context, request UpdateDatabaseEndpointRequest) (*Operation, error)

	// Update a Database Project.
	UpdateDatabaseProject(ctx context.Context, request UpdateDatabaseProjectRequest) (*Operation, error)
}
