// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func createFederationPolicyRequestToPb(st *CreateFederationPolicyRequest) (*createFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFederationPolicyRequestPb{}
	pb.Policy = st.Policy
	pb.RecipientName = st.RecipientName

	return pb, nil
}

type createFederationPolicyRequestPb struct {
	Policy        FederationPolicy `json:"policy"`
	RecipientName string           `json:"-" url:"-"`
}

func createFederationPolicyRequestFromPb(pb *createFederationPolicyRequestPb) (*CreateFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFederationPolicyRequest{}
	st.Policy = pb.Policy
	st.RecipientName = pb.RecipientName

	return st, nil
}

func createProviderToPb(st *CreateProvider) (*createProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createProviderPb{}
	pb.AuthenticationType = st.AuthenticationType
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.RecipientProfileStr = st.RecipientProfileStr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createProviderPb struct {
	AuthenticationType  AuthenticationType `json:"authentication_type"`
	Comment             string             `json:"comment,omitempty"`
	Name                string             `json:"name"`
	RecipientProfileStr string             `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createProviderFromPb(pb *createProviderPb) (*CreateProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProvider{}
	st.AuthenticationType = pb.AuthenticationType
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.RecipientProfileStr = pb.RecipientProfileStr

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createProviderPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createProviderPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createRecipientToPb(st *CreateRecipient) (*createRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRecipientPb{}
	pb.AuthenticationType = st.AuthenticationType
	pb.Comment = st.Comment
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.ExpirationTime = st.ExpirationTime
	pb.IpAccessList = st.IpAccessList
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.PropertiesKvpairs = st.PropertiesKvpairs
	pb.SharingCode = st.SharingCode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createRecipientPb struct {
	AuthenticationType             AuthenticationType          `json:"authentication_type"`
	Comment                        string                      `json:"comment,omitempty"`
	DataRecipientGlobalMetastoreId string                      `json:"data_recipient_global_metastore_id,omitempty"`
	ExpirationTime                 int64                       `json:"expiration_time,omitempty"`
	IpAccessList                   *IpAccessList               `json:"ip_access_list,omitempty"`
	Name                           string                      `json:"name"`
	Owner                          string                      `json:"owner,omitempty"`
	PropertiesKvpairs              *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	SharingCode                    string                      `json:"sharing_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRecipientFromPb(pb *createRecipientPb) (*CreateRecipient, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRecipient{}
	st.AuthenticationType = pb.AuthenticationType
	st.Comment = pb.Comment
	st.DataRecipientGlobalMetastoreId = pb.DataRecipientGlobalMetastoreId
	st.ExpirationTime = pb.ExpirationTime
	st.IpAccessList = pb.IpAccessList
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PropertiesKvpairs = pb.PropertiesKvpairs
	st.SharingCode = pb.SharingCode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRecipientPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRecipientPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createShareToPb(st *CreateShare) (*createSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSharePb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createSharePb struct {
	Comment     string `json:"comment,omitempty"`
	Name        string `json:"name"`
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createShareFromPb(pb *createSharePb) (*CreateShare, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateShare{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteFederationPolicyRequestToPb(st *DeleteFederationPolicyRequest) (*deleteFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFederationPolicyRequestPb{}
	pb.Name = st.Name
	pb.RecipientName = st.RecipientName

	return pb, nil
}

type deleteFederationPolicyRequestPb struct {
	Name          string `json:"-" url:"-"`
	RecipientName string `json:"-" url:"-"`
}

func deleteFederationPolicyRequestFromPb(pb *deleteFederationPolicyRequestPb) (*DeleteFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFederationPolicyRequest{}
	st.Name = pb.Name
	st.RecipientName = pb.RecipientName

	return st, nil
}

func deleteProviderRequestToPb(st *DeleteProviderRequest) (*deleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteProviderRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteProviderRequestFromPb(pb *deleteProviderRequestPb) (*DeleteProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteProviderRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteRecipientRequestToPb(st *DeleteRecipientRequest) (*deleteRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRecipientRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteRecipientRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteRecipientRequestFromPb(pb *deleteRecipientRequestPb) (*DeleteRecipientRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRecipientRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteShareRequestToPb(st *DeleteShareRequest) (*deleteShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteShareRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteShareRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteShareRequestFromPb(pb *deleteShareRequestPb) (*DeleteShareRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteShareRequest{}
	st.Name = pb.Name

	return st, nil
}

func deltaSharingDependencyToPb(st *DeltaSharingDependency) (*deltaSharingDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingDependencyPb{}
	pb.Function = st.Function
	pb.Table = st.Table

	return pb, nil
}

type deltaSharingDependencyPb struct {
	Function *DeltaSharingFunctionDependency `json:"function,omitempty"`
	Table    *DeltaSharingTableDependency    `json:"table,omitempty"`
}

func deltaSharingDependencyFromPb(pb *deltaSharingDependencyPb) (*DeltaSharingDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingDependency{}
	st.Function = pb.Function
	st.Table = pb.Table

	return st, nil
}

func deltaSharingDependencyListToPb(st *DeltaSharingDependencyList) (*deltaSharingDependencyListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingDependencyListPb{}
	pb.Dependencies = st.Dependencies

	return pb, nil
}

type deltaSharingDependencyListPb struct {
	Dependencies []DeltaSharingDependency `json:"dependencies,omitempty"`
}

func deltaSharingDependencyListFromPb(pb *deltaSharingDependencyListPb) (*DeltaSharingDependencyList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingDependencyList{}
	st.Dependencies = pb.Dependencies

	return st, nil
}

func deltaSharingFunctionToPb(st *DeltaSharingFunction) (*deltaSharingFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingFunctionPb{}
	pb.Aliases = st.Aliases
	pb.Comment = st.Comment
	pb.DataType = st.DataType
	pb.DependencyList = st.DependencyList
	pb.FullDataType = st.FullDataType
	pb.Id = st.Id
	pb.InputParams = st.InputParams
	pb.Name = st.Name
	pb.Properties = st.Properties
	pb.RoutineDefinition = st.RoutineDefinition
	pb.Schema = st.Schema
	pb.SecurableKind = st.SecurableKind
	pb.Share = st.Share
	pb.ShareId = st.ShareId
	pb.StorageLocation = st.StorageLocation
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deltaSharingFunctionPb struct {
	Aliases           []RegisteredModelAlias      `json:"aliases,omitempty"`
	Comment           string                      `json:"comment,omitempty"`
	DataType          ColumnTypeName              `json:"data_type,omitempty"`
	DependencyList    *DeltaSharingDependencyList `json:"dependency_list,omitempty"`
	FullDataType      string                      `json:"full_data_type,omitempty"`
	Id                string                      `json:"id,omitempty"`
	InputParams       *FunctionParameterInfos     `json:"input_params,omitempty"`
	Name              string                      `json:"name,omitempty"`
	Properties        string                      `json:"properties,omitempty"`
	RoutineDefinition string                      `json:"routine_definition,omitempty"`
	Schema            string                      `json:"schema,omitempty"`
	SecurableKind     SharedSecurableKind         `json:"securable_kind,omitempty"`
	Share             string                      `json:"share,omitempty"`
	ShareId           string                      `json:"share_id,omitempty"`
	StorageLocation   string                      `json:"storage_location,omitempty"`
	Tags              []catalog.TagKeyValue       `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSharingFunctionFromPb(pb *deltaSharingFunctionPb) (*DeltaSharingFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingFunction{}
	st.Aliases = pb.Aliases
	st.Comment = pb.Comment
	st.DataType = pb.DataType
	st.DependencyList = pb.DependencyList
	st.FullDataType = pb.FullDataType
	st.Id = pb.Id
	st.InputParams = pb.InputParams
	st.Name = pb.Name
	st.Properties = pb.Properties
	st.RoutineDefinition = pb.RoutineDefinition
	st.Schema = pb.Schema
	st.SecurableKind = pb.SecurableKind
	st.Share = pb.Share
	st.ShareId = pb.ShareId
	st.StorageLocation = pb.StorageLocation
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deltaSharingFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaSharingFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deltaSharingFunctionDependencyToPb(st *DeltaSharingFunctionDependency) (*deltaSharingFunctionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingFunctionDependencyPb{}
	pb.FunctionName = st.FunctionName
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deltaSharingFunctionDependencyPb struct {
	FunctionName string `json:"function_name,omitempty"`
	SchemaName   string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSharingFunctionDependencyFromPb(pb *deltaSharingFunctionDependencyPb) (*DeltaSharingFunctionDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingFunctionDependency{}
	st.FunctionName = pb.FunctionName
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deltaSharingFunctionDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaSharingFunctionDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deltaSharingTableDependencyToPb(st *DeltaSharingTableDependency) (*deltaSharingTableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingTableDependencyPb{}
	pb.SchemaName = st.SchemaName
	pb.TableName = st.TableName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deltaSharingTableDependencyPb struct {
	SchemaName string `json:"schema_name,omitempty"`
	TableName  string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSharingTableDependencyFromPb(pb *deltaSharingTableDependencyPb) (*DeltaSharingTableDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingTableDependency{}
	st.SchemaName = pb.SchemaName
	st.TableName = pb.TableName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deltaSharingTableDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaSharingTableDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func federationPolicyToPb(st *FederationPolicy) (*federationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &federationPolicyPb{}
	pb.Comment = st.Comment
	pb.CreateTime = st.CreateTime
	pb.Id = st.Id
	pb.Name = st.Name
	pb.OidcPolicy = st.OidcPolicy
	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type federationPolicyPb struct {
	Comment    string                `json:"comment,omitempty"`
	CreateTime string                `json:"create_time,omitempty"`
	Id         string                `json:"id,omitempty"`
	Name       string                `json:"name,omitempty"`
	OidcPolicy *OidcFederationPolicy `json:"oidc_policy,omitempty"`
	UpdateTime string                `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func federationPolicyFromPb(pb *federationPolicyPb) (*FederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FederationPolicy{}
	st.Comment = pb.Comment
	st.CreateTime = pb.CreateTime
	st.Id = pb.Id
	st.Name = pb.Name
	st.OidcPolicy = pb.OidcPolicy
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *federationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st federationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func functionParameterInfoToPb(st *FunctionParameterInfo) (*functionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfoPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.ParameterDefault = st.ParameterDefault
	pb.ParameterMode = st.ParameterMode
	pb.ParameterType = st.ParameterType
	pb.Position = st.Position
	pb.TypeIntervalType = st.TypeIntervalType
	pb.TypeJson = st.TypeJson
	pb.TypeName = st.TypeName
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type functionParameterInfoPb struct {
	Comment          string                `json:"comment,omitempty"`
	Name             string                `json:"name,omitempty"`
	ParameterDefault string                `json:"parameter_default,omitempty"`
	ParameterMode    FunctionParameterMode `json:"parameter_mode,omitempty"`
	ParameterType    FunctionParameterType `json:"parameter_type,omitempty"`
	Position         int                   `json:"position,omitempty"`
	TypeIntervalType string                `json:"type_interval_type,omitempty"`
	TypeJson         string                `json:"type_json,omitempty"`
	TypeName         ColumnTypeName        `json:"type_name,omitempty"`
	TypePrecision    int                   `json:"type_precision,omitempty"`
	TypeScale        int                   `json:"type_scale,omitempty"`
	TypeText         string                `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func functionParameterInfoFromPb(pb *functionParameterInfoPb) (*FunctionParameterInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfo{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.ParameterDefault = pb.ParameterDefault
	st.ParameterMode = pb.ParameterMode
	st.ParameterType = pb.ParameterType
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	st.TypeName = pb.TypeName
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *functionParameterInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st functionParameterInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func functionParameterInfosToPb(st *FunctionParameterInfos) (*functionParameterInfosPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfosPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

type functionParameterInfosPb struct {
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func functionParameterInfosFromPb(pb *functionParameterInfosPb) (*FunctionParameterInfos, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfos{}
	st.Parameters = pb.Parameters

	return st, nil
}

func getActivationUrlInfoRequestToPb(st *GetActivationUrlInfoRequest) (*getActivationUrlInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getActivationUrlInfoRequestPb{}
	pb.ActivationUrl = st.ActivationUrl

	return pb, nil
}

type getActivationUrlInfoRequestPb struct {
	ActivationUrl string `json:"-" url:"-"`
}

func getActivationUrlInfoRequestFromPb(pb *getActivationUrlInfoRequestPb) (*GetActivationUrlInfoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetActivationUrlInfoRequest{}
	st.ActivationUrl = pb.ActivationUrl

	return st, nil
}

func getFederationPolicyRequestToPb(st *GetFederationPolicyRequest) (*getFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFederationPolicyRequestPb{}
	pb.Name = st.Name
	pb.RecipientName = st.RecipientName

	return pb, nil
}

type getFederationPolicyRequestPb struct {
	Name          string `json:"-" url:"-"`
	RecipientName string `json:"-" url:"-"`
}

func getFederationPolicyRequestFromPb(pb *getFederationPolicyRequestPb) (*GetFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFederationPolicyRequest{}
	st.Name = pb.Name
	st.RecipientName = pb.RecipientName

	return st, nil
}

func getProviderRequestToPb(st *GetProviderRequest) (*getProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getProviderRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getProviderRequestFromPb(pb *getProviderRequestPb) (*GetProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderRequest{}
	st.Name = pb.Name

	return st, nil
}

func getRecipientRequestToPb(st *GetRecipientRequest) (*getRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRecipientRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getRecipientRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getRecipientRequestFromPb(pb *getRecipientRequestPb) (*GetRecipientRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRecipientRequest{}
	st.Name = pb.Name

	return st, nil
}

func getRecipientSharePermissionsResponseToPb(st *GetRecipientSharePermissionsResponse) (*getRecipientSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRecipientSharePermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PermissionsOut = st.PermissionsOut

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getRecipientSharePermissionsResponsePb struct {
	NextPageToken  string                       `json:"next_page_token,omitempty"`
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRecipientSharePermissionsResponseFromPb(pb *getRecipientSharePermissionsResponsePb) (*GetRecipientSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRecipientSharePermissionsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PermissionsOut = pb.PermissionsOut

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRecipientSharePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRecipientSharePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getSharePermissionsResponseToPb(st *GetSharePermissionsResponse) (*getSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSharePermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PrivilegeAssignments = st.PrivilegeAssignments

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getSharePermissionsResponsePb struct {
	NextPageToken        string                `json:"next_page_token,omitempty"`
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getSharePermissionsResponseFromPb(pb *getSharePermissionsResponsePb) (*GetSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSharePermissionsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PrivilegeAssignments = pb.PrivilegeAssignments

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getSharePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getSharePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getShareRequestToPb(st *GetShareRequest) (*getShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getShareRequestPb{}
	pb.IncludeSharedData = st.IncludeSharedData
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getShareRequestPb struct {
	IncludeSharedData bool   `json:"-" url:"include_shared_data,omitempty"`
	Name              string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getShareRequestFromPb(pb *getShareRequestPb) (*GetShareRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetShareRequest{}
	st.IncludeSharedData = pb.IncludeSharedData
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getShareRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getShareRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func ipAccessListToPb(st *IpAccessList) (*ipAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ipAccessListPb{}
	pb.AllowedIpAddresses = st.AllowedIpAddresses

	return pb, nil
}

type ipAccessListPb struct {
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

func ipAccessListFromPb(pb *ipAccessListPb) (*IpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IpAccessList{}
	st.AllowedIpAddresses = pb.AllowedIpAddresses

	return st, nil
}

func listFederationPoliciesRequestToPb(st *ListFederationPoliciesRequest) (*listFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFederationPoliciesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.RecipientName = st.RecipientName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFederationPoliciesRequestPb struct {
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	RecipientName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFederationPoliciesRequestFromPb(pb *listFederationPoliciesRequestPb) (*ListFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFederationPoliciesRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.RecipientName = pb.RecipientName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFederationPoliciesResponseToPb(st *ListFederationPoliciesResponse) (*listFederationPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFederationPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Policies = st.Policies

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFederationPoliciesResponsePb struct {
	NextPageToken string             `json:"next_page_token,omitempty"`
	Policies      []FederationPolicy `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFederationPoliciesResponseFromPb(pb *listFederationPoliciesResponsePb) (*ListFederationPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFederationPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Policies = pb.Policies

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFederationPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFederationPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProviderShareAssetsRequestToPb(st *ListProviderShareAssetsRequest) (*listProviderShareAssetsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderShareAssetsRequestPb{}
	pb.FunctionMaxResults = st.FunctionMaxResults
	pb.NotebookMaxResults = st.NotebookMaxResults
	pb.ProviderName = st.ProviderName
	pb.ShareName = st.ShareName
	pb.TableMaxResults = st.TableMaxResults
	pb.VolumeMaxResults = st.VolumeMaxResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProviderShareAssetsRequestPb struct {
	FunctionMaxResults int    `json:"-" url:"function_max_results,omitempty"`
	NotebookMaxResults int    `json:"-" url:"notebook_max_results,omitempty"`
	ProviderName       string `json:"-" url:"-"`
	ShareName          string `json:"-" url:"-"`
	TableMaxResults    int    `json:"-" url:"table_max_results,omitempty"`
	VolumeMaxResults   int    `json:"-" url:"volume_max_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProviderShareAssetsRequestFromPb(pb *listProviderShareAssetsRequestPb) (*ListProviderShareAssetsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderShareAssetsRequest{}
	st.FunctionMaxResults = pb.FunctionMaxResults
	st.NotebookMaxResults = pb.NotebookMaxResults
	st.ProviderName = pb.ProviderName
	st.ShareName = pb.ShareName
	st.TableMaxResults = pb.TableMaxResults
	st.VolumeMaxResults = pb.VolumeMaxResults

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProviderShareAssetsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProviderShareAssetsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProviderShareAssetsResponseToPb(st *ListProviderShareAssetsResponse) (*listProviderShareAssetsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderShareAssetsResponsePb{}
	pb.Functions = st.Functions
	pb.Notebooks = st.Notebooks
	pb.Tables = st.Tables
	pb.Volumes = st.Volumes

	return pb, nil
}

type listProviderShareAssetsResponsePb struct {
	Functions []DeltaSharingFunction `json:"functions,omitempty"`
	Notebooks []NotebookFile         `json:"notebooks,omitempty"`
	Tables    []Table                `json:"tables,omitempty"`
	Volumes   []Volume               `json:"volumes,omitempty"`
}

func listProviderShareAssetsResponseFromPb(pb *listProviderShareAssetsResponsePb) (*ListProviderShareAssetsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderShareAssetsResponse{}
	st.Functions = pb.Functions
	st.Notebooks = pb.Notebooks
	st.Tables = pb.Tables
	st.Volumes = pb.Volumes

	return st, nil
}

func listProviderSharesResponseToPb(st *ListProviderSharesResponse) (*listProviderSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderSharesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Shares = st.Shares

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProviderSharesResponsePb struct {
	NextPageToken string          `json:"next_page_token,omitempty"`
	Shares        []ProviderShare `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProviderSharesResponseFromPb(pb *listProviderSharesResponsePb) (*ListProviderSharesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderSharesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Shares = pb.Shares

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProviderSharesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProviderSharesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProvidersRequestToPb(st *ListProvidersRequest) (*listProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersRequestPb{}
	pb.DataProviderGlobalMetastoreId = st.DataProviderGlobalMetastoreId
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProvidersRequestPb struct {
	DataProviderGlobalMetastoreId string `json:"-" url:"data_provider_global_metastore_id,omitempty"`
	MaxResults                    int    `json:"-" url:"max_results,omitempty"`
	PageToken                     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProvidersRequestFromPb(pb *listProvidersRequestPb) (*ListProvidersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersRequest{}
	st.DataProviderGlobalMetastoreId = pb.DataProviderGlobalMetastoreId
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProvidersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProvidersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listProvidersResponseToPb(st *ListProvidersResponse) (*listProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Providers = st.Providers

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listProvidersResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Providers     []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProvidersResponseFromPb(pb *listProvidersResponsePb) (*ListProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Providers = pb.Providers

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listProvidersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProvidersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRecipientsRequestToPb(st *ListRecipientsRequest) (*listRecipientsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRecipientsRequestPb{}
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRecipientsRequestPb struct {
	DataRecipientGlobalMetastoreId string `json:"-" url:"data_recipient_global_metastore_id,omitempty"`
	MaxResults                     int    `json:"-" url:"max_results,omitempty"`
	PageToken                      string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRecipientsRequestFromPb(pb *listRecipientsRequestPb) (*ListRecipientsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRecipientsRequest{}
	st.DataRecipientGlobalMetastoreId = pb.DataRecipientGlobalMetastoreId
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRecipientsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRecipientsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRecipientsResponseToPb(st *ListRecipientsResponse) (*listRecipientsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRecipientsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Recipients = st.Recipients

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRecipientsResponsePb struct {
	NextPageToken string          `json:"next_page_token,omitempty"`
	Recipients    []RecipientInfo `json:"recipients,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRecipientsResponseFromPb(pb *listRecipientsResponsePb) (*ListRecipientsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRecipientsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Recipients = pb.Recipients

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRecipientsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRecipientsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSharesRequestToPb(st *ListSharesRequest) (*listSharesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSharesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.Name = st.Name
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSharesRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	Name       string `json:"-" url:"-"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSharesRequestFromPb(pb *listSharesRequestPb) (*ListSharesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSharesRequest{}
	st.MaxResults = pb.MaxResults
	st.Name = pb.Name
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSharesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSharesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSharesResponseToPb(st *ListSharesResponse) (*listSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSharesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Shares = st.Shares

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSharesResponsePb struct {
	NextPageToken string      `json:"next_page_token,omitempty"`
	Shares        []ShareInfo `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSharesResponseFromPb(pb *listSharesResponsePb) (*ListSharesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSharesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Shares = pb.Shares

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSharesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSharesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func notebookFileToPb(st *NotebookFile) (*notebookFilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookFilePb{}
	pb.Comment = st.Comment
	pb.Id = st.Id
	pb.Name = st.Name
	pb.Share = st.Share
	pb.ShareId = st.ShareId
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type notebookFilePb struct {
	Comment string                `json:"comment,omitempty"`
	Id      string                `json:"id,omitempty"`
	Name    string                `json:"name,omitempty"`
	Share   string                `json:"share,omitempty"`
	ShareId string                `json:"share_id,omitempty"`
	Tags    []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookFileFromPb(pb *notebookFilePb) (*NotebookFile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookFile{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	st.Name = pb.Name
	st.Share = pb.Share
	st.ShareId = pb.ShareId
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notebookFilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookFilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func oidcFederationPolicyToPb(st *OidcFederationPolicy) (*oidcFederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oidcFederationPolicyPb{}
	pb.Audiences = st.Audiences
	pb.Issuer = st.Issuer
	pb.Subject = st.Subject
	pb.SubjectClaim = st.SubjectClaim

	return pb, nil
}

type oidcFederationPolicyPb struct {
	Audiences    []string `json:"audiences,omitempty"`
	Issuer       string   `json:"issuer"`
	Subject      string   `json:"subject"`
	SubjectClaim string   `json:"subject_claim"`
}

func oidcFederationPolicyFromPb(pb *oidcFederationPolicyPb) (*OidcFederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OidcFederationPolicy{}
	st.Audiences = pb.Audiences
	st.Issuer = pb.Issuer
	st.Subject = pb.Subject
	st.SubjectClaim = pb.SubjectClaim

	return st, nil
}

func partitionToPb(st *Partition) (*partitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partitionPb{}
	pb.Values = st.Values

	return pb, nil
}

type partitionPb struct {
	Values []PartitionValue `json:"values,omitempty"`
}

func partitionFromPb(pb *partitionPb) (*Partition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Partition{}
	st.Values = pb.Values

	return st, nil
}

func partitionValueToPb(st *PartitionValue) (*partitionValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partitionValuePb{}
	pb.Name = st.Name
	pb.Op = st.Op
	pb.RecipientPropertyKey = st.RecipientPropertyKey
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type partitionValuePb struct {
	Name                 string           `json:"name,omitempty"`
	Op                   PartitionValueOp `json:"op,omitempty"`
	RecipientPropertyKey string           `json:"recipient_property_key,omitempty"`
	Value                string           `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func partitionValueFromPb(pb *partitionValuePb) (*PartitionValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartitionValue{}
	st.Name = pb.Name
	st.Op = pb.Op
	st.RecipientPropertyKey = pb.RecipientPropertyKey
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *partitionValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st partitionValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func permissionsChangeToPb(st *PermissionsChange) (*permissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsChangePb{}
	pb.Add = st.Add
	pb.Principal = st.Principal
	pb.Remove = st.Remove

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type permissionsChangePb struct {
	Add       []string `json:"add,omitempty"`
	Principal string   `json:"principal,omitempty"`
	Remove    []string `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionsChangeFromPb(pb *permissionsChangePb) (*PermissionsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsChange{}
	st.Add = pb.Add
	st.Principal = pb.Principal
	st.Remove = pb.Remove

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func privilegeAssignmentToPb(st *PrivilegeAssignment) (*privilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privilegeAssignmentPb{}
	pb.Principal = st.Principal
	pb.Privileges = st.Privileges

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type privilegeAssignmentPb struct {
	Principal  string      `json:"principal,omitempty"`
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func privilegeAssignmentFromPb(pb *privilegeAssignmentPb) (*PrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivilegeAssignment{}
	st.Principal = pb.Principal
	st.Privileges = pb.Privileges

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *privilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st privilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func providerInfoToPb(st *ProviderInfo) (*providerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerInfoPb{}
	pb.AuthenticationType = st.AuthenticationType
	pb.Cloud = st.Cloud
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DataProviderGlobalMetastoreId = st.DataProviderGlobalMetastoreId
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.RecipientProfile = st.RecipientProfile
	pb.RecipientProfileStr = st.RecipientProfileStr
	pb.Region = st.Region
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type providerInfoPb struct {
	AuthenticationType            AuthenticationType `json:"authentication_type,omitempty"`
	Cloud                         string             `json:"cloud,omitempty"`
	Comment                       string             `json:"comment,omitempty"`
	CreatedAt                     int64              `json:"created_at,omitempty"`
	CreatedBy                     string             `json:"created_by,omitempty"`
	DataProviderGlobalMetastoreId string             `json:"data_provider_global_metastore_id,omitempty"`
	MetastoreId                   string             `json:"metastore_id,omitempty"`
	Name                          string             `json:"name,omitempty"`
	Owner                         string             `json:"owner,omitempty"`
	RecipientProfile              *RecipientProfile  `json:"recipient_profile,omitempty"`
	RecipientProfileStr           string             `json:"recipient_profile_str,omitempty"`
	Region                        string             `json:"region,omitempty"`
	UpdatedAt                     int64              `json:"updated_at,omitempty"`
	UpdatedBy                     string             `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func providerInfoFromPb(pb *providerInfoPb) (*ProviderInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderInfo{}
	st.AuthenticationType = pb.AuthenticationType
	st.Cloud = pb.Cloud
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataProviderGlobalMetastoreId = pb.DataProviderGlobalMetastoreId
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.RecipientProfile = pb.RecipientProfile
	st.RecipientProfileStr = pb.RecipientProfileStr
	st.Region = pb.Region
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *providerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st providerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func providerShareToPb(st *ProviderShare) (*providerSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerSharePb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type providerSharePb struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func providerShareFromPb(pb *providerSharePb) (*ProviderShare, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderShare{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *providerSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st providerSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func recipientInfoToPb(st *RecipientInfo) (*recipientInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &recipientInfoPb{}
	pb.Activated = st.Activated
	pb.ActivationUrl = st.ActivationUrl
	pb.AuthenticationType = st.AuthenticationType
	pb.Cloud = st.Cloud
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.ExpirationTime = st.ExpirationTime
	pb.IpAccessList = st.IpAccessList
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.PropertiesKvpairs = st.PropertiesKvpairs
	pb.Region = st.Region
	pb.SharingCode = st.SharingCode
	pb.Tokens = st.Tokens
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type recipientInfoPb struct {
	Activated                      bool                        `json:"activated,omitempty"`
	ActivationUrl                  string                      `json:"activation_url,omitempty"`
	AuthenticationType             AuthenticationType          `json:"authentication_type,omitempty"`
	Cloud                          string                      `json:"cloud,omitempty"`
	Comment                        string                      `json:"comment,omitempty"`
	CreatedAt                      int64                       `json:"created_at,omitempty"`
	CreatedBy                      string                      `json:"created_by,omitempty"`
	DataRecipientGlobalMetastoreId string                      `json:"data_recipient_global_metastore_id,omitempty"`
	ExpirationTime                 int64                       `json:"expiration_time,omitempty"`
	IpAccessList                   *IpAccessList               `json:"ip_access_list,omitempty"`
	MetastoreId                    string                      `json:"metastore_id,omitempty"`
	Name                           string                      `json:"name,omitempty"`
	Owner                          string                      `json:"owner,omitempty"`
	PropertiesKvpairs              *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	Region                         string                      `json:"region,omitempty"`
	SharingCode                    string                      `json:"sharing_code,omitempty"`
	Tokens                         []RecipientTokenInfo        `json:"tokens,omitempty"`
	UpdatedAt                      int64                       `json:"updated_at,omitempty"`
	UpdatedBy                      string                      `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func recipientInfoFromPb(pb *recipientInfoPb) (*RecipientInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RecipientInfo{}
	st.Activated = pb.Activated
	st.ActivationUrl = pb.ActivationUrl
	st.AuthenticationType = pb.AuthenticationType
	st.Cloud = pb.Cloud
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataRecipientGlobalMetastoreId = pb.DataRecipientGlobalMetastoreId
	st.ExpirationTime = pb.ExpirationTime
	st.IpAccessList = pb.IpAccessList
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PropertiesKvpairs = pb.PropertiesKvpairs
	st.Region = pb.Region
	st.SharingCode = pb.SharingCode
	st.Tokens = pb.Tokens
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *recipientInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st recipientInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func recipientProfileToPb(st *RecipientProfile) (*recipientProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &recipientProfilePb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type recipientProfilePb struct {
	BearerToken             string `json:"bearer_token,omitempty"`
	Endpoint                string `json:"endpoint,omitempty"`
	ShareCredentialsVersion int    `json:"share_credentials_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func recipientProfileFromPb(pb *recipientProfilePb) (*RecipientProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RecipientProfile{}
	st.BearerToken = pb.BearerToken
	st.Endpoint = pb.Endpoint
	st.ShareCredentialsVersion = pb.ShareCredentialsVersion

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *recipientProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st recipientProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func recipientTokenInfoToPb(st *RecipientTokenInfo) (*recipientTokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &recipientTokenInfoPb{}
	pb.ActivationUrl = st.ActivationUrl
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.ExpirationTime = st.ExpirationTime
	pb.Id = st.Id
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type recipientTokenInfoPb struct {
	ActivationUrl  string `json:"activation_url,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
	ExpirationTime int64  `json:"expiration_time,omitempty"`
	Id             string `json:"id,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	UpdatedBy      string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func recipientTokenInfoFromPb(pb *recipientTokenInfoPb) (*RecipientTokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RecipientTokenInfo{}
	st.ActivationUrl = pb.ActivationUrl
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.ExpirationTime = pb.ExpirationTime
	st.Id = pb.Id
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *recipientTokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st recipientTokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelAliasToPb(st *RegisteredModelAlias) (*registeredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAliasPb{}
	pb.AliasName = st.AliasName
	pb.VersionNum = st.VersionNum

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelAliasPb struct {
	AliasName  string `json:"alias_name,omitempty"`
	VersionNum int64  `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelAliasFromPb(pb *registeredModelAliasPb) (*RegisteredModelAlias, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAlias{}
	st.AliasName = pb.AliasName
	st.VersionNum = pb.VersionNum

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAliasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAliasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func retrieveTokenRequestToPb(st *RetrieveTokenRequest) (*retrieveTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &retrieveTokenRequestPb{}
	pb.ActivationUrl = st.ActivationUrl

	return pb, nil
}

type retrieveTokenRequestPb struct {
	ActivationUrl string `json:"-" url:"-"`
}

func retrieveTokenRequestFromPb(pb *retrieveTokenRequestPb) (*RetrieveTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RetrieveTokenRequest{}
	st.ActivationUrl = pb.ActivationUrl

	return st, nil
}

func retrieveTokenResponseToPb(st *RetrieveTokenResponse) (*retrieveTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &retrieveTokenResponsePb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ExpirationTime = st.ExpirationTime
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type retrieveTokenResponsePb struct {
	BearerToken             string `json:"bearerToken,omitempty"`
	Endpoint                string `json:"endpoint,omitempty"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	ShareCredentialsVersion int    `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func retrieveTokenResponseFromPb(pb *retrieveTokenResponsePb) (*RetrieveTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RetrieveTokenResponse{}
	st.BearerToken = pb.BearerToken
	st.Endpoint = pb.Endpoint
	st.ExpirationTime = pb.ExpirationTime
	st.ShareCredentialsVersion = pb.ShareCredentialsVersion

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *retrieveTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st retrieveTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func rotateRecipientTokenToPb(st *RotateRecipientToken) (*rotateRecipientTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rotateRecipientTokenPb{}
	pb.ExistingTokenExpireInSeconds = st.ExistingTokenExpireInSeconds
	pb.Name = st.Name

	return pb, nil
}

type rotateRecipientTokenPb struct {
	ExistingTokenExpireInSeconds int64  `json:"existing_token_expire_in_seconds"`
	Name                         string `json:"-" url:"-"`
}

func rotateRecipientTokenFromPb(pb *rotateRecipientTokenPb) (*RotateRecipientToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RotateRecipientToken{}
	st.ExistingTokenExpireInSeconds = pb.ExistingTokenExpireInSeconds
	st.Name = pb.Name

	return st, nil
}

func securablePropertiesKvPairsToPb(st *SecurablePropertiesKvPairs) (*securablePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &securablePropertiesKvPairsPb{}
	pb.Properties = st.Properties

	return pb, nil
}

type securablePropertiesKvPairsPb struct {
	Properties map[string]string `json:"properties"`
}

func securablePropertiesKvPairsFromPb(pb *securablePropertiesKvPairsPb) (*SecurablePropertiesKvPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecurablePropertiesKvPairs{}
	st.Properties = pb.Properties

	return st, nil
}

func shareInfoToPb(st *ShareInfo) (*shareInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &shareInfoPb{}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Name = st.Name
	pb.Objects = st.Objects
	pb.Owner = st.Owner
	pb.StorageLocation = st.StorageLocation
	pb.StorageRoot = st.StorageRoot
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type shareInfoPb struct {
	Comment         string             `json:"comment,omitempty"`
	CreatedAt       int64              `json:"created_at,omitempty"`
	CreatedBy       string             `json:"created_by,omitempty"`
	Name            string             `json:"name,omitempty"`
	Objects         []SharedDataObject `json:"objects,omitempty"`
	Owner           string             `json:"owner,omitempty"`
	StorageLocation string             `json:"storage_location,omitempty"`
	StorageRoot     string             `json:"storage_root,omitempty"`
	UpdatedAt       int64              `json:"updated_at,omitempty"`
	UpdatedBy       string             `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func shareInfoFromPb(pb *shareInfoPb) (*ShareInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareInfo{}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Name = pb.Name
	st.Objects = pb.Objects
	st.Owner = pb.Owner
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *shareInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st shareInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sharePermissionsRequestToPb(st *SharePermissionsRequest) (*sharePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharePermissionsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.Name = st.Name
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sharePermissionsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	Name       string `json:"-" url:"-"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sharePermissionsRequestFromPb(pb *sharePermissionsRequestPb) (*SharePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharePermissionsRequest{}
	st.MaxResults = pb.MaxResults
	st.Name = pb.Name
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sharePermissionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sharePermissionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func shareToPrivilegeAssignmentToPb(st *ShareToPrivilegeAssignment) (*shareToPrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &shareToPrivilegeAssignmentPb{}
	pb.PrivilegeAssignments = st.PrivilegeAssignments
	pb.ShareName = st.ShareName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type shareToPrivilegeAssignmentPb struct {
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	ShareName            string                `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func shareToPrivilegeAssignmentFromPb(pb *shareToPrivilegeAssignmentPb) (*ShareToPrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareToPrivilegeAssignment{}
	st.PrivilegeAssignments = pb.PrivilegeAssignments
	st.ShareName = pb.ShareName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *shareToPrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st shareToPrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sharedDataObjectToPb(st *SharedDataObject) (*sharedDataObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharedDataObjectPb{}
	pb.AddedAt = st.AddedAt
	pb.AddedBy = st.AddedBy
	pb.CdfEnabled = st.CdfEnabled
	pb.Comment = st.Comment
	pb.Content = st.Content
	pb.DataObjectType = st.DataObjectType
	pb.HistoryDataSharingStatus = st.HistoryDataSharingStatus
	pb.Name = st.Name
	pb.Partitions = st.Partitions
	pb.SharedAs = st.SharedAs
	pb.StartVersion = st.StartVersion
	pb.Status = st.Status
	pb.StringSharedAs = st.StringSharedAs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sharedDataObjectPb struct {
	AddedAt                  int64                                    `json:"added_at,omitempty"`
	AddedBy                  string                                   `json:"added_by,omitempty"`
	CdfEnabled               bool                                     `json:"cdf_enabled,omitempty"`
	Comment                  string                                   `json:"comment,omitempty"`
	Content                  string                                   `json:"content,omitempty"`
	DataObjectType           SharedDataObjectDataObjectType           `json:"data_object_type,omitempty"`
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus `json:"history_data_sharing_status,omitempty"`
	Name                     string                                   `json:"name"`
	Partitions               []Partition                              `json:"partitions,omitempty"`
	SharedAs                 string                                   `json:"shared_as,omitempty"`
	StartVersion             int64                                    `json:"start_version,omitempty"`
	Status                   SharedDataObjectStatus                   `json:"status,omitempty"`
	StringSharedAs           string                                   `json:"string_shared_as,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sharedDataObjectFromPb(pb *sharedDataObjectPb) (*SharedDataObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObject{}
	st.AddedAt = pb.AddedAt
	st.AddedBy = pb.AddedBy
	st.CdfEnabled = pb.CdfEnabled
	st.Comment = pb.Comment
	st.Content = pb.Content
	st.DataObjectType = pb.DataObjectType
	st.HistoryDataSharingStatus = pb.HistoryDataSharingStatus
	st.Name = pb.Name
	st.Partitions = pb.Partitions
	st.SharedAs = pb.SharedAs
	st.StartVersion = pb.StartVersion
	st.Status = pb.Status
	st.StringSharedAs = pb.StringSharedAs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sharedDataObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sharedDataObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sharedDataObjectUpdateToPb(st *SharedDataObjectUpdate) (*sharedDataObjectUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharedDataObjectUpdatePb{}
	pb.Action = st.Action
	pb.DataObject = st.DataObject

	return pb, nil
}

type sharedDataObjectUpdatePb struct {
	Action     SharedDataObjectUpdateAction `json:"action,omitempty"`
	DataObject *SharedDataObject            `json:"data_object,omitempty"`
}

func sharedDataObjectUpdateFromPb(pb *sharedDataObjectUpdatePb) (*SharedDataObjectUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObjectUpdate{}
	st.Action = pb.Action
	st.DataObject = pb.DataObject

	return st, nil
}

func tableToPb(st *Table) (*tablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tablePb{}
	pb.Comment = st.Comment
	pb.Id = st.Id
	pb.InternalAttributes = st.InternalAttributes
	pb.MaterializationNamespace = st.MaterializationNamespace
	pb.MaterializedTableName = st.MaterializedTableName
	pb.Name = st.Name
	pb.Schema = st.Schema
	pb.Share = st.Share
	pb.ShareId = st.ShareId
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tablePb struct {
	Comment                  string                   `json:"comment,omitempty"`
	Id                       string                   `json:"id,omitempty"`
	InternalAttributes       *TableInternalAttributes `json:"internal_attributes,omitempty"`
	MaterializationNamespace string                   `json:"materialization_namespace,omitempty"`
	MaterializedTableName    string                   `json:"materialized_table_name,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	Schema                   string                   `json:"schema,omitempty"`
	Share                    string                   `json:"share,omitempty"`
	ShareId                  string                   `json:"share_id,omitempty"`
	Tags                     []catalog.TagKeyValue    `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableFromPb(pb *tablePb) (*Table, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Table{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	st.InternalAttributes = pb.InternalAttributes
	st.MaterializationNamespace = pb.MaterializationNamespace
	st.MaterializedTableName = pb.MaterializedTableName
	st.Name = pb.Name
	st.Schema = pb.Schema
	st.Share = pb.Share
	st.ShareId = pb.ShareId
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tableInternalAttributesToPb(st *TableInternalAttributes) (*tableInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableInternalAttributesPb{}
	pb.ParentStorageLocation = st.ParentStorageLocation
	pb.StorageLocation = st.StorageLocation
	pb.Type = st.Type
	pb.ViewDefinition = st.ViewDefinition

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableInternalAttributesPb struct {
	ParentStorageLocation string                                 `json:"parent_storage_location,omitempty"`
	StorageLocation       string                                 `json:"storage_location,omitempty"`
	Type                  TableInternalAttributesSharedTableType `json:"type,omitempty"`
	ViewDefinition        string                                 `json:"view_definition,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableInternalAttributesFromPb(pb *tableInternalAttributesPb) (*TableInternalAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableInternalAttributes{}
	st.ParentStorageLocation = pb.ParentStorageLocation
	st.StorageLocation = pb.StorageLocation
	st.Type = pb.Type
	st.ViewDefinition = pb.ViewDefinition

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableInternalAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableInternalAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateFederationPolicyRequestToPb(st *UpdateFederationPolicyRequest) (*updateFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateFederationPolicyRequestPb{}
	pb.Name = st.Name
	pb.Policy = st.Policy
	pb.RecipientName = st.RecipientName
	pb.UpdateMask = st.UpdateMask

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateFederationPolicyRequestPb struct {
	Name          string           `json:"-" url:"-"`
	Policy        FederationPolicy `json:"policy"`
	RecipientName string           `json:"-" url:"-"`
	UpdateMask    string           `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateFederationPolicyRequestFromPb(pb *updateFederationPolicyRequestPb) (*UpdateFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateFederationPolicyRequest{}
	st.Name = pb.Name
	st.Policy = pb.Policy
	st.RecipientName = pb.RecipientName
	st.UpdateMask = pb.UpdateMask

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateProviderToPb(st *UpdateProvider) (*updateProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.RecipientProfileStr = st.RecipientProfileStr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateProviderPb struct {
	Comment             string `json:"comment,omitempty"`
	Name                string `json:"-" url:"-"`
	NewName             string `json:"new_name,omitempty"`
	Owner               string `json:"owner,omitempty"`
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateProviderFromPb(pb *updateProviderPb) (*UpdateProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProvider{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.RecipientProfileStr = pb.RecipientProfileStr

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateProviderPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateProviderPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateRecipientToPb(st *UpdateRecipient) (*updateRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRecipientPb{}
	pb.Comment = st.Comment
	pb.ExpirationTime = st.ExpirationTime
	pb.IpAccessList = st.IpAccessList
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.PropertiesKvpairs = st.PropertiesKvpairs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateRecipientPb struct {
	Comment           string                      `json:"comment,omitempty"`
	ExpirationTime    int64                       `json:"expiration_time,omitempty"`
	IpAccessList      *IpAccessList               `json:"ip_access_list,omitempty"`
	Name              string                      `json:"-" url:"-"`
	NewName           string                      `json:"new_name,omitempty"`
	Owner             string                      `json:"owner,omitempty"`
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRecipientFromPb(pb *updateRecipientPb) (*UpdateRecipient, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRecipient{}
	st.Comment = pb.Comment
	st.ExpirationTime = pb.ExpirationTime
	st.IpAccessList = pb.IpAccessList
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.PropertiesKvpairs = pb.PropertiesKvpairs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRecipientPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRecipientPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateShareToPb(st *UpdateShare) (*updateSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSharePb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.StorageRoot = st.StorageRoot
	pb.Updates = st.Updates

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateSharePb struct {
	Comment     string                   `json:"comment,omitempty"`
	Name        string                   `json:"-" url:"-"`
	NewName     string                   `json:"new_name,omitempty"`
	Owner       string                   `json:"owner,omitempty"`
	StorageRoot string                   `json:"storage_root,omitempty"`
	Updates     []SharedDataObjectUpdate `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateShareFromPb(pb *updateSharePb) (*UpdateShare, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateShare{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.StorageRoot = pb.StorageRoot
	st.Updates = pb.Updates

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateSharePermissionsToPb(st *UpdateSharePermissions) (*updateSharePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSharePermissionsPb{}
	pb.Changes = st.Changes
	pb.Name = st.Name
	pb.OmitPermissionsList = st.OmitPermissionsList

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateSharePermissionsPb struct {
	Changes             []PermissionsChange `json:"changes,omitempty"`
	Name                string              `json:"-" url:"-"`
	OmitPermissionsList bool                `json:"omit_permissions_list,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateSharePermissionsFromPb(pb *updateSharePermissionsPb) (*UpdateSharePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSharePermissions{}
	st.Changes = pb.Changes
	st.Name = pb.Name
	st.OmitPermissionsList = pb.OmitPermissionsList

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateSharePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateSharePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateSharePermissionsResponseToPb(st *UpdateSharePermissionsResponse) (*updateSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSharePermissionsResponsePb{}
	pb.PrivilegeAssignments = st.PrivilegeAssignments

	return pb, nil
}

type updateSharePermissionsResponsePb struct {
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func updateSharePermissionsResponseFromPb(pb *updateSharePermissionsResponsePb) (*UpdateSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSharePermissionsResponse{}
	st.PrivilegeAssignments = pb.PrivilegeAssignments

	return st, nil
}

func volumeToPb(st *Volume) (*volumePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumePb{}
	pb.Comment = st.Comment
	pb.Id = st.Id
	pb.InternalAttributes = st.InternalAttributes
	pb.Name = st.Name
	pb.Schema = st.Schema
	pb.Share = st.Share
	pb.ShareId = st.ShareId
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type volumePb struct {
	Comment            string                    `json:"comment,omitempty"`
	Id                 string                    `json:"id,omitempty"`
	InternalAttributes *VolumeInternalAttributes `json:"internal_attributes,omitempty"`
	Name               string                    `json:"name,omitempty"`
	Schema             string                    `json:"schema,omitempty"`
	Share              string                    `json:"share,omitempty"`
	ShareId            string                    `json:"share_id,omitempty"`
	Tags               []catalog.TagKeyValue     `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func volumeFromPb(pb *volumePb) (*Volume, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Volume{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	st.InternalAttributes = pb.InternalAttributes
	st.Name = pb.Name
	st.Schema = pb.Schema
	st.Share = pb.Share
	st.ShareId = pb.ShareId
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *volumePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st volumePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func volumeInternalAttributesToPb(st *VolumeInternalAttributes) (*volumeInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumeInternalAttributesPb{}
	pb.StorageLocation = st.StorageLocation
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type volumeInternalAttributesPb struct {
	StorageLocation string `json:"storage_location,omitempty"`
	Type            string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func volumeInternalAttributesFromPb(pb *volumeInternalAttributesPb) (*VolumeInternalAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumeInternalAttributes{}
	st.StorageLocation = pb.StorageLocation
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *volumeInternalAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st volumeInternalAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	s := t.Format(time.RFC3339Nano)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339Nano, *s)
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
	if *s == "" {
		return &[]string{}, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
