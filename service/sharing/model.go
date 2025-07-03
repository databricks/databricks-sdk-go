// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharing

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/google/go-querystring/query"
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

type CreateFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be created.
	// Wire name: 'policy'
	Policy FederationPolicy `json:"policy"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being created.
	RecipientName string `json:"-" tf:"-"`
}

func (st *CreateFederationPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createFederationPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateProvider) EncodeValues(key string, v *url.Values) error {
	pb, err := createProviderToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	SharingCode string `json:"sharing_code,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateRecipient) EncodeValues(key string, v *url.Values) error {
	pb, err := createRecipientToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateShare struct {
	// User-provided free-form text description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// Name of the share.
	// Wire name: 'name'
	Name string `json:"name"`
	// Storage root URL for the share.
	// Wire name: 'storage_root'
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateShare) EncodeValues(key string, v *url.Values) error {
	pb, err := createShareToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be deleted.
	Name string `json:"-" tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being deleted.
	RecipientName string `json:"-" tf:"-"`
}

func (st *DeleteFederationPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteFederationPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteProviderRequest struct {
	// Name of the provider.
	Name string `json:"-" tf:"-"`
}

func (st *DeleteProviderRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteProviderRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name string `json:"-" tf:"-"`
}

func (st *DeleteRecipientRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRecipientRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteShareRequest struct {
	// The name of the share.
	Name string `json:"-" tf:"-"`
}

func (st *DeleteShareRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteShareRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Represents a UC dependency.
type DeltaSharingDependency struct {

	// Wire name: 'function'
	Function *DeltaSharingFunctionDependency `json:"function,omitempty"`

	// Wire name: 'table'
	Table *DeltaSharingTableDependency `json:"table,omitempty"`
}

func (st *DeltaSharingDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaSharingDependencyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Represents a list of dependencies.
type DeltaSharingDependencyList struct {
	// An array of Dependency.
	// Wire name: 'dependencies'
	Dependencies []DeltaSharingDependency `json:"dependencies,omitempty"`
}

func (st *DeltaSharingDependencyList) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaSharingDependencyListToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeltaSharingFunction) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaSharingFunctionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A Function in UC as a dependency.
type DeltaSharingFunctionDependency struct {

	// Wire name: 'function_name'
	FunctionName string `json:"function_name,omitempty"`

	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeltaSharingFunctionDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaSharingFunctionDependencyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A Table in UC as a dependency.
type DeltaSharingTableDependency struct {

	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`

	// Wire name: 'table_name'
	TableName string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeltaSharingTableDependency) EncodeValues(key string, v *url.Values) error {
	pb, err := deltaSharingTableDependencyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FederationPolicy struct {
	// Description of the policy. This is a user-provided description.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// System-generated timestamp indicating when the policy was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
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
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *FederationPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := federationPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *FederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &federationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := federationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := federationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *FunctionParameterInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := functionParameterInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FunctionParameterInfos struct {
	// The list of parameters of the function.
	// Wire name: 'parameters'
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func (st *FunctionParameterInfos) EncodeValues(key string, v *url.Values) error {
	pb, err := functionParameterInfosToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetActivationUrlInfoRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" tf:"-"`
}

func (st *GetActivationUrlInfoRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getActivationUrlInfoRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetFederationPolicyRequest struct {
	// Name of the policy. This is the name of the policy to be retrieved.
	Name string `json:"-" tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policy is being retrieved.
	RecipientName string `json:"-" tf:"-"`
}

func (st *GetFederationPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getFederationPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetProviderRequest struct {
	// Name of the provider.
	Name string `json:"-" tf:"-"`
}

func (st *GetProviderRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getProviderRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetRecipientRequest struct {
	// Name of the recipient.
	Name string `json:"-" tf:"-"`
}

func (st *GetRecipientRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRecipientRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share permissions for a recipient.
	// Wire name: 'permissions_out'
	PermissionsOut []ShareToPrivilegeAssignment `json:"permissions_out,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetRecipientSharePermissionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getRecipientSharePermissionsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetSharePermissionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getSharePermissionsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetShareRequest struct {
	// Query for data to include in the share.
	IncludeSharedData bool `json:"-" tf:"-"`
	// The name of the share.
	Name string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetShareRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getShareRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type IpAccessList struct {
	// Allowed IP Addresses in CIDR notation. Limit of 100.
	// Wire name: 'allowed_ip_addresses'
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

func (st *IpAccessList) EncodeValues(key string, v *url.Values) error {
	pb, err := ipAccessListToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListFederationPoliciesRequest struct {
	MaxResults int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`
	// Name of the recipient. This is the name of the recipient for which the
	// policies are being listed.
	RecipientName string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListFederationPoliciesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listFederationPoliciesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFederationPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFederationPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listFederationPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListFederationPoliciesResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'policies'
	Policies []FederationPolicy `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListFederationPoliciesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listFederationPoliciesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFederationPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFederationPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listFederationPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	VolumeMaxResults int `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProviderShareAssetsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listProviderShareAssetsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

func (st *ListProviderShareAssetsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listProviderShareAssetsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListProviderSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider shares.
	// Wire name: 'shares'
	Shares []ProviderShare `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProviderSharesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listProviderSharesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProvidersRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listProvidersRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListProvidersResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of provider information objects.
	// Wire name: 'providers'
	Providers []ProviderInfo `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListProvidersResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listProvidersResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListRecipientsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listRecipientsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListRecipientsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of recipient information objects.
	// Wire name: 'recipients'
	Recipients []RecipientInfo `json:"recipients,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListRecipientsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listRecipientsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListSharesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listSharesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListSharesResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// An array of data share information objects.
	// Wire name: 'shares'
	Shares []ShareInfo `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListSharesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listSharesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NotebookFile) EncodeValues(key string, v *url.Values) error {
	pb, err := notebookFileToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

func (st *OidcFederationPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := oidcFederationPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *OidcFederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oidcFederationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := oidcFederationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OidcFederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := oidcFederationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Partition struct {
	// An array of partition values.
	// Wire name: 'values'
	Values []PartitionValue `json:"values,omitempty"`
}

func (st *Partition) EncodeValues(key string, v *url.Values) error {
	pb, err := partitionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Partition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &partitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := partitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Partition) MarshalJSON() ([]byte, error) {
	pb, err := partitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PartitionValue) EncodeValues(key string, v *url.Values) error {
	pb, err := partitionValueToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type PermissionsChange struct {
	// The set of privileges to add.
	// Wire name: 'add'
	Add []string `json:"add,omitempty"`
	// The principal whose privileges we are changing.
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// The set of privileges to remove.
	// Wire name: 'remove'
	Remove []string `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PermissionsChange) EncodeValues(key string, v *url.Values) error {
	pb, err := permissionsChangeToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type PrivilegeAssignment struct {
	// The principal (user email address or group name). For deleted principals,
	// `principal` is empty while `principal_id` is populated.
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// The privileges assigned to the principal.
	// Wire name: 'privileges'
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PrivilegeAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := privilegeAssignmentToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ProviderInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := providerInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ProviderShare struct {
	// The name of the Provider Share.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ProviderShare) EncodeValues(key string, v *url.Values) error {
	pb, err := providerShareToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RecipientInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := recipientInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RecipientProfile struct {
	// The token used to authorize the recipient.
	// Wire name: 'bearer_token'
	BearerToken string `json:"bearer_token,omitempty"`
	// The endpoint for the share to be used by the recipient.
	// Wire name: 'endpoint'
	Endpoint string `json:"endpoint,omitempty"`
	// The version number of the recipient's credentials on a share.
	// Wire name: 'share_credentials_version'
	ShareCredentialsVersion int `json:"share_credentials_version,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RecipientProfile) EncodeValues(key string, v *url.Values) error {
	pb, err := recipientProfileToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RecipientTokenInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := recipientTokenInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RegisteredModelAlias struct {
	// Name of the alias.
	// Wire name: 'alias_name'
	AliasName string `json:"alias_name,omitempty"`
	// Numeric model version that alias will reference.
	// Wire name: 'version_num'
	VersionNum int64 `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RegisteredModelAlias) EncodeValues(key string, v *url.Values) error {
	pb, err := registeredModelAliasToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RetrieveTokenRequest struct {
	// The one time activation url. It also accepts activation token.
	ActivationUrl string `json:"-" tf:"-"`
}

func (st *RetrieveTokenRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := retrieveTokenRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ShareCredentialsVersion int `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RetrieveTokenResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := retrieveTokenResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

func (st *RotateRecipientToken) EncodeValues(key string, v *url.Values) error {
	pb, err := rotateRecipientTokenToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	// Wire name: 'properties'
	Properties map[string]string `json:"properties"`
}

func (st *SecurablePropertiesKvPairs) EncodeValues(key string, v *url.Values) error {
	pb, err := securablePropertiesKvPairsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ShareInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := shareInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SharePermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := sharePermissionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ShareToPrivilegeAssignment struct {
	// The privileges assigned to the principal.
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
	// The share name.
	// Wire name: 'share_name'
	ShareName string `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ShareToPrivilegeAssignment) EncodeValues(key string, v *url.Values) error {
	pb, err := shareToPrivilegeAssignmentToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	StringSharedAs string `json:"string_shared_as,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SharedDataObject) EncodeValues(key string, v *url.Values) error {
	pb, err := sharedDataObjectToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SharedDataObjectUpdate struct {
	// One of: **ADD**, **REMOVE**, **UPDATE**.
	// Wire name: 'action'
	Action SharedDataObjectUpdateAction `json:"action,omitempty"`
	// The data object that is being added, removed, or updated.
	// Wire name: 'data_object'
	DataObject *SharedDataObject `json:"data_object,omitempty"`
}

func (st *SharedDataObjectUpdate) EncodeValues(key string, v *url.Values) error {
	pb, err := sharedDataObjectUpdateToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Table) EncodeValues(key string, v *url.Values) error {
	pb, err := tableToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ViewDefinition string `json:"view_definition,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TableInternalAttributes) EncodeValues(key string, v *url.Values) error {
	pb, err := tableInternalAttributesToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	UpdateMask string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateFederationPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateFederationPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateProvider) EncodeValues(key string, v *url.Values) error {
	pb, err := updateProviderToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateRecipient) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRecipientToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Updates []SharedDataObjectUpdate `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateShare) EncodeValues(key string, v *url.Values) error {
	pb, err := updateShareToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateSharePermissions struct {
	// Array of permissions change objects.
	// Wire name: 'changes'
	Changes []PermissionsChange `json:"changes,omitempty"`
	// The name of the share.
	Name string `json:"-" tf:"-"`
	// Optional. Whether to return the latest permissions list of the share in
	// the response.
	// Wire name: 'omit_permissions_list'
	OmitPermissionsList bool `json:"omit_permissions_list,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateSharePermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := updateSharePermissionsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
	// Wire name: 'privilege_assignments'
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func (st *UpdateSharePermissionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := updateSharePermissionsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Tags []catalog.TagKeyValue `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Volume) EncodeValues(key string, v *url.Values) error {
	pb, err := volumeToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Internal information for D2D sharing that should not be disclosed to external
// users.
type VolumeInternalAttributes struct {
	// The cloud storage location of the volume
	// Wire name: 'storage_location'
	StorageLocation string `json:"storage_location,omitempty"`
	// The type of the shared volume.
	// Wire name: 'type'
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *VolumeInternalAttributes) EncodeValues(key string, v *url.Values) error {
	pb, err := volumeInternalAttributesToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
