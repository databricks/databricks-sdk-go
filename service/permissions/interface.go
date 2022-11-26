// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
)

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
type PermissionsService interface {

	// Get permission levels
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error)

	// Get object permissions
	//
	// Gets the permission of an object. Objects can inherit permissions from
	// their parent objects or root objects.
	GetPermissions(ctx context.Context, request GetPermissionsRequest) (*ObjectPermissions, error)

	// Set permissions
	//
	// Sets permissions on object. Objects can inherit permissions from their
	// parent objects and root objects.
	Set(ctx context.Context, request PermissionsRequest) error

	// Update permission
	//
	// Updates the permissions on an object.
	Update(ctx context.Context, request PermissionsRequest) error
}

// Databricks Workspace Assignment REST API
type WorkspaceAssignmentService interface {

	// Create permission assignments
	//
	// Create new permission assignments for the specified account and
	// workspace.
	Create(ctx context.Context, request CreateWorkspaceAssignments) (*WorkspaceAssignmentsCreated, error)

	// Delete permissions assignment
	//
	// Deletes the workspace permissions assignment for a given account and
	// workspace using the specified service principal.
	Delete(ctx context.Context, request DeleteRequest) error

	// List workspace permissions
	//
	// Get an array of workspace permissions for the specified account and
	// workspace.
	Get(ctx context.Context, request GetRequest) (*WorkspacePermissions, error)

	// Get permission assignments
	//
	// Get the permission assignments for the specified Databricks Account and
	// Databricks Workspace.
	//
	// Use ListAll() to get all PermissionAssignment instances
	List(ctx context.Context, request ListRequest) (*PermissionAssignments, error)

	// Update permissions assignment
	//
	// Updates the workspace permissions assignment for a given account and
	// workspace using the specified service principal.
	Update(ctx context.Context, request UpdateWorkspaceAssignments) error
}
