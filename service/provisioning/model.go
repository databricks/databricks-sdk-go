// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AwsCredentials struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

type AwsDbManagedNetworkExtraInfo struct {
	// This field is need to populate worker env for DB managed VPC. It is
	// likely only for resource tracking/deletion purpose.
	DhcpOptionsId string `json:"dhcp_options_id,omitempty"`
	// This is the internal gateway which is different from the NAT gateway in
	// the NPIP VPC Infra. It is likely only for resource tracking/deletion
	// purpose.
	GatewayId string `json:"gateway_id,omitempty"`
	// Security group which the Vault will control, ensuring that
	// worker_opened_ports are actually open.
	ManagedSecurityGroup string `json:"managed_security_group,omitempty"`
	// Resources description for no public IP shard environment.
	NpipVpcInfra *NpipVpcInfra `json:"npipVpcInfra,omitempty"`
	// Security group which is given to the user to manage without Databricks
	// interference.
	UnmanagedSecurityGroup string `json:"unmanaged_security_group,omitempty"`
	// Contents of the secret key which gives ssh access to the workers.
	WorkerKeyContents string `json:"worker_key_contents,omitempty"`
	// Name of the keypair in AWS which allows sshing into the workers.
	WorkerKeypairName string `json:"worker_keypair_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsDbManagedNetworkExtraInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsDbManagedNetworkExtraInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AwsKeyInfo struct {
	// The alias name of the KMS key.
	KeyAlias string `json:"key_alias,omitempty"`
	// The the arn of the KMS key.
	KeyArn string `json:"key_arn"`
	// The region of the KMS key.
	KeyRegion string `json:"key_region"`
	// Indicates if the key should be used for cluster volumes. Can only be set
	// if the CMK can be used as a data plane key (use case storage)
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AwsNetworkInfo struct {
	// Additional information for DB managed VPC, which is mainly used to
	// populate WorkerEnvironment.
	DbManagedVpcExtraInfo *AwsDbManagedNetworkExtraInfo `json:"db_managed_vpc_extra_info,omitempty"`
	// The cloud-provided Security Group IDs that will be determine ingress and
	// egress rules for Cluster nodes.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// The cloud-provided Subnet IDs that will be available to Clusters in
	// Workspaces using this Network.
	SubnetIds []string `json:"subnet_ids,omitempty"`
	// Details information of each individual subnet, including
	// availability_zone and address_space. This field is populated during
	// workspace creation and used for WorkerEnvironment.
	Subnets []SubnetInfo `json:"subnets,omitempty"`
	// CIDR that used for routing tables and security groups. Example:
	// 10.0.0.0/16. CIDR blocks can now be inferred from instance metadata
	// during setup so theoretically it is no longer necessary to populate the
	// `vpcAddressSpace` field. But there is a unknown bug which causes errors
	// when listing existing clusters and preventing customers from creating new
	// clusters under workspace `Compute` page. This field is populated during
	// workspace creation and used for WorkerEnvironment.
	VpcAddressSpace string `json:"vpc_address_space,omitempty"`
	// The cloud-provided VPC ID.
	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsNetworkInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsNetworkInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureKeyInfo struct {
	// The Disk Encryption Set id that is used to represent the key info used
	// for Managed Disk BYOK use case
	DiskEncryptionSetId string `json:"disk_encryption_set_id,omitempty"`
	// The structure to store key access credential This is set if the Managed
	// Identity is being used to access the Azure Key Vault key.
	KeyAccessConfiguration *KeyAccessConfiguration `json:"key_access_configuration,omitempty"`
	// The name of the key in KeyVault.
	KeyName string `json:"key_name,omitempty"`
	// The base URI of the KeyVault.
	KeyVaultUri string `json:"key_vault_uri,omitempty"`
	// The tenant id where the KeyVault lives.
	TenantId string `json:"tenant_id,omitempty"`
	// The current key version.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AzureWorkspaceInfo struct {
	// Azure Resource Group name
	ResourceGroup string `json:"resource_group,omitempty"`
	// Azure Subscription ID
	SubscriptionId string `json:"subscription_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureWorkspaceInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureWorkspaceInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CloudResourceContainer struct {
	Gcp *CustomerFacingGcpCloudResourceContainer `json:"gcp,omitempty"`
}

