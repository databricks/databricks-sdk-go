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

	// Create a Branch.
	CreateBranch(ctx context.Context, request CreateBranchRequest) (*Operation, error)

	// Create an Endpoint.
	CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Operation, error)

	// Create a Project.
	CreateProject(ctx context.Context, request CreateProjectRequest) (*Operation, error)

	// Delete a Branch.
	DeleteBranch(ctx context.Context, request DeleteBranchRequest) error

	// Delete an Endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) error

	// Delete a Project.
	DeleteProject(ctx context.Context, request DeleteProjectRequest) error

	// Get a Branch.
	GetBranch(ctx context.Context, request GetBranchRequest) (*Branch, error)

	// Get an Endpoint.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error)

	// Get an Operation.
	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Get a Project.
	GetProject(ctx context.Context, request GetProjectRequest) (*Project, error)

	// List Branches.
	ListBranches(ctx context.Context, request ListBranchesRequest) (*ListBranchesResponse, error)

	// List Endpoints.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error)

	// List Projects.
	ListProjects(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error)

	// Update a Branch.
	UpdateBranch(ctx context.Context, request UpdateBranchRequest) (*Operation, error)

	// Update an Endpoint.
	UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Operation, error)

	// Update a Project.
	UpdateProject(ctx context.Context, request UpdateProjectRequest) (*Operation, error)
}
