// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// Budget policy of this pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`

	DryRun bool `json:"dry_run,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *RunAs `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePipeline) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePipeline) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	EffectiveSettings *PipelineSpec `json:"effective_settings,omitempty"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CronTrigger struct {
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`

	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CronTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CronTrigger) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	Instance string `json:"instance,omitempty"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo int `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataPlaneId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Days of week in which the restart is allowed to happen (within a five-hour
// window starting at start_hour). If not specified all days of the week will be
// used.
type DayOfWeek string

const DayOfWeekFriday DayOfWeek = `FRIDAY`

const DayOfWeekMonday DayOfWeek = `MONDAY`

const DayOfWeekSaturday DayOfWeek = `SATURDAY`

const DayOfWeekSunday DayOfWeek = `SUNDAY`

const DayOfWeekThursday DayOfWeek = `THURSDAY`

const DayOfWeekTuesday DayOfWeek = `TUESDAY`

const DayOfWeekWednesday DayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *DayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DayOfWeek) Set(v string) error {
	switch v {
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = DayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Type always returns DayOfWeek to satisfy [pflag.Value] interface
func (f *DayOfWeek) Type() string {
	return "DayOfWeek"
}

// Delete a pipeline
type DeletePipelineRequest struct {
	PipelineId string `json:"-" url:"-"`
}

type DeletePipelineResponse struct {
}

// The deployment method that manages the pipeline: - BUNDLE: The pipeline is
// managed by a Databricks Asset Bundle.
type DeploymentKind string

const DeploymentKindBundle DeploymentKind = `BUNDLE`

// String representation for [fmt.Print]
func (f *DeploymentKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentKind) Set(v string) error {
	switch v {
	case `BUNDLE`:
		*f = DeploymentKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUNDLE"`, v)
	}
}

// Type always returns DeploymentKind to satisfy [pflag.Value] interface
func (f *DeploymentKind) Type() string {
	return "DeploymentKind"
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// Budget policy of this pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified int64 `json:"expected_last_modified,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Unique identifier for this pipeline.
	PipelineId string `json:"pipeline_id,omitempty" url:"-"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *RunAs `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditPipeline) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditPipeline) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EditPipelineResponse struct {
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions []SerializedException `json:"exceptions,omitempty"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal bool `json:"fatal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ErrorDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ErrorDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The severity level of the event.
type EventLevel string

const EventLevelError EventLevel = `ERROR`

const EventLevelInfo EventLevel = `INFO`

const EventLevelMetrics EventLevel = `METRICS`

const EventLevelWarn EventLevel = `WARN`

// String representation for [fmt.Print]
func (f *EventLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EventLevel) Set(v string) error {
	switch v {
	case `ERROR`, `INFO`, `METRICS`, `WARN`:
		*f = EventLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "INFO", "METRICS", "WARN"`, v)
	}
}

// Type always returns EventLevel to satisfy [pflag.Value] interface
func (f *EventLevel) Type() string {
	return "EventLevel"
}

type FileLibrary struct {
	// The absolute path of the file.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FileLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Filters struct {
	// Paths to exclude.
	Exclude []string `json:"exclude,omitempty"`
	// Paths to include.
	Include []string `json:"include,omitempty"`
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []PipelinePermissionsDescription `json:"permission_levels,omitempty"`
}

// Get pipeline permissions
type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

// Get a pipeline
type GetPipelineRequest struct {
	PipelineId string `json:"-" url:"-"`
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause string `json:"cause,omitempty"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Serverless budget policy ID of this pipeline.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// The health of a pipeline.
	Health GetPipelineResponseHealth `json:"health,omitempty"`
	// The last time the pipeline settings were modified or created.
	LastModified int64 `json:"last_modified,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name string `json:"name,omitempty"`
	// The ID of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec *PipelineSpec `json:"spec,omitempty"`
	// The pipeline state.
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The health of a pipeline.
type GetPipelineResponseHealth string

const GetPipelineResponseHealthHealthy GetPipelineResponseHealth = `HEALTHY`

const GetPipelineResponseHealthUnhealthy GetPipelineResponseHealth = `UNHEALTHY`

// String representation for [fmt.Print]
func (f *GetPipelineResponseHealth) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetPipelineResponseHealth) Set(v string) error {
	switch v {
	case `HEALTHY`, `UNHEALTHY`:
		*f = GetPipelineResponseHealth(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HEALTHY", "UNHEALTHY"`, v)
	}
}

// Type always returns GetPipelineResponseHealth to satisfy [pflag.Value] interface
func (f *GetPipelineResponseHealth) Type() string {
	return "GetPipelineResponseHealth"
}

// Get a pipeline update
type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" url:"-"`
	// The ID of the update.
	UpdateId string `json:"-" url:"-"`
}

