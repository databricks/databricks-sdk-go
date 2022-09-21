// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbsql

// all definitions in this file are in alphabetical order

type AccessControl struct {
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

type Alert struct {
	CreatedAt string `json:"created_at,omitempty"`
	// ID of the alert.
	Id string `json:"id,omitempty"`

	LastTriggeredAt string `json:"last_triggered_at,omitempty"`
	// Name of the alert.
	Name string `json:"name,omitempty"`

	Options *AlertOptions `json:"options,omitempty"`

	Query *Query `json:"query,omitempty"`

	Rearm int `json:"rearm,omitempty"`

	State AlertState `json:"state,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	User *User `json:"user,omitempty"`
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	Column string `json:"column"`
	// Custom body of alert notification, if it exists. See `%(alertsLink)s` for
	// custom templating instructions.
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See `%(alertsLink)s` for custom
	// templating instructions.
	CustomSubject string `json:"custom_subject,omitempty"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and alert destinations when triggered.
	Muted bool `json:"muted,omitempty"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op string `json:"op"`
	// Number of failures encountered during alert refresh. This counter is used
	// for sending aggregated alert failure email notifications.
	ScheduleFailures int `json:"schedule_failures,omitempty"`
	// Value used to compare in alert evaluation.
	Value string `json:"value"`
}

type CreateRefreshSchedule struct {
	AlertId string `json:"-" path:"alert_id"`

	Cron string `json:"cron"`

	DataSourceId string `json:"data_source_id,omitempty"`
}

type CreateSubscription struct {
	AlertId string `json:"alert_id" path:"alert_id"`

	DestinationId string `json:"destination_id,omitempty"`

	UserId int `json:"user_id,omitempty"`
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit bool `json:"can_edit,omitempty"`
	// Timestamp when this dashboard was created.
	CreatedAt string `json:"created_at,omitempty"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// The ID for this dashboard.
	Id string `json:"id,omitempty"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft bool `json:"is_draft,omitempty"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name,omitempty"`

	Options *DashboardOptions `json:"options,omitempty"`

	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug string `json:"slug,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *User `json:"user,omitempty"`
	// The ID of the user that created and owns this dashboard.
	UserId int `json:"user_id,omitempty"`
	// Unused field.
	Version string `json:"version,omitempty"`

	Widgets []Widget `json:"widgets,omitempty"`
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// The unique identifier for this data source / SQL warehouse. Can be used
	// when creating / modifying queries and dashboards.
	Id string `json:"id,omitempty"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name string `json:"name,omitempty"`
	// <needs content>
	PauseReason string `json:"pause_reason,omitempty"`
	// <needs content>
	Paused any/* MISSING TYPE */ `json:"paused,omitempty"`
	// <needs content>
	SupportsAutoLimit bool `json:"supports_auto_limit,omitempty"`
	// <needs content>
	Syntax string `json:"syntax,omitempty"`
	// <needs content>
	Type string `json:"type,omitempty"`
	// <needs content>
	ViewOnly bool `json:"view_only,omitempty"`
	// <needs content>
	WarehouseId string `json:"warehouse_id,omitempty"`
}

// Alert destination subscribed to the alert, if it exists. Alert destinations
// can be configured by admins through the UI. See `%(alertDestinationsLink)s`.
type Destination struct {
	// ID of the alert destination.
	Id string `json:"id,omitempty"`
	// Name of the alert destination.
	Name string `json:"name,omitempty"`
	// Type of the alert destination.
	Type DestinationType `json:"type,omitempty"`
}

// Type of the alert destination.
type DestinationType string

const DestinationTypeEmail DestinationType = `email`

const DestinationTypeHangoutsChat DestinationType = `hangouts_chat`

const DestinationTypeMattermost DestinationType = `mattermost`

const DestinationTypeMicrosoftTeams DestinationType = `microsoft_teams`

const DestinationTypePagerduty DestinationType = `pagerduty`

const DestinationTypeSlack DestinationType = `slack`

const DestinationTypeWebhook DestinationType = `webhook`

type EditAlert struct {
	AlertId string `json:"-" path:"alert_id"`

	Name string `json:"name"`

	Options AlertOptions `json:"options"`

	QueryId string `json:"query_id"`

	Rearm int `json:"rearm,omitempty"`
}

type Parameter struct {
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name string `json:"name,omitempty"`
	// The text displayed in a parameter picking widget.
	Title string `json:"title,omitempty"`
	// Parameters can have several different types.
	Type ParameterType `json:"type,omitempty"`
	// The default value for this parameter.
	Value string `json:"value,omitempty"`
}

// Parameters can have several different types.
type ParameterType string

const ParameterTypeDatetime ParameterType = `datetime`

const ParameterTypeNumber ParameterType = `number`

const ParameterTypeText ParameterType = `text`

// This describes an enum
type PermissionLevel string

// Can manage the query
const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

// Can run the query
const PermissionLevelCanRun PermissionLevel = `CAN_RUN`

// Can view the query
const PermissionLevelCanView PermissionLevel = `CAN_VIEW`

type Query struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit bool `json:"can_edit,omitempty"`
	// The timestamp when this query was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Data Source ID. The UUID that uniquely identifies this data source / SQL
	// warehouse across the API.
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool `json:"is_draft,omitempty"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool `json:"is_safe,omitempty"`

	LastModifiedBy *User `json:"last_modified_by,omitempty"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById int `json:"last_modified_by_id,omitempty"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId string `json:"latest_query_data_id,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`

	Options *QueryOptions `json:"options,omitempty"`

	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash string `json:"query_hash,omitempty"`

	Schedule *QueryInterval `json:"schedule,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *User `json:"user,omitempty"`
	// The ID of the user who created this query.
	UserId int `json:"user_id,omitempty"`

	Visualizations []Visualization `json:"visualizations,omitempty"`
}

type QueryInterval struct {
	// For weekly runs, the day of the week to start the run.
	DayOfWeek any/* MISSING TYPE */ `json:"day_of_week,omitempty"`
	// Integer number of seconds between runs.
	Interval int `json:"interval,omitempty"`
	// For daily, weekly, and monthly runs, the time of day to start the run.
	Time any/* MISSING TYPE */ `json:"time,omitempty"`
	// A date after which this schedule no longer applies.
	Until any/* MISSING TYPE */ `json:"until,omitempty"`
}

type QueryOptions struct {
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	Parameters []Parameter `json:"parameters,omitempty"`
}

type QueryPostContent struct {
	// The ID of the data source / SQL warehouse where this query will run.
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that can convey additional information about this
	// query such as usage notes.
	Description string `json:"description,omitempty"`
	// The name or title of this query to display in list views.
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any/* MISSING TYPE */ `json:"options,omitempty"`
	// The text of the query.
	Query string `json:"query,omitempty"`

	QueryId string `json:"-" path:"query_id"`
	// JSON object that describes the scheduled execution frequency. A schedule
	// object includes `interval`, `time`, `day_of_week`, and `until` fields. If
	// a scheduled is supplied, then only `interval` is required. All other
	// field can be `null`.
	Schedule *QueryInterval `json:"schedule,omitempty"`
}

type RefreshSchedule struct {
	Cron string `json:"cron,omitempty"`

	DataSourceId string `json:"data_source_id,omitempty"`

	Id string `json:"id,omitempty"`
}

type Subscription struct {
	AlertId string `json:"alert_id,omitempty"`

	Destination *Destination `json:"destination,omitempty"`

	Id string `json:"id,omitempty"`

	User *User `json:"user,omitempty"`
}

type Success struct {
	Message SuccessMessage `json:"message,omitempty"`
}

type SuccessMessage string

const SuccessMessageSuccess SuccessMessage = `Success`

// Tags can be applied to dashboards and queries. They are used for filtering
// list views.

type User struct {
	Email string `json:"email,omitempty"`

	Id int `json:"id,omitempty"`
	// Whether this user is an admin in the Databricks workspace.
	IsDbAdmin bool `json:"is_db_admin,omitempty"`

	Name string `json:"name,omitempty"`
	// The URL for the gravatar profile picture tied to this user's email
	// address.
	ProfileImageUrl string `json:"profile_image_url,omitempty"`
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type Visualization struct {
	CreatedAt string `json:"created_at,omitempty"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description string `json:"description,omitempty"`
	// The UUID for this visualization.
	Id string `json:"id,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any/* MISSING TYPE */ `json:"options,omitempty"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type string `json:"type,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`
}

