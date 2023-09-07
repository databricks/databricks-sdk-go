// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"fmt"

	marshal "github.com/databricks/databricks-sdk-go/json"
)

// all definitions in this file are in alphabetical order

type AclItem struct {
	// The permission level applied to the principal.
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	Principal string `json:"principal"`

	ForceSendFields []string `json:"-"`
}

func (s *AclItem) UnmarshalJSON(b []byte) error {
	type C AclItem
	return marshal.Unmarshal(b, (*C)(s))
}

func (s AclItem) MarshalJSON() ([]byte, error) {
	type C AclItem
	return marshal.Marshal((C)(s))
}

type AclPermission string

const AclPermissionManage AclPermission = `MANAGE`

const AclPermissionRead AclPermission = `READ`

const AclPermissionWrite AclPermission = `WRITE`

// String representation for [fmt.Print]
func (f *AclPermission) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AclPermission) Set(v string) error {
	switch v {
	case `MANAGE`, `READ`, `WRITE`:
		*f = AclPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGE", "READ", "WRITE"`, v)
	}
}

// Type always returns AclPermission to satisfy [pflag.Value] interface
func (f *AclPermission) Type() string {
	return "AclPermission"
}

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName string `json:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId string `json:"resource_id"`

	ForceSendFields []string `json:"-"`
}

func (s *AzureKeyVaultSecretScopeMetadata) UnmarshalJSON(b []byte) error {
	type C AzureKeyVaultSecretScopeMetadata
	return marshal.Unmarshal(b, (*C)(s))
}

func (s AzureKeyVaultSecretScopeMetadata) MarshalJSON() ([]byte, error) {
	type C AzureKeyVaultSecretScopeMetadata
	return marshal.Marshal((C)(s))
}

type CreateCredentials struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `json:"git_provider"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider.
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateCredentials) UnmarshalJSON(b []byte) error {
	type C CreateCredentials
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateCredentials) MarshalJSON() ([]byte, error) {
	type C CreateCredentials
	return marshal.Marshal((C)(s))
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateCredentialsResponse) UnmarshalJSON(b []byte) error {
	type C CreateCredentialsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateCredentialsResponse) MarshalJSON() ([]byte, error) {
	type C CreateCredentialsResponse
	return marshal.Marshal((C)(s))
}

type CreateRepo struct {
	// Desired path for the repo in the workspace. Must be in the format
	// /Repos/{folder}/{repo-name}.
	Path string `json:"path,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider string `json:"provider"`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the Git repository to be linked.
	Url string `json:"url"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateRepo) UnmarshalJSON(b []byte) error {
	type C CreateRepo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateRepo) MarshalJSON() ([]byte, error) {
	type C CreateRepo
	return marshal.Marshal((C)(s))
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata `json:"backend_azure_keyvault,omitempty"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal string `json:"initial_manage_principal,omitempty"`
	// Scope name requested by the user. Scope names are unique.
	Scope string `json:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType ScopeBackendType `json:"scope_backend_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateScope) UnmarshalJSON(b []byte) error {
	type C CreateScope
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CreateScope) MarshalJSON() ([]byte, error) {
	type C CreateScope
	return marshal.Marshal((C)(s))
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, gitHubOAuth, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	type C CredentialInfo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	type C CredentialInfo
	return marshal.Marshal((C)(s))
}

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Delete) UnmarshalJSON(b []byte) error {
	type C Delete
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Delete) MarshalJSON() ([]byte, error) {
	type C Delete
	return marshal.Marshal((C)(s))
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal string `json:"principal"`
	// The name of the scope to remove permissions from.
	Scope string `json:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAcl) UnmarshalJSON(b []byte) error {
	type C DeleteAcl
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteAcl) MarshalJSON() ([]byte, error) {
	type C DeleteAcl
	return marshal.Marshal((C)(s))
}

// Delete a credential
type DeleteGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteGitCredentialRequest) UnmarshalJSON(b []byte) error {
	type C DeleteGitCredentialRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteGitCredentialRequest) MarshalJSON() ([]byte, error) {
	type C DeleteGitCredentialRequest
	return marshal.Marshal((C)(s))
}

