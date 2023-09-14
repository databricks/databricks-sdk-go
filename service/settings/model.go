// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type AccountNetworkPolicyMessage struct {
	// Whether or not serverless UDF can access the internet. When false, access
	// to the internet will be blocked from serverless clusters. Trusted traffic
	// required by clusters for basic functionality will not be affected.
	ServerlessInternetAccessEnabled bool `json:"serverless_internet_access_enabled,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AccountNetworkPolicyMessage) UnmarshalJSON(b []byte) error {
	type C AccountNetworkPolicyMessage
	return marshal.Unmarshal(b, (*C)(s))
}

func (s AccountNetworkPolicyMessage) MarshalJSON() ([]byte, error) {
	type C AccountNetworkPolicyMessage
	return marshal.Marshal((C)(s))
}

type CreateIpAccessList struct {
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// This describes an enum
	ListType ListType `json:"list_type"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateIpAccessList) UnmarshalJSON(b []byte) error {
	type C CreateIpAccessList
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateIpAccessList) MarshalJSON() ([]byte, error) {
	type C CreateIpAccessList
	return marshal.Marshal((C)(s))
}

type CreateIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateIpAccessListResponse) UnmarshalJSON(b []byte) error {
	type C CreateIpAccessListResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateIpAccessListResponse) MarshalJSON() ([]byte, error) {
	type C CreateIpAccessListResponse
	return marshal.Marshal((C)(s))
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
	type C CreateOboTokenRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	type C CreateOboTokenRequest
	return marshal.Marshal((C)(s))
}

type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	type C CreateOboTokenResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	type C CreateOboTokenResponse
	return marshal.Marshal((C)(s))
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
	type C CreateTokenRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateTokenRequest) MarshalJSON() ([]byte, error) {
	type C CreateTokenRequest
	return marshal.Marshal((C)(s))
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	type C CreateTokenResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateTokenResponse) MarshalJSON() ([]byte, error) {
	type C CreateTokenResponse
	return marshal.Marshal((C)(s))
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list.
	IpAccessListId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	type C DeleteAccountIpAccessListRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAccountIpAccessListRequest) MarshalJSON() ([]byte, error) {
	type C DeleteAccountIpAccessListRequest
	return marshal.Marshal((C)(s))
}

// Delete Account Network Policy
type DeleteAccountNetworkPolicyRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAccountNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	type C DeleteAccountNetworkPolicyRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAccountNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	type C DeleteAccountNetworkPolicyRequest
	return marshal.Marshal((C)(s))
}

type DeleteAccountNetworkPolicyResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAccountNetworkPolicyResponse) UnmarshalJSON(b []byte) error {
	type C DeleteAccountNetworkPolicyResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAccountNetworkPolicyResponse) MarshalJSON() ([]byte, error) {
	type C DeleteAccountNetworkPolicyResponse
	return marshal.Marshal((C)(s))
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteIpAccessListRequest) UnmarshalJSON(b []byte) error {
	type C DeleteIpAccessListRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteIpAccessListRequest) MarshalJSON() ([]byte, error) {
	type C DeleteIpAccessListRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *DeletePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	type C DeletePersonalComputeSettingRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeletePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	type C DeletePersonalComputeSettingRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *DeletePersonalComputeSettingResponse) UnmarshalJSON(b []byte) error {
	type C DeletePersonalComputeSettingResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeletePersonalComputeSettingResponse) MarshalJSON() ([]byte, error) {
	type C DeletePersonalComputeSettingResponse
	return marshal.Marshal((C)(s))
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteTokenManagementRequest) UnmarshalJSON(b []byte) error {
	type C DeleteTokenManagementRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteTokenManagementRequest) MarshalJSON() ([]byte, error) {
	type C DeleteTokenManagementRequest
	return marshal.Marshal((C)(s))
}

type FetchIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *FetchIpAccessListResponse) UnmarshalJSON(b []byte) error {
	type C FetchIpAccessListResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s FetchIpAccessListResponse) MarshalJSON() ([]byte, error) {
	type C FetchIpAccessListResponse
	return marshal.Marshal((C)(s))
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list.
	IpAccessListId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	type C GetAccountIpAccessListRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAccountIpAccessListRequest) MarshalJSON() ([]byte, error) {
	type C GetAccountIpAccessListRequest
	return marshal.Marshal((C)(s))
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetIpAccessListRequest) UnmarshalJSON(b []byte) error {
	type C GetIpAccessListRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetIpAccessListRequest) MarshalJSON() ([]byte, error) {
	type C GetIpAccessListRequest
	return marshal.Marshal((C)(s))
}

type GetIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetIpAccessListResponse) UnmarshalJSON(b []byte) error {
	type C GetIpAccessListResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetIpAccessListResponse) MarshalJSON() ([]byte, error) {
	type C GetIpAccessListResponse
	return marshal.Marshal((C)(s))
}

type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetIpAccessListsResponse) UnmarshalJSON(b []byte) error {
	type C GetIpAccessListsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetIpAccessListsResponse) MarshalJSON() ([]byte, error) {
	type C GetIpAccessListsResponse
	return marshal.Marshal((C)(s))
}