type Widget struct {
	// The unique ID for this widget.
	Id int `json:"id,omitempty"`

	Options *WidgetOptions `json:"options,omitempty"`

	Visualization *Visualization `json:"visualization,omitempty"`
	// Unused field.
	Width int `json:"width,omitempty"`
}

type WidgetOptions struct {
	// Timestamp when this object was created
	CreatedAt string `json:"created_at,omitempty"`
	// The dashboard ID to which this widget belongs. Each widget can belong to
	// one dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// Whether this widget is hidden on the dashboard.
	IsHidden bool `json:"isHidden,omitempty"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings any/* MISSING TYPE */ `json:"parameterMappings,omitempty"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position any/* MISSING TYPE */ `json:"position,omitempty"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text string `json:"text,omitempty"`
	// Timestamp of the last time this object was updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Timestamp when the alert was created.

// ID of the alert.

// Timestamp when the alert was last triggered.

// Name of the alert.

// ID of the query evaluated by the alert.

// Number of seconds after being triggered before the alert rearms itself and
// can be triggered again. If `null`, alert will never be triggered again.

// State of the alert. Possible values are: `unknown` (yet to be evaluated),
// `triggered` (evaluated and fulfilled trigger conditions), or `ok` (evaluated
// and did not fulfill trigger conditions).
type AlertState string

const AlertStateOk AlertState = `ok`

const AlertStateTriggered AlertState = `triggered`

const AlertStateUnknown AlertState = `unknown`

// ID of the alert subscriber (if subscribing an alert destination). Alert
// destinations can be configured by admins through the UI. See
// `%(alertDestinationsLink)s`.

// ID of the alert subscription.

// ID of the alert subscriber (if subscribing a user).

// Timestamp when the alert was last updated.

type CreateDashboardRequest struct {
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is true.
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// Draft dashboards only appear in list views for their owners.
	IsDraft bool `json:"is_draft,omitempty"`
	// Indicates whether the dashboard is trashed. Trashed dashboards don't
	// appear in list views.
	IsTrashed bool `json:"is_trashed,omitempty"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// An array of widget objects. A complete description of widget objects can
	// be found in the response to [Retrieve A Dashboard
	// Definition](#operation/sql-analytics-fetch-dashboard). Databricks does
	// not recommend creating new widgets via this API.
	Widgets []Widget `json:"widgets,omitempty"`
}

