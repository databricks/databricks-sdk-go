// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iampb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccessControlRequestPb struct {
	GroupName            string            `json:"group_name,omitempty"`
	PermissionLevel      PermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string            `json:"service_principal_name,omitempty"`
	UserName             string            `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AccessControlResponsePb struct {
	AllPermissions       []PermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string         `json:"display_name,omitempty"`
	GroupName            string         `json:"group_name,omitempty"`
	ServicePrincipalName string         `json:"service_principal_name,omitempty"`
	UserName             string         `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ActorPb struct {
	ActorId int64 `json:"actor_id,omitempty" url:"actor_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ActorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ActorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CheckPolicyRequestPb struct {
	Actor            ActorPb                `json:"-" url:"actor"`
	AuthzIdentity    RequestAuthzIdentityPb `json:"-" url:"authz_identity"`
	ConsistencyToken ConsistencyTokenPb     `json:"-" url:"consistency_token"`
	Permission       string                 `json:"-" url:"permission"`
	Resource         string                 `json:"-" url:"resource"`
	ResourceInfo     *ResourceInfoPb        `json:"-" url:"resource_info,omitempty"`
}

type CheckPolicyResponsePb struct {
	ConsistencyToken ConsistencyTokenPb `json:"consistency_token"`
	IsPermitted      bool               `json:"is_permitted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CheckPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CheckPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComplexValuePb struct {
	Display string `json:"display,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComplexValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComplexValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ConsistencyTokenPb struct {
	Value string `json:"value" url:"value"`
}

type DeleteAccountGroupRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteAccountServicePrincipalRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteAccountUserRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteGroupRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteServicePrincipalRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteUserRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteWorkspaceAssignmentRequestPb struct {
	PrincipalId int64 `json:"-" url:"-"`
	WorkspaceId int64 `json:"-" url:"-"`
}

type DeleteWorkspacePermissionAssignmentResponsePb struct {
}

type GetAccountGroupRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetAccountServicePrincipalRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetAccountUserRequestPb struct {
	Attributes         string         `json:"-" url:"attributes,omitempty"`
	Count              int            `json:"-" url:"count,omitempty"`
	ExcludedAttributes string         `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string         `json:"-" url:"filter,omitempty"`
	Id                 string         `json:"-" url:"-"`
	SortBy             string         `json:"-" url:"sortBy,omitempty"`
	SortOrder          GetSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int            `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetAccountUserRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetAccountUserRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAssignableRolesForResourceRequestPb struct {
	Resource string `json:"-" url:"resource"`
}

type GetAssignableRolesForResourceResponsePb struct {
	Roles []RolePb `json:"roles,omitempty"`
}

type GetGroupRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetPasswordPermissionLevelsRequestPb struct {
}

type GetPasswordPermissionLevelsResponsePb struct {
	PermissionLevels []PasswordPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetPasswordPermissionsRequestPb struct {
}

type GetPermissionLevelsRequestPb struct {
	RequestObjectId   string `json:"-" url:"-"`
	RequestObjectType string `json:"-" url:"-"`
}

type GetPermissionLevelsResponsePb struct {
	PermissionLevels []PermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetPermissionRequestPb struct {
	RequestObjectId   string `json:"-" url:"-"`
	RequestObjectType string `json:"-" url:"-"`
}

type GetRuleSetRequestPb struct {
	Etag string `json:"-" url:"etag"`
	Name string `json:"-" url:"name"`
}

type GetServicePrincipalRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetSortOrderPb string

const GetSortOrderAscending GetSortOrderPb = `ascending`

const GetSortOrderDescending GetSortOrderPb = `descending`

type GetUserRequestPb struct {
	Attributes         string         `json:"-" url:"attributes,omitempty"`
	Count              int            `json:"-" url:"count,omitempty"`
	ExcludedAttributes string         `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string         `json:"-" url:"filter,omitempty"`
	Id                 string         `json:"-" url:"-"`
	SortBy             string         `json:"-" url:"sortBy,omitempty"`
	SortOrder          GetSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int            `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetUserRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetUserRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetWorkspaceAssignmentRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

type GrantRulePb struct {
	Principals []string `json:"principals,omitempty"`
	Role       string   `json:"role"`
}

type GroupPb struct {
	DisplayName  string           `json:"displayName,omitempty"`
	Entitlements []ComplexValuePb `json:"entitlements,omitempty"`
	ExternalId   string           `json:"externalId,omitempty"`
	Groups       []ComplexValuePb `json:"groups,omitempty"`
	Id           string           `json:"id,omitempty"`
	Members      []ComplexValuePb `json:"members,omitempty"`
	Meta         *ResourceMetaPb  `json:"meta,omitempty"`
	Roles        []ComplexValuePb `json:"roles,omitempty"`
	Schemas      []GroupSchemaPb  `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GroupPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GroupPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GroupSchemaPb string

const GroupSchemaUrnIetfParamsScimSchemasCore20Group GroupSchemaPb = `urn:ietf:params:scim:schemas:core:2.0:Group`

type ListAccountGroupsRequestPb struct {
	Attributes         string          `json:"-" url:"attributes,omitempty"`
	Count              int64           `json:"-" url:"count,omitempty"`
	ExcludedAttributes string          `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string          `json:"-" url:"filter,omitempty"`
	SortBy             string          `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64           `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAccountGroupsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAccountGroupsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAccountServicePrincipalsRequestPb struct {
	Attributes         string          `json:"-" url:"attributes,omitempty"`
	Count              int64           `json:"-" url:"count,omitempty"`
	ExcludedAttributes string          `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string          `json:"-" url:"filter,omitempty"`
	SortBy             string          `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64           `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAccountServicePrincipalsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAccountServicePrincipalsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAccountUsersRequestPb struct {
	Attributes         string          `json:"-" url:"attributes,omitempty"`
	Count              int64           `json:"-" url:"count,omitempty"`
	ExcludedAttributes string          `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string          `json:"-" url:"filter,omitempty"`
	SortBy             string          `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64           `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAccountUsersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAccountUsersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListGroupsRequestPb struct {
	Attributes         string          `json:"-" url:"attributes,omitempty"`
	Count              int64           `json:"-" url:"count,omitempty"`
	ExcludedAttributes string          `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string          `json:"-" url:"filter,omitempty"`
	SortBy             string          `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64           `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListGroupsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListGroupsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListGroupsResponsePb struct {
	ItemsPerPage int64                  `json:"itemsPerPage,omitempty"`
	Resources    []GroupPb              `json:"Resources,omitempty"`
	Schemas      []ListResponseSchemaPb `json:"schemas,omitempty"`
	StartIndex   int64                  `json:"startIndex,omitempty"`
	TotalResults int64                  `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListGroupsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListGroupsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListResponseSchemaPb string

const ListResponseSchemaUrnIetfParamsScimApiMessages20ListResponse ListResponseSchemaPb = `urn:ietf:params:scim:api:messages:2.0:ListResponse`

type ListServicePrincipalResponsePb struct {
	ItemsPerPage int64                  `json:"itemsPerPage,omitempty"`
	Resources    []ServicePrincipalPb   `json:"Resources,omitempty"`
	Schemas      []ListResponseSchemaPb `json:"schemas,omitempty"`
	StartIndex   int64                  `json:"startIndex,omitempty"`
	TotalResults int64                  `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListServicePrincipalResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListServicePrincipalResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListServicePrincipalsRequestPb struct {
	Attributes         string          `json:"-" url:"attributes,omitempty"`
	Count              int64           `json:"-" url:"count,omitempty"`
	ExcludedAttributes string          `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string          `json:"-" url:"filter,omitempty"`
	SortBy             string          `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64           `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListServicePrincipalsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListServicePrincipalsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSortOrderPb string

const ListSortOrderAscending ListSortOrderPb = `ascending`

const ListSortOrderDescending ListSortOrderPb = `descending`

type ListUsersRequestPb struct {
	Attributes         string          `json:"-" url:"attributes,omitempty"`
	Count              int64           `json:"-" url:"count,omitempty"`
	ExcludedAttributes string          `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string          `json:"-" url:"filter,omitempty"`
	SortBy             string          `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrderPb `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64           `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListUsersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListUsersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListUsersResponsePb struct {
	ItemsPerPage int64                  `json:"itemsPerPage,omitempty"`
	Resources    []UserPb               `json:"Resources,omitempty"`
	Schemas      []ListResponseSchemaPb `json:"schemas,omitempty"`
	StartIndex   int64                  `json:"startIndex,omitempty"`
	TotalResults int64                  `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListUsersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListUsersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListWorkspaceAssignmentRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

type MeRequestPb struct {
}

type MigratePermissionsRequestPb struct {
	FromWorkspaceGroupName string `json:"from_workspace_group_name"`
	Size                   int    `json:"size,omitempty"`
	ToAccountGroupName     string `json:"to_account_group_name"`
	WorkspaceId            int64  `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MigratePermissionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MigratePermissionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MigratePermissionsResponsePb struct {
	PermissionsMigrated int `json:"permissions_migrated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MigratePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MigratePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NamePb struct {
	FamilyName string `json:"familyName,omitempty"`
	GivenName  string `json:"givenName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NamePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NamePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ObjectPermissionsPb struct {
	AccessControlList []AccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                    `json:"object_id,omitempty"`
	ObjectType        string                    `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ObjectPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ObjectPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PartialUpdatePb struct {
	Id         string          `json:"-" url:"-"`
	Operations []PatchPb       `json:"Operations,omitempty"`
	Schemas    []PatchSchemaPb `json:"schemas,omitempty"`
}

type PasswordAccessControlRequestPb struct {
	GroupName            string                    `json:"group_name,omitempty"`
	PermissionLevel      PasswordPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                    `json:"service_principal_name,omitempty"`
	UserName             string                    `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PasswordAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PasswordAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordAccessControlResponsePb struct {
	AllPermissions       []PasswordPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                 `json:"display_name,omitempty"`
	GroupName            string                 `json:"group_name,omitempty"`
	ServicePrincipalName string                 `json:"service_principal_name,omitempty"`
	UserName             string                 `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PasswordAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PasswordAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordPermissionPb struct {
	Inherited           bool                      `json:"inherited,omitempty"`
	InheritedFromObject []string                  `json:"inherited_from_object,omitempty"`
	PermissionLevel     PasswordPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PasswordPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PasswordPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordPermissionLevelPb string

const PasswordPermissionLevelCanUse PasswordPermissionLevelPb = `CAN_USE`

type PasswordPermissionsPb struct {
	AccessControlList []PasswordAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                            `json:"object_id,omitempty"`
	ObjectType        string                            `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PasswordPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PasswordPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordPermissionsDescriptionPb struct {
	Description     string                    `json:"description,omitempty"`
	PermissionLevel PasswordPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PasswordPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PasswordPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordPermissionsRequestPb struct {
	AccessControlList []PasswordAccessControlRequestPb `json:"access_control_list,omitempty"`
}

type PatchPb struct {
	Op    PatchOpPb `json:"op,omitempty"`
	Path  string    `json:"path,omitempty"`
	Value any       `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PatchPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PatchPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PatchOpPb string

const PatchOpAdd PatchOpPb = `add`

const PatchOpRemove PatchOpPb = `remove`

const PatchOpReplace PatchOpPb = `replace`

type PatchSchemaPb string

const PatchSchemaUrnIetfParamsScimApiMessages20PatchOp PatchSchemaPb = `urn:ietf:params:scim:api:messages:2.0:PatchOp`

type PermissionPb struct {
	Inherited           bool              `json:"inherited,omitempty"`
	InheritedFromObject []string          `json:"inherited_from_object,omitempty"`
	PermissionLevel     PermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionAssignmentPb struct {
	Error       string                  `json:"error,omitempty"`
	Permissions []WorkspacePermissionPb `json:"permissions,omitempty"`
	Principal   *PrincipalOutputPb      `json:"principal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PermissionAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PermissionAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionAssignmentsPb struct {
	PermissionAssignments []PermissionAssignmentPb `json:"permission_assignments,omitempty"`
}

type PermissionLevelPb string

const PermissionLevelCanAttachTo PermissionLevelPb = `CAN_ATTACH_TO`

const PermissionLevelCanBind PermissionLevelPb = `CAN_BIND`

const PermissionLevelCanCreate PermissionLevelPb = `CAN_CREATE`

const PermissionLevelCanEdit PermissionLevelPb = `CAN_EDIT`

const PermissionLevelCanEditMetadata PermissionLevelPb = `CAN_EDIT_METADATA`

const PermissionLevelCanManage PermissionLevelPb = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevelPb = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageRun PermissionLevelPb = `CAN_MANAGE_RUN`

const PermissionLevelCanManageStagingVersions PermissionLevelPb = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanMonitor PermissionLevelPb = `CAN_MONITOR`

const PermissionLevelCanMonitorOnly PermissionLevelPb = `CAN_MONITOR_ONLY`

const PermissionLevelCanQuery PermissionLevelPb = `CAN_QUERY`

const PermissionLevelCanRead PermissionLevelPb = `CAN_READ`

const PermissionLevelCanRestart PermissionLevelPb = `CAN_RESTART`

const PermissionLevelCanRun PermissionLevelPb = `CAN_RUN`

const PermissionLevelCanUse PermissionLevelPb = `CAN_USE`

const PermissionLevelCanView PermissionLevelPb = `CAN_VIEW`

const PermissionLevelCanViewMetadata PermissionLevelPb = `CAN_VIEW_METADATA`

const PermissionLevelIsOwner PermissionLevelPb = `IS_OWNER`

type PermissionOutputPb struct {
	Description     string                `json:"description,omitempty"`
	PermissionLevel WorkspacePermissionPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PermissionOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PermissionOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionsDescriptionPb struct {
	Description     string            `json:"description,omitempty"`
	PermissionLevel PermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PrincipalOutputPb struct {
	DisplayName          string `json:"display_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
	PrincipalId          int64  `json:"principal_id,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PrincipalOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PrincipalOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RequestAuthzIdentityPb string

const RequestAuthzIdentityRequestAuthzIdentityServiceIdentity RequestAuthzIdentityPb = `REQUEST_AUTHZ_IDENTITY_SERVICE_IDENTITY`

const RequestAuthzIdentityRequestAuthzIdentityUserContext RequestAuthzIdentityPb = `REQUEST_AUTHZ_IDENTITY_USER_CONTEXT`

type ResourceInfoPb struct {
	Id                 string          `json:"id" url:"id"`
	LegacyAclPath      string          `json:"legacy_acl_path,omitempty" url:"legacy_acl_path,omitempty"`
	ParentResourceInfo *ResourceInfoPb `json:"parent_resource_info,omitempty" url:"parent_resource_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResourceInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResourceInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResourceMetaPb struct {
	ResourceType string `json:"resourceType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResourceMetaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResourceMetaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RolePb struct {
	Name string `json:"name"`
}

type RuleSetResponsePb struct {
	Etag       string        `json:"etag"`
	GrantRules []GrantRulePb `json:"grant_rules,omitempty"`
	Name       string        `json:"name"`
}

type RuleSetUpdateRequestPb struct {
	Etag       string        `json:"etag"`
	GrantRules []GrantRulePb `json:"grant_rules,omitempty"`
	Name       string        `json:"name"`
}

type ServicePrincipalPb struct {
	Active        bool                       `json:"active,omitempty"`
	ApplicationId string                     `json:"applicationId,omitempty"`
	DisplayName   string                     `json:"displayName,omitempty"`
	Entitlements  []ComplexValuePb           `json:"entitlements,omitempty"`
	ExternalId    string                     `json:"externalId,omitempty"`
	Groups        []ComplexValuePb           `json:"groups,omitempty"`
	Id            string                     `json:"id,omitempty"`
	Roles         []ComplexValuePb           `json:"roles,omitempty"`
	Schemas       []ServicePrincipalSchemaPb `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServicePrincipalPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServicePrincipalPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServicePrincipalSchemaPb string

const ServicePrincipalSchemaUrnIetfParamsScimSchemasCore20ServicePrincipal ServicePrincipalSchemaPb = `urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal`

type SetObjectPermissionsPb struct {
	AccessControlList []AccessControlRequestPb `json:"access_control_list,omitempty"`
	RequestObjectId   string                   `json:"-" url:"-"`
	RequestObjectType string                   `json:"-" url:"-"`
}

type UpdateObjectPermissionsPb struct {
	AccessControlList []AccessControlRequestPb `json:"access_control_list,omitempty"`
	RequestObjectId   string                   `json:"-" url:"-"`
	RequestObjectType string                   `json:"-" url:"-"`
}

type UpdateRuleSetRequestPb struct {
	Name    string                 `json:"name"`
	RuleSet RuleSetUpdateRequestPb `json:"rule_set"`
}

type UpdateWorkspaceAssignmentsPb struct {
	Permissions []WorkspacePermissionPb `json:"permissions,omitempty"`
	PrincipalId int64                   `json:"-" url:"-"`
	WorkspaceId int64                   `json:"-" url:"-"`
}

type UserPb struct {
	Active       bool             `json:"active,omitempty"`
	DisplayName  string           `json:"displayName,omitempty"`
	Emails       []ComplexValuePb `json:"emails,omitempty"`
	Entitlements []ComplexValuePb `json:"entitlements,omitempty"`
	ExternalId   string           `json:"externalId,omitempty"`
	Groups       []ComplexValuePb `json:"groups,omitempty"`
	Id           string           `json:"id,omitempty"`
	Name         *NamePb          `json:"name,omitempty"`
	Roles        []ComplexValuePb `json:"roles,omitempty"`
	Schemas      []UserSchemaPb   `json:"schemas,omitempty"`
	UserName     string           `json:"userName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UserPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UserPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UserSchemaPb string

const UserSchemaUrnIetfParamsScimSchemasCore20User UserSchemaPb = `urn:ietf:params:scim:schemas:core:2.0:User`

const UserSchemaUrnIetfParamsScimSchemasExtensionWorkspace20User UserSchemaPb = `urn:ietf:params:scim:schemas:extension:workspace:2.0:User`

type WorkspacePermissionPb string

const WorkspacePermissionAdmin WorkspacePermissionPb = `ADMIN`

const WorkspacePermissionUnknown WorkspacePermissionPb = `UNKNOWN`

const WorkspacePermissionUser WorkspacePermissionPb = `USER`

type WorkspacePermissionsPb struct {
	Permissions []PermissionOutputPb `json:"permissions,omitempty"`
}
