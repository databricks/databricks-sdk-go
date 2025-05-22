// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspace

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func aclItemToPb(st *AclItem) (*aclItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aclItemPb{}
	pb.Permission = st.Permission
	pb.Principal = st.Principal

	return pb, nil
}

type aclItemPb struct {
	Permission AclPermission `json:"permission"`
	Principal  string        `json:"principal"`
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

func azureKeyVaultSecretScopeMetadataToPb(st *AzureKeyVaultSecretScopeMetadata) (*azureKeyVaultSecretScopeMetadataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureKeyVaultSecretScopeMetadataPb{}
	pb.DnsName = st.DnsName
	pb.ResourceId = st.ResourceId

	return pb, nil
}

type azureKeyVaultSecretScopeMetadataPb struct {
	DnsName    string `json:"dns_name"`
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

type createCredentialsRequestPb struct {
	GitProvider         string `json:"git_provider"`
	GitUsername         string `json:"git_username,omitempty"`
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

type createCredentialsResponsePb struct {
	CredentialId int64  `json:"credential_id"`
	GitProvider  string `json:"git_provider"`
	GitUsername  string `json:"git_username,omitempty"`

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

func createRepoRequestToPb(st *CreateRepoRequest) (*createRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRepoRequestPb{}
	pb.Path = st.Path
	pb.Provider = st.Provider
	pb.SparseCheckout = st.SparseCheckout
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createRepoRequestPb struct {
	Path           string          `json:"path,omitempty"`
	Provider       string          `json:"provider"`
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	Url            string          `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRepoRequestFromPb(pb *createRepoRequestPb) (*CreateRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRepoRequest{}
	st.Path = pb.Path
	st.Provider = pb.Provider
	st.SparseCheckout = pb.SparseCheckout
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
	pb.SparseCheckout = st.SparseCheckout
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createRepoResponsePb struct {
	Branch         string          `json:"branch,omitempty"`
	HeadCommitId   string          `json:"head_commit_id,omitempty"`
	Id             int64           `json:"id,omitempty"`
	Path           string          `json:"path,omitempty"`
	Provider       string          `json:"provider,omitempty"`
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	Url            string          `json:"url,omitempty"`

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
	st.SparseCheckout = pb.SparseCheckout
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

func createScopeToPb(st *CreateScope) (*createScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createScopePb{}
	pb.BackendAzureKeyvault = st.BackendAzureKeyvault
	pb.InitialManagePrincipal = st.InitialManagePrincipal
	pb.Scope = st.Scope
	pb.ScopeBackendType = st.ScopeBackendType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createScopePb struct {
	BackendAzureKeyvault   *AzureKeyVaultSecretScopeMetadata `json:"backend_azure_keyvault,omitempty"`
	InitialManagePrincipal string                            `json:"initial_manage_principal,omitempty"`
	Scope                  string                            `json:"scope"`
	ScopeBackendType       ScopeBackendType                  `json:"scope_backend_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createScopeFromPb(pb *createScopePb) (*CreateScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateScope{}
	st.BackendAzureKeyvault = pb.BackendAzureKeyvault
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

func createScopeResponseToPb(st *CreateScopeResponse) (*createScopeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createScopeResponsePb{}

	return pb, nil
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

type credentialInfoPb struct {
	CredentialId int64  `json:"credential_id"`
	GitProvider  string `json:"git_provider,omitempty"`
	GitUsername  string `json:"git_username,omitempty"`

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

type deletePb struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive,omitempty"`

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

func deleteAclToPb(st *DeleteAcl) (*deleteAclPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAclPb{}
	pb.Principal = st.Principal
	pb.Scope = st.Scope

	return pb, nil
}

type deleteAclPb struct {
	Principal string `json:"principal"`
	Scope     string `json:"scope"`
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

func deleteAclResponseToPb(st *DeleteAclResponse) (*deleteAclResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAclResponsePb{}

	return pb, nil
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

func deleteCredentialsRequestToPb(st *DeleteCredentialsRequest) (*deleteCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	return pb, nil
}

type deleteCredentialsRequestPb struct {
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

func deleteCredentialsResponseToPb(st *DeleteCredentialsResponse) (*deleteCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialsResponsePb{}

	return pb, nil
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

func deleteRepoRequestToPb(st *DeleteRepoRequest) (*deleteRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRepoRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

type deleteRepoRequestPb struct {
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

func deleteRepoResponseToPb(st *DeleteRepoResponse) (*deleteRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRepoResponsePb{}

	return pb, nil
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

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

func deleteScopeToPb(st *DeleteScope) (*deleteScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScopePb{}
	pb.Scope = st.Scope

	return pb, nil
}

type deleteScopePb struct {
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

func deleteScopeResponseToPb(st *DeleteScopeResponse) (*deleteScopeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteScopeResponsePb{}

	return pb, nil
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

func deleteSecretToPb(st *DeleteSecret) (*deleteSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSecretPb{}
	pb.Key = st.Key
	pb.Scope = st.Scope

	return pb, nil
}

type deleteSecretPb struct {
	Key   string `json:"key"`
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

func deleteSecretResponseToPb(st *DeleteSecretResponse) (*deleteSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSecretResponsePb{}

	return pb, nil
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

func exportRequestToPb(st *ExportRequest) (*exportRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportRequestPb{}
	pb.Format = st.Format
	pb.Path = st.Path

	return pb, nil
}

type exportRequestPb struct {
	Format ExportFormat `json:"-" url:"format,omitempty"`
	Path   string       `json:"-" url:"path"`
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

type exportResponsePb struct {
	Content  string `json:"content,omitempty"`
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

func getAclRequestToPb(st *GetAclRequest) (*getAclRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAclRequestPb{}
	pb.Principal = st.Principal
	pb.Scope = st.Scope

	return pb, nil
}

type getAclRequestPb struct {
	Principal string `json:"-" url:"principal"`
	Scope     string `json:"-" url:"scope"`
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

func getCredentialsRequestToPb(st *GetCredentialsRequest) (*getCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsRequestPb{}
	pb.CredentialId = st.CredentialId

	return pb, nil
}

type getCredentialsRequestPb struct {
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

type getCredentialsResponsePb struct {
	CredentialId int64  `json:"credential_id"`
	GitProvider  string `json:"git_provider,omitempty"`
	GitUsername  string `json:"git_username,omitempty"`

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

func getRepoPermissionLevelsRequestToPb(st *GetRepoPermissionLevelsRequest) (*getRepoPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoPermissionLevelsRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

type getRepoPermissionLevelsRequestPb struct {
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

func getRepoPermissionLevelsResponseToPb(st *GetRepoPermissionLevelsResponse) (*getRepoPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getRepoPermissionLevelsResponsePb struct {
	PermissionLevels []RepoPermissionsDescription `json:"permission_levels,omitempty"`
}

func getRepoPermissionLevelsResponseFromPb(pb *getRepoPermissionLevelsResponsePb) (*GetRepoPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRepoPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getRepoPermissionsRequestToPb(st *GetRepoPermissionsRequest) (*getRepoPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoPermissionsRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

type getRepoPermissionsRequestPb struct {
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

func getRepoRequestToPb(st *GetRepoRequest) (*getRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRepoRequestPb{}
	pb.RepoId = st.RepoId

	return pb, nil
}

type getRepoRequestPb struct {
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
	pb.SparseCheckout = st.SparseCheckout
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getRepoResponsePb struct {
	Branch         string          `json:"branch,omitempty"`
	HeadCommitId   string          `json:"head_commit_id,omitempty"`
	Id             int64           `json:"id,omitempty"`
	Path           string          `json:"path,omitempty"`
	Provider       string          `json:"provider,omitempty"`
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	Url            string          `json:"url,omitempty"`

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
	st.SparseCheckout = pb.SparseCheckout
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

func getSecretRequestToPb(st *GetSecretRequest) (*getSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSecretRequestPb{}
	pb.Key = st.Key
	pb.Scope = st.Scope

	return pb, nil
}

type getSecretRequestPb struct {
	Key   string `json:"-" url:"key"`
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

type getSecretResponsePb struct {
	Key   string `json:"key,omitempty"`
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

func getStatusRequestToPb(st *GetStatusRequest) (*getStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatusRequestPb{}
	pb.Path = st.Path

	return pb, nil
}

type getStatusRequestPb struct {
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

func getWorkspaceObjectPermissionLevelsRequestToPb(st *GetWorkspaceObjectPermissionLevelsRequest) (*getWorkspaceObjectPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceObjectPermissionLevelsRequestPb{}
	pb.WorkspaceObjectId = st.WorkspaceObjectId
	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
}

type getWorkspaceObjectPermissionLevelsRequestPb struct {
	WorkspaceObjectId   string `json:"-" url:"-"`
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

func getWorkspaceObjectPermissionLevelsResponseToPb(st *GetWorkspaceObjectPermissionLevelsResponse) (*getWorkspaceObjectPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceObjectPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getWorkspaceObjectPermissionLevelsResponsePb struct {
	PermissionLevels []WorkspaceObjectPermissionsDescription `json:"permission_levels,omitempty"`
}

func getWorkspaceObjectPermissionLevelsResponseFromPb(pb *getWorkspaceObjectPermissionLevelsResponsePb) (*GetWorkspaceObjectPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceObjectPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
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

type getWorkspaceObjectPermissionsRequestPb struct {
	WorkspaceObjectId   string `json:"-" url:"-"`
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

type importPb struct {
	Content   string       `json:"content,omitempty"`
	Format    ImportFormat `json:"format,omitempty"`
	Language  Language     `json:"language,omitempty"`
	Overwrite bool         `json:"overwrite,omitempty"`
	Path      string       `json:"path"`

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

func importResponseToPb(st *ImportResponse) (*importResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &importResponsePb{}

	return pb, nil
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

func listAclsRequestToPb(st *ListAclsRequest) (*listAclsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAclsRequestPb{}
	pb.Scope = st.Scope

	return pb, nil
}

type listAclsRequestPb struct {
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

func listAclsResponseToPb(st *ListAclsResponse) (*listAclsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAclsResponsePb{}
	pb.Items = st.Items

	return pb, nil
}

type listAclsResponsePb struct {
	Items []AclItem `json:"items,omitempty"`
}

func listAclsResponseFromPb(pb *listAclsResponsePb) (*ListAclsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAclsResponse{}
	st.Items = pb.Items

	return st, nil
}

func listCredentialsResponseToPb(st *ListCredentialsResponse) (*listCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsResponsePb{}
	pb.Credentials = st.Credentials

	return pb, nil
}

type listCredentialsResponsePb struct {
	Credentials []CredentialInfo `json:"credentials,omitempty"`
}

func listCredentialsResponseFromPb(pb *listCredentialsResponsePb) (*ListCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsResponse{}
	st.Credentials = pb.Credentials

	return st, nil
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

type listReposRequestPb struct {
	NextPageToken string `json:"-" url:"next_page_token,omitempty"`
	PathPrefix    string `json:"-" url:"path_prefix,omitempty"`

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

func listReposResponseToPb(st *ListReposResponse) (*listReposResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listReposResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Repos = st.Repos

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listReposResponsePb struct {
	NextPageToken string     `json:"next_page_token,omitempty"`
	Repos         []RepoInfo `json:"repos,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listReposResponseFromPb(pb *listReposResponsePb) (*ListReposResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListReposResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Repos = pb.Repos

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listReposResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listReposResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listResponseToPb(st *ListResponse) (*listResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listResponsePb{}
	pb.Objects = st.Objects

	return pb, nil
}

type listResponsePb struct {
	Objects []ObjectInfo `json:"objects,omitempty"`
}

func listResponseFromPb(pb *listResponsePb) (*ListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListResponse{}
	st.Objects = pb.Objects

	return st, nil
}

func listScopesResponseToPb(st *ListScopesResponse) (*listScopesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listScopesResponsePb{}
	pb.Scopes = st.Scopes

	return pb, nil
}

type listScopesResponsePb struct {
	Scopes []SecretScope `json:"scopes,omitempty"`
}

func listScopesResponseFromPb(pb *listScopesResponsePb) (*ListScopesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListScopesResponse{}
	st.Scopes = pb.Scopes

	return st, nil
}

func listSecretsRequestToPb(st *ListSecretsRequest) (*listSecretsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSecretsRequestPb{}
	pb.Scope = st.Scope

	return pb, nil
}

type listSecretsRequestPb struct {
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

func listSecretsResponseToPb(st *ListSecretsResponse) (*listSecretsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSecretsResponsePb{}
	pb.Secrets = st.Secrets

	return pb, nil
}

type listSecretsResponsePb struct {
	Secrets []SecretMetadata `json:"secrets,omitempty"`
}

func listSecretsResponseFromPb(pb *listSecretsResponsePb) (*ListSecretsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSecretsResponse{}
	st.Secrets = pb.Secrets

	return st, nil
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

type listWorkspaceRequestPb struct {
	NotebooksModifiedAfter int64  `json:"-" url:"notebooks_modified_after,omitempty"`
	Path                   string `json:"-" url:"path"`

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

func mkdirsToPb(st *Mkdirs) (*mkdirsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkdirsPb{}
	pb.Path = st.Path

	return pb, nil
}

type mkdirsPb struct {
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

func mkdirsResponseToPb(st *MkdirsResponse) (*mkdirsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkdirsResponsePb{}

	return pb, nil
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

type objectInfoPb struct {
	CreatedAt  int64      `json:"created_at,omitempty"`
	Language   Language   `json:"language,omitempty"`
	ModifiedAt int64      `json:"modified_at,omitempty"`
	ObjectId   int64      `json:"object_id,omitempty"`
	ObjectType ObjectType `json:"object_type,omitempty"`
	Path       string     `json:"path,omitempty"`
	ResourceId string     `json:"resource_id,omitempty"`
	Size       int64      `json:"size,omitempty"`

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

type putAclPb struct {
	Permission AclPermission `json:"permission"`
	Principal  string        `json:"principal"`
	Scope      string        `json:"scope"`
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

func putAclResponseToPb(st *PutAclResponse) (*putAclResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAclResponsePb{}

	return pb, nil
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

type putSecretPb struct {
	BytesValue  string `json:"bytes_value,omitempty"`
	Key         string `json:"key"`
	Scope       string `json:"scope"`
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

func putSecretResponseToPb(st *PutSecretResponse) (*putSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putSecretResponsePb{}

	return pb, nil
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

type repoAccessControlRequestPb struct {
	GroupName            string              `json:"group_name,omitempty"`
	PermissionLevel      RepoPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string              `json:"service_principal_name,omitempty"`
	UserName             string              `json:"user_name,omitempty"`

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

func repoAccessControlResponseToPb(st *RepoAccessControlResponse) (*repoAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type repoAccessControlResponsePb struct {
	AllPermissions       []RepoPermission `json:"all_permissions,omitempty"`
	DisplayName          string           `json:"display_name,omitempty"`
	GroupName            string           `json:"group_name,omitempty"`
	ServicePrincipalName string           `json:"service_principal_name,omitempty"`
	UserName             string           `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoAccessControlResponseFromPb(pb *repoAccessControlResponsePb) (*RepoAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
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
	pb.SparseCheckout = st.SparseCheckout
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type repoInfoPb struct {
	Branch         string          `json:"branch,omitempty"`
	HeadCommitId   string          `json:"head_commit_id,omitempty"`
	Id             int64           `json:"id,omitempty"`
	Path           string          `json:"path,omitempty"`
	Provider       string          `json:"provider,omitempty"`
	SparseCheckout *SparseCheckout `json:"sparse_checkout,omitempty"`
	Url            string          `json:"url,omitempty"`

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
	st.SparseCheckout = pb.SparseCheckout
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

type repoPermissionPb struct {
	Inherited           bool                `json:"inherited,omitempty"`
	InheritedFromObject []string            `json:"inherited_from_object,omitempty"`
	PermissionLevel     RepoPermissionLevel `json:"permission_level,omitempty"`

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

func repoPermissionsToPb(st *RepoPermissions) (*repoPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type repoPermissionsPb struct {
	AccessControlList []RepoAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                      `json:"object_id,omitempty"`
	ObjectType        string                      `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repoPermissionsFromPb(pb *repoPermissionsPb) (*RepoPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissions{}
	st.AccessControlList = pb.AccessControlList
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

type repoPermissionsDescriptionPb struct {
	Description     string              `json:"description,omitempty"`
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

func repoPermissionsRequestToPb(st *RepoPermissionsRequest) (*repoPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repoPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.RepoId = st.RepoId

	return pb, nil
}

type repoPermissionsRequestPb struct {
	AccessControlList []RepoAccessControlRequest `json:"access_control_list,omitempty"`
	RepoId            string                     `json:"-" url:"-"`
}

func repoPermissionsRequestFromPb(pb *repoPermissionsRequestPb) (*RepoPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepoPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.RepoId = pb.RepoId

	return st, nil
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

type secretMetadataPb struct {
	Key                  string `json:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`

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

func secretScopeToPb(st *SecretScope) (*secretScopePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &secretScopePb{}
	pb.BackendType = st.BackendType
	pb.KeyvaultMetadata = st.KeyvaultMetadata
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type secretScopePb struct {
	BackendType      ScopeBackendType                  `json:"backend_type,omitempty"`
	KeyvaultMetadata *AzureKeyVaultSecretScopeMetadata `json:"keyvault_metadata,omitempty"`
	Name             string                            `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func secretScopeFromPb(pb *secretScopePb) (*SecretScope, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretScope{}
	st.BackendType = pb.BackendType
	st.KeyvaultMetadata = pb.KeyvaultMetadata
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

func sparseCheckoutToPb(st *SparseCheckout) (*sparseCheckoutPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparseCheckoutPb{}
	pb.Patterns = st.Patterns

	return pb, nil
}

type sparseCheckoutPb struct {
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

func sparseCheckoutUpdateToPb(st *SparseCheckoutUpdate) (*sparseCheckoutUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparseCheckoutUpdatePb{}
	pb.Patterns = st.Patterns

	return pb, nil
}

type sparseCheckoutUpdatePb struct {
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

type updateCredentialsRequestPb struct {
	CredentialId        int64  `json:"-" url:"-"`
	GitProvider         string `json:"git_provider"`
	GitUsername         string `json:"git_username,omitempty"`
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

func updateCredentialsResponseToPb(st *UpdateCredentialsResponse) (*updateCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCredentialsResponsePb{}

	return pb, nil
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

func updateRepoRequestToPb(st *UpdateRepoRequest) (*updateRepoRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRepoRequestPb{}
	pb.Branch = st.Branch
	pb.RepoId = st.RepoId
	pb.SparseCheckout = st.SparseCheckout
	pb.Tag = st.Tag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateRepoRequestPb struct {
	Branch         string                `json:"branch,omitempty"`
	RepoId         int64                 `json:"-" url:"-"`
	SparseCheckout *SparseCheckoutUpdate `json:"sparse_checkout,omitempty"`
	Tag            string                `json:"tag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRepoRequestFromPb(pb *updateRepoRequestPb) (*UpdateRepoRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRepoRequest{}
	st.Branch = pb.Branch
	st.RepoId = pb.RepoId
	st.SparseCheckout = pb.SparseCheckout
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

func updateRepoResponseToPb(st *UpdateRepoResponse) (*updateRepoResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRepoResponsePb{}

	return pb, nil
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

type workspaceObjectAccessControlRequestPb struct {
	GroupName            string                         `json:"group_name,omitempty"`
	PermissionLevel      WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                         `json:"service_principal_name,omitempty"`
	UserName             string                         `json:"user_name,omitempty"`

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

func workspaceObjectAccessControlResponseToPb(st *WorkspaceObjectAccessControlResponse) (*workspaceObjectAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type workspaceObjectAccessControlResponsePb struct {
	AllPermissions       []WorkspaceObjectPermission `json:"all_permissions,omitempty"`
	DisplayName          string                      `json:"display_name,omitempty"`
	GroupName            string                      `json:"group_name,omitempty"`
	ServicePrincipalName string                      `json:"service_principal_name,omitempty"`
	UserName             string                      `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceObjectAccessControlResponseFromPb(pb *workspaceObjectAccessControlResponsePb) (*WorkspaceObjectAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
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

type workspaceObjectPermissionPb struct {
	Inherited           bool                           `json:"inherited,omitempty"`
	InheritedFromObject []string                       `json:"inherited_from_object,omitempty"`
	PermissionLevel     WorkspaceObjectPermissionLevel `json:"permission_level,omitempty"`

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

func workspaceObjectPermissionsToPb(st *WorkspaceObjectPermissions) (*workspaceObjectPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type workspaceObjectPermissionsPb struct {
	AccessControlList []WorkspaceObjectAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                                 `json:"object_id,omitempty"`
	ObjectType        string                                 `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceObjectPermissionsFromPb(pb *workspaceObjectPermissionsPb) (*WorkspaceObjectPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissions{}
	st.AccessControlList = pb.AccessControlList
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

type workspaceObjectPermissionsDescriptionPb struct {
	Description     string                         `json:"description,omitempty"`
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

func workspaceObjectPermissionsRequestToPb(st *WorkspaceObjectPermissionsRequest) (*workspaceObjectPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceObjectPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.WorkspaceObjectId = st.WorkspaceObjectId
	pb.WorkspaceObjectType = st.WorkspaceObjectType

	return pb, nil
}

type workspaceObjectPermissionsRequestPb struct {
	AccessControlList   []WorkspaceObjectAccessControlRequest `json:"access_control_list,omitempty"`
	WorkspaceObjectId   string                                `json:"-" url:"-"`
	WorkspaceObjectType string                                `json:"-" url:"-"`
}

func workspaceObjectPermissionsRequestFromPb(pb *workspaceObjectPermissionsRequestPb) (*WorkspaceObjectPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceObjectPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
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
