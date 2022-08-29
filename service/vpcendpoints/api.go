// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package vpcendpoints

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage VPC endpoint configurations for this account. This object
// registers an AWS VPC endpoint in your Databricks account so your workspace
// can use it with AWS PrivateLink. Your VPC endpoint connects to one of two VPC
// endpoint services -- one for workspace (both for front-end connection and for
// back-end connection to REST APIs) and one for the back-end secure cluster
// connectivity relay from the data plane. Your account must be enabled for
// PrivateLink to use these APIs. Before configuring PrivateLink, it is
// important to read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
type VpcendpointsService interface {
    // Create a VPC endpoint configuration, which represents a [VPC
    // endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
    // object in AWS used to communicate privately with Databricks over [AWS
    // PrivateLink](https://aws.amazon.com/privatelink). **IMPORTANT**: When you
    // register a VPC endpoint to the Databricks workspace VPC endpoint service
    // for any workspace, **in this release &lt;Databricks&gt; enables front-end (web
    // application and REST API) access from the source network of the VPC
    // endpoint to all workspaces in that AWS region in your &lt;Databricks&gt;
    // account if the workspaces have any PrivateLink connections in their
    // workspace configuration**. If you have questions about this behavior,
    // contact your Databricks representative. Within AWS, your VPC endpoint
    // stays in `pendingAcceptance` state until you register it in a VPC
    // endpoint configuration through the Account API. Upon doing so, the
    // Databricks [endpoint
    // service](https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html)
    // automatically accepts the VPC endpoint and it eventually transitions to
    // the `available` state. Before configuring PrivateLink, it is important to
    // read the [Databricks article about
    // PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
    // This operation is available only if your account is on the E2 version of
    // the platform and your Databricks account is enabled for PrivateLink
    // (Public Preview). Contact your Databricks representative to enable your
    // account for PrivateLink.
    CreateVpcEndpoint(ctx context.Context, createVPCEndpointRequest CreateVPCEndpointRequest) (*VPCEndpoint, error)
    // Delete a VPC endpoint configuration, which represents an [AWS VPC
    // endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
    // that can communicate privately with Databricks over [AWS
    // PrivateLink](https://aws.amazon.com/privatelink). Upon deleting a VPC
    // endpoint configuration, the VPC endpoint in AWS changes its state from
    // `accepted` to `rejected`, meaning it will no longer be usable from your
    // VPC. Before configuring PrivateLink, it is important to read the
    // [Databricks article about
    // PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
    // This operation is available only if your account is on the E2 version of
    // the platform and your Databricks account is enabled for PrivateLink
    // (Public Preview). Contact your Databricks representative to enable your
    // account for PrivateLink.
    DeleteVpcEndpoint(ctx context.Context, deleteVpcEndpointRequest DeleteVpcEndpointRequest) error
    // Get a list of all VPC endpoints for an account, specified by ID. Before
    // configuring PrivateLink, it is important to read the [Databricks article
    // about
    // PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
    // This operation is available only if your account is on the E2 version of
    // the platform and your Databricks account is enabled for PrivateLink
    // (Public Preview). Contact your Databricks representative to enable your
    // account for PrivateLink.
    GetAllVpcEndpoints(ctx context.Context, getAllVpcEndpointsRequest GetAllVpcEndpointsRequest) (*ListVPCEndpointsResponse, error)
    // Get a VPC endpoint configuration, which represents a [VPC
    // endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
    // object in AWS used to communicate privately with Databricks over [AWS
    // PrivateLink](https://aws.amazon.com/privatelink). This operation is
    // available only if your account is on the E2 version of the platform and
    // your Databricks account is enabled for PrivateLink (Public Preview).
    // Contact your Databricks representative to enable your account for
    // PrivateLink.
    GetVpcEndpoint(ctx context.Context, getVpcEndpointRequest GetVpcEndpointRequest) (*VPCEndpoint, error)
	GetVpcEndpointByAccountIdAndVpcEndpointId(ctx context.Context, accountId string, vpcEndpointId string) (*VPCEndpoint, error)
	DeleteVpcEndpointByAccountIdAndVpcEndpointId(ctx context.Context, accountId string, vpcEndpointId string) error
	GetAllVpcEndpointsByAccountId(ctx context.Context, accountId string) (*ListVPCEndpointsResponse, error)
}

