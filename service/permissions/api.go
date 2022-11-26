// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

func NewPermissions(client *client.DatabricksClient) *PermissionsAPI {
	return &PermissionsAPI{
		impl: &permissionsImpl{
			client: client,
		},
	}
}

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
type PermissionsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PermissionsService)
	impl PermissionsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PermissionsAPI) WithImpl(impl PermissionsService) *PermissionsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Permissions API implementation
func (a *PermissionsAPI) Impl() PermissionsService {
	return a.impl
}

// Get permission levels
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, request)
}

// Get permission levels
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevelsByObjectTypeAndId(ctx context.Context, objectType string, id string) (*GetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		ObjectType: objectType,
		Id:         id,
	})
}

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*ObjectPermissions, error) {
	return a.impl.GetPermissions(ctx, request)
}

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetPermissionsByObjectTypeAndId(ctx context.Context, objectType string, id string) (*ObjectPermissions, error) {
	return a.impl.GetPermissions(ctx, GetPermissionsRequest{
		ObjectType: objectType,
		Id:         id,
	})
}

// Set permissions
//
// Sets permissions on object. Objects can inherit permissions from their parent
// objects and root objects.
func (a *PermissionsAPI) Set(ctx context.Context, request PermissionsRequest) error {
	return a.impl.Set(ctx, request)
}

// Update permission
//
// Updates the permissions on an object.
func (a *PermissionsAPI) Update(ctx context.Context, request PermissionsRequest) error {
	return a.impl.Update(ctx, request)
}

func NewWorkspaceAssignment(client *client.DatabricksClient) *WorkspaceAssignmentAPI {
	return &WorkspaceAssignmentAPI{
		impl: &workspaceAssignmentImpl{
			client: client,
		},
	}
}

// Databricks Workspace Assignment REST API
type WorkspaceAssignmentAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspaceAssignmentService)
	impl WorkspaceAssignmentService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WorkspaceAssignmentAPI) WithImpl(impl WorkspaceAssignmentService) *WorkspaceAssignmentAPI {
	a.impl = impl
	return a
}

// Impl returns low-level WorkspaceAssignment API implementation
func (a *WorkspaceAssignmentAPI) Impl() WorkspaceAssignmentService {
	return a.impl
}

// Create permission assignments
//
// Create new permission assignments for the specified account and workspace.
func (a *WorkspaceAssignmentAPI) Create(ctx context.Context, request CreateWorkspaceAssignments) (*WorkspaceAssignmentsCreated, error) {
	return a.impl.Create(ctx, request)
}

// Delete permissions assignment
//
// Deletes the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) Delete(ctx context.Context, request DeleteRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete permissions assignment
//
// Deletes the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) DeleteByWorkspaceIdAndPrincipalId(ctx context.Context, workspaceId int64, principalId int64) error {
	return a.impl.Delete(ctx, DeleteRequest{
		WorkspaceId: workspaceId,
		PrincipalId: principalId,
	})
}

// List workspace permissions
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) Get(ctx context.Context, request GetRequest) (*WorkspacePermissions, error) {
	return a.impl.Get(ctx, request)
}

// List workspace permissions
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*WorkspacePermissions, error) {
	return a.impl.Get(ctx, GetRequest{
		WorkspaceId: workspaceId,
	})
}

// Get permission assignments
//
// Get the permission assignments for the specified Databricks Account and
// Databricks Workspace.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAssignmentAPI) ListAll(ctx context.Context, request ListRequest) ([]PermissionAssignment, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.PermissionAssignments, nil
}

// Get permission assignments
//
// Get the permission assignments for the specified Databricks Account and
// Databricks Workspace.
func (a *WorkspaceAssignmentAPI) ListByWorkspaceId(ctx context.Context, workspaceId int64) (*PermissionAssignments, error) {
	return a.impl.List(ctx, ListRequest{
		WorkspaceId: workspaceId,
	})
}

// Update permissions assignment
//
// Updates the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	return a.impl.Update(ctx, request)
}