type GetUpdateResponse struct {
	// The current update info.
	Update *UpdateInfo `json:"update,omitempty"`
}

type IngestionConfig struct {
	// Select a specific source report.
	Report *ReportSpec `json:"report,omitempty"`
	// Select all tables from a specific source schema.
	Schema *SchemaSpec `json:"schema,omitempty"`
	// Select a specific source table.
	Table *TableSpec `json:"table,omitempty"`
}

type IngestionGatewayPipelineDefinition struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	ConnectionId string `json:"connection_id,omitempty"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName string `json:"connection_name,omitempty"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog string `json:"gateway_storage_catalog,omitempty"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName string `json:"gateway_storage_name,omitempty"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema string `json:"gateway_storage_schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionGatewayPipelineDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionGatewayPipelineDefinition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection that this ingestion pipeline uses
	// to communicate with the source. This is used with connectors for
	// applications like Salesforce, Workday, and so on.
	ConnectionName string `json:"connection_name,omitempty"`
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	IngestionGatewayId string `json:"ingestion_gateway_id,omitempty"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects []IngestionConfig `json:"objects,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IngestionPipelineDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List pipeline events
type ListPipelineEventsRequest struct {
	// Criteria to select a subset of results, expressed using a SQL-like
	// syntax. The supported filters are: 1. level='INFO' (or WARN or ERROR) 2.
	// level in ('INFO', 'WARN') 3. id='[event-id]' 4. timestamp > 'TIMESTAMP'
	// (or >=,<,<=,=)
	//
	// Composite expressions are supported, for example: level in ('ERROR',
	// 'WARN') AND timestamp> '2021-07-22T06:37:33.083Z'
	Filter string `json:"-" url:"filter,omitempty"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken string `json:"-" url:"page_token,omitempty"`

	PipelineId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelineEventsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelineEventsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	Events []PipelineEvent `json:"events,omitempty"`
	// If present, a token to fetch the next page of events.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelineEventsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelineEventsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List pipelines
type ListPipelinesRequest struct {
	// Select a subset of results based on the specified criteria. The supported
	// filters are:
	//
	// * `notebook='<path>'` to select pipelines that reference the provided
	// notebook path. * `name LIKE '[pattern]'` to select pipelines with a name
	// that matches pattern. Wildcards are supported, for example: `name LIKE
	// '%shopping%'`
	//
	// Composite filters are not supported. This field is optional.
	Filter string `json:"-" url:"filter,omitempty"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Page token returned by previous call
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelinesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelinesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of events matching the request criteria.
	Statuses []PipelineStateInfo `json:"statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPipelinesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelinesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List pipeline updates
type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Page token returned by previous call
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The pipeline to return updates for.
	PipelineId string `json:"-" url:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId string `json:"-" url:"until_update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUpdatesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUpdatesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	Updates []UpdateInfo `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListUpdatesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUpdatesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ManualTrigger struct {
}

// Maturity level for EventDetails.
type MaturityLevel string

const MaturityLevelDeprecated MaturityLevel = `DEPRECATED`

const MaturityLevelEvolving MaturityLevel = `EVOLVING`

const MaturityLevelStable MaturityLevel = `STABLE`

// String representation for [fmt.Print]
func (f *MaturityLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *MaturityLevel) Set(v string) error {
	switch v {
	case `DEPRECATED`, `EVOLVING`, `STABLE`:
		*f = MaturityLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPRECATED", "EVOLVING", "STABLE"`, v)
	}
}

// Type always returns MaturityLevel to satisfy [pflag.Value] interface
func (f *MaturityLevel) Type() string {
	return "MaturityLevel"
}

