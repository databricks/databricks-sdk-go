// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
)

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or notification
// destinations if the condition was met. Alerts can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
type AlertsService interface {

	// Create an alert.
	//
	// Creates an alert. An alert is a Databricks SQL object that periodically
	// runs a query, evaluates a condition of its result, and notifies users or
	// notification destinations if the condition was met.
	Create(ctx context.Context, request CreateAlert) (*Alert, error)

	// Delete an alert.
	//
	// Deletes an alert. Deleted alerts are no longer accessible and cannot be
	// restored. **Note:** Unlike queries and dashboards, alerts cannot be moved
	// to the trash.
	Delete(ctx context.Context, request DeleteAlertRequest) error

	// Get an alert.
	//
	// Gets an alert.
	Get(ctx context.Context, request GetAlertRequest) (*Alert, error)

	// Get alerts.
	//
	// Gets a list of alerts.
	List(ctx context.Context) ([]Alert, error)

	// Update an alert.
	//
	// Updates an alert.
	Update(ctx context.Context, request EditAlert) error
}

// This is an evolving API that facilitates the addition and removal of widgets
// from existing dashboards within the Databricks Workspace. Data structures may
// change over time.
type DashboardWidgetsService interface {

	// Add widget to a dashboard.
	Create(ctx context.Context, request CreateWidget) (*Widget, error)

	// Remove widget.
	Delete(ctx context.Context, request DeleteDashboardWidgetRequest) error

	// Update existing widget.
	Update(ctx context.Context, request CreateWidget) (*Widget, error)
}

