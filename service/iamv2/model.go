// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateGroupProxyRequest struct {
	// Required. Group to be created in <Databricks>
	Group Group `json:"group"`
}

type CreateGroupRequest struct {
	// Required. Group to be created in <Databricks>
	Group Group `json:"group"`
}

type CreateServicePrincipalProxyRequest struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
}

type CreateServicePrincipalRequest struct {
	// Required. Service principal to be created in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
}

type CreateUserProxyRequest struct {
	// Required. User to be created in <Databricks>
	User User `json:"user"`
}

type CreateUserRequest struct {
	// Required. User to be created in <Databricks>
	User User `json:"user"`
}

type CreateWorkspaceAccessDetailLocalRequest struct {
	// Required. Workspace access detail to be created in <Databricks>.
	WorkspaceAccessDetail WorkspaceAccessDetail `json:"workspace_access_detail"`
}

type CreateWorkspaceAccessDetailRequest struct {
	// Required. The parent path for workspace access detail.
	Parent string `json:"-" url:"-"`
	// Required. Workspace access detail to be created in <Databricks>.
	WorkspaceAccessDetail WorkspaceAccessDetail `json:"workspace_access_detail"`
}

type DeleteGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type DeleteGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type DeleteServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type DeleteServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type DeleteUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type DeleteUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type DeleteWorkspaceAccessDetailLocalRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
}

type DeleteWorkspaceAccessDetailRequest struct {
	// Required. ID of the principal in Databricks to delete workspace access
	// for.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"-" url:"-"`
}

type GetGroupProxyRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type GetGroupRequest struct {
	// Required. Internal ID of the group in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type GetServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type GetServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type GetUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type GetUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId int64 `json:"-" url:"-"`
}

type GetWorkspaceAccessDetailLocalRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// access details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Controls what fields are returned.
	View WorkspaceAccessDetailView `json:"-" url:"view,omitempty"`
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

// The details of a Group resource.
type Group struct {
	// The parent account ID for group in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// ExternalId of the group in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`
	// Display name of the group.
	GroupName string `json:"group_name,omitempty"`
	// Internal group ID of the group in Databricks.
	InternalId int64 `json:"internal_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Group) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Group) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListGroupsProxyRequest struct {
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

// TODO: Write description later when this method is implemented
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
	// The maximum number of sps to return. The service may return fewer than
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

// TODO: Write description later when this method is implemented
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

type ListUsersProxyRequest struct {
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

// TODO: Write description later when this method is implemented
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

type ListWorkspaceAccessDetailsLocalRequest struct {
	// The maximum number of workspace access details to return. The service may
	// return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListWorkspaceAccessDetails call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAccessDetailsLocalRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAccessDetailsLocalRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceAccessDetailsRequest struct {
	// The maximum number of workspace access details to return. The service may
	// return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous ListWorkspaceAccessDetails call.
	// Provide this to retrieve the subsequent page.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The workspace ID for which the workspace access details are being
	// fetched.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAccessDetailsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAccessDetailsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// TODO: Write description later when this method is implemented
type ListWorkspaceAccessDetailsResponse struct {
	// A token, which can be sent as page_token to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	WorkspaceAccessDetails []WorkspaceAccessDetail `json:"workspace_access_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceAccessDetailsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceAccessDetailsResponse) MarshalJSON() ([]byte, error) {
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

// The details of a ServicePrincipal resource.
type ServicePrincipal struct {
	// The parent account ID for the service principal in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// The activity status of a sp in a Databricks account.
	AccountSpStatus State `json:"account_sp_status,omitempty"`
	// Application ID of the service principal.
	ApplicationId string `json:"application_id,omitempty"`
	// Display name of the service principal.
	DisplayName string `json:"display_name,omitempty"`
	// ExternalId of the service principal in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`
	// Internal service principal ID of the service principal in Databricks.
	InternalId int64 `json:"internal_id,omitempty"`

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

// Request message for syncing a group with the given external ID from the
// customer's IdP into Databricks. Will sync metadata such as the group's
// groupname, and inherited parent groups.
type SyncGroupProxyRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId string `json:"external_id"`
}

// Request message for syncing a group with the given external ID from the
// customer's IdP into Databricks. Will sync metadata such as the group's
// groupname, and inherited parent groups.
type SyncGroupRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId string `json:"external_id"`
}

