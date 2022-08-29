// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package logdelivery

// all definitions in this file are in alphabetical order

type GetAllLogDeliveryConfigsRequest struct {
    // Databricks account ID of any type. For non-E2 account types, get your
    // account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    // Filter by credential configuration ID.
    CredentialsId string ` url:"credentials_id,omitempty"`
    // Filter by status `ENABLED` or `DISABLED`.
    Status LogDeliveryConfigStatus ` url:"status,omitempty"`
    // Filter by storage configuration ID.
    StorageConfigurationId string ` url:"storage_configuration_id,omitempty"`
}


type GetLogDeliveryConfigRequest struct {
    // Databricks account ID of any type. For non-E2 account types, get your
    // account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    // Databricks log delivery configuration ID
    LogDeliveryConfigurationId string ` path:"log_delivery_configuration_id"`
}

// Status of log delivery configuration. Set to `ENABLED` (enabled) or
// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable the
// configuration](#operation/patch-log-delivery-config-status) later. Deletion
// of a configuration is not supported, so disable a log delivery configuration
// that is no longer needed.
type LogDeliveryConfigStatus string


const LogDeliveryConfigStatusEnabled LogDeliveryConfigStatus = `ENABLED`

const LogDeliveryConfigStatusDisabled LogDeliveryConfigStatus = `DISABLED`

type LogDeliveryConfiguration struct {
    // The Databricks account ID that hosts the log delivery configuration.
    AccountId string `json:"account_id,omitempty"`
    // Databricks log delivery configuration ID.
    ConfigId string `json:"config_id,omitempty"`
    // Time in epoch milliseconds when the log delivery configuration was
    // created.
    CreationTime int64 `json:"creation_time,omitempty"`
    // Databricks log delivery status.
    LogDeliveryStatus *LogDeliveryConfigurationLogDeliveryStatus `json:"log_delivery_status,omitempty"`
    // Time in epoch milliseconds when the log delivery configuration was
    // updated.
    UpdateTime int64 `json:"update_time,omitempty"`
}

// Databricks log delivery status.
type LogDeliveryConfigurationLogDeliveryStatus struct {
    // The UTC time for the latest log delivery attempt.
    LastAttemptTime string `json:"last_attempt_time,omitempty"`
    // The UTC time for the latest successful log delivery.
    LastSuccessfulAttemptTime string `json:"last_successful_attempt_time,omitempty"`
    // Informative message about the latest log delivery attempt. If the log
    // delivery fails with USER_FAILURE, error details will be provided for
    // fixing misconfigurations in cloud permissions.
    Message string `json:"message,omitempty"`
    // The status string for log delivery. Possible values are: * `CREATED`:
    // There were no log delivery attempts since the config was created. *
    // `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
    // * `USER_FAILURE`: The latest attempt of log delivery failed because of
    // misconfiguration of customer provided permissions on role or storage. *
    // `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
    // Databricks internal error. Contact support if it doesn&#39;t go away soon. *
    // `NOT_FOUND`: The log delivery status as the configuration has been
    // disabled since the release of this feature or there are no workspaces in
    // the account.
    Status LogDeliveryConfigurationLogDeliveryStatusStatus `json:"status,omitempty"`
}

// The status string for log delivery. Possible values are: * `CREATED`: There
// were no log delivery attempts since the config was created. * `SUCCEEDED`:
// The latest attempt of log delivery has succeeded completely. *
// `USER_FAILURE`: The latest attempt of log delivery failed because of
// misconfiguration of customer provided permissions on role or storage. *
// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
// Databricks internal error. Contact support if it doesn&#39;t go away soon. *
// `NOT_FOUND`: The log delivery status as the configuration has been disabled
// since the release of this feature or there are no workspaces in the account.
type LogDeliveryConfigurationLogDeliveryStatusStatus string


const LogDeliveryConfigurationLogDeliveryStatusStatusCreated LogDeliveryConfigurationLogDeliveryStatusStatus = `CREATED`

const LogDeliveryConfigurationLogDeliveryStatusStatusSucceeded LogDeliveryConfigurationLogDeliveryStatusStatus = `SUCCEEDED`

const LogDeliveryConfigurationLogDeliveryStatusStatusUserFailure LogDeliveryConfigurationLogDeliveryStatusStatus = `USER_FAILURE`

const LogDeliveryConfigurationLogDeliveryStatusStatusSystemFailure LogDeliveryConfigurationLogDeliveryStatusStatus = `SYSTEM_FAILURE`

const LogDeliveryConfigurationLogDeliveryStatusStatusNotFound LogDeliveryConfigurationLogDeliveryStatusStatus = `NOT_FOUND`

type UpdateLogDeliveryConfigurationStatusRequest struct {
    // Databricks account ID of any type. For non-E2 account types, get your
    // account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    // Databricks log delivery configuration ID
    LogDeliveryConfigurationId string ` path:"log_delivery_configuration_id"`
    
    Status LogDeliveryConfigStatus `json:"status"`
}


type WrappedCreateLogDeliveryConfiguration struct {
    // Databricks account ID of any type. For non-E2 account types, get your
    // account ID from the [Accounts
    // Console](https://docs.databricks.com/administration-guide/account-settings/usage.html).
    AccountId string ` path:"account_id"`
    
    LogDeliveryConfiguration any /* MISSING TYPE */ `json:"log_delivery_configuration,omitempty"`
}


type WrappedLogDeliveryConfiguration struct {
    
    LogDeliveryConfiguration any /* MISSING TYPE */ `json:"log_delivery_configuration,omitempty"`
}


type WrappedLogDeliveryConfigurations struct {
    
    LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
}