// Delete a repo
type DeleteRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteRepoRequest) UnmarshalJSON(b []byte) error {
	type C DeleteRepoRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteRepoRequest) MarshalJSON() ([]byte, error) {
	type C DeleteRepoRequest
	return marshal.Marshal((C)(s))
}

type DeleteScope struct {
	// Name of the scope to delete.
	Scope string `json:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteScope) UnmarshalJSON(b []byte) error {
	type C DeleteScope
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteScope) MarshalJSON() ([]byte, error) {
	type C DeleteScope
	return marshal.Marshal((C)(s))
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key string `json:"key"`
	// The name of the scope that contains the secret to delete.
	Scope string `json:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteSecret) UnmarshalJSON(b []byte) error {
	type C DeleteSecret
	return marshal.Unmarshal(b, (*C)(s))
}

func (s DeleteSecret) MarshalJSON() ([]byte, error) {
	type C DeleteSecret
	return marshal.Marshal((C)(s))
}

type ExportFormat string

const ExportFormatDbc ExportFormat = `DBC`

const ExportFormatHtml ExportFormat = `HTML`

const ExportFormatJupyter ExportFormat = `JUPYTER`

const ExportFormatRMarkdown ExportFormat = `R_MARKDOWN`

const ExportFormatSource ExportFormat = `SOURCE`

// String representation for [fmt.Print]
func (f *ExportFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExportFormat) Set(v string) error {
	switch v {
	case `DBC`, `HTML`, `JUPYTER`, `R_MARKDOWN`, `SOURCE`:
		*f = ExportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBC", "HTML", "JUPYTER", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Type always returns ExportFormat to satisfy [pflag.Value] interface
func (f *ExportFormat) Type() string {
	return "ExportFormat"
}

// Export a workspace object
type ExportRequest struct {
	// This specifies the format of the exported file. By default, this is
	// `SOURCE`.
	//
	// The value is case sensitive.
	//
	// - `SOURCE`: The notebook is exported as source code. - `HTML`: The
	// notebook is exported as an HTML file. - `JUPYTER`: The notebook is
	// exported as a Jupyter/IPython Notebook file. - `DBC`: The notebook is
	// exported in Databricks archive format. - `R_MARKDOWN`: The notebook is
	// exported to R Markdown format.
	Format ExportFormat `json:"-" url:"format,omitempty"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC` and `SOURCE` format.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *ExportRequest) UnmarshalJSON(b []byte) error {
	type C ExportRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ExportRequest) MarshalJSON() ([]byte, error) {
	type C ExportRequest
	return marshal.Marshal((C)(s))
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content string `json:"content,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ExportResponse) UnmarshalJSON(b []byte) error {
	type C ExportResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ExportResponse) MarshalJSON() ([]byte, error) {
	type C ExportResponse
	return marshal.Marshal((C)(s))
}

// Get secret ACL details
type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal string `json:"-" url:"principal"`
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAclRequest) UnmarshalJSON(b []byte) error {
	type C GetAclRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetAclRequest) MarshalJSON() ([]byte, error) {
	type C GetAclRequest
	return marshal.Marshal((C)(s))
}

type GetCredentialsResponse struct {
	Credentials []CredentialInfo `json:"credentials,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetCredentialsResponse) UnmarshalJSON(b []byte) error {
	type C GetCredentialsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetCredentialsResponse) MarshalJSON() ([]byte, error) {
	type C GetCredentialsResponse
	return marshal.Marshal((C)(s))
}

// Get a credential entry
type GetGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetGitCredentialRequest) UnmarshalJSON(b []byte) error {
	type C GetGitCredentialRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetGitCredentialRequest) MarshalJSON() ([]byte, error) {
	type C GetGitCredentialRequest
	return marshal.Marshal((C)(s))
}

// Get repo permission levels
type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetRepoPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	type C GetRepoPermissionLevelsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetRepoPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	type C GetRepoPermissionLevelsRequest
	return marshal.Marshal((C)(s))
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RepoPermissionsDescription `json:"permission_levels,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetRepoPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	type C GetRepoPermissionLevelsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetRepoPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	type C GetRepoPermissionLevelsResponse
	return marshal.Marshal((C)(s))
}

// Get repo permissions
type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetRepoPermissionsRequest) UnmarshalJSON(b []byte) error {
	type C GetRepoPermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetRepoPermissionsRequest) MarshalJSON() ([]byte, error) {
	type C GetRepoPermissionsRequest
	return marshal.Marshal((C)(s))
}

// Get a repo
type GetRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetRepoRequest) UnmarshalJSON(b []byte) error {
	type C GetRepoRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetRepoRequest) MarshalJSON() ([]byte, error) {
	type C GetRepoRequest
	return marshal.Marshal((C)(s))
}

// Get a secret
type GetSecretRequest struct {
	// The key to fetch secret for.
	Key string `json:"-" url:"key"`
	// The name of the scope to fetch secret information from.
	Scope string `json:"-" url:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *GetSecretRequest) UnmarshalJSON(b []byte) error {
	type C GetSecretRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetSecretRequest) MarshalJSON() ([]byte, error) {
	type C GetSecretRequest
	return marshal.Marshal((C)(s))
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The value of the secret in its byte representation.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetSecretResponse) UnmarshalJSON(b []byte) error {
	type C GetSecretResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetSecretResponse) MarshalJSON() ([]byte, error) {
	type C GetSecretResponse
	return marshal.Marshal((C)(s))
}