// Check configuration status
type GetStatusRequest struct {
	Keys string `json:"-" url:"keys"`

	ForceSendFields []string `json:"-"`
}

func (s *GetStatusRequest) UnmarshalJSON(b []byte) error {
	type C GetStatusRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetStatusRequest) MarshalJSON() ([]byte, error) {
	type C GetStatusRequest
	return marshal.Marshal((C)(s))
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetTokenManagementRequest) UnmarshalJSON(b []byte) error {
	type C GetTokenManagementRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetTokenManagementRequest) MarshalJSON() ([]byte, error) {
	type C GetTokenManagementRequest
	return marshal.Marshal((C)(s))
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetTokenPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	type C GetTokenPermissionLevelsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetTokenPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	type C GetTokenPermissionLevelsResponse
	return marshal.Marshal((C)(s))
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
	// This describes an enum
	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	type C IpAccessListInfo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s IpAccessListInfo) MarshalJSON() ([]byte, error) {
	type C IpAccessListInfo
	return marshal.Marshal((C)(s))
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
	type C ListTokenManagementRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	type C ListTokenManagementRequest
	return marshal.Marshal((C)(s))
}

type ListTokensResponse struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListTokensResponse) UnmarshalJSON(b []byte) error {
	type C ListTokensResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListTokensResponse) MarshalJSON() ([]byte, error) {
	type C ListTokensResponse
	return marshal.Marshal((C)(s))
}

// This describes an enum
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

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum `json:"value"`

	ForceSendFields []string `json:"-"`
}

func (s *PersonalComputeMessage) UnmarshalJSON(b []byte) error {
	type C PersonalComputeMessage
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PersonalComputeMessage) MarshalJSON() ([]byte, error) {
	type C PersonalComputeMessage
	return marshal.Marshal((C)(s))
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
	type C PersonalComputeSetting
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	type C PersonalComputeSetting
	return marshal.Marshal((C)(s))
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
	type C PublicTokenInfo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PublicTokenInfo) MarshalJSON() ([]byte, error) {
	type C PublicTokenInfo
	return marshal.Marshal((C)(s))
}

// Get Account Network Policy
type ReadAccountNetworkPolicyRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag"`

	ForceSendFields []string `json:"-"`
}

func (s *ReadAccountNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	type C ReadAccountNetworkPolicyRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ReadAccountNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	type C ReadAccountNetworkPolicyRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ReadPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	type C ReadPersonalComputeSettingRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ReadPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	type C ReadPersonalComputeSettingRequest
	return marshal.Marshal((C)(s))
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
	// This describes an enum
	ListType ListType `json:"list_type"`

	ForceSendFields []string `json:"-"`
}

func (s *ReplaceIpAccessList) UnmarshalJSON(b []byte) error {
	type C ReplaceIpAccessList
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ReplaceIpAccessList) MarshalJSON() ([]byte, error) {
	type C ReplaceIpAccessList
	return marshal.Marshal((C)(s))
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`

	ForceSendFields []string `json:"-"`
}

func (s *RevokeTokenRequest) UnmarshalJSON(b []byte) error {
	type C RevokeTokenRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RevokeTokenRequest) MarshalJSON() ([]byte, error) {
	type C RevokeTokenRequest
	return marshal.Marshal((C)(s))
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
	type C TokenAccessControlRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	type C TokenAccessControlRequest
	return marshal.Marshal((C)(s))
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
	type C TokenAccessControlResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	type C TokenAccessControlResponse
	return marshal.Marshal((C)(s))
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
	type C TokenInfo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	type C TokenInfo
	return marshal.Marshal((C)(s))
}

type TokenPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermission) UnmarshalJSON(b []byte) error {
	type C TokenPermission
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenPermission) MarshalJSON() ([]byte, error) {
	type C TokenPermission
	return marshal.Marshal((C)(s))
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
	type C TokenPermissions
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenPermissions) MarshalJSON() ([]byte, error) {
	type C TokenPermissions
	return marshal.Marshal((C)(s))
}

type TokenPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	type C TokenPermissionsDescription
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	type C TokenPermissionsDescription
	return marshal.Marshal((C)(s))
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenPermissionsRequest) UnmarshalJSON(b []byte) error {
	type C TokenPermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s TokenPermissionsRequest) MarshalJSON() ([]byte, error) {
	type C TokenPermissionsRequest
	return marshal.Marshal((C)(s))
}

// Update Account Network Policy
type UpdateAccountNetworkPolicyRequest struct {
	// This should always be set to true for Settings RPCs. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing,omitempty"`

	Setting *AccountNetworkPolicyMessage `json:"setting,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateAccountNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	type C UpdateAccountNetworkPolicyRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdateAccountNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	type C UpdateAccountNetworkPolicyRequest
	return marshal.Marshal((C)(s))
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
	// This describes an enum
	ListType ListType `json:"list_type"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	type C UpdateIpAccessList
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	type C UpdateIpAccessList
	return marshal.Marshal((C)(s))
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
	type C UpdatePersonalComputeSettingRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdatePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	type C UpdatePersonalComputeSettingRequest
	return marshal.Marshal((C)(s))
}

type WorkspaceConf map[string]string
