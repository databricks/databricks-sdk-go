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

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames bool
	// Budget policy of this pipeline.
	BudgetPolicyId string
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog string
	// DLT Release Channel that specifies which version to use.
	Channel string
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster
	// String-String configuration for this pipeline execution.
	Configuration map[string]string
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool

	DryRun bool
	// Pipeline product edition.
	Edition string
	// Event log configuration for this pipeline
	EventLog *EventLogSpec
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition
	// Unique identifier for this pipeline.
	Id string
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary
	// Friendly identifier for this pipeline.
	Name string
	// List of notification settings for this pipeline.
	Notifications []Notifications
	// Whether Photon is enabled for this pipeline.
	Photon bool
	// Restart window of this pipeline.
	RestartWindow *RestartWindow
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *RunAs
	// The default schema (database) where tables are read from or published to.
	Schema string
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool
	// DBFS root directory for storing checkpoints and tables.
	Storage string
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger

	ForceSendFields []string
}

func createPipelineToPb(st *CreatePipeline) (*createPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPipelinePb{}
	allowDuplicateNamesPb, err := identity(&st.AllowDuplicateNames)
	if err != nil {
		return nil, err
	}
	if allowDuplicateNamesPb != nil {
		pb.AllowDuplicateNames = *allowDuplicateNamesPb
	}

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	channelPb, err := identity(&st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = *channelPb
	}

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

	configurationPb := map[string]string{}
	for k, v := range st.Configuration {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			configurationPb[k] = *itemPb
		}
	}
	pb.Configuration = configurationPb

	continuousPb, err := identity(&st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = *continuousPb
	}

	deploymentPb, err := pipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	developmentPb, err := identity(&st.Development)
	if err != nil {
		return nil, err
	}
	if developmentPb != nil {
		pb.Development = *developmentPb
	}

	dryRunPb, err := identity(&st.DryRun)
	if err != nil {
		return nil, err
	}
	if dryRunPb != nil {
		pb.DryRun = *dryRunPb
	}

	editionPb, err := identity(&st.Edition)
	if err != nil {
		return nil, err
	}
	if editionPb != nil {
		pb.Edition = *editionPb
	}

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

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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

	photonPb, err := identity(&st.Photon)
	if err != nil {
		return nil, err
	}
	if photonPb != nil {
		pb.Photon = *photonPb
	}

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

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	serverlessPb, err := identity(&st.Serverless)
	if err != nil {
		return nil, err
	}
	if serverlessPb != nil {
		pb.Serverless = *serverlessPb
	}

	storagePb, err := identity(&st.Storage)
	if err != nil {
		return nil, err
	}
	if storagePb != nil {
		pb.Storage = *storagePb
	}

	targetPb, err := identity(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

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
	allowDuplicateNamesField, err := identity(&pb.AllowDuplicateNames)
	if err != nil {
		return nil, err
	}
	if allowDuplicateNamesField != nil {
		st.AllowDuplicateNames = *allowDuplicateNamesField
	}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	channelField, err := identity(&pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = *channelField
	}

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

	configurationField := map[string]string{}
	for k, v := range pb.Configuration {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			configurationField[k] = *item
		}
	}
	st.Configuration = configurationField
	continuousField, err := identity(&pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = *continuousField
	}
	deploymentField, err := pipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	developmentField, err := identity(&pb.Development)
	if err != nil {
		return nil, err
	}
	if developmentField != nil {
		st.Development = *developmentField
	}
	dryRunField, err := identity(&pb.DryRun)
	if err != nil {
		return nil, err
	}
	if dryRunField != nil {
		st.DryRun = *dryRunField
	}
	editionField, err := identity(&pb.Edition)
	if err != nil {
		return nil, err
	}
	if editionField != nil {
		st.Edition = *editionField
	}
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	photonField, err := identity(&pb.Photon)
	if err != nil {
		return nil, err
	}
	if photonField != nil {
		st.Photon = *photonField
	}
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
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	serverlessField, err := identity(&pb.Serverless)
	if err != nil {
		return nil, err
	}
	if serverlessField != nil {
		st.Serverless = *serverlessField
	}
	storageField, err := identity(&pb.Storage)
	if err != nil {
		return nil, err
	}
	if storageField != nil {
		st.Storage = *storageField
	}
	targetField, err := identity(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}
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
	EffectiveSettings *PipelineSpec
	// The unique identifier for the newly created pipeline. Only returned when
	// dry_run is false.
	PipelineId string

	ForceSendFields []string
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

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

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
	QuartzCronSchedule string

	TimezoneId string

	ForceSendFields []string
}

