// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

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

type ClusterAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type ClusterAccessControlResponse struct {
	AllPermissions []ClusterPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type ClusterGetPermissionLevelsResponse struct {
	PermissionLevels []ClusterPermissionsDescription `json:"permission_levels,omitempty"`
}

type ClusterObjectPermissions struct {
	AccessControlList []ClusterAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the cluster.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType ClusterObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type ClusterObjectPermissionsObjectType string

const ClusterObjectPermissionsObjectTypeCluster ClusterObjectPermissionsObjectType = `cluster`

// String representation for [fmt.Print]
func (copot *ClusterObjectPermissionsObjectType) String() string {
	return string(*copot)
}

// Set raw string value and validate it against allowed values
func (copot *ClusterObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `cluster`:
		*copot = ClusterObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "cluster"`, v)
	}
}

// Type always returns ClusterObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (copot *ClusterObjectPermissionsObjectType) Type() string {
	return "ClusterObjectPermissionsObjectType"
}

type ClusterPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`
}

type ClusterPermissionLevel string

const ClusterPermissionLevelCanAttachTo ClusterPermissionLevel = `CAN_ATTACH_TO`

const ClusterPermissionLevelCanRestart ClusterPermissionLevel = `CAN_RESTART`

const ClusterPermissionLevelCanUse ClusterPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (cpl *ClusterPermissionLevel) String() string {
	return string(*cpl)
}

// Set raw string value and validate it against allowed values
func (cpl *ClusterPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_RESTART`, `CAN_USE`:
		*cpl = ClusterPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_RESTART", "CAN_USE"`, v)
	}
}

// Type always returns ClusterPermissionLevel to satisfy [pflag.Value] interface
func (cpl *ClusterPermissionLevel) Type() string {
	return "ClusterPermissionLevel"
}

type ClusterPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`
}

type ClusterPermissionsRequest struct {
	AccessControlList []ClusterAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a cluster
	ClusterId string `json:"-" url:"-"`
}

type ClusterPolicyAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type ClusterPolicyAccessControlResponse struct {
	AllPermissions []ClusterPolicyPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type ClusterPolicyGetPermissionLevelsResponse struct {
	PermissionLevels []ClusterPolicyPermissionsDescription `json:"permission_levels,omitempty"`
}

type ClusterPolicyObjectPermissions struct {
	AccessControlList []ClusterPolicyAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the cluster policy.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType ClusterPolicyObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type ClusterPolicyObjectPermissionsObjectType string

const ClusterPolicyObjectPermissionsObjectTypeClusterPolicy ClusterPolicyObjectPermissionsObjectType = `cluster-policy`

// String representation for [fmt.Print]
func (cpopot *ClusterPolicyObjectPermissionsObjectType) String() string {
	return string(*cpopot)
}

// Set raw string value and validate it against allowed values
func (cpopot *ClusterPolicyObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `cluster-policy`:
		*cpopot = ClusterPolicyObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "cluster-policy"`, v)
	}
}

// Type always returns ClusterPolicyObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (cpopot *ClusterPolicyObjectPermissionsObjectType) Type() string {
	return "ClusterPolicyObjectPermissionsObjectType"
}

type ClusterPolicyPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
}

type ClusterPolicyPermissionLevel string

const ClusterPolicyPermissionLevelCanUse ClusterPolicyPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (cppl *ClusterPolicyPermissionLevel) String() string {
	return string(*cppl)
}

// Set raw string value and validate it against allowed values
func (cppl *ClusterPolicyPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*cppl = ClusterPolicyPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns ClusterPolicyPermissionLevel to satisfy [pflag.Value] interface
func (cppl *ClusterPolicyPermissionLevel) Type() string {
	return "ClusterPolicyPermissionLevel"
}

type ClusterPolicyPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
}

type ClusterPolicyPermissionsRequest struct {
	AccessControlList []ClusterPolicyAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of the cluster policy.
	ClusterPolicyId string `json:"-" url:"-"`
}

type ComplexValue struct {
	Display string `json:"display,omitempty"`

	Primary bool `json:"primary,omitempty"`

	Ref string `json:"$ref,omitempty"`

	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`
}

// Delete a group
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`
}

// Delete a service principal
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`
}

// Delete a user
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`
}

// Delete a group
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

// Delete a service principal
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

