// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"encoding/json"
	"fmt"

	"github.com/databricks/databricks-sdk-go/service/compute"
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

// Type always returns AuthenticationMethod to satisfy [pflag.Value] interface
func (f *AuthenticationMethod) Type() string {
	return "AuthenticationMethod"
}

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	// Wire name: 'created_time'
	CreatedTime int64
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId int64
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	// Wire name: 'settings'
	Settings *JobSettings

	ForceSendFields []string `tf:"-"`
}

func (st *BaseJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &baseJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := baseJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BaseJob) MarshalJSON() ([]byte, error) {
	pb, err := baseJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	AttemptNumber int
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	// Wire name: 'cleanup_duration'
	CleanupDuration int64
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	// Wire name: 'cluster_instance'
	ClusterInstance *ClusterInstance
	// A snapshot of the job’s cluster specification when this run was
	// created.
	// Wire name: 'cluster_spec'
	ClusterSpec *ClusterSpec
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// Description of the run
	// Wire name: 'description'
	Description string
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	// Wire name: 'end_time'
	EndTime int64
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	// Wire name: 'execution_duration'
	ExecutionDuration int64
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
	GitSource *GitSource
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster
	// The canonical identifier of the job that contains this run.
	// Wire name: 'job_id'
	JobId int64
	// Job-level parameters used in the run
	// Wire name: 'job_parameters'
	JobParameters []JobParameter
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	// Wire name: 'job_run_id'
	JobRunId int64
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	// Wire name: 'number_in_job'
	NumberInJob int64
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	// Wire name: 'original_attempt_run_id'
	OriginalAttemptRunId int64
	// The parameters used for this run.
	// Wire name: 'overriding_parameters'
	OverridingParameters *RunParameters
	// The time in milliseconds that the run has spent in the queue.
	// Wire name: 'queue_duration'
	QueueDuration int64
	// The repair history of the run.
	// Wire name: 'repair_history'
	RepairHistory []RepairHistoryItem
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	// Wire name: 'run_duration'
	RunDuration int64
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'run_id'
	RunId int64
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'run_name'
	RunName string
	// The URL to the detail page of the run.
	// Wire name: 'run_page_url'
	RunPageUrl string
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	// Wire name: 'run_type'
	RunType RunType
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	// Wire name: 'schedule'
	Schedule *CronSchedule
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	// Wire name: 'setup_duration'
	SetupDuration int64
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	// Wire name: 'start_time'
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState
	// The current status of the run
	// Wire name: 'status'
	Status *RunStatus
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	// Wire name: 'tasks'
	Tasks []RunTask
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	// Wire name: 'trigger'
	Trigger TriggerType
	// Additional details about what triggered the run
	// Wire name: 'trigger_info'
	TriggerInfo *TriggerInfo

	ForceSendFields []string `tf:"-"`
}

func (st *BaseRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &baseRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := baseRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BaseRun) MarshalJSON() ([]byte, error) {
	pb, err := baseRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	// Wire name: 'all_queued_runs'
	AllQueuedRuns bool
	// The canonical identifier of the job to cancel all runs of.
	// Wire name: 'job_id'
	JobId int64

	ForceSendFields []string `tf:"-"`
}

