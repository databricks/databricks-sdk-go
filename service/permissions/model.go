// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

// all definitions in this file are in alphabetical order

type AccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`
}

type AccessControlResponse struct {
	// All permissions.
	AllPermissions []Permission `json:"all_permissions,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`
}

type CreateWorkspaceAssignments struct {
	// Array of permissions assignments to apply to a workspace.
	PermissionAssignments []PermissionAssignmentInput `json:"permission_assignments,omitempty"`
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the service principal.
	PrincipalId int64 `json:"-" path:"principal_id"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

// Get object permissions
type Get struct {
	RequestObjectId string `json:"-" path:"request_object_id"`
	// <needs content>
	RequestObjectType string `json:"-" path:"request_object_type"`
}

// Get permission levels
type GetPermissionLevels struct {
	// <needs content>
	RequestObjectId string `json:"-" path:"request_object_id"`
	// <needs content>
	RequestObjectType string `json:"-" path:"request_object_type"`
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

type ObjectPermissions struct {
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`
}

type Permission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
}

type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `json:"error,omitempty"`
	// The permissions level of the service principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// Information about the service principal assigned for the workspace.
	Principal *PrincipalOutput `json:"principal,omitempty"`
}

type PermissionAssignmentInput struct {
	// The group name for the service principal.
	GroupName string `json:"group_name,omitempty"`
	// Array of permissions to apply to the workspace for the service principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// The name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The username of the owner of the service principal.
	UserName string `json:"user_name,omitempty"`
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`
}

// Permission level
type PermissionLevel string

const PermissionLevelCanAttachTo PermissionLevel = `CAN_ATTACH_TO`

const PermissionLevelCanBind PermissionLevel = `CAN_BIND`

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanEditMetadata PermissionLevel = `CAN_EDIT_METADATA`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageRun PermissionLevel = `CAN_MANAGE_RUN`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

const PermissionLevelCanRestart PermissionLevel = `CAN_RESTART`

const PermissionLevelCanRun PermissionLevel = `CAN_RUN`

const PermissionLevelCanUse PermissionLevel = `CAN_USE`

const PermissionLevelCanView PermissionLevel = `CAN_VIEW`

const PermissionLevelCanViewMetadata PermissionLevel = `CAN_VIEW_METADATA`

const PermissionLevelIsOwner PermissionLevel = `IS_OWNER`

type PermissionOutput struct {
	// The results of a permissions query.
	Description string `json:"description,omitempty"`

	PermissionLevel WorkspacePermission `json:"permission_level,omitempty"`
}

type PermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
}

type PermissionsRequest struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`

	RequestObjectId string `json:"-" path:"request_object_id"`
	// <needs content>
	RequestObjectType string `json:"-" path:"request_object_type"`
}

type PrincipalOutput struct {
	// The display name of the service principal.
	DisplayName string `json:"display_name,omitempty"`
	// The group name for the service principal.
	GroupName string `json:"group_name,omitempty"`
	// The unique, opaque id of the principal.
	PrincipalId int64 `json:"principal_id,omitempty"`
	// The name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The username of the owner of the service principal.
	UserName string `json:"user_name,omitempty"`
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// The ID of the service principal.
	PrincipalId int64 `json:"-" path:"principal_id"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

type WorkspaceAssignmentsCreated struct {
	// Array of permissions assignments applied to a workspace.
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`
}

type WorkspacePermission string

const WorkspacePermissionAdmin WorkspacePermission = `ADMIN`

const WorkspacePermissionUnknown WorkspacePermission = `UNKNOWN`

const WorkspacePermissionUser WorkspacePermission = `USER`

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions []PermissionOutput `json:"permissions,omitempty"`
}
