// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deployment

// all definitions in this file are in alphabetical order

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
}

// The general workspace configurations that are specific to cloud providers.
type CloudResourceBucket struct {
	// The general workspace configurations that are specific to Google Cloud.
	Gcp *GcpProjectContainer `json:"gcp,omitempty"`
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
}

type CreateCredentialRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`

	AwsCredentials AwsCredentials `json:"aws_credentials"`
	// The human-readable name of the credential configuration object.
	CredentialsName string `json:"credentials_name"`
}

type CreateCustomerManagedKeyRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`

	AwsKeyInfo CreateAwsKeyInfo `json:"aws_key_info"`
	// The cases that the key can be used for. Include one or both of these
	// options: * `MANAGED_SERVICES`: Encrypts notebook and secret data in the
	// control plane * `STORAGE`: Encrypts the workspace's root S3 bucket (root
	// DBFS and system data) and, optionally, cluster EBS volumes.
	UseCases []string `json:"use_cases"`
}

// The network configurations for the workspace. If you provide a network
// configuration ID for a new workspace, Databricks deploys the new workspace
// into that associated customer-managed VPC. If omitted, by default Databricks
// creates a new Databricks-managed VPC for the workspace in your Google account
// and manages its lifecycle.
//
// All the IP range configurations must be mutually exclusive. An attempt to
// create a workspace fails if Databricks detects an IP range overlap.
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
// workspace](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html).
type CreateGcpNetwork struct {
	// The common network configuration fields that can be used by both
	// Databricks-managed VPCs and customer-managed VPCs.
	GcpCommonNetworkConfig *GcpCommonNetworkConfig `json:"gcp_common_network_config,omitempty"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.
	GcpManagedNetworkConfig *GcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`
	// The network configuration ID that is attached to the workspace. If you
	// provide a network configuration ID for a new workspace, Databricks
	// validates the network resources and deploys the new workspace into your
	// associated customer-managed VPC that is specified in this network
	// configuration. If omitted, by default Databricks creates a new
	// Databricks-managed VPC for the workspace in your Google account and
	// manages its lifecycle.
	NetworkId string `json:"network_id,omitempty"`
}

type CreateNetworkRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name"`
	// IDs of one to five security groups associated with this network. Security
	// group IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least two subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink/).
	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id,omitempty"`
}

type CreateStorageConfigurationRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Root S3 bucket information.
	RootBucketInfo RootBucketInfo `json:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name"`
}

type CreateVpcEndpointRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name"`
}

type CreateWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// The AWS region of the workspace's data plane.
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud string `json:"cloud,omitempty"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceBucket *CloudResourceBucket `json:"cloud_resource_bucket,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<workspace-deployment-name>.cloud.databricks.com`. For example, if the
	// deployment name is `abcsales`, your workspace URL will be
	// `https://abcsales.cloud.databricks.com`. Hyphens are allowed. This
	// property supports only the set of characters that are allowed in a
	// subdomain.
	//
	// If your account has a non-empty deployment name prefix at workspace
	// creation time, the workspace deployment name changes so that the
	// beginning has the account prefix and a hyphen. For example, if your
	// account's deployment prefix is `acme` and the workspace deployment name
	// is `workspace-1`, the `deployment_name` field becomes `acme-workspace-1`
	// and that is the value that is returned in JSON responses for the
	// `deployment_name` field. The workspace URL is
	// `acme-workspace-1.cloud.databricks.com`.
	//
	// If your account has a non-empty deployment name prefix and you set
	// `deployment_name` to the reserved keyword `EMPTY`, `deployment_name` is
	// just the account prefix only. For example, if your account's deployment
	// prefix is `acme` and the workspace deployment name is `EMPTY`,
	// `deployment_name` becomes `acme` only and the workspace URL is
	// `acme.cloud.databricks.com`.
	//
	// Contact your Databricks representatives to add an account deployment name
	// prefix to your account. If you do not have a deployment name prefix, the
	// special deployment name value `EMPTY` is invalid.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	//
	// If a new workspace omits this property, the server generates a unique
	// deployment name for you with the pattern `dbc-xxxxxxxx-xxxx`.
	DeploymentName string `json:"deployment_name,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location string `json:"location,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This is used to encrypt the workspace's notebook and secret data
	// in the control plane, in addition to Databricks SQL queries and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configurations for the workspace. If you provide a network
	// configuration ID for a new workspace, Databricks deploys the new
	// workspace into that associated customer-managed VPC. If omitted, by
	// default Databricks creates a new Databricks-managed VPC for the workspace
	// in your Google account and manages its lifecycle.
	//
	// All the IP range configurations must be mutually exclusive. An attempt to
	// create a workspace fails if Databricks detects an IP range overlap.
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
	// for a new
	// workspace](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html).
	Network *CreateGcpNetwork `json:"network,omitempty"`
	// The ID of the workspace's network configuration object. To use [AWS
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html)
	// (Public Preview), this field is required.
	NetworkId string `json:"network_id,omitempty"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing](https://databricks.com/product/aws-pricing).
	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink (Public Preview). This ID must be specified for customers
	// using [AWS PrivateLink](https://aws.amazon.com/privatelink/) for either
	// front-end (user-to-workspace connection), back-end (data plane to control
	// plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
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
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`
	// The cases that the key can be used for. Include one or both of these
	// options: * `MANAGED_SERVICES`: Encrypts notebook and secret data in the
	// control plane * `STORAGE`: Encrypts the workspace's root S3 bucket (root
	// DBFS and system data) and optionally cluster EBS volumes.
	UseCases []string `json:"use_cases,omitempty"`
}