// Get status
type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *GetStatusRequest) UnmarshalJSON(b []byte) error {
	type C GetStatusRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetStatusRequest) MarshalJSON() ([]byte, error) {
	type C GetStatusRequest
	return marshal.Marshal((C)(s))
}

// Get workspace object permission levels
type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetWorkspaceObjectPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	type C GetWorkspaceObjectPermissionLevelsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetWorkspaceObjectPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	type C GetWorkspaceObjectPermissionLevelsRequest
	return marshal.Marshal((C)(s))
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []WorkspaceObjectPermissionsDescription `json:"permission_levels,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetWorkspaceObjectPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	type C GetWorkspaceObjectPermissionLevelsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetWorkspaceObjectPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	type C GetWorkspaceObjectPermissionLevelsResponse
	return marshal.Marshal((C)(s))
}

// Get workspace object permissions
type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetWorkspaceObjectPermissionsRequest) UnmarshalJSON(b []byte) error {
	type C GetWorkspaceObjectPermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s GetWorkspaceObjectPermissionsRequest) MarshalJSON() ([]byte, error) {
	type C GetWorkspaceObjectPermissionsRequest
	return marshal.Marshal((C)(s))
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown. This parameter might be absent,
	// and instead a posted file is used.
	Content string `json:"content,omitempty"`
	// This specifies the format of the file to be imported.
	//
	// The value is case sensitive.
	//
	// - `AUTO`: The item is imported depending on an analysis of the item's
	// extension and the header content provided in the request. If the item is
	// imported as a notebook, then the item's extension is automatically
	// removed. - `SOURCE`: The notebook is imported as source code. - `HTML`:
	// The notebook is imported as an HTML file. - `JUPYTER`: The notebook is
	// imported as a Jupyter/IPython Notebook file. - `DBC`: The notebook is
	// imported in Databricks archive format. Required for directories. -
	// `R_MARKDOWN`: The notebook is imported from R Markdown format.
	Format ImportFormat `json:"format,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `json:"language,omitempty"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite bool `json:"overwrite,omitempty"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` format.
	Path string `json:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *Import) UnmarshalJSON(b []byte) error {
	type C Import
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Import) MarshalJSON() ([]byte, error) {
	type C Import
	return marshal.Marshal((C)(s))
}

// This specifies the format of the file to be imported.
//
// The value is case sensitive.
//
// - `AUTO`: The item is imported depending on an analysis of the item's
// extension and the header content provided in the request. If the item is
// imported as a notebook, then the item's extension is automatically removed. -
// `SOURCE`: The notebook is imported as source code. - `HTML`: The notebook is
// imported as an HTML file. - `JUPYTER`: The notebook is imported as a
// Jupyter/IPython Notebook file. - `DBC`: The notebook is imported in
// Databricks archive format. Required for directories. - `R_MARKDOWN`: The
// notebook is imported from R Markdown format.
type ImportFormat string

const ImportFormatAuto ImportFormat = `AUTO`

const ImportFormatDbc ImportFormat = `DBC`

const ImportFormatHtml ImportFormat = `HTML`

const ImportFormatJupyter ImportFormat = `JUPYTER`

const ImportFormatRMarkdown ImportFormat = `R_MARKDOWN`

const ImportFormatSource ImportFormat = `SOURCE`

// String representation for [fmt.Print]
func (f *ImportFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ImportFormat) Set(v string) error {
	switch v {
	case `AUTO`, `DBC`, `HTML`, `JUPYTER`, `R_MARKDOWN`, `SOURCE`:
		*f = ImportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO", "DBC", "HTML", "JUPYTER", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Type always returns ImportFormat to satisfy [pflag.Value] interface
func (f *ImportFormat) Type() string {
	return "ImportFormat"
}

// The language of the object. This value is set only if the object type is
// `NOTEBOOK`.
type Language string

const LanguagePython Language = `PYTHON`

const LanguageR Language = `R`

const LanguageScala Language = `SCALA`

const LanguageSql Language = `SQL`

// String representation for [fmt.Print]
func (f *Language) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Language) Set(v string) error {
	switch v {
	case `PYTHON`, `R`, `SCALA`, `SQL`:
		*f = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PYTHON", "R", "SCALA", "SQL"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (f *Language) Type() string {
	return "Language"
}

// Lists ACLs
type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *ListAclsRequest) UnmarshalJSON(b []byte) error {
	type C ListAclsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListAclsRequest) MarshalJSON() ([]byte, error) {
	type C ListAclsRequest
	return marshal.Marshal((C)(s))
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items []AclItem `json:"items,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListAclsResponse) UnmarshalJSON(b []byte) error {
	type C ListAclsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListAclsResponse) MarshalJSON() ([]byte, error) {
	type C ListAclsResponse
	return marshal.Marshal((C)(s))
}

