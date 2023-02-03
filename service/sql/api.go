// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Alerts, Dashboards, Data Sources, Dbsql Permissions, Queries, Query History, Statement Execution, Warehouses, etc.
package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewAlerts(client *client.DatabricksClient) *AlertsAPI {
	return &AlertsAPI{
		impl: &alertsImpl{
			client: client,
		},
	}
}

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or alert destinations if
// the condition was met.
//
// **Note**: Programmatic operations on refresh schedules via the Databricks SQL
// API are deprecated. Alert refresh schedules can be created, updated, fetched
// and deleted using Jobs API, e.g. :method:jobs/create.
type AlertsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(AlertsService)
	impl AlertsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *AlertsAPI) WithImpl(impl AlertsService) *AlertsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Alerts API implementation
func (a *AlertsAPI) Impl() AlertsService {
	return a.impl
}

// Create an alert.
//
// Creates an alert. An alert is a Databricks SQL object that periodically runs
// a query, evaluates a condition of its result, and notifies users or alert
// destinations if the condition was met.
func (a *AlertsAPI) Create(ctx context.Context, request CreateAlert) (*Alert, error) {
	return a.impl.Create(ctx, request)
}

// [DEPRECATED] Create a refresh schedule.
//
// Creates a new refresh schedule for an alert.
//
// **Note:** The structure of refresh schedules is subject to change.
//
// **Note:** This API is deprecated: Use :method:jobs/create to create a job
// with the alert.
func (a *AlertsAPI) CreateSchedule(ctx context.Context, request CreateRefreshSchedule) (*RefreshSchedule, error) {
	return a.impl.CreateSchedule(ctx, request)
}

// Delete an alert.
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved to
// the trash.
func (a *AlertsAPI) Delete(ctx context.Context, request DeleteAlertRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete an alert.
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved to
// the trash.
func (a *AlertsAPI) DeleteByAlertId(ctx context.Context, alertId string) error {
	return a.impl.Delete(ctx, DeleteAlertRequest{
		AlertId: alertId,
	})
}

// [DEPRECATED] Delete a refresh schedule.
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
//
// **Note:** This API is deprecated: Use :method:jobs/delete to delete a job for
// the alert.
func (a *AlertsAPI) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	return a.impl.DeleteSchedule(ctx, request)
}

// [DEPRECATED] Delete a refresh schedule.
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
//
// **Note:** This API is deprecated: Use :method:jobs/delete to delete a job for
// the alert.
func (a *AlertsAPI) DeleteScheduleByAlertIdAndScheduleId(ctx context.Context, alertId string, scheduleId string) error {
	return a.impl.DeleteSchedule(ctx, DeleteScheduleRequest{
		AlertId:    alertId,
		ScheduleId: scheduleId,
	})
}

// Get an alert.
//
// Gets an alert.
func (a *AlertsAPI) Get(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	return a.impl.Get(ctx, request)
}

// Get an alert.
//
// Gets an alert.
func (a *AlertsAPI) GetByAlertId(ctx context.Context, alertId string) (*Alert, error) {
	return a.impl.Get(ctx, GetAlertRequest{
		AlertId: alertId,
	})
}

// [DEPRECATED] Get an alert's subscriptions.
//
// Get the subscriptions for an alert. An alert subscription represents exactly
// one recipient being notified whenever the alert is triggered. The alert
// recipient is specified by either the `user` field or the `destination` field.
// The `user` field is ignored if `destination` is non-`null`.
//
// **Note:** This API is deprecated: Use :method:jobs/get to get the
// subscriptions associated with a job for an alert.
func (a *AlertsAPI) GetSubscriptions(ctx context.Context, request GetSubscriptionsRequest) ([]Subscription, error) {
	return a.impl.GetSubscriptions(ctx, request)
}

// [DEPRECATED] Get an alert's subscriptions.
//
// Get the subscriptions for an alert. An alert subscription represents exactly
// one recipient being notified whenever the alert is triggered. The alert
// recipient is specified by either the `user` field or the `destination` field.
// The `user` field is ignored if `destination` is non-`null`.
//
// **Note:** This API is deprecated: Use :method:jobs/get to get the
// subscriptions associated with a job for an alert.
func (a *AlertsAPI) GetSubscriptionsByAlertId(ctx context.Context, alertId string) ([]Subscription, error) {
	return a.impl.GetSubscriptions(ctx, GetSubscriptionsRequest{
		AlertId: alertId,
	})
}

// Get alerts.
//
// Gets a list of alerts.
func (a *AlertsAPI) List(ctx context.Context) ([]Alert, error) {
	return a.impl.List(ctx)
}