// This enumeration represents the type of Databricks VPC [endpoint
// service](https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html)
// that was used when creating this VPC endpoint.
//
// If the VPC endpoint connects to the Databricks control plane for either the
// front-end connection or the back-end REST API connection, the value is
// `WORKSPACE_ACCESS`.
//
// If the VPC endpoint connects to the Databricks workspace for the back-end
// [secure cluster
// connectivity](https://docs.databricks.com/security/secure-cluster-connectivity.html)
// relay, the value is `DATAPLANE_RELAY_ACCESS`.
type EndpointUseCase string

const EndpointUseCaseDataplaneRelayAccess EndpointUseCase = `DATAPLANE_RELAY_ACCESS`

const EndpointUseCaseWorkspaceAccess EndpointUseCase = `WORKSPACE_ACCESS`

// The AWS resource associated with this error: credentials, VPC, subnet,
// security group, or network ACL.
type ErrorType string

const ErrorTypeCredentials ErrorType = `credentials`

const ErrorTypeNetworkacl ErrorType = `networkAcl`

const ErrorTypeSecuritygroup ErrorType = `securityGroup`

const ErrorTypeSubnet ErrorType = `subnet`

const ErrorTypeVpc ErrorType = `vpc`

// Specifies the network connectivity types for the GKE nodes and the GKE master
// network. Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for
// the workspace. The GKE nodes will not have public IPs. Set to
// `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a public
// GKE cluster have public IP addresses.
type GkeConnectivityType string

const GkeConnectivityTypePrivateNodePublicMaster GkeConnectivityType = `PRIVATE_NODE_PUBLIC_MASTER`

const GkeConnectivityTypePublicNodePublicMaster GkeConnectivityType = `PUBLIC_NODE_PUBLIC_MASTER`

// The common network configuration fields that can be used by both
// Databricks-managed VPCs and customer-managed VPCs.
type GcpCommonNetworkConfig struct {
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled.
	//
	// It must be exactly as big as `/28`.
	GkeClusterMasterIpRange string `json:"gke_cluster_master_ip_range,omitempty"`
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network. Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE
	// cluster for the workspace. The GKE nodes will not have public IPs. Set to
	// `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a
	// public GKE cluster have public IP addresses.
	GkeConnectivityType GkeConnectivityType `json:"gke_connectivity_type,omitempty"`
}

// The network settings for the workspace. The configurations are only for
// Databricks-managed VPCs. It is ignored if you specify a customer-managed VPC
// in the `network_id` field.
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
}

type GcpNetwork struct {
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId string `json:"network_id,omitempty"`
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
	SubnetId string `json:"subnet_id,omitempty"`
	// The Google Cloud region of the workspace data plane (for example,
	// `us-east4`).
	SubnetRegion string `json:"subnet_region"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id"`
}

// The general workspace configurations that are specific to Google Cloud.
type GcpProjectContainer struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId string `json:"project_id,omitempty"`
}

type KeyStatus string

const KeyStatusAttached KeyStatus = `ATTACHED`

const KeyStatusDetached KeyStatus = `DETACHED`

const KeyStatusUnknown KeyStatus = `UNKNOWN`

// This describes an enum
type KeyUseCase string

// Encrypts notebook and secret data in the control plane
const KeyUseCaseManagedServices KeyUseCase = `MANAGED_SERVICES`

// Encrypts the workspace's root S3 bucket (root DBFS and system data) and,
// optionally, cluster EBS volumes.
const KeyUseCaseStorage KeyUseCase = `STORAGE`

