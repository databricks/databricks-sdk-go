// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package storageconfigurations

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

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
