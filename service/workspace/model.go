// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import "fmt"

// all definitions in this file are in alphabetical order

type AclItem struct {
	// The permission level applied to the principal.
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	Principal string `json:"principal"`
}

type AclPermission string

const AclPermissionManage AclPermission = `MANAGE`

const AclPermissionRead AclPermission = `READ`

const AclPermissionWrite AclPermission = `WRITE`

// String representation for [fmt.Print]
func (ap *AclPermission) String() string {
	return string(*ap)
}

// Set raw string value and validate it against allowed values
func (ap *AclPermission) Set(v string) error {
	switch v {
	case `MANAGE`, `READ`, `WRITE`:
		*ap = AclPermission(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGE", "READ", "WRITE"`, v)
	}
}

// Type always returns AclPermission to satisfy [pflag.Value] interface
func (ap *AclPermission) Type() string {
	return "AclPermission"
}

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	DnsName string `json:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId string `json:"resource_id"`
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
}

type CreateScope struct {
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal string `json:"initial_manage_principal,omitempty"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
	// Scope name requested by the user. Scope names are unique.
	Scope string `json:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType ScopeBackendType `json:"scope_backend_type,omitempty"`
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are gitHub, bitbucketCloud, gitLab, azureDevOpsServices,
	// gitHubEnterprise, bitbucketServer, gitLabEnterpriseEdition and
	// awsCodeCommit.
	GitProvider string `json:"git_provider,omitempty"`
	// Git username.
	GitUsername string `json:"git_username,omitempty"`
}

type Delete struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	Principal string `json:"principal"`
	// The name of the scope to remove permissions from.
	Scope string `json:"scope"`
}

// Delete a credential
type DeleteGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
}

// Delete a repo
type DeleteRepoRequest struct {
	// The ID for the corresponding repo to access.
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

// This specifies the format of the file to be imported. By default, this is
// `SOURCE`.
//
// If using `AUTO` the item is imported or exported as either a workspace file
// or a notebook,depending on an analysis of the item’s extension and the
// header content provided in the request. The value is case sensitive. In
// addition, if the item is imported as a notebook, then the item’s extension
// is automatically removed.
type ExportFormat string

const ExportFormatAuto ExportFormat = `AUTO`

const ExportFormatDbc ExportFormat = `DBC`

const ExportFormatHtml ExportFormat = `HTML`

const ExportFormatJupyter ExportFormat = `JUPYTER`

const ExportFormatRMarkdown ExportFormat = `R_MARKDOWN`

const ExportFormatSource ExportFormat = `SOURCE`

// String representation for [fmt.Print]
func (ef *ExportFormat) String() string {
	return string(*ef)
}

// Set raw string value and validate it against allowed values
func (ef *ExportFormat) Set(v string) error {
	switch v {
	case `AUTO`, `DBC`, `HTML`, `JUPYTER`, `R_MARKDOWN`, `SOURCE`:
		*ef = ExportFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTO", "DBC", "HTML", "JUPYTER", "R_MARKDOWN", "SOURCE"`, v)
	}
}

// Type always returns ExportFormat to satisfy [pflag.Value] interface
func (ef *ExportFormat) Type() string {
	return "ExportFormat"
}

// Export a workspace object
type ExportRequest struct {
	// Flag to enable direct download. If it is `true`, the response will be the
	// exported file itself. Otherwise, the response contains content as base64
	// encoded string.
	DirectDownload bool `json:"-" url:"direct_download,omitempty"`
	// This specifies the format of the exported file. By default, this is
	// `SOURCE`. However it may be one of: `SOURCE`, `HTML`, `JUPYTER`, `DBC`.
	//
	// The value is case sensitive.
	Format ExportFormat `json:"-" url:"format,omitempty"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC` format.
	Path string `json:"-" url:"path"`
}

type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown.
	Content string `json:"content,omitempty"`
}