// AlertNameToIdMap calls [AlertsAPI.List] and creates a map of results with [Alert].Name as key and [Alert].Id as value.
//
// Returns an error if there's more than one [Alert] with the same .Name.
//
// Note: All [Alert] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AlertsAPI) AlertNameToIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByName calls [AlertsAPI.AlertNameToIdMap] and returns a single [Alert].
//
// Returns an error if there's more than one [Alert] with the same .Name.
//
// Note: All [Alert] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AlertsAPI) GetByName(ctx context.Context, name string) (*Alert, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]Alert{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("Alert named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of Alert named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// [DEPRECATED] Get refresh schedules.
//
// Gets the refresh schedules for the specified alert. Alerts can have refresh
// schedules that specify when to refresh and evaluate the associated query
// result.
//
// **Note:** Although refresh schedules are returned in a list, only one refresh
// schedule per alert is currently supported. The structure of refresh schedules
// is subject to change.
//
// **Note:** This API is deprecated: Use :method:jobs/list to list jobs and
// filter by the alert.
func (a *AlertsAPI) ListSchedules(ctx context.Context, request ListSchedulesRequest) ([]RefreshSchedule, error) {
	return a.impl.ListSchedules(ctx, request)
}

// [DEPRECATED] Get refresh schedules.
//
// Gets the refresh schedules for the specified alert. Alerts can have refresh
// schedules that specify when to refresh and evaluate the associated query
// result.
//
// **Note:** Although refresh schedules are returned in a list, only one refresh
// schedule per alert is currently supported. The structure of refresh schedules
// is subject to change.
//
// **Note:** This API is deprecated: Use :method:jobs/list to list jobs and
// filter by the alert.
func (a *AlertsAPI) ListSchedulesByAlertId(ctx context.Context, alertId string) ([]RefreshSchedule, error) {
	return a.impl.ListSchedules(ctx, ListSchedulesRequest{
		AlertId: alertId,
	})
}

// [DEPRECATED] Subscribe to an alert.
//
// **Note:** This API is deprecated: Use :method:jobs/update to subscribe to a
// job for an alert.
func (a *AlertsAPI) Subscribe(ctx context.Context, request CreateSubscription) (*Subscription, error) {
	return a.impl.Subscribe(ctx, request)
}

// [DEPRECATED] Unsubscribe to an alert.
//
// Unsubscribes a user or a destination to an alert.
//
// **Note:** This API is deprecated: Use :method:jobs/update to unsubscribe to a
// job for an alert.
func (a *AlertsAPI) Unsubscribe(ctx context.Context, request UnsubscribeRequest) error {
	return a.impl.Unsubscribe(ctx, request)
}

// [DEPRECATED] Unsubscribe to an alert.
//
// Unsubscribes a user or a destination to an alert.
//
// **Note:** This API is deprecated: Use :method:jobs/update to unsubscribe to a
// job for an alert.
func (a *AlertsAPI) UnsubscribeByAlertIdAndSubscriptionId(ctx context.Context, alertId string, subscriptionId string) error {
	return a.impl.Unsubscribe(ctx, UnsubscribeRequest{
		AlertId:        alertId,
		SubscriptionId: subscriptionId,
	})
}

// Update an alert.
//
// Updates an alert.
func (a *AlertsAPI) Update(ctx context.Context, request EditAlert) error {
	return a.impl.Update(ctx, request)
}

func NewDashboards(client *client.DatabricksClient) *DashboardsAPI {
	return &DashboardsAPI{
		impl: &dashboardsImpl{
			client: client,
		},
	}
}

// In general, there is little need to modify dashboards using the API. However,
// it can be useful to use dashboard objects to look-up a collection of related
// query IDs. The API can also be used to duplicate multiple dashboards at once
// since you can get a dashboard definition with a GET request and then POST it
// to create a new one.
//
// **Note**: Programmatic operations on refresh schedules via the Databricks SQL
// API are deprecated. Dashboard refresh schedules can be created, updated,
// fetched and deleted using Jobs API, e.g. :method:jobs/create.
type DashboardsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(DashboardsService)
	impl DashboardsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *DashboardsAPI) WithImpl(impl DashboardsService) *DashboardsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Dashboards API implementation
func (a *DashboardsAPI) Impl() DashboardsService {
	return a.impl
}

// Create a dashboard object.
func (a *DashboardsAPI) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	return a.impl.Create(ctx, request)
}

// Remove a dashboard.
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) Delete(ctx context.Context, request DeleteDashboardRequest) error {
	return a.impl.Delete(ctx, request)
}

// Remove a dashboard.
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteByDashboardId(ctx context.Context, dashboardId string) error {
	return a.impl.Delete(ctx, DeleteDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Retrieve a definition.
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	return a.impl.Get(ctx, request)
}

// Retrieve a definition.
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.impl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard objects.
//
// Fetch a paginated list of dashboard objects.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DashboardsAPI) ListAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	var results []Dashboard
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	// deduplicate items that may have been added during iteration
	seen := map[string]bool{}
	request.Page = 1 // start iterating from the first page
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Results) == 0 {
			break
		}
		for _, v := range response.Results {
			id := v.Id
			if seen[id] {
				// item was added during iteration
				continue
			}
			seen[id] = true
			results = append(results, v)
		}
		request.Page++
	}
	return results, nil
}

