// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	// Wire name: 'allow_duplicate_names'
	AllowDuplicateNames bool
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool

	// Wire name: 'dry_run'
	DryRun bool
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	// Wire name: 'run_as'
	RunAs *RunAs
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger *PipelineTrigger

	ForceSendFields []string `tf:"-"`
}

func createPipelineToPb(st *CreatePipeline) (*createPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPipelinePb{}
	pb.AllowDuplicateNames = st.AllowDuplicateNames

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Catalog = st.Catalog

	pb.Channel = st.Channel

	var clustersPb []pipelineClusterPb
	for _, item := range st.Clusters {
		itemPb, err := pipelineClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb

	pb.Configuration = st.Configuration

	pb.Continuous = st.Continuous

	deploymentPb, err := pipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	pb.Development = st.Development

	pb.DryRun = st.DryRun

	pb.Edition = st.Edition

	eventLogPb, err := eventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}

	filtersPb, err := filtersToPb(st.Filters)
	if err != nil {
		return nil, err
	}
	if filtersPb != nil {
		pb.Filters = filtersPb
	}

	gatewayDefinitionPb, err := ingestionGatewayPipelineDefinitionToPb(st.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionPb != nil {
		pb.GatewayDefinition = gatewayDefinitionPb
	}

	pb.Id = st.Id

	ingestionDefinitionPb, err := ingestionPipelineDefinitionToPb(st.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionPb != nil {
		pb.IngestionDefinition = ingestionDefinitionPb
	}

	var librariesPb []pipelineLibraryPb
	for _, item := range st.Libraries {
		itemPb, err := pipelineLibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	pb.Name = st.Name

	var notificationsPb []notificationsPb
	for _, item := range st.Notifications {
		itemPb, err := notificationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notificationsPb = append(notificationsPb, *itemPb)
		}
	}
	pb.Notifications = notificationsPb

	pb.Photon = st.Photon

	restartWindowPb, err := restartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}

	runAsPb, err := runAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}

	pb.Schema = st.Schema

	pb.Serverless = st.Serverless

	pb.Storage = st.Storage

	pb.Target = st.Target

	triggerPb, err := pipelineTriggerToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createPipelinePb struct {
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
	Clusters []pipelineClusterPb `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *pipelineDeploymentPb `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`

	DryRun bool `json:"dry_run,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Event log configuration for this pipeline
	EventLog *eventLogSpecPb `json:"event_log,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *filtersPb `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *ingestionGatewayPipelineDefinitionPb `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *ingestionPipelineDefinitionPb `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []pipelineLibraryPb `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []notificationsPb `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	RestartWindow *restartWindowPb `json:"restart_window,omitempty"`
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *runAsPb `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *pipelineTriggerPb `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPipelineFromPb(pb *createPipelinePb) (*CreatePipeline, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePipeline{}
	st.AllowDuplicateNames = pb.AllowDuplicateNames
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel

	var clustersField []PipelineCluster
	for _, itemPb := range pb.Clusters {
		item, err := pipelineClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	st.Configuration = pb.Configuration
	st.Continuous = pb.Continuous
	deploymentField, err := pipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Development = pb.Development
	st.DryRun = pb.DryRun
	st.Edition = pb.Edition
	eventLogField, err := eventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	filtersField, err := filtersFromPb(pb.Filters)
	if err != nil {
		return nil, err
	}
	if filtersField != nil {
		st.Filters = filtersField
	}
	gatewayDefinitionField, err := ingestionGatewayPipelineDefinitionFromPb(pb.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionField != nil {
		st.GatewayDefinition = gatewayDefinitionField
	}
	st.Id = pb.Id
	ingestionDefinitionField, err := ingestionPipelineDefinitionFromPb(pb.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionField != nil {
		st.IngestionDefinition = ingestionDefinitionField
	}

	var librariesField []PipelineLibrary
	for _, itemPb := range pb.Libraries {
		item, err := pipelineLibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.Name = pb.Name

	var notificationsField []Notifications
	for _, itemPb := range pb.Notifications {
		item, err := notificationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notificationsField = append(notificationsField, *item)
		}
	}
	st.Notifications = notificationsField
	st.Photon = pb.Photon
	restartWindowField, err := restartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	runAsField, err := runAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Target = pb.Target
	triggerField, err := pipelineTriggerFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPipelinePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPipelinePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	// Wire name: 'effective_settings'
	EffectiveSettings *PipelineSpec
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	// Wire name: 'pipeline_id'
	PipelineId string

	ForceSendFields []string `tf:"-"`
}

func createPipelineResponseToPb(st *CreatePipelineResponse) (*createPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPipelineResponsePb{}
	effectiveSettingsPb, err := pipelineSpecToPb(st.EffectiveSettings)
	if err != nil {
		return nil, err
	}
	if effectiveSettingsPb != nil {
		pb.EffectiveSettings = effectiveSettingsPb
	}

	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createPipelineResponsePb struct {
	// Only returned when dry_run is true.
	EffectiveSettings *pipelineSpecPb `json:"effective_settings,omitempty"`
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPipelineResponseFromPb(pb *createPipelineResponsePb) (*CreatePipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePipelineResponse{}
	effectiveSettingsField, err := pipelineSpecFromPb(pb.EffectiveSettings)
	if err != nil {
		return nil, err
	}
	if effectiveSettingsField != nil {
		st.EffectiveSettings = effectiveSettingsField
	}
	st.PipelineId = pb.PipelineId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPipelineResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPipelineResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CronTrigger struct {

	// Wire name: 'quartz_cron_schedule'
	QuartzCronSchedule string

	// Wire name: 'timezone_id'
	TimezoneId string

	ForceSendFields []string `tf:"-"`
}

func cronTriggerToPb(st *CronTrigger) (*cronTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronTriggerPb{}
	pb.QuartzCronSchedule = st.QuartzCronSchedule

	pb.TimezoneId = st.TimezoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type cronTriggerPb struct {
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`

	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cronTriggerFromPb(pb *cronTriggerPb) (*CronTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronTrigger{}
	st.QuartzCronSchedule = pb.QuartzCronSchedule
	st.TimezoneId = pb.TimezoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cronTriggerPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cronTriggerPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	// Wire name: 'instance'
	Instance string
	// A sequence number, unique and increasing within the data plane instance.
	// Wire name: 'seq_no'
	SeqNo int

	ForceSendFields []string `tf:"-"`
}

func dataPlaneIdToPb(st *DataPlaneId) (*dataPlaneIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneIdPb{}
	pb.Instance = st.Instance

	pb.SeqNo = st.SeqNo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dataPlaneIdPb struct {
	// The instance name of the data plane emitting an event.
	Instance string `json:"instance,omitempty"`
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo int `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneIdFromPb(pb *dataPlaneIdPb) (*DataPlaneId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneId{}
	st.Instance = pb.Instance
	st.SeqNo = pb.SeqNo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataPlaneIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataPlaneIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Days of week in which the restart is allowed to happen (within a five-hour
// window starting at start_hour). If not specified all days of the week will be
// used.
type DayOfWeek string
type dayOfWeekPb string

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

func dayOfWeekToPb(st *DayOfWeek) (*dayOfWeekPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dayOfWeekPb(*st)
	return &pb, nil
}

func dayOfWeekFromPb(pb *dayOfWeekPb) (*DayOfWeek, error) {
	if pb == nil {
		return nil, nil
	}
	st := DayOfWeek(*pb)
	return &st, nil
}

// Delete a pipeline
type DeletePipelineRequest struct {

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func deletePipelineRequestToPb(st *DeletePipelineRequest) (*deletePipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePipelineRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
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

type deletePipelineRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func deletePipelineRequestFromPb(pb *deletePipelineRequestPb) (*DeletePipelineRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePipelineRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

type DeletePipelineResponse struct {
}

func deletePipelineResponseToPb(st *DeletePipelineResponse) (*deletePipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePipelineResponsePb{}

	return pb, nil
}

func (st *DeletePipelineResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePipelineResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePipelineResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePipelineResponse) MarshalJSON() ([]byte, error) {
	pb, err := deletePipelineResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deletePipelineResponsePb struct {
}

func deletePipelineResponseFromPb(pb *deletePipelineResponsePb) (*DeletePipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePipelineResponse{}

	return st, nil
}

// The deployment method that manages the pipeline: - BUNDLE: The pipeline is
// managed by a Databricks Asset Bundle.
type DeploymentKind string
type deploymentKindPb string

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

func deploymentKindToPb(st *DeploymentKind) (*deploymentKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := deploymentKindPb(*st)
	return &pb, nil
}

func deploymentKindFromPb(pb *deploymentKindPb) (*DeploymentKind, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeploymentKind(*pb)
	return &st, nil
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	// Wire name: 'allow_duplicate_names'
	AllowDuplicateNames bool
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	// Wire name: 'expected_last_modified'
	ExpectedLastModified int64
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool
	// Unique identifier for this pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	// Wire name: 'run_as'
	RunAs *RunAs
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger *PipelineTrigger

	ForceSendFields []string `tf:"-"`
}

func editPipelineToPb(st *EditPipeline) (*editPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPipelinePb{}
	pb.AllowDuplicateNames = st.AllowDuplicateNames

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Catalog = st.Catalog

	pb.Channel = st.Channel

	var clustersPb []pipelineClusterPb
	for _, item := range st.Clusters {
		itemPb, err := pipelineClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb

	pb.Configuration = st.Configuration

	pb.Continuous = st.Continuous

	deploymentPb, err := pipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	pb.Development = st.Development

	pb.Edition = st.Edition

	eventLogPb, err := eventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}

	pb.ExpectedLastModified = st.ExpectedLastModified

	filtersPb, err := filtersToPb(st.Filters)
	if err != nil {
		return nil, err
	}
	if filtersPb != nil {
		pb.Filters = filtersPb
	}

	gatewayDefinitionPb, err := ingestionGatewayPipelineDefinitionToPb(st.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionPb != nil {
		pb.GatewayDefinition = gatewayDefinitionPb
	}

	pb.Id = st.Id

	ingestionDefinitionPb, err := ingestionPipelineDefinitionToPb(st.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionPb != nil {
		pb.IngestionDefinition = ingestionDefinitionPb
	}

	var librariesPb []pipelineLibraryPb
	for _, item := range st.Libraries {
		itemPb, err := pipelineLibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	pb.Name = st.Name

	var notificationsPb []notificationsPb
	for _, item := range st.Notifications {
		itemPb, err := notificationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notificationsPb = append(notificationsPb, *itemPb)
		}
	}
	pb.Notifications = notificationsPb

	pb.Photon = st.Photon

	pb.PipelineId = st.PipelineId

	restartWindowPb, err := restartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}

	runAsPb, err := runAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}

	pb.Schema = st.Schema

	pb.Serverless = st.Serverless

	pb.Storage = st.Storage

	pb.Target = st.Target

	triggerPb, err := pipelineTriggerToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type editPipelinePb struct {
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
	Clusters []pipelineClusterPb `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *pipelineDeploymentPb `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Event log configuration for this pipeline
	EventLog *eventLogSpecPb `json:"event_log,omitempty"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified int64 `json:"expected_last_modified,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *filtersPb `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *ingestionGatewayPipelineDefinitionPb `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *ingestionPipelineDefinitionPb `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []pipelineLibraryPb `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []notificationsPb `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Unique identifier for this pipeline.
	PipelineId string `json:"pipeline_id,omitempty" url:"-"`
	// Restart window of this pipeline.
	RestartWindow *restartWindowPb `json:"restart_window,omitempty"`
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *runAsPb `json:"run_as,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *pipelineTriggerPb `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editPipelineFromPb(pb *editPipelinePb) (*EditPipeline, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPipeline{}
	st.AllowDuplicateNames = pb.AllowDuplicateNames
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel

	var clustersField []PipelineCluster
	for _, itemPb := range pb.Clusters {
		item, err := pipelineClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	st.Configuration = pb.Configuration
	st.Continuous = pb.Continuous
	deploymentField, err := pipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Development = pb.Development
	st.Edition = pb.Edition
	eventLogField, err := eventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	st.ExpectedLastModified = pb.ExpectedLastModified
	filtersField, err := filtersFromPb(pb.Filters)
	if err != nil {
		return nil, err
	}
	if filtersField != nil {
		st.Filters = filtersField
	}
	gatewayDefinitionField, err := ingestionGatewayPipelineDefinitionFromPb(pb.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionField != nil {
		st.GatewayDefinition = gatewayDefinitionField
	}
	st.Id = pb.Id
	ingestionDefinitionField, err := ingestionPipelineDefinitionFromPb(pb.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionField != nil {
		st.IngestionDefinition = ingestionDefinitionField
	}

	var librariesField []PipelineLibrary
	for _, itemPb := range pb.Libraries {
		item, err := pipelineLibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.Name = pb.Name

	var notificationsField []Notifications
	for _, itemPb := range pb.Notifications {
		item, err := notificationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notificationsField = append(notificationsField, *item)
		}
	}
	st.Notifications = notificationsField
	st.Photon = pb.Photon
	st.PipelineId = pb.PipelineId
	restartWindowField, err := restartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	runAsField, err := runAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Target = pb.Target
	triggerField, err := pipelineTriggerFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editPipelinePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editPipelinePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditPipelineResponse struct {
}

func editPipelineResponseToPb(st *EditPipelineResponse) (*editPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPipelineResponsePb{}

	return pb, nil
}

func (st *EditPipelineResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editPipelineResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editPipelineResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditPipelineResponse) MarshalJSON() ([]byte, error) {
	pb, err := editPipelineResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type editPipelineResponsePb struct {
}

func editPipelineResponseFromPb(pb *editPipelineResponsePb) (*EditPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPipelineResponse{}

	return st, nil
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	// Wire name: 'exceptions'
	Exceptions []SerializedException
	// Whether this error is considered fatal, that is, unrecoverable.
	// Wire name: 'fatal'
	Fatal bool

	ForceSendFields []string `tf:"-"`
}

func errorDetailToPb(st *ErrorDetail) (*errorDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &errorDetailPb{}

	var exceptionsPb []serializedExceptionPb
	for _, item := range st.Exceptions {
		itemPb, err := serializedExceptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			exceptionsPb = append(exceptionsPb, *itemPb)
		}
	}
	pb.Exceptions = exceptionsPb

	pb.Fatal = st.Fatal

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type errorDetailPb struct {
	// The exception thrown for this error, with its chain of cause.
	Exceptions []serializedExceptionPb `json:"exceptions,omitempty"`
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal bool `json:"fatal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func errorDetailFromPb(pb *errorDetailPb) (*ErrorDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ErrorDetail{}

	var exceptionsField []SerializedException
	for _, itemPb := range pb.Exceptions {
		item, err := serializedExceptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			exceptionsField = append(exceptionsField, *item)
		}
	}
	st.Exceptions = exceptionsField
	st.Fatal = pb.Fatal

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *errorDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st errorDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The severity level of the event.
type EventLevel string
type eventLevelPb string

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

func eventLevelToPb(st *EventLevel) (*eventLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := eventLevelPb(*st)
	return &pb, nil
}

func eventLevelFromPb(pb *eventLevelPb) (*EventLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := EventLevel(*pb)
	return &st, nil
}

// Configurable event log parameters.
type EventLogSpec struct {
	// The UC catalog the event log is published under.
	// Wire name: 'catalog'
	Catalog string
	// The name the event log is published to in UC.
	// Wire name: 'name'
	Name string
	// The UC schema the event log is published under.
	// Wire name: 'schema'
	Schema string

	ForceSendFields []string `tf:"-"`
}

func eventLogSpecToPb(st *EventLogSpec) (*eventLogSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &eventLogSpecPb{}
	pb.Catalog = st.Catalog

	pb.Name = st.Name

	pb.Schema = st.Schema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type eventLogSpecPb struct {
	// The UC catalog the event log is published under.
	Catalog string `json:"catalog,omitempty"`
	// The name the event log is published to in UC.
	Name string `json:"name,omitempty"`
	// The UC schema the event log is published under.
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func eventLogSpecFromPb(pb *eventLogSpecPb) (*EventLogSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EventLogSpec{}
	st.Catalog = pb.Catalog
	st.Name = pb.Name
	st.Schema = pb.Schema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *eventLogSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st eventLogSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileLibrary struct {
	// The absolute path of the file.
	// Wire name: 'path'
	Path string

	ForceSendFields []string `tf:"-"`
}

func fileLibraryToPb(st *FileLibrary) (*fileLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileLibraryPb{}
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type fileLibraryPb struct {
	// The absolute path of the file.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileLibraryFromPb(pb *fileLibraryPb) (*FileLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileLibrary{}
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Filters struct {
	// Paths to exclude.
	// Wire name: 'exclude'
	Exclude []string
	// Paths to include.
	// Wire name: 'include'
	Include []string
}

func filtersToPb(st *Filters) (*filtersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filtersPb{}
	pb.Exclude = st.Exclude

	pb.Include = st.Include

	return pb, nil
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

type filtersPb struct {
	// Paths to exclude.
	Exclude []string `json:"exclude,omitempty"`
	// Paths to include.
	Include []string `json:"include,omitempty"`
}

func filtersFromPb(pb *filtersPb) (*Filters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Filters{}
	st.Exclude = pb.Exclude
	st.Include = pb.Include

	return st, nil
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func getPipelinePermissionLevelsRequestToPb(st *GetPipelinePermissionLevelsRequest) (*getPipelinePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionLevelsRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
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

type getPipelinePermissionLevelsRequestPb struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

func getPipelinePermissionLevelsRequestFromPb(pb *getPipelinePermissionLevelsRequestPb) (*GetPipelinePermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionLevelsRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []PipelinePermissionsDescription
}

func getPipelinePermissionLevelsResponseToPb(st *GetPipelinePermissionLevelsResponse) (*getPipelinePermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionLevelsResponsePb{}

	var permissionLevelsPb []pipelinePermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := pipelinePermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
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

type getPipelinePermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []pipelinePermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getPipelinePermissionLevelsResponseFromPb(pb *getPipelinePermissionLevelsResponsePb) (*GetPipelinePermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionLevelsResponse{}

	var permissionLevelsField []PipelinePermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := pipelinePermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

// Get pipeline permissions
type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func getPipelinePermissionsRequestToPb(st *GetPipelinePermissionsRequest) (*getPipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionsRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
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

type getPipelinePermissionsRequestPb struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

func getPipelinePermissionsRequestFromPb(pb *getPipelinePermissionsRequestPb) (*GetPipelinePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionsRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

// Get a pipeline
type GetPipelineRequest struct {

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func getPipelineRequestToPb(st *GetPipelineRequest) (*getPipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelineRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
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

type getPipelineRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func getPipelineRequestFromPb(pb *getPipelineRequestPb) (*GetPipelineRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelineRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	// Wire name: 'cause'
	Cause string
	// The ID of the cluster that the pipeline is running on.
	// Wire name: 'cluster_id'
	ClusterId string
	// The username of the pipeline creator.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// Serverless budget policy ID of this pipeline.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string
	// The health of a pipeline.
	// Wire name: 'health'
	Health GetPipelineResponseHealth
	// The last time the pipeline settings were modified or created.
	// Wire name: 'last_modified'
	LastModified int64
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	// Wire name: 'latest_updates'
	LatestUpdates []UpdateStateInfo
	// A human friendly identifier for the pipeline, taken from the `spec`.
	// Wire name: 'name'
	Name string
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string
	// Username of the user that the pipeline will run on behalf of.
	// Wire name: 'run_as_user_name'
	RunAsUserName string
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	// Wire name: 'spec'
	Spec *PipelineSpec
	// The pipeline state.
	// Wire name: 'state'
	State PipelineState

	ForceSendFields []string `tf:"-"`
}

func getPipelineResponseToPb(st *GetPipelineResponse) (*getPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelineResponsePb{}
	pb.Cause = st.Cause

	pb.ClusterId = st.ClusterId

	pb.CreatorUserName = st.CreatorUserName

	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId

	pb.Health = st.Health

	pb.LastModified = st.LastModified

	var latestUpdatesPb []updateStateInfoPb
	for _, item := range st.LatestUpdates {
		itemPb, err := updateStateInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			latestUpdatesPb = append(latestUpdatesPb, *itemPb)
		}
	}
	pb.LatestUpdates = latestUpdatesPb

	pb.Name = st.Name

	pb.PipelineId = st.PipelineId

	pb.RunAsUserName = st.RunAsUserName

	specPb, err := pipelineSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getPipelineResponsePb struct {
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
	LatestUpdates []updateStateInfoPb `json:"latest_updates,omitempty"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name string `json:"name,omitempty"`
	// The ID of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec *pipelineSpecPb `json:"spec,omitempty"`
	// The pipeline state.
	State PipelineState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPipelineResponseFromPb(pb *getPipelineResponsePb) (*GetPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelineResponse{}
	st.Cause = pb.Cause
	st.ClusterId = pb.ClusterId
	st.CreatorUserName = pb.CreatorUserName
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.Health = pb.Health
	st.LastModified = pb.LastModified

	var latestUpdatesField []UpdateStateInfo
	for _, itemPb := range pb.LatestUpdates {
		item, err := updateStateInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			latestUpdatesField = append(latestUpdatesField, *item)
		}
	}
	st.LatestUpdates = latestUpdatesField
	st.Name = pb.Name
	st.PipelineId = pb.PipelineId
	st.RunAsUserName = pb.RunAsUserName
	specField, err := pipelineSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPipelineResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPipelineResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The health of a pipeline.
type GetPipelineResponseHealth string
type getPipelineResponseHealthPb string

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

func getPipelineResponseHealthToPb(st *GetPipelineResponseHealth) (*getPipelineResponseHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getPipelineResponseHealthPb(*st)
	return &pb, nil
}

func getPipelineResponseHealthFromPb(pb *getPipelineResponseHealthPb) (*GetPipelineResponseHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetPipelineResponseHealth(*pb)
	return &st, nil
}

// Get a pipeline update
type GetUpdateRequest struct {
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
	// The ID of the update.
	// Wire name: 'update_id'
	UpdateId string `tf:"-"`
}

func getUpdateRequestToPb(st *GetUpdateRequest) (*getUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getUpdateRequestPb{}
	pb.PipelineId = st.PipelineId

	pb.UpdateId = st.UpdateId

	return pb, nil
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

type getUpdateRequestPb struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" url:"-"`
	// The ID of the update.
	UpdateId string `json:"-" url:"-"`
}

func getUpdateRequestFromPb(pb *getUpdateRequestPb) (*GetUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetUpdateRequest{}
	st.PipelineId = pb.PipelineId
	st.UpdateId = pb.UpdateId

	return st, nil
}

type GetUpdateResponse struct {
	// The current update info.
	// Wire name: 'update'
	Update *UpdateInfo
}

func getUpdateResponseToPb(st *GetUpdateResponse) (*getUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getUpdateResponsePb{}
	updatePb, err := updateInfoToPb(st.Update)
	if err != nil {
		return nil, err
	}
	if updatePb != nil {
		pb.Update = updatePb
	}

	return pb, nil
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

type getUpdateResponsePb struct {
	// The current update info.
	Update *updateInfoPb `json:"update,omitempty"`
}

func getUpdateResponseFromPb(pb *getUpdateResponsePb) (*GetUpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetUpdateResponse{}
	updateField, err := updateInfoFromPb(pb.Update)
	if err != nil {
		return nil, err
	}
	if updateField != nil {
		st.Update = updateField
	}

	return st, nil
}

type IngestionConfig struct {
	// Select a specific source report.
	// Wire name: 'report'
	Report *ReportSpec
	// Select all tables from a specific source schema.
	// Wire name: 'schema'
	Schema *SchemaSpec
	// Select a specific source table.
	// Wire name: 'table'
	Table *TableSpec
}

func ingestionConfigToPb(st *IngestionConfig) (*ingestionConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionConfigPb{}
	reportPb, err := reportSpecToPb(st.Report)
	if err != nil {
		return nil, err
	}
	if reportPb != nil {
		pb.Report = reportPb
	}

	schemaPb, err := schemaSpecToPb(st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = schemaPb
	}

	tablePb, err := tableSpecToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
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

type ingestionConfigPb struct {
	// Select a specific source report.
	Report *reportSpecPb `json:"report,omitempty"`
	// Select all tables from a specific source schema.
	Schema *schemaSpecPb `json:"schema,omitempty"`
	// Select a specific source table.
	Table *tableSpecPb `json:"table,omitempty"`
}

func ingestionConfigFromPb(pb *ingestionConfigPb) (*IngestionConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionConfig{}
	reportField, err := reportSpecFromPb(pb.Report)
	if err != nil {
		return nil, err
	}
	if reportField != nil {
		st.Report = reportField
	}
	schemaField, err := schemaSpecFromPb(pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = schemaField
	}
	tableField, err := tableSpecFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}

	return st, nil
}

type IngestionGatewayPipelineDefinition struct {
	// [Deprecated, use connection_name instead] Immutable. The Unity Catalog
	// connection that this gateway pipeline uses to communicate with the
	// source.
	// Wire name: 'connection_id'
	ConnectionId string
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	// Wire name: 'connection_name'
	ConnectionName string
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	// Wire name: 'gateway_storage_catalog'
	GatewayStorageCatalog string
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	// Wire name: 'gateway_storage_name'
	GatewayStorageName string
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	// Wire name: 'gateway_storage_schema'
	GatewayStorageSchema string

	ForceSendFields []string `tf:"-"`
}

func ingestionGatewayPipelineDefinitionToPb(st *IngestionGatewayPipelineDefinition) (*ingestionGatewayPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionGatewayPipelineDefinitionPb{}
	pb.ConnectionId = st.ConnectionId

	pb.ConnectionName = st.ConnectionName

	pb.GatewayStorageCatalog = st.GatewayStorageCatalog

	pb.GatewayStorageName = st.GatewayStorageName

	pb.GatewayStorageSchema = st.GatewayStorageSchema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type ingestionGatewayPipelineDefinitionPb struct {
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

func ingestionGatewayPipelineDefinitionFromPb(pb *ingestionGatewayPipelineDefinitionPb) (*IngestionGatewayPipelineDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionGatewayPipelineDefinition{}
	st.ConnectionId = pb.ConnectionId
	st.ConnectionName = pb.ConnectionName
	st.GatewayStorageCatalog = pb.GatewayStorageCatalog
	st.GatewayStorageName = pb.GatewayStorageName
	st.GatewayStorageSchema = pb.GatewayStorageSchema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ingestionGatewayPipelineDefinitionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ingestionGatewayPipelineDefinitionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection that this ingestion pipeline uses
	// to communicate with the source. This is used with connectors for
	// applications like Salesforce, Workday, and so on.
	// Wire name: 'connection_name'
	ConnectionName string
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	// Wire name: 'ingestion_gateway_id'
	IngestionGatewayId string
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	// Wire name: 'objects'
	Objects []IngestionConfig
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string `tf:"-"`
}

func ingestionPipelineDefinitionToPb(st *IngestionPipelineDefinition) (*ingestionPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionPipelineDefinitionPb{}
	pb.ConnectionName = st.ConnectionName

	pb.IngestionGatewayId = st.IngestionGatewayId

	var objectsPb []ingestionConfigPb
	for _, item := range st.Objects {
		itemPb, err := ingestionConfigToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			objectsPb = append(objectsPb, *itemPb)
		}
	}
	pb.Objects = objectsPb

	tableConfigurationPb, err := tableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type ingestionPipelineDefinitionPb struct {
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
	Objects []ingestionConfigPb `json:"objects,omitempty"`
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration *tableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ingestionPipelineDefinitionFromPb(pb *ingestionPipelineDefinitionPb) (*IngestionPipelineDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionPipelineDefinition{}
	st.ConnectionName = pb.ConnectionName
	st.IngestionGatewayId = pb.IngestionGatewayId

	var objectsField []IngestionConfig
	for _, itemPb := range pb.Objects {
		item, err := ingestionConfigFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			objectsField = append(objectsField, *item)
		}
	}
	st.Objects = objectsField
	tableConfigurationField, err := tableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ingestionPipelineDefinitionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ingestionPipelineDefinitionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listPipelineEventsRequestToPb(st *ListPipelineEventsRequest) (*listPipelineEventsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelineEventsRequestPb{}
	pb.Filter = st.Filter

	pb.MaxResults = st.MaxResults

	pb.OrderBy = st.OrderBy

	pb.PageToken = st.PageToken

	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPipelineEventsRequestPb struct {
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

func listPipelineEventsRequestFromPb(pb *listPipelineEventsRequestPb) (*ListPipelineEventsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelineEventsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	st.PipelineId = pb.PipelineId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelineEventsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelineEventsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	// Wire name: 'events'
	Events []PipelineEvent
	// If present, a token to fetch the next page of events.
	// Wire name: 'next_page_token'
	NextPageToken string
	// If present, a token to fetch the previous page of events.
	// Wire name: 'prev_page_token'
	PrevPageToken string

	ForceSendFields []string `tf:"-"`
}

func listPipelineEventsResponseToPb(st *ListPipelineEventsResponse) (*listPipelineEventsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelineEventsResponsePb{}

	var eventsPb []pipelineEventPb
	for _, item := range st.Events {
		itemPb, err := pipelineEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb

	pb.NextPageToken = st.NextPageToken

	pb.PrevPageToken = st.PrevPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPipelineEventsResponsePb struct {
	// The list of events matching the request criteria.
	Events []pipelineEventPb `json:"events,omitempty"`
	// If present, a token to fetch the next page of events.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, a token to fetch the previous page of events.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelineEventsResponseFromPb(pb *listPipelineEventsResponsePb) (*ListPipelineEventsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelineEventsResponse{}

	var eventsField []PipelineEvent
	for _, itemPb := range pb.Events {
		item, err := pipelineEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelineEventsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelineEventsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Page token returned by previous call
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listPipelinesRequestToPb(st *ListPipelinesRequest) (*listPipelinesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelinesRequestPb{}
	pb.Filter = st.Filter

	pb.MaxResults = st.MaxResults

	pb.OrderBy = st.OrderBy

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPipelinesRequestPb struct {
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

func listPipelinesRequestFromPb(pb *listPipelinesRequestPb) (*ListPipelinesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelinesRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelinesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelinesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	// Wire name: 'next_page_token'
	NextPageToken string
	// The list of events matching the request criteria.
	// Wire name: 'statuses'
	Statuses []PipelineStateInfo

	ForceSendFields []string `tf:"-"`
}

func listPipelinesResponseToPb(st *ListPipelinesResponse) (*listPipelinesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelinesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var statusesPb []pipelineStateInfoPb
	for _, item := range st.Statuses {
		itemPb, err := pipelineStateInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statusesPb = append(statusesPb, *itemPb)
		}
	}
	pb.Statuses = statusesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPipelinesResponsePb struct {
	// If present, a token to fetch the next page of events.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of events matching the request criteria.
	Statuses []pipelineStateInfoPb `json:"statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelinesResponseFromPb(pb *listPipelinesResponsePb) (*ListPipelinesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelinesResponse{}
	st.NextPageToken = pb.NextPageToken

	var statusesField []PipelineStateInfo
	for _, itemPb := range pb.Statuses {
		item, err := pipelineStateInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statusesField = append(statusesField, *item)
		}
	}
	st.Statuses = statusesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelinesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelinesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List pipeline updates
type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Page token returned by previous call
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The pipeline to return updates for.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
	// If present, returns updates until and including this update_id.
	// Wire name: 'until_update_id'
	UntilUpdateId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listUpdatesRequestToPb(st *ListUpdatesRequest) (*listUpdatesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUpdatesRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.PipelineId = st.PipelineId

	pb.UntilUpdateId = st.UntilUpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listUpdatesRequestPb struct {
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

func listUpdatesRequestFromPb(pb *listUpdatesRequestPb) (*ListUpdatesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUpdatesRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.PipelineId = pb.PipelineId
	st.UntilUpdateId = pb.UntilUpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listUpdatesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listUpdatesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	// Wire name: 'next_page_token'
	NextPageToken string
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	// Wire name: 'prev_page_token'
	PrevPageToken string

	// Wire name: 'updates'
	Updates []UpdateInfo

	ForceSendFields []string `tf:"-"`
}

func listUpdatesResponseToPb(st *ListUpdatesResponse) (*listUpdatesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUpdatesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	pb.PrevPageToken = st.PrevPageToken

	var updatesPb []updateInfoPb
	for _, item := range st.Updates {
		itemPb, err := updateInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			updatesPb = append(updatesPb, *itemPb)
		}
	}
	pb.Updates = updatesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listUpdatesResponsePb struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	Updates []updateInfoPb `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listUpdatesResponseFromPb(pb *listUpdatesResponsePb) (*ListUpdatesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUpdatesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	var updatesField []UpdateInfo
	for _, itemPb := range pb.Updates {
		item, err := updateInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			updatesField = append(updatesField, *item)
		}
	}
	st.Updates = updatesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listUpdatesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listUpdatesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ManualTrigger struct {
}

func manualTriggerToPb(st *ManualTrigger) (*manualTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &manualTriggerPb{}

	return pb, nil
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

type manualTriggerPb struct {
}

func manualTriggerFromPb(pb *manualTriggerPb) (*ManualTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ManualTrigger{}

	return st, nil
}

// Maturity level for EventDetails.
type MaturityLevel string
type maturityLevelPb string

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

func maturityLevelToPb(st *MaturityLevel) (*maturityLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := maturityLevelPb(*st)
	return &pb, nil
}

func maturityLevelFromPb(pb *maturityLevelPb) (*MaturityLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := MaturityLevel(*pb)
	return &st, nil
}

type NotebookLibrary struct {
	// The absolute path of the notebook.
	// Wire name: 'path'
	Path string

	ForceSendFields []string `tf:"-"`
}

func notebookLibraryToPb(st *NotebookLibrary) (*notebookLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookLibraryPb{}
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type notebookLibraryPb struct {
	// The absolute path of the notebook.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookLibraryFromPb(pb *notebookLibraryPb) (*NotebookLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookLibrary{}
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notebookLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Alerts []string
	// A list of email addresses notified when a configured alert is triggered.
	// Wire name: 'email_recipients'
	EmailRecipients []string
}

func notificationsToPb(st *Notifications) (*notificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notificationsPb{}
	pb.Alerts = st.Alerts

	pb.EmailRecipients = st.EmailRecipients

	return pb, nil
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

type notificationsPb struct {
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

func notificationsFromPb(pb *notificationsPb) (*Notifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Notifications{}
	st.Alerts = pb.Alerts
	st.EmailRecipients = pb.EmailRecipients

	return st, nil
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	// Wire name: 'batch_id'
	BatchId int
	// The cloud provider, e.g., AWS or Azure.
	// Wire name: 'cloud'
	Cloud string
	// The id of the cluster where an execution happens. Unique within a region.
	// Wire name: 'cluster_id'
	ClusterId string
	// The name of a dataset. Unique within a pipeline.
	// Wire name: 'dataset_name'
	DatasetName string
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	// Wire name: 'flow_id'
	FlowId string
	// The name of the flow. Not unique.
	// Wire name: 'flow_name'
	FlowName string
	// The optional host name where the event was triggered
	// Wire name: 'host'
	Host string
	// The id of a maintenance run. Globally unique.
	// Wire name: 'maintenance_id'
	MaintenanceId string
	// Materialization name.
	// Wire name: 'materialization_name'
	MaterializationName string
	// The org id of the user. Unique within a cloud.
	// Wire name: 'org_id'
	OrgId int
	// The id of the pipeline. Globally unique.
	// Wire name: 'pipeline_id'
	PipelineId string
	// The name of the pipeline. Not unique.
	// Wire name: 'pipeline_name'
	PipelineName string
	// The cloud region.
	// Wire name: 'region'
	Region string
	// The id of the request that caused an update.
	// Wire name: 'request_id'
	RequestId string
	// The id of a (delta) table. Globally unique.
	// Wire name: 'table_id'
	TableId string
	// The Unity Catalog id of the MV or ST being updated.
	// Wire name: 'uc_resource_id'
	UcResourceId string
	// The id of an execution. Globally unique.
	// Wire name: 'update_id'
	UpdateId string

	ForceSendFields []string `tf:"-"`
}

func originToPb(st *Origin) (*originPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &originPb{}
	pb.BatchId = st.BatchId

	pb.Cloud = st.Cloud

	pb.ClusterId = st.ClusterId

	pb.DatasetName = st.DatasetName

	pb.FlowId = st.FlowId

	pb.FlowName = st.FlowName

	pb.Host = st.Host

	pb.MaintenanceId = st.MaintenanceId

	pb.MaterializationName = st.MaterializationName

	pb.OrgId = st.OrgId

	pb.PipelineId = st.PipelineId

	pb.PipelineName = st.PipelineName

	pb.Region = st.Region

	pb.RequestId = st.RequestId

	pb.TableId = st.TableId

	pb.UcResourceId = st.UcResourceId

	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type originPb struct {
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

func originFromPb(pb *originPb) (*Origin, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Origin{}
	st.BatchId = pb.BatchId
	st.Cloud = pb.Cloud
	st.ClusterId = pb.ClusterId
	st.DatasetName = pb.DatasetName
	st.FlowId = pb.FlowId
	st.FlowName = pb.FlowName
	st.Host = pb.Host
	st.MaintenanceId = pb.MaintenanceId
	st.MaterializationName = pb.MaterializationName
	st.OrgId = pb.OrgId
	st.PipelineId = pb.PipelineId
	st.PipelineName = pb.PipelineName
	st.Region = pb.Region
	st.RequestId = pb.RequestId
	st.TableId = pb.TableId
	st.UcResourceId = pb.UcResourceId
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *originPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st originPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func pipelineAccessControlRequestToPb(st *PipelineAccessControlRequest) (*pipelineAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineAccessControlRequestPb struct {
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

func pipelineAccessControlRequestFromPb(pb *pipelineAccessControlRequestPb) (*PipelineAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []PipelinePermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func pipelineAccessControlResponseToPb(st *PipelineAccessControlResponse) (*pipelineAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineAccessControlResponsePb{}

	var allPermissionsPb []pipelinePermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := pipelinePermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	pb.DisplayName = st.DisplayName

	pb.GroupName = st.GroupName

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []pipelinePermissionPb `json:"all_permissions,omitempty"`
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

func pipelineAccessControlResponseFromPb(pb *pipelineAccessControlResponsePb) (*PipelineAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineAccessControlResponse{}

	var allPermissionsField []PipelinePermission
	for _, itemPb := range pb.AllPermissions {
		item, err := pipelinePermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *PipelineClusterAutoscale
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *compute.AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *compute.AzureAttributes
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *compute.ClusterLogConf
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string
	// Whether to enable local disk encryption for the cluster.
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *compute.GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []compute.InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	// Wire name: 'label'
	Label string
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	// Wire name: 'spark_conf'
	SparkConf map[string]string
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
	SparkEnvVars map[string]string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string

	ForceSendFields []string `tf:"-"`
}

func pipelineClusterToPb(st *PipelineCluster) (*pipelineClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues

	autoscalePb, err := pipelineClusterAutoscaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	awsAttributesPb, err := compute.AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}

	azureAttributesPb, err := compute.AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}

	clusterLogConfPb, err := compute.ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}

	pb.CustomTags = st.CustomTags

	pb.DriverInstancePoolId = st.DriverInstancePoolId

	pb.DriverNodeTypeId = st.DriverNodeTypeId

	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption

	gcpAttributesPb, err := compute.GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []compute.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := compute.InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb

	pb.InstancePoolId = st.InstancePoolId

	pb.Label = st.Label

	pb.NodeTypeId = st.NodeTypeId

	pb.NumWorkers = st.NumWorkers

	pb.PolicyId = st.PolicyId

	pb.SparkConf = st.SparkConf

	pb.SparkEnvVars = st.SparkEnvVars

	pb.SshPublicKeys = st.SshPublicKeys

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineClusterPb struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *pipelineClusterAutoscalePb `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *compute.AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *compute.AzureAttributesPb `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *compute.ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
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
	GcpAttributes *compute.GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []compute.InitScriptInfoPb `json:"init_scripts,omitempty"`
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

func pipelineClusterFromPb(pb *pipelineClusterPb) (*PipelineCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	autoscaleField, err := pipelineClusterAutoscaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	awsAttributesField, err := compute.AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := compute.AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	clusterLogConfField, err := compute.ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	st.CustomTags = pb.CustomTags
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	gcpAttributesField, err := compute.GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []compute.InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := compute.InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.Label = pb.Label
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SshPublicKeys = pb.SshPublicKeys

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	// Wire name: 'max_workers'
	MaxWorkers int
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	// Wire name: 'min_workers'
	MinWorkers int
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	// Wire name: 'mode'
	Mode PipelineClusterAutoscaleMode
}

func pipelineClusterAutoscaleToPb(st *PipelineClusterAutoscale) (*pipelineClusterAutoscalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineClusterAutoscalePb{}
	pb.MaxWorkers = st.MaxWorkers

	pb.MinWorkers = st.MinWorkers

	pb.Mode = st.Mode

	return pb, nil
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

type pipelineClusterAutoscalePb struct {
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

func pipelineClusterAutoscaleFromPb(pb *pipelineClusterAutoscalePb) (*PipelineClusterAutoscale, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineClusterAutoscale{}
	st.MaxWorkers = pb.MaxWorkers
	st.MinWorkers = pb.MinWorkers
	st.Mode = pb.Mode

	return st, nil
}

// Databricks Enhanced Autoscaling optimizes cluster utilization by
// automatically allocating cluster resources based on workload volume, with
// minimal impact to the data processing latency of your pipelines. Enhanced
// Autoscaling is available for `updates` clusters only. The legacy autoscaling
// feature is used for `maintenance` clusters.
type PipelineClusterAutoscaleMode string
type pipelineClusterAutoscaleModePb string

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

func pipelineClusterAutoscaleModeToPb(st *PipelineClusterAutoscaleMode) (*pipelineClusterAutoscaleModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelineClusterAutoscaleModePb(*st)
	return &pb, nil
}

func pipelineClusterAutoscaleModeFromPb(pb *pipelineClusterAutoscaleModePb) (*PipelineClusterAutoscaleMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineClusterAutoscaleMode(*pb)
	return &st, nil
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	// Wire name: 'kind'
	Kind DeploymentKind
	// The path to the file containing metadata about the deployment.
	// Wire name: 'metadata_file_path'
	MetadataFilePath string

	ForceSendFields []string `tf:"-"`
}

func pipelineDeploymentToPb(st *PipelineDeployment) (*pipelineDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineDeploymentPb{}
	pb.Kind = st.Kind

	pb.MetadataFilePath = st.MetadataFilePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineDeploymentPb struct {
	// The deployment method that manages the pipeline.
	Kind DeploymentKind `json:"kind,omitempty"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineDeploymentFromPb(pb *pipelineDeploymentPb) (*PipelineDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineDeployment{}
	st.Kind = pb.Kind
	st.MetadataFilePath = pb.MetadataFilePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	// Wire name: 'error'
	Error *ErrorDetail
	// The event type. Should always correspond to the details
	// Wire name: 'event_type'
	EventType string
	// A time-based, globally unique id.
	// Wire name: 'id'
	Id string
	// The severity level of the event.
	// Wire name: 'level'
	Level EventLevel
	// Maturity level for event_type.
	// Wire name: 'maturity_level'
	MaturityLevel MaturityLevel
	// The display message associated with the event.
	// Wire name: 'message'
	Message string
	// Describes where the event originates from.
	// Wire name: 'origin'
	Origin *Origin
	// A sequencing object to identify and order events.
	// Wire name: 'sequence'
	Sequence *Sequencing
	// The time of the event.
	// Wire name: 'timestamp'
	Timestamp string

	ForceSendFields []string `tf:"-"`
}

func pipelineEventToPb(st *PipelineEvent) (*pipelineEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineEventPb{}
	errorPb, err := errorDetailToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}

	pb.EventType = st.EventType

	pb.Id = st.Id

	pb.Level = st.Level

	pb.MaturityLevel = st.MaturityLevel

	pb.Message = st.Message

	originPb, err := originToPb(st.Origin)
	if err != nil {
		return nil, err
	}
	if originPb != nil {
		pb.Origin = originPb
	}

	sequencePb, err := sequencingToPb(st.Sequence)
	if err != nil {
		return nil, err
	}
	if sequencePb != nil {
		pb.Sequence = sequencePb
	}

	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineEventPb struct {
	// Information about an error captured by the event.
	Error *errorDetailPb `json:"error,omitempty"`
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
	Origin *originPb `json:"origin,omitempty"`
	// A sequencing object to identify and order events.
	Sequence *sequencingPb `json:"sequence,omitempty"`
	// The time of the event.
	Timestamp string `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineEventFromPb(pb *pipelineEventPb) (*PipelineEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineEvent{}
	errorField, err := errorDetailFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	st.EventType = pb.EventType
	st.Id = pb.Id
	st.Level = pb.Level
	st.MaturityLevel = pb.MaturityLevel
	st.Message = pb.Message
	originField, err := originFromPb(pb.Origin)
	if err != nil {
		return nil, err
	}
	if originField != nil {
		st.Origin = originField
	}
	sequenceField, err := sequencingFromPb(pb.Sequence)
	if err != nil {
		return nil, err
	}
	if sequenceField != nil {
		st.Sequence = sequenceField
	}
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineEventPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineEventPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	// Wire name: 'file'
	File *FileLibrary
	// URI of the jar to be installed. Currently only DBFS is supported.
	// Wire name: 'jar'
	Jar string
	// Specification of a maven library to be installed.
	// Wire name: 'maven'
	Maven *compute.MavenLibrary
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	// Wire name: 'notebook'
	Notebook *NotebookLibrary
	// URI of the whl to be installed.
	// Wire name: 'whl'
	Whl string

	ForceSendFields []string `tf:"-"`
}

func pipelineLibraryToPb(st *PipelineLibrary) (*pipelineLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineLibraryPb{}
	filePb, err := fileLibraryToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}

	pb.Jar = st.Jar

	mavenPb, err := compute.MavenLibraryToPb(st.Maven)
	if err != nil {
		return nil, err
	}
	if mavenPb != nil {
		pb.Maven = mavenPb
	}

	notebookPb, err := notebookLibraryToPb(st.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookPb != nil {
		pb.Notebook = notebookPb
	}

	pb.Whl = st.Whl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineLibraryPb struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	File *fileLibraryPb `json:"file,omitempty"`
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed.
	Maven *compute.MavenLibraryPb `json:"maven,omitempty"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook *notebookLibraryPb `json:"notebook,omitempty"`
	// URI of the whl to be installed.
	Whl string `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineLibraryFromPb(pb *pipelineLibraryPb) (*PipelineLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineLibrary{}
	fileField, err := fileLibraryFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	st.Jar = pb.Jar
	mavenField, err := compute.MavenLibraryFromPb(pb.Maven)
	if err != nil {
		return nil, err
	}
	if mavenField != nil {
		st.Maven = mavenField
	}
	notebookField, err := notebookLibraryFromPb(pb.Notebook)
	if err != nil {
		return nil, err
	}
	if notebookField != nil {
		st.Notebook = notebookField
	}
	st.Whl = pb.Whl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel

	ForceSendFields []string `tf:"-"`
}

func pipelinePermissionToPb(st *PipelinePermission) (*pipelinePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelinePermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelinePermissionFromPb(pb *pipelinePermissionPb) (*PipelinePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelinePermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelinePermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type PipelinePermissionLevel string
type pipelinePermissionLevelPb string

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

func pipelinePermissionLevelToPb(st *PipelinePermissionLevel) (*pipelinePermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinePermissionLevelPb(*st)
	return &pb, nil
}

func pipelinePermissionLevelFromPb(pb *pipelinePermissionLevelPb) (*PipelinePermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelinePermissionLevel(*pb)
	return &st, nil
}

type PipelinePermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []PipelineAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func pipelinePermissionsToPb(st *PipelinePermissions) (*pipelinePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsPb{}

	var accessControlListPb []pipelineAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := pipelineAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelinePermissionsPb struct {
	AccessControlList []pipelineAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelinePermissionsFromPb(pb *pipelinePermissionsPb) (*PipelinePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissions{}

	var accessControlListField []PipelineAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := pipelineAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelinePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelinePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel

	ForceSendFields []string `tf:"-"`
}

func pipelinePermissionsDescriptionToPb(st *PipelinePermissionsDescription) (*pipelinePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelinePermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelinePermissionsDescriptionFromPb(pb *pipelinePermissionsDescriptionPb) (*PipelinePermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelinePermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelinePermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []PipelineAccessControlRequest
	// The pipeline for which to get or manage permissions.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func pipelinePermissionsRequestToPb(st *PipelinePermissionsRequest) (*pipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsRequestPb{}

	var accessControlListPb []pipelineAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := pipelineAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.PipelineId = st.PipelineId

	return pb, nil
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

type pipelinePermissionsRequestPb struct {
	AccessControlList []pipelineAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The pipeline for which to get or manage permissions.
	PipelineId string `json:"-" url:"-"`
}

func pipelinePermissionsRequestFromPb(pb *pipelinePermissionsRequestPb) (*PipelinePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissionsRequest{}

	var accessControlListField []PipelineAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := pipelineAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.PipelineId = pb.PipelineId

	return st, nil
}

type PipelineSpec struct {
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger *PipelineTrigger

	ForceSendFields []string `tf:"-"`
}

func pipelineSpecToPb(st *PipelineSpec) (*pipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineSpecPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Catalog = st.Catalog

	pb.Channel = st.Channel

	var clustersPb []pipelineClusterPb
	for _, item := range st.Clusters {
		itemPb, err := pipelineClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb

	pb.Configuration = st.Configuration

	pb.Continuous = st.Continuous

	deploymentPb, err := pipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	pb.Development = st.Development

	pb.Edition = st.Edition

	eventLogPb, err := eventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}

	filtersPb, err := filtersToPb(st.Filters)
	if err != nil {
		return nil, err
	}
	if filtersPb != nil {
		pb.Filters = filtersPb
	}

	gatewayDefinitionPb, err := ingestionGatewayPipelineDefinitionToPb(st.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionPb != nil {
		pb.GatewayDefinition = gatewayDefinitionPb
	}

	pb.Id = st.Id

	ingestionDefinitionPb, err := ingestionPipelineDefinitionToPb(st.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionPb != nil {
		pb.IngestionDefinition = ingestionDefinitionPb
	}

	var librariesPb []pipelineLibraryPb
	for _, item := range st.Libraries {
		itemPb, err := pipelineLibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	pb.Name = st.Name

	var notificationsPb []notificationsPb
	for _, item := range st.Notifications {
		itemPb, err := notificationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notificationsPb = append(notificationsPb, *itemPb)
		}
	}
	pb.Notifications = notificationsPb

	pb.Photon = st.Photon

	restartWindowPb, err := restartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}

	pb.Schema = st.Schema

	pb.Serverless = st.Serverless

	pb.Storage = st.Storage

	pb.Target = st.Target

	triggerPb, err := pipelineTriggerToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineSpecPb struct {
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
	Clusters []pipelineClusterPb `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Deployment type of this pipeline.
	Deployment *pipelineDeploymentPb `json:"deployment,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Event log configuration for this pipeline
	EventLog *eventLogSpecPb `json:"event_log,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *filtersPb `json:"filters,omitempty"`
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *ingestionGatewayPipelineDefinitionPb `json:"gateway_definition,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *ingestionPipelineDefinitionPb `json:"ingestion_definition,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []pipelineLibraryPb `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// List of notification settings for this pipeline.
	Notifications []notificationsPb `json:"notifications,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Restart window of this pipeline.
	RestartWindow *restartWindowPb `json:"restart_window,omitempty"`
	// The default schema (database) where tables are read from or published to.
	Schema string `json:"schema,omitempty"`
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool `json:"serverless,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *pipelineTriggerPb `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineSpecFromPb(pb *pipelineSpecPb) (*PipelineSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineSpec{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel

	var clustersField []PipelineCluster
	for _, itemPb := range pb.Clusters {
		item, err := pipelineClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	st.Configuration = pb.Configuration
	st.Continuous = pb.Continuous
	deploymentField, err := pipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Development = pb.Development
	st.Edition = pb.Edition
	eventLogField, err := eventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	filtersField, err := filtersFromPb(pb.Filters)
	if err != nil {
		return nil, err
	}
	if filtersField != nil {
		st.Filters = filtersField
	}
	gatewayDefinitionField, err := ingestionGatewayPipelineDefinitionFromPb(pb.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionField != nil {
		st.GatewayDefinition = gatewayDefinitionField
	}
	st.Id = pb.Id
	ingestionDefinitionField, err := ingestionPipelineDefinitionFromPb(pb.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionField != nil {
		st.IngestionDefinition = ingestionDefinitionField
	}

	var librariesField []PipelineLibrary
	for _, itemPb := range pb.Libraries {
		item, err := pipelineLibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.Name = pb.Name

	var notificationsField []Notifications
	for _, itemPb := range pb.Notifications {
		item, err := notificationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notificationsField = append(notificationsField, *item)
		}
	}
	st.Notifications = notificationsField
	st.Photon = pb.Photon
	restartWindowField, err := restartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Target = pb.Target
	triggerField, err := pipelineTriggerFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The pipeline state.
type PipelineState string
type pipelineStatePb string

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

func pipelineStateToPb(st *PipelineState) (*pipelineStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelineStatePb(*st)
	return &pb, nil
}

func pipelineStateFromPb(pb *pipelineStatePb) (*PipelineState, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineState(*pb)
	return &st, nil
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	// Wire name: 'cluster_id'
	ClusterId string
	// The username of the pipeline creator.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// The health of a pipeline.
	// Wire name: 'health'
	Health PipelineStateInfoHealth
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	// Wire name: 'latest_updates'
	LatestUpdates []UpdateStateInfo
	// The user-friendly name of the pipeline.
	// Wire name: 'name'
	Name string
	// The unique identifier of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	// Wire name: 'run_as_user_name'
	RunAsUserName string
	// The pipeline state.
	// Wire name: 'state'
	State PipelineState

	ForceSendFields []string `tf:"-"`
}

func pipelineStateInfoToPb(st *PipelineStateInfo) (*pipelineStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineStateInfoPb{}
	pb.ClusterId = st.ClusterId

	pb.CreatorUserName = st.CreatorUserName

	pb.Health = st.Health

	var latestUpdatesPb []updateStateInfoPb
	for _, item := range st.LatestUpdates {
		itemPb, err := updateStateInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			latestUpdatesPb = append(latestUpdatesPb, *itemPb)
		}
	}
	pb.LatestUpdates = latestUpdatesPb

	pb.Name = st.Name

	pb.PipelineId = st.PipelineId

	pb.RunAsUserName = st.RunAsUserName

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pipelineStateInfoPb struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	Health PipelineStateInfoHealth `json:"health,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []updateStateInfoPb `json:"latest_updates,omitempty"`
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

func pipelineStateInfoFromPb(pb *pipelineStateInfoPb) (*PipelineStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineStateInfo{}
	st.ClusterId = pb.ClusterId
	st.CreatorUserName = pb.CreatorUserName
	st.Health = pb.Health

	var latestUpdatesField []UpdateStateInfo
	for _, itemPb := range pb.LatestUpdates {
		item, err := updateStateInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			latestUpdatesField = append(latestUpdatesField, *item)
		}
	}
	st.LatestUpdates = latestUpdatesField
	st.Name = pb.Name
	st.PipelineId = pb.PipelineId
	st.RunAsUserName = pb.RunAsUserName
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The health of a pipeline.
type PipelineStateInfoHealth string
type pipelineStateInfoHealthPb string

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

func pipelineStateInfoHealthToPb(st *PipelineStateInfoHealth) (*pipelineStateInfoHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelineStateInfoHealthPb(*st)
	return &pb, nil
}

func pipelineStateInfoHealthFromPb(pb *pipelineStateInfoHealthPb) (*PipelineStateInfoHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineStateInfoHealth(*pb)
	return &st, nil
}

type PipelineTrigger struct {

	// Wire name: 'cron'
	Cron *CronTrigger

	// Wire name: 'manual'
	Manual *ManualTrigger
}

func pipelineTriggerToPb(st *PipelineTrigger) (*pipelineTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineTriggerPb{}
	cronPb, err := cronTriggerToPb(st.Cron)
	if err != nil {
		return nil, err
	}
	if cronPb != nil {
		pb.Cron = cronPb
	}

	manualPb, err := manualTriggerToPb(st.Manual)
	if err != nil {
		return nil, err
	}
	if manualPb != nil {
		pb.Manual = manualPb
	}

	return pb, nil
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

type pipelineTriggerPb struct {
	Cron *cronTriggerPb `json:"cron,omitempty"`

	Manual *manualTriggerPb `json:"manual,omitempty"`
}

func pipelineTriggerFromPb(pb *pipelineTriggerPb) (*PipelineTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineTrigger{}
	cronField, err := cronTriggerFromPb(pb.Cron)
	if err != nil {
		return nil, err
	}
	if cronField != nil {
		st.Cron = cronField
	}
	manualField, err := manualTriggerFromPb(pb.Manual)
	if err != nil {
		return nil, err
	}
	if manualField != nil {
		st.Manual = manualField
	}

	return st, nil
}

type ReportSpec struct {
	// Required. Destination catalog to store table.
	// Wire name: 'destination_catalog'
	DestinationCatalog string
	// Required. Destination schema to store table.
	// Wire name: 'destination_schema'
	DestinationSchema string
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	// Wire name: 'destination_table'
	DestinationTable string
	// Required. Report URL in the source system.
	// Wire name: 'source_url'
	SourceUrl string
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string `tf:"-"`
}

func reportSpecToPb(st *ReportSpec) (*reportSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &reportSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog

	pb.DestinationSchema = st.DestinationSchema

	pb.DestinationTable = st.DestinationTable

	pb.SourceUrl = st.SourceUrl

	tableConfigurationPb, err := tableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type reportSpecPb struct {
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
	TableConfiguration *tableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func reportSpecFromPb(pb *reportSpecPb) (*ReportSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReportSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.DestinationTable = pb.DestinationTable
	st.SourceUrl = pb.SourceUrl
	tableConfigurationField, err := tableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *reportSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st reportSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RestartWindow struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	// Wire name: 'days_of_week'
	DaysOfWeek []DayOfWeek
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	// Wire name: 'start_hour'
	StartHour int
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	// Wire name: 'time_zone_id'
	TimeZoneId string

	ForceSendFields []string `tf:"-"`
}

func restartWindowToPb(st *RestartWindow) (*restartWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restartWindowPb{}
	pb.DaysOfWeek = st.DaysOfWeek

	pb.StartHour = st.StartHour

	pb.TimeZoneId = st.TimeZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type restartWindowPb struct {
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

func restartWindowFromPb(pb *restartWindowPb) (*RestartWindow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartWindow{}
	st.DaysOfWeek = pb.DaysOfWeek
	st.StartHour = pb.StartHour
	st.TimeZoneId = pb.TimeZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restartWindowPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restartWindowPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	ServicePrincipalName string
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func runAsToPb(st *RunAs) (*runAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runAsPb{}
	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type runAsPb struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runAsFromPb(pb *runAsPb) (*RunAs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunAs{}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runAsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runAsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	// Wire name: 'destination_catalog'
	DestinationCatalog string
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	// Wire name: 'destination_schema'
	DestinationSchema string
	// The source catalog name. Might be optional depending on the type of
	// source.
	// Wire name: 'source_catalog'
	SourceCatalog string
	// Required. Schema name in the source database.
	// Wire name: 'source_schema'
	SourceSchema string
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string `tf:"-"`
}

func schemaSpecToPb(st *SchemaSpec) (*schemaSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schemaSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog

	pb.DestinationSchema = st.DestinationSchema

	pb.SourceCatalog = st.SourceCatalog

	pb.SourceSchema = st.SourceSchema

	tableConfigurationPb, err := tableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type schemaSpecPb struct {
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
	TableConfiguration *tableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func schemaSpecFromPb(pb *schemaSpecPb) (*SchemaSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SchemaSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.SourceCatalog = pb.SourceCatalog
	st.SourceSchema = pb.SourceSchema
	tableConfigurationField, err := tableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *schemaSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st schemaSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Sequencing struct {
	// A sequence number, unique and increasing within the control plane.
	// Wire name: 'control_plane_seq_no'
	ControlPlaneSeqNo int
	// the ID assigned by the data plane.
	// Wire name: 'data_plane_id'
	DataPlaneId *DataPlaneId

	ForceSendFields []string `tf:"-"`
}

func sequencingToPb(st *Sequencing) (*sequencingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sequencingPb{}
	pb.ControlPlaneSeqNo = st.ControlPlaneSeqNo

	dataPlaneIdPb, err := dataPlaneIdToPb(st.DataPlaneId)
	if err != nil {
		return nil, err
	}
	if dataPlaneIdPb != nil {
		pb.DataPlaneId = dataPlaneIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type sequencingPb struct {
	// A sequence number, unique and increasing within the control plane.
	ControlPlaneSeqNo int `json:"control_plane_seq_no,omitempty"`
	// the ID assigned by the data plane.
	DataPlaneId *dataPlaneIdPb `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sequencingFromPb(pb *sequencingPb) (*Sequencing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Sequencing{}
	st.ControlPlaneSeqNo = pb.ControlPlaneSeqNo
	dataPlaneIdField, err := dataPlaneIdFromPb(pb.DataPlaneId)
	if err != nil {
		return nil, err
	}
	if dataPlaneIdField != nil {
		st.DataPlaneId = dataPlaneIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sequencingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sequencingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SerializedException struct {
	// Runtime class of the exception
	// Wire name: 'class_name'
	ClassName string
	// Exception message
	// Wire name: 'message'
	Message string
	// Stack trace consisting of a list of stack frames
	// Wire name: 'stack'
	Stack []StackFrame

	ForceSendFields []string `tf:"-"`
}

func serializedExceptionToPb(st *SerializedException) (*serializedExceptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serializedExceptionPb{}
	pb.ClassName = st.ClassName

	pb.Message = st.Message

	var stackPb []stackFramePb
	for _, item := range st.Stack {
		itemPb, err := stackFrameToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			stackPb = append(stackPb, *itemPb)
		}
	}
	pb.Stack = stackPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type serializedExceptionPb struct {
	// Runtime class of the exception
	ClassName string `json:"class_name,omitempty"`
	// Exception message
	Message string `json:"message,omitempty"`
	// Stack trace consisting of a list of stack frames
	Stack []stackFramePb `json:"stack,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func serializedExceptionFromPb(pb *serializedExceptionPb) (*SerializedException, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SerializedException{}
	st.ClassName = pb.ClassName
	st.Message = pb.Message

	var stackField []StackFrame
	for _, itemPb := range pb.Stack {
		item, err := stackFrameFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			stackField = append(stackField, *item)
		}
	}
	st.Stack = stackField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *serializedExceptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st serializedExceptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StackFrame struct {
	// Class from which the method call originated
	// Wire name: 'declaring_class'
	DeclaringClass string
	// File where the method is defined
	// Wire name: 'file_name'
	FileName string
	// Line from which the method was called
	// Wire name: 'line_number'
	LineNumber int
	// Name of the method which was called
	// Wire name: 'method_name'
	MethodName string

	ForceSendFields []string `tf:"-"`
}

func stackFrameToPb(st *StackFrame) (*stackFramePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stackFramePb{}
	pb.DeclaringClass = st.DeclaringClass

	pb.FileName = st.FileName

	pb.LineNumber = st.LineNumber

	pb.MethodName = st.MethodName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type stackFramePb struct {
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

func stackFrameFromPb(pb *stackFramePb) (*StackFrame, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StackFrame{}
	st.DeclaringClass = pb.DeclaringClass
	st.FileName = pb.FileName
	st.LineNumber = pb.LineNumber
	st.MethodName = pb.MethodName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *stackFramePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st stackFramePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartUpdate struct {

	// Wire name: 'cause'
	Cause StartUpdateCause
	// If true, this update will reset all tables before running.
	// Wire name: 'full_refresh'
	FullRefresh bool
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'full_refresh_selection'
	FullRefreshSelection []string

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'refresh_selection'
	RefreshSelection []string
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	// Wire name: 'validate_only'
	ValidateOnly bool

	ForceSendFields []string `tf:"-"`
}

func startUpdateToPb(st *StartUpdate) (*startUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startUpdatePb{}
	pb.Cause = st.Cause

	pb.FullRefresh = st.FullRefresh

	pb.FullRefreshSelection = st.FullRefreshSelection

	pb.PipelineId = st.PipelineId

	pb.RefreshSelection = st.RefreshSelection

	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type startUpdatePb struct {
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

func startUpdateFromPb(pb *startUpdatePb) (*StartUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartUpdate{}
	st.Cause = pb.Cause
	st.FullRefresh = pb.FullRefresh
	st.FullRefreshSelection = pb.FullRefreshSelection
	st.PipelineId = pb.PipelineId
	st.RefreshSelection = pb.RefreshSelection
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *startUpdatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st startUpdatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartUpdateCause string
type startUpdateCausePb string

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

func startUpdateCauseToPb(st *StartUpdateCause) (*startUpdateCausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := startUpdateCausePb(*st)
	return &pb, nil
}

func startUpdateCauseFromPb(pb *startUpdateCausePb) (*StartUpdateCause, error) {
	if pb == nil {
		return nil, nil
	}
	st := StartUpdateCause(*pb)
	return &st, nil
}

type StartUpdateResponse struct {

	// Wire name: 'update_id'
	UpdateId string

	ForceSendFields []string `tf:"-"`
}

func startUpdateResponseToPb(st *StartUpdateResponse) (*startUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startUpdateResponsePb{}
	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type startUpdateResponsePb struct {
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func startUpdateResponseFromPb(pb *startUpdateResponsePb) (*StartUpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartUpdateResponse{}
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *startUpdateResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st startUpdateResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StopPipelineResponse struct {
}

func stopPipelineResponseToPb(st *StopPipelineResponse) (*stopPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopPipelineResponsePb{}

	return pb, nil
}

func (st *StopPipelineResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stopPipelineResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stopPipelineResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StopPipelineResponse) MarshalJSON() ([]byte, error) {
	pb, err := stopPipelineResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type stopPipelineResponsePb struct {
}

func stopPipelineResponseFromPb(pb *stopPipelineResponsePb) (*StopPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopPipelineResponse{}

	return st, nil
}

// Stop a pipeline
type StopRequest struct {

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func stopRequestToPb(st *StopRequest) (*stopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
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

type stopRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func stopRequestFromPb(pb *stopRequestPb) (*StopRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	// Wire name: 'destination_catalog'
	DestinationCatalog string
	// Required. Destination schema to store table.
	// Wire name: 'destination_schema'
	DestinationSchema string
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	// Wire name: 'destination_table'
	DestinationTable string
	// Source catalog name. Might be optional depending on the type of source.
	// Wire name: 'source_catalog'
	SourceCatalog string
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	// Wire name: 'source_schema'
	SourceSchema string
	// Required. Table name in the source database.
	// Wire name: 'source_table'
	SourceTable string
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string `tf:"-"`
}

func tableSpecToPb(st *TableSpec) (*tableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog

	pb.DestinationSchema = st.DestinationSchema

	pb.DestinationTable = st.DestinationTable

	pb.SourceCatalog = st.SourceCatalog

	pb.SourceSchema = st.SourceSchema

	pb.SourceTable = st.SourceTable

	tableConfigurationPb, err := tableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tableSpecPb struct {
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
	TableConfiguration *tableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableSpecFromPb(pb *tableSpecPb) (*TableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.DestinationTable = pb.DestinationTable
	st.SourceCatalog = pb.SourceCatalog
	st.SourceSchema = pb.SourceSchema
	st.SourceTable = pb.SourceTable
	tableConfigurationField, err := tableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableSpecificConfig struct {
	// The primary key of the table used to apply changes.
	// Wire name: 'primary_keys'
	PrimaryKeys []string
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	// Wire name: 'salesforce_include_formula_fields'
	SalesforceIncludeFormulaFields bool
	// The SCD type to use to ingest the table.
	// Wire name: 'scd_type'
	ScdType TableSpecificConfigScdType
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	// Wire name: 'sequence_by'
	SequenceBy []string

	ForceSendFields []string `tf:"-"`
}

func tableSpecificConfigToPb(st *TableSpecificConfig) (*tableSpecificConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSpecificConfigPb{}
	pb.PrimaryKeys = st.PrimaryKeys

	pb.SalesforceIncludeFormulaFields = st.SalesforceIncludeFormulaFields

	pb.ScdType = st.ScdType

	pb.SequenceBy = st.SequenceBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tableSpecificConfigPb struct {
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

func tableSpecificConfigFromPb(pb *tableSpecificConfigPb) (*TableSpecificConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSpecificConfig{}
	st.PrimaryKeys = pb.PrimaryKeys
	st.SalesforceIncludeFormulaFields = pb.SalesforceIncludeFormulaFields
	st.ScdType = pb.ScdType
	st.SequenceBy = pb.SequenceBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableSpecificConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableSpecificConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The SCD type to use to ingest the table.
type TableSpecificConfigScdType string
type tableSpecificConfigScdTypePb string

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

func tableSpecificConfigScdTypeToPb(st *TableSpecificConfigScdType) (*tableSpecificConfigScdTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := tableSpecificConfigScdTypePb(*st)
	return &pb, nil
}

func tableSpecificConfigScdTypeFromPb(pb *tableSpecificConfigScdTypePb) (*TableSpecificConfigScdType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableSpecificConfigScdType(*pb)
	return &st, nil
}

type UpdateInfo struct {
	// What triggered this update.
	// Wire name: 'cause'
	Cause UpdateInfoCause
	// The ID of the cluster that the update is running on.
	// Wire name: 'cluster_id'
	ClusterId string
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	// Wire name: 'config'
	Config *PipelineSpec
	// The time when this update was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// If true, this update will reset all tables before running.
	// Wire name: 'full_refresh'
	FullRefresh bool
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'full_refresh_selection'
	FullRefreshSelection []string
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'refresh_selection'
	RefreshSelection []string
	// The update state.
	// Wire name: 'state'
	State UpdateInfoState
	// The ID of this update.
	// Wire name: 'update_id'
	UpdateId string
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	// Wire name: 'validate_only'
	ValidateOnly bool

	ForceSendFields []string `tf:"-"`
}

func updateInfoToPb(st *UpdateInfo) (*updateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInfoPb{}
	pb.Cause = st.Cause

	pb.ClusterId = st.ClusterId

	configPb, err := pipelineSpecToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	pb.CreationTime = st.CreationTime

	pb.FullRefresh = st.FullRefresh

	pb.FullRefreshSelection = st.FullRefreshSelection

	pb.PipelineId = st.PipelineId

	pb.RefreshSelection = st.RefreshSelection

	pb.State = st.State

	pb.UpdateId = st.UpdateId

	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateInfoPb struct {
	// What triggered this update.
	Cause UpdateInfoCause `json:"cause,omitempty"`
	// The ID of the cluster that the update is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config *pipelineSpecPb `json:"config,omitempty"`
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

func updateInfoFromPb(pb *updateInfoPb) (*UpdateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInfo{}
	st.Cause = pb.Cause
	st.ClusterId = pb.ClusterId
	configField, err := pipelineSpecFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	st.CreationTime = pb.CreationTime
	st.FullRefresh = pb.FullRefresh
	st.FullRefreshSelection = pb.FullRefreshSelection
	st.PipelineId = pb.PipelineId
	st.RefreshSelection = pb.RefreshSelection
	st.State = pb.State
	st.UpdateId = pb.UpdateId
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// What triggered this update.
type UpdateInfoCause string
type updateInfoCausePb string

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

func updateInfoCauseToPb(st *UpdateInfoCause) (*updateInfoCausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateInfoCausePb(*st)
	return &pb, nil
}

func updateInfoCauseFromPb(pb *updateInfoCausePb) (*UpdateInfoCause, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateInfoCause(*pb)
	return &st, nil
}

// The update state.
type UpdateInfoState string
type updateInfoStatePb string

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

func updateInfoStateToPb(st *UpdateInfoState) (*updateInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateInfoStatePb(*st)
	return &pb, nil
}

func updateInfoStateFromPb(pb *updateInfoStatePb) (*UpdateInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateInfoState(*pb)
	return &st, nil
}

type UpdateStateInfo struct {

	// Wire name: 'creation_time'
	CreationTime string

	// Wire name: 'state'
	State UpdateStateInfoState

	// Wire name: 'update_id'
	UpdateId string

	ForceSendFields []string `tf:"-"`
}

func updateStateInfoToPb(st *UpdateStateInfo) (*updateStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateStateInfoPb{}
	pb.CreationTime = st.CreationTime

	pb.State = st.State

	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateStateInfoPb struct {
	CreationTime string `json:"creation_time,omitempty"`

	State UpdateStateInfoState `json:"state,omitempty"`

	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateStateInfoFromPb(pb *updateStateInfoPb) (*UpdateStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStateInfo{}
	st.CreationTime = pb.CreationTime
	st.State = pb.State
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateStateInfoState string
type updateStateInfoStatePb string

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

func updateStateInfoStateToPb(st *UpdateStateInfoState) (*updateStateInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateStateInfoStatePb(*st)
	return &pb, nil
}

func updateStateInfoStateFromPb(pb *updateStateInfoStatePb) (*UpdateStateInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateStateInfoState(*pb)
	return &st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