type ListWorkspaceEncryptionKeyRecordsResponse struct {
	WorkspaceEncryptionKeyRecords []WorkspaceEncryptionKeyRecord `json:"workspaceEncryptionKeyRecords,omitempty"`
}

type Network struct {
	// The Databricks account ID associated with this network configuration.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the network was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Array of error messages about the network configuration.
	ErrorMessages []NetworkHealth `json:"error_messages,omitempty"`
	// The Google Cloud specific information for this network (for example, the
	// VPC ID, subnet ID, and secondary IP ranges).
	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The Databricks network configuration ID.
	NetworkId string `json:"network_id,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name,omitempty"`

	SecurityGroupIds []string `json:"security_group_ids,omitempty"`

	SubnetIds []string `json:"subnet_ids,omitempty"`
	// If specified, contains the VPC endpoints used to allow cluster
	// communication from this VPC over [AWS
	// PrivateLink](https://aws.amazon.com/privatelink/).
	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network configuration. VPC IDs can
	// be used in multiple networks.
	VpcId string `json:"vpc_id,omitempty"`
	// This describes an enum
	VpcStatus VpcStatus `json:"vpc_status,omitempty"`
	// Array of warning messages about the network configuration.
	WarningMessages []NetworkWarning `json:"warning_messages,omitempty"`
	// Workspace ID associated with this network configuration.
	WorkspaceId int64 `json:"workspace_id,omitempty"`
}

type NetworkHealth struct {
	// Details of the error.
	ErrorMessage string `json:"error_message,omitempty"`
	// The AWS resource associated with this error: credentials, VPC, subnet,
	// security group, or network ACL.
	ErrorType ErrorType `json:"error_type,omitempty"`
}

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink](https://aws.amazon.com/privatelink/).
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this network to access the Databricks secure
	// cluster connectivity relay. See [Secure Cluster
	// Connectivity](https://docs.databricks.com/security/secure-cluster-connectivity.html).
	//
	// This is a list type for future compatibility, but currently only one VPC
	// endpoint ID should be supplied.
	//
	// **Note**: This is the Databricks-specific ID of the VPC endpoint object
	// in the Account API, not the AWS VPC endpoint ID that you see for your
	// endpoint in the AWS Console.
	DataplaneRelay []string `json:"dataplane_relay"`
	// The VPC endpoint ID used by this network to access the Databricks REST
	// API. Databricks clusters make calls to our REST API as part of cluster
	// creation, mlflow tracking, and many other features. Thus, this is
	// required even if your workspace allows public access to the REST API.
	//
	// This is a list type for future compatibility, but currently only one VPC
	// endpoint ID should be supplied.
	//
	// **Note**: This is the Databricks-specific ID of the VPC endpoint object
	// in the Account API, not the AWS VPC endpoint ID that you see for your
	// endpoint in the AWS Console.
	RestApi []string `json:"rest_api"`
}

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage string `json:"warning_message,omitempty"`
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType WarningType `json:"warning_type,omitempty"`
}

// The pricing tier of the workspace. For pricing tier information, see [AWS
// Pricing](https://databricks.com/product/aws-pricing).
type PricingTier string

const PricingTierCommunityEdition PricingTier = `COMMUNITY_EDITION`

const PricingTierDedicated PricingTier = `DEDICATED`

const PricingTierEnterprise PricingTier = `ENTERPRISE`

const PricingTierPremium PricingTier = `PREMIUM`

const PricingTierStandard PricingTier = `STANDARD`

const PricingTierUnknown PricingTier = `UNKNOWN`

// The private access level controls which VPC endpoints can connect to the UI
// or API of any workspace that attaches this private access settings object. *
// `ANY` (deprecated): Any VPC endpoint can connect to your workspace. *
// `ACCOUNT` level access (the default) allows only VPC endpoints that are
// registered in your Databricks account connect to your workspace. * `ENDPOINT`
// level access allows only specified VPC endpoints connect to your workspace.
// For details, see `allowed_vpc_endpoint_ids`.
type PrivateAccessLevel string

const PrivateAccessLevelAccount PrivateAccessLevel = `ACCOUNT`

const PrivateAccessLevelAny PrivateAccessLevel = `ANY`

const PrivateAccessLevelEndpoint PrivateAccessLevel = `ENDPOINT`

