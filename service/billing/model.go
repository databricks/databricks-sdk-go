// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

// all definitions in this file are in alphabetical order

type BudgetAlert struct {
	// List of email addresses to be notified when budget percentage is exceeded
	// in the given period.
	EmailNotifications []string `json:"email_notifications,omitempty"`
	// Percentage of the target amount used in the currect period that will
	// trigger a notification.
	MinPercentage int `json:"min_percentage,omitempty"`
}

// Budget configuration to be created.
type BudgetCreateRequest struct {
	Alerts []BudgetAlert `json:"alerts,omitempty"`
	// Optional end date of the budget.
	EndDate string `json:"end_date,omitempty"`
	// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
	// account that matches this expression will be counted in this budget.
	//
	// Supported properties on left-hand side of comparison: * `workspaceId` -
	// the ID of the workspace * `sku` - SKU of the cluster, e.g.
	// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of
	// the cluster
	//
	// Supported comparison operators: * `=` - equal * `!=` - not equal
	//
	// Supported logical operators: `AND`, `OR`.
	//
	// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
	// tag.'my tag'='my value')` * `workspaceId!=456` *
	// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'`
	// * `tag.name1='value1' AND tag.name2='value2'`
	Filter string `json:"filter"`
	// Human-readable name of the budget.
	Name string `json:"name"`
	// Period length in years, months, weeks and/or days. Examples: `1 month`,
	// `30 days`, `1 year, 2 months, 1 week, 2 days`
	Period string `json:"period"`
	// Start date of the budget period calculation.
	StartDate string `json:"start_date"`
	// Target amount of the budget per period in USD.
	TargetAmount string `json:"target_amount"`
}

// List of budgets.
type BudgetList struct {
	Budgets []BudgetWithStatus `json:"budgets,omitempty"`
}

// Budget configuration with daily status.
type BudgetWithStatus struct {
	Alerts []BudgetAlert `json:"alerts,omitempty"`

	BudgetId string `json:"budget_id,omitempty"`

	CreationTime string `json:"creation_time,omitempty"`
	// Optional end date of the budget.
	EndDate string `json:"end_date,omitempty"`
	// SQL-like filter expression with workspaceId, SKU and tag. Usage in your
	// account that matches this expression will be counted in this budget.
	//
	// Supported properties on left-hand side of comparison: * `workspaceId` -
	// the ID of the workspace * `sku` - SKU of the cluster, e.g.
	// `STANDARD_ALL_PURPOSE_COMPUTE` * `tag.tagName`, `tag.'tag name'` - tag of
	// the cluster
	//
	// Supported comparison operators: * `=` - equal * `!=` - not equal
	//
	// Supported logical operators: `AND`, `OR`.
	//
	// Examples: * `workspaceId=123 OR (sku='STANDARD_ALL_PURPOSE_COMPUTE' AND
	// tag.'my tag'='my value')` * `workspaceId!=456` *
	// `sku='STANDARD_ALL_PURPOSE_COMPUTE' OR sku='PREMIUM_ALL_PURPOSE_COMPUTE'`
	// * `tag.name1='value1' AND tag.name2='value2'`
	Filter string `json:"filter,omitempty"`
	// Human-readable name of the budget.
	Name string `json:"name,omitempty"`
	// Period length in years, months, weeks and/or days. Examples: `1 month`,
	// `30 days`, `1 year, 2 months, 1 week, 2 days`
	Period string `json:"period,omitempty"`
	// Start date of the budget period calculation.
	StartDate string `json:"start_date,omitempty"`
	// Amount used in the budget for each day (noncumulative).
	StatusDaily []BudgetWithStatusStatusDailyItem `json:"status_daily,omitempty"`
	// Target amount of the budget per period in USD.
	TargetAmount string `json:"target_amount,omitempty"`

	UpdateTime string `json:"update_time,omitempty"`
}

type BudgetWithStatusStatusDailyItem struct {
	// Amount used in this day in USD.
	Amount string `json:"amount,omitempty"`

	Date string `json:"date,omitempty"`
}

type CreateBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget configuration to be created.
	Budget BudgetCreateRequest `json:"budget"`
}

