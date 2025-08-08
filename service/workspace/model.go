// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

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

func (st AclItem) MarshalJSON() ([]byte, error) {
	pb, err := AclItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AclItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.AclItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AclItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st AzureKeyVaultSecretScopeMetadata) MarshalJSON() ([]byte, error) {
	pb, err := AzureKeyVaultSecretScopeMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AzureKeyVaultSecretScopeMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.AzureKeyVaultSecretScopeMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AzureKeyVaultSecretScopeMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st CreateCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.CreateCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st CreateCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.CreateCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st CreateRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.CreateRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st CreateRepoResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateRepoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateRepoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.CreateRepoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateRepoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st CreateScope) MarshalJSON() ([]byte, error) {
	pb, err := CreateScopeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateScope) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.CreateScopePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateScopeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st CredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := CredentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.CredentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CredentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st Delete) MarshalJSON() ([]byte, error) {
	pb, err := DeleteToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Delete) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.DeletePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteToPb(st *Delete) (*workspacepb.DeletePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.DeletePb{}
	pb.Path = st.Path
	pb.Recursive = st.Recursive

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteFromPb(pb *workspacepb.DeletePb) (*Delete, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Delete{}
	st.Path = pb.Path
	st.Recursive = pb.Recursive

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st DeleteAcl) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAclToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAcl) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.DeleteAclPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAclFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st DeleteCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.DeleteCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st DeleteRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.DeleteRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st DeleteScope) MarshalJSON() ([]byte, error) {
	pb, err := DeleteScopeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteScope) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.DeleteScopePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteScopeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st DeleteSecret) MarshalJSON() ([]byte, error) {
	pb, err := DeleteSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.DeleteSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ExportRequest) MarshalJSON() ([]byte, error) {
	pb, err := ExportRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExportRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ExportRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExportRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ExportResponse) MarshalJSON() ([]byte, error) {
	pb, err := ExportResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExportResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ExportResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExportResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExportResponseToPb(st *ExportResponse) (*workspacepb.ExportResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ExportResponsePb{}
	pb.Content = st.Content
	pb.FileType = st.FileType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExportResponseFromPb(pb *workspacepb.ExportResponsePb) (*ExportResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportResponse{}
	st.Content = pb.Content
	st.FileType = pb.FileType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st GetAclRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAclRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAclRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetAclRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAclRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func (st GetRepoPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRepoPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRepoPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetRepoPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRepoPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetRepoPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetRepoPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRepoPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetRepoPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRepoPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetRepoPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRepoPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRepoPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetRepoPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRepoPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetRepoResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetRepoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRepoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetRepoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRepoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st GetSecretRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetSecretRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetSecretRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetSecretRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetSecretRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetSecretResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetSecretResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetSecretResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetSecretResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetSecretResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetSecretResponseToPb(st *GetSecretResponse) (*workspacepb.GetSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.GetSecretResponsePb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetSecretResponseFromPb(pb *workspacepb.GetSecretResponsePb) (*GetSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSecretResponse{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string `tf:"-"`
}

func (st GetStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetWorkspaceObjectPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetWorkspaceObjectPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetWorkspaceObjectPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetWorkspaceObjectPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetWorkspaceObjectPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetWorkspaceObjectPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetWorkspaceObjectPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetWorkspaceObjectPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetWorkspaceObjectPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetWorkspaceObjectPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st GetWorkspaceObjectPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetWorkspaceObjectPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetWorkspaceObjectPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.GetWorkspaceObjectPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetWorkspaceObjectPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st Import) MarshalJSON() ([]byte, error) {
	pb, err := ImportToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Import) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ImportPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ImportFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st ListAclsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAclsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAclsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListAclsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAclsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListAclsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListAclsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAclsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListAclsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAclsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListReposRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListReposRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListReposRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListReposRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListReposRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListReposRequestToPb(st *ListReposRequest) (*workspacepb.ListReposRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListReposRequestPb{}
	pb.NextPageToken = st.NextPageToken
	pb.PathPrefix = st.PathPrefix

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListReposRequestFromPb(pb *workspacepb.ListReposRequestPb) (*ListReposRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListReposRequest{}
	st.NextPageToken = pb.NextPageToken
	st.PathPrefix = pb.PathPrefix

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st ListReposResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListReposResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListReposResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListReposResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListReposResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListResponse struct {
	// List of objects.
	// Wire name: 'objects'
	Objects []ObjectInfo ``
}

func (st ListResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListScopesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListScopesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListScopesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListScopesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListScopesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListSecretsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListSecretsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListSecretsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListSecretsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListSecretsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListSecretsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListSecretsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListSecretsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListSecretsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListSecretsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ListWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ListWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListWorkspaceRequestToPb(st *ListWorkspaceRequest) (*workspacepb.ListWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.ListWorkspaceRequestPb{}
	pb.NotebooksModifiedAfter = st.NotebooksModifiedAfter
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListWorkspaceRequestFromPb(pb *workspacepb.ListWorkspaceRequestPb) (*ListWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWorkspaceRequest{}
	st.NotebooksModifiedAfter = pb.NotebooksModifiedAfter
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	// Wire name: 'path'
	Path string ``
}

func (st Mkdirs) MarshalJSON() ([]byte, error) {
	pb, err := MkdirsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Mkdirs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.MkdirsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MkdirsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st ObjectInfo) MarshalJSON() ([]byte, error) {
	pb, err := ObjectInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ObjectInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.ObjectInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ObjectInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st PutAcl) MarshalJSON() ([]byte, error) {
	pb, err := PutAclToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PutAcl) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.PutAclPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PutAclFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st PutSecret) MarshalJSON() ([]byte, error) {
	pb, err := PutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.PutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st RepoAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := RepoAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st RepoAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := RepoAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st RepoInfo) MarshalJSON() ([]byte, error) {
	pb, err := RepoInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st RepoPermission) MarshalJSON() ([]byte, error) {
	pb, err := RepoPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st RepoPermissions) MarshalJSON() ([]byte, error) {
	pb, err := RepoPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RepoPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel ``
	ForceSendFields []string            `tf:"-"`
}

func (st RepoPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := RepoPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RepoPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlRequest ``
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func (st RepoPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := RepoPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepoPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.RepoPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepoPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st SecretMetadata) MarshalJSON() ([]byte, error) {
	pb, err := SecretMetadataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SecretMetadata) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.SecretMetadataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SecretMetadataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SecretMetadataToPb(st *SecretMetadata) (*workspacepb.SecretMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacepb.SecretMetadataPb{}
	pb.Key = st.Key
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SecretMetadataFromPb(pb *workspacepb.SecretMetadataPb) (*SecretMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretMetadata{}
	st.Key = pb.Key
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st SecretScope) MarshalJSON() ([]byte, error) {
	pb, err := SecretScopeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SecretScope) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.SecretScopePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SecretScopeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st SparseCheckout) MarshalJSON() ([]byte, error) {
	pb, err := SparseCheckoutToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparseCheckout) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.SparseCheckoutPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparseCheckoutFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st SparseCheckoutUpdate) MarshalJSON() ([]byte, error) {
	pb, err := SparseCheckoutUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparseCheckoutUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.SparseCheckoutUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparseCheckoutUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

func (st UpdateCredentialsRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCredentialsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCredentialsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.UpdateCredentialsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCredentialsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st UpdateRepoRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRepoRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRepoRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.UpdateRepoRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRepoRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st WorkspaceObjectAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceObjectAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceObjectAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.WorkspaceObjectAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceObjectAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st WorkspaceObjectAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceObjectAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceObjectAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.WorkspaceObjectAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceObjectAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st WorkspaceObjectPermission) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceObjectPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceObjectPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.WorkspaceObjectPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceObjectPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st WorkspaceObjectPermissions) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceObjectPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceObjectPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.WorkspaceObjectPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceObjectPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type WorkspaceObjectPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (st WorkspaceObjectPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceObjectPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceObjectPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.WorkspaceObjectPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceObjectPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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

func (st WorkspaceObjectPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceObjectPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceObjectPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacepb.WorkspaceObjectPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceObjectPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
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