// DashboardNameToIdMap calls [DashboardsAPI.ListAll] and creates a map of results with [Dashboard].Name as key and [Dashboard].Id as value.
//
// Returns an error if there's more than one [Dashboard] with the same .Name.
//
// Note: All [Dashboard] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DashboardsAPI) DashboardNameToIdMap(ctx context.Context, request ListDashboardsRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByName calls [DashboardsAPI.DashboardNameToIdMap] and returns a single [Dashboard].
//
// Returns an error if there's more than one [Dashboard] with the same .Name.
//
// Note: All [Dashboard] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DashboardsAPI) GetByName(ctx context.Context, name string) (*Dashboard, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListDashboardsRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]Dashboard{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("Dashboard named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of Dashboard named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Restore a dashboard.
//
// A restored dashboard appears in list views and searches and can be shared.
func (a *DashboardsAPI) Restore(ctx context.Context, request RestoreDashboardRequest) error {
	return a.impl.Restore(ctx, request)
}

func NewDataSources(client *client.DatabricksClient) *DataSourcesAPI {
	return &DataSourcesAPI{
		impl: &dataSourcesImpl{
			client: client,
		},
	}
}

// This API is provided to assist you in making new query objects. When creating
// a query object, you may optionally specify a `data_source_id` for the SQL
// warehouse against which it will run. If you don't already know the
// `data_source_id` for your desired SQL warehouse, this API will help you find
// it.
//
// This API does not support searches. It returns the full list of SQL
// warehouses in your workspace. We advise you to use any text editor, REST
// client, or `grep` to search the response from this API for the name of your
// SQL warehouse as it appears in Databricks SQL.
type DataSourcesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(DataSourcesService)
	impl DataSourcesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *DataSourcesAPI) WithImpl(impl DataSourcesService) *DataSourcesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level DataSources API implementation
func (a *DataSourcesAPI) Impl() DataSourcesService {
	return a.impl
}

// Get a list of SQL warehouses.
//
// Retrieves a full list of SQL warehouses available in this workspace. All
// fields that appear in this API response are enumerated for clarity. However,
// you need only a SQL warehouse's `id` to create new queries against it.
func (a *DataSourcesAPI) List(ctx context.Context) ([]DataSource, error) {
	return a.impl.List(ctx)
}

// DataSourceNameToIdMap calls [DataSourcesAPI.List] and creates a map of results with [DataSource].Name as key and [DataSource].Id as value.
//
// Returns an error if there's more than one [DataSource] with the same .Name.
//
// Note: All [DataSource] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DataSourcesAPI) DataSourceNameToIdMap(ctx context.Context) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByName calls [DataSourcesAPI.DataSourceNameToIdMap] and returns a single [DataSource].
//
// Returns an error if there's more than one [DataSource] with the same .Name.
//
// Note: All [DataSource] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DataSourcesAPI) GetByName(ctx context.Context, name string) (*DataSource, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.List(ctx)
	if err != nil {
		return nil, err
	}
	tmp := map[string][]DataSource{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("DataSource named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of DataSource named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

func NewDbsqlPermissions(client *client.DatabricksClient) *DbsqlPermissionsAPI {
	return &DbsqlPermissionsAPI{
		impl: &dbsqlPermissionsImpl{
			client: client,
		},
	}
}

// The SQL Permissions API is similar to the endpoints of the
// :method:permissions/set. However, this exposes only one endpoint, which gets
// the Access Control List for a given object. You cannot modify any permissions
// using this API.
//
// There are three levels of permission:
//
// - `CAN_VIEW`: Allows read-only access
//
// - `CAN_RUN`: Allows read access and run access (superset of `CAN_VIEW`)
//
// - `CAN_MANAGE`: Allows all actions: read, run, edit, delete, modify
// permissions (superset of `CAN_RUN`)
type DbsqlPermissionsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(DbsqlPermissionsService)
	impl DbsqlPermissionsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *DbsqlPermissionsAPI) WithImpl(impl DbsqlPermissionsService) *DbsqlPermissionsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level DbsqlPermissions API implementation
func (a *DbsqlPermissionsAPI) Impl() DbsqlPermissionsService {
	return a.impl
}

// Get object ACL.
//
// Gets a JSON representation of the access control list (ACL) for a specified
// object.
func (a *DbsqlPermissionsAPI) Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error) {
	return a.impl.Get(ctx, request)
}

// Get object ACL.
//
// Gets a JSON representation of the access control list (ACL) for a specified
// object.
func (a *DbsqlPermissionsAPI) GetByObjectTypeAndObjectId(ctx context.Context, objectType ObjectTypePlural, objectId string) (*GetResponse, error) {
	return a.impl.Get(ctx, GetDbsqlPermissionRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
}

// Set object ACL.
//
// Sets the access control list (ACL) for a specified object. This operation
// will complete rewrite the ACL.
func (a *DbsqlPermissionsAPI) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	return a.impl.Set(ctx, request)
}