// Get repos
type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken string `json:"-" url:"next_page_token,omitempty"`
	// Filters repos that have paths starting with the given path prefix.
	PathPrefix string `json:"-" url:"path_prefix,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListReposRequest) UnmarshalJSON(b []byte) error {
	type C ListReposRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListReposRequest) MarshalJSON() ([]byte, error) {
	type C ListReposRequest
	return marshal.Marshal((C)(s))
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the GET /repos
	// endpoint to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Repos []RepoInfo `json:"repos,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListReposResponse) UnmarshalJSON(b []byte) error {
	type C ListReposResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListReposResponse) MarshalJSON() ([]byte, error) {
	type C ListReposResponse
	return marshal.Marshal((C)(s))
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `json:"objects,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListResponse) UnmarshalJSON(b []byte) error {
	type C ListResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListResponse) MarshalJSON() ([]byte, error) {
	type C ListResponse
	return marshal.Marshal((C)(s))
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes []SecretScope `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListScopesResponse) UnmarshalJSON(b []byte) error {
	type C ListScopesResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListScopesResponse) MarshalJSON() ([]byte, error) {
	type C ListScopesResponse
	return marshal.Marshal((C)(s))
}

// List secret keys
type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope string `json:"-" url:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *ListSecretsRequest) UnmarshalJSON(b []byte) error {
	type C ListSecretsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListSecretsRequest) MarshalJSON() ([]byte, error) {
	type C ListSecretsRequest
	return marshal.Marshal((C)(s))
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets []SecretMetadata `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListSecretsResponse) UnmarshalJSON(b []byte) error {
	type C ListSecretsResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListSecretsResponse) MarshalJSON() ([]byte, error) {
	type C ListSecretsResponse
	return marshal.Marshal((C)(s))
}

