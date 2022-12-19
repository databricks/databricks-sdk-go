// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Alerts, Dashboards, Data Sources, Dbsql Permissions, Queries, Query History, Warehouses, etc.
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
func (a *AlertsAPI) Create(ctx context.Context, request EditAlert) (*Alert, error) {
	return a.impl.Create(ctx, request)
}

// Create a refresh schedule.
//
// Creates a new refresh schedule for an alert.
//
// **Note:** The structure of refresh schedules is subject to change.
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

// Delete a refresh schedule.
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
func (a *AlertsAPI) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	return a.impl.DeleteSchedule(ctx, request)
}

// Delete a refresh schedule.
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
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

// Get an alert's subscriptions.
//
// Get the subscriptions for an alert. An alert subscription represents exactly
// one recipient being notified whenever the alert is triggered. The alert
// recipient is specified by either the `user` field or the `destination` field.
// The `user` field is ignored if `destination` is non-`null`.
func (a *AlertsAPI) GetSubscriptions(ctx context.Context, request GetSubscriptionsRequest) ([]Subscription, error) {
	return a.impl.GetSubscriptions(ctx, request)
}

// Get an alert's subscriptions.
//
// Get the subscriptions for an alert. An alert subscription represents exactly
// one recipient being notified whenever the alert is triggered. The alert
// recipient is specified by either the `user` field or the `destination` field.
// The `user` field is ignored if `destination` is non-`null`.
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

// Get refresh schedules.
//
// Gets the refresh schedules for the specified alert. Alerts can have refresh
// schedules that specify when to refresh and evaluate the associated query
// result.
//
// **Note:** Although refresh schedules are returned in a list, only one refresh
// schedule per alert is currently supported. The structure of refresh schedules
// is subject to change.
func (a *AlertsAPI) ListSchedules(ctx context.Context, request ListSchedulesRequest) ([]RefreshSchedule, error) {
	return a.impl.ListSchedules(ctx, request)
}

// Get refresh schedules.
//
// Gets the refresh schedules for the specified alert. Alerts can have refresh
// schedules that specify when to refresh and evaluate the associated query
// result.
//
// **Note:** Although refresh schedules are returned in a list, only one refresh
// schedule per alert is currently supported. The structure of refresh schedules
// is subject to change.
func (a *AlertsAPI) ListSchedulesByAlertId(ctx context.Context, alertId string) ([]RefreshSchedule, error) {
	return a.impl.ListSchedules(ctx, ListSchedulesRequest{
		AlertId: alertId,
	})
}

// Subscribe to an alert.
func (a *AlertsAPI) Subscribe(ctx context.Context, request CreateSubscription) (*Subscription, error) {
	return a.impl.Subscribe(ctx, request)
}

// Unsubscribe to an alert.
//
// Unsubscribes a user or a destination to an alert.
func (a *AlertsAPI) Unsubscribe(ctx context.Context, request UnsubscribeRequest) error {
	return a.impl.Unsubscribe(ctx, request)
}

// Unsubscribe to an alert.
//
// Unsubscribes a user or a destination to an alert.
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
// :method:permissions/setobjectpermissions. However, this exposes only one
// endpoint, which gets the Access Control List for a given object. You cannot
// modify any permissions using this API.
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
func (a *QueriesAPI) Update(ctx context.Context, request QueryPostContent) (*Query, error) {
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
				Polling: true,
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
				Polling: true,
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
				Polling: true,
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
				Polling: true,
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
				Polling: true,
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
				Polling: true,
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