// Transfer object ownership.
//
// Transfers ownership of a dashboard, query, or alert to an active user.
// Requires an admin API key.
func (a *DbsqlPermissionsAPI) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	return a.impl.TransferOwnership(ctx, request)
}

func NewQueries(client *client.DatabricksClient) *QueriesAPI {
	return &QueriesAPI{
		impl: &queriesImpl{
			client: client,
		},
	}
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, execution schedule, parameters, and visualizations.
//
// **Note**: Programmatic operations on refresh schedules via the Databricks SQL
// API are deprecated. Query refresh schedules can be created, updated, fetched
// and deleted using Jobs API, e.g. :method:jobs/create.
type QueriesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(QueriesService)
	impl QueriesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *QueriesAPI) WithImpl(impl QueriesService) *QueriesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Queries API implementation
func (a *QueriesAPI) Impl() QueriesService {
	return a.impl
}

// Create a new query definition.
//
// Creates a new query definition. Queries created with this endpoint belong to
// the authenticated user making the request.
//
// The `data_source_id` field specifies the ID of the SQL warehouse to run this
// query against. You can use the Data Sources API to see a complete list of
// available SQL warehouses. Or you can copy the `data_source_id` from an
// existing query.
//
// **Note**: You cannot add a visualization until you create the query.
func (a *QueriesAPI) Create(ctx context.Context, request QueryPostContent) (*Query, error) {
	return a.impl.Create(ctx, request)
}

// Delete a query.
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) Delete(ctx context.Context, request DeleteQueryRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a query.
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteByQueryId(ctx context.Context, queryId string) error {
	return a.impl.Delete(ctx, DeleteQueryRequest{
		QueryId: queryId,
	})
}

// Get a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) Get(ctx context.Context, request GetQueryRequest) (*Query, error) {
	return a.impl.Get(ctx, request)
}

// Get a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetByQueryId(ctx context.Context, queryId string) (*Query, error) {
	return a.impl.Get(ctx, GetQueryRequest{
		QueryId: queryId,
	})
}

// Get a list of queries.
//
// Gets a list of queries. Optionally, this list can be filtered by a search
// term.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueriesAPI) ListAll(ctx context.Context, request ListQueriesRequest) ([]Query, error) {
	var results []Query
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	// deduplicate items that may have been added during iteration
	seen := map[string]bool{}
	request.Page = 1 // start iterating from the first page
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Results) == 0 {
			break
		}
		for _, v := range response.Results {
			id := v.Id
			if seen[id] {
				// item was added during iteration
				continue
			}
			seen[id] = true
			results = append(results, v)
		}
		request.Page++
	}
	return results, nil
}

// QueryNameToIdMap calls [QueriesAPI.ListAll] and creates a map of results with [Query].Name as key and [Query].Id as value.
//
// Returns an error if there's more than one [Query] with the same .Name.
//
// Note: All [Query] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueriesAPI) QueryNameToIdMap(ctx context.Context, request ListQueriesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByName calls [QueriesAPI.QueryNameToIdMap] and returns a single [Query].
//
// Returns an error if there's more than one [Query] with the same .Name.
//
// Note: All [Query] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueriesAPI) GetByName(ctx context.Context, name string) (*Query, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListQueriesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]Query{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("Query named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of Query named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Restore a query.
//
// Restore a query that has been moved to the trash. A restored query appears in
// list views and searches. You can use restored queries for alerts.
func (a *QueriesAPI) Restore(ctx context.Context, request RestoreQueryRequest) error {
	return a.impl.Restore(ctx, request)
}

// Change a query definition.
//
// Modify this query definition.
//
// **Note**: You cannot undo this operation.
func (a *QueriesAPI) Update(ctx context.Context, request QueryEditContent) (*Query, error) {
	return a.impl.Update(ctx, request)
}

func NewQueryHistory(client *client.DatabricksClient) *QueryHistoryAPI {
	return &QueryHistoryAPI{
		impl: &queryHistoryImpl{
			client: client,
		},
	}
}

// Access the history of queries through SQL warehouses.
type QueryHistoryAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(QueryHistoryService)
	impl QueryHistoryService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *QueryHistoryAPI) WithImpl(impl QueryHistoryService) *QueryHistoryAPI {
	a.impl = impl
	return a
}

// Impl returns low-level QueryHistory API implementation
func (a *QueryHistoryAPI) Impl() QueryHistoryService {
	return a.impl
}

// List Queries.
//
// List the history of queries through SQL warehouses.
//
// You can filter by user ID, warehouse ID, status, and time range.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueryHistoryAPI) ListAll(ctx context.Context, request ListQueryHistoryRequest) ([]QueryInfo, error) {
	var results []QueryInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Res) == 0 {
			break
		}
		for _, v := range response.Res {
			results = append(results, v)
		}
		request.PageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

