// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// represents an identity trying to access a resource - user or a service
// principal group can be a principal of a permission set assignment but an
// actor is always a user or a service principal
type Actor struct {
	ActorId int64 `json:"actor_id,omitempty" url:"actor_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Actor) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Actor) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CheckPolicyRequest struct {
	Actor Actor `json:"-" url:"actor"`

	AuthzIdentity RequestAuthzIdentity `json:"-" url:"authz_identity"`

	ConsistencyToken ConsistencyToken `json:"-" url:"consistency_token"`

	Permission string `json:"-" url:"permission"`
	// Ex: (servicePrincipal/use,
	// accounts/<account-id>/servicePrincipals/<sp-id>) Ex:
	// (servicePrincipal.ruleSet/update,
	// accounts/<account-id>/servicePrincipals/<sp-id>/ruleSets/default)
	Resource string `json:"-" url:"resource"`

	ResourceInfo *ResourceInfo `json:"-" url:"resource_info,omitempty"`
}

type CheckPolicyResponse struct {
	ConsistencyToken ConsistencyToken `json:"consistency_token"`

	IsPermitted bool `json:"is_permitted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CheckPolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CheckPolicyResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ComplexValue struct {
	Display string `json:"display,omitempty"`

	Primary bool `json:"primary,omitempty"`

	Ref string `json:"$ref,omitempty"`

	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplexValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplexValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ConsistencyToken struct {
	Value string `json:"value" url:"value"`
}

type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`
}

type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`
}

type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`
}

type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`
}

type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`
}

type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`
}

type GetAccountUserRequest struct {
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
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder GetSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAccountUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAccountUserRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	//
	// Examples | Summary :--- | :--- `resource=accounts/<ACCOUNT_ID>` | A
	// resource name for the account.
	// `resource=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>` | A resource name for
	// the group. `resource=accounts/<ACCOUNT_ID>/servicePrincipals/<SP_ID>` | A
	// resource name for the service principal.
	Resource string `json:"-" url:"resource"`
}

type GetAssignableRolesForResourceResponse struct {
	Roles []Role `json:"roles,omitempty"`
}

type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PasswordPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetPermissionLevelsRequest struct {
	RequestObjectId string `json:"-" url:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" url:"-"`
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

type GetPermissionRequest struct {
	// The id of the request object.
	RequestObjectId string `json:"-" url:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" url:"-"`
}

type GetRuleSetRequest struct {
	// Etag used for versioning. The response is at least as fresh as the eTag
	// provided. Etag is used for optimistic concurrency control as a way to
	// help prevent simultaneous updates of a rule set from overwriting each
	// other. It is strongly suggested that systems make use of the etag in the
	// read -> modify -> write pattern to perform rule set updates in order to
	// avoid race conditions that is get an etag from a GET rule set request,
	// and pass it with the PUT update request to identify the rule set version
	// you are updating.
	//
	// Examples | Summary :--- | :--- `etag=` | An empty etag can only be used
	// in GET to indicate no freshness requirements.
	// `etag=RENUAAABhSweA4NvVmmUYdiU717H3Tgy0UJdor3gE4a+mq/oj9NjAf8ZsQ==` | An
	// etag encoded a specific version of the rule set to get or to be updated.
	Etag string `json:"-" url:"etag"`
	// The ruleset name associated with the request.
	//
	// Examples | Summary :--- | :---
	// `name=accounts/<ACCOUNT_ID>/ruleSets/default` | A name for a rule set on
	// the account.
	// `name=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>/ruleSets/default` | A name
	// for a rule set on the group.
	// `name=accounts/<ACCOUNT_ID>/servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>/ruleSets/default`
	// | A name for a rule set on the service principal.
	Name string `json:"-" url:"name"`
}

type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

type GetSortOrder string

const GetSortOrderAscending GetSortOrder = `ascending`

const GetSortOrderDescending GetSortOrder = `descending`

// String representation for [fmt.Print]
func (f *GetSortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetSortOrder) Set(v string) error {
	switch v {
	case `ascending`, `descending`:
		*f = GetSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ascending", "descending"`, v)
	}
}

// Values returns all possible values for GetSortOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetSortOrder) Values() []GetSortOrder {
	return []GetSortOrder{
		GetSortOrderAscending,
		GetSortOrderDescending,
	}
}

