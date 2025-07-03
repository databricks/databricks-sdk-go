// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/google/go-querystring/query"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	// Wire name: 'allow_duplicate_names'
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool `json:"development,omitempty"`

	// Wire name: 'dry_run'
	DryRun bool `json:"dry_run,omitempty"`
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	// Wire name: 'environment'
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	// Wire name: 'root_path'
	RootPath string `json:"root_path,omitempty"`

	// Wire name: 'run_as'
	RunAs *RunAs `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	// Wire name: 'tags'
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger *PipelineTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePipeline) EncodeValues(key string, v *url.Values) error {
	pb, err := createPipelineToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreatePipeline) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPipelinePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPipelineFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePipeline) MarshalJSON() ([]byte, error) {
	pb, err := createPipelineToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	// Wire name: 'effective_settings'
	EffectiveSettings *PipelineSpec `json:"effective_settings,omitempty"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePipelineResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createPipelineResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreatePipelineResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPipelineResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPipelineResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePipelineResponse) MarshalJSON() ([]byte, error) {
	pb, err := createPipelineResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CronTrigger struct {

	// Wire name: 'quartz_cron_schedule'
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`

	// Wire name: 'timezone_id'
	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CronTrigger) EncodeValues(key string, v *url.Values) error {
	pb, err := cronTriggerToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CronTrigger) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cronTriggerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cronTriggerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CronTrigger) MarshalJSON() ([]byte, error) {
	pb, err := cronTriggerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	// Wire name: 'instance'
	Instance string `json:"instance,omitempty"`
	// A sequence number, unique and increasing within the data plane instance.
	// Wire name: 'seq_no'
	SeqNo int64 `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DataPlaneId) EncodeValues(key string, v *url.Values) error {
	pb, err := dataPlaneIdToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DataPlaneId) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dataPlaneIdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dataPlaneIdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DataPlaneId) MarshalJSON() ([]byte, error) {
	pb, err := dataPlaneIdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for DayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// Type always returns DayOfWeek to satisfy [pflag.Value] interface
func (f *DayOfWeek) Type() string {
	return "DayOfWeek"
}

type DeletePipelineRequest struct {
	PipelineId string `json:"-" tf:"-"`
}

func (st *DeletePipelineRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deletePipelineRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeletePipelineRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePipelineRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePipelineRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePipelineRequest) MarshalJSON() ([]byte, error) {
	pb, err := deletePipelineRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for DeploymentKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeploymentKind) Values() []DeploymentKind {
	return []DeploymentKind{
		DeploymentKindBundle,
	}
}

// Type always returns DeploymentKind to satisfy [pflag.Value] interface
func (f *DeploymentKind) Type() string {
	return "DeploymentKind"
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	// Wire name: 'allow_duplicate_names'
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	// Wire name: 'environment'
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	// Wire name: 'expected_last_modified'
	ExpectedLastModified int64 `json:"expected_last_modified,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool `json:"photon,omitempty"`
	// Unique identifier for this pipeline.
	PipelineId string `json:"-" tf:"-"`
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	// Wire name: 'root_path'
	RootPath string `json:"root_path,omitempty"`

	// Wire name: 'run_as'
	RunAs *RunAs `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	// Wire name: 'tags'
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger *PipelineTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EditPipeline) EncodeValues(key string, v *url.Values) error {
	pb, err := editPipelineToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EditPipeline) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editPipelinePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editPipelineFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditPipeline) MarshalJSON() ([]byte, error) {
	pb, err := editPipelineToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	// Wire name: 'exceptions'
	Exceptions []SerializedException `json:"exceptions,omitempty"`
	// Whether this error is considered fatal, that is, unrecoverable.
	// Wire name: 'fatal'
	Fatal bool `json:"fatal,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ErrorDetail) EncodeValues(key string, v *url.Values) error {
	pb, err := errorDetailToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ErrorDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &errorDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := errorDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ErrorDetail) MarshalJSON() ([]byte, error) {
	pb, err := errorDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for EventLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *EventLevel) Values() []EventLevel {
	return []EventLevel{
		EventLevelError,
		EventLevelInfo,
		EventLevelMetrics,
		EventLevelWarn,
	}
}

// Type always returns EventLevel to satisfy [pflag.Value] interface
func (f *EventLevel) Type() string {
	return "EventLevel"
}

// Configurable event log parameters.
type EventLogSpec struct {
	// The UC catalog the event log is published under.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// The name the event log is published to in UC.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The UC schema the event log is published under.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EventLogSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := eventLogSpecToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EventLogSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &eventLogSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := eventLogSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EventLogSpec) MarshalJSON() ([]byte, error) {
	pb, err := eventLogSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FileLibrary struct {
	// The absolute path of the source code.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *FileLibrary) EncodeValues(key string, v *url.Values) error {
	pb, err := fileLibraryToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *FileLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileLibrary) MarshalJSON() ([]byte, error) {
	pb, err := fileLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Filters struct {
	// Paths to exclude.
	// Wire name: 'exclude'
	Exclude []string `json:"exclude,omitempty"`
	// Paths to include.
	// Wire name: 'include'
	Include []string `json:"include,omitempty"`
}

func (st *Filters) EncodeValues(key string, v *url.Values) error {
	pb, err := filtersToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Filters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filtersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := filtersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Filters) MarshalJSON() ([]byte, error) {
	pb, err := filtersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" tf:"-"`
}

