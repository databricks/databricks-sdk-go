// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewPermissions(client *client.DatabricksClient) *PermissionsAPI {
	return &PermissionsAPI{
		PermissionsService: &permissionsAPI{
			client: client,
		},
	}
}

// Permissions API are used to create read, write, edit, update and manage
// access for various users on different objects and endpoints.
type PermissionsAPI struct {
	// PermissionsService contains low-level REST API interface.
	PermissionsService
}

// Get object permissions
//
// Gets the permission of an object. Objects can inherit permissions from their
// parent objects or root objects.
func (a *PermissionsAPI) GetObjectPermissions(ctx context.Context, request GetObjectPermissionsRequest) (*ObjectPermissions, error) {
	return a.PermissionsService.GetObjectPermissions(ctx, request)
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
	return a.PermissionsService.GetPermissionLevels(ctx, request)
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
	return a.PermissionsService.SetObjectPermissions(ctx, request)
}

// Update permission
//
// Updates the permissions on an object.
func (a *PermissionsAPI) UpdateObjectPermissions(ctx context.Context, request UpdateObjectPermissions) error {
	return a.PermissionsService.UpdateObjectPermissions(ctx, request)
}

// unexported type that holds implementations of just Permissions API methods
type permissionsAPI struct {
	client *client.DatabricksClient
}

func (a *permissionsAPI) GetObjectPermissions(ctx context.Context, request GetObjectPermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Get(ctx, path, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsAPI) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Get(ctx, path, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

func (a *permissionsAPI) SetObjectPermissions(ctx context.Context, request SetObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Put(ctx, path, request)
	return err
}

func (a *permissionsAPI) UpdateObjectPermissions(ctx context.Context, request UpdateObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewWorkspaceAssignment(client *client.DatabricksClient) *WorkspaceAssignmentAPI {
	return &WorkspaceAssignmentAPI{
		WorkspaceAssignmentService: &workspaceAssignmentAPI{
			client: client,
		},
	}
}

// Databricks Workspace Assignment REST API
type WorkspaceAssignmentAPI struct {
	// WorkspaceAssignmentService contains low-level REST API interface.
	WorkspaceAssignmentService
}

// Create permission assignments
//
// Create new permission assignments for the specified account and workspace.
func (a *WorkspaceAssignmentAPI) Create(ctx context.Context, request CreateWorkspaceAssignments) (*WorkspaceAssignmentsCreated, error) {
	return a.WorkspaceAssignmentService.Create(ctx, request)
}

// Delete permissions assignment
//
// Deletes the workspace permissions assignment for a given account and
// workspace using the specified service principal.
func (a *WorkspaceAssignmentAPI) Delete(ctx context.Context, request DeleteRequest) error {
	return a.WorkspaceAssignmentService.Delete(ctx, request)
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
	return a.WorkspaceAssignmentService.Get(ctx, request)
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
	return a.WorkspaceAssignmentService.Update(ctx, request)
}

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentAPI struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentAPI) Create(ctx context.Context, request CreateWorkspaceAssignments) (*WorkspaceAssignmentsCreated, error) {
	var workspaceAssignmentsCreated WorkspaceAssignmentsCreated
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Post(ctx, path, request, &workspaceAssignmentsCreated)
	return &workspaceAssignmentsCreated, err
}

func (a *workspaceAssignmentAPI) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.Config.AccountID, request.WorkspaceId, request.PrincipalId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *workspaceAssignmentAPI) Get(ctx context.Context, request GetRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &workspacePermissions)
	return &workspacePermissions, err
}

func (a *workspaceAssignmentAPI) List(ctx context.Context, request ListRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentAPI) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.Config.AccountID, request.WorkspaceId, request.PrincipalId)
	err := a.client.Put(ctx, path, request)
	return err
}
