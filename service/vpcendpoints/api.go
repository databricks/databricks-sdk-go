// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vpcendpoints

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewVpcEndpoints(client *client.DatabricksClient) VpcEndpointsService {
	return &VpcEndpointsAPI{
		client: client,
	}
}

type VpcEndpointsAPI struct {
	client *client.DatabricksClient
}

// Create VPC endpoint configuration
//
// Creates a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// **Important**: When you register a VPC endpoint to the Databricks workspace
// VPC endpoint service for any workspace, **in this release <Databricks>
// enables front-end (web application and REST API) access from the source
// network of the VPC endpoint to all workspaces in that AWS region in your
// <Databricks> account if the workspaces have any PrivateLink connections in
// their workspace configuration**. If you have questions about this behavior,
// contact your Databricks representative.
//
// Within AWS, your VPC endpoint stays in `pendingAcceptance` state until you
// register it in a VPC endpoint configuration through the Account API. After
// you register the VPC endpoint configuration, the Databricks [endpoint
// service](https://docs.aws.amazon.com/vpc/latest/privatelink/privatelink-share-your-services.html)
// automatically accepts the VPC endpoint and it eventually transitions to the
// `available` state.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) CreateVpcEndpoint(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", request.AccountId)
	err := a.client.Post(ctx, path, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

// Delete VPC endpoint configuration
//
// Deletes a VPC endpoint configuration, which represents an [AWS VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// that can communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// Upon deleting a VPC endpoint configuration, the VPC endpoint in AWS changes
// its state from `accepted` to `rejected`, which means that it is no longer
// usable from your VPC.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) DeleteVpcEndpoint(ctx context.Context, request DeleteVpcEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", request.AccountId, request.VpcEndpointId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete VPC endpoint configuration
//
// Deletes a VPC endpoint configuration, which represents an [AWS VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// that can communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// Upon deleting a VPC endpoint configuration, the VPC endpoint in AWS changes
// its state from `accepted` to `rejected`, which means that it is no longer
// usable from your VPC.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) DeleteVpcEndpointByAccountIdAndVpcEndpointId(ctx context.Context, accountId string, vpcEndpointId string) error {
	return a.DeleteVpcEndpoint(ctx, DeleteVpcEndpointRequest{
		AccountId:     accountId,
		VpcEndpointId: vpcEndpointId,
	})
}

// Get all VPC endpoint configurations
//
// Gets a list of all VPC endpoints for an account, specified by ID.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) GetAllVpcEndpoints(ctx context.Context, request GetAllVpcEndpointsRequest) ([]VpcEndpoint, error) {
	var vpcEndpointList []VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", request.AccountId)
	err := a.client.Get(ctx, path, request, &vpcEndpointList)
	return vpcEndpointList, err
}

// Get all VPC endpoint configurations
//
// Gets a list of all VPC endpoints for an account, specified by ID.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) GetAllVpcEndpointsByAccountId(ctx context.Context, accountId string) ([]VpcEndpoint, error) {
	return a.GetAllVpcEndpoints(ctx, GetAllVpcEndpointsRequest{
		AccountId: accountId,
	})
}

// Get a VPC endpoint configuration
//
// Gets a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) GetVpcEndpoint(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", request.AccountId, request.VpcEndpointId)
	err := a.client.Get(ctx, path, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

// Get a VPC endpoint configuration
//
// Gets a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) GetVpcEndpointByAccountIdAndVpcEndpointId(ctx context.Context, accountId string, vpcEndpointId string) (*VpcEndpoint, error) {
	return a.GetVpcEndpoint(ctx, GetVpcEndpointRequest{
		AccountId:     accountId,
		VpcEndpointId: vpcEndpointId,
	})
}
