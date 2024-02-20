// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/iam"
)

// all definitions in this file are in alphabetical order

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `json:"settings,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *BaseJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseJob) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BaseRun struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \> 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
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
	GitSource *GitSource `json:"git_source,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
	// Job-level parameters used in the run
	JobParameters []JobParameter `json:"job_parameters,omitempty"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64 `json:"original_attempt_run_id,omitempty"`
	// The parameters used for this run.
	OverridingParameters *RunParameters `json:"overriding_parameters,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int `json:"run_duration,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`
	// * `JOB_RUN`: Normal job run. A run created with :method:jobs/runNow. *
	// `WORKFLOW_RUN`: Workflow run. A run created with [dbutils.notebook.run].
	// * `SUBMIT_RUN`: Submit run. A run created with :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// The current state of the run.
	State *RunState `json:"state,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks []RunTask `json:"tasks,omitempty"`
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
	Trigger TriggerType `json:"trigger,omitempty"`

	TriggerInfo *TriggerInfo `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *BaseRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns bool `json:"all_queued_runs,omitempty"`
	// The canonical identifier of the job to cancel all runs of.
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CancelAllRuns) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CancelAllRuns) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CancelRun struct {
	// This field is required.
	RunId int64 `json:"run_id"`
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
	ClusterId string `json:"cluster_id,omitempty"`
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	SparkContextId string `json:"spark_context_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ClusterInstance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterInstance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this job. When running jobs on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the job. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ClusterSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Left string `json:"left,omitempty"`
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
	Op ConditionTaskOp `json:"op,omitempty"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string `json:"right,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ConditionTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ConditionTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
}

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList []iam.AccessControlRequest `json:"access_control_list,omitempty"`
	// A list of compute requirements that can be referenced by tasks of this
	// job.
	Compute []JobCompute `json:"compute,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous `json:"continuous,omitempty"`
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment `json:"deployment,omitempty"`
	// An optional description for the job. The maximum length is 1024
	// characters in UTF-8 encoding.
	Description string `json:"description,omitempty"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode CreateJobEditMode `json:"edit_mode,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `json:"format,omitempty"`
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
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job.
	//
	// Set this value if you want to be able to execute multiple runs of the
	// same job concurrently. This is useful for example if you trigger your job
	// on a frequent schedule and want to allow consecutive runs to overlap with
	// each other, or if you want to trigger multiple runs which differ by their
	// input parameters.
	//
	// This setting affects only new runs. For example, suppose the job’s
	// concurrency is 4 and there are 4 concurrent active runs. Then setting the
	// concurrency to 3 won’t kill any of the active runs. However, from then
	// on, new runs are skipped unless there are fewer than 3 active runs.
	//
	// This value cannot exceed 1000. Setting this value to `0` causes all new
	// runs to be skipped.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// Job-level parameter definitions
	Parameters []JobParameterDefinition `json:"parameters,omitempty"`
	// The queue settings of the job.
	Queue *QueueSettings `json:"queue,omitempty"`
	// Write-only setting, available only in Create/Update/Reset and Submit
	// calls. Specifies the user or service principal that the job runs as. If
	// not specified, the job runs as the user who created the job.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *JobRunAs `json:"run_as,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job.
	Tasks []Task `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings `json:"trigger,omitempty"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateJob) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateJob) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Edit mode of the job.
//
// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
// `EDITABLE`: The job is in an editable state and can be modified.
type CreateJobEditMode string

// The job is in an editable state and can be modified.
const CreateJobEditModeEditable CreateJobEditMode = `EDITABLE`

// The job is in a locked UI state and cannot be modified.
const CreateJobEditModeUiLocked CreateJobEditMode = `UI_LOCKED`

// String representation for [fmt.Print]
func (f *CreateJobEditMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateJobEditMode) Set(v string) error {
	switch v {
	case `EDITABLE`, `UI_LOCKED`:
		*f = CreateJobEditMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EDITABLE", "UI_LOCKED"`, v)
	}
}

// Type always returns CreateJobEditMode to satisfy [pflag.Value] interface
func (f *CreateJobEditMode) Type() string {
	return "CreateJobEditMode"
}

type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required."
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders map[string]string `json:"artifacts_headers,omitempty"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink string `json:"artifacts_link,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DbtOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	Catalog string `json:"catalog,omitempty"`
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	Commands []string `json:"commands"`
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	ProfilesDirectory string `json:"profiles_directory,omitempty"`
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	ProjectDirectory string `json:"project_directory,omitempty"`
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema string `json:"schema,omitempty"`
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local <Databricks> workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in <Databricks> workspace. * `GIT`:
	// Project is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DbtTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DbtTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId int64 `json:"job_id"`
}

