// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dataquality

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// The granularity for aggregating data into time windows based on their
// timestamp.
type AggregationGranularity string

const AggregationGranularityAggregationGranularity1Day AggregationGranularity = `AGGREGATION_GRANULARITY_1_DAY`

const AggregationGranularityAggregationGranularity1Hour AggregationGranularity = `AGGREGATION_GRANULARITY_1_HOUR`

const AggregationGranularityAggregationGranularity1Month AggregationGranularity = `AGGREGATION_GRANULARITY_1_MONTH`

const AggregationGranularityAggregationGranularity1Week AggregationGranularity = `AGGREGATION_GRANULARITY_1_WEEK`

const AggregationGranularityAggregationGranularity1Year AggregationGranularity = `AGGREGATION_GRANULARITY_1_YEAR`

const AggregationGranularityAggregationGranularity2Weeks AggregationGranularity = `AGGREGATION_GRANULARITY_2_WEEKS`

const AggregationGranularityAggregationGranularity30Minutes AggregationGranularity = `AGGREGATION_GRANULARITY_30_MINUTES`

const AggregationGranularityAggregationGranularity3Weeks AggregationGranularity = `AGGREGATION_GRANULARITY_3_WEEKS`

const AggregationGranularityAggregationGranularity4Weeks AggregationGranularity = `AGGREGATION_GRANULARITY_4_WEEKS`

const AggregationGranularityAggregationGranularity5Minutes AggregationGranularity = `AGGREGATION_GRANULARITY_5_MINUTES`

// String representation for [fmt.Print]
func (f *AggregationGranularity) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AggregationGranularity) Set(v string) error {
	switch v {
	case `AGGREGATION_GRANULARITY_1_DAY`, `AGGREGATION_GRANULARITY_1_HOUR`, `AGGREGATION_GRANULARITY_1_MONTH`, `AGGREGATION_GRANULARITY_1_WEEK`, `AGGREGATION_GRANULARITY_1_YEAR`, `AGGREGATION_GRANULARITY_2_WEEKS`, `AGGREGATION_GRANULARITY_30_MINUTES`, `AGGREGATION_GRANULARITY_3_WEEKS`, `AGGREGATION_GRANULARITY_4_WEEKS`, `AGGREGATION_GRANULARITY_5_MINUTES`:
		*f = AggregationGranularity(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AGGREGATION_GRANULARITY_1_DAY", "AGGREGATION_GRANULARITY_1_HOUR", "AGGREGATION_GRANULARITY_1_MONTH", "AGGREGATION_GRANULARITY_1_WEEK", "AGGREGATION_GRANULARITY_1_YEAR", "AGGREGATION_GRANULARITY_2_WEEKS", "AGGREGATION_GRANULARITY_30_MINUTES", "AGGREGATION_GRANULARITY_3_WEEKS", "AGGREGATION_GRANULARITY_4_WEEKS", "AGGREGATION_GRANULARITY_5_MINUTES"`, v)
	}
}

// Values returns all possible values for AggregationGranularity.
//
// There is no guarantee on the order of the values in the slice.
func (f *AggregationGranularity) Values() []AggregationGranularity {
	return []AggregationGranularity{
		AggregationGranularityAggregationGranularity1Day,
		AggregationGranularityAggregationGranularity1Hour,
		AggregationGranularityAggregationGranularity1Month,
		AggregationGranularityAggregationGranularity1Week,
		AggregationGranularityAggregationGranularity1Year,
		AggregationGranularityAggregationGranularity2Weeks,
		AggregationGranularityAggregationGranularity30Minutes,
		AggregationGranularityAggregationGranularity3Weeks,
		AggregationGranularityAggregationGranularity4Weeks,
		AggregationGranularityAggregationGranularity5Minutes,
	}
}

// Type always returns AggregationGranularity to satisfy [pflag.Value] interface
func (f *AggregationGranularity) Type() string {
	return "AggregationGranularity"
}

// Anomaly Detection Configurations.
type AnomalyDetectionConfig struct {
}

// Request to cancel a refresh.
type CancelRefreshRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
}

