// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AclItem struct {
	// The permission level applied to the principal.
	// Wire name: 'permission'
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	// Wire name: 'principal'
	Principal string `json:"principal"`
}

func (st *AclItem) EncodeValues(key string, v *url.Values) error {
	pb, err := aclItemToPb(st)
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

func (st *AclItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aclItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aclItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AclItem) MarshalJSON() ([]byte, error) {
	pb, err := aclItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	// Wire name: 'dns_name'
	DnsName string `json:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	// Wire name: 'resource_id'
	ResourceId string `json:"resource_id"`
}

func (st *AzureKeyVaultSecretScopeMetadata) EncodeValues(key string, v *url.Values) error {
	pb, err := azureKeyVaultSecretScopeMetadataToPb(st)
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

func (st *AzureKeyVaultSecretScopeMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureKeyVaultSecretScopeMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureKeyVaultSecretScopeMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureKeyVaultSecretScopeMetadata) MarshalJSON() ([]byte, error) {
	pb, err := azureKeyVaultSecretScopeMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialsRequest struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'git_provider'
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	// Wire name: 'git_username'
	GitUsername string `json:"git_username,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	// Wire name: 'personal_access_token'
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCredentialsRequestToPb(st)
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

func (st *CreateCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateCredentialsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createCredentialsResponseToPb(st)
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

func (st *CreateCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateRepoRequest struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in `/Repos`, path must be in
	// the format `/Repos/{folder}/{repo-name}`.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'provider'
	Provider string `json:"provider"`
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the Git repository to be linked.
	// Wire name: 'url'
	Url string `json:"url"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateRepoRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createRepoRequestToPb(st)
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

func (st *CreateRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := createRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateRepoResponse struct {
	// Branch that the Git folder (repo) is checked out to.
	// Wire name: 'branch'
	Branch string `json:"branch,omitempty"`
	// SHA-1 hash representing the commit ID of the current HEAD of the Git
	// folder (repo).
	// Wire name: 'head_commit_id'
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'id'
	Id int64 `json:"id,omitempty"`
	// Path of the Git folder (repo) in the workspace.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// Git provider of the linked Git repository.
	// Wire name: 'provider'
	Provider string `json:"provider,omitempty"`
	// Sparse checkout settings for the Git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the linked Git repository.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateRepoResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createRepoResponseToPb(st)
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

func (st *CreateRepoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRepoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRepoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRepoResponse) MarshalJSON() ([]byte, error) {
	pb, err := createRepoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	// Wire name: 'backend_azure_keyvault'
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata `json:"backend_azure_keyvault,omitempty"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	// Wire name: 'initial_manage_principal'
	InitialManagePrincipal string `json:"initial_manage_principal,omitempty"`
	// Scope name requested by the user. Scope names are unique.
	// Wire name: 'scope'
	Scope string `json:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	// Wire name: 'scope_backend_type'
	ScopeBackendType ScopeBackendType `json:"scope_backend_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateScope) EncodeValues(key string, v *url.Values) error {
	pb, err := createScopeToPb(st)
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

func (st *CreateScope) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createScopePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createScopeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateScope) MarshalJSON() ([]byte, error) {
	pb, err := createScopeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string `json:"git_provider,omitempty"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CredentialInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := credentialInfoToPb(st)
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

func (st *CredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &credentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := credentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := credentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Delete struct {
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	// Wire name: 'recursive'
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Delete) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteToPb(st)
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

func (st *Delete) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Delete) MarshalJSON() ([]byte, error) {
	pb, err := deleteToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	// Wire name: 'principal'
	Principal string `json:"principal"`
	// The name of the scope to remove permissions from.
	// Wire name: 'scope'
	Scope string `json:"scope"`
}

func (st *DeleteAcl) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAclToPb(st)
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

func (st *DeleteAcl) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAclPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAclFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAcl) MarshalJSON() ([]byte, error) {
	pb, err := deleteAclToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" tf:"-"`
}

func (st *DeleteCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCredentialsRequestToPb(st)
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

func (st *DeleteCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	RepoId int64 `json:"-" tf:"-"`
}

func (st *DeleteRepoRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRepoRequestToPb(st)
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

func (st *DeleteRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteScope struct {
	// Name of the scope to delete.
	// Wire name: 'scope'
	Scope string `json:"scope"`
}

func (st *DeleteScope) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteScopeToPb(st)
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

func (st *DeleteScope) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteScopePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteScopeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteScope) MarshalJSON() ([]byte, error) {
	pb, err := deleteScopeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteSecret struct {
	// Name of the secret to delete.
	// Wire name: 'key'
	Key string `json:"key"`
	// The name of the scope that contains the secret to delete.
	// Wire name: 'scope'
	Scope string `json:"scope"`
}

func (st *DeleteSecret) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteSecretToPb(st)
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

func (st *DeleteSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSecret) MarshalJSON() ([]byte, error) {
	pb, err := deleteSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	Format ExportFormat `json:"-" tf:"-"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC`, `SOURCE`, and `AUTO` format.
	Path string `json:"-" tf:"-"`
}

func (st *ExportRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := exportRequestToPb(st)
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

func (st *ExportRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportRequest) MarshalJSON() ([]byte, error) {
	pb, err := exportRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The request field `direct_download` determines whether a JSON response or
// binary contents are returned by this endpoint.
type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	// Wire name: 'content'
	Content string `json:"content,omitempty"`
	// The file type of the exported file.
	// Wire name: 'file_type'
	FileType string `json:"file_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ExportResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := exportResponseToPb(st)
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

func (st *ExportResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportResponse) MarshalJSON() ([]byte, error) {
	pb, err := exportResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetAclRequest struct {
	// The principal to fetch ACL information for.
	Principal string `json:"-" tf:"-"`
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" tf:"-"`
}

func (st *GetAclRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAclRequestToPb(st)
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

func (st *GetAclRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAclRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAclRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAclRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAclRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" tf:"-"`
}

func (st *GetCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCredentialsRequestToPb(st)
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

func (st *GetCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetCredentialsResponse struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string `json:"git_provider,omitempty"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetCredentialsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getCredentialsResponseToPb(st)
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

func (st *GetCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" tf:"-"`
}

func (st *GetRepoPermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRepoPermissionLevelsRequestToPb(st)
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

func (st *GetRepoPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRepoPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRepoPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRepoPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRepoPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []RepoPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetRepoPermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getRepoPermissionLevelsResponseToPb(st)
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

func (st *GetRepoPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRepoPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRepoPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRepoPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRepoPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" tf:"-"`
}

func (st *GetRepoPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRepoPermissionsRequestToPb(st)
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

func (st *GetRepoPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRepoPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRepoPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRepoPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRepoPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRepoRequest struct {
	// ID of the Git folder (repo) object in the workspace.
	RepoId int64 `json:"-" tf:"-"`
}

func (st *GetRepoRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRepoRequestToPb(st)
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

func (st *GetRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetRepoResponse struct {
	// Branch that the local version of the repo is checked out to.
	// Wire name: 'branch'
	Branch string `json:"branch,omitempty"`
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	// Wire name: 'head_commit_id'
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'id'
	Id int64 `json:"id,omitempty"`
	// Path of the Git folder (repo) in the workspace.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// Git provider of the linked Git repository.
	// Wire name: 'provider'
	Provider string `json:"provider,omitempty"`
	// Sparse checkout settings for the Git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the linked Git repository.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetRepoResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getRepoResponseToPb(st)
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

func (st *GetRepoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRepoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRepoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRepoResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRepoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetSecretRequest struct {
	// The key to fetch secret for.
	Key string `json:"-" tf:"-"`
	// The name of the scope to fetch secret information from.
	Scope string `json:"-" tf:"-"`
}

func (st *GetSecretRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getSecretRequestToPb(st)
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

func (st *GetSecretRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSecretRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSecretRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSecretRequest) MarshalJSON() ([]byte, error) {
	pb, err := getSecretRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// The value of the secret in its byte representation.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetSecretResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getSecretResponseToPb(st)
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

func (st *GetSecretResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSecretResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSecretResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSecretResponse) MarshalJSON() ([]byte, error) {
	pb, err := getSecretResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" tf:"-"`
}

func (st *GetStatusRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStatusRequestToPb(st)
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

func (st *GetStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" tf:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" tf:"-"`
}

func (st *GetWorkspaceObjectPermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceObjectPermissionLevelsRequestToPb(st)
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

func (st *GetWorkspaceObjectPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceObjectPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceObjectPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceObjectPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceObjectPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []WorkspaceObjectPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetWorkspaceObjectPermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceObjectPermissionLevelsResponseToPb(st)
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

func (st *GetWorkspaceObjectPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceObjectPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceObjectPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceObjectPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceObjectPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" tf:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" tf:"-"`
}

func (st *GetWorkspaceObjectPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceObjectPermissionsRequestToPb(st)
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

func (st *GetWorkspaceObjectPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceObjectPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceObjectPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceObjectPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceObjectPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown. This parameter might be absent,
	// and instead a posted file is used.
	// Wire name: 'content'
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
	// Wire name: 'format'
	Format ImportFormat `json:"format,omitempty"`
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	// Wire name: 'language'
	Language Language `json:"language,omitempty"`
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	// Wire name: 'overwrite'
	Overwrite bool `json:"overwrite,omitempty"`
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	// Wire name: 'path'
	Path string `json:"path"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Import) EncodeValues(key string, v *url.Values) error {
	pb, err := importToPb(st)
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

func (st *Import) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &importPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := importFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Import) MarshalJSON() ([]byte, error) {
	pb, err := importToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	Scope string `json:"-" tf:"-"`
}

func (st *ListAclsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listAclsRequestToPb(st)
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

func (st *ListAclsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAclsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAclsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAclsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAclsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	// Wire name: 'items'
	Items []AclItem `json:"items,omitempty"`
}

func (st *ListAclsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listAclsResponseToPb(st)
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

func (st *ListAclsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAclsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAclsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAclsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAclsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListCredentialsResponse struct {
	// List of credentials.
	// Wire name: 'credentials'
	Credentials []CredentialInfo `json:"credentials,omitempty"`
}

func (st *ListCredentialsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listCredentialsResponseToPb(st)
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

func (st *ListCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	NextPageToken string `json:"-" tf:"-"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided or when provided an effectively empty prefix (`/` or
	// `/Workspace`) Git folders (repos) from `/Workspace/Repos` will be served.
	PathPrefix string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListReposRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listReposRequestToPb(st)
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

func (st *ListReposRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listReposRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listReposRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListReposRequest) MarshalJSON() ([]byte, error) {
	pb, err := listReposRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Git folders (repos).
	// Wire name: 'repos'
	Repos []RepoInfo `json:"repos,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListReposResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listReposResponseToPb(st)
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

func (st *ListReposResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listReposResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listReposResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListReposResponse) MarshalJSON() ([]byte, error) {
	pb, err := listReposResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListResponse struct {
	// List of objects.
	// Wire name: 'objects'
	Objects []ObjectInfo `json:"objects,omitempty"`
}

func (st *ListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listResponseToPb(st)
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

func (st *ListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListResponse) MarshalJSON() ([]byte, error) {
	pb, err := listResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListScopesResponse struct {
	// The available secret scopes.
	// Wire name: 'scopes'
	Scopes []SecretScope `json:"scopes,omitempty"`
}

func (st *ListScopesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listScopesResponseToPb(st)
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

func (st *ListScopesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listScopesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listScopesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListScopesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listScopesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	Scope string `json:"-" tf:"-"`
}

func (st *ListSecretsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listSecretsRequestToPb(st)
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

func (st *ListSecretsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSecretsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSecretsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSecretsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listSecretsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	// Wire name: 'secrets'
	Secrets []SecretMetadata `json:"secrets,omitempty"`
}

func (st *ListSecretsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listSecretsResponseToPb(st)
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

func (st *ListSecretsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listSecretsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listSecretsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListSecretsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listSecretsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter int64 `json:"-" tf:"-"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listWorkspaceRequestToPb(st)
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

func (st *ListWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := listWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	// Wire name: 'path'
	Path string `json:"path"`
}

func (st *Mkdirs) EncodeValues(key string, v *url.Values) error {
	pb, err := mkdirsToPb(st)
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

func (st *Mkdirs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mkdirsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := mkdirsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Mkdirs) MarshalJSON() ([]byte, error) {
	pb, err := mkdirsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The information of the object in workspace. It will be returned by “list“
// and “get-status“.
type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	// Wire name: 'language'
	Language Language `json:"language,omitempty"`
	// Only applicable to files, the last modified UTC timestamp.
	// Wire name: 'modified_at'
	ModifiedAt int64 `json:"modified_at,omitempty"`
	// Unique identifier for the object.
	// Wire name: 'object_id'
	ObjectId int64 `json:"object_id,omitempty"`
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	// Wire name: 'object_type'
	ObjectType ObjectType `json:"object_type,omitempty"`
	// The absolute path of the object.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	// Wire name: 'resource_id'
	ResourceId string `json:"resource_id,omitempty"`
	// Only applicable to files. The file size in bytes can be returned.
	// Wire name: 'size'
	Size int64 `json:"size,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ObjectInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := objectInfoToPb(st)
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

func (st *ObjectInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &objectInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := objectInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ObjectInfo) MarshalJSON() ([]byte, error) {
	pb, err := objectInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'permission'
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	// Wire name: 'principal'
	Principal string `json:"principal"`
	// The name of the scope to apply permissions to.
	// Wire name: 'scope'
	Scope string `json:"scope"`
}

func (st *PutAcl) EncodeValues(key string, v *url.Values) error {
	pb, err := putAclToPb(st)
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

func (st *PutAcl) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putAclPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putAclFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutAcl) MarshalJSON() ([]byte, error) {
	pb, err := putAclToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	// Wire name: 'bytes_value'
	BytesValue string `json:"bytes_value,omitempty"`
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string `json:"key"`
	// The name of the scope to which the secret will be associated with.
	// Wire name: 'scope'
	Scope string `json:"scope"`
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	// Wire name: 'string_value'
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PutSecret) EncodeValues(key string, v *url.Values) error {
	pb, err := putSecretToPb(st)
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

func (st *PutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutSecret) MarshalJSON() ([]byte, error) {
	pb, err := putSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RepoAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RepoAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := repoAccessControlRequestToPb(st)
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

func (st *RepoAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := repoAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RepoAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []RepoPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RepoAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := repoAccessControlResponseToPb(st)
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

func (st *RepoAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := repoAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Git folder (repo) information.
type RepoInfo struct {
	// Name of the current git branch of the git folder (repo).
	// Wire name: 'branch'
	Branch string `json:"branch,omitempty"`
	// Current git commit id of the git folder (repo).
	// Wire name: 'head_commit_id'
	HeadCommitId string `json:"head_commit_id,omitempty"`
	// Id of the git folder (repo) in the Workspace.
	// Wire name: 'id'
	Id int64 `json:"id,omitempty"`
	// Root path of the git folder (repo) in the Workspace.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`
	// Git provider of the remote git repository, e.g. `gitHub`.
	// Wire name: 'provider'
	Provider string `json:"provider,omitempty"`
	// Sparse checkout config for the git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	// URL of the remote git repository.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RepoInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := repoInfoToPb(st)
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

func (st *RepoInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoInfo) MarshalJSON() ([]byte, error) {
	pb, err := repoInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RepoPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RepoPermission) EncodeValues(key string, v *url.Values) error {
	pb, err := repoPermissionToPb(st)
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

func (st *RepoPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoPermission) MarshalJSON() ([]byte, error) {
	pb, err := repoPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RepoPermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := repoPermissionsToPb(st)
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

func (st *RepoPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoPermissions) MarshalJSON() ([]byte, error) {
	pb, err := repoPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RepoPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RepoPermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := repoPermissionsDescriptionToPb(st)
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

func (st *RepoPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := repoPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RepoPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlRequest `json:"access_control_list,omitempty"`
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" tf:"-"`
}

func (st *RepoPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := repoPermissionsRequestToPb(st)
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

func (st *RepoPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repoPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repoPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepoPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := repoPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type SecretMetadata struct {
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// The last updated timestamp (in milliseconds) for the secret.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SecretMetadata) EncodeValues(key string, v *url.Values) error {
	pb, err := secretMetadataToPb(st)
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

func (st *SecretMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &secretMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := secretMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SecretMetadata) MarshalJSON() ([]byte, error) {
	pb, err := secretMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SecretScope struct {
	// The type of secret scope backend.
	// Wire name: 'backend_type'
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	// Wire name: 'keyvault_metadata'
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
	// A unique name to identify the secret scope.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SecretScope) EncodeValues(key string, v *url.Values) error {
	pb, err := secretScopeToPb(st)
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

func (st *SecretScope) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &secretScopePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := secretScopeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SecretScope) MarshalJSON() ([]byte, error) {
	pb, err := secretScopeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	// Wire name: 'patterns'
	Patterns []string `json:"patterns,omitempty"`
}

func (st *SparseCheckout) EncodeValues(key string, v *url.Values) error {
	pb, err := sparseCheckoutToPb(st)
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

func (st *SparseCheckout) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparseCheckoutPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparseCheckoutFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparseCheckout) MarshalJSON() ([]byte, error) {
	pb, err := sparseCheckoutToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	// Wire name: 'patterns'
	Patterns []string `json:"patterns,omitempty"`
}

func (st *SparseCheckoutUpdate) EncodeValues(key string, v *url.Values) error {
	pb, err := sparseCheckoutUpdateToPb(st)
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

func (st *SparseCheckoutUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparseCheckoutUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparseCheckoutUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparseCheckoutUpdate) MarshalJSON() ([]byte, error) {
	pb, err := sparseCheckoutUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" tf:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'git_provider'
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	// Wire name: 'git_username'
	GitUsername string `json:"git_username,omitempty"`
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	// Wire name: 'personal_access_token'
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateCredentialsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCredentialsRequestToPb(st)
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

func (st *UpdateCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateRepoRequest struct {
	// Branch that the local version of the repo is checked out to.
	// Wire name: 'branch'
	Branch string `json:"branch,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	RepoId int64 `json:"-" tf:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckoutUpdate `json:"sparse_checkout,omitempty"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	// Wire name: 'tag'
	Tag string `json:"tag,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateRepoRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRepoRequestToPb(st)
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

func (st *UpdateRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *WorkspaceObjectAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceObjectAccessControlRequestToPb(st)
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

func (st *WorkspaceObjectAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceObjectAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceObjectAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceObjectAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := workspaceObjectAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WorkspaceObjectAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []WorkspaceObjectPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *WorkspaceObjectAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceObjectAccessControlResponseToPb(st)
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

func (st *WorkspaceObjectAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceObjectAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceObjectAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceObjectAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := workspaceObjectAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WorkspaceObjectPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *WorkspaceObjectPermission) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceObjectPermissionToPb(st)
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

func (st *WorkspaceObjectPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceObjectPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceObjectPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceObjectPermission) MarshalJSON() ([]byte, error) {
	pb, err := workspaceObjectPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []WorkspaceObjectAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *WorkspaceObjectPermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceObjectPermissionsToPb(st)
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

func (st *WorkspaceObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := workspaceObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WorkspaceObjectPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *WorkspaceObjectPermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceObjectPermissionsDescriptionToPb(st)
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

func (st *WorkspaceObjectPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceObjectPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceObjectPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceObjectPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := workspaceObjectPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WorkspaceObjectPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []WorkspaceObjectAccessControlRequest `json:"access_control_list,omitempty"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" tf:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" tf:"-"`
}

func (st *WorkspaceObjectPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceObjectPermissionsRequestToPb(st)
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

func (st *WorkspaceObjectPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceObjectPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceObjectPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceObjectPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := workspaceObjectPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
