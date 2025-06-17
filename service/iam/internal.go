// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package iam

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

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

type accessControlRequestPb struct {
	GroupName            string          `json:"group_name,omitempty"`
	PermissionLevel      PermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string          `json:"service_principal_name,omitempty"`
	UserName             string          `json:"user_name,omitempty"`

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

func accessControlResponseToPb(st *AccessControlResponse) (*accessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type accessControlResponsePb struct {
	AllPermissions       []Permission `json:"all_permissions,omitempty"`
	DisplayName          string       `json:"display_name,omitempty"`
	GroupName            string       `json:"group_name,omitempty"`
	ServicePrincipalName string       `json:"service_principal_name,omitempty"`
	UserName             string       `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func accessControlResponseFromPb(pb *accessControlResponsePb) (*AccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
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

func actorToPb(st *Actor) (*actorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &actorPb{}
	pb.ActorId = st.ActorId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

func checkPolicyRequestToPb(st *CheckPolicyRequest) (*checkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &checkPolicyRequestPb{}
	pb.Actor = st.Actor
	pb.AuthzIdentity = st.AuthzIdentity
	pb.ConsistencyToken = st.ConsistencyToken
	pb.Permission = st.Permission
	pb.Resource = st.Resource
	pb.ResourceInfo = st.ResourceInfo

	return pb, nil
}

type checkPolicyRequestPb struct {
	Actor            Actor                `json:"-" url:"actor"`
	AuthzIdentity    RequestAuthzIdentity `json:"-" url:"authz_identity"`
	ConsistencyToken ConsistencyToken     `json:"-" url:"consistency_token"`
	Permission       string               `json:"-" url:"permission"`
	Resource         string               `json:"-" url:"resource"`
	ResourceInfo     *ResourceInfo        `json:"-" url:"resource_info,omitempty"`
}

func checkPolicyRequestFromPb(pb *checkPolicyRequestPb) (*CheckPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CheckPolicyRequest{}
	st.Actor = pb.Actor
	st.AuthzIdentity = pb.AuthzIdentity
	st.ConsistencyToken = pb.ConsistencyToken
	st.Permission = pb.Permission
	st.Resource = pb.Resource
	st.ResourceInfo = pb.ResourceInfo

	return st, nil
}

func checkPolicyResponseToPb(st *CheckPolicyResponse) (*checkPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &checkPolicyResponsePb{}
	pb.ConsistencyToken = st.ConsistencyToken
	pb.IsPermitted = st.IsPermitted

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type checkPolicyResponsePb struct {
	ConsistencyToken ConsistencyToken `json:"consistency_token"`
	IsPermitted      bool             `json:"is_permitted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func checkPolicyResponseFromPb(pb *checkPolicyResponsePb) (*CheckPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CheckPolicyResponse{}
	st.ConsistencyToken = pb.ConsistencyToken
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

type complexValuePb struct {
	Display string `json:"display,omitempty"`
	Primary bool   `json:"primary,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`

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

func consistencyTokenToPb(st *ConsistencyToken) (*consistencyTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &consistencyTokenPb{}
	pb.Value = st.Value

	return pb, nil
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

func deleteAccountGroupRequestToPb(st *DeleteAccountGroupRequest) (*deleteAccountGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteAccountGroupRequestPb struct {
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

func deleteAccountServicePrincipalRequestToPb(st *DeleteAccountServicePrincipalRequest) (*deleteAccountServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteAccountServicePrincipalRequestPb struct {
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

func deleteAccountUserRequestToPb(st *DeleteAccountUserRequest) (*deleteAccountUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountUserRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteAccountUserRequestPb struct {
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

func deleteGroupRequestToPb(st *DeleteGroupRequest) (*deleteGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteGroupRequestPb struct {
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

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

func deleteServicePrincipalRequestToPb(st *DeleteServicePrincipalRequest) (*deleteServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteServicePrincipalRequestPb struct {
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

func deleteUserRequestToPb(st *DeleteUserRequest) (*deleteUserRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteUserRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteUserRequestPb struct {
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

func deleteWorkspaceAssignmentRequestToPb(st *DeleteWorkspaceAssignmentRequest) (*deleteWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspaceAssignmentRequestPb{}
	pb.PrincipalId = st.PrincipalId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type deleteWorkspaceAssignmentRequestPb struct {
	PrincipalId int64 `json:"-" url:"-"`
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

func deleteWorkspacePermissionAssignmentResponseToPb(st *DeleteWorkspacePermissionAssignmentResponse) (*deleteWorkspacePermissionAssignmentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspacePermissionAssignmentResponsePb{}

	return pb, nil
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

func getAccountGroupRequestToPb(st *GetAccountGroupRequest) (*getAccountGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getAccountGroupRequestPb struct {
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

func getAccountServicePrincipalRequestToPb(st *GetAccountServicePrincipalRequest) (*getAccountServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getAccountServicePrincipalRequestPb struct {
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

type getAccountUserRequestPb struct {
	Attributes         string       `json:"-" url:"attributes,omitempty"`
	Count              int          `json:"-" url:"count,omitempty"`
	ExcludedAttributes string       `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string       `json:"-" url:"filter,omitempty"`
	Id                 string       `json:"-" url:"-"`
	SortBy             string       `json:"-" url:"sortBy,omitempty"`
	SortOrder          GetSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int          `json:"-" url:"startIndex,omitempty"`

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

func getAssignableRolesForResourceRequestToPb(st *GetAssignableRolesForResourceRequest) (*getAssignableRolesForResourceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAssignableRolesForResourceRequestPb{}
	pb.Resource = st.Resource

	return pb, nil
}

type getAssignableRolesForResourceRequestPb struct {
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

func getAssignableRolesForResourceResponseToPb(st *GetAssignableRolesForResourceResponse) (*getAssignableRolesForResourceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAssignableRolesForResourceResponsePb{}
	pb.Roles = st.Roles

	return pb, nil
}

type getAssignableRolesForResourceResponsePb struct {
	Roles []Role `json:"roles,omitempty"`
}

func getAssignableRolesForResourceResponseFromPb(pb *getAssignableRolesForResourceResponsePb) (*GetAssignableRolesForResourceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAssignableRolesForResourceResponse{}
	st.Roles = pb.Roles

	return st, nil
}

func getGroupRequestToPb(st *GetGroupRequest) (*getGroupRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGroupRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getGroupRequestPb struct {
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

func getPasswordPermissionLevelsResponseToPb(st *GetPasswordPermissionLevelsResponse) (*getPasswordPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPasswordPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getPasswordPermissionLevelsResponsePb struct {
	PermissionLevels []PasswordPermissionsDescription `json:"permission_levels,omitempty"`
}

func getPasswordPermissionLevelsResponseFromPb(pb *getPasswordPermissionLevelsResponsePb) (*GetPasswordPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPasswordPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
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

type getPermissionLevelsRequestPb struct {
	RequestObjectId   string `json:"-" url:"-"`
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

func getPermissionLevelsResponseToPb(st *GetPermissionLevelsResponse) (*getPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getPermissionLevelsResponsePb struct {
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

func getPermissionLevelsResponseFromPb(pb *getPermissionLevelsResponsePb) (*GetPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
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

type getPermissionRequestPb struct {
	RequestObjectId   string `json:"-" url:"-"`
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

func getRuleSetRequestToPb(st *GetRuleSetRequest) (*getRuleSetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRuleSetRequestPb{}
	pb.Etag = st.Etag
	pb.Name = st.Name

	return pb, nil
}

type getRuleSetRequestPb struct {
	Etag string `json:"-" url:"etag"`
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

func getServicePrincipalRequestToPb(st *GetServicePrincipalRequest) (*getServicePrincipalRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServicePrincipalRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getServicePrincipalRequestPb struct {
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

type getUserRequestPb struct {
	Attributes         string       `json:"-" url:"attributes,omitempty"`
	Count              int          `json:"-" url:"count,omitempty"`
	ExcludedAttributes string       `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string       `json:"-" url:"filter,omitempty"`
	Id                 string       `json:"-" url:"-"`
	SortBy             string       `json:"-" url:"sortBy,omitempty"`
	SortOrder          GetSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int          `json:"-" url:"startIndex,omitempty"`

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

func getWorkspaceAssignmentRequestToPb(st *GetWorkspaceAssignmentRequest) (*getWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type getWorkspaceAssignmentRequestPb struct {
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

func grantRuleToPb(st *GrantRule) (*grantRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &grantRulePb{}
	pb.Principals = st.Principals
	pb.Role = st.Role

	return pb, nil
}

type grantRulePb struct {
	Principals []string `json:"principals,omitempty"`
	Role       string   `json:"role"`
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

func groupToPb(st *Group) (*groupPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &groupPb{}
	pb.DisplayName = st.DisplayName
	pb.Entitlements = st.Entitlements
	pb.ExternalId = st.ExternalId
	pb.Groups = st.Groups
	pb.Id = st.Id
	pb.Members = st.Members
	pb.Meta = st.Meta
	pb.Roles = st.Roles
	pb.Schemas = st.Schemas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type groupPb struct {
	DisplayName  string         `json:"displayName,omitempty"`
	Entitlements []ComplexValue `json:"entitlements,omitempty"`
	ExternalId   string         `json:"externalId,omitempty"`
	Groups       []ComplexValue `json:"groups,omitempty"`
	Id           string         `json:"id,omitempty" url:"-"`
	Members      []ComplexValue `json:"members,omitempty"`
	Meta         *ResourceMeta  `json:"meta,omitempty"`
	Roles        []ComplexValue `json:"roles,omitempty"`
	Schemas      []GroupSchema  `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func groupFromPb(pb *groupPb) (*Group, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Group{}
	st.DisplayName = pb.DisplayName
	st.Entitlements = pb.Entitlements
	st.ExternalId = pb.ExternalId
	st.Groups = pb.Groups
	st.Id = pb.Id
	st.Members = pb.Members
	st.Meta = pb.Meta
	st.Roles = pb.Roles
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

type listAccountGroupsRequestPb struct {
	Attributes         string        `json:"-" url:"attributes,omitempty"`
	Count              int64         `json:"-" url:"count,omitempty"`
	ExcludedAttributes string        `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string        `json:"-" url:"filter,omitempty"`
	SortBy             string        `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64         `json:"-" url:"startIndex,omitempty"`

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

type listAccountServicePrincipalsRequestPb struct {
	Attributes         string        `json:"-" url:"attributes,omitempty"`
	Count              int64         `json:"-" url:"count,omitempty"`
	ExcludedAttributes string        `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string        `json:"-" url:"filter,omitempty"`
	SortBy             string        `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64         `json:"-" url:"startIndex,omitempty"`

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

type listAccountUsersRequestPb struct {
	Attributes         string        `json:"-" url:"attributes,omitempty"`
	Count              int64         `json:"-" url:"count,omitempty"`
	ExcludedAttributes string        `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string        `json:"-" url:"filter,omitempty"`
	SortBy             string        `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64         `json:"-" url:"startIndex,omitempty"`

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

type listGroupsRequestPb struct {
	Attributes         string        `json:"-" url:"attributes,omitempty"`
	Count              int64         `json:"-" url:"count,omitempty"`
	ExcludedAttributes string        `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string        `json:"-" url:"filter,omitempty"`
	SortBy             string        `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64         `json:"-" url:"startIndex,omitempty"`

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

func listGroupsResponseToPb(st *ListGroupsResponse) (*listGroupsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listGroupsResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage
	pb.Resources = st.Resources
	pb.Schemas = st.Schemas
	pb.StartIndex = st.StartIndex
	pb.TotalResults = st.TotalResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listGroupsResponsePb struct {
	ItemsPerPage int64                `json:"itemsPerPage,omitempty"`
	Resources    []Group              `json:"Resources,omitempty"`
	Schemas      []ListResponseSchema `json:"schemas,omitempty"`
	StartIndex   int64                `json:"startIndex,omitempty"`
	TotalResults int64                `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listGroupsResponseFromPb(pb *listGroupsResponsePb) (*ListGroupsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGroupsResponse{}
	st.ItemsPerPage = pb.ItemsPerPage
	st.Resources = pb.Resources
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

func listServicePrincipalResponseToPb(st *ListServicePrincipalResponse) (*listServicePrincipalResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage
	pb.Resources = st.Resources
	pb.Schemas = st.Schemas
	pb.StartIndex = st.StartIndex
	pb.TotalResults = st.TotalResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listServicePrincipalResponsePb struct {
	ItemsPerPage int64                `json:"itemsPerPage,omitempty"`
	Resources    []ServicePrincipal   `json:"Resources,omitempty"`
	Schemas      []ListResponseSchema `json:"schemas,omitempty"`
	StartIndex   int64                `json:"startIndex,omitempty"`
	TotalResults int64                `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalResponseFromPb(pb *listServicePrincipalResponsePb) (*ListServicePrincipalResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalResponse{}
	st.ItemsPerPage = pb.ItemsPerPage
	st.Resources = pb.Resources
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

type listServicePrincipalsRequestPb struct {
	Attributes         string        `json:"-" url:"attributes,omitempty"`
	Count              int64         `json:"-" url:"count,omitempty"`
	ExcludedAttributes string        `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string        `json:"-" url:"filter,omitempty"`
	SortBy             string        `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64         `json:"-" url:"startIndex,omitempty"`

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

type listUsersRequestPb struct {
	Attributes         string        `json:"-" url:"attributes,omitempty"`
	Count              int64         `json:"-" url:"count,omitempty"`
	ExcludedAttributes string        `json:"-" url:"excludedAttributes,omitempty"`
	Filter             string        `json:"-" url:"filter,omitempty"`
	SortBy             string        `json:"-" url:"sortBy,omitempty"`
	SortOrder          ListSortOrder `json:"-" url:"sortOrder,omitempty"`
	StartIndex         int64         `json:"-" url:"startIndex,omitempty"`

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

func listUsersResponseToPb(st *ListUsersResponse) (*listUsersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUsersResponsePb{}
	pb.ItemsPerPage = st.ItemsPerPage
	pb.Resources = st.Resources
	pb.Schemas = st.Schemas
	pb.StartIndex = st.StartIndex
	pb.TotalResults = st.TotalResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listUsersResponsePb struct {
	ItemsPerPage int64                `json:"itemsPerPage,omitempty"`
	Resources    []User               `json:"Resources,omitempty"`
	Schemas      []ListResponseSchema `json:"schemas,omitempty"`
	StartIndex   int64                `json:"startIndex,omitempty"`
	TotalResults int64                `json:"totalResults,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listUsersResponseFromPb(pb *listUsersResponsePb) (*ListUsersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUsersResponse{}
	st.ItemsPerPage = pb.ItemsPerPage
	st.Resources = pb.Resources
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

func listWorkspaceAssignmentRequestToPb(st *ListWorkspaceAssignmentRequest) (*listWorkspaceAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWorkspaceAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type listWorkspaceAssignmentRequestPb struct {
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

type migratePermissionsRequestPb struct {
	FromWorkspaceGroupName string `json:"from_workspace_group_name"`
	Size                   int    `json:"size,omitempty"`
	ToAccountGroupName     string `json:"to_account_group_name"`
	WorkspaceId            int64  `json:"workspace_id"`

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

func migratePermissionsResponseToPb(st *MigratePermissionsResponse) (*migratePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &migratePermissionsResponsePb{}
	pb.PermissionsMigrated = st.PermissionsMigrated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type migratePermissionsResponsePb struct {
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

type namePb struct {
	FamilyName string `json:"familyName,omitempty"`
	GivenName  string `json:"givenName,omitempty"`

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

func objectPermissionsToPb(st *ObjectPermissions) (*objectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &objectPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type objectPermissionsPb struct {
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                  `json:"object_id,omitempty"`
	ObjectType        string                  `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func objectPermissionsFromPb(pb *objectPermissionsPb) (*ObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ObjectPermissions{}
	st.AccessControlList = pb.AccessControlList
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

func partialUpdateToPb(st *PartialUpdate) (*partialUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partialUpdatePb{}
	pb.Id = st.Id
	pb.Operations = st.Operations
	pb.Schemas = st.Schemas

	return pb, nil
}

type partialUpdatePb struct {
	Id         string        `json:"-" url:"-"`
	Operations []Patch       `json:"Operations,omitempty"`
	Schemas    []PatchSchema `json:"schemas,omitempty"`
}

func partialUpdateFromPb(pb *partialUpdatePb) (*PartialUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartialUpdate{}
	st.Id = pb.Id
	st.Operations = pb.Operations
	st.Schemas = pb.Schemas

	return st, nil
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

type passwordAccessControlRequestPb struct {
	GroupName            string                  `json:"group_name,omitempty"`
	PermissionLevel      PasswordPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                  `json:"service_principal_name,omitempty"`
	UserName             string                  `json:"user_name,omitempty"`

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

func passwordAccessControlResponseToPb(st *PasswordAccessControlResponse) (*passwordAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type passwordAccessControlResponsePb struct {
	AllPermissions       []PasswordPermission `json:"all_permissions,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	GroupName            string               `json:"group_name,omitempty"`
	ServicePrincipalName string               `json:"service_principal_name,omitempty"`
	UserName             string               `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func passwordAccessControlResponseFromPb(pb *passwordAccessControlResponsePb) (*PasswordAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
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

type passwordPermissionPb struct {
	Inherited           bool                    `json:"inherited,omitempty"`
	InheritedFromObject []string                `json:"inherited_from_object,omitempty"`
	PermissionLevel     PasswordPermissionLevel `json:"permission_level,omitempty"`

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

func passwordPermissionsToPb(st *PasswordPermissions) (*passwordPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type passwordPermissionsPb struct {
	AccessControlList []PasswordAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                          `json:"object_id,omitempty"`
	ObjectType        string                          `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func passwordPermissionsFromPb(pb *passwordPermissionsPb) (*PasswordPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissions{}
	st.AccessControlList = pb.AccessControlList
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

type passwordPermissionsDescriptionPb struct {
	Description     string                  `json:"description,omitempty"`
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

func passwordPermissionsRequestToPb(st *PasswordPermissionsRequest) (*passwordPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &passwordPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList

	return pb, nil
}

type passwordPermissionsRequestPb struct {
	AccessControlList []PasswordAccessControlRequest `json:"access_control_list,omitempty"`
}

func passwordPermissionsRequestFromPb(pb *passwordPermissionsRequestPb) (*PasswordPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PasswordPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList

	return st, nil
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

type patchPb struct {
	Op    PatchOp `json:"op,omitempty"`
	Path  string  `json:"path,omitempty"`
	Value any     `json:"value,omitempty"`

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

func patchResponseToPb(st *PatchResponse) (*patchResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchResponsePb{}

	return pb, nil
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

type permissionPb struct {
	Inherited           bool            `json:"inherited,omitempty"`
	InheritedFromObject []string        `json:"inherited_from_object,omitempty"`
	PermissionLevel     PermissionLevel `json:"permission_level,omitempty"`

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

func permissionAssignmentToPb(st *PermissionAssignment) (*permissionAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionAssignmentPb{}
	pb.Error = st.Error
	pb.Permissions = st.Permissions
	pb.Principal = st.Principal

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type permissionAssignmentPb struct {
	Error       string                `json:"error,omitempty"`
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	Principal   *PrincipalOutput      `json:"principal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionAssignmentFromPb(pb *permissionAssignmentPb) (*PermissionAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionAssignment{}
	st.Error = pb.Error
	st.Permissions = pb.Permissions
	st.Principal = pb.Principal

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func permissionAssignmentsToPb(st *PermissionAssignments) (*permissionAssignmentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionAssignmentsPb{}
	pb.PermissionAssignments = st.PermissionAssignments

	return pb, nil
}

type permissionAssignmentsPb struct {
	PermissionAssignments []PermissionAssignment `json:"permission_assignments,omitempty"`
}

func permissionAssignmentsFromPb(pb *permissionAssignmentsPb) (*PermissionAssignments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionAssignments{}
	st.PermissionAssignments = pb.PermissionAssignments

	return st, nil
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

type permissionOutputPb struct {
	Description     string              `json:"description,omitempty"`
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

type permissionsDescriptionPb struct {
	Description     string          `json:"description,omitempty"`
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

type principalOutputPb struct {
	DisplayName          string `json:"display_name,omitempty"`
	GroupName            string `json:"group_name,omitempty"`
	PrincipalId          int64  `json:"principal_id,omitempty"`
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`

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

func resourceInfoToPb(st *ResourceInfo) (*resourceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resourceInfoPb{}
	pb.Id = st.Id
	pb.LegacyAclPath = st.LegacyAclPath
	pb.ParentResourceInfo = st.ParentResourceInfo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resourceInfoPb struct {
	Id                 string        `json:"id" url:"id"`
	LegacyAclPath      string        `json:"legacy_acl_path,omitempty" url:"legacy_acl_path,omitempty"`
	ParentResourceInfo *ResourceInfo `json:"parent_resource_info,omitempty" url:"parent_resource_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resourceInfoFromPb(pb *resourceInfoPb) (*ResourceInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResourceInfo{}
	st.Id = pb.Id
	st.LegacyAclPath = pb.LegacyAclPath
	st.ParentResourceInfo = pb.ParentResourceInfo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resourceInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resourceInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type resourceMetaPb struct {
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

func roleToPb(st *Role) (*rolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rolePb{}
	pb.Name = st.Name

	return pb, nil
}

type rolePb struct {
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

func ruleSetResponseToPb(st *RuleSetResponse) (*ruleSetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ruleSetResponsePb{}
	pb.Etag = st.Etag
	pb.GrantRules = st.GrantRules
	pb.Name = st.Name

	return pb, nil
}

type ruleSetResponsePb struct {
	Etag       string      `json:"etag"`
	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	Name       string      `json:"name"`
}

func ruleSetResponseFromPb(pb *ruleSetResponsePb) (*RuleSetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RuleSetResponse{}
	st.Etag = pb.Etag
	st.GrantRules = pb.GrantRules
	st.Name = pb.Name

	return st, nil
}

func ruleSetUpdateRequestToPb(st *RuleSetUpdateRequest) (*ruleSetUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ruleSetUpdateRequestPb{}
	pb.Etag = st.Etag
	pb.GrantRules = st.GrantRules
	pb.Name = st.Name

	return pb, nil
}

type ruleSetUpdateRequestPb struct {
	Etag       string      `json:"etag"`
	GrantRules []GrantRule `json:"grant_rules,omitempty"`
	Name       string      `json:"name"`
}

func ruleSetUpdateRequestFromPb(pb *ruleSetUpdateRequestPb) (*RuleSetUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RuleSetUpdateRequest{}
	st.Etag = pb.Etag
	st.GrantRules = pb.GrantRules
	st.Name = pb.Name

	return st, nil
}

func servicePrincipalToPb(st *ServicePrincipal) (*servicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servicePrincipalPb{}
	pb.Active = st.Active
	pb.ApplicationId = st.ApplicationId
	pb.DisplayName = st.DisplayName
	pb.Entitlements = st.Entitlements
	pb.ExternalId = st.ExternalId
	pb.Groups = st.Groups
	pb.Id = st.Id
	pb.Roles = st.Roles
	pb.Schemas = st.Schemas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servicePrincipalPb struct {
	Active        bool                     `json:"active,omitempty"`
	ApplicationId string                   `json:"applicationId,omitempty"`
	DisplayName   string                   `json:"displayName,omitempty"`
	Entitlements  []ComplexValue           `json:"entitlements,omitempty"`
	ExternalId    string                   `json:"externalId,omitempty"`
	Groups        []ComplexValue           `json:"groups,omitempty"`
	Id            string                   `json:"id,omitempty" url:"-"`
	Roles         []ComplexValue           `json:"roles,omitempty"`
	Schemas       []ServicePrincipalSchema `json:"schemas,omitempty"`

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
	st.Entitlements = pb.Entitlements
	st.ExternalId = pb.ExternalId
	st.Groups = pb.Groups
	st.Id = pb.Id
	st.Roles = pb.Roles
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

func setObjectPermissionsToPb(st *SetObjectPermissions) (*setObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setObjectPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.RequestObjectId = st.RequestObjectId
	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
}

type setObjectPermissionsPb struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`
	RequestObjectId   string                 `json:"-" url:"-"`
	RequestObjectType string                 `json:"-" url:"-"`
}

func setObjectPermissionsFromPb(pb *setObjectPermissionsPb) (*SetObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetObjectPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
}

func updateObjectPermissionsToPb(st *UpdateObjectPermissions) (*updateObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateObjectPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.RequestObjectId = st.RequestObjectId
	pb.RequestObjectType = st.RequestObjectType

	return pb, nil
}

type updateObjectPermissionsPb struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`
	RequestObjectId   string                 `json:"-" url:"-"`
	RequestObjectType string                 `json:"-" url:"-"`
}

func updateObjectPermissionsFromPb(pb *updateObjectPermissionsPb) (*UpdateObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateObjectPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.RequestObjectId = pb.RequestObjectId
	st.RequestObjectType = pb.RequestObjectType

	return st, nil
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
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

func updateRuleSetRequestToPb(st *UpdateRuleSetRequest) (*updateRuleSetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRuleSetRequestPb{}
	pb.Name = st.Name
	pb.RuleSet = st.RuleSet

	return pb, nil
}

type updateRuleSetRequestPb struct {
	Name    string               `json:"name"`
	RuleSet RuleSetUpdateRequest `json:"rule_set"`
}

func updateRuleSetRequestFromPb(pb *updateRuleSetRequestPb) (*UpdateRuleSetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRuleSetRequest{}
	st.Name = pb.Name
	st.RuleSet = pb.RuleSet

	return st, nil
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

type updateWorkspaceAssignmentsPb struct {
	Permissions []WorkspacePermission `json:"permissions,omitempty"`
	PrincipalId int64                 `json:"-" url:"-"`
	WorkspaceId int64                 `json:"-" url:"-"`
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

func userToPb(st *User) (*userPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &userPb{}
	pb.Active = st.Active
	pb.DisplayName = st.DisplayName
	pb.Emails = st.Emails
	pb.Entitlements = st.Entitlements
	pb.ExternalId = st.ExternalId
	pb.Groups = st.Groups
	pb.Id = st.Id
	pb.Name = st.Name
	pb.Roles = st.Roles
	pb.Schemas = st.Schemas
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type userPb struct {
	Active       bool           `json:"active,omitempty"`
	DisplayName  string         `json:"displayName,omitempty"`
	Emails       []ComplexValue `json:"emails,omitempty"`
	Entitlements []ComplexValue `json:"entitlements,omitempty"`
	ExternalId   string         `json:"externalId,omitempty"`
	Groups       []ComplexValue `json:"groups,omitempty"`
	Id           string         `json:"id,omitempty" url:"-"`
	Name         *Name          `json:"name,omitempty"`
	Roles        []ComplexValue `json:"roles,omitempty"`
	Schemas      []UserSchema   `json:"schemas,omitempty"`
	UserName     string         `json:"userName,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func userFromPb(pb *userPb) (*User, error) {
	if pb == nil {
		return nil, nil
	}
	st := &User{}
	st.Active = pb.Active
	st.DisplayName = pb.DisplayName
	st.Emails = pb.Emails
	st.Entitlements = pb.Entitlements
	st.ExternalId = pb.ExternalId
	st.Groups = pb.Groups
	st.Id = pb.Id
	st.Name = pb.Name
	st.Roles = pb.Roles
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

func workspacePermissionsToPb(st *WorkspacePermissions) (*workspacePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacePermissionsPb{}
	pb.Permissions = st.Permissions

	return pb, nil
}

type workspacePermissionsPb struct {
	Permissions []PermissionOutput `json:"permissions,omitempty"`
}

func workspacePermissionsFromPb(pb *workspacePermissionsPb) (*WorkspacePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspacePermissions{}
	st.Permissions = pb.Permissions

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
