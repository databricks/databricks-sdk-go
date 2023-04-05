// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"context"
)

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or notification
// destinations if the condition was met.
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

// In general, there is little need to modify dashboards using the API. However,
// it can be useful to use dashboard objects to look-up a collection of related
// query IDs. The API can also be used to duplicate multiple dashboards at once
// since you can get a dashboard definition with a GET request and then POST it
// to create a new one.
type DashboardsService interface {

	// Create a dashboard object.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

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
	// Use ListAll() to get all Dashboard instances, which will iterate over every result page.
	List(ctx context.Context, request ListDashboardsRequest) (*ListResponse, error)

	// Restore a dashboard.
	//
	// A restored dashboard appears in list views and searches and can be
	// shared.
	Restore(ctx context.Context, request RestoreDashboardRequest) error
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
// tags, parameters, and visualizations.
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

// The SQL Statement Execution API manages the execution of arbitrary SQL
// statements and the fetching of result data.
//
// **Release status**
//
// This feature is in [Public Preview].
//
// **Getting started**
//
// We suggest beginning with the [SQL Statement Execution API tutorial].
//
// **Overview of statement execution and result fetching**
//
// Statement execution begins by issuing a
// :method:statementexecution/executeStatement request with a valid SQL
// statement and warehouse ID, along with optional parameters such as the data
// catalog and output format.
//
// When submitting the statement, the call can behave synchronously or
// asynchronously, based on the `wait_timeout` setting. When set between 5-50
// seconds (default: 10) the call behaves synchronously and waits for results up
// to the specified timeout; when set to `0s`, the call is asynchronous and
// responds immediately with a statement ID that can be used to poll for status
// or fetch the results in a separate call.
//
// **Call mode: synchronous**
//
// In synchronous mode, when statement execution completes within the `wait
// timeout`, the result data is returned directly in the response. This response
// will contain `statement_id`, `status`, `manifest`, and `result` fields. The
// `status` field confirms success whereas the `manifest` field contains the
// result data column schema and metadata about the result set. The `result`
// field contains the first chunk of result data according to the specified
// `disposition`, and links to fetch any remaining chunks.
//
// If the execution does not complete before `wait_timeout`, the setting
// `on_wait_timeout` determines how the system responds.
//
// By default, `on_wait_timeout=CONTINUE`, and after reaching `wait_timeout`, a
// response is returned and statement execution continues asynchronously. The
// response will contain only `statement_id` and `status` fields, and the caller
// must now follow the flow described for asynchronous call mode to poll and
// fetch the result.
//
// Alternatively, `on_wait_timeout` can also be set to `CANCEL`; in this case if
// the timeout is reached before execution completes, the underlying statement
// execution is canceled, and a `CANCELED` status is returned in the response.
//
// **Call mode: asynchronous**
//
// In asynchronous mode, or after a timed-out synchronous request continues, a
// `statement_id` and `status` will be returned. In this case polling
// :method:statementexecution/getStatement calls are required to fetch the
// result and metadata.
//
// Next, a caller must poll until execution completes (`SUCCEEDED`, `FAILED`,
// etc.) by issuing :method:statementexecution/getStatement requests for the
// given `statement_id`.
//
// When execution has succeeded, the response will contain `status`, `manifest`,
// and `result` fields. These fields and the structure are identical to those in
// the response to a successful synchronous submission. The `result` field will
// contain the first chunk of result data, either `INLINE` or as
// `EXTERNAL_LINKS` depending on `disposition`. Additional chunks of result data
// can be fetched by checking for the presence of the `next_chunk_internal_link`
// field, and iteratively `GET` those paths until that field is unset: `GET
// https://$DATABRICKS_HOST/{next_chunk_internal_link}`.
//
// **Fetching result data: format and disposition**
//
// Result data from statement execution is available in two formats: JSON, and
// [Apache Arrow Columnar]. Statements producing a result set smaller than 16
// MiB can be fetched as `format=JSON_ARRAY`, using the `disposition=INLINE`.
// When a statement executed in `INLINE` disposition exceeds this limit, the
// execution is aborted, and no result can be fetched. Using
// `format=ARROW_STREAM` and `disposition=EXTERNAL_LINKS` allows large result
// sets, and with higher throughput.
//
// The API uses defaults of `format=JSON_ARRAY` and `disposition=INLINE`. `We
// advise explicitly setting format and disposition in all production use cases.
//
// **Statement response: statement_id, status, manifest, and result**
//
// The base call :method:statementexecution/getStatement returns a single
// response combining `statement_id`, `status`, a result `manifest`, and a
// `result` data chunk or link, depending on the `disposition`. The `manifest`
// contains the result schema definition and the result summary metadata. When
// using `disposition=EXTERNAL_LINKS`, it also contains a full listing of all
// chunks and their summary metadata.
//
// **Use case: small result sets with INLINE + JSON_ARRAY**
//
// For flows that generate small and predictable result sets (<= 16 MiB),
// `INLINE` downloads of `JSON_ARRAY` result data are typically the simplest way
// to execute and fetch result data.
//
// When the result set with `disposition=INLINE` is larger, the result can be
// transferred in chunks. After receiving the initial chunk with
// :method:statementexecution/executeStatement or
// :method:statementexecution/getStatement subsequent calls are required to
// iteratively fetch each chunk. Each result response contains a link to the
// next chunk, when there are additional chunks to fetch; it can be found in the
// field `.next_chunk_internal_link`. This link is an absolute `path` to be
// joined with your `$DATABRICKS_HOST`, and of the form
// `/api/2.0/sql/statements/{statement_id}/result/chunks/{chunk_index}`. The
// next chunk can be fetched by issuing a
// :method:statementexecution/getStatementResultChunkN request.
//
// When using this mode, each chunk may be fetched once, and in order. A chunk
// without a field `next_chunk_internal_link` indicates the last chunk was
// reached and all chunks have been fetched from the result set.
//
// **Use case: large result sets with EXTERNAL_LINKS + ARROW_STREAM**
//
// Using `EXTERNAL_LINKS` to fetch result data in Arrow format allows you to
// fetch large result sets efficiently. The primary difference from using
// `INLINE` disposition is that fetched result chunks contain resolved
// `external_links` URLs, which can be fetched with standard HTTP.
//
// **Presigned URLs**
//
// External links point to data stored within your workspace's internal DBFS, in
// the form of a presigned URL. The URLs are valid for only a short period, <=
// 15 minutes. Alongside each `external_link` is an expiration field indicating
// the time at which the URL is no longer valid. In `EXTERNAL_LINKS` mode,
// chunks can be resolved and fetched multiple times and in parallel.
//
// ----
//
// ### **Warning: We recommend you protect the URLs in the EXTERNAL_LINKS.**
//
// When using the EXTERNAL_LINKS disposition, a short-lived pre-signed URL is
// generated, which the client can use to download the result chunk directly
// from cloud storage. As the short-lived credential is embedded in a pre-signed
// URL, this URL should be protected.
//
// Since pre-signed URLs are generated with embedded temporary credentials, you
// need to remove the authorization header from the fetch requests.
//
// ----
//
// Similar to `INLINE` mode, callers can iterate through the result set, by
// using the `next_chunk_internal_link` field. Each internal link response will
// contain an external link to the raw chunk data, and additionally contain the
// `next_chunk_internal_link` if there are more chunks.
//
// Unlike `INLINE` mode, when using `EXTERNAL_LINKS`, chunks may be fetched out
// of order, and in parallel to achieve higher throughput.
//
// **Limits and limitations**
//
// Note: All byte limits are calculated based on internal storage metrics and
// will not match byte counts of actual payloads.
//
// - Statements with `disposition=INLINE` are limited to 16 MiB and will abort
// when this limit is exceeded. - Statements with `disposition=EXTERNAL_LINKS`
// are limited to 100 GiB. - The maximum query text size is 16 MiB. -
// Cancelation may silently fail. A successful response from a cancel request
// indicates that the cancel request was successfully received and sent to the
// processing engine. However, for example, an outstanding statement may
// complete execution during signal delivery, with the cancel signal arriving
// too late to be meaningful. Polling for status until a terminal state is
// reached is a reliable way to determine the final state. - Wait timeouts are
// approximate, occur server-side, and cannot account for caller delays, network
// latency from caller to service, and similarly. - After a statement has been
// submitted and a statement_id is returned, that statement's status and result
// will automatically close after either of 2 conditions: - The last result
// chunk is fetched (or resolved to an external link). - One hour passes with no
// calls to get the status or fetch the result. Best practice: in asynchronous
// clients, poll for status regularly (and with backoff) to keep the statement
// open and alive. - After fetching the last result chunk (including
// chunk_index=0) the statement is automatically closed.
//
// [Apache Arrow Columnar]: https://arrow.apache.org/overview/
// [Public Preview]: https://docs.databricks.com/release-notes/release-types.html
// [SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
type StatementExecutionService interface {

	// Cancel statement execution.
	//
	// Requests that an executing statement be canceled. Callers must poll for
	// status to see the terminal state.
	CancelExecution(ctx context.Context, request CancelExecutionRequest) error

	// Execute a SQL statement.
	//
	// Execute a SQL statement, and if flagged as such, await its result for a
	// specified time.
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
	// **NOTE** This call currently may take up to 5 seconds to get the latest
	// status and result.
	GetStatement(ctx context.Context, request GetStatementRequest) (*GetStatementResponse, error)

	// Get result chunk by index.
	//
	// After the statement execution has `SUCCEEDED`, the result data can be
	// fetched by chunks. Whereas the first chuck with `chunk_index=0` is
	// typically fetched through a `get status` request, subsequent chunks can
	// be fetched using a `get result` request. The response structure is
	// identical to the nested `result` element described in the `get status`
	// request, and similarly includes the `next_chunk_index` and
	// `next_chunk_internal_link` fields for simple iteration through the result
	// set.
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
}
