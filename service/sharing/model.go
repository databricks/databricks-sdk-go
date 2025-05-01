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

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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
	AuthenticationType AuthenticationType
	// Description about the provider.
	Comment string
	// The name of the Provider.
	Name string
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string

	ForceSendFields []string
}

func createProviderToPb(st *CreateProvider) (*createProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createProviderPb{}
	authenticationTypePb, err := identity(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	recipientProfileStrPb, err := identity(&st.RecipientProfileStr)
	if err != nil {
		return nil, err
	}
	if recipientProfileStrPb != nil {
		pb.RecipientProfileStr = *recipientProfileStrPb
	}

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
	authenticationTypeField, err := identity(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	recipientProfileStrField, err := identity(&pb.RecipientProfileStr)
	if err != nil {
		return nil, err
	}
	if recipientProfileStrField != nil {
		st.RecipientProfileStr = *recipientProfileStrField
	}

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
	AuthenticationType AuthenticationType
	// Description about the recipient.
	Comment string
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64
	// IP Access List
	IpAccessList *IpAccessList
	// Name of Recipient.
	Name string
	// Username of the recipient owner.
	Owner string
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string

	ForceSendFields []string
}

func createRecipientToPb(st *CreateRecipient) (*createRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRecipientPb{}
	authenticationTypePb, err := identity(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	dataRecipientGlobalMetastoreIdPb, err := identity(&st.DataRecipientGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataRecipientGlobalMetastoreIdPb != nil {
		pb.DataRecipientGlobalMetastoreId = *dataRecipientGlobalMetastoreIdPb
	}

	expirationTimePb, err := identity(&st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

	ipAccessListPb, err := ipAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesKvpairsPb, err := securablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}

	sharingCodePb, err := identity(&st.SharingCode)
	if err != nil {
		return nil, err
	}
	if sharingCodePb != nil {
		pb.SharingCode = *sharingCodePb
	}

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
	authenticationTypeField, err := identity(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	dataRecipientGlobalMetastoreIdField, err := identity(&pb.DataRecipientGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataRecipientGlobalMetastoreIdField != nil {
		st.DataRecipientGlobalMetastoreId = *dataRecipientGlobalMetastoreIdField
	}
	expirationTimeField, err := identity(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
	ipAccessListField, err := ipAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	propertiesKvpairsField, err := securablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsField != nil {
		st.PropertiesKvpairs = propertiesKvpairsField
	}
	sharingCodeField, err := identity(&pb.SharingCode)
	if err != nil {
		return nil, err
	}
	if sharingCodeField != nil {
		st.SharingCode = *sharingCodeField
	}

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
	Comment string
	// Name of the share.
	Name string
	// Storage root URL for the share.
	StorageRoot string

	ForceSendFields []string
}

func createShareToPb(st *CreateShare) (*createSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSharePb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	storageRootPb, err := identity(&st.StorageRoot)
	if err != nil {
		return nil, err
	}
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	storageRootField, err := identity(&pb.StorageRoot)
	if err != nil {
		return nil, err
	}
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}

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
	Name string
}

func deleteProviderRequestToPb(st *DeleteProviderRequest) (*deleteProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteProviderRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Delete a share recipient
type DeleteRecipientRequest struct {
	// Name of the recipient.
	Name string
}

func deleteRecipientRequestToPb(st *DeleteRecipientRequest) (*deleteRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRecipientRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	Name string
}

func deleteShareRequestToPb(st *DeleteShareRequest) (*deleteShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteShareRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Represents a UC dependency.
type DeltaSharingDependency struct {
	// A Function in UC as a dependency.
	Function *DeltaSharingFunctionDependency
	// A Table in UC as a dependency.
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
	Aliases []RegisteredModelAlias
	// The comment of the function.
	Comment string
	// The data type of the function.
	DataType ColumnTypeName
	// The dependency list of the function.
	DependencyList *DeltaSharingDependencyList
	// The full data type of the function.
	FullDataType string
	// The id of the function.
	Id string
	// The function parameter information.
	InputParams *FunctionParameterInfos
	// The name of the function.
	Name string
	// The properties of the function.
	Properties string
	// The routine definition of the function.
	RoutineDefinition string
	// The name of the schema that the function belongs to.
	Schema string
	// The securable kind of the function.
	SecurableKind SharedSecurableKind
	// The name of the share that the function belongs to.
	Share string
	// The id of the share that the function belongs to.
	ShareId string
	// The storage location of the function.
	StorageLocation string
	// The tags of the function.
	Tags []catalog.TagKeyValue

	ForceSendFields []string
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

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	dataTypePb, err := identity(&st.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypePb != nil {
		pb.DataType = *dataTypePb
	}

	dependencyListPb, err := deltaSharingDependencyListToPb(st.DependencyList)
	if err != nil {
		return nil, err
	}
	if dependencyListPb != nil {
		pb.DependencyList = dependencyListPb
	}

	fullDataTypePb, err := identity(&st.FullDataType)
	if err != nil {
		return nil, err
	}
	if fullDataTypePb != nil {
		pb.FullDataType = *fullDataTypePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	inputParamsPb, err := functionParameterInfosToPb(st.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsPb != nil {
		pb.InputParams = inputParamsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	propertiesPb, err := identity(&st.Properties)
	if err != nil {
		return nil, err
	}
	if propertiesPb != nil {
		pb.Properties = *propertiesPb
	}

	routineDefinitionPb, err := identity(&st.RoutineDefinition)
	if err != nil {
		return nil, err
	}
	if routineDefinitionPb != nil {
		pb.RoutineDefinition = *routineDefinitionPb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	securableKindPb, err := identity(&st.SecurableKind)
	if err != nil {
		return nil, err
	}
	if securableKindPb != nil {
		pb.SecurableKind = *securableKindPb
	}

	sharePb, err := identity(&st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = *sharePb
	}

	shareIdPb, err := identity(&st.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdPb != nil {
		pb.ShareId = *shareIdPb
	}

	storageLocationPb, err := identity(&st.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	dataTypeField, err := identity(&pb.DataType)
	if err != nil {
		return nil, err
	}
	if dataTypeField != nil {
		st.DataType = *dataTypeField
	}
	dependencyListField, err := deltaSharingDependencyListFromPb(pb.DependencyList)
	if err != nil {
		return nil, err
	}
	if dependencyListField != nil {
		st.DependencyList = dependencyListField
	}
	fullDataTypeField, err := identity(&pb.FullDataType)
	if err != nil {
		return nil, err
	}
	if fullDataTypeField != nil {
		st.FullDataType = *fullDataTypeField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	inputParamsField, err := functionParameterInfosFromPb(pb.InputParams)
	if err != nil {
		return nil, err
	}
	if inputParamsField != nil {
		st.InputParams = inputParamsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	propertiesField, err := identity(&pb.Properties)
	if err != nil {
		return nil, err
	}
	if propertiesField != nil {
		st.Properties = *propertiesField
	}
	routineDefinitionField, err := identity(&pb.RoutineDefinition)
	if err != nil {
		return nil, err
	}
	if routineDefinitionField != nil {
		st.RoutineDefinition = *routineDefinitionField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	securableKindField, err := identity(&pb.SecurableKind)
	if err != nil {
		return nil, err
	}
	if securableKindField != nil {
		st.SecurableKind = *securableKindField
	}
	shareField, err := identity(&pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = *shareField
	}
	shareIdField, err := identity(&pb.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdField != nil {
		st.ShareId = *shareIdField
	}
	storageLocationField, err := identity(&pb.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}

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
	FunctionName string

	SchemaName string

	ForceSendFields []string
}

func deltaSharingFunctionDependencyToPb(st *DeltaSharingFunctionDependency) (*deltaSharingFunctionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingFunctionDependencyPb{}
	functionNamePb, err := identity(&st.FunctionName)
	if err != nil {
		return nil, err
	}
	if functionNamePb != nil {
		pb.FunctionName = *functionNamePb
	}

	schemaNamePb, err := identity(&st.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

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
	functionNameField, err := identity(&pb.FunctionName)
	if err != nil {
		return nil, err
	}
	if functionNameField != nil {
		st.FunctionName = *functionNameField
	}
	schemaNameField, err := identity(&pb.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

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
	SchemaName string

	TableName string

	ForceSendFields []string
}

func deltaSharingTableDependencyToPb(st *DeltaSharingTableDependency) (*deltaSharingTableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaSharingTableDependencyPb{}
	schemaNamePb, err := identity(&st.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	tableNamePb, err := identity(&st.TableName)
	if err != nil {
		return nil, err
	}
	if tableNamePb != nil {
		pb.TableName = *tableNamePb
	}

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
	schemaNameField, err := identity(&pb.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	tableNameField, err := identity(&pb.TableName)
	if err != nil {
		return nil, err
	}
	if tableNameField != nil {
		st.TableName = *tableNameField
	}

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
	Comment string
	// The name of the parameter.
	Name string
	// The default value of the parameter.
	ParameterDefault string
	// The mode of the function parameter.
	ParameterMode FunctionParameterMode
	// The type of the function parameter.
	ParameterType FunctionParameterType
	// The position of the parameter.
	Position int
	// The interval type of the parameter type.
	TypeIntervalType string
	// The type of the parameter in JSON format.
	TypeJson string
	// The type of the parameter in Enum format.
	TypeName ColumnTypeName
	// The precision of the parameter type.
	TypePrecision int
	// The scale of the parameter type.
	TypeScale int
	// The type of the parameter in text format.
	TypeText string

	ForceSendFields []string
}

func functionParameterInfoToPb(st *FunctionParameterInfo) (*functionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfoPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	parameterDefaultPb, err := identity(&st.ParameterDefault)
	if err != nil {
		return nil, err
	}
	if parameterDefaultPb != nil {
		pb.ParameterDefault = *parameterDefaultPb
	}

	parameterModePb, err := identity(&st.ParameterMode)
	if err != nil {
		return nil, err
	}
	if parameterModePb != nil {
		pb.ParameterMode = *parameterModePb
	}

	parameterTypePb, err := identity(&st.ParameterType)
	if err != nil {
		return nil, err
	}
	if parameterTypePb != nil {
		pb.ParameterType = *parameterTypePb
	}

	positionPb, err := identity(&st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = *positionPb
	}

	typeIntervalTypePb, err := identity(&st.TypeIntervalType)
	if err != nil {
		return nil, err
	}
	if typeIntervalTypePb != nil {
		pb.TypeIntervalType = *typeIntervalTypePb
	}

	typeJsonPb, err := identity(&st.TypeJson)
	if err != nil {
		return nil, err
	}
	if typeJsonPb != nil {
		pb.TypeJson = *typeJsonPb
	}

	typeNamePb, err := identity(&st.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}

	typePrecisionPb, err := identity(&st.TypePrecision)
	if err != nil {
		return nil, err
	}
	if typePrecisionPb != nil {
		pb.TypePrecision = *typePrecisionPb
	}

	typeScalePb, err := identity(&st.TypeScale)
	if err != nil {
		return nil, err
	}
	if typeScalePb != nil {
		pb.TypeScale = *typeScalePb
	}

	typeTextPb, err := identity(&st.TypeText)
	if err != nil {
		return nil, err
	}
	if typeTextPb != nil {
		pb.TypeText = *typeTextPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	parameterDefaultField, err := identity(&pb.ParameterDefault)
	if err != nil {
		return nil, err
	}
	if parameterDefaultField != nil {
		st.ParameterDefault = *parameterDefaultField
	}
	parameterModeField, err := identity(&pb.ParameterMode)
	if err != nil {
		return nil, err
	}
	if parameterModeField != nil {
		st.ParameterMode = *parameterModeField
	}
	parameterTypeField, err := identity(&pb.ParameterType)
	if err != nil {
		return nil, err
	}
	if parameterTypeField != nil {
		st.ParameterType = *parameterTypeField
	}
	positionField, err := identity(&pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = *positionField
	}
	typeIntervalTypeField, err := identity(&pb.TypeIntervalType)
	if err != nil {
		return nil, err
	}
	if typeIntervalTypeField != nil {
		st.TypeIntervalType = *typeIntervalTypeField
	}
	typeJsonField, err := identity(&pb.TypeJson)
	if err != nil {
		return nil, err
	}
	if typeJsonField != nil {
		st.TypeJson = *typeJsonField
	}
	typeNameField, err := identity(&pb.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	typePrecisionField, err := identity(&pb.TypePrecision)
	if err != nil {
		return nil, err
	}
	if typePrecisionField != nil {
		st.TypePrecision = *typePrecisionField
	}
	typeScaleField, err := identity(&pb.TypeScale)
	if err != nil {
		return nil, err
	}
	if typeScaleField != nil {
		st.TypeScale = *typeScaleField
	}
	typeTextField, err := identity(&pb.TypeText)
	if err != nil {
		return nil, err
	}
	if typeTextField != nil {
		st.TypeText = *typeTextField
	}

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
	ActivationUrl string
}

func getActivationUrlInfoRequestToPb(st *GetActivationUrlInfoRequest) (*getActivationUrlInfoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getActivationUrlInfoRequestPb{}
	activationUrlPb, err := identity(&st.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlPb != nil {
		pb.ActivationUrl = *activationUrlPb
	}

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
	activationUrlField, err := identity(&pb.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlField != nil {
		st.ActivationUrl = *activationUrlField
	}

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
	Name string
}

func getProviderRequestToPb(st *GetProviderRequest) (*getProviderRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getProviderRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// Get a share recipient
type GetRecipientRequest struct {
	// Name of the recipient.
	Name string
}

func getRecipientRequestToPb(st *GetRecipientRequest) (*getRecipientRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRecipientRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type GetRecipientSharePermissionsResponse struct {
	// Opaque token to retrieve the next page of results. Absent if there are no
	// more pages. __page_token__ should be set to this value for the next
	// request (for the next page of results).
	NextPageToken string
	// An array of data share permissions for a recipient.
	PermissionsOut []ShareToPrivilegeAssignment

	ForceSendFields []string
}

func getRecipientSharePermissionsResponseToPb(st *GetRecipientSharePermissionsResponse) (*getRecipientSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRecipientSharePermissionsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	NextPageToken string
	// The privileges assigned to each principal
	PrivilegeAssignments []PrivilegeAssignment

	ForceSendFields []string
}

func getSharePermissionsResponseToPb(st *GetSharePermissionsResponse) (*getSharePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSharePermissionsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	IncludeSharedData bool
	// The name of the share.
	Name string

	ForceSendFields []string
}

func getShareRequestToPb(st *GetShareRequest) (*getShareRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getShareRequestPb{}
	includeSharedDataPb, err := identity(&st.IncludeSharedData)
	if err != nil {
		return nil, err
	}
	if includeSharedDataPb != nil {
		pb.IncludeSharedData = *includeSharedDataPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	includeSharedDataField, err := identity(&pb.IncludeSharedData)
	if err != nil {
		return nil, err
	}
	if includeSharedDataField != nil {
		st.IncludeSharedData = *includeSharedDataField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	AllowedIpAddresses []string
}

func ipAccessListToPb(st *IpAccessList) (*ipAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ipAccessListPb{}

	var allowedIpAddressesPb []string
	for _, item := range st.AllowedIpAddresses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedIpAddressesPb = append(allowedIpAddressesPb, *itemPb)
		}
	}
	pb.AllowedIpAddresses = allowedIpAddressesPb

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

	var allowedIpAddressesField []string
	for _, itemPb := range pb.AllowedIpAddresses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedIpAddressesField = append(allowedIpAddressesField, *item)
		}
	}
	st.AllowedIpAddresses = allowedIpAddressesField

	return st, nil
}

// List assets by provider share
type ListProviderShareAssetsRequest struct {
	// Maximum number of functions to return.
	FunctionMaxResults int
	// Maximum number of notebooks to return.
	NotebookMaxResults int
	// The name of the provider who owns the share.
	ProviderName string
	// The name of the share.
	ShareName string
	// Maximum number of tables to return.
	TableMaxResults int
	// Maximum number of volumes to return.
	VolumeMaxResults int

	ForceSendFields []string
}

func listProviderShareAssetsRequestToPb(st *ListProviderShareAssetsRequest) (*listProviderShareAssetsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderShareAssetsRequestPb{}
	functionMaxResultsPb, err := identity(&st.FunctionMaxResults)
	if err != nil {
		return nil, err
	}
	if functionMaxResultsPb != nil {
		pb.FunctionMaxResults = *functionMaxResultsPb
	}

	notebookMaxResultsPb, err := identity(&st.NotebookMaxResults)
	if err != nil {
		return nil, err
	}
	if notebookMaxResultsPb != nil {
		pb.NotebookMaxResults = *notebookMaxResultsPb
	}

	providerNamePb, err := identity(&st.ProviderName)
	if err != nil {
		return nil, err
	}
	if providerNamePb != nil {
		pb.ProviderName = *providerNamePb
	}

	shareNamePb, err := identity(&st.ShareName)
	if err != nil {
		return nil, err
	}
	if shareNamePb != nil {
		pb.ShareName = *shareNamePb
	}

	tableMaxResultsPb, err := identity(&st.TableMaxResults)
	if err != nil {
		return nil, err
	}
	if tableMaxResultsPb != nil {
		pb.TableMaxResults = *tableMaxResultsPb
	}

	volumeMaxResultsPb, err := identity(&st.VolumeMaxResults)
	if err != nil {
		return nil, err
	}
	if volumeMaxResultsPb != nil {
		pb.VolumeMaxResults = *volumeMaxResultsPb
	}

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
	functionMaxResultsField, err := identity(&pb.FunctionMaxResults)
	if err != nil {
		return nil, err
	}
	if functionMaxResultsField != nil {
		st.FunctionMaxResults = *functionMaxResultsField
	}
	notebookMaxResultsField, err := identity(&pb.NotebookMaxResults)
	if err != nil {
		return nil, err
	}
	if notebookMaxResultsField != nil {
		st.NotebookMaxResults = *notebookMaxResultsField
	}
	providerNameField, err := identity(&pb.ProviderName)
	if err != nil {
		return nil, err
	}
	if providerNameField != nil {
		st.ProviderName = *providerNameField
	}
	shareNameField, err := identity(&pb.ShareName)
	if err != nil {
		return nil, err
	}
	if shareNameField != nil {
		st.ShareName = *shareNameField
	}
	tableMaxResultsField, err := identity(&pb.TableMaxResults)
	if err != nil {
		return nil, err
	}
	if tableMaxResultsField != nil {
		st.TableMaxResults = *tableMaxResultsField
	}
	volumeMaxResultsField, err := identity(&pb.VolumeMaxResults)
	if err != nil {
		return nil, err
	}
	if volumeMaxResultsField != nil {
		st.VolumeMaxResults = *volumeMaxResultsField
	}

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
	Functions []DeltaSharingFunction
	// The list of notebooks in the share.
	Notebooks []NotebookFile
	// The list of tables in the share.
	Tables []Table
	// The list of volumes in the share.
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
	NextPageToken string
	// An array of provider shares.
	Shares []ProviderShare

	ForceSendFields []string
}

func listProviderSharesResponseToPb(st *ListProviderSharesResponse) (*listProviderSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProviderSharesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	DataProviderGlobalMetastoreId string
	// Maximum number of providers to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid providers are returned (not
	// recommended). - Note: The number of returned providers might be less than
	// the specified max_results size, even zero. The only definitive indication
	// that no further providers can be fetched is when the next_page_token is
	// unset from the response.
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listProvidersRequestToPb(st *ListProvidersRequest) (*listProvidersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersRequestPb{}
	dataProviderGlobalMetastoreIdPb, err := identity(&st.DataProviderGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataProviderGlobalMetastoreIdPb != nil {
		pb.DataProviderGlobalMetastoreId = *dataProviderGlobalMetastoreIdPb
	}

	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	dataProviderGlobalMetastoreIdField, err := identity(&pb.DataProviderGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataProviderGlobalMetastoreIdField != nil {
		st.DataProviderGlobalMetastoreId = *dataProviderGlobalMetastoreIdField
	}
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// An array of provider information objects.
	Providers []ProviderInfo

	ForceSendFields []string
}

func listProvidersResponseToPb(st *ListProvidersResponse) (*listProvidersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listProvidersResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	DataRecipientGlobalMetastoreId string
	// Maximum number of recipients to return. - when set to 0, the page length
	// is set to a server configured value (recommended); - when set to a value
	// greater than 0, the page length is the minimum of this value and a server
	// configured value; - when set to a value less than 0, an invalid parameter
	// error is returned; - If not set, all valid recipients are returned (not
	// recommended). - Note: The number of returned recipients might be less
	// than the specified max_results size, even zero. The only definitive
	// indication that no further recipients can be fetched is when the
	// next_page_token is unset from the response.
	MaxResults int
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listRecipientsRequestToPb(st *ListRecipientsRequest) (*listRecipientsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRecipientsRequestPb{}
	dataRecipientGlobalMetastoreIdPb, err := identity(&st.DataRecipientGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataRecipientGlobalMetastoreIdPb != nil {
		pb.DataRecipientGlobalMetastoreId = *dataRecipientGlobalMetastoreIdPb
	}

	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	dataRecipientGlobalMetastoreIdField, err := identity(&pb.DataRecipientGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataRecipientGlobalMetastoreIdField != nil {
		st.DataRecipientGlobalMetastoreId = *dataRecipientGlobalMetastoreIdField
	}
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// An array of recipient information objects.
	Recipients []RecipientInfo

	ForceSendFields []string
}

func listRecipientsResponseToPb(st *ListRecipientsResponse) (*listRecipientsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRecipientsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	MaxResults int
	// Name of the provider in which to list shares.
	Name string
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listSharesRequestToPb(st *ListSharesRequest) (*listSharesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSharesRequestPb{}
	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// An array of data share information objects.
	Shares []ShareInfo

	ForceSendFields []string
}

func listSharesResponseToPb(st *ListSharesResponse) (*listSharesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSharesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	Comment string
	// The id of the notebook file.
	Id string
	// Name of the notebook file.
	Name string
	// The name of the share that the notebook file belongs to.
	Share string
	// The id of the share that the notebook file belongs to.
	ShareId string
	// The tags of the notebook file.
	Tags []catalog.TagKeyValue

	ForceSendFields []string
}

func notebookFileToPb(st *NotebookFile) (*notebookFilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookFilePb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	sharePb, err := identity(&st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = *sharePb
	}

	shareIdPb, err := identity(&st.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdPb != nil {
		pb.ShareId = *shareIdPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	shareField, err := identity(&pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = *shareField
	}
	shareIdField, err := identity(&pb.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdField != nil {
		st.ShareId = *shareIdField
	}

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
	Name string
	// The operator to apply for the value.
	Op PartitionValueOp
	// The key of a Delta Sharing recipient's property. For example
	// "databricks-account-id". When this field is set, field `value` can not be
	// set.
	RecipientPropertyKey string
	// The value of the partition column. When this value is not set, it means
	// `null` value. When this field is set, field `recipient_property_key` can
	// not be set.
	Value string

	ForceSendFields []string
}

func partitionValueToPb(st *PartitionValue) (*partitionValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partitionValuePb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	opPb, err := identity(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}

	recipientPropertyKeyPb, err := identity(&st.RecipientPropertyKey)
	if err != nil {
		return nil, err
	}
	if recipientPropertyKeyPb != nil {
		pb.RecipientPropertyKey = *recipientPropertyKeyPb
	}

	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	opField, err := identity(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	recipientPropertyKeyField, err := identity(&pb.RecipientPropertyKey)
	if err != nil {
		return nil, err
	}
	if recipientPropertyKeyField != nil {
		st.RecipientPropertyKey = *recipientPropertyKeyField
	}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

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
	Add []string
	// The principal whose privileges we are changing.
	Principal string
	// The set of privileges to remove.
	Remove []string

	ForceSendFields []string
}

func permissionsChangeToPb(st *PermissionsChange) (*permissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsChangePb{}

	var addPb []string
	for _, item := range st.Add {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addPb = append(addPb, *itemPb)
		}
	}
	pb.Add = addPb

	principalPb, err := identity(&st.Principal)
	if err != nil {
		return nil, err
	}
	if principalPb != nil {
		pb.Principal = *principalPb
	}

	var removePb []string
	for _, item := range st.Remove {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			removePb = append(removePb, *itemPb)
		}
	}
	pb.Remove = removePb

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

	var addField []string
	for _, itemPb := range pb.Add {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			addField = append(addField, *item)
		}
	}
	st.Add = addField
	principalField, err := identity(&pb.Principal)
	if err != nil {
		return nil, err
	}
	if principalField != nil {
		st.Principal = *principalField
	}

	var removeField []string
	for _, itemPb := range pb.Remove {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			removeField = append(removeField, *item)
		}
	}
	st.Remove = removeField

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
	Principal string
	// The privileges assigned to the principal.
	Privileges []Privilege

	ForceSendFields []string
}

func privilegeAssignmentToPb(st *PrivilegeAssignment) (*privilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privilegeAssignmentPb{}
	principalPb, err := identity(&st.Principal)
	if err != nil {
		return nil, err
	}
	if principalPb != nil {
		pb.Principal = *principalPb
	}

	var privilegesPb []Privilege
	for _, item := range st.Privileges {
		itemPb, err := identity(&item)
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
	principalField, err := identity(&pb.Principal)
	if err != nil {
		return nil, err
	}
	if principalField != nil {
		st.Principal = *principalField
	}

	var privilegesField []Privilege
	for _, itemPb := range pb.Privileges {
		item, err := identity(&itemPb)
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

func (st *privilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st privilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ProviderInfo struct {
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType
	// Cloud vendor of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Cloud string
	// Description about the provider.
	Comment string
	// Time at which this Provider was created, in epoch milliseconds.
	CreatedAt int64
	// Username of Provider creator.
	CreatedBy string
	// The global UC metastore id of the data provider. This field is only
	// present when the __authentication_type__ is **DATABRICKS**. The
	// identifier is of format __cloud__:__region__:__metastore-uuid__.
	DataProviderGlobalMetastoreId string
	// UUID of the provider's UC metastore. This field is only present when the
	// __authentication_type__ is **DATABRICKS**.
	MetastoreId string
	// The name of the Provider.
	Name string
	// Username of Provider owner.
	Owner string
	// The recipient profile. This field is only present when the
	// authentication_type is `TOKEN` or `OAUTH_CLIENT_CREDENTIALS`.
	RecipientProfile *RecipientProfile
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string
	// Cloud region of the provider's UC metastore. This field is only present
	// when the __authentication_type__ is **DATABRICKS**.
	Region string
	// Time at which this Provider was created, in epoch milliseconds.
	UpdatedAt int64
	// Username of user who last modified Provider.
	UpdatedBy string

	ForceSendFields []string
}

func providerInfoToPb(st *ProviderInfo) (*providerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerInfoPb{}
	authenticationTypePb, err := identity(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}

	cloudPb, err := identity(&st.Cloud)
	if err != nil {
		return nil, err
	}
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	dataProviderGlobalMetastoreIdPb, err := identity(&st.DataProviderGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataProviderGlobalMetastoreIdPb != nil {
		pb.DataProviderGlobalMetastoreId = *dataProviderGlobalMetastoreIdPb
	}

	metastoreIdPb, err := identity(&st.MetastoreId)
	if err != nil {
		return nil, err
	}
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	recipientProfilePb, err := recipientProfileToPb(st.RecipientProfile)
	if err != nil {
		return nil, err
	}
	if recipientProfilePb != nil {
		pb.RecipientProfile = recipientProfilePb
	}

	recipientProfileStrPb, err := identity(&st.RecipientProfileStr)
	if err != nil {
		return nil, err
	}
	if recipientProfileStrPb != nil {
		pb.RecipientProfileStr = *recipientProfileStrPb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	authenticationTypeField, err := identity(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	cloudField, err := identity(&pb.Cloud)
	if err != nil {
		return nil, err
	}
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	dataProviderGlobalMetastoreIdField, err := identity(&pb.DataProviderGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataProviderGlobalMetastoreIdField != nil {
		st.DataProviderGlobalMetastoreId = *dataProviderGlobalMetastoreIdField
	}
	metastoreIdField, err := identity(&pb.MetastoreId)
	if err != nil {
		return nil, err
	}
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	recipientProfileField, err := recipientProfileFromPb(pb.RecipientProfile)
	if err != nil {
		return nil, err
	}
	if recipientProfileField != nil {
		st.RecipientProfile = recipientProfileField
	}
	recipientProfileStrField, err := identity(&pb.RecipientProfileStr)
	if err != nil {
		return nil, err
	}
	if recipientProfileStrField != nil {
		st.RecipientProfileStr = *recipientProfileStrField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	Name string

	ForceSendFields []string
}

func providerShareToPb(st *ProviderShare) (*providerSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &providerSharePb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	Activated bool
	// Full activation url to retrieve the access token. It will be empty if the
	// token is already retrieved.
	ActivationUrl string
	// The delta sharing authentication type.
	AuthenticationType AuthenticationType
	// Cloud vendor of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Cloud string
	// Description about the recipient.
	Comment string
	// Time at which this recipient was created, in epoch milliseconds.
	CreatedAt int64
	// Username of recipient creator.
	CreatedBy string
	// The global Unity Catalog metastore id provided by the data recipient.
	// This field is only present when the __authentication_type__ is
	// **DATABRICKS**. The identifier is of format
	// __cloud__:__region__:__metastore-uuid__.
	DataRecipientGlobalMetastoreId string
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64
	// IP Access List
	IpAccessList *IpAccessList
	// Unique identifier of recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	MetastoreId string
	// Name of Recipient.
	Name string
	// Username of the recipient owner.
	Owner string
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs
	// Cloud region of the recipient's Unity Catalog Metastore. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	Region string
	// The one-time sharing code provided by the data recipient. This field is
	// only present when the __authentication_type__ is **DATABRICKS**.
	SharingCode string
	// This field is only present when the __authentication_type__ is **TOKEN**.
	Tokens []RecipientTokenInfo
	// Time at which the recipient was updated, in epoch milliseconds.
	UpdatedAt int64
	// Username of recipient updater.
	UpdatedBy string

	ForceSendFields []string
}

func recipientInfoToPb(st *RecipientInfo) (*recipientInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &recipientInfoPb{}
	activatedPb, err := identity(&st.Activated)
	if err != nil {
		return nil, err
	}
	if activatedPb != nil {
		pb.Activated = *activatedPb
	}

	activationUrlPb, err := identity(&st.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlPb != nil {
		pb.ActivationUrl = *activationUrlPb
	}

	authenticationTypePb, err := identity(&st.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypePb != nil {
		pb.AuthenticationType = *authenticationTypePb
	}

	cloudPb, err := identity(&st.Cloud)
	if err != nil {
		return nil, err
	}
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	dataRecipientGlobalMetastoreIdPb, err := identity(&st.DataRecipientGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataRecipientGlobalMetastoreIdPb != nil {
		pb.DataRecipientGlobalMetastoreId = *dataRecipientGlobalMetastoreIdPb
	}

	expirationTimePb, err := identity(&st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

	ipAccessListPb, err := ipAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	metastoreIdPb, err := identity(&st.MetastoreId)
	if err != nil {
		return nil, err
	}
	if metastoreIdPb != nil {
		pb.MetastoreId = *metastoreIdPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	propertiesKvpairsPb, err := securablePropertiesKvPairsToPb(st.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsPb != nil {
		pb.PropertiesKvpairs = propertiesKvpairsPb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	sharingCodePb, err := identity(&st.SharingCode)
	if err != nil {
		return nil, err
	}
	if sharingCodePb != nil {
		pb.SharingCode = *sharingCodePb
	}

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

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	activatedField, err := identity(&pb.Activated)
	if err != nil {
		return nil, err
	}
	if activatedField != nil {
		st.Activated = *activatedField
	}
	activationUrlField, err := identity(&pb.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlField != nil {
		st.ActivationUrl = *activationUrlField
	}
	authenticationTypeField, err := identity(&pb.AuthenticationType)
	if err != nil {
		return nil, err
	}
	if authenticationTypeField != nil {
		st.AuthenticationType = *authenticationTypeField
	}
	cloudField, err := identity(&pb.Cloud)
	if err != nil {
		return nil, err
	}
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	dataRecipientGlobalMetastoreIdField, err := identity(&pb.DataRecipientGlobalMetastoreId)
	if err != nil {
		return nil, err
	}
	if dataRecipientGlobalMetastoreIdField != nil {
		st.DataRecipientGlobalMetastoreId = *dataRecipientGlobalMetastoreIdField
	}
	expirationTimeField, err := identity(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
	ipAccessListField, err := ipAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	metastoreIdField, err := identity(&pb.MetastoreId)
	if err != nil {
		return nil, err
	}
	if metastoreIdField != nil {
		st.MetastoreId = *metastoreIdField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	propertiesKvpairsField, err := securablePropertiesKvPairsFromPb(pb.PropertiesKvpairs)
	if err != nil {
		return nil, err
	}
	if propertiesKvpairsField != nil {
		st.PropertiesKvpairs = propertiesKvpairsField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}
	sharingCodeField, err := identity(&pb.SharingCode)
	if err != nil {
		return nil, err
	}
	if sharingCodeField != nil {
		st.SharingCode = *sharingCodeField
	}

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
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	BearerToken string
	// The endpoint for the share to be used by the recipient.
	Endpoint string
	// The version number of the recipient's credentials on a share.
	ShareCredentialsVersion int

	ForceSendFields []string
}

func recipientProfileToPb(st *RecipientProfile) (*recipientProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &recipientProfilePb{}
	bearerTokenPb, err := identity(&st.BearerToken)
	if err != nil {
		return nil, err
	}
	if bearerTokenPb != nil {
		pb.BearerToken = *bearerTokenPb
	}

	endpointPb, err := identity(&st.Endpoint)
	if err != nil {
		return nil, err
	}
	if endpointPb != nil {
		pb.Endpoint = *endpointPb
	}

	shareCredentialsVersionPb, err := identity(&st.ShareCredentialsVersion)
	if err != nil {
		return nil, err
	}
	if shareCredentialsVersionPb != nil {
		pb.ShareCredentialsVersion = *shareCredentialsVersionPb
	}

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
	bearerTokenField, err := identity(&pb.BearerToken)
	if err != nil {
		return nil, err
	}
	if bearerTokenField != nil {
		st.BearerToken = *bearerTokenField
	}
	endpointField, err := identity(&pb.Endpoint)
	if err != nil {
		return nil, err
	}
	if endpointField != nil {
		st.Endpoint = *endpointField
	}
	shareCredentialsVersionField, err := identity(&pb.ShareCredentialsVersion)
	if err != nil {
		return nil, err
	}
	if shareCredentialsVersionField != nil {
		st.ShareCredentialsVersion = *shareCredentialsVersionField
	}

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
	ActivationUrl string
	// Time at which this recipient token was created, in epoch milliseconds.
	CreatedAt int64
	// Username of recipient token creator.
	CreatedBy string
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime int64
	// Unique ID of the recipient token.
	Id string
	// Time at which this recipient token was updated, in epoch milliseconds.
	UpdatedAt int64
	// Username of recipient token updater.
	UpdatedBy string

	ForceSendFields []string
}

func recipientTokenInfoToPb(st *RecipientTokenInfo) (*recipientTokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &recipientTokenInfoPb{}
	activationUrlPb, err := identity(&st.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlPb != nil {
		pb.ActivationUrl = *activationUrlPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	expirationTimePb, err := identity(&st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	activationUrlField, err := identity(&pb.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlField != nil {
		st.ActivationUrl = *activationUrlField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	expirationTimeField, err := identity(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	AliasName string
	// Numeric model version that alias will reference.
	VersionNum int64

	ForceSendFields []string
}

func registeredModelAliasToPb(st *RegisteredModelAlias) (*registeredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAliasPb{}
	aliasNamePb, err := identity(&st.AliasName)
	if err != nil {
		return nil, err
	}
	if aliasNamePb != nil {
		pb.AliasName = *aliasNamePb
	}

	versionNumPb, err := identity(&st.VersionNum)
	if err != nil {
		return nil, err
	}
	if versionNumPb != nil {
		pb.VersionNum = *versionNumPb
	}

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
	aliasNameField, err := identity(&pb.AliasName)
	if err != nil {
		return nil, err
	}
	if aliasNameField != nil {
		st.AliasName = *aliasNameField
	}
	versionNumField, err := identity(&pb.VersionNum)
	if err != nil {
		return nil, err
	}
	if versionNumField != nil {
		st.VersionNum = *versionNumField
	}

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
	ActivationUrl string
}

func retrieveTokenRequestToPb(st *RetrieveTokenRequest) (*retrieveTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &retrieveTokenRequestPb{}
	activationUrlPb, err := identity(&st.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlPb != nil {
		pb.ActivationUrl = *activationUrlPb
	}

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
	activationUrlField, err := identity(&pb.ActivationUrl)
	if err != nil {
		return nil, err
	}
	if activationUrlField != nil {
		st.ActivationUrl = *activationUrlField
	}

	return st, nil
}

type RetrieveTokenResponse struct {
	// The token used to authorize the recipient.
	BearerToken string
	// The endpoint for the share to be used by the recipient.
	Endpoint string
	// Expiration timestamp of the token in epoch milliseconds.
	ExpirationTime string
	// These field names must follow the delta sharing protocol.
	ShareCredentialsVersion int

	ForceSendFields []string
}

func retrieveTokenResponseToPb(st *RetrieveTokenResponse) (*retrieveTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &retrieveTokenResponsePb{}
	bearerTokenPb, err := identity(&st.BearerToken)
	if err != nil {
		return nil, err
	}
	if bearerTokenPb != nil {
		pb.BearerToken = *bearerTokenPb
	}

	endpointPb, err := identity(&st.Endpoint)
	if err != nil {
		return nil, err
	}
	if endpointPb != nil {
		pb.Endpoint = *endpointPb
	}

	expirationTimePb, err := identity(&st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

	shareCredentialsVersionPb, err := identity(&st.ShareCredentialsVersion)
	if err != nil {
		return nil, err
	}
	if shareCredentialsVersionPb != nil {
		pb.ShareCredentialsVersion = *shareCredentialsVersionPb
	}

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
	bearerTokenField, err := identity(&pb.BearerToken)
	if err != nil {
		return nil, err
	}
	if bearerTokenField != nil {
		st.BearerToken = *bearerTokenField
	}
	endpointField, err := identity(&pb.Endpoint)
	if err != nil {
		return nil, err
	}
	if endpointField != nil {
		st.Endpoint = *endpointField
	}
	expirationTimeField, err := identity(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
	shareCredentialsVersionField, err := identity(&pb.ShareCredentialsVersion)
	if err != nil {
		return nil, err
	}
	if shareCredentialsVersionField != nil {
		st.ShareCredentialsVersion = *shareCredentialsVersionField
	}

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
	ExistingTokenExpireInSeconds int64
	// The name of the Recipient.
	Name string
}

func rotateRecipientTokenToPb(st *RotateRecipientToken) (*rotateRecipientTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rotateRecipientTokenPb{}
	existingTokenExpireInSecondsPb, err := identity(&st.ExistingTokenExpireInSeconds)
	if err != nil {
		return nil, err
	}
	if existingTokenExpireInSecondsPb != nil {
		pb.ExistingTokenExpireInSeconds = *existingTokenExpireInSecondsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	existingTokenExpireInSecondsField, err := identity(&pb.ExistingTokenExpireInSeconds)
	if err != nil {
		return nil, err
	}
	if existingTokenExpireInSecondsField != nil {
		st.ExistingTokenExpireInSeconds = *existingTokenExpireInSecondsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

// An object with __properties__ containing map of key-value properties attached
// to the securable.
type SecurablePropertiesKvPairs struct {
	// A map of key-value properties attached to the securable.
	Properties map[string]string
}

func securablePropertiesKvPairsToPb(st *SecurablePropertiesKvPairs) (*securablePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &securablePropertiesKvPairsPb{}

	propertiesPb := map[string]string{}
	for k, v := range st.Properties {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			propertiesPb[k] = *itemPb
		}
	}
	pb.Properties = propertiesPb

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

	propertiesField := map[string]string{}
	for k, v := range pb.Properties {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			propertiesField[k] = *item
		}
	}
	st.Properties = propertiesField

	return st, nil
}

type ShareInfo struct {
	// User-provided free-form text description.
	Comment string
	// Time at which this share was created, in epoch milliseconds.
	CreatedAt int64
	// Username of share creator.
	CreatedBy string
	// Name of the share.
	Name string
	// A list of shared data objects within the share.
	Objects []SharedDataObject
	// Username of current owner of share.
	Owner string
	// Storage Location URL (full path) for the share.
	StorageLocation string
	// Storage root URL for the share.
	StorageRoot string
	// Time at which this share was updated, in epoch milliseconds.
	UpdatedAt int64
	// Username of share updater.
	UpdatedBy string

	ForceSendFields []string
}

func shareInfoToPb(st *ShareInfo) (*shareInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &shareInfoPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	storageLocationPb, err := identity(&st.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	storageRootPb, err := identity(&st.StorageRoot)
	if err != nil {
		return nil, err
	}
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	storageLocationField, err := identity(&pb.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	storageRootField, err := identity(&pb.StorageRoot)
	if err != nil {
		return nil, err
	}
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

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
	MaxResults int
	// The name of the Recipient.
	Name string
	// Opaque pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func sharePermissionsRequestToPb(st *SharePermissionsRequest) (*sharePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharePermissionsRequestPb{}
	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	PrivilegeAssignments []PrivilegeAssignment
	// The share name.
	ShareName string

	ForceSendFields []string
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

	shareNamePb, err := identity(&st.ShareName)
	if err != nil {
		return nil, err
	}
	if shareNamePb != nil {
		pb.ShareName = *shareNamePb
	}

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
	shareNameField, err := identity(&pb.ShareName)
	if err != nil {
		return nil, err
	}
	if shareNameField != nil {
		st.ShareName = *shareNameField
	}

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
	AddedAt int64
	// Username of the sharer.
	AddedBy string
	// Whether to enable cdf or indicate if cdf is enabled on the shared object.
	CdfEnabled bool
	// A user-provided comment when adding the data object to the share.
	Comment string
	// The content of the notebook file when the data object type is
	// NOTEBOOK_FILE. This should be base64 encoded. Required for adding a
	// NOTEBOOK_FILE, optional for updating, ignored for other types.
	Content string
	// The type of the data object.
	DataObjectType SharedDataObjectDataObjectType
	// Whether to enable or disable sharing of data history. If not specified,
	// the default is **DISABLED**.
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatus
	// A fully qualified name that uniquely identifies a data object. For
	// example, a table's fully qualified name is in the format of
	// `<catalog>.<schema>.<table>`,
	Name string
	// Array of partitions for the shared data.
	Partitions []Partition
	// A user-provided new name for the data object within the share. If this
	// new name is not provided, the object's original name will be used as the
	// `shared_as` name. The `shared_as` name must be unique within a share. For
	// tables, the new name must follow the format of `<schema>.<table>`.
	SharedAs string
	// The start version associated with the object. This allows data providers
	// to control the lowest object version that is accessible by clients. If
	// specified, clients can query snapshots or changes for versions >=
	// start_version. If not specified, clients can only query starting from the
	// version of the object at the time it was added to the share.
	//
	// NOTE: The start_version should be <= the `current` version of the object.
	StartVersion int64
	// One of: **ACTIVE**, **PERMISSION_DENIED**.
	Status SharedDataObjectStatus
	// A user-provided new name for the shared object within the share. If this
	// new name is not not provided, the object's original name will be used as
	// the `string_shared_as` name. The `string_shared_as` name must be unique
	// for objects of the same type within a Share. For notebooks, the new name
	// should be the new notebook file name.
	StringSharedAs string

	ForceSendFields []string
}

func sharedDataObjectToPb(st *SharedDataObject) (*sharedDataObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharedDataObjectPb{}
	addedAtPb, err := identity(&st.AddedAt)
	if err != nil {
		return nil, err
	}
	if addedAtPb != nil {
		pb.AddedAt = *addedAtPb
	}

	addedByPb, err := identity(&st.AddedBy)
	if err != nil {
		return nil, err
	}
	if addedByPb != nil {
		pb.AddedBy = *addedByPb
	}

	cdfEnabledPb, err := identity(&st.CdfEnabled)
	if err != nil {
		return nil, err
	}
	if cdfEnabledPb != nil {
		pb.CdfEnabled = *cdfEnabledPb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	dataObjectTypePb, err := identity(&st.DataObjectType)
	if err != nil {
		return nil, err
	}
	if dataObjectTypePb != nil {
		pb.DataObjectType = *dataObjectTypePb
	}

	historyDataSharingStatusPb, err := identity(&st.HistoryDataSharingStatus)
	if err != nil {
		return nil, err
	}
	if historyDataSharingStatusPb != nil {
		pb.HistoryDataSharingStatus = *historyDataSharingStatusPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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

	sharedAsPb, err := identity(&st.SharedAs)
	if err != nil {
		return nil, err
	}
	if sharedAsPb != nil {
		pb.SharedAs = *sharedAsPb
	}

	startVersionPb, err := identity(&st.StartVersion)
	if err != nil {
		return nil, err
	}
	if startVersionPb != nil {
		pb.StartVersion = *startVersionPb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	stringSharedAsPb, err := identity(&st.StringSharedAs)
	if err != nil {
		return nil, err
	}
	if stringSharedAsPb != nil {
		pb.StringSharedAs = *stringSharedAsPb
	}

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
	addedAtField, err := identity(&pb.AddedAt)
	if err != nil {
		return nil, err
	}
	if addedAtField != nil {
		st.AddedAt = *addedAtField
	}
	addedByField, err := identity(&pb.AddedBy)
	if err != nil {
		return nil, err
	}
	if addedByField != nil {
		st.AddedBy = *addedByField
	}
	cdfEnabledField, err := identity(&pb.CdfEnabled)
	if err != nil {
		return nil, err
	}
	if cdfEnabledField != nil {
		st.CdfEnabled = *cdfEnabledField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	dataObjectTypeField, err := identity(&pb.DataObjectType)
	if err != nil {
		return nil, err
	}
	if dataObjectTypeField != nil {
		st.DataObjectType = *dataObjectTypeField
	}
	historyDataSharingStatusField, err := identity(&pb.HistoryDataSharingStatus)
	if err != nil {
		return nil, err
	}
	if historyDataSharingStatusField != nil {
		st.HistoryDataSharingStatus = *historyDataSharingStatusField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	sharedAsField, err := identity(&pb.SharedAs)
	if err != nil {
		return nil, err
	}
	if sharedAsField != nil {
		st.SharedAs = *sharedAsField
	}
	startVersionField, err := identity(&pb.StartVersion)
	if err != nil {
		return nil, err
	}
	if startVersionField != nil {
		st.StartVersion = *startVersionField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	stringSharedAsField, err := identity(&pb.StringSharedAs)
	if err != nil {
		return nil, err
	}
	if stringSharedAsField != nil {
		st.StringSharedAs = *stringSharedAsField
	}

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
	Action SharedDataObjectUpdateAction
	// The data object that is being added, removed, or updated.
	DataObject *SharedDataObject
}

func sharedDataObjectUpdateToPb(st *SharedDataObjectUpdate) (*sharedDataObjectUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sharedDataObjectUpdatePb{}
	actionPb, err := identity(&st.Action)
	if err != nil {
		return nil, err
	}
	if actionPb != nil {
		pb.Action = *actionPb
	}

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
	actionField, err := identity(&pb.Action)
	if err != nil {
		return nil, err
	}
	if actionField != nil {
		st.Action = *actionField
	}
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
	Comment string
	// The id of the table.
	Id string
	// Internal information for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes *TableInternalAttributes
	// The name of a materialized table.
	MaterializedTableName string
	// The name of the table.
	Name string
	// The name of the schema that the table belongs to.
	Schema string
	// The name of the share that the table belongs to.
	Share string
	// The id of the share that the table belongs to.
	ShareId string
	// The Tags of the table.
	Tags []catalog.TagKeyValue

	ForceSendFields []string
}

func tableToPb(st *Table) (*tablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tablePb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	internalAttributesPb, err := tableInternalAttributesToPb(st.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesPb != nil {
		pb.InternalAttributes = internalAttributesPb
	}

	materializedTableNamePb, err := identity(&st.MaterializedTableName)
	if err != nil {
		return nil, err
	}
	if materializedTableNamePb != nil {
		pb.MaterializedTableName = *materializedTableNamePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	sharePb, err := identity(&st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = *sharePb
	}

	shareIdPb, err := identity(&st.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdPb != nil {
		pb.ShareId = *shareIdPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	internalAttributesField, err := tableInternalAttributesFromPb(pb.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesField != nil {
		st.InternalAttributes = internalAttributesField
	}
	materializedTableNameField, err := identity(&pb.MaterializedTableName)
	if err != nil {
		return nil, err
	}
	if materializedTableNameField != nil {
		st.MaterializedTableName = *materializedTableNameField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	shareField, err := identity(&pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = *shareField
	}
	shareIdField, err := identity(&pb.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdField != nil {
		st.ShareId = *shareIdField
	}

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
	ParentStorageLocation string
	// The cloud storage location of a shard table with DIRECTORY_BASED_TABLE
	// type.
	StorageLocation string
	// The type of the shared table.
	Type TableInternalAttributesSharedTableType
	// The view definition of a shared view. DEPRECATED.
	ViewDefinition string

	ForceSendFields []string
}

func tableInternalAttributesToPb(st *TableInternalAttributes) (*tableInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableInternalAttributesPb{}
	parentStorageLocationPb, err := identity(&st.ParentStorageLocation)
	if err != nil {
		return nil, err
	}
	if parentStorageLocationPb != nil {
		pb.ParentStorageLocation = *parentStorageLocationPb
	}

	storageLocationPb, err := identity(&st.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	viewDefinitionPb, err := identity(&st.ViewDefinition)
	if err != nil {
		return nil, err
	}
	if viewDefinitionPb != nil {
		pb.ViewDefinition = *viewDefinitionPb
	}

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
	parentStorageLocationField, err := identity(&pb.ParentStorageLocation)
	if err != nil {
		return nil, err
	}
	if parentStorageLocationField != nil {
		st.ParentStorageLocation = *parentStorageLocationField
	}
	storageLocationField, err := identity(&pb.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}
	viewDefinitionField, err := identity(&pb.ViewDefinition)
	if err != nil {
		return nil, err
	}
	if viewDefinitionField != nil {
		st.ViewDefinition = *viewDefinitionField
	}

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
	Comment string
	// Name of the provider.
	Name string
	// New name for the provider.
	NewName string
	// Username of Provider owner.
	Owner string
	// This field is required when the __authentication_type__ is **TOKEN**,
	// **OAUTH_CLIENT_CREDENTIALS** or not provided.
	RecipientProfileStr string

	ForceSendFields []string
}

func updateProviderToPb(st *UpdateProvider) (*updateProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProviderPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb, err := identity(&st.NewName)
	if err != nil {
		return nil, err
	}
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	recipientProfileStrPb, err := identity(&st.RecipientProfileStr)
	if err != nil {
		return nil, err
	}
	if recipientProfileStrPb != nil {
		pb.RecipientProfileStr = *recipientProfileStrPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField, err := identity(&pb.NewName)
	if err != nil {
		return nil, err
	}
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	recipientProfileStrField, err := identity(&pb.RecipientProfileStr)
	if err != nil {
		return nil, err
	}
	if recipientProfileStrField != nil {
		st.RecipientProfileStr = *recipientProfileStrField
	}

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
	Comment string
	// Expiration timestamp of the token, in epoch milliseconds.
	ExpirationTime int64
	// IP Access List
	IpAccessList *IpAccessList
	// Name of the recipient.
	Name string
	// New name for the recipient. .
	NewName string
	// Username of the recipient owner.
	Owner string
	// Recipient properties as map of string key-value pairs. When provided in
	// update request, the specified properties will override the existing
	// properties. To add and remove properties, one would need to perform a
	// read-modify-write.
	PropertiesKvpairs *SecurablePropertiesKvPairs

	ForceSendFields []string
}

func updateRecipientToPb(st *UpdateRecipient) (*updateRecipientPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRecipientPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	expirationTimePb, err := identity(&st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

	ipAccessListPb, err := ipAccessListToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb, err := identity(&st.NewName)
	if err != nil {
		return nil, err
	}
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	expirationTimeField, err := identity(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
	ipAccessListField, err := ipAccessListFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField, err := identity(&pb.NewName)
	if err != nil {
		return nil, err
	}
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
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
	Comment string
	// The name of the share.
	Name string
	// New name for the share.
	NewName string
	// Username of current owner of share.
	Owner string
	// Storage root URL for the share.
	StorageRoot string
	// Array of shared data object updates.
	Updates []SharedDataObjectUpdate

	ForceSendFields []string
}

func updateShareToPb(st *UpdateShare) (*updateSharePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSharePb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	newNamePb, err := identity(&st.NewName)
	if err != nil {
		return nil, err
	}
	if newNamePb != nil {
		pb.NewName = *newNamePb
	}

	ownerPb, err := identity(&st.Owner)
	if err != nil {
		return nil, err
	}
	if ownerPb != nil {
		pb.Owner = *ownerPb
	}

	storageRootPb, err := identity(&st.StorageRoot)
	if err != nil {
		return nil, err
	}
	if storageRootPb != nil {
		pb.StorageRoot = *storageRootPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	newNameField, err := identity(&pb.NewName)
	if err != nil {
		return nil, err
	}
	if newNameField != nil {
		st.NewName = *newNameField
	}
	ownerField, err := identity(&pb.Owner)
	if err != nil {
		return nil, err
	}
	if ownerField != nil {
		st.Owner = *ownerField
	}
	storageRootField, err := identity(&pb.StorageRoot)
	if err != nil {
		return nil, err
	}
	if storageRootField != nil {
		st.StorageRoot = *storageRootField
	}

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
	Changes []PermissionsChange
	// The name of the share.
	Name string
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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type UpdateSharePermissionsResponse struct {
	// The privileges assigned to each principal
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
	Comment string
	// This id maps to the shared_volume_id in database Recipient needs
	// shared_volume_id for recon to check if this volume is already in
	// recipient's DB or not.
	Id string
	// Internal attributes for D2D sharing that should not be disclosed to
	// external users.
	InternalAttributes *VolumeInternalAttributes
	// The name of the volume.
	Name string
	// The name of the schema that the volume belongs to.
	Schema string
	// The name of the share that the volume belongs to.
	Share string
	// / The id of the share that the volume belongs to.
	ShareId string
	// The tags of the volume.
	Tags []catalog.TagKeyValue

	ForceSendFields []string
}

func volumeToPb(st *Volume) (*volumePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumePb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	internalAttributesPb, err := volumeInternalAttributesToPb(st.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesPb != nil {
		pb.InternalAttributes = internalAttributesPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	sharePb, err := identity(&st.Share)
	if err != nil {
		return nil, err
	}
	if sharePb != nil {
		pb.Share = *sharePb
	}

	shareIdPb, err := identity(&st.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdPb != nil {
		pb.ShareId = *shareIdPb
	}

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
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	internalAttributesField, err := volumeInternalAttributesFromPb(pb.InternalAttributes)
	if err != nil {
		return nil, err
	}
	if internalAttributesField != nil {
		st.InternalAttributes = internalAttributesField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	shareField, err := identity(&pb.Share)
	if err != nil {
		return nil, err
	}
	if shareField != nil {
		st.Share = *shareField
	}
	shareIdField, err := identity(&pb.ShareId)
	if err != nil {
		return nil, err
	}
	if shareIdField != nil {
		st.ShareId = *shareIdField
	}

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
	StorageLocation string
	// The type of the shared volume.
	Type string

	ForceSendFields []string
}

func volumeInternalAttributesToPb(st *VolumeInternalAttributes) (*volumeInternalAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumeInternalAttributesPb{}
	storageLocationPb, err := identity(&st.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationPb != nil {
		pb.StorageLocation = *storageLocationPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

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
	storageLocationField, err := identity(&pb.StorageLocation)
	if err != nil {
		return nil, err
	}
	if storageLocationField != nil {
		st.StorageLocation = *storageLocationField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *volumeInternalAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st volumeInternalAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
