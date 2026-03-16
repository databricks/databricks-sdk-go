// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package environments

import (
	"context"
)

// APIs to manage environment resources.
//
// The Environments API provides management capabilities for different types of
// environments including workspace-level base environments that define the
// environment version and dependencies to be used in serverless notebooks and
// jobs.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type EnvironmentsService interface {

	// Creates a new WorkspaceBaseEnvironment. This is a long-running operation.
	// The operation will asynchronously generate a materialized environment to
	// optimize dependency resolution and is only marked as done when the
	// materialized environment has been successfully generated or has failed.
	CreateWorkspaceBaseEnvironment(ctx context.Context, request CreateWorkspaceBaseEnvironmentRequest) (*Operation, error)

	// Deletes a WorkspaceBaseEnvironment. Deleting a base environment may
	// impact linked notebooks and jobs. This operation is irreversible and
	// should be performed only when you are certain the environment is no
	// longer needed.
	DeleteWorkspaceBaseEnvironment(ctx context.Context, request DeleteWorkspaceBaseEnvironmentRequest) error

	// Gets the default WorkspaceBaseEnvironment configuration for the
	// workspace. Returns the current default base environment settings for both
	// CPU and GPU compute.
	GetDefaultWorkspaceBaseEnvironment(ctx context.Context, request GetDefaultWorkspaceBaseEnvironmentRequest) (*DefaultWorkspaceBaseEnvironment, error)

	// Gets the status of a long-running operation. Clients can use this method
	// to poll the operation result.
	GetOperation(ctx context.Context, request GetOperationRequest) (*Operation, error)

	// Retrieves a WorkspaceBaseEnvironment by its name.
	GetWorkspaceBaseEnvironment(ctx context.Context, request GetWorkspaceBaseEnvironmentRequest) (*WorkspaceBaseEnvironment, error)

	// Lists all WorkspaceBaseEnvironments in the workspace.
	ListWorkspaceBaseEnvironments(ctx context.Context, request ListWorkspaceBaseEnvironmentsRequest) (*ListWorkspaceBaseEnvironmentsResponse, error)

	// Refreshes the materialized environment for a WorkspaceBaseEnvironment.
	// This is a long-running operation. The operation will asynchronously
	// regenerate the materialized environment and is only marked as done when
	// the materialized environment has been successfully generated or has
	// failed. The existing materialized environment remains available until it
	// expires.
	RefreshWorkspaceBaseEnvironment(ctx context.Context, request RefreshWorkspaceBaseEnvironmentRequest) (*Operation, error)

	// Updates the default WorkspaceBaseEnvironment configuration for the
	// workspace. Sets the specified base environments as the workspace defaults
	// for CPU and/or GPU compute.
	UpdateDefaultWorkspaceBaseEnvironment(ctx context.Context, request UpdateDefaultWorkspaceBaseEnvironmentRequest) (*DefaultWorkspaceBaseEnvironment, error)

	// Updates an existing WorkspaceBaseEnvironment. This is a long-running
	// operation. The operation will asynchronously regenerate the materialized
	// environment and is only marked as done when the materialized environment
	// has been successfully generated or has failed. The existing materialized
	// environment remains available until it expires.
	UpdateWorkspaceBaseEnvironment(ctx context.Context, request UpdateWorkspaceBaseEnvironmentRequest) (*Operation, error)
}
