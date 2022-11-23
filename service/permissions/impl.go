// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just Permissions API methods
type permissionsImpl struct {
	client *client.DatabricksClient
}

func (a *permissionsImpl) GetObjectPermissions(ctx context.Context, request GetObjectPermissionsRequest) (*ObjectPermissions, error) {
	var objectPermissions ObjectPermissions
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &objectPermissions)
	return &objectPermissions, err
}

func (a *permissionsImpl) GetPermissionLevels(ctx context.Context, request GetPermissionLevelsRequest) (*GetPermissionLevelsResponse, error) {
	var getPermissionLevelsResponse GetPermissionLevelsResponse
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v/permissionLevels", request.RequestObjectType, request.RequestObjectId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getPermissionLevelsResponse)
	return &getPermissionLevelsResponse, err
}

func (a *permissionsImpl) SetObjectPermissions(ctx context.Context, request SetObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

func (a *permissionsImpl) UpdateObjectPermissions(ctx context.Context, request UpdateObjectPermissions) error {
	path := fmt.Sprintf("/api/2.0/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just WorkspaceAssignment API methods
type workspaceAssignmentImpl struct {
	client *client.DatabricksClient
}

func (a *workspaceAssignmentImpl) Create(ctx context.Context, request CreateWorkspaceAssignments) (*WorkspaceAssignmentsCreated, error) {
	var workspaceAssignmentsCreated WorkspaceAssignmentsCreated
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodPost, path, request, &workspaceAssignmentsCreated)
	return &workspaceAssignmentsCreated, err
}

func (a *workspaceAssignmentImpl) Delete(ctx context.Context, request DeleteRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *workspaceAssignmentImpl) Get(ctx context.Context, request GetRequest) (*WorkspacePermissions, error) {
	var workspacePermissions WorkspacePermissions
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/permissions", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &workspacePermissions)
	return &workspacePermissions, err
}

func (a *workspaceAssignmentImpl) List(ctx context.Context, request ListRequest) (*PermissionAssignments, error) {
	var permissionAssignments PermissionAssignments
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &permissionAssignments)
	return &permissionAssignments, err
}

func (a *workspaceAssignmentImpl) Update(ctx context.Context, request UpdateWorkspaceAssignments) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", a.client.ConfiguredAccountID(), request.WorkspaceId, request.PrincipalId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}