type CreateAwsKeyInfo struct {
	// The alias name of the KMS key.
	KeyAlias string `json:"key_alias,omitempty"`
	// The the arn of the KMS key.
	KeyArn string `json:"key_arn"`
	// The region of the KMS key.
	KeyRegion string `json:"key_region,omitempty"`
	// Indicates if the key should be used for cluster volumes. Can only be set
	// if the CMK can be used as a data plane key (use case storage)
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAwsKeyInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCredentialAwsCredentials struct {
	StsRole *CreateCredentialStsRole `json:"sts_role,omitempty"`
}

type CreateCredentialRequest struct {
	AwsCredentials *CreateCredentialAwsCredentials `json:"aws_credentials,omitempty"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * Use Amazon's STS service to assume a specified IAM role. The
// `longLivedProvider` is required to grant permission to assume `roleArn`. As
// an example, consider the vault creating the vpc in the customer account. The
// customer may provide her credentials as a role that we can assume. To create
// the VPC, the vault will use the "sts:AssumeRole" permission in its IAM role
// to assume the customer role. In this case, the vault's role is the long lived
// provider. @param roleArn The role to assume @param externalId An identifier
// that enables cross account role assumption @param longLivedProvider The
// credentials with which to assume the role
type CreateCredentialStsRole struct {
	// Note: This must match the external_id on the parent object.
	//
	// TODO(j): Add validation to ensure this cannot be updated. If the user can
	// override the external_id, that defeats the purpose.
	ExternalId string `json:"external_id,omitempty"`

	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCredentialStsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCredentialStsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCustomerManagedKeyRequest struct {
	AwsKeyInfo *CreateAwsKeyInfo `json:"aws_key_info,omitempty"`

	GcpKeyInfo *CreateGcpKeyInfo `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	UseCases []KeyUseCase `json:"use_cases,omitempty"`
}

type CreateGcpKeyInfo struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount *GcpServiceAccount `json:"gcp_service_account,omitempty"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId string `json:"kms_key_id"`
}

type CreateNetworkRequest struct {
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateNetworkRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePrivateAccessSettingsRequest struct {
	// The MWS API ID of VPC Endpoints that can access this workspace - only
	// filled if privateAccessLevel is ENDPOINT
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The level of isolation of a workspace attached to this settings object
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// The friendly user-facing name of the Private Access Settings (i.e. jake's
	// private access settings)
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Whether or not public traffic can enter this workspace. True for hybrid
	// workspaces, false otherwise.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The region in which this private access settings is valid
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePrivateAccessSettingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStorageConfigurationRequest struct {
	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`

	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateStorageConfigurationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateStorageConfigurationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The cloud info of this vpc endpoint.
	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateWorkspaceRequest struct {
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DeploymentName string `json:"deployment_name,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// Whether No Public IP is enabled for the workspace
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink. You must specify this ID if you are using [AWS PrivateLink]
	// for either front-end (user-to-workspace connection), back-end (data plane
	// to control plane connection), or both connection types. Before
	// configuring PrivateLink, read the [Databricks article about
	// PrivateLink].",
	//
	// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
	// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Credential struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`

	AwsCredentials *AwsCredentials `json:"aws_credentials,omitempty"`
	// Time in epoch milliseconds when the credential was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Databricks credential configuration ID.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Credential) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Credential) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Corresponds to compute mode defined here:
// https://src.dev.databricks.com/databricks/universe@9076536b18479afd639d1c1f9dd5a59f72215e69/-/blob/central/api/common.proto?L872
type CustomerFacingComputeMode string

const CustomerFacingComputeModeHybrid CustomerFacingComputeMode = `HYBRID`

const CustomerFacingComputeModeServerless CustomerFacingComputeMode = `SERVERLESS`

// String representation for [fmt.Print]
func (f *CustomerFacingComputeMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingComputeMode) Set(v string) error {
	switch v {
	case `HYBRID`, `SERVERLESS`:
		*f = CustomerFacingComputeMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HYBRID", "SERVERLESS"`, v)
	}
}

// Values returns all possible values for CustomerFacingComputeMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingComputeMode) Values() []CustomerFacingComputeMode {
	return []CustomerFacingComputeMode{
		CustomerFacingComputeModeHybrid,
		CustomerFacingComputeModeServerless,
	}
}

// Type always returns CustomerFacingComputeMode to satisfy [pflag.Value] interface
func (f *CustomerFacingComputeMode) Type() string {
	return "CustomerFacingComputeMode"
}

type CustomerFacingGcpCloudResourceContainer struct {
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingGcpCloudResourceContainer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingGcpCloudResourceContainer) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerFacingStorageMode string

const CustomerFacingStorageModeCustomerHosted CustomerFacingStorageMode = `CUSTOMER_HOSTED`

const CustomerFacingStorageModeDefaultStorage CustomerFacingStorageMode = `DEFAULT_STORAGE`

// String representation for [fmt.Print]
func (f *CustomerFacingStorageMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingStorageMode) Set(v string) error {
	switch v {
	case `CUSTOMER_HOSTED`, `DEFAULT_STORAGE`:
		*f = CustomerFacingStorageMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUSTOMER_HOSTED", "DEFAULT_STORAGE"`, v)
	}
}

// Values returns all possible values for CustomerFacingStorageMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingStorageMode) Values() []CustomerFacingStorageMode {
	return []CustomerFacingStorageMode{
		CustomerFacingStorageModeCustomerHosted,
		CustomerFacingStorageModeDefaultStorage,
	}
}