// Delete a user
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type DeltaLiveTableAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel DeltaLiveTablePermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type DeltaLiveTableAccessControlResponse struct {
	AllPermissions []DeltaLiveTablePermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type DeltaLiveTableGetPermissionLevelsResponse struct {
	PermissionLevels []DeltaLiveTablePermissionsDescription `json:"permission_levels,omitempty"`
}

type DeltaLiveTableObjectPermissions struct {
	AccessControlList []DeltaLiveTableAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the pipeline.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType DeltaLiveTableObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type DeltaLiveTableObjectPermissionsObjectType string

const DeltaLiveTableObjectPermissionsObjectTypePipeline DeltaLiveTableObjectPermissionsObjectType = `pipeline`

// String representation for [fmt.Print]
func (dltopot *DeltaLiveTableObjectPermissionsObjectType) String() string {
	return string(*dltopot)
}

// Set raw string value and validate it against allowed values
func (dltopot *DeltaLiveTableObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `pipeline`:
		*dltopot = DeltaLiveTableObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "pipeline"`, v)
	}
}

// Type always returns DeltaLiveTableObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (dltopot *DeltaLiveTableObjectPermissionsObjectType) Type() string {
	return "DeltaLiveTableObjectPermissionsObjectType"
}

type DeltaLiveTablePermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel DeltaLiveTablePermissionLevel `json:"permission_level,omitempty"`
}

type DeltaLiveTablePermissionLevel string

const DeltaLiveTablePermissionLevelCanManage DeltaLiveTablePermissionLevel = `CAN_MANAGE`

const DeltaLiveTablePermissionLevelCanRun DeltaLiveTablePermissionLevel = `CAN_RUN`

const DeltaLiveTablePermissionLevelCanView DeltaLiveTablePermissionLevel = `CAN_VIEW`

const DeltaLiveTablePermissionLevelIsOwner DeltaLiveTablePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (dltpl *DeltaLiveTablePermissionLevel) String() string {
	return string(*dltpl)
}

// Set raw string value and validate it against allowed values
func (dltpl *DeltaLiveTablePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*dltpl = DeltaLiveTablePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Type always returns DeltaLiveTablePermissionLevel to satisfy [pflag.Value] interface
func (dltpl *DeltaLiveTablePermissionLevel) Type() string {
	return "DeltaLiveTablePermissionLevel"
}

type DeltaLiveTablePermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel DeltaLiveTablePermissionLevel `json:"permission_level,omitempty"`
}

type DeltaLiveTablePermissionsRequest struct {
	AccessControlList []DeltaLiveTableAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a pipeline.
	PipelineId string `json:"-" url:"-"`
}

type DirectoryAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel DirectoryPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type DirectoryAccessControlResponse struct {
	AllPermissions []DirectoryPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type DirectoryGetPermissionLevelsResponse struct {
	PermissionLevels []DirectoryPermissionsDescription `json:"permission_levels,omitempty"`
}

type DirectoryObjectPermissions struct {
	AccessControlList []DirectoryAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the directory.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType DirectoryObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type DirectoryObjectPermissionsObjectType string

const DirectoryObjectPermissionsObjectTypeDirectory DirectoryObjectPermissionsObjectType = `directory`

// String representation for [fmt.Print]
func (dopot *DirectoryObjectPermissionsObjectType) String() string {
	return string(*dopot)
}

// Set raw string value and validate it against allowed values
func (dopot *DirectoryObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `directory`:
		*dopot = DirectoryObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "directory"`, v)
	}
}

// Type always returns DirectoryObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (dopot *DirectoryObjectPermissionsObjectType) Type() string {
	return "DirectoryObjectPermissionsObjectType"
}

type DirectoryPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel DirectoryPermissionLevel `json:"permission_level,omitempty"`
}

type DirectoryPermissionLevel string

const DirectoryPermissionLevelCanEdit DirectoryPermissionLevel = `CAN_EDIT`

const DirectoryPermissionLevelCanManage DirectoryPermissionLevel = `CAN_MANAGE`

const DirectoryPermissionLevelCanRead DirectoryPermissionLevel = `CAN_READ`

const DirectoryPermissionLevelCanRun DirectoryPermissionLevel = `CAN_RUN`

