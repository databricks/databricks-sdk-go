// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package accountpermissionassignments

// all definitions in this file are in alphabetical order

type CreateWorkspacePermissionAssignmentsRequest struct {
	AccountId string ` path:"account_id"`

	PermissionAssignments []WorkspacePermissionAssignmentInput `json:"permission_assignments,omitempty"`

	WorkspaceId int64 ` path:"workspace_id"`
}

type CreateWorkspacePermissionAssignmentsResponse struct {
	PermissionAssignments []WorkspacePermissionAssignmentOutput `json:"permission_assignments,omitempty"`
}

type DeleteWorkspacePermissionAssignmentRequest struct {
	AccountId string ` path:"account_id"`

	PrincipalId int64 ` path:"principal_id"`

	WorkspaceId int64 ` path:"workspace_id"`
}

type GetWorkspacePermissionAssignmentsRequest struct {
	AccountId string ` path:"account_id"`

	WorkspaceId int64 ` path:"workspace_id"`
}

type GetWorkspacePermissionAssignmentsResponse struct {
	PermissionAssignments []WorkspacePermissionAssignmentOutput `json:"permission_assignments,omitempty"`
}

type ListWorkspacePermissionsRequest struct {
	AccountId string ` path:"account_id"`

	WorkspaceId int64 ` path:"workspace_id"`
}

type ListWorkspacePermissionsResponse struct {
	Permissions []PermissionOutput `json:"permissions,omitempty"`
}

type PermissionOutput struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PermissionOutputPermissionLevel `json:"permission_level,omitempty"`
}

type PermissionOutputPermissionLevel string

const PermissionOutputPermissionLevelUnknown PermissionOutputPermissionLevel = `UNKNOWN`

const PermissionOutputPermissionLevelUser PermissionOutputPermissionLevel = `USER`

const PermissionOutputPermissionLevelAdmin PermissionOutputPermissionLevel = `ADMIN`

type PrincipalOutput struct {
	DisplayName string `json:"display_name,omitempty"`

	GroupName string `json:"group_name,omitempty"`
	// The unique, opaque id of the principal.
	PrincipalId int64 `json:"principal_id,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

type UpdateWorkspacePermissionAssignmentRequest struct {
	AccountId string ` path:"account_id"`

	Permissions []UpdateWorkspacePermissionAssignmentRequestPermissionsItem `json:"permissions,omitempty"`

	PrincipalId int64 ` path:"principal_id"`

	WorkspaceId int64 ` path:"workspace_id"`
}

type UpdateWorkspacePermissionAssignmentRequestPermissionsItem string

const UpdateWorkspacePermissionAssignmentRequestPermissionsItemUnknown UpdateWorkspacePermissionAssignmentRequestPermissionsItem = `UNKNOWN`

const UpdateWorkspacePermissionAssignmentRequestPermissionsItemUser UpdateWorkspacePermissionAssignmentRequestPermissionsItem = `USER`

const UpdateWorkspacePermissionAssignmentRequestPermissionsItemAdmin UpdateWorkspacePermissionAssignmentRequestPermissionsItem = `ADMIN`

type WorkspacePermissionAssignmentInput struct {
	GroupName string `json:"group_name,omitempty"`

	Permissions []WorkspacePermissionAssignmentInputPermissionsItem `json:"permissions,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

type WorkspacePermissionAssignmentInputPermissionsItem string

const WorkspacePermissionAssignmentInputPermissionsItemUnknown WorkspacePermissionAssignmentInputPermissionsItem = `UNKNOWN`

const WorkspacePermissionAssignmentInputPermissionsItemUser WorkspacePermissionAssignmentInputPermissionsItem = `USER`

const WorkspacePermissionAssignmentInputPermissionsItemAdmin WorkspacePermissionAssignmentInputPermissionsItem = `ADMIN`

type WorkspacePermissionAssignmentOutput struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `json:"error,omitempty"`

	Permissions []WorkspacePermissionAssignmentOutputPermissionsItem `json:"permissions,omitempty"`

	Principal *PrincipalOutput `json:"principal,omitempty"`
}

type WorkspacePermissionAssignmentOutputPermissionsItem string

const WorkspacePermissionAssignmentOutputPermissionsItemUnknown WorkspacePermissionAssignmentOutputPermissionsItem = `UNKNOWN`

const WorkspacePermissionAssignmentOutputPermissionsItemUser WorkspacePermissionAssignmentOutputPermissionsItem = `USER`

const WorkspacePermissionAssignmentOutputPermissionsItemAdmin WorkspacePermissionAssignmentOutputPermissionsItem = `ADMIN`