type SyncGroupResponse struct {
	// The group that was synced.
	Group *Group `json:"group,omitempty"`
}

// Request message for syncing a service principal with the given external ID
// from the customer's IdP into Databricks. Will sync metadata such as the
// service principal's displayname, status, and inherited parent groups.
type SyncServicePrincipalProxyRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId string `json:"external_id"`
}

// Request message for syncing a service principal with the given external ID
// from the customer's IdP into Databricks. Will sync metadata such as the
// service principal's displayname, status, and inherited parent groups.
type SyncServicePrincipalRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId string `json:"external_id"`
}

type SyncServicePrincipalResponse struct {
	// The service principal that was synced.
	ServicePrincipal *ServicePrincipal `json:"service_principal,omitempty"`
}

// Request message for syncing a user with the given external ID from the
// customer's IdP into Databricks. Will sync metadata such as the user's
// displayname, status, and inherited parent groups.
type SyncUserProxyRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId string `json:"external_id"`
}

// Request message for syncing a user with the given external ID from the
// customer's IdP into Databricks. Will sync metadata such as the user's
// displayname, status, and inherited parent groups.
type SyncUserRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId string `json:"external_id"`
}

type SyncUserResponse struct {
	// The user that was synced.
	User *User `json:"user,omitempty"`
}

type UpdateGroupProxyRequest struct {
	// Required. Group to be updated in <Databricks>
	Group Group `json:"group"`
	// Required. Internal ID of the group in Databricks.
	InternalId int64 `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateGroupRequest struct {
	// Required. Group to be updated in <Databricks>
	Group Group `json:"group"`
	// Required. Internal ID of the group in Databricks.
	InternalId int64 `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateServicePrincipalProxyRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId int64 `json:"-" url:"-"`
	// Required. Service principal to be updated in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateServicePrincipalRequest struct {
	// Required. Internal ID of the service principal in Databricks.
	InternalId int64 `json:"-" url:"-"`
	// Required. Service Principal to be updated in <Databricks>
	ServicePrincipal ServicePrincipal `json:"service_principal"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateUserProxyRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId int64 `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
	// Required. User to be updated in <Databricks>
	User User `json:"user"`
}

type UpdateUserRequest struct {
	// Required. Internal ID of the user in Databricks.
	InternalId int64 `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
	// Required. User to be updated in <Databricks>
	User User `json:"user"`
}

type UpdateWorkspaceAccessDetailLocalRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
	// Required. WorkspaceAccessDetail to be updated in <Databricks>
	WorkspaceAccessDetail WorkspaceAccessDetail `json:"workspace_access_detail"`
}

type UpdateWorkspaceAccessDetailRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Optional. The list of fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
	// Required. Workspace access detail to be updated in <Databricks>
	WorkspaceAccessDetail WorkspaceAccessDetail `json:"workspace_access_detail"`
	// Required. The workspace ID for which the workspace access detail is being
	// updated.
	WorkspaceId int64 `json:"-" url:"-"`
}

// The details of a User resource.
type User struct {
	// The accountId parent of the user in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// The activity status of a user in a Databricks account.
	AccountUserStatus State `json:"account_user_status,omitempty"`
	// ExternalId of the user in the customer's IdP.
	ExternalId string `json:"external_id,omitempty"`
	// Internal userId of the user in Databricks.
	InternalId int64 `json:"internal_id,omitempty"`

	Name *UserName `json:"name,omitempty"`
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

type UserName struct {
	FamilyName string `json:"family_name,omitempty"`

	GivenName string `json:"given_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UserName) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UserName) MarshalJSON() ([]byte, error) {
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
