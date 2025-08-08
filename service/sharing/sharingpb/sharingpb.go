// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sharingpb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/catalog/catalogpb"
)

type AuthenticationTypePb string

const AuthenticationTypeDatabricks AuthenticationTypePb = `DATABRICKS`

const AuthenticationTypeOauthClientCredentials AuthenticationTypePb = `OAUTH_CLIENT_CREDENTIALS`

const AuthenticationTypeOidcFederation AuthenticationTypePb = `OIDC_FEDERATION`

const AuthenticationTypeToken AuthenticationTypePb = `TOKEN`

type ColumnTypeNamePb string

const ColumnTypeNameArray ColumnTypeNamePb = `ARRAY`

const ColumnTypeNameBinary ColumnTypeNamePb = `BINARY`

const ColumnTypeNameBoolean ColumnTypeNamePb = `BOOLEAN`

const ColumnTypeNameByte ColumnTypeNamePb = `BYTE`

const ColumnTypeNameChar ColumnTypeNamePb = `CHAR`

const ColumnTypeNameDate ColumnTypeNamePb = `DATE`

const ColumnTypeNameDecimal ColumnTypeNamePb = `DECIMAL`

const ColumnTypeNameDouble ColumnTypeNamePb = `DOUBLE`

const ColumnTypeNameFloat ColumnTypeNamePb = `FLOAT`

const ColumnTypeNameInt ColumnTypeNamePb = `INT`

const ColumnTypeNameInterval ColumnTypeNamePb = `INTERVAL`

const ColumnTypeNameLong ColumnTypeNamePb = `LONG`

const ColumnTypeNameMap ColumnTypeNamePb = `MAP`

const ColumnTypeNameNull ColumnTypeNamePb = `NULL`

const ColumnTypeNameShort ColumnTypeNamePb = `SHORT`

const ColumnTypeNameString ColumnTypeNamePb = `STRING`

const ColumnTypeNameStruct ColumnTypeNamePb = `STRUCT`

const ColumnTypeNameTableType ColumnTypeNamePb = `TABLE_TYPE`

const ColumnTypeNameTimestamp ColumnTypeNamePb = `TIMESTAMP`

const ColumnTypeNameTimestampNtz ColumnTypeNamePb = `TIMESTAMP_NTZ`

const ColumnTypeNameUserDefinedType ColumnTypeNamePb = `USER_DEFINED_TYPE`

const ColumnTypeNameVariant ColumnTypeNamePb = `VARIANT`

type CreateFederationPolicyRequestPb struct {
	Policy        FederationPolicyPb `json:"policy"`
	RecipientName string             `json:"-" url:"-"`
}

