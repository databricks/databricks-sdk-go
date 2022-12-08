// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

import "fmt"

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
	WorkspaceId int64 `json:"-" url:"-"`
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the service principal.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// Get object permissions
type Get struct {
	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`
}

// Get permission levels
type GetPermissionLevels struct {
	// <needs content>
	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`
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

// String representation for [fmt.Print]
func (pl *PermissionLevel) String() string {
	return string(*pl)
}

// Set raw string value and validate it against allowed values
func (pl *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_BIND`, `CAN_EDIT`, `CAN_EDIT_METADATA`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_RUN`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`, `CAN_RESTART`, `CAN_RUN`, `CAN_USE`, `CAN_VIEW`, `CAN_VIEW_METADATA`, `IS_OWNER`:
		*pl = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_BIND", "CAN_EDIT", "CAN_EDIT_METADATA", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_RUN", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ", "CAN_RESTART", "CAN_RUN", "CAN_USE", "CAN_VIEW", "CAN_VIEW_METADATA", "IS_OWNER"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (pl *PermissionLevel) Type() string {
	return "PermissionLevel"
}

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

	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`
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
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type WorkspaceAssignmentsCreated struct {
	// Array of permissions assignments applied to a workspace.
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`
}

type WorkspacePermission string

const WorkspacePermissionAdmin WorkspacePermission = `ADMIN`

const WorkspacePermissionUnknown WorkspacePermission = `UNKNOWN`

const WorkspacePermissionUser WorkspacePermission = `USER`

// String representation for [fmt.Print]
func (wp *WorkspacePermission) String() string {
	return string(*wp)
}

// Set raw string value and validate it against allowed values
func (wp *WorkspacePermission) Set(v string) error {
	switch v {
	case `ADMIN`, `UNKNOWN`, `USER`:
		*wp = WorkspacePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADMIN", "UNKNOWN", "USER"`, v)
	}
}

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (wp *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions []PermissionOutput `json:"permissions,omitempty"`
}
