// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import "fmt"

// all definitions in this file are in alphabetical order

type CreateIpAccessList struct {
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// This describes an enum
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
}

type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the ifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list.
	IpAccessListId string `json:"-" url:"-"`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
}

// Delete Personal Compute setting
type DeletePersonalComputeSettingRequest struct {
	// TBD
	Etag string `json:"-" url:"etag,omitempty"`
}

type DeletePersonalComputeSettingResponse struct {
	// TBD
	Etag string `json:"etag"`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
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
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
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
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById string `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`
}

type ListTokensResponse struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
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
	// TBD
	Value PersonalComputeMessageEnum `json:"value"`
}

// TBD
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
	// TBD
	Etag string `json:"etag,omitempty"`

	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	// Name of the corresponding setting. Needs to be 'default' if the setting
	// is a singleton.
	SettingName string `json:"setting_name,omitempty"`
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
}

// Get Personal Compute setting
type ReadPersonalComputeSettingRequest struct {
	// TBD
	Etag string `json:"-" url:"etag,omitempty"`
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
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`
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
}

// Update Personal Compute setting
type UpdatePersonalComputeSettingRequest struct {
	// TBD
	AllowMissing bool `json:"allow_missing,omitempty"`

	Setting *PersonalComputeSetting `json:"setting,omitempty"`
}

type WorkspaceConf map[string]string
