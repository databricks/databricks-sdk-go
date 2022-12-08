// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ipaccesslists

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

// Delete access list
type Delete struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
}

type FetchIpAccessListResponse struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

// Get access list
type Get struct {
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
}

type GetIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
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
	// Universally unique identifier(UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// This describes an enum
	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `json:"updated_by,omitempty"`
}

// This describes an enum
type ListType string

// An allow list. Include this IP or range.
const ListTypeAllow ListType = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.",
const ListTypeBlock ListType = `BLOCK`

// String representation for [fmt.Print]
func (lt *ListType) String() string {
	return string(*lt)
}

// Set raw string value and validate it against allowed values
func (lt *ListType) Set(v string) error {
	switch v {
	case `ALLOW`, `BLOCK`:
		*lt = ListType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW", "BLOCK"`, v)
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (lt *ListType) Type() string {
	return "ListType"
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
	// Universally unique identifier(UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// This describes an enum
	ListType ListType `json:"list_type"`
}

type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list to modify.
	IpAccessListId string `json:"-" url:"-"`
	// Array of IP addresses or CIDR values to be added to the IP access list.
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Universally unique identifier(UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// This describes an enum
	ListType ListType `json:"list_type,omitempty"`
}
