// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Billable Usage, Budget Policy, Budgets, Log Delivery, Usage Dashboards, etc.
package billing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type billableUsageBaseClient struct {
	billableUsageImpl
}

type budgetPolicyBaseClient struct {
	budgetPolicyImpl
}

// Delete a budget policy.
//
// Deletes a policy
func (a *budgetPolicyBaseClient) DeleteByPolicyId(ctx context.Context, policyId string) (*DeleteResponse, error) {
	return a.budgetPolicyImpl.Delete(ctx, DeleteBudgetPolicyRequest{
		PolicyId: policyId,
	})
}

// Get a budget policy.
//
// Retrieves a policy by it's ID.
func (a *budgetPolicyBaseClient) GetByPolicyId(ctx context.Context, policyId string) (*BudgetPolicy, error) {
	return a.budgetPolicyImpl.Get(ctx, GetBudgetPolicyRequest{
		PolicyId: policyId,
	})
}

type budgetsBaseClient struct {
	budgetsImpl
}

// Delete budget.
//
// Deletes a budget configuration for an account. Both account and budget
// configuration are specified by ID. This cannot be undone.
func (a *budgetsBaseClient) DeleteByBudgetId(ctx context.Context, budgetId string) (*DeleteBudgetConfigurationResponse, error) {
	return a.budgetsImpl.Delete(ctx, DeleteBudgetConfigurationRequest{
		BudgetId: budgetId,
	})
}

// Get budget.
//
// Gets a budget configuration for an account. Both account and budget
// configuration are specified by ID.
func (a *budgetsBaseClient) GetByBudgetId(ctx context.Context, budgetId string) (*GetBudgetConfigurationResponse, error) {
	return a.budgetsImpl.Get(ctx, GetBudgetConfigurationRequest{
		BudgetId: budgetId,
	})
}

type logDeliveryBaseClient struct {
	logDeliveryImpl
}

// Get log delivery configuration.
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *logDeliveryBaseClient) GetByLogDeliveryConfigurationId(ctx context.Context, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error) {
	return a.logDeliveryImpl.Get(ctx, GetLogDeliveryRequest{
		LogDeliveryConfigurationId: logDeliveryConfigurationId,
	})
}

// LogDeliveryConfigurationConfigNameToConfigIdMap calls [logDeliveryBaseClient.ListAll] and creates a map of results with [LogDeliveryConfiguration].ConfigName as key and [LogDeliveryConfiguration].ConfigId as value.
//
// Returns an error if there's more than one [LogDeliveryConfiguration] with the same .ConfigName.
//
// Note: All [LogDeliveryConfiguration] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *logDeliveryBaseClient) LogDeliveryConfigurationConfigNameToConfigIdMap(ctx context.Context, request ListLogDeliveryRequest) (map[string]string, error) {
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

// GetByConfigName calls [logDeliveryBaseClient.LogDeliveryConfigurationConfigNameToConfigIdMap] and returns a single [LogDeliveryConfiguration].
//
// Returns an error if there's more than one [LogDeliveryConfiguration] with the same .ConfigName.
//
// Note: All [LogDeliveryConfiguration] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *logDeliveryBaseClient) GetByConfigName(ctx context.Context, name string) (*LogDeliveryConfiguration, error) {
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

type usageDashboardsBaseClient struct {
	usageDashboardsImpl
}
