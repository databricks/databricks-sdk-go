// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

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

	ForceSendFields []string `json:"-"`
}

func (s *AccessControlRequest) UnmarshalJSON(b []byte) error {
	type C AccessControlRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s AccessControlRequest) MarshalJSON() ([]byte, error) {
	type C AccessControlRequest
	return marshal.Marshal((C)(s))
}

type AccessControlResponse struct {
	// All permissions.
	AllPermissions []Permission `json:"all_permissions,omitempty"`
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

func (s *AccessControlResponse) UnmarshalJSON(b []byte) error {
	type C AccessControlResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s AccessControlResponse) MarshalJSON() ([]byte, error) {
	type C AccessControlResponse
	return marshal.Marshal((C)(s))
}

type ComplexValue struct {
	Display string `json:"display,omitempty"`

	Primary bool `json:"primary,omitempty"`

	Ref string `json:"$ref,omitempty"`

	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ComplexValue) UnmarshalJSON(b []byte) error {
	type C ComplexValue
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ComplexValue) MarshalJSON() ([]byte, error) {
	type C ComplexValue
	return marshal.Marshal((C)(s))
}

// Delete a group
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAccountGroupRequest) UnmarshalJSON(b []byte) error {
	type C DeleteAccountGroupRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAccountGroupRequest) MarshalJSON() ([]byte, error) {
	type C DeleteAccountGroupRequest
	return marshal.Marshal((C)(s))
}

// Delete a service principal
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAccountServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	type C DeleteAccountServicePrincipalRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAccountServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	type C DeleteAccountServicePrincipalRequest
	return marshal.Marshal((C)(s))
}

// Delete a user
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAccountUserRequest) UnmarshalJSON(b []byte) error {
	type C DeleteAccountUserRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAccountUserRequest) MarshalJSON() ([]byte, error) {
	type C DeleteAccountUserRequest
	return marshal.Marshal((C)(s))
}

// Delete a group
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteGroupRequest) UnmarshalJSON(b []byte) error {
	type C DeleteGroupRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteGroupRequest) MarshalJSON() ([]byte, error) {
	type C DeleteGroupRequest
	return marshal.Marshal((C)(s))
}

// Delete a service principal
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	type C DeleteServicePrincipalRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	type C DeleteServicePrincipalRequest
	return marshal.Marshal((C)(s))
}

// Delete a user
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteUserRequest) UnmarshalJSON(b []byte) error {
	type C DeleteUserRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteUserRequest) MarshalJSON() ([]byte, error) {
	type C DeleteUserRequest
	return marshal.Marshal((C)(s))
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	type C DeleteWorkspaceAssignmentRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	type C DeleteWorkspaceAssignmentRequest
	return marshal.Marshal((C)(s))
}

// Get group details
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAccountGroupRequest) UnmarshalJSON(b []byte) error {
	type C GetAccountGroupRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAccountGroupRequest) MarshalJSON() ([]byte, error) {
	type C GetAccountGroupRequest
	return marshal.Marshal((C)(s))
}

// Get service principal details
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAccountServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	type C GetAccountServicePrincipalRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAccountServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	type C GetAccountServicePrincipalRequest
	return marshal.Marshal((C)(s))
}

// Get user details
type GetAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAccountUserRequest) UnmarshalJSON(b []byte) error {
	type C GetAccountUserRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAccountUserRequest) MarshalJSON() ([]byte, error) {
	type C GetAccountUserRequest
	return marshal.Marshal((C)(s))
}

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	Resource string `json:"-" url:"resource"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAssignableRolesForResourceRequest) UnmarshalJSON(b []byte) error {
	type C GetAssignableRolesForResourceRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAssignableRolesForResourceRequest) MarshalJSON() ([]byte, error) {
	type C GetAssignableRolesForResourceRequest
	return marshal.Marshal((C)(s))
}

type GetAssignableRolesForResourceResponse struct {
	Roles []Role `json:"roles,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAssignableRolesForResourceResponse) UnmarshalJSON(b []byte) error {
	type C GetAssignableRolesForResourceResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAssignableRolesForResourceResponse) MarshalJSON() ([]byte, error) {
	type C GetAssignableRolesForResourceResponse
	return marshal.Marshal((C)(s))
}

