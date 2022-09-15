// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbsql

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func NewAlerts(client *client.DatabricksClient) AlertsService {
	return &AlertsAPI{
		client: client,
	}
}

type AlertsAPI struct {
	client *client.DatabricksClient
}

// Create an alert
//
// Creates an alert. An alert is a Databricks SQL object that periodically runs
// a query, evaluates a condition of its result, and notifies users or alert
// destinations if the condition was met.
func (a *AlertsAPI) CreateAlert(ctx context.Context, request EditAlert) (*Alert, error) {
	var alert Alert
	path := "/api/2.0/preview/sql/alerts"
	err := a.client.Post(ctx, path, request, &alert)
	return &alert, err
}

// Create an alert's refresh schedule
//
// **Note:** The structure of refresh schedules is subject to change.
func (a *AlertsAPI) CreateSchedule(ctx context.Context, request CreateRefreshSchedule) (*RefreshSchedule, error) {
	var refreshSchedule RefreshSchedule
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/refresh-schedules", request.AlertId)
	err := a.client.Post(ctx, path, request, &refreshSchedule)
	return &refreshSchedule, err
}

// Delete an alert
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored.
//
// **Note:** Unlike queries and dashboards, alerts cannot be moved to the trash.
func (a *AlertsAPI) DeleteAlert(ctx context.Context, request DeleteAlertRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete an alert
//
// Deletes an alert. Deleted alerts are no longer accessible and cannot be
// restored.
//
// **Note:** Unlike queries and dashboards, alerts cannot be moved to the trash.
func (a *AlertsAPI) DeleteAlertByAlertId(ctx context.Context, alertId string) error {
	return a.DeleteAlert(ctx, DeleteAlertRequest{
		AlertId: alertId,
	})
}

// Delete an alert's refresh schedule
//
// Deletes alert refresh schedule that specifies when to refresh and evaluate
// the associated query result.
func (a *AlertsAPI) DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/refresh-schedules/%v", request.AlertId, request.ScheduleId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Delete an alert's refresh schedule
//
// Deletes alert refresh schedule that specifies when to refresh and evaluate
// the associated query result.
func (a *AlertsAPI) DeleteScheduleByAlertIdAndScheduleId(ctx context.Context, alertId string, scheduleId string) error {
	return a.DeleteSchedule(ctx, DeleteScheduleRequest{
		AlertId:    alertId,
		ScheduleId: scheduleId,
	})
}

// Fetch an alert
//
// Gets an alert.
func (a *AlertsAPI) GetAlert(ctx context.Context, request GetAlertRequest) (*Alert, error) {
	var alert Alert
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Get(ctx, path, request, &alert)
	return &alert, err
}

// Fetch an alert
//
// Gets an alert.
func (a *AlertsAPI) GetAlertByAlertId(ctx context.Context, alertId string) (*Alert, error) {
	return a.GetAlert(ctx, GetAlertRequest{
		AlertId: alertId,
	})
}

// Fetch an alert's subscriptions
//
// An alert subscription represents exactly one recipient being notified
// whenever the alert is triggered. The alert recipient is specified by either
// the `user` field or the `destination` field. The `user` field is ignored if
// `destination` is non-`null`.
func (a *AlertsAPI) GetSubscriptions(ctx context.Context, request GetSubscriptionsRequest) ([]Subscription, error) {
	var subscriptionList []Subscription
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/subscriptions", request.AlertId)
	err := a.client.Get(ctx, path, request, &subscriptionList)
	return subscriptionList, err
}

// Fetch an alert's subscriptions
//
// An alert subscription represents exactly one recipient being notified
// whenever the alert is triggered. The alert recipient is specified by either
// the `user` field or the `destination` field. The `user` field is ignored if
// `destination` is non-`null`.
func (a *AlertsAPI) GetSubscriptionsByAlertId(ctx context.Context, alertId string) ([]Subscription, error) {
	return a.GetSubscriptions(ctx, GetSubscriptionsRequest{
		AlertId: alertId,
	})
}

// Fetch a list of alerts.
//
// Get a list of alerts.
func (a *AlertsAPI) ListAlerts(ctx context.Context) ([]Alert, error) {
	var alertList []Alert
	path := "/api/2.0/preview/sql/alerts"
	err := a.client.Get(ctx, path, nil, &alertList)
	return alertList, err
}

<<<<<<< HEAD
// Fetch an alert's refresh schedules
//
=======
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

>>>>>>> 499b182 (regenerate)
// Alerts can have refresh schedules that specify when to refresh and evaluate
// the associated query result.
//
// **Note:** Although refresh schedules are returned in a list, only one refresh
// schedule per alert is currently supported. The structure of refresh schedules
// is subject to change.
func (a *AlertsAPI) ListSchedules(ctx context.Context, request ListSchedulesRequest) ([]RefreshSchedule, error) {
	var refreshScheduleList []RefreshSchedule
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/refresh-schedules", request.AlertId)
	err := a.client.Get(ctx, path, request, &refreshScheduleList)
	return refreshScheduleList, err
}

// Fetch an alert's refresh schedules
//
// Alerts can have refresh schedules that specify when to refresh and evaluate
// the associated query result.
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
	var subscription Subscription
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/subscriptions", request.AlertId)
	err := a.client.Post(ctx, path, request, &subscription)
	return &subscription, err
}

