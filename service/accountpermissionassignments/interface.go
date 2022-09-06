// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package accountpermissionassignments

import (
	"context"
)

type AccountPermissionAssignmentsService interface {
	CreateWorkspacePermissionAssignments(ctx context.Context, createWorkspacePermissionAssignmentsRequest CreateWorkspacePermissionAssignmentsRequest) (*CreateWorkspacePermissionAssignmentsResponse, error)

	DeleteWorkspacePermissionAssignment(ctx context.Context, deleteWorkspacePermissionAssignmentRequest DeleteWorkspacePermissionAssignmentRequest) error
	DeleteWorkspacePermissionAssignmentByAccountIdAndWorkspaceIdAndPrincipalId(ctx context.Context, accountId string, workspaceId int64, principalId int64) error

	GetWorkspacePermissionAssignments(ctx context.Context, getWorkspacePermissionAssignmentsRequest GetWorkspacePermissionAssignmentsRequest) (*GetWorkspacePermissionAssignmentsResponse, error)
	GetWorkspacePermissionAssignmentsByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*GetWorkspacePermissionAssignmentsResponse, error)

	ListWorkspacePermissions(ctx context.Context, listWorkspacePermissionsRequest ListWorkspacePermissionsRequest) (*ListWorkspacePermissionsResponse, error)
	ListWorkspacePermissionsByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*ListWorkspacePermissionsResponse, error)

	UpdateWorkspacePermissionAssignment(ctx context.Context, updateWorkspacePermissionAssignmentRequest UpdateWorkspacePermissionAssignmentRequest) error
}
