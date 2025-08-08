// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/databricks-sdk-go/service/catalog/catalogpb"
	"github.com/databricks/databricks-sdk-go/service/sharing/sharingpb"
)

// The delta sharing authentication type.
type AuthenticationType string

const AuthenticationTypeDatabricks AuthenticationType = `DATABRICKS`

const AuthenticationTypeOauthClientCredentials AuthenticationType = `OAUTH_CLIENT_CREDENTIALS`

const AuthenticationTypeOidcFederation AuthenticationType = `OIDC_FEDERATION`

const AuthenticationTypeToken AuthenticationType = `TOKEN`

// String representation for [fmt.Print]
func (f *AuthenticationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationType) Set(v string) error {
	switch v {
	case `DATABRICKS`, `OAUTH_CLIENT_CREDENTIALS`, `OIDC_FEDERATION`, `TOKEN`:
		*f = AuthenticationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS", "OAUTH_CLIENT_CREDENTIALS", "OIDC_FEDERATION", "TOKEN"`, v)
	}
}

// Values returns all possible values for AuthenticationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		AuthenticationTypeDatabricks,
		AuthenticationTypeOauthClientCredentials,
		AuthenticationTypeOidcFederation,
		AuthenticationTypeToken,
	}
}

// Type always returns AuthenticationType to satisfy [pflag.Value] interface
func (f *AuthenticationType) Type() string {
	return "AuthenticationType"
}

func AuthenticationTypeToPb(st *AuthenticationType) (*sharingpb.AuthenticationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.AuthenticationTypePb(*st)
	return &pb, nil
}

func AuthenticationTypeFromPb(pb *sharingpb.AuthenticationTypePb) (*AuthenticationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AuthenticationType(*pb)
	return &st, nil
}

// UC supported column types Copied from
// https://src.dev.databricks.com/databricks/universe@23a85902bb58695ab9293adc9f327b0714b55e72/-/blob/managed-catalog/api/messages/table.proto?L68
type ColumnTypeName string

const ColumnTypeNameArray ColumnTypeName = `ARRAY`

const ColumnTypeNameBinary ColumnTypeName = `BINARY`

const ColumnTypeNameBoolean ColumnTypeName = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeName = `BYTE`

const ColumnTypeNameChar ColumnTypeName = `CHAR`

const ColumnTypeNameDate ColumnTypeName = `DATE`

const ColumnTypeNameDecimal ColumnTypeName = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeName = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeName = `FLOAT`

const ColumnTypeNameInt ColumnTypeName = `INT`

const ColumnTypeNameInterval ColumnTypeName = `INTERVAL`

const ColumnTypeNameLong ColumnTypeName = `LONG`

const ColumnTypeNameMap ColumnTypeName = `MAP`

const ColumnTypeNameNull ColumnTypeName = `NULL`

const ColumnTypeNameShort ColumnTypeName = `SHORT`

const ColumnTypeNameString ColumnTypeName = `STRING`

const ColumnTypeNameStruct ColumnTypeName = `STRUCT`

const ColumnTypeNameTableType ColumnTypeName = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeName = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeName = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeName = `USER_DEFINED_TYPE`

const ColumnTypeNameVariant ColumnTypeName = `VARIANT`

// String representation for [fmt.Print]
func (f *ColumnTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`:
		*f = ColumnTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TABLE_TYPE", "TIMESTAMP", "TIMESTAMP_NTZ", "USER_DEFINED_TYPE", "VARIANT"`, v)
	}
}

// Values returns all possible values for ColumnTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnTypeName) Values() []ColumnTypeName {
	return []ColumnTypeName{
		ColumnTypeNameArray,
		ColumnTypeNameBinary,
		ColumnTypeNameBoolean,
		ColumnTypeNameByte,
		ColumnTypeNameChar,
		ColumnTypeNameDate,
		ColumnTypeNameDecimal,
		ColumnTypeNameDouble,
		ColumnTypeNameFloat,
		ColumnTypeNameInt,
		ColumnTypeNameInterval,
		ColumnTypeNameLong,
		ColumnTypeNameMap,
		ColumnTypeNameNull,
		ColumnTypeNameShort,
		ColumnTypeNameString,
		ColumnTypeNameStruct,
		ColumnTypeNameTableType,
		ColumnTypeNameTimestamp,
		ColumnTypeNameTimestampNtz,
		ColumnTypeNameUserDefinedType,
		ColumnTypeNameVariant,
	}
}

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

func ColumnTypeNameToPb(st *ColumnTypeName) (*sharingpb.ColumnTypeNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.ColumnTypeNamePb(*st)
	return &pb, nil
}

func ColumnTypeNameFromPb(pb *sharingpb.ColumnTypeNamePb) (*ColumnTypeName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ColumnTypeName(*pb)
	return &st, nil
}

type CreateFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be created.
	// Wire name: 'policy'
	Policy FederationPolicy ``
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being created.
	// Wire name: 'recipient_name'
	RecipientName string `tf:"-"`
}

func CreateFederationPolicyRequestToPb(st *CreateFederationPolicyRequest) (*sharingpb.CreateFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.CreateFederationPolicyRequestPb{}
	policyPb, err := FederationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.RecipientName = st.RecipientName

	return pb, nil
}

func CreateFederationPolicyRequestFromPb(pb *sharingpb.CreateFederationPolicyRequestPb) (*CreateFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFederationPolicyRequest{}
	policyField, err := FederationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.RecipientName = pb.RecipientName

	return st, nil
}

type CreateProvider struct {

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType ``
	// Description about the provider.
	// Wire name: 'comment'
	Comment string ``
	// The name of the Provider.
	// Wire name: 'name'
	Name string ``
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string   ``
	ForceSendFields     []string `tf:"-"`
}