func NewStatementExecution(client *client.DatabricksClient) *StatementExecutionAPI {
	return &StatementExecutionAPI{
		impl: &statementExecutionImpl{
			client: client,
		},
	}
}

// The SQL Statement Execution API manages the execution of arbitrary SQL
// statements and the fetching of result data.
//
// # Release Status
//
// This feature is in [Private Preview]. To try it, reach out to your Databricks
// contact.
//
// # Getting started
//
// We suggest beginning with the [SQL Statement Execution API tutorial].
//
// # Overview of statement execution and result fetching
//
// Statement execution begins by calling
// :method:StatementExecution/executeStatement with a valid SQL statement and
// warehouse ID, along with optional parameters such as the data catalog and
// output format.
//
// When submitting the statement, the call can behave synchronously or
// asynchronously, based on the `wait_timeout` setting. When set between 5-50
// seconds (default: 10) the call behaves synchronously; when set to `0s`, the
// call is asynchronous and responds immediately if accepted.
//
// ----
//
// ### **Warning: drop authorization header when fetching data through external
// links**
//
// External link URLs do not require an Authorization header or token, and thus
// all calls to fetch external links must remove the Authorization header.
//
// ----
//
// Similar to INLINE mode, callers can iterate through the result set, by using
// the field `next_chunk_internal_link`. Each internal link response will
// contain an external link to the raw chunk data, and additionally contain the
// next_chunk_internal_link if there are more chunks.
//
// Unlike INLINE mode, when using EXTERNAL_LINKS, chunks may be fetched out of
// order, and in parallel to achieve higher throughput.
//
// # Limits and limitations
//
// - All byte limits are calculated based on internal storage metrics, and will
// not match byte counts of actual payloads. - INLINE mode statements limited to
// 16 MiB, and will abort when this limit is exceeded. - Cancelation may
// silently fail: A successful response from a cancel request indicates that the
// cancel request was successfully received and sent to the processing engine.
// However, for example, an outstanding statement may complete execution during
// signal delivery, with cancel signal arriving too late to be meaningful.
// Polling for status until a terminal state is reached a reliable way to see
// final state. - Wait timeouts are approximate, occur server-side, and cannot
// account for caller delays, network latency from caller to service, and
// similarly. - After a statement has been submitted and a statement_id
// produced, that statement's status and result will automatically close after
// either of 2 conditions: - The last result chunk is fetched (or resolved to an
// external link). - Ten (10) minutes pass with no calls to get status or fetch
// result data. Best practice: in asynchronous clients, poll for status
// regularly (and with backoff) to keep the statement open and alive.
//
// # Private Preview limitations
//
// - `EXTERNAL_LINKS` mode will fail for result sets < 5MB. - After any cancel
// or close operation, the statement will no longer be visible from the API,
// specifically - After fetching last result chunk (including `chunk_index=0`),
// the statement is closed; a short time after closure, the statement will no
// longer be visible to the API, and further calls may return 404. Thus calling
// GET .../{statement_id} will return a 404 NOT FOUND error. - In practice, this
// means that a CANCEL and subsequent poll will often return a NOT FOUND. This
// will be addressed in a future update.
//
// [Private Preview]: https://docs.databricks.com/release-notes/release-types.html
// [SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
type StatementExecutionAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(StatementExecutionService)
	impl StatementExecutionService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *StatementExecutionAPI) WithImpl(impl StatementExecutionService) *StatementExecutionAPI {
	a.impl = impl
	return a
}

// Impl returns low-level StatementExecution API implementation
func (a *StatementExecutionAPI) Impl() StatementExecutionService {
	return a.impl
}

// Cancel statement execution.
//
// Requests that an executing statement be cancelled. Callers must poll for
// status to see terminal state.
func (a *StatementExecutionAPI) CancelExecution(ctx context.Context, request CancelExecutionRequest) error {
	return a.impl.CancelExecution(ctx, request)
}

