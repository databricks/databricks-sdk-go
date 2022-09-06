// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaces

// all definitions in this file are in alphabetical order
// The general workspace configurations that are specific to cloud providers.
type CloudResourceBucket struct {
	// The general workspace configurations that are specific to Google Cloud.
	Gcp *CloudResourceBucketGcp `json:"gcp,omitempty"`
}

// The general workspace configurations that are specific to Google Cloud.
type CloudResourceBucketGcp struct {
	// The Google Cloud project ID, which the workspace uses to instantiate
	// cloud resources for your workspace.
	ProjectId string `json:"project_id,omitempty"`
}

type CreateWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// The AWS region of the workspace&#39;s Data Plane.
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud provider which the workspace uses. For Google Cloud workspaces,
	// always set this field to `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceBucket *CloudResourceBucket `json:"cloud_resource_bucket,omitempty"`
	// ID of the workspace&#39;s credential configuration object
	CredentialsId string `json:"credentials_id,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `&lt;workspace-deployment-name&gt;.cloud.databricks.com`. For example, if the
	// deployment name is `abcsales`, your workspace URL will be
	// `https://abcsales.cloud.databricks.com`. Hyphens are allowed. This
	// property supports only the set of characters that are allowed in a
	// subdomain. If your account has a non-empty deployment name prefix at
	// workspace creation time, the workspace deployment name changes so that
	// the beginning has the account prefix and a hyphen. For example, if your
	// account&#39;s deployment prefix is `acme` and the workspace deployment name
	// is `workspace-1`, the `deployment_name` field becomes `acme-workspace-1`
	// and that is the value that will be returned in JSON responses for the
	// `deployment_name` field. The workspace URL is
	// `acme-workspace-1.cloud.databricks.com`. If your account has a non-empty
	// deployment name prefix and you set `deployment_name` to the reserved
	// keyword `EMPTY`, `deployment_name` is just the account prefix only. For
	// example, if your account&#39;s deployment prefix is `acme` and the workspace
	// deployment name is `EMPTY`, `deployment_name` becomes `acme` only and the
	// workspace URL is `acme.cloud.databricks.com`. Contact your Databricks
	// representatives to add an account deployment name prefix to your account.
	// If you do not have a deployment name prefix, the special deployment name
	// value `EMPTY` is invalid. This value must be unique across all
	// non-deleted deployments across all AWS regions. If a new workspace omits
	// this property, the server generates a unique deployment name for you with
	// the pattern `dbc-xxxxxxxx-xxxx`.
	DeploymentName string `json:"deployment_name,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location string `json:"location,omitempty"`
	// The ID of the workspace&#39;s managed services encryption key configuration
	// object. This is used to encrypt the workspace&#39;s notebook and secret data
	// in the control plane, as well as Databricks SQL queries and query
	// history. The provided key configuration object property `use_cases` must
	// contain `MANAGED_SERVICES`.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The network configurations for the workspace. If you provide a network
	// configuration ID for a new workspace, Databricks deploys the new
	// workspace into that associated customer-managed VPC. If omitted, by
	// default Databricks creates a new Databricks-managed VPC for the workspace
	// in your Google account and manages its lifecycle. All the IP range
	// configurations must be mutually exclusive. An attempt to create a
	// workspace fails if Databricks detects an IP range overlap. Specify custom
	// IP ranges in CIDR format. The IP ranges for these fields must not
	// overlap, and all IP addresses must be entirely within the following
	// ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`, `192.168.0.0/16`,
	// and `240.0.0.0/4`. The sizes of these IP ranges affect the maximum number
	// of nodes for the workspace. **Important**: Confirm the IP ranges used by
	// your Databricks workspace before creating the workspace. You cannot
	// change them after your workspace is deployed. If the IP address ranges
	// for your Databricks are too small, IP exhaustion can occur, causing your
	// Databricks jobs to fail. To determine the address range sizes that you
	// need, Databricks provides a calculator as a Microsoft Excel spreadsheet.
	// See [calculate subnet sizes for a new
	// workspace](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html).
	Network *CreateWorkspaceRequestNetwork `json:"network,omitempty"`
	// The ID of the workspace&#39;s network configuration object. To use [AWS
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html)
	// (Public Preview), this field is required.
	NetworkId string `json:"network_id,omitempty"`
	// The pricing tier of the workspace. If you do not provide this, the API
	// will default to the highest pricing tier available to your account. See
	// https://databricks.com/product/aws-pricing for available pricing tier
	// information.
	PricingTier CreateWorkspaceRequestPricingTier `json:"pricing_tier,omitempty"`
	// Only used for PrivateLink, which is in Public Preview. This is the ID of
	// the workspace&#39;s private access settings object. This ID must be specified
	// for customers using [AWS
	// PrivateLink](https://aws.amazon.com/privatelink/) for either front-end
	// (user-to-workspace connection), back-end (data plane to control plane
	// connection), or both connection types. Before configuring PrivateLink, it
	// is important to read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// The ID of the workspace&#39;s storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the workspace&#39;s storage encryption key configuration object.
	// This is used to encrypt the workspace&#39;s root S3 bucket (root DBFS and
	// system data) and optionally cluster EBS volumes. The provided key
	// configuration object property `use_cases` must contain `STORAGE`.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// The workspace&#39;s human-readable name.
	WorkspaceName string `json:"workspace_name"`
}

