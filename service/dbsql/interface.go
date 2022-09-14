// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbsql

import (
	"context"
)

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or alert destinations if
// the condition was met.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type AlertsService interface {
	CreateAlert(ctx context.Context, editAlert EditAlert) (*Alert, error)

	// **Note:** The structure of refresh schedules is subject to change.
	CreateSchedule(ctx context.Context, createRefreshSchedule CreateRefreshSchedule) (*RefreshSchedule, error)

	// Deleted alerts are no longer accessible. Once deleted, alerts cannot be
	// restored. **Note:** Unlike queries and dashboards, alerts do not support
	// the trash functionality.
	DeleteAlert(ctx context.Context, deleteAlertRequest DeleteAlertRequest) error

	DeleteAlertByAlertId(ctx context.Context, alertId string) error

	DeleteSchedule(ctx context.Context, deleteScheduleRequest DeleteScheduleRequest) error

	DeleteScheduleByAlertIdAndScheduleId(ctx context.Context, alertId string, scheduleId string) error

	GetAlert(ctx context.Context, getAlertRequest GetAlertRequest) (*Alert, error)

	GetAlertByAlertId(ctx context.Context, alertId string) (*Alert, error)

	// An alert subscription represents exactly one recipient being notified
	// whenever the alert is triggered. The alert recipient is specified by
	// either the `user` field or the `destination` field. The `user` field is
	// ignored if `destination` is non-`null`.
	GetSubscriptions(ctx context.Context, getSubscriptionsRequest GetSubscriptionsRequest) ([]Subscription, error)

	GetSubscriptionsByAlertId(ctx context.Context, alertId string) ([]Subscription, error)

	ListAlerts(ctx context.Context) ([]Alert, error)

	AlertNameToIdMap(ctx context.Context) (map[string]string, error)

	GetAlertByName(ctx context.Context, name string) (*Alert, error)
	// Alerts can have refresh schedules that specify when to refresh and
	// evaluate the associated query result. **Note:** Although refresh
	// schedules are returned in a list, only one refresh schedule per alert is
	// currently supported. The structure of refresh schedules is subject to
	// change.
	ListSchedules(ctx context.Context, listSchedulesRequest ListSchedulesRequest) ([]RefreshSchedule, error)

	ListSchedulesByAlertId(ctx context.Context, alertId string) ([]RefreshSchedule, error)

	Subscribe(ctx context.Context, createSubscription CreateSubscription) (*Subscription, error)

	Unsubscribe(ctx context.Context, unsubscribeRequest UnsubscribeRequest) error

	UnsubscribeByAlertIdAndSubscriptionId(ctx context.Context, alertId string, subscriptionId string) error

	UpdateAlert(ctx context.Context, editAlert EditAlert) error
}

// In general, there is little need to modify dashboards using the API. However,
// it can be useful to use dashboard objects to look-up a collection of related
// query IDs. The API can also be used to duplicate multiple dashboards at once
// since you can get a dashboard definition with a GET request and then POST it
// to create a new one.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type DashboardsService interface {
	CreateDashboard(ctx context.Context, createDashboardRequest CreateDashboardRequest) (*Dashboard, error)

	// Trashed dashboards do not appear in list views or searches and cannot be
	// shared.
	DeleteDashboard(ctx context.Context, deleteDashboardRequest DeleteDashboardRequest) error

	DeleteDashboardByDashboardId(ctx context.Context, dashboardId string) error

	// Returns a JSON representation of a dashboard object, including its
	// visualization and query objects.
	GetDashboard(ctx context.Context, getDashboardRequest GetDashboardRequest) (*Dashboard, error)

	GetDashboardByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	ListDashboards(ctx context.Context, listDashboardsRequest ListDashboardsRequest) (*ListDashboardsResponse, error)
	// ListDashboardsAll retrieves all available results from the platform
	ListDashboardsAll(ctx context.Context, request ListDashboardsRequest) ([]Dashboard, error)

	DashboardNameToIdMap(ctx context.Context, request ListDashboardsRequest) (map[string]string, error)

	GetDashboardByName(ctx context.Context, name string) (*Dashboard, error)
	// A restored dashboard appears in list views and searches and can be
	// shared.
	RestoreDashboard(ctx context.Context, restoreDashboardRequest RestoreDashboardRequest) error
}

