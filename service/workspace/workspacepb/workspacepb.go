// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspacepb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AclItemPb struct {
	Permission AclPermissionPb `json:"permission"`
	Principal  string          `json:"principal"`
}

type AclPermissionPb string

const AclPermissionManage AclPermissionPb = `MANAGE`

const AclPermissionRead AclPermissionPb = `READ`

const AclPermissionWrite AclPermissionPb = `WRITE`

type AzureKeyVaultSecretScopeMetadataPb struct {
	DnsName    string `json:"dns_name"`
	ResourceId string `json:"resource_id"`
}

type CreateCredentialsRequestPb struct {
	GitProvider          string `json:"git_provider"`
	GitUsername          string `json:"git_username,omitempty"`
	IsDefaultForProvider bool   `json:"is_default_for_provider,omitempty"`
	Name                 string `json:"name,omitempty"`
	PersonalAccessToken  string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCredentialsResponsePb struct {
	CredentialId         int64  `json:"credential_id"`
	GitProvider          string `json:"git_provider"`
	GitUsername          string `json:"git_username,omitempty"`
	IsDefaultForProvider bool   `json:"is_default_for_provider,omitempty"`
	Name                 string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRepoRequestPb struct {
	Path           string            `json:"path,omitempty"`
	Provider       string            `json:"provider"`
	SparseCheckout *SparseCheckoutPb `json:"sparse_checkout,omitempty"`
	Url            string            `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRepoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRepoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRepoResponsePb struct {
	Branch         string            `json:"branch,omitempty"`
	HeadCommitId   string            `json:"head_commit_id,omitempty"`
	Id             int64             `json:"id,omitempty"`
	Path           string            `json:"path,omitempty"`
	Provider       string            `json:"provider,omitempty"`
	SparseCheckout *SparseCheckoutPb `json:"sparse_checkout,omitempty"`
	Url            string            `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateRepoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateRepoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateScopePb struct {
	BackendAzureKeyvault   *AzureKeyVaultSecretScopeMetadataPb `json:"backend_azure_keyvault,omitempty"`
	InitialManagePrincipal string                              `json:"initial_manage_principal,omitempty"`
	Scope                  string                              `json:"scope"`
	ScopeBackendType       ScopeBackendTypePb                  `json:"scope_backend_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateScopePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateScopePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CredentialInfoPb struct {
	CredentialId         int64  `json:"credential_id"`
	GitProvider          string `json:"git_provider,omitempty"`
	GitUsername          string `json:"git_username,omitempty"`
	IsDefaultForProvider bool   `json:"is_default_for_provider,omitempty"`
	Name                 string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CredentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CredentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeletePb struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeletePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeletePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAclPb struct {
	Principal string `json:"principal"`
	Scope     string `json:"scope"`
}

type DeleteCredentialsRequestPb struct {
	CredentialId int64 `json:"-" url:"-"`
}

type DeleteCredentialsResponsePb struct {
}

type DeleteRepoRequestPb struct {
	RepoId int64 `json:"-" url:"-"`
}

type DeleteRepoResponsePb struct {
}

type DeleteResponsePb struct {
}

type DeleteScopePb struct {
	Scope string `json:"scope"`
}

type DeleteSecretPb struct {
	Key   string `json:"key"`
	Scope string `json:"scope"`
}

type DeleteSecretResponsePb struct {
}

type ExportFormatPb string

const ExportFormatAuto ExportFormatPb = `AUTO`

const ExportFormatDbc ExportFormatPb = `DBC`

const ExportFormatHtml ExportFormatPb = `HTML`

const ExportFormatJupyter ExportFormatPb = `JUPYTER`

const ExportFormatRaw ExportFormatPb = `RAW`

const ExportFormatRMarkdown ExportFormatPb = `R_MARKDOWN`

const ExportFormatSource ExportFormatPb = `SOURCE`

type ExportRequestPb struct {
	Format ExportFormatPb `json:"-" url:"format,omitempty"`
	Path   string         `json:"-" url:"path"`
}

type ExportResponsePb struct {
	Content  string `json:"content,omitempty"`
	FileType string `json:"file_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExportResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExportResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAclRequestPb struct {
	Principal string `json:"-" url:"principal"`
	Scope     string `json:"-" url:"scope"`
}

type GetCredentialsRequestPb struct {
	CredentialId int64 `json:"-" url:"-"`
}

type GetCredentialsResponsePb struct {
	CredentialId         int64  `json:"credential_id"`
	GitProvider          string `json:"git_provider,omitempty"`
	GitUsername          string `json:"git_username,omitempty"`
	IsDefaultForProvider bool   `json:"is_default_for_provider,omitempty"`
	Name                 string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetRepoPermissionLevelsRequestPb struct {
	RepoId string `json:"-" url:"-"`
}

type GetRepoPermissionLevelsResponsePb struct {
	PermissionLevels []RepoPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetRepoPermissionsRequestPb struct {
	RepoId string `json:"-" url:"-"`
}

type GetRepoRequestPb struct {
	RepoId int64 `json:"-" url:"-"`
}

type GetRepoResponsePb struct {
	Branch         string            `json:"branch,omitempty"`
	HeadCommitId   string            `json:"head_commit_id,omitempty"`
	Id             int64             `json:"id,omitempty"`
	Path           string            `json:"path,omitempty"`
	Provider       string            `json:"provider,omitempty"`
	SparseCheckout *SparseCheckoutPb `json:"sparse_checkout,omitempty"`
	Url            string            `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetRepoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetRepoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSecretRequestPb struct {
	Key   string `json:"-" url:"key"`
	Scope string `json:"-" url:"scope"`
}

type GetSecretResponsePb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetSecretResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetSecretResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetStatusRequestPb struct {
	Path string `json:"-" url:"path"`
}

type GetWorkspaceObjectPermissionLevelsRequestPb struct {
	WorkspaceObjectId   string `json:"-" url:"-"`
	WorkspaceObjectType string `json:"-" url:"-"`
}

type GetWorkspaceObjectPermissionLevelsResponsePb struct {
	PermissionLevels []WorkspaceObjectPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetWorkspaceObjectPermissionsRequestPb struct {
	WorkspaceObjectId   string `json:"-" url:"-"`
	WorkspaceObjectType string `json:"-" url:"-"`
}

type ImportPb struct {
	Content   string         `json:"content,omitempty"`
	Format    ImportFormatPb `json:"format,omitempty"`
	Language  LanguagePb     `json:"language,omitempty"`
	Overwrite bool           `json:"overwrite,omitempty"`
	Path      string         `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ImportPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ImportPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ImportFormatPb string

const ImportFormatAuto ImportFormatPb = `AUTO`

const ImportFormatDbc ImportFormatPb = `DBC`

const ImportFormatHtml ImportFormatPb = `HTML`

const ImportFormatJupyter ImportFormatPb = `JUPYTER`

const ImportFormatRaw ImportFormatPb = `RAW`

const ImportFormatRMarkdown ImportFormatPb = `R_MARKDOWN`

const ImportFormatSource ImportFormatPb = `SOURCE`

type ImportResponsePb struct {
}

type LanguagePb string

const LanguagePython LanguagePb = `PYTHON`

const LanguageR LanguagePb = `R`

const LanguageScala LanguagePb = `SCALA`

const LanguageSql LanguagePb = `SQL`

type ListAclsRequestPb struct {
	Scope string `json:"-" url:"scope"`
}

type ListAclsResponsePb struct {
	Items []AclItemPb `json:"items,omitempty"`
}

type ListCredentialsRequestPb struct {
}

type ListCredentialsResponsePb struct {
	Credentials []CredentialInfoPb `json:"credentials,omitempty"`
}

type ListReposRequestPb struct {
	NextPageToken string `json:"-" url:"next_page_token,omitempty"`
	PathPrefix    string `json:"-" url:"path_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListReposRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListReposRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListReposResponsePb struct {
	NextPageToken string       `json:"next_page_token,omitempty"`
	Repos         []RepoInfoPb `json:"repos,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListReposResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListReposResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListResponsePb struct {
	Objects []ObjectInfoPb `json:"objects,omitempty"`
}

type ListScopesRequestPb struct {
}

type ListScopesResponsePb struct {
	Scopes []SecretScopePb `json:"scopes,omitempty"`
}

type ListSecretsRequestPb struct {
	Scope string `json:"-" url:"scope"`
}

type ListSecretsResponsePb struct {
	Secrets []SecretMetadataPb `json:"secrets,omitempty"`
}

type ListWorkspaceRequestPb struct {
	NotebooksModifiedAfter int64  `json:"-" url:"notebooks_modified_after,omitempty"`
	Path                   string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MkdirsPb struct {
	Path string `json:"path"`
}

type MkdirsResponsePb struct {
}

type ObjectInfoPb struct {
	CreatedAt  int64        `json:"created_at,omitempty"`
	Language   LanguagePb   `json:"language,omitempty"`
	ModifiedAt int64        `json:"modified_at,omitempty"`
	ObjectId   int64        `json:"object_id,omitempty"`
	ObjectType ObjectTypePb `json:"object_type,omitempty"`
	Path       string       `json:"path,omitempty"`
	ResourceId string       `json:"resource_id,omitempty"`
	Size       int64        `json:"size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ObjectInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ObjectInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ObjectTypePb string

const ObjectTypeDashboard ObjectTypePb = `DASHBOARD`

const ObjectTypeDirectory ObjectTypePb = `DIRECTORY`

const ObjectTypeFile ObjectTypePb = `FILE`

const ObjectTypeLibrary ObjectTypePb = `LIBRARY`

const ObjectTypeNotebook ObjectTypePb = `NOTEBOOK`

const ObjectTypeRepo ObjectTypePb = `REPO`

type PutAclPb struct {
	Permission AclPermissionPb `json:"permission"`
	Principal  string          `json:"principal"`
	Scope      string          `json:"scope"`
}

type PutSecretPb struct {
	BytesValue  string `json:"bytes_value,omitempty"`
	Key         string `json:"key"`
	Scope       string `json:"scope"`
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoAccessControlRequestPb struct {
	GroupName            string                `json:"group_name,omitempty"`
	PermissionLevel      RepoPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                `json:"service_principal_name,omitempty"`
	UserName             string                `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepoAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepoAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoAccessControlResponsePb struct {
	AllPermissions       []RepoPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
	GroupName            string             `json:"group_name,omitempty"`
	ServicePrincipalName string             `json:"service_principal_name,omitempty"`
	UserName             string             `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepoAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepoAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoInfoPb struct {
	Branch         string            `json:"branch,omitempty"`
	HeadCommitId   string            `json:"head_commit_id,omitempty"`
	Id             int64             `json:"id,omitempty"`
	Path           string            `json:"path,omitempty"`
	Provider       string            `json:"provider,omitempty"`
	SparseCheckout *SparseCheckoutPb `json:"sparse_checkout,omitempty"`
	Url            string            `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepoInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepoInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermissionPb struct {
	Inherited           bool                  `json:"inherited,omitempty"`
	InheritedFromObject []string              `json:"inherited_from_object,omitempty"`
	PermissionLevel     RepoPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepoPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepoPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermissionLevelPb string

const RepoPermissionLevelCanEdit RepoPermissionLevelPb = `CAN_EDIT`

const RepoPermissionLevelCanManage RepoPermissionLevelPb = `CAN_MANAGE`

const RepoPermissionLevelCanRead RepoPermissionLevelPb = `CAN_READ`

const RepoPermissionLevelCanRun RepoPermissionLevelPb = `CAN_RUN`

type RepoPermissionsPb struct {
	AccessControlList []RepoAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                        `json:"object_id,omitempty"`
	ObjectType        string                        `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepoPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepoPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermissionsDescriptionPb struct {
	Description     string                `json:"description,omitempty"`
	PermissionLevel RepoPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepoPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepoPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermissionsRequestPb struct {
	AccessControlList []RepoAccessControlRequestPb `json:"access_control_list,omitempty"`
	RepoId            string                       `json:"-" url:"-"`
}

type ScopeBackendTypePb string

const ScopeBackendTypeAzureKeyvault ScopeBackendTypePb = `AZURE_KEYVAULT`

const ScopeBackendTypeDatabricks ScopeBackendTypePb = `DATABRICKS`

type SecretMetadataPb struct {
	Key                  string `json:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SecretMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SecretMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SecretScopePb struct {
	BackendType      ScopeBackendTypePb                  `json:"backend_type,omitempty"`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadataPb `json:"keyvault_metadata,omitempty"`
	Name             string                              `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SecretScopePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SecretScopePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparseCheckoutPb struct {
	Patterns []string `json:"patterns,omitempty"`
}

type SparseCheckoutUpdatePb struct {
	Patterns []string `json:"patterns,omitempty"`
}

type UpdateCredentialsRequestPb struct {
	CredentialId         int64  `json:"-" url:"-"`
	GitProvider          string `json:"git_provider"`
	GitUsername          string `json:"git_username,omitempty"`
	IsDefaultForProvider bool   `json:"is_default_for_provider,omitempty"`
	Name                 string `json:"name,omitempty"`
	PersonalAccessToken  string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCredentialsResponsePb struct {
}

type UpdateRepoRequestPb struct {
	Branch         string                  `json:"branch,omitempty"`
	RepoId         int64                   `json:"-" url:"-"`
	SparseCheckout *SparseCheckoutUpdatePb `json:"sparse_checkout,omitempty"`
	Tag            string                  `json:"tag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateRepoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateRepoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRepoResponsePb struct {
}

type WorkspaceObjectAccessControlRequestPb struct {
	GroupName            string                           `json:"group_name,omitempty"`
	PermissionLevel      WorkspaceObjectPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                           `json:"service_principal_name,omitempty"`
	UserName             string                           `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspaceObjectAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspaceObjectAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectAccessControlResponsePb struct {
	AllPermissions       []WorkspaceObjectPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                        `json:"display_name,omitempty"`
	GroupName            string                        `json:"group_name,omitempty"`
	ServicePrincipalName string                        `json:"service_principal_name,omitempty"`
	UserName             string                        `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspaceObjectAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspaceObjectAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermissionPb struct {
	Inherited           bool                             `json:"inherited,omitempty"`
	InheritedFromObject []string                         `json:"inherited_from_object,omitempty"`
	PermissionLevel     WorkspaceObjectPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspaceObjectPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspaceObjectPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermissionLevelPb string

const WorkspaceObjectPermissionLevelCanEdit WorkspaceObjectPermissionLevelPb = `CAN_EDIT`

const WorkspaceObjectPermissionLevelCanManage WorkspaceObjectPermissionLevelPb = `CAN_MANAGE`

const WorkspaceObjectPermissionLevelCanRead WorkspaceObjectPermissionLevelPb = `CAN_READ`

const WorkspaceObjectPermissionLevelCanRun WorkspaceObjectPermissionLevelPb = `CAN_RUN`

type WorkspaceObjectPermissionsPb struct {
	AccessControlList []WorkspaceObjectAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                                   `json:"object_id,omitempty"`
	ObjectType        string                                   `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspaceObjectPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspaceObjectPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermissionsDescriptionPb struct {
	Description     string                           `json:"description,omitempty"`
	PermissionLevel WorkspaceObjectPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspaceObjectPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspaceObjectPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermissionsRequestPb struct {
	AccessControlList   []WorkspaceObjectAccessControlRequestPb `json:"access_control_list,omitempty"`
	WorkspaceObjectId   string                                  `json:"-" url:"-"`
	WorkspaceObjectType string                                  `json:"-" url:"-"`
}