func cronTriggerToPb(st *CronTrigger) (*cronTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronTriggerPb{}
	quartzCronSchedulePb, err := identity(&st.QuartzCronSchedule)
	if err != nil {
		return nil, err
	}
	if quartzCronSchedulePb != nil {
		pb.QuartzCronSchedule = *quartzCronSchedulePb
	}

	timezoneIdPb, err := identity(&st.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdPb != nil {
		pb.TimezoneId = *timezoneIdPb
	}

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
	quartzCronScheduleField, err := identity(&pb.QuartzCronSchedule)
	if err != nil {
		return nil, err
	}
	if quartzCronScheduleField != nil {
		st.QuartzCronSchedule = *quartzCronScheduleField
	}
	timezoneIdField, err := identity(&pb.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdField != nil {
		st.TimezoneId = *timezoneIdField
	}

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
	Instance string
	// A sequence number, unique and increasing within the data plane instance.
	SeqNo int64

	ForceSendFields []string
}

func dataPlaneIdToPb(st *DataPlaneId) (*dataPlaneIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneIdPb{}
	instancePb, err := identity(&st.Instance)
	if err != nil {
		return nil, err
	}
	if instancePb != nil {
		pb.Instance = *instancePb
	}

	seqNoPb, err := identity(&st.SeqNo)
	if err != nil {
		return nil, err
	}
	if seqNoPb != nil {
		pb.SeqNo = *seqNoPb
	}

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
	SeqNo int64 `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneIdFromPb(pb *dataPlaneIdPb) (*DataPlaneId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneId{}
	instanceField, err := identity(&pb.Instance)
	if err != nil {
		return nil, err
	}
	if instanceField != nil {
		st.Instance = *instanceField
	}
	seqNoField, err := identity(&pb.SeqNo)
	if err != nil {
		return nil, err
	}
	if seqNoField != nil {
		st.SeqNo = *seqNoField
	}

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
	PipelineId string
}

func deletePipelineRequestToPb(st *DeletePipelineRequest) (*deletePipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePipelineRequestPb{}
	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

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
	AllowDuplicateNames bool
	// Budget policy of this pipeline.
	BudgetPolicyId string
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog string
	// DLT Release Channel that specifies which version to use.
	Channel string
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster
	// String-String configuration for this pipeline execution.
	Configuration map[string]string
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool
	// Pipeline product edition.
	Edition string
	// Event log configuration for this pipeline
	EventLog *EventLogSpec
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified int64
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition
	// Unique identifier for this pipeline.
	Id string
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary
	// Friendly identifier for this pipeline.
	Name string
	// List of notification settings for this pipeline.
	Notifications []Notifications
	// Whether Photon is enabled for this pipeline.
	Photon bool
	// Unique identifier for this pipeline.
	PipelineId string
	// Restart window of this pipeline.
	RestartWindow *RestartWindow
	// Write-only setting, available only in Create/Update calls. Specifies the
	// user or service principal that the pipeline runs as. If not specified,
	// the pipeline runs as the user who created the pipeline.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *RunAs
	// The default schema (database) where tables are read from or published to.
	Schema string
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool
	// DBFS root directory for storing checkpoints and tables.
	Storage string
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger

	ForceSendFields []string
}

func editPipelineToPb(st *EditPipeline) (*editPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPipelinePb{}
	allowDuplicateNamesPb, err := identity(&st.AllowDuplicateNames)
	if err != nil {
		return nil, err
	}
	if allowDuplicateNamesPb != nil {
		pb.AllowDuplicateNames = *allowDuplicateNamesPb
	}

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	channelPb, err := identity(&st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = *channelPb
	}

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

	configurationPb := map[string]string{}
	for k, v := range st.Configuration {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			configurationPb[k] = *itemPb
		}
	}
	pb.Configuration = configurationPb

	continuousPb, err := identity(&st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = *continuousPb
	}

	deploymentPb, err := pipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	developmentPb, err := identity(&st.Development)
	if err != nil {
		return nil, err
	}
	if developmentPb != nil {
		pb.Development = *developmentPb
	}

	editionPb, err := identity(&st.Edition)
	if err != nil {
		return nil, err
	}
	if editionPb != nil {
		pb.Edition = *editionPb
	}

	eventLogPb, err := eventLogSpecToPb(st.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogPb != nil {
		pb.EventLog = eventLogPb
	}

	expectedLastModifiedPb, err := identity(&st.ExpectedLastModified)
	if err != nil {
		return nil, err
	}
	if expectedLastModifiedPb != nil {
		pb.ExpectedLastModified = *expectedLastModifiedPb
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

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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

	photonPb, err := identity(&st.Photon)
	if err != nil {
		return nil, err
	}
	if photonPb != nil {
		pb.Photon = *photonPb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	serverlessPb, err := identity(&st.Serverless)
	if err != nil {
		return nil, err
	}
	if serverlessPb != nil {
		pb.Serverless = *serverlessPb
	}

	storagePb, err := identity(&st.Storage)
	if err != nil {
		return nil, err
	}
	if storagePb != nil {
		pb.Storage = *storagePb
	}

	targetPb, err := identity(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

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
	PipelineId string `json:"-" url:"-"`
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
	allowDuplicateNamesField, err := identity(&pb.AllowDuplicateNames)
	if err != nil {
		return nil, err
	}
	if allowDuplicateNamesField != nil {
		st.AllowDuplicateNames = *allowDuplicateNamesField
	}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	channelField, err := identity(&pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = *channelField
	}

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

	configurationField := map[string]string{}
	for k, v := range pb.Configuration {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			configurationField[k] = *item
		}
	}
	st.Configuration = configurationField
	continuousField, err := identity(&pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = *continuousField
	}
	deploymentField, err := pipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	developmentField, err := identity(&pb.Development)
	if err != nil {
		return nil, err
	}
	if developmentField != nil {
		st.Development = *developmentField
	}
	editionField, err := identity(&pb.Edition)
	if err != nil {
		return nil, err
	}
	if editionField != nil {
		st.Edition = *editionField
	}
	eventLogField, err := eventLogSpecFromPb(pb.EventLog)
	if err != nil {
		return nil, err
	}
	if eventLogField != nil {
		st.EventLog = eventLogField
	}
	expectedLastModifiedField, err := identity(&pb.ExpectedLastModified)
	if err != nil {
		return nil, err
	}
	if expectedLastModifiedField != nil {
		st.ExpectedLastModified = *expectedLastModifiedField
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	photonField, err := identity(&pb.Photon)
	if err != nil {
		return nil, err
	}
	if photonField != nil {
		st.Photon = *photonField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
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
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	serverlessField, err := identity(&pb.Serverless)
	if err != nil {
		return nil, err
	}
	if serverlessField != nil {
		st.Serverless = *serverlessField
	}
	storageField, err := identity(&pb.Storage)
	if err != nil {
		return nil, err
	}
	if storageField != nil {
		st.Storage = *storageField
	}
	targetField, err := identity(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}
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
	Exceptions []SerializedException
	// Whether this error is considered fatal, that is, unrecoverable.
	Fatal bool

	ForceSendFields []string
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

	fatalPb, err := identity(&st.Fatal)
	if err != nil {
		return nil, err
	}
	if fatalPb != nil {
		pb.Fatal = *fatalPb
	}

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
	fatalField, err := identity(&pb.Fatal)
	if err != nil {
		return nil, err
	}
	if fatalField != nil {
		st.Fatal = *fatalField
	}

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
	Catalog string
	// The name the event log is published to in UC.
	Name string
	// The UC schema the event log is published under.
	Schema string

	ForceSendFields []string
}

func eventLogSpecToPb(st *EventLogSpec) (*eventLogSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &eventLogSpecPb{}
	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

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
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}

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
	// The absolute path of the source code.
	Path string

	ForceSendFields []string
}

func fileLibraryToPb(st *FileLibrary) (*fileLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileLibraryPb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

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
	// The absolute path of the source code.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileLibraryFromPb(pb *fileLibraryPb) (*FileLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileLibrary{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

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
	Exclude []string
	// Paths to include.
	Include []string
}

func filtersToPb(st *Filters) (*filtersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filtersPb{}

	var excludePb []string
	for _, item := range st.Exclude {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			excludePb = append(excludePb, *itemPb)
		}
	}
	pb.Exclude = excludePb

	var includePb []string
	for _, item := range st.Include {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			includePb = append(includePb, *itemPb)
		}
	}
	pb.Include = includePb

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

	var excludeField []string
	for _, itemPb := range pb.Exclude {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			excludeField = append(excludeField, *item)
		}
	}
	st.Exclude = excludeField

	var includeField []string
	for _, itemPb := range pb.Include {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			includeField = append(includeField, *item)
		}
	}
	st.Include = includeField

	return st, nil
}

// Get pipeline permission levels
type GetPipelinePermissionLevelsRequest struct {
	// The pipeline for which to get or manage permissions.
	PipelineId string
}

func getPipelinePermissionLevelsRequestToPb(st *GetPipelinePermissionLevelsRequest) (*getPipelinePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionLevelsRequestPb{}
	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	return st, nil
}

type GetPipelinePermissionLevelsResponse struct {
	// Specific permission levels
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
	PipelineId string
}

func getPipelinePermissionsRequestToPb(st *GetPipelinePermissionsRequest) (*getPipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionsRequestPb{}
	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	return st, nil
}

// Get a pipeline
type GetPipelineRequest struct {
	PipelineId string
}

func getPipelineRequestToPb(st *GetPipelineRequest) (*getPipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelineRequestPb{}
	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	return st, nil
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause string
	// The ID of the cluster that the pipeline is running on.
	ClusterId string
	// The username of the pipeline creator.
	CreatorUserName string
	// Serverless budget policy ID of this pipeline.
	EffectiveBudgetPolicyId string
	// The health of a pipeline.
	Health GetPipelineResponseHealth
	// The last time the pipeline settings were modified or created.
	LastModified int64
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name string
	// The ID of the pipeline.
	PipelineId string
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec *PipelineSpec
	// The pipeline state.
	State PipelineState

	ForceSendFields []string
}

func getPipelineResponseToPb(st *GetPipelineResponse) (*getPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelineResponsePb{}
	causePb, err := identity(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	effectiveBudgetPolicyIdPb, err := identity(&st.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdPb != nil {
		pb.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdPb
	}

	healthPb, err := identity(&st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = *healthPb
	}

	lastModifiedPb, err := identity(&st.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedPb != nil {
		pb.LastModified = *lastModifiedPb
	}

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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	runAsUserNamePb, err := identity(&st.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNamePb != nil {
		pb.RunAsUserName = *runAsUserNamePb
	}

	specPb, err := pipelineSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

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
	causeField, err := identity(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	effectiveBudgetPolicyIdField, err := identity(&pb.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdField != nil {
		st.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdField
	}
	healthField, err := identity(&pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = *healthField
	}
	lastModifiedField, err := identity(&pb.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedField != nil {
		st.LastModified = *lastModifiedField
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
	runAsUserNameField, err := identity(&pb.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNameField != nil {
		st.RunAsUserName = *runAsUserNameField
	}
	specField, err := pipelineSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

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
	PipelineId string
	// The ID of the update.
	UpdateId string
}

func getUpdateRequestToPb(st *GetUpdateRequest) (*getUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getUpdateRequestPb{}
	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	updateIdPb, err := identity(&st.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdPb != nil {
		pb.UpdateId = *updateIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
	updateIdField, err := identity(&pb.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdField != nil {
		st.UpdateId = *updateIdField
	}

	return st, nil
}

type GetUpdateResponse struct {
	// The current update info.
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
	Report *ReportSpec
	// Select all tables from a specific source schema.
	Schema *SchemaSpec
	// Select a specific source table.
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
	ConnectionId string
	// Immutable. The Unity Catalog connection that this gateway pipeline uses
	// to communicate with the source.
	ConnectionName string
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog string
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName string
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema string

	ForceSendFields []string
}

func ingestionGatewayPipelineDefinitionToPb(st *IngestionGatewayPipelineDefinition) (*ingestionGatewayPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionGatewayPipelineDefinitionPb{}
	connectionIdPb, err := identity(&st.ConnectionId)
	if err != nil {
		return nil, err
	}
	if connectionIdPb != nil {
		pb.ConnectionId = *connectionIdPb
	}

	connectionNamePb, err := identity(&st.ConnectionName)
	if err != nil {
		return nil, err
	}
	if connectionNamePb != nil {
		pb.ConnectionName = *connectionNamePb
	}

	gatewayStorageCatalogPb, err := identity(&st.GatewayStorageCatalog)
	if err != nil {
		return nil, err
	}
	if gatewayStorageCatalogPb != nil {
		pb.GatewayStorageCatalog = *gatewayStorageCatalogPb
	}

	gatewayStorageNamePb, err := identity(&st.GatewayStorageName)
	if err != nil {
		return nil, err
	}
	if gatewayStorageNamePb != nil {
		pb.GatewayStorageName = *gatewayStorageNamePb
	}

	gatewayStorageSchemaPb, err := identity(&st.GatewayStorageSchema)
	if err != nil {
		return nil, err
	}
	if gatewayStorageSchemaPb != nil {
		pb.GatewayStorageSchema = *gatewayStorageSchemaPb
	}

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
	ConnectionName string `json:"connection_name"`
	// Required, Immutable. The name of the catalog for the gateway pipeline's
	// storage location.
	GatewayStorageCatalog string `json:"gateway_storage_catalog"`
	// Optional. The Unity Catalog-compatible name for the gateway storage
	// location. This is the destination to use for the data that is extracted
	// by the gateway. Delta Live Tables system will automatically create the
	// storage location under the catalog and schema.
	GatewayStorageName string `json:"gateway_storage_name,omitempty"`
	// Required, Immutable. The name of the schema for the gateway pipelines's
	// storage location.
	GatewayStorageSchema string `json:"gateway_storage_schema"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ingestionGatewayPipelineDefinitionFromPb(pb *ingestionGatewayPipelineDefinitionPb) (*IngestionGatewayPipelineDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionGatewayPipelineDefinition{}
	connectionIdField, err := identity(&pb.ConnectionId)
	if err != nil {
		return nil, err
	}
	if connectionIdField != nil {
		st.ConnectionId = *connectionIdField
	}
	connectionNameField, err := identity(&pb.ConnectionName)
	if err != nil {
		return nil, err
	}
	if connectionNameField != nil {
		st.ConnectionName = *connectionNameField
	}
	gatewayStorageCatalogField, err := identity(&pb.GatewayStorageCatalog)
	if err != nil {
		return nil, err
	}
	if gatewayStorageCatalogField != nil {
		st.GatewayStorageCatalog = *gatewayStorageCatalogField
	}
	gatewayStorageNameField, err := identity(&pb.GatewayStorageName)
	if err != nil {
		return nil, err
	}
	if gatewayStorageNameField != nil {
		st.GatewayStorageName = *gatewayStorageNameField
	}
	gatewayStorageSchemaField, err := identity(&pb.GatewayStorageSchema)
	if err != nil {
		return nil, err
	}
	if gatewayStorageSchemaField != nil {
		st.GatewayStorageSchema = *gatewayStorageSchemaField
	}

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
	ConnectionName string
	// Immutable. Identifier for the gateway that is used by this ingestion
	// pipeline to communicate with the source database. This is used with
	// connectors to databases like SQL Server.
	IngestionGatewayId string
	// Required. Settings specifying tables to replicate and the destination for
	// the replicated tables.
	Objects []IngestionConfig
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in the pipeline.
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string
}

func ingestionPipelineDefinitionToPb(st *IngestionPipelineDefinition) (*ingestionPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionPipelineDefinitionPb{}
	connectionNamePb, err := identity(&st.ConnectionName)
	if err != nil {
		return nil, err
	}
	if connectionNamePb != nil {
		pb.ConnectionName = *connectionNamePb
	}

	ingestionGatewayIdPb, err := identity(&st.IngestionGatewayId)
	if err != nil {
		return nil, err
	}
	if ingestionGatewayIdPb != nil {
		pb.IngestionGatewayId = *ingestionGatewayIdPb
	}

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
	connectionNameField, err := identity(&pb.ConnectionName)
	if err != nil {
		return nil, err
	}
	if connectionNameField != nil {
		st.ConnectionName = *connectionNameField
	}
	ingestionGatewayIdField, err := identity(&pb.IngestionGatewayId)
	if err != nil {
		return nil, err
	}
	if ingestionGatewayIdField != nil {
		st.IngestionGatewayId = *ingestionGatewayIdField
	}

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
	Filter string
	// Max number of entries to return in a single page. The system may return
	// fewer than max_results events in a response, even if there are more
	// events available.
	MaxResults int
	// A string indicating a sort order by timestamp for the results, for
	// example, ["timestamp asc"]. The sort order can be ascending or
	// descending. By default, events are returned in descending order by
	// timestamp.
	OrderBy []string
	// Page token returned by previous call. This field is mutually exclusive
	// with all fields in this request except max_results. An error is returned
	// if any fields other than max_results are set when this field is set.
	PageToken string
	// The pipeline to return events for.
	PipelineId string

	ForceSendFields []string
}

func listPipelineEventsRequestToPb(st *ListPipelineEventsRequest) (*listPipelineEventsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelineEventsRequestPb{}
	filterPb, err := identity(&st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = *filterPb
	}

	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	var orderByPb []string
	for _, item := range st.OrderBy {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			orderByPb = append(orderByPb, *itemPb)
		}
	}
	pb.OrderBy = orderByPb

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	// The pipeline to return events for.
	PipelineId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelineEventsRequestFromPb(pb *listPipelineEventsRequestPb) (*ListPipelineEventsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelineEventsRequest{}
	filterField, err := identity(&pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = *filterField
	}
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}

	var orderByField []string
	for _, itemPb := range pb.OrderBy {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			orderByField = append(orderByField, *item)
		}
	}
	st.OrderBy = orderByField
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

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
	Events []PipelineEvent
	// If present, a token to fetch the next page of events.
	NextPageToken string
	// If present, a token to fetch the previous page of events.
	PrevPageToken string

	ForceSendFields []string
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

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	prevPageTokenPb, err := identity(&st.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenPb != nil {
		pb.PrevPageToken = *prevPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	prevPageTokenField, err := identity(&pb.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenField != nil {
		st.PrevPageToken = *prevPageTokenField
	}

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
	Filter string
	// The maximum number of entries to return in a single page. The system may
	// return fewer than max_results events in a response, even if there are
	// more events available. This field is optional. The default value is 25.
	// The maximum value is 100. An error is returned if the value of
	// max_results is greater than 100.
	MaxResults int
	// A list of strings specifying the order of results. Supported order_by
	// fields are id and name. The default is id asc. This field is optional.
	OrderBy []string
	// Page token returned by previous call
	PageToken string

	ForceSendFields []string
}

func listPipelinesRequestToPb(st *ListPipelinesRequest) (*listPipelinesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelinesRequestPb{}
	filterPb, err := identity(&st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = *filterPb
	}

	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	var orderByPb []string
	for _, item := range st.OrderBy {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			orderByPb = append(orderByPb, *itemPb)
		}
	}
	pb.OrderBy = orderByPb

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

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
	filterField, err := identity(&pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = *filterField
	}
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}

	var orderByField []string
	for _, itemPb := range pb.OrderBy {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			orderByField = append(orderByField, *item)
		}
	}
	st.OrderBy = orderByField
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

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
	NextPageToken string
	// The list of events matching the request criteria.
	Statuses []PipelineStateInfo

	ForceSendFields []string
}

func listPipelinesResponseToPb(st *ListPipelinesResponse) (*listPipelinesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelinesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	MaxResults int
	// Page token returned by previous call
	PageToken string
	// The pipeline to return updates for.
	PipelineId string
	// If present, returns updates until and including this update_id.
	UntilUpdateId string

	ForceSendFields []string
}

func listUpdatesRequestToPb(st *ListUpdatesRequest) (*listUpdatesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUpdatesRequestPb{}
	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	untilUpdateIdPb, err := identity(&st.UntilUpdateId)
	if err != nil {
		return nil, err
	}
	if untilUpdateIdPb != nil {
		pb.UntilUpdateId = *untilUpdateIdPb
	}

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
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
	untilUpdateIdField, err := identity(&pb.UntilUpdateId)
	if err != nil {
		return nil, err
	}
	if untilUpdateIdField != nil {
		st.UntilUpdateId = *untilUpdateIdField
	}

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
	NextPageToken string
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken string

	Updates []UpdateInfo

	ForceSendFields []string
}

func listUpdatesResponseToPb(st *ListUpdatesResponse) (*listUpdatesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUpdatesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	prevPageTokenPb, err := identity(&st.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenPb != nil {
		pb.PrevPageToken = *prevPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	prevPageTokenField, err := identity(&pb.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenField != nil {
		st.PrevPageToken = *prevPageTokenField
	}

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
	// The absolute path of the source code.
	Path string

	ForceSendFields []string
}

func notebookLibraryToPb(st *NotebookLibrary) (*notebookLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookLibraryPb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

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
	// The absolute path of the source code.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookLibraryFromPb(pb *notebookLibraryPb) (*NotebookLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookLibrary{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

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
	Alerts []string
	// A list of email addresses notified when a configured alert is triggered.
	EmailRecipients []string
}

func notificationsToPb(st *Notifications) (*notificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notificationsPb{}

	var alertsPb []string
	for _, item := range st.Alerts {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			alertsPb = append(alertsPb, *itemPb)
		}
	}
	pb.Alerts = alertsPb

	var emailRecipientsPb []string
	for _, item := range st.EmailRecipients {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			emailRecipientsPb = append(emailRecipientsPb, *itemPb)
		}
	}
	pb.EmailRecipients = emailRecipientsPb

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

	var alertsField []string
	for _, itemPb := range pb.Alerts {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			alertsField = append(alertsField, *item)
		}
	}
	st.Alerts = alertsField

	var emailRecipientsField []string
	for _, itemPb := range pb.EmailRecipients {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			emailRecipientsField = append(emailRecipientsField, *item)
		}
	}
	st.EmailRecipients = emailRecipientsField

	return st, nil
}