func (st *GetPipelinePermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPipelinePermissionLevelsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPipelinePermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPipelinePermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPipelinePermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPipelinePermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPipelinePermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []PipelinePermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetPipelinePermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getPipelinePermissionLevelsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPipelinePermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPipelinePermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPipelinePermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPipelinePermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPipelinePermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" tf:"-"`
}

func (st *GetPipelinePermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPipelinePermissionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPipelinePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPipelinePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPipelinePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPipelinePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPipelinePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPipelineRequest struct {
	PipelineId string `json:"-" tf:"-"`
}

func (st *GetPipelineRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPipelineRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPipelineRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPipelineRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPipelineRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPipelineRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPipelineRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	// Wire name: 'cause'
	Cause string `json:"cause,omitempty"`
	// The ID of the cluster that the pipeline is running on.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	// Wire name: 'creator_user_name'
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Serverless budget policy ID of this pipeline.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// The health of a pipeline.
	// Wire name: 'health'
	Health GetPipelineResponseHealth `json:"health,omitempty"`
	// The last time the pipeline settings were modified or created.
	// Wire name: 'last_modified'
	LastModified int64 `json:"last_modified,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	// Wire name: 'latest_updates'
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// The user or service principal that the pipeline runs as, if specified in
	// the request. This field indicates the explicit configuration of `run_as`
	// for the pipeline. To find the value in all cases, explicit or implicit,
	// use `run_as_user_name`.
	// Wire name: 'run_as'
	RunAs *RunAs `json:"run_as,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	// Wire name: 'run_as_user_name'
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	// Wire name: 'spec'
	Spec *PipelineSpec `json:"spec,omitempty"`
	// The pipeline state.
	// Wire name: 'state'
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetPipelineResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getPipelineResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetPipelineResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPipelineResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPipelineResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPipelineResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPipelineResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for GetPipelineResponseHealth.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetPipelineResponseHealth) Values() []GetPipelineResponseHealth {
	return []GetPipelineResponseHealth{
		GetPipelineResponseHealthHealthy,
		GetPipelineResponseHealthUnhealthy,
	}
}

// Type always returns GetPipelineResponseHealth to satisfy [pflag.Value] interface
func (f *GetPipelineResponseHealth) Type() string {
	return "GetPipelineResponseHealth"
}

type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" tf:"-"`
	// The ID of the update.
	UpdateId string `json:"-" tf:"-"`
}

func (st *GetUpdateRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getUpdateRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetUpdateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getUpdateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getUpdateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetUpdateRequest) MarshalJSON() ([]byte, error) {
	pb, err := getUpdateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetUpdateResponse struct {
	// The current update info.
	// Wire name: 'update'
	Update *UpdateInfo `json:"update,omitempty"`
}

func (st *GetUpdateResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getUpdateResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetUpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getUpdateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getUpdateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetUpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := getUpdateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type IngestionConfig struct {
	// Select a specific source report.
	// Wire name: 'report'
	Report *ReportSpec `json:"report,omitempty"`
	// Select all tables from a specific source schema.
	// Wire name: 'schema'
	Schema *SchemaSpec `json:"schema,omitempty"`
	// Select a specific source table.
	// Wire name: 'table'
	Table *TableSpec `json:"table,omitempty"`
}

func (st *IngestionConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := ingestionConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *IngestionConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ingestionConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ingestionConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st IngestionConfig) MarshalJSON() ([]byte, error) {
	pb, err := ingestionConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type IngestionGatewayPipelineDefinition struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	// Wire name: 'connection_id'
	ConnectionId string `json:"connection_id,omitempty"`
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	// Wire name: 'connection_name'
	ConnectionName string `json:"connection_name"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	// Wire name: 'gateway_storage_catalog'
	GatewayStorageCatalog string `json:"gateway_storage_catalog"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	// Wire name: 'gateway_storage_name'
	GatewayStorageName string `json:"gateway_storage_name,omitempty"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	// Wire name: 'gateway_storage_schema'
	GatewayStorageSchema string `json:"gateway_storage_schema"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *IngestionGatewayPipelineDefinition) EncodeValues(key string, v *url.Values) error {
	pb, err := ingestionGatewayPipelineDefinitionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *IngestionGatewayPipelineDefinition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ingestionGatewayPipelineDefinitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ingestionGatewayPipelineDefinitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st IngestionGatewayPipelineDefinition) MarshalJSON() ([]byte, error) {
	pb, err := ingestionGatewayPipelineDefinitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection that this ingestion pipeline uses
	// to communicate with the source. This is used with connectors for
	// applications like Salesforce, Workday, and so on.
	// Wire name: 'connection_name'
	ConnectionName string `json:"connection_name,omitempty"`
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	// Wire name: 'ingestion_gateway_id'
	IngestionGatewayId string `json:"ingestion_gateway_id,omitempty"`
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	// Wire name: 'objects'
	Objects []IngestionConfig `json:"objects,omitempty"`
	// The type of the foreign source. The source type will be inferred from the
	// source connection or ingestion gateway. This field is output only and
	// will be ignored if provided.
	// Wire name: 'source_type'
	SourceType IngestionSourceType `json:"source_type,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *IngestionPipelineDefinition) EncodeValues(key string, v *url.Values) error {
	pb, err := ingestionPipelineDefinitionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *IngestionPipelineDefinition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ingestionPipelineDefinitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ingestionPipelineDefinitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st IngestionPipelineDefinition) MarshalJSON() ([]byte, error) {
	pb, err := ingestionPipelineDefinitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type IngestionSourceType string

const IngestionSourceTypeBigquery IngestionSourceType = `BIGQUERY`

const IngestionSourceTypeDynamics365 IngestionSourceType = `DYNAMICS365`

const IngestionSourceTypeGa4RawData IngestionSourceType = `GA4_RAW_DATA`

const IngestionSourceTypeManagedPostgresql IngestionSourceType = `MANAGED_POSTGRESQL`

const IngestionSourceTypeMysql IngestionSourceType = `MYSQL`

const IngestionSourceTypeNetsuite IngestionSourceType = `NETSUITE`

const IngestionSourceTypeOracle IngestionSourceType = `ORACLE`

const IngestionSourceTypePostgresql IngestionSourceType = `POSTGRESQL`

const IngestionSourceTypeSalesforce IngestionSourceType = `SALESFORCE`

const IngestionSourceTypeServicenow IngestionSourceType = `SERVICENOW`

const IngestionSourceTypeSharepoint IngestionSourceType = `SHAREPOINT`

const IngestionSourceTypeSqlserver IngestionSourceType = `SQLSERVER`

const IngestionSourceTypeTeradata IngestionSourceType = `TERADATA`

const IngestionSourceTypeWorkdayRaas IngestionSourceType = `WORKDAY_RAAS`

// String representation for [fmt.Print]
func (f *IngestionSourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *IngestionSourceType) Set(v string) error {
	switch v {
	case `BIGQUERY`, `DYNAMICS365`, `GA4_RAW_DATA`, `MANAGED_POSTGRESQL`, `MYSQL`, `NETSUITE`, `ORACLE`, `POSTGRESQL`, `SALESFORCE`, `SERVICENOW`, `SHAREPOINT`, `SQLSERVER`, `TERADATA`, `WORKDAY_RAAS`:
		*f = IngestionSourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "DYNAMICS365", "GA4_RAW_DATA", "MANAGED_POSTGRESQL", "MYSQL", "NETSUITE", "ORACLE", "POSTGRESQL", "SALESFORCE", "SERVICENOW", "SHAREPOINT", "SQLSERVER", "TERADATA", "WORKDAY_RAAS"`, v)
	}
}

