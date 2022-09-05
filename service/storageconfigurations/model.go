// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package storageconfigurations

// all definitions in this file are in alphabetical order

type CreateStorageConfigurationRequest struct {
	// Databricks account ID. When you create or manage workspaces, your account
	// must be on the E2 version of the platform or on a select custom plan that
	// allows multiple workspaces per account. If you are configuring log
	// delivery, all account types are supported. For non-E2 account types, get
	// your account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Root S3 bucket information.
	RootBucketInfo CreateStorageConfigurationRequestRootBucketInfo `json:"root_bucket_info"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name"`
}

// Root S3 bucket information.
type CreateStorageConfigurationRequestRootBucketInfo struct {
	// The name of the S3 bucket.
	BucketName string `json:"bucket_name,omitempty"`
}

type DeleteStorageConfigRequest struct {
	// Databricks account ID. When you create or manage workspaces, your account
	// must be on the E2 version of the platform or on a select custom plan that
	// allows multiple workspaces per account. If you are configuring log
	// delivery, all account types are supported. For non-E2 account types, get
	// your account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string ` path:"storage_configuration_id"`
}

type GetAllStorageConfigsRequest struct {
	// Databricks account ID. When you create or manage workspaces, your account
	// must be on the E2 version of the platform or on a select custom plan that
	// allows multiple workspaces per account. If you are configuring log
	// delivery, all account types are supported. For non-E2 account types, get
	// your account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
}

type GetStorageConfigRequest struct {
	// Databricks account ID. When you create or manage workspaces, your account
	// must be on the E2 version of the platform or on a select custom plan that
	// allows multiple workspaces per account. If you are configuring log
	// delivery, all account types are supported. For non-E2 account types, get
	// your account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string ` path:"account_id"`
	// Databricks Account API storage configuration ID.
	StorageConfigurationId string ` path:"storage_configuration_id"`
}

// Storage configuration objects.
type ListStorageConfigurationsResponse []StorageConfiguration

type StorageConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when the storage configuration was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Root S3 bucket information.
	RootBucketInfo *StorageConfigurationRootBucketInfo `json:"root_bucket_info,omitempty"`
	// Databricks storage configuration ID.
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// The human-readable name of the storage configuration.
	StorageConfigurationName string `json:"storage_configuration_name,omitempty"`
}

// Root S3 bucket information.
type StorageConfigurationRootBucketInfo struct {
	// The name of the S3 bucket.
	BucketName string `json:"bucket_name,omitempty"`
}
