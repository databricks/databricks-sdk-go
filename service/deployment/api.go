// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deployment

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/retries"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

func NewCredentialConfigurations(client *client.DatabricksClient) CredentialConfigurationsService {
	return &CredentialConfigurationsAPI{
		client: client,
	}
}

type CredentialConfigurationsAPI struct {
	client *client.DatabricksClient
}

// Create credential configuration
//
// Creates a Databricks credential configuration that represents cloud
// cross-account credentials for a specified account. Databricks uses this to
// set up network infrastructure properly to host Databricks clusters. For your
// AWS IAM role, you need to trust the External ID (the Databricks Account API
// account ID) in the returned credential object, and configure the required
// access policy.
//
// Save the response's `credentials_id` field, which is the ID for your new
// credential configuration object.
//
// For information about how to create a new workspace with this API, see
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *CredentialConfigurationsAPI) CreateCredentialConfig(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", request.AccountId)
	err := a.client.Post(ctx, path, request, &credential)
	return &credential, err
}

// Delete credential configuration
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialConfigurationsAPI) DeleteCredentialConfig(ctx context.Context, request DeleteCredentialConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", request.AccountId, request.CredentialsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete credential configuration
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialConfigurationsAPI) DeleteCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) error {
	return a.DeleteCredentialConfig(ctx, DeleteCredentialConfigRequest{
		AccountId:     accountId,
		CredentialsId: credentialsId,
	})
}

// Get credential configuration
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialConfigurationsAPI) GetCredentialConfig(ctx context.Context, request GetCredentialConfigRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", request.AccountId, request.CredentialsId)
	err := a.client.Get(ctx, path, request, &credential)
	return &credential, err
}

// Get credential configuration
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialConfigurationsAPI) GetCredentialConfigByAccountIdAndCredentialsId(ctx context.Context, accountId string, credentialsId string) (*Credential, error) {
	return a.GetCredentialConfig(ctx, GetCredentialConfigRequest{
		AccountId:     accountId,
		CredentialsId: credentialsId,
	})
}

// Get all credential configurations
//
// Gets all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialConfigurationsAPI) ListCredentials(ctx context.Context, request ListCredentialsRequest) ([]Credential, error) {
	var credentialList []Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", request.AccountId)
	err := a.client.Get(ctx, path, request, &credentialList)
	return credentialList, err
}

func (a *CredentialConfigurationsAPI) CredentialCredentialsNameToCredentialsIdMap(ctx context.Context, request ListCredentialsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListCredentials(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.CredentialsName] = v.CredentialsId
	}
	return mapping, nil
}

func (a *CredentialConfigurationsAPI) GetCredentialByCredentialsName(ctx context.Context, name string) (*Credential, error) {
	result, err := a.ListCredentials(ctx, ListCredentialsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.CredentialsName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Credential named '%s' does not exist", name)
}

// Get all credential configurations
//
// Gets all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialConfigurationsAPI) ListCredentialsByAccountId(ctx context.Context, accountId string) ([]Credential, error) {
	return a.ListCredentials(ctx, ListCredentialsRequest{
		AccountId: accountId,
	})
}

func NewKeyConfigurations(client *client.DatabricksClient) KeyConfigurationsService {
	return &KeyConfigurationsAPI{
		client: client,
	}
}

type KeyConfigurationsAPI struct {
	client *client.DatabricksClient
}