// Type always returns GetSortOrder to satisfy [pflag.Value] interface
func (f *GetSortOrder) Type() string {
	return "GetSortOrder"
}

type GetUserRequest struct {
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
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" url:"sortBy,omitempty"`
	// The order to sort the results.
	SortOrder GetSortOrder `json:"-" url:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetUserRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type GrantRule struct {
	// Principals this grant rule applies to. A principal can be a user (for end
	// users), a service principal (for applications and compute workloads), or
	// an account group. Each principal has its own identifier format: *
	// users/<USERNAME> * groups/<GROUP_NAME> *
	// servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>
	Principals []string `json:"principals,omitempty"`
	// Role that is assigned to the list of principals.
	Role string `json:"role"`
}

type Group struct {
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks group ID
	Id string `json:"id,omitempty"`

	Members []ComplexValue `json:"members,omitempty"`
	// Container for the group identifier. Workspace local versus account.
	Meta *ResourceMeta `json:"meta,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	Roles []ComplexValue `json:"roles,omitempty"`
	// The schema of the group.
	Schemas []GroupSchema `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Group) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Group) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GroupSchema string

const GroupSchemaUrnIetfParamsScimSchemasCore20Group GroupSchema = `urn:ietf:params:scim:schemas:core:2.0:Group`

// String representation for [fmt.Print]
func (f *GroupSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GroupSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:schemas:core:2.0:Group`:
		*f = GroupSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:schemas:core:2.0:Group"`, v)
	}
}

// Values returns all possible values for GroupSchema.
//
// There is no guarantee on the order of the values in the slice.
func (f *GroupSchema) Values() []GroupSchema {
	return []GroupSchema{
		GroupSchemaUrnIetfParamsScimSchemasCore20Group,
	}
}

// Type always returns GroupSchema to satisfy [pflag.Value] interface
func (f *GroupSchema) Type() string {
	return "GroupSchema"
}

type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int64 `json:"-" url:"count,omitempty"`
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
	StartIndex int64 `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int64 `json:"-" url:"count,omitempty"`
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
	StartIndex int64 `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page. Default is 10000.
	Count int64 `json:"-" url:"count,omitempty"`
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
	StartIndex int64 `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountUsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountUsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int64 `json:"-" url:"count,omitempty"`
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
	StartIndex int64 `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []Group `json:"Resources,omitempty"`
	// The schema of the service principal.
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListResponseSchema string

const ListResponseSchemaUrnIetfParamsScimApiMessages20ListResponse ListResponseSchema = `urn:ietf:params:scim:api:messages:2.0:ListResponse`

// String representation for [fmt.Print]
func (f *ListResponseSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListResponseSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:api:messages:2.0:ListResponse`:
		*f = ListResponseSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:api:messages:2.0:ListResponse"`, v)
	}
}

// Values returns all possible values for ListResponseSchema.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListResponseSchema) Values() []ListResponseSchema {
	return []ListResponseSchema{
		ListResponseSchemaUrnIetfParamsScimApiMessages20ListResponse,
	}
}

// Type always returns ListResponseSchema to satisfy [pflag.Value] interface
func (f *ListResponseSchema) Type() string {
	return "ListResponseSchema"
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []ServicePrincipal `json:"Resources,omitempty"`
	// The schema of the List response.
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int64 `json:"-" url:"count,omitempty"`
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
	StartIndex int64 `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for ListSortOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListSortOrder) Values() []ListSortOrder {
	return []ListSortOrder{
		ListSortOrderAscending,
		ListSortOrderDescending,
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (f *ListSortOrder) Type() string {
	return "ListSortOrder"
}

type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" url:"attributes,omitempty"`
	// Desired number of results per page.
	Count int64 `json:"-" url:"count,omitempty"`
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
	StartIndex int64 `json:"-" url:"startIndex,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersResponse struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []User `json:"Resources,omitempty"`
	// The schema of the List response.
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`
}

