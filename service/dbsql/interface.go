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

	// Creates an alert. An alert is a Databricks SQL object that periodically
	// runs a query, evaluates a condition of its result, and notifies users or
	// alert destinations if the condition was met.
	CreateAlert(ctx context.Context, request EditAlert) (*Alert, error)

	// **Note:** The structure of refresh schedules is subject to change.
	CreateSchedule(ctx context.Context, request CreateRefreshSchedule) (*RefreshSchedule, error)

	// Deletes an alert. Deleted alerts are no longer accessible and cannot be
	// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved
	// to the trash.
	DeleteAlert(ctx context.Context, request DeleteAlertRequest) error

	// DeleteAlertByAlertId calls DeleteAlert, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteAlertByAlertId(ctx context.Context, alertId string) error

	// Deletes alert refresh schedule that specifies when to refresh and
	// evaluate the associated query result.
	DeleteSchedule(ctx context.Context, request DeleteScheduleRequest) error

	// DeleteScheduleByAlertIdAndScheduleId calls DeleteSchedule, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteScheduleByAlertIdAndScheduleId(ctx context.Context, alertId string, scheduleId string) error

	// Gets an alert.
	GetAlert(ctx context.Context, request GetAlertRequest) (*Alert, error)

	// GetAlertByAlertId calls GetAlert, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetAlertByAlertId(ctx context.Context, alertId string) (*Alert, error)

	// An alert subscription represents exactly one recipient being notified
	// whenever the alert is triggered. The alert recipient is specified by
	// either the `user` field or the `destination` field. The `user` field is
	// ignored if `destination` is non-`null`.
	GetSubscriptions(ctx context.Context, request GetSubscriptionsRequest) ([]Subscription, error)

	// GetSubscriptionsByAlertId calls GetSubscriptions, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetSubscriptionsByAlertId(ctx context.Context, alertId string) ([]Subscription, error)

	// Get a list of alerts.
	ListAlerts(ctx context.Context) ([]Alert, error)

	// Alerts can have refresh schedules that specify when to refresh and
	// evaluate the associated query result. **Note:** Although refresh
	// schedules are returned in a list, only one refresh schedule per alert is
	// currently supported. The structure of refresh schedules is subject to
	// change.
	ListSchedules(ctx context.Context, request ListSchedulesRequest) ([]RefreshSchedule, error)

	// ListSchedulesByAlertId calls ListSchedules, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListSchedulesByAlertId(ctx context.Context, alertId string) ([]RefreshSchedule, error)

	Subscribe(ctx context.Context, request CreateSubscription) (*Subscription, error)

	// Unsubscribes a user or a destination to an alert.
	Unsubscribe(ctx context.Context, request UnsubscribeRequest) error

	// UnsubscribeByAlertIdAndSubscriptionId calls Unsubscribe, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	UnsubscribeByAlertIdAndSubscriptionId(ctx context.Context, alertId string, subscriptionId string) error

	// Updates an alert.
	UpdateAlert(ctx context.Context, request EditAlert) error
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
	CreateDashboard(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Moves a dashboard to the trash. Trashed dashboards do not appear in list
	// views or searches, and cannot be shared.
	DeleteDashboard(ctx context.Context, request DeleteDashboardRequest) error

	// DeleteDashboardByDashboardId calls DeleteDashboard, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteDashboardByDashboardId(ctx context.Context, dashboardId string) error

	// Returns a JSON representation of a dashboard object, including its
	// visualization and query objects.
	GetDashboard(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// GetDashboardByDashboardId calls GetDashboard, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetDashboardByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	ListDashboards(ctx context.Context, request ListDashboardsRequest) (*ListDashboardsResponse, error)

	// A restored dashboard appears in list views and searches and can be
	// shared.
	RestoreDashboard(ctx context.Context, request RestoreDashboardRequest) error
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

	// Retrieves a full list of SQL warehouses available in this workspace. All
	// fields that appear in this API response are enumerated for clarity.
	// However, you need only a SQL warehouse's `id` to create new queries
	// against it.
	ListDataSources(ctx context.Context) ([]DataSource, error)
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
	GetPermissions(ctx context.Context, request GetPermissionsRequest) (*GetPermissionsResponse, error)

	// GetPermissionsByObjectTypeAndObjectId calls GetPermissions, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetPermissionsByObjectTypeAndObjectId(ctx context.Context, objectType ObjectTypePlural, objectId string) (*GetPermissionsResponse, error)

	// Completely rewrite the access control list for a specified object.
	SetPermissions(ctx context.Context, request SetPermissionsRequest) (*SetPermissionsResponse, error)

	// Transfer ownership of a dashboard, query, or alert to an active user.
	// Requires an admin API key.
	TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error)
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, execution schedule, parameters, and visualizations.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type QueriesService interface {

	// Creates a new query definition. Queries created with this endpoint belong
	// to the authenticated user making the request. The `data_source_id` field
	// specifies the ID of the SQL warehouse to run this query against. You can
	// use the Data Sources API to see a complete list of available SQL
	// warehouses. Or you can copy the `data_source_id` from an existing query.
	// **Note**: You cannot add a visualization until you create the query.
	CreateQuery(ctx context.Context, request QueryPostContent) (*Query, error)

	// Moves a query to the trash. Trashed queries immediately disappear from
	// searches and list views, and they cannot be used for alerts. The trash is
	// deleted after 30 days.
	DeleteQuery(ctx context.Context, request DeleteQueryRequest) error

	// DeleteQueryByQueryId calls DeleteQuery, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteQueryByQueryId(ctx context.Context, queryId string) error

	// Retrieve a query object definition along with contextual permissions
	// information about the currently authenticated user.
	GetQuery(ctx context.Context, request GetQueryRequest) (*Query, error)

	// GetQueryByQueryId calls GetQuery, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetQueryByQueryId(ctx context.Context, queryId string) (*Query, error)

	// Retrieves a list of queries. Optionally, this list can be filtered by a
	// search term.
	ListQueries(ctx context.Context, request ListQueriesRequest) (*ListQueriesResponse, error)

	// A restored query appears in list views and searches. You can use restored
	// queries for alerts.
	RestoreQuery(ctx context.Context, request RestoreQueryRequest) error

	// Modify this query definition. **Note**: You cannot undo this operation.
	Update(ctx context.Context, request QueryPostContent) (*Query, error)
}
