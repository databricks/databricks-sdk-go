// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
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

// unexported type that holds implementations of just budgets API methods
type budgetsImpl struct {
	client *client.DatabricksClient
}

func (a *budgetsImpl) Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error) {
	var createBudgetConfigurationResponse CreateBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createBudgetConfigurationResponse)
	return &createBudgetConfigurationResponse, err
}

func (a *budgetsImpl) Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error {
	var deleteBudgetConfigurationResponse DeleteBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, request, &deleteBudgetConfigurationResponse)
	return err
}

func (a *budgetsImpl) Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error) {
	var getBudgetConfigurationResponse GetBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getBudgetConfigurationResponse)
	return &getBudgetConfigurationResponse, err
}

func (a *budgetsImpl) List(ctx context.Context, request ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {
	var listBudgetConfigurationsResponse ListBudgetConfigurationsResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &listBudgetConfigurationsResponse)
	return &listBudgetConfigurationsResponse, err
}

func (a *budgetsImpl) Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error) {
	var updateBudgetConfigurationResponse UpdateBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, request, &updateBudgetConfigurationResponse)
	return &updateBudgetConfigurationResponse, err
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
	var patchStatusResponse PatchStatusResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, request, &patchStatusResponse)
	return err
}

// unexported type that holds implementations of just UsageDashboards API methods
type usageDashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *usageDashboardsImpl) Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error) {
	var createBillingUsageDashboardResponse CreateBillingUsageDashboardResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, request, &createBillingUsageDashboardResponse)
	return &createBillingUsageDashboardResponse, err
}

func (a *usageDashboardsImpl) Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error) {
	var getBillingUsageDashboardResponse GetBillingUsageDashboardResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, request, &getBillingUsageDashboardResponse)
	return &getBillingUsageDashboardResponse, err
}