type CreateLogDeliveryConfigurationParams struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `json:"config_name,omitempty"`
	// The ID for a [Databricks credential
	// configuration](#operation/create-credential-config) that represents the
	// AWS IAM role with policy and trust relationship as described in the main
	// billable usage documentation page. See [Configure billable usage
	// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
	CredentialsId string `json:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log
	// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
	// For the CSV schema, see the [View billable
	// usage](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	//
	// * `AUDIT_LOGS` — Configure [audit log
	// delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
	// For the JSON schema, see [Configure audit
	// logging](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html)
	LogType LogType `json:"log_type"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable
	// usage](https://docs.databricks.com/administration-guide/account-settings/usage.html)
	// * If `log_type` is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON
	// (JavaScript Object Notation) format is supported. For the schema, see the
	// [Configuring audit
	// logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
	OutputFormat OutputFormat `json:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// "The ID for a [Databricks storage
	// configuration](#operation/create-storage-config) that represents the S3
	// bucket with bucket policy as described in the main billable usage
	// documentation page. See [Configure billable usage
	// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html)."
	StorageConfigurationId string `json:"storage_configuration_id"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`
}

type DeleteBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget ID
	BudgetId string `json:"-" path:"budget_id"`
}

// This describes an enum
type DeliveryStatus string

// There were no log delivery attempts since the config was created.
const DeliveryStatusCreated DeliveryStatus = `CREATED`

// The log delivery status as the configuration has been disabled since the
// release of this feature or there are no workspaces in the account.
const DeliveryStatusNotFound DeliveryStatus = `NOT_FOUND`

// The latest attempt of log delivery has succeeded completely.
const DeliveryStatusSucceeded DeliveryStatus = `SUCCEEDED`

// The latest attempt of log delivery failed because of an Databricks internal
// error. Contact support if it doesn't go away soon.
const DeliveryStatusSystemFailure DeliveryStatus = `SYSTEM_FAILURE`

// The latest attempt of log delivery failed because of misconfiguration of
// customer provided permissions on role or storage.
const DeliveryStatusUserFailure DeliveryStatus = `USER_FAILURE`

type DownloadBillableUsageRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth any/* ERROR */ `json:"-" url:"end_month,omitempty"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool `json:"-" url:"personal_data,omitempty"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	StartMonth any/* ERROR */ `json:"-" url:"start_month,omitempty"`
}

type GetAllBudgetsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
}

type GetAllLogDeliveryConfigsRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Filter by credential configuration ID.
	CredentialsId string `json:"-" url:"credentials_id,omitempty"`
	// Filter by status `ENABLED` or `DISABLED`.
	Status LogDeliveryConfigStatus `json:"-" url:"status,omitempty"`
	// Filter by storage configuration ID.
	StorageConfigurationId string `json:"-" url:"storage_configuration_id,omitempty"`
}

type GetBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget ID
	BudgetId string `json:"-" path:"budget_id"`
}

type GetLogDeliveryConfigRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `json:"-" path:"log_delivery_configuration_id"`
}

// Status of log delivery configuration. Set to `ENABLED` (enabled) or
// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable the
// configuration](#operation/patch-log-delivery-config-status) later. Deletion
// of a configuration is not supported, so disable a log delivery configuration
// that is no longer needed.
type LogDeliveryConfigStatus string

const LogDeliveryConfigStatusDisabled LogDeliveryConfigStatus = `DISABLED`

const LogDeliveryConfigStatusEnabled LogDeliveryConfigStatus = `ENABLED`

type LogDeliveryConfiguration struct {
	// The Databricks account ID that hosts the log delivery configuration.
	AccountId string `json:"account_id,omitempty"`
	// Databricks log delivery configuration ID.
	ConfigId string `json:"config_id,omitempty"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `json:"config_name,omitempty"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The ID for a [Databricks credential
	// configuration](#operation/create-credential-config) that represents the
	// AWS IAM role with policy and trust relationship as described in the main
	// billable usage documentation page. See [Configure billable usage
	// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
	CredentialsId string `json:"credentials_id,omitempty"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// Databricks log delivery status.
	LogDeliveryStatus *LogDeliveryStatus `json:"log_delivery_status,omitempty"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log
	// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
	// For the CSV schema, see the [View billable
	// usage](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	//
	// * `AUDIT_LOGS` — Configure [audit log
	// delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
	// For the JSON schema, see [Configure audit
	// logging](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html)
	LogType LogType `json:"log_type,omitempty"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable
	// usage](https://docs.databricks.com/administration-guide/account-settings/usage.html)
	// * If `log_type` is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON
	// (JavaScript Object Notation) format is supported. For the schema, see the
	// [Configuring audit
	// logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
	OutputFormat OutputFormat `json:"output_format,omitempty"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// "The ID for a [Databricks storage
	// configuration](#operation/create-storage-config) that represents the S3
	// bucket with bucket policy as described in the main billable usage
	// documentation page. See [Configure billable usage
	// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html)."
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	UpdateTime int64 `json:"update_time,omitempty"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`
}

// Databricks log delivery status.
type LogDeliveryStatus struct {
	// The UTC time for the latest log delivery attempt.
	LastAttemptTime string `json:"last_attempt_time,omitempty"`
	// The UTC time for the latest successful log delivery.
	LastSuccessfulAttemptTime string `json:"last_successful_attempt_time,omitempty"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	Message string `json:"message,omitempty"`
	// This describes an enum
	Status DeliveryStatus `json:"status,omitempty"`
}

// Log delivery type. Supported values are:
//
// * `BILLABLE_USAGE` — Configure [billable usage log
// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
// For the CSV schema, see the [View billable
// usage](https://docs.databricks.com/administration-guide/account-settings/usage.html).
//
// * `AUDIT_LOGS` — Configure [audit log
// delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
// For the JSON schema, see [Configure audit
// logging](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html)
type LogType string

const LogTypeAuditLogs LogType = `AUDIT_LOGS`

const LogTypeBillableUsage LogType = `BILLABLE_USAGE`

type ModifyBudgetRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Budget configuration to be created.
	Budget BudgetCreateRequest `json:"budget"`
	// Budget ID
	BudgetId string `json:"-" path:"budget_id"`
}

// The file type of log delivery.
//
// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the CSV
// (comma-separated values) format is supported. For the schema, see the [View
// billable
// usage](https://docs.databricks.com/administration-guide/account-settings/usage.html)
// * If `log_type` is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON
// (JavaScript Object Notation) format is supported. For the schema, see the
// [Configuring audit
// logs](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
type OutputFormat string

const OutputFormatCsv OutputFormat = `CSV`

const OutputFormatJson OutputFormat = `JSON`

type UpdateLogDeliveryConfigurationStatusRequest struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `json:"-" path:"log_delivery_configuration_id"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status"`
}

type WrappedCreateLogDeliveryConfiguration struct {
	// Databricks account ID of any type. For non-E2 account types, get your
	// account ID from the [Accounts
	// Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
	AccountId string `json:"-" path:"account_id"`

	LogDeliveryConfiguration *CreateLogDeliveryConfigurationParams `json:"log_delivery_configuration,omitempty"`
}

type WrappedLogDeliveryConfiguration struct {
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

type WrappedLogDeliveryConfigurations struct {
	LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
}