// Response to cancelling a refresh.
type CancelRefreshResponse struct {
	// The refresh to cancel.
	Refresh *Refresh `json:"refresh,omitempty"`
}

type CreateMonitorRequest struct {
	// The monitor to create.
	Monitor Monitor `json:"monitor"`
}

type CreateRefreshRequest struct {
	// The UUID of the request object. For example, table id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType string `json:"-" url:"-"`
	// The refresh to create
	Refresh Refresh `json:"refresh"`
}

// The data quality monitoring workflow cron schedule.
type CronSchedule struct {
	// Read only field that indicates whether the schedule is paused or not.
	PauseStatus CronSchedulePauseStatus `json:"pause_status,omitempty"`
	// The expression that determines when to run the monitor. See [examples].
	//
	// [examples]: https://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone id. The schedule for a job will be resolved with respect
	// to this timezone. See `Java TimeZone
	// <http://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html>`_ for
	// details. The timezone id (e.g., ``America/Los_Angeles``) in which to
	// evaluate the quartz expression.
	TimezoneId string `json:"timezone_id"`
}

// The data quality monitoring workflow cron schedule pause status.
type CronSchedulePauseStatus string

const CronSchedulePauseStatusCronSchedulePauseStatusPaused CronSchedulePauseStatus = `CRON_SCHEDULE_PAUSE_STATUS_PAUSED`

const CronSchedulePauseStatusCronSchedulePauseStatusUnpaused CronSchedulePauseStatus = `CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED`

// String representation for [fmt.Print]
func (f *CronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `CRON_SCHEDULE_PAUSE_STATUS_PAUSED`, `CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED`:
		*f = CronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CRON_SCHEDULE_PAUSE_STATUS_PAUSED", "CRON_SCHEDULE_PAUSE_STATUS_UNPAUSED"`, v)
	}
}

// Values returns all possible values for CronSchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *CronSchedulePauseStatus) Values() []CronSchedulePauseStatus {
	return []CronSchedulePauseStatus{
		CronSchedulePauseStatusCronSchedulePauseStatusPaused,
		CronSchedulePauseStatusCronSchedulePauseStatusUnpaused,
	}
}

// Type always returns CronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *CronSchedulePauseStatus) Type() string {
	return "CronSchedulePauseStatus"
}

