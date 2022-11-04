// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package customermanagedkeys

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

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
