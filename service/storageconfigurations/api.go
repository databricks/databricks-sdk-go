package storageconfigurations

import (
	"context"
	"fmt"

	"github.com/databricks/sdk-go/databricks"
	"github.com/databricks/sdk-go/databricks/client"
)

// New creates ClustersAPI instance from provider meta
func New(c ...*databricks.Config) *StorageConfigurationsAPI {
	return &StorageConfigurationsAPI{
		client: client.New(c...),
	}
}

type StorageConfigurationsAPI struct {
	client *client.DatabricksClient
}

// Create creates a configuration for the root s3 bucket
func (a StorageConfigurationsAPI) Create(ctx context.Context, storageConfigurationName string, bucketName string) (StorageConfiguration, error) {
	var conf StorageConfiguration
	path := fmt.Sprintf("/accounts/%s/storage-configurations", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, StorageConfiguration{
		StorageConfigurationName: storageConfigurationName,
		RootBucketInfo: &RootBucketInfo{
			BucketName: bucketName,
		},
	}, &conf)
	return conf, err
}

// Read returns the configuration for the root s3 bucket and metadata for the storage configuration
func (a StorageConfigurationsAPI) Read(ctx context.Context, storageConfigurationID string) (StorageConfiguration, error) {
	var conf StorageConfiguration
	path := fmt.Sprintf("/accounts/%s/storage-configurations/%s", a.client.Config.AccountID, storageConfigurationID)
	err := a.client.Get(ctx, path, nil, &conf)
	return conf, err
}

// Delete deletes the configuration for the root s3 bucket
func (a StorageConfigurationsAPI) Delete(ctx context.Context, storageConfigurationID string) error {
	storageConfigurationAPIPath := fmt.Sprintf("/accounts/%s/storage-configurations/%s", a.client.Config.AccountID, storageConfigurationID)
	return a.client.Delete(ctx, storageConfigurationAPIPath, nil)
}

// List lists all the storage configurations for the root s3 buckets in the account ID provided to the client config
func (a StorageConfigurationsAPI) List(ctx context.Context) ([]StorageConfiguration, error) {
	var confList []StorageConfiguration
	path := fmt.Sprintf("/accounts/%s/storage-configurations", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &confList)
	return confList, err
}