// String representation for [fmt.Print]
func (dpl *DirectoryPermissionLevel) String() string {
	return string(*dpl)
}

// Set raw string value and validate it against allowed values
func (dpl *DirectoryPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`, `CAN_RUN`:
		*dpl = DirectoryPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ", "CAN_RUN"`, v)
	}
}

// Type always returns DirectoryPermissionLevel to satisfy [pflag.Value] interface
func (dpl *DirectoryPermissionLevel) Type() string {
	return "DirectoryPermissionLevel"
}

type DirectoryPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel DirectoryPermissionLevel `json:"permission_level,omitempty"`
}

type DirectoryPermissionsRequest struct {
	AccessControlList []DirectoryAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a directory.
	JobId string `json:"-" url:"-"`
}

// Get group details
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`
}

// Get service principal details
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`
}

// Get user details
type GetAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`
}

// List assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name of the rule set to get or update.
	Name string `json:"-" url:"name"`
}

type GetAssignableRolesForResourceResponse struct {
	Roles []string `json:"roles,omitempty"`
}

// Get group details
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

// Get permission levels
type GetPermissionLevelsRequest struct {
	// The ID of a cluster policy
	ClusterPolicyId string `json:"-" url:"-"`
}

// Get cluster policy permissions.
type GetPermissionsClusterPolicyRequest struct {
	// The ID of the cluster policy.
	ClusterPolicyId string `json:"-" url:"-"`
}

// Get cluster permissions
type GetPermissionsClusterRequest struct {
	// The ID of a cluster
	ClusterId string `json:"-" url:"-"`
}

// Get pipeline permissions
type GetPermissionsDeltaLiveTableRequest struct {
	// The ID of a pipeline.
	PipelineId string `json:"-" url:"-"`
}

// Get directory permissions
type GetPermissionsDirectoryRequest struct {
	DirectoryId string `json:"-" url:"-"`
	// The ID of a directory.
	JobId string `json:"-" url:"-"`
}

// Get job permissions
type GetPermissionsJobRequest struct {
	// The ID of a job.
	JobId string `json:"-" url:"-"`
}

// Get experiment permissions
type GetPermissionsMLflowExperimentRequest struct {
	// The ID of an experiment.
	ExperimentId string `json:"-" url:"-"`
}

// Get registered model permissions
type GetPermissionsMlFlowRegisteredModelRequest struct {
	// The ID of a registered model.
	RegisteredModelId string `json:"-" url:"-"`
}

// Get notebook permissions
type GetPermissionsNotebookRequest struct {
	// The ID of a notebook.
	NotebookId string `json:"-" url:"-"`
}

// Get instance pool permissions
type GetPermissionsPoolRequest struct {
	InstancePoolId string `json:"-" url:"-"`
	// The ID of an instance pool.
	RegisteredModelId string `json:"-" url:"-"`
}

// Get repo permissions
type GetPermissionsRepoRequest struct {
	// The ID of a repo.
	RepoId string `json:"-" url:"-"`
}

// Get SQL warehouse permissions
type GetPermissionsSqlWarehousRequest struct {
	// The ID of a warehouse.
	WarehouseId string `json:"-" url:"-"`
}

// Get a rule set
type GetRuleSetRequest struct {
	// Etag used for versioning. The response is at least as fresh as the eTag
	// provided. Etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a rule set from overwriting each
	// other. It is strongly suggested that systems make use of the etag in the
	// read -> modify -> write pattern to perform rule set updates in order to
	// avoid race conditions that is get an etag from a GET rule set request,
	// and pass it with the PUT update request to identify the rule set version
	// you are updating.
	Etag string `json:"-" url:"etag"`
	// The ruleset name associated with the request.
	Name string `json:"-" url:"name"`
}

// Get service principal details
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

// Get user details
type GetUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type GrantRule struct {
	// Principals this grant rule applies to.
	Principals []string `json:"principals,omitempty"`
	// Role that is assigned to the list of principals.
	Role string `json:"role"`
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`

	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks group ID
	Id string `json:"id,omitempty" url:"-"`

	Members []ComplexValue `json:"members,omitempty"`

	Roles []ComplexValue `json:"roles,omitempty"`
}

type JobAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type JobAccessControlResponse struct {
	AllPermissions []JobPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type JobGetPermissionLevelsResponse struct {
	PermissionLevels []JobPermissionsDescription `json:"permission_levels,omitempty"`
}

type JobObjectPermissions struct {
	AccessControlList []JobAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the job.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType JobObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type JobObjectPermissionsObjectType string

const JobObjectPermissionsObjectTypeJob JobObjectPermissionsObjectType = `job`

// String representation for [fmt.Print]
func (jopot *JobObjectPermissionsObjectType) String() string {
	return string(*jopot)
}

// Set raw string value and validate it against allowed values
func (jopot *JobObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `job`:
		*jopot = JobObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "job"`, v)
	}
}

// Type always returns JobObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (jopot *JobObjectPermissionsObjectType) Type() string {
	return "JobObjectPermissionsObjectType"
}

type JobPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`
}

type JobPermissionLevel string

const JobPermissionLevelCanManage JobPermissionLevel = `CAN_MANAGE`

const JobPermissionLevelCanManageRun JobPermissionLevel = `CAN_MANAGE_RUN`

const JobPermissionLevelCanView JobPermissionLevel = `CAN_VIEW`

const JobPermissionLevelIsOwner JobPermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (jpl *JobPermissionLevel) String() string {
	return string(*jpl)
}

// Set raw string value and validate it against allowed values
func (jpl *JobPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*jpl = JobPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MANAGE_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Type always returns JobPermissionLevel to satisfy [pflag.Value] interface
func (jpl *JobPermissionLevel) Type() string {
	return "JobPermissionLevel"
}

type JobPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`
}

type JobPermissionsRequest struct {
	AccessControlList []JobAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a job.
	JobId string `json:"-" url:"-"`
	// The ID of a notebook.
	NotebookId string `json:"-" url:"-"`
}

// List group details
type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int `json:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

// List service principals
type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int `json:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

// List users
type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int `json:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

// List group details
type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int `json:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []Group `json:"Resources,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []ServicePrincipal `json:"Resources,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`
}

// List service principals
type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int `json:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

type ListSortOrder string

const ListSortOrderAscending ListSortOrder = `ascending`

const ListSortOrderDescending ListSortOrder = `descending`

// String representation for [fmt.Print]
func (lso *ListSortOrder) String() string {
	return string(*lso)
}

// Set raw string value and validate it against allowed values
func (lso *ListSortOrder) Set(v string) error {
	switch v {
	case `ascending`, `descending`:
		*lso = ListSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ascending", "descending"`, v)
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (lso *ListSortOrder) Type() string {
	return "ListSortOrder"
}

// List users
type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int `json:"-" url:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" url:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

type ListUsersResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []User `json:"Resources,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`
}

type MLflowExperimentAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel MLflowExperimentPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type MLflowExperimentAccessControlResponse struct {
	AllPermissions []MLflowExperimentPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type MLflowExperimentGetPermissionLevelsResponse struct {
	PermissionLevels []MLflowExperimentPermissionsDescription `json:"permission_levels,omitempty"`
}

type MLflowExperimentObjectPermissions struct {
	AccessControlList []MLflowExperimentAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the experiment.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType MLflowExperimentObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type MLflowExperimentObjectPermissionsObjectType string

const MLflowExperimentObjectPermissionsObjectTypeExperiment MLflowExperimentObjectPermissionsObjectType = `experiment`

// String representation for [fmt.Print]
func (mleopot *MLflowExperimentObjectPermissionsObjectType) String() string {
	return string(*mleopot)
}

// Set raw string value and validate it against allowed values
func (mleopot *MLflowExperimentObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `experiment`:
		*mleopot = MLflowExperimentObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "experiment"`, v)
	}
}

// Type always returns MLflowExperimentObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (mleopot *MLflowExperimentObjectPermissionsObjectType) Type() string {
	return "MLflowExperimentObjectPermissionsObjectType"
}

type MLflowExperimentPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel MLflowExperimentPermissionLevel `json:"permission_level,omitempty"`
}

type MLflowExperimentPermissionLevel string

const MLflowExperimentPermissionLevelCanEdit MLflowExperimentPermissionLevel = `CAN_EDIT`

const MLflowExperimentPermissionLevelCanManage MLflowExperimentPermissionLevel = `CAN_MANAGE`

const MLflowExperimentPermissionLevelCanRead MLflowExperimentPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (mlepl *MLflowExperimentPermissionLevel) String() string {
	return string(*mlepl)
}

// Set raw string value and validate it against allowed values
func (mlepl *MLflowExperimentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*mlepl = MLflowExperimentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Type always returns MLflowExperimentPermissionLevel to satisfy [pflag.Value] interface
func (mlepl *MLflowExperimentPermissionLevel) Type() string {
	return "MLflowExperimentPermissionLevel"
}

type MLflowExperimentPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel MLflowExperimentPermissionLevel `json:"permission_level,omitempty"`
}

type MLflowExperimentPermissionsRequest struct {
	AccessControlList []MLflowExperimentAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of an experiment.
	ExperimentId string `json:"-" url:"-"`
}

type MLflowRegisteredModelAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel MLflowRegisteredModelPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type MLflowRegisteredModelAccessControlResponse struct {
	AllPermissions []MLflowRegisteredModelPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type MLflowRegisteredModelGetPermissionLevelsResponse struct {
	PermissionLevels []MLflowRegisteredModelPermissionsDescription `json:"permission_levels,omitempty"`
}

type MLflowRegisteredModelObjectPermissions struct {
	AccessControlList []MLflowRegisteredModelAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the registered model.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType MLflowRegisteredModelObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type MLflowRegisteredModelObjectPermissionsObjectType string

const MLflowRegisteredModelObjectPermissionsObjectTypeRegisteredModel MLflowRegisteredModelObjectPermissionsObjectType = `registered-model`

// String representation for [fmt.Print]
func (mlrmopot *MLflowRegisteredModelObjectPermissionsObjectType) String() string {
	return string(*mlrmopot)
}

// Set raw string value and validate it against allowed values
func (mlrmopot *MLflowRegisteredModelObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `registered-model`:
		*mlrmopot = MLflowRegisteredModelObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "registered-model"`, v)
	}
}

// Type always returns MLflowRegisteredModelObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (mlrmopot *MLflowRegisteredModelObjectPermissionsObjectType) Type() string {
	return "MLflowRegisteredModelObjectPermissionsObjectType"
}

type MLflowRegisteredModelPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel MLflowRegisteredModelPermissionLevel `json:"permission_level,omitempty"`
}

type MLflowRegisteredModelPermissionLevel string

const MLflowRegisteredModelPermissionLevelCanEdit MLflowRegisteredModelPermissionLevel = `CAN_EDIT`

const MLflowRegisteredModelPermissionLevelCanManage MLflowRegisteredModelPermissionLevel = `CAN_MANAGE`

const MLflowRegisteredModelPermissionLevelCanManageProductionVersions MLflowRegisteredModelPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const MLflowRegisteredModelPermissionLevelCanManageStagingVersions MLflowRegisteredModelPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const MLflowRegisteredModelPermissionLevelCanRead MLflowRegisteredModelPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (mlrmpl *MLflowRegisteredModelPermissionLevel) String() string {
	return string(*mlrmpl)
}

// Set raw string value and validate it against allowed values
func (mlrmpl *MLflowRegisteredModelPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*mlrmpl = MLflowRegisteredModelPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Type always returns MLflowRegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (mlrmpl *MLflowRegisteredModelPermissionLevel) Type() string {
	return "MLflowRegisteredModelPermissionLevel"
}

type MLflowRegisteredModelPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel MLflowRegisteredModelPermissionLevel `json:"permission_level,omitempty"`
}

type MLflowRegisteredModelPermissionsRequest struct {
	AccessControlList []MLflowRegisteredModelAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a registered model.
	RegisteredModelId string `json:"-" url:"-"`
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName string `json:"familyName,omitempty"`
	// Given name of the Databricks user.
	GivenName string `json:"givenName,omitempty"`
}

type PartialUpdate struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`

	Operations []Patch `json:"operations,omitempty"`
}

type PasswordAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type PasswordAccessControlResponse struct {
	AllPermissions []PasswordPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type PasswordGetPermissionLevelsResponse struct {
	PermissionLevels []PasswordPermissionsDescription `json:"permission_levels,omitempty"`
}

type PasswordObjectPermissions struct {
	AccessControlList []PasswordAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType PasswordObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type PasswordObjectPermissionsObjectType string

const PasswordObjectPermissionsObjectTypePassword PasswordObjectPermissionsObjectType = `password`

// String representation for [fmt.Print]
func (popot *PasswordObjectPermissionsObjectType) String() string {
	return string(*popot)
}

// Set raw string value and validate it against allowed values
func (popot *PasswordObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `password`:
		*popot = PasswordObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "password"`, v)
	}
}

