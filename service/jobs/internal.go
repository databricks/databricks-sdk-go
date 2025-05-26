// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func baseJobToPb(st *BaseJob) (*baseJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &baseJobPb{}
	pb.CreatedTime = st.CreatedTime
	pb.CreatorUserName = st.CreatorUserName
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	pb.HasMore = st.HasMore
	pb.JobId = st.JobId
	pb.Settings = st.Settings

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type baseJobPb struct {
	CreatedTime             int64        `json:"created_time,omitempty"`
	CreatorUserName         string       `json:"creator_user_name,omitempty"`
	EffectiveBudgetPolicyId string       `json:"effective_budget_policy_id,omitempty"`
	HasMore                 bool         `json:"has_more,omitempty"`
	JobId                   int64        `json:"job_id,omitempty"`
	Settings                *JobSettings `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func baseJobFromPb(pb *baseJobPb) (*BaseJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseJob{}
	st.CreatedTime = pb.CreatedTime
	st.CreatorUserName = pb.CreatorUserName
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.HasMore = pb.HasMore
	st.JobId = pb.JobId
	st.Settings = pb.Settings

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *baseJobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st baseJobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func baseRunToPb(st *BaseRun) (*baseRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &baseRunPb{}
	pb.AttemptNumber = st.AttemptNumber
	pb.CleanupDuration = st.CleanupDuration
	pb.ClusterInstance = st.ClusterInstance
	pb.ClusterSpec = st.ClusterSpec
	pb.CreatorUserName = st.CreatorUserName
	pb.Description = st.Description
	pb.EffectivePerformanceTarget = st.EffectivePerformanceTarget
	pb.EndTime = st.EndTime
	pb.ExecutionDuration = st.ExecutionDuration
	pb.GitSource = st.GitSource
	pb.HasMore = st.HasMore
	pb.JobClusters = st.JobClusters
	pb.JobId = st.JobId
	pb.JobParameters = st.JobParameters
	pb.JobRunId = st.JobRunId
	pb.NumberInJob = st.NumberInJob
	pb.OriginalAttemptRunId = st.OriginalAttemptRunId
	pb.OverridingParameters = st.OverridingParameters
	pb.QueueDuration = st.QueueDuration
	pb.RepairHistory = st.RepairHistory
	pb.RunDuration = st.RunDuration
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunPageUrl = st.RunPageUrl
	pb.RunType = st.RunType
	pb.Schedule = st.Schedule
	pb.SetupDuration = st.SetupDuration
	pb.StartTime = st.StartTime
	pb.State = st.State
	pb.Status = st.Status
	pb.Tasks = st.Tasks
	pb.Trigger = st.Trigger
	pb.TriggerInfo = st.TriggerInfo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type baseRunPb struct {
	AttemptNumber              int                 `json:"attempt_number,omitempty"`
	CleanupDuration            int64               `json:"cleanup_duration,omitempty"`
	ClusterInstance            *ClusterInstance    `json:"cluster_instance,omitempty"`
	ClusterSpec                *ClusterSpec        `json:"cluster_spec,omitempty"`
	CreatorUserName            string              `json:"creator_user_name,omitempty"`
	Description                string              `json:"description,omitempty"`
	EffectivePerformanceTarget PerformanceTarget   `json:"effective_performance_target,omitempty"`
	EndTime                    int64               `json:"end_time,omitempty"`
	ExecutionDuration          int64               `json:"execution_duration,omitempty"`
	GitSource                  *GitSource          `json:"git_source,omitempty"`
	HasMore                    bool                `json:"has_more,omitempty"`
	JobClusters                []JobCluster        `json:"job_clusters,omitempty"`
	JobId                      int64               `json:"job_id,omitempty"`
	JobParameters              []JobParameter      `json:"job_parameters,omitempty"`
	JobRunId                   int64               `json:"job_run_id,omitempty"`
	NumberInJob                int64               `json:"number_in_job,omitempty"`
	OriginalAttemptRunId       int64               `json:"original_attempt_run_id,omitempty"`
	OverridingParameters       *RunParameters      `json:"overriding_parameters,omitempty"`
	QueueDuration              int64               `json:"queue_duration,omitempty"`
	RepairHistory              []RepairHistoryItem `json:"repair_history,omitempty"`
	RunDuration                int64               `json:"run_duration,omitempty"`
	RunId                      int64               `json:"run_id,omitempty"`
	RunName                    string              `json:"run_name,omitempty"`
	RunPageUrl                 string              `json:"run_page_url,omitempty"`
	RunType                    RunType             `json:"run_type,omitempty"`
	Schedule                   *CronSchedule       `json:"schedule,omitempty"`
	SetupDuration              int64               `json:"setup_duration,omitempty"`
	StartTime                  int64               `json:"start_time,omitempty"`
	State                      *RunState           `json:"state,omitempty"`
	Status                     *RunStatus          `json:"status,omitempty"`
	Tasks                      []RunTask           `json:"tasks,omitempty"`
	Trigger                    TriggerType         `json:"trigger,omitempty"`
	TriggerInfo                *TriggerInfo        `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func baseRunFromPb(pb *baseRunPb) (*BaseRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseRun{}
	st.AttemptNumber = pb.AttemptNumber
	st.CleanupDuration = pb.CleanupDuration
	st.ClusterInstance = pb.ClusterInstance
	st.ClusterSpec = pb.ClusterSpec
	st.CreatorUserName = pb.CreatorUserName
	st.Description = pb.Description
	st.EffectivePerformanceTarget = pb.EffectivePerformanceTarget
	st.EndTime = pb.EndTime
	st.ExecutionDuration = pb.ExecutionDuration
	st.GitSource = pb.GitSource
	st.HasMore = pb.HasMore
	st.JobClusters = pb.JobClusters
	st.JobId = pb.JobId
	st.JobParameters = pb.JobParameters
	st.JobRunId = pb.JobRunId
	st.NumberInJob = pb.NumberInJob
	st.OriginalAttemptRunId = pb.OriginalAttemptRunId
	st.OverridingParameters = pb.OverridingParameters
	st.QueueDuration = pb.QueueDuration
	st.RepairHistory = pb.RepairHistory
	st.RunDuration = pb.RunDuration
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunPageUrl = pb.RunPageUrl
	st.RunType = pb.RunType
	st.Schedule = pb.Schedule
	st.SetupDuration = pb.SetupDuration
	st.StartTime = pb.StartTime
	st.State = pb.State
	st.Status = pb.Status
	st.Tasks = pb.Tasks
	st.Trigger = pb.Trigger
	st.TriggerInfo = pb.TriggerInfo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *baseRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st baseRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cancelAllRunsToPb(st *CancelAllRuns) (*cancelAllRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelAllRunsPb{}
	pb.AllQueuedRuns = st.AllQueuedRuns
	pb.JobId = st.JobId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cancelAllRunsPb struct {
	AllQueuedRuns bool  `json:"all_queued_runs,omitempty"`
	JobId         int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cancelAllRunsFromPb(pb *cancelAllRunsPb) (*CancelAllRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelAllRuns{}
	st.AllQueuedRuns = pb.AllQueuedRuns
	st.JobId = pb.JobId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cancelAllRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cancelAllRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cancelAllRunsResponseToPb(st *CancelAllRunsResponse) (*cancelAllRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelAllRunsResponsePb{}

	return pb, nil
}

type cancelAllRunsResponsePb struct {
}

func cancelAllRunsResponseFromPb(pb *cancelAllRunsResponsePb) (*CancelAllRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelAllRunsResponse{}

	return st, nil
}

func cancelRunToPb(st *CancelRun) (*cancelRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

type cancelRunPb struct {
	RunId int64 `json:"run_id"`
}

func cancelRunFromPb(pb *cancelRunPb) (*CancelRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRun{}
	st.RunId = pb.RunId

	return st, nil
}

func cancelRunResponseToPb(st *CancelRunResponse) (*cancelRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRunResponsePb{}

	return pb, nil
}

type cancelRunResponsePb struct {
}

func cancelRunResponseFromPb(pb *cancelRunResponsePb) (*CancelRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRunResponse{}

	return st, nil
}

func cleanRoomTaskRunStateToPb(st *CleanRoomTaskRunState) (*cleanRoomTaskRunStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomTaskRunStatePb{}
	pb.LifeCycleState = st.LifeCycleState
	pb.ResultState = st.ResultState

	return pb, nil
}

type cleanRoomTaskRunStatePb struct {
	LifeCycleState CleanRoomTaskRunLifeCycleState `json:"life_cycle_state,omitempty"`
	ResultState    CleanRoomTaskRunResultState    `json:"result_state,omitempty"`
}

func cleanRoomTaskRunStateFromPb(pb *cleanRoomTaskRunStatePb) (*CleanRoomTaskRunState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomTaskRunState{}
	st.LifeCycleState = pb.LifeCycleState
	st.ResultState = pb.ResultState

	return st, nil
}

func cleanRoomsNotebookTaskToPb(st *CleanRoomsNotebookTask) (*cleanRoomsNotebookTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomsNotebookTaskPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.Etag = st.Etag
	pb.NotebookBaseParameters = st.NotebookBaseParameters
	pb.NotebookName = st.NotebookName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cleanRoomsNotebookTaskPb struct {
	CleanRoomName          string            `json:"clean_room_name"`
	Etag                   string            `json:"etag,omitempty"`
	NotebookBaseParameters map[string]string `json:"notebook_base_parameters,omitempty"`
	NotebookName           string            `json:"notebook_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomsNotebookTaskFromPb(pb *cleanRoomsNotebookTaskPb) (*CleanRoomsNotebookTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomsNotebookTask{}
	st.CleanRoomName = pb.CleanRoomName
	st.Etag = pb.Etag
	st.NotebookBaseParameters = pb.NotebookBaseParameters
	st.NotebookName = pb.NotebookName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomsNotebookTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomsNotebookTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(st *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) (*cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb{}
	pb.CleanRoomJobRunState = st.CleanRoomJobRunState
	pb.NotebookOutput = st.NotebookOutput
	pb.OutputSchemaInfo = st.OutputSchemaInfo

	return pb, nil
}

type cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb struct {
	CleanRoomJobRunState *CleanRoomTaskRunState `json:"clean_room_job_run_state,omitempty"`
	NotebookOutput       *NotebookOutput        `json:"notebook_output,omitempty"`
	OutputSchemaInfo     *OutputSchemaInfo      `json:"output_schema_info,omitempty"`
}

func cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb *cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb) (*CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput{}
	st.CleanRoomJobRunState = pb.CleanRoomJobRunState
	st.NotebookOutput = pb.NotebookOutput
	st.OutputSchemaInfo = pb.OutputSchemaInfo

	return st, nil
}

func clusterInstanceToPb(st *ClusterInstance) (*clusterInstancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterInstancePb{}
	pb.ClusterId = st.ClusterId
	pb.SparkContextId = st.SparkContextId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterInstancePb struct {
	ClusterId      string `json:"cluster_id,omitempty"`
	SparkContextId string `json:"spark_context_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterInstanceFromPb(pb *clusterInstancePb) (*ClusterInstance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterInstance{}
	st.ClusterId = pb.ClusterId
	st.SparkContextId = pb.SparkContextId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterInstancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterInstancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterSpecToPb(st *ClusterSpec) (*clusterSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSpecPb{}
	pb.ExistingClusterId = st.ExistingClusterId
	pb.JobClusterKey = st.JobClusterKey
	pb.Libraries = st.Libraries
	pb.NewCluster = st.NewCluster

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterSpecPb struct {
	ExistingClusterId string               `json:"existing_cluster_id,omitempty"`
	JobClusterKey     string               `json:"job_cluster_key,omitempty"`
	Libraries         []compute.Library    `json:"libraries,omitempty"`
	NewCluster        *compute.ClusterSpec `json:"new_cluster,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterSpecFromPb(pb *clusterSpecPb) (*ClusterSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSpec{}
	st.ExistingClusterId = pb.ExistingClusterId
	st.JobClusterKey = pb.JobClusterKey
	st.Libraries = pb.Libraries
	st.NewCluster = pb.NewCluster

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func computeConfigToPb(st *ComputeConfig) (*computeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computeConfigPb{}
	pb.GpuNodePoolId = st.GpuNodePoolId
	pb.GpuType = st.GpuType
	pb.NumGpus = st.NumGpus

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type computeConfigPb struct {
	GpuNodePoolId string `json:"gpu_node_pool_id,omitempty"`
	GpuType       string `json:"gpu_type,omitempty"`
	NumGpus       int    `json:"num_gpus"`

	ForceSendFields []string `json:"-" url:"-"`
}

func computeConfigFromPb(pb *computeConfigPb) (*ComputeConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComputeConfig{}
	st.GpuNodePoolId = pb.GpuNodePoolId
	st.GpuType = pb.GpuType
	st.NumGpus = pb.NumGpus

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *computeConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st computeConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func conditionTaskToPb(st *ConditionTask) (*conditionTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &conditionTaskPb{}
	pb.Left = st.Left
	pb.Op = st.Op
	pb.Right = st.Right

	return pb, nil
}

type conditionTaskPb struct {
	Left  string          `json:"left"`
	Op    ConditionTaskOp `json:"op"`
	Right string          `json:"right"`
}

func conditionTaskFromPb(pb *conditionTaskPb) (*ConditionTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConditionTask{}
	st.Left = pb.Left
	st.Op = pb.Op
	st.Right = pb.Right

	return st, nil
}

func continuousToPb(st *Continuous) (*continuousPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &continuousPb{}
	pb.PauseStatus = st.PauseStatus

	return pb, nil
}

type continuousPb struct {
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
}

func continuousFromPb(pb *continuousPb) (*Continuous, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Continuous{}
	st.PauseStatus = pb.PauseStatus

	return st, nil
}

func createJobToPb(st *CreateJob) (*createJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createJobPb{}
	pb.AccessControlList = st.AccessControlList
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Continuous = st.Continuous
	pb.Deployment = st.Deployment
	pb.Description = st.Description
	pb.EditMode = st.EditMode
	pb.EmailNotifications = st.EmailNotifications
	pb.Environments = st.Environments
	pb.Format = st.Format
	pb.GitSource = st.GitSource
	pb.Health = st.Health
	pb.JobClusters = st.JobClusters
	pb.MaxConcurrentRuns = st.MaxConcurrentRuns
	pb.Name = st.Name
	pb.NotificationSettings = st.NotificationSettings
	pb.Parameters = st.Parameters
	pb.PerformanceTarget = st.PerformanceTarget
	pb.Queue = st.Queue
	pb.RunAs = st.RunAs
	pb.Schedule = st.Schedule
	pb.Tags = st.Tags
	pb.Tasks = st.Tasks
	pb.TimeoutSeconds = st.TimeoutSeconds
	pb.Trigger = st.Trigger
	pb.WebhookNotifications = st.WebhookNotifications

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createJobPb struct {
	AccessControlList    []JobAccessControlRequest `json:"access_control_list,omitempty"`
	BudgetPolicyId       string                    `json:"budget_policy_id,omitempty"`
	Continuous           *Continuous               `json:"continuous,omitempty"`
	Deployment           *JobDeployment            `json:"deployment,omitempty"`
	Description          string                    `json:"description,omitempty"`
	EditMode             JobEditMode               `json:"edit_mode,omitempty"`
	EmailNotifications   *JobEmailNotifications    `json:"email_notifications,omitempty"`
	Environments         []JobEnvironment          `json:"environments,omitempty"`
	Format               Format                    `json:"format,omitempty"`
	GitSource            *GitSource                `json:"git_source,omitempty"`
	Health               *JobsHealthRules          `json:"health,omitempty"`
	JobClusters          []JobCluster              `json:"job_clusters,omitempty"`
	MaxConcurrentRuns    int                       `json:"max_concurrent_runs,omitempty"`
	Name                 string                    `json:"name,omitempty"`
	NotificationSettings *JobNotificationSettings  `json:"notification_settings,omitempty"`
	Parameters           []JobParameterDefinition  `json:"parameters,omitempty"`
	PerformanceTarget    PerformanceTarget         `json:"performance_target,omitempty"`
	Queue                *QueueSettings            `json:"queue,omitempty"`
	RunAs                *JobRunAs                 `json:"run_as,omitempty"`
	Schedule             *CronSchedule             `json:"schedule,omitempty"`
	Tags                 map[string]string         `json:"tags,omitempty"`
	Tasks                []Task                    `json:"tasks,omitempty"`
	TimeoutSeconds       int                       `json:"timeout_seconds,omitempty"`
	Trigger              *TriggerSettings          `json:"trigger,omitempty"`
	WebhookNotifications *WebhookNotifications     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createJobFromPb(pb *createJobPb) (*CreateJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateJob{}
	st.AccessControlList = pb.AccessControlList
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Continuous = pb.Continuous
	st.Deployment = pb.Deployment
	st.Description = pb.Description
	st.EditMode = pb.EditMode
	st.EmailNotifications = pb.EmailNotifications
	st.Environments = pb.Environments
	st.Format = pb.Format
	st.GitSource = pb.GitSource
	st.Health = pb.Health
	st.JobClusters = pb.JobClusters
	st.MaxConcurrentRuns = pb.MaxConcurrentRuns
	st.Name = pb.Name
	st.NotificationSettings = pb.NotificationSettings
	st.Parameters = pb.Parameters
	st.PerformanceTarget = pb.PerformanceTarget
	st.Queue = pb.Queue
	st.RunAs = pb.RunAs
	st.Schedule = pb.Schedule
	st.Tags = pb.Tags
	st.Tasks = pb.Tasks
	st.TimeoutSeconds = pb.TimeoutSeconds
	st.Trigger = pb.Trigger
	st.WebhookNotifications = pb.WebhookNotifications

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createJobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createJobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}
	pb.JobId = st.JobId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createResponsePb struct {
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	st.JobId = pb.JobId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cronScheduleToPb(st *CronSchedule) (*cronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronSchedulePb{}
	pb.PauseStatus = st.PauseStatus
	pb.QuartzCronExpression = st.QuartzCronExpression
	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

type cronSchedulePb struct {
	PauseStatus          PauseStatus `json:"pause_status,omitempty"`
	QuartzCronExpression string      `json:"quartz_cron_expression"`
	TimezoneId           string      `json:"timezone_id"`
}

func cronScheduleFromPb(pb *cronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	st.PauseStatus = pb.PauseStatus
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

func dashboardPageSnapshotToPb(st *DashboardPageSnapshot) (*dashboardPageSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPageSnapshotPb{}
	pb.PageDisplayName = st.PageDisplayName
	pb.WidgetErrorDetails = st.WidgetErrorDetails

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dashboardPageSnapshotPb struct {
	PageDisplayName    string              `json:"page_display_name,omitempty"`
	WidgetErrorDetails []WidgetErrorDetail `json:"widget_error_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardPageSnapshotFromPb(pb *dashboardPageSnapshotPb) (*DashboardPageSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardPageSnapshot{}
	st.PageDisplayName = pb.PageDisplayName
	st.WidgetErrorDetails = pb.WidgetErrorDetails

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardPageSnapshotPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPageSnapshotPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dashboardTaskToPb(st *DashboardTask) (*dashboardTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardTaskPb{}
	pb.DashboardId = st.DashboardId
	pb.Subscription = st.Subscription
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dashboardTaskPb struct {
	DashboardId  string        `json:"dashboard_id,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	WarehouseId  string        `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardTaskFromPb(pb *dashboardTaskPb) (*DashboardTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardTask{}
	st.DashboardId = pb.DashboardId
	st.Subscription = pb.Subscription
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dashboardTaskOutputToPb(st *DashboardTaskOutput) (*dashboardTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardTaskOutputPb{}
	pb.PageSnapshots = st.PageSnapshots

	return pb, nil
}

type dashboardTaskOutputPb struct {
	PageSnapshots []DashboardPageSnapshot `json:"page_snapshots,omitempty"`
}

func dashboardTaskOutputFromPb(pb *dashboardTaskOutputPb) (*DashboardTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardTaskOutput{}
	st.PageSnapshots = pb.PageSnapshots

	return st, nil
}

func dbtOutputToPb(st *DbtOutput) (*dbtOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dbtOutputPb{}
	pb.ArtifactsHeaders = st.ArtifactsHeaders
	pb.ArtifactsLink = st.ArtifactsLink

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dbtOutputPb struct {
	ArtifactsHeaders map[string]string `json:"artifacts_headers,omitempty"`
	ArtifactsLink    string            `json:"artifacts_link,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dbtOutputFromPb(pb *dbtOutputPb) (*DbtOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtOutput{}
	st.ArtifactsHeaders = pb.ArtifactsHeaders
	st.ArtifactsLink = pb.ArtifactsLink

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dbtOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dbtOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dbtTaskToPb(st *DbtTask) (*dbtTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dbtTaskPb{}
	pb.Catalog = st.Catalog
	pb.Commands = st.Commands
	pb.ProfilesDirectory = st.ProfilesDirectory
	pb.ProjectDirectory = st.ProjectDirectory
	pb.Schema = st.Schema
	pb.Source = st.Source
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dbtTaskPb struct {
	Catalog           string   `json:"catalog,omitempty"`
	Commands          []string `json:"commands"`
	ProfilesDirectory string   `json:"profiles_directory,omitempty"`
	ProjectDirectory  string   `json:"project_directory,omitempty"`
	Schema            string   `json:"schema,omitempty"`
	Source            Source   `json:"source,omitempty"`
	WarehouseId       string   `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dbtTaskFromPb(pb *dbtTaskPb) (*DbtTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtTask{}
	st.Catalog = pb.Catalog
	st.Commands = pb.Commands
	st.ProfilesDirectory = pb.ProfilesDirectory
	st.ProjectDirectory = pb.ProjectDirectory
	st.Schema = pb.Schema
	st.Source = pb.Source
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dbtTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dbtTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteJobToPb(st *DeleteJob) (*deleteJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteJobPb{}
	pb.JobId = st.JobId

	return pb, nil
}

type deleteJobPb struct {
	JobId int64 `json:"job_id"`
}

func deleteJobFromPb(pb *deleteJobPb) (*DeleteJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteJob{}
	st.JobId = pb.JobId

	return st, nil
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
}

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

func deleteRunToPb(st *DeleteRun) (*deleteRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

type deleteRunPb struct {
	RunId int64 `json:"run_id"`
}

func deleteRunFromPb(pb *deleteRunPb) (*DeleteRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRun{}
	st.RunId = pb.RunId

	return st, nil
}

func deleteRunResponseToPb(st *DeleteRunResponse) (*deleteRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunResponsePb{}

	return pb, nil
}

type deleteRunResponsePb struct {
}

func deleteRunResponseFromPb(pb *deleteRunResponsePb) (*DeleteRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunResponse{}

	return st, nil
}

func enforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) (*enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb{}
	pb.Field = st.Field
	pb.NewValue = st.NewValue
	pb.PreviousValue = st.PreviousValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb struct {
	Field         string `json:"field,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(pb *enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) (*EnforcePolicyComplianceForJobResponseJobClusterSettingsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceForJobResponseJobClusterSettingsChange{}
	st.Field = pb.Field
	st.NewValue = pb.NewValue
	st.PreviousValue = pb.PreviousValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enforcePolicyComplianceRequestToPb(st *EnforcePolicyComplianceRequest) (*enforcePolicyComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforcePolicyComplianceRequestPb{}
	pb.JobId = st.JobId
	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enforcePolicyComplianceRequestPb struct {
	JobId        int64 `json:"job_id"`
	ValidateOnly bool  `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforcePolicyComplianceRequestFromPb(pb *enforcePolicyComplianceRequestPb) (*EnforcePolicyComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceRequest{}
	st.JobId = pb.JobId
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforcePolicyComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforcePolicyComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enforcePolicyComplianceResponseToPb(st *EnforcePolicyComplianceResponse) (*enforcePolicyComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforcePolicyComplianceResponsePb{}
	pb.HasChanges = st.HasChanges
	pb.JobClusterChanges = st.JobClusterChanges
	pb.Settings = st.Settings

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enforcePolicyComplianceResponsePb struct {
	HasChanges        bool                                                            `json:"has_changes,omitempty"`
	JobClusterChanges []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange `json:"job_cluster_changes,omitempty"`
	Settings          *JobSettings                                                    `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforcePolicyComplianceResponseFromPb(pb *enforcePolicyComplianceResponsePb) (*EnforcePolicyComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceResponse{}
	st.HasChanges = pb.HasChanges
	st.JobClusterChanges = pb.JobClusterChanges
	st.Settings = pb.Settings

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforcePolicyComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforcePolicyComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func exportRunOutputToPb(st *ExportRunOutput) (*exportRunOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportRunOutputPb{}
	pb.Views = st.Views

	return pb, nil
}

type exportRunOutputPb struct {
	Views []ViewItem `json:"views,omitempty"`
}

func exportRunOutputFromPb(pb *exportRunOutputPb) (*ExportRunOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRunOutput{}
	st.Views = pb.Views

	return st, nil
}

func exportRunRequestToPb(st *ExportRunRequest) (*exportRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportRunRequestPb{}
	pb.RunId = st.RunId
	pb.ViewsToExport = st.ViewsToExport

	return pb, nil
}

type exportRunRequestPb struct {
	RunId         int64         `json:"-" url:"run_id"`
	ViewsToExport ViewsToExport `json:"-" url:"views_to_export,omitempty"`
}

func exportRunRequestFromPb(pb *exportRunRequestPb) (*ExportRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRunRequest{}
	st.RunId = pb.RunId
	st.ViewsToExport = pb.ViewsToExport

	return st, nil
}

func fileArrivalTriggerConfigurationToPb(st *FileArrivalTriggerConfiguration) (*fileArrivalTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileArrivalTriggerConfigurationPb{}
	pb.MinTimeBetweenTriggersSeconds = st.MinTimeBetweenTriggersSeconds
	pb.Url = st.Url
	pb.WaitAfterLastChangeSeconds = st.WaitAfterLastChangeSeconds

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type fileArrivalTriggerConfigurationPb struct {
	MinTimeBetweenTriggersSeconds int    `json:"min_time_between_triggers_seconds,omitempty"`
	Url                           string `json:"url"`
	WaitAfterLastChangeSeconds    int    `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileArrivalTriggerConfigurationFromPb(pb *fileArrivalTriggerConfigurationPb) (*FileArrivalTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileArrivalTriggerConfiguration{}
	st.MinTimeBetweenTriggersSeconds = pb.MinTimeBetweenTriggersSeconds
	st.Url = pb.Url
	st.WaitAfterLastChangeSeconds = pb.WaitAfterLastChangeSeconds

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileArrivalTriggerConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileArrivalTriggerConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func forEachStatsToPb(st *ForEachStats) (*forEachStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachStatsPb{}
	pb.ErrorMessageStats = st.ErrorMessageStats
	pb.TaskRunStats = st.TaskRunStats

	return pb, nil
}

type forEachStatsPb struct {
	ErrorMessageStats []ForEachTaskErrorMessageStats `json:"error_message_stats,omitempty"`
	TaskRunStats      *ForEachTaskTaskRunStats       `json:"task_run_stats,omitempty"`
}

func forEachStatsFromPb(pb *forEachStatsPb) (*ForEachStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachStats{}
	st.ErrorMessageStats = pb.ErrorMessageStats
	st.TaskRunStats = pb.TaskRunStats

	return st, nil
}

func forEachTaskToPb(st *ForEachTask) (*forEachTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachTaskPb{}
	pb.Concurrency = st.Concurrency
	pb.Inputs = st.Inputs
	pb.Task = st.Task

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type forEachTaskPb struct {
	Concurrency int    `json:"concurrency,omitempty"`
	Inputs      string `json:"inputs"`
	Task        Task   `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forEachTaskFromPb(pb *forEachTaskPb) (*ForEachTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTask{}
	st.Concurrency = pb.Concurrency
	st.Inputs = pb.Inputs
	st.Task = pb.Task

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forEachTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forEachTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func forEachTaskErrorMessageStatsToPb(st *ForEachTaskErrorMessageStats) (*forEachTaskErrorMessageStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachTaskErrorMessageStatsPb{}
	pb.Count = st.Count
	pb.ErrorMessage = st.ErrorMessage
	pb.TerminationCategory = st.TerminationCategory

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type forEachTaskErrorMessageStatsPb struct {
	Count               int    `json:"count,omitempty"`
	ErrorMessage        string `json:"error_message,omitempty"`
	TerminationCategory string `json:"termination_category,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forEachTaskErrorMessageStatsFromPb(pb *forEachTaskErrorMessageStatsPb) (*ForEachTaskErrorMessageStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTaskErrorMessageStats{}
	st.Count = pb.Count
	st.ErrorMessage = pb.ErrorMessage
	st.TerminationCategory = pb.TerminationCategory

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forEachTaskErrorMessageStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forEachTaskErrorMessageStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func forEachTaskTaskRunStatsToPb(st *ForEachTaskTaskRunStats) (*forEachTaskTaskRunStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachTaskTaskRunStatsPb{}
	pb.ActiveIterations = st.ActiveIterations
	pb.CompletedIterations = st.CompletedIterations
	pb.FailedIterations = st.FailedIterations
	pb.ScheduledIterations = st.ScheduledIterations
	pb.SucceededIterations = st.SucceededIterations
	pb.TotalIterations = st.TotalIterations

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type forEachTaskTaskRunStatsPb struct {
	ActiveIterations    int `json:"active_iterations,omitempty"`
	CompletedIterations int `json:"completed_iterations,omitempty"`
	FailedIterations    int `json:"failed_iterations,omitempty"`
	ScheduledIterations int `json:"scheduled_iterations,omitempty"`
	SucceededIterations int `json:"succeeded_iterations,omitempty"`
	TotalIterations     int `json:"total_iterations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forEachTaskTaskRunStatsFromPb(pb *forEachTaskTaskRunStatsPb) (*ForEachTaskTaskRunStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTaskTaskRunStats{}
	st.ActiveIterations = pb.ActiveIterations
	st.CompletedIterations = pb.CompletedIterations
	st.FailedIterations = pb.FailedIterations
	st.ScheduledIterations = pb.ScheduledIterations
	st.SucceededIterations = pb.SucceededIterations
	st.TotalIterations = pb.TotalIterations

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forEachTaskTaskRunStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forEachTaskTaskRunStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func genAiComputeTaskToPb(st *GenAiComputeTask) (*genAiComputeTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genAiComputeTaskPb{}
	pb.Command = st.Command
	pb.Compute = st.Compute
	pb.DlRuntimeImage = st.DlRuntimeImage
	pb.MlflowExperimentName = st.MlflowExperimentName
	pb.Source = st.Source
	pb.TrainingScriptPath = st.TrainingScriptPath
	pb.YamlParameters = st.YamlParameters
	pb.YamlParametersFilePath = st.YamlParametersFilePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genAiComputeTaskPb struct {
	Command                string         `json:"command,omitempty"`
	Compute                *ComputeConfig `json:"compute,omitempty"`
	DlRuntimeImage         string         `json:"dl_runtime_image"`
	MlflowExperimentName   string         `json:"mlflow_experiment_name,omitempty"`
	Source                 Source         `json:"source,omitempty"`
	TrainingScriptPath     string         `json:"training_script_path,omitempty"`
	YamlParameters         string         `json:"yaml_parameters,omitempty"`
	YamlParametersFilePath string         `json:"yaml_parameters_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genAiComputeTaskFromPb(pb *genAiComputeTaskPb) (*GenAiComputeTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenAiComputeTask{}
	st.Command = pb.Command
	st.Compute = pb.Compute
	st.DlRuntimeImage = pb.DlRuntimeImage
	st.MlflowExperimentName = pb.MlflowExperimentName
	st.Source = pb.Source
	st.TrainingScriptPath = pb.TrainingScriptPath
	st.YamlParameters = pb.YamlParameters
	st.YamlParametersFilePath = pb.YamlParametersFilePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genAiComputeTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genAiComputeTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getJobPermissionLevelsRequestToPb(st *GetJobPermissionLevelsRequest) (*getJobPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobPermissionLevelsRequestPb{}
	pb.JobId = st.JobId

	return pb, nil
}

type getJobPermissionLevelsRequestPb struct {
	JobId string `json:"-" url:"-"`
}

func getJobPermissionLevelsRequestFromPb(pb *getJobPermissionLevelsRequestPb) (*GetJobPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionLevelsRequest{}
	st.JobId = pb.JobId

	return st, nil
}

func getJobPermissionLevelsResponseToPb(st *GetJobPermissionLevelsResponse) (*getJobPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getJobPermissionLevelsResponsePb struct {
	PermissionLevels []JobPermissionsDescription `json:"permission_levels,omitempty"`
}

func getJobPermissionLevelsResponseFromPb(pb *getJobPermissionLevelsResponsePb) (*GetJobPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getJobPermissionsRequestToPb(st *GetJobPermissionsRequest) (*getJobPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobPermissionsRequestPb{}
	pb.JobId = st.JobId

	return pb, nil
}

type getJobPermissionsRequestPb struct {
	JobId string `json:"-" url:"-"`
}

func getJobPermissionsRequestFromPb(pb *getJobPermissionsRequestPb) (*GetJobPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionsRequest{}
	st.JobId = pb.JobId

	return st, nil
}

func getJobRequestToPb(st *GetJobRequest) (*getJobRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobRequestPb{}
	pb.JobId = st.JobId
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getJobRequestPb struct {
	JobId     int64  `json:"-" url:"job_id"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getJobRequestFromPb(pb *getJobRequestPb) (*GetJobRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobRequest{}
	st.JobId = pb.JobId
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getJobRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getJobRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPolicyComplianceRequestToPb(st *GetPolicyComplianceRequest) (*getPolicyComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPolicyComplianceRequestPb{}
	pb.JobId = st.JobId

	return pb, nil
}

type getPolicyComplianceRequestPb struct {
	JobId int64 `json:"-" url:"job_id"`
}

func getPolicyComplianceRequestFromPb(pb *getPolicyComplianceRequestPb) (*GetPolicyComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyComplianceRequest{}
	st.JobId = pb.JobId

	return st, nil
}

func getPolicyComplianceResponseToPb(st *GetPolicyComplianceResponse) (*getPolicyComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPolicyComplianceResponsePb{}
	pb.IsCompliant = st.IsCompliant
	pb.Violations = st.Violations

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPolicyComplianceResponsePb struct {
	IsCompliant bool              `json:"is_compliant,omitempty"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPolicyComplianceResponseFromPb(pb *getPolicyComplianceResponsePb) (*GetPolicyComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyComplianceResponse{}
	st.IsCompliant = pb.IsCompliant
	st.Violations = pb.Violations

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPolicyComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPolicyComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getRunOutputRequestToPb(st *GetRunOutputRequest) (*getRunOutputRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunOutputRequestPb{}
	pb.RunId = st.RunId

	return pb, nil
}

type getRunOutputRequestPb struct {
	RunId int64 `json:"-" url:"run_id"`
}

func getRunOutputRequestFromPb(pb *getRunOutputRequestPb) (*GetRunOutputRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunOutputRequest{}
	st.RunId = pb.RunId

	return st, nil
}

func getRunRequestToPb(st *GetRunRequest) (*getRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunRequestPb{}
	pb.IncludeHistory = st.IncludeHistory
	pb.IncludeResolvedValues = st.IncludeResolvedValues
	pb.PageToken = st.PageToken
	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getRunRequestPb struct {
	IncludeHistory        bool   `json:"-" url:"include_history,omitempty"`
	IncludeResolvedValues bool   `json:"-" url:"include_resolved_values,omitempty"`
	PageToken             string `json:"-" url:"page_token,omitempty"`
	RunId                 int64  `json:"-" url:"run_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRunRequestFromPb(pb *getRunRequestPb) (*GetRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunRequest{}
	st.IncludeHistory = pb.IncludeHistory
	st.IncludeResolvedValues = pb.IncludeResolvedValues
	st.PageToken = pb.PageToken
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRunRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRunRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gitSnapshotToPb(st *GitSnapshot) (*gitSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gitSnapshotPb{}
	pb.UsedCommit = st.UsedCommit

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gitSnapshotPb struct {
	UsedCommit string `json:"used_commit,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gitSnapshotFromPb(pb *gitSnapshotPb) (*GitSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GitSnapshot{}
	st.UsedCommit = pb.UsedCommit

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gitSnapshotPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gitSnapshotPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gitSourceToPb(st *GitSource) (*gitSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gitSourcePb{}
	pb.GitBranch = st.GitBranch
	pb.GitCommit = st.GitCommit
	pb.GitProvider = st.GitProvider
	pb.GitSnapshot = st.GitSnapshot
	pb.GitTag = st.GitTag
	pb.GitUrl = st.GitUrl
	pb.JobSource = st.JobSource

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gitSourcePb struct {
	GitBranch   string       `json:"git_branch,omitempty"`
	GitCommit   string       `json:"git_commit,omitempty"`
	GitProvider GitProvider  `json:"git_provider"`
	GitSnapshot *GitSnapshot `json:"git_snapshot,omitempty"`
	GitTag      string       `json:"git_tag,omitempty"`
	GitUrl      string       `json:"git_url"`
	JobSource   *JobSource   `json:"job_source,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gitSourceFromPb(pb *gitSourcePb) (*GitSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GitSource{}
	st.GitBranch = pb.GitBranch
	st.GitCommit = pb.GitCommit
	st.GitProvider = pb.GitProvider
	st.GitSnapshot = pb.GitSnapshot
	st.GitTag = pb.GitTag
	st.GitUrl = pb.GitUrl
	st.JobSource = pb.JobSource

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gitSourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gitSourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobToPb(st *Job) (*jobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPb{}
	pb.CreatedTime = st.CreatedTime
	pb.CreatorUserName = st.CreatorUserName
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	pb.HasMore = st.HasMore
	pb.JobId = st.JobId
	pb.NextPageToken = st.NextPageToken
	pb.RunAsUserName = st.RunAsUserName
	pb.Settings = st.Settings

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobPb struct {
	CreatedTime             int64        `json:"created_time,omitempty"`
	CreatorUserName         string       `json:"creator_user_name,omitempty"`
	EffectiveBudgetPolicyId string       `json:"effective_budget_policy_id,omitempty"`
	HasMore                 bool         `json:"has_more,omitempty"`
	JobId                   int64        `json:"job_id,omitempty"`
	NextPageToken           string       `json:"next_page_token,omitempty"`
	RunAsUserName           string       `json:"run_as_user_name,omitempty"`
	Settings                *JobSettings `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobFromPb(pb *jobPb) (*Job, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Job{}
	st.CreatedTime = pb.CreatedTime
	st.CreatorUserName = pb.CreatorUserName
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.HasMore = pb.HasMore
	st.JobId = pb.JobId
	st.NextPageToken = pb.NextPageToken
	st.RunAsUserName = pb.RunAsUserName
	st.Settings = pb.Settings

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobAccessControlRequestToPb(st *JobAccessControlRequest) (*jobAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobAccessControlRequestPb struct {
	GroupName            string             `json:"group_name,omitempty"`
	PermissionLevel      JobPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string             `json:"service_principal_name,omitempty"`
	UserName             string             `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobAccessControlRequestFromPb(pb *jobAccessControlRequestPb) (*JobAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobAccessControlResponseToPb(st *JobAccessControlResponse) (*jobAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobAccessControlResponsePb struct {
	AllPermissions       []JobPermission `json:"all_permissions,omitempty"`
	DisplayName          string          `json:"display_name,omitempty"`
	GroupName            string          `json:"group_name,omitempty"`
	ServicePrincipalName string          `json:"service_principal_name,omitempty"`
	UserName             string          `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobAccessControlResponseFromPb(pb *jobAccessControlResponsePb) (*JobAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobClusterToPb(st *JobCluster) (*jobClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobClusterPb{}
	pb.JobClusterKey = st.JobClusterKey
	pb.NewCluster = st.NewCluster

	return pb, nil
}

type jobClusterPb struct {
	JobClusterKey string              `json:"job_cluster_key"`
	NewCluster    compute.ClusterSpec `json:"new_cluster"`
}

func jobClusterFromPb(pb *jobClusterPb) (*JobCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobCluster{}
	st.JobClusterKey = pb.JobClusterKey
	st.NewCluster = pb.NewCluster

	return st, nil
}

func jobComplianceToPb(st *JobCompliance) (*jobCompliancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobCompliancePb{}
	pb.IsCompliant = st.IsCompliant
	pb.JobId = st.JobId
	pb.Violations = st.Violations

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobCompliancePb struct {
	IsCompliant bool              `json:"is_compliant,omitempty"`
	JobId       int64             `json:"job_id"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobComplianceFromPb(pb *jobCompliancePb) (*JobCompliance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobCompliance{}
	st.IsCompliant = pb.IsCompliant
	st.JobId = pb.JobId
	st.Violations = pb.Violations

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobCompliancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobCompliancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobDeploymentToPb(st *JobDeployment) (*jobDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobDeploymentPb{}
	pb.Kind = st.Kind
	pb.MetadataFilePath = st.MetadataFilePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobDeploymentPb struct {
	Kind             JobDeploymentKind `json:"kind"`
	MetadataFilePath string            `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobDeploymentFromPb(pb *jobDeploymentPb) (*JobDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobDeployment{}
	st.Kind = pb.Kind
	st.MetadataFilePath = pb.MetadataFilePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobEmailNotificationsToPb(st *JobEmailNotifications) (*jobEmailNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobEmailNotificationsPb{}
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns
	pb.OnDurationWarningThresholdExceeded = st.OnDurationWarningThresholdExceeded
	pb.OnFailure = st.OnFailure
	pb.OnStart = st.OnStart
	pb.OnStreamingBacklogExceeded = st.OnStreamingBacklogExceeded
	pb.OnSuccess = st.OnSuccess

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobEmailNotificationsPb struct {
	NoAlertForSkippedRuns              bool     `json:"no_alert_for_skipped_runs,omitempty"`
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	OnFailure                          []string `json:"on_failure,omitempty"`
	OnStart                            []string `json:"on_start,omitempty"`
	OnStreamingBacklogExceeded         []string `json:"on_streaming_backlog_exceeded,omitempty"`
	OnSuccess                          []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobEmailNotificationsFromPb(pb *jobEmailNotificationsPb) (*JobEmailNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobEmailNotifications{}
	st.NoAlertForSkippedRuns = pb.NoAlertForSkippedRuns
	st.OnDurationWarningThresholdExceeded = pb.OnDurationWarningThresholdExceeded
	st.OnFailure = pb.OnFailure
	st.OnStart = pb.OnStart
	st.OnStreamingBacklogExceeded = pb.OnStreamingBacklogExceeded
	st.OnSuccess = pb.OnSuccess

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobEmailNotificationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobEmailNotificationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobEnvironmentToPb(st *JobEnvironment) (*jobEnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobEnvironmentPb{}
	pb.EnvironmentKey = st.EnvironmentKey
	pb.Spec = st.Spec

	return pb, nil
}

type jobEnvironmentPb struct {
	EnvironmentKey string               `json:"environment_key"`
	Spec           *compute.Environment `json:"spec,omitempty"`
}

func jobEnvironmentFromPb(pb *jobEnvironmentPb) (*JobEnvironment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobEnvironment{}
	st.EnvironmentKey = pb.EnvironmentKey
	st.Spec = pb.Spec

	return st, nil
}

func jobNotificationSettingsToPb(st *JobNotificationSettings) (*jobNotificationSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobNotificationSettingsPb{}
	pb.NoAlertForCanceledRuns = st.NoAlertForCanceledRuns
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobNotificationSettingsPb struct {
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	NoAlertForSkippedRuns  bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobNotificationSettingsFromPb(pb *jobNotificationSettingsPb) (*JobNotificationSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobNotificationSettings{}
	st.NoAlertForCanceledRuns = pb.NoAlertForCanceledRuns
	st.NoAlertForSkippedRuns = pb.NoAlertForSkippedRuns

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobNotificationSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobNotificationSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobParameterToPb(st *JobParameter) (*jobParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobParameterPb{}
	pb.Default = st.Default
	pb.Name = st.Name
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobParameterPb struct {
	Default string `json:"default,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobParameterFromPb(pb *jobParameterPb) (*JobParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobParameter{}
	st.Default = pb.Default
	st.Name = pb.Name
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobParameterDefinitionToPb(st *JobParameterDefinition) (*jobParameterDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobParameterDefinitionPb{}
	pb.Default = st.Default
	pb.Name = st.Name

	return pb, nil
}

type jobParameterDefinitionPb struct {
	Default string `json:"default"`
	Name    string `json:"name"`
}

func jobParameterDefinitionFromPb(pb *jobParameterDefinitionPb) (*JobParameterDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobParameterDefinition{}
	st.Default = pb.Default
	st.Name = pb.Name

	return st, nil
}

func jobPermissionToPb(st *JobPermission) (*jobPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobPermissionPb struct {
	Inherited           bool               `json:"inherited,omitempty"`
	InheritedFromObject []string           `json:"inherited_from_object,omitempty"`
	PermissionLevel     JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobPermissionFromPb(pb *jobPermissionPb) (*JobPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobPermissionsToPb(st *JobPermissions) (*jobPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobPermissionsPb struct {
	AccessControlList []JobAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                     `json:"object_id,omitempty"`
	ObjectType        string                     `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobPermissionsFromPb(pb *jobPermissionsPb) (*JobPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobPermissionsDescriptionToPb(st *JobPermissionsDescription) (*jobPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobPermissionsDescriptionPb struct {
	Description     string             `json:"description,omitempty"`
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobPermissionsDescriptionFromPb(pb *jobPermissionsDescriptionPb) (*JobPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobPermissionsRequestToPb(st *JobPermissionsRequest) (*jobPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.JobId = st.JobId

	return pb, nil
}

type jobPermissionsRequestPb struct {
	AccessControlList []JobAccessControlRequest `json:"access_control_list,omitempty"`
	JobId             string                    `json:"-" url:"-"`
}

func jobPermissionsRequestFromPb(pb *jobPermissionsRequestPb) (*JobPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.JobId = pb.JobId

	return st, nil
}

func jobRunAsToPb(st *JobRunAs) (*jobRunAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobRunAsPb{}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobRunAsPb struct {
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobRunAsFromPb(pb *jobRunAsPb) (*JobRunAs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobRunAs{}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobRunAsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobRunAsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobSettingsToPb(st *JobSettings) (*jobSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSettingsPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Continuous = st.Continuous
	pb.Deployment = st.Deployment
	pb.Description = st.Description
	pb.EditMode = st.EditMode
	pb.EmailNotifications = st.EmailNotifications
	pb.Environments = st.Environments
	pb.Format = st.Format
	pb.GitSource = st.GitSource
	pb.Health = st.Health
	pb.JobClusters = st.JobClusters
	pb.MaxConcurrentRuns = st.MaxConcurrentRuns
	pb.Name = st.Name
	pb.NotificationSettings = st.NotificationSettings
	pb.Parameters = st.Parameters
	pb.PerformanceTarget = st.PerformanceTarget
	pb.Queue = st.Queue
	pb.RunAs = st.RunAs
	pb.Schedule = st.Schedule
	pb.Tags = st.Tags
	pb.Tasks = st.Tasks
	pb.TimeoutSeconds = st.TimeoutSeconds
	pb.Trigger = st.Trigger
	pb.WebhookNotifications = st.WebhookNotifications

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type jobSettingsPb struct {
	BudgetPolicyId       string                   `json:"budget_policy_id,omitempty"`
	Continuous           *Continuous              `json:"continuous,omitempty"`
	Deployment           *JobDeployment           `json:"deployment,omitempty"`
	Description          string                   `json:"description,omitempty"`
	EditMode             JobEditMode              `json:"edit_mode,omitempty"`
	EmailNotifications   *JobEmailNotifications   `json:"email_notifications,omitempty"`
	Environments         []JobEnvironment         `json:"environments,omitempty"`
	Format               Format                   `json:"format,omitempty"`
	GitSource            *GitSource               `json:"git_source,omitempty"`
	Health               *JobsHealthRules         `json:"health,omitempty"`
	JobClusters          []JobCluster             `json:"job_clusters,omitempty"`
	MaxConcurrentRuns    int                      `json:"max_concurrent_runs,omitempty"`
	Name                 string                   `json:"name,omitempty"`
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	Parameters           []JobParameterDefinition `json:"parameters,omitempty"`
	PerformanceTarget    PerformanceTarget        `json:"performance_target,omitempty"`
	Queue                *QueueSettings           `json:"queue,omitempty"`
	RunAs                *JobRunAs                `json:"run_as,omitempty"`
	Schedule             *CronSchedule            `json:"schedule,omitempty"`
	Tags                 map[string]string        `json:"tags,omitempty"`
	Tasks                []Task                   `json:"tasks,omitempty"`
	TimeoutSeconds       int                      `json:"timeout_seconds,omitempty"`
	Trigger              *TriggerSettings         `json:"trigger,omitempty"`
	WebhookNotifications *WebhookNotifications    `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobSettingsFromPb(pb *jobSettingsPb) (*JobSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSettings{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Continuous = pb.Continuous
	st.Deployment = pb.Deployment
	st.Description = pb.Description
	st.EditMode = pb.EditMode
	st.EmailNotifications = pb.EmailNotifications
	st.Environments = pb.Environments
	st.Format = pb.Format
	st.GitSource = pb.GitSource
	st.Health = pb.Health
	st.JobClusters = pb.JobClusters
	st.MaxConcurrentRuns = pb.MaxConcurrentRuns
	st.Name = pb.Name
	st.NotificationSettings = pb.NotificationSettings
	st.Parameters = pb.Parameters
	st.PerformanceTarget = pb.PerformanceTarget
	st.Queue = pb.Queue
	st.RunAs = pb.RunAs
	st.Schedule = pb.Schedule
	st.Tags = pb.Tags
	st.Tasks = pb.Tasks
	st.TimeoutSeconds = pb.TimeoutSeconds
	st.Trigger = pb.Trigger
	st.WebhookNotifications = pb.WebhookNotifications

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func jobSourceToPb(st *JobSource) (*jobSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSourcePb{}
	pb.DirtyState = st.DirtyState
	pb.ImportFromGitBranch = st.ImportFromGitBranch
	pb.JobConfigPath = st.JobConfigPath

	return pb, nil
}

type jobSourcePb struct {
	DirtyState          JobSourceDirtyState `json:"dirty_state,omitempty"`
	ImportFromGitBranch string              `json:"import_from_git_branch"`
	JobConfigPath       string              `json:"job_config_path"`
}

func jobSourceFromPb(pb *jobSourcePb) (*JobSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSource{}
	st.DirtyState = pb.DirtyState
	st.ImportFromGitBranch = pb.ImportFromGitBranch
	st.JobConfigPath = pb.JobConfigPath

	return st, nil
}

func jobsHealthRuleToPb(st *JobsHealthRule) (*jobsHealthRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobsHealthRulePb{}
	pb.Metric = st.Metric
	pb.Op = st.Op
	pb.Value = st.Value

	return pb, nil
}

type jobsHealthRulePb struct {
	Metric JobsHealthMetric   `json:"metric"`
	Op     JobsHealthOperator `json:"op"`
	Value  int64              `json:"value"`
}

func jobsHealthRuleFromPb(pb *jobsHealthRulePb) (*JobsHealthRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobsHealthRule{}
	st.Metric = pb.Metric
	st.Op = pb.Op
	st.Value = pb.Value

	return st, nil
}

func jobsHealthRulesToPb(st *JobsHealthRules) (*jobsHealthRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobsHealthRulesPb{}
	pb.Rules = st.Rules

	return pb, nil
}

type jobsHealthRulesPb struct {
	Rules []JobsHealthRule `json:"rules,omitempty"`
}

func jobsHealthRulesFromPb(pb *jobsHealthRulesPb) (*JobsHealthRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobsHealthRules{}
	st.Rules = pb.Rules

	return st, nil
}

func listJobComplianceForPolicyResponseToPb(st *ListJobComplianceForPolicyResponse) (*listJobComplianceForPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobComplianceForPolicyResponsePb{}
	pb.Jobs = st.Jobs
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listJobComplianceForPolicyResponsePb struct {
	Jobs          []JobCompliance `json:"jobs,omitempty"`
	NextPageToken string          `json:"next_page_token,omitempty"`
	PrevPageToken string          `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobComplianceForPolicyResponseFromPb(pb *listJobComplianceForPolicyResponsePb) (*ListJobComplianceForPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobComplianceForPolicyResponse{}
	st.Jobs = pb.Jobs
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listJobComplianceForPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobComplianceForPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listJobComplianceRequestToPb(st *ListJobComplianceRequest) (*listJobComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobComplianceRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listJobComplianceRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`
	PolicyId  string `json:"-" url:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobComplianceRequestFromPb(pb *listJobComplianceRequestPb) (*ListJobComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobComplianceRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listJobComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listJobsRequestToPb(st *ListJobsRequest) (*listJobsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobsRequestPb{}
	pb.ExpandTasks = st.ExpandTasks
	pb.Limit = st.Limit
	pb.Name = st.Name
	pb.Offset = st.Offset
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listJobsRequestPb struct {
	ExpandTasks bool   `json:"-" url:"expand_tasks,omitempty"`
	Limit       int    `json:"-" url:"limit,omitempty"`
	Name        string `json:"-" url:"name,omitempty"`
	Offset      int    `json:"-" url:"offset,omitempty"`
	PageToken   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobsRequestFromPb(pb *listJobsRequestPb) (*ListJobsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobsRequest{}
	st.ExpandTasks = pb.ExpandTasks
	st.Limit = pb.Limit
	st.Name = pb.Name
	st.Offset = pb.Offset
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listJobsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listJobsResponseToPb(st *ListJobsResponse) (*listJobsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobsResponsePb{}
	pb.HasMore = st.HasMore
	pb.Jobs = st.Jobs
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listJobsResponsePb struct {
	HasMore       bool      `json:"has_more,omitempty"`
	Jobs          []BaseJob `json:"jobs,omitempty"`
	NextPageToken string    `json:"next_page_token,omitempty"`
	PrevPageToken string    `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobsResponseFromPb(pb *listJobsResponsePb) (*ListJobsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobsResponse{}
	st.HasMore = pb.HasMore
	st.Jobs = pb.Jobs
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listJobsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRunsRequestToPb(st *ListRunsRequest) (*listRunsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRunsRequestPb{}
	pb.ActiveOnly = st.ActiveOnly
	pb.CompletedOnly = st.CompletedOnly
	pb.ExpandTasks = st.ExpandTasks
	pb.JobId = st.JobId
	pb.Limit = st.Limit
	pb.Offset = st.Offset
	pb.PageToken = st.PageToken
	pb.RunType = st.RunType
	pb.StartTimeFrom = st.StartTimeFrom
	pb.StartTimeTo = st.StartTimeTo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRunsRequestPb struct {
	ActiveOnly    bool    `json:"-" url:"active_only,omitempty"`
	CompletedOnly bool    `json:"-" url:"completed_only,omitempty"`
	ExpandTasks   bool    `json:"-" url:"expand_tasks,omitempty"`
	JobId         int64   `json:"-" url:"job_id,omitempty"`
	Limit         int     `json:"-" url:"limit,omitempty"`
	Offset        int     `json:"-" url:"offset,omitempty"`
	PageToken     string  `json:"-" url:"page_token,omitempty"`
	RunType       RunType `json:"-" url:"run_type,omitempty"`
	StartTimeFrom int64   `json:"-" url:"start_time_from,omitempty"`
	StartTimeTo   int64   `json:"-" url:"start_time_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRunsRequestFromPb(pb *listRunsRequestPb) (*ListRunsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRunsRequest{}
	st.ActiveOnly = pb.ActiveOnly
	st.CompletedOnly = pb.CompletedOnly
	st.ExpandTasks = pb.ExpandTasks
	st.JobId = pb.JobId
	st.Limit = pb.Limit
	st.Offset = pb.Offset
	st.PageToken = pb.PageToken
	st.RunType = pb.RunType
	st.StartTimeFrom = pb.StartTimeFrom
	st.StartTimeTo = pb.StartTimeTo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRunsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRunsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listRunsResponseToPb(st *ListRunsResponse) (*listRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRunsResponsePb{}
	pb.HasMore = st.HasMore
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken
	pb.Runs = st.Runs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listRunsResponsePb struct {
	HasMore       bool      `json:"has_more,omitempty"`
	NextPageToken string    `json:"next_page_token,omitempty"`
	PrevPageToken string    `json:"prev_page_token,omitempty"`
	Runs          []BaseRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRunsResponseFromPb(pb *listRunsResponsePb) (*ListRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRunsResponse{}
	st.HasMore = pb.HasMore
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken
	st.Runs = pb.Runs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func notebookOutputToPb(st *NotebookOutput) (*notebookOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookOutputPb{}
	pb.Result = st.Result
	pb.Truncated = st.Truncated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type notebookOutputPb struct {
	Result    string `json:"result,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookOutputFromPb(pb *notebookOutputPb) (*NotebookOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookOutput{}
	st.Result = pb.Result
	st.Truncated = pb.Truncated

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notebookOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func notebookTaskToPb(st *NotebookTask) (*notebookTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookTaskPb{}
	pb.BaseParameters = st.BaseParameters
	pb.NotebookPath = st.NotebookPath
	pb.Source = st.Source
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type notebookTaskPb struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
	NotebookPath   string            `json:"notebook_path"`
	Source         Source            `json:"source,omitempty"`
	WarehouseId    string            `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookTaskFromPb(pb *notebookTaskPb) (*NotebookTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookTask{}
	st.BaseParameters = pb.BaseParameters
	st.NotebookPath = pb.NotebookPath
	st.Source = pb.Source
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notebookTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func outputSchemaInfoToPb(st *OutputSchemaInfo) (*outputSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &outputSchemaInfoPb{}
	pb.CatalogName = st.CatalogName
	pb.ExpirationTime = st.ExpirationTime
	pb.SchemaName = st.SchemaName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type outputSchemaInfoPb struct {
	CatalogName    string `json:"catalog_name,omitempty"`
	ExpirationTime int64  `json:"expiration_time,omitempty"`
	SchemaName     string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func outputSchemaInfoFromPb(pb *outputSchemaInfoPb) (*OutputSchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OutputSchemaInfo{}
	st.CatalogName = pb.CatalogName
	st.ExpirationTime = pb.ExpirationTime
	st.SchemaName = pb.SchemaName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *outputSchemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st outputSchemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func periodicTriggerConfigurationToPb(st *PeriodicTriggerConfiguration) (*periodicTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &periodicTriggerConfigurationPb{}
	pb.Interval = st.Interval
	pb.Unit = st.Unit

	return pb, nil
}

type periodicTriggerConfigurationPb struct {
	Interval int                                  `json:"interval"`
	Unit     PeriodicTriggerConfigurationTimeUnit `json:"unit"`
}

func periodicTriggerConfigurationFromPb(pb *periodicTriggerConfigurationPb) (*PeriodicTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PeriodicTriggerConfiguration{}
	st.Interval = pb.Interval
	st.Unit = pb.Unit

	return st, nil
}

func pipelineParamsToPb(st *PipelineParams) (*pipelineParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineParamsPb{}
	pb.FullRefresh = st.FullRefresh

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineParamsPb struct {
	FullRefresh bool `json:"full_refresh,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineParamsFromPb(pb *pipelineParamsPb) (*PipelineParams, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineParams{}
	st.FullRefresh = pb.FullRefresh

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineTaskToPb(st *PipelineTask) (*pipelineTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineTaskPb{}
	pb.FullRefresh = st.FullRefresh
	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineTaskPb struct {
	FullRefresh bool   `json:"full_refresh,omitempty"`
	PipelineId  string `json:"pipeline_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineTaskFromPb(pb *pipelineTaskPb) (*PipelineTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineTask{}
	st.FullRefresh = pb.FullRefresh
	st.PipelineId = pb.PipelineId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func powerBiModelToPb(st *PowerBiModel) (*powerBiModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &powerBiModelPb{}
	pb.AuthenticationMethod = st.AuthenticationMethod
	pb.ModelName = st.ModelName
	pb.OverwriteExisting = st.OverwriteExisting
	pb.StorageMode = st.StorageMode
	pb.WorkspaceName = st.WorkspaceName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type powerBiModelPb struct {
	AuthenticationMethod AuthenticationMethod `json:"authentication_method,omitempty"`
	ModelName            string               `json:"model_name,omitempty"`
	OverwriteExisting    bool                 `json:"overwrite_existing,omitempty"`
	StorageMode          StorageMode          `json:"storage_mode,omitempty"`
	WorkspaceName        string               `json:"workspace_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func powerBiModelFromPb(pb *powerBiModelPb) (*PowerBiModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiModel{}
	st.AuthenticationMethod = pb.AuthenticationMethod
	st.ModelName = pb.ModelName
	st.OverwriteExisting = pb.OverwriteExisting
	st.StorageMode = pb.StorageMode
	st.WorkspaceName = pb.WorkspaceName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *powerBiModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st powerBiModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func powerBiTableToPb(st *PowerBiTable) (*powerBiTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &powerBiTablePb{}
	pb.Catalog = st.Catalog
	pb.Name = st.Name
	pb.Schema = st.Schema
	pb.StorageMode = st.StorageMode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type powerBiTablePb struct {
	Catalog     string      `json:"catalog,omitempty"`
	Name        string      `json:"name,omitempty"`
	Schema      string      `json:"schema,omitempty"`
	StorageMode StorageMode `json:"storage_mode,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func powerBiTableFromPb(pb *powerBiTablePb) (*PowerBiTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiTable{}
	st.Catalog = pb.Catalog
	st.Name = pb.Name
	st.Schema = pb.Schema
	st.StorageMode = pb.StorageMode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *powerBiTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st powerBiTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func powerBiTaskToPb(st *PowerBiTask) (*powerBiTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &powerBiTaskPb{}
	pb.ConnectionResourceName = st.ConnectionResourceName
	pb.PowerBiModel = st.PowerBiModel
	pb.RefreshAfterUpdate = st.RefreshAfterUpdate
	pb.Tables = st.Tables
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type powerBiTaskPb struct {
	ConnectionResourceName string         `json:"connection_resource_name,omitempty"`
	PowerBiModel           *PowerBiModel  `json:"power_bi_model,omitempty"`
	RefreshAfterUpdate     bool           `json:"refresh_after_update,omitempty"`
	Tables                 []PowerBiTable `json:"tables,omitempty"`
	WarehouseId            string         `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func powerBiTaskFromPb(pb *powerBiTaskPb) (*PowerBiTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiTask{}
	st.ConnectionResourceName = pb.ConnectionResourceName
	st.PowerBiModel = pb.PowerBiModel
	st.RefreshAfterUpdate = pb.RefreshAfterUpdate
	st.Tables = pb.Tables
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *powerBiTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st powerBiTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pythonWheelTaskToPb(st *PythonWheelTask) (*pythonWheelTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pythonWheelTaskPb{}
	pb.EntryPoint = st.EntryPoint
	pb.NamedParameters = st.NamedParameters
	pb.PackageName = st.PackageName
	pb.Parameters = st.Parameters

	return pb, nil
}

type pythonWheelTaskPb struct {
	EntryPoint      string            `json:"entry_point"`
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	PackageName     string            `json:"package_name"`
	Parameters      []string          `json:"parameters,omitempty"`
}

func pythonWheelTaskFromPb(pb *pythonWheelTaskPb) (*PythonWheelTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PythonWheelTask{}
	st.EntryPoint = pb.EntryPoint
	st.NamedParameters = pb.NamedParameters
	st.PackageName = pb.PackageName
	st.Parameters = pb.Parameters

	return st, nil
}

func queueDetailsToPb(st *QueueDetails) (*queueDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queueDetailsPb{}
	pb.Code = st.Code
	pb.Message = st.Message

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queueDetailsPb struct {
	Code    QueueDetailsCodeCode `json:"code,omitempty"`
	Message string               `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queueDetailsFromPb(pb *queueDetailsPb) (*QueueDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueueDetails{}
	st.Code = pb.Code
	st.Message = pb.Message

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queueDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queueDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func queueSettingsToPb(st *QueueSettings) (*queueSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queueSettingsPb{}
	pb.Enabled = st.Enabled

	return pb, nil
}

type queueSettingsPb struct {
	Enabled bool `json:"enabled"`
}

func queueSettingsFromPb(pb *queueSettingsPb) (*QueueSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueueSettings{}
	st.Enabled = pb.Enabled

	return st, nil
}

func repairHistoryItemToPb(st *RepairHistoryItem) (*repairHistoryItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repairHistoryItemPb{}
	pb.EffectivePerformanceTarget = st.EffectivePerformanceTarget
	pb.EndTime = st.EndTime
	pb.Id = st.Id
	pb.StartTime = st.StartTime
	pb.State = st.State
	pb.Status = st.Status
	pb.TaskRunIds = st.TaskRunIds
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type repairHistoryItemPb struct {
	EffectivePerformanceTarget PerformanceTarget     `json:"effective_performance_target,omitempty"`
	EndTime                    int64                 `json:"end_time,omitempty"`
	Id                         int64                 `json:"id,omitempty"`
	StartTime                  int64                 `json:"start_time,omitempty"`
	State                      *RunState             `json:"state,omitempty"`
	Status                     *RunStatus            `json:"status,omitempty"`
	TaskRunIds                 []int64               `json:"task_run_ids,omitempty"`
	Type                       RepairHistoryItemType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repairHistoryItemFromPb(pb *repairHistoryItemPb) (*RepairHistoryItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairHistoryItem{}
	st.EffectivePerformanceTarget = pb.EffectivePerformanceTarget
	st.EndTime = pb.EndTime
	st.Id = pb.Id
	st.StartTime = pb.StartTime
	st.State = pb.State
	st.Status = pb.Status
	st.TaskRunIds = pb.TaskRunIds
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repairHistoryItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repairHistoryItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func repairRunToPb(st *RepairRun) (*repairRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repairRunPb{}
	pb.DbtCommands = st.DbtCommands
	pb.JarParams = st.JarParams
	pb.JobParameters = st.JobParameters
	pb.LatestRepairId = st.LatestRepairId
	pb.NotebookParams = st.NotebookParams
	pb.PerformanceTarget = st.PerformanceTarget
	pb.PipelineParams = st.PipelineParams
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.RerunAllFailedTasks = st.RerunAllFailedTasks
	pb.RerunDependentTasks = st.RerunDependentTasks
	pb.RerunTasks = st.RerunTasks
	pb.RunId = st.RunId
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type repairRunPb struct {
	DbtCommands         []string          `json:"dbt_commands,omitempty"`
	JarParams           []string          `json:"jar_params,omitempty"`
	JobParameters       map[string]string `json:"job_parameters,omitempty"`
	LatestRepairId      int64             `json:"latest_repair_id,omitempty"`
	NotebookParams      map[string]string `json:"notebook_params,omitempty"`
	PerformanceTarget   PerformanceTarget `json:"performance_target,omitempty"`
	PipelineParams      *PipelineParams   `json:"pipeline_params,omitempty"`
	PythonNamedParams   map[string]string `json:"python_named_params,omitempty"`
	PythonParams        []string          `json:"python_params,omitempty"`
	RerunAllFailedTasks bool              `json:"rerun_all_failed_tasks,omitempty"`
	RerunDependentTasks bool              `json:"rerun_dependent_tasks,omitempty"`
	RerunTasks          []string          `json:"rerun_tasks,omitempty"`
	RunId               int64             `json:"run_id"`
	SparkSubmitParams   []string          `json:"spark_submit_params,omitempty"`
	SqlParams           map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repairRunFromPb(pb *repairRunPb) (*RepairRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairRun{}
	st.DbtCommands = pb.DbtCommands
	st.JarParams = pb.JarParams
	st.JobParameters = pb.JobParameters
	st.LatestRepairId = pb.LatestRepairId
	st.NotebookParams = pb.NotebookParams
	st.PerformanceTarget = pb.PerformanceTarget
	st.PipelineParams = pb.PipelineParams
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.RerunAllFailedTasks = pb.RerunAllFailedTasks
	st.RerunDependentTasks = pb.RerunDependentTasks
	st.RerunTasks = pb.RerunTasks
	st.RunId = pb.RunId
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repairRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repairRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func repairRunResponseToPb(st *RepairRunResponse) (*repairRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repairRunResponsePb{}
	pb.RepairId = st.RepairId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type repairRunResponsePb struct {
	RepairId int64 `json:"repair_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repairRunResponseFromPb(pb *repairRunResponsePb) (*RepairRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairRunResponse{}
	st.RepairId = pb.RepairId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repairRunResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repairRunResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func resetJobToPb(st *ResetJob) (*resetJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resetJobPb{}
	pb.JobId = st.JobId
	pb.NewSettings = st.NewSettings

	return pb, nil
}

type resetJobPb struct {
	JobId       int64       `json:"job_id"`
	NewSettings JobSettings `json:"new_settings"`
}

func resetJobFromPb(pb *resetJobPb) (*ResetJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResetJob{}
	st.JobId = pb.JobId
	st.NewSettings = pb.NewSettings

	return st, nil
}

func resetResponseToPb(st *ResetResponse) (*resetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resetResponsePb{}

	return pb, nil
}

type resetResponsePb struct {
}

func resetResponseFromPb(pb *resetResponsePb) (*ResetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResetResponse{}

	return st, nil
}

func resolvedConditionTaskValuesToPb(st *ResolvedConditionTaskValues) (*resolvedConditionTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedConditionTaskValuesPb{}
	pb.Left = st.Left
	pb.Right = st.Right

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resolvedConditionTaskValuesPb struct {
	Left  string `json:"left,omitempty"`
	Right string `json:"right,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resolvedConditionTaskValuesFromPb(pb *resolvedConditionTaskValuesPb) (*ResolvedConditionTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedConditionTaskValues{}
	st.Left = pb.Left
	st.Right = pb.Right

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resolvedConditionTaskValuesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resolvedConditionTaskValuesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func resolvedDbtTaskValuesToPb(st *ResolvedDbtTaskValues) (*resolvedDbtTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedDbtTaskValuesPb{}
	pb.Commands = st.Commands

	return pb, nil
}

type resolvedDbtTaskValuesPb struct {
	Commands []string `json:"commands,omitempty"`
}

func resolvedDbtTaskValuesFromPb(pb *resolvedDbtTaskValuesPb) (*ResolvedDbtTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedDbtTaskValues{}
	st.Commands = pb.Commands

	return st, nil
}

func resolvedNotebookTaskValuesToPb(st *ResolvedNotebookTaskValues) (*resolvedNotebookTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedNotebookTaskValuesPb{}
	pb.BaseParameters = st.BaseParameters

	return pb, nil
}

type resolvedNotebookTaskValuesPb struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

func resolvedNotebookTaskValuesFromPb(pb *resolvedNotebookTaskValuesPb) (*ResolvedNotebookTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedNotebookTaskValues{}
	st.BaseParameters = pb.BaseParameters

	return st, nil
}

func resolvedParamPairValuesToPb(st *ResolvedParamPairValues) (*resolvedParamPairValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedParamPairValuesPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

type resolvedParamPairValuesPb struct {
	Parameters map[string]string `json:"parameters,omitempty"`
}

func resolvedParamPairValuesFromPb(pb *resolvedParamPairValuesPb) (*ResolvedParamPairValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedParamPairValues{}
	st.Parameters = pb.Parameters

	return st, nil
}

func resolvedPythonWheelTaskValuesToPb(st *ResolvedPythonWheelTaskValues) (*resolvedPythonWheelTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedPythonWheelTaskValuesPb{}
	pb.NamedParameters = st.NamedParameters
	pb.Parameters = st.Parameters

	return pb, nil
}

type resolvedPythonWheelTaskValuesPb struct {
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	Parameters      []string          `json:"parameters,omitempty"`
}

func resolvedPythonWheelTaskValuesFromPb(pb *resolvedPythonWheelTaskValuesPb) (*ResolvedPythonWheelTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedPythonWheelTaskValues{}
	st.NamedParameters = pb.NamedParameters
	st.Parameters = pb.Parameters

	return st, nil
}

func resolvedRunJobTaskValuesToPb(st *ResolvedRunJobTaskValues) (*resolvedRunJobTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedRunJobTaskValuesPb{}
	pb.JobParameters = st.JobParameters
	pb.Parameters = st.Parameters

	return pb, nil
}

type resolvedRunJobTaskValuesPb struct {
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	Parameters    map[string]string `json:"parameters,omitempty"`
}

func resolvedRunJobTaskValuesFromPb(pb *resolvedRunJobTaskValuesPb) (*ResolvedRunJobTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedRunJobTaskValues{}
	st.JobParameters = pb.JobParameters
	st.Parameters = pb.Parameters

	return st, nil
}

func resolvedStringParamsValuesToPb(st *ResolvedStringParamsValues) (*resolvedStringParamsValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedStringParamsValuesPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

type resolvedStringParamsValuesPb struct {
	Parameters []string `json:"parameters,omitempty"`
}

func resolvedStringParamsValuesFromPb(pb *resolvedStringParamsValuesPb) (*ResolvedStringParamsValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedStringParamsValues{}
	st.Parameters = pb.Parameters

	return st, nil
}

func resolvedValuesToPb(st *ResolvedValues) (*resolvedValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedValuesPb{}
	pb.ConditionTask = st.ConditionTask
	pb.DbtTask = st.DbtTask
	pb.NotebookTask = st.NotebookTask
	pb.PythonWheelTask = st.PythonWheelTask
	pb.RunJobTask = st.RunJobTask
	pb.SimulationTask = st.SimulationTask
	pb.SparkJarTask = st.SparkJarTask
	pb.SparkPythonTask = st.SparkPythonTask
	pb.SparkSubmitTask = st.SparkSubmitTask
	pb.SqlTask = st.SqlTask

	return pb, nil
}

type resolvedValuesPb struct {
	ConditionTask   *ResolvedConditionTaskValues   `json:"condition_task,omitempty"`
	DbtTask         *ResolvedDbtTaskValues         `json:"dbt_task,omitempty"`
	NotebookTask    *ResolvedNotebookTaskValues    `json:"notebook_task,omitempty"`
	PythonWheelTask *ResolvedPythonWheelTaskValues `json:"python_wheel_task,omitempty"`
	RunJobTask      *ResolvedRunJobTaskValues      `json:"run_job_task,omitempty"`
	SimulationTask  *ResolvedParamPairValues       `json:"simulation_task,omitempty"`
	SparkJarTask    *ResolvedStringParamsValues    `json:"spark_jar_task,omitempty"`
	SparkPythonTask *ResolvedStringParamsValues    `json:"spark_python_task,omitempty"`
	SparkSubmitTask *ResolvedStringParamsValues    `json:"spark_submit_task,omitempty"`
	SqlTask         *ResolvedParamPairValues       `json:"sql_task,omitempty"`
}

func resolvedValuesFromPb(pb *resolvedValuesPb) (*ResolvedValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedValues{}
	st.ConditionTask = pb.ConditionTask
	st.DbtTask = pb.DbtTask
	st.NotebookTask = pb.NotebookTask
	st.PythonWheelTask = pb.PythonWheelTask
	st.RunJobTask = pb.RunJobTask
	st.SimulationTask = pb.SimulationTask
	st.SparkJarTask = pb.SparkJarTask
	st.SparkPythonTask = pb.SparkPythonTask
	st.SparkSubmitTask = pb.SparkSubmitTask
	st.SqlTask = pb.SqlTask

	return st, nil
}

func runToPb(st *Run) (*runPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runPb{}
	pb.AttemptNumber = st.AttemptNumber
	pb.CleanupDuration = st.CleanupDuration
	pb.ClusterInstance = st.ClusterInstance
	pb.ClusterSpec = st.ClusterSpec
	pb.CreatorUserName = st.CreatorUserName
	pb.Description = st.Description
	pb.EffectivePerformanceTarget = st.EffectivePerformanceTarget
	pb.EndTime = st.EndTime
	pb.ExecutionDuration = st.ExecutionDuration
	pb.GitSource = st.GitSource
	pb.HasMore = st.HasMore
	pb.Iterations = st.Iterations
	pb.JobClusters = st.JobClusters
	pb.JobId = st.JobId
	pb.JobParameters = st.JobParameters
	pb.JobRunId = st.JobRunId
	pb.NextPageToken = st.NextPageToken
	pb.NumberInJob = st.NumberInJob
	pb.OriginalAttemptRunId = st.OriginalAttemptRunId
	pb.OverridingParameters = st.OverridingParameters
	pb.QueueDuration = st.QueueDuration
	pb.RepairHistory = st.RepairHistory
	pb.RunDuration = st.RunDuration
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunPageUrl = st.RunPageUrl
	pb.RunType = st.RunType
	pb.Schedule = st.Schedule
	pb.SetupDuration = st.SetupDuration
	pb.StartTime = st.StartTime
	pb.State = st.State
	pb.Status = st.Status
	pb.Tasks = st.Tasks
	pb.Trigger = st.Trigger
	pb.TriggerInfo = st.TriggerInfo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runPb struct {
	AttemptNumber              int                 `json:"attempt_number,omitempty"`
	CleanupDuration            int64               `json:"cleanup_duration,omitempty"`
	ClusterInstance            *ClusterInstance    `json:"cluster_instance,omitempty"`
	ClusterSpec                *ClusterSpec        `json:"cluster_spec,omitempty"`
	CreatorUserName            string              `json:"creator_user_name,omitempty"`
	Description                string              `json:"description,omitempty"`
	EffectivePerformanceTarget PerformanceTarget   `json:"effective_performance_target,omitempty"`
	EndTime                    int64               `json:"end_time,omitempty"`
	ExecutionDuration          int64               `json:"execution_duration,omitempty"`
	GitSource                  *GitSource          `json:"git_source,omitempty"`
	HasMore                    bool                `json:"has_more,omitempty"`
	Iterations                 []RunTask           `json:"iterations,omitempty"`
	JobClusters                []JobCluster        `json:"job_clusters,omitempty"`
	JobId                      int64               `json:"job_id,omitempty"`
	JobParameters              []JobParameter      `json:"job_parameters,omitempty"`
	JobRunId                   int64               `json:"job_run_id,omitempty"`
	NextPageToken              string              `json:"next_page_token,omitempty"`
	NumberInJob                int64               `json:"number_in_job,omitempty"`
	OriginalAttemptRunId       int64               `json:"original_attempt_run_id,omitempty"`
	OverridingParameters       *RunParameters      `json:"overriding_parameters,omitempty"`
	QueueDuration              int64               `json:"queue_duration,omitempty"`
	RepairHistory              []RepairHistoryItem `json:"repair_history,omitempty"`
	RunDuration                int64               `json:"run_duration,omitempty"`
	RunId                      int64               `json:"run_id,omitempty"`
	RunName                    string              `json:"run_name,omitempty"`
	RunPageUrl                 string              `json:"run_page_url,omitempty"`
	RunType                    RunType             `json:"run_type,omitempty"`
	Schedule                   *CronSchedule       `json:"schedule,omitempty"`
	SetupDuration              int64               `json:"setup_duration,omitempty"`
	StartTime                  int64               `json:"start_time,omitempty"`
	State                      *RunState           `json:"state,omitempty"`
	Status                     *RunStatus          `json:"status,omitempty"`
	Tasks                      []RunTask           `json:"tasks,omitempty"`
	Trigger                    TriggerType         `json:"trigger,omitempty"`
	TriggerInfo                *TriggerInfo        `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runFromPb(pb *runPb) (*Run, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Run{}
	st.AttemptNumber = pb.AttemptNumber
	st.CleanupDuration = pb.CleanupDuration
	st.ClusterInstance = pb.ClusterInstance
	st.ClusterSpec = pb.ClusterSpec
	st.CreatorUserName = pb.CreatorUserName
	st.Description = pb.Description
	st.EffectivePerformanceTarget = pb.EffectivePerformanceTarget
	st.EndTime = pb.EndTime
	st.ExecutionDuration = pb.ExecutionDuration
	st.GitSource = pb.GitSource
	st.HasMore = pb.HasMore
	st.Iterations = pb.Iterations
	st.JobClusters = pb.JobClusters
	st.JobId = pb.JobId
	st.JobParameters = pb.JobParameters
	st.JobRunId = pb.JobRunId
	st.NextPageToken = pb.NextPageToken
	st.NumberInJob = pb.NumberInJob
	st.OriginalAttemptRunId = pb.OriginalAttemptRunId
	st.OverridingParameters = pb.OverridingParameters
	st.QueueDuration = pb.QueueDuration
	st.RepairHistory = pb.RepairHistory
	st.RunDuration = pb.RunDuration
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunPageUrl = pb.RunPageUrl
	st.RunType = pb.RunType
	st.Schedule = pb.Schedule
	st.SetupDuration = pb.SetupDuration
	st.StartTime = pb.StartTime
	st.State = pb.State
	st.Status = pb.Status
	st.Tasks = pb.Tasks
	st.Trigger = pb.Trigger
	st.TriggerInfo = pb.TriggerInfo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runConditionTaskToPb(st *RunConditionTask) (*runConditionTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runConditionTaskPb{}
	pb.Left = st.Left
	pb.Op = st.Op
	pb.Outcome = st.Outcome
	pb.Right = st.Right

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runConditionTaskPb struct {
	Left    string          `json:"left"`
	Op      ConditionTaskOp `json:"op"`
	Outcome string          `json:"outcome,omitempty"`
	Right   string          `json:"right"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runConditionTaskFromPb(pb *runConditionTaskPb) (*RunConditionTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunConditionTask{}
	st.Left = pb.Left
	st.Op = pb.Op
	st.Outcome = pb.Outcome
	st.Right = pb.Right

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runConditionTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runConditionTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runForEachTaskToPb(st *RunForEachTask) (*runForEachTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runForEachTaskPb{}
	pb.Concurrency = st.Concurrency
	pb.Inputs = st.Inputs
	pb.Stats = st.Stats
	pb.Task = st.Task

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runForEachTaskPb struct {
	Concurrency int           `json:"concurrency,omitempty"`
	Inputs      string        `json:"inputs"`
	Stats       *ForEachStats `json:"stats,omitempty"`
	Task        Task          `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runForEachTaskFromPb(pb *runForEachTaskPb) (*RunForEachTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunForEachTask{}
	st.Concurrency = pb.Concurrency
	st.Inputs = pb.Inputs
	st.Stats = pb.Stats
	st.Task = pb.Task

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runForEachTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runForEachTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runJobOutputToPb(st *RunJobOutput) (*runJobOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runJobOutputPb{}
	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runJobOutputPb struct {
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runJobOutputFromPb(pb *runJobOutputPb) (*RunJobOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunJobOutput{}
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runJobOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runJobOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runJobTaskToPb(st *RunJobTask) (*runJobTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runJobTaskPb{}
	pb.DbtCommands = st.DbtCommands
	pb.JarParams = st.JarParams
	pb.JobId = st.JobId
	pb.JobParameters = st.JobParameters
	pb.NotebookParams = st.NotebookParams
	pb.PipelineParams = st.PipelineParams
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	return pb, nil
}

type runJobTaskPb struct {
	DbtCommands       []string          `json:"dbt_commands,omitempty"`
	JarParams         []string          `json:"jar_params,omitempty"`
	JobId             int64             `json:"job_id"`
	JobParameters     map[string]string `json:"job_parameters,omitempty"`
	NotebookParams    map[string]string `json:"notebook_params,omitempty"`
	PipelineParams    *PipelineParams   `json:"pipeline_params,omitempty"`
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	PythonParams      []string          `json:"python_params,omitempty"`
	SparkSubmitParams []string          `json:"spark_submit_params,omitempty"`
	SqlParams         map[string]string `json:"sql_params,omitempty"`
}

func runJobTaskFromPb(pb *runJobTaskPb) (*RunJobTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunJobTask{}
	st.DbtCommands = pb.DbtCommands
	st.JarParams = pb.JarParams
	st.JobId = pb.JobId
	st.JobParameters = pb.JobParameters
	st.NotebookParams = pb.NotebookParams
	st.PipelineParams = pb.PipelineParams
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	return st, nil
}

func runNowToPb(st *RunNow) (*runNowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runNowPb{}
	pb.DbtCommands = st.DbtCommands
	pb.IdempotencyToken = st.IdempotencyToken
	pb.JarParams = st.JarParams
	pb.JobId = st.JobId
	pb.JobParameters = st.JobParameters
	pb.NotebookParams = st.NotebookParams
	pb.Only = st.Only
	pb.PerformanceTarget = st.PerformanceTarget
	pb.PipelineParams = st.PipelineParams
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.Queue = st.Queue
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runNowPb struct {
	DbtCommands       []string          `json:"dbt_commands,omitempty"`
	IdempotencyToken  string            `json:"idempotency_token,omitempty"`
	JarParams         []string          `json:"jar_params,omitempty"`
	JobId             int64             `json:"job_id"`
	JobParameters     map[string]string `json:"job_parameters,omitempty"`
	NotebookParams    map[string]string `json:"notebook_params,omitempty"`
	Only              []string          `json:"only,omitempty"`
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	PipelineParams    *PipelineParams   `json:"pipeline_params,omitempty"`
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	PythonParams      []string          `json:"python_params,omitempty"`
	Queue             *QueueSettings    `json:"queue,omitempty"`
	SparkSubmitParams []string          `json:"spark_submit_params,omitempty"`
	SqlParams         map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runNowFromPb(pb *runNowPb) (*RunNow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunNow{}
	st.DbtCommands = pb.DbtCommands
	st.IdempotencyToken = pb.IdempotencyToken
	st.JarParams = pb.JarParams
	st.JobId = pb.JobId
	st.JobParameters = pb.JobParameters
	st.NotebookParams = pb.NotebookParams
	st.Only = pb.Only
	st.PerformanceTarget = pb.PerformanceTarget
	st.PipelineParams = pb.PipelineParams
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.Queue = pb.Queue
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runNowPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runNowPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runNowResponseToPb(st *RunNowResponse) (*runNowResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runNowResponsePb{}
	pb.NumberInJob = st.NumberInJob
	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runNowResponsePb struct {
	NumberInJob int64 `json:"number_in_job,omitempty"`
	RunId       int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runNowResponseFromPb(pb *runNowResponsePb) (*RunNowResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunNowResponse{}
	st.NumberInJob = pb.NumberInJob
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runNowResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runNowResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runOutputToPb(st *RunOutput) (*runOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runOutputPb{}
	pb.CleanRoomsNotebookOutput = st.CleanRoomsNotebookOutput
	pb.DashboardOutput = st.DashboardOutput
	pb.DbtOutput = st.DbtOutput
	pb.Error = st.Error
	pb.ErrorTrace = st.ErrorTrace
	pb.Info = st.Info
	pb.Logs = st.Logs
	pb.LogsTruncated = st.LogsTruncated
	pb.Metadata = st.Metadata
	pb.NotebookOutput = st.NotebookOutput
	pb.RunJobOutput = st.RunJobOutput
	pb.SqlOutput = st.SqlOutput

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runOutputPb struct {
	CleanRoomsNotebookOutput *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput `json:"clean_rooms_notebook_output,omitempty"`
	DashboardOutput          *DashboardTaskOutput                                `json:"dashboard_output,omitempty"`
	DbtOutput                *DbtOutput                                          `json:"dbt_output,omitempty"`
	Error                    string                                              `json:"error,omitempty"`
	ErrorTrace               string                                              `json:"error_trace,omitempty"`
	Info                     string                                              `json:"info,omitempty"`
	Logs                     string                                              `json:"logs,omitempty"`
	LogsTruncated            bool                                                `json:"logs_truncated,omitempty"`
	Metadata                 *Run                                                `json:"metadata,omitempty"`
	NotebookOutput           *NotebookOutput                                     `json:"notebook_output,omitempty"`
	RunJobOutput             *RunJobOutput                                       `json:"run_job_output,omitempty"`
	SqlOutput                *SqlOutput                                          `json:"sql_output,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runOutputFromPb(pb *runOutputPb) (*RunOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunOutput{}
	st.CleanRoomsNotebookOutput = pb.CleanRoomsNotebookOutput
	st.DashboardOutput = pb.DashboardOutput
	st.DbtOutput = pb.DbtOutput
	st.Error = pb.Error
	st.ErrorTrace = pb.ErrorTrace
	st.Info = pb.Info
	st.Logs = pb.Logs
	st.LogsTruncated = pb.LogsTruncated
	st.Metadata = pb.Metadata
	st.NotebookOutput = pb.NotebookOutput
	st.RunJobOutput = pb.RunJobOutput
	st.SqlOutput = pb.SqlOutput

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runParametersToPb(st *RunParameters) (*runParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runParametersPb{}
	pb.DbtCommands = st.DbtCommands
	pb.JarParams = st.JarParams
	pb.NotebookParams = st.NotebookParams
	pb.PipelineParams = st.PipelineParams
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	return pb, nil
}

type runParametersPb struct {
	DbtCommands       []string          `json:"dbt_commands,omitempty"`
	JarParams         []string          `json:"jar_params,omitempty"`
	NotebookParams    map[string]string `json:"notebook_params,omitempty"`
	PipelineParams    *PipelineParams   `json:"pipeline_params,omitempty"`
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	PythonParams      []string          `json:"python_params,omitempty"`
	SparkSubmitParams []string          `json:"spark_submit_params,omitempty"`
	SqlParams         map[string]string `json:"sql_params,omitempty"`
}

func runParametersFromPb(pb *runParametersPb) (*RunParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunParameters{}
	st.DbtCommands = pb.DbtCommands
	st.JarParams = pb.JarParams
	st.NotebookParams = pb.NotebookParams
	st.PipelineParams = pb.PipelineParams
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	return st, nil
}

func runStateToPb(st *RunState) (*runStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runStatePb{}
	pb.LifeCycleState = st.LifeCycleState
	pb.QueueReason = st.QueueReason
	pb.ResultState = st.ResultState
	pb.StateMessage = st.StateMessage
	pb.UserCancelledOrTimedout = st.UserCancelledOrTimedout

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runStatePb struct {
	LifeCycleState          RunLifeCycleState `json:"life_cycle_state,omitempty"`
	QueueReason             string            `json:"queue_reason,omitempty"`
	ResultState             RunResultState    `json:"result_state,omitempty"`
	StateMessage            string            `json:"state_message,omitempty"`
	UserCancelledOrTimedout bool              `json:"user_cancelled_or_timedout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runStateFromPb(pb *runStatePb) (*RunState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunState{}
	st.LifeCycleState = pb.LifeCycleState
	st.QueueReason = pb.QueueReason
	st.ResultState = pb.ResultState
	st.StateMessage = pb.StateMessage
	st.UserCancelledOrTimedout = pb.UserCancelledOrTimedout

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runStatusToPb(st *RunStatus) (*runStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runStatusPb{}
	pb.QueueDetails = st.QueueDetails
	pb.State = st.State
	pb.TerminationDetails = st.TerminationDetails

	return pb, nil
}

type runStatusPb struct {
	QueueDetails       *QueueDetails            `json:"queue_details,omitempty"`
	State              RunLifecycleStateV2State `json:"state,omitempty"`
	TerminationDetails *TerminationDetails      `json:"termination_details,omitempty"`
}

func runStatusFromPb(pb *runStatusPb) (*RunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunStatus{}
	st.QueueDetails = pb.QueueDetails
	st.State = pb.State
	st.TerminationDetails = pb.TerminationDetails

	return st, nil
}

func runTaskToPb(st *RunTask) (*runTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runTaskPb{}
	pb.AttemptNumber = st.AttemptNumber
	pb.CleanRoomsNotebookTask = st.CleanRoomsNotebookTask
	pb.CleanupDuration = st.CleanupDuration
	pb.ClusterInstance = st.ClusterInstance
	pb.ConditionTask = st.ConditionTask
	pb.DashboardTask = st.DashboardTask
	pb.DbtTask = st.DbtTask
	pb.DependsOn = st.DependsOn
	pb.Description = st.Description
	pb.Disabled = st.Disabled
	pb.EffectivePerformanceTarget = st.EffectivePerformanceTarget
	pb.EmailNotifications = st.EmailNotifications
	pb.EndTime = st.EndTime
	pb.EnvironmentKey = st.EnvironmentKey
	pb.ExecutionDuration = st.ExecutionDuration
	pb.ExistingClusterId = st.ExistingClusterId
	pb.ForEachTask = st.ForEachTask
	pb.GenAiComputeTask = st.GenAiComputeTask
	pb.GitSource = st.GitSource
	pb.JobClusterKey = st.JobClusterKey
	pb.Libraries = st.Libraries
	pb.NewCluster = st.NewCluster
	pb.NotebookTask = st.NotebookTask
	pb.NotificationSettings = st.NotificationSettings
	pb.PipelineTask = st.PipelineTask
	pb.PowerBiTask = st.PowerBiTask
	pb.PythonWheelTask = st.PythonWheelTask
	pb.QueueDuration = st.QueueDuration
	pb.ResolvedValues = st.ResolvedValues
	pb.RunDuration = st.RunDuration
	pb.RunId = st.RunId
	pb.RunIf = st.RunIf
	pb.RunJobTask = st.RunJobTask
	pb.RunPageUrl = st.RunPageUrl
	pb.SetupDuration = st.SetupDuration
	pb.SparkJarTask = st.SparkJarTask
	pb.SparkPythonTask = st.SparkPythonTask
	pb.SparkSubmitTask = st.SparkSubmitTask
	pb.SqlTask = st.SqlTask
	pb.StartTime = st.StartTime
	pb.State = st.State
	pb.Status = st.Status
	pb.TaskKey = st.TaskKey
	pb.TimeoutSeconds = st.TimeoutSeconds
	pb.WebhookNotifications = st.WebhookNotifications

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runTaskPb struct {
	AttemptNumber              int                       `json:"attempt_number,omitempty"`
	CleanRoomsNotebookTask     *CleanRoomsNotebookTask   `json:"clean_rooms_notebook_task,omitempty"`
	CleanupDuration            int64                     `json:"cleanup_duration,omitempty"`
	ClusterInstance            *ClusterInstance          `json:"cluster_instance,omitempty"`
	ConditionTask              *RunConditionTask         `json:"condition_task,omitempty"`
	DashboardTask              *DashboardTask            `json:"dashboard_task,omitempty"`
	DbtTask                    *DbtTask                  `json:"dbt_task,omitempty"`
	DependsOn                  []TaskDependency          `json:"depends_on,omitempty"`
	Description                string                    `json:"description,omitempty"`
	Disabled                   bool                      `json:"disabled,omitempty"`
	EffectivePerformanceTarget PerformanceTarget         `json:"effective_performance_target,omitempty"`
	EmailNotifications         *JobEmailNotifications    `json:"email_notifications,omitempty"`
	EndTime                    int64                     `json:"end_time,omitempty"`
	EnvironmentKey             string                    `json:"environment_key,omitempty"`
	ExecutionDuration          int64                     `json:"execution_duration,omitempty"`
	ExistingClusterId          string                    `json:"existing_cluster_id,omitempty"`
	ForEachTask                *RunForEachTask           `json:"for_each_task,omitempty"`
	GenAiComputeTask           *GenAiComputeTask         `json:"gen_ai_compute_task,omitempty"`
	GitSource                  *GitSource                `json:"git_source,omitempty"`
	JobClusterKey              string                    `json:"job_cluster_key,omitempty"`
	Libraries                  []compute.Library         `json:"libraries,omitempty"`
	NewCluster                 *compute.ClusterSpec      `json:"new_cluster,omitempty"`
	NotebookTask               *NotebookTask             `json:"notebook_task,omitempty"`
	NotificationSettings       *TaskNotificationSettings `json:"notification_settings,omitempty"`
	PipelineTask               *PipelineTask             `json:"pipeline_task,omitempty"`
	PowerBiTask                *PowerBiTask              `json:"power_bi_task,omitempty"`
	PythonWheelTask            *PythonWheelTask          `json:"python_wheel_task,omitempty"`
	QueueDuration              int64                     `json:"queue_duration,omitempty"`
	ResolvedValues             *ResolvedValues           `json:"resolved_values,omitempty"`
	RunDuration                int64                     `json:"run_duration,omitempty"`
	RunId                      int64                     `json:"run_id,omitempty"`
	RunIf                      RunIf                     `json:"run_if,omitempty"`
	RunJobTask                 *RunJobTask               `json:"run_job_task,omitempty"`
	RunPageUrl                 string                    `json:"run_page_url,omitempty"`
	SetupDuration              int64                     `json:"setup_duration,omitempty"`
	SparkJarTask               *SparkJarTask             `json:"spark_jar_task,omitempty"`
	SparkPythonTask            *SparkPythonTask          `json:"spark_python_task,omitempty"`
	SparkSubmitTask            *SparkSubmitTask          `json:"spark_submit_task,omitempty"`
	SqlTask                    *SqlTask                  `json:"sql_task,omitempty"`
	StartTime                  int64                     `json:"start_time,omitempty"`
	State                      *RunState                 `json:"state,omitempty"`
	Status                     *RunStatus                `json:"status,omitempty"`
	TaskKey                    string                    `json:"task_key"`
	TimeoutSeconds             int                       `json:"timeout_seconds,omitempty"`
	WebhookNotifications       *WebhookNotifications     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runTaskFromPb(pb *runTaskPb) (*RunTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunTask{}
	st.AttemptNumber = pb.AttemptNumber
	st.CleanRoomsNotebookTask = pb.CleanRoomsNotebookTask
	st.CleanupDuration = pb.CleanupDuration
	st.ClusterInstance = pb.ClusterInstance
	st.ConditionTask = pb.ConditionTask
	st.DashboardTask = pb.DashboardTask
	st.DbtTask = pb.DbtTask
	st.DependsOn = pb.DependsOn
	st.Description = pb.Description
	st.Disabled = pb.Disabled
	st.EffectivePerformanceTarget = pb.EffectivePerformanceTarget
	st.EmailNotifications = pb.EmailNotifications
	st.EndTime = pb.EndTime
	st.EnvironmentKey = pb.EnvironmentKey
	st.ExecutionDuration = pb.ExecutionDuration
	st.ExistingClusterId = pb.ExistingClusterId
	st.ForEachTask = pb.ForEachTask
	st.GenAiComputeTask = pb.GenAiComputeTask
	st.GitSource = pb.GitSource
	st.JobClusterKey = pb.JobClusterKey
	st.Libraries = pb.Libraries
	st.NewCluster = pb.NewCluster
	st.NotebookTask = pb.NotebookTask
	st.NotificationSettings = pb.NotificationSettings
	st.PipelineTask = pb.PipelineTask
	st.PowerBiTask = pb.PowerBiTask
	st.PythonWheelTask = pb.PythonWheelTask
	st.QueueDuration = pb.QueueDuration
	st.ResolvedValues = pb.ResolvedValues
	st.RunDuration = pb.RunDuration
	st.RunId = pb.RunId
	st.RunIf = pb.RunIf
	st.RunJobTask = pb.RunJobTask
	st.RunPageUrl = pb.RunPageUrl
	st.SetupDuration = pb.SetupDuration
	st.SparkJarTask = pb.SparkJarTask
	st.SparkPythonTask = pb.SparkPythonTask
	st.SparkSubmitTask = pb.SparkSubmitTask
	st.SqlTask = pb.SqlTask
	st.StartTime = pb.StartTime
	st.State = pb.State
	st.Status = pb.Status
	st.TaskKey = pb.TaskKey
	st.TimeoutSeconds = pb.TimeoutSeconds
	st.WebhookNotifications = pb.WebhookNotifications

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sparkJarTaskToPb(st *SparkJarTask) (*sparkJarTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkJarTaskPb{}
	pb.JarUri = st.JarUri
	pb.MainClassName = st.MainClassName
	pb.Parameters = st.Parameters
	pb.RunAsRepl = st.RunAsRepl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sparkJarTaskPb struct {
	JarUri        string   `json:"jar_uri,omitempty"`
	MainClassName string   `json:"main_class_name,omitempty"`
	Parameters    []string `json:"parameters,omitempty"`
	RunAsRepl     bool     `json:"run_as_repl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sparkJarTaskFromPb(pb *sparkJarTaskPb) (*SparkJarTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkJarTask{}
	st.JarUri = pb.JarUri
	st.MainClassName = pb.MainClassName
	st.Parameters = pb.Parameters
	st.RunAsRepl = pb.RunAsRepl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkJarTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkJarTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sparkPythonTaskToPb(st *SparkPythonTask) (*sparkPythonTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkPythonTaskPb{}
	pb.Parameters = st.Parameters
	pb.PythonFile = st.PythonFile
	pb.Source = st.Source

	return pb, nil
}

type sparkPythonTaskPb struct {
	Parameters []string `json:"parameters,omitempty"`
	PythonFile string   `json:"python_file"`
	Source     Source   `json:"source,omitempty"`
}

func sparkPythonTaskFromPb(pb *sparkPythonTaskPb) (*SparkPythonTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkPythonTask{}
	st.Parameters = pb.Parameters
	st.PythonFile = pb.PythonFile
	st.Source = pb.Source

	return st, nil
}

func sparkSubmitTaskToPb(st *SparkSubmitTask) (*sparkSubmitTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkSubmitTaskPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

type sparkSubmitTaskPb struct {
	Parameters []string `json:"parameters,omitempty"`
}

func sparkSubmitTaskFromPb(pb *sparkSubmitTaskPb) (*SparkSubmitTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkSubmitTask{}
	st.Parameters = pb.Parameters

	return st, nil
}

func sqlAlertOutputToPb(st *SqlAlertOutput) (*sqlAlertOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlAlertOutputPb{}
	pb.AlertState = st.AlertState
	pb.OutputLink = st.OutputLink
	pb.QueryText = st.QueryText
	pb.SqlStatements = st.SqlStatements
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlAlertOutputPb struct {
	AlertState    SqlAlertState        `json:"alert_state,omitempty"`
	OutputLink    string               `json:"output_link,omitempty"`
	QueryText     string               `json:"query_text,omitempty"`
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	WarehouseId   string               `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlAlertOutputFromPb(pb *sqlAlertOutputPb) (*SqlAlertOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlAlertOutput{}
	st.AlertState = pb.AlertState
	st.OutputLink = pb.OutputLink
	st.QueryText = pb.QueryText
	st.SqlStatements = pb.SqlStatements
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlAlertOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlAlertOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlDashboardOutputToPb(st *SqlDashboardOutput) (*sqlDashboardOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlDashboardOutputPb{}
	pb.WarehouseId = st.WarehouseId
	pb.Widgets = st.Widgets

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlDashboardOutputPb struct {
	WarehouseId string                     `json:"warehouse_id,omitempty"`
	Widgets     []SqlDashboardWidgetOutput `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlDashboardOutputFromPb(pb *sqlDashboardOutputPb) (*SqlDashboardOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlDashboardOutput{}
	st.WarehouseId = pb.WarehouseId
	st.Widgets = pb.Widgets

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlDashboardOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlDashboardOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlDashboardWidgetOutputToPb(st *SqlDashboardWidgetOutput) (*sqlDashboardWidgetOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlDashboardWidgetOutputPb{}
	pb.EndTime = st.EndTime
	pb.Error = st.Error
	pb.OutputLink = st.OutputLink
	pb.StartTime = st.StartTime
	pb.Status = st.Status
	pb.WidgetId = st.WidgetId
	pb.WidgetTitle = st.WidgetTitle

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlDashboardWidgetOutputPb struct {
	EndTime     int64                          `json:"end_time,omitempty"`
	Error       *SqlOutputError                `json:"error,omitempty"`
	OutputLink  string                         `json:"output_link,omitempty"`
	StartTime   int64                          `json:"start_time,omitempty"`
	Status      SqlDashboardWidgetOutputStatus `json:"status,omitempty"`
	WidgetId    string                         `json:"widget_id,omitempty"`
	WidgetTitle string                         `json:"widget_title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlDashboardWidgetOutputFromPb(pb *sqlDashboardWidgetOutputPb) (*SqlDashboardWidgetOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlDashboardWidgetOutput{}
	st.EndTime = pb.EndTime
	st.Error = pb.Error
	st.OutputLink = pb.OutputLink
	st.StartTime = pb.StartTime
	st.Status = pb.Status
	st.WidgetId = pb.WidgetId
	st.WidgetTitle = pb.WidgetTitle

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlDashboardWidgetOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlDashboardWidgetOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlOutputToPb(st *SqlOutput) (*sqlOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlOutputPb{}
	pb.AlertOutput = st.AlertOutput
	pb.DashboardOutput = st.DashboardOutput
	pb.QueryOutput = st.QueryOutput

	return pb, nil
}

type sqlOutputPb struct {
	AlertOutput     *SqlAlertOutput     `json:"alert_output,omitempty"`
	DashboardOutput *SqlDashboardOutput `json:"dashboard_output,omitempty"`
	QueryOutput     *SqlQueryOutput     `json:"query_output,omitempty"`
}

func sqlOutputFromPb(pb *sqlOutputPb) (*SqlOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlOutput{}
	st.AlertOutput = pb.AlertOutput
	st.DashboardOutput = pb.DashboardOutput
	st.QueryOutput = pb.QueryOutput

	return st, nil
}

func sqlOutputErrorToPb(st *SqlOutputError) (*sqlOutputErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlOutputErrorPb{}
	pb.Message = st.Message

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlOutputErrorPb struct {
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlOutputErrorFromPb(pb *sqlOutputErrorPb) (*SqlOutputError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlOutputError{}
	st.Message = pb.Message

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlOutputErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlOutputErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlQueryOutputToPb(st *SqlQueryOutput) (*sqlQueryOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlQueryOutputPb{}
	pb.EndpointId = st.EndpointId
	pb.OutputLink = st.OutputLink
	pb.QueryText = st.QueryText
	pb.SqlStatements = st.SqlStatements
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlQueryOutputPb struct {
	EndpointId    string               `json:"endpoint_id,omitempty"`
	OutputLink    string               `json:"output_link,omitempty"`
	QueryText     string               `json:"query_text,omitempty"`
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	WarehouseId   string               `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlQueryOutputFromPb(pb *sqlQueryOutputPb) (*SqlQueryOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlQueryOutput{}
	st.EndpointId = pb.EndpointId
	st.OutputLink = pb.OutputLink
	st.QueryText = pb.QueryText
	st.SqlStatements = pb.SqlStatements
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlQueryOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlQueryOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlStatementOutputToPb(st *SqlStatementOutput) (*sqlStatementOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlStatementOutputPb{}
	pb.LookupKey = st.LookupKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlStatementOutputPb struct {
	LookupKey string `json:"lookup_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlStatementOutputFromPb(pb *sqlStatementOutputPb) (*SqlStatementOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlStatementOutput{}
	st.LookupKey = pb.LookupKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlStatementOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlStatementOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlTaskToPb(st *SqlTask) (*sqlTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskPb{}
	pb.Alert = st.Alert
	pb.Dashboard = st.Dashboard
	pb.File = st.File
	pb.Parameters = st.Parameters
	pb.Query = st.Query
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

type sqlTaskPb struct {
	Alert       *SqlTaskAlert     `json:"alert,omitempty"`
	Dashboard   *SqlTaskDashboard `json:"dashboard,omitempty"`
	File        *SqlTaskFile      `json:"file,omitempty"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	Query       *SqlTaskQuery     `json:"query,omitempty"`
	WarehouseId string            `json:"warehouse_id"`
}

func sqlTaskFromPb(pb *sqlTaskPb) (*SqlTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTask{}
	st.Alert = pb.Alert
	st.Dashboard = pb.Dashboard
	st.File = pb.File
	st.Parameters = pb.Parameters
	st.Query = pb.Query
	st.WarehouseId = pb.WarehouseId

	return st, nil
}

func sqlTaskAlertToPb(st *SqlTaskAlert) (*sqlTaskAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskAlertPb{}
	pb.AlertId = st.AlertId
	pb.PauseSubscriptions = st.PauseSubscriptions
	pb.Subscriptions = st.Subscriptions

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlTaskAlertPb struct {
	AlertId            string                `json:"alert_id"`
	PauseSubscriptions bool                  `json:"pause_subscriptions,omitempty"`
	Subscriptions      []SqlTaskSubscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlTaskAlertFromPb(pb *sqlTaskAlertPb) (*SqlTaskAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskAlert{}
	st.AlertId = pb.AlertId
	st.PauseSubscriptions = pb.PauseSubscriptions
	st.Subscriptions = pb.Subscriptions

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlTaskAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlTaskAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlTaskDashboardToPb(st *SqlTaskDashboard) (*sqlTaskDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskDashboardPb{}
	pb.CustomSubject = st.CustomSubject
	pb.DashboardId = st.DashboardId
	pb.PauseSubscriptions = st.PauseSubscriptions
	pb.Subscriptions = st.Subscriptions

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlTaskDashboardPb struct {
	CustomSubject      string                `json:"custom_subject,omitempty"`
	DashboardId        string                `json:"dashboard_id"`
	PauseSubscriptions bool                  `json:"pause_subscriptions,omitempty"`
	Subscriptions      []SqlTaskSubscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlTaskDashboardFromPb(pb *sqlTaskDashboardPb) (*SqlTaskDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskDashboard{}
	st.CustomSubject = pb.CustomSubject
	st.DashboardId = pb.DashboardId
	st.PauseSubscriptions = pb.PauseSubscriptions
	st.Subscriptions = pb.Subscriptions

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlTaskDashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlTaskDashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlTaskFileToPb(st *SqlTaskFile) (*sqlTaskFilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskFilePb{}
	pb.Path = st.Path
	pb.Source = st.Source

	return pb, nil
}

type sqlTaskFilePb struct {
	Path   string `json:"path"`
	Source Source `json:"source,omitempty"`
}

func sqlTaskFileFromPb(pb *sqlTaskFilePb) (*SqlTaskFile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskFile{}
	st.Path = pb.Path
	st.Source = pb.Source

	return st, nil
}

func sqlTaskQueryToPb(st *SqlTaskQuery) (*sqlTaskQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskQueryPb{}
	pb.QueryId = st.QueryId

	return pb, nil
}

type sqlTaskQueryPb struct {
	QueryId string `json:"query_id"`
}

func sqlTaskQueryFromPb(pb *sqlTaskQueryPb) (*SqlTaskQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskQuery{}
	st.QueryId = pb.QueryId

	return st, nil
}

func sqlTaskSubscriptionToPb(st *SqlTaskSubscription) (*sqlTaskSubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskSubscriptionPb{}
	pb.DestinationId = st.DestinationId
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlTaskSubscriptionPb struct {
	DestinationId string `json:"destination_id,omitempty"`
	UserName      string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlTaskSubscriptionFromPb(pb *sqlTaskSubscriptionPb) (*SqlTaskSubscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskSubscription{}
	st.DestinationId = pb.DestinationId
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlTaskSubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlTaskSubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func submitRunToPb(st *SubmitRun) (*submitRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &submitRunPb{}
	pb.AccessControlList = st.AccessControlList
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.EmailNotifications = st.EmailNotifications
	pb.Environments = st.Environments
	pb.GitSource = st.GitSource
	pb.Health = st.Health
	pb.IdempotencyToken = st.IdempotencyToken
	pb.NotificationSettings = st.NotificationSettings
	pb.Queue = st.Queue
	pb.RunAs = st.RunAs
	pb.RunName = st.RunName
	pb.Tasks = st.Tasks
	pb.TimeoutSeconds = st.TimeoutSeconds
	pb.WebhookNotifications = st.WebhookNotifications

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type submitRunPb struct {
	AccessControlList    []JobAccessControlRequest `json:"access_control_list,omitempty"`
	BudgetPolicyId       string                    `json:"budget_policy_id,omitempty"`
	EmailNotifications   *JobEmailNotifications    `json:"email_notifications,omitempty"`
	Environments         []JobEnvironment          `json:"environments,omitempty"`
	GitSource            *GitSource                `json:"git_source,omitempty"`
	Health               *JobsHealthRules          `json:"health,omitempty"`
	IdempotencyToken     string                    `json:"idempotency_token,omitempty"`
	NotificationSettings *JobNotificationSettings  `json:"notification_settings,omitempty"`
	Queue                *QueueSettings            `json:"queue,omitempty"`
	RunAs                *JobRunAs                 `json:"run_as,omitempty"`
	RunName              string                    `json:"run_name,omitempty"`
	Tasks                []SubmitTask              `json:"tasks,omitempty"`
	TimeoutSeconds       int                       `json:"timeout_seconds,omitempty"`
	WebhookNotifications *WebhookNotifications     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func submitRunFromPb(pb *submitRunPb) (*SubmitRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitRun{}
	st.AccessControlList = pb.AccessControlList
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.EmailNotifications = pb.EmailNotifications
	st.Environments = pb.Environments
	st.GitSource = pb.GitSource
	st.Health = pb.Health
	st.IdempotencyToken = pb.IdempotencyToken
	st.NotificationSettings = pb.NotificationSettings
	st.Queue = pb.Queue
	st.RunAs = pb.RunAs
	st.RunName = pb.RunName
	st.Tasks = pb.Tasks
	st.TimeoutSeconds = pb.TimeoutSeconds
	st.WebhookNotifications = pb.WebhookNotifications

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *submitRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st submitRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func submitRunResponseToPb(st *SubmitRunResponse) (*submitRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &submitRunResponsePb{}
	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type submitRunResponsePb struct {
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func submitRunResponseFromPb(pb *submitRunResponsePb) (*SubmitRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitRunResponse{}
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *submitRunResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st submitRunResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func submitTaskToPb(st *SubmitTask) (*submitTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &submitTaskPb{}
	pb.CleanRoomsNotebookTask = st.CleanRoomsNotebookTask
	pb.ConditionTask = st.ConditionTask
	pb.DashboardTask = st.DashboardTask
	pb.DbtTask = st.DbtTask
	pb.DependsOn = st.DependsOn
	pb.Description = st.Description
	pb.EmailNotifications = st.EmailNotifications
	pb.EnvironmentKey = st.EnvironmentKey
	pb.ExistingClusterId = st.ExistingClusterId
	pb.ForEachTask = st.ForEachTask
	pb.GenAiComputeTask = st.GenAiComputeTask
	pb.Health = st.Health
	pb.Libraries = st.Libraries
	pb.NewCluster = st.NewCluster
	pb.NotebookTask = st.NotebookTask
	pb.NotificationSettings = st.NotificationSettings
	pb.PipelineTask = st.PipelineTask
	pb.PowerBiTask = st.PowerBiTask
	pb.PythonWheelTask = st.PythonWheelTask
	pb.RunIf = st.RunIf
	pb.RunJobTask = st.RunJobTask
	pb.SparkJarTask = st.SparkJarTask
	pb.SparkPythonTask = st.SparkPythonTask
	pb.SparkSubmitTask = st.SparkSubmitTask
	pb.SqlTask = st.SqlTask
	pb.TaskKey = st.TaskKey
	pb.TimeoutSeconds = st.TimeoutSeconds
	pb.WebhookNotifications = st.WebhookNotifications

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type submitTaskPb struct {
	CleanRoomsNotebookTask *CleanRoomsNotebookTask   `json:"clean_rooms_notebook_task,omitempty"`
	ConditionTask          *ConditionTask            `json:"condition_task,omitempty"`
	DashboardTask          *DashboardTask            `json:"dashboard_task,omitempty"`
	DbtTask                *DbtTask                  `json:"dbt_task,omitempty"`
	DependsOn              []TaskDependency          `json:"depends_on,omitempty"`
	Description            string                    `json:"description,omitempty"`
	EmailNotifications     *JobEmailNotifications    `json:"email_notifications,omitempty"`
	EnvironmentKey         string                    `json:"environment_key,omitempty"`
	ExistingClusterId      string                    `json:"existing_cluster_id,omitempty"`
	ForEachTask            *ForEachTask              `json:"for_each_task,omitempty"`
	GenAiComputeTask       *GenAiComputeTask         `json:"gen_ai_compute_task,omitempty"`
	Health                 *JobsHealthRules          `json:"health,omitempty"`
	Libraries              []compute.Library         `json:"libraries,omitempty"`
	NewCluster             *compute.ClusterSpec      `json:"new_cluster,omitempty"`
	NotebookTask           *NotebookTask             `json:"notebook_task,omitempty"`
	NotificationSettings   *TaskNotificationSettings `json:"notification_settings,omitempty"`
	PipelineTask           *PipelineTask             `json:"pipeline_task,omitempty"`
	PowerBiTask            *PowerBiTask              `json:"power_bi_task,omitempty"`
	PythonWheelTask        *PythonWheelTask          `json:"python_wheel_task,omitempty"`
	RunIf                  RunIf                     `json:"run_if,omitempty"`
	RunJobTask             *RunJobTask               `json:"run_job_task,omitempty"`
	SparkJarTask           *SparkJarTask             `json:"spark_jar_task,omitempty"`
	SparkPythonTask        *SparkPythonTask          `json:"spark_python_task,omitempty"`
	SparkSubmitTask        *SparkSubmitTask          `json:"spark_submit_task,omitempty"`
	SqlTask                *SqlTask                  `json:"sql_task,omitempty"`
	TaskKey                string                    `json:"task_key"`
	TimeoutSeconds         int                       `json:"timeout_seconds,omitempty"`
	WebhookNotifications   *WebhookNotifications     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func submitTaskFromPb(pb *submitTaskPb) (*SubmitTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitTask{}
	st.CleanRoomsNotebookTask = pb.CleanRoomsNotebookTask
	st.ConditionTask = pb.ConditionTask
	st.DashboardTask = pb.DashboardTask
	st.DbtTask = pb.DbtTask
	st.DependsOn = pb.DependsOn
	st.Description = pb.Description
	st.EmailNotifications = pb.EmailNotifications
	st.EnvironmentKey = pb.EnvironmentKey
	st.ExistingClusterId = pb.ExistingClusterId
	st.ForEachTask = pb.ForEachTask
	st.GenAiComputeTask = pb.GenAiComputeTask
	st.Health = pb.Health
	st.Libraries = pb.Libraries
	st.NewCluster = pb.NewCluster
	st.NotebookTask = pb.NotebookTask
	st.NotificationSettings = pb.NotificationSettings
	st.PipelineTask = pb.PipelineTask
	st.PowerBiTask = pb.PowerBiTask
	st.PythonWheelTask = pb.PythonWheelTask
	st.RunIf = pb.RunIf
	st.RunJobTask = pb.RunJobTask
	st.SparkJarTask = pb.SparkJarTask
	st.SparkPythonTask = pb.SparkPythonTask
	st.SparkSubmitTask = pb.SparkSubmitTask
	st.SqlTask = pb.SqlTask
	st.TaskKey = pb.TaskKey
	st.TimeoutSeconds = pb.TimeoutSeconds
	st.WebhookNotifications = pb.WebhookNotifications

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *submitTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st submitTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func subscriptionToPb(st *Subscription) (*subscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionPb{}
	pb.CustomSubject = st.CustomSubject
	pb.Paused = st.Paused
	pb.Subscribers = st.Subscribers

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type subscriptionPb struct {
	CustomSubject string                   `json:"custom_subject,omitempty"`
	Paused        bool                     `json:"paused,omitempty"`
	Subscribers   []SubscriptionSubscriber `json:"subscribers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func subscriptionFromPb(pb *subscriptionPb) (*Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscription{}
	st.CustomSubject = pb.CustomSubject
	st.Paused = pb.Paused
	st.Subscribers = pb.Subscribers

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *subscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st subscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func subscriptionSubscriberToPb(st *SubscriptionSubscriber) (*subscriptionSubscriberPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionSubscriberPb{}
	pb.DestinationId = st.DestinationId
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type subscriptionSubscriberPb struct {
	DestinationId string `json:"destination_id,omitempty"`
	UserName      string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func subscriptionSubscriberFromPb(pb *subscriptionSubscriberPb) (*SubscriptionSubscriber, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriber{}
	st.DestinationId = pb.DestinationId
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *subscriptionSubscriberPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st subscriptionSubscriberPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tableUpdateTriggerConfigurationToPb(st *TableUpdateTriggerConfiguration) (*tableUpdateTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableUpdateTriggerConfigurationPb{}
	pb.Condition = st.Condition
	pb.MinTimeBetweenTriggersSeconds = st.MinTimeBetweenTriggersSeconds
	pb.TableNames = st.TableNames
	pb.WaitAfterLastChangeSeconds = st.WaitAfterLastChangeSeconds

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableUpdateTriggerConfigurationPb struct {
	Condition                     Condition `json:"condition,omitempty"`
	MinTimeBetweenTriggersSeconds int       `json:"min_time_between_triggers_seconds,omitempty"`
	TableNames                    []string  `json:"table_names,omitempty"`
	WaitAfterLastChangeSeconds    int       `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableUpdateTriggerConfigurationFromPb(pb *tableUpdateTriggerConfigurationPb) (*TableUpdateTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableUpdateTriggerConfiguration{}
	st.Condition = pb.Condition
	st.MinTimeBetweenTriggersSeconds = pb.MinTimeBetweenTriggersSeconds
	st.TableNames = pb.TableNames
	st.WaitAfterLastChangeSeconds = pb.WaitAfterLastChangeSeconds

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableUpdateTriggerConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableUpdateTriggerConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func taskToPb(st *Task) (*taskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskPb{}
	pb.CleanRoomsNotebookTask = st.CleanRoomsNotebookTask
	pb.ConditionTask = st.ConditionTask
	pb.DashboardTask = st.DashboardTask
	pb.DbtTask = st.DbtTask
	pb.DependsOn = st.DependsOn
	pb.Description = st.Description
	pb.DisableAutoOptimization = st.DisableAutoOptimization
	pb.EmailNotifications = st.EmailNotifications
	pb.EnvironmentKey = st.EnvironmentKey
	pb.ExistingClusterId = st.ExistingClusterId
	pb.ForEachTask = st.ForEachTask
	pb.GenAiComputeTask = st.GenAiComputeTask
	pb.Health = st.Health
	pb.JobClusterKey = st.JobClusterKey
	pb.Libraries = st.Libraries
	pb.MaxRetries = st.MaxRetries
	pb.MinRetryIntervalMillis = st.MinRetryIntervalMillis
	pb.NewCluster = st.NewCluster
	pb.NotebookTask = st.NotebookTask
	pb.NotificationSettings = st.NotificationSettings
	pb.PipelineTask = st.PipelineTask
	pb.PowerBiTask = st.PowerBiTask
	pb.PythonWheelTask = st.PythonWheelTask
	pb.RetryOnTimeout = st.RetryOnTimeout
	pb.RunIf = st.RunIf
	pb.RunJobTask = st.RunJobTask
	pb.SparkJarTask = st.SparkJarTask
	pb.SparkPythonTask = st.SparkPythonTask
	pb.SparkSubmitTask = st.SparkSubmitTask
	pb.SqlTask = st.SqlTask
	pb.TaskKey = st.TaskKey
	pb.TimeoutSeconds = st.TimeoutSeconds
	pb.WebhookNotifications = st.WebhookNotifications

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type taskPb struct {
	CleanRoomsNotebookTask  *CleanRoomsNotebookTask   `json:"clean_rooms_notebook_task,omitempty"`
	ConditionTask           *ConditionTask            `json:"condition_task,omitempty"`
	DashboardTask           *DashboardTask            `json:"dashboard_task,omitempty"`
	DbtTask                 *DbtTask                  `json:"dbt_task,omitempty"`
	DependsOn               []TaskDependency          `json:"depends_on,omitempty"`
	Description             string                    `json:"description,omitempty"`
	DisableAutoOptimization bool                      `json:"disable_auto_optimization,omitempty"`
	EmailNotifications      *TaskEmailNotifications   `json:"email_notifications,omitempty"`
	EnvironmentKey          string                    `json:"environment_key,omitempty"`
	ExistingClusterId       string                    `json:"existing_cluster_id,omitempty"`
	ForEachTask             *ForEachTask              `json:"for_each_task,omitempty"`
	GenAiComputeTask        *GenAiComputeTask         `json:"gen_ai_compute_task,omitempty"`
	Health                  *JobsHealthRules          `json:"health,omitempty"`
	JobClusterKey           string                    `json:"job_cluster_key,omitempty"`
	Libraries               []compute.Library         `json:"libraries,omitempty"`
	MaxRetries              int                       `json:"max_retries,omitempty"`
	MinRetryIntervalMillis  int                       `json:"min_retry_interval_millis,omitempty"`
	NewCluster              *compute.ClusterSpec      `json:"new_cluster,omitempty"`
	NotebookTask            *NotebookTask             `json:"notebook_task,omitempty"`
	NotificationSettings    *TaskNotificationSettings `json:"notification_settings,omitempty"`
	PipelineTask            *PipelineTask             `json:"pipeline_task,omitempty"`
	PowerBiTask             *PowerBiTask              `json:"power_bi_task,omitempty"`
	PythonWheelTask         *PythonWheelTask          `json:"python_wheel_task,omitempty"`
	RetryOnTimeout          bool                      `json:"retry_on_timeout,omitempty"`
	RunIf                   RunIf                     `json:"run_if,omitempty"`
	RunJobTask              *RunJobTask               `json:"run_job_task,omitempty"`
	SparkJarTask            *SparkJarTask             `json:"spark_jar_task,omitempty"`
	SparkPythonTask         *SparkPythonTask          `json:"spark_python_task,omitempty"`
	SparkSubmitTask         *SparkSubmitTask          `json:"spark_submit_task,omitempty"`
	SqlTask                 *SqlTask                  `json:"sql_task,omitempty"`
	TaskKey                 string                    `json:"task_key"`
	TimeoutSeconds          int                       `json:"timeout_seconds,omitempty"`
	WebhookNotifications    *WebhookNotifications     `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskFromPb(pb *taskPb) (*Task, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Task{}
	st.CleanRoomsNotebookTask = pb.CleanRoomsNotebookTask
	st.ConditionTask = pb.ConditionTask
	st.DashboardTask = pb.DashboardTask
	st.DbtTask = pb.DbtTask
	st.DependsOn = pb.DependsOn
	st.Description = pb.Description
	st.DisableAutoOptimization = pb.DisableAutoOptimization
	st.EmailNotifications = pb.EmailNotifications
	st.EnvironmentKey = pb.EnvironmentKey
	st.ExistingClusterId = pb.ExistingClusterId
	st.ForEachTask = pb.ForEachTask
	st.GenAiComputeTask = pb.GenAiComputeTask
	st.Health = pb.Health
	st.JobClusterKey = pb.JobClusterKey
	st.Libraries = pb.Libraries
	st.MaxRetries = pb.MaxRetries
	st.MinRetryIntervalMillis = pb.MinRetryIntervalMillis
	st.NewCluster = pb.NewCluster
	st.NotebookTask = pb.NotebookTask
	st.NotificationSettings = pb.NotificationSettings
	st.PipelineTask = pb.PipelineTask
	st.PowerBiTask = pb.PowerBiTask
	st.PythonWheelTask = pb.PythonWheelTask
	st.RetryOnTimeout = pb.RetryOnTimeout
	st.RunIf = pb.RunIf
	st.RunJobTask = pb.RunJobTask
	st.SparkJarTask = pb.SparkJarTask
	st.SparkPythonTask = pb.SparkPythonTask
	st.SparkSubmitTask = pb.SparkSubmitTask
	st.SqlTask = pb.SqlTask
	st.TaskKey = pb.TaskKey
	st.TimeoutSeconds = pb.TimeoutSeconds
	st.WebhookNotifications = pb.WebhookNotifications

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func taskDependencyToPb(st *TaskDependency) (*taskDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskDependencyPb{}
	pb.Outcome = st.Outcome
	pb.TaskKey = st.TaskKey

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type taskDependencyPb struct {
	Outcome string `json:"outcome,omitempty"`
	TaskKey string `json:"task_key"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskDependencyFromPb(pb *taskDependencyPb) (*TaskDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskDependency{}
	st.Outcome = pb.Outcome
	st.TaskKey = pb.TaskKey

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func taskEmailNotificationsToPb(st *TaskEmailNotifications) (*taskEmailNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskEmailNotificationsPb{}
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns
	pb.OnDurationWarningThresholdExceeded = st.OnDurationWarningThresholdExceeded
	pb.OnFailure = st.OnFailure
	pb.OnStart = st.OnStart
	pb.OnStreamingBacklogExceeded = st.OnStreamingBacklogExceeded
	pb.OnSuccess = st.OnSuccess

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type taskEmailNotificationsPb struct {
	NoAlertForSkippedRuns              bool     `json:"no_alert_for_skipped_runs,omitempty"`
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	OnFailure                          []string `json:"on_failure,omitempty"`
	OnStart                            []string `json:"on_start,omitempty"`
	OnStreamingBacklogExceeded         []string `json:"on_streaming_backlog_exceeded,omitempty"`
	OnSuccess                          []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskEmailNotificationsFromPb(pb *taskEmailNotificationsPb) (*TaskEmailNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskEmailNotifications{}
	st.NoAlertForSkippedRuns = pb.NoAlertForSkippedRuns
	st.OnDurationWarningThresholdExceeded = pb.OnDurationWarningThresholdExceeded
	st.OnFailure = pb.OnFailure
	st.OnStart = pb.OnStart
	st.OnStreamingBacklogExceeded = pb.OnStreamingBacklogExceeded
	st.OnSuccess = pb.OnSuccess

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskEmailNotificationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskEmailNotificationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func taskNotificationSettingsToPb(st *TaskNotificationSettings) (*taskNotificationSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskNotificationSettingsPb{}
	pb.AlertOnLastAttempt = st.AlertOnLastAttempt
	pb.NoAlertForCanceledRuns = st.NoAlertForCanceledRuns
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type taskNotificationSettingsPb struct {
	AlertOnLastAttempt     bool `json:"alert_on_last_attempt,omitempty"`
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	NoAlertForSkippedRuns  bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskNotificationSettingsFromPb(pb *taskNotificationSettingsPb) (*TaskNotificationSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskNotificationSettings{}
	st.AlertOnLastAttempt = pb.AlertOnLastAttempt
	st.NoAlertForCanceledRuns = pb.NoAlertForCanceledRuns
	st.NoAlertForSkippedRuns = pb.NoAlertForSkippedRuns

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskNotificationSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskNotificationSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func terminationDetailsToPb(st *TerminationDetails) (*terminationDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &terminationDetailsPb{}
	pb.Code = st.Code
	pb.Message = st.Message
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type terminationDetailsPb struct {
	Code    TerminationCodeCode `json:"code,omitempty"`
	Message string              `json:"message,omitempty"`
	Type    TerminationTypeType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func terminationDetailsFromPb(pb *terminationDetailsPb) (*TerminationDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TerminationDetails{}
	st.Code = pb.Code
	st.Message = pb.Message
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *terminationDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st terminationDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func triggerInfoToPb(st *TriggerInfo) (*triggerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggerInfoPb{}
	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type triggerInfoPb struct {
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func triggerInfoFromPb(pb *triggerInfoPb) (*TriggerInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerInfo{}
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *triggerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st triggerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func triggerSettingsToPb(st *TriggerSettings) (*triggerSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggerSettingsPb{}
	pb.FileArrival = st.FileArrival
	pb.PauseStatus = st.PauseStatus
	pb.Periodic = st.Periodic
	pb.Table = st.Table
	pb.TableUpdate = st.TableUpdate

	return pb, nil
}

type triggerSettingsPb struct {
	FileArrival *FileArrivalTriggerConfiguration `json:"file_arrival,omitempty"`
	PauseStatus PauseStatus                      `json:"pause_status,omitempty"`
	Periodic    *PeriodicTriggerConfiguration    `json:"periodic,omitempty"`
	Table       *TableUpdateTriggerConfiguration `json:"table,omitempty"`
	TableUpdate *TableUpdateTriggerConfiguration `json:"table_update,omitempty"`
}

func triggerSettingsFromPb(pb *triggerSettingsPb) (*TriggerSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerSettings{}
	st.FileArrival = pb.FileArrival
	st.PauseStatus = pb.PauseStatus
	st.Periodic = pb.Periodic
	st.Table = pb.Table
	st.TableUpdate = pb.TableUpdate

	return st, nil
}

func updateJobToPb(st *UpdateJob) (*updateJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateJobPb{}
	pb.FieldsToRemove = st.FieldsToRemove
	pb.JobId = st.JobId
	pb.NewSettings = st.NewSettings

	return pb, nil
}

type updateJobPb struct {
	FieldsToRemove []string     `json:"fields_to_remove,omitempty"`
	JobId          int64        `json:"job_id"`
	NewSettings    *JobSettings `json:"new_settings,omitempty"`
}

func updateJobFromPb(pb *updateJobPb) (*UpdateJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateJob{}
	st.FieldsToRemove = pb.FieldsToRemove
	st.JobId = pb.JobId
	st.NewSettings = pb.NewSettings

	return st, nil
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
}

type updateResponsePb struct {
}

func updateResponseFromPb(pb *updateResponsePb) (*UpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResponse{}

	return st, nil
}

func viewItemToPb(st *ViewItem) (*viewItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &viewItemPb{}
	pb.Content = st.Content
	pb.Name = st.Name
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type viewItemPb struct {
	Content string   `json:"content,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    ViewType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func viewItemFromPb(pb *viewItemPb) (*ViewItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ViewItem{}
	st.Content = pb.Content
	st.Name = pb.Name
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *viewItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st viewItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func webhookToPb(st *Webhook) (*webhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &webhookPb{}
	pb.Id = st.Id

	return pb, nil
}

type webhookPb struct {
	Id string `json:"id"`
}

func webhookFromPb(pb *webhookPb) (*Webhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Webhook{}
	st.Id = pb.Id

	return st, nil
}

func webhookNotificationsToPb(st *WebhookNotifications) (*webhookNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &webhookNotificationsPb{}
	pb.OnDurationWarningThresholdExceeded = st.OnDurationWarningThresholdExceeded
	pb.OnFailure = st.OnFailure
	pb.OnStart = st.OnStart
	pb.OnStreamingBacklogExceeded = st.OnStreamingBacklogExceeded
	pb.OnSuccess = st.OnSuccess

	return pb, nil
}

type webhookNotificationsPb struct {
	OnDurationWarningThresholdExceeded []Webhook `json:"on_duration_warning_threshold_exceeded,omitempty"`
	OnFailure                          []Webhook `json:"on_failure,omitempty"`
	OnStart                            []Webhook `json:"on_start,omitempty"`
	OnStreamingBacklogExceeded         []Webhook `json:"on_streaming_backlog_exceeded,omitempty"`
	OnSuccess                          []Webhook `json:"on_success,omitempty"`
}

func webhookNotificationsFromPb(pb *webhookNotificationsPb) (*WebhookNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WebhookNotifications{}
	st.OnDurationWarningThresholdExceeded = pb.OnDurationWarningThresholdExceeded
	st.OnFailure = pb.OnFailure
	st.OnStart = pb.OnStart
	st.OnStreamingBacklogExceeded = pb.OnStreamingBacklogExceeded
	st.OnSuccess = pb.OnSuccess

	return st, nil
}

func widgetErrorDetailToPb(st *WidgetErrorDetail) (*widgetErrorDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetErrorDetailPb{}
	pb.Message = st.Message

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type widgetErrorDetailPb struct {
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func widgetErrorDetailFromPb(pb *widgetErrorDetailPb) (*WidgetErrorDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WidgetErrorDetail{}
	st.Message = pb.Message

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *widgetErrorDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st widgetErrorDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
