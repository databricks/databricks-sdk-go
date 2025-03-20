// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or notification
// destinations if the condition was met. Alerts can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
type AlertsClient struct {
	AlertsAPI
}

func NewAlertsClient(cfg *config.Config) (*AlertsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AlertsClient{
		AlertsAPI: AlertsAPI{
			alertsImpl: alertsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The alerts API can be used to perform CRUD operations on alerts. An alert is
// a Databricks SQL object that periodically runs a query, evaluates a condition
// of its result, and notifies one or more users and/or notification
// destinations if the condition was met. Alerts can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
//
// **Note**: A new version of the Databricks SQL API is now available. Please
// see the latest version. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type AlertsLegacyClient struct {
	AlertsLegacyAPI
}

func NewAlertsLegacyClient(cfg *config.Config) (*AlertsLegacyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &AlertsLegacyClient{
		AlertsLegacyAPI: AlertsLegacyAPI{
			alertsLegacyImpl: alertsLegacyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// This is an evolving API that facilitates the addition and removal of widgets
// from existing dashboards within the Databricks Workspace. Data structures may
// change over time.
type DashboardWidgetsClient struct {
	DashboardWidgetsAPI
}

func NewDashboardWidgetsClient(cfg *config.Config) (*DashboardWidgetsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DashboardWidgetsClient{
		DashboardWidgetsAPI: DashboardWidgetsAPI{
			dashboardWidgetsImpl: dashboardWidgetsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// In general, there is little need to modify dashboards using the API. However,
// it can be useful to use dashboard objects to look-up a collection of related
// query IDs. The API can also be used to duplicate multiple dashboards at once
// since you can get a dashboard definition with a GET request and then POST it
// to create a new one. Dashboards can be scheduled using the `sql_task` type of
// the Jobs API, e.g. :method:jobs/create.
type DashboardsClient struct {
	DashboardsAPI
}

func NewDashboardsClient(cfg *config.Config) (*DashboardsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DashboardsClient{
		DashboardsAPI: DashboardsAPI{
			dashboardsImpl: dashboardsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
//
// **Note**: A new version of the Databricks SQL API is now available. [Learn
// more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type DataSourcesClient struct {
	DataSourcesAPI
}

func NewDataSourcesClient(cfg *config.Config) (*DataSourcesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DataSourcesClient{
		DataSourcesAPI: DataSourcesAPI{
			dataSourcesImpl: dataSourcesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
//
// **Note**: A new version of the Databricks SQL API is now available. [Learn
// more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type DbsqlPermissionsClient struct {
	DbsqlPermissionsAPI
}

func NewDbsqlPermissionsClient(cfg *config.Config) (*DbsqlPermissionsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &DbsqlPermissionsClient{
		DbsqlPermissionsAPI: DbsqlPermissionsAPI{
			dbsqlPermissionsImpl: dbsqlPermissionsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The queries API can be used to perform CRUD operations on queries. A query is
// a Databricks SQL object that includes the target SQL warehouse, query text,
// name, description, tags, and parameters. Queries can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
type QueriesClient struct {
	QueriesAPI
}

func NewQueriesClient(cfg *config.Config) (*QueriesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueriesClient{
		QueriesAPI: QueriesAPI{
			queriesImpl: queriesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These endpoints are used for CRUD operations on query definitions. Query
// definitions include the target SQL warehouse, query text, name, description,
// tags, parameters, and visualizations. Queries can be scheduled using the
// `sql_task` type of the Jobs API, e.g. :method:jobs/create.
//
// **Note**: A new version of the Databricks SQL API is now available. Please
// see the latest version. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type QueriesLegacyClient struct {
	QueriesLegacyAPI
}

func NewQueriesLegacyClient(cfg *config.Config) (*QueriesLegacyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueriesLegacyClient{
		QueriesLegacyAPI: QueriesLegacyAPI{
			queriesLegacyImpl: queriesLegacyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// A service responsible for storing and retrieving the list of queries run
// against SQL endpoints and serverless compute.
type QueryHistoryClient struct {
	QueryHistoryAPI
}

func NewQueryHistoryClient(cfg *config.Config) (*QueryHistoryClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryHistoryClient{
		QueryHistoryAPI: QueryHistoryAPI{
			queryHistoryImpl: queryHistoryImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// This is an evolving API that facilitates the addition and removal of
// visualizations from existing queries in the Databricks Workspace. Data
// structures can change over time.
type QueryVisualizationsClient struct {
	QueryVisualizationsAPI
}

func NewQueryVisualizationsClient(cfg *config.Config) (*QueryVisualizationsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsClient{
		QueryVisualizationsAPI: QueryVisualizationsAPI{
			queryVisualizationsImpl: queryVisualizationsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// This is an evolving API that facilitates the addition and removal of
// vizualisations from existing queries within the Databricks Workspace. Data
// structures may change over time.
//
// **Note**: A new version of the Databricks SQL API is now available. Please
// see the latest version. [Learn more]
//
// [Learn more]: https://docs.databricks.com/en/sql/dbsql-api-latest.html
type QueryVisualizationsLegacyClient struct {
	QueryVisualizationsLegacyAPI
}

func NewQueryVisualizationsLegacyClient(cfg *config.Config) (*QueryVisualizationsLegacyClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &QueryVisualizationsLegacyClient{
		QueryVisualizationsLegacyAPI: QueryVisualizationsLegacyAPI{
			queryVisualizationsLegacyImpl: queryVisualizationsLegacyImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// Redash V2 service for workspace configurations (internal)
type RedashConfigClient struct {
	RedashConfigAPI
}

func NewRedashConfigClient(cfg *config.Config) (*RedashConfigClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &RedashConfigClient{
		RedashConfigAPI: RedashConfigAPI{
			redashConfigImpl: redashConfigImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
// network latency from caller to service. - To guarantee that the statement is
// kept alive, you must poll at least once every 15 minutes. - The results are
// only available for one hour after success; polling does not extend this. -
// The SQL Execution API must be used for the entire lifecycle of the statement.
// For example, you cannot use the Jobs API to execute the command, and then the
// SQL Execution API to cancel it.
//
// [Apache Arrow Columnar]: https://arrow.apache.org/overview/
// [Databricks SQL Statement Execution API tutorial]: https://docs.databricks.com/sql/api/sql-execution-tutorial.html
type StatementExecutionClient struct {
	StatementExecutionAPI
}

func NewStatementExecutionClient(cfg *config.Config) (*StatementExecutionClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &StatementExecutionClient{
		StatementExecutionAPI: StatementExecutionAPI{
			statementExecutionImpl: statementExecutionImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// A SQL warehouse is a compute resource that lets you run SQL commands on data
// objects within Databricks SQL. Compute resources are infrastructure resources
// that provide processing capabilities in the cloud.
type WarehousesClient struct {
	WarehousesAPI
}

func NewWarehousesClient(cfg *config.Config) (*WarehousesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WarehousesClient{
		WarehousesAPI: WarehousesAPI{
			warehousesImpl: warehousesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