func (st *CancelAllRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelAllRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelAllRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelAllRuns) MarshalJSON() ([]byte, error) {
	pb, err := cancelAllRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelAllRunsResponse struct {
}

func (st *CancelAllRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelAllRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelAllRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelAllRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelAllRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelRun struct {
	// This field is required.
	// Wire name: 'run_id'
	RunId int64
}

func (st *CancelRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelRun) MarshalJSON() ([]byte, error) {
	pb, err := cancelRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelRunResponse struct {
}

func (st *CancelRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns CleanRoomTaskRunLifeCycleState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunLifeCycleState) Type() string {
	return "CleanRoomTaskRunLifeCycleState"
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

// Type always returns CleanRoomTaskRunResultState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunResultState) Type() string {
	return "CleanRoomTaskRunResultState"
}

// Stores the run state of the clean rooms notebook task.
type CleanRoomTaskRunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	// Wire name: 'life_cycle_state'
	LifeCycleState CleanRoomTaskRunLifeCycleState
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	// Wire name: 'result_state'
	ResultState CleanRoomTaskRunResultState
}

func (st *CleanRoomTaskRunState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomTaskRunStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomTaskRunStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomTaskRunState) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomTaskRunStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CleanRoomsNotebookTask struct {
	// The clean room that the notebook belongs to.
	// Wire name: 'clean_room_name'
	CleanRoomName string
	// Checksum to validate the freshness of the notebook resource (i.e. the
	// notebook being run is the latest version). It can be fetched by calling
	// the :method:cleanroomassets/get API.
	// Wire name: 'etag'
	Etag string
	// Base parameters to be used for the clean room notebook job.
	// Wire name: 'notebook_base_parameters'
	NotebookBaseParameters map[string]string
	// Name of the notebook being run.
	// Wire name: 'notebook_name'
	NotebookName string

	ForceSendFields []string `tf:"-"`
}

func (st *CleanRoomsNotebookTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomsNotebookTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomsNotebookTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomsNotebookTask) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomsNotebookTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput struct {
	// The run state of the clean rooms notebook task.
	// Wire name: 'clean_room_job_run_state'
	CleanRoomJobRunState *CleanRoomTaskRunState
	// The notebook output for the clean room run
	// Wire name: 'notebook_output'
	NotebookOutput *NotebookOutput
	// Information on how to access the output schema for the clean room run
	// Wire name: 'output_schema_info'
	OutputSchemaInfo *OutputSchemaInfo
}

func (st *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	ClusterId string
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	// Wire name: 'spark_context_id'
	SparkContextId string

	ForceSendFields []string `tf:"-"`
}

func (st *ClusterInstance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterInstancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterInstanceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterInstance) MarshalJSON() ([]byte, error) {
	pb, err := clusterInstanceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	// Wire name: 'job_cluster_key'
	JobClusterKey string
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec

	ForceSendFields []string `tf:"-"`
}

func (st *ClusterSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterSpec) MarshalJSON() ([]byte, error) {
	pb, err := clusterSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ComputeConfig struct {
	// IDof the GPU pool to use.
	// Wire name: 'gpu_node_pool_id'
	GpuNodePoolId string
	// GPU type.
	// Wire name: 'gpu_type'
	GpuType string
	// Number of GPUs.
	// Wire name: 'num_gpus'
	NumGpus int

	ForceSendFields []string `tf:"-"`
}

func (st *ComputeConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computeConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := computeConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComputeConfig) MarshalJSON() ([]byte, error) {
	pb, err := computeConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns Condition to satisfy [pflag.Value] interface
func (f *Condition) Type() string {
	return "Condition"
}

type ConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	// Wire name: 'left'
	Left string
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
	Op ConditionTaskOp
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	// Wire name: 'right'
	Right string
}

func (st *ConditionTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &conditionTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := conditionTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ConditionTask) MarshalJSON() ([]byte, error) {
	pb, err := conditionTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns ConditionTaskOp to satisfy [pflag.Value] interface
func (f *ConditionTaskOp) Type() string {
	return "ConditionTaskOp"
}

type Continuous struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	// Wire name: 'pause_status'
	PauseStatus PauseStatus
}

func (st *Continuous) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &continuousPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := continuousFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Continuous) MarshalJSON() ([]byte, error) {
	pb, err := continuousToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateJob struct {
	// List of permissions to set on the job.
	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlRequest
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	// Wire name: 'continuous'
	Continuous *Continuous
	// Deployment information for jobs managed by external sources.
	// Wire name: 'deployment'
	Deployment *JobDeployment
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	// Wire name: 'description'
	Description string
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	// Wire name: 'edit_mode'
	EditMode JobEditMode
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	// Wire name: 'environments'
	Environments []JobEnvironment
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	// Wire name: 'format'
	Format Format
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
	GitSource *GitSource
	// An optional set of health rules that can be defined for this job.
	// Wire name: 'health'
	Health *JobsHealthRules
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster
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
	MaxConcurrentRuns int
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'name'
	Name string
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	// Wire name: 'notification_settings'
	NotificationSettings *JobNotificationSettings
	// Job-level parameter definitions
	// Wire name: 'parameters'
	Parameters []JobParameterDefinition
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget
	// The queue settings of the job.
	// Wire name: 'queue'
	Queue *QueueSettings
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	// Wire name: 'run_as'
	RunAs *JobRunAs
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	// Wire name: 'schedule'
	Schedule *CronSchedule
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	// Wire name: 'tags'
	Tags map[string]string
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	// Wire name: 'tasks'
	Tasks []Task
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	// Wire name: 'trigger'
	Trigger *TriggerSettings
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string `tf:"-"`
}

func (st *CreateJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateJob) MarshalJSON() ([]byte, error) {
	pb, err := createJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Job was created successfully
type CreateResponse struct {
	// The canonical identifier for the newly created job.
	// Wire name: 'job_id'
	JobId int64

	ForceSendFields []string `tf:"-"`
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := createResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus PauseStatus
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	// Wire name: 'quartz_cron_expression'
	QuartzCronExpression string
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	// Wire name: 'timezone_id'
	TimezoneId string
}

func (st *CronSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cronSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cronScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CronSchedule) MarshalJSON() ([]byte, error) {
	pb, err := cronScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DashboardPageSnapshot struct {

	// Wire name: 'page_display_name'
	PageDisplayName string

	// Wire name: 'widget_error_details'
	WidgetErrorDetails []WidgetErrorDetail

	ForceSendFields []string `tf:"-"`
}

func (st *DashboardPageSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardPageSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardPageSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardPageSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := dashboardPageSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Configures the Lakeview Dashboard job task type.
type DashboardTask struct {
	// The identifier of the dashboard to refresh.
	// Wire name: 'dashboard_id'
	DashboardId string
	// Optional: subscription configuration for sending the dashboard snapshot.
	// Wire name: 'subscription'
	Subscription *Subscription
	// Optional: The warehouse id to execute the dashboard with for the
	// schedule. If not specified, the default warehouse of the dashboard will
	// be used.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *DashboardTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardTask) MarshalJSON() ([]byte, error) {
	pb, err := dashboardTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DashboardTaskOutput struct {
	// Should only be populated for manual PDF download jobs.
	// Wire name: 'page_snapshots'
	PageSnapshots []DashboardPageSnapshot
}

func (st *DashboardTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := dashboardTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	// Wire name: 'artifacts_headers'
	ArtifactsHeaders map[string]string
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	// Wire name: 'artifacts_link'
	ArtifactsLink string

	ForceSendFields []string `tf:"-"`
}

func (st *DbtOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dbtOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dbtOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DbtOutput) MarshalJSON() ([]byte, error) {
	pb, err := dbtOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	// Wire name: 'catalog'
	Catalog string
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	// Wire name: 'commands'
	Commands []string
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	// Wire name: 'profiles_directory'
	ProfilesDirectory string
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	// Wire name: 'project_directory'
	ProjectDirectory string
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	// Wire name: 'schema'
	Schema string
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	// Wire name: 'source'
	Source Source
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *DbtTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dbtTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dbtTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DbtTask) MarshalJSON() ([]byte, error) {
	pb, err := dbtTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	// Wire name: 'job_id'
	JobId int64
}

func (st *DeleteJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteJob) MarshalJSON() ([]byte, error) {
	pb, err := deleteJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteResponse struct {
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRun struct {
	// ID of the run to delete.
	// Wire name: 'run_id'
	RunId int64
}

func (st *DeleteRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRun) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteRunResponse struct {
}

func (st *DeleteRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents a change to the job cluster's settings that would be required for
// the job clusters to become compliant with their policies.
type EnforcePolicyComplianceForJobResponseJobClusterSettingsChange struct {
	// The field where this change would be made, prepended with the job cluster
	// key.
	// Wire name: 'field'
	Field string
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	// Wire name: 'new_value'
	NewValue string
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	// Wire name: 'previous_value'
	PreviousValue string

	ForceSendFields []string `tf:"-"`
}

func (st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) MarshalJSON() ([]byte, error) {
	pb, err := enforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnforcePolicyComplianceRequest struct {
	// The ID of the job you want to enforce policy compliance on.
	// Wire name: 'job_id'
	JobId int64
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	// Wire name: 'validate_only'
	ValidateOnly bool

	ForceSendFields []string `tf:"-"`
}

func (st *EnforcePolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforcePolicyComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforcePolicyComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforcePolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := enforcePolicyComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnforcePolicyComplianceResponse struct {
	// Whether any changes have been made to the job cluster settings for the
	// job to become compliant with its policies.
	// Wire name: 'has_changes'
	HasChanges bool
	// A list of job cluster changes that have been made to the job’s cluster
	// settings in order for all job clusters to become compliant with their
	// policies.
	// Wire name: 'job_cluster_changes'
	JobClusterChanges []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange
	// Updated job settings after policy enforcement. Policy enforcement only
	// applies to job clusters that are created when running the job (which are
	// specified in new_cluster) and does not apply to existing all-purpose
	// clusters. Updated job settings are derived by applying policy default
	// values to the existing job clusters in order to satisfy policy
	// requirements.
	// Wire name: 'settings'
	Settings *JobSettings

	ForceSendFields []string `tf:"-"`
}

func (st *EnforcePolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforcePolicyComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforcePolicyComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforcePolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := enforcePolicyComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Run was exported successfully.
type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	// Wire name: 'views'
	Views []ViewItem
}

func (st *ExportRunOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportRunOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportRunOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportRunOutput) MarshalJSON() ([]byte, error) {
	pb, err := exportRunOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Export and retrieve a job run
type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	// Wire name: 'run_id'
	RunId int64 `tf:"-"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	// Wire name: 'views_to_export'
	ViewsToExport ViewsToExport `tf:"-"`
}

func (st *ExportRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := exportRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	// Wire name: 'min_time_between_triggers_seconds'
	MinTimeBetweenTriggersSeconds int
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	// Wire name: 'url'
	Url string
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	// Wire name: 'wait_after_last_change_seconds'
	WaitAfterLastChangeSeconds int

	ForceSendFields []string `tf:"-"`
}

func (st *FileArrivalTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileArrivalTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileArrivalTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileArrivalTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := fileArrivalTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	// Wire name: 'error_message_stats'
	ErrorMessageStats []ForEachTaskErrorMessageStats
	// Describes stats of the iteration. Only latest retries are considered.
	// Wire name: 'task_run_stats'
	TaskRunStats *ForEachTaskTaskRunStats
}

func (st *ForEachStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachStats) MarshalJSON() ([]byte, error) {
	pb, err := forEachStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	// Wire name: 'concurrency'
	Concurrency int
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	// Wire name: 'inputs'
	Inputs string
	// Configuration for the task that will be run for each element in the array
	// Wire name: 'task'
	Task Task

	ForceSendFields []string `tf:"-"`
}

func (st *ForEachTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachTask) MarshalJSON() ([]byte, error) {
	pb, err := forEachTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	// Wire name: 'count'
	Count int
	// Describes the error message occured during the iterations.
	// Wire name: 'error_message'
	ErrorMessage string
	// Describes the termination reason for the error message.
	// Wire name: 'termination_category'
	TerminationCategory string

	ForceSendFields []string `tf:"-"`
}

func (st *ForEachTaskErrorMessageStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachTaskErrorMessageStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachTaskErrorMessageStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachTaskErrorMessageStats) MarshalJSON() ([]byte, error) {
	pb, err := forEachTaskErrorMessageStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	// Wire name: 'active_iterations'
	ActiveIterations int
	// Describes the number of failed and succeeded iteration runs.
	// Wire name: 'completed_iterations'
	CompletedIterations int
	// Describes the number of failed iteration runs.
	// Wire name: 'failed_iterations'
	FailedIterations int
	// Describes the number of iteration runs that have been scheduled.
	// Wire name: 'scheduled_iterations'
	ScheduledIterations int
	// Describes the number of succeeded iteration runs.
	// Wire name: 'succeeded_iterations'
	SucceededIterations int
	// Describes the length of the list of items to iterate over.
	// Wire name: 'total_iterations'
	TotalIterations int

	ForceSendFields []string `tf:"-"`
}

func (st *ForEachTaskTaskRunStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachTaskTaskRunStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachTaskTaskRunStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachTaskTaskRunStats) MarshalJSON() ([]byte, error) {
	pb, err := forEachTaskTaskRunStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

type GenAiComputeTask struct {
	// Command launcher to run the actual script, e.g. bash, python etc.
	// Wire name: 'command'
	Command string

	// Wire name: 'compute'
	Compute *ComputeConfig
	// Runtime image
	// Wire name: 'dl_runtime_image'
	DlRuntimeImage string
	// Optional string containing the name of the MLflow experiment to log the
	// run to. If name is not found, backend will create the mlflow experiment
	// using the name.
	// Wire name: 'mlflow_experiment_name'
	MlflowExperimentName string
	// Optional location type of the training script. When set to `WORKSPACE`,
	// the script will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the script will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`: Script
	// is located in Databricks workspace. * `GIT`: Script is located in cloud
	// Git provider.
	// Wire name: 'source'
	Source Source
	// The training script file path to be executed. Cloud file URIs (such as
	// dbfs:/, s3:/, adls:/, gcs:/) and workspace paths are supported. For
	// python files stored in the Databricks workspace, the path must be
	// absolute and begin with `/`. For files stored in a remote repository, the
	// path must be relative. This field is required.
	// Wire name: 'training_script_path'
	TrainingScriptPath string
	// Optional string containing model parameters passed to the training script
	// in yaml format. If present, then the content in yaml_parameters_file_path
	// will be ignored.
	// Wire name: 'yaml_parameters'
	YamlParameters string
	// Optional path to a YAML file containing model parameters passed to the
	// training script.
	// Wire name: 'yaml_parameters_file_path'
	YamlParametersFilePath string

	ForceSendFields []string `tf:"-"`
}

func (st *GenAiComputeTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genAiComputeTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genAiComputeTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenAiComputeTask) MarshalJSON() ([]byte, error) {
	pb, err := genAiComputeTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get job permission levels
type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	// Wire name: 'job_id'
	JobId string `tf:"-"`
}

func (st *GetJobPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getJobPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []JobPermissionsDescription
}

func (st *GetJobPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getJobPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get job permissions
type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	// Wire name: 'job_id'
	JobId string `tf:"-"`
}

func (st *GetJobPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getJobPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a single job
type GetJobRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	// Wire name: 'job_id'
	JobId int64 `tf:"-"`
	// Use `next_page_token` returned from the previous GetJob response to
	// request the next page of the job's array properties.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetJobRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobRequest) MarshalJSON() ([]byte, error) {
	pb, err := getJobRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get job policy compliance
type GetPolicyComplianceRequest struct {
	// The ID of the job whose compliance status you are requesting.
	// Wire name: 'job_id'
	JobId int64 `tf:"-"`
}

func (st *GetPolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPolicyComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPolicyComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPolicyComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetPolicyComplianceResponse struct {
	// Whether the job is compliant with its policies or not. Jobs could be out
	// of compliance if a policy they are using was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	// Wire name: 'is_compliant'
	IsCompliant bool
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	// Wire name: 'violations'
	Violations map[string]string

	ForceSendFields []string `tf:"-"`
}

func (st *GetPolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPolicyComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPolicyComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPolicyComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get the output for a single run
type GetRunOutputRequest struct {
	// The canonical identifier for the run.
	// Wire name: 'run_id'
	RunId int64 `tf:"-"`
}

func (st *GetRunOutputRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunOutputRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunOutputRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunOutputRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRunOutputRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a single job run
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
	RunId int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns GitProvider to satisfy [pflag.Value] interface
func (f *GitProvider) Type() string {
	return "GitProvider"
}

// Read-only state of the remote repository at the time the job was run. This
// field is only included on job runs.
type GitSnapshot struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	// Wire name: 'used_commit'
	UsedCommit string

	ForceSendFields []string `tf:"-"`
}

func (st *GitSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gitSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gitSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GitSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := gitSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	GitBranch string
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	// Wire name: 'git_commit'
	GitCommit string
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	// Wire name: 'git_provider'
	GitProvider GitProvider
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	// Wire name: 'git_snapshot'
	GitSnapshot *GitSnapshot
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	// Wire name: 'git_tag'
	GitTag string
	// URL of the repository to be cloned by this job.
	// Wire name: 'git_url'
	GitUrl string
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	// Wire name: 'job_source'
	JobSource *JobSource

	ForceSendFields []string `tf:"-"`
}

func (st *GitSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gitSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gitSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GitSource) MarshalJSON() ([]byte, error) {
	pb, err := gitSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Job was retrieved successfully.
type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	// Wire name: 'created_time'
	CreatedTime int64
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	// Wire name: 'effective_budget_policy_id'
	EffectiveBudgetPolicyId string
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId int64
	// A token that can be used to list the next page of array properties.
	// Wire name: 'next_page_token'
	NextPageToken string
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	// Wire name: 'run_as_user_name'
	RunAsUserName string
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	// Wire name: 'settings'
	Settings *JobSettings

	ForceSendFields []string `tf:"-"`
}

func (st *Job) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Job) MarshalJSON() ([]byte, error) {
	pb, err := jobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel JobPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *JobAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := jobAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []JobPermission
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

func (st *JobAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := jobAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	// Wire name: 'job_cluster_key'
	JobClusterKey string
	// If new_cluster, a description of a cluster that is created for each task.
	// Wire name: 'new_cluster'
	NewCluster compute.ClusterSpec
}

func (st *JobCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobCluster) MarshalJSON() ([]byte, error) {
	pb, err := jobClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobCompliance struct {
	// Whether this job is in compliance with the latest version of its policy.
	// Wire name: 'is_compliant'
	IsCompliant bool
	// Canonical unique identifier for a job.
	// Wire name: 'job_id'
	JobId int64
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	// Wire name: 'violations'
	Violations map[string]string

	ForceSendFields []string `tf:"-"`
}

func (st *JobCompliance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobCompliancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobComplianceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobCompliance) MarshalJSON() ([]byte, error) {
	pb, err := jobComplianceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobDeployment struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	// Wire name: 'kind'
	Kind JobDeploymentKind
	// Path of the file that contains deployment metadata.
	// Wire name: 'metadata_file_path'
	MetadataFilePath string

	ForceSendFields []string `tf:"-"`
}

func (st *JobDeployment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobDeploymentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobDeploymentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobDeployment) MarshalJSON() ([]byte, error) {
	pb, err := jobDeploymentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns JobDeploymentKind to satisfy [pflag.Value] interface
func (f *JobDeploymentKind) Type() string {
	return "JobDeploymentKind"
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

// Type always returns JobEditMode to satisfy [pflag.Value] interface
func (f *JobEditMode) Type() string {
	return "JobEditMode"
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	// Wire name: 'on_duration_warning_threshold_exceeded'
	OnDurationWarningThresholdExceeded []string
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	// Wire name: 'on_failure'
	OnFailure []string
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_start'
	OnStart []string
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	// Wire name: 'on_streaming_backlog_exceeded'
	OnStreamingBacklogExceeded []string
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_success'
	OnSuccess []string

	ForceSendFields []string `tf:"-"`
}

func (st *JobEmailNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobEmailNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobEmailNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobEmailNotifications) MarshalJSON() ([]byte, error) {
	pb, err := jobEmailNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobEnvironment struct {
	// The key of an environment. It has to be unique within a job.
	// Wire name: 'environment_key'
	EnvironmentKey string
	// The environment entity used to preserve serverless environment side
	// panel, jobs' environment for non-notebook task, and DLT's environment for
	// classic and serverless pipelines. In this minimal environment spec, only
	// pip dependencies are supported.
	// Wire name: 'spec'
	Spec *compute.Environment
}

func (st *JobEnvironment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobEnvironmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobEnvironmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobEnvironment) MarshalJSON() ([]byte, error) {
	pb, err := jobEnvironmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	// Wire name: 'no_alert_for_canceled_runs'
	NoAlertForCanceledRuns bool
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool

	ForceSendFields []string `tf:"-"`
}

func (st *JobNotificationSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobNotificationSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobNotificationSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobNotificationSettings) MarshalJSON() ([]byte, error) {
	pb, err := jobNotificationSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobParameter struct {
	// The optional default value of the parameter
	// Wire name: 'default'
	Default string
	// The name of the parameter
	// Wire name: 'name'
	Name string
	// The value used in the run
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *JobParameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobParameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobParameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobParameter) MarshalJSON() ([]byte, error) {
	pb, err := jobParameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	// Wire name: 'default'
	Default string
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	// Wire name: 'name'
	Name string
}

func (st *JobParameterDefinition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobParameterDefinitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobParameterDefinitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobParameterDefinition) MarshalJSON() ([]byte, error) {
	pb, err := jobParameterDefinitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel JobPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *JobPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermission) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns JobPermissionLevel to satisfy [pflag.Value] interface
func (f *JobPermissionLevel) Type() string {
	return "JobPermissionLevel"
}

type JobPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func (st *JobPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermissions) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel JobPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *JobPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlRequest
	// The job for which to get or manage permissions.
	// Wire name: 'job_id'
	JobId string `tf:"-"`
}

func (st *JobPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	ServicePrincipalName string
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *JobRunAs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobRunAsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobRunAsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobRunAs) MarshalJSON() ([]byte, error) {
	pb, err := jobRunAsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type JobSettings struct {
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	// Wire name: 'continuous'
	Continuous *Continuous
	// Deployment information for jobs managed by external sources.
	// Wire name: 'deployment'
	Deployment *JobDeployment
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	// Wire name: 'description'
	Description string
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	// Wire name: 'edit_mode'
	EditMode JobEditMode
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	// Wire name: 'environments'
	Environments []JobEnvironment
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	// Wire name: 'format'
	Format Format
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
	GitSource *GitSource
	// An optional set of health rules that can be defined for this job.
	// Wire name: 'health'
	Health *JobsHealthRules
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster
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
	MaxConcurrentRuns int
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'name'
	Name string
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	// Wire name: 'notification_settings'
	NotificationSettings *JobNotificationSettings
	// Job-level parameter definitions
	// Wire name: 'parameters'
	Parameters []JobParameterDefinition
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget
	// The queue settings of the job.
	// Wire name: 'queue'
	Queue *QueueSettings
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	// Wire name: 'run_as'
	RunAs *JobRunAs
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	// Wire name: 'schedule'
	Schedule *CronSchedule
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	// Wire name: 'tags'
	Tags map[string]string
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	// Wire name: 'tasks'
	Tasks []Task
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	// Wire name: 'trigger'
	Trigger *TriggerSettings
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string `tf:"-"`
}

func (st *JobSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSettings) MarshalJSON() ([]byte, error) {
	pb, err := jobSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	DirtyState JobSourceDirtyState
	// Name of the branch which the job is imported from.
	// Wire name: 'import_from_git_branch'
	ImportFromGitBranch string
	// Path of the job YAML file that contains the job specification.
	// Wire name: 'job_config_path'
	JobConfigPath string
}

func (st *JobSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSource) MarshalJSON() ([]byte, error) {
	pb, err := jobSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns JobSourceDirtyState to satisfy [pflag.Value] interface
func (f *JobSourceDirtyState) Type() string {
	return "JobSourceDirtyState"
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

// Type always returns JobsHealthMetric to satisfy [pflag.Value] interface
func (f *JobsHealthMetric) Type() string {
	return "JobsHealthMetric"
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

// Type always returns JobsHealthOperator to satisfy [pflag.Value] interface
func (f *JobsHealthOperator) Type() string {
	return "JobsHealthOperator"
}

type JobsHealthRule struct {
	// Specifies the health metric that is being evaluated for a particular
	// health rule.
	//
	// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
	// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data
	// waiting to be consumed across all streams. This metric is in Public
	// Preview. * `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset
	// lag across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_SECONDS`: An estimate of the maximum consumer delay
	// across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_FILES`: An estimate of the maximum number of
	// outstanding files across all streams. This metric is in Public Preview.
	// Wire name: 'metric'
	Metric JobsHealthMetric
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	// Wire name: 'op'
	Op JobsHealthOperator
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	// Wire name: 'value'
	Value int64
}

func (st *JobsHealthRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobsHealthRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobsHealthRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobsHealthRule) MarshalJSON() ([]byte, error) {
	pb, err := jobsHealthRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {

	// Wire name: 'rules'
	Rules []JobsHealthRule
}

func (st *JobsHealthRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobsHealthRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobsHealthRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobsHealthRules) MarshalJSON() ([]byte, error) {
	pb, err := jobsHealthRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListJobComplianceForPolicyResponse struct {
	// A list of jobs and their policy compliance statuses.
	// Wire name: 'jobs'
	Jobs []JobCompliance
	// This field represents the pagination token to retrieve the next page of
	// results. If this field is not in the response, it means no further
	// results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string
	// This field represents the pagination token to retrieve the previous page
	// of results. If this field is not in the response, it means no further
	// results for the request.
	// Wire name: 'prev_page_token'
	PrevPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListJobComplianceForPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobComplianceForPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobComplianceForPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobComplianceForPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := listJobComplianceForPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List job policy compliance
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
	PolicyId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListJobComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := listJobComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List jobs
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
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListJobsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listJobsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List of jobs was retrieved successfully.
type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	// Wire name: 'has_more'
	HasMore bool
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	// Wire name: 'jobs'
	Jobs []BaseJob
	// A token that can be used to list the next page of jobs (if applicable).
	// Wire name: 'next_page_token'
	NextPageToken string
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	// Wire name: 'prev_page_token'
	PrevPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListJobsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listJobsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List job runs
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
	StartTimeTo int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListRunsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRunsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRunsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRunsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listRunsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List of runs was retrieved successfully.
type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	// Wire name: 'has_more'
	HasMore bool
	// A token that can be used to list the next page of runs (if applicable).
	// Wire name: 'next_page_token'
	NextPageToken string
	// A token that can be used to list the previous page of runs (if
	// applicable).
	// Wire name: 'prev_page_token'
	PrevPageToken string
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	// Wire name: 'runs'
	Runs []BaseRun

	ForceSendFields []string `tf:"-"`
}

func (st *ListRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	// Wire name: 'result'
	Result string
	// Whether or not the result was truncated.
	// Wire name: 'truncated'
	Truncated bool

	ForceSendFields []string `tf:"-"`
}

func (st *NotebookOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notebookOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notebookOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotebookOutput) MarshalJSON() ([]byte, error) {
	pb, err := notebookOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	BaseParameters map[string]string
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	// Wire name: 'notebook_path'
	NotebookPath string
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	// Wire name: 'source'
	Source Source
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *NotebookTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notebookTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notebookTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotebookTask) MarshalJSON() ([]byte, error) {
	pb, err := notebookTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Stores the catalog name, schema name, and the output schema expiration time
// for the clean room run.
type OutputSchemaInfo struct {

	// Wire name: 'catalog_name'
	CatalogName string
	// The expiration time for the output schema as a Unix timestamp in
	// milliseconds.
	// Wire name: 'expiration_time'
	ExpirationTime int64

	// Wire name: 'schema_name'
	SchemaName string

	ForceSendFields []string `tf:"-"`
}

func (st *OutputSchemaInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &outputSchemaInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := outputSchemaInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OutputSchemaInfo) MarshalJSON() ([]byte, error) {
	pb, err := outputSchemaInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns PauseStatus to satisfy [pflag.Value] interface
func (f *PauseStatus) Type() string {
	return "PauseStatus"
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

// Type always returns PerformanceTarget to satisfy [pflag.Value] interface
func (f *PerformanceTarget) Type() string {
	return "PerformanceTarget"
}

type PeriodicTriggerConfiguration struct {
	// The interval at which the trigger should run.
	// Wire name: 'interval'
	Interval int
	// The unit of time for the interval.
	// Wire name: 'unit'
	Unit PeriodicTriggerConfigurationTimeUnit
}

func (st *PeriodicTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &periodicTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := periodicTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PeriodicTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := periodicTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns PeriodicTriggerConfigurationTimeUnit to satisfy [pflag.Value] interface
func (f *PeriodicTriggerConfigurationTimeUnit) Type() string {
	return "PeriodicTriggerConfigurationTimeUnit"
}

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	// Wire name: 'full_refresh'
	FullRefresh bool

	ForceSendFields []string `tf:"-"`
}

func (st *PipelineParams) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineParamsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineParamsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineParams) MarshalJSON() ([]byte, error) {
	pb, err := pipelineParamsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PipelineTask struct {
	// If true, triggers a full refresh on the delta live table.
	// Wire name: 'full_refresh'
	FullRefresh bool
	// The full name of the pipeline task to execute.
	// Wire name: 'pipeline_id'
	PipelineId string

	ForceSendFields []string `tf:"-"`
}

func (st *PipelineTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineTask) MarshalJSON() ([]byte, error) {
	pb, err := pipelineTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PowerBiModel struct {
	// How the published Power BI model authenticates to Databricks
	// Wire name: 'authentication_method'
	AuthenticationMethod AuthenticationMethod
	// The name of the Power BI model
	// Wire name: 'model_name'
	ModelName string
	// Whether to overwrite existing Power BI models
	// Wire name: 'overwrite_existing'
	OverwriteExisting bool
	// The default storage mode of the Power BI model
	// Wire name: 'storage_mode'
	StorageMode StorageMode
	// The name of the Power BI workspace of the model
	// Wire name: 'workspace_name'
	WorkspaceName string

	ForceSendFields []string `tf:"-"`
}

func (st *PowerBiModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &powerBiModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := powerBiModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PowerBiModel) MarshalJSON() ([]byte, error) {
	pb, err := powerBiModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PowerBiTable struct {
	// The catalog name in Databricks
	// Wire name: 'catalog'
	Catalog string
	// The table name in Databricks
	// Wire name: 'name'
	Name string
	// The schema name in Databricks
	// Wire name: 'schema'
	Schema string
	// The Power BI storage mode of the table
	// Wire name: 'storage_mode'
	StorageMode StorageMode

	ForceSendFields []string `tf:"-"`
}

func (st *PowerBiTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &powerBiTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := powerBiTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PowerBiTable) MarshalJSON() ([]byte, error) {
	pb, err := powerBiTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PowerBiTask struct {
	// The resource name of the UC connection to authenticate from Databricks to
	// Power BI
	// Wire name: 'connection_resource_name'
	ConnectionResourceName string
	// The semantic model to update
	// Wire name: 'power_bi_model'
	PowerBiModel *PowerBiModel
	// Whether the model should be refreshed after the update
	// Wire name: 'refresh_after_update'
	RefreshAfterUpdate bool
	// The tables to be exported to Power BI
	// Wire name: 'tables'
	Tables []PowerBiTable
	// The SQL warehouse ID to use as the Power BI data source
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *PowerBiTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &powerBiTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := powerBiTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PowerBiTask) MarshalJSON() ([]byte, error) {
	pb, err := powerBiTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	// Wire name: 'entry_point'
	EntryPoint string
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	// Wire name: 'named_parameters'
	NamedParameters map[string]string
	// Name of the package to execute
	// Wire name: 'package_name'
	PackageName string
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	// Wire name: 'parameters'
	Parameters []string
}

func (st *PythonWheelTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pythonWheelTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pythonWheelTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PythonWheelTask) MarshalJSON() ([]byte, error) {
	pb, err := pythonWheelTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueueDetails struct {
	// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run
	// was queued due to reaching the workspace limit of active task runs. *
	// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the
	// per-job limit of concurrent job runs. *
	// `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run was queued due to reaching
	// the workspace limit of active run job tasks.
	// Wire name: 'code'
	Code QueueDetailsCodeCode
	// A descriptive message with the queuing details. This field is
	// unstructured, and its exact format is subject to change.
	// Wire name: 'message'
	Message string

	ForceSendFields []string `tf:"-"`
}

func (st *QueueDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queueDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queueDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueueDetails) MarshalJSON() ([]byte, error) {
	pb, err := queueDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns QueueDetailsCodeCode to satisfy [pflag.Value] interface
func (f *QueueDetailsCodeCode) Type() string {
	return "QueueDetailsCodeCode"
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	// Wire name: 'enabled'
	Enabled bool
}

func (st *QueueSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queueSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queueSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueueSettings) MarshalJSON() ([]byte, error) {
	pb, err := queueSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	EffectivePerformanceTarget PerformanceTarget
	// The end time of the (repaired) run.
	// Wire name: 'end_time'
	EndTime int64
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	// Wire name: 'id'
	Id int64
	// The start time of the (repaired) run.
	// Wire name: 'start_time'
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState
	// The current status of the run
	// Wire name: 'status'
	Status *RunStatus
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	// Wire name: 'task_run_ids'
	TaskRunIds []int64
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	// Wire name: 'type'
	Type RepairHistoryItemType

	ForceSendFields []string `tf:"-"`
}

func (st *RepairHistoryItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repairHistoryItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repairHistoryItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepairHistoryItem) MarshalJSON() ([]byte, error) {
	pb, err := repairHistoryItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns RepairHistoryItemType to satisfy [pflag.Value] interface
func (f *RepairHistoryItemType) Type() string {
	return "RepairHistoryItemType"
}

type RepairRun struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string
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
	JarParams []string
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	// Wire name: 'job_parameters'
	JobParameters map[string]string
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	// Wire name: 'latest_repair_id'
	LatestRepairId int64
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
	NotebookParams map[string]string
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string
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
	PythonParams []string
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	// Wire name: 'rerun_all_failed_tasks'
	RerunAllFailedTasks bool
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	// Wire name: 'rerun_dependent_tasks'
	RerunDependentTasks bool
	// The task keys of the task runs to repair.
	// Wire name: 'rerun_tasks'
	RerunTasks []string
	// The job run ID of the run to repair. The run must not be in progress.
	// Wire name: 'run_id'
	RunId int64
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
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams map[string]string

	ForceSendFields []string `tf:"-"`
}

func (st *RepairRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repairRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repairRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepairRun) MarshalJSON() ([]byte, error) {
	pb, err := repairRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Run repair was initiated.
type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	// Wire name: 'repair_id'
	RepairId int64

	ForceSendFields []string `tf:"-"`
}

func (st *RepairRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repairRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repairRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepairRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := repairRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	// Wire name: 'job_id'
	JobId int64
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	// Wire name: 'new_settings'
	NewSettings JobSettings
}

func (st *ResetJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resetJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resetJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResetJob) MarshalJSON() ([]byte, error) {
	pb, err := resetJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResetResponse struct {
}

func (st *ResetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resetResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resetResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResetResponse) MarshalJSON() ([]byte, error) {
	pb, err := resetResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedConditionTaskValues struct {

	// Wire name: 'left'
	Left string

	// Wire name: 'right'
	Right string

	ForceSendFields []string `tf:"-"`
}

func (st *ResolvedConditionTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedConditionTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedConditionTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedConditionTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedConditionTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedDbtTaskValues struct {

	// Wire name: 'commands'
	Commands []string
}

func (st *ResolvedDbtTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedDbtTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedDbtTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedDbtTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedDbtTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedNotebookTaskValues struct {

	// Wire name: 'base_parameters'
	BaseParameters map[string]string
}

func (st *ResolvedNotebookTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedNotebookTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedNotebookTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedNotebookTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedNotebookTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedParamPairValues struct {

	// Wire name: 'parameters'
	Parameters map[string]string
}

func (st *ResolvedParamPairValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedParamPairValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedParamPairValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedParamPairValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedParamPairValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedPythonWheelTaskValues struct {

	// Wire name: 'named_parameters'
	NamedParameters map[string]string

	// Wire name: 'parameters'
	Parameters []string
}

func (st *ResolvedPythonWheelTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedPythonWheelTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedPythonWheelTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedPythonWheelTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedPythonWheelTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedRunJobTaskValues struct {

	// Wire name: 'job_parameters'
	JobParameters map[string]string

	// Wire name: 'parameters'
	Parameters map[string]string
}

func (st *ResolvedRunJobTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedRunJobTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedRunJobTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedRunJobTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedRunJobTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedStringParamsValues struct {

	// Wire name: 'parameters'
	Parameters []string
}

func (st *ResolvedStringParamsValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedStringParamsValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedStringParamsValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedStringParamsValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedStringParamsValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResolvedValues struct {

	// Wire name: 'condition_task'
	ConditionTask *ResolvedConditionTaskValues

	// Wire name: 'dbt_task'
	DbtTask *ResolvedDbtTaskValues

	// Wire name: 'notebook_task'
	NotebookTask *ResolvedNotebookTaskValues

	// Wire name: 'python_wheel_task'
	PythonWheelTask *ResolvedPythonWheelTaskValues

	// Wire name: 'run_job_task'
	RunJobTask *ResolvedRunJobTaskValues

	// Wire name: 'simulation_task'
	SimulationTask *ResolvedParamPairValues

	// Wire name: 'spark_jar_task'
	SparkJarTask *ResolvedStringParamsValues

	// Wire name: 'spark_python_task'
	SparkPythonTask *ResolvedStringParamsValues

	// Wire name: 'spark_submit_task'
	SparkSubmitTask *ResolvedStringParamsValues

	// Wire name: 'sql_task'
	SqlTask *ResolvedParamPairValues
}

func (st *ResolvedValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	AttemptNumber int
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	// Wire name: 'cleanup_duration'
	CleanupDuration int64
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	// Wire name: 'cluster_instance'
	ClusterInstance *ClusterInstance
	// A snapshot of the job’s cluster specification when this run was
	// created.
	// Wire name: 'cluster_spec'
	ClusterSpec *ClusterSpec
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// Description of the run
	// Wire name: 'description'
	Description string
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	// Wire name: 'end_time'
	EndTime int64
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	// Wire name: 'execution_duration'
	ExecutionDuration int64
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
	GitSource *GitSource
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	// Wire name: 'has_more'
	HasMore bool
	// Only populated by for-each iterations. The parent for-each task is
	// located in tasks array.
	// Wire name: 'iterations'
	Iterations []RunTask
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	// Wire name: 'job_clusters'
	JobClusters []JobCluster
	// The canonical identifier of the job that contains this run.
	// Wire name: 'job_id'
	JobId int64
	// Job-level parameters used in the run
	// Wire name: 'job_parameters'
	JobParameters []JobParameter
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	// Wire name: 'job_run_id'
	JobRunId int64
	// A token that can be used to list the next page of array properties.
	// Wire name: 'next_page_token'
	NextPageToken string
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	// Wire name: 'number_in_job'
	NumberInJob int64
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	// Wire name: 'original_attempt_run_id'
	OriginalAttemptRunId int64
	// The parameters used for this run.
	// Wire name: 'overriding_parameters'
	OverridingParameters *RunParameters
	// The time in milliseconds that the run has spent in the queue.
	// Wire name: 'queue_duration'
	QueueDuration int64
	// The repair history of the run.
	// Wire name: 'repair_history'
	RepairHistory []RepairHistoryItem
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	// Wire name: 'run_duration'
	RunDuration int64
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'run_id'
	RunId int64
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	// Wire name: 'run_name'
	RunName string
	// The URL to the detail page of the run.
	// Wire name: 'run_page_url'
	RunPageUrl string
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	// Wire name: 'run_type'
	RunType RunType
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	// Wire name: 'schedule'
	Schedule *CronSchedule
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	// Wire name: 'setup_duration'
	SetupDuration int64
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	// Wire name: 'start_time'
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState
	// The current status of the run
	// Wire name: 'status'
	Status *RunStatus
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	// Wire name: 'tasks'
	Tasks []RunTask
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	// Wire name: 'trigger'
	Trigger TriggerType
	// Additional details about what triggered the run
	// Wire name: 'trigger_info'
	TriggerInfo *TriggerInfo

	ForceSendFields []string `tf:"-"`
}

func (st *Run) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Run) MarshalJSON() ([]byte, error) {
	pb, err := runToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RunConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	// Wire name: 'left'
	Left string
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
	Op ConditionTaskOp
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	// Wire name: 'outcome'
	Outcome string
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	// Wire name: 'right'
	Right string

	ForceSendFields []string `tf:"-"`
}

func (st *RunConditionTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runConditionTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runConditionTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunConditionTask) MarshalJSON() ([]byte, error) {
	pb, err := runConditionTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RunForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	// Wire name: 'concurrency'
	Concurrency int
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	// Wire name: 'inputs'
	Inputs string
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of an For each task
	// Wire name: 'stats'
	Stats *ForEachStats
	// Configuration for the task that will be run for each element in the array
	// Wire name: 'task'
	Task Task

	ForceSendFields []string `tf:"-"`
}

func (st *RunForEachTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runForEachTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runForEachTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunForEachTask) MarshalJSON() ([]byte, error) {
	pb, err := runForEachTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns RunIf to satisfy [pflag.Value] interface
func (f *RunIf) Type() string {
	return "RunIf"
}

type RunJobOutput struct {
	// The run id of the triggered job run
	// Wire name: 'run_id'
	RunId int64

	ForceSendFields []string `tf:"-"`
}

func (st *RunJobOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runJobOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runJobOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunJobOutput) MarshalJSON() ([]byte, error) {
	pb, err := runJobOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RunJobTask struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string
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
	JarParams []string
	// ID of the job to trigger.
	// Wire name: 'job_id'
	JobId int64
	// Job-level parameters used to trigger the job.
	// Wire name: 'job_parameters'
	JobParameters map[string]string
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
	NotebookParams map[string]string
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string
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
	PythonParams []string
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
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams map[string]string
}

func (st *RunJobTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runJobTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runJobTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunJobTask) MarshalJSON() ([]byte, error) {
	pb, err := runJobTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns RunLifeCycleState to satisfy [pflag.Value] interface
func (f *RunLifeCycleState) Type() string {
	return "RunLifeCycleState"
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

// Type always returns RunLifecycleStateV2State to satisfy [pflag.Value] interface
func (f *RunLifecycleStateV2State) Type() string {
	return "RunLifecycleStateV2State"
}

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string
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
	IdempotencyToken string
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
	JarParams []string
	// The ID of the job to be executed
	// Wire name: 'job_id'
	JobId int64
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	// Wire name: 'job_parameters'
	JobParameters map[string]string
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
	NotebookParams map[string]string
	// A list of task keys to run inside of the job. If this field is not
	// provided, all tasks in the job will be run.
	// Wire name: 'only'
	Only []string
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'performance_target'
	PerformanceTarget PerformanceTarget
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string
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
	PythonParams []string
	// The queue settings of the run.
	// Wire name: 'queue'
	Queue *QueueSettings
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
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams map[string]string

	ForceSendFields []string `tf:"-"`
}

func (st *RunNow) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runNowPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runNowFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunNow) MarshalJSON() ([]byte, error) {
	pb, err := runNowToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Run was started successfully.
type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	// Wire name: 'number_in_job'
	NumberInJob int64
	// The globally unique ID of the newly triggered run.
	// Wire name: 'run_id'
	RunId int64

	ForceSendFields []string `tf:"-"`
}

func (st *RunNowResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runNowResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runNowResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunNowResponse) MarshalJSON() ([]byte, error) {
	pb, err := runNowResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Run output was retrieved successfully.
type RunOutput struct {
	// The output of a clean rooms notebook task, if available
	// Wire name: 'clean_rooms_notebook_output'
	CleanRoomsNotebookOutput *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput
	// The output of a dashboard task, if available
	// Wire name: 'dashboard_output'
	DashboardOutput *DashboardTaskOutput
	// The output of a dbt task, if available.
	// Wire name: 'dbt_output'
	DbtOutput *DbtOutput
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	// Wire name: 'error'
	Error string
	// If there was an error executing the run, this field contains any
	// available stack traces.
	// Wire name: 'error_trace'
	ErrorTrace string

	// Wire name: 'info'
	Info string
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	// Wire name: 'logs'
	Logs string
	// Whether the logs are truncated.
	// Wire name: 'logs_truncated'
	LogsTruncated bool
	// All details of the run except for its output.
	// Wire name: 'metadata'
	Metadata *Run
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	// Wire name: 'notebook_output'
	NotebookOutput *NotebookOutput
	// The output of a run job task, if available
	// Wire name: 'run_job_output'
	RunJobOutput *RunJobOutput
	// The output of a SQL task, if available.
	// Wire name: 'sql_output'
	SqlOutput *SqlOutput

	ForceSendFields []string `tf:"-"`
}

func (st *RunOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunOutput) MarshalJSON() ([]byte, error) {
	pb, err := runOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	// Wire name: 'dbt_commands'
	DbtCommands []string
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
	JarParams []string
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
	NotebookParams map[string]string
	// Controls whether the pipeline should perform a full refresh
	// Wire name: 'pipeline_params'
	PipelineParams *PipelineParams

	// Wire name: 'python_named_params'
	PythonNamedParams map[string]string
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
	PythonParams []string
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
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	// Wire name: 'sql_params'
	SqlParams map[string]string
}

func (st *RunParameters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runParametersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runParametersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunParameters) MarshalJSON() ([]byte, error) {
	pb, err := runParametersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	// Wire name: 'life_cycle_state'
	LifeCycleState RunLifeCycleState
	// The reason indicating why the run was queued.
	// Wire name: 'queue_reason'
	QueueReason string
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	// Wire name: 'result_state'
	ResultState RunResultState
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	// Wire name: 'state_message'
	StateMessage string
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	// Wire name: 'user_cancelled_or_timedout'
	UserCancelledOrTimedout bool

	ForceSendFields []string `tf:"-"`
}

func (st *RunState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunState) MarshalJSON() ([]byte, error) {
	pb, err := runStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The current status of the run
type RunStatus struct {
	// If the run was queued, details about the reason for queuing the run.
	// Wire name: 'queue_details'
	QueueDetails *QueueDetails
	// The current state of the run.
	// Wire name: 'state'
	State RunLifecycleStateV2State
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	// Wire name: 'termination_details'
	TerminationDetails *TerminationDetails
}

func (st *RunStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunStatus) MarshalJSON() ([]byte, error) {
	pb, err := runStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	AttemptNumber int
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	// Wire name: 'clean_rooms_notebook_task'
	CleanRoomsNotebookTask *CleanRoomsNotebookTask
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	// Wire name: 'cleanup_duration'
	CleanupDuration int64
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	// Wire name: 'cluster_instance'
	ClusterInstance *ClusterInstance
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	// Wire name: 'condition_task'
	ConditionTask *RunConditionTask
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	// Wire name: 'dashboard_task'
	DashboardTask *DashboardTask
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	// Wire name: 'dbt_task'
	DbtTask *DbtTask
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	// Wire name: 'depends_on'
	DependsOn []TaskDependency
	// An optional description for this task.
	// Wire name: 'description'
	Description string
	// Deprecated, field was never used in production.
	// Wire name: 'disabled'
	Disabled bool
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	// Wire name: 'effective_performance_target'
	EffectivePerformanceTarget PerformanceTarget
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	// Wire name: 'end_time'
	EndTime int64
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	// Wire name: 'environment_key'
	EnvironmentKey string
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	// Wire name: 'execution_duration'
	ExecutionDuration int64
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	// Wire name: 'for_each_task'
	ForEachTask *RunForEachTask

	// Wire name: 'gen_ai_compute_task'
	GenAiComputeTask *GenAiComputeTask
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	// Wire name: 'git_source'
	GitSource *GitSource
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	// Wire name: 'job_cluster_key'
	JobClusterKey string
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec
	// The task runs a notebook when the `notebook_task` field is present.
	// Wire name: 'notebook_task'
	NotebookTask *NotebookTask
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	// Wire name: 'notification_settings'
	NotificationSettings *TaskNotificationSettings
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	// Wire name: 'pipeline_task'
	PipelineTask *PipelineTask
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	// Wire name: 'power_bi_task'
	PowerBiTask *PowerBiTask
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	// Wire name: 'python_wheel_task'
	PythonWheelTask *PythonWheelTask
	// The time in milliseconds that the run has spent in the queue.
	// Wire name: 'queue_duration'
	QueueDuration int64
	// Parameter values including resolved references
	// Wire name: 'resolved_values'
	ResolvedValues *ResolvedValues
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	// Wire name: 'run_duration'
	RunDuration int64
	// The ID of the task run.
	// Wire name: 'run_id'
	RunId int64
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	// Wire name: 'run_if'
	RunIf RunIf
	// The task triggers another job when the `run_job_task` field is present.
	// Wire name: 'run_job_task'
	RunJobTask *RunJobTask

	// Wire name: 'run_page_url'
	RunPageUrl string
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	// Wire name: 'setup_duration'
	SetupDuration int64
	// The task runs a JAR when the `spark_jar_task` field is present.
	// Wire name: 'spark_jar_task'
	SparkJarTask *SparkJarTask
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	// Wire name: 'spark_python_task'
	SparkPythonTask *SparkPythonTask
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
	SparkSubmitTask *SparkSubmitTask
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	// Wire name: 'sql_task'
	SqlTask *SqlTask
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	// Wire name: 'start_time'
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	// Wire name: 'state'
	State *RunState
	// The current status of the run
	// Wire name: 'status'
	Status *RunStatus
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	// Wire name: 'task_key'
	TaskKey string
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string `tf:"-"`
}

func (st *RunTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunTask) MarshalJSON() ([]byte, error) {
	pb, err := runTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns RunType to satisfy [pflag.Value] interface
func (f *RunType) Type() string {
	return "RunType"
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

// Type always returns Source to satisfy [pflag.Value] interface
func (f *Source) Type() string {
	return "Source"
}

type SparkJarTask struct {
	// Deprecated since 04/2016. Provide a `jar` through the `libraries` field
	// instead. For an example, see :method:jobs/create.
	// Wire name: 'jar_uri'
	JarUri string
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	// Wire name: 'main_class_name'
	MainClassName string
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'parameters'
	Parameters []string
	// Deprecated. A value of `false` is no longer supported.
	// Wire name: 'run_as_repl'
	RunAsRepl bool

	ForceSendFields []string `tf:"-"`
}

func (st *SparkJarTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkJarTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkJarTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkJarTask) MarshalJSON() ([]byte, error) {
	pb, err := sparkJarTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'parameters'
	Parameters []string
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	// Wire name: 'python_file'
	PythonFile string
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
	Source Source
}

func (st *SparkPythonTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkPythonTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkPythonTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkPythonTask) MarshalJSON() ([]byte, error) {
	pb, err := sparkPythonTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// Wire name: 'parameters'
	Parameters []string
}

func (st *SparkSubmitTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkSubmitTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkSubmitTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkSubmitTask) MarshalJSON() ([]byte, error) {
	pb, err := sparkSubmitTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlAlertOutput struct {
	// The state of the SQL alert.
	//
	// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
	// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled
	// trigger conditions
	// Wire name: 'alert_state'
	AlertState SqlAlertState
	// The link to find the output results.
	// Wire name: 'output_link'
	OutputLink string
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	// Wire name: 'query_text'
	QueryText string
	// Information about SQL statements executed in the run.
	// Wire name: 'sql_statements'
	SqlStatements []SqlStatementOutput
	// The canonical identifier of the SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *SqlAlertOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlAlertOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlAlertOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlAlertOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlAlertOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns SqlAlertState to satisfy [pflag.Value] interface
func (f *SqlAlertState) Type() string {
	return "SqlAlertState"
}

type SqlDashboardOutput struct {
	// The canonical identifier of the SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId string
	// Widgets executed in the run. Only SQL query based widgets are listed.
	// Wire name: 'widgets'
	Widgets []SqlDashboardWidgetOutput

	ForceSendFields []string `tf:"-"`
}

func (st *SqlDashboardOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlDashboardOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlDashboardOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlDashboardOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlDashboardOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	// Wire name: 'end_time'
	EndTime int64
	// The information about the error when execution fails.
	// Wire name: 'error'
	Error *SqlOutputError
	// The link to find the output results.
	// Wire name: 'output_link'
	OutputLink string
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	// Wire name: 'start_time'
	StartTime int64
	// The execution status of the SQL widget.
	// Wire name: 'status'
	Status SqlDashboardWidgetOutputStatus
	// The canonical identifier of the SQL widget.
	// Wire name: 'widget_id'
	WidgetId string
	// The title of the SQL widget.
	// Wire name: 'widget_title'
	WidgetTitle string

	ForceSendFields []string `tf:"-"`
}

func (st *SqlDashboardWidgetOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlDashboardWidgetOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlDashboardWidgetOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlDashboardWidgetOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlDashboardWidgetOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns SqlDashboardWidgetOutputStatus to satisfy [pflag.Value] interface
func (f *SqlDashboardWidgetOutputStatus) Type() string {
	return "SqlDashboardWidgetOutputStatus"
}

type SqlOutput struct {
	// The output of a SQL alert task, if available.
	// Wire name: 'alert_output'
	AlertOutput *SqlAlertOutput
	// The output of a SQL dashboard task, if available.
	// Wire name: 'dashboard_output'
	DashboardOutput *SqlDashboardOutput
	// The output of a SQL query task, if available.
	// Wire name: 'query_output'
	QueryOutput *SqlQueryOutput
}

func (st *SqlOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlOutputError struct {
	// The error message when execution fails.
	// Wire name: 'message'
	Message string

	ForceSendFields []string `tf:"-"`
}

func (st *SqlOutputError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlOutputErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlOutputErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlOutputError) MarshalJSON() ([]byte, error) {
	pb, err := sqlOutputErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlQueryOutput struct {

	// Wire name: 'endpoint_id'
	EndpointId string
	// The link to find the output results.
	// Wire name: 'output_link'
	OutputLink string
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	// Wire name: 'query_text'
	QueryText string
	// Information about SQL statements executed in the run.
	// Wire name: 'sql_statements'
	SqlStatements []SqlStatementOutput
	// The canonical identifier of the SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func (st *SqlQueryOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlQueryOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlQueryOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlQueryOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlQueryOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	// Wire name: 'lookup_key'
	LookupKey string

	ForceSendFields []string `tf:"-"`
}

func (st *SqlStatementOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlStatementOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlStatementOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlStatementOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlStatementOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	// Wire name: 'alert'
	Alert *SqlTaskAlert
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	// Wire name: 'dashboard'
	Dashboard *SqlTaskDashboard
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	// Wire name: 'file'
	File *SqlTaskFile
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	// Wire name: 'parameters'
	Parameters map[string]string
	// If query, indicates that this job must execute a SQL query.
	// Wire name: 'query'
	Query *SqlTaskQuery
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	// Wire name: 'warehouse_id'
	WarehouseId string
}

func (st *SqlTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTask) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	// Wire name: 'alert_id'
	AlertId string
	// If true, the alert notifications are not sent to subscribers.
	// Wire name: 'pause_subscriptions'
	PauseSubscriptions bool
	// If specified, alert notifications are sent to subscribers.
	// Wire name: 'subscriptions'
	Subscriptions []SqlTaskSubscription

	ForceSendFields []string `tf:"-"`
}

func (st *SqlTaskAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskAlert) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	// Wire name: 'custom_subject'
	CustomSubject string
	// The canonical identifier of the SQL dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	// Wire name: 'pause_subscriptions'
	PauseSubscriptions bool
	// If specified, dashboard snapshots are sent to subscriptions.
	// Wire name: 'subscriptions'
	Subscriptions []SqlTaskSubscription

	ForceSendFields []string `tf:"-"`
}

func (st *SqlTaskDashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskDashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskDashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskDashboard) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskDashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	// Wire name: 'path'
	Path string
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	// Wire name: 'source'
	Source Source
}

func (st *SqlTaskFile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskFilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskFileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskFile) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskFileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	// Wire name: 'query_id'
	QueryId string
}

func (st *SqlTaskQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskQuery) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	// Wire name: 'destination_id'
	DestinationId string
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *SqlTaskSubscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskSubscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskSubscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskSubscription) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskSubscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns StorageMode to satisfy [pflag.Value] interface
func (f *StorageMode) Type() string {
	return "StorageMode"
}

type SubmitRun struct {
	// List of permissions to set on the job.
	// Wire name: 'access_control_list'
	AccessControlList []JobAccessControlRequest
	// The user specified id of the budget policy to use for this one-time run.
	// If not specified, the run will be not be attributed to any budget policy.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string
	// An optional set of email addresses notified when the run begins or
	// completes.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	// Wire name: 'environments'
	Environments []JobEnvironment
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
	GitSource *GitSource
	// An optional set of health rules that can be defined for this job.
	// Wire name: 'health'
	Health *JobsHealthRules
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
	IdempotencyToken string
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	// Wire name: 'notification_settings'
	NotificationSettings *JobNotificationSettings
	// The queue settings of the one-time run.
	// Wire name: 'queue'
	Queue *QueueSettings
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	// Wire name: 'run_as'
	RunAs *JobRunAs
	// An optional name for the run. The default value is `Untitled`.
	// Wire name: 'run_name'
	RunName string

	// Wire name: 'tasks'
	Tasks []SubmitTask
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string `tf:"-"`
}

func (st *SubmitRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &submitRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := submitRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubmitRun) MarshalJSON() ([]byte, error) {
	pb, err := submitRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Run was created and started successfully.
type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	// Wire name: 'run_id'
	RunId int64

	ForceSendFields []string `tf:"-"`
}

func (st *SubmitRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &submitRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := submitRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubmitRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := submitRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SubmitTask struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	// Wire name: 'clean_rooms_notebook_task'
	CleanRoomsNotebookTask *CleanRoomsNotebookTask
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	// Wire name: 'condition_task'
	ConditionTask *ConditionTask
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	// Wire name: 'dashboard_task'
	DashboardTask *DashboardTask
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	// Wire name: 'dbt_task'
	DbtTask *DbtTask
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	// Wire name: 'depends_on'
	DependsOn []TaskDependency
	// An optional description for this task.
	// Wire name: 'description'
	Description string
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	// Wire name: 'email_notifications'
	EmailNotifications *JobEmailNotifications
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	// Wire name: 'environment_key'
	EnvironmentKey string
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	// Wire name: 'for_each_task'
	ForEachTask *ForEachTask

	// Wire name: 'gen_ai_compute_task'
	GenAiComputeTask *GenAiComputeTask
	// An optional set of health rules that can be defined for this job.
	// Wire name: 'health'
	Health *JobsHealthRules
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec
	// The task runs a notebook when the `notebook_task` field is present.
	// Wire name: 'notebook_task'
	NotebookTask *NotebookTask
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	// Wire name: 'notification_settings'
	NotificationSettings *TaskNotificationSettings
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	// Wire name: 'pipeline_task'
	PipelineTask *PipelineTask
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	// Wire name: 'power_bi_task'
	PowerBiTask *PowerBiTask
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	// Wire name: 'python_wheel_task'
	PythonWheelTask *PythonWheelTask
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	// Wire name: 'run_if'
	RunIf RunIf
	// The task triggers another job when the `run_job_task` field is present.
	// Wire name: 'run_job_task'
	RunJobTask *RunJobTask
	// The task runs a JAR when the `spark_jar_task` field is present.
	// Wire name: 'spark_jar_task'
	SparkJarTask *SparkJarTask
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	// Wire name: 'spark_python_task'
	SparkPythonTask *SparkPythonTask
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
	SparkSubmitTask *SparkSubmitTask
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	// Wire name: 'sql_task'
	SqlTask *SqlTask
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	// Wire name: 'task_key'
	TaskKey string
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string `tf:"-"`
}

func (st *SubmitTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &submitTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := submitTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubmitTask) MarshalJSON() ([]byte, error) {
	pb, err := submitTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Subscription struct {
	// Optional: Allows users to specify a custom subject line on the email sent
	// to subscribers.
	// Wire name: 'custom_subject'
	CustomSubject string
	// When true, the subscription will not send emails.
	// Wire name: 'paused'
	Paused bool
	// The list of subscribers to send the snapshot of the dashboard to.
	// Wire name: 'subscribers'
	Subscribers []SubscriptionSubscriber

	ForceSendFields []string `tf:"-"`
}

func (st *Subscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Subscription) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SubscriptionSubscriber struct {
	// A snapshot of the dashboard will be sent to the destination when the
	// `destination_id` field is present.
	// Wire name: 'destination_id'
	DestinationId string
	// A snapshot of the dashboard will be sent to the user's email when the
	// `user_name` field is present.
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *SubscriptionSubscriber) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionSubscriberPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionSubscriberFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubscriptionSubscriber) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionSubscriberToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TableUpdateTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	// Wire name: 'condition'
	Condition Condition
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	// Wire name: 'min_time_between_triggers_seconds'
	MinTimeBetweenTriggersSeconds int
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	// Wire name: 'table_names'
	TableNames []string
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	// Wire name: 'wait_after_last_change_seconds'
	WaitAfterLastChangeSeconds int

	ForceSendFields []string `tf:"-"`
}

func (st *TableUpdateTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableUpdateTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableUpdateTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableUpdateTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := tableUpdateTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Task struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	// Wire name: 'clean_rooms_notebook_task'
	CleanRoomsNotebookTask *CleanRoomsNotebookTask
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	// Wire name: 'condition_task'
	ConditionTask *ConditionTask
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	// Wire name: 'dashboard_task'
	DashboardTask *DashboardTask
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	// Wire name: 'dbt_task'
	DbtTask *DbtTask
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	// Wire name: 'depends_on'
	DependsOn []TaskDependency
	// An optional description for this task.
	// Wire name: 'description'
	Description string
	// An option to disable auto optimization in serverless
	// Wire name: 'disable_auto_optimization'
	DisableAutoOptimization bool
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	// Wire name: 'email_notifications'
	EmailNotifications *TaskEmailNotifications
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	// Wire name: 'environment_key'
	EnvironmentKey string
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	// Wire name: 'existing_cluster_id'
	ExistingClusterId string
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	// Wire name: 'for_each_task'
	ForEachTask *ForEachTask

	// Wire name: 'gen_ai_compute_task'
	GenAiComputeTask *GenAiComputeTask
	// An optional set of health rules that can be defined for this job.
	// Wire name: 'health'
	Health *JobsHealthRules
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	// Wire name: 'job_cluster_key'
	JobClusterKey string
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	// Wire name: 'libraries'
	Libraries []compute.Library
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	// Wire name: 'max_retries'
	MaxRetries int
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	// Wire name: 'min_retry_interval_millis'
	MinRetryIntervalMillis int
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	// Wire name: 'new_cluster'
	NewCluster *compute.ClusterSpec
	// The task runs a notebook when the `notebook_task` field is present.
	// Wire name: 'notebook_task'
	NotebookTask *NotebookTask
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	// Wire name: 'notification_settings'
	NotificationSettings *TaskNotificationSettings
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	// Wire name: 'pipeline_task'
	PipelineTask *PipelineTask
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	// Wire name: 'power_bi_task'
	PowerBiTask *PowerBiTask
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	// Wire name: 'python_wheel_task'
	PythonWheelTask *PythonWheelTask
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	// Wire name: 'retry_on_timeout'
	RetryOnTimeout bool
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
	RunIf RunIf
	// The task triggers another job when the `run_job_task` field is present.
	// Wire name: 'run_job_task'
	RunJobTask *RunJobTask
	// The task runs a JAR when the `spark_jar_task` field is present.
	// Wire name: 'spark_jar_task'
	SparkJarTask *SparkJarTask
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	// Wire name: 'spark_python_task'
	SparkPythonTask *SparkPythonTask
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
	SparkSubmitTask *SparkSubmitTask
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	// Wire name: 'sql_task'
	SqlTask *SqlTask
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	// Wire name: 'task_key'
	TaskKey string
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	// Wire name: 'timeout_seconds'
	TimeoutSeconds int
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	// Wire name: 'webhook_notifications'
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string `tf:"-"`
}

func (st *Task) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Task) MarshalJSON() ([]byte, error) {
	pb, err := taskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	// Wire name: 'outcome'
	Outcome string
	// The name of the task this task depends on.
	// Wire name: 'task_key'
	TaskKey string

	ForceSendFields []string `tf:"-"`
}

func (st *TaskDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskDependency) MarshalJSON() ([]byte, error) {
	pb, err := taskDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TaskEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	// Wire name: 'on_duration_warning_threshold_exceeded'
	OnDurationWarningThresholdExceeded []string
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	// Wire name: 'on_failure'
	OnFailure []string
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_start'
	OnStart []string
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	// Wire name: 'on_streaming_backlog_exceeded'
	OnStreamingBacklogExceeded []string
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	// Wire name: 'on_success'
	OnSuccess []string

	ForceSendFields []string `tf:"-"`
}

func (st *TaskEmailNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskEmailNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskEmailNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskEmailNotifications) MarshalJSON() ([]byte, error) {
	pb, err := taskEmailNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	// Wire name: 'alert_on_last_attempt'
	AlertOnLastAttempt bool
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	// Wire name: 'no_alert_for_canceled_runs'
	NoAlertForCanceledRuns bool
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	// Wire name: 'no_alert_for_skipped_runs'
	NoAlertForSkippedRuns bool

	ForceSendFields []string `tf:"-"`
}

func (st *TaskNotificationSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskNotificationSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskNotificationSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskNotificationSettings) MarshalJSON() ([]byte, error) {
	pb, err := taskNotificationSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns TerminationCodeCode to satisfy [pflag.Value] interface
func (f *TerminationCodeCode) Type() string {
	return "TerminationCodeCode"
}

type TerminationDetails struct {
	// The code indicates why the run was terminated. Additional codes might be
	// introduced in future releases. * `SUCCESS`: The run was completed
	// successfully. * `SUCCESS_WITH_FAILURES`: The run was completed
	// successfully but some child runs failed. * `USER_CANCELED`: The run was
	// successfully canceled during execution by a user. * `CANCELED`: The run
	// was canceled during execution by the Databricks platform; for example, if
	// the maximum run duration was exceeded. * `SKIPPED`: Run was never
	// executed, for example, if the upstream task run failed, the dependency
	// type condition was not met, or there were no material tasks to execute. *
	// `INTERNAL_ERROR`: The run encountered an unexpected error. Refer to the
	// state message for further details. * `DRIVER_ERROR`: The run encountered
	// an error while communicating with the Spark Driver. * `CLUSTER_ERROR`:
	// The run failed due to a cluster error. Refer to the state message for
	// further details. * `REPOSITORY_CHECKOUT_FAILED`: Failed to complete the
	// checkout due to an error when communicating with the third party service.
	// * `INVALID_CLUSTER_REQUEST`: The run failed because it issued an invalid
	// request to start the cluster. * `WORKSPACE_RUN_LIMIT_EXCEEDED`: The
	// workspace has reached the quota for the maximum number of concurrent
	// active runs. Consider scheduling the runs over a larger time frame. *
	// `FEATURE_DISABLED`: The run failed because it tried to access a feature
	// unavailable for the workspace. * `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The
	// number of cluster creation, start, and upsize requests have exceeded the
	// allotted rate limit. Consider spreading the run execution over a larger
	// time frame. * `STORAGE_ACCESS_ERROR`: The run failed due to an error when
	// accessing the customer blob storage. Refer to the state message for
	// further details. * `RUN_EXECUTION_ERROR`: The run was completed with task
	// failures. For more details, refer to the state message or run output. *
	// `UNAUTHORIZED_ERROR`: The run failed due to a permission issue while
	// accessing a resource. Refer to the state message for further details. *
	// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
	// user-requested library. Refer to the state message for further details.
	// The causes might include, but are not limited to: The provided library is
	// invalid, there are insufficient permissions to install the library, and
	// so forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the
	// limit of maximum concurrent runs set for the job. *
	// `MAX_SPARK_CONTEXTS_EXCEEDED`: The run is scheduled on a cluster that has
	// already reached the maximum number of contexts it is configured to
	// create. See: [Link]. * `RESOURCE_NOT_FOUND`: A resource necessary for run
	// execution does not exist. Refer to the state message for further details.
	// * `INVALID_RUN_CONFIGURATION`: The run failed due to an invalid
	// configuration. Refer to the state message for further details. *
	// `CLOUD_FAILURE`: The run failed due to a cloud provider issue. Refer to
	// the state message for further details. * `MAX_JOB_QUEUE_SIZE_EXCEEDED`:
	// The run was skipped due to reaching the job level queue size limit. *
	// `DISABLED`: The run was never executed because it was disabled explicitly
	// by the user.
	//
	// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
	// Wire name: 'code'
	Code TerminationCodeCode
	// A descriptive message with the termination details. This field is
	// unstructured and the format might change.
	// Wire name: 'message'
	Message string
	// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
	// error occurred in the Databricks platform. Please look at the [status
	// page] or contact support if the issue persists. * `CLIENT_ERROR`: The run
	// was terminated because of an error caused by user input or the job
	// configuration. * `CLOUD_FAILURE`: The run was terminated because of an
	// issue with your cloud provider.
	//
	// [status page]: https://status.databricks.com/
	// Wire name: 'type'
	Type TerminationTypeType

	ForceSendFields []string `tf:"-"`
}

func (st *TerminationDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &terminationDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := terminationDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TerminationDetails) MarshalJSON() ([]byte, error) {
	pb, err := terminationDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns TerminationTypeType to satisfy [pflag.Value] interface
func (f *TerminationTypeType) Type() string {
	return "TerminationTypeType"
}

// Additional details about what triggered the run
type TriggerInfo struct {
	// The run id of the Run Job task run
	// Wire name: 'run_id'
	RunId int64

	ForceSendFields []string `tf:"-"`
}

func (st *TriggerInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &triggerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := triggerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TriggerInfo) MarshalJSON() ([]byte, error) {
	pb, err := triggerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TriggerSettings struct {
	// File arrival trigger settings.
	// Wire name: 'file_arrival'
	FileArrival *FileArrivalTriggerConfiguration
	// Whether this trigger is paused or not.
	// Wire name: 'pause_status'
	PauseStatus PauseStatus
	// Periodic trigger settings.
	// Wire name: 'periodic'
	Periodic *PeriodicTriggerConfiguration
	// Old table trigger settings name. Deprecated in favor of `table_update`.
	// Wire name: 'table'
	Table *TableUpdateTriggerConfiguration

	// Wire name: 'table_update'
	TableUpdate *TableUpdateTriggerConfiguration
}

func (st *TriggerSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &triggerSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := triggerSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TriggerSettings) MarshalJSON() ([]byte, error) {
	pb, err := triggerSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
// `TABLE`: Indicates a run that is triggered by a table update. *
// `CONTINUOUS_RESTART`: Indicates a run created by user to manually restart a
// continuous job run.
type TriggerType string

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
	case `FILE_ARRIVAL`, `ONE_TIME`, `PERIODIC`, `RETRY`, `RUN_JOB_TASK`, `TABLE`:
		*f = TriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_ARRIVAL", "ONE_TIME", "PERIODIC", "RETRY", "RUN_JOB_TASK", "TABLE"`, v)
	}
}

// Type always returns TriggerType to satisfy [pflag.Value] interface
func (f *TriggerType) Type() string {
	return "TriggerType"
}

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported, except for tasks and job clusters (`tasks/task_1`). This
	// field is optional.
	// Wire name: 'fields_to_remove'
	FieldsToRemove []string
	// The canonical identifier of the job to update. This field is required.
	// Wire name: 'job_id'
	JobId int64
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
	NewSettings *JobSettings
}

func (st *UpdateJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateJob) MarshalJSON() ([]byte, error) {
	pb, err := updateJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateResponse struct {
}

func (st *UpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ViewItem struct {
	// Content of the view.
	// Wire name: 'content'
	Content string
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	// Wire name: 'name'
	Name string
	// Type of the view item.
	// Wire name: 'type'
	Type ViewType

	ForceSendFields []string `tf:"-"`
}

func (st *ViewItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &viewItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := viewItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ViewItem) MarshalJSON() ([]byte, error) {
	pb, err := viewItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
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

// Type always returns ViewsToExport to satisfy [pflag.Value] interface
func (f *ViewsToExport) Type() string {
	return "ViewsToExport"
}

type Webhook struct {

	// Wire name: 'id'
	Id string
}

func (st *Webhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &webhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := webhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Webhook) MarshalJSON() ([]byte, error) {
	pb, err := webhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	// Wire name: 'on_duration_warning_threshold_exceeded'
	OnDurationWarningThresholdExceeded []Webhook
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	// Wire name: 'on_failure'
	OnFailure []Webhook
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	// Wire name: 'on_start'
	OnStart []Webhook
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	// Wire name: 'on_streaming_backlog_exceeded'
	OnStreamingBacklogExceeded []Webhook
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	// Wire name: 'on_success'
	OnSuccess []Webhook
}

func (st *WebhookNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &webhookNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := webhookNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WebhookNotifications) MarshalJSON() ([]byte, error) {
	pb, err := webhookNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WidgetErrorDetail struct {

	// Wire name: 'message'
	Message string

	ForceSendFields []string `tf:"-"`
}

func (st *WidgetErrorDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &widgetErrorDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := widgetErrorDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WidgetErrorDetail) MarshalJSON() ([]byte, error) {
	pb, err := widgetErrorDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