type MigratePermissionsRequest struct {
	// The name of the workspace group that permissions will be migrated from.
	FromWorkspaceGroupName string `json:"from_workspace_group_name"`
	// The maximum number of permissions that will be migrated.
	Size int `json:"size,omitempty"`
	// The name of the account group that permissions will be migrated to.
	ToAccountGroupName string `json:"to_account_group_name"`
	// WorkspaceId of the associated workspace where the permission migration
	// will occur.
	WorkspaceId int64 `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigratePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigratePermissionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	PermissionsMigrated int `json:"permissions_migrated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MigratePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MigratePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Name struct {
	// Family name of the Databricks user.
	FamilyName string `json:"familyName,omitempty"`
	// Given name of the Databricks user.
	GivenName string `json:"givenName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Name) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Name) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ObjectPermissions struct {
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ObjectPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ObjectPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PartialUpdate struct {
	// Unique ID in the Databricks workspace.
	Id string `json:"-" url:"-"`

	Operations []Patch `json:"Operations,omitempty"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas []PatchSchema `json:"schemas,omitempty"`
}

type PasswordAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PasswordAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PasswordAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PasswordPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PasswordPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for PasswordPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PasswordPermissionLevel) Values() []PasswordPermissionLevel {
	return []PasswordPermissionLevel{
		PasswordPermissionLevelCanUse,
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PasswordPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PasswordPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PasswordPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PasswordPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Value any `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Patch) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Patch) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for PatchOp.
//
// There is no guarantee on the order of the values in the slice.
func (f *PatchOp) Values() []PatchOp {
	return []PatchOp{
		PatchOpAdd,
		PatchOpRemove,
		PatchOpReplace,
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

// Values returns all possible values for PatchSchema.
//
// There is no guarantee on the order of the values in the slice.
func (f *PatchSchema) Values() []PatchSchema {
	return []PatchSchema{
		PatchSchemaUrnIetfParamsScimApiMessages20PatchOp,
	}
}

// Type always returns PatchSchema to satisfy [pflag.Value] interface
func (f *PatchSchema) Type() string {
	return "PatchSchema"
}

type Permission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Permission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Permission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The output format for existing workspace PermissionAssignment records, which
// contains some info for user consumption.
type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `json:"error,omitempty"`
	// The permissions level of the principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// Information about the principal assigned to the workspace.
	Principal *PrincipalOutput `json:"principal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PermissionAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`
}

// Permission level
type PermissionLevel string

const PermissionLevelCanAttachTo PermissionLevel = `CAN_ATTACH_TO`

const PermissionLevelCanBind PermissionLevel = `CAN_BIND`

const PermissionLevelCanCreate PermissionLevel = `CAN_CREATE`

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanEditMetadata PermissionLevel = `CAN_EDIT_METADATA`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageRun PermissionLevel = `CAN_MANAGE_RUN`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanMonitor PermissionLevel = `CAN_MONITOR`

const PermissionLevelCanMonitorOnly PermissionLevel = `CAN_MONITOR_ONLY`

const PermissionLevelCanQuery PermissionLevel = `CAN_QUERY`

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
	case `CAN_ATTACH_TO`, `CAN_BIND`, `CAN_CREATE`, `CAN_EDIT`, `CAN_EDIT_METADATA`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_RUN`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_MONITOR`, `CAN_MONITOR_ONLY`, `CAN_QUERY`, `CAN_READ`, `CAN_RESTART`, `CAN_RUN`, `CAN_USE`, `CAN_VIEW`, `CAN_VIEW_METADATA`, `IS_OWNER`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_BIND", "CAN_CREATE", "CAN_EDIT", "CAN_EDIT_METADATA", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_RUN", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MONITOR", "CAN_MONITOR_ONLY", "CAN_QUERY", "CAN_READ", "CAN_RESTART", "CAN_RUN", "CAN_USE", "CAN_VIEW", "CAN_VIEW_METADATA", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelCanAttachTo,
		PermissionLevelCanBind,
		PermissionLevelCanCreate,
		PermissionLevelCanEdit,
		PermissionLevelCanEditMetadata,
		PermissionLevelCanManage,
		PermissionLevelCanManageProductionVersions,
		PermissionLevelCanManageRun,
		PermissionLevelCanManageStagingVersions,
		PermissionLevelCanMonitor,
		PermissionLevelCanMonitorOnly,
		PermissionLevelCanQuery,
		PermissionLevelCanRead,
		PermissionLevelCanRestart,
		PermissionLevelCanRun,
		PermissionLevelCanUse,
		PermissionLevelCanView,
		PermissionLevelCanViewMetadata,
		PermissionLevelIsOwner,
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PermissionOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Information about the principal assigned to the workspace.
type PrincipalOutput struct {
	// The display name of the principal.
	DisplayName string `json:"display_name,omitempty"`
	// The group name of the group. Present only if the principal is a group.
	GroupName string `json:"group_name,omitempty"`
	// The unique, opaque id of the principal.
	PrincipalId int64 `json:"principal_id,omitempty"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The username of the user. Present only if the principal is a user.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrincipalOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrincipalOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Defines the identity to be used for authZ of the request on the server side.
// See one pager for for more information: http://go/acl/service-identity
type RequestAuthzIdentity string

const RequestAuthzIdentityRequestAuthzIdentityServiceIdentity RequestAuthzIdentity = `REQUEST_AUTHZ_IDENTITY_SERVICE_IDENTITY`

const RequestAuthzIdentityRequestAuthzIdentityUserContext RequestAuthzIdentity = `REQUEST_AUTHZ_IDENTITY_USER_CONTEXT`

// String representation for [fmt.Print]
func (f *RequestAuthzIdentity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RequestAuthzIdentity) Set(v string) error {
	switch v {
	case `REQUEST_AUTHZ_IDENTITY_SERVICE_IDENTITY`, `REQUEST_AUTHZ_IDENTITY_USER_CONTEXT`:
		*f = RequestAuthzIdentity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "REQUEST_AUTHZ_IDENTITY_SERVICE_IDENTITY", "REQUEST_AUTHZ_IDENTITY_USER_CONTEXT"`, v)
	}
}