type CreateProviderPb struct {
	AuthenticationType  AuthenticationTypePb `json:"authentication_type"`
	Comment             string               `json:"comment,omitempty"`
	Name                string               `json:"name"`
	RecipientProfileStr string               `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateProviderPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateProviderPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRecipientPb struct {
	AuthenticationType             AuthenticationTypePb          `json:"authentication_type"`
	Comment                        string                        `json:"comment,omitempty"`
	DataRecipientGlobalMetastoreId string                        `json:"data_recipient_global_metastore_id,omitempty"`
	ExpirationTime                 int64                         `json:"expiration_time,omitempty"`
	IpAccessList                   *IpAccessListPb               `json:"ip_access_list,omitempty"`
	Name                           string                        `json:"name"`
	Owner                          string                        `json:"owner,omitempty"`
	PropertiesKvpairs              *SecurablePropertiesKvPairsPb `json:"properties_kvpairs,omitempty"`
	SharingCode                    string                        `json:"sharing_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRecipientPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRecipientPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateSharePb struct {
	Comment     string `json:"comment,omitempty"`
	Name        string `json:"name"`
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteFederationPolicyRequestPb struct {
	Name          string `json:"-" url:"-"`
	RecipientName string `json:"-" url:"-"`
}

type DeleteProviderRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteRecipientRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeleteShareRequestPb struct {
	Name string `json:"-" url:"-"`
}

type DeltaSharingDependencyPb struct {
	Function *DeltaSharingFunctionDependencyPb `json:"function,omitempty"`
	Table    *DeltaSharingTableDependencyPb    `json:"table,omitempty"`
}

type DeltaSharingDependencyListPb struct {
	Dependencies []DeltaSharingDependencyPb `json:"dependencies,omitempty"`
}

type DeltaSharingFunctionPb struct {
	Aliases           []RegisteredModelAliasPb      `json:"aliases,omitempty"`
	Comment           string                        `json:"comment,omitempty"`
	DataType          ColumnTypeNamePb              `json:"data_type,omitempty"`
	DependencyList    *DeltaSharingDependencyListPb `json:"dependency_list,omitempty"`
	FullDataType      string                        `json:"full_data_type,omitempty"`
	Id                string                        `json:"id,omitempty"`
	InputParams       *FunctionParameterInfosPb     `json:"input_params,omitempty"`
	Name              string                        `json:"name,omitempty"`
	Properties        string                        `json:"properties,omitempty"`
	RoutineDefinition string                        `json:"routine_definition,omitempty"`
	Schema            string                        `json:"schema,omitempty"`
	SecurableKind     SharedSecurableKindPb         `json:"securable_kind,omitempty"`
	Share             string                        `json:"share,omitempty"`
	ShareId           string                        `json:"share_id,omitempty"`
	StorageLocation   string                        `json:"storage_location,omitempty"`
	Tags              []catalogpb.TagKeyValuePb     `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeltaSharingFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeltaSharingFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeltaSharingFunctionDependencyPb struct {
	FunctionName string `json:"function_name,omitempty"`
	SchemaName   string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeltaSharingFunctionDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeltaSharingFunctionDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeltaSharingTableDependencyPb struct {
	SchemaName string `json:"schema_name,omitempty"`
	TableName  string `json:"table_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeltaSharingTableDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeltaSharingTableDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FederationPolicyPb struct {
	Comment    string                  `json:"comment,omitempty"`
	CreateTime string                  `json:"create_time,omitempty"`
	Id         string                  `json:"id,omitempty"`
	Name       string                  `json:"name,omitempty"`
	OidcPolicy *OidcFederationPolicyPb `json:"oidc_policy,omitempty"`
	UpdateTime string                  `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FederationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FederationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FunctionParameterInfoPb struct {
	Comment          string                  `json:"comment,omitempty"`
	Name             string                  `json:"name,omitempty"`
	ParameterDefault string                  `json:"parameter_default,omitempty"`
	ParameterMode    FunctionParameterModePb `json:"parameter_mode,omitempty"`
	ParameterType    FunctionParameterTypePb `json:"parameter_type,omitempty"`
	Position         int                     `json:"position,omitempty"`
	TypeIntervalType string                  `json:"type_interval_type,omitempty"`
	TypeJson         string                  `json:"type_json,omitempty"`
	TypeName         ColumnTypeNamePb        `json:"type_name,omitempty"`
	TypePrecision    int                     `json:"type_precision,omitempty"`
	TypeScale        int                     `json:"type_scale,omitempty"`
	TypeText         string                  `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FunctionParameterInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FunctionParameterInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FunctionParameterInfosPb struct {
	Parameters []FunctionParameterInfoPb `json:"parameters,omitempty"`
}

type FunctionParameterModePb string

const FunctionParameterModeIn FunctionParameterModePb = `IN`

const FunctionParameterModeInout FunctionParameterModePb = `INOUT`

const FunctionParameterModeOut FunctionParameterModePb = `OUT`

type FunctionParameterTypePb string

const FunctionParameterTypeColumn FunctionParameterTypePb = `COLUMN`

const FunctionParameterTypeParam FunctionParameterTypePb = `PARAM`

type GetActivationUrlInfoRequestPb struct {
	ActivationUrl string `json:"-" url:"-"`
}

type GetActivationUrlInfoResponsePb struct {
}

type GetFederationPolicyRequestPb struct {
	Name          string `json:"-" url:"-"`
	RecipientName string `json:"-" url:"-"`
}

type GetProviderRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetRecipientRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetRecipientSharePermissionsResponsePb struct {
	NextPageToken  string                         `json:"next_page_token,omitempty"`
	PermissionsOut []ShareToPrivilegeAssignmentPb `json:"permissions_out,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetRecipientSharePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetRecipientSharePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSharePermissionsResponsePb struct {
	NextPageToken        string                  `json:"next_page_token,omitempty"`
	PrivilegeAssignments []PrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetSharePermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetSharePermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetShareRequestPb struct {
	IncludeSharedData bool   `json:"-" url:"include_shared_data,omitempty"`
	Name              string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetShareRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetShareRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type IpAccessListPb struct {
	AllowedIpAddresses []string `json:"allowed_ip_addresses,omitempty"`
}

type ListFederationPoliciesRequestPb struct {
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	RecipientName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFederationPoliciesResponsePb struct {
	NextPageToken string               `json:"next_page_token,omitempty"`
	Policies      []FederationPolicyPb `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFederationPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFederationPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProviderShareAssetsRequestPb struct {
	FunctionMaxResults int    `json:"-" url:"function_max_results,omitempty"`
	NotebookMaxResults int    `json:"-" url:"notebook_max_results,omitempty"`
	ProviderName       string `json:"-" url:"-"`
	ShareName          string `json:"-" url:"-"`
	TableMaxResults    int    `json:"-" url:"table_max_results,omitempty"`
	VolumeMaxResults   int    `json:"-" url:"volume_max_results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProviderShareAssetsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProviderShareAssetsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProviderShareAssetsResponsePb struct {
	Functions []DeltaSharingFunctionPb `json:"functions,omitempty"`
	Notebooks []NotebookFilePb         `json:"notebooks,omitempty"`
	Tables    []TablePb                `json:"tables,omitempty"`
	Volumes   []VolumePb               `json:"volumes,omitempty"`
}

type ListProviderSharesResponsePb struct {
	NextPageToken string            `json:"next_page_token,omitempty"`
	Shares        []ProviderSharePb `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProviderSharesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProviderSharesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProvidersRequestPb struct {
	DataProviderGlobalMetastoreId string `json:"-" url:"data_provider_global_metastore_id,omitempty"`
	MaxResults                    int    `json:"-" url:"max_results,omitempty"`
	PageToken                     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProvidersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProvidersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListProvidersResponsePb struct {
	NextPageToken string           `json:"next_page_token,omitempty"`
	Providers     []ProviderInfoPb `json:"providers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListProvidersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListProvidersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRecipientsRequestPb struct {
	DataRecipientGlobalMetastoreId string `json:"-" url:"data_recipient_global_metastore_id,omitempty"`
	MaxResults                     int    `json:"-" url:"max_results,omitempty"`
	PageToken                      string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRecipientsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRecipientsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRecipientsResponsePb struct {
	NextPageToken string            `json:"next_page_token,omitempty"`
	Recipients    []RecipientInfoPb `json:"recipients,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRecipientsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRecipientsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSharesRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	Name       string `json:"-" url:"-"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSharesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSharesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSharesResponsePb struct {
	NextPageToken string        `json:"next_page_token,omitempty"`
	Shares        []ShareInfoPb `json:"shares,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListSharesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListSharesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookFilePb struct {
	Comment string                    `json:"comment,omitempty"`
	Id      string                    `json:"id,omitempty"`
	Name    string                    `json:"name,omitempty"`
	Share   string                    `json:"share,omitempty"`
	ShareId string                    `json:"share_id,omitempty"`
	Tags    []catalogpb.TagKeyValuePb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NotebookFilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NotebookFilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OidcFederationPolicyPb struct {
	Audiences    []string `json:"audiences,omitempty"`
	Issuer       string   `json:"issuer"`
	Subject      string   `json:"subject"`
	SubjectClaim string   `json:"subject_claim"`
}

type PartitionPb struct {
	Values []PartitionValuePb `json:"values,omitempty"`
}

type PartitionValuePb struct {
	Name                 string             `json:"name,omitempty"`
	Op                   PartitionValueOpPb `json:"op,omitempty"`
	RecipientPropertyKey string             `json:"recipient_property_key,omitempty"`
	Value                string             `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PartitionValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PartitionValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PartitionValueOpPb string

const PartitionValueOpEqual PartitionValueOpPb = `EQUAL`

const PartitionValueOpLike PartitionValueOpPb = `LIKE`

type PermissionsChangePb struct {
	Add       []string `json:"add,omitempty"`
	Principal string   `json:"principal,omitempty"`
	Remove    []string `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PermissionsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PermissionsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PrivilegePb string

const PrivilegeAccess PrivilegePb = `ACCESS`

const PrivilegeAllPrivileges PrivilegePb = `ALL_PRIVILEGES`

const PrivilegeApplyTag PrivilegePb = `APPLY_TAG`

const PrivilegeCreate PrivilegePb = `CREATE`

const PrivilegeCreateCatalog PrivilegePb = `CREATE_CATALOG`

const PrivilegeCreateConnection PrivilegePb = `CREATE_CONNECTION`

const PrivilegeCreateExternalLocation PrivilegePb = `CREATE_EXTERNAL_LOCATION`

const PrivilegeCreateExternalTable PrivilegePb = `CREATE_EXTERNAL_TABLE`

const PrivilegeCreateExternalVolume PrivilegePb = `CREATE_EXTERNAL_VOLUME`

const PrivilegeCreateForeignCatalog PrivilegePb = `CREATE_FOREIGN_CATALOG`

const PrivilegeCreateForeignSecurable PrivilegePb = `CREATE_FOREIGN_SECURABLE`

const PrivilegeCreateFunction PrivilegePb = `CREATE_FUNCTION`

const PrivilegeCreateManagedStorage PrivilegePb = `CREATE_MANAGED_STORAGE`

const PrivilegeCreateMaterializedView PrivilegePb = `CREATE_MATERIALIZED_VIEW`

const PrivilegeCreateModel PrivilegePb = `CREATE_MODEL`

const PrivilegeCreateProvider PrivilegePb = `CREATE_PROVIDER`

const PrivilegeCreateRecipient PrivilegePb = `CREATE_RECIPIENT`

const PrivilegeCreateSchema PrivilegePb = `CREATE_SCHEMA`

const PrivilegeCreateServiceCredential PrivilegePb = `CREATE_SERVICE_CREDENTIAL`

const PrivilegeCreateShare PrivilegePb = `CREATE_SHARE`

const PrivilegeCreateStorageCredential PrivilegePb = `CREATE_STORAGE_CREDENTIAL`

const PrivilegeCreateTable PrivilegePb = `CREATE_TABLE`

const PrivilegeCreateView PrivilegePb = `CREATE_VIEW`

const PrivilegeCreateVolume PrivilegePb = `CREATE_VOLUME`

const PrivilegeExecute PrivilegePb = `EXECUTE`

const PrivilegeManage PrivilegePb = `MANAGE`

const PrivilegeManageAllowlist PrivilegePb = `MANAGE_ALLOWLIST`

const PrivilegeModify PrivilegePb = `MODIFY`

const PrivilegeReadFiles PrivilegePb = `READ_FILES`

const PrivilegeReadPrivateFiles PrivilegePb = `READ_PRIVATE_FILES`

const PrivilegeReadVolume PrivilegePb = `READ_VOLUME`

const PrivilegeRefresh PrivilegePb = `REFRESH`

const PrivilegeSelect PrivilegePb = `SELECT`

const PrivilegeSetSharePermission PrivilegePb = `SET_SHARE_PERMISSION`

const PrivilegeUsage PrivilegePb = `USAGE`

const PrivilegeUseCatalog PrivilegePb = `USE_CATALOG`

const PrivilegeUseConnection PrivilegePb = `USE_CONNECTION`

const PrivilegeUseMarketplaceAssets PrivilegePb = `USE_MARKETPLACE_ASSETS`

const PrivilegeUseProvider PrivilegePb = `USE_PROVIDER`

const PrivilegeUseRecipient PrivilegePb = `USE_RECIPIENT`

const PrivilegeUseSchema PrivilegePb = `USE_SCHEMA`

const PrivilegeUseShare PrivilegePb = `USE_SHARE`

const PrivilegeWriteFiles PrivilegePb = `WRITE_FILES`

const PrivilegeWritePrivateFiles PrivilegePb = `WRITE_PRIVATE_FILES`

const PrivilegeWriteVolume PrivilegePb = `WRITE_VOLUME`

type PrivilegeAssignmentPb struct {
	Principal  string        `json:"principal,omitempty"`
	Privileges []PrivilegePb `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ProviderInfoPb struct {
	AuthenticationType            AuthenticationTypePb `json:"authentication_type,omitempty"`
	Cloud                         string               `json:"cloud,omitempty"`
	Comment                       string               `json:"comment,omitempty"`
	CreatedAt                     int64                `json:"created_at,omitempty"`
	CreatedBy                     string               `json:"created_by,omitempty"`
	DataProviderGlobalMetastoreId string               `json:"data_provider_global_metastore_id,omitempty"`
	MetastoreId                   string               `json:"metastore_id,omitempty"`
	Name                          string               `json:"name,omitempty"`
	Owner                         string               `json:"owner,omitempty"`
	RecipientProfile              *RecipientProfilePb  `json:"recipient_profile,omitempty"`
	RecipientProfileStr           string               `json:"recipient_profile_str,omitempty"`
	Region                        string               `json:"region,omitempty"`
	UpdatedAt                     int64                `json:"updated_at,omitempty"`
	UpdatedBy                     string               `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ProviderInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ProviderInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ProviderSharePb struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ProviderSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ProviderSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RecipientInfoPb struct {
	Activated                      bool                          `json:"activated,omitempty"`
	ActivationUrl                  string                        `json:"activation_url,omitempty"`
	AuthenticationType             AuthenticationTypePb          `json:"authentication_type,omitempty"`
	Cloud                          string                        `json:"cloud,omitempty"`
	Comment                        string                        `json:"comment,omitempty"`
	CreatedAt                      int64                         `json:"created_at,omitempty"`
	CreatedBy                      string                        `json:"created_by,omitempty"`
	DataRecipientGlobalMetastoreId string                        `json:"data_recipient_global_metastore_id,omitempty"`
	ExpirationTime                 int64                         `json:"expiration_time,omitempty"`
	IpAccessList                   *IpAccessListPb               `json:"ip_access_list,omitempty"`
	MetastoreId                    string                        `json:"metastore_id,omitempty"`
	Name                           string                        `json:"name,omitempty"`
	Owner                          string                        `json:"owner,omitempty"`
	PropertiesKvpairs              *SecurablePropertiesKvPairsPb `json:"properties_kvpairs,omitempty"`
	Region                         string                        `json:"region,omitempty"`
	SharingCode                    string                        `json:"sharing_code,omitempty"`
	Tokens                         []RecipientTokenInfoPb        `json:"tokens,omitempty"`
	UpdatedAt                      int64                         `json:"updated_at,omitempty"`
	UpdatedBy                      string                        `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RecipientInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RecipientInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RecipientProfilePb struct {
	BearerToken             string `json:"bearer_token,omitempty"`
	Endpoint                string `json:"endpoint,omitempty"`
	ShareCredentialsVersion int    `json:"share_credentials_version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RecipientProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RecipientProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RecipientTokenInfoPb struct {
	ActivationUrl  string `json:"activation_url,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	CreatedBy      string `json:"created_by,omitempty"`
	ExpirationTime int64  `json:"expiration_time,omitempty"`
	Id             string `json:"id,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	UpdatedBy      string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RecipientTokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RecipientTokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelAliasPb struct {
	AliasName  string `json:"alias_name,omitempty"`
	VersionNum int64  `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RegisteredModelAliasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RegisteredModelAliasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RetrieveTokenRequestPb struct {
	ActivationUrl string `json:"-" url:"-"`
}

type RetrieveTokenResponsePb struct {
	BearerToken             string `json:"bearerToken,omitempty"`
	Endpoint                string `json:"endpoint,omitempty"`
	ExpirationTime          string `json:"expirationTime,omitempty"`
	ShareCredentialsVersion int    `json:"shareCredentialsVersion,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RetrieveTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RetrieveTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RotateRecipientTokenPb struct {
	ExistingTokenExpireInSeconds int64  `json:"existing_token_expire_in_seconds"`
	Name                         string `json:"-" url:"-"`
}

type SecurablePropertiesKvPairsPb struct {
	Properties map[string]string `json:"properties"`
}

type ShareInfoPb struct {
	Comment         string               `json:"comment,omitempty"`
	CreatedAt       int64                `json:"created_at,omitempty"`
	CreatedBy       string               `json:"created_by,omitempty"`
	Name            string               `json:"name,omitempty"`
	Objects         []SharedDataObjectPb `json:"objects,omitempty"`
	Owner           string               `json:"owner,omitempty"`
	StorageLocation string               `json:"storage_location,omitempty"`
	StorageRoot     string               `json:"storage_root,omitempty"`
	UpdatedAt       int64                `json:"updated_at,omitempty"`
	UpdatedBy       string               `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ShareInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ShareInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SharePermissionsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	Name       string `json:"-" url:"-"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SharePermissionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SharePermissionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ShareToPrivilegeAssignmentPb struct {
	PrivilegeAssignments []PrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`
	ShareName            string                  `json:"share_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ShareToPrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ShareToPrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SharedDataObjectPb struct {
	AddedAt                  int64                                      `json:"added_at,omitempty"`
	AddedBy                  string                                     `json:"added_by,omitempty"`
	CdfEnabled               bool                                       `json:"cdf_enabled,omitempty"`
	Comment                  string                                     `json:"comment,omitempty"`
	Content                  string                                     `json:"content,omitempty"`
	DataObjectType           SharedDataObjectDataObjectTypePb           `json:"data_object_type,omitempty"`
	HistoryDataSharingStatus SharedDataObjectHistoryDataSharingStatusPb `json:"history_data_sharing_status,omitempty"`
	Name                     string                                     `json:"name"`
	Partitions               []PartitionPb                              `json:"partitions,omitempty"`
	SharedAs                 string                                     `json:"shared_as,omitempty"`
	StartVersion             int64                                      `json:"start_version,omitempty"`
	Status                   SharedDataObjectStatusPb                   `json:"status,omitempty"`
	StringSharedAs           string                                     `json:"string_shared_as,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SharedDataObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SharedDataObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SharedDataObjectDataObjectTypePb string

const SharedDataObjectDataObjectTypeFeatureSpec SharedDataObjectDataObjectTypePb = `FEATURE_SPEC`

const SharedDataObjectDataObjectTypeFunction SharedDataObjectDataObjectTypePb = `FUNCTION`

const SharedDataObjectDataObjectTypeMaterializedView SharedDataObjectDataObjectTypePb = `MATERIALIZED_VIEW`

const SharedDataObjectDataObjectTypeModel SharedDataObjectDataObjectTypePb = `MODEL`

const SharedDataObjectDataObjectTypeNotebookFile SharedDataObjectDataObjectTypePb = `NOTEBOOK_FILE`

const SharedDataObjectDataObjectTypeSchema SharedDataObjectDataObjectTypePb = `SCHEMA`

const SharedDataObjectDataObjectTypeStreamingTable SharedDataObjectDataObjectTypePb = `STREAMING_TABLE`

const SharedDataObjectDataObjectTypeTable SharedDataObjectDataObjectTypePb = `TABLE`

const SharedDataObjectDataObjectTypeView SharedDataObjectDataObjectTypePb = `VIEW`

type SharedDataObjectHistoryDataSharingStatusPb string

const SharedDataObjectHistoryDataSharingStatusDisabled SharedDataObjectHistoryDataSharingStatusPb = `DISABLED`

const SharedDataObjectHistoryDataSharingStatusEnabled SharedDataObjectHistoryDataSharingStatusPb = `ENABLED`

type SharedDataObjectStatusPb string

const SharedDataObjectStatusActive SharedDataObjectStatusPb = `ACTIVE`

const SharedDataObjectStatusPermissionDenied SharedDataObjectStatusPb = `PERMISSION_DENIED`

type SharedDataObjectUpdatePb struct {
	Action     SharedDataObjectUpdateActionPb `json:"action,omitempty"`
	DataObject *SharedDataObjectPb            `json:"data_object,omitempty"`
}

type SharedDataObjectUpdateActionPb string

const SharedDataObjectUpdateActionAdd SharedDataObjectUpdateActionPb = `ADD`

const SharedDataObjectUpdateActionRemove SharedDataObjectUpdateActionPb = `REMOVE`

const SharedDataObjectUpdateActionUpdate SharedDataObjectUpdateActionPb = `UPDATE`

type SharedSecurableKindPb string

const SharedSecurableKindFunctionFeatureSpec SharedSecurableKindPb = `FUNCTION_FEATURE_SPEC`

const SharedSecurableKindFunctionRegisteredModel SharedSecurableKindPb = `FUNCTION_REGISTERED_MODEL`

const SharedSecurableKindFunctionStandard SharedSecurableKindPb = `FUNCTION_STANDARD`

type TablePb struct {
	Comment                  string                     `json:"comment,omitempty"`
	Id                       string                     `json:"id,omitempty"`
	InternalAttributes       *TableInternalAttributesPb `json:"internal_attributes,omitempty"`
	MaterializationNamespace string                     `json:"materialization_namespace,omitempty"`
	MaterializedTableName    string                     `json:"materialized_table_name,omitempty"`
	Name                     string                     `json:"name,omitempty"`
	Schema                   string                     `json:"schema,omitempty"`
	Share                    string                     `json:"share,omitempty"`
	ShareId                  string                     `json:"share_id,omitempty"`
	Tags                     []catalogpb.TagKeyValuePb  `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableInternalAttributesPb struct {
	ParentStorageLocation string                                   `json:"parent_storage_location,omitempty"`
	StorageLocation       string                                   `json:"storage_location,omitempty"`
	Type                  TableInternalAttributesSharedTableTypePb `json:"type,omitempty"`
	ViewDefinition        string                                   `json:"view_definition,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableInternalAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableInternalAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableInternalAttributesSharedTableTypePb string

const TableInternalAttributesSharedTableTypeDeltaIcebergTable TableInternalAttributesSharedTableTypePb = `DELTA_ICEBERG_TABLE`

const TableInternalAttributesSharedTableTypeDirectoryBasedTable TableInternalAttributesSharedTableTypePb = `DIRECTORY_BASED_TABLE`

const TableInternalAttributesSharedTableTypeFileBasedTable TableInternalAttributesSharedTableTypePb = `FILE_BASED_TABLE`

const TableInternalAttributesSharedTableTypeForeignTable TableInternalAttributesSharedTableTypePb = `FOREIGN_TABLE`

const TableInternalAttributesSharedTableTypeMaterializedView TableInternalAttributesSharedTableTypePb = `MATERIALIZED_VIEW`

const TableInternalAttributesSharedTableTypeStreamingTable TableInternalAttributesSharedTableTypePb = `STREAMING_TABLE`

const TableInternalAttributesSharedTableTypeView TableInternalAttributesSharedTableTypePb = `VIEW`

type UpdateFederationPolicyRequestPb struct {
	Name          string             `json:"-" url:"-"`
	Policy        FederationPolicyPb `json:"policy"`
	RecipientName string             `json:"-" url:"-"`
	UpdateMask    string             `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateProviderPb struct {
	Comment             string `json:"comment,omitempty"`
	Name                string `json:"-" url:"-"`
	NewName             string `json:"new_name,omitempty"`
	Owner               string `json:"owner,omitempty"`
	RecipientProfileStr string `json:"recipient_profile_str,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateProviderPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateProviderPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRecipientPb struct {
	Comment           string                        `json:"comment,omitempty"`
	ExpirationTime    int64                         `json:"expiration_time,omitempty"`
	IpAccessList      *IpAccessListPb               `json:"ip_access_list,omitempty"`
	Name              string                        `json:"-" url:"-"`
	NewName           string                        `json:"new_name,omitempty"`
	Owner             string                        `json:"owner,omitempty"`
	PropertiesKvpairs *SecurablePropertiesKvPairsPb `json:"properties_kvpairs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateRecipientPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateRecipientPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateSharePb struct {
	Comment     string                     `json:"comment,omitempty"`
	Name        string                     `json:"-" url:"-"`
	NewName     string                     `json:"new_name,omitempty"`
	Owner       string                     `json:"owner,omitempty"`
	StorageRoot string                     `json:"storage_root,omitempty"`
	Updates     []SharedDataObjectUpdatePb `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateSharePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateSharePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateSharePermissionsPb struct {
	Changes             []PermissionsChangePb `json:"changes,omitempty"`
	Name                string                `json:"-" url:"-"`
	OmitPermissionsList bool                  `json:"omit_permissions_list,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateSharePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateSharePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateSharePermissionsResponsePb struct {
	PrivilegeAssignments []PrivilegeAssignmentPb `json:"privilege_assignments,omitempty"`
}

type VolumePb struct {
	Comment            string                      `json:"comment,omitempty"`
	Id                 string                      `json:"id,omitempty"`
	InternalAttributes *VolumeInternalAttributesPb `json:"internal_attributes,omitempty"`
	Name               string                      `json:"name,omitempty"`
	Schema             string                      `json:"schema,omitempty"`
	Share              string                      `json:"share,omitempty"`
	ShareId            string                      `json:"share_id,omitempty"`
	Tags               []catalogpb.TagKeyValuePb   `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *VolumePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VolumePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VolumeInternalAttributesPb struct {
	StorageLocation string `json:"storage_location,omitempty"`
	Type            string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *VolumeInternalAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VolumeInternalAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