// Get group details
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetGroupRequest) UnmarshalJSON(b []byte) error {
	type C GetGroupRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetGroupRequest) MarshalJSON() ([]byte, error) {
	type C GetGroupRequest
	return marshal.Marshal((C)(s))
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PasswordPermissionsDescription `json:"permission_levels,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPasswordPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	type C GetPasswordPermissionLevelsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetPasswordPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	type C GetPasswordPermissionLevelsResponse
	return marshal.Marshal((C)(s))
}

// Get object permission levels
type GetPermissionLevelsRequest struct {
	// <needs content>
	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	type C GetPermissionLevelsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	type C GetPermissionLevelsRequest
	return marshal.Marshal((C)(s))
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	type C GetPermissionLevelsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	type C GetPermissionLevelsResponse
	return marshal.Marshal((C)(s))
}

// Get object permissions
type GetPermissionRequest struct {
	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPermissionRequest) UnmarshalJSON(b []byte) error {
	type C GetPermissionRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetPermissionRequest) MarshalJSON() ([]byte, error) {
	type C GetPermissionRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *GetRuleSetRequest) UnmarshalJSON(b []byte) error {
	type C GetRuleSetRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetRuleSetRequest) MarshalJSON() ([]byte, error) {
	type C GetRuleSetRequest
	return marshal.Marshal((C)(s))
}

// Get service principal details
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	type C GetServicePrincipalRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	type C GetServicePrincipalRequest
	return marshal.Marshal((C)(s))
}

// Get user details
type GetUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetUserRequest) UnmarshalJSON(b []byte) error {
	type C GetUserRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetUserRequest) MarshalJSON() ([]byte, error) {
	type C GetUserRequest
	return marshal.Marshal((C)(s))
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	type C GetWorkspaceAssignmentRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	type C GetWorkspaceAssignmentRequest
	return marshal.Marshal((C)(s))
}

type GrantRule struct {
	// Principals this grant rule applies to.
	Principals []string `json:"principals,omitempty"`
	// Role that is assigned to the list of principals.
	Role string `json:"role"`

	ForceSendFields []string `json:"-"`
}

func (s *GrantRule) UnmarshalJSON(b []byte) error {
	type C GrantRule
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GrantRule) MarshalJSON() ([]byte, error) {
	type C GrantRule
	return marshal.Marshal((C)(s))
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
	// Container for the group identifier. Workspace local versus account.
	Meta *ResourceMeta `json:"meta,omitempty"`

	Roles []ComplexValue `json:"roles,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Group) UnmarshalJSON(b []byte) error {
	type C Group
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Group) MarshalJSON() ([]byte, error) {
	type C Group
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListAccountGroupsRequest) UnmarshalJSON(b []byte) error {
	type C ListAccountGroupsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	type C ListAccountGroupsRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListAccountServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	type C ListAccountServicePrincipalsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListAccountServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	type C ListAccountServicePrincipalsRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListAccountUsersRequest) UnmarshalJSON(b []byte) error {
	type C ListAccountUsersRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListAccountUsersRequest) MarshalJSON() ([]byte, error) {
	type C ListAccountUsersRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	type C ListGroupsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListGroupsRequest) MarshalJSON() ([]byte, error) {
	type C ListGroupsRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	type C ListGroupsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListGroupsResponse) MarshalJSON() ([]byte, error) {
	type C ListGroupsResponse
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	type C ListServicePrincipalResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListServicePrincipalResponse) MarshalJSON() ([]byte, error) {
	type C ListServicePrincipalResponse
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	type C ListServicePrincipalsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	type C ListServicePrincipalsRequest
	return marshal.Marshal((C)(s))
}