// Execute an SQL statement.
//
// Execute an SQL statement, and if flagged as such, await its result for a
// specified time.
//
// ## Call mode: synchronous
//
// In synchronous mode, when statement execution completes, making the result
// available within the wait timeout, result data is returned in the response.
// This response will contain `statement_id`, `status`, `manifest`, and `result`
// fields. `status` will confirm success, and `manifest` contains both the
// result data column schema, and metadata about the result set. `result` will
// contain the first chunk of result data according to the specified
// disposition, and links to fetch any remaining chunks.
//
// If execution does not complete before `wait_timeout`, a response will be
// returned immediately. The setting `on_wait_timeout` determines how the system
// responds.
//
// By default, `on_wait_timeout=CONTINUE`, and after reaching timeout, a
// response is sent and statement execution continues asynchronously. The
// response will contain only `statement_id` and `status` fields, and caller
// must now follow the flow described for asynchronous call mode to poll and
// fetch result.
//
// Alternatively, `on_wait_timeout` can also be set to `CANCEL`; in this case if
// the timeout is reached before execution completes, the underlying statement
// execution is canceled, and a `CANCELED` status is returned in the response.
//
// ## Call mode: asynchronous
//
// In asynchronous mode, or after a timed-out synchronous request continues, a
// `statement_id` and `status` will be returned. In this case polling
// :method:StatementExecution/getStatement calls are required to fetch result
// and metadata.
//
// ### GET statement calls
//
// Next a caller must poll until execution completes (SUCCEEDED, FAILED, etc.).
// Given a `statement_id`, poll by calling
// :method:StatementExecution/getStatement.
//
// When execution has succeeded, the response will contain `status`, `manifest`,
// and `result` fields. These fields and structure are identical to those in the
// response to a successful synchronous submission. `result` will contain the
// first chunk of result data, either inline or as external links depending on
// disposition. Additional chunks of result data can be fetched by checking for
// the presence of the `next_chunk_internal_link` field, and iteratively
// :method:StatementExecution/getStatement those paths until that field is
// unset: `GET https://$DATABRICKS_HOST/{next_chunk_internal_link}`.
//
// ## Fetching result data: format and disposition
//
// Result data from statement execution is available in two formats: JSON, and
// [Apache Arrow Columnar]. Statements producing a result set smaller than 16
// MiB can be fetched as `format=JSON_ARRAY`, using the `disposition=INLINE`.
// When a statement executed in INLINE disposition exceeds this limit, execution
// is aborted, and no result can be fetched. Using `format=ARROW_STREAM` and
// `disposition=EXTERNAL_LINKS` allows large result sets to be fetched, and with
// higher throughput.
//
// The API uses defaults of `format=JSON_ARRAY` and `disposition=INLINE`. We
// advise explicitly setting format and disposition in all production use cases.
//
// ## Statement response: statement_id, status, manifest, and result
//
// The base call :method:StatementExecution/getStatement returns a single
// response combining statement_id, status, a result manifest, and a result data
// chunk or link. The manifest contains the result schema definition, and result
// summary metadata. When using EXTERNAL_LINKS disposition, it also contains a
// full listing of all chunks and their summary metadata.
//
// # Use case: small result sets with INLINE + JSON_ARRAY
//
// For flows which will generate small and predictable result sets (<= 16 MiB),
// INLINE downloads of JSON_ARRAY result data is typically the simplest way to
// execute and fetch result data. In this case,
// `:method:StatementExecution/executeStatement, along with a `warehouse_id`
// (required) and any other desired options. With default parameters, (noteably
// `wait_timeout=10s`), execution and result fetch are synchronous: a small
// result will be returned in the response, if completed within 10 seconds.
// `wait_timeout` can be extended up to 50 seconds.
//
// When result set in INLINE mode becomes larger, it will be transfer results in
// chunks, each up to 4 MiB. After receiving the initial chunk get
// :method:StatementExecution/executeStatement or
// :method:StatementExecution/getStatement, subsequent calls are required to
// iteratively fetch each chunk. Each result response contains link to the next
// chunk, when there are additiona chunks remaining; it can be found in the
// field `.next_chunk_internal_link`. This link is an absolute `path` to be
// joined with your `$DATABRICKS_HOST`, and of the form
// `/api/2.0/sql/statements/{statement_id}/result/chunks/...`. The next chunk
// can be fetched like this `GET
// https://$DATABRICKS_HOST/{next_chunk_internal_link}`.
//
// When using this mode, each chunk may be fetched once, and in order. If a
// chunk has no field `.next_chunk_internal_link`, that indicates it to be the
// last chunk, and all chunks have been fetched from the result set.
//
// # Use case: large result sets with EXTERNAL_LINKS + ARROW_STREAM
//
// Using EXTERNAL_LINKS to fetch result data in Arrow format allows you to fetch
// large result sets efficiently. The primary difference from using INLINE
// disposition is that fetched result chunks contain resolved `external_links`
// URLs, which can be fetched with standard HTTP.
//
// ## Presigned URLs
//
// External links point to data stored within your workspace's internal DBFS, in
// the form of a presigned URL. The URLs are valid for only a short period, <=
// 15 minutes. Alongside each external_link is an expiration field indicating
// the time at which the URL is no longer valid. In EXTERNAL_LINKS mode, chunks
// be resolved and fetched multiple time, and in parallel.
//
// [Apache Arrow Columnar]: https://arrow.apache.org/overview/
func (a *StatementExecutionAPI) ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*ExecuteStatementResponse, error) {
	return a.impl.ExecuteStatement(ctx, request)
}

// Get status, manifest, and result first chunk.
//
// Polls for statement status; when status.state=SUCCEEDED will also return
// result manifest, and the first chunk of result data.
//
// **NOTE** This call currently may take up to 5 seconds to get latest status
// and result.
func (a *StatementExecutionAPI) GetStatement(ctx context.Context, request GetStatementRequest) (*GetStatementResponse, error) {
	return a.impl.GetStatement(ctx, request)
}