// Type always returns CustomerFacingStorageMode to satisfy [pflag.Value] interface
func (f *CustomerFacingStorageMode) Type() string {
	return "CustomerFacingStorageMode"
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`

	AzureKeyInfo *AzureKeyInfo `json:"azure_key_info,omitempty"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	GcpKeyInfo *GcpKeyInfo `json:"gcp_key_info,omitempty"`
	// The cases that the key can be used for.
	UseCases []KeyUseCase `json:"use_cases,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerManagedKey) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerManagedKey) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteCredentialRequest struct {
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

type DeleteEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" url:"-"`
}

type DeleteNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" url:"-"`
}

type DeletePrivateAccesRequest struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type DeleteStorageRequest struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

type DeleteVpcEndpointRequest struct {
	VpcEndpointId string `json:"-" url:"-"`
}

type DeleteWorkspaceRequest struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

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

// ErrorType and WarningType are used to represent the type of error or warning
// by NetworkHealth and NetworkWarning defined in
// central/api/accounts/accounts.proto
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

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	AuthoritativeUserEmail string `json:"authoritative_user_email,omitempty"`
	// The authoritative user full name.
	AuthoritativeUserFullName string `json:"authoritative_user_full_name,omitempty"`
	// The legal entity name for the external workspace
	CustomerName string `json:"customer_name,omitempty"`

	OptOutExternalCustomerTosWorkflow bool `json:"opt_out_external_customer_tos_workflow,omitempty"`
	// The email of the authoritative user that signed the Terms of service.
	TosAcceptedByEmail string `json:"tos_accepted_by_email,omitempty"`
	// The full name of the authoritative user that signed the Terms of service.
	TosAcceptedByFullName string `json:"tos_accepted_by_full_name,omitempty"`
	// Indicates when the Terms of service was signed. None if it has not been
	// signed.
	TosAcceptedTimestamp int64 `json:"tos_accepted_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalCustomerInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalCustomerInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The shared network config for GCP workspace. This object has common network
// configurations that are network attributions of a workspace. DEPRECATED. Use
// GkeConfig instead.
type GcpCommonNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	GkeClusterMasterIpRange string `json:"gke_cluster_master_ip_range,omitempty"`
	// The type of network connectivity of the GKE cluster.
	GkeConnectivityType GkeConfigConnectivityType `json:"gke_connectivity_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpCommonNetworkConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpCommonNetworkConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpKeyInfo struct {
	// Globally unique service account email that has access to the KMS key. The
	// service account exists within the Databricks CP project.
	GcpServiceAccount *GcpServiceAccount `json:"gcp_service_account,omitempty"`
	// Globally unique kms key resource id of the form
	// projects/testProjectId/locations/us-east4/keyRings/gcpCmkKeyRing/cryptoKeys/cmk-eastus4
	KmsKeyId string `json:"kms_key_id"`
}

