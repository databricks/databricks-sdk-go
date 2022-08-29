// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package logdelivery

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// These APIs manage log delivery configurations for this account. The two
// supported log types for this API are _billable usage logs_ and _audit logs_.
// This feature is in Public Preview. This feature works with all account ID
// types. Log delivery works with all account types. However, if your account is
// on the E2 version of the platform or on a select custom plan that allows
// multiple workspaces per account, you can optionally configure different
// storage destinations for each workspace. Log delivery status is also provided
// to know the latest status of log delivery attempts. The high-level flow of
// billable usage delivery: 1. **Create storage**: In AWS, [create a new AWS S3
// bucket](https://docs.databricks.com/administration-guide/account-api/aws-storage.html)
// with a specific bucket policy. Using Databricks APIs, call the Account API to
// create a [storage configuration object](#operation/create-storage-config)
// that uses the bucket name. 2. **Create credentials**: In AWS, create the
// appropriate AWS IAM role. For full details, including the required IAM role
// policies and trust relationship, see [Billable usage log
// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
// Using Databricks APIs, call the Account API to create a [credential
// configuration object](#operation/create-credential-config) that uses the IAM
// role&#39;s ARN. 3. **Create log delivery configuration**: Using Databricks APIs,
// call the Account API to [create a log delivery
// configuration](#operation/create-log-delivery-config) that uses the
// credential and storage configuration objects from previous steps. You can
// specify if the logs should include all events of that log type in your
// account (_Account level_ delivery) or only events for a specific set of
// workspaces (_workspace level_ delivery). Account level log delivery applies
// to all current and future workspaces plus account level logs, while workspace
// level log delivery solely delivers logs related to the specified workspaces.
// You can create multiple types of delivery configurations per account. For
// billable usage delivery: * For more information about billable usage logs,
// see [Billable usage log
// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html).
// For the CSV schema, see the [Usage
// page](https://docs.databricks.com/administration-guide/account-settings/usage.html).
// * The delivery location is `&lt;bucket-name&gt;/&lt;prefix&gt;/billable-usage/csv/`,
// where `&lt;prefix&gt;` is the name of the optional delivery path prefix you set up
// during log delivery configuration. Files are named
// `workspaceId=&lt;workspace-id&gt;-usageMonth=&lt;month&gt;.csv`. * All billable usage
// logs apply to specific workspaces (_workspace level_ logs). You can aggregate
// usage for your entire account by creating an _account level_ delivery
// configuration that delivers logs for all current and future workspaces in
// your account. * The files are delivered daily by overwriting the month&#39;s CSV
// file for each workspace. For audit log delivery: * For more information about
// about audit log delivery, see [Audit log
// delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html),
// which includes information about the used JSON schema. * The delivery
// location is
// `&lt;bucket-name&gt;/&lt;delivery-path-prefix&gt;/workspaceId=&lt;workspaceId&gt;/date=&lt;yyyy-mm-dd&gt;/auditlogs_&lt;internal-id&gt;.json`.
// Files may get overwritten with the same content multiple times to achieve
// exactly-once delivery. * If the audit log delivery configuration included
// specific workspace IDs, only _workspace-level_ audit logs for those
// workspaces are delivered. If the log delivery configuration applies to the
// entire account (_account level_ delivery configuration), the audit log
// delivery includes workspace-level audit logs for all workspaces in the
// account as well as account-level audit logs. See [Audit log
// delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html)
// for details. * Auditable events are typically available in logs within 15
// minutes.
type LogdeliveryService interface {
    // Create a new Databricks log delivery configuration to enable delivery of
    // the specified type of logs to your storage location. This requires that
    // you already created a [credential
    // object](#operation/create-credential-config) (which encapsulates a
    // cross-account service IAM role) and a [storage configuration
    // object](#operation/create-storage-config) (which encapsulates an S3
    // bucket). For full details, including the required IAM role policies and
    // bucket policies, see [Billable usage log
    // delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html)
    // or [Audit log
    // delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
    // Note: There is a limit on the number of log delivery configurations
    // available per account (each limit applies separately to each log type
    // including billable usage and audit logs). You can create a maximum of two
    // enabled account-level delivery configurations (configurations without a
    // workspace filter) per type. Additionally, you can create two enabled
    // workspace level delivery configurations per workspace for each log type,
    // meaning the same workspace ID can occur in the workspace filter for no
    // more than two delivery configurations per log type. You cannot delete a
    // log delivery configuration, but you can disable it (see [Enable or
    // disable log delivery
    // configuration](#operation/patch-log-delivery-config-status)).
    CreateLogDeliveryConfig(ctx context.Context, wrappedCreateLogDeliveryConfiguration WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error)
    // Get all Databricks log delivery configurations associated with an account
    // specified by ID.
    GetAllLogDeliveryConfigs(ctx context.Context, getAllLogDeliveryConfigsRequest GetAllLogDeliveryConfigsRequest) (*WrappedLogDeliveryConfigurations, error)
    // Get a Databricks log delivery configuration object for an account, both
    // specified by ID.
    GetLogDeliveryConfig(ctx context.Context, getLogDeliveryConfigRequest GetLogDeliveryConfigRequest) (*WrappedLogDeliveryConfiguration, error)
    // Enable or disable a log delivery configuration. Deletion of delivery
    // configurations is not supported, so disable log delivery configurations
    // that are no longer needed. Note that you can&#39;t re-enable a delivery
    // configuration if this would violate the delivery configuration limits
    // described under [Create log
    // delivery](#operation/create-log-delivery-config).
    PatchLogDeliveryConfigStatus(ctx context.Context, updateLogDeliveryConfigurationStatusRequest UpdateLogDeliveryConfigurationStatusRequest) error
	GetAllLogDeliveryConfigsByAccountId(ctx context.Context, accountId string) (*WrappedLogDeliveryConfigurations, error)
	GetLogDeliveryConfigByAccountIdAndLogDeliveryConfigurationId(ctx context.Context, accountId string, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error)
}

