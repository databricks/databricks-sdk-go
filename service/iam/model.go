// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/iam/iampb"
)

type AccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := AccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.AccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AccessControlRequestToPb(st *AccessControlRequest) (*iampb.AccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.AccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := PermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AccessControlRequestFromPb(pb *iampb.AccessControlRequestPb) (*AccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := PermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []Permission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := AccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.AccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AccessControlResponseToPb(st *AccessControlResponse) (*iampb.AccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.AccessControlResponsePb{}

	var allPermissionsPb []iampb.PermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := PermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AccessControlResponseFromPb(pb *iampb.AccessControlResponsePb) (*AccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControlResponse{}

	var allPermissionsField []Permission
	for _, itemPb := range pb.AllPermissions {
		item, err := PermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// represents an identity trying to access a resource - user or a service
// principal group can be a principal of a permission set assignment but an
// actor is always a user or a service principal
type Actor struct {

	// Wire name: 'actor_id'
	ActorId         int64    `json:"actor_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Actor) MarshalJSON() ([]byte, error) {
	pb, err := ActorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Actor) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ActorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ActorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ActorToPb(st *Actor) (*iampb.ActorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ActorPb{}
	pb.ActorId = st.ActorId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ActorFromPb(pb *iampb.ActorPb) (*Actor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Actor{}
	st.ActorId = pb.ActorId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CheckPolicyRequest struct {
	Actor Actor `json:"-" tf:"-"`

	AuthzIdentity RequestAuthzIdentity `json:"-" tf:"-"`

	ConsistencyToken ConsistencyToken `json:"-" tf:"-"`

	Permission string `json:"-" tf:"-"`
	// Ex: (servicePrincipal/use,
	// accounts/<account-id>/servicePrincipals/<sp-id>) Ex:
	// (servicePrincipal.ruleSet/update,
	// accounts/<account-id>/servicePrincipals/<sp-id>/ruleSets/default)
	Resource string `json:"-" tf:"-"`

	ResourceInfo *ResourceInfo `json:"-" tf:"-"`
}

func (st CheckPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := CheckPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CheckPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.CheckPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CheckPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CheckPolicyRequestToPb(st *CheckPolicyRequest) (*iampb.CheckPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.CheckPolicyRequestPb{}
	actorPb, err := ActorToPb(&st.Actor)
	if err != nil {
		return nil, err
	}
	if actorPb != nil {
		pb.Actor = *actorPb
	}
	authzIdentityPb, err := RequestAuthzIdentityToPb(&st.AuthzIdentity)
	if err != nil {
		return nil, err
	}
	if authzIdentityPb != nil {
		pb.AuthzIdentity = *authzIdentityPb
	}
	consistencyTokenPb, err := ConsistencyTokenToPb(&st.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenPb != nil {
		pb.ConsistencyToken = *consistencyTokenPb
	}
	pb.Permission = st.Permission
	pb.Resource = st.Resource
	resourceInfoPb, err := ResourceInfoToPb(st.ResourceInfo)
	if err != nil {
		return nil, err
	}
	if resourceInfoPb != nil {
		pb.ResourceInfo = resourceInfoPb
	}

	return pb, nil
}

func CheckPolicyRequestFromPb(pb *iampb.CheckPolicyRequestPb) (*CheckPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CheckPolicyRequest{}
	actorField, err := ActorFromPb(&pb.Actor)
	if err != nil {
		return nil, err
	}
	if actorField != nil {
		st.Actor = *actorField
	}
	authzIdentityField, err := RequestAuthzIdentityFromPb(&pb.AuthzIdentity)
	if err != nil {
		return nil, err
	}
	if authzIdentityField != nil {
		st.AuthzIdentity = *authzIdentityField
	}
	consistencyTokenField, err := ConsistencyTokenFromPb(&pb.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenField != nil {
		st.ConsistencyToken = *consistencyTokenField
	}
	st.Permission = pb.Permission
	st.Resource = pb.Resource
	resourceInfoField, err := ResourceInfoFromPb(pb.ResourceInfo)
	if err != nil {
		return nil, err
	}
	if resourceInfoField != nil {
		st.ResourceInfo = resourceInfoField
	}

	return st, nil
}

type CheckPolicyResponse struct {

	// Wire name: 'consistency_token'
	ConsistencyToken ConsistencyToken `json:"consistency_token"`

	// Wire name: 'is_permitted'
	IsPermitted     bool     `json:"is_permitted,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CheckPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := CheckPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CheckPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.CheckPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CheckPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CheckPolicyResponseToPb(st *CheckPolicyResponse) (*iampb.CheckPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.CheckPolicyResponsePb{}
	consistencyTokenPb, err := ConsistencyTokenToPb(&st.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenPb != nil {
		pb.ConsistencyToken = *consistencyTokenPb
	}
	pb.IsPermitted = st.IsPermitted

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CheckPolicyResponseFromPb(pb *iampb.CheckPolicyResponsePb) (*CheckPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CheckPolicyResponse{}
	consistencyTokenField, err := ConsistencyTokenFromPb(&pb.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenField != nil {
		st.ConsistencyToken = *consistencyTokenField
	}
	st.IsPermitted = pb.IsPermitted

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ComplexValue struct {

	// Wire name: 'display'
	Display string `json:"display,omitempty"`

	// Wire name: 'primary'
	Primary bool `json:"primary,omitempty"`

	// Wire name: '$ref'
	Ref string `json:"$ref,omitempty"`

	// Wire name: 'type'
	Type string `json:"type,omitempty"`

	// Wire name: 'value'
	Value           string   `json:"value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ComplexValue) MarshalJSON() ([]byte, error) {
	pb, err := ComplexValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComplexValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ComplexValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComplexValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComplexValueToPb(st *ComplexValue) (*iampb.ComplexValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ComplexValuePb{}
	pb.Display = st.Display
	pb.Primary = st.Primary
	pb.Ref = st.Ref
	pb.Type = st.Type
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComplexValueFromPb(pb *iampb.ComplexValuePb) (*ComplexValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplexValue{}
	st.Display = pb.Display
	st.Primary = pb.Primary
	st.Ref = pb.Ref
	st.Type = pb.Type
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ConsistencyToken struct {

	// Wire name: 'value'
	Value string `json:"value"`
}

func (st ConsistencyToken) MarshalJSON() ([]byte, error) {
	pb, err := ConsistencyTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ConsistencyToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ConsistencyTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ConsistencyTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ConsistencyTokenToPb(st *ConsistencyToken) (*iampb.ConsistencyTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ConsistencyTokenPb{}
	pb.Value = st.Value

	return pb, nil
}

func ConsistencyTokenFromPb(pb *iampb.ConsistencyTokenPb) (*ConsistencyToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConsistencyToken{}
	st.Value = pb.Value

	return st, nil
}

type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" tf:"-"`
}

func (st DeleteAccountGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteAccountGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountGroupRequestToPb(st *DeleteAccountGroupRequest) (*iampb.DeleteAccountGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteAccountGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteAccountGroupRequestFromPb(pb *iampb.DeleteAccountGroupRequestPb) (*DeleteAccountGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" tf:"-"`
}

func (st DeleteAccountServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteAccountServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountServicePrincipalRequestToPb(st *DeleteAccountServicePrincipalRequest) (*iampb.DeleteAccountServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteAccountServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteAccountServicePrincipalRequestFromPb(pb *iampb.DeleteAccountServicePrincipalRequestPb) (*DeleteAccountServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" tf:"-"`
}

func (st DeleteAccountUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteAccountUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountUserRequestToPb(st *DeleteAccountUserRequest) (*iampb.DeleteAccountUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteAccountUserRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteAccountUserRequestFromPb(pb *iampb.DeleteAccountUserRequestPb) (*DeleteAccountUserRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountUserRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" tf:"-"`
}

func (st DeleteGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteGroupRequestToPb(st *DeleteGroupRequest) (*iampb.DeleteGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteGroupRequestFromPb(pb *iampb.DeleteGroupRequestPb) (*DeleteGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" tf:"-"`
}

func (st DeleteServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteServicePrincipalRequestToPb(st *DeleteServicePrincipalRequest) (*iampb.DeleteServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteServicePrincipalRequestFromPb(pb *iampb.DeleteServicePrincipalRequestPb) (*DeleteServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" tf:"-"`
}

func (st DeleteUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteUserRequestToPb(st *DeleteUserRequest) (*iampb.DeleteUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteUserRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteUserRequestFromPb(pb *iampb.DeleteUserRequestPb) (*DeleteUserRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteUserRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteWorkspaceAssignmentRequest struct {
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" tf:"-"`
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st DeleteWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteWorkspaceAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.DeleteWorkspaceAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteWorkspaceAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteWorkspaceAssignmentRequestToPb(st *DeleteWorkspaceAssignmentRequest) (*iampb.DeleteWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.DeleteWorkspaceAssignmentRequestPb{}
	pb.PrincipalId = st.PrincipalId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func DeleteWorkspaceAssignmentRequestFromPb(pb *iampb.DeleteWorkspaceAssignmentRequestPb) (*DeleteWorkspaceAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWorkspaceAssignmentRequest{}
	st.PrincipalId = pb.PrincipalId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" tf:"-"`
}

func (st GetAccountGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAccountGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAccountGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetAccountGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAccountGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAccountGroupRequestToPb(st *GetAccountGroupRequest) (*iampb.GetAccountGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetAccountGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetAccountGroupRequestFromPb(pb *iampb.GetAccountGroupRequestPb) (*GetAccountGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" tf:"-"`
}

func (st GetAccountServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAccountServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAccountServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetAccountServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAccountServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAccountServicePrincipalRequestToPb(st *GetAccountServicePrincipalRequest) (*iampb.GetAccountServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetAccountServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetAccountServicePrincipalRequestFromPb(pb *iampb.GetAccountServicePrincipalRequestPb) (*GetAccountServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetAccountUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page. Default is 10000.
	Count int `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder GetSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int      `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetAccountUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAccountUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAccountUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetAccountUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAccountUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAccountUserRequestToPb(st *GetAccountUserRequest) (*iampb.GetAccountUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetAccountUserRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.Id = st.Id
	pb.SortBy = st.SortBy
	sortOrderPb, err := GetSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetAccountUserRequestFromPb(pb *iampb.GetAccountUserRequestPb) (*GetAccountUserRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountUserRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.Id = pb.Id
	st.SortBy = pb.SortBy
	sortOrderField, err := GetSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	//
	// Examples | Summary :--- | :--- `resource=accounts/<ACCOUNT_ID>` | A
	// resource name for the account.
	// `resource=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>` | A resource name for
	// the group. `resource=accounts/<ACCOUNT_ID>/servicePrincipals/<SP_ID>` | A
	// resource name for the service principal.
	Resource string `json:"-" tf:"-"`
}

func (st GetAssignableRolesForResourceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAssignableRolesForResourceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAssignableRolesForResourceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetAssignableRolesForResourceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAssignableRolesForResourceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAssignableRolesForResourceRequestToPb(st *GetAssignableRolesForResourceRequest) (*iampb.GetAssignableRolesForResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetAssignableRolesForResourceRequestPb{}
	pb.Resource = st.Resource

	return pb, nil
}

func GetAssignableRolesForResourceRequestFromPb(pb *iampb.GetAssignableRolesForResourceRequestPb) (*GetAssignableRolesForResourceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAssignableRolesForResourceRequest{}
	st.Resource = pb.Resource

	return st, nil
}

type GetAssignableRolesForResourceResponse struct {

	// Wire name: 'roles'
	Roles []Role `json:"roles,omitempty"`
}

func (st GetAssignableRolesForResourceResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetAssignableRolesForResourceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAssignableRolesForResourceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetAssignableRolesForResourceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAssignableRolesForResourceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAssignableRolesForResourceResponseToPb(st *GetAssignableRolesForResourceResponse) (*iampb.GetAssignableRolesForResourceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetAssignableRolesForResourceResponsePb{}

	var rolesPb []iampb.RolePb
	for _, item := range st.Roles {
		itemPb, err := RoleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	return pb, nil
}

func GetAssignableRolesForResourceResponseFromPb(pb *iampb.GetAssignableRolesForResourceResponsePb) (*GetAssignableRolesForResourceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAssignableRolesForResourceResponse{}

	var rolesField []Role
	for _, itemPb := range pb.Roles {
		item, err := RoleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField

	return st, nil
}

type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" tf:"-"`
}

func (st GetGroupRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetGroupRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetGroupRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetGroupRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetGroupRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetGroupRequestToPb(st *GetGroupRequest) (*iampb.GetGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetGroupRequestFromPb(pb *iampb.GetGroupRequestPb) (*GetGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetPasswordPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []PasswordPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetPasswordPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetPasswordPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPasswordPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetPasswordPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPasswordPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPasswordPermissionLevelsResponseToPb(st *GetPasswordPermissionLevelsResponse) (*iampb.GetPasswordPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetPasswordPermissionLevelsResponsePb{}

	var permissionLevelsPb []iampb.PasswordPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := PasswordPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetPasswordPermissionLevelsResponseFromPb(pb *iampb.GetPasswordPermissionLevelsResponsePb) (*GetPasswordPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPasswordPermissionLevelsResponse{}

	var permissionLevelsField []PasswordPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := PasswordPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetPermissionLevelsRequest struct {
	RequestObjectId string `json:"-" tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" tf:"-"`
}

func (st GetPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPermissionLevelsRequestToPb(st *GetPermissionLevelsRequest) (*iampb.GetPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetPermissionLevelsRequestPb{}
	pb.RequestObjectId = st.RequestObjectId
	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
}

func GetPermissionLevelsRequestFromPb(pb *iampb.GetPermissionLevelsRequestPb) (*GetPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionLevelsRequest{}
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
}

type GetPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPermissionLevelsResponseToPb(st *GetPermissionLevelsResponse) (*iampb.GetPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetPermissionLevelsResponsePb{}

	var permissionLevelsPb []iampb.PermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := PermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetPermissionLevelsResponseFromPb(pb *iampb.GetPermissionLevelsResponsePb) (*GetPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionLevelsResponse{}

	var permissionLevelsField []PermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := PermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetPermissionRequest struct {
	// The id of the request object.
	RequestObjectId string `json:"-" tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" tf:"-"`
}

func (st GetPermissionRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPermissionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPermissionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetPermissionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPermissionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPermissionRequestToPb(st *GetPermissionRequest) (*iampb.GetPermissionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetPermissionRequestPb{}
	pb.RequestObjectId = st.RequestObjectId
	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
}

func GetPermissionRequestFromPb(pb *iampb.GetPermissionRequestPb) (*GetPermissionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionRequest{}
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
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
	Etag string `json:"-" tf:"-"`
	// The ruleset name associated with the request.
	//
	// Examples | Summary :--- | :---
	// `name=accounts/<ACCOUNT_ID>/ruleSets/default` | A name for a rule set on
	// the account.
	// `name=accounts/<ACCOUNT_ID>/groups/<GROUP_ID>/ruleSets/default` | A name
	// for a rule set on the group.
	// `name=accounts/<ACCOUNT_ID>/servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>/ruleSets/default`
	// | A name for a rule set on the service principal.
	Name string `json:"-" tf:"-"`
}

func (st GetRuleSetRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRuleSetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRuleSetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetRuleSetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRuleSetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRuleSetRequestToPb(st *GetRuleSetRequest) (*iampb.GetRuleSetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetRuleSetRequestPb{}
	pb.Etag = st.Etag
	pb.Name = st.Name

	return pb, nil
}

func GetRuleSetRequestFromPb(pb *iampb.GetRuleSetRequestPb) (*GetRuleSetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRuleSetRequest{}
	st.Etag = pb.Etag
	st.Name = pb.Name

	return st, nil
}

type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" tf:"-"`
}

func (st GetServicePrincipalRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetServicePrincipalRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetServicePrincipalRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetServicePrincipalRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetServicePrincipalRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetServicePrincipalRequestToPb(st *GetServicePrincipalRequest) (*iampb.GetServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetServicePrincipalRequestFromPb(pb *iampb.GetServicePrincipalRequestPb) (*GetServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
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

func GetSortOrderToPb(st *GetSortOrder) (*iampb.GetSortOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.GetSortOrderPb(*st)
	return &pb, nil
}

func GetSortOrderFromPb(pb *iampb.GetSortOrderPb) (*GetSortOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetSortOrder(*pb)
	return &st, nil
}

type GetUserRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page.
	Count int `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder GetSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int      `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetUserRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetUserRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetUserRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetUserRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetUserRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetUserRequestToPb(st *GetUserRequest) (*iampb.GetUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetUserRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.Id = st.Id
	pb.SortBy = st.SortBy
	sortOrderPb, err := GetSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetUserRequestFromPb(pb *iampb.GetUserRequestPb) (*GetUserRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetUserRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.Id = pb.Id
	st.SortBy = pb.SortBy
	sortOrderField, err := GetSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st GetWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetWorkspaceAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GetWorkspaceAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetWorkspaceAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetWorkspaceAssignmentRequestToPb(st *GetWorkspaceAssignmentRequest) (*iampb.GetWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GetWorkspaceAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func GetWorkspaceAssignmentRequestFromPb(pb *iampb.GetWorkspaceAssignmentRequestPb) (*GetWorkspaceAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceAssignmentRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type GrantRule struct {
	// Principals this grant rule applies to. A principal can be a user (for end
	// users), a service principal (for applications and compute workloads), or
	// an account group. Each principal has its own identifier format: *
	// users/<USERNAME> * groups/<GROUP_NAME> *
	// servicePrincipals/<SERVICE_PRINCIPAL_APPLICATION_ID>
	// Wire name: 'principals'
	Principals []string `json:"principals,omitempty"`
	// Role that is assigned to the list of principals.
	// Wire name: 'role'
	Role string `json:"role"`
}

func (st GrantRule) MarshalJSON() ([]byte, error) {
	pb, err := GrantRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GrantRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GrantRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GrantRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GrantRuleToPb(st *GrantRule) (*iampb.GrantRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GrantRulePb{}
	pb.Principals = st.Principals
	pb.Role = st.Role

	return pb, nil
}

func GrantRuleFromPb(pb *iampb.GrantRulePb) (*GrantRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GrantRule{}
	st.Principals = pb.Principals
	st.Role = pb.Role

	return st, nil
}

type Group struct {
	// String that represents a human-readable group name
	// Wire name: 'displayName'
	DisplayName string `json:"displayName,omitempty"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	// Wire name: 'entitlements'
	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	// Wire name: 'externalId'
	ExternalId string `json:"externalId,omitempty"`

	// Wire name: 'groups'
	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks group ID
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'members'
	Members []ComplexValue `json:"members,omitempty"`
	// Container for the group identifier. Workspace local versus account.
	// Wire name: 'meta'
	Meta *ResourceMeta `json:"meta,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	// Wire name: 'roles'
	Roles []ComplexValue `json:"roles,omitempty"`
	// The schema of the group.
	// Wire name: 'schemas'
	Schemas         []GroupSchema `json:"schemas,omitempty"`
	ForceSendFields []string      `json:"-" tf:"-"`
}

func (st Group) MarshalJSON() ([]byte, error) {
	pb, err := GroupToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Group) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.GroupPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GroupFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GroupToPb(st *Group) (*iampb.GroupPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.GroupPb{}
	pb.DisplayName = st.DisplayName

	var entitlementsPb []iampb.ComplexValuePb
	for _, item := range st.Entitlements {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entitlementsPb = append(entitlementsPb, *itemPb)
		}
	}
	pb.Entitlements = entitlementsPb
	pb.ExternalId = st.ExternalId

	var groupsPb []iampb.ComplexValuePb
	for _, item := range st.Groups {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			groupsPb = append(groupsPb, *itemPb)
		}
	}
	pb.Groups = groupsPb
	pb.Id = st.Id

	var membersPb []iampb.ComplexValuePb
	for _, item := range st.Members {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			membersPb = append(membersPb, *itemPb)
		}
	}
	pb.Members = membersPb
	metaPb, err := ResourceMetaToPb(st.Meta)
	if err != nil {
		return nil, err
	}
	if metaPb != nil {
		pb.Meta = metaPb
	}

	var rolesPb []iampb.ComplexValuePb
	for _, item := range st.Roles {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	var schemasPb []iampb.GroupSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := GroupSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GroupFromPb(pb *iampb.GroupPb) (*Group, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Group{}
	st.DisplayName = pb.DisplayName

	var entitlementsField []ComplexValue
	for _, itemPb := range pb.Entitlements {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			entitlementsField = append(entitlementsField, *item)
		}
	}
	st.Entitlements = entitlementsField
	st.ExternalId = pb.ExternalId

	var groupsField []ComplexValue
	for _, itemPb := range pb.Groups {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			groupsField = append(groupsField, *item)
		}
	}
	st.Groups = groupsField
	st.Id = pb.Id

	var membersField []ComplexValue
	for _, itemPb := range pb.Members {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			membersField = append(membersField, *item)
		}
	}
	st.Members = membersField
	metaField, err := ResourceMetaFromPb(pb.Meta)
	if err != nil {
		return nil, err
	}
	if metaField != nil {
		st.Meta = metaField
	}

	var rolesField []ComplexValue
	for _, itemPb := range pb.Roles {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField

	var schemasField []GroupSchema
	for _, itemPb := range pb.Schemas {
		item, err := GroupSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func GroupSchemaToPb(st *GroupSchema) (*iampb.GroupSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.GroupSchemaPb(*st)
	return &pb, nil
}

func GroupSchemaFromPb(pb *iampb.GroupSchemaPb) (*GroupSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := GroupSchema(*pb)
	return &st, nil
}

type ListAccountGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page. Default is 10000.
	Count int64 `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Attribute to sort the results.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAccountGroupsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAccountGroupsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAccountGroupsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListAccountGroupsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAccountGroupsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAccountGroupsRequestToPb(st *ListAccountGroupsRequest) (*iampb.ListAccountGroupsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListAccountGroupsRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.SortBy = st.SortBy
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAccountGroupsRequestFromPb(pb *iampb.ListAccountGroupsRequestPb) (*ListAccountGroupsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountGroupsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListAccountServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page. Default is 10000.
	Count int64 `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Attribute to sort the results.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAccountServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAccountServicePrincipalsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAccountServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListAccountServicePrincipalsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAccountServicePrincipalsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAccountServicePrincipalsRequestToPb(st *ListAccountServicePrincipalsRequest) (*iampb.ListAccountServicePrincipalsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListAccountServicePrincipalsRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.SortBy = st.SortBy
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAccountServicePrincipalsRequestFromPb(pb *iampb.ListAccountServicePrincipalsRequestPb) (*ListAccountServicePrincipalsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountServicePrincipalsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListAccountUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page. Default is 10000.
	Count int64 `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAccountUsersRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAccountUsersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAccountUsersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListAccountUsersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAccountUsersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAccountUsersRequestToPb(st *ListAccountUsersRequest) (*iampb.ListAccountUsersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListAccountUsersRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.SortBy = st.SortBy
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAccountUsersRequestFromPb(pb *iampb.ListAccountUsersRequestPb) (*ListAccountUsersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountUsersRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page.
	Count int64 `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Attribute to sort the results.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListGroupsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListGroupsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListGroupsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListGroupsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListGroupsRequestToPb(st *ListGroupsRequest) (*iampb.ListGroupsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListGroupsRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.SortBy = st.SortBy
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListGroupsRequestFromPb(pb *iampb.ListGroupsRequestPb) (*ListGroupsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGroupsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListGroupsResponse struct {
	// Total results returned in the response.
	// Wire name: 'itemsPerPage'
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	// Wire name: 'Resources'
	Resources []Group `json:"Resources,omitempty"`
	// The schema of the service principal.
	// Wire name: 'schemas'
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	// Wire name: 'totalResults'
	TotalResults    int64    `json:"totalResults,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListGroupsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListGroupsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListGroupsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListGroupsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListGroupsResponseToPb(st *ListGroupsResponse) (*iampb.ListGroupsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListGroupsResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage

	var resourcesPb []iampb.GroupPb
	for _, item := range st.Resources {
		itemPb, err := GroupToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	var schemasPb []iampb.ListResponseSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := ListResponseSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb
	pb.StartIndex = st.StartIndex
	pb.TotalResults = st.TotalResults

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListGroupsResponseFromPb(pb *iampb.ListGroupsResponsePb) (*ListGroupsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGroupsResponse{}
	st.ItemsPerPage = pb.ItemsPerPage

	var resourcesField []Group
	for _, itemPb := range pb.Resources {
		item, err := GroupFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField

	var schemasField []ListResponseSchema
	for _, itemPb := range pb.Schemas {
		item, err := ListResponseSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField
	st.StartIndex = pb.StartIndex
	st.TotalResults = pb.TotalResults

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ListResponseSchemaToPb(st *ListResponseSchema) (*iampb.ListResponseSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.ListResponseSchemaPb(*st)
	return &pb, nil
}

func ListResponseSchemaFromPb(pb *iampb.ListResponseSchemaPb) (*ListResponseSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListResponseSchema(*pb)
	return &st, nil
}

type ListServicePrincipalResponse struct {
	// Total results returned in the response.
	// Wire name: 'itemsPerPage'
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	// Wire name: 'Resources'
	Resources []ServicePrincipal `json:"Resources,omitempty"`
	// The schema of the List response.
	// Wire name: 'schemas'
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	// Wire name: 'totalResults'
	TotalResults    int64    `json:"totalResults,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListServicePrincipalResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListServicePrincipalResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListServicePrincipalResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListServicePrincipalResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListServicePrincipalResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListServicePrincipalResponseToPb(st *ListServicePrincipalResponse) (*iampb.ListServicePrincipalResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListServicePrincipalResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage

	var resourcesPb []iampb.ServicePrincipalPb
	for _, item := range st.Resources {
		itemPb, err := ServicePrincipalToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	var schemasPb []iampb.ListResponseSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := ListResponseSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb
	pb.StartIndex = st.StartIndex
	pb.TotalResults = st.TotalResults

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListServicePrincipalResponseFromPb(pb *iampb.ListServicePrincipalResponsePb) (*ListServicePrincipalResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalResponse{}
	st.ItemsPerPage = pb.ItemsPerPage

	var resourcesField []ServicePrincipal
	for _, itemPb := range pb.Resources {
		item, err := ServicePrincipalFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField

	var schemasField []ListResponseSchema
	for _, itemPb := range pb.Schemas {
		item, err := ListResponseSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField
	st.StartIndex = pb.StartIndex
	st.TotalResults = pb.TotalResults

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListServicePrincipalsRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page.
	Count int64 `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Attribute to sort the results.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListServicePrincipalsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListServicePrincipalsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListServicePrincipalsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListServicePrincipalsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListServicePrincipalsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListServicePrincipalsRequestToPb(st *ListServicePrincipalsRequest) (*iampb.ListServicePrincipalsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListServicePrincipalsRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.SortBy = st.SortBy
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListServicePrincipalsRequestFromPb(pb *iampb.ListServicePrincipalsRequestPb) (*ListServicePrincipalsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ListSortOrderToPb(st *ListSortOrder) (*iampb.ListSortOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.ListSortOrderPb(*st)
	return &pb, nil
}

func ListSortOrderFromPb(pb *iampb.ListSortOrderPb) (*ListSortOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListSortOrder(*pb)
	return &st, nil
}

type ListUsersRequest struct {
	// Comma-separated list of attributes to return in response.
	Attributes string `json:"-" tf:"-"`
	// Desired number of results per page.
	Count int64 `json:"-" tf:"-"`
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `json:"-" tf:"-"`
	// Query by which the results have to be filtered. Supported operators are
	// equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`).
	// Additionally, simple expressions can be formed using logical operators -
	// `and` and `or`. The [SCIM RFC] has more details but we currently only
	// support simple expressions.
	//
	// [SCIM RFC]: https://tools.ietf.org/html/rfc7644#section-3.4.2.2
	Filter string `json:"-" tf:"-"`
	// Attribute to sort the results. Multi-part paths are supported. For
	// example, `userName`, `name.givenName`, and `emails`.
	SortBy string `json:"-" tf:"-"`
	// The order to sort the results.
	SortOrder ListSortOrder `json:"-" tf:"-"`
	// Specifies the index of the first result. First item is number 1.
	StartIndex      int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListUsersRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListUsersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListUsersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListUsersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListUsersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListUsersRequestToPb(st *ListUsersRequest) (*iampb.ListUsersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListUsersRequestPb{}
	pb.Attributes = st.Attributes
	pb.Count = st.Count
	pb.ExcludedAttributes = st.ExcludedAttributes
	pb.Filter = st.Filter
	pb.SortBy = st.SortBy
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}
	pb.StartIndex = st.StartIndex

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListUsersRequestFromPb(pb *iampb.ListUsersRequestPb) (*ListUsersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUsersRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}
	st.StartIndex = pb.StartIndex

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListUsersResponse struct {
	// Total results returned in the response.
	// Wire name: 'itemsPerPage'
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	// Wire name: 'Resources'
	Resources []User `json:"Resources,omitempty"`
	// The schema of the List response.
	// Wire name: 'schemas'
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	// Wire name: 'startIndex'
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	// Wire name: 'totalResults'
	TotalResults    int64    `json:"totalResults,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListUsersResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListUsersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListUsersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListUsersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListUsersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListUsersResponseToPb(st *ListUsersResponse) (*iampb.ListUsersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListUsersResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage

	var resourcesPb []iampb.UserPb
	for _, item := range st.Resources {
		itemPb, err := UserToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	var schemasPb []iampb.ListResponseSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := ListResponseSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb
	pb.StartIndex = st.StartIndex
	pb.TotalResults = st.TotalResults

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListUsersResponseFromPb(pb *iampb.ListUsersResponsePb) (*ListUsersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUsersResponse{}
	st.ItemsPerPage = pb.ItemsPerPage

	var resourcesField []User
	for _, itemPb := range pb.Resources {
		item, err := UserFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField

	var schemasField []ListResponseSchema
	for _, itemPb := range pb.Schemas {
		item, err := ListResponseSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField
	st.StartIndex = pb.StartIndex
	st.TotalResults = pb.TotalResults

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st ListWorkspaceAssignmentRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListWorkspaceAssignmentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListWorkspaceAssignmentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ListWorkspaceAssignmentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListWorkspaceAssignmentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListWorkspaceAssignmentRequestToPb(st *ListWorkspaceAssignmentRequest) (*iampb.ListWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ListWorkspaceAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func ListWorkspaceAssignmentRequestFromPb(pb *iampb.ListWorkspaceAssignmentRequestPb) (*ListWorkspaceAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWorkspaceAssignmentRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type MigratePermissionsRequest struct {
	// The name of the workspace group that permissions will be migrated from.
	// Wire name: 'from_workspace_group_name'
	FromWorkspaceGroupName string `json:"from_workspace_group_name"`
	// The maximum number of permissions that will be migrated.
	// Wire name: 'size'
	Size int `json:"size,omitempty"`
	// The name of the account group that permissions will be migrated to.
	// Wire name: 'to_account_group_name'
	ToAccountGroupName string `json:"to_account_group_name"`
	// WorkspaceId of the associated workspace where the permission migration
	// will occur.
	// Wire name: 'workspace_id'
	WorkspaceId     int64    `json:"workspace_id"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MigratePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := MigratePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *MigratePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.MigratePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MigratePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MigratePermissionsRequestToPb(st *MigratePermissionsRequest) (*iampb.MigratePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.MigratePermissionsRequestPb{}
	pb.FromWorkspaceGroupName = st.FromWorkspaceGroupName
	pb.Size = st.Size
	pb.ToAccountGroupName = st.ToAccountGroupName
	pb.WorkspaceId = st.WorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func MigratePermissionsRequestFromPb(pb *iampb.MigratePermissionsRequestPb) (*MigratePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MigratePermissionsRequest{}
	st.FromWorkspaceGroupName = pb.FromWorkspaceGroupName
	st.Size = pb.Size
	st.ToAccountGroupName = pb.ToAccountGroupName
	st.WorkspaceId = pb.WorkspaceId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	// Wire name: 'permissions_migrated'
	PermissionsMigrated int      `json:"permissions_migrated,omitempty"`
	ForceSendFields     []string `json:"-" tf:"-"`
}

func (st MigratePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := MigratePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *MigratePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.MigratePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MigratePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MigratePermissionsResponseToPb(st *MigratePermissionsResponse) (*iampb.MigratePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.MigratePermissionsResponsePb{}
	pb.PermissionsMigrated = st.PermissionsMigrated

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func MigratePermissionsResponseFromPb(pb *iampb.MigratePermissionsResponsePb) (*MigratePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MigratePermissionsResponse{}
	st.PermissionsMigrated = pb.PermissionsMigrated

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Name struct {
	// Family name of the Databricks user.
	// Wire name: 'familyName'
	FamilyName string `json:"familyName,omitempty"`
	// Given name of the Databricks user.
	// Wire name: 'givenName'
	GivenName       string   `json:"givenName,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Name) MarshalJSON() ([]byte, error) {
	pb, err := NameToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Name) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.NamePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NameFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NameToPb(st *Name) (*iampb.NamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.NamePb{}
	pb.FamilyName = st.FamilyName
	pb.GivenName = st.GivenName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NameFromPb(pb *iampb.NamePb) (*Name, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Name{}
	st.FamilyName = pb.FamilyName
	st.GivenName = pb.GivenName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType      string   `json:"object_type,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := ObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ObjectPermissionsToPb(st *ObjectPermissions) (*iampb.ObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ObjectPermissionsPb{}

	var accessControlListPb []iampb.AccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := AccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ObjectPermissionsFromPb(pb *iampb.ObjectPermissionsPb) (*ObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ObjectPermissions{}

	var accessControlListField []AccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := AccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PartialUpdate struct {
	// Unique ID in the Databricks workspace.
	Id string `json:"-" tf:"-"`

	// Wire name: 'Operations'
	Operations []Patch `json:"Operations,omitempty"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	// Wire name: 'schemas'
	Schemas []PatchSchema `json:"schemas,omitempty"`
}

func (st PartialUpdate) MarshalJSON() ([]byte, error) {
	pb, err := PartialUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PartialUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PartialUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PartialUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PartialUpdateToPb(st *PartialUpdate) (*iampb.PartialUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PartialUpdatePb{}
	pb.Id = st.Id

	var operationsPb []iampb.PatchPb
	for _, item := range st.Operations {
		itemPb, err := PatchToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			operationsPb = append(operationsPb, *itemPb)
		}
	}
	pb.Operations = operationsPb

	var schemasPb []iampb.PatchSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := PatchSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb

	return pb, nil
}

func PartialUpdateFromPb(pb *iampb.PartialUpdatePb) (*PartialUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartialUpdate{}
	st.Id = pb.Id

	var operationsField []Patch
	for _, itemPb := range pb.Operations {
		item, err := PatchFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			operationsField = append(operationsField, *item)
		}
	}
	st.Operations = operationsField

	var schemasField []PatchSchema
	for _, itemPb := range pb.Schemas {
		item, err := PatchSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField

	return st, nil
}

type PasswordAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PasswordAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := PasswordAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PasswordAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PasswordAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PasswordAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PasswordAccessControlRequestToPb(st *PasswordAccessControlRequest) (*iampb.PasswordAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PasswordAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := PasswordPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PasswordAccessControlRequestFromPb(pb *iampb.PasswordAccessControlRequestPb) (*PasswordAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := PasswordPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PasswordAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []PasswordPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PasswordAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := PasswordAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PasswordAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PasswordAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PasswordAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PasswordAccessControlResponseToPb(st *PasswordAccessControlResponse) (*iampb.PasswordAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PasswordAccessControlResponsePb{}

	var allPermissionsPb []iampb.PasswordPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := PasswordPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PasswordAccessControlResponseFromPb(pb *iampb.PasswordAccessControlResponsePb) (*PasswordAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordAccessControlResponse{}

	var allPermissionsField []PasswordPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := PasswordPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PasswordPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string                `json:"-" tf:"-"`
}

func (st PasswordPermission) MarshalJSON() ([]byte, error) {
	pb, err := PasswordPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PasswordPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PasswordPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PasswordPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PasswordPermissionToPb(st *PasswordPermission) (*iampb.PasswordPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PasswordPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := PasswordPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PasswordPermissionFromPb(pb *iampb.PasswordPermissionPb) (*PasswordPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := PasswordPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func PasswordPermissionLevelToPb(st *PasswordPermissionLevel) (*iampb.PasswordPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.PasswordPermissionLevelPb(*st)
	return &pb, nil
}

func PasswordPermissionLevelFromPb(pb *iampb.PasswordPermissionLevelPb) (*PasswordPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PasswordPermissionLevel(*pb)
	return &st, nil
}

type PasswordPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []PasswordAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType      string   `json:"object_type,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PasswordPermissions) MarshalJSON() ([]byte, error) {
	pb, err := PasswordPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PasswordPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PasswordPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PasswordPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PasswordPermissionsToPb(st *PasswordPermissions) (*iampb.PasswordPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PasswordPermissionsPb{}

	var accessControlListPb []iampb.PasswordAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := PasswordAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PasswordPermissionsFromPb(pb *iampb.PasswordPermissionsPb) (*PasswordPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissions{}

	var accessControlListField []PasswordAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := PasswordAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PasswordPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string                `json:"-" tf:"-"`
}

func (st PasswordPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := PasswordPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PasswordPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PasswordPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PasswordPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PasswordPermissionsDescriptionToPb(st *PasswordPermissionsDescription) (*iampb.PasswordPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PasswordPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := PasswordPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PasswordPermissionsDescriptionFromPb(pb *iampb.PasswordPermissionsDescriptionPb) (*PasswordPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := PasswordPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PasswordPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []PasswordAccessControlRequest `json:"access_control_list,omitempty"`
}

func (st PasswordPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := PasswordPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PasswordPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PasswordPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PasswordPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PasswordPermissionsRequestToPb(st *PasswordPermissionsRequest) (*iampb.PasswordPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PasswordPermissionsRequestPb{}

	var accessControlListPb []iampb.PasswordAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := PasswordAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	return pb, nil
}

func PasswordPermissionsRequestFromPb(pb *iampb.PasswordPermissionsRequestPb) (*PasswordPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissionsRequest{}

	var accessControlListField []PasswordAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := PasswordAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField

	return st, nil
}

type Patch struct {
	// Type of patch operation.
	// Wire name: 'op'
	Op PatchOp `json:"op,omitempty"`
	// Selection of patch operation
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// Value to modify
	// Wire name: 'value'
	Value           any      `json:"value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Patch) MarshalJSON() ([]byte, error) {
	pb, err := PatchToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Patch) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PatchPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PatchFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PatchToPb(st *Patch) (*iampb.PatchPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PatchPb{}
	opPb, err := PatchOpToPb(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}
	pb.Path = st.Path
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PatchFromPb(pb *iampb.PatchPb) (*Patch, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Patch{}
	opField, err := PatchOpFromPb(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	st.Path = pb.Path
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func PatchOpToPb(st *PatchOp) (*iampb.PatchOpPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.PatchOpPb(*st)
	return &pb, nil
}

func PatchOpFromPb(pb *iampb.PatchOpPb) (*PatchOp, error) {
	if pb == nil {
		return nil, nil
	}
	st := PatchOp(*pb)
	return &st, nil
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

func PatchSchemaToPb(st *PatchSchema) (*iampb.PatchSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.PatchSchemaPb(*st)
	return &pb, nil
}

func PatchSchemaFromPb(pb *iampb.PatchSchemaPb) (*PatchSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := PatchSchema(*pb)
	return &st, nil
}

type Permission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string        `json:"-" tf:"-"`
}

func (st Permission) MarshalJSON() ([]byte, error) {
	pb, err := PermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Permission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermissionToPb(st *Permission) (*iampb.PermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := PermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PermissionFromPb(pb *iampb.PermissionPb) (*Permission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Permission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := PermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The output format for existing workspace PermissionAssignment records, which
// contains some info for user consumption.
type PermissionAssignment struct {
	// Error response associated with a workspace permission assignment, if any.
	// Wire name: 'error'
	Error string `json:"error,omitempty"`
	// The permissions level of the principal.
	// Wire name: 'permissions'
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// Information about the principal assigned to the workspace.
	// Wire name: 'principal'
	Principal       *PrincipalOutput `json:"principal,omitempty"`
	ForceSendFields []string         `json:"-" tf:"-"`
}

func (st PermissionAssignment) MarshalJSON() ([]byte, error) {
	pb, err := PermissionAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PermissionAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PermissionAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermissionAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermissionAssignmentToPb(st *PermissionAssignment) (*iampb.PermissionAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PermissionAssignmentPb{}
	pb.Error = st.Error

	var permissionsPb []iampb.WorkspacePermissionPb
	for _, item := range st.Permissions {
		itemPb, err := WorkspacePermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionsPb = append(permissionsPb, *itemPb)
		}
	}
	pb.Permissions = permissionsPb
	principalPb, err := PrincipalOutputToPb(st.Principal)
	if err != nil {
		return nil, err
	}
	if principalPb != nil {
		pb.Principal = principalPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PermissionAssignmentFromPb(pb *iampb.PermissionAssignmentPb) (*PermissionAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionAssignment{}
	st.Error = pb.Error

	var permissionsField []WorkspacePermission
	for _, itemPb := range pb.Permissions {
		item, err := WorkspacePermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionsField = append(permissionsField, *item)
		}
	}
	st.Permissions = permissionsField
	principalField, err := PrincipalOutputFromPb(pb.Principal)
	if err != nil {
		return nil, err
	}
	if principalField != nil {
		st.Principal = principalField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	// Wire name: 'permission_assignments'
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`
}

func (st PermissionAssignments) MarshalJSON() ([]byte, error) {
	pb, err := PermissionAssignmentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PermissionAssignments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PermissionAssignmentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermissionAssignmentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermissionAssignmentsToPb(st *PermissionAssignments) (*iampb.PermissionAssignmentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PermissionAssignmentsPb{}

	var permissionAssignmentsPb []iampb.PermissionAssignmentPb
	for _, item := range st.PermissionAssignments {
		itemPb, err := PermissionAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionAssignmentsPb = append(permissionAssignmentsPb, *itemPb)
		}
	}
	pb.PermissionAssignments = permissionAssignmentsPb

	return pb, nil
}

func PermissionAssignmentsFromPb(pb *iampb.PermissionAssignmentsPb) (*PermissionAssignments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionAssignments{}

	var permissionAssignmentsField []PermissionAssignment
	for _, itemPb := range pb.PermissionAssignments {
		item, err := PermissionAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionAssignmentsField = append(permissionAssignmentsField, *item)
		}
	}
	st.PermissionAssignments = permissionAssignmentsField

	return st, nil
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

func PermissionLevelToPb(st *PermissionLevel) (*iampb.PermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.PermissionLevelPb(*st)
	return &pb, nil
}

func PermissionLevelFromPb(pb *iampb.PermissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
}

type PermissionOutput struct {
	// The results of a permissions query.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WorkspacePermission `json:"permission_level,omitempty"`
	ForceSendFields []string            `json:"-" tf:"-"`
}

func (st PermissionOutput) MarshalJSON() ([]byte, error) {
	pb, err := PermissionOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PermissionOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PermissionOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermissionOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermissionOutputToPb(st *PermissionOutput) (*iampb.PermissionOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PermissionOutputPb{}
	pb.Description = st.Description
	permissionLevelPb, err := WorkspacePermissionToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PermissionOutputFromPb(pb *iampb.PermissionOutputPb) (*PermissionOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionOutput{}
	st.Description = pb.Description
	permissionLevelField, err := WorkspacePermissionFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string        `json:"-" tf:"-"`
}

func (st PermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := PermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermissionsDescriptionToPb(st *PermissionsDescription) (*iampb.PermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := PermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PermissionsDescriptionFromPb(pb *iampb.PermissionsDescriptionPb) (*PermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := PermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Information about the principal assigned to the workspace.
type PrincipalOutput struct {
	// The display name of the principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// The group name of the group. Present only if the principal is a group.
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// The unique, opaque id of the principal.
	// Wire name: 'principal_id'
	PrincipalId int64 `json:"principal_id,omitempty"`
	// The name of the service principal. Present only if the principal is a
	// service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The username of the user. Present only if the principal is a user.
	// Wire name: 'user_name'
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PrincipalOutput) MarshalJSON() ([]byte, error) {
	pb, err := PrincipalOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PrincipalOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.PrincipalOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PrincipalOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PrincipalOutputToPb(st *PrincipalOutput) (*iampb.PrincipalOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.PrincipalOutputPb{}
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.PrincipalId = st.PrincipalId
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PrincipalOutputFromPb(pb *iampb.PrincipalOutputPb) (*PrincipalOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrincipalOutput{}
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.PrincipalId = pb.PrincipalId
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func RequestAuthzIdentityToPb(st *RequestAuthzIdentity) (*iampb.RequestAuthzIdentityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.RequestAuthzIdentityPb(*st)
	return &pb, nil
}

func RequestAuthzIdentityFromPb(pb *iampb.RequestAuthzIdentityPb) (*RequestAuthzIdentity, error) {
	if pb == nil {
		return nil, nil
	}
	st := RequestAuthzIdentity(*pb)
	return &st, nil
}

type ResourceInfo struct {
	// Id of the current resource.
	// Wire name: 'id'
	Id string `json:"id"`
	// The legacy acl path of the current resource.
	// Wire name: 'legacy_acl_path'
	LegacyAclPath string `json:"legacy_acl_path,omitempty"`
	// Parent resource info for the current resource. The parent may have
	// another parent.
	// Wire name: 'parent_resource_info'
	ParentResourceInfo *ResourceInfo `json:"parent_resource_info,omitempty"`
	ForceSendFields    []string      `json:"-" tf:"-"`
}

func (st ResourceInfo) MarshalJSON() ([]byte, error) {
	pb, err := ResourceInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResourceInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ResourceInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResourceInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResourceInfoToPb(st *ResourceInfo) (*iampb.ResourceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ResourceInfoPb{}
	pb.Id = st.Id
	pb.LegacyAclPath = st.LegacyAclPath
	parentResourceInfoPb, err := ResourceInfoToPb(st.ParentResourceInfo)
	if err != nil {
		return nil, err
	}
	if parentResourceInfoPb != nil {
		pb.ParentResourceInfo = parentResourceInfoPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ResourceInfoFromPb(pb *iampb.ResourceInfoPb) (*ResourceInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResourceInfo{}
	st.Id = pb.Id
	st.LegacyAclPath = pb.LegacyAclPath
	parentResourceInfoField, err := ResourceInfoFromPb(pb.ParentResourceInfo)
	if err != nil {
		return nil, err
	}
	if parentResourceInfoField != nil {
		st.ParentResourceInfo = parentResourceInfoField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	// Wire name: 'resourceType'
	ResourceType    string   `json:"resourceType,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ResourceMeta) MarshalJSON() ([]byte, error) {
	pb, err := ResourceMetaToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResourceMeta) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ResourceMetaPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResourceMetaFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResourceMetaToPb(st *ResourceMeta) (*iampb.ResourceMetaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ResourceMetaPb{}
	pb.ResourceType = st.ResourceType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ResourceMetaFromPb(pb *iampb.ResourceMetaPb) (*ResourceMeta, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResourceMeta{}
	st.ResourceType = pb.ResourceType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	// Wire name: 'name'
	Name string `json:"name"`
}

func (st Role) MarshalJSON() ([]byte, error) {
	pb, err := RoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Role) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.RolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RoleToPb(st *Role) (*iampb.RolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.RolePb{}
	pb.Name = st.Name

	return pb, nil
}

func RoleFromPb(pb *iampb.RolePb) (*Role, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Role{}
	st.Name = pb.Name

	return st, nil
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
	Etag string `json:"etag"`

	// Wire name: 'grant_rules'
	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	// Wire name: 'name'
	Name string `json:"name"`
}

func (st RuleSetResponse) MarshalJSON() ([]byte, error) {
	pb, err := RuleSetResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RuleSetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.RuleSetResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RuleSetResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RuleSetResponseToPb(st *RuleSetResponse) (*iampb.RuleSetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.RuleSetResponsePb{}
	pb.Etag = st.Etag

	var grantRulesPb []iampb.GrantRulePb
	for _, item := range st.GrantRules {
		itemPb, err := GrantRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			grantRulesPb = append(grantRulesPb, *itemPb)
		}
	}
	pb.GrantRules = grantRulesPb
	pb.Name = st.Name

	return pb, nil
}

func RuleSetResponseFromPb(pb *iampb.RuleSetResponsePb) (*RuleSetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RuleSetResponse{}
	st.Etag = pb.Etag

	var grantRulesField []GrantRule
	for _, itemPb := range pb.GrantRules {
		item, err := GrantRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			grantRulesField = append(grantRulesField, *item)
		}
	}
	st.GrantRules = grantRulesField
	st.Name = pb.Name

	return st, nil
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
	Etag string `json:"etag"`

	// Wire name: 'grant_rules'
	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	// Name of the rule set.
	// Wire name: 'name'
	Name string `json:"name"`
}

func (st RuleSetUpdateRequest) MarshalJSON() ([]byte, error) {
	pb, err := RuleSetUpdateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RuleSetUpdateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.RuleSetUpdateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RuleSetUpdateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RuleSetUpdateRequestToPb(st *RuleSetUpdateRequest) (*iampb.RuleSetUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.RuleSetUpdateRequestPb{}
	pb.Etag = st.Etag

	var grantRulesPb []iampb.GrantRulePb
	for _, item := range st.GrantRules {
		itemPb, err := GrantRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			grantRulesPb = append(grantRulesPb, *itemPb)
		}
	}
	pb.GrantRules = grantRulesPb
	pb.Name = st.Name

	return pb, nil
}

func RuleSetUpdateRequestFromPb(pb *iampb.RuleSetUpdateRequestPb) (*RuleSetUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RuleSetUpdateRequest{}
	st.Etag = pb.Etag

	var grantRulesField []GrantRule
	for _, itemPb := range pb.GrantRules {
		item, err := GrantRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			grantRulesField = append(grantRulesField, *item)
		}
	}
	st.GrantRules = grantRulesField
	st.Name = pb.Name

	return st, nil
}

type ServicePrincipal struct {
	// If this user is active
	// Wire name: 'active'
	Active bool `json:"active,omitempty"`
	// UUID relating to the service principal
	// Wire name: 'applicationId'
	ApplicationId string `json:"applicationId,omitempty"`
	// String that represents a concatenation of given and family names.
	// Wire name: 'displayName'
	DisplayName string `json:"displayName,omitempty"`
	// Entitlements assigned to the service principal. See [assigning
	// entitlements] for a full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	// Wire name: 'entitlements'
	Entitlements []ComplexValue `json:"entitlements,omitempty"`

	// Wire name: 'externalId'
	ExternalId string `json:"externalId,omitempty"`

	// Wire name: 'groups'
	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks service principal ID.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	// Wire name: 'roles'
	Roles []ComplexValue `json:"roles,omitempty"`
	// The schema of the List response.
	// Wire name: 'schemas'
	Schemas         []ServicePrincipalSchema `json:"schemas,omitempty"`
	ForceSendFields []string                 `json:"-" tf:"-"`
}

func (st ServicePrincipal) MarshalJSON() ([]byte, error) {
	pb, err := ServicePrincipalToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ServicePrincipal) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.ServicePrincipalPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ServicePrincipalFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ServicePrincipalToPb(st *ServicePrincipal) (*iampb.ServicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.ServicePrincipalPb{}
	pb.Active = st.Active
	pb.ApplicationId = st.ApplicationId
	pb.DisplayName = st.DisplayName

	var entitlementsPb []iampb.ComplexValuePb
	for _, item := range st.Entitlements {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entitlementsPb = append(entitlementsPb, *itemPb)
		}
	}
	pb.Entitlements = entitlementsPb
	pb.ExternalId = st.ExternalId

	var groupsPb []iampb.ComplexValuePb
	for _, item := range st.Groups {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			groupsPb = append(groupsPb, *itemPb)
		}
	}
	pb.Groups = groupsPb
	pb.Id = st.Id

	var rolesPb []iampb.ComplexValuePb
	for _, item := range st.Roles {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	var schemasPb []iampb.ServicePrincipalSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := ServicePrincipalSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ServicePrincipalFromPb(pb *iampb.ServicePrincipalPb) (*ServicePrincipal, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServicePrincipal{}
	st.Active = pb.Active
	st.ApplicationId = pb.ApplicationId
	st.DisplayName = pb.DisplayName

	var entitlementsField []ComplexValue
	for _, itemPb := range pb.Entitlements {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			entitlementsField = append(entitlementsField, *item)
		}
	}
	st.Entitlements = entitlementsField
	st.ExternalId = pb.ExternalId

	var groupsField []ComplexValue
	for _, itemPb := range pb.Groups {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			groupsField = append(groupsField, *item)
		}
	}
	st.Groups = groupsField
	st.Id = pb.Id

	var rolesField []ComplexValue
	for _, itemPb := range pb.Roles {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField

	var schemasField []ServicePrincipalSchema
	for _, itemPb := range pb.Schemas {
		item, err := ServicePrincipalSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ServicePrincipalSchemaToPb(st *ServicePrincipalSchema) (*iampb.ServicePrincipalSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.ServicePrincipalSchemaPb(*st)
	return &pb, nil
}

func ServicePrincipalSchemaFromPb(pb *iampb.ServicePrincipalSchemaPb) (*ServicePrincipalSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServicePrincipalSchema(*pb)
	return &st, nil
}

type SetObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`
	// The id of the request object.
	RequestObjectId string `json:"-" tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" tf:"-"`
}

func (st SetObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := SetObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SetObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.SetObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SetObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SetObjectPermissionsToPb(st *SetObjectPermissions) (*iampb.SetObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.SetObjectPermissionsPb{}

	var accessControlListPb []iampb.AccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := AccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.RequestObjectId = st.RequestObjectId
	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
}

func SetObjectPermissionsFromPb(pb *iampb.SetObjectPermissionsPb) (*SetObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetObjectPermissions{}

	var accessControlListField []AccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := AccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
}

type UpdateObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`
	// The id of the request object.
	RequestObjectId string `json:"-" tf:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// alertsv2, authorization, clusters, cluster-policies, dashboards,
	// dbsql-dashboards, directories, experiments, files, instance-pools, jobs,
	// notebooks, pipelines, queries, registered-models, repos,
	// serving-endpoints, or warehouses.
	RequestObjectType string `json:"-" tf:"-"`
}

func (st UpdateObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := UpdateObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.UpdateObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateObjectPermissionsToPb(st *UpdateObjectPermissions) (*iampb.UpdateObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.UpdateObjectPermissionsPb{}

	var accessControlListPb []iampb.AccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := AccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.RequestObjectId = st.RequestObjectId
	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
}

func UpdateObjectPermissionsFromPb(pb *iampb.UpdateObjectPermissionsPb) (*UpdateObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateObjectPermissions{}

	var accessControlListField []AccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := AccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	// Wire name: 'name'
	Name string `json:"name"`

	// Wire name: 'rule_set'
	RuleSet RuleSetUpdateRequest `json:"rule_set"`
}

func (st UpdateRuleSetRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRuleSetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRuleSetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.UpdateRuleSetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRuleSetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateRuleSetRequestToPb(st *UpdateRuleSetRequest) (*iampb.UpdateRuleSetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.UpdateRuleSetRequestPb{}
	pb.Name = st.Name
	ruleSetPb, err := RuleSetUpdateRequestToPb(&st.RuleSet)
	if err != nil {
		return nil, err
	}
	if ruleSetPb != nil {
		pb.RuleSet = *ruleSetPb
	}

	return pb, nil
}

func UpdateRuleSetRequestFromPb(pb *iampb.UpdateRuleSetRequestPb) (*UpdateRuleSetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRuleSetRequest{}
	st.Name = pb.Name
	ruleSetField, err := RuleSetUpdateRequestFromPb(&pb.RuleSet)
	if err != nil {
		return nil, err
	}
	if ruleSetField != nil {
		st.RuleSet = *ruleSetField
	}

	return st, nil
}

type UpdateWorkspaceAssignments struct {
	// Array of permissions assignments to update on the workspace. Valid values
	// are "USER" and "ADMIN" (case-sensitive). If both "USER" and "ADMIN" are
	// provided, "ADMIN" takes precedence. Other values will be ignored. Note
	// that excluding this field, or providing unsupported values, will have the
	// same effect as providing an empty list, which will result in the deletion
	// of all permissions for the principal.
	// Wire name: 'permissions'
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" tf:"-"`
	// The workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st UpdateWorkspaceAssignments) MarshalJSON() ([]byte, error) {
	pb, err := UpdateWorkspaceAssignmentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateWorkspaceAssignments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.UpdateWorkspaceAssignmentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateWorkspaceAssignmentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateWorkspaceAssignmentsToPb(st *UpdateWorkspaceAssignments) (*iampb.UpdateWorkspaceAssignmentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.UpdateWorkspaceAssignmentsPb{}

	var permissionsPb []iampb.WorkspacePermissionPb
	for _, item := range st.Permissions {
		itemPb, err := WorkspacePermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionsPb = append(permissionsPb, *itemPb)
		}
	}
	pb.Permissions = permissionsPb
	pb.PrincipalId = st.PrincipalId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func UpdateWorkspaceAssignmentsFromPb(pb *iampb.UpdateWorkspaceAssignmentsPb) (*UpdateWorkspaceAssignments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceAssignments{}

	var permissionsField []WorkspacePermission
	for _, itemPb := range pb.Permissions {
		item, err := WorkspacePermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionsField = append(permissionsField, *item)
		}
	}
	st.Permissions = permissionsField
	st.PrincipalId = pb.PrincipalId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type User struct {
	// If this user is active
	// Wire name: 'active'
	Active bool `json:"active,omitempty"`
	// String that represents a concatenation of given and family names. For
	// example `John Smith`. This field cannot be updated through the Workspace
	// SCIM APIs when [identity federation is enabled]. Use Account SCIM APIs to
	// update `displayName`.
	//
	// [identity federation is enabled]: https://docs.databricks.com/administration-guide/users-groups/best-practices.html#enable-identity-federation
	// Wire name: 'displayName'
	DisplayName string `json:"displayName,omitempty"`
	// All the emails associated with the Databricks user.
	// Wire name: 'emails'
	Emails []ComplexValue `json:"emails,omitempty"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	// Wire name: 'entitlements'
	Entitlements []ComplexValue `json:"entitlements,omitempty"`
	// External ID is not currently supported. It is reserved for future use.
	// Wire name: 'externalId'
	ExternalId string `json:"externalId,omitempty"`

	// Wire name: 'groups'
	Groups []ComplexValue `json:"groups,omitempty"`
	// Databricks user ID.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'name'
	Name *Name `json:"name,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	// Wire name: 'roles'
	Roles []ComplexValue `json:"roles,omitempty"`
	// The schema of the user.
	// Wire name: 'schemas'
	Schemas []UserSchema `json:"schemas,omitempty"`
	// Email address of the Databricks user.
	// Wire name: 'userName'
	UserName        string   `json:"userName,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st User) MarshalJSON() ([]byte, error) {
	pb, err := UserToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *User) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.UserPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UserFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UserToPb(st *User) (*iampb.UserPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.UserPb{}
	pb.Active = st.Active
	pb.DisplayName = st.DisplayName

	var emailsPb []iampb.ComplexValuePb
	for _, item := range st.Emails {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			emailsPb = append(emailsPb, *itemPb)
		}
	}
	pb.Emails = emailsPb

	var entitlementsPb []iampb.ComplexValuePb
	for _, item := range st.Entitlements {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entitlementsPb = append(entitlementsPb, *itemPb)
		}
	}
	pb.Entitlements = entitlementsPb
	pb.ExternalId = st.ExternalId

	var groupsPb []iampb.ComplexValuePb
	for _, item := range st.Groups {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			groupsPb = append(groupsPb, *itemPb)
		}
	}
	pb.Groups = groupsPb
	pb.Id = st.Id
	namePb, err := NameToPb(st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = namePb
	}

	var rolesPb []iampb.ComplexValuePb
	for _, item := range st.Roles {
		itemPb, err := ComplexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	var schemasPb []iampb.UserSchemaPb
	for _, item := range st.Schemas {
		itemPb, err := UserSchemaToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemasPb = append(schemasPb, *itemPb)
		}
	}
	pb.Schemas = schemasPb
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UserFromPb(pb *iampb.UserPb) (*User, error) {
	if pb == nil {
		return nil, nil
	}
	st := &User{}
	st.Active = pb.Active
	st.DisplayName = pb.DisplayName

	var emailsField []ComplexValue
	for _, itemPb := range pb.Emails {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			emailsField = append(emailsField, *item)
		}
	}
	st.Emails = emailsField

	var entitlementsField []ComplexValue
	for _, itemPb := range pb.Entitlements {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			entitlementsField = append(entitlementsField, *item)
		}
	}
	st.Entitlements = entitlementsField
	st.ExternalId = pb.ExternalId

	var groupsField []ComplexValue
	for _, itemPb := range pb.Groups {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			groupsField = append(groupsField, *item)
		}
	}
	st.Groups = groupsField
	st.Id = pb.Id
	nameField, err := NameFromPb(pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = nameField
	}

	var rolesField []ComplexValue
	for _, itemPb := range pb.Roles {
		item, err := ComplexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField

	var schemasField []UserSchema
	for _, itemPb := range pb.Schemas {
		item, err := UserSchemaFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemasField = append(schemasField, *item)
		}
	}
	st.Schemas = schemasField
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func UserSchemaToPb(st *UserSchema) (*iampb.UserSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.UserSchemaPb(*st)
	return &pb, nil
}

func UserSchemaFromPb(pb *iampb.UserSchemaPb) (*UserSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := UserSchema(*pb)
	return &st, nil
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

func WorkspacePermissionToPb(st *WorkspacePermission) (*iampb.WorkspacePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := iampb.WorkspacePermissionPb(*st)
	return &pb, nil
}

func WorkspacePermissionFromPb(pb *iampb.WorkspacePermissionPb) (*WorkspacePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspacePermission(*pb)
	return &st, nil
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	// Wire name: 'permissions'
	Permissions []PermissionOutput `json:"permissions,omitempty"`
}

func (st WorkspacePermissions) MarshalJSON() ([]byte, error) {
	pb, err := WorkspacePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspacePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &iampb.WorkspacePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspacePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WorkspacePermissionsToPb(st *WorkspacePermissions) (*iampb.WorkspacePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &iampb.WorkspacePermissionsPb{}

	var permissionsPb []iampb.PermissionOutputPb
	for _, item := range st.Permissions {
		itemPb, err := PermissionOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionsPb = append(permissionsPb, *itemPb)
		}
	}
	pb.Permissions = permissionsPb

	return pb, nil
}

func WorkspacePermissionsFromPb(pb *iampb.WorkspacePermissionsPb) (*WorkspacePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspacePermissions{}

	var permissionsField []PermissionOutput
	for _, itemPb := range pb.Permissions {
		item, err := PermissionOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionsField = append(permissionsField, *item)
		}
	}
	st.Permissions = permissionsField

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
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
