// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type AccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *AccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := accessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []Permission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *AccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := accessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// represents an identity trying to access a resource - user or a service
// principal group can be a principal of a permission set assignment but an
// actor is always a user or a service principal
type Actor struct {

	// Wire name: 'actor_id'
	ActorId int64

	ForceSendFields []string `tf:"-"`
}

func (st *Actor) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &actorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := actorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Actor) MarshalJSON() ([]byte, error) {
	pb, err := actorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Check access policy to a resource
type CheckPolicyRequest struct {

	// Wire name: 'actor'
	Actor Actor `tf:"-"`

	// Wire name: 'authz_identity'
	AuthzIdentity RequestAuthzIdentity `tf:"-"`

	// Wire name: 'consistency_token'
	ConsistencyToken ConsistencyToken `tf:"-"`

	// Wire name: 'permission'
	Permission string `tf:"-"`
	// Ex: (servicePrincipal/use,
	// accounts/<account-id>/servicePrincipals/<sp-id>) Ex:
	// (servicePrincipal.ruleSet/update,
	// accounts/<account-id>/servicePrincipals/<sp-id>/ruleSets/default)
	// Wire name: 'resource'
	Resource string `tf:"-"`

	// Wire name: 'resource_info'
	ResourceInfo *ResourceInfo `tf:"-"`
}

func (st *CheckPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &checkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := checkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CheckPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := checkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CheckPolicyResponse struct {

	// Wire name: 'consistency_token'
	ConsistencyToken ConsistencyToken

	// Wire name: 'is_permitted'
	IsPermitted bool

	ForceSendFields []string `tf:"-"`
}

func (st *CheckPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &checkPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := checkPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CheckPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := checkPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ComplexValue struct {

	// Wire name: 'display'
	Display string

	// Wire name: 'primary'
	Primary bool

	// Wire name: '$ref'
	Ref string

	// Wire name: 'type'
	Type string

	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *ComplexValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &complexValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := complexValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComplexValue) MarshalJSON() ([]byte, error) {
	pb, err := complexValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ConsistencyToken struct {

	// Wire name: 'value'
	Value string
}

func (st *ConsistencyToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &consistencyTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := consistencyTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ConsistencyToken) MarshalJSON() ([]byte, error) {
	pb, err := consistencyTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a group.
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteAccountGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a service principal.
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteAccountServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a user.
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteAccountUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a group.
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteResponse struct {
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a service principal.
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a user.
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete permissions assignment
type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	// Wire name: 'principal_id'
	PrincipalId int64 `tf:"-"`
	// The workspace ID for the account.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *DeleteWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWorkspaceAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWorkspaceAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteWorkspaceAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteWorkspacePermissionAssignmentResponse struct {
}

func (st *DeleteWorkspacePermissionAssignmentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWorkspacePermissionAssignmentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWorkspacePermissionAssignmentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWorkspacePermissionAssignmentResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteWorkspacePermissionAssignmentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get group details.
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *GetAccountGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get service principal details.
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *GetAccountServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get user details.
type GetAccountUserRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page. Default is 10000.
	// Wire name: 'count'
	Count int `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Unique ID for a user in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder GetSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetAccountUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	//
	// Examples | Summary :--- | :--- `resource=accounts/<ACCOUNT_ID>` | A
	// resource name for the account.
	// `resource=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>` | A resource name for
	// the group. `resource=accounts/<ACCOUNT_ID>/servicePrincipals/<SP_ID>` | A
	// resource name for the service principal.
	// Wire name: 'resource'
	Resource string `tf:"-"`
}

func (st *GetAssignableRolesForResourceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAssignableRolesForResourceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAssignableRolesForResourceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAssignableRolesForResourceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAssignableRolesForResourceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetAssignableRolesForResourceResponse struct {

	// Wire name: 'roles'
	Roles []Role
}

func (st *GetAssignableRolesForResourceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAssignableRolesForResourceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAssignableRolesForResourceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAssignableRolesForResourceResponse) MarshalJSON() ([]byte, error) {
	pb, err := getAssignableRolesForResourceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get group details.
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *GetGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := getGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []PasswordPermissionsDescription
}

func (st *GetPasswordPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPasswordPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPasswordPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPasswordPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPasswordPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get object permission levels
type GetPermissionLevelsRequest struct {

	// Wire name: 'request_object_id'
	RequestObjectId string `tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	// Wire name: 'request_object_type'
	RequestObjectType string `tf:"-"`
}

func (st *GetPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []PermissionsDescription
}

func (st *GetPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get object permissions
type GetPermissionRequest struct {
	// The id of the request object.
	// Wire name: 'request_object_id'
	RequestObjectId string `tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	// Wire name: 'request_object_type'
	RequestObjectType string `tf:"-"`
}

func (st *GetPermissionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPermissionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPermissionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPermissionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPermissionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	//
	// Examples | Summary :--- | :--- `etag=` | An empty etag can only be used
	// in GET to indicate no freshness requirements.
	// `etag=RENUAAABhSweA4NvVmmUYdiU717H3Tgy0UJdor3gE4a+mq/oj9NjAf8ZsQ==` | An
	// etag encoded a specific version of the rule set to get or to be updated.
	// Wire name: 'etag'
	Etag string `tf:"-"`
	// The ruleset name associated with the request.
	//
	// Examples | Summary :--- | :---
	// `name=accounts/<ACCOUNT_ID>/ruleSets/default` | A name for a rule set on
	// the account.
	// `name=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>/ruleSets/default` | A name
	// for a rule set on the group.
	// `name=accounts/<ACCOUNT_ID>/servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>/ruleSets/default`
	// | A name for a rule set on the service principal.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st *GetRuleSetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRuleSetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRuleSetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRuleSetRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRuleSetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get service principal details.
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *GetServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := getServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns GetSortOrder to satisfy [pflag.Value] interface
func (f *GetSortOrder) Type() string {
	return "GetSortOrder"
}

// Get user details.
type GetUserRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page.
	// Wire name: 'count'
	Count int `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Unique ID for a user in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder GetSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := getUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *GetWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GrantRule struct {
	// Principals this grant rule applies to. A principal can be a user (for end
	// users), a service principal (for applications and compute workloads), or
	// an account group. Each principal has its own identifier format: *
	// users/<USERNAME> * groups/<GROUP_NAME> *
	// servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>
	// Wire name: 'principals'
	Principals []string
	// Role that is assigned to the list of principals.
	// Wire name: 'role'
	Role string
}

func (st *GrantRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &grantRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := grantRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GrantRule) MarshalJSON() ([]byte, error) {
	pb, err := grantRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Group struct {
	// String that represents a human-readable group name
	// Wire name: 'displayName'
	DisplayName string
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	// Wire name: 'entitlements'
	Entitlements []ComplexValue

	// Wire name: 'externalId'
	ExternalId string

	// Wire name: 'groups'
	Groups []ComplexValue
	// Databricks group ID
	// Wire name: 'id'
	Id string

	// Wire name: 'members'
	Members []ComplexValue
	// Container for the group identifier. Workspace local versus account.
	// Wire name: 'meta'
	Meta *ResourceMeta
	// Corresponds to AWS instance profile/arn role.
	// Wire name: 'roles'
	Roles []ComplexValue
	// The schema of the group.
	// Wire name: 'schemas'
	Schemas []GroupSchema

	ForceSendFields []string `tf:"-"`
}

func (st *Group) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &groupPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := groupFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Group) MarshalJSON() ([]byte, error) {
	pb, err := groupToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns GroupSchema to satisfy [pflag.Value] interface
func (f *GroupSchema) Type() string {
	return "GroupSchema"
}

// List group details.
type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page. Default is 10000.
	// Wire name: 'count'
	Count int64 `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Attribute to sort the results.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder ListSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListAccountGroupsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountGroupsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountGroupsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAccountGroupsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List service principals.
type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page. Default is 10000.
	// Wire name: 'count'
	Count int64 `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Attribute to sort the results.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder ListSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListAccountServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountServicePrincipalsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountServicePrincipalsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAccountServicePrincipalsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List users.
type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page. Default is 10000.
	// Wire name: 'count'
	Count int64 `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder ListSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListAccountUsersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountUsersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountUsersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountUsersRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAccountUsersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List group details.
type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page.
	// Wire name: 'count'
	Count int64 `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Attribute to sort the results.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder ListSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listGroupsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listGroupsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListGroupsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listGroupsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	// Wire name: 'itemsPerPage'
	ItemsPerPage int64
	// User objects returned in the response.
	// Wire name: 'Resources'
	Resources []Group
	// The schema of the service principal.
	// Wire name: 'schemas'
	Schemas []ListResponseSchema
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64
	// Total results that match the request filters.
	// Wire name: 'totalResults'
	TotalResults int64

	ForceSendFields []string `tf:"-"`
}

func (st *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listGroupsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listGroupsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListGroupsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listGroupsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns ListResponseSchema to satisfy [pflag.Value] interface
func (f *ListResponseSchema) Type() string {
	return "ListResponseSchema"
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	// Wire name: 'itemsPerPage'
	ItemsPerPage int64
	// User objects returned in the response.
	// Wire name: 'Resources'
	Resources []ServicePrincipal
	// The schema of the List response.
	// Wire name: 'schemas'
	Schemas []ListResponseSchema
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64
	// Total results that match the request filters.
	// Wire name: 'totalResults'
	TotalResults int64

	ForceSendFields []string `tf:"-"`
}

func (st *ListServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listServicePrincipalResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listServicePrincipalResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListServicePrincipalResponse) MarshalJSON() ([]byte, error) {
	pb, err := listServicePrincipalResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List service principals.
type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page.
	// Wire name: 'count'
	Count int64 `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Attribute to sort the results.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder ListSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listServicePrincipalsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listServicePrincipalsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listServicePrincipalsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// List users.
type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	// Wire name: 'attributes'
	Attributes string `tf:"-"`
	// Desired number of results per page.
	// Wire name: 'count'
	Count int64 `tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	// Wire name: 'excludedAttributes'
	ExcludedAttributes string `tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	// Wire name: 'sortBy'
	SortBy string `tf:"-"`
	// The order to sort the results.
	// Wire name: 'sortOrder'
	SortOrder ListSortOrder `tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListUsersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listUsersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listUsersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListUsersRequest) MarshalJSON() ([]byte, error) {
	pb, err := listUsersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListUsersResponse struct {
	// Total results returned in the response.
	// Wire name: 'itemsPerPage'
	ItemsPerPage int64
	// User objects returned in the response.
	// Wire name: 'Resources'
	Resources []User
	// The schema of the List response.
	// Wire name: 'schemas'
	Schemas []ListResponseSchema
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64
	// Total results that match the request filters.
	// Wire name: 'totalResults'
	TotalResults int64

	ForceSendFields []string `tf:"-"`
}

func (st *ListUsersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listUsersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listUsersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListUsersResponse) MarshalJSON() ([]byte, error) {
	pb, err := listUsersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *ListWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listWorkspaceAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listWorkspaceAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := listWorkspaceAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MigratePermissionsRequest struct {
	// The name of the workspace group that permissions will be migrated from.
	// Wire name: 'from_workspace_group_name'
	FromWorkspaceGroupName string
	// The maximum number of permissions that will be migrated.
	// Wire name: 'size'
	Size int
	// The name of the account group that permissions will be migrated to.
	// Wire name: 'to_account_group_name'
	ToAccountGroupName string
	// WorkspaceId of the associated workspace where the permission migration
	// will occur.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func (st *MigratePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &migratePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := migratePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MigratePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := migratePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	// Wire name: 'permissions_migrated'
	PermissionsMigrated int

	ForceSendFields []string `tf:"-"`
}

func (st *MigratePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &migratePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := migratePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MigratePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := migratePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Name struct {
	// Family name of the Databricks user.
	// Wire name: 'familyName'
	FamilyName string
	// Given name of the Databricks user.
	// Wire name: 'givenName'
	GivenName string

	ForceSendFields []string `tf:"-"`
}

func (st *Name) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &namePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nameFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Name) MarshalJSON() ([]byte, error) {
	pb, err := nameToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func (st *ObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &objectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := objectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := objectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PartialUpdate struct {
	// Unique ID in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'Operations'
	Operations []Patch
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	// Wire name: 'schemas'
	Schemas []PatchSchema
}

func (st *PartialUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &partialUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := partialUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PartialUpdate) MarshalJSON() ([]byte, error) {
	pb, err := partialUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PasswordAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *PasswordAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &passwordAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := passwordAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PasswordAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := passwordAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PasswordAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []PasswordPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *PasswordAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &passwordAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := passwordAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PasswordAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := passwordAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PasswordPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *PasswordPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &passwordPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := passwordPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PasswordPermission) MarshalJSON() ([]byte, error) {
	pb, err := passwordPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []PasswordAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func (st *PasswordPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &passwordPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := passwordPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PasswordPermissions) MarshalJSON() ([]byte, error) {
	pb, err := passwordPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PasswordPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *PasswordPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &passwordPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := passwordPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PasswordPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := passwordPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PasswordPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []PasswordAccessControlRequest
}

func (st *PasswordPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &passwordPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := passwordPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PasswordPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := passwordPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Patch struct {
	// Type of patch operation.
	// Wire name: 'op'
	Op PatchOp
	// Selection of patch operation
	// Wire name: 'path'
	Path string
	// Value to modify
	// Wire name: 'value'
	Value any

	ForceSendFields []string `tf:"-"`
}

func (st *Patch) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &patchPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := patchFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Patch) MarshalJSON() ([]byte, error) {
	pb, err := patchToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type PatchResponse struct {
}

func (st *PatchResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &patchResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := patchResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PatchResponse) MarshalJSON() ([]byte, error) {
	pb, err := patchResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *Permission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Permission) MarshalJSON() ([]byte, error) {
	pb, err := permissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The output format for existing workspace PermissionAssignment records, which
// contains some info for user consumption.
type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	// Wire name: 'error'
	Error string
	// The permissions level of the principal.
	// Wire name: 'permissions'
	Permissions []WorkspacePermission
	// Information about the principal assigned to the workspace.
	// Wire name: 'principal'
	Principal *PrincipalOutput

	ForceSendFields []string `tf:"-"`
}

func (st *PermissionAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionAssignment) MarshalJSON() ([]byte, error) {
	pb, err := permissionAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	// Wire name: 'permission_assignments'
	PermissionAssignments []PermissionAssignment
}

func (st *PermissionAssignments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionAssignmentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionAssignmentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionAssignments) MarshalJSON() ([]byte, error) {
	pb, err := permissionAssignmentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

type PermissionOutput struct {
	// The results of a permissions query.
	// Wire name: 'description'
	Description string

	// Wire name: 'permission_level'
	PermissionLevel WorkspacePermission

	ForceSendFields []string `tf:"-"`
}

func (st *PermissionOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionOutput) MarshalJSON() ([]byte, error) {
	pb, err := permissionOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *PermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := permissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Information about the principal assigned to the workspace.
type PrincipalOutput struct {
	// The display name of the principal.
	// Wire name: 'display_name'
	DisplayName string
	// The group name of the group. Present only if the principal is a group.
	// Wire name: 'group_name'
	GroupName string
	// The unique, opaque id of the principal.
	// Wire name: 'principal_id'
	PrincipalId int64
	// The name of the service principal. Present only if the principal is a
	// service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// The username of the user. Present only if the principal is a user.
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *PrincipalOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &principalOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := principalOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PrincipalOutput) MarshalJSON() ([]byte, error) {
	pb, err := principalOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns RequestAuthzIdentity to satisfy [pflag.Value] interface
func (f *RequestAuthzIdentity) Type() string {
	return "RequestAuthzIdentity"
}

type ResourceInfo struct {
	// Id of the current resource.
	// Wire name: 'id'
	Id string
	// The legacy acl path of the current resource.
	// Wire name: 'legacy_acl_path'
	LegacyAclPath string
	// Parent resource info for the current resource. The parent may have
	// another parent.
	// Wire name: 'parent_resource_info'
	ParentResourceInfo *ResourceInfo

	ForceSendFields []string `tf:"-"`
}

func (st *ResourceInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resourceInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resourceInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResourceInfo) MarshalJSON() ([]byte, error) {
	pb, err := resourceInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	// Wire name: 'resourceType'
	ResourceType string

	ForceSendFields []string `tf:"-"`
}

func (st *ResourceMeta) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resourceMetaPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resourceMetaFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResourceMeta) MarshalJSON() ([]byte, error) {
	pb, err := resourceMetaToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	// Wire name: 'name'
	Name string
}

func (st *Role) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := roleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Role) MarshalJSON() ([]byte, error) {
	pb, err := roleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string

	// Wire name: 'grant_rules'
	GrantRules []GrantRule
	// Name of the rule set.
	// Wire name: 'name'
	Name string
}

func (st *RuleSetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ruleSetResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ruleSetResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RuleSetResponse) MarshalJSON() ([]byte, error) {
	pb, err := ruleSetResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string

	// Wire name: 'grant_rules'
	GrantRules []GrantRule
	// Name of the rule set.
	// Wire name: 'name'
	Name string
}

func (st *RuleSetUpdateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ruleSetUpdateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ruleSetUpdateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RuleSetUpdateRequest) MarshalJSON() ([]byte, error) {
	pb, err := ruleSetUpdateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServicePrincipal struct {
	// If this user is active
	// Wire name: 'active'
	Active bool
	// UUID relating to the service principal
	// Wire name: 'applicationId'
	ApplicationId string
	// String that represents a concatenation of given and family names.
	// Wire name: 'displayName'
	DisplayName string
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	// Wire name: 'entitlements'
	Entitlements []ComplexValue

	// Wire name: 'externalId'
	ExternalId string

	// Wire name: 'groups'
	Groups []ComplexValue
	// Databricks service principal ID.
	// Wire name: 'id'
	Id string
	// Corresponds to AWS instance profile/arn role.
	// Wire name: 'roles'
	Roles []ComplexValue
	// The schema of the List response.
	// Wire name: 'schemas'
	Schemas []ServicePrincipalSchema

	ForceSendFields []string `tf:"-"`
}

func (st *ServicePrincipal) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servicePrincipalPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servicePrincipalFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServicePrincipal) MarshalJSON() ([]byte, error) {
	pb, err := servicePrincipalToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns ServicePrincipalSchema to satisfy [pflag.Value] interface
func (f *ServicePrincipalSchema) Type() string {
	return "ServicePrincipalSchema"
}

type SetObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControlRequest
	// The id of the request object.
	// Wire name: 'request_object_id'
	RequestObjectId string `tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	// Wire name: 'request_object_type'
	RequestObjectType string `tf:"-"`
}

func (st *SetObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := setObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControlRequest
	// The id of the request object.
	// Wire name: 'request_object_id'
	RequestObjectId string `tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	// Wire name: 'request_object_type'
	RequestObjectType string `tf:"-"`
}

func (st *UpdateObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := updateObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateResponse struct {
}

func (st *UpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	// Wire name: 'name'
	Name string

	// Wire name: 'rule_set'
	RuleSet RuleSetUpdateRequest
}

func (st *UpdateRuleSetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRuleSetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRuleSetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRuleSetRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateRuleSetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace. Valid values
	// are "USER" and "ADMIN" (case-sensitive). If both "USER" and "ADMIN" are
	// provided, "ADMIN" takes precedence. Other values will be ignored. Note
	// that excluding this field, or providing unsupported values, will have the
	// same effect as providing an empty list, which will result in the deletion
	// of all permissions for the principal.
	// Wire name: 'permissions'
	Permissions []WorkspacePermission
	// The ID of the user, service principal, or group.
	// Wire name: 'principal_id'
	PrincipalId int64 `tf:"-"`
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *UpdateWorkspaceAssignments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceAssignmentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceAssignmentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceAssignments) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceAssignmentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type User struct {
	// If this user is active
	// Wire name: 'active'
	Active bool
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	// Wire name: 'displayName'
	DisplayName string
	// All the emails associated with the Databricks user.
	// Wire name: 'emails'
	Emails []ComplexValue
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	// Wire name: 'entitlements'
	Entitlements []ComplexValue
	// External ID is not currently supported. It is reserved for future use.
	// Wire name: 'externalId'
	ExternalId string

	// Wire name: 'groups'
	Groups []ComplexValue
	// Databricks user ID.
	// Wire name: 'id'
	Id string

	// Wire name: 'name'
	Name *Name
	// Corresponds to AWS instance profile/arn role.
	// Wire name: 'roles'
	Roles []ComplexValue
	// The schema of the user.
	// Wire name: 'schemas'
	Schemas []UserSchema
	// Email address of the Databricks user.
	// Wire name: 'userName'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *User) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &userPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := userFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st User) MarshalJSON() ([]byte, error) {
	pb, err := userToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (f *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	// Wire name: 'permissions'
	Permissions []PermissionOutput
}

func (st *WorkspacePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspacePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspacePermissions) MarshalJSON() ([]byte, error) {
	pb, err := workspacePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
