// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"encoding/json"
	"fmt"
)

type AwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *StsRole
}

func (st *AwsCredentials) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsCredentialsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsCredentialsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsCredentials) MarshalJSON() ([]byte, error) {
	pb, err := awsCredentialsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string
	// The AWS KMS key's Amazon Resource Name (ARN).
	// Wire name: 'key_arn'
	KeyArn string
	// The AWS KMS key region.
	// Wire name: 'key_region'
	KeyRegion string
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool

	ForceSendFields []string `tf:"-"`
}

func (st *AwsKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := awsKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	// Wire name: 'resource_group'
	ResourceGroup string
	// Azure Subscription ID
	// Wire name: 'subscription_id'
	SubscriptionId string

	ForceSendFields []string `tf:"-"`
}

func (st *AzureWorkspaceInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureWorkspaceInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureWorkspaceInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureWorkspaceInfo) MarshalJSON() ([]byte, error) {
	pb, err := azureWorkspaceInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {
	// The general workspace configurations that are specific to Google Cloud.
	// Wire name: 'gcp'
	Gcp *CustomerFacingGcpCloudResourceContainer
}

func (st *CloudResourceContainer) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cloudResourceContainerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cloudResourceContainerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CloudResourceContainer) MarshalJSON() ([]byte, error) {
	pb, err := cloudResourceContainerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	// Wire name: 'key_alias'
	KeyAlias string
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	// Wire name: 'key_arn'
	KeyArn string
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	// Wire name: 'reuse_key_for_cluster_volumes'
	ReuseKeyForClusterVolumes bool

	ForceSendFields []string `tf:"-"`
}

func (st *CreateAwsKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAwsKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAwsKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAwsKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := createAwsKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialAwsCredentials struct {

	// Wire name: 'sts_role'
	StsRole *CreateCredentialStsRole
}

func (st *CreateCredentialAwsCredentials) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialAwsCredentialsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialAwsCredentialsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialAwsCredentials) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialAwsCredentialsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialRequest struct {

	// Wire name: 'aws_credentials'
	AwsCredentials CreateCredentialAwsCredentials
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string
}

func (st *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateCredentialStsRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCredentialStsRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCredentialStsRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCredentialStsRole) MarshalJSON() ([]byte, error) {
	pb, err := createCredentialStsRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCustomerManagedKeyRequest struct {

	// Wire name: 'aws_key_info'
	AwsKeyInfo *CreateAwsKeyInfo

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *CreateGcpKeyInfo
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase
}

func (st *CreateCustomerManagedKeyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCustomerManagedKeyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCustomerManagedKeyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCustomerManagedKeyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCustomerManagedKeyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string
}

func (st *CreateGcpKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createGcpKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createGcpKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateGcpKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := createGcpKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateNetworkRequest struct {
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	// Wire name: 'security_group_ids'
	SecurityGroupIds []string
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	// Wire name: 'subnet_ids'
	SubnetIds []string
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateNetworkRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createNetworkRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createNetworkRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	pb, err := createNetworkRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateStorageConfigurationRequest struct {
	// Root S3 bucket information.
	// Wire name: 'root_bucket_info'
	RootBucketInfo RootBucketInfo
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string
}

func (st *CreateStorageConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createStorageConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createStorageConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateStorageConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createStorageConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createVpcEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createVpcEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := createVpcEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	// Wire name: 'aws_region'
	AwsRegion string
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	// Wire name: 'cloud'
	Cloud string
	// The general workspace configurations that are specific to cloud
	// providers.
	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string
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
	DeploymentName string
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig
	// The configurations for the GKE cluster of a Databricks workspace.
	// Wire name: 'gke_config'
	GkeConfig *GkeConfig
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	// Wire name: 'location'
	Location string
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string

	// Wire name: 'network_id'
	NetworkId string
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	// Wire name: 'pricing_tier'
	PricingTier PricingTier
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
	PrivateAccessSettingsId string
	// The ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string
	// The workspace's human-readable name.
	// Wire name: 'workspace_name'
	WorkspaceName string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := createWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string

	// Wire name: 'aws_credentials'
	AwsCredentials *AwsCredentials
	// Time in epoch milliseconds when the credential was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Databricks credential configuration ID.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The human-readable name of the credential configuration object.
	// Wire name: 'credentials_name'
	CredentialsName string

	ForceSendFields []string `tf:"-"`
}

func (st *Credential) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &credentialPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := credentialFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Credential) MarshalJSON() ([]byte, error) {
	pb, err := credentialToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The custom tags key-value pairing that is attached to this workspace. The
// key-value pair is a string of utf-8 characters. The value can be an empty
// string, with maximum length of 255 characters. The key can be of maximum
// length of 127 characters, and cannot be empty.
type CustomTags map[string]string

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	// Wire name: 'project_id'
	ProjectId string

	ForceSendFields []string `tf:"-"`
}

func (st *CustomerFacingGcpCloudResourceContainer) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customerFacingGcpCloudResourceContainerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customerFacingGcpCloudResourceContainerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomerFacingGcpCloudResourceContainer) MarshalJSON() ([]byte, error) {
	pb, err := customerFacingGcpCloudResourceContainerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	// Wire name: 'account_id'
	AccountId string

	// Wire name: 'aws_key_info'
	AwsKeyInfo *AwsKeyInfo
	// Time in epoch milliseconds when the customer key was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// ID of the encryption key configuration object.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string

	// Wire name: 'gcp_key_info'
	GcpKeyInfo *GcpKeyInfo
	// The cases that the key can be used for.
	// Wire name: 'use_cases'
	UseCases []KeyUseCase

	ForceSendFields []string `tf:"-"`
}

func (st *CustomerManagedKey) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customerManagedKeyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customerManagedKeyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomerManagedKey) MarshalJSON() ([]byte, error) {
	pb, err := customerManagedKeyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete credential configuration
type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
}