// Get status, manifest, and result first chunk.
//
// Polls for statement status; when status.state=SUCCEEDED will also return
// result manifest, and the first chunk of result data.
//
// **NOTE** This call currently may take up to 5 seconds to get latest status
// and result.
func (a *StatementExecutionAPI) GetStatementByStatementId(ctx context.Context, statementId string) (*GetStatementResponse, error) {
	return a.impl.GetStatement(ctx, GetStatementRequest{
		StatementId: statementId,
	})
}

// Get result chunk by index.
//
// After statement execution has SUCCEEDED, result data can be fetched by
// chunks.
//
// The first chunk (`chunk_index=0`) is typically fetched through
// `getStatementResult`, and subsequent chunks with this call. The response
// structure is identical to the nested `result` element described in
// getStatementResult, and similarly includes `next_chunk_index` and
// `next_chunk_internal_link` for simple iteration through the result set.
func (a *StatementExecutionAPI) GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error) {
	return a.impl.GetStatementResultChunkN(ctx, request)
}

// Get result chunk by index.
//
// After statement execution has SUCCEEDED, result data can be fetched by
// chunks.
//
// The first chunk (`chunk_index=0`) is typically fetched through
// `getStatementResult`, and subsequent chunks with this call. The response
// structure is identical to the nested `result` element described in
// getStatementResult, and similarly includes `next_chunk_index` and
// `next_chunk_internal_link` for simple iteration through the result set.
func (a *StatementExecutionAPI) GetStatementResultChunkNByStatementIdAndChunkIndex(ctx context.Context, statementId string, chunkIndex int) (*ResultData, error) {
	return a.impl.GetStatementResultChunkN(ctx, GetStatementResultChunkNRequest{
		StatementId: statementId,
		ChunkIndex:  chunkIndex,
	})
}

func NewWarehouses(client *client.DatabricksClient) *WarehousesAPI {
	return &WarehousesAPI{
		impl: &warehousesImpl{
			client: client,
		},
	}
}

// A SQL warehouse is a compute resource that lets you run SQL commands on data
// objects within Databricks SQL. Compute resources are infrastructure resources
// that provide processing capabilities in the cloud.
type WarehousesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WarehousesService)
	impl WarehousesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WarehousesAPI) WithImpl(impl WarehousesService) *WarehousesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Warehouses API implementation
func (a *WarehousesAPI) Impl() WarehousesService {
	return a.impl
}

// Create a warehouse.
//
// Creates a new SQL warehouse.
func (a *WarehousesAPI) Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	return a.impl.Create(ctx, request)
}

// Calls [WarehousesAPI.Create] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) CreateAndWait(ctx context.Context, createWarehouseRequest CreateWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	createWarehouseResponse, err := a.Create(ctx, createWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.Get(ctx, GetWarehouseRequest{
			Id: createWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if getWarehouseResponse.Health != nil {
			statusMessage = getWarehouseResponse.Health.Summary
		}
		switch status {
		case StateRunning: // target state
			return getWarehouseResponse, nil
		case StateStopped, StateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete a warehouse.
//
// Deletes a SQL warehouse.
func (a *WarehousesAPI) Delete(ctx context.Context, request DeleteWarehouseRequest) error {
	return a.impl.Delete(ctx, request)
}

// Calls [WarehousesAPI.Delete] and waits to reach DELETED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) DeleteAndWait(ctx context.Context, deleteWarehouseRequest DeleteWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Delete(ctx, deleteWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.Get(ctx, GetWarehouseRequest{
			Id: deleteWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if getWarehouseResponse.Health != nil {
			statusMessage = getWarehouseResponse.Health.Summary
		}
		switch status {
		case StateDeleted: // target state
			return getWarehouseResponse, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Delete a warehouse.
//
// Deletes a SQL warehouse.
func (a *WarehousesAPI) DeleteById(ctx context.Context, id string) error {
	return a.impl.Delete(ctx, DeleteWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) DeleteByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	return a.DeleteAndWait(ctx, DeleteWarehouseRequest{
		Id: id,
	}, options...)
}

// Update a warehouse.
//
// Updates the configuration for a SQL warehouse.
func (a *WarehousesAPI) Edit(ctx context.Context, request EditWarehouseRequest) error {
	return a.impl.Edit(ctx, request)
}

// Calls [WarehousesAPI.Edit] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) EditAndWait(ctx context.Context, editWarehouseRequest EditWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Edit(ctx, editWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.Get(ctx, GetWarehouseRequest{
			Id: editWarehouseRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if getWarehouseResponse.Health != nil {
			statusMessage = getWarehouseResponse.Health.Summary
		}
		switch status {
		case StateRunning: // target state
			return getWarehouseResponse, nil
		case StateStopped, StateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get warehouse info.
//
// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error) {
	return a.impl.Get(ctx, request)
}

// Calls [WarehousesAPI.Get] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) GetAndWait(ctx context.Context, getWarehouseRequest GetWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	getWarehouseResponse, err := a.Get(ctx, getWarehouseRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.Get(ctx, GetWarehouseRequest{
			Id: getWarehouseResponse.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if getWarehouseResponse.Health != nil {
			statusMessage = getWarehouseResponse.Health.Summary
		}
		switch status {
		case StateRunning: // target state
			return getWarehouseResponse, nil
		case StateStopped, StateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Get warehouse info.
//
// Gets the information for a single SQL warehouse.
func (a *WarehousesAPI) GetById(ctx context.Context, id string) (*GetWarehouseResponse, error) {
	return a.impl.Get(ctx, GetWarehouseRequest{
		Id: id,
	})
}

func (a *WarehousesAPI) GetByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	return a.GetAndWait(ctx, GetWarehouseRequest{
		Id: id,
	}, options...)
}

// Get the workspace configuration.
//
// Gets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error) {
	return a.impl.GetWorkspaceWarehouseConfig(ctx)
}

// List warehouses.
//
// Lists all SQL warehouses that a user has manager permissions on.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WarehousesAPI) ListAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error) {
	response, err := a.impl.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.Warehouses, nil
}