// List contents
type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter int `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *ListWorkspaceRequest) UnmarshalJSON(b []byte) error {
	type C ListWorkspaceRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ListWorkspaceRequest) MarshalJSON() ([]byte, error) {
	type C ListWorkspaceRequest
	return marshal.Marshal((C)(s))
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *Mkdirs) UnmarshalJSON(b []byte) error {
	type C Mkdirs
	return marshal.Unmarshal(b, (*C)(s))
}

func (s Mkdirs) MarshalJSON() ([]byte, error) {
	type C Mkdirs
	return marshal.Marshal((C)(s))
}

type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `json:"language,omitempty"`
	// Only applicable to files, the last modified UTC timestamp.
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// Unique identifier for the object.
	ObjectId int64 `json:"object_id,omitempty"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository
	ObjectType ObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	Path string `json:"path,omitempty"`
	// Only applicable to files. The file size in bytes can be returned.
	Size int64 `json:"size,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ObjectInfo) UnmarshalJSON(b []byte) error {
	type C ObjectInfo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s ObjectInfo) MarshalJSON() ([]byte, error) {
	type C ObjectInfo
	return marshal.Marshal((C)(s))
}

// The type of the object in workspace.
//
// - `NOTEBOOK`: document that contains runnable code, visualizations, and
// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
// file - `REPO`: repository
type ObjectType string

const ObjectTypeDirectory ObjectType = `DIRECTORY`

const ObjectTypeFile ObjectType = `FILE`

const ObjectTypeLibrary ObjectType = `LIBRARY`

const ObjectTypeNotebook ObjectType = `NOTEBOOK`

const ObjectTypeRepo ObjectType = `REPO`

// String representation for [fmt.Print]
func (f *ObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ObjectType) Set(v string) error {
	switch v {
	case `DIRECTORY`, `FILE`, `LIBRARY`, `NOTEBOOK`, `REPO`:
		*f = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECTORY", "FILE", "LIBRARY", "NOTEBOOK", "REPO"`, v)
	}
}

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (f *ObjectType) Type() string {
	return "ObjectType"
}

type PutAcl struct {
	// The permission level applied to the principal.
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	Principal string `json:"principal"`
	// The name of the scope to apply permissions to.
	Scope string `json:"scope"`

	ForceSendFields []string `json:"-"`
}

func (s *PutAcl) UnmarshalJSON(b []byte) error {
	type C PutAcl
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PutAcl) MarshalJSON() ([]byte, error) {
	type C PutAcl
	return marshal.Marshal((C)(s))
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	BytesValue string `json:"bytes_value,omitempty"`
	// A unique name to identify the secret.
	Key string `json:"key"`
	// The name of the scope to which the secret will be associated with.
	Scope string `json:"scope"`
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PutSecret) UnmarshalJSON(b []byte) error {
	type C PutSecret
	return marshal.Unmarshal(b, (*C)(s))
}

func (s PutSecret) MarshalJSON() ([]byte, error) {
	type C PutSecret
	return marshal.Marshal((C)(s))
}

type RepoAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoAccessControlRequest) UnmarshalJSON(b []byte) error {
	type C RepoAccessControlRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoAccessControlRequest) MarshalJSON() ([]byte, error) {
	type C RepoAccessControlRequest
	return marshal.Marshal((C)(s))
}

type RepoAccessControlResponse struct {
	// All permissions.
	AllPermissions []RepoPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoAccessControlResponse) UnmarshalJSON(b []byte) error {
	type C RepoAccessControlResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoAccessControlResponse) MarshalJSON() ([]byte, error) {
	type C RepoAccessControlResponse
	return marshal.Marshal((C)(s))
}

