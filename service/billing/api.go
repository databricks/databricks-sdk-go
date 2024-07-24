// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Billable Usage, Budgets, Log Delivery, Usage Dashboards, etc.
package billing

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/useragent"
)

type BillableUsageInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockBillableUsageInterface instead.
	WithImpl(impl BillableUsageService) BillableUsageInterface

	// Impl returns low-level BillableUsage API implementation
	// Deprecated: use MockBillableUsageInterface instead.
	Impl() BillableUsageService

	// Return billable usage logs.
	//
	// Returns billable usage logs in CSV format for the specified account and date
	// range. For the data schema, see [CSV file schema]. Note that this method
	// might take multiple minutes to complete.
	//
	// **Warning**: Depending on the queried date range, the number of workspaces in
	// the account, the size of the response and the internet speed of the caller,
	// this API may hit a timeout after a few minutes. If you experience this, try
	// to mitigate by calling the API with narrower date ranges.
	//
	// [CSV file schema]: https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema
	Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error)
}

func NewBillableUsage(client *client.DatabricksClient) *BillableUsageAPI {
	return &BillableUsageAPI{
		impl: &billableUsageImpl{
			client: client,
		},
	}
}

// This API allows you to download billable usage logs for the specified account
// and date range. This feature works with all account types.
type BillableUsageAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(BillableUsageService)
	impl BillableUsageService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockBillableUsageInterface instead.
func (a *BillableUsageAPI) WithImpl(impl BillableUsageService) BillableUsageInterface {
	a.impl = impl
	return a
}

// Impl returns low-level BillableUsage API implementation
// Deprecated: use MockBillableUsageInterface instead.
func (a *BillableUsageAPI) Impl() BillableUsageService {
	return a.impl
}

// Return billable usage logs.
//
// Returns billable usage logs in CSV format for the specified account and date
// range. For the data schema, see [CSV file schema]. Note that this method
// might take multiple minutes to complete.
//
// **Warning**: Depending on the queried date range, the number of workspaces in
// the account, the size of the response and the internet speed of the caller,
// this API may hit a timeout after a few minutes. If you experience this, try
// to mitigate by calling the API with narrower date ranges.
//
// [CSV file schema]: https://docs.databricks.com/administration-guide/account-settings/usage-analysis.html#schema
func (a *BillableUsageAPI) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	return a.impl.Download(ctx, request)
}

type BudgetsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockBudgetsInterface instead.
	WithImpl(impl BudgetsService) BudgetsInterface

	// Impl returns low-level Budgets API implementation
	// Deprecated: use MockBudgetsInterface instead.
	Impl() BudgetsService

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

	// Delete budget.
	//
	// Deletes a budget configuration for an account. Both account and budget
	// configuration are specified by ID. This cannot be undone.
	DeleteByBudgetId(ctx context.Context, budgetId string) error

	// Get budget.
	//
	// Gets a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error)

	// Get budget.
	//
	// Gets a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	GetByBudgetId(ctx context.Context, budgetId string) (*GetBudgetConfigurationResponse, error)

	// Get all budgets.
	//
	// Gets all budgets associated with this account.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListBudgetConfigurationsRequest) listing.Iterator[BudgetConfiguration]

	// Get all budgets.
	//
	// Gets all budgets associated with this account.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListBudgetConfigurationsRequest) ([]BudgetConfiguration, error)

	// Modify budget.
	//
	// Updates a budget configuration for an account. Both account and budget
	// configuration are specified by ID.
	Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error)
}

func NewBudgets(client *client.DatabricksClient) *BudgetsAPI {
	return &BudgetsAPI{
		impl: &budgetsImpl{
			client: client,
		},
	}
}

// These APIs manage budget configurations for this account. Budgets enable you
// to monitor usage across your account. You can set up budgets to either track
// account-wide spending, or apply filters to track the spending of specific
// teams, projects, or workspaces.
type BudgetsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(BudgetsService)
	impl BudgetsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockBudgetsInterface instead.
func (a *BudgetsAPI) WithImpl(impl BudgetsService) BudgetsInterface {
	a.impl = impl
	return a
}

// Impl returns low-level Budgets API implementation
// Deprecated: use MockBudgetsInterface instead.
func (a *BudgetsAPI) Impl() BudgetsService {
	return a.impl
}

// Create new budget.
//
// Create a new budget configuration for an account. For full details, see
// https://docs.databricks.com/en/admin/account-settings/budgets.html.
func (a *BudgetsAPI) Create(ctx context.Context, request CreateBudgetConfigurationRequest) (*CreateBudgetConfigurationResponse, error) {
	return a.impl.Create(ctx, request)
}

