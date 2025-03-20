// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Billable Usage, Budget Policy, Budgets, Log Delivery, Usage Dashboards, etc.
package billing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// This API allows you to download billable usage logs for the specified account
// and date range. This feature works with all account types.
type BillableUsageAPI struct {
	billableUsageImpl
}

// A service serves REST API about Budget policies
type BudgetPolicyAPI struct {
	budgetPolicyImpl
}

// Delete a budget policy.
//
// Deletes a policy
func (a *BudgetPolicyAPI) DeleteByPolicyId(ctx context.Context, policyId string) (*DeleteResponse, error) {
	return a.budgetPolicyImpl.Delete(ctx, DeleteBudgetPolicyRequest{
		PolicyId: policyId,
	})
}

// Get a budget policy.
//
// Retrieves a policy by it's ID.
func (a *BudgetPolicyAPI) GetByPolicyId(ctx context.Context, policyId string) (*BudgetPolicy, error) {
	return a.budgetPolicyImpl.Get(ctx, GetBudgetPolicyRequest{
		PolicyId: policyId,
	})
}

// These APIs manage budget configurations for this account. Budgets enable you
// to monitor usage across your account. You can set up budgets to either track
// account-wide spending, or apply filters to track the spending of specific
// teams, projects, or workspaces.
type BudgetsAPI struct {
	budgetsImpl
}

// Delete budget.
//
// Deletes a budget configuration for an account. Both account and budget
// configuration are specified by ID. This cannot be undone.
func (a *BudgetsAPI) DeleteByBudgetId(ctx context.Context, budgetId string) (*DeleteBudgetConfigurationResponse, error) {
	return a.budgetsImpl.Delete(ctx, DeleteBudgetConfigurationRequest{
		BudgetId: budgetId,
	})
}

// Get budget.
//
// Gets a budget configuration for an account. Both account and budget
// configuration are specified by ID.
func (a *BudgetsAPI) GetByBudgetId(ctx context.Context, budgetId string) (*GetBudgetConfigurationResponse, error) {
	return a.budgetsImpl.Get(ctx, GetBudgetConfigurationRequest{
		BudgetId: budgetId,
	})
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
type LogDeliveryAPI struct {
	logDeliveryImpl
}

// Get log delivery configuration.
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) GetByLogDeliveryConfigurationId(ctx context.Context, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error) {
	return a.logDeliveryImpl.Get(ctx, GetLogDeliveryRequest{
		LogDeliveryConfigurationId: logDeliveryConfigurationId,
	})
}

// LogDeliveryConfigurationConfigNameToConfigIdMap calls [LogDeliveryAPI.ListAll] and creates a map of results with [LogDeliveryConfiguration].ConfigName as key and [LogDeliveryConfiguration].ConfigId as value.
//
// Returns an error if there's more than one [LogDeliveryConfiguration] with the same .ConfigName.
//
// Note: All [LogDeliveryConfiguration] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LogDeliveryAPI) LogDeliveryConfigurationConfigNameToConfigIdMap(ctx context.Context, request ListLogDeliveryRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.ConfigName
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .ConfigName: %s", key)
		}
		mapping[key] = v.ConfigId
	}
	return mapping, nil
}

// GetByConfigName calls [LogDeliveryAPI.LogDeliveryConfigurationConfigNameToConfigIdMap] and returns a single [LogDeliveryConfiguration].
//
// Returns an error if there's more than one [LogDeliveryConfiguration] with the same .ConfigName.
//
// Note: All [LogDeliveryConfiguration] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LogDeliveryAPI) GetByConfigName(ctx context.Context, name string) (*LogDeliveryConfiguration, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListLogDeliveryRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]LogDeliveryConfiguration{}
	for _, v := range result {
		key := v.ConfigName
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("LogDeliveryConfiguration named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of LogDeliveryConfiguration named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// These APIs manage usage dashboards for this account. Usage dashboards enable
// you to gain insights into your usage with pre-built dashboards: visualize
// breakdowns, analyze tag attributions, and identify cost drivers.
type UsageDashboardsAPI struct {
	usageDashboardsImpl
}
