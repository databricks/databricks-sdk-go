// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
)

// This API allows you to download billable usage logs for the specified account
// and date range. This feature works with all account types.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type BillableUsageService interface {

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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type BudgetPolicyService interface {

	// Creates a new policy.
	Create(ctx context.Context, request CreateBudgetPolicyRequest) (*BudgetPolicy, error)

	// Deletes a policy
	Delete(ctx context.Context, request DeleteBudgetPolicyRequest) error

	// Retrieves a policy by it's ID.
	Get(ctx context.Context, request GetBudgetPolicyRequest) (*BudgetPolicy, error)

	// Lists all policies. Policies are returned in the alphabetically ascending
	// order of their names.
	List(ctx context.Context, request ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error)

	// Updates a policy
	Update(ctx context.Context, request UpdateBudgetPolicyRequest) (*BudgetPolicy, error)
}

// These APIs manage budget configurations for this account. Budgets enable you
// to monitor usage across your account. You can set up budgets to either track
// account-wide spending, or apply filters to track the spending of specific
// teams, projects, or workspaces.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type BudgetsService interface {

	// Create a new budget configuration for an account. For full details, see
	// https://docs.databricks.com/en/admin/account-settings/budgets.html.
	Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error)

	// Deletes a budget configuration for an account. Both account and budget
	// configuration are specified by ID. This cannot be undone.
	Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error

	// Gets a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error)

	// Gets all budgets associated with this account.
	List(ctx context.Context, request ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error)

	// Updates a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error)
}

// These APIs manage Log delivery configurations for this account. Log delivery
// configs enable you to configure the delivery of the specified type of logs to
// your storage account.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LogDeliveryService interface {

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

	// Gets a Databricks log delivery configuration object for an account, both
	// specified by ID.
	Get(ctx context.Context, request GetLogDeliveryRequest) (*GetLogDeliveryConfigurationResponse, error)

	// Gets all Databricks log delivery configurations associated with an
	// account specified by ID.
	List(ctx context.Context, request ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error)

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
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type UsageDashboardsService interface {

	// Create a usage dashboard specified by workspaceId, accountId, and
	// dashboard type.
	Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error)

	// Get a usage dashboard specified by workspaceId, accountId, and dashboard
	// type.
	Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error)
}
