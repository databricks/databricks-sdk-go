// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbsql

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
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

// Create an alert
//
// Creates an alert. An alert is a Databricks SQL object that periodically runs
// a query, evaluates a condition of its result, and notifies users or alert
// destinations if the condition was met.
func (a *AlertsAPI) CreateAlert(ctx context.Context, request EditAlert) (*Alert, error) {
	return a.impl.CreateAlert(ctx, request)
}

// Create a refresh schedule
//
// Creates a new refresh schedule for an alert.
//
// **Note:** The structure of refresh schedules is subject to change.
func (a *AlertsAPI) CreateSchedule(ctx context.Context, request CreateRefreshSchedule) (*RefreshSchedule, error) {
	return a.impl.CreateSchedule(ctx, request)
}

// Delete an alert
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved to
// the trash.
func (a *AlertsAPI) DeleteAlert(ctx context.Context, request DeleteAlertRequest) error {
	return a.impl.DeleteAlert(ctx, request)
}

// Delete an alert
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved to
// the trash.
func (a *AlertsAPI) DeleteAlertByAlertId(ctx context.Context, alertId string) error {
	return a.impl.DeleteAlert(ctx, DeleteAlertRequest{
		AlertId: alertId,
	})
}

// Delete a refresh schedule
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
func (a *AlertsAPI) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	return a.impl.DeleteSchedule(ctx, request)
}

// Delete a refresh schedule
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
func (a *AlertsAPI) DeleteScheduleByAlertIdAndScheduleId(ctx context.Context, alertId string, scheduleId string) error {
	return a.impl.DeleteSchedule(ctx, DeleteScheduleRequest{
		AlertId:    alertId,
		ScheduleId: scheduleId,
	})
}

// Get an alert
//
// Gets an alert.
func (a *AlertsAPI) GetAlert(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	return a.impl.GetAlert(ctx, request)
}

// Get an alert
//
// Gets an alert.
func (a *AlertsAPI) GetAlertByAlertId(ctx context.Context, alertId string) (*Alert, error) {
	return a.impl.GetAlert(ctx, GetAlertRequest{
		AlertId: alertId,
	})
}

// Get an alert's subscriptions
//
// Get the subscriptions for an alert. An alert subscription represents exactly
// one recipient being notified whenever the alert is triggered. The alert
// recipient is specified by either the `user` field or the `destination` field.
// The `user` field is ignored if `destination` is non-`null`.
func (a *AlertsAPI) GetSubscriptions(ctx context.Context, request GetSubscriptionsRequest) ([]Subscription, error) {
	return a.impl.GetSubscriptions(ctx, request)
}

// Get an alert's subscriptions
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

// Get alerts
//
// Gets a list of alerts.
func (a *AlertsAPI) ListAlerts(ctx context.Context) ([]Alert, error) {
	return a.impl.ListAlerts(ctx)
}

// AlertNameToIdMap calls [AlertsAPI.ListAlerts] and creates a map of results with [Alert].Name as key and [Alert].Id as value.
//
// Note: All [Alert] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AlertsAPI) AlertNameToIdMap(ctx context.Context) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListAlerts(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.Name] = v.Id
	}
	return mapping, nil
}

// GetAlertByName calls [AlertsAPI.AlertNameToIdMap] and returns a single [Alert].
//
// Note: All [Alert] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *AlertsAPI) GetAlertByName(ctx context.Context, name string) (*Alert, error) {
	result, err := a.ListAlerts(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.Name != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Alert named '%s' does not exist", name)
}

// Get refresh schedules
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

// Get refresh schedules
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

// Subscribe to an alert
func (a *AlertsAPI) Subscribe(ctx context.Context, request CreateSubscription) (*Subscription, error) {
	return a.impl.Subscribe(ctx, request)
}

// Unsubscribe to an alert
//
// Unsubscribes a user or a destination to an alert.
func (a *AlertsAPI) Unsubscribe(ctx context.Context, request UnsubscribeRequest) error {
	return a.impl.Unsubscribe(ctx, request)
}

// Unsubscribe to an alert
//
// Unsubscribes a user or a destination to an alert.
func (a *AlertsAPI) UnsubscribeByAlertIdAndSubscriptionId(ctx context.Context, alertId string, subscriptionId string) error {
	return a.impl.Unsubscribe(ctx, UnsubscribeRequest{
		AlertId:        alertId,
		SubscriptionId: subscriptionId,
	})
}

