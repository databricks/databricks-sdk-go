// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// An item representing an ACL rule applied to the given principal (user or
// group) on the associated scope point.
type AclItem struct {
	// The permission level applied to the principal.
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	Principal string `json:"principal"`
}

// The ACL permission levels for Secret ACLs applied to secret scopes.
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

// Values returns all possible values for AclPermission.
//
// There is no guarantee on the order of the values in the slice.
func (f *AclPermission) Values() []AclPermission {
	return []AclPermission{
		AclPermissionManage,
		AclPermissionRead,
		AclPermissionWrite,
	}
}

// Type always returns AclPermission to satisfy [pflag.Value] interface
func (f *AclPermission) Type() string {
	return "AclPermission"
}

// The metadata of the Azure KeyVault for a secret scope of type
// `AZURE_KEYVAULT`
type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName string `json:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId string `json:"resource_id"`
}

type CreateCredentialsRequest struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername string `json:"git_username,omitempty"`
	// if the credential is the default for the given provider
	IsDefaultForProvider bool `json:"is_default_for_provider,omitempty"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name string `json:"name,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername string `json:"git_username,omitempty"`
	// if the credential is the default for the given provider
	IsDefaultForProvider bool `json:"is_default_for_provider,omitempty"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRepoRequest struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in `/Repos`, path must be in
	// the format `/Repos/{folder}/{repo-name}`.
	Path string `json:"path,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	Provider string `json:"provider"`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the Git repository to be linked.
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRepoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRepoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateRepoResponse struct {
	// Branch that the Git folder (repo) is checked out to.
	Branch string `json:"branch,omitempty"`
	// SHA-1 hash representing the commit ID of the current HEAD of the Git
	// folder (repo).
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	Id int64 `json:"id,omitempty"`
	// Path of the Git folder (repo) in the workspace.
	Path string `json:"path,omitempty"`
	// Git provider of the linked Git repository.
	Provider string `json:"provider,omitempty"`
	// Sparse checkout settings for the Git folder (repo).
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the linked Git repository.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateRepoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRepoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateScope struct {
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata `json:"backend_azure_keyvault,omitempty"`
	// The principal that is initially granted ``MANAGE`` permission to the
	// created scope.
	InitialManagePrincipal string `json:"initial_manage_principal,omitempty"`
	// Scope name requested by the user. Scope names are unique.
	Scope string `json:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to ``DATABRICKS``
	ScopeBackendType ScopeBackendType `json:"scope_backend_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateScope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateScope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider string `json:"git_provider,omitempty"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername string `json:"git_username,omitempty"`
	// if the credential is the default for the given provider
	IsDefaultForProvider bool `json:"is_default_for_provider,omitempty"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Delete) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Delete) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal string `json:"principal"`
	// The name of the scope to remove permissions from.
	Scope string `json:"scope"`
}

type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
}

type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	RepoId int64 `json:"-" url:"-"`
}

type DeleteScope struct {
	// Name of the scope to delete.
	Scope string `json:"scope"`
}

type DeleteSecret struct {
	// Name of the secret to delete.
	Key string `json:"key"`
	// The name of the scope that contains the secret to delete.
	Scope string `json:"scope"`
}

// The format for workspace import and export.
type ExportFormat string

const ExportFormatAuto ExportFormat = `AUTO`

const ExportFormatDbc ExportFormat = `DBC`

const ExportFormatHtml ExportFormat = `HTML`

const ExportFormatJupyter ExportFormat = `JUPYTER`

const ExportFormatRaw ExportFormat = `RAW`

const ExportFormatRMarkdown ExportFormat = `R_MARKDOWN`

const ExportFormatSource ExportFormat = `SOURCE`

// String representation for [fmt.Print]
func (f *ExportFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExportFormat) Set(v string) error {
	switch v {
	case `AUTO`, `DBC`, `HTML`, `JUPYTER`, `RAW`, `R_MARKDOWN`, `SOURCE`:
		*f = ExportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO", "DBC", "HTML", "JUPYTER", "RAW", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Values returns all possible values for ExportFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExportFormat) Values() []ExportFormat {
	return []ExportFormat{
		ExportFormatAuto,
		ExportFormatDbc,
		ExportFormatHtml,
		ExportFormatJupyter,
		ExportFormatRaw,
		ExportFormatRMarkdown,
		ExportFormatSource,
	}
}

// Type always returns ExportFormat to satisfy [pflag.Value] interface
func (f *ExportFormat) Type() string {
	return "ExportFormat"
}

type ExportRequest struct {
	// This specifies the format of the exported file. By default, this is
	// `SOURCE`.
	//
	// The value is case sensitive.
	//
	// - `SOURCE`: The notebook is exported as source code. Directory exports
	// will not include non-notebook entries. - `HTML`: The notebook is exported
	// as an HTML file. - `JUPYTER`: The notebook is exported as a
	// Jupyter/IPython Notebook file. - `DBC`: The notebook is exported in
	// Databricks archive format. Directory exports will not include
	// non-notebook entries. - `R_MARKDOWN`: The notebook is exported to R
	// Markdown format. - `AUTO`: The object or directory is exported depending
	// on the objects type. Directory exports will include notebooks and
	// workspace files.
	Format ExportFormat `json:"-" url:"format,omitempty"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC`, `SOURCE`, and `AUTO` format.
	Path string `json:"-" url:"path"`
}

// The request field `direct_download` determines whether a JSON response or
// binary contents are returned by this endpoint.
type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content string `json:"content,omitempty"`
	// The file type of the exported file.
	FileType string `json:"file_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExportResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExportResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal string `json:"-" url:"principal"`
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`
}

type GetCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
}

type GetCredentialsResponse struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider string `json:"git_provider,omitempty"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername string `json:"git_username,omitempty"`
	// if the credential is the default for the given provider
	IsDefaultForProvider bool `json:"is_default_for_provider,omitempty"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []RepoPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`
}