type NotebookLibrary struct {
	// The absolute path of the notebook.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotebookLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Notifications struct {
	// A list of alerts that trigger the sending of notifications to the
	// configured destinations. The supported alerts are:
	//
	// * `on-update-success`: A pipeline update completes successfully. *
	// `on-update-failure`: Each time a pipeline update fails. *
	// `on-update-fatal-failure`: A pipeline update fails with a non-retryable
	// (fatal) error. * `on-flow-failure`: A single data flow fails.
	Alerts []string `json:"alerts,omitempty"`
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients []string `json:"email_recipients,omitempty"`
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	BatchId int `json:"batch_id,omitempty"`
	// The cloud provider, e.g., AWS or Azure.
	Cloud string `json:"cloud,omitempty"`
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId string `json:"cluster_id,omitempty"`
	// The name of a dataset. Unique within a pipeline.
	DatasetName string `json:"dataset_name,omitempty"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId string `json:"flow_id,omitempty"`
	// The name of the flow. Not unique.
	FlowName string `json:"flow_name,omitempty"`
	// The optional host name where the event was triggered
	Host string `json:"host,omitempty"`
	// The id of a maintenance run. Globally unique.
	MaintenanceId string `json:"maintenance_id,omitempty"`
	// Materialization name.
	MaterializationName string `json:"materialization_name,omitempty"`
	// The org id of the user. Unique within a cloud.
	OrgId int `json:"org_id,omitempty"`
	// The id of the pipeline. Globally unique.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The name of the pipeline. Not unique.
	PipelineName string `json:"pipeline_name,omitempty"`
	// The cloud region.
	Region string `json:"region,omitempty"`
	// The id of the request that caused an update.
	RequestId string `json:"request_id,omitempty"`
	// The id of a (delta) table. Globally unique.
	TableId string `json:"table_id,omitempty"`
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId string `json:"uc_resource_id,omitempty"`
	// The id of an execution. Globally unique.
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Origin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Origin) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineAccessControlResponse struct {
	// All permissions.
	AllPermissions []PipelinePermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *PipelineClusterAutoscale `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *compute.AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *compute.AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *compute.ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Whether to enable local disk encryption for the cluster.
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *compute.GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []compute.InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	Label string `json:"label,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	MaxWorkers int `json:"max_workers"`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers int `json:"min_workers"`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode PipelineClusterAutoscaleMode `json:"mode,omitempty"`
}

// Databricks Enhanced Autoscaling optimizes cluster utilization by
// automatically allocating cluster resources based on workload volume, with
// minimal impact to the data processing latency of your pipelines. Enhanced
// Autoscaling is available for `updates` clusters only. The legacy autoscaling
// feature is used for `maintenance` clusters.
type PipelineClusterAutoscaleMode string

const PipelineClusterAutoscaleModeEnhanced PipelineClusterAutoscaleMode = `ENHANCED`

const PipelineClusterAutoscaleModeLegacy PipelineClusterAutoscaleMode = `LEGACY`

// String representation for [fmt.Print]
func (f *PipelineClusterAutoscaleMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineClusterAutoscaleMode) Set(v string) error {
	switch v {
	case `ENHANCED`, `LEGACY`:
		*f = PipelineClusterAutoscaleMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ENHANCED", "LEGACY"`, v)
	}
}