func (st *DeleteCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete encryption key configuration
type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `tf:"-"`
}

func (st *DeleteEncryptionKeyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteEncryptionKeyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteEncryptionKeyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteEncryptionKeyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteEncryptionKeyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a network configuration
type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `tf:"-"`
}

func (st *DeleteNetworkRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNetworkRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNetworkRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNetworkRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteNetworkRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a private access settings object
type DeletePrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
}

func (st *DeletePrivateAccesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePrivateAccesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePrivateAccesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePrivateAccesRequest) MarshalJSON() ([]byte, error) {
	pb, err := deletePrivateAccesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteResponse struct {
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

// Delete storage configuration
type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`
}

func (st *DeleteStorageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteStorageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteStorageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteStorageRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteStorageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete VPC endpoint configuration
type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `tf:"-"`
}

func (st *DeleteVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteVpcEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteVpcEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteVpcEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a workspace
type DeleteWorkspaceRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *DeleteWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns EndpointUseCase to satisfy [pflag.Value] interface
func (f *EndpointUseCase) Type() string {
	return "EndpointUseCase"
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

// Type always returns ErrorType to satisfy [pflag.Value] interface
func (f *ErrorType) Type() string {
	return "ErrorType"
}

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	// Wire name: 'authoritative_user_email'
	AuthoritativeUserEmail string
	// The authoritative user full name.
	// Wire name: 'authoritative_user_full_name'
	AuthoritativeUserFullName string
	// The legal entity name for the external workspace
	// Wire name: 'customer_name'
	CustomerName string

	ForceSendFields []string `tf:"-"`
}

func (st *ExternalCustomerInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalCustomerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalCustomerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalCustomerInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalCustomerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	// Wire name: 'kms_key_id'
	KmsKeyId string
}

func (st *GcpKeyInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpKeyInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpKeyInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpKeyInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcpKeyInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	GkeClusterPodIpRange string
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	// Wire name: 'gke_cluster_service_ip_range'
	GkeClusterServiceIpRange string
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	// Wire name: 'subnet_cidr'
	SubnetCidr string

	ForceSendFields []string `tf:"-"`
}

func (st *GcpManagedNetworkConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpManagedNetworkConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpManagedNetworkConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpManagedNetworkConfig) MarshalJSON() ([]byte, error) {
	pb, err := gcpManagedNetworkConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	// Wire name: 'network_project_id'
	NetworkProjectId string
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	// Wire name: 'pod_ip_range_name'
	PodIpRangeName string
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	// Wire name: 'service_ip_range_name'
	ServiceIpRangeName string
	// The ID of the subnet associated with this network.
	// Wire name: 'subnet_id'
	SubnetId string
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	// Wire name: 'subnet_region'
	SubnetRegion string
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	// Wire name: 'vpc_id'
	VpcId string
}

func (st *GcpNetworkInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpNetworkInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpNetworkInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpNetworkInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcpNetworkInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	// Wire name: 'endpoint_region'
	EndpointRegion string
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	// Wire name: 'project_id'
	ProjectId string
	// The unique ID of this PSC connection.
	// Wire name: 'psc_connection_id'
	PscConnectionId string
	// The name of the PSC endpoint in the Google Cloud project.
	// Wire name: 'psc_endpoint_name'
	PscEndpointName string
	// The service attachment this PSC connection connects to.
	// Wire name: 'service_attachment_id'
	ServiceAttachmentId string

	ForceSendFields []string `tf:"-"`
}

func (st *GcpVpcEndpointInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpVpcEndpointInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpVpcEndpointInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpVpcEndpointInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcpVpcEndpointInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get credential configuration
type GetCredentialRequest struct {
	// Databricks Account API credential configuration ID
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
}

func (st *GetCredentialRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get encryption key configuration
type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	// Wire name: 'customer_managed_key_id'
	CustomerManagedKeyId string `tf:"-"`
}

func (st *GetEncryptionKeyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEncryptionKeyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEncryptionKeyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEncryptionKeyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getEncryptionKeyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a network configuration
type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	// Wire name: 'network_id'
	NetworkId string `tf:"-"`
}

func (st *GetNetworkRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getNetworkRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getNetworkRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetNetworkRequest) MarshalJSON() ([]byte, error) {
	pb, err := getNetworkRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a private access settings object
type GetPrivateAccesRequest struct {
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
}

func (st *GetPrivateAccesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPrivateAccesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPrivateAccesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPrivateAccesRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPrivateAccesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get storage configuration
type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`
}

func (st *GetStorageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStorageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStorageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStorageRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStorageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a VPC endpoint configuration
type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `tf:"-"`
}

func (st *GetVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getVpcEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getVpcEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := getVpcEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a workspace
type GetWorkspaceRequest struct {
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *GetWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	ConnectivityType GkeConfigConnectivityType
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	// Wire name: 'master_ip_range'
	MasterIpRange string

	ForceSendFields []string `tf:"-"`
}

func (st *GkeConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gkeConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gkeConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GkeConfig) MarshalJSON() ([]byte, error) {
	pb, err := gkeConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns GkeConfigConnectivityType to satisfy [pflag.Value] interface
func (f *GkeConfigConnectivityType) Type() string {
	return "GkeConfigConnectivityType"
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

// Type always returns KeyUseCase to satisfy [pflag.Value] interface
func (f *KeyUseCase) Type() string {
	return "KeyUseCase"
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	// Wire name: 'account_id'
	AccountId string
	// Time in epoch milliseconds when the network was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Array of error messages about the network configuration.
	// Wire name: 'error_messages'
	ErrorMessages []NetworkHealth
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	// Wire name: 'gcp_network_info'
	GcpNetworkInfo *GcpNetworkInfo
	// The Databricks network configuration ID.
	// Wire name: 'network_id'
	NetworkId string
	// The human-readable name of the network configuration.
	// Wire name: 'network_name'
	NetworkName string

	// Wire name: 'security_group_ids'
	SecurityGroupIds []string

	// Wire name: 'subnet_ids'
	SubnetIds []string
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS PrivateLink].
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// Wire name: 'vpc_endpoints'
	VpcEndpoints *NetworkVpcEndpoints
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	// Wire name: 'vpc_id'
	VpcId string
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	// Wire name: 'vpc_status'
	VpcStatus VpcStatus
	// Array of warning messages about the network configuration.
	// Wire name: 'warning_messages'
	WarningMessages []NetworkWarning
	// Workspace ID associated with this network configuration.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func (st *Network) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Network) MarshalJSON() ([]byte, error) {
	pb, err := networkToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NetworkHealth struct {
	// Details of the error.
	// Wire name: 'error_message'
	ErrorMessage string
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	// Wire name: 'error_type'
	ErrorType ErrorType

	ForceSendFields []string `tf:"-"`
}

func (st *NetworkHealth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkHealthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkHealthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkHealth) MarshalJSON() ([]byte, error) {
	pb, err := networkHealthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	// Wire name: 'dataplane_relay'
	DataplaneRelay []string
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	// Wire name: 'rest_api'
	RestApi []string
}

func (st *NetworkVpcEndpoints) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkVpcEndpointsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkVpcEndpointsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkVpcEndpoints) MarshalJSON() ([]byte, error) {
	pb, err := networkVpcEndpointsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NetworkWarning struct {
	// Details of the warning.
	// Wire name: 'warning_message'
	WarningMessage string
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	// Wire name: 'warning_type'
	WarningType WarningType

	ForceSendFields []string `tf:"-"`
}

func (st *NetworkWarning) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkWarningPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkWarningFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkWarning) MarshalJSON() ([]byte, error) {
	pb, err := networkWarningToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns PricingTier to satisfy [pflag.Value] interface
func (f *PricingTier) Type() string {
	return "PricingTier"
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

// Type always returns PrivateAccessLevel to satisfy [pflag.Value] interface
func (f *PrivateAccessLevel) Type() string {
	return "PrivateAccessLevel"
}

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string
	// An array of Databricks VPC endpoint IDs.
	// Wire name: 'allowed_vpc_endpoint_ids'
	AllowedVpcEndpointIds []string
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel
	// Databricks private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool
	// The cloud region for workspaces attached to this private access settings
	// object.
	// Wire name: 'region'
	Region string

	ForceSendFields []string `tf:"-"`
}

func (st *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &privateAccessSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := privateAccessSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	pb, err := privateAccessSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ReplaceResponse struct {
}

func (st *ReplaceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &replaceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := replaceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReplaceResponse) MarshalJSON() ([]byte, error) {
	pb, err := replaceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	// Wire name: 'bucket_name'
	BucketName string

	ForceSendFields []string `tf:"-"`
}

func (st *RootBucketInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rootBucketInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rootBucketInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RootBucketInfo) MarshalJSON() ([]byte, error) {
	pb, err := rootBucketInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string
	// Time in epoch milliseconds when the storage configuration was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Root S3 bucket information.
	// Wire name: 'root_bucket_info'
	RootBucketInfo *RootBucketInfo
	// Databricks storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// The human-readable name of the storage configuration.
	// Wire name: 'storage_configuration_name'
	StorageConfigurationName string

	ForceSendFields []string `tf:"-"`
}

func (st *StorageConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &storageConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := storageConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StorageConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := storageConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	// Wire name: 'external_id'
	ExternalId string
	// The Amazon Resource Name (ARN) of the cross account role.
	// Wire name: 'role_arn'
	RoleArn string

	ForceSendFields []string `tf:"-"`
}

func (st *StsRole) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stsRolePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stsRoleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StsRole) MarshalJSON() ([]byte, error) {
	pb, err := stsRoleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateResponse struct {
}

func (st *UpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'aws_region'
	AwsRegion string
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string

	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	// Wire name: 'network_id'
	NetworkId string
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string
	// Workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	AllowedVpcEndpointIds []string
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ACCOUNT` level access (the default) allows only VPC endpoints
	// that are registered in your Databricks account connect to your workspace.
	// * `ENDPOINT` level access allows only specified VPC endpoints connect to
	// your workspace. For details, see `allowed_vpc_endpoint_ids`.
	// Wire name: 'private_access_level'
	PrivateAccessLevel PrivateAccessLevel
	// Databricks Account API private access settings ID.
	// Wire name: 'private_access_settings_id'
	PrivateAccessSettingsId string `tf:"-"`
	// The human-readable name of the private access settings object.
	// Wire name: 'private_access_settings_name'
	PrivateAccessSettingsName string
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	// Wire name: 'public_access_enabled'
	PublicAccessEnabled bool
	// The cloud region for workspaces associated with this private access
	// settings object.
	// Wire name: 'region'
	Region string

	ForceSendFields []string `tf:"-"`
}

func (st *UpsertPrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &upsertPrivateAccessSettingsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := upsertPrivateAccessSettingsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpsertPrivateAccessSettingsRequest) MarshalJSON() ([]byte, error) {
	pb, err := upsertPrivateAccessSettingsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	// Wire name: 'account_id'
	AccountId string
	// The AWS Account in which the VPC endpoint object exists.
	// Wire name: 'aws_account_id'
	AwsAccountId string
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	// Wire name: 'aws_endpoint_service_id'
	AwsEndpointServiceId string
	// The ID of the VPC endpoint object in AWS.
	// Wire name: 'aws_vpc_endpoint_id'
	AwsVpcEndpointId string
	// The Google Cloud specific information for this Private Service Connect
	// endpoint.
	// Wire name: 'gcp_vpc_endpoint_info'
	GcpVpcEndpointInfo *GcpVpcEndpointInfo
	// The AWS region in which this VPC endpoint object exists.
	// Wire name: 'region'
	Region string
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	// Wire name: 'state'
	State string
	// This enumeration represents the type of Databricks VPC [endpoint service]
	// that was used when creating this VPC endpoint.
	//
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	// Wire name: 'use_case'
	UseCase EndpointUseCase
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string
	// The human-readable name of the storage configuration.
	// Wire name: 'vpc_endpoint_name'
	VpcEndpointName string

	ForceSendFields []string `tf:"-"`
}

func (st *VpcEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &vpcEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := vpcEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VpcEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := vpcEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns VpcStatus to satisfy [pflag.Value] interface
func (f *VpcStatus) Type() string {
	return "VpcStatus"
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

// Type always returns WarningType to satisfy [pflag.Value] interface
func (f *WarningType) Type() string {
	return "WarningType"
}

type Workspace struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	// Wire name: 'aws_region'
	AwsRegion string

	// Wire name: 'azure_workspace_info'
	AzureWorkspaceInfo *AzureWorkspaceInfo
	// The cloud name. This field always has the value `gcp`.
	// Wire name: 'cloud'
	Cloud string
	// The general workspace configurations that are specific to cloud
	// providers.
	// Wire name: 'cloud_resource_container'
	CloudResourceContainer *CloudResourceContainer
	// Time in epoch milliseconds when the workspace was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// ID of the workspace's credential configuration object.
	// Wire name: 'credentials_id'
	CredentialsId string
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	// Wire name: 'custom_tags'
	CustomTags map[string]string
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	// Wire name: 'deployment_name'
	DeploymentName string
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	// Wire name: 'external_customer_info'
	ExternalCustomerInfo *ExternalCustomerInfo
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.", All the IP range configurations must be
	// mutually exclusive. An attempt to create a workspace fails if Databricks
	// detects an IP range overlap.
	//
	// Specify custom IP ranges in CIDR format. The IP ranges for these fields
	// must not overlap, and all IP addresses must be entirely within the
	// following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
	// `192.168.0.0/16`, and `240.0.0.0/4`.
	//
	// The sizes of these IP ranges affect the maximum number of nodes for the
	// workspace.
	//
	// **Important**: Confirm the IP ranges used by your Databricks workspace
	// before creating the workspace. You cannot change them after your
	// workspace is deployed. If the IP address ranges for your Databricks are
	// too small, IP exhaustion can occur, causing your Databricks jobs to fail.
	// To determine the address range sizes that you need, Databricks provides a
	// calculator as a Microsoft Excel spreadsheet. See [calculate subnet sizes
	// for a new workspace].
	//
	// [calculate subnet sizes for a new workspace]: https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html
	// Wire name: 'gcp_managed_network_config'
	GcpManagedNetworkConfig *GcpManagedNetworkConfig
	// The configurations for the GKE cluster of a Databricks workspace.
	// Wire name: 'gke_config'
	GkeConfig *GkeConfig
	// Whether no public IP is enabled for the workspace.
	// Wire name: 'is_no_public_ip_enabled'
	IsNoPublicIpEnabled bool
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	// Wire name: 'location'
	Location string
	// ID of the key configuration for encrypting managed services.
	// Wire name: 'managed_services_customer_managed_key_id'
	ManagedServicesCustomerManagedKeyId string
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	// Wire name: 'network_id'
	NetworkId string
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing].
	//
	// [AWS Pricing]: https://databricks.com/product/aws-pricing
	// Wire name: 'pricing_tier'
	PricingTier PricingTier
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
	PrivateAccessSettingsId string
	// ID of the workspace's storage configuration object.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// ID of the key configuration for encrypting workspace storage.
	// Wire name: 'storage_customer_managed_key_id'
	StorageCustomerManagedKeyId string
	// A unique integer ID for the workspace
	// Wire name: 'workspace_id'
	WorkspaceId int64
	// The human-readable name of the workspace.
	// Wire name: 'workspace_name'
	WorkspaceName string
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	// Wire name: 'workspace_status'
	WorkspaceStatus WorkspaceStatus
	// Message describing the current workspace status.
	// Wire name: 'workspace_status_message'
	WorkspaceStatusMessage string

	ForceSendFields []string `tf:"-"`
}

func (st *Workspace) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspacePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Workspace) MarshalJSON() ([]byte, error) {
	pb, err := workspaceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns WorkspaceStatus to satisfy [pflag.Value] interface
func (f *WorkspaceStatus) Type() string {
	return "WorkspaceStatus"
}