func (s *CreateProvider) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateProvider) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateProviderToPb(st *CreateProvider) (*sharingpb.CreateProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.CreateProviderPb{}
	authenticationTypePb, err := AuthenticationTypeToPb(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.RecipientProfileStr = st.RecipientProfileStr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateProviderFromPb(pb *sharingpb.CreateProviderPb) (*CreateProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateProvider{}
	authenticationTypeField, err := AuthenticationTypeFromPb(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.RecipientProfileStr = pb.RecipientProfileStr

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateRecipient struct {

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType ``
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string ``
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string ``
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList ``
	// Name of Recipient.
	// Wire name: 'name'
	Name string ``
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string ``
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs ``
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'sharing_code'
	SharingCode     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateRecipient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRecipient) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateRecipientToPb(st *CreateRecipient) (*sharingpb.CreateRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.CreateRecipientPb{}
	authenticationTypePb, err := AuthenticationTypeToPb(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}
	pb.Comment = st.Comment
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.ExpirationTime = st.ExpirationTime
	ipAccessListPb, err := IpAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}
	pb.Name = st.Name
	pb.Owner = st.Owner
	propertiesKvpairsPb, err := SecurablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}
	pb.SharingCode = st.SharingCode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateRecipientFromPb(pb *sharingpb.CreateRecipientPb) (*CreateRecipient, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRecipient{}
	authenticationTypeField, err := AuthenticationTypeFromPb(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	st.Comment = pb.Comment
	st.DataRecipientGlobalMetastoreId = pb.DataRecipientGlobalMetastoreId
	st.ExpirationTime = pb.ExpirationTime
	ipAccessListField, err := IpAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	st.Name = pb.Name
	st.Owner = pb.Owner
	propertiesKvpairsField, err := SecurablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsField != nil {
		st.PropertiesKvpairs = propertiesKvpairsField
	}
	st.SharingCode = pb.SharingCode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Name of the share.
	// Wire name: 'name'
	Name string ``
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateShareToPb(st *CreateShare) (*sharingpb.CreateSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.CreateSharePb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateShareFromPb(pb *sharingpb.CreateSharePb) (*CreateShare, error) {
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

type DeleteFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be deleted.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being deleted.
	// Wire name: 'recipient_name'
	RecipientName string `tf:"-"`
}

func DeleteFederationPolicyRequestToPb(st *DeleteFederationPolicyRequest) (*sharingpb.DeleteFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeleteFederationPolicyRequestPb{}
	pb.Name = st.Name
	pb.RecipientName = st.RecipientName

	return pb, nil
}

func DeleteFederationPolicyRequestFromPb(pb *sharingpb.DeleteFederationPolicyRequestPb) (*DeleteFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFederationPolicyRequest{}
	st.Name = pb.Name
	st.RecipientName = pb.RecipientName

	return st, nil
}

type DeleteProviderRequest struct {
	// Name of the provider.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteProviderRequestToPb(st *DeleteProviderRequest) (*sharingpb.DeleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeleteProviderRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteProviderRequestFromPb(pb *sharingpb.DeleteProviderRequestPb) (*DeleteProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteProviderRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteRecipientRequest struct {
	// Name of the recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteRecipientRequestToPb(st *DeleteRecipientRequest) (*sharingpb.DeleteRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeleteRecipientRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteRecipientRequestFromPb(pb *sharingpb.DeleteRecipientRequestPb) (*DeleteRecipientRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRecipientRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteShareRequest struct {
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteShareRequestToPb(st *DeleteShareRequest) (*sharingpb.DeleteShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeleteShareRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteShareRequestFromPb(pb *sharingpb.DeleteShareRequestPb) (*DeleteShareRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteShareRequest{}
	st.Name = pb.Name

	return st, nil
}

// Represents a UC dependency.
type DeltaSharingDependency struct {

	// Wire name: 'function'
	Function *DeltaSharingFunctionDependency ``

	// Wire name: 'table'
	Table *DeltaSharingTableDependency ``
}

func DeltaSharingDependencyToPb(st *DeltaSharingDependency) (*sharingpb.DeltaSharingDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingDependencyPb{}
	functionPb, err := DeltaSharingFunctionDependencyToPb(st.Function)
	if err != nil {
		return nil, err
	}
	if functionPb != nil {
		pb.Function = functionPb
	}
	tablePb, err := DeltaSharingTableDependencyToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
}

func DeltaSharingDependencyFromPb(pb *sharingpb.DeltaSharingDependencyPb) (*DeltaSharingDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingDependency{}
	functionField, err := DeltaSharingFunctionDependencyFromPb(pb.Function)
	if err != nil {
		return nil, err
	}
	if functionField != nil {
		st.Function = functionField
	}
	tableField, err := DeltaSharingTableDependencyFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}

	return st, nil
}

// Represents a list of dependencies.
type DeltaSharingDependencyList struct {
	// An array of Dependency.
	// Wire name: 'dependencies'
	Dependencies []DeltaSharingDependency ``
}

func DeltaSharingDependencyListToPb(st *DeltaSharingDependencyList) (*sharingpb.DeltaSharingDependencyListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingDependencyListPb{}

	var dependenciesPb []sharingpb.DeltaSharingDependencyPb
	for _, item := range st.Dependencies {
		itemPb, err := DeltaSharingDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependenciesPb = append(dependenciesPb, *itemPb)
		}
	}
	pb.Dependencies = dependenciesPb

	return pb, nil
}

func DeltaSharingDependencyListFromPb(pb *sharingpb.DeltaSharingDependencyListPb) (*DeltaSharingDependencyList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingDependencyList{}

	var dependenciesField []DeltaSharingDependency
	for _, itemPb := range pb.Dependencies {
		item, err := DeltaSharingDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependenciesField = append(dependenciesField, *item)
		}
	}
	st.Dependencies = dependenciesField

	return st, nil
}

type DeltaSharingFunction struct {
	// The aliass of registered model.
	// Wire name: 'aliases'
	Aliases []RegisteredModelAlias ``
	// The comment of the function.
	// Wire name: 'comment'
	Comment string ``
	// The data type of the function.
	// Wire name: 'data_type'
	DataType ColumnTypeName ``
	// The dependency list of the function.
	// Wire name: 'dependency_list'
	DependencyList *DeltaSharingDependencyList ``
	// The full data type of the function.
	// Wire name: 'full_data_type'
	FullDataType string ``
	// The id of the function.
	// Wire name: 'id'
	Id string ``
	// The function parameter information.
	// Wire name: 'input_params'
	InputParams *FunctionParameterInfos ``
	// The name of the function.
	// Wire name: 'name'
	Name string ``
	// The properties of the function.
	// Wire name: 'properties'
	Properties string ``
	// The routine definition of the function.
	// Wire name: 'routine_definition'
	RoutineDefinition string ``
	// The name of the schema that the function belongs to.
	// Wire name: 'schema'
	Schema string ``
	// The securable kind of the function.
	// Wire name: 'securable_kind'
	SecurableKind SharedSecurableKind ``
	// The name of the share that the function belongs to.
	// Wire name: 'share'
	Share string ``
	// The id of the share that the function belongs to.
	// Wire name: 'share_id'
	ShareId string ``
	// The storage location of the function.
	// Wire name: 'storage_location'
	StorageLocation string ``
	// The tags of the function.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue ``
	ForceSendFields []string              `tf:"-"`
}

func (s *DeltaSharingFunction) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSharingFunction) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeltaSharingFunctionToPb(st *DeltaSharingFunction) (*sharingpb.DeltaSharingFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingFunctionPb{}

	var aliasesPb []sharingpb.RegisteredModelAliasPb
	for _, item := range st.Aliases {
		itemPb, err := RegisteredModelAliasToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			aliasesPb = append(aliasesPb, *itemPb)
		}
	}
	pb.Aliases = aliasesPb
	pb.Comment = st.Comment
	dataTypePb, err := ColumnTypeNameToPb(&st.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypePb != nil {
		pb.DataType = *dataTypePb
	}
	dependencyListPb, err := DeltaSharingDependencyListToPb(st.DependencyList)
	if err != nil {
		return nil, err
	}
	if dependencyListPb != nil {
		pb.DependencyList = dependencyListPb
	}
	pb.FullDataType = st.FullDataType
	pb.Id = st.Id
	inputParamsPb, err := FunctionParameterInfosToPb(st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = inputParamsPb
	}
	pb.Name = st.Name
	pb.Properties = st.Properties
	pb.RoutineDefinition = st.RoutineDefinition
	pb.Schema = st.Schema
	securableKindPb, err := SharedSecurableKindToPb(&st.SecurableKind)
	if err != nil {
		return nil, err
	}
	if securableKindPb != nil {
		pb.SecurableKind = *securableKindPb
	}
	pb.Share = st.Share
	pb.ShareId = st.ShareId
	pb.StorageLocation = st.StorageLocation

	var tagsPb []catalogpb.TagKeyValuePb
	for _, item := range st.Tags {
		itemPb, err := catalog.TagKeyValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeltaSharingFunctionFromPb(pb *sharingpb.DeltaSharingFunctionPb) (*DeltaSharingFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingFunction{}

	var aliasesField []RegisteredModelAlias
	for _, itemPb := range pb.Aliases {
		item, err := RegisteredModelAliasFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			aliasesField = append(aliasesField, *item)
		}
	}
	st.Aliases = aliasesField
	st.Comment = pb.Comment
	dataTypeField, err := ColumnTypeNameFromPb(&pb.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypeField != nil {
		st.DataType = *dataTypeField
	}
	dependencyListField, err := DeltaSharingDependencyListFromPb(pb.DependencyList)
	if err != nil {
		return nil, err
	}
	if dependencyListField != nil {
		st.DependencyList = dependencyListField
	}
	st.FullDataType = pb.FullDataType
	st.Id = pb.Id
	inputParamsField, err := FunctionParameterInfosFromPb(pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = inputParamsField
	}
	st.Name = pb.Name
	st.Properties = pb.Properties
	st.RoutineDefinition = pb.RoutineDefinition
	st.Schema = pb.Schema
	securableKindField, err := SharedSecurableKindFromPb(&pb.SecurableKind)
	if err != nil {
		return nil, err
	}
	if securableKindField != nil {
		st.SecurableKind = *securableKindField
	}
	st.Share = pb.Share
	st.ShareId = pb.ShareId
	st.StorageLocation = pb.StorageLocation

	var tagsField []catalog.TagKeyValue
	for _, itemPb := range pb.Tags {
		item, err := catalog.TagKeyValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency struct {

	// Wire name: 'function_name'
	FunctionName string ``

	// Wire name: 'schema_name'
	SchemaName      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DeltaSharingFunctionDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSharingFunctionDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeltaSharingFunctionDependencyToPb(st *DeltaSharingFunctionDependency) (*sharingpb.DeltaSharingFunctionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingFunctionDependencyPb{}
	pb.FunctionName = st.FunctionName
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeltaSharingFunctionDependencyFromPb(pb *sharingpb.DeltaSharingFunctionDependencyPb) (*DeltaSharingFunctionDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingFunctionDependency{}
	st.FunctionName = pb.FunctionName
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A Table in UC as a dependency.
type DeltaSharingTableDependency struct {

	// Wire name: 'schema_name'
	SchemaName string ``

	// Wire name: 'table_name'
	TableName       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DeltaSharingTableDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeltaSharingTableDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeltaSharingTableDependencyToPb(st *DeltaSharingTableDependency) (*sharingpb.DeltaSharingTableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingTableDependencyPb{}
	pb.SchemaName = st.SchemaName
	pb.TableName = st.TableName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeltaSharingTableDependencyFromPb(pb *sharingpb.DeltaSharingTableDependencyPb) (*DeltaSharingTableDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingTableDependency{}
	st.SchemaName = pb.SchemaName
	st.TableName = pb.TableName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type FederationPolicy struct {
	// Description of the policy. This is a user-provided description.
	// Wire name: 'comment'
	Comment string ``
	// System-generated timestamp indicating when the policy was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// Unique, immutable system-generated identifier for the federation policy.
	// Wire name: 'id'
	Id string ``
	// Name of the federation policy. A recipient can have multiple policies
	// with different names. The name must contain only lowercase alphanumeric
	// characters, numbers, and hyphens.
	// Wire name: 'name'
	Name string ``
	// Specifies the policy to use for validating OIDC claims in the federated
	// tokens.
	// Wire name: 'oidc_policy'
	OidcPolicy *OidcFederationPolicy ``
	// System-generated timestamp indicating when the policy was last updated.
	// Wire name: 'update_time'
	UpdateTime      *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *FederationPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FederationPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FederationPolicyToPb(st *FederationPolicy) (*sharingpb.FederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.FederationPolicyPb{}
	pb.Comment = st.Comment
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.Id = st.Id
	pb.Name = st.Name
	oidcPolicyPb, err := OidcFederationPolicyToPb(st.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyPb != nil {
		pb.OidcPolicy = oidcPolicyPb
	}
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FederationPolicyFromPb(pb *sharingpb.FederationPolicyPb) (*FederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FederationPolicy{}
	st.Comment = pb.Comment
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.Id = pb.Id
	st.Name = pb.Name
	oidcPolicyField, err := OidcFederationPolicyFromPb(pb.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyField != nil {
		st.OidcPolicy = oidcPolicyField
	}
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Represents a parameter of a function. The same message is used for both input
// and output columns.
type FunctionParameterInfo struct {
	// The comment of the parameter.
	// Wire name: 'comment'
	Comment string ``
	// The name of the parameter.
	// Wire name: 'name'
	Name string ``
	// The default value of the parameter.
	// Wire name: 'parameter_default'
	ParameterDefault string ``
	// The mode of the function parameter.
	// Wire name: 'parameter_mode'
	ParameterMode FunctionParameterMode ``
	// The type of the function parameter.
	// Wire name: 'parameter_type'
	ParameterType FunctionParameterType ``
	// The position of the parameter.
	// Wire name: 'position'
	Position int ``
	// The interval type of the parameter type.
	// Wire name: 'type_interval_type'
	TypeIntervalType string ``
	// The type of the parameter in JSON format.
	// Wire name: 'type_json'
	TypeJson string ``
	// The type of the parameter in Enum format.
	// Wire name: 'type_name'
	TypeName ColumnTypeName ``
	// The precision of the parameter type.
	// Wire name: 'type_precision'
	TypePrecision int ``
	// The scale of the parameter type.
	// Wire name: 'type_scale'
	TypeScale int ``
	// The type of the parameter in text format.
	// Wire name: 'type_text'
	TypeText        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FunctionParameterInfoToPb(st *FunctionParameterInfo) (*sharingpb.FunctionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.FunctionParameterInfoPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.ParameterDefault = st.ParameterDefault
	parameterModePb, err := FunctionParameterModeToPb(&st.ParameterMode)
	if err != nil {
		return nil, err
	}
	if parameterModePb != nil {
		pb.ParameterMode = *parameterModePb
	}
	parameterTypePb, err := FunctionParameterTypeToPb(&st.ParameterType)
	if err != nil {
		return nil, err
	}
	if parameterTypePb != nil {
		pb.ParameterType = *parameterTypePb
	}
	pb.Position = st.Position
	pb.TypeIntervalType = st.TypeIntervalType
	pb.TypeJson = st.TypeJson
	typeNamePb, err := ColumnTypeNameToPb(&st.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FunctionParameterInfoFromPb(pb *sharingpb.FunctionParameterInfoPb) (*FunctionParameterInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfo{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.ParameterDefault = pb.ParameterDefault
	parameterModeField, err := FunctionParameterModeFromPb(&pb.ParameterMode)
	if err != nil {
		return nil, err
	}
	if parameterModeField != nil {
		st.ParameterMode = *parameterModeField
	}
	parameterTypeField, err := FunctionParameterTypeFromPb(&pb.ParameterType)
	if err != nil {
		return nil, err
	}
	if parameterTypeField != nil {
		st.ParameterType = *parameterTypeField
	}
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	typeNameField, err := ColumnTypeNameFromPb(&pb.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type FunctionParameterInfos struct {
	// The list of parameters of the function.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo ``
}

func FunctionParameterInfosToPb(st *FunctionParameterInfos) (*sharingpb.FunctionParameterInfosPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.FunctionParameterInfosPb{}

	var parametersPb []sharingpb.FunctionParameterInfoPb
	for _, item := range st.Parameters {
		itemPb, err := FunctionParameterInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	return pb, nil
}

func FunctionParameterInfosFromPb(pb *sharingpb.FunctionParameterInfosPb) (*FunctionParameterInfos, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfos{}

	var parametersField []FunctionParameterInfo
	for _, itemPb := range pb.Parameters {
		item, err := FunctionParameterInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField

	return st, nil
}

type FunctionParameterMode string

const FunctionParameterModeIn FunctionParameterMode = `IN`

const FunctionParameterModeInout FunctionParameterMode = `INOUT`

const FunctionParameterModeOut FunctionParameterMode = `OUT`

// String representation for [fmt.Print]
func (f *FunctionParameterMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterMode) Set(v string) error {
	switch v {
	case `IN`, `INOUT`, `OUT`:
		*f = FunctionParameterMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN", "INOUT", "OUT"`, v)
	}
}

// Values returns all possible values for FunctionParameterMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterMode) Values() []FunctionParameterMode {
	return []FunctionParameterMode{
		FunctionParameterModeIn,
		FunctionParameterModeInout,
		FunctionParameterModeOut,
	}
}

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

func FunctionParameterModeToPb(st *FunctionParameterMode) (*sharingpb.FunctionParameterModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.FunctionParameterModePb(*st)
	return &pb, nil
}

func FunctionParameterModeFromPb(pb *sharingpb.FunctionParameterModePb) (*FunctionParameterMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterMode(*pb)
	return &st, nil
}

type FunctionParameterType string

const FunctionParameterTypeColumn FunctionParameterType = `COLUMN`

const FunctionParameterTypeParam FunctionParameterType = `PARAM`

// String representation for [fmt.Print]
func (f *FunctionParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *FunctionParameterType) Set(v string) error {
	switch v {
	case `COLUMN`, `PARAM`:
		*f = FunctionParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COLUMN", "PARAM"`, v)
	}
}

// Values returns all possible values for FunctionParameterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *FunctionParameterType) Values() []FunctionParameterType {
	return []FunctionParameterType{
		FunctionParameterTypeColumn,
		FunctionParameterTypeParam,
	}
}

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

func FunctionParameterTypeToPb(st *FunctionParameterType) (*sharingpb.FunctionParameterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.FunctionParameterTypePb(*st)
	return &pb, nil
}

func FunctionParameterTypeFromPb(pb *sharingpb.FunctionParameterTypePb) (*FunctionParameterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterType(*pb)
	return &st, nil
}

type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	// Wire name: 'activation_url'
	ActivationUrl string `tf:"-"`
}

func GetActivationUrlInfoRequestToPb(st *GetActivationUrlInfoRequest) (*sharingpb.GetActivationUrlInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetActivationUrlInfoRequestPb{}
	pb.ActivationUrl = st.ActivationUrl

	return pb, nil
}

func GetActivationUrlInfoRequestFromPb(pb *sharingpb.GetActivationUrlInfoRequestPb) (*GetActivationUrlInfoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetActivationUrlInfoRequest{}
	st.ActivationUrl = pb.ActivationUrl

	return st, nil
}

type GetFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be retrieved.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being retrieved.
	// Wire name: 'recipient_name'
	RecipientName string `tf:"-"`
}

func GetFederationPolicyRequestToPb(st *GetFederationPolicyRequest) (*sharingpb.GetFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetFederationPolicyRequestPb{}
	pb.Name = st.Name
	pb.RecipientName = st.RecipientName

	return pb, nil
}

func GetFederationPolicyRequestFromPb(pb *sharingpb.GetFederationPolicyRequestPb) (*GetFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFederationPolicyRequest{}
	st.Name = pb.Name
	st.RecipientName = pb.RecipientName

	return st, nil
}

type GetProviderRequest struct {
	// Name of the provider.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetProviderRequestToPb(st *GetProviderRequest) (*sharingpb.GetProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetProviderRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetProviderRequestFromPb(pb *sharingpb.GetProviderRequestPb) (*GetProviderRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetProviderRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetRecipientRequest struct {
	// Name of the recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetRecipientRequestToPb(st *GetRecipientRequest) (*sharingpb.GetRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetRecipientRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetRecipientRequestFromPb(pb *sharingpb.GetRecipientRequestPb) (*GetRecipientRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRecipientRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of data share permissions for a recipient.
	// Wire name: 'permissions_out'
	PermissionsOut  []ShareToPrivilegeAssignment ``
	ForceSendFields []string                     `tf:"-"`
}

func (s *GetRecipientSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRecipientSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetRecipientSharePermissionsResponseToPb(st *GetRecipientSharePermissionsResponse) (*sharingpb.GetRecipientSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetRecipientSharePermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var permissionsOutPb []sharingpb.ShareToPrivilegeAssignmentPb
	for _, item := range st.PermissionsOut {
		itemPb, err := ShareToPrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionsOutPb = append(permissionsOutPb, *itemPb)
		}
	}
	pb.PermissionsOut = permissionsOutPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetRecipientSharePermissionsResponseFromPb(pb *sharingpb.GetRecipientSharePermissionsResponsePb) (*GetRecipientSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRecipientSharePermissionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var permissionsOutField []ShareToPrivilegeAssignment
	for _, itemPb := range pb.PermissionsOut {
		item, err := ShareToPrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionsOutField = append(permissionsOutField, *item)
		}
	}
	st.PermissionsOut = permissionsOutField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment ``
	ForceSendFields      []string              `tf:"-"`
}

func (s *GetSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetSharePermissionsResponseToPb(st *GetSharePermissionsResponse) (*sharingpb.GetSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetSharePermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var privilegeAssignmentsPb []sharingpb.PrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := PrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetSharePermissionsResponseFromPb(pb *sharingpb.GetSharePermissionsResponsePb) (*GetSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSharePermissionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := PrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetShareRequest struct {
	// Query for data to include in the share.
	// Wire name: 'include_shared_data'
	IncludeSharedData bool `tf:"-"`
	// The name of the share.
	// Wire name: 'name'
	Name            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *GetShareRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetShareRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetShareRequestToPb(st *GetShareRequest) (*sharingpb.GetShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetShareRequestPb{}
	pb.IncludeSharedData = st.IncludeSharedData
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetShareRequestFromPb(pb *sharingpb.GetShareRequestPb) (*GetShareRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetShareRequest{}
	st.IncludeSharedData = pb.IncludeSharedData
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	// Wire name: 'allowed_ip_addresses'
	AllowedIpAddresses []string ``
}

func IpAccessListToPb(st *IpAccessList) (*sharingpb.IpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.IpAccessListPb{}
	pb.AllowedIpAddresses = st.AllowedIpAddresses

	return pb, nil
}

func IpAccessListFromPb(pb *sharingpb.IpAccessListPb) (*IpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IpAccessList{}
	st.AllowedIpAddresses = pb.AllowedIpAddresses

	return st, nil
}

type ListFederationPoliciesRequest struct {

	// Wire name: 'max_results'
	MaxResults int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policies are being listed.
	// Wire name: 'recipient_name'
	RecipientName   string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFederationPoliciesRequestToPb(st *ListFederationPoliciesRequest) (*sharingpb.ListFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListFederationPoliciesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.RecipientName = st.RecipientName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListFederationPoliciesRequestFromPb(pb *sharingpb.ListFederationPoliciesRequestPb) (*ListFederationPoliciesRequest, error) {
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

type ListFederationPoliciesResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'policies'
	Policies        []FederationPolicy ``
	ForceSendFields []string           `tf:"-"`
}

func (s *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListFederationPoliciesResponseToPb(st *ListFederationPoliciesResponse) (*sharingpb.ListFederationPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListFederationPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var policiesPb []sharingpb.FederationPolicyPb
	for _, item := range st.Policies {
		itemPb, err := FederationPolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policiesPb = append(policiesPb, *itemPb)
		}
	}
	pb.Policies = policiesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListFederationPoliciesResponseFromPb(pb *sharingpb.ListFederationPoliciesResponsePb) (*ListFederationPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFederationPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken

	var policiesField []FederationPolicy
	for _, itemPb := range pb.Policies {
		item, err := FederationPolicyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policiesField = append(policiesField, *item)
		}
	}
	st.Policies = policiesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListProviderShareAssetsRequest struct {
	// Maximum number of functions to return.
	// Wire name: 'function_max_results'
	FunctionMaxResults int `tf:"-"`
	// Maximum number of notebooks to return.
	// Wire name: 'notebook_max_results'
	NotebookMaxResults int `tf:"-"`
	// The name of the provider who owns the share.
	// Wire name: 'provider_name'
	ProviderName string `tf:"-"`
	// The name of the share.
	// Wire name: 'share_name'
	ShareName string `tf:"-"`
	// Maximum number of tables to return.
	// Wire name: 'table_max_results'
	TableMaxResults int `tf:"-"`
	// Maximum number of volumes to return.
	// Wire name: 'volume_max_results'
	VolumeMaxResults int      `tf:"-"`
	ForceSendFields  []string `tf:"-"`
}

func (s *ListProviderShareAssetsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderShareAssetsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProviderShareAssetsRequestToPb(st *ListProviderShareAssetsRequest) (*sharingpb.ListProviderShareAssetsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListProviderShareAssetsRequestPb{}
	pb.FunctionMaxResults = st.FunctionMaxResults
	pb.NotebookMaxResults = st.NotebookMaxResults
	pb.ProviderName = st.ProviderName
	pb.ShareName = st.ShareName
	pb.TableMaxResults = st.TableMaxResults
	pb.VolumeMaxResults = st.VolumeMaxResults

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListProviderShareAssetsRequestFromPb(pb *sharingpb.ListProviderShareAssetsRequestPb) (*ListProviderShareAssetsRequest, error) {
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

// Response to ListProviderShareAssets, which contains the list of assets of a
// share.
type ListProviderShareAssetsResponse struct {
	// The list of functions in the share.
	// Wire name: 'functions'
	Functions []DeltaSharingFunction ``
	// The list of notebooks in the share.
	// Wire name: 'notebooks'
	Notebooks []NotebookFile ``
	// The list of tables in the share.
	// Wire name: 'tables'
	Tables []Table ``
	// The list of volumes in the share.
	// Wire name: 'volumes'
	Volumes []Volume ``
}

func ListProviderShareAssetsResponseToPb(st *ListProviderShareAssetsResponse) (*sharingpb.ListProviderShareAssetsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListProviderShareAssetsResponsePb{}

	var functionsPb []sharingpb.DeltaSharingFunctionPb
	for _, item := range st.Functions {
		itemPb, err := DeltaSharingFunctionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			functionsPb = append(functionsPb, *itemPb)
		}
	}
	pb.Functions = functionsPb

	var notebooksPb []sharingpb.NotebookFilePb
	for _, item := range st.Notebooks {
		itemPb, err := NotebookFileToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebooksPb = append(notebooksPb, *itemPb)
		}
	}
	pb.Notebooks = notebooksPb

	var tablesPb []sharingpb.TablePb
	for _, item := range st.Tables {
		itemPb, err := TableToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tablesPb = append(tablesPb, *itemPb)
		}
	}
	pb.Tables = tablesPb

	var volumesPb []sharingpb.VolumePb
	for _, item := range st.Volumes {
		itemPb, err := VolumeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			volumesPb = append(volumesPb, *itemPb)
		}
	}
	pb.Volumes = volumesPb

	return pb, nil
}

func ListProviderShareAssetsResponseFromPb(pb *sharingpb.ListProviderShareAssetsResponsePb) (*ListProviderShareAssetsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderShareAssetsResponse{}

	var functionsField []DeltaSharingFunction
	for _, itemPb := range pb.Functions {
		item, err := DeltaSharingFunctionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			functionsField = append(functionsField, *item)
		}
	}
	st.Functions = functionsField

	var notebooksField []NotebookFile
	for _, itemPb := range pb.Notebooks {
		item, err := NotebookFileFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebooksField = append(notebooksField, *item)
		}
	}
	st.Notebooks = notebooksField

	var tablesField []Table
	for _, itemPb := range pb.Tables {
		item, err := TableFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tablesField = append(tablesField, *item)
		}
	}
	st.Tables = tablesField

	var volumesField []Volume
	for _, itemPb := range pb.Volumes {
		item, err := VolumeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			volumesField = append(volumesField, *item)
		}
	}
	st.Volumes = volumesField

	return st, nil
}

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of provider shares.
	// Wire name: 'shares'
	Shares          []ProviderShare ``
	ForceSendFields []string        `tf:"-"`
}

func (s *ListProviderSharesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProviderSharesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProviderSharesResponseToPb(st *ListProviderSharesResponse) (*sharingpb.ListProviderSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListProviderSharesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var sharesPb []sharingpb.ProviderSharePb
	for _, item := range st.Shares {
		itemPb, err := ProviderShareToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sharesPb = append(sharesPb, *itemPb)
		}
	}
	pb.Shares = sharesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListProviderSharesResponseFromPb(pb *sharingpb.ListProviderSharesResponsePb) (*ListProviderSharesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderSharesResponse{}
	st.NextPageToken = pb.NextPageToken

	var sharesField []ProviderShare
	for _, itemPb := range pb.Shares {
		item, err := ProviderShareFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sharesField = append(sharesField, *item)
		}
	}
	st.Shares = sharesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	// Wire name: 'data_provider_global_metastore_id'
	DataProviderGlobalMetastoreId string `tf:"-"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProvidersRequestToPb(st *ListProvidersRequest) (*sharingpb.ListProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListProvidersRequestPb{}
	pb.DataProviderGlobalMetastoreId = st.DataProviderGlobalMetastoreId
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListProvidersRequestFromPb(pb *sharingpb.ListProvidersRequestPb) (*ListProvidersRequest, error) {
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

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of provider information objects.
	// Wire name: 'providers'
	Providers       []ProviderInfo ``
	ForceSendFields []string       `tf:"-"`
}

func (s *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListProvidersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListProvidersResponseToPb(st *ListProvidersResponse) (*sharingpb.ListProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListProvidersResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var providersPb []sharingpb.ProviderInfoPb
	for _, item := range st.Providers {
		itemPb, err := ProviderInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			providersPb = append(providersPb, *itemPb)
		}
	}
	pb.Providers = providersPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListProvidersResponseFromPb(pb *sharingpb.ListProvidersResponsePb) (*ListProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersResponse{}
	st.NextPageToken = pb.NextPageToken

	var providersField []ProviderInfo
	for _, itemPb := range pb.Providers {
		item, err := ProviderInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			providersField = append(providersField, *item)
		}
	}
	st.Providers = providersField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string `tf:"-"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListRecipientsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRecipientsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListRecipientsRequestToPb(st *ListRecipientsRequest) (*sharingpb.ListRecipientsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListRecipientsRequestPb{}
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListRecipientsRequestFromPb(pb *sharingpb.ListRecipientsRequestPb) (*ListRecipientsRequest, error) {
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

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of recipient information objects.
	// Wire name: 'recipients'
	Recipients      []RecipientInfo ``
	ForceSendFields []string        `tf:"-"`
}

func (s *ListRecipientsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRecipientsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListRecipientsResponseToPb(st *ListRecipientsResponse) (*sharingpb.ListRecipientsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListRecipientsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var recipientsPb []sharingpb.RecipientInfoPb
	for _, item := range st.Recipients {
		itemPb, err := RecipientInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			recipientsPb = append(recipientsPb, *itemPb)
		}
	}
	pb.Recipients = recipientsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListRecipientsResponseFromPb(pb *sharingpb.ListRecipientsResponsePb) (*ListRecipientsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRecipientsResponse{}
	st.NextPageToken = pb.NextPageToken

	var recipientsField []RecipientInfo
	for _, itemPb := range pb.Recipients {
		item, err := RecipientInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			recipientsField = append(recipientsField, *item)
		}
	}
	st.Recipients = recipientsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListSharesRequest struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Name of the provider in which to list shares.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListSharesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSharesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSharesRequestToPb(st *ListSharesRequest) (*sharingpb.ListSharesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListSharesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.Name = st.Name
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSharesRequestFromPb(pb *sharingpb.ListSharesRequestPb) (*ListSharesRequest, error) {
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

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// An array of data share information objects.
	// Wire name: 'shares'
	Shares          []ShareInfo ``
	ForceSendFields []string    `tf:"-"`
}

func (s *ListSharesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSharesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListSharesResponseToPb(st *ListSharesResponse) (*sharingpb.ListSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListSharesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var sharesPb []sharingpb.ShareInfoPb
	for _, item := range st.Shares {
		itemPb, err := ShareInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sharesPb = append(sharesPb, *itemPb)
		}
	}
	pb.Shares = sharesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListSharesResponseFromPb(pb *sharingpb.ListSharesResponsePb) (*ListSharesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSharesResponse{}
	st.NextPageToken = pb.NextPageToken

	var sharesField []ShareInfo
	for _, itemPb := range pb.Shares {
		item, err := ShareInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sharesField = append(sharesField, *item)
		}
	}
	st.Shares = sharesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type NotebookFile struct {
	// The comment of the notebook file.
	// Wire name: 'comment'
	Comment string ``
	// The id of the notebook file.
	// Wire name: 'id'
	Id string ``
	// Name of the notebook file.
	// Wire name: 'name'
	Name string ``
	// The name of the share that the notebook file belongs to.
	// Wire name: 'share'
	Share string ``
	// The id of the share that the notebook file belongs to.
	// Wire name: 'share_id'
	ShareId string ``
	// The tags of the notebook file.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue ``
	ForceSendFields []string              `tf:"-"`
}

func (s *NotebookFile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookFile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NotebookFileToPb(st *NotebookFile) (*sharingpb.NotebookFilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.NotebookFilePb{}
	pb.Comment = st.Comment
	pb.Id = st.Id
	pb.Name = st.Name
	pb.Share = st.Share
	pb.ShareId = st.ShareId

	var tagsPb []catalogpb.TagKeyValuePb
	for _, item := range st.Tags {
		itemPb, err := catalog.TagKeyValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NotebookFileFromPb(pb *sharingpb.NotebookFilePb) (*NotebookFile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookFile{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	st.Name = pb.Name
	st.Share = pb.Share
	st.ShareId = pb.ShareId

	var tagsField []catalog.TagKeyValue
	for _, itemPb := range pb.Tags {
		item, err := catalog.TagKeyValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Specifies the policy to use for validating OIDC claims in your federated
// tokens from Delta Sharing Clients. Refer to
// https://docs.databricks.com/en/delta-sharing/create-recipient-oidc-fed for
// more details.
type OidcFederationPolicy struct {
	// The allowed token audiences, as specified in the 'aud' claim of federated
	// tokens. The audience identifier is intended to represent the recipient of
	// the token. Can be any non-empty string value. As long as the audience in
	// the token matches at least one audience in the policy,
	// Wire name: 'audiences'
	Audiences []string ``
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	// Wire name: 'issuer'
	Issuer string ``
	// The required token subject, as specified in the subject claim of
	// federated tokens. The subject claim identifies the identity of the user
	// or machine accessing the resource. Examples for Entra ID (AAD): - U2M
	// flow (group access): If the subject claim is `groups`, this must be the
	// Object ID of the group in Entra ID. - U2M flow (user access): If the
	// subject claim is `oid`, this must be the Object ID of the user in Entra
	// ID. - M2M flow (OAuth App access): If the subject claim is `azp`, this
	// must be the client ID of the OAuth app registered in Entra ID.
	// Wire name: 'subject'
	Subject string ``
	// The claim that contains the subject of the token. Depending on the
	// identity provider and the use case (U2M or M2M), this can vary: - For
	// Entra ID (AAD): * U2M flow (group access): Use `groups`. * U2M flow (user
	// access): Use `oid`. * M2M flow (OAuth App access): Use `azp`. - For other
	// IdPs, refer to the specific IdP documentation.
	//
	// Supported `subject_claim` values are: - `oid`: Object ID of the user. -
	// `azp`: Client ID of the OAuth app. - `groups`: Object ID of the group. -
	// `sub`: Subject identifier for other use cases.
	// Wire name: 'subject_claim'
	SubjectClaim string ``
}

func OidcFederationPolicyToPb(st *OidcFederationPolicy) (*sharingpb.OidcFederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.OidcFederationPolicyPb{}
	pb.Audiences = st.Audiences
	pb.Issuer = st.Issuer
	pb.Subject = st.Subject
	pb.SubjectClaim = st.SubjectClaim

	return pb, nil
}

func OidcFederationPolicyFromPb(pb *sharingpb.OidcFederationPolicyPb) (*OidcFederationPolicy, error) {
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

type Partition struct {
	// An array of partition values.
	// Wire name: 'values'
	Values []PartitionValue ``
}

func PartitionToPb(st *Partition) (*sharingpb.PartitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.PartitionPb{}

	var valuesPb []sharingpb.PartitionValuePb
	for _, item := range st.Values {
		itemPb, err := PartitionValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			valuesPb = append(valuesPb, *itemPb)
		}
	}
	pb.Values = valuesPb

	return pb, nil
}

func PartitionFromPb(pb *sharingpb.PartitionPb) (*Partition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Partition{}

	var valuesField []PartitionValue
	for _, itemPb := range pb.Values {
		item, err := PartitionValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			valuesField = append(valuesField, *item)
		}
	}
	st.Values = valuesField

	return st, nil
}

type PartitionValue struct {
	// The name of the partition column.
	// Wire name: 'name'
	Name string ``
	// The operator to apply for the value.
	// Wire name: 'op'
	Op PartitionValueOp ``
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	// Wire name: 'recipient_property_key'
	RecipientPropertyKey string ``
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PartitionValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PartitionValueToPb(st *PartitionValue) (*sharingpb.PartitionValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.PartitionValuePb{}
	pb.Name = st.Name
	opPb, err := PartitionValueOpToPb(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}
	pb.RecipientPropertyKey = st.RecipientPropertyKey
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PartitionValueFromPb(pb *sharingpb.PartitionValuePb) (*PartitionValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartitionValue{}
	st.Name = pb.Name
	opField, err := PartitionValueOpFromPb(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	st.RecipientPropertyKey = pb.RecipientPropertyKey
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PartitionValueOp string

const PartitionValueOpEqual PartitionValueOp = `EQUAL`

const PartitionValueOpLike PartitionValueOp = `LIKE`

// String representation for [fmt.Print]
func (f *PartitionValueOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PartitionValueOp) Set(v string) error {
	switch v {
	case `EQUAL`, `LIKE`:
		*f = PartitionValueOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "LIKE"`, v)
	}
}

// Values returns all possible values for PartitionValueOp.
//
// There is no guarantee on the order of the values in the slice.
func (f *PartitionValueOp) Values() []PartitionValueOp {
	return []PartitionValueOp{
		PartitionValueOpEqual,
		PartitionValueOpLike,
	}
}

// Type always returns PartitionValueOp to satisfy [pflag.Value] interface
func (f *PartitionValueOp) Type() string {
	return "PartitionValueOp"
}

func PartitionValueOpToPb(st *PartitionValueOp) (*sharingpb.PartitionValueOpPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.PartitionValueOpPb(*st)
	return &pb, nil
}

func PartitionValueOpFromPb(pb *sharingpb.PartitionValueOpPb) (*PartitionValueOp, error) {
	if pb == nil {
		return nil, nil
	}
	st := PartitionValueOp(*pb)
	return &st, nil
}

type PermissionsChange struct {
	// The set of privileges to add.
	// Wire name: 'add'
	Add []string ``
	// The principal whose privileges we are changing. Only one of principal or
	// principal_id should be specified, never both at the same time.
	// Wire name: 'principal'
	Principal string ``
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove          []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *PermissionsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PermissionsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PermissionsChangeToPb(st *PermissionsChange) (*sharingpb.PermissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.PermissionsChangePb{}
	pb.Add = st.Add
	pb.Principal = st.Principal
	pb.Remove = st.Remove

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PermissionsChangeFromPb(pb *sharingpb.PermissionsChangePb) (*PermissionsChange, error) {
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

type Privilege string

const PrivilegeAccess Privilege = `ACCESS`

const PrivilegeAllPrivileges Privilege = `ALL_PRIVILEGES`

const PrivilegeApplyTag Privilege = `APPLY_TAG`

const PrivilegeCreate Privilege = `CREATE`

const PrivilegeCreateCatalog Privilege = `CREATE_CATALOG`

const PrivilegeCreateConnection Privilege = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation Privilege = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable Privilege = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume Privilege = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog Privilege = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable Privilege = `CREATE_FOREIGN_SECURABLE`

const PrivilegeCreateFunction Privilege = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage Privilege = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView Privilege = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel Privilege = `CREATE_MODEL`

const PrivilegeCreateProvider Privilege = `CREATE_PROVIDER`

const PrivilegeCreateRecipient Privilege = `CREATE_RECIPIENT`

const PrivilegeCreateSchema Privilege = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential Privilege = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare Privilege = `CREATE_SHARE`

const PrivilegeCreateStorageCredential Privilege = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable Privilege = `CREATE_TABLE`

const PrivilegeCreateView Privilege = `CREATE_VIEW`

const PrivilegeCreateVolume Privilege = `CREATE_VOLUME`

const PrivilegeExecute Privilege = `EXECUTE`

const PrivilegeManage Privilege = `MANAGE`

const PrivilegeManageAllowlist Privilege = `MANAGE_ALLOWLIST`

const PrivilegeModify Privilege = `MODIFY`

const PrivilegeReadFiles Privilege = `READ_FILES`

const PrivilegeReadPrivateFiles Privilege = `READ_PRIVATE_FILES`

const PrivilegeReadVolume Privilege = `READ_VOLUME`

const PrivilegeRefresh Privilege = `REFRESH`

const PrivilegeSelect Privilege = `SELECT`

const PrivilegeSetSharePermission Privilege = `SET_SHARE_PERMISSION`

const PrivilegeUsage Privilege = `USAGE`

const PrivilegeUseCatalog Privilege = `USE_CATALOG`

const PrivilegeUseConnection Privilege = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets Privilege = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider Privilege = `USE_PROVIDER`

const PrivilegeUseRecipient Privilege = `USE_RECIPIENT`

const PrivilegeUseSchema Privilege = `USE_SCHEMA`

const PrivilegeUseShare Privilege = `USE_SHARE`

const PrivilegeWriteFiles Privilege = `WRITE_FILES`

const PrivilegeWritePrivateFiles Privilege = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume Privilege = `WRITE_VOLUME`

// String representation for [fmt.Print]
func (f *Privilege) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Privilege) Set(v string) error {
	switch v {
	case `ACCESS`, `ALL_PRIVILEGES`, `APPLY_TAG`, `CREATE`, `CREATE_CATALOG`, `CREATE_CONNECTION`, `CREATE_EXTERNAL_LOCATION`, `CREATE_EXTERNAL_TABLE`, `CREATE_EXTERNAL_VOLUME`, `CREATE_FOREIGN_CATALOG`, `CREATE_FOREIGN_SECURABLE`, `CREATE_FUNCTION`, `CREATE_MANAGED_STORAGE`, `CREATE_MATERIALIZED_VIEW`, `CREATE_MODEL`, `CREATE_PROVIDER`, `CREATE_RECIPIENT`, `CREATE_SCHEMA`, `CREATE_SERVICE_CREDENTIAL`, `CREATE_SHARE`, `CREATE_STORAGE_CREDENTIAL`, `CREATE_TABLE`, `CREATE_VIEW`, `CREATE_VOLUME`, `EXECUTE`, `MANAGE`, `MANAGE_ALLOWLIST`, `MODIFY`, `READ_FILES`, `READ_PRIVATE_FILES`, `READ_VOLUME`, `REFRESH`, `SELECT`, `SET_SHARE_PERMISSION`, `USAGE`, `USE_CATALOG`, `USE_CONNECTION`, `USE_MARKETPLACE_ASSETS`, `USE_PROVIDER`, `USE_RECIPIENT`, `USE_SCHEMA`, `USE_SHARE`, `WRITE_FILES`, `WRITE_PRIVATE_FILES`, `WRITE_VOLUME`:
		*f = Privilege(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCESS", "ALL_PRIVILEGES", "APPLY_TAG", "CREATE", "CREATE_CATALOG", "CREATE_CONNECTION", "CREATE_EXTERNAL_LOCATION", "CREATE_EXTERNAL_TABLE", "CREATE_EXTERNAL_VOLUME", "CREATE_FOREIGN_CATALOG", "CREATE_FOREIGN_SECURABLE", "CREATE_FUNCTION", "CREATE_MANAGED_STORAGE", "CREATE_MATERIALIZED_VIEW", "CREATE_MODEL", "CREATE_PROVIDER", "CREATE_RECIPIENT", "CREATE_SCHEMA", "CREATE_SERVICE_CREDENTIAL", "CREATE_SHARE", "CREATE_STORAGE_CREDENTIAL", "CREATE_TABLE", "CREATE_VIEW", "CREATE_VOLUME", "EXECUTE", "MANAGE", "MANAGE_ALLOWLIST", "MODIFY", "READ_FILES", "READ_PRIVATE_FILES", "READ_VOLUME", "REFRESH", "SELECT", "SET_SHARE_PERMISSION", "USAGE", "USE_CATALOG", "USE_CONNECTION", "USE_MARKETPLACE_ASSETS", "USE_PROVIDER", "USE_RECIPIENT", "USE_SCHEMA", "USE_SHARE", "WRITE_FILES", "WRITE_PRIVATE_FILES", "WRITE_VOLUME"`, v)
	}
}

// Values returns all possible values for Privilege.
//
// There is no guarantee on the order of the values in the slice.
func (f *Privilege) Values() []Privilege {
	return []Privilege{
		PrivilegeAccess,
		PrivilegeAllPrivileges,
		PrivilegeApplyTag,
		PrivilegeCreate,
		PrivilegeCreateCatalog,
		PrivilegeCreateConnection,
		PrivilegeCreateExternalLocation,
		PrivilegeCreateExternalTable,
		PrivilegeCreateExternalVolume,
		PrivilegeCreateForeignCatalog,
		PrivilegeCreateForeignSecurable,
		PrivilegeCreateFunction,
		PrivilegeCreateManagedStorage,
		PrivilegeCreateMaterializedView,
		PrivilegeCreateModel,
		PrivilegeCreateProvider,
		PrivilegeCreateRecipient,
		PrivilegeCreateSchema,
		PrivilegeCreateServiceCredential,
		PrivilegeCreateShare,
		PrivilegeCreateStorageCredential,
		PrivilegeCreateTable,
		PrivilegeCreateView,
		PrivilegeCreateVolume,
		PrivilegeExecute,
		PrivilegeManage,
		PrivilegeManageAllowlist,
		PrivilegeModify,
		PrivilegeReadFiles,
		PrivilegeReadPrivateFiles,
		PrivilegeReadVolume,
		PrivilegeRefresh,
		PrivilegeSelect,
		PrivilegeSetSharePermission,
		PrivilegeUsage,
		PrivilegeUseCatalog,
		PrivilegeUseConnection,
		PrivilegeUseMarketplaceAssets,
		PrivilegeUseProvider,
		PrivilegeUseRecipient,
		PrivilegeUseSchema,
		PrivilegeUseShare,
		PrivilegeWriteFiles,
		PrivilegeWritePrivateFiles,
		PrivilegeWriteVolume,
	}
}

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

func PrivilegeToPb(st *Privilege) (*sharingpb.PrivilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.PrivilegePb(*st)
	return &pb, nil
}

func PrivilegeFromPb(pb *sharingpb.PrivilegePb) (*Privilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := Privilege(*pb)
	return &st, nil
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name). For deleted principals,
	// `principal` is empty while `principal_id` is populated.
	// Wire name: 'principal'
	Principal string ``
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges      []Privilege ``
	ForceSendFields []string    `tf:"-"`
}

func (s *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PrivilegeAssignmentToPb(st *PrivilegeAssignment) (*sharingpb.PrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.PrivilegeAssignmentPb{}
	pb.Principal = st.Principal

	var privilegesPb []sharingpb.PrivilegePb
	for _, item := range st.Privileges {
		itemPb, err := PrivilegeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegesPb = append(privilegesPb, *itemPb)
		}
	}
	pb.Privileges = privilegesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PrivilegeAssignmentFromPb(pb *sharingpb.PrivilegeAssignmentPb) (*PrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivilegeAssignment{}
	st.Principal = pb.Principal

	var privilegesField []Privilege
	for _, itemPb := range pb.Privileges {
		item, err := PrivilegeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegesField = append(privilegesField, *item)
		}
	}
	st.Privileges = privilegesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ProviderInfo struct {

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType ``
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'cloud'
	Cloud string ``
	// Description about the provider.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this Provider was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of Provider creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_provider_global_metastore_id'
	DataProviderGlobalMetastoreId string ``
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// The name of the Provider.
	// Wire name: 'name'
	Name string ``
	// Username of Provider owner.
	// Wire name: 'owner'
	Owner string ``
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	// Wire name: 'recipient_profile'
	RecipientProfile *RecipientProfile ``
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string ``
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'region'
	Region string ``
	// Time at which this Provider was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of user who last modified Provider.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ProviderInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ProviderInfoToPb(st *ProviderInfo) (*sharingpb.ProviderInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ProviderInfoPb{}
	authenticationTypePb, err := AuthenticationTypeToPb(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}
	pb.Cloud = st.Cloud
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DataProviderGlobalMetastoreId = st.DataProviderGlobalMetastoreId
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	recipientProfilePb, err := RecipientProfileToPb(st.RecipientProfile)
	if err != nil {
		return nil, err
	}
	if recipientProfilePb != nil {
		pb.RecipientProfile = recipientProfilePb
	}
	pb.RecipientProfileStr = st.RecipientProfileStr
	pb.Region = st.Region
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ProviderInfoFromPb(pb *sharingpb.ProviderInfoPb) (*ProviderInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderInfo{}
	authenticationTypeField, err := AuthenticationTypeFromPb(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	st.Cloud = pb.Cloud
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataProviderGlobalMetastoreId = pb.DataProviderGlobalMetastoreId
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	recipientProfileField, err := RecipientProfileFromPb(pb.RecipientProfile)
	if err != nil {
		return nil, err
	}
	if recipientProfileField != nil {
		st.RecipientProfile = recipientProfileField
	}
	st.RecipientProfileStr = pb.RecipientProfileStr
	st.Region = pb.Region
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ProviderShare struct {
	// The name of the Provider Share.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ProviderShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ProviderShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ProviderShareToPb(st *ProviderShare) (*sharingpb.ProviderSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ProviderSharePb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ProviderShareFromPb(pb *sharingpb.ProviderSharePb) (*ProviderShare, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderShare{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	// Wire name: 'activated'
	Activated bool ``
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string ``

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType ``
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'cloud'
	Cloud string ``
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this recipient was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of recipient creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string ``
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList ``
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'metastore_id'
	MetastoreId string ``
	// Name of Recipient.
	// Wire name: 'name'
	Name string ``
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string ``
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs ``
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'region'
	Region string ``
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'sharing_code'
	SharingCode string ``
	// This field is only present when the __authentication_type__ is **TOKEN**.
	// Wire name: 'tokens'
	Tokens []RecipientTokenInfo ``
	// Time at which the recipient was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of recipient updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RecipientInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RecipientInfoToPb(st *RecipientInfo) (*sharingpb.RecipientInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RecipientInfoPb{}
	pb.Activated = st.Activated
	pb.ActivationUrl = st.ActivationUrl
	authenticationTypePb, err := AuthenticationTypeToPb(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}
	pb.Cloud = st.Cloud
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.ExpirationTime = st.ExpirationTime
	ipAccessListPb, err := IpAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	propertiesKvpairsPb, err := SecurablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}
	pb.Region = st.Region
	pb.SharingCode = st.SharingCode

	var tokensPb []sharingpb.RecipientTokenInfoPb
	for _, item := range st.Tokens {
		itemPb, err := RecipientTokenInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokensPb = append(tokensPb, *itemPb)
		}
	}
	pb.Tokens = tokensPb
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RecipientInfoFromPb(pb *sharingpb.RecipientInfoPb) (*RecipientInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RecipientInfo{}
	st.Activated = pb.Activated
	st.ActivationUrl = pb.ActivationUrl
	authenticationTypeField, err := AuthenticationTypeFromPb(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	st.Cloud = pb.Cloud
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataRecipientGlobalMetastoreId = pb.DataRecipientGlobalMetastoreId
	st.ExpirationTime = pb.ExpirationTime
	ipAccessListField, err := IpAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	propertiesKvpairsField, err := SecurablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsField != nil {
		st.PropertiesKvpairs = propertiesKvpairsField
	}
	st.Region = pb.Region
	st.SharingCode = pb.SharingCode

	var tokensField []RecipientTokenInfo
	for _, itemPb := range pb.Tokens {
		item, err := RecipientTokenInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokensField = append(tokensField, *item)
		}
	}
	st.Tokens = tokensField
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	// Wire name: 'bearer_token'
	BearerToken string ``
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string ``
	// The version number of the recipient's credentials on a share.
	// Wire name: 'share_credentials_version'
	ShareCredentialsVersion int      ``
	ForceSendFields         []string `tf:"-"`
}

func (s *RecipientProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RecipientProfileToPb(st *RecipientProfile) (*sharingpb.RecipientProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RecipientProfilePb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RecipientProfileFromPb(pb *sharingpb.RecipientProfilePb) (*RecipientProfile, error) {
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

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string ``
	// Time at which this recipient token was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of recipient token creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``
	// Unique ID of the recipient token.
	// Wire name: 'id'
	Id string ``
	// Time at which this recipient token was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of recipient token updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RecipientTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RecipientTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RecipientTokenInfoToPb(st *RecipientTokenInfo) (*sharingpb.RecipientTokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RecipientTokenInfoPb{}
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

func RecipientTokenInfoFromPb(pb *sharingpb.RecipientTokenInfoPb) (*RecipientTokenInfo, error) {
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

type RegisteredModelAlias struct {
	// Name of the alias.
	// Wire name: 'alias_name'
	AliasName string ``
	// Numeric model version that alias will reference.
	// Wire name: 'version_num'
	VersionNum      int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RegisteredModelAliasToPb(st *RegisteredModelAlias) (*sharingpb.RegisteredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RegisteredModelAliasPb{}
	pb.AliasName = st.AliasName
	pb.VersionNum = st.VersionNum

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RegisteredModelAliasFromPb(pb *sharingpb.RegisteredModelAliasPb) (*RegisteredModelAlias, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAlias{}
	st.AliasName = pb.AliasName
	st.VersionNum = pb.VersionNum

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	// Wire name: 'activation_url'
	ActivationUrl string `tf:"-"`
}

func RetrieveTokenRequestToPb(st *RetrieveTokenRequest) (*sharingpb.RetrieveTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RetrieveTokenRequestPb{}
	pb.ActivationUrl = st.ActivationUrl

	return pb, nil
}

func RetrieveTokenRequestFromPb(pb *sharingpb.RetrieveTokenRequestPb) (*RetrieveTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RetrieveTokenRequest{}
	st.ActivationUrl = pb.ActivationUrl

	return st, nil
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	// Wire name: 'bearerToken'
	BearerToken string ``
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string ``
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expirationTime'
	ExpirationTime string ``
	// These field names must follow the delta sharing protocol.
	// Wire name: 'shareCredentialsVersion'
	ShareCredentialsVersion int      ``
	ForceSendFields         []string `tf:"-"`
}

func (s *RetrieveTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RetrieveTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RetrieveTokenResponseToPb(st *RetrieveTokenResponse) (*sharingpb.RetrieveTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RetrieveTokenResponsePb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ExpirationTime = st.ExpirationTime
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RetrieveTokenResponseFromPb(pb *sharingpb.RetrieveTokenResponsePb) (*RetrieveTokenResponse, error) {
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

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	// Wire name: 'existing_token_expire_in_seconds'
	ExistingTokenExpireInSeconds int64 ``
	// The name of the Recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func RotateRecipientTokenToPb(st *RotateRecipientToken) (*sharingpb.RotateRecipientTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RotateRecipientTokenPb{}
	pb.ExistingTokenExpireInSeconds = st.ExistingTokenExpireInSeconds
	pb.Name = st.Name

	return pb, nil
}

func RotateRecipientTokenFromPb(pb *sharingpb.RotateRecipientTokenPb) (*RotateRecipientToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RotateRecipientToken{}
	st.ExistingTokenExpireInSeconds = pb.ExistingTokenExpireInSeconds
	st.Name = pb.Name

	return st, nil
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string ``
}

func SecurablePropertiesKvPairsToPb(st *SecurablePropertiesKvPairs) (*sharingpb.SecurablePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.SecurablePropertiesKvPairsPb{}
	pb.Properties = st.Properties

	return pb, nil
}

func SecurablePropertiesKvPairsFromPb(pb *sharingpb.SecurablePropertiesKvPairsPb) (*SecurablePropertiesKvPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecurablePropertiesKvPairs{}
	st.Properties = pb.Properties

	return st, nil
}

type ShareInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// Time at which this share was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// Username of share creator.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Name of the share.
	// Wire name: 'name'
	Name string ``
	// A list of shared data objects within the share.
	// Wire name: 'objects'
	Objects []SharedDataObject ``
	// Username of current owner of share.
	// Wire name: 'owner'
	Owner string ``
	// Storage Location URL (full path) for the share.
	// Wire name: 'storage_location'
	StorageLocation string ``
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string ``
	// Time at which this share was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 ``
	// Username of share updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ShareInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ShareInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ShareInfoToPb(st *ShareInfo) (*sharingpb.ShareInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ShareInfoPb{}
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Name = st.Name

	var objectsPb []sharingpb.SharedDataObjectPb
	for _, item := range st.Objects {
		itemPb, err := SharedDataObjectToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			objectsPb = append(objectsPb, *itemPb)
		}
	}
	pb.Objects = objectsPb
	pb.Owner = st.Owner
	pb.StorageLocation = st.StorageLocation
	pb.StorageRoot = st.StorageRoot
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ShareInfoFromPb(pb *sharingpb.ShareInfoPb) (*ShareInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareInfo{}
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Name = pb.Name

	var objectsField []SharedDataObject
	for _, itemPb := range pb.Objects {
		item, err := SharedDataObjectFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			objectsField = append(objectsField, *item)
		}
	}
	st.Objects = objectsField
	st.Owner = pb.Owner
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SharePermissionsRequest struct {
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// The name of the Recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *SharePermissionsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharePermissionsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SharePermissionsRequestToPb(st *SharePermissionsRequest) (*sharingpb.SharePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.SharePermissionsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.Name = st.Name
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SharePermissionsRequestFromPb(pb *sharingpb.SharePermissionsRequestPb) (*SharePermissionsRequest, error) {
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

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment ``
	// The share name.
	// Wire name: 'share_name'
	ShareName       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ShareToPrivilegeAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ShareToPrivilegeAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ShareToPrivilegeAssignmentToPb(st *ShareToPrivilegeAssignment) (*sharingpb.ShareToPrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ShareToPrivilegeAssignmentPb{}

	var privilegeAssignmentsPb []sharingpb.PrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := PrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb
	pb.ShareName = st.ShareName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ShareToPrivilegeAssignmentFromPb(pb *sharingpb.ShareToPrivilegeAssignmentPb) (*ShareToPrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareToPrivilegeAssignment{}

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := PrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField
	st.ShareName = pb.ShareName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	// Wire name: 'added_at'
	AddedAt int64 ``
	// Username of the sharer.
	// Wire name: 'added_by'
	AddedBy string ``
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	// Wire name: 'cdf_enabled'
	CdfEnabled bool ``
	// A user-provided comment when adding the data object to the share.
	// Wire name: 'comment'
	Comment string ``
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	// Wire name: 'content'
	Content string ``
	// The type of the data object.
	// Wire name: 'data_object_type'
	DataObjectType SharedDataObjectDataObjectType ``
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	// Wire name: 'history_data_sharing_status'
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus ``
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`,
	// Wire name: 'name'
	Name string ``
	// Array of partitions for the shared data.
	// Wire name: 'partitions'
	Partitions []Partition ``
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	// Wire name: 'shared_as'
	SharedAs string ``
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	// Wire name: 'start_version'
	StartVersion int64 ``
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	// Wire name: 'status'
	Status SharedDataObjectStatus ``
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
	// Wire name: 'string_shared_as'
	StringSharedAs  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SharedDataObject) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SharedDataObject) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SharedDataObjectToPb(st *SharedDataObject) (*sharingpb.SharedDataObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.SharedDataObjectPb{}
	pb.AddedAt = st.AddedAt
	pb.AddedBy = st.AddedBy
	pb.CdfEnabled = st.CdfEnabled
	pb.Comment = st.Comment
	pb.Content = st.Content
	dataObjectTypePb, err := SharedDataObjectDataObjectTypeToPb(&st.DataObjectType)
	if err != nil {
		return nil, err
	}
	if dataObjectTypePb != nil {
		pb.DataObjectType = *dataObjectTypePb
	}
	historyDataSharingStatusPb, err := SharedDataObjectHistoryDataSharingStatusToPb(&st.HistoryDataSharingStatus)
	if err != nil {
		return nil, err
	}
	if historyDataSharingStatusPb != nil {
		pb.HistoryDataSharingStatus = *historyDataSharingStatusPb
	}
	pb.Name = st.Name

	var partitionsPb []sharingpb.PartitionPb
	for _, item := range st.Partitions {
		itemPb, err := PartitionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			partitionsPb = append(partitionsPb, *itemPb)
		}
	}
	pb.Partitions = partitionsPb
	pb.SharedAs = st.SharedAs
	pb.StartVersion = st.StartVersion
	statusPb, err := SharedDataObjectStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StringSharedAs = st.StringSharedAs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SharedDataObjectFromPb(pb *sharingpb.SharedDataObjectPb) (*SharedDataObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObject{}
	st.AddedAt = pb.AddedAt
	st.AddedBy = pb.AddedBy
	st.CdfEnabled = pb.CdfEnabled
	st.Comment = pb.Comment
	st.Content = pb.Content
	dataObjectTypeField, err := SharedDataObjectDataObjectTypeFromPb(&pb.DataObjectType)
	if err != nil {
		return nil, err
	}
	if dataObjectTypeField != nil {
		st.DataObjectType = *dataObjectTypeField
	}
	historyDataSharingStatusField, err := SharedDataObjectHistoryDataSharingStatusFromPb(&pb.HistoryDataSharingStatus)
	if err != nil {
		return nil, err
	}
	if historyDataSharingStatusField != nil {
		st.HistoryDataSharingStatus = *historyDataSharingStatusField
	}
	st.Name = pb.Name

	var partitionsField []Partition
	for _, itemPb := range pb.Partitions {
		item, err := PartitionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			partitionsField = append(partitionsField, *item)
		}
	}
	st.Partitions = partitionsField
	st.SharedAs = pb.SharedAs
	st.StartVersion = pb.StartVersion
	statusField, err := SharedDataObjectStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StringSharedAs = pb.StringSharedAs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SharedDataObjectDataObjectType string

const SharedDataObjectDataObjectTypeFeatureSpec SharedDataObjectDataObjectType = `FEATURE_SPEC`

const SharedDataObjectDataObjectTypeFunction SharedDataObjectDataObjectType = `FUNCTION`

const SharedDataObjectDataObjectTypeMaterializedView SharedDataObjectDataObjectType = `MATERIALIZED_VIEW`

const SharedDataObjectDataObjectTypeModel SharedDataObjectDataObjectType = `MODEL`

const SharedDataObjectDataObjectTypeNotebookFile SharedDataObjectDataObjectType = `NOTEBOOK_FILE`

const SharedDataObjectDataObjectTypeSchema SharedDataObjectDataObjectType = `SCHEMA`

const SharedDataObjectDataObjectTypeStreamingTable SharedDataObjectDataObjectType = `STREAMING_TABLE`

const SharedDataObjectDataObjectTypeTable SharedDataObjectDataObjectType = `TABLE`

const SharedDataObjectDataObjectTypeView SharedDataObjectDataObjectType = `VIEW`

// String representation for [fmt.Print]
func (f *SharedDataObjectDataObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectDataObjectType) Set(v string) error {
	switch v {
	case `FEATURE_SPEC`, `FUNCTION`, `MATERIALIZED_VIEW`, `MODEL`, `NOTEBOOK_FILE`, `SCHEMA`, `STREAMING_TABLE`, `TABLE`, `VIEW`:
		*f = SharedDataObjectDataObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FEATURE_SPEC", "FUNCTION", "MATERIALIZED_VIEW", "MODEL", "NOTEBOOK_FILE", "SCHEMA", "STREAMING_TABLE", "TABLE", "VIEW"`, v)
	}
}

// Values returns all possible values for SharedDataObjectDataObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectDataObjectType) Values() []SharedDataObjectDataObjectType {
	return []SharedDataObjectDataObjectType{
		SharedDataObjectDataObjectTypeFeatureSpec,
		SharedDataObjectDataObjectTypeFunction,
		SharedDataObjectDataObjectTypeMaterializedView,
		SharedDataObjectDataObjectTypeModel,
		SharedDataObjectDataObjectTypeNotebookFile,
		SharedDataObjectDataObjectTypeSchema,
		SharedDataObjectDataObjectTypeStreamingTable,
		SharedDataObjectDataObjectTypeTable,
		SharedDataObjectDataObjectTypeView,
	}
}

// Type always returns SharedDataObjectDataObjectType to satisfy [pflag.Value] interface
func (f *SharedDataObjectDataObjectType) Type() string {
	return "SharedDataObjectDataObjectType"
}

func SharedDataObjectDataObjectTypeToPb(st *SharedDataObjectDataObjectType) (*sharingpb.SharedDataObjectDataObjectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.SharedDataObjectDataObjectTypePb(*st)
	return &pb, nil
}

func SharedDataObjectDataObjectTypeFromPb(pb *sharingpb.SharedDataObjectDataObjectTypePb) (*SharedDataObjectDataObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectDataObjectType(*pb)
	return &st, nil
}

type SharedDataObjectHistoryDataSharingStatus string

const SharedDataObjectHistoryDataSharingStatusDisabled SharedDataObjectHistoryDataSharingStatus = `DISABLED`

const SharedDataObjectHistoryDataSharingStatusEnabled SharedDataObjectHistoryDataSharingStatus = `ENABLED`

// String representation for [fmt.Print]
func (f *SharedDataObjectHistoryDataSharingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectHistoryDataSharingStatus) Set(v string) error {
	switch v {
	case `DISABLED`, `ENABLED`:
		*f = SharedDataObjectHistoryDataSharingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLED", "ENABLED"`, v)
	}
}

// Values returns all possible values for SharedDataObjectHistoryDataSharingStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectHistoryDataSharingStatus) Values() []SharedDataObjectHistoryDataSharingStatus {
	return []SharedDataObjectHistoryDataSharingStatus{
		SharedDataObjectHistoryDataSharingStatusDisabled,
		SharedDataObjectHistoryDataSharingStatusEnabled,
	}
}

// Type always returns SharedDataObjectHistoryDataSharingStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectHistoryDataSharingStatus) Type() string {
	return "SharedDataObjectHistoryDataSharingStatus"
}

func SharedDataObjectHistoryDataSharingStatusToPb(st *SharedDataObjectHistoryDataSharingStatus) (*sharingpb.SharedDataObjectHistoryDataSharingStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.SharedDataObjectHistoryDataSharingStatusPb(*st)
	return &pb, nil
}

func SharedDataObjectHistoryDataSharingStatusFromPb(pb *sharingpb.SharedDataObjectHistoryDataSharingStatusPb) (*SharedDataObjectHistoryDataSharingStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectHistoryDataSharingStatus(*pb)
	return &st, nil
}

type SharedDataObjectStatus string

const SharedDataObjectStatusActive SharedDataObjectStatus = `ACTIVE`

const SharedDataObjectStatusPermissionDenied SharedDataObjectStatus = `PERMISSION_DENIED`

// String representation for [fmt.Print]
func (f *SharedDataObjectStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `PERMISSION_DENIED`:
		*f = SharedDataObjectStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "PERMISSION_DENIED"`, v)
	}
}

// Values returns all possible values for SharedDataObjectStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectStatus) Values() []SharedDataObjectStatus {
	return []SharedDataObjectStatus{
		SharedDataObjectStatusActive,
		SharedDataObjectStatusPermissionDenied,
	}
}

// Type always returns SharedDataObjectStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectStatus) Type() string {
	return "SharedDataObjectStatus"
}

func SharedDataObjectStatusToPb(st *SharedDataObjectStatus) (*sharingpb.SharedDataObjectStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.SharedDataObjectStatusPb(*st)
	return &pb, nil
}

func SharedDataObjectStatusFromPb(pb *sharingpb.SharedDataObjectStatusPb) (*SharedDataObjectStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectStatus(*pb)
	return &st, nil
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	// Wire name: 'action'
	Action SharedDataObjectUpdateAction ``
	// The data object that is being added, removed, or updated. The maximum
	// number update data objects allowed is a 100.
	// Wire name: 'data_object'
	DataObject *SharedDataObject ``
}

func SharedDataObjectUpdateToPb(st *SharedDataObjectUpdate) (*sharingpb.SharedDataObjectUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.SharedDataObjectUpdatePb{}
	actionPb, err := SharedDataObjectUpdateActionToPb(&st.Action)
	if err != nil {
		return nil, err
	}
	if actionPb != nil {
		pb.Action = *actionPb
	}
	dataObjectPb, err := SharedDataObjectToPb(st.DataObject)
	if err != nil {
		return nil, err
	}
	if dataObjectPb != nil {
		pb.DataObject = dataObjectPb
	}

	return pb, nil
}

func SharedDataObjectUpdateFromPb(pb *sharingpb.SharedDataObjectUpdatePb) (*SharedDataObjectUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObjectUpdate{}
	actionField, err := SharedDataObjectUpdateActionFromPb(&pb.Action)
	if err != nil {
		return nil, err
	}
	if actionField != nil {
		st.Action = *actionField
	}
	dataObjectField, err := SharedDataObjectFromPb(pb.DataObject)
	if err != nil {
		return nil, err
	}
	if dataObjectField != nil {
		st.DataObject = dataObjectField
	}

	return st, nil
}

type SharedDataObjectUpdateAction string

const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateAction = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateAction = `REMOVE`

const SharedDataObjectUpdateActionUpdate SharedDataObjectUpdateAction = `UPDATE`

// String representation for [fmt.Print]
func (f *SharedDataObjectUpdateAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedDataObjectUpdateAction) Set(v string) error {
	switch v {
	case `ADD`, `REMOVE`, `UPDATE`:
		*f = SharedDataObjectUpdateAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD", "REMOVE", "UPDATE"`, v)
	}
}

// Values returns all possible values for SharedDataObjectUpdateAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedDataObjectUpdateAction) Values() []SharedDataObjectUpdateAction {
	return []SharedDataObjectUpdateAction{
		SharedDataObjectUpdateActionAdd,
		SharedDataObjectUpdateActionRemove,
		SharedDataObjectUpdateActionUpdate,
	}
}

// Type always returns SharedDataObjectUpdateAction to satisfy [pflag.Value] interface
func (f *SharedDataObjectUpdateAction) Type() string {
	return "SharedDataObjectUpdateAction"
}

func SharedDataObjectUpdateActionToPb(st *SharedDataObjectUpdateAction) (*sharingpb.SharedDataObjectUpdateActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.SharedDataObjectUpdateActionPb(*st)
	return &pb, nil
}

func SharedDataObjectUpdateActionFromPb(pb *sharingpb.SharedDataObjectUpdateActionPb) (*SharedDataObjectUpdateAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectUpdateAction(*pb)
	return &st, nil
}

// The SecurableKind of a delta-shared object.
type SharedSecurableKind string

const SharedSecurableKindFunctionFeatureSpec SharedSecurableKind = `FUNCTION_FEATURE_SPEC`

const SharedSecurableKindFunctionRegisteredModel SharedSecurableKind = `FUNCTION_REGISTERED_MODEL`

const SharedSecurableKindFunctionStandard SharedSecurableKind = `FUNCTION_STANDARD`

// String representation for [fmt.Print]
func (f *SharedSecurableKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SharedSecurableKind) Set(v string) error {
	switch v {
	case `FUNCTION_FEATURE_SPEC`, `FUNCTION_REGISTERED_MODEL`, `FUNCTION_STANDARD`:
		*f = SharedSecurableKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FUNCTION_FEATURE_SPEC", "FUNCTION_REGISTERED_MODEL", "FUNCTION_STANDARD"`, v)
	}
}

// Values returns all possible values for SharedSecurableKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *SharedSecurableKind) Values() []SharedSecurableKind {
	return []SharedSecurableKind{
		SharedSecurableKindFunctionFeatureSpec,
		SharedSecurableKindFunctionRegisteredModel,
		SharedSecurableKindFunctionStandard,
	}
}

// Type always returns SharedSecurableKind to satisfy [pflag.Value] interface
func (f *SharedSecurableKind) Type() string {
	return "SharedSecurableKind"
}

func SharedSecurableKindToPb(st *SharedSecurableKind) (*sharingpb.SharedSecurableKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.SharedSecurableKindPb(*st)
	return &pb, nil
}

func SharedSecurableKindFromPb(pb *sharingpb.SharedSecurableKindPb) (*SharedSecurableKind, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedSecurableKind(*pb)
	return &st, nil
}

type Table struct {
	// The comment of the table.
	// Wire name: 'comment'
	Comment string ``
	// The id of the table.
	// Wire name: 'id'
	Id string ``
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	// Wire name: 'internal_attributes'
	InternalAttributes *TableInternalAttributes ``
	// The catalog and schema of the materialized table
	// Wire name: 'materialization_namespace'
	MaterializationNamespace string ``
	// The name of a materialized table.
	// Wire name: 'materialized_table_name'
	MaterializedTableName string ``
	// The name of the table.
	// Wire name: 'name'
	Name string ``
	// The name of the schema that the table belongs to.
	// Wire name: 'schema'
	Schema string ``
	// The name of the share that the table belongs to.
	// Wire name: 'share'
	Share string ``
	// The id of the share that the table belongs to.
	// Wire name: 'share_id'
	ShareId string ``
	// The Tags of the table.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue ``
	ForceSendFields []string              `tf:"-"`
}

func (s *Table) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Table) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableToPb(st *Table) (*sharingpb.TablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.TablePb{}
	pb.Comment = st.Comment
	pb.Id = st.Id
	internalAttributesPb, err := TableInternalAttributesToPb(st.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesPb != nil {
		pb.InternalAttributes = internalAttributesPb
	}
	pb.MaterializationNamespace = st.MaterializationNamespace
	pb.MaterializedTableName = st.MaterializedTableName
	pb.Name = st.Name
	pb.Schema = st.Schema
	pb.Share = st.Share
	pb.ShareId = st.ShareId

	var tagsPb []catalogpb.TagKeyValuePb
	for _, item := range st.Tags {
		itemPb, err := catalog.TagKeyValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableFromPb(pb *sharingpb.TablePb) (*Table, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Table{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	internalAttributesField, err := TableInternalAttributesFromPb(pb.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesField != nil {
		st.InternalAttributes = internalAttributesField
	}
	st.MaterializationNamespace = pb.MaterializationNamespace
	st.MaterializedTableName = pb.MaterializedTableName
	st.Name = pb.Name
	st.Schema = pb.Schema
	st.Share = pb.Share
	st.ShareId = pb.ShareId

	var tagsField []catalog.TagKeyValue
	for _, itemPb := range pb.Tags {
		item, err := catalog.TagKeyValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type TableInternalAttributes struct {
	// Will be populated in the reconciliation response for VIEW and
	// FOREIGN_TABLE, with the value of the parent UC entity's storage_location,
	// following the same logic as getManagedEntityPath in
	// CreateStagingTableHandler, which is used to store the materialized table
	// for a shared VIEW/FOREIGN_TABLE for D2O queries. The value will be used
	// on the recipient side to be whitelisted when SEG is enabled on the
	// workspace of the recipient, to allow the recipient users to query this
	// shared VIEW/FOREIGN_TABLE.
	// Wire name: 'parent_storage_location'
	ParentStorageLocation string ``
	// The cloud storage location of a shard table with DIRECTORY_BASED_TABLE
	// type.
	// Wire name: 'storage_location'
	StorageLocation string ``
	// The type of the shared table.
	// Wire name: 'type'
	Type TableInternalAttributesSharedTableType ``
	// The view definition of a shared view. DEPRECATED.
	// Wire name: 'view_definition'
	ViewDefinition  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *TableInternalAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableInternalAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableInternalAttributesToPb(st *TableInternalAttributes) (*sharingpb.TableInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.TableInternalAttributesPb{}
	pb.ParentStorageLocation = st.ParentStorageLocation
	pb.StorageLocation = st.StorageLocation
	typePb, err := TableInternalAttributesSharedTableTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}
	pb.ViewDefinition = st.ViewDefinition

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableInternalAttributesFromPb(pb *sharingpb.TableInternalAttributesPb) (*TableInternalAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableInternalAttributes{}
	st.ParentStorageLocation = pb.ParentStorageLocation
	st.StorageLocation = pb.StorageLocation
	typeField, err := TableInternalAttributesSharedTableTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}
	st.ViewDefinition = pb.ViewDefinition

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TableInternalAttributesSharedTableType string

const TableInternalAttributesSharedTableTypeDeltaIcebergTable TableInternalAttributesSharedTableType = `DELTA_ICEBERG_TABLE`

const TableInternalAttributesSharedTableTypeDirectoryBasedTable TableInternalAttributesSharedTableType = `DIRECTORY_BASED_TABLE`

const TableInternalAttributesSharedTableTypeFileBasedTable TableInternalAttributesSharedTableType = `FILE_BASED_TABLE`

const TableInternalAttributesSharedTableTypeForeignTable TableInternalAttributesSharedTableType = `FOREIGN_TABLE`

const TableInternalAttributesSharedTableTypeMaterializedView TableInternalAttributesSharedTableType = `MATERIALIZED_VIEW`

const TableInternalAttributesSharedTableTypeStreamingTable TableInternalAttributesSharedTableType = `STREAMING_TABLE`

const TableInternalAttributesSharedTableTypeView TableInternalAttributesSharedTableType = `VIEW`

// String representation for [fmt.Print]
func (f *TableInternalAttributesSharedTableType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableInternalAttributesSharedTableType) Set(v string) error {
	switch v {
	case `DELTA_ICEBERG_TABLE`, `DIRECTORY_BASED_TABLE`, `FILE_BASED_TABLE`, `FOREIGN_TABLE`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableInternalAttributesSharedTableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELTA_ICEBERG_TABLE", "DIRECTORY_BASED_TABLE", "FILE_BASED_TABLE", "FOREIGN_TABLE", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Values returns all possible values for TableInternalAttributesSharedTableType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableInternalAttributesSharedTableType) Values() []TableInternalAttributesSharedTableType {
	return []TableInternalAttributesSharedTableType{
		TableInternalAttributesSharedTableTypeDeltaIcebergTable,
		TableInternalAttributesSharedTableTypeDirectoryBasedTable,
		TableInternalAttributesSharedTableTypeFileBasedTable,
		TableInternalAttributesSharedTableTypeForeignTable,
		TableInternalAttributesSharedTableTypeMaterializedView,
		TableInternalAttributesSharedTableTypeStreamingTable,
		TableInternalAttributesSharedTableTypeView,
	}
}

// Type always returns TableInternalAttributesSharedTableType to satisfy [pflag.Value] interface
func (f *TableInternalAttributesSharedTableType) Type() string {
	return "TableInternalAttributesSharedTableType"
}

func TableInternalAttributesSharedTableTypeToPb(st *TableInternalAttributesSharedTableType) (*sharingpb.TableInternalAttributesSharedTableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharingpb.TableInternalAttributesSharedTableTypePb(*st)
	return &pb, nil
}

func TableInternalAttributesSharedTableTypeFromPb(pb *sharingpb.TableInternalAttributesSharedTableTypePb) (*TableInternalAttributesSharedTableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableInternalAttributesSharedTableType(*pb)
	return &st, nil
}

type UpdateFederationPolicyRequest struct {
	// Name of the policy. This is the name of the current name of the policy.
	// Wire name: 'name'
	Name string `tf:"-"`

	// Wire name: 'policy'
	Policy FederationPolicy ``
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being updated.
	// Wire name: 'recipient_name'
	RecipientName string `tf:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'comment,oidc_policy.audiences'.
	// Wire name: 'update_mask'
	UpdateMask      *[]string `tf:"-"`
	ForceSendFields []string  `tf:"-"`
}

func (s *UpdateFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateFederationPolicyRequestToPb(st *UpdateFederationPolicyRequest) (*sharingpb.UpdateFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.UpdateFederationPolicyRequestPb{}
	pb.Name = st.Name
	policyPb, err := FederationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.RecipientName = st.RecipientName
	updateMaskPb, err := fieldMaskToPb(st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateFederationPolicyRequestFromPb(pb *sharingpb.UpdateFederationPolicyRequestPb) (*UpdateFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateFederationPolicyRequest{}
	st.Name = pb.Name
	policyField, err := FederationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.RecipientName = pb.RecipientName
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = updateMaskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateProvider struct {
	// Description about the provider.
	// Wire name: 'comment'
	Comment string ``
	// Name of the provider.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the provider.
	// Wire name: 'new_name'
	NewName string ``
	// Username of Provider owner.
	// Wire name: 'owner'
	Owner string ``
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string   ``
	ForceSendFields     []string `tf:"-"`
}

func (s *UpdateProvider) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateProvider) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateProviderToPb(st *UpdateProvider) (*sharingpb.UpdateProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.UpdateProviderPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.RecipientProfileStr = st.RecipientProfileStr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateProviderFromPb(pb *sharingpb.UpdateProviderPb) (*UpdateProvider, error) {
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

type UpdateRecipient struct {
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string ``
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList ``
	// Name of the recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the recipient. .
	// Wire name: 'new_name'
	NewName string ``
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string ``
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs ``
	ForceSendFields   []string                    `tf:"-"`
}

func (s *UpdateRecipient) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRecipient) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateRecipientToPb(st *UpdateRecipient) (*sharingpb.UpdateRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.UpdateRecipientPb{}
	pb.Comment = st.Comment
	pb.ExpirationTime = st.ExpirationTime
	ipAccessListPb, err := IpAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	propertiesKvpairsPb, err := SecurablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateRecipientFromPb(pb *sharingpb.UpdateRecipientPb) (*UpdateRecipient, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRecipient{}
	st.Comment = pb.Comment
	st.ExpirationTime = pb.ExpirationTime
	ipAccessListField, err := IpAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	propertiesKvpairsField, err := SecurablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsField != nil {
		st.PropertiesKvpairs = propertiesKvpairsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string ``
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the share.
	// Wire name: 'new_name'
	NewName string ``
	// Username of current owner of share.
	// Wire name: 'owner'
	Owner string ``
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string ``
	// Array of shared data object updates.
	// Wire name: 'updates'
	Updates         []SharedDataObjectUpdate ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *UpdateShare) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateShare) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateShareToPb(st *UpdateShare) (*sharingpb.UpdateSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.UpdateSharePb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.StorageRoot = st.StorageRoot

	var updatesPb []sharingpb.SharedDataObjectUpdatePb
	for _, item := range st.Updates {
		itemPb, err := SharedDataObjectUpdateToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			updatesPb = append(updatesPb, *itemPb)
		}
	}
	pb.Updates = updatesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateShareFromPb(pb *sharingpb.UpdateSharePb) (*UpdateShare, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateShare{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.StorageRoot = pb.StorageRoot

	var updatesField []SharedDataObjectUpdate
	for _, itemPb := range pb.Updates {
		item, err := SharedDataObjectUpdateFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			updatesField = append(updatesField, *item)
		}
	}
	st.Updates = updatesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateSharePermissions struct {
	// Array of permissions change objects.
	// Wire name: 'changes'
	Changes []PermissionsChange ``
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Optional. Whether to return the latest permissions list of the share in
	// the response.
	// Wire name: 'omit_permissions_list'
	OmitPermissionsList bool     ``
	ForceSendFields     []string `tf:"-"`
}

func (s *UpdateSharePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateSharePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateSharePermissionsToPb(st *UpdateSharePermissions) (*sharingpb.UpdateSharePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.UpdateSharePermissionsPb{}

	var changesPb []sharingpb.PermissionsChangePb
	for _, item := range st.Changes {
		itemPb, err := PermissionsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			changesPb = append(changesPb, *itemPb)
		}
	}
	pb.Changes = changesPb
	pb.Name = st.Name
	pb.OmitPermissionsList = st.OmitPermissionsList

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateSharePermissionsFromPb(pb *sharingpb.UpdateSharePermissionsPb) (*UpdateSharePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSharePermissions{}

	var changesField []PermissionsChange
	for _, itemPb := range pb.Changes {
		item, err := PermissionsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			changesField = append(changesField, *item)
		}
	}
	st.Changes = changesField
	st.Name = pb.Name
	st.OmitPermissionsList = pb.OmitPermissionsList

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment ``
}

func UpdateSharePermissionsResponseToPb(st *UpdateSharePermissionsResponse) (*sharingpb.UpdateSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.UpdateSharePermissionsResponsePb{}

	var privilegeAssignmentsPb []sharingpb.PrivilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := PrivilegeAssignmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			privilegeAssignmentsPb = append(privilegeAssignmentsPb, *itemPb)
		}
	}
	pb.PrivilegeAssignments = privilegeAssignmentsPb

	return pb, nil
}

func UpdateSharePermissionsResponseFromPb(pb *sharingpb.UpdateSharePermissionsResponsePb) (*UpdateSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSharePermissionsResponse{}

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := PrivilegeAssignmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			privilegeAssignmentsField = append(privilegeAssignmentsField, *item)
		}
	}
	st.PrivilegeAssignments = privilegeAssignmentsField

	return st, nil
}

type Volume struct {
	// The comment of the volume.
	// Wire name: 'comment'
	Comment string ``
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	// Wire name: 'id'
	Id string ``
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	// Wire name: 'internal_attributes'
	InternalAttributes *VolumeInternalAttributes ``
	// The name of the volume.
	// Wire name: 'name'
	Name string ``
	// The name of the schema that the volume belongs to.
	// Wire name: 'schema'
	Schema string ``
	// The name of the share that the volume belongs to.
	// Wire name: 'share'
	Share string ``
	// / The id of the share that the volume belongs to.
	// Wire name: 'share_id'
	ShareId string ``
	// The tags of the volume.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue ``
	ForceSendFields []string              `tf:"-"`
}

func (s *Volume) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Volume) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VolumeToPb(st *Volume) (*sharingpb.VolumePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.VolumePb{}
	pb.Comment = st.Comment
	pb.Id = st.Id
	internalAttributesPb, err := VolumeInternalAttributesToPb(st.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesPb != nil {
		pb.InternalAttributes = internalAttributesPb
	}
	pb.Name = st.Name
	pb.Schema = st.Schema
	pb.Share = st.Share
	pb.ShareId = st.ShareId

	var tagsPb []catalogpb.TagKeyValuePb
	for _, item := range st.Tags {
		itemPb, err := catalog.TagKeyValueToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VolumeFromPb(pb *sharingpb.VolumePb) (*Volume, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Volume{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	internalAttributesField, err := VolumeInternalAttributesFromPb(pb.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesField != nil {
		st.InternalAttributes = internalAttributesField
	}
	st.Name = pb.Name
	st.Schema = pb.Schema
	st.Share = pb.Share
	st.ShareId = pb.ShareId

	var tagsField []catalog.TagKeyValue
	for _, itemPb := range pb.Tags {
		item, err := catalog.TagKeyValueFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type VolumeInternalAttributes struct {
	// The cloud storage location of the volume
	// Wire name: 'storage_location'
	StorageLocation string ``
	// The type of the shared volume.
	// Wire name: 'type'
	Type            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *VolumeInternalAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VolumeInternalAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VolumeInternalAttributesToPb(st *VolumeInternalAttributes) (*sharingpb.VolumeInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.VolumeInternalAttributesPb{}
	pb.StorageLocation = st.StorageLocation
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VolumeInternalAttributesFromPb(pb *sharingpb.VolumeInternalAttributesPb) (*VolumeInternalAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumeInternalAttributes{}
	st.StorageLocation = pb.StorageLocation
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
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
