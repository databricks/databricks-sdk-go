// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package customermanagedkeys

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage encryption key configurations for this workspace
// (optional). A key configuration encapsulates the AWS KMS key information and
// some information about how the key configuration can be used. There are two
// possible uses for key configurations: * Managed services: A key configuration
// can be used to encrypt a workspace&#39;s notebook and secret data in the control
// plane, as well as Databricks SQL queries and query history. * Storage: A key
// configuration can be used to encrypt a workspace&#39;s DBFS and EBS data in the
// data plane. In both of these cases, the key configuration&#39;s ID is used when
// creating a new workspace. This Preview feature is available if your account
// is on the E2 version of the platform. Updating a running workspace with
// workspace storage encryption requires that the workspace is on the E2 version
// of the platform. If you have an older workspace, it might not be on the E2
// version of the platform. If you are not sure, contact your Databricks
// reprsentative.
type CustomermanagedkeysService interface {
    // Create a customer-managed key configuration object for an account,
    // specified by ID. This operation uploads a reference to a customer-managed
    // key to Databricks. If the key is assigned as a workspace&#39;s
    // customer-managed key for managed services, Databricks uses the key to
    // encrypt the workspaces notebooks and secrets in the control plane, as
    // well as Databricks SQL queries and query history. If it is specified as a
    // workspace&#39;s customer-managed key for workspace storage, the key encrypts
    // the workspace&#39;s root S3 bucket (which contains the workspace&#39;s root DBFS
    // and system data) and optionally cluster EBS volume data. **Important**:
    // Customer-managed keys are supported only for some deployment types,
    // subscription types, and AWS regions. This operation is available only if
    // your account is on the E2 version of the platform or on a select custom
    // plan that allows multiple workspaces per account.
    CreateKeyConfig(ctx context.Context, createCustomerManagedKeyRequest CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error)
    // Delete a customer-managed key configuration object for an account. You
    // cannot delete a configuration that is associated with a running
    // workspace.
    DeleteKeyConfig(ctx context.Context, deleteKeyConfigRequest DeleteKeyConfigRequest) error
    // Get all customer-managed key configuration objects for an account. If the
    // key is specified as a workspace&#39;s managed services customer-managed key,
    // Databricks will use the key to encrypt the workspace&#39;s notebooks and
    // secrets in the control plane, as well as Databricks SQL queries and query
    // history. If the key is specified as a workspace&#39;s storage
    // customer-managed key, the key is used to encrypt the workspace&#39;s root S3
    // bucket and optionally can encrypt cluster EBS volumes data in the data
    // plane. **Important**: Customer-managed keys are supported only for some
    // deployment types, subscription types, and AWS regions. This operation is
    // available only if your account is on the E2 version of the platform.
    GetAllKeyConfigs(ctx context.Context, getAllKeyConfigsRequest GetAllKeyConfigsRequest) (*ListCustomerManagedKeysResponse, error)
    // Get a customer-managed key configuration object for an account, specified
    // by ID. This operation uploads a reference to a customer-managed key to
    // Databricks. If assigned as a workspace&#39;s customer-managed key for managed
    // services, Databricks uses the key to encrypt the workspaces notebooks and
    // secrets in the control plane, as well as Databricks SQL queries and query
    // history. If it is specified as a workspace&#39;s customer-managed key for
    // storage, the key encrypts the workspace&#39;s root S3 bucket (which contains
    // the workspace&#39;s root DBFS and system data) and optionally cluster EBS
    // volume data. **Important**: Customer-managed keys are supported only for
    // some deployment types, subscription types, and AWS regions. This
    // operation is available only if your account is on the E2 version of the
    // platform.
    GetKeyConfig(ctx context.Context, getKeyConfigRequest GetKeyConfigRequest) (*CustomerManagedKey, error)
    // Get a list of records of how key configurations were associated with
    // workspaces. **Important**: Customer-managed keys are supported only for
    // some deployment types, subscription types, and AWS regions. This
    // operation is available only if your account is on the E2 version of the
    // platform.
    GetKeyWorkspaceHistory(ctx context.Context, getKeyWorkspaceHistoryRequest GetKeyWorkspaceHistoryRequest) (*ListWorkspaceEncryptionKeyRecordsResponse, error)
	GetAllKeyConfigsByAccountId(ctx context.Context, accountId string) (*ListCustomerManagedKeysResponse, error)
	GetKeyWorkspaceHistoryByAccountId(ctx context.Context, accountId string) (*ListWorkspaceEncryptionKeyRecordsResponse, error)
	GetKeyConfigByAccountIdAndCustomerManagedKeyId(ctx context.Context, accountId string, customerManagedKeyId string) (*CustomerManagedKey, error)
	DeleteKeyConfigByAccountIdAndCustomerManagedKeyId(ctx context.Context, accountId string, customerManagedKeyId string) error
}

func New(client *client.DatabricksClient) CustomermanagedkeysService {
	return &CustomermanagedkeysAPI{
		client: client,
	}
}

type CustomermanagedkeysAPI struct {
	client *client.DatabricksClient
}

