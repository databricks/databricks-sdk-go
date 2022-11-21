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
		AlertsService: &alertsAPI{
			client: client,
		},
	}
}

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or alert destinations if
// the condition was met.
type AlertsAPI struct {
	// AlertsService contains low-level REST API interface.
	AlertsService
}

// Create an alert
//
// Creates an alert. An alert is a Databricks SQL object that periodically runs
// a query, evaluates a condition of its result, and notifies users or alert
// destinations if the condition was met.
func (a *AlertsAPI) CreateAlert(ctx context.Context, request EditAlert) (*Alert, error) {
	return a.AlertsService.CreateAlert(ctx, request)
}

// Create a refresh schedule
//
// Creates a new refresh schedule for an alert.
//
// **Note:** The structure of refresh schedules is subject to change.
func (a *AlertsAPI) CreateSchedule(ctx context.Context, request CreateRefreshSchedule) (*RefreshSchedule, error) {
	return a.AlertsService.CreateSchedule(ctx, request)
}

// Delete an alert
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved to
// the trash.
func (a *AlertsAPI) DeleteAlert(ctx context.Context, request DeleteAlertRequest) error {
	return a.AlertsService.DeleteAlert(ctx, request)
}

// Delete an alert
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved to
// the trash.
func (a *AlertsAPI) DeleteAlertByAlertId(ctx context.Context, alertId string) error {
	return a.DeleteAlert(ctx, DeleteAlertRequest{
		AlertId: alertId,
	})
}

// Delete a refresh schedule
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
func (a *AlertsAPI) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	return a.AlertsService.DeleteSchedule(ctx, request)
}

// Delete a refresh schedule
//
// Deletes an alert's refresh schedule. The refresh schedule specifies when to
// refresh and evaluate the associated query result.
func (a *AlertsAPI) DeleteScheduleByAlertIdAndScheduleId(ctx context.Context, alertId string, scheduleId string) error {
	return a.DeleteSchedule(ctx, DeleteScheduleRequest{
		AlertId:    alertId,
		ScheduleId: scheduleId,
	})
}

// Get an alert
//
// Gets an alert.
func (a *AlertsAPI) GetAlert(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	return a.AlertsService.GetAlert(ctx, request)
}