func New(client *client.DatabricksClient) LogdeliveryService {
	return &LogdeliveryAPI{
		client: client,
	}
}

type LogdeliveryAPI struct {
	client *client.DatabricksClient
}

// Create a new Databricks log delivery configuration to enable delivery of the
// specified type of logs to your storage location. This requires that you
// already created a [credential object](#operation/create-credential-config)
// (which encapsulates a cross-account service IAM role) and a [storage
// configuration object](#operation/create-storage-config) (which encapsulates
// an S3 bucket). For full details, including the required IAM role policies and
// bucket policies, see [Billable usage log
// delivery](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html)
// or [Audit log
// delivery](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
// Note: There is a limit on the number of log delivery configurations available
// per account (each limit applies separately to each log type including
// billable usage and audit logs). You can create a maximum of two enabled
// account-level delivery configurations (configurations without a workspace
// filter) per type. Additionally, you can create two enabled workspace level
// delivery configurations per workspace for each log type, meaning the same
// workspace ID can occur in the workspace filter for no more than two delivery
// configurations per log type. You cannot delete a log delivery configuration,
// but you can disable it (see [Enable or disable log delivery
// configuration](#operation/patch-log-delivery-config-status)).
func (a *LogdeliveryAPI) CreateLogDeliveryConfig(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/accounts/%v/log-delivery", request.AccountId)
	err := a.client.Post(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Get all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *LogdeliveryAPI) GetAllLogDeliveryConfigs(ctx context.Context, request GetAllLogDeliveryConfigsRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/accounts/%v/log-delivery", request.AccountId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

// Get a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogdeliveryAPI) GetLogDeliveryConfig(ctx context.Context, request GetLogDeliveryConfigRequest) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/accounts/%v/log-delivery/%v", request.AccountId, request.LogDeliveryConfigurationId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Enable or disable a log delivery configuration. Deletion of delivery
// configurations is not supported, so disable log delivery configurations that
// are no longer needed. Note that you can&#39;t re-enable a delivery configuration
// if this would violate the delivery configuration limits described under
// [Create log delivery](#operation/create-log-delivery-config).
func (a *LogdeliveryAPI) PatchLogDeliveryConfigStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	path := fmt.Sprintf("/accounts/%v/log-delivery/%v", request.AccountId, request.LogDeliveryConfigurationId)
	err := a.client.Patch(ctx, path, request)
	return err
}


func (a *LogdeliveryAPI) GetAllLogDeliveryConfigsByAccountId(ctx context.Context, accountId string) (*WrappedLogDeliveryConfigurations, error) {
	return a.GetAllLogDeliveryConfigs(ctx, GetAllLogDeliveryConfigsRequest{
		AccountId: accountId,
	})
}

func (a *LogdeliveryAPI) GetLogDeliveryConfigByAccountIdAndLogDeliveryConfigurationId(ctx context.Context, accountId string, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error) {
	return a.GetLogDeliveryConfig(ctx, GetLogDeliveryConfigRequest{
		AccountId: accountId,
		LogDeliveryConfigurationId: logDeliveryConfigurationId,
	})
}