// Values returns all possible values for IngestionSourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *IngestionSourceType) Values() []IngestionSourceType {
	return []IngestionSourceType{
		IngestionSourceTypeBigquery,
		IngestionSourceTypeDynamics365,
		IngestionSourceTypeGa4RawData,
		IngestionSourceTypeManagedPostgresql,
		IngestionSourceTypeMysql,
		IngestionSourceTypeNetsuite,
		IngestionSourceTypeOracle,
		IngestionSourceTypePostgresql,
		IngestionSourceTypeSalesforce,
		IngestionSourceTypeServicenow,
		IngestionSourceTypeSharepoint,
		IngestionSourceTypeSqlserver,
		IngestionSourceTypeTeradata,
		IngestionSourceTypeWorkdayRaas,
	}
}

// Type always returns IngestionSourceType to satisfy [pflag.Value] interface
func (f *IngestionSourceType) Type() string {
	return "IngestionSourceType"
}

type ListPipelineEventsRequest struct {
	// Criteria to select a subset of results, expressed using a SQL-like
	// syntax. The supported filters are: 1. level='INFO' (or WARN or ERROR) 2.
	// level in ('INFO', 'WARN') 3. id='[event-id]' 4. timestamp > 'TIMESTAMP'
	// (or >=,<,<=,=)
	//
	// Composite expressions are supported, for example: level in ('ERROR',
	// 'WARN') AND timestamp> '2021-07-22T06:37:33.083Z'
	Filter string `json:"-" tf:"-"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults int `json:"-" tf:"-"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy []string `json:"-" tf:"-"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken string `json:"-" tf:"-"`
	// The pipeline to return events for.
	PipelineId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPipelineEventsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listPipelineEventsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListPipelineEventsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPipelineEventsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPipelineEventsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPipelineEventsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listPipelineEventsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	// Wire name: 'events'
	Events []PipelineEvent `json:"events,omitempty"`
	// If present, a token to fetch the next page of events.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, a token to fetch the previous page of events.
	// Wire name: 'prev_page_token'
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPipelineEventsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listPipelineEventsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListPipelineEventsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPipelineEventsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPipelineEventsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPipelineEventsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listPipelineEventsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

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
	Filter string `json:"-" tf:"-"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults int `json:"-" tf:"-"`
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	OrderBy []string `json:"-" tf:"-"`
	// Page token returned by previous call
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPipelinesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listPipelinesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListPipelinesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPipelinesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPipelinesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPipelinesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listPipelinesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of events matching the request criteria.
	// Wire name: 'statuses'
	Statuses []PipelineStateInfo `json:"statuses,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPipelinesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listPipelinesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListPipelinesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPipelinesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPipelinesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPipelinesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listPipelinesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults int `json:"-" tf:"-"`
	// Page token returned by previous call
	PageToken string `json:"-" tf:"-"`
	// The pipeline to return updates for.
	PipelineId string `json:"-" tf:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListUpdatesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listUpdatesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListUpdatesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listUpdatesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listUpdatesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListUpdatesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listUpdatesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	// Wire name: 'prev_page_token'
	PrevPageToken string `json:"prev_page_token,omitempty"`

	// Wire name: 'updates'
	Updates []UpdateInfo `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListUpdatesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listUpdatesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListUpdatesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listUpdatesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listUpdatesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListUpdatesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listUpdatesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ManualTrigger struct {
}

func (st *ManualTrigger) EncodeValues(key string, v *url.Values) error {
	pb, err := manualTriggerToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ManualTrigger) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &manualTriggerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := manualTriggerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ManualTrigger) MarshalJSON() ([]byte, error) {
	pb, err := manualTriggerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for MaturityLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *MaturityLevel) Values() []MaturityLevel {
	return []MaturityLevel{
		MaturityLevelDeprecated,
		MaturityLevelEvolving,
		MaturityLevelStable,
	}
}

// Type always returns MaturityLevel to satisfy [pflag.Value] interface
func (f *MaturityLevel) Type() string {
	return "MaturityLevel"
}

type NotebookLibrary struct {
	// The absolute path of the source code.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NotebookLibrary) EncodeValues(key string, v *url.Values) error {
	pb, err := notebookLibraryToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NotebookLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notebookLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notebookLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotebookLibrary) MarshalJSON() ([]byte, error) {
	pb, err := notebookLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Notifications struct {
	// A list of alerts that trigger the sending of notifications to the
	// configured destinations. The supported alerts are:
	//
	// * `on-update-success`: A pipeline update completes successfully. *
	// `on-update-failure`: Each time a pipeline update fails. *
	// `on-update-fatal-failure`: A pipeline update fails with a non-retryable
	// (fatal) error. * `on-flow-failure`: A single data flow fails.
	// Wire name: 'alerts'
	Alerts []string `json:"alerts,omitempty"`
	// A list of email addresses notified when a configured alert is triggered.
	// Wire name: 'email_recipients'
	EmailRecipients []string `json:"email_recipients,omitempty"`
}

func (st *Notifications) EncodeValues(key string, v *url.Values) error {
	pb, err := notificationsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Notifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Notifications) MarshalJSON() ([]byte, error) {
	pb, err := notificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	// Wire name: 'batch_id'
	BatchId int64 `json:"batch_id,omitempty"`
	// The cloud provider, e.g., AWS or Azure.
	// Wire name: 'cloud'
	Cloud string `json:"cloud,omitempty"`
	// The id of the cluster where an execution happens. Unique within a region.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`
	// The name of a dataset. Unique within a pipeline.
	// Wire name: 'dataset_name'
	DatasetName string `json:"dataset_name,omitempty"`
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	// Wire name: 'flow_id'
	FlowId string `json:"flow_id,omitempty"`
	// The name of the flow. Not unique.
	// Wire name: 'flow_name'
	FlowName string `json:"flow_name,omitempty"`
	// The optional host name where the event was triggered
	// Wire name: 'host'
	Host string `json:"host,omitempty"`
	// The id of a maintenance run. Globally unique.
	// Wire name: 'maintenance_id'
	MaintenanceId string `json:"maintenance_id,omitempty"`
	// Materialization name.
	// Wire name: 'materialization_name'
	MaterializationName string `json:"materialization_name,omitempty"`
	// The org id of the user. Unique within a cloud.
	// Wire name: 'org_id'
	OrgId int64 `json:"org_id,omitempty"`
	// The id of the pipeline. Globally unique.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// The name of the pipeline. Not unique.
	// Wire name: 'pipeline_name'
	PipelineName string `json:"pipeline_name,omitempty"`
	// The cloud region.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The id of the request that caused an update.
	// Wire name: 'request_id'
	RequestId string `json:"request_id,omitempty"`
	// The id of a (delta) table. Globally unique.
	// Wire name: 'table_id'
	TableId string `json:"table_id,omitempty"`
	// The Unity Catalog id of the MV or ST being updated.
	// Wire name: 'uc_resource_id'
	UcResourceId string `json:"uc_resource_id,omitempty"`
	// The id of an execution. Globally unique.
	// Wire name: 'update_id'
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Origin) EncodeValues(key string, v *url.Values) error {
	pb, err := originToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Origin) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &originPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := originFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Origin) MarshalJSON() ([]byte, error) {
	pb, err := originToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PathPattern struct {
	// The source code to include for pipelines
	// Wire name: 'include'
	Include string `json:"include,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PathPattern) EncodeValues(key string, v *url.Values) error {
	pb, err := pathPatternToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PathPattern) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pathPatternPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pathPatternFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PathPattern) MarshalJSON() ([]byte, error) {
	pb, err := pathPatternToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineAccessControlRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := pipelineAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []PipelinePermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineAccessControlResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := pipelineAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *PipelineClusterAutoscale `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *compute.AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *compute.AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *compute.ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Whether to enable local disk encryption for the cluster.
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *compute.GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []compute.InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	// Wire name: 'label'
	Label string `json:"label,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	// Wire name: 'spark_conf'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineCluster) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineClusterToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineCluster) MarshalJSON() ([]byte, error) {
	pb, err := pipelineClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	// Wire name: 'max_workers'
	MaxWorkers int `json:"max_workers"`
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	// Wire name: 'min_workers'
	MinWorkers int `json:"min_workers"`
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	// Wire name: 'mode'
	Mode PipelineClusterAutoscaleMode `json:"mode,omitempty"`
}

func (st *PipelineClusterAutoscale) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineClusterAutoscaleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineClusterAutoscale) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineClusterAutoscalePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineClusterAutoscaleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineClusterAutoscale) MarshalJSON() ([]byte, error) {
	pb, err := pipelineClusterAutoscaleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for PipelineClusterAutoscaleMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineClusterAutoscaleMode) Values() []PipelineClusterAutoscaleMode {
	return []PipelineClusterAutoscaleMode{
		PipelineClusterAutoscaleModeEnhanced,
		PipelineClusterAutoscaleModeLegacy,
	}
}

// Type always returns PipelineClusterAutoscaleMode to satisfy [pflag.Value] interface
func (f *PipelineClusterAutoscaleMode) Type() string {
	return "PipelineClusterAutoscaleMode"
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	// Wire name: 'kind'
	Kind DeploymentKind `json:"kind"`
	// The path to the file containing metadata about the deployment.
	// Wire name: 'metadata_file_path'
	MetadataFilePath string `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineDeployment) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineDeploymentToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineDeployment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineDeploymentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineDeploymentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineDeployment) MarshalJSON() ([]byte, error) {
	pb, err := pipelineDeploymentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	// Wire name: 'error'
	Error *ErrorDetail `json:"error,omitempty"`
	// The event type. Should always correspond to the details
	// Wire name: 'event_type'
	EventType string `json:"event_type,omitempty"`
	// A time-based, globally unique id.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The severity level of the event.
	// Wire name: 'level'
	Level EventLevel `json:"level,omitempty"`
	// Maturity level for event_type.
	// Wire name: 'maturity_level'
	MaturityLevel MaturityLevel `json:"maturity_level,omitempty"`
	// The display message associated with the event.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// Describes where the event originates from.
	// Wire name: 'origin'
	Origin *Origin `json:"origin,omitempty"`
	// A sequencing object to identify and order events.
	// Wire name: 'sequence'
	Sequence *Sequencing `json:"sequence,omitempty"`
	// The time of the event.
	// Wire name: 'timestamp'
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineEvent) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineEventToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineEvent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineEventPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineEventFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineEvent) MarshalJSON() ([]byte, error) {
	pb, err := pipelineEventToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	// Wire name: 'file'
	File *FileLibrary `json:"file,omitempty"`
	// The unified field to include source codes. Each entry can be a notebook
	// path, a file path, or a folder path that ends `/**`. This field cannot be
	// used together with `notebook` or `file`.
	// Wire name: 'glob'
	Glob *PathPattern `json:"glob,omitempty"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	// Wire name: 'jar'
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed.
	// Wire name: 'maven'
	Maven *compute.MavenLibrary `json:"maven,omitempty"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	// Wire name: 'notebook'
	Notebook *NotebookLibrary `json:"notebook,omitempty"`
	// URI of the whl to be installed.
	// Wire name: 'whl'
	Whl string `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineLibrary) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineLibraryToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineLibrary) MarshalJSON() ([]byte, error) {
	pb, err := pipelineLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelinePermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelinePermission) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelinePermissionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelinePermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelinePermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelinePermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelinePermission) MarshalJSON() ([]byte, error) {
	pb, err := pipelinePermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for PipelinePermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelinePermissionLevel) Values() []PipelinePermissionLevel {
	return []PipelinePermissionLevel{
		PipelinePermissionLevelCanManage,
		PipelinePermissionLevelCanRun,
		PipelinePermissionLevelCanView,
		PipelinePermissionLevelIsOwner,
	}
}

// Type always returns PipelinePermissionLevel to satisfy [pflag.Value] interface
func (f *PipelinePermissionLevel) Type() string {
	return "PipelinePermissionLevel"
}

type PipelinePermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []PipelineAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelinePermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelinePermissionsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelinePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelinePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelinePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelinePermissions) MarshalJSON() ([]byte, error) {
	pb, err := pipelinePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelinePermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelinePermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelinePermissionsDescriptionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelinePermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelinePermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelinePermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelinePermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := pipelinePermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelinePermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []PipelineAccessControlRequest `json:"access_control_list,omitempty"`
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" tf:"-"`
}