type ListSortOrder string

const ListSortOrderAscending ListSortOrder = `ascending`

const ListSortOrderDescending ListSortOrder = `descending`

// String representation for [fmt.Print]
func (f *ListSortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListSortOrder) Set(v string) error {
	switch v {
	case `ascending`, `descending`:
		*f = ListSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ascending", "descending"`, v)
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (f *ListSortOrder) Type() string {
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

	ForceSendFields []string `json:"-"`
}

func (s *ListUsersRequest) UnmarshalJSON(b []byte) error {
	type C ListUsersRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListUsersRequest) MarshalJSON() ([]byte, error) {
	type C ListUsersRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ListUsersResponse) UnmarshalJSON(b []byte) error {
	type C ListUsersResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListUsersResponse) MarshalJSON() ([]byte, error) {
	type C ListUsersResponse
	return marshal.Marshal((C)(s))
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *ListWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	type C ListWorkspaceAssignmentRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	type C ListWorkspaceAssignmentRequest
	return marshal.Marshal((C)(s))
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName string `json:"familyName,omitempty"`
	// Given name of the Databricks user.
	GivenName string `json:"givenName,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Name) UnmarshalJSON(b []byte) error {
	type C Name
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Name) MarshalJSON() ([]byte, error) {
	type C Name
	return marshal.Marshal((C)(s))
}

type ObjectPermissions struct {
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ObjectPermissions) UnmarshalJSON(b []byte) error {
	type C ObjectPermissions
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ObjectPermissions) MarshalJSON() ([]byte, error) {
	type C ObjectPermissions
	return marshal.Marshal((C)(s))
}

type PartialUpdate struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`

	Operations []Patch `json:"Operations,omitempty"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas []PatchSchema `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PartialUpdate) UnmarshalJSON(b []byte) error {
	type C PartialUpdate
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PartialUpdate) MarshalJSON() ([]byte, error) {
	type C PartialUpdate
	return marshal.Marshal((C)(s))
}

type PasswordAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PasswordAccessControlRequest) UnmarshalJSON(b []byte) error {
	type C PasswordAccessControlRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PasswordAccessControlRequest) MarshalJSON() ([]byte, error) {
	type C PasswordAccessControlRequest
	return marshal.Marshal((C)(s))
}

type PasswordAccessControlResponse struct {
	// All permissions.
	AllPermissions []PasswordPermission `json:"all_permissions,omitempty"`
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

func (s *PasswordAccessControlResponse) UnmarshalJSON(b []byte) error {
	type C PasswordAccessControlResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PasswordAccessControlResponse) MarshalJSON() ([]byte, error) {
	type C PasswordAccessControlResponse
	return marshal.Marshal((C)(s))
}

type PasswordPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PasswordPermission) UnmarshalJSON(b []byte) error {
	type C PasswordPermission
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PasswordPermission) MarshalJSON() ([]byte, error) {
	type C PasswordPermission
	return marshal.Marshal((C)(s))
}

// Permission level
type PasswordPermissionLevel string

const PasswordPermissionLevelCanUse PasswordPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *PasswordPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PasswordPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = PasswordPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns PasswordPermissionLevel to satisfy [pflag.Value] interface
func (f *PasswordPermissionLevel) Type() string {
	return "PasswordPermissionLevel"
}

type PasswordPermissions struct {
	AccessControlList []PasswordAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PasswordPermissions) UnmarshalJSON(b []byte) error {
	type C PasswordPermissions
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PasswordPermissions) MarshalJSON() ([]byte, error) {
	type C PasswordPermissions
	return marshal.Marshal((C)(s))
}

type PasswordPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PasswordPermissionsDescription) UnmarshalJSON(b []byte) error {
	type C PasswordPermissionsDescription
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PasswordPermissionsDescription) MarshalJSON() ([]byte, error) {
	type C PasswordPermissionsDescription
	return marshal.Marshal((C)(s))
}