// Type always returns PasswordObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (popot *PasswordObjectPermissionsObjectType) Type() string {
	return "PasswordObjectPermissionsObjectType"
}

type PasswordPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
}

type PasswordPermissionLevel string

const PasswordPermissionLevelCanManage PasswordPermissionLevel = `CAN_MANAGE`

const PasswordPermissionLevelCanUse PasswordPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (ppl *PasswordPermissionLevel) String() string {
	return string(*ppl)
}

// Set raw string value and validate it against allowed values
func (ppl *PasswordPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`:
		*ppl = PasswordPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE"`, v)
	}
}

// Type always returns PasswordPermissionLevel to satisfy [pflag.Value] interface
func (ppl *PasswordPermissionLevel) Type() string {
	return "PasswordPermissionLevel"
}

type PasswordPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
}

type PasswordPermissionsRequest struct {
	AccessControlList []PasswordAccessControlRequest `json:"access_control_list,omitempty"`
}

type Patch struct {
	// Type of patch operation.
	Op PatchOp `json:"op,omitempty"`
	// Selection of patch operation
	Path string `json:"path,omitempty"`
	// Value to modify
	Value string `json:"value,omitempty"`
}

// Type of patch operation.
type PatchOp string

const PatchOpAdd PatchOp = `add`

