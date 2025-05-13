// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog"
)

// The delta sharing authentication type.
type AuthenticationType string
type authenticationTypePb string

const AuthenticationTypeDatabricks AuthenticationType = `DATABRICKS`

const AuthenticationTypeOauthClientCredentials AuthenticationType = `OAUTH_CLIENT_CREDENTIALS`

const AuthenticationTypeToken AuthenticationType = `TOKEN`

// String representation for [fmt.Print]
func (f *AuthenticationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationType) Set(v string) error {
	switch v {
	case `DATABRICKS`, `OAUTH_CLIENT_CREDENTIALS`, `TOKEN`:
		*f = AuthenticationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATABRICKS", "OAUTH_CLIENT_CREDENTIALS", "TOKEN"`, v)
	}
}

// Type always returns AuthenticationType to satisfy [pflag.Value] interface
func (f *AuthenticationType) Type() string {
	return "AuthenticationType"
}

func authenticationTypeToPb(st *AuthenticationType) (*authenticationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := authenticationTypePb(*st)
	return &pb, nil
}

func authenticationTypeFromPb(pb *authenticationTypePb) (*AuthenticationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AuthenticationType(*pb)
	return &st, nil
}

// UC supported column types Copied from
// https://src.dev.databricks.com/databricks/universe@23a85902bb58695ab9293adc9f327b0714b55e72/-/blob/managed-catalog/api/messages/table.proto?L68
type ColumnTypeName string
type columnTypeNamePb string

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

// Type always returns ColumnTypeName to satisfy [pflag.Value] interface
func (f *ColumnTypeName) Type() string {
	return "ColumnTypeName"
}

func columnTypeNameToPb(st *ColumnTypeName) (*columnTypeNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := columnTypeNamePb(*st)
	return &pb, nil
}

func columnTypeNameFromPb(pb *columnTypeNamePb) (*ColumnTypeName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ColumnTypeName(*pb)
	return &st, nil
}

