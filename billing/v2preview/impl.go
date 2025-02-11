// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billingpreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/listing"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

// unexported type that holds implementations of just BillableUsagePreview API methods
type billableUsagePreviewImpl struct {
	client *client.DatabricksClient
}

func (a *billableUsagePreviewImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	var downloadResponse DownloadResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/usage/download", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &downloadResponse)
	return &downloadResponse, err
}

// unexported type that holds implementations of just BudgetPolicyPreview API methods
type budgetPolicyPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *budgetPolicyPreviewImpl) Create(ctx context.Context, request CreateBudgetPolicyRequest) (*BudgetPolicy, error) {
	var budgetPolicy BudgetPolicy
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budget-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &budgetPolicy)
	return &budgetPolicy, err
}

func (a *budgetPolicyPreviewImpl) Delete(ctx context.Context, request DeleteBudgetPolicyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *budgetPolicyPreviewImpl) Get(ctx context.Context, request GetBudgetPolicyRequest) (*BudgetPolicy, error) {
	var budgetPolicy BudgetPolicy
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &budgetPolicy)
	return &budgetPolicy, err
}

// List policies.
//
// Lists all policies. Policies are returned in the alphabetically ascending
// order of their names.
func (a *budgetPolicyPreviewImpl) List(ctx context.Context, request ListBudgetPoliciesRequest) listing.Iterator[BudgetPolicy] {

	getNextPage := func(ctx context.Context, req ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListBudgetPoliciesResponse) []BudgetPolicy {
		return resp.Policies
	}
	getNextReq := func(resp *ListBudgetPoliciesResponse) *ListBudgetPoliciesRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// List policies.
//
// Lists all policies. Policies are returned in the alphabetically ascending
// order of their names.
func (a *budgetPolicyPreviewImpl) ListAll(ctx context.Context, request ListBudgetPoliciesRequest) ([]BudgetPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetPolicy](ctx, iterator)
}
func (a *budgetPolicyPreviewImpl) internalList(ctx context.Context, request ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error) {
	var listBudgetPoliciesResponse ListBudgetPoliciesResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budget-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBudgetPoliciesResponse)
	return &listBudgetPoliciesResponse, err
}

func (a *budgetPolicyPreviewImpl) Update(ctx context.Context, request UpdateBudgetPolicyRequest) (*BudgetPolicy, error) {
	var budgetPolicy BudgetPolicy
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), request.PolicyId)
	queryParams := make(map[string]any)
	if request.LimitConfig != nil {
		queryParams["limit_config"] = request.LimitConfig
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request.Policy, &budgetPolicy)
	return &budgetPolicy, err
}

// unexported type that holds implementations of just BudgetsPreview API methods
type budgetsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *budgetsPreviewImpl) Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error) {
	var createBudgetConfigurationResponse CreateBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budgets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createBudgetConfigurationResponse)
	return &createBudgetConfigurationResponse, err
}

func (a *budgetsPreviewImpl) Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error {
	var deleteBudgetConfigurationResponse DeleteBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteBudgetConfigurationResponse)
	return err
}

func (a *budgetsPreviewImpl) Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error) {
	var getBudgetConfigurationResponse GetBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getBudgetConfigurationResponse)
	return &getBudgetConfigurationResponse, err
}

// Get all budgets.
//
// Gets all budgets associated with this account.
func (a *budgetsPreviewImpl) List(ctx context.Context, request ListBudgetConfigurationsRequest) listing.Iterator[BudgetConfiguration] {

	getNextPage := func(ctx context.Context, req ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *ListBudgetConfigurationsResponse) []BudgetConfiguration {
		return resp.Budgets
	}
	getNextReq := func(resp *ListBudgetConfigurationsResponse) *ListBudgetConfigurationsRequest {
		if resp.NextPageToken == "" {
			return nil
		}
		request.PageToken = resp.NextPageToken
		return &request
	}
	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		getNextReq)
	return iterator
}

// Get all budgets.
//
// Gets all budgets associated with this account.
func (a *budgetsPreviewImpl) ListAll(ctx context.Context, request ListBudgetConfigurationsRequest) ([]BudgetConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetConfiguration](ctx, iterator)
}
func (a *budgetsPreviewImpl) internalList(ctx context.Context, request ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {
	var listBudgetConfigurationsResponse ListBudgetConfigurationsResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budgets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &listBudgetConfigurationsResponse)
	return &listBudgetConfigurationsResponse, err
}

func (a *budgetsPreviewImpl) Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error) {
	var updateBudgetConfigurationResponse UpdateBudgetConfigurationResponse
	path := fmt.Sprintf("/api/2.1preview/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), request.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &updateBudgetConfigurationResponse)
	return &updateBudgetConfigurationResponse, err
}

// unexported type that holds implementations of just LogDeliveryPreview API methods
type logDeliveryPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *logDeliveryPreviewImpl) Create(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

func (a *logDeliveryPreviewImpl) Get(ctx context.Context, request GetLogDeliveryRequest) (*WrappedLogDeliveryConfiguration, error) {
	var wrappedLogDeliveryConfiguration WrappedLogDeliveryConfiguration
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &wrappedLogDeliveryConfiguration)
	return &wrappedLogDeliveryConfiguration, err
}

// Get all log delivery configurations.
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *logDeliveryPreviewImpl) List(ctx context.Context, request ListLogDeliveryRequest) listing.Iterator[LogDeliveryConfiguration] {

	getNextPage := func(ctx context.Context, req ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *WrappedLogDeliveryConfigurations) []LogDeliveryConfiguration {
		return resp.LogDeliveryConfigurations
	}

	iterator := listing.NewIterator(
		&request,
		getNextPage,
		getItems,
		nil)
	return iterator
}

// Get all log delivery configurations.
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *logDeliveryPreviewImpl) ListAll(ctx context.Context, request ListLogDeliveryRequest) ([]LogDeliveryConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[LogDeliveryConfiguration](ctx, iterator)
}
func (a *logDeliveryPreviewImpl) internalList(ctx context.Context, request ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
	var wrappedLogDeliveryConfigurations WrappedLogDeliveryConfigurations
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &wrappedLogDeliveryConfigurations)
	return &wrappedLogDeliveryConfigurations, err
}

func (a *logDeliveryPreviewImpl) PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	var patchStatusResponse PatchStatusResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), request.LogDeliveryConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &patchStatusResponse)
	return err
}

// unexported type that holds implementations of just UsageDashboardsPreview API methods
type usageDashboardsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *usageDashboardsPreviewImpl) Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error) {
	var createBillingUsageDashboardResponse CreateBillingUsageDashboardResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &createBillingUsageDashboardResponse)
	return &createBillingUsageDashboardResponse, err
}

func (a *usageDashboardsPreviewImpl) Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error) {
	var getBillingUsageDashboardResponse GetBillingUsageDashboardResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &getBillingUsageDashboardResponse)
	return &getBillingUsageDashboardResponse, err
}