func (st *PipelinePermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelinePermissionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelinePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelinePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelinePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelinePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := pipelinePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineSpec struct {
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string `json:"edition,omitempty"`
	// Environment specification for this pipeline used to install dependencies.
	// Wire name: 'environment'
	Environment *PipelinesEnvironment `json:"environment,omitempty"`
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec `json:"event_log,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow `json:"restart_window,omitempty"`
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	// Wire name: 'root_path'
	RootPath string `json:"root_path,omitempty"`
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string `json:"storage,omitempty"`
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	// Wire name: 'tags'
	Tags map[string]string `json:"tags,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger *PipelineTrigger `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineSpecToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineSpec) MarshalJSON() ([]byte, error) {
	pb, err := pipelineSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for PipelineState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineState) Values() []PipelineState {
	return []PipelineState{
		PipelineStateDeleted,
		PipelineStateDeploying,
		PipelineStateFailed,
		PipelineStateIdle,
		PipelineStateRecovering,
		PipelineStateResetting,
		PipelineStateRunning,
		PipelineStateStarting,
		PipelineStateStopping,
	}
}

// Type always returns PipelineState to satisfy [pflag.Value] interface
func (f *PipelineState) Type() string {
	return "PipelineState"
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	// Wire name: 'creator_user_name'
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	// Wire name: 'health'
	Health PipelineStateInfoHealth `json:"health,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	// Wire name: 'latest_updates'
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// The user-friendly name of the pipeline.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The unique identifier of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	// Wire name: 'run_as_user_name'
	RunAsUserName string `json:"run_as_user_name,omitempty"`

	// Wire name: 'state'
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PipelineStateInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineStateInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineStateInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineStateInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineStateInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineStateInfo) MarshalJSON() ([]byte, error) {
	pb, err := pipelineStateInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for PipelineStateInfoHealth.
//
// There is no guarantee on the order of the values in the slice.
func (f *PipelineStateInfoHealth) Values() []PipelineStateInfoHealth {
	return []PipelineStateInfoHealth{
		PipelineStateInfoHealthHealthy,
		PipelineStateInfoHealthUnhealthy,
	}
}

// Type always returns PipelineStateInfoHealth to satisfy [pflag.Value] interface
func (f *PipelineStateInfoHealth) Type() string {
	return "PipelineStateInfoHealth"
}

type PipelineTrigger struct {

	// Wire name: 'cron'
	Cron *CronTrigger `json:"cron,omitempty"`

	// Wire name: 'manual'
	Manual *ManualTrigger `json:"manual,omitempty"`
}

func (st *PipelineTrigger) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelineTriggerToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelineTrigger) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineTriggerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineTriggerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineTrigger) MarshalJSON() ([]byte, error) {
	pb, err := pipelineTriggerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type PipelinesEnvironment struct {
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url>
	// Wire name: 'dependencies'
	Dependencies []string `json:"dependencies,omitempty"`
}

func (st *PipelinesEnvironment) EncodeValues(key string, v *url.Values) error {
	pb, err := pipelinesEnvironmentToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PipelinesEnvironment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelinesEnvironmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelinesEnvironmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelinesEnvironment) MarshalJSON() ([]byte, error) {
	pb, err := pipelinesEnvironmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ReportSpec struct {
	// Required. Destination catalog to store table.
	// Wire name: 'destination_catalog'
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store table.
	// Wire name: 'destination_schema'
	DestinationSchema string `json:"destination_schema"`
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	// Wire name: 'destination_table'
	DestinationTable string `json:"destination_table,omitempty"`
	// Required. Report URL in the source system.
	// Wire name: 'source_url'
	SourceUrl string `json:"source_url"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ReportSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := reportSpecToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ReportSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &reportSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := reportSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReportSpec) MarshalJSON() ([]byte, error) {
	pb, err := reportSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestartWindow struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	// Wire name: 'days_of_week'
	DaysOfWeek []DayOfWeek `json:"days_of_week,omitempty"`
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	// Wire name: 'start_hour'
	StartHour int `json:"start_hour"`
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	// Wire name: 'time_zone_id'
	TimeZoneId string `json:"time_zone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RestartWindow) EncodeValues(key string, v *url.Values) error {
	pb, err := restartWindowToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RestartWindow) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restartWindowPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restartWindowFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestartWindow) MarshalJSON() ([]byte, error) {
	pb, err := restartWindowToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RunAs) EncodeValues(key string, v *url.Values) error {
	pb, err := runAsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RunAs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runAsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runAsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunAs) MarshalJSON() ([]byte, error) {
	pb, err := runAsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	// Wire name: 'destination_catalog'
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	// Wire name: 'destination_schema'
	DestinationSchema string `json:"destination_schema"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	// Wire name: 'source_catalog'
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Required. Schema name in the source database.
	// Wire name: 'source_schema'
	SourceSchema string `json:"source_schema"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SchemaSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := schemaSpecToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SchemaSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &schemaSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := schemaSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SchemaSpec) MarshalJSON() ([]byte, error) {
	pb, err := schemaSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Sequencing struct {
	// A sequence number, unique and increasing per pipeline.
	// Wire name: 'control_plane_seq_no'
	ControlPlaneSeqNo int64 `json:"control_plane_seq_no,omitempty"`
	// the ID assigned by the data plane.
	// Wire name: 'data_plane_id'
	DataPlaneId *DataPlaneId `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Sequencing) EncodeValues(key string, v *url.Values) error {
	pb, err := sequencingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Sequencing) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sequencingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sequencingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Sequencing) MarshalJSON() ([]byte, error) {
	pb, err := sequencingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SerializedException struct {
	// Runtime class of the exception
	// Wire name: 'class_name'
	ClassName string `json:"class_name,omitempty"`
	// Exception message
	// Wire name: 'message'
	Message string `json:"message,omitempty"`
	// Stack trace consisting of a list of stack frames
	// Wire name: 'stack'
	Stack []StackFrame `json:"stack,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SerializedException) EncodeValues(key string, v *url.Values) error {
	pb, err := serializedExceptionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SerializedException) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &serializedExceptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := serializedExceptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SerializedException) MarshalJSON() ([]byte, error) {
	pb, err := serializedExceptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StackFrame struct {
	// Class from which the method call originated
	// Wire name: 'declaring_class'
	DeclaringClass string `json:"declaring_class,omitempty"`
	// File where the method is defined
	// Wire name: 'file_name'
	FileName string `json:"file_name,omitempty"`
	// Line from which the method was called
	// Wire name: 'line_number'
	LineNumber int `json:"line_number,omitempty"`
	// Name of the method which was called
	// Wire name: 'method_name'
	MethodName string `json:"method_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *StackFrame) EncodeValues(key string, v *url.Values) error {
	pb, err := stackFrameToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StackFrame) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stackFramePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stackFrameFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StackFrame) MarshalJSON() ([]byte, error) {
	pb, err := stackFrameToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StartUpdate struct {

	// Wire name: 'cause'
	Cause StartUpdateCause `json:"cause,omitempty"`
	// If true, this update will reset all tables before running.
	// Wire name: 'full_refresh'
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'full_refresh_selection'
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`

	PipelineId string `json:"-" tf:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'refresh_selection'
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	// Wire name: 'validate_only'
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *StartUpdate) EncodeValues(key string, v *url.Values) error {
	pb, err := startUpdateToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StartUpdate) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startUpdatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startUpdateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartUpdate) MarshalJSON() ([]byte, error) {
	pb, err := startUpdateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// What triggered this update.
type StartUpdateCause string

const StartUpdateCauseApiCall StartUpdateCause = `API_CALL`

const StartUpdateCauseInfrastructureMaintenance StartUpdateCause = `INFRASTRUCTURE_MAINTENANCE`

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
	case `API_CALL`, `INFRASTRUCTURE_MAINTENANCE`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = StartUpdateCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "INFRASTRUCTURE_MAINTENANCE", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Values returns all possible values for StartUpdateCause.
//
// There is no guarantee on the order of the values in the slice.
func (f *StartUpdateCause) Values() []StartUpdateCause {
	return []StartUpdateCause{
		StartUpdateCauseApiCall,
		StartUpdateCauseInfrastructureMaintenance,
		StartUpdateCauseJobTask,
		StartUpdateCauseRetryOnFailure,
		StartUpdateCauseSchemaChange,
		StartUpdateCauseServiceUpgrade,
		StartUpdateCauseUserAction,
	}
}

// Type always returns StartUpdateCause to satisfy [pflag.Value] interface
func (f *StartUpdateCause) Type() string {
	return "StartUpdateCause"
}

type StartUpdateResponse struct {

	// Wire name: 'update_id'
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *StartUpdateResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := startUpdateResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StartUpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startUpdateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startUpdateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartUpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := startUpdateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StopRequest struct {
	PipelineId string `json:"-" tf:"-"`
}

func (st *StopRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := stopRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StopRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stopRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stopRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StopRequest) MarshalJSON() ([]byte, error) {
	pb, err := stopRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	// Wire name: 'destination_catalog'
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store table.
	// Wire name: 'destination_schema'
	DestinationSchema string `json:"destination_schema"`
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	// Wire name: 'destination_table'
	DestinationTable string `json:"destination_table,omitempty"`
	// Source catalog name. Might be optional depending on the type of source.
	// Wire name: 'source_catalog'
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	// Wire name: 'source_schema'
	SourceSchema string `json:"source_schema,omitempty"`
	// Required. Table name in the source database.
	// Wire name: 'source_table'
	SourceTable string `json:"source_table"`
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TableSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := tableSpecToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TableSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableSpec) MarshalJSON() ([]byte, error) {
	pb, err := tableSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TableSpecificConfig struct {
	// A list of column names to be excluded for the ingestion. When not
	// specified, include_columns fully controls what columns to be ingested.
	// When specified, all other columns including future ones will be
	// automatically included for ingestion. This field in mutually exclusive
	// with `include_columns`.
	// Wire name: 'exclude_columns'
	ExcludeColumns []string `json:"exclude_columns,omitempty"`
	// A list of column names to be included for the ingestion. When not
	// specified, all columns except ones in exclude_columns will be included.
	// Future columns will be automatically included. When specified, all other
	// future columns will be automatically excluded from ingestion. This field
	// in mutually exclusive with `exclude_columns`.
	// Wire name: 'include_columns'
	IncludeColumns []string `json:"include_columns,omitempty"`
	// The primary key of the table used to apply changes.
	// Wire name: 'primary_keys'
	PrimaryKeys []string `json:"primary_keys,omitempty"`
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	// Wire name: 'salesforce_include_formula_fields'
	SalesforceIncludeFormulaFields bool `json:"salesforce_include_formula_fields,omitempty"`
	// The SCD type to use to ingest the table.
	// Wire name: 'scd_type'
	ScdType TableSpecificConfigScdType `json:"scd_type,omitempty"`
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	// Wire name: 'sequence_by'
	SequenceBy []string `json:"sequence_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TableSpecificConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := tableSpecificConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TableSpecificConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableSpecificConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableSpecificConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableSpecificConfig) MarshalJSON() ([]byte, error) {
	pb, err := tableSpecificConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The SCD type to use to ingest the table.
type TableSpecificConfigScdType string

const TableSpecificConfigScdTypeAppendOnly TableSpecificConfigScdType = `APPEND_ONLY`

const TableSpecificConfigScdTypeScdType1 TableSpecificConfigScdType = `SCD_TYPE_1`

const TableSpecificConfigScdTypeScdType2 TableSpecificConfigScdType = `SCD_TYPE_2`

// String representation for [fmt.Print]
func (f *TableSpecificConfigScdType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TableSpecificConfigScdType) Set(v string) error {
	switch v {
	case `APPEND_ONLY`, `SCD_TYPE_1`, `SCD_TYPE_2`:
		*f = TableSpecificConfigScdType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPEND_ONLY", "SCD_TYPE_1", "SCD_TYPE_2"`, v)
	}
}

// Values returns all possible values for TableSpecificConfigScdType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TableSpecificConfigScdType) Values() []TableSpecificConfigScdType {
	return []TableSpecificConfigScdType{
		TableSpecificConfigScdTypeAppendOnly,
		TableSpecificConfigScdTypeScdType1,
		TableSpecificConfigScdTypeScdType2,
	}
}

