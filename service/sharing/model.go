// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
	Policy FederationPolicy `json:"policy"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being created.
	RecipientName string `json:"-" tf:"-"`
}

func (st CreateFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.CreateFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the provider.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	// Wire name: 'name'
	Name string `json:"name"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string   `json:"recipient_profile_str,omitempty"`
	ForceSendFields     []string `json:"-" tf:"-"`
}

func (st CreateProvider) MarshalJSON() ([]byte, error) {
	pb, err := CreateProviderToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateProvider) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.CreateProviderPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateProviderFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateRecipient struct {

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	// Wire name: 'name'
	Name string `json:"name"`
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'sharing_code'
	SharingCode     string   `json:"sharing_code,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateRecipient) MarshalJSON() ([]byte, error) {
	pb, err := CreateRecipientToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateRecipient) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.CreateRecipientPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateRecipientFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the share.
	// Wire name: 'name'
	Name string `json:"name"`
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot     string   `json:"storage_root,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateShare) MarshalJSON() ([]byte, error) {
	pb, err := CreateShareToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateShare) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.CreateSharePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateShareFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateShareToPb(st *CreateShare) (*sharingpb.CreateSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.CreateSharePb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.StorageRoot = st.StorageRoot

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be deleted.
	Name string `json:"-" tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being deleted.
	RecipientName string `json:"-" tf:"-"`
}

func (st DeleteFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeleteFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st DeleteProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeleteProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st DeleteRecipientRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRecipientRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRecipientRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeleteRecipientRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRecipientRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st DeleteShareRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteShareRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteShareRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeleteShareRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteShareRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Function *DeltaSharingFunctionDependency `json:"function,omitempty"`

	// Wire name: 'table'
	Table *DeltaSharingTableDependency `json:"table,omitempty"`
}

func (st DeltaSharingDependency) MarshalJSON() ([]byte, error) {
	pb, err := DeltaSharingDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeltaSharingDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeltaSharingDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeltaSharingDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Dependencies []DeltaSharingDependency `json:"dependencies,omitempty"`
}

func (st DeltaSharingDependencyList) MarshalJSON() ([]byte, error) {
	pb, err := DeltaSharingDependencyListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeltaSharingDependencyList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeltaSharingDependencyListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeltaSharingDependencyListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Aliases []RegisteredModelAlias `json:"aliases,omitempty"`
	// The comment of the function.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The data type of the function.
	// Wire name: 'data_type'
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// The dependency list of the function.
	// Wire name: 'dependency_list'
	DependencyList *DeltaSharingDependencyList `json:"dependency_list,omitempty"`
	// The full data type of the function.
	// Wire name: 'full_data_type'
	FullDataType string `json:"full_data_type,omitempty"`
	// The id of the function.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The function parameter information.
	// Wire name: 'input_params'
	InputParams *FunctionParameterInfos `json:"input_params,omitempty"`
	// The name of the function.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The properties of the function.
	// Wire name: 'properties'
	Properties string `json:"properties,omitempty"`
	// The routine definition of the function.
	// Wire name: 'routine_definition'
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// The name of the schema that the function belongs to.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// The securable kind of the function.
	// Wire name: 'securable_kind'
	SecurableKind SharedSecurableKind `json:"securable_kind,omitempty"`
	// The name of the share that the function belongs to.
	// Wire name: 'share'
	Share string `json:"share,omitempty"`
	// The id of the share that the function belongs to.
	// Wire name: 'share_id'
	ShareId string `json:"share_id,omitempty"`
	// The storage location of the function.
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// The tags of the function.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue `json:"tags,omitempty"`
	ForceSendFields []string              `json:"-" tf:"-"`
}

func (st DeltaSharingFunction) MarshalJSON() ([]byte, error) {
	pb, err := DeltaSharingFunctionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeltaSharingFunction) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeltaSharingFunctionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeltaSharingFunctionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency struct {

	// Wire name: 'function_name'
	FunctionName string `json:"function_name,omitempty"`

	// Wire name: 'schema_name'
	SchemaName      string   `json:"schema_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeltaSharingFunctionDependency) MarshalJSON() ([]byte, error) {
	pb, err := DeltaSharingFunctionDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeltaSharingFunctionDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeltaSharingFunctionDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeltaSharingFunctionDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeltaSharingFunctionDependencyToPb(st *DeltaSharingFunctionDependency) (*sharingpb.DeltaSharingFunctionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingFunctionDependencyPb{}
	pb.FunctionName = st.FunctionName
	pb.SchemaName = st.SchemaName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeltaSharingFunctionDependencyFromPb(pb *sharingpb.DeltaSharingFunctionDependencyPb) (*DeltaSharingFunctionDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingFunctionDependency{}
	st.FunctionName = pb.FunctionName
	st.SchemaName = pb.SchemaName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A Table in UC as a dependency.
type DeltaSharingTableDependency struct {

	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`

	// Wire name: 'table_name'
	TableName       string   `json:"table_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeltaSharingTableDependency) MarshalJSON() ([]byte, error) {
	pb, err := DeltaSharingTableDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeltaSharingTableDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.DeltaSharingTableDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeltaSharingTableDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeltaSharingTableDependencyToPb(st *DeltaSharingTableDependency) (*sharingpb.DeltaSharingTableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.DeltaSharingTableDependencyPb{}
	pb.SchemaName = st.SchemaName
	pb.TableName = st.TableName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeltaSharingTableDependencyFromPb(pb *sharingpb.DeltaSharingTableDependencyPb) (*DeltaSharingTableDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingTableDependency{}
	st.SchemaName = pb.SchemaName
	st.TableName = pb.TableName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FederationPolicy struct {
	// Description of the policy. This is a user-provided description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// System-generated timestamp indicating when the policy was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"` //legacy
	// Unique, immutable system-generated identifier for the federation policy.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Name of the federation policy. A recipient can have multiple policies
	// with different names. The name must contain only lowercase alphanumeric
	// characters, numbers, and hyphens.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Specifies the policy to use for validating OIDC claims in the federated
	// tokens.
	// Wire name: 'oidc_policy'
	OidcPolicy *OidcFederationPolicy `json:"oidc_policy,omitempty"`
	// System-generated timestamp indicating when the policy was last updated.
	// Wire name: 'update_time'
	UpdateTime      string   `json:"update_time,omitempty"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := FederationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.FederationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FederationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FederationPolicyToPb(st *FederationPolicy) (*sharingpb.FederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.FederationPolicyPb{}
	pb.Comment = st.Comment
	pb.CreateTime = st.CreateTime
	pb.Id = st.Id
	pb.Name = st.Name
	oidcPolicyPb, err := OidcFederationPolicyToPb(st.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyPb != nil {
		pb.OidcPolicy = oidcPolicyPb
	}
	pb.UpdateTime = st.UpdateTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FederationPolicyFromPb(pb *sharingpb.FederationPolicyPb) (*FederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FederationPolicy{}
	st.Comment = pb.Comment
	st.CreateTime = pb.CreateTime
	st.Id = pb.Id
	st.Name = pb.Name
	oidcPolicyField, err := OidcFederationPolicyFromPb(pb.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyField != nil {
		st.OidcPolicy = oidcPolicyField
	}
	st.UpdateTime = pb.UpdateTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Represents a parameter of a function. The same message is used for both input
// and output columns.
type FunctionParameterInfo struct {
	// The comment of the parameter.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the parameter.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The default value of the parameter.
	// Wire name: 'parameter_default'
	ParameterDefault string `json:"parameter_default,omitempty"`
	// The mode of the function parameter.
	// Wire name: 'parameter_mode'
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// The type of the function parameter.
	// Wire name: 'parameter_type'
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// The position of the parameter.
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// The interval type of the parameter type.
	// Wire name: 'type_interval_type'
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The type of the parameter in JSON format.
	// Wire name: 'type_json'
	TypeJson string `json:"type_json,omitempty"`
	// The type of the parameter in Enum format.
	// Wire name: 'type_name'
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// The precision of the parameter type.
	// Wire name: 'type_precision'
	TypePrecision int `json:"type_precision,omitempty"`
	// The scale of the parameter type.
	// Wire name: 'type_scale'
	TypeScale int `json:"type_scale,omitempty"`
	// The type of the parameter in text format.
	// Wire name: 'type_text'
	TypeText        string   `json:"type_text,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	pb, err := FunctionParameterInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.FunctionParameterInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FunctionParameterInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FunctionParameterInfos struct {
	// The list of parameters of the function.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func (st FunctionParameterInfos) MarshalJSON() ([]byte, error) {
	pb, err := FunctionParameterInfosToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FunctionParameterInfos) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.FunctionParameterInfosPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FunctionParameterInfosFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	ActivationUrl string `json:"-" tf:"-"`
}

func (st GetActivationUrlInfoRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetActivationUrlInfoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetActivationUrlInfoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetActivationUrlInfoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetActivationUrlInfoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being retrieved.
	RecipientName string `json:"-" tf:"-"`
}

func (st GetFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st GetProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"-" tf:"-"`
}

func (st GetRecipientRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRecipientRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRecipientRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetRecipientRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRecipientRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share permissions for a recipient.
	// Wire name: 'permissions_out'
	PermissionsOut  []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`
	ForceSendFields []string                     `json:"-" tf:"-"`
}

func (st GetRecipientSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetRecipientSharePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRecipientSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetRecipientSharePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRecipientSharePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	ForceSendFields      []string              `json:"-" tf:"-"`
}

func (st GetSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetSharePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetSharePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetSharePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData bool `json:"-" tf:"-"`
	// The name of the share.
	Name            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetShareRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetShareRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetShareRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.GetShareRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetShareRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetShareRequestToPb(st *GetShareRequest) (*sharingpb.GetShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.GetShareRequestPb{}
	pb.IncludeSharedData = st.IncludeSharedData
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetShareRequestFromPb(pb *sharingpb.GetShareRequestPb) (*GetShareRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetShareRequest{}
	st.IncludeSharedData = pb.IncludeSharedData
	st.Name = pb.Name

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	// Wire name: 'allowed_ip_addresses'
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

func (st IpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := IpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *IpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.IpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := IpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	MaxResults int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policies are being listed.
	RecipientName   string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListFederationPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListFederationPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListFederationPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListFederationPoliciesRequestToPb(st *ListFederationPoliciesRequest) (*sharingpb.ListFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListFederationPoliciesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.RecipientName = st.RecipientName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListFederationPoliciesResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'policies'
	Policies        []FederationPolicy `json:"policies,omitempty"`
	ForceSendFields []string           `json:"-" tf:"-"`
}

func (st ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListFederationPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListFederationPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListFederationPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListProviderShareAssetsRequest struct {
	// Maximum number of functions to return.
	FunctionMaxResults int `json:"-" tf:"-"`
	// Maximum number of notebooks to return.
	NotebookMaxResults int `json:"-" tf:"-"`
	// The name of the provider who owns the share.
	ProviderName string `json:"-" tf:"-"`
	// The name of the share.
	ShareName string `json:"-" tf:"-"`
	// Maximum number of tables to return.
	TableMaxResults int `json:"-" tf:"-"`
	// Maximum number of volumes to return.
	VolumeMaxResults int      `json:"-" tf:"-"`
	ForceSendFields  []string `json:"-" tf:"-"`
}

func (st ListProviderShareAssetsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListProviderShareAssetsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListProviderShareAssetsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListProviderShareAssetsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListProviderShareAssetsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Response to ListProviderShareAssets, which contains the list of assets of a
// share.
type ListProviderShareAssetsResponse struct {
	// The list of functions in the share.
	// Wire name: 'functions'
	Functions []DeltaSharingFunction `json:"functions,omitempty"`
	// The list of notebooks in the share.
	// Wire name: 'notebooks'
	Notebooks []NotebookFile `json:"notebooks,omitempty"`
	// The list of tables in the share.
	// Wire name: 'tables'
	Tables []Table `json:"tables,omitempty"`
	// The list of volumes in the share.
	// Wire name: 'volumes'
	Volumes []Volume `json:"volumes,omitempty"`
}

func (st ListProviderShareAssetsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListProviderShareAssetsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListProviderShareAssetsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListProviderShareAssetsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListProviderShareAssetsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider shares.
	// Wire name: 'shares'
	Shares          []ProviderShare `json:"shares,omitempty"`
	ForceSendFields []string        `json:"-" tf:"-"`
}

func (st ListProviderSharesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListProviderSharesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListProviderSharesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListProviderSharesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListProviderSharesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListProvidersRequest struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId string `json:"-" tf:"-"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListProvidersRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListProvidersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListProvidersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListProvidersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListProvidersRequestToPb(st *ListProvidersRequest) (*sharingpb.ListProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListProvidersRequestPb{}
	pb.DataProviderGlobalMetastoreId = st.DataProviderGlobalMetastoreId
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider information objects.
	// Wire name: 'providers'
	Providers       []ProviderInfo `json:"providers,omitempty"`
	ForceSendFields []string       `json:"-" tf:"-"`
}

func (st ListProvidersResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListProvidersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListProvidersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListProvidersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListRecipientsRequest struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId string `json:"-" tf:"-"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults int `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListRecipientsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListRecipientsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListRecipientsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListRecipientsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListRecipientsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListRecipientsRequestToPb(st *ListRecipientsRequest) (*sharingpb.ListRecipientsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListRecipientsRequestPb{}
	pb.DataRecipientGlobalMetastoreId = st.DataRecipientGlobalMetastoreId
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of recipient information objects.
	// Wire name: 'recipients'
	Recipients      []RecipientInfo `json:"recipients,omitempty"`
	ForceSendFields []string        `json:"-" tf:"-"`
}

func (st ListRecipientsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListRecipientsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListRecipientsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListRecipientsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListRecipientsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	MaxResults int `json:"-" tf:"-"`
	// Name of the provider in which to list shares.
	Name string `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListSharesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListSharesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListSharesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListSharesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListSharesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListSharesRequestToPb(st *ListSharesRequest) (*sharingpb.ListSharesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ListSharesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.Name = st.Name
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share information objects.
	// Wire name: 'shares'
	Shares          []ShareInfo `json:"shares,omitempty"`
	ForceSendFields []string    `json:"-" tf:"-"`
}

func (st ListSharesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListSharesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListSharesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ListSharesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListSharesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type NotebookFile struct {
	// The comment of the notebook file.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The id of the notebook file.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Name of the notebook file.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The name of the share that the notebook file belongs to.
	// Wire name: 'share'
	Share string `json:"share,omitempty"`
	// The id of the share that the notebook file belongs to.
	// Wire name: 'share_id'
	ShareId string `json:"share_id,omitempty"`
	// The tags of the notebook file.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue `json:"tags,omitempty"`
	ForceSendFields []string              `json:"-" tf:"-"`
}

func (st NotebookFile) MarshalJSON() ([]byte, error) {
	pb, err := NotebookFileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NotebookFile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.NotebookFilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NotebookFileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Audiences []string `json:"audiences,omitempty"`
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	// Wire name: 'issuer'
	Issuer string `json:"issuer"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. The subject claim identifies the identity of the user
	// or machine accessing the resource. Examples for Entra ID (AAD): - U2M
	// flow (group access): If the subject claim is `groups`, this must be the
	// Object ID of the group in Entra ID. - U2M flow (user access): If the
	// subject claim is `oid`, this must be the Object ID of the user in Entra
	// ID. - M2M flow (OAuth App access): If the subject claim is `azp`, this
	// must be the client ID of the OAuth app registered in Entra ID.
	// Wire name: 'subject'
	Subject string `json:"subject"`
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
	SubjectClaim string `json:"subject_claim"`
}

func (st OidcFederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := OidcFederationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *OidcFederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.OidcFederationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := OidcFederationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Values []PartitionValue `json:"values,omitempty"`
}

func (st Partition) MarshalJSON() ([]byte, error) {
	pb, err := PartitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Partition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.PartitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PartitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Name string `json:"name,omitempty"`
	// The operator to apply for the value.
	// Wire name: 'op'
	Op PartitionValueOp `json:"op,omitempty"`
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	// Wire name: 'recipient_property_key'
	RecipientPropertyKey string `json:"recipient_property_key,omitempty"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	// Wire name: 'value'
	Value           string   `json:"value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PartitionValue) MarshalJSON() ([]byte, error) {
	pb, err := PartitionValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PartitionValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.PartitionValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PartitionValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Add []string `json:"add,omitempty"`
	// The principal whose privileges we are changing. Only one of principal or
	// principal_id should be specified, never both at the same time.
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove          []string `json:"remove,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PermissionsChange) MarshalJSON() ([]byte, error) {
	pb, err := PermissionsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PermissionsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.PermissionsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermissionsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermissionsChangeToPb(st *PermissionsChange) (*sharingpb.PermissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.PermissionsChangePb{}
	pb.Add = st.Add
	pb.Principal = st.Principal
	pb.Remove = st.Remove

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges      []Privilege `json:"privileges,omitempty"`
	ForceSendFields []string    `json:"-" tf:"-"`
}

func (st PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	pb, err := PrivilegeAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.PrivilegeAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PrivilegeAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ProviderInfo struct {

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`
	// Description about the provider.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Provider creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_provider_global_metastore_id'
	DataProviderGlobalMetastoreId string `json:"data_provider_global_metastore_id,omitempty"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the Provider.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of Provider owner.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	// Wire name: 'recipient_profile'
	RecipientProfile *RecipientProfile `json:"recipient_profile,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified Provider.
	// Wire name: 'updated_by'
	UpdatedBy       string   `json:"updated_by,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ProviderInfo) MarshalJSON() ([]byte, error) {
	pb, err := ProviderInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ProviderInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ProviderInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ProviderInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ProviderShare struct {
	// The name of the Provider Share.
	// Wire name: 'name'
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ProviderShare) MarshalJSON() ([]byte, error) {
	pb, err := ProviderShareToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ProviderShare) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ProviderSharePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ProviderShareFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ProviderShareToPb(st *ProviderShare) (*sharingpb.ProviderSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.ProviderSharePb{}
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ProviderShareFromPb(pb *sharingpb.ProviderSharePb) (*ProviderShare, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProviderShare{}
	st.Name = pb.Name

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	// Wire name: 'activated'
	Activated bool `json:"activated,omitempty"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string `json:"activation_url,omitempty"`

	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Time at which this recipient was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'metastore_id'
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Recipient.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'sharing_code'
	SharingCode string `json:"sharing_code,omitempty"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	// Wire name: 'tokens'
	Tokens []RecipientTokenInfo `json:"tokens,omitempty"`
	// Time at which the recipient was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   `json:"updated_by,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RecipientInfo) MarshalJSON() ([]byte, error) {
	pb, err := RecipientInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RecipientInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RecipientInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RecipientInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	// Wire name: 'bearer_token'
	BearerToken string `json:"bearer_token,omitempty"`
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string `json:"endpoint,omitempty"`
	// The version number of the recipient's credentials on a share.
	// Wire name: 'share_credentials_version'
	ShareCredentialsVersion int      `json:"share_credentials_version,omitempty"`
	ForceSendFields         []string `json:"-" tf:"-"`
}

func (st RecipientProfile) MarshalJSON() ([]byte, error) {
	pb, err := RecipientProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RecipientProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RecipientProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RecipientProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RecipientProfileToPb(st *RecipientProfile) (*sharingpb.RecipientProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RecipientProfilePb{}
	pb.BearerToken = st.BearerToken
	pb.Endpoint = st.Endpoint
	pb.ShareCredentialsVersion = st.ShareCredentialsVersion

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this recipient token was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient token creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique ID of the recipient token.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Time at which this recipient token was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient token updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   `json:"updated_by,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RecipientTokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := RecipientTokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RecipientTokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RecipientTokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RecipientTokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegisteredModelAlias struct {
	// Name of the alias.
	// Wire name: 'alias_name'
	AliasName string `json:"alias_name,omitempty"`
	// Numeric model version that alias will reference.
	// Wire name: 'version_num'
	VersionNum      int64    `json:"version_num,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelAliasToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RegisteredModelAliasPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelAliasFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelAliasToPb(st *RegisteredModelAlias) (*sharingpb.RegisteredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.RegisteredModelAliasPb{}
	pb.AliasName = st.AliasName
	pb.VersionNum = st.VersionNum

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegisteredModelAliasFromPb(pb *sharingpb.RegisteredModelAliasPb) (*RegisteredModelAlias, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAlias{}
	st.AliasName = pb.AliasName
	st.VersionNum = pb.VersionNum

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" tf:"-"`
}

func (st RetrieveTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := RetrieveTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RetrieveTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RetrieveTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RetrieveTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	BearerToken string `json:"bearerToken,omitempty"`
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string `json:"endpoint,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expirationTime'
	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol.
	// Wire name: 'shareCredentialsVersion'
	ShareCredentialsVersion int      `json:"shareCredentialsVersion,omitempty"`
	ForceSendFields         []string `json:"-" tf:"-"`
}

func (st RetrieveTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := RetrieveTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RetrieveTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RetrieveTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RetrieveTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	// Wire name: 'existing_token_expire_in_seconds'
	ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds"`
	// The name of the Recipient.
	Name string `json:"-" tf:"-"`
}

func (st RotateRecipientToken) MarshalJSON() ([]byte, error) {
	pb, err := RotateRecipientTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RotateRecipientToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.RotateRecipientTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RotateRecipientTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Properties map[string]string `json:"properties"`
}

func (st SecurablePropertiesKvPairs) MarshalJSON() ([]byte, error) {
	pb, err := SecurablePropertiesKvPairsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SecurablePropertiesKvPairs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.SecurablePropertiesKvPairsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SecurablePropertiesKvPairsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Comment string `json:"comment,omitempty"`
	// Time at which this share was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of share creator.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Name of the share.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// A list of shared data objects within the share.
	// Wire name: 'objects'
	Objects []SharedDataObject `json:"objects,omitempty"`
	// Username of current owner of share.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Storage Location URL (full path) for the share.
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this share was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of share updater.
	// Wire name: 'updated_by'
	UpdatedBy       string   `json:"updated_by,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ShareInfo) MarshalJSON() ([]byte, error) {
	pb, err := ShareInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ShareInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ShareInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ShareInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	MaxResults int `json:"-" tf:"-"`
	// The name of the Recipient.
	Name string `json:"-" tf:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SharePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := SharePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SharePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.SharePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SharePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SharePermissionsRequestToPb(st *SharePermissionsRequest) (*sharingpb.SharePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.SharePermissionsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.Name = st.Name
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	// The share name.
	// Wire name: 'share_name'
	ShareName       string   `json:"share_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ShareToPrivilegeAssignment) MarshalJSON() ([]byte, error) {
	pb, err := ShareToPrivilegeAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ShareToPrivilegeAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.ShareToPrivilegeAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ShareToPrivilegeAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	// Wire name: 'added_at'
	AddedAt int64 `json:"added_at,omitempty"`
	// Username of the sharer.
	// Wire name: 'added_by'
	AddedBy string `json:"added_by,omitempty"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	// Wire name: 'cdf_enabled'
	CdfEnabled bool `json:"cdf_enabled,omitempty"`
	// A user-provided comment when adding the data object to the share.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	// Wire name: 'content'
	Content string `json:"content,omitempty"`
	// The type of the data object.
	// Wire name: 'data_object_type'
	DataObjectType SharedDataObjectDataObjectType `json:"data_object_type,omitempty"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	// Wire name: 'history_data_sharing_status'
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus `json:"history_data_sharing_status,omitempty"`
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`,
	// Wire name: 'name'
	Name string `json:"name"`
	// Array of partitions for the shared data.
	// Wire name: 'partitions'
	Partitions []Partition `json:"partitions,omitempty"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	// Wire name: 'shared_as'
	SharedAs string `json:"shared_as,omitempty"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	// Wire name: 'start_version'
	StartVersion int64 `json:"start_version,omitempty"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	// Wire name: 'status'
	Status SharedDataObjectStatus `json:"status,omitempty"`
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
	// Wire name: 'string_shared_as'
	StringSharedAs  string   `json:"string_shared_as,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SharedDataObject) MarshalJSON() ([]byte, error) {
	pb, err := SharedDataObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SharedDataObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.SharedDataObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SharedDataObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Action SharedDataObjectUpdateAction `json:"action,omitempty"`
	// The data object that is being added, removed, or updated. The maximum
	// number update data objects allowed is a 100.
	// Wire name: 'data_object'
	DataObject *SharedDataObject `json:"data_object,omitempty"`
}

func (st SharedDataObjectUpdate) MarshalJSON() ([]byte, error) {
	pb, err := SharedDataObjectUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SharedDataObjectUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.SharedDataObjectUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SharedDataObjectUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Comment string `json:"comment,omitempty"`
	// The id of the table.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	// Wire name: 'internal_attributes'
	InternalAttributes *TableInternalAttributes `json:"internal_attributes,omitempty"`
	// The catalog and schema of the materialized table
	// Wire name: 'materialization_namespace'
	MaterializationNamespace string `json:"materialization_namespace,omitempty"`
	// The name of a materialized table.
	// Wire name: 'materialized_table_name'
	MaterializedTableName string `json:"materialized_table_name,omitempty"`
	// The name of the table.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The name of the schema that the table belongs to.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// The name of the share that the table belongs to.
	// Wire name: 'share'
	Share string `json:"share,omitempty"`
	// The id of the share that the table belongs to.
	// Wire name: 'share_id'
	ShareId string `json:"share_id,omitempty"`
	// The Tags of the table.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue `json:"tags,omitempty"`
	ForceSendFields []string              `json:"-" tf:"-"`
}

func (st Table) MarshalJSON() ([]byte, error) {
	pb, err := TableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Table) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.TablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	ParentStorageLocation string `json:"parent_storage_location,omitempty"`
	// The cloud storage location of a shard table with DIRECTORY_BASED_TABLE
	// type.
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the shared table.
	// Wire name: 'type'
	Type TableInternalAttributesSharedTableType `json:"type,omitempty"`
	// The view definition of a shared view. DEPRECATED.
	// Wire name: 'view_definition'
	ViewDefinition  string   `json:"view_definition,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TableInternalAttributes) MarshalJSON() ([]byte, error) {
	pb, err := TableInternalAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TableInternalAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.TableInternalAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TableInternalAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
	Name string `json:"-" tf:"-"`

	// Wire name: 'policy'
	Policy FederationPolicy `json:"policy"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being updated.
	RecipientName string `json:"-" tf:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'comment,oidc_policy.audiences'.
	UpdateMask      string   `json:"-" tf:"-"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.UpdateFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	pb.UpdateMask = st.UpdateMask

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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
	st.UpdateMask = pb.UpdateMask

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateProvider struct {
	// Description about the provider.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the provider.
	Name string `json:"-" tf:"-"`
	// New name for the provider.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// Username of Provider owner.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string   `json:"recipient_profile_str,omitempty"`
	ForceSendFields     []string `json:"-" tf:"-"`
}

func (st UpdateProvider) MarshalJSON() ([]byte, error) {
	pb, err := UpdateProviderToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateProvider) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.UpdateProviderPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateProviderFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateRecipient struct {
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList `json:"ip_access_list,omitempty"`
	// Name of the recipient.
	Name string `json:"-" tf:"-"`
	// New name for the recipient. .
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs `json:"properties_kvpairs,omitempty"`
	ForceSendFields   []string                    `json:"-" tf:"-"`
}

func (st UpdateRecipient) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRecipientToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRecipient) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.UpdateRecipientPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRecipientFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The name of the share.
	Name string `json:"-" tf:"-"`
	// New name for the share.
	// Wire name: 'new_name'
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of share.
	// Wire name: 'owner'
	Owner string `json:"owner,omitempty"`
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`
	// Array of shared data object updates.
	// Wire name: 'updates'
	Updates         []SharedDataObjectUpdate `json:"updates,omitempty"`
	ForceSendFields []string                 `json:"-" tf:"-"`
}

func (st UpdateShare) MarshalJSON() ([]byte, error) {
	pb, err := UpdateShareToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateShare) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.UpdateSharePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateShareFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateSharePermissions struct {
	// Array of permissions change objects.
	// Wire name: 'changes'
	Changes []PermissionsChange `json:"changes,omitempty"`
	// The name of the share.
	Name string `json:"-" tf:"-"`
	// Optional. Whether to return the latest permissions list of the share in
	// the response.
	// Wire name: 'omit_permissions_list'
	OmitPermissionsList bool     `json:"omit_permissions_list,omitempty"`
	ForceSendFields     []string `json:"-" tf:"-"`
}

func (st UpdateSharePermissions) MarshalJSON() ([]byte, error) {
	pb, err := UpdateSharePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateSharePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.UpdateSharePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateSharePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func (st UpdateSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateSharePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.UpdateSharePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateSharePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
	Comment string `json:"comment,omitempty"`
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	// Wire name: 'internal_attributes'
	InternalAttributes *VolumeInternalAttributes `json:"internal_attributes,omitempty"`
	// The name of the volume.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The name of the schema that the volume belongs to.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// The name of the share that the volume belongs to.
	// Wire name: 'share'
	Share string `json:"share,omitempty"`
	// / The id of the share that the volume belongs to.
	// Wire name: 'share_id'
	ShareId string `json:"share_id,omitempty"`
	// The tags of the volume.
	// Wire name: 'tags'
	Tags            []catalog.TagKeyValue `json:"tags,omitempty"`
	ForceSendFields []string              `json:"-" tf:"-"`
}

func (st Volume) MarshalJSON() ([]byte, error) {
	pb, err := VolumeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Volume) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.VolumePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := VolumeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type VolumeInternalAttributes struct {
	// The cloud storage location of the volume
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the shared volume.
	// Wire name: 'type'
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st VolumeInternalAttributes) MarshalJSON() ([]byte, error) {
	pb, err := VolumeInternalAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *VolumeInternalAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharingpb.VolumeInternalAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := VolumeInternalAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func VolumeInternalAttributesToPb(st *VolumeInternalAttributes) (*sharingpb.VolumeInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharingpb.VolumeInternalAttributesPb{}
	pb.StorageLocation = st.StorageLocation
	pb.Type = st.Type

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func VolumeInternalAttributesFromPb(pb *sharingpb.VolumeInternalAttributesPb) (*VolumeInternalAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumeInternalAttributes{}
	st.StorageLocation = pb.StorageLocation
	st.Type = pb.Type

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
