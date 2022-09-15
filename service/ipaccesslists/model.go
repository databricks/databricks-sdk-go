// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

// all definitions in this file are in alphabetical order

type CreateIPAccessListRequest struct {
	IpAddresses []string `json:"ip_addresses"`

	Label string `json:"label"`

	ListType ListType `json:"list_type"`
}

type CreateIPAccessListResponse struct {
	AddressCount int `json:"address_count,omitempty"`

	CreatedAt int64 `json:"created_at,omitempty"`

	CreatedBy int64 `json:"created_by,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	IpAddresses []string `json:"ip_addresses,omitempty"`

	Label string `json:"label,omitempty"`

	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type,omitempty"`

	UpdatedAt int64 `json:"updated_at,omitempty"`

	UpdatedBy int64 `json:"updated_by,omitempty"`
}

type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" path:"ip_access_list_id"`
}

type FetchIpAccessListRequest struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" path:"ip_access_list_id"`
}

type GetIPAccessListResponse struct {
	IpAccessLists []CreateIPAccessListResponse `json:"ip_access_lists,omitempty"`
}

type ReplaceIPAccessListRequest struct {
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" path:"ip_access_list_id"`

	IpAddresses []string `json:"ip_addresses"`

	Label string `json:"label"`

	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type"`
}

type UpdateIPAccessListRequest struct {
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" path:"ip_access_list_id"`

	IpAddresses []string `json:"ip_addresses,omitempty"`

	Label string `json:"label,omitempty"`

	ListId string `json:"list_id,omitempty"`

	ListType ListType `json:"list_type,omitempty"`
}

// Total number of IP or CIDR values.

// Creation timestamp in milliseconds

// User ID of the user who created this list

// Specifies whether this IP access list is enabled.

// Label for the IP access list. This **cannot** be empty.

// UUID of the IP access list

// This describes an enum
type ListType string

// An allow list. Include this IP or range.
const ListTypeAllow ListType = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.
const ListTypeBlock ListType = `BLOCK`

// Update timestamp in milliseconds

// User ID of the user who updated this list
