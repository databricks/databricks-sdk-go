// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AwsCredentials struct {
	StsRole *StsRole `json:"sts_role,omitempty"`
}

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn string `json:"key_arn"`
	// The AWS KMS key region.
	KeyRegion string `json:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsKeyInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsKeyInfo) MarshalJSON() ([]byte, error) {
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

// The general workspace configurations that are specific to cloud providers.
type CloudResourceContainer struct {
	Gcp *CustomerFacingGcpCloudResourceContainer `json:"gcp,omitempty"`
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn string `json:"key_arn"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
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
	AwsCredentials CreateCredentialAwsCredentials `json:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name"`
}

type CreateCredentialStsRole struct {
	// The Amazon Resource Name (ARN) of the cross account role.
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
	UseCases []KeyUseCase `json:"use_cases"`
}

type CreateGcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId string `json:"kms_key_id"`
}

type CreateNetworkRequest struct {
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateNetworkRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateNetworkRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateStorageConfigurationRequest struct {
	RootBucketInfo RootBucketInfo `json:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name"`
}

type CreateVpcEndpointRequest struct {
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`

	GcpVpcEndpointInfo *GcpVpcEndpointInfo `json:"gcp_vpc_endpoint_info,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateVpcEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVpcEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane.
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
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
	DeploymentName string `json:"deployment_name,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location string `json:"location,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to help protect and control access to the
	// workspace's notebooks, secrets, Databricks SQL queries, and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkId string `json:"network_id,omitempty"`

	PricingTier PricingTier `json:"pricing_tier,omitempty"`
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
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the workspace's storage encryption key configuration object.
	// This is used to encrypt the workspace's root S3 bucket (root DBFS and
	// system data) and, optionally, cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The workspace's human-readable name.
	WorkspaceName string `json:"workspace_name"`

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

// The general workspace configurations that are specific to Google Cloud.
type CustomerFacingGcpCloudResourceContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId string `json:"project_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomerFacingGcpCloudResourceContainer) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomerFacingGcpCloudResourceContainer) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`
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
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type DeleteStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" url:"-"`
}

type DeleteVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" url:"-"`
}

type DeleteWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
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

type ExternalCustomerInfo struct {
	// Email of the authoritative user.
	AuthoritativeUserEmail string `json:"authoritative_user_email,omitempty"`
	// The authoritative user full name.
	AuthoritativeUserFullName string `json:"authoritative_user_full_name,omitempty"`
	// The legal entity name for the external workspace
	CustomerName string `json:"customer_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExternalCustomerInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalCustomerInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GcpKeyInfo struct {
	// The GCP KMS key's resource name
	KmsKeyId string `json:"kms_key_id"`
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
	GkeClusterPodIpRange string `json:"gke_cluster_pod_ip_range,omitempty"`
	// The IP range from which to allocate GKE cluster services. No bigger than
	// `/16` and no smaller than `/27`.
	GkeClusterServiceIpRange string `json:"gke_cluster_service_ip_range,omitempty"`
	// The IP range from which to allocate GKE cluster nodes. No bigger than
	// `/9` and no smaller than `/29`.
	SubnetCidr string `json:"subnet_cidr,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpManagedNetworkConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpManagedNetworkConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The Google Cloud specific information for this network (for example, the VPC
// ID, subnet ID, and secondary IP ranges).
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	NetworkProjectId string `json:"network_project_id"`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can be
	// used by only one workspace.
	PodIpRangeName string `json:"pod_ip_range_name"`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// be used by only one workspace.
	ServiceIpRangeName string `json:"service_ip_range_name"`
	// The ID of the subnet associated with this network.
	SubnetId string `json:"subnet_id"`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion string `json:"subnet_region"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id"`
}

// The Google Cloud specific information for this Private Service Connect
// endpoint.
type GcpVpcEndpointInfo struct {
	// Region of the PSC endpoint.
	EndpointRegion string `json:"endpoint_region"`
	// The Google Cloud project ID of the VPC network where the PSC connection
	// resides.
	ProjectId string `json:"project_id"`
	// The unique ID of this PSC connection.
	PscConnectionId string `json:"psc_connection_id,omitempty"`
	// The name of the PSC endpoint in the Google Cloud project.
	PscEndpointName string `json:"psc_endpoint_name"`
	// The service attachment this PSC connection connects to.
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
	// Databricks Account API credential configuration ID
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
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
}

type GetStorageRequest struct {
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" url:"-"`
}

type GetVpcEndpointRequest struct {
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" url:"-"`
}

type GetWorkspaceRequest struct {
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`
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
	ConnectivityType GkeConfigConnectivityType `json:"connectivity_type,omitempty"`
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
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

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the network was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Array of error messages about the network configuration.
	ErrorMessages []NetworkHealth `json:"error_messages,omitempty"`

	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The Databricks network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`

	SecurityGroupIds []string `json:"security_group_ids,omitempty"`

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

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink/
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay.
	DataplaneRelay []string `json:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API.
	RestApi []string `json:"rest_api"`
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

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// An array of Databricks VPC endpoint IDs.
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`

	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks private access settings ID.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name,omitempty"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The cloud region for workspaces attached to this private access settings
	// object.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PrivateAccessSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PrivateAccessSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
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
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`

	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`
	// Databricks storage configuration ID.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StorageConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StorageConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string `json:"role_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StsRole) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StsRole) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateWorkspaceRequest struct {
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	AwsRegion string `json:"aws_region,omitempty"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. For failed workspaces only, you can
	// switch from a Databricks-managed VPC to a customer-managed VPC by
	// updating the workspace to add a network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The ID of the workspace's private access settings configuration object.
	// This parameter is available only for updating failed workspaces.
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateWorkspaceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`

	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" url:"-"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The cloud region for workspaces associated with this private access
	// settings object.
	Region string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpsertPrivateAccessSettingsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpsertPrivateAccessSettingsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
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

type Workspace struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion string `json:"aws_region,omitempty"`

	AzureWorkspaceInfo *AzureWorkspaceInfo `json:"azure_workspace_info,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceContainer *CloudResourceContainer `json:"cloud_resource_container,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The custom tags key-value pairing that is attached to this workspace. The
	// key-value pair is a string of utf-8 characters. The value can be an empty
	// string, with maximum length of 255 characters. The key can be of maximum
	// length of 127 characters, and cannot be empty.
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName string `json:"deployment_name,omitempty"`
	// If this workspace is for a external customer, then external_customer_info
	// is populated. If this workspace is not for a external customer, then
	// external_customer_info is empty.
	ExternalCustomerInfo *ExternalCustomerInfo `json:"external_customer_info,omitempty"`

	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`

	GkeConfig *GkeConfig `json:"gke_config,omitempty"`
	// Whether no public IP is enabled for the workspace.
	IsNoPublicIpEnabled bool `json:"is_no_public_ip_enabled,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
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
	// A unique integer ID for the workspace
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`

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