// Delete budget.
//
// Deletes a budget configuration for an account. Both account and budget
// configuration are specified by ID. This cannot be undone.
func (a *BudgetsAPI) Delete(ctx context.Context, request DeleteBudgetConfigurationRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete budget.
//
// Deletes a budget configuration for an account. Both account and budget
// configuration are specified by ID. This cannot be undone.
func (a *BudgetsAPI) DeleteByBudgetId(ctx context.Context, budgetId string) error {
	return a.impl.Delete(ctx, DeleteBudgetConfigurationRequest{
		BudgetId: budgetId,
	})
}

// Get budget.
//
// Gets a budget configuration for an account. Both account and budget
// configuration are specified by ID.
func (a *BudgetsAPI) Get(ctx context.Context, request GetBudgetConfigurationRequest) (*GetBudgetConfigurationResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get budget.
//
// Gets a budget configuration for an account. Both account and budget
// configuration are specified by ID.
func (a *BudgetsAPI) GetByBudgetId(ctx context.Context, budgetId string) (*GetBudgetConfigurationResponse, error) {
	return a.impl.Get(ctx, GetBudgetConfigurationRequest{
		BudgetId: budgetId,
	})
}

// Get all budgets.
//
// Gets all budgets associated with this account.
//
// This method is generated by Databricks SDK Code Generator.
func (a *BudgetsAPI) List(ctx context.Context, request ListBudgetConfigurationsRequest) listing.Iterator[BudgetConfiguration] {

	getNextPage := func(ctx context.Context, req ListBudgetConfigurationsRequest) (*ListBudgetConfigurationsResponse, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *BudgetsAPI) ListAll(ctx context.Context, request ListBudgetConfigurationsRequest) ([]BudgetConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[BudgetConfiguration](ctx, iterator)
}

// Modify budget.
//
// Updates a budget configuration for an account. Both account and budget
// configuration are specified by ID.
func (a *BudgetsAPI) Update(ctx context.Context, request UpdateBudgetConfigurationRequest) (*UpdateBudgetConfigurationResponse, error) {
	return a.impl.Update(ctx, request)
}

type LogDeliveryInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockLogDeliveryInterface instead.
	WithImpl(impl LogDeliveryService) LogDeliveryInterface

	// Impl returns low-level LogDelivery API implementation
	// Deprecated: use MockLogDeliveryInterface instead.
	Impl() LogDeliveryService

	// Create a new log delivery configuration.
	//
	// Creates a new Databricks log delivery configuration to enable delivery of the
	// specified type of logs to your storage location. This requires that you
	// already created a [credential object](:method:Credentials/Create) (which
	// encapsulates a cross-account service IAM role) and a [storage configuration
	// object](:method:Storage/Create) (which encapsulates an S3 bucket).
	//
	// For full details, including the required IAM role policies and bucket
	// policies, see [Deliver and access billable usage logs] or [Configure audit
	// logging].
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

	// Get log delivery configuration.
	//
	// Gets a Databricks log delivery configuration object for an account, both
	// specified by ID.
	GetByLogDeliveryConfigurationId(ctx context.Context, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error)

	// Get all log delivery configurations.
	//
	// Gets all Databricks log delivery configurations associated with an account
	// specified by ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	List(ctx context.Context, request ListLogDeliveryRequest) listing.Iterator[LogDeliveryConfiguration]

	// Get all log delivery configurations.
	//
	// Gets all Databricks log delivery configurations associated with an account
	// specified by ID.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListAll(ctx context.Context, request ListLogDeliveryRequest) ([]LogDeliveryConfiguration, error)

	// LogDeliveryConfigurationConfigNameToConfigIdMap calls [LogDeliveryAPI.ListAll] and creates a map of results with [LogDeliveryConfiguration].ConfigName as key and [LogDeliveryConfiguration].ConfigId as value.
	//
	// Returns an error if there's more than one [LogDeliveryConfiguration] with the same .ConfigName.
	//
	// Note: All [LogDeliveryConfiguration] instances are loaded into memory before creating a map.
	//
	// This method is generated by Databricks SDK Code Generator.
	LogDeliveryConfigurationConfigNameToConfigIdMap(ctx context.Context, request ListLogDeliveryRequest) (map[string]string, error)

	// GetByConfigName calls [LogDeliveryAPI.LogDeliveryConfigurationConfigNameToConfigIdMap] and returns a single [LogDeliveryConfiguration].
	//
	// Returns an error if there's more than one [LogDeliveryConfiguration] with the same .ConfigName.
	//
	// Note: All [LogDeliveryConfiguration] instances are loaded into memory before returning matching by name.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetByConfigName(ctx context.Context, name string) (*LogDeliveryConfiguration, error)

	// Enable or disable log delivery configuration.
	//
	// Enables or disables a log delivery configuration. Deletion of delivery
	// configurations is not supported, so disable log delivery configurations that
	// are no longer needed. Note that you can't re-enable a delivery configuration
	// if this would violate the delivery configuration limits described under
	// [Create log delivery](:method:LogDelivery/Create).
	PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error
}

func NewLogDelivery(client *client.DatabricksClient) *LogDeliveryAPI {
	return &LogDeliveryAPI{
		impl: &logDeliveryImpl{
			client: client,
		},
	}
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
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(LogDeliveryService)
	impl LogDeliveryService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockLogDeliveryInterface instead.
func (a *LogDeliveryAPI) WithImpl(impl LogDeliveryService) LogDeliveryInterface {
	a.impl = impl
	return a
}

// Impl returns low-level LogDelivery API implementation
// Deprecated: use MockLogDeliveryInterface instead.
func (a *LogDeliveryAPI) Impl() LogDeliveryService {
	return a.impl
}

// Create a new log delivery configuration.
//
// Creates a new Databricks log delivery configuration to enable delivery of the
// specified type of logs to your storage location. This requires that you
// already created a [credential object](:method:Credentials/Create) (which
// encapsulates a cross-account service IAM role) and a [storage configuration
// object](:method:Storage/Create) (which encapsulates an S3 bucket).
//
// For full details, including the required IAM role policies and bucket
// policies, see [Deliver and access billable usage logs] or [Configure audit
// logging].
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
// configuration](:method:LogDelivery/PatchStatus)).
//
// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [Deliver and access billable usage logs]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
func (a *LogDeliveryAPI) Create(ctx context.Context, request WrappedCreateLogDeliveryConfiguration) (*WrappedLogDeliveryConfiguration, error) {
	return a.impl.Create(ctx, request)
}

// Get log delivery configuration.
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) Get(ctx context.Context, request GetLogDeliveryRequest) (*WrappedLogDeliveryConfiguration, error) {
	return a.impl.Get(ctx, request)
}

// Get log delivery configuration.
//
// Gets a Databricks log delivery configuration object for an account, both
// specified by ID.
func (a *LogDeliveryAPI) GetByLogDeliveryConfigurationId(ctx context.Context, logDeliveryConfigurationId string) (*WrappedLogDeliveryConfiguration, error) {
	return a.impl.Get(ctx, GetLogDeliveryRequest{
		LogDeliveryConfigurationId: logDeliveryConfigurationId,
	})
}

// Get all log delivery configurations.
//
// Gets all Databricks log delivery configurations associated with an account
// specified by ID.
//
// This method is generated by Databricks SDK Code Generator.
func (a *LogDeliveryAPI) List(ctx context.Context, request ListLogDeliveryRequest) listing.Iterator[LogDeliveryConfiguration] {

	getNextPage := func(ctx context.Context, req ListLogDeliveryRequest) (*WrappedLogDeliveryConfigurations, error) {
		ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
		return a.impl.List(ctx, req)
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
//
// This method is generated by Databricks SDK Code Generator.
func (a *LogDeliveryAPI) ListAll(ctx context.Context, request ListLogDeliveryRequest) ([]LogDeliveryConfiguration, error) {
	iterator := a.List(ctx, request)
	return listing.ToSlice[LogDeliveryConfiguration](ctx, iterator)
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

// Enable or disable log delivery configuration.
//
// Enables or disables a log delivery configuration. Deletion of delivery
// configurations is not supported, so disable log delivery configurations that
// are no longer needed. Note that you can't re-enable a delivery configuration
// if this would violate the delivery configuration limits described under
// [Create log delivery](:method:LogDelivery/Create).
func (a *LogDeliveryAPI) PatchStatus(ctx context.Context, request UpdateLogDeliveryConfigurationStatusRequest) error {
	return a.impl.PatchStatus(ctx, request)
}

type UsageDashboardsInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockUsageDashboardsInterface instead.
	WithImpl(impl UsageDashboardsService) UsageDashboardsInterface

	// Impl returns low-level UsageDashboards API implementation
	// Deprecated: use MockUsageDashboardsInterface instead.
	Impl() UsageDashboardsService

	// Create new usage dashboard.
	//
	// Create a usage dashboard specified by workspaceId, accountId, and dashboard
	// type.
	Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error)

	// Get usage dashboard.
	//
	// Get a usage dashboard specified by workspaceId, accountId, and dashboard
	// type.
	Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error)
}

func NewUsageDashboards(client *client.DatabricksClient) *UsageDashboardsAPI {
	return &UsageDashboardsAPI{
		impl: &usageDashboardsImpl{
			client: client,
		},
	}
}

// These APIs manage usage dashboards for this account. Usage dashboards enable
// you to gain insights into your usage with pre-built dashboards: visualize
// breakdowns, analyze tag attributions, and identify cost drivers.
type UsageDashboardsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(UsageDashboardsService)
	impl UsageDashboardsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockUsageDashboardsInterface instead.
func (a *UsageDashboardsAPI) WithImpl(impl UsageDashboardsService) UsageDashboardsInterface {
	a.impl = impl
	return a
}

// Impl returns low-level UsageDashboards API implementation
// Deprecated: use MockUsageDashboardsInterface instead.
func (a *UsageDashboardsAPI) Impl() UsageDashboardsService {
	return a.impl
}

// Create new usage dashboard.
//
// Create a usage dashboard specified by workspaceId, accountId, and dashboard
// type.
func (a *UsageDashboardsAPI) Create(ctx context.Context, request CreateBillingUsageDashboardRequest) (*CreateBillingUsageDashboardResponse, error) {
	return a.impl.Create(ctx, request)
}

// Get usage dashboard.
//
// Get a usage dashboard specified by workspaceId, accountId, and dashboard
// type.
func (a *UsageDashboardsAPI) Get(ctx context.Context, request GetBillingUsageDashboardRequest) (*GetBillingUsageDashboardResponse, error) {
	return a.impl.Get(ctx, request)
}