// This API is provided to assist you in making new query objects. When creating
// a query object, you may optionally specify a `data_source_id` for the SQL
// warehouse against which it will run. If you don't already know the
// `data_source_id` for your desired SQL warehouse, this API will help you find
// it. This API does not support searches. It returns the full list of SQL
// warehouses in your workspace. We advise you to use any text editor, REST
// client, or `grep` to search the response from this API for the name of your
// SQL warehouse as it appears in Databricks SQL.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type DataSourcesService interface {
	// Retrieve a full list of SQL warehouses available in this workspace. All
	// fields that appear in this API response are enumerated for clarity.
	// However, you will only need a SQL warehouse's `id` to create new queries
	// against it.
	ListDataSources(ctx context.Context) ([]DataSource, error)

	DataSourceNameToIdMap(ctx context.Context) (map[string]string, error)
}

// The SQL Permissions API is similar to the endpoints of the [Permissions
// API](%(permissionsApiLink)s). However, this exposes only one endpoint, which
// gets the Access Control List for a given object. You cannot modify any
// permissions using this API. There are three levels of permission: -
// `CAN_VIEW`: Allows read-only access - `CAN_RUN`: Allows read access and run
// access (superset of `CAN_VIEW`) - `CAN_MANAGE`: Allows all actions: read,
// run, edit, delete, modify permissions (superset of `CAN_RUN`)
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type DbsqlPermissionsService interface {
	// Returns a JSON representation of the access control list (ACL) for a
	// specified object.
	GetPermissions(ctx context.Context, getPermissionsRequest GetPermissionsRequest) (*GetPermissionsResponse, error)

	GetPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType ObjectTypePlural, objectId string) (*GetPermissionsResponse, error)

	// Completely rewrite the access control list for a specified object.
	SetPermissions(ctx context.Context, setPermissionsRequest SetPermissionsRequest) (*SetPermissionsResponse, error)

	// Transfer ownership of a dashboard, query, or alert to an active user.
	// Requires an admin API key.
	TransferOwnership(ctx context.Context, transferOwnershipRequest TransferOwnershipRequest) (*Success, error)
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, execution schedule, parameters, and visualizations.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type QueriesService interface {
	// Queries created with this endpoint belong to the authenticated user
	// making the request. The `data_source_id` field specifies the id of the
	// SQL warehouse against which this query will run. You can use the Data
	// Sources API to see a complete list of available SQL warehouses. Or you
	// can copy the `data_source_id` from an existing query. **Note**: You
	// cannot add a visualization until you create the query.
	CreateQuery(ctx context.Context, queryPostContent QueryPostContent) (*Query, error)

	// Trashed queries immediately disappear from searches and list views and
	// cannot be used for alerts. The trash is deleted after 30 days.
	DeleteQuery(ctx context.Context, deleteQueryRequest DeleteQueryRequest) error

	DeleteQueryByQueryId(ctx context.Context, queryId string) error

	// Retrieve a query object definition along with contextual permissions
	// information about the currently authenticated user.
	GetQuery(ctx context.Context, getQueryRequest GetQueryRequest) (*Query, error)

	GetQueryByQueryId(ctx context.Context, queryId string) (*Query, error)

	// Optionally this list can be filtered by a search term.
	ListQueries(ctx context.Context, listQueriesRequest ListQueriesRequest) (*ListQueriesResponse, error)
	// ListQueriesAll retrieves all available results from the platform
	ListQueriesAll(ctx context.Context, request ListQueriesRequest) ([]Query, error)

	QueryNameToIdMap(ctx context.Context, request ListQueriesRequest) (map[string]string, error)

	GetQueryByName(ctx context.Context, name string) (*Query, error)
	// A restored query appears in list views and searches. You can use restored
	// queries for alerts.
	RestoreQuery(ctx context.Context, restoreQueryRequest RestoreQueryRequest) error

	// Modify this query definition. **Note**: You cannot undo this operation.
	Update(ctx context.Context, queryPostContent QueryPostContent) (*Query, error)
}
