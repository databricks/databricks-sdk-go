// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package accountpermissionassignments

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type AccountpermissionassignmentsService interface {
    
    CreateWorkspacePermissionAssignments(ctx context.Context, createWorkspacePermissionAssignmentsRequest CreateWorkspacePermissionAssignmentsRequest) (*CreateWorkspacePermissionAssignmentsResponse, error)
    
    DeleteWorkspacePermissionAssignment(ctx context.Context, deleteWorkspacePermissionAssignmentRequest DeleteWorkspacePermissionAssignmentRequest) error
    
    GetWorkspacePermissionAssignments(ctx context.Context, getWorkspacePermissionAssignmentsRequest GetWorkspacePermissionAssignmentsRequest) (*GetWorkspacePermissionAssignmentsResponse, error)
    
    ListWorkspacePermissions(ctx context.Context, listWorkspacePermissionsRequest ListWorkspacePermissionsRequest) (*ListWorkspacePermissionsResponse, error)
    
    UpdateWorkspacePermissionAssignment(ctx context.Context, updateWorkspacePermissionAssignmentRequest UpdateWorkspacePermissionAssignmentRequest) error
	GetWorkspacePermissionAssignmentsByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*GetWorkspacePermissionAssignmentsResponse, error)
	ListWorkspacePermissionsByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*ListWorkspacePermissionsResponse, error)
	DeleteWorkspacePermissionAssignmentByAccountIdAndWorkspaceIdAndPrincipalId(ctx context.Context, accountId string, workspaceId int64, principalId int64) error
}

func New(client *client.DatabricksClient) AccountpermissionassignmentsService {
	return &AccountpermissionassignmentsAPI{
		client: client,
	}
}

type AccountpermissionassignmentsAPI struct {
	client *client.DatabricksClient
}


func (a *AccountpermissionassignmentsAPI) CreateWorkspacePermissionAssignments(ctx context.Context, request CreateWorkspacePermissionAssignmentsRequest) (*CreateWorkspacePermissionAssignmentsResponse, error) {
	var createWorkspacePermissionAssignmentsResponse CreateWorkspacePermissionAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", request.AccountId, request.WorkspaceId)
	err := a.client.Post(ctx, path, request, &createWorkspacePermissionAssignmentsResponse)
	return &createWorkspacePermissionAssignmentsResponse, err
}


func (a *AccountpermissionassignmentsAPI) DeleteWorkspacePermissionAssignment(ctx context.Context, request DeleteWorkspacePermissionAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", request.AccountId, request.WorkspaceId, request.PrincipalId)
	err := a.client.Delete(ctx, path, request)
	return err
}


func (a *AccountpermissionassignmentsAPI) GetWorkspacePermissionAssignments(ctx context.Context, request GetWorkspacePermissionAssignmentsRequest) (*GetWorkspacePermissionAssignmentsResponse, error) {
	var getWorkspacePermissionAssignmentsResponse GetWorkspacePermissionAssignmentsResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments", request.AccountId, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &getWorkspacePermissionAssignmentsResponse)
	return &getWorkspacePermissionAssignmentsResponse, err
}


func (a *AccountpermissionassignmentsAPI) ListWorkspacePermissions(ctx context.Context, request ListWorkspacePermissionsRequest) (*ListWorkspacePermissionsResponse, error) {
	var listWorkspacePermissionsResponse ListWorkspacePermissionsResponse
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/permissions", request.AccountId, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &listWorkspacePermissionsResponse)
	return &listWorkspacePermissionsResponse, err
}


func (a *AccountpermissionassignmentsAPI) UpdateWorkspacePermissionAssignment(ctx context.Context, request UpdateWorkspacePermissionAssignmentRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/accounts/%v/workspaces/%v/permissionassignments/principals/%v", request.AccountId, request.WorkspaceId, request.PrincipalId)
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *AccountpermissionassignmentsAPI) GetWorkspacePermissionAssignmentsByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*GetWorkspacePermissionAssignmentsResponse, error) {
	return a.GetWorkspacePermissionAssignments(ctx, GetWorkspacePermissionAssignmentsRequest{
		AccountId: accountId,
		WorkspaceId: workspaceId,
	})
}

func (a *AccountpermissionassignmentsAPI) ListWorkspacePermissionsByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*ListWorkspacePermissionsResponse, error) {
	return a.ListWorkspacePermissions(ctx, ListWorkspacePermissionsRequest{
		AccountId: accountId,
		WorkspaceId: workspaceId,
	})
}

func (a *AccountpermissionassignmentsAPI) DeleteWorkspacePermissionAssignmentByAccountIdAndWorkspaceIdAndPrincipalId(ctx context.Context, accountId string, workspaceId int64, principalId int64) error {
	return a.DeleteWorkspacePermissionAssignment(ctx, DeleteWorkspacePermissionAssignmentRequest{
		AccountId: accountId,
		WorkspaceId: workspaceId,
		PrincipalId: principalId,
	})
}
