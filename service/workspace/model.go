// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/workspace/workspacepb"
)

// An item representing an ACL rule applied to the given principal (user or
// group) on the associated scope point.
type AclItem struct {
	// The permission level applied to the principal.
	// Wire name: 'permission'
	Permission AclPermission ``
	// The principal in which the permission is applied.
	// Wire name: 'principal'
	Principal string ``
}

func AclItemToPb(st *AclItem) (*workspacepb.AclItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.AclItemPb{}
	permissionPb, err := AclPermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}
	pb.Principal = st.Principal

	return pb, nil
}

func AclItemFromPb(pb *workspacepb.AclItemPb) (*AclItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AclItem{}
	permissionField, err := AclPermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}
	st.Principal = pb.Principal

	return st, nil
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

func AclPermissionToPb(st *AclPermission) (*workspacepb.AclPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.AclPermissionPb(*st)
	return &pb, nil
}

func AclPermissionFromPb(pb *workspacepb.AclPermissionPb) (*AclPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AclPermission(*pb)
	return &st, nil
}

// The metadata of the Azure KeyVault for a secret scope of type
// `AZURE_KEYVAULT`
type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	// Wire name: 'dns_name'
	DnsName string ``
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	// Wire name: 'resource_id'
	ResourceId string ``
}

func AzureKeyVaultSecretScopeMetadataToPb(st *AzureKeyVaultSecretScopeMetadata) (*workspacepb.AzureKeyVaultSecretScopeMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.AzureKeyVaultSecretScopeMetadataPb{}
	pb.DnsName = st.DnsName
	pb.ResourceId = st.ResourceId

	return pb, nil
}

func AzureKeyVaultSecretScopeMetadataFromPb(pb *workspacepb.AzureKeyVaultSecretScopeMetadataPb) (*AzureKeyVaultSecretScopeMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureKeyVaultSecretScopeMetadata{}
	st.DnsName = pb.DnsName
	st.ResourceId = pb.ResourceId

	return st, nil
}

type CreateCredentialsRequest struct {
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'git_provider'
	GitProvider string ``
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	// Wire name: 'git_username'
	GitUsername string ``
	// if the credential is the default for the given provider
	// Wire name: 'is_default_for_provider'
	IsDefaultForProvider bool ``
	// the name of the git credential, used for identification and ease of
	// lookup
	// Wire name: 'name'
	Name string ``
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	// Wire name: 'personal_access_token'
	PersonalAccessToken string   ``
	ForceSendFields     []string `tf:"-"`
}