type DeleteAlertRequest struct {
	AlertId string `json:"-" path:"alert_id"`
}

type DeleteDashboardRequest struct {
	DashboardId string `json:"-" path:"dashboard_id"`
}

type DeleteQueryRequest struct {
	QueryId string `json:"-" path:"query_id"`
}

type DeleteScheduleRequest struct {
	AlertId string `json:"-" path:"alert_id"`

	ScheduleId string `json:"-" path:"schedule_id"`
}

type GetAlertRequest struct {
	AlertId string `json:"-" path:"alert_id"`
}

type GetDashboardRequest struct {
	DashboardId string `json:"-" path:"dashboard_id"`
}

type GetPermissionsRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId string `json:"-" path:"objectId"`
	// The type of object permissions to check.
	ObjectType ObjectTypePlural `json:"-" path:"objectType"`
}

type GetPermissionsResponse struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`

	ObjectId ObjectType `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`
}

type GetQueryRequest struct {
	QueryId string `json:"-" path:"query_id"`
}

type GetSubscriptionsRequest struct {
	AlertId string `json:"-" path:"alert_id"`
}

type ListDashboardsOrder string

const ListDashboardsOrderCreatedAt ListDashboardsOrder = `created_at`

const ListDashboardsOrderName ListDashboardsOrder = `name`

type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order ListDashboardsOrder `json:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page int `json:"-" url:"page,omitempty"`
	// Number of dashboards to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Full text search term.
	Q string `json:"-" url:"q,omitempty"`
}