// The network configuration for the workspace.
type GcpManagedNetworkConfig struct {
	// The IP range that will be used to allocate GKE cluster Pods from.
	GkeClusterPodIpRange string `json:"gke_cluster_pod_ip_range,omitempty"`
	// The IP range that will be used to allocate GKE cluster Services from.
	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`
	// The IP range which will be used to allocate GKE cluster nodes from. Note:
	// Pods, services and master IP range must be mutually exclusive.
	SubnetCidr string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpManagedNetworkConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpManagedNetworkConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpNetworkInfo struct {
	// The GCP project ID for network resources. This project is where the VPC
	// and subnet resides.
	NetworkProjectId string `json:"network_project_id"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Pod IP range. This is BYO VPC specific. DB VPC uses
	// network.getGcpManagedNetworkConfig.getGkeClusterPodIpRange
	PodIpRangeName string `json:"pod_ip_range_name"`
	// Name of the secondary range within the subnet that will be used by GKE as
	// Service IP range.
	ServiceIpRangeName string `json:"service_ip_range_name"`
	// The customer-provided Subnet ID that will be available to Clusters in
	// Workspaces using this Network.
	SubnetId string `json:"subnet_id"`

	SubnetRegion string `json:"subnet_region"`
	// The customer-provided VPC ID.
	VpcId string `json:"vpc_id"`
}

type GcpServiceAccount struct {
	ServiceAccountEmail string `json:"service_account_email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpServiceAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpServiceAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpVpcEndpointInfo struct {
	EndpointRegion string `json:"endpoint_region"`

	ProjectId string `json:"project_id"`

	PscConnectionId string `json:"psc_connection_id,omitempty"`

	PscEndpointName string `json:"psc_endpoint_name"`

	ServiceAttachmentId string `json:"service_attachment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpVpcEndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpVpcEndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCredentialRequest struct {
	// Credential configuration ID
	CredentialsId string `json:"-" url:"-"`
}

type GetEncryptionKeyRequest struct {
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" url:"-"`
}

type GetNetworkRequest struct {
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" url:"-"`
}

type GetPrivateAccesRequest struct {
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type GetStorageRequest struct {
	StorageConfigurationId string `json:"-" url:"-"`
}

type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" url:"-"`
}

type GetWorkspaceRequest struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

// The configurations of the GKE cluster used by the GCP workspace.
type GkeConfig struct {
	// The type of network connectivity of the GKE cluster.
	ConnectivityType GkeConfigConnectivityType `json:"connectivity_type,omitempty"`
	// The IP range that will be used to allocate GKE cluster master resources
	// from. This field must not be set if
	// gke_cluster_type=PUBLIC_NODE_PUBLIC_MASTER.
	MasterIpRange string `json:"master_ip_range,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GkeConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GkeConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// The credential ID that is used to access the key vault.
type KeyAccessConfiguration struct {
	CredentialId string `json:"credential_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KeyAccessConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KeyAccessConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KeyUseCase string

const KeyUseCaseManagedServices KeyUseCase = `MANAGED_SERVICES`

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

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId string `json:"account_id,omitempty"`

	AwsNetworkInfo *AwsNetworkInfo `json:"aws_network_info,omitempty"`
	// Time in epoch milliseconds when the network was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Array of error messages about the network configuration.
	ErrorMessages []NetworkHealth `json:"error_messages,omitempty"`

	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The Databricks network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string `json:"vpc_id,omitempty"`

	VpcStatus VpcStatus `json:"vpc_status,omitempty"`
	// Array of warning messages about the network configuration.
	WarningMessages []NetworkWarning `json:"warning_messages,omitempty"`
	// Workspace ID associated with this network configuration.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Network) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Network) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage string `json:"error_message,omitempty"`

	ErrorType ErrorType `json:"error_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkHealth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkHealth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay []string `json:"dataplane_relay,omitempty"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi []string `json:"rest_api,omitempty"`
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage string `json:"warning_message,omitempty"`

	WarningType WarningType `json:"warning_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkWarning) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkWarning) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Describes AWS resources allocations for NPIP shard environments. Used to
// track and delete resources during worker (and shard) environment deletion.
// Should only be used for MT NPIP shard environments currently.
type NpipVpcInfra struct {
	// Elastic IP allocation id. Example: eipalloc-0df89abd3b5a548af
	NatEipAllocationId string `json:"natEipAllocationId,omitempty"`
	// NAT gateway id. Example: nat-0ae5b2f027fe7221a
	NatGatewayId string `json:"natGatewayId,omitempty"`
	// Route table association id. Example: rtbassoc-089a9a9037542a912
	NatRouteTableAssociationId string `json:"natRouteTableAssociationId,omitempty"`
	// Route table id. Example: rtb-06118dc3003ee809b
	NatRouteTableId string `json:"natRouteTableId,omitempty"`
	// Subnet id. Example: subnet-0f6f001e243e00c10
	NatSubnetId string `json:"natSubnetId,omitempty"`
	// VPC endpoint id. Example: vpce-08f210093b4e5ecb5
	NatVpcEndpointId string `json:"natVpcEndpointId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NpipVpcInfra) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NpipVpcInfra) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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

type PrivateAccessLevel string

const PrivateAccessLevelAccount PrivateAccessLevel = `ACCOUNT`

const PrivateAccessLevelAny PrivateAccessLevel = `ANY`

const PrivateAccessLevelEndpoint PrivateAccessLevel = `ENDPOINT`

const PrivateAccessLevelUnknownAccessLevel PrivateAccessLevel = `UNKNOWN_ACCESS_LEVEL`

// String representation for [fmt.Print]
func (f *PrivateAccessLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PrivateAccessLevel) Set(v string) error {
	switch v {
	case `ACCOUNT`, `ANY`, `ENDPOINT`, `UNKNOWN_ACCESS_LEVEL`:
		*f = PrivateAccessLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACCOUNT", "ANY", "ENDPOINT", "UNKNOWN_ACCESS_LEVEL"`, v)
	}
}

// Values returns all possible values for PrivateAccessLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PrivateAccessLevel) Values() []PrivateAccessLevel {
	return []PrivateAccessLevel{
		PrivateAccessLevelAccount,
		PrivateAccessLevelAny,
		PrivateAccessLevelEndpoint,
		PrivateAccessLevelUnknownAccessLevel,
	}
}