// Data Profiling Configurations.
type DataProfilingConfig struct {
	// Field for specifying the absolute path to a custom directory to store
	// data-monitoring assets. Normally prepopulated to a default user location
	// via UI and Python APIs.
	AssetsDir string `json:"assets_dir,omitempty"`
	// Baseline table name. Baseline data is used to compute drift from the data
	// in the monitored `table_name`. The baseline table and the monitored table
	// shall have the same schema.
	BaselineTableName string `json:"baseline_table_name,omitempty"`
	// Custom metrics.
	CustomMetrics []DataProfilingCustomMetric `json:"custom_metrics,omitempty"`
	// Id of dashboard that visualizes the computed metrics. This can be empty
	// if the monitor is in PENDING state.
	DashboardId string `json:"dashboard_id,omitempty"`
	// Table that stores drift metrics data. Format:
	// `catalog.schema.table_name`.
	DriftMetricsTableName string `json:"drift_metrics_table_name,omitempty"`
	// The warehouse for dashboard creation
	EffectiveWarehouseId string `json:"effective_warehouse_id,omitempty"`
	// Configuration for monitoring inference log tables.
	InferenceLog *InferenceLogConfig `json:"inference_log,omitempty"`
	// The latest error message for a monitor failure.
	LatestMonitorFailureMessage string `json:"latest_monitor_failure_message,omitempty"`
	// Represents the current monitor configuration version in use. The version
	// will be represented in a numeric fashion (1,2,3...). The field has
	// flexibility to take on negative values, which can indicate corrupted
	// monitor_version numbers.
	MonitorVersion int64 `json:"monitor_version,omitempty"`
	// Unity Catalog table to monitor. Format: `catalog.schema.table_name`
	MonitoredTableName string `json:"monitored_table_name,omitempty"`
	// Field for specifying notification settings.
	NotificationSettings *NotificationSettings `json:"notification_settings,omitempty"`
	// ID of the schema where output tables are created.
	OutputSchemaId string `json:"output_schema_id"`
	// Table that stores profile metrics data. Format:
	// `catalog.schema.table_name`.
	ProfileMetricsTableName string `json:"profile_metrics_table_name,omitempty"`
	// The cron schedule.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// Whether to skip creating a default dashboard summarizing data quality
	// metrics.
	SkipBuiltinDashboard bool `json:"skip_builtin_dashboard,omitempty"`
	// List of column expressions to slice data with for targeted analysis. The
	// data is grouped by each expression independently, resulting in a separate
	// slice for each predicate and its complements. For example
	// `slicing_exprs=[“col_1”, “col_2 > 10”]` will generate the
	// following slices: two slices for `col_2 > 10` (True and False), and one
	// slice per unique value in `col1`. For high-cardinality columns, only the
	// top 100 unique values by frequency will generate slices.
	SlicingExprs []string `json:"slicing_exprs,omitempty"`
	// Configuration for monitoring snapshot tables.
	Snapshot *SnapshotConfig `json:"snapshot,omitempty"`
	// The data profiling monitor status.
	Status DataProfilingStatus `json:"status,omitempty"`
	// Configuration for monitoring time series tables.
	TimeSeries *TimeSeriesConfig `json:"time_series,omitempty"`
	// Optional argument to specify the warehouse for dashboard creation. If not
	// specified, the first running warehouse will be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataProfilingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataProfilingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Custom metric definition.
type DataProfilingCustomMetric struct {
	// Jinja template for a SQL expression that specifies how to compute the
	// metric. See [create metric definition].
	//
	// [create metric definition]: https://docs.databricks.com/en/lakehouse-monitoring/custom-metrics.html#create-definition
	Definition string `json:"definition"`
	// A list of column names in the input table the metric should be computed
	// for. Can use ``":table"`` to indicate that the metric needs information
	// from multiple columns.
	InputColumns []string `json:"input_columns"`
	// Name of the metric in the output tables.
	Name string `json:"name"`
	// The output type of the custom metric.
	OutputDataType string `json:"output_data_type"`
	// The type of the custom metric.
	Type DataProfilingCustomMetricType `json:"type"`
}

// The custom metric type.
type DataProfilingCustomMetricType string

const DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeAggregate DataProfilingCustomMetricType = `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`

const DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDerived DataProfilingCustomMetricType = `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`

const DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDrift DataProfilingCustomMetricType = `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`

// String representation for [fmt.Print]
func (f *DataProfilingCustomMetricType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataProfilingCustomMetricType) Set(v string) error {
	switch v {
	case `DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED`, `DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT`:
		*f = DataProfilingCustomMetricType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_PROFILING_CUSTOM_METRIC_TYPE_AGGREGATE", "DATA_PROFILING_CUSTOM_METRIC_TYPE_DERIVED", "DATA_PROFILING_CUSTOM_METRIC_TYPE_DRIFT"`, v)
	}
}

// Values returns all possible values for DataProfilingCustomMetricType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataProfilingCustomMetricType) Values() []DataProfilingCustomMetricType {
	return []DataProfilingCustomMetricType{
		DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeAggregate,
		DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDerived,
		DataProfilingCustomMetricTypeDataProfilingCustomMetricTypeDrift,
	}
}

// Type always returns DataProfilingCustomMetricType to satisfy [pflag.Value] interface
func (f *DataProfilingCustomMetricType) Type() string {
	return "DataProfilingCustomMetricType"
}

