// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vpcendpoints

// all definitions in this file are in alphabetical order

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

type DeleteVpcEndpointRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" path:"vpc_endpoint_id"`
}

type GetAllVpcEndpointsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetVpcEndpointRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks VPC endpoint ID.
	VpcEndpointId string `json:"-" path:"vpc_endpoint_id"`
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
	UseCase VpcEndpointUseCase `json:"use_case,omitempty"`
	// Databricks VPC endpoint ID. This is the Databricks-specific name of the
	// VPC endpoint. Do not confuse this with the `aws_vpc_endpoint_id`, which
	// is the ID within AWS of the VPC endpoint.
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`
	// The human-readable name of the storage configuration.
	VpcEndpointName string `json:"vpc_endpoint_name,omitempty"`
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
type VpcEndpointUseCase string

const VpcEndpointUseCaseDataplaneRelayAccess VpcEndpointUseCase = `DATAPLANE_RELAY_ACCESS`

const VpcEndpointUseCaseWorkspaceAccess VpcEndpointUseCase = `WORKSPACE_ACCESS`
