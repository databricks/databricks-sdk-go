// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iamv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
}

type CreateWorkspaceAssignmentDetailRequest struct {
	// Required. Workspace assignment detail to be created in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
	// Required. The workspace ID for which the workspace assignment detail is
	// being created.
	WorkspaceId int64 `json:"-" url:"-"`
}

type DeleteWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
}

type DeleteWorkspaceAssignmentDetailRequest struct {
	// Required. ID of the principal in Databricks to delete workspace
	// assignment for.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID where the principal has access.
	WorkspaceId int64 `json:"-" url:"-"`
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

type GetWorkspaceAssignmentDetailProxyRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
}

type GetWorkspaceAssignmentDetailRequest struct {
	// Required. The internal ID of the principal (user/sp/group) for which the
	// assignment details are being requested.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The workspace ID for which the assignment details are being
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

type ListWorkspaceAssignmentDetailsProxyRequest struct {
	// The maximum number of workspace assignment details to return. The service
	// may return fewer than this value.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// ListWorkspaceAssignmentDetailsProxy call. Provide this to retrieve the
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
// groupname, and inherited parent groups.
type ResolveGroupProxyRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId string `json:"external_id"`
}

// Request message for resolving a group with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the group's
// groupname, and inherited parent groups.
type ResolveGroupRequest struct {
	// Required. The external ID of the group in the customer's IdP.
	ExternalId string `json:"external_id"`
}

type ResolveGroupResponse struct {
	// The group that was resolved.
	Group *Group `json:"group,omitempty"`
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalProxyRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId string `json:"external_id"`
}

// Request message for resolving a service principal with the given external ID
// from the customer's IdP into Databricks. Will resolve metadata such as the
// service principal's displayname, status, and inherited parent groups.
type ResolveServicePrincipalRequest struct {
	// Required. The external ID of the service principal in the customer's IdP.
	ExternalId string `json:"external_id"`
}

type ResolveServicePrincipalResponse struct {
	// The service principal that was resolved.
	ServicePrincipal *ServicePrincipal `json:"service_principal,omitempty"`
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserProxyRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId string `json:"external_id"`
}

// Request message for resolving a user with the given external ID from the
// customer's IdP into Databricks. Will resolve metadata such as the user's
// displayname, status, and inherited parent groups.
type ResolveUserRequest struct {
	// Required. The external ID of the user in the customer's IdP.
	ExternalId string `json:"external_id"`
}

type ResolveUserResponse struct {
	// The user that was resolved.
	User *User `json:"user,omitempty"`
}

// The details of a ServicePrincipal resource.
type ServicePrincipal struct {
	// The parent account ID for the service principal in Databricks.
	AccountId string `json:"account_id,omitempty"`
	// The activity status of a service principal in a Databricks account.
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

type UpdateWorkspaceAssignmentDetailProxyRequest struct {
	// Required. ID of the principal in Databricks.
	PrincipalId int64 `json:"-" url:"-"`
	// Required. The list of fields to update.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
	// Required. Workspace assignment detail to be updated in <Databricks>.
	WorkspaceAssignmentDetail WorkspaceAssignmentDetail `json:"workspace_assignment_detail"`
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
	Username string `json:"username,omitempty"`

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

// The details of a principal's assignment to a workspace.
type WorkspaceAssignmentDetail struct {
	// The account ID parent of the workspace where the principal is assigned
	AccountId string `json:"account_id,omitempty"`
	// Entitlements granted directly to the principal on this workspace. The
	// only client-settable field: create and update manage exactly this set
	// (including entitlements the principal also holds via a group).
	Entitlements []Entitlement `json:"entitlements,omitempty"`
	// The internal ID of the principal (user/sp/group) in Databricks.
	PrincipalId int64 `json:"principal_id"`

	PrincipalType PrincipalType `json:"principal_type,omitempty"`
	// The workspace ID where the principal is assigned
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceAssignmentDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceAssignmentDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
