// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type CreateIpAccessList struct {
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`
}

type CreateIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId string `json:"application_id"`
	// Comment that describes the purpose of the token.
	Comment string `json:"comment,omitempty"`
	// The number of seconds before the token expires.
	LifetimeSeconds int64 `json:"lifetime_seconds"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the ifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Default namespace setting.
type DefaultNamespaceSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	Namespace StringMessage `json:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list.
	IpAccessListId string `json:"-" url:"-"`
}

// Delete the default namespace
type DeleteDefaultWorkspaceNamespaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag"`
}

type DeleteDefaultWorkspaceNamespaceResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag"`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
}

// Delete Personal Compute setting
type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag"`
}

type DeletePersonalComputeSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag"`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

type ExchangeToken struct {
	// The requested token.
	Credential string `json:"credential,omitempty"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime int64 `json:"credentialEolTime,omitempty"`
	// User ID of the user that owns this token.
	OwnerId int64 `json:"ownerId,omitempty"`
	// The scopes of access granted in the token.
	Scopes []string `json:"scopes,omitempty"`
	// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN`
	// is supported.
	TokenType TokenType `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ExchangeToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ExchangeTokenRequest struct {
	PartitionId PartitionId `json:"partitionId"`
	// Array of scopes for the token request.
	Scopes []string `json:"scopes"`

	TokenType []TokenType `json:"tokenType"`
}

type ExchangeTokenResponse struct {
	Values []ExchangeToken `json:"values,omitempty"`
}

type FetchIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list.
	IpAccessListId string `json:"-" url:"-"`
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
}

type GetIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

// Check configuration status
type GetStatusRequest struct {
	Keys string `json:"-" url:"keys"`
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount int `json:"address_count,omitempty"`
	// Creation timestamp in milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// User ID of the user who created this list.
	CreatedBy int64 `json:"created_by,omitempty"`
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IpAccessListInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById string `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTokensResponse struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

// Type of IP access list. Valid values are as follows and are case-sensitive:
//
// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
type ListType string

// An allow list. Include this IP or range.
const ListTypeAllow ListType = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.
const ListTypeBlock ListType = `BLOCK`

// String representation for [fmt.Print]
func (f *ListType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListType) Set(v string) error {
	switch v {
	case `ALLOW`, `BLOCK`:
		*f = ListType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW", "BLOCK"`, v)
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (f *ListType) Type() string {
	return "ListType"
}

type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PartitionId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum `json:"value"`
}

// ON: Grants all users in all workspaces access to the Personal Compute default
// policy, allowing all users to create single-machine compute resources.
// DELEGATE: Moves access control for the Personal Compute default policy to
// individual workspaces and requires a workspace’s users or groups to be
// added to the ACLs of that workspace’s Personal Compute default policy
// before they will be able to create compute resources through that policy.
type PersonalComputeMessageEnum string

const PersonalComputeMessageEnumDelegate PersonalComputeMessageEnum = `DELEGATE`

const PersonalComputeMessageEnumOn PersonalComputeMessageEnum = `ON`

// String representation for [fmt.Print]
func (f *PersonalComputeMessageEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalComputeMessageEnum) Set(v string) error {
	switch v {
	case `DELEGATE`, `ON`:
		*f = PersonalComputeMessageEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELEGATE", "ON"`, v)
	}
}

// Type always returns PersonalComputeMessageEnum to satisfy [pflag.Value] interface
func (f *PersonalComputeMessageEnum) Type() string {
	return "PersonalComputeMessageEnum"
}

type PersonalComputeSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	// Name of the corresponding setting. Needs to be 'default' if there is only
	// one setting instance per account.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment string `json:"comment,omitempty"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// The ID of this token.
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PublicTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublicTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the default namespace
type ReadDefaultWorkspaceNamespaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag"`
}

// Get Personal Compute setting
type ReadPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag"`
}

type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`

	ForceSendFields []string `json:"-"`
}

func (s *ReplaceIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReplaceIpAccessList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string `json:"comment,omitempty"`
	// User ID of the user that created the token.
	CreatedById int64 `json:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"created_by_username,omitempty"`
	// Timestamp when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Timestamp when the token expires.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// User ID of the user that owns the token.
	OwnerId int64 `json:"owner_id,omitempty"`
	// ID of the token.
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type TokenPermissionLevel string

const TokenPermissionLevelCanUse TokenPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *TokenPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = TokenPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (f *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissions struct {
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN` is
// supported.
type TokenType string

const TokenTypeAzureActiveDirectoryToken TokenType = `AZURE_ACTIVE_DIRECTORY_TOKEN`

// String representation for [fmt.Print]
func (f *TokenType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenType) Set(v string) error {
	switch v {
	case `AZURE_ACTIVE_DIRECTORY_TOKEN`:
		*f = TokenType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_ACTIVE_DIRECTORY_TOKEN"`, v)
	}
}

// Type always returns TokenType to satisfy [pflag.Value] interface
func (f *TokenType) Type() string {
	return "TokenType"
}

// Updates the default namespace setting
type UpdateDefaultWorkspaceNamespaceRequest struct {
	// This should always be set to true for Settings RPCs. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing,omitempty"`
	// Field mask required to be passed into the PATCH request. Field mask
	// specifies which fields of the setting payload will be updated. For
	// example, for Default Namespace setting, the field mask is supposed to
	// contain fields from the DefaultNamespaceSetting.namespace schema.
	//
	// The field mask needs to supplied as single string. To specify multiple
	// fields in the field mask, use comma as the seperator (no space).
	FieldMask string `json:"field_mask,omitempty"`
	// Default namespace setting.
	Setting *DefaultNamespaceSetting `json:"setting,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateDefaultWorkspaceNamespaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateDefaultWorkspaceNamespaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update Personal Compute setting
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings RPCs. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing,omitempty"`

	Setting *PersonalComputeSetting `json:"setting,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdatePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceConf map[string]string