// Values returns all possible values for RequestAuthzIdentity.
//
// There is no guarantee on the order of the values in the slice.
func (f *RequestAuthzIdentity) Values() []RequestAuthzIdentity {
	return []RequestAuthzIdentity{
		RequestAuthzIdentityRequestAuthzIdentityServiceIdentity,
		RequestAuthzIdentityRequestAuthzIdentityUserContext,
	}
}

// Type always returns RequestAuthzIdentity to satisfy [pflag.Value] interface
func (f *RequestAuthzIdentity) Type() string {
	return "RequestAuthzIdentity"
}

type ResourceInfo struct {
	// Id of the current resource.
	Id string `json:"id" url:"id"`
	// The legacy acl path of the current resource.
	LegacyAclPath string `json:"legacy_acl_path,omitempty" url:"legacy_acl_path,omitempty"`
	// Parent resource info for the current resource. The parent may have
	// another parent.
	ParentResourceInfo *ResourceInfo `json:"parent_resource_info,omitempty" url:"parent_resource_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResourceInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResourceInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType string `json:"resourceType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResourceMeta) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResourceMeta) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name string `json:"name"`
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned. Etag used for
	// versioning. The response is at least as fresh as the eTag provided. Etag
	// is used for optimistic concurrency control as a way to help prevent
	// simultaneous updates of a rule set from overwriting each other. It is
	// strongly suggested that systems make use of the etag in the read ->
	// modify -> write pattern to perform rule set updates in order to avoid
	// race conditions that is get an etag from a GET rule set request, and pass
	// it with the PUT update request to identify the rule set version you are
	// updating.
	Etag string `json:"etag"`

	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name"`
}