const PatchOpRemove PatchOp = `remove`

const PatchOpReplace PatchOp = `replace`

// String representation for [fmt.Print]
func (po *PatchOp) String() string {
	return string(*po)
}

// Set raw string value and validate it against allowed values
func (po *PatchOp) Set(v string) error {
	switch v {
	case `add`, `remove`, `replace`:
		*po = PatchOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "add", "remove", "replace"`, v)
	}
}

// Type always returns PatchOp to satisfy [pflag.Value] interface
func (po *PatchOp) Type() string {
	return "PatchOp"
}

type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `json:"error,omitempty"`
	// The permissions level of the principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// Information about the principal assigned to the workspace.
	Principal *PrincipalOutput `json:"principal,omitempty"`
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

type PoolAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel PoolPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type PoolAccessControlResponse struct {
	AllPermissions []PoolPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type PoolGetPermissionLevelsResponse struct {
	PermissionLevels []PoolPermissionsDescription `json:"permission_levels,omitempty"`
}

type PoolObjectPermissions struct {
	AccessControlList []PoolAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the instance pool.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType PoolObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type PoolObjectPermissionsObjectType string

const PoolObjectPermissionsObjectTypeInstancePool PoolObjectPermissionsObjectType = `instance-pool`

// String representation for [fmt.Print]
func (popot *PoolObjectPermissionsObjectType) String() string {
	return string(*popot)
}

// Set raw string value and validate it against allowed values
func (popot *PoolObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `instance-pool`:
		*popot = PoolObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "instance-pool"`, v)
	}
}

// Type always returns PoolObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (popot *PoolObjectPermissionsObjectType) Type() string {
	return "PoolObjectPermissionsObjectType"
}

type PoolPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel PoolPermissionLevel `json:"permission_level,omitempty"`
}

type PoolPermissionLevel string

const PoolPermissionLevelCanAttachTo PoolPermissionLevel = `CAN_ATTACH_TO`

const PoolPermissionLevelCanManage PoolPermissionLevel = `CAN_MANAGE`

// String representation for [fmt.Print]
func (ppl *PoolPermissionLevel) String() string {
	return string(*ppl)
}

// Set raw string value and validate it against allowed values
func (ppl *PoolPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_MANAGE`:
		*ppl = PoolPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_MANAGE"`, v)
	}
}

// Type always returns PoolPermissionLevel to satisfy [pflag.Value] interface
func (ppl *PoolPermissionLevel) Type() string {
	return "PoolPermissionLevel"
}

type PoolPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PoolPermissionLevel `json:"permission_level,omitempty"`
}

type PoolPermissionsRequest struct {
	AccessControlList []PoolAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of an instance pool.
	RegisteredModelId string `json:"-" url:"-"`
}

