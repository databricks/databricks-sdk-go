// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package postgres

import (
	"context"
)

// Use the Postgres API to create and manage Lakebase Autoscaling Postgres
// infrastructure, including projects, branches, compute endpoints, and roles.
//
// This API manages database infrastructure only. To query or modify data, use
// the Data API or direct SQL connections.
//
// **About resource IDs and names**
//
// Resources are identified by hierarchical resource names like
// `projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id}`. The
// `name` field on each resource contains this full path and is output-only.
// Note that `name` refers to this resource path, not the user-visible
// `display_name`.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PostgresService interface {

	// Creates a new database branch in the project.
	CreateBranch(ctx context.Context, request CreateBranchRequest) (*Operation, error)

	// Creates a new compute endpoint in the branch.
	CreateEndpoint(ctx context.Context, request CreateEndpointRequest) (*Operation, error)

	// Creates a new Lakebase Autoscaling Postgres database project, which
	// contains branches and compute endpoints.
	CreateProject(ctx context.Context, request CreateProjectRequest) (*Operation, error)

	// Creates a new Postgres role in the branch.
	CreateRole(ctx context.Context, request CreateRoleRequest) (*Operation, error)

	// Deletes the specified database branch.
	DeleteBranch(ctx context.Context, request DeleteBranchRequest) (*Operation, error)

	// Deletes the specified compute endpoint.
	DeleteEndpoint(ctx context.Context, request DeleteEndpointRequest) (*Operation, error)

	// Deletes the specified database project.
	DeleteProject(ctx context.Context, request DeleteProjectRequest) (*Operation, error)

	// Deletes the specified Postgres role.
	DeleteRole(ctx context.Context, request DeleteRoleRequest) (*Operation, error)

	// Generate OAuth credentials for a Postgres database.
	GenerateDatabaseCredential(ctx context.Context, request GenerateDatabaseCredentialRequest) (*DatabaseCredential, error)

	// Retrieves information about the specified database branch.
	GetBranch(ctx context.Context, request GetBranchRequest) (*Branch, error)

	// Retrieves information about the specified compute endpoint, including its
	// connection details and operational state.
	GetEndpoint(ctx context.Context, request GetEndpointRequest) (*Endpoint, error)

	// Retrieves the status of a long-running operation.
	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Retrieves information about the specified database project.
	GetProject(ctx context.Context, request GetProjectRequest) (*Project, error)

	// Retrieves information about the specified Postgres role, including its
	// authentication method and permissions.
	GetRole(ctx context.Context, request GetRoleRequest) (*Role, error)

	// Returns a paginated list of database branches in the project.
	ListBranches(ctx context.Context, request ListBranchesRequest) (*ListBranchesResponse, error)

	// Returns a paginated list of compute endpoints in the branch.
	ListEndpoints(ctx context.Context, request ListEndpointsRequest) (*ListEndpointsResponse, error)

	// Returns a paginated list of database projects in the workspace that the
	// user has permission to access.
	ListProjects(ctx context.Context, request ListProjectsRequest) (*ListProjectsResponse, error)

	// Returns a paginated list of Postgres roles in the branch.
	ListRoles(ctx context.Context, request ListRolesRequest) (*ListRolesResponse, error)

	// Updates the specified database branch. You can set this branch as the
	// project's default branch, or protect/unprotect it.
	UpdateBranch(ctx context.Context, request UpdateBranchRequest) (*Operation, error)

	// Updates the specified compute endpoint. You can update autoscaling
	// limits, suspend timeout, or enable/disable the compute endpoint.
	UpdateEndpoint(ctx context.Context, request UpdateEndpointRequest) (*Operation, error)

	// Updates the specified database project.
	UpdateProject(ctx context.Context, request UpdateProjectRequest) (*Operation, error)
}