// The status of the data profiling monitor.
type DataProfilingStatus string

const DataProfilingStatusDataProfilingStatusActive DataProfilingStatus = `DATA_PROFILING_STATUS_ACTIVE`

const DataProfilingStatusDataProfilingStatusDeletePending DataProfilingStatus = `DATA_PROFILING_STATUS_DELETE_PENDING`

const DataProfilingStatusDataProfilingStatusError DataProfilingStatus = `DATA_PROFILING_STATUS_ERROR`

const DataProfilingStatusDataProfilingStatusFailed DataProfilingStatus = `DATA_PROFILING_STATUS_FAILED`

const DataProfilingStatusDataProfilingStatusPending DataProfilingStatus = `DATA_PROFILING_STATUS_PENDING`

// String representation for [fmt.Print]
func (f *DataProfilingStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataProfilingStatus) Set(v string) error {
	switch v {
	case `DATA_PROFILING_STATUS_ACTIVE`, `DATA_PROFILING_STATUS_DELETE_PENDING`, `DATA_PROFILING_STATUS_ERROR`, `DATA_PROFILING_STATUS_FAILED`, `DATA_PROFILING_STATUS_PENDING`:
		*f = DataProfilingStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_PROFILING_STATUS_ACTIVE", "DATA_PROFILING_STATUS_DELETE_PENDING", "DATA_PROFILING_STATUS_ERROR", "DATA_PROFILING_STATUS_FAILED", "DATA_PROFILING_STATUS_PENDING"`, v)
	}
}

// Values returns all possible values for DataProfilingStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataProfilingStatus) Values() []DataProfilingStatus {
	return []DataProfilingStatus{
		DataProfilingStatusDataProfilingStatusActive,
		DataProfilingStatusDataProfilingStatusDeletePending,
		DataProfilingStatusDataProfilingStatusError,
		DataProfilingStatusDataProfilingStatusFailed,
		DataProfilingStatusDataProfilingStatusPending,
	}
}

// Type always returns DataProfilingStatus to satisfy [pflag.Value] interface
func (f *DataProfilingStatus) Type() string {
	return "DataProfilingStatus"
}

type DeleteMonitorRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
}

type DeleteRefreshRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
}

type GetMonitorRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
}

type GetRefreshRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
}

// Inference log configuration.
type InferenceLogConfig struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities []AggregationGranularity `json:"granularities"`
	// Column for the label.
	LabelColumn string `json:"label_column,omitempty"`
	// Column for the model identifier.
	ModelIdColumn string `json:"model_id_column"`
	// Column for the prediction.
	PredictionColumn string `json:"prediction_column"`
	// Problem type the model aims to solve.
	ProblemType InferenceProblemType `json:"problem_type"`
	// Column for the timestamp.
	TimestampColumn string `json:"timestamp_column"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InferenceLogConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InferenceLogConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Inference problem type the model aims to solve.
type InferenceProblemType string

const InferenceProblemTypeInferenceProblemTypeClassification InferenceProblemType = `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`

const InferenceProblemTypeInferenceProblemTypeRegression InferenceProblemType = `INFERENCE_PROBLEM_TYPE_REGRESSION`

// String representation for [fmt.Print]
func (f *InferenceProblemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InferenceProblemType) Set(v string) error {
	switch v {
	case `INFERENCE_PROBLEM_TYPE_CLASSIFICATION`, `INFERENCE_PROBLEM_TYPE_REGRESSION`:
		*f = InferenceProblemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "INFERENCE_PROBLEM_TYPE_CLASSIFICATION", "INFERENCE_PROBLEM_TYPE_REGRESSION"`, v)
	}
}

// Values returns all possible values for InferenceProblemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *InferenceProblemType) Values() []InferenceProblemType {
	return []InferenceProblemType{
		InferenceProblemTypeInferenceProblemTypeClassification,
		InferenceProblemTypeInferenceProblemTypeRegression,
	}
}