type RuleSetUpdateRequest struct {
	// Identifies the version of the rule set returned. Etag used for
	// versioning. The response is at least as fresh as the eTag provided. Etag
	// is used for optimistic concurrency control as a way to help prevent
	// simultaneous updates of a rule set from overwriting each other. It is
	// strongly suggested that systems make use of the etag in the read ->
	// modify -> write pattern to perform rule set updates in order to avoid
	// race conditions that is get an etag from a GET rule set request, and pass
	// it with the PUT update request to identify the rule set version you are
	// updating.
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
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks service principal ID.
	Id string `json:"id,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	Roles []ComplexValue `json:"roles,omitempty"`
	// The schema of the List response.
	Schemas []ServicePrincipalSchema `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServicePrincipal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServicePrincipalSchema string

const ServicePrincipalSchemaUrnIetfParamsScimSchemasCore20ServicePrincipal ServicePrincipalSchema = `urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal`

// String representation for [fmt.Print]
func (f *ServicePrincipalSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServicePrincipalSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal`:
		*f = ServicePrincipalSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal"`, v)
	}
}

// Values returns all possible values for ServicePrincipalSchema.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServicePrincipalSchema) Values() []ServicePrincipalSchema {
	return []ServicePrincipalSchema{
		ServicePrincipalSchemaUrnIetfParamsScimSchemasCore20ServicePrincipal,
	}
}

// Type always returns ServicePrincipalSchema to satisfy [pflag.Value] interface
func (f *ServicePrincipalSchema) Type() string {
	return "ServicePrincipalSchema"
}

type SetObjectPermissions struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`
	// The id of the request object.
	RequestObjectId string `json:"-" url:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" url:"-"`
}

type UpdateObjectPermissions struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`
	// The id of the request object.
	RequestObjectId string `json:"-" url:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" url:"-"`
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	Name string `json:"name"`

	RuleSet RuleSetUpdateRequest `json:"rule_set"`
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace. Valid values
	// are "USER" and "ADMIN" (case-sensitive). If both "USER" and "ADMIN" are
	// provided, "ADMIN" takes precedence. Other values will be ignored. Note
	// that excluding this field, or providing unsupported values, will have the
	// same effect as providing an empty list, which will result in the deletion
	// of all permissions for the principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

type User struct {
	// If this user is active
	Active bool `json:"active,omitempty"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	DisplayName string `json:"displayName,omitempty"`
	// All the emails associated with the Databricks user.
	Emails []ComplexValue `json:"emails,omitempty"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []ComplexValue `json:"entitlements,omitempty"`
	// External ID is not currently supported. It is reserved for future use.
	ExternalId string `json:"externalId,omitempty"`

	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks user ID.
	Id string `json:"id,omitempty"`

	Name *Name `json:"name,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	Roles []ComplexValue `json:"roles,omitempty"`
	// The schema of the user.
	Schemas []UserSchema `json:"schemas,omitempty"`
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s User) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UserSchema string

const UserSchemaUrnIetfParamsScimSchemasCore20User UserSchema = `urn:ietf:params:scim:schemas:core:2.0:User`

const UserSchemaUrnIetfParamsScimSchemasExtensionWorkspace20User UserSchema = `urn:ietf:params:scim:schemas:extension:workspace:2.0:User`

// String representation for [fmt.Print]
func (f *UserSchema) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UserSchema) Set(v string) error {
	switch v {
	case `urn:ietf:params:scim:schemas:core:2.0:User`, `urn:ietf:params:scim:schemas:extension:workspace:2.0:User`:
		*f = UserSchema(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:workspace:2.0:User"`, v)
	}
}

// Values returns all possible values for UserSchema.
//
// There is no guarantee on the order of the values in the slice.
func (f *UserSchema) Values() []UserSchema {
	return []UserSchema{
		UserSchemaUrnIetfParamsScimSchemasCore20User,
		UserSchemaUrnIetfParamsScimSchemasExtensionWorkspace20User,
	}
}

// Type always returns UserSchema to satisfy [pflag.Value] interface
func (f *UserSchema) Type() string {
	return "UserSchema"
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

// Values returns all possible values for WorkspacePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspacePermission) Values() []WorkspacePermission {
	return []WorkspacePermission{
		WorkspacePermissionAdmin,
		WorkspacePermissionUnknown,
		WorkspacePermissionUser,
	}
}

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (f *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	Permissions []PermissionOutput `json:"permissions,omitempty"`
}