type ListDashboardsResponse struct {
	// The total number of dashboards.
	Count int `json:"count,omitempty"`
	// The current page being displayed.
	Page int `json:"page,omitempty"`
	// The number of dashboards per page.
	PageSize int `json:"page_size,omitempty"`
	// List of dashboards returned.
	Results []Dashboard `json:"results,omitempty"`
}

type ListQueriesRequest struct {
	// Name of query attribute to order by. Default sort order is ascending.
	// Append a dash (`-`) to order descending instead. - `name`: The name of
	// the query. - `created_at`: The timestamp the query was created. -
	// `schedule`: The refresh interval for each query. For example: "Every 5
	// Hours" or "Every 5 Minutes". "Never" is treated as the highest value for
	// sorting. - `runtime`: The time it took to run this query. This is blank
	// for parameterized queries. A blank value is treated as the highest value
	// for sorting. - `executed_at`: The timestamp when the query was last run.
	// - `created_by`: The user name of the user that created the query.
	Order string `json:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page int `json:"-" url:"page,omitempty"`
	// Number of queries to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Full text search term
	Q string `json:"-" url:"q,omitempty"`
}

type ListQueriesResponse struct {
	// The total number of queries.
	Count int `json:"count,omitempty"`
	// The page number that is currently displayed.
	Page int `json:"page,omitempty"`
	// The number of queries per page.
	PageSize int `json:"page_size,omitempty"`
	// List of queries returned.
	Results []Query `json:"results,omitempty"`
}

type ListSchedulesRequest struct {
	AlertId string `json:"-" path:"alert_id"`
}

// A UUID generated by the application.

// A singular noun object type.
type ObjectType string

const ObjectTypeAlert ObjectType = `alert`

const ObjectTypeDashboard ObjectType = `dashboard`

const ObjectTypeDataSource ObjectType = `data_source`

const ObjectTypeQuery ObjectType = `query`

// An object's type and UUID, separated by a forward slash (/) character.

// Always a plural of the object type.
type ObjectTypePlural string

const ObjectTypePluralAlerts ObjectTypePlural = `alerts`

const ObjectTypePluralDashboards ObjectTypePlural = `dashboards`

const ObjectTypePluralDataSources ObjectTypePlural = `data_sources`

const ObjectTypePluralQueries ObjectTypePlural = `queries`

// The singular form of the type of object which can be owned.
type OwnableObjectType string

const OwnableObjectTypeAlert OwnableObjectType = `alert`

const OwnableObjectTypeDashboard OwnableObjectType = `dashboard`

const OwnableObjectTypeQuery OwnableObjectType = `query`

// Cron string representing the refresh schedule.

// ID of the SQL warehouse to refresh with. If `null`, query's SQL warehouse
// will be used to refresh.

// ID of the refresh schedule.

type RestoreDashboardRequest struct {
	DashboardId string `json:"-" path:"dashboard_id"`
}

type RestoreQueryRequest struct {
	QueryId string `json:"-" path:"query_id"`
}

type SetPermissionsRequest struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId string `json:"-" path:"objectId"`
	// The type of object permission to set.
	ObjectType ObjectTypePlural `json:"-" path:"objectType"`
}

type SetPermissionsResponse struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`

	ObjectId ObjectType `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner string `json:"new_owner,omitempty"`
}

type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner string `json:"new_owner,omitempty"`
	// The ID of the object on which to change ownership.
	ObjectId TransferOwnershipObjectId `json:"-" path:"objectId"`
	// The type of object on which to change ownership.
	ObjectType OwnableObjectType `json:"-" path:"objectType"`
}

type UnsubscribeRequest struct {
	AlertId string `json:"-" path:"alert_id"`

	SubscriptionId string `json:"-" path:"subscription_id"`
}