type GetRepoRequest struct {
	// ID of the Git folder (repo) object in the workspace.
	RepoId int64 `json:"-" url:"-"`
}

type GetRepoResponse struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	Id int64 `json:"id,omitempty"`
	// Path of the Git folder (repo) in the workspace.
	Path string `json:"path,omitempty"`
	// Git provider of the linked Git repository.
	Provider string `json:"provider,omitempty"`
	// Sparse checkout settings for the Git folder (repo).
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the linked Git repository.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRepoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRepoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSecretRequest struct {
	// Name of the secret to fetch value information.
	Key string `json:"-" url:"key"`
	// The name of the scope that contains the secret.
	Scope string `json:"-" url:"scope"`
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The value of the secret in its byte representation.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetSecretResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSecretResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`
}

type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []WorkspaceObjectPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`
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
	// removed. - `SOURCE`: The notebook or directory is imported as source
	// code. - `HTML`: The notebook is imported as an HTML file. - `JUPYTER`:
	// The notebook is imported as a Jupyter/IPython Notebook file. - `DBC`: The
	// notebook is imported in Databricks archive format. Required for
	// directories. - `R_MARKDOWN`: The notebook is imported from R Markdown
	// format.
	Format ImportFormat `json:"format,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `json:"language,omitempty"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	Overwrite bool `json:"overwrite,omitempty"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Import) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Import) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The format for workspace import and export.
type ImportFormat string

const ImportFormatAuto ImportFormat = `AUTO`

const ImportFormatDbc ImportFormat = `DBC`

const ImportFormatHtml ImportFormat = `HTML`

const ImportFormatJupyter ImportFormat = `JUPYTER`

const ImportFormatRaw ImportFormat = `RAW`

const ImportFormatRMarkdown ImportFormat = `R_MARKDOWN`

const ImportFormatSource ImportFormat = `SOURCE`

// String representation for [fmt.Print]
func (f *ImportFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ImportFormat) Set(v string) error {
	switch v {
	case `AUTO`, `DBC`, `HTML`, `JUPYTER`, `RAW`, `R_MARKDOWN`, `SOURCE`:
		*f = ImportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO", "DBC", "HTML", "JUPYTER", "RAW", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Values returns all possible values for ImportFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *ImportFormat) Values() []ImportFormat {
	return []ImportFormat{
		ImportFormatAuto,
		ImportFormatDbc,
		ImportFormatHtml,
		ImportFormatJupyter,
		ImportFormatRaw,
		ImportFormatRMarkdown,
		ImportFormatSource,
	}
}

// Type always returns ImportFormat to satisfy [pflag.Value] interface
func (f *ImportFormat) Type() string {
	return "ImportFormat"
}

// The language of notebook.
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

// Values returns all possible values for Language.
//
// There is no guarantee on the order of the values in the slice.
func (f *Language) Values() []Language {
	return []Language{
		LanguagePython,
		LanguageR,
		LanguageScala,
		LanguageSql,
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (f *Language) Type() string {
	return "Language"
}

type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items []AclItem `json:"items,omitempty"`
}

type ListCredentialsResponse struct {
	// List of credentials.
	Credentials []CredentialInfo `json:"credentials,omitempty"`
}

type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken string `json:"-" url:"next_page_token,omitempty"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided or when provided an effectively empty prefix (`/` or
	// `/Workspace`) Git folders (repos) from `/Workspace/Repos` will be served.
	PathPrefix string `json:"-" url:"path_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListReposRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListReposRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Git folders (repos).
	Repos []RepoInfo `json:"repos,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListReposResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListReposResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `json:"objects,omitempty"`
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes []SecretScope `json:"scopes,omitempty"`
}

type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope string `json:"-" url:"scope"`
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets []SecretMetadata `json:"secrets,omitempty"`
}