func New(client *client.DatabricksClient) VpcendpointsService {
	return &VpcendpointsAPI{
		client: client,
	}
}

type VpcendpointsAPI struct {
	client *client.DatabricksClient
}

// Create a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). **IMPORTANT**: When you
// register a VPC endpoint to the Databricks workspace VPC endpoint service for
// any workspace, **in this release &lt;Databricks&gt; enables front-end (web
// application and REST API) access from the source network of the VPC endpoint
// to all workspaces in that AWS region in your &lt;Databricks&gt; account if the
// workspaces have any PrivateLink connections in their workspace
// configuration**. If you have questions about this behavior, contact your
// Databricks representative. Within AWS, your VPC endpoint stays in
// `pendingAcceptance` state until you register it in a VPC endpoint
// configuration through the Account API. Upon doing so, the Databricks
// [endpoint
// service](https://docs.aws.amazon.com/vpc/latest/privatelink/endpoint-service.html)
// automatically accepts the VPC endpoint and it eventually transitions to the
// `available` state. Before configuring PrivateLink, it is important to read
// the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcendpointsAPI) CreateVpcEndpoint(ctx context.Context, request CreateVPCEndpointRequest) (*VPCEndpoint, error) {
	var vPCEndpoint VPCEndpoint
	path := fmt.Sprintf("/accounts/%v/vpc-endpoints", request.AccountId)
	err := a.client.Post(ctx, path, request, &vPCEndpoint)
	return &vPCEndpoint, err
}

// Delete a VPC endpoint configuration, which represents an [AWS VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
// that can communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). Upon deleting a VPC
// endpoint configuration, the VPC endpoint in AWS changes its state from
// `accepted` to `rejected`, meaning it will no longer be usable from your VPC.
// Before configuring PrivateLink, it is important to read the [Databricks
// article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcendpointsAPI) DeleteVpcEndpoint(ctx context.Context, request DeleteVpcEndpointRequest) error {
	path := fmt.Sprintf("/accounts/%v/vpc-endpoints/%v", request.AccountId, request.VpcEndpointId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a list of all VPC endpoints for an account, specified by ID. Before
// configuring PrivateLink, it is important to read the [Databricks article
// about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcendpointsAPI) GetAllVpcEndpoints(ctx context.Context, request GetAllVpcEndpointsRequest) (*ListVPCEndpointsResponse, error) {
	var listVPCEndpointsResponse ListVPCEndpointsResponse
	path := fmt.Sprintf("/accounts/%v/vpc-endpoints", request.AccountId)
	err := a.client.Get(ctx, path, request, &listVPCEndpointsResponse)
	return &listVPCEndpointsResponse, err
}

// Get a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). This operation is available
// only if your account is on the E2 version of the platform and your Databricks
// account is enabled for PrivateLink (Public Preview). Contact your Databricks
// representative to enable your account for PrivateLink.
func (a *VpcendpointsAPI) GetVpcEndpoint(ctx context.Context, request GetVpcEndpointRequest) (*VPCEndpoint, error) {
	var vPCEndpoint VPCEndpoint
	path := fmt.Sprintf("/accounts/%v/vpc-endpoints/%v", request.AccountId, request.VpcEndpointId)
	err := a.client.Get(ctx, path, request, &vPCEndpoint)
	return &vPCEndpoint, err
}


func (a *VpcendpointsAPI) GetVpcEndpointByAccountIdAndVpcEndpointId(ctx context.Context, accountId string, vpcEndpointId string) (*VPCEndpoint, error) {
	return a.GetVpcEndpoint(ctx, GetVpcEndpointRequest{
		AccountId: accountId,
		VpcEndpointId: vpcEndpointId,
	})
}

func (a *VpcendpointsAPI) DeleteVpcEndpointByAccountIdAndVpcEndpointId(ctx context.Context, accountId string, vpcEndpointId string) error {
	return a.DeleteVpcEndpoint(ctx, DeleteVpcEndpointRequest{
		AccountId: accountId,
		VpcEndpointId: vpcEndpointId,
	})
}

func (a *VpcendpointsAPI) GetAllVpcEndpointsByAccountId(ctx context.Context, accountId string) (*ListVPCEndpointsResponse, error) {
	return a.GetAllVpcEndpoints(ctx, GetAllVpcEndpointsRequest{
		AccountId: accountId,
	})
}