// Type always returns TableSpecificConfigScdType to satisfy [pflag.Value] interface
func (f *TableSpecificConfigScdType) Type() string {
	return "TableSpecificConfigScdType"
}

type UpdateInfo struct {
	// What triggered this update.
	// Wire name: 'cause'
	Cause UpdateInfoCause `json:"cause,omitempty"`
	// The ID of the cluster that the update is running on.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	// Wire name: 'config'
	Config *PipelineSpec `json:"config,omitempty"`
	// The time when this update was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// If true, this update will reset all tables before running.
	// Wire name: 'full_refresh'
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'full_refresh_selection'
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string `json:"pipeline_id,omitempty"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'refresh_selection'
	RefreshSelection []string `json:"refresh_selection,omitempty"`
	// The update state.
	// Wire name: 'state'
	State UpdateInfoState `json:"state,omitempty"`
	// The ID of this update.
	// Wire name: 'update_id'
	UpdateId string `json:"update_id,omitempty"`
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	// Wire name: 'validate_only'
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := updateInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateInfo) MarshalJSON() ([]byte, error) {
	pb, err := updateInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// What triggered this update.
type UpdateInfoCause string

const UpdateInfoCauseApiCall UpdateInfoCause = `API_CALL`

const UpdateInfoCauseInfrastructureMaintenance UpdateInfoCause = `INFRASTRUCTURE_MAINTENANCE`

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
	case `API_CALL`, `INFRASTRUCTURE_MAINTENANCE`, `JOB_TASK`, `RETRY_ON_FAILURE`, `SCHEMA_CHANGE`, `SERVICE_UPGRADE`, `USER_ACTION`:
		*f = UpdateInfoCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API_CALL", "INFRASTRUCTURE_MAINTENANCE", "JOB_TASK", "RETRY_ON_FAILURE", "SCHEMA_CHANGE", "SERVICE_UPGRADE", "USER_ACTION"`, v)
	}
}