type PasswordPermissionsRequest struct {
	AccessControlList []PasswordAccessControlRequest `json:"access_control_list,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PasswordPermissionsRequest) UnmarshalJSON(b []byte) error {
	type C PasswordPermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PasswordPermissionsRequest) MarshalJSON() ([]byte, error) {
	type C PasswordPermissionsRequest
	return marshal.Marshal((C)(s))
}

type Patch struct {
	// Type of patch operation.
	Op PatchOp `json:"op,omitempty"`
	// Selection of patch operation
	Path string `json:"path,omitempty"`
	// Value to modify
	Value any `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Patch) UnmarshalJSON(b []byte) error {
	type C Patch
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Patch) MarshalJSON() ([]byte, error) {
	type C Patch
	return marshal.Marshal((C)(s))
}

// Type of patch operation.
type PatchOp string

const PatchOpAdd PatchOp = `add`

const PatchOpRemove PatchOp = `remove`

const PatchOpReplace PatchOp = `replace`

// String representation for [fmt.Print]
func (f *PatchOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PatchOp) Set(v string) error {
	switch v {
	case `add`, `remove`, `replace`:
		*f = PatchOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "add", "remove", "replace"`, v)
	}
}

// Type always returns PatchOp to satisfy [pflag.Value] interface
func (f *PatchOp) Type() string {
	return "PatchOp"
}

type PatchSchema string

const PatchSchemaUrnIetfParamsScimApiMessages20PatchOp PatchSchema = `urn:ietf:params:scim:api:messages:2.0:PatchOp`

// String representation for [fmt.Print]
func (f *PatchSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PatchSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:api:messages:2.0:PatchOp`:
		*f = PatchSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:api:messages:2.0:PatchOp"`, v)
	}
}

// Type always returns PatchSchema to satisfy [pflag.Value] interface
func (f *PatchSchema) Type() string {
	return "PatchSchema"
}

type Permission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Permission) UnmarshalJSON(b []byte) error {
	type C Permission
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Permission) MarshalJSON() ([]byte, error) {
	type C Permission
	return marshal.Marshal((C)(s))
}

type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `json:"error,omitempty"`
	// The permissions level of the principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// Information about the principal assigned to the workspace.
	Principal *PrincipalOutput `json:"principal,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PermissionAssignment) UnmarshalJSON(b []byte) error {
	type C PermissionAssignment
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PermissionAssignment) MarshalJSON() ([]byte, error) {
	type C PermissionAssignment
	return marshal.Marshal((C)(s))
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PermissionAssignments) UnmarshalJSON(b []byte) error {
	type C PermissionAssignments
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PermissionAssignments) MarshalJSON() ([]byte, error) {
	type C PermissionAssignments
	return marshal.Marshal((C)(s))
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
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_BIND`, `CAN_EDIT`, `CAN_EDIT_METADATA`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_RUN`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`, `CAN_RESTART`, `CAN_RUN`, `CAN_USE`, `CAN_VIEW`, `CAN_VIEW_METADATA`, `IS_OWNER`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_BIND", "CAN_EDIT", "CAN_EDIT_METADATA", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_RUN", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ", "CAN_RESTART", "CAN_RUN", "CAN_USE", "CAN_VIEW", "CAN_VIEW_METADATA", "IS_OWNER"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

