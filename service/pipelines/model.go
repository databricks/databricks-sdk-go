// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
	"github.com/databricks/databricks-sdk-go/service/pipelines/pipelinespb"
)

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	// Wire name: 'allow_duplicate_names'
	AllowDuplicateNames bool ``
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string ``
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string ``
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster ``
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string ``
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool ``
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment ``
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool ``

	// Wire name: 'dry_run'
	DryRun bool ``
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string ``
	// Environment specification for this pipeline used to install dependencies.
	// Wire name: 'environment'
	Environment *PipelinesEnvironment ``
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec ``
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters ``
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition ``
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string ``
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition ``
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary ``
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string ``
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications ``
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool ``
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow ``
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	// Wire name: 'root_path'
	RootPath string ``

	// Wire name: 'run_as'
	RunAs *RunAs ``
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string ``
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool ``
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string ``
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	// Wire name: 'tags'
	Tags map[string]string ``
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string ``
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger         *PipelineTrigger ``
	ForceSendFields []string         `tf:"-"`
}

func (s *CreatePipeline) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePipeline) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreatePipelineToPb(st *CreatePipeline) (*pipelinespb.CreatePipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.CreatePipelinePb{}
	pb.AllowDuplicateNames = st.AllowDuplicateNames
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Catalog = st.Catalog
	pb.Channel = st.Channel

	var clustersPb []pipelinespb.PipelineClusterPb
	for _, item := range st.Clusters {
		itemPb, err := PipelineClusterToPb(&item)
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
	deploymentPb, err := PipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}
	pb.Development = st.Development
	pb.DryRun = st.DryRun
	pb.Edition = st.Edition
	environmentPb, err := PipelinesEnvironmentToPb(st.Environment)
	if err != nil {
		return nil, err
	}
	if environmentPb != nil {
		pb.Environment = environmentPb
	}
	eventLogPb, err := EventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}
	filtersPb, err := FiltersToPb(st.Filters)
	if err != nil {
		return nil, err
	}
	if filtersPb != nil {
		pb.Filters = filtersPb
	}
	gatewayDefinitionPb, err := IngestionGatewayPipelineDefinitionToPb(st.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionPb != nil {
		pb.GatewayDefinition = gatewayDefinitionPb
	}
	pb.Id = st.Id
	ingestionDefinitionPb, err := IngestionPipelineDefinitionToPb(st.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionPb != nil {
		pb.IngestionDefinition = ingestionDefinitionPb
	}

	var librariesPb []pipelinespb.PipelineLibraryPb
	for _, item := range st.Libraries {
		itemPb, err := PipelineLibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.Name = st.Name

	var notificationsPb []pipelinespb.NotificationsPb
	for _, item := range st.Notifications {
		itemPb, err := NotificationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notificationsPb = append(notificationsPb, *itemPb)
		}
	}
	pb.Notifications = notificationsPb
	pb.Photon = st.Photon
	restartWindowPb, err := RestartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}
	pb.RootPath = st.RootPath
	runAsPb, err := RunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}
	pb.Schema = st.Schema
	pb.Serverless = st.Serverless
	pb.Storage = st.Storage
	pb.Tags = st.Tags
	pb.Target = st.Target
	triggerPb, err := PipelineTriggerToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreatePipelineFromPb(pb *pipelinespb.CreatePipelinePb) (*CreatePipeline, error) {
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
		item, err := PipelineClusterFromPb(&itemPb)
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
	deploymentField, err := PipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Development = pb.Development
	st.DryRun = pb.DryRun
	st.Edition = pb.Edition
	environmentField, err := PipelinesEnvironmentFromPb(pb.Environment)
	if err != nil {
		return nil, err
	}
	if environmentField != nil {
		st.Environment = environmentField
	}
	eventLogField, err := EventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	filtersField, err := FiltersFromPb(pb.Filters)
	if err != nil {
		return nil, err
	}
	if filtersField != nil {
		st.Filters = filtersField
	}
	gatewayDefinitionField, err := IngestionGatewayPipelineDefinitionFromPb(pb.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionField != nil {
		st.GatewayDefinition = gatewayDefinitionField
	}
	st.Id = pb.Id
	ingestionDefinitionField, err := IngestionPipelineDefinitionFromPb(pb.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionField != nil {
		st.IngestionDefinition = ingestionDefinitionField
	}

	var librariesField []PipelineLibrary
	for _, itemPb := range pb.Libraries {
		item, err := PipelineLibraryFromPb(&itemPb)
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
		item, err := NotificationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notificationsField = append(notificationsField, *item)
		}
	}
	st.Notifications = notificationsField
	st.Photon = pb.Photon
	restartWindowField, err := RestartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	st.RootPath = pb.RootPath
	runAsField, err := RunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Tags = pb.Tags
	st.Target = pb.Target
	triggerField, err := PipelineTriggerFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true.
	// Wire name: 'effective_settings'
	EffectiveSettings *PipelineSpec ``
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	// Wire name: 'pipeline_id'
	PipelineId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreatePipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreatePipelineResponseToPb(st *CreatePipelineResponse) (*pipelinespb.CreatePipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.CreatePipelineResponsePb{}
	effectiveSettingsPb, err := PipelineSpecToPb(st.EffectiveSettings)
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

func CreatePipelineResponseFromPb(pb *pipelinespb.CreatePipelineResponsePb) (*CreatePipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePipelineResponse{}
	effectiveSettingsField, err := PipelineSpecFromPb(pb.EffectiveSettings)
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

type CronTrigger struct {

	// Wire name: 'quartz_cron_schedule'
	QuartzCronSchedule string ``

	// Wire name: 'timezone_id'
	TimezoneId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CronTrigger) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CronTrigger) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CronTriggerToPb(st *CronTrigger) (*pipelinespb.CronTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.CronTriggerPb{}
	pb.QuartzCronSchedule = st.QuartzCronSchedule
	pb.TimezoneId = st.TimezoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CronTriggerFromPb(pb *pipelinespb.CronTriggerPb) (*CronTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronTrigger{}
	st.QuartzCronSchedule = pb.QuartzCronSchedule
	st.TimezoneId = pb.TimezoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DataPlaneId struct {
	// The instance name of the data plane emitting an event.
	// Wire name: 'instance'
	Instance string ``
	// A sequence number, unique and increasing within the data plane instance.
	// Wire name: 'seq_no'
	SeqNo           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *DataPlaneId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DataPlaneIdToPb(st *DataPlaneId) (*pipelinespb.DataPlaneIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.DataPlaneIdPb{}
	pb.Instance = st.Instance
	pb.SeqNo = st.SeqNo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DataPlaneIdFromPb(pb *pipelinespb.DataPlaneIdPb) (*DataPlaneId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneId{}
	st.Instance = pb.Instance
	st.SeqNo = pb.SeqNo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func DayOfWeekToPb(st *DayOfWeek) (*pipelinespb.DayOfWeekPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.DayOfWeekPb(*st)
	return &pb, nil
}

func DayOfWeekFromPb(pb *pipelinespb.DayOfWeekPb) (*DayOfWeek, error) {
	if pb == nil {
		return nil, nil
	}
	st := DayOfWeek(*pb)
	return &st, nil
}

type DeletePipelineRequest struct {

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func DeletePipelineRequestToPb(st *DeletePipelineRequest) (*pipelinespb.DeletePipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.DeletePipelineRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

func DeletePipelineRequestFromPb(pb *pipelinespb.DeletePipelineRequestPb) (*DeletePipelineRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePipelineRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
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

func DeploymentKindToPb(st *DeploymentKind) (*pipelinespb.DeploymentKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.DeploymentKindPb(*st)
	return &pb, nil
}

func DeploymentKindFromPb(pb *pipelinespb.DeploymentKindPb) (*DeploymentKind, error) {
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
	AllowDuplicateNames bool ``
	// Budget policy of this pipeline.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string ``
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string ``
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster ``
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string ``
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool ``
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment ``
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool ``
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string ``
	// Environment specification for this pipeline used to install dependencies.
	// Wire name: 'environment'
	Environment *PipelinesEnvironment ``
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec ``
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	// Wire name: 'expected_last_modified'
	ExpectedLastModified int64 ``
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters ``
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition ``
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string ``
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition ``
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary ``
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string ``
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications ``
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool ``
	// Unique identifier for this pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow ``
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	// Wire name: 'root_path'
	RootPath string ``

	// Wire name: 'run_as'
	RunAs *RunAs ``
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string ``
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool ``
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string ``
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	// Wire name: 'tags'
	Tags map[string]string ``
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string ``
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger         *PipelineTrigger ``
	ForceSendFields []string         `tf:"-"`
}

func (s *EditPipeline) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditPipeline) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EditPipelineToPb(st *EditPipeline) (*pipelinespb.EditPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.EditPipelinePb{}
	pb.AllowDuplicateNames = st.AllowDuplicateNames
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Catalog = st.Catalog
	pb.Channel = st.Channel

	var clustersPb []pipelinespb.PipelineClusterPb
	for _, item := range st.Clusters {
		itemPb, err := PipelineClusterToPb(&item)
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
	deploymentPb, err := PipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}
	pb.Development = st.Development
	pb.Edition = st.Edition
	environmentPb, err := PipelinesEnvironmentToPb(st.Environment)
	if err != nil {
		return nil, err
	}
	if environmentPb != nil {
		pb.Environment = environmentPb
	}
	eventLogPb, err := EventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}
	pb.ExpectedLastModified = st.ExpectedLastModified
	filtersPb, err := FiltersToPb(st.Filters)
	if err != nil {
		return nil, err
	}
	if filtersPb != nil {
		pb.Filters = filtersPb
	}
	gatewayDefinitionPb, err := IngestionGatewayPipelineDefinitionToPb(st.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionPb != nil {
		pb.GatewayDefinition = gatewayDefinitionPb
	}
	pb.Id = st.Id
	ingestionDefinitionPb, err := IngestionPipelineDefinitionToPb(st.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionPb != nil {
		pb.IngestionDefinition = ingestionDefinitionPb
	}

	var librariesPb []pipelinespb.PipelineLibraryPb
	for _, item := range st.Libraries {
		itemPb, err := PipelineLibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.Name = st.Name

	var notificationsPb []pipelinespb.NotificationsPb
	for _, item := range st.Notifications {
		itemPb, err := NotificationsToPb(&item)
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
	restartWindowPb, err := RestartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}
	pb.RootPath = st.RootPath
	runAsPb, err := RunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}
	pb.Schema = st.Schema
	pb.Serverless = st.Serverless
	pb.Storage = st.Storage
	pb.Tags = st.Tags
	pb.Target = st.Target
	triggerPb, err := PipelineTriggerToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EditPipelineFromPb(pb *pipelinespb.EditPipelinePb) (*EditPipeline, error) {
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
		item, err := PipelineClusterFromPb(&itemPb)
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
	deploymentField, err := PipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Development = pb.Development
	st.Edition = pb.Edition
	environmentField, err := PipelinesEnvironmentFromPb(pb.Environment)
	if err != nil {
		return nil, err
	}
	if environmentField != nil {
		st.Environment = environmentField
	}
	eventLogField, err := EventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	st.ExpectedLastModified = pb.ExpectedLastModified
	filtersField, err := FiltersFromPb(pb.Filters)
	if err != nil {
		return nil, err
	}
	if filtersField != nil {
		st.Filters = filtersField
	}
	gatewayDefinitionField, err := IngestionGatewayPipelineDefinitionFromPb(pb.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionField != nil {
		st.GatewayDefinition = gatewayDefinitionField
	}
	st.Id = pb.Id
	ingestionDefinitionField, err := IngestionPipelineDefinitionFromPb(pb.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionField != nil {
		st.IngestionDefinition = ingestionDefinitionField
	}

	var librariesField []PipelineLibrary
	for _, itemPb := range pb.Libraries {
		item, err := PipelineLibraryFromPb(&itemPb)
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
		item, err := NotificationsFromPb(&itemPb)
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
	restartWindowField, err := RestartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	st.RootPath = pb.RootPath
	runAsField, err := RunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Tags = pb.Tags
	st.Target = pb.Target
	triggerField, err := PipelineTriggerFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ErrorDetail struct {
	// The exception thrown for this error, with its chain of cause.
	// Wire name: 'exceptions'
	Exceptions []SerializedException ``
	// Whether this error is considered fatal, that is, unrecoverable.
	// Wire name: 'fatal'
	Fatal           bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *ErrorDetail) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ErrorDetail) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ErrorDetailToPb(st *ErrorDetail) (*pipelinespb.ErrorDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ErrorDetailPb{}

	var exceptionsPb []pipelinespb.SerializedExceptionPb
	for _, item := range st.Exceptions {
		itemPb, err := SerializedExceptionToPb(&item)
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

func ErrorDetailFromPb(pb *pipelinespb.ErrorDetailPb) (*ErrorDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ErrorDetail{}

	var exceptionsField []SerializedException
	for _, itemPb := range pb.Exceptions {
		item, err := SerializedExceptionFromPb(&itemPb)
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

func EventLevelToPb(st *EventLevel) (*pipelinespb.EventLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.EventLevelPb(*st)
	return &pb, nil
}

func EventLevelFromPb(pb *pipelinespb.EventLevelPb) (*EventLevel, error) {
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
	Catalog string ``
	// The name the event log is published to in UC.
	// Wire name: 'name'
	Name string ``
	// The UC schema the event log is published under.
	// Wire name: 'schema'
	Schema          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EventLogSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EventLogSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EventLogSpecToPb(st *EventLogSpec) (*pipelinespb.EventLogSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.EventLogSpecPb{}
	pb.Catalog = st.Catalog
	pb.Name = st.Name
	pb.Schema = st.Schema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EventLogSpecFromPb(pb *pipelinespb.EventLogSpecPb) (*EventLogSpec, error) {
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

type FileLibrary struct {
	// The absolute path of the source code.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *FileLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FileLibraryToPb(st *FileLibrary) (*pipelinespb.FileLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.FileLibraryPb{}
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FileLibraryFromPb(pb *pipelinespb.FileLibraryPb) (*FileLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileLibrary{}
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Filters struct {
	// Paths to exclude.
	// Wire name: 'exclude'
	Exclude []string ``
	// Paths to include.
	// Wire name: 'include'
	Include []string ``
}

func FiltersToPb(st *Filters) (*pipelinespb.FiltersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.FiltersPb{}
	pb.Exclude = st.Exclude
	pb.Include = st.Include

	return pb, nil
}

func FiltersFromPb(pb *pipelinespb.FiltersPb) (*Filters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Filters{}
	st.Exclude = pb.Exclude
	st.Include = pb.Include

	return st, nil
}

type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func GetPipelinePermissionLevelsRequestToPb(st *GetPipelinePermissionLevelsRequest) (*pipelinespb.GetPipelinePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetPipelinePermissionLevelsRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

func GetPipelinePermissionLevelsRequestFromPb(pb *pipelinespb.GetPipelinePermissionLevelsRequestPb) (*GetPipelinePermissionLevelsRequest, error) {
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
	PermissionLevels []PipelinePermissionsDescription ``
}

func GetPipelinePermissionLevelsResponseToPb(st *GetPipelinePermissionLevelsResponse) (*pipelinespb.GetPipelinePermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetPipelinePermissionLevelsResponsePb{}

	var permissionLevelsPb []pipelinespb.PipelinePermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := PipelinePermissionsDescriptionToPb(&item)
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

func GetPipelinePermissionLevelsResponseFromPb(pb *pipelinespb.GetPipelinePermissionLevelsResponsePb) (*GetPipelinePermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionLevelsResponse{}

	var permissionLevelsField []PipelinePermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := PipelinePermissionsDescriptionFromPb(&itemPb)
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

type GetPipelinePermissionsRequest struct {
	// The pipeline for which to get or manage permissions.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func GetPipelinePermissionsRequestToPb(st *GetPipelinePermissionsRequest) (*pipelinespb.GetPipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetPipelinePermissionsRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

func GetPipelinePermissionsRequestFromPb(pb *pipelinespb.GetPipelinePermissionsRequestPb) (*GetPipelinePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionsRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

type GetPipelineRequest struct {

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func GetPipelineRequestToPb(st *GetPipelineRequest) (*pipelinespb.GetPipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetPipelineRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

func GetPipelineRequestFromPb(pb *pipelinespb.GetPipelineRequestPb) (*GetPipelineRequest, error) {
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
	Cause string ``
	// The ID of the cluster that the pipeline is running on.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The username of the pipeline creator.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// Serverless budget policy ID of this pipeline.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string ``
	// The health of a pipeline.
	// Wire name: 'health'
	Health GetPipelineResponseHealth ``
	// The last time the pipeline settings were modified or created.
	// Wire name: 'last_modified'
	LastModified int64 ``
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	// Wire name: 'latest_updates'
	LatestUpdates []UpdateStateInfo ``
	// A human friendly identifier for the pipeline, taken from the `spec`.
	// Wire name: 'name'
	Name string ``
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// The user or service principal that the pipeline runs as, if specified in
	// the request. This field indicates the explicit configuration of `run_as`
	// for the pipeline. To find the value in all cases, explicit or implicit,
	// use `run_as_user_name`.
	// Wire name: 'run_as'
	RunAs *RunAs ``
	// Username of the user that the pipeline will run on behalf of.
	// Wire name: 'run_as_user_name'
	RunAsUserName string ``
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	// Wire name: 'spec'
	Spec *PipelineSpec ``
	// The pipeline state.
	// Wire name: 'state'
	State           PipelineState ``
	ForceSendFields []string      `tf:"-"`
}

func (s *GetPipelineResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPipelineResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetPipelineResponseToPb(st *GetPipelineResponse) (*pipelinespb.GetPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetPipelineResponsePb{}
	pb.Cause = st.Cause
	pb.ClusterId = st.ClusterId
	pb.CreatorUserName = st.CreatorUserName
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	healthPb, err := GetPipelineResponseHealthToPb(&st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = *healthPb
	}
	pb.LastModified = st.LastModified

	var latestUpdatesPb []pipelinespb.UpdateStateInfoPb
	for _, item := range st.LatestUpdates {
		itemPb, err := UpdateStateInfoToPb(&item)
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
	runAsPb, err := RunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}
	pb.RunAsUserName = st.RunAsUserName
	specPb, err := PipelineSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}
	statePb, err := PipelineStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetPipelineResponseFromPb(pb *pipelinespb.GetPipelineResponsePb) (*GetPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelineResponse{}
	st.Cause = pb.Cause
	st.ClusterId = pb.ClusterId
	st.CreatorUserName = pb.CreatorUserName
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	healthField, err := GetPipelineResponseHealthFromPb(&pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = *healthField
	}
	st.LastModified = pb.LastModified

	var latestUpdatesField []UpdateStateInfo
	for _, itemPb := range pb.LatestUpdates {
		item, err := UpdateStateInfoFromPb(&itemPb)
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
	runAsField, err := RunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	st.RunAsUserName = pb.RunAsUserName
	specField, err := PipelineSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	stateField, err := PipelineStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func GetPipelineResponseHealthToPb(st *GetPipelineResponseHealth) (*pipelinespb.GetPipelineResponseHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.GetPipelineResponseHealthPb(*st)
	return &pb, nil
}

func GetPipelineResponseHealthFromPb(pb *pipelinespb.GetPipelineResponseHealthPb) (*GetPipelineResponseHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetPipelineResponseHealth(*pb)
	return &st, nil
}

type GetUpdateRequest struct {
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
	// The ID of the update.
	// Wire name: 'update_id'
	UpdateId string `tf:"-"`
}

func GetUpdateRequestToPb(st *GetUpdateRequest) (*pipelinespb.GetUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetUpdateRequestPb{}
	pb.PipelineId = st.PipelineId
	pb.UpdateId = st.UpdateId

	return pb, nil
}

func GetUpdateRequestFromPb(pb *pipelinespb.GetUpdateRequestPb) (*GetUpdateRequest, error) {
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
	Update *UpdateInfo ``
}

func GetUpdateResponseToPb(st *GetUpdateResponse) (*pipelinespb.GetUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.GetUpdateResponsePb{}
	updatePb, err := UpdateInfoToPb(st.Update)
	if err != nil {
		return nil, err
	}
	if updatePb != nil {
		pb.Update = updatePb
	}

	return pb, nil
}

func GetUpdateResponseFromPb(pb *pipelinespb.GetUpdateResponsePb) (*GetUpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetUpdateResponse{}
	updateField, err := UpdateInfoFromPb(pb.Update)
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
	Report *ReportSpec ``
	// Select all tables from a specific source schema.
	// Wire name: 'schema'
	Schema *SchemaSpec ``
	// Select a specific source table.
	// Wire name: 'table'
	Table *TableSpec ``
}

func IngestionConfigToPb(st *IngestionConfig) (*pipelinespb.IngestionConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.IngestionConfigPb{}
	reportPb, err := ReportSpecToPb(st.Report)
	if err != nil {
		return nil, err
	}
	if reportPb != nil {
		pb.Report = reportPb
	}
	schemaPb, err := SchemaSpecToPb(st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = schemaPb
	}
	tablePb, err := TableSpecToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	return pb, nil
}

func IngestionConfigFromPb(pb *pipelinespb.IngestionConfigPb) (*IngestionConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionConfig{}
	reportField, err := ReportSpecFromPb(pb.Report)
	if err != nil {
		return nil, err
	}
	if reportField != nil {
		st.Report = reportField
	}
	schemaField, err := SchemaSpecFromPb(pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = schemaField
	}
	tableField, err := TableSpecFromPb(pb.Table)
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
	ConnectionId string ``
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	// Wire name: 'connection_name'
	ConnectionName string ``
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	// Wire name: 'gateway_storage_catalog'
	GatewayStorageCatalog string ``
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	// Wire name: 'gateway_storage_name'
	GatewayStorageName string ``
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	// Wire name: 'gateway_storage_schema'
	GatewayStorageSchema string   ``
	ForceSendFields      []string `tf:"-"`
}

func (s *IngestionGatewayPipelineDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionGatewayPipelineDefinition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func IngestionGatewayPipelineDefinitionToPb(st *IngestionGatewayPipelineDefinition) (*pipelinespb.IngestionGatewayPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.IngestionGatewayPipelineDefinitionPb{}
	pb.ConnectionId = st.ConnectionId
	pb.ConnectionName = st.ConnectionName
	pb.GatewayStorageCatalog = st.GatewayStorageCatalog
	pb.GatewayStorageName = st.GatewayStorageName
	pb.GatewayStorageSchema = st.GatewayStorageSchema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func IngestionGatewayPipelineDefinitionFromPb(pb *pipelinespb.IngestionGatewayPipelineDefinitionPb) (*IngestionGatewayPipelineDefinition, error) {
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

type IngestionPipelineDefinition struct {
	// Immutable. The Unity Catalog connection that this ingestion pipeline uses
	// to communicate with the source. This is used with connectors for
	// applications like Salesforce, Workday, and so on.
	// Wire name: 'connection_name'
	ConnectionName string ``
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	// Wire name: 'ingestion_gateway_id'
	IngestionGatewayId string ``
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	// Wire name: 'objects'
	Objects []IngestionConfig ``
	// The type of the foreign source. The source type will be inferred from the
	// source connection or ingestion gateway. This field is output only and
	// will be ignored if provided.
	// Wire name: 'source_type'
	SourceType IngestionSourceType ``
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig ``
	ForceSendFields    []string             `tf:"-"`
}

func (s *IngestionPipelineDefinition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func IngestionPipelineDefinitionToPb(st *IngestionPipelineDefinition) (*pipelinespb.IngestionPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.IngestionPipelineDefinitionPb{}
	pb.ConnectionName = st.ConnectionName
	pb.IngestionGatewayId = st.IngestionGatewayId

	var objectsPb []pipelinespb.IngestionConfigPb
	for _, item := range st.Objects {
		itemPb, err := IngestionConfigToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			objectsPb = append(objectsPb, *itemPb)
		}
	}
	pb.Objects = objectsPb
	sourceTypePb, err := IngestionSourceTypeToPb(&st.SourceType)
	if err != nil {
		return nil, err
	}
	if sourceTypePb != nil {
		pb.SourceType = *sourceTypePb
	}
	tableConfigurationPb, err := TableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func IngestionPipelineDefinitionFromPb(pb *pipelinespb.IngestionPipelineDefinitionPb) (*IngestionPipelineDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionPipelineDefinition{}
	st.ConnectionName = pb.ConnectionName
	st.IngestionGatewayId = pb.IngestionGatewayId

	var objectsField []IngestionConfig
	for _, itemPb := range pb.Objects {
		item, err := IngestionConfigFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			objectsField = append(objectsField, *item)
		}
	}
	st.Objects = objectsField
	sourceTypeField, err := IngestionSourceTypeFromPb(&pb.SourceType)
	if err != nil {
		return nil, err
	}
	if sourceTypeField != nil {
		st.SourceType = *sourceTypeField
	}
	tableConfigurationField, err := TableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Configurations that are only applicable for query-based ingestion connectors.
type IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig struct {
	// The names of the monotonically increasing columns in the source table
	// that are used to enable the table to be read and ingested incrementally
	// through structured streaming. The columns are allowed to have repeated
	// values but have to be non-decreasing. If the source data is merged into
	// the destination (e.g., using SCD Type 1 or Type 2), these columns will
	// implicitly define the `sequence_by` behavior. You can still explicitly
	// set `sequence_by` to override this default.
	// Wire name: 'cursor_columns'
	CursorColumns []string ``
	// Specifies a SQL WHERE condition that specifies that the source row has
	// been deleted. This is sometimes referred to as "soft-deletes". For
	// example: "Operation = 'DELETE'" or "is_deleted = true". This field is
	// orthogonal to `hard_deletion_sync_interval_in_seconds`, one for
	// soft-deletes and the other for hard-deletes. See also the
	// hard_deletion_sync_min_interval_in_seconds field for handling of "hard
	// deletes" where the source rows are physically removed from the table.
	// Wire name: 'deletion_condition'
	DeletionCondition string ``
	// Specifies the minimum interval (in seconds) between snapshots on primary
	// keys for detecting and synchronizing hard deletions—i.e., rows that
	// have been physically removed from the source table. This interval acts as
	// a lower bound. If ingestion runs less frequently than this value, hard
	// deletion synchronization will align with the actual ingestion frequency
	// instead of happening more often. If not set, hard deletion
	// synchronization via snapshots is disabled. This field is mutable and can
	// be updated without triggering a full snapshot.
	// Wire name: 'hard_deletion_sync_min_interval_in_seconds'
	HardDeletionSyncMinIntervalInSeconds int64    ``
	ForceSendFields                      []string `tf:"-"`
}

func (s *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigToPb(st *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig) (*pipelinespb.IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb{}
	pb.CursorColumns = st.CursorColumns
	pb.DeletionCondition = st.DeletionCondition
	pb.HardDeletionSyncMinIntervalInSeconds = st.HardDeletionSyncMinIntervalInSeconds

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigFromPb(pb *pipelinespb.IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb) (*IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig{}
	st.CursorColumns = pb.CursorColumns
	st.DeletionCondition = pb.DeletionCondition
	st.HardDeletionSyncMinIntervalInSeconds = pb.HardDeletionSyncMinIntervalInSeconds

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type IngestionSourceType string

const IngestionSourceTypeBigquery IngestionSourceType = `BIGQUERY`

const IngestionSourceTypeConfluence IngestionSourceType = `CONFLUENCE`

const IngestionSourceTypeDynamics365 IngestionSourceType = `DYNAMICS365`

const IngestionSourceTypeGa4RawData IngestionSourceType = `GA4_RAW_DATA`

const IngestionSourceTypeManagedPostgresql IngestionSourceType = `MANAGED_POSTGRESQL`

const IngestionSourceTypeMetaMarketing IngestionSourceType = `META_MARKETING`

const IngestionSourceTypeMysql IngestionSourceType = `MYSQL`

const IngestionSourceTypeNetsuite IngestionSourceType = `NETSUITE`

const IngestionSourceTypeOracle IngestionSourceType = `ORACLE`

const IngestionSourceTypePostgresql IngestionSourceType = `POSTGRESQL`

const IngestionSourceTypeRedshift IngestionSourceType = `REDSHIFT`

const IngestionSourceTypeSalesforce IngestionSourceType = `SALESFORCE`

const IngestionSourceTypeServicenow IngestionSourceType = `SERVICENOW`

const IngestionSourceTypeSharepoint IngestionSourceType = `SHAREPOINT`

const IngestionSourceTypeSqldw IngestionSourceType = `SQLDW`

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
	case `BIGQUERY`, `CONFLUENCE`, `DYNAMICS365`, `GA4_RAW_DATA`, `MANAGED_POSTGRESQL`, `META_MARKETING`, `MYSQL`, `NETSUITE`, `ORACLE`, `POSTGRESQL`, `REDSHIFT`, `SALESFORCE`, `SERVICENOW`, `SHAREPOINT`, `SQLDW`, `SQLSERVER`, `TERADATA`, `WORKDAY_RAAS`:
		*f = IngestionSourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BIGQUERY", "CONFLUENCE", "DYNAMICS365", "GA4_RAW_DATA", "MANAGED_POSTGRESQL", "META_MARKETING", "MYSQL", "NETSUITE", "ORACLE", "POSTGRESQL", "REDSHIFT", "SALESFORCE", "SERVICENOW", "SHAREPOINT", "SQLDW", "SQLSERVER", "TERADATA", "WORKDAY_RAAS"`, v)
	}
}

// Values returns all possible values for IngestionSourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *IngestionSourceType) Values() []IngestionSourceType {
	return []IngestionSourceType{
		IngestionSourceTypeBigquery,
		IngestionSourceTypeConfluence,
		IngestionSourceTypeDynamics365,
		IngestionSourceTypeGa4RawData,
		IngestionSourceTypeManagedPostgresql,
		IngestionSourceTypeMetaMarketing,
		IngestionSourceTypeMysql,
		IngestionSourceTypeNetsuite,
		IngestionSourceTypeOracle,
		IngestionSourceTypePostgresql,
		IngestionSourceTypeRedshift,
		IngestionSourceTypeSalesforce,
		IngestionSourceTypeServicenow,
		IngestionSourceTypeSharepoint,
		IngestionSourceTypeSqldw,
		IngestionSourceTypeSqlserver,
		IngestionSourceTypeTeradata,
		IngestionSourceTypeWorkdayRaas,
	}
}

// Type always returns IngestionSourceType to satisfy [pflag.Value] interface
func (f *IngestionSourceType) Type() string {
	return "IngestionSourceType"
}

func IngestionSourceTypeToPb(st *IngestionSourceType) (*pipelinespb.IngestionSourceTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.IngestionSourceTypePb(*st)
	return &pb, nil
}

func IngestionSourceTypeFromPb(pb *pipelinespb.IngestionSourceTypePb) (*IngestionSourceType, error) {
	if pb == nil {
		return nil, nil
	}
	st := IngestionSourceType(*pb)
	return &st, nil
}

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
	// The pipeline to return events for.
	// Wire name: 'pipeline_id'
	PipelineId      string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListPipelineEventsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelineEventsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListPipelineEventsRequestToPb(st *ListPipelineEventsRequest) (*pipelinespb.ListPipelineEventsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ListPipelineEventsRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken
	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListPipelineEventsRequestFromPb(pb *pipelinespb.ListPipelineEventsRequestPb) (*ListPipelineEventsRequest, error) {
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

type ListPipelineEventsResponse struct {
	// The list of events matching the request criteria.
	// Wire name: 'events'
	Events []PipelineEvent ``
	// If present, a token to fetch the next page of events.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// If present, a token to fetch the previous page of events.
	// Wire name: 'prev_page_token'
	PrevPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListPipelineEventsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelineEventsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListPipelineEventsResponseToPb(st *ListPipelineEventsResponse) (*pipelinespb.ListPipelineEventsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ListPipelineEventsResponsePb{}

	var eventsPb []pipelinespb.PipelineEventPb
	for _, item := range st.Events {
		itemPb, err := PipelineEventToPb(&item)
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

func ListPipelineEventsResponseFromPb(pb *pipelinespb.ListPipelineEventsResponsePb) (*ListPipelineEventsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelineEventsResponse{}

	var eventsField []PipelineEvent
	for _, itemPb := range pb.Events {
		item, err := PipelineEventFromPb(&itemPb)
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListPipelinesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelinesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListPipelinesRequestToPb(st *ListPipelinesRequest) (*pipelinespb.ListPipelinesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ListPipelinesRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListPipelinesRequestFromPb(pb *pipelinespb.ListPipelinesRequestPb) (*ListPipelinesRequest, error) {
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

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// The list of events matching the request criteria.
	// Wire name: 'statuses'
	Statuses        []PipelineStateInfo ``
	ForceSendFields []string            `tf:"-"`
}

func (s *ListPipelinesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPipelinesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListPipelinesResponseToPb(st *ListPipelinesResponse) (*pipelinespb.ListPipelinesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ListPipelinesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var statusesPb []pipelinespb.PipelineStateInfoPb
	for _, item := range st.Statuses {
		itemPb, err := PipelineStateInfoToPb(&item)
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

func ListPipelinesResponseFromPb(pb *pipelinespb.ListPipelinesResponsePb) (*ListPipelinesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelinesResponse{}
	st.NextPageToken = pb.NextPageToken

	var statusesField []PipelineStateInfo
	for _, itemPb := range pb.Statuses {
		item, err := PipelineStateInfoFromPb(&itemPb)
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
	UntilUpdateId   string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListUpdatesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUpdatesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListUpdatesRequestToPb(st *ListUpdatesRequest) (*pipelinespb.ListUpdatesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ListUpdatesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.PipelineId = st.PipelineId
	pb.UntilUpdateId = st.UntilUpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListUpdatesRequestFromPb(pb *pipelinespb.ListUpdatesRequestPb) (*ListUpdatesRequest, error) {
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

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	// Wire name: 'prev_page_token'
	PrevPageToken string ``

	// Wire name: 'updates'
	Updates         []UpdateInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (s *ListUpdatesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListUpdatesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListUpdatesResponseToPb(st *ListUpdatesResponse) (*pipelinespb.ListUpdatesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ListUpdatesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	var updatesPb []pipelinespb.UpdateInfoPb
	for _, item := range st.Updates {
		itemPb, err := UpdateInfoToPb(&item)
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

func ListUpdatesResponseFromPb(pb *pipelinespb.ListUpdatesResponsePb) (*ListUpdatesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUpdatesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	var updatesField []UpdateInfo
	for _, itemPb := range pb.Updates {
		item, err := UpdateInfoFromPb(&itemPb)
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

type ManualTrigger struct {
}

func ManualTriggerToPb(st *ManualTrigger) (*pipelinespb.ManualTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ManualTriggerPb{}

	return pb, nil
}

func ManualTriggerFromPb(pb *pipelinespb.ManualTriggerPb) (*ManualTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ManualTrigger{}

	return st, nil
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

func MaturityLevelToPb(st *MaturityLevel) (*pipelinespb.MaturityLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.MaturityLevelPb(*st)
	return &pb, nil
}

func MaturityLevelFromPb(pb *pipelinespb.MaturityLevelPb) (*MaturityLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := MaturityLevel(*pb)
	return &st, nil
}

type NotebookLibrary struct {
	// The absolute path of the source code.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *NotebookLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NotebookLibraryToPb(st *NotebookLibrary) (*pipelinespb.NotebookLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.NotebookLibraryPb{}
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NotebookLibraryFromPb(pb *pipelinespb.NotebookLibraryPb) (*NotebookLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookLibrary{}
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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
	Alerts []string ``
	// A list of email addresses notified when a configured alert is triggered.
	// Wire name: 'email_recipients'
	EmailRecipients []string ``
}

func NotificationsToPb(st *Notifications) (*pipelinespb.NotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.NotificationsPb{}
	pb.Alerts = st.Alerts
	pb.EmailRecipients = st.EmailRecipients

	return pb, nil
}

func NotificationsFromPb(pb *pipelinespb.NotificationsPb) (*Notifications, error) {
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
	BatchId int64 ``
	// The cloud provider, e.g., AWS or Azure.
	// Wire name: 'cloud'
	Cloud string ``
	// The id of the cluster where an execution happens. Unique within a region.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The name of a dataset. Unique within a pipeline.
	// Wire name: 'dataset_name'
	DatasetName string ``
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	// Wire name: 'flow_id'
	FlowId string ``
	// The name of the flow. Not unique.
	// Wire name: 'flow_name'
	FlowName string ``
	// The optional host name where the event was triggered
	// Wire name: 'host'
	Host string ``
	// The id of a maintenance run. Globally unique.
	// Wire name: 'maintenance_id'
	MaintenanceId string ``
	// Materialization name.
	// Wire name: 'materialization_name'
	MaterializationName string ``
	// The org id of the user. Unique within a cloud.
	// Wire name: 'org_id'
	OrgId int64 ``
	// The id of the pipeline. Globally unique.
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// The name of the pipeline. Not unique.
	// Wire name: 'pipeline_name'
	PipelineName string ``
	// The cloud region.
	// Wire name: 'region'
	Region string ``
	// The id of the request that caused an update.
	// Wire name: 'request_id'
	RequestId string ``
	// The id of a (delta) table. Globally unique.
	// Wire name: 'table_id'
	TableId string ``
	// The Unity Catalog id of the MV or ST being updated.
	// Wire name: 'uc_resource_id'
	UcResourceId string ``
	// The id of an execution. Globally unique.
	// Wire name: 'update_id'
	UpdateId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Origin) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Origin) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OriginToPb(st *Origin) (*pipelinespb.OriginPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.OriginPb{}
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

func OriginFromPb(pb *pipelinespb.OriginPb) (*Origin, error) {
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

type PathPattern struct {
	// The source code to include for pipelines
	// Wire name: 'include'
	Include         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PathPattern) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PathPattern) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PathPatternToPb(st *PathPattern) (*pipelinespb.PathPatternPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PathPatternPb{}
	pb.Include = st.Include

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PathPatternFromPb(pb *pipelinespb.PathPatternPb) (*PathPattern, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PathPattern{}
	st.Include = pb.Include

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PipelineAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelineAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineAccessControlRequestToPb(st *PipelineAccessControlRequest) (*pipelinespb.PipelineAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := PipelinePermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelineAccessControlRequestFromPb(pb *pipelinespb.PipelineAccessControlRequestPb) (*PipelineAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := PipelinePermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PipelineAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []PipelinePermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelineAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineAccessControlResponseToPb(st *PipelineAccessControlResponse) (*pipelinespb.PipelineAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineAccessControlResponsePb{}

	var allPermissionsPb []pipelinespb.PipelinePermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := PipelinePermissionToPb(&item)
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

func PipelineAccessControlResponseFromPb(pb *pipelinespb.PipelineAccessControlResponsePb) (*PipelineAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineAccessControlResponse{}

	var allPermissionsField []PipelinePermission
	for _, itemPb := range pb.AllPermissions {
		item, err := PipelinePermissionFromPb(&itemPb)
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

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool ``
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *PipelineClusterAutoscale ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *compute.AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *compute.AzureAttributes ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *compute.ClusterLogConf ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Whether to enable local disk encryption for the cluster.
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *compute.GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []compute.InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	// Wire name: 'label'
	Label string ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
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
	NumWorkers int ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
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
	SparkEnvVars map[string]string ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys   []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelineCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineClusterToPb(st *PipelineCluster) (*pipelinespb.PipelineClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	autoscalePb, err := PipelineClusterAutoscaleToPb(st.Autoscale)
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

	var initScriptsPb []computepb.InitScriptInfoPb
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

func PipelineClusterFromPb(pb *pipelinespb.PipelineClusterPb) (*PipelineCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	autoscaleField, err := PipelineClusterAutoscaleFromPb(pb.Autoscale)
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

type PipelineClusterAutoscale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. `max_workers` must be strictly greater than `min_workers`.
	// Wire name: 'max_workers'
	MaxWorkers int ``
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	// Wire name: 'min_workers'
	MinWorkers int ``
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	// Wire name: 'mode'
	Mode PipelineClusterAutoscaleMode ``
}

func PipelineClusterAutoscaleToPb(st *PipelineClusterAutoscale) (*pipelinespb.PipelineClusterAutoscalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineClusterAutoscalePb{}
	pb.MaxWorkers = st.MaxWorkers
	pb.MinWorkers = st.MinWorkers
	modePb, err := PipelineClusterAutoscaleModeToPb(&st.Mode)
	if err != nil {
		return nil, err
	}
	if modePb != nil {
		pb.Mode = *modePb
	}

	return pb, nil
}

func PipelineClusterAutoscaleFromPb(pb *pipelinespb.PipelineClusterAutoscalePb) (*PipelineClusterAutoscale, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineClusterAutoscale{}
	st.MaxWorkers = pb.MaxWorkers
	st.MinWorkers = pb.MinWorkers
	modeField, err := PipelineClusterAutoscaleModeFromPb(&pb.Mode)
	if err != nil {
		return nil, err
	}
	if modeField != nil {
		st.Mode = *modeField
	}

	return st, nil
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

func PipelineClusterAutoscaleModeToPb(st *PipelineClusterAutoscaleMode) (*pipelinespb.PipelineClusterAutoscaleModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.PipelineClusterAutoscaleModePb(*st)
	return &pb, nil
}

func PipelineClusterAutoscaleModeFromPb(pb *pipelinespb.PipelineClusterAutoscaleModePb) (*PipelineClusterAutoscaleMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineClusterAutoscaleMode(*pb)
	return &st, nil
}

type PipelineDeployment struct {
	// The deployment method that manages the pipeline.
	// Wire name: 'kind'
	Kind DeploymentKind ``
	// The path to the file containing metadata about the deployment.
	// Wire name: 'metadata_file_path'
	MetadataFilePath string   ``
	ForceSendFields  []string `tf:"-"`
}

func (s *PipelineDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineDeploymentToPb(st *PipelineDeployment) (*pipelinespb.PipelineDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineDeploymentPb{}
	kindPb, err := DeploymentKindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.MetadataFilePath = st.MetadataFilePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelineDeploymentFromPb(pb *pipelinespb.PipelineDeploymentPb) (*PipelineDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineDeployment{}
	kindField, err := DeploymentKindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.MetadataFilePath = pb.MetadataFilePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PipelineEvent struct {
	// Information about an error captured by the event.
	// Wire name: 'error'
	Error *ErrorDetail ``
	// The event type. Should always correspond to the details
	// Wire name: 'event_type'
	EventType string ``
	// A time-based, globally unique id.
	// Wire name: 'id'
	Id string ``
	// The severity level of the event.
	// Wire name: 'level'
	Level EventLevel ``
	// Maturity level for event_type.
	// Wire name: 'maturity_level'
	MaturityLevel MaturityLevel ``
	// The display message associated with the event.
	// Wire name: 'message'
	Message string ``
	// Describes where the event originates from.
	// Wire name: 'origin'
	Origin *Origin ``
	// A sequencing object to identify and order events.
	// Wire name: 'sequence'
	Sequence *Sequencing ``
	// The time of the event.
	// Wire name: 'timestamp'
	Timestamp       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelineEvent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineEvent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineEventToPb(st *PipelineEvent) (*pipelinespb.PipelineEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineEventPb{}
	errorPb, err := ErrorDetailToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}
	pb.EventType = st.EventType
	pb.Id = st.Id
	levelPb, err := EventLevelToPb(&st.Level)
	if err != nil {
		return nil, err
	}
	if levelPb != nil {
		pb.Level = *levelPb
	}
	maturityLevelPb, err := MaturityLevelToPb(&st.MaturityLevel)
	if err != nil {
		return nil, err
	}
	if maturityLevelPb != nil {
		pb.MaturityLevel = *maturityLevelPb
	}
	pb.Message = st.Message
	originPb, err := OriginToPb(st.Origin)
	if err != nil {
		return nil, err
	}
	if originPb != nil {
		pb.Origin = originPb
	}
	sequencePb, err := SequencingToPb(st.Sequence)
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

func PipelineEventFromPb(pb *pipelinespb.PipelineEventPb) (*PipelineEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineEvent{}
	errorField, err := ErrorDetailFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	st.EventType = pb.EventType
	st.Id = pb.Id
	levelField, err := EventLevelFromPb(&pb.Level)
	if err != nil {
		return nil, err
	}
	if levelField != nil {
		st.Level = *levelField
	}
	maturityLevelField, err := MaturityLevelFromPb(&pb.MaturityLevel)
	if err != nil {
		return nil, err
	}
	if maturityLevelField != nil {
		st.MaturityLevel = *maturityLevelField
	}
	st.Message = pb.Message
	originField, err := OriginFromPb(pb.Origin)
	if err != nil {
		return nil, err
	}
	if originField != nil {
		st.Origin = originField
	}
	sequenceField, err := SequencingFromPb(pb.Sequence)
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

type PipelineLibrary struct {
	// The path to a file that defines a pipeline and is stored in the
	// Databricks Repos.
	// Wire name: 'file'
	File *FileLibrary ``
	// The unified field to include source codes. Each entry can be a notebook
	// path, a file path, or a folder path that ends `/**`. This field cannot be
	// used together with `notebook` or `file`.
	// Wire name: 'glob'
	Glob *PathPattern ``
	// URI of the jar to be installed. Currently only DBFS is supported.
	// Wire name: 'jar'
	Jar string ``
	// Specification of a maven library to be installed.
	// Wire name: 'maven'
	Maven *compute.MavenLibrary ``
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	// Wire name: 'notebook'
	Notebook *NotebookLibrary ``
	// URI of the whl to be installed.
	// Wire name: 'whl'
	Whl             string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelineLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineLibraryToPb(st *PipelineLibrary) (*pipelinespb.PipelineLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineLibraryPb{}
	filePb, err := FileLibraryToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}
	globPb, err := PathPatternToPb(st.Glob)
	if err != nil {
		return nil, err
	}
	if globPb != nil {
		pb.Glob = globPb
	}
	pb.Jar = st.Jar
	mavenPb, err := compute.MavenLibraryToPb(st.Maven)
	if err != nil {
		return nil, err
	}
	if mavenPb != nil {
		pb.Maven = mavenPb
	}
	notebookPb, err := NotebookLibraryToPb(st.Notebook)
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

func PipelineLibraryFromPb(pb *pipelinespb.PipelineLibraryPb) (*PipelineLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineLibrary{}
	fileField, err := FileLibraryFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	globField, err := PathPatternFromPb(pb.Glob)
	if err != nil {
		return nil, err
	}
	if globField != nil {
		st.Glob = globField
	}
	st.Jar = pb.Jar
	mavenField, err := compute.MavenLibraryFromPb(pb.Maven)
	if err != nil {
		return nil, err
	}
	if mavenField != nil {
		st.Maven = mavenField
	}
	notebookField, err := NotebookLibraryFromPb(pb.Notebook)
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

type PipelinePermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel ``
	ForceSendFields []string                `tf:"-"`
}

func (s *PipelinePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelinePermissionToPb(st *PipelinePermission) (*pipelinespb.PipelinePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelinePermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := PipelinePermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelinePermissionFromPb(pb *pipelinespb.PipelinePermissionPb) (*PipelinePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := PipelinePermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func PipelinePermissionLevelToPb(st *PipelinePermissionLevel) (*pipelinespb.PipelinePermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.PipelinePermissionLevelPb(*st)
	return &pb, nil
}

func PipelinePermissionLevelFromPb(pb *pipelinespb.PipelinePermissionLevelPb) (*PipelinePermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelinePermissionLevel(*pb)
	return &st, nil
}

type PipelinePermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []PipelineAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PipelinePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelinePermissionsToPb(st *PipelinePermissions) (*pipelinespb.PipelinePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelinePermissionsPb{}

	var accessControlListPb []pipelinespb.PipelineAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := PipelineAccessControlResponseToPb(&item)
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

func PipelinePermissionsFromPb(pb *pipelinespb.PipelinePermissionsPb) (*PipelinePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissions{}

	var accessControlListField []PipelineAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := PipelineAccessControlResponseFromPb(&itemPb)
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

type PipelinePermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel PipelinePermissionLevel ``
	ForceSendFields []string                `tf:"-"`
}

func (s *PipelinePermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelinePermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelinePermissionsDescriptionToPb(st *PipelinePermissionsDescription) (*pipelinespb.PipelinePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelinePermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := PipelinePermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelinePermissionsDescriptionFromPb(pb *pipelinespb.PipelinePermissionsDescriptionPb) (*PipelinePermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := PipelinePermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PipelinePermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []PipelineAccessControlRequest ``
	// The pipeline for which to get or manage permissions.
	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func PipelinePermissionsRequestToPb(st *PipelinePermissionsRequest) (*pipelinespb.PipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelinePermissionsRequestPb{}

	var accessControlListPb []pipelinespb.PipelineAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := PipelineAccessControlRequestToPb(&item)
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

func PipelinePermissionsRequestFromPb(pb *pipelinespb.PipelinePermissionsRequestPb) (*PipelinePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissionsRequest{}

	var accessControlListField []PipelineAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := PipelineAccessControlRequestFromPb(&itemPb)
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
	BudgetPolicyId string ``
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	// Wire name: 'catalog'
	Catalog string ``
	// DLT Release Channel that specifies which version to use.
	// Wire name: 'channel'
	Channel string ``
	// Cluster settings for this pipeline deployment.
	// Wire name: 'clusters'
	Clusters []PipelineCluster ``
	// String-String configuration for this pipeline execution.
	// Wire name: 'configuration'
	Configuration map[string]string ``
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	// Wire name: 'continuous'
	Continuous bool ``
	// Deployment type of this pipeline.
	// Wire name: 'deployment'
	Deployment *PipelineDeployment ``
	// Whether the pipeline is in Development mode. Defaults to false.
	// Wire name: 'development'
	Development bool ``
	// Pipeline product edition.
	// Wire name: 'edition'
	Edition string ``
	// Environment specification for this pipeline used to install dependencies.
	// Wire name: 'environment'
	Environment *PipelinesEnvironment ``
	// Event log configuration for this pipeline
	// Wire name: 'event_log'
	EventLog *EventLogSpec ``
	// Filters on which Pipeline packages to include in the deployed graph.
	// Wire name: 'filters'
	Filters *Filters ``
	// The definition of a gateway pipeline to support change data capture.
	// Wire name: 'gateway_definition'
	GatewayDefinition *IngestionGatewayPipelineDefinition ``
	// Unique identifier for this pipeline.
	// Wire name: 'id'
	Id string ``
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	// Wire name: 'ingestion_definition'
	IngestionDefinition *IngestionPipelineDefinition ``
	// Libraries or code needed by this deployment.
	// Wire name: 'libraries'
	Libraries []PipelineLibrary ``
	// Friendly identifier for this pipeline.
	// Wire name: 'name'
	Name string ``
	// List of notification settings for this pipeline.
	// Wire name: 'notifications'
	Notifications []Notifications ``
	// Whether Photon is enabled for this pipeline.
	// Wire name: 'photon'
	Photon bool ``
	// Restart window of this pipeline.
	// Wire name: 'restart_window'
	RestartWindow *RestartWindow ``
	// Root path for this pipeline. This is used as the root directory when
	// editing the pipeline in the Databricks user interface and it is added to
	// sys.path when executing Python sources during pipeline execution.
	// Wire name: 'root_path'
	RootPath string ``
	// The default schema (database) where tables are read from or published to.
	// Wire name: 'schema'
	Schema string ``
	// Whether serverless compute is enabled for this pipeline.
	// Wire name: 'serverless'
	Serverless bool ``
	// DBFS root directory for storing checkpoints and tables.
	// Wire name: 'storage'
	Storage string ``
	// A map of tags associated with the pipeline. These are forwarded to the
	// cluster as cluster tags, and are therefore subject to the same
	// limitations. A maximum of 25 tags can be added to the pipeline.
	// Wire name: 'tags'
	Tags map[string]string ``
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	// Wire name: 'target'
	Target string ``
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	// Wire name: 'trigger'
	Trigger         *PipelineTrigger ``
	ForceSendFields []string         `tf:"-"`
}

func (s *PipelineSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineSpecToPb(st *PipelineSpec) (*pipelinespb.PipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineSpecPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Catalog = st.Catalog
	pb.Channel = st.Channel

	var clustersPb []pipelinespb.PipelineClusterPb
	for _, item := range st.Clusters {
		itemPb, err := PipelineClusterToPb(&item)
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
	deploymentPb, err := PipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}
	pb.Development = st.Development
	pb.Edition = st.Edition
	environmentPb, err := PipelinesEnvironmentToPb(st.Environment)
	if err != nil {
		return nil, err
	}
	if environmentPb != nil {
		pb.Environment = environmentPb
	}
	eventLogPb, err := EventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}
	filtersPb, err := FiltersToPb(st.Filters)
	if err != nil {
		return nil, err
	}
	if filtersPb != nil {
		pb.Filters = filtersPb
	}
	gatewayDefinitionPb, err := IngestionGatewayPipelineDefinitionToPb(st.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionPb != nil {
		pb.GatewayDefinition = gatewayDefinitionPb
	}
	pb.Id = st.Id
	ingestionDefinitionPb, err := IngestionPipelineDefinitionToPb(st.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionPb != nil {
		pb.IngestionDefinition = ingestionDefinitionPb
	}

	var librariesPb []pipelinespb.PipelineLibraryPb
	for _, item := range st.Libraries {
		itemPb, err := PipelineLibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.Name = st.Name

	var notificationsPb []pipelinespb.NotificationsPb
	for _, item := range st.Notifications {
		itemPb, err := NotificationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notificationsPb = append(notificationsPb, *itemPb)
		}
	}
	pb.Notifications = notificationsPb
	pb.Photon = st.Photon
	restartWindowPb, err := RestartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}
	pb.RootPath = st.RootPath
	pb.Schema = st.Schema
	pb.Serverless = st.Serverless
	pb.Storage = st.Storage
	pb.Tags = st.Tags
	pb.Target = st.Target
	triggerPb, err := PipelineTriggerToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelineSpecFromPb(pb *pipelinespb.PipelineSpecPb) (*PipelineSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineSpec{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel

	var clustersField []PipelineCluster
	for _, itemPb := range pb.Clusters {
		item, err := PipelineClusterFromPb(&itemPb)
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
	deploymentField, err := PipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Development = pb.Development
	st.Edition = pb.Edition
	environmentField, err := PipelinesEnvironmentFromPb(pb.Environment)
	if err != nil {
		return nil, err
	}
	if environmentField != nil {
		st.Environment = environmentField
	}
	eventLogField, err := EventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	filtersField, err := FiltersFromPb(pb.Filters)
	if err != nil {
		return nil, err
	}
	if filtersField != nil {
		st.Filters = filtersField
	}
	gatewayDefinitionField, err := IngestionGatewayPipelineDefinitionFromPb(pb.GatewayDefinition)
	if err != nil {
		return nil, err
	}
	if gatewayDefinitionField != nil {
		st.GatewayDefinition = gatewayDefinitionField
	}
	st.Id = pb.Id
	ingestionDefinitionField, err := IngestionPipelineDefinitionFromPb(pb.IngestionDefinition)
	if err != nil {
		return nil, err
	}
	if ingestionDefinitionField != nil {
		st.IngestionDefinition = ingestionDefinitionField
	}

	var librariesField []PipelineLibrary
	for _, itemPb := range pb.Libraries {
		item, err := PipelineLibraryFromPb(&itemPb)
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
		item, err := NotificationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notificationsField = append(notificationsField, *item)
		}
	}
	st.Notifications = notificationsField
	st.Photon = pb.Photon
	restartWindowField, err := RestartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	st.RootPath = pb.RootPath
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Tags = pb.Tags
	st.Target = pb.Target
	triggerField, err := PipelineTriggerFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func PipelineStateToPb(st *PipelineState) (*pipelinespb.PipelineStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.PipelineStatePb(*st)
	return &pb, nil
}

func PipelineStateFromPb(pb *pipelinespb.PipelineStatePb) (*PipelineState, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineState(*pb)
	return &st, nil
}

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The username of the pipeline creator.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// The health of a pipeline.
	// Wire name: 'health'
	Health PipelineStateInfoHealth ``
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	// Wire name: 'latest_updates'
	LatestUpdates []UpdateStateInfo ``
	// The user-friendly name of the pipeline.
	// Wire name: 'name'
	Name string ``
	// The unique identifier of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	// Wire name: 'run_as_user_name'
	RunAsUserName string ``

	// Wire name: 'state'
	State           PipelineState ``
	ForceSendFields []string      `tf:"-"`
}

func (s *PipelineStateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineStateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PipelineStateInfoToPb(st *PipelineStateInfo) (*pipelinespb.PipelineStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineStateInfoPb{}
	pb.ClusterId = st.ClusterId
	pb.CreatorUserName = st.CreatorUserName
	healthPb, err := PipelineStateInfoHealthToPb(&st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = *healthPb
	}

	var latestUpdatesPb []pipelinespb.UpdateStateInfoPb
	for _, item := range st.LatestUpdates {
		itemPb, err := UpdateStateInfoToPb(&item)
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
	statePb, err := PipelineStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PipelineStateInfoFromPb(pb *pipelinespb.PipelineStateInfoPb) (*PipelineStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineStateInfo{}
	st.ClusterId = pb.ClusterId
	st.CreatorUserName = pb.CreatorUserName
	healthField, err := PipelineStateInfoHealthFromPb(&pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = *healthField
	}

	var latestUpdatesField []UpdateStateInfo
	for _, itemPb := range pb.LatestUpdates {
		item, err := UpdateStateInfoFromPb(&itemPb)
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
	stateField, err := PipelineStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func PipelineStateInfoHealthToPb(st *PipelineStateInfoHealth) (*pipelinespb.PipelineStateInfoHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.PipelineStateInfoHealthPb(*st)
	return &pb, nil
}

func PipelineStateInfoHealthFromPb(pb *pipelinespb.PipelineStateInfoHealthPb) (*PipelineStateInfoHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := PipelineStateInfoHealth(*pb)
	return &st, nil
}

type PipelineTrigger struct {

	// Wire name: 'cron'
	Cron *CronTrigger ``

	// Wire name: 'manual'
	Manual *ManualTrigger ``
}

func PipelineTriggerToPb(st *PipelineTrigger) (*pipelinespb.PipelineTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelineTriggerPb{}
	cronPb, err := CronTriggerToPb(st.Cron)
	if err != nil {
		return nil, err
	}
	if cronPb != nil {
		pb.Cron = cronPb
	}
	manualPb, err := ManualTriggerToPb(st.Manual)
	if err != nil {
		return nil, err
	}
	if manualPb != nil {
		pb.Manual = manualPb
	}

	return pb, nil
}

func PipelineTriggerFromPb(pb *pipelinespb.PipelineTriggerPb) (*PipelineTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineTrigger{}
	cronField, err := CronTriggerFromPb(pb.Cron)
	if err != nil {
		return nil, err
	}
	if cronField != nil {
		st.Cron = cronField
	}
	manualField, err := ManualTriggerFromPb(pb.Manual)
	if err != nil {
		return nil, err
	}
	if manualField != nil {
		st.Manual = manualField
	}

	return st, nil
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
	Dependencies []string ``
}

func PipelinesEnvironmentToPb(st *PipelinesEnvironment) (*pipelinespb.PipelinesEnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.PipelinesEnvironmentPb{}
	pb.Dependencies = st.Dependencies

	return pb, nil
}

func PipelinesEnvironmentFromPb(pb *pipelinespb.PipelinesEnvironmentPb) (*PipelinesEnvironment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinesEnvironment{}
	st.Dependencies = pb.Dependencies

	return st, nil
}

type ReportSpec struct {
	// Required. Destination catalog to store table.
	// Wire name: 'destination_catalog'
	DestinationCatalog string ``
	// Required. Destination schema to store table.
	// Wire name: 'destination_schema'
	DestinationSchema string ``
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	// Wire name: 'destination_table'
	DestinationTable string ``
	// Required. Report URL in the source system.
	// Wire name: 'source_url'
	SourceUrl string ``
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig ``
	ForceSendFields    []string             `tf:"-"`
}

func (s *ReportSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReportSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ReportSpecToPb(st *ReportSpec) (*pipelinespb.ReportSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.ReportSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog
	pb.DestinationSchema = st.DestinationSchema
	pb.DestinationTable = st.DestinationTable
	pb.SourceUrl = st.SourceUrl
	tableConfigurationPb, err := TableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ReportSpecFromPb(pb *pipelinespb.ReportSpecPb) (*ReportSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReportSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.DestinationTable = pb.DestinationTable
	st.SourceUrl = pb.SourceUrl
	tableConfigurationField, err := TableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type RestartWindow struct {
	// Days of week in which the restart is allowed to happen (within a
	// five-hour window starting at start_hour). If not specified all days of
	// the week will be used.
	// Wire name: 'days_of_week'
	DaysOfWeek []DayOfWeek ``
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	// Wire name: 'start_hour'
	StartHour int ``
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	// Wire name: 'time_zone_id'
	TimeZoneId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RestartWindow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestartWindow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RestartWindowToPb(st *RestartWindow) (*pipelinespb.RestartWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.RestartWindowPb{}

	var daysOfWeekPb []pipelinespb.DayOfWeekPb
	for _, item := range st.DaysOfWeek {
		itemPb, err := DayOfWeekToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			daysOfWeekPb = append(daysOfWeekPb, *itemPb)
		}
	}
	pb.DaysOfWeek = daysOfWeekPb
	pb.StartHour = st.StartHour
	pb.TimeZoneId = st.TimeZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RestartWindowFromPb(pb *pipelinespb.RestartWindowPb) (*RestartWindow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartWindow{}

	var daysOfWeekField []DayOfWeek
	for _, itemPb := range pb.DaysOfWeek {
		item, err := DayOfWeekFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			daysOfWeekField = append(daysOfWeekField, *item)
		}
	}
	st.DaysOfWeek = daysOfWeekField
	st.StartHour = pb.StartHour
	st.TimeZoneId = pb.TimeZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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
	ServicePrincipalName string ``
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *RunAs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunAs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RunAsToPb(st *RunAs) (*pipelinespb.RunAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.RunAsPb{}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RunAsFromPb(pb *pipelinespb.RunAsPb) (*RunAs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunAs{}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SchemaSpec struct {
	// Required. Destination catalog to store tables.
	// Wire name: 'destination_catalog'
	DestinationCatalog string ``
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	// Wire name: 'destination_schema'
	DestinationSchema string ``
	// The source catalog name. Might be optional depending on the type of
	// source.
	// Wire name: 'source_catalog'
	SourceCatalog string ``
	// Required. Schema name in the source database.
	// Wire name: 'source_schema'
	SourceSchema string ``
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig ``
	ForceSendFields    []string             `tf:"-"`
}

func (s *SchemaSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SchemaSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SchemaSpecToPb(st *SchemaSpec) (*pipelinespb.SchemaSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.SchemaSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog
	pb.DestinationSchema = st.DestinationSchema
	pb.SourceCatalog = st.SourceCatalog
	pb.SourceSchema = st.SourceSchema
	tableConfigurationPb, err := TableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SchemaSpecFromPb(pb *pipelinespb.SchemaSpecPb) (*SchemaSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SchemaSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.SourceCatalog = pb.SourceCatalog
	st.SourceSchema = pb.SourceSchema
	tableConfigurationField, err := TableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type Sequencing struct {
	// A sequence number, unique and increasing per pipeline.
	// Wire name: 'control_plane_seq_no'
	ControlPlaneSeqNo int64 ``
	// the ID assigned by the data plane.
	// Wire name: 'data_plane_id'
	DataPlaneId     *DataPlaneId ``
	ForceSendFields []string     `tf:"-"`
}

func (s *Sequencing) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Sequencing) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SequencingToPb(st *Sequencing) (*pipelinespb.SequencingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.SequencingPb{}
	pb.ControlPlaneSeqNo = st.ControlPlaneSeqNo
	dataPlaneIdPb, err := DataPlaneIdToPb(st.DataPlaneId)
	if err != nil {
		return nil, err
	}
	if dataPlaneIdPb != nil {
		pb.DataPlaneId = dataPlaneIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SequencingFromPb(pb *pipelinespb.SequencingPb) (*Sequencing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Sequencing{}
	st.ControlPlaneSeqNo = pb.ControlPlaneSeqNo
	dataPlaneIdField, err := DataPlaneIdFromPb(pb.DataPlaneId)
	if err != nil {
		return nil, err
	}
	if dataPlaneIdField != nil {
		st.DataPlaneId = dataPlaneIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SerializedException struct {
	// Runtime class of the exception
	// Wire name: 'class_name'
	ClassName string ``
	// Exception message
	// Wire name: 'message'
	Message string ``
	// Stack trace consisting of a list of stack frames
	// Wire name: 'stack'
	Stack           []StackFrame ``
	ForceSendFields []string     `tf:"-"`
}

func (s *SerializedException) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SerializedException) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SerializedExceptionToPb(st *SerializedException) (*pipelinespb.SerializedExceptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.SerializedExceptionPb{}
	pb.ClassName = st.ClassName
	pb.Message = st.Message

	var stackPb []pipelinespb.StackFramePb
	for _, item := range st.Stack {
		itemPb, err := StackFrameToPb(&item)
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

func SerializedExceptionFromPb(pb *pipelinespb.SerializedExceptionPb) (*SerializedException, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SerializedException{}
	st.ClassName = pb.ClassName
	st.Message = pb.Message

	var stackField []StackFrame
	for _, itemPb := range pb.Stack {
		item, err := StackFrameFromPb(&itemPb)
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

type StackFrame struct {
	// Class from which the method call originated
	// Wire name: 'declaring_class'
	DeclaringClass string ``
	// File where the method is defined
	// Wire name: 'file_name'
	FileName string ``
	// Line from which the method was called
	// Wire name: 'line_number'
	LineNumber int ``
	// Name of the method which was called
	// Wire name: 'method_name'
	MethodName      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *StackFrame) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StackFrame) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StackFrameToPb(st *StackFrame) (*pipelinespb.StackFramePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.StackFramePb{}
	pb.DeclaringClass = st.DeclaringClass
	pb.FileName = st.FileName
	pb.LineNumber = st.LineNumber
	pb.MethodName = st.MethodName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StackFrameFromPb(pb *pipelinespb.StackFramePb) (*StackFrame, error) {
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

type StartUpdate struct {

	// Wire name: 'cause'
	Cause StartUpdateCause ``
	// If true, this update will reset all tables before running.
	// Wire name: 'full_refresh'
	FullRefresh bool ``
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'full_refresh_selection'
	FullRefreshSelection []string ``

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'refresh_selection'
	RefreshSelection []string ``
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	// Wire name: 'validate_only'
	ValidateOnly    bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *StartUpdate) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StartUpdate) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StartUpdateToPb(st *StartUpdate) (*pipelinespb.StartUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.StartUpdatePb{}
	causePb, err := StartUpdateCauseToPb(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}
	pb.FullRefresh = st.FullRefresh
	pb.FullRefreshSelection = st.FullRefreshSelection
	pb.PipelineId = st.PipelineId
	pb.RefreshSelection = st.RefreshSelection
	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StartUpdateFromPb(pb *pipelinespb.StartUpdatePb) (*StartUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartUpdate{}
	causeField, err := StartUpdateCauseFromPb(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	st.FullRefresh = pb.FullRefresh
	st.FullRefreshSelection = pb.FullRefreshSelection
	st.PipelineId = pb.PipelineId
	st.RefreshSelection = pb.RefreshSelection
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func StartUpdateCauseToPb(st *StartUpdateCause) (*pipelinespb.StartUpdateCausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.StartUpdateCausePb(*st)
	return &pb, nil
}

func StartUpdateCauseFromPb(pb *pipelinespb.StartUpdateCausePb) (*StartUpdateCause, error) {
	if pb == nil {
		return nil, nil
	}
	st := StartUpdateCause(*pb)
	return &st, nil
}

type StartUpdateResponse struct {

	// Wire name: 'update_id'
	UpdateId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *StartUpdateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StartUpdateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StartUpdateResponseToPb(st *StartUpdateResponse) (*pipelinespb.StartUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.StartUpdateResponsePb{}
	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StartUpdateResponseFromPb(pb *pipelinespb.StartUpdateResponsePb) (*StartUpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartUpdateResponse{}
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type StopRequest struct {

	// Wire name: 'pipeline_id'
	PipelineId string `tf:"-"`
}

func StopRequestToPb(st *StopRequest) (*pipelinespb.StopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.StopRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

func StopRequestFromPb(pb *pipelinespb.StopRequestPb) (*StopRequest, error) {
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
	DestinationCatalog string ``
	// Required. Destination schema to store table.
	// Wire name: 'destination_schema'
	DestinationSchema string ``
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	// Wire name: 'destination_table'
	DestinationTable string ``
	// Source catalog name. Might be optional depending on the type of source.
	// Wire name: 'source_catalog'
	SourceCatalog string ``
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	// Wire name: 'source_schema'
	SourceSchema string ``
	// Required. Table name in the source database.
	// Wire name: 'source_table'
	SourceTable string ``
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	// Wire name: 'table_configuration'
	TableConfiguration *TableSpecificConfig ``
	ForceSendFields    []string             `tf:"-"`
}

func (s *TableSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableSpecToPb(st *TableSpec) (*pipelinespb.TableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.TableSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog
	pb.DestinationSchema = st.DestinationSchema
	pb.DestinationTable = st.DestinationTable
	pb.SourceCatalog = st.SourceCatalog
	pb.SourceSchema = st.SourceSchema
	pb.SourceTable = st.SourceTable
	tableConfigurationPb, err := TableSpecificConfigToPb(st.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationPb != nil {
		pb.TableConfiguration = tableConfigurationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableSpecFromPb(pb *pipelinespb.TableSpecPb) (*TableSpec, error) {
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
	tableConfigurationField, err := TableSpecificConfigFromPb(pb.TableConfiguration)
	if err != nil {
		return nil, err
	}
	if tableConfigurationField != nil {
		st.TableConfiguration = tableConfigurationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TableSpecificConfig struct {
	// A list of column names to be excluded for the ingestion. When not
	// specified, include_columns fully controls what columns to be ingested.
	// When specified, all other columns including future ones will be
	// automatically included for ingestion. This field in mutually exclusive
	// with `include_columns`.
	// Wire name: 'exclude_columns'
	ExcludeColumns []string ``
	// A list of column names to be included for the ingestion. When not
	// specified, all columns except ones in exclude_columns will be included.
	// Future columns will be automatically included. When specified, all other
	// future columns will be automatically excluded from ingestion. This field
	// in mutually exclusive with `exclude_columns`.
	// Wire name: 'include_columns'
	IncludeColumns []string ``
	// The primary key of the table used to apply changes.
	// Wire name: 'primary_keys'
	PrimaryKeys []string ``

	// Wire name: 'query_based_connector_config'
	QueryBasedConnectorConfig *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfig ``
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	// Wire name: 'salesforce_include_formula_fields'
	SalesforceIncludeFormulaFields bool ``
	// The SCD type to use to ingest the table.
	// Wire name: 'scd_type'
	ScdType TableSpecificConfigScdType ``
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	// Wire name: 'sequence_by'
	SequenceBy      []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *TableSpecificConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableSpecificConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TableSpecificConfigToPb(st *TableSpecificConfig) (*pipelinespb.TableSpecificConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.TableSpecificConfigPb{}
	pb.ExcludeColumns = st.ExcludeColumns
	pb.IncludeColumns = st.IncludeColumns
	pb.PrimaryKeys = st.PrimaryKeys
	queryBasedConnectorConfigPb, err := IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigToPb(st.QueryBasedConnectorConfig)
	if err != nil {
		return nil, err
	}
	if queryBasedConnectorConfigPb != nil {
		pb.QueryBasedConnectorConfig = queryBasedConnectorConfigPb
	}
	pb.SalesforceIncludeFormulaFields = st.SalesforceIncludeFormulaFields
	scdTypePb, err := TableSpecificConfigScdTypeToPb(&st.ScdType)
	if err != nil {
		return nil, err
	}
	if scdTypePb != nil {
		pb.ScdType = *scdTypePb
	}
	pb.SequenceBy = st.SequenceBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TableSpecificConfigFromPb(pb *pipelinespb.TableSpecificConfigPb) (*TableSpecificConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSpecificConfig{}
	st.ExcludeColumns = pb.ExcludeColumns
	st.IncludeColumns = pb.IncludeColumns
	st.PrimaryKeys = pb.PrimaryKeys
	queryBasedConnectorConfigField, err := IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigFromPb(pb.QueryBasedConnectorConfig)
	if err != nil {
		return nil, err
	}
	if queryBasedConnectorConfigField != nil {
		st.QueryBasedConnectorConfig = queryBasedConnectorConfigField
	}
	st.SalesforceIncludeFormulaFields = pb.SalesforceIncludeFormulaFields
	scdTypeField, err := TableSpecificConfigScdTypeFromPb(&pb.ScdType)
	if err != nil {
		return nil, err
	}
	if scdTypeField != nil {
		st.ScdType = *scdTypeField
	}
	st.SequenceBy = pb.SequenceBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func TableSpecificConfigScdTypeToPb(st *TableSpecificConfigScdType) (*pipelinespb.TableSpecificConfigScdTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.TableSpecificConfigScdTypePb(*st)
	return &pb, nil
}

func TableSpecificConfigScdTypeFromPb(pb *pipelinespb.TableSpecificConfigScdTypePb) (*TableSpecificConfigScdType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TableSpecificConfigScdType(*pb)
	return &st, nil
}

type UpdateInfo struct {
	// What triggered this update.
	// Wire name: 'cause'
	Cause UpdateInfoCause ``
	// The ID of the cluster that the update is running on.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	// Wire name: 'config'
	Config *PipelineSpec ``
	// The time when this update was created.
	// Wire name: 'creation_time'
	CreationTime int64 ``
	// If true, this update will reset all tables before running.
	// Wire name: 'full_refresh'
	FullRefresh bool ``
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'full_refresh_selection'
	FullRefreshSelection []string ``
	// The ID of the pipeline.
	// Wire name: 'pipeline_id'
	PipelineId string ``
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	// Wire name: 'refresh_selection'
	RefreshSelection []string ``
	// The update state.
	// Wire name: 'state'
	State UpdateInfoState ``
	// The ID of this update.
	// Wire name: 'update_id'
	UpdateId string ``
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	// Wire name: 'validate_only'
	ValidateOnly    bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateInfoToPb(st *UpdateInfo) (*pipelinespb.UpdateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.UpdateInfoPb{}
	causePb, err := UpdateInfoCauseToPb(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}
	pb.ClusterId = st.ClusterId
	configPb, err := PipelineSpecToPb(st.Config)
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
	statePb, err := UpdateInfoStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	pb.UpdateId = st.UpdateId
	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateInfoFromPb(pb *pipelinespb.UpdateInfoPb) (*UpdateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInfo{}
	causeField, err := UpdateInfoCauseFromPb(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	st.ClusterId = pb.ClusterId
	configField, err := PipelineSpecFromPb(pb.Config)
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
	stateField, err := UpdateInfoStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	st.UpdateId = pb.UpdateId
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func UpdateInfoCauseToPb(st *UpdateInfoCause) (*pipelinespb.UpdateInfoCausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.UpdateInfoCausePb(*st)
	return &pb, nil
}

func UpdateInfoCauseFromPb(pb *pipelinespb.UpdateInfoCausePb) (*UpdateInfoCause, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateInfoCause(*pb)
	return &st, nil
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

func UpdateInfoStateToPb(st *UpdateInfoState) (*pipelinespb.UpdateInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.UpdateInfoStatePb(*st)
	return &pb, nil
}

func UpdateInfoStateFromPb(pb *pipelinespb.UpdateInfoStatePb) (*UpdateInfoState, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateInfoState(*pb)
	return &st, nil
}

type UpdateStateInfo struct {

	// Wire name: 'creation_time'
	CreationTime string ``

	// Wire name: 'state'
	State UpdateStateInfoState ``

	// Wire name: 'update_id'
	UpdateId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateStateInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateStateInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateStateInfoToPb(st *UpdateStateInfo) (*pipelinespb.UpdateStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinespb.UpdateStateInfoPb{}
	pb.CreationTime = st.CreationTime
	statePb, err := UpdateStateInfoStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateStateInfoFromPb(pb *pipelinespb.UpdateStateInfoPb) (*UpdateStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStateInfo{}
	st.CreationTime = pb.CreationTime
	stateField, err := UpdateStateInfoStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func UpdateStateInfoStateToPb(st *UpdateStateInfoState) (*pipelinespb.UpdateStateInfoStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pipelinespb.UpdateStateInfoStatePb(*st)
	return &pb, nil
}

func UpdateStateInfoStateFromPb(pb *pipelinespb.UpdateStateInfoStatePb) (*UpdateStateInfoState, error) {
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
	s := fmt.Sprintf("%.9fs", d.Seconds())
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