type Origin struct {
	// The id of a batch. Unique within a flow.
	BatchId int64
	// The cloud provider, e.g., AWS or Azure.
	Cloud string
	// The id of the cluster where an execution happens. Unique within a region.
	ClusterId string
	// The name of a dataset. Unique within a pipeline.
	DatasetName string
	// The id of the flow. Globally unique. Incremental queries will generally
	// reuse the same id while complete queries will have a new id per update.
	FlowId string
	// The name of the flow. Not unique.
	FlowName string
	// The optional host name where the event was triggered
	Host string
	// The id of a maintenance run. Globally unique.
	MaintenanceId string
	// Materialization name.
	MaterializationName string
	// The org id of the user. Unique within a cloud.
	OrgId int64
	// The id of the pipeline. Globally unique.
	PipelineId string
	// The name of the pipeline. Not unique.
	PipelineName string
	// The cloud region.
	Region string
	// The id of the request that caused an update.
	RequestId string
	// The id of a (delta) table. Globally unique.
	TableId string
	// The Unity Catalog id of the MV or ST being updated.
	UcResourceId string
	// The id of an execution. Globally unique.
	UpdateId string

	ForceSendFields []string
}

func originToPb(st *Origin) (*originPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &originPb{}
	batchIdPb, err := identity(&st.BatchId)
	if err != nil {
		return nil, err
	}
	if batchIdPb != nil {
		pb.BatchId = *batchIdPb
	}

	cloudPb, err := identity(&st.Cloud)
	if err != nil {
		return nil, err
	}
	if cloudPb != nil {
		pb.Cloud = *cloudPb
	}

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	datasetNamePb, err := identity(&st.DatasetName)
	if err != nil {
		return nil, err
	}
	if datasetNamePb != nil {
		pb.DatasetName = *datasetNamePb
	}

	flowIdPb, err := identity(&st.FlowId)
	if err != nil {
		return nil, err
	}
	if flowIdPb != nil {
		pb.FlowId = *flowIdPb
	}

	flowNamePb, err := identity(&st.FlowName)
	if err != nil {
		return nil, err
	}
	if flowNamePb != nil {
		pb.FlowName = *flowNamePb
	}

	hostPb, err := identity(&st.Host)
	if err != nil {
		return nil, err
	}
	if hostPb != nil {
		pb.Host = *hostPb
	}

	maintenanceIdPb, err := identity(&st.MaintenanceId)
	if err != nil {
		return nil, err
	}
	if maintenanceIdPb != nil {
		pb.MaintenanceId = *maintenanceIdPb
	}

	materializationNamePb, err := identity(&st.MaterializationName)
	if err != nil {
		return nil, err
	}
	if materializationNamePb != nil {
		pb.MaterializationName = *materializationNamePb
	}

	orgIdPb, err := identity(&st.OrgId)
	if err != nil {
		return nil, err
	}
	if orgIdPb != nil {
		pb.OrgId = *orgIdPb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	pipelineNamePb, err := identity(&st.PipelineName)
	if err != nil {
		return nil, err
	}
	if pipelineNamePb != nil {
		pb.PipelineName = *pipelineNamePb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	requestIdPb, err := identity(&st.RequestId)
	if err != nil {
		return nil, err
	}
	if requestIdPb != nil {
		pb.RequestId = *requestIdPb
	}

	tableIdPb, err := identity(&st.TableId)
	if err != nil {
		return nil, err
	}
	if tableIdPb != nil {
		pb.TableId = *tableIdPb
	}

	ucResourceIdPb, err := identity(&st.UcResourceId)
	if err != nil {
		return nil, err
	}
	if ucResourceIdPb != nil {
		pb.UcResourceId = *ucResourceIdPb
	}

	updateIdPb, err := identity(&st.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdPb != nil {
		pb.UpdateId = *updateIdPb
	}

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
	BatchId int64 `json:"batch_id,omitempty"`
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
	OrgId int64 `json:"org_id,omitempty"`
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
	batchIdField, err := identity(&pb.BatchId)
	if err != nil {
		return nil, err
	}
	if batchIdField != nil {
		st.BatchId = *batchIdField
	}
	cloudField, err := identity(&pb.Cloud)
	if err != nil {
		return nil, err
	}
	if cloudField != nil {
		st.Cloud = *cloudField
	}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	datasetNameField, err := identity(&pb.DatasetName)
	if err != nil {
		return nil, err
	}
	if datasetNameField != nil {
		st.DatasetName = *datasetNameField
	}
	flowIdField, err := identity(&pb.FlowId)
	if err != nil {
		return nil, err
	}
	if flowIdField != nil {
		st.FlowId = *flowIdField
	}
	flowNameField, err := identity(&pb.FlowName)
	if err != nil {
		return nil, err
	}
	if flowNameField != nil {
		st.FlowName = *flowNameField
	}
	hostField, err := identity(&pb.Host)
	if err != nil {
		return nil, err
	}
	if hostField != nil {
		st.Host = *hostField
	}
	maintenanceIdField, err := identity(&pb.MaintenanceId)
	if err != nil {
		return nil, err
	}
	if maintenanceIdField != nil {
		st.MaintenanceId = *maintenanceIdField
	}
	materializationNameField, err := identity(&pb.MaterializationName)
	if err != nil {
		return nil, err
	}
	if materializationNameField != nil {
		st.MaterializationName = *materializationNameField
	}
	orgIdField, err := identity(&pb.OrgId)
	if err != nil {
		return nil, err
	}
	if orgIdField != nil {
		st.OrgId = *orgIdField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
	pipelineNameField, err := identity(&pb.PipelineName)
	if err != nil {
		return nil, err
	}
	if pipelineNameField != nil {
		st.PipelineName = *pipelineNameField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}
	requestIdField, err := identity(&pb.RequestId)
	if err != nil {
		return nil, err
	}
	if requestIdField != nil {
		st.RequestId = *requestIdField
	}
	tableIdField, err := identity(&pb.TableId)
	if err != nil {
		return nil, err
	}
	if tableIdField != nil {
		st.TableId = *tableIdField
	}
	ucResourceIdField, err := identity(&pb.UcResourceId)
	if err != nil {
		return nil, err
	}
	if ucResourceIdField != nil {
		st.UcResourceId = *ucResourceIdField
	}
	updateIdField, err := identity(&pb.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdField != nil {
		st.UpdateId = *updateIdField
	}

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
	GroupName string
	// Permission level
	PermissionLevel PipelinePermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func pipelineAccessControlRequestToPb(st *PipelineAccessControlRequest) (*pipelineAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

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
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

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
	AllPermissions []PipelinePermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
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

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

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
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

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
	ApplyPolicyDefaultValues bool
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *PipelineClusterAutoscale
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *compute.AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *compute.AzureAttributes
	// The configuration for delivering spark logs to a long-term storage
	// destination. Only dbfs destinations are supported. Only one destination
	// can be specified for one cluster. If the conf is given, the logs will be
	// delivered to the destination every `5 mins`. The destination of driver
	// logs is `$destination/$clusterId/driver`, while the destination of
	// executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *compute.ClusterLogConf
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId string
	// Whether to enable local disk encryption for the cluster.
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *compute.GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []compute.InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// A label for the cluster specification, either `default` to configure the
	// default cluster, or `maintenance` to configure the maintenance cluster.
	// This field is optional. The default value is `default`.
	Label string
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. See :method:clusters/create for more
	// details.
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
	SparkEnvVars map[string]string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string

	ForceSendFields []string
}

func pipelineClusterToPb(st *PipelineCluster) (*pipelineClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineClusterPb{}
	applyPolicyDefaultValuesPb, err := identity(&st.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesPb != nil {
		pb.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesPb
	}

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

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	labelPb, err := identity(&st.Label)
	if err != nil {
		return nil, err
	}
	if labelPb != nil {
		pb.Label = *labelPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

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
	applyPolicyDefaultValuesField, err := identity(&pb.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesField != nil {
		st.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesField
	}
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

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	labelField, err := identity(&pb.Label)
	if err != nil {
		return nil, err
	}
	if labelField != nil {
		st.Label = *labelField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField

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
	MaxWorkers int
	// The minimum number of workers the cluster can scale down to when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers int
	// Databricks Enhanced Autoscaling optimizes cluster utilization by
	// automatically allocating cluster resources based on workload volume, with
	// minimal impact to the data processing latency of your pipelines. Enhanced
	// Autoscaling is available for `updates` clusters only. The legacy
	// autoscaling feature is used for `maintenance` clusters.
	Mode PipelineClusterAutoscaleMode
}

func pipelineClusterAutoscaleToPb(st *PipelineClusterAutoscale) (*pipelineClusterAutoscalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineClusterAutoscalePb{}
	maxWorkersPb, err := identity(&st.MaxWorkers)
	if err != nil {
		return nil, err
	}
	if maxWorkersPb != nil {
		pb.MaxWorkers = *maxWorkersPb
	}

	minWorkersPb, err := identity(&st.MinWorkers)
	if err != nil {
		return nil, err
	}
	if minWorkersPb != nil {
		pb.MinWorkers = *minWorkersPb
	}

	modePb, err := identity(&st.Mode)
	if err != nil {
		return nil, err
	}
	if modePb != nil {
		pb.Mode = *modePb
	}

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
	maxWorkersField, err := identity(&pb.MaxWorkers)
	if err != nil {
		return nil, err
	}
	if maxWorkersField != nil {
		st.MaxWorkers = *maxWorkersField
	}
	minWorkersField, err := identity(&pb.MinWorkers)
	if err != nil {
		return nil, err
	}
	if minWorkersField != nil {
		st.MinWorkers = *minWorkersField
	}
	modeField, err := identity(&pb.Mode)
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
	Kind DeploymentKind
	// The path to the file containing metadata about the deployment.
	MetadataFilePath string

	ForceSendFields []string
}

func pipelineDeploymentToPb(st *PipelineDeployment) (*pipelineDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineDeploymentPb{}
	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	metadataFilePathPb, err := identity(&st.MetadataFilePath)
	if err != nil {
		return nil, err
	}
	if metadataFilePathPb != nil {
		pb.MetadataFilePath = *metadataFilePathPb
	}

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
	Kind DeploymentKind `json:"kind"`
	// The path to the file containing metadata about the deployment.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineDeploymentFromPb(pb *pipelineDeploymentPb) (*PipelineDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineDeployment{}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	metadataFilePathField, err := identity(&pb.MetadataFilePath)
	if err != nil {
		return nil, err
	}
	if metadataFilePathField != nil {
		st.MetadataFilePath = *metadataFilePathField
	}

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
	Error *ErrorDetail
	// The event type. Should always correspond to the details
	EventType string
	// A time-based, globally unique id.
	Id string
	// The severity level of the event.
	Level EventLevel
	// Maturity level for event_type.
	MaturityLevel MaturityLevel
	// The display message associated with the event.
	Message string
	// Describes where the event originates from.
	Origin *Origin
	// A sequencing object to identify and order events.
	Sequence *Sequencing
	// The time of the event.
	Timestamp string

	ForceSendFields []string
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

	eventTypePb, err := identity(&st.EventType)
	if err != nil {
		return nil, err
	}
	if eventTypePb != nil {
		pb.EventType = *eventTypePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	levelPb, err := identity(&st.Level)
	if err != nil {
		return nil, err
	}
	if levelPb != nil {
		pb.Level = *levelPb
	}

	maturityLevelPb, err := identity(&st.MaturityLevel)
	if err != nil {
		return nil, err
	}
	if maturityLevelPb != nil {
		pb.MaturityLevel = *maturityLevelPb
	}

	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

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

	timestampPb, err := identity(&st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

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
	eventTypeField, err := identity(&pb.EventType)
	if err != nil {
		return nil, err
	}
	if eventTypeField != nil {
		st.EventType = *eventTypeField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	levelField, err := identity(&pb.Level)
	if err != nil {
		return nil, err
	}
	if levelField != nil {
		st.Level = *levelField
	}
	maturityLevelField, err := identity(&pb.MaturityLevel)
	if err != nil {
		return nil, err
	}
	if maturityLevelField != nil {
		st.MaturityLevel = *maturityLevelField
	}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
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
	timestampField, err := identity(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = *timestampField
	}

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
	File *FileLibrary
	// URI of the jar to be installed. Currently only DBFS is supported.
	Jar string
	// Specification of a maven library to be installed.
	Maven *compute.MavenLibrary
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace.
	Notebook *NotebookLibrary
	// URI of the whl to be installed.
	Whl string

	ForceSendFields []string
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

	jarPb, err := identity(&st.Jar)
	if err != nil {
		return nil, err
	}
	if jarPb != nil {
		pb.Jar = *jarPb
	}

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

	whlPb, err := identity(&st.Whl)
	if err != nil {
		return nil, err
	}
	if whlPb != nil {
		pb.Whl = *whlPb
	}

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
	jarField, err := identity(&pb.Jar)
	if err != nil {
		return nil, err
	}
	if jarField != nil {
		st.Jar = *jarField
	}
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
	whlField, err := identity(&pb.Whl)
	if err != nil {
		return nil, err
	}
	if whlField != nil {
		st.Whl = *whlField
	}

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
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel PipelinePermissionLevel

	ForceSendFields []string
}

func pipelinePermissionToPb(st *PipelinePermission) (*pipelinePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

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
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

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
	AccessControlList []PipelineAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
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

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

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
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

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
	Description string
	// Permission level
	PermissionLevel PipelinePermissionLevel

	ForceSendFields []string
}

func pipelinePermissionsDescriptionToPb(st *PipelinePermissionsDescription) (*pipelinePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

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
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

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
	AccessControlList []PipelineAccessControlRequest
	// The pipeline for which to get or manage permissions.
	PipelineId string
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

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	return st, nil
}

type PipelineSpec struct {
	// Budget policy of this pipeline.
	BudgetPolicyId string
	// A catalog in Unity Catalog to publish data from this pipeline to. If
	// `target` is specified, tables in this pipeline are published to a
	// `target` schema inside `catalog` (for example,
	// `catalog`.`target`.`table`). If `target` is not specified, no data is
	// published to Unity Catalog.
	Catalog string
	// DLT Release Channel that specifies which version to use.
	Channel string
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster
	// String-String configuration for this pipeline execution.
	Configuration map[string]string
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool
	// Deployment type of this pipeline.
	Deployment *PipelineDeployment
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool
	// Pipeline product edition.
	Edition string
	// Event log configuration for this pipeline
	EventLog *EventLogSpec
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters
	// The definition of a gateway pipeline to support change data capture.
	GatewayDefinition *IngestionGatewayPipelineDefinition
	// Unique identifier for this pipeline.
	Id string
	// The configuration for a managed ingestion pipeline. These settings cannot
	// be used with the 'libraries', 'schema', 'target', or 'catalog' settings.
	IngestionDefinition *IngestionPipelineDefinition
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary
	// Friendly identifier for this pipeline.
	Name string
	// List of notification settings for this pipeline.
	Notifications []Notifications
	// Whether Photon is enabled for this pipeline.
	Photon bool
	// Restart window of this pipeline.
	RestartWindow *RestartWindow
	// The default schema (database) where tables are read from or published to.
	Schema string
	// Whether serverless compute is enabled for this pipeline.
	Serverless bool
	// DBFS root directory for storing checkpoints and tables.
	Storage string
	// Target schema (database) to add tables in this pipeline to. Exactly one
	// of `schema` or `target` must be specified. To publish to Unity Catalog,
	// also specify `catalog`. This legacy field is deprecated for pipeline
	// creation in favor of the `schema` field.
	Target string
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger

	ForceSendFields []string
}

func pipelineSpecToPb(st *PipelineSpec) (*pipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineSpecPb{}
	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	channelPb, err := identity(&st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = *channelPb
	}

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

	configurationPb := map[string]string{}
	for k, v := range st.Configuration {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			configurationPb[k] = *itemPb
		}
	}
	pb.Configuration = configurationPb

	continuousPb, err := identity(&st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = *continuousPb
	}

	deploymentPb, err := pipelineDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	developmentPb, err := identity(&st.Development)
	if err != nil {
		return nil, err
	}
	if developmentPb != nil {
		pb.Development = *developmentPb
	}

	editionPb, err := identity(&st.Edition)
	if err != nil {
		return nil, err
	}
	if editionPb != nil {
		pb.Edition = *editionPb
	}

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

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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

	photonPb, err := identity(&st.Photon)
	if err != nil {
		return nil, err
	}
	if photonPb != nil {
		pb.Photon = *photonPb
	}

	restartWindowPb, err := restartWindowToPb(st.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowPb != nil {
		pb.RestartWindow = restartWindowPb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	serverlessPb, err := identity(&st.Serverless)
	if err != nil {
		return nil, err
	}
	if serverlessPb != nil {
		pb.Serverless = *serverlessPb
	}

	storagePb, err := identity(&st.Storage)
	if err != nil {
		return nil, err
	}
	if storagePb != nil {
		pb.Storage = *storagePb
	}

	targetPb, err := identity(&st.Target)
	if err != nil {
		return nil, err
	}
	if targetPb != nil {
		pb.Target = *targetPb
	}

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
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	channelField, err := identity(&pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = *channelField
	}

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

	configurationField := map[string]string{}
	for k, v := range pb.Configuration {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			configurationField[k] = *item
		}
	}
	st.Configuration = configurationField
	continuousField, err := identity(&pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = *continuousField
	}
	deploymentField, err := pipelineDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	developmentField, err := identity(&pb.Development)
	if err != nil {
		return nil, err
	}
	if developmentField != nil {
		st.Development = *developmentField
	}
	editionField, err := identity(&pb.Edition)
	if err != nil {
		return nil, err
	}
	if editionField != nil {
		st.Edition = *editionField
	}
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	photonField, err := identity(&pb.Photon)
	if err != nil {
		return nil, err
	}
	if photonField != nil {
		st.Photon = *photonField
	}
	restartWindowField, err := restartWindowFromPb(pb.RestartWindow)
	if err != nil {
		return nil, err
	}
	if restartWindowField != nil {
		st.RestartWindow = restartWindowField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	serverlessField, err := identity(&pb.Serverless)
	if err != nil {
		return nil, err
	}
	if serverlessField != nil {
		st.Serverless = *serverlessField
	}
	storageField, err := identity(&pb.Storage)
	if err != nil {
		return nil, err
	}
	if storageField != nil {
		st.Storage = *storageField
	}
	targetField, err := identity(&pb.Target)
	if err != nil {
		return nil, err
	}
	if targetField != nil {
		st.Target = *targetField
	}
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
	ClusterId string
	// The username of the pipeline creator.
	CreatorUserName string
	// The health of a pipeline.
	Health PipelineStateInfoHealth
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo
	// The user-friendly name of the pipeline.
	Name string
	// The unique identifier of the pipeline.
	PipelineId string
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName string
	// The pipeline state.
	State PipelineState

	ForceSendFields []string
}

func pipelineStateInfoToPb(st *PipelineStateInfo) (*pipelineStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineStateInfoPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	healthPb, err := identity(&st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = *healthPb
	}

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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	runAsUserNamePb, err := identity(&st.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNamePb != nil {
		pb.RunAsUserName = *runAsUserNamePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

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
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	healthField, err := identity(&pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = *healthField
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}
	runAsUserNameField, err := identity(&pb.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNameField != nil {
		st.RunAsUserName = *runAsUserNameField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

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
	Cron *CronTrigger

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
	DestinationCatalog string
	// Required. Destination schema to store table.
	DestinationSchema string
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	DestinationTable string
	// Required. Report URL in the source system.
	SourceUrl string
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object.
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string
}

func reportSpecToPb(st *ReportSpec) (*reportSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &reportSpecPb{}
	destinationCatalogPb, err := identity(&st.DestinationCatalog)
	if err != nil {
		return nil, err
	}
	if destinationCatalogPb != nil {
		pb.DestinationCatalog = *destinationCatalogPb
	}

	destinationSchemaPb, err := identity(&st.DestinationSchema)
	if err != nil {
		return nil, err
	}
	if destinationSchemaPb != nil {
		pb.DestinationSchema = *destinationSchemaPb
	}

	destinationTablePb, err := identity(&st.DestinationTable)
	if err != nil {
		return nil, err
	}
	if destinationTablePb != nil {
		pb.DestinationTable = *destinationTablePb
	}

	sourceUrlPb, err := identity(&st.SourceUrl)
	if err != nil {
		return nil, err
	}
	if sourceUrlPb != nil {
		pb.SourceUrl = *sourceUrlPb
	}

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
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema string `json:"destination_schema"`
	// Required. Destination table name. The pipeline fails if a table with that
	// name already exists.
	DestinationTable string `json:"destination_table,omitempty"`
	// Required. Report URL in the source system.
	SourceUrl string `json:"source_url"`
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
	destinationCatalogField, err := identity(&pb.DestinationCatalog)
	if err != nil {
		return nil, err
	}
	if destinationCatalogField != nil {
		st.DestinationCatalog = *destinationCatalogField
	}
	destinationSchemaField, err := identity(&pb.DestinationSchema)
	if err != nil {
		return nil, err
	}
	if destinationSchemaField != nil {
		st.DestinationSchema = *destinationSchemaField
	}
	destinationTableField, err := identity(&pb.DestinationTable)
	if err != nil {
		return nil, err
	}
	if destinationTableField != nil {
		st.DestinationTable = *destinationTableField
	}
	sourceUrlField, err := identity(&pb.SourceUrl)
	if err != nil {
		return nil, err
	}
	if sourceUrlField != nil {
		st.SourceUrl = *sourceUrlField
	}
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
	DaysOfWeek []DayOfWeek
	// An integer between 0 and 23 denoting the start hour for the restart
	// window in the 24-hour day. Continuous pipeline restart is triggered only
	// within a five-hour window starting at this hour.
	StartHour int
	// Time zone id of restart window. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details. If not specified, UTC will be used.
	TimeZoneId string

	ForceSendFields []string
}

func restartWindowToPb(st *RestartWindow) (*restartWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restartWindowPb{}

	var daysOfWeekPb []DayOfWeek
	for _, item := range st.DaysOfWeek {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			daysOfWeekPb = append(daysOfWeekPb, *itemPb)
		}
	}
	pb.DaysOfWeek = daysOfWeekPb

	startHourPb, err := identity(&st.StartHour)
	if err != nil {
		return nil, err
	}
	if startHourPb != nil {
		pb.StartHour = *startHourPb
	}

	timeZoneIdPb, err := identity(&st.TimeZoneId)
	if err != nil {
		return nil, err
	}
	if timeZoneIdPb != nil {
		pb.TimeZoneId = *timeZoneIdPb
	}

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

	var daysOfWeekField []DayOfWeek
	for _, itemPb := range pb.DaysOfWeek {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			daysOfWeekField = append(daysOfWeekField, *item)
		}
	}
	st.DaysOfWeek = daysOfWeekField
	startHourField, err := identity(&pb.StartHour)
	if err != nil {
		return nil, err
	}
	if startHourField != nil {
		st.StartHour = *startHourField
	}
	timeZoneIdField, err := identity(&pb.TimeZoneId)
	if err != nil {
		return nil, err
	}
	if timeZoneIdField != nil {
		st.TimeZoneId = *timeZoneIdField
	}

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
	ServicePrincipalName string
	// The email of an active workspace user. Users can only set this field to
	// their own email.
	UserName string

	ForceSendFields []string
}

func runAsToPb(st *RunAs) (*runAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runAsPb{}
	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

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
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

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
	DestinationCatalog string
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema string
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog string
	// Required. Schema name in the source database.
	SourceSchema string
	// Configuration settings to control the ingestion of tables. These settings
	// are applied to all tables in this schema and override the
	// table_configuration defined in the IngestionPipelineDefinition object.
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string
}

func schemaSpecToPb(st *SchemaSpec) (*schemaSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schemaSpecPb{}
	destinationCatalogPb, err := identity(&st.DestinationCatalog)
	if err != nil {
		return nil, err
	}
	if destinationCatalogPb != nil {
		pb.DestinationCatalog = *destinationCatalogPb
	}

	destinationSchemaPb, err := identity(&st.DestinationSchema)
	if err != nil {
		return nil, err
	}
	if destinationSchemaPb != nil {
		pb.DestinationSchema = *destinationSchemaPb
	}

	sourceCatalogPb, err := identity(&st.SourceCatalog)
	if err != nil {
		return nil, err
	}
	if sourceCatalogPb != nil {
		pb.SourceCatalog = *sourceCatalogPb
	}

	sourceSchemaPb, err := identity(&st.SourceSchema)
	if err != nil {
		return nil, err
	}
	if sourceSchemaPb != nil {
		pb.SourceSchema = *sourceSchemaPb
	}

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
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store tables in. Tables with the same
	// name as the source tables are created in this destination schema. The
	// pipeline fails If a table with the same name already exists.
	DestinationSchema string `json:"destination_schema"`
	// The source catalog name. Might be optional depending on the type of
	// source.
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Required. Schema name in the source database.
	SourceSchema string `json:"source_schema"`
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
	destinationCatalogField, err := identity(&pb.DestinationCatalog)
	if err != nil {
		return nil, err
	}
	if destinationCatalogField != nil {
		st.DestinationCatalog = *destinationCatalogField
	}
	destinationSchemaField, err := identity(&pb.DestinationSchema)
	if err != nil {
		return nil, err
	}
	if destinationSchemaField != nil {
		st.DestinationSchema = *destinationSchemaField
	}
	sourceCatalogField, err := identity(&pb.SourceCatalog)
	if err != nil {
		return nil, err
	}
	if sourceCatalogField != nil {
		st.SourceCatalog = *sourceCatalogField
	}
	sourceSchemaField, err := identity(&pb.SourceSchema)
	if err != nil {
		return nil, err
	}
	if sourceSchemaField != nil {
		st.SourceSchema = *sourceSchemaField
	}
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
	ControlPlaneSeqNo int64
	// the ID assigned by the data plane.
	DataPlaneId *DataPlaneId

	ForceSendFields []string
}

func sequencingToPb(st *Sequencing) (*sequencingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sequencingPb{}
	controlPlaneSeqNoPb, err := identity(&st.ControlPlaneSeqNo)
	if err != nil {
		return nil, err
	}
	if controlPlaneSeqNoPb != nil {
		pb.ControlPlaneSeqNo = *controlPlaneSeqNoPb
	}

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
	ControlPlaneSeqNo int64 `json:"control_plane_seq_no,omitempty"`
	// the ID assigned by the data plane.
	DataPlaneId *dataPlaneIdPb `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sequencingFromPb(pb *sequencingPb) (*Sequencing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Sequencing{}
	controlPlaneSeqNoField, err := identity(&pb.ControlPlaneSeqNo)
	if err != nil {
		return nil, err
	}
	if controlPlaneSeqNoField != nil {
		st.ControlPlaneSeqNo = *controlPlaneSeqNoField
	}
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
	ClassName string
	// Exception message
	Message string
	// Stack trace consisting of a list of stack frames
	Stack []StackFrame

	ForceSendFields []string
}

func serializedExceptionToPb(st *SerializedException) (*serializedExceptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serializedExceptionPb{}
	classNamePb, err := identity(&st.ClassName)
	if err != nil {
		return nil, err
	}
	if classNamePb != nil {
		pb.ClassName = *classNamePb
	}

	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

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
	classNameField, err := identity(&pb.ClassName)
	if err != nil {
		return nil, err
	}
	if classNameField != nil {
		st.ClassName = *classNameField
	}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}

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
	DeclaringClass string
	// File where the method is defined
	FileName string
	// Line from which the method was called
	LineNumber int
	// Name of the method which was called
	MethodName string

	ForceSendFields []string
}

func stackFrameToPb(st *StackFrame) (*stackFramePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stackFramePb{}
	declaringClassPb, err := identity(&st.DeclaringClass)
	if err != nil {
		return nil, err
	}
	if declaringClassPb != nil {
		pb.DeclaringClass = *declaringClassPb
	}

	fileNamePb, err := identity(&st.FileName)
	if err != nil {
		return nil, err
	}
	if fileNamePb != nil {
		pb.FileName = *fileNamePb
	}

	lineNumberPb, err := identity(&st.LineNumber)
	if err != nil {
		return nil, err
	}
	if lineNumberPb != nil {
		pb.LineNumber = *lineNumberPb
	}

	methodNamePb, err := identity(&st.MethodName)
	if err != nil {
		return nil, err
	}
	if methodNamePb != nil {
		pb.MethodName = *methodNamePb
	}

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
	declaringClassField, err := identity(&pb.DeclaringClass)
	if err != nil {
		return nil, err
	}
	if declaringClassField != nil {
		st.DeclaringClass = *declaringClassField
	}
	fileNameField, err := identity(&pb.FileName)
	if err != nil {
		return nil, err
	}
	if fileNameField != nil {
		st.FileName = *fileNameField
	}
	lineNumberField, err := identity(&pb.LineNumber)
	if err != nil {
		return nil, err
	}
	if lineNumberField != nil {
		st.LineNumber = *lineNumberField
	}
	methodNameField, err := identity(&pb.MethodName)
	if err != nil {
		return nil, err
	}
	if methodNameField != nil {
		st.MethodName = *methodNameField
	}

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
	// What triggered this update.
	Cause StartUpdateCause
	// If true, this update will reset all tables before running.
	FullRefresh bool
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string

	PipelineId string
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly bool

	ForceSendFields []string
}

func startUpdateToPb(st *StartUpdate) (*startUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startUpdatePb{}
	causePb, err := identity(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}

	fullRefreshPb, err := identity(&st.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshPb != nil {
		pb.FullRefresh = *fullRefreshPb
	}

	var fullRefreshSelectionPb []string
	for _, item := range st.FullRefreshSelection {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fullRefreshSelectionPb = append(fullRefreshSelectionPb, *itemPb)
		}
	}
	pb.FullRefreshSelection = fullRefreshSelectionPb

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	var refreshSelectionPb []string
	for _, item := range st.RefreshSelection {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			refreshSelectionPb = append(refreshSelectionPb, *itemPb)
		}
	}
	pb.RefreshSelection = refreshSelectionPb

	validateOnlyPb, err := identity(&st.ValidateOnly)
	if err != nil {
		return nil, err
	}
	if validateOnlyPb != nil {
		pb.ValidateOnly = *validateOnlyPb
	}

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
	// What triggered this update.
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
	causeField, err := identity(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	fullRefreshField, err := identity(&pb.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshField != nil {
		st.FullRefresh = *fullRefreshField
	}

	var fullRefreshSelectionField []string
	for _, itemPb := range pb.FullRefreshSelection {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fullRefreshSelectionField = append(fullRefreshSelectionField, *item)
		}
	}
	st.FullRefreshSelection = fullRefreshSelectionField
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	var refreshSelectionField []string
	for _, itemPb := range pb.RefreshSelection {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			refreshSelectionField = append(refreshSelectionField, *item)
		}
	}
	st.RefreshSelection = refreshSelectionField
	validateOnlyField, err := identity(&pb.ValidateOnly)
	if err != nil {
		return nil, err
	}
	if validateOnlyField != nil {
		st.ValidateOnly = *validateOnlyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *startUpdatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st startUpdatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// What triggered this update.
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
	UpdateId string

	ForceSendFields []string
}

func startUpdateResponseToPb(st *StartUpdateResponse) (*startUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startUpdateResponsePb{}
	updateIdPb, err := identity(&st.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdPb != nil {
		pb.UpdateId = *updateIdPb
	}

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
	updateIdField, err := identity(&pb.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdField != nil {
		st.UpdateId = *updateIdField
	}

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
	PipelineId string
}

func stopRequestToPb(st *StopRequest) (*stopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopRequestPb{}
	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

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
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	return st, nil
}

type TableSpec struct {
	// Required. Destination catalog to store table.
	DestinationCatalog string
	// Required. Destination schema to store table.
	DestinationSchema string
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable string
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog string
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema string
	// Required. Table name in the source database.
	SourceTable string
	// Configuration settings to control the ingestion of tables. These settings
	// override the table_configuration defined in the
	// IngestionPipelineDefinition object and the SchemaSpec.
	TableConfiguration *TableSpecificConfig

	ForceSendFields []string
}

func tableSpecToPb(st *TableSpec) (*tableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSpecPb{}
	destinationCatalogPb, err := identity(&st.DestinationCatalog)
	if err != nil {
		return nil, err
	}
	if destinationCatalogPb != nil {
		pb.DestinationCatalog = *destinationCatalogPb
	}

	destinationSchemaPb, err := identity(&st.DestinationSchema)
	if err != nil {
		return nil, err
	}
	if destinationSchemaPb != nil {
		pb.DestinationSchema = *destinationSchemaPb
	}

	destinationTablePb, err := identity(&st.DestinationTable)
	if err != nil {
		return nil, err
	}
	if destinationTablePb != nil {
		pb.DestinationTable = *destinationTablePb
	}

	sourceCatalogPb, err := identity(&st.SourceCatalog)
	if err != nil {
		return nil, err
	}
	if sourceCatalogPb != nil {
		pb.SourceCatalog = *sourceCatalogPb
	}

	sourceSchemaPb, err := identity(&st.SourceSchema)
	if err != nil {
		return nil, err
	}
	if sourceSchemaPb != nil {
		pb.SourceSchema = *sourceSchemaPb
	}

	sourceTablePb, err := identity(&st.SourceTable)
	if err != nil {
		return nil, err
	}
	if sourceTablePb != nil {
		pb.SourceTable = *sourceTablePb
	}

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
	DestinationCatalog string `json:"destination_catalog"`
	// Required. Destination schema to store table.
	DestinationSchema string `json:"destination_schema"`
	// Optional. Destination table name. The pipeline fails if a table with that
	// name already exists. If not set, the source table name is used.
	DestinationTable string `json:"destination_table,omitempty"`
	// Source catalog name. Might be optional depending on the type of source.
	SourceCatalog string `json:"source_catalog,omitempty"`
	// Schema name in the source database. Might be optional depending on the
	// type of source.
	SourceSchema string `json:"source_schema,omitempty"`
	// Required. Table name in the source database.
	SourceTable string `json:"source_table"`
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
	destinationCatalogField, err := identity(&pb.DestinationCatalog)
	if err != nil {
		return nil, err
	}
	if destinationCatalogField != nil {
		st.DestinationCatalog = *destinationCatalogField
	}
	destinationSchemaField, err := identity(&pb.DestinationSchema)
	if err != nil {
		return nil, err
	}
	if destinationSchemaField != nil {
		st.DestinationSchema = *destinationSchemaField
	}
	destinationTableField, err := identity(&pb.DestinationTable)
	if err != nil {
		return nil, err
	}
	if destinationTableField != nil {
		st.DestinationTable = *destinationTableField
	}
	sourceCatalogField, err := identity(&pb.SourceCatalog)
	if err != nil {
		return nil, err
	}
	if sourceCatalogField != nil {
		st.SourceCatalog = *sourceCatalogField
	}
	sourceSchemaField, err := identity(&pb.SourceSchema)
	if err != nil {
		return nil, err
	}
	if sourceSchemaField != nil {
		st.SourceSchema = *sourceSchemaField
	}
	sourceTableField, err := identity(&pb.SourceTable)
	if err != nil {
		return nil, err
	}
	if sourceTableField != nil {
		st.SourceTable = *sourceTableField
	}
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
	PrimaryKeys []string
	// If true, formula fields defined in the table are included in the
	// ingestion. This setting is only valid for the Salesforce connector
	SalesforceIncludeFormulaFields bool
	// The SCD type to use to ingest the table.
	ScdType TableSpecificConfigScdType
	// The column names specifying the logical order of events in the source
	// data. Delta Live Tables uses this sequencing to handle change events that
	// arrive out of order.
	SequenceBy []string

	ForceSendFields []string
}

func tableSpecificConfigToPb(st *TableSpecificConfig) (*tableSpecificConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSpecificConfigPb{}

	var primaryKeysPb []string
	for _, item := range st.PrimaryKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			primaryKeysPb = append(primaryKeysPb, *itemPb)
		}
	}
	pb.PrimaryKeys = primaryKeysPb

	salesforceIncludeFormulaFieldsPb, err := identity(&st.SalesforceIncludeFormulaFields)
	if err != nil {
		return nil, err
	}
	if salesforceIncludeFormulaFieldsPb != nil {
		pb.SalesforceIncludeFormulaFields = *salesforceIncludeFormulaFieldsPb
	}

	scdTypePb, err := identity(&st.ScdType)
	if err != nil {
		return nil, err
	}
	if scdTypePb != nil {
		pb.ScdType = *scdTypePb
	}

	var sequenceByPb []string
	for _, item := range st.SequenceBy {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sequenceByPb = append(sequenceByPb, *itemPb)
		}
	}
	pb.SequenceBy = sequenceByPb

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

	var primaryKeysField []string
	for _, itemPb := range pb.PrimaryKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			primaryKeysField = append(primaryKeysField, *item)
		}
	}
	st.PrimaryKeys = primaryKeysField
	salesforceIncludeFormulaFieldsField, err := identity(&pb.SalesforceIncludeFormulaFields)
	if err != nil {
		return nil, err
	}
	if salesforceIncludeFormulaFieldsField != nil {
		st.SalesforceIncludeFormulaFields = *salesforceIncludeFormulaFieldsField
	}
	scdTypeField, err := identity(&pb.ScdType)
	if err != nil {
		return nil, err
	}
	if scdTypeField != nil {
		st.ScdType = *scdTypeField
	}

	var sequenceByField []string
	for _, itemPb := range pb.SequenceBy {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sequenceByField = append(sequenceByField, *item)
		}
	}
	st.SequenceBy = sequenceByField

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
	Cause UpdateInfoCause
	// The ID of the cluster that the update is running on.
	ClusterId string
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config *PipelineSpec
	// The time when this update was created.
	CreationTime int64
	// If true, this update will reset all tables before running.
	FullRefresh bool
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string
	// The ID of the pipeline.
	PipelineId string
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string
	// The update state.
	State UpdateInfoState
	// The ID of this update.
	UpdateId string
	// If true, this update only validates the correctness of pipeline source
	// code but does not materialize or publish any datasets.
	ValidateOnly bool

	ForceSendFields []string
}

func updateInfoToPb(st *UpdateInfo) (*updateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInfoPb{}
	causePb, err := identity(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	configPb, err := pipelineSpecToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	creationTimePb, err := identity(&st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	fullRefreshPb, err := identity(&st.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshPb != nil {
		pb.FullRefresh = *fullRefreshPb
	}

	var fullRefreshSelectionPb []string
	for _, item := range st.FullRefreshSelection {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fullRefreshSelectionPb = append(fullRefreshSelectionPb, *itemPb)
		}
	}
	pb.FullRefreshSelection = fullRefreshSelectionPb

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	var refreshSelectionPb []string
	for _, item := range st.RefreshSelection {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			refreshSelectionPb = append(refreshSelectionPb, *itemPb)
		}
	}
	pb.RefreshSelection = refreshSelectionPb

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	updateIdPb, err := identity(&st.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdPb != nil {
		pb.UpdateId = *updateIdPb
	}

	validateOnlyPb, err := identity(&st.ValidateOnly)
	if err != nil {
		return nil, err
	}
	if validateOnlyPb != nil {
		pb.ValidateOnly = *validateOnlyPb
	}

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
	causeField, err := identity(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	configField, err := pipelineSpecFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	creationTimeField, err := identity(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	fullRefreshField, err := identity(&pb.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshField != nil {
		st.FullRefresh = *fullRefreshField
	}

	var fullRefreshSelectionField []string
	for _, itemPb := range pb.FullRefreshSelection {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fullRefreshSelectionField = append(fullRefreshSelectionField, *item)
		}
	}
	st.FullRefreshSelection = fullRefreshSelectionField
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	var refreshSelectionField []string
	for _, itemPb := range pb.RefreshSelection {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			refreshSelectionField = append(refreshSelectionField, *item)
		}
	}
	st.RefreshSelection = refreshSelectionField
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	updateIdField, err := identity(&pb.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdField != nil {
		st.UpdateId = *updateIdField
	}
	validateOnlyField, err := identity(&pb.ValidateOnly)
	if err != nil {
		return nil, err
	}
	if validateOnlyField != nil {
		st.ValidateOnly = *validateOnlyField
	}

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
	CreationTime string
	// The update state.
	State UpdateStateInfoState

	UpdateId string

	ForceSendFields []string
}

func updateStateInfoToPb(st *UpdateStateInfo) (*updateStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateStateInfoPb{}
	creationTimePb, err := identity(&st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	updateIdPb, err := identity(&st.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdPb != nil {
		pb.UpdateId = *updateIdPb
	}

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
	// The update state.
	State UpdateStateInfoState `json:"state,omitempty"`

	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateStateInfoFromPb(pb *updateStateInfoPb) (*UpdateStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStateInfo{}
	creationTimeField, err := identity(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	updateIdField, err := identity(&pb.UpdateId)
	if err != nil {
		return nil, err
	}
	if updateIdField != nil {
		st.UpdateId = *updateIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The update state.
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