// Type always returns InferenceProblemType to satisfy [pflag.Value] interface
func (f *InferenceProblemType) Type() string {
	return "InferenceProblemType"
}

type ListMonitorRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMonitorRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing Monitors.
type ListMonitorResponse struct {
	Monitors []Monitor `json:"monitors,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListMonitorResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListMonitorResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRefreshRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRefreshRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRefreshRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Response for listing refreshes.
type ListRefreshResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Refreshes []Refresh `json:"refreshes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListRefreshResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRefreshResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Monitor for the data quality of unity catalog entities such as schema or
// table.
type Monitor struct {
	// Anomaly Detection Configuration, applicable to `schema` object types.
	AnomalyDetectionConfig *AnomalyDetectionConfig `json:"anomaly_detection_config,omitempty"`
	// Data Profiling Configuration, applicable to `table` object types
	DataProfilingConfig *DataProfilingConfig `json:"data_profiling_config,omitempty"`
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"object_type"`
}

// Destination of the data quality monitoring notification.
type NotificationDestination struct {
	// The list of email addresses to send the notification to. A maximum of 5
	// email addresses is supported.
	EmailAddresses []string `json:"email_addresses,omitempty"`
}

// Settings for sending notifications on the data quality monitoring.
type NotificationSettings struct {
	// Destinations to send notifications on failure/timeout.
	OnFailure *NotificationDestination `json:"on_failure,omitempty"`
}

// The Refresh object gives information on a refresh of the data quality
// monitoring pipeline.
type Refresh struct {
	// Time when the refresh ended (milliseconds since 1/1/1970 UTC).
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// An optional message to give insight into the current state of the refresh
	// (e.g. FAILURE messages).
	Message string `json:"message,omitempty"`
	// The UUID of the request object. For example, table id.
	ObjectId string `json:"object_id"`
	// The type of the monitored object. Can be one of the following: `schema`or
	// `table`.
	ObjectType string `json:"object_type"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"refresh_id,omitempty"`
	// Time when the refresh started (milliseconds since 1/1/1970 UTC).
	StartTimeMs int64 `json:"start_time_ms,omitempty"`
	// The current state of the refresh.
	State RefreshState `json:"state,omitempty"`
	// What triggered the refresh.
	Trigger RefreshTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Refresh) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Refresh) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the refresh.
type RefreshState string

const RefreshStateMonitorRefreshStateCanceled RefreshState = `MONITOR_REFRESH_STATE_CANCELED`

const RefreshStateMonitorRefreshStateFailed RefreshState = `MONITOR_REFRESH_STATE_FAILED`

const RefreshStateMonitorRefreshStatePending RefreshState = `MONITOR_REFRESH_STATE_PENDING`

const RefreshStateMonitorRefreshStateRunning RefreshState = `MONITOR_REFRESH_STATE_RUNNING`

const RefreshStateMonitorRefreshStateSuccess RefreshState = `MONITOR_REFRESH_STATE_SUCCESS`

const RefreshStateMonitorRefreshStateUnknown RefreshState = `MONITOR_REFRESH_STATE_UNKNOWN`

// String representation for [fmt.Print]
func (f *RefreshState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RefreshState) Set(v string) error {
	switch v {
	case `MONITOR_REFRESH_STATE_CANCELED`, `MONITOR_REFRESH_STATE_FAILED`, `MONITOR_REFRESH_STATE_PENDING`, `MONITOR_REFRESH_STATE_RUNNING`, `MONITOR_REFRESH_STATE_SUCCESS`, `MONITOR_REFRESH_STATE_UNKNOWN`:
		*f = RefreshState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_REFRESH_STATE_CANCELED", "MONITOR_REFRESH_STATE_FAILED", "MONITOR_REFRESH_STATE_PENDING", "MONITOR_REFRESH_STATE_RUNNING", "MONITOR_REFRESH_STATE_SUCCESS", "MONITOR_REFRESH_STATE_UNKNOWN"`, v)
	}
}