// Create encryption key configuration
//
// Creates a customer-managed key configuration object for an account, specified
// by ID. This operation uploads a reference to a customer-managed key to
// Databricks. If the key is assigned as a workspace's customer-managed key for
// managed services, Databricks uses the key to encrypt the workspaces notebooks
// and secrets in the control plane, in addition to Databricks SQL queries and
// query history. If it is specified as a workspace's customer-managed key for
// workspace storage, the key encrypts the workspace's root S3 bucket (which
// contains the workspace's root DBFS and system data) and, optionally, cluster
// EBS volume data.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *KeyConfigurationsAPI) CreateKeyConfig(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", request.AccountId)
	err := a.client.Post(ctx, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

// Delete encryption key configuration
//
// Deletes a customer-managed key configuration object for an account. You
// cannot delete a configuration that is associated with a running workspace.
func (a *KeyConfigurationsAPI) DeleteKeyConfig(ctx context.Context, request DeleteKeyConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", request.AccountId, request.CustomerManagedKeyId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete encryption key configuration
//
// Deletes a customer-managed key configuration object for an account. You
// cannot delete a configuration that is associated with a running workspace.
func (a *KeyConfigurationsAPI) DeleteKeyConfigByAccountIdAndCustomerManagedKeyId(ctx context.Context, accountId string, customerManagedKeyId string) error {
	return a.DeleteKeyConfig(ctx, DeleteKeyConfigRequest{
		AccountId:            accountId,
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

// Get all encryption key configurations
//
// Gets all customer-managed key configuration objects for an account. If the
// key is specified as a workspace's managed services customer-managed key,
// Databricks uses the key to encrypt the workspace's notebooks and secrets in
// the control plane, in addition to Databricks SQL queries and query history.
// If the key is specified as a workspace's storage customer-managed key, the
// key is used to encrypt the workspace's root S3 bucket and optionally can
// encrypt cluster EBS volumes data in the data plane.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetAllKeyConfigs(ctx context.Context, request GetAllKeyConfigsRequest) ([]CustomerManagedKey, error) {
	var customerManagedKeyList []CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", request.AccountId)
	err := a.client.Get(ctx, path, request, &customerManagedKeyList)
	return customerManagedKeyList, err
}

// Get all encryption key configurations
//
// Gets all customer-managed key configuration objects for an account. If the
// key is specified as a workspace's managed services customer-managed key,
// Databricks uses the key to encrypt the workspace's notebooks and secrets in
// the control plane, in addition to Databricks SQL queries and query history.
// If the key is specified as a workspace's storage customer-managed key, the
// key is used to encrypt the workspace's root S3 bucket and optionally can
// encrypt cluster EBS volumes data in the data plane.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetAllKeyConfigsByAccountId(ctx context.Context, accountId string) ([]CustomerManagedKey, error) {
	return a.GetAllKeyConfigs(ctx, GetAllKeyConfigsRequest{
		AccountId: accountId,
	})
}

// Get encryption key configuration
//
// Gets a customer-managed key configuration object for an account, specified by
// ID. This operation uploads a reference to a customer-managed key to
// Databricks. If assigned as a workspace's customer-managed key for managed
// services, Databricks uses the key to encrypt the workspaces notebooks and
// secrets in the control plane, in addition to Databricks SQL queries and query
// history. If it is specified as a workspace's customer-managed key for
// storage, the key encrypts the workspace's root S3 bucket (which contains the
// workspace's root DBFS and system data) and, optionally, cluster EBS volume
// data.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetKeyConfig(ctx context.Context, request GetKeyConfigRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", request.AccountId, request.CustomerManagedKeyId)
	err := a.client.Get(ctx, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

// Get encryption key configuration
//
// Gets a customer-managed key configuration object for an account, specified by
// ID. This operation uploads a reference to a customer-managed key to
// Databricks. If assigned as a workspace's customer-managed key for managed
// services, Databricks uses the key to encrypt the workspaces notebooks and
// secrets in the control plane, in addition to Databricks SQL queries and query
// history. If it is specified as a workspace's customer-managed key for
// storage, the key encrypts the workspace's root S3 bucket (which contains the
// workspace's root DBFS and system data) and, optionally, cluster EBS volume
// data.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetKeyConfigByAccountIdAndCustomerManagedKeyId(ctx context.Context, accountId string, customerManagedKeyId string) (*CustomerManagedKey, error) {
	return a.GetKeyConfig(ctx, GetKeyConfigRequest{
		AccountId:            accountId,
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

// Get history of a key's associations with workspaces
//
// Gets a list of records that show how key configurations are associated with
// workspaces.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetKeyWorkspaceHistory(ctx context.Context, request GetKeyWorkspaceHistoryRequest) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	var listWorkspaceEncryptionKeyRecordsResponse ListWorkspaceEncryptionKeyRecordsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-key-history", request.AccountId)
	err := a.client.Get(ctx, path, request, &listWorkspaceEncryptionKeyRecordsResponse)
	return &listWorkspaceEncryptionKeyRecordsResponse, err
}

// Get history of a key's associations with workspaces
//
// Gets a list of records that show how key configurations are associated with
// workspaces.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetKeyWorkspaceHistoryByAccountId(ctx context.Context, accountId string) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	return a.GetKeyWorkspaceHistory(ctx, GetKeyWorkspaceHistoryRequest{
		AccountId: accountId,
	})
}

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

func NewPrivateAccessSettings(client *client.DatabricksClient) PrivateAccessSettingsService {
	return &PrivateAccessSettingsAPI{
		client: client,
	}
}

type PrivateAccessSettingsAPI struct {
	client *client.DatabricksClient
}

// Create private access settings
//
// Creates a private access settings object, which specifies how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink). To
// use AWS PrivateLink, a workspace must have a private access settings object
// referenced by ID in the workspace's `private_access_settings_id` property.
//
// You can share one private access settings with multiple workspaces in a
// single account. However, private access settings are specific to AWS regions,
// so only workspaces in the same AWS region can use a given private access
// settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) CreatePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", request.AccountId)
	err := a.client.Post(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

// Delete a private access settings object
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) DeletePrivateAccessSettings(ctx context.Context, request DeletePrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete a private access settings object
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) DeletePrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) error {
	return a.DeletePrivateAccessSettings(ctx, DeletePrivateAccessSettingsRequest{
		AccountId:               accountId,
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Get all private access settings objects
//
// Gets a list of all private access settings objects for an account, specified
// by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for AWS PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetAllPrivateAccessSettings(ctx context.Context, request GetAllPrivateAccessSettingsRequest) ([]PrivateAccessSettings, error) {
	var privateAccessSettingsList []PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", request.AccountId)
	err := a.client.Get(ctx, path, request, &privateAccessSettingsList)
	return privateAccessSettingsList, err
}

func (a *PrivateAccessSettingsAPI) PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx context.Context, request GetAllPrivateAccessSettingsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.GetAllPrivateAccessSettings(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.PrivateAccessSettingsName] = v.PrivateAccessSettingsId
	}
	return mapping, nil
}

func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettingsByPrivateAccessSettingsName(ctx context.Context, name string) (*PrivateAccessSettings, error) {
	result, err := a.GetAllPrivateAccessSettings(ctx, GetAllPrivateAccessSettingsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.PrivateAccessSettingsName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("PrivateAccessSettings named '%s' does not exist", name)
}

// Get all private access settings objects
//
// Gets a list of all private access settings objects for an account, specified
// by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for AWS PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetAllPrivateAccessSettingsByAccountId(ctx context.Context, accountId string) ([]PrivateAccessSettings, error) {
	return a.GetAllPrivateAccessSettings(ctx, GetAllPrivateAccessSettingsRequest{
		AccountId: accountId,
	})
}

// Get a private access settings object
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettings(ctx context.Context, request GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Get(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

// Get a private access settings object
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettingsByAccountIdAndPrivateAccessSettingsId(ctx context.Context, accountId string, privateAccessSettingsId string) (*PrivateAccessSettings, error) {
	return a.GetPrivateAccessSettings(ctx, GetPrivateAccessSettingsRequest{
		AccountId:               accountId,
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Replace private access settings
//
// Updates an existing private access settings object, which specifies how your
// workspace is accessed over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink, a
// workspace must have a private access settings object referenced by ID in the
// workspace's `private_access_settings_id` property.
//
// This operation completely overwrites your existing private access settings
// object attached to your workspaces. All workspaces attached to the private
// access settings are affected by any change. If `public_access_enabled`,
// `private_access_level`, or `allowed_vpc_endpoint_ids` are updated, effects of
// these changes might take several minutes to propagate to the workspace API.
// You can share one private access settings object with multiple workspaces in
// a single account. However, private access settings are specific to AWS
// regions, so only workspaces in the same AWS region can use a given private
// access settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) ReplacePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", request.AccountId, request.PrivateAccessSettingsId)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewStorageConfigurations(client *client.DatabricksClient) StorageConfigurationsService {
	return &StorageConfigurationsAPI{
		client: client,
	}
}

type StorageConfigurationsAPI struct {
	client *client.DatabricksClient
}

// Create new storage configuration
//
// Creates new storage configuration for an account, specified by ID. Uploads a
// storage configuration object that represents the root AWS S3 bucket in your
// account. Databricks stores related workspace assets including DBFS, cluster
// logs, and job results. For the AWS S3 bucket, you need to configure the
// required bucket policy.
//
// For information about how to create a new workspace with this API, see
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *StorageConfigurationsAPI) CreateStorageConfig(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", request.AccountId)
	err := a.client.Post(ctx, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

// Delete storage configuration
//
// Deletes a Databricks storage configuration. You cannot delete a storage
// configuration that is associated with any workspace.
func (a *StorageConfigurationsAPI) DeleteStorageConfig(ctx context.Context, request DeleteStorageConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", request.AccountId, request.StorageConfigurationId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete storage configuration
//
// Deletes a Databricks storage configuration. You cannot delete a storage
// configuration that is associated with any workspace.
func (a *StorageConfigurationsAPI) DeleteStorageConfigByAccountIdAndStorageConfigurationId(ctx context.Context, accountId string, storageConfigurationId string) error {
	return a.DeleteStorageConfig(ctx, DeleteStorageConfigRequest{
		AccountId:              accountId,
		StorageConfigurationId: storageConfigurationId,
	})
}

// Get all storage configurations
//
// Gets a list of all Databricks storage configurations for your account,
// specified by ID.
func (a *StorageConfigurationsAPI) GetAllStorageConfigs(ctx context.Context, request GetAllStorageConfigsRequest) ([]StorageConfiguration, error) {
	var storageConfigurationList []StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", request.AccountId)
	err := a.client.Get(ctx, path, request, &storageConfigurationList)
	return storageConfigurationList, err
}

func (a *StorageConfigurationsAPI) StorageConfigurationStorageConfigurationNameToStorageConfigurationIdMap(ctx context.Context, request GetAllStorageConfigsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.GetAllStorageConfigs(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.StorageConfigurationName] = v.StorageConfigurationId
	}
	return mapping, nil
}

func (a *StorageConfigurationsAPI) GetStorageConfigurationByStorageConfigurationName(ctx context.Context, name string) (*StorageConfiguration, error) {
	result, err := a.GetAllStorageConfigs(ctx, GetAllStorageConfigsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.StorageConfigurationName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("StorageConfiguration named '%s' does not exist", name)
}

// Get all storage configurations
//
// Gets a list of all Databricks storage configurations for your account,
// specified by ID.
func (a *StorageConfigurationsAPI) GetAllStorageConfigsByAccountId(ctx context.Context, accountId string) ([]StorageConfiguration, error) {
	return a.GetAllStorageConfigs(ctx, GetAllStorageConfigsRequest{
		AccountId: accountId,
	})
}

// Get storage configuration
//
// Gets a Databricks storage configuration for an account, both specified by ID.
func (a *StorageConfigurationsAPI) GetStorageConfig(ctx context.Context, request GetStorageConfigRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", request.AccountId, request.StorageConfigurationId)
	err := a.client.Get(ctx, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

// Get storage configuration
//
// Gets a Databricks storage configuration for an account, both specified by ID.
func (a *StorageConfigurationsAPI) GetStorageConfigByAccountIdAndStorageConfigurationId(ctx context.Context, accountId string, storageConfigurationId string) (*StorageConfiguration, error) {
	return a.GetStorageConfig(ctx, GetStorageConfigRequest{
		AccountId:              accountId,
		StorageConfigurationId: storageConfigurationId,
	})
}

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

func NewWorkspaces(client *client.DatabricksClient) WorkspacesService {
	return &WorkspacesAPI{
		client: client,
	}
}

type WorkspacesAPI struct {
	client *client.DatabricksClient
}

// Create a new workspace
//
// Creates a new workspace using a credential configuration and a storage
// configuration, an optional network configuration (if using a customer-managed
// VPC), an optional managed services key configuration (if using
// customer-managed keys for managed services), and an optional storage key
// configuration (if using customer-managed keys for storage). The key
// configurations used for managed services and storage encryption can be the
// same or different.
//
// **Important**: This operation is asynchronous. A response with HTTP status
// code 200 means the request has been accepted and is in progress, but does not
// mean that the workspace deployed successfully and is running. The initial
// workspace status is typically `PROVISIONING`. Use the workspace ID
// (`workspace_id`) field in the response to identify the new workspace and make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// You can share one customer-managed VPC with multiple workspaces in a single
// account. It is not required to create a new VPC for each workspace. However,
// you **cannot** reuse subnets or Security Groups between workspaces. If you
// plan to share one VPC with multiple workspaces, make sure you size your VPC
// and subnets accordingly. Because a Databricks Account API network
// configuration encapsulates this information, you cannot reuse a Databricks
// Account API network configuration across workspaces.\nFor information about
// how to create a new workspace with this API **including error handling**, see
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// **Important**: Customer-managed VPCs, PrivateLink, and customer-managed keys
// are supported on a limited set of deployment and subscription types. If you
// have questions about availability, contact your Databricks
// representative.\n\nThis operation is available only if your account is on the
// E2 version of the platform or on a select custom plan that allows multiple
// workspaces per account.
func (a *WorkspacesAPI) CreateWorkspace(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", request.AccountId)
	err := a.client.Post(ctx, path, request, &workspace)
	return &workspace, err
}

// CreateWorkspace and wait to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Workspace](60*time.Minute) functional option.
func (a *WorkspacesAPI) CreateWorkspaceAndWait(ctx context.Context, createWorkspaceRequest CreateWorkspaceRequest, options ...retries.Option[Workspace]) (*Workspace, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	workspace, err := a.CreateWorkspace(ctx, createWorkspaceRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Workspace]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Workspace](ctx, i.Timeout, func() (*Workspace, *retries.Err) {
		workspace, err := a.GetWorkspace(ctx, GetWorkspaceRequest{
			WorkspaceId: workspace.WorkspaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Workspace]{
				Info:    *workspace,
				Timeout: i.Timeout,
			})
		}
		status := workspace.WorkspaceStatus
		statusMessage := workspace.WorkspaceStatusMessage
		switch status {
		case WorkspaceStatusRunning: // target state
			return workspace, nil
		case WorkspaceStatusBanned, WorkspaceStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				WorkspaceStatusRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete workspace
//
// Terminates and deletes a Databricks workspace. From an API perspective,
// deletion is immediate. However, it might take a few minutes for all
// workspaces resources to be deleted, depending on the size and number of
// workspace resources.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) DeleteWorkspace(ctx context.Context, request DeleteWorkspaceRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", request.AccountId, request.WorkspaceId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete workspace
//
// Terminates and deletes a Databricks workspace. From an API perspective,
// deletion is immediate. However, it might take a few minutes for all
// workspaces resources to be deleted, depending on the size and number of
// workspace resources.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) DeleteWorkspaceByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) error {
	return a.DeleteWorkspace(ctx, DeleteWorkspaceRequest{
		AccountId:   accountId,
		WorkspaceId: workspaceId,
	})
}

// Get all workspaces
//
// Gets a list of all workspaces associated with an account, specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetAllWorkspaces(ctx context.Context, request GetAllWorkspacesRequest) ([]Workspace, error) {
	var workspaceList []Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", request.AccountId)
	err := a.client.Get(ctx, path, request, &workspaceList)
	return workspaceList, err
}

func (a *WorkspacesAPI) WorkspaceWorkspaceNameToWorkspaceIdMap(ctx context.Context, request GetAllWorkspacesRequest) (map[string]int64, error) {
	mapping := map[string]int64{}
	result, err := a.GetAllWorkspaces(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.WorkspaceName] = v.WorkspaceId
	}
	return mapping, nil
}

func (a *WorkspacesAPI) GetWorkspaceByWorkspaceName(ctx context.Context, name string) (*Workspace, error) {
	result, err := a.GetAllWorkspaces(ctx, GetAllWorkspacesRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.WorkspaceName != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Workspace named '%s' does not exist", name)
}

// Get all workspaces
//
// Gets a list of all workspaces associated with an account, specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetAllWorkspacesByAccountId(ctx context.Context, accountId string) ([]Workspace, error) {
	return a.GetAllWorkspaces(ctx, GetAllWorkspacesRequest{
		AccountId: accountId,
	})
}

// Get workspace
//
// Gets information including status for a Databricks workspace, specified by
// ID. In the response, the `workspace_status` field indicates the current
// status. After initial workspace creation (which is asynchronous), make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetWorkspace(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", request.AccountId, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &workspace)
	return &workspace, err
}

// Get workspace
//
// Gets information including status for a Databricks workspace, specified by
// ID. In the response, the `workspace_status` field indicates the current
// status. After initial workspace creation (which is asynchronous), make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetWorkspaceByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*Workspace, error) {
	return a.GetWorkspace(ctx, GetWorkspaceRequest{
		AccountId:   accountId,
		WorkspaceId: workspaceId,
	})
}

// Get the history of a workspace's associations with keys
//
// Gets a list of all associations with key configuration objects for the
// specified workspace that encapsulate customer-managed keys that encrypt
// managed services, workspace storage, or in some cases both.
//
// **Important**: In the current implementation, keys cannot be rotated or
// removed from a workspace. It is possible for a workspace to show a storage
// customer-managed key having been attached and then detached if the workspace
// was updated to use the key and the update operation failed.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types and subscription types.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *WorkspacesAPI) GetWorkspaceKeyHistory(ctx context.Context, request GetWorkspaceKeyHistoryRequest) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	var listWorkspaceEncryptionKeyRecordsResponse ListWorkspaceEncryptionKeyRecordsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/customer-managed-key-history", request.AccountId, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &listWorkspaceEncryptionKeyRecordsResponse)
	return &listWorkspaceEncryptionKeyRecordsResponse, err
}

// Get the history of a workspace's associations with keys
//
// Gets a list of all associations with key configuration objects for the
// specified workspace that encapsulate customer-managed keys that encrypt
// managed services, workspace storage, or in some cases both.
//
// **Important**: In the current implementation, keys cannot be rotated or
// removed from a workspace. It is possible for a workspace to show a storage
// customer-managed key having been attached and then detached if the workspace
// was updated to use the key and the update operation failed.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types and subscription types.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *WorkspacesAPI) GetWorkspaceKeyHistoryByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	return a.GetWorkspaceKeyHistory(ctx, GetWorkspaceKeyHistoryRequest{
		AccountId:   accountId,
		WorkspaceId: workspaceId,
	})
}

// Update workspace configuration
//
// Updates a workspace configuration for either a running workspace or a failed
// workspace. The elements that can be updated varies between these two use
// cases.
//
// ### Update a failed workspace You can update a Databricks workspace
// configuration for failed workspace deployment for some fields, but not all
// fields. For a failed workspace, this request supports updates to the
// following fields only: - Credential configuration ID - Storage configuration
// ID - Network configuration ID. Used only if you use customer-managed VPC. -
// Key configuration ID for managed services (control plane storage, such as
// notebook source and Databricks SQL queries). Used only if you use
// customer-managed keys for managed services. - Key configuration ID for
// workspace storage (root S3 bucket and, optionally, EBS volumes). Used only if
// you use customer-managed keys for workspace storage. **Important**: If the
// workspace was ever in the running state, even if briefly before becoming a
// failed workspace, you cannot add a new key configuration ID for workspace
// storage.
//
// After calling the `PATCH` operation to update the workspace configuration,
// make repeated `GET` requests with the workspace ID and check the workspace
// status. The workspace is successful if the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// ### Update a running workspace You can update a Databricks workspace
// configuration for running workspaces for some fields, but not all fields. For
// a running workspace, this request supports updating the following fields
// only: - Credential configuration ID
//
// - Network configuration ID. Used only if you already use use customer-managed
// VPC. This change is supported only if you specified a network configuration
// ID in your original workspace creation. In other words, you cannot switch
// from a Databricks-managed VPC to a customer-managed VPC. **Note**: You cannot
// use a network configuration update in this API to add support for PrivateLink
// (in Public Preview). To add PrivateLink to an existing workspace, contact
// your Databricks representative.
//
// - Key configuration ID for managed services (control plane storage, such as
// notebook source and Databricks SQL queries). Databricks does not directly
// encrypt the data with the customer-managed key (CMK). Databricks uses both
// the CMK and the Databricks managed key (DMK) that is unique to your workspace
// to encrypt the Data Encryption Key (DEK). Databricks uses the DEK to encrypt
// your workspace's managed services persisted data. If the workspace does not
// already have a CMK for managed services, adding this ID enables managed
// services encryption for new or updated data. Existing managed services data
// that existed before adding the key remains not encrypted with the DEK until
// it is modified. If the workspace already has customer-managed keys for
// managed services, this request rotates (changes) the CMK keys and the DEK is
// re-encrypted with the DMK and the new CMK. - Key configuration ID for
// workspace storage (root S3 bucket and, optionally, EBS volumes). You can set
// this only if the workspace does not already have a customer-managed key
// configuration for workspace storage.
//
// **Important**: For updating running workspaces, this API is unavailable on
// Mondays, Tuesdays, and Thursdays from 4:30pm-7:30pm PST due to routine
// maintenance. Plan your workspace updates accordingly. For questions about
// this schedule, contact your Databricks representative.
//
// **Important**: To update a running workspace, your workspace must have no
// running cluster instances, which includes all-purpose clusters, job clusters,
// and pools that might have running clusters. Terminate all cluster instances
// in the workspace before calling this API.
//
// ### Wait until changes take effect. After calling the `PATCH` operation to
// update the workspace configuration, make repeated `GET` requests with the
// workspace ID and check the workspace status and the status of the fields. *
// For workspaces with a Databricks-managed VPC, the workspace status becomes
// `PROVISIONING` temporarily (typically under 20 minutes). If the workspace
// update is successful, the workspace status changes to `RUNNING`. Note that
// you can also check the workspace status in the [Account
// Console](https://docs.databricks.com/administration-guide/account-settings-e2/account-console-e2.html).
// However, you cannot use or create clusters for another 20 minutes after that
// status change. This results in a total of up to 40 minutes in which you
// cannot create clusters. If you create or use clusters before this time
// interval elapses, clusters do not launch successfully, fail, or could cause
// other unexpected behavior.
//
// * For workspaces with a customer-managed VPC, the workspace status stays at
// status `RUNNING` and the VPC change happens immediately. A change to the
// storage customer-managed key configuration ID might take a few minutes to
// update, so continue to check the workspace until you observe that it has been
// updated. If the update fails, the workspace might revert silently to its
// original configuration. After the workspace has been updated, you cannot use
// or create clusters for another 20 minutes. If you create or use clusters
// before this time interval elapses, clusters do not launch successfully, fail,
// or could cause other unexpected behavior.
//
// If you update the _storage_ customer-managed key configurations, it takes 20
// minutes for the changes to fully take effect. During the 20 minute wait, it
// is important that you stop all REST API calls to the DBFS API. If you are
// modifying _only the managed services key configuration_, you can omit the 20
// minute wait.
//
// **Important**: Customer-managed keys and customer-managed VPCs are supported
// by only some deployment types and subscription types. If you have questions
// about availability, contact your Databricks representative.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) PatchWorkspace(ctx context.Context, request UpdateWorkspaceRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", request.AccountId, request.WorkspaceId)
	err := a.client.Patch(ctx, path, request)
	return err
}
