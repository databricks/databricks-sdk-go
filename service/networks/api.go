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

// Create network configuration
//
// Creates a Databricks network configuration that represents an AWS VPC and its
// resources. The VPC will be used for new Databricks clusters. This requires a
// pre-existing VPC and subnets. For VPC requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
//
// **Important**: You can share one customer-managed VPC with multiple
// workspaces in a single account. Therefore, you can share one VPC across
// multiple Account API network configurations. However, you **cannot** reuse
// subnets or Security Groups between workspaces. Because a Databricks Account
// API network configuration encapsulates this information, you cannot reuse a
// Databricks Account API network configuration across workspaces. If you plan
// to share one VPC with multiple workspaces, make sure you size your VPC and
// subnets accordingly. For information about how to create a new workspace with
// this API, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) CreateNetworkConfig(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", request.AccountId)
	err := a.client.Post(ctx, path, request, &network)
	return &network, err
}

// Delete network configuration
//
// Deletes a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) DeleteNetworkConfig(ctx context.Context, request DeleteNetworkConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", request.AccountId, request.NetworkId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete network configuration
//
// Deletes a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) DeleteNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) error {
	return a.DeleteNetworkConfig(ctx, DeleteNetworkConfigRequest{
		AccountId: accountId,
		NetworkId: networkId,
	})
}

// Get all network configurations
//
// Gets a list of all Databricks network configurations for an account,
// specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetAllNetworkConfigs(ctx context.Context, request GetAllNetworkConfigsRequest) ([]Network, error) {
	var networkList []Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", request.AccountId)
	err := a.client.Get(ctx, path, request, &networkList)
	return networkList, err
}

func (a *NetworkConfigurationsAPI) NetworkNetworkNameToNetworkIdMap(ctx context.Context, request GetAllNetworkConfigsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.GetAllNetworkConfigs(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.NetworkName] = v.NetworkId
	}
	return mapping, nil
}

func (a *NetworkConfigurationsAPI) GetNetworkByNetworkName(ctx context.Context, name string) (*Network, error) {
	result, err := a.GetAllNetworkConfigs(ctx, GetAllNetworkConfigsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.NetworkName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Network named '%s' does not exist", name)
}

// Get all network configurations
//
// Gets a list of all Databricks network configurations for an account,
// specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetAllNetworkConfigsByAccountId(ctx context.Context, accountId string) ([]Network, error) {
	return a.GetAllNetworkConfigs(ctx, GetAllNetworkConfigsRequest{
		AccountId: accountId,
	})
}

// Get a network configuration
//
// Gets a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-existing VPC and subnets. For VPC
// requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetNetworkConfig(ctx context.Context, request GetNetworkConfigRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", request.AccountId, request.NetworkId)
	err := a.client.Get(ctx, path, request, &network)
	return &network, err
}

// Get a network configuration
//
// Gets a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-existing VPC and subnets. For VPC
// requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetNetworkConfigByAccountIdAndNetworkId(ctx context.Context, accountId string, networkId string) (*Network, error) {
	return a.GetNetworkConfig(ctx, GetNetworkConfigRequest{
		AccountId: accountId,
		NetworkId: networkId,
	})
}
