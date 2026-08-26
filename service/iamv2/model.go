// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateDirectGroupMemberProxyRequest struct {
	// Required. The group membership to create.
	DirectGroupMember DirectGroupMember `json:"direct_group_member"`
	// Required. Internal ID of the group in Databricks.
	GroupId int64 `json:"-" url:"-"`
}

func (s *CreateDirectGroupMemberProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateDirectGroupMemberRequest struct {
	// Required. The direct group member to be added to the group.
	DirectGroupMember DirectGroupMember `json:"direct_group_member"`
	// Required. Internal ID of the group in Databricks.
	GroupId int64 `json:"-" url:"-"`
}

func (s *CreateDirectGroupMemberRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateGroupProxyRequest struct {
	// Required. Group to be created in <Databricks>
	Group Group `json:"group"`
}

func (s *CreateGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateGroupRequest struct {
	// Required. Group to be created in <Databricks>
	Group Group `json:"group"`
}

func (s *CreateGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateServicePrincipalProxyRequest struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
}

func (s *CreateServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateServicePrincipalRequest struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
}

func (s *CreateServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateUserProxyRequest struct {
	// Required. User to be created in <Databricks>
	User User `json:"user"`
}

func (s *CreateUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateUserRequest struct {
	// Required. User to be created in <Databricks>
	User User `json:"user"`
}

func (s *CreateUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
}

func (s *CreateWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentDetailRequest struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being created.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *CreateWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentProxyRequest struct {
	// Required. Workspace assignment to be created in <Databricks>.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
}

func (s *CreateWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type CreateWorkspaceAssignmentRequest struct {
	// Required. Workspace assignment to be created in <Databricks>.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
	// Required. The workspace ID for which the workspace assignment is being
	// created.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *CreateWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDirectGroupMemberProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId int64 `json:"-" url:"-"`
	// Required. Internal ID of the principal to be unassigned from the group.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *DeleteDirectGroupMemberProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteDirectGroupMemberRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId int64 `json:"-" url:"-"`
	// Required. Internal ID of the principal to be unassigned from the group.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *DeleteDirectGroupMemberRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId string `json:"-" url:"-"`
}

func (s *DeleteGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId string `json:"-" url:"-"`
}

func (s *DeleteGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"-" url:"-"`
}

func (s *DeleteServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"-" url:"-"`
}

func (s *DeleteServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId string `json:"-" url:"-"`
}

func (s *DeleteUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId string `json:"-" url:"-"`
}

func (s *DeleteUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentProxyRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type DeleteWorkspaceAssignmentRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *DeleteWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Represents a principal that is a direct member of a group, with its source of
// membership.
type DirectGroupMember struct {
	// Display name of the principal.
	DisplayName string `json:"display_name,omitempty"`
	// The external ID of the principal in Databricks.
	ExternalId string `json:"external_id,omitempty"`
	// The internal ID of the group this member belongs to.
	GroupId int64 `json:"group_id,omitempty"`
	// The source of group membership (internal or from identity provider).
	MembershipSource GroupMembershipSource `json:"membership_source,omitempty"`
	// Internal ID of the principal in Databricks.
	PrincipalId int64 `json:"principal_id"`
	// The type of the principal (user/service principal/group).
	PrincipalType PrincipalType `json:"principal_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DirectGroupMember) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DirectGroupMember) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Entitlement string

const EntitlementAllowClusterCreate Entitlement = `ALLOW_CLUSTER_CREATE`

const EntitlementAllowInstancePoolCreate Entitlement = `ALLOW_INSTANCE_POOL_CREATE`

const EntitlementDatabricksSqlAccess Entitlement = `DATABRICKS_SQL_ACCESS`

const EntitlementWorkspaceAccess Entitlement = `WORKSPACE_ACCESS`

const EntitlementWorkspaceAdmin Entitlement = `WORKSPACE_ADMIN`

const EntitlementWorkspaceConsume Entitlement = `WORKSPACE_CONSUME`

// String representation for [fmt.Print]
func (f *Entitlement) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Entitlement) Set(v string) error {
	switch v {
	case `ALLOW_CLUSTER_CREATE`, `ALLOW_INSTANCE_POOL_CREATE`, `DATABRICKS_SQL_ACCESS`, `WORKSPACE_ACCESS`, `WORKSPACE_ADMIN`, `WORKSPACE_CONSUME`:
		*f = Entitlement(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_CLUSTER_CREATE", "ALLOW_INSTANCE_POOL_CREATE", "DATABRICKS_SQL_ACCESS", "WORKSPACE_ACCESS", "WORKSPACE_ADMIN", "WORKSPACE_CONSUME"`, v)
	}
}

// Values returns all possible values for Entitlement.
//
// There is no guarantee on the order of the values in the slice.
func (f *Entitlement) Values() []Entitlement {
	return []Entitlement{
		EntitlementAllowClusterCreate,
		EntitlementAllowInstancePoolCreate,
		EntitlementDatabricksSqlAccess,
		EntitlementWorkspaceAccess,
		EntitlementWorkspaceAdmin,
		EntitlementWorkspaceConsume,
	}
}

// Type always returns Entitlement to satisfy [pflag.Value] interface
func (f *Entitlement) Type() string {
	return "Entitlement"
}

type GetDirectGroupMemberProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId int64 `json:"-" url:"-"`
	// Required. Internal ID of the principal belonging to the group in
	// Databricks.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *GetDirectGroupMemberProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetDirectGroupMemberRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId int64 `json:"-" url:"-"`
	// Required. Internal ID of the principal belonging to the group in
	// Databricks.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *GetDirectGroupMemberRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId string `json:"-" url:"-"`
}

func (s *GetGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	GroupId string `json:"-" url:"-"`
}

func (s *GetGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"-" url:"-"`
}

func (s *GetServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"-" url:"-"`
}

func (s *GetServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId string `json:"-" url:"-"`
}

func (s *GetUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	UserId string `json:"-" url:"-"`
}

func (s *GetUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAccessDetailLocalRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Controls what fields are returned.
	View WorkspaceAccessDetailView `json:"-" url:"view,omitempty"`
}

func (s *GetWorkspaceAccessDetailLocalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAccessDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Controls what fields are returned.
	View WorkspaceAccessDetailView `json:"-" url:"view,omitempty"`
	// Required. The workspace ID for which the access details are being
	// requested.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAccessDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentDetailProxyRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The workspace ID for which the assignment details are being
	// requested.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentProxyRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment is being requested.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceAssignmentRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment is being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The workspace ID for which the assignment is being requested.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type GetWorkspaceIdentityDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// identity details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
}

func (s *GetWorkspaceIdentityDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The details of a Group resource.
type Group struct {
	// The parent account ID for group in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// ExternalId of the group in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`
	// Internal group ID of the group in Databricks.
	GroupId string `json:"group_id,omitempty"`
	// Display name of the group.
	GroupName string `json:"group_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Group) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Group) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The source of the group membership (internal or from identity provider).
type GroupMembershipSource string

const GroupMembershipSourceIdentityProvider GroupMembershipSource = `IDENTITY_PROVIDER`

const GroupMembershipSourceInternal GroupMembershipSource = `INTERNAL`

// String representation for [fmt.Print]
func (f *GroupMembershipSource) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GroupMembershipSource) Set(v string) error {
	switch v {
	case `IDENTITY_PROVIDER`, `INTERNAL`:
		*f = GroupMembershipSource(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IDENTITY_PROVIDER", "INTERNAL"`, v)
	}
}

// Values returns all possible values for GroupMembershipSource.
//
// There is no guarantee on the order of the values in the slice.
func (f *GroupMembershipSource) Values() []GroupMembershipSource {
	return []GroupMembershipSource{
		GroupMembershipSourceIdentityProvider,
		GroupMembershipSourceInternal,
	}
}

// Type always returns GroupMembershipSource to satisfy [pflag.Value] interface
func (f *GroupMembershipSource) Type() string {
	return "GroupMembershipSource"
}

type ListDirectGroupMembersProxyRequest struct {
	// Required. Internal ID of the group in Databricks whose direct members are
	// being listed.
	GroupId int64 `json:"-" url:"-"`
	// The maximum number of members to return. The service may return fewer
	// than this value. If not provided, defaults to 1000 (also the maximum
	// allowed).
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectGroupMembersProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectGroupMembersProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListDirectGroupMembersRequest struct {
	// Required. Internal ID of the group in Databricks whose direct members are
	// being listed.
	GroupId int64 `json:"-" url:"-"`
	// The maximum number of members to return. The service may return fewer
	// than this value. If not provided, defaults to 1000 (also the maximum
	// allowed).
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListDirectGroupMembers call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectGroupMembersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectGroupMembersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for listing direct group members.
type ListDirectGroupMembersResponse struct {
	// The list of direct group members with their membership source type.
	DirectGroupMembers []DirectGroupMember `json:"direct_group_members,omitempty"`
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListDirectGroupMembersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDirectGroupMembersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsProxyRequest struct {
	// Optional. Allows filtering groups by group name or external id.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of groups to return. The service may return fewer than
	// this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListGroups call. Provide this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsRequest struct {
	// Optional. Allows filtering groups by group name or external id.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of groups to return. The service may return fewer than
	// this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListGroups call. Provide this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message containing a page of groups in the account.
type ListGroupsResponse struct {
	Groups []Group `json:"groups,omitempty"`
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalsProxyRequest struct {
	// Optional. Allows filtering service principals by application id or
	// external id.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of SPs to return. The service may return fewer than
	// this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListServicePrincipals call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalsRequest struct {
	// Optional. Allows filtering service principals by application id or
	// external id.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of service principals to return. The service may
	// return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListServicePrincipals call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message containing a page of service principals in the account.
type ListServicePrincipalsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ServicePrincipals []ServicePrincipal `json:"service_principals,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTransitiveParentGroupsProxyRequest struct {
	// The maximum number of parent groups to return. The service may return
	// fewer than this value. If not provided, defaults to 1000 (also the
	// maximum allowed).
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListTransitiveParentGroups call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Required. Internal ID of the principal in Databricks whose transitive
	// parent groups are being listed.
	PrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTransitiveParentGroupsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTransitiveParentGroupsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTransitiveParentGroupsRequest struct {
	// The maximum number of parent groups to return. The service may return
	// fewer than this value. If not provided, defaults to 1000 (also the
	// maximum allowed).
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListTransitiveParentGroups call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Required. Internal ID of the principal in Databricks whose transitive
	// parent groups are being listed.
	PrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTransitiveParentGroupsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTransitiveParentGroupsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for listing all transitive parent groups of a principal.
type ListTransitiveParentGroupsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of transitive parent groups.
	TransitiveParentGroups []TransitiveParentGroup `json:"transitive_parent_groups,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTransitiveParentGroupsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTransitiveParentGroupsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersProxyRequest struct {
	// Optional. Allows filtering users by username or external id.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of users to return. The service may return fewer than
	// this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListUsers call. Provide this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersRequest struct {
	// Optional. Allows filtering users by username or external id.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of users to return. The service may return fewer than
	// this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListUsers call. Provide this to
	// retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUsersResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	Users []User `json:"users,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUsersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUsersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentDetailsProxyRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentDetailsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentDetailsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentDetailsRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListWorkspaceAssignmentDetails
	// call. Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Required. The workspace ID for which the workspace assignment details are
	// being fetched.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentDetailsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentDetailsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for listing workspace assignment details.
type ListWorkspaceAssignmentDetailsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	WorkspaceAssignmentDetails []WorkspaceAssignmentDetail `json:"workspace_assignment_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentDetailsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentDetailsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentsProxyRequest struct {
	// The maximum number of workspace assignments to return. The service may
	// return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token from a previous list call. Provide this to retrieve the
	// subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentsProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentsProxyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAssignmentsRequest struct {
	// The maximum number of workspace assignments to return. The service may
	// return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListWorkspaceAssignments call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Required. The workspace ID for which the workspace assignments are being
	// fetched.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response message for listing workspace assignments.
type ListWorkspaceAssignmentsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	WorkspaceAssignments []WorkspaceAssignment `json:"workspace_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of the principal (user/sp/group).
type PrincipalType string

const PrincipalTypeGroup PrincipalType = `GROUP`

const PrincipalTypeServicePrincipal PrincipalType = `SERVICE_PRINCIPAL`

const PrincipalTypeUser PrincipalType = `USER`

// String representation for [fmt.Print]
func (f *PrincipalType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrincipalType) Set(v string) error {
	switch v {
	case `GROUP`, `SERVICE_PRINCIPAL`, `USER`:
		*f = PrincipalType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GROUP", "SERVICE_PRINCIPAL", "USER"`, v)
	}
}

// Values returns all possible values for PrincipalType.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		PrincipalTypeGroup,
		PrincipalTypeServicePrincipal,
		PrincipalTypeUser,
	}
}

// Type always returns PrincipalType to satisfy [pflag.Value] interface
func (f *PrincipalType) Type() string {
	return "PrincipalType"
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// name and inherited parent groups.
type ResolveGroupProxyRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// name and inherited parent groups.
type ResolveGroupRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolveGroupResponse struct {
	// The group that was resolved.
	Group *Group `json:"group,omitempty"`
}

func (s *ResolveGroupResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's display name, status, and inherited parent groups.
type ResolveServicePrincipalProxyRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's display name, status, and inherited parent groups.
type ResolveServicePrincipalRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolveServicePrincipalResponse struct {
	// The service principal that was resolved.
	ServicePrincipal *ServicePrincipal `json:"service_principal,omitempty"`
}

func (s *ResolveServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// display name, status, and inherited parent groups.
type ResolveUserProxyRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// display name, status, and inherited parent groups.
type ResolveUserRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId string `json:"external_id"`
}

func (s *ResolveUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type ResolveUserResponse struct {
	// The user that was resolved.
	User *User `json:"user,omitempty"`
}

func (s *ResolveUserResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The details of a ServicePrincipal resource.
type ServicePrincipal struct {
	// The parent account ID for the service principal in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// The activity status of a service principal in a Databricks account.
	AccountSpStatus State `json:"account_sp_status"`
	// Application ID of the service principal. Set at creation time and cannot
	// be changed afterwards; when omitted, the server generates one.
	ApplicationId string `json:"application_id,omitempty"`
	// Display name of the service principal.
	DisplayName string `json:"display_name"`
	// ExternalId of the service principal in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`
	// Internal service principal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"service_principal_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ServicePrincipal) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServicePrincipal) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The activity status of a user or service principal in a Databricks account or
// workspace.
type State string

const StateActive State = `ACTIVE`

const StateInactive State = `INACTIVE`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `ACTIVE`, `INACTIVE`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "INACTIVE"`, v)
	}
}

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateActive,
		StateInactive,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

// Represents a group that is a transitive parent of a principal.
type TransitiveParentGroup struct {
	// The parent account ID for group in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// ExternalId of the group in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`
	// Internal group ID of the group in Databricks.
	GroupId string `json:"group_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TransitiveParentGroup) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransitiveParentGroup) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateGroupProxyRequest struct {
	// Required. Group to be updated in <Databricks>
	Group Group `json:"group"`
	// Required. Internal ID of the group in Databricks.
	GroupId string `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateGroupProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateGroupRequest struct {
	// Required. Group to be updated in <Databricks>
	Group Group `json:"group"`
	// Required. Internal ID of the group in Databricks.
	GroupId string `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateGroupRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateServicePrincipalProxyRequest struct {
	// Required. Service principal to be updated in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateServicePrincipalProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateServicePrincipalRequest struct {
	// Required. Service Principal to be updated in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
	// Required. Internal ID of the service principal in Databricks.
	ServicePrincipalId string `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

func (s *UpdateServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateUserProxyRequest struct {
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
	// Required. User to be updated in <Databricks>
	User User `json:"user"`
	// Required. Internal ID of the user in Databricks.
	UserId string `json:"-" url:"-"`
}

func (s *UpdateUserProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateUserRequest struct {
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
	// Required. User to be updated in <Databricks>
	User User `json:"user"`
	// Required. Internal ID of the user in Databricks.
	UserId string `json:"-" url:"-"`
}

func (s *UpdateUserRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment detail to be updated in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
}

func (s *UpdateWorkspaceAssignmentDetailProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment detail to be updated in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being updated.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *UpdateWorkspaceAssignmentDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentProxyRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment to be updated in <Databricks>.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
}

func (s *UpdateWorkspaceAssignmentProxyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceAssignmentRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment to be updated in <Databricks>.
	WorkspaceAssignment WorkspaceAssignment `json:"workspace_assignment"`
	// Required. The workspace ID for which the workspace assignment is being
	// updated.
	WorkspaceId int64 `json:"-" url:"-"`
}

func (s *UpdateWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

type UpdateWorkspaceIdentityDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace identity detail to be updated in <Databricks>.
	WorkspaceIdentityDetail WorkspaceIdentityDetail `json:"workspace_identity_detail"`
}

func (s *UpdateWorkspaceIdentityDetailRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

// The details of a User resource.
type User struct {
	// The accountId parent of the user in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// The activity status of a user in a Databricks account.
	AccountUserStatus State `json:"account_user_status"`
	// ExternalId of the user in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`

	FullName UserFullName `json:"full_name"`
	// Internal userId of the user in Databricks.
	UserId string `json:"user_id,omitempty"`
	// Username/email of the user.
	Username string `json:"username"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s User) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UserFullName struct {
	FamilyName string `json:"family_name,omitempty"`

	GivenName string `json:"given_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UserFullName) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UserFullName) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The details of a principal's access to a workspace.
type WorkspaceAccessDetail struct {
	AccessType WorkspaceAccessDetailAccessType `json:"access_type,omitempty"`
	// The account ID parent of the workspace where the principal has access.
	AccountId string `json:"account_id,omitempty"`
	// The permissions granted to the principal in the workspace.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id,omitempty"`

	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The activity status of the principal in the workspace. Not applicable for
	// groups at the moment.
	Status State `json:"status,omitempty"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAccessDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAccessDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of access the principal has to the workspace.
type WorkspaceAccessDetailAccessType string

const WorkspaceAccessDetailAccessTypeDirect WorkspaceAccessDetailAccessType = `DIRECT`

const WorkspaceAccessDetailAccessTypeIndirect WorkspaceAccessDetailAccessType = `INDIRECT`

// String representation for [fmt.Print]
func (f *WorkspaceAccessDetailAccessType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceAccessDetailAccessType) Set(v string) error {
	switch v {
	case `DIRECT`, `INDIRECT`:
		*f = WorkspaceAccessDetailAccessType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT", "INDIRECT"`, v)
	}
}

// Values returns all possible values for WorkspaceAccessDetailAccessType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceAccessDetailAccessType) Values() []WorkspaceAccessDetailAccessType {
	return []WorkspaceAccessDetailAccessType{
		WorkspaceAccessDetailAccessTypeDirect,
		WorkspaceAccessDetailAccessTypeIndirect,
	}
}

// Type always returns WorkspaceAccessDetailAccessType to satisfy [pflag.Value] interface
func (f *WorkspaceAccessDetailAccessType) Type() string {
	return "WorkspaceAccessDetailAccessType"
}

// Controls what fields are returned in the GetWorkspaceAccessDetail response.
type WorkspaceAccessDetailView string

const WorkspaceAccessDetailViewBasic WorkspaceAccessDetailView = `BASIC`

const WorkspaceAccessDetailViewFull WorkspaceAccessDetailView = `FULL`

// String representation for [fmt.Print]
func (f *WorkspaceAccessDetailView) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceAccessDetailView) Set(v string) error {
	switch v {
	case `BASIC`, `FULL`:
		*f = WorkspaceAccessDetailView(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BASIC", "FULL"`, v)
	}
}

// Values returns all possible values for WorkspaceAccessDetailView.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceAccessDetailView) Values() []WorkspaceAccessDetailView {
	return []WorkspaceAccessDetailView{
		WorkspaceAccessDetailViewBasic,
		WorkspaceAccessDetailViewFull,
	}
}

// Type always returns WorkspaceAccessDetailView to satisfy [pflag.Value] interface
func (f *WorkspaceAccessDetailView) Type() string {
	return "WorkspaceAccessDetailView"
}

// The direct assignment of a provisioned account-level principal (user, service
// principal, or group) to a workspace, together with the entitlements that
// assignment grants in the workspace.
//
// This resource covers only principals assigned directly to the workspace.
// Principals that inherit workspace access through a group are not represented
// here. See WorkspaceAccessDetail and WorkspaceIdentityDetail for the
// effective, direct-or-indirect view. Creating the resource assigns the
// principal to the workspace, and deleting it removes the assignment.
//
// `entitlements` is the only client-settable field. It holds the entitlements
// granted directly on this assignment, including any the principal also holds
// through a group. `effective_entitlements` is the read-only union of those and
// any granted through group membership.
//
// A direct assignment always carries at least one directly-assigned
// entitlement, because the assignment is what grants it. Create and update both
// reject an empty `entitlements` set. To remove a principal's assignment
// entirely, delete the resource.
//
// This resource replaces workspace assignment previously managed through the
// workspace SCIM and permission-assignment APIs, and is intended for account
// and workspace admins.
type WorkspaceAssignment struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId string `json:"account_id,omitempty"`
	// Every entitlement the principal holds in this workspace, whether granted
	// directly or through group membership. Get responses populate this field.
	// List responses leave it empty.
	EffectiveEntitlements []Entitlement `json:"effective_entitlements,omitempty"`
	// Entitlements granted directly to the principal on this workspace. This is
	// the only client-settable field. Create and update manage exactly this
	// set, including entitlements the principal also holds through a group.
	// List responses leave this field empty. Get a single principal to read its
	// entitlements.
	Entitlements []Entitlement `json:"entitlements,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id"`
	// The type of the principal (user/service principal/group) that is
	// assigned.
	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The workspace ID where the principal is assigned.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The direct assignment of a provisioned account-level principal (user, service
// principal, or group) to a workspace, together with the entitlements that
// assignment grants in the workspace.
//
// This resource covers only principals assigned directly to the workspace.
// Principals that inherit workspace access through a group are not represented
// here. See WorkspaceAccessDetail and WorkspaceIdentityDetail for the
// effective, direct-or-indirect view. Creating the resource assigns the
// principal to the workspace, and deleting it removes the assignment.
//
// `entitlements` is the only client-settable field. It holds the entitlements
// granted directly on this assignment, including any the principal also holds
// through a group. `effective_entitlements` is the read-only union of those and
// any granted through group membership.
//
// A direct assignment always carries at least one directly-assigned
// entitlement, because the assignment is what grants it. Create and update both
// reject an empty `entitlements` set. To remove a principal's assignment
// entirely, delete the resource.
//
// This resource replaces workspace assignment previously managed through the
// workspace SCIM and permission-assignment APIs, and is intended for account
// and workspace admins.
type WorkspaceAssignmentDetail struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId string `json:"account_id,omitempty"`
	// Every entitlement the principal holds in this workspace, whether granted
	// directly or through group membership. Get responses populate this field.
	// List responses leave it empty.
	EffectiveEntitlements []Entitlement `json:"effective_entitlements,omitempty"`
	// Entitlements granted directly to the principal on this workspace. This is
	// the only client-settable field. Create and update manage exactly this
	// set, including entitlements the principal also holds through a group.
	// List responses leave this field empty. Get a single principal to read its
	// entitlements.
	Entitlements []Entitlement `json:"entitlements,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id"`

	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The workspace ID where the principal is assigned.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAssignmentDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAssignmentDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The details of a directly or indirectly assigned principal's details in a
// workspace.
type WorkspaceIdentityDetail struct {
	// The type of assignment the principal has to the workspace (direct or
	// indirect).
	AssignmentType WorkspaceIdentityDetailAssignmentType `json:"assignment_type,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id,omitempty"`
	// The type of the principal (user/service principal/group).
	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The activity status of an identity in a Databricks workspace.
	WorkspaceIdentityStatus State `json:"workspace_identity_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceIdentityDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceIdentityDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of assignment the principal has to the workspace.
type WorkspaceIdentityDetailAssignmentType string

const WorkspaceIdentityDetailAssignmentTypeDirect WorkspaceIdentityDetailAssignmentType = `DIRECT`

const WorkspaceIdentityDetailAssignmentTypeIndirect WorkspaceIdentityDetailAssignmentType = `INDIRECT`

// String representation for [fmt.Print]
func (f *WorkspaceIdentityDetailAssignmentType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceIdentityDetailAssignmentType) Set(v string) error {
	switch v {
	case `DIRECT`, `INDIRECT`:
		*f = WorkspaceIdentityDetailAssignmentType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT", "INDIRECT"`, v)
	}
}

// Values returns all possible values for WorkspaceIdentityDetailAssignmentType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceIdentityDetailAssignmentType) Values() []WorkspaceIdentityDetailAssignmentType {
	return []WorkspaceIdentityDetailAssignmentType{
		WorkspaceIdentityDetailAssignmentTypeDirect,
		WorkspaceIdentityDetailAssignmentTypeIndirect,
	}
}

// Type always returns WorkspaceIdentityDetailAssignmentType to satisfy [pflag.Value] interface
func (f *WorkspaceIdentityDetailAssignmentType) Type() string {
	return "WorkspaceIdentityDetailAssignmentType"
}

// The type of permission a principal has to a workspace (admin/user).
type WorkspacePermission string

const WorkspacePermissionAdminPermission WorkspacePermission = `ADMIN_PERMISSION`

const WorkspacePermissionUserPermission WorkspacePermission = `USER_PERMISSION`

// String representation for [fmt.Print]
func (f *WorkspacePermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspacePermission) Set(v string) error {
	switch v {
	case `ADMIN_PERMISSION`, `USER_PERMISSION`:
		*f = WorkspacePermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADMIN_PERMISSION", "USER_PERMISSION"`, v)
	}
}

// Values returns all possible values for WorkspacePermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspacePermission) Values() []WorkspacePermission {
	return []WorkspacePermission{
		WorkspacePermissionAdminPermission,
		WorkspacePermissionUserPermission,
	}
}

// Type always returns WorkspacePermission to satisfy [pflag.Value] interface
func (f *WorkspacePermission) Type() string {
	return "WorkspacePermission"
}
