// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just BillableUsageDownload API methods
type billableUsageDownloadImpl struct {
	client *client.DatabricksClient
}

func (a *billableUsageDownloadImpl) DownloadBillableUsage(ctx context.Context, request DownloadBillableUsageRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/usage/download", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, request, nil)
	return err
}

// unexported type that holds implementations of just Budgets API methods
type budgetsImpl struct {
	client *client.DatabricksClient
}

func (a *budgetsImpl) CreateBudget(ctx context.Context, request CreateBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

func (a *budgetsImpl) DeleteBudget(ctx context.Context, request DeleteBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", a.client.Config.AccountID, request.BudgetId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *budgetsImpl) GetBudget(ctx context.Context, request GetBudgetRequest) (*BudgetWithStatus, error) {
	var budgetWithStatus BudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", a.client.Config.AccountID, request.BudgetId)
	err := a.client.Get(ctx, path, request, &budgetWithStatus)
	return &budgetWithStatus, err
}

func (a *budgetsImpl) ListBudgets(ctx context.Context) (*BudgetList, error) {
	var budgetList BudgetList
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &budgetList)
	return &budgetList, err
}

func (a *budgetsImpl) UpdateBudget(ctx context.Context, request UpdateBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", a.client.Config.AccountID, request.BudgetId)
	err := a.client.Patch(ctx, path, request)
	return err
}

// unexported type that holds implementations of just LogDelivery API methods
type logDeliveryImpl struct {
	client *client.DatabricksClient
}

func (a *logDeliveryImpl) CreateLogDeliveryConfig(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

func (a *logDeliveryImpl) GetLogDeliveryConfig(ctx context.Context, request GetLogDeliveryConfigRequest) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.Config.AccountID, request.LogDeliveryConfigurationId)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

func (a *logDeliveryImpl) ListLogDeliveryConfigs(ctx context.Context, request ListLogDeliveryConfigsRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

func (a *logDeliveryImpl) PatchLogDeliveryConfigStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.Config.AccountID, request.LogDeliveryConfigurationId)
	err := a.client.Patch(ctx, path, request)
	return err
}