// Values returns all possible values for RefreshState.
//
// There is no guarantee on the order of the values in the slice.
func (f *RefreshState) Values() []RefreshState {
	return []RefreshState{
		RefreshStateMonitorRefreshStateCanceled,
		RefreshStateMonitorRefreshStateFailed,
		RefreshStateMonitorRefreshStatePending,
		RefreshStateMonitorRefreshStateRunning,
		RefreshStateMonitorRefreshStateSuccess,
		RefreshStateMonitorRefreshStateUnknown,
	}
}

// Type always returns RefreshState to satisfy [pflag.Value] interface
func (f *RefreshState) Type() string {
	return "RefreshState"
}

// The trigger of the refresh.
type RefreshTrigger string

const RefreshTriggerMonitorRefreshTriggerDataChange RefreshTrigger = `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`

const RefreshTriggerMonitorRefreshTriggerManual RefreshTrigger = `MONITOR_REFRESH_TRIGGER_MANUAL`

const RefreshTriggerMonitorRefreshTriggerSchedule RefreshTrigger = `MONITOR_REFRESH_TRIGGER_SCHEDULE`

const RefreshTriggerMonitorRefreshTriggerUnknown RefreshTrigger = `MONITOR_REFRESH_TRIGGER_UNKNOWN`

// String representation for [fmt.Print]
func (f *RefreshTrigger) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RefreshTrigger) Set(v string) error {
	switch v {
	case `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`, `MONITOR_REFRESH_TRIGGER_MANUAL`, `MONITOR_REFRESH_TRIGGER_SCHEDULE`, `MONITOR_REFRESH_TRIGGER_UNKNOWN`:
		*f = RefreshTrigger(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONITOR_REFRESH_TRIGGER_DATA_CHANGE", "MONITOR_REFRESH_TRIGGER_MANUAL", "MONITOR_REFRESH_TRIGGER_SCHEDULE", "MONITOR_REFRESH_TRIGGER_UNKNOWN"`, v)
	}
}

// Values returns all possible values for RefreshTrigger.
//
// There is no guarantee on the order of the values in the slice.
func (f *RefreshTrigger) Values() []RefreshTrigger {
	return []RefreshTrigger{
		RefreshTriggerMonitorRefreshTriggerDataChange,
		RefreshTriggerMonitorRefreshTriggerManual,
		RefreshTriggerMonitorRefreshTriggerSchedule,
		RefreshTriggerMonitorRefreshTriggerUnknown,
	}
}

// Type always returns RefreshTrigger to satisfy [pflag.Value] interface
func (f *RefreshTrigger) Type() string {
	return "RefreshTrigger"
}

// Snapshot analysis configuration.
type SnapshotConfig struct {
}

// Time series analysis configuration.
type TimeSeriesConfig struct {
	// List of granularities to use when aggregating data into time windows
	// based on their timestamp.
	Granularities []AggregationGranularity `json:"granularities"`
	// Column for the timestamp.
	TimestampColumn string `json:"timestamp_column"`
}

type UpdateMonitorRequest struct {
	// The monitor to update.
	Monitor Monitor `json:"monitor"`
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// The field mask to specify which fields to update as a comma-separated
	// list. Example value:
	// `data_profiling_config.custom_metrics,data_profiling_config.schedule.quartz_cron_expression`
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateRefreshRequest struct {
	// The UUID of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: `schema`
	// or `table`.
	ObjectType string `json:"-" url:"-"`
	// The refresh to update.
	Refresh Refresh `json:"refresh"`
	// Unique id of the refresh operation.
	RefreshId int64 `json:"-" url:"-"`
	// The field mask to specify which fields to update.
	UpdateMask string `json:"-" url:"update_mask"`
}
