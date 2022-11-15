// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package scim

// all definitions in this file are in alphabetical order

type ComplexValue struct {
	Display string `json:"display,omitempty"`

	Primary bool `json:"primary,omitempty"`

	Ref string `json:"$ref,omitempty"`

	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`
}

type DeleteGroupRequest struct {
	// Unique ID for a group in the <Workspace>.
	Id string `json:"-" path:"id"`
}

type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the <Workspace>.
	Id string `json:"-" path:"id"`
}

type DeleteUserRequest struct {
	// Unique ID for a user in the <Workspace>.
	Id string `json:"-" path:"id"`
}

type GetGroupRequest struct {
	// Unique ID for a group in the <Workspace>.
	Id string `json:"-" path:"id"`
}

type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the <Workspace>.
	Id string `json:"-" path:"id"`
}

type GetUserRequest struct {
	// Unique ID for a user in the <Workspace>.
	Id string `json:"-" path:"id"`
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`

	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks group ID
	Id string `json:"id,omitempty" path:"id"`

	Members []ComplexValue `json:"members,omitempty"`

	Roles []ComplexValue `json:"roles,omitempty"`
}

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
	// `and` and `or`. The [SCIM
	// RFC](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) has more
	// details but we currently only support simple expressions.
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListGroupsSortOrder `json:"-" url:"sortOrder,omitempty"`
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

type ListGroupsSortOrder string

const ListGroupsSortOrderAscending ListGroupsSortOrder = `ascending`

const ListGroupsSortOrderDescending ListGroupsSortOrder = `descending`

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
	// `and` and `or`. The [SCIM
	// RFC](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) has more
	// details but we currently only support simple expressions.
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListServicePrincipalsSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`
}

type ListServicePrincipalsSortOrder string

const ListServicePrincipalsSortOrderAscending ListServicePrincipalsSortOrder = `ascending`

const ListServicePrincipalsSortOrderDescending ListServicePrincipalsSortOrder = `descending`

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
	// `and` and `or`. The [SCIM
	// RFC](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) has more
	// details but we currently only support simple expressions.
	Filter string `json:"-" url:"filter,omitempty"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder ListUsersSortOrder `json:"-" url:"sortOrder,omitempty"`
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

type ListUsersSortOrder string

const ListUsersSortOrderAscending ListUsersSortOrder = `ascending`

const ListUsersSortOrderDescending ListUsersSortOrder = `descending`

type Name struct {
	// Family name of the Databricks user.
	FamilyName string `json:"familyName,omitempty"`
	// Given name of the Databricks user.
	GivenName string `json:"givenName,omitempty"`
}

type PartialUpdate struct {
	// Unique ID for a group in the <Workspace>.
	Id string `json:"-" path:"id"`

	Operations []Patch `json:"operations,omitempty"`
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

type ServicePrincipal struct {
	// If this user is active
	Active bool `json:"active,omitempty"`
	// UUID relating to the service principal
	ApplicationId string `json:"applicationId,omitempty"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`.
	DisplayName string `json:"displayName,omitempty"`

	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks service principal ID.
	Id string `json:"id,omitempty" path:"id"`

	Roles []ComplexValue `json:"roles,omitempty"`
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
	Id string `json:"id,omitempty" path:"id"`

	Name *Name `json:"name,omitempty"`

	Roles []ComplexValue `json:"roles,omitempty"`
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`
}
