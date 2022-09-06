// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networks

// all definitions in this file are in alphabetical order

type CreateNetworkRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`

	GcpNetworkInfo *GcpNetworkInfo `json:"gcp_network_info,omitempty"`
	// The human-readable name of the network configuration.
	NetworkName string `json:"network_name"`
	// IDs of 1 to 5 security groups associated with this network. Security
	// groups IDs **cannot** be used in multiple network configurations.
	SecurityGroupIds []string `json:"security_group_ids,omitempty"`
	// IDs of at least 2 subnets associated with this network. Subnet IDs
	// **cannot** be used in multiple network configurations.
	SubnetIds []string `json:"subnet_ids,omitempty"`

	VpcEndpoints *NetworkVpcEndpoints `json:"vpc_endpoints,omitempty"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id,omitempty"`
}

type DeleteNetworkConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Databricks Account API network configuration ID.
	NetworkId string ` path:"network_id"`
}

// The Google Cloud specific information for this network, for example the VPC
// ID, subnet ID, and secondary IP ranges.
type GcpNetworkInfo struct {
	// The Google Cloud project ID of the VPC network.
	NetworkProjectId string `json:"network_project_id"`
	// The name of the secondary IP range for pods. A Databricks-managed GKE
	// cluster uses this IP range for its pods. This secondary IP range can only
	// be used by one workspace.
	PodIpRangeName string `json:"pod_ip_range_name"`
	// The name of the secondary IP range for services. A Databricks-managed GKE
	// cluster uses this IP range for its services. This secondary IP range can
	// only be used by one workspace.
	ServiceIpRangeName string `json:"service_ip_range_name"`
	// The ID of the subnet associated with this network.
	SubnetId string `json:"subnet_id,omitempty"`
	// The Google Cloud region of the workspace data plane. For example,
	// `us-east4`.
	SubnetRegion string `json:"subnet_region"`
	// The ID of the VPC associated with this network. VPC IDs can be used in
	// multiple network configurations.
	VpcId string `json:"vpc_id"`
}

type GetAllNetworkConfigsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
}

type GetNetworkConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Databricks Account API network configuration ID.
	NetworkId string ` path:"network_id"`
}

// Array of network configuration objects.
type ListNetworksResponse []Network

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
	// The status of this network configuration object in terms of its use in a
	// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`:
	// Broken. * `WARNED`: Warned.
	VpcStatus NetworkVpcStatus `json:"vpc_status,omitempty"`
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
	ErrorType NetworkHealthErrorType `json:"error_type,omitempty"`
}

// The AWS resource associated with this error: credentials, VPC, subnet,
// security group, or network ACL.
type NetworkHealthErrorType string

const NetworkHealthErrorTypeCredentials NetworkHealthErrorType = `credentials`

const NetworkHealthErrorTypeNetworkacl NetworkHealthErrorType = `networkAcl`

const NetworkHealthErrorTypeSecuritygroup NetworkHealthErrorType = `securityGroup`

const NetworkHealthErrorTypeSubnet NetworkHealthErrorType = `subnet`

const NetworkHealthErrorTypeVpc NetworkHealthErrorType = `vpc`

// If specified, contains the VPC endpoints used to allow cluster communication
// from this VPC over [AWS PrivateLink](https://aws.amazon.com/privatelink/).
type NetworkVpcEndpoints struct {
	// The VPC endpoint ID used by this Network to access the Databricks secure
	// cluster connectivity relay. See [Secure Cluster
	// Connectivity](https://docs.databricks.com/security/secure-cluster-connectivity.html).
	// This is a list type for future compatibility, but currently only one VPC
	// endpoint ID should be supplied. Note: This is the Databricks-specific ID
	// of the VPC endpoint object in the Account API, not the AWS VPC endpoint
	// ID that you see for your endpoint in the AWS Console.
	DataplaneRelay []string `json:"dataplane_relay"`
	// The VPC endpoint ID used by this Network to access the Databricks REST
	// API. Databricks clusters make calls to our REST API as part of cluster
	// creation, mlflow tracking, and many other features. Thus, this is
	// required even if your workspace allows public access to the REST API.
	// This is a list type for future compatibility, but currently only one VPC
	// endpoint ID should be supplied. Note: This is the Databricks-specific ID
	// of the VPC endpoint object in the Account API, not the AWS VPC endpoint
	// ID that you see for your endpoint in the AWS Console.
	RestApi []string `json:"rest_api"`
}

// The status of this network configuration object in terms of its use in a
// workspace: * `UNATTACHED`: Unattached. * `VALID`: Valid. * `BROKEN`: Broken.
// * `WARNED`: Warned.
type NetworkVpcStatus string

const NetworkVpcStatusBroken NetworkVpcStatus = `BROKEN`

const NetworkVpcStatusUnattached NetworkVpcStatus = `UNATTACHED`

const NetworkVpcStatusValid NetworkVpcStatus = `VALID`

const NetworkVpcStatusWarned NetworkVpcStatus = `WARNED`

type NetworkWarning struct {
	// Details of the warning.
	WarningMessage string `json:"warning_message,omitempty"`
	// The AWS resource associated with this warning: a subnet or a security
	// group.
	WarningType NetworkWarningWarningType `json:"warning_type,omitempty"`
}

// The AWS resource associated with this warning: a subnet or a security group.
type NetworkWarningWarningType string

const NetworkWarningWarningTypeSecuritygroup NetworkWarningWarningType = `securityGroup`

const NetworkWarningWarningTypeSubnet NetworkWarningWarningType = `subnet`