type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter int64 `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`
}

// The information of the object in workspace. It will be returned by “list“
// and “get-status“.
type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	Language Language `json:"language,omitempty"`
	// Only applicable to files, the last modified UTC timestamp.
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// Unique identifier for the object.
	ObjectId int64 `json:"object_id,omitempty"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	ObjectType ObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	Path string `json:"path,omitempty"`
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	ResourceId string `json:"resource_id,omitempty"`
	// Only applicable to files. The file size in bytes can be returned.
	Size int64 `json:"size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ObjectInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ObjectInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of the object in workspace.
type ObjectType string

const ObjectTypeDashboard ObjectType = `DASHBOARD`

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
	case `DASHBOARD`, `DIRECTORY`, `FILE`, `LIBRARY`, `NOTEBOOK`, `REPO`:
		*f = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD", "DIRECTORY", "FILE", "LIBRARY", "NOTEBOOK", "REPO"`, v)
	}
}

// Values returns all possible values for ObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ObjectType) Values() []ObjectType {
	return []ObjectType{
		ObjectTypeDashboard,
		ObjectTypeDirectory,
		ObjectTypeFile,
		ObjectTypeLibrary,
		ObjectTypeNotebook,
		ObjectTypeRepo,
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepoAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepoAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Git folder (repo) information.
type RepoInfo struct {
	// Name of the current git branch of the git folder (repo).
	Branch string `json:"branch,omitempty"`
	// Current git commit id of the git folder (repo).
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// Id of the git folder (repo) in the Workspace.
	Id int64 `json:"id,omitempty"`
	// Root path of the git folder (repo) in the Workspace.
	Path string `json:"path,omitempty"`
	// Git provider of the remote git repository, e.g. `gitHub`.
	Provider string `json:"provider,omitempty"`
	// Sparse checkout config for the git folder (repo).
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the remote git repository.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepoInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepoPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for RepoPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *RepoPermissionLevel) Values() []RepoPermissionLevel {
	return []RepoPermissionLevel{
		RepoPermissionLevelCanEdit,
		RepoPermissionLevelCanManage,
		RepoPermissionLevelCanRead,
		RepoPermissionLevelCanRun,
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepoPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RepoPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepoPermissionsRequest struct {
	AccessControlList []RepoAccessControlRequest `json:"access_control_list,omitempty"`
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`
}

// The types of secret scope backends in the Secret Manager. Azure KeyVault
// backed secret scopes will be supported in a later release.
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

// Values returns all possible values for ScopeBackendType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ScopeBackendType) Values() []ScopeBackendType {
	return []ScopeBackendType{
		ScopeBackendTypeAzureKeyvault,
		ScopeBackendTypeDatabricks,
	}
}

// Type always returns ScopeBackendType to satisfy [pflag.Value] interface
func (f *ScopeBackendType) Type() string {
	return "ScopeBackendType"
}

// The metadata about a secret. Returned when listing secrets. Does not contain
// the actual secret value.
type SecretMetadata struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SecretMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An organizational resource for storing secrets. Secret scopes can be
// different types (Databricks-managed, Azure KeyVault backed, etc), and ACLs
// can be applied to control permissions for all secrets within a scope.
type SecretScope struct {
	// The type of secret scope backend.
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
	// A unique name to identify the secret scope.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SecretScope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretScope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns []string `json:"patterns,omitempty"`
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns []string `json:"patterns,omitempty"`
}

type UpdateCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	GitUsername string `json:"git_username,omitempty"`
	// if the credential is the default for the given provider
	IsDefaultForProvider bool `json:"is_default_for_provider,omitempty"`
	// the name of the git credential, used for identification and ease of
	// lookup
	Name string `json:"name,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateRepoRequest struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	RepoId int64 `json:"-" url:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout *SparseCheckoutUpdate `json:"sparse_checkout,omitempty"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag string `json:"tag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateRepoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRepoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceObjectAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceObjectAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceObjectPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Values returns all possible values for WorkspaceObjectPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceObjectPermissionLevel) Values() []WorkspaceObjectPermissionLevel {
	return []WorkspaceObjectPermissionLevel{
		WorkspaceObjectPermissionLevelCanEdit,
		WorkspaceObjectPermissionLevelCanManage,
		WorkspaceObjectPermissionLevelCanRead,
		WorkspaceObjectPermissionLevelCanRun,
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceObjectPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceObjectPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WorkspaceObjectPermissionsRequest struct {
	AccessControlList []WorkspaceObjectAccessControlRequest `json:"access_control_list,omitempty"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`
}