// The network configurations for the workspace. If you provide a network
// configuration ID for a new workspace, Databricks deploys the new workspace
// into that associated customer-managed VPC. If omitted, by default Databricks
// creates a new Databricks-managed VPC for the workspace in your Google account
// and manages its lifecycle. All the IP range configurations must be mutually
// exclusive. An attempt to create a workspace fails if Databricks detects an IP
// range overlap. Specify custom IP ranges in CIDR format. The IP ranges for
// these fields must not overlap, and all IP addresses must be entirely within
// the following ranges: `10.0.0.0/8`, `100.64.0.0/10`, `172.16.0.0/12`,
// `192.168.0.0/16`, and `240.0.0.0/4`. The sizes of these IP ranges affect the
// maximum number of nodes for the workspace. **Important**: Confirm the IP
// ranges used by your Databricks workspace before creating the workspace. You
// cannot change them after your workspace is deployed. If the IP address ranges
// for your Databricks are too small, IP exhaustion can occur, causing your
// Databricks jobs to fail. To determine the address range sizes that you need,
// Databricks provides a calculator as a Microsoft Excel spreadsheet. See
// [calculate subnet sizes for a new
// workspace](https://docs.gcp.databricks.com/administration-guide/cloud-configurations/gcp/network-sizing.html).
type CreateWorkspaceRequestNetwork struct {
	// The common network configuration fields that can be used by both
	// Databricks-managed VPCs and customer-managed VPCs.
	GcpCommonNetworkConfig *CreateWorkspaceRequestNetworkGcpCommonNetworkConfig `json:"gcp_common_network_config,omitempty"`
	// The network settings for the workspace. The configurations are only for
	// Databricks-managed VPCs. It is ignored if you specify a customer-managed
	// VPC in the `network_id` field.
	GcpManagedNetworkConfig *CreateWorkspaceRequestNetworkGcpManagedNetworkConfig `json:"gcp_managed_network_config,omitempty"`
	// The network configuration ID that is attached to the workspace. If you
	// provide a network configuration ID for a new workspace, Databricks
	// validates the network resources and deploys the new workspace into your
	// associated customer-managed VPC that is specified in this network
	// configuration. If omitted, by default Databricks creates a new
	// Databricks-managed VPC for the workspace in your Google account and
	// manages its lifecycle.
	NetworkId string `json:"network_id,omitempty"`
}

// The common network configuration fields that can be used by both
// Databricks-managed VPCs and customer-managed VPCs.
type CreateWorkspaceRequestNetworkGcpCommonNetworkConfig struct {
	// The IP range from which to allocate GKE cluster master resources. This
	// field will be ignored if GKE private cluster is not enabled. It must be
	// exactly as big as `/28`.
	GkeClusterMasterIpRange string `json:"gke_cluster_master_ip_range,omitempty"`
	// Specifies the network connectivity types for the GKE nodes and the GKE
	// master network. Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE
	// cluster for the workspace. The GKE nodes will not have public IPs. Set to
	// `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a
	// public GKE cluster have public IP addresses.
	GkeConnectivityType CreateWorkspaceRequestNetworkGcpCommonNetworkConfigGkeConnectivityType `json:"gke_connectivity_type,omitempty"`
}

