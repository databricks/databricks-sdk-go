// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
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

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetObjectPermissions(ctx context.Context, request GetObjectPermissionsRequest) (*ObjectPermissions, error) {
	return a.impl.GetObjectPermissions(ctx, request)
}

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetObjectPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType string, objectId string) (*ObjectPermissions, error) {
	return a.impl.GetObjectPermissions(ctx, GetObjectPermissionsRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
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
func (a *PermissionsAPI) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error) {
	return a.impl.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// Set permissions
//
// Sets permissions on object. Objects can inherit permissions from their parent
// objects and root objects.
func (a *PermissionsAPI) SetObjectPermissions(ctx context.Context, request SetObjectPermissions) error {
	return a.impl.SetObjectPermissions(ctx, request)
}

// Update permission
//
// Updates the permissions on an object.
func (a *PermissionsAPI) UpdateObjectPermissions(ctx context.Context, request UpdateObjectPermissions) error {
	return a.impl.UpdateObjectPermissions(ctx, request)
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
