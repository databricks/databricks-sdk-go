// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// unexported type that holds implementations of just BillableUsage API methods
type billableUsageImpl struct {
	client *client.DatabricksClient
}

func (a *billableUsageImpl) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {

	requestPb, pbErr := downloadRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var downloadResponsePb downloadResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/usage/download", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "text/plain"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&downloadResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := downloadResponseFromPb(&downloadResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just BudgetPolicy API methods
type budgetPolicyImpl struct {
	client *client.DatabricksClient
}

func (a *budgetPolicyImpl) Create(ctx context.Context, request CreateBudgetPolicyRequest) (*BudgetPolicy, error) {

	requestPb, pbErr := createBudgetPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var budgetPolicyPb budgetPolicyPb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&budgetPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := budgetPolicyFromPb(&budgetPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *budgetPolicyImpl) Delete(ctx context.Context, request DeleteBudgetPolicyRequest) error {

	requestPb, pbErr := deleteBudgetPolicyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), requestPb.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *budgetPolicyImpl) Get(ctx context.Context, request GetBudgetPolicyRequest) (*BudgetPolicy, error) {

	requestPb, pbErr := getBudgetPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var budgetPolicyPb budgetPolicyPb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), requestPb.PolicyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&budgetPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := budgetPolicyFromPb(&budgetPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Lists all policies. Policies are returned in the alphabetically ascending
// order of their names.
func (a *budgetPolicyImpl) List(ctx context.Context, request ListBudgetPoliciesRequest) listing.Iterator[BudgetPolicy] {

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

// Lists all policies. Policies are returned in the alphabetically ascending
// order of their names.
func (a *budgetPolicyImpl) ListAll(ctx context.Context, request ListBudgetPoliciesRequest) ([]BudgetPolicy, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetPolicy](ctx, iterator)
}

func (a *budgetPolicyImpl) internalList(ctx context.Context, request ListBudgetPoliciesRequest) (*ListBudgetPoliciesResponse, error) {

	requestPb, pbErr := listBudgetPoliciesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listBudgetPoliciesResponsePb listBudgetPoliciesResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listBudgetPoliciesResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listBudgetPoliciesResponseFromPb(&listBudgetPoliciesResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *budgetPolicyImpl) Update(ctx context.Context, request UpdateBudgetPolicyRequest) (*BudgetPolicy, error) {

	requestPb, pbErr := updateBudgetPolicyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var budgetPolicyPb budgetPolicyPb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budget-policies/%v", a.client.ConfiguredAccountID(), requestPb.PolicyId)
	queryParams := make(map[string]any)
	if requestPb.LimitConfig != nil {
		queryParams["limit_config"] = requestPb.LimitConfig
	}
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb).Policy,
		&budgetPolicyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := budgetPolicyFromPb(&budgetPolicyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just budgets API methods
type budgetsImpl struct {
	client *client.DatabricksClient
}

func (a *budgetsImpl) Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error) {

	requestPb, pbErr := createBudgetConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createBudgetConfigurationResponsePb createBudgetConfigurationResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createBudgetConfigurationResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createBudgetConfigurationResponseFromPb(&createBudgetConfigurationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *budgetsImpl) Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error {

	requestPb, pbErr := deleteBudgetConfigurationRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteBudgetConfigurationResponsePb deleteBudgetConfigurationResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), requestPb.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteBudgetConfigurationResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *budgetsImpl) Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error) {

	requestPb, pbErr := getBudgetConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getBudgetConfigurationResponsePb getBudgetConfigurationResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), requestPb.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getBudgetConfigurationResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getBudgetConfigurationResponseFromPb(&getBudgetConfigurationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets all budgets associated with this account.
func (a *budgetsImpl) List(ctx context.Context, request ListBudgetConfigurationsRequest) listing.Iterator[BudgetConfiguration] {

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

// Gets all budgets associated with this account.
func (a *budgetsImpl) ListAll(ctx context.Context, request ListBudgetConfigurationsRequest) ([]BudgetConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetConfiguration](ctx, iterator)
}

func (a *budgetsImpl) internalList(ctx context.Context, request ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {

	requestPb, pbErr := listBudgetConfigurationsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var listBudgetConfigurationsResponsePb listBudgetConfigurationsResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&listBudgetConfigurationsResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := listBudgetConfigurationsResponseFromPb(&listBudgetConfigurationsResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *budgetsImpl) Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error) {

	requestPb, pbErr := updateBudgetConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var updateBudgetConfigurationResponsePb updateBudgetConfigurationResponsePb
	path := fmt.Sprintf("/api/2.1/accounts/%v/budgets/%v", a.client.ConfiguredAccountID(), requestPb.BudgetId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateBudgetConfigurationResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := updateBudgetConfigurationResponseFromPb(&updateBudgetConfigurationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// unexported type that holds implementations of just LogDelivery API methods
type logDeliveryImpl struct {
	client *client.DatabricksClient
}

func (a *logDeliveryImpl) Create(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {

	requestPb, pbErr := wrappedCreateLogDeliveryConfigurationToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var wrappedLogDeliveryConfigurationPb wrappedLogDeliveryConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&wrappedLogDeliveryConfigurationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := wrappedLogDeliveryConfigurationFromPb(&wrappedLogDeliveryConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *logDeliveryImpl) Get(ctx context.Context, request GetLogDeliveryRequest) (*GetLogDeliveryConfigurationResponse, error) {

	requestPb, pbErr := getLogDeliveryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getLogDeliveryConfigurationResponsePb getLogDeliveryConfigurationResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), requestPb.LogDeliveryConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getLogDeliveryConfigurationResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getLogDeliveryConfigurationResponseFromPb(&getLogDeliveryConfigurationResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *logDeliveryImpl) List(ctx context.Context, request ListLogDeliveryRequest) listing.Iterator[LogDeliveryConfiguration] {

	getNextPage := func(ctx context.Context, req ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.internalList(ctx, req)
	}
	getItems := func(resp *WrappedLogDeliveryConfigurations) []LogDeliveryConfiguration {
		return resp.LogDeliveryConfigurations
	}
	getNextReq := func(resp *WrappedLogDeliveryConfigurations) *ListLogDeliveryRequest {
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

// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
func (a *logDeliveryImpl) ListAll(ctx context.Context, request ListLogDeliveryRequest) ([]LogDeliveryConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[LogDeliveryConfiguration](ctx, iterator)
}

func (a *logDeliveryImpl) internalList(ctx context.Context, request ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {

	requestPb, pbErr := listLogDeliveryRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var wrappedLogDeliveryConfigurationsPb wrappedLogDeliveryConfigurationsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&wrappedLogDeliveryConfigurationsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := wrappedLogDeliveryConfigurationsFromPb(&wrappedLogDeliveryConfigurationsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *logDeliveryImpl) PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {

	requestPb, pbErr := updateLogDeliveryConfigurationStatusRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var patchStatusResponsePb patchStatusResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/log-delivery/%v", a.client.ConfiguredAccountID(), requestPb.LogDeliveryConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&patchStatusResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just UsageDashboards API methods
type usageDashboardsImpl struct {
	client *client.DatabricksClient
}

func (a *usageDashboardsImpl) Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error) {

	requestPb, pbErr := createBillingUsageDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var createBillingUsageDashboardResponsePb createBillingUsageDashboardResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&createBillingUsageDashboardResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := createBillingUsageDashboardResponseFromPb(&createBillingUsageDashboardResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *usageDashboardsImpl) Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error) {

	requestPb, pbErr := getBillingUsageDashboardRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var getBillingUsageDashboardResponsePb getBillingUsageDashboardResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/dashboard", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&getBillingUsageDashboardResponsePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := getBillingUsageDashboardResponseFromPb(&getBillingUsageDashboardResponsePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}