// Create a customer-managed key configuration object for an account, specified
// by ID. This operation uploads a reference to a customer-managed key to
// Databricks. If the key is assigned as a workspace&#39;s customer-managed key for
// managed services, Databricks uses the key to encrypt the workspaces notebooks
// and secrets in the control plane, as well as Databricks SQL queries and query
// history. If it is specified as a workspace&#39;s customer-managed key for
// workspace storage, the key encrypts the workspace&#39;s root S3 bucket (which
// contains the workspace&#39;s root DBFS and system data) and optionally cluster
// EBS volume data. **Important**: Customer-managed keys are supported only for
// some deployment types, subscription types, and AWS regions. This operation is
// available only if your account is on the E2 version of the platform or on a
// select custom plan that allows multiple workspaces per account.
func (a *CustomermanagedkeysAPI) CreateKeyConfig(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/accounts/%v/customer-managed-keys", request.AccountId)
	err := a.client.Post(ctx, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

// Delete a customer-managed key configuration object for an account. You cannot
// delete a configuration that is associated with a running workspace.
func (a *CustomermanagedkeysAPI) DeleteKeyConfig(ctx context.Context, request DeleteKeyConfigRequest) error {
	path := fmt.Sprintf("/accounts/%v/customer-managed-keys/%v", request.AccountId, request.CustomerManagedKeyId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get all customer-managed key configuration objects for an account. If the key
// is specified as a workspace&#39;s managed services customer-managed key,
// Databricks will use the key to encrypt the workspace&#39;s notebooks and secrets
// in the control plane, as well as Databricks SQL queries and query history. If
// the key is specified as a workspace&#39;s storage customer-managed key, the key
// is used to encrypt the workspace&#39;s root S3 bucket and optionally can encrypt
// cluster EBS volumes data in the data plane. **Important**: Customer-managed
// keys are supported only for some deployment types, subscription types, and
// AWS regions. This operation is available only if your account is on the E2
// version of the platform.
func (a *CustomermanagedkeysAPI) GetAllKeyConfigs(ctx context.Context, request GetAllKeyConfigsRequest) (*ListCustomerManagedKeysResponse, error) {
	var listCustomerManagedKeysResponse ListCustomerManagedKeysResponse
	path := fmt.Sprintf("/accounts/%v/customer-managed-keys", request.AccountId)
	err := a.client.Get(ctx, path, request, &listCustomerManagedKeysResponse)
	return &listCustomerManagedKeysResponse, err
}

// Get a customer-managed key configuration object for an account, specified by
// ID. This operation uploads a reference to a customer-managed key to
// Databricks. If assigned as a workspace&#39;s customer-managed key for managed
// services, Databricks uses the key to encrypt the workspaces notebooks and
// secrets in the control plane, as well as Databricks SQL queries and query
// history. If it is specified as a workspace&#39;s customer-managed key for
// storage, the key encrypts the workspace&#39;s root S3 bucket (which contains the
// workspace&#39;s root DBFS and system data) and optionally cluster EBS volume
// data. **Important**: Customer-managed keys are supported only for some
// deployment types, subscription types, and AWS regions. This operation is
// available only if your account is on the E2 version of the platform.
func (a *CustomermanagedkeysAPI) GetKeyConfig(ctx context.Context, request GetKeyConfigRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/accounts/%v/customer-managed-keys/%v", request.AccountId, request.CustomerManagedKeyId)
	err := a.client.Get(ctx, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

// Get a list of records of how key configurations were associated with
// workspaces. **Important**: Customer-managed keys are supported only for some
// deployment types, subscription types, and AWS regions. This operation is
// available only if your account is on the E2 version of the platform.
func (a *CustomermanagedkeysAPI) GetKeyWorkspaceHistory(ctx context.Context, request GetKeyWorkspaceHistoryRequest) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	var listWorkspaceEncryptionKeyRecordsResponse ListWorkspaceEncryptionKeyRecordsResponse
	path := fmt.Sprintf("/accounts/%v/customer-managed-key-history", request.AccountId)
	err := a.client.Get(ctx, path, request, &listWorkspaceEncryptionKeyRecordsResponse)
	return &listWorkspaceEncryptionKeyRecordsResponse, err
}


func (a *CustomermanagedkeysAPI) GetAllKeyConfigsByAccountId(ctx context.Context, accountId string) (*ListCustomerManagedKeysResponse, error) {
	return a.GetAllKeyConfigs(ctx, GetAllKeyConfigsRequest{
		AccountId: accountId,
	})
}

func (a *CustomermanagedkeysAPI) GetKeyWorkspaceHistoryByAccountId(ctx context.Context, accountId string) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	return a.GetKeyWorkspaceHistory(ctx, GetKeyWorkspaceHistoryRequest{
		AccountId: accountId,
	})
}

func (a *CustomermanagedkeysAPI) GetKeyConfigByAccountIdAndCustomerManagedKeyId(ctx context.Context, accountId string, customerManagedKeyId string) (*CustomerManagedKey, error) {
	return a.GetKeyConfig(ctx, GetKeyConfigRequest{
		AccountId: accountId,
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

func (a *CustomermanagedkeysAPI) DeleteKeyConfigByAccountIdAndCustomerManagedKeyId(ctx context.Context, accountId string, customerManagedKeyId string) error {
	return a.DeleteKeyConfig(ctx, DeleteKeyConfigRequest{
		AccountId: accountId,
		CustomerManagedKeyId: customerManagedKeyId,
	})
}
