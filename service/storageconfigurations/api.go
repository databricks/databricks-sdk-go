// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package storageconfigurations

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage storage configurations for this workspace. A root storage
// S3 bucket in your account is required to store objects like cluster logs,
// notebook revisions, and job results. You can also use the root storage S3
// bucket for storage of non-production DBFS data. A storage configuration
// encapsulates this bucket information, and its ID is used when creating a new
// workspace. 
type StorageconfigurationsService interface {
    // Create new storage configuration for an account, specified by ID. Uploads
    // a storage configuration object that represents the root AWS S3 bucket in
    // your account. Databricks stores related workspace assets including DBFS,
    // cluster logs, and job results. For AWS S3 bucket, you need to configure
    // the required bucket policy. For detailed instructions of creating a new
    // workspace with this API, see [Create a new workspace using the Account
    // API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
    CreateStorageConfig(ctx context.Context, createStorageConfigurationRequest CreateStorageConfigurationRequest) (*StorageConfiguration, error)
    // Delete a Databricks storage configuration. You cannot delete a storage
    // configuration that is currently being associated to any workspace.
    DeleteStorageConfig(ctx context.Context, deleteStorageConfigRequest DeleteStorageConfigRequest) error
    // Get a list of all Databricks storage configurations for your account,
    // specified by ID.
    GetAllStorageConfigs(ctx context.Context, getAllStorageConfigsRequest GetAllStorageConfigsRequest) (*ListStorageConfigurationsResponse, error)
    // Get a Databricks storage configuration for an account, both specified by
    // ID.
    GetStorageConfig(ctx context.Context, getStorageConfigRequest GetStorageConfigRequest) (*StorageConfiguration, error)
	GetStorageConfigByAccountIdAndStorageConfigurationId(ctx context.Context, accountId string, storageConfigurationId string) (*StorageConfiguration, error)
	DeleteStorageConfigByAccountIdAndStorageConfigurationId(ctx context.Context, accountId string, storageConfigurationId string) error
	GetAllStorageConfigsByAccountId(ctx context.Context, accountId string) (*ListStorageConfigurationsResponse, error)
}

func New(client *client.DatabricksClient) StorageconfigurationsService {
	return &StorageconfigurationsAPI{
		client: client,
	}
}

type StorageconfigurationsAPI struct {
	client *client.DatabricksClient
}

// Create new storage configuration for an account, specified by ID. Uploads a
// storage configuration object that represents the root AWS S3 bucket in your
// account. Databricks stores related workspace assets including DBFS, cluster
// logs, and job results. For AWS S3 bucket, you need to configure the required
// bucket policy. For detailed instructions of creating a new workspace with
// this API, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *StorageconfigurationsAPI) CreateStorageConfig(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/accounts/%v/storage-configurations", request.AccountId)
	err := a.client.Post(ctx, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

// Delete a Databricks storage configuration. You cannot delete a storage
// configuration that is currently being associated to any workspace.
func (a *StorageconfigurationsAPI) DeleteStorageConfig(ctx context.Context, request DeleteStorageConfigRequest) error {
	path := fmt.Sprintf("/accounts/%v/storage-configurations/%v", request.AccountId, request.StorageConfigurationId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Get a list of all Databricks storage configurations for your account,
// specified by ID.
func (a *StorageconfigurationsAPI) GetAllStorageConfigs(ctx context.Context, request GetAllStorageConfigsRequest) (*ListStorageConfigurationsResponse, error) {
	var listStorageConfigurationsResponse ListStorageConfigurationsResponse
	path := fmt.Sprintf("/accounts/%v/storage-configurations", request.AccountId)
	err := a.client.Get(ctx, path, request, &listStorageConfigurationsResponse)
	return &listStorageConfigurationsResponse, err
}

// Get a Databricks storage configuration for an account, both specified by ID.
func (a *StorageconfigurationsAPI) GetStorageConfig(ctx context.Context, request GetStorageConfigRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/accounts/%v/storage-configurations/%v", request.AccountId, request.StorageConfigurationId)
	err := a.client.Get(ctx, path, request, &storageConfiguration)
	return &storageConfiguration, err
}


func (a *StorageconfigurationsAPI) GetStorageConfigByAccountIdAndStorageConfigurationId(ctx context.Context, accountId string, storageConfigurationId string) (*StorageConfiguration, error) {
	return a.GetStorageConfig(ctx, GetStorageConfigRequest{
		AccountId: accountId,
		StorageConfigurationId: storageConfigurationId,
	})
}

func (a *StorageconfigurationsAPI) DeleteStorageConfigByAccountIdAndStorageConfigurationId(ctx context.Context, accountId string, storageConfigurationId string) error {
	return a.DeleteStorageConfig(ctx, DeleteStorageConfigRequest{
		AccountId: accountId,
		StorageConfigurationId: storageConfigurationId,
	})
}

func (a *StorageconfigurationsAPI) GetAllStorageConfigsByAccountId(ctx context.Context, accountId string) (*ListStorageConfigurationsResponse, error) {
	return a.GetAllStorageConfigs(ctx, GetAllStorageConfigsRequest{
		AccountId: accountId,
	})
}
