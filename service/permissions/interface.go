// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
)

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
type PermissionsService interface {

	// Get object permissions.
	//
	// Gets the permission of an object. Objects can inherit permissions from
	// their parent objects or root objects.
	Get(ctx context.Context, request Get) (*ObjectPermissions, error)

	// Get permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevels) (*GetPermissionLevelsResponse, error)

	// Set permissions.
	//
	// Sets permissions on object. Objects can inherit permissions from their
	// parent objects and root objects.
	Set(ctx context.Context, request PermissionsRequest) error

	// Update permission.
	//
	// Updates the permissions on an object.
	Update(ctx context.Context, request PermissionsRequest) error
}

// The Workspace Permission Assignment API allows you to manage workspace
// permissions for principals in your account.
type WorkspaceAssignmentService interface {

	// Delete permissions assignment.
	//
	// Deletes the workspace permissions assignment in a given account and
	// workspace for the specified principal.
	Delete(ctx context.Context, request DeleteWorkspaceAssignmentRequest) error

	// List workspace permissions.
	//
	// Get an array of workspace permissions for the specified account and
	// workspace.
	Get(ctx context.Context, request GetWorkspaceAssignmentRequest) (*WorkspacePermissions, error)

	// Get permission assignments.
	//
	// Get the permission assignments for the specified Databricks Account and
	// Databricks Workspace.
	//
	// Use ListAll() to get all PermissionAssignment instances
	List(ctx context.Context, request ListWorkspaceAssignmentRequest) (*PermissionAssignments, error)

	// Create or update permissions assignment.
	//
	// Creates or updates the workspace permissions assignment in a given
	// account and workspace for the specified principal.
	Update(ctx context.Context, request UpdateWorkspaceAssignments) error
}
