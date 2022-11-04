// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package customermanagedkeys

// all definitions in this file are in alphabetical order

type AwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN).
	KeyArn string `json:"key_arn"`
	// The AWS KMS key region.
	KeyRegion string `json:"key_region"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. If you do not want to use this key for encrypting
	// EBS volumes, set to `false`.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`
}

type CreateAwsKeyInfo struct {
	// The AWS KMS key alias.
	KeyAlias string `json:"key_alias,omitempty"`
	// The AWS KMS key's Amazon Resource Name (ARN). Note that the key's AWS
	// region is inferred from the ARN.
	KeyArn string `json:"key_arn"`
	// This field applies only if the `use_cases` property includes `STORAGE`.
	// If this is set to `true` or omitted, the key is also used to encrypt
	// cluster EBS volumes. To not use this key also for encrypting EBS volumes,
	// set this to `false`.
	ReuseKeyForClusterVolumes bool `json:"reuse_key_for_cluster_volumes,omitempty"`
}

type CreateCustomerManagedKeyRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`

	AwsKeyInfo CreateAwsKeyInfo `json:"aws_key_info"`
	// The cases that the key can be used for. Include one or both of these
	// options: * `MANAGED_SERVICES`: Encrypts notebook and secret data in the
	// control plane * `STORAGE`: Encrypts the workspace's root S3 bucket (root
	// DBFS and system data) and, optionally, cluster EBS volumes.
	UseCases []string `json:"use_cases"`
}

type CustomerManagedKey struct {
	// The Databricks account ID that holds the customer-managed key.
	AccountId string `json:"account_id,omitempty"`

	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`
	// Time in epoch milliseconds when the customer key was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`
	// The cases that the key can be used for. Include one or both of these
	// options: * `MANAGED_SERVICES`: Encrypts notebook and secret data in the
	// control plane * `STORAGE`: Encrypts the workspace's root S3 bucket (root
	// DBFS and system data) and optionally cluster EBS volumes.
	UseCases []string `json:"use_cases,omitempty"`
}

type DeleteKeyConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" path:"customer_managed_key_id"`
}

type GetAllKeyConfigsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetKeyConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks encryption key configuration ID.
	CustomerManagedKeyId string `json:"-" path:"customer_managed_key_id"`
}

type GetKeyWorkspaceHistoryRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type KeyStatus string

const KeyStatusAttached KeyStatus = `ATTACHED`

const KeyStatusDetached KeyStatus = `DETACHED`

const KeyStatusUnknown KeyStatus = `UNKNOWN`

// This describes an enum
type KeyUseCase string

// Encrypts notebook and secret data in the control plane
const KeyUseCaseManagedServices KeyUseCase = `MANAGED_SERVICES`

// Encrypts the workspace's root S3 bucket (root DBFS and system data) and,
// optionally, cluster EBS volumes.
const KeyUseCaseStorage KeyUseCase = `STORAGE`

type ListWorkspaceEncryptionKeyRecordsResponse struct {
	WorkspaceEncryptionKeyRecords []WorkspaceEncryptionKeyRecord `json:"workspaceEncryptionKeyRecords,omitempty"`
}

type WorkspaceEncryptionKeyRecord struct {
	AwsKeyInfo *AwsKeyInfo `json:"aws_key_info,omitempty"`
	// ID of the encryption key configuration object.
	CustomerManagedKeyId string `json:"customer_managed_key_id,omitempty"`

	KeyStatus KeyStatus `json:"key_status,omitempty"`
	// ID for the workspace-key mapping record.
	RecordId string `json:"record_id,omitempty"`
	// Time in epoch milliseconds when the record was added.
	UpdateTime int64 `json:"update_time,omitempty"`
	// This describes an enum
	UseCase KeyUseCase `json:"use_case,omitempty"`
	// Workspace ID.
	WorkspaceId int64 `json:"workspace_id,omitempty"`
}