// Get secret ACL details
type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal string `json:"-" url:"principal"`
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`
}

type GetCredentialsResponse struct {
	Credentials []CredentialInfo `json:"credentials,omitempty"`
}

// Get a credential entry
type GetGitCredentialRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
}

// Get a repo
type GetRepoRequest struct {
	// The ID for the corresponding repo to access.
	RepoId int64 `json:"-" url:"-"`
}

// Get status
type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** will be thrown. This parameter might be
	// absent, and instead a posted file will be used.
	Content string `json:"content,omitempty"`
	// This specifies the format of the file to be imported. By default, this is
	// `SOURCE`.
	//
	// If using `AUTO` the item is imported or exported as either a workspace
	// file or a notebook,depending on an analysis of the item’s extension and
	// the header content provided in the request. The value is case sensitive.
	// In addition, if the item is imported as a notebook, then the item’s
	// extension is automatically removed.
	Format ExportFormat `json:"format,omitempty"`
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
}

// The language of the object. This value is set only if the object type is
// `NOTEBOOK`.
type Language string

const LanguagePython Language = `PYTHON`

const LanguageR Language = `R`

const LanguageScala Language = `SCALA`

const LanguageSql Language = `SQL`

// String representation for [fmt.Print]
func (l *Language) String() string {
	return string(*l)
}

// Set raw string value and validate it against allowed values
func (l *Language) Set(v string) error {
	switch v {
	case `PYTHON`, `R`, `SCALA`, `SQL`:
		*l = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PYTHON", "R", "SCALA", "SQL"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (l *Language) Type() string {
	return "Language"
}

// Lists ACLs
type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items []AclItem `json:"items,omitempty"`
}

// Get repos
type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken string `json:"-" url:"next_page_token,omitempty"`
	// Filters repos that have paths starting with the given path prefix.
	PathPrefix string `json:"-" url:"path_prefix,omitempty"`
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the GET /repos
	// endpoint to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Repos []RepoInfo `json:"repos,omitempty"`
}

type ListResponse struct {
	// List of objects.
	Objects []ObjectInfo `json:"objects,omitempty"`
}

type ListScopesResponse struct {
	// The available secret scopes.
	Scopes []SecretScope `json:"scopes,omitempty"`
}

// List secret keys
type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope string `json:"-" url:"scope"`
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets []SecretMetadata `json:"secrets,omitempty"`
}

// List contents
type ListWorkspaceRequest struct {
	// <content needed>
	NotebooksModifiedAfter int `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`
}

type ObjectInfo struct {
	// <content needed>
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	Language Language `json:"language,omitempty"`
	// <content needed>
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// <content needed>
	ObjectId int64 `json:"object_id,omitempty"`
	// The type of the object in workspace.
	ObjectType ObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	Path string `json:"path,omitempty"`
	// <content needed>
	Size int64 `json:"size,omitempty"`
}

// The type of the object in workspace.
type ObjectType string

const ObjectTypeDirectory ObjectType = `DIRECTORY`

const ObjectTypeFile ObjectType = `FILE`

const ObjectTypeLibrary ObjectType = `LIBRARY`

const ObjectTypeNotebook ObjectType = `NOTEBOOK`

const ObjectTypeRepo ObjectType = `REPO`

// String representation for [fmt.Print]
func (ot *ObjectType) String() string {
	return string(*ot)
}

// Set raw string value and validate it against allowed values
func (ot *ObjectType) Set(v string) error {
	switch v {
	case `DIRECTORY`, `FILE`, `LIBRARY`, `NOTEBOOK`, `REPO`:
		*ot = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECTORY", "FILE", "LIBRARY", "NOTEBOOK", "REPO"`, v)
	}
}

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (ot *ObjectType) Type() string {
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
}

type ScopeBackendType string

const ScopeBackendTypeAzureKeyvault ScopeBackendType = `AZURE_KEYVAULT`

const ScopeBackendTypeDatabricks ScopeBackendType = `DATABRICKS`

// String representation for [fmt.Print]
func (sbt *ScopeBackendType) String() string {
	return string(*sbt)
}

// Set raw string value and validate it against allowed values
func (sbt *ScopeBackendType) Set(v string) error {
	switch v {
	case `AZURE_KEYVAULT`, `DATABRICKS`:
		*sbt = ScopeBackendType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_KEYVAULT", "DATABRICKS"`, v)
	}
}

// Type always returns ScopeBackendType to satisfy [pflag.Value] interface
func (sbt *ScopeBackendType) Type() string {
	return "ScopeBackendType"
}

type SecretMetadata struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
}

type SecretScope struct {
	// The type of secret scope backend.
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
	// A unique name to identify the secret scope.
	Name string `json:"name,omitempty"`
}

type SparseCheckout struct {
	// List of patterns to include for sparse checkout.
	Patterns []string `json:"patterns,omitempty"`
}

type SparseCheckoutUpdate struct {
	// List of patterns to include for sparse checkout.
	Patterns []string `json:"patterns,omitempty"`
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
}