type PrivateAccessSettings struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// An array of Databricks VPC endpoint IDs. This is the Databricks ID
	// returned when registering the VPC endpoint configuration in your
	// Databricks account. This is _not_ the ID of the VPC endpoint in AWS.
	//
	// Only used when `private_access_level` is set to `ENDPOINT`. This is an
	// allow list of VPC endpoints registered in your Databricks account that
	// can connect to your workspace over AWS PrivateLink.
	//
	// **Note**: If hybrid access to your workspace is enabled by setting
	// `public_access_enabled` to `true`, this control only works for
	// PrivateLink connections. To control how your workspace is accessed via
	// public internet, see [IP access
	// lists](https://docs.databricks.com/security/network/ip-access-list.html).
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ANY` (deprecated): Any VPC endpoint can connect to your
	// workspace. * `ACCOUNT` level access (the default) allows only VPC
	// endpoints that are registered in your Databricks account connect to your
	// workspace. * `ENDPOINT` level access allows only specified VPC endpoints
	// connect to your workspace. For details, see `allowed_vpc_endpoint_ids`.
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
	// The AWS region for workspaces attached to this private access settings
	// object.
	Region string `json:"region,omitempty"`
}

// Root S3 bucket information.
type RootBucketInfo struct {
	// The name of the S3 bucket.
	BucketName string `json:"bucket_name,omitempty"`
}

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Root S3 bucket information.
	RootBucketInfo *RootBucketInfo `json:"root_bucket_info,omitempty"`
	// Databricks storage configuration ID.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`
}

type StsRole struct {
	// The external ID that needs to be trusted by the cross-account role. This
	// is always your Databricks account ID.
	ExternalId string `json:"external_id,omitempty"`
	// The Amazon Resource Name (ARN) of the cross account role.
	RoleArn string `json:"role_arn,omitempty"`
}

type UpdateWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// The AWS region of the workspace's data plane (for example, `us-west-2`).
	// This parameter is available only for updating failed workspaces.
	AwsRegion string `json:"aws_region,omitempty"`
	// ID of the workspace's credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The ID of the workspace's managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The ID of the workspace's network configuration object. Used only if you
	// already use a customer-managed VPC. This change is supported only if you
	// specified a network configuration ID when the workspace was created. In
	// other words, you cannot switch from a Databricks-managed VPC to a
	// customer-managed VPC. This parameter is available for updating both
	// failed and running workspaces. **Note**: You cannot use a network
	// configuration update in this API to add support for PrivateLink (Public
	// Preview). To add PrivateLink to an existing workspace, contact your
	// Databricks representative.
	NetworkId string `json:"network_id,omitempty"`
	// The ID of the workspace's storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

type UpsertPrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
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
	// public internet, see [IP access
	// lists](https://docs.databricks.com/security/network/ip-access-list.html).
	AllowedVpcEndpointIds []string `json:"allowed_vpc_endpoint_ids,omitempty"`
	// The private access level controls which VPC endpoints can connect to the
	// UI or API of any workspace that attaches this private access settings
	// object. * `ANY` (deprecated): Any VPC endpoint can connect to your
	// workspace. * `ACCOUNT` level access (the default) allows only VPC
	// endpoints that are registered in your Databricks account connect to your
	// workspace. * `ENDPOINT` level access allows only specified VPC endpoints
	// connect to your workspace. For details, see `allowed_vpc_endpoint_ids`.
	PrivateAccessLevel PrivateAccessLevel `json:"private_access_level,omitempty"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" path:"private_access_settings_id"`
	// The human-readable name of the private access settings object.
	PrivateAccessSettingsName string `json:"private_access_settings_name"`
	// Determines if the workspace can be accessed over public internet. For
	// fully private workspaces, you can optionally specify `false`, but only if
	// you implement both the front-end and the back-end PrivateLink
	// connections. Otherwise, specify `true`, which means that public access is
	// enabled.
	PublicAccessEnabled bool `json:"public_access_enabled,omitempty"`
	// The AWS region for workspaces associated with this private access
	// settings object. This must be a [region that Databricks supports for
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/regions.html).
	Region string `json:"region"`
}

type VpcEndpoint struct {
	// The Databricks account ID that hosts the VPC endpoint configuration.
	AccountId string `json:"account_id,omitempty"`
	// The AWS Account in which the VPC endpoint object exists.
	AwsAccountId string `json:"aws_account_id,omitempty"`
	// The ID of the Databricks [endpoint
	// service](https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html)
	// that this VPC endpoint is connected to. For a list of endpoint service
	// IDs for each supported AWS region, see the [Databricks PrivateLink
	// documentation](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	AwsEndpointServiceId string `json:"aws_endpoint_service_id,omitempty"`
	// The ID of the VPC endpoint object in AWS.
	AwsVpcEndpointId string `json:"aws_vpc_endpoint_id,omitempty"`
	// The AWS region in which this VPC endpoint object exists.
	Region string `json:"region,omitempty"`
	// The current state (such as `available` or `rejected`) of the VPC
	// endpoint. Derived from AWS. For the full set of values, see [AWS
	// DescribeVpcEndpoint
	// documentation](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-vpc-endpoints.html).
	State string `json:"state,omitempty"`
	// This enumeration represents the type of Databricks VPC [endpoint
	// service](https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html)
	// that was used when creating this VPC endpoint.
	//
	// If the VPC endpoint connects to the Databricks control plane for either
	// the front-end connection or the back-end REST API connection, the value
	// is `WORKSPACE_ACCESS`.
	//
	// If the VPC endpoint connects to the Databricks workspace for the back-end
	// [secure cluster
	// connectivity](https://docs.databricks.com/security/secure-cluster-connectivity.html)
	// relay, the value is `DATAPLANE_RELAY_ACCESS`.
	UseCase EndpointUseCase `json:"use_case,omitempty"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`
}

