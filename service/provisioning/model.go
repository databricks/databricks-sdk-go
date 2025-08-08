// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/provisioning/provisioningpb"
)

type AwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *StsRole ``
}

func AwsCredentialsToPb(st *AwsCredentials) (*provisioningpb.AwsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.AwsCredentialsPb{}
	stsRolePb, err := StsRoleToPb(st.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRolePb != nil {
		pb.StsRole = stsRolePb
	}

	return pb, nil
}

func AwsCredentialsFromPb(pb *provisioningpb.AwsCredentialsPb) (*AwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsCredentials{}
	stsRoleField, err := StsRoleFromPb(pb.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRoleField != nil {
		st.StsRole = stsRoleField
	}

	return st, nil
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string ``
	// The AWS KMS key's Amazon Resource Name (ARN).
	// Wire name: 'key_arn'
	KeyArn string ``
	// The AWS KMS key region.
	// Wire name: 'key_region'
	KeyRegion string ``
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool     ``
	ForceSendFields           []string `tf:"-"`
}

func (s *AwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AwsKeyInfoToPb(st *AwsKeyInfo) (*provisioningpb.AwsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.AwsKeyInfoPb{}
	pb.KeyAlias = st.KeyAlias
	pb.KeyArn = st.KeyArn
	pb.KeyRegion = st.KeyRegion
	pb.ReuseKeyForClusterVolumes = st.ReuseKeyForClusterVolumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AwsKeyInfoFromPb(pb *provisioningpb.AwsKeyInfoPb) (*AwsKeyInfo, error) {
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

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	// Wire name: 'resource_group'
	ResourceGroup string ``
	// Azure Subscription ID
	// Wire name: 'subscription_id'
	SubscriptionId  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AzureWorkspaceInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureWorkspaceInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AzureWorkspaceInfoToPb(st *AzureWorkspaceInfo) (*provisioningpb.AzureWorkspaceInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.AzureWorkspaceInfoPb{}
	pb.ResourceGroup = st.ResourceGroup
	pb.SubscriptionId = st.SubscriptionId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AzureWorkspaceInfoFromPb(pb *provisioningpb.AzureWorkspaceInfoPb) (*AzureWorkspaceInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureWorkspaceInfo{}
	st.ResourceGroup = pb.ResourceGroup
	st.SubscriptionId = pb.SubscriptionId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {

	// Wire name: 'gcp'
	Gcp *CustomerFacingGcpCloudResourceContainer ``
}

func CloudResourceContainerToPb(st *CloudResourceContainer) (*provisioningpb.CloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CloudResourceContainerPb{}
	gcpPb, err := CustomerFacingGcpCloudResourceContainerToPb(st.Gcp)
	if err != nil {
		return nil, err
	}
	if gcpPb != nil {
		pb.Gcp = gcpPb
	}

	return pb, nil
}

func CloudResourceContainerFromPb(pb *provisioningpb.CloudResourceContainerPb) (*CloudResourceContainer, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudResourceContainer{}
	gcpField, err := CustomerFacingGcpCloudResourceContainerFromPb(pb.Gcp)
	if err != nil {
		return nil, err
	}
	if gcpField != nil {
		st.Gcp = gcpField
	}

	return st, nil
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string ``
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	// Wire name: 'key_arn'
	KeyArn string ``
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool     ``
	ForceSendFields           []string `tf:"-"`
}

func (s *CreateAwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAwsKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateAwsKeyInfoToPb(st *CreateAwsKeyInfo) (*provisioningpb.CreateAwsKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateAwsKeyInfoPb{}
	pb.KeyAlias = st.KeyAlias
	pb.KeyArn = st.KeyArn
	pb.ReuseKeyForClusterVolumes = st.ReuseKeyForClusterVolumes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateAwsKeyInfoFromPb(pb *provisioningpb.CreateAwsKeyInfoPb) (*CreateAwsKeyInfo, error) {
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

type CreateCredentialAwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *CreateCredentialStsRole ``
}

func CreateCredentialAwsCredentialsToPb(st *CreateCredentialAwsCredentials) (*provisioningpb.CreateCredentialAwsCredentialsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateCredentialAwsCredentialsPb{}
	stsRolePb, err := CreateCredentialStsRoleToPb(st.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRolePb != nil {
		pb.StsRole = stsRolePb
	}

	return pb, nil
}

func CreateCredentialAwsCredentialsFromPb(pb *provisioningpb.CreateCredentialAwsCredentialsPb) (*CreateCredentialAwsCredentials, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialAwsCredentials{}
	stsRoleField, err := CreateCredentialStsRoleFromPb(pb.StsRole)
	if err != nil {
		return nil, err
	}
	if stsRoleField != nil {
		st.StsRole = stsRoleField
	}

	return st, nil
}

type CreateCredentialRequest struct {

	// Wire name: 'aws_credentials'
	AwsCredentials CreateCredentialAwsCredentials ``
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string ``
}

func CreateCredentialRequestToPb(st *CreateCredentialRequest) (*provisioningpb.CreateCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateCredentialRequestPb{}
	awsCredentialsPb, err := CreateCredentialAwsCredentialsToPb(&st.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsPb != nil {
		pb.AwsCredentials = *awsCredentialsPb
	}
	pb.CredentialsName = st.CredentialsName

	return pb, nil
}

func CreateCredentialRequestFromPb(pb *provisioningpb.CreateCredentialRequestPb) (*CreateCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialRequest{}
	awsCredentialsField, err := CreateCredentialAwsCredentialsFromPb(&pb.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsField != nil {
		st.AwsCredentials = *awsCredentialsField
	}
	st.CredentialsName = pb.CredentialsName

	return st, nil
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateCredentialStsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialStsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateCredentialStsRoleToPb(st *CreateCredentialStsRole) (*provisioningpb.CreateCredentialStsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateCredentialStsRolePb{}
	pb.RoleArn = st.RoleArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateCredentialStsRoleFromPb(pb *provisioningpb.CreateCredentialStsRolePb) (*CreateCredentialStsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCredentialStsRole{}
	st.RoleArn = pb.RoleArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateCustomerManagedKeyRequest struct {

	// Wire name: 'aws_key_info'
	AwsKeyInfo *CreateAwsKeyInfo ``

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *CreateGcpKeyInfo ``
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase ``
}

func CreateCustomerManagedKeyRequestToPb(st *CreateCustomerManagedKeyRequest) (*provisioningpb.CreateCustomerManagedKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateCustomerManagedKeyRequestPb{}
	awsKeyInfoPb, err := CreateAwsKeyInfoToPb(st.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoPb != nil {
		pb.AwsKeyInfo = awsKeyInfoPb
	}
	gcpKeyInfoPb, err := CreateGcpKeyInfoToPb(st.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoPb != nil {
		pb.GcpKeyInfo = gcpKeyInfoPb
	}

	var useCasesPb []provisioningpb.KeyUseCasePb
	for _, item := range st.UseCases {
		itemPb, err := KeyUseCaseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			useCasesPb = append(useCasesPb, *itemPb)
		}
	}
	pb.UseCases = useCasesPb

	return pb, nil
}

func CreateCustomerManagedKeyRequestFromPb(pb *provisioningpb.CreateCustomerManagedKeyRequestPb) (*CreateCustomerManagedKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomerManagedKeyRequest{}
	awsKeyInfoField, err := CreateAwsKeyInfoFromPb(pb.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoField != nil {
		st.AwsKeyInfo = awsKeyInfoField
	}
	gcpKeyInfoField, err := CreateGcpKeyInfoFromPb(pb.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoField != nil {
		st.GcpKeyInfo = gcpKeyInfoField
	}

	var useCasesField []KeyUseCase
	for _, itemPb := range pb.UseCases {
		item, err := KeyUseCaseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			useCasesField = append(useCasesField, *item)
		}
	}
	st.UseCases = useCasesField

	return st, nil
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string ``
}

func CreateGcpKeyInfoToPb(st *CreateGcpKeyInfo) (*provisioningpb.CreateGcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateGcpKeyInfoPb{}
	pb.KmsKeyId = st.KmsKeyId

	return pb, nil
}

func CreateGcpKeyInfoFromPb(pb *provisioningpb.CreateGcpKeyInfoPb) (*CreateGcpKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateGcpKeyInfo{}
	st.KmsKeyId = pb.KmsKeyId

	return st, nil
}

type CreateNetworkRequest struct {

	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo ``
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string ``
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	// Wire name: 'security_group_ids'
	SecurityGroupIds []string ``
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	// Wire name: 'subnet_ids'
	SubnetIds []string ``

	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints ``
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateNetworkRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateNetworkRequestToPb(st *CreateNetworkRequest) (*provisioningpb.CreateNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateNetworkRequestPb{}
	gcpNetworkInfoPb, err := GcpNetworkInfoToPb(st.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoPb != nil {
		pb.GcpNetworkInfo = gcpNetworkInfoPb
	}
	pb.NetworkName = st.NetworkName
	pb.SecurityGroupIds = st.SecurityGroupIds
	pb.SubnetIds = st.SubnetIds
	vpcEndpointsPb, err := NetworkVpcEndpointsToPb(st.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsPb != nil {
		pb.VpcEndpoints = vpcEndpointsPb
	}
	pb.VpcId = st.VpcId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateNetworkRequestFromPb(pb *provisioningpb.CreateNetworkRequestPb) (*CreateNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkRequest{}
	gcpNetworkInfoField, err := GcpNetworkInfoFromPb(pb.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoField != nil {
		st.GcpNetworkInfo = gcpNetworkInfoField
	}
	st.NetworkName = pb.NetworkName
	st.SecurityGroupIds = pb.SecurityGroupIds
	st.SubnetIds = pb.SubnetIds
	vpcEndpointsField, err := NetworkVpcEndpointsFromPb(pb.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsField != nil {
		st.VpcEndpoints = vpcEndpointsField
	}
	st.VpcId = pb.VpcId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateStorageConfigurationRequest struct {

	// Wire name: 'root_bucket_info'
	RootBucketInfo RootBucketInfo ``
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string ``
}

func CreateStorageConfigurationRequestToPb(st *CreateStorageConfigurationRequest) (*provisioningpb.CreateStorageConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateStorageConfigurationRequestPb{}
	rootBucketInfoPb, err := RootBucketInfoToPb(&st.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoPb != nil {
		pb.RootBucketInfo = *rootBucketInfoPb
	}
	pb.StorageConfigurationName = st.StorageConfigurationName

	return pb, nil
}

func CreateStorageConfigurationRequestFromPb(pb *provisioningpb.CreateStorageConfigurationRequestPb) (*CreateStorageConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateStorageConfigurationRequest{}
	rootBucketInfoField, err := RootBucketInfoFromPb(&pb.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoField != nil {
		st.RootBucketInfo = *rootBucketInfoField
	}
	st.StorageConfigurationName = pb.StorageConfigurationName

	return st, nil
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string ``

	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo ``
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string ``
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateVpcEndpointRequestToPb(st *CreateVpcEndpointRequest) (*provisioningpb.CreateVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateVpcEndpointRequestPb{}
	pb.AwsVpcEndpointId = st.AwsVpcEndpointId
	gcpVpcEndpointInfoPb, err := GcpVpcEndpointInfoToPb(st.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoPb != nil {
		pb.GcpVpcEndpointInfo = gcpVpcEndpointInfoPb
	}
	pb.Region = st.Region
	pb.VpcEndpointName = st.VpcEndpointName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateVpcEndpointRequestFromPb(pb *provisioningpb.CreateVpcEndpointRequestPb) (*CreateVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVpcEndpointRequest{}
	st.AwsVpcEndpointId = pb.AwsVpcEndpointId
	gcpVpcEndpointInfoField, err := GcpVpcEndpointInfoFromPb(pb.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoField != nil {
		st.GcpVpcEndpointInfo = gcpVpcEndpointInfoField
	}
	st.Region = pb.Region
	st.VpcEndpointName = pb.VpcEndpointName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	// Wire name: 'aws_region'
	AwsRegion string ``
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	// Wire name: 'cloud'
	Cloud string ``

	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer ``
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string ``
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for the web application and REST APIs is
	// `<workspace-deployment-name>.cloud.databricks.com`. For example, if the
	// deployment name is `abcsales`, your workspace URL will be
	// `https://abcsales.cloud.databricks.com`. Hyphens are allowed. This
	// property supports only the set of characters that are allowed in a
	// subdomain.
	//
	// To set this value, you must have a deployment name prefix. Contact your
	// Databricks account team to add an account deployment name prefix to your
	// account.
	//
	// Workspace deployment names follow the account prefix and a hyphen. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `workspace-1`, the JSON response for the
	// `deployment_name` field becomes `acme-workspace-1`. The workspace URL
	// would be `acme-workspace-1.cloud.databricks.com`.
	//
	// You can also set the `deployment_name` to the reserved keyword `EMPTY` if
	// you want the deployment name to only include the deployment prefix. For
	// example, if your account's deployment prefix is `acme` and the workspace
	// deployment name is `EMPTY`, the `deployment_name` becomes `acme` only and
	// the workspace URL is `acme.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	//
	// If a new workspace omits this property, the server generates a unique
	// deployment name for you with the pattern `dbc-xxxxxxxx-xxxx`.
	// Wire name: 'deployment_name'
	DeploymentName string ``

	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig ``

	// Wire name: 'gke_config'
	GkeConfig *GkeConfig ``
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool ``
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	// Wire name: 'location'
	Location string ``
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string ``

	// Wire name: 'network_id'
	NetworkId string ``

	// Wire name: 'pricing_tier'
	PricingTier PricingTier ``
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. This ID must be specified for customers using [AWS
	// PrivateLink] for either front-end (user-to-workspace connection),
	// back-end (data plane to control plane connection), or both connection
	// types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string ``
	// The ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string ``
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string ``
	// The workspace's human-readable name.
	// Wire name: 'workspace_name'
	WorkspaceName   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateWorkspaceRequestToPb(st *CreateWorkspaceRequest) (*provisioningpb.CreateWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CreateWorkspaceRequestPb{}
	pb.AwsRegion = st.AwsRegion
	pb.Cloud = st.Cloud
	cloudResourceContainerPb, err := CloudResourceContainerToPb(st.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerPb != nil {
		pb.CloudResourceContainer = cloudResourceContainerPb
	}
	pb.CredentialsId = st.CredentialsId
	pb.CustomTags = st.CustomTags
	pb.DeploymentName = st.DeploymentName
	gcpManagedNetworkConfigPb, err := GcpManagedNetworkConfigToPb(st.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigPb != nil {
		pb.GcpManagedNetworkConfig = gcpManagedNetworkConfigPb
	}
	gkeConfigPb, err := GkeConfigToPb(st.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigPb != nil {
		pb.GkeConfig = gkeConfigPb
	}
	pb.IsNoPublicIpEnabled = st.IsNoPublicIpEnabled
	pb.Location = st.Location
	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId
	pb.NetworkId = st.NetworkId
	pricingTierPb, err := PricingTierToPb(&st.PricingTier)
	if err != nil {
		return nil, err
	}
	if pricingTierPb != nil {
		pb.PricingTier = *pricingTierPb
	}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId
	pb.WorkspaceName = st.WorkspaceName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateWorkspaceRequestFromPb(pb *provisioningpb.CreateWorkspaceRequestPb) (*CreateWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWorkspaceRequest{}
	st.AwsRegion = pb.AwsRegion
	st.Cloud = pb.Cloud
	cloudResourceContainerField, err := CloudResourceContainerFromPb(pb.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerField != nil {
		st.CloudResourceContainer = cloudResourceContainerField
	}
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.DeploymentName = pb.DeploymentName
	gcpManagedNetworkConfigField, err := GcpManagedNetworkConfigFromPb(pb.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigField != nil {
		st.GcpManagedNetworkConfig = gcpManagedNetworkConfigField
	}
	gkeConfigField, err := GkeConfigFromPb(pb.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigField != nil {
		st.GkeConfig = gkeConfigField
	}
	st.IsNoPublicIpEnabled = pb.IsNoPublicIpEnabled
	st.Location = pb.Location
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkId = pb.NetworkId
	pricingTierField, err := PricingTierFromPb(&pb.PricingTier)
	if err != nil {
		return nil, err
	}
	if pricingTierField != nil {
		st.PricingTier = *pricingTierField
	}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceName = pb.WorkspaceName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string ``

	// Wire name: 'aws_credentials'
	AwsCredentials *AwsCredentials ``
	// Time in epoch milliseconds when the credential was created.
	// Wire name: 'creation_time'
	CreationTime int64 ``
	// Databricks credential configuration ID.
	// Wire name: 'credentials_id'
	CredentialsId string ``
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Credential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Credential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CredentialToPb(st *Credential) (*provisioningpb.CredentialPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CredentialPb{}
	pb.AccountId = st.AccountId
	awsCredentialsPb, err := AwsCredentialsToPb(st.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsPb != nil {
		pb.AwsCredentials = awsCredentialsPb
	}
	pb.CreationTime = st.CreationTime
	pb.CredentialsId = st.CredentialsId
	pb.CredentialsName = st.CredentialsName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CredentialFromPb(pb *provisioningpb.CredentialPb) (*Credential, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Credential{}
	st.AccountId = pb.AccountId
	awsCredentialsField, err := AwsCredentialsFromPb(pb.AwsCredentials)
	if err != nil {
		return nil, err
	}
	if awsCredentialsField != nil {
		st.AwsCredentials = awsCredentialsField
	}
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.CredentialsName = pb.CredentialsName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	// Wire name: 'project_id'
	ProjectId       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CustomerFacingGcpCloudResourceContainer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingGcpCloudResourceContainer) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CustomerFacingGcpCloudResourceContainerToPb(st *CustomerFacingGcpCloudResourceContainer) (*provisioningpb.CustomerFacingGcpCloudResourceContainerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CustomerFacingGcpCloudResourceContainerPb{}
	pb.ProjectId = st.ProjectId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CustomerFacingGcpCloudResourceContainerFromPb(pb *provisioningpb.CustomerFacingGcpCloudResourceContainerPb) (*CustomerFacingGcpCloudResourceContainer, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerFacingGcpCloudResourceContainer{}
	st.ProjectId = pb.ProjectId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	// Wire name: 'account_id'
	AccountId string ``

	// Wire name: 'aws_key_info'
	AwsKeyInfo *AwsKeyInfo ``
	// Time in epoch milliseconds when the customer key was created.
	// Wire name: 'creation_time'
	CreationTime int64 ``
	// ID of the encryption key configuration object.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string ``

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *GcpKeyInfo ``
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases        []KeyUseCase ``
	ForceSendFields []string     `tf:"-"`
}

func (s *CustomerManagedKey) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerManagedKey) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CustomerManagedKeyToPb(st *CustomerManagedKey) (*provisioningpb.CustomerManagedKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.CustomerManagedKeyPb{}
	pb.AccountId = st.AccountId
	awsKeyInfoPb, err := AwsKeyInfoToPb(st.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoPb != nil {
		pb.AwsKeyInfo = awsKeyInfoPb
	}
	pb.CreationTime = st.CreationTime
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId
	gcpKeyInfoPb, err := GcpKeyInfoToPb(st.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoPb != nil {
		pb.GcpKeyInfo = gcpKeyInfoPb
	}

	var useCasesPb []provisioningpb.KeyUseCasePb
	for _, item := range st.UseCases {
		itemPb, err := KeyUseCaseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			useCasesPb = append(useCasesPb, *itemPb)
		}
	}
	pb.UseCases = useCasesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CustomerManagedKeyFromPb(pb *provisioningpb.CustomerManagedKeyPb) (*CustomerManagedKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerManagedKey{}
	st.AccountId = pb.AccountId
	awsKeyInfoField, err := AwsKeyInfoFromPb(pb.AwsKeyInfo)
	if err != nil {
		return nil, err
	}
	if awsKeyInfoField != nil {
		st.AwsKeyInfo = awsKeyInfoField
	}
	st.CreationTime = pb.CreationTime
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId
	gcpKeyInfoField, err := GcpKeyInfoFromPb(pb.GcpKeyInfo)
	if err != nil {
		return nil, err
	}
	if gcpKeyInfoField != nil {
		st.GcpKeyInfo = gcpKeyInfoField
	}

	var useCasesField []KeyUseCase
	for _, itemPb := range pb.UseCases {
		item, err := KeyUseCaseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			useCasesField = append(useCasesField, *item)
		}
	}
	st.UseCases = useCasesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
}

func DeleteCredentialRequestToPb(st *DeleteCredentialRequest) (*provisioningpb.DeleteCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeleteCredentialRequestPb{}
	pb.CredentialsId = st.CredentialsId

	return pb, nil
}

func DeleteCredentialRequestFromPb(pb *provisioningpb.DeleteCredentialRequestPb) (*DeleteCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCredentialRequest{}
	st.CredentialsId = pb.CredentialsId

	return st, nil
}

type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `tf:"-"`
}

func DeleteEncryptionKeyRequestToPb(st *DeleteEncryptionKeyRequest) (*provisioningpb.DeleteEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeleteEncryptionKeyRequestPb{}
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	return pb, nil
}

func DeleteEncryptionKeyRequestFromPb(pb *provisioningpb.DeleteEncryptionKeyRequestPb) (*DeleteEncryptionKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteEncryptionKeyRequest{}
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId

	return st, nil
}

type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `tf:"-"`
}

func DeleteNetworkRequestToPb(st *DeleteNetworkRequest) (*provisioningpb.DeleteNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeleteNetworkRequestPb{}
	pb.NetworkId = st.NetworkId

	return pb, nil
}

func DeleteNetworkRequestFromPb(pb *provisioningpb.DeleteNetworkRequestPb) (*DeleteNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkRequest{}
	st.NetworkId = pb.NetworkId

	return st, nil
}

type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
}

func DeletePrivateAccesRequestToPb(st *DeletePrivateAccesRequest) (*provisioningpb.DeletePrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeletePrivateAccesRequestPb{}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	return pb, nil
}

func DeletePrivateAccesRequestFromPb(pb *provisioningpb.DeletePrivateAccesRequestPb) (*DeletePrivateAccesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePrivateAccesRequest{}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId

	return st, nil
}

type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`
}

func DeleteStorageRequestToPb(st *DeleteStorageRequest) (*provisioningpb.DeleteStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeleteStorageRequestPb{}
	pb.StorageConfigurationId = st.StorageConfigurationId

	return pb, nil
}

func DeleteStorageRequestFromPb(pb *provisioningpb.DeleteStorageRequestPb) (*DeleteStorageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteStorageRequest{}
	st.StorageConfigurationId = pb.StorageConfigurationId

	return st, nil
}

type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `tf:"-"`
}

func DeleteVpcEndpointRequestToPb(st *DeleteVpcEndpointRequest) (*provisioningpb.DeleteVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeleteVpcEndpointRequestPb{}
	pb.VpcEndpointId = st.VpcEndpointId

	return pb, nil
}

func DeleteVpcEndpointRequestFromPb(pb *provisioningpb.DeleteVpcEndpointRequestPb) (*DeleteVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVpcEndpointRequest{}
	st.VpcEndpointId = pb.VpcEndpointId

	return st, nil
}

type DeleteWorkspaceRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func DeleteWorkspaceRequestToPb(st *DeleteWorkspaceRequest) (*provisioningpb.DeleteWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.DeleteWorkspaceRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func DeleteWorkspaceRequestFromPb(pb *provisioningpb.DeleteWorkspaceRequestPb) (*DeleteWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWorkspaceRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

// This enumeration represents the type of Databricks VPC [endpoint service]
// that was used when creating this VPC endpoint.
//
// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
type EndpointUseCase string

const EndpointUseCaseDataplaneRelayAccess EndpointUseCase = `DATAPLANE_RELAY_ACCESS`

const EndpointUseCaseWorkspaceAccess EndpointUseCase = `WORKSPACE_ACCESS`

// String representation for [fmt.Print]
func (f *EndpointUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointUseCase) Set(v string) error {
	switch v {
	case `DATAPLANE_RELAY_ACCESS`, `WORKSPACE_ACCESS`:
		*f = EndpointUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATAPLANE_RELAY_ACCESS", "WORKSPACE_ACCESS"`, v)
	}
}

// Values returns all possible values for EndpointUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointUseCase) Values() []EndpointUseCase {
	return []EndpointUseCase{
		EndpointUseCaseDataplaneRelayAccess,
		EndpointUseCaseWorkspaceAccess,
	}
}

// Type always returns EndpointUseCase to satisfy [pflag.Value] interface
func (f *EndpointUseCase) Type() string {
	return "EndpointUseCase"
}

func EndpointUseCaseToPb(st *EndpointUseCase) (*provisioningpb.EndpointUseCasePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.EndpointUseCasePb(*st)
	return &pb, nil
}

func EndpointUseCaseFromPb(pb *provisioningpb.EndpointUseCasePb) (*EndpointUseCase, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointUseCase(*pb)
	return &st, nil
}

// The AWS resource associated with this error: credentials, VPC, subnet,
// security group, or network ACL.
type ErrorType string

const ErrorTypeCredentials ErrorType = `credentials`

const ErrorTypeNetworkAcl ErrorType = `networkAcl`

const ErrorTypeSecurityGroup ErrorType = `securityGroup`

const ErrorTypeSubnet ErrorType = `subnet`

const ErrorTypeVpc ErrorType = `vpc`

// String representation for [fmt.Print]
func (f *ErrorType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ErrorType) Set(v string) error {
	switch v {
	case `credentials`, `networkAcl`, `securityGroup`, `subnet`, `vpc`:
		*f = ErrorType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "credentials", "networkAcl", "securityGroup", "subnet", "vpc"`, v)
	}
}

// Values returns all possible values for ErrorType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ErrorType) Values() []ErrorType {
	return []ErrorType{
		ErrorTypeCredentials,
		ErrorTypeNetworkAcl,
		ErrorTypeSecurityGroup,
		ErrorTypeSubnet,
		ErrorTypeVpc,
	}
}

// Type always returns ErrorType to satisfy [pflag.Value] interface
func (f *ErrorType) Type() string {
	return "ErrorType"
}

func ErrorTypeToPb(st *ErrorType) (*provisioningpb.ErrorTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.ErrorTypePb(*st)
	return &pb, nil
}

func ErrorTypeFromPb(pb *provisioningpb.ErrorTypePb) (*ErrorType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ErrorType(*pb)
	return &st, nil
}

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	// Wire name: 'authoritative_user_email'
	AuthoritativeUserEmail string ``
	// The authoritative user full name.
	// Wire name: 'authoritative_user_full_name'
	AuthoritativeUserFullName string ``
	// The legal entity name for the external workspace
	// Wire name: 'customer_name'
	CustomerName    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalCustomerInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalCustomerInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalCustomerInfoToPb(st *ExternalCustomerInfo) (*provisioningpb.ExternalCustomerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.ExternalCustomerInfoPb{}
	pb.AuthoritativeUserEmail = st.AuthoritativeUserEmail
	pb.AuthoritativeUserFullName = st.AuthoritativeUserFullName
	pb.CustomerName = st.CustomerName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalCustomerInfoFromPb(pb *provisioningpb.ExternalCustomerInfoPb) (*ExternalCustomerInfo, error) {
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

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string ``
}

func GcpKeyInfoToPb(st *GcpKeyInfo) (*provisioningpb.GcpKeyInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GcpKeyInfoPb{}
	pb.KmsKeyId = st.KmsKeyId

	return pb, nil
}

func GcpKeyInfoFromPb(pb *provisioningpb.GcpKeyInfoPb) (*GcpKeyInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpKeyInfo{}
	st.KmsKeyId = pb.KmsKeyId

	return st, nil
}

// The network settings for the workspace. The configurations are only for
// Databricks-managed VPCs. It is ignored if you specify a customer-managed VPC
// in the `network_id` field.", All the IP range configurations must be mutually
// exclusive. An attempt to create a workspace fails if Databricks detects an IP
// range overlap.
//
// Specify custom IP ranges in CIDR format. The IP ranges for these fields must
// not overlap, and all IP addresses must be entirely within the following
// ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`, `192.168.0.0/16`, and
// `240.0.0.0/4`.
//
// The sizes of these IP ranges affect the maximum number of nodes for the
// workspace.
//
// **Important**: Confirm the IP ranges used by your Databricks workspace before
// creating the workspace. You cannot change them after your workspace is
// deployed. If the IP address ranges for your Databricks are too small, IP
// exhaustion can occur, causing your Databricks jobs to fail. To determine the
// address range sizes that you need, Databricks provides a calculator as a
// Microsoft Excel spreadsheet. See [calculate subnet sizes for a new
// workspace].
//
// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
type GcpManagedNetworkConfig struct {
	// The IP range from which to allocate GKE cluster pods. No bigger than `/9`
	// and no smaller than `/21`.
	// Wire name: 'gke_cluster_pod_ip_range'
	GkeClusterPodIpRange string ``
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	// Wire name: 'gke_cluster_service_ip_range'
	GkeClusterServiceIpRange string ``
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	// Wire name: 'subnet_cidr'
	SubnetCidr      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GcpManagedNetworkConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpManagedNetworkConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GcpManagedNetworkConfigToPb(st *GcpManagedNetworkConfig) (*provisioningpb.GcpManagedNetworkConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GcpManagedNetworkConfigPb{}
	pb.GkeClusterPodIpRange = st.GkeClusterPodIpRange
	pb.GkeClusterServiceIpRange = st.GkeClusterServiceIpRange
	pb.SubnetCidr = st.SubnetCidr

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GcpManagedNetworkConfigFromPb(pb *provisioningpb.GcpManagedNetworkConfigPb) (*GcpManagedNetworkConfig, error) {
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

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	// Wire name: 'network_project_id'
	NetworkProjectId string ``
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	// Wire name: 'pod_ip_range_name'
	PodIpRangeName string ``
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	// Wire name: 'service_ip_range_name'
	ServiceIpRangeName string ``
	// The ID of the subnet associated with this network.
	// Wire name: 'subnet_id'
	SubnetId string ``
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	// Wire name: 'subnet_region'
	SubnetRegion string ``
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string ``
}

func GcpNetworkInfoToPb(st *GcpNetworkInfo) (*provisioningpb.GcpNetworkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GcpNetworkInfoPb{}
	pb.NetworkProjectId = st.NetworkProjectId
	pb.PodIpRangeName = st.PodIpRangeName
	pb.ServiceIpRangeName = st.ServiceIpRangeName
	pb.SubnetId = st.SubnetId
	pb.SubnetRegion = st.SubnetRegion
	pb.VpcId = st.VpcId

	return pb, nil
}

func GcpNetworkInfoFromPb(pb *provisioningpb.GcpNetworkInfoPb) (*GcpNetworkInfo, error) {
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

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	// Wire name: 'endpoint_region'
	EndpointRegion string ``
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	// Wire name: 'project_id'
	ProjectId string ``
	// The unique ID of this PSC connection.
	// Wire name: 'psc_connection_id'
	PscConnectionId string ``
	// The name of the PSC endpoint in the Google Cloud project.
	// Wire name: 'psc_endpoint_name'
	PscEndpointName string ``
	// The service attachment this PSC connection connects to.
	// Wire name: 'service_attachment_id'
	ServiceAttachmentId string   ``
	ForceSendFields     []string `tf:"-"`
}

func (s *GcpVpcEndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpVpcEndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GcpVpcEndpointInfoToPb(st *GcpVpcEndpointInfo) (*provisioningpb.GcpVpcEndpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GcpVpcEndpointInfoPb{}
	pb.EndpointRegion = st.EndpointRegion
	pb.ProjectId = st.ProjectId
	pb.PscConnectionId = st.PscConnectionId
	pb.PscEndpointName = st.PscEndpointName
	pb.ServiceAttachmentId = st.ServiceAttachmentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GcpVpcEndpointInfoFromPb(pb *provisioningpb.GcpVpcEndpointInfoPb) (*GcpVpcEndpointInfo, error) {
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

type GetCredentialRequest struct {
	// Databricks Account API credential configuration ID
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
}

func GetCredentialRequestToPb(st *GetCredentialRequest) (*provisioningpb.GetCredentialRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetCredentialRequestPb{}
	pb.CredentialsId = st.CredentialsId

	return pb, nil
}

func GetCredentialRequestFromPb(pb *provisioningpb.GetCredentialRequestPb) (*GetCredentialRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialRequest{}
	st.CredentialsId = pb.CredentialsId

	return st, nil
}

type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `tf:"-"`
}

func GetEncryptionKeyRequestToPb(st *GetEncryptionKeyRequest) (*provisioningpb.GetEncryptionKeyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetEncryptionKeyRequestPb{}
	pb.CustomerManagedKeyId = st.CustomerManagedKeyId

	return pb, nil
}

func GetEncryptionKeyRequestFromPb(pb *provisioningpb.GetEncryptionKeyRequestPb) (*GetEncryptionKeyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEncryptionKeyRequest{}
	st.CustomerManagedKeyId = pb.CustomerManagedKeyId

	return st, nil
}

type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `tf:"-"`
}

func GetNetworkRequestToPb(st *GetNetworkRequest) (*provisioningpb.GetNetworkRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetNetworkRequestPb{}
	pb.NetworkId = st.NetworkId

	return pb, nil
}

func GetNetworkRequestFromPb(pb *provisioningpb.GetNetworkRequestPb) (*GetNetworkRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkRequest{}
	st.NetworkId = pb.NetworkId

	return st, nil
}

type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
}

func GetPrivateAccesRequestToPb(st *GetPrivateAccesRequest) (*provisioningpb.GetPrivateAccesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetPrivateAccesRequestPb{}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId

	return pb, nil
}

func GetPrivateAccesRequestFromPb(pb *provisioningpb.GetPrivateAccesRequestPb) (*GetPrivateAccesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPrivateAccesRequest{}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId

	return st, nil
}

type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`
}

func GetStorageRequestToPb(st *GetStorageRequest) (*provisioningpb.GetStorageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetStorageRequestPb{}
	pb.StorageConfigurationId = st.StorageConfigurationId

	return pb, nil
}

func GetStorageRequestFromPb(pb *provisioningpb.GetStorageRequestPb) (*GetStorageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStorageRequest{}
	st.StorageConfigurationId = pb.StorageConfigurationId

	return st, nil
}

type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `tf:"-"`
}

func GetVpcEndpointRequestToPb(st *GetVpcEndpointRequest) (*provisioningpb.GetVpcEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetVpcEndpointRequestPb{}
	pb.VpcEndpointId = st.VpcEndpointId

	return pb, nil
}

func GetVpcEndpointRequestFromPb(pb *provisioningpb.GetVpcEndpointRequestPb) (*GetVpcEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetVpcEndpointRequest{}
	st.VpcEndpointId = pb.VpcEndpointId

	return st, nil
}

type GetWorkspaceRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func GetWorkspaceRequestToPb(st *GetWorkspaceRequest) (*provisioningpb.GetWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GetWorkspaceRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func GetWorkspaceRequestFromPb(pb *provisioningpb.GetWorkspaceRequestPb) (*GetWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

// The configurations for the GKE cluster of a Databricks workspace.
type GkeConfig struct {
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network.
	//
	// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
	// workspace. The GKE nodes will not have public IPs.
	//
	// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of
	// a public GKE cluster have public IP addresses.
	// Wire name: 'connectivity_type'
	ConnectivityType GkeConfigConnectivityType ``
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	// Wire name: 'master_ip_range'
	MasterIpRange   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GkeConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GkeConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GkeConfigToPb(st *GkeConfig) (*provisioningpb.GkeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.GkeConfigPb{}
	connectivityTypePb, err := GkeConfigConnectivityTypeToPb(&st.ConnectivityType)
	if err != nil {
		return nil, err
	}
	if connectivityTypePb != nil {
		pb.ConnectivityType = *connectivityTypePb
	}
	pb.MasterIpRange = st.MasterIpRange

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GkeConfigFromPb(pb *provisioningpb.GkeConfigPb) (*GkeConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GkeConfig{}
	connectivityTypeField, err := GkeConfigConnectivityTypeFromPb(&pb.ConnectivityType)
	if err != nil {
		return nil, err
	}
	if connectivityTypeField != nil {
		st.ConnectivityType = *connectivityTypeField
	}
	st.MasterIpRange = pb.MasterIpRange

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Specifies the network connectivity types for the GKE nodes and the GKE master
// network.
//
// Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for the
// workspace. The GKE nodes will not have public IPs.
//
// Set to `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a
// public GKE cluster have public IP addresses.
type GkeConfigConnectivityType string

const GkeConfigConnectivityTypePrivateNodePublicMaster GkeConfigConnectivityType = `PRIVATE_NODE_PUBLIC_MASTER`

const GkeConfigConnectivityTypePublicNodePublicMaster GkeConfigConnectivityType = `PUBLIC_NODE_PUBLIC_MASTER`

// String representation for [fmt.Print]
func (f *GkeConfigConnectivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GkeConfigConnectivityType) Set(v string) error {
	switch v {
	case `PRIVATE_NODE_PUBLIC_MASTER`, `PUBLIC_NODE_PUBLIC_MASTER`:
		*f = GkeConfigConnectivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PRIVATE_NODE_PUBLIC_MASTER", "PUBLIC_NODE_PUBLIC_MASTER"`, v)
	}
}

// Values returns all possible values for GkeConfigConnectivityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GkeConfigConnectivityType) Values() []GkeConfigConnectivityType {
	return []GkeConfigConnectivityType{
		GkeConfigConnectivityTypePrivateNodePublicMaster,
		GkeConfigConnectivityTypePublicNodePublicMaster,
	}
}

// Type always returns GkeConfigConnectivityType to satisfy [pflag.Value] interface
func (f *GkeConfigConnectivityType) Type() string {
	return "GkeConfigConnectivityType"
}

func GkeConfigConnectivityTypeToPb(st *GkeConfigConnectivityType) (*provisioningpb.GkeConfigConnectivityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.GkeConfigConnectivityTypePb(*st)
	return &pb, nil
}

func GkeConfigConnectivityTypeFromPb(pb *provisioningpb.GkeConfigConnectivityTypePb) (*GkeConfigConnectivityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := GkeConfigConnectivityType(*pb)
	return &st, nil
}

// Possible values are: * `MANAGED_SERVICES`: Encrypts notebook and secret data
// in the control plane * `STORAGE`: Encrypts the workspace's root S3 bucket
// (root DBFS and system data) and, optionally, cluster EBS volumes.
type KeyUseCase string

// Encrypts notebook and secret data in the control plane
const KeyUseCaseManagedServices KeyUseCase = `MANAGED_SERVICES`

// Encrypts the workspace's root S3 bucket (root DBFS and system data) and,
// optionally, cluster EBS volumes.
const KeyUseCaseStorage KeyUseCase = `STORAGE`

// String representation for [fmt.Print]
func (f *KeyUseCase) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KeyUseCase) Set(v string) error {
	switch v {
	case `MANAGED_SERVICES`, `STORAGE`:
		*f = KeyUseCase(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MANAGED_SERVICES", "STORAGE"`, v)
	}
}

// Values returns all possible values for KeyUseCase.
//
// There is no guarantee on the order of the values in the slice.
func (f *KeyUseCase) Values() []KeyUseCase {
	return []KeyUseCase{
		KeyUseCaseManagedServices,
		KeyUseCaseStorage,
	}
}

// Type always returns KeyUseCase to satisfy [pflag.Value] interface
func (f *KeyUseCase) Type() string {
	return "KeyUseCase"
}

func KeyUseCaseToPb(st *KeyUseCase) (*provisioningpb.KeyUseCasePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.KeyUseCasePb(*st)
	return &pb, nil
}

func KeyUseCaseFromPb(pb *provisioningpb.KeyUseCasePb) (*KeyUseCase, error) {
	if pb == nil {
		return nil, nil
	}
	st := KeyUseCase(*pb)
	return &st, nil
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	// Wire name: 'account_id'
	AccountId string ``
	// Time in epoch milliseconds when the network was created.
	// Wire name: 'creation_time'
	CreationTime int64 ``
	// Array of error messages about the network configuration.
	// Wire name: 'error_messages'
	ErrorMessages []NetworkHealth ``

	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo ``
	// The Databricks network configuration ID.
	// Wire name: 'network_id'
	NetworkId string ``
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string ``

	// Wire name: 'security_group_ids'
	SecurityGroupIds []string ``

	// Wire name: 'subnet_ids'
	SubnetIds []string ``

	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints ``
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	// Wire name: 'vpc_id'
	VpcId string ``

	// Wire name: 'vpc_status'
	VpcStatus VpcStatus ``
	// Array of warning messages about the network configuration.
	// Wire name: 'warning_messages'
	WarningMessages []NetworkWarning ``
	// Workspace ID associated with this network configuration.
	// Wire name: 'workspace_id'
	WorkspaceId     int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *Network) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Network) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NetworkToPb(st *Network) (*provisioningpb.NetworkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.NetworkPb{}
	pb.AccountId = st.AccountId
	pb.CreationTime = st.CreationTime

	var errorMessagesPb []provisioningpb.NetworkHealthPb
	for _, item := range st.ErrorMessages {
		itemPb, err := NetworkHealthToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			errorMessagesPb = append(errorMessagesPb, *itemPb)
		}
	}
	pb.ErrorMessages = errorMessagesPb
	gcpNetworkInfoPb, err := GcpNetworkInfoToPb(st.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoPb != nil {
		pb.GcpNetworkInfo = gcpNetworkInfoPb
	}
	pb.NetworkId = st.NetworkId
	pb.NetworkName = st.NetworkName
	pb.SecurityGroupIds = st.SecurityGroupIds
	pb.SubnetIds = st.SubnetIds
	vpcEndpointsPb, err := NetworkVpcEndpointsToPb(st.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsPb != nil {
		pb.VpcEndpoints = vpcEndpointsPb
	}
	pb.VpcId = st.VpcId
	vpcStatusPb, err := VpcStatusToPb(&st.VpcStatus)
	if err != nil {
		return nil, err
	}
	if vpcStatusPb != nil {
		pb.VpcStatus = *vpcStatusPb
	}

	var warningMessagesPb []provisioningpb.NetworkWarningPb
	for _, item := range st.WarningMessages {
		itemPb, err := NetworkWarningToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			warningMessagesPb = append(warningMessagesPb, *itemPb)
		}
	}
	pb.WarningMessages = warningMessagesPb
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NetworkFromPb(pb *provisioningpb.NetworkPb) (*Network, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Network{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime

	var errorMessagesField []NetworkHealth
	for _, itemPb := range pb.ErrorMessages {
		item, err := NetworkHealthFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			errorMessagesField = append(errorMessagesField, *item)
		}
	}
	st.ErrorMessages = errorMessagesField
	gcpNetworkInfoField, err := GcpNetworkInfoFromPb(pb.GcpNetworkInfo)
	if err != nil {
		return nil, err
	}
	if gcpNetworkInfoField != nil {
		st.GcpNetworkInfo = gcpNetworkInfoField
	}
	st.NetworkId = pb.NetworkId
	st.NetworkName = pb.NetworkName
	st.SecurityGroupIds = pb.SecurityGroupIds
	st.SubnetIds = pb.SubnetIds
	vpcEndpointsField, err := NetworkVpcEndpointsFromPb(pb.VpcEndpoints)
	if err != nil {
		return nil, err
	}
	if vpcEndpointsField != nil {
		st.VpcEndpoints = vpcEndpointsField
	}
	st.VpcId = pb.VpcId
	vpcStatusField, err := VpcStatusFromPb(&pb.VpcStatus)
	if err != nil {
		return nil, err
	}
	if vpcStatusField != nil {
		st.VpcStatus = *vpcStatusField
	}

	var warningMessagesField []NetworkWarning
	for _, itemPb := range pb.WarningMessages {
		item, err := NetworkWarningFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			warningMessagesField = append(warningMessagesField, *item)
		}
	}
	st.WarningMessages = warningMessagesField
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type NetworkHealth struct {
	// Details of the error.
	// Wire name: 'error_message'
	ErrorMessage string ``

	// Wire name: 'error_type'
	ErrorType       ErrorType ``
	ForceSendFields []string  `tf:"-"`
}

func (s *NetworkHealth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkHealth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NetworkHealthToPb(st *NetworkHealth) (*provisioningpb.NetworkHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.NetworkHealthPb{}
	pb.ErrorMessage = st.ErrorMessage
	errorTypePb, err := ErrorTypeToPb(&st.ErrorType)
	if err != nil {
		return nil, err
	}
	if errorTypePb != nil {
		pb.ErrorType = *errorTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NetworkHealthFromPb(pb *provisioningpb.NetworkHealthPb) (*NetworkHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkHealth{}
	st.ErrorMessage = pb.ErrorMessage
	errorTypeField, err := ErrorTypeFromPb(&pb.ErrorType)
	if err != nil {
		return nil, err
	}
	if errorTypeField != nil {
		st.ErrorType = *errorTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	// Wire name: 'dataplane_relay'
	DataplaneRelay []string ``
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	// Wire name: 'rest_api'
	RestApi []string ``
}

func NetworkVpcEndpointsToPb(st *NetworkVpcEndpoints) (*provisioningpb.NetworkVpcEndpointsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.NetworkVpcEndpointsPb{}
	pb.DataplaneRelay = st.DataplaneRelay
	pb.RestApi = st.RestApi

	return pb, nil
}

func NetworkVpcEndpointsFromPb(pb *provisioningpb.NetworkVpcEndpointsPb) (*NetworkVpcEndpoints, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkVpcEndpoints{}
	st.DataplaneRelay = pb.DataplaneRelay
	st.RestApi = pb.RestApi

	return st, nil
}

type NetworkWarning struct {
	// Details of the warning.
	// Wire name: 'warning_message'
	WarningMessage string ``

	// Wire name: 'warning_type'
	WarningType     WarningType ``
	ForceSendFields []string    `tf:"-"`
}

func (s *NetworkWarning) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkWarning) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NetworkWarningToPb(st *NetworkWarning) (*provisioningpb.NetworkWarningPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.NetworkWarningPb{}
	pb.WarningMessage = st.WarningMessage
	warningTypePb, err := WarningTypeToPb(&st.WarningType)
	if err != nil {
		return nil, err
	}
	if warningTypePb != nil {
		pb.WarningType = *warningTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NetworkWarningFromPb(pb *provisioningpb.NetworkWarningPb) (*NetworkWarning, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkWarning{}
	st.WarningMessage = pb.WarningMessage
	warningTypeField, err := WarningTypeFromPb(&pb.WarningType)
	if err != nil {
		return nil, err
	}
	if warningTypeField != nil {
		st.WarningType = *warningTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The pricing tier of the workspace. For pricing tier information, see [AWS
// Pricing].
//
// [AWS Pricing]: https://databricks.com/product/aws-pricing
type PricingTier string

const PricingTierCommunityEdition PricingTier = `COMMUNITY_EDITION`

const PricingTierDedicated PricingTier = `DEDICATED`

const PricingTierEnterprise PricingTier = `ENTERPRISE`

const PricingTierPremium PricingTier = `PREMIUM`

const PricingTierStandard PricingTier = `STANDARD`

const PricingTierUnknown PricingTier = `UNKNOWN`

// String representation for [fmt.Print]
func (f *PricingTier) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PricingTier) Set(v string) error {
	switch v {
	case `COMMUNITY_EDITION`, `DEDICATED`, `ENTERPRISE`, `PREMIUM`, `STANDARD`, `UNKNOWN`:
		*f = PricingTier(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMUNITY_EDITION", "DEDICATED", "ENTERPRISE", "PREMIUM", "STANDARD", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for PricingTier.
//
// There is no guarantee on the order of the values in the slice.
func (f *PricingTier) Values() []PricingTier {
	return []PricingTier{
		PricingTierCommunityEdition,
		PricingTierDedicated,
		PricingTierEnterprise,
		PricingTierPremium,
		PricingTierStandard,
		PricingTierUnknown,
	}
}

// Type always returns PricingTier to satisfy [pflag.Value] interface
func (f *PricingTier) Type() string {
	return "PricingTier"
}

func PricingTierToPb(st *PricingTier) (*provisioningpb.PricingTierPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.PricingTierPb(*st)
	return &pb, nil
}

func PricingTierFromPb(pb *provisioningpb.PricingTierPb) (*PricingTier, error) {
	if pb == nil {
		return nil, nil
	}
	st := PricingTier(*pb)
	return &st, nil
}

// The private access level controls which VPC endpoints can connect to the UI
// or API of any workspace that attaches this private access settings object. *
// `ACCOUNT` level access (the default) allows only VPC endpoints that are
// registered in your Databricks account connect to your workspace. * `ENDPOINT`
// level access allows only specified VPC endpoints connect to your workspace.
// For details, see `allowed_vpc_endpoint_ids`.
type PrivateAccessLevel string

const PrivateAccessLevelAccount PrivateAccessLevel = `ACCOUNT`

const PrivateAccessLevelEndpoint PrivateAccessLevel = `ENDPOINT`

// String representation for [fmt.Print]
func (f *PrivateAccessLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateAccessLevel) Set(v string) error {
	switch v {
	case `ACCOUNT`, `ENDPOINT`:
		*f = PrivateAccessLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCOUNT", "ENDPOINT"`, v)
	}
}

// Values returns all possible values for PrivateAccessLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateAccessLevel) Values() []PrivateAccessLevel {
	return []PrivateAccessLevel{
		PrivateAccessLevelAccount,
		PrivateAccessLevelEndpoint,
	}
}

// Type always returns PrivateAccessLevel to satisfy [pflag.Value] interface
func (f *PrivateAccessLevel) Type() string {
	return "PrivateAccessLevel"
}

func PrivateAccessLevelToPb(st *PrivateAccessLevel) (*provisioningpb.PrivateAccessLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.PrivateAccessLevelPb(*st)
	return &pb, nil
}

func PrivateAccessLevelFromPb(pb *provisioningpb.PrivateAccessLevelPb) (*PrivateAccessLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PrivateAccessLevel(*pb)
	return &st, nil
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string ``
	// An array of Databricks VPC endpoint IDs.
	// Wire name: 'allowed_vpc_endpoint_ids'
	AllowedVpcEndpointIds []string ``

	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel ``
	// Databricks private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string ``
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string ``
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool ``
	// The cloud region for workspaces attached to this private access settings
	// object.
	// Wire name: 'region'
	Region          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PrivateAccessSettingsToPb(st *PrivateAccessSettings) (*provisioningpb.PrivateAccessSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.PrivateAccessSettingsPb{}
	pb.AccountId = st.AccountId
	pb.AllowedVpcEndpointIds = st.AllowedVpcEndpointIds
	privateAccessLevelPb, err := PrivateAccessLevelToPb(&st.PrivateAccessLevel)
	if err != nil {
		return nil, err
	}
	if privateAccessLevelPb != nil {
		pb.PrivateAccessLevel = *privateAccessLevelPb
	}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId
	pb.PrivateAccessSettingsName = st.PrivateAccessSettingsName
	pb.PublicAccessEnabled = st.PublicAccessEnabled
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PrivateAccessSettingsFromPb(pb *provisioningpb.PrivateAccessSettingsPb) (*PrivateAccessSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PrivateAccessSettings{}
	st.AccountId = pb.AccountId
	st.AllowedVpcEndpointIds = pb.AllowedVpcEndpointIds
	privateAccessLevelField, err := PrivateAccessLevelFromPb(&pb.PrivateAccessLevel)
	if err != nil {
		return nil, err
	}
	if privateAccessLevelField != nil {
		st.PrivateAccessLevel = *privateAccessLevelField
	}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.PrivateAccessSettingsName = pb.PrivateAccessSettingsName
	st.PublicAccessEnabled = pb.PublicAccessEnabled
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	// Wire name: 'bucket_name'
	BucketName      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RootBucketInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RootBucketInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RootBucketInfoToPb(st *RootBucketInfo) (*provisioningpb.RootBucketInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.RootBucketInfoPb{}
	pb.BucketName = st.BucketName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RootBucketInfoFromPb(pb *provisioningpb.RootBucketInfoPb) (*RootBucketInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RootBucketInfo{}
	st.BucketName = pb.BucketName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string ``
	// Time in epoch milliseconds when the storage configuration was created.
	// Wire name: 'creation_time'
	CreationTime int64 ``

	// Wire name: 'root_bucket_info'
	RootBucketInfo *RootBucketInfo ``
	// Databricks storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string ``
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string   ``
	ForceSendFields          []string `tf:"-"`
}

func (s *StorageConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StorageConfigurationToPb(st *StorageConfiguration) (*provisioningpb.StorageConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.StorageConfigurationPb{}
	pb.AccountId = st.AccountId
	pb.CreationTime = st.CreationTime
	rootBucketInfoPb, err := RootBucketInfoToPb(st.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoPb != nil {
		pb.RootBucketInfo = rootBucketInfoPb
	}
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.StorageConfigurationName = st.StorageConfigurationName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StorageConfigurationFromPb(pb *provisioningpb.StorageConfigurationPb) (*StorageConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StorageConfiguration{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime
	rootBucketInfoField, err := RootBucketInfoFromPb(pb.RootBucketInfo)
	if err != nil {
		return nil, err
	}
	if rootBucketInfoField != nil {
		st.RootBucketInfo = rootBucketInfoField
	}
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageConfigurationName = pb.StorageConfigurationName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	// Wire name: 'external_id'
	ExternalId string ``
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *StsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StsRoleToPb(st *StsRole) (*provisioningpb.StsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.StsRolePb{}
	pb.ExternalId = st.ExternalId
	pb.RoleArn = st.RoleArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StsRoleFromPb(pb *provisioningpb.StsRolePb) (*StsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StsRole{}
	st.ExternalId = pb.ExternalId
	st.RoleArn = pb.RoleArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'aws_region'
	AwsRegion string ``
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	// Wire name: 'credentials_id'
	CredentialsId string ``
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string ``

	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string ``
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	// Wire name: 'network_id'
	NetworkId string ``
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string ``
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string ``
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string ``
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId     int64    `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateWorkspaceRequestToPb(st *UpdateWorkspaceRequest) (*provisioningpb.UpdateWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.UpdateWorkspaceRequestPb{}
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

func UpdateWorkspaceRequestFromPb(pb *provisioningpb.UpdateWorkspaceRequestPb) (*UpdateWorkspaceRequest, error) {
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

type UpsertPrivateAccessSettingsRequest struct {
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID that
	// is returned when registering the VPC endpoint configuration in your
	// Databricks account. This is not the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints that in your account that can connect to your
	// workspace over AWS PrivateLink.
	//
	// If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access lists].
	//
	// [IP access lists]: https://docs.databricks.com/security/network/ip-access-list.html
	// Wire name: 'allowed_vpc_endpoint_ids'
	AllowedVpcEndpointIds []string ``

	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel ``
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string ``
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool ``
	// The cloud region for workspaces associated with this private access
	// settings object.
	// Wire name: 'region'
	Region          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpsertPrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertPrivateAccessSettingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpsertPrivateAccessSettingsRequestToPb(st *UpsertPrivateAccessSettingsRequest) (*provisioningpb.UpsertPrivateAccessSettingsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.UpsertPrivateAccessSettingsRequestPb{}
	pb.AllowedVpcEndpointIds = st.AllowedVpcEndpointIds
	privateAccessLevelPb, err := PrivateAccessLevelToPb(&st.PrivateAccessLevel)
	if err != nil {
		return nil, err
	}
	if privateAccessLevelPb != nil {
		pb.PrivateAccessLevel = *privateAccessLevelPb
	}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId
	pb.PrivateAccessSettingsName = st.PrivateAccessSettingsName
	pb.PublicAccessEnabled = st.PublicAccessEnabled
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpsertPrivateAccessSettingsRequestFromPb(pb *provisioningpb.UpsertPrivateAccessSettingsRequestPb) (*UpsertPrivateAccessSettingsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpsertPrivateAccessSettingsRequest{}
	st.AllowedVpcEndpointIds = pb.AllowedVpcEndpointIds
	privateAccessLevelField, err := PrivateAccessLevelFromPb(&pb.PrivateAccessLevel)
	if err != nil {
		return nil, err
	}
	if privateAccessLevelField != nil {
		st.PrivateAccessLevel = *privateAccessLevelField
	}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.PrivateAccessSettingsName = pb.PrivateAccessSettingsName
	st.PublicAccessEnabled = pb.PublicAccessEnabled
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	// Wire name: 'account_id'
	AccountId string ``
	// The AWS Account in which the VPC endpoint object exists.
	// Wire name: 'aws_account_id'
	AwsAccountId string ``
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	// Wire name: 'aws_endpoint_service_id'
	AwsEndpointServiceId string ``
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string ``

	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo ``
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string ``
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	// Wire name: 'state'
	State string ``

	// Wire name: 'use_case'
	UseCase EndpointUseCase ``
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string ``
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *VpcEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VpcEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VpcEndpointToPb(st *VpcEndpoint) (*provisioningpb.VpcEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.VpcEndpointPb{}
	pb.AccountId = st.AccountId
	pb.AwsAccountId = st.AwsAccountId
	pb.AwsEndpointServiceId = st.AwsEndpointServiceId
	pb.AwsVpcEndpointId = st.AwsVpcEndpointId
	gcpVpcEndpointInfoPb, err := GcpVpcEndpointInfoToPb(st.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoPb != nil {
		pb.GcpVpcEndpointInfo = gcpVpcEndpointInfoPb
	}
	pb.Region = st.Region
	pb.State = st.State
	useCasePb, err := EndpointUseCaseToPb(&st.UseCase)
	if err != nil {
		return nil, err
	}
	if useCasePb != nil {
		pb.UseCase = *useCasePb
	}
	pb.VpcEndpointId = st.VpcEndpointId
	pb.VpcEndpointName = st.VpcEndpointName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VpcEndpointFromPb(pb *provisioningpb.VpcEndpointPb) (*VpcEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VpcEndpoint{}
	st.AccountId = pb.AccountId
	st.AwsAccountId = pb.AwsAccountId
	st.AwsEndpointServiceId = pb.AwsEndpointServiceId
	st.AwsVpcEndpointId = pb.AwsVpcEndpointId
	gcpVpcEndpointInfoField, err := GcpVpcEndpointInfoFromPb(pb.GcpVpcEndpointInfo)
	if err != nil {
		return nil, err
	}
	if gcpVpcEndpointInfoField != nil {
		st.GcpVpcEndpointInfo = gcpVpcEndpointInfoField
	}
	st.Region = pb.Region
	st.State = pb.State
	useCaseField, err := EndpointUseCaseFromPb(&pb.UseCase)
	if err != nil {
		return nil, err
	}
	if useCaseField != nil {
		st.UseCase = *useCaseField
	}
	st.VpcEndpointId = pb.VpcEndpointId
	st.VpcEndpointName = pb.VpcEndpointName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The status of this network configuration object in terms of its use in a
// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`: Broken.
// * `WARNED`: Warned.
type VpcStatus string

// Broken.
const VpcStatusBroken VpcStatus = `BROKEN`

// Unattached.
const VpcStatusUnattached VpcStatus = `UNATTACHED`

// Valid.
const VpcStatusValid VpcStatus = `VALID`

// Warned.
const VpcStatusWarned VpcStatus = `WARNED`

// String representation for [fmt.Print]
func (f *VpcStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *VpcStatus) Set(v string) error {
	switch v {
	case `BROKEN`, `UNATTACHED`, `VALID`, `WARNED`:
		*f = VpcStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BROKEN", "UNATTACHED", "VALID", "WARNED"`, v)
	}
}

// Values returns all possible values for VpcStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *VpcStatus) Values() []VpcStatus {
	return []VpcStatus{
		VpcStatusBroken,
		VpcStatusUnattached,
		VpcStatusValid,
		VpcStatusWarned,
	}
}

// Type always returns VpcStatus to satisfy [pflag.Value] interface
func (f *VpcStatus) Type() string {
	return "VpcStatus"
}

func VpcStatusToPb(st *VpcStatus) (*provisioningpb.VpcStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.VpcStatusPb(*st)
	return &pb, nil
}

func VpcStatusFromPb(pb *provisioningpb.VpcStatusPb) (*VpcStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := VpcStatus(*pb)
	return &st, nil
}

// The AWS resource associated with this warning: a subnet or a security group.
type WarningType string

const WarningTypeSecurityGroup WarningType = `securityGroup`

const WarningTypeSubnet WarningType = `subnet`

// String representation for [fmt.Print]
func (f *WarningType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarningType) Set(v string) error {
	switch v {
	case `securityGroup`, `subnet`:
		*f = WarningType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "securityGroup", "subnet"`, v)
	}
}

// Values returns all possible values for WarningType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarningType) Values() []WarningType {
	return []WarningType{
		WarningTypeSecurityGroup,
		WarningTypeSubnet,
	}
}

// Type always returns WarningType to satisfy [pflag.Value] interface
func (f *WarningType) Type() string {
	return "WarningType"
}

func WarningTypeToPb(st *WarningType) (*provisioningpb.WarningTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.WarningTypePb(*st)
	return &pb, nil
}

func WarningTypeFromPb(pb *provisioningpb.WarningTypePb) (*WarningType, error) {
	if pb == nil {
		return nil, nil
	}
	st := WarningType(*pb)
	return &st, nil
}

type Workspace struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string ``
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	// Wire name: 'aws_region'
	AwsRegion string ``

	// Wire name: 'azure_workspace_info'
	AzureWorkspaceInfo *AzureWorkspaceInfo ``
	// The cloud name. This field always has the value `gcp`.
	// Wire name: 'cloud'
	Cloud string ``

	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer ``
	// Time in epoch milliseconds when the workspace was created.
	// Wire name: 'creation_time'
	CreationTime int64 ``
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string ``
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	// Wire name: 'deployment_name'
	DeploymentName string ``
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	// Wire name: 'external_customer_info'
	ExternalCustomerInfo *ExternalCustomerInfo ``

	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig ``

	// Wire name: 'gke_config'
	GkeConfig *GkeConfig ``
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool ``
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	// Wire name: 'location'
	Location string ``
	// ID of the key configuration for encrypting managed services.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string ``
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	// Wire name: 'network_id'
	NetworkId string ``

	// Wire name: 'pricing_tier'
	PricingTier PricingTier ``
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string ``
	// ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string ``
	// ID of the key configuration for encrypting workspace storage.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string ``
	// A unique integer ID for the workspace
	// Wire name: 'workspace_id'
	WorkspaceId int64 ``
	// The human-readable name of the workspace.
	// Wire name: 'workspace_name'
	WorkspaceName string ``

	// Wire name: 'workspace_status'
	WorkspaceStatus WorkspaceStatus ``
	// Message describing the current workspace status.
	// Wire name: 'workspace_status_message'
	WorkspaceStatusMessage string   ``
	ForceSendFields        []string `tf:"-"`
}

func (s *Workspace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Workspace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WorkspaceToPb(st *Workspace) (*provisioningpb.WorkspacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &provisioningpb.WorkspacePb{}
	pb.AccountId = st.AccountId
	pb.AwsRegion = st.AwsRegion
	azureWorkspaceInfoPb, err := AzureWorkspaceInfoToPb(st.AzureWorkspaceInfo)
	if err != nil {
		return nil, err
	}
	if azureWorkspaceInfoPb != nil {
		pb.AzureWorkspaceInfo = azureWorkspaceInfoPb
	}
	pb.Cloud = st.Cloud
	cloudResourceContainerPb, err := CloudResourceContainerToPb(st.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerPb != nil {
		pb.CloudResourceContainer = cloudResourceContainerPb
	}
	pb.CreationTime = st.CreationTime
	pb.CredentialsId = st.CredentialsId
	pb.CustomTags = st.CustomTags
	pb.DeploymentName = st.DeploymentName
	externalCustomerInfoPb, err := ExternalCustomerInfoToPb(st.ExternalCustomerInfo)
	if err != nil {
		return nil, err
	}
	if externalCustomerInfoPb != nil {
		pb.ExternalCustomerInfo = externalCustomerInfoPb
	}
	gcpManagedNetworkConfigPb, err := GcpManagedNetworkConfigToPb(st.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigPb != nil {
		pb.GcpManagedNetworkConfig = gcpManagedNetworkConfigPb
	}
	gkeConfigPb, err := GkeConfigToPb(st.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigPb != nil {
		pb.GkeConfig = gkeConfigPb
	}
	pb.IsNoPublicIpEnabled = st.IsNoPublicIpEnabled
	pb.Location = st.Location
	pb.ManagedServicesCustomerManagedKeyId = st.ManagedServicesCustomerManagedKeyId
	pb.NetworkId = st.NetworkId
	pricingTierPb, err := PricingTierToPb(&st.PricingTier)
	if err != nil {
		return nil, err
	}
	if pricingTierPb != nil {
		pb.PricingTier = *pricingTierPb
	}
	pb.PrivateAccessSettingsId = st.PrivateAccessSettingsId
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.StorageCustomerManagedKeyId = st.StorageCustomerManagedKeyId
	pb.WorkspaceId = st.WorkspaceId
	pb.WorkspaceName = st.WorkspaceName
	workspaceStatusPb, err := WorkspaceStatusToPb(&st.WorkspaceStatus)
	if err != nil {
		return nil, err
	}
	if workspaceStatusPb != nil {
		pb.WorkspaceStatus = *workspaceStatusPb
	}
	pb.WorkspaceStatusMessage = st.WorkspaceStatusMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WorkspaceFromPb(pb *provisioningpb.WorkspacePb) (*Workspace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Workspace{}
	st.AccountId = pb.AccountId
	st.AwsRegion = pb.AwsRegion
	azureWorkspaceInfoField, err := AzureWorkspaceInfoFromPb(pb.AzureWorkspaceInfo)
	if err != nil {
		return nil, err
	}
	if azureWorkspaceInfoField != nil {
		st.AzureWorkspaceInfo = azureWorkspaceInfoField
	}
	st.Cloud = pb.Cloud
	cloudResourceContainerField, err := CloudResourceContainerFromPb(pb.CloudResourceContainer)
	if err != nil {
		return nil, err
	}
	if cloudResourceContainerField != nil {
		st.CloudResourceContainer = cloudResourceContainerField
	}
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.CustomTags = pb.CustomTags
	st.DeploymentName = pb.DeploymentName
	externalCustomerInfoField, err := ExternalCustomerInfoFromPb(pb.ExternalCustomerInfo)
	if err != nil {
		return nil, err
	}
	if externalCustomerInfoField != nil {
		st.ExternalCustomerInfo = externalCustomerInfoField
	}
	gcpManagedNetworkConfigField, err := GcpManagedNetworkConfigFromPb(pb.GcpManagedNetworkConfig)
	if err != nil {
		return nil, err
	}
	if gcpManagedNetworkConfigField != nil {
		st.GcpManagedNetworkConfig = gcpManagedNetworkConfigField
	}
	gkeConfigField, err := GkeConfigFromPb(pb.GkeConfig)
	if err != nil {
		return nil, err
	}
	if gkeConfigField != nil {
		st.GkeConfig = gkeConfigField
	}
	st.IsNoPublicIpEnabled = pb.IsNoPublicIpEnabled
	st.Location = pb.Location
	st.ManagedServicesCustomerManagedKeyId = pb.ManagedServicesCustomerManagedKeyId
	st.NetworkId = pb.NetworkId
	pricingTierField, err := PricingTierFromPb(&pb.PricingTier)
	if err != nil {
		return nil, err
	}
	if pricingTierField != nil {
		st.PricingTier = *pricingTierField
	}
	st.PrivateAccessSettingsId = pb.PrivateAccessSettingsId
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.StorageCustomerManagedKeyId = pb.StorageCustomerManagedKeyId
	st.WorkspaceId = pb.WorkspaceId
	st.WorkspaceName = pb.WorkspaceName
	workspaceStatusField, err := WorkspaceStatusFromPb(&pb.WorkspaceStatus)
	if err != nil {
		return nil, err
	}
	if workspaceStatusField != nil {
		st.WorkspaceStatus = *workspaceStatusField
	}
	st.WorkspaceStatusMessage = pb.WorkspaceStatusMessage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// The status of the workspace. For workspace creation, usually it is set to
// `PROVISIONING` initially. Continue to check the status until the status is
// `RUNNING`.
type WorkspaceStatus string

const WorkspaceStatusBanned WorkspaceStatus = `BANNED`

const WorkspaceStatusCancelling WorkspaceStatus = `CANCELLING`

const WorkspaceStatusFailed WorkspaceStatus = `FAILED`

const WorkspaceStatusNotProvisioned WorkspaceStatus = `NOT_PROVISIONED`

const WorkspaceStatusProvisioning WorkspaceStatus = `PROVISIONING`

const WorkspaceStatusRunning WorkspaceStatus = `RUNNING`

// String representation for [fmt.Print]
func (f *WorkspaceStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WorkspaceStatus) Set(v string) error {
	switch v {
	case `BANNED`, `CANCELLING`, `FAILED`, `NOT_PROVISIONED`, `PROVISIONING`, `RUNNING`:
		*f = WorkspaceStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BANNED", "CANCELLING", "FAILED", "NOT_PROVISIONED", "PROVISIONING", "RUNNING"`, v)
	}
}

// Values returns all possible values for WorkspaceStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *WorkspaceStatus) Values() []WorkspaceStatus {
	return []WorkspaceStatus{
		WorkspaceStatusBanned,
		WorkspaceStatusCancelling,
		WorkspaceStatusFailed,
		WorkspaceStatusNotProvisioned,
		WorkspaceStatusProvisioning,
		WorkspaceStatusRunning,
	}
}

// Type always returns WorkspaceStatus to satisfy [pflag.Value] interface
func (f *WorkspaceStatus) Type() string {
	return "WorkspaceStatus"
}

func WorkspaceStatusToPb(st *WorkspaceStatus) (*provisioningpb.WorkspaceStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := provisioningpb.WorkspaceStatusPb(*st)
	return &pb, nil
}

func WorkspaceStatusFromPb(pb *provisioningpb.WorkspaceStatusPb) (*WorkspaceStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := WorkspaceStatus(*pb)
	return &st, nil
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