type PrincipalOutput struct {
	// The display name of the principal.
	DisplayName string `json:"display_name,omitempty"`
	// The group name of the groupl. Present only if the principal is a group.
	GroupName string `json:"group_name,omitempty"`
	// The unique, opaque id of the principal.
	PrincipalId int64 `json:"principal_id,omitempty"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The username of the user. Present only if the principal is a user.
	UserName string `json:"user_name,omitempty"`
}

type RepoAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type RepoAccessControlResponse struct {
	AllPermissions []RepoPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type RepoGetPermissionLevelsResponse struct {
	PermissionLevels []RepoPermissionsDescription `json:"permission_levels,omitempty"`
}

type RepoObjectPermissions struct {
	AccessControlList []RepoAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the repo.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType RepoObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type RepoObjectPermissionsObjectType string

const RepoObjectPermissionsObjectTypeRepo RepoObjectPermissionsObjectType = `repo`

// String representation for [fmt.Print]
func (ropot *RepoObjectPermissionsObjectType) String() string {
	return string(*ropot)
}

// Set raw string value and validate it against allowed values
func (ropot *RepoObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `repo`:
		*ropot = RepoObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "repo"`, v)
	}
}

// Type always returns RepoObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (ropot *RepoObjectPermissionsObjectType) Type() string {
	return "RepoObjectPermissionsObjectType"
}

type RepoPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
}

type RepoPermissionLevel string

const RepoPermissionLevelCanEdit RepoPermissionLevel = `CAN_EDIT`

const RepoPermissionLevelCanManage RepoPermissionLevel = `CAN_MANAGE`

const RepoPermissionLevelCanRead RepoPermissionLevel = `CAN_READ`

const RepoPermissionLevelCanRun RepoPermissionLevel = `CAN_RUN`

// String representation for [fmt.Print]
func (rpl *RepoPermissionLevel) String() string {
	return string(*rpl)
}

// Set raw string value and validate it against allowed values
func (rpl *RepoPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`, `CAN_RUN`:
		*rpl = RepoPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ", "CAN_RUN"`, v)
	}
}

// Type always returns RepoPermissionLevel to satisfy [pflag.Value] interface
func (rpl *RepoPermissionLevel) Type() string {
	return "RepoPermissionLevel"
}

type RepoPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
}

type RepoPermissionsRequest struct {
	AccessControlList []RepoAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a repo.
	RepoId string `json:"-" url:"-"`
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned.
	Etag string `json:"etag,omitempty"`

	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name,omitempty"`
}

type RuleSetUpdateRequest struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	Etag string `json:"etag"`

	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name"`
}

type ServicePrincipal struct {
	// If this user is active
	Active bool `json:"active,omitempty"`
	// UUID relating to the service principal
	ApplicationId string `json:"applicationId,omitempty"`
	// String that represents a concatenation of given and family names.
	DisplayName string `json:"displayName,omitempty"`

	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks service principal ID.
	Id string `json:"id,omitempty" url:"-"`

	Roles []ComplexValue `json:"roles,omitempty"`
}

type SqlWarehouseAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel SqlWarehousePermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type SqlWarehouseAccessControlResponse struct {
	AllPermissions []SqlWarehousePermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type SqlWarehouseGetPermissionLevelsResponse struct {
	PermissionLevels []SqlWarehousePermissionsDescription `json:"permission_levels,omitempty"`
}

type SqlWarehouseObjectPermissions struct {
	AccessControlList []SqlWarehouseAccessControlResponse `json:"access_control_list,omitempty"`
	// The ID of the SQL warehouse.
	ObjectId string `json:"object_id,omitempty"`

	ObjectType SqlWarehouseObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type SqlWarehouseObjectPermissionsObjectType string

const SqlWarehouseObjectPermissionsObjectTypeWarehouse SqlWarehouseObjectPermissionsObjectType = `warehouse`

// String representation for [fmt.Print]
func (swopot *SqlWarehouseObjectPermissionsObjectType) String() string {
	return string(*swopot)
}

// Set raw string value and validate it against allowed values
func (swopot *SqlWarehouseObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `warehouse`:
		*swopot = SqlWarehouseObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "warehouse"`, v)
	}
}

// Type always returns SqlWarehouseObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (swopot *SqlWarehouseObjectPermissionsObjectType) Type() string {
	return "SqlWarehouseObjectPermissionsObjectType"
}

type SqlWarehousePermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel SqlWarehousePermissionLevel `json:"permission_level,omitempty"`
}

type SqlWarehousePermissionLevel string

const SqlWarehousePermissionLevelCanManage SqlWarehousePermissionLevel = `CAN_MANAGE`

const SqlWarehousePermissionLevelCanUse SqlWarehousePermissionLevel = `CAN_USE`

const SqlWarehousePermissionLevelIsOwner SqlWarehousePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (swpl *SqlWarehousePermissionLevel) String() string {
	return string(*swpl)
}

// Set raw string value and validate it against allowed values
func (swpl *SqlWarehousePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`, `IS_OWNER`:
		*swpl = SqlWarehousePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE", "IS_OWNER"`, v)
	}
}

// Type always returns SqlWarehousePermissionLevel to satisfy [pflag.Value] interface
func (swpl *SqlWarehousePermissionLevel) Type() string {
	return "SqlWarehousePermissionLevel"
}

type SqlWarehousePermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel SqlWarehousePermissionLevel `json:"permission_level,omitempty"`
}