// This describes an enum
type VpcStatus string

// Broken.
const VpcStatusBroken VpcStatus = `BROKEN`

// Unattached.
const VpcStatusUnattached VpcStatus = `UNATTACHED`

// Valid.
const VpcStatusValid VpcStatus = `VALID`

// Warned.
const VpcStatusWarned VpcStatus = `WARNED`

// The AWS resource associated with this warning: a subnet or a security group.
type WarningType string

const WarningTypeSecuritygroup WarningType = `securityGroup`

const WarningTypeSubnet WarningType = `subnet`

type Workspace struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// The AWS region of the workspace data plane (for example, `us-west-2`).
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`
	// The general workspace configurations that are specific to cloud
	// providers.
	CloudResourceBucket *CloudResourceBucket `json:"cloud_resource_bucket,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace's credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `<deployment-name>.cloud.databricks.com`.
	//
	// This value must be unique across all non-deleted deployments across all
	// AWS regions.
	DeploymentName string `json:"deployment_name,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account (for example, `us-east4`).
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	Network *GcpNetwork `json:"network,omitempty"`
	// The pricing tier of the workspace. For pricing tier information, see [AWS
	// Pricing](https://databricks.com/product/aws-pricing).
	PricingTier PricingTier `json:"pricing_tier,omitempty"`
	// ID of the workspace's private access settings object. Only used for
	// PrivateLink (Public Preview). You must specify this ID if you are using
	// [AWS PrivateLink](https://aws.amazon.com/privatelink/) for either
	// front-end (user-to-workspace connection), back-end (data plane to control
	// plane connection), or both connection types.
	//
	// Before configuring PrivateLink, read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace's storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`
	// The status of the workspace. For workspace creation, usually it is set to
	// `PROVISIONING` initially. Continue to check the status until the status
	// is `RUNNING`.
	WorkspaceStatus WorkspaceStatus `json:"workspace_status,omitempty"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`
}

type WorkspaceEncryptionKeyRecord struct {
	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	KeyStatus KeyStatus `json:"key_status,omitempty"`
	// ID for the workspace-key mapping record.
	RecordId string `json:"record_id,omitempty"`
	// Time in epoch milliseconds when the record was added.
	UpdateTime int64 `json:"update_time,omitempty"`
	// This describes an enum
	UseCase KeyUseCase `json:"use_case,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"workspace_id,omitempty"`
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

type DeleteCredentialConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" path:"credentials_id"`
}

type DeleteKeyConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" path:"customer_managed_key_id"`
}

type DeleteNetworkConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" path:"network_id"`
}

type DeletePrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" path:"private_access_settings_id"`
}

type DeleteStorageConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" path:"storage_configuration_id"`
}

type DeleteVpcEndpointRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" path:"vpc_endpoint_id"`
}

type DeleteWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

type GetAllKeyConfigsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetAllNetworkConfigsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetAllPrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetAllStorageConfigsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetAllVpcEndpointsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetAllWorkspacesRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetCredentialConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API credential configuration ID
	CredentialsId string `json:"-" path:"credentials_id"`
}

type GetKeyConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" path:"customer_managed_key_id"`
}

type GetKeyWorkspaceHistoryRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetNetworkConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API network configuration ID.
	NetworkId string `json:"-" path:"network_id"`
}

type GetPrivateAccessSettingsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API private access settings ID.
	PrivateAccessSettingsId string `json:"-" path:"private_access_settings_id"`
}

type GetStorageConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string `json:"-" path:"storage_configuration_id"`
}

type GetVpcEndpointRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" path:"vpc_endpoint_id"`
}

type GetWorkspaceKeyHistoryRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

type GetWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Workspace ID.
	WorkspaceId int64 `json:"-" path:"workspace_id"`
}

type ListCredentialsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}
