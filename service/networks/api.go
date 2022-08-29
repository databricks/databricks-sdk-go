// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networks

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage network configurations for customer-managed VPCs
// (optional). A network configuration encapsulates the IDs for AWS VPCs,
// subnets, and security groups. Its ID is used when creating a new workspace if
// you use customer-managed VPCs.
type NetworksService interface {
    // Create a Databricks network configuration that represents an AWS VPC and
    // its resources. The VPC will be used for new Databricks clusters. This
    // requires a pre-created VPC and subnets. For VPC requirements, see
    // [Customer-managed
    // VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
    // **Important**: You can share one customer-managed VPC with multiple
    // workspaces in a single account. Therefore, you can share one VPC across
    // multiple Account API network configurations. However, you **cannot**
    // reuse subnets or Security Groups between workspaces. Because a Databricks
    // Account API network configuration encapsulates this information, you
    // cannot reuse a Databricks Account API network configuration across
    // workspaces. If you plan to share one VPC with multiple workspaces, be
    // sure to size your VPC and subnets accordingly. For detailed instructions
    // of creating a new workspace with this API, see [Create a new workspace
    // using the Account
    // API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
    // This operation is available only if your account is on the E2 version of
    // the platform.
    CreateNetworkConfig(ctx context.Context, createNetworkRequest CreateNetworkRequest) (*Network, error)
    // Delete a Databricks network configuration, which represents an AWS VPC
    // and its resources. You cannot delete a network that is associated with a
    // workspace. This operation is available only if your account is on the E2
    // version of the platform.
    DeleteNetworkConfig(ctx context.Context, deleteNetworkConfigRequest DeleteNetworkConfigRequest) error
    // Get a list of all Databricks network configurations for an account,
    // specified by ID. This operation is available only if your account is on
    // the E2 version of the platform.
    GetAllNetworkConfigs(ctx context.Context, getAllNetworkConfigsRequest GetAllNetworkConfigsRequest) (*ListNetworksResponse, error)
    // Get a Databricks network configuration, which represents an AWS VPC and
    // its resources. This requires a pre-created VPC and subnets. For VPC
    // requirements, see [Customer-managed
    // VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
    // This operation is available only if your account is on the E2 version of
    // the platform.
    GetNetworkConfig(ctx context.Context, getNetworkConfigRequest GetNetworkConfigRequest) (*Network, error)
	GetAllNetworkConfigsByAccountId(ctx context.Context, accountId string) (*ListNetworksResponse, error)
	GetNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) (*Network, error)
	DeleteNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) error
}

func New(client *client.DatabricksClient) NetworksService {
	return &NetworksAPI{
		client: client,
	}
}

type NetworksAPI struct {
	client *client.DatabricksClient
}

// Create a Databricks network configuration that represents an AWS VPC and its
// resources. The VPC will be used for new Databricks clusters. This requires a
// pre-created VPC and subnets. For VPC requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
// **Important**: You can share one customer-managed VPC with multiple
// workspaces in a single account. Therefore, you can share one VPC across
// multiple Account API network configurations. However, you **cannot** reuse
// subnets or Security Groups between workspaces. Because a Databricks Account
// API network configuration encapsulates this information, you cannot reuse a
// Databricks Account API network configuration across workspaces. If you plan
// to share one VPC with multiple workspaces, be sure to size your VPC and
// subnets accordingly. For detailed instructions of creating a new workspace
// with this API, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworksAPI) CreateNetworkConfig(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/accounts/%v/networks", request.AccountId)
	err := a.client.Post(ctx, path, request, &network)
	return &network, err
}

// Delete a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace. This operation is available only if your account is on the E2
// version of the platform.
func (a *NetworksAPI) DeleteNetworkConfig(ctx context.Context, request DeleteNetworkConfigRequest) error {
	path := fmt.Sprintf("/accounts/%v/networks/%v", request.AccountId, request.NetworkId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a list of all Databricks network configurations for an account, specified
// by ID. This operation is available only if your account is on the E2 version
// of the platform.
func (a *NetworksAPI) GetAllNetworkConfigs(ctx context.Context, request GetAllNetworkConfigsRequest) (*ListNetworksResponse, error) {
	var listNetworksResponse ListNetworksResponse
	path := fmt.Sprintf("/accounts/%v/networks", request.AccountId)
	err := a.client.Get(ctx, path, request, &listNetworksResponse)
	return &listNetworksResponse, err
}

// Get a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-created VPC and subnets. For VPC requirements,
// see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworksAPI) GetNetworkConfig(ctx context.Context, request GetNetworkConfigRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/accounts/%v/networks/%v", request.AccountId, request.NetworkId)
	err := a.client.Get(ctx, path, request, &network)
	return &network, err
}


func (a *NetworksAPI) GetAllNetworkConfigsByAccountId(ctx context.Context, accountId string) (*ListNetworksResponse, error) {
	return a.GetAllNetworkConfigs(ctx, GetAllNetworkConfigsRequest{
		AccountId: accountId,
	})
}

func (a *NetworksAPI) GetNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) (*Network, error) {
	return a.GetNetworkConfig(ctx, GetNetworkConfigRequest{
		AccountId: accountId,
		NetworkId: networkId,
	})
}

func (a *NetworksAPI) DeleteNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) error {
	return a.DeleteNetworkConfig(ctx, DeleteNetworkConfigRequest{
		AccountId: accountId,
		NetworkId: networkId,
	})
}
