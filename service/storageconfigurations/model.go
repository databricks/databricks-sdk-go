package storageconfigurations

// RootBucketInfo points to a bucket name
type RootBucketInfo struct {
	BucketName string `json:"bucket_name,omitempty"`
}

// StorageConfiguration is the object that contains all the information for the root storage bucket
type StorageConfiguration struct {
	StorageConfigurationID   string          `json:"storage_configuration_id,omitempty"`
	StorageConfigurationName string          `json:"storage_configuration_name,omitempty"`
	RootBucketInfo           *RootBucketInfo `json:"root_bucket_info,omitempty"`
	AccountID                string          `json:"account_id,omitempty"`
	CreationTime             int64           `json:"creation_time,omitempty"`
}