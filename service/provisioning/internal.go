// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

func awsCredentialsToPb(st *AwsCredentials) (*awsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsCredentialsPb{}
	pb.StsRole = st.StsRole

	return pb, nil
}

type awsCredentialsPb struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

func awsCredentialsFromPb(pb *awsCredentialsPb) (*AwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsCredentials{}
	st.StsRole = pb.StsRole

	return st, nil
}

func awsKeyInfoToPb(st *AwsKeyInfo) (*awsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsKeyInfoPb{}
	pb.KeyAlias = st.KeyAlias

	pb.KeyArn = st.KeyArn

	pb.KeyRegion = st.KeyRegion

	pb.ReuseKeyForClusterVolumes = st.ReuseKeyForClusterVolumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type awsKeyInfoPb struct {
	KeyAlias string `json:"key_alias,omitempty"`

	KeyArn string `json:"key_arn"`

	KeyRegion string `json:"key_region"`

	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsKeyInfoFromPb(pb *awsKeyInfoPb) (*AwsKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsKeyInfo{}
	st.KeyAlias = pb.KeyAlias
	st.KeyArn = pb.KeyArn
	st.KeyRegion = pb.KeyRegion
	st.ReuseKeyForClusterVolumes = pb.ReuseKeyForClusterVolumes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsKeyInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsKeyInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureWorkspaceInfoToPb(st *AzureWorkspaceInfo) (*azureWorkspaceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureWorkspaceInfoPb{}
	pb.ResourceGroup = st.ResourceGroup

	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureWorkspaceInfoPb struct {
	ResourceGroup string `json:"resource_group,omitempty"`

	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureWorkspaceInfoFromPb(pb *azureWorkspaceInfoPb) (*AzureWorkspaceInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureWorkspaceInfo{}
	st.ResourceGroup = pb.ResourceGroup
	st.SubscriptionId = pb.SubscriptionId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureWorkspaceInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureWorkspaceInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cloudResourceContainerToPb(st *CloudResourceContainer) (*cloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudResourceContainerPb{}
	pb.Gcp = st.Gcp

	return pb, nil
}

type cloudResourceContainerPb struct {
	Gcp *CustomerFacingGcpCloudResourceContainer `json:"gcp,omitempty"`
}

func cloudResourceContainerFromPb(pb *cloudResourceContainerPb) (*CloudResourceContainer, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudResourceContainer{}
	st.Gcp = pb.Gcp

	return st, nil
}

func createAwsKeyInfoToPb(st *CreateAwsKeyInfo) (*createAwsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAwsKeyInfoPb{}
	pb.KeyAlias = st.KeyAlias

	pb.KeyArn = st.KeyArn

	pb.ReuseKeyForClusterVolumes = st.ReuseKeyForClusterVolumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createAwsKeyInfoPb struct {
	KeyAlias string `json:"key_alias,omitempty"`

	KeyArn string `json:"key_arn"`

	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAwsKeyInfoFromPb(pb *createAwsKeyInfoPb) (*CreateAwsKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAwsKeyInfo{}
	st.KeyAlias = pb.KeyAlias
	st.KeyArn = pb.KeyArn
	st.ReuseKeyForClusterVolumes = pb.ReuseKeyForClusterVolumes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAwsKeyInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAwsKeyInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCredentialAwsCredentialsToPb(st *CreateCredentialAwsCredentials) (*createCredentialAwsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialAwsCredentialsPb{}
	pb.StsRole = st.StsRole

	return pb, nil
}

type createCredentialAwsCredentialsPb struct {
	StsRole *CreateCredentialStsRole `json:"sts_role,omitempty"`
}

func createCredentialAwsCredentialsFromPb(pb *createCredentialAwsCredentialsPb) (*CreateCredentialAwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialAwsCredentials{}
	st.StsRole = pb.StsRole

	return st, nil
}

func createCredentialRequestToPb(st *CreateCredentialRequest) (*createCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialRequestPb{}
	pb.AwsCredentials = st.AwsCredentials

	pb.CredentialsName = st.CredentialsName

	return pb, nil
}

type createCredentialRequestPb struct {
	AwsCredentials CreateCredentialAwsCredentials `json:"aws_credentials"`

	CredentialsName string `json:"credentials_name"`
}

func createCredentialRequestFromPb(pb *createCredentialRequestPb) (*CreateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialRequest{}
	st.AwsCredentials = pb.AwsCredentials
	st.CredentialsName = pb.CredentialsName

	return st, nil
}

func createCredentialStsRoleToPb(st *CreateCredentialStsRole) (*createCredentialStsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCredentialStsRolePb{}
	pb.RoleArn = st.RoleArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createCredentialStsRolePb struct {
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCredentialStsRoleFromPb(pb *createCredentialStsRolePb) (*CreateCredentialStsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialStsRole{}
	st.RoleArn = pb.RoleArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCredentialStsRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCredentialStsRolePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCustomerManagedKeyRequestToPb(st *CreateCustomerManagedKeyRequest) (*createCustomerManagedKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomerManagedKeyRequestPb{}
	pb.AwsKeyInfo = st.AwsKeyInfo

	pb.GcpKeyInfo = st.GcpKeyInfo

	pb.UseCases = st.UseCases

	return pb, nil
}

type createCustomerManagedKeyRequestPb struct {
	AwsKeyInfo *CreateAwsKeyInfo `json:"aws_key_info,omitempty"`

	GcpKeyInfo *CreateGcpKeyInfo `json:"gcp_key_info,omitempty"`

	UseCases []KeyUseCase `json:"use_cases"`
}

func createCustomerManagedKeyRequestFromPb(pb *createCustomerManagedKeyRequestPb) (*CreateCustomerManagedKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomerManagedKeyRequest{}
	st.AwsKeyInfo = pb.AwsKeyInfo
	st.GcpKeyInfo = pb.GcpKeyInfo
	st.UseCases = pb.UseCases

	return st, nil
}

func createGcpKeyInfoToPb(st *CreateGcpKeyInfo) (*createGcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createGcpKeyInfoPb{}
	pb.KmsKeyId = st.KmsKeyId

	return pb, nil
}

type createGcpKeyInfoPb struct {
	KmsKeyId string `json:"kms_key_id"`
}

func createGcpKeyInfoFromPb(pb *createGcpKeyInfoPb) (*CreateGcpKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateGcpKeyInfo{}
	st.KmsKeyId = pb.KmsKeyId

	return st, nil
}

func createNetworkRequestToPb(st *CreateNetworkRequest) (*createNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkRequestPb{}
	pb.GcpNetworkInfo = st.GcpNetworkInfo

	pb.NetworkName = st.NetworkName

	pb.SecurityGroupIds = st.SecurityGroupIds

	pb.SubnetIds = st.SubnetIds

	pb.VpcEndpoints = st.VpcEndpoints

	pb.VpcId = st.VpcId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createNetworkRequestPb struct {
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`

	NetworkName string `json:"network_name"`

	SecurityGroupIds []string `json:"security_group_ids,omitempty"`

	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`

	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createNetworkRequestFromPb(pb *createNetworkRequestPb) (*CreateNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkRequest{}
	st.GcpNetworkInfo = pb.GcpNetworkInfo
	st.NetworkName = pb.NetworkName
	st.SecurityGroupIds = pb.SecurityGroupIds
	st.SubnetIds = pb.SubnetIds
	st.VpcEndpoints = pb.VpcEndpoints
	st.VpcId = pb.VpcId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createNetworkRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createNetworkRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createStorageConfigurationRequestToPb(st *CreateStorageConfigurationRequest) (*createStorageConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createStorageConfigurationRequestPb{}
	pb.RootBucketInfo = st.RootBucketInfo

	pb.StorageConfigurationName = st.StorageConfigurationName

	return pb, nil
}

type createStorageConfigurationRequestPb struct {
	RootBucketInfo RootBucketInfo `json:"root_bucket_info"`

	StorageConfigurationName string `json:"storage_configuration_name"`
}

func createStorageConfigurationRequestFromPb(pb *createStorageConfigurationRequestPb) (*CreateStorageConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateStorageConfigurationRequest{}
	st.RootBucketInfo = pb.RootBucketInfo
	st.StorageConfigurationName = pb.StorageConfigurationName

	return st, nil
}

func createVpcEndpointRequestToPb(st *CreateVpcEndpointRequest) (*createVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVpcEndpointRequestPb{}
	pb.AwsVpcEndpointId = st.AwsVpcEndpointId

	pb.GcpVpcEndpointInfo = st.GcpVpcEndpointInfo

	pb.Region = st.Region

	pb.VpcEndpointName = st.VpcEndpointName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createVpcEndpointRequestPb struct {
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`

	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`

	Region string `json:"region,omitempty"`

	VpcEndpointName string `json:"vpc_endpoint_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createVpcEndpointRequestFromPb(pb *createVpcEndpointRequestPb) (*CreateVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVpcEndpointRequest{}
	st.AwsVpcEndpointId = pb.AwsVpcEndpointId
	st.GcpVpcEndpointInfo = pb.GcpVpcEndpointInfo
	st.Region = pb.Region
	st.VpcEndpointName = pb.VpcEndpointName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createVpcEndpointRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createVpcEndpointRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createWorkspaceRequestToPb(st *CreateWorkspaceRequest) (*createWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWorkspaceRequestPb{}
	pb.AwsRegion = st.AwsRegion

	pb.Cloud = st.Cloud

	pb.CloudResourceContainer = st.CloudResourceContainer

	pb.CredentialsId = st.CredentialsId

	pb.CustomTags = st.CustomTags

	pb.DeploymentName = st.DeploymentName

	pb.GcpManagedNetworkConfig = st.GcpManagedNetworkConfig

	pb.GkeConfig = st.GkeConfig

	pb.IsNoPublicIpEnabled = st.IsNoPublicIpEnabled

	pb.Location = st.Location

	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId

	pb.NetworkId = st.NetworkId

	pb.PricingTier = st.PricingTier

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId

	pb.WorkspaceName = st.WorkspaceName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createWorkspaceRequestPb struct {
	AwsRegion string `json:"aws_region,omitempty"`

	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`

	CredentialsId string `json:"credentials_id,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DeploymentName string `json:"deployment_name,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`

	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`

	Location string `json:"location,omitempty"`

	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`

	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`

	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`

	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`

	WorkspaceName string `json:"workspace_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWorkspaceRequestFromPb(pb *createWorkspaceRequestPb) (*CreateWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWorkspaceRequest{}
	st.AwsRegion = pb.AwsRegion
	st.Cloud = pb.Cloud
	st.CloudResourceContainer = pb.CloudResourceContainer
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.DeploymentName = pb.DeploymentName
	st.GcpManagedNetworkConfig = pb.GcpManagedNetworkConfig
	st.GkeConfig = pb.GkeConfig
	st.IsNoPublicIpEnabled = pb.IsNoPublicIpEnabled
	st.Location = pb.Location
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkId = pb.NetworkId
	st.PricingTier = pb.PricingTier
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceName = pb.WorkspaceName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func credentialToPb(st *Credential) (*credentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &credentialPb{}
	pb.AccountId = st.AccountId

	pb.AwsCredentials = st.AwsCredentials

	pb.CreationTime = st.CreationTime

	pb.CredentialsId = st.CredentialsId

	pb.CredentialsName = st.CredentialsName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type credentialPb struct {
	AccountId string `json:"account_id,omitempty"`

	AwsCredentials *AwsCredentials `json:"aws_credentials,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`

	CredentialsId string `json:"credentials_id,omitempty"`

	CredentialsName string `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func credentialFromPb(pb *credentialPb) (*Credential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Credential{}
	st.AccountId = pb.AccountId
	st.AwsCredentials = pb.AwsCredentials
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.CredentialsName = pb.CredentialsName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *credentialPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st credentialPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type customTagsPb CustomTags

func customTagsToPb(st *CustomTags) (*customTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := customTagsPb(*st)
	return &stPb, nil
}

func customTagsFromPb(stPb *customTagsPb) (*CustomTags, error) {
	if stPb == nil {
		return nil, nil
	}
	st := CustomTags(*stPb)
	return &st, nil
}

func customerFacingGcpCloudResourceContainerToPb(st *CustomerFacingGcpCloudResourceContainer) (*customerFacingGcpCloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerFacingGcpCloudResourceContainerPb{}
	pb.ProjectId = st.ProjectId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type customerFacingGcpCloudResourceContainerPb struct {
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customerFacingGcpCloudResourceContainerFromPb(pb *customerFacingGcpCloudResourceContainerPb) (*CustomerFacingGcpCloudResourceContainer, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerFacingGcpCloudResourceContainer{}
	st.ProjectId = pb.ProjectId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customerFacingGcpCloudResourceContainerPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customerFacingGcpCloudResourceContainerPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func customerManagedKeyToPb(st *CustomerManagedKey) (*customerManagedKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerManagedKeyPb{}
	pb.AccountId = st.AccountId

	pb.AwsKeyInfo = st.AwsKeyInfo

	pb.CreationTime = st.CreationTime

	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	pb.GcpKeyInfo = st.GcpKeyInfo

	pb.UseCases = st.UseCases

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type customerManagedKeyPb struct {
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`

	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	GcpKeyInfo *GcpKeyInfo `json:"gcp_key_info,omitempty"`

	UseCases []KeyUseCase `json:"use_cases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customerManagedKeyFromPb(pb *customerManagedKeyPb) (*CustomerManagedKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerManagedKey{}
	st.AccountId = pb.AccountId
	st.AwsKeyInfo = pb.AwsKeyInfo
	st.CreationTime = pb.CreationTime
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId
	st.GcpKeyInfo = pb.GcpKeyInfo
	st.UseCases = pb.UseCases

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customerManagedKeyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customerManagedKeyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteCredentialRequestToPb(st *DeleteCredentialRequest) (*deleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCredentialRequestPb{}
	pb.CredentialsId = st.CredentialsId

	return pb, nil
}

type deleteCredentialRequestPb struct {
	CredentialsId string `json:"-" url:"-"`
}

func deleteCredentialRequestFromPb(pb *deleteCredentialRequestPb) (*DeleteCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialRequest{}
	st.CredentialsId = pb.CredentialsId

	return st, nil
}

func deleteEncryptionKeyRequestToPb(st *DeleteEncryptionKeyRequest) (*deleteEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteEncryptionKeyRequestPb{}
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	return pb, nil
}

type deleteEncryptionKeyRequestPb struct {
	CustomerManagedKeyId string `json:"-" url:"-"`
}

func deleteEncryptionKeyRequestFromPb(pb *deleteEncryptionKeyRequestPb) (*DeleteEncryptionKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteEncryptionKeyRequest{}
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId

	return st, nil
}

func deleteNetworkRequestToPb(st *DeleteNetworkRequest) (*deleteNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkRequestPb{}
	pb.NetworkId = st.NetworkId

	return pb, nil
}

type deleteNetworkRequestPb struct {
	NetworkId string `json:"-" url:"-"`
}

func deleteNetworkRequestFromPb(pb *deleteNetworkRequestPb) (*DeleteNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkRequest{}
	st.NetworkId = pb.NetworkId

	return st, nil
}

func deletePrivateAccesRequestToPb(st *DeletePrivateAccesRequest) (*deletePrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePrivateAccesRequestPb{}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	return pb, nil
}

type deletePrivateAccesRequestPb struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func deletePrivateAccesRequestFromPb(pb *deletePrivateAccesRequestPb) (*DeletePrivateAccesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePrivateAccesRequest{}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId

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

func deleteStorageRequestToPb(st *DeleteStorageRequest) (*deleteStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteStorageRequestPb{}
	pb.StorageConfigurationId = st.StorageConfigurationId

	return pb, nil
}

type deleteStorageRequestPb struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

func deleteStorageRequestFromPb(pb *deleteStorageRequestPb) (*DeleteStorageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteStorageRequest{}
	st.StorageConfigurationId = pb.StorageConfigurationId

	return st, nil
}

func deleteVpcEndpointRequestToPb(st *DeleteVpcEndpointRequest) (*deleteVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVpcEndpointRequestPb{}
	pb.VpcEndpointId = st.VpcEndpointId

	return pb, nil
}

type deleteVpcEndpointRequestPb struct {
	VpcEndpointId string `json:"-" url:"-"`
}

func deleteVpcEndpointRequestFromPb(pb *deleteVpcEndpointRequestPb) (*DeleteVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVpcEndpointRequest{}
	st.VpcEndpointId = pb.VpcEndpointId

	return st, nil
}

func deleteWorkspaceRequestToPb(st *DeleteWorkspaceRequest) (*deleteWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWorkspaceRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type deleteWorkspaceRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

func deleteWorkspaceRequestFromPb(pb *deleteWorkspaceRequestPb) (*DeleteWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWorkspaceRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func externalCustomerInfoToPb(st *ExternalCustomerInfo) (*externalCustomerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalCustomerInfoPb{}
	pb.AuthoritativeUserEmail = st.AuthoritativeUserEmail

	pb.AuthoritativeUserFullName = st.AuthoritativeUserFullName

	pb.CustomerName = st.CustomerName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type externalCustomerInfoPb struct {
	AuthoritativeUserEmail string `json:"authoritative_user_email,omitempty"`

	AuthoritativeUserFullName string `json:"authoritative_user_full_name,omitempty"`

	CustomerName string `json:"customer_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalCustomerInfoFromPb(pb *externalCustomerInfoPb) (*ExternalCustomerInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalCustomerInfo{}
	st.AuthoritativeUserEmail = pb.AuthoritativeUserEmail
	st.AuthoritativeUserFullName = pb.AuthoritativeUserFullName
	st.CustomerName = pb.CustomerName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalCustomerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalCustomerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gcpKeyInfoToPb(st *GcpKeyInfo) (*gcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpKeyInfoPb{}
	pb.KmsKeyId = st.KmsKeyId

	return pb, nil
}

type gcpKeyInfoPb struct {
	KmsKeyId string `json:"kms_key_id"`
}

func gcpKeyInfoFromPb(pb *gcpKeyInfoPb) (*GcpKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpKeyInfo{}
	st.KmsKeyId = pb.KmsKeyId

	return st, nil
}

func gcpManagedNetworkConfigToPb(st *GcpManagedNetworkConfig) (*gcpManagedNetworkConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpManagedNetworkConfigPb{}
	pb.GkeClusterPodIpRange = st.GkeClusterPodIpRange

	pb.GkeClusterServiceIpRange = st.GkeClusterServiceIpRange

	pb.SubnetCidr = st.SubnetCidr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gcpManagedNetworkConfigPb struct {
	GkeClusterPodIpRange string `json:"gke_cluster_pod_ip_range,omitempty"`

	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`

	SubnetCidr string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpManagedNetworkConfigFromPb(pb *gcpManagedNetworkConfigPb) (*GcpManagedNetworkConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpManagedNetworkConfig{}
	st.GkeClusterPodIpRange = pb.GkeClusterPodIpRange
	st.GkeClusterServiceIpRange = pb.GkeClusterServiceIpRange
	st.SubnetCidr = pb.SubnetCidr

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpManagedNetworkConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpManagedNetworkConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gcpNetworkInfoToPb(st *GcpNetworkInfo) (*gcpNetworkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpNetworkInfoPb{}
	pb.NetworkProjectId = st.NetworkProjectId

	pb.PodIpRangeName = st.PodIpRangeName

	pb.ServiceIpRangeName = st.ServiceIpRangeName

	pb.SubnetId = st.SubnetId

	pb.SubnetRegion = st.SubnetRegion

	pb.VpcId = st.VpcId

	return pb, nil
}

type gcpNetworkInfoPb struct {
	NetworkProjectId string `json:"network_project_id"`

	PodIpRangeName string `json:"pod_ip_range_name"`

	ServiceIpRangeName string `json:"service_ip_range_name"`

	SubnetId string `json:"subnet_id"`

	SubnetRegion string `json:"subnet_region"`

	VpcId string `json:"vpc_id"`
}

func gcpNetworkInfoFromPb(pb *gcpNetworkInfoPb) (*GcpNetworkInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpNetworkInfo{}
	st.NetworkProjectId = pb.NetworkProjectId
	st.PodIpRangeName = pb.PodIpRangeName
	st.ServiceIpRangeName = pb.ServiceIpRangeName
	st.SubnetId = pb.SubnetId
	st.SubnetRegion = pb.SubnetRegion
	st.VpcId = pb.VpcId

	return st, nil
}

func gcpVpcEndpointInfoToPb(st *GcpVpcEndpointInfo) (*gcpVpcEndpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpVpcEndpointInfoPb{}
	pb.EndpointRegion = st.EndpointRegion

	pb.ProjectId = st.ProjectId

	pb.PscConnectionId = st.PscConnectionId

	pb.PscEndpointName = st.PscEndpointName

	pb.ServiceAttachmentId = st.ServiceAttachmentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gcpVpcEndpointInfoPb struct {
	EndpointRegion string `json:"endpoint_region"`

	ProjectId string `json:"project_id"`

	PscConnectionId string `json:"psc_connection_id,omitempty"`

	PscEndpointName string `json:"psc_endpoint_name"`

	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpVpcEndpointInfoFromPb(pb *gcpVpcEndpointInfoPb) (*GcpVpcEndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpVpcEndpointInfo{}
	st.EndpointRegion = pb.EndpointRegion
	st.ProjectId = pb.ProjectId
	st.PscConnectionId = pb.PscConnectionId
	st.PscEndpointName = pb.PscEndpointName
	st.ServiceAttachmentId = pb.ServiceAttachmentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpVpcEndpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpVpcEndpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getCredentialRequestToPb(st *GetCredentialRequest) (*getCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialRequestPb{}
	pb.CredentialsId = st.CredentialsId

	return pb, nil
}

type getCredentialRequestPb struct {
	CredentialsId string `json:"-" url:"-"`
}

func getCredentialRequestFromPb(pb *getCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	st.CredentialsId = pb.CredentialsId

	return st, nil
}

func getEncryptionKeyRequestToPb(st *GetEncryptionKeyRequest) (*getEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEncryptionKeyRequestPb{}
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	return pb, nil
}

type getEncryptionKeyRequestPb struct {
	CustomerManagedKeyId string `json:"-" url:"-"`
}

func getEncryptionKeyRequestFromPb(pb *getEncryptionKeyRequestPb) (*GetEncryptionKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEncryptionKeyRequest{}
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId

	return st, nil
}

func getNetworkRequestToPb(st *GetNetworkRequest) (*getNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNetworkRequestPb{}
	pb.NetworkId = st.NetworkId

	return pb, nil
}

type getNetworkRequestPb struct {
	NetworkId string `json:"-" url:"-"`
}

func getNetworkRequestFromPb(pb *getNetworkRequestPb) (*GetNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkRequest{}
	st.NetworkId = pb.NetworkId

	return st, nil
}

func getPrivateAccesRequestToPb(st *GetPrivateAccesRequest) (*getPrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPrivateAccesRequestPb{}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	return pb, nil
}

type getPrivateAccesRequestPb struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

func getPrivateAccesRequestFromPb(pb *getPrivateAccesRequestPb) (*GetPrivateAccesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPrivateAccesRequest{}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId

	return st, nil
}

func getStorageRequestToPb(st *GetStorageRequest) (*getStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStorageRequestPb{}
	pb.StorageConfigurationId = st.StorageConfigurationId

	return pb, nil
}

type getStorageRequestPb struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

func getStorageRequestFromPb(pb *getStorageRequestPb) (*GetStorageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStorageRequest{}
	st.StorageConfigurationId = pb.StorageConfigurationId

	return st, nil
}

func getVpcEndpointRequestToPb(st *GetVpcEndpointRequest) (*getVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getVpcEndpointRequestPb{}
	pb.VpcEndpointId = st.VpcEndpointId

	return pb, nil
}

type getVpcEndpointRequestPb struct {
	VpcEndpointId string `json:"-" url:"-"`
}

func getVpcEndpointRequestFromPb(pb *getVpcEndpointRequestPb) (*GetVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetVpcEndpointRequest{}
	st.VpcEndpointId = pb.VpcEndpointId

	return st, nil
}

func getWorkspaceRequestToPb(st *GetWorkspaceRequest) (*getWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type getWorkspaceRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

func getWorkspaceRequestFromPb(pb *getWorkspaceRequestPb) (*GetWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func gkeConfigToPb(st *GkeConfig) (*gkeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gkeConfigPb{}
	pb.ConnectivityType = st.ConnectivityType

	pb.MasterIpRange = st.MasterIpRange

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gkeConfigPb struct {
	ConnectivityType GkeConfigConnectivityType `json:"connectivity_type,omitempty"`

	MasterIpRange string `json:"master_ip_range,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gkeConfigFromPb(pb *gkeConfigPb) (*GkeConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GkeConfig{}
	st.ConnectivityType = pb.ConnectivityType
	st.MasterIpRange = pb.MasterIpRange

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gkeConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gkeConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func networkToPb(st *Network) (*networkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkPb{}
	pb.AccountId = st.AccountId

	pb.CreationTime = st.CreationTime

	pb.ErrorMessages = st.ErrorMessages

	pb.GcpNetworkInfo = st.GcpNetworkInfo

	pb.NetworkId = st.NetworkId

	pb.NetworkName = st.NetworkName

	pb.SecurityGroupIds = st.SecurityGroupIds

	pb.SubnetIds = st.SubnetIds

	pb.VpcEndpoints = st.VpcEndpoints

	pb.VpcId = st.VpcId

	pb.VpcStatus = st.VpcStatus

	pb.WarningMessages = st.WarningMessages

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type networkPb struct {
	AccountId string `json:"account_id,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`

	ErrorMessages []NetworkHealth `json:"error_messages,omitempty"`

	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`

	NetworkId string `json:"network_id,omitempty"`

	NetworkName string `json:"network_name,omitempty"`

	SecurityGroupIds []string `json:"security_group_ids,omitempty"`

	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`

	VpcId string `json:"vpc_id,omitempty"`

	VpcStatus VpcStatus `json:"vpc_status,omitempty"`

	WarningMessages []NetworkWarning `json:"warning_messages,omitempty"`

	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkFromPb(pb *networkPb) (*Network, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Network{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime
	st.ErrorMessages = pb.ErrorMessages
	st.GcpNetworkInfo = pb.GcpNetworkInfo
	st.NetworkId = pb.NetworkId
	st.NetworkName = pb.NetworkName
	st.SecurityGroupIds = pb.SecurityGroupIds
	st.SubnetIds = pb.SubnetIds
	st.VpcEndpoints = pb.VpcEndpoints
	st.VpcId = pb.VpcId
	st.VpcStatus = pb.VpcStatus
	st.WarningMessages = pb.WarningMessages
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func networkHealthToPb(st *NetworkHealth) (*networkHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkHealthPb{}
	pb.ErrorMessage = st.ErrorMessage

	pb.ErrorType = st.ErrorType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type networkHealthPb struct {
	ErrorMessage string `json:"error_message,omitempty"`

	ErrorType ErrorType `json:"error_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkHealthFromPb(pb *networkHealthPb) (*NetworkHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkHealth{}
	st.ErrorMessage = pb.ErrorMessage
	st.ErrorType = pb.ErrorType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkHealthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkHealthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func networkVpcEndpointsToPb(st *NetworkVpcEndpoints) (*networkVpcEndpointsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkVpcEndpointsPb{}
	pb.DataplaneRelay = st.DataplaneRelay

	pb.RestApi = st.RestApi

	return pb, nil
}

type networkVpcEndpointsPb struct {
	DataplaneRelay []string `json:"dataplane_relay"`

	RestApi []string `json:"rest_api"`
}

func networkVpcEndpointsFromPb(pb *networkVpcEndpointsPb) (*NetworkVpcEndpoints, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkVpcEndpoints{}
	st.DataplaneRelay = pb.DataplaneRelay
	st.RestApi = pb.RestApi

	return st, nil
}

func networkWarningToPb(st *NetworkWarning) (*networkWarningPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkWarningPb{}
	pb.WarningMessage = st.WarningMessage

	pb.WarningType = st.WarningType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type networkWarningPb struct {
	WarningMessage string `json:"warning_message,omitempty"`

	WarningType WarningType `json:"warning_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkWarningFromPb(pb *networkWarningPb) (*NetworkWarning, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkWarning{}
	st.WarningMessage = pb.WarningMessage
	st.WarningType = pb.WarningType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkWarningPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkWarningPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func privateAccessSettingsToPb(st *PrivateAccessSettings) (*privateAccessSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &privateAccessSettingsPb{}
	pb.AccountId = st.AccountId

	pb.AllowedVpcEndpointIds = st.AllowedVpcEndpointIds

	pb.PrivateAccessLevel = st.PrivateAccessLevel

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.PrivateAccessSettingsName = st.PrivateAccessSettingsName

	pb.PublicAccessEnabled = st.PublicAccessEnabled

	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type privateAccessSettingsPb struct {
	AccountId string `json:"account_id,omitempty"`

	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`

	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`

	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`

	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`

	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`

	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func privateAccessSettingsFromPb(pb *privateAccessSettingsPb) (*PrivateAccessSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivateAccessSettings{}
	st.AccountId = pb.AccountId
	st.AllowedVpcEndpointIds = pb.AllowedVpcEndpointIds
	st.PrivateAccessLevel = pb.PrivateAccessLevel
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.PrivateAccessSettingsName = pb.PrivateAccessSettingsName
	st.PublicAccessEnabled = pb.PublicAccessEnabled
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *privateAccessSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st privateAccessSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func replaceResponseToPb(st *ReplaceResponse) (*replaceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &replaceResponsePb{}

	return pb, nil
}

type replaceResponsePb struct {
}

func replaceResponseFromPb(pb *replaceResponsePb) (*ReplaceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReplaceResponse{}

	return st, nil
}

func rootBucketInfoToPb(st *RootBucketInfo) (*rootBucketInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rootBucketInfoPb{}
	pb.BucketName = st.BucketName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type rootBucketInfoPb struct {
	BucketName string `json:"bucket_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func rootBucketInfoFromPb(pb *rootBucketInfoPb) (*RootBucketInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RootBucketInfo{}
	st.BucketName = pb.BucketName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *rootBucketInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st rootBucketInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func storageConfigurationToPb(st *StorageConfiguration) (*storageConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &storageConfigurationPb{}
	pb.AccountId = st.AccountId

	pb.CreationTime = st.CreationTime

	pb.RootBucketInfo = st.RootBucketInfo

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageConfigurationName = st.StorageConfigurationName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type storageConfigurationPb struct {
	AccountId string `json:"account_id,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`

	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`

	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`

	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func storageConfigurationFromPb(pb *storageConfigurationPb) (*StorageConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StorageConfiguration{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime
	st.RootBucketInfo = pb.RootBucketInfo
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageConfigurationName = pb.StorageConfigurationName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *storageConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st storageConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func stsRoleToPb(st *StsRole) (*stsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stsRolePb{}
	pb.ExternalId = st.ExternalId

	pb.RoleArn = st.RoleArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type stsRolePb struct {
	ExternalId string `json:"external_id,omitempty"`

	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func stsRoleFromPb(pb *stsRolePb) (*StsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StsRole{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *stsRolePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st stsRolePb) MarshalJSON() ([]byte, error) {
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

func updateWorkspaceRequestToPb(st *UpdateWorkspaceRequest) (*updateWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceRequestPb{}
	pb.AwsRegion = st.AwsRegion

	pb.CredentialsId = st.CredentialsId

	pb.CustomTags = st.CustomTags

	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId

	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId

	pb.NetworkId = st.NetworkId

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId

	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateWorkspaceRequestPb struct {
	AwsRegion string `json:"aws_region,omitempty"`

	CredentialsId string `json:"credentials_id,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`

	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`

	NetworkId string `json:"network_id,omitempty"`

	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`

	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`

	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`

	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateWorkspaceRequestFromPb(pb *updateWorkspaceRequestPb) (*UpdateWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceRequest{}
	st.AwsRegion = pb.AwsRegion
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.NetworkId = pb.NetworkId
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func upsertPrivateAccessSettingsRequestToPb(st *UpsertPrivateAccessSettingsRequest) (*upsertPrivateAccessSettingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &upsertPrivateAccessSettingsRequestPb{}
	pb.AllowedVpcEndpointIds = st.AllowedVpcEndpointIds

	pb.PrivateAccessLevel = st.PrivateAccessLevel

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.PrivateAccessSettingsName = st.PrivateAccessSettingsName

	pb.PublicAccessEnabled = st.PublicAccessEnabled

	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type upsertPrivateAccessSettingsRequestPb struct {
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`

	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`

	PrivateAccessSettingsId string `json:"-" url:"-"`

	PrivateAccessSettingsName string `json:"private_access_settings_name"`

	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`

	Region string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func upsertPrivateAccessSettingsRequestFromPb(pb *upsertPrivateAccessSettingsRequestPb) (*UpsertPrivateAccessSettingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertPrivateAccessSettingsRequest{}
	st.AllowedVpcEndpointIds = pb.AllowedVpcEndpointIds
	st.PrivateAccessLevel = pb.PrivateAccessLevel
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.PrivateAccessSettingsName = pb.PrivateAccessSettingsName
	st.PublicAccessEnabled = pb.PublicAccessEnabled
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *upsertPrivateAccessSettingsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st upsertPrivateAccessSettingsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func vpcEndpointToPb(st *VpcEndpoint) (*vpcEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &vpcEndpointPb{}
	pb.AccountId = st.AccountId

	pb.AwsAccountId = st.AwsAccountId

	pb.AwsEndpointServiceId = st.AwsEndpointServiceId

	pb.AwsVpcEndpointId = st.AwsVpcEndpointId

	pb.GcpVpcEndpointInfo = st.GcpVpcEndpointInfo

	pb.Region = st.Region

	pb.State = st.State

	pb.UseCase = st.UseCase

	pb.VpcEndpointId = st.VpcEndpointId

	pb.VpcEndpointName = st.VpcEndpointName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type vpcEndpointPb struct {
	AccountId string `json:"account_id,omitempty"`

	AwsAccountId string `json:"aws_account_id,omitempty"`

	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`

	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`

	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`

	Region string `json:"region,omitempty"`

	State string `json:"state,omitempty"`

	UseCase EndpointUseCase `json:"use_case,omitempty"`

	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`

	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func vpcEndpointFromPb(pb *vpcEndpointPb) (*VpcEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VpcEndpoint{}
	st.AccountId = pb.AccountId
	st.AwsAccountId = pb.AwsAccountId
	st.AwsEndpointServiceId = pb.AwsEndpointServiceId
	st.AwsVpcEndpointId = pb.AwsVpcEndpointId
	st.GcpVpcEndpointInfo = pb.GcpVpcEndpointInfo
	st.Region = pb.Region
	st.State = pb.State
	st.UseCase = pb.UseCase
	st.VpcEndpointId = pb.VpcEndpointId
	st.VpcEndpointName = pb.VpcEndpointName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *vpcEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st vpcEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func workspaceToPb(st *Workspace) (*workspacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspacePb{}
	pb.AccountId = st.AccountId

	pb.AwsRegion = st.AwsRegion

	pb.AzureWorkspaceInfo = st.AzureWorkspaceInfo

	pb.Cloud = st.Cloud

	pb.CloudResourceContainer = st.CloudResourceContainer

	pb.CreationTime = st.CreationTime

	pb.CredentialsId = st.CredentialsId

	pb.CustomTags = st.CustomTags

	pb.DeploymentName = st.DeploymentName

	pb.ExternalCustomerInfo = st.ExternalCustomerInfo

	pb.GcpManagedNetworkConfig = st.GcpManagedNetworkConfig

	pb.GkeConfig = st.GkeConfig

	pb.IsNoPublicIpEnabled = st.IsNoPublicIpEnabled

	pb.Location = st.Location

	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId

	pb.NetworkId = st.NetworkId

	pb.PricingTier = st.PricingTier

	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId

	pb.WorkspaceId = st.WorkspaceId

	pb.WorkspaceName = st.WorkspaceName

	pb.WorkspaceStatus = st.WorkspaceStatus

	pb.WorkspaceStatusMessage = st.WorkspaceStatusMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type workspacePb struct {
	AccountId string `json:"account_id,omitempty"`

	AwsRegion string `json:"aws_region,omitempty"`

	AzureWorkspaceInfo *AzureWorkspaceInfo `json:"azure_workspace_info,omitempty"`

	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`

	CredentialsId string `json:"credentials_id,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DeploymentName string `json:"deployment_name,omitempty"`

	ExternalCustomerInfo *ExternalCustomerInfo `json:"external_customer_info,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`

	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`

	Location string `json:"location,omitempty"`

	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`

	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`

	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`

	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`

	WorkspaceId int64 `json:"workspace_id,omitempty"`

	WorkspaceName string `json:"workspace_name,omitempty"`

	WorkspaceStatus WorkspaceStatus `json:"workspace_status,omitempty"`

	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceFromPb(pb *workspacePb) (*Workspace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Workspace{}
	st.AccountId = pb.AccountId
	st.AwsRegion = pb.AwsRegion
	st.AzureWorkspaceInfo = pb.AzureWorkspaceInfo
	st.Cloud = pb.Cloud
	st.CloudResourceContainer = pb.CloudResourceContainer
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.DeploymentName = pb.DeploymentName
	st.ExternalCustomerInfo = pb.ExternalCustomerInfo
	st.GcpManagedNetworkConfig = pb.GcpManagedNetworkConfig
	st.GkeConfig = pb.GkeConfig
	st.IsNoPublicIpEnabled = pb.IsNoPublicIpEnabled
	st.Location = pb.Location
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkId = pb.NetworkId
	st.PricingTier = pb.PricingTier
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceId = pb.WorkspaceId
	st.WorkspaceName = pb.WorkspaceName
	st.WorkspaceStatus = pb.WorkspaceStatus
	st.WorkspaceStatusMessage = pb.WorkspaceStatusMessage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