// In general, there is little need to modify dashboards using the API. However,
// it can be useful to use dashboard objects to look-up a collection of related
// query IDs. The API can also be used to duplicate multiple dashboards at once
// since you can get a dashboard definition with a GET request and then POST it
// to create a new one. Dashboards can be scheduled using the `sql_task` type of
// the Jobs API, e.g. :method:jobs/create.
type DashboardsService interface {

	// Create a dashboard object.
	Create(ctx context.Context, request DashboardPostContent) (*Dashboard, error)

	// Remove a dashboard.
	//
	// Moves a dashboard to the trash. Trashed dashboards do not appear in list
	// views or searches, and cannot be shared.
	Delete(ctx context.Context, request DeleteDashboardRequest) error

	// Retrieve a definition.
	//
	// Returns a JSON representation of a dashboard object, including its
	// visualization and query objects.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get dashboard objects.
	//
	// Fetch a paginated list of dashboard objects.
	//
	// ### **Warning: Calling this API concurrently 10 or more times could
	// result in throttling, service degradation, or a temporary ban.**
	//
	// Use ListAll() to get all Dashboard instances, which will iterate over every result page.
	List(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error)

	// Restore a dashboard.
	//
	// A restored dashboard appears in list views and searches and can be
	// shared.
	Restore(ctx context.Context, request RestoreDashboardRequest) error

	// Change a dashboard definition.
	//
	// Modify this dashboard definition. This operation only affects attributes
	// of the dashboard object. It does not add, modify, or remove widgets.
	//
	// **Note**: You cannot undo this operation.
	Update(ctx context.Context, request DashboardEditContent) (*Dashboard, error)
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
type DataSourcesService interface {

	// Get a list of SQL warehouses.
	//
	// Retrieves a full list of SQL warehouses available in this workspace. All
	// fields that appear in this API response are enumerated for clarity.
	// However, you need only a SQL warehouse's `id` to create new queries
	// against it.
	List(ctx context.Context) ([]DataSource, error)
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
type DbsqlPermissionsService interface {

	// Get object ACL.
	//
	// Gets a JSON representation of the access control list (ACL) for a
	// specified object.
	Get(ctx context.Context, request GetDbsqlPermissionRequest) (*GetResponse, error)

	// Set object ACL.
	//
	// Sets the access control list (ACL) for a specified object. This operation
	// will complete rewrite the ACL.
	Set(ctx context.Context, request SetRequest) (*SetResponse, error)

	// Transfer object ownership.
	//
	// Transfers ownership of a dashboard, query, or alert to an active user.
	// Requires an admin API key.
	TransferOwnership(ctx context.Context, request TransferOwnershipRequest) (*Success, error)
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, parameters, and visualizations. Queries can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
type QueriesService interface {

	// Create a new query definition.
	//
	// Creates a new query definition. Queries created with this endpoint belong
	// to the authenticated user making the request.
	//
	// The `data_source_id` field specifies the ID of the SQL warehouse to run
	// this query against. You can use the Data Sources API to see a complete
	// list of available SQL warehouses. Or you can copy the `data_source_id`
	// from an existing query.
	//
	// **Note**: You cannot add a visualization until you create the query.
	Create(ctx context.Context, request QueryPostContent) (*Query, error)

	// Delete a query.
	//
	// Moves a query to the trash. Trashed queries immediately disappear from
	// searches and list views, and they cannot be used for alerts. The trash is
	// deleted after 30 days.
	Delete(ctx context.Context, request DeleteQueryRequest) error

	// Get a query definition.
	//
	// Retrieve a query object definition along with contextual permissions
	// information about the currently authenticated user.
	Get(ctx context.Context, request GetQueryRequest) (*Query, error)

	// Get a list of queries.
	//
	// Gets a list of queries. Optionally, this list can be filtered by a search
	// term.
	//
	// ### **Warning: Calling this API concurrently 10 or more times could
	// result in throttling, service degradation, or a temporary ban.**
	//
	// Use ListAll() to get all Query instances, which will iterate over every result page.
	List(ctx context.Context, request ListQueriesRequest) (*QueryList, error)

	// Restore a query.
	//
	// Restore a query that has been moved to the trash. A restored query
	// appears in list views and searches. You can use restored queries for
	// alerts.
	Restore(ctx context.Context, request RestoreQueryRequest) error

	// Change a query definition.
	//
	// Modify this query definition.
	//
	// **Note**: You cannot undo this operation.
	Update(ctx context.Context, request QueryEditContent) (*Query, error)
}

// Access the history of queries through SQL warehouses.
type QueryHistoryService interface {

	// List Queries.
	//
	// List the history of queries through SQL warehouses.
	//
	// You can filter by user ID, warehouse ID, status, and time range.
	//
	// Use ListAll() to get all QueryInfo instances, which will iterate over every result page.
	List(ctx context.Context, request ListQueryHistoryRequest) (*ListQueriesResponse, error)
}

// This is an evolving API that facilitates the addition and removal of
// vizualisations from existing queries within the Databricks Workspace. Data
// structures may change over time.
type QueryVisualizationsService interface {

	// Add visualization to a query.
	Create(ctx context.Context, request CreateQueryVisualizationRequest) (*Visualization, error)

	// Remove visualization.
	Delete(ctx context.Context, request DeleteQueryVisualizationRequest) error

	// Edit existing visualization.
	Update(ctx context.Context, request Visualization) (*Visualization, error)
}

// The Databricks SQL Statement Execution API can be used to execute SQL
// statements on a SQL warehouse and fetch the result.
//
// **Getting started**
//
// We suggest beginning with the [Databricks SQL Statement Execution API
// tutorial].
//
// **Overview of statement execution and result fetching**
//
// Statement execution begins by issuing a
// :method:statementexecution/executeStatement request with a valid SQL
// statement and warehouse ID, along with optional parameters such as the data
// catalog and output format. If no other parameters are specified, the server
// will wait for up to 10s before returning a response. If the statement has
// completed within this timespan, the response will include the result data as
// a JSON array and metadata. Otherwise, if no result is available after the 10s
// timeout expired, the response will provide the statement ID that can be used
// to poll for results by using a :method:statementexecution/getStatement
// request.
//
// You can specify whether the call should behave synchronously, asynchronously
// or start synchronously with a fallback to asynchronous execution. This is
// controlled with the `wait_timeout` and `on_wait_timeout` settings. If
// `wait_timeout` is set between 5-50 seconds (default: 10s), the call waits for
// results up to the specified timeout; when set to `0s`, the call is
// asynchronous and responds immediately with a statement ID. The
// `on_wait_timeout` setting specifies what should happen when the timeout is
// reached while the statement execution has not yet finished. This can be set
// to either `CONTINUE`, to fallback to asynchronous mode, or it can be set to
// `CANCEL`, which cancels the statement.
//
// In summary: - Synchronous mode - `wait_timeout=30s` and
// `on_wait_timeout=CANCEL` - The call waits up to 30 seconds; if the statement
// execution finishes within this time, the result data is returned directly in
// the response. If the execution takes longer than 30 seconds, the execution is
// canceled and the call returns with a `CANCELED` state. - Asynchronous mode -
// `wait_timeout=0s` (`on_wait_timeout` is ignored) - The call doesn't wait for
// the statement to finish but returns directly with a statement ID. The status
// of the statement execution can be polled by issuing
// :method:statementexecution/getStatement with the statement ID. Once the
// execution has succeeded, this call also returns the result and metadata in
// the response. - Hybrid mode (default) - `wait_timeout=10s` and
// `on_wait_timeout=CONTINUE` - The call waits for up to 10 seconds; if the
// statement execution finishes within this time, the result data is returned
// directly in the response. If the execution takes longer than 10 seconds, a
// statement ID is returned. The statement ID can be used to fetch status and
// results in the same way as in the asynchronous mode.
//
// Depending on the size, the result can be split into multiple chunks. If the
// statement execution is successful, the statement response contains a manifest
// and the first chunk of the result. The manifest contains schema information
// and provides metadata for each chunk in the result. Result chunks can be
// retrieved by index with :method:statementexecution/getStatementResultChunkN
// which may be called in any order and in parallel. For sequential fetching,
// each chunk, apart from the last, also contains a `next_chunk_index` and
// `next_chunk_internal_link` that point to the next chunk.
//
// A statement can be canceled with :method:statementexecution/cancelExecution.
//
// **Fetching result data: format and disposition**
//
// To specify the format of the result data, use the `format` field, which can
// be set to one of the following options: `JSON_ARRAY` (JSON), `ARROW_STREAM`
// ([Apache Arrow Columnar]), or `CSV`.
//
// There are two ways to receive statement results, controlled by the
// `disposition` setting, which can be either `INLINE` or `EXTERNAL_LINKS`:
//
// - `INLINE`: In this mode, the result data is directly included in the
// response. It's best suited for smaller results. This mode can only be used
// with the `JSON_ARRAY` format.
//
// - `EXTERNAL_LINKS`: In this mode, the response provides links that can be
// used to download the result data in chunks separately. This approach is ideal
// for larger results and offers higher throughput. This mode can be used with
// all the formats: `JSON_ARRAY`, `ARROW_STREAM`, and `CSV`.
//
// By default, the API uses `format=JSON_ARRAY` and `disposition=INLINE`.
//
// **Limits and limitations**
//
// Note: The byte limit for INLINE disposition is based on internal storage
// metrics and will not exactly match the byte count of the actual payload.
//
// - Statements with `disposition=INLINE` are limited to 25 MiB and will fail
// when this limit is exceeded. - Statements with `disposition=EXTERNAL_LINKS`
// are limited to 100 GiB. Result sets larger than this limit will be truncated.
// Truncation is indicated by the `truncated` field in the result manifest. -
// The maximum query text size is 16 MiB. - Cancelation might silently fail. A
// successful response from a cancel request indicates that the cancel request
// was successfully received and sent to the processing engine. However, an
// outstanding statement might have already completed execution when the cancel
// request arrives. Polling for status until a terminal state is reached is a
// reliable way to determine the final state. - Wait timeouts are approximate,
// occur server-side, and cannot account for things such as caller delays and
// network latency from caller to service. - The system will auto-close a
// statement after one hour if the client stops polling and thus you must poll
// at least once an hour. - The results are only available for one hour after
// success; polling does not extend this.
//
// [Apache Arrow Columnar]: https://arrow.apache.org/overview/
// [Databricks SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
type StatementExecutionService interface {

	// Cancel statement execution.
	//
	// Requests that an executing statement be canceled. Callers must poll for
	// status to see the terminal state.
	CancelExecution(ctx context.Context, request CancelExecutionRequest) error

	// Execute a SQL statement.
	ExecuteStatement(ctx context.Context, request ExecuteStatementRequest) (*ExecuteStatementResponse, error)

	// Get status, manifest, and result first chunk.
	//
	// This request can be used to poll for the statement's status. When the
	// `status.state` field is `SUCCEEDED` it will also return the result
	// manifest and the first chunk of the result data. When the statement is in
	// the terminal states `CANCELED`, `CLOSED` or `FAILED`, it returns HTTP 200
	// with the state set. After at least 12 hours in terminal state, the
	// statement is removed from the warehouse and further calls will receive an
	// HTTP 404 response.
	//
	// **NOTE** This call currently might take up to 5 seconds to get the latest
	// status and result.
	GetStatement(ctx context.Context, request GetStatementRequest) (*GetStatementResponse, error)

	// Get result chunk by index.
	//
	// After the statement execution has `SUCCEEDED`, this request can be used
	// to fetch any chunk by index. Whereas the first chunk with `chunk_index=0`
	// is typically fetched with :method:statementexecution/executeStatement or
	// :method:statementexecution/getStatement, this request can be used to
	// fetch subsequent chunks. The response structure is identical to the
	// nested `result` element described in the
	// :method:statementexecution/getStatement request, and similarly includes
	// the `next_chunk_index` and `next_chunk_internal_link` fields for simple
	// iteration through the result set.
	GetStatementResultChunkN(ctx context.Context, request GetStatementResultChunkNRequest) (*ResultData, error)
}

// A SQL warehouse is a compute resource that lets you run SQL commands on data
// objects within Databricks SQL. Compute resources are infrastructure resources
// that provide processing capabilities in the cloud.
type WarehousesService interface {

	// Create a warehouse.
	//
	// Creates a new SQL warehouse.
	Create(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error)

	// Delete a warehouse.
	//
	// Deletes a SQL warehouse.
	Delete(ctx context.Context, request DeleteWarehouseRequest) error

	// Update a warehouse.
	//
	// Updates the configuration for a SQL warehouse.
	Edit(ctx context.Context, request EditWarehouseRequest) error

	// Get warehouse info.
	//
	// Gets the information for a single SQL warehouse.
	Get(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error)

	// Get SQL warehouse permission levels.
	//
	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetWarehousePermissionLevelsRequest) (*GetWarehousePermissionLevelsResponse, error)

	// Get SQL warehouse permissions.
	//
	// Gets the permissions of a SQL warehouse. SQL warehouses can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetWarehousePermissionsRequest) (*WarehousePermissions, error)

	// Get the workspace configuration.
	//
	// Gets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error)

	// List warehouses.
	//
	// Lists all SQL warehouses that a user has manager permissions on.
	//
	// Use ListAll() to get all EndpointInfo instances
	List(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error)

	// Set SQL warehouse permissions.
	//
	// Sets permissions on a SQL warehouse. SQL warehouses can inherit
	// permissions from their root object.
	SetPermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error)

	// Set the workspace configuration.
	//
	// Sets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error

	// Start a warehouse.
	//
	// Starts a SQL warehouse.
	Start(ctx context.Context, request StartRequest) error

	// Stop a warehouse.
	//
	// Stops a SQL warehouse.
	Stop(ctx context.Context, request StopRequest) error

	// Update SQL warehouse permissions.
	//
	// Updates the permissions on a SQL warehouse. SQL warehouses can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request WarehousePermissionsRequest) (*WarehousePermissions, error)
}
