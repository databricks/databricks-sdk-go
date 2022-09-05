// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

// all definitions in this file are in alphabetical order

type CreateIPAccessListRequest struct {
	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`

	ListType ListType `json:"list_type"`
}

type CreateIPAccessListResponse struct {
	// Total number of IP or CIDR values.
	AddressCount int `json:"address_count,omitempty"`
	// Creation timestamp in milliseconds
	CreatedAt int64 `json:"created_at,omitempty"`
	// User ID of the user who created this list
	CreatedBy int64 `json:"created_by,omitempty"`
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// UUID of the IP access list
	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list
	UpdatedBy int64 `json:"updated_by,omitempty"`
}

type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string ` path:"ip_access_list_id"`
}

type FetchIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string ` path:"ip_access_list_id"`
}

type GetIPAccessListResponse struct {
	IpAccessLists []CreateIPAccessListResponse `json:"ip_access_lists,omitempty"`
}

type ReplaceIPAccessListRequest struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string ` path:"ip_access_list_id"`

	IpAddresses []string `json:"ip_addresses"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// UUID of the IP access list
	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type"`
}

type UpdateIPAccessListRequest struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string ` path:"ip_access_list_id"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// UUID of the IP access list
	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type,omitempty"`
}

// Array of IP addresses or CIDR values to be added to the IP access list.
type IpAddresses []string

// Type of IP access list. Valid values are as follows and are case-sensitive: *
// `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
type ListType string

// Type of IP access list. Valid values are as follows and are case-sensitive: *
// `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
const ListTypeAllow ListType = `ALLOW`

// Type of IP access list. Valid values are as follows and are case-sensitive: *
// `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
const ListTypeBlock ListType = `BLOCK`