// Update an alert
//
// Updates an alert.
func (a *AlertsAPI) UpdateAlert(ctx context.Context, request EditAlert) error {
	return a.impl.UpdateAlert(ctx, request)
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

// Create a dashboard object
func (a *DashboardsAPI) CreateDashboard(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	return a.impl.CreateDashboard(ctx, request)
}

// Remove a dashboard
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteDashboard(ctx context.Context, request DeleteDashboardRequest) error {
	return a.impl.DeleteDashboard(ctx, request)
}

// Remove a dashboard
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteDashboardByDashboardId(ctx context.Context, dashboardId string) error {
	return a.impl.DeleteDashboard(ctx, DeleteDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Retrieve a definition
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetDashboard(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	return a.impl.GetDashboard(ctx, request)
}

// Retrieve a definition
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetDashboardByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.impl.GetDashboard(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get dashboard objects
//
// Fetch a paginated list of dashboard objects.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DashboardsAPI) ListDashboardsAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	var results []Dashboard
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	request.Page = 1 // start iterating from the first page
	for {
		response, err := a.impl.ListDashboards(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Results) == 0 {
			break
		}
		for _, v := range response.Results {
			results = append(results, v)
		}
		request.Page++
	}
	return results, nil
}

// DashboardNameToIdMap calls [DashboardsAPI.ListDashboardsAll] and creates a map of results with [Dashboard].Name as key and [Dashboard].Id as value.
//
// Note: All [Dashboard] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DashboardsAPI) DashboardNameToIdMap(ctx context.Context, request ListDashboardsRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListDashboardsAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.Name] = v.Id
	}
	return mapping, nil
}

// GetDashboardByName calls [DashboardsAPI.DashboardNameToIdMap] and returns a single [Dashboard].
//
// Note: All [Dashboard] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DashboardsAPI) GetDashboardByName(ctx context.Context, name string) (*Dashboard, error) {
	result, err := a.ListDashboardsAll(ctx, ListDashboardsRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.Name != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Dashboard named '%s' does not exist", name)
}

// Restore a dashboard
//
// A restored dashboard appears in list views and searches and can be shared.
func (a *DashboardsAPI) RestoreDashboard(ctx context.Context, request RestoreDashboardRequest) error {
	return a.impl.RestoreDashboard(ctx, request)
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

// Get a list of SQL warehouses
//
// Retrieves a full list of SQL warehouses available in this workspace. All
// fields that appear in this API response are enumerated for clarity. However,
// you need only a SQL warehouse's `id` to create new queries against it.
func (a *DataSourcesAPI) ListDataSources(ctx context.Context) ([]DataSource, error) {
	return a.impl.ListDataSources(ctx)
}

// DataSourceNameToIdMap calls [DataSourcesAPI.ListDataSources] and creates a map of results with [DataSource].Name as key and [DataSource].Id as value.
//
// Note: All [DataSource] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DataSourcesAPI) DataSourceNameToIdMap(ctx context.Context) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListDataSources(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.Name] = v.Id
	}
	return mapping, nil
}

// GetDataSourceByName calls [DataSourcesAPI.DataSourceNameToIdMap] and returns a single [DataSource].
//
// Note: All [DataSource] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *DataSourcesAPI) GetDataSourceByName(ctx context.Context, name string) (*DataSource, error) {
	result, err := a.ListDataSources(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.Name != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("DataSource named '%s' does not exist", name)
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

// Get object ACL
//
// Gets a JSON representation of the access control list (ACL) for a specified
// object.
func (a *DbsqlPermissionsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return a.impl.GetPermissions(ctx, request)
}

// Get object ACL
//
// Gets a JSON representation of the access control list (ACL) for a specified
// object.
func (a *DbsqlPermissionsAPI) GetPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType ObjectTypePlural, objectId string) (*GetPermissionsResponse, error) {
	return a.impl.GetPermissions(ctx, GetPermissionsRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
}

// Set object ACL
//
// Sets the access control list (ACL) for a specified object. This operation
// will complete rewrite the ACL.
func (a *DbsqlPermissionsAPI) SetPermissions(ctx context.Context, request SetPermissionsRequest) (*SetPermissionsResponse, error) {
	return a.impl.SetPermissions(ctx, request)
}

// Transfer object ownership
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

// Create a new query definition
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
func (a *QueriesAPI) CreateQuery(ctx context.Context, request QueryPostContent) (*Query, error) {
	return a.impl.CreateQuery(ctx, request)
}

// Delete a query
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteQuery(ctx context.Context, request DeleteQueryRequest) error {
	return a.impl.DeleteQuery(ctx, request)
}

// Delete a query
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteQueryByQueryId(ctx context.Context, queryId string) error {
	return a.impl.DeleteQuery(ctx, DeleteQueryRequest{
		QueryId: queryId,
	})
}

// Get a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetQuery(ctx context.Context, request GetQueryRequest) (*Query, error) {
	return a.impl.GetQuery(ctx, request)
}

// Get a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetQueryByQueryId(ctx context.Context, queryId string) (*Query, error) {
	return a.impl.GetQuery(ctx, GetQueryRequest{
		QueryId: queryId,
	})
}

// Get a list of queries
//
// Gets a list of queries. Optionally, this list can be filtered by a search
// term.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueriesAPI) ListQueriesAll(ctx context.Context, request ListQueriesRequest) ([]Query, error) {
	var results []Query
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	request.Page = 1 // start iterating from the first page
	for {
		response, err := a.impl.ListQueries(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Results) == 0 {
			break
		}
		for _, v := range response.Results {
			results = append(results, v)
		}
		request.Page++
	}
	return results, nil
}

// QueryNameToIdMap calls [QueriesAPI.ListQueriesAll] and creates a map of results with [Query].Name as key and [Query].Id as value.
//
// Note: All [Query] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueriesAPI) QueryNameToIdMap(ctx context.Context, request ListQueriesRequest) (map[string]string, error) {
	mapping := map[string]string{}
	result, err := a.ListQueriesAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		mapping[v.Name] = v.Id
	}
	return mapping, nil
}

// GetQueryByName calls [QueriesAPI.QueryNameToIdMap] and returns a single [Query].
//
// Note: All [Query] are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *QueriesAPI) GetQueryByName(ctx context.Context, name string) (*Query, error) {
	result, err := a.ListQueriesAll(ctx, ListQueriesRequest{})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		if v.Name != name {
			continue
		}
		return &v, nil
	}
	return nil, fmt.Errorf("Query named '%s' does not exist", name)
}

// Restore a query
//
// Restore a query that has been moved to the trash. A restored query appears in
// list views and searches. You can use restored queries for alerts.
func (a *QueriesAPI) RestoreQuery(ctx context.Context, request RestoreQueryRequest) error {
	return a.impl.RestoreQuery(ctx, request)
}

// Change a query definition
//
// Modify this query definition.
//
// **Note**: You cannot undo this operation.
func (a *QueriesAPI) UpdateQuery(ctx context.Context, request QueryPostContent) (*Query, error) {
	return a.impl.UpdateQuery(ctx, request)
}
