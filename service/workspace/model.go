// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AclItem struct {
	// The permission level applied to the principal.
	// Wire name: 'permission'
	Permission AclPermission
	// The principal in which the permission is applied.
	// Wire name: 'principal'
	Principal string
}

func aclItemToPb(st *AclItem) (*aclItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aclItemPb{}
	pb.Permission = st.Permission

	pb.Principal = st.Principal

	return pb, nil
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

type aclItemPb struct {
	// The permission level applied to the principal.
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	Principal string `json:"principal"`
}

func aclItemFromPb(pb *aclItemPb) (*AclItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AclItem{}
	st.Permission = pb.Permission
	st.Principal = pb.Principal

	return st, nil
}

type AclPermission string
type aclPermissionPb string

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

func aclPermissionToPb(st *AclPermission) (*aclPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := aclPermissionPb(*st)
	return &pb, nil
}

func aclPermissionFromPb(pb *aclPermissionPb) (*AclPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := AclPermission(*pb)
	return &st, nil
}

type AzureKeyVaultSecretScopeMetadata struct {
	// The DNS of the KeyVault
	// Wire name: 'dns_name'
	DnsName string
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	// Wire name: 'resource_id'
	ResourceId string
}

func azureKeyVaultSecretScopeMetadataToPb(st *AzureKeyVaultSecretScopeMetadata) (*azureKeyVaultSecretScopeMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureKeyVaultSecretScopeMetadataPb{}
	pb.DnsName = st.DnsName

	pb.ResourceId = st.ResourceId

	return pb, nil
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

type azureKeyVaultSecretScopeMetadataPb struct {
	// The DNS of the KeyVault
	DnsName string `json:"dns_name"`
	// The resource id of the azure KeyVault that user wants to associate the
	// scope with.
	ResourceId string `json:"resource_id"`
}

func azureKeyVaultSecretScopeMetadataFromPb(pb *azureKeyVaultSecretScopeMetadataPb) (*AzureKeyVaultSecretScopeMetadata, error) {
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
	GitProvider string
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	// Wire name: 'git_username'
	GitUsername string
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	// Wire name: 'personal_access_token'
	PersonalAccessToken string

	ForceSendFields []string `tf:"-"`
}

func createCredentialsRequestToPb(st *CreateCredentialsRequest) (*createCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialsRequestPb{}
	pb.GitProvider = st.GitProvider

	pb.GitUsername = st.GitUsername

	pb.PersonalAccessToken = st.PersonalAccessToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createCredentialsRequestPb struct {
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
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCredentialsRequestFromPb(pb *createCredentialsRequestPb) (*CreateCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialsRequest{}
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.PersonalAccessToken = pb.PersonalAccessToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCredentialsResponse struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string

	ForceSendFields []string `tf:"-"`
}

func createCredentialsResponseToPb(st *CreateCredentialsResponse) (*createCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialsResponsePb{}
	pb.CredentialId = st.CredentialId

	pb.GitProvider = st.GitProvider

	pb.GitUsername = st.GitUsername

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createCredentialsResponsePb struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider string `json:"git_provider"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCredentialsResponseFromPb(pb *createCredentialsResponsePb) (*CreateCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialsResponse{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRepoRequest struct {
	// Desired path for the repo in the workspace. Almost any path in the
	// workspace can be chosen. If repo is created in `/Repos`, path must be in
	// the format `/Repos/{folder}/{repo-name}`.
	// Wire name: 'path'
	Path string
	// Git provider. This field is case-insensitive. The available Git providers
	// are `gitHub`, `bitbucketCloud`, `gitLab`, `azureDevOpsServices`,
	// `gitHubEnterprise`, `bitbucketServer`, `gitLabEnterpriseEdition` and
	// `awsCodeCommit`.
	// Wire name: 'provider'
	Provider string
	// If specified, the repo will be created with sparse checkout enabled. You
	// cannot enable/disable sparse checkout after the repo is created.
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout
	// URL of the Git repository to be linked.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func createRepoRequestToPb(st *CreateRepoRequest) (*createRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRepoRequestPb{}
	pb.Path = st.Path

	pb.Provider = st.Provider

	sparseCheckoutPb, err := sparseCheckoutToPb(st.SparseCheckout)
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

type createRepoRequestPb struct {
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
	SparseCheckout *sparseCheckoutPb `json:"sparse_checkout,omitempty"`
	// URL of the Git repository to be linked.
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRepoRequestFromPb(pb *createRepoRequestPb) (*CreateRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRepoRequest{}
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := sparseCheckoutFromPb(pb.SparseCheckout)
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

func (st *createRepoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRepoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRepoResponse struct {
	// Branch that the Git folder (repo) is checked out to.
	// Wire name: 'branch'
	Branch string
	// SHA-1 hash representing the commit ID of the current HEAD of the Git
	// folder (repo).
	// Wire name: 'head_commit_id'
	HeadCommitId string
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'id'
	Id int64
	// Path of the Git folder (repo) in the workspace.
	// Wire name: 'path'
	Path string
	// Git provider of the linked Git repository.
	// Wire name: 'provider'
	Provider string
	// Sparse checkout settings for the Git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout
	// URL of the linked Git repository.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func createRepoResponseToPb(st *CreateRepoResponse) (*createRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRepoResponsePb{}
	pb.Branch = st.Branch

	pb.HeadCommitId = st.HeadCommitId

	pb.Id = st.Id

	pb.Path = st.Path

	pb.Provider = st.Provider

	sparseCheckoutPb, err := sparseCheckoutToPb(st.SparseCheckout)
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

type createRepoResponsePb struct {
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
	SparseCheckout *sparseCheckoutPb `json:"sparse_checkout,omitempty"`
	// URL of the linked Git repository.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRepoResponseFromPb(pb *createRepoResponsePb) (*CreateRepoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRepoResponse{}
	st.Branch = pb.Branch
	st.HeadCommitId = pb.HeadCommitId
	st.Id = pb.Id
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := sparseCheckoutFromPb(pb.SparseCheckout)
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

func (st *createRepoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRepoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateScope struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	// Wire name: 'backend_azure_keyvault'
	BackendAzureKeyvault *AzureKeyVaultSecretScopeMetadata
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	// Wire name: 'initial_manage_principal'
	InitialManagePrincipal string
	// Scope name requested by the user. Scope names are unique.
	// Wire name: 'scope'
	Scope string
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	// Wire name: 'scope_backend_type'
	ScopeBackendType ScopeBackendType

	ForceSendFields []string `tf:"-"`
}

func createScopeToPb(st *CreateScope) (*createScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createScopePb{}
	backendAzureKeyvaultPb, err := azureKeyVaultSecretScopeMetadataToPb(st.BackendAzureKeyvault)
	if err != nil {
		return nil, err
	}
	if backendAzureKeyvaultPb != nil {
		pb.BackendAzureKeyvault = backendAzureKeyvaultPb
	}

	pb.InitialManagePrincipal = st.InitialManagePrincipal

	pb.Scope = st.Scope

	pb.ScopeBackendType = st.ScopeBackendType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createScopePb struct {
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	BackendAzureKeyvault *azureKeyVaultSecretScopeMetadataPb `json:"backend_azure_keyvault,omitempty"`
	// The principal that is initially granted `MANAGE` permission to the
	// created scope.
	InitialManagePrincipal string `json:"initial_manage_principal,omitempty"`
	// Scope name requested by the user. Scope names are unique.
	Scope string `json:"scope"`
	// The backend type the scope will be created with. If not specified, will
	// default to `DATABRICKS`
	ScopeBackendType ScopeBackendType `json:"scope_backend_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createScopeFromPb(pb *createScopePb) (*CreateScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScope{}
	backendAzureKeyvaultField, err := azureKeyVaultSecretScopeMetadataFromPb(pb.BackendAzureKeyvault)
	if err != nil {
		return nil, err
	}
	if backendAzureKeyvaultField != nil {
		st.BackendAzureKeyvault = backendAzureKeyvaultField
	}
	st.InitialManagePrincipal = pb.InitialManagePrincipal
	st.Scope = pb.Scope
	st.ScopeBackendType = pb.ScopeBackendType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createScopePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createScopePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateScopeResponse struct {
}

func createScopeResponseToPb(st *CreateScopeResponse) (*createScopeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createScopeResponsePb{}

	return pb, nil
}

func (st *CreateScopeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createScopeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createScopeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateScopeResponse) MarshalJSON() ([]byte, error) {
	pb, err := createScopeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createScopeResponsePb struct {
}

func createScopeResponseFromPb(pb *createScopeResponsePb) (*CreateScopeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScopeResponse{}

	return st, nil
}

type CredentialInfo struct {
	// ID of the credential object in the workspace.
	// Wire name: 'credential_id'
	CredentialId int64
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string

	ForceSendFields []string `tf:"-"`
}

func credentialInfoToPb(st *CredentialInfo) (*credentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialInfoPb{}
	pb.CredentialId = st.CredentialId

	pb.GitProvider = st.GitProvider

	pb.GitUsername = st.GitUsername

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type credentialInfoPb struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider string `json:"git_provider,omitempty"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialInfoFromPb(pb *credentialInfoPb) (*CredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialInfo{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Delete struct {
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	// Wire name: 'recursive'
	Recursive bool

	ForceSendFields []string `tf:"-"`
}

func deleteToPb(st *Delete) (*deletePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePb{}
	pb.Path = st.Path

	pb.Recursive = st.Recursive

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deletePb struct {
	// The absolute path of the notebook or directory.
	Path string `json:"path"`
	// The flag that specifies whether to delete the object recursively. It is
	// `false` by default. Please note this deleting directory is not atomic. If
	// it fails in the middle, some of objects under this directory may be
	// deleted and cannot be undone.
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteFromPb(pb *deletePb) (*Delete, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Delete{}
	st.Path = pb.Path
	st.Recursive = pb.Recursive

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deletePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deletePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAcl struct {
	// The principal to remove an existing ACL from.
	// Wire name: 'principal'
	Principal string
	// The name of the scope to remove permissions from.
	// Wire name: 'scope'
	Scope string
}

func deleteAclToPb(st *DeleteAcl) (*deleteAclPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAclPb{}
	pb.Principal = st.Principal

	pb.Scope = st.Scope

	return pb, nil
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

type deleteAclPb struct {
	// The principal to remove an existing ACL from.
	Principal string `json:"principal"`
	// The name of the scope to remove permissions from.
	Scope string `json:"scope"`
}

func deleteAclFromPb(pb *deleteAclPb) (*DeleteAcl, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAcl{}
	st.Principal = pb.Principal
	st.Scope = pb.Scope

	return st, nil
}

type DeleteAclResponse struct {
}

func deleteAclResponseToPb(st *DeleteAclResponse) (*deleteAclResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAclResponsePb{}

	return pb, nil
}

func (st *DeleteAclResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAclResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAclResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAclResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteAclResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAclResponsePb struct {
}

func deleteAclResponseFromPb(pb *deleteAclResponsePb) (*DeleteAclResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAclResponse{}

	return st, nil
}

// Delete a credential
type DeleteCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	// Wire name: 'credential_id'
	CredentialId int64 `tf:"-"`
}

func deleteCredentialsRequestToPb(st *DeleteCredentialsRequest) (*deleteCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	return pb, nil
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

type deleteCredentialsRequestPb struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
}

func deleteCredentialsRequestFromPb(pb *deleteCredentialsRequestPb) (*DeleteCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialsRequest{}
	st.CredentialId = pb.CredentialId

	return st, nil
}

type DeleteCredentialsResponse struct {
}

func deleteCredentialsResponseToPb(st *DeleteCredentialsResponse) (*deleteCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialsResponsePb{}

	return pb, nil
}

func (st *DeleteCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCredentialsResponsePb struct {
}

func deleteCredentialsResponseFromPb(pb *deleteCredentialsResponsePb) (*DeleteCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialsResponse{}

	return st, nil
}

// Delete a repo
type DeleteRepoRequest struct {
	// The ID for the corresponding repo to delete.
	// Wire name: 'repo_id'
	RepoId int64 `tf:"-"`
}

func deleteRepoRequestToPb(st *DeleteRepoRequest) (*deleteRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRepoRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
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

type deleteRepoRequestPb struct {
	// The ID for the corresponding repo to delete.
	RepoId int64 `json:"-" url:"-"`
}

func deleteRepoRequestFromPb(pb *deleteRepoRequestPb) (*DeleteRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRepoRequest{}
	st.RepoId = pb.RepoId

	return st, nil
}

type DeleteRepoResponse struct {
}

func deleteRepoResponseToPb(st *DeleteRepoResponse) (*deleteRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRepoResponsePb{}

	return pb, nil
}

func (st *DeleteRepoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRepoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRepoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRepoResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRepoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRepoResponsePb struct {
}

func deleteRepoResponseFromPb(pb *deleteRepoResponsePb) (*DeleteRepoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRepoResponse{}

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

type DeleteScope struct {
	// Name of the scope to delete.
	// Wire name: 'scope'
	Scope string
}

func deleteScopeToPb(st *DeleteScope) (*deleteScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScopePb{}
	pb.Scope = st.Scope

	return pb, nil
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

type deleteScopePb struct {
	// Name of the scope to delete.
	Scope string `json:"scope"`
}

func deleteScopeFromPb(pb *deleteScopePb) (*DeleteScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteScope{}
	st.Scope = pb.Scope

	return st, nil
}

type DeleteScopeResponse struct {
}

func deleteScopeResponseToPb(st *DeleteScopeResponse) (*deleteScopeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScopeResponsePb{}

	return pb, nil
}

func (st *DeleteScopeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteScopeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteScopeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteScopeResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteScopeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteScopeResponsePb struct {
}

func deleteScopeResponseFromPb(pb *deleteScopeResponsePb) (*DeleteScopeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteScopeResponse{}

	return st, nil
}

type DeleteSecret struct {
	// Name of the secret to delete.
	// Wire name: 'key'
	Key string
	// The name of the scope that contains the secret to delete.
	// Wire name: 'scope'
	Scope string
}

func deleteSecretToPb(st *DeleteSecret) (*deleteSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSecretPb{}
	pb.Key = st.Key

	pb.Scope = st.Scope

	return pb, nil
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

type deleteSecretPb struct {
	// Name of the secret to delete.
	Key string `json:"key"`
	// The name of the scope that contains the secret to delete.
	Scope string `json:"scope"`
}

func deleteSecretFromPb(pb *deleteSecretPb) (*DeleteSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSecret{}
	st.Key = pb.Key
	st.Scope = pb.Scope

	return st, nil
}

type DeleteSecretResponse struct {
}

func deleteSecretResponseToPb(st *DeleteSecretResponse) (*deleteSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSecretResponsePb{}

	return pb, nil
}

func (st *DeleteSecretResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSecretResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSecretResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSecretResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteSecretResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteSecretResponsePb struct {
}

func deleteSecretResponseFromPb(pb *deleteSecretResponsePb) (*DeleteSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSecretResponse{}

	return st, nil
}

// The format for workspace import and export.
type ExportFormat string
type exportFormatPb string

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

// Type always returns ExportFormat to satisfy [pflag.Value] interface
func (f *ExportFormat) Type() string {
	return "ExportFormat"
}

func exportFormatToPb(st *ExportFormat) (*exportFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := exportFormatPb(*st)
	return &pb, nil
}

func exportFormatFromPb(pb *exportFormatPb) (*ExportFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExportFormat(*pb)
	return &st, nil
}

// Export a workspace object
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

func exportRequestToPb(st *ExportRequest) (*exportRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportRequestPb{}
	pb.Format = st.Format

	pb.Path = st.Path

	return pb, nil
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

type exportRequestPb struct {
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

func exportRequestFromPb(pb *exportRequestPb) (*ExportRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRequest{}
	st.Format = pb.Format
	st.Path = pb.Path

	return st, nil
}

// The request field `direct_download` determines whether a JSON response or
// binary contents are returned by this endpoint.
type ExportResponse struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	// Wire name: 'content'
	Content string
	// The file type of the exported file.
	// Wire name: 'file_type'
	FileType string

	ForceSendFields []string `tf:"-"`
}

func exportResponseToPb(st *ExportResponse) (*exportResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportResponsePb{}
	pb.Content = st.Content

	pb.FileType = st.FileType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type exportResponsePb struct {
	// The base64-encoded content. If the limit (10MB) is exceeded, exception
	// with error code **MAX_NOTEBOOK_SIZE_EXCEEDED** is thrown.
	Content string `json:"content,omitempty"`
	// The file type of the exported file.
	FileType string `json:"file_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func exportResponseFromPb(pb *exportResponsePb) (*ExportResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportResponse{}
	st.Content = pb.Content
	st.FileType = pb.FileType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *exportResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st exportResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get secret ACL details
type GetAclRequest struct {
	// The principal to fetch ACL information for.
	// Wire name: 'principal'
	Principal string `tf:"-"`
	// The name of the scope to fetch ACL information from.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func getAclRequestToPb(st *GetAclRequest) (*getAclRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAclRequestPb{}
	pb.Principal = st.Principal

	pb.Scope = st.Scope

	return pb, nil
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

type getAclRequestPb struct {
	// The principal to fetch ACL information for.
	Principal string `json:"-" url:"principal"`
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`
}

func getAclRequestFromPb(pb *getAclRequestPb) (*GetAclRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAclRequest{}
	st.Principal = pb.Principal
	st.Scope = pb.Scope

	return st, nil
}

// Get a credential entry
type GetCredentialsRequest struct {
	// The ID for the corresponding credential to access.
	// Wire name: 'credential_id'
	CredentialId int64 `tf:"-"`
}

func getCredentialsRequestToPb(st *GetCredentialsRequest) (*getCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	return pb, nil
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

type getCredentialsRequestPb struct {
	// The ID for the corresponding credential to access.
	CredentialId int64 `json:"-" url:"-"`
}

func getCredentialsRequestFromPb(pb *getCredentialsRequestPb) (*GetCredentialsRequest, error) {
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
	CredentialId int64
	// The Git provider associated with the credential.
	// Wire name: 'git_provider'
	GitProvider string
	// The username or email provided with your Git provider account and
	// associated with the credential.
	// Wire name: 'git_username'
	GitUsername string

	ForceSendFields []string `tf:"-"`
}

func getCredentialsResponseToPb(st *GetCredentialsResponse) (*getCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsResponsePb{}
	pb.CredentialId = st.CredentialId

	pb.GitProvider = st.GitProvider

	pb.GitUsername = st.GitUsername

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getCredentialsResponsePb struct {
	// ID of the credential object in the workspace.
	CredentialId int64 `json:"credential_id"`
	// The Git provider associated with the credential.
	GitProvider string `json:"git_provider,omitempty"`
	// The username or email provided with your Git provider account and
	// associated with the credential.
	GitUsername string `json:"git_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCredentialsResponseFromPb(pb *getCredentialsResponsePb) (*GetCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsResponse{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get repo permission levels
type GetRepoPermissionLevelsRequest struct {
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func getRepoPermissionLevelsRequestToPb(st *GetRepoPermissionLevelsRequest) (*getRepoPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoPermissionLevelsRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
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

type getRepoPermissionLevelsRequestPb struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`
}

func getRepoPermissionLevelsRequestFromPb(pb *getRepoPermissionLevelsRequestPb) (*GetRepoPermissionLevelsRequest, error) {
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
	PermissionLevels []RepoPermissionsDescription
}

func getRepoPermissionLevelsResponseToPb(st *GetRepoPermissionLevelsResponse) (*getRepoPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoPermissionLevelsResponsePb{}

	var permissionLevelsPb []repoPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := repoPermissionsDescriptionToPb(&item)
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

type getRepoPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []repoPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getRepoPermissionLevelsResponseFromPb(pb *getRepoPermissionLevelsResponsePb) (*GetRepoPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoPermissionLevelsResponse{}

	var permissionLevelsField []RepoPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := repoPermissionsDescriptionFromPb(&itemPb)
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

// Get repo permissions
type GetRepoPermissionsRequest struct {
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func getRepoPermissionsRequestToPb(st *GetRepoPermissionsRequest) (*getRepoPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoPermissionsRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
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

type getRepoPermissionsRequestPb struct {
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`
}

func getRepoPermissionsRequestFromPb(pb *getRepoPermissionsRequestPb) (*GetRepoPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoPermissionsRequest{}
	st.RepoId = pb.RepoId

	return st, nil
}

// Get a repo
type GetRepoRequest struct {
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'repo_id'
	RepoId int64 `tf:"-"`
}

func getRepoRequestToPb(st *GetRepoRequest) (*getRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
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

type getRepoRequestPb struct {
	// ID of the Git folder (repo) object in the workspace.
	RepoId int64 `json:"-" url:"-"`
}

func getRepoRequestFromPb(pb *getRepoRequestPb) (*GetRepoRequest, error) {
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
	Branch string
	// SHA-1 hash representing the commit ID of the current HEAD of the repo.
	// Wire name: 'head_commit_id'
	HeadCommitId string
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'id'
	Id int64
	// Path of the Git folder (repo) in the workspace.
	// Wire name: 'path'
	Path string
	// Git provider of the linked Git repository.
	// Wire name: 'provider'
	Provider string
	// Sparse checkout settings for the Git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout
	// URL of the linked Git repository.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func getRepoResponseToPb(st *GetRepoResponse) (*getRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoResponsePb{}
	pb.Branch = st.Branch

	pb.HeadCommitId = st.HeadCommitId

	pb.Id = st.Id

	pb.Path = st.Path

	pb.Provider = st.Provider

	sparseCheckoutPb, err := sparseCheckoutToPb(st.SparseCheckout)
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

type getRepoResponsePb struct {
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
	SparseCheckout *sparseCheckoutPb `json:"sparse_checkout,omitempty"`
	// URL of the linked Git repository.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRepoResponseFromPb(pb *getRepoResponsePb) (*GetRepoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoResponse{}
	st.Branch = pb.Branch
	st.HeadCommitId = pb.HeadCommitId
	st.Id = pb.Id
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := sparseCheckoutFromPb(pb.SparseCheckout)
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

func (st *getRepoResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRepoResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a secret
type GetSecretRequest struct {
	// The key to fetch secret for.
	// Wire name: 'key'
	Key string `tf:"-"`
	// The name of the scope to fetch secret information from.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func getSecretRequestToPb(st *GetSecretRequest) (*getSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSecretRequestPb{}
	pb.Key = st.Key

	pb.Scope = st.Scope

	return pb, nil
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

type getSecretRequestPb struct {
	// The key to fetch secret for.
	Key string `json:"-" url:"key"`
	// The name of the scope to fetch secret information from.
	Scope string `json:"-" url:"scope"`
}

func getSecretRequestFromPb(pb *getSecretRequestPb) (*GetSecretRequest, error) {
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
	Key string
	// The value of the secret in its byte representation.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func getSecretResponseToPb(st *GetSecretResponse) (*getSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSecretResponsePb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getSecretResponsePb struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The value of the secret in its byte representation.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getSecretResponseFromPb(pb *getSecretResponsePb) (*GetSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSecretResponse{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getSecretResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getSecretResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get status
type GetStatusRequest struct {
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string `tf:"-"`
}

func getStatusRequestToPb(st *GetStatusRequest) (*getStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatusRequestPb{}
	pb.Path = st.Path

	return pb, nil
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

type getStatusRequestPb struct {
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`
}

func getStatusRequestFromPb(pb *getStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	st.Path = pb.Path

	return st, nil
}

// Get workspace object permission levels
type GetWorkspaceObjectPermissionLevelsRequest struct {
	// The workspace object for which to get or manage permissions.
	// Wire name: 'workspace_object_id'
	WorkspaceObjectId string `tf:"-"`
	// The workspace object type for which to get or manage permissions.
	// Wire name: 'workspace_object_type'
	WorkspaceObjectType string `tf:"-"`
}

func getWorkspaceObjectPermissionLevelsRequestToPb(st *GetWorkspaceObjectPermissionLevelsRequest) (*getWorkspaceObjectPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceObjectPermissionLevelsRequestPb{}
	pb.WorkspaceObjectId = st.WorkspaceObjectId

	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
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

type getWorkspaceObjectPermissionLevelsRequestPb struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`
}

func getWorkspaceObjectPermissionLevelsRequestFromPb(pb *getWorkspaceObjectPermissionLevelsRequestPb) (*GetWorkspaceObjectPermissionLevelsRequest, error) {
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
	PermissionLevels []WorkspaceObjectPermissionsDescription
}

func getWorkspaceObjectPermissionLevelsResponseToPb(st *GetWorkspaceObjectPermissionLevelsResponse) (*getWorkspaceObjectPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceObjectPermissionLevelsResponsePb{}

	var permissionLevelsPb []workspaceObjectPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := workspaceObjectPermissionsDescriptionToPb(&item)
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

type getWorkspaceObjectPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []workspaceObjectPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getWorkspaceObjectPermissionLevelsResponseFromPb(pb *getWorkspaceObjectPermissionLevelsResponsePb) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceObjectPermissionLevelsResponse{}

	var permissionLevelsField []WorkspaceObjectPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := workspaceObjectPermissionsDescriptionFromPb(&itemPb)
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

// Get workspace object permissions
type GetWorkspaceObjectPermissionsRequest struct {
	// The workspace object for which to get or manage permissions.
	// Wire name: 'workspace_object_id'
	WorkspaceObjectId string `tf:"-"`
	// The workspace object type for which to get or manage permissions.
	// Wire name: 'workspace_object_type'
	WorkspaceObjectType string `tf:"-"`
}

func getWorkspaceObjectPermissionsRequestToPb(st *GetWorkspaceObjectPermissionsRequest) (*getWorkspaceObjectPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceObjectPermissionsRequestPb{}
	pb.WorkspaceObjectId = st.WorkspaceObjectId

	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
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

type getWorkspaceObjectPermissionsRequestPb struct {
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`
}

func getWorkspaceObjectPermissionsRequestFromPb(pb *getWorkspaceObjectPermissionsRequestPb) (*GetWorkspaceObjectPermissionsRequest, error) {
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
	Content string
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
	Format ImportFormat
	// The language of the object. This value is set only if the object type is
	// `NOTEBOOK`.
	// Wire name: 'language'
	Language Language
	// The flag that specifies whether to overwrite existing object. It is
	// `false` by default. For `DBC` format, `overwrite` is not supported since
	// it may contain a directory.
	// Wire name: 'overwrite'
	Overwrite bool
	// The absolute path of the object or directory. Importing a directory is
	// only supported for the `DBC` and `SOURCE` formats.
	// Wire name: 'path'
	Path string

	ForceSendFields []string `tf:"-"`
}

func importToPb(st *Import) (*importPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &importPb{}
	pb.Content = st.Content

	pb.Format = st.Format

	pb.Language = st.Language

	pb.Overwrite = st.Overwrite

	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type importPb struct {
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

func importFromPb(pb *importPb) (*Import, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Import{}
	st.Content = pb.Content
	st.Format = pb.Format
	st.Language = pb.Language
	st.Overwrite = pb.Overwrite
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *importPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st importPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The format for workspace import and export.
type ImportFormat string
type importFormatPb string

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

// Type always returns ImportFormat to satisfy [pflag.Value] interface
func (f *ImportFormat) Type() string {
	return "ImportFormat"
}

func importFormatToPb(st *ImportFormat) (*importFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := importFormatPb(*st)
	return &pb, nil
}

func importFormatFromPb(pb *importFormatPb) (*ImportFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := ImportFormat(*pb)
	return &st, nil
}

type ImportResponse struct {
}

func importResponseToPb(st *ImportResponse) (*importResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &importResponsePb{}

	return pb, nil
}

func (st *ImportResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &importResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := importResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ImportResponse) MarshalJSON() ([]byte, error) {
	pb, err := importResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type importResponsePb struct {
}

func importResponseFromPb(pb *importResponsePb) (*ImportResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ImportResponse{}

	return st, nil
}

// The language of notebook.
type Language string
type languagePb string

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

func languageToPb(st *Language) (*languagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := languagePb(*st)
	return &pb, nil
}

func languageFromPb(pb *languagePb) (*Language, error) {
	if pb == nil {
		return nil, nil
	}
	st := Language(*pb)
	return &st, nil
}

// Lists ACLs
type ListAclsRequest struct {
	// The name of the scope to fetch ACL information from.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func listAclsRequestToPb(st *ListAclsRequest) (*listAclsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAclsRequestPb{}
	pb.Scope = st.Scope

	return pb, nil
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

type listAclsRequestPb struct {
	// The name of the scope to fetch ACL information from.
	Scope string `json:"-" url:"scope"`
}

func listAclsRequestFromPb(pb *listAclsRequestPb) (*ListAclsRequest, error) {
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
	Items []AclItem
}

func listAclsResponseToPb(st *ListAclsResponse) (*listAclsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAclsResponsePb{}

	var itemsPb []aclItemPb
	for _, item := range st.Items {
		itemPb, err := aclItemToPb(&item)
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

type listAclsResponsePb struct {
	// The associated ACLs rule applied to principals in the given scope.
	Items []aclItemPb `json:"items,omitempty"`
}

func listAclsResponseFromPb(pb *listAclsResponsePb) (*ListAclsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAclsResponse{}

	var itemsField []AclItem
	for _, itemPb := range pb.Items {
		item, err := aclItemFromPb(&itemPb)
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
	Credentials []CredentialInfo
}

func listCredentialsResponseToPb(st *ListCredentialsResponse) (*listCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsResponsePb{}

	var credentialsPb []credentialInfoPb
	for _, item := range st.Credentials {
		itemPb, err := credentialInfoToPb(&item)
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

type listCredentialsResponsePb struct {
	// List of credentials.
	Credentials []credentialInfoPb `json:"credentials,omitempty"`
}

func listCredentialsResponseFromPb(pb *listCredentialsResponsePb) (*ListCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsResponse{}

	var credentialsField []CredentialInfo
	for _, itemPb := range pb.Credentials {
		item, err := credentialInfoFromPb(&itemPb)
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

// Get repos
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
	PathPrefix string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listReposRequestToPb(st *ListReposRequest) (*listReposRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listReposRequestPb{}
	pb.NextPageToken = st.NextPageToken

	pb.PathPrefix = st.PathPrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listReposRequestPb struct {
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

func listReposRequestFromPb(pb *listReposRequestPb) (*ListReposRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListReposRequest{}
	st.NextPageToken = pb.NextPageToken
	st.PathPrefix = pb.PathPrefix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listReposRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listReposRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListReposResponse struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string
	// List of Git folders (repos).
	// Wire name: 'repos'
	Repos []RepoInfo

	ForceSendFields []string `tf:"-"`
}

func listReposResponseToPb(st *ListReposResponse) (*listReposResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listReposResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var reposPb []repoInfoPb
	for _, item := range st.Repos {
		itemPb, err := repoInfoToPb(&item)
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

type listReposResponsePb struct {
	// Token that can be specified as a query parameter to the `GET /repos`
	// endpoint to retrieve the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of Git folders (repos).
	Repos []repoInfoPb `json:"repos,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listReposResponseFromPb(pb *listReposResponsePb) (*ListReposResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListReposResponse{}
	st.NextPageToken = pb.NextPageToken

	var reposField []RepoInfo
	for _, itemPb := range pb.Repos {
		item, err := repoInfoFromPb(&itemPb)
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

func (st *listReposResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listReposResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListResponse struct {
	// List of objects.
	// Wire name: 'objects'
	Objects []ObjectInfo
}

func listResponseToPb(st *ListResponse) (*listResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listResponsePb{}

	var objectsPb []objectInfoPb
	for _, item := range st.Objects {
		itemPb, err := objectInfoToPb(&item)
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

type listResponsePb struct {
	// List of objects.
	Objects []objectInfoPb `json:"objects,omitempty"`
}

func listResponseFromPb(pb *listResponsePb) (*ListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListResponse{}

	var objectsField []ObjectInfo
	for _, itemPb := range pb.Objects {
		item, err := objectInfoFromPb(&itemPb)
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
	Scopes []SecretScope
}

func listScopesResponseToPb(st *ListScopesResponse) (*listScopesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listScopesResponsePb{}

	var scopesPb []secretScopePb
	for _, item := range st.Scopes {
		itemPb, err := secretScopeToPb(&item)
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

type listScopesResponsePb struct {
	// The available secret scopes.
	Scopes []secretScopePb `json:"scopes,omitempty"`
}

func listScopesResponseFromPb(pb *listScopesResponsePb) (*ListScopesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListScopesResponse{}

	var scopesField []SecretScope
	for _, itemPb := range pb.Scopes {
		item, err := secretScopeFromPb(&itemPb)
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

// List secret keys
type ListSecretsRequest struct {
	// The name of the scope to list secrets within.
	// Wire name: 'scope'
	Scope string `tf:"-"`
}

func listSecretsRequestToPb(st *ListSecretsRequest) (*listSecretsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSecretsRequestPb{}
	pb.Scope = st.Scope

	return pb, nil
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

type listSecretsRequestPb struct {
	// The name of the scope to list secrets within.
	Scope string `json:"-" url:"scope"`
}

func listSecretsRequestFromPb(pb *listSecretsRequestPb) (*ListSecretsRequest, error) {
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
	Secrets []SecretMetadata
}

func listSecretsResponseToPb(st *ListSecretsResponse) (*listSecretsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSecretsResponsePb{}

	var secretsPb []secretMetadataPb
	for _, item := range st.Secrets {
		itemPb, err := secretMetadataToPb(&item)
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

type listSecretsResponsePb struct {
	// Metadata information of all secrets contained within the given scope.
	Secrets []secretMetadataPb `json:"secrets,omitempty"`
}

func listSecretsResponseFromPb(pb *listSecretsResponsePb) (*ListSecretsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSecretsResponse{}

	var secretsField []SecretMetadata
	for _, itemPb := range pb.Secrets {
		item, err := secretMetadataFromPb(&itemPb)
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

// List contents
type ListWorkspaceRequest struct {
	// UTC timestamp in milliseconds
	// Wire name: 'notebooks_modified_after'
	NotebooksModifiedAfter int64 `tf:"-"`
	// The absolute path of the notebook or directory.
	// Wire name: 'path'
	Path string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listWorkspaceRequestToPb(st *ListWorkspaceRequest) (*listWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWorkspaceRequestPb{}
	pb.NotebooksModifiedAfter = st.NotebooksModifiedAfter

	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listWorkspaceRequestPb struct {
	// UTC timestamp in milliseconds
	NotebooksModifiedAfter int64 `json:"-" url:"notebooks_modified_after,omitempty"`
	// The absolute path of the notebook or directory.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listWorkspaceRequestFromPb(pb *listWorkspaceRequestPb) (*ListWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWorkspaceRequest{}
	st.NotebooksModifiedAfter = pb.NotebooksModifiedAfter
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Mkdirs struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	// Wire name: 'path'
	Path string
}

func mkdirsToPb(st *Mkdirs) (*mkdirsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkdirsPb{}
	pb.Path = st.Path

	return pb, nil
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

type mkdirsPb struct {
	// The absolute path of the directory. If the parent directories do not
	// exist, it will also create them. If the directory already exists, this
	// command will do nothing and succeed.
	Path string `json:"path"`
}

func mkdirsFromPb(pb *mkdirsPb) (*Mkdirs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Mkdirs{}
	st.Path = pb.Path

	return st, nil
}

type MkdirsResponse struct {
}

func mkdirsResponseToPb(st *MkdirsResponse) (*mkdirsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkdirsResponsePb{}

	return pb, nil
}

func (st *MkdirsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mkdirsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := mkdirsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MkdirsResponse) MarshalJSON() ([]byte, error) {
	pb, err := mkdirsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type mkdirsResponsePb struct {
}

func mkdirsResponseFromPb(pb *mkdirsResponsePb) (*MkdirsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MkdirsResponse{}

	return st, nil
}

// The information of the object in workspace. It will be returned by “list“
// and “get-status“.
type ObjectInfo struct {
	// Only applicable to files. The creation UTC timestamp.
	// Wire name: 'created_at'
	CreatedAt int64
	// The language of the object. This value is set only if the object type is
	// ``NOTEBOOK``.
	// Wire name: 'language'
	Language Language
	// Only applicable to files, the last modified UTC timestamp.
	// Wire name: 'modified_at'
	ModifiedAt int64
	// Unique identifier for the object.
	// Wire name: 'object_id'
	ObjectId int64
	// The type of the object in workspace.
	//
	// - `NOTEBOOK`: document that contains runnable code, visualizations, and
	// explanatory text. - `DIRECTORY`: directory - `LIBRARY`: library - `FILE`:
	// file - `REPO`: repository - `DASHBOARD`: Lakeview dashboard
	// Wire name: 'object_type'
	ObjectType ObjectType
	// The absolute path of the object.
	// Wire name: 'path'
	Path string
	// A unique identifier for the object that is consistent across all
	// Databricks APIs.
	// Wire name: 'resource_id'
	ResourceId string
	// Only applicable to files. The file size in bytes can be returned.
	// Wire name: 'size'
	Size int64

	ForceSendFields []string `tf:"-"`
}

func objectInfoToPb(st *ObjectInfo) (*objectInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &objectInfoPb{}
	pb.CreatedAt = st.CreatedAt

	pb.Language = st.Language

	pb.ModifiedAt = st.ModifiedAt

	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	pb.Path = st.Path

	pb.ResourceId = st.ResourceId

	pb.Size = st.Size

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type objectInfoPb struct {
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

func objectInfoFromPb(pb *objectInfoPb) (*ObjectInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ObjectInfo{}
	st.CreatedAt = pb.CreatedAt
	st.Language = pb.Language
	st.ModifiedAt = pb.ModifiedAt
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType
	st.Path = pb.Path
	st.ResourceId = pb.ResourceId
	st.Size = pb.Size

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *objectInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st objectInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of the object in workspace.
type ObjectType string
type objectTypePb string

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

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (f *ObjectType) Type() string {
	return "ObjectType"
}

func objectTypeToPb(st *ObjectType) (*objectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := objectTypePb(*st)
	return &pb, nil
}

func objectTypeFromPb(pb *objectTypePb) (*ObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ObjectType(*pb)
	return &st, nil
}

type PutAcl struct {
	// The permission level applied to the principal.
	// Wire name: 'permission'
	Permission AclPermission
	// The principal in which the permission is applied.
	// Wire name: 'principal'
	Principal string
	// The name of the scope to apply permissions to.
	// Wire name: 'scope'
	Scope string
}

func putAclToPb(st *PutAcl) (*putAclPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAclPb{}
	pb.Permission = st.Permission

	pb.Principal = st.Principal

	pb.Scope = st.Scope

	return pb, nil
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

type putAclPb struct {
	// The permission level applied to the principal.
	Permission AclPermission `json:"permission"`
	// The principal in which the permission is applied.
	Principal string `json:"principal"`
	// The name of the scope to apply permissions to.
	Scope string `json:"scope"`
}

func putAclFromPb(pb *putAclPb) (*PutAcl, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAcl{}
	st.Permission = pb.Permission
	st.Principal = pb.Principal
	st.Scope = pb.Scope

	return st, nil
}

type PutAclResponse struct {
}

func putAclResponseToPb(st *PutAclResponse) (*putAclResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAclResponsePb{}

	return pb, nil
}

func (st *PutAclResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putAclResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putAclResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutAclResponse) MarshalJSON() ([]byte, error) {
	pb, err := putAclResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type putAclResponsePb struct {
}

func putAclResponseFromPb(pb *putAclResponsePb) (*PutAclResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAclResponse{}

	return st, nil
}

type PutSecret struct {
	// If specified, value will be stored as bytes.
	// Wire name: 'bytes_value'
	BytesValue string
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string
	// The name of the scope to which the secret will be associated with.
	// Wire name: 'scope'
	Scope string
	// If specified, note that the value will be stored in UTF-8 (MB4) form.
	// Wire name: 'string_value'
	StringValue string

	ForceSendFields []string `tf:"-"`
}

func putSecretToPb(st *PutSecret) (*putSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putSecretPb{}
	pb.BytesValue = st.BytesValue

	pb.Key = st.Key

	pb.Scope = st.Scope

	pb.StringValue = st.StringValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type putSecretPb struct {
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

func putSecretFromPb(pb *putSecretPb) (*PutSecret, error) {
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

func (st *putSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st putSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PutSecretResponse struct {
}

func putSecretResponseToPb(st *PutSecretResponse) (*putSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putSecretResponsePb{}

	return pb, nil
}

func (st *PutSecretResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putSecretResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putSecretResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutSecretResponse) MarshalJSON() ([]byte, error) {
	pb, err := putSecretResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type putSecretResponsePb struct {
}

func putSecretResponseFromPb(pb *putSecretResponsePb) (*PutSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutSecretResponse{}

	return st, nil
}

type RepoAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func repoAccessControlRequestToPb(st *RepoAccessControlRequest) (*repoAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type repoAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoAccessControlRequestFromPb(pb *repoAccessControlRequestPb) (*RepoAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repoAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repoAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []RepoPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func repoAccessControlResponseToPb(st *RepoAccessControlResponse) (*repoAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoAccessControlResponsePb{}

	var allPermissionsPb []repoPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := repoPermissionToPb(&item)
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

type repoAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []repoPermissionPb `json:"all_permissions,omitempty"`
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

func repoAccessControlResponseFromPb(pb *repoAccessControlResponsePb) (*RepoAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoAccessControlResponse{}

	var allPermissionsField []RepoPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := repoPermissionFromPb(&itemPb)
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

func (st *repoAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repoAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Git folder (repo) information.
type RepoInfo struct {
	// Name of the current git branch of the git folder (repo).
	// Wire name: 'branch'
	Branch string
	// Current git commit id of the git folder (repo).
	// Wire name: 'head_commit_id'
	HeadCommitId string
	// Id of the git folder (repo) in the Workspace.
	// Wire name: 'id'
	Id int64
	// Root path of the git folder (repo) in the Workspace.
	// Wire name: 'path'
	Path string
	// Git provider of the remote git repository, e.g. `gitHub`.
	// Wire name: 'provider'
	Provider string
	// Sparse checkout config for the git folder (repo).
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckout
	// URL of the remote git repository.
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func repoInfoToPb(st *RepoInfo) (*repoInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoInfoPb{}
	pb.Branch = st.Branch

	pb.HeadCommitId = st.HeadCommitId

	pb.Id = st.Id

	pb.Path = st.Path

	pb.Provider = st.Provider

	sparseCheckoutPb, err := sparseCheckoutToPb(st.SparseCheckout)
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

type repoInfoPb struct {
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
	SparseCheckout *sparseCheckoutPb `json:"sparse_checkout,omitempty"`
	// URL of the remote git repository.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoInfoFromPb(pb *repoInfoPb) (*RepoInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoInfo{}
	st.Branch = pb.Branch
	st.HeadCommitId = pb.HeadCommitId
	st.Id = pb.Id
	st.Path = pb.Path
	st.Provider = pb.Provider
	sparseCheckoutField, err := sparseCheckoutFromPb(pb.SparseCheckout)
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

func (st *repoInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repoInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func repoPermissionToPb(st *RepoPermission) (*repoPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type repoPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoPermissionFromPb(pb *repoPermissionPb) (*RepoPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repoPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repoPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type RepoPermissionLevel string
type repoPermissionLevelPb string

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

func repoPermissionLevelToPb(st *RepoPermissionLevel) (*repoPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := repoPermissionLevelPb(*st)
	return &pb, nil
}

func repoPermissionLevelFromPb(pb *repoPermissionLevelPb) (*RepoPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := RepoPermissionLevel(*pb)
	return &st, nil
}

type RepoPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func repoPermissionsToPb(st *RepoPermissions) (*repoPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoPermissionsPb{}

	var accessControlListPb []repoAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := repoAccessControlResponseToPb(&item)
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

type repoPermissionsPb struct {
	AccessControlList []repoAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoPermissionsFromPb(pb *repoPermissionsPb) (*RepoPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissions{}

	var accessControlListField []RepoAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := repoAccessControlResponseFromPb(&itemPb)
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

func (st *repoPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repoPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RepoPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func repoPermissionsDescriptionToPb(st *RepoPermissionsDescription) (*repoPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type repoPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel RepoPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoPermissionsDescriptionFromPb(pb *repoPermissionsDescriptionPb) (*RepoPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repoPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repoPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepoPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RepoAccessControlRequest
	// The repo for which to get or manage permissions.
	// Wire name: 'repo_id'
	RepoId string `tf:"-"`
}

func repoPermissionsRequestToPb(st *RepoPermissionsRequest) (*repoPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoPermissionsRequestPb{}

	var accessControlListPb []repoAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := repoAccessControlRequestToPb(&item)
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

type repoPermissionsRequestPb struct {
	AccessControlList []repoAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The repo for which to get or manage permissions.
	RepoId string `json:"-" url:"-"`
}

func repoPermissionsRequestFromPb(pb *repoPermissionsRequestPb) (*RepoPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissionsRequest{}

	var accessControlListField []RepoAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := repoAccessControlRequestFromPb(&itemPb)
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

type ScopeBackendType string
type scopeBackendTypePb string

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

func scopeBackendTypeToPb(st *ScopeBackendType) (*scopeBackendTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := scopeBackendTypePb(*st)
	return &pb, nil
}

func scopeBackendTypeFromPb(pb *scopeBackendTypePb) (*ScopeBackendType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ScopeBackendType(*pb)
	return &st, nil
}

type SecretMetadata struct {
	// A unique name to identify the secret.
	// Wire name: 'key'
	Key string
	// The last updated timestamp (in milliseconds) for the secret.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64

	ForceSendFields []string `tf:"-"`
}

func secretMetadataToPb(st *SecretMetadata) (*secretMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &secretMetadataPb{}
	pb.Key = st.Key

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type secretMetadataPb struct {
	// A unique name to identify the secret.
	Key string `json:"key,omitempty"`
	// The last updated timestamp (in milliseconds) for the secret.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func secretMetadataFromPb(pb *secretMetadataPb) (*SecretMetadata, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretMetadata{}
	st.Key = pb.Key
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *secretMetadataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st secretMetadataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SecretScope struct {
	// The type of secret scope backend.
	// Wire name: 'backend_type'
	BackendType ScopeBackendType
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	// Wire name: 'keyvault_metadata'
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata
	// A unique name to identify the secret scope.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func secretScopeToPb(st *SecretScope) (*secretScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &secretScopePb{}
	pb.BackendType = st.BackendType

	keyvaultMetadataPb, err := azureKeyVaultSecretScopeMetadataToPb(st.KeyvaultMetadata)
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

type secretScopePb struct {
	// The type of secret scope backend.
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
	// The metadata for the secret scope if the type is `AZURE_KEYVAULT`
	KeyvaultMetadata *azureKeyVaultSecretScopeMetadataPb `json:"keyvault_metadata,omitempty"`
	// A unique name to identify the secret scope.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func secretScopeFromPb(pb *secretScopePb) (*SecretScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretScope{}
	st.BackendType = pb.BackendType
	keyvaultMetadataField, err := azureKeyVaultSecretScopeMetadataFromPb(pb.KeyvaultMetadata)
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

func (st *secretScopePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st secretScopePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Sparse checkout configuration, it contains options like cone patterns.
type SparseCheckout struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	// Wire name: 'patterns'
	Patterns []string
}

func sparseCheckoutToPb(st *SparseCheckout) (*sparseCheckoutPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparseCheckoutPb{}
	pb.Patterns = st.Patterns

	return pb, nil
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

type sparseCheckoutPb struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns []string `json:"patterns,omitempty"`
}

func sparseCheckoutFromPb(pb *sparseCheckoutPb) (*SparseCheckout, error) {
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
	Patterns []string
}

func sparseCheckoutUpdateToPb(st *SparseCheckoutUpdate) (*sparseCheckoutUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparseCheckoutUpdatePb{}
	pb.Patterns = st.Patterns

	return pb, nil
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

type sparseCheckoutUpdatePb struct {
	// List of sparse checkout cone patterns, see [cone mode handling] for
	// details.
	//
	// [cone mode handling]: https://git-scm.com/docs/git-sparse-checkout#_internalscone_mode_handling
	Patterns []string `json:"patterns,omitempty"`
}

func sparseCheckoutUpdateFromPb(pb *sparseCheckoutUpdatePb) (*SparseCheckoutUpdate, error) {
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
	GitProvider string
	// The username or email provided with your Git provider account, depending
	// on which provider you are using. For GitHub, GitHub Enterprise Server, or
	// Azure DevOps Services, either email or username may be used. For GitLab,
	// GitLab Enterprise Edition, email must be used. For AWS CodeCommit,
	// BitBucket or BitBucket Server, username must be used. For all other
	// providers please see your provider's Personal Access Token authentication
	// documentation to see what is supported.
	// Wire name: 'git_username'
	GitUsername string
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	// Wire name: 'personal_access_token'
	PersonalAccessToken string

	ForceSendFields []string `tf:"-"`
}

func updateCredentialsRequestToPb(st *UpdateCredentialsRequest) (*updateCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	pb.GitProvider = st.GitProvider

	pb.GitUsername = st.GitUsername

	pb.PersonalAccessToken = st.PersonalAccessToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateCredentialsRequestPb struct {
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
	// The personal access token used to authenticate to the corresponding Git
	// provider. For certain providers, support may exist for other types of
	// scoped access tokens. [Learn more].
	//
	// [Learn more]: https://docs.databricks.com/repos/get-access-tokens-from-git-provider.html
	PersonalAccessToken string `json:"personal_access_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateCredentialsRequestFromPb(pb *updateCredentialsRequestPb) (*UpdateCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCredentialsRequest{}
	st.CredentialId = pb.CredentialId
	st.GitProvider = pb.GitProvider
	st.GitUsername = pb.GitUsername
	st.PersonalAccessToken = pb.PersonalAccessToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCredentialsResponse struct {
}

func updateCredentialsResponseToPb(st *UpdateCredentialsResponse) (*updateCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCredentialsResponsePb{}

	return pb, nil
}

func (st *UpdateCredentialsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCredentialsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCredentialsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCredentialsResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateCredentialsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCredentialsResponsePb struct {
}

func updateCredentialsResponseFromPb(pb *updateCredentialsResponsePb) (*UpdateCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCredentialsResponse{}

	return st, nil
}

type UpdateRepoRequest struct {
	// Branch that the local version of the repo is checked out to.
	// Wire name: 'branch'
	Branch string
	// ID of the Git folder (repo) object in the workspace.
	// Wire name: 'repo_id'
	RepoId int64 `tf:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	// Wire name: 'sparse_checkout'
	SparseCheckout *SparseCheckoutUpdate
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	// Wire name: 'tag'
	Tag string

	ForceSendFields []string `tf:"-"`
}

func updateRepoRequestToPb(st *UpdateRepoRequest) (*updateRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRepoRequestPb{}
	pb.Branch = st.Branch

	pb.RepoId = st.RepoId

	sparseCheckoutPb, err := sparseCheckoutUpdateToPb(st.SparseCheckout)
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

type updateRepoRequestPb struct {
	// Branch that the local version of the repo is checked out to.
	Branch string `json:"branch,omitempty"`
	// ID of the Git folder (repo) object in the workspace.
	RepoId int64 `json:"-" url:"-"`
	// If specified, update the sparse checkout settings. The update will fail
	// if sparse checkout is not enabled for the repo.
	SparseCheckout *sparseCheckoutUpdatePb `json:"sparse_checkout,omitempty"`
	// Tag that the local version of the repo is checked out to. Updating the
	// repo to a tag puts the repo in a detached HEAD state. Before committing
	// new changes, you must update the repo to a branch instead of the detached
	// HEAD.
	Tag string `json:"tag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRepoRequestFromPb(pb *updateRepoRequestPb) (*UpdateRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRepoRequest{}
	st.Branch = pb.Branch
	st.RepoId = pb.RepoId
	sparseCheckoutField, err := sparseCheckoutUpdateFromPb(pb.SparseCheckout)
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

func (st *updateRepoRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRepoRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRepoResponse struct {
}

func updateRepoResponseToPb(st *UpdateRepoResponse) (*updateRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRepoResponsePb{}

	return pb, nil
}

func (st *UpdateRepoResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRepoResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRepoResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRepoResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateRepoResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateRepoResponsePb struct {
}

func updateRepoResponseFromPb(pb *updateRepoResponsePb) (*UpdateRepoResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRepoResponse{}

	return st, nil
}

type WorkspaceObjectAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func workspaceObjectAccessControlRequestToPb(st *WorkspaceObjectAccessControlRequest) (*workspaceObjectAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type workspaceObjectAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceObjectAccessControlRequestFromPb(pb *workspaceObjectAccessControlRequestPb) (*WorkspaceObjectAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceObjectAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceObjectAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []WorkspaceObjectPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func workspaceObjectAccessControlResponseToPb(st *WorkspaceObjectAccessControlResponse) (*workspaceObjectAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectAccessControlResponsePb{}

	var allPermissionsPb []workspaceObjectPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := workspaceObjectPermissionToPb(&item)
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

type workspaceObjectAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []workspaceObjectPermissionPb `json:"all_permissions,omitempty"`
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

func workspaceObjectAccessControlResponseFromPb(pb *workspaceObjectAccessControlResponsePb) (*WorkspaceObjectAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectAccessControlResponse{}

	var allPermissionsField []WorkspaceObjectPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := workspaceObjectPermissionFromPb(&itemPb)
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

func (st *workspaceObjectAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceObjectAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func workspaceObjectPermissionToPb(st *WorkspaceObjectPermission) (*workspaceObjectPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type workspaceObjectPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceObjectPermissionFromPb(pb *workspaceObjectPermissionPb) (*WorkspaceObjectPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceObjectPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceObjectPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type WorkspaceObjectPermissionLevel string
type workspaceObjectPermissionLevelPb string

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

func workspaceObjectPermissionLevelToPb(st *WorkspaceObjectPermissionLevel) (*workspaceObjectPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := workspaceObjectPermissionLevelPb(*st)
	return &pb, nil
}

func workspaceObjectPermissionLevelFromPb(pb *workspaceObjectPermissionLevelPb) (*WorkspaceObjectPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspaceObjectPermissionLevel(*pb)
	return &st, nil
}

type WorkspaceObjectPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []WorkspaceObjectAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func workspaceObjectPermissionsToPb(st *WorkspaceObjectPermissions) (*workspaceObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectPermissionsPb{}

	var accessControlListPb []workspaceObjectAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := workspaceObjectAccessControlResponseToPb(&item)
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

type workspaceObjectPermissionsPb struct {
	AccessControlList []workspaceObjectAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceObjectPermissionsFromPb(pb *workspaceObjectPermissionsPb) (*WorkspaceObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissions{}

	var accessControlListField []WorkspaceObjectAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := workspaceObjectAccessControlResponseFromPb(&itemPb)
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

func (st *workspaceObjectPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceObjectPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel WorkspaceObjectPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func workspaceObjectPermissionsDescriptionToPb(st *WorkspaceObjectPermissionsDescription) (*workspaceObjectPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type workspaceObjectPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceObjectPermissionsDescriptionFromPb(pb *workspaceObjectPermissionsDescriptionPb) (*WorkspaceObjectPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceObjectPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceObjectPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WorkspaceObjectPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []WorkspaceObjectAccessControlRequest
	// The workspace object for which to get or manage permissions.
	// Wire name: 'workspace_object_id'
	WorkspaceObjectId string `tf:"-"`
	// The workspace object type for which to get or manage permissions.
	// Wire name: 'workspace_object_type'
	WorkspaceObjectType string `tf:"-"`
}

func workspaceObjectPermissionsRequestToPb(st *WorkspaceObjectPermissionsRequest) (*workspaceObjectPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectPermissionsRequestPb{}

	var accessControlListPb []workspaceObjectAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := workspaceObjectAccessControlRequestToPb(&item)
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

type workspaceObjectPermissionsRequestPb struct {
	AccessControlList []workspaceObjectAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The workspace object for which to get or manage permissions.
	WorkspaceObjectId string `json:"-" url:"-"`
	// The workspace object type for which to get or manage permissions.
	WorkspaceObjectType string `json:"-" url:"-"`
}

func workspaceObjectPermissionsRequestFromPb(pb *workspaceObjectPermissionsRequestPb) (*WorkspaceObjectPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissionsRequest{}

	var accessControlListField []WorkspaceObjectAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := workspaceObjectAccessControlRequestFromPb(&itemPb)
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
	s := fmt.Sprintf("%fs", d.Seconds())
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