type SqlWarehousePermissionsRequest struct {
	AccessControlList []SqlWarehouseAccessControlRequest `json:"access_control_list,omitempty"`
	// The ID of a warehouse.
	WarehouseId string `json:"-" url:"-"`
}

type TokenAccessControlRequest struct {
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type TokenAccessControlResponse struct {
	AllPermissions []TokenPermission `json:"all_permissions,omitempty"`
	// A customizable name for the user or service principal. This is not
	// defined for groups.
	DisplayName string `json:"display_name,omitempty"`
	// Group name. There are two built-in groups: `users` for all users, and
	// `admins` for administrators.
	GroupName string `json:"group_name,omitempty"`
	// Value that uniquely identifies a service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// Email address for a user.
	UserName string `json:"user_name,omitempty"`
}

type TokenGetPermissionLevelsResponse struct {
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

type TokenObjectPermissions struct {
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType TokenObjectPermissionsObjectType `json:"object_type,omitempty"`
}

type TokenObjectPermissionsObjectType string

const TokenObjectPermissionsObjectTypeToken TokenObjectPermissionsObjectType = `token`

// String representation for [fmt.Print]
func (topot *TokenObjectPermissionsObjectType) String() string {
	return string(*topot)
}

// Set raw string value and validate it against allowed values
func (topot *TokenObjectPermissionsObjectType) Set(v string) error {
	switch v {
	case `token`:
		*topot = TokenObjectPermissionsObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "token"`, v)
	}
}

// Type always returns TokenObjectPermissionsObjectType to satisfy [pflag.Value] interface
func (topot *TokenObjectPermissionsObjectType) Type() string {
	return "TokenObjectPermissionsObjectType"
}

type TokenPermission struct {
	// Specifies whether the permission is inherited from a parent ACL rather
	// than set explicitly. See related property `inherited_from_object`.
	Inherited bool `json:"inherited,omitempty"`
	// The list of parent ACL object IDs that contribute to inherited permission
	// on an ACL object. This is only defined if related property `inherited` is
	// set to `true`.
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
}

type TokenPermissionLevel string

const TokenPermissionLevelCanManage TokenPermissionLevel = `CAN_MANAGE`

const TokenPermissionLevelCanUse TokenPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (tpl *TokenPermissionLevel) String() string {
	return string(*tpl)
}

// Set raw string value and validate it against allowed values
func (tpl *TokenPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_USE`:
		*tpl = TokenPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_USE"`, v)
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (tpl *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name string `json:"name"`

	RuleSet RuleSetUpdateRequest `json:"rule_set"`
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace.
	Permissions []WorkspacePermission `json:"permissions"`
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type User struct {
	// If this user is active
	Active bool `json:"active,omitempty"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`.
	DisplayName string `json:"displayName,omitempty"`
	// All the emails associated with the Databricks user.
	Emails []ComplexValue `json:"emails,omitempty"`

	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks user ID.
	Id string `json:"id,omitempty" url:"-"`

	Name *Name `json:"name,omitempty"`

	Roles []ComplexValue `json:"roles,omitempty"`
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`
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