type DeleteRun struct {
	// The canonical identifier of the run for which to retrieve the metadata.
	RunId int64 `json:"run_id"`
}

type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	Views []ViewItem `json:"views,omitempty"`
}

// Export and retrieve a job run
type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 `json:"-" url:"run_id"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport `json:"-" url:"views_to_export,omitempty"`
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// The storage location to monitor for file arrivals. The value must point
	// to the root or a subpath of an external location URL or the root or
	// subpath of a Unity Catalog volume.
	Url string `json:"url,omitempty"`
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *FileArrivalTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileArrivalTriggerConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats *ForEachTaskErrorMessageStats `json:"error_message_stats,omitempty"`
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats *ForEachTaskTaskRunStats `json:"task_run_stats,omitempty"`
}

type ForEachTask struct {
	// Controls the number of active iterations task runs. Default is 20,
	// maximum allowed is 100.
	Concurrency int `json:"concurrency,omitempty"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string `json:"inputs"`

	Task Task `json:"task"`

	ForceSendFields []string `json:"-"`
}

func (s *ForEachTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForEachTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count string `json:"count,omitempty"`
	// Describes the error message occured during the iterations.
	ErrorMessage string `json:"error_message,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ForEachTaskErrorMessageStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForEachTaskErrorMessageStats) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	ActiveIterations int `json:"active_iterations,omitempty"`
	// Describes the number of failed and succeeded iteration runs.
	CompletedIterations int `json:"completed_iterations,omitempty"`
	// Describes the number of failed iteration runs.
	FailedIterations int `json:"failed_iterations,omitempty"`
	// Describes the number of iteration runs that have been scheduled.
	ScheduledIterations int `json:"scheduled_iterations,omitempty"`
	// Describes the number of succeeded iteration runs.
	SucceededIterations int `json:"succeeded_iterations,omitempty"`
	// Describes the length of the list of items to iterate over.
	TotalIterations int `json:"total_iterations,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ForEachTaskTaskRunStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ForEachTaskTaskRunStats) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Get job permission levels
type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []JobPermissionsDescription `json:"permission_levels,omitempty"`
}

// Get job permissions
type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

// Get a single job
type GetJobRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId int64 `json:"-" url:"job_id"`
}

// Get the output for a single run
type GetRunOutputRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 `json:"-" url:"run_id"`
}

// Get a single job run
type GetRunRequest struct {
	// Whether to include the repair history in the response.
	IncludeHistory bool `json:"-" url:"include_history,omitempty"`
	// Whether to include resolved parameter values in the response.
	IncludeResolvedValues bool `json:"-" url:"include_resolved_values,omitempty"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId int64 `json:"-" url:"run_id"`

	ForceSendFields []string `json:"-"`
}

func (s *GetRunRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRunRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	UsedCommit string `json:"used_commit,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GitSnapshot) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GitSnapshot) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	GitBranch string `json:"git_branch,omitempty"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	GitCommit string `json:"git_commit,omitempty"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitProvider `json:"git_provider"`
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot *GitSnapshot `json:"git_snapshot,omitempty"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	GitTag string `json:"git_tag,omitempty"`
	// URL of the repository to be cloned by this job.
	GitUrl string `json:"git_url"`
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	JobSource *JobSource `json:"job_source,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GitSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GitSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `json:"settings,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Job) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Job) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobAccessControlResponse struct {
	// All permissions.
	AllPermissions []JobPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey string `json:"job_cluster_key"`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
}

type JobCompute struct {
	// A unique name for the compute requirement. This field is required and
	// must be unique within the job. `JobTaskSettings` may refer to this field
	// to determine the compute requirements for the task execution.
	ComputeKey string `json:"compute_key"`

	Spec compute.ComputeSpec `json:"spec"`
}

type JobDeployment struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	Kind JobDeploymentKind `json:"kind"`
	// Path of the file that contains deployment metadata.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobDeployment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobDeployment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The kind of deployment that manages the job.
//
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

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobEmailNotifications) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobEmailNotifications) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobNotificationSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobNotificationSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobParameter struct {
	// The optional default value of the parameter
	Default string `json:"default,omitempty"`
	// The name of the parameter
	Name string `json:"name,omitempty"`
	// The value used in the run
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	Default string `json:"default"`
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name string `json:"name"`
}

type JobPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	AccessControlList []JobAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobPermissionsRequest struct {
	AccessControlList []JobAccessControlRequest `json:"access_control_list,omitempty"`
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

// Write-only setting, available only in Create/Update/Reset and Submit calls.
// Specifies the user or service principal that the job runs as. If not
// specified, the job runs as the user who created the job.
//
// Only `user_name` or `service_principal_name` can be specified. If both are
// specified, an error is thrown.
type JobRunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobRunAs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobRunAs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type JobSettings struct {
	// A list of compute requirements that can be referenced by tasks of this
	// job.
	Compute []JobCompute `json:"compute,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous `json:"continuous,omitempty"`
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment `json:"deployment,omitempty"`
	// An optional description for the job. The maximum length is 1024
	// characters in UTF-8 encoding.
	Description string `json:"description,omitempty"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobSettingsEditMode `json:"edit_mode,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `json:"format,omitempty"`
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
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job.
	//
	// Set this value if you want to be able to execute multiple runs of the
	// same job concurrently. This is useful for example if you trigger your job
	// on a frequent schedule and want to allow consecutive runs to overlap with
	// each other, or if you want to trigger multiple runs which differ by their
	// input parameters.
	//
	// This setting affects only new runs. For example, suppose the job’s
	// concurrency is 4 and there are 4 concurrent active runs. Then setting the
	// concurrency to 3 won’t kill any of the active runs. However, from then
	// on, new runs are skipped unless there are fewer than 3 active runs.
	//
	// This value cannot exceed 1000. Setting this value to `0` causes all new
	// runs to be skipped.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// Job-level parameter definitions
	Parameters []JobParameterDefinition `json:"parameters,omitempty"`
	// The queue settings of the job.
	Queue *QueueSettings `json:"queue,omitempty"`
	// Write-only setting, available only in Create/Update/Reset and Submit
	// calls. Specifies the user or service principal that the job runs as. If
	// not specified, the job runs as the user who created the job.
	//
	// Only `user_name` or `service_principal_name` can be specified. If both
	// are specified, an error is thrown.
	RunAs *JobRunAs `json:"run_as,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job.
	Tasks []Task `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings `json:"trigger,omitempty"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Edit mode of the job.
//
// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
// `EDITABLE`: The job is in an editable state and can be modified.
type JobSettingsEditMode string

// The job is in an editable state and can be modified.
const JobSettingsEditModeEditable JobSettingsEditMode = `EDITABLE`

// The job is in a locked UI state and cannot be modified.
const JobSettingsEditModeUiLocked JobSettingsEditMode = `UI_LOCKED`

// String representation for [fmt.Print]
func (f *JobSettingsEditMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobSettingsEditMode) Set(v string) error {
	switch v {
	case `EDITABLE`, `UI_LOCKED`:
		*f = JobSettingsEditMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EDITABLE", "UI_LOCKED"`, v)
	}
}

// Type always returns JobSettingsEditMode to satisfy [pflag.Value] interface
func (f *JobSettingsEditMode) Type() string {
	return "JobSettingsEditMode"
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
	DirtyState JobSourceDirtyState `json:"dirty_state,omitempty"`
	// Name of the branch which the job is imported from.
	ImportFromGitBranch string `json:"import_from_git_branch"`
	// Path of the job YAML file that contains the job specification.
	JobConfigPath string `json:"job_config_path"`
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
type JobsHealthMetric string

const JobsHealthMetricRunDurationSeconds JobsHealthMetric = `RUN_DURATION_SECONDS`

// String representation for [fmt.Print]
func (f *JobsHealthMetric) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthMetric) Set(v string) error {
	switch v {
	case `RUN_DURATION_SECONDS`:
		*f = JobsHealthMetric(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RUN_DURATION_SECONDS"`, v)
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
	Metric JobsHealthMetric `json:"metric,omitempty"`
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op JobsHealthOperator `json:"op,omitempty"`
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value int64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *JobsHealthRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s JobsHealthRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {
	Rules []JobsHealthRule `json:"rules,omitempty"`
}

// List jobs
type ListJobsRequest struct {
	// Whether to include task and cluster details in the response.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	Limit int `json:"-" url:"limit,omitempty"`
	// A filter on the list based on the exact (case insensitive) job name.
	Name string `json:"-" url:"name,omitempty"`
	// The offset of the first job to return, relative to the most recently
	// created job.
	//
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset int `json:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListJobsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// The list of jobs.
	Jobs []BaseJob `json:"jobs,omitempty"`
	// A token that can be used to list the next page of jobs.
	NextPageToken string `json:"next_page_token,omitempty"`
	// A token that can be used to list the previous page of jobs.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListJobsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List job runs
type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	ActiveOnly bool `json:"-" url:"active_only,omitempty"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly bool `json:"-" url:"completed_only,omitempty"`
	// Whether to include task and cluster details in the response.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId int64 `json:"-" url:"job_id,omitempty"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit int `json:"-" url:"limit,omitempty"`
	// The offset of the first run to return, relative to the most recent run.
	//
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset int `json:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	RunType ListRunsRunType `json:"-" url:"run_type,omitempty"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom int `json:"-" url:"start_time_from,omitempty"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo int `json:"-" url:"start_time_to,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListRunsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRunsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// A token that can be used to list the next page of runs.
	NextPageToken string `json:"next_page_token,omitempty"`
	// A token that can be used to list the previous page of runs.
	PrevPageToken string `json:"prev_page_token,omitempty"`
	// A list of runs, from most recently started to least.
	Runs []BaseRun `json:"runs,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListRunsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListRunsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * `JOB_RUN`: Normal job run. A run created with :method:jobs/runNow. *
// `WORKFLOW_RUN`: Workflow run. A run created with [dbutils.notebook.run]. *
// `SUBMIT_RUN`: Submit run. A run created with :method:jobs/submit.
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
type ListRunsRunType string

// Normal job run. A run created with :method:jobs/runNow.
const ListRunsRunTypeJobRun ListRunsRunType = `JOB_RUN`

// Submit run. A run created with :method:jobs/submit.
const ListRunsRunTypeSubmitRun ListRunsRunType = `SUBMIT_RUN`

// Workflow run. A run created with
// [dbutils.notebook.run](/dev-tools/databricks-utils.html#dbutils-workflow).
const ListRunsRunTypeWorkflowRun ListRunsRunType = `WORKFLOW_RUN`

// String representation for [fmt.Print]
func (f *ListRunsRunType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListRunsRunType) Set(v string) error {
	switch v {
	case `JOB_RUN`, `SUBMIT_RUN`, `WORKFLOW_RUN`:
		*f = ListRunsRunType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "JOB_RUN", "SUBMIT_RUN", "WORKFLOW_RUN"`, v)
	}
}

// Type always returns ListRunsRunType to satisfy [pflag.Value] interface
func (f *ListRunsRunType) Type() string {
	return "ListRunsRunType"
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	Result string `json:"result,omitempty"`
	// Whether or not the result was truncated.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *NotebookOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotebookOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NotebookTask struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to :method:jobs/runNow with parameters specified, the
	// two parameters maps are merged. If the same key is specified in
	// `base_parameters` and in `run-now`, the value from `run-now` is used.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
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
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-widgets
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath string `json:"notebook_path"`
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local <Databricks> workspace. When
	// set to `GIT`, the notebook will be retrieved from a Git repository
	// defined in `git_source`. If the value is empty, the task will use `GIT`
	// if `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Notebook is located in <Databricks> workspace. * `GIT`:
	// Notebook is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
}

type ParamPairs map[string]string

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

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PipelineParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PipelineTask struct {
	// If true, a full refresh will be triggered on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// The full name of the pipeline task to execute.
	PipelineId string `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PipelineTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PipelineTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint string `json:"entry_point,omitempty"`
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	// Name of the package to execute
	PackageName string `json:"package_name,omitempty"`
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters []string `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PythonWheelTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PythonWheelTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled bool `json:"enabled"`
}

type RepairHistoryItem struct {
	// The end time of the (repaired) run.
	EndTime int64 `json:"end_time,omitempty"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id int64 `json:"id,omitempty"`
	// The start time of the (repaired) run.
	StartTime int64 `json:"start_time,omitempty"`
	// The current state of the run.
	State *RunState `json:"state,omitempty"`
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds []int64 `json:"task_run_ids,omitempty"`
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type RepairHistoryItemType `json:"type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepairHistoryItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepairHistoryItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	JarParams []string `json:"jar_params,omitempty"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId int64 `json:"latest_repair_id,omitempty"`
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
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `"python_named_params": {"name": "task", "data":
	// "dbfs:/path/to/data.json"}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	PythonParams []string `json:"python_params,omitempty"`
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	RerunAllFailedTasks bool `json:"rerun_all_failed_tasks,omitempty"`
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	RerunDependentTasks bool `json:"rerun_dependent_tasks,omitempty"`
	// The task keys of the task runs to repair.
	RerunTasks []string `json:"rerun_tasks,omitempty"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId int64 `json:"run_id"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepairRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepairRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId int64 `json:"repair_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RepairRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RepairRunResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings JobSettings `json:"new_settings"`
}

type ResolvedConditionTaskValues struct {
	Left string `json:"left,omitempty"`

	Right string `json:"right,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ResolvedConditionTaskValues) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResolvedConditionTaskValues) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResolvedDbtTaskValues struct {
	Commands []string `json:"commands,omitempty"`
}

type ResolvedNotebookTaskValues struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

type ResolvedParamPairValues struct {
	Parameters map[string]string `json:"parameters,omitempty"`
}

type ResolvedPythonWheelTaskValues struct {
	NamedParameters map[string]string `json:"named_parameters,omitempty"`

	Parameters []string `json:"parameters,omitempty"`
}

type ResolvedRunJobTaskValues struct {
	NamedParameters map[string]string `json:"named_parameters,omitempty"`

	Parameters map[string]string `json:"parameters,omitempty"`
}

type ResolvedStringParamsValues struct {
	Parameters []string `json:"parameters,omitempty"`
}

type ResolvedValues struct {
	ConditionTask *ResolvedConditionTaskValues `json:"condition_task,omitempty"`

	DbtTask *ResolvedDbtTaskValues `json:"dbt_task,omitempty"`

	NotebookTask *ResolvedNotebookTaskValues `json:"notebook_task,omitempty"`

	PythonWheelTask *ResolvedPythonWheelTaskValues `json:"python_wheel_task,omitempty"`

	RunJobTask *ResolvedRunJobTaskValues `json:"run_job_task,omitempty"`

	SimulationTask *ResolvedParamPairValues `json:"simulation_task,omitempty"`

	SparkJarTask *ResolvedStringParamsValues `json:"spark_jar_task,omitempty"`

	SparkPythonTask *ResolvedStringParamsValues `json:"spark_python_task,omitempty"`

	SparkSubmitTask *ResolvedStringParamsValues `json:"spark_submit_task,omitempty"`

	SqlTask *ResolvedParamPairValues `json:"sql_task,omitempty"`
}

type Run struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \> 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
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
	GitSource *GitSource `json:"git_source,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
	// Job-level parameters used in the run
	JobParameters []JobParameter `json:"job_parameters,omitempty"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64 `json:"original_attempt_run_id,omitempty"`
	// The parameters used for this run.
	OverridingParameters *RunParameters `json:"overriding_parameters,omitempty"`
	// The repair history of the run.
	RepairHistory []RepairHistoryItem `json:"repair_history,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int `json:"run_duration,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`
	// * `JOB_RUN`: Normal job run. A run created with :method:jobs/runNow. *
	// `WORKFLOW_RUN`: Workflow run. A run created with [dbutils.notebook.run].
	// * `SUBMIT_RUN`: Submit run. A run created with :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// The current state of the run.
	State *RunState `json:"state,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks []RunTask `json:"tasks,omitempty"`
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
	Trigger TriggerType `json:"trigger,omitempty"`

	TriggerInfo *TriggerInfo `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Run) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Run) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunConditionTask struct {
	// The left operand of the condition task.
	Left string `json:"left"`
	// The condtion task operator.
	Op RunConditionTaskOp `json:"op"`
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	Outcome string `json:"outcome,omitempty"`
	// The right operand of the condition task.
	Right string `json:"right"`

	ForceSendFields []string `json:"-"`
}

func (s *RunConditionTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunConditionTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The condtion task operator.
type RunConditionTaskOp string

const RunConditionTaskOpEqualTo RunConditionTaskOp = `EQUAL_TO`

const RunConditionTaskOpGreaterThan RunConditionTaskOp = `GREATER_THAN`

const RunConditionTaskOpGreaterThanOrEqual RunConditionTaskOp = `GREATER_THAN_OR_EQUAL`

const RunConditionTaskOpLessThan RunConditionTaskOp = `LESS_THAN`

const RunConditionTaskOpLessThanOrEqual RunConditionTaskOp = `LESS_THAN_OR_EQUAL`

const RunConditionTaskOpNotEqual RunConditionTaskOp = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *RunConditionTaskOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunConditionTaskOp) Set(v string) error {
	switch v {
	case `EQUAL_TO`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = RunConditionTaskOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL_TO", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Type always returns RunConditionTaskOp to satisfy [pflag.Value] interface
func (f *RunConditionTaskOp) Type() string {
	return "RunConditionTaskOp"
}

type RunForEachTask struct {
	// Controls the number of active iterations task runs. Default is 20,
	// maximum allowed is 100.
	Concurrency int `json:"concurrency,omitempty"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string `json:"inputs,omitempty"`

	Stats *ForEachStats `json:"stats,omitempty"`

	Task *Task `json:"task,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunForEachTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunForEachTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	RunId int `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunJobOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunJobOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunJobTask struct {
	// ID of the job to trigger.
	JobId int `json:"job_id"`
	// Job-level parameters used to trigger the job.
	JobParameters map[string]string `json:"job_parameters,omitempty"`
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

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
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
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the job to be executed
	JobId int64 `json:"job_id"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string `json:"job_parameters,omitempty"`
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
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `"python_named_params": {"name": "task", "data":
	// "dbfs:/path/to/data.json"}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	PythonParams []string `json:"python_params,omitempty"`
	// The queue settings of the run.
	Queue *QueueSettings `json:"queue,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunNow) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunNow) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// The globally unique ID of the newly triggered run.
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunNowResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunNowResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunOutput struct {
	// The output of a dbt task, if available.
	DbtOutput *DbtOutput `json:"dbt_output,omitempty"`
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error string `json:"error,omitempty"`
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace string `json:"error_trace,omitempty"`
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	Logs string `json:"logs,omitempty"`
	// Whether the logs are truncated.
	LogsTruncated bool `json:"logs_truncated,omitempty"`
	// All details of the run except for its output.
	Metadata *Run `json:"metadata,omitempty"`
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. <Databricks> restricts this
	// API to return the first 5 MB of the output. To return a larger result,
	// use the
	// [ClusterLogConf](/dev-tools/api/latest/clusters.html#clusterlogconf)
	// field to configure log storage for the job cluster.
	NotebookOutput *NotebookOutput `json:"notebook_output,omitempty"`
	// The output of a run job task, if available
	RunJobOutput *RunJobOutput `json:"run_job_output,omitempty"`
	// The output of a SQL task, if available.
	SqlOutput *SqlOutput `json:"sql_output,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	JarParams []string `json:"jar_params,omitempty"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string `json:"job_parameters,omitempty"`
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
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `"python_named_params": {"name": "task", "data":
	// "dbfs:/path/to/data.json"}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`
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
// canceled.
type RunResultState string

// The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

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
	case `CANCELED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_FAILED`:
		*f = RunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_FAILED"`, v)
	}
}

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response.
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty"`
	// The reason indicating why the run was queued.
	QueueReason string `json:"queue_reason,omitempty"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states.
	ResultState RunResultState `json:"result_state,omitempty"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage string `json:"state_message,omitempty"`
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	UserCancelledOrTimedout bool `json:"user_cancelled_or_timedout,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RunTask struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \> 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// If condition_task, specifies a condition with an outcome that can be used
	// to control the execution of other tasks. Does not require a cluster to
	// execute and does not support retries or notifications.
	ConditionTask *RunConditionTask `json:"condition_task,omitempty"`
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this job. When running jobs on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// If for_each_task, indicates that this task must execute the nested task
	// within it.
	ForEachTask *RunForEachTask `json:"for_each_task,omitempty"`
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
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the job. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a new cluster that is created only for
	// this task.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this job must run a notebook. This field
	// may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// If pipeline_task, indicates that this job must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// Parameter values including resolved references
	ResolvedValues *ResolvedValues `json:"resolved_values,omitempty"`
	// The ID of the task run.
	RunId int64 `json:"run_id,omitempty"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `json:"run_if,omitempty"`
	// If run_job_task, indicates that this task must execute another job.
	RunJobTask *RunJobTask `json:"run_job_task,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// If spark_jar_task, indicates that this job must run a JAR.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// If spark_python_task, indicates that this job must run a Python file.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// If `spark_submit_task`, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
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
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// If sql_task, indicates that this job must execute a SQL.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// The current state of the run.
	State *RunState `json:"state,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *RunTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RunTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// * `JOB_RUN`: Normal job run. A run created with :method:jobs/runNow. *
// `WORKFLOW_RUN`: Workflow run. A run created with [dbutils.notebook.run]. *
// `SUBMIT_RUN`: Submit run. A run created with :method:jobs/submit.
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
type RunType string

// Normal job run. A run created with :method:jobs/runNow.
const RunTypeJobRun RunType = `JOB_RUN`

// Submit run. A run created with :method:jobs/submit.
const RunTypeSubmitRun RunType = `SUBMIT_RUN`

// Workflow run. A run created with
// [dbutils.notebook.run](/dev-tools/databricks-utils.html#dbutils-workflow).
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

type Source string

const SourceGit Source = `GIT`

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
	JarUri string `json:"jar_uri,omitempty"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	MainClassName string `json:"main_class_name,omitempty"`
	// Parameters passed to the main method.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	Parameters []string `json:"parameters,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SparkJarTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SparkJarTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	Parameters []string `json:"parameters,omitempty"`
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile string `json:"python_file"`
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local <Databricks>
	// workspace or cloud location (if the `python_file` has a URI format). When
	// set to `GIT`, the Python file will be retrieved from a Git repository
	// defined in `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a <Databricks> workspace or
	// at a cloud filesystem URI. * `GIT`: The Python file is located in a
	// remote Git repository.
	Source Source `json:"source,omitempty"`
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [task parameter variables] such as `{{job.id}}` to pass context about
	// job runs.
	//
	// [task parameter variables]: https://docs.databricks.com/workflows/jobs/parameter-value-references.html
	Parameters []string `json:"parameters,omitempty"`
}

type SqlAlertOutput struct {
	// The state of the SQL alert.
	//
	// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
	// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled
	// trigger conditions
	AlertState SqlAlertState `json:"alert_state,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlAlertOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlAlertOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	WarehouseId string `json:"warehouse_id,omitempty"`
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets []SqlDashboardWidgetOutput `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlDashboardOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlDashboardOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	EndTime int64 `json:"end_time,omitempty"`
	// The information about the error when execution fails.
	Error *SqlOutputError `json:"error,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	StartTime int64 `json:"start_time,omitempty"`
	// The execution status of the SQL widget.
	Status SqlDashboardWidgetOutputStatus `json:"status,omitempty"`
	// The canonical identifier of the SQL widget.
	WidgetId string `json:"widget_id,omitempty"`
	// The title of the SQL widget.
	WidgetTitle string `json:"widget_title,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlDashboardWidgetOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlDashboardWidgetOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The execution status of the SQL widget.
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
	AlertOutput *SqlAlertOutput `json:"alert_output,omitempty"`
	// The output of a SQL dashboard task, if available.
	DashboardOutput *SqlDashboardOutput `json:"dashboard_output,omitempty"`
	// The output of a SQL query task, if available.
	QueryOutput *SqlQueryOutput `json:"query_output,omitempty"`
}

type SqlOutputError struct {
	// The error message when execution fails.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlOutputError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlOutputError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlQueryOutput struct {
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlQueryOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlQueryOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlStatementOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlStatementOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert *SqlTaskAlert `json:"alert,omitempty"`
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard *SqlTaskDashboard `json:"dashboard,omitempty"`
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository. Only one SQL statement is supported in a file. Multiple SQL
	// statements separated by semicolons (;) are not permitted.
	File *SqlTaskFile `json:"file,omitempty"`
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters map[string]string `json:"parameters,omitempty"`
	// If query, indicates that this job must execute a SQL query.
	Query *SqlTaskQuery `json:"query,omitempty"`
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	WarehouseId string `json:"warehouse_id"`
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	AlertId string `json:"alert_id"`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions []SqlTaskSubscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlTaskAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	CustomSubject string `json:"custom_subject,omitempty"`
	// The canonical identifier of the SQL dashboard.
	DashboardId string `json:"dashboard_id"`
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, dashboard snapshots are sent to subscriptions.
	Subscriptions []SqlTaskSubscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlTaskDashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskDashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	Path string `json:"path"`
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local <Databricks> workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in <Databricks> workspace. * `GIT`:
	// SQL file is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	QueryId string `json:"query_id"`
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	DestinationId string `json:"destination_id,omitempty"`
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SqlTaskSubscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SqlTaskSubscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList []iam.AccessControlRequest `json:"access_control_list,omitempty"`
	// An optional set of email addresses notified when the run begins or
	// completes.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
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
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
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
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// The queue settings of the one-time run.
	Queue *QueueSettings `json:"queue,omitempty"`
	// An optional name for the run. The default value is `Untitled`.
	RunName string `json:"run_name,omitempty"`

	Tasks []SubmitTask `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SubmitRun) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubmitRun) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SubmitRunResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubmitRunResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SubmitTask struct {
	// If condition_task, specifies a condition with an outcome that can be used
	// to control the execution of other tasks. Does not require a cluster to
	// execute and does not support retries or notifications.
	ConditionTask *ConditionTask `json:"condition_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency `json:"depends_on,omitempty"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this task. Only all-purpose clusters are supported. When
	// running tasks on an existing cluster, you may need to manually restart
	// the cluster if it stops responding. We suggest running jobs on new
	// clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// If for_each_task, indicates that this must execute the nested task within
	// it for the inputs provided.
	ForEachTask *ForEachTask `json:"for_each_task,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the task. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings `json:"notification_settings,omitempty"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `json:"run_if,omitempty"`
	// If run_job_task, indicates that this job must execute another job.
	RunJobTask *RunJobTask `json:"run_job_task,omitempty"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// If `spark_submit_task`, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
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
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// If sql_task, indicates that this job must execute a SQL.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SubmitTask) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SubmitTask) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TableTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	Condition Condition `json:"condition,omitempty"`
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	TableNames []string `json:"table_names,omitempty"`
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TableTriggerConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TableTriggerConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Task struct {
	// The key of the compute requirement, specified in `job.settings.compute`,
	// to use for execution of this task.
	ComputeKey string `json:"compute_key,omitempty"`
	// If condition_task, specifies a condition with an outcome that can be used
	// to control the execution of other tasks. Does not require a cluster to
	// execute and does not support retries or notifications.
	ConditionTask *ConditionTask `json:"condition_task,omitempty"`
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	DependsOn []TaskDependency `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *TaskEmailNotifications `json:"email_notifications,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this task. Only all-purpose clusters are supported. When
	// running tasks on an existing cluster, you may need to manually restart
	// the cluster if it stops responding. We suggest running jobs on new
	// clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// If for_each_task, indicates that this must execute the nested task within
	// it for the inputs provided.
	ForEachTask *ForEachTask `json:"for_each_task,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules `json:"health,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the task. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries int `json:"max_retries,omitempty"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int `json:"min_retry_interval_millis,omitempty"`
	// If new_cluster, a description of a cluster that is created for only for
	// this task.
	NewCluster *compute.ClusterSpec `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	NotificationSettings *TaskNotificationSettings `json:"notification_settings,omitempty"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// An optional policy to specify whether to retry a task when it times out.
	RetryOnTimeout bool `json:"retry_on_timeout,omitempty"`
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	RunIf RunIf `json:"run_if,omitempty"`
	// If run_job_task, indicates that this task must execute another job.
	RunJobTask *RunJobTask `json:"run_job_task,omitempty"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// If `spark_submit_task`, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
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
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// If sql_task, indicates that this job must execute a SQL task.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	WebhookNotifications *WebhookNotifications `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Task) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Task) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome string `json:"outcome,omitempty"`
	// The name of the task this task depends on.
	TaskKey string `json:"task_key"`

	ForceSendFields []string `json:"-"`
}

func (s *TaskDependency) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskDependency) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TaskEmailNotifications struct {
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	AlertOnLastAttempt bool `json:"alert_on_last_attempt,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TaskNotificationSettings) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskNotificationSettings) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TriggerInfo struct {
	// The run id of the Run Job task run
	RunId int `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TriggerInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TriggerInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TriggerSettings struct {
	// File arrival trigger settings.
	FileArrival *FileArrivalTriggerConfiguration `json:"file_arrival,omitempty"`
	// Whether this trigger is paused or not.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// Table trigger settings.
	Table *TableTriggerConfiguration `json:"table,omitempty"`
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
// `TABLE`: Indicates a run that is triggered by a table update.
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
	FieldsToRemove []string `json:"fields_to_remove,omitempty"`
	// The canonical identifier of the job to update. This field is required.
	JobId int64 `json:"job_id"`
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
	NewSettings *JobSettings `json:"new_settings,omitempty"`
}

type ViewItem struct {
	// Content of the view.
	Content string `json:"content,omitempty"`
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name string `json:"name,omitempty"`
	// Type of the view item.
	Type ViewType `json:"type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ViewItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ViewItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Webhook) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Webhook) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	OnDurationWarningThresholdExceeded []WebhookNotificationsOnDurationWarningThresholdExceededItem `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure []Webhook `json:"on_failure,omitempty"`
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart []Webhook `json:"on_start,omitempty"`
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess []Webhook `json:"on_success,omitempty"`
}

type WebhookNotificationsOnDurationWarningThresholdExceededItem struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *WebhookNotificationsOnDurationWarningThresholdExceededItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WebhookNotificationsOnDurationWarningThresholdExceededItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