// Get an alert
//
// Gets an alert.
func (a *AlertsAPI) GetAlertByAlertId(ctx context.Context, alertId string) (*Alert, error) {
	return a.GetAlert(ctx, GetAlertRequest{
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
	return a.AlertsService.GetSubscriptions(ctx, request)
}

// Get an alert's subscriptions
//
// Get the subscriptions for an alert. An alert subscription represents exactly
// one recipient being notified whenever the alert is triggered. The alert
// recipient is specified by either the `user` field or the `destination` field.
// The `user` field is ignored if `destination` is non-`null`.
func (a *AlertsAPI) GetSubscriptionsByAlertId(ctx context.Context, alertId string) ([]Subscription, error) {
	return a.GetSubscriptions(ctx, GetSubscriptionsRequest{
		AlertId: alertId,
	})
}

// Get alerts
//
// Gets a list of alerts.
func (a *AlertsAPI) ListAlerts(ctx context.Context) ([]Alert, error) {
	return a.AlertsService.ListAlerts(ctx)
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
	return a.AlertsService.ListSchedules(ctx, request)
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
	return a.ListSchedules(ctx, ListSchedulesRequest{
		AlertId: alertId,
	})
}

// Subscribe to an alert
func (a *AlertsAPI) Subscribe(ctx context.Context, request CreateSubscription) (*Subscription, error) {
	return a.AlertsService.Subscribe(ctx, request)
}

// Unsubscribe to an alert
//
// Unsubscribes a user or a destination to an alert.
func (a *AlertsAPI) Unsubscribe(ctx context.Context, request UnsubscribeRequest) error {
	return a.AlertsService.Unsubscribe(ctx, request)
}

// Unsubscribe to an alert
//
// Unsubscribes a user or a destination to an alert.
func (a *AlertsAPI) UnsubscribeByAlertIdAndSubscriptionId(ctx context.Context, alertId string, subscriptionId string) error {
	return a.Unsubscribe(ctx, UnsubscribeRequest{
		AlertId:        alertId,
		SubscriptionId: subscriptionId,
	})
}

// Update an alert
//
// Updates an alert.
func (a *AlertsAPI) UpdateAlert(ctx context.Context, request EditAlert) error {
	return a.AlertsService.UpdateAlert(ctx, request)
}

// unexported type that holds implementations of just Alerts API methods
type alertsAPI struct {
	client *client.DatabricksClient
}

func (a *alertsAPI) CreateAlert(ctx context.Context, request EditAlert) (*Alert, error) {
	var alert Alert
	path := "/api/2.0/preview/sql/alerts"
	err := a.client.Post(ctx, path, request, &alert)
	return &alert, err
}

func (a *alertsAPI) CreateSchedule(ctx context.Context, request CreateRefreshSchedule) (*RefreshSchedule, error) {
	var refreshSchedule RefreshSchedule
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/refresh-schedules", request.AlertId)
	err := a.client.Post(ctx, path, request, &refreshSchedule)
	return &refreshSchedule, err
}

func (a *alertsAPI) DeleteAlert(ctx context.Context, request DeleteAlertRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *alertsAPI) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/refresh-schedules/%v", request.AlertId, request.ScheduleId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *alertsAPI) GetAlert(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Get(ctx, path, request, &alert)
	return &alert, err
}

func (a *alertsAPI) GetSubscriptions(ctx context.Context, request GetSubscriptionsRequest) ([]Subscription, error) {
	var subscriptionList []Subscription
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/subscriptions", request.AlertId)
	err := a.client.Get(ctx, path, request, &subscriptionList)
	return subscriptionList, err
}

func (a *alertsAPI) ListAlerts(ctx context.Context) ([]Alert, error) {
	var alertList []Alert
	path := "/api/2.0/preview/sql/alerts"
	err := a.client.Get(ctx, path, nil, &alertList)
	return alertList, err
}

func (a *alertsAPI) ListSchedules(ctx context.Context, request ListSchedulesRequest) ([]RefreshSchedule, error) {
	var refreshScheduleList []RefreshSchedule
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/refresh-schedules", request.AlertId)
	err := a.client.Get(ctx, path, request, &refreshScheduleList)
	return refreshScheduleList, err
}

func (a *alertsAPI) Subscribe(ctx context.Context, request CreateSubscription) (*Subscription, error) {
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/subscriptions", request.AlertId)
	err := a.client.Post(ctx, path, request, &subscription)
	return &subscription, err
}

func (a *alertsAPI) Unsubscribe(ctx context.Context, request UnsubscribeRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/subscriptions/%v", request.AlertId, request.SubscriptionId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *alertsAPI) UpdateAlert(ctx context.Context, request EditAlert) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewDashboards(client *client.DatabricksClient) *DashboardsAPI {
	return &DashboardsAPI{
		DashboardsService: &dashboardsAPI{
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
	// DashboardsService contains low-level REST API interface.
	DashboardsService
}

// Create a dashboard object
func (a *DashboardsAPI) CreateDashboard(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	return a.DashboardsService.CreateDashboard(ctx, request)
}

// Remove a dashboard
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteDashboard(ctx context.Context, request DeleteDashboardRequest) error {
	return a.DashboardsService.DeleteDashboard(ctx, request)
}

// Remove a dashboard
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteDashboardByDashboardId(ctx context.Context, dashboardId string) error {
	return a.DeleteDashboard(ctx, DeleteDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Retrieve a definition
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetDashboard(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	return a.DashboardsService.GetDashboard(ctx, request)
}

// Retrieve a definition
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetDashboardByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.GetDashboard(ctx, GetDashboardRequest{
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
		response, err := a.ListDashboards(ctx, request)
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

// Restore a dashboard
//
// A restored dashboard appears in list views and searches and can be shared.
func (a *DashboardsAPI) RestoreDashboard(ctx context.Context, request RestoreDashboardRequest) error {
	return a.DashboardsService.RestoreDashboard(ctx, request)
}

// unexported type that holds implementations of just Dashboards API methods
type dashboardsAPI struct {
	client *client.DatabricksClient
}

func (a *dashboardsAPI) CreateDashboard(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/preview/sql/dashboards"
	err := a.client.Post(ctx, path, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsAPI) DeleteDashboard(ctx context.Context, request DeleteDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *dashboardsAPI) GetDashboard(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	err := a.client.Get(ctx, path, request, &dashboard)
	return &dashboard, err
}

func (a *dashboardsAPI) ListDashboards(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error) {
	var listDashboardsResponse ListDashboardsResponse
	path := "/api/2.0/preview/sql/dashboards"
	err := a.client.Get(ctx, path, request, &listDashboardsResponse)
	return &listDashboardsResponse, err
}

func (a *dashboardsAPI) RestoreDashboard(ctx context.Context, request RestoreDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func NewDataSources(client *client.DatabricksClient) *DataSourcesAPI {
	return &DataSourcesAPI{
		DataSourcesService: &dataSourcesAPI{
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
	// DataSourcesService contains low-level REST API interface.
	DataSourcesService
}

// Get a list of SQL warehouses
//
// Retrieves a full list of SQL warehouses available in this workspace. All
// fields that appear in this API response are enumerated for clarity. However,
// you need only a SQL warehouse's `id` to create new queries against it.
func (a *DataSourcesAPI) ListDataSources(ctx context.Context) ([]DataSource, error) {
	return a.DataSourcesService.ListDataSources(ctx)
}

// unexported type that holds implementations of just DataSources API methods
type dataSourcesAPI struct {
	client *client.DatabricksClient
}

func (a *dataSourcesAPI) ListDataSources(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0/preview/sql/data_sources"
	err := a.client.Get(ctx, path, nil, &dataSourceList)
	return dataSourceList, err
}

func NewDbsqlPermissions(client *client.DatabricksClient) *DbsqlPermissionsAPI {
	return &DbsqlPermissionsAPI{
		DbsqlPermissionsService: &dbsqlPermissionsAPI{
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
	// DbsqlPermissionsService contains low-level REST API interface.
	DbsqlPermissionsService
}

// Get object ACL
//
// Gets a JSON representation of the access control list (ACL) for a specified
// object.
func (a *DbsqlPermissionsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	return a.DbsqlPermissionsService.GetPermissions(ctx, request)
}

// Get object ACL
//
// Gets a JSON representation of the access control list (ACL) for a specified
// object.
func (a *DbsqlPermissionsAPI) GetPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType ObjectTypePlural, objectId string) (*GetPermissionsResponse, error) {
	return a.GetPermissions(ctx, GetPermissionsRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
}

// Set object ACL
//
// Sets the access control list (ACL) for a specified object. This operation
// will complete rewrite the ACL.
func (a *DbsqlPermissionsAPI) SetPermissions(ctx context.Context, request SetPermissionsRequest) (*SetPermissionsResponse, error) {
	return a.DbsqlPermissionsService.SetPermissions(ctx, request)
}

// Transfer object ownership
//
// Transfers ownership of a dashboard, query, or alert to an active user.
// Requires an admin API key.
func (a *DbsqlPermissionsAPI) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	return a.DbsqlPermissionsService.TransferOwnership(ctx, request)
}

// unexported type that holds implementations of just DbsqlPermissions API methods
type dbsqlPermissionsAPI struct {
	client *client.DatabricksClient
}

func (a *dbsqlPermissionsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Get(ctx, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

func (a *dbsqlPermissionsAPI) SetPermissions(ctx context.Context, request SetPermissionsRequest) (*SetPermissionsResponse, error) {
	var setPermissionsResponse SetPermissionsResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Post(ctx, path, request, &setPermissionsResponse)
	return &setPermissionsResponse, err
}

func (a *dbsqlPermissionsAPI) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	err := a.client.Post(ctx, path, request, &success)
	return &success, err
}

func NewQueries(client *client.DatabricksClient) *QueriesAPI {
	return &QueriesAPI{
		QueriesService: &queriesAPI{
			client: client,
		},
	}
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, execution schedule, parameters, and visualizations.
type QueriesAPI struct {
	// QueriesService contains low-level REST API interface.
	QueriesService
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
	return a.QueriesService.CreateQuery(ctx, request)
}

// Delete a query
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteQuery(ctx context.Context, request DeleteQueryRequest) error {
	return a.QueriesService.DeleteQuery(ctx, request)
}

// Delete a query
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteQueryByQueryId(ctx context.Context, queryId string) error {
	return a.DeleteQuery(ctx, DeleteQueryRequest{
		QueryId: queryId,
	})
}

// Get a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetQuery(ctx context.Context, request GetQueryRequest) (*Query, error) {
	return a.QueriesService.GetQuery(ctx, request)
}

// Get a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetQueryByQueryId(ctx context.Context, queryId string) (*Query, error) {
	return a.GetQuery(ctx, GetQueryRequest{
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
		response, err := a.ListQueries(ctx, request)
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

// Restore a query
//
// Restore a query that has been moved to the trash. A restored query appears in
// list views and searches. You can use restored queries for alerts.
func (a *QueriesAPI) RestoreQuery(ctx context.Context, request RestoreQueryRequest) error {
	return a.QueriesService.RestoreQuery(ctx, request)
}

// Change a query definition
//
// Modify this query definition.
//
// **Note**: You cannot undo this operation.
func (a *QueriesAPI) UpdateQuery(ctx context.Context, request QueryPostContent) (*Query, error) {
	return a.QueriesService.UpdateQuery(ctx, request)
}

// unexported type that holds implementations of just Queries API methods
type queriesAPI struct {
	client *client.DatabricksClient
}

func (a *queriesAPI) CreateQuery(ctx context.Context, request QueryPostContent) (*Query, error) {
	var query Query
	path := "/api/2.0/preview/sql/queries"
	err := a.client.Post(ctx, path, request, &query)
	return &query, err
}

func (a *queriesAPI) DeleteQuery(ctx context.Context, request DeleteQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *queriesAPI) GetQuery(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Get(ctx, path, request, &query)
	return &query, err
}

func (a *queriesAPI) ListQueries(ctx context.Context, request ListQueriesRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/preview/sql/queries"
	err := a.client.Get(ctx, path, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

func (a *queriesAPI) RestoreQuery(ctx context.Context, request RestoreQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func (a *queriesAPI) UpdateQuery(ctx context.Context, request QueryPostContent) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Post(ctx, path, request, &query)
	return &query, err
}