// Values returns all possible values for UpdateInfoCause.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateInfoCause) Values() []UpdateInfoCause {
	return []UpdateInfoCause{
		UpdateInfoCauseApiCall,
		UpdateInfoCauseInfrastructureMaintenance,
		UpdateInfoCauseJobTask,
		UpdateInfoCauseRetryOnFailure,
		UpdateInfoCauseSchemaChange,
		UpdateInfoCauseServiceUpgrade,
		UpdateInfoCauseUserAction,
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

// Values returns all possible values for UpdateInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateInfoState) Values() []UpdateInfoState {
	return []UpdateInfoState{
		UpdateInfoStateCanceled,
		UpdateInfoStateCompleted,
		UpdateInfoStateCreated,
		UpdateInfoStateFailed,
		UpdateInfoStateInitializing,
		UpdateInfoStateQueued,
		UpdateInfoStateResetting,
		UpdateInfoStateRunning,
		UpdateInfoStateSettingUpTables,
		UpdateInfoStateStopping,
		UpdateInfoStateWaitingForResources,
	}
}

// Type always returns UpdateInfoState to satisfy [pflag.Value] interface
func (f *UpdateInfoState) Type() string {
	return "UpdateInfoState"
}

type UpdateStateInfo struct {

	// Wire name: 'creation_time'
	CreationTime string `json:"creation_time,omitempty"`

	// Wire name: 'state'
	State UpdateStateInfoState `json:"state,omitempty"`

	// Wire name: 'update_id'
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateStateInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := updateStateInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateStateInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateStateInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateStateInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateStateInfo) MarshalJSON() ([]byte, error) {
	pb, err := updateStateInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The update state.
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

// Values returns all possible values for UpdateStateInfoState.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateStateInfoState) Values() []UpdateStateInfoState {
	return []UpdateStateInfoState{
		UpdateStateInfoStateCanceled,
		UpdateStateInfoStateCompleted,
		UpdateStateInfoStateCreated,
		UpdateStateInfoStateFailed,
		UpdateStateInfoStateInitializing,
		UpdateStateInfoStateQueued,
		UpdateStateInfoStateResetting,
		UpdateStateInfoStateRunning,
		UpdateStateInfoStateSettingUpTables,
		UpdateStateInfoStateStopping,
		UpdateStateInfoStateWaitingForResources,
	}
}

// Type always returns UpdateStateInfoState to satisfy [pflag.Value] interface
func (f *UpdateStateInfoState) Type() string {
	return "UpdateStateInfoState"
}
