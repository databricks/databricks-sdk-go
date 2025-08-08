// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
	"github.com/databricks/databricks-sdk-go/service/jobs/jobspb"
)

type AuthenticationMethod string

const AuthenticationMethodOauth AuthenticationMethod = `OAUTH`

const AuthenticationMethodPat AuthenticationMethod = `PAT`

// String representation for [fmt.Print]
func (f *AuthenticationMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationMethod) Set(v string) error {
	switch v {
	case `OAUTH`, `PAT`:
		*f = AuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OAUTH", "PAT"`, v)
	}
}

// Values returns all possible values for AuthenticationMethod.
//
// There is no guarantee on the order of the values in the slice.
func (f *AuthenticationMethod) Values() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodOauth,
		AuthenticationMethodPat,
	}
}

// Type always returns AuthenticationMethod to satisfy [pflag.Value] interface
func (f *AuthenticationMethod) Type() string {
	return "AuthenticationMethod"
}

func AuthenticationMethodToPb(st *AuthenticationMethod) (*jobspb.AuthenticationMethodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.AuthenticationMethodPb(*st)
	return &pb, nil
}

func AuthenticationMethodFromPb(pb *jobspb.AuthenticationMethodPb) (*AuthenticationMethod, error) {
	if pb == nil {
		return nil, nil
	}
	st := AuthenticationMethod(*pb)
	return &st, nil
}

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	// Wire name: 'created_time'
	CreatedTime int64 ``
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string ``
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool ``
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId int64 ``
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	// Wire name: 'settings'
	Settings *JobSettings ``
	// State of the trigger associated with the job.
	// Wire name: 'trigger_state'
	TriggerState    *TriggerStateProto ``
	ForceSendFields []string           `tf:"-"`
}