func (s *CreateCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateCredentialsRequestToPb(st *CreateCredentialsRequest) (*workspacepb.CreateCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.CreateCredentialsRequestPb{}
	pb.GitProvider = st.GitProvider
	pb.GitUsername = st.GitUsername
	pb.IsDefaultForProvider = st.IsDefaultForProvider
	pb.Name = st.Name
	pb.PersonalAccessToken = st.PersonalAccessToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateCredentialsRequestFromPb(pb *workspacepb.CreateCredentialsRequestPb) (*CreateCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialsRequest{}
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.IsDefaultForProvider = pb.IsDefaultForProvider
	st.Name = pb.Name
	st.PersonalAccessToken = pb.PersonalAccessToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64 ``
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string ``
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string ``
	// if the credential is the default for the given provider
	// Wire name: 'is_default_for_provider'
	IsDefaultForProvider bool ``
	// the name of the git credential, used for identification and ease of
	// lookup
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateCredentialsResponseToPb(st *CreateCredentialsResponse) (*workspacepb.CreateCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.CreateCredentialsResponsePb{}
	pb.CredentialId = st.CredentialId
	pb.GitProvider = st.GitProvider
	pb.GitUsername = st.GitUsername
	pb.IsDefaultForProvider = st.IsDefaultForProvider
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateCredentialsResponseFromPb(pb *workspacepb.CreateCredentialsResponsePb) (*CreateCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialsResponse{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.IsDefaultForProvider = pb.IsDefaultForProvider
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateRepoRequest struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in `/Repos`, path must be in
	// the format `/Repos/{folder}/{repo-name}`.
	// Wire name: 'path'
	Path string ``
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'provider'
	Provider string ``
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout ``
	// URL of the Git repository to be linked.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateRepoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRepoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateRepoRequestToPb(st *CreateRepoRequest) (*workspacepb.CreateRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.CreateRepoRequestPb{}
	pb.Path = st.Path
	pb.Provider = st.Provider
	sparseCheckoutPb, err := SparseCheckoutToPb(st.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutPb != nil {
		pb.SparseCheckout = sparseCheckoutPb
	}
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateRepoRequestFromPb(pb *workspacepb.CreateRepoRequestPb) (*CreateRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRepoRequest{}
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := SparseCheckoutFromPb(pb.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutField != nil {
		st.SparseCheckout = sparseCheckoutField
	}
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateRepoResponse struct {
	// Branch that the Git folder (repo) is checked out to.
	// Wire name: 'branch'
	Branch string ``
	// SHA-1 hash representing the commit ID of the current HEAD of the Git
	// folder (repo).
	// Wire name: 'head_commit_id'
	HeadCommitId string ``
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'id'
	Id int64 ``
	// Path of the Git folder (repo) in the workspace.
	// Wire name: 'path'
	Path string ``
	// Git provider of the linked Git repository.
	// Wire name: 'provider'
	Provider string ``
	// Sparse checkout settings for the Git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout ``
	// URL of the linked Git repository.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateRepoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateRepoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateRepoResponseToPb(st *CreateRepoResponse) (*workspacepb.CreateRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.CreateRepoResponsePb{}
	pb.Branch = st.Branch
	pb.HeadCommitId = st.HeadCommitId
	pb.Id = st.Id
	pb.Path = st.Path
	pb.Provider = st.Provider
	sparseCheckoutPb, err := SparseCheckoutToPb(st.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutPb != nil {
		pb.SparseCheckout = sparseCheckoutPb
	}
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateRepoResponseFromPb(pb *workspacepb.CreateRepoResponsePb) (*CreateRepoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRepoResponse{}
	st.Branch = pb.Branch
	st.HeadCommitId = pb.HeadCommitId
	st.Id = pb.Id
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := SparseCheckoutFromPb(pb.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutField != nil {
		st.SparseCheckout = sparseCheckoutField
	}
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateScope struct {
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	// Wire name: 'backend_azure_keyvault'
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata ``
	// The principal that is initially granted ``MANAGE`` permission to the
	// created scope.
	// Wire name: 'initial_manage_principal'
	InitialManagePrincipal string ``
	// Scope name requested by the user. Scope names are unique.
	// Wire name: 'scope'
	Scope string ``
	// The backend type the scope will be created with. If not specified, will
	// default to ``DATABRICKS``
	// Wire name: 'scope_backend_type'
	ScopeBackendType ScopeBackendType ``
	ForceSendFields  []string         `tf:"-"`
}

func (s *CreateScope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateScope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateScopeToPb(st *CreateScope) (*workspacepb.CreateScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.CreateScopePb{}
	backendAzureKeyvaultPb, err := AzureKeyVaultSecretScopeMetadataToPb(st.BackendAzureKeyvault)
	if err != nil {
		return nil, err
	}
	if backendAzureKeyvaultPb != nil {
		pb.BackendAzureKeyvault = backendAzureKeyvaultPb
	}
	pb.InitialManagePrincipal = st.InitialManagePrincipal
	pb.Scope = st.Scope
	scopeBackendTypePb, err := ScopeBackendTypeToPb(&st.ScopeBackendType)
	if err != nil {
		return nil, err
	}
	if scopeBackendTypePb != nil {
		pb.ScopeBackendType = *scopeBackendTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateScopeFromPb(pb *workspacepb.CreateScopePb) (*CreateScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScope{}
	backendAzureKeyvaultField, err := AzureKeyVaultSecretScopeMetadataFromPb(pb.BackendAzureKeyvault)
	if err != nil {
		return nil, err
	}
	if backendAzureKeyvaultField != nil {
		st.BackendAzureKeyvault = backendAzureKeyvaultField
	}
	st.InitialManagePrincipal = pb.InitialManagePrincipal
	st.Scope = pb.Scope
	scopeBackendTypeField, err := ScopeBackendTypeFromPb(&pb.ScopeBackendType)
	if err != nil {
		return nil, err
	}
	if scopeBackendTypeField != nil {
		st.ScopeBackendType = *scopeBackendTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64 ``
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string ``
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string ``
	// if the credential is the default for the given provider
	// Wire name: 'is_default_for_provider'
	IsDefaultForProvider bool ``
	// the name of the git credential, used for identification and ease of
	// lookup
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CredentialInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CredentialInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CredentialInfoToPb(st *CredentialInfo) (*workspacepb.CredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.CredentialInfoPb{}
	pb.CredentialId = st.CredentialId
	pb.GitProvider = st.GitProvider
	pb.GitUsername = st.GitUsername
	pb.IsDefaultForProvider = st.IsDefaultForProvider
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CredentialInfoFromPb(pb *workspacepb.CredentialInfoPb) (*CredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialInfo{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.IsDefaultForProvider = pb.IsDefaultForProvider
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Delete struct {
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string ``
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	// Wire name: 'recursive'
	Recursive       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *Delete) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Delete) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DeleteToPb(st *Delete) (*workspacepb.DeletePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeletePb{}
	pb.Path = st.Path
	pb.Recursive = st.Recursive

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DeleteFromPb(pb *workspacepb.DeletePb) (*Delete, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Delete{}
	st.Path = pb.Path
	st.Recursive = pb.Recursive

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	// Wire name: 'principal'
	Principal string ``
	// The name of the scope to remove permissions from.
	// Wire name: 'scope'
	Scope string ``
}

func DeleteAclToPb(st *DeleteAcl) (*workspacepb.DeleteAclPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeleteAclPb{}
	pb.Principal = st.Principal
	pb.Scope = st.Scope

	return pb, nil
}

func DeleteAclFromPb(pb *workspacepb.DeleteAclPb) (*DeleteAcl, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAcl{}
	st.Principal = pb.Principal
	st.Scope = pb.Scope

	return st, nil
}

type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	// Wire name: 'credential_id'
	CredentialId int64 `tf:"-"`
}

func DeleteCredentialsRequestToPb(st *DeleteCredentialsRequest) (*workspacepb.DeleteCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeleteCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	return pb, nil
}

func DeleteCredentialsRequestFromPb(pb *workspacepb.DeleteCredentialsRequestPb) (*DeleteCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialsRequest{}
	st.CredentialId = pb.CredentialId

	return st, nil
}

type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	// Wire name: 'repo_id'
	RepoId int64 `tf:"-"`
}

func DeleteRepoRequestToPb(st *DeleteRepoRequest) (*workspacepb.DeleteRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeleteRepoRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

func DeleteRepoRequestFromPb(pb *workspacepb.DeleteRepoRequestPb) (*DeleteRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRepoRequest{}
	st.RepoId = pb.RepoId

	return st, nil
}

type DeleteScope struct {
	// Name of the scope to delete.
	// Wire name: 'scope'
	Scope string ``
}

func DeleteScopeToPb(st *DeleteScope) (*workspacepb.DeleteScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeleteScopePb{}
	pb.Scope = st.Scope

	return pb, nil
}

func DeleteScopeFromPb(pb *workspacepb.DeleteScopePb) (*DeleteScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteScope{}
	st.Scope = pb.Scope

	return st, nil
}

type DeleteSecret struct {
	// Name of the secret to delete.
	// Wire name: 'key'
	Key string ``
	// The name of the scope that contains the secret to delete.
	// Wire name: 'scope'
	Scope string ``
}

func DeleteSecretToPb(st *DeleteSecret) (*workspacepb.DeleteSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeleteSecretPb{}
	pb.Key = st.Key
	pb.Scope = st.Scope

	return pb, nil
}

func DeleteSecretFromPb(pb *workspacepb.DeleteSecretPb) (*DeleteSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSecret{}
	st.Key = pb.Key
	st.Scope = pb.Scope

	return st, nil
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

func ExportFormatToPb(st *ExportFormat) (*workspacepb.ExportFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.ExportFormatPb(*st)
	return &pb, nil
}

func ExportFormatFromPb(pb *workspacepb.ExportFormatPb) (*ExportFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExportFormat(*pb)
	return &st, nil
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
	// Wire name: 'format'
	Format ExportFormat `tf:"-"`
	// The absolute path of the object or directory. Exporting a directory is
	// only supported for the `DBC`, `SOURCE`, and `AUTO` format.
	// Wire name: 'path'
	Path string `tf:"-"`
}

func ExportRequestToPb(st *ExportRequest) (*workspacepb.ExportRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ExportRequestPb{}
	formatPb, err := ExportFormatToPb(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}
	pb.Path = st.Path

	return pb, nil
}

func ExportRequestFromPb(pb *workspacepb.ExportRequestPb) (*ExportRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRequest{}
	formatField, err := ExportFormatFromPb(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	st.Path = pb.Path

	return st, nil
}

// The request field `direct_download` determines whether a JSON response or
// binary contents are returned by this endpoint.
type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	// Wire name: 'content'
	Content string ``
	// The file type of the exported file.
	// Wire name: 'file_type'
	FileType        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExportResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExportResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExportResponseToPb(st *ExportResponse) (*workspacepb.ExportResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ExportResponsePb{}
	pb.Content = st.Content
	pb.FileType = st.FileType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExportResponseFromPb(pb *workspacepb.ExportResponsePb) (*ExportResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportResponse{}
	st.Content = pb.Content
	st.FileType = pb.FileType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetAclRequest struct {
	// The principal to fetch ACL information for.
	// Wire name: 'principal'
	Principal string `tf:"-"`
	// The name of the scope to fetch ACL information from.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func GetAclRequestToPb(st *GetAclRequest) (*workspacepb.GetAclRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetAclRequestPb{}
	pb.Principal = st.Principal
	pb.Scope = st.Scope

	return pb, nil
}

func GetAclRequestFromPb(pb *workspacepb.GetAclRequestPb) (*GetAclRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAclRequest{}
	st.Principal = pb.Principal
	st.Scope = pb.Scope

	return st, nil
}

type GetCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	// Wire name: 'credential_id'
	CredentialId int64 `tf:"-"`
}

func GetCredentialsRequestToPb(st *GetCredentialsRequest) (*workspacepb.GetCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	return pb, nil
}

func GetCredentialsRequestFromPb(pb *workspacepb.GetCredentialsRequestPb) (*GetCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsRequest{}
	st.CredentialId = pb.CredentialId

	return st, nil
}

type GetCredentialsResponse struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64 ``
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string ``
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string ``
	// if the credential is the default for the given provider
	// Wire name: 'is_default_for_provider'
	IsDefaultForProvider bool ``
	// the name of the git credential, used for identification and ease of
	// lookup
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetCredentialsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCredentialsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetCredentialsResponseToPb(st *GetCredentialsResponse) (*workspacepb.GetCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetCredentialsResponsePb{}
	pb.CredentialId = st.CredentialId
	pb.GitProvider = st.GitProvider
	pb.GitUsername = st.GitUsername
	pb.IsDefaultForProvider = st.IsDefaultForProvider
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetCredentialsResponseFromPb(pb *workspacepb.GetCredentialsResponsePb) (*GetCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsResponse{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.IsDefaultForProvider = pb.IsDefaultForProvider
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func GetRepoPermissionLevelsRequestToPb(st *GetRepoPermissionLevelsRequest) (*workspacepb.GetRepoPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetRepoPermissionLevelsRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

func GetRepoPermissionLevelsRequestFromPb(pb *workspacepb.GetRepoPermissionLevelsRequestPb) (*GetRepoPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoPermissionLevelsRequest{}
	st.RepoId = pb.RepoId

	return st, nil
}

type GetRepoPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []RepoPermissionsDescription ``
}

func GetRepoPermissionLevelsResponseToPb(st *GetRepoPermissionLevelsResponse) (*workspacepb.GetRepoPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetRepoPermissionLevelsResponsePb{}

	var permissionLevelsPb []workspacepb.RepoPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := RepoPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetRepoPermissionLevelsResponseFromPb(pb *workspacepb.GetRepoPermissionLevelsResponsePb) (*GetRepoPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoPermissionLevelsResponse{}

	var permissionLevelsField []RepoPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := RepoPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func GetRepoPermissionsRequestToPb(st *GetRepoPermissionsRequest) (*workspacepb.GetRepoPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetRepoPermissionsRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

func GetRepoPermissionsRequestFromPb(pb *workspacepb.GetRepoPermissionsRequestPb) (*GetRepoPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoPermissionsRequest{}
	st.RepoId = pb.RepoId

	return st, nil
}

type GetRepoRequest struct {
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'repo_id'
	RepoId int64 `tf:"-"`
}

func GetRepoRequestToPb(st *GetRepoRequest) (*workspacepb.GetRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetRepoRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

func GetRepoRequestFromPb(pb *workspacepb.GetRepoRequestPb) (*GetRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoRequest{}
	st.RepoId = pb.RepoId

	return st, nil
}

type GetRepoResponse struct {
	// Branch that the local version of the repo is checked out to.
	// Wire name: 'branch'
	Branch string ``
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	// Wire name: 'head_commit_id'
	HeadCommitId string ``
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'id'
	Id int64 ``
	// Path of the Git folder (repo) in the workspace.
	// Wire name: 'path'
	Path string ``
	// Git provider of the linked Git repository.
	// Wire name: 'provider'
	Provider string ``
	// Sparse checkout settings for the Git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout ``
	// URL of the linked Git repository.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetRepoResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRepoResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetRepoResponseToPb(st *GetRepoResponse) (*workspacepb.GetRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetRepoResponsePb{}
	pb.Branch = st.Branch
	pb.HeadCommitId = st.HeadCommitId
	pb.Id = st.Id
	pb.Path = st.Path
	pb.Provider = st.Provider
	sparseCheckoutPb, err := SparseCheckoutToPb(st.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutPb != nil {
		pb.SparseCheckout = sparseCheckoutPb
	}
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetRepoResponseFromPb(pb *workspacepb.GetRepoResponsePb) (*GetRepoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoResponse{}
	st.Branch = pb.Branch
	st.HeadCommitId = pb.HeadCommitId
	st.Id = pb.Id
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := SparseCheckoutFromPb(pb.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutField != nil {
		st.SparseCheckout = sparseCheckoutField
	}
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetSecretRequest struct {
	// Name of the secret to fetch value information.
	// Wire name: 'key'
	Key string `tf:"-"`
	// The name of the scope that contains the secret.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func GetSecretRequestToPb(st *GetSecretRequest) (*workspacepb.GetSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetSecretRequestPb{}
	pb.Key = st.Key
	pb.Scope = st.Scope

	return pb, nil
}

func GetSecretRequestFromPb(pb *workspacepb.GetSecretRequestPb) (*GetSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSecretRequest{}
	st.Key = pb.Key
	st.Scope = pb.Scope

	return st, nil
}

type GetSecretResponse struct {
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string ``
	// The value of the secret in its byte representation.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GetSecretResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetSecretResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetSecretResponseToPb(st *GetSecretResponse) (*workspacepb.GetSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetSecretResponsePb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetSecretResponseFromPb(pb *workspacepb.GetSecretResponsePb) (*GetSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSecretResponse{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string `tf:"-"`
}

func GetStatusRequestToPb(st *GetStatusRequest) (*workspacepb.GetStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetStatusRequestPb{}
	pb.Path = st.Path

	return pb, nil
}

func GetStatusRequestFromPb(pb *workspacepb.GetStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	st.Path = pb.Path

	return st, nil
}

type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	// Wire name: 'workspace_object_id'
	WorkspaceObjectId string `tf:"-"`
	// The workspace object type for which to get or manage permissions.
	// Wire name: 'workspace_object_type'
	WorkspaceObjectType string `tf:"-"`
}

func GetWorkspaceObjectPermissionLevelsRequestToPb(st *GetWorkspaceObjectPermissionLevelsRequest) (*workspacepb.GetWorkspaceObjectPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetWorkspaceObjectPermissionLevelsRequestPb{}
	pb.WorkspaceObjectId = st.WorkspaceObjectId
	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
}

func GetWorkspaceObjectPermissionLevelsRequestFromPb(pb *workspacepb.GetWorkspaceObjectPermissionLevelsRequestPb) (*GetWorkspaceObjectPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceObjectPermissionLevelsRequest{}
	st.WorkspaceObjectId = pb.WorkspaceObjectId
	st.WorkspaceObjectType = pb.WorkspaceObjectType

	return st, nil
}

type GetWorkspaceObjectPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []WorkspaceObjectPermissionsDescription ``
}

func GetWorkspaceObjectPermissionLevelsResponseToPb(st *GetWorkspaceObjectPermissionLevelsResponse) (*workspacepb.GetWorkspaceObjectPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetWorkspaceObjectPermissionLevelsResponsePb{}

	var permissionLevelsPb []workspacepb.WorkspaceObjectPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := WorkspaceObjectPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetWorkspaceObjectPermissionLevelsResponseFromPb(pb *workspacepb.GetWorkspaceObjectPermissionLevelsResponsePb) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceObjectPermissionLevelsResponse{}

	var permissionLevelsField []WorkspaceObjectPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := WorkspaceObjectPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	// Wire name: 'workspace_object_id'
	WorkspaceObjectId string `tf:"-"`
	// The workspace object type for which to get or manage permissions.
	// Wire name: 'workspace_object_type'
	WorkspaceObjectType string `tf:"-"`
}

func GetWorkspaceObjectPermissionsRequestToPb(st *GetWorkspaceObjectPermissionsRequest) (*workspacepb.GetWorkspaceObjectPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetWorkspaceObjectPermissionsRequestPb{}
	pb.WorkspaceObjectId = st.WorkspaceObjectId
	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
}

func GetWorkspaceObjectPermissionsRequestFromPb(pb *workspacepb.GetWorkspaceObjectPermissionsRequestPb) (*GetWorkspaceObjectPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceObjectPermissionsRequest{}
	st.WorkspaceObjectId = pb.WorkspaceObjectId
	st.WorkspaceObjectType = pb.WorkspaceObjectType

	return st, nil
}

type Import struct {
	// The base64-encoded content. This has a limit of 10 MB.
	//
	// If the limit (10MB) is exceeded, exception with error code
	// **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown. This parameter might be absent,
	// and instead a posted file is used.
	// Wire name: 'content'
	Content string ``
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
	Format ImportFormat ``
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	// Wire name: 'language'
	Language Language ``
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	// Wire name: 'overwrite'
	Overwrite bool ``
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Import) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Import) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ImportToPb(st *Import) (*workspacepb.ImportPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ImportPb{}
	pb.Content = st.Content
	formatPb, err := ImportFormatToPb(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}
	languagePb, err := LanguageToPb(&st.Language)
	if err != nil {
		return nil, err
	}
	if languagePb != nil {
		pb.Language = *languagePb
	}
	pb.Overwrite = st.Overwrite
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ImportFromPb(pb *workspacepb.ImportPb) (*Import, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Import{}
	st.Content = pb.Content
	formatField, err := ImportFormatFromPb(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	languageField, err := LanguageFromPb(&pb.Language)
	if err != nil {
		return nil, err
	}
	if languageField != nil {
		st.Language = *languageField
	}
	st.Overwrite = pb.Overwrite
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ImportFormatToPb(st *ImportFormat) (*workspacepb.ImportFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.ImportFormatPb(*st)
	return &pb, nil
}

func ImportFormatFromPb(pb *workspacepb.ImportFormatPb) (*ImportFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := ImportFormat(*pb)
	return &st, nil
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

func LanguageToPb(st *Language) (*workspacepb.LanguagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.LanguagePb(*st)
	return &pb, nil
}

func LanguageFromPb(pb *workspacepb.LanguagePb) (*Language, error) {
	if pb == nil {
		return nil, nil
	}
	st := Language(*pb)
	return &st, nil
}

type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func ListAclsRequestToPb(st *ListAclsRequest) (*workspacepb.ListAclsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListAclsRequestPb{}
	pb.Scope = st.Scope

	return pb, nil
}

func ListAclsRequestFromPb(pb *workspacepb.ListAclsRequestPb) (*ListAclsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAclsRequest{}
	st.Scope = pb.Scope

	return st, nil
}

type ListAclsResponse struct {
	// The associated ACLs rule applied to principals in the given scope.
	// Wire name: 'items'
	Items []AclItem ``
}

func ListAclsResponseToPb(st *ListAclsResponse) (*workspacepb.ListAclsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListAclsResponsePb{}

	var itemsPb []workspacepb.AclItemPb
	for _, item := range st.Items {
		itemPb, err := AclItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			itemsPb = append(itemsPb, *itemPb)
		}
	}
	pb.Items = itemsPb

	return pb, nil
}

func ListAclsResponseFromPb(pb *workspacepb.ListAclsResponsePb) (*ListAclsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAclsResponse{}

	var itemsField []AclItem
	for _, itemPb := range pb.Items {
		item, err := AclItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			itemsField = append(itemsField, *item)
		}
	}
	st.Items = itemsField

	return st, nil
}

type ListCredentialsResponse struct {
	// List of credentials.
	// Wire name: 'credentials'
	Credentials []CredentialInfo ``
}

func ListCredentialsResponseToPb(st *ListCredentialsResponse) (*workspacepb.ListCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListCredentialsResponsePb{}

	var credentialsPb []workspacepb.CredentialInfoPb
	for _, item := range st.Credentials {
		itemPb, err := CredentialInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			credentialsPb = append(credentialsPb, *itemPb)
		}
	}
	pb.Credentials = credentialsPb

	return pb, nil
}

func ListCredentialsResponseFromPb(pb *workspacepb.ListCredentialsResponsePb) (*ListCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsResponse{}

	var credentialsField []CredentialInfo
	for _, itemPb := range pb.Credentials {
		item, err := CredentialInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			credentialsField = append(credentialsField, *item)
		}
	}
	st.Credentials = credentialsField

	return st, nil
}

type ListReposRequest struct {
	// Token used to get the next page of results. If not specified, returns the
	// first page of results as well as a next page token if there are more
	// results.
	// Wire name: 'next_page_token'
	NextPageToken string `tf:"-"`
	// Filters repos that have paths starting with the given path prefix. If not
	// provided or when provided an effectively empty prefix (`/` or
	// `/Workspace`) Git folders (repos) from `/Workspace/Repos` will be served.
	// Wire name: 'path_prefix'
	PathPrefix      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListReposRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListReposRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListReposRequestToPb(st *ListReposRequest) (*workspacepb.ListReposRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListReposRequestPb{}
	pb.NextPageToken = st.NextPageToken
	pb.PathPrefix = st.PathPrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListReposRequestFromPb(pb *workspacepb.ListReposRequestPb) (*ListReposRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListReposRequest{}
	st.NextPageToken = pb.NextPageToken
	st.PathPrefix = pb.PathPrefix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// List of Git folders (repos).
	// Wire name: 'repos'
	Repos           []RepoInfo ``
	ForceSendFields []string   `tf:"-"`
}

func (s *ListReposResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListReposResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListReposResponseToPb(st *ListReposResponse) (*workspacepb.ListReposResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListReposResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var reposPb []workspacepb.RepoInfoPb
	for _, item := range st.Repos {
		itemPb, err := RepoInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			reposPb = append(reposPb, *itemPb)
		}
	}
	pb.Repos = reposPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListReposResponseFromPb(pb *workspacepb.ListReposResponsePb) (*ListReposResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListReposResponse{}
	st.NextPageToken = pb.NextPageToken

	var reposField []RepoInfo
	for _, itemPb := range pb.Repos {
		item, err := RepoInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			reposField = append(reposField, *item)
		}
	}
	st.Repos = reposField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListResponse struct {
	// List of objects.
	// Wire name: 'objects'
	Objects []ObjectInfo ``
}

func ListResponseToPb(st *ListResponse) (*workspacepb.ListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListResponsePb{}

	var objectsPb []workspacepb.ObjectInfoPb
	for _, item := range st.Objects {
		itemPb, err := ObjectInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			objectsPb = append(objectsPb, *itemPb)
		}
	}
	pb.Objects = objectsPb

	return pb, nil
}

func ListResponseFromPb(pb *workspacepb.ListResponsePb) (*ListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListResponse{}

	var objectsField []ObjectInfo
	for _, itemPb := range pb.Objects {
		item, err := ObjectInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			objectsField = append(objectsField, *item)
		}
	}
	st.Objects = objectsField

	return st, nil
}

type ListScopesResponse struct {
	// The available secret scopes.
	// Wire name: 'scopes'
	Scopes []SecretScope ``
}

func ListScopesResponseToPb(st *ListScopesResponse) (*workspacepb.ListScopesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListScopesResponsePb{}

	var scopesPb []workspacepb.SecretScopePb
	for _, item := range st.Scopes {
		itemPb, err := SecretScopeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	return pb, nil
}

func ListScopesResponseFromPb(pb *workspacepb.ListScopesResponsePb) (*ListScopesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListScopesResponse{}

	var scopesField []SecretScope
	for _, itemPb := range pb.Scopes {
		item, err := SecretScopeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			scopesField = append(scopesField, *item)
		}
	}
	st.Scopes = scopesField

	return st, nil
}

type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func ListSecretsRequestToPb(st *ListSecretsRequest) (*workspacepb.ListSecretsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListSecretsRequestPb{}
	pb.Scope = st.Scope

	return pb, nil
}

func ListSecretsRequestFromPb(pb *workspacepb.ListSecretsRequestPb) (*ListSecretsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSecretsRequest{}
	st.Scope = pb.Scope

	return st, nil
}

type ListSecretsResponse struct {
	// Metadata information of all secrets contained within the given scope.
	// Wire name: 'secrets'
	Secrets []SecretMetadata ``
}

func ListSecretsResponseToPb(st *ListSecretsResponse) (*workspacepb.ListSecretsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListSecretsResponsePb{}

	var secretsPb []workspacepb.SecretMetadataPb
	for _, item := range st.Secrets {
		itemPb, err := SecretMetadataToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			secretsPb = append(secretsPb, *itemPb)
		}
	}
	pb.Secrets = secretsPb

	return pb, nil
}

func ListSecretsResponseFromPb(pb *workspacepb.ListSecretsResponsePb) (*ListSecretsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSecretsResponse{}

	var secretsField []SecretMetadata
	for _, itemPb := range pb.Secrets {
		item, err := SecretMetadataFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			secretsField = append(secretsField, *item)
		}
	}
	st.Secrets = secretsField

	return st, nil
}

type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	// Wire name: 'notebooks_modified_after'
	NotebooksModifiedAfter int64 `tf:"-"`
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListWorkspaceRequestToPb(st *ListWorkspaceRequest) (*workspacepb.ListWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListWorkspaceRequestPb{}
	pb.NotebooksModifiedAfter = st.NotebooksModifiedAfter
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListWorkspaceRequestFromPb(pb *workspacepb.ListWorkspaceRequestPb) (*ListWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWorkspaceRequest{}
	st.NotebooksModifiedAfter = pb.NotebooksModifiedAfter
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	// Wire name: 'path'
	Path string ``
}

func MkdirsToPb(st *Mkdirs) (*workspacepb.MkdirsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.MkdirsPb{}
	pb.Path = st.Path

	return pb, nil
}

func MkdirsFromPb(pb *workspacepb.MkdirsPb) (*Mkdirs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Mkdirs{}
	st.Path = pb.Path

	return st, nil
}

// The information of the object in workspace. It will be returned by “list“
// and “get-status“.
type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	// Wire name: 'created_at'
	CreatedAt int64 ``
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	// Wire name: 'language'
	Language Language ``
	// Only applicable to files, the last modified UTC timestamp.
	// Wire name: 'modified_at'
	ModifiedAt int64 ``
	// Unique identifier for the object.
	// Wire name: 'object_id'
	ObjectId int64 ``
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	// Wire name: 'object_type'
	ObjectType ObjectType ``
	// The absolute path of the object.
	// Wire name: 'path'
	Path string ``
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	// Wire name: 'resource_id'
	ResourceId string ``
	// Only applicable to files. The file size in bytes can be returned.
	// Wire name: 'size'
	Size            int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *ObjectInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ObjectInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ObjectInfoToPb(st *ObjectInfo) (*workspacepb.ObjectInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ObjectInfoPb{}
	pb.CreatedAt = st.CreatedAt
	languagePb, err := LanguageToPb(&st.Language)
	if err != nil {
		return nil, err
	}
	if languagePb != nil {
		pb.Language = *languagePb
	}
	pb.ModifiedAt = st.ModifiedAt
	pb.ObjectId = st.ObjectId
	objectTypePb, err := ObjectTypeToPb(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}
	pb.Path = st.Path
	pb.ResourceId = st.ResourceId
	pb.Size = st.Size

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ObjectInfoFromPb(pb *workspacepb.ObjectInfoPb) (*ObjectInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ObjectInfo{}
	st.CreatedAt = pb.CreatedAt
	languageField, err := LanguageFromPb(&pb.Language)
	if err != nil {
		return nil, err
	}
	if languageField != nil {
		st.Language = *languageField
	}
	st.ModifiedAt = pb.ModifiedAt
	st.ObjectId = pb.ObjectId
	objectTypeField, err := ObjectTypeFromPb(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}
	st.Path = pb.Path
	st.ResourceId = pb.ResourceId
	st.Size = pb.Size

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ObjectTypeToPb(st *ObjectType) (*workspacepb.ObjectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.ObjectTypePb(*st)
	return &pb, nil
}

func ObjectTypeFromPb(pb *workspacepb.ObjectTypePb) (*ObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ObjectType(*pb)
	return &st, nil
}

type PutAcl struct {
	// The permission level applied to the principal.
	// Wire name: 'permission'
	Permission AclPermission ``
	// The principal in which the permission is applied.
	// Wire name: 'principal'
	Principal string ``
	// The name of the scope to apply permissions to.
	// Wire name: 'scope'
	Scope string ``
}

func PutAclToPb(st *PutAcl) (*workspacepb.PutAclPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.PutAclPb{}
	permissionPb, err := AclPermissionToPb(&st.Permission)
	if err != nil {
		return nil, err
	}
	if permissionPb != nil {
		pb.Permission = *permissionPb
	}
	pb.Principal = st.Principal
	pb.Scope = st.Scope

	return pb, nil
}

func PutAclFromPb(pb *workspacepb.PutAclPb) (*PutAcl, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAcl{}
	permissionField, err := AclPermissionFromPb(&pb.Permission)
	if err != nil {
		return nil, err
	}
	if permissionField != nil {
		st.Permission = *permissionField
	}
	st.Principal = pb.Principal
	st.Scope = pb.Scope

	return st, nil
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	// Wire name: 'bytes_value'
	BytesValue string ``
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string ``
	// The name of the scope to which the secret will be associated with.
	// Wire name: 'scope'
	Scope string ``
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	// Wire name: 'string_value'
	StringValue     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PutSecret) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PutSecret) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PutSecretToPb(st *PutSecret) (*workspacepb.PutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.PutSecretPb{}
	pb.BytesValue = st.BytesValue
	pb.Key = st.Key
	pb.Scope = st.Scope
	pb.StringValue = st.StringValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PutSecretFromPb(pb *workspacepb.PutSecretPb) (*PutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutSecret{}
	st.BytesValue = pb.BytesValue
	st.Key = pb.Key
	st.Scope = pb.Scope
	st.StringValue = pb.StringValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RepoAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RepoAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RepoAccessControlRequestToPb(st *RepoAccessControlRequest) (*workspacepb.RepoAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := RepoPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RepoAccessControlRequestFromPb(pb *workspacepb.RepoAccessControlRequestPb) (*RepoAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := RepoPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RepoAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []RepoPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RepoAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RepoAccessControlResponseToPb(st *RepoAccessControlResponse) (*workspacepb.RepoAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoAccessControlResponsePb{}

	var allPermissionsPb []workspacepb.RepoPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := RepoPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RepoAccessControlResponseFromPb(pb *workspacepb.RepoAccessControlResponsePb) (*RepoAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoAccessControlResponse{}

	var allPermissionsField []RepoPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := RepoPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Git folder (repo) information.
type RepoInfo struct {
	// Name of the current git branch of the git folder (repo).
	// Wire name: 'branch'
	Branch string ``
	// Current git commit id of the git folder (repo).
	// Wire name: 'head_commit_id'
	HeadCommitId string ``
	// Id of the git folder (repo) in the Workspace.
	// Wire name: 'id'
	Id int64 ``
	// Root path of the git folder (repo) in the Workspace.
	// Wire name: 'path'
	Path string ``
	// Git provider of the remote git repository, e.g. `gitHub`.
	// Wire name: 'provider'
	Provider string ``
	// Sparse checkout config for the git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout ``
	// URL of the remote git repository.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RepoInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RepoInfoToPb(st *RepoInfo) (*workspacepb.RepoInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoInfoPb{}
	pb.Branch = st.Branch
	pb.HeadCommitId = st.HeadCommitId
	pb.Id = st.Id
	pb.Path = st.Path
	pb.Provider = st.Provider
	sparseCheckoutPb, err := SparseCheckoutToPb(st.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutPb != nil {
		pb.SparseCheckout = sparseCheckoutPb
	}
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RepoInfoFromPb(pb *workspacepb.RepoInfoPb) (*RepoInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoInfo{}
	st.Branch = pb.Branch
	st.HeadCommitId = pb.HeadCommitId
	st.Id = pb.Id
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := SparseCheckoutFromPb(pb.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutField != nil {
		st.SparseCheckout = sparseCheckoutField
	}
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RepoPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel ``
	ForceSendFields []string            `tf:"-"`
}

func (s *RepoPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RepoPermissionToPb(st *RepoPermission) (*workspacepb.RepoPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := RepoPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RepoPermissionFromPb(pb *workspacepb.RepoPermissionPb) (*RepoPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := RepoPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func RepoPermissionLevelToPb(st *RepoPermissionLevel) (*workspacepb.RepoPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.RepoPermissionLevelPb(*st)
	return &pb, nil
}

func RepoPermissionLevelFromPb(pb *workspacepb.RepoPermissionLevelPb) (*RepoPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := RepoPermissionLevel(*pb)
	return &st, nil
}

type RepoPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RepoPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RepoPermissionsToPb(st *RepoPermissions) (*workspacepb.RepoPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoPermissionsPb{}

	var accessControlListPb []workspacepb.RepoAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := RepoAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RepoPermissionsFromPb(pb *workspacepb.RepoPermissionsPb) (*RepoPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissions{}

	var accessControlListField []RepoAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := RepoAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RepoPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel ``
	ForceSendFields []string            `tf:"-"`
}

func (s *RepoPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepoPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RepoPermissionsDescriptionToPb(st *RepoPermissionsDescription) (*workspacepb.RepoPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := RepoPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RepoPermissionsDescriptionFromPb(pb *workspacepb.RepoPermissionsDescriptionPb) (*RepoPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := RepoPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RepoPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlRequest ``
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func RepoPermissionsRequestToPb(st *RepoPermissionsRequest) (*workspacepb.RepoPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.RepoPermissionsRequestPb{}

	var accessControlListPb []workspacepb.RepoAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := RepoAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.RepoId = st.RepoId

	return pb, nil
}

func RepoPermissionsRequestFromPb(pb *workspacepb.RepoPermissionsRequestPb) (*RepoPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissionsRequest{}

	var accessControlListField []RepoAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := RepoAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.RepoId = pb.RepoId

	return st, nil
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

func ScopeBackendTypeToPb(st *ScopeBackendType) (*workspacepb.ScopeBackendTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.ScopeBackendTypePb(*st)
	return &pb, nil
}

func ScopeBackendTypeFromPb(pb *workspacepb.ScopeBackendTypePb) (*ScopeBackendType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ScopeBackendType(*pb)
	return &st, nil
}

// The metadata about a secret. Returned when listing secrets. Does not contain
// the actual secret value.
type SecretMetadata struct {
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string ``
	// The last updated timestamp (in milliseconds) for the secret.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64    ``
	ForceSendFields      []string `tf:"-"`
}

func (s *SecretMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretMetadata) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SecretMetadataToPb(st *SecretMetadata) (*workspacepb.SecretMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.SecretMetadataPb{}
	pb.Key = st.Key
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SecretMetadataFromPb(pb *workspacepb.SecretMetadataPb) (*SecretMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretMetadata{}
	st.Key = pb.Key
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// An organizational resource for storing secrets. Secret scopes can be
// different types (Databricks-managed, Azure KeyVault backed, etc), and ACLs
// can be applied to control permissions for all secrets within a scope.
type SecretScope struct {
	// The type of secret scope backend.
	// Wire name: 'backend_type'
	BackendType ScopeBackendType ``
	// The metadata for the secret scope if the type is ``AZURE_KEYVAULT``
	// Wire name: 'keyvault_metadata'
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata ``
	// A unique name to identify the secret scope.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *SecretScope) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretScope) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SecretScopeToPb(st *SecretScope) (*workspacepb.SecretScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.SecretScopePb{}
	backendTypePb, err := ScopeBackendTypeToPb(&st.BackendType)
	if err != nil {
		return nil, err
	}
	if backendTypePb != nil {
		pb.BackendType = *backendTypePb
	}
	keyvaultMetadataPb, err := AzureKeyVaultSecretScopeMetadataToPb(st.KeyvaultMetadata)
	if err != nil {
		return nil, err
	}
	if keyvaultMetadataPb != nil {
		pb.KeyvaultMetadata = keyvaultMetadataPb
	}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SecretScopeFromPb(pb *workspacepb.SecretScopePb) (*SecretScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretScope{}
	backendTypeField, err := ScopeBackendTypeFromPb(&pb.BackendType)
	if err != nil {
		return nil, err
	}
	if backendTypeField != nil {
		st.BackendType = *backendTypeField
	}
	keyvaultMetadataField, err := AzureKeyVaultSecretScopeMetadataFromPb(pb.KeyvaultMetadata)
	if err != nil {
		return nil, err
	}
	if keyvaultMetadataField != nil {
		st.KeyvaultMetadata = keyvaultMetadataField
	}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	// Wire name: 'patterns'
	Patterns []string ``
}

func SparseCheckoutToPb(st *SparseCheckout) (*workspacepb.SparseCheckoutPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.SparseCheckoutPb{}
	pb.Patterns = st.Patterns

	return pb, nil
}

func SparseCheckoutFromPb(pb *workspacepb.SparseCheckoutPb) (*SparseCheckout, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparseCheckout{}
	st.Patterns = pb.Patterns

	return st, nil
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckoutUpdate struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	// Wire name: 'patterns'
	Patterns []string ``
}

func SparseCheckoutUpdateToPb(st *SparseCheckoutUpdate) (*workspacepb.SparseCheckoutUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.SparseCheckoutUpdatePb{}
	pb.Patterns = st.Patterns

	return pb, nil
}

func SparseCheckoutUpdateFromPb(pb *workspacepb.SparseCheckoutUpdatePb) (*SparseCheckoutUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparseCheckoutUpdate{}
	st.Patterns = pb.Patterns

	return st, nil
}

type UpdateCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	// Wire name: 'credential_id'
	CredentialId int64 `tf:"-"`
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'git_provider'
	GitProvider string ``
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	// Wire name: 'git_username'
	GitUsername string ``
	// if the credential is the default for the given provider
	// Wire name: 'is_default_for_provider'
	IsDefaultForProvider bool ``
	// the name of the git credential, used for identification and ease of
	// lookup
	// Wire name: 'name'
	Name string ``
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	// Wire name: 'personal_access_token'
	PersonalAccessToken string   ``
	ForceSendFields     []string `tf:"-"`
}

func (s *UpdateCredentialsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateCredentialsRequestToPb(st *UpdateCredentialsRequest) (*workspacepb.UpdateCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.UpdateCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId
	pb.GitProvider = st.GitProvider
	pb.GitUsername = st.GitUsername
	pb.IsDefaultForProvider = st.IsDefaultForProvider
	pb.Name = st.Name
	pb.PersonalAccessToken = st.PersonalAccessToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateCredentialsRequestFromPb(pb *workspacepb.UpdateCredentialsRequestPb) (*UpdateCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCredentialsRequest{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.IsDefaultForProvider = pb.IsDefaultForProvider
	st.Name = pb.Name
	st.PersonalAccessToken = pb.PersonalAccessToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateRepoRequest struct {
	// Branch that the local version of the repo is checked out to.
	// Wire name: 'branch'
	Branch string ``
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'repo_id'
	RepoId int64 `tf:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckoutUpdate ``
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	// Wire name: 'tag'
	Tag             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateRepoRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateRepoRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateRepoRequestToPb(st *UpdateRepoRequest) (*workspacepb.UpdateRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.UpdateRepoRequestPb{}
	pb.Branch = st.Branch
	pb.RepoId = st.RepoId
	sparseCheckoutPb, err := SparseCheckoutUpdateToPb(st.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutPb != nil {
		pb.SparseCheckout = sparseCheckoutPb
	}
	pb.Tag = st.Tag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateRepoRequestFromPb(pb *workspacepb.UpdateRepoRequestPb) (*UpdateRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRepoRequest{}
	st.Branch = pb.Branch
	st.RepoId = pb.RepoId
	sparseCheckoutField, err := SparseCheckoutUpdateFromPb(pb.SparseCheckout)
	if err != nil {
		return nil, err
	}
	if sparseCheckoutField != nil {
		st.SparseCheckout = sparseCheckoutField
	}
	st.Tag = pb.Tag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *WorkspaceObjectAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WorkspaceObjectAccessControlRequestToPb(st *WorkspaceObjectAccessControlRequest) (*workspacepb.WorkspaceObjectAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.WorkspaceObjectAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := WorkspaceObjectPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WorkspaceObjectAccessControlRequestFromPb(pb *workspacepb.WorkspaceObjectAccessControlRequestPb) (*WorkspaceObjectAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := WorkspaceObjectPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WorkspaceObjectAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []WorkspaceObjectPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *WorkspaceObjectAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WorkspaceObjectAccessControlResponseToPb(st *WorkspaceObjectAccessControlResponse) (*workspacepb.WorkspaceObjectAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.WorkspaceObjectAccessControlResponsePb{}

	var allPermissionsPb []workspacepb.WorkspaceObjectPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := WorkspaceObjectPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WorkspaceObjectAccessControlResponseFromPb(pb *workspacepb.WorkspaceObjectAccessControlResponsePb) (*WorkspaceObjectAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectAccessControlResponse{}

	var allPermissionsField []WorkspaceObjectPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := WorkspaceObjectPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WorkspaceObjectPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (s *WorkspaceObjectPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WorkspaceObjectPermissionToPb(st *WorkspaceObjectPermission) (*workspacepb.WorkspaceObjectPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.WorkspaceObjectPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := WorkspaceObjectPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WorkspaceObjectPermissionFromPb(pb *workspacepb.WorkspaceObjectPermissionPb) (*WorkspaceObjectPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := WorkspaceObjectPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func WorkspaceObjectPermissionLevelToPb(st *WorkspaceObjectPermissionLevel) (*workspacepb.WorkspaceObjectPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspacepb.WorkspaceObjectPermissionLevelPb(*st)
	return &pb, nil
}

func WorkspaceObjectPermissionLevelFromPb(pb *workspacepb.WorkspaceObjectPermissionLevelPb) (*WorkspaceObjectPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspaceObjectPermissionLevel(*pb)
	return &st, nil
}

type WorkspaceObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []WorkspaceObjectAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *WorkspaceObjectPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WorkspaceObjectPermissionsToPb(st *WorkspaceObjectPermissions) (*workspacepb.WorkspaceObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.WorkspaceObjectPermissionsPb{}

	var accessControlListPb []workspacepb.WorkspaceObjectAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := WorkspaceObjectAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WorkspaceObjectPermissionsFromPb(pb *workspacepb.WorkspaceObjectPermissionsPb) (*WorkspaceObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissions{}

	var accessControlListField []WorkspaceObjectAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := WorkspaceObjectAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WorkspaceObjectPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (s *WorkspaceObjectPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceObjectPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WorkspaceObjectPermissionsDescriptionToPb(st *WorkspaceObjectPermissionsDescription) (*workspacepb.WorkspaceObjectPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.WorkspaceObjectPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := WorkspaceObjectPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WorkspaceObjectPermissionsDescriptionFromPb(pb *workspacepb.WorkspaceObjectPermissionsDescriptionPb) (*WorkspaceObjectPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := WorkspaceObjectPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WorkspaceObjectPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []WorkspaceObjectAccessControlRequest ``
	// The workspace object for which to get or manage permissions.
	// Wire name: 'workspace_object_id'
	WorkspaceObjectId string `tf:"-"`
	// The workspace object type for which to get or manage permissions.
	// Wire name: 'workspace_object_type'
	WorkspaceObjectType string `tf:"-"`
}

func WorkspaceObjectPermissionsRequestToPb(st *WorkspaceObjectPermissionsRequest) (*workspacepb.WorkspaceObjectPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.WorkspaceObjectPermissionsRequestPb{}

	var accessControlListPb []workspacepb.WorkspaceObjectAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := WorkspaceObjectAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.WorkspaceObjectId = st.WorkspaceObjectId
	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
}

func WorkspaceObjectPermissionsRequestFromPb(pb *workspacepb.WorkspaceObjectPermissionsRequestPb) (*WorkspaceObjectPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissionsRequest{}

	var accessControlListField []WorkspaceObjectAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := WorkspaceObjectAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.WorkspaceObjectId = pb.WorkspaceObjectId
	st.WorkspaceObjectType = pb.WorkspaceObjectType

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
