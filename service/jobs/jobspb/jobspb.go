// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobspb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
)

type AuthenticationMethodPb string

const AuthenticationMethodOauth AuthenticationMethodPb = `OAUTH`

const AuthenticationMethodPat AuthenticationMethodPb = `PAT`

type BaseJobPb struct {
	CreatedTime             int64                `json:"created_time,omitempty"`
	CreatorUserName         string               `json:"creator_user_name,omitempty"`
	EffectiveBudgetPolicyId string               `json:"effective_budget_policy_id,omitempty"`
	HasMore                 bool                 `json:"has_more,omitempty"`
	JobId                   int64                `json:"job_id,omitempty"`
	Settings                *JobSettingsPb       `json:"settings,omitempty"`
	TriggerState            *TriggerStateProtoPb `json:"trigger_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BaseJobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BaseJobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BaseRunPb struct {
	AttemptNumber              int                   `json:"attempt_number,omitempty"`
	CleanupDuration            int64                 `json:"cleanup_duration,omitempty"`
	ClusterInstance            *ClusterInstancePb    `json:"cluster_instance,omitempty"`
	ClusterSpec                *ClusterSpecPb        `json:"cluster_spec,omitempty"`
	CreatorUserName            string                `json:"creator_user_name,omitempty"`
	Description                string                `json:"description,omitempty"`
	EffectivePerformanceTarget PerformanceTargetPb   `json:"effective_performance_target,omitempty"`
	EndTime                    int64                 `json:"end_time,omitempty"`
	ExecutionDuration          int64                 `json:"execution_duration,omitempty"`
	GitSource                  *GitSourcePb          `json:"git_source,omitempty"`
	HasMore                    bool                  `json:"has_more,omitempty"`
	JobClusters                []JobClusterPb        `json:"job_clusters,omitempty"`
	JobId                      int64                 `json:"job_id,omitempty"`
	JobParameters              []JobParameterPb      `json:"job_parameters,omitempty"`
	JobRunId                   int64                 `json:"job_run_id,omitempty"`
	NumberInJob                int64                 `json:"number_in_job,omitempty"`
	OriginalAttemptRunId       int64                 `json:"original_attempt_run_id,omitempty"`
	OverridingParameters       *RunParametersPb      `json:"overriding_parameters,omitempty"`
	QueueDuration              int64                 `json:"queue_duration,omitempty"`
	RepairHistory              []RepairHistoryItemPb `json:"repair_history,omitempty"`
	RunDuration                int64                 `json:"run_duration,omitempty"`
	RunId                      int64                 `json:"run_id,omitempty"`
	RunName                    string                `json:"run_name,omitempty"`
	RunPageUrl                 string                `json:"run_page_url,omitempty"`
	RunType                    RunTypePb             `json:"run_type,omitempty"`
	Schedule                   *CronSchedulePb       `json:"schedule,omitempty"`
	SetupDuration              int64                 `json:"setup_duration,omitempty"`
	StartTime                  int64                 `json:"start_time,omitempty"`
	State                      *RunStatePb           `json:"state,omitempty"`
	Status                     *RunStatusPb          `json:"status,omitempty"`
	Tasks                      []RunTaskPb           `json:"tasks,omitempty"`
	Trigger                    TriggerTypePb         `json:"trigger,omitempty"`
	TriggerInfo                *TriggerInfoPb        `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BaseRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BaseRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelAllRunsPb struct {
	AllQueuedRuns bool  `json:"all_queued_runs,omitempty"`
	JobId         int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CancelAllRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CancelAllRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelRunPb struct {
	RunId int64 `json:"run_id"`
}

type CleanRoomTaskRunLifeCycleStatePb string

const CleanRoomTaskRunLifeCycleStateBlocked CleanRoomTaskRunLifeCycleStatePb = `BLOCKED`

const CleanRoomTaskRunLifeCycleStateInternalError CleanRoomTaskRunLifeCycleStatePb = `INTERNAL_ERROR`

const CleanRoomTaskRunLifeCycleStatePending CleanRoomTaskRunLifeCycleStatePb = `PENDING`

const CleanRoomTaskRunLifeCycleStateQueued CleanRoomTaskRunLifeCycleStatePb = `QUEUED`

const CleanRoomTaskRunLifeCycleStateRunning CleanRoomTaskRunLifeCycleStatePb = `RUNNING`

const CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified CleanRoomTaskRunLifeCycleStatePb = `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`

const CleanRoomTaskRunLifeCycleStateSkipped CleanRoomTaskRunLifeCycleStatePb = `SKIPPED`

const CleanRoomTaskRunLifeCycleStateTerminated CleanRoomTaskRunLifeCycleStatePb = `TERMINATED`

const CleanRoomTaskRunLifeCycleStateTerminating CleanRoomTaskRunLifeCycleStatePb = `TERMINATING`

const CleanRoomTaskRunLifeCycleStateWaitingForRetry CleanRoomTaskRunLifeCycleStatePb = `WAITING_FOR_RETRY`

type CleanRoomTaskRunResultStatePb string

const CleanRoomTaskRunResultStateCanceled CleanRoomTaskRunResultStatePb = `CANCELED`

const CleanRoomTaskRunResultStateDisabled CleanRoomTaskRunResultStatePb = `DISABLED`

const CleanRoomTaskRunResultStateEvicted CleanRoomTaskRunResultStatePb = `EVICTED`

const CleanRoomTaskRunResultStateExcluded CleanRoomTaskRunResultStatePb = `EXCLUDED`

const CleanRoomTaskRunResultStateFailed CleanRoomTaskRunResultStatePb = `FAILED`

const CleanRoomTaskRunResultStateMaximumConcurrentRunsReached CleanRoomTaskRunResultStatePb = `MAXIMUM_CONCURRENT_RUNS_REACHED`

const CleanRoomTaskRunResultStateRunResultStateUnspecified CleanRoomTaskRunResultStatePb = `RUN_RESULT_STATE_UNSPECIFIED`

const CleanRoomTaskRunResultStateSuccess CleanRoomTaskRunResultStatePb = `SUCCESS`

const CleanRoomTaskRunResultStateSuccessWithFailures CleanRoomTaskRunResultStatePb = `SUCCESS_WITH_FAILURES`

const CleanRoomTaskRunResultStateTimedout CleanRoomTaskRunResultStatePb = `TIMEDOUT`

const CleanRoomTaskRunResultStateUpstreamCanceled CleanRoomTaskRunResultStatePb = `UPSTREAM_CANCELED`

const CleanRoomTaskRunResultStateUpstreamEvicted CleanRoomTaskRunResultStatePb = `UPSTREAM_EVICTED`

const CleanRoomTaskRunResultStateUpstreamFailed CleanRoomTaskRunResultStatePb = `UPSTREAM_FAILED`

type CleanRoomTaskRunStatePb struct {
	LifeCycleState CleanRoomTaskRunLifeCycleStatePb `json:"life_cycle_state,omitempty"`
	ResultState    CleanRoomTaskRunResultStatePb    `json:"result_state,omitempty"`
}

type CleanRoomsNotebookTaskPb struct {
	CleanRoomName          string            `json:"clean_room_name"`
	Etag                   string            `json:"etag,omitempty"`
	NotebookBaseParameters map[string]string `json:"notebook_base_parameters,omitempty"`
	NotebookName           string            `json:"notebook_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CleanRoomsNotebookTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CleanRoomsNotebookTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb struct {
	CleanRoomJobRunState *CleanRoomTaskRunStatePb `json:"clean_room_job_run_state,omitempty"`
	NotebookOutput       *NotebookOutputPb        `json:"notebook_output,omitempty"`
	OutputSchemaInfo     *OutputSchemaInfoPb      `json:"output_schema_info,omitempty"`
}

type ClusterInstancePb struct {
	ClusterId      string `json:"cluster_id,omitempty"`
	SparkContextId string `json:"spark_context_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterInstancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterInstancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterSpecPb struct {
	ExistingClusterId string                   `json:"existing_cluster_id,omitempty"`
	JobClusterKey     string                   `json:"job_cluster_key,omitempty"`
	Libraries         []computepb.LibraryPb    `json:"libraries,omitempty"`
	NewCluster        *computepb.ClusterSpecPb `json:"new_cluster,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComputeConfigPb struct {
	GpuNodePoolId string `json:"gpu_node_pool_id,omitempty"`
	GpuType       string `json:"gpu_type,omitempty"`
	NumGpus       int    `json:"num_gpus"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComputeConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComputeConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ConditionPb string

const ConditionAllUpdated ConditionPb = `ALL_UPDATED`

const ConditionAnyUpdated ConditionPb = `ANY_UPDATED`

type ConditionTaskPb struct {
	Left  string            `json:"left"`
	Op    ConditionTaskOpPb `json:"op"`
	Right string            `json:"right"`
}

type ConditionTaskOpPb string

const ConditionTaskOpEqualTo ConditionTaskOpPb = `EQUAL_TO`

const ConditionTaskOpGreaterThan ConditionTaskOpPb = `GREATER_THAN`

const ConditionTaskOpGreaterThanOrEqual ConditionTaskOpPb = `GREATER_THAN_OR_EQUAL`

const ConditionTaskOpLessThan ConditionTaskOpPb = `LESS_THAN`

const ConditionTaskOpLessThanOrEqual ConditionTaskOpPb = `LESS_THAN_OR_EQUAL`

const ConditionTaskOpNotEqual ConditionTaskOpPb = `NOT_EQUAL`

type ContinuousPb struct {
	PauseStatus PauseStatusPb `json:"pause_status,omitempty"`
}

type CreateJobPb struct {
	AccessControlList    []JobAccessControlRequestPb `json:"access_control_list,omitempty"`
	BudgetPolicyId       string                      `json:"budget_policy_id,omitempty"`
	Continuous           *ContinuousPb               `json:"continuous,omitempty"`
	Deployment           *JobDeploymentPb            `json:"deployment,omitempty"`
	Description          string                      `json:"description,omitempty"`
	EditMode             JobEditModePb               `json:"edit_mode,omitempty"`
	EmailNotifications   *JobEmailNotificationsPb    `json:"email_notifications,omitempty"`
	Environments         []JobEnvironmentPb          `json:"environments,omitempty"`
	Format               FormatPb                    `json:"format,omitempty"`
	GitSource            *GitSourcePb                `json:"git_source,omitempty"`
	Health               *JobsHealthRulesPb          `json:"health,omitempty"`
	JobClusters          []JobClusterPb              `json:"job_clusters,omitempty"`
	MaxConcurrentRuns    int                         `json:"max_concurrent_runs,omitempty"`
	Name                 string                      `json:"name,omitempty"`
	NotificationSettings *JobNotificationSettingsPb  `json:"notification_settings,omitempty"`
	Parameters           []JobParameterDefinitionPb  `json:"parameters,omitempty"`
	PerformanceTarget    PerformanceTargetPb         `json:"performance_target,omitempty"`
	Queue                *QueueSettingsPb            `json:"queue,omitempty"`
	RunAs                *JobRunAsPb                 `json:"run_as,omitempty"`
	Schedule             *CronSchedulePb             `json:"schedule,omitempty"`
	Tags                 map[string]string           `json:"tags,omitempty"`
	Tasks                []TaskPb                    `json:"tasks,omitempty"`
	TimeoutSeconds       int                         `json:"timeout_seconds,omitempty"`
	Trigger              *TriggerSettingsPb          `json:"trigger,omitempty"`
	WebhookNotifications *WebhookNotificationsPb     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateJobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateJobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateResponsePb struct {
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CronSchedulePb struct {
	PauseStatus          PauseStatusPb `json:"pause_status,omitempty"`
	QuartzCronExpression string        `json:"quartz_cron_expression"`
	TimezoneId           string        `json:"timezone_id"`
}

type DashboardPageSnapshotPb struct {
	PageDisplayName    string                `json:"page_display_name,omitempty"`
	WidgetErrorDetails []WidgetErrorDetailPb `json:"widget_error_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardPageSnapshotPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardPageSnapshotPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardTaskPb struct {
	DashboardId  string          `json:"dashboard_id,omitempty"`
	Subscription *SubscriptionPb `json:"subscription,omitempty"`
	WarehouseId  string          `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardTaskOutputPb struct {
	PageSnapshots []DashboardPageSnapshotPb `json:"page_snapshots,omitempty"`
}

type DbtCloudJobRunStepPb struct {
	Index  int                    `json:"index,omitempty"`
	Logs   string                 `json:"logs,omitempty"`
	Name   string                 `json:"name,omitempty"`
	Status DbtPlatformRunStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtCloudJobRunStepPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtCloudJobRunStepPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtCloudTaskPb struct {
	ConnectionResourceName string `json:"connection_resource_name,omitempty"`
	DbtCloudJobId          int64  `json:"dbt_cloud_job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtCloudTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtCloudTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtCloudTaskOutputPb struct {
	DbtCloudJobRunId     int64                  `json:"dbt_cloud_job_run_id,omitempty"`
	DbtCloudJobRunOutput []DbtCloudJobRunStepPb `json:"dbt_cloud_job_run_output,omitempty"`
	DbtCloudJobRunUrl    string                 `json:"dbt_cloud_job_run_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtCloudTaskOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtCloudTaskOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtOutputPb struct {
	ArtifactsHeaders map[string]string `json:"artifacts_headers,omitempty"`
	ArtifactsLink    string            `json:"artifacts_link,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtPlatformJobRunStepPb struct {
	Index         int                    `json:"index,omitempty"`
	Logs          string                 `json:"logs,omitempty"`
	LogsTruncated bool                   `json:"logs_truncated,omitempty"`
	Name          string                 `json:"name,omitempty"`
	NameTruncated bool                   `json:"name_truncated,omitempty"`
	Status        DbtPlatformRunStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtPlatformJobRunStepPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtPlatformJobRunStepPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtPlatformRunStatusPb string

const DbtPlatformRunStatusCancelled DbtPlatformRunStatusPb = `CANCELLED`

const DbtPlatformRunStatusError DbtPlatformRunStatusPb = `ERROR`

const DbtPlatformRunStatusQueued DbtPlatformRunStatusPb = `QUEUED`

const DbtPlatformRunStatusRunning DbtPlatformRunStatusPb = `RUNNING`

const DbtPlatformRunStatusStarting DbtPlatformRunStatusPb = `STARTING`

const DbtPlatformRunStatusSuccess DbtPlatformRunStatusPb = `SUCCESS`

type DbtPlatformTaskPb struct {
	ConnectionResourceName string `json:"connection_resource_name,omitempty"`
	DbtPlatformJobId       string `json:"dbt_platform_job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtPlatformTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtPlatformTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtPlatformTaskOutputPb struct {
	DbtPlatformJobRunId     string                    `json:"dbt_platform_job_run_id,omitempty"`
	DbtPlatformJobRunOutput []DbtPlatformJobRunStepPb `json:"dbt_platform_job_run_output,omitempty"`
	DbtPlatformJobRunUrl    string                    `json:"dbt_platform_job_run_url,omitempty"`
	StepsTruncated          bool                      `json:"steps_truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtPlatformTaskOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtPlatformTaskOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtTaskPb struct {
	Catalog           string   `json:"catalog,omitempty"`
	Commands          []string `json:"commands"`
	ProfilesDirectory string   `json:"profiles_directory,omitempty"`
	ProjectDirectory  string   `json:"project_directory,omitempty"`
	Schema            string   `json:"schema,omitempty"`
	Source            SourcePb `json:"source,omitempty"`
	WarehouseId       string   `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DbtTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DbtTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteJobPb struct {
	JobId int64 `json:"job_id"`
}

type DeleteRunPb struct {
	RunId int64 `json:"run_id"`
}

type EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb struct {
	Field         string `json:"field,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnforcePolicyComplianceRequestPb struct {
	JobId        int64 `json:"job_id"`
	ValidateOnly bool  `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnforcePolicyComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnforcePolicyComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnforcePolicyComplianceResponsePb struct {
	HasChanges        bool                                                              `json:"has_changes,omitempty"`
	JobClusterChanges []EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb `json:"job_cluster_changes,omitempty"`
	Settings          *JobSettingsPb                                                    `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnforcePolicyComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnforcePolicyComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExportRunOutputPb struct {
	Views []ViewItemPb `json:"views,omitempty"`
}

type ExportRunRequestPb struct {
	RunId         int64           `json:"-" url:"run_id"`
	ViewsToExport ViewsToExportPb `json:"-" url:"views_to_export,omitempty"`
}

type FileArrivalTriggerConfigurationPb struct {
	MinTimeBetweenTriggersSeconds int    `json:"min_time_between_triggers_seconds,omitempty"`
	Url                           string `json:"url"`
	WaitAfterLastChangeSeconds    int    `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileArrivalTriggerConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileArrivalTriggerConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileArrivalTriggerStatePb struct {
	UsingFileEvents bool `json:"using_file_events,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileArrivalTriggerStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileArrivalTriggerStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForEachStatsPb struct {
	ErrorMessageStats []ForEachTaskErrorMessageStatsPb `json:"error_message_stats,omitempty"`
	TaskRunStats      *ForEachTaskTaskRunStatsPb       `json:"task_run_stats,omitempty"`
}

type ForEachTaskPb struct {
	Concurrency int    `json:"concurrency,omitempty"`
	Inputs      string `json:"inputs"`
	Task        TaskPb `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ForEachTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ForEachTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForEachTaskErrorMessageStatsPb struct {
	Count               int    `json:"count,omitempty"`
	ErrorMessage        string `json:"error_message,omitempty"`
	TerminationCategory string `json:"termination_category,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ForEachTaskErrorMessageStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ForEachTaskErrorMessageStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForEachTaskTaskRunStatsPb struct {
	ActiveIterations    int `json:"active_iterations,omitempty"`
	CompletedIterations int `json:"completed_iterations,omitempty"`
	FailedIterations    int `json:"failed_iterations,omitempty"`
	ScheduledIterations int `json:"scheduled_iterations,omitempty"`
	SucceededIterations int `json:"succeeded_iterations,omitempty"`
	TotalIterations     int `json:"total_iterations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ForEachTaskTaskRunStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ForEachTaskTaskRunStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FormatPb string

const FormatMultiTask FormatPb = `MULTI_TASK`

const FormatSingleTask FormatPb = `SINGLE_TASK`

type GenAiComputeTaskPb struct {
	Command                string           `json:"command,omitempty"`
	Compute                *ComputeConfigPb `json:"compute,omitempty"`
	DlRuntimeImage         string           `json:"dl_runtime_image"`
	MlflowExperimentName   string           `json:"mlflow_experiment_name,omitempty"`
	Source                 SourcePb         `json:"source,omitempty"`
	TrainingScriptPath     string           `json:"training_script_path,omitempty"`
	YamlParameters         string           `json:"yaml_parameters,omitempty"`
	YamlParametersFilePath string           `json:"yaml_parameters_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenAiComputeTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenAiComputeTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetJobPermissionLevelsRequestPb struct {
	JobId string `json:"-" url:"-"`
}

type GetJobPermissionLevelsResponsePb struct {
	PermissionLevels []JobPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetJobPermissionsRequestPb struct {
	JobId string `json:"-" url:"-"`
}

type GetJobRequestPb struct {
	JobId     int64  `json:"-" url:"job_id"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetJobRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetJobRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPolicyComplianceRequestPb struct {
	JobId int64 `json:"-" url:"job_id"`
}

type GetPolicyComplianceResponsePb struct {
	IsCompliant bool              `json:"is_compliant,omitempty"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPolicyComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPolicyComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetRunOutputRequestPb struct {
	RunId int64 `json:"-" url:"run_id"`
}

type GetRunRequestPb struct {
	IncludeHistory        bool   `json:"-" url:"include_history,omitempty"`
	IncludeResolvedValues bool   `json:"-" url:"include_resolved_values,omitempty"`
	PageToken             string `json:"-" url:"page_token,omitempty"`
	RunId                 int64  `json:"-" url:"run_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetRunRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetRunRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GitProviderPb string

const GitProviderAwsCodeCommit GitProviderPb = `awsCodeCommit`

const GitProviderAzureDevOpsServices GitProviderPb = `azureDevOpsServices`

const GitProviderBitbucketCloud GitProviderPb = `bitbucketCloud`

const GitProviderBitbucketServer GitProviderPb = `bitbucketServer`

const GitProviderGitHub GitProviderPb = `gitHub`

const GitProviderGitHubEnterprise GitProviderPb = `gitHubEnterprise`

const GitProviderGitLab GitProviderPb = `gitLab`

const GitProviderGitLabEnterpriseEdition GitProviderPb = `gitLabEnterpriseEdition`

type GitSnapshotPb struct {
	UsedCommit string `json:"used_commit,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GitSnapshotPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GitSnapshotPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GitSourcePb struct {
	GitBranch   string         `json:"git_branch,omitempty"`
	GitCommit   string         `json:"git_commit,omitempty"`
	GitProvider GitProviderPb  `json:"git_provider"`
	GitSnapshot *GitSnapshotPb `json:"git_snapshot,omitempty"`
	GitTag      string         `json:"git_tag,omitempty"`
	GitUrl      string         `json:"git_url"`
	JobSource   *JobSourcePb   `json:"job_source,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GitSourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GitSourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobPb struct {
	CreatedTime             int64                `json:"created_time,omitempty"`
	CreatorUserName         string               `json:"creator_user_name,omitempty"`
	EffectiveBudgetPolicyId string               `json:"effective_budget_policy_id,omitempty"`
	HasMore                 bool                 `json:"has_more,omitempty"`
	JobId                   int64                `json:"job_id,omitempty"`
	NextPageToken           string               `json:"next_page_token,omitempty"`
	RunAsUserName           string               `json:"run_as_user_name,omitempty"`
	Settings                *JobSettingsPb       `json:"settings,omitempty"`
	TriggerState            *TriggerStateProtoPb `json:"trigger_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobAccessControlRequestPb struct {
	GroupName            string               `json:"group_name,omitempty"`
	PermissionLevel      JobPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string               `json:"service_principal_name,omitempty"`
	UserName             string               `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobAccessControlResponsePb struct {
	AllPermissions       []JobPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string            `json:"display_name,omitempty"`
	GroupName            string            `json:"group_name,omitempty"`
	ServicePrincipalName string            `json:"service_principal_name,omitempty"`
	UserName             string            `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobClusterPb struct {
	JobClusterKey string                  `json:"job_cluster_key"`
	NewCluster    computepb.ClusterSpecPb `json:"new_cluster"`
}

type JobCompliancePb struct {
	IsCompliant bool              `json:"is_compliant,omitempty"`
	JobId       int64             `json:"job_id"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobCompliancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobCompliancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobDeploymentPb struct {
	Kind             JobDeploymentKindPb `json:"kind"`
	MetadataFilePath string              `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobDeploymentKindPb string

// The job is managed by Databricks Asset Bundle.
const JobDeploymentKindBundle JobDeploymentKindPb = `BUNDLE`

type JobEditModePb string

// The job is in an editable state and can be modified.
const JobEditModeEditable JobEditModePb = `EDITABLE`

// The job is in a locked UI state and cannot be modified.
const JobEditModeUiLocked JobEditModePb = `UI_LOCKED`

type JobEmailNotificationsPb struct {
	NoAlertForSkippedRuns              bool     `json:"no_alert_for_skipped_runs,omitempty"`
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	OnFailure                          []string `json:"on_failure,omitempty"`
	OnStart                            []string `json:"on_start,omitempty"`
	OnStreamingBacklogExceeded         []string `json:"on_streaming_backlog_exceeded,omitempty"`
	OnSuccess                          []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobEmailNotificationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobEmailNotificationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobEnvironmentPb struct {
	EnvironmentKey string                   `json:"environment_key"`
	Spec           *computepb.EnvironmentPb `json:"spec,omitempty"`
}

type JobNotificationSettingsPb struct {
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	NoAlertForSkippedRuns  bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobNotificationSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobNotificationSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobParameterPb struct {
	Default string `json:"default,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobParameterDefinitionPb struct {
	Default string `json:"default"`
	Name    string `json:"name"`
}

type JobPermissionPb struct {
	Inherited           bool                 `json:"inherited,omitempty"`
	InheritedFromObject []string             `json:"inherited_from_object,omitempty"`
	PermissionLevel     JobPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobPermissionLevelPb string

const JobPermissionLevelCanManage JobPermissionLevelPb = `CAN_MANAGE`

const JobPermissionLevelCanManageRun JobPermissionLevelPb = `CAN_MANAGE_RUN`

const JobPermissionLevelCanView JobPermissionLevelPb = `CAN_VIEW`

const JobPermissionLevelIsOwner JobPermissionLevelPb = `IS_OWNER`

type JobPermissionsPb struct {
	AccessControlList []JobAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                       `json:"object_id,omitempty"`
	ObjectType        string                       `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobPermissionsDescriptionPb struct {
	Description     string               `json:"description,omitempty"`
	PermissionLevel JobPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobPermissionsRequestPb struct {
	AccessControlList []JobAccessControlRequestPb `json:"access_control_list,omitempty"`
	JobId             string                      `json:"-" url:"-"`
}

type JobRunAsPb struct {
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobRunAsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobRunAsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobSettingsPb struct {
	BudgetPolicyId       string                     `json:"budget_policy_id,omitempty"`
	Continuous           *ContinuousPb              `json:"continuous,omitempty"`
	Deployment           *JobDeploymentPb           `json:"deployment,omitempty"`
	Description          string                     `json:"description,omitempty"`
	EditMode             JobEditModePb              `json:"edit_mode,omitempty"`
	EmailNotifications   *JobEmailNotificationsPb   `json:"email_notifications,omitempty"`
	Environments         []JobEnvironmentPb         `json:"environments,omitempty"`
	Format               FormatPb                   `json:"format,omitempty"`
	GitSource            *GitSourcePb               `json:"git_source,omitempty"`
	Health               *JobsHealthRulesPb         `json:"health,omitempty"`
	JobClusters          []JobClusterPb             `json:"job_clusters,omitempty"`
	MaxConcurrentRuns    int                        `json:"max_concurrent_runs,omitempty"`
	Name                 string                     `json:"name,omitempty"`
	NotificationSettings *JobNotificationSettingsPb `json:"notification_settings,omitempty"`
	Parameters           []JobParameterDefinitionPb `json:"parameters,omitempty"`
	PerformanceTarget    PerformanceTargetPb        `json:"performance_target,omitempty"`
	Queue                *QueueSettingsPb           `json:"queue,omitempty"`
	RunAs                *JobRunAsPb                `json:"run_as,omitempty"`
	Schedule             *CronSchedulePb            `json:"schedule,omitempty"`
	Tags                 map[string]string          `json:"tags,omitempty"`
	Tasks                []TaskPb                   `json:"tasks,omitempty"`
	TimeoutSeconds       int                        `json:"timeout_seconds,omitempty"`
	Trigger              *TriggerSettingsPb         `json:"trigger,omitempty"`
	WebhookNotifications *WebhookNotificationsPb    `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *JobSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st JobSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobSourcePb struct {
	DirtyState          JobSourceDirtyStatePb `json:"dirty_state,omitempty"`
	ImportFromGitBranch string                `json:"import_from_git_branch"`
	JobConfigPath       string                `json:"job_config_path"`
}

type JobSourceDirtyStatePb string

// The job is temporary disconnected from the remote job specification and is
// allowed for live edit. Import the remote job specification again from UI to
// make the job fully synced.
const JobSourceDirtyStateDisconnected JobSourceDirtyStatePb = `DISCONNECTED`

// The job is not yet synced with the remote job specification. Import the
// remote job specification from UI to make the job fully synced.
const JobSourceDirtyStateNotSynced JobSourceDirtyStatePb = `NOT_SYNCED`

type JobsHealthMetricPb string

// Expected total time for a run in seconds.
const JobsHealthMetricRunDurationSeconds JobsHealthMetricPb = `RUN_DURATION_SECONDS`

// An estimate of the maximum bytes of data waiting to be consumed across all
// streams. This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogBytes JobsHealthMetricPb = `STREAMING_BACKLOG_BYTES`

// An estimate of the maximum number of outstanding files across all streams.
// This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogFiles JobsHealthMetricPb = `STREAMING_BACKLOG_FILES`

// An estimate of the maximum offset lag across all streams. This metric is in
// Public Preview.
const JobsHealthMetricStreamingBacklogRecords JobsHealthMetricPb = `STREAMING_BACKLOG_RECORDS`

// An estimate of the maximum consumer delay across all streams. This metric is
// in Public Preview.
const JobsHealthMetricStreamingBacklogSeconds JobsHealthMetricPb = `STREAMING_BACKLOG_SECONDS`

type JobsHealthOperatorPb string

const JobsHealthOperatorGreaterThan JobsHealthOperatorPb = `GREATER_THAN`

type JobsHealthRulePb struct {
	Metric JobsHealthMetricPb   `json:"metric"`
	Op     JobsHealthOperatorPb `json:"op"`
	Value  int64                `json:"value"`
}

type JobsHealthRulesPb struct {
	Rules []JobsHealthRulePb `json:"rules,omitempty"`
}

type ListJobComplianceForPolicyResponsePb struct {
	Jobs          []JobCompliancePb `json:"jobs,omitempty"`
	NextPageToken string            `json:"next_page_token,omitempty"`
	PrevPageToken string            `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListJobComplianceForPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListJobComplianceForPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListJobComplianceRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`
	PolicyId  string `json:"-" url:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListJobComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListJobComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListJobsRequestPb struct {
	ExpandTasks bool   `json:"-" url:"expand_tasks,omitempty"`
	Limit       int    `json:"-" url:"limit,omitempty"`
	Name        string `json:"-" url:"name,omitempty"`
	Offset      int    `json:"-" url:"offset,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListJobsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListJobsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListJobsResponsePb struct {
	HasMore       bool        `json:"has_more,omitempty"`
	Jobs          []BaseJobPb `json:"jobs,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`
	PrevPageToken string      `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListJobsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListJobsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRunsRequestPb struct {
	ActiveOnly    bool      `json:"-" url:"active_only,omitempty"`
	CompletedOnly bool      `json:"-" url:"completed_only,omitempty"`
	ExpandTasks   bool      `json:"-" url:"expand_tasks,omitempty"`
	JobId         int64     `json:"-" url:"job_id,omitempty"`
	Limit         int       `json:"-" url:"limit,omitempty"`
	Offset        int       `json:"-" url:"offset,omitempty"`
	PageToken     string    `json:"-" url:"page_token,omitempty"`
	RunType       RunTypePb `json:"-" url:"run_type,omitempty"`
	StartTimeFrom int64     `json:"-" url:"start_time_from,omitempty"`
	StartTimeTo   int64     `json:"-" url:"start_time_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRunsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRunsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRunsResponsePb struct {
	HasMore       bool        `json:"has_more,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`
	PrevPageToken string      `json:"prev_page_token,omitempty"`
	Runs          []BaseRunPb `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookOutputPb struct {
	Result    string `json:"result,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NotebookOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NotebookOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookTaskPb struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
	NotebookPath   string            `json:"notebook_path"`
	Source         SourcePb          `json:"source,omitempty"`
	WarehouseId    string            `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NotebookTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NotebookTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OutputSchemaInfoPb struct {
	CatalogName    string `json:"catalog_name,omitempty"`
	ExpirationTime int64  `json:"expiration_time,omitempty"`
	SchemaName     string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OutputSchemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OutputSchemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PauseStatusPb string

const PauseStatusPaused PauseStatusPb = `PAUSED`

const PauseStatusUnpaused PauseStatusPb = `UNPAUSED`

type PerformanceTargetPb string

const PerformanceTargetPerformanceOptimized PerformanceTargetPb = `PERFORMANCE_OPTIMIZED`

const PerformanceTargetStandard PerformanceTargetPb = `STANDARD`

type PeriodicTriggerConfigurationPb struct {
	Interval int                                    `json:"interval"`
	Unit     PeriodicTriggerConfigurationTimeUnitPb `json:"unit"`
}

type PeriodicTriggerConfigurationTimeUnitPb string

const PeriodicTriggerConfigurationTimeUnitDays PeriodicTriggerConfigurationTimeUnitPb = `DAYS`

const PeriodicTriggerConfigurationTimeUnitHours PeriodicTriggerConfigurationTimeUnitPb = `HOURS`

const PeriodicTriggerConfigurationTimeUnitWeeks PeriodicTriggerConfigurationTimeUnitPb = `WEEKS`

type PipelineParamsPb struct {
	FullRefresh bool `json:"full_refresh,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineTaskPb struct {
	FullRefresh bool   `json:"full_refresh,omitempty"`
	PipelineId  string `json:"pipeline_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PowerBiModelPb struct {
	AuthenticationMethod AuthenticationMethodPb `json:"authentication_method,omitempty"`
	ModelName            string                 `json:"model_name,omitempty"`
	OverwriteExisting    bool                   `json:"overwrite_existing,omitempty"`
	StorageMode          StorageModePb          `json:"storage_mode,omitempty"`
	WorkspaceName        string                 `json:"workspace_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PowerBiModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PowerBiModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PowerBiTablePb struct {
	Catalog     string        `json:"catalog,omitempty"`
	Name        string        `json:"name,omitempty"`
	Schema      string        `json:"schema,omitempty"`
	StorageMode StorageModePb `json:"storage_mode,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PowerBiTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PowerBiTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PowerBiTaskPb struct {
	ConnectionResourceName string           `json:"connection_resource_name,omitempty"`
	PowerBiModel           *PowerBiModelPb  `json:"power_bi_model,omitempty"`
	RefreshAfterUpdate     bool             `json:"refresh_after_update,omitempty"`
	Tables                 []PowerBiTablePb `json:"tables,omitempty"`
	WarehouseId            string           `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PowerBiTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PowerBiTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PythonWheelTaskPb struct {
	EntryPoint      string            `json:"entry_point"`
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	PackageName     string            `json:"package_name"`
	Parameters      []string          `json:"parameters,omitempty"`
}

type QueueDetailsPb struct {
	Code    QueueDetailsCodeCodePb `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueueDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueueDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueueDetailsCodeCodePb string

// The run was queued due to reaching the workspace limit of active task runs.
const QueueDetailsCodeCodeActiveRunsLimitReached QueueDetailsCodeCodePb = `ACTIVE_RUNS_LIMIT_REACHED`

// The run was queued due to reaching the workspace limit of active run job
// tasks.
const QueueDetailsCodeCodeActiveRunJobTasksLimitReached QueueDetailsCodeCodePb = `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`

// The run was queued due to reaching the per-job limit of concurrent job runs.
const QueueDetailsCodeCodeMaxConcurrentRunsReached QueueDetailsCodeCodePb = `MAX_CONCURRENT_RUNS_REACHED`

type QueueSettingsPb struct {
	Enabled bool `json:"enabled"`
}

type RepairHistoryItemPb struct {
	EffectivePerformanceTarget PerformanceTargetPb     `json:"effective_performance_target,omitempty"`
	EndTime                    int64                   `json:"end_time,omitempty"`
	Id                         int64                   `json:"id,omitempty"`
	StartTime                  int64                   `json:"start_time,omitempty"`
	State                      *RunStatePb             `json:"state,omitempty"`
	Status                     *RunStatusPb            `json:"status,omitempty"`
	TaskRunIds                 []int64                 `json:"task_run_ids,omitempty"`
	Type                       RepairHistoryItemTypePb `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepairHistoryItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepairHistoryItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepairHistoryItemTypePb string

const RepairHistoryItemTypeOriginal RepairHistoryItemTypePb = `ORIGINAL`

const RepairHistoryItemTypeRepair RepairHistoryItemTypePb = `REPAIR`

type RepairRunPb struct {
	DbtCommands         []string            `json:"dbt_commands,omitempty"`
	JarParams           []string            `json:"jar_params,omitempty"`
	JobParameters       map[string]string   `json:"job_parameters,omitempty"`
	LatestRepairId      int64               `json:"latest_repair_id,omitempty"`
	NotebookParams      map[string]string   `json:"notebook_params,omitempty"`
	PerformanceTarget   PerformanceTargetPb `json:"performance_target,omitempty"`
	PipelineParams      *PipelineParamsPb   `json:"pipeline_params,omitempty"`
	PythonNamedParams   map[string]string   `json:"python_named_params,omitempty"`
	PythonParams        []string            `json:"python_params,omitempty"`
	RerunAllFailedTasks bool                `json:"rerun_all_failed_tasks,omitempty"`
	RerunDependentTasks bool                `json:"rerun_dependent_tasks,omitempty"`
	RerunTasks          []string            `json:"rerun_tasks,omitempty"`
	RunId               int64               `json:"run_id"`
	SparkSubmitParams   []string            `json:"spark_submit_params,omitempty"`
	SqlParams           map[string]string   `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepairRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepairRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RepairRunResponsePb struct {
	RepairId int64 `json:"repair_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RepairRunResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RepairRunResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResetJobPb struct {
	JobId       int64         `json:"job_id"`
	NewSettings JobSettingsPb `json:"new_settings"`
}

type ResolvedConditionTaskValuesPb struct {
	Left  string `json:"left,omitempty"`
	Right string `json:"right,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResolvedConditionTaskValuesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResolvedConditionTaskValuesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResolvedDbtTaskValuesPb struct {
	Commands []string `json:"commands,omitempty"`
}

type ResolvedNotebookTaskValuesPb struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

type ResolvedParamPairValuesPb struct {
	Parameters map[string]string `json:"parameters,omitempty"`
}

type ResolvedPythonWheelTaskValuesPb struct {
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	Parameters      []string          `json:"parameters,omitempty"`
}

type ResolvedRunJobTaskValuesPb struct {
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	Parameters    map[string]string `json:"parameters,omitempty"`
}

type ResolvedStringParamsValuesPb struct {
	Parameters []string `json:"parameters,omitempty"`
}

type ResolvedValuesPb struct {
	ConditionTask   *ResolvedConditionTaskValuesPb   `json:"condition_task,omitempty"`
	DbtTask         *ResolvedDbtTaskValuesPb         `json:"dbt_task,omitempty"`
	NotebookTask    *ResolvedNotebookTaskValuesPb    `json:"notebook_task,omitempty"`
	PythonWheelTask *ResolvedPythonWheelTaskValuesPb `json:"python_wheel_task,omitempty"`
	RunJobTask      *ResolvedRunJobTaskValuesPb      `json:"run_job_task,omitempty"`
	SimulationTask  *ResolvedParamPairValuesPb       `json:"simulation_task,omitempty"`
	SparkJarTask    *ResolvedStringParamsValuesPb    `json:"spark_jar_task,omitempty"`
	SparkPythonTask *ResolvedStringParamsValuesPb    `json:"spark_python_task,omitempty"`
	SparkSubmitTask *ResolvedStringParamsValuesPb    `json:"spark_submit_task,omitempty"`
	SqlTask         *ResolvedParamPairValuesPb       `json:"sql_task,omitempty"`
}

type RunPb struct {
	AttemptNumber              int                   `json:"attempt_number,omitempty"`
	CleanupDuration            int64                 `json:"cleanup_duration,omitempty"`
	ClusterInstance            *ClusterInstancePb    `json:"cluster_instance,omitempty"`
	ClusterSpec                *ClusterSpecPb        `json:"cluster_spec,omitempty"`
	CreatorUserName            string                `json:"creator_user_name,omitempty"`
	Description                string                `json:"description,omitempty"`
	EffectivePerformanceTarget PerformanceTargetPb   `json:"effective_performance_target,omitempty"`
	EndTime                    int64                 `json:"end_time,omitempty"`
	ExecutionDuration          int64                 `json:"execution_duration,omitempty"`
	GitSource                  *GitSourcePb          `json:"git_source,omitempty"`
	HasMore                    bool                  `json:"has_more,omitempty"`
	Iterations                 []RunTaskPb           `json:"iterations,omitempty"`
	JobClusters                []JobClusterPb        `json:"job_clusters,omitempty"`
	JobId                      int64                 `json:"job_id,omitempty"`
	JobParameters              []JobParameterPb      `json:"job_parameters,omitempty"`
	JobRunId                   int64                 `json:"job_run_id,omitempty"`
	NextPageToken              string                `json:"next_page_token,omitempty"`
	NumberInJob                int64                 `json:"number_in_job,omitempty"`
	OriginalAttemptRunId       int64                 `json:"original_attempt_run_id,omitempty"`
	OverridingParameters       *RunParametersPb      `json:"overriding_parameters,omitempty"`
	QueueDuration              int64                 `json:"queue_duration,omitempty"`
	RepairHistory              []RepairHistoryItemPb `json:"repair_history,omitempty"`
	RunDuration                int64                 `json:"run_duration,omitempty"`
	RunId                      int64                 `json:"run_id,omitempty"`
	RunName                    string                `json:"run_name,omitempty"`
	RunPageUrl                 string                `json:"run_page_url,omitempty"`
	RunType                    RunTypePb             `json:"run_type,omitempty"`
	Schedule                   *CronSchedulePb       `json:"schedule,omitempty"`
	SetupDuration              int64                 `json:"setup_duration,omitempty"`
	StartTime                  int64                 `json:"start_time,omitempty"`
	State                      *RunStatePb           `json:"state,omitempty"`
	Status                     *RunStatusPb          `json:"status,omitempty"`
	Tasks                      []RunTaskPb           `json:"tasks,omitempty"`
	Trigger                    TriggerTypePb         `json:"trigger,omitempty"`
	TriggerInfo                *TriggerInfoPb        `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunConditionTaskPb struct {
	Left    string            `json:"left"`
	Op      ConditionTaskOpPb `json:"op"`
	Outcome string            `json:"outcome,omitempty"`
	Right   string            `json:"right"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunConditionTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunConditionTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunForEachTaskPb struct {
	Concurrency int             `json:"concurrency,omitempty"`
	Inputs      string          `json:"inputs"`
	Stats       *ForEachStatsPb `json:"stats,omitempty"`
	Task        TaskPb          `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunForEachTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunForEachTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunIfPb string

// All dependencies have been completed
const RunIfAllDone RunIfPb = `ALL_DONE`

// ALl dependencies have failed
const RunIfAllFailed RunIfPb = `ALL_FAILED`

// All dependencies have executed and succeeded
const RunIfAllSuccess RunIfPb = `ALL_SUCCESS`

// At least one dependency failed
const RunIfAtLeastOneFailed RunIfPb = `AT_LEAST_ONE_FAILED`

// At least one dependency has succeeded
const RunIfAtLeastOneSuccess RunIfPb = `AT_LEAST_ONE_SUCCESS`

// None of the dependencies have failed and at least one was executed
const RunIfNoneFailed RunIfPb = `NONE_FAILED`

type RunJobOutputPb struct {
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunJobOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunJobOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunJobTaskPb struct {
	DbtCommands       []string          `json:"dbt_commands,omitempty"`
	JarParams         []string          `json:"jar_params,omitempty"`
	JobId             int64             `json:"job_id"`
	JobParameters     map[string]string `json:"job_parameters,omitempty"`
	NotebookParams    map[string]string `json:"notebook_params,omitempty"`
	PipelineParams    *PipelineParamsPb `json:"pipeline_params,omitempty"`
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	PythonParams      []string          `json:"python_params,omitempty"`
	SparkSubmitParams []string          `json:"spark_submit_params,omitempty"`
	SqlParams         map[string]string `json:"sql_params,omitempty"`
}

type RunLifeCycleStatePb string

// The run is blocked on an upstream dependency.
const RunLifeCycleStateBlocked RunLifeCycleStatePb = `BLOCKED`

// An exceptional state that indicates a failure in the Jobs service, such as
// network failure over a long period. If a run on a new cluster ends in the
// `INTERNAL_ERROR` state, the Jobs service terminates the cluster as soon as
// possible. This state is terminal.
const RunLifeCycleStateInternalError RunLifeCycleStatePb = `INTERNAL_ERROR`

// The run is waiting to be executed while the cluster and execution context are
// being prepared.
const RunLifeCycleStatePending RunLifeCycleStatePb = `PENDING`

// The run is queued.
const RunLifeCycleStateQueued RunLifeCycleStatePb = `QUEUED`

// The task of this run is being executed.
const RunLifeCycleStateRunning RunLifeCycleStatePb = `RUNNING`

// This run was aborted because a previous run of the same job was already
// active. This state is terminal.
const RunLifeCycleStateSkipped RunLifeCycleStatePb = `SKIPPED`

// The task of this run has completed, and the cluster and execution context
// have been cleaned up. This state is terminal.
const RunLifeCycleStateTerminated RunLifeCycleStatePb = `TERMINATED`

// The task of this run has completed, and the cluster and execution context are
// being cleaned up.
const RunLifeCycleStateTerminating RunLifeCycleStatePb = `TERMINATING`

// The run is waiting for a retry.
const RunLifeCycleStateWaitingForRetry RunLifeCycleStatePb = `WAITING_FOR_RETRY`

type RunLifecycleStateV2StatePb string

const RunLifecycleStateV2StateBlocked RunLifecycleStateV2StatePb = `BLOCKED`

const RunLifecycleStateV2StatePending RunLifecycleStateV2StatePb = `PENDING`

const RunLifecycleStateV2StateQueued RunLifecycleStateV2StatePb = `QUEUED`

const RunLifecycleStateV2StateRunning RunLifecycleStateV2StatePb = `RUNNING`

const RunLifecycleStateV2StateTerminated RunLifecycleStateV2StatePb = `TERMINATED`

const RunLifecycleStateV2StateTerminating RunLifecycleStateV2StatePb = `TERMINATING`

const RunLifecycleStateV2StateWaiting RunLifecycleStateV2StatePb = `WAITING`

type RunNowPb struct {
	DbtCommands       []string            `json:"dbt_commands,omitempty"`
	IdempotencyToken  string              `json:"idempotency_token,omitempty"`
	JarParams         []string            `json:"jar_params,omitempty"`
	JobId             int64               `json:"job_id"`
	JobParameters     map[string]string   `json:"job_parameters,omitempty"`
	NotebookParams    map[string]string   `json:"notebook_params,omitempty"`
	Only              []string            `json:"only,omitempty"`
	PerformanceTarget PerformanceTargetPb `json:"performance_target,omitempty"`
	PipelineParams    *PipelineParamsPb   `json:"pipeline_params,omitempty"`
	PythonNamedParams map[string]string   `json:"python_named_params,omitempty"`
	PythonParams      []string            `json:"python_params,omitempty"`
	Queue             *QueueSettingsPb    `json:"queue,omitempty"`
	SparkSubmitParams []string            `json:"spark_submit_params,omitempty"`
	SqlParams         map[string]string   `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunNowPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunNowPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunNowResponsePb struct {
	NumberInJob int64 `json:"number_in_job,omitempty"`
	RunId       int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunNowResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunNowResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunOutputPb struct {
	CleanRoomsNotebookOutput *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb `json:"clean_rooms_notebook_output,omitempty"`
	DashboardOutput          *DashboardTaskOutputPb                                `json:"dashboard_output,omitempty"`
	DbtCloudOutput           *DbtCloudTaskOutputPb                                 `json:"dbt_cloud_output,omitempty"`
	DbtOutput                *DbtOutputPb                                          `json:"dbt_output,omitempty"`
	DbtPlatformOutput        *DbtPlatformTaskOutputPb                              `json:"dbt_platform_output,omitempty"`
	Error                    string                                                `json:"error,omitempty"`
	ErrorTrace               string                                                `json:"error_trace,omitempty"`
	Info                     string                                                `json:"info,omitempty"`
	Logs                     string                                                `json:"logs,omitempty"`
	LogsTruncated            bool                                                  `json:"logs_truncated,omitempty"`
	Metadata                 *RunPb                                                `json:"metadata,omitempty"`
	NotebookOutput           *NotebookOutputPb                                     `json:"notebook_output,omitempty"`
	RunJobOutput             *RunJobOutputPb                                       `json:"run_job_output,omitempty"`
	SqlOutput                *SqlOutputPb                                          `json:"sql_output,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunParametersPb struct {
	DbtCommands       []string          `json:"dbt_commands,omitempty"`
	JarParams         []string          `json:"jar_params,omitempty"`
	NotebookParams    map[string]string `json:"notebook_params,omitempty"`
	PipelineParams    *PipelineParamsPb `json:"pipeline_params,omitempty"`
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	PythonParams      []string          `json:"python_params,omitempty"`
	SparkSubmitParams []string          `json:"spark_submit_params,omitempty"`
	SqlParams         map[string]string `json:"sql_params,omitempty"`
}

type RunResultStatePb string

// The run was canceled at user request.
const RunResultStateCanceled RunResultStatePb = `CANCELED`

// The run was skipped because it was disabled explicitly by the user.
const RunResultStateDisabled RunResultStatePb = `DISABLED`

// The run was skipped because the necessary conditions were not met.
const RunResultStateExcluded RunResultStatePb = `EXCLUDED`

// The task completed with an error.
const RunResultStateFailed RunResultStatePb = `FAILED`

// The run was skipped because the maximum concurrent runs were reached.
const RunResultStateMaximumConcurrentRunsReached RunResultStatePb = `MAXIMUM_CONCURRENT_RUNS_REACHED`

// The task completed successfully.
const RunResultStateSuccess RunResultStatePb = `SUCCESS`

// The job run completed successfully with some failures; leaf tasks were
// successful.
const RunResultStateSuccessWithFailures RunResultStatePb = `SUCCESS_WITH_FAILURES`

// The run was stopped after reaching the timeout.
const RunResultStateTimedout RunResultStatePb = `TIMEDOUT`

// The run was skipped because an upstream task was canceled.
const RunResultStateUpstreamCanceled RunResultStatePb = `UPSTREAM_CANCELED`

// The run was skipped because of an upstream failure.
const RunResultStateUpstreamFailed RunResultStatePb = `UPSTREAM_FAILED`

type RunStatePb struct {
	LifeCycleState          RunLifeCycleStatePb `json:"life_cycle_state,omitempty"`
	QueueReason             string              `json:"queue_reason,omitempty"`
	ResultState             RunResultStatePb    `json:"result_state,omitempty"`
	StateMessage            string              `json:"state_message,omitempty"`
	UserCancelledOrTimedout bool                `json:"user_cancelled_or_timedout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunStatusPb struct {
	QueueDetails       *QueueDetailsPb            `json:"queue_details,omitempty"`
	State              RunLifecycleStateV2StatePb `json:"state,omitempty"`
	TerminationDetails *TerminationDetailsPb      `json:"termination_details,omitempty"`
}

type RunTaskPb struct {
	AttemptNumber              int                         `json:"attempt_number,omitempty"`
	CleanRoomsNotebookTask     *CleanRoomsNotebookTaskPb   `json:"clean_rooms_notebook_task,omitempty"`
	CleanupDuration            int64                       `json:"cleanup_duration,omitempty"`
	ClusterInstance            *ClusterInstancePb          `json:"cluster_instance,omitempty"`
	ConditionTask              *RunConditionTaskPb         `json:"condition_task,omitempty"`
	DashboardTask              *DashboardTaskPb            `json:"dashboard_task,omitempty"`
	DbtCloudTask               *DbtCloudTaskPb             `json:"dbt_cloud_task,omitempty"`
	DbtPlatformTask            *DbtPlatformTaskPb          `json:"dbt_platform_task,omitempty"`
	DbtTask                    *DbtTaskPb                  `json:"dbt_task,omitempty"`
	DependsOn                  []TaskDependencyPb          `json:"depends_on,omitempty"`
	Description                string                      `json:"description,omitempty"`
	Disabled                   bool                        `json:"disabled,omitempty"`
	EffectivePerformanceTarget PerformanceTargetPb         `json:"effective_performance_target,omitempty"`
	EmailNotifications         *JobEmailNotificationsPb    `json:"email_notifications,omitempty"`
	EndTime                    int64                       `json:"end_time,omitempty"`
	EnvironmentKey             string                      `json:"environment_key,omitempty"`
	ExecutionDuration          int64                       `json:"execution_duration,omitempty"`
	ExistingClusterId          string                      `json:"existing_cluster_id,omitempty"`
	ForEachTask                *RunForEachTaskPb           `json:"for_each_task,omitempty"`
	GenAiComputeTask           *GenAiComputeTaskPb         `json:"gen_ai_compute_task,omitempty"`
	GitSource                  *GitSourcePb                `json:"git_source,omitempty"`
	JobClusterKey              string                      `json:"job_cluster_key,omitempty"`
	Libraries                  []computepb.LibraryPb       `json:"libraries,omitempty"`
	NewCluster                 *computepb.ClusterSpecPb    `json:"new_cluster,omitempty"`
	NotebookTask               *NotebookTaskPb             `json:"notebook_task,omitempty"`
	NotificationSettings       *TaskNotificationSettingsPb `json:"notification_settings,omitempty"`
	PipelineTask               *PipelineTaskPb             `json:"pipeline_task,omitempty"`
	PowerBiTask                *PowerBiTaskPb              `json:"power_bi_task,omitempty"`
	PythonWheelTask            *PythonWheelTaskPb          `json:"python_wheel_task,omitempty"`
	QueueDuration              int64                       `json:"queue_duration,omitempty"`
	ResolvedValues             *ResolvedValuesPb           `json:"resolved_values,omitempty"`
	RunDuration                int64                       `json:"run_duration,omitempty"`
	RunId                      int64                       `json:"run_id,omitempty"`
	RunIf                      RunIfPb                     `json:"run_if,omitempty"`
	RunJobTask                 *RunJobTaskPb               `json:"run_job_task,omitempty"`
	RunPageUrl                 string                      `json:"run_page_url,omitempty"`
	SetupDuration              int64                       `json:"setup_duration,omitempty"`
	SparkJarTask               *SparkJarTaskPb             `json:"spark_jar_task,omitempty"`
	SparkPythonTask            *SparkPythonTaskPb          `json:"spark_python_task,omitempty"`
	SparkSubmitTask            *SparkSubmitTaskPb          `json:"spark_submit_task,omitempty"`
	SqlTask                    *SqlTaskPb                  `json:"sql_task,omitempty"`
	StartTime                  int64                       `json:"start_time,omitempty"`
	State                      *RunStatePb                 `json:"state,omitempty"`
	Status                     *RunStatusPb                `json:"status,omitempty"`
	TaskKey                    string                      `json:"task_key"`
	TimeoutSeconds             int                         `json:"timeout_seconds,omitempty"`
	WebhookNotifications       *WebhookNotificationsPb     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunTypePb string

// Normal job run. A run created with :method:jobs/runNow.
const RunTypeJobRun RunTypePb = `JOB_RUN`

// Submit run. A run created with :method:jobs/submit.
const RunTypeSubmitRun RunTypePb = `SUBMIT_RUN`

// Workflow run. A run created with [dbutils.notebook.run].
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
const RunTypeWorkflowRun RunTypePb = `WORKFLOW_RUN`

type SourcePb string

// SQL file is located in cloud Git provider.
const SourceGit SourcePb = `GIT`

// SQL file is located in <Databricks> workspace.
const SourceWorkspace SourcePb = `WORKSPACE`

type SparkJarTaskPb struct {
	JarUri        string   `json:"jar_uri,omitempty"`
	MainClassName string   `json:"main_class_name,omitempty"`
	Parameters    []string `json:"parameters,omitempty"`
	RunAsRepl     bool     `json:"run_as_repl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SparkJarTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SparkJarTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkPythonTaskPb struct {
	Parameters []string `json:"parameters,omitempty"`
	PythonFile string   `json:"python_file"`
	Source     SourcePb `json:"source,omitempty"`
}

type SparkSubmitTaskPb struct {
	Parameters []string `json:"parameters,omitempty"`
}

type SqlAlertOutputPb struct {
	AlertState    SqlAlertStatePb        `json:"alert_state,omitempty"`
	OutputLink    string                 `json:"output_link,omitempty"`
	QueryText     string                 `json:"query_text,omitempty"`
	SqlStatements []SqlStatementOutputPb `json:"sql_statements,omitempty"`
	WarehouseId   string                 `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlAlertOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlAlertOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlAlertStatePb string

const SqlAlertStateOk SqlAlertStatePb = `OK`

const SqlAlertStateTriggered SqlAlertStatePb = `TRIGGERED`

const SqlAlertStateUnknown SqlAlertStatePb = `UNKNOWN`

type SqlDashboardOutputPb struct {
	WarehouseId string                       `json:"warehouse_id,omitempty"`
	Widgets     []SqlDashboardWidgetOutputPb `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlDashboardOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlDashboardOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlDashboardWidgetOutputPb struct {
	EndTime     int64                            `json:"end_time,omitempty"`
	Error       *SqlOutputErrorPb                `json:"error,omitempty"`
	OutputLink  string                           `json:"output_link,omitempty"`
	StartTime   int64                            `json:"start_time,omitempty"`
	Status      SqlDashboardWidgetOutputStatusPb `json:"status,omitempty"`
	WidgetId    string                           `json:"widget_id,omitempty"`
	WidgetTitle string                           `json:"widget_title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlDashboardWidgetOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlDashboardWidgetOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlDashboardWidgetOutputStatusPb string

const SqlDashboardWidgetOutputStatusCancelled SqlDashboardWidgetOutputStatusPb = `CANCELLED`

const SqlDashboardWidgetOutputStatusFailed SqlDashboardWidgetOutputStatusPb = `FAILED`

const SqlDashboardWidgetOutputStatusPending SqlDashboardWidgetOutputStatusPb = `PENDING`

const SqlDashboardWidgetOutputStatusRunning SqlDashboardWidgetOutputStatusPb = `RUNNING`

const SqlDashboardWidgetOutputStatusSuccess SqlDashboardWidgetOutputStatusPb = `SUCCESS`

type SqlOutputPb struct {
	AlertOutput     *SqlAlertOutputPb     `json:"alert_output,omitempty"`
	DashboardOutput *SqlDashboardOutputPb `json:"dashboard_output,omitempty"`
	QueryOutput     *SqlQueryOutputPb     `json:"query_output,omitempty"`
}

type SqlOutputErrorPb struct {
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlOutputErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlOutputErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlQueryOutputPb struct {
	EndpointId    string                 `json:"endpoint_id,omitempty"`
	OutputLink    string                 `json:"output_link,omitempty"`
	QueryText     string                 `json:"query_text,omitempty"`
	SqlStatements []SqlStatementOutputPb `json:"sql_statements,omitempty"`
	WarehouseId   string                 `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlQueryOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlQueryOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlStatementOutputPb struct {
	LookupKey string `json:"lookup_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlStatementOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlStatementOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlTaskPb struct {
	Alert       *SqlTaskAlertPb     `json:"alert,omitempty"`
	Dashboard   *SqlTaskDashboardPb `json:"dashboard,omitempty"`
	File        *SqlTaskFilePb      `json:"file,omitempty"`
	Parameters  map[string]string   `json:"parameters,omitempty"`
	Query       *SqlTaskQueryPb     `json:"query,omitempty"`
	WarehouseId string              `json:"warehouse_id"`
}

type SqlTaskAlertPb struct {
	AlertId            string                  `json:"alert_id"`
	PauseSubscriptions bool                    `json:"pause_subscriptions,omitempty"`
	Subscriptions      []SqlTaskSubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlTaskAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlTaskAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlTaskDashboardPb struct {
	CustomSubject      string                  `json:"custom_subject,omitempty"`
	DashboardId        string                  `json:"dashboard_id"`
	PauseSubscriptions bool                    `json:"pause_subscriptions,omitempty"`
	Subscriptions      []SqlTaskSubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlTaskDashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlTaskDashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlTaskFilePb struct {
	Path   string   `json:"path"`
	Source SourcePb `json:"source,omitempty"`
}

type SqlTaskQueryPb struct {
	QueryId string `json:"query_id"`
}

type SqlTaskSubscriptionPb struct {
	DestinationId string `json:"destination_id,omitempty"`
	UserName      string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlTaskSubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlTaskSubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StorageModePb string

const StorageModeDirectQuery StorageModePb = `DIRECT_QUERY`

const StorageModeDual StorageModePb = `DUAL`

const StorageModeImport StorageModePb = `IMPORT`

type SubmitRunPb struct {
	AccessControlList    []JobAccessControlRequestPb `json:"access_control_list,omitempty"`
	BudgetPolicyId       string                      `json:"budget_policy_id,omitempty"`
	EmailNotifications   *JobEmailNotificationsPb    `json:"email_notifications,omitempty"`
	Environments         []JobEnvironmentPb          `json:"environments,omitempty"`
	GitSource            *GitSourcePb                `json:"git_source,omitempty"`
	Health               *JobsHealthRulesPb          `json:"health,omitempty"`
	IdempotencyToken     string                      `json:"idempotency_token,omitempty"`
	NotificationSettings *JobNotificationSettingsPb  `json:"notification_settings,omitempty"`
	Queue                *QueueSettingsPb            `json:"queue,omitempty"`
	RunAs                *JobRunAsPb                 `json:"run_as,omitempty"`
	RunName              string                      `json:"run_name,omitempty"`
	Tasks                []SubmitTaskPb              `json:"tasks,omitempty"`
	TimeoutSeconds       int                         `json:"timeout_seconds,omitempty"`
	WebhookNotifications *WebhookNotificationsPb     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SubmitRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SubmitRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubmitRunResponsePb struct {
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SubmitRunResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SubmitRunResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubmitTaskPb struct {
	CleanRoomsNotebookTask *CleanRoomsNotebookTaskPb   `json:"clean_rooms_notebook_task,omitempty"`
	ConditionTask          *ConditionTaskPb            `json:"condition_task,omitempty"`
	DashboardTask          *DashboardTaskPb            `json:"dashboard_task,omitempty"`
	DbtCloudTask           *DbtCloudTaskPb             `json:"dbt_cloud_task,omitempty"`
	DbtPlatformTask        *DbtPlatformTaskPb          `json:"dbt_platform_task,omitempty"`
	DbtTask                *DbtTaskPb                  `json:"dbt_task,omitempty"`
	DependsOn              []TaskDependencyPb          `json:"depends_on,omitempty"`
	Description            string                      `json:"description,omitempty"`
	EmailNotifications     *JobEmailNotificationsPb    `json:"email_notifications,omitempty"`
	EnvironmentKey         string                      `json:"environment_key,omitempty"`
	ExistingClusterId      string                      `json:"existing_cluster_id,omitempty"`
	ForEachTask            *ForEachTaskPb              `json:"for_each_task,omitempty"`
	GenAiComputeTask       *GenAiComputeTaskPb         `json:"gen_ai_compute_task,omitempty"`
	Health                 *JobsHealthRulesPb          `json:"health,omitempty"`
	Libraries              []computepb.LibraryPb       `json:"libraries,omitempty"`
	NewCluster             *computepb.ClusterSpecPb    `json:"new_cluster,omitempty"`
	NotebookTask           *NotebookTaskPb             `json:"notebook_task,omitempty"`
	NotificationSettings   *TaskNotificationSettingsPb `json:"notification_settings,omitempty"`
	PipelineTask           *PipelineTaskPb             `json:"pipeline_task,omitempty"`
	PowerBiTask            *PowerBiTaskPb              `json:"power_bi_task,omitempty"`
	PythonWheelTask        *PythonWheelTaskPb          `json:"python_wheel_task,omitempty"`
	RunIf                  RunIfPb                     `json:"run_if,omitempty"`
	RunJobTask             *RunJobTaskPb               `json:"run_job_task,omitempty"`
	SparkJarTask           *SparkJarTaskPb             `json:"spark_jar_task,omitempty"`
	SparkPythonTask        *SparkPythonTaskPb          `json:"spark_python_task,omitempty"`
	SparkSubmitTask        *SparkSubmitTaskPb          `json:"spark_submit_task,omitempty"`
	SqlTask                *SqlTaskPb                  `json:"sql_task,omitempty"`
	TaskKey                string                      `json:"task_key"`
	TimeoutSeconds         int                         `json:"timeout_seconds,omitempty"`
	WebhookNotifications   *WebhookNotificationsPb     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SubmitTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SubmitTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubscriptionPb struct {
	CustomSubject string                     `json:"custom_subject,omitempty"`
	Paused        bool                       `json:"paused,omitempty"`
	Subscribers   []SubscriptionSubscriberPb `json:"subscribers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubscriptionSubscriberPb struct {
	DestinationId string `json:"destination_id,omitempty"`
	UserName      string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SubscriptionSubscriberPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SubscriptionSubscriberPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableUpdateTriggerConfigurationPb struct {
	Condition                     ConditionPb `json:"condition,omitempty"`
	MinTimeBetweenTriggersSeconds int         `json:"min_time_between_triggers_seconds,omitempty"`
	TableNames                    []string    `json:"table_names,omitempty"`
	WaitAfterLastChangeSeconds    int         `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableUpdateTriggerConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableUpdateTriggerConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskPb struct {
	CleanRoomsNotebookTask  *CleanRoomsNotebookTaskPb   `json:"clean_rooms_notebook_task,omitempty"`
	ConditionTask           *ConditionTaskPb            `json:"condition_task,omitempty"`
	DashboardTask           *DashboardTaskPb            `json:"dashboard_task,omitempty"`
	DbtCloudTask            *DbtCloudTaskPb             `json:"dbt_cloud_task,omitempty"`
	DbtPlatformTask         *DbtPlatformTaskPb          `json:"dbt_platform_task,omitempty"`
	DbtTask                 *DbtTaskPb                  `json:"dbt_task,omitempty"`
	DependsOn               []TaskDependencyPb          `json:"depends_on,omitempty"`
	Description             string                      `json:"description,omitempty"`
	DisableAutoOptimization bool                        `json:"disable_auto_optimization,omitempty"`
	EmailNotifications      *TaskEmailNotificationsPb   `json:"email_notifications,omitempty"`
	EnvironmentKey          string                      `json:"environment_key,omitempty"`
	ExistingClusterId       string                      `json:"existing_cluster_id,omitempty"`
	ForEachTask             *ForEachTaskPb              `json:"for_each_task,omitempty"`
	GenAiComputeTask        *GenAiComputeTaskPb         `json:"gen_ai_compute_task,omitempty"`
	Health                  *JobsHealthRulesPb          `json:"health,omitempty"`
	JobClusterKey           string                      `json:"job_cluster_key,omitempty"`
	Libraries               []computepb.LibraryPb       `json:"libraries,omitempty"`
	MaxRetries              int                         `json:"max_retries,omitempty"`
	MinRetryIntervalMillis  int                         `json:"min_retry_interval_millis,omitempty"`
	NewCluster              *computepb.ClusterSpecPb    `json:"new_cluster,omitempty"`
	NotebookTask            *NotebookTaskPb             `json:"notebook_task,omitempty"`
	NotificationSettings    *TaskNotificationSettingsPb `json:"notification_settings,omitempty"`
	PipelineTask            *PipelineTaskPb             `json:"pipeline_task,omitempty"`
	PowerBiTask             *PowerBiTaskPb              `json:"power_bi_task,omitempty"`
	PythonWheelTask         *PythonWheelTaskPb          `json:"python_wheel_task,omitempty"`
	RetryOnTimeout          bool                        `json:"retry_on_timeout,omitempty"`
	RunIf                   RunIfPb                     `json:"run_if,omitempty"`
	RunJobTask              *RunJobTaskPb               `json:"run_job_task,omitempty"`
	SparkJarTask            *SparkJarTaskPb             `json:"spark_jar_task,omitempty"`
	SparkPythonTask         *SparkPythonTaskPb          `json:"spark_python_task,omitempty"`
	SparkSubmitTask         *SparkSubmitTaskPb          `json:"spark_submit_task,omitempty"`
	SqlTask                 *SqlTaskPb                  `json:"sql_task,omitempty"`
	TaskKey                 string                      `json:"task_key"`
	TimeoutSeconds          int                         `json:"timeout_seconds,omitempty"`
	WebhookNotifications    *WebhookNotificationsPb     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskDependencyPb struct {
	Outcome string `json:"outcome,omitempty"`
	TaskKey string `json:"task_key"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TaskDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TaskDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskEmailNotificationsPb struct {
	NoAlertForSkippedRuns              bool     `json:"no_alert_for_skipped_runs,omitempty"`
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	OnFailure                          []string `json:"on_failure,omitempty"`
	OnStart                            []string `json:"on_start,omitempty"`
	OnStreamingBacklogExceeded         []string `json:"on_streaming_backlog_exceeded,omitempty"`
	OnSuccess                          []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TaskEmailNotificationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TaskEmailNotificationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskNotificationSettingsPb struct {
	AlertOnLastAttempt     bool `json:"alert_on_last_attempt,omitempty"`
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	NoAlertForSkippedRuns  bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TaskNotificationSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TaskNotificationSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TerminationCodeCodePb string

const TerminationCodeCodeBudgetPolicyLimitExceeded TerminationCodeCodePb = `BUDGET_POLICY_LIMIT_EXCEEDED`

// The run was canceled during execution by the <Databricks> platform; for
// example, if the maximum run duration was exceeded.
const TerminationCodeCodeCanceled TerminationCodeCodePb = `CANCELED`

// The run failed due to a cloud provider issue. Refer to the state message for
// further details.
const TerminationCodeCodeCloudFailure TerminationCodeCodePb = `CLOUD_FAILURE`

// The run failed due to a cluster error. Refer to the state message for further
// details.
const TerminationCodeCodeClusterError TerminationCodeCodePb = `CLUSTER_ERROR`

// The number of cluster creation, start, and upsize requests have exceeded the
// allotted rate limit. Consider spreading the run execution over a larger time
// frame.
const TerminationCodeCodeClusterRequestLimitExceeded TerminationCodeCodePb = `CLUSTER_REQUEST_LIMIT_EXCEEDED`

// The run was never executed because it was disabled explicitly by the user.
const TerminationCodeCodeDisabled TerminationCodeCodePb = `DISABLED`

// The run encountered an error while communicating with the Spark Driver.
const TerminationCodeCodeDriverError TerminationCodeCodePb = `DRIVER_ERROR`

// The run failed because it tried to access a feature unavailable for the
// workspace.
const TerminationCodeCodeFeatureDisabled TerminationCodeCodePb = `FEATURE_DISABLED`

// The run encountered an unexpected error. Refer to the state message for
// further details.
const TerminationCodeCodeInternalError TerminationCodeCodePb = `INTERNAL_ERROR`

// The run failed because it issued an invalid request to start the cluster.
const TerminationCodeCodeInvalidClusterRequest TerminationCodeCodePb = `INVALID_CLUSTER_REQUEST`

// The run failed due to an invalid configuration. Refer to the state message
// for further details.
const TerminationCodeCodeInvalidRunConfiguration TerminationCodeCodePb = `INVALID_RUN_CONFIGURATION`

// The run failed while installing the user-requested library. Refer to the
// state message for further details. The causes might include, but are not
// limited to: The provided library is invalid, there are insufficient
// permissions to install the library, and so forth.
const TerminationCodeCodeLibraryInstallationError TerminationCodeCodePb = `LIBRARY_INSTALLATION_ERROR`

// The scheduled run exceeds the limit of maximum concurrent runs set for the
// job.
const TerminationCodeCodeMaxConcurrentRunsExceeded TerminationCodeCodePb = `MAX_CONCURRENT_RUNS_EXCEEDED`

// The run was skipped due to reaching the job level queue size limit.
const TerminationCodeCodeMaxJobQueueSizeExceeded TerminationCodeCodePb = `MAX_JOB_QUEUE_SIZE_EXCEEDED`

// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link].
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
const TerminationCodeCodeMaxSparkContextsExceeded TerminationCodeCodePb = `MAX_SPARK_CONTEXTS_EXCEEDED`

// Failed to complete the checkout due to an error when communicating with the
// third party service.
const TerminationCodeCodeRepositoryCheckoutFailed TerminationCodeCodePb = `REPOSITORY_CHECKOUT_FAILED`

// A resource necessary for run execution does not exist. Refer to the state
// message for further details.
const TerminationCodeCodeResourceNotFound TerminationCodeCodePb = `RESOURCE_NOT_FOUND`

// The run was completed with task failures. For more details, refer to the
// state message or run output.
const TerminationCodeCodeRunExecutionError TerminationCodeCodePb = `RUN_EXECUTION_ERROR`

// Run was never executed, for example, if the upstream task run failed, the
// dependency type condition was not met, or there were no material tasks to
// execute.
const TerminationCodeCodeSkipped TerminationCodeCodePb = `SKIPPED`

// The run failed due to an error when accessing the customer blob storage.
// Refer to the state message for further details.
const TerminationCodeCodeStorageAccessError TerminationCodeCodePb = `STORAGE_ACCESS_ERROR`

// The run was completed successfully.
const TerminationCodeCodeSuccess TerminationCodeCodePb = `SUCCESS`

// The run was completed successfully but some child runs failed.
const TerminationCodeCodeSuccessWithFailures TerminationCodeCodePb = `SUCCESS_WITH_FAILURES`

// The run failed due to a permission issue while accessing a resource. Refer to
// the state message for further details.
const TerminationCodeCodeUnauthorizedError TerminationCodeCodePb = `UNAUTHORIZED_ERROR`

// The run was successfully canceled during execution by a user.
const TerminationCodeCodeUserCanceled TerminationCodeCodePb = `USER_CANCELED`

// The workspace has reached the quota for the maximum number of concurrent
// active runs. Consider scheduling the runs over a larger time frame.
const TerminationCodeCodeWorkspaceRunLimitExceeded TerminationCodeCodePb = `WORKSPACE_RUN_LIMIT_EXCEEDED`

type TerminationDetailsPb struct {
	Code    TerminationCodeCodePb `json:"code,omitempty"`
	Message string                `json:"message,omitempty"`
	Type    TerminationTypeTypePb `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TerminationDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TerminationDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TerminationTypeTypePb string

// The run was terminated because of an error caused by user input or the job
// configuration.
const TerminationTypeTypeClientError TerminationTypeTypePb = `CLIENT_ERROR`

// The run was terminated because of an issue with your cloud provider.
const TerminationTypeTypeCloudFailure TerminationTypeTypePb = `CLOUD_FAILURE`

// An error occurred in the <Databricks> platform. Please look at the [status
// page] or contact support if the issue persists.
//
// [status page]: https://status.databricks.com/
const TerminationTypeTypeInternalError TerminationTypeTypePb = `INTERNAL_ERROR`

// The run terminated without any issues
const TerminationTypeTypeSuccess TerminationTypeTypePb = `SUCCESS`

type TriggerInfoPb struct {
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TriggerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TriggerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TriggerSettingsPb struct {
	FileArrival *FileArrivalTriggerConfigurationPb `json:"file_arrival,omitempty"`
	PauseStatus PauseStatusPb                      `json:"pause_status,omitempty"`
	Periodic    *PeriodicTriggerConfigurationPb    `json:"periodic,omitempty"`
	Table       *TableUpdateTriggerConfigurationPb `json:"table,omitempty"`
	TableUpdate *TableUpdateTriggerConfigurationPb `json:"table_update,omitempty"`
}

type TriggerStateProtoPb struct {
	FileArrival *FileArrivalTriggerStatePb `json:"file_arrival,omitempty"`
}

type TriggerTypePb string

// Indicates a run that is triggered by a continuous job.
const TriggerTypeContinuous TriggerTypePb = `CONTINUOUS`

// Indicates a run created by user to manually restart a continuous job run.
const TriggerTypeContinuousRestart TriggerTypePb = `CONTINUOUS_RESTART`

// Indicates a run that is triggered by a file arrival.
const TriggerTypeFileArrival TriggerTypePb = `FILE_ARRIVAL`

// One time triggers that fire a single run. This occurs you triggered a single
// run on demand through the UI or the API.
const TriggerTypeOneTime TriggerTypePb = `ONE_TIME`

// Schedules that periodically trigger runs, such as a cron scheduler.
const TriggerTypePeriodic TriggerTypePb = `PERIODIC`

// Indicates a run that is triggered as a retry of a previously failed run. This
// occurs when you request to re-run the job in case of failures.
const TriggerTypeRetry TriggerTypePb = `RETRY`

// Indicates a run that is triggered using a Run Job task.
const TriggerTypeRunJobTask TriggerTypePb = `RUN_JOB_TASK`

// Indicates a run that is triggered by a table update.
const TriggerTypeTable TriggerTypePb = `TABLE`

type UpdateJobPb struct {
	FieldsToRemove []string       `json:"fields_to_remove,omitempty"`
	JobId          int64          `json:"job_id"`
	NewSettings    *JobSettingsPb `json:"new_settings,omitempty"`
}

type ViewItemPb struct {
	Content string     `json:"content,omitempty"`
	Name    string     `json:"name,omitempty"`
	Type    ViewTypePb `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ViewItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ViewItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ViewTypePb string

// Dashboard view item.
const ViewTypeDashboard ViewTypePb = `DASHBOARD`

// Notebook view item.
const ViewTypeNotebook ViewTypePb = `NOTEBOOK`

type ViewsToExportPb string

// All views of the notebook.
const ViewsToExportAll ViewsToExportPb = `ALL`

// Code view of the notebook.
const ViewsToExportCode ViewsToExportPb = `CODE`

// All dashboard views of the notebook.
const ViewsToExportDashboards ViewsToExportPb = `DASHBOARDS`

type WebhookPb struct {
	Id string `json:"id"`
}

type WebhookNotificationsPb struct {
	OnDurationWarningThresholdExceeded []WebhookPb `json:"on_duration_warning_threshold_exceeded,omitempty"`
	OnFailure                          []WebhookPb `json:"on_failure,omitempty"`
	OnStart                            []WebhookPb `json:"on_start,omitempty"`
	OnStreamingBacklogExceeded         []WebhookPb `json:"on_streaming_backlog_exceeded,omitempty"`
	OnSuccess                          []WebhookPb `json:"on_success,omitempty"`
}

type WidgetErrorDetailPb struct {
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WidgetErrorDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WidgetErrorDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
