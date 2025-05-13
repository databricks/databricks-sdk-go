// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
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

func accessControlRequestToPb(st *AccessControlRequest) (*accessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type accessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func accessControlRequestFromPb(pb *accessControlRequestPb) (*AccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *accessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st accessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func accessControlResponseToPb(st *AccessControlResponse) (*accessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accessControlResponsePb{}

	var allPermissionsPb []permissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := permissionToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type accessControlResponsePb struct {
	// All permissions.
	AllPermissions []permissionPb `json:"all_permissions,omitempty"`
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

func accessControlResponseFromPb(pb *accessControlResponsePb) (*AccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControlResponse{}

	var allPermissionsField []Permission
	for _, itemPb := range pb.AllPermissions {
		item, err := permissionFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *accessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st accessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// represents an identity trying to access a resource - user or a service
// principal group can be a principal of a permission set assignment but an
// actor is always a user or a service principal
type Actor struct {

	// Wire name: 'actor_id'
	ActorId int64

	ForceSendFields []string `tf:"-"`
}

func actorToPb(st *Actor) (*actorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &actorPb{}
	pb.ActorId = st.ActorId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type actorPb struct {
	ActorId int64 `json:"actor_id,omitempty" url:"actor_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func actorFromPb(pb *actorPb) (*Actor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Actor{}
	st.ActorId = pb.ActorId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *actorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st actorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func checkPolicyRequestToPb(st *CheckPolicyRequest) (*checkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &checkPolicyRequestPb{}
	actorPb, err := actorToPb(&st.Actor)
	if err != nil {
		return nil, err
	}
	if actorPb != nil {
		pb.Actor = *actorPb
	}

	pb.AuthzIdentity = st.AuthzIdentity

	consistencyTokenPb, err := consistencyTokenToPb(&st.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenPb != nil {
		pb.ConsistencyToken = *consistencyTokenPb
	}

	pb.Permission = st.Permission

	pb.Resource = st.Resource

	resourceInfoPb, err := resourceInfoToPb(st.ResourceInfo)
	if err != nil {
		return nil, err
	}
	if resourceInfoPb != nil {
		pb.ResourceInfo = resourceInfoPb
	}

	return pb, nil
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

type checkPolicyRequestPb struct {
	Actor actorPb `json:"-" url:"actor"`

	AuthzIdentity RequestAuthzIdentity `json:"-" url:"authz_identity"`

	ConsistencyToken consistencyTokenPb `json:"-" url:"consistency_token"`

	Permission string `json:"-" url:"permission"`
	// Ex: (servicePrincipal/use,
	// accounts/<account-id>/servicePrincipals/<sp-id>) Ex:
	// (servicePrincipal.ruleSet/update,
	// accounts/<account-id>/servicePrincipals/<sp-id>/ruleSets/default)
	Resource string `json:"-" url:"resource"`

	ResourceInfo *resourceInfoPb `json:"-" url:"resource_info,omitempty"`
}

func checkPolicyRequestFromPb(pb *checkPolicyRequestPb) (*CheckPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CheckPolicyRequest{}
	actorField, err := actorFromPb(&pb.Actor)
	if err != nil {
		return nil, err
	}
	if actorField != nil {
		st.Actor = *actorField
	}
	st.AuthzIdentity = pb.AuthzIdentity
	consistencyTokenField, err := consistencyTokenFromPb(&pb.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenField != nil {
		st.ConsistencyToken = *consistencyTokenField
	}
	st.Permission = pb.Permission
	st.Resource = pb.Resource
	resourceInfoField, err := resourceInfoFromPb(pb.ResourceInfo)
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
	ConsistencyToken ConsistencyToken

	// Wire name: 'is_permitted'
	IsPermitted bool

	ForceSendFields []string `tf:"-"`
}

func checkPolicyResponseToPb(st *CheckPolicyResponse) (*checkPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &checkPolicyResponsePb{}
	consistencyTokenPb, err := consistencyTokenToPb(&st.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenPb != nil {
		pb.ConsistencyToken = *consistencyTokenPb
	}

	pb.IsPermitted = st.IsPermitted

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type checkPolicyResponsePb struct {
	ConsistencyToken consistencyTokenPb `json:"consistency_token"`

	IsPermitted bool `json:"is_permitted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func checkPolicyResponseFromPb(pb *checkPolicyResponsePb) (*CheckPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CheckPolicyResponse{}
	consistencyTokenField, err := consistencyTokenFromPb(&pb.ConsistencyToken)
	if err != nil {
		return nil, err
	}
	if consistencyTokenField != nil {
		st.ConsistencyToken = *consistencyTokenField
	}
	st.IsPermitted = pb.IsPermitted

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *checkPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st checkPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func complexValueToPb(st *ComplexValue) (*complexValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complexValuePb{}
	pb.Display = st.Display

	pb.Primary = st.Primary

	pb.Ref = st.Ref

	pb.Type = st.Type

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type complexValuePb struct {
	Display string `json:"display,omitempty"`

	Primary bool `json:"primary,omitempty"`

	Ref string `json:"$ref,omitempty"`

	Type string `json:"type,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func complexValueFromPb(pb *complexValuePb) (*ComplexValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplexValue{}
	st.Display = pb.Display
	st.Primary = pb.Primary
	st.Ref = pb.Ref
	st.Type = pb.Type
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *complexValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st complexValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ConsistencyToken struct {

	// Wire name: 'value'
	Value string
}

func consistencyTokenToPb(st *ConsistencyToken) (*consistencyTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &consistencyTokenPb{}
	pb.Value = st.Value

	return pb, nil
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

type consistencyTokenPb struct {
	Value string `json:"value"`
}

func consistencyTokenFromPb(pb *consistencyTokenPb) (*ConsistencyToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConsistencyToken{}
	st.Value = pb.Value

	return st, nil
}

// Delete a group.
type DeleteAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteAccountGroupRequestToPb(st *DeleteAccountGroupRequest) (*deleteAccountGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteAccountGroupRequestPb struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`
}

func deleteAccountGroupRequestFromPb(pb *deleteAccountGroupRequestPb) (*DeleteAccountGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete a service principal.
type DeleteAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteAccountServicePrincipalRequestToPb(st *DeleteAccountServicePrincipalRequest) (*deleteAccountServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteAccountServicePrincipalRequestPb struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`
}

func deleteAccountServicePrincipalRequestFromPb(pb *deleteAccountServicePrincipalRequestPb) (*DeleteAccountServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete a user.
type DeleteAccountUserRequest struct {
	// Unique ID for a user in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteAccountUserRequestToPb(st *DeleteAccountUserRequest) (*deleteAccountUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountUserRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteAccountUserRequestPb struct {
	// Unique ID for a user in the Databricks account.
	Id string `json:"-" url:"-"`
}

func deleteAccountUserRequestFromPb(pb *deleteAccountUserRequestPb) (*DeleteAccountUserRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountUserRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete a group.
type DeleteGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteGroupRequestToPb(st *DeleteGroupRequest) (*deleteGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteGroupRequestPb struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

func deleteGroupRequestFromPb(pb *deleteGroupRequestPb) (*DeleteGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

// Delete a service principal.
type DeleteServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteServicePrincipalRequestToPb(st *DeleteServicePrincipalRequest) (*deleteServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteServicePrincipalRequestPb struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

func deleteServicePrincipalRequestFromPb(pb *deleteServicePrincipalRequestPb) (*DeleteServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete a user.
type DeleteUserRequest struct {
	// Unique ID for a user in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteUserRequestToPb(st *DeleteUserRequest) (*deleteUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteUserRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteUserRequestPb struct {
	// Unique ID for a user in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

func deleteUserRequestFromPb(pb *deleteUserRequestPb) (*DeleteUserRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteUserRequest{}
	st.Id = pb.Id

	return st, nil
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

func deleteWorkspaceAssignmentRequestToPb(st *DeleteWorkspaceAssignmentRequest) (*deleteWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspaceAssignmentRequestPb{}
	pb.PrincipalId = st.PrincipalId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
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

type deleteWorkspaceAssignmentRequestPb struct {
	// The ID of the user, service principal, or group.
	PrincipalId int64 `json:"-" url:"-"`
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`
}

func deleteWorkspaceAssignmentRequestFromPb(pb *deleteWorkspaceAssignmentRequestPb) (*DeleteWorkspaceAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWorkspaceAssignmentRequest{}
	st.PrincipalId = pb.PrincipalId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type DeleteWorkspacePermissionAssignmentResponse struct {
}

func deleteWorkspacePermissionAssignmentResponseToPb(st *DeleteWorkspacePermissionAssignmentResponse) (*deleteWorkspacePermissionAssignmentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspacePermissionAssignmentResponsePb{}

	return pb, nil
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

type deleteWorkspacePermissionAssignmentResponsePb struct {
}

func deleteWorkspacePermissionAssignmentResponseFromPb(pb *deleteWorkspacePermissionAssignmentResponsePb) (*DeleteWorkspacePermissionAssignmentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWorkspacePermissionAssignmentResponse{}

	return st, nil
}

// Get group details.
type GetAccountGroupRequest struct {
	// Unique ID for a group in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func getAccountGroupRequestToPb(st *GetAccountGroupRequest) (*getAccountGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getAccountGroupRequestPb struct {
	// Unique ID for a group in the Databricks account.
	Id string `json:"-" url:"-"`
}

func getAccountGroupRequestFromPb(pb *getAccountGroupRequestPb) (*GetAccountGroupRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountGroupRequest{}
	st.Id = pb.Id

	return st, nil
}

// Get service principal details.
type GetAccountServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks account.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func getAccountServicePrincipalRequestToPb(st *GetAccountServicePrincipalRequest) (*getAccountServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getAccountServicePrincipalRequestPb struct {
	// Unique ID for a service principal in the Databricks account.
	Id string `json:"-" url:"-"`
}

func getAccountServicePrincipalRequestFromPb(pb *getAccountServicePrincipalRequestPb) (*GetAccountServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
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

func getAccountUserRequestToPb(st *GetAccountUserRequest) (*getAccountUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountUserRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.Id = st.Id

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getAccountUserRequestPb struct {
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

func getAccountUserRequestFromPb(pb *getAccountUserRequestPb) (*GetAccountUserRequest, error) {
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
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAccountUserRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAccountUserRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get assignable roles for a resource
type GetAssignableRolesForResourceRequest struct {
	// The resource name for which assignable roles will be listed.
	// Wire name: 'resource'
	Resource string `tf:"-"`
}

func getAssignableRolesForResourceRequestToPb(st *GetAssignableRolesForResourceRequest) (*getAssignableRolesForResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAssignableRolesForResourceRequestPb{}
	pb.Resource = st.Resource

	return pb, nil
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

type getAssignableRolesForResourceRequestPb struct {
	// The resource name for which assignable roles will be listed.
	Resource string `json:"-" url:"resource"`
}

func getAssignableRolesForResourceRequestFromPb(pb *getAssignableRolesForResourceRequestPb) (*GetAssignableRolesForResourceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAssignableRolesForResourceRequest{}
	st.Resource = pb.Resource

	return st, nil
}

type GetAssignableRolesForResourceResponse struct {

	// Wire name: 'roles'
	Roles []Role
}

func getAssignableRolesForResourceResponseToPb(st *GetAssignableRolesForResourceResponse) (*getAssignableRolesForResourceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAssignableRolesForResourceResponsePb{}

	var rolesPb []rolePb
	for _, item := range st.Roles {
		itemPb, err := roleToPb(&item)
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

type getAssignableRolesForResourceResponsePb struct {
	Roles []rolePb `json:"roles,omitempty"`
}

func getAssignableRolesForResourceResponseFromPb(pb *getAssignableRolesForResourceResponsePb) (*GetAssignableRolesForResourceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAssignableRolesForResourceResponse{}

	var rolesField []Role
	for _, itemPb := range pb.Roles {
		item, err := roleFromPb(&itemPb)
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

// Get group details.
type GetGroupRequest struct {
	// Unique ID for a group in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func getGroupRequestToPb(st *GetGroupRequest) (*getGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getGroupRequestPb struct {
	// Unique ID for a group in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

func getGroupRequestFromPb(pb *getGroupRequestPb) (*GetGroupRequest, error) {
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
	PermissionLevels []PasswordPermissionsDescription
}

func getPasswordPermissionLevelsResponseToPb(st *GetPasswordPermissionLevelsResponse) (*getPasswordPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPasswordPermissionLevelsResponsePb{}

	var permissionLevelsPb []passwordPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := passwordPermissionsDescriptionToPb(&item)
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

type getPasswordPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []passwordPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getPasswordPermissionLevelsResponseFromPb(pb *getPasswordPermissionLevelsResponsePb) (*GetPasswordPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPasswordPermissionLevelsResponse{}

	var permissionLevelsField []PasswordPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := passwordPermissionsDescriptionFromPb(&itemPb)
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

// Get object permission levels
type GetPermissionLevelsRequest struct {
	// <needs content>
	// Wire name: 'request_object_id'
	RequestObjectId string `tf:"-"`
	// <needs content>
	// Wire name: 'request_object_type'
	RequestObjectType string `tf:"-"`
}

func getPermissionLevelsRequestToPb(st *GetPermissionLevelsRequest) (*getPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPermissionLevelsRequestPb{}
	pb.RequestObjectId = st.RequestObjectId

	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
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

type getPermissionLevelsRequestPb struct {
	// <needs content>
	RequestObjectId string `json:"-" url:"-"`
	// <needs content>
	RequestObjectType string `json:"-" url:"-"`
}

func getPermissionLevelsRequestFromPb(pb *getPermissionLevelsRequestPb) (*GetPermissionLevelsRequest, error) {
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
	PermissionLevels []PermissionsDescription
}

func getPermissionLevelsResponseToPb(st *GetPermissionLevelsResponse) (*getPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPermissionLevelsResponsePb{}

	var permissionLevelsPb []permissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := permissionsDescriptionToPb(&item)
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

type getPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []permissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getPermissionLevelsResponseFromPb(pb *getPermissionLevelsResponsePb) (*GetPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionLevelsResponse{}

	var permissionLevelsField []PermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := permissionsDescriptionFromPb(&itemPb)
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

func getPermissionRequestToPb(st *GetPermissionRequest) (*getPermissionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPermissionRequestPb{}
	pb.RequestObjectId = st.RequestObjectId

	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
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

type getPermissionRequestPb struct {
	// The id of the request object.
	RequestObjectId string `json:"-" url:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType string `json:"-" url:"-"`
}

func getPermissionRequestFromPb(pb *getPermissionRequestPb) (*GetPermissionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionRequest{}
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
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
	// Wire name: 'etag'
	Etag string `tf:"-"`
	// The ruleset name associated with the request.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getRuleSetRequestToPb(st *GetRuleSetRequest) (*getRuleSetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRuleSetRequestPb{}
	pb.Etag = st.Etag

	pb.Name = st.Name

	return pb, nil
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

type getRuleSetRequestPb struct {
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
}

func getRuleSetRequestFromPb(pb *getRuleSetRequestPb) (*GetRuleSetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRuleSetRequest{}
	st.Etag = pb.Etag
	st.Name = pb.Name

	return st, nil
}

// Get service principal details.
type GetServicePrincipalRequest struct {
	// Unique ID for a service principal in the Databricks workspace.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func getServicePrincipalRequestToPb(st *GetServicePrincipalRequest) (*getServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getServicePrincipalRequestPb struct {
	// Unique ID for a service principal in the Databricks workspace.
	Id string `json:"-" url:"-"`
}

func getServicePrincipalRequestFromPb(pb *getServicePrincipalRequestPb) (*GetServicePrincipalRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServicePrincipalRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetSortOrder string
type getSortOrderPb string

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

func getSortOrderToPb(st *GetSortOrder) (*getSortOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getSortOrderPb(*st)
	return &pb, nil
}

func getSortOrderFromPb(pb *getSortOrderPb) (*GetSortOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetSortOrder(*pb)
	return &st, nil
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

func getUserRequestToPb(st *GetUserRequest) (*getUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getUserRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.Id = st.Id

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getUserRequestPb struct {
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

func getUserRequestFromPb(pb *getUserRequestPb) (*GetUserRequest, error) {
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
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getUserRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getUserRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List workspace permissions
type GetWorkspaceAssignmentRequest struct {
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func getWorkspaceAssignmentRequestToPb(st *GetWorkspaceAssignmentRequest) (*getWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
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

type getWorkspaceAssignmentRequestPb struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
}

func getWorkspaceAssignmentRequestFromPb(pb *getWorkspaceAssignmentRequestPb) (*GetWorkspaceAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceAssignmentRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

type GrantRule struct {
	// Principals this grant rule applies to.
	// Wire name: 'principals'
	Principals []string
	// Role that is assigned to the list of principals.
	// Wire name: 'role'
	Role string
}

func grantRuleToPb(st *GrantRule) (*grantRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &grantRulePb{}
	pb.Principals = st.Principals

	pb.Role = st.Role

	return pb, nil
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

type grantRulePb struct {
	// Principals this grant rule applies to.
	Principals []string `json:"principals,omitempty"`
	// Role that is assigned to the list of principals.
	Role string `json:"role"`
}

func grantRuleFromPb(pb *grantRulePb) (*GrantRule, error) {
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

func groupToPb(st *Group) (*groupPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &groupPb{}
	pb.DisplayName = st.DisplayName

	var entitlementsPb []complexValuePb
	for _, item := range st.Entitlements {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entitlementsPb = append(entitlementsPb, *itemPb)
		}
	}
	pb.Entitlements = entitlementsPb

	pb.ExternalId = st.ExternalId

	var groupsPb []complexValuePb
	for _, item := range st.Groups {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			groupsPb = append(groupsPb, *itemPb)
		}
	}
	pb.Groups = groupsPb

	pb.Id = st.Id

	var membersPb []complexValuePb
	for _, item := range st.Members {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			membersPb = append(membersPb, *itemPb)
		}
	}
	pb.Members = membersPb

	metaPb, err := resourceMetaToPb(st.Meta)
	if err != nil {
		return nil, err
	}
	if metaPb != nil {
		pb.Meta = metaPb
	}

	var rolesPb []complexValuePb
	for _, item := range st.Roles {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	pb.Schemas = st.Schemas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type groupPb struct {
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	// Entitlements assigned to the group. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []complexValuePb `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []complexValuePb `json:"groups,omitempty"`
	// Databricks group ID
	Id string `json:"id,omitempty" url:"-"`

	Members []complexValuePb `json:"members,omitempty"`
	// Container for the group identifier. Workspace local versus account.
	Meta *resourceMetaPb `json:"meta,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	Roles []complexValuePb `json:"roles,omitempty"`
	// The schema of the group.
	Schemas []GroupSchema `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func groupFromPb(pb *groupPb) (*Group, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Group{}
	st.DisplayName = pb.DisplayName

	var entitlementsField []ComplexValue
	for _, itemPb := range pb.Entitlements {
		item, err := complexValueFromPb(&itemPb)
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
		item, err := complexValueFromPb(&itemPb)
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
		item, err := complexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			membersField = append(membersField, *item)
		}
	}
	st.Members = membersField
	metaField, err := resourceMetaFromPb(pb.Meta)
	if err != nil {
		return nil, err
	}
	if metaField != nil {
		st.Meta = metaField
	}

	var rolesField []ComplexValue
	for _, itemPb := range pb.Roles {
		item, err := complexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField
	st.Schemas = pb.Schemas

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *groupPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st groupPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GroupSchema string
type groupSchemaPb string

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

func groupSchemaToPb(st *GroupSchema) (*groupSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := groupSchemaPb(*st)
	return &pb, nil
}

func groupSchemaFromPb(pb *groupSchemaPb) (*GroupSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := GroupSchema(*pb)
	return &st, nil
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

func listAccountGroupsRequestToPb(st *ListAccountGroupsRequest) (*listAccountGroupsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountGroupsRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAccountGroupsRequestPb struct {
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

func listAccountGroupsRequestFromPb(pb *listAccountGroupsRequestPb) (*ListAccountGroupsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountGroupsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAccountGroupsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAccountGroupsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func listAccountServicePrincipalsRequestToPb(st *ListAccountServicePrincipalsRequest) (*listAccountServicePrincipalsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountServicePrincipalsRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAccountServicePrincipalsRequestPb struct {
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

func listAccountServicePrincipalsRequestFromPb(pb *listAccountServicePrincipalsRequestPb) (*ListAccountServicePrincipalsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountServicePrincipalsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAccountServicePrincipalsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAccountServicePrincipalsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func listAccountUsersRequestToPb(st *ListAccountUsersRequest) (*listAccountUsersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountUsersRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAccountUsersRequestPb struct {
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

func listAccountUsersRequestFromPb(pb *listAccountUsersRequestPb) (*ListAccountUsersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountUsersRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAccountUsersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAccountUsersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func listGroupsRequestToPb(st *ListGroupsRequest) (*listGroupsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listGroupsRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listGroupsRequestPb struct {
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

func listGroupsRequestFromPb(pb *listGroupsRequestPb) (*ListGroupsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGroupsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listGroupsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listGroupsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func listGroupsResponseToPb(st *ListGroupsResponse) (*listGroupsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listGroupsResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage

	var resourcesPb []groupPb
	for _, item := range st.Resources {
		itemPb, err := groupToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	pb.Schemas = st.Schemas

	pb.StartIndex = st.StartIndex

	pb.TotalResults = st.TotalResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listGroupsResponsePb struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []groupPb `json:"Resources,omitempty"`
	// The schema of the service principal.
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listGroupsResponseFromPb(pb *listGroupsResponsePb) (*ListGroupsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGroupsResponse{}
	st.ItemsPerPage = pb.ItemsPerPage

	var resourcesField []Group
	for _, itemPb := range pb.Resources {
		item, err := groupFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField
	st.Schemas = pb.Schemas
	st.StartIndex = pb.StartIndex
	st.TotalResults = pb.TotalResults

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listGroupsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listGroupsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListResponseSchema string
type listResponseSchemaPb string

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

func listResponseSchemaToPb(st *ListResponseSchema) (*listResponseSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listResponseSchemaPb(*st)
	return &pb, nil
}

func listResponseSchemaFromPb(pb *listResponseSchemaPb) (*ListResponseSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListResponseSchema(*pb)
	return &st, nil
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

func listServicePrincipalResponseToPb(st *ListServicePrincipalResponse) (*listServicePrincipalResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage

	var resourcesPb []servicePrincipalPb
	for _, item := range st.Resources {
		itemPb, err := servicePrincipalToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	pb.Schemas = st.Schemas

	pb.StartIndex = st.StartIndex

	pb.TotalResults = st.TotalResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listServicePrincipalResponsePb struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []servicePrincipalPb `json:"Resources,omitempty"`
	// The schema of the List response.
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalResponseFromPb(pb *listServicePrincipalResponsePb) (*ListServicePrincipalResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalResponse{}
	st.ItemsPerPage = pb.ItemsPerPage

	var resourcesField []ServicePrincipal
	for _, itemPb := range pb.Resources {
		item, err := servicePrincipalFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField
	st.Schemas = pb.Schemas
	st.StartIndex = pb.StartIndex
	st.TotalResults = pb.TotalResults

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func listServicePrincipalsRequestToPb(st *ListServicePrincipalsRequest) (*listServicePrincipalsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalsRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listServicePrincipalsRequestPb struct {
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

func listServicePrincipalsRequestFromPb(pb *listServicePrincipalsRequestPb) (*ListServicePrincipalsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalsRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSortOrder string
type listSortOrderPb string

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

func listSortOrderToPb(st *ListSortOrder) (*listSortOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listSortOrderPb(*st)
	return &pb, nil
}

func listSortOrderFromPb(pb *listSortOrderPb) (*ListSortOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListSortOrder(*pb)
	return &st, nil
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

func listUsersRequestToPb(st *ListUsersRequest) (*listUsersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUsersRequestPb{}
	pb.Attributes = st.Attributes

	pb.Count = st.Count

	pb.ExcludedAttributes = st.ExcludedAttributes

	pb.Filter = st.Filter

	pb.SortBy = st.SortBy

	pb.SortOrder = st.SortOrder

	pb.StartIndex = st.StartIndex

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listUsersRequestPb struct {
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

func listUsersRequestFromPb(pb *listUsersRequestPb) (*ListUsersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUsersRequest{}
	st.Attributes = pb.Attributes
	st.Count = pb.Count
	st.ExcludedAttributes = pb.ExcludedAttributes
	st.Filter = pb.Filter
	st.SortBy = pb.SortBy
	st.SortOrder = pb.SortOrder
	st.StartIndex = pb.StartIndex

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listUsersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listUsersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func listUsersResponseToPb(st *ListUsersResponse) (*listUsersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUsersResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage

	var resourcesPb []userPb
	for _, item := range st.Resources {
		itemPb, err := userToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resourcesPb = append(resourcesPb, *itemPb)
		}
	}
	pb.Resources = resourcesPb

	pb.Schemas = st.Schemas

	pb.StartIndex = st.StartIndex

	pb.TotalResults = st.TotalResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listUsersResponsePb struct {
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// User objects returned in the response.
	Resources []userPb `json:"Resources,omitempty"`
	// The schema of the List response.
	Schemas []ListResponseSchema `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. First
	// item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listUsersResponseFromPb(pb *listUsersResponsePb) (*ListUsersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUsersResponse{}
	st.ItemsPerPage = pb.ItemsPerPage

	var resourcesField []User
	for _, itemPb := range pb.Resources {
		item, err := userFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resourcesField = append(resourcesField, *item)
		}
	}
	st.Resources = resourcesField
	st.Schemas = pb.Schemas
	st.StartIndex = pb.StartIndex
	st.TotalResults = pb.TotalResults

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listUsersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listUsersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get permission assignments
type ListWorkspaceAssignmentRequest struct {
	// The workspace ID for the account.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func listWorkspaceAssignmentRequestToPb(st *ListWorkspaceAssignmentRequest) (*listWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWorkspaceAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
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

type listWorkspaceAssignmentRequestPb struct {
	// The workspace ID for the account.
	WorkspaceId int64 `json:"-" url:"-"`
}

func listWorkspaceAssignmentRequestFromPb(pb *listWorkspaceAssignmentRequestPb) (*ListWorkspaceAssignmentRequest, error) {
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

func migratePermissionsRequestToPb(st *MigratePermissionsRequest) (*migratePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &migratePermissionsRequestPb{}
	pb.FromWorkspaceGroupName = st.FromWorkspaceGroupName

	pb.Size = st.Size

	pb.ToAccountGroupName = st.ToAccountGroupName

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type migratePermissionsRequestPb struct {
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

func migratePermissionsRequestFromPb(pb *migratePermissionsRequestPb) (*MigratePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MigratePermissionsRequest{}
	st.FromWorkspaceGroupName = pb.FromWorkspaceGroupName
	st.Size = pb.Size
	st.ToAccountGroupName = pb.ToAccountGroupName
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *migratePermissionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st migratePermissionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MigratePermissionsResponse struct {
	// Number of permissions migrated.
	// Wire name: 'permissions_migrated'
	PermissionsMigrated int

	ForceSendFields []string `tf:"-"`
}

func migratePermissionsResponseToPb(st *MigratePermissionsResponse) (*migratePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &migratePermissionsResponsePb{}
	pb.PermissionsMigrated = st.PermissionsMigrated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type migratePermissionsResponsePb struct {
	// Number of permissions migrated.
	PermissionsMigrated int `json:"permissions_migrated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func migratePermissionsResponseFromPb(pb *migratePermissionsResponsePb) (*MigratePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MigratePermissionsResponse{}
	st.PermissionsMigrated = pb.PermissionsMigrated

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *migratePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st migratePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func nameToPb(st *Name) (*namePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &namePb{}
	pb.FamilyName = st.FamilyName

	pb.GivenName = st.GivenName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type namePb struct {
	// Family name of the Databricks user.
	FamilyName string `json:"familyName,omitempty"`
	// Given name of the Databricks user.
	GivenName string `json:"givenName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nameFromPb(pb *namePb) (*Name, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Name{}
	st.FamilyName = pb.FamilyName
	st.GivenName = pb.GivenName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *namePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st namePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func objectPermissionsToPb(st *ObjectPermissions) (*objectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &objectPermissionsPb{}

	var accessControlListPb []accessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := accessControlResponseToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type objectPermissionsPb struct {
	AccessControlList []accessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func objectPermissionsFromPb(pb *objectPermissionsPb) (*ObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ObjectPermissions{}

	var accessControlListField []AccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := accessControlResponseFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *objectPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st objectPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func partialUpdateToPb(st *PartialUpdate) (*partialUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partialUpdatePb{}
	pb.Id = st.Id

	var operationsPb []patchPb
	for _, item := range st.Operations {
		itemPb, err := patchToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			operationsPb = append(operationsPb, *itemPb)
		}
	}
	pb.Operations = operationsPb

	pb.Schemas = st.Schemas

	return pb, nil
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

type partialUpdatePb struct {
	// Unique ID in the Databricks workspace.
	Id string `json:"-" url:"-"`

	Operations []patchPb `json:"Operations,omitempty"`
	// The schema of the patch request. Must be
	// ["urn:ietf:params:scim:api:messages:2.0:PatchOp"].
	Schemas []PatchSchema `json:"schemas,omitempty"`
}

func partialUpdateFromPb(pb *partialUpdatePb) (*PartialUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartialUpdate{}
	st.Id = pb.Id

	var operationsField []Patch
	for _, itemPb := range pb.Operations {
		item, err := patchFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			operationsField = append(operationsField, *item)
		}
	}
	st.Operations = operationsField
	st.Schemas = pb.Schemas

	return st, nil
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

func passwordAccessControlRequestToPb(st *PasswordAccessControlRequest) (*passwordAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type passwordAccessControlRequestPb struct {
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

func passwordAccessControlRequestFromPb(pb *passwordAccessControlRequestPb) (*PasswordAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *passwordAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st passwordAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func passwordAccessControlResponseToPb(st *PasswordAccessControlResponse) (*passwordAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordAccessControlResponsePb{}

	var allPermissionsPb []passwordPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := passwordPermissionToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type passwordAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []passwordPermissionPb `json:"all_permissions,omitempty"`
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

func passwordAccessControlResponseFromPb(pb *passwordAccessControlResponsePb) (*PasswordAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordAccessControlResponse{}

	var allPermissionsField []PasswordPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := passwordPermissionFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *passwordAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st passwordAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func passwordPermissionToPb(st *PasswordPermission) (*passwordPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type passwordPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func passwordPermissionFromPb(pb *passwordPermissionPb) (*PasswordPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *passwordPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st passwordPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type PasswordPermissionLevel string
type passwordPermissionLevelPb string

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

func passwordPermissionLevelToPb(st *PasswordPermissionLevel) (*passwordPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := passwordPermissionLevelPb(*st)
	return &pb, nil
}

func passwordPermissionLevelFromPb(pb *passwordPermissionLevelPb) (*PasswordPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PasswordPermissionLevel(*pb)
	return &st, nil
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

func passwordPermissionsToPb(st *PasswordPermissions) (*passwordPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordPermissionsPb{}

	var accessControlListPb []passwordAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := passwordAccessControlResponseToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type passwordPermissionsPb struct {
	AccessControlList []passwordAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func passwordPermissionsFromPb(pb *passwordPermissionsPb) (*PasswordPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissions{}

	var accessControlListField []PasswordAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := passwordAccessControlResponseFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *passwordPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st passwordPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PasswordPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func passwordPermissionsDescriptionToPb(st *PasswordPermissionsDescription) (*passwordPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type passwordPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PasswordPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func passwordPermissionsDescriptionFromPb(pb *passwordPermissionsDescriptionPb) (*PasswordPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *passwordPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st passwordPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PasswordPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []PasswordAccessControlRequest
}

func passwordPermissionsRequestToPb(st *PasswordPermissionsRequest) (*passwordPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordPermissionsRequestPb{}

	var accessControlListPb []passwordAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := passwordAccessControlRequestToPb(&item)
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

type passwordPermissionsRequestPb struct {
	AccessControlList []passwordAccessControlRequestPb `json:"access_control_list,omitempty"`
}

func passwordPermissionsRequestFromPb(pb *passwordPermissionsRequestPb) (*PasswordPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissionsRequest{}

	var accessControlListField []PasswordAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := passwordAccessControlRequestFromPb(&itemPb)
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
	Op PatchOp
	// Selection of patch operation
	// Wire name: 'path'
	Path string
	// Value to modify
	// Wire name: 'value'
	Value any

	ForceSendFields []string `tf:"-"`
}

func patchToPb(st *Patch) (*patchPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchPb{}
	pb.Op = st.Op

	pb.Path = st.Path

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type patchPb struct {
	// Type of patch operation.
	Op PatchOp `json:"op,omitempty"`
	// Selection of patch operation
	Path string `json:"path,omitempty"`
	// Value to modify
	Value any `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func patchFromPb(pb *patchPb) (*Patch, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Patch{}
	st.Op = pb.Op
	st.Path = pb.Path
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *patchPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st patchPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Type of patch operation.
type PatchOp string
type patchOpPb string

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

func patchOpToPb(st *PatchOp) (*patchOpPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := patchOpPb(*st)
	return &pb, nil
}

func patchOpFromPb(pb *patchOpPb) (*PatchOp, error) {
	if pb == nil {
		return nil, nil
	}
	st := PatchOp(*pb)
	return &st, nil
}

type PatchResponse struct {
}

func patchResponseToPb(st *PatchResponse) (*patchResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchResponsePb{}

	return pb, nil
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

type patchResponsePb struct {
}

func patchResponseFromPb(pb *patchResponsePb) (*PatchResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchResponse{}

	return st, nil
}

type PatchSchema string
type patchSchemaPb string

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

func patchSchemaToPb(st *PatchSchema) (*patchSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := patchSchemaPb(*st)
	return &pb, nil
}

func patchSchemaFromPb(pb *patchSchemaPb) (*PatchSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := PatchSchema(*pb)
	return &st, nil
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

func permissionToPb(st *Permission) (*permissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type permissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionFromPb(pb *permissionPb) (*Permission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Permission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func permissionAssignmentToPb(st *PermissionAssignment) (*permissionAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionAssignmentPb{}
	pb.Error = st.Error

	pb.Permissions = st.Permissions

	principalPb, err := principalOutputToPb(st.Principal)
	if err != nil {
		return nil, err
	}
	if principalPb != nil {
		pb.Principal = principalPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type permissionAssignmentPb struct {
	// Error response associated with a workspace permission assignment, if any.
	Error string `json:"error,omitempty"`
	// The permissions level of the principal.
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	// Information about the principal assigned to the workspace.
	Principal *principalOutputPb `json:"principal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionAssignmentFromPb(pb *permissionAssignmentPb) (*PermissionAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionAssignment{}
	st.Error = pb.Error
	st.Permissions = pb.Permissions
	principalField, err := principalOutputFromPb(pb.Principal)
	if err != nil {
		return nil, err
	}
	if principalField != nil {
		st.Principal = principalField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionAssignments struct {
	// Array of permissions assignments defined for a workspace.
	// Wire name: 'permission_assignments'
	PermissionAssignments []PermissionAssignment
}

func permissionAssignmentsToPb(st *PermissionAssignments) (*permissionAssignmentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionAssignmentsPb{}

	var permissionAssignmentsPb []permissionAssignmentPb
	for _, item := range st.PermissionAssignments {
		itemPb, err := permissionAssignmentToPb(&item)
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

type permissionAssignmentsPb struct {
	// Array of permissions assignments defined for a workspace.
	PermissionAssignments []permissionAssignmentPb `json:"permission_assignments,omitempty"`
}

func permissionAssignmentsFromPb(pb *permissionAssignmentsPb) (*PermissionAssignments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionAssignments{}

	var permissionAssignmentsField []PermissionAssignment
	for _, itemPb := range pb.PermissionAssignments {
		item, err := permissionAssignmentFromPb(&itemPb)
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
type permissionLevelPb string

const PermissionLevelCanAttachTo PermissionLevel = `CAN_ATTACH_TO`

const PermissionLevelCanBind PermissionLevel = `CAN_BIND`

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanEditMetadata PermissionLevel = `CAN_EDIT_METADATA`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageRun PermissionLevel = `CAN_MANAGE_RUN`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanMonitor PermissionLevel = `CAN_MONITOR`

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
	case `CAN_ATTACH_TO`, `CAN_BIND`, `CAN_EDIT`, `CAN_EDIT_METADATA`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_RUN`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_MONITOR`, `CAN_QUERY`, `CAN_READ`, `CAN_RESTART`, `CAN_RUN`, `CAN_USE`, `CAN_VIEW`, `CAN_VIEW_METADATA`, `IS_OWNER`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_BIND", "CAN_EDIT", "CAN_EDIT_METADATA", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_RUN", "CAN_MANAGE_STAGING_VERSIONS", "CAN_MONITOR", "CAN_QUERY", "CAN_READ", "CAN_RESTART", "CAN_RUN", "CAN_USE", "CAN_VIEW", "CAN_VIEW_METADATA", "IS_OWNER"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

func permissionLevelToPb(st *PermissionLevel) (*permissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := permissionLevelPb(*st)
	return &pb, nil
}

func permissionLevelFromPb(pb *permissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
}

type PermissionOutput struct {
	// The results of a permissions query.
	// Wire name: 'description'
	Description string

	// Wire name: 'permission_level'
	PermissionLevel WorkspacePermission

	ForceSendFields []string `tf:"-"`
}

func permissionOutputToPb(st *PermissionOutput) (*permissionOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionOutputPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type permissionOutputPb struct {
	// The results of a permissions query.
	Description string `json:"description,omitempty"`

	PermissionLevel WorkspacePermission `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionOutputFromPb(pb *permissionOutputPb) (*PermissionOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionOutput{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel

	ForceSendFields []string `tf:"-"`
}

func permissionsDescriptionToPb(st *PermissionsDescription) (*permissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type permissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionsDescriptionFromPb(pb *permissionsDescriptionPb) (*PermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermissionsRequest struct {

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

func permissionsRequestToPb(st *PermissionsRequest) (*permissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsRequestPb{}

	var accessControlListPb []accessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := accessControlRequestToPb(&item)
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

func (st *PermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := permissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type permissionsRequestPb struct {
	AccessControlList []accessControlRequestPb `json:"access_control_list,omitempty"`
	// The id of the request object.
	RequestObjectId string `json:"-" url:"-"`
	// The type of the request object. Can be one of the following: alerts,
	// authorization, clusters, cluster-policies, dashboards, dbsql-dashboards,
	// directories, experiments, files, instance-pools, jobs, notebooks,
	// pipelines, queries, registered-models, repos, serving-endpoints, or
	// warehouses.
	RequestObjectType string `json:"-" url:"-"`
}

func permissionsRequestFromPb(pb *permissionsRequestPb) (*PermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsRequest{}

	var accessControlListField []AccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := accessControlRequestFromPb(&itemPb)
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

func principalOutputToPb(st *PrincipalOutput) (*principalOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &principalOutputPb{}
	pb.DisplayName = st.DisplayName

	pb.GroupName = st.GroupName

	pb.PrincipalId = st.PrincipalId

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type principalOutputPb struct {
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

func principalOutputFromPb(pb *principalOutputPb) (*PrincipalOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrincipalOutput{}
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.PrincipalId = pb.PrincipalId
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *principalOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st principalOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Defines the identity to be used for authZ of the request on the server side.
// See one pager for for more information: http://go/acl/service-identity
type RequestAuthzIdentity string
type requestAuthzIdentityPb string

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

func requestAuthzIdentityToPb(st *RequestAuthzIdentity) (*requestAuthzIdentityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := requestAuthzIdentityPb(*st)
	return &pb, nil
}

func requestAuthzIdentityFromPb(pb *requestAuthzIdentityPb) (*RequestAuthzIdentity, error) {
	if pb == nil {
		return nil, nil
	}
	st := RequestAuthzIdentity(*pb)
	return &st, nil
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

func resourceInfoToPb(st *ResourceInfo) (*resourceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resourceInfoPb{}
	pb.Id = st.Id

	pb.LegacyAclPath = st.LegacyAclPath

	parentResourceInfoPb, err := resourceInfoToPb(st.ParentResourceInfo)
	if err != nil {
		return nil, err
	}
	if parentResourceInfoPb != nil {
		pb.ParentResourceInfo = parentResourceInfoPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resourceInfoPb struct {
	// Id of the current resource.
	Id string `json:"id" url:"id"`
	// The legacy acl path of the current resource.
	LegacyAclPath string `json:"legacy_acl_path,omitempty" url:"legacy_acl_path,omitempty"`
	// Parent resource info for the current resource. The parent may have
	// another parent.
	ParentResourceInfo *resourceInfoPb `json:"parent_resource_info,omitempty" url:"parent_resource_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resourceInfoFromPb(pb *resourceInfoPb) (*ResourceInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResourceInfo{}
	st.Id = pb.Id
	st.LegacyAclPath = pb.LegacyAclPath
	parentResourceInfoField, err := resourceInfoFromPb(pb.ParentResourceInfo)
	if err != nil {
		return nil, err
	}
	if parentResourceInfoField != nil {
		st.ParentResourceInfo = parentResourceInfoField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resourceInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resourceInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResourceMeta struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	// Wire name: 'resourceType'
	ResourceType string

	ForceSendFields []string `tf:"-"`
}

func resourceMetaToPb(st *ResourceMeta) (*resourceMetaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resourceMetaPb{}
	pb.ResourceType = st.ResourceType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resourceMetaPb struct {
	// Identifier for group type. Can be local workspace group
	// (`WorkspaceGroup`) or account group (`Group`).
	ResourceType string `json:"resourceType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resourceMetaFromPb(pb *resourceMetaPb) (*ResourceMeta, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResourceMeta{}
	st.ResourceType = pb.ResourceType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resourceMetaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resourceMetaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Role struct {
	// Role to assign to a principal or a list of principals on a resource.
	// Wire name: 'name'
	Name string
}

func roleToPb(st *Role) (*rolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rolePb{}
	pb.Name = st.Name

	return pb, nil
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

type rolePb struct {
	// Role to assign to a principal or a list of principals on a resource.
	Name string `json:"name"`
}

func roleFromPb(pb *rolePb) (*Role, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Role{}
	st.Name = pb.Name

	return st, nil
}

type RuleSetResponse struct {
	// Identifies the version of the rule set returned.
	// Wire name: 'etag'
	Etag string

	// Wire name: 'grant_rules'
	GrantRules []GrantRule
	// Name of the rule set.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func ruleSetResponseToPb(st *RuleSetResponse) (*ruleSetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ruleSetResponsePb{}
	pb.Etag = st.Etag

	var grantRulesPb []grantRulePb
	for _, item := range st.GrantRules {
		itemPb, err := grantRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			grantRulesPb = append(grantRulesPb, *itemPb)
		}
	}
	pb.GrantRules = grantRulesPb

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type ruleSetResponsePb struct {
	// Identifies the version of the rule set returned.
	Etag string `json:"etag,omitempty"`

	GrantRules []grantRulePb `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ruleSetResponseFromPb(pb *ruleSetResponsePb) (*RuleSetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RuleSetResponse{}
	st.Etag = pb.Etag

	var grantRulesField []GrantRule
	for _, itemPb := range pb.GrantRules {
		item, err := grantRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			grantRulesField = append(grantRulesField, *item)
		}
	}
	st.GrantRules = grantRulesField
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ruleSetResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ruleSetResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RuleSetUpdateRequest struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	// Wire name: 'etag'
	Etag string

	// Wire name: 'grant_rules'
	GrantRules []GrantRule
	// Name of the rule set.
	// Wire name: 'name'
	Name string
}

func ruleSetUpdateRequestToPb(st *RuleSetUpdateRequest) (*ruleSetUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ruleSetUpdateRequestPb{}
	pb.Etag = st.Etag

	var grantRulesPb []grantRulePb
	for _, item := range st.GrantRules {
		itemPb, err := grantRuleToPb(&item)
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

type ruleSetUpdateRequestPb struct {
	// The expected etag of the rule set to update. The update will fail if the
	// value does not match the value that is stored in account access control
	// service.
	Etag string `json:"etag"`

	GrantRules []grantRulePb `json:"grant_rules,omitempty"`
	// Name of the rule set.
	Name string `json:"name"`
}

func ruleSetUpdateRequestFromPb(pb *ruleSetUpdateRequestPb) (*RuleSetUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RuleSetUpdateRequest{}
	st.Etag = pb.Etag

	var grantRulesField []GrantRule
	for _, itemPb := range pb.GrantRules {
		item, err := grantRuleFromPb(&itemPb)
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

func servicePrincipalToPb(st *ServicePrincipal) (*servicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servicePrincipalPb{}
	pb.Active = st.Active

	pb.ApplicationId = st.ApplicationId

	pb.DisplayName = st.DisplayName

	var entitlementsPb []complexValuePb
	for _, item := range st.Entitlements {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entitlementsPb = append(entitlementsPb, *itemPb)
		}
	}
	pb.Entitlements = entitlementsPb

	pb.ExternalId = st.ExternalId

	var groupsPb []complexValuePb
	for _, item := range st.Groups {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			groupsPb = append(groupsPb, *itemPb)
		}
	}
	pb.Groups = groupsPb

	pb.Id = st.Id

	var rolesPb []complexValuePb
	for _, item := range st.Roles {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	pb.Schemas = st.Schemas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servicePrincipalPb struct {
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
	Entitlements []complexValuePb `json:"entitlements,omitempty"`

	ExternalId string `json:"externalId,omitempty"`

	Groups []complexValuePb `json:"groups,omitempty"`
	// Databricks service principal ID.
	Id string `json:"id,omitempty" url:"-"`
	// Corresponds to AWS instance profile/arn role.
	Roles []complexValuePb `json:"roles,omitempty"`
	// The schema of the List response.
	Schemas []ServicePrincipalSchema `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servicePrincipalFromPb(pb *servicePrincipalPb) (*ServicePrincipal, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServicePrincipal{}
	st.Active = pb.Active
	st.ApplicationId = pb.ApplicationId
	st.DisplayName = pb.DisplayName

	var entitlementsField []ComplexValue
	for _, itemPb := range pb.Entitlements {
		item, err := complexValueFromPb(&itemPb)
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
		item, err := complexValueFromPb(&itemPb)
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
		item, err := complexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField
	st.Schemas = pb.Schemas

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servicePrincipalPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servicePrincipalPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServicePrincipalSchema string
type servicePrincipalSchemaPb string

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

func servicePrincipalSchemaToPb(st *ServicePrincipalSchema) (*servicePrincipalSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servicePrincipalSchemaPb(*st)
	return &pb, nil
}

func servicePrincipalSchemaFromPb(pb *servicePrincipalSchemaPb) (*ServicePrincipalSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServicePrincipalSchema(*pb)
	return &st, nil
}

type UpdateResponse struct {
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
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

type updateResponsePb struct {
}

func updateResponseFromPb(pb *updateResponsePb) (*UpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResponse{}

	return st, nil
}

type UpdateRuleSetRequest struct {
	// Name of the rule set.
	// Wire name: 'name'
	Name string

	// Wire name: 'rule_set'
	RuleSet RuleSetUpdateRequest
}

func updateRuleSetRequestToPb(st *UpdateRuleSetRequest) (*updateRuleSetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRuleSetRequestPb{}
	pb.Name = st.Name

	ruleSetPb, err := ruleSetUpdateRequestToPb(&st.RuleSet)
	if err != nil {
		return nil, err
	}
	if ruleSetPb != nil {
		pb.RuleSet = *ruleSetPb
	}

	return pb, nil
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

type updateRuleSetRequestPb struct {
	// Name of the rule set.
	Name string `json:"name"`

	RuleSet ruleSetUpdateRequestPb `json:"rule_set"`
}

func updateRuleSetRequestFromPb(pb *updateRuleSetRequestPb) (*UpdateRuleSetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRuleSetRequest{}
	st.Name = pb.Name
	ruleSetField, err := ruleSetUpdateRequestFromPb(&pb.RuleSet)
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
	Permissions []WorkspacePermission
	// The ID of the user, service principal, or group.
	// Wire name: 'principal_id'
	PrincipalId int64 `tf:"-"`
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func updateWorkspaceAssignmentsToPb(st *UpdateWorkspaceAssignments) (*updateWorkspaceAssignmentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceAssignmentsPb{}
	pb.Permissions = st.Permissions

	pb.PrincipalId = st.PrincipalId

	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
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

type updateWorkspaceAssignmentsPb struct {
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

func updateWorkspaceAssignmentsFromPb(pb *updateWorkspaceAssignmentsPb) (*UpdateWorkspaceAssignments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceAssignments{}
	st.Permissions = pb.Permissions
	st.PrincipalId = pb.PrincipalId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
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

func userToPb(st *User) (*userPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &userPb{}
	pb.Active = st.Active

	pb.DisplayName = st.DisplayName

	var emailsPb []complexValuePb
	for _, item := range st.Emails {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			emailsPb = append(emailsPb, *itemPb)
		}
	}
	pb.Emails = emailsPb

	var entitlementsPb []complexValuePb
	for _, item := range st.Entitlements {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entitlementsPb = append(entitlementsPb, *itemPb)
		}
	}
	pb.Entitlements = entitlementsPb

	pb.ExternalId = st.ExternalId

	var groupsPb []complexValuePb
	for _, item := range st.Groups {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			groupsPb = append(groupsPb, *itemPb)
		}
	}
	pb.Groups = groupsPb

	pb.Id = st.Id

	namePb, err := nameToPb(st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = namePb
	}

	var rolesPb []complexValuePb
	for _, item := range st.Roles {
		itemPb, err := complexValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rolesPb = append(rolesPb, *itemPb)
		}
	}
	pb.Roles = rolesPb

	pb.Schemas = st.Schemas

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type userPb struct {
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
	Emails []complexValuePb `json:"emails,omitempty"`
	// Entitlements assigned to the user. See [assigning entitlements] for a
	// full list of supported values.
	//
	// [assigning entitlements]: https://docs.databricks.com/administration-guide/users-groups/index.html#assigning-entitlements
	Entitlements []complexValuePb `json:"entitlements,omitempty"`
	// External ID is not currently supported. It is reserved for future use.
	ExternalId string `json:"externalId,omitempty"`

	Groups []complexValuePb `json:"groups,omitempty"`
	// Databricks user ID.
	Id string `json:"id,omitempty" url:"-"`

	Name *namePb `json:"name,omitempty"`
	// Corresponds to AWS instance profile/arn role.
	Roles []complexValuePb `json:"roles,omitempty"`
	// The schema of the user.
	Schemas []UserSchema `json:"schemas,omitempty"`
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func userFromPb(pb *userPb) (*User, error) {
	if pb == nil {
		return nil, nil
	}
	st := &User{}
	st.Active = pb.Active
	st.DisplayName = pb.DisplayName

	var emailsField []ComplexValue
	for _, itemPb := range pb.Emails {
		item, err := complexValueFromPb(&itemPb)
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
		item, err := complexValueFromPb(&itemPb)
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
		item, err := complexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			groupsField = append(groupsField, *item)
		}
	}
	st.Groups = groupsField
	st.Id = pb.Id
	nameField, err := nameFromPb(pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = nameField
	}

	var rolesField []ComplexValue
	for _, itemPb := range pb.Roles {
		item, err := complexValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rolesField = append(rolesField, *item)
		}
	}
	st.Roles = rolesField
	st.Schemas = pb.Schemas
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *userPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st userPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UserSchema string
type userSchemaPb string

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

func userSchemaToPb(st *UserSchema) (*userSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := userSchemaPb(*st)
	return &pb, nil
}

func userSchemaFromPb(pb *userSchemaPb) (*UserSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := UserSchema(*pb)
	return &st, nil
}

type WorkspacePermission string
type workspacePermissionPb string

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

func workspacePermissionToPb(st *WorkspacePermission) (*workspacePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacePermissionPb(*st)
	return &pb, nil
}

func workspacePermissionFromPb(pb *workspacePermissionPb) (*WorkspacePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspacePermission(*pb)
	return &st, nil
}

type WorkspacePermissions struct {
	// Array of permissions defined for a workspace.
	// Wire name: 'permissions'
	Permissions []PermissionOutput
}

func workspacePermissionsToPb(st *WorkspacePermissions) (*workspacePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacePermissionsPb{}

	var permissionsPb []permissionOutputPb
	for _, item := range st.Permissions {
		itemPb, err := permissionOutputToPb(&item)
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

type workspacePermissionsPb struct {
	// Array of permissions defined for a workspace.
	Permissions []permissionOutputPb `json:"permissions,omitempty"`
}

func workspacePermissionsFromPb(pb *workspacePermissionsPb) (*WorkspacePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspacePermissions{}

	var permissionsField []PermissionOutput
	for _, itemPb := range pb.Permissions {
		item, err := permissionOutputFromPb(&itemPb)
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