type RepoInfo struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// ID of the repo object in the workspace.
	Id int64 `json:"id,omitempty"`
	// Desired path for the repo in the workspace. Must be in the format
	// /Repos/{folder}/{repo-name}.
	Path string `json:"path,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	Provider string `json:"provider,omitempty"`

	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the Git repository to be linked.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoInfo) UnmarshalJSON(b []byte) error {
	type C RepoInfo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoInfo) MarshalJSON() ([]byte, error) {
	type C RepoInfo
	return marshal.Marshal((C)(s))
}

type RepoPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoPermission) UnmarshalJSON(b []byte) error {
	type C RepoPermission
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoPermission) MarshalJSON() ([]byte, error) {
	type C RepoPermission
	return marshal.Marshal((C)(s))
}

// Permission level
type RepoPermissionLevel string

const RepoPermissionLevelCanEdit RepoPermissionLevel = `CAN_EDIT`

const RepoPermissionLevelCanManage RepoPermissionLevel = `CAN_MANAGE`

const RepoPermissionLevelCanRead RepoPermissionLevel = `CAN_READ`

const RepoPermissionLevelCanRun RepoPermissionLevel = `CAN_RUN`

// String representation for [fmt.Print]
func (f *RepoPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RepoPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`, `CAN_RUN`:
		*f = RepoPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ", "CAN_RUN"`, v)
	}
}

// Type always returns RepoPermissionLevel to satisfy [pflag.Value] interface
func (f *RepoPermissionLevel) Type() string {
	return "RepoPermissionLevel"
}

type RepoPermissions struct {
	AccessControlList []RepoAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoPermissions) UnmarshalJSON(b []byte) error {
	type C RepoPermissions
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoPermissions) MarshalJSON() ([]byte, error) {
	type C RepoPermissions
	return marshal.Marshal((C)(s))
}

type RepoPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoPermissionsDescription) UnmarshalJSON(b []byte) error {
	type C RepoPermissionsDescription
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoPermissionsDescription) MarshalJSON() ([]byte, error) {
	type C RepoPermissionsDescription
	return marshal.Marshal((C)(s))
}

type RepoPermissionsRequest struct {
	AccessControlList []RepoAccessControlRequest `json:"access_control_list,omitempty"`
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *RepoPermissionsRequest) UnmarshalJSON(b []byte) error {
	type C RepoPermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s RepoPermissionsRequest) MarshalJSON() ([]byte, error) {
	type C RepoPermissionsRequest
	return marshal.Marshal((C)(s))
}

type ScopeBackendType string

const ScopeBackendTypeAzureKeyvault ScopeBackendType = `AZURE_KEYVAULT`

const ScopeBackendTypeDatabricks ScopeBackendType = `DATABRICKS`

// String representation for [fmt.Print]
func (f *ScopeBackendType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ScopeBackendType) Set(v string) error {
	switch v {
	case `AZURE_KEYVAULT`, `DATABRICKS`:
		*f = ScopeBackendType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_KEYVAULT", "DATABRICKS"`, v)
	}
}

// Type always returns ScopeBackendType to satisfy [pflag.Value] interface
func (f *ScopeBackendType) Type() string {
	return "ScopeBackendType"
}

type SecretMetadata struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SecretMetadata) UnmarshalJSON(b []byte) error {
	type C SecretMetadata
	return marshal.Unmarshal(b, (*C)(s))
}

func (s SecretMetadata) MarshalJSON() ([]byte, error) {
	type C SecretMetadata
	return marshal.Marshal((C)(s))
}

type SecretScope struct {
	// The type of secret scope backend.
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
	// A unique name to identify the secret scope.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SecretScope) UnmarshalJSON(b []byte) error {
	type C SecretScope
	return marshal.Unmarshal(b, (*C)(s))
}

func (s SecretScope) MarshalJSON() ([]byte, error) {
	type C SecretScope
	return marshal.Marshal((C)(s))
}

type SparseCheckout struct {
	// List of patterns to include for sparse checkout.
	Patterns []string `json:"patterns,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SparseCheckout) UnmarshalJSON(b []byte) error {
	type C SparseCheckout
	return marshal.Unmarshal(b, (*C)(s))
}

func (s SparseCheckout) MarshalJSON() ([]byte, error) {
	type C SparseCheckout
	return marshal.Marshal((C)(s))
}

type SparseCheckoutUpdate struct {
	// List of patterns to include for sparse checkout.
	Patterns []string `json:"patterns,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SparseCheckoutUpdate) UnmarshalJSON(b []byte) error {
	type C SparseCheckoutUpdate
	return marshal.Unmarshal(b, (*C)(s))
}

func (s SparseCheckoutUpdate) MarshalJSON() ([]byte, error) {
	type C SparseCheckoutUpdate
	return marshal.Marshal((C)(s))
}