// Unsubscribe to an alert
//
// Unsubscribes a user or a destination to an alert.
func (a *AlertsAPI) Unsubscribe(ctx context.Context, request UnsubscribeRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v/subscriptions/%v", request.AlertId, request.SubscriptionId)
	err := a.client.Delete(ctx, path, request)
	return err
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

// Edit an alert
//
// Updates an alert.
func (a *AlertsAPI) UpdateAlert(ctx context.Context, request EditAlert) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/alerts/%v", request.AlertId)
	err := a.client.Put(ctx, path, request)
	return err
}

func NewDashboards(client *client.DatabricksClient) DashboardsService {
	return &DashboardsAPI{
		client: client,
	}
}

type DashboardsAPI struct {
	client *client.DatabricksClient
}

// Create a new dashboard object
func (a *DashboardsAPI) CreateDashboard(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := "/api/2.0/preview/sql/dashboards"
	err := a.client.Post(ctx, path, request, &dashboard)
	return &dashboard, err
}

// Move a dashboard to the trash
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteDashboard(ctx context.Context, request DeleteDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Move a dashboard to the trash
//
// Moves a dashboard to the trash. Trashed dashboards do not appear in list
// views or searches, and cannot be shared.
func (a *DashboardsAPI) DeleteDashboardByDashboardId(ctx context.Context, dashboardId string) error {
	return a.DeleteDashboard(ctx, DeleteDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Retrieve a dashboard definition
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetDashboard(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	var dashboard Dashboard
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/%v", request.DashboardId)
	err := a.client.Get(ctx, path, request, &dashboard)
	return &dashboard, err
}

// Retrieve a dashboard definition
//
// Returns a JSON representation of a dashboard object, including its
// visualization and query objects.
func (a *DashboardsAPI) GetDashboardByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.GetDashboard(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Fetch a paginated list of dashboard objects
func (a *DashboardsAPI) ListDashboards(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error) {
	var listDashboardsResponse ListDashboardsResponse
	path := "/api/2.0/preview/sql/dashboards"
	err := a.client.Get(ctx, path, request, &listDashboardsResponse)
	return &listDashboardsResponse, err
}

<<<<<<< HEAD
// Restore a trashed dashboard
//
=======
func (a *DashboardsAPI) ListDashboardsAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error) {
	var results []Dashboard
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
		request.Page += 1
	}
	return results, nil
}

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

>>>>>>> 499b182 (regenerate)
// A restored dashboard appears in list views and searches and can be shared.
func (a *DashboardsAPI) RestoreDashboard(ctx context.Context, request RestoreDashboardRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/dashboards/trash/%v", request.DashboardId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

func NewDataSources(client *client.DatabricksClient) DataSourcesService {
	return &DataSourcesAPI{
		client: client,
	}
}

type DataSourcesAPI struct {
	client *client.DatabricksClient
}

// Fetch a list of SQL warehouses
//
// Retrieves a full list of SQL warehouses available in this workspace. All
// fields that appear in this API response are enumerated for clarity. However,
// you need only a SQL warehouse's `id` to create new queries against it.
func (a *DataSourcesAPI) ListDataSources(ctx context.Context) ([]DataSource, error) {
	var dataSourceList []DataSource
	path := "/api/2.0/preview/sql/data_sources"
	err := a.client.Get(ctx, path, nil, &dataSourceList)
	return dataSourceList, err
}

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

func NewDbsqlPermissions(client *client.DatabricksClient) DbsqlPermissionsService {
	return &DbsqlPermissionsAPI{
		client: client,
	}
}

type DbsqlPermissionsAPI struct {
	client *client.DatabricksClient
}

// Get object access control list
//
// Returns a JSON representation of the access control list (ACL) for a
// specified object.
func (a *DbsqlPermissionsAPI) GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error) {
	var getPermissionsResponse GetPermissionsResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Get(ctx, path, request, &getPermissionsResponse)
	return &getPermissionsResponse, err
}

// Get object access control list
//
// Returns a JSON representation of the access control list (ACL) for a
// specified object.
func (a *DbsqlPermissionsAPI) GetPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType ObjectTypePlural, objectId string) (*GetPermissionsResponse, error) {
	return a.GetPermissions(ctx, GetPermissionsRequest{
		ObjectType: objectType,
		ObjectId:   objectId,
	})
}

// Set object access control list
//
// Completely rewrite the access control list for a specified object.
func (a *DbsqlPermissionsAPI) SetPermissions(ctx context.Context, request SetPermissionsRequest) (*SetPermissionsResponse, error) {
	var setPermissionsResponse SetPermissionsResponse
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v", request.ObjectType, request.ObjectId)
	err := a.client.Post(ctx, path, request, &setPermissionsResponse)
	return &setPermissionsResponse, err
}

// Transfer object ownership
//
// Transfer ownership of a dashboard, query, or alert to an active user.
// Requires an admin API key.
func (a *DbsqlPermissionsAPI) TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error) {
	var success Success
	path := fmt.Sprintf("/api/2.0/preview/sql/permissions/%v/%v/transfer", request.ObjectType, request.ObjectId)
	err := a.client.Post(ctx, path, request, &success)
	return &success, err
}

