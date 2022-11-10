// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewBillableUsageDownload(client *client.DatabricksClient) BillableUsageDownloadService {
	return &BillableUsageDownloadAPI{
		client: client,
	}
}

type BillableUsageDownloadAPI struct {
	client *client.DatabricksClient
}

// Return billable usage logs
//
// Returns billable usage logs in CSV format for the specified account and date
// range. For the data schema, see [CSV file
// schema](https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema).
// Note that this method might take multiple seconds to complete.
func (a *BillableUsageDownloadAPI) DownloadBillableUsage(ctx context.Context, request DownloadBillableUsageRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/usage/download", request.AccountId)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

// Return billable usage logs
//
// Returns billable usage logs in CSV format for the specified account and date
// range. For the data schema, see [CSV file
// schema](https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema).
// Note that this method might take multiple seconds to complete.
func (a *BillableUsageDownloadAPI) DownloadBillableUsageByAccountId(ctx context.Context, accountId string) error {
	return a.DownloadBillableUsage(ctx, DownloadBillableUsageRequest{
		AccountId: accountId,
	})
}

func NewBudgets(client *client.DatabricksClient) BudgetsService {
	return &BudgetsAPI{
		client: client,
	}
}

type BudgetsAPI struct {
	client *client.DatabricksClient
}

// Create a new budget
//
// Creates a new budget in the specified account.
func (a *BudgetsAPI) CreateBudget(ctx context.Context, request CreateBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", request.AccountId)
	err := a.client.Post(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

// Delete budget
//
// Deletes the budget specified by its UUID.
func (a *BudgetsAPI) DeleteBudget(ctx context.Context, request DeleteBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete budget
//
// Deletes the budget specified by its UUID.
func (a *BudgetsAPI) DeleteBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) error {
	return a.DeleteBudget(ctx, DeleteBudgetRequest{
		AccountId: accountId,
		BudgetId:  budgetId,
	})
}

// Get all budgets
//
// Gets all budgets associated with this account, including noncumulative status
// for each day that the budget is configured to include.
func (a *BudgetsAPI) GetAllBudgets(ctx context.Context, request GetAllBudgetsRequest) (*BudgetList, error) {
	var budgetList BudgetList
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", request.AccountId)
	err := a.client.Get(ctx, path, request, &budgetList)
	return &budgetList, err
}

// Get all budgets
//
// Gets all budgets associated with this account, including noncumulative status
// for each day that the budget is configured to include.
func (a *BudgetsAPI) GetAllBudgetsByAccountId(ctx context.Context, accountId string) (*BudgetList, error) {
	return a.GetAllBudgets(ctx, GetAllBudgetsRequest{
		AccountId: accountId,
	})
}

// Get budget and its status
//
// Gets the budget specified by its UUID, including noncumulative status for
// each day that the budget is configured to include.
func (a *BudgetsAPI) GetBudget(ctx context.Context, request GetBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Get(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

// Get budget and its status
//
// Gets the budget specified by its UUID, including noncumulative status for
// each day that the budget is configured to include.
func (a *BudgetsAPI) GetBudgetByAccountIdAndBudgetId(ctx context.Context, accountId string, budgetId string) (*BudgetWithStatus, error) {
	return a.GetBudget(ctx, GetBudgetRequest{
		AccountId: accountId,
		BudgetId:  budgetId,
	})
}

// Modify budget
//
// Modifies a budget in this account. Budget properties are completely
// overwritten.
func (a *BudgetsAPI) ModifyBudget(ctx context.Context, request ModifyBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", request.AccountId, request.BudgetId)
	err := a.client.Patch(ctx, path, request)
	return err
}

func NewLogDelivery(client *client.DatabricksClient) LogDeliveryService {
	return &LogDeliveryAPI{
		client: client,
	}
}

type LogDeliveryAPI struct {
	client *client.DatabricksClient
}

// Create a new log delivery configuration
//
// Creates a new Databricks log delivery configuration to enable delivery of the
// specified type of logs to your storage location. This requires that you
// already created a [credential object](#operation/create-credential-config)
// (which encapsulates a cross-account service IAM role) and a [storage
// configuration object](#operation/create-storage-config) (which encapsulates
// an S3 bucket).
//
// For full details, including the required IAM role policies and bucket
// policies, see [Deliver and access billable usage
// logs](https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html)
// or [Configure audit
// logging](https://docs.databricks.com/administration-guide/account-settings/audit-logs.html).
//
// **Note**: There is a limit on the number of log delivery configurations
// available per account (each limit applies separately to each log type
// including billable usage and audit logs). You can create a maximum of two
// enabled account-level delivery configurations (configurations without a
// workspace filter) per type. Additionally, you can create two enabled
// workspace-level delivery configurations per workspace for each log type,
// which means that the same workspace ID can occur in the workspace filter for
// no more than two delivery configurations per log type.
//
// You cannot delete a log delivery configuration, but you can disable it (see
// [Enable or disable log delivery
// configuration](#operation/patch-log-delivery-config-status)).
func (a *LogDeliveryAPI) CreateLogDeliveryConfig(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", request.AccountId)
	err := a.client.Post(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Get all log delivery configurations
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
//
// Use GetAllLogDeliveryConfigsAll() to get all LogDeliveryConfiguration instances
func (a *LogDeliveryAPI) GetAllLogDeliveryConfigs(ctx context.Context, request GetAllLogDeliveryConfigsRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", request.AccountId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

// GetAllLogDeliveryConfigsAll returns all LogDeliveryConfiguration instances. This method exists for consistency purposes.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LogDeliveryAPI) GetAllLogDeliveryConfigsAll(ctx context.Context, request GetAllLogDeliveryConfigsRequest) ([]LogDeliveryConfiguration, error) {
	response, err := a.GetAllLogDeliveryConfigs(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.LogDeliveryConfigurations, nil
}

// Get all log delivery configurations
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *LogDeliveryAPI) GetAllLogDeliveryConfigsByAccountId(ctx context.Context, accountId string) (*WrappedLogDeliveryConfigurations, error) {
	return a.GetAllLogDeliveryConfigs(ctx, GetAllLogDeliveryConfigsRequest{
		AccountId: accountId,
	})
}

// Get log delivery configuration
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) GetLogDeliveryConfig(ctx context.Context, request GetLogDeliveryConfigRequest) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", request.AccountId, request.LogDeliveryConfigurationId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Get log delivery configuration
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) GetLogDeliveryConfigByAccountIdAndLogDeliveryConfigurationId(ctx context.Context, accountId string, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error) {
	return a.GetLogDeliveryConfig(ctx, GetLogDeliveryConfigRequest{
		AccountId:                  accountId,
		LogDeliveryConfigurationId: logDeliveryConfigurationId,
	})
}

// Enable or disable log delivery configuration
//
// Enables or disables a log delivery configuration. Deletion of delivery
// configurations is not supported, so disable log delivery configurations that
// are no longer needed. Note that you can't re-enable a delivery configuration
// if this would violate the delivery configuration limits described under
// [Create log delivery](#operation/create-log-delivery-config).
func (a *LogDeliveryAPI) PatchLogDeliveryConfigStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", request.AccountId, request.LogDeliveryConfigurationId)
	err := a.client.Patch(ctx, path, request)
	return err
}