// Type always returns PrivateAccessLevel to satisfy [pflag.Value] interface
func (f *PrivateAccessLevel) Type() string {
	return "PrivateAccessLevel"
}

// *
type PrivateAccessSettings struct {
	// The MWS Account in which the Private Access Settings exists.
	AccountId string `json:"account_id,omitempty"`
	// The MWS API ID of VPC Endpoints that can access this workspace - only
	// filled if privateAccessLevel is ENDPOINT
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The level of isolation of a workspace attached to this settings object
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// The ID in the MWS API of the Private Access Settings.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The friendly user-facing name of the Private Access Settings (i.e. jake's
	// private access settings)
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Whether or not public traffic can enter this workspace. True for hybrid
	// workspaces, false otherwise.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The region in which this private access settings is valid
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ReplacePrivateAccessSettingsRequest struct {
	CustomerFacingPrivateAccessSettings PrivateAccessSettings `json:"customer_facing_private_access_settings"`
	// The ID in the MWS API of the Private Access Settings.
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type RootBucketInfo struct {
	// Name of the bucket
	BucketName string `json:"bucket_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RootBucketInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RootBucketInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StorageConfiguration struct {
	AccountId string `json:"account_id,omitempty"`

	CreationTime int64 `json:"creation_time,omitempty"`
	// The IAM role that is used to access the workspace catalog which is
	// created during workspace creation for UC by Default. If a storage
	// configuration that has this field populated is used to create a
	// workspace, then a workspace catalog is created together with the
	// workspace. The workspace catalog shares the root bucket with internal
	// workspace storage (including DBFS root) but uses a dedicated bucket path
	// prefix.
	RoleArn string `json:"role_arn,omitempty"`

	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`

	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`

	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StorageConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * Use Amazon's STS service to assume a specified IAM role. The
// `longLivedProvider` is required to grant permission to assume `roleArn`. As
// an example, consider the vault creating the vpc in the customer account. The
// customer may provide her credentials as a role that we can assume. To create
// the VPC, the vault will use the "sts:AssumeRole" permission in its IAM role
// to assume the customer role. In this case, the vault's role is the long lived
// provider. @param roleArn The role to assume @param externalId An identifier
// that enables cross account role assumption @param longLivedProvider The
// credentials with which to assume the role
type StsRole struct {
	// Note: This must match the external_id on the parent object.
	//
	// TODO(j): Add validation to ensure this cannot be updated. If the user can
	// override the external_id, that defeats the purpose.
	ExternalId string `json:"external_id,omitempty"`

	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Describes a single subnet, which is associated with a particular AWS AZ and a
// particular address space which is a subset of the overall vpc_address_space.
type SubnetInfo struct {
	// Example: us-west-2a
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// Example: 10.0.0.0/17.
	SubnetAddressSpace string `json:"subnet_address_space,omitempty"`

	SubnetId string `json:"subnet_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SubnetInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubnetInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWorkspaceRequest struct {
	CustomerFacingWorkspace Workspace `json:"customer_facing_workspace"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`
	// A unique integer ID for the workspace
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// *
type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration. TODO
	// - This may signal an OpenAPI diff; it does not show up in the generated
	// spec
	AccountId string `json:"account_id,omitempty"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId string `json:"aws_account_id,omitempty"`
	// The ID of the Databricks [endpoint service] that this VPC endpoint is
	// connected to. For a list of endpoint service IDs for each supported AWS
	// region, see the [Databricks PrivateLink documentation].
	//
	// [Databricks PrivateLink documentation]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
	// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html
	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The cloud info of this vpc endpoint. Info for a GCP vpc endpoint.
	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint documentation].
	//
	// [AWS DescribeVpcEndpoint documentation]: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html
	State string `json:"state,omitempty"`

	UseCase EndpointUseCase `json:"use_case,omitempty"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *VpcEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s VpcEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VpcStatus string

const VpcStatusBroken VpcStatus = `BROKEN`

const VpcStatusUnattached VpcStatus = `UNATTACHED`

const VpcStatusValid VpcStatus = `VALID`

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

type Workspace struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`

	AwsRegion string `json:"aws_region,omitempty"`

	AzureWorkspaceInfo *AzureWorkspaceInfo `json:"azure_workspace_info,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// The compute mode of the workspace.
	ComputeMode CustomerFacingComputeMode `json:"compute_mode,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DeploymentName string `json:"deployment_name,omitempty"`
	// maps to external_customer_info from workspace proto this will contains
	// fields for the customers
	ExternalCustomerInfo *ExternalCustomerInfo `json:"external_customer_info,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// Whether No Public IP is enabled for the workspace
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configuration for the workspace.
	//
	// DEPRECATED. Use `network_id` instead.
	Network *WorkspaceNetwork `json:"network,omitempty"`
	// The object ID of network connectivity config.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// If this workspace is BYO VPC, then the network_id will be populated. If
	// this workspace is not BYO VPC, then the network_id will be empty.
	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`
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
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The storage mode of the workspace.
	StorageMode CustomerFacingStorageMode `json:"storage_mode,omitempty"`
	// A unique integer ID for the workspace
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`
	// The status of a workspace
	WorkspaceStatus WorkspaceStatus `json:"workspace_status,omitempty"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Workspace) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Workspace) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network configuration for workspaces.
type WorkspaceNetwork struct {
	// The shared network config for GCP workspace. This object has common
	// network configurations that are network attributions of a workspace. This
	// object is input-only.
	GcpCommonNetworkConfig *GcpCommonNetworkConfig `json:"gcp_common_network_config,omitempty"`
	// The mutually exclusive network deployment modes. The option decides which
	// network mode the workspace will use. The network config for GCP workspace
	// with Databricks managed network. This object is input-only and will not
	// be provided when listing workspaces. See go/gcp-byovpc-alpha-design for
	// interface decisions.
	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`
	// The ID of the network object, if the workspace is a BYOVPC workspace.
	// This should apply to workspaces on all clouds in internal services. In
	// accounts-rest-api, user will use workspace.network_id for input and
	// output instead. Currently (2021-06-19) the network ID is only used by
	// GCP.
	NetworkId string `json:"network_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *WorkspaceNetwork) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WorkspaceNetwork) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The different statuses of a workspace. The following represents the current
// set of valid transitions from status to status: NOT_PROVISIONED ->
// PROVISIONING -> CANCELLED PROVISIONING -> RUNNING -> FAILED -> CANCELLED
// (note that this transition is disallowed in the MultiWorkspace Project)
// RUNNING -> PROVISIONING -> BANNED -> CANCELLED FAILED -> PROVISIONING ->
// CANCELLED BANNED -> RUNNING -> CANCELLED Note that a transition from any
// state to itself is also valid. TODO(PLAT-5867): add a transition from
// CANCELLED to some other value (e.g. RECOVERING)
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