type CreateProvider struct {
	// The delta sharing authentication type.
	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType
	// Description about the provider.
	// Wire name: 'comment'
	Comment string
	// The name of the Provider.
	// Wire name: 'name'
	Name string
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateProvider) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createProviderPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createProviderFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateProvider) MarshalJSON() ([]byte, error) {
	pb, err := createProviderToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createProviderPb struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// The name of the Provider.
	Name string `json:"name"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

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

type CreateRecipient struct {
	// The delta sharing authentication type.
	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList
	// Name of Recipient.
	// Wire name: 'name'
	Name string
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'sharing_code'
	SharingCode string

	ForceSendFields []string `tf:"-"`
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

	ipAccessListPb, err := ipAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	pb.Name = st.Name

	pb.Owner = st.Owner

	propertiesKvpairsPb, err := securablePropertiesKvPairsToPb(st.PropertiesKvpairs)
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

func (st *CreateRecipient) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRecipientPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRecipientFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRecipient) MarshalJSON() ([]byte, error) {
	pb, err := createRecipientToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createRecipientPb struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// IP Access List
	IpAccessList *ipAccessListPb `json:"ip_access_list,omitempty"`
	// Name of Recipient.
	Name string `json:"name"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *securablePropertiesKvPairsPb `json:"properties_kvpairs,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `json:"sharing_code,omitempty"`

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
	ipAccessListField, err := ipAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	st.Name = pb.Name
	st.Owner = pb.Owner
	propertiesKvpairsField, err := securablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
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

func (st *createRecipientPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRecipientPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Name of the share.
	// Wire name: 'name'
	Name string
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateShare) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createSharePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createShareFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateShare) MarshalJSON() ([]byte, error) {
	pb, err := createShareToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createSharePb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Name of the share.
	Name string `json:"name"`
	// Storage root URL for the share.
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

// Delete a provider
type DeleteProviderRequest struct {
	// Name of the provider.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteProviderRequestToPb(st *DeleteProviderRequest) (*deleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteProviderRequestPb struct {
	// Name of the provider.
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

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteRecipientRequestToPb(st *DeleteRecipientRequest) (*deleteRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRecipientRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteRecipientRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRecipientRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRecipientRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRecipientRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteRecipientRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRecipientRequestPb struct {
	// Name of the recipient.
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

// Delete a share
type DeleteShareRequest struct {
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteShareRequestToPb(st *DeleteShareRequest) (*deleteShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteShareRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteShareRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteShareRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteShareRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteShareRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteShareRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteShareRequestPb struct {
	// The name of the share.
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

// Represents a UC dependency.
type DeltaSharingDependency struct {
	// A Function in UC as a dependency.
	// Wire name: 'function'
	Function *DeltaSharingFunctionDependency
	// A Table in UC as a dependency.
	// Wire name: 'table'
	Table *DeltaSharingTableDependency
}

func deltaSharingDependencyToPb(st *DeltaSharingDependency) (*deltaSharingDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingDependencyPb{}
	functionPb, err := deltaSharingFunctionDependencyToPb(st.Function)
	if err != nil {
		return nil, err
	}
	if functionPb != nil {
		pb.Function = functionPb
	}

	tablePb, err := deltaSharingTableDependencyToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
}

func (st *DeltaSharingDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSharingDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSharingDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSharingDependency) MarshalJSON() ([]byte, error) {
	pb, err := deltaSharingDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSharingDependencyPb struct {
	// A Function in UC as a dependency.
	Function *deltaSharingFunctionDependencyPb `json:"function,omitempty"`
	// A Table in UC as a dependency.
	Table *deltaSharingTableDependencyPb `json:"table,omitempty"`
}

func deltaSharingDependencyFromPb(pb *deltaSharingDependencyPb) (*DeltaSharingDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingDependency{}
	functionField, err := deltaSharingFunctionDependencyFromPb(pb.Function)
	if err != nil {
		return nil, err
	}
	if functionField != nil {
		st.Function = functionField
	}
	tableField, err := deltaSharingTableDependencyFromPb(pb.Table)
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
	Dependencies []DeltaSharingDependency
}

func deltaSharingDependencyListToPb(st *DeltaSharingDependencyList) (*deltaSharingDependencyListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingDependencyListPb{}

	var dependenciesPb []deltaSharingDependencyPb
	for _, item := range st.Dependencies {
		itemPb, err := deltaSharingDependencyToPb(&item)
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

func (st *DeltaSharingDependencyList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSharingDependencyListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSharingDependencyListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSharingDependencyList) MarshalJSON() ([]byte, error) {
	pb, err := deltaSharingDependencyListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSharingDependencyListPb struct {
	// An array of Dependency.
	Dependencies []deltaSharingDependencyPb `json:"dependencies,omitempty"`
}

func deltaSharingDependencyListFromPb(pb *deltaSharingDependencyListPb) (*DeltaSharingDependencyList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingDependencyList{}

	var dependenciesField []DeltaSharingDependency
	for _, itemPb := range pb.Dependencies {
		item, err := deltaSharingDependencyFromPb(&itemPb)
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
	Aliases []RegisteredModelAlias
	// The comment of the function.
	// Wire name: 'comment'
	Comment string
	// The data type of the function.
	// Wire name: 'data_type'
	DataType ColumnTypeName
	// The dependency list of the function.
	// Wire name: 'dependency_list'
	DependencyList *DeltaSharingDependencyList
	// The full data type of the function.
	// Wire name: 'full_data_type'
	FullDataType string
	// The id of the function.
	// Wire name: 'id'
	Id string
	// The function parameter information.
	// Wire name: 'input_params'
	InputParams *FunctionParameterInfos
	// The name of the function.
	// Wire name: 'name'
	Name string
	// The properties of the function.
	// Wire name: 'properties'
	Properties string
	// The routine definition of the function.
	// Wire name: 'routine_definition'
	RoutineDefinition string
	// The name of the schema that the function belongs to.
	// Wire name: 'schema'
	Schema string
	// The securable kind of the function.
	// Wire name: 'securable_kind'
	SecurableKind SharedSecurableKind
	// The name of the share that the function belongs to.
	// Wire name: 'share'
	Share string
	// The id of the share that the function belongs to.
	// Wire name: 'share_id'
	ShareId string
	// The storage location of the function.
	// Wire name: 'storage_location'
	StorageLocation string
	// The tags of the function.
	// Wire name: 'tags'
	Tags []catalog.TagKeyValue

	ForceSendFields []string `tf:"-"`
}

func deltaSharingFunctionToPb(st *DeltaSharingFunction) (*deltaSharingFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingFunctionPb{}

	var aliasesPb []registeredModelAliasPb
	for _, item := range st.Aliases {
		itemPb, err := registeredModelAliasToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			aliasesPb = append(aliasesPb, *itemPb)
		}
	}
	pb.Aliases = aliasesPb

	pb.Comment = st.Comment

	pb.DataType = st.DataType

	dependencyListPb, err := deltaSharingDependencyListToPb(st.DependencyList)
	if err != nil {
		return nil, err
	}
	if dependencyListPb != nil {
		pb.DependencyList = dependencyListPb
	}

	pb.FullDataType = st.FullDataType

	pb.Id = st.Id

	inputParamsPb, err := functionParameterInfosToPb(st.InputParams)
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

	pb.SecurableKind = st.SecurableKind

	pb.Share = st.Share

	pb.ShareId = st.ShareId

	pb.StorageLocation = st.StorageLocation

	var tagsPb []catalog.TagKeyValuePb
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

func (st *DeltaSharingFunction) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSharingFunctionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSharingFunctionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSharingFunction) MarshalJSON() ([]byte, error) {
	pb, err := deltaSharingFunctionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSharingFunctionPb struct {
	// The aliass of registered model.
	Aliases []registeredModelAliasPb `json:"aliases,omitempty"`
	// The comment of the function.
	Comment string `json:"comment,omitempty"`
	// The data type of the function.
	DataType ColumnTypeName `json:"data_type,omitempty"`
	// The dependency list of the function.
	DependencyList *deltaSharingDependencyListPb `json:"dependency_list,omitempty"`
	// The full data type of the function.
	FullDataType string `json:"full_data_type,omitempty"`
	// The id of the function.
	Id string `json:"id,omitempty"`
	// The function parameter information.
	InputParams *functionParameterInfosPb `json:"input_params,omitempty"`
	// The name of the function.
	Name string `json:"name,omitempty"`
	// The properties of the function.
	Properties string `json:"properties,omitempty"`
	// The routine definition of the function.
	RoutineDefinition string `json:"routine_definition,omitempty"`
	// The name of the schema that the function belongs to.
	Schema string `json:"schema,omitempty"`
	// The securable kind of the function.
	SecurableKind SharedSecurableKind `json:"securable_kind,omitempty"`
	// The name of the share that the function belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the function belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The storage location of the function.
	StorageLocation string `json:"storage_location,omitempty"`
	// The tags of the function.
	Tags []catalog.TagKeyValuePb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deltaSharingFunctionFromPb(pb *deltaSharingFunctionPb) (*DeltaSharingFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaSharingFunction{}

	var aliasesField []RegisteredModelAlias
	for _, itemPb := range pb.Aliases {
		item, err := registeredModelAliasFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			aliasesField = append(aliasesField, *item)
		}
	}
	st.Aliases = aliasesField
	st.Comment = pb.Comment
	st.DataType = pb.DataType
	dependencyListField, err := deltaSharingDependencyListFromPb(pb.DependencyList)
	if err != nil {
		return nil, err
	}
	if dependencyListField != nil {
		st.DependencyList = dependencyListField
	}
	st.FullDataType = pb.FullDataType
	st.Id = pb.Id
	inputParamsField, err := functionParameterInfosFromPb(pb.InputParams)
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
	st.SecurableKind = pb.SecurableKind
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

func (st *deltaSharingFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deltaSharingFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency struct {

	// Wire name: 'function_name'
	FunctionName string

	// Wire name: 'schema_name'
	SchemaName string

	ForceSendFields []string `tf:"-"`
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

func (st *DeltaSharingFunctionDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSharingFunctionDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSharingFunctionDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSharingFunctionDependency) MarshalJSON() ([]byte, error) {
	pb, err := deltaSharingFunctionDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSharingFunctionDependencyPb struct {
	FunctionName string `json:"function_name,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

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

// A Table in UC as a dependency.
type DeltaSharingTableDependency struct {

	// Wire name: 'schema_name'
	SchemaName string

	// Wire name: 'table_name'
	TableName string

	ForceSendFields []string `tf:"-"`
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

func (st *DeltaSharingTableDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deltaSharingTableDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deltaSharingTableDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeltaSharingTableDependency) MarshalJSON() ([]byte, error) {
	pb, err := deltaSharingTableDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deltaSharingTableDependencyPb struct {
	SchemaName string `json:"schema_name,omitempty"`

	TableName string `json:"table_name,omitempty"`

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

// Represents a parameter of a function. The same message is used for both input
// and output columns.
type FunctionParameterInfo struct {
	// The comment of the parameter.
	// Wire name: 'comment'
	Comment string
	// The name of the parameter.
	// Wire name: 'name'
	Name string
	// The default value of the parameter.
	// Wire name: 'parameter_default'
	ParameterDefault string
	// The mode of the function parameter.
	// Wire name: 'parameter_mode'
	ParameterMode FunctionParameterMode
	// The type of the function parameter.
	// Wire name: 'parameter_type'
	ParameterType FunctionParameterType
	// The position of the parameter.
	// Wire name: 'position'
	Position int
	// The interval type of the parameter type.
	// Wire name: 'type_interval_type'
	TypeIntervalType string
	// The type of the parameter in JSON format.
	// Wire name: 'type_json'
	TypeJson string
	// The type of the parameter in Enum format.
	// Wire name: 'type_name'
	TypeName ColumnTypeName
	// The precision of the parameter type.
	// Wire name: 'type_precision'
	TypePrecision int
	// The scale of the parameter type.
	// Wire name: 'type_scale'
	TypeScale int
	// The type of the parameter in text format.
	// Wire name: 'type_text'
	TypeText string

	ForceSendFields []string `tf:"-"`
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

func (st *FunctionParameterInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &functionParameterInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := functionParameterInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FunctionParameterInfo) MarshalJSON() ([]byte, error) {
	pb, err := functionParameterInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type functionParameterInfoPb struct {
	// The comment of the parameter.
	Comment string `json:"comment,omitempty"`
	// The name of the parameter.
	Name string `json:"name,omitempty"`
	// The default value of the parameter.
	ParameterDefault string `json:"parameter_default,omitempty"`
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode `json:"parameter_mode,omitempty"`
	// The type of the function parameter.
	ParameterType FunctionParameterType `json:"parameter_type,omitempty"`
	// The position of the parameter.
	Position int `json:"position,omitempty"`
	// The interval type of the parameter type.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The type of the parameter in JSON format.
	TypeJson string `json:"type_json,omitempty"`
	// The type of the parameter in Enum format.
	TypeName ColumnTypeName `json:"type_name,omitempty"`
	// The precision of the parameter type.
	TypePrecision int `json:"type_precision,omitempty"`
	// The scale of the parameter type.
	TypeScale int `json:"type_scale,omitempty"`
	// The type of the parameter in text format.
	TypeText string `json:"type_text,omitempty"`

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

type FunctionParameterInfos struct {
	// The list of parameters of the function.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo
}

func functionParameterInfosToPb(st *FunctionParameterInfos) (*functionParameterInfosPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfosPb{}

	var parametersPb []functionParameterInfoPb
	for _, item := range st.Parameters {
		itemPb, err := functionParameterInfoToPb(&item)
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

func (st *FunctionParameterInfos) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &functionParameterInfosPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := functionParameterInfosFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FunctionParameterInfos) MarshalJSON() ([]byte, error) {
	pb, err := functionParameterInfosToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type functionParameterInfosPb struct {
	// The list of parameters of the function.
	Parameters []functionParameterInfoPb `json:"parameters,omitempty"`
}

func functionParameterInfosFromPb(pb *functionParameterInfosPb) (*FunctionParameterInfos, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfos{}

	var parametersField []FunctionParameterInfo
	for _, itemPb := range pb.Parameters {
		item, err := functionParameterInfoFromPb(&itemPb)
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
type functionParameterModePb string

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

// Type always returns FunctionParameterMode to satisfy [pflag.Value] interface
func (f *FunctionParameterMode) Type() string {
	return "FunctionParameterMode"
}

func functionParameterModeToPb(st *FunctionParameterMode) (*functionParameterModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionParameterModePb(*st)
	return &pb, nil
}

func functionParameterModeFromPb(pb *functionParameterModePb) (*FunctionParameterMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterMode(*pb)
	return &st, nil
}

type FunctionParameterType string
type functionParameterTypePb string

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

// Type always returns FunctionParameterType to satisfy [pflag.Value] interface
func (f *FunctionParameterType) Type() string {
	return "FunctionParameterType"
}

func functionParameterTypeToPb(st *FunctionParameterType) (*functionParameterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := functionParameterTypePb(*st)
	return &pb, nil
}

func functionParameterTypeFromPb(pb *functionParameterTypePb) (*FunctionParameterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := FunctionParameterType(*pb)
	return &st, nil
}

// Get a share activation URL
type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	// Wire name: 'activation_url'
	ActivationUrl string `tf:"-"`
}

func getActivationUrlInfoRequestToPb(st *GetActivationUrlInfoRequest) (*getActivationUrlInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getActivationUrlInfoRequestPb{}
	pb.ActivationUrl = st.ActivationUrl

	return pb, nil
}

func (st *GetActivationUrlInfoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getActivationUrlInfoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getActivationUrlInfoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetActivationUrlInfoRequest) MarshalJSON() ([]byte, error) {
	pb, err := getActivationUrlInfoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getActivationUrlInfoRequestPb struct {
	// The one time activation url. It also accepts activation token.
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

type GetActivationUrlInfoResponse struct {
}

func getActivationUrlInfoResponseToPb(st *GetActivationUrlInfoResponse) (*getActivationUrlInfoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getActivationUrlInfoResponsePb{}

	return pb, nil
}

func (st *GetActivationUrlInfoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getActivationUrlInfoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getActivationUrlInfoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetActivationUrlInfoResponse) MarshalJSON() ([]byte, error) {
	pb, err := getActivationUrlInfoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getActivationUrlInfoResponsePb struct {
}

func getActivationUrlInfoResponseFromPb(pb *getActivationUrlInfoResponsePb) (*GetActivationUrlInfoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetActivationUrlInfoResponse{}

	return st, nil
}

// Get a provider
type GetProviderRequest struct {
	// Name of the provider.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getProviderRequestToPb(st *GetProviderRequest) (*getProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetProviderRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getProviderRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getProviderRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetProviderRequest) MarshalJSON() ([]byte, error) {
	pb, err := getProviderRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getProviderRequestPb struct {
	// Name of the provider.
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

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getRecipientRequestToPb(st *GetRecipientRequest) (*getRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRecipientRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetRecipientRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRecipientRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRecipientRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRecipientRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRecipientRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRecipientRequestPb struct {
	// Name of the recipient.
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

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of data share permissions for a recipient.
	// Wire name: 'permissions_out'
	PermissionsOut []ShareToPrivilegeAssignment

	ForceSendFields []string `tf:"-"`
}

func getRecipientSharePermissionsResponseToPb(st *GetRecipientSharePermissionsResponse) (*getRecipientSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRecipientSharePermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var permissionsOutPb []shareToPrivilegeAssignmentPb
	for _, item := range st.PermissionsOut {
		itemPb, err := shareToPrivilegeAssignmentToPb(&item)
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

func (st *GetRecipientSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRecipientSharePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRecipientSharePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRecipientSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRecipientSharePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRecipientSharePermissionsResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share permissions for a recipient.
	PermissionsOut []shareToPrivilegeAssignmentPb `json:"permissions_out,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRecipientSharePermissionsResponseFromPb(pb *getRecipientSharePermissionsResponsePb) (*GetRecipientSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRecipientSharePermissionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var permissionsOutField []ShareToPrivilegeAssignment
	for _, itemPb := range pb.PermissionsOut {
		item, err := shareToPrivilegeAssignmentFromPb(&itemPb)
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

func (st *getRecipientSharePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRecipientSharePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment

	ForceSendFields []string `tf:"-"`
}

func getSharePermissionsResponseToPb(st *GetSharePermissionsResponse) (*getSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSharePermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var privilegeAssignmentsPb []privilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := privilegeAssignmentToPb(&item)
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

func (st *GetSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSharePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSharePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getSharePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getSharePermissionsResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges assigned to each principal
	PrivilegeAssignments []privilegeAssignmentPb `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getSharePermissionsResponseFromPb(pb *getSharePermissionsResponsePb) (*GetSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSharePermissionsResponse{}
	st.NextPageToken = pb.NextPageToken

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := privilegeAssignmentFromPb(&itemPb)
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

func (st *getSharePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getSharePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a share
type GetShareRequest struct {
	// Query for data to include in the share.
	// Wire name: 'include_shared_data'
	IncludeSharedData bool `tf:"-"`
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *GetShareRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getShareRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getShareRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetShareRequest) MarshalJSON() ([]byte, error) {
	pb, err := getShareRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getShareRequestPb struct {
	// Query for data to include in the share.
	IncludeSharedData bool `json:"-" url:"include_shared_data,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`

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

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	// Wire name: 'allowed_ip_addresses'
	AllowedIpAddresses []string
}

func ipAccessListToPb(st *IpAccessList) (*ipAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ipAccessListPb{}
	pb.AllowedIpAddresses = st.AllowedIpAddresses

	return pb, nil
}

func (st *IpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ipAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ipAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st IpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := ipAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ipAccessListPb struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
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

// List assets by provider share
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
	VolumeMaxResults int `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListProviderShareAssetsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProviderShareAssetsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProviderShareAssetsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProviderShareAssetsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listProviderShareAssetsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProviderShareAssetsRequestPb struct {
	// Maximum number of functions to return.
	FunctionMaxResults int `json:"-" url:"function_max_results,omitempty"`
	// Maximum number of notebooks to return.
	NotebookMaxResults int `json:"-" url:"notebook_max_results,omitempty"`
	// The name of the provider who owns the share.
	ProviderName string `json:"-" url:"-"`
	// The name of the share.
	ShareName string `json:"-" url:"-"`
	// Maximum number of tables to return.
	TableMaxResults int `json:"-" url:"table_max_results,omitempty"`
	// Maximum number of volumes to return.
	VolumeMaxResults int `json:"-" url:"volume_max_results,omitempty"`

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

// Response to ListProviderShareAssets, which contains the list of assets of a
// share.
type ListProviderShareAssetsResponse struct {
	// The list of functions in the share.
	// Wire name: 'functions'
	Functions []DeltaSharingFunction
	// The list of notebooks in the share.
	// Wire name: 'notebooks'
	Notebooks []NotebookFile
	// The list of tables in the share.
	// Wire name: 'tables'
	Tables []Table
	// The list of volumes in the share.
	// Wire name: 'volumes'
	Volumes []Volume
}

func listProviderShareAssetsResponseToPb(st *ListProviderShareAssetsResponse) (*listProviderShareAssetsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderShareAssetsResponsePb{}

	var functionsPb []deltaSharingFunctionPb
	for _, item := range st.Functions {
		itemPb, err := deltaSharingFunctionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			functionsPb = append(functionsPb, *itemPb)
		}
	}
	pb.Functions = functionsPb

	var notebooksPb []notebookFilePb
	for _, item := range st.Notebooks {
		itemPb, err := notebookFileToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebooksPb = append(notebooksPb, *itemPb)
		}
	}
	pb.Notebooks = notebooksPb

	var tablesPb []tablePb
	for _, item := range st.Tables {
		itemPb, err := tableToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tablesPb = append(tablesPb, *itemPb)
		}
	}
	pb.Tables = tablesPb

	var volumesPb []volumePb
	for _, item := range st.Volumes {
		itemPb, err := volumeToPb(&item)
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

func (st *ListProviderShareAssetsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProviderShareAssetsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProviderShareAssetsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProviderShareAssetsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listProviderShareAssetsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProviderShareAssetsResponsePb struct {
	// The list of functions in the share.
	Functions []deltaSharingFunctionPb `json:"functions,omitempty"`
	// The list of notebooks in the share.
	Notebooks []notebookFilePb `json:"notebooks,omitempty"`
	// The list of tables in the share.
	Tables []tablePb `json:"tables,omitempty"`
	// The list of volumes in the share.
	Volumes []volumePb `json:"volumes,omitempty"`
}

func listProviderShareAssetsResponseFromPb(pb *listProviderShareAssetsResponsePb) (*ListProviderShareAssetsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderShareAssetsResponse{}

	var functionsField []DeltaSharingFunction
	for _, itemPb := range pb.Functions {
		item, err := deltaSharingFunctionFromPb(&itemPb)
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
		item, err := notebookFileFromPb(&itemPb)
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
		item, err := tableFromPb(&itemPb)
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
		item, err := volumeFromPb(&itemPb)
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
	NextPageToken string
	// An array of provider shares.
	// Wire name: 'shares'
	Shares []ProviderShare

	ForceSendFields []string `tf:"-"`
}

func listProviderSharesResponseToPb(st *ListProviderSharesResponse) (*listProviderSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderSharesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var sharesPb []providerSharePb
	for _, item := range st.Shares {
		itemPb, err := providerShareToPb(&item)
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

func (st *ListProviderSharesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProviderSharesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProviderSharesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProviderSharesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listProviderSharesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProviderSharesResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider shares.
	Shares []providerSharePb `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProviderSharesResponseFromPb(pb *listProviderSharesResponsePb) (*ListProviderSharesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProviderSharesResponse{}
	st.NextPageToken = pb.NextPageToken

	var sharesField []ProviderShare
	for _, itemPb := range pb.Shares {
		item, err := providerShareFromPb(&itemPb)
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

func (st *listProviderSharesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProviderSharesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List providers
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
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListProvidersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProvidersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProvidersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProvidersRequest) MarshalJSON() ([]byte, error) {
	pb, err := listProvidersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProvidersRequestPb struct {
	// If not provided, all providers will be returned. If no providers exist
	// with this ID, no results will be returned.
	DataProviderGlobalMetastoreId string `json:"-" url:"data_provider_global_metastore_id,omitempty"`
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of provider information objects.
	// Wire name: 'providers'
	Providers []ProviderInfo

	ForceSendFields []string `tf:"-"`
}

func listProvidersResponseToPb(st *ListProvidersResponse) (*listProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var providersPb []providerInfoPb
	for _, item := range st.Providers {
		itemPb, err := providerInfoToPb(&item)
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

func (st *ListProvidersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listProvidersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listProvidersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListProvidersResponse) MarshalJSON() ([]byte, error) {
	pb, err := listProvidersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listProvidersResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider information objects.
	Providers []providerInfoPb `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listProvidersResponseFromPb(pb *listProvidersResponsePb) (*ListProvidersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListProvidersResponse{}
	st.NextPageToken = pb.NextPageToken

	var providersField []ProviderInfo
	for _, itemPb := range pb.Providers {
		item, err := providerInfoFromPb(&itemPb)
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

func (st *listProvidersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listProvidersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List share recipients
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
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListRecipientsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRecipientsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRecipientsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRecipientsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listRecipientsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRecipientsRequestPb struct {
	// If not provided, all recipients will be returned. If no recipients exist
	// with this ID, no results will be returned.
	DataRecipientGlobalMetastoreId string `json:"-" url:"data_recipient_global_metastore_id,omitempty"`
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of recipient information objects.
	// Wire name: 'recipients'
	Recipients []RecipientInfo

	ForceSendFields []string `tf:"-"`
}

func listRecipientsResponseToPb(st *ListRecipientsResponse) (*listRecipientsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRecipientsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var recipientsPb []recipientInfoPb
	for _, item := range st.Recipients {
		itemPb, err := recipientInfoToPb(&item)
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

func (st *ListRecipientsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRecipientsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRecipientsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRecipientsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listRecipientsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRecipientsResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of recipient information objects.
	Recipients []recipientInfoPb `json:"recipients,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRecipientsResponseFromPb(pb *listRecipientsResponsePb) (*ListRecipientsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRecipientsResponse{}
	st.NextPageToken = pb.NextPageToken

	var recipientsField []RecipientInfo
	for _, itemPb := range pb.Recipients {
		item, err := recipientInfoFromPb(&itemPb)
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

func (st *listRecipientsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRecipientsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List shares by Provider
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
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListSharesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSharesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSharesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSharesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSharesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSharesRequestPb struct {
	// Maximum number of shares to return. - when set to 0, the page length is
	// set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid shares are returned (not
	// recommended). - Note: The number of returned shares might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further shares can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Name of the provider in which to list shares.
	Name string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string
	// An array of data share information objects.
	// Wire name: 'shares'
	Shares []ShareInfo

	ForceSendFields []string `tf:"-"`
}

func listSharesResponseToPb(st *ListSharesResponse) (*listSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSharesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var sharesPb []shareInfoPb
	for _, item := range st.Shares {
		itemPb, err := shareInfoToPb(&item)
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

func (st *ListSharesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSharesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSharesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSharesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listSharesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listSharesResponsePb struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share information objects.
	Shares []shareInfoPb `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSharesResponseFromPb(pb *listSharesResponsePb) (*ListSharesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSharesResponse{}
	st.NextPageToken = pb.NextPageToken

	var sharesField []ShareInfo
	for _, itemPb := range pb.Shares {
		item, err := shareInfoFromPb(&itemPb)
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

func (st *listSharesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSharesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookFile struct {
	// The comment of the notebook file.
	// Wire name: 'comment'
	Comment string
	// The id of the notebook file.
	// Wire name: 'id'
	Id string
	// Name of the notebook file.
	// Wire name: 'name'
	Name string
	// The name of the share that the notebook file belongs to.
	// Wire name: 'share'
	Share string
	// The id of the share that the notebook file belongs to.
	// Wire name: 'share_id'
	ShareId string
	// The tags of the notebook file.
	// Wire name: 'tags'
	Tags []catalog.TagKeyValue

	ForceSendFields []string `tf:"-"`
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

	var tagsPb []catalog.TagKeyValuePb
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

func (st *NotebookFile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notebookFilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notebookFileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotebookFile) MarshalJSON() ([]byte, error) {
	pb, err := notebookFileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type notebookFilePb struct {
	// The comment of the notebook file.
	Comment string `json:"comment,omitempty"`
	// The id of the notebook file.
	Id string `json:"id,omitempty"`
	// Name of the notebook file.
	Name string `json:"name,omitempty"`
	// The name of the share that the notebook file belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the notebook file belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The tags of the notebook file.
	Tags []catalog.TagKeyValuePb `json:"tags,omitempty"`

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

func (st *notebookFilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookFilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Partition struct {
	// An array of partition values.
	// Wire name: 'values'
	Values []PartitionValue
}

func PartitionToPb(st *Partition) (*PartitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &PartitionPb{}

	var valuesPb []partitionValuePb
	for _, item := range st.Values {
		itemPb, err := partitionValueToPb(&item)
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

func (st *Partition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &PartitionPb{}
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

func (st Partition) MarshalJSON() ([]byte, error) {
	pb, err := PartitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PartitionPb struct {
	// An array of partition values.
	Values []partitionValuePb `json:"values,omitempty"`
}

func PartitionFromPb(pb *PartitionPb) (*Partition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Partition{}

	var valuesField []PartitionValue
	for _, itemPb := range pb.Values {
		item, err := partitionValueFromPb(&itemPb)
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
	Name string
	// The operator to apply for the value.
	// Wire name: 'op'
	Op PartitionValueOp
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	// Wire name: 'recipient_property_key'
	RecipientPropertyKey string
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
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

func (st *PartitionValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &partitionValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := partitionValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PartitionValue) MarshalJSON() ([]byte, error) {
	pb, err := partitionValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type partitionValuePb struct {
	// The name of the partition column.
	Name string `json:"name,omitempty"`
	// The operator to apply for the value.
	Op PartitionValueOp `json:"op,omitempty"`
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey string `json:"recipient_property_key,omitempty"`
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value string `json:"value,omitempty"`

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

type PartitionValueOp string
type partitionValueOpPb string

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

// Type always returns PartitionValueOp to satisfy [pflag.Value] interface
func (f *PartitionValueOp) Type() string {
	return "PartitionValueOp"
}

func partitionValueOpToPb(st *PartitionValueOp) (*partitionValueOpPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := partitionValueOpPb(*st)
	return &pb, nil
}

func partitionValueOpFromPb(pb *partitionValueOpPb) (*PartitionValueOp, error) {
	if pb == nil {
		return nil, nil
	}
	st := PartitionValueOp(*pb)
	return &st, nil
}

type PermissionsChange struct {
	// The set of privileges to add.
	// Wire name: 'add'
	Add []string
	// The principal whose privileges we are changing.
	// Wire name: 'principal'
	Principal string
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove []string

	ForceSendFields []string `tf:"-"`
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

func (st *PermissionsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permissionsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permissionsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermissionsChange) MarshalJSON() ([]byte, error) {
	pb, err := permissionsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type permissionsChangePb struct {
	// The set of privileges to add.
	Add []string `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	Remove []string `json:"remove,omitempty"`

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

type Privilege string
type privilegePb string

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

// Type always returns Privilege to satisfy [pflag.Value] interface
func (f *Privilege) Type() string {
	return "Privilege"
}

func privilegeToPb(st *Privilege) (*privilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := privilegePb(*st)
	return &pb, nil
}

func privilegeFromPb(pb *privilegePb) (*Privilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := Privilege(*pb)
	return &st, nil
}

type PrivilegeAssignment struct {
	// The principal (user email address or group name).
	// Wire name: 'principal'
	Principal string
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges []Privilege

	ForceSendFields []string `tf:"-"`
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

func (st *PrivilegeAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &privilegeAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := privilegeAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PrivilegeAssignment) MarshalJSON() ([]byte, error) {
	pb, err := privilegeAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type privilegeAssignmentPb struct {
	// The principal (user email address or group name).
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
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

type ProviderInfo struct {
	// The delta sharing authentication type.
	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'cloud'
	Cloud string
	// Description about the provider.
	// Wire name: 'comment'
	Comment string
	// Time at which this Provider was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of Provider creator.
	// Wire name: 'created_by'
	CreatedBy string
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_provider_global_metastore_id'
	DataProviderGlobalMetastoreId string
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	// Wire name: 'metastore_id'
	MetastoreId string
	// The name of the Provider.
	// Wire name: 'name'
	Name string
	// Username of Provider owner.
	// Wire name: 'owner'
	Owner string
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	// Wire name: 'recipient_profile'
	RecipientProfile *RecipientProfile
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'region'
	Region string
	// Time at which this Provider was created, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of user who last modified Provider.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
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

	recipientProfilePb, err := recipientProfileToPb(st.RecipientProfile)
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

func (st *ProviderInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &providerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := providerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ProviderInfo) MarshalJSON() ([]byte, error) {
	pb, err := providerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type providerInfoPb struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud string `json:"cloud,omitempty"`
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of Provider creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	DataProviderGlobalMetastoreId string `json:"data_provider_global_metastore_id,omitempty"`
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId string `json:"metastore_id,omitempty"`
	// The name of the Provider.
	Name string `json:"name,omitempty"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	RecipientProfile *recipientProfilePb `json:"recipient_profile,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region string `json:"region,omitempty"`
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of user who last modified Provider.
	UpdatedBy string `json:"updated_by,omitempty"`

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
	recipientProfileField, err := recipientProfileFromPb(pb.RecipientProfile)
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

func (st *providerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st providerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ProviderShare struct {
	// The name of the Provider Share.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
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

func (st *ProviderShare) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &providerSharePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := providerShareFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ProviderShare) MarshalJSON() ([]byte, error) {
	pb, err := providerShareToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type providerSharePb struct {
	// The name of the Provider Share.
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

type RecipientInfo struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	// Wire name: 'activated'
	Activated bool
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string
	// The delta sharing authentication type.
	// Wire name: 'authentication_type'
	AuthenticationType AuthenticationType
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'cloud'
	Cloud string
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string
	// Time at which this recipient was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of recipient creator.
	// Wire name: 'created_by'
	CreatedBy string
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	// Wire name: 'data_recipient_global_metastore_id'
	DataRecipientGlobalMetastoreId string
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'metastore_id'
	MetastoreId string
	// Name of Recipient.
	// Wire name: 'name'
	Name string
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'region'
	Region string
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	// Wire name: 'sharing_code'
	SharingCode string
	// This field is only present when the __authentication_type__ is **TOKEN**.
	// Wire name: 'tokens'
	Tokens []RecipientTokenInfo
	// Time at which the recipient was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of recipient updater.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
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

	ipAccessListPb, err := ipAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	pb.MetastoreId = st.MetastoreId

	pb.Name = st.Name

	pb.Owner = st.Owner

	propertiesKvpairsPb, err := securablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}

	pb.Region = st.Region

	pb.SharingCode = st.SharingCode

	var tokensPb []recipientTokenInfoPb
	for _, item := range st.Tokens {
		itemPb, err := recipientTokenInfoToPb(&item)
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

func (st *RecipientInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &recipientInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := recipientInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RecipientInfo) MarshalJSON() ([]byte, error) {
	pb, err := recipientInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type recipientInfoPb struct {
	// A boolean status field showing whether the Recipient's activation URL has
	// been exercised or not.
	Activated bool `json:"activated,omitempty"`
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType `json:"authentication_type,omitempty"`
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Cloud string `json:"cloud,omitempty"`
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient creator.
	CreatedBy string `json:"created_by,omitempty"`
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string `json:"data_recipient_global_metastore_id,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// IP Access List
	IpAccessList *ipAccessListPb `json:"ip_access_list,omitempty"`
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	MetastoreId string `json:"metastore_id,omitempty"`
	// Name of Recipient.
	Name string `json:"name,omitempty"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *securablePropertiesKvPairsPb `json:"properties_kvpairs,omitempty"`
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region string `json:"region,omitempty"`
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string `json:"sharing_code,omitempty"`
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []recipientTokenInfoPb `json:"tokens,omitempty"`
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient updater.
	UpdatedBy string `json:"updated_by,omitempty"`

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
	ipAccessListField, err := ipAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	propertiesKvpairsField, err := securablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
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
		item, err := recipientTokenInfoFromPb(&itemPb)
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

func (st *recipientInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st recipientInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RecipientProfile struct {
	// The token used to authorize the recipient.
	// Wire name: 'bearer_token'
	BearerToken string
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string
	// The version number of the recipient's credentials on a share.
	// Wire name: 'share_credentials_version'
	ShareCredentialsVersion int

	ForceSendFields []string `tf:"-"`
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

func (st *RecipientProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &recipientProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := recipientProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RecipientProfile) MarshalJSON() ([]byte, error) {
	pb, err := recipientProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type recipientProfilePb struct {
	// The token used to authorize the recipient.
	BearerToken string `json:"bearer_token,omitempty"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `json:"endpoint,omitempty"`
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion int `json:"share_credentials_version,omitempty"`

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

type RecipientTokenInfo struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	// Wire name: 'activation_url'
	ActivationUrl string
	// Time at which this recipient token was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of recipient token creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// Unique ID of the recipient token.
	// Wire name: 'id'
	Id string
	// Time at which this recipient token was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of recipient token updater.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
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

func (st *RecipientTokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &recipientTokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := recipientTokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RecipientTokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := recipientTokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type recipientTokenInfoPb struct {
	// Full activation URL to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string `json:"activation_url,omitempty"`
	// Time at which this recipient token was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of recipient token creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// Unique ID of the recipient token.
	Id string `json:"id,omitempty"`
	// Time at which this recipient token was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of recipient token updater.
	UpdatedBy string `json:"updated_by,omitempty"`

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

type RegisteredModelAlias struct {
	// Name of the alias.
	// Wire name: 'alias_name'
	AliasName string
	// Numeric model version that alias will reference.
	// Wire name: 'version_num'
	VersionNum int64

	ForceSendFields []string `tf:"-"`
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

func (st *RegisteredModelAlias) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelAliasPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelAliasFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelAlias) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelAliasToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelAliasPb struct {
	// Name of the alias.
	AliasName string `json:"alias_name,omitempty"`
	// Numeric model version that alias will reference.
	VersionNum int64 `json:"version_num,omitempty"`

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

// Get an access token
type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	// Wire name: 'activation_url'
	ActivationUrl string `tf:"-"`
}

func retrieveTokenRequestToPb(st *RetrieveTokenRequest) (*retrieveTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &retrieveTokenRequestPb{}
	pb.ActivationUrl = st.ActivationUrl

	return pb, nil
}

func (st *RetrieveTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &retrieveTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := retrieveTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RetrieveTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := retrieveTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type retrieveTokenRequestPb struct {
	// The one time activation url. It also accepts activation token.
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

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	// Wire name: 'bearerToken'
	BearerToken string
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string
	// Expiration timestamp of the token in epoch milliseconds.
	// Wire name: 'expirationTime'
	ExpirationTime string
	// These field names must follow the delta sharing protocol.
	// Wire name: 'shareCredentialsVersion'
	ShareCredentialsVersion int

	ForceSendFields []string `tf:"-"`
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

func (st *RetrieveTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &retrieveTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := retrieveTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RetrieveTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := retrieveTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type retrieveTokenResponsePb struct {
	// The token used to authorize the recipient.
	BearerToken string `json:"bearerToken,omitempty"`
	// The endpoint for the share to be used by the recipient.
	Endpoint string `json:"endpoint,omitempty"`
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime string `json:"expirationTime,omitempty"`
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

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

type RotateRecipientToken struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	// Wire name: 'existing_token_expire_in_seconds'
	ExistingTokenExpireInSeconds int64
	// The name of the Recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
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

func (st *RotateRecipientToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rotateRecipientTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rotateRecipientTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RotateRecipientToken) MarshalJSON() ([]byte, error) {
	pb, err := rotateRecipientTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type rotateRecipientTokenPb struct {
	// The expiration time of the bearer token in ISO 8601 format. This will set
	// the expiration_time of existing token only to a smaller timestamp, it
	// cannot extend the expiration_time. Use 0 to expire the existing token
	// immediately, negative number will return an error.
	ExistingTokenExpireInSeconds int64 `json:"existing_token_expire_in_seconds"`
	// The name of the Recipient.
	Name string `json:"-" url:"-"`
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

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string
}

func securablePropertiesKvPairsToPb(st *SecurablePropertiesKvPairs) (*securablePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &securablePropertiesKvPairsPb{}
	pb.Properties = st.Properties

	return pb, nil
}

func (st *SecurablePropertiesKvPairs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &securablePropertiesKvPairsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := securablePropertiesKvPairsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SecurablePropertiesKvPairs) MarshalJSON() ([]byte, error) {
	pb, err := securablePropertiesKvPairsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type securablePropertiesKvPairsPb struct {
	// A map of key-value properties attached to the securable.
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

type ShareInfo struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// Time at which this share was created, in epoch milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// Username of share creator.
	// Wire name: 'created_by'
	CreatedBy string
	// Name of the share.
	// Wire name: 'name'
	Name string
	// A list of shared data objects within the share.
	// Wire name: 'objects'
	Objects []SharedDataObject
	// Username of current owner of share.
	// Wire name: 'owner'
	Owner string
	// Storage Location URL (full path) for the share.
	// Wire name: 'storage_location'
	StorageLocation string
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string
	// Time at which this share was updated, in epoch milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// Username of share updater.
	// Wire name: 'updated_by'
	UpdatedBy string

	ForceSendFields []string `tf:"-"`
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

	var objectsPb []sharedDataObjectPb
	for _, item := range st.Objects {
		itemPb, err := sharedDataObjectToPb(&item)
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

func (st *ShareInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &shareInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := shareInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ShareInfo) MarshalJSON() ([]byte, error) {
	pb, err := shareInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type shareInfoPb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// Username of share creator.
	CreatedBy string `json:"created_by,omitempty"`
	// Name of the share.
	Name string `json:"name,omitempty"`
	// A list of shared data objects within the share.
	Objects []sharedDataObjectPb `json:"objects,omitempty"`
	// Username of current owner of share.
	Owner string `json:"owner,omitempty"`
	// Storage Location URL (full path) for the share.
	StorageLocation string `json:"storage_location,omitempty"`
	// Storage root URL for the share.
	StorageRoot string `json:"storage_root,omitempty"`
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Username of share updater.
	UpdatedBy string `json:"updated_by,omitempty"`

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

	var objectsField []SharedDataObject
	for _, itemPb := range pb.Objects {
		item, err := sharedDataObjectFromPb(&itemPb)
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

func (st *shareInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st shareInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get recipient share permissions
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
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *SharePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sharePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SharePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := sharePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sharePermissionsRequestPb struct {
	// Maximum number of permissions to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid permissions are returned (not
	// recommended). - Note: The number of returned permissions might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further permissions can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// The name of the Recipient.
	Name string `json:"-" url:"-"`
	// Opaque pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

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

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment
	// The share name.
	// Wire name: 'share_name'
	ShareName string

	ForceSendFields []string `tf:"-"`
}

func shareToPrivilegeAssignmentToPb(st *ShareToPrivilegeAssignment) (*shareToPrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &shareToPrivilegeAssignmentPb{}

	var privilegeAssignmentsPb []privilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := privilegeAssignmentToPb(&item)
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

func (st *ShareToPrivilegeAssignment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &shareToPrivilegeAssignmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := shareToPrivilegeAssignmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ShareToPrivilegeAssignment) MarshalJSON() ([]byte, error) {
	pb, err := shareToPrivilegeAssignmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type shareToPrivilegeAssignmentPb struct {
	// The privileges assigned to the principal.
	PrivilegeAssignments []privilegeAssignmentPb `json:"privilege_assignments,omitempty"`
	// The share name.
	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func shareToPrivilegeAssignmentFromPb(pb *shareToPrivilegeAssignmentPb) (*ShareToPrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ShareToPrivilegeAssignment{}

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := privilegeAssignmentFromPb(&itemPb)
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

func (st *shareToPrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st shareToPrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SharedDataObject struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	// Wire name: 'added_at'
	AddedAt int64
	// Username of the sharer.
	// Wire name: 'added_by'
	AddedBy string
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	// Wire name: 'cdf_enabled'
	CdfEnabled bool
	// A user-provided comment when adding the data object to the share.
	// Wire name: 'comment'
	Comment string
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	// Wire name: 'content'
	Content string
	// The type of the data object.
	// Wire name: 'data_object_type'
	DataObjectType SharedDataObjectDataObjectType
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	// Wire name: 'history_data_sharing_status'
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`,
	// Wire name: 'name'
	Name string
	// Array of partitions for the shared data.
	// Wire name: 'partitions'
	Partitions []Partition
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	// Wire name: 'shared_as'
	SharedAs string
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	// Wire name: 'start_version'
	StartVersion int64
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	// Wire name: 'status'
	Status SharedDataObjectStatus
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
	// Wire name: 'string_shared_as'
	StringSharedAs string

	ForceSendFields []string `tf:"-"`
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

	var partitionsPb []PartitionPb
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

	pb.Status = st.Status

	pb.StringSharedAs = st.StringSharedAs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SharedDataObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharedDataObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sharedDataObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SharedDataObject) MarshalJSON() ([]byte, error) {
	pb, err := sharedDataObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sharedDataObjectPb struct {
	// The time when this data object is added to the share, in epoch
	// milliseconds.
	AddedAt int64 `json:"added_at,omitempty"`
	// Username of the sharer.
	AddedBy string `json:"added_by,omitempty"`
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled bool `json:"cdf_enabled,omitempty"`
	// A user-provided comment when adding the data object to the share.
	Comment string `json:"comment,omitempty"`
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content string `json:"content,omitempty"`
	// The type of the data object.
	DataObjectType SharedDataObjectDataObjectType `json:"data_object_type,omitempty"`
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus `json:"history_data_sharing_status,omitempty"`
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`,
	Name string `json:"name"`
	// Array of partitions for the shared data.
	Partitions []PartitionPb `json:"partitions,omitempty"`
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs string `json:"shared_as,omitempty"`
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion int64 `json:"start_version,omitempty"`
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status SharedDataObjectStatus `json:"status,omitempty"`
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
	StringSharedAs string `json:"string_shared_as,omitempty"`

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

type SharedDataObjectDataObjectType string
type sharedDataObjectDataObjectTypePb string

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

// Type always returns SharedDataObjectDataObjectType to satisfy [pflag.Value] interface
func (f *SharedDataObjectDataObjectType) Type() string {
	return "SharedDataObjectDataObjectType"
}

func sharedDataObjectDataObjectTypeToPb(st *SharedDataObjectDataObjectType) (*sharedDataObjectDataObjectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharedDataObjectDataObjectTypePb(*st)
	return &pb, nil
}

func sharedDataObjectDataObjectTypeFromPb(pb *sharedDataObjectDataObjectTypePb) (*SharedDataObjectDataObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectDataObjectType(*pb)
	return &st, nil
}

type SharedDataObjectHistoryDataSharingStatus string
type sharedDataObjectHistoryDataSharingStatusPb string

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

// Type always returns SharedDataObjectHistoryDataSharingStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectHistoryDataSharingStatus) Type() string {
	return "SharedDataObjectHistoryDataSharingStatus"
}

func sharedDataObjectHistoryDataSharingStatusToPb(st *SharedDataObjectHistoryDataSharingStatus) (*sharedDataObjectHistoryDataSharingStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharedDataObjectHistoryDataSharingStatusPb(*st)
	return &pb, nil
}

func sharedDataObjectHistoryDataSharingStatusFromPb(pb *sharedDataObjectHistoryDataSharingStatusPb) (*SharedDataObjectHistoryDataSharingStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectHistoryDataSharingStatus(*pb)
	return &st, nil
}

type SharedDataObjectStatus string
type sharedDataObjectStatusPb string

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

// Type always returns SharedDataObjectStatus to satisfy [pflag.Value] interface
func (f *SharedDataObjectStatus) Type() string {
	return "SharedDataObjectStatus"
}

func sharedDataObjectStatusToPb(st *SharedDataObjectStatus) (*sharedDataObjectStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharedDataObjectStatusPb(*st)
	return &pb, nil
}

func sharedDataObjectStatusFromPb(pb *sharedDataObjectStatusPb) (*SharedDataObjectStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectStatus(*pb)
	return &st, nil
}

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	// Wire name: 'action'
	Action SharedDataObjectUpdateAction
	// The data object that is being added, removed, or updated.
	// Wire name: 'data_object'
	DataObject *SharedDataObject
}

func sharedDataObjectUpdateToPb(st *SharedDataObjectUpdate) (*sharedDataObjectUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharedDataObjectUpdatePb{}
	pb.Action = st.Action

	dataObjectPb, err := sharedDataObjectToPb(st.DataObject)
	if err != nil {
		return nil, err
	}
	if dataObjectPb != nil {
		pb.DataObject = dataObjectPb
	}

	return pb, nil
}

func (st *SharedDataObjectUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sharedDataObjectUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sharedDataObjectUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SharedDataObjectUpdate) MarshalJSON() ([]byte, error) {
	pb, err := sharedDataObjectUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sharedDataObjectUpdatePb struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	Action SharedDataObjectUpdateAction `json:"action,omitempty"`
	// The data object that is being added, removed, or updated.
	DataObject *sharedDataObjectPb `json:"data_object,omitempty"`
}

func sharedDataObjectUpdateFromPb(pb *sharedDataObjectUpdatePb) (*SharedDataObjectUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SharedDataObjectUpdate{}
	st.Action = pb.Action
	dataObjectField, err := sharedDataObjectFromPb(pb.DataObject)
	if err != nil {
		return nil, err
	}
	if dataObjectField != nil {
		st.DataObject = dataObjectField
	}

	return st, nil
}

type SharedDataObjectUpdateAction string
type sharedDataObjectUpdateActionPb string

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

// Type always returns SharedDataObjectUpdateAction to satisfy [pflag.Value] interface
func (f *SharedDataObjectUpdateAction) Type() string {
	return "SharedDataObjectUpdateAction"
}

func sharedDataObjectUpdateActionToPb(st *SharedDataObjectUpdateAction) (*sharedDataObjectUpdateActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharedDataObjectUpdateActionPb(*st)
	return &pb, nil
}

func sharedDataObjectUpdateActionFromPb(pb *sharedDataObjectUpdateActionPb) (*SharedDataObjectUpdateAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedDataObjectUpdateAction(*pb)
	return &st, nil
}

// The SecurableKind of a delta-shared object.
type SharedSecurableKind string
type sharedSecurableKindPb string

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

// Type always returns SharedSecurableKind to satisfy [pflag.Value] interface
func (f *SharedSecurableKind) Type() string {
	return "SharedSecurableKind"
}

func sharedSecurableKindToPb(st *SharedSecurableKind) (*sharedSecurableKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sharedSecurableKindPb(*st)
	return &pb, nil
}

func sharedSecurableKindFromPb(pb *sharedSecurableKindPb) (*SharedSecurableKind, error) {
	if pb == nil {
		return nil, nil
	}
	st := SharedSecurableKind(*pb)
	return &st, nil
}

type Table struct {
	// The comment of the table.
	// Wire name: 'comment'
	Comment string
	// The id of the table.
	// Wire name: 'id'
	Id string
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	// Wire name: 'internal_attributes'
	InternalAttributes *TableInternalAttributes
	// The name of a materialized table.
	// Wire name: 'materialized_table_name'
	MaterializedTableName string
	// The name of the table.
	// Wire name: 'name'
	Name string
	// The name of the schema that the table belongs to.
	// Wire name: 'schema'
	Schema string
	// The name of the share that the table belongs to.
	// Wire name: 'share'
	Share string
	// The id of the share that the table belongs to.
	// Wire name: 'share_id'
	ShareId string
	// The Tags of the table.
	// Wire name: 'tags'
	Tags []catalog.TagKeyValue

	ForceSendFields []string `tf:"-"`
}

func tableToPb(st *Table) (*tablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tablePb{}
	pb.Comment = st.Comment

	pb.Id = st.Id

	internalAttributesPb, err := tableInternalAttributesToPb(st.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesPb != nil {
		pb.InternalAttributes = internalAttributesPb
	}

	pb.MaterializedTableName = st.MaterializedTableName

	pb.Name = st.Name

	pb.Schema = st.Schema

	pb.Share = st.Share

	pb.ShareId = st.ShareId

	var tagsPb []catalog.TagKeyValuePb
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

func (st *Table) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Table) MarshalJSON() ([]byte, error) {
	pb, err := tableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tablePb struct {
	// The comment of the table.
	Comment string `json:"comment,omitempty"`
	// The id of the table.
	Id string `json:"id,omitempty"`
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes *tableInternalAttributesPb `json:"internal_attributes,omitempty"`
	// The name of a materialized table.
	MaterializedTableName string `json:"materialized_table_name,omitempty"`
	// The name of the table.
	Name string `json:"name,omitempty"`
	// The name of the schema that the table belongs to.
	Schema string `json:"schema,omitempty"`
	// The name of the share that the table belongs to.
	Share string `json:"share,omitempty"`
	// The id of the share that the table belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The Tags of the table.
	Tags []catalog.TagKeyValuePb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableFromPb(pb *tablePb) (*Table, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Table{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	internalAttributesField, err := tableInternalAttributesFromPb(pb.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesField != nil {
		st.InternalAttributes = internalAttributesField
	}
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

func (st *tablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	ParentStorageLocation string
	// The cloud storage location of a shard table with DIRECTORY_BASED_TABLE
	// type.
	// Wire name: 'storage_location'
	StorageLocation string
	// The type of the shared table.
	// Wire name: 'type'
	Type TableInternalAttributesSharedTableType
	// The view definition of a shared view. DEPRECATED.
	// Wire name: 'view_definition'
	ViewDefinition string

	ForceSendFields []string `tf:"-"`
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

func (st *TableInternalAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableInternalAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableInternalAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableInternalAttributes) MarshalJSON() ([]byte, error) {
	pb, err := tableInternalAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableInternalAttributesPb struct {
	// Will be populated in the reconciliation response for VIEW and
	// FOREIGN_TABLE, with the value of the parent UC entity's storage_location,
	// following the same logic as getManagedEntityPath in
	// CreateStagingTableHandler, which is used to store the materialized table
	// for a shared VIEW/FOREIGN_TABLE for D2O queries. The value will be used
	// on the recipient side to be whitelisted when SEG is enabled on the
	// workspace of the recipient, to allow the recipient users to query this
	// shared VIEW/FOREIGN_TABLE.
	ParentStorageLocation string `json:"parent_storage_location,omitempty"`
	// The cloud storage location of a shard table with DIRECTORY_BASED_TABLE
	// type.
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the shared table.
	Type TableInternalAttributesSharedTableType `json:"type,omitempty"`
	// The view definition of a shared view. DEPRECATED.
	ViewDefinition string `json:"view_definition,omitempty"`

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

type TableInternalAttributesSharedTableType string
type tableInternalAttributesSharedTableTypePb string

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
	case `DIRECTORY_BASED_TABLE`, `FILE_BASED_TABLE`, `FOREIGN_TABLE`, `MATERIALIZED_VIEW`, `STREAMING_TABLE`, `VIEW`:
		*f = TableInternalAttributesSharedTableType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECTORY_BASED_TABLE", "FILE_BASED_TABLE", "FOREIGN_TABLE", "MATERIALIZED_VIEW", "STREAMING_TABLE", "VIEW"`, v)
	}
}

// Type always returns TableInternalAttributesSharedTableType to satisfy [pflag.Value] interface
func (f *TableInternalAttributesSharedTableType) Type() string {
	return "TableInternalAttributesSharedTableType"
}

func tableInternalAttributesSharedTableTypeToPb(st *TableInternalAttributesSharedTableType) (*tableInternalAttributesSharedTableTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := tableInternalAttributesSharedTableTypePb(*st)
	return &pb, nil
}

func tableInternalAttributesSharedTableTypeFromPb(pb *tableInternalAttributesSharedTableTypePb) (*TableInternalAttributesSharedTableType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableInternalAttributesSharedTableType(*pb)
	return &st, nil
}

type UpdateProvider struct {
	// Description about the provider.
	// Wire name: 'comment'
	Comment string
	// Name of the provider.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the provider.
	// Wire name: 'new_name'
	NewName string
	// Username of Provider owner.
	// Wire name: 'owner'
	Owner string
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	// Wire name: 'recipient_profile_str'
	RecipientProfileStr string

	ForceSendFields []string `tf:"-"`
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

func (st *UpdateProvider) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateProviderPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateProviderFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateProvider) MarshalJSON() ([]byte, error) {
	pb, err := updateProviderToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateProviderPb struct {
	// Description about the provider.
	Comment string `json:"comment,omitempty"`
	// Name of the provider.
	Name string `json:"-" url:"-"`
	// New name for the provider.
	NewName string `json:"new_name,omitempty"`
	// Username of Provider owner.
	Owner string `json:"owner,omitempty"`
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
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

type UpdateRecipient struct {
	// Description about the recipient.
	// Wire name: 'comment'
	Comment string
	// Expiration timestamp of the token, in epoch milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64
	// IP Access List
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessList
	// Name of the recipient.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the recipient. .
	// Wire name: 'new_name'
	NewName string
	// Username of the recipient owner.
	// Wire name: 'owner'
	Owner string
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	// Wire name: 'properties_kvpairs'
	PropertiesKvpairs *SecurablePropertiesKvPairs

	ForceSendFields []string `tf:"-"`
}

func updateRecipientToPb(st *UpdateRecipient) (*updateRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRecipientPb{}
	pb.Comment = st.Comment

	pb.ExpirationTime = st.ExpirationTime

	ipAccessListPb, err := ipAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.Owner = st.Owner

	propertiesKvpairsPb, err := securablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateRecipient) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRecipientPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRecipientFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRecipient) MarshalJSON() ([]byte, error) {
	pb, err := updateRecipientToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateRecipientPb struct {
	// Description about the recipient.
	Comment string `json:"comment,omitempty"`
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`
	// IP Access List
	IpAccessList *ipAccessListPb `json:"ip_access_list,omitempty"`
	// Name of the recipient.
	Name string `json:"-" url:"-"`
	// New name for the recipient. .
	NewName string `json:"new_name,omitempty"`
	// Username of the recipient owner.
	Owner string `json:"owner,omitempty"`
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *securablePropertiesKvPairsPb `json:"properties_kvpairs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRecipientFromPb(pb *updateRecipientPb) (*UpdateRecipient, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRecipient{}
	st.Comment = pb.Comment
	st.ExpirationTime = pb.ExpirationTime
	ipAccessListField, err := ipAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	propertiesKvpairsField, err := securablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsField != nil {
		st.PropertiesKvpairs = propertiesKvpairsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRecipientPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRecipientPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`
	// New name for the share.
	// Wire name: 'new_name'
	NewName string
	// Username of current owner of share.
	// Wire name: 'owner'
	Owner string
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string
	// Array of shared data object updates.
	// Wire name: 'updates'
	Updates []SharedDataObjectUpdate

	ForceSendFields []string `tf:"-"`
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

	var updatesPb []sharedDataObjectUpdatePb
	for _, item := range st.Updates {
		itemPb, err := sharedDataObjectUpdateToPb(&item)
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

func (st *UpdateShare) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateSharePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateShareFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateShare) MarshalJSON() ([]byte, error) {
	pb, err := updateShareToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateSharePb struct {
	// User-provided free-form text description.
	Comment string `json:"comment,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
	// New name for the share.
	NewName string `json:"new_name,omitempty"`
	// Username of current owner of share.
	Owner string `json:"owner,omitempty"`
	// Storage root URL for the share.
	StorageRoot string `json:"storage_root,omitempty"`
	// Array of shared data object updates.
	Updates []sharedDataObjectUpdatePb `json:"updates,omitempty"`

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

	var updatesField []SharedDataObjectUpdate
	for _, itemPb := range pb.Updates {
		item, err := sharedDataObjectUpdateFromPb(&itemPb)
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

func (st *updateSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateSharePermissions struct {
	// Array of permission changes.
	// Wire name: 'changes'
	Changes []PermissionsChange
	// The name of the share.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func updateSharePermissionsToPb(st *UpdateSharePermissions) (*updateSharePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSharePermissionsPb{}

	var changesPb []permissionsChangePb
	for _, item := range st.Changes {
		itemPb, err := permissionsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			changesPb = append(changesPb, *itemPb)
		}
	}
	pb.Changes = changesPb

	pb.Name = st.Name

	return pb, nil
}

func (st *UpdateSharePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateSharePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateSharePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateSharePermissions) MarshalJSON() ([]byte, error) {
	pb, err := updateSharePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateSharePermissionsPb struct {
	// Array of permission changes.
	Changes []permissionsChangePb `json:"changes,omitempty"`
	// The name of the share.
	Name string `json:"-" url:"-"`
}

func updateSharePermissionsFromPb(pb *updateSharePermissionsPb) (*UpdateSharePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSharePermissions{}

	var changesField []PermissionsChange
	for _, itemPb := range pb.Changes {
		item, err := permissionsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			changesField = append(changesField, *item)
		}
	}
	st.Changes = changesField
	st.Name = pb.Name

	return st, nil
}

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment
}

func updateSharePermissionsResponseToPb(st *UpdateSharePermissionsResponse) (*updateSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSharePermissionsResponsePb{}

	var privilegeAssignmentsPb []privilegeAssignmentPb
	for _, item := range st.PrivilegeAssignments {
		itemPb, err := privilegeAssignmentToPb(&item)
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

func (st *UpdateSharePermissionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateSharePermissionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateSharePermissionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateSharePermissionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateSharePermissionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateSharePermissionsResponsePb struct {
	// The privileges assigned to each principal
	PrivilegeAssignments []privilegeAssignmentPb `json:"privilege_assignments,omitempty"`
}

func updateSharePermissionsResponseFromPb(pb *updateSharePermissionsResponsePb) (*UpdateSharePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSharePermissionsResponse{}

	var privilegeAssignmentsField []PrivilegeAssignment
	for _, itemPb := range pb.PrivilegeAssignments {
		item, err := privilegeAssignmentFromPb(&itemPb)
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
	Comment string
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	// Wire name: 'id'
	Id string
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	// Wire name: 'internal_attributes'
	InternalAttributes *VolumeInternalAttributes
	// The name of the volume.
	// Wire name: 'name'
	Name string
	// The name of the schema that the volume belongs to.
	// Wire name: 'schema'
	Schema string
	// The name of the share that the volume belongs to.
	// Wire name: 'share'
	Share string
	// / The id of the share that the volume belongs to.
	// Wire name: 'share_id'
	ShareId string
	// The tags of the volume.
	// Wire name: 'tags'
	Tags []catalog.TagKeyValue

	ForceSendFields []string `tf:"-"`
}

func volumeToPb(st *Volume) (*volumePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumePb{}
	pb.Comment = st.Comment

	pb.Id = st.Id

	internalAttributesPb, err := volumeInternalAttributesToPb(st.InternalAttributes)
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

	var tagsPb []catalog.TagKeyValuePb
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

func (st *Volume) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &volumePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := volumeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Volume) MarshalJSON() ([]byte, error) {
	pb, err := volumeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type volumePb struct {
	// The comment of the volume.
	Comment string `json:"comment,omitempty"`
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	Id string `json:"id,omitempty"`
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes *volumeInternalAttributesPb `json:"internal_attributes,omitempty"`
	// The name of the volume.
	Name string `json:"name,omitempty"`
	// The name of the schema that the volume belongs to.
	Schema string `json:"schema,omitempty"`
	// The name of the share that the volume belongs to.
	Share string `json:"share,omitempty"`
	// / The id of the share that the volume belongs to.
	ShareId string `json:"share_id,omitempty"`
	// The tags of the volume.
	Tags []catalog.TagKeyValuePb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func volumeFromPb(pb *volumePb) (*Volume, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Volume{}
	st.Comment = pb.Comment
	st.Id = pb.Id
	internalAttributesField, err := volumeInternalAttributesFromPb(pb.InternalAttributes)
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

func (st *volumePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st volumePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Internal information for D2D sharing that should not be disclosed to external
// users.
type VolumeInternalAttributes struct {
	// The cloud storage location of the volume
	// Wire name: 'storage_location'
	StorageLocation string
	// The type of the shared volume.
	// Wire name: 'type'
	Type string

	ForceSendFields []string `tf:"-"`
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

func (st *VolumeInternalAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &volumeInternalAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := volumeInternalAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VolumeInternalAttributes) MarshalJSON() ([]byte, error) {
	pb, err := volumeInternalAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type volumeInternalAttributesPb struct {
	// The cloud storage location of the volume
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the shared volume.
	Type string `json:"type,omitempty"`

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