// Specifies the network connectivity types for the GKE nodes and the GKE master
// network. Set to `PRIVATE_NODE_PUBLIC_MASTER` for a private GKE cluster for
// the workspace. The GKE nodes will not have public IPs. Set to
// `PUBLIC_NODE_PUBLIC_MASTER` for a public GKE cluster. The nodes of a public
// GKE cluster have public IP addresses.
type CreateWorkspaceRequestNetworkGcpCommonNetworkConfigGkeConnectivityType string

const CreateWorkspaceRequestNetworkGcpCommonNetworkConfigGkeConnectivityTypePrivateNodePublicMaster CreateWorkspaceRequestNetworkGcpCommonNetworkConfigGkeConnectivityType = `PRIVATE_NODE_PUBLIC_MASTER`

const CreateWorkspaceRequestNetworkGcpCommonNetworkConfigGkeConnectivityTypePublicNodePublicMaster CreateWorkspaceRequestNetworkGcpCommonNetworkConfigGkeConnectivityType = `PUBLIC_NODE_PUBLIC_MASTER`

// The network settings for the workspace. The configurations are only for
// Databricks-managed VPCs. It is ignored if you specify a customer-managed VPC
// in the `network_id` field.
type CreateWorkspaceRequestNetworkGcpManagedNetworkConfig struct {
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

// The pricing tier of the workspace. If you do not provide this, the API will
// default to the highest pricing tier available to your account. See
// https://databricks.com/product/aws-pricing for available pricing tier
// information.
type CreateWorkspaceRequestPricingTier string

const CreateWorkspaceRequestPricingTierEnterprise CreateWorkspaceRequestPricingTier = `ENTERPRISE`

const CreateWorkspaceRequestPricingTierPremium CreateWorkspaceRequestPricingTier = `PREMIUM`

const CreateWorkspaceRequestPricingTierStandard CreateWorkspaceRequestPricingTier = `STANDARD`

type DeleteWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Workspace ID.
	WorkspaceId int64 ` path:"workspace_id"`
}

type GetAllWorkspacesRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
}

type GetWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Workspace ID.
	WorkspaceId int64 ` path:"workspace_id"`
}

// An array of workspaces.
type ListWorkspacesResponse []Workspace

type UpdateWorkspaceRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// The AWS region of the workspace&#39;s Data Plane. For example, `us-west-2`.
	// This parameter is available only for updating failed workspaces.
	AwsRegion string `json:"aws_region,omitempty"`
	// ID of the workspace&#39;s credential configuration object. This parameter is
	// available for updating both failed and running workspaces.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The ID of the workspace&#39;s managed services encryption key configuration
	// object. This parameter is available only for updating failed workspaces.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`
	// The ID of the workspace&#39;s network configuration object. Used only if you
	// already use a customer-managed VPC. This change is supported only if you
	// specified a network configuration ID when the workspace was created. In
	// other words, you cannot switch from a Databricks-managed VPC to a
	// customer-managed VPC. This parameter is available for updating both
	// failed and running workspaces. Note: You cannot use a network
	// configuration update in this API to add support for PrivateLink (in
	// Public Preview). To add PrivateLink to an existing workspace, contact
	// your Databricks representative.
	NetworkId string `json:"network_id,omitempty"`
	// The ID of the workspace&#39;s storage configuration object. This parameter is
	// available only for updating failed workspaces.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The ID of the key configuration object for workspace storage. This
	// parameter is available for updating both failed and running workspaces.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 ` path:"workspace_id"`
}

