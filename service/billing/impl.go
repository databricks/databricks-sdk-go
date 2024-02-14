// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just BillableUsage API methods
type billableUsageImpl struct {
	client *client.DatabricksClient
}

func (a *billableUsageImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/usage/download", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &downloadResponse)
	return &downloadResponse, err
}

// unexported type that holds implementations of just Budgets API methods
type budgetsImpl struct {
	client *client.DatabricksClient
}

func (a *budgetsImpl) Create(ctx context.Context, request WrappedBudget) (*WrappedBudgetWithStatus, error) {
	var wrappedBudgetWithStatus WrappedBudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &wrappedBudgetWithStatus)
	return &wrappedBudgetWithStatus, err
}

func (a *budgetsImpl) Delete(ctx context.Context, request DeleteBudgetRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, nil)
	return err
}

func (a *budgetsImpl) Get(ctx context.Context, request GetBudgetRequest) (*WrappedBudgetWithStatus, error) {
	var wrappedBudgetWithStatus WrappedBudgetWithStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &wrappedBudgetWithStatus)
	return &wrappedBudgetWithStatus, err
}

func (a *budgetsImpl) List(ctx context.Context) (*BudgetList, error) {
	var budgetList BudgetList
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, &budgetList)
	return &budgetList, err
}

func (a *budgetsImpl) Update(ctx context.Context, request WrappedBudget) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/budget/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, nil)
	return err
}

// unexported type that holds implementations of just LogDelivery API methods
type logDeliveryImpl struct {
	client *client.DatabricksClient
}

func (a *logDeliveryImpl) Create(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

func (a *logDeliveryImpl) Get(ctx context.Context, request GetLogDeliveryRequest) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

func (a *logDeliveryImpl) List(ctx context.Context, request ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

func (a *logDeliveryImpl) PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, nil)
	return err
}
