// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package networks

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewNetworkConfigurations(client *client.DatabricksClient) NetworkConfigurationsService {
	return &NetworkConfigurationsAPI{
		client: client,
	}
}

type NetworkConfigurationsAPI struct {
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
func (a *NetworkConfigurationsAPI) CreateNetworkConfig(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", request.AccountId)
	err := a.client.Post(ctx, path, request, &network)
	return &network, err
}

// Delete a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace. This operation is available only if your account is on the E2
// version of the platform.
func (a *NetworkConfigurationsAPI) DeleteNetworkConfig(ctx context.Context, request DeleteNetworkConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", request.AccountId, request.NetworkId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace. This operation is available only if your account is on the E2
// version of the platform.
func (a *NetworkConfigurationsAPI) DeleteNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) error {
	return a.DeleteNetworkConfig(ctx, DeleteNetworkConfigRequest{
		AccountId: accountId,
		NetworkId: networkId,
	})
}

// Get a list of all Databricks network configurations for an account, specified
// by ID. This operation is available only if your account is on the E2 version
// of the platform.
func (a *NetworkConfigurationsAPI) GetAllNetworkConfigs(ctx context.Context, request GetAllNetworkConfigsRequest) (*ListNetworksResponse, error) {
	var listNetworksResponse ListNetworksResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", request.AccountId)
	err := a.client.Get(ctx, path, request, &listNetworksResponse)
	return &listNetworksResponse, err
}

// Get a list of all Databricks network configurations for an account, specified
// by ID. This operation is available only if your account is on the E2 version
// of the platform.
func (a *NetworkConfigurationsAPI) GetAllNetworkConfigsByAccountId(ctx context.Context, accountId string) (*ListNetworksResponse, error) {
	return a.GetAllNetworkConfigs(ctx, GetAllNetworkConfigsRequest{
		AccountId: accountId,
	})
}

// Get a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-created VPC and subnets. For VPC requirements,
// see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetNetworkConfig(ctx context.Context, request GetNetworkConfigRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", request.AccountId, request.NetworkId)
	err := a.client.Get(ctx, path, request, &network)
	return &network, err
}

// Get a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-created VPC and subnets. For VPC requirements,
// see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) (*Network, error) {
	return a.GetNetworkConfig(ctx, GetNetworkConfigRequest{
		AccountId: accountId,
		NetworkId: networkId,
	})
}
