// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewPermissions(client *client.DatabricksClient) PermissionsService {
	return &PermissionsAPI{
		client: client,
	}
}

type PermissionsAPI struct {
	client *client.DatabricksClient
}

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetObjectPermissions(ctx context.Context, request GetObjectPermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Get(ctx, path, request, &objectPermissions)
	return &objectPermissions, err
}

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetObjectPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType string, objectId string) (*ObjectPermissions, error) {
	return a.GetObjectPermissions(ctx, GetObjectPermissionsRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
}

// Get permission levels
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Get(ctx, path, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

// Get permission levels
//
// Gets the permission levels that a user can have on an object.
func (a *PermissionsAPI) GetPermissionLevelsByRequestObjectTypeAndRequestObjectId(ctx context.Context, requestObjectType string, requestObjectId string) (*GetPermissionLevelsResponse, error) {
	return a.GetPermissionLevels(ctx, GetPermissionLevelsRequest{
		RequestObjectType: requestObjectType,
		RequestObjectId:   requestObjectId,
	})
}

// Set permissions
//
// Sets permissions on object. Objects can inherit permissions from their parent
// objects and root objects.
func (a *PermissionsAPI) SetObjectPermissions(ctx context.Context, request SetObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Put(ctx, path, request)
	return err
}

// Update permission
//
// Updates the permissions on an object.
func (a *PermissionsAPI) UpdateObjectPermissions(ctx context.Context, request UpdateObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewWorkspaceAssignment(client *client.DatabricksClient) WorkspaceAssignmentService {
	return &WorkspaceAssignmentAPI{
		client: client,
	}
}

type WorkspaceAssignmentAPI struct {
	client *client.DatabricksClient
}

// Create permission assignments
//
// Create new permission assignments for the specified account and workspace.
func (a *WorkspaceAssignmentAPI) Create(ctx context.Context, request CreateWorkspaceAssignments) (*WorkspaceAssignmentsCreated, error) {
	var workspaceAssignmentsCreated WorkspaceAssignmentsCreated
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Post(ctx, path, request, &workspaceAssignmentsCreated)
	return &workspaceAssignmentsCreated, err
}

// Delete permissions assignment
//
// Deletes the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.Config.AccountID, request.WorkspaceId, request.PrincipalId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete permissions assignment
//
// Deletes the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) DeleteByWorkspaceIdAndPrincipalId(ctx context.Context, workspaceId int64, principalId int64) error {
	return a.Delete(ctx, DeleteRequest{
		WorkspaceId: workspaceId,
		PrincipalId: principalId,
	})
}

// List workspace permissions
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) Get(ctx context.Context, request GetRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &workspacePermissions)
	return &workspacePermissions, err
}

// List workspace permissions
//
// Get an array of workspace permissions for the specified account and
// workspace.
func (a *WorkspaceAssignmentAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*WorkspacePermissions, error) {
	return a.Get(ctx, GetRequest{
		WorkspaceId: workspaceId,
	})
}

// Get permission assignments
//
// Get the permission assignments for the specified Databricks Account and
// Databricks Workspace.
//
// Use ListAll() to get all PermissionAssignment instances
func (a *WorkspaceAssignmentAPI) List(ctx context.Context, request ListRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &permissionAssignments)
	return &permissionAssignments, err
}

// ListAll returns all PermissionAssignment instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WorkspaceAssignmentAPI) ListAll(ctx context.Context, request ListRequest) ([]PermissionAssignment, error) {
	response, err := a.List(ctx, request)
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
	return a.List(ctx, ListRequest{
		WorkspaceId: workspaceId,
	})
}

// Update permissions assignment
//
// Updates the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.Config.AccountID, request.WorkspaceId, request.PrincipalId)
	err := a.client.Put(ctx, path, request)
	return err
}
