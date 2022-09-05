// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

// all definitions in this file are in alphabetical order

type AccessControlRequest struct {
	GroupName string `json:"group_name,omitempty"`
	// getPermissionLevel defaults to CAN_ATTACH_TO when it&#39;s not defined (for
	// PUT/PATCH) so you should use .permissionLevel.isDefined to verify that
	// it actually exists
	PermissionLevel AccessControlRequestPermissionLevel `json:"permission_level,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

// getPermissionLevel defaults to CAN_ATTACH_TO when it&#39;s not defined (for
// PUT/PATCH) so you should use .permissionLevel.isDefined to verify that it
// actually exists
type AccessControlRequestPermissionLevel string

const AccessControlRequestPermissionLevelCanManage AccessControlRequestPermissionLevel = `CAN_MANAGE`

const AccessControlRequestPermissionLevelCanRestart AccessControlRequestPermissionLevel = `CAN_RESTART`

const AccessControlRequestPermissionLevelCanAttachTo AccessControlRequestPermissionLevel = `CAN_ATTACH_TO`

const AccessControlRequestPermissionLevelIsOwner AccessControlRequestPermissionLevel = `IS_OWNER`

const AccessControlRequestPermissionLevelCanManageRun AccessControlRequestPermissionLevel = `CAN_MANAGE_RUN`

const AccessControlRequestPermissionLevelCanView AccessControlRequestPermissionLevel = `CAN_VIEW`

const AccessControlRequestPermissionLevelCanRead AccessControlRequestPermissionLevel = `CAN_READ`

const AccessControlRequestPermissionLevelCanRun AccessControlRequestPermissionLevel = `CAN_RUN`

const AccessControlRequestPermissionLevelCanEdit AccessControlRequestPermissionLevel = `CAN_EDIT`

const AccessControlRequestPermissionLevelCanUse AccessControlRequestPermissionLevel = `CAN_USE`

const AccessControlRequestPermissionLevelCanManageStagingVersions AccessControlRequestPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const AccessControlRequestPermissionLevelCanManageProductionVersions AccessControlRequestPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const AccessControlRequestPermissionLevelCanEditMetadata AccessControlRequestPermissionLevel = `CAN_EDIT_METADATA`

const AccessControlRequestPermissionLevelCanViewMetadata AccessControlRequestPermissionLevel = `CAN_VIEW_METADATA`

const AccessControlRequestPermissionLevelCanBind AccessControlRequestPermissionLevel = `CAN_BIND`

type AccessControlResponse struct {
	AllPermissions []Permission `json:"all_permissions,omitempty"`

	GroupName string `json:"group_name,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

type GetObjectPermissionsRequest struct {
	ObjectId string ` path:"object_id"`

	ObjectType string ` path:"object_type"`
}

type GetPermissionLevelsRequest struct {
	RequestObjectId string ` path:"request_object_id"`

	RequestObjectType string ` path:"request_object_type"`
}

type GetPermissionLevelsResponse struct {
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

type ObjectPermissions struct {
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`
}

type Permission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel PermissionPermissionLevel `json:"permission_level,omitempty"`
}

type PermissionPermissionLevel string

const PermissionPermissionLevelCanManage PermissionPermissionLevel = `CAN_MANAGE`

const PermissionPermissionLevelCanRestart PermissionPermissionLevel = `CAN_RESTART`

const PermissionPermissionLevelCanAttachTo PermissionPermissionLevel = `CAN_ATTACH_TO`

const PermissionPermissionLevelIsOwner PermissionPermissionLevel = `IS_OWNER`

const PermissionPermissionLevelCanManageRun PermissionPermissionLevel = `CAN_MANAGE_RUN`

const PermissionPermissionLevelCanView PermissionPermissionLevel = `CAN_VIEW`

const PermissionPermissionLevelCanRead PermissionPermissionLevel = `CAN_READ`

const PermissionPermissionLevelCanRun PermissionPermissionLevel = `CAN_RUN`

const PermissionPermissionLevelCanEdit PermissionPermissionLevel = `CAN_EDIT`

const PermissionPermissionLevelCanUse PermissionPermissionLevel = `CAN_USE`

const PermissionPermissionLevelCanManageStagingVersions PermissionPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionPermissionLevelCanManageProductionVersions PermissionPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionPermissionLevelCanEditMetadata PermissionPermissionLevel = `CAN_EDIT_METADATA`

const PermissionPermissionLevelCanViewMetadata PermissionPermissionLevel = `CAN_VIEW_METADATA`

const PermissionPermissionLevelCanBind PermissionPermissionLevel = `CAN_BIND`

type PermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PermissionsDescriptionPermissionLevel `json:"permission_level,omitempty"`
}

type PermissionsDescriptionPermissionLevel string

const PermissionsDescriptionPermissionLevelCanManage PermissionsDescriptionPermissionLevel = `CAN_MANAGE`

const PermissionsDescriptionPermissionLevelCanRestart PermissionsDescriptionPermissionLevel = `CAN_RESTART`

const PermissionsDescriptionPermissionLevelCanAttachTo PermissionsDescriptionPermissionLevel = `CAN_ATTACH_TO`

const PermissionsDescriptionPermissionLevelIsOwner PermissionsDescriptionPermissionLevel = `IS_OWNER`

const PermissionsDescriptionPermissionLevelCanManageRun PermissionsDescriptionPermissionLevel = `CAN_MANAGE_RUN`

const PermissionsDescriptionPermissionLevelCanView PermissionsDescriptionPermissionLevel = `CAN_VIEW`

const PermissionsDescriptionPermissionLevelCanRead PermissionsDescriptionPermissionLevel = `CAN_READ`

const PermissionsDescriptionPermissionLevelCanRun PermissionsDescriptionPermissionLevel = `CAN_RUN`

const PermissionsDescriptionPermissionLevelCanEdit PermissionsDescriptionPermissionLevel = `CAN_EDIT`

const PermissionsDescriptionPermissionLevelCanUse PermissionsDescriptionPermissionLevel = `CAN_USE`

const PermissionsDescriptionPermissionLevelCanManageStagingVersions PermissionsDescriptionPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionsDescriptionPermissionLevelCanManageProductionVersions PermissionsDescriptionPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionsDescriptionPermissionLevelCanEditMetadata PermissionsDescriptionPermissionLevel = `CAN_EDIT_METADATA`

const PermissionsDescriptionPermissionLevelCanViewMetadata PermissionsDescriptionPermissionLevel = `CAN_VIEW_METADATA`

const PermissionsDescriptionPermissionLevelCanBind PermissionsDescriptionPermissionLevel = `CAN_BIND`

type SetObjectPermissionsRequest struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty" path:"object_id"`

	ObjectType string `json:"object_type,omitempty" path:"object_type"`
}

type UpdateObjectPermissionsRequest struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty" path:"object_id"`

	ObjectType string `json:"object_type,omitempty" path:"object_type"`
}