func (st BaseJob) MarshalJSON() ([]byte, error) {
	pb, err := BaseJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BaseJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.BaseJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BaseJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BaseJobToPb(st *BaseJob) (*jobspb.BaseJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.BaseJobPb{}
	pb.CreatedTime = st.CreatedTime
	pb.CreatorUserName = st.CreatorUserName
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	pb.HasMore = st.HasMore
	pb.JobId = st.JobId
	settingsPb, err := JobSettingsToPb(st.Settings)
	if err != nil {
		return nil, err
	}
	if settingsPb != nil {
		pb.Settings = settingsPb
	}
	triggerStatePb, err := TriggerStateProtoToPb(st.TriggerState)
	if err != nil {
		return nil, err
	}
	if triggerStatePb != nil {
		pb.TriggerState = triggerStatePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func BaseJobFromPb(pb *jobspb.BaseJobPb) (*BaseJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseJob{}
	st.CreatedTime = pb.CreatedTime
	st.CreatorUserName = pb.CreatorUserName
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.HasMore = pb.HasMore
	st.JobId = pb.JobId
	settingsField, err := JobSettingsFromPb(pb.Settings)
	if err != nil {
		return nil, err
	}
	if settingsField != nil {
		st.Settings = settingsField
	}
	triggerStateField, err := TriggerStateProtoFromPb(pb.TriggerState)
	if err != nil {
		return nil, err
	}
	if triggerStateField != nil {
		st.TriggerState = triggerStateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type BaseRun struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	// Wire name: 'attempt_number'
	AttemptNumber int ``
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	// Wire name: 'cleanup_duration'
	CleanupDuration int64 ``
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	// Wire name: 'cluster_instance'
	ClusterInstance *ClusterInstance ``
	// A snapshot of the job’s cluster specification when this run was
	// created.
	// Wire name: 'cluster_spec'
	ClusterSpec *ClusterSpec ``
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// Description of the run
	// Wire name: 'description'
	Description string ``
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget ``
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	// Wire name: 'end_time'
	EndTime int64 ``
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	// Wire name: 'execution_duration'
	ExecutionDuration int64 ``
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource ``
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool ``
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster ``
	// The canonical identifier of the job that contains this run.
	// Wire name: 'job_id'
	JobId int64 ``
	// Job-level parameters used in the run
	// Wire name: 'job_parameters'
	JobParameters []JobParameter ``
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	// Wire name: 'job_run_id'
	JobRunId int64 ``
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	// Wire name: 'number_in_job'
	NumberInJob int64 ``
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	// Wire name: 'original_attempt_run_id'
	OriginalAttemptRunId int64 ``
	// The parameters used for this run.
	// Wire name: 'overriding_parameters'
	OverridingParameters *RunParameters ``
	// The time in milliseconds that the run has spent in the queue.
	// Wire name: 'queue_duration'
	QueueDuration int64 ``
	// The repair history of the run.
	// Wire name: 'repair_history'
	RepairHistory []RepairHistoryItem ``
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	// Wire name: 'run_duration'
	RunDuration int64 ``
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'run_id'
	RunId int64 ``
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'run_name'
	RunName string ``
	// The URL to the detail page of the run.
	// Wire name: 'run_page_url'
	RunPageUrl string ``

	// Wire name: 'run_type'
	RunType RunType ``
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	// Wire name: 'schedule'
	Schedule *CronSchedule ``
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	// Wire name: 'setup_duration'
	SetupDuration int64 ``
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	// Wire name: 'start_time'
	StartTime int64 ``
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState ``

	// Wire name: 'status'
	Status *RunStatus ``
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	// Wire name: 'tasks'
	Tasks []RunTask ``

	// Wire name: 'trigger'
	Trigger TriggerType ``

	// Wire name: 'trigger_info'
	TriggerInfo     *TriggerInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (st BaseRun) MarshalJSON() ([]byte, error) {
	pb, err := BaseRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BaseRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.BaseRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BaseRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BaseRunToPb(st *BaseRun) (*jobspb.BaseRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.BaseRunPb{}
	pb.AttemptNumber = st.AttemptNumber
	pb.CleanupDuration = st.CleanupDuration
	clusterInstancePb, err := ClusterInstanceToPb(st.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstancePb != nil {
		pb.ClusterInstance = clusterInstancePb
	}
	clusterSpecPb, err := ClusterSpecToPb(st.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecPb != nil {
		pb.ClusterSpec = clusterSpecPb
	}
	pb.CreatorUserName = st.CreatorUserName
	pb.Description = st.Description
	effectivePerformanceTargetPb, err := PerformanceTargetToPb(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}
	pb.EndTime = st.EndTime
	pb.ExecutionDuration = st.ExecutionDuration
	gitSourcePb, err := GitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}
	pb.HasMore = st.HasMore

	var jobClustersPb []jobspb.JobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := JobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb
	pb.JobId = st.JobId

	var jobParametersPb []jobspb.JobParameterPb
	for _, item := range st.JobParameters {
		itemPb, err := JobParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb = append(jobParametersPb, *itemPb)
		}
	}
	pb.JobParameters = jobParametersPb
	pb.JobRunId = st.JobRunId
	pb.NumberInJob = st.NumberInJob
	pb.OriginalAttemptRunId = st.OriginalAttemptRunId
	overridingParametersPb, err := RunParametersToPb(st.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersPb != nil {
		pb.OverridingParameters = overridingParametersPb
	}
	pb.QueueDuration = st.QueueDuration

	var repairHistoryPb []jobspb.RepairHistoryItemPb
	for _, item := range st.RepairHistory {
		itemPb, err := RepairHistoryItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repairHistoryPb = append(repairHistoryPb, *itemPb)
		}
	}
	pb.RepairHistory = repairHistoryPb
	pb.RunDuration = st.RunDuration
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunPageUrl = st.RunPageUrl
	runTypePb, err := RunTypeToPb(&st.RunType)
	if err != nil {
		return nil, err
	}
	if runTypePb != nil {
		pb.RunType = *runTypePb
	}
	schedulePb, err := CronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.SetupDuration = st.SetupDuration
	pb.StartTime = st.StartTime
	statePb, err := RunStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	statusPb, err := RunStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	var tasksPb []jobspb.RunTaskPb
	for _, item := range st.Tasks {
		itemPb, err := RunTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb
	triggerPb, err := TriggerTypeToPb(&st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = *triggerPb
	}
	triggerInfoPb, err := TriggerInfoToPb(st.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoPb != nil {
		pb.TriggerInfo = triggerInfoPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func BaseRunFromPb(pb *jobspb.BaseRunPb) (*BaseRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseRun{}
	st.AttemptNumber = pb.AttemptNumber
	st.CleanupDuration = pb.CleanupDuration
	clusterInstanceField, err := ClusterInstanceFromPb(pb.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstanceField != nil {
		st.ClusterInstance = clusterInstanceField
	}
	clusterSpecField, err := ClusterSpecFromPb(pb.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecField != nil {
		st.ClusterSpec = clusterSpecField
	}
	st.CreatorUserName = pb.CreatorUserName
	st.Description = pb.Description
	effectivePerformanceTargetField, err := PerformanceTargetFromPb(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	st.EndTime = pb.EndTime
	st.ExecutionDuration = pb.ExecutionDuration
	gitSourceField, err := GitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	st.HasMore = pb.HasMore

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := JobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	st.JobId = pb.JobId

	var jobParametersField []JobParameter
	for _, itemPb := range pb.JobParameters {
		item, err := JobParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField = append(jobParametersField, *item)
		}
	}
	st.JobParameters = jobParametersField
	st.JobRunId = pb.JobRunId
	st.NumberInJob = pb.NumberInJob
	st.OriginalAttemptRunId = pb.OriginalAttemptRunId
	overridingParametersField, err := RunParametersFromPb(pb.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersField != nil {
		st.OverridingParameters = overridingParametersField
	}
	st.QueueDuration = pb.QueueDuration

	var repairHistoryField []RepairHistoryItem
	for _, itemPb := range pb.RepairHistory {
		item, err := RepairHistoryItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repairHistoryField = append(repairHistoryField, *item)
		}
	}
	st.RepairHistory = repairHistoryField
	st.RunDuration = pb.RunDuration
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunPageUrl = pb.RunPageUrl
	runTypeField, err := RunTypeFromPb(&pb.RunType)
	if err != nil {
		return nil, err
	}
	if runTypeField != nil {
		st.RunType = *runTypeField
	}
	scheduleField, err := CronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SetupDuration = pb.SetupDuration
	st.StartTime = pb.StartTime
	stateField, err := RunStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := RunStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	var tasksField []RunTask
	for _, itemPb := range pb.Tasks {
		item, err := RunTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	triggerField, err := TriggerTypeFromPb(&pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = *triggerField
	}
	triggerInfoField, err := TriggerInfoFromPb(pb.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoField != nil {
		st.TriggerInfo = triggerInfoField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	// Wire name: 'all_queued_runs'
	AllQueuedRuns bool ``
	// The canonical identifier of the job to cancel all runs of.
	// Wire name: 'job_id'
	JobId           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st CancelAllRuns) MarshalJSON() ([]byte, error) {
	pb, err := CancelAllRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CancelAllRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CancelAllRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CancelAllRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CancelAllRunsToPb(st *CancelAllRuns) (*jobspb.CancelAllRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CancelAllRunsPb{}
	pb.AllQueuedRuns = st.AllQueuedRuns
	pb.JobId = st.JobId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CancelAllRunsFromPb(pb *jobspb.CancelAllRunsPb) (*CancelAllRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelAllRuns{}
	st.AllQueuedRuns = pb.AllQueuedRuns
	st.JobId = pb.JobId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CancelRun struct {
	// This field is required.
	// Wire name: 'run_id'
	RunId int64 ``
}

func (st CancelRun) MarshalJSON() ([]byte, error) {
	pb, err := CancelRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CancelRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CancelRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CancelRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CancelRunToPb(st *CancelRun) (*jobspb.CancelRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CancelRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func CancelRunFromPb(pb *jobspb.CancelRunPb) (*CancelRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRun{}
	st.RunId = pb.RunId

	return st, nil
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to remove coupling with jobs API definition
type CleanRoomTaskRunLifeCycleState string

const CleanRoomTaskRunLifeCycleStateBlocked CleanRoomTaskRunLifeCycleState = `BLOCKED`

const CleanRoomTaskRunLifeCycleStateInternalError CleanRoomTaskRunLifeCycleState = `INTERNAL_ERROR`

const CleanRoomTaskRunLifeCycleStatePending CleanRoomTaskRunLifeCycleState = `PENDING`

const CleanRoomTaskRunLifeCycleStateQueued CleanRoomTaskRunLifeCycleState = `QUEUED`

const CleanRoomTaskRunLifeCycleStateRunning CleanRoomTaskRunLifeCycleState = `RUNNING`

const CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified CleanRoomTaskRunLifeCycleState = `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`

const CleanRoomTaskRunLifeCycleStateSkipped CleanRoomTaskRunLifeCycleState = `SKIPPED`

const CleanRoomTaskRunLifeCycleStateTerminated CleanRoomTaskRunLifeCycleState = `TERMINATED`

const CleanRoomTaskRunLifeCycleStateTerminating CleanRoomTaskRunLifeCycleState = `TERMINATING`

const CleanRoomTaskRunLifeCycleStateWaitingForRetry CleanRoomTaskRunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = CleanRoomTaskRunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "RUN_LIFE_CYCLE_STATE_UNSPECIFIED", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Values returns all possible values for CleanRoomTaskRunLifeCycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomTaskRunLifeCycleState) Values() []CleanRoomTaskRunLifeCycleState {
	return []CleanRoomTaskRunLifeCycleState{
		CleanRoomTaskRunLifeCycleStateBlocked,
		CleanRoomTaskRunLifeCycleStateInternalError,
		CleanRoomTaskRunLifeCycleStatePending,
		CleanRoomTaskRunLifeCycleStateQueued,
		CleanRoomTaskRunLifeCycleStateRunning,
		CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified,
		CleanRoomTaskRunLifeCycleStateSkipped,
		CleanRoomTaskRunLifeCycleStateTerminated,
		CleanRoomTaskRunLifeCycleStateTerminating,
		CleanRoomTaskRunLifeCycleStateWaitingForRetry,
	}
}

// Type always returns CleanRoomTaskRunLifeCycleState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunLifeCycleState) Type() string {
	return "CleanRoomTaskRunLifeCycleState"
}

func CleanRoomTaskRunLifeCycleStateToPb(st *CleanRoomTaskRunLifeCycleState) (*jobspb.CleanRoomTaskRunLifeCycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.CleanRoomTaskRunLifeCycleStatePb(*st)
	return &pb, nil
}

func CleanRoomTaskRunLifeCycleStateFromPb(pb *jobspb.CleanRoomTaskRunLifeCycleStatePb) (*CleanRoomTaskRunLifeCycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomTaskRunLifeCycleState(*pb)
	return &st, nil
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to avoid cyclic dependency.
type CleanRoomTaskRunResultState string

const CleanRoomTaskRunResultStateCanceled CleanRoomTaskRunResultState = `CANCELED`

const CleanRoomTaskRunResultStateDisabled CleanRoomTaskRunResultState = `DISABLED`

const CleanRoomTaskRunResultStateEvicted CleanRoomTaskRunResultState = `EVICTED`

const CleanRoomTaskRunResultStateExcluded CleanRoomTaskRunResultState = `EXCLUDED`

const CleanRoomTaskRunResultStateFailed CleanRoomTaskRunResultState = `FAILED`

const CleanRoomTaskRunResultStateMaximumConcurrentRunsReached CleanRoomTaskRunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

const CleanRoomTaskRunResultStateRunResultStateUnspecified CleanRoomTaskRunResultState = `RUN_RESULT_STATE_UNSPECIFIED`

const CleanRoomTaskRunResultStateSuccess CleanRoomTaskRunResultState = `SUCCESS`

const CleanRoomTaskRunResultStateSuccessWithFailures CleanRoomTaskRunResultState = `SUCCESS_WITH_FAILURES`

const CleanRoomTaskRunResultStateTimedout CleanRoomTaskRunResultState = `TIMEDOUT`

const CleanRoomTaskRunResultStateUpstreamCanceled CleanRoomTaskRunResultState = `UPSTREAM_CANCELED`

const CleanRoomTaskRunResultStateUpstreamEvicted CleanRoomTaskRunResultState = `UPSTREAM_EVICTED`

const CleanRoomTaskRunResultStateUpstreamFailed CleanRoomTaskRunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EVICTED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `RUN_RESULT_STATE_UNSPECIFIED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_EVICTED`, `UPSTREAM_FAILED`:
		*f = CleanRoomTaskRunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EVICTED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "RUN_RESULT_STATE_UNSPECIFIED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_EVICTED", "UPSTREAM_FAILED"`, v)
	}
}

// Values returns all possible values for CleanRoomTaskRunResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CleanRoomTaskRunResultState) Values() []CleanRoomTaskRunResultState {
	return []CleanRoomTaskRunResultState{
		CleanRoomTaskRunResultStateCanceled,
		CleanRoomTaskRunResultStateDisabled,
		CleanRoomTaskRunResultStateEvicted,
		CleanRoomTaskRunResultStateExcluded,
		CleanRoomTaskRunResultStateFailed,
		CleanRoomTaskRunResultStateMaximumConcurrentRunsReached,
		CleanRoomTaskRunResultStateRunResultStateUnspecified,
		CleanRoomTaskRunResultStateSuccess,
		CleanRoomTaskRunResultStateSuccessWithFailures,
		CleanRoomTaskRunResultStateTimedout,
		CleanRoomTaskRunResultStateUpstreamCanceled,
		CleanRoomTaskRunResultStateUpstreamEvicted,
		CleanRoomTaskRunResultStateUpstreamFailed,
	}
}

// Type always returns CleanRoomTaskRunResultState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunResultState) Type() string {
	return "CleanRoomTaskRunResultState"
}

func CleanRoomTaskRunResultStateToPb(st *CleanRoomTaskRunResultState) (*jobspb.CleanRoomTaskRunResultStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.CleanRoomTaskRunResultStatePb(*st)
	return &pb, nil
}

func CleanRoomTaskRunResultStateFromPb(pb *jobspb.CleanRoomTaskRunResultStatePb) (*CleanRoomTaskRunResultState, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomTaskRunResultState(*pb)
	return &st, nil
}

// Stores the run state of the clean rooms notebook task.
type CleanRoomTaskRunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	// Wire name: 'life_cycle_state'
	LifeCycleState CleanRoomTaskRunLifeCycleState ``
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	// Wire name: 'result_state'
	ResultState CleanRoomTaskRunResultState ``
}

func (st CleanRoomTaskRunState) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomTaskRunStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomTaskRunState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CleanRoomTaskRunStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomTaskRunStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomTaskRunStateToPb(st *CleanRoomTaskRunState) (*jobspb.CleanRoomTaskRunStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CleanRoomTaskRunStatePb{}
	lifeCycleStatePb, err := CleanRoomTaskRunLifeCycleStateToPb(&st.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStatePb != nil {
		pb.LifeCycleState = *lifeCycleStatePb
	}
	resultStatePb, err := CleanRoomTaskRunResultStateToPb(&st.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStatePb != nil {
		pb.ResultState = *resultStatePb
	}

	return pb, nil
}

func CleanRoomTaskRunStateFromPb(pb *jobspb.CleanRoomTaskRunStatePb) (*CleanRoomTaskRunState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomTaskRunState{}
	lifeCycleStateField, err := CleanRoomTaskRunLifeCycleStateFromPb(&pb.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStateField != nil {
		st.LifeCycleState = *lifeCycleStateField
	}
	resultStateField, err := CleanRoomTaskRunResultStateFromPb(&pb.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStateField != nil {
		st.ResultState = *resultStateField
	}

	return st, nil
}

type CleanRoomsNotebookTask struct {
	// The clean room that the notebook belongs to.
	// Wire name: 'clean_room_name'
	CleanRoomName string ``
	// Checksum to validate the freshness of the notebook resource (i.e. the
	// notebook being run is the latest version). It can be fetched by calling
	// the :method:cleanroomassets/get API.
	// Wire name: 'etag'
	Etag string ``
	// Base parameters to be used for the clean room notebook job.
	// Wire name: 'notebook_base_parameters'
	NotebookBaseParameters map[string]string ``
	// Name of the notebook being run.
	// Wire name: 'notebook_name'
	NotebookName    string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CleanRoomsNotebookTask) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomsNotebookTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomsNotebookTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CleanRoomsNotebookTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomsNotebookTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomsNotebookTaskToPb(st *CleanRoomsNotebookTask) (*jobspb.CleanRoomsNotebookTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CleanRoomsNotebookTaskPb{}
	pb.CleanRoomName = st.CleanRoomName
	pb.Etag = st.Etag
	pb.NotebookBaseParameters = st.NotebookBaseParameters
	pb.NotebookName = st.NotebookName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CleanRoomsNotebookTaskFromPb(pb *jobspb.CleanRoomsNotebookTaskPb) (*CleanRoomsNotebookTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomsNotebookTask{}
	st.CleanRoomName = pb.CleanRoomName
	st.Etag = pb.Etag
	st.NotebookBaseParameters = pb.NotebookBaseParameters
	st.NotebookName = pb.NotebookName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput struct {
	// The run state of the clean rooms notebook task.
	// Wire name: 'clean_room_job_run_state'
	CleanRoomJobRunState *CleanRoomTaskRunState ``
	// The notebook output for the clean room run
	// Wire name: 'notebook_output'
	NotebookOutput *NotebookOutput ``
	// Information on how to access the output schema for the clean room run
	// Wire name: 'output_schema_info'
	OutputSchemaInfo *OutputSchemaInfo ``
}

func (st CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(st *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) (*jobspb.CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb{}
	cleanRoomJobRunStatePb, err := CleanRoomTaskRunStateToPb(st.CleanRoomJobRunState)
	if err != nil {
		return nil, err
	}
	if cleanRoomJobRunStatePb != nil {
		pb.CleanRoomJobRunState = cleanRoomJobRunStatePb
	}
	notebookOutputPb, err := NotebookOutputToPb(st.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputPb != nil {
		pb.NotebookOutput = notebookOutputPb
	}
	outputSchemaInfoPb, err := OutputSchemaInfoToPb(st.OutputSchemaInfo)
	if err != nil {
		return nil, err
	}
	if outputSchemaInfoPb != nil {
		pb.OutputSchemaInfo = outputSchemaInfoPb
	}

	return pb, nil
}

func CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb *jobspb.CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb) (*CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput{}
	cleanRoomJobRunStateField, err := CleanRoomTaskRunStateFromPb(pb.CleanRoomJobRunState)
	if err != nil {
		return nil, err
	}
	if cleanRoomJobRunStateField != nil {
		st.CleanRoomJobRunState = cleanRoomJobRunStateField
	}
	notebookOutputField, err := NotebookOutputFromPb(pb.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputField != nil {
		st.NotebookOutput = notebookOutputField
	}
	outputSchemaInfoField, err := OutputSchemaInfoFromPb(pb.OutputSchemaInfo)
	if err != nil {
		return nil, err
	}
	if outputSchemaInfoField != nil {
		st.OutputSchemaInfo = outputSchemaInfoField
	}

	return st, nil
}

type ClusterInstance struct {
	// The canonical identifier for the cluster used by a run. This field is
	// always available for runs on existing clusters. For runs on new clusters,
	// it becomes available once the cluster is created. This value can be used
	// to view logs by browsing to `/#setting/sparkui/$cluster_id/driver-logs`.
	// The logs continue to be available after the run completes.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	// Wire name: 'spark_context_id'
	SparkContextId  string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterInstance) MarshalJSON() ([]byte, error) {
	pb, err := ClusterInstanceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterInstance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ClusterInstancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterInstanceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterInstanceToPb(st *ClusterInstance) (*jobspb.ClusterInstancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ClusterInstancePb{}
	pb.ClusterId = st.ClusterId
	pb.SparkContextId = st.SparkContextId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterInstanceFromPb(pb *jobspb.ClusterInstancePb) (*ClusterInstance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterInstance{}
	st.ClusterId = pb.ClusterId
	st.SparkContextId = pb.SparkContextId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string ``
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	// Wire name: 'job_cluster_key'
	JobClusterKey string ``
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library ``
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster      *compute.ClusterSpec ``
	ForceSendFields []string             `tf:"-"`
}

func (st ClusterSpec) MarshalJSON() ([]byte, error) {
	pb, err := ClusterSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ClusterSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterSpecToPb(st *ClusterSpec) (*jobspb.ClusterSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ClusterSpecPb{}
	pb.ExistingClusterId = st.ExistingClusterId
	pb.JobClusterKey = st.JobClusterKey

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterSpecFromPb(pb *jobspb.ClusterSpecPb) (*ClusterSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSpec{}
	st.ExistingClusterId = pb.ExistingClusterId
	st.JobClusterKey = pb.JobClusterKey

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ComputeConfig struct {
	// IDof the GPU pool to use.
	// Wire name: 'gpu_node_pool_id'
	GpuNodePoolId string ``
	// GPU type.
	// Wire name: 'gpu_type'
	GpuType string ``
	// Number of GPUs.
	// Wire name: 'num_gpus'
	NumGpus         int      ``
	ForceSendFields []string `tf:"-"`
}

func (st ComputeConfig) MarshalJSON() ([]byte, error) {
	pb, err := ComputeConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComputeConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ComputeConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComputeConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComputeConfigToPb(st *ComputeConfig) (*jobspb.ComputeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ComputeConfigPb{}
	pb.GpuNodePoolId = st.GpuNodePoolId
	pb.GpuType = st.GpuType
	pb.NumGpus = st.NumGpus

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComputeConfigFromPb(pb *jobspb.ComputeConfigPb) (*ComputeConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComputeConfig{}
	st.GpuNodePoolId = pb.GpuNodePoolId
	st.GpuType = pb.GpuType
	st.NumGpus = pb.NumGpus

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Condition string

const ConditionAllUpdated Condition = `ALL_UPDATED`

const ConditionAnyUpdated Condition = `ANY_UPDATED`

// String representation for [fmt.Print]
func (f *Condition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Condition) Set(v string) error {
	switch v {
	case `ALL_UPDATED`, `ANY_UPDATED`:
		*f = Condition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_UPDATED", "ANY_UPDATED"`, v)
	}
}

// Values returns all possible values for Condition.
//
// There is no guarantee on the order of the values in the slice.
func (f *Condition) Values() []Condition {
	return []Condition{
		ConditionAllUpdated,
		ConditionAnyUpdated,
	}
}

// Type always returns Condition to satisfy [pflag.Value] interface
func (f *Condition) Type() string {
	return "Condition"
}

func ConditionToPb(st *Condition) (*jobspb.ConditionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.ConditionPb(*st)
	return &pb, nil
}

func ConditionFromPb(pb *jobspb.ConditionPb) (*Condition, error) {
	if pb == nil {
		return nil, nil
	}
	st := Condition(*pb)
	return &st, nil
}

type ConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	// Wire name: 'left'
	Left string ``
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	// Wire name: 'op'
	Op ConditionTaskOp ``
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	// Wire name: 'right'
	Right string ``
}

func (st ConditionTask) MarshalJSON() ([]byte, error) {
	pb, err := ConditionTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ConditionTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ConditionTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ConditionTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ConditionTaskToPb(st *ConditionTask) (*jobspb.ConditionTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ConditionTaskPb{}
	pb.Left = st.Left
	opPb, err := ConditionTaskOpToPb(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}
	pb.Right = st.Right

	return pb, nil
}

func ConditionTaskFromPb(pb *jobspb.ConditionTaskPb) (*ConditionTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConditionTask{}
	st.Left = pb.Left
	opField, err := ConditionTaskOpFromPb(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	st.Right = pb.Right

	return st, nil
}

// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
// operands. This means that `“12.0” == “12”` will evaluate to `false`.
// * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`
// operators perform numeric comparison of their operands. `“12.0” >=
// “12”` will evaluate to `true`, `“10.0” >= “12”` will evaluate to
// `false`.
//
// The boolean comparison to task values can be implemented with operators
// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it will
// be serialized to `“true”` or `“false”` for the comparison.
type ConditionTaskOp string

const ConditionTaskOpEqualTo ConditionTaskOp = `EQUAL_TO`

const ConditionTaskOpGreaterThan ConditionTaskOp = `GREATER_THAN`

const ConditionTaskOpGreaterThanOrEqual ConditionTaskOp = `GREATER_THAN_OR_EQUAL`

const ConditionTaskOpLessThan ConditionTaskOp = `LESS_THAN`

const ConditionTaskOpLessThanOrEqual ConditionTaskOp = `LESS_THAN_OR_EQUAL`

const ConditionTaskOpNotEqual ConditionTaskOp = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *ConditionTaskOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConditionTaskOp) Set(v string) error {
	switch v {
	case `EQUAL_TO`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = ConditionTaskOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL_TO", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Values returns all possible values for ConditionTaskOp.
//
// There is no guarantee on the order of the values in the slice.
func (f *ConditionTaskOp) Values() []ConditionTaskOp {
	return []ConditionTaskOp{
		ConditionTaskOpEqualTo,
		ConditionTaskOpGreaterThan,
		ConditionTaskOpGreaterThanOrEqual,
		ConditionTaskOpLessThan,
		ConditionTaskOpLessThanOrEqual,
		ConditionTaskOpNotEqual,
	}
}

// Type always returns ConditionTaskOp to satisfy [pflag.Value] interface
func (f *ConditionTaskOp) Type() string {
	return "ConditionTaskOp"
}

func ConditionTaskOpToPb(st *ConditionTaskOp) (*jobspb.ConditionTaskOpPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.ConditionTaskOpPb(*st)
	return &pb, nil
}

func ConditionTaskOpFromPb(pb *jobspb.ConditionTaskOpPb) (*ConditionTaskOp, error) {
	if pb == nil {
		return nil, nil
	}
	st := ConditionTaskOp(*pb)
	return &st, nil
}

type Continuous struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	// Wire name: 'pause_status'
	PauseStatus PauseStatus ``
}

func (st Continuous) MarshalJSON() ([]byte, error) {
	pb, err := ContinuousToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Continuous) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ContinuousPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ContinuousFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ContinuousToPb(st *Continuous) (*jobspb.ContinuousPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ContinuousPb{}
	pauseStatusPb, err := PauseStatusToPb(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}

	return pb, nil
}

func ContinuousFromPb(pb *jobspb.ContinuousPb) (*Continuous, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Continuous{}
	pauseStatusField, err := PauseStatusFromPb(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}

	return st, nil
}

type CreateJob struct {
	// List of permissions to set on the job.
	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlRequest ``
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	// Wire name: 'continuous'
	Continuous *Continuous ``
	// Deployment information for jobs managed by external sources.
	// Wire name: 'deployment'
	Deployment *JobDeployment ``
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	// Wire name: 'description'
	Description string ``
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	// Wire name: 'edit_mode'
	EditMode JobEditMode ``
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications ``
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	// Wire name: 'environments'
	Environments []JobEnvironment ``
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	// Wire name: 'format'
	Format Format ``
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource ``

	// Wire name: 'health'
	Health *JobsHealthRules ``
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster ``
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	// Wire name: 'max_concurrent_runs'
	MaxConcurrentRuns int ``
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'name'
	Name string ``
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	// Wire name: 'notification_settings'
	NotificationSettings *JobNotificationSettings ``
	// Job-level parameter definitions
	// Wire name: 'parameters'
	Parameters []JobParameterDefinition ``
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget ``
	// The queue settings of the job.
	// Wire name: 'queue'
	Queue *QueueSettings ``
	// The user or service principal that the job runs as, if specified in the
	// request. This field indicates the explicit configuration of `run_as` for
	// the job. To find the value in all cases, explicit or implicit, use
	// `run_as_user_name`.
	// Wire name: 'run_as'
	RunAs *JobRunAs ``
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	// Wire name: 'schedule'
	Schedule *CronSchedule ``
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	// Wire name: 'tags'
	Tags map[string]string ``
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	// Wire name: 'tasks'
	Tasks []Task ``
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int ``
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	// Wire name: 'trigger'
	Trigger *TriggerSettings ``
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications ``
	ForceSendFields      []string              `tf:"-"`
}

func (st CreateJob) MarshalJSON() ([]byte, error) {
	pb, err := CreateJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CreateJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateJobToPb(st *CreateJob) (*jobspb.CreateJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CreateJobPb{}

	var accessControlListPb []jobspb.JobAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := JobAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.BudgetPolicyId = st.BudgetPolicyId
	continuousPb, err := ContinuousToPb(st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = continuousPb
	}
	deploymentPb, err := JobDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}
	pb.Description = st.Description
	editModePb, err := JobEditModeToPb(&st.EditMode)
	if err != nil {
		return nil, err
	}
	if editModePb != nil {
		pb.EditMode = *editModePb
	}
	emailNotificationsPb, err := JobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	var environmentsPb []jobspb.JobEnvironmentPb
	for _, item := range st.Environments {
		itemPb, err := JobEnvironmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentsPb = append(environmentsPb, *itemPb)
		}
	}
	pb.Environments = environmentsPb
	formatPb, err := FormatToPb(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}
	gitSourcePb, err := GitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}
	healthPb, err := JobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	var jobClustersPb []jobspb.JobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := JobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb
	pb.MaxConcurrentRuns = st.MaxConcurrentRuns
	pb.Name = st.Name
	notificationSettingsPb, err := JobNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	var parametersPb []jobspb.JobParameterDefinitionPb
	for _, item := range st.Parameters {
		itemPb, err := JobParameterDefinitionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb
	performanceTargetPb, err := PerformanceTargetToPb(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}
	queuePb, err := QueueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}
	runAsPb, err := JobRunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}
	schedulePb, err := CronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.Tags = st.Tags

	var tasksPb []jobspb.TaskPb
	for _, item := range st.Tasks {
		itemPb, err := TaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb
	pb.TimeoutSeconds = st.TimeoutSeconds
	triggerPb, err := TriggerSettingsToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}
	webhookNotificationsPb, err := WebhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateJobFromPb(pb *jobspb.CreateJobPb) (*CreateJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateJob{}

	var accessControlListField []JobAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := JobAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.BudgetPolicyId = pb.BudgetPolicyId
	continuousField, err := ContinuousFromPb(pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = continuousField
	}
	deploymentField, err := JobDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Description = pb.Description
	editModeField, err := JobEditModeFromPb(&pb.EditMode)
	if err != nil {
		return nil, err
	}
	if editModeField != nil {
		st.EditMode = *editModeField
	}
	emailNotificationsField, err := JobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}

	var environmentsField []JobEnvironment
	for _, itemPb := range pb.Environments {
		item, err := JobEnvironmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentsField = append(environmentsField, *item)
		}
	}
	st.Environments = environmentsField
	formatField, err := FormatFromPb(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	gitSourceField, err := GitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	healthField, err := JobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := JobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	st.MaxConcurrentRuns = pb.MaxConcurrentRuns
	st.Name = pb.Name
	notificationSettingsField, err := JobNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}

	var parametersField []JobParameterDefinition
	for _, itemPb := range pb.Parameters {
		item, err := JobParameterDefinitionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	performanceTargetField, err := PerformanceTargetFromPb(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	queueField, err := QueueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	runAsField, err := JobRunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	scheduleField, err := CronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.Tags = pb.Tags

	var tasksField []Task
	for _, itemPb := range pb.Tasks {
		item, err := TaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	st.TimeoutSeconds = pb.TimeoutSeconds
	triggerField, err := TriggerSettingsFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}
	webhookNotificationsField, err := WebhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Job was created successfully
type CreateResponse struct {
	// The canonical identifier for the newly created job.
	// Wire name: 'job_id'
	JobId           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CreateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateResponseToPb(st *CreateResponse) (*jobspb.CreateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CreateResponsePb{}
	pb.JobId = st.JobId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateResponseFromPb(pb *jobspb.CreateResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	st.JobId = pb.JobId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus PauseStatus ``
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string ``
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	// Wire name: 'timezone_id'
	TimezoneId string ``
}

func (st CronSchedule) MarshalJSON() ([]byte, error) {
	pb, err := CronScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CronSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.CronSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CronScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CronScheduleToPb(st *CronSchedule) (*jobspb.CronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.CronSchedulePb{}
	pauseStatusPb, err := PauseStatusToPb(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}
	pb.QuartzCronExpression = st.QuartzCronExpression
	pb.TimezoneId = st.TimezoneId

	return pb, nil
}

func CronScheduleFromPb(pb *jobspb.CronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	pauseStatusField, err := PauseStatusFromPb(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	st.QuartzCronExpression = pb.QuartzCronExpression
	st.TimezoneId = pb.TimezoneId

	return st, nil
}

type DashboardPageSnapshot struct {

	// Wire name: 'page_display_name'
	PageDisplayName string ``

	// Wire name: 'widget_error_details'
	WidgetErrorDetails []WidgetErrorDetail ``
	ForceSendFields    []string            `tf:"-"`
}

func (st DashboardPageSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := DashboardPageSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DashboardPageSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DashboardPageSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DashboardPageSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DashboardPageSnapshotToPb(st *DashboardPageSnapshot) (*jobspb.DashboardPageSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DashboardPageSnapshotPb{}
	pb.PageDisplayName = st.PageDisplayName

	var widgetErrorDetailsPb []jobspb.WidgetErrorDetailPb
	for _, item := range st.WidgetErrorDetails {
		itemPb, err := WidgetErrorDetailToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			widgetErrorDetailsPb = append(widgetErrorDetailsPb, *itemPb)
		}
	}
	pb.WidgetErrorDetails = widgetErrorDetailsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DashboardPageSnapshotFromPb(pb *jobspb.DashboardPageSnapshotPb) (*DashboardPageSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardPageSnapshot{}
	st.PageDisplayName = pb.PageDisplayName

	var widgetErrorDetailsField []WidgetErrorDetail
	for _, itemPb := range pb.WidgetErrorDetails {
		item, err := WidgetErrorDetailFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			widgetErrorDetailsField = append(widgetErrorDetailsField, *item)
		}
	}
	st.WidgetErrorDetails = widgetErrorDetailsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Configures the Lakeview Dashboard job task type.
type DashboardTask struct {
	// The identifier of the dashboard to refresh.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// Optional: subscription configuration for sending the dashboard snapshot.
	// Wire name: 'subscription'
	Subscription *Subscription ``
	// Optional: The warehouse id to execute the dashboard with for the
	// schedule. If not specified, the default warehouse of the dashboard will
	// be used.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st DashboardTask) MarshalJSON() ([]byte, error) {
	pb, err := DashboardTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DashboardTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DashboardTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DashboardTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DashboardTaskToPb(st *DashboardTask) (*jobspb.DashboardTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DashboardTaskPb{}
	pb.DashboardId = st.DashboardId
	subscriptionPb, err := SubscriptionToPb(st.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionPb != nil {
		pb.Subscription = subscriptionPb
	}
	pb.WarehouseId = st.WarehouseId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DashboardTaskFromPb(pb *jobspb.DashboardTaskPb) (*DashboardTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardTask{}
	st.DashboardId = pb.DashboardId
	subscriptionField, err := SubscriptionFromPb(pb.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionField != nil {
		st.Subscription = subscriptionField
	}
	st.WarehouseId = pb.WarehouseId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DashboardTaskOutput struct {
	// Should only be populated for manual PDF download jobs.
	// Wire name: 'page_snapshots'
	PageSnapshots []DashboardPageSnapshot ``
}

func (st DashboardTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := DashboardTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DashboardTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DashboardTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DashboardTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DashboardTaskOutputToPb(st *DashboardTaskOutput) (*jobspb.DashboardTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DashboardTaskOutputPb{}

	var pageSnapshotsPb []jobspb.DashboardPageSnapshotPb
	for _, item := range st.PageSnapshots {
		itemPb, err := DashboardPageSnapshotToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pageSnapshotsPb = append(pageSnapshotsPb, *itemPb)
		}
	}
	pb.PageSnapshots = pageSnapshotsPb

	return pb, nil
}

func DashboardTaskOutputFromPb(pb *jobspb.DashboardTaskOutputPb) (*DashboardTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardTaskOutput{}

	var pageSnapshotsField []DashboardPageSnapshot
	for _, itemPb := range pb.PageSnapshots {
		item, err := DashboardPageSnapshotFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pageSnapshotsField = append(pageSnapshotsField, *item)
		}
	}
	st.PageSnapshots = pageSnapshotsField

	return st, nil
}

// Format of response retrieved from dbt Cloud, for inclusion in output
// Deprecated in favor of DbtPlatformJobRunStep
type DbtCloudJobRunStep struct {
	// Orders the steps in the job
	// Wire name: 'index'
	Index int ``
	// Output of the step
	// Wire name: 'logs'
	Logs string ``
	// Name of the step in the job
	// Wire name: 'name'
	Name string ``
	// State of the step
	// Wire name: 'status'
	Status          DbtPlatformRunStatus ``
	ForceSendFields []string             `tf:"-"`
}

func (st DbtCloudJobRunStep) MarshalJSON() ([]byte, error) {
	pb, err := DbtCloudJobRunStepToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtCloudJobRunStep) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtCloudJobRunStepPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtCloudJobRunStepFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtCloudJobRunStepToPb(st *DbtCloudJobRunStep) (*jobspb.DbtCloudJobRunStepPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtCloudJobRunStepPb{}
	pb.Index = st.Index
	pb.Logs = st.Logs
	pb.Name = st.Name
	statusPb, err := DbtPlatformRunStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtCloudJobRunStepFromPb(pb *jobspb.DbtCloudJobRunStepPb) (*DbtCloudJobRunStep, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtCloudJobRunStep{}
	st.Index = pb.Index
	st.Logs = pb.Logs
	st.Name = pb.Name
	statusField, err := DbtPlatformRunStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Deprecated in favor of DbtPlatformTask
type DbtCloudTask struct {
	// The resource name of the UC connection that authenticates the dbt Cloud
	// for this task
	// Wire name: 'connection_resource_name'
	ConnectionResourceName string ``
	// Id of the dbt Cloud job to be triggered
	// Wire name: 'dbt_cloud_job_id'
	DbtCloudJobId   int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st DbtCloudTask) MarshalJSON() ([]byte, error) {
	pb, err := DbtCloudTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtCloudTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtCloudTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtCloudTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtCloudTaskToPb(st *DbtCloudTask) (*jobspb.DbtCloudTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtCloudTaskPb{}
	pb.ConnectionResourceName = st.ConnectionResourceName
	pb.DbtCloudJobId = st.DbtCloudJobId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtCloudTaskFromPb(pb *jobspb.DbtCloudTaskPb) (*DbtCloudTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtCloudTask{}
	st.ConnectionResourceName = pb.ConnectionResourceName
	st.DbtCloudJobId = pb.DbtCloudJobId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Deprecated in favor of DbtPlatformTaskOutput
type DbtCloudTaskOutput struct {
	// Id of the job run in dbt Cloud
	// Wire name: 'dbt_cloud_job_run_id'
	DbtCloudJobRunId int64 ``
	// Steps of the job run as received from dbt Cloud
	// Wire name: 'dbt_cloud_job_run_output'
	DbtCloudJobRunOutput []DbtCloudJobRunStep ``
	// Url where full run details can be viewed
	// Wire name: 'dbt_cloud_job_run_url'
	DbtCloudJobRunUrl string   ``
	ForceSendFields   []string `tf:"-"`
}

func (st DbtCloudTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := DbtCloudTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtCloudTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtCloudTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtCloudTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtCloudTaskOutputToPb(st *DbtCloudTaskOutput) (*jobspb.DbtCloudTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtCloudTaskOutputPb{}
	pb.DbtCloudJobRunId = st.DbtCloudJobRunId

	var dbtCloudJobRunOutputPb []jobspb.DbtCloudJobRunStepPb
	for _, item := range st.DbtCloudJobRunOutput {
		itemPb, err := DbtCloudJobRunStepToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dbtCloudJobRunOutputPb = append(dbtCloudJobRunOutputPb, *itemPb)
		}
	}
	pb.DbtCloudJobRunOutput = dbtCloudJobRunOutputPb
	pb.DbtCloudJobRunUrl = st.DbtCloudJobRunUrl

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtCloudTaskOutputFromPb(pb *jobspb.DbtCloudTaskOutputPb) (*DbtCloudTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtCloudTaskOutput{}
	st.DbtCloudJobRunId = pb.DbtCloudJobRunId

	var dbtCloudJobRunOutputField []DbtCloudJobRunStep
	for _, itemPb := range pb.DbtCloudJobRunOutput {
		item, err := DbtCloudJobRunStepFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dbtCloudJobRunOutputField = append(dbtCloudJobRunOutputField, *item)
		}
	}
	st.DbtCloudJobRunOutput = dbtCloudJobRunOutputField
	st.DbtCloudJobRunUrl = pb.DbtCloudJobRunUrl

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	// Wire name: 'artifacts_headers'
	ArtifactsHeaders map[string]string ``
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	// Wire name: 'artifacts_link'
	ArtifactsLink   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st DbtOutput) MarshalJSON() ([]byte, error) {
	pb, err := DbtOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtOutputToPb(st *DbtOutput) (*jobspb.DbtOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtOutputPb{}
	pb.ArtifactsHeaders = st.ArtifactsHeaders
	pb.ArtifactsLink = st.ArtifactsLink

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtOutputFromPb(pb *jobspb.DbtOutputPb) (*DbtOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtOutput{}
	st.ArtifactsHeaders = pb.ArtifactsHeaders
	st.ArtifactsLink = pb.ArtifactsLink

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Format of response retrieved from dbt platform, for inclusion in output
type DbtPlatformJobRunStep struct {
	// Orders the steps in the job
	// Wire name: 'index'
	Index int ``
	// Output of the step
	// Wire name: 'logs'
	Logs string ``
	// Whether the logs of this step have been truncated. If true, the logs has
	// been truncated to 10000 characters.
	// Wire name: 'logs_truncated'
	LogsTruncated bool ``
	// Name of the step in the job
	// Wire name: 'name'
	Name string ``
	// Whether the name of the job has been truncated. If true, the name has
	// been truncated to 100 characters.
	// Wire name: 'name_truncated'
	NameTruncated bool ``
	// State of the step
	// Wire name: 'status'
	Status          DbtPlatformRunStatus ``
	ForceSendFields []string             `tf:"-"`
}

func (st DbtPlatformJobRunStep) MarshalJSON() ([]byte, error) {
	pb, err := DbtPlatformJobRunStepToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtPlatformJobRunStep) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtPlatformJobRunStepPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtPlatformJobRunStepFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtPlatformJobRunStepToPb(st *DbtPlatformJobRunStep) (*jobspb.DbtPlatformJobRunStepPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtPlatformJobRunStepPb{}
	pb.Index = st.Index
	pb.Logs = st.Logs
	pb.LogsTruncated = st.LogsTruncated
	pb.Name = st.Name
	pb.NameTruncated = st.NameTruncated
	statusPb, err := DbtPlatformRunStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtPlatformJobRunStepFromPb(pb *jobspb.DbtPlatformJobRunStepPb) (*DbtPlatformJobRunStep, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtPlatformJobRunStep{}
	st.Index = pb.Index
	st.Logs = pb.Logs
	st.LogsTruncated = pb.LogsTruncated
	st.Name = pb.Name
	st.NameTruncated = pb.NameTruncated
	statusField, err := DbtPlatformRunStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Response enumeration from calling the dbt platform API, for inclusion in
// output
type DbtPlatformRunStatus string

const DbtPlatformRunStatusCancelled DbtPlatformRunStatus = `CANCELLED`

const DbtPlatformRunStatusError DbtPlatformRunStatus = `ERROR`

const DbtPlatformRunStatusQueued DbtPlatformRunStatus = `QUEUED`

const DbtPlatformRunStatusRunning DbtPlatformRunStatus = `RUNNING`

const DbtPlatformRunStatusStarting DbtPlatformRunStatus = `STARTING`

const DbtPlatformRunStatusSuccess DbtPlatformRunStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *DbtPlatformRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DbtPlatformRunStatus) Set(v string) error {
	switch v {
	case `CANCELLED`, `ERROR`, `QUEUED`, `RUNNING`, `STARTING`, `SUCCESS`:
		*f = DbtPlatformRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "ERROR", "QUEUED", "RUNNING", "STARTING", "SUCCESS"`, v)
	}
}

// Values returns all possible values for DbtPlatformRunStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DbtPlatformRunStatus) Values() []DbtPlatformRunStatus {
	return []DbtPlatformRunStatus{
		DbtPlatformRunStatusCancelled,
		DbtPlatformRunStatusError,
		DbtPlatformRunStatusQueued,
		DbtPlatformRunStatusRunning,
		DbtPlatformRunStatusStarting,
		DbtPlatformRunStatusSuccess,
	}
}

// Type always returns DbtPlatformRunStatus to satisfy [pflag.Value] interface
func (f *DbtPlatformRunStatus) Type() string {
	return "DbtPlatformRunStatus"
}

func DbtPlatformRunStatusToPb(st *DbtPlatformRunStatus) (*jobspb.DbtPlatformRunStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.DbtPlatformRunStatusPb(*st)
	return &pb, nil
}

func DbtPlatformRunStatusFromPb(pb *jobspb.DbtPlatformRunStatusPb) (*DbtPlatformRunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := DbtPlatformRunStatus(*pb)
	return &st, nil
}

type DbtPlatformTask struct {
	// The resource name of the UC connection that authenticates the dbt
	// platform for this task
	// Wire name: 'connection_resource_name'
	ConnectionResourceName string ``
	// Id of the dbt platform job to be triggered. Specified as a string for
	// maximum compatibility with clients.
	// Wire name: 'dbt_platform_job_id'
	DbtPlatformJobId string   ``
	ForceSendFields  []string `tf:"-"`
}

func (st DbtPlatformTask) MarshalJSON() ([]byte, error) {
	pb, err := DbtPlatformTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtPlatformTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtPlatformTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtPlatformTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtPlatformTaskToPb(st *DbtPlatformTask) (*jobspb.DbtPlatformTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtPlatformTaskPb{}
	pb.ConnectionResourceName = st.ConnectionResourceName
	pb.DbtPlatformJobId = st.DbtPlatformJobId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtPlatformTaskFromPb(pb *jobspb.DbtPlatformTaskPb) (*DbtPlatformTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtPlatformTask{}
	st.ConnectionResourceName = pb.ConnectionResourceName
	st.DbtPlatformJobId = pb.DbtPlatformJobId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DbtPlatformTaskOutput struct {
	// Id of the job run in dbt platform. Specified as a string for maximum
	// compatibility with clients.
	// Wire name: 'dbt_platform_job_run_id'
	DbtPlatformJobRunId string ``
	// Steps of the job run as received from dbt platform
	// Wire name: 'dbt_platform_job_run_output'
	DbtPlatformJobRunOutput []DbtPlatformJobRunStep ``
	// Url where full run details can be viewed
	// Wire name: 'dbt_platform_job_run_url'
	DbtPlatformJobRunUrl string ``
	// Whether the number of steps in the output has been truncated. If true,
	// the output will contain the first 20 steps of the output.
	// Wire name: 'steps_truncated'
	StepsTruncated  bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st DbtPlatformTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := DbtPlatformTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtPlatformTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtPlatformTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtPlatformTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtPlatformTaskOutputToPb(st *DbtPlatformTaskOutput) (*jobspb.DbtPlatformTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtPlatformTaskOutputPb{}
	pb.DbtPlatformJobRunId = st.DbtPlatformJobRunId

	var dbtPlatformJobRunOutputPb []jobspb.DbtPlatformJobRunStepPb
	for _, item := range st.DbtPlatformJobRunOutput {
		itemPb, err := DbtPlatformJobRunStepToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dbtPlatformJobRunOutputPb = append(dbtPlatformJobRunOutputPb, *itemPb)
		}
	}
	pb.DbtPlatformJobRunOutput = dbtPlatformJobRunOutputPb
	pb.DbtPlatformJobRunUrl = st.DbtPlatformJobRunUrl
	pb.StepsTruncated = st.StepsTruncated

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtPlatformTaskOutputFromPb(pb *jobspb.DbtPlatformTaskOutputPb) (*DbtPlatformTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtPlatformTaskOutput{}
	st.DbtPlatformJobRunId = pb.DbtPlatformJobRunId

	var dbtPlatformJobRunOutputField []DbtPlatformJobRunStep
	for _, itemPb := range pb.DbtPlatformJobRunOutput {
		item, err := DbtPlatformJobRunStepFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dbtPlatformJobRunOutputField = append(dbtPlatformJobRunOutputField, *item)
		}
	}
	st.DbtPlatformJobRunOutput = dbtPlatformJobRunOutputField
	st.DbtPlatformJobRunUrl = pb.DbtPlatformJobRunUrl
	st.StepsTruncated = pb.StepsTruncated

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	// Wire name: 'catalog'
	Catalog string ``
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	// Wire name: 'commands'
	Commands []string ``
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	// Wire name: 'profiles_directory'
	ProfilesDirectory string ``
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	// Wire name: 'project_directory'
	ProjectDirectory string ``
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	// Wire name: 'schema'
	Schema string ``
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	// Wire name: 'source'
	Source Source ``
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st DbtTask) MarshalJSON() ([]byte, error) {
	pb, err := DbtTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbtTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DbtTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbtTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbtTaskToPb(st *DbtTask) (*jobspb.DbtTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DbtTaskPb{}
	pb.Catalog = st.Catalog
	pb.Commands = st.Commands
	pb.ProfilesDirectory = st.ProfilesDirectory
	pb.ProjectDirectory = st.ProjectDirectory
	pb.Schema = st.Schema
	sourcePb, err := SourceToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	pb.WarehouseId = st.WarehouseId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DbtTaskFromPb(pb *jobspb.DbtTaskPb) (*DbtTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtTask{}
	st.Catalog = pb.Catalog
	st.Commands = pb.Commands
	st.ProfilesDirectory = pb.ProfilesDirectory
	st.ProjectDirectory = pb.ProjectDirectory
	st.Schema = pb.Schema
	sourceField, err := SourceFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	st.WarehouseId = pb.WarehouseId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	// Wire name: 'job_id'
	JobId int64 ``
}

func (st DeleteJob) MarshalJSON() ([]byte, error) {
	pb, err := DeleteJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DeleteJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteJobToPb(st *DeleteJob) (*jobspb.DeleteJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DeleteJobPb{}
	pb.JobId = st.JobId

	return pb, nil
}

func DeleteJobFromPb(pb *jobspb.DeleteJobPb) (*DeleteJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteJob{}
	st.JobId = pb.JobId

	return st, nil
}

type DeleteRun struct {
	// ID of the run to delete.
	// Wire name: 'run_id'
	RunId int64 ``
}

func (st DeleteRun) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.DeleteRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteRunToPb(st *DeleteRun) (*jobspb.DeleteRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.DeleteRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func DeleteRunFromPb(pb *jobspb.DeleteRunPb) (*DeleteRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRun{}
	st.RunId = pb.RunId

	return st, nil
}

// Represents a change to the job cluster's settings that would be required for
// the job clusters to become compliant with their policies.
type EnforcePolicyComplianceForJobResponseJobClusterSettingsChange struct {
	// The field where this change would be made, prepended with the job cluster
	// key.
	// Wire name: 'field'
	Field string ``
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	// Wire name: 'new_value'
	NewValue string ``
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	// Wire name: 'previous_value'
	PreviousValue   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) MarshalJSON() ([]byte, error) {
	pb, err := EnforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) (*jobspb.EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb{}
	pb.Field = st.Field
	pb.NewValue = st.NewValue
	pb.PreviousValue = st.PreviousValue

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(pb *jobspb.EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) (*EnforcePolicyComplianceForJobResponseJobClusterSettingsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceForJobResponseJobClusterSettingsChange{}
	st.Field = pb.Field
	st.NewValue = pb.NewValue
	st.PreviousValue = pb.PreviousValue

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EnforcePolicyComplianceRequest struct {
	// The ID of the job you want to enforce policy compliance on.
	// Wire name: 'job_id'
	JobId int64 ``
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	// Wire name: 'validate_only'
	ValidateOnly    bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st EnforcePolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := EnforcePolicyComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnforcePolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.EnforcePolicyComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnforcePolicyComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnforcePolicyComplianceRequestToPb(st *EnforcePolicyComplianceRequest) (*jobspb.EnforcePolicyComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.EnforcePolicyComplianceRequestPb{}
	pb.JobId = st.JobId
	pb.ValidateOnly = st.ValidateOnly

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnforcePolicyComplianceRequestFromPb(pb *jobspb.EnforcePolicyComplianceRequestPb) (*EnforcePolicyComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceRequest{}
	st.JobId = pb.JobId
	st.ValidateOnly = pb.ValidateOnly

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EnforcePolicyComplianceResponse struct {
	// Whether any changes have been made to the job cluster settings for the
	// job to become compliant with its policies.
	// Wire name: 'has_changes'
	HasChanges bool ``
	// A list of job cluster changes that have been made to the job’s cluster
	// settings in order for all job clusters to become compliant with their
	// policies.
	// Wire name: 'job_cluster_changes'
	JobClusterChanges []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange ``
	// Updated job settings after policy enforcement. Policy enforcement only
	// applies to job clusters that are created when running the job (which are
	// specified in new_cluster) and does not apply to existing all-purpose
	// clusters. Updated job settings are derived by applying policy default
	// values to the existing job clusters in order to satisfy policy
	// requirements.
	// Wire name: 'settings'
	Settings        *JobSettings ``
	ForceSendFields []string     `tf:"-"`
}

func (st EnforcePolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := EnforcePolicyComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnforcePolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.EnforcePolicyComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnforcePolicyComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnforcePolicyComplianceResponseToPb(st *EnforcePolicyComplianceResponse) (*jobspb.EnforcePolicyComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.EnforcePolicyComplianceResponsePb{}
	pb.HasChanges = st.HasChanges

	var jobClusterChangesPb []jobspb.EnforcePolicyComplianceForJobResponseJobClusterSettingsChangePb
	for _, item := range st.JobClusterChanges {
		itemPb, err := EnforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClusterChangesPb = append(jobClusterChangesPb, *itemPb)
		}
	}
	pb.JobClusterChanges = jobClusterChangesPb
	settingsPb, err := JobSettingsToPb(st.Settings)
	if err != nil {
		return nil, err
	}
	if settingsPb != nil {
		pb.Settings = settingsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnforcePolicyComplianceResponseFromPb(pb *jobspb.EnforcePolicyComplianceResponsePb) (*EnforcePolicyComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceResponse{}
	st.HasChanges = pb.HasChanges

	var jobClusterChangesField []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange
	for _, itemPb := range pb.JobClusterChanges {
		item, err := EnforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClusterChangesField = append(jobClusterChangesField, *item)
		}
	}
	st.JobClusterChanges = jobClusterChangesField
	settingsField, err := JobSettingsFromPb(pb.Settings)
	if err != nil {
		return nil, err
	}
	if settingsField != nil {
		st.Settings = settingsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Run was exported successfully.
type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	// Wire name: 'views'
	Views []ViewItem ``
}

func (st ExportRunOutput) MarshalJSON() ([]byte, error) {
	pb, err := ExportRunOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExportRunOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ExportRunOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExportRunOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExportRunOutputToPb(st *ExportRunOutput) (*jobspb.ExportRunOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ExportRunOutputPb{}

	var viewsPb []jobspb.ViewItemPb
	for _, item := range st.Views {
		itemPb, err := ViewItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			viewsPb = append(viewsPb, *itemPb)
		}
	}
	pb.Views = viewsPb

	return pb, nil
}

func ExportRunOutputFromPb(pb *jobspb.ExportRunOutputPb) (*ExportRunOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRunOutput{}

	var viewsField []ViewItem
	for _, itemPb := range pb.Views {
		item, err := ViewItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			viewsField = append(viewsField, *item)
		}
	}
	st.Views = viewsField

	return st, nil
}

type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	// Wire name: 'run_id'
	RunId int64 `tf:"-"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	// Wire name: 'views_to_export'
	ViewsToExport ViewsToExport `tf:"-"`
}

func (st ExportRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := ExportRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExportRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ExportRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExportRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExportRunRequestToPb(st *ExportRunRequest) (*jobspb.ExportRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ExportRunRequestPb{}
	pb.RunId = st.RunId
	viewsToExportPb, err := ViewsToExportToPb(&st.ViewsToExport)
	if err != nil {
		return nil, err
	}
	if viewsToExportPb != nil {
		pb.ViewsToExport = *viewsToExportPb
	}

	return pb, nil
}

func ExportRunRequestFromPb(pb *jobspb.ExportRunRequestPb) (*ExportRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRunRequest{}
	st.RunId = pb.RunId
	viewsToExportField, err := ViewsToExportFromPb(&pb.ViewsToExport)
	if err != nil {
		return nil, err
	}
	if viewsToExportField != nil {
		st.ViewsToExport = *viewsToExportField
	}

	return st, nil
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	// Wire name: 'min_time_between_triggers_seconds'
	MinTimeBetweenTriggersSeconds int ``
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	// Wire name: 'url'
	Url string ``
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	// Wire name: 'wait_after_last_change_seconds'
	WaitAfterLastChangeSeconds int      ``
	ForceSendFields            []string `tf:"-"`
}

func (st FileArrivalTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := FileArrivalTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FileArrivalTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.FileArrivalTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FileArrivalTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FileArrivalTriggerConfigurationToPb(st *FileArrivalTriggerConfiguration) (*jobspb.FileArrivalTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.FileArrivalTriggerConfigurationPb{}
	pb.MinTimeBetweenTriggersSeconds = st.MinTimeBetweenTriggersSeconds
	pb.Url = st.Url
	pb.WaitAfterLastChangeSeconds = st.WaitAfterLastChangeSeconds

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FileArrivalTriggerConfigurationFromPb(pb *jobspb.FileArrivalTriggerConfigurationPb) (*FileArrivalTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileArrivalTriggerConfiguration{}
	st.MinTimeBetweenTriggersSeconds = pb.MinTimeBetweenTriggersSeconds
	st.Url = pb.Url
	st.WaitAfterLastChangeSeconds = pb.WaitAfterLastChangeSeconds

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FileArrivalTriggerState struct {
	// Indicates whether the trigger leverages file events to detect file
	// arrivals.
	// Wire name: 'using_file_events'
	UsingFileEvents bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st FileArrivalTriggerState) MarshalJSON() ([]byte, error) {
	pb, err := FileArrivalTriggerStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FileArrivalTriggerState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.FileArrivalTriggerStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FileArrivalTriggerStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FileArrivalTriggerStateToPb(st *FileArrivalTriggerState) (*jobspb.FileArrivalTriggerStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.FileArrivalTriggerStatePb{}
	pb.UsingFileEvents = st.UsingFileEvents

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FileArrivalTriggerStateFromPb(pb *jobspb.FileArrivalTriggerStatePb) (*FileArrivalTriggerState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileArrivalTriggerState{}
	st.UsingFileEvents = pb.UsingFileEvents

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	// Wire name: 'error_message_stats'
	ErrorMessageStats []ForEachTaskErrorMessageStats ``
	// Describes stats of the iteration. Only latest retries are considered.
	// Wire name: 'task_run_stats'
	TaskRunStats *ForEachTaskTaskRunStats ``
}

func (st ForEachStats) MarshalJSON() ([]byte, error) {
	pb, err := ForEachStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ForEachStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ForEachStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ForEachStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ForEachStatsToPb(st *ForEachStats) (*jobspb.ForEachStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ForEachStatsPb{}

	var errorMessageStatsPb []jobspb.ForEachTaskErrorMessageStatsPb
	for _, item := range st.ErrorMessageStats {
		itemPb, err := ForEachTaskErrorMessageStatsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			errorMessageStatsPb = append(errorMessageStatsPb, *itemPb)
		}
	}
	pb.ErrorMessageStats = errorMessageStatsPb
	taskRunStatsPb, err := ForEachTaskTaskRunStatsToPb(st.TaskRunStats)
	if err != nil {
		return nil, err
	}
	if taskRunStatsPb != nil {
		pb.TaskRunStats = taskRunStatsPb
	}

	return pb, nil
}

func ForEachStatsFromPb(pb *jobspb.ForEachStatsPb) (*ForEachStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachStats{}

	var errorMessageStatsField []ForEachTaskErrorMessageStats
	for _, itemPb := range pb.ErrorMessageStats {
		item, err := ForEachTaskErrorMessageStatsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			errorMessageStatsField = append(errorMessageStatsField, *item)
		}
	}
	st.ErrorMessageStats = errorMessageStatsField
	taskRunStatsField, err := ForEachTaskTaskRunStatsFromPb(pb.TaskRunStats)
	if err != nil {
		return nil, err
	}
	if taskRunStatsField != nil {
		st.TaskRunStats = taskRunStatsField
	}

	return st, nil
}

type ForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	// Wire name: 'concurrency'
	Concurrency int ``
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	// Wire name: 'inputs'
	Inputs string ``
	// Configuration for the task that will be run for each element in the array
	// Wire name: 'task'
	Task            Task     ``
	ForceSendFields []string `tf:"-"`
}

func (st ForEachTask) MarshalJSON() ([]byte, error) {
	pb, err := ForEachTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ForEachTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ForEachTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ForEachTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ForEachTaskToPb(st *ForEachTask) (*jobspb.ForEachTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ForEachTaskPb{}
	pb.Concurrency = st.Concurrency
	pb.Inputs = st.Inputs
	taskPb, err := TaskToPb(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ForEachTaskFromPb(pb *jobspb.ForEachTaskPb) (*ForEachTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTask{}
	st.Concurrency = pb.Concurrency
	st.Inputs = pb.Inputs
	taskField, err := TaskFromPb(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	// Wire name: 'count'
	Count int ``
	// Describes the error message occured during the iterations.
	// Wire name: 'error_message'
	ErrorMessage string ``
	// Describes the termination reason for the error message.
	// Wire name: 'termination_category'
	TerminationCategory string   ``
	ForceSendFields     []string `tf:"-"`
}

func (st ForEachTaskErrorMessageStats) MarshalJSON() ([]byte, error) {
	pb, err := ForEachTaskErrorMessageStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ForEachTaskErrorMessageStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ForEachTaskErrorMessageStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ForEachTaskErrorMessageStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ForEachTaskErrorMessageStatsToPb(st *ForEachTaskErrorMessageStats) (*jobspb.ForEachTaskErrorMessageStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ForEachTaskErrorMessageStatsPb{}
	pb.Count = st.Count
	pb.ErrorMessage = st.ErrorMessage
	pb.TerminationCategory = st.TerminationCategory

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ForEachTaskErrorMessageStatsFromPb(pb *jobspb.ForEachTaskErrorMessageStatsPb) (*ForEachTaskErrorMessageStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTaskErrorMessageStats{}
	st.Count = pb.Count
	st.ErrorMessage = pb.ErrorMessage
	st.TerminationCategory = pb.TerminationCategory

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	// Wire name: 'active_iterations'
	ActiveIterations int ``
	// Describes the number of failed and succeeded iteration runs.
	// Wire name: 'completed_iterations'
	CompletedIterations int ``
	// Describes the number of failed iteration runs.
	// Wire name: 'failed_iterations'
	FailedIterations int ``
	// Describes the number of iteration runs that have been scheduled.
	// Wire name: 'scheduled_iterations'
	ScheduledIterations int ``
	// Describes the number of succeeded iteration runs.
	// Wire name: 'succeeded_iterations'
	SucceededIterations int ``
	// Describes the length of the list of items to iterate over.
	// Wire name: 'total_iterations'
	TotalIterations int      ``
	ForceSendFields []string `tf:"-"`
}

func (st ForEachTaskTaskRunStats) MarshalJSON() ([]byte, error) {
	pb, err := ForEachTaskTaskRunStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ForEachTaskTaskRunStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ForEachTaskTaskRunStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ForEachTaskTaskRunStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ForEachTaskTaskRunStatsToPb(st *ForEachTaskTaskRunStats) (*jobspb.ForEachTaskTaskRunStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ForEachTaskTaskRunStatsPb{}
	pb.ActiveIterations = st.ActiveIterations
	pb.CompletedIterations = st.CompletedIterations
	pb.FailedIterations = st.FailedIterations
	pb.ScheduledIterations = st.ScheduledIterations
	pb.SucceededIterations = st.SucceededIterations
	pb.TotalIterations = st.TotalIterations

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ForEachTaskTaskRunStatsFromPb(pb *jobspb.ForEachTaskTaskRunStatsPb) (*ForEachTaskTaskRunStats, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Format string

const FormatMultiTask Format = `MULTI_TASK`

const FormatSingleTask Format = `SINGLE_TASK`

// String representation for [fmt.Print]
func (f *Format) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Format) Set(v string) error {
	switch v {
	case `MULTI_TASK`, `SINGLE_TASK`:
		*f = Format(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MULTI_TASK", "SINGLE_TASK"`, v)
	}
}

// Values returns all possible values for Format.
//
// There is no guarantee on the order of the values in the slice.
func (f *Format) Values() []Format {
	return []Format{
		FormatMultiTask,
		FormatSingleTask,
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

func FormatToPb(st *Format) (*jobspb.FormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.FormatPb(*st)
	return &pb, nil
}

func FormatFromPb(pb *jobspb.FormatPb) (*Format, error) {
	if pb == nil {
		return nil, nil
	}
	st := Format(*pb)
	return &st, nil
}

type GenAiComputeTask struct {
	// Command launcher to run the actual script, e.g. bash, python etc.
	// Wire name: 'command'
	Command string ``

	// Wire name: 'compute'
	Compute *ComputeConfig ``
	// Runtime image
	// Wire name: 'dl_runtime_image'
	DlRuntimeImage string ``
	// Optional string containing the name of the MLflow experiment to log the
	// run to. If name is not found, backend will create the mlflow experiment
	// using the name.
	// Wire name: 'mlflow_experiment_name'
	MlflowExperimentName string ``
	// Optional location type of the training script. When set to `WORKSPACE`,
	// the script will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the script will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`: Script
	// is located in Databricks workspace. * `GIT`: Script is located in cloud
	// Git provider.
	// Wire name: 'source'
	Source Source ``
	// The training script file path to be executed. Cloud file URIs (such as
	// dbfs:/, s3:/, adls:/, gcs:/) and workspace paths are supported. For
	// python files stored in the Databricks workspace, the path must be
	// absolute and begin with `/`. For files stored in a remote repository, the
	// path must be relative. This field is required.
	// Wire name: 'training_script_path'
	TrainingScriptPath string ``
	// Optional string containing model parameters passed to the training script
	// in yaml format. If present, then the content in yaml_parameters_file_path
	// will be ignored.
	// Wire name: 'yaml_parameters'
	YamlParameters string ``
	// Optional path to a YAML file containing model parameters passed to the
	// training script.
	// Wire name: 'yaml_parameters_file_path'
	YamlParametersFilePath string   ``
	ForceSendFields        []string `tf:"-"`
}

func (st GenAiComputeTask) MarshalJSON() ([]byte, error) {
	pb, err := GenAiComputeTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GenAiComputeTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GenAiComputeTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GenAiComputeTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GenAiComputeTaskToPb(st *GenAiComputeTask) (*jobspb.GenAiComputeTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GenAiComputeTaskPb{}
	pb.Command = st.Command
	computePb, err := ComputeConfigToPb(st.Compute)
	if err != nil {
		return nil, err
	}
	if computePb != nil {
		pb.Compute = computePb
	}
	pb.DlRuntimeImage = st.DlRuntimeImage
	pb.MlflowExperimentName = st.MlflowExperimentName
	sourcePb, err := SourceToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	pb.TrainingScriptPath = st.TrainingScriptPath
	pb.YamlParameters = st.YamlParameters
	pb.YamlParametersFilePath = st.YamlParametersFilePath

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GenAiComputeTaskFromPb(pb *jobspb.GenAiComputeTaskPb) (*GenAiComputeTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenAiComputeTask{}
	st.Command = pb.Command
	computeField, err := ComputeConfigFromPb(pb.Compute)
	if err != nil {
		return nil, err
	}
	if computeField != nil {
		st.Compute = computeField
	}
	st.DlRuntimeImage = pb.DlRuntimeImage
	st.MlflowExperimentName = pb.MlflowExperimentName
	sourceField, err := SourceFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	st.TrainingScriptPath = pb.TrainingScriptPath
	st.YamlParameters = pb.YamlParameters
	st.YamlParametersFilePath = pb.YamlParametersFilePath

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	// Wire name: 'job_id'
	JobId string `tf:"-"`
}

func (st GetJobPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetJobPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetJobPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetJobPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetJobPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetJobPermissionLevelsRequestToPb(st *GetJobPermissionLevelsRequest) (*jobspb.GetJobPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetJobPermissionLevelsRequestPb{}
	pb.JobId = st.JobId

	return pb, nil
}

func GetJobPermissionLevelsRequestFromPb(pb *jobspb.GetJobPermissionLevelsRequestPb) (*GetJobPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionLevelsRequest{}
	st.JobId = pb.JobId

	return st, nil
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []JobPermissionsDescription ``
}

func (st GetJobPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetJobPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetJobPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetJobPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetJobPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetJobPermissionLevelsResponseToPb(st *GetJobPermissionLevelsResponse) (*jobspb.GetJobPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetJobPermissionLevelsResponsePb{}

	var permissionLevelsPb []jobspb.JobPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := JobPermissionsDescriptionToPb(&item)
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

func GetJobPermissionLevelsResponseFromPb(pb *jobspb.GetJobPermissionLevelsResponsePb) (*GetJobPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionLevelsResponse{}

	var permissionLevelsField []JobPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := JobPermissionsDescriptionFromPb(&itemPb)
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

type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	// Wire name: 'job_id'
	JobId string `tf:"-"`
}

func (st GetJobPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetJobPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetJobPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetJobPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetJobPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetJobPermissionsRequestToPb(st *GetJobPermissionsRequest) (*jobspb.GetJobPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetJobPermissionsRequestPb{}
	pb.JobId = st.JobId

	return pb, nil
}

func GetJobPermissionsRequestFromPb(pb *jobspb.GetJobPermissionsRequestPb) (*GetJobPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionsRequest{}
	st.JobId = pb.JobId

	return st, nil
}

type GetJobRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	// Wire name: 'job_id'
	JobId int64 `tf:"-"`
	// Use `next_page_token` returned from the previous GetJob response to
	// request the next page of the job's array properties.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GetJobRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetJobRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetJobRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetJobRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetJobRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetJobRequestToPb(st *GetJobRequest) (*jobspb.GetJobRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetJobRequestPb{}
	pb.JobId = st.JobId
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetJobRequestFromPb(pb *jobspb.GetJobRequestPb) (*GetJobRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobRequest{}
	st.JobId = pb.JobId
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetPolicyComplianceRequest struct {
	// The ID of the job whose compliance status you are requesting.
	// Wire name: 'job_id'
	JobId int64 `tf:"-"`
}

func (st GetPolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPolicyComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetPolicyComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPolicyComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPolicyComplianceRequestToPb(st *GetPolicyComplianceRequest) (*jobspb.GetPolicyComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetPolicyComplianceRequestPb{}
	pb.JobId = st.JobId

	return pb, nil
}

func GetPolicyComplianceRequestFromPb(pb *jobspb.GetPolicyComplianceRequestPb) (*GetPolicyComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyComplianceRequest{}
	st.JobId = pb.JobId

	return st, nil
}

type GetPolicyComplianceResponse struct {
	// Whether the job is compliant with its policies or not. Jobs could be out
	// of compliance if a policy they are using was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	// Wire name: 'is_compliant'
	IsCompliant bool ``
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	// Wire name: 'violations'
	Violations      map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (st GetPolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetPolicyComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetPolicyComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPolicyComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPolicyComplianceResponseToPb(st *GetPolicyComplianceResponse) (*jobspb.GetPolicyComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetPolicyComplianceResponsePb{}
	pb.IsCompliant = st.IsCompliant
	pb.Violations = st.Violations

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetPolicyComplianceResponseFromPb(pb *jobspb.GetPolicyComplianceResponsePb) (*GetPolicyComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyComplianceResponse{}
	st.IsCompliant = pb.IsCompliant
	st.Violations = pb.Violations

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetRunOutputRequest struct {
	// The canonical identifier for the run.
	// Wire name: 'run_id'
	RunId int64 `tf:"-"`
}

func (st GetRunOutputRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRunOutputRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRunOutputRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetRunOutputRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRunOutputRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRunOutputRequestToPb(st *GetRunOutputRequest) (*jobspb.GetRunOutputRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetRunOutputRequestPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func GetRunOutputRequestFromPb(pb *jobspb.GetRunOutputRequestPb) (*GetRunOutputRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunOutputRequest{}
	st.RunId = pb.RunId

	return st, nil
}

type GetRunRequest struct {
	// Whether to include the repair history in the response.
	// Wire name: 'include_history'
	IncludeHistory bool `tf:"-"`
	// Whether to include resolved parameter values in the response.
	// Wire name: 'include_resolved_values'
	IncludeResolvedValues bool `tf:"-"`
	// Use `next_page_token` returned from the previous GetRun response to
	// request the next page of the run's array properties.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	// Wire name: 'run_id'
	RunId           int64    `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GetRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GetRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRunRequestToPb(st *GetRunRequest) (*jobspb.GetRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GetRunRequestPb{}
	pb.IncludeHistory = st.IncludeHistory
	pb.IncludeResolvedValues = st.IncludeResolvedValues
	pb.PageToken = st.PageToken
	pb.RunId = st.RunId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetRunRequestFromPb(pb *jobspb.GetRunRequestPb) (*GetRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunRequest{}
	st.IncludeHistory = pb.IncludeHistory
	st.IncludeResolvedValues = pb.IncludeResolvedValues
	st.PageToken = pb.PageToken
	st.RunId = pb.RunId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GitProvider string

const GitProviderAwsCodeCommit GitProvider = `awsCodeCommit`

const GitProviderAzureDevOpsServices GitProvider = `azureDevOpsServices`

const GitProviderBitbucketCloud GitProvider = `bitbucketCloud`

const GitProviderBitbucketServer GitProvider = `bitbucketServer`

const GitProviderGitHub GitProvider = `gitHub`

const GitProviderGitHubEnterprise GitProvider = `gitHubEnterprise`

const GitProviderGitLab GitProvider = `gitLab`

const GitProviderGitLabEnterpriseEdition GitProvider = `gitLabEnterpriseEdition`

// String representation for [fmt.Print]
func (f *GitProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GitProvider) Set(v string) error {
	switch v {
	case `awsCodeCommit`, `azureDevOpsServices`, `bitbucketCloud`, `bitbucketServer`, `gitHub`, `gitHubEnterprise`, `gitLab`, `gitLabEnterpriseEdition`:
		*f = GitProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "awsCodeCommit", "azureDevOpsServices", "bitbucketCloud", "bitbucketServer", "gitHub", "gitHubEnterprise", "gitLab", "gitLabEnterpriseEdition"`, v)
	}
}

// Values returns all possible values for GitProvider.
//
// There is no guarantee on the order of the values in the slice.
func (f *GitProvider) Values() []GitProvider {
	return []GitProvider{
		GitProviderAwsCodeCommit,
		GitProviderAzureDevOpsServices,
		GitProviderBitbucketCloud,
		GitProviderBitbucketServer,
		GitProviderGitHub,
		GitProviderGitHubEnterprise,
		GitProviderGitLab,
		GitProviderGitLabEnterpriseEdition,
	}
}

// Type always returns GitProvider to satisfy [pflag.Value] interface
func (f *GitProvider) Type() string {
	return "GitProvider"
}

func GitProviderToPb(st *GitProvider) (*jobspb.GitProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.GitProviderPb(*st)
	return &pb, nil
}

func GitProviderFromPb(pb *jobspb.GitProviderPb) (*GitProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := GitProvider(*pb)
	return &st, nil
}

// Read-only state of the remote repository at the time the job was run. This
// field is only included on job runs.
type GitSnapshot struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	// Wire name: 'used_commit'
	UsedCommit      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GitSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := GitSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GitSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GitSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GitSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GitSnapshotToPb(st *GitSnapshot) (*jobspb.GitSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GitSnapshotPb{}
	pb.UsedCommit = st.UsedCommit

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GitSnapshotFromPb(pb *jobspb.GitSnapshotPb) (*GitSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GitSnapshot{}
	st.UsedCommit = pb.UsedCommit

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// An optional specification for a remote Git repository containing the source
// code used by tasks. Version-controlled source code is supported by notebook,
// dbt, Python script, and SQL File tasks.
//
// If `git_source` is set, these tasks retrieve the file from the remote
// repository by default. However, this behavior can be overridden by setting
// `source` to `WORKSPACE` on the task.
//
// Note: dbt and SQL File tasks support only version-controlled sources. If dbt
// or SQL File tasks are used, `git_source` must be defined on the job.
type GitSource struct {
	// Name of the branch to be checked out and used by this job. This field
	// cannot be specified in conjunction with git_tag or git_commit.
	// Wire name: 'git_branch'
	GitBranch string ``
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	// Wire name: 'git_commit'
	GitCommit string ``
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	// Wire name: 'git_provider'
	GitProvider GitProvider ``

	// Wire name: 'git_snapshot'
	GitSnapshot *GitSnapshot ``
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	// Wire name: 'git_tag'
	GitTag string ``
	// URL of the repository to be cloned by this job.
	// Wire name: 'git_url'
	GitUrl string ``
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	// Wire name: 'job_source'
	JobSource       *JobSource ``
	ForceSendFields []string   `tf:"-"`
}

func (st GitSource) MarshalJSON() ([]byte, error) {
	pb, err := GitSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GitSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.GitSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GitSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GitSourceToPb(st *GitSource) (*jobspb.GitSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.GitSourcePb{}
	pb.GitBranch = st.GitBranch
	pb.GitCommit = st.GitCommit
	gitProviderPb, err := GitProviderToPb(&st.GitProvider)
	if err != nil {
		return nil, err
	}
	if gitProviderPb != nil {
		pb.GitProvider = *gitProviderPb
	}
	gitSnapshotPb, err := GitSnapshotToPb(st.GitSnapshot)
	if err != nil {
		return nil, err
	}
	if gitSnapshotPb != nil {
		pb.GitSnapshot = gitSnapshotPb
	}
	pb.GitTag = st.GitTag
	pb.GitUrl = st.GitUrl
	jobSourcePb, err := JobSourceToPb(st.JobSource)
	if err != nil {
		return nil, err
	}
	if jobSourcePb != nil {
		pb.JobSource = jobSourcePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GitSourceFromPb(pb *jobspb.GitSourcePb) (*GitSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GitSource{}
	st.GitBranch = pb.GitBranch
	st.GitCommit = pb.GitCommit
	gitProviderField, err := GitProviderFromPb(&pb.GitProvider)
	if err != nil {
		return nil, err
	}
	if gitProviderField != nil {
		st.GitProvider = *gitProviderField
	}
	gitSnapshotField, err := GitSnapshotFromPb(pb.GitSnapshot)
	if err != nil {
		return nil, err
	}
	if gitSnapshotField != nil {
		st.GitSnapshot = gitSnapshotField
	}
	st.GitTag = pb.GitTag
	st.GitUrl = pb.GitUrl
	jobSourceField, err := JobSourceFromPb(pb.JobSource)
	if err != nil {
		return nil, err
	}
	if jobSourceField != nil {
		st.JobSource = jobSourceField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Job was retrieved successfully.
type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	// Wire name: 'created_time'
	CreatedTime int64 ``
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string ``
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool ``
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId int64 ``
	// A token that can be used to list the next page of array properties.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	// Wire name: 'run_as_user_name'
	RunAsUserName string ``
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	// Wire name: 'settings'
	Settings *JobSettings ``
	// State of the trigger associated with the job.
	// Wire name: 'trigger_state'
	TriggerState    *TriggerStateProto ``
	ForceSendFields []string           `tf:"-"`
}

func (st Job) MarshalJSON() ([]byte, error) {
	pb, err := JobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Job) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobToPb(st *Job) (*jobspb.JobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobPb{}
	pb.CreatedTime = st.CreatedTime
	pb.CreatorUserName = st.CreatorUserName
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	pb.HasMore = st.HasMore
	pb.JobId = st.JobId
	pb.NextPageToken = st.NextPageToken
	pb.RunAsUserName = st.RunAsUserName
	settingsPb, err := JobSettingsToPb(st.Settings)
	if err != nil {
		return nil, err
	}
	if settingsPb != nil {
		pb.Settings = settingsPb
	}
	triggerStatePb, err := TriggerStateProtoToPb(st.TriggerState)
	if err != nil {
		return nil, err
	}
	if triggerStatePb != nil {
		pb.TriggerState = triggerStatePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobFromPb(pb *jobspb.JobPb) (*Job, error) {
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
	settingsField, err := JobSettingsFromPb(pb.Settings)
	if err != nil {
		return nil, err
	}
	if settingsField != nil {
		st.Settings = settingsField
	}
	triggerStateField, err := TriggerStateProtoFromPb(pb.TriggerState)
	if err != nil {
		return nil, err
	}
	if triggerStateField != nil {
		st.TriggerState = triggerStateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel JobPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st JobAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := JobAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobAccessControlRequestToPb(st *JobAccessControlRequest) (*jobspb.JobAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := JobPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobAccessControlRequestFromPb(pb *jobspb.JobAccessControlRequestPb) (*JobAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := JobPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []JobPermission ``
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

func (st JobAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := JobAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobAccessControlResponseToPb(st *JobAccessControlResponse) (*jobspb.JobAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobAccessControlResponsePb{}

	var allPermissionsPb []jobspb.JobPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := JobPermissionToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobAccessControlResponseFromPb(pb *jobspb.JobAccessControlResponsePb) (*JobAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobAccessControlResponse{}

	var allPermissionsField []JobPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := JobPermissionFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	// Wire name: 'job_cluster_key'
	JobClusterKey string ``
	// If new_cluster, a description of a cluster that is created for each task.
	// Wire name: 'new_cluster'
	NewCluster compute.ClusterSpec ``
}

func (st JobCluster) MarshalJSON() ([]byte, error) {
	pb, err := JobClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobClusterToPb(st *JobCluster) (*jobspb.JobClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobClusterPb{}
	pb.JobClusterKey = st.JobClusterKey
	newClusterPb, err := compute.ClusterSpecToPb(&st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = *newClusterPb
	}

	return pb, nil
}

func JobClusterFromPb(pb *jobspb.JobClusterPb) (*JobCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobCluster{}
	st.JobClusterKey = pb.JobClusterKey
	newClusterField, err := compute.ClusterSpecFromPb(&pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = *newClusterField
	}

	return st, nil
}

type JobCompliance struct {
	// Whether this job is in compliance with the latest version of its policy.
	// Wire name: 'is_compliant'
	IsCompliant bool ``
	// Canonical unique identifier for a job.
	// Wire name: 'job_id'
	JobId int64 ``
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	// Wire name: 'violations'
	Violations      map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (st JobCompliance) MarshalJSON() ([]byte, error) {
	pb, err := JobComplianceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobCompliance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobCompliancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobComplianceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobComplianceToPb(st *JobCompliance) (*jobspb.JobCompliancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobCompliancePb{}
	pb.IsCompliant = st.IsCompliant
	pb.JobId = st.JobId
	pb.Violations = st.Violations

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobComplianceFromPb(pb *jobspb.JobCompliancePb) (*JobCompliance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobCompliance{}
	st.IsCompliant = pb.IsCompliant
	st.JobId = pb.JobId
	st.Violations = pb.Violations

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobDeployment struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	// Wire name: 'kind'
	Kind JobDeploymentKind ``
	// Path of the file that contains deployment metadata.
	// Wire name: 'metadata_file_path'
	MetadataFilePath string   ``
	ForceSendFields  []string `tf:"-"`
}

func (st JobDeployment) MarshalJSON() ([]byte, error) {
	pb, err := JobDeploymentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobDeployment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobDeploymentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobDeploymentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobDeploymentToPb(st *JobDeployment) (*jobspb.JobDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobDeploymentPb{}
	kindPb, err := JobDeploymentKindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.MetadataFilePath = st.MetadataFilePath

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobDeploymentFromPb(pb *jobspb.JobDeploymentPb) (*JobDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobDeployment{}
	kindField, err := JobDeploymentKindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.MetadataFilePath = pb.MetadataFilePath

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
type JobDeploymentKind string

// The job is managed by Databricks Asset Bundle.
const JobDeploymentKindBundle JobDeploymentKind = `BUNDLE`

// String representation for [fmt.Print]
func (f *JobDeploymentKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobDeploymentKind) Set(v string) error {
	switch v {
	case `BUNDLE`:
		*f = JobDeploymentKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUNDLE"`, v)
	}
}

// Values returns all possible values for JobDeploymentKind.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobDeploymentKind) Values() []JobDeploymentKind {
	return []JobDeploymentKind{
		JobDeploymentKindBundle,
	}
}

// Type always returns JobDeploymentKind to satisfy [pflag.Value] interface
func (f *JobDeploymentKind) Type() string {
	return "JobDeploymentKind"
}

func JobDeploymentKindToPb(st *JobDeploymentKind) (*jobspb.JobDeploymentKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.JobDeploymentKindPb(*st)
	return &pb, nil
}

func JobDeploymentKindFromPb(pb *jobspb.JobDeploymentKindPb) (*JobDeploymentKind, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobDeploymentKind(*pb)
	return &st, nil
}

// Edit mode of the job.
//
// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
// `EDITABLE`: The job is in an editable state and can be modified.
type JobEditMode string

// The job is in an editable state and can be modified.
const JobEditModeEditable JobEditMode = `EDITABLE`

// The job is in a locked UI state and cannot be modified.
const JobEditModeUiLocked JobEditMode = `UI_LOCKED`

// String representation for [fmt.Print]
func (f *JobEditMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobEditMode) Set(v string) error {
	switch v {
	case `EDITABLE`, `UI_LOCKED`:
		*f = JobEditMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EDITABLE", "UI_LOCKED"`, v)
	}
}

// Values returns all possible values for JobEditMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobEditMode) Values() []JobEditMode {
	return []JobEditMode{
		JobEditModeEditable,
		JobEditModeUiLocked,
	}
}

// Type always returns JobEditMode to satisfy [pflag.Value] interface
func (f *JobEditMode) Type() string {
	return "JobEditMode"
}

func JobEditModeToPb(st *JobEditMode) (*jobspb.JobEditModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.JobEditModePb(*st)
	return &pb, nil
}

func JobEditModeFromPb(pb *jobspb.JobEditModePb) (*JobEditMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobEditMode(*pb)
	return &st, nil
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool ``
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	// Wire name: 'on_duration_warning_threshold_exceeded'
	OnDurationWarningThresholdExceeded []string ``
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	// Wire name: 'on_failure'
	OnFailure []string ``
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_start'
	OnStart []string ``
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	// Wire name: 'on_streaming_backlog_exceeded'
	OnStreamingBacklogExceeded []string ``
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_success'
	OnSuccess       []string ``
	ForceSendFields []string `tf:"-"`
}

func (st JobEmailNotifications) MarshalJSON() ([]byte, error) {
	pb, err := JobEmailNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobEmailNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobEmailNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobEmailNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobEmailNotificationsToPb(st *JobEmailNotifications) (*jobspb.JobEmailNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobEmailNotificationsPb{}
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns
	pb.OnDurationWarningThresholdExceeded = st.OnDurationWarningThresholdExceeded
	pb.OnFailure = st.OnFailure
	pb.OnStart = st.OnStart
	pb.OnStreamingBacklogExceeded = st.OnStreamingBacklogExceeded
	pb.OnSuccess = st.OnSuccess

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobEmailNotificationsFromPb(pb *jobspb.JobEmailNotificationsPb) (*JobEmailNotifications, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobEnvironment struct {
	// The key of an environment. It has to be unique within a job.
	// Wire name: 'environment_key'
	EnvironmentKey string ``

	// Wire name: 'spec'
	Spec *compute.Environment ``
}

func (st JobEnvironment) MarshalJSON() ([]byte, error) {
	pb, err := JobEnvironmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobEnvironment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobEnvironmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobEnvironmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobEnvironmentToPb(st *JobEnvironment) (*jobspb.JobEnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobEnvironmentPb{}
	pb.EnvironmentKey = st.EnvironmentKey
	specPb, err := compute.EnvironmentToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}

	return pb, nil
}

func JobEnvironmentFromPb(pb *jobspb.JobEnvironmentPb) (*JobEnvironment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobEnvironment{}
	st.EnvironmentKey = pb.EnvironmentKey
	specField, err := compute.EnvironmentFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}

	return st, nil
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	// Wire name: 'no_alert_for_canceled_runs'
	NoAlertForCanceledRuns bool ``
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (st JobNotificationSettings) MarshalJSON() ([]byte, error) {
	pb, err := JobNotificationSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobNotificationSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobNotificationSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobNotificationSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobNotificationSettingsToPb(st *JobNotificationSettings) (*jobspb.JobNotificationSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobNotificationSettingsPb{}
	pb.NoAlertForCanceledRuns = st.NoAlertForCanceledRuns
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobNotificationSettingsFromPb(pb *jobspb.JobNotificationSettingsPb) (*JobNotificationSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobNotificationSettings{}
	st.NoAlertForCanceledRuns = pb.NoAlertForCanceledRuns
	st.NoAlertForSkippedRuns = pb.NoAlertForSkippedRuns

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobParameter struct {
	// The optional default value of the parameter
	// Wire name: 'default'
	Default string ``
	// The name of the parameter
	// Wire name: 'name'
	Name string ``
	// The value used in the run
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st JobParameter) MarshalJSON() ([]byte, error) {
	pb, err := JobParameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobParameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobParameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobParameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobParameterToPb(st *JobParameter) (*jobspb.JobParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobParameterPb{}
	pb.Default = st.Default
	pb.Name = st.Name
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobParameterFromPb(pb *jobspb.JobParameterPb) (*JobParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobParameter{}
	st.Default = pb.Default
	st.Name = pb.Name
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	// Wire name: 'default'
	Default string ``
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	// Wire name: 'name'
	Name string ``
}

func (st JobParameterDefinition) MarshalJSON() ([]byte, error) {
	pb, err := JobParameterDefinitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobParameterDefinition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobParameterDefinitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobParameterDefinitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobParameterDefinitionToPb(st *JobParameterDefinition) (*jobspb.JobParameterDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobParameterDefinitionPb{}
	pb.Default = st.Default
	pb.Name = st.Name

	return pb, nil
}

func JobParameterDefinitionFromPb(pb *jobspb.JobParameterDefinitionPb) (*JobParameterDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobParameterDefinition{}
	st.Default = pb.Default
	st.Name = pb.Name

	return st, nil
}

type JobPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel JobPermissionLevel ``
	ForceSendFields []string           `tf:"-"`
}

func (st JobPermission) MarshalJSON() ([]byte, error) {
	pb, err := JobPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobPermissionToPb(st *JobPermission) (*jobspb.JobPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := JobPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobPermissionFromPb(pb *jobspb.JobPermissionPb) (*JobPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := JobPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Permission level
type JobPermissionLevel string

const JobPermissionLevelCanManage JobPermissionLevel = `CAN_MANAGE`

const JobPermissionLevelCanManageRun JobPermissionLevel = `CAN_MANAGE_RUN`

const JobPermissionLevelCanView JobPermissionLevel = `CAN_VIEW`

const JobPermissionLevelIsOwner JobPermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *JobPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = JobPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MANAGE_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for JobPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobPermissionLevel) Values() []JobPermissionLevel {
	return []JobPermissionLevel{
		JobPermissionLevelCanManage,
		JobPermissionLevelCanManageRun,
		JobPermissionLevelCanView,
		JobPermissionLevelIsOwner,
	}
}

// Type always returns JobPermissionLevel to satisfy [pflag.Value] interface
func (f *JobPermissionLevel) Type() string {
	return "JobPermissionLevel"
}

func JobPermissionLevelToPb(st *JobPermissionLevel) (*jobspb.JobPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.JobPermissionLevelPb(*st)
	return &pb, nil
}

func JobPermissionLevelFromPb(pb *jobspb.JobPermissionLevelPb) (*JobPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobPermissionLevel(*pb)
	return &st, nil
}

type JobPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st JobPermissions) MarshalJSON() ([]byte, error) {
	pb, err := JobPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobPermissionsToPb(st *JobPermissions) (*jobspb.JobPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobPermissionsPb{}

	var accessControlListPb []jobspb.JobAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := JobAccessControlResponseToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobPermissionsFromPb(pb *jobspb.JobPermissionsPb) (*JobPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissions{}

	var accessControlListField []JobAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := JobAccessControlResponseFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel JobPermissionLevel ``
	ForceSendFields []string           `tf:"-"`
}

func (st JobPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := JobPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobPermissionsDescriptionToPb(st *JobPermissionsDescription) (*jobspb.JobPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := JobPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobPermissionsDescriptionFromPb(pb *jobspb.JobPermissionsDescriptionPb) (*JobPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := JobPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlRequest ``
	// The job for which to get or manage permissions.
	// Wire name: 'job_id'
	JobId string `tf:"-"`
}

func (st JobPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := JobPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobPermissionsRequestToPb(st *JobPermissionsRequest) (*jobspb.JobPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobPermissionsRequestPb{}

	var accessControlListPb []jobspb.JobAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := JobAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.JobId = st.JobId

	return pb, nil
}

func JobPermissionsRequestFromPb(pb *jobspb.JobPermissionsRequestPb) (*JobPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissionsRequest{}

	var accessControlListField []JobAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := JobAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.JobId = pb.JobId

	return st, nil
}

// Write-only setting. Specifies the user or service principal that the job runs
// as. If not specified, the job runs as the user who created the job.
//
// Either `user_name` or `service_principal_name` should be specified. If not,
// an error is thrown.
type JobRunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st JobRunAs) MarshalJSON() ([]byte, error) {
	pb, err := JobRunAsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobRunAs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobRunAsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobRunAsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobRunAsToPb(st *JobRunAs) (*jobspb.JobRunAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobRunAsPb{}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobRunAsFromPb(pb *jobspb.JobRunAsPb) (*JobRunAs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobRunAs{}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobSettings struct {
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	// Wire name: 'continuous'
	Continuous *Continuous ``
	// Deployment information for jobs managed by external sources.
	// Wire name: 'deployment'
	Deployment *JobDeployment ``
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	// Wire name: 'description'
	Description string ``
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	// Wire name: 'edit_mode'
	EditMode JobEditMode ``
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications ``
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	// Wire name: 'environments'
	Environments []JobEnvironment ``
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	// Wire name: 'format'
	Format Format ``
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource ``

	// Wire name: 'health'
	Health *JobsHealthRules ``
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster ``
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	// Wire name: 'max_concurrent_runs'
	MaxConcurrentRuns int ``
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'name'
	Name string ``
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	// Wire name: 'notification_settings'
	NotificationSettings *JobNotificationSettings ``
	// Job-level parameter definitions
	// Wire name: 'parameters'
	Parameters []JobParameterDefinition ``
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget ``
	// The queue settings of the job.
	// Wire name: 'queue'
	Queue *QueueSettings ``
	// The user or service principal that the job runs as, if specified in the
	// request. This field indicates the explicit configuration of `run_as` for
	// the job. To find the value in all cases, explicit or implicit, use
	// `run_as_user_name`.
	// Wire name: 'run_as'
	RunAs *JobRunAs ``
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	// Wire name: 'schedule'
	Schedule *CronSchedule ``
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	// Wire name: 'tags'
	Tags map[string]string ``
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	// Wire name: 'tasks'
	Tasks []Task ``
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int ``
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	// Wire name: 'trigger'
	Trigger *TriggerSettings ``
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications ``
	ForceSendFields      []string              `tf:"-"`
}

func (st JobSettings) MarshalJSON() ([]byte, error) {
	pb, err := JobSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobSettingsToPb(st *JobSettings) (*jobspb.JobSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobSettingsPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	continuousPb, err := ContinuousToPb(st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = continuousPb
	}
	deploymentPb, err := JobDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}
	pb.Description = st.Description
	editModePb, err := JobEditModeToPb(&st.EditMode)
	if err != nil {
		return nil, err
	}
	if editModePb != nil {
		pb.EditMode = *editModePb
	}
	emailNotificationsPb, err := JobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	var environmentsPb []jobspb.JobEnvironmentPb
	for _, item := range st.Environments {
		itemPb, err := JobEnvironmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentsPb = append(environmentsPb, *itemPb)
		}
	}
	pb.Environments = environmentsPb
	formatPb, err := FormatToPb(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}
	gitSourcePb, err := GitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}
	healthPb, err := JobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	var jobClustersPb []jobspb.JobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := JobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb
	pb.MaxConcurrentRuns = st.MaxConcurrentRuns
	pb.Name = st.Name
	notificationSettingsPb, err := JobNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	var parametersPb []jobspb.JobParameterDefinitionPb
	for _, item := range st.Parameters {
		itemPb, err := JobParameterDefinitionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb
	performanceTargetPb, err := PerformanceTargetToPb(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}
	queuePb, err := QueueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}
	runAsPb, err := JobRunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}
	schedulePb, err := CronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.Tags = st.Tags

	var tasksPb []jobspb.TaskPb
	for _, item := range st.Tasks {
		itemPb, err := TaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb
	pb.TimeoutSeconds = st.TimeoutSeconds
	triggerPb, err := TriggerSettingsToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}
	webhookNotificationsPb, err := WebhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobSettingsFromPb(pb *jobspb.JobSettingsPb) (*JobSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSettings{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	continuousField, err := ContinuousFromPb(pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = continuousField
	}
	deploymentField, err := JobDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	st.Description = pb.Description
	editModeField, err := JobEditModeFromPb(&pb.EditMode)
	if err != nil {
		return nil, err
	}
	if editModeField != nil {
		st.EditMode = *editModeField
	}
	emailNotificationsField, err := JobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}

	var environmentsField []JobEnvironment
	for _, itemPb := range pb.Environments {
		item, err := JobEnvironmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentsField = append(environmentsField, *item)
		}
	}
	st.Environments = environmentsField
	formatField, err := FormatFromPb(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	gitSourceField, err := GitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	healthField, err := JobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := JobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	st.MaxConcurrentRuns = pb.MaxConcurrentRuns
	st.Name = pb.Name
	notificationSettingsField, err := JobNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}

	var parametersField []JobParameterDefinition
	for _, itemPb := range pb.Parameters {
		item, err := JobParameterDefinitionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	performanceTargetField, err := PerformanceTargetFromPb(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	queueField, err := QueueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	runAsField, err := JobRunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	scheduleField, err := CronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.Tags = pb.Tags

	var tasksField []Task
	for _, itemPb := range pb.Tasks {
		item, err := TaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	st.TimeoutSeconds = pb.TimeoutSeconds
	triggerField, err := TriggerSettingsFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}
	webhookNotificationsField, err := WebhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The source of the job specification in the remote repository when the job is
// source controlled.
type JobSource struct {
	// Dirty state indicates the job is not fully synced with the job
	// specification in the remote repository.
	//
	// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
	// remote job specification. Import the remote job specification from UI to
	// make the job fully synced. * `DISCONNECTED`: The job is temporary
	// disconnected from the remote job specification and is allowed for live
	// edit. Import the remote job specification again from UI to make the job
	// fully synced.
	// Wire name: 'dirty_state'
	DirtyState JobSourceDirtyState ``
	// Name of the branch which the job is imported from.
	// Wire name: 'import_from_git_branch'
	ImportFromGitBranch string ``
	// Path of the job YAML file that contains the job specification.
	// Wire name: 'job_config_path'
	JobConfigPath string ``
}

func (st JobSource) MarshalJSON() ([]byte, error) {
	pb, err := JobSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobSourceToPb(st *JobSource) (*jobspb.JobSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobSourcePb{}
	dirtyStatePb, err := JobSourceDirtyStateToPb(&st.DirtyState)
	if err != nil {
		return nil, err
	}
	if dirtyStatePb != nil {
		pb.DirtyState = *dirtyStatePb
	}
	pb.ImportFromGitBranch = st.ImportFromGitBranch
	pb.JobConfigPath = st.JobConfigPath

	return pb, nil
}

func JobSourceFromPb(pb *jobspb.JobSourcePb) (*JobSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSource{}
	dirtyStateField, err := JobSourceDirtyStateFromPb(&pb.DirtyState)
	if err != nil {
		return nil, err
	}
	if dirtyStateField != nil {
		st.DirtyState = *dirtyStateField
	}
	st.ImportFromGitBranch = pb.ImportFromGitBranch
	st.JobConfigPath = pb.JobConfigPath

	return st, nil
}

// Dirty state indicates the job is not fully synced with the job specification
// in the remote repository.
//
// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
// remote job specification. Import the remote job specification from UI to make
// the job fully synced. * `DISCONNECTED`: The job is temporary disconnected
// from the remote job specification and is allowed for live edit. Import the
// remote job specification again from UI to make the job fully synced.
type JobSourceDirtyState string

// The job is temporary disconnected from the remote job specification and is
// allowed for live edit. Import the remote job specification again from UI to
// make the job fully synced.
const JobSourceDirtyStateDisconnected JobSourceDirtyState = `DISCONNECTED`

// The job is not yet synced with the remote job specification. Import the
// remote job specification from UI to make the job fully synced.
const JobSourceDirtyStateNotSynced JobSourceDirtyState = `NOT_SYNCED`

// String representation for [fmt.Print]
func (f *JobSourceDirtyState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobSourceDirtyState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `NOT_SYNCED`:
		*f = JobSourceDirtyState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "NOT_SYNCED"`, v)
	}
}

// Values returns all possible values for JobSourceDirtyState.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobSourceDirtyState) Values() []JobSourceDirtyState {
	return []JobSourceDirtyState{
		JobSourceDirtyStateDisconnected,
		JobSourceDirtyStateNotSynced,
	}
}

// Type always returns JobSourceDirtyState to satisfy [pflag.Value] interface
func (f *JobSourceDirtyState) Type() string {
	return "JobSourceDirtyState"
}

func JobSourceDirtyStateToPb(st *JobSourceDirtyState) (*jobspb.JobSourceDirtyStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.JobSourceDirtyStatePb(*st)
	return &pb, nil
}

func JobSourceDirtyStateFromPb(pb *jobspb.JobSourceDirtyStatePb) (*JobSourceDirtyState, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobSourceDirtyState(*pb)
	return &st, nil
}

// Specifies the health metric that is being evaluated for a particular health
// rule.
//
// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data waiting
// to be consumed across all streams. This metric is in Public Preview. *
// `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset lag across all
// streams. This metric is in Public Preview. * `STREAMING_BACKLOG_SECONDS`: An
// estimate of the maximum consumer delay across all streams. This metric is in
// Public Preview. * `STREAMING_BACKLOG_FILES`: An estimate of the maximum
// number of outstanding files across all streams. This metric is in Public
// Preview.
type JobsHealthMetric string

// Expected total time for a run in seconds.
const JobsHealthMetricRunDurationSeconds JobsHealthMetric = `RUN_DURATION_SECONDS`

// An estimate of the maximum bytes of data waiting to be consumed across all
// streams. This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogBytes JobsHealthMetric = `STREAMING_BACKLOG_BYTES`

// An estimate of the maximum number of outstanding files across all streams.
// This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogFiles JobsHealthMetric = `STREAMING_BACKLOG_FILES`

// An estimate of the maximum offset lag across all streams. This metric is in
// Public Preview.
const JobsHealthMetricStreamingBacklogRecords JobsHealthMetric = `STREAMING_BACKLOG_RECORDS`

// An estimate of the maximum consumer delay across all streams. This metric is
// in Public Preview.
const JobsHealthMetricStreamingBacklogSeconds JobsHealthMetric = `STREAMING_BACKLOG_SECONDS`

// String representation for [fmt.Print]
func (f *JobsHealthMetric) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthMetric) Set(v string) error {
	switch v {
	case `RUN_DURATION_SECONDS`, `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_FILES`, `STREAMING_BACKLOG_RECORDS`, `STREAMING_BACKLOG_SECONDS`:
		*f = JobsHealthMetric(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RUN_DURATION_SECONDS", "STREAMING_BACKLOG_BYTES", "STREAMING_BACKLOG_FILES", "STREAMING_BACKLOG_RECORDS", "STREAMING_BACKLOG_SECONDS"`, v)
	}
}

// Values returns all possible values for JobsHealthMetric.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobsHealthMetric) Values() []JobsHealthMetric {
	return []JobsHealthMetric{
		JobsHealthMetricRunDurationSeconds,
		JobsHealthMetricStreamingBacklogBytes,
		JobsHealthMetricStreamingBacklogFiles,
		JobsHealthMetricStreamingBacklogRecords,
		JobsHealthMetricStreamingBacklogSeconds,
	}
}

// Type always returns JobsHealthMetric to satisfy [pflag.Value] interface
func (f *JobsHealthMetric) Type() string {
	return "JobsHealthMetric"
}

func JobsHealthMetricToPb(st *JobsHealthMetric) (*jobspb.JobsHealthMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.JobsHealthMetricPb(*st)
	return &pb, nil
}

func JobsHealthMetricFromPb(pb *jobspb.JobsHealthMetricPb) (*JobsHealthMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobsHealthMetric(*pb)
	return &st, nil
}

// Specifies the operator used to compare the health metric value with the
// specified threshold.
type JobsHealthOperator string

const JobsHealthOperatorGreaterThan JobsHealthOperator = `GREATER_THAN`

// String representation for [fmt.Print]
func (f *JobsHealthOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthOperator) Set(v string) error {
	switch v {
	case `GREATER_THAN`:
		*f = JobsHealthOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GREATER_THAN"`, v)
	}
}

// Values returns all possible values for JobsHealthOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *JobsHealthOperator) Values() []JobsHealthOperator {
	return []JobsHealthOperator{
		JobsHealthOperatorGreaterThan,
	}
}

// Type always returns JobsHealthOperator to satisfy [pflag.Value] interface
func (f *JobsHealthOperator) Type() string {
	return "JobsHealthOperator"
}

func JobsHealthOperatorToPb(st *JobsHealthOperator) (*jobspb.JobsHealthOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.JobsHealthOperatorPb(*st)
	return &pb, nil
}

func JobsHealthOperatorFromPb(pb *jobspb.JobsHealthOperatorPb) (*JobsHealthOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobsHealthOperator(*pb)
	return &st, nil
}

type JobsHealthRule struct {

	// Wire name: 'metric'
	Metric JobsHealthMetric ``

	// Wire name: 'op'
	Op JobsHealthOperator ``
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	// Wire name: 'value'
	Value int64 ``
}

func (st JobsHealthRule) MarshalJSON() ([]byte, error) {
	pb, err := JobsHealthRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobsHealthRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobsHealthRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobsHealthRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobsHealthRuleToPb(st *JobsHealthRule) (*jobspb.JobsHealthRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobsHealthRulePb{}
	metricPb, err := JobsHealthMetricToPb(&st.Metric)
	if err != nil {
		return nil, err
	}
	if metricPb != nil {
		pb.Metric = *metricPb
	}
	opPb, err := JobsHealthOperatorToPb(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}
	pb.Value = st.Value

	return pb, nil
}

func JobsHealthRuleFromPb(pb *jobspb.JobsHealthRulePb) (*JobsHealthRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobsHealthRule{}
	metricField, err := JobsHealthMetricFromPb(&pb.Metric)
	if err != nil {
		return nil, err
	}
	if metricField != nil {
		st.Metric = *metricField
	}
	opField, err := JobsHealthOperatorFromPb(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	st.Value = pb.Value

	return st, nil
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {

	// Wire name: 'rules'
	Rules []JobsHealthRule ``
}

func (st JobsHealthRules) MarshalJSON() ([]byte, error) {
	pb, err := JobsHealthRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobsHealthRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.JobsHealthRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobsHealthRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobsHealthRulesToPb(st *JobsHealthRules) (*jobspb.JobsHealthRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.JobsHealthRulesPb{}

	var rulesPb []jobspb.JobsHealthRulePb
	for _, item := range st.Rules {
		itemPb, err := JobsHealthRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rulesPb = append(rulesPb, *itemPb)
		}
	}
	pb.Rules = rulesPb

	return pb, nil
}

func JobsHealthRulesFromPb(pb *jobspb.JobsHealthRulesPb) (*JobsHealthRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobsHealthRules{}

	var rulesField []JobsHealthRule
	for _, itemPb := range pb.Rules {
		item, err := JobsHealthRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rulesField = append(rulesField, *item)
		}
	}
	st.Rules = rulesField

	return st, nil
}

type ListJobComplianceForPolicyResponse struct {
	// A list of jobs and their policy compliance statuses.
	// Wire name: 'jobs'
	Jobs []JobCompliance ``
	// This field represents the pagination token to retrieve the next page of
	// results. If this field is not in the response, it means no further
	// results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// This field represents the pagination token to retrieve the previous page
	// of results. If this field is not in the response, it means no further
	// results for the request.
	// Wire name: 'prev_page_token'
	PrevPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListJobComplianceForPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListJobComplianceForPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListJobComplianceForPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ListJobComplianceForPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListJobComplianceForPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListJobComplianceForPolicyResponseToPb(st *ListJobComplianceForPolicyResponse) (*jobspb.ListJobComplianceForPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ListJobComplianceForPolicyResponsePb{}

	var jobsPb []jobspb.JobCompliancePb
	for _, item := range st.Jobs {
		itemPb, err := JobComplianceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobsPb = append(jobsPb, *itemPb)
		}
	}
	pb.Jobs = jobsPb
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListJobComplianceForPolicyResponseFromPb(pb *jobspb.ListJobComplianceForPolicyResponsePb) (*ListJobComplianceForPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobComplianceForPolicyResponse{}

	var jobsField []JobCompliance
	for _, itemPb := range pb.Jobs {
		item, err := JobComplianceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobsField = append(jobsField, *item)
		}
	}
	st.Jobs = jobsField
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListJobComplianceRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Canonical unique identifier for the cluster policy.
	// Wire name: 'policy_id'
	PolicyId        string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListJobComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListJobComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListJobComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ListJobComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListJobComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListJobComplianceRequestToPb(st *ListJobComplianceRequest) (*jobspb.ListJobComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ListJobComplianceRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListJobComplianceRequestFromPb(pb *jobspb.ListJobComplianceRequestPb) (*ListJobComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobComplianceRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListJobsRequest struct {
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/get to
	// paginate through all tasks and clusters.
	// Wire name: 'expand_tasks'
	ExpandTasks bool `tf:"-"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	// Wire name: 'limit'
	Limit int `tf:"-"`
	// A filter on the list based on the exact (case insensitive) job name.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The offset of the first job to return, relative to the most recently
	// created job. Deprecated since June 2023. Use `page_token` to iterate
	// through the pages instead.
	// Wire name: 'offset'
	Offset int `tf:"-"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListJobsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListJobsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListJobsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ListJobsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListJobsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListJobsRequestToPb(st *ListJobsRequest) (*jobspb.ListJobsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ListJobsRequestPb{}
	pb.ExpandTasks = st.ExpandTasks
	pb.Limit = st.Limit
	pb.Name = st.Name
	pb.Offset = st.Offset
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListJobsRequestFromPb(pb *jobspb.ListJobsRequestPb) (*ListJobsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobsRequest{}
	st.ExpandTasks = pb.ExpandTasks
	st.Limit = pb.Limit
	st.Name = pb.Name
	st.Offset = pb.Offset
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// List of jobs was retrieved successfully.
type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	// Wire name: 'has_more'
	HasMore bool ``
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	// Wire name: 'jobs'
	Jobs []BaseJob ``
	// A token that can be used to list the next page of jobs (if applicable).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	// Wire name: 'prev_page_token'
	PrevPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListJobsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListJobsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListJobsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ListJobsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListJobsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListJobsResponseToPb(st *ListJobsResponse) (*jobspb.ListJobsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ListJobsResponsePb{}
	pb.HasMore = st.HasMore

	var jobsPb []jobspb.BaseJobPb
	for _, item := range st.Jobs {
		itemPb, err := BaseJobToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobsPb = append(jobsPb, *itemPb)
		}
	}
	pb.Jobs = jobsPb
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListJobsResponseFromPb(pb *jobspb.ListJobsResponsePb) (*ListJobsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobsResponse{}
	st.HasMore = pb.HasMore

	var jobsField []BaseJob
	for _, itemPb := range pb.Jobs {
		item, err := BaseJobFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobsField = append(jobsField, *item)
		}
	}
	st.Jobs = jobsField
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	// Wire name: 'active_only'
	ActiveOnly bool `tf:"-"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	// Wire name: 'completed_only'
	CompletedOnly bool `tf:"-"`
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/getrun to
	// paginate through all tasks and clusters.
	// Wire name: 'expand_tasks'
	ExpandTasks bool `tf:"-"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	// Wire name: 'job_id'
	JobId int64 `tf:"-"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	// Wire name: 'limit'
	Limit int `tf:"-"`
	// The offset of the first run to return, relative to the most recent run.
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	// Wire name: 'offset'
	Offset int `tf:"-"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	// Wire name: 'run_type'
	RunType RunType `tf:"-"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	// Wire name: 'start_time_from'
	StartTimeFrom int64 `tf:"-"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	// Wire name: 'start_time_to'
	StartTimeTo     int64    `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListRunsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListRunsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListRunsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ListRunsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListRunsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListRunsRequestToPb(st *ListRunsRequest) (*jobspb.ListRunsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ListRunsRequestPb{}
	pb.ActiveOnly = st.ActiveOnly
	pb.CompletedOnly = st.CompletedOnly
	pb.ExpandTasks = st.ExpandTasks
	pb.JobId = st.JobId
	pb.Limit = st.Limit
	pb.Offset = st.Offset
	pb.PageToken = st.PageToken
	runTypePb, err := RunTypeToPb(&st.RunType)
	if err != nil {
		return nil, err
	}
	if runTypePb != nil {
		pb.RunType = *runTypePb
	}
	pb.StartTimeFrom = st.StartTimeFrom
	pb.StartTimeTo = st.StartTimeTo

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListRunsRequestFromPb(pb *jobspb.ListRunsRequestPb) (*ListRunsRequest, error) {
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
	runTypeField, err := RunTypeFromPb(&pb.RunType)
	if err != nil {
		return nil, err
	}
	if runTypeField != nil {
		st.RunType = *runTypeField
	}
	st.StartTimeFrom = pb.StartTimeFrom
	st.StartTimeTo = pb.StartTimeTo

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// List of runs was retrieved successfully.
type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	// Wire name: 'has_more'
	HasMore bool ``
	// A token that can be used to list the next page of runs (if applicable).
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// A token that can be used to list the previous page of runs (if
	// applicable).
	// Wire name: 'prev_page_token'
	PrevPageToken string ``
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	// Wire name: 'runs'
	Runs            []BaseRun ``
	ForceSendFields []string  `tf:"-"`
}

func (st ListRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ListRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListRunsResponseToPb(st *ListRunsResponse) (*jobspb.ListRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ListRunsResponsePb{}
	pb.HasMore = st.HasMore
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	var runsPb []jobspb.BaseRunPb
	for _, item := range st.Runs {
		itemPb, err := BaseRunToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			runsPb = append(runsPb, *itemPb)
		}
	}
	pb.Runs = runsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListRunsResponseFromPb(pb *jobspb.ListRunsResponsePb) (*ListRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRunsResponse{}
	st.HasMore = pb.HasMore
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	var runsField []BaseRun
	for _, itemPb := range pb.Runs {
		item, err := BaseRunFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			runsField = append(runsField, *item)
		}
	}
	st.Runs = runsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	// Wire name: 'result'
	Result string ``
	// Whether or not the result was truncated.
	// Wire name: 'truncated'
	Truncated       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st NotebookOutput) MarshalJSON() ([]byte, error) {
	pb, err := NotebookOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NotebookOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.NotebookOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NotebookOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NotebookOutputToPb(st *NotebookOutput) (*jobspb.NotebookOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.NotebookOutputPb{}
	pb.Result = st.Result
	pb.Truncated = st.Truncated

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NotebookOutputFromPb(pb *jobspb.NotebookOutputPb) (*NotebookOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookOutput{}
	st.Result = pb.Result
	st.Truncated = pb.Truncated

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type NotebookTask struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to :method:jobs/run Now with parameters specified,
	// the two parameters maps are merged. If the same key is specified in
	// `base_parameters` and in `run-now`, the value from `run-now` is used. Use
	// [Task parameter variables] to set parameters containing information about
	// job runs.
	//
	// If the notebook takes a parameter that is not specified in the job’s
	// `base_parameters` or the `run-now` override parameters, the default value
	// from the notebook is used.
	//
	// Retrieve these parameters in a notebook using [dbutils.widgets.get].
	//
	// The JSON representation of this field cannot exceed 1MB.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-widgets
	// Wire name: 'base_parameters'
	BaseParameters map[string]string ``
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	// Wire name: 'notebook_path'
	NotebookPath string ``
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	// Wire name: 'source'
	Source Source ``
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st NotebookTask) MarshalJSON() ([]byte, error) {
	pb, err := NotebookTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NotebookTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.NotebookTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NotebookTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NotebookTaskToPb(st *NotebookTask) (*jobspb.NotebookTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.NotebookTaskPb{}
	pb.BaseParameters = st.BaseParameters
	pb.NotebookPath = st.NotebookPath
	sourcePb, err := SourceToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}
	pb.WarehouseId = st.WarehouseId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NotebookTaskFromPb(pb *jobspb.NotebookTaskPb) (*NotebookTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookTask{}
	st.BaseParameters = pb.BaseParameters
	st.NotebookPath = pb.NotebookPath
	sourceField, err := SourceFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	st.WarehouseId = pb.WarehouseId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Stores the catalog name, schema name, and the output schema expiration time
// for the clean room run.
type OutputSchemaInfo struct {

	// Wire name: 'catalog_name'
	CatalogName string ``
	// The expiration time for the output schema as a Unix timestamp in
	// milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64 ``

	// Wire name: 'schema_name'
	SchemaName      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st OutputSchemaInfo) MarshalJSON() ([]byte, error) {
	pb, err := OutputSchemaInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *OutputSchemaInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.OutputSchemaInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := OutputSchemaInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func OutputSchemaInfoToPb(st *OutputSchemaInfo) (*jobspb.OutputSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.OutputSchemaInfoPb{}
	pb.CatalogName = st.CatalogName
	pb.ExpirationTime = st.ExpirationTime
	pb.SchemaName = st.SchemaName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func OutputSchemaInfoFromPb(pb *jobspb.OutputSchemaInfoPb) (*OutputSchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OutputSchemaInfo{}
	st.CatalogName = pb.CatalogName
	st.ExpirationTime = pb.ExpirationTime
	st.SchemaName = pb.SchemaName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PauseStatus string

const PauseStatusPaused PauseStatus = `PAUSED`

const PauseStatusUnpaused PauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *PauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = PauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Values returns all possible values for PauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *PauseStatus) Values() []PauseStatus {
	return []PauseStatus{
		PauseStatusPaused,
		PauseStatusUnpaused,
	}
}

// Type always returns PauseStatus to satisfy [pflag.Value] interface
func (f *PauseStatus) Type() string {
	return "PauseStatus"
}

func PauseStatusToPb(st *PauseStatus) (*jobspb.PauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.PauseStatusPb(*st)
	return &pb, nil
}

func PauseStatusFromPb(pb *jobspb.PauseStatusPb) (*PauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := PauseStatus(*pb)
	return &st, nil
}

// PerformanceTarget defines how performant (lower latency) or cost efficient
// the execution of run on serverless compute should be. The performance mode on
// the job or pipeline should map to a performance setting that is passed to
// Cluster Manager (see cluster-common PerformanceTarget).
type PerformanceTarget string

const PerformanceTargetPerformanceOptimized PerformanceTarget = `PERFORMANCE_OPTIMIZED`

const PerformanceTargetStandard PerformanceTarget = `STANDARD`

// String representation for [fmt.Print]
func (f *PerformanceTarget) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PerformanceTarget) Set(v string) error {
	switch v {
	case `PERFORMANCE_OPTIMIZED`, `STANDARD`:
		*f = PerformanceTarget(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PERFORMANCE_OPTIMIZED", "STANDARD"`, v)
	}
}

// Values returns all possible values for PerformanceTarget.
//
// There is no guarantee on the order of the values in the slice.
func (f *PerformanceTarget) Values() []PerformanceTarget {
	return []PerformanceTarget{
		PerformanceTargetPerformanceOptimized,
		PerformanceTargetStandard,
	}
}

// Type always returns PerformanceTarget to satisfy [pflag.Value] interface
func (f *PerformanceTarget) Type() string {
	return "PerformanceTarget"
}

func PerformanceTargetToPb(st *PerformanceTarget) (*jobspb.PerformanceTargetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.PerformanceTargetPb(*st)
	return &pb, nil
}

func PerformanceTargetFromPb(pb *jobspb.PerformanceTargetPb) (*PerformanceTarget, error) {
	if pb == nil {
		return nil, nil
	}
	st := PerformanceTarget(*pb)
	return &st, nil
}

type PeriodicTriggerConfiguration struct {
	// The interval at which the trigger should run.
	// Wire name: 'interval'
	Interval int ``
	// The unit of time for the interval.
	// Wire name: 'unit'
	Unit PeriodicTriggerConfigurationTimeUnit ``
}

func (st PeriodicTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := PeriodicTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PeriodicTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PeriodicTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PeriodicTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PeriodicTriggerConfigurationToPb(st *PeriodicTriggerConfiguration) (*jobspb.PeriodicTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PeriodicTriggerConfigurationPb{}
	pb.Interval = st.Interval
	unitPb, err := PeriodicTriggerConfigurationTimeUnitToPb(&st.Unit)
	if err != nil {
		return nil, err
	}
	if unitPb != nil {
		pb.Unit = *unitPb
	}

	return pb, nil
}

func PeriodicTriggerConfigurationFromPb(pb *jobspb.PeriodicTriggerConfigurationPb) (*PeriodicTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PeriodicTriggerConfiguration{}
	st.Interval = pb.Interval
	unitField, err := PeriodicTriggerConfigurationTimeUnitFromPb(&pb.Unit)
	if err != nil {
		return nil, err
	}
	if unitField != nil {
		st.Unit = *unitField
	}

	return st, nil
}

type PeriodicTriggerConfigurationTimeUnit string

const PeriodicTriggerConfigurationTimeUnitDays PeriodicTriggerConfigurationTimeUnit = `DAYS`

const PeriodicTriggerConfigurationTimeUnitHours PeriodicTriggerConfigurationTimeUnit = `HOURS`

const PeriodicTriggerConfigurationTimeUnitWeeks PeriodicTriggerConfigurationTimeUnit = `WEEKS`

// String representation for [fmt.Print]
func (f *PeriodicTriggerConfigurationTimeUnit) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PeriodicTriggerConfigurationTimeUnit) Set(v string) error {
	switch v {
	case `DAYS`, `HOURS`, `WEEKS`:
		*f = PeriodicTriggerConfigurationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAYS", "HOURS", "WEEKS"`, v)
	}
}

// Values returns all possible values for PeriodicTriggerConfigurationTimeUnit.
//
// There is no guarantee on the order of the values in the slice.
func (f *PeriodicTriggerConfigurationTimeUnit) Values() []PeriodicTriggerConfigurationTimeUnit {
	return []PeriodicTriggerConfigurationTimeUnit{
		PeriodicTriggerConfigurationTimeUnitDays,
		PeriodicTriggerConfigurationTimeUnitHours,
		PeriodicTriggerConfigurationTimeUnitWeeks,
	}
}

// Type always returns PeriodicTriggerConfigurationTimeUnit to satisfy [pflag.Value] interface
func (f *PeriodicTriggerConfigurationTimeUnit) Type() string {
	return "PeriodicTriggerConfigurationTimeUnit"
}

func PeriodicTriggerConfigurationTimeUnitToPb(st *PeriodicTriggerConfigurationTimeUnit) (*jobspb.PeriodicTriggerConfigurationTimeUnitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.PeriodicTriggerConfigurationTimeUnitPb(*st)
	return &pb, nil
}

func PeriodicTriggerConfigurationTimeUnitFromPb(pb *jobspb.PeriodicTriggerConfigurationTimeUnitPb) (*PeriodicTriggerConfigurationTimeUnit, error) {
	if pb == nil {
		return nil, nil
	}
	st := PeriodicTriggerConfigurationTimeUnit(*pb)
	return &st, nil
}

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	// Wire name: 'full_refresh'
	FullRefresh     bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st PipelineParams) MarshalJSON() ([]byte, error) {
	pb, err := PipelineParamsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PipelineParams) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PipelineParamsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PipelineParamsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PipelineParamsToPb(st *PipelineParams) (*jobspb.PipelineParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PipelineParamsPb{}
	pb.FullRefresh = st.FullRefresh

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PipelineParamsFromPb(pb *jobspb.PipelineParamsPb) (*PipelineParams, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineParams{}
	st.FullRefresh = pb.FullRefresh

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PipelineTask struct {
	// If true, triggers a full refresh on the delta live table.
	// Wire name: 'full_refresh'
	FullRefresh bool ``
	// The full name of the pipeline task to execute.
	// Wire name: 'pipeline_id'
	PipelineId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PipelineTask) MarshalJSON() ([]byte, error) {
	pb, err := PipelineTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PipelineTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PipelineTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PipelineTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PipelineTaskToPb(st *PipelineTask) (*jobspb.PipelineTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PipelineTaskPb{}
	pb.FullRefresh = st.FullRefresh
	pb.PipelineId = st.PipelineId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PipelineTaskFromPb(pb *jobspb.PipelineTaskPb) (*PipelineTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineTask{}
	st.FullRefresh = pb.FullRefresh
	st.PipelineId = pb.PipelineId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PowerBiModel struct {
	// How the published Power BI model authenticates to Databricks
	// Wire name: 'authentication_method'
	AuthenticationMethod AuthenticationMethod ``
	// The name of the Power BI model
	// Wire name: 'model_name'
	ModelName string ``
	// Whether to overwrite existing Power BI models
	// Wire name: 'overwrite_existing'
	OverwriteExisting bool ``
	// The default storage mode of the Power BI model
	// Wire name: 'storage_mode'
	StorageMode StorageMode ``
	// The name of the Power BI workspace of the model
	// Wire name: 'workspace_name'
	WorkspaceName   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PowerBiModel) MarshalJSON() ([]byte, error) {
	pb, err := PowerBiModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PowerBiModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PowerBiModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PowerBiModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PowerBiModelToPb(st *PowerBiModel) (*jobspb.PowerBiModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PowerBiModelPb{}
	authenticationMethodPb, err := AuthenticationMethodToPb(&st.AuthenticationMethod)
	if err != nil {
		return nil, err
	}
	if authenticationMethodPb != nil {
		pb.AuthenticationMethod = *authenticationMethodPb
	}
	pb.ModelName = st.ModelName
	pb.OverwriteExisting = st.OverwriteExisting
	storageModePb, err := StorageModeToPb(&st.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModePb != nil {
		pb.StorageMode = *storageModePb
	}
	pb.WorkspaceName = st.WorkspaceName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PowerBiModelFromPb(pb *jobspb.PowerBiModelPb) (*PowerBiModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiModel{}
	authenticationMethodField, err := AuthenticationMethodFromPb(&pb.AuthenticationMethod)
	if err != nil {
		return nil, err
	}
	if authenticationMethodField != nil {
		st.AuthenticationMethod = *authenticationMethodField
	}
	st.ModelName = pb.ModelName
	st.OverwriteExisting = pb.OverwriteExisting
	storageModeField, err := StorageModeFromPb(&pb.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModeField != nil {
		st.StorageMode = *storageModeField
	}
	st.WorkspaceName = pb.WorkspaceName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PowerBiTable struct {
	// The catalog name in Databricks
	// Wire name: 'catalog'
	Catalog string ``
	// The table name in Databricks
	// Wire name: 'name'
	Name string ``
	// The schema name in Databricks
	// Wire name: 'schema'
	Schema string ``
	// The Power BI storage mode of the table
	// Wire name: 'storage_mode'
	StorageMode     StorageMode ``
	ForceSendFields []string    `tf:"-"`
}

func (st PowerBiTable) MarshalJSON() ([]byte, error) {
	pb, err := PowerBiTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PowerBiTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PowerBiTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PowerBiTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PowerBiTableToPb(st *PowerBiTable) (*jobspb.PowerBiTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PowerBiTablePb{}
	pb.Catalog = st.Catalog
	pb.Name = st.Name
	pb.Schema = st.Schema
	storageModePb, err := StorageModeToPb(&st.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModePb != nil {
		pb.StorageMode = *storageModePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PowerBiTableFromPb(pb *jobspb.PowerBiTablePb) (*PowerBiTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiTable{}
	st.Catalog = pb.Catalog
	st.Name = pb.Name
	st.Schema = pb.Schema
	storageModeField, err := StorageModeFromPb(&pb.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModeField != nil {
		st.StorageMode = *storageModeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PowerBiTask struct {
	// The resource name of the UC connection to authenticate from Databricks to
	// Power BI
	// Wire name: 'connection_resource_name'
	ConnectionResourceName string ``
	// The semantic model to update
	// Wire name: 'power_bi_model'
	PowerBiModel *PowerBiModel ``
	// Whether the model should be refreshed after the update
	// Wire name: 'refresh_after_update'
	RefreshAfterUpdate bool ``
	// The tables to be exported to Power BI
	// Wire name: 'tables'
	Tables []PowerBiTable ``
	// The SQL warehouse ID to use as the Power BI data source
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PowerBiTask) MarshalJSON() ([]byte, error) {
	pb, err := PowerBiTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PowerBiTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PowerBiTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PowerBiTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PowerBiTaskToPb(st *PowerBiTask) (*jobspb.PowerBiTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PowerBiTaskPb{}
	pb.ConnectionResourceName = st.ConnectionResourceName
	powerBiModelPb, err := PowerBiModelToPb(st.PowerBiModel)
	if err != nil {
		return nil, err
	}
	if powerBiModelPb != nil {
		pb.PowerBiModel = powerBiModelPb
	}
	pb.RefreshAfterUpdate = st.RefreshAfterUpdate

	var tablesPb []jobspb.PowerBiTablePb
	for _, item := range st.Tables {
		itemPb, err := PowerBiTableToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tablesPb = append(tablesPb, *itemPb)
		}
	}
	pb.Tables = tablesPb
	pb.WarehouseId = st.WarehouseId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PowerBiTaskFromPb(pb *jobspb.PowerBiTaskPb) (*PowerBiTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiTask{}
	st.ConnectionResourceName = pb.ConnectionResourceName
	powerBiModelField, err := PowerBiModelFromPb(pb.PowerBiModel)
	if err != nil {
		return nil, err
	}
	if powerBiModelField != nil {
		st.PowerBiModel = powerBiModelField
	}
	st.RefreshAfterUpdate = pb.RefreshAfterUpdate

	var tablesField []PowerBiTable
	for _, itemPb := range pb.Tables {
		item, err := PowerBiTableFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tablesField = append(tablesField, *item)
		}
	}
	st.Tables = tablesField
	st.WarehouseId = pb.WarehouseId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	// Wire name: 'entry_point'
	EntryPoint string ``
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	// Wire name: 'named_parameters'
	NamedParameters map[string]string ``
	// Name of the package to execute
	// Wire name: 'package_name'
	PackageName string ``
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	// Wire name: 'parameters'
	Parameters []string ``
}

func (st PythonWheelTask) MarshalJSON() ([]byte, error) {
	pb, err := PythonWheelTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PythonWheelTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.PythonWheelTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PythonWheelTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PythonWheelTaskToPb(st *PythonWheelTask) (*jobspb.PythonWheelTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.PythonWheelTaskPb{}
	pb.EntryPoint = st.EntryPoint
	pb.NamedParameters = st.NamedParameters
	pb.PackageName = st.PackageName
	pb.Parameters = st.Parameters

	return pb, nil
}

func PythonWheelTaskFromPb(pb *jobspb.PythonWheelTaskPb) (*PythonWheelTask, error) {
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

type QueueDetails struct {

	// Wire name: 'code'
	Code QueueDetailsCodeCode ``
	// A descriptive message with the queuing details. This field is
	// unstructured, and its exact format is subject to change.
	// Wire name: 'message'
	Message         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st QueueDetails) MarshalJSON() ([]byte, error) {
	pb, err := QueueDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *QueueDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.QueueDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := QueueDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func QueueDetailsToPb(st *QueueDetails) (*jobspb.QueueDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.QueueDetailsPb{}
	codePb, err := QueueDetailsCodeCodeToPb(&st.Code)
	if err != nil {
		return nil, err
	}
	if codePb != nil {
		pb.Code = *codePb
	}
	pb.Message = st.Message

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func QueueDetailsFromPb(pb *jobspb.QueueDetailsPb) (*QueueDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueueDetails{}
	codeField, err := QueueDetailsCodeCodeFromPb(&pb.Code)
	if err != nil {
		return nil, err
	}
	if codeField != nil {
		st.Code = *codeField
	}
	st.Message = pb.Message

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run was
// queued due to reaching the workspace limit of active task runs. *
// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the per-job
// limit of concurrent job runs. * `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run
// was queued due to reaching the workspace limit of active run job tasks.
type QueueDetailsCodeCode string

// The run was queued due to reaching the workspace limit of active task runs.
const QueueDetailsCodeCodeActiveRunsLimitReached QueueDetailsCodeCode = `ACTIVE_RUNS_LIMIT_REACHED`

// The run was queued due to reaching the workspace limit of active run job
// tasks.
const QueueDetailsCodeCodeActiveRunJobTasksLimitReached QueueDetailsCodeCode = `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`

// The run was queued due to reaching the per-job limit of concurrent job runs.
const QueueDetailsCodeCodeMaxConcurrentRunsReached QueueDetailsCodeCode = `MAX_CONCURRENT_RUNS_REACHED`

// String representation for [fmt.Print]
func (f *QueueDetailsCodeCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueueDetailsCodeCode) Set(v string) error {
	switch v {
	case `ACTIVE_RUNS_LIMIT_REACHED`, `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`, `MAX_CONCURRENT_RUNS_REACHED`:
		*f = QueueDetailsCodeCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_RUNS_LIMIT_REACHED", "ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED", "MAX_CONCURRENT_RUNS_REACHED"`, v)
	}
}

// Values returns all possible values for QueueDetailsCodeCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueueDetailsCodeCode) Values() []QueueDetailsCodeCode {
	return []QueueDetailsCodeCode{
		QueueDetailsCodeCodeActiveRunsLimitReached,
		QueueDetailsCodeCodeActiveRunJobTasksLimitReached,
		QueueDetailsCodeCodeMaxConcurrentRunsReached,
	}
}

// Type always returns QueueDetailsCodeCode to satisfy [pflag.Value] interface
func (f *QueueDetailsCodeCode) Type() string {
	return "QueueDetailsCodeCode"
}

func QueueDetailsCodeCodeToPb(st *QueueDetailsCodeCode) (*jobspb.QueueDetailsCodeCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.QueueDetailsCodeCodePb(*st)
	return &pb, nil
}

func QueueDetailsCodeCodeFromPb(pb *jobspb.QueueDetailsCodeCodePb) (*QueueDetailsCodeCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueueDetailsCodeCode(*pb)
	return &st, nil
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	// Wire name: 'enabled'
	Enabled bool ``
}

func (st QueueSettings) MarshalJSON() ([]byte, error) {
	pb, err := QueueSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *QueueSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.QueueSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := QueueSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func QueueSettingsToPb(st *QueueSettings) (*jobspb.QueueSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.QueueSettingsPb{}
	pb.Enabled = st.Enabled

	return pb, nil
}

func QueueSettingsFromPb(pb *jobspb.QueueSettingsPb) (*QueueSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueueSettings{}
	st.Enabled = pb.Enabled

	return st, nil
}

type RepairHistoryItem struct {
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget ``
	// The end time of the (repaired) run.
	// Wire name: 'end_time'
	EndTime int64 ``
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	// Wire name: 'id'
	Id int64 ``
	// The start time of the (repaired) run.
	// Wire name: 'start_time'
	StartTime int64 ``
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState ``

	// Wire name: 'status'
	Status *RunStatus ``
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	// Wire name: 'task_run_ids'
	TaskRunIds []int64 ``
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	// Wire name: 'type'
	Type            RepairHistoryItemType ``
	ForceSendFields []string              `tf:"-"`
}

func (st RepairHistoryItem) MarshalJSON() ([]byte, error) {
	pb, err := RepairHistoryItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepairHistoryItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RepairHistoryItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepairHistoryItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RepairHistoryItemToPb(st *RepairHistoryItem) (*jobspb.RepairHistoryItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RepairHistoryItemPb{}
	effectivePerformanceTargetPb, err := PerformanceTargetToPb(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}
	pb.EndTime = st.EndTime
	pb.Id = st.Id
	pb.StartTime = st.StartTime
	statePb, err := RunStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	statusPb, err := RunStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}
	pb.TaskRunIds = st.TaskRunIds
	typePb, err := RepairHistoryItemTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RepairHistoryItemFromPb(pb *jobspb.RepairHistoryItemPb) (*RepairHistoryItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairHistoryItem{}
	effectivePerformanceTargetField, err := PerformanceTargetFromPb(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	st.EndTime = pb.EndTime
	st.Id = pb.Id
	st.StartTime = pb.StartTime
	stateField, err := RunStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := RunStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.TaskRunIds = pb.TaskRunIds
	typeField, err := RepairHistoryItemTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The repair history item type. Indicates whether a run is the original run or
// a repair run.
type RepairHistoryItemType string

const RepairHistoryItemTypeOriginal RepairHistoryItemType = `ORIGINAL`

const RepairHistoryItemTypeRepair RepairHistoryItemType = `REPAIR`

// String representation for [fmt.Print]
func (f *RepairHistoryItemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RepairHistoryItemType) Set(v string) error {
	switch v {
	case `ORIGINAL`, `REPAIR`:
		*f = RepairHistoryItemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ORIGINAL", "REPAIR"`, v)
	}
}

// Values returns all possible values for RepairHistoryItemType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RepairHistoryItemType) Values() []RepairHistoryItemType {
	return []RepairHistoryItemType{
		RepairHistoryItemTypeOriginal,
		RepairHistoryItemTypeRepair,
	}
}

// Type always returns RepairHistoryItemType to satisfy [pflag.Value] interface
func (f *RepairHistoryItemType) Type() string {
	return "RepairHistoryItemType"
}

func RepairHistoryItemTypeToPb(st *RepairHistoryItemType) (*jobspb.RepairHistoryItemTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.RepairHistoryItemTypePb(*st)
	return &pb, nil
}

func RepairHistoryItemTypeFromPb(pb *jobspb.RepairHistoryItemTypePb) (*RepairHistoryItemType, error) {
	if pb == nil {
		return nil, nil
	}
	st := RepairHistoryItemType(*pb)
	return &st, nil
}

type RepairRun struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string ``
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'jar_params'
	JarParams []string ``
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	// Wire name: 'job_parameters'
	JobParameters map[string]string ``
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	// Wire name: 'latest_repair_id'
	LatestRepairId int64 ``
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// Wire name: 'notebook_params'
	NotebookParams map[string]string ``
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget ``
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams ``

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string ``
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'python_params'
	PythonParams []string ``
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	// Wire name: 'rerun_all_failed_tasks'
	RerunAllFailedTasks bool ``
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	// Wire name: 'rerun_dependent_tasks'
	RerunDependentTasks bool ``
	// The task keys of the task runs to repair.
	// Wire name: 'rerun_tasks'
	RerunTasks []string ``
	// The job run ID of the run to repair. The run must not be in progress.
	// Wire name: 'run_id'
	RunId int64 ``
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'spark_submit_params'
	SparkSubmitParams []string ``
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams       map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (st RepairRun) MarshalJSON() ([]byte, error) {
	pb, err := RepairRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepairRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RepairRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepairRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RepairRunToPb(st *RepairRun) (*jobspb.RepairRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RepairRunPb{}
	pb.DbtCommands = st.DbtCommands
	pb.JarParams = st.JarParams
	pb.JobParameters = st.JobParameters
	pb.LatestRepairId = st.LatestRepairId
	pb.NotebookParams = st.NotebookParams
	performanceTargetPb, err := PerformanceTargetToPb(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}
	pipelineParamsPb, err := PipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.RerunAllFailedTasks = st.RerunAllFailedTasks
	pb.RerunDependentTasks = st.RerunDependentTasks
	pb.RerunTasks = st.RerunTasks
	pb.RunId = st.RunId
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RepairRunFromPb(pb *jobspb.RepairRunPb) (*RepairRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairRun{}
	st.DbtCommands = pb.DbtCommands
	st.JarParams = pb.JarParams
	st.JobParameters = pb.JobParameters
	st.LatestRepairId = pb.LatestRepairId
	st.NotebookParams = pb.NotebookParams
	performanceTargetField, err := PerformanceTargetFromPb(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	pipelineParamsField, err := PipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.RerunAllFailedTasks = pb.RerunAllFailedTasks
	st.RerunDependentTasks = pb.RerunDependentTasks
	st.RerunTasks = pb.RerunTasks
	st.RunId = pb.RunId
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Run repair was initiated.
type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	// Wire name: 'repair_id'
	RepairId        int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st RepairRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := RepairRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RepairRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RepairRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RepairRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RepairRunResponseToPb(st *RepairRunResponse) (*jobspb.RepairRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RepairRunResponsePb{}
	pb.RepairId = st.RepairId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RepairRunResponseFromPb(pb *jobspb.RepairRunResponsePb) (*RepairRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairRunResponse{}
	st.RepairId = pb.RepairId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	// Wire name: 'job_id'
	JobId int64 ``
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	// Wire name: 'new_settings'
	NewSettings JobSettings ``
}

func (st ResetJob) MarshalJSON() ([]byte, error) {
	pb, err := ResetJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResetJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResetJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResetJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResetJobToPb(st *ResetJob) (*jobspb.ResetJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResetJobPb{}
	pb.JobId = st.JobId
	newSettingsPb, err := JobSettingsToPb(&st.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsPb != nil {
		pb.NewSettings = *newSettingsPb
	}

	return pb, nil
}

func ResetJobFromPb(pb *jobspb.ResetJobPb) (*ResetJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResetJob{}
	st.JobId = pb.JobId
	newSettingsField, err := JobSettingsFromPb(&pb.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsField != nil {
		st.NewSettings = *newSettingsField
	}

	return st, nil
}

type ResolvedConditionTaskValues struct {

	// Wire name: 'left'
	Left string ``

	// Wire name: 'right'
	Right           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ResolvedConditionTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedConditionTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedConditionTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedConditionTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedConditionTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedConditionTaskValuesToPb(st *ResolvedConditionTaskValues) (*jobspb.ResolvedConditionTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedConditionTaskValuesPb{}
	pb.Left = st.Left
	pb.Right = st.Right

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ResolvedConditionTaskValuesFromPb(pb *jobspb.ResolvedConditionTaskValuesPb) (*ResolvedConditionTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedConditionTaskValues{}
	st.Left = pb.Left
	st.Right = pb.Right

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ResolvedDbtTaskValues struct {

	// Wire name: 'commands'
	Commands []string ``
}

func (st ResolvedDbtTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedDbtTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedDbtTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedDbtTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedDbtTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedDbtTaskValuesToPb(st *ResolvedDbtTaskValues) (*jobspb.ResolvedDbtTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedDbtTaskValuesPb{}
	pb.Commands = st.Commands

	return pb, nil
}

func ResolvedDbtTaskValuesFromPb(pb *jobspb.ResolvedDbtTaskValuesPb) (*ResolvedDbtTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedDbtTaskValues{}
	st.Commands = pb.Commands

	return st, nil
}

type ResolvedNotebookTaskValues struct {

	// Wire name: 'base_parameters'
	BaseParameters map[string]string ``
}

func (st ResolvedNotebookTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedNotebookTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedNotebookTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedNotebookTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedNotebookTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedNotebookTaskValuesToPb(st *ResolvedNotebookTaskValues) (*jobspb.ResolvedNotebookTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedNotebookTaskValuesPb{}
	pb.BaseParameters = st.BaseParameters

	return pb, nil
}

func ResolvedNotebookTaskValuesFromPb(pb *jobspb.ResolvedNotebookTaskValuesPb) (*ResolvedNotebookTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedNotebookTaskValues{}
	st.BaseParameters = pb.BaseParameters

	return st, nil
}

type ResolvedParamPairValues struct {

	// Wire name: 'parameters'
	Parameters map[string]string ``
}

func (st ResolvedParamPairValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedParamPairValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedParamPairValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedParamPairValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedParamPairValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedParamPairValuesToPb(st *ResolvedParamPairValues) (*jobspb.ResolvedParamPairValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedParamPairValuesPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

func ResolvedParamPairValuesFromPb(pb *jobspb.ResolvedParamPairValuesPb) (*ResolvedParamPairValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedParamPairValues{}
	st.Parameters = pb.Parameters

	return st, nil
}

type ResolvedPythonWheelTaskValues struct {

	// Wire name: 'named_parameters'
	NamedParameters map[string]string ``

	// Wire name: 'parameters'
	Parameters []string ``
}

func (st ResolvedPythonWheelTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedPythonWheelTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedPythonWheelTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedPythonWheelTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedPythonWheelTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedPythonWheelTaskValuesToPb(st *ResolvedPythonWheelTaskValues) (*jobspb.ResolvedPythonWheelTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedPythonWheelTaskValuesPb{}
	pb.NamedParameters = st.NamedParameters
	pb.Parameters = st.Parameters

	return pb, nil
}

func ResolvedPythonWheelTaskValuesFromPb(pb *jobspb.ResolvedPythonWheelTaskValuesPb) (*ResolvedPythonWheelTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedPythonWheelTaskValues{}
	st.NamedParameters = pb.NamedParameters
	st.Parameters = pb.Parameters

	return st, nil
}

type ResolvedRunJobTaskValues struct {

	// Wire name: 'job_parameters'
	JobParameters map[string]string ``

	// Wire name: 'parameters'
	Parameters map[string]string ``
}

func (st ResolvedRunJobTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedRunJobTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedRunJobTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedRunJobTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedRunJobTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedRunJobTaskValuesToPb(st *ResolvedRunJobTaskValues) (*jobspb.ResolvedRunJobTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedRunJobTaskValuesPb{}
	pb.JobParameters = st.JobParameters
	pb.Parameters = st.Parameters

	return pb, nil
}

func ResolvedRunJobTaskValuesFromPb(pb *jobspb.ResolvedRunJobTaskValuesPb) (*ResolvedRunJobTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedRunJobTaskValues{}
	st.JobParameters = pb.JobParameters
	st.Parameters = pb.Parameters

	return st, nil
}

type ResolvedStringParamsValues struct {

	// Wire name: 'parameters'
	Parameters []string ``
}

func (st ResolvedStringParamsValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedStringParamsValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedStringParamsValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedStringParamsValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedStringParamsValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedStringParamsValuesToPb(st *ResolvedStringParamsValues) (*jobspb.ResolvedStringParamsValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedStringParamsValuesPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

func ResolvedStringParamsValuesFromPb(pb *jobspb.ResolvedStringParamsValuesPb) (*ResolvedStringParamsValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedStringParamsValues{}
	st.Parameters = pb.Parameters

	return st, nil
}

type ResolvedValues struct {

	// Wire name: 'condition_task'
	ConditionTask *ResolvedConditionTaskValues ``

	// Wire name: 'dbt_task'
	DbtTask *ResolvedDbtTaskValues ``

	// Wire name: 'notebook_task'
	NotebookTask *ResolvedNotebookTaskValues ``

	// Wire name: 'python_wheel_task'
	PythonWheelTask *ResolvedPythonWheelTaskValues ``

	// Wire name: 'run_job_task'
	RunJobTask *ResolvedRunJobTaskValues ``

	// Wire name: 'simulation_task'
	SimulationTask *ResolvedParamPairValues ``

	// Wire name: 'spark_jar_task'
	SparkJarTask *ResolvedStringParamsValues ``

	// Wire name: 'spark_python_task'
	SparkPythonTask *ResolvedStringParamsValues ``

	// Wire name: 'spark_submit_task'
	SparkSubmitTask *ResolvedStringParamsValues ``

	// Wire name: 'sql_task'
	SqlTask *ResolvedParamPairValues ``
}

func (st ResolvedValues) MarshalJSON() ([]byte, error) {
	pb, err := ResolvedValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResolvedValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ResolvedValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResolvedValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResolvedValuesToPb(st *ResolvedValues) (*jobspb.ResolvedValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ResolvedValuesPb{}
	conditionTaskPb, err := ResolvedConditionTaskValuesToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}
	dbtTaskPb, err := ResolvedDbtTaskValuesToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}
	notebookTaskPb, err := ResolvedNotebookTaskValuesToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}
	pythonWheelTaskPb, err := ResolvedPythonWheelTaskValuesToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}
	runJobTaskPb, err := ResolvedRunJobTaskValuesToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}
	simulationTaskPb, err := ResolvedParamPairValuesToPb(st.SimulationTask)
	if err != nil {
		return nil, err
	}
	if simulationTaskPb != nil {
		pb.SimulationTask = simulationTaskPb
	}
	sparkJarTaskPb, err := ResolvedStringParamsValuesToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}
	sparkPythonTaskPb, err := ResolvedStringParamsValuesToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}
	sparkSubmitTaskPb, err := ResolvedStringParamsValuesToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}
	sqlTaskPb, err := ResolvedParamPairValuesToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}

	return pb, nil
}

func ResolvedValuesFromPb(pb *jobspb.ResolvedValuesPb) (*ResolvedValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedValues{}
	conditionTaskField, err := ResolvedConditionTaskValuesFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dbtTaskField, err := ResolvedDbtTaskValuesFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}
	notebookTaskField, err := ResolvedNotebookTaskValuesFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	pythonWheelTaskField, err := ResolvedPythonWheelTaskValuesFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	runJobTaskField, err := ResolvedRunJobTaskValuesFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	simulationTaskField, err := ResolvedParamPairValuesFromPb(pb.SimulationTask)
	if err != nil {
		return nil, err
	}
	if simulationTaskField != nil {
		st.SimulationTask = simulationTaskField
	}
	sparkJarTaskField, err := ResolvedStringParamsValuesFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := ResolvedStringParamsValuesFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := ResolvedStringParamsValuesFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := ResolvedParamPairValuesFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}

	return st, nil
}

// Run was retrieved successfully
type Run struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	// Wire name: 'attempt_number'
	AttemptNumber int ``
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	// Wire name: 'cleanup_duration'
	CleanupDuration int64 ``
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	// Wire name: 'cluster_instance'
	ClusterInstance *ClusterInstance ``
	// A snapshot of the job’s cluster specification when this run was
	// created.
	// Wire name: 'cluster_spec'
	ClusterSpec *ClusterSpec ``
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// Description of the run
	// Wire name: 'description'
	Description string ``
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget ``
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	// Wire name: 'end_time'
	EndTime int64 ``
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	// Wire name: 'execution_duration'
	ExecutionDuration int64 ``
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource ``
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool ``
	// Only populated by for-each iterations. The parent for-each task is
	// located in tasks array.
	// Wire name: 'iterations'
	Iterations []RunTask ``
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster ``
	// The canonical identifier of the job that contains this run.
	// Wire name: 'job_id'
	JobId int64 ``
	// Job-level parameters used in the run
	// Wire name: 'job_parameters'
	JobParameters []JobParameter ``
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	// Wire name: 'job_run_id'
	JobRunId int64 ``
	// A token that can be used to list the next page of array properties.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	// Wire name: 'number_in_job'
	NumberInJob int64 ``
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	// Wire name: 'original_attempt_run_id'
	OriginalAttemptRunId int64 ``
	// The parameters used for this run.
	// Wire name: 'overriding_parameters'
	OverridingParameters *RunParameters ``
	// The time in milliseconds that the run has spent in the queue.
	// Wire name: 'queue_duration'
	QueueDuration int64 ``
	// The repair history of the run.
	// Wire name: 'repair_history'
	RepairHistory []RepairHistoryItem ``
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	// Wire name: 'run_duration'
	RunDuration int64 ``
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'run_id'
	RunId int64 ``
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'run_name'
	RunName string ``
	// The URL to the detail page of the run.
	// Wire name: 'run_page_url'
	RunPageUrl string ``

	// Wire name: 'run_type'
	RunType RunType ``
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	// Wire name: 'schedule'
	Schedule *CronSchedule ``
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	// Wire name: 'setup_duration'
	SetupDuration int64 ``
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	// Wire name: 'start_time'
	StartTime int64 ``
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState ``

	// Wire name: 'status'
	Status *RunStatus ``
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	// Wire name: 'tasks'
	Tasks []RunTask ``

	// Wire name: 'trigger'
	Trigger TriggerType ``

	// Wire name: 'trigger_info'
	TriggerInfo     *TriggerInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (st Run) MarshalJSON() ([]byte, error) {
	pb, err := RunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Run) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunToPb(st *Run) (*jobspb.RunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunPb{}
	pb.AttemptNumber = st.AttemptNumber
	pb.CleanupDuration = st.CleanupDuration
	clusterInstancePb, err := ClusterInstanceToPb(st.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstancePb != nil {
		pb.ClusterInstance = clusterInstancePb
	}
	clusterSpecPb, err := ClusterSpecToPb(st.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecPb != nil {
		pb.ClusterSpec = clusterSpecPb
	}
	pb.CreatorUserName = st.CreatorUserName
	pb.Description = st.Description
	effectivePerformanceTargetPb, err := PerformanceTargetToPb(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}
	pb.EndTime = st.EndTime
	pb.ExecutionDuration = st.ExecutionDuration
	gitSourcePb, err := GitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}
	pb.HasMore = st.HasMore

	var iterationsPb []jobspb.RunTaskPb
	for _, item := range st.Iterations {
		itemPb, err := RunTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			iterationsPb = append(iterationsPb, *itemPb)
		}
	}
	pb.Iterations = iterationsPb

	var jobClustersPb []jobspb.JobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := JobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb
	pb.JobId = st.JobId

	var jobParametersPb []jobspb.JobParameterPb
	for _, item := range st.JobParameters {
		itemPb, err := JobParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb = append(jobParametersPb, *itemPb)
		}
	}
	pb.JobParameters = jobParametersPb
	pb.JobRunId = st.JobRunId
	pb.NextPageToken = st.NextPageToken
	pb.NumberInJob = st.NumberInJob
	pb.OriginalAttemptRunId = st.OriginalAttemptRunId
	overridingParametersPb, err := RunParametersToPb(st.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersPb != nil {
		pb.OverridingParameters = overridingParametersPb
	}
	pb.QueueDuration = st.QueueDuration

	var repairHistoryPb []jobspb.RepairHistoryItemPb
	for _, item := range st.RepairHistory {
		itemPb, err := RepairHistoryItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repairHistoryPb = append(repairHistoryPb, *itemPb)
		}
	}
	pb.RepairHistory = repairHistoryPb
	pb.RunDuration = st.RunDuration
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunPageUrl = st.RunPageUrl
	runTypePb, err := RunTypeToPb(&st.RunType)
	if err != nil {
		return nil, err
	}
	if runTypePb != nil {
		pb.RunType = *runTypePb
	}
	schedulePb, err := CronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	pb.SetupDuration = st.SetupDuration
	pb.StartTime = st.StartTime
	statePb, err := RunStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	statusPb, err := RunStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	var tasksPb []jobspb.RunTaskPb
	for _, item := range st.Tasks {
		itemPb, err := RunTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb
	triggerPb, err := TriggerTypeToPb(&st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = *triggerPb
	}
	triggerInfoPb, err := TriggerInfoToPb(st.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoPb != nil {
		pb.TriggerInfo = triggerInfoPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunFromPb(pb *jobspb.RunPb) (*Run, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Run{}
	st.AttemptNumber = pb.AttemptNumber
	st.CleanupDuration = pb.CleanupDuration
	clusterInstanceField, err := ClusterInstanceFromPb(pb.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstanceField != nil {
		st.ClusterInstance = clusterInstanceField
	}
	clusterSpecField, err := ClusterSpecFromPb(pb.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecField != nil {
		st.ClusterSpec = clusterSpecField
	}
	st.CreatorUserName = pb.CreatorUserName
	st.Description = pb.Description
	effectivePerformanceTargetField, err := PerformanceTargetFromPb(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	st.EndTime = pb.EndTime
	st.ExecutionDuration = pb.ExecutionDuration
	gitSourceField, err := GitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	st.HasMore = pb.HasMore

	var iterationsField []RunTask
	for _, itemPb := range pb.Iterations {
		item, err := RunTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			iterationsField = append(iterationsField, *item)
		}
	}
	st.Iterations = iterationsField

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := JobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	st.JobId = pb.JobId

	var jobParametersField []JobParameter
	for _, itemPb := range pb.JobParameters {
		item, err := JobParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField = append(jobParametersField, *item)
		}
	}
	st.JobParameters = jobParametersField
	st.JobRunId = pb.JobRunId
	st.NextPageToken = pb.NextPageToken
	st.NumberInJob = pb.NumberInJob
	st.OriginalAttemptRunId = pb.OriginalAttemptRunId
	overridingParametersField, err := RunParametersFromPb(pb.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersField != nil {
		st.OverridingParameters = overridingParametersField
	}
	st.QueueDuration = pb.QueueDuration

	var repairHistoryField []RepairHistoryItem
	for _, itemPb := range pb.RepairHistory {
		item, err := RepairHistoryItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repairHistoryField = append(repairHistoryField, *item)
		}
	}
	st.RepairHistory = repairHistoryField
	st.RunDuration = pb.RunDuration
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunPageUrl = pb.RunPageUrl
	runTypeField, err := RunTypeFromPb(&pb.RunType)
	if err != nil {
		return nil, err
	}
	if runTypeField != nil {
		st.RunType = *runTypeField
	}
	scheduleField, err := CronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.SetupDuration = pb.SetupDuration
	st.StartTime = pb.StartTime
	stateField, err := RunStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := RunStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	var tasksField []RunTask
	for _, itemPb := range pb.Tasks {
		item, err := RunTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	triggerField, err := TriggerTypeFromPb(&pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = *triggerField
	}
	triggerInfoField, err := TriggerInfoFromPb(pb.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoField != nil {
		st.TriggerInfo = triggerInfoField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RunConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	// Wire name: 'left'
	Left string ``
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	// Wire name: 'op'
	Op ConditionTaskOp ``
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	// Wire name: 'outcome'
	Outcome string ``
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	// Wire name: 'right'
	Right           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RunConditionTask) MarshalJSON() ([]byte, error) {
	pb, err := RunConditionTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunConditionTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunConditionTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunConditionTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunConditionTaskToPb(st *RunConditionTask) (*jobspb.RunConditionTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunConditionTaskPb{}
	pb.Left = st.Left
	opPb, err := ConditionTaskOpToPb(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}
	pb.Outcome = st.Outcome
	pb.Right = st.Right

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunConditionTaskFromPb(pb *jobspb.RunConditionTaskPb) (*RunConditionTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunConditionTask{}
	st.Left = pb.Left
	opField, err := ConditionTaskOpFromPb(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	st.Outcome = pb.Outcome
	st.Right = pb.Right

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RunForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	// Wire name: 'concurrency'
	Concurrency int ``
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	// Wire name: 'inputs'
	Inputs string ``
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of an For each task
	// Wire name: 'stats'
	Stats *ForEachStats ``
	// Configuration for the task that will be run for each element in the array
	// Wire name: 'task'
	Task            Task     ``
	ForceSendFields []string `tf:"-"`
}

func (st RunForEachTask) MarshalJSON() ([]byte, error) {
	pb, err := RunForEachTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunForEachTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunForEachTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunForEachTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunForEachTaskToPb(st *RunForEachTask) (*jobspb.RunForEachTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunForEachTaskPb{}
	pb.Concurrency = st.Concurrency
	pb.Inputs = st.Inputs
	statsPb, err := ForEachStatsToPb(st.Stats)
	if err != nil {
		return nil, err
	}
	if statsPb != nil {
		pb.Stats = statsPb
	}
	taskPb, err := TaskToPb(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunForEachTaskFromPb(pb *jobspb.RunForEachTaskPb) (*RunForEachTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunForEachTask{}
	st.Concurrency = pb.Concurrency
	st.Inputs = pb.Inputs
	statsField, err := ForEachStatsFromPb(pb.Stats)
	if err != nil {
		return nil, err
	}
	if statsField != nil {
		st.Stats = statsField
	}
	taskField, err := TaskFromPb(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// An optional value indicating the condition that determines whether the task
// should be run once its dependencies have been completed. When omitted,
// defaults to `ALL_SUCCESS`.
//
// Possible values are: * `ALL_SUCCESS`: All dependencies have executed and
// succeeded * `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
// `NONE_FAILED`: None of the dependencies have failed and at least one was
// executed * `ALL_DONE`: All dependencies have been completed *
// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
// dependencies have failed
type RunIf string

// All dependencies have been completed
const RunIfAllDone RunIf = `ALL_DONE`

// ALl dependencies have failed
const RunIfAllFailed RunIf = `ALL_FAILED`

// All dependencies have executed and succeeded
const RunIfAllSuccess RunIf = `ALL_SUCCESS`

// At least one dependency failed
const RunIfAtLeastOneFailed RunIf = `AT_LEAST_ONE_FAILED`

// At least one dependency has succeeded
const RunIfAtLeastOneSuccess RunIf = `AT_LEAST_ONE_SUCCESS`

// None of the dependencies have failed and at least one was executed
const RunIfNoneFailed RunIf = `NONE_FAILED`

// String representation for [fmt.Print]
func (f *RunIf) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunIf) Set(v string) error {
	switch v {
	case `ALL_DONE`, `ALL_FAILED`, `ALL_SUCCESS`, `AT_LEAST_ONE_FAILED`, `AT_LEAST_ONE_SUCCESS`, `NONE_FAILED`:
		*f = RunIf(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_DONE", "ALL_FAILED", "ALL_SUCCESS", "AT_LEAST_ONE_FAILED", "AT_LEAST_ONE_SUCCESS", "NONE_FAILED"`, v)
	}
}

// Values returns all possible values for RunIf.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunIf) Values() []RunIf {
	return []RunIf{
		RunIfAllDone,
		RunIfAllFailed,
		RunIfAllSuccess,
		RunIfAtLeastOneFailed,
		RunIfAtLeastOneSuccess,
		RunIfNoneFailed,
	}
}

// Type always returns RunIf to satisfy [pflag.Value] interface
func (f *RunIf) Type() string {
	return "RunIf"
}

func RunIfToPb(st *RunIf) (*jobspb.RunIfPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.RunIfPb(*st)
	return &pb, nil
}

func RunIfFromPb(pb *jobspb.RunIfPb) (*RunIf, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunIf(*pb)
	return &st, nil
}

type RunJobOutput struct {
	// The run id of the triggered job run
	// Wire name: 'run_id'
	RunId           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st RunJobOutput) MarshalJSON() ([]byte, error) {
	pb, err := RunJobOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunJobOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunJobOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunJobOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunJobOutputToPb(st *RunJobOutput) (*jobspb.RunJobOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunJobOutputPb{}
	pb.RunId = st.RunId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunJobOutputFromPb(pb *jobspb.RunJobOutputPb) (*RunJobOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunJobOutput{}
	st.RunId = pb.RunId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RunJobTask struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string ``
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'jar_params'
	JarParams []string ``
	// ID of the job to trigger.
	// Wire name: 'job_id'
	JobId int64 ``
	// Job-level parameters used to trigger the job.
	// Wire name: 'job_parameters'
	JobParameters map[string]string ``
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// Wire name: 'notebook_params'
	NotebookParams map[string]string ``
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams ``

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string ``
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'python_params'
	PythonParams []string ``
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'spark_submit_params'
	SparkSubmitParams []string ``
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams map[string]string ``
}

func (st RunJobTask) MarshalJSON() ([]byte, error) {
	pb, err := RunJobTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunJobTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunJobTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunJobTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunJobTaskToPb(st *RunJobTask) (*jobspb.RunJobTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunJobTaskPb{}
	pb.DbtCommands = st.DbtCommands
	pb.JarParams = st.JarParams
	pb.JobId = st.JobId
	pb.JobParameters = st.JobParameters
	pb.NotebookParams = st.NotebookParams
	pipelineParamsPb, err := PipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	return pb, nil
}

func RunJobTaskFromPb(pb *jobspb.RunJobTaskPb) (*RunJobTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunJobTask{}
	st.DbtCommands = pb.DbtCommands
	st.JarParams = pb.JarParams
	st.JobId = pb.JobId
	st.JobParameters = pb.JobParameters
	st.NotebookParams = pb.NotebookParams
	pipelineParamsField, err := PipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	return st, nil
}

// A value indicating the run's lifecycle state. The possible values are: *
// `QUEUED`: The run is queued. * `PENDING`: The run is waiting to be executed
// while the cluster and execution context are being prepared. * `RUNNING`: The
// task of this run is being executed. * `TERMINATING`: The task of this run has
// completed, and the cluster and execution context are being cleaned up. *
// `TERMINATED`: The task of this run has completed, and the cluster and
// execution context have been cleaned up. This state is terminal. * `SKIPPED`:
// This run was aborted because a previous run of the same job was already
// active. This state is terminal. * `INTERNAL_ERROR`: An exceptional state that
// indicates a failure in the Jobs service, such as network failure over a long
// period. If a run on a new cluster ends in the `INTERNAL_ERROR` state, the
// Jobs service terminates the cluster as soon as possible. This state is
// terminal. * `BLOCKED`: The run is blocked on an upstream dependency. *
// `WAITING_FOR_RETRY`: The run is waiting for a retry.
type RunLifeCycleState string

// The run is blocked on an upstream dependency.
const RunLifeCycleStateBlocked RunLifeCycleState = `BLOCKED`

// An exceptional state that indicates a failure in the Jobs service, such as
// network failure over a long period. If a run on a new cluster ends in the
// `INTERNAL_ERROR` state, the Jobs service terminates the cluster as soon as
// possible. This state is terminal.
const RunLifeCycleStateInternalError RunLifeCycleState = `INTERNAL_ERROR`

// The run is waiting to be executed while the cluster and execution context are
// being prepared.
const RunLifeCycleStatePending RunLifeCycleState = `PENDING`

// The run is queued.
const RunLifeCycleStateQueued RunLifeCycleState = `QUEUED`

// The task of this run is being executed.
const RunLifeCycleStateRunning RunLifeCycleState = `RUNNING`

// This run was aborted because a previous run of the same job was already
// active. This state is terminal.
const RunLifeCycleStateSkipped RunLifeCycleState = `SKIPPED`

// The task of this run has completed, and the cluster and execution context
// have been cleaned up. This state is terminal.
const RunLifeCycleStateTerminated RunLifeCycleState = `TERMINATED`

// The task of this run has completed, and the cluster and execution context are
// being cleaned up.
const RunLifeCycleStateTerminating RunLifeCycleState = `TERMINATING`

// The run is waiting for a retry.
const RunLifeCycleStateWaitingForRetry RunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *RunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = RunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Values returns all possible values for RunLifeCycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunLifeCycleState) Values() []RunLifeCycleState {
	return []RunLifeCycleState{
		RunLifeCycleStateBlocked,
		RunLifeCycleStateInternalError,
		RunLifeCycleStatePending,
		RunLifeCycleStateQueued,
		RunLifeCycleStateRunning,
		RunLifeCycleStateSkipped,
		RunLifeCycleStateTerminated,
		RunLifeCycleStateTerminating,
		RunLifeCycleStateWaitingForRetry,
	}
}

// Type always returns RunLifeCycleState to satisfy [pflag.Value] interface
func (f *RunLifeCycleState) Type() string {
	return "RunLifeCycleState"
}

func RunLifeCycleStateToPb(st *RunLifeCycleState) (*jobspb.RunLifeCycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.RunLifeCycleStatePb(*st)
	return &pb, nil
}

func RunLifeCycleStateFromPb(pb *jobspb.RunLifeCycleStatePb) (*RunLifeCycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunLifeCycleState(*pb)
	return &st, nil
}

// The current state of the run.
type RunLifecycleStateV2State string

const RunLifecycleStateV2StateBlocked RunLifecycleStateV2State = `BLOCKED`

const RunLifecycleStateV2StatePending RunLifecycleStateV2State = `PENDING`

const RunLifecycleStateV2StateQueued RunLifecycleStateV2State = `QUEUED`

const RunLifecycleStateV2StateRunning RunLifecycleStateV2State = `RUNNING`

const RunLifecycleStateV2StateTerminated RunLifecycleStateV2State = `TERMINATED`

const RunLifecycleStateV2StateTerminating RunLifecycleStateV2State = `TERMINATING`

const RunLifecycleStateV2StateWaiting RunLifecycleStateV2State = `WAITING`

// String representation for [fmt.Print]
func (f *RunLifecycleStateV2State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunLifecycleStateV2State) Set(v string) error {
	switch v {
	case `BLOCKED`, `PENDING`, `QUEUED`, `RUNNING`, `TERMINATED`, `TERMINATING`, `WAITING`:
		*f = RunLifecycleStateV2State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "PENDING", "QUEUED", "RUNNING", "TERMINATED", "TERMINATING", "WAITING"`, v)
	}
}

// Values returns all possible values for RunLifecycleStateV2State.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunLifecycleStateV2State) Values() []RunLifecycleStateV2State {
	return []RunLifecycleStateV2State{
		RunLifecycleStateV2StateBlocked,
		RunLifecycleStateV2StatePending,
		RunLifecycleStateV2StateQueued,
		RunLifecycleStateV2StateRunning,
		RunLifecycleStateV2StateTerminated,
		RunLifecycleStateV2StateTerminating,
		RunLifecycleStateV2StateWaiting,
	}
}

// Type always returns RunLifecycleStateV2State to satisfy [pflag.Value] interface
func (f *RunLifecycleStateV2State) Type() string {
	return "RunLifecycleStateV2State"
}

func RunLifecycleStateV2StateToPb(st *RunLifecycleStateV2State) (*jobspb.RunLifecycleStateV2StatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.RunLifecycleStateV2StatePb(*st)
	return &pb, nil
}

func RunLifecycleStateV2StateFromPb(pb *jobspb.RunLifecycleStateV2StatePb) (*RunLifecycleStateV2State, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunLifecycleStateV2State(*pb)
	return &st, nil
}

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string ``
	// An optional token to guarantee the idempotency of job run requests. If a
	// run with the provided token already exists, the request does not create a
	// new run but returns the ID of the existing run instead. If a run with the
	// provided token is deleted, an error is returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	// Wire name: 'idempotency_token'
	IdempotencyToken string ``
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'jar_params'
	JarParams []string ``
	// The ID of the job to be executed
	// Wire name: 'job_id'
	JobId int64 ``
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	// Wire name: 'job_parameters'
	JobParameters map[string]string ``
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// Wire name: 'notebook_params'
	NotebookParams map[string]string ``
	// A list of task keys to run inside of the job. If this field is not
	// provided, all tasks in the job will be run.
	// Wire name: 'only'
	Only []string ``
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget ``
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams ``

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string ``
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'python_params'
	PythonParams []string ``
	// The queue settings of the run.
	// Wire name: 'queue'
	Queue *QueueSettings ``
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'spark_submit_params'
	SparkSubmitParams []string ``
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams       map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (st RunNow) MarshalJSON() ([]byte, error) {
	pb, err := RunNowToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunNow) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunNowPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunNowFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunNowToPb(st *RunNow) (*jobspb.RunNowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunNowPb{}
	pb.DbtCommands = st.DbtCommands
	pb.IdempotencyToken = st.IdempotencyToken
	pb.JarParams = st.JarParams
	pb.JobId = st.JobId
	pb.JobParameters = st.JobParameters
	pb.NotebookParams = st.NotebookParams
	pb.Only = st.Only
	performanceTargetPb, err := PerformanceTargetToPb(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}
	pipelineParamsPb, err := PipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	queuePb, err := QueueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunNowFromPb(pb *jobspb.RunNowPb) (*RunNow, error) {
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
	performanceTargetField, err := PerformanceTargetFromPb(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	pipelineParamsField, err := PipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	queueField, err := QueueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Run was started successfully.
type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	// Wire name: 'number_in_job'
	NumberInJob int64 ``
	// The globally unique ID of the newly triggered run.
	// Wire name: 'run_id'
	RunId           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st RunNowResponse) MarshalJSON() ([]byte, error) {
	pb, err := RunNowResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunNowResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunNowResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunNowResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunNowResponseToPb(st *RunNowResponse) (*jobspb.RunNowResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunNowResponsePb{}
	pb.NumberInJob = st.NumberInJob
	pb.RunId = st.RunId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunNowResponseFromPb(pb *jobspb.RunNowResponsePb) (*RunNowResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunNowResponse{}
	st.NumberInJob = pb.NumberInJob
	st.RunId = pb.RunId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Run output was retrieved successfully.
type RunOutput struct {
	// The output of a clean rooms notebook task, if available
	// Wire name: 'clean_rooms_notebook_output'
	CleanRoomsNotebookOutput *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput ``
	// The output of a dashboard task, if available
	// Wire name: 'dashboard_output'
	DashboardOutput *DashboardTaskOutput ``
	// Deprecated in favor of the new dbt_platform_output
	// Wire name: 'dbt_cloud_output'
	DbtCloudOutput *DbtCloudTaskOutput ``
	// The output of a dbt task, if available.
	// Wire name: 'dbt_output'
	DbtOutput *DbtOutput ``

	// Wire name: 'dbt_platform_output'
	DbtPlatformOutput *DbtPlatformTaskOutput ``
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	// Wire name: 'error'
	Error string ``
	// If there was an error executing the run, this field contains any
	// available stack traces.
	// Wire name: 'error_trace'
	ErrorTrace string ``

	// Wire name: 'info'
	Info string ``
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	// Wire name: 'logs'
	Logs string ``
	// Whether the logs are truncated.
	// Wire name: 'logs_truncated'
	LogsTruncated bool ``
	// All details of the run except for its output.
	// Wire name: 'metadata'
	Metadata *Run ``
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	// Wire name: 'notebook_output'
	NotebookOutput *NotebookOutput ``
	// The output of a run job task, if available
	// Wire name: 'run_job_output'
	RunJobOutput *RunJobOutput ``
	// The output of a SQL task, if available.
	// Wire name: 'sql_output'
	SqlOutput       *SqlOutput ``
	ForceSendFields []string   `tf:"-"`
}

func (st RunOutput) MarshalJSON() ([]byte, error) {
	pb, err := RunOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunOutputToPb(st *RunOutput) (*jobspb.RunOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunOutputPb{}
	cleanRoomsNotebookOutputPb, err := CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(st.CleanRoomsNotebookOutput)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookOutputPb != nil {
		pb.CleanRoomsNotebookOutput = cleanRoomsNotebookOutputPb
	}
	dashboardOutputPb, err := DashboardTaskOutputToPb(st.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputPb != nil {
		pb.DashboardOutput = dashboardOutputPb
	}
	dbtCloudOutputPb, err := DbtCloudTaskOutputToPb(st.DbtCloudOutput)
	if err != nil {
		return nil, err
	}
	if dbtCloudOutputPb != nil {
		pb.DbtCloudOutput = dbtCloudOutputPb
	}
	dbtOutputPb, err := DbtOutputToPb(st.DbtOutput)
	if err != nil {
		return nil, err
	}
	if dbtOutputPb != nil {
		pb.DbtOutput = dbtOutputPb
	}
	dbtPlatformOutputPb, err := DbtPlatformTaskOutputToPb(st.DbtPlatformOutput)
	if err != nil {
		return nil, err
	}
	if dbtPlatformOutputPb != nil {
		pb.DbtPlatformOutput = dbtPlatformOutputPb
	}
	pb.Error = st.Error
	pb.ErrorTrace = st.ErrorTrace
	pb.Info = st.Info
	pb.Logs = st.Logs
	pb.LogsTruncated = st.LogsTruncated
	metadataPb, err := RunToPb(st.Metadata)
	if err != nil {
		return nil, err
	}
	if metadataPb != nil {
		pb.Metadata = metadataPb
	}
	notebookOutputPb, err := NotebookOutputToPb(st.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputPb != nil {
		pb.NotebookOutput = notebookOutputPb
	}
	runJobOutputPb, err := RunJobOutputToPb(st.RunJobOutput)
	if err != nil {
		return nil, err
	}
	if runJobOutputPb != nil {
		pb.RunJobOutput = runJobOutputPb
	}
	sqlOutputPb, err := SqlOutputToPb(st.SqlOutput)
	if err != nil {
		return nil, err
	}
	if sqlOutputPb != nil {
		pb.SqlOutput = sqlOutputPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunOutputFromPb(pb *jobspb.RunOutputPb) (*RunOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunOutput{}
	cleanRoomsNotebookOutputField, err := CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb.CleanRoomsNotebookOutput)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookOutputField != nil {
		st.CleanRoomsNotebookOutput = cleanRoomsNotebookOutputField
	}
	dashboardOutputField, err := DashboardTaskOutputFromPb(pb.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputField != nil {
		st.DashboardOutput = dashboardOutputField
	}
	dbtCloudOutputField, err := DbtCloudTaskOutputFromPb(pb.DbtCloudOutput)
	if err != nil {
		return nil, err
	}
	if dbtCloudOutputField != nil {
		st.DbtCloudOutput = dbtCloudOutputField
	}
	dbtOutputField, err := DbtOutputFromPb(pb.DbtOutput)
	if err != nil {
		return nil, err
	}
	if dbtOutputField != nil {
		st.DbtOutput = dbtOutputField
	}
	dbtPlatformOutputField, err := DbtPlatformTaskOutputFromPb(pb.DbtPlatformOutput)
	if err != nil {
		return nil, err
	}
	if dbtPlatformOutputField != nil {
		st.DbtPlatformOutput = dbtPlatformOutputField
	}
	st.Error = pb.Error
	st.ErrorTrace = pb.ErrorTrace
	st.Info = pb.Info
	st.Logs = pb.Logs
	st.LogsTruncated = pb.LogsTruncated
	metadataField, err := RunFromPb(pb.Metadata)
	if err != nil {
		return nil, err
	}
	if metadataField != nil {
		st.Metadata = metadataField
	}
	notebookOutputField, err := NotebookOutputFromPb(pb.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputField != nil {
		st.NotebookOutput = notebookOutputField
	}
	runJobOutputField, err := RunJobOutputFromPb(pb.RunJobOutput)
	if err != nil {
		return nil, err
	}
	if runJobOutputField != nil {
		st.RunJobOutput = runJobOutputField
	}
	sqlOutputField, err := SqlOutputFromPb(pb.SqlOutput)
	if err != nil {
		return nil, err
	}
	if sqlOutputField != nil {
		st.SqlOutput = sqlOutputField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string ``
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'jar_params'
	JarParams []string ``
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// Wire name: 'notebook_params'
	NotebookParams map[string]string ``
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams ``

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string ``
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'python_params'
	PythonParams []string ``
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'spark_submit_params'
	SparkSubmitParams []string ``
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams map[string]string ``
}

func (st RunParameters) MarshalJSON() ([]byte, error) {
	pb, err := RunParametersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunParameters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunParametersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunParametersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunParametersToPb(st *RunParameters) (*jobspb.RunParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunParametersPb{}
	pb.DbtCommands = st.DbtCommands
	pb.JarParams = st.JarParams
	pb.NotebookParams = st.NotebookParams
	pipelineParamsPb, err := PipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}
	pb.PythonNamedParams = st.PythonNamedParams
	pb.PythonParams = st.PythonParams
	pb.SparkSubmitParams = st.SparkSubmitParams
	pb.SqlParams = st.SqlParams

	return pb, nil
}

func RunParametersFromPb(pb *jobspb.RunParametersPb) (*RunParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunParameters{}
	st.DbtCommands = pb.DbtCommands
	st.JarParams = pb.JarParams
	st.NotebookParams = pb.NotebookParams
	pipelineParamsField, err := PipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}
	st.PythonNamedParams = pb.PythonNamedParams
	st.PythonParams = pb.PythonParams
	st.SparkSubmitParams = pb.SparkSubmitParams
	st.SqlParams = pb.SqlParams

	return st, nil
}

// A value indicating the run's result. The possible values are: * `SUCCESS`:
// The task completed successfully. * `FAILED`: The task completed with an
// error. * `TIMEDOUT`: The run was stopped after reaching the timeout. *
// `CANCELED`: The run was canceled at user request. *
// `MAXIMUM_CONCURRENT_RUNS_REACHED`: The run was skipped because the maximum
// concurrent runs were reached. * `EXCLUDED`: The run was skipped because the
// necessary conditions were not met. * `SUCCESS_WITH_FAILURES`: The job run
// completed successfully with some failures; leaf tasks were successful. *
// `UPSTREAM_FAILED`: The run was skipped because of an upstream failure. *
// `UPSTREAM_CANCELED`: The run was skipped because an upstream task was
// canceled. * `DISABLED`: The run was skipped because it was disabled
// explicitly by the user.
type RunResultState string

// The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

// The run was skipped because it was disabled explicitly by the user.
const RunResultStateDisabled RunResultState = `DISABLED`

// The run was skipped because the necessary conditions were not met.
const RunResultStateExcluded RunResultState = `EXCLUDED`

// The task completed with an error.
const RunResultStateFailed RunResultState = `FAILED`

// The run was skipped because the maximum concurrent runs were reached.
const RunResultStateMaximumConcurrentRunsReached RunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

// The task completed successfully.
const RunResultStateSuccess RunResultState = `SUCCESS`

// The job run completed successfully with some failures; leaf tasks were
// successful.
const RunResultStateSuccessWithFailures RunResultState = `SUCCESS_WITH_FAILURES`

// The run was stopped after reaching the timeout.
const RunResultStateTimedout RunResultState = `TIMEDOUT`

// The run was skipped because an upstream task was canceled.
const RunResultStateUpstreamCanceled RunResultState = `UPSTREAM_CANCELED`

// The run was skipped because of an upstream failure.
const RunResultStateUpstreamFailed RunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *RunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_FAILED`:
		*f = RunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_FAILED"`, v)
	}
}

// Values returns all possible values for RunResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunResultState) Values() []RunResultState {
	return []RunResultState{
		RunResultStateCanceled,
		RunResultStateDisabled,
		RunResultStateExcluded,
		RunResultStateFailed,
		RunResultStateMaximumConcurrentRunsReached,
		RunResultStateSuccess,
		RunResultStateSuccessWithFailures,
		RunResultStateTimedout,
		RunResultStateUpstreamCanceled,
		RunResultStateUpstreamFailed,
	}
}

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

func RunResultStateToPb(st *RunResultState) (*jobspb.RunResultStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.RunResultStatePb(*st)
	return &pb, nil
}

func RunResultStateFromPb(pb *jobspb.RunResultStatePb) (*RunResultState, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunResultState(*pb)
	return &st, nil
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	// Wire name: 'life_cycle_state'
	LifeCycleState RunLifeCycleState ``
	// The reason indicating why the run was queued.
	// Wire name: 'queue_reason'
	QueueReason string ``
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	// Wire name: 'result_state'
	ResultState RunResultState ``
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	// Wire name: 'state_message'
	StateMessage string ``
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	// Wire name: 'user_cancelled_or_timedout'
	UserCancelledOrTimedout bool     ``
	ForceSendFields         []string `tf:"-"`
}

func (st RunState) MarshalJSON() ([]byte, error) {
	pb, err := RunStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunStateToPb(st *RunState) (*jobspb.RunStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunStatePb{}
	lifeCycleStatePb, err := RunLifeCycleStateToPb(&st.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStatePb != nil {
		pb.LifeCycleState = *lifeCycleStatePb
	}
	pb.QueueReason = st.QueueReason
	resultStatePb, err := RunResultStateToPb(&st.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStatePb != nil {
		pb.ResultState = *resultStatePb
	}
	pb.StateMessage = st.StateMessage
	pb.UserCancelledOrTimedout = st.UserCancelledOrTimedout

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunStateFromPb(pb *jobspb.RunStatePb) (*RunState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunState{}
	lifeCycleStateField, err := RunLifeCycleStateFromPb(&pb.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStateField != nil {
		st.LifeCycleState = *lifeCycleStateField
	}
	st.QueueReason = pb.QueueReason
	resultStateField, err := RunResultStateFromPb(&pb.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStateField != nil {
		st.ResultState = *resultStateField
	}
	st.StateMessage = pb.StateMessage
	st.UserCancelledOrTimedout = pb.UserCancelledOrTimedout

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The current status of the run
type RunStatus struct {
	// If the run was queued, details about the reason for queuing the run.
	// Wire name: 'queue_details'
	QueueDetails *QueueDetails ``

	// Wire name: 'state'
	State RunLifecycleStateV2State ``
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	// Wire name: 'termination_details'
	TerminationDetails *TerminationDetails ``
}

func (st RunStatus) MarshalJSON() ([]byte, error) {
	pb, err := RunStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunStatusToPb(st *RunStatus) (*jobspb.RunStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunStatusPb{}
	queueDetailsPb, err := QueueDetailsToPb(st.QueueDetails)
	if err != nil {
		return nil, err
	}
	if queueDetailsPb != nil {
		pb.QueueDetails = queueDetailsPb
	}
	statePb, err := RunLifecycleStateV2StateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	terminationDetailsPb, err := TerminationDetailsToPb(st.TerminationDetails)
	if err != nil {
		return nil, err
	}
	if terminationDetailsPb != nil {
		pb.TerminationDetails = terminationDetailsPb
	}

	return pb, nil
}

func RunStatusFromPb(pb *jobspb.RunStatusPb) (*RunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunStatus{}
	queueDetailsField, err := QueueDetailsFromPb(pb.QueueDetails)
	if err != nil {
		return nil, err
	}
	if queueDetailsField != nil {
		st.QueueDetails = queueDetailsField
	}
	stateField, err := RunLifecycleStateV2StateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	terminationDetailsField, err := TerminationDetailsFromPb(pb.TerminationDetails)
	if err != nil {
		return nil, err
	}
	if terminationDetailsField != nil {
		st.TerminationDetails = terminationDetailsField
	}

	return st, nil
}

// Used when outputting a child run, in GetRun or ListRuns.
type RunTask struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	// Wire name: 'attempt_number'
	AttemptNumber int ``
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	// Wire name: 'clean_rooms_notebook_task'
	CleanRoomsNotebookTask *CleanRoomsNotebookTask ``
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	// Wire name: 'cleanup_duration'
	CleanupDuration int64 ``
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	// Wire name: 'cluster_instance'
	ClusterInstance *ClusterInstance ``
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	// Wire name: 'condition_task'
	ConditionTask *RunConditionTask ``
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	// Wire name: 'dashboard_task'
	DashboardTask *DashboardTask ``
	// Task type for dbt cloud, deprecated in favor of the new name
	// dbt_platform_task
	// Wire name: 'dbt_cloud_task'
	DbtCloudTask *DbtCloudTask ``

	// Wire name: 'dbt_platform_task'
	DbtPlatformTask *DbtPlatformTask ``
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	// Wire name: 'dbt_task'
	DbtTask *DbtTask ``
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	// Wire name: 'depends_on'
	DependsOn []TaskDependency ``
	// An optional description for this task.
	// Wire name: 'description'
	Description string ``
	// Deprecated, field was never used in production.
	// Wire name: 'disabled'
	Disabled bool ``
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget ``
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications ``
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	// Wire name: 'end_time'
	EndTime int64 ``
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	// Wire name: 'environment_key'
	EnvironmentKey string ``
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	// Wire name: 'execution_duration'
	ExecutionDuration int64 ``
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string ``
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	// Wire name: 'for_each_task'
	ForEachTask *RunForEachTask ``

	// Wire name: 'gen_ai_compute_task'
	GenAiComputeTask *GenAiComputeTask ``
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource ``
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	// Wire name: 'job_cluster_key'
	JobClusterKey string ``
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library ``
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec ``
	// The task runs a notebook when the `notebook_task` field is present.
	// Wire name: 'notebook_task'
	NotebookTask *NotebookTask ``
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	// Wire name: 'notification_settings'
	NotificationSettings *TaskNotificationSettings ``
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	// Wire name: 'pipeline_task'
	PipelineTask *PipelineTask ``
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	// Wire name: 'power_bi_task'
	PowerBiTask *PowerBiTask ``
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	// Wire name: 'python_wheel_task'
	PythonWheelTask *PythonWheelTask ``
	// The time in milliseconds that the run has spent in the queue.
	// Wire name: 'queue_duration'
	QueueDuration int64 ``
	// Parameter values including resolved references
	// Wire name: 'resolved_values'
	ResolvedValues *ResolvedValues ``
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	// Wire name: 'run_duration'
	RunDuration int64 ``
	// The ID of the task run.
	// Wire name: 'run_id'
	RunId int64 ``
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	// Wire name: 'run_if'
	RunIf RunIf ``
	// The task triggers another job when the `run_job_task` field is present.
	// Wire name: 'run_job_task'
	RunJobTask *RunJobTask ``

	// Wire name: 'run_page_url'
	RunPageUrl string ``
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	// Wire name: 'setup_duration'
	SetupDuration int64 ``
	// The task runs a JAR when the `spark_jar_task` field is present.
	// Wire name: 'spark_jar_task'
	SparkJarTask *SparkJarTask ``
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	// Wire name: 'spark_python_task'
	SparkPythonTask *SparkPythonTask ``
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	// Wire name: 'spark_submit_task'
	SparkSubmitTask *SparkSubmitTask ``
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	// Wire name: 'sql_task'
	SqlTask *SqlTask ``
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	// Wire name: 'start_time'
	StartTime int64 ``
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState ``

	// Wire name: 'status'
	Status *RunStatus ``
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	// Wire name: 'task_key'
	TaskKey string ``
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int ``
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications ``
	ForceSendFields      []string              `tf:"-"`
}

func (st RunTask) MarshalJSON() ([]byte, error) {
	pb, err := RunTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.RunTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunTaskToPb(st *RunTask) (*jobspb.RunTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.RunTaskPb{}
	pb.AttemptNumber = st.AttemptNumber
	cleanRoomsNotebookTaskPb, err := CleanRoomsNotebookTaskToPb(st.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskPb != nil {
		pb.CleanRoomsNotebookTask = cleanRoomsNotebookTaskPb
	}
	pb.CleanupDuration = st.CleanupDuration
	clusterInstancePb, err := ClusterInstanceToPb(st.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstancePb != nil {
		pb.ClusterInstance = clusterInstancePb
	}
	conditionTaskPb, err := RunConditionTaskToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}
	dashboardTaskPb, err := DashboardTaskToPb(st.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskPb != nil {
		pb.DashboardTask = dashboardTaskPb
	}
	dbtCloudTaskPb, err := DbtCloudTaskToPb(st.DbtCloudTask)
	if err != nil {
		return nil, err
	}
	if dbtCloudTaskPb != nil {
		pb.DbtCloudTask = dbtCloudTaskPb
	}
	dbtPlatformTaskPb, err := DbtPlatformTaskToPb(st.DbtPlatformTask)
	if err != nil {
		return nil, err
	}
	if dbtPlatformTaskPb != nil {
		pb.DbtPlatformTask = dbtPlatformTaskPb
	}
	dbtTaskPb, err := DbtTaskToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	var dependsOnPb []jobspb.TaskDependencyPb
	for _, item := range st.DependsOn {
		itemPb, err := TaskDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependsOnPb = append(dependsOnPb, *itemPb)
		}
	}
	pb.DependsOn = dependsOnPb
	pb.Description = st.Description
	pb.Disabled = st.Disabled
	effectivePerformanceTargetPb, err := PerformanceTargetToPb(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}
	emailNotificationsPb, err := JobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}
	pb.EndTime = st.EndTime
	pb.EnvironmentKey = st.EnvironmentKey
	pb.ExecutionDuration = st.ExecutionDuration
	pb.ExistingClusterId = st.ExistingClusterId
	forEachTaskPb, err := RunForEachTaskToPb(st.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskPb != nil {
		pb.ForEachTask = forEachTaskPb
	}
	genAiComputeTaskPb, err := GenAiComputeTaskToPb(st.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskPb != nil {
		pb.GenAiComputeTask = genAiComputeTaskPb
	}
	gitSourcePb, err := GitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}
	pb.JobClusterKey = st.JobClusterKey

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}
	notebookTaskPb, err := NotebookTaskToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}
	notificationSettingsPb, err := TaskNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}
	pipelineTaskPb, err := PipelineTaskToPb(st.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskPb != nil {
		pb.PipelineTask = pipelineTaskPb
	}
	powerBiTaskPb, err := PowerBiTaskToPb(st.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskPb != nil {
		pb.PowerBiTask = powerBiTaskPb
	}
	pythonWheelTaskPb, err := PythonWheelTaskToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}
	pb.QueueDuration = st.QueueDuration
	resolvedValuesPb, err := ResolvedValuesToPb(st.ResolvedValues)
	if err != nil {
		return nil, err
	}
	if resolvedValuesPb != nil {
		pb.ResolvedValues = resolvedValuesPb
	}
	pb.RunDuration = st.RunDuration
	pb.RunId = st.RunId
	runIfPb, err := RunIfToPb(&st.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfPb != nil {
		pb.RunIf = *runIfPb
	}
	runJobTaskPb, err := RunJobTaskToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}
	pb.RunPageUrl = st.RunPageUrl
	pb.SetupDuration = st.SetupDuration
	sparkJarTaskPb, err := SparkJarTaskToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}
	sparkPythonTaskPb, err := SparkPythonTaskToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}
	sparkSubmitTaskPb, err := SparkSubmitTaskToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}
	sqlTaskPb, err := SqlTaskToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}
	pb.StartTime = st.StartTime
	statePb, err := RunStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	statusPb, err := RunStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}
	pb.TaskKey = st.TaskKey
	pb.TimeoutSeconds = st.TimeoutSeconds
	webhookNotificationsPb, err := WebhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunTaskFromPb(pb *jobspb.RunTaskPb) (*RunTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunTask{}
	st.AttemptNumber = pb.AttemptNumber
	cleanRoomsNotebookTaskField, err := CleanRoomsNotebookTaskFromPb(pb.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskField != nil {
		st.CleanRoomsNotebookTask = cleanRoomsNotebookTaskField
	}
	st.CleanupDuration = pb.CleanupDuration
	clusterInstanceField, err := ClusterInstanceFromPb(pb.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstanceField != nil {
		st.ClusterInstance = clusterInstanceField
	}
	conditionTaskField, err := RunConditionTaskFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dashboardTaskField, err := DashboardTaskFromPb(pb.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskField != nil {
		st.DashboardTask = dashboardTaskField
	}
	dbtCloudTaskField, err := DbtCloudTaskFromPb(pb.DbtCloudTask)
	if err != nil {
		return nil, err
	}
	if dbtCloudTaskField != nil {
		st.DbtCloudTask = dbtCloudTaskField
	}
	dbtPlatformTaskField, err := DbtPlatformTaskFromPb(pb.DbtPlatformTask)
	if err != nil {
		return nil, err
	}
	if dbtPlatformTaskField != nil {
		st.DbtPlatformTask = dbtPlatformTaskField
	}
	dbtTaskField, err := DbtTaskFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}

	var dependsOnField []TaskDependency
	for _, itemPb := range pb.DependsOn {
		item, err := TaskDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependsOnField = append(dependsOnField, *item)
		}
	}
	st.DependsOn = dependsOnField
	st.Description = pb.Description
	st.Disabled = pb.Disabled
	effectivePerformanceTargetField, err := PerformanceTargetFromPb(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	emailNotificationsField, err := JobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}
	st.EndTime = pb.EndTime
	st.EnvironmentKey = pb.EnvironmentKey
	st.ExecutionDuration = pb.ExecutionDuration
	st.ExistingClusterId = pb.ExistingClusterId
	forEachTaskField, err := RunForEachTaskFromPb(pb.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskField != nil {
		st.ForEachTask = forEachTaskField
	}
	genAiComputeTaskField, err := GenAiComputeTaskFromPb(pb.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskField != nil {
		st.GenAiComputeTask = genAiComputeTaskField
	}
	gitSourceField, err := GitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	st.JobClusterKey = pb.JobClusterKey

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}
	notebookTaskField, err := NotebookTaskFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	notificationSettingsField, err := TaskNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	pipelineTaskField, err := PipelineTaskFromPb(pb.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskField != nil {
		st.PipelineTask = pipelineTaskField
	}
	powerBiTaskField, err := PowerBiTaskFromPb(pb.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskField != nil {
		st.PowerBiTask = powerBiTaskField
	}
	pythonWheelTaskField, err := PythonWheelTaskFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	st.QueueDuration = pb.QueueDuration
	resolvedValuesField, err := ResolvedValuesFromPb(pb.ResolvedValues)
	if err != nil {
		return nil, err
	}
	if resolvedValuesField != nil {
		st.ResolvedValues = resolvedValuesField
	}
	st.RunDuration = pb.RunDuration
	st.RunId = pb.RunId
	runIfField, err := RunIfFromPb(&pb.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfField != nil {
		st.RunIf = *runIfField
	}
	runJobTaskField, err := RunJobTaskFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	st.RunPageUrl = pb.RunPageUrl
	st.SetupDuration = pb.SetupDuration
	sparkJarTaskField, err := SparkJarTaskFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := SparkPythonTaskFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := SparkSubmitTaskFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := SqlTaskFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}
	st.StartTime = pb.StartTime
	stateField, err := RunStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := RunStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.TaskKey = pb.TaskKey
	st.TimeoutSeconds = pb.TimeoutSeconds
	webhookNotificationsField, err := WebhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The type of a run. * `JOB_RUN`: Normal job run. A run created with
// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
// :method:jobs/submit.
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
type RunType string

// Normal job run. A run created with :method:jobs/runNow.
const RunTypeJobRun RunType = `JOB_RUN`

// Submit run. A run created with :method:jobs/submit.
const RunTypeSubmitRun RunType = `SUBMIT_RUN`

// Workflow run. A run created with [dbutils.notebook.run].
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
const RunTypeWorkflowRun RunType = `WORKFLOW_RUN`

// String representation for [fmt.Print]
func (f *RunType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunType) Set(v string) error {
	switch v {
	case `JOB_RUN`, `SUBMIT_RUN`, `WORKFLOW_RUN`:
		*f = RunType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "JOB_RUN", "SUBMIT_RUN", "WORKFLOW_RUN"`, v)
	}
}

// Values returns all possible values for RunType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunType) Values() []RunType {
	return []RunType{
		RunTypeJobRun,
		RunTypeSubmitRun,
		RunTypeWorkflowRun,
	}
}

// Type always returns RunType to satisfy [pflag.Value] interface
func (f *RunType) Type() string {
	return "RunType"
}

func RunTypeToPb(st *RunType) (*jobspb.RunTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.RunTypePb(*st)
	return &pb, nil
}

func RunTypeFromPb(pb *jobspb.RunTypePb) (*RunType, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunType(*pb)
	return &st, nil
}

// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL file
// will be retrieved\ from the local Databricks workspace. When set to `GIT`,
// the SQL file will be retrieved from a Git repository defined in `git_source`.
// If the value is empty, the task will use `GIT` if `git_source` is defined and
// `WORKSPACE` otherwise.
//
// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL file
// is located in cloud Git provider.
type Source string

// SQL file is located in cloud Git provider.
const SourceGit Source = `GIT`

// SQL file is located in <Databricks> workspace.
const SourceWorkspace Source = `WORKSPACE`

// String representation for [fmt.Print]
func (f *Source) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Source) Set(v string) error {
	switch v {
	case `GIT`, `WORKSPACE`:
		*f = Source(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GIT", "WORKSPACE"`, v)
	}
}

// Values returns all possible values for Source.
//
// There is no guarantee on the order of the values in the slice.
func (f *Source) Values() []Source {
	return []Source{
		SourceGit,
		SourceWorkspace,
	}
}

// Type always returns Source to satisfy [pflag.Value] interface
func (f *Source) Type() string {
	return "Source"
}

func SourceToPb(st *Source) (*jobspb.SourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.SourcePb(*st)
	return &pb, nil
}

func SourceFromPb(pb *jobspb.SourcePb) (*Source, error) {
	if pb == nil {
		return nil, nil
	}
	st := Source(*pb)
	return &st, nil
}

type SparkJarTask struct {
	// Deprecated since 04/2016. Provide a `jar` through the `libraries` field
	// instead. For an example, see :method:jobs/create.
	// Wire name: 'jar_uri'
	JarUri string ``
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	// Wire name: 'main_class_name'
	MainClassName string ``
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'parameters'
	Parameters []string ``
	// Deprecated. A value of `false` is no longer supported.
	// Wire name: 'run_as_repl'
	RunAsRepl       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st SparkJarTask) MarshalJSON() ([]byte, error) {
	pb, err := SparkJarTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparkJarTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SparkJarTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparkJarTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SparkJarTaskToPb(st *SparkJarTask) (*jobspb.SparkJarTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SparkJarTaskPb{}
	pb.JarUri = st.JarUri
	pb.MainClassName = st.MainClassName
	pb.Parameters = st.Parameters
	pb.RunAsRepl = st.RunAsRepl

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SparkJarTaskFromPb(pb *jobspb.SparkJarTaskPb) (*SparkJarTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkJarTask{}
	st.JarUri = pb.JarUri
	st.MainClassName = pb.MainClassName
	st.Parameters = pb.Parameters
	st.RunAsRepl = pb.RunAsRepl

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'parameters'
	Parameters []string ``
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	// Wire name: 'python_file'
	PythonFile string ``
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local Databricks workspace
	// or cloud location (if the `python_file` has a URI format). When set to
	// `GIT`, the Python file will be retrieved from a Git repository defined in
	// `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a Databricks workspace or at
	// a cloud filesystem URI. * `GIT`: The Python file is located in a remote
	// Git repository.
	// Wire name: 'source'
	Source Source ``
}

func (st SparkPythonTask) MarshalJSON() ([]byte, error) {
	pb, err := SparkPythonTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparkPythonTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SparkPythonTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparkPythonTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SparkPythonTaskToPb(st *SparkPythonTask) (*jobspb.SparkPythonTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SparkPythonTaskPb{}
	pb.Parameters = st.Parameters
	pb.PythonFile = st.PythonFile
	sourcePb, err := SourceToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	return pb, nil
}

func SparkPythonTaskFromPb(pb *jobspb.SparkPythonTaskPb) (*SparkPythonTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkPythonTask{}
	st.Parameters = pb.Parameters
	st.PythonFile = pb.PythonFile
	sourceField, err := SourceFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}

	return st, nil
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'parameters'
	Parameters []string ``
}

func (st SparkSubmitTask) MarshalJSON() ([]byte, error) {
	pb, err := SparkSubmitTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparkSubmitTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SparkSubmitTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparkSubmitTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SparkSubmitTaskToPb(st *SparkSubmitTask) (*jobspb.SparkSubmitTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SparkSubmitTaskPb{}
	pb.Parameters = st.Parameters

	return pb, nil
}

func SparkSubmitTaskFromPb(pb *jobspb.SparkSubmitTaskPb) (*SparkSubmitTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkSubmitTask{}
	st.Parameters = pb.Parameters

	return st, nil
}

type SqlAlertOutput struct {

	// Wire name: 'alert_state'
	AlertState SqlAlertState ``
	// The link to find the output results.
	// Wire name: 'output_link'
	OutputLink string ``
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	// Wire name: 'query_text'
	QueryText string ``
	// Information about SQL statements executed in the run.
	// Wire name: 'sql_statements'
	SqlStatements []SqlStatementOutput ``
	// The canonical identifier of the SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SqlAlertOutput) MarshalJSON() ([]byte, error) {
	pb, err := SqlAlertOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlAlertOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlAlertOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlAlertOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlAlertOutputToPb(st *SqlAlertOutput) (*jobspb.SqlAlertOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlAlertOutputPb{}
	alertStatePb, err := SqlAlertStateToPb(&st.AlertState)
	if err != nil {
		return nil, err
	}
	if alertStatePb != nil {
		pb.AlertState = *alertStatePb
	}
	pb.OutputLink = st.OutputLink
	pb.QueryText = st.QueryText

	var sqlStatementsPb []jobspb.SqlStatementOutputPb
	for _, item := range st.SqlStatements {
		itemPb, err := SqlStatementOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlStatementsPb = append(sqlStatementsPb, *itemPb)
		}
	}
	pb.SqlStatements = sqlStatementsPb
	pb.WarehouseId = st.WarehouseId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlAlertOutputFromPb(pb *jobspb.SqlAlertOutputPb) (*SqlAlertOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlAlertOutput{}
	alertStateField, err := SqlAlertStateFromPb(&pb.AlertState)
	if err != nil {
		return nil, err
	}
	if alertStateField != nil {
		st.AlertState = *alertStateField
	}
	st.OutputLink = pb.OutputLink
	st.QueryText = pb.QueryText

	var sqlStatementsField []SqlStatementOutput
	for _, itemPb := range pb.SqlStatements {
		item, err := SqlStatementOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlStatementsField = append(sqlStatementsField, *item)
		}
	}
	st.SqlStatements = sqlStatementsField
	st.WarehouseId = pb.WarehouseId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The state of the SQL alert.
//
// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled trigger
// conditions
type SqlAlertState string

const SqlAlertStateOk SqlAlertState = `OK`

const SqlAlertStateTriggered SqlAlertState = `TRIGGERED`

const SqlAlertStateUnknown SqlAlertState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *SqlAlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlAlertState) Set(v string) error {
	switch v {
	case `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = SqlAlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for SqlAlertState.
//
// There is no guarantee on the order of the values in the slice.
func (f *SqlAlertState) Values() []SqlAlertState {
	return []SqlAlertState{
		SqlAlertStateOk,
		SqlAlertStateTriggered,
		SqlAlertStateUnknown,
	}
}

// Type always returns SqlAlertState to satisfy [pflag.Value] interface
func (f *SqlAlertState) Type() string {
	return "SqlAlertState"
}

func SqlAlertStateToPb(st *SqlAlertState) (*jobspb.SqlAlertStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.SqlAlertStatePb(*st)
	return &pb, nil
}

func SqlAlertStateFromPb(pb *jobspb.SqlAlertStatePb) (*SqlAlertState, error) {
	if pb == nil {
		return nil, nil
	}
	st := SqlAlertState(*pb)
	return &st, nil
}

type SqlDashboardOutput struct {
	// The canonical identifier of the SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId string ``
	// Widgets executed in the run. Only SQL query based widgets are listed.
	// Wire name: 'widgets'
	Widgets         []SqlDashboardWidgetOutput ``
	ForceSendFields []string                   `tf:"-"`
}

func (st SqlDashboardOutput) MarshalJSON() ([]byte, error) {
	pb, err := SqlDashboardOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlDashboardOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlDashboardOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlDashboardOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlDashboardOutputToPb(st *SqlDashboardOutput) (*jobspb.SqlDashboardOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlDashboardOutputPb{}
	pb.WarehouseId = st.WarehouseId

	var widgetsPb []jobspb.SqlDashboardWidgetOutputPb
	for _, item := range st.Widgets {
		itemPb, err := SqlDashboardWidgetOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			widgetsPb = append(widgetsPb, *itemPb)
		}
	}
	pb.Widgets = widgetsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlDashboardOutputFromPb(pb *jobspb.SqlDashboardOutputPb) (*SqlDashboardOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlDashboardOutput{}
	st.WarehouseId = pb.WarehouseId

	var widgetsField []SqlDashboardWidgetOutput
	for _, itemPb := range pb.Widgets {
		item, err := SqlDashboardWidgetOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			widgetsField = append(widgetsField, *item)
		}
	}
	st.Widgets = widgetsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	// Wire name: 'end_time'
	EndTime int64 ``
	// The information about the error when execution fails.
	// Wire name: 'error'
	Error *SqlOutputError ``
	// The link to find the output results.
	// Wire name: 'output_link'
	OutputLink string ``
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	// Wire name: 'start_time'
	StartTime int64 ``
	// The execution status of the SQL widget.
	// Wire name: 'status'
	Status SqlDashboardWidgetOutputStatus ``
	// The canonical identifier of the SQL widget.
	// Wire name: 'widget_id'
	WidgetId string ``
	// The title of the SQL widget.
	// Wire name: 'widget_title'
	WidgetTitle     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SqlDashboardWidgetOutput) MarshalJSON() ([]byte, error) {
	pb, err := SqlDashboardWidgetOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlDashboardWidgetOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlDashboardWidgetOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlDashboardWidgetOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlDashboardWidgetOutputToPb(st *SqlDashboardWidgetOutput) (*jobspb.SqlDashboardWidgetOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlDashboardWidgetOutputPb{}
	pb.EndTime = st.EndTime
	errorPb, err := SqlOutputErrorToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}
	pb.OutputLink = st.OutputLink
	pb.StartTime = st.StartTime
	statusPb, err := SqlDashboardWidgetOutputStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.WidgetId = st.WidgetId
	pb.WidgetTitle = st.WidgetTitle

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlDashboardWidgetOutputFromPb(pb *jobspb.SqlDashboardWidgetOutputPb) (*SqlDashboardWidgetOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlDashboardWidgetOutput{}
	st.EndTime = pb.EndTime
	errorField, err := SqlOutputErrorFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	st.OutputLink = pb.OutputLink
	st.StartTime = pb.StartTime
	statusField, err := SqlDashboardWidgetOutputStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.WidgetId = pb.WidgetId
	st.WidgetTitle = pb.WidgetTitle

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlDashboardWidgetOutputStatus string

const SqlDashboardWidgetOutputStatusCancelled SqlDashboardWidgetOutputStatus = `CANCELLED`

const SqlDashboardWidgetOutputStatusFailed SqlDashboardWidgetOutputStatus = `FAILED`

const SqlDashboardWidgetOutputStatusPending SqlDashboardWidgetOutputStatus = `PENDING`

const SqlDashboardWidgetOutputStatusRunning SqlDashboardWidgetOutputStatus = `RUNNING`

const SqlDashboardWidgetOutputStatusSuccess SqlDashboardWidgetOutputStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *SqlDashboardWidgetOutputStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlDashboardWidgetOutputStatus) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`:
		*f = SqlDashboardWidgetOutputStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCESS"`, v)
	}
}

// Values returns all possible values for SqlDashboardWidgetOutputStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SqlDashboardWidgetOutputStatus) Values() []SqlDashboardWidgetOutputStatus {
	return []SqlDashboardWidgetOutputStatus{
		SqlDashboardWidgetOutputStatusCancelled,
		SqlDashboardWidgetOutputStatusFailed,
		SqlDashboardWidgetOutputStatusPending,
		SqlDashboardWidgetOutputStatusRunning,
		SqlDashboardWidgetOutputStatusSuccess,
	}
}

// Type always returns SqlDashboardWidgetOutputStatus to satisfy [pflag.Value] interface
func (f *SqlDashboardWidgetOutputStatus) Type() string {
	return "SqlDashboardWidgetOutputStatus"
}

func SqlDashboardWidgetOutputStatusToPb(st *SqlDashboardWidgetOutputStatus) (*jobspb.SqlDashboardWidgetOutputStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.SqlDashboardWidgetOutputStatusPb(*st)
	return &pb, nil
}

func SqlDashboardWidgetOutputStatusFromPb(pb *jobspb.SqlDashboardWidgetOutputStatusPb) (*SqlDashboardWidgetOutputStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SqlDashboardWidgetOutputStatus(*pb)
	return &st, nil
}

type SqlOutput struct {
	// The output of a SQL alert task, if available.
	// Wire name: 'alert_output'
	AlertOutput *SqlAlertOutput ``
	// The output of a SQL dashboard task, if available.
	// Wire name: 'dashboard_output'
	DashboardOutput *SqlDashboardOutput ``
	// The output of a SQL query task, if available.
	// Wire name: 'query_output'
	QueryOutput *SqlQueryOutput ``
}

func (st SqlOutput) MarshalJSON() ([]byte, error) {
	pb, err := SqlOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlOutputToPb(st *SqlOutput) (*jobspb.SqlOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlOutputPb{}
	alertOutputPb, err := SqlAlertOutputToPb(st.AlertOutput)
	if err != nil {
		return nil, err
	}
	if alertOutputPb != nil {
		pb.AlertOutput = alertOutputPb
	}
	dashboardOutputPb, err := SqlDashboardOutputToPb(st.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputPb != nil {
		pb.DashboardOutput = dashboardOutputPb
	}
	queryOutputPb, err := SqlQueryOutputToPb(st.QueryOutput)
	if err != nil {
		return nil, err
	}
	if queryOutputPb != nil {
		pb.QueryOutput = queryOutputPb
	}

	return pb, nil
}

func SqlOutputFromPb(pb *jobspb.SqlOutputPb) (*SqlOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlOutput{}
	alertOutputField, err := SqlAlertOutputFromPb(pb.AlertOutput)
	if err != nil {
		return nil, err
	}
	if alertOutputField != nil {
		st.AlertOutput = alertOutputField
	}
	dashboardOutputField, err := SqlDashboardOutputFromPb(pb.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputField != nil {
		st.DashboardOutput = dashboardOutputField
	}
	queryOutputField, err := SqlQueryOutputFromPb(pb.QueryOutput)
	if err != nil {
		return nil, err
	}
	if queryOutputField != nil {
		st.QueryOutput = queryOutputField
	}

	return st, nil
}

type SqlOutputError struct {
	// The error message when execution fails.
	// Wire name: 'message'
	Message         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SqlOutputError) MarshalJSON() ([]byte, error) {
	pb, err := SqlOutputErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlOutputError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlOutputErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlOutputErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlOutputErrorToPb(st *SqlOutputError) (*jobspb.SqlOutputErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlOutputErrorPb{}
	pb.Message = st.Message

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlOutputErrorFromPb(pb *jobspb.SqlOutputErrorPb) (*SqlOutputError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlOutputError{}
	st.Message = pb.Message

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlQueryOutput struct {

	// Wire name: 'endpoint_id'
	EndpointId string ``
	// The link to find the output results.
	// Wire name: 'output_link'
	OutputLink string ``
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	// Wire name: 'query_text'
	QueryText string ``
	// Information about SQL statements executed in the run.
	// Wire name: 'sql_statements'
	SqlStatements []SqlStatementOutput ``
	// The canonical identifier of the SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SqlQueryOutput) MarshalJSON() ([]byte, error) {
	pb, err := SqlQueryOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlQueryOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlQueryOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlQueryOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlQueryOutputToPb(st *SqlQueryOutput) (*jobspb.SqlQueryOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlQueryOutputPb{}
	pb.EndpointId = st.EndpointId
	pb.OutputLink = st.OutputLink
	pb.QueryText = st.QueryText

	var sqlStatementsPb []jobspb.SqlStatementOutputPb
	for _, item := range st.SqlStatements {
		itemPb, err := SqlStatementOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlStatementsPb = append(sqlStatementsPb, *itemPb)
		}
	}
	pb.SqlStatements = sqlStatementsPb
	pb.WarehouseId = st.WarehouseId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlQueryOutputFromPb(pb *jobspb.SqlQueryOutputPb) (*SqlQueryOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlQueryOutput{}
	st.EndpointId = pb.EndpointId
	st.OutputLink = pb.OutputLink
	st.QueryText = pb.QueryText

	var sqlStatementsField []SqlStatementOutput
	for _, itemPb := range pb.SqlStatements {
		item, err := SqlStatementOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlStatementsField = append(sqlStatementsField, *item)
		}
	}
	st.SqlStatements = sqlStatementsField
	st.WarehouseId = pb.WarehouseId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	// Wire name: 'lookup_key'
	LookupKey       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SqlStatementOutput) MarshalJSON() ([]byte, error) {
	pb, err := SqlStatementOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlStatementOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlStatementOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlStatementOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlStatementOutputToPb(st *SqlStatementOutput) (*jobspb.SqlStatementOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlStatementOutputPb{}
	pb.LookupKey = st.LookupKey

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlStatementOutputFromPb(pb *jobspb.SqlStatementOutputPb) (*SqlStatementOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlStatementOutput{}
	st.LookupKey = pb.LookupKey

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	// Wire name: 'alert'
	Alert *SqlTaskAlert ``
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	// Wire name: 'dashboard'
	Dashboard *SqlTaskDashboard ``
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	// Wire name: 'file'
	File *SqlTaskFile ``
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	// Wire name: 'parameters'
	Parameters map[string]string ``
	// If query, indicates that this job must execute a SQL query.
	// Wire name: 'query'
	Query *SqlTaskQuery ``
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	// Wire name: 'warehouse_id'
	WarehouseId string ``
}

func (st SqlTask) MarshalJSON() ([]byte, error) {
	pb, err := SqlTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlTaskToPb(st *SqlTask) (*jobspb.SqlTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlTaskPb{}
	alertPb, err := SqlTaskAlertToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}
	dashboardPb, err := SqlTaskDashboardToPb(st.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardPb != nil {
		pb.Dashboard = dashboardPb
	}
	filePb, err := SqlTaskFileToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}
	pb.Parameters = st.Parameters
	queryPb, err := SqlTaskQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

func SqlTaskFromPb(pb *jobspb.SqlTaskPb) (*SqlTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTask{}
	alertField, err := SqlTaskAlertFromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}
	dashboardField, err := SqlTaskDashboardFromPb(pb.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardField != nil {
		st.Dashboard = dashboardField
	}
	fileField, err := SqlTaskFileFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	st.Parameters = pb.Parameters
	queryField, err := SqlTaskQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	st.WarehouseId = pb.WarehouseId

	return st, nil
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	// Wire name: 'alert_id'
	AlertId string ``
	// If true, the alert notifications are not sent to subscribers.
	// Wire name: 'pause_subscriptions'
	PauseSubscriptions bool ``
	// If specified, alert notifications are sent to subscribers.
	// Wire name: 'subscriptions'
	Subscriptions   []SqlTaskSubscription ``
	ForceSendFields []string              `tf:"-"`
}

func (st SqlTaskAlert) MarshalJSON() ([]byte, error) {
	pb, err := SqlTaskAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlTaskAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlTaskAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlTaskAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlTaskAlertToPb(st *SqlTaskAlert) (*jobspb.SqlTaskAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlTaskAlertPb{}
	pb.AlertId = st.AlertId
	pb.PauseSubscriptions = st.PauseSubscriptions

	var subscriptionsPb []jobspb.SqlTaskSubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := SqlTaskSubscriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscriptionsPb = append(subscriptionsPb, *itemPb)
		}
	}
	pb.Subscriptions = subscriptionsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlTaskAlertFromPb(pb *jobspb.SqlTaskAlertPb) (*SqlTaskAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskAlert{}
	st.AlertId = pb.AlertId
	st.PauseSubscriptions = pb.PauseSubscriptions

	var subscriptionsField []SqlTaskSubscription
	for _, itemPb := range pb.Subscriptions {
		item, err := SqlTaskSubscriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscriptionsField = append(subscriptionsField, *item)
		}
	}
	st.Subscriptions = subscriptionsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// The canonical identifier of the SQL dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	// Wire name: 'pause_subscriptions'
	PauseSubscriptions bool ``
	// If specified, dashboard snapshots are sent to subscriptions.
	// Wire name: 'subscriptions'
	Subscriptions   []SqlTaskSubscription ``
	ForceSendFields []string              `tf:"-"`
}

func (st SqlTaskDashboard) MarshalJSON() ([]byte, error) {
	pb, err := SqlTaskDashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlTaskDashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlTaskDashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlTaskDashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlTaskDashboardToPb(st *SqlTaskDashboard) (*jobspb.SqlTaskDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlTaskDashboardPb{}
	pb.CustomSubject = st.CustomSubject
	pb.DashboardId = st.DashboardId
	pb.PauseSubscriptions = st.PauseSubscriptions

	var subscriptionsPb []jobspb.SqlTaskSubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := SqlTaskSubscriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscriptionsPb = append(subscriptionsPb, *itemPb)
		}
	}
	pb.Subscriptions = subscriptionsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlTaskDashboardFromPb(pb *jobspb.SqlTaskDashboardPb) (*SqlTaskDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskDashboard{}
	st.CustomSubject = pb.CustomSubject
	st.DashboardId = pb.DashboardId
	st.PauseSubscriptions = pb.PauseSubscriptions

	var subscriptionsField []SqlTaskSubscription
	for _, itemPb := range pb.Subscriptions {
		item, err := SqlTaskSubscriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscriptionsField = append(subscriptionsField, *item)
		}
	}
	st.Subscriptions = subscriptionsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	// Wire name: 'path'
	Path string ``
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	// Wire name: 'source'
	Source Source ``
}

func (st SqlTaskFile) MarshalJSON() ([]byte, error) {
	pb, err := SqlTaskFileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlTaskFile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlTaskFilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlTaskFileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlTaskFileToPb(st *SqlTaskFile) (*jobspb.SqlTaskFilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlTaskFilePb{}
	pb.Path = st.Path
	sourcePb, err := SourceToPb(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	return pb, nil
}

func SqlTaskFileFromPb(pb *jobspb.SqlTaskFilePb) (*SqlTaskFile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskFile{}
	st.Path = pb.Path
	sourceField, err := SourceFromPb(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}

	return st, nil
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	// Wire name: 'query_id'
	QueryId string ``
}

func (st SqlTaskQuery) MarshalJSON() ([]byte, error) {
	pb, err := SqlTaskQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlTaskQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlTaskQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlTaskQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlTaskQueryToPb(st *SqlTaskQuery) (*jobspb.SqlTaskQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlTaskQueryPb{}
	pb.QueryId = st.QueryId

	return pb, nil
}

func SqlTaskQueryFromPb(pb *jobspb.SqlTaskQueryPb) (*SqlTaskQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskQuery{}
	st.QueryId = pb.QueryId

	return st, nil
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	// Wire name: 'destination_id'
	DestinationId string ``
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SqlTaskSubscription) MarshalJSON() ([]byte, error) {
	pb, err := SqlTaskSubscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlTaskSubscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SqlTaskSubscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlTaskSubscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlTaskSubscriptionToPb(st *SqlTaskSubscription) (*jobspb.SqlTaskSubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SqlTaskSubscriptionPb{}
	pb.DestinationId = st.DestinationId
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlTaskSubscriptionFromPb(pb *jobspb.SqlTaskSubscriptionPb) (*SqlTaskSubscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskSubscription{}
	st.DestinationId = pb.DestinationId
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type StorageMode string

const StorageModeDirectQuery StorageMode = `DIRECT_QUERY`

const StorageModeDual StorageMode = `DUAL`

const StorageModeImport StorageMode = `IMPORT`

// String representation for [fmt.Print]
func (f *StorageMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StorageMode) Set(v string) error {
	switch v {
	case `DIRECT_QUERY`, `DUAL`, `IMPORT`:
		*f = StorageMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT_QUERY", "DUAL", "IMPORT"`, v)
	}
}

// Values returns all possible values for StorageMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *StorageMode) Values() []StorageMode {
	return []StorageMode{
		StorageModeDirectQuery,
		StorageModeDual,
		StorageModeImport,
	}
}

// Type always returns StorageMode to satisfy [pflag.Value] interface
func (f *StorageMode) Type() string {
	return "StorageMode"
}

func StorageModeToPb(st *StorageMode) (*jobspb.StorageModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.StorageModePb(*st)
	return &pb, nil
}

func StorageModeFromPb(pb *jobspb.StorageModePb) (*StorageMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := StorageMode(*pb)
	return &st, nil
}

type SubmitRun struct {
	// List of permissions to set on the job.
	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlRequest ``
	// The user specified id of the budget policy to use for this one-time run.
	// If not specified, the run will be not be attributed to any budget policy.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// An optional set of email addresses notified when the run begins or
	// completes.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications ``
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	// Wire name: 'environments'
	Environments []JobEnvironment ``
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource ``

	// Wire name: 'health'
	Health *JobsHealthRules ``
	// An optional token that can be used to guarantee the idempotency of job
	// run requests. If a run with the provided token already exists, the
	// request does not create a new run but returns the ID of the existing run
	// instead. If a run with the provided token is deleted, an error is
	// returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	// Wire name: 'idempotency_token'
	IdempotencyToken string ``
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	// Wire name: 'notification_settings'
	NotificationSettings *JobNotificationSettings ``
	// The queue settings of the one-time run.
	// Wire name: 'queue'
	Queue *QueueSettings ``
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	// Wire name: 'run_as'
	RunAs *JobRunAs ``
	// An optional name for the run. The default value is `Untitled`.
	// Wire name: 'run_name'
	RunName string ``

	// Wire name: 'tasks'
	Tasks []SubmitTask ``
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int ``
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications ``
	ForceSendFields      []string              `tf:"-"`
}

func (st SubmitRun) MarshalJSON() ([]byte, error) {
	pb, err := SubmitRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SubmitRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SubmitRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SubmitRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SubmitRunToPb(st *SubmitRun) (*jobspb.SubmitRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SubmitRunPb{}

	var accessControlListPb []jobspb.JobAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := JobAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.BudgetPolicyId = st.BudgetPolicyId
	emailNotificationsPb, err := JobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	var environmentsPb []jobspb.JobEnvironmentPb
	for _, item := range st.Environments {
		itemPb, err := JobEnvironmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentsPb = append(environmentsPb, *itemPb)
		}
	}
	pb.Environments = environmentsPb
	gitSourcePb, err := GitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}
	healthPb, err := JobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}
	pb.IdempotencyToken = st.IdempotencyToken
	notificationSettingsPb, err := JobNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}
	queuePb, err := QueueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}
	runAsPb, err := JobRunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}
	pb.RunName = st.RunName

	var tasksPb []jobspb.SubmitTaskPb
	for _, item := range st.Tasks {
		itemPb, err := SubmitTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb
	pb.TimeoutSeconds = st.TimeoutSeconds
	webhookNotificationsPb, err := WebhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SubmitRunFromPb(pb *jobspb.SubmitRunPb) (*SubmitRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitRun{}

	var accessControlListField []JobAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := JobAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.BudgetPolicyId = pb.BudgetPolicyId
	emailNotificationsField, err := JobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}

	var environmentsField []JobEnvironment
	for _, itemPb := range pb.Environments {
		item, err := JobEnvironmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentsField = append(environmentsField, *item)
		}
	}
	st.Environments = environmentsField
	gitSourceField, err := GitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	healthField, err := JobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	st.IdempotencyToken = pb.IdempotencyToken
	notificationSettingsField, err := JobNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	queueField, err := QueueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	runAsField, err := JobRunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	st.RunName = pb.RunName

	var tasksField []SubmitTask
	for _, itemPb := range pb.Tasks {
		item, err := SubmitTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	st.TimeoutSeconds = pb.TimeoutSeconds
	webhookNotificationsField, err := WebhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Run was created and started successfully.
type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	// Wire name: 'run_id'
	RunId           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st SubmitRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := SubmitRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SubmitRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SubmitRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SubmitRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SubmitRunResponseToPb(st *SubmitRunResponse) (*jobspb.SubmitRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SubmitRunResponsePb{}
	pb.RunId = st.RunId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SubmitRunResponseFromPb(pb *jobspb.SubmitRunResponsePb) (*SubmitRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitRunResponse{}
	st.RunId = pb.RunId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SubmitTask struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	// Wire name: 'clean_rooms_notebook_task'
	CleanRoomsNotebookTask *CleanRoomsNotebookTask ``
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	// Wire name: 'condition_task'
	ConditionTask *ConditionTask ``
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	// Wire name: 'dashboard_task'
	DashboardTask *DashboardTask ``
	// Task type for dbt cloud, deprecated in favor of the new name
	// dbt_platform_task
	// Wire name: 'dbt_cloud_task'
	DbtCloudTask *DbtCloudTask ``

	// Wire name: 'dbt_platform_task'
	DbtPlatformTask *DbtPlatformTask ``
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	// Wire name: 'dbt_task'
	DbtTask *DbtTask ``
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	// Wire name: 'depends_on'
	DependsOn []TaskDependency ``
	// An optional description for this task.
	// Wire name: 'description'
	Description string ``
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications ``
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	// Wire name: 'environment_key'
	EnvironmentKey string ``
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string ``
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	// Wire name: 'for_each_task'
	ForEachTask *ForEachTask ``

	// Wire name: 'gen_ai_compute_task'
	GenAiComputeTask *GenAiComputeTask ``

	// Wire name: 'health'
	Health *JobsHealthRules ``
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library ``
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec ``
	// The task runs a notebook when the `notebook_task` field is present.
	// Wire name: 'notebook_task'
	NotebookTask *NotebookTask ``
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	// Wire name: 'notification_settings'
	NotificationSettings *TaskNotificationSettings ``
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	// Wire name: 'pipeline_task'
	PipelineTask *PipelineTask ``
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	// Wire name: 'power_bi_task'
	PowerBiTask *PowerBiTask ``
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	// Wire name: 'python_wheel_task'
	PythonWheelTask *PythonWheelTask ``
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	// Wire name: 'run_if'
	RunIf RunIf ``
	// The task triggers another job when the `run_job_task` field is present.
	// Wire name: 'run_job_task'
	RunJobTask *RunJobTask ``
	// The task runs a JAR when the `spark_jar_task` field is present.
	// Wire name: 'spark_jar_task'
	SparkJarTask *SparkJarTask ``
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	// Wire name: 'spark_python_task'
	SparkPythonTask *SparkPythonTask ``
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	// Wire name: 'spark_submit_task'
	SparkSubmitTask *SparkSubmitTask ``
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	// Wire name: 'sql_task'
	SqlTask *SqlTask ``
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	// Wire name: 'task_key'
	TaskKey string ``
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int ``
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications ``
	ForceSendFields      []string              `tf:"-"`
}

func (st SubmitTask) MarshalJSON() ([]byte, error) {
	pb, err := SubmitTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SubmitTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SubmitTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SubmitTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SubmitTaskToPb(st *SubmitTask) (*jobspb.SubmitTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SubmitTaskPb{}
	cleanRoomsNotebookTaskPb, err := CleanRoomsNotebookTaskToPb(st.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskPb != nil {
		pb.CleanRoomsNotebookTask = cleanRoomsNotebookTaskPb
	}
	conditionTaskPb, err := ConditionTaskToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}
	dashboardTaskPb, err := DashboardTaskToPb(st.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskPb != nil {
		pb.DashboardTask = dashboardTaskPb
	}
	dbtCloudTaskPb, err := DbtCloudTaskToPb(st.DbtCloudTask)
	if err != nil {
		return nil, err
	}
	if dbtCloudTaskPb != nil {
		pb.DbtCloudTask = dbtCloudTaskPb
	}
	dbtPlatformTaskPb, err := DbtPlatformTaskToPb(st.DbtPlatformTask)
	if err != nil {
		return nil, err
	}
	if dbtPlatformTaskPb != nil {
		pb.DbtPlatformTask = dbtPlatformTaskPb
	}
	dbtTaskPb, err := DbtTaskToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	var dependsOnPb []jobspb.TaskDependencyPb
	for _, item := range st.DependsOn {
		itemPb, err := TaskDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependsOnPb = append(dependsOnPb, *itemPb)
		}
	}
	pb.DependsOn = dependsOnPb
	pb.Description = st.Description
	emailNotificationsPb, err := JobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}
	pb.EnvironmentKey = st.EnvironmentKey
	pb.ExistingClusterId = st.ExistingClusterId
	forEachTaskPb, err := ForEachTaskToPb(st.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskPb != nil {
		pb.ForEachTask = forEachTaskPb
	}
	genAiComputeTaskPb, err := GenAiComputeTaskToPb(st.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskPb != nil {
		pb.GenAiComputeTask = genAiComputeTaskPb
	}
	healthPb, err := JobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}
	notebookTaskPb, err := NotebookTaskToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}
	notificationSettingsPb, err := TaskNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}
	pipelineTaskPb, err := PipelineTaskToPb(st.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskPb != nil {
		pb.PipelineTask = pipelineTaskPb
	}
	powerBiTaskPb, err := PowerBiTaskToPb(st.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskPb != nil {
		pb.PowerBiTask = powerBiTaskPb
	}
	pythonWheelTaskPb, err := PythonWheelTaskToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}
	runIfPb, err := RunIfToPb(&st.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfPb != nil {
		pb.RunIf = *runIfPb
	}
	runJobTaskPb, err := RunJobTaskToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}
	sparkJarTaskPb, err := SparkJarTaskToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}
	sparkPythonTaskPb, err := SparkPythonTaskToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}
	sparkSubmitTaskPb, err := SparkSubmitTaskToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}
	sqlTaskPb, err := SqlTaskToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}
	pb.TaskKey = st.TaskKey
	pb.TimeoutSeconds = st.TimeoutSeconds
	webhookNotificationsPb, err := WebhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SubmitTaskFromPb(pb *jobspb.SubmitTaskPb) (*SubmitTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitTask{}
	cleanRoomsNotebookTaskField, err := CleanRoomsNotebookTaskFromPb(pb.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskField != nil {
		st.CleanRoomsNotebookTask = cleanRoomsNotebookTaskField
	}
	conditionTaskField, err := ConditionTaskFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dashboardTaskField, err := DashboardTaskFromPb(pb.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskField != nil {
		st.DashboardTask = dashboardTaskField
	}
	dbtCloudTaskField, err := DbtCloudTaskFromPb(pb.DbtCloudTask)
	if err != nil {
		return nil, err
	}
	if dbtCloudTaskField != nil {
		st.DbtCloudTask = dbtCloudTaskField
	}
	dbtPlatformTaskField, err := DbtPlatformTaskFromPb(pb.DbtPlatformTask)
	if err != nil {
		return nil, err
	}
	if dbtPlatformTaskField != nil {
		st.DbtPlatformTask = dbtPlatformTaskField
	}
	dbtTaskField, err := DbtTaskFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}

	var dependsOnField []TaskDependency
	for _, itemPb := range pb.DependsOn {
		item, err := TaskDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependsOnField = append(dependsOnField, *item)
		}
	}
	st.DependsOn = dependsOnField
	st.Description = pb.Description
	emailNotificationsField, err := JobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}
	st.EnvironmentKey = pb.EnvironmentKey
	st.ExistingClusterId = pb.ExistingClusterId
	forEachTaskField, err := ForEachTaskFromPb(pb.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskField != nil {
		st.ForEachTask = forEachTaskField
	}
	genAiComputeTaskField, err := GenAiComputeTaskFromPb(pb.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskField != nil {
		st.GenAiComputeTask = genAiComputeTaskField
	}
	healthField, err := JobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}
	notebookTaskField, err := NotebookTaskFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	notificationSettingsField, err := TaskNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	pipelineTaskField, err := PipelineTaskFromPb(pb.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskField != nil {
		st.PipelineTask = pipelineTaskField
	}
	powerBiTaskField, err := PowerBiTaskFromPb(pb.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskField != nil {
		st.PowerBiTask = powerBiTaskField
	}
	pythonWheelTaskField, err := PythonWheelTaskFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	runIfField, err := RunIfFromPb(&pb.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfField != nil {
		st.RunIf = *runIfField
	}
	runJobTaskField, err := RunJobTaskFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	sparkJarTaskField, err := SparkJarTaskFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := SparkPythonTaskFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := SparkSubmitTaskFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := SqlTaskFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}
	st.TaskKey = pb.TaskKey
	st.TimeoutSeconds = pb.TimeoutSeconds
	webhookNotificationsField, err := WebhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Subscription struct {
	// Optional: Allows users to specify a custom subject line on the email sent
	// to subscribers.
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// When true, the subscription will not send emails.
	// Wire name: 'paused'
	Paused bool ``
	// The list of subscribers to send the snapshot of the dashboard to.
	// Wire name: 'subscribers'
	Subscribers     []SubscriptionSubscriber ``
	ForceSendFields []string                 `tf:"-"`
}

func (st Subscription) MarshalJSON() ([]byte, error) {
	pb, err := SubscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Subscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SubscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SubscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SubscriptionToPb(st *Subscription) (*jobspb.SubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SubscriptionPb{}
	pb.CustomSubject = st.CustomSubject
	pb.Paused = st.Paused

	var subscribersPb []jobspb.SubscriptionSubscriberPb
	for _, item := range st.Subscribers {
		itemPb, err := SubscriptionSubscriberToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscribersPb = append(subscribersPb, *itemPb)
		}
	}
	pb.Subscribers = subscribersPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SubscriptionFromPb(pb *jobspb.SubscriptionPb) (*Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscription{}
	st.CustomSubject = pb.CustomSubject
	st.Paused = pb.Paused

	var subscribersField []SubscriptionSubscriber
	for _, itemPb := range pb.Subscribers {
		item, err := SubscriptionSubscriberFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscribersField = append(subscribersField, *item)
		}
	}
	st.Subscribers = subscribersField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SubscriptionSubscriber struct {
	// A snapshot of the dashboard will be sent to the destination when the
	// `destination_id` field is present.
	// Wire name: 'destination_id'
	DestinationId string ``
	// A snapshot of the dashboard will be sent to the user's email when the
	// `user_name` field is present.
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SubscriptionSubscriber) MarshalJSON() ([]byte, error) {
	pb, err := SubscriptionSubscriberToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SubscriptionSubscriber) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.SubscriptionSubscriberPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SubscriptionSubscriberFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SubscriptionSubscriberToPb(st *SubscriptionSubscriber) (*jobspb.SubscriptionSubscriberPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.SubscriptionSubscriberPb{}
	pb.DestinationId = st.DestinationId
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SubscriptionSubscriberFromPb(pb *jobspb.SubscriptionSubscriberPb) (*SubscriptionSubscriber, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriber{}
	st.DestinationId = pb.DestinationId
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TableUpdateTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	// Wire name: 'condition'
	Condition Condition ``
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	// Wire name: 'min_time_between_triggers_seconds'
	MinTimeBetweenTriggersSeconds int ``
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	// Wire name: 'table_names'
	TableNames []string ``
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	// Wire name: 'wait_after_last_change_seconds'
	WaitAfterLastChangeSeconds int      ``
	ForceSendFields            []string `tf:"-"`
}

func (st TableUpdateTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := TableUpdateTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TableUpdateTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TableUpdateTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TableUpdateTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TableUpdateTriggerConfigurationToPb(st *TableUpdateTriggerConfiguration) (*jobspb.TableUpdateTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TableUpdateTriggerConfigurationPb{}
	conditionPb, err := ConditionToPb(&st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = *conditionPb
	}
	pb.MinTimeBetweenTriggersSeconds = st.MinTimeBetweenTriggersSeconds
	pb.TableNames = st.TableNames
	pb.WaitAfterLastChangeSeconds = st.WaitAfterLastChangeSeconds

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TableUpdateTriggerConfigurationFromPb(pb *jobspb.TableUpdateTriggerConfigurationPb) (*TableUpdateTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableUpdateTriggerConfiguration{}
	conditionField, err := ConditionFromPb(&pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = *conditionField
	}
	st.MinTimeBetweenTriggersSeconds = pb.MinTimeBetweenTriggersSeconds
	st.TableNames = pb.TableNames
	st.WaitAfterLastChangeSeconds = pb.WaitAfterLastChangeSeconds

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Task struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	// Wire name: 'clean_rooms_notebook_task'
	CleanRoomsNotebookTask *CleanRoomsNotebookTask ``
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	// Wire name: 'condition_task'
	ConditionTask *ConditionTask ``
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	// Wire name: 'dashboard_task'
	DashboardTask *DashboardTask ``
	// Task type for dbt cloud, deprecated in favor of the new name
	// dbt_platform_task
	// Wire name: 'dbt_cloud_task'
	DbtCloudTask *DbtCloudTask ``

	// Wire name: 'dbt_platform_task'
	DbtPlatformTask *DbtPlatformTask ``
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	// Wire name: 'dbt_task'
	DbtTask *DbtTask ``
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	// Wire name: 'depends_on'
	DependsOn []TaskDependency ``
	// An optional description for this task.
	// Wire name: 'description'
	Description string ``
	// An option to disable auto optimization in serverless
	// Wire name: 'disable_auto_optimization'
	DisableAutoOptimization bool ``
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	// Wire name: 'email_notifications'
	EmailNotifications *TaskEmailNotifications ``
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	// Wire name: 'environment_key'
	EnvironmentKey string ``
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string ``
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	// Wire name: 'for_each_task'
	ForEachTask *ForEachTask ``

	// Wire name: 'gen_ai_compute_task'
	GenAiComputeTask *GenAiComputeTask ``

	// Wire name: 'health'
	Health *JobsHealthRules ``
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	// Wire name: 'job_cluster_key'
	JobClusterKey string ``
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library ``
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	// Wire name: 'max_retries'
	MaxRetries int ``
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	// Wire name: 'min_retry_interval_millis'
	MinRetryIntervalMillis int ``
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec ``
	// The task runs a notebook when the `notebook_task` field is present.
	// Wire name: 'notebook_task'
	NotebookTask *NotebookTask ``
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	// Wire name: 'notification_settings'
	NotificationSettings *TaskNotificationSettings ``
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	// Wire name: 'pipeline_task'
	PipelineTask *PipelineTask ``
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	// Wire name: 'power_bi_task'
	PowerBiTask *PowerBiTask ``
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	// Wire name: 'python_wheel_task'
	PythonWheelTask *PythonWheelTask ``
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	// Wire name: 'retry_on_timeout'
	RetryOnTimeout bool ``
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	// Wire name: 'run_if'
	RunIf RunIf ``
	// The task triggers another job when the `run_job_task` field is present.
	// Wire name: 'run_job_task'
	RunJobTask *RunJobTask ``
	// The task runs a JAR when the `spark_jar_task` field is present.
	// Wire name: 'spark_jar_task'
	SparkJarTask *SparkJarTask ``
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	// Wire name: 'spark_python_task'
	SparkPythonTask *SparkPythonTask ``
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	// Wire name: 'spark_submit_task'
	SparkSubmitTask *SparkSubmitTask ``
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	// Wire name: 'sql_task'
	SqlTask *SqlTask ``
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	// Wire name: 'task_key'
	TaskKey string ``
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int ``
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications ``
	ForceSendFields      []string              `tf:"-"`
}

func (st Task) MarshalJSON() ([]byte, error) {
	pb, err := TaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Task) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TaskToPb(st *Task) (*jobspb.TaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TaskPb{}
	cleanRoomsNotebookTaskPb, err := CleanRoomsNotebookTaskToPb(st.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskPb != nil {
		pb.CleanRoomsNotebookTask = cleanRoomsNotebookTaskPb
	}
	conditionTaskPb, err := ConditionTaskToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}
	dashboardTaskPb, err := DashboardTaskToPb(st.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskPb != nil {
		pb.DashboardTask = dashboardTaskPb
	}
	dbtCloudTaskPb, err := DbtCloudTaskToPb(st.DbtCloudTask)
	if err != nil {
		return nil, err
	}
	if dbtCloudTaskPb != nil {
		pb.DbtCloudTask = dbtCloudTaskPb
	}
	dbtPlatformTaskPb, err := DbtPlatformTaskToPb(st.DbtPlatformTask)
	if err != nil {
		return nil, err
	}
	if dbtPlatformTaskPb != nil {
		pb.DbtPlatformTask = dbtPlatformTaskPb
	}
	dbtTaskPb, err := DbtTaskToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	var dependsOnPb []jobspb.TaskDependencyPb
	for _, item := range st.DependsOn {
		itemPb, err := TaskDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependsOnPb = append(dependsOnPb, *itemPb)
		}
	}
	pb.DependsOn = dependsOnPb
	pb.Description = st.Description
	pb.DisableAutoOptimization = st.DisableAutoOptimization
	emailNotificationsPb, err := TaskEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}
	pb.EnvironmentKey = st.EnvironmentKey
	pb.ExistingClusterId = st.ExistingClusterId
	forEachTaskPb, err := ForEachTaskToPb(st.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskPb != nil {
		pb.ForEachTask = forEachTaskPb
	}
	genAiComputeTaskPb, err := GenAiComputeTaskToPb(st.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskPb != nil {
		pb.GenAiComputeTask = genAiComputeTaskPb
	}
	healthPb, err := JobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}
	pb.JobClusterKey = st.JobClusterKey

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.MaxRetries = st.MaxRetries
	pb.MinRetryIntervalMillis = st.MinRetryIntervalMillis
	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}
	notebookTaskPb, err := NotebookTaskToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}
	notificationSettingsPb, err := TaskNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}
	pipelineTaskPb, err := PipelineTaskToPb(st.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskPb != nil {
		pb.PipelineTask = pipelineTaskPb
	}
	powerBiTaskPb, err := PowerBiTaskToPb(st.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskPb != nil {
		pb.PowerBiTask = powerBiTaskPb
	}
	pythonWheelTaskPb, err := PythonWheelTaskToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}
	pb.RetryOnTimeout = st.RetryOnTimeout
	runIfPb, err := RunIfToPb(&st.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfPb != nil {
		pb.RunIf = *runIfPb
	}
	runJobTaskPb, err := RunJobTaskToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}
	sparkJarTaskPb, err := SparkJarTaskToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}
	sparkPythonTaskPb, err := SparkPythonTaskToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}
	sparkSubmitTaskPb, err := SparkSubmitTaskToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}
	sqlTaskPb, err := SqlTaskToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}
	pb.TaskKey = st.TaskKey
	pb.TimeoutSeconds = st.TimeoutSeconds
	webhookNotificationsPb, err := WebhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TaskFromPb(pb *jobspb.TaskPb) (*Task, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Task{}
	cleanRoomsNotebookTaskField, err := CleanRoomsNotebookTaskFromPb(pb.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskField != nil {
		st.CleanRoomsNotebookTask = cleanRoomsNotebookTaskField
	}
	conditionTaskField, err := ConditionTaskFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dashboardTaskField, err := DashboardTaskFromPb(pb.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskField != nil {
		st.DashboardTask = dashboardTaskField
	}
	dbtCloudTaskField, err := DbtCloudTaskFromPb(pb.DbtCloudTask)
	if err != nil {
		return nil, err
	}
	if dbtCloudTaskField != nil {
		st.DbtCloudTask = dbtCloudTaskField
	}
	dbtPlatformTaskField, err := DbtPlatformTaskFromPb(pb.DbtPlatformTask)
	if err != nil {
		return nil, err
	}
	if dbtPlatformTaskField != nil {
		st.DbtPlatformTask = dbtPlatformTaskField
	}
	dbtTaskField, err := DbtTaskFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}

	var dependsOnField []TaskDependency
	for _, itemPb := range pb.DependsOn {
		item, err := TaskDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependsOnField = append(dependsOnField, *item)
		}
	}
	st.DependsOn = dependsOnField
	st.Description = pb.Description
	st.DisableAutoOptimization = pb.DisableAutoOptimization
	emailNotificationsField, err := TaskEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}
	st.EnvironmentKey = pb.EnvironmentKey
	st.ExistingClusterId = pb.ExistingClusterId
	forEachTaskField, err := ForEachTaskFromPb(pb.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskField != nil {
		st.ForEachTask = forEachTaskField
	}
	genAiComputeTaskField, err := GenAiComputeTaskFromPb(pb.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskField != nil {
		st.GenAiComputeTask = genAiComputeTaskField
	}
	healthField, err := JobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	st.JobClusterKey = pb.JobClusterKey

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.MaxRetries = pb.MaxRetries
	st.MinRetryIntervalMillis = pb.MinRetryIntervalMillis
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}
	notebookTaskField, err := NotebookTaskFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	notificationSettingsField, err := TaskNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	pipelineTaskField, err := PipelineTaskFromPb(pb.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskField != nil {
		st.PipelineTask = pipelineTaskField
	}
	powerBiTaskField, err := PowerBiTaskFromPb(pb.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskField != nil {
		st.PowerBiTask = powerBiTaskField
	}
	pythonWheelTaskField, err := PythonWheelTaskFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	st.RetryOnTimeout = pb.RetryOnTimeout
	runIfField, err := RunIfFromPb(&pb.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfField != nil {
		st.RunIf = *runIfField
	}
	runJobTaskField, err := RunJobTaskFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	sparkJarTaskField, err := SparkJarTaskFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := SparkPythonTaskFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := SparkSubmitTaskFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := SqlTaskFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}
	st.TaskKey = pb.TaskKey
	st.TimeoutSeconds = pb.TimeoutSeconds
	webhookNotificationsField, err := WebhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	// Wire name: 'outcome'
	Outcome string ``
	// The name of the task this task depends on.
	// Wire name: 'task_key'
	TaskKey         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st TaskDependency) MarshalJSON() ([]byte, error) {
	pb, err := TaskDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TaskDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TaskDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TaskDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TaskDependencyToPb(st *TaskDependency) (*jobspb.TaskDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TaskDependencyPb{}
	pb.Outcome = st.Outcome
	pb.TaskKey = st.TaskKey

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TaskDependencyFromPb(pb *jobspb.TaskDependencyPb) (*TaskDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskDependency{}
	st.Outcome = pb.Outcome
	st.TaskKey = pb.TaskKey

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TaskEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool ``
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	// Wire name: 'on_duration_warning_threshold_exceeded'
	OnDurationWarningThresholdExceeded []string ``
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	// Wire name: 'on_failure'
	OnFailure []string ``
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_start'
	OnStart []string ``
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	// Wire name: 'on_streaming_backlog_exceeded'
	OnStreamingBacklogExceeded []string ``
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_success'
	OnSuccess       []string ``
	ForceSendFields []string `tf:"-"`
}

func (st TaskEmailNotifications) MarshalJSON() ([]byte, error) {
	pb, err := TaskEmailNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TaskEmailNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TaskEmailNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TaskEmailNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TaskEmailNotificationsToPb(st *TaskEmailNotifications) (*jobspb.TaskEmailNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TaskEmailNotificationsPb{}
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns
	pb.OnDurationWarningThresholdExceeded = st.OnDurationWarningThresholdExceeded
	pb.OnFailure = st.OnFailure
	pb.OnStart = st.OnStart
	pb.OnStreamingBacklogExceeded = st.OnStreamingBacklogExceeded
	pb.OnSuccess = st.OnSuccess

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TaskEmailNotificationsFromPb(pb *jobspb.TaskEmailNotificationsPb) (*TaskEmailNotifications, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	// Wire name: 'alert_on_last_attempt'
	AlertOnLastAttempt bool ``
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	// Wire name: 'no_alert_for_canceled_runs'
	NoAlertForCanceledRuns bool ``
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (st TaskNotificationSettings) MarshalJSON() ([]byte, error) {
	pb, err := TaskNotificationSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TaskNotificationSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TaskNotificationSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TaskNotificationSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TaskNotificationSettingsToPb(st *TaskNotificationSettings) (*jobspb.TaskNotificationSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TaskNotificationSettingsPb{}
	pb.AlertOnLastAttempt = st.AlertOnLastAttempt
	pb.NoAlertForCanceledRuns = st.NoAlertForCanceledRuns
	pb.NoAlertForSkippedRuns = st.NoAlertForSkippedRuns

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TaskNotificationSettingsFromPb(pb *jobspb.TaskNotificationSettingsPb) (*TaskNotificationSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskNotificationSettings{}
	st.AlertOnLastAttempt = pb.AlertOnLastAttempt
	st.NoAlertForCanceledRuns = pb.NoAlertForCanceledRuns
	st.NoAlertForSkippedRuns = pb.NoAlertForSkippedRuns

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The code indicates why the run was terminated. Additional codes might be
// introduced in future releases. * `SUCCESS`: The run was completed
// successfully. * `SUCCESS_WITH_FAILURES`: The run was completed successfully
// but some child runs failed. * `USER_CANCELED`: The run was successfully
// canceled during execution by a user. * `CANCELED`: The run was canceled
// during execution by the Databricks platform; for example, if the maximum run
// duration was exceeded. * `SKIPPED`: Run was never executed, for example, if
// the upstream task run failed, the dependency type condition was not met, or
// there were no material tasks to execute. * `INTERNAL_ERROR`: The run
// encountered an unexpected error. Refer to the state message for further
// details. * `DRIVER_ERROR`: The run encountered an error while communicating
// with the Spark Driver. * `CLUSTER_ERROR`: The run failed due to a cluster
// error. Refer to the state message for further details. *
// `REPOSITORY_CHECKOUT_FAILED`: Failed to complete the checkout due to an error
// when communicating with the third party service. * `INVALID_CLUSTER_REQUEST`:
// The run failed because it issued an invalid request to start the cluster. *
// `WORKSPACE_RUN_LIMIT_EXCEEDED`: The workspace has reached the quota for the
// maximum number of concurrent active runs. Consider scheduling the runs over a
// larger time frame. * `FEATURE_DISABLED`: The run failed because it tried to
// access a feature unavailable for the workspace. *
// `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The number of cluster creation, start, and
// upsize requests have exceeded the allotted rate limit. Consider spreading the
// run execution over a larger time frame. * `STORAGE_ACCESS_ERROR`: The run
// failed due to an error when accessing the customer blob storage. Refer to the
// state message for further details. * `RUN_EXECUTION_ERROR`: The run was
// completed with task failures. For more details, refer to the state message or
// run output. * `UNAUTHORIZED_ERROR`: The run failed due to a permission issue
// while accessing a resource. Refer to the state message for further details. *
// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
// user-requested library. Refer to the state message for further details. The
// causes might include, but are not limited to: The provided library is
// invalid, there are insufficient permissions to install the library, and so
// forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the limit
// of maximum concurrent runs set for the job. * `MAX_SPARK_CONTEXTS_EXCEEDED`:
// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link]. * `RESOURCE_NOT_FOUND`:
// A resource necessary for run execution does not exist. Refer to the state
// message for further details. * `INVALID_RUN_CONFIGURATION`: The run failed
// due to an invalid configuration. Refer to the state message for further
// details. * `CLOUD_FAILURE`: The run failed due to a cloud provider issue.
// Refer to the state message for further details. *
// `MAX_JOB_QUEUE_SIZE_EXCEEDED`: The run was skipped due to reaching the job
// level queue size limit. * `DISABLED`: The run was never executed because it
// was disabled explicitly by the user.
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
type TerminationCodeCode string

const TerminationCodeCodeBudgetPolicyLimitExceeded TerminationCodeCode = `BUDGET_POLICY_LIMIT_EXCEEDED`

// The run was canceled during execution by the <Databricks> platform; for
// example, if the maximum run duration was exceeded.
const TerminationCodeCodeCanceled TerminationCodeCode = `CANCELED`

// The run failed due to a cloud provider issue. Refer to the state message for
// further details.
const TerminationCodeCodeCloudFailure TerminationCodeCode = `CLOUD_FAILURE`

// The run failed due to a cluster error. Refer to the state message for further
// details.
const TerminationCodeCodeClusterError TerminationCodeCode = `CLUSTER_ERROR`

// The number of cluster creation, start, and upsize requests have exceeded the
// allotted rate limit. Consider spreading the run execution over a larger time
// frame.
const TerminationCodeCodeClusterRequestLimitExceeded TerminationCodeCode = `CLUSTER_REQUEST_LIMIT_EXCEEDED`

// The run was never executed because it was disabled explicitly by the user.
const TerminationCodeCodeDisabled TerminationCodeCode = `DISABLED`

// The run encountered an error while communicating with the Spark Driver.
const TerminationCodeCodeDriverError TerminationCodeCode = `DRIVER_ERROR`

// The run failed because it tried to access a feature unavailable for the
// workspace.
const TerminationCodeCodeFeatureDisabled TerminationCodeCode = `FEATURE_DISABLED`

// The run encountered an unexpected error. Refer to the state message for
// further details.
const TerminationCodeCodeInternalError TerminationCodeCode = `INTERNAL_ERROR`

// The run failed because it issued an invalid request to start the cluster.
const TerminationCodeCodeInvalidClusterRequest TerminationCodeCode = `INVALID_CLUSTER_REQUEST`

// The run failed due to an invalid configuration. Refer to the state message
// for further details.
const TerminationCodeCodeInvalidRunConfiguration TerminationCodeCode = `INVALID_RUN_CONFIGURATION`

// The run failed while installing the user-requested library. Refer to the
// state message for further details. The causes might include, but are not
// limited to: The provided library is invalid, there are insufficient
// permissions to install the library, and so forth.
const TerminationCodeCodeLibraryInstallationError TerminationCodeCode = `LIBRARY_INSTALLATION_ERROR`

// The scheduled run exceeds the limit of maximum concurrent runs set for the
// job.
const TerminationCodeCodeMaxConcurrentRunsExceeded TerminationCodeCode = `MAX_CONCURRENT_RUNS_EXCEEDED`

// The run was skipped due to reaching the job level queue size limit.
const TerminationCodeCodeMaxJobQueueSizeExceeded TerminationCodeCode = `MAX_JOB_QUEUE_SIZE_EXCEEDED`

// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link].
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
const TerminationCodeCodeMaxSparkContextsExceeded TerminationCodeCode = `MAX_SPARK_CONTEXTS_EXCEEDED`

// Failed to complete the checkout due to an error when communicating with the
// third party service.
const TerminationCodeCodeRepositoryCheckoutFailed TerminationCodeCode = `REPOSITORY_CHECKOUT_FAILED`

// A resource necessary for run execution does not exist. Refer to the state
// message for further details.
const TerminationCodeCodeResourceNotFound TerminationCodeCode = `RESOURCE_NOT_FOUND`

// The run was completed with task failures. For more details, refer to the
// state message or run output.
const TerminationCodeCodeRunExecutionError TerminationCodeCode = `RUN_EXECUTION_ERROR`

// Run was never executed, for example, if the upstream task run failed, the
// dependency type condition was not met, or there were no material tasks to
// execute.
const TerminationCodeCodeSkipped TerminationCodeCode = `SKIPPED`

// The run failed due to an error when accessing the customer blob storage.
// Refer to the state message for further details.
const TerminationCodeCodeStorageAccessError TerminationCodeCode = `STORAGE_ACCESS_ERROR`

// The run was completed successfully.
const TerminationCodeCodeSuccess TerminationCodeCode = `SUCCESS`

// The run was completed successfully but some child runs failed.
const TerminationCodeCodeSuccessWithFailures TerminationCodeCode = `SUCCESS_WITH_FAILURES`

// The run failed due to a permission issue while accessing a resource. Refer to
// the state message for further details.
const TerminationCodeCodeUnauthorizedError TerminationCodeCode = `UNAUTHORIZED_ERROR`

// The run was successfully canceled during execution by a user.
const TerminationCodeCodeUserCanceled TerminationCodeCode = `USER_CANCELED`

// The workspace has reached the quota for the maximum number of concurrent
// active runs. Consider scheduling the runs over a larger time frame.
const TerminationCodeCodeWorkspaceRunLimitExceeded TerminationCodeCode = `WORKSPACE_RUN_LIMIT_EXCEEDED`

// String representation for [fmt.Print]
func (f *TerminationCodeCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationCodeCode) Set(v string) error {
	switch v {
	case `BUDGET_POLICY_LIMIT_EXCEEDED`, `CANCELED`, `CLOUD_FAILURE`, `CLUSTER_ERROR`, `CLUSTER_REQUEST_LIMIT_EXCEEDED`, `DISABLED`, `DRIVER_ERROR`, `FEATURE_DISABLED`, `INTERNAL_ERROR`, `INVALID_CLUSTER_REQUEST`, `INVALID_RUN_CONFIGURATION`, `LIBRARY_INSTALLATION_ERROR`, `MAX_CONCURRENT_RUNS_EXCEEDED`, `MAX_JOB_QUEUE_SIZE_EXCEEDED`, `MAX_SPARK_CONTEXTS_EXCEEDED`, `REPOSITORY_CHECKOUT_FAILED`, `RESOURCE_NOT_FOUND`, `RUN_EXECUTION_ERROR`, `SKIPPED`, `STORAGE_ACCESS_ERROR`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `UNAUTHORIZED_ERROR`, `USER_CANCELED`, `WORKSPACE_RUN_LIMIT_EXCEEDED`:
		*f = TerminationCodeCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUDGET_POLICY_LIMIT_EXCEEDED", "CANCELED", "CLOUD_FAILURE", "CLUSTER_ERROR", "CLUSTER_REQUEST_LIMIT_EXCEEDED", "DISABLED", "DRIVER_ERROR", "FEATURE_DISABLED", "INTERNAL_ERROR", "INVALID_CLUSTER_REQUEST", "INVALID_RUN_CONFIGURATION", "LIBRARY_INSTALLATION_ERROR", "MAX_CONCURRENT_RUNS_EXCEEDED", "MAX_JOB_QUEUE_SIZE_EXCEEDED", "MAX_SPARK_CONTEXTS_EXCEEDED", "REPOSITORY_CHECKOUT_FAILED", "RESOURCE_NOT_FOUND", "RUN_EXECUTION_ERROR", "SKIPPED", "STORAGE_ACCESS_ERROR", "SUCCESS", "SUCCESS_WITH_FAILURES", "UNAUTHORIZED_ERROR", "USER_CANCELED", "WORKSPACE_RUN_LIMIT_EXCEEDED"`, v)
	}
}

// Values returns all possible values for TerminationCodeCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationCodeCode) Values() []TerminationCodeCode {
	return []TerminationCodeCode{
		TerminationCodeCodeBudgetPolicyLimitExceeded,
		TerminationCodeCodeCanceled,
		TerminationCodeCodeCloudFailure,
		TerminationCodeCodeClusterError,
		TerminationCodeCodeClusterRequestLimitExceeded,
		TerminationCodeCodeDisabled,
		TerminationCodeCodeDriverError,
		TerminationCodeCodeFeatureDisabled,
		TerminationCodeCodeInternalError,
		TerminationCodeCodeInvalidClusterRequest,
		TerminationCodeCodeInvalidRunConfiguration,
		TerminationCodeCodeLibraryInstallationError,
		TerminationCodeCodeMaxConcurrentRunsExceeded,
		TerminationCodeCodeMaxJobQueueSizeExceeded,
		TerminationCodeCodeMaxSparkContextsExceeded,
		TerminationCodeCodeRepositoryCheckoutFailed,
		TerminationCodeCodeResourceNotFound,
		TerminationCodeCodeRunExecutionError,
		TerminationCodeCodeSkipped,
		TerminationCodeCodeStorageAccessError,
		TerminationCodeCodeSuccess,
		TerminationCodeCodeSuccessWithFailures,
		TerminationCodeCodeUnauthorizedError,
		TerminationCodeCodeUserCanceled,
		TerminationCodeCodeWorkspaceRunLimitExceeded,
	}
}

// Type always returns TerminationCodeCode to satisfy [pflag.Value] interface
func (f *TerminationCodeCode) Type() string {
	return "TerminationCodeCode"
}

func TerminationCodeCodeToPb(st *TerminationCodeCode) (*jobspb.TerminationCodeCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.TerminationCodeCodePb(*st)
	return &pb, nil
}

func TerminationCodeCodeFromPb(pb *jobspb.TerminationCodeCodePb) (*TerminationCodeCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationCodeCode(*pb)
	return &st, nil
}

type TerminationDetails struct {

	// Wire name: 'code'
	Code TerminationCodeCode ``
	// A descriptive message with the termination details. This field is
	// unstructured and the format might change.
	// Wire name: 'message'
	Message string ``

	// Wire name: 'type'
	Type            TerminationTypeType ``
	ForceSendFields []string            `tf:"-"`
}

func (st TerminationDetails) MarshalJSON() ([]byte, error) {
	pb, err := TerminationDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TerminationDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TerminationDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TerminationDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TerminationDetailsToPb(st *TerminationDetails) (*jobspb.TerminationDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TerminationDetailsPb{}
	codePb, err := TerminationCodeCodeToPb(&st.Code)
	if err != nil {
		return nil, err
	}
	if codePb != nil {
		pb.Code = *codePb
	}
	pb.Message = st.Message
	typePb, err := TerminationTypeTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TerminationDetailsFromPb(pb *jobspb.TerminationDetailsPb) (*TerminationDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TerminationDetails{}
	codeField, err := TerminationCodeCodeFromPb(&pb.Code)
	if err != nil {
		return nil, err
	}
	if codeField != nil {
		st.Code = *codeField
	}
	st.Message = pb.Message
	typeField, err := TerminationTypeTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
// error occurred in the Databricks platform. Please look at the [status page]
// or contact support if the issue persists. * `CLIENT_ERROR`: The run was
// terminated because of an error caused by user input or the job configuration.
// * `CLOUD_FAILURE`: The run was terminated because of an issue with your cloud
// provider.
//
// [status page]: https://status.databricks.com/
type TerminationTypeType string

// The run was terminated because of an error caused by user input or the job
// configuration.
const TerminationTypeTypeClientError TerminationTypeType = `CLIENT_ERROR`

// The run was terminated because of an issue with your cloud provider.
const TerminationTypeTypeCloudFailure TerminationTypeType = `CLOUD_FAILURE`

// An error occurred in the <Databricks> platform. Please look at the [status
// page] or contact support if the issue persists.
//
// [status page]: https://status.databricks.com/
const TerminationTypeTypeInternalError TerminationTypeType = `INTERNAL_ERROR`

// The run terminated without any issues
const TerminationTypeTypeSuccess TerminationTypeType = `SUCCESS`

// String representation for [fmt.Print]
func (f *TerminationTypeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationTypeType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `INTERNAL_ERROR`, `SUCCESS`:
		*f = TerminationTypeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "INTERNAL_ERROR", "SUCCESS"`, v)
	}
}

// Values returns all possible values for TerminationTypeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationTypeType) Values() []TerminationTypeType {
	return []TerminationTypeType{
		TerminationTypeTypeClientError,
		TerminationTypeTypeCloudFailure,
		TerminationTypeTypeInternalError,
		TerminationTypeTypeSuccess,
	}
}

// Type always returns TerminationTypeType to satisfy [pflag.Value] interface
func (f *TerminationTypeType) Type() string {
	return "TerminationTypeType"
}

func TerminationTypeTypeToPb(st *TerminationTypeType) (*jobspb.TerminationTypeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.TerminationTypeTypePb(*st)
	return &pb, nil
}

func TerminationTypeTypeFromPb(pb *jobspb.TerminationTypeTypePb) (*TerminationTypeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationTypeType(*pb)
	return &st, nil
}

// Additional details about what triggered the run
type TriggerInfo struct {
	// The run id of the Run Job task run
	// Wire name: 'run_id'
	RunId           int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st TriggerInfo) MarshalJSON() ([]byte, error) {
	pb, err := TriggerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TriggerInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TriggerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TriggerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TriggerInfoToPb(st *TriggerInfo) (*jobspb.TriggerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TriggerInfoPb{}
	pb.RunId = st.RunId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TriggerInfoFromPb(pb *jobspb.TriggerInfoPb) (*TriggerInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerInfo{}
	st.RunId = pb.RunId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TriggerSettings struct {
	// File arrival trigger settings.
	// Wire name: 'file_arrival'
	FileArrival *FileArrivalTriggerConfiguration ``
	// Whether this trigger is paused or not.
	// Wire name: 'pause_status'
	PauseStatus PauseStatus ``
	// Periodic trigger settings.
	// Wire name: 'periodic'
	Periodic *PeriodicTriggerConfiguration ``
	// Old table trigger settings name. Deprecated in favor of `table_update`.
	// Wire name: 'table'
	Table *TableUpdateTriggerConfiguration ``

	// Wire name: 'table_update'
	TableUpdate *TableUpdateTriggerConfiguration ``
}

func (st TriggerSettings) MarshalJSON() ([]byte, error) {
	pb, err := TriggerSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TriggerSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TriggerSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TriggerSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TriggerSettingsToPb(st *TriggerSettings) (*jobspb.TriggerSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TriggerSettingsPb{}
	fileArrivalPb, err := FileArrivalTriggerConfigurationToPb(st.FileArrival)
	if err != nil {
		return nil, err
	}
	if fileArrivalPb != nil {
		pb.FileArrival = fileArrivalPb
	}
	pauseStatusPb, err := PauseStatusToPb(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}
	periodicPb, err := PeriodicTriggerConfigurationToPb(st.Periodic)
	if err != nil {
		return nil, err
	}
	if periodicPb != nil {
		pb.Periodic = periodicPb
	}
	tablePb, err := TableUpdateTriggerConfigurationToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}
	tableUpdatePb, err := TableUpdateTriggerConfigurationToPb(st.TableUpdate)
	if err != nil {
		return nil, err
	}
	if tableUpdatePb != nil {
		pb.TableUpdate = tableUpdatePb
	}

	return pb, nil
}

func TriggerSettingsFromPb(pb *jobspb.TriggerSettingsPb) (*TriggerSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerSettings{}
	fileArrivalField, err := FileArrivalTriggerConfigurationFromPb(pb.FileArrival)
	if err != nil {
		return nil, err
	}
	if fileArrivalField != nil {
		st.FileArrival = fileArrivalField
	}
	pauseStatusField, err := PauseStatusFromPb(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	periodicField, err := PeriodicTriggerConfigurationFromPb(pb.Periodic)
	if err != nil {
		return nil, err
	}
	if periodicField != nil {
		st.Periodic = periodicField
	}
	tableField, err := TableUpdateTriggerConfigurationFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}
	tableUpdateField, err := TableUpdateTriggerConfigurationFromPb(pb.TableUpdate)
	if err != nil {
		return nil, err
	}
	if tableUpdateField != nil {
		st.TableUpdate = tableUpdateField
	}

	return st, nil
}

type TriggerStateProto struct {

	// Wire name: 'file_arrival'
	FileArrival *FileArrivalTriggerState ``
}

func (st TriggerStateProto) MarshalJSON() ([]byte, error) {
	pb, err := TriggerStateProtoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TriggerStateProto) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.TriggerStateProtoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TriggerStateProtoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TriggerStateProtoToPb(st *TriggerStateProto) (*jobspb.TriggerStateProtoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.TriggerStateProtoPb{}
	fileArrivalPb, err := FileArrivalTriggerStateToPb(st.FileArrival)
	if err != nil {
		return nil, err
	}
	if fileArrivalPb != nil {
		pb.FileArrival = fileArrivalPb
	}

	return pb, nil
}

func TriggerStateProtoFromPb(pb *jobspb.TriggerStateProtoPb) (*TriggerStateProto, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerStateProto{}
	fileArrivalField, err := FileArrivalTriggerStateFromPb(pb.FileArrival)
	if err != nil {
		return nil, err
	}
	if fileArrivalField != nil {
		st.FileArrival = fileArrivalField
	}

	return st, nil
}

// The type of trigger that fired this run.
//
// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures. *
// `RUN_JOB_TASK`: Indicates a run that is triggered using a Run Job task. *
// `FILE_ARRIVAL`: Indicates a run that is triggered by a file arrival. *
// `CONTINUOUS`: Indicates a run that is triggered by a continuous job. *
// `TABLE`: Indicates a run that is triggered by a table update. *
// `CONTINUOUS_RESTART`: Indicates a run created by user to manually restart a
// continuous job run. * `MODEL`: Indicates a run that is triggered by a model
// update.
type TriggerType string

// Indicates a run that is triggered by a continuous job.
const TriggerTypeContinuous TriggerType = `CONTINUOUS`

// Indicates a run created by user to manually restart a continuous job run.
const TriggerTypeContinuousRestart TriggerType = `CONTINUOUS_RESTART`

// Indicates a run that is triggered by a file arrival.
const TriggerTypeFileArrival TriggerType = `FILE_ARRIVAL`

// One time triggers that fire a single run. This occurs you triggered a single
// run on demand through the UI or the API.
const TriggerTypeOneTime TriggerType = `ONE_TIME`

// Schedules that periodically trigger runs, such as a cron scheduler.
const TriggerTypePeriodic TriggerType = `PERIODIC`

// Indicates a run that is triggered as a retry of a previously failed run. This
// occurs when you request to re-run the job in case of failures.
const TriggerTypeRetry TriggerType = `RETRY`

// Indicates a run that is triggered using a Run Job task.
const TriggerTypeRunJobTask TriggerType = `RUN_JOB_TASK`

// Indicates a run that is triggered by a table update.
const TriggerTypeTable TriggerType = `TABLE`

// String representation for [fmt.Print]
func (f *TriggerType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TriggerType) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `CONTINUOUS_RESTART`, `FILE_ARRIVAL`, `ONE_TIME`, `PERIODIC`, `RETRY`, `RUN_JOB_TASK`, `TABLE`:
		*f = TriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "CONTINUOUS_RESTART", "FILE_ARRIVAL", "ONE_TIME", "PERIODIC", "RETRY", "RUN_JOB_TASK", "TABLE"`, v)
	}
}

// Values returns all possible values for TriggerType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TriggerType) Values() []TriggerType {
	return []TriggerType{
		TriggerTypeContinuous,
		TriggerTypeContinuousRestart,
		TriggerTypeFileArrival,
		TriggerTypeOneTime,
		TriggerTypePeriodic,
		TriggerTypeRetry,
		TriggerTypeRunJobTask,
		TriggerTypeTable,
	}
}

// Type always returns TriggerType to satisfy [pflag.Value] interface
func (f *TriggerType) Type() string {
	return "TriggerType"
}

func TriggerTypeToPb(st *TriggerType) (*jobspb.TriggerTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.TriggerTypePb(*st)
	return &pb, nil
}

func TriggerTypeFromPb(pb *jobspb.TriggerTypePb) (*TriggerType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TriggerType(*pb)
	return &st, nil
}

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported, except for tasks and job clusters (`tasks/task_1`). This
	// field is optional.
	// Wire name: 'fields_to_remove'
	FieldsToRemove []string ``
	// The canonical identifier of the job to update. This field is required.
	// Wire name: 'job_id'
	JobId int64 ``
	// The new settings for the job.
	//
	// Top-level fields specified in `new_settings` are completely replaced,
	// except for arrays which are merged. That is, new and existing entries are
	// completely replaced based on the respective key fields, i.e. `task_key`
	// or `job_cluster_key`, while previous entries are kept.
	//
	// Partially updating nested fields is not supported.
	//
	// Changes to the field `JobSettings.timeout_seconds` are applied to active
	// runs. Changes to other fields are applied to future runs only.
	// Wire name: 'new_settings'
	NewSettings *JobSettings ``
}

func (st UpdateJob) MarshalJSON() ([]byte, error) {
	pb, err := UpdateJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.UpdateJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateJobToPb(st *UpdateJob) (*jobspb.UpdateJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.UpdateJobPb{}
	pb.FieldsToRemove = st.FieldsToRemove
	pb.JobId = st.JobId
	newSettingsPb, err := JobSettingsToPb(st.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsPb != nil {
		pb.NewSettings = newSettingsPb
	}

	return pb, nil
}

func UpdateJobFromPb(pb *jobspb.UpdateJobPb) (*UpdateJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateJob{}
	st.FieldsToRemove = pb.FieldsToRemove
	st.JobId = pb.JobId
	newSettingsField, err := JobSettingsFromPb(pb.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsField != nil {
		st.NewSettings = newSettingsField
	}

	return st, nil
}

type ViewItem struct {
	// Content of the view.
	// Wire name: 'content'
	Content string ``
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	// Wire name: 'name'
	Name string ``
	// Type of the view item.
	// Wire name: 'type'
	Type            ViewType ``
	ForceSendFields []string `tf:"-"`
}

func (st ViewItem) MarshalJSON() ([]byte, error) {
	pb, err := ViewItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ViewItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.ViewItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ViewItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ViewItemToPb(st *ViewItem) (*jobspb.ViewItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.ViewItemPb{}
	pb.Content = st.Content
	pb.Name = st.Name
	typePb, err := ViewTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ViewItemFromPb(pb *jobspb.ViewItemPb) (*ViewItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ViewItem{}
	st.Content = pb.Content
	st.Name = pb.Name
	typeField, err := ViewTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// * `NOTEBOOK`: Notebook view item. * `DASHBOARD`: Dashboard view item.
type ViewType string

// Dashboard view item.
const ViewTypeDashboard ViewType = `DASHBOARD`

// Notebook view item.
const ViewTypeNotebook ViewType = `NOTEBOOK`

// String representation for [fmt.Print]
func (f *ViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewType) Set(v string) error {
	switch v {
	case `DASHBOARD`, `NOTEBOOK`:
		*f = ViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD", "NOTEBOOK"`, v)
	}
}

// Values returns all possible values for ViewType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewType) Values() []ViewType {
	return []ViewType{
		ViewTypeDashboard,
		ViewTypeNotebook,
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}

func ViewTypeToPb(st *ViewType) (*jobspb.ViewTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.ViewTypePb(*st)
	return &pb, nil
}

func ViewTypeFromPb(pb *jobspb.ViewTypePb) (*ViewType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ViewType(*pb)
	return &st, nil
}

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
type ViewsToExport string

// All views of the notebook.
const ViewsToExportAll ViewsToExport = `ALL`

// Code view of the notebook.
const ViewsToExportCode ViewsToExport = `CODE`

// All dashboard views of the notebook.
const ViewsToExportDashboards ViewsToExport = `DASHBOARDS`

// String representation for [fmt.Print]
func (f *ViewsToExport) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewsToExport) Set(v string) error {
	switch v {
	case `ALL`, `CODE`, `DASHBOARDS`:
		*f = ViewsToExport(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL", "CODE", "DASHBOARDS"`, v)
	}
}

// Values returns all possible values for ViewsToExport.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewsToExport) Values() []ViewsToExport {
	return []ViewsToExport{
		ViewsToExportAll,
		ViewsToExportCode,
		ViewsToExportDashboards,
	}
}

// Type always returns ViewsToExport to satisfy [pflag.Value] interface
func (f *ViewsToExport) Type() string {
	return "ViewsToExport"
}

func ViewsToExportToPb(st *ViewsToExport) (*jobspb.ViewsToExportPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobspb.ViewsToExportPb(*st)
	return &pb, nil
}

func ViewsToExportFromPb(pb *jobspb.ViewsToExportPb) (*ViewsToExport, error) {
	if pb == nil {
		return nil, nil
	}
	st := ViewsToExport(*pb)
	return &st, nil
}

type Webhook struct {

	// Wire name: 'id'
	Id string ``
}

func (st Webhook) MarshalJSON() ([]byte, error) {
	pb, err := WebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Webhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.WebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WebhookToPb(st *Webhook) (*jobspb.WebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.WebhookPb{}
	pb.Id = st.Id

	return pb, nil
}

func WebhookFromPb(pb *jobspb.WebhookPb) (*Webhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Webhook{}
	st.Id = pb.Id

	return st, nil
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	// Wire name: 'on_duration_warning_threshold_exceeded'
	OnDurationWarningThresholdExceeded []Webhook ``
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	// Wire name: 'on_failure'
	OnFailure []Webhook ``
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	// Wire name: 'on_start'
	OnStart []Webhook ``
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	// Wire name: 'on_streaming_backlog_exceeded'
	OnStreamingBacklogExceeded []Webhook ``
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	// Wire name: 'on_success'
	OnSuccess []Webhook ``
}

func (st WebhookNotifications) MarshalJSON() ([]byte, error) {
	pb, err := WebhookNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WebhookNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.WebhookNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WebhookNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WebhookNotificationsToPb(st *WebhookNotifications) (*jobspb.WebhookNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.WebhookNotificationsPb{}

	var onDurationWarningThresholdExceededPb []jobspb.WebhookPb
	for _, item := range st.OnDurationWarningThresholdExceeded {
		itemPb, err := WebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onDurationWarningThresholdExceededPb = append(onDurationWarningThresholdExceededPb, *itemPb)
		}
	}
	pb.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededPb

	var onFailurePb []jobspb.WebhookPb
	for _, item := range st.OnFailure {
		itemPb, err := WebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onFailurePb = append(onFailurePb, *itemPb)
		}
	}
	pb.OnFailure = onFailurePb

	var onStartPb []jobspb.WebhookPb
	for _, item := range st.OnStart {
		itemPb, err := WebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStartPb = append(onStartPb, *itemPb)
		}
	}
	pb.OnStart = onStartPb

	var onStreamingBacklogExceededPb []jobspb.WebhookPb
	for _, item := range st.OnStreamingBacklogExceeded {
		itemPb, err := WebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStreamingBacklogExceededPb = append(onStreamingBacklogExceededPb, *itemPb)
		}
	}
	pb.OnStreamingBacklogExceeded = onStreamingBacklogExceededPb

	var onSuccessPb []jobspb.WebhookPb
	for _, item := range st.OnSuccess {
		itemPb, err := WebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onSuccessPb = append(onSuccessPb, *itemPb)
		}
	}
	pb.OnSuccess = onSuccessPb

	return pb, nil
}

func WebhookNotificationsFromPb(pb *jobspb.WebhookNotificationsPb) (*WebhookNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WebhookNotifications{}

	var onDurationWarningThresholdExceededField []Webhook
	for _, itemPb := range pb.OnDurationWarningThresholdExceeded {
		item, err := WebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onDurationWarningThresholdExceededField = append(onDurationWarningThresholdExceededField, *item)
		}
	}
	st.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededField

	var onFailureField []Webhook
	for _, itemPb := range pb.OnFailure {
		item, err := WebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onFailureField = append(onFailureField, *item)
		}
	}
	st.OnFailure = onFailureField

	var onStartField []Webhook
	for _, itemPb := range pb.OnStart {
		item, err := WebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStartField = append(onStartField, *item)
		}
	}
	st.OnStart = onStartField

	var onStreamingBacklogExceededField []Webhook
	for _, itemPb := range pb.OnStreamingBacklogExceeded {
		item, err := WebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStreamingBacklogExceededField = append(onStreamingBacklogExceededField, *item)
		}
	}
	st.OnStreamingBacklogExceeded = onStreamingBacklogExceededField

	var onSuccessField []Webhook
	for _, itemPb := range pb.OnSuccess {
		item, err := WebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onSuccessField = append(onSuccessField, *item)
		}
	}
	st.OnSuccess = onSuccessField

	return st, nil
}

type WidgetErrorDetail struct {

	// Wire name: 'message'
	Message         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st WidgetErrorDetail) MarshalJSON() ([]byte, error) {
	pb, err := WidgetErrorDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WidgetErrorDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobspb.WidgetErrorDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WidgetErrorDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WidgetErrorDetailToPb(st *WidgetErrorDetail) (*jobspb.WidgetErrorDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobspb.WidgetErrorDetailPb{}
	pb.Message = st.Message

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func WidgetErrorDetailFromPb(pb *jobspb.WidgetErrorDetailPb) (*WidgetErrorDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WidgetErrorDetail{}
	st.Message = pb.Message

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