// Type always returns PipelineClusterAutoscaleMode to satisfy [pflag.Value] interface
func (f *PipelineClusterAutoscaleMode) Type() string {
	return "PipelineClusterAutoscaleMode"
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	Kind DeploymentKind `json:"kind,omitempty"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	Error *ErrorDetail `json:"error,omitempty"`
	// The event type. Should always correspond to the details
	EventType string `json:"event_type,omitempty"`
	// A time-based, globally unique id.
	Id string `json:"id,omitempty"`
	// The severity level of the event.
	Level EventLevel `json:"level,omitempty"`
	// Maturity level for event_type.
	MaturityLevel MaturityLevel `json:"maturity_level,omitempty"`
	// The display message associated with the event.
	Message string `json:"message,omitempty"`
	// Describes where the event originates from.
	Origin *Origin `json:"origin,omitempty"`
	// A sequencing object to identify and order events.
	Sequence *Sequencing `json:"sequence,omitempty"`
	// The time of the event.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineEvent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineEvent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File *FileLibrary `json:"file,omitempty"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed.
	Maven *compute.MavenLibrary `json:"maven,omitempty"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook *NotebookLibrary `json:"notebook,omitempty"`
	// URI of the whl to be installed.
	Whl string `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelinePermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type PipelinePermissionLevel string

const PipelinePermissionLevelCanManage PipelinePermissionLevel = `CAN_MANAGE`

const PipelinePermissionLevelCanRun PipelinePermissionLevel = `CAN_RUN`

const PipelinePermissionLevelCanView PipelinePermissionLevel = `CAN_VIEW`

const PipelinePermissionLevelIsOwner PipelinePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *PipelinePermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelinePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = PipelinePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Type always returns PipelinePermissionLevel to satisfy [pflag.Value] interface
func (f *PipelinePermissionLevel) Type() string {
	return "PipelinePermissionLevel"
}

type PipelinePermissions struct {
	AccessControlList []PipelineAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelinePermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelinePermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelinePermissionsRequest struct {
	AccessControlList []PipelineAccessControlRequest `json:"access_control_list,omitempty"`
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

type PipelineSpec struct {
	// Budget policy of this pipeline.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'target' or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []Notifications `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// The default schema (database) where tables are read from or published to.
	// The presence of this field implies that the pipeline is in direct
	// publishing mode.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to. If not
	// specified, no data is published to the Hive metastore or Unity Catalog.
	// To publish to Unity Catalog, also specify `catalog`.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The pipeline state.
type PipelineState string

const PipelineStateDeleted PipelineState = `DELETED`

const PipelineStateDeploying PipelineState = `DEPLOYING`

const PipelineStateFailed PipelineState = `FAILED`

const PipelineStateIdle PipelineState = `IDLE`

const PipelineStateRecovering PipelineState = `RECOVERING`

const PipelineStateResetting PipelineState = `RESETTING`

const PipelineStateRunning PipelineState = `RUNNING`

const PipelineStateStarting PipelineState = `STARTING`

const PipelineStateStopping PipelineState = `STOPPING`

// String representation for [fmt.Print]
func (f *PipelineState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineState) Set(v string) error {
	switch v {
	case `DELETED`, `DEPLOYING`, `FAILED`, `IDLE`, `RECOVERING`, `RESETTING`, `RUNNING`, `STARTING`, `STOPPING`:
		*f = PipelineState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "DEPLOYING", "FAILED", "IDLE", "RECOVERING", "RESETTING", "RUNNING", "STARTING", "STOPPING"`, v)
	}
}

// Type always returns PipelineState to satisfy [pflag.Value] interface
func (f *PipelineState) Type() string {
	return "PipelineState"
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	Health PipelineStateInfoHealth `json:"health,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// The user-friendly name of the pipeline.
	Name string `json:"name,omitempty"`
	// The unique identifier of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline state.
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PipelineStateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineStateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The health of a pipeline.
type PipelineStateInfoHealth string

const PipelineStateInfoHealthHealthy PipelineStateInfoHealth = `HEALTHY`

const PipelineStateInfoHealthUnhealthy PipelineStateInfoHealth = `UNHEALTHY`

// String representation for [fmt.Print]
func (f *PipelineStateInfoHealth) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PipelineStateInfoHealth) Set(v string) error {
	switch v {
	case `HEALTHY`, `UNHEALTHY`:
		*f = PipelineStateInfoHealth(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "HEALTHY", "UNHEALTHY"`, v)
	}
}

// Type always returns PipelineStateInfoHealth to satisfy [pflag.Value] interface
func (f *PipelineStateInfoHealth) Type() string {
	return "PipelineStateInfoHealth"
}

type PipelineTrigger struct {
	Cron *CronTrigger `json:"cron,omitempty"`

	Manual *ManualTrigger `json:"manual,omitempty"`
}

type ReportSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog string `json:"destination_catalog,omitempty"`
	// Required. Destination schema to store table.
	DestinationSchema string `json:"destination_schema,omitempty"`
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	DestinationTable string `json:"destination_table,omitempty"`
	// Required. Report URL in the source system.
	SourceUrl string `json:"source_url,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ReportSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReportSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestartWindow struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	DaysOfWeek []DayOfWeek `json:"days_of_week,omitempty"`
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	StartHour int `json:"start_hour"`
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId string `json:"time_zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestartWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestartWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Write-only setting, available only in Create/Update calls. Specifies the user
// or service principal that the pipeline runs as. If not specified, the
// pipeline runs as the user who created the pipeline.
//
// Only `user_name` or `service_principal_name` can be specified. If both are
// specified, an error is thrown.
type RunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RunAs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunAs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	DestinationCatalog string `json:"destination_catalog,omitempty"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema string `json:"destination_schema,omitempty"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Required. Schema name in the source database.
	SourceSchema string `json:"source_schema,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SchemaSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Sequencing struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo int `json:"control_plane_seq_no,omitempty"`
	// the ID assigned by the data plane.
	DataPlaneId *DataPlaneId `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Sequencing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Sequencing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SerializedException struct {
	// Runtime class of the exception
	ClassName string `json:"class_name,omitempty"`
	// Exception message
	Message string `json:"message,omitempty"`
	// Stack trace consisting of a list of stack frames
	Stack []StackFrame `json:"stack,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SerializedException) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SerializedException) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StackFrame struct {
	// Class from which the method call originated
	DeclaringClass string `json:"declaring_class,omitempty"`
	// File where the method is defined
	FileName string `json:"file_name,omitempty"`
	// Line from which the method was called
	LineNumber int `json:"line_number,omitempty"`
	// Name of the method which was called
	MethodName string `json:"method_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StackFrame) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StackFrame) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StartUpdate struct {
	Cause StartUpdateCause `json:"cause,omitempty"`
	// If true, this update will reset all tables before running.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`

	PipelineId string `json:"-" url:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StartUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StartUpdate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StartUpdateCause string

const StartUpdateCauseApiCall StartUpdateCause = `API_CALL`

const StartUpdateCauseJobTask StartUpdateCause = `JOB_TASK`

const StartUpdateCauseRetryOnFailure StartUpdateCause = `RETRY_ON_FAILURE`

const StartUpdateCauseSchemaChange StartUpdateCause = `SCHEMA_CHANGE`

const StartUpdateCauseServiceUpgrade StartUpdateCause = `SERVICE_UPGRADE`

const StartUpdateCauseUserAction StartUpdateCause = `USER_ACTION`

// String representation for [fmt.Print]
func (f *StartUpdateCause) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StartUpdateCause) Set(v string) error {
	switch v {
	case `API_CALL`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = StartUpdateCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Type always returns StartUpdateCause to satisfy [pflag.Value] interface
func (f *StartUpdateCause) Type() string {
	return "StartUpdateCause"
}

type StartUpdateResponse struct {
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StartUpdateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StartUpdateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StopPipelineResponse struct {
}

// Stop a pipeline
type StopRequest struct {
	PipelineId string `json:"-" url:"-"`
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog string `json:"destination_catalog,omitempty"`
	// Required. Destination schema to store table.
	DestinationSchema string `json:"destination_schema,omitempty"`
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable string `json:"destination_table,omitempty"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema string `json:"source_schema,omitempty"`
	// Required. Table name in the source database.
	SourceTable string `json:"source_table,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableSpecificConfig struct {
	// The primary key of the table used to apply changes.
	PrimaryKeys []string `json:"primary_keys,omitempty"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields bool `json:"salesforce_include_formula_fields,omitempty"`
	// The SCD type to use to ingest the table.
	ScdType TableSpecificConfigScdType `json:"scd_type,omitempty"`
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	SequenceBy []string `json:"sequence_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TableSpecificConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSpecificConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The SCD type to use to ingest the table.
type TableSpecificConfigScdType string

const TableSpecificConfigScdTypeScdType1 TableSpecificConfigScdType = `SCD_TYPE_1`

const TableSpecificConfigScdTypeScdType2 TableSpecificConfigScdType = `SCD_TYPE_2`

// String representation for [fmt.Print]
func (f *TableSpecificConfigScdType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableSpecificConfigScdType) Set(v string) error {
	switch v {
	case `SCD_TYPE_1`, `SCD_TYPE_2`:
		*f = TableSpecificConfigScdType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "SCD_TYPE_1", "SCD_TYPE_2"`, v)
	}
}

// Type always returns TableSpecificConfigScdType to satisfy [pflag.Value] interface
func (f *TableSpecificConfigScdType) Type() string {
	return "TableSpecificConfigScdType"
}

type UpdateInfo struct {
	// What triggered this update.
	Cause UpdateInfoCause `json:"cause,omitempty"`
	// The ID of the cluster that the update is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config *PipelineSpec `json:"config,omitempty"`
	// The time when this update was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// If true, this update will reset all tables before running.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`
	// The ID of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// The update state.
	State UpdateInfoState `json:"state,omitempty"`
	// The ID of this update.
	UpdateId string `json:"update_id,omitempty"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// What triggered this update.
type UpdateInfoCause string

const UpdateInfoCauseApiCall UpdateInfoCause = `API_CALL`

const UpdateInfoCauseJobTask UpdateInfoCause = `JOB_TASK`

const UpdateInfoCauseRetryOnFailure UpdateInfoCause = `RETRY_ON_FAILURE`

const UpdateInfoCauseSchemaChange UpdateInfoCause = `SCHEMA_CHANGE`

const UpdateInfoCauseServiceUpgrade UpdateInfoCause = `SERVICE_UPGRADE`

const UpdateInfoCauseUserAction UpdateInfoCause = `USER_ACTION`

// String representation for [fmt.Print]
func (f *UpdateInfoCause) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateInfoCause) Set(v string) error {
	switch v {
	case `API_CALL`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = UpdateInfoCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Type always returns UpdateInfoCause to satisfy [pflag.Value] interface
func (f *UpdateInfoCause) Type() string {
	return "UpdateInfoCause"
}

// The update state.
type UpdateInfoState string

const UpdateInfoStateCanceled UpdateInfoState = `CANCELED`

const UpdateInfoStateCompleted UpdateInfoState = `COMPLETED`

const UpdateInfoStateCreated UpdateInfoState = `CREATED`

const UpdateInfoStateFailed UpdateInfoState = `FAILED`

const UpdateInfoStateInitializing UpdateInfoState = `INITIALIZING`

const UpdateInfoStateQueued UpdateInfoState = `QUEUED`

const UpdateInfoStateResetting UpdateInfoState = `RESETTING`

const UpdateInfoStateRunning UpdateInfoState = `RUNNING`

const UpdateInfoStateSettingUpTables UpdateInfoState = `SETTING_UP_TABLES`

const UpdateInfoStateStopping UpdateInfoState = `STOPPING`

const UpdateInfoStateWaitingForResources UpdateInfoState = `WAITING_FOR_RESOURCES`

// String representation for [fmt.Print]
func (f *UpdateInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `COMPLETED`, `CREATED`, `FAILED`, `INITIALIZING`, `QUEUED`, `RESETTING`, `RUNNING`, `SETTING_UP_TABLES`, `STOPPING`, `WAITING_FOR_RESOURCES`:
		*f = UpdateInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "COMPLETED", "CREATED", "FAILED", "INITIALIZING", "QUEUED", "RESETTING", "RUNNING", "SETTING_UP_TABLES", "STOPPING", "WAITING_FOR_RESOURCES"`, v)
	}
}

// Type always returns UpdateInfoState to satisfy [pflag.Value] interface
func (f *UpdateInfoState) Type() string {
	return "UpdateInfoState"
}

type UpdateStateInfo struct {
	CreationTime string `json:"creation_time,omitempty"`

	State UpdateStateInfoState `json:"state,omitempty"`

	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateStateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateStateInfoState string

const UpdateStateInfoStateCanceled UpdateStateInfoState = `CANCELED`

const UpdateStateInfoStateCompleted UpdateStateInfoState = `COMPLETED`

const UpdateStateInfoStateCreated UpdateStateInfoState = `CREATED`

const UpdateStateInfoStateFailed UpdateStateInfoState = `FAILED`

const UpdateStateInfoStateInitializing UpdateStateInfoState = `INITIALIZING`

const UpdateStateInfoStateQueued UpdateStateInfoState = `QUEUED`

const UpdateStateInfoStateResetting UpdateStateInfoState = `RESETTING`

const UpdateStateInfoStateRunning UpdateStateInfoState = `RUNNING`

const UpdateStateInfoStateSettingUpTables UpdateStateInfoState = `SETTING_UP_TABLES`

const UpdateStateInfoStateStopping UpdateStateInfoState = `STOPPING`

const UpdateStateInfoStateWaitingForResources UpdateStateInfoState = `WAITING_FOR_RESOURCES`

// String representation for [fmt.Print]
func (f *UpdateStateInfoState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateStateInfoState) Set(v string) error {
	switch v {
	case `CANCELED`, `COMPLETED`, `CREATED`, `FAILED`, `INITIALIZING`, `QUEUED`, `RESETTING`, `RUNNING`, `SETTING_UP_TABLES`, `STOPPING`, `WAITING_FOR_RESOURCES`:
		*f = UpdateStateInfoState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "COMPLETED", "CREATED", "FAILED", "INITIALIZING", "QUEUED", "RESETTING", "RUNNING", "SETTING_UP_TABLES", "STOPPING", "WAITING_FOR_RESOURCES"`, v)
	}
}

// Type always returns UpdateStateInfoState to satisfy [pflag.Value] interface
func (f *UpdateStateInfoState) Type() string {
	return "UpdateStateInfoState"
}
