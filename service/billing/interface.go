// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
)

// This API allows you to download billable usage logs for the specified account
// and date range. This feature works with all account types.
type BillableUsageService interface {

	// Return billable usage logs.
	//
	// Returns billable usage logs in CSV format for the specified account and
	// date range. For the data schema, see [CSV file schema]. Note that this
	// method might take multiple minutes to complete.
	//
	// **Warning**: Depending on the queried date range, the number of
	// workspaces in the account, the size of the response and the internet
	// speed of the caller, this API may hit a timeout after a few minutes. If
	// you experience this, try to mitigate by calling the API with narrower
	// date ranges.
	//
	// [CSV file schema]: https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema
	Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error)
}

// A service serves REST API about Budget policies
type BudgetPolicyService interface {

	// Create a budget policy.
	//
	// Creates a new policy.
	Create(ctx context.Context, request CreateBudgetPolicyRequest) (*BudgetPolicy, error)

	// Delete a budget policy.
	//
	// Deletes a policy
	Delete(ctx context.Context, request DeleteBudgetPolicyRequest) error

	// Get a budget policy.
	//
	// Retrieves a policy by it's ID.
	Get(ctx context.Context, request GetBudgetPolicyRequest) (*BudgetPolicy, error)

	// List policies.
	//
	// Lists all policies. Policies are returned in the alphabetically ascending
	// order of their names.
	//
	// Use ListAll() to get all BudgetPolicy instances, which will iterate over every result page.
	List(ctx context.Context, request ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error)

	// Update a budget policy.
	//
	// Updates a policy
	Update(ctx context.Context, request UpdateBudgetPolicyRequest) (*BudgetPolicy, error)
}

// These APIs manage budget configurations for this account. Budgets enable you
// to monitor usage across your account. You can set up budgets to either track
// account-wide spending, or apply filters to track the spending of specific
// teams, projects, or workspaces.
type BudgetsService interface {

	// Create new budget.
	//
	// Create a new budget configuration for an account. For full details, see
	// https://docs.databricks.com/en/admin/account-settings/budgets.html.
	Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error)

	// Delete budget.
	//
	// Deletes a budget configuration for an account. Both account and budget
	// configuration are specified by ID. This cannot be undone.
	Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error

	// Get budget.
	//
	// Gets a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error)

	// Get all budgets.
	//
	// Gets all budgets associated with this account.
	//
	// Use ListAll() to get all BudgetConfiguration instances, which will iterate over every result page.
	List(ctx context.Context, request ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error)

	// Modify budget.
	//
	// Updates a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error)
}