type PermissionOutput struct {
	// The results of a permissions query.
	Description string `json:"description,omitempty"`

	PermissionLevel WorkspacePermission `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PermissionOutput) UnmarshalJSON(b []byte) error {
	type C PermissionOutput
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PermissionOutput) MarshalJSON() ([]byte, error) {
	type C PermissionOutput
	return marshal.Marshal((C)(s))
}

type PermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PermissionsDescription) UnmarshalJSON(b []byte) error {
	type C PermissionsDescription
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PermissionsDescription) MarshalJSON() ([]byte, error) {
	type C PermissionsDescription
	return marshal.Marshal((C)(s))
}

type PermissionsRequest struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`

	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *PermissionsRequest) UnmarshalJSON(b []byte) error {
	type C PermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PermissionsRequest) MarshalJSON() ([]byte, error) {
	type C PermissionsRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *PrincipalOutput) UnmarshalJSON(b []byte) error {
	type C PrincipalOutput
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PrincipalOutput) MarshalJSON() ([]byte, error) {
	type C PrincipalOutput
	return marshal.Marshal((C)(s))
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType string `json:"resourceType,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ResourceMeta) UnmarshalJSON(b []byte) error {
	type C ResourceMeta
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ResourceMeta) MarshalJSON() ([]byte, error) {
	type C ResourceMeta
	return marshal.Marshal((C)(s))
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name string `json:"name"`

	ForceSendFields []string `json:"-"`
}

func (s *Role) UnmarshalJSON(b []byte) error {
	type C Role
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Role) MarshalJSON() ([]byte, error) {
	type C Role
	return marshal.Marshal((C)(s))
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned.
	Etag string `json:"etag,omitempty"`

	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RuleSetResponse) UnmarshalJSON(b []byte) error {
	type C RuleSetResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RuleSetResponse) MarshalJSON() ([]byte, error) {
	type C RuleSetResponse
	return marshal.Marshal((C)(s))
}

type RuleSetUpdateRequest struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	Etag string `json:"etag"`

	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name"`

	ForceSendFields []string `json:"-"`
}

func (s *RuleSetUpdateRequest) UnmarshalJSON(b []byte) error {
	type C RuleSetUpdateRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RuleSetUpdateRequest) MarshalJSON() ([]byte, error) {
	type C RuleSetUpdateRequest
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *ServicePrincipal) UnmarshalJSON(b []byte) error {
	type C ServicePrincipal
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ServicePrincipal) MarshalJSON() ([]byte, error) {
	type C ServicePrincipal
	return marshal.Marshal((C)(s))
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name string `json:"name"`

	RuleSet RuleSetUpdateRequest `json:"rule_set"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateRuleSetRequest) UnmarshalJSON(b []byte) error {
	type C UpdateRuleSetRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdateRuleSetRequest) MarshalJSON() ([]byte, error) {
	type C UpdateRuleSetRequest
	return marshal.Marshal((C)(s))
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace.
	Permissions []WorkspacePermission `json:"permissions"`
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateWorkspaceAssignments) UnmarshalJSON(b []byte) error {
	type C UpdateWorkspaceAssignments
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdateWorkspaceAssignments) MarshalJSON() ([]byte, error) {
	type C UpdateWorkspaceAssignments
	return marshal.Marshal((C)(s))
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

	ForceSendFields []string `json:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	type C User
	return marshal.Unmarshal(b, (*C)(s))
}

func (s User) MarshalJSON() ([]byte, error) {
	type C User
	return marshal.Marshal((C)(s))
}

type WorkspacePermission string

const WorkspacePermissionAdmin WorkspacePermission = `ADMIN`

const WorkspacePermissionUnknown WorkspacePermission = `UNKNOWN`

const WorkspacePermissionUser WorkspacePermission = `USER`

// String representation for [fmt.Print]
func (f *WorkspacePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspacePermission) Set(v string) error {
	switch v {
	case `ADMIN`, `UNKNOWN`, `USER`:
		*f = WorkspacePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADMIN", "UNKNOWN", "USER"`, v)
	}
}

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (f *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions []PermissionOutput `json:"permissions,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspacePermissions) UnmarshalJSON(b []byte) error {
	type C WorkspacePermissions
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspacePermissions) MarshalJSON() ([]byte, error) {
	type C WorkspacePermissions
	return marshal.Marshal((C)(s))
}