type UpdateCredentials struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider.
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateCredentials) UnmarshalJSON(b []byte) error {
	type C UpdateCredentials
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdateCredentials) MarshalJSON() ([]byte, error) {
	type C UpdateCredentials
	return marshal.Marshal((C)(s))
}

type UpdateRepo struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" url:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout *SparseCheckoutUpdate `json:"sparse_checkout,omitempty"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag string `json:"tag,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateRepo) UnmarshalJSON(b []byte) error {
	type C UpdateRepo
	return marshal.Unmarshal(b, (*C)(s))
}

func (s UpdateRepo) MarshalJSON() ([]byte, error) {
	type C UpdateRepo
	return marshal.Marshal((C)(s))
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`
	// name of the service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspaceObjectAccessControlRequest) UnmarshalJSON(b []byte) error {
	type C WorkspaceObjectAccessControlRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspaceObjectAccessControlRequest) MarshalJSON() ([]byte, error) {
	type C WorkspaceObjectAccessControlRequest
	return marshal.Marshal((C)(s))
}

type WorkspaceObjectAccessControlResponse struct {
	// All permissions.
	AllPermissions []WorkspaceObjectPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspaceObjectAccessControlResponse) UnmarshalJSON(b []byte) error {
	type C WorkspaceObjectAccessControlResponse
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspaceObjectAccessControlResponse) MarshalJSON() ([]byte, error) {
	type C WorkspaceObjectAccessControlResponse
	return marshal.Marshal((C)(s))
}

type WorkspaceObjectPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspaceObjectPermission) UnmarshalJSON(b []byte) error {
	type C WorkspaceObjectPermission
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspaceObjectPermission) MarshalJSON() ([]byte, error) {
	type C WorkspaceObjectPermission
	return marshal.Marshal((C)(s))
}

// Permission level
type WorkspaceObjectPermissionLevel string

const WorkspaceObjectPermissionLevelCanEdit WorkspaceObjectPermissionLevel = `CAN_EDIT`

const WorkspaceObjectPermissionLevelCanManage WorkspaceObjectPermissionLevel = `CAN_MANAGE`

const WorkspaceObjectPermissionLevelCanRead WorkspaceObjectPermissionLevel = `CAN_READ`

const WorkspaceObjectPermissionLevelCanRun WorkspaceObjectPermissionLevel = `CAN_RUN`

// String representation for [fmt.Print]
func (f *WorkspaceObjectPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceObjectPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`, `CAN_RUN`:
		*f = WorkspaceObjectPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ", "CAN_RUN"`, v)
	}
}

// Type always returns WorkspaceObjectPermissionLevel to satisfy [pflag.Value] interface
func (f *WorkspaceObjectPermissionLevel) Type() string {
	return "WorkspaceObjectPermissionLevel"
}

type WorkspaceObjectPermissions struct {
	AccessControlList []WorkspaceObjectAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspaceObjectPermissions) UnmarshalJSON(b []byte) error {
	type C WorkspaceObjectPermissions
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspaceObjectPermissions) MarshalJSON() ([]byte, error) {
	type C WorkspaceObjectPermissions
	return marshal.Marshal((C)(s))
}

type WorkspaceObjectPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspaceObjectPermissionsDescription) UnmarshalJSON(b []byte) error {
	type C WorkspaceObjectPermissionsDescription
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspaceObjectPermissionsDescription) MarshalJSON() ([]byte, error) {
	type C WorkspaceObjectPermissionsDescription
	return marshal.Marshal((C)(s))
}

type WorkspaceObjectPermissionsRequest struct {
	AccessControlList []WorkspaceObjectAccessControlRequest `json:"access_control_list,omitempty"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *WorkspaceObjectPermissionsRequest) UnmarshalJSON(b []byte) error {
	type C WorkspaceObjectPermissionsRequest
	return marshal.Unmarshal(b, (*C)(s))
}

func (s WorkspaceObjectPermissionsRequest) MarshalJSON() ([]byte, error) {
	type C WorkspaceObjectPermissionsRequest
	return marshal.Marshal((C)(s))
}