func NewQueries(client *client.DatabricksClient) QueriesService {
	return &QueriesAPI{
		client: client,
	}
}

type QueriesAPI struct {
	client *client.DatabricksClient
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
	var query Query
	path := "/api/2.0/preview/sql/queries"
	err := a.client.Post(ctx, path, request, &query)
	return &query, err
}

// Move a query to the trash
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteQuery(ctx context.Context, request DeleteQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Delete(ctx, path, request)
	return err
}

// Move a query to the trash
//
// Moves a query to the trash. Trashed queries immediately disappear from
// searches and list views, and they cannot be used for alerts. The trash is
// deleted after 30 days.
func (a *QueriesAPI) DeleteQueryByQueryId(ctx context.Context, queryId string) error {
	return a.DeleteQuery(ctx, DeleteQueryRequest{
		QueryId: queryId,
	})
}

// Retrieve a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetQuery(ctx context.Context, request GetQueryRequest) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Get(ctx, path, request, &query)
	return &query, err
}

// Retrieve a query definition.
//
// Retrieve a query object definition along with contextual permissions
// information about the currently authenticated user.
func (a *QueriesAPI) GetQueryByQueryId(ctx context.Context, queryId string) (*Query, error) {
	return a.GetQuery(ctx, GetQueryRequest{
		QueryId: queryId,
	})
}

// Retrieve list of queries
//
// Retrieves a list of queries. Optionally, this list can be filtered by a
// search term.
func (a *QueriesAPI) ListQueries(ctx context.Context, request ListQueriesRequest) (*ListQueriesResponse, error) {
	var listQueriesResponse ListQueriesResponse
	path := "/api/2.0/preview/sql/queries"
	err := a.client.Get(ctx, path, request, &listQueriesResponse)
	return &listQueriesResponse, err
}

<<<<<<< HEAD
// Restore a trashed query
//
=======
func (a *QueriesAPI) ListQueriesAll(ctx context.Context, request ListQueriesRequest) ([]Query, error) {
	var results []Query
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
		request.Page += 1
	}
	return results, nil
}

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

>>>>>>> 499b182 (regenerate)
// A restored query appears in list views and searches. You can use restored
// queries for alerts.
func (a *QueriesAPI) RestoreQuery(ctx context.Context, request RestoreQueryRequest) error {
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/trash/%v", request.QueryId)
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Change a query definition
//
// Modify this query definition.
//
// **Note**: You cannot undo this operation.
func (a *QueriesAPI) Update(ctx context.Context, request QueryPostContent) (*Query, error) {
	var query Query
	path := fmt.Sprintf("/api/2.0/preview/sql/queries/%v", request.QueryId)
	err := a.client.Post(ctx, path, request, &query)
	return &query, err
}