type Workspace struct {
	// Databricks account ID
	AccountId string `json:"account_id,omitempty"`
	// The AWS region of the workspace Data Plane. For example, `us-west-2`.
	AwsRegion string `json:"aws_region,omitempty"`
	// The cloud name. This field always has the value `gcp`.
	Cloud string `json:"cloud,omitempty"`

	CloudResourceBucket *CloudResourceBucket `json:"cloud_resource_bucket,omitempty"`
	// Time in epoch milliseconds when the workspace was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the workspace&#39;s credential configuration object.
	CredentialsId string `json:"credentials_id,omitempty"`
	// The deployment name defines part of the subdomain for the workspace. The
	// workspace URL for web application and REST APIs is
	// `&lt;deployment-name&gt;.cloud.databricks.com`. This value must be unique
	// across all non-deleted deployments across all AWS regions.
	DeploymentName string `json:"deployment_name,omitempty"`
	// The Google Cloud region of the workspace data plane in your Google
	// account. For example, `us-east4`.
	Location string `json:"location,omitempty"`
	// ID of the key configuration for encrypting managed services.
	ManagedServicesCustomerManagedKeyId string `json:"managed_services_customer_managed_key_id,omitempty"`

	Network *WorkspaceNetwork `json:"network,omitempty"`
	// The pricing tier of the workspace. See
	// https://databricks.com/product/aws-pricing for available pricing tier
	// information.
	PricingTier WorkspacePricingTier `json:"pricing_tier,omitempty"`
	// Only used for PrivateLink, which is in Public Preview. This is the ID of
	// the workspace&#39;s private access settings object. This ID must be specified
	// for customers using [AWS
	// PrivateLink](https://aws.amazon.com/privatelink/) for either front-end
	// (user-to-workspace connection), back-end (data plane to control plane
	// connection), or both connection types. Before configuring PrivateLink, it
	// is important to read the [Databricks article about
	// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
	PrivateAccessSettingsId string `json:"private_access_settings_id,omitempty"`
	// ID of the workspace&#39;s storage configuration object.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// ID of the key configuration for encrypting workspace storage.
	StorageCustomerManagedKeyId string `json:"storage_customer_managed_key_id,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"workspace_id,omitempty"`
	// The human-readable name of the workspace.
	WorkspaceName string `json:"workspace_name,omitempty"`
	// The status of the workspace. For workspace creation, it is typically
	// initially `PROVISIONING`. Continue to check the status until the status
	// is `RUNNING`. For detailed instructions of creating a new workspace with
	// this API **including error handling** see [Create a new workspace using
	// the Account Management
	// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
	WorkspaceStatus WorkspaceWorkspaceStatus `json:"workspace_status,omitempty"`
	// Message describing the current workspace status.
	WorkspaceStatusMessage string `json:"workspace_status_message,omitempty"`
}

type WorkspaceNetwork struct {
	// The network configuration ID that is attached to the workspace. This
	// field is available only if the network is a customer-managed network.
	NetworkId string `json:"network_id,omitempty"`
}

// The pricing tier of the workspace. See
// https://databricks.com/product/aws-pricing for available pricing tier
// information.
type WorkspacePricingTier string

const WorkspacePricingTierCommunityEdition WorkspacePricingTier = `COMMUNITY_EDITION`

const WorkspacePricingTierDedicated WorkspacePricingTier = `DEDICATED`

const WorkspacePricingTierEnterprise WorkspacePricingTier = `ENTERPRISE`

const WorkspacePricingTierPremium WorkspacePricingTier = `PREMIUM`

const WorkspacePricingTierStandard WorkspacePricingTier = `STANDARD`

const WorkspacePricingTierUnknown WorkspacePricingTier = `UNKNOWN`

// The status of the workspace. For workspace creation, it is typically
// initially `PROVISIONING`. Continue to check the status until the status is
// `RUNNING`. For detailed instructions of creating a new workspace with this
// API **including error handling** see [Create a new workspace using the
// Account Management
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
type WorkspaceWorkspaceStatus string

const WorkspaceWorkspaceStatusBanned WorkspaceWorkspaceStatus = `BANNED`

const WorkspaceWorkspaceStatusCancelling WorkspaceWorkspaceStatus = `CANCELLING`

const WorkspaceWorkspaceStatusFailed WorkspaceWorkspaceStatus = `FAILED`

const WorkspaceWorkspaceStatusNotProvisioned WorkspaceWorkspaceStatus = `NOT_PROVISIONED`

const WorkspaceWorkspaceStatusProvisioning WorkspaceWorkspaceStatus = `PROVISIONING`

const WorkspaceWorkspaceStatusRunning WorkspaceWorkspaceStatus = `RUNNING`