// These APIs manage log delivery configurations for this account. The two
// supported log types for this API are _billable usage logs_ and _audit logs_.
// This feature is in Public Preview. This feature works with all account ID
// types.
//
// Log delivery works with all account types. However, if your account is on the
// E2 version of the platform or on a select custom plan that allows multiple
// workspaces per account, you can optionally configure different storage
// destinations for each workspace. Log delivery status is also provided to know
// the latest status of log delivery attempts. The high-level flow of billable
// usage delivery:
//
// 1. **Create storage**: In AWS, [create a new AWS S3 bucket] with a specific
// bucket policy. Using Databricks APIs, call the Account API to create a
// [storage configuration object](:method:Storage/Create) that uses the bucket
// name. 2. **Create credentials**: In AWS, create the appropriate AWS IAM role.
// For full details, including the required IAM role policies and trust
// relationship, see [Billable usage log delivery]. Using Databricks APIs, call
// the Account API to create a [credential configuration
// object](:method:Credentials/Create) that uses the IAM role"s ARN. 3. **Create
// log delivery configuration**: Using Databricks APIs, call the Account API to
// [create a log delivery configuration](:method:LogDelivery/Create) that uses
// the credential and storage configuration objects from previous steps. You can
// specify if the logs should include all events of that log type in your
// account (_Account level_ delivery) or only events for a specific set of
// workspaces (_workspace level_ delivery). Account level log delivery applies
// to all current and future workspaces plus account level logs, while workspace
// level log delivery solely delivers logs related to the specified workspaces.
// You can create multiple types of delivery configurations per account.
//
// For billable usage delivery: * For more information about billable usage
// logs, see [Billable usage log delivery]. For the CSV schema, see the [Usage
// page]. * The delivery location is
// `<bucket-name>/<prefix>/billable-usage/csv/`, where `<prefix>` is the name of
// the optional delivery path prefix you set up during log delivery
// configuration. Files are named
// `workspaceId=<workspace-id>-usageMonth=<month>.csv`. * All billable usage
// logs apply to specific workspaces (_workspace level_ logs). You can aggregate
// usage for your entire account by creating an _account level_ delivery
// configuration that delivers logs for all current and future workspaces in
// your account. * The files are delivered daily by overwriting the month's CSV
// file for each workspace.
//
// For audit log delivery: * For more information about about audit log
// delivery, see [Audit log delivery], which includes information about the used
// JSON schema. * The delivery location is
// `<bucket-name>/<delivery-path-prefix>/workspaceId=<workspaceId>/date=<yyyy-mm-dd>/auditlogs_<internal-id>.json`.
// Files may get overwritten with the same content multiple times to achieve
// exactly-once delivery. * If the audit log delivery configuration included
// specific workspace IDs, only _workspace-level_ audit logs for those
// workspaces are delivered. If the log delivery configuration applies to the
// entire account (_account level_ delivery configuration), the audit log
// delivery includes workspace-level audit logs for all workspaces in the
// account as well as account-level audit logs. See [Audit log delivery] for
// details. * Auditable events are typically available in logs within 15
// minutes.
//
// [Audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [Billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
// [Usage page]: https://docs.databricks.com/administration-guide/account-settings/usage.html
// [create a new AWS S3 bucket]: https://docs.databricks.com/administration-guide/account-api/aws-storage.html
type LogDeliveryService interface {

	// Create a new log delivery configuration.
	//
	// Creates a new Databricks log delivery configuration to enable delivery of
	// the specified type of logs to your storage location. This requires that
	// you already created a [credential object](:method:Credentials/Create)
	// (which encapsulates a cross-account service IAM role) and a [storage
	// configuration object](:method:Storage/Create) (which encapsulates an S3
	// bucket).
	//
	// For full details, including the required IAM role policies and bucket
	// policies, see [Deliver and access billable usage logs] or [Configure
	// audit logging].
	//
	// **Note**: There is a limit on the number of log delivery configurations
	// available per account (each limit applies separately to each log type
	// including billable usage and audit logs). You can create a maximum of two
	// enabled account-level delivery configurations (configurations without a
	// workspace filter) per type. Additionally, you can create two enabled
	// workspace-level delivery configurations per workspace for each log type,
	// which means that the same workspace ID can occur in the workspace filter
	// for no more than two delivery configurations per log type.
	//
	// You cannot delete a log delivery configuration, but you can disable it
	// (see [Enable or disable log delivery
	// configuration](:method:LogDelivery/PatchStatus)).
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [Deliver and access billable usage logs]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	Create(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error)

	// Get log delivery configuration.
	//
	// Gets a Databricks log delivery configuration object for an account, both
	// specified by ID.
	Get(ctx context.Context, request GetLogDeliveryRequest) (*WrappedLogDeliveryConfiguration, error)

	// Get all log delivery configurations.
	//
	// Gets all Databricks log delivery configurations associated with an
	// account specified by ID.
	//
	// Use ListAll() to get all LogDeliveryConfiguration instances
	List(ctx context.Context, request ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error)

	// Enable or disable log delivery configuration.
	//
	// Enables or disables a log delivery configuration. Deletion of delivery
	// configurations is not supported, so disable log delivery configurations
	// that are no longer needed. Note that you can't re-enable a delivery
	// configuration if this would violate the delivery configuration limits
	// described under [Create log delivery](:method:LogDelivery/Create).
	PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error
}

// These APIs manage usage dashboards for this account. Usage dashboards enable
// you to gain insights into your usage with pre-built dashboards: visualize
// breakdowns, analyze tag attributions, and identify cost drivers.
type UsageDashboardsService interface {

	// Create new usage dashboard.
	//
	// Create a usage dashboard specified by workspaceId, accountId, and
	// dashboard type.
	Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error)

	// Get usage dashboard.
	//
	// Get a usage dashboard specified by workspaceId, accountId, and dashboard
	// type.
	Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error)
}
