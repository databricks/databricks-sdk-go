// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package catalog

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func accountsCreateMetastoreToPb(st *AccountsCreateMetastore) (*accountsCreateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsCreateMetastorePb{}
	pb.MetastoreInfo = st.MetastoreInfo

	return pb, nil
}

type accountsCreateMetastorePb struct {
	MetastoreInfo *CreateMetastore `json:"metastore_info,omitempty"`
}

func accountsCreateMetastoreFromPb(pb *accountsCreateMetastorePb) (*AccountsCreateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateMetastore{}
	st.MetastoreInfo = pb.MetastoreInfo

	return st, nil
}

func accountsCreateMetastoreAssignmentToPb(st *AccountsCreateMetastoreAssignment) (*accountsCreateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsCreateMetastoreAssignmentPb{}
	pb.MetastoreAssignment = st.MetastoreAssignment
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type accountsCreateMetastoreAssignmentPb struct {
	MetastoreAssignment *CreateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	MetastoreId         string                     `json:"-" url:"-"`
	WorkspaceId         int64                      `json:"-" url:"-"`
}

func accountsCreateMetastoreAssignmentFromPb(pb *accountsCreateMetastoreAssignmentPb) (*AccountsCreateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateMetastoreAssignment{}
	st.MetastoreAssignment = pb.MetastoreAssignment
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func accountsCreateStorageCredentialToPb(st *AccountsCreateStorageCredential) (*accountsCreateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsCreateStorageCredentialPb{}
	pb.CredentialInfo = st.CredentialInfo
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

type accountsCreateStorageCredentialPb struct {
	CredentialInfo *CreateStorageCredential `json:"credential_info,omitempty"`
	MetastoreId    string                   `json:"-" url:"-"`
}

func accountsCreateStorageCredentialFromPb(pb *accountsCreateStorageCredentialPb) (*AccountsCreateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsCreateStorageCredential{}
	st.CredentialInfo = pb.CredentialInfo
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

func accountsMetastoreAssignmentToPb(st *AccountsMetastoreAssignment) (*accountsMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsMetastoreAssignmentPb{}
	pb.MetastoreAssignment = st.MetastoreAssignment

	return pb, nil
}

type accountsMetastoreAssignmentPb struct {
	MetastoreAssignment *MetastoreAssignment `json:"metastore_assignment,omitempty"`
}

func accountsMetastoreAssignmentFromPb(pb *accountsMetastoreAssignmentPb) (*AccountsMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsMetastoreAssignment{}
	st.MetastoreAssignment = pb.MetastoreAssignment

	return st, nil
}

func accountsMetastoreInfoToPb(st *AccountsMetastoreInfo) (*accountsMetastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsMetastoreInfoPb{}
	pb.MetastoreInfo = st.MetastoreInfo

	return pb, nil
}

type accountsMetastoreInfoPb struct {
	MetastoreInfo *MetastoreInfo `json:"metastore_info,omitempty"`
}

func accountsMetastoreInfoFromPb(pb *accountsMetastoreInfoPb) (*AccountsMetastoreInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsMetastoreInfo{}
	st.MetastoreInfo = pb.MetastoreInfo

	return st, nil
}

func accountsStorageCredentialInfoToPb(st *AccountsStorageCredentialInfo) (*accountsStorageCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsStorageCredentialInfoPb{}
	pb.CredentialInfo = st.CredentialInfo

	return pb, nil
}

type accountsStorageCredentialInfoPb struct {
	CredentialInfo *StorageCredentialInfo `json:"credential_info,omitempty"`
}

func accountsStorageCredentialInfoFromPb(pb *accountsStorageCredentialInfoPb) (*AccountsStorageCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsStorageCredentialInfo{}
	st.CredentialInfo = pb.CredentialInfo

	return st, nil
}

func accountsUpdateMetastoreToPb(st *AccountsUpdateMetastore) (*accountsUpdateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateMetastorePb{}
	pb.MetastoreId = st.MetastoreId
	pb.MetastoreInfo = st.MetastoreInfo

	return pb, nil
}

type accountsUpdateMetastorePb struct {
	MetastoreId   string           `json:"-" url:"-"`
	MetastoreInfo *UpdateMetastore `json:"metastore_info,omitempty"`
}

func accountsUpdateMetastoreFromPb(pb *accountsUpdateMetastorePb) (*AccountsUpdateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateMetastore{}
	st.MetastoreId = pb.MetastoreId
	st.MetastoreInfo = pb.MetastoreInfo

	return st, nil
}

func accountsUpdateMetastoreAssignmentToPb(st *AccountsUpdateMetastoreAssignment) (*accountsUpdateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateMetastoreAssignmentPb{}
	pb.MetastoreAssignment = st.MetastoreAssignment
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type accountsUpdateMetastoreAssignmentPb struct {
	MetastoreAssignment *UpdateMetastoreAssignment `json:"metastore_assignment,omitempty"`
	MetastoreId         string                     `json:"-" url:"-"`
	WorkspaceId         int64                      `json:"-" url:"-"`
}

func accountsUpdateMetastoreAssignmentFromPb(pb *accountsUpdateMetastoreAssignmentPb) (*AccountsUpdateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateMetastoreAssignment{}
	st.MetastoreAssignment = pb.MetastoreAssignment
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func accountsUpdateStorageCredentialToPb(st *AccountsUpdateStorageCredential) (*accountsUpdateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountsUpdateStorageCredentialPb{}
	pb.CredentialInfo = st.CredentialInfo
	pb.MetastoreId = st.MetastoreId
	pb.StorageCredentialName = st.StorageCredentialName

	return pb, nil
}

type accountsUpdateStorageCredentialPb struct {
	CredentialInfo        *UpdateStorageCredential `json:"credential_info,omitempty"`
	MetastoreId           string                   `json:"-" url:"-"`
	StorageCredentialName string                   `json:"-" url:"-"`
}

func accountsUpdateStorageCredentialFromPb(pb *accountsUpdateStorageCredentialPb) (*AccountsUpdateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountsUpdateStorageCredential{}
	st.CredentialInfo = pb.CredentialInfo
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	return st, nil
}

func artifactAllowlistInfoToPb(st *ArtifactAllowlistInfo) (*artifactAllowlistInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactAllowlistInfoPb{}
	pb.ArtifactMatchers = st.ArtifactMatchers
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type artifactAllowlistInfoPb struct {
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers,omitempty"`
	CreatedAt        int64             `json:"created_at,omitempty"`
	CreatedBy        string            `json:"created_by,omitempty"`
	MetastoreId      string            `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func artifactAllowlistInfoFromPb(pb *artifactAllowlistInfoPb) (*ArtifactAllowlistInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactAllowlistInfo{}
	st.ArtifactMatchers = pb.ArtifactMatchers
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *artifactAllowlistInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st artifactAllowlistInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func artifactMatcherToPb(st *ArtifactMatcher) (*artifactMatcherPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactMatcherPb{}
	pb.Artifact = st.Artifact
	pb.MatchType = st.MatchType

	return pb, nil
}

type artifactMatcherPb struct {
	Artifact  string    `json:"artifact"`
	MatchType MatchType `json:"match_type"`
}

func artifactMatcherFromPb(pb *artifactMatcherPb) (*ArtifactMatcher, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactMatcher{}
	st.Artifact = pb.Artifact
	st.MatchType = pb.MatchType

	return st, nil
}

func assignResponseToPb(st *AssignResponse) (*assignResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &assignResponsePb{}

	return pb, nil
}

type assignResponsePb struct {
}

func assignResponseFromPb(pb *assignResponsePb) (*AssignResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AssignResponse{}

	return st, nil
}

func awsCredentialsToPb(st *AwsCredentials) (*awsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsCredentialsPb{}
	pb.AccessKeyId = st.AccessKeyId
	pb.AccessPoint = st.AccessPoint
	pb.SecretAccessKey = st.SecretAccessKey
	pb.SessionToken = st.SessionToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type awsCredentialsPb struct {
	AccessKeyId     string `json:"access_key_id,omitempty"`
	AccessPoint     string `json:"access_point,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	SessionToken    string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsCredentialsFromPb(pb *awsCredentialsPb) (*AwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsCredentials{}
	st.AccessKeyId = pb.AccessKeyId
	st.AccessPoint = pb.AccessPoint
	st.SecretAccessKey = pb.SecretAccessKey
	st.SessionToken = pb.SessionToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsCredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsCredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func awsIamRoleToPb(st *AwsIamRole) (*awsIamRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRolePb{}
	pb.ExternalId = st.ExternalId
	pb.RoleArn = st.RoleArn
	pb.UnityCatalogIamArn = st.UnityCatalogIamArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type awsIamRolePb struct {
	ExternalId         string `json:"external_id,omitempty"`
	RoleArn            string `json:"role_arn,omitempty"`
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsIamRoleFromPb(pb *awsIamRolePb) (*AwsIamRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRole{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn
	st.UnityCatalogIamArn = pb.UnityCatalogIamArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsIamRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsIamRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func awsIamRoleRequestToPb(st *AwsIamRoleRequest) (*awsIamRoleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRoleRequestPb{}
	pb.RoleArn = st.RoleArn

	return pb, nil
}

type awsIamRoleRequestPb struct {
	RoleArn string `json:"role_arn"`
}

func awsIamRoleRequestFromPb(pb *awsIamRoleRequestPb) (*AwsIamRoleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRoleRequest{}
	st.RoleArn = pb.RoleArn

	return st, nil
}

func awsIamRoleResponseToPb(st *AwsIamRoleResponse) (*awsIamRoleResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsIamRoleResponsePb{}
	pb.ExternalId = st.ExternalId
	pb.RoleArn = st.RoleArn
	pb.UnityCatalogIamArn = st.UnityCatalogIamArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type awsIamRoleResponsePb struct {
	ExternalId         string `json:"external_id,omitempty"`
	RoleArn            string `json:"role_arn"`
	UnityCatalogIamArn string `json:"unity_catalog_iam_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsIamRoleResponseFromPb(pb *awsIamRoleResponsePb) (*AwsIamRoleResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsIamRoleResponse{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn
	st.UnityCatalogIamArn = pb.UnityCatalogIamArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsIamRoleResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsIamRoleResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func awsSqsQueueToPb(st *AwsSqsQueue) (*awsSqsQueuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsSqsQueuePb{}
	pb.ManagedResourceId = st.ManagedResourceId
	pb.QueueUrl = st.QueueUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type awsSqsQueuePb struct {
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	QueueUrl          string `json:"queue_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsSqsQueueFromPb(pb *awsSqsQueuePb) (*AwsSqsQueue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsSqsQueue{}
	st.ManagedResourceId = pb.ManagedResourceId
	st.QueueUrl = pb.QueueUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsSqsQueuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsSqsQueuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureActiveDirectoryTokenToPb(st *AzureActiveDirectoryToken) (*azureActiveDirectoryTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureActiveDirectoryTokenPb{}
	pb.AadToken = st.AadToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureActiveDirectoryTokenPb struct {
	AadToken string `json:"aad_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureActiveDirectoryTokenFromPb(pb *azureActiveDirectoryTokenPb) (*AzureActiveDirectoryToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureActiveDirectoryToken{}
	st.AadToken = pb.AadToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureActiveDirectoryTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureActiveDirectoryTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureManagedIdentityToPb(st *AzureManagedIdentity) (*azureManagedIdentityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityPb{}
	pb.AccessConnectorId = st.AccessConnectorId
	pb.CredentialId = st.CredentialId
	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureManagedIdentityPb struct {
	AccessConnectorId string `json:"access_connector_id"`
	CredentialId      string `json:"credential_id,omitempty"`
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureManagedIdentityFromPb(pb *azureManagedIdentityPb) (*AzureManagedIdentity, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentity{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.CredentialId = pb.CredentialId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureManagedIdentityPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureManagedIdentityPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureManagedIdentityRequestToPb(st *AzureManagedIdentityRequest) (*azureManagedIdentityRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityRequestPb{}
	pb.AccessConnectorId = st.AccessConnectorId
	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureManagedIdentityRequestPb struct {
	AccessConnectorId string `json:"access_connector_id"`
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureManagedIdentityRequestFromPb(pb *azureManagedIdentityRequestPb) (*AzureManagedIdentityRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentityRequest{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureManagedIdentityRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureManagedIdentityRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureManagedIdentityResponseToPb(st *AzureManagedIdentityResponse) (*azureManagedIdentityResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureManagedIdentityResponsePb{}
	pb.AccessConnectorId = st.AccessConnectorId
	pb.CredentialId = st.CredentialId
	pb.ManagedIdentityId = st.ManagedIdentityId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureManagedIdentityResponsePb struct {
	AccessConnectorId string `json:"access_connector_id"`
	CredentialId      string `json:"credential_id,omitempty"`
	ManagedIdentityId string `json:"managed_identity_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureManagedIdentityResponseFromPb(pb *azureManagedIdentityResponsePb) (*AzureManagedIdentityResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureManagedIdentityResponse{}
	st.AccessConnectorId = pb.AccessConnectorId
	st.CredentialId = pb.CredentialId
	st.ManagedIdentityId = pb.ManagedIdentityId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureManagedIdentityResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureManagedIdentityResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureQueueStorageToPb(st *AzureQueueStorage) (*azureQueueStoragePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureQueueStoragePb{}
	pb.ManagedResourceId = st.ManagedResourceId
	pb.QueueUrl = st.QueueUrl
	pb.ResourceGroup = st.ResourceGroup
	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureQueueStoragePb struct {
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	QueueUrl          string `json:"queue_url,omitempty"`
	ResourceGroup     string `json:"resource_group,omitempty"`
	SubscriptionId    string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureQueueStorageFromPb(pb *azureQueueStoragePb) (*AzureQueueStorage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureQueueStorage{}
	st.ManagedResourceId = pb.ManagedResourceId
	st.QueueUrl = pb.QueueUrl
	st.ResourceGroup = pb.ResourceGroup
	st.SubscriptionId = pb.SubscriptionId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureQueueStoragePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureQueueStoragePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureServicePrincipalToPb(st *AzureServicePrincipal) (*azureServicePrincipalPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureServicePrincipalPb{}
	pb.ApplicationId = st.ApplicationId
	pb.ClientSecret = st.ClientSecret
	pb.DirectoryId = st.DirectoryId

	return pb, nil
}

type azureServicePrincipalPb struct {
	ApplicationId string `json:"application_id"`
	ClientSecret  string `json:"client_secret"`
	DirectoryId   string `json:"directory_id"`
}

func azureServicePrincipalFromPb(pb *azureServicePrincipalPb) (*AzureServicePrincipal, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureServicePrincipal{}
	st.ApplicationId = pb.ApplicationId
	st.ClientSecret = pb.ClientSecret
	st.DirectoryId = pb.DirectoryId

	return st, nil
}

func azureUserDelegationSasToPb(st *AzureUserDelegationSas) (*azureUserDelegationSasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureUserDelegationSasPb{}
	pb.SasToken = st.SasToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureUserDelegationSasPb struct {
	SasToken string `json:"sas_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureUserDelegationSasFromPb(pb *azureUserDelegationSasPb) (*AzureUserDelegationSas, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureUserDelegationSas{}
	st.SasToken = pb.SasToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureUserDelegationSasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureUserDelegationSasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cancelRefreshRequestToPb(st *CancelRefreshRequest) (*cancelRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRefreshRequestPb{}
	pb.RefreshId = st.RefreshId
	pb.TableName = st.TableName

	return pb, nil
}

type cancelRefreshRequestPb struct {
	RefreshId string `json:"-" url:"-"`
	TableName string `json:"-" url:"-"`
}

func cancelRefreshRequestFromPb(pb *cancelRefreshRequestPb) (*CancelRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRefreshRequest{}
	st.RefreshId = pb.RefreshId
	st.TableName = pb.TableName

	return st, nil
}

func cancelRefreshResponseToPb(st *CancelRefreshResponse) (*cancelRefreshResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRefreshResponsePb{}

	return pb, nil
}

type cancelRefreshResponsePb struct {
}

func cancelRefreshResponseFromPb(pb *cancelRefreshResponsePb) (*CancelRefreshResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRefreshResponse{}

	return st, nil
}

func catalogInfoToPb(st *CatalogInfo) (*catalogInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &catalogInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogType = st.CatalogType
	pb.Comment = st.Comment
	pb.ConnectionName = st.ConnectionName
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.EffectivePredictiveOptimizationFlag = st.EffectivePredictiveOptimizationFlag
	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization
	pb.FullName = st.FullName
	pb.IsolationMode = st.IsolationMode
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Owner = st.Owner
	pb.Properties = st.Properties
	pb.ProviderName = st.ProviderName
	pb.ProvisioningInfo = st.ProvisioningInfo
	pb.SecurableType = st.SecurableType
	pb.ShareName = st.ShareName
	pb.StorageLocation = st.StorageLocation
	pb.StorageRoot = st.StorageRoot
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type catalogInfoPb struct {
	BrowseOnly                          bool                                 `json:"browse_only,omitempty"`
	CatalogType                         CatalogType                          `json:"catalog_type,omitempty"`
	Comment                             string                               `json:"comment,omitempty"`
	ConnectionName                      string                               `json:"connection_name,omitempty"`
	CreatedAt                           int64                                `json:"created_at,omitempty"`
	CreatedBy                           string                               `json:"created_by,omitempty"`
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	EnablePredictiveOptimization        EnablePredictiveOptimization         `json:"enable_predictive_optimization,omitempty"`
	FullName                            string                               `json:"full_name,omitempty"`
	IsolationMode                       CatalogIsolationMode                 `json:"isolation_mode,omitempty"`
	MetastoreId                         string                               `json:"metastore_id,omitempty"`
	Name                                string                               `json:"name,omitempty"`
	Options                             map[string]string                    `json:"options,omitempty"`
	Owner                               string                               `json:"owner,omitempty"`
	Properties                          map[string]string                    `json:"properties,omitempty"`
	ProviderName                        string                               `json:"provider_name,omitempty"`
	ProvisioningInfo                    *ProvisioningInfo                    `json:"provisioning_info,omitempty"`
	SecurableType                       SecurableType                        `json:"securable_type,omitempty"`
	ShareName                           string                               `json:"share_name,omitempty"`
	StorageLocation                     string                               `json:"storage_location,omitempty"`
	StorageRoot                         string                               `json:"storage_root,omitempty"`
	UpdatedAt                           int64                                `json:"updated_at,omitempty"`
	UpdatedBy                           string                               `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func catalogInfoFromPb(pb *catalogInfoPb) (*CatalogInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CatalogInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogType = pb.CatalogType
	st.Comment = pb.Comment
	st.ConnectionName = pb.ConnectionName
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.EffectivePredictiveOptimizationFlag = pb.EffectivePredictiveOptimizationFlag
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.FullName = pb.FullName
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	st.ProviderName = pb.ProviderName
	st.ProvisioningInfo = pb.ProvisioningInfo
	st.SecurableType = pb.SecurableType
	st.ShareName = pb.ShareName
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *catalogInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st catalogInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cloudflareApiTokenToPb(st *CloudflareApiToken) (*cloudflareApiTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudflareApiTokenPb{}
	pb.AccessKeyId = st.AccessKeyId
	pb.AccountId = st.AccountId
	pb.SecretAccessKey = st.SecretAccessKey

	return pb, nil
}

type cloudflareApiTokenPb struct {
	AccessKeyId     string `json:"access_key_id"`
	AccountId       string `json:"account_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

func cloudflareApiTokenFromPb(pb *cloudflareApiTokenPb) (*CloudflareApiToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudflareApiToken{}
	st.AccessKeyId = pb.AccessKeyId
	st.AccountId = pb.AccountId
	st.SecretAccessKey = pb.SecretAccessKey

	return st, nil
}

func columnInfoToPb(st *ColumnInfo) (*columnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnInfoPb{}
	pb.Comment = st.Comment
	pb.Mask = st.Mask
	pb.Name = st.Name
	pb.Nullable = st.Nullable
	pb.PartitionIndex = st.PartitionIndex
	pb.Position = st.Position
	pb.TypeIntervalType = st.TypeIntervalType
	pb.TypeJson = st.TypeJson
	pb.TypeName = st.TypeName
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type columnInfoPb struct {
	Comment          string         `json:"comment,omitempty"`
	Mask             *ColumnMask    `json:"mask,omitempty"`
	Name             string         `json:"name,omitempty"`
	Nullable         bool           `json:"nullable,omitempty"`
	PartitionIndex   int            `json:"partition_index,omitempty"`
	Position         int            `json:"position,omitempty"`
	TypeIntervalType string         `json:"type_interval_type,omitempty"`
	TypeJson         string         `json:"type_json,omitempty"`
	TypeName         ColumnTypeName `json:"type_name,omitempty"`
	TypePrecision    int            `json:"type_precision,omitempty"`
	TypeScale        int            `json:"type_scale,omitempty"`
	TypeText         string         `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func columnInfoFromPb(pb *columnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Comment = pb.Comment
	st.Mask = pb.Mask
	st.Name = pb.Name
	st.Nullable = pb.Nullable
	st.PartitionIndex = pb.PartitionIndex
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	st.TypeName = pb.TypeName
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *columnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st columnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func columnMaskToPb(st *ColumnMask) (*columnMaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnMaskPb{}
	pb.FunctionName = st.FunctionName
	pb.UsingColumnNames = st.UsingColumnNames

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type columnMaskPb struct {
	FunctionName     string   `json:"function_name,omitempty"`
	UsingColumnNames []string `json:"using_column_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func columnMaskFromPb(pb *columnMaskPb) (*ColumnMask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnMask{}
	st.FunctionName = pb.FunctionName
	st.UsingColumnNames = pb.UsingColumnNames

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *columnMaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st columnMaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func connectionInfoToPb(st *ConnectionInfo) (*connectionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &connectionInfoPb{}
	pb.Comment = st.Comment
	pb.ConnectionId = st.ConnectionId
	pb.ConnectionType = st.ConnectionType
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.CredentialType = st.CredentialType
	pb.FullName = st.FullName
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Owner = st.Owner
	pb.Properties = st.Properties
	pb.ProvisioningInfo = st.ProvisioningInfo
	pb.ReadOnly = st.ReadOnly
	pb.SecurableType = st.SecurableType
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type connectionInfoPb struct {
	Comment          string            `json:"comment,omitempty"`
	ConnectionId     string            `json:"connection_id,omitempty"`
	ConnectionType   ConnectionType    `json:"connection_type,omitempty"`
	CreatedAt        int64             `json:"created_at,omitempty"`
	CreatedBy        string            `json:"created_by,omitempty"`
	CredentialType   CredentialType    `json:"credential_type,omitempty"`
	FullName         string            `json:"full_name,omitempty"`
	MetastoreId      string            `json:"metastore_id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Options          map[string]string `json:"options,omitempty"`
	Owner            string            `json:"owner,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
	ProvisioningInfo *ProvisioningInfo `json:"provisioning_info,omitempty"`
	ReadOnly         bool              `json:"read_only,omitempty"`
	SecurableType    SecurableType     `json:"securable_type,omitempty"`
	UpdatedAt        int64             `json:"updated_at,omitempty"`
	UpdatedBy        string            `json:"updated_by,omitempty"`
	Url              string            `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func connectionInfoFromPb(pb *connectionInfoPb) (*ConnectionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConnectionInfo{}
	st.Comment = pb.Comment
	st.ConnectionId = pb.ConnectionId
	st.ConnectionType = pb.ConnectionType
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CredentialType = pb.CredentialType
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	st.ProvisioningInfo = pb.ProvisioningInfo
	st.ReadOnly = pb.ReadOnly
	st.SecurableType = pb.SecurableType
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *connectionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st connectionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func continuousUpdateStatusToPb(st *ContinuousUpdateStatus) (*continuousUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &continuousUpdateStatusPb{}
	pb.InitialPipelineSyncProgress = st.InitialPipelineSyncProgress
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type continuousUpdateStatusPb struct {
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
	LastProcessedCommitVersion  int64             `json:"last_processed_commit_version,omitempty"`
	Timestamp                   string            `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func continuousUpdateStatusFromPb(pb *continuousUpdateStatusPb) (*ContinuousUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContinuousUpdateStatus{}
	st.InitialPipelineSyncProgress = pb.InitialPipelineSyncProgress
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *continuousUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st continuousUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCatalogToPb(st *CreateCatalog) (*createCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCatalogPb{}
	pb.Comment = st.Comment
	pb.ConnectionName = st.ConnectionName
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Properties = st.Properties
	pb.ProviderName = st.ProviderName
	pb.ShareName = st.ShareName
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createCatalogPb struct {
	Comment        string            `json:"comment,omitempty"`
	ConnectionName string            `json:"connection_name,omitempty"`
	Name           string            `json:"name"`
	Options        map[string]string `json:"options,omitempty"`
	Properties     map[string]string `json:"properties,omitempty"`
	ProviderName   string            `json:"provider_name,omitempty"`
	ShareName      string            `json:"share_name,omitempty"`
	StorageRoot    string            `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCatalogFromPb(pb *createCatalogPb) (*CreateCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCatalog{}
	st.Comment = pb.Comment
	st.ConnectionName = pb.ConnectionName
	st.Name = pb.Name
	st.Options = pb.Options
	st.Properties = pb.Properties
	st.ProviderName = pb.ProviderName
	st.ShareName = pb.ShareName
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createConnectionToPb(st *CreateConnection) (*createConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createConnectionPb{}
	pb.Comment = st.Comment
	pb.ConnectionType = st.ConnectionType
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Properties = st.Properties
	pb.ReadOnly = st.ReadOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createConnectionPb struct {
	Comment        string            `json:"comment,omitempty"`
	ConnectionType ConnectionType    `json:"connection_type"`
	Name           string            `json:"name"`
	Options        map[string]string `json:"options"`
	Properties     map[string]string `json:"properties,omitempty"`
	ReadOnly       bool              `json:"read_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createConnectionFromPb(pb *createConnectionPb) (*CreateConnection, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateConnection{}
	st.Comment = pb.Comment
	st.ConnectionType = pb.ConnectionType
	st.Name = pb.Name
	st.Options = pb.Options
	st.Properties = pb.Properties
	st.ReadOnly = pb.ReadOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createConnectionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createConnectionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCredentialRequestToPb(st *CreateCredentialRequest) (*createCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialRequestPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.Comment = st.Comment
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.Name = st.Name
	pb.Purpose = st.Purpose
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createCredentialRequestPb struct {
	AwsIamRole                  *AwsIamRole                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentity        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal       `json:"azure_service_principal,omitempty"`
	Comment                     string                       `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	Name                        string                       `json:"name"`
	Purpose                     CredentialPurpose            `json:"purpose,omitempty"`
	ReadOnly                    bool                         `json:"read_only,omitempty"`
	SkipValidation              bool                         `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCredentialRequestFromPb(pb *createCredentialRequestPb) (*CreateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialRequest{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.Comment = pb.Comment
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.Name = pb.Name
	st.Purpose = pb.Purpose
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createExternalLocationToPb(st *CreateExternalLocation) (*createExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExternalLocationPb{}
	pb.Comment = st.Comment
	pb.CredentialName = st.CredentialName
	pb.EnableFileEvents = st.EnableFileEvents
	pb.EncryptionDetails = st.EncryptionDetails
	pb.Fallback = st.Fallback
	pb.FileEventQueue = st.FileEventQueue
	pb.Name = st.Name
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createExternalLocationPb struct {
	Comment           string             `json:"comment,omitempty"`
	CredentialName    string             `json:"credential_name"`
	EnableFileEvents  bool               `json:"enable_file_events,omitempty"`
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	Fallback          bool               `json:"fallback,omitempty"`
	FileEventQueue    *FileEventQueue    `json:"file_event_queue,omitempty"`
	Name              string             `json:"name"`
	ReadOnly          bool               `json:"read_only,omitempty"`
	SkipValidation    bool               `json:"skip_validation,omitempty"`
	Url               string             `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExternalLocationFromPb(pb *createExternalLocationPb) (*CreateExternalLocation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExternalLocation{}
	st.Comment = pb.Comment
	st.CredentialName = pb.CredentialName
	st.EnableFileEvents = pb.EnableFileEvents
	st.EncryptionDetails = pb.EncryptionDetails
	st.Fallback = pb.Fallback
	st.FileEventQueue = pb.FileEventQueue
	st.Name = pb.Name
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExternalLocationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExternalLocationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createFunctionToPb(st *CreateFunction) (*createFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFunctionPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.DataType = st.DataType
	pb.ExternalLanguage = st.ExternalLanguage
	pb.ExternalName = st.ExternalName
	pb.FullDataType = st.FullDataType
	pb.InputParams = st.InputParams
	pb.IsDeterministic = st.IsDeterministic
	pb.IsNullCall = st.IsNullCall
	pb.Name = st.Name
	pb.ParameterStyle = st.ParameterStyle
	pb.Properties = st.Properties
	pb.ReturnParams = st.ReturnParams
	pb.RoutineBody = st.RoutineBody
	pb.RoutineDefinition = st.RoutineDefinition
	pb.RoutineDependencies = st.RoutineDependencies
	pb.SchemaName = st.SchemaName
	pb.SecurityType = st.SecurityType
	pb.SpecificName = st.SpecificName
	pb.SqlDataAccess = st.SqlDataAccess
	pb.SqlPath = st.SqlPath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createFunctionPb struct {
	CatalogName         string                       `json:"catalog_name"`
	Comment             string                       `json:"comment,omitempty"`
	DataType            ColumnTypeName               `json:"data_type"`
	ExternalLanguage    string                       `json:"external_language,omitempty"`
	ExternalName        string                       `json:"external_name,omitempty"`
	FullDataType        string                       `json:"full_data_type"`
	InputParams         FunctionParameterInfos       `json:"input_params"`
	IsDeterministic     bool                         `json:"is_deterministic"`
	IsNullCall          bool                         `json:"is_null_call"`
	Name                string                       `json:"name"`
	ParameterStyle      CreateFunctionParameterStyle `json:"parameter_style"`
	Properties          string                       `json:"properties,omitempty"`
	ReturnParams        *FunctionParameterInfos      `json:"return_params,omitempty"`
	RoutineBody         CreateFunctionRoutineBody    `json:"routine_body"`
	RoutineDefinition   string                       `json:"routine_definition"`
	RoutineDependencies *DependencyList              `json:"routine_dependencies,omitempty"`
	SchemaName          string                       `json:"schema_name"`
	SecurityType        CreateFunctionSecurityType   `json:"security_type"`
	SpecificName        string                       `json:"specific_name"`
	SqlDataAccess       CreateFunctionSqlDataAccess  `json:"sql_data_access"`
	SqlPath             string                       `json:"sql_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFunctionFromPb(pb *createFunctionPb) (*CreateFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFunction{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.DataType = pb.DataType
	st.ExternalLanguage = pb.ExternalLanguage
	st.ExternalName = pb.ExternalName
	st.FullDataType = pb.FullDataType
	st.InputParams = pb.InputParams
	st.IsDeterministic = pb.IsDeterministic
	st.IsNullCall = pb.IsNullCall
	st.Name = pb.Name
	st.ParameterStyle = pb.ParameterStyle
	st.Properties = pb.Properties
	st.ReturnParams = pb.ReturnParams
	st.RoutineBody = pb.RoutineBody
	st.RoutineDefinition = pb.RoutineDefinition
	st.RoutineDependencies = pb.RoutineDependencies
	st.SchemaName = pb.SchemaName
	st.SecurityType = pb.SecurityType
	st.SpecificName = pb.SpecificName
	st.SqlDataAccess = pb.SqlDataAccess
	st.SqlPath = pb.SqlPath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createFunctionRequestToPb(st *CreateFunctionRequest) (*createFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createFunctionRequestPb{}
	pb.FunctionInfo = st.FunctionInfo

	return pb, nil
}

type createFunctionRequestPb struct {
	FunctionInfo CreateFunction `json:"function_info"`
}

func createFunctionRequestFromPb(pb *createFunctionRequestPb) (*CreateFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFunctionRequest{}
	st.FunctionInfo = pb.FunctionInfo

	return st, nil
}

func createMetastoreToPb(st *CreateMetastore) (*createMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMetastorePb{}
	pb.Name = st.Name
	pb.Region = st.Region
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createMetastorePb struct {
	Name        string `json:"name"`
	Region      string `json:"region,omitempty"`
	StorageRoot string `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createMetastoreFromPb(pb *createMetastorePb) (*CreateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMetastore{}
	st.Name = pb.Name
	st.Region = pb.Region
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createMetastorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createMetastorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createMetastoreAssignmentToPb(st *CreateMetastoreAssignment) (*createMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type createMetastoreAssignmentPb struct {
	DefaultCatalogName string `json:"default_catalog_name"`
	MetastoreId        string `json:"metastore_id"`
	WorkspaceId        int64  `json:"-" url:"-"`
}

func createMetastoreAssignmentFromPb(pb *createMetastoreAssignmentPb) (*CreateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMetastoreAssignment{}
	st.DefaultCatalogName = pb.DefaultCatalogName
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func createMonitorToPb(st *CreateMonitor) (*createMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createMonitorPb{}
	pb.AssetsDir = st.AssetsDir
	pb.BaselineTableName = st.BaselineTableName
	pb.CustomMetrics = st.CustomMetrics
	pb.DataClassificationConfig = st.DataClassificationConfig
	pb.InferenceLog = st.InferenceLog
	pb.Notifications = st.Notifications
	pb.OutputSchemaName = st.OutputSchemaName
	pb.Schedule = st.Schedule
	pb.SkipBuiltinDashboard = st.SkipBuiltinDashboard
	pb.SlicingExprs = st.SlicingExprs
	pb.Snapshot = st.Snapshot
	pb.TableName = st.TableName
	pb.TimeSeries = st.TimeSeries
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createMonitorPb struct {
	AssetsDir                string                           `json:"assets_dir"`
	BaselineTableName        string                           `json:"baseline_table_name,omitempty"`
	CustomMetrics            []MonitorMetric                  `json:"custom_metrics,omitempty"`
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	InferenceLog             *MonitorInferenceLog             `json:"inference_log,omitempty"`
	Notifications            *MonitorNotifications            `json:"notifications,omitempty"`
	OutputSchemaName         string                           `json:"output_schema_name"`
	Schedule                 *MonitorCronSchedule             `json:"schedule,omitempty"`
	SkipBuiltinDashboard     bool                             `json:"skip_builtin_dashboard,omitempty"`
	SlicingExprs             []string                         `json:"slicing_exprs,omitempty"`
	Snapshot                 *MonitorSnapshot                 `json:"snapshot,omitempty"`
	TableName                string                           `json:"-" url:"-"`
	TimeSeries               *MonitorTimeSeries               `json:"time_series,omitempty"`
	WarehouseId              string                           `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createMonitorFromPb(pb *createMonitorPb) (*CreateMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateMonitor{}
	st.AssetsDir = pb.AssetsDir
	st.BaselineTableName = pb.BaselineTableName
	st.CustomMetrics = pb.CustomMetrics
	st.DataClassificationConfig = pb.DataClassificationConfig
	st.InferenceLog = pb.InferenceLog
	st.Notifications = pb.Notifications
	st.OutputSchemaName = pb.OutputSchemaName
	st.Schedule = pb.Schedule
	st.SkipBuiltinDashboard = pb.SkipBuiltinDashboard
	st.SlicingExprs = pb.SlicingExprs
	st.Snapshot = pb.Snapshot
	st.TableName = pb.TableName
	st.TimeSeries = pb.TimeSeries
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createMonitorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createMonitorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createOnlineTableRequestToPb(st *CreateOnlineTableRequest) (*createOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOnlineTableRequestPb{}
	pb.Table = st.Table

	return pb, nil
}

type createOnlineTableRequestPb struct {
	Table OnlineTable `json:"table"`
}

func createOnlineTableRequestFromPb(pb *createOnlineTableRequestPb) (*CreateOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOnlineTableRequest{}
	st.Table = pb.Table

	return st, nil
}

func createRegisteredModelRequestToPb(st *CreateRegisteredModelRequest) (*createRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRegisteredModelRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.SchemaName = st.SchemaName
	pb.StorageLocation = st.StorageLocation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createRegisteredModelRequestPb struct {
	CatalogName     string `json:"catalog_name"`
	Comment         string `json:"comment,omitempty"`
	Name            string `json:"name"`
	SchemaName      string `json:"schema_name"`
	StorageLocation string `json:"storage_location,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRegisteredModelRequestFromPb(pb *createRegisteredModelRequestPb) (*CreateRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRegisteredModelRequest{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}

	return pb, nil
}

type createResponsePb struct {
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}

	return st, nil
}

func createSchemaToPb(st *CreateSchema) (*createSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createSchemaPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Properties = st.Properties
	pb.StorageRoot = st.StorageRoot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createSchemaPb struct {
	CatalogName string            `json:"catalog_name"`
	Comment     string            `json:"comment,omitempty"`
	Name        string            `json:"name"`
	Properties  map[string]string `json:"properties,omitempty"`
	StorageRoot string            `json:"storage_root,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createSchemaFromPb(pb *createSchemaPb) (*CreateSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateSchema{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Properties = pb.Properties
	st.StorageRoot = pb.StorageRoot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createStorageCredentialToPb(st *CreateStorageCredential) (*createStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createStorageCredentialPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.CloudflareApiToken = st.CloudflareApiToken
	pb.Comment = st.Comment
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.Name = st.Name
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createStorageCredentialPb struct {
	AwsIamRole                  *AwsIamRoleRequest                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityRequest        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal              `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiToken                 `json:"cloudflare_api_token,omitempty"`
	Comment                     string                              `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	Name                        string                              `json:"name"`
	ReadOnly                    bool                                `json:"read_only,omitempty"`
	SkipValidation              bool                                `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createStorageCredentialFromPb(pb *createStorageCredentialPb) (*CreateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateStorageCredential{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.CloudflareApiToken = pb.CloudflareApiToken
	st.Comment = pb.Comment
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.Name = pb.Name
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createTableConstraintToPb(st *CreateTableConstraint) (*createTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTableConstraintPb{}
	pb.Constraint = st.Constraint
	pb.FullNameArg = st.FullNameArg

	return pb, nil
}

type createTableConstraintPb struct {
	Constraint  TableConstraint `json:"constraint"`
	FullNameArg string          `json:"full_name_arg"`
}

func createTableConstraintFromPb(pb *createTableConstraintPb) (*CreateTableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTableConstraint{}
	st.Constraint = pb.Constraint
	st.FullNameArg = pb.FullNameArg

	return st, nil
}

func createVolumeRequestContentToPb(st *CreateVolumeRequestContent) (*createVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVolumeRequestContentPb{}
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.SchemaName = st.SchemaName
	pb.StorageLocation = st.StorageLocation
	pb.VolumeType = st.VolumeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createVolumeRequestContentPb struct {
	CatalogName     string     `json:"catalog_name"`
	Comment         string     `json:"comment,omitempty"`
	Name            string     `json:"name"`
	SchemaName      string     `json:"schema_name"`
	StorageLocation string     `json:"storage_location,omitempty"`
	VolumeType      VolumeType `json:"volume_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createVolumeRequestContentFromPb(pb *createVolumeRequestContentPb) (*CreateVolumeRequestContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVolumeRequestContent{}
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	st.VolumeType = pb.VolumeType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createVolumeRequestContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createVolumeRequestContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func credentialInfoToPb(st *CredentialInfo) (*credentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialInfoPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.FullName = st.FullName
	pb.Id = st.Id
	pb.IsolationMode = st.IsolationMode
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.Purpose = st.Purpose
	pb.ReadOnly = st.ReadOnly
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.UsedForManagedStorage = st.UsedForManagedStorage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type credentialInfoPb struct {
	AwsIamRole                  *AwsIamRole                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentity        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal       `json:"azure_service_principal,omitempty"`
	Comment                     string                       `json:"comment,omitempty"`
	CreatedAt                   int64                        `json:"created_at,omitempty"`
	CreatedBy                   string                       `json:"created_by,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	FullName                    string                       `json:"full_name,omitempty"`
	Id                          string                       `json:"id,omitempty"`
	IsolationMode               IsolationMode                `json:"isolation_mode,omitempty"`
	MetastoreId                 string                       `json:"metastore_id,omitempty"`
	Name                        string                       `json:"name,omitempty"`
	Owner                       string                       `json:"owner,omitempty"`
	Purpose                     CredentialPurpose            `json:"purpose,omitempty"`
	ReadOnly                    bool                         `json:"read_only,omitempty"`
	UpdatedAt                   int64                        `json:"updated_at,omitempty"`
	UpdatedBy                   string                       `json:"updated_by,omitempty"`
	UsedForManagedStorage       bool                         `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialInfoFromPb(pb *credentialInfoPb) (*CredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialInfo{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.FullName = pb.FullName
	st.Id = pb.Id
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.Purpose = pb.Purpose
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UsedForManagedStorage = pb.UsedForManagedStorage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func credentialValidationResultToPb(st *CredentialValidationResult) (*credentialValidationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialValidationResultPb{}
	pb.Message = st.Message
	pb.Result = st.Result

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type credentialValidationResultPb struct {
	Message string                   `json:"message,omitempty"`
	Result  ValidateCredentialResult `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialValidationResultFromPb(pb *credentialValidationResultPb) (*CredentialValidationResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CredentialValidationResult{}
	st.Message = pb.Message
	st.Result = pb.Result

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialValidationResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialValidationResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databricksGcpServiceAccountToPb(st *DatabricksGcpServiceAccount) (*databricksGcpServiceAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountPb{}
	pb.CredentialId = st.CredentialId
	pb.Email = st.Email
	pb.PrivateKeyId = st.PrivateKeyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databricksGcpServiceAccountPb struct {
	CredentialId string `json:"credential_id,omitempty"`
	Email        string `json:"email,omitempty"`
	PrivateKeyId string `json:"private_key_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databricksGcpServiceAccountFromPb(pb *databricksGcpServiceAccountPb) (*DatabricksGcpServiceAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccount{}
	st.CredentialId = pb.CredentialId
	st.Email = pb.Email
	st.PrivateKeyId = pb.PrivateKeyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databricksGcpServiceAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databricksGcpServiceAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databricksGcpServiceAccountRequestToPb(st *DatabricksGcpServiceAccountRequest) (*databricksGcpServiceAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountRequestPb{}

	return pb, nil
}

type databricksGcpServiceAccountRequestPb struct {
}

func databricksGcpServiceAccountRequestFromPb(pb *databricksGcpServiceAccountRequestPb) (*DatabricksGcpServiceAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccountRequest{}

	return st, nil
}

func databricksGcpServiceAccountResponseToPb(st *DatabricksGcpServiceAccountResponse) (*databricksGcpServiceAccountResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksGcpServiceAccountResponsePb{}
	pb.CredentialId = st.CredentialId
	pb.Email = st.Email

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databricksGcpServiceAccountResponsePb struct {
	CredentialId string `json:"credential_id,omitempty"`
	Email        string `json:"email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databricksGcpServiceAccountResponseFromPb(pb *databricksGcpServiceAccountResponsePb) (*DatabricksGcpServiceAccountResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksGcpServiceAccountResponse{}
	st.CredentialId = pb.CredentialId
	st.Email = pb.Email

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databricksGcpServiceAccountResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databricksGcpServiceAccountResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAccountMetastoreAssignmentRequestToPb(st *DeleteAccountMetastoreAssignmentRequest) (*deleteAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountMetastoreAssignmentRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type deleteAccountMetastoreAssignmentRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
	WorkspaceId int64  `json:"-" url:"-"`
}

func deleteAccountMetastoreAssignmentRequestFromPb(pb *deleteAccountMetastoreAssignmentRequestPb) (*DeleteAccountMetastoreAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountMetastoreAssignmentRequest{}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func deleteAccountMetastoreRequestToPb(st *DeleteAccountMetastoreRequest) (*deleteAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountMetastoreRequestPb{}
	pb.Force = st.Force
	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteAccountMetastoreRequestPb struct {
	Force       bool   `json:"-" url:"force,omitempty"`
	MetastoreId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAccountMetastoreRequestFromPb(pb *deleteAccountMetastoreRequestPb) (*DeleteAccountMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountMetastoreRequest{}
	st.Force = pb.Force
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAccountMetastoreRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAccountMetastoreRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAccountStorageCredentialRequestToPb(st *DeleteAccountStorageCredentialRequest) (*deleteAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountStorageCredentialRequestPb{}
	pb.Force = st.Force
	pb.MetastoreId = st.MetastoreId
	pb.StorageCredentialName = st.StorageCredentialName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteAccountStorageCredentialRequestPb struct {
	Force                 bool   `json:"-" url:"force,omitempty"`
	MetastoreId           string `json:"-" url:"-"`
	StorageCredentialName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAccountStorageCredentialRequestFromPb(pb *deleteAccountStorageCredentialRequestPb) (*DeleteAccountStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountStorageCredentialRequest{}
	st.Force = pb.Force
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAccountStorageCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAccountStorageCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAliasRequestToPb(st *DeleteAliasRequest) (*deleteAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAliasRequestPb{}
	pb.Alias = st.Alias
	pb.FullName = st.FullName

	return pb, nil
}

type deleteAliasRequestPb struct {
	Alias    string `json:"-" url:"-"`
	FullName string `json:"-" url:"-"`
}

func deleteAliasRequestFromPb(pb *deleteAliasRequestPb) (*DeleteAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName

	return st, nil
}

func deleteAliasResponseToPb(st *DeleteAliasResponse) (*deleteAliasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAliasResponsePb{}

	return pb, nil
}

type deleteAliasResponsePb struct {
}

func deleteAliasResponseFromPb(pb *deleteAliasResponsePb) (*DeleteAliasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAliasResponse{}

	return st, nil
}

func deleteCatalogRequestToPb(st *DeleteCatalogRequest) (*deleteCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCatalogRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteCatalogRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteCatalogRequestFromPb(pb *deleteCatalogRequestPb) (*DeleteCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCatalogRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteCatalogRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteCatalogRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteConnectionRequestToPb(st *DeleteConnectionRequest) (*deleteConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteConnectionRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteConnectionRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteConnectionRequestFromPb(pb *deleteConnectionRequestPb) (*DeleteConnectionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteConnectionRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteCredentialRequestToPb(st *DeleteCredentialRequest) (*deleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialRequestPb{}
	pb.Force = st.Force
	pb.NameArg = st.NameArg

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteCredentialRequestPb struct {
	Force   bool   `json:"-" url:"force,omitempty"`
	NameArg string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteCredentialRequestFromPb(pb *deleteCredentialRequestPb) (*DeleteCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialRequest{}
	st.Force = pb.Force
	st.NameArg = pb.NameArg

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteCredentialResponseToPb(st *DeleteCredentialResponse) (*deleteCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialResponsePb{}

	return pb, nil
}

type deleteCredentialResponsePb struct {
}

func deleteCredentialResponseFromPb(pb *deleteCredentialResponsePb) (*DeleteCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialResponse{}

	return st, nil
}

func deleteExternalLocationRequestToPb(st *DeleteExternalLocationRequest) (*deleteExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExternalLocationRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteExternalLocationRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteExternalLocationRequestFromPb(pb *deleteExternalLocationRequestPb) (*DeleteExternalLocationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExternalLocationRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteExternalLocationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteExternalLocationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteFunctionRequestToPb(st *DeleteFunctionRequest) (*deleteFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFunctionRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteFunctionRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteFunctionRequestFromPb(pb *deleteFunctionRequestPb) (*DeleteFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFunctionRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteMetastoreRequestToPb(st *DeleteMetastoreRequest) (*deleteMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteMetastoreRequestPb{}
	pb.Force = st.Force
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteMetastoreRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Id    string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteMetastoreRequestFromPb(pb *deleteMetastoreRequestPb) (*DeleteMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteMetastoreRequest{}
	st.Force = pb.Force
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteMetastoreRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteMetastoreRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*deleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionRequestPb{}
	pb.FullName = st.FullName
	pb.Version = st.Version

	return pb, nil
}

type deleteModelVersionRequestPb struct {
	FullName string `json:"-" url:"-"`
	Version  int    `json:"-" url:"-"`
}

func deleteModelVersionRequestFromPb(pb *deleteModelVersionRequestPb) (*DeleteModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionRequest{}
	st.FullName = pb.FullName
	st.Version = pb.Version

	return st, nil
}

func deleteOnlineTableRequestToPb(st *DeleteOnlineTableRequest) (*deleteOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteOnlineTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteOnlineTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteOnlineTableRequestFromPb(pb *deleteOnlineTableRequestPb) (*DeleteOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteOnlineTableRequest{}
	st.Name = pb.Name

	return st, nil
}

func deleteQualityMonitorRequestToPb(st *DeleteQualityMonitorRequest) (*deleteQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQualityMonitorRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

type deleteQualityMonitorRequestPb struct {
	TableName string `json:"-" url:"-"`
}

func deleteQualityMonitorRequestFromPb(pb *deleteQualityMonitorRequestPb) (*DeleteQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQualityMonitorRequest{}
	st.TableName = pb.TableName

	return st, nil
}

func deleteRegisteredModelRequestToPb(st *DeleteRegisteredModelRequest) (*deleteRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRegisteredModelRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

type deleteRegisteredModelRequestPb struct {
	FullName string `json:"-" url:"-"`
}

func deleteRegisteredModelRequestFromPb(pb *deleteRegisteredModelRequestPb) (*DeleteRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRegisteredModelRequest{}
	st.FullName = pb.FullName

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

func deleteSchemaRequestToPb(st *DeleteSchemaRequest) (*deleteSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSchemaRequestPb{}
	pb.Force = st.Force
	pb.FullName = st.FullName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteSchemaRequestPb struct {
	Force    bool   `json:"-" url:"force,omitempty"`
	FullName string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteSchemaRequestFromPb(pb *deleteSchemaRequestPb) (*DeleteSchemaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSchemaRequest{}
	st.Force = pb.Force
	st.FullName = pb.FullName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteSchemaRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteSchemaRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteStorageCredentialRequestToPb(st *DeleteStorageCredentialRequest) (*deleteStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteStorageCredentialRequestPb{}
	pb.Force = st.Force
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteStorageCredentialRequestPb struct {
	Force bool   `json:"-" url:"force,omitempty"`
	Name  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteStorageCredentialRequestFromPb(pb *deleteStorageCredentialRequestPb) (*DeleteStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteStorageCredentialRequest{}
	st.Force = pb.Force
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteStorageCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteStorageCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteTableConstraintRequestToPb(st *DeleteTableConstraintRequest) (*deleteTableConstraintRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTableConstraintRequestPb{}
	pb.Cascade = st.Cascade
	pb.ConstraintName = st.ConstraintName
	pb.FullName = st.FullName

	return pb, nil
}

type deleteTableConstraintRequestPb struct {
	Cascade        bool   `json:"-" url:"cascade"`
	ConstraintName string `json:"-" url:"constraint_name"`
	FullName       string `json:"-" url:"-"`
}

func deleteTableConstraintRequestFromPb(pb *deleteTableConstraintRequestPb) (*DeleteTableConstraintRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTableConstraintRequest{}
	st.Cascade = pb.Cascade
	st.ConstraintName = pb.ConstraintName
	st.FullName = pb.FullName

	return st, nil
}

func deleteTableRequestToPb(st *DeleteTableRequest) (*deleteTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTableRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

type deleteTableRequestPb struct {
	FullName string `json:"-" url:"-"`
}

func deleteTableRequestFromPb(pb *deleteTableRequestPb) (*DeleteTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTableRequest{}
	st.FullName = pb.FullName

	return st, nil
}

func deleteVolumeRequestToPb(st *DeleteVolumeRequest) (*deleteVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVolumeRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteVolumeRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteVolumeRequestFromPb(pb *deleteVolumeRequestPb) (*DeleteVolumeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVolumeRequest{}
	st.Name = pb.Name

	return st, nil
}

func deltaRuntimePropertiesKvPairsToPb(st *DeltaRuntimePropertiesKvPairs) (*deltaRuntimePropertiesKvPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deltaRuntimePropertiesKvPairsPb{}
	pb.DeltaRuntimeProperties = st.DeltaRuntimeProperties

	return pb, nil
}

type deltaRuntimePropertiesKvPairsPb struct {
	DeltaRuntimeProperties map[string]string `json:"delta_runtime_properties"`
}

func deltaRuntimePropertiesKvPairsFromPb(pb *deltaRuntimePropertiesKvPairsPb) (*DeltaRuntimePropertiesKvPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeltaRuntimePropertiesKvPairs{}
	st.DeltaRuntimeProperties = pb.DeltaRuntimeProperties

	return st, nil
}

func dependencyToPb(st *Dependency) (*dependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dependencyPb{}
	pb.Function = st.Function
	pb.Table = st.Table

	return pb, nil
}

type dependencyPb struct {
	Function *FunctionDependency `json:"function,omitempty"`
	Table    *TableDependency    `json:"table,omitempty"`
}

func dependencyFromPb(pb *dependencyPb) (*Dependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dependency{}
	st.Function = pb.Function
	st.Table = pb.Table

	return st, nil
}

func dependencyListToPb(st *DependencyList) (*dependencyListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dependencyListPb{}
	pb.Dependencies = st.Dependencies

	return pb, nil
}

type dependencyListPb struct {
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

func dependencyListFromPb(pb *dependencyListPb) (*DependencyList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DependencyList{}
	st.Dependencies = pb.Dependencies

	return st, nil
}

func disableRequestToPb(st *DisableRequest) (*disableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.SchemaName = st.SchemaName

	return pb, nil
}

type disableRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
	SchemaName  string `json:"-" url:"-"`
}

func disableRequestFromPb(pb *disableRequestPb) (*DisableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableRequest{}
	st.MetastoreId = pb.MetastoreId
	st.SchemaName = pb.SchemaName

	return st, nil
}

func disableResponseToPb(st *DisableResponse) (*disableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableResponsePb{}

	return pb, nil
}

type disableResponsePb struct {
}

func disableResponseFromPb(pb *disableResponsePb) (*DisableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableResponse{}

	return st, nil
}

func effectivePermissionsListToPb(st *EffectivePermissionsList) (*effectivePermissionsListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePermissionsListPb{}
	pb.NextPageToken = st.NextPageToken
	pb.PrivilegeAssignments = st.PrivilegeAssignments

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type effectivePermissionsListPb struct {
	NextPageToken        string                         `json:"next_page_token,omitempty"`
	PrivilegeAssignments []EffectivePrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePermissionsListFromPb(pb *effectivePermissionsListPb) (*EffectivePermissionsList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePermissionsList{}
	st.NextPageToken = pb.NextPageToken
	st.PrivilegeAssignments = pb.PrivilegeAssignments

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePermissionsListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePermissionsListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func effectivePredictiveOptimizationFlagToPb(st *EffectivePredictiveOptimizationFlag) (*effectivePredictiveOptimizationFlagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePredictiveOptimizationFlagPb{}
	pb.InheritedFromName = st.InheritedFromName
	pb.InheritedFromType = st.InheritedFromType
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type effectivePredictiveOptimizationFlagPb struct {
	InheritedFromName string                                               `json:"inherited_from_name,omitempty"`
	InheritedFromType EffectivePredictiveOptimizationFlagInheritedFromType `json:"inherited_from_type,omitempty"`
	Value             EnablePredictiveOptimization                         `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePredictiveOptimizationFlagFromPb(pb *effectivePredictiveOptimizationFlagPb) (*EffectivePredictiveOptimizationFlag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePredictiveOptimizationFlag{}
	st.InheritedFromName = pb.InheritedFromName
	st.InheritedFromType = pb.InheritedFromType
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePredictiveOptimizationFlagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePredictiveOptimizationFlagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func effectivePrivilegeToPb(st *EffectivePrivilege) (*effectivePrivilegePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePrivilegePb{}
	pb.InheritedFromName = st.InheritedFromName
	pb.InheritedFromType = st.InheritedFromType
	pb.Privilege = st.Privilege

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type effectivePrivilegePb struct {
	InheritedFromName string        `json:"inherited_from_name,omitempty"`
	InheritedFromType SecurableType `json:"inherited_from_type,omitempty"`
	Privilege         Privilege     `json:"privilege,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePrivilegeFromPb(pb *effectivePrivilegePb) (*EffectivePrivilege, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePrivilege{}
	st.InheritedFromName = pb.InheritedFromName
	st.InheritedFromType = pb.InheritedFromType
	st.Privilege = pb.Privilege

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePrivilegePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePrivilegePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func effectivePrivilegeAssignmentToPb(st *EffectivePrivilegeAssignment) (*effectivePrivilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &effectivePrivilegeAssignmentPb{}
	pb.Principal = st.Principal
	pb.Privileges = st.Privileges

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type effectivePrivilegeAssignmentPb struct {
	Principal  string               `json:"principal,omitempty"`
	Privileges []EffectivePrivilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func effectivePrivilegeAssignmentFromPb(pb *effectivePrivilegeAssignmentPb) (*EffectivePrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EffectivePrivilegeAssignment{}
	st.Principal = pb.Principal
	st.Privileges = pb.Privileges

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *effectivePrivilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st effectivePrivilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enableRequestToPb(st *EnableRequest) (*enableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.MetastoreId = st.MetastoreId
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enableRequestPb struct {
	CatalogName string `json:"catalog_name,omitempty"`
	MetastoreId string `json:"-" url:"-"`
	SchemaName  string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableRequestFromPb(pb *enableRequestPb) (*EnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableRequest{}
	st.CatalogName = pb.CatalogName
	st.MetastoreId = pb.MetastoreId
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enableResponseToPb(st *EnableResponse) (*enableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableResponsePb{}

	return pb, nil
}

type enableResponsePb struct {
}

func enableResponseFromPb(pb *enableResponsePb) (*EnableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableResponse{}

	return st, nil
}

func encryptionDetailsToPb(st *EncryptionDetails) (*encryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &encryptionDetailsPb{}
	pb.SseEncryptionDetails = st.SseEncryptionDetails

	return pb, nil
}

type encryptionDetailsPb struct {
	SseEncryptionDetails *SseEncryptionDetails `json:"sse_encryption_details,omitempty"`
}

func encryptionDetailsFromPb(pb *encryptionDetailsPb) (*EncryptionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EncryptionDetails{}
	st.SseEncryptionDetails = pb.SseEncryptionDetails

	return st, nil
}

func existsRequestToPb(st *ExistsRequest) (*existsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &existsRequestPb{}
	pb.FullName = st.FullName

	return pb, nil
}

type existsRequestPb struct {
	FullName string `json:"-" url:"-"`
}

func existsRequestFromPb(pb *existsRequestPb) (*ExistsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExistsRequest{}
	st.FullName = pb.FullName

	return st, nil
}

func externalLocationInfoToPb(st *ExternalLocationInfo) (*externalLocationInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalLocationInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.CredentialId = st.CredentialId
	pb.CredentialName = st.CredentialName
	pb.EnableFileEvents = st.EnableFileEvents
	pb.EncryptionDetails = st.EncryptionDetails
	pb.Fallback = st.Fallback
	pb.FileEventQueue = st.FileEventQueue
	pb.IsolationMode = st.IsolationMode
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type externalLocationInfoPb struct {
	BrowseOnly        bool               `json:"browse_only,omitempty"`
	Comment           string             `json:"comment,omitempty"`
	CreatedAt         int64              `json:"created_at,omitempty"`
	CreatedBy         string             `json:"created_by,omitempty"`
	CredentialId      string             `json:"credential_id,omitempty"`
	CredentialName    string             `json:"credential_name,omitempty"`
	EnableFileEvents  bool               `json:"enable_file_events,omitempty"`
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	Fallback          bool               `json:"fallback,omitempty"`
	FileEventQueue    *FileEventQueue    `json:"file_event_queue,omitempty"`
	IsolationMode     IsolationMode      `json:"isolation_mode,omitempty"`
	MetastoreId       string             `json:"metastore_id,omitempty"`
	Name              string             `json:"name,omitempty"`
	Owner             string             `json:"owner,omitempty"`
	ReadOnly          bool               `json:"read_only,omitempty"`
	UpdatedAt         int64              `json:"updated_at,omitempty"`
	UpdatedBy         string             `json:"updated_by,omitempty"`
	Url               string             `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalLocationInfoFromPb(pb *externalLocationInfoPb) (*ExternalLocationInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLocationInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.CredentialId = pb.CredentialId
	st.CredentialName = pb.CredentialName
	st.EnableFileEvents = pb.EnableFileEvents
	st.EncryptionDetails = pb.EncryptionDetails
	st.Fallback = pb.Fallback
	st.FileEventQueue = pb.FileEventQueue
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalLocationInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalLocationInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func failedStatusToPb(st *FailedStatus) (*failedStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &failedStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type failedStatusPb struct {
	LastProcessedCommitVersion int64  `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func failedStatusFromPb(pb *failedStatusPb) (*FailedStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FailedStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *failedStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st failedStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fileEventQueueToPb(st *FileEventQueue) (*fileEventQueuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileEventQueuePb{}
	pb.ManagedAqs = st.ManagedAqs
	pb.ManagedPubsub = st.ManagedPubsub
	pb.ManagedSqs = st.ManagedSqs
	pb.ProvidedAqs = st.ProvidedAqs
	pb.ProvidedPubsub = st.ProvidedPubsub
	pb.ProvidedSqs = st.ProvidedSqs

	return pb, nil
}

type fileEventQueuePb struct {
	ManagedAqs     *AzureQueueStorage `json:"managed_aqs,omitempty"`
	ManagedPubsub  *GcpPubsub         `json:"managed_pubsub,omitempty"`
	ManagedSqs     *AwsSqsQueue       `json:"managed_sqs,omitempty"`
	ProvidedAqs    *AzureQueueStorage `json:"provided_aqs,omitempty"`
	ProvidedPubsub *GcpPubsub         `json:"provided_pubsub,omitempty"`
	ProvidedSqs    *AwsSqsQueue       `json:"provided_sqs,omitempty"`
}

func fileEventQueueFromPb(pb *fileEventQueuePb) (*FileEventQueue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileEventQueue{}
	st.ManagedAqs = pb.ManagedAqs
	st.ManagedPubsub = pb.ManagedPubsub
	st.ManagedSqs = pb.ManagedSqs
	st.ProvidedAqs = pb.ProvidedAqs
	st.ProvidedPubsub = pb.ProvidedPubsub
	st.ProvidedSqs = pb.ProvidedSqs

	return st, nil
}

func foreignKeyConstraintToPb(st *ForeignKeyConstraint) (*foreignKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &foreignKeyConstraintPb{}
	pb.ChildColumns = st.ChildColumns
	pb.Name = st.Name
	pb.ParentColumns = st.ParentColumns
	pb.ParentTable = st.ParentTable

	return pb, nil
}

type foreignKeyConstraintPb struct {
	ChildColumns  []string `json:"child_columns"`
	Name          string   `json:"name"`
	ParentColumns []string `json:"parent_columns"`
	ParentTable   string   `json:"parent_table"`
}

func foreignKeyConstraintFromPb(pb *foreignKeyConstraintPb) (*ForeignKeyConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForeignKeyConstraint{}
	st.ChildColumns = pb.ChildColumns
	st.Name = pb.Name
	st.ParentColumns = pb.ParentColumns
	st.ParentTable = pb.ParentTable

	return st, nil
}

func functionDependencyToPb(st *FunctionDependency) (*functionDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionDependencyPb{}
	pb.FunctionFullName = st.FunctionFullName

	return pb, nil
}

type functionDependencyPb struct {
	FunctionFullName string `json:"function_full_name"`
}

func functionDependencyFromPb(pb *functionDependencyPb) (*FunctionDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionDependency{}
	st.FunctionFullName = pb.FunctionFullName

	return st, nil
}

func functionInfoToPb(st *FunctionInfo) (*functionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DataType = st.DataType
	pb.ExternalLanguage = st.ExternalLanguage
	pb.ExternalName = st.ExternalName
	pb.FullDataType = st.FullDataType
	pb.FullName = st.FullName
	pb.FunctionId = st.FunctionId
	pb.InputParams = st.InputParams
	pb.IsDeterministic = st.IsDeterministic
	pb.IsNullCall = st.IsNullCall
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.ParameterStyle = st.ParameterStyle
	pb.Properties = st.Properties
	pb.ReturnParams = st.ReturnParams
	pb.RoutineBody = st.RoutineBody
	pb.RoutineDefinition = st.RoutineDefinition
	pb.RoutineDependencies = st.RoutineDependencies
	pb.SchemaName = st.SchemaName
	pb.SecurityType = st.SecurityType
	pb.SpecificName = st.SpecificName
	pb.SqlDataAccess = st.SqlDataAccess
	pb.SqlPath = st.SqlPath
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type functionInfoPb struct {
	BrowseOnly          bool                       `json:"browse_only,omitempty"`
	CatalogName         string                     `json:"catalog_name,omitempty"`
	Comment             string                     `json:"comment,omitempty"`
	CreatedAt           int64                      `json:"created_at,omitempty"`
	CreatedBy           string                     `json:"created_by,omitempty"`
	DataType            ColumnTypeName             `json:"data_type,omitempty"`
	ExternalLanguage    string                     `json:"external_language,omitempty"`
	ExternalName        string                     `json:"external_name,omitempty"`
	FullDataType        string                     `json:"full_data_type,omitempty"`
	FullName            string                     `json:"full_name,omitempty"`
	FunctionId          string                     `json:"function_id,omitempty"`
	InputParams         *FunctionParameterInfos    `json:"input_params,omitempty"`
	IsDeterministic     bool                       `json:"is_deterministic,omitempty"`
	IsNullCall          bool                       `json:"is_null_call,omitempty"`
	MetastoreId         string                     `json:"metastore_id,omitempty"`
	Name                string                     `json:"name,omitempty"`
	Owner               string                     `json:"owner,omitempty"`
	ParameterStyle      FunctionInfoParameterStyle `json:"parameter_style,omitempty"`
	Properties          string                     `json:"properties,omitempty"`
	ReturnParams        *FunctionParameterInfos    `json:"return_params,omitempty"`
	RoutineBody         FunctionInfoRoutineBody    `json:"routine_body,omitempty"`
	RoutineDefinition   string                     `json:"routine_definition,omitempty"`
	RoutineDependencies *DependencyList            `json:"routine_dependencies,omitempty"`
	SchemaName          string                     `json:"schema_name,omitempty"`
	SecurityType        FunctionInfoSecurityType   `json:"security_type,omitempty"`
	SpecificName        string                     `json:"specific_name,omitempty"`
	SqlDataAccess       FunctionInfoSqlDataAccess  `json:"sql_data_access,omitempty"`
	SqlPath             string                     `json:"sql_path,omitempty"`
	UpdatedAt           int64                      `json:"updated_at,omitempty"`
	UpdatedBy           string                     `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func functionInfoFromPb(pb *functionInfoPb) (*FunctionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataType = pb.DataType
	st.ExternalLanguage = pb.ExternalLanguage
	st.ExternalName = pb.ExternalName
	st.FullDataType = pb.FullDataType
	st.FullName = pb.FullName
	st.FunctionId = pb.FunctionId
	st.InputParams = pb.InputParams
	st.IsDeterministic = pb.IsDeterministic
	st.IsNullCall = pb.IsNullCall
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.ParameterStyle = pb.ParameterStyle
	st.Properties = pb.Properties
	st.ReturnParams = pb.ReturnParams
	st.RoutineBody = pb.RoutineBody
	st.RoutineDefinition = pb.RoutineDefinition
	st.RoutineDependencies = pb.RoutineDependencies
	st.SchemaName = pb.SchemaName
	st.SecurityType = pb.SecurityType
	st.SpecificName = pb.SpecificName
	st.SqlDataAccess = pb.SqlDataAccess
	st.SqlPath = pb.SqlPath
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *functionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st functionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func functionParameterInfoToPb(st *FunctionParameterInfo) (*functionParameterInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfoPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.ParameterDefault = st.ParameterDefault
	pb.ParameterMode = st.ParameterMode
	pb.ParameterType = st.ParameterType
	pb.Position = st.Position
	pb.TypeIntervalType = st.TypeIntervalType
	pb.TypeJson = st.TypeJson
	pb.TypeName = st.TypeName
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type functionParameterInfoPb struct {
	Comment          string                `json:"comment,omitempty"`
	Name             string                `json:"name"`
	ParameterDefault string                `json:"parameter_default,omitempty"`
	ParameterMode    FunctionParameterMode `json:"parameter_mode,omitempty"`
	ParameterType    FunctionParameterType `json:"parameter_type,omitempty"`
	Position         int                   `json:"position"`
	TypeIntervalType string                `json:"type_interval_type,omitempty"`
	TypeJson         string                `json:"type_json,omitempty"`
	TypeName         ColumnTypeName        `json:"type_name"`
	TypePrecision    int                   `json:"type_precision,omitempty"`
	TypeScale        int                   `json:"type_scale,omitempty"`
	TypeText         string                `json:"type_text"`

	ForceSendFields []string `json:"-" url:"-"`
}

func functionParameterInfoFromPb(pb *functionParameterInfoPb) (*FunctionParameterInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfo{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.ParameterDefault = pb.ParameterDefault
	st.ParameterMode = pb.ParameterMode
	st.ParameterType = pb.ParameterType
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeJson = pb.TypeJson
	st.TypeName = pb.TypeName
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *functionParameterInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st functionParameterInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func functionParameterInfosToPb(st *FunctionParameterInfos) (*functionParameterInfosPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &functionParameterInfosPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

type functionParameterInfosPb struct {
	Parameters []FunctionParameterInfo `json:"parameters,omitempty"`
}

func functionParameterInfosFromPb(pb *functionParameterInfosPb) (*FunctionParameterInfos, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FunctionParameterInfos{}
	st.Parameters = pb.Parameters

	return st, nil
}

func gcpOauthTokenToPb(st *GcpOauthToken) (*gcpOauthTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpOauthTokenPb{}
	pb.OauthToken = st.OauthToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gcpOauthTokenPb struct {
	OauthToken string `json:"oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpOauthTokenFromPb(pb *gcpOauthTokenPb) (*GcpOauthToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpOauthToken{}
	st.OauthToken = pb.OauthToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpOauthTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpOauthTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gcpPubsubToPb(st *GcpPubsub) (*gcpPubsubPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpPubsubPb{}
	pb.ManagedResourceId = st.ManagedResourceId
	pb.SubscriptionName = st.SubscriptionName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gcpPubsubPb struct {
	ManagedResourceId string `json:"managed_resource_id,omitempty"`
	SubscriptionName  string `json:"subscription_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpPubsubFromPb(pb *gcpPubsubPb) (*GcpPubsub, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpPubsub{}
	st.ManagedResourceId = pb.ManagedResourceId
	st.SubscriptionName = pb.SubscriptionName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpPubsubPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpPubsubPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func generateTemporaryServiceCredentialAzureOptionsToPb(st *GenerateTemporaryServiceCredentialAzureOptions) (*generateTemporaryServiceCredentialAzureOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialAzureOptionsPb{}
	pb.Resources = st.Resources

	return pb, nil
}

type generateTemporaryServiceCredentialAzureOptionsPb struct {
	Resources []string `json:"resources,omitempty"`
}

func generateTemporaryServiceCredentialAzureOptionsFromPb(pb *generateTemporaryServiceCredentialAzureOptionsPb) (*GenerateTemporaryServiceCredentialAzureOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialAzureOptions{}
	st.Resources = pb.Resources

	return st, nil
}

func generateTemporaryServiceCredentialGcpOptionsToPb(st *GenerateTemporaryServiceCredentialGcpOptions) (*generateTemporaryServiceCredentialGcpOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialGcpOptionsPb{}
	pb.Scopes = st.Scopes

	return pb, nil
}

type generateTemporaryServiceCredentialGcpOptionsPb struct {
	Scopes []string `json:"scopes,omitempty"`
}

func generateTemporaryServiceCredentialGcpOptionsFromPb(pb *generateTemporaryServiceCredentialGcpOptionsPb) (*GenerateTemporaryServiceCredentialGcpOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialGcpOptions{}
	st.Scopes = pb.Scopes

	return st, nil
}

func generateTemporaryServiceCredentialRequestToPb(st *GenerateTemporaryServiceCredentialRequest) (*generateTemporaryServiceCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryServiceCredentialRequestPb{}
	pb.AzureOptions = st.AzureOptions
	pb.CredentialName = st.CredentialName
	pb.GcpOptions = st.GcpOptions

	return pb, nil
}

type generateTemporaryServiceCredentialRequestPb struct {
	AzureOptions   *GenerateTemporaryServiceCredentialAzureOptions `json:"azure_options,omitempty"`
	CredentialName string                                          `json:"credential_name"`
	GcpOptions     *GenerateTemporaryServiceCredentialGcpOptions   `json:"gcp_options,omitempty"`
}

func generateTemporaryServiceCredentialRequestFromPb(pb *generateTemporaryServiceCredentialRequestPb) (*GenerateTemporaryServiceCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryServiceCredentialRequest{}
	st.AzureOptions = pb.AzureOptions
	st.CredentialName = pb.CredentialName
	st.GcpOptions = pb.GcpOptions

	return st, nil
}

func generateTemporaryTableCredentialRequestToPb(st *GenerateTemporaryTableCredentialRequest) (*generateTemporaryTableCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryTableCredentialRequestPb{}
	pb.Operation = st.Operation
	pb.TableId = st.TableId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type generateTemporaryTableCredentialRequestPb struct {
	Operation TableOperation `json:"operation,omitempty"`
	TableId   string         `json:"table_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func generateTemporaryTableCredentialRequestFromPb(pb *generateTemporaryTableCredentialRequestPb) (*GenerateTemporaryTableCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryTableCredentialRequest{}
	st.Operation = pb.Operation
	st.TableId = pb.TableId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *generateTemporaryTableCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st generateTemporaryTableCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func generateTemporaryTableCredentialResponseToPb(st *GenerateTemporaryTableCredentialResponse) (*generateTemporaryTableCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &generateTemporaryTableCredentialResponsePb{}
	pb.AwsTempCredentials = st.AwsTempCredentials
	pb.AzureAad = st.AzureAad
	pb.AzureUserDelegationSas = st.AzureUserDelegationSas
	pb.ExpirationTime = st.ExpirationTime
	pb.GcpOauthToken = st.GcpOauthToken
	pb.R2TempCredentials = st.R2TempCredentials
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type generateTemporaryTableCredentialResponsePb struct {
	AwsTempCredentials     *AwsCredentials            `json:"aws_temp_credentials,omitempty"`
	AzureAad               *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	AzureUserDelegationSas *AzureUserDelegationSas    `json:"azure_user_delegation_sas,omitempty"`
	ExpirationTime         int64                      `json:"expiration_time,omitempty"`
	GcpOauthToken          *GcpOauthToken             `json:"gcp_oauth_token,omitempty"`
	R2TempCredentials      *R2Credentials             `json:"r2_temp_credentials,omitempty"`
	Url                    string                     `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func generateTemporaryTableCredentialResponseFromPb(pb *generateTemporaryTableCredentialResponsePb) (*GenerateTemporaryTableCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenerateTemporaryTableCredentialResponse{}
	st.AwsTempCredentials = pb.AwsTempCredentials
	st.AzureAad = pb.AzureAad
	st.AzureUserDelegationSas = pb.AzureUserDelegationSas
	st.ExpirationTime = pb.ExpirationTime
	st.GcpOauthToken = pb.GcpOauthToken
	st.R2TempCredentials = pb.R2TempCredentials
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *generateTemporaryTableCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st generateTemporaryTableCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getAccountMetastoreAssignmentRequestToPb(st *GetAccountMetastoreAssignmentRequest) (*getAccountMetastoreAssignmentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountMetastoreAssignmentRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type getAccountMetastoreAssignmentRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

func getAccountMetastoreAssignmentRequestFromPb(pb *getAccountMetastoreAssignmentRequestPb) (*GetAccountMetastoreAssignmentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountMetastoreAssignmentRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func getAccountMetastoreRequestToPb(st *GetAccountMetastoreRequest) (*getAccountMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountMetastoreRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

type getAccountMetastoreRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
}

func getAccountMetastoreRequestFromPb(pb *getAccountMetastoreRequestPb) (*GetAccountMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountMetastoreRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

func getAccountStorageCredentialRequestToPb(st *GetAccountStorageCredentialRequest) (*getAccountStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountStorageCredentialRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.StorageCredentialName = st.StorageCredentialName

	return pb, nil
}

type getAccountStorageCredentialRequestPb struct {
	MetastoreId           string `json:"-" url:"-"`
	StorageCredentialName string `json:"-" url:"-"`
}

func getAccountStorageCredentialRequestFromPb(pb *getAccountStorageCredentialRequestPb) (*GetAccountStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountStorageCredentialRequest{}
	st.MetastoreId = pb.MetastoreId
	st.StorageCredentialName = pb.StorageCredentialName

	return st, nil
}

func getArtifactAllowlistRequestToPb(st *GetArtifactAllowlistRequest) (*getArtifactAllowlistRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getArtifactAllowlistRequestPb{}
	pb.ArtifactType = st.ArtifactType

	return pb, nil
}

type getArtifactAllowlistRequestPb struct {
	ArtifactType ArtifactType `json:"-" url:"-"`
}

func getArtifactAllowlistRequestFromPb(pb *getArtifactAllowlistRequestPb) (*GetArtifactAllowlistRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetArtifactAllowlistRequest{}
	st.ArtifactType = pb.ArtifactType

	return st, nil
}

func getBindingsRequestToPb(st *GetBindingsRequest) (*getBindingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBindingsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SecurableName = st.SecurableName
	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getBindingsRequestPb struct {
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SecurableName string `json:"-" url:"-"`
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getBindingsRequestFromPb(pb *getBindingsRequestPb) (*GetBindingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBindingsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SecurableName = pb.SecurableName
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getBindingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getBindingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getByAliasRequestToPb(st *GetByAliasRequest) (*getByAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getByAliasRequestPb{}
	pb.Alias = st.Alias
	pb.FullName = st.FullName
	pb.IncludeAliases = st.IncludeAliases

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getByAliasRequestPb struct {
	Alias          string `json:"-" url:"-"`
	FullName       string `json:"-" url:"-"`
	IncludeAliases bool   `json:"-" url:"include_aliases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getByAliasRequestFromPb(pb *getByAliasRequestPb) (*GetByAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetByAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName
	st.IncludeAliases = pb.IncludeAliases

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getByAliasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getByAliasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getCatalogRequestToPb(st *GetCatalogRequest) (*getCatalogRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCatalogRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getCatalogRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCatalogRequestFromPb(pb *getCatalogRequestPb) (*GetCatalogRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCatalogRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCatalogRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCatalogRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getCatalogWorkspaceBindingsResponseToPb(st *GetCatalogWorkspaceBindingsResponse) (*getCatalogWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCatalogWorkspaceBindingsResponsePb{}
	pb.Workspaces = st.Workspaces

	return pb, nil
}

type getCatalogWorkspaceBindingsResponsePb struct {
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func getCatalogWorkspaceBindingsResponseFromPb(pb *getCatalogWorkspaceBindingsResponsePb) (*GetCatalogWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCatalogWorkspaceBindingsResponse{}
	st.Workspaces = pb.Workspaces

	return st, nil
}

func getConnectionRequestToPb(st *GetConnectionRequest) (*getConnectionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getConnectionRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getConnectionRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getConnectionRequestFromPb(pb *getConnectionRequestPb) (*GetConnectionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetConnectionRequest{}
	st.Name = pb.Name

	return st, nil
}

func getCredentialRequestToPb(st *GetCredentialRequest) (*getCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialRequestPb{}
	pb.NameArg = st.NameArg

	return pb, nil
}

type getCredentialRequestPb struct {
	NameArg string `json:"-" url:"-"`
}

func getCredentialRequestFromPb(pb *getCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	st.NameArg = pb.NameArg

	return st, nil
}

func getEffectiveRequestToPb(st *GetEffectiveRequest) (*getEffectiveRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEffectiveRequestPb{}
	pb.FullName = st.FullName
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.Principal = st.Principal
	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getEffectiveRequestPb struct {
	FullName      string `json:"-" url:"-"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	Principal     string `json:"-" url:"principal,omitempty"`
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getEffectiveRequestFromPb(pb *getEffectiveRequestPb) (*GetEffectiveRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEffectiveRequest{}
	st.FullName = pb.FullName
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.Principal = pb.Principal
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEffectiveRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEffectiveRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getExternalLocationRequestToPb(st *GetExternalLocationRequest) (*getExternalLocationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExternalLocationRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getExternalLocationRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getExternalLocationRequestFromPb(pb *getExternalLocationRequestPb) (*GetExternalLocationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExternalLocationRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getExternalLocationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getExternalLocationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getFunctionRequestToPb(st *GetFunctionRequest) (*getFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getFunctionRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getFunctionRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getFunctionRequestFromPb(pb *getFunctionRequestPb) (*GetFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFunctionRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getGrantRequestToPb(st *GetGrantRequest) (*getGrantRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGrantRequestPb{}
	pb.FullName = st.FullName
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.Principal = st.Principal
	pb.SecurableType = st.SecurableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getGrantRequestPb struct {
	FullName      string `json:"-" url:"-"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	Principal     string `json:"-" url:"principal,omitempty"`
	SecurableType string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getGrantRequestFromPb(pb *getGrantRequestPb) (*GetGrantRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGrantRequest{}
	st.FullName = pb.FullName
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.Principal = pb.Principal
	st.SecurableType = pb.SecurableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getGrantRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getGrantRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getMetastoreRequestToPb(st *GetMetastoreRequest) (*getMetastoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetastoreRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getMetastoreRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getMetastoreRequestFromPb(pb *getMetastoreRequestPb) (*GetMetastoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetastoreRequest{}
	st.Id = pb.Id

	return st, nil
}

func getMetastoreSummaryResponseToPb(st *GetMetastoreSummaryResponse) (*getMetastoreSummaryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetastoreSummaryResponsePb{}
	pb.Cloud = st.Cloud
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DefaultDataAccessConfigId = st.DefaultDataAccessConfigId
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName
	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds
	pb.DeltaSharingScope = st.DeltaSharingScope
	pb.ExternalAccessEnabled = st.ExternalAccessEnabled
	pb.GlobalMetastoreId = st.GlobalMetastoreId
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.PrivilegeModelVersion = st.PrivilegeModelVersion
	pb.Region = st.Region
	pb.StorageRoot = st.StorageRoot
	pb.StorageRootCredentialId = st.StorageRootCredentialId
	pb.StorageRootCredentialName = st.StorageRootCredentialName
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getMetastoreSummaryResponsePb struct {
	Cloud                                       string                `json:"cloud,omitempty"`
	CreatedAt                                   int64                 `json:"created_at,omitempty"`
	CreatedBy                                   string                `json:"created_by,omitempty"`
	DefaultDataAccessConfigId                   string                `json:"default_data_access_config_id,omitempty"`
	DeltaSharingOrganizationName                string                `json:"delta_sharing_organization_name,omitempty"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64                 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingScope                           DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	ExternalAccessEnabled                       bool                  `json:"external_access_enabled,omitempty"`
	GlobalMetastoreId                           string                `json:"global_metastore_id,omitempty"`
	MetastoreId                                 string                `json:"metastore_id,omitempty"`
	Name                                        string                `json:"name,omitempty"`
	Owner                                       string                `json:"owner,omitempty"`
	PrivilegeModelVersion                       string                `json:"privilege_model_version,omitempty"`
	Region                                      string                `json:"region,omitempty"`
	StorageRoot                                 string                `json:"storage_root,omitempty"`
	StorageRootCredentialId                     string                `json:"storage_root_credential_id,omitempty"`
	StorageRootCredentialName                   string                `json:"storage_root_credential_name,omitempty"`
	UpdatedAt                                   int64                 `json:"updated_at,omitempty"`
	UpdatedBy                                   string                `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getMetastoreSummaryResponseFromPb(pb *getMetastoreSummaryResponsePb) (*GetMetastoreSummaryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetastoreSummaryResponse{}
	st.Cloud = pb.Cloud
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DefaultDataAccessConfigId = pb.DefaultDataAccessConfigId
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	st.DeltaSharingScope = pb.DeltaSharingScope
	st.ExternalAccessEnabled = pb.ExternalAccessEnabled
	st.GlobalMetastoreId = pb.GlobalMetastoreId
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.Region = pb.Region
	st.StorageRoot = pb.StorageRoot
	st.StorageRootCredentialId = pb.StorageRootCredentialId
	st.StorageRootCredentialName = pb.StorageRootCredentialName
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getMetastoreSummaryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getMetastoreSummaryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getModelVersionRequestToPb(st *GetModelVersionRequest) (*getModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeAliases = st.IncludeAliases
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getModelVersionRequestPb struct {
	FullName       string `json:"-" url:"-"`
	IncludeAliases bool   `json:"-" url:"include_aliases,omitempty"`
	IncludeBrowse  bool   `json:"-" url:"include_browse,omitempty"`
	Version        int    `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getModelVersionRequestFromPb(pb *getModelVersionRequestPb) (*GetModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionRequest{}
	st.FullName = pb.FullName
	st.IncludeAliases = pb.IncludeAliases
	st.IncludeBrowse = pb.IncludeBrowse
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getOnlineTableRequestToPb(st *GetOnlineTableRequest) (*getOnlineTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOnlineTableRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getOnlineTableRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getOnlineTableRequestFromPb(pb *getOnlineTableRequestPb) (*GetOnlineTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOnlineTableRequest{}
	st.Name = pb.Name

	return st, nil
}

func getPermissionsResponseToPb(st *GetPermissionsResponse) (*getPermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPermissionsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PrivilegeAssignments = st.PrivilegeAssignments

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPermissionsResponsePb struct {
	NextPageToken        string                `json:"next_page_token,omitempty"`
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPermissionsResponseFromPb(pb *getPermissionsResponsePb) (*GetPermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPermissionsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PrivilegeAssignments = pb.PrivilegeAssignments

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPermissionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPermissionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getQualityMonitorRequestToPb(st *GetQualityMonitorRequest) (*getQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQualityMonitorRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

type getQualityMonitorRequestPb struct {
	TableName string `json:"-" url:"-"`
}

func getQualityMonitorRequestFromPb(pb *getQualityMonitorRequestPb) (*GetQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQualityMonitorRequest{}
	st.TableName = pb.TableName

	return st, nil
}

func getQuotaRequestToPb(st *GetQuotaRequest) (*getQuotaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQuotaRequestPb{}
	pb.ParentFullName = st.ParentFullName
	pb.ParentSecurableType = st.ParentSecurableType
	pb.QuotaName = st.QuotaName

	return pb, nil
}

type getQuotaRequestPb struct {
	ParentFullName      string `json:"-" url:"-"`
	ParentSecurableType string `json:"-" url:"-"`
	QuotaName           string `json:"-" url:"-"`
}

func getQuotaRequestFromPb(pb *getQuotaRequestPb) (*GetQuotaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQuotaRequest{}
	st.ParentFullName = pb.ParentFullName
	st.ParentSecurableType = pb.ParentSecurableType
	st.QuotaName = pb.QuotaName

	return st, nil
}

func getQuotaResponseToPb(st *GetQuotaResponse) (*getQuotaResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQuotaResponsePb{}
	pb.QuotaInfo = st.QuotaInfo

	return pb, nil
}

type getQuotaResponsePb struct {
	QuotaInfo *QuotaInfo `json:"quota_info,omitempty"`
}

func getQuotaResponseFromPb(pb *getQuotaResponsePb) (*GetQuotaResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQuotaResponse{}
	st.QuotaInfo = pb.QuotaInfo

	return st, nil
}

func getRefreshRequestToPb(st *GetRefreshRequest) (*getRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRefreshRequestPb{}
	pb.RefreshId = st.RefreshId
	pb.TableName = st.TableName

	return pb, nil
}

type getRefreshRequestPb struct {
	RefreshId string `json:"-" url:"-"`
	TableName string `json:"-" url:"-"`
}

func getRefreshRequestFromPb(pb *getRefreshRequestPb) (*GetRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRefreshRequest{}
	st.RefreshId = pb.RefreshId
	st.TableName = pb.TableName

	return st, nil
}

func getRegisteredModelRequestToPb(st *GetRegisteredModelRequest) (*getRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeAliases = st.IncludeAliases
	pb.IncludeBrowse = st.IncludeBrowse

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getRegisteredModelRequestPb struct {
	FullName       string `json:"-" url:"-"`
	IncludeAliases bool   `json:"-" url:"include_aliases,omitempty"`
	IncludeBrowse  bool   `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRegisteredModelRequestFromPb(pb *getRegisteredModelRequestPb) (*GetRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelRequest{}
	st.FullName = pb.FullName
	st.IncludeAliases = pb.IncludeAliases
	st.IncludeBrowse = pb.IncludeBrowse

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getSchemaRequestToPb(st *GetSchemaRequest) (*getSchemaRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSchemaRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeBrowse = st.IncludeBrowse

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getSchemaRequestPb struct {
	FullName      string `json:"-" url:"-"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getSchemaRequestFromPb(pb *getSchemaRequestPb) (*GetSchemaRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSchemaRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getSchemaRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getSchemaRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getStorageCredentialRequestToPb(st *GetStorageCredentialRequest) (*getStorageCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStorageCredentialRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getStorageCredentialRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getStorageCredentialRequestFromPb(pb *getStorageCredentialRequestPb) (*GetStorageCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStorageCredentialRequest{}
	st.Name = pb.Name

	return st, nil
}

func getTableRequestToPb(st *GetTableRequest) (*getTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTableRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.IncludeDeltaMetadata = st.IncludeDeltaMetadata
	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getTableRequestPb struct {
	FullName                    string `json:"-" url:"-"`
	IncludeBrowse               bool   `json:"-" url:"include_browse,omitempty"`
	IncludeDeltaMetadata        bool   `json:"-" url:"include_delta_metadata,omitempty"`
	IncludeManifestCapabilities bool   `json:"-" url:"include_manifest_capabilities,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getTableRequestFromPb(pb *getTableRequestPb) (*GetTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTableRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse
	st.IncludeDeltaMetadata = pb.IncludeDeltaMetadata
	st.IncludeManifestCapabilities = pb.IncludeManifestCapabilities

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getTableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getTableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getWorkspaceBindingRequestToPb(st *GetWorkspaceBindingRequest) (*getWorkspaceBindingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceBindingRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getWorkspaceBindingRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getWorkspaceBindingRequestFromPb(pb *getWorkspaceBindingRequestPb) (*GetWorkspaceBindingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceBindingRequest{}
	st.Name = pb.Name

	return st, nil
}

func getWorkspaceBindingsResponseToPb(st *GetWorkspaceBindingsResponse) (*getWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceBindingsResponsePb{}
	pb.Bindings = st.Bindings
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getWorkspaceBindingsResponsePb struct {
	Bindings      []WorkspaceBinding `json:"bindings,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getWorkspaceBindingsResponseFromPb(pb *getWorkspaceBindingsResponsePb) (*GetWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceBindingsResponse{}
	st.Bindings = pb.Bindings
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getWorkspaceBindingsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getWorkspaceBindingsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAccountMetastoreAssignmentsRequestToPb(st *ListAccountMetastoreAssignmentsRequest) (*listAccountMetastoreAssignmentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountMetastoreAssignmentsRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

type listAccountMetastoreAssignmentsRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
}

func listAccountMetastoreAssignmentsRequestFromPb(pb *listAccountMetastoreAssignmentsRequestPb) (*ListAccountMetastoreAssignmentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountMetastoreAssignmentsRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

func listAccountMetastoreAssignmentsResponseToPb(st *ListAccountMetastoreAssignmentsResponse) (*listAccountMetastoreAssignmentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountMetastoreAssignmentsResponsePb{}
	pb.WorkspaceIds = st.WorkspaceIds

	return pb, nil
}

type listAccountMetastoreAssignmentsResponsePb struct {
	WorkspaceIds []int64 `json:"workspace_ids,omitempty"`
}

func listAccountMetastoreAssignmentsResponseFromPb(pb *listAccountMetastoreAssignmentsResponsePb) (*ListAccountMetastoreAssignmentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountMetastoreAssignmentsResponse{}
	st.WorkspaceIds = pb.WorkspaceIds

	return st, nil
}

func listAccountStorageCredentialsRequestToPb(st *ListAccountStorageCredentialsRequest) (*listAccountStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountStorageCredentialsRequestPb{}
	pb.MetastoreId = st.MetastoreId

	return pb, nil
}

type listAccountStorageCredentialsRequestPb struct {
	MetastoreId string `json:"-" url:"-"`
}

func listAccountStorageCredentialsRequestFromPb(pb *listAccountStorageCredentialsRequestPb) (*ListAccountStorageCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountStorageCredentialsRequest{}
	st.MetastoreId = pb.MetastoreId

	return st, nil
}

func listAccountStorageCredentialsResponseToPb(st *ListAccountStorageCredentialsResponse) (*listAccountStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountStorageCredentialsResponsePb{}
	pb.StorageCredentials = st.StorageCredentials

	return pb, nil
}

type listAccountStorageCredentialsResponsePb struct {
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`
}

func listAccountStorageCredentialsResponseFromPb(pb *listAccountStorageCredentialsResponsePb) (*ListAccountStorageCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountStorageCredentialsResponse{}
	st.StorageCredentials = pb.StorageCredentials

	return st, nil
}

func listCatalogsRequestToPb(st *ListCatalogsRequest) (*listCatalogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCatalogsRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCatalogsRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCatalogsRequestFromPb(pb *listCatalogsRequestPb) (*ListCatalogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCatalogsRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCatalogsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCatalogsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listCatalogsResponseToPb(st *ListCatalogsResponse) (*listCatalogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCatalogsResponsePb{}
	pb.Catalogs = st.Catalogs
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCatalogsResponsePb struct {
	Catalogs      []CatalogInfo `json:"catalogs,omitempty"`
	NextPageToken string        `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCatalogsResponseFromPb(pb *listCatalogsResponsePb) (*ListCatalogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCatalogsResponse{}
	st.Catalogs = pb.Catalogs
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCatalogsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCatalogsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listConnectionsRequestToPb(st *ListConnectionsRequest) (*listConnectionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listConnectionsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listConnectionsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listConnectionsRequestFromPb(pb *listConnectionsRequestPb) (*ListConnectionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListConnectionsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listConnectionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listConnectionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listConnectionsResponseToPb(st *ListConnectionsResponse) (*listConnectionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listConnectionsResponsePb{}
	pb.Connections = st.Connections
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listConnectionsResponsePb struct {
	Connections   []ConnectionInfo `json:"connections,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listConnectionsResponseFromPb(pb *listConnectionsResponsePb) (*ListConnectionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListConnectionsResponse{}
	st.Connections = pb.Connections
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listConnectionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listConnectionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listCredentialsRequestToPb(st *ListCredentialsRequest) (*listCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.Purpose = st.Purpose

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCredentialsRequestPb struct {
	MaxResults int               `json:"-" url:"max_results,omitempty"`
	PageToken  string            `json:"-" url:"page_token,omitempty"`
	Purpose    CredentialPurpose `json:"-" url:"purpose,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCredentialsRequestFromPb(pb *listCredentialsRequestPb) (*ListCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.Purpose = pb.Purpose

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listCredentialsResponseToPb(st *ListCredentialsResponse) (*listCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCredentialsResponsePb{}
	pb.Credentials = st.Credentials
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCredentialsResponsePb struct {
	Credentials   []CredentialInfo `json:"credentials,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCredentialsResponseFromPb(pb *listCredentialsResponsePb) (*ListCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCredentialsResponse{}
	st.Credentials = pb.Credentials
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExternalLocationsRequestToPb(st *ListExternalLocationsRequest) (*listExternalLocationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExternalLocationsRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExternalLocationsRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExternalLocationsRequestFromPb(pb *listExternalLocationsRequestPb) (*ListExternalLocationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLocationsRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExternalLocationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExternalLocationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listExternalLocationsResponseToPb(st *ListExternalLocationsResponse) (*listExternalLocationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExternalLocationsResponsePb{}
	pb.ExternalLocations = st.ExternalLocations
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listExternalLocationsResponsePb struct {
	ExternalLocations []ExternalLocationInfo `json:"external_locations,omitempty"`
	NextPageToken     string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExternalLocationsResponseFromPb(pb *listExternalLocationsResponsePb) (*ListExternalLocationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExternalLocationsResponse{}
	st.ExternalLocations = pb.ExternalLocations
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExternalLocationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExternalLocationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFunctionsRequestToPb(st *ListFunctionsRequest) (*listFunctionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFunctionsRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFunctionsRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SchemaName    string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFunctionsRequestFromPb(pb *listFunctionsRequestPb) (*ListFunctionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFunctionsRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFunctionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFunctionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFunctionsResponseToPb(st *ListFunctionsResponse) (*listFunctionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFunctionsResponsePb{}
	pb.Functions = st.Functions
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFunctionsResponsePb struct {
	Functions     []FunctionInfo `json:"functions,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFunctionsResponseFromPb(pb *listFunctionsResponsePb) (*ListFunctionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFunctionsResponse{}
	st.Functions = pb.Functions
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFunctionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFunctionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listMetastoresRequestToPb(st *ListMetastoresRequest) (*listMetastoresRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listMetastoresRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listMetastoresRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listMetastoresRequestFromPb(pb *listMetastoresRequestPb) (*ListMetastoresRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListMetastoresRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listMetastoresRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listMetastoresRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listMetastoresResponseToPb(st *ListMetastoresResponse) (*listMetastoresResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listMetastoresResponsePb{}
	pb.Metastores = st.Metastores
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listMetastoresResponsePb struct {
	Metastores    []MetastoreInfo `json:"metastores,omitempty"`
	NextPageToken string          `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listMetastoresResponseFromPb(pb *listMetastoresResponsePb) (*ListMetastoresResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListMetastoresResponse{}
	st.Metastores = pb.Metastores
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listMetastoresResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listMetastoresResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listModelVersionsRequestToPb(st *ListModelVersionsRequest) (*listModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelVersionsRequestPb{}
	pb.FullName = st.FullName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listModelVersionsRequestPb struct {
	FullName      string `json:"-" url:"-"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelVersionsRequestFromPb(pb *listModelVersionsRequestPb) (*ListModelVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelVersionsRequest{}
	st.FullName = pb.FullName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelVersionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelVersionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listModelVersionsResponseToPb(st *ListModelVersionsResponse) (*listModelVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelVersionsResponsePb{}
	pb.ModelVersions = st.ModelVersions
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listModelVersionsResponsePb struct {
	ModelVersions []ModelVersionInfo `json:"model_versions,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelVersionsResponseFromPb(pb *listModelVersionsResponsePb) (*ListModelVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelVersionsResponse{}
	st.ModelVersions = pb.ModelVersions
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelVersionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelVersionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listQuotasRequestToPb(st *ListQuotasRequest) (*listQuotasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQuotasRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQuotasRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQuotasRequestFromPb(pb *listQuotasRequestPb) (*ListQuotasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQuotasRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQuotasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQuotasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listQuotasResponseToPb(st *ListQuotasResponse) (*listQuotasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQuotasResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Quotas = st.Quotas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQuotasResponsePb struct {
	NextPageToken string      `json:"next_page_token,omitempty"`
	Quotas        []QuotaInfo `json:"quotas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQuotasResponseFromPb(pb *listQuotasResponsePb) (*ListQuotasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQuotasResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Quotas = pb.Quotas

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQuotasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQuotasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRefreshesRequestToPb(st *ListRefreshesRequest) (*listRefreshesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRefreshesRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

type listRefreshesRequestPb struct {
	TableName string `json:"-" url:"-"`
}

func listRefreshesRequestFromPb(pb *listRefreshesRequestPb) (*ListRefreshesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRefreshesRequest{}
	st.TableName = pb.TableName

	return st, nil
}

func listRegisteredModelsRequestToPb(st *ListRegisteredModelsRequest) (*listRegisteredModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegisteredModelsRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRegisteredModelsRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name,omitempty"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SchemaName    string `json:"-" url:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRegisteredModelsRequestFromPb(pb *listRegisteredModelsRequestPb) (*ListRegisteredModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegisteredModelsRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRegisteredModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRegisteredModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRegisteredModelsResponseToPb(st *ListRegisteredModelsResponse) (*listRegisteredModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegisteredModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.RegisteredModels = st.RegisteredModels

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRegisteredModelsResponsePb struct {
	NextPageToken    string                `json:"next_page_token,omitempty"`
	RegisteredModels []RegisteredModelInfo `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRegisteredModelsResponseFromPb(pb *listRegisteredModelsResponsePb) (*ListRegisteredModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegisteredModelsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.RegisteredModels = pb.RegisteredModels

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRegisteredModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRegisteredModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSchemasRequestToPb(st *ListSchemasRequest) (*listSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchemasRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSchemasRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchemasRequestFromPb(pb *listSchemasRequestPb) (*ListSchemasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchemasRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSchemasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchemasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSchemasResponseToPb(st *ListSchemasResponse) (*listSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSchemasResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Schemas = st.Schemas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSchemasResponsePb struct {
	NextPageToken string       `json:"next_page_token,omitempty"`
	Schemas       []SchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSchemasResponseFromPb(pb *listSchemasResponsePb) (*ListSchemasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSchemasResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Schemas = pb.Schemas

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSchemasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSchemasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listStorageCredentialsRequestToPb(st *ListStorageCredentialsRequest) (*listStorageCredentialsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStorageCredentialsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listStorageCredentialsRequestPb struct {
	MaxResults int    `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listStorageCredentialsRequestFromPb(pb *listStorageCredentialsRequestPb) (*ListStorageCredentialsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStorageCredentialsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listStorageCredentialsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listStorageCredentialsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listStorageCredentialsResponseToPb(st *ListStorageCredentialsResponse) (*listStorageCredentialsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStorageCredentialsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.StorageCredentials = st.StorageCredentials

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listStorageCredentialsResponsePb struct {
	NextPageToken      string                  `json:"next_page_token,omitempty"`
	StorageCredentials []StorageCredentialInfo `json:"storage_credentials,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listStorageCredentialsResponseFromPb(pb *listStorageCredentialsResponsePb) (*ListStorageCredentialsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStorageCredentialsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.StorageCredentials = pb.StorageCredentials

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listStorageCredentialsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listStorageCredentialsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSummariesRequestToPb(st *ListSummariesRequest) (*listSummariesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSummariesRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaNamePattern = st.SchemaNamePattern
	pb.TableNamePattern = st.TableNamePattern

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSummariesRequestPb struct {
	CatalogName                 string `json:"-" url:"catalog_name"`
	IncludeManifestCapabilities bool   `json:"-" url:"include_manifest_capabilities,omitempty"`
	MaxResults                  int    `json:"-" url:"max_results,omitempty"`
	PageToken                   string `json:"-" url:"page_token,omitempty"`
	SchemaNamePattern           string `json:"-" url:"schema_name_pattern,omitempty"`
	TableNamePattern            string `json:"-" url:"table_name_pattern,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSummariesRequestFromPb(pb *listSummariesRequestPb) (*ListSummariesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSummariesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeManifestCapabilities = pb.IncludeManifestCapabilities
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaNamePattern = pb.SchemaNamePattern
	st.TableNamePattern = pb.TableNamePattern

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSummariesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSummariesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSystemSchemasRequestToPb(st *ListSystemSchemasRequest) (*listSystemSchemasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSystemSchemasRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.MetastoreId = st.MetastoreId
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSystemSchemasRequestPb struct {
	MaxResults  int    `json:"-" url:"max_results,omitempty"`
	MetastoreId string `json:"-" url:"-"`
	PageToken   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSystemSchemasRequestFromPb(pb *listSystemSchemasRequestPb) (*ListSystemSchemasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSystemSchemasRequest{}
	st.MaxResults = pb.MaxResults
	st.MetastoreId = pb.MetastoreId
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSystemSchemasRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSystemSchemasRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listSystemSchemasResponseToPb(st *ListSystemSchemasResponse) (*listSystemSchemasResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listSystemSchemasResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Schemas = st.Schemas

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listSystemSchemasResponsePb struct {
	NextPageToken string             `json:"next_page_token,omitempty"`
	Schemas       []SystemSchemaInfo `json:"schemas,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listSystemSchemasResponseFromPb(pb *listSystemSchemasResponsePb) (*ListSystemSchemasResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListSystemSchemasResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Schemas = pb.Schemas

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listSystemSchemasResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listSystemSchemasResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listTableSummariesResponseToPb(st *ListTableSummariesResponse) (*listTableSummariesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTableSummariesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Tables = st.Tables

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listTableSummariesResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Tables        []TableSummary `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTableSummariesResponseFromPb(pb *listTableSummariesResponsePb) (*ListTableSummariesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTableSummariesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Tables = pb.Tables

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTableSummariesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTableSummariesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listTablesRequestToPb(st *ListTablesRequest) (*listTablesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTablesRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.IncludeDeltaMetadata = st.IncludeDeltaMetadata
	pb.IncludeManifestCapabilities = st.IncludeManifestCapabilities
	pb.MaxResults = st.MaxResults
	pb.OmitColumns = st.OmitColumns
	pb.OmitProperties = st.OmitProperties
	pb.OmitUsername = st.OmitUsername
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listTablesRequestPb struct {
	CatalogName                 string `json:"-" url:"catalog_name"`
	IncludeBrowse               bool   `json:"-" url:"include_browse,omitempty"`
	IncludeDeltaMetadata        bool   `json:"-" url:"include_delta_metadata,omitempty"`
	IncludeManifestCapabilities bool   `json:"-" url:"include_manifest_capabilities,omitempty"`
	MaxResults                  int    `json:"-" url:"max_results,omitempty"`
	OmitColumns                 bool   `json:"-" url:"omit_columns,omitempty"`
	OmitProperties              bool   `json:"-" url:"omit_properties,omitempty"`
	OmitUsername                bool   `json:"-" url:"omit_username,omitempty"`
	PageToken                   string `json:"-" url:"page_token,omitempty"`
	SchemaName                  string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTablesRequestFromPb(pb *listTablesRequestPb) (*ListTablesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTablesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.IncludeDeltaMetadata = pb.IncludeDeltaMetadata
	st.IncludeManifestCapabilities = pb.IncludeManifestCapabilities
	st.MaxResults = pb.MaxResults
	st.OmitColumns = pb.OmitColumns
	st.OmitProperties = pb.OmitProperties
	st.OmitUsername = pb.OmitUsername
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTablesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTablesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listTablesResponseToPb(st *ListTablesResponse) (*listTablesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTablesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Tables = st.Tables

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listTablesResponsePb struct {
	NextPageToken string      `json:"next_page_token,omitempty"`
	Tables        []TableInfo `json:"tables,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTablesResponseFromPb(pb *listTablesResponsePb) (*ListTablesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTablesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Tables = pb.Tables

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTablesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTablesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listVolumesRequestToPb(st *ListVolumesRequest) (*listVolumesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVolumesRequestPb{}
	pb.CatalogName = st.CatalogName
	pb.IncludeBrowse = st.IncludeBrowse
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listVolumesRequestPb struct {
	CatalogName   string `json:"-" url:"catalog_name"`
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	SchemaName    string `json:"-" url:"schema_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVolumesRequestFromPb(pb *listVolumesRequestPb) (*ListVolumesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVolumesRequest{}
	st.CatalogName = pb.CatalogName
	st.IncludeBrowse = pb.IncludeBrowse
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVolumesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVolumesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listVolumesResponseContentToPb(st *ListVolumesResponseContent) (*listVolumesResponseContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVolumesResponseContentPb{}
	pb.NextPageToken = st.NextPageToken
	pb.Volumes = st.Volumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listVolumesResponseContentPb struct {
	NextPageToken string       `json:"next_page_token,omitempty"`
	Volumes       []VolumeInfo `json:"volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVolumesResponseContentFromPb(pb *listVolumesResponseContentPb) (*ListVolumesResponseContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVolumesResponseContent{}
	st.NextPageToken = pb.NextPageToken
	st.Volumes = pb.Volumes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVolumesResponseContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVolumesResponseContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func metastoreAssignmentToPb(st *MetastoreAssignment) (*metastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type metastoreAssignmentPb struct {
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	MetastoreId        string `json:"metastore_id"`
	WorkspaceId        int64  `json:"workspace_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func metastoreAssignmentFromPb(pb *metastoreAssignmentPb) (*MetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MetastoreAssignment{}
	st.DefaultCatalogName = pb.DefaultCatalogName
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *metastoreAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st metastoreAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func metastoreInfoToPb(st *MetastoreInfo) (*metastoreInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metastoreInfoPb{}
	pb.Cloud = st.Cloud
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DefaultDataAccessConfigId = st.DefaultDataAccessConfigId
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName
	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds
	pb.DeltaSharingScope = st.DeltaSharingScope
	pb.ExternalAccessEnabled = st.ExternalAccessEnabled
	pb.GlobalMetastoreId = st.GlobalMetastoreId
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.PrivilegeModelVersion = st.PrivilegeModelVersion
	pb.Region = st.Region
	pb.StorageRoot = st.StorageRoot
	pb.StorageRootCredentialId = st.StorageRootCredentialId
	pb.StorageRootCredentialName = st.StorageRootCredentialName
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type metastoreInfoPb struct {
	Cloud                                       string                `json:"cloud,omitempty"`
	CreatedAt                                   int64                 `json:"created_at,omitempty"`
	CreatedBy                                   string                `json:"created_by,omitempty"`
	DefaultDataAccessConfigId                   string                `json:"default_data_access_config_id,omitempty"`
	DeltaSharingOrganizationName                string                `json:"delta_sharing_organization_name,omitempty"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64                 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingScope                           DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	ExternalAccessEnabled                       bool                  `json:"external_access_enabled,omitempty"`
	GlobalMetastoreId                           string                `json:"global_metastore_id,omitempty"`
	MetastoreId                                 string                `json:"metastore_id,omitempty"`
	Name                                        string                `json:"name,omitempty"`
	Owner                                       string                `json:"owner,omitempty"`
	PrivilegeModelVersion                       string                `json:"privilege_model_version,omitempty"`
	Region                                      string                `json:"region,omitempty"`
	StorageRoot                                 string                `json:"storage_root,omitempty"`
	StorageRootCredentialId                     string                `json:"storage_root_credential_id,omitempty"`
	StorageRootCredentialName                   string                `json:"storage_root_credential_name,omitempty"`
	UpdatedAt                                   int64                 `json:"updated_at,omitempty"`
	UpdatedBy                                   string                `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func metastoreInfoFromPb(pb *metastoreInfoPb) (*MetastoreInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MetastoreInfo{}
	st.Cloud = pb.Cloud
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DefaultDataAccessConfigId = pb.DefaultDataAccessConfigId
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	st.DeltaSharingScope = pb.DeltaSharingScope
	st.ExternalAccessEnabled = pb.ExternalAccessEnabled
	st.GlobalMetastoreId = pb.GlobalMetastoreId
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.Region = pb.Region
	st.StorageRoot = pb.StorageRoot
	st.StorageRootCredentialId = pb.StorageRootCredentialId
	st.StorageRootCredentialName = pb.StorageRootCredentialName
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *metastoreInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st metastoreInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func modelVersionInfoToPb(st *ModelVersionInfo) (*modelVersionInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionInfoPb{}
	pb.Aliases = st.Aliases
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Id = st.Id
	pb.MetastoreId = st.MetastoreId
	pb.ModelName = st.ModelName
	pb.ModelVersionDependencies = st.ModelVersionDependencies
	pb.RunId = st.RunId
	pb.RunWorkspaceId = st.RunWorkspaceId
	pb.SchemaName = st.SchemaName
	pb.Source = st.Source
	pb.Status = st.Status
	pb.StorageLocation = st.StorageLocation
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type modelVersionInfoPb struct {
	Aliases                  []RegisteredModelAlias `json:"aliases,omitempty"`
	BrowseOnly               bool                   `json:"browse_only,omitempty"`
	CatalogName              string                 `json:"catalog_name,omitempty"`
	Comment                  string                 `json:"comment,omitempty"`
	CreatedAt                int64                  `json:"created_at,omitempty"`
	CreatedBy                string                 `json:"created_by,omitempty"`
	Id                       string                 `json:"id,omitempty"`
	MetastoreId              string                 `json:"metastore_id,omitempty"`
	ModelName                string                 `json:"model_name,omitempty"`
	ModelVersionDependencies *DependencyList        `json:"model_version_dependencies,omitempty"`
	RunId                    string                 `json:"run_id,omitempty"`
	RunWorkspaceId           int                    `json:"run_workspace_id,omitempty"`
	SchemaName               string                 `json:"schema_name,omitempty"`
	Source                   string                 `json:"source,omitempty"`
	Status                   ModelVersionInfoStatus `json:"status,omitempty"`
	StorageLocation          string                 `json:"storage_location,omitempty"`
	UpdatedAt                int64                  `json:"updated_at,omitempty"`
	UpdatedBy                string                 `json:"updated_by,omitempty"`
	Version                  int                    `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionInfoFromPb(pb *modelVersionInfoPb) (*ModelVersionInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionInfo{}
	st.Aliases = pb.Aliases
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Id = pb.Id
	st.MetastoreId = pb.MetastoreId
	st.ModelName = pb.ModelName
	st.ModelVersionDependencies = pb.ModelVersionDependencies
	st.RunId = pb.RunId
	st.RunWorkspaceId = pb.RunWorkspaceId
	st.SchemaName = pb.SchemaName
	st.Source = pb.Source
	st.Status = pb.Status
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func monitorCronScheduleToPb(st *MonitorCronSchedule) (*monitorCronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorCronSchedulePb{}
	pb.PauseStatus = st.PauseStatus
	pb.QuartzCronExpression = st.QuartzCronExpression
	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

type monitorCronSchedulePb struct {
	PauseStatus          MonitorCronSchedulePauseStatus `json:"pause_status,omitempty"`
	QuartzCronExpression string                         `json:"quartz_cron_expression"`
	TimezoneId           string                         `json:"timezone_id"`
}

func monitorCronScheduleFromPb(pb *monitorCronSchedulePb) (*MonitorCronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorCronSchedule{}
	st.PauseStatus = pb.PauseStatus
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

func monitorDataClassificationConfigToPb(st *MonitorDataClassificationConfig) (*monitorDataClassificationConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorDataClassificationConfigPb{}
	pb.Enabled = st.Enabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type monitorDataClassificationConfigPb struct {
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorDataClassificationConfigFromPb(pb *monitorDataClassificationConfigPb) (*MonitorDataClassificationConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorDataClassificationConfig{}
	st.Enabled = pb.Enabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorDataClassificationConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorDataClassificationConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func monitorDestinationToPb(st *MonitorDestination) (*monitorDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorDestinationPb{}
	pb.EmailAddresses = st.EmailAddresses

	return pb, nil
}

type monitorDestinationPb struct {
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

func monitorDestinationFromPb(pb *monitorDestinationPb) (*MonitorDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorDestination{}
	st.EmailAddresses = pb.EmailAddresses

	return st, nil
}

func monitorInferenceLogToPb(st *MonitorInferenceLog) (*monitorInferenceLogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorInferenceLogPb{}
	pb.Granularities = st.Granularities
	pb.LabelCol = st.LabelCol
	pb.ModelIdCol = st.ModelIdCol
	pb.PredictionCol = st.PredictionCol
	pb.PredictionProbaCol = st.PredictionProbaCol
	pb.ProblemType = st.ProblemType
	pb.TimestampCol = st.TimestampCol

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type monitorInferenceLogPb struct {
	Granularities      []string                       `json:"granularities"`
	LabelCol           string                         `json:"label_col,omitempty"`
	ModelIdCol         string                         `json:"model_id_col"`
	PredictionCol      string                         `json:"prediction_col"`
	PredictionProbaCol string                         `json:"prediction_proba_col,omitempty"`
	ProblemType        MonitorInferenceLogProblemType `json:"problem_type"`
	TimestampCol       string                         `json:"timestamp_col"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorInferenceLogFromPb(pb *monitorInferenceLogPb) (*MonitorInferenceLog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorInferenceLog{}
	st.Granularities = pb.Granularities
	st.LabelCol = pb.LabelCol
	st.ModelIdCol = pb.ModelIdCol
	st.PredictionCol = pb.PredictionCol
	st.PredictionProbaCol = pb.PredictionProbaCol
	st.ProblemType = pb.ProblemType
	st.TimestampCol = pb.TimestampCol

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorInferenceLogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorInferenceLogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func monitorInfoToPb(st *MonitorInfo) (*monitorInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorInfoPb{}
	pb.AssetsDir = st.AssetsDir
	pb.BaselineTableName = st.BaselineTableName
	pb.CustomMetrics = st.CustomMetrics
	pb.DashboardId = st.DashboardId
	pb.DataClassificationConfig = st.DataClassificationConfig
	pb.DriftMetricsTableName = st.DriftMetricsTableName
	pb.InferenceLog = st.InferenceLog
	pb.LatestMonitorFailureMsg = st.LatestMonitorFailureMsg
	pb.MonitorVersion = st.MonitorVersion
	pb.Notifications = st.Notifications
	pb.OutputSchemaName = st.OutputSchemaName
	pb.ProfileMetricsTableName = st.ProfileMetricsTableName
	pb.Schedule = st.Schedule
	pb.SlicingExprs = st.SlicingExprs
	pb.Snapshot = st.Snapshot
	pb.Status = st.Status
	pb.TableName = st.TableName
	pb.TimeSeries = st.TimeSeries

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type monitorInfoPb struct {
	AssetsDir                string                           `json:"assets_dir,omitempty"`
	BaselineTableName        string                           `json:"baseline_table_name,omitempty"`
	CustomMetrics            []MonitorMetric                  `json:"custom_metrics,omitempty"`
	DashboardId              string                           `json:"dashboard_id,omitempty"`
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	DriftMetricsTableName    string                           `json:"drift_metrics_table_name"`
	InferenceLog             *MonitorInferenceLog             `json:"inference_log,omitempty"`
	LatestMonitorFailureMsg  string                           `json:"latest_monitor_failure_msg,omitempty"`
	MonitorVersion           string                           `json:"monitor_version"`
	Notifications            *MonitorNotifications            `json:"notifications,omitempty"`
	OutputSchemaName         string                           `json:"output_schema_name,omitempty"`
	ProfileMetricsTableName  string                           `json:"profile_metrics_table_name"`
	Schedule                 *MonitorCronSchedule             `json:"schedule,omitempty"`
	SlicingExprs             []string                         `json:"slicing_exprs,omitempty"`
	Snapshot                 *MonitorSnapshot                 `json:"snapshot,omitempty"`
	Status                   MonitorInfoStatus                `json:"status"`
	TableName                string                           `json:"table_name"`
	TimeSeries               *MonitorTimeSeries               `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorInfoFromPb(pb *monitorInfoPb) (*MonitorInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorInfo{}
	st.AssetsDir = pb.AssetsDir
	st.BaselineTableName = pb.BaselineTableName
	st.CustomMetrics = pb.CustomMetrics
	st.DashboardId = pb.DashboardId
	st.DataClassificationConfig = pb.DataClassificationConfig
	st.DriftMetricsTableName = pb.DriftMetricsTableName
	st.InferenceLog = pb.InferenceLog
	st.LatestMonitorFailureMsg = pb.LatestMonitorFailureMsg
	st.MonitorVersion = pb.MonitorVersion
	st.Notifications = pb.Notifications
	st.OutputSchemaName = pb.OutputSchemaName
	st.ProfileMetricsTableName = pb.ProfileMetricsTableName
	st.Schedule = pb.Schedule
	st.SlicingExprs = pb.SlicingExprs
	st.Snapshot = pb.Snapshot
	st.Status = pb.Status
	st.TableName = pb.TableName
	st.TimeSeries = pb.TimeSeries

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func monitorMetricToPb(st *MonitorMetric) (*monitorMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorMetricPb{}
	pb.Definition = st.Definition
	pb.InputColumns = st.InputColumns
	pb.Name = st.Name
	pb.OutputDataType = st.OutputDataType
	pb.Type = st.Type

	return pb, nil
}

type monitorMetricPb struct {
	Definition     string            `json:"definition"`
	InputColumns   []string          `json:"input_columns"`
	Name           string            `json:"name"`
	OutputDataType string            `json:"output_data_type"`
	Type           MonitorMetricType `json:"type"`
}

func monitorMetricFromPb(pb *monitorMetricPb) (*MonitorMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorMetric{}
	st.Definition = pb.Definition
	st.InputColumns = pb.InputColumns
	st.Name = pb.Name
	st.OutputDataType = pb.OutputDataType
	st.Type = pb.Type

	return st, nil
}

func monitorNotificationsToPb(st *MonitorNotifications) (*monitorNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorNotificationsPb{}
	pb.OnFailure = st.OnFailure
	pb.OnNewClassificationTagDetected = st.OnNewClassificationTagDetected

	return pb, nil
}

type monitorNotificationsPb struct {
	OnFailure                      *MonitorDestination `json:"on_failure,omitempty"`
	OnNewClassificationTagDetected *MonitorDestination `json:"on_new_classification_tag_detected,omitempty"`
}

func monitorNotificationsFromPb(pb *monitorNotificationsPb) (*MonitorNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorNotifications{}
	st.OnFailure = pb.OnFailure
	st.OnNewClassificationTagDetected = pb.OnNewClassificationTagDetected

	return st, nil
}

func monitorRefreshInfoToPb(st *MonitorRefreshInfo) (*monitorRefreshInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorRefreshInfoPb{}
	pb.EndTimeMs = st.EndTimeMs
	pb.Message = st.Message
	pb.RefreshId = st.RefreshId
	pb.StartTimeMs = st.StartTimeMs
	pb.State = st.State
	pb.Trigger = st.Trigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type monitorRefreshInfoPb struct {
	EndTimeMs   int64                     `json:"end_time_ms,omitempty"`
	Message     string                    `json:"message,omitempty"`
	RefreshId   int64                     `json:"refresh_id"`
	StartTimeMs int64                     `json:"start_time_ms"`
	State       MonitorRefreshInfoState   `json:"state"`
	Trigger     MonitorRefreshInfoTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func monitorRefreshInfoFromPb(pb *monitorRefreshInfoPb) (*MonitorRefreshInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorRefreshInfo{}
	st.EndTimeMs = pb.EndTimeMs
	st.Message = pb.Message
	st.RefreshId = pb.RefreshId
	st.StartTimeMs = pb.StartTimeMs
	st.State = pb.State
	st.Trigger = pb.Trigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *monitorRefreshInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st monitorRefreshInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func monitorRefreshListResponseToPb(st *MonitorRefreshListResponse) (*monitorRefreshListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorRefreshListResponsePb{}
	pb.Refreshes = st.Refreshes

	return pb, nil
}

type monitorRefreshListResponsePb struct {
	Refreshes []MonitorRefreshInfo `json:"refreshes,omitempty"`
}

func monitorRefreshListResponseFromPb(pb *monitorRefreshListResponsePb) (*MonitorRefreshListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorRefreshListResponse{}
	st.Refreshes = pb.Refreshes

	return st, nil
}

func monitorSnapshotToPb(st *MonitorSnapshot) (*monitorSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorSnapshotPb{}

	return pb, nil
}

type monitorSnapshotPb struct {
}

func monitorSnapshotFromPb(pb *monitorSnapshotPb) (*MonitorSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorSnapshot{}

	return st, nil
}

func monitorTimeSeriesToPb(st *MonitorTimeSeries) (*monitorTimeSeriesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &monitorTimeSeriesPb{}
	pb.Granularities = st.Granularities
	pb.TimestampCol = st.TimestampCol

	return pb, nil
}

type monitorTimeSeriesPb struct {
	Granularities []string `json:"granularities"`
	TimestampCol  string   `json:"timestamp_col"`
}

func monitorTimeSeriesFromPb(pb *monitorTimeSeriesPb) (*MonitorTimeSeries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MonitorTimeSeries{}
	st.Granularities = pb.Granularities
	st.TimestampCol = pb.TimestampCol

	return st, nil
}

func namedTableConstraintToPb(st *NamedTableConstraint) (*namedTableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &namedTableConstraintPb{}
	pb.Name = st.Name

	return pb, nil
}

type namedTableConstraintPb struct {
	Name string `json:"name"`
}

func namedTableConstraintFromPb(pb *namedTableConstraintPb) (*NamedTableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NamedTableConstraint{}
	st.Name = pb.Name

	return st, nil
}

func onlineTableToPb(st *OnlineTable) (*onlineTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTablePb{}
	pb.Name = st.Name
	pb.Spec = st.Spec
	pb.Status = st.Status
	pb.TableServingUrl = st.TableServingUrl
	pb.UnityCatalogProvisioningState = st.UnityCatalogProvisioningState

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type onlineTablePb struct {
	Name                          string                `json:"name,omitempty"`
	Spec                          *OnlineTableSpec      `json:"spec,omitempty"`
	Status                        *OnlineTableStatus    `json:"status,omitempty"`
	TableServingUrl               string                `json:"table_serving_url,omitempty"`
	UnityCatalogProvisioningState ProvisioningInfoState `json:"unity_catalog_provisioning_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineTableFromPb(pb *onlineTablePb) (*OnlineTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTable{}
	st.Name = pb.Name
	st.Spec = pb.Spec
	st.Status = pb.Status
	st.TableServingUrl = pb.TableServingUrl
	st.UnityCatalogProvisioningState = pb.UnityCatalogProvisioningState

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func onlineTableSpecToPb(st *OnlineTableSpec) (*onlineTableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecPb{}
	pb.PerformFullCopy = st.PerformFullCopy
	pb.PipelineId = st.PipelineId
	pb.PrimaryKeyColumns = st.PrimaryKeyColumns
	pb.RunContinuously = st.RunContinuously
	pb.RunTriggered = st.RunTriggered
	pb.SourceTableFullName = st.SourceTableFullName
	pb.TimeseriesKey = st.TimeseriesKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type onlineTableSpecPb struct {
	PerformFullCopy     bool                                       `json:"perform_full_copy,omitempty"`
	PipelineId          string                                     `json:"pipeline_id,omitempty"`
	PrimaryKeyColumns   []string                                   `json:"primary_key_columns,omitempty"`
	RunContinuously     *OnlineTableSpecContinuousSchedulingPolicy `json:"run_continuously,omitempty"`
	RunTriggered        *OnlineTableSpecTriggeredSchedulingPolicy  `json:"run_triggered,omitempty"`
	SourceTableFullName string                                     `json:"source_table_full_name,omitempty"`
	TimeseriesKey       string                                     `json:"timeseries_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineTableSpecFromPb(pb *onlineTableSpecPb) (*OnlineTableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpec{}
	st.PerformFullCopy = pb.PerformFullCopy
	st.PipelineId = pb.PipelineId
	st.PrimaryKeyColumns = pb.PrimaryKeyColumns
	st.RunContinuously = pb.RunContinuously
	st.RunTriggered = pb.RunTriggered
	st.SourceTableFullName = pb.SourceTableFullName
	st.TimeseriesKey = pb.TimeseriesKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineTableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineTableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func onlineTableSpecContinuousSchedulingPolicyToPb(st *OnlineTableSpecContinuousSchedulingPolicy) (*onlineTableSpecContinuousSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecContinuousSchedulingPolicyPb{}

	return pb, nil
}

type onlineTableSpecContinuousSchedulingPolicyPb struct {
}

func onlineTableSpecContinuousSchedulingPolicyFromPb(pb *onlineTableSpecContinuousSchedulingPolicyPb) (*OnlineTableSpecContinuousSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpecContinuousSchedulingPolicy{}

	return st, nil
}

func onlineTableSpecTriggeredSchedulingPolicyToPb(st *OnlineTableSpecTriggeredSchedulingPolicy) (*onlineTableSpecTriggeredSchedulingPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableSpecTriggeredSchedulingPolicyPb{}

	return pb, nil
}

type onlineTableSpecTriggeredSchedulingPolicyPb struct {
}

func onlineTableSpecTriggeredSchedulingPolicyFromPb(pb *onlineTableSpecTriggeredSchedulingPolicyPb) (*OnlineTableSpecTriggeredSchedulingPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableSpecTriggeredSchedulingPolicy{}

	return st, nil
}

func onlineTableStatusToPb(st *OnlineTableStatus) (*onlineTableStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &onlineTableStatusPb{}
	pb.ContinuousUpdateStatus = st.ContinuousUpdateStatus
	pb.DetailedState = st.DetailedState
	pb.FailedStatus = st.FailedStatus
	pb.Message = st.Message
	pb.ProvisioningStatus = st.ProvisioningStatus
	pb.TriggeredUpdateStatus = st.TriggeredUpdateStatus

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type onlineTableStatusPb struct {
	ContinuousUpdateStatus *ContinuousUpdateStatus `json:"continuous_update_status,omitempty"`
	DetailedState          OnlineTableState        `json:"detailed_state,omitempty"`
	FailedStatus           *FailedStatus           `json:"failed_status,omitempty"`
	Message                string                  `json:"message,omitempty"`
	ProvisioningStatus     *ProvisioningStatus     `json:"provisioning_status,omitempty"`
	TriggeredUpdateStatus  *TriggeredUpdateStatus  `json:"triggered_update_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func onlineTableStatusFromPb(pb *onlineTableStatusPb) (*OnlineTableStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineTableStatus{}
	st.ContinuousUpdateStatus = pb.ContinuousUpdateStatus
	st.DetailedState = pb.DetailedState
	st.FailedStatus = pb.FailedStatus
	st.Message = pb.Message
	st.ProvisioningStatus = pb.ProvisioningStatus
	st.TriggeredUpdateStatus = pb.TriggeredUpdateStatus

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *onlineTableStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st onlineTableStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func permissionsChangeToPb(st *PermissionsChange) (*permissionsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permissionsChangePb{}
	pb.Add = st.Add
	pb.Principal = st.Principal
	pb.Remove = st.Remove

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type permissionsChangePb struct {
	Add       []Privilege `json:"add,omitempty"`
	Principal string      `json:"principal,omitempty"`
	Remove    []Privilege `json:"remove,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func permissionsChangeFromPb(pb *permissionsChangePb) (*PermissionsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermissionsChange{}
	st.Add = pb.Add
	st.Principal = pb.Principal
	st.Remove = pb.Remove

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *permissionsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st permissionsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineProgressToPb(st *PipelineProgress) (*pipelineProgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineProgressPb{}
	pb.EstimatedCompletionTimeSeconds = st.EstimatedCompletionTimeSeconds
	pb.LatestVersionCurrentlyProcessing = st.LatestVersionCurrentlyProcessing
	pb.SyncProgressCompletion = st.SyncProgressCompletion
	pb.SyncedRowCount = st.SyncedRowCount
	pb.TotalRowCount = st.TotalRowCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineProgressPb struct {
	EstimatedCompletionTimeSeconds   float64 `json:"estimated_completion_time_seconds,omitempty"`
	LatestVersionCurrentlyProcessing int64   `json:"latest_version_currently_processing,omitempty"`
	SyncProgressCompletion           float64 `json:"sync_progress_completion,omitempty"`
	SyncedRowCount                   int64   `json:"synced_row_count,omitempty"`
	TotalRowCount                    int64   `json:"total_row_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineProgressFromPb(pb *pipelineProgressPb) (*PipelineProgress, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineProgress{}
	st.EstimatedCompletionTimeSeconds = pb.EstimatedCompletionTimeSeconds
	st.LatestVersionCurrentlyProcessing = pb.LatestVersionCurrentlyProcessing
	st.SyncProgressCompletion = pb.SyncProgressCompletion
	st.SyncedRowCount = pb.SyncedRowCount
	st.TotalRowCount = pb.TotalRowCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineProgressPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineProgressPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func primaryKeyConstraintToPb(st *PrimaryKeyConstraint) (*primaryKeyConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &primaryKeyConstraintPb{}
	pb.ChildColumns = st.ChildColumns
	pb.Name = st.Name
	pb.TimeseriesColumns = st.TimeseriesColumns

	return pb, nil
}

type primaryKeyConstraintPb struct {
	ChildColumns      []string `json:"child_columns"`
	Name              string   `json:"name"`
	TimeseriesColumns []string `json:"timeseries_columns,omitempty"`
}

func primaryKeyConstraintFromPb(pb *primaryKeyConstraintPb) (*PrimaryKeyConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrimaryKeyConstraint{}
	st.ChildColumns = pb.ChildColumns
	st.Name = pb.Name
	st.TimeseriesColumns = pb.TimeseriesColumns

	return st, nil
}

func privilegeAssignmentToPb(st *PrivilegeAssignment) (*privilegeAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privilegeAssignmentPb{}
	pb.Principal = st.Principal
	pb.Privileges = st.Privileges

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type privilegeAssignmentPb struct {
	Principal  string      `json:"principal,omitempty"`
	Privileges []Privilege `json:"privileges,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func privilegeAssignmentFromPb(pb *privilegeAssignmentPb) (*PrivilegeAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivilegeAssignment{}
	st.Principal = pb.Principal
	st.Privileges = pb.Privileges

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *privilegeAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st privilegeAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func provisioningInfoToPb(st *ProvisioningInfo) (*provisioningInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningInfoPb{}
	pb.State = st.State

	return pb, nil
}

type provisioningInfoPb struct {
	State ProvisioningInfoState `json:"state,omitempty"`
}

func provisioningInfoFromPb(pb *provisioningInfoPb) (*ProvisioningInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProvisioningInfo{}
	st.State = pb.State

	return st, nil
}

func provisioningStatusToPb(st *ProvisioningStatus) (*provisioningStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningStatusPb{}
	pb.InitialPipelineSyncProgress = st.InitialPipelineSyncProgress

	return pb, nil
}

type provisioningStatusPb struct {
	InitialPipelineSyncProgress *PipelineProgress `json:"initial_pipeline_sync_progress,omitempty"`
}

func provisioningStatusFromPb(pb *provisioningStatusPb) (*ProvisioningStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ProvisioningStatus{}
	st.InitialPipelineSyncProgress = pb.InitialPipelineSyncProgress

	return st, nil
}

func quotaInfoToPb(st *QuotaInfo) (*quotaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &quotaInfoPb{}
	pb.LastRefreshedAt = st.LastRefreshedAt
	pb.ParentFullName = st.ParentFullName
	pb.ParentSecurableType = st.ParentSecurableType
	pb.QuotaCount = st.QuotaCount
	pb.QuotaLimit = st.QuotaLimit
	pb.QuotaName = st.QuotaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type quotaInfoPb struct {
	LastRefreshedAt     int64         `json:"last_refreshed_at,omitempty"`
	ParentFullName      string        `json:"parent_full_name,omitempty"`
	ParentSecurableType SecurableType `json:"parent_securable_type,omitempty"`
	QuotaCount          int           `json:"quota_count,omitempty"`
	QuotaLimit          int           `json:"quota_limit,omitempty"`
	QuotaName           string        `json:"quota_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func quotaInfoFromPb(pb *quotaInfoPb) (*QuotaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QuotaInfo{}
	st.LastRefreshedAt = pb.LastRefreshedAt
	st.ParentFullName = pb.ParentFullName
	st.ParentSecurableType = pb.ParentSecurableType
	st.QuotaCount = pb.QuotaCount
	st.QuotaLimit = pb.QuotaLimit
	st.QuotaName = pb.QuotaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *quotaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st quotaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func r2CredentialsToPb(st *R2Credentials) (*r2CredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &r2CredentialsPb{}
	pb.AccessKeyId = st.AccessKeyId
	pb.SecretAccessKey = st.SecretAccessKey
	pb.SessionToken = st.SessionToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type r2CredentialsPb struct {
	AccessKeyId     string `json:"access_key_id,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`
	SessionToken    string `json:"session_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func r2CredentialsFromPb(pb *r2CredentialsPb) (*R2Credentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &R2Credentials{}
	st.AccessKeyId = pb.AccessKeyId
	st.SecretAccessKey = pb.SecretAccessKey
	st.SessionToken = pb.SessionToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *r2CredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st r2CredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func readVolumeRequestToPb(st *ReadVolumeRequest) (*readVolumeRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readVolumeRequestPb{}
	pb.IncludeBrowse = st.IncludeBrowse
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type readVolumeRequestPb struct {
	IncludeBrowse bool   `json:"-" url:"include_browse,omitempty"`
	Name          string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func readVolumeRequestFromPb(pb *readVolumeRequestPb) (*ReadVolumeRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadVolumeRequest{}
	st.IncludeBrowse = pb.IncludeBrowse
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *readVolumeRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st readVolumeRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func regenerateDashboardRequestToPb(st *RegenerateDashboardRequest) (*regenerateDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regenerateDashboardRequestPb{}
	pb.TableName = st.TableName
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type regenerateDashboardRequestPb struct {
	TableName   string `json:"-" url:"-"`
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func regenerateDashboardRequestFromPb(pb *regenerateDashboardRequestPb) (*RegenerateDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegenerateDashboardRequest{}
	st.TableName = pb.TableName
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *regenerateDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st regenerateDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func regenerateDashboardResponseToPb(st *RegenerateDashboardResponse) (*regenerateDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &regenerateDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.ParentFolder = st.ParentFolder

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type regenerateDashboardResponsePb struct {
	DashboardId  string `json:"dashboard_id,omitempty"`
	ParentFolder string `json:"parent_folder,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func regenerateDashboardResponseFromPb(pb *regenerateDashboardResponsePb) (*RegenerateDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegenerateDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.ParentFolder = pb.ParentFolder

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *regenerateDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st regenerateDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelAliasToPb(st *RegisteredModelAlias) (*registeredModelAliasPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAliasPb{}
	pb.AliasName = st.AliasName
	pb.VersionNum = st.VersionNum

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelAliasPb struct {
	AliasName  string `json:"alias_name,omitempty"`
	VersionNum int    `json:"version_num,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelAliasFromPb(pb *registeredModelAliasPb) (*RegisteredModelAlias, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAlias{}
	st.AliasName = pb.AliasName
	st.VersionNum = pb.VersionNum

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAliasPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAliasPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func registeredModelInfoToPb(st *RegisteredModelInfo) (*registeredModelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelInfoPb{}
	pb.Aliases = st.Aliases
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.FullName = st.FullName
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.SchemaName = st.SchemaName
	pb.StorageLocation = st.StorageLocation
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type registeredModelInfoPb struct {
	Aliases         []RegisteredModelAlias `json:"aliases,omitempty"`
	BrowseOnly      bool                   `json:"browse_only,omitempty"`
	CatalogName     string                 `json:"catalog_name,omitempty"`
	Comment         string                 `json:"comment,omitempty"`
	CreatedAt       int64                  `json:"created_at,omitempty"`
	CreatedBy       string                 `json:"created_by,omitempty"`
	FullName        string                 `json:"full_name,omitempty"`
	MetastoreId     string                 `json:"metastore_id,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Owner           string                 `json:"owner,omitempty"`
	SchemaName      string                 `json:"schema_name,omitempty"`
	StorageLocation string                 `json:"storage_location,omitempty"`
	UpdatedAt       int64                  `json:"updated_at,omitempty"`
	UpdatedBy       string                 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelInfoFromPb(pb *registeredModelInfoPb) (*RegisteredModelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelInfo{}
	st.Aliases = pb.Aliases
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runRefreshRequestToPb(st *RunRefreshRequest) (*runRefreshRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runRefreshRequestPb{}
	pb.TableName = st.TableName

	return pb, nil
}

type runRefreshRequestPb struct {
	TableName string `json:"-" url:"-"`
}

func runRefreshRequestFromPb(pb *runRefreshRequestPb) (*RunRefreshRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunRefreshRequest{}
	st.TableName = pb.TableName

	return st, nil
}

func schemaInfoToPb(st *SchemaInfo) (*schemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schemaInfoPb{}
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.CatalogType = st.CatalogType
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.EffectivePredictiveOptimizationFlag = st.EffectivePredictiveOptimizationFlag
	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization
	pb.FullName = st.FullName
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.Properties = st.Properties
	pb.SchemaId = st.SchemaId
	pb.StorageLocation = st.StorageLocation
	pb.StorageRoot = st.StorageRoot
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type schemaInfoPb struct {
	BrowseOnly                          bool                                 `json:"browse_only,omitempty"`
	CatalogName                         string                               `json:"catalog_name,omitempty"`
	CatalogType                         CatalogType                          `json:"catalog_type,omitempty"`
	Comment                             string                               `json:"comment,omitempty"`
	CreatedAt                           int64                                `json:"created_at,omitempty"`
	CreatedBy                           string                               `json:"created_by,omitempty"`
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	EnablePredictiveOptimization        EnablePredictiveOptimization         `json:"enable_predictive_optimization,omitempty"`
	FullName                            string                               `json:"full_name,omitempty"`
	MetastoreId                         string                               `json:"metastore_id,omitempty"`
	Name                                string                               `json:"name,omitempty"`
	Owner                               string                               `json:"owner,omitempty"`
	Properties                          map[string]string                    `json:"properties,omitempty"`
	SchemaId                            string                               `json:"schema_id,omitempty"`
	StorageLocation                     string                               `json:"storage_location,omitempty"`
	StorageRoot                         string                               `json:"storage_root,omitempty"`
	UpdatedAt                           int64                                `json:"updated_at,omitempty"`
	UpdatedBy                           string                               `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func schemaInfoFromPb(pb *schemaInfoPb) (*SchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SchemaInfo{}
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.CatalogType = pb.CatalogType
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.EffectivePredictiveOptimizationFlag = pb.EffectivePredictiveOptimizationFlag
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.Properties = pb.Properties
	st.SchemaId = pb.SchemaId
	st.StorageLocation = pb.StorageLocation
	st.StorageRoot = pb.StorageRoot
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *schemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st schemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func setArtifactAllowlistToPb(st *SetArtifactAllowlist) (*setArtifactAllowlistPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setArtifactAllowlistPb{}
	pb.ArtifactMatchers = st.ArtifactMatchers
	pb.ArtifactType = st.ArtifactType
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.MetastoreId = st.MetastoreId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type setArtifactAllowlistPb struct {
	ArtifactMatchers []ArtifactMatcher `json:"artifact_matchers"`
	ArtifactType     ArtifactType      `json:"-" url:"-"`
	CreatedAt        int64             `json:"created_at,omitempty"`
	CreatedBy        string            `json:"created_by,omitempty"`
	MetastoreId      string            `json:"metastore_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setArtifactAllowlistFromPb(pb *setArtifactAllowlistPb) (*SetArtifactAllowlist, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetArtifactAllowlist{}
	st.ArtifactMatchers = pb.ArtifactMatchers
	st.ArtifactType = pb.ArtifactType
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.MetastoreId = pb.MetastoreId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setArtifactAllowlistPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setArtifactAllowlistPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func setRegisteredModelAliasRequestToPb(st *SetRegisteredModelAliasRequest) (*setRegisteredModelAliasRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setRegisteredModelAliasRequestPb{}
	pb.Alias = st.Alias
	pb.FullName = st.FullName
	pb.VersionNum = st.VersionNum

	return pb, nil
}

type setRegisteredModelAliasRequestPb struct {
	Alias      string `json:"alias" url:"-"`
	FullName   string `json:"full_name" url:"-"`
	VersionNum int    `json:"version_num"`
}

func setRegisteredModelAliasRequestFromPb(pb *setRegisteredModelAliasRequestPb) (*SetRegisteredModelAliasRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetRegisteredModelAliasRequest{}
	st.Alias = pb.Alias
	st.FullName = pb.FullName
	st.VersionNum = pb.VersionNum

	return st, nil
}

func sseEncryptionDetailsToPb(st *SseEncryptionDetails) (*sseEncryptionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sseEncryptionDetailsPb{}
	pb.Algorithm = st.Algorithm
	pb.AwsKmsKeyArn = st.AwsKmsKeyArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sseEncryptionDetailsPb struct {
	Algorithm    SseEncryptionDetailsAlgorithm `json:"algorithm,omitempty"`
	AwsKmsKeyArn string                        `json:"aws_kms_key_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sseEncryptionDetailsFromPb(pb *sseEncryptionDetailsPb) (*SseEncryptionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SseEncryptionDetails{}
	st.Algorithm = pb.Algorithm
	st.AwsKmsKeyArn = pb.AwsKmsKeyArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sseEncryptionDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sseEncryptionDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func storageCredentialInfoToPb(st *StorageCredentialInfo) (*storageCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &storageCredentialInfoPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.CloudflareApiToken = st.CloudflareApiToken
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.FullName = st.FullName
	pb.Id = st.Id
	pb.IsolationMode = st.IsolationMode
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.UsedForManagedStorage = st.UsedForManagedStorage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type storageCredentialInfoPb struct {
	AwsIamRole                  *AwsIamRoleResponse                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityResponse        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal               `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiToken                  `json:"cloudflare_api_token,omitempty"`
	Comment                     string                               `json:"comment,omitempty"`
	CreatedAt                   int64                                `json:"created_at,omitempty"`
	CreatedBy                   string                               `json:"created_by,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountResponse `json:"databricks_gcp_service_account,omitempty"`
	FullName                    string                               `json:"full_name,omitempty"`
	Id                          string                               `json:"id,omitempty"`
	IsolationMode               IsolationMode                        `json:"isolation_mode,omitempty"`
	MetastoreId                 string                               `json:"metastore_id,omitempty"`
	Name                        string                               `json:"name,omitempty"`
	Owner                       string                               `json:"owner,omitempty"`
	ReadOnly                    bool                                 `json:"read_only,omitempty"`
	UpdatedAt                   int64                                `json:"updated_at,omitempty"`
	UpdatedBy                   string                               `json:"updated_by,omitempty"`
	UsedForManagedStorage       bool                                 `json:"used_for_managed_storage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func storageCredentialInfoFromPb(pb *storageCredentialInfoPb) (*StorageCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StorageCredentialInfo{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.CloudflareApiToken = pb.CloudflareApiToken
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.FullName = pb.FullName
	st.Id = pb.Id
	st.IsolationMode = pb.IsolationMode
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.UsedForManagedStorage = pb.UsedForManagedStorage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *storageCredentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st storageCredentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func systemSchemaInfoToPb(st *SystemSchemaInfo) (*systemSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &systemSchemaInfoPb{}
	pb.Schema = st.Schema
	pb.State = st.State

	return pb, nil
}

type systemSchemaInfoPb struct {
	Schema string `json:"schema"`
	State  string `json:"state"`
}

func systemSchemaInfoFromPb(pb *systemSchemaInfoPb) (*SystemSchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SystemSchemaInfo{}
	st.Schema = pb.Schema
	st.State = pb.State

	return st, nil
}

func tableConstraintToPb(st *TableConstraint) (*tableConstraintPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableConstraintPb{}
	pb.ForeignKeyConstraint = st.ForeignKeyConstraint
	pb.NamedTableConstraint = st.NamedTableConstraint
	pb.PrimaryKeyConstraint = st.PrimaryKeyConstraint

	return pb, nil
}

type tableConstraintPb struct {
	ForeignKeyConstraint *ForeignKeyConstraint `json:"foreign_key_constraint,omitempty"`
	NamedTableConstraint *NamedTableConstraint `json:"named_table_constraint,omitempty"`
	PrimaryKeyConstraint *PrimaryKeyConstraint `json:"primary_key_constraint,omitempty"`
}

func tableConstraintFromPb(pb *tableConstraintPb) (*TableConstraint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableConstraint{}
	st.ForeignKeyConstraint = pb.ForeignKeyConstraint
	st.NamedTableConstraint = pb.NamedTableConstraint
	st.PrimaryKeyConstraint = pb.PrimaryKeyConstraint

	return st, nil
}

func tableDependencyToPb(st *TableDependency) (*tableDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableDependencyPb{}
	pb.TableFullName = st.TableFullName

	return pb, nil
}

type tableDependencyPb struct {
	TableFullName string `json:"table_full_name"`
}

func tableDependencyFromPb(pb *tableDependencyPb) (*TableDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableDependency{}
	st.TableFullName = pb.TableFullName

	return st, nil
}

func tableExistsResponseToPb(st *TableExistsResponse) (*tableExistsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableExistsResponsePb{}
	pb.TableExists = st.TableExists

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableExistsResponsePb struct {
	TableExists bool `json:"table_exists,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableExistsResponseFromPb(pb *tableExistsResponsePb) (*TableExistsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableExistsResponse{}
	st.TableExists = pb.TableExists

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableExistsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableExistsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tableInfoToPb(st *TableInfo) (*tableInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableInfoPb{}
	pb.AccessPoint = st.AccessPoint
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Columns = st.Columns
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.DataAccessConfigurationId = st.DataAccessConfigurationId
	pb.DataSourceFormat = st.DataSourceFormat
	pb.DeletedAt = st.DeletedAt
	pb.DeltaRuntimePropertiesKvpairs = st.DeltaRuntimePropertiesKvpairs
	pb.EffectivePredictiveOptimizationFlag = st.EffectivePredictiveOptimizationFlag
	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization
	pb.EncryptionDetails = st.EncryptionDetails
	pb.FullName = st.FullName
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.PipelineId = st.PipelineId
	pb.Properties = st.Properties
	pb.RowFilter = st.RowFilter
	pb.SchemaName = st.SchemaName
	pb.SqlPath = st.SqlPath
	pb.StorageCredentialName = st.StorageCredentialName
	pb.StorageLocation = st.StorageLocation
	pb.TableConstraints = st.TableConstraints
	pb.TableId = st.TableId
	pb.TableType = st.TableType
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.ViewDefinition = st.ViewDefinition
	pb.ViewDependencies = st.ViewDependencies

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableInfoPb struct {
	AccessPoint                         string                               `json:"access_point,omitempty"`
	BrowseOnly                          bool                                 `json:"browse_only,omitempty"`
	CatalogName                         string                               `json:"catalog_name,omitempty"`
	Columns                             []ColumnInfo                         `json:"columns,omitempty"`
	Comment                             string                               `json:"comment,omitempty"`
	CreatedAt                           int64                                `json:"created_at,omitempty"`
	CreatedBy                           string                               `json:"created_by,omitempty"`
	DataAccessConfigurationId           string                               `json:"data_access_configuration_id,omitempty"`
	DataSourceFormat                    DataSourceFormat                     `json:"data_source_format,omitempty"`
	DeletedAt                           int64                                `json:"deleted_at,omitempty"`
	DeltaRuntimePropertiesKvpairs       *DeltaRuntimePropertiesKvPairs       `json:"delta_runtime_properties_kvpairs,omitempty"`
	EffectivePredictiveOptimizationFlag *EffectivePredictiveOptimizationFlag `json:"effective_predictive_optimization_flag,omitempty"`
	EnablePredictiveOptimization        EnablePredictiveOptimization         `json:"enable_predictive_optimization,omitempty"`
	EncryptionDetails                   *EncryptionDetails                   `json:"encryption_details,omitempty"`
	FullName                            string                               `json:"full_name,omitempty"`
	MetastoreId                         string                               `json:"metastore_id,omitempty"`
	Name                                string                               `json:"name,omitempty"`
	Owner                               string                               `json:"owner,omitempty"`
	PipelineId                          string                               `json:"pipeline_id,omitempty"`
	Properties                          map[string]string                    `json:"properties,omitempty"`
	RowFilter                           *TableRowFilter                      `json:"row_filter,omitempty"`
	SchemaName                          string                               `json:"schema_name,omitempty"`
	SqlPath                             string                               `json:"sql_path,omitempty"`
	StorageCredentialName               string                               `json:"storage_credential_name,omitempty"`
	StorageLocation                     string                               `json:"storage_location,omitempty"`
	TableConstraints                    []TableConstraint                    `json:"table_constraints,omitempty"`
	TableId                             string                               `json:"table_id,omitempty"`
	TableType                           TableType                            `json:"table_type,omitempty"`
	UpdatedAt                           int64                                `json:"updated_at,omitempty"`
	UpdatedBy                           string                               `json:"updated_by,omitempty"`
	ViewDefinition                      string                               `json:"view_definition,omitempty"`
	ViewDependencies                    *DependencyList                      `json:"view_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableInfoFromPb(pb *tableInfoPb) (*TableInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableInfo{}
	st.AccessPoint = pb.AccessPoint
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Columns = pb.Columns
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.DataAccessConfigurationId = pb.DataAccessConfigurationId
	st.DataSourceFormat = pb.DataSourceFormat
	st.DeletedAt = pb.DeletedAt
	st.DeltaRuntimePropertiesKvpairs = pb.DeltaRuntimePropertiesKvpairs
	st.EffectivePredictiveOptimizationFlag = pb.EffectivePredictiveOptimizationFlag
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.EncryptionDetails = pb.EncryptionDetails
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.PipelineId = pb.PipelineId
	st.Properties = pb.Properties
	st.RowFilter = pb.RowFilter
	st.SchemaName = pb.SchemaName
	st.SqlPath = pb.SqlPath
	st.StorageCredentialName = pb.StorageCredentialName
	st.StorageLocation = pb.StorageLocation
	st.TableConstraints = pb.TableConstraints
	st.TableId = pb.TableId
	st.TableType = pb.TableType
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.ViewDefinition = pb.ViewDefinition
	st.ViewDependencies = pb.ViewDependencies

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tableRowFilterToPb(st *TableRowFilter) (*tableRowFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableRowFilterPb{}
	pb.FunctionName = st.FunctionName
	pb.InputColumnNames = st.InputColumnNames

	return pb, nil
}

type tableRowFilterPb struct {
	FunctionName     string   `json:"function_name"`
	InputColumnNames []string `json:"input_column_names"`
}

func tableRowFilterFromPb(pb *tableRowFilterPb) (*TableRowFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableRowFilter{}
	st.FunctionName = pb.FunctionName
	st.InputColumnNames = pb.InputColumnNames

	return st, nil
}

func tableSummaryToPb(st *TableSummary) (*tableSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSummaryPb{}
	pb.FullName = st.FullName
	pb.TableType = st.TableType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableSummaryPb struct {
	FullName  string    `json:"full_name,omitempty"`
	TableType TableType `json:"table_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableSummaryFromPb(pb *tableSummaryPb) (*TableSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSummary{}
	st.FullName = pb.FullName
	st.TableType = pb.TableType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableSummaryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableSummaryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tagKeyValueToPb(st *TagKeyValue) (*tagKeyValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tagKeyValuePb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tagKeyValuePb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tagKeyValueFromPb(pb *tagKeyValuePb) (*TagKeyValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TagKeyValue{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tagKeyValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tagKeyValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func temporaryCredentialsToPb(st *TemporaryCredentials) (*temporaryCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &temporaryCredentialsPb{}
	pb.AwsTempCredentials = st.AwsTempCredentials
	pb.AzureAad = st.AzureAad
	pb.ExpirationTime = st.ExpirationTime
	pb.GcpOauthToken = st.GcpOauthToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type temporaryCredentialsPb struct {
	AwsTempCredentials *AwsCredentials            `json:"aws_temp_credentials,omitempty"`
	AzureAad           *AzureActiveDirectoryToken `json:"azure_aad,omitempty"`
	ExpirationTime     int64                      `json:"expiration_time,omitempty"`
	GcpOauthToken      *GcpOauthToken             `json:"gcp_oauth_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func temporaryCredentialsFromPb(pb *temporaryCredentialsPb) (*TemporaryCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TemporaryCredentials{}
	st.AwsTempCredentials = pb.AwsTempCredentials
	st.AzureAad = pb.AzureAad
	st.ExpirationTime = pb.ExpirationTime
	st.GcpOauthToken = pb.GcpOauthToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *temporaryCredentialsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st temporaryCredentialsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func triggeredUpdateStatusToPb(st *TriggeredUpdateStatus) (*triggeredUpdateStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggeredUpdateStatusPb{}
	pb.LastProcessedCommitVersion = st.LastProcessedCommitVersion
	pb.Timestamp = st.Timestamp
	pb.TriggeredUpdateProgress = st.TriggeredUpdateProgress

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type triggeredUpdateStatusPb struct {
	LastProcessedCommitVersion int64             `json:"last_processed_commit_version,omitempty"`
	Timestamp                  string            `json:"timestamp,omitempty"`
	TriggeredUpdateProgress    *PipelineProgress `json:"triggered_update_progress,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func triggeredUpdateStatusFromPb(pb *triggeredUpdateStatusPb) (*TriggeredUpdateStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggeredUpdateStatus{}
	st.LastProcessedCommitVersion = pb.LastProcessedCommitVersion
	st.Timestamp = pb.Timestamp
	st.TriggeredUpdateProgress = pb.TriggeredUpdateProgress

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *triggeredUpdateStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st triggeredUpdateStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func unassignRequestToPb(st *UnassignRequest) (*unassignRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unassignRequestPb{}
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type unassignRequestPb struct {
	MetastoreId string `json:"-" url:"metastore_id"`
	WorkspaceId int64  `json:"-" url:"-"`
}

func unassignRequestFromPb(pb *unassignRequestPb) (*UnassignRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnassignRequest{}
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func unassignResponseToPb(st *UnassignResponse) (*unassignResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unassignResponsePb{}

	return pb, nil
}

type unassignResponsePb struct {
}

func unassignResponseFromPb(pb *unassignResponsePb) (*UnassignResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnassignResponse{}

	return st, nil
}

func updateAssignmentResponseToPb(st *UpdateAssignmentResponse) (*updateAssignmentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAssignmentResponsePb{}

	return pb, nil
}

type updateAssignmentResponsePb struct {
}

func updateAssignmentResponseFromPb(pb *updateAssignmentResponsePb) (*UpdateAssignmentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAssignmentResponse{}

	return st, nil
}

func updateCatalogToPb(st *UpdateCatalog) (*updateCatalogPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCatalogPb{}
	pb.Comment = st.Comment
	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization
	pb.IsolationMode = st.IsolationMode
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Options = st.Options
	pb.Owner = st.Owner
	pb.Properties = st.Properties

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateCatalogPb struct {
	Comment                      string                       `json:"comment,omitempty"`
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	IsolationMode                CatalogIsolationMode         `json:"isolation_mode,omitempty"`
	Name                         string                       `json:"-" url:"-"`
	NewName                      string                       `json:"new_name,omitempty"`
	Options                      map[string]string            `json:"options,omitempty"`
	Owner                        string                       `json:"owner,omitempty"`
	Properties                   map[string]string            `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateCatalogFromPb(pb *updateCatalogPb) (*UpdateCatalog, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCatalog{}
	st.Comment = pb.Comment
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.IsolationMode = pb.IsolationMode
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Options = pb.Options
	st.Owner = pb.Owner
	st.Properties = pb.Properties

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateCatalogPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateCatalogPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateCatalogWorkspaceBindingsResponseToPb(st *UpdateCatalogWorkspaceBindingsResponse) (*updateCatalogWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCatalogWorkspaceBindingsResponsePb{}
	pb.Workspaces = st.Workspaces

	return pb, nil
}

type updateCatalogWorkspaceBindingsResponsePb struct {
	Workspaces []int64 `json:"workspaces,omitempty"`
}

func updateCatalogWorkspaceBindingsResponseFromPb(pb *updateCatalogWorkspaceBindingsResponsePb) (*UpdateCatalogWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCatalogWorkspaceBindingsResponse{}
	st.Workspaces = pb.Workspaces

	return st, nil
}

func updateConnectionToPb(st *UpdateConnection) (*updateConnectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateConnectionPb{}
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Options = st.Options
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateConnectionPb struct {
	Name    string            `json:"-" url:"-"`
	NewName string            `json:"new_name,omitempty"`
	Options map[string]string `json:"options"`
	Owner   string            `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateConnectionFromPb(pb *updateConnectionPb) (*UpdateConnection, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateConnection{}
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Options = pb.Options
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateConnectionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateConnectionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateCredentialRequestToPb(st *UpdateCredentialRequest) (*updateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCredentialRequestPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.Comment = st.Comment
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.Force = st.Force
	pb.IsolationMode = st.IsolationMode
	pb.NameArg = st.NameArg
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateCredentialRequestPb struct {
	AwsIamRole                  *AwsIamRole                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentity        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal       `json:"azure_service_principal,omitempty"`
	Comment                     string                       `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	Force                       bool                         `json:"force,omitempty"`
	IsolationMode               IsolationMode                `json:"isolation_mode,omitempty"`
	NameArg                     string                       `json:"-" url:"-"`
	NewName                     string                       `json:"new_name,omitempty"`
	Owner                       string                       `json:"owner,omitempty"`
	ReadOnly                    bool                         `json:"read_only,omitempty"`
	SkipValidation              bool                         `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateCredentialRequestFromPb(pb *updateCredentialRequestPb) (*UpdateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCredentialRequest{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.Comment = pb.Comment
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.Force = pb.Force
	st.IsolationMode = pb.IsolationMode
	st.NameArg = pb.NameArg
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateExternalLocationToPb(st *UpdateExternalLocation) (*updateExternalLocationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExternalLocationPb{}
	pb.Comment = st.Comment
	pb.CredentialName = st.CredentialName
	pb.EnableFileEvents = st.EnableFileEvents
	pb.EncryptionDetails = st.EncryptionDetails
	pb.Fallback = st.Fallback
	pb.FileEventQueue = st.FileEventQueue
	pb.Force = st.Force
	pb.IsolationMode = st.IsolationMode
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateExternalLocationPb struct {
	Comment           string             `json:"comment,omitempty"`
	CredentialName    string             `json:"credential_name,omitempty"`
	EnableFileEvents  bool               `json:"enable_file_events,omitempty"`
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	Fallback          bool               `json:"fallback,omitempty"`
	FileEventQueue    *FileEventQueue    `json:"file_event_queue,omitempty"`
	Force             bool               `json:"force,omitempty"`
	IsolationMode     IsolationMode      `json:"isolation_mode,omitempty"`
	Name              string             `json:"-" url:"-"`
	NewName           string             `json:"new_name,omitempty"`
	Owner             string             `json:"owner,omitempty"`
	ReadOnly          bool               `json:"read_only,omitempty"`
	SkipValidation    bool               `json:"skip_validation,omitempty"`
	Url               string             `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateExternalLocationFromPb(pb *updateExternalLocationPb) (*UpdateExternalLocation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExternalLocation{}
	st.Comment = pb.Comment
	st.CredentialName = pb.CredentialName
	st.EnableFileEvents = pb.EnableFileEvents
	st.EncryptionDetails = pb.EncryptionDetails
	st.Fallback = pb.Fallback
	st.FileEventQueue = pb.FileEventQueue
	st.Force = pb.Force
	st.IsolationMode = pb.IsolationMode
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateExternalLocationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateExternalLocationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateFunctionToPb(st *UpdateFunction) (*updateFunctionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateFunctionPb{}
	pb.Name = st.Name
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateFunctionPb struct {
	Name  string `json:"-" url:"-"`
	Owner string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateFunctionFromPb(pb *updateFunctionPb) (*UpdateFunction, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateFunction{}
	st.Name = pb.Name
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateFunctionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateFunctionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateMetastoreToPb(st *UpdateMetastore) (*updateMetastorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMetastorePb{}
	pb.DeltaSharingOrganizationName = st.DeltaSharingOrganizationName
	pb.DeltaSharingRecipientTokenLifetimeInSeconds = st.DeltaSharingRecipientTokenLifetimeInSeconds
	pb.DeltaSharingScope = st.DeltaSharingScope
	pb.Id = st.Id
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.PrivilegeModelVersion = st.PrivilegeModelVersion
	pb.StorageRootCredentialId = st.StorageRootCredentialId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateMetastorePb struct {
	DeltaSharingOrganizationName                string                `json:"delta_sharing_organization_name,omitempty"`
	DeltaSharingRecipientTokenLifetimeInSeconds int64                 `json:"delta_sharing_recipient_token_lifetime_in_seconds,omitempty"`
	DeltaSharingScope                           DeltaSharingScopeEnum `json:"delta_sharing_scope,omitempty"`
	Id                                          string                `json:"-" url:"-"`
	NewName                                     string                `json:"new_name,omitempty"`
	Owner                                       string                `json:"owner,omitempty"`
	PrivilegeModelVersion                       string                `json:"privilege_model_version,omitempty"`
	StorageRootCredentialId                     string                `json:"storage_root_credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateMetastoreFromPb(pb *updateMetastorePb) (*UpdateMetastore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMetastore{}
	st.DeltaSharingOrganizationName = pb.DeltaSharingOrganizationName
	st.DeltaSharingRecipientTokenLifetimeInSeconds = pb.DeltaSharingRecipientTokenLifetimeInSeconds
	st.DeltaSharingScope = pb.DeltaSharingScope
	st.Id = pb.Id
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.PrivilegeModelVersion = pb.PrivilegeModelVersion
	st.StorageRootCredentialId = pb.StorageRootCredentialId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateMetastorePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateMetastorePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateMetastoreAssignmentToPb(st *UpdateMetastoreAssignment) (*updateMetastoreAssignmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMetastoreAssignmentPb{}
	pb.DefaultCatalogName = st.DefaultCatalogName
	pb.MetastoreId = st.MetastoreId
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateMetastoreAssignmentPb struct {
	DefaultCatalogName string `json:"default_catalog_name,omitempty"`
	MetastoreId        string `json:"metastore_id,omitempty"`
	WorkspaceId        int64  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateMetastoreAssignmentFromPb(pb *updateMetastoreAssignmentPb) (*UpdateMetastoreAssignment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMetastoreAssignment{}
	st.DefaultCatalogName = pb.DefaultCatalogName
	st.MetastoreId = pb.MetastoreId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateMetastoreAssignmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateMetastoreAssignmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*updateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionRequestPb{}
	pb.Comment = st.Comment
	pb.FullName = st.FullName
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateModelVersionRequestPb struct {
	Comment  string `json:"comment,omitempty"`
	FullName string `json:"-" url:"-"`
	Version  int    `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateModelVersionRequestFromPb(pb *updateModelVersionRequestPb) (*UpdateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionRequest{}
	st.Comment = pb.Comment
	st.FullName = pb.FullName
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateMonitorToPb(st *UpdateMonitor) (*updateMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateMonitorPb{}
	pb.BaselineTableName = st.BaselineTableName
	pb.CustomMetrics = st.CustomMetrics
	pb.DashboardId = st.DashboardId
	pb.DataClassificationConfig = st.DataClassificationConfig
	pb.InferenceLog = st.InferenceLog
	pb.Notifications = st.Notifications
	pb.OutputSchemaName = st.OutputSchemaName
	pb.Schedule = st.Schedule
	pb.SlicingExprs = st.SlicingExprs
	pb.Snapshot = st.Snapshot
	pb.TableName = st.TableName
	pb.TimeSeries = st.TimeSeries

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateMonitorPb struct {
	BaselineTableName        string                           `json:"baseline_table_name,omitempty"`
	CustomMetrics            []MonitorMetric                  `json:"custom_metrics,omitempty"`
	DashboardId              string                           `json:"dashboard_id,omitempty"`
	DataClassificationConfig *MonitorDataClassificationConfig `json:"data_classification_config,omitempty"`
	InferenceLog             *MonitorInferenceLog             `json:"inference_log,omitempty"`
	Notifications            *MonitorNotifications            `json:"notifications,omitempty"`
	OutputSchemaName         string                           `json:"output_schema_name"`
	Schedule                 *MonitorCronSchedule             `json:"schedule,omitempty"`
	SlicingExprs             []string                         `json:"slicing_exprs,omitempty"`
	Snapshot                 *MonitorSnapshot                 `json:"snapshot,omitempty"`
	TableName                string                           `json:"-" url:"-"`
	TimeSeries               *MonitorTimeSeries               `json:"time_series,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateMonitorFromPb(pb *updateMonitorPb) (*UpdateMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateMonitor{}
	st.BaselineTableName = pb.BaselineTableName
	st.CustomMetrics = pb.CustomMetrics
	st.DashboardId = pb.DashboardId
	st.DataClassificationConfig = pb.DataClassificationConfig
	st.InferenceLog = pb.InferenceLog
	st.Notifications = pb.Notifications
	st.OutputSchemaName = pb.OutputSchemaName
	st.Schedule = pb.Schedule
	st.SlicingExprs = pb.SlicingExprs
	st.Snapshot = pb.Snapshot
	st.TableName = pb.TableName
	st.TimeSeries = pb.TimeSeries

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateMonitorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateMonitorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updatePermissionsToPb(st *UpdatePermissions) (*updatePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePermissionsPb{}
	pb.Changes = st.Changes
	pb.FullName = st.FullName
	pb.SecurableType = st.SecurableType

	return pb, nil
}

type updatePermissionsPb struct {
	Changes       []PermissionsChange `json:"changes,omitempty"`
	FullName      string              `json:"-" url:"-"`
	SecurableType string              `json:"-" url:"-"`
}

func updatePermissionsFromPb(pb *updatePermissionsPb) (*UpdatePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePermissions{}
	st.Changes = pb.Changes
	st.FullName = pb.FullName
	st.SecurableType = pb.SecurableType

	return st, nil
}

func updatePermissionsResponseToPb(st *UpdatePermissionsResponse) (*updatePermissionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePermissionsResponsePb{}
	pb.PrivilegeAssignments = st.PrivilegeAssignments

	return pb, nil
}

type updatePermissionsResponsePb struct {
	PrivilegeAssignments []PrivilegeAssignment `json:"privilege_assignments,omitempty"`
}

func updatePermissionsResponseFromPb(pb *updatePermissionsResponsePb) (*UpdatePermissionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePermissionsResponse{}
	st.PrivilegeAssignments = pb.PrivilegeAssignments

	return st, nil
}

func updateRegisteredModelRequestToPb(st *UpdateRegisteredModelRequest) (*updateRegisteredModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRegisteredModelRequestPb{}
	pb.Comment = st.Comment
	pb.FullName = st.FullName
	pb.NewName = st.NewName
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateRegisteredModelRequestPb struct {
	Comment  string `json:"comment,omitempty"`
	FullName string `json:"-" url:"-"`
	NewName  string `json:"new_name,omitempty"`
	Owner    string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRegisteredModelRequestFromPb(pb *updateRegisteredModelRequestPb) (*UpdateRegisteredModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRegisteredModelRequest{}
	st.Comment = pb.Comment
	st.FullName = pb.FullName
	st.NewName = pb.NewName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRegisteredModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRegisteredModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
}

type updateResponsePb struct {
}

func updateResponseFromPb(pb *updateResponsePb) (*UpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResponse{}

	return st, nil
}

func updateSchemaToPb(st *UpdateSchema) (*updateSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSchemaPb{}
	pb.Comment = st.Comment
	pb.EnablePredictiveOptimization = st.EnablePredictiveOptimization
	pb.FullName = st.FullName
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.Properties = st.Properties

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateSchemaPb struct {
	Comment                      string                       `json:"comment,omitempty"`
	EnablePredictiveOptimization EnablePredictiveOptimization `json:"enable_predictive_optimization,omitempty"`
	FullName                     string                       `json:"-" url:"-"`
	NewName                      string                       `json:"new_name,omitempty"`
	Owner                        string                       `json:"owner,omitempty"`
	Properties                   map[string]string            `json:"properties,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateSchemaFromPb(pb *updateSchemaPb) (*UpdateSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSchema{}
	st.Comment = pb.Comment
	st.EnablePredictiveOptimization = pb.EnablePredictiveOptimization
	st.FullName = pb.FullName
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.Properties = pb.Properties

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateStorageCredentialToPb(st *UpdateStorageCredential) (*updateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateStorageCredentialPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.CloudflareApiToken = st.CloudflareApiToken
	pb.Comment = st.Comment
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.Force = st.Force
	pb.IsolationMode = st.IsolationMode
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner
	pb.ReadOnly = st.ReadOnly
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateStorageCredentialPb struct {
	AwsIamRole                  *AwsIamRoleRequest                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityResponse       `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal              `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiToken                 `json:"cloudflare_api_token,omitempty"`
	Comment                     string                              `json:"comment,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	Force                       bool                                `json:"force,omitempty"`
	IsolationMode               IsolationMode                       `json:"isolation_mode,omitempty"`
	Name                        string                              `json:"-" url:"-"`
	NewName                     string                              `json:"new_name,omitempty"`
	Owner                       string                              `json:"owner,omitempty"`
	ReadOnly                    bool                                `json:"read_only,omitempty"`
	SkipValidation              bool                                `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateStorageCredentialFromPb(pb *updateStorageCredentialPb) (*UpdateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStorageCredential{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.CloudflareApiToken = pb.CloudflareApiToken
	st.Comment = pb.Comment
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.Force = pb.Force
	st.IsolationMode = pb.IsolationMode
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner
	st.ReadOnly = pb.ReadOnly
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateTableRequestToPb(st *UpdateTableRequest) (*updateTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateTableRequestPb{}
	pb.FullName = st.FullName
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateTableRequestPb struct {
	FullName string `json:"-" url:"-"`
	Owner    string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateTableRequestFromPb(pb *updateTableRequestPb) (*UpdateTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateTableRequest{}
	st.FullName = pb.FullName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateTableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateTableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateVolumeRequestContentToPb(st *UpdateVolumeRequestContent) (*updateVolumeRequestContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVolumeRequestContentPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.NewName = st.NewName
	pb.Owner = st.Owner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateVolumeRequestContentPb struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"-" url:"-"`
	NewName string `json:"new_name,omitempty"`
	Owner   string `json:"owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateVolumeRequestContentFromPb(pb *updateVolumeRequestContentPb) (*UpdateVolumeRequestContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateVolumeRequestContent{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.NewName = pb.NewName
	st.Owner = pb.Owner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateVolumeRequestContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateVolumeRequestContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateWorkspaceBindingsToPb(st *UpdateWorkspaceBindings) (*updateWorkspaceBindingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceBindingsPb{}
	pb.AssignWorkspaces = st.AssignWorkspaces
	pb.Name = st.Name
	pb.UnassignWorkspaces = st.UnassignWorkspaces

	return pb, nil
}

type updateWorkspaceBindingsPb struct {
	AssignWorkspaces   []int64 `json:"assign_workspaces,omitempty"`
	Name               string  `json:"-" url:"-"`
	UnassignWorkspaces []int64 `json:"unassign_workspaces,omitempty"`
}

func updateWorkspaceBindingsFromPb(pb *updateWorkspaceBindingsPb) (*UpdateWorkspaceBindings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindings{}
	st.AssignWorkspaces = pb.AssignWorkspaces
	st.Name = pb.Name
	st.UnassignWorkspaces = pb.UnassignWorkspaces

	return st, nil
}

func updateWorkspaceBindingsParametersToPb(st *UpdateWorkspaceBindingsParameters) (*updateWorkspaceBindingsParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceBindingsParametersPb{}
	pb.Add = st.Add
	pb.Remove = st.Remove
	pb.SecurableName = st.SecurableName
	pb.SecurableType = st.SecurableType

	return pb, nil
}

type updateWorkspaceBindingsParametersPb struct {
	Add           []WorkspaceBinding `json:"add,omitempty"`
	Remove        []WorkspaceBinding `json:"remove,omitempty"`
	SecurableName string             `json:"-" url:"-"`
	SecurableType string             `json:"-" url:"-"`
}

func updateWorkspaceBindingsParametersFromPb(pb *updateWorkspaceBindingsParametersPb) (*UpdateWorkspaceBindingsParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindingsParameters{}
	st.Add = pb.Add
	st.Remove = pb.Remove
	st.SecurableName = pb.SecurableName
	st.SecurableType = pb.SecurableType

	return st, nil
}

func updateWorkspaceBindingsResponseToPb(st *UpdateWorkspaceBindingsResponse) (*updateWorkspaceBindingsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceBindingsResponsePb{}
	pb.Bindings = st.Bindings

	return pb, nil
}

type updateWorkspaceBindingsResponsePb struct {
	Bindings []WorkspaceBinding `json:"bindings,omitempty"`
}

func updateWorkspaceBindingsResponseFromPb(pb *updateWorkspaceBindingsResponsePb) (*UpdateWorkspaceBindingsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceBindingsResponse{}
	st.Bindings = pb.Bindings

	return st, nil
}

func validateCredentialRequestToPb(st *ValidateCredentialRequest) (*validateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateCredentialRequestPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.CredentialName = st.CredentialName
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.ExternalLocationName = st.ExternalLocationName
	pb.Purpose = st.Purpose
	pb.ReadOnly = st.ReadOnly
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type validateCredentialRequestPb struct {
	AwsIamRole                  *AwsIamRole                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentity        `json:"azure_managed_identity,omitempty"`
	CredentialName              string                       `json:"credential_name,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccount `json:"databricks_gcp_service_account,omitempty"`
	ExternalLocationName        string                       `json:"external_location_name,omitempty"`
	Purpose                     CredentialPurpose            `json:"purpose,omitempty"`
	ReadOnly                    bool                         `json:"read_only,omitempty"`
	Url                         string                       `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateCredentialRequestFromPb(pb *validateCredentialRequestPb) (*ValidateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateCredentialRequest{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.CredentialName = pb.CredentialName
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.ExternalLocationName = pb.ExternalLocationName
	st.Purpose = pb.Purpose
	st.ReadOnly = pb.ReadOnly
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateCredentialRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateCredentialRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func validateCredentialResponseToPb(st *ValidateCredentialResponse) (*validateCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateCredentialResponsePb{}
	pb.IsDir = st.IsDir
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type validateCredentialResponsePb struct {
	IsDir   bool                         `json:"isDir,omitempty"`
	Results []CredentialValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateCredentialResponseFromPb(pb *validateCredentialResponsePb) (*ValidateCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateCredentialResponse{}
	st.IsDir = pb.IsDir
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func validateStorageCredentialToPb(st *ValidateStorageCredential) (*validateStorageCredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateStorageCredentialPb{}
	pb.AwsIamRole = st.AwsIamRole
	pb.AzureManagedIdentity = st.AzureManagedIdentity
	pb.AzureServicePrincipal = st.AzureServicePrincipal
	pb.CloudflareApiToken = st.CloudflareApiToken
	pb.DatabricksGcpServiceAccount = st.DatabricksGcpServiceAccount
	pb.ExternalLocationName = st.ExternalLocationName
	pb.ReadOnly = st.ReadOnly
	pb.StorageCredentialName = st.StorageCredentialName
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type validateStorageCredentialPb struct {
	AwsIamRole                  *AwsIamRoleRequest                  `json:"aws_iam_role,omitempty"`
	AzureManagedIdentity        *AzureManagedIdentityRequest        `json:"azure_managed_identity,omitempty"`
	AzureServicePrincipal       *AzureServicePrincipal              `json:"azure_service_principal,omitempty"`
	CloudflareApiToken          *CloudflareApiToken                 `json:"cloudflare_api_token,omitempty"`
	DatabricksGcpServiceAccount *DatabricksGcpServiceAccountRequest `json:"databricks_gcp_service_account,omitempty"`
	ExternalLocationName        string                              `json:"external_location_name,omitempty"`
	ReadOnly                    bool                                `json:"read_only,omitempty"`
	StorageCredentialName       string                              `json:"storage_credential_name,omitempty"`
	Url                         string                              `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateStorageCredentialFromPb(pb *validateStorageCredentialPb) (*ValidateStorageCredential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateStorageCredential{}
	st.AwsIamRole = pb.AwsIamRole
	st.AzureManagedIdentity = pb.AzureManagedIdentity
	st.AzureServicePrincipal = pb.AzureServicePrincipal
	st.CloudflareApiToken = pb.CloudflareApiToken
	st.DatabricksGcpServiceAccount = pb.DatabricksGcpServiceAccount
	st.ExternalLocationName = pb.ExternalLocationName
	st.ReadOnly = pb.ReadOnly
	st.StorageCredentialName = pb.StorageCredentialName
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateStorageCredentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateStorageCredentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func validateStorageCredentialResponseToPb(st *ValidateStorageCredentialResponse) (*validateStorageCredentialResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validateStorageCredentialResponsePb{}
	pb.IsDir = st.IsDir
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type validateStorageCredentialResponsePb struct {
	IsDir   bool               `json:"isDir,omitempty"`
	Results []ValidationResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validateStorageCredentialResponseFromPb(pb *validateStorageCredentialResponsePb) (*ValidateStorageCredentialResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidateStorageCredentialResponse{}
	st.IsDir = pb.IsDir
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validateStorageCredentialResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validateStorageCredentialResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func validationResultToPb(st *ValidationResult) (*validationResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &validationResultPb{}
	pb.Message = st.Message
	pb.Operation = st.Operation
	pb.Result = st.Result

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type validationResultPb struct {
	Message   string                    `json:"message,omitempty"`
	Operation ValidationResultOperation `json:"operation,omitempty"`
	Result    ValidationResultResult    `json:"result,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func validationResultFromPb(pb *validationResultPb) (*ValidationResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ValidationResult{}
	st.Message = pb.Message
	st.Operation = pb.Operation
	st.Result = pb.Result

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *validationResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st validationResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func volumeInfoToPb(st *VolumeInfo) (*volumeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumeInfoPb{}
	pb.AccessPoint = st.AccessPoint
	pb.BrowseOnly = st.BrowseOnly
	pb.CatalogName = st.CatalogName
	pb.Comment = st.Comment
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.EncryptionDetails = st.EncryptionDetails
	pb.FullName = st.FullName
	pb.MetastoreId = st.MetastoreId
	pb.Name = st.Name
	pb.Owner = st.Owner
	pb.SchemaName = st.SchemaName
	pb.StorageLocation = st.StorageLocation
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy
	pb.VolumeId = st.VolumeId
	pb.VolumeType = st.VolumeType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type volumeInfoPb struct {
	AccessPoint       string             `json:"access_point,omitempty"`
	BrowseOnly        bool               `json:"browse_only,omitempty"`
	CatalogName       string             `json:"catalog_name,omitempty"`
	Comment           string             `json:"comment,omitempty"`
	CreatedAt         int64              `json:"created_at,omitempty"`
	CreatedBy         string             `json:"created_by,omitempty"`
	EncryptionDetails *EncryptionDetails `json:"encryption_details,omitempty"`
	FullName          string             `json:"full_name,omitempty"`
	MetastoreId       string             `json:"metastore_id,omitempty"`
	Name              string             `json:"name,omitempty"`
	Owner             string             `json:"owner,omitempty"`
	SchemaName        string             `json:"schema_name,omitempty"`
	StorageLocation   string             `json:"storage_location,omitempty"`
	UpdatedAt         int64              `json:"updated_at,omitempty"`
	UpdatedBy         string             `json:"updated_by,omitempty"`
	VolumeId          string             `json:"volume_id,omitempty"`
	VolumeType        VolumeType         `json:"volume_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func volumeInfoFromPb(pb *volumeInfoPb) (*VolumeInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumeInfo{}
	st.AccessPoint = pb.AccessPoint
	st.BrowseOnly = pb.BrowseOnly
	st.CatalogName = pb.CatalogName
	st.Comment = pb.Comment
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.EncryptionDetails = pb.EncryptionDetails
	st.FullName = pb.FullName
	st.MetastoreId = pb.MetastoreId
	st.Name = pb.Name
	st.Owner = pb.Owner
	st.SchemaName = pb.SchemaName
	st.StorageLocation = pb.StorageLocation
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy
	st.VolumeId = pb.VolumeId
	st.VolumeType = pb.VolumeType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *volumeInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st volumeInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func workspaceBindingToPb(st *WorkspaceBinding) (*workspaceBindingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceBindingPb{}
	pb.BindingType = st.BindingType
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type workspaceBindingPb struct {
	BindingType WorkspaceBindingBindingType `json:"binding_type,omitempty"`
	WorkspaceId int64                       `json:"workspace_id"`
}

func workspaceBindingFromPb(pb *workspaceBindingPb) (*WorkspaceBinding, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceBinding{}
	st.BindingType = pb.BindingType
	st.WorkspaceId = pb.WorkspaceId

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