// EndpointInfoNameToIdMap calls [WarehousesAPI.ListAll] and creates a map of results with [EndpointInfo].Name as key and [EndpointInfo].Id as value.
//
// Returns an error if there's more than one [EndpointInfo] with the same .Name.
//
// Note: All [EndpointInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WarehousesAPI) EndpointInfoNameToIdMap(ctx context.Context, request ListWarehousesRequest) (map[string]string, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]string{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Name
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Name: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetByName calls [WarehousesAPI.EndpointInfoNameToIdMap] and returns a single [EndpointInfo].
//
// Returns an error if there's more than one [EndpointInfo] with the same .Name.
//
// Note: All [EndpointInfo] instances are loaded into memory before returning matching by name.
//
// This method is generated by Databricks SDK Code Generator.
func (a *WarehousesAPI) GetByName(ctx context.Context, name string) (*EndpointInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListWarehousesRequest{})
	if err != nil {
		return nil, err
	}
	tmp := map[string][]EndpointInfo{}
	for _, v := range result {
		key := v.Name
		tmp[key] = append(tmp[key], v)
	}
	alternatives, ok := tmp[name]
	if !ok || len(alternatives) == 0 {
		return nil, fmt.Errorf("EndpointInfo named '%s' does not exist", name)
	}
	if len(alternatives) > 1 {
		return nil, fmt.Errorf("there are %d instances of EndpointInfo named '%s'", len(alternatives), name)
	}
	return &alternatives[0], nil
}

// Set the workspace configuration.
//
// Sets the workspace level configuration that is shared by all SQL warehouses
// in a workspace.
func (a *WarehousesAPI) SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error {
	return a.impl.SetWorkspaceWarehouseConfig(ctx, request)
}

// Start a warehouse.
//
// Starts a SQL warehouse.
func (a *WarehousesAPI) Start(ctx context.Context, request StartRequest) error {
	return a.impl.Start(ctx, request)
}

// Calls [WarehousesAPI.Start] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) StartAndWait(ctx context.Context, startRequest StartRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Start(ctx, startRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.Get(ctx, GetWarehouseRequest{
			Id: startRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if getWarehouseResponse.Health != nil {
			statusMessage = getWarehouseResponse.Health.Summary
		}
		switch status {
		case StateRunning: // target state
			return getWarehouseResponse, nil
		case StateStopped, StateDeleted:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				StateRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// Stop a warehouse.
//
// Stops a SQL warehouse.
func (a *WarehousesAPI) Stop(ctx context.Context, request StopRequest) error {
	return a.impl.Stop(ctx, request)
}

// Calls [WarehousesAPI.Stop] and waits to reach STOPPED state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[GetWarehouseResponse](60*time.Minute) functional option.
func (a *WarehousesAPI) StopAndWait(ctx context.Context, stopRequest StopRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	err := a.Stop(ctx, stopRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[GetWarehouseResponse]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[GetWarehouseResponse](ctx, i.Timeout, func() (*GetWarehouseResponse, *retries.Err) {
		getWarehouseResponse, err := a.Get(ctx, GetWarehouseRequest{
			Id: stopRequest.Id,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[GetWarehouseResponse]{
				Info:    getWarehouseResponse,
				Timeout: i.Timeout,
			})
		}
		status := getWarehouseResponse.State
		statusMessage := fmt.Sprintf("current status: %s", status)
		if getWarehouseResponse.Health != nil {
			statusMessage = getWarehouseResponse.Health.Summary
		}
		switch status {
		case StateStopped: // target state
			return getWarehouseResponse, nil
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}
