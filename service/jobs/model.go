// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"fmt"

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
	// The continuous trigger that triggered this run.
	Continuous *Continuous `json:"continuous,omitempty"`
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
	// An optional specification for a remote repository containing the
	// notebooks used by this job's notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
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
	// An optional name for the run. The maximum allowed length is 4096 bytes in
	// UTF-8 encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`
	// This describes an enum
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
	// The result and lifecycle states of the run.
	State *RunState `json:"state,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks []RunTask `json:"tasks,omitempty"`
	// This describes an enum
	Trigger TriggerType `json:"trigger,omitempty"`
}

type CancelAllRuns struct {
	// The canonical identifier of the job to cancel all runs of. This field is
	// required.
	JobId int64 `json:"job_id"`
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
	NewCluster *compute.BaseClusterInfo `json:"new_cluster,omitempty"`
}

type Continuous struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus ContinuousPauseStatus `json:"pause_status,omitempty"`
}

// Indicate whether the continuous execution of the job is paused or not.
// Defaults to UNPAUSED.
type ContinuousPauseStatus string

const ContinuousPauseStatusPaused ContinuousPauseStatus = `PAUSED`

const ContinuousPauseStatusUnpaused ContinuousPauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *ContinuousPauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ContinuousPauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = ContinuousPauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns ContinuousPauseStatus to satisfy [pflag.Value] interface
func (f *ContinuousPauseStatus) Type() string {
	return "ContinuousPauseStatus"
}

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList []iam.AccessControlRequest `json:"access_control_list,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous `json:"continuous,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format CreateJobFormat `json:"format,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job's notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
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
	// This value cannot exceed 1000\. Setting this value to 0 causes all new
	// runs to be skipped. The default behavior is to allow only 1 concurrent
	// run.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
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
	Tasks []JobTaskSettings `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. The default behavior
	// is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// Trigger settings for the job. Can be used to trigger a run when new files
	// arrive in an external location. The default behavior is that the job runs
	// only when triggered by clicking “Run Now” in the Jobs UI or sending
	// an API request to `runNow`.
	Trigger *TriggerSettings `json:"trigger,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	WebhookNotifications *JobWebhookNotifications `json:"webhook_notifications,omitempty"`
}

// Used to tell what is the format of the job. This field is ignored in
// Create/Update/Reset calls. When using the Jobs API 2.1 this value is always
// set to `"MULTI_TASK"`.
type CreateJobFormat string

const CreateJobFormatMultiTask CreateJobFormat = `MULTI_TASK`

const CreateJobFormatSingleTask CreateJobFormat = `SINGLE_TASK`

// String representation for [fmt.Print]
func (f *CreateJobFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateJobFormat) Set(v string) error {
	switch v {
	case `MULTI_TASK`, `SINGLE_TASK`:
		*f = CreateJobFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MULTI_TASK", "SINGLE_TASK"`, v)
	}
}

// Type always returns CreateJobFormat to satisfy [pflag.Value] interface
func (f *CreateJobFormat) Type() string {
	return "CreateJobFormat"
}

type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId int64 `json:"job_id,omitempty"`
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus CronSchedulePauseStatus `json:"pause_status,omitempty"`
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

// Indicate whether this schedule is paused or not.
type CronSchedulePauseStatus string

const CronSchedulePauseStatusPaused CronSchedulePauseStatus = `PAUSED`

const CronSchedulePauseStatusUnpaused CronSchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *CronSchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CronSchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = CronSchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns CronSchedulePauseStatus to satisfy [pflag.Value] interface
func (f *CronSchedulePauseStatus) Type() string {
	return "CronSchedulePauseStatus"
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders map[string]string `json:"artifacts_headers,omitempty"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink string `json:"artifacts_link,omitempty"`
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
	// Optional (relative) path to the project directory, if no value is
	// provided, the root of the git repository is used.
	ProjectDirectory string `json:"project_directory,omitempty"`
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema string `json:"schema,omitempty"`
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId string `json:"warehouse_id,omitempty"`
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
	// The exported content in HTML format (one for every view item).
	Views []ViewItem `json:"views,omitempty"`
}

// Export and retrieve a job run
type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 `json:"-" url:"run_id"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport `json:"-" url:"views_to_export,omitempty"`
}

type FileArrivalTriggerSettings struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggerSeconds int `json:"min_time_between_trigger_seconds,omitempty"`
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	Url string `json:"url,omitempty"`
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`
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
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId int64 `json:"-" url:"run_id"`
}

// Read-only state of the remote repository at the time the job was run. This
// field is only included on job runs.
type GitSnapshot struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	UsedCommit string `json:"used_commit,omitempty"`
}

// An optional specification for a remote repository containing the notebooks
// used by this job's notebook tasks.
type GitSource struct {
	// Name of the branch to be checked out and used by this job. This field
	// cannot be specified in conjunction with git_tag or git_commit.
	//
	// The maximum length is 255 characters.
	GitBranch string `json:"git_branch,omitempty"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag. The maximum length
	// is 64 characters.
	GitCommit string `json:"git_commit,omitempty"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitSourceGitProvider `json:"git_provider"`
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot *GitSnapshot `json:"git_snapshot,omitempty"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	//
	// The maximum length is 255 characters.
	GitTag string `json:"git_tag,omitempty"`
	// URL of the repository to be cloned by this job. The maximum length is 300
	// characters.
	GitUrl string `json:"git_url"`
}

// Unique identifier of the service used to host the Git repository. The value
// is case insensitive.
type GitSourceGitProvider string

const GitSourceGitProviderAwscodecommit GitSourceGitProvider = `awsCodeCommit`

const GitSourceGitProviderAzuredevopsservices GitSourceGitProvider = `azureDevOpsServices`

const GitSourceGitProviderBitbucketcloud GitSourceGitProvider = `bitbucketCloud`

const GitSourceGitProviderBitbucketserver GitSourceGitProvider = `bitbucketServer`

const GitSourceGitProviderGithub GitSourceGitProvider = `gitHub`

const GitSourceGitProviderGithubenterprise GitSourceGitProvider = `gitHubEnterprise`

const GitSourceGitProviderGitlab GitSourceGitProvider = `gitLab`

const GitSourceGitProviderGitlabenterpriseedition GitSourceGitProvider = `gitLabEnterpriseEdition`

// String representation for [fmt.Print]
func (f *GitSourceGitProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GitSourceGitProvider) Set(v string) error {
	switch v {
	case `awsCodeCommit`, `azureDevOpsServices`, `bitbucketCloud`, `bitbucketServer`, `gitHub`, `gitHubEnterprise`, `gitLab`, `gitLabEnterpriseEdition`:
		*f = GitSourceGitProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "awsCodeCommit", "azureDevOpsServices", "bitbucketCloud", "bitbucketServer", "gitHub", "gitHubEnterprise", "gitLab", "gitLabEnterpriseEdition"`, v)
	}
}

// Type always returns GitSourceGitProvider to satisfy [pflag.Value] interface
func (f *GitSourceGitProvider) Type() string {
	return "GitSourceGitProvider"
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
	// History of the file arrival trigger associated with the job.
	TriggerHistory *TriggerHistory `json:"trigger_history,omitempty"`
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey string `json:"job_cluster_key"`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster *compute.BaseClusterInfo `json:"new_cluster,omitempty"`
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `SKIPPED`,
	// `FAILED`, or `TIMED_OUT` result_state. If this is not specified on job
	// creation, reset, or update the list is empty, and notifications are not
	// sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESSFUL` result_state.
	// If not specified on job creation, reset, or update, the list is empty,
	// and notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
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
}

type JobSettings struct {
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous `json:"continuous,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format JobSettingsFormat `json:"format,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job's notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
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
	// This value cannot exceed 1000\. Setting this value to 0 causes all new
	// runs to be skipped. The default behavior is to allow only 1 concurrent
	// run.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
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
	Tasks []JobTaskSettings `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. The default behavior
	// is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// Trigger settings for the job. Can be used to trigger a run when new files
	// arrive in an external location. The default behavior is that the job runs
	// only when triggered by clicking “Run Now” in the Jobs UI or sending
	// an API request to `runNow`.
	Trigger *TriggerSettings `json:"trigger,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	WebhookNotifications *JobWebhookNotifications `json:"webhook_notifications,omitempty"`
}

// Used to tell what is the format of the job. This field is ignored in
// Create/Update/Reset calls. When using the Jobs API 2.1 this value is always
// set to `"MULTI_TASK"`.
type JobSettingsFormat string

const JobSettingsFormatMultiTask JobSettingsFormat = `MULTI_TASK`

const JobSettingsFormatSingleTask JobSettingsFormat = `SINGLE_TASK`

// String representation for [fmt.Print]
func (f *JobSettingsFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobSettingsFormat) Set(v string) error {
	switch v {
	case `MULTI_TASK`, `SINGLE_TASK`:
		*f = JobSettingsFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MULTI_TASK", "SINGLE_TASK"`, v)
	}
}

// Type always returns JobSettingsFormat to satisfy [pflag.Value] interface
func (f *JobSettingsFormat) Type() string {
	return "JobSettingsFormat"
}

type JobTaskSettings struct {
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task. This field is required when a job
	// consists of more than one task.
	DependsOn []TaskDependenciesItem `json:"depends_on,omitempty"`
	// An optional description for this task. The maximum length is 4096 bytes.
	Description string `json:"description,omitempty"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *TaskEmailNotifications `json:"email_notifications,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this task. When running tasks on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the task. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value -1 means
	// to retry indefinitely and the value 0 means to never retry. The default
	// behavior is to never retry.
	MaxRetries int `json:"max_retries,omitempty"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int `json:"min_retry_interval_millis,omitempty"`
	// If new_cluster, a description of a cluster that is created for only for
	// this task.
	NewCluster *compute.BaseClusterInfo `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` for this task.
	NotificationSettings *TaskNotificationSettings `json:"notification_settings,omitempty"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// An optional policy to specify whether to retry a task when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout bool `json:"retry_on_timeout,omitempty"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// If spark_submit_task, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// If sql_task, indicates that this job must execute a SQL task.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset. The maximum length is 100 characters.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. The default
	// behavior is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

type JobWebhookNotifications struct {
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure []JobWebhookNotificationsOnFailureItem `json:"on_failure,omitempty"`
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart []JobWebhookNotificationsOnStartItem `json:"on_start,omitempty"`
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess []JobWebhookNotificationsOnSuccessItem `json:"on_success,omitempty"`
}

type JobWebhookNotificationsOnFailureItem struct {
	Id string `json:"id,omitempty"`
}

type JobWebhookNotificationsOnStartItem struct {
	Id string `json:"id,omitempty"`
}

type JobWebhookNotificationsOnSuccessItem struct {
	Id string `json:"id,omitempty"`
}

// List all jobs
type ListJobsRequest struct {
	// Whether to include task and cluster details in the response.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 25. The default value is 20.
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
}

// List runs for a job
type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `PENDING`, `RUNNING`, or `TERMINATING`. This field cannot be
	// `true` when completed_only is `true`.
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
	// than 25. The default value is 25. If a request specifies a limit of 0,
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
}

// This describes an enum
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
}

type NotebookTask struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to :method:jobs/runNow with parameters specified, the
	// two parameters maps are merged. If the same key is specified in
	// `base_parameters` and in `run-now`, the value from `run-now` is used.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// If the notebook takes a parameter that is not specified in the job’s
	// `base_parameters` or the `run-now` override parameters, the default value
	// from the notebook is used.
	//
	// Retrieve these parameters in a notebook using [dbutils.widgets.get].
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-widgets
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath string `json:"notebook_path"`
	// This describes an enum
	Source NotebookTaskSource `json:"source,omitempty"`
}

// This describes an enum
type NotebookTaskSource string

// Notebook is located in cloud Git provider.
const NotebookTaskSourceGit NotebookTaskSource = `GIT`

// Notebook is located in <Databricks> workspace.
const NotebookTaskSourceWorkspace NotebookTaskSource = `WORKSPACE`

// String representation for [fmt.Print]
func (f *NotebookTaskSource) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NotebookTaskSource) Set(v string) error {
	switch v {
	case `GIT`, `WORKSPACE`:
		*f = NotebookTaskSource(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GIT", "WORKSPACE"`, v)
	}
}

// Type always returns NotebookTaskSource to satisfy [pflag.Value] interface
func (f *NotebookTaskSource) Type() string {
	return "NotebookTaskSource"
}

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
}

type PipelineTask struct {
	// If true, a full refresh will be triggered on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// The full name of the pipeline task to execute.
	PipelineId string `json:"pipeline_id,omitempty"`
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
}

type RepairHistoryItem struct {
	// The end time of the (repaired) run.
	EndTime int64 `json:"end_time,omitempty"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id int64 `json:"id,omitempty"`
	// The start time of the (repaired) run.
	StartTime int64 `json:"start_time,omitempty"`
	// The result and lifecycle state of the run.
	State *RunState `json:"state,omitempty"`
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds []int64 `json:"task_run_ids,omitempty"`
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type RepairHistoryItemType `json:"type,omitempty"`
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
	// `\"jar_params\": [\"john doe\", \"35\"]`. The parameters are used to
	// invoke the main function of the main class specified in the Spark JAR
	// task. If not specified upon `run-now`, it defaults to an empty list.
	// jar_params cannot be specified in conjunction with notebook_params. The
	// JSON representation of this field (for example `{\"jar_params\":[\"john
	// doe\",\"35\"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html"#parameter-variables") to set
	// parameters containing information about job runs.
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId int64 `json:"latest_repair_id,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `\"notebook_params\": {\"name\": \"john doe\", \"age\": \"35\"}`. The map
	// is passed to the notebook and is accessible through the
	// [dbutils.widgets.get] function.
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
	// `{\"notebook_params\":{\"name\":\"john doe\",\"age\":\"35\"}}`) cannot
	// exceed 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `"python_named_params": {"name": "task", "data":
	// "dbfs:/path/to/data.json"}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `\"python_params\": [\"john doe\", \"35\"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{\"python_params\":[\"john
	// doe\",\"35\"]}`) cannot exceed 10,000 bytes.
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
	PythonParams []string `json:"python_params,omitempty"`
	// If true, repair all failed tasks. Only one of rerun_tasks or
	// rerun_all_failed_tasks can be used.
	RerunAllFailedTasks bool `json:"rerun_all_failed_tasks,omitempty"`
	// The task keys of the task runs to repair.
	RerunTasks []string `json:"rerun_tasks,omitempty"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId int64 `json:"run_id"`
	// A list of parameters for jobs with spark submit task, for example
	// `\"spark_submit_params\": [\"--class\",
	// \"org.apache.spark.examples.SparkPi\"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{\"python_params\":[\"john doe\",\"35\"]}`) cannot exceed 10,000 bytes.
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
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId int64 `json:"repair_id,omitempty"`
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
	// The continuous trigger that triggered this run.
	Continuous *Continuous `json:"continuous,omitempty"`
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
	// An optional specification for a remote repository containing the
	// notebooks used by this job's notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
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
	// An optional name for the run. The maximum allowed length is 4096 bytes in
	// UTF-8 encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`
	// This describes an enum
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
	// The result and lifecycle states of the run.
	State *RunState `json:"state,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls.
	Tasks []RunTask `json:"tasks,omitempty"`
	// This describes an enum
	Trigger TriggerType `json:"trigger,omitempty"`
}

// This describes an enum
type RunLifeCycleState string

// The run is blocked on an upstream dependency.
const RunLifeCycleStateBlocked RunLifeCycleState = `BLOCKED`

// An exceptional state that indicates a failure in the Jobs service, such as
// network failure over a long period. If a run on a new cluster ends in the
// `INTERNAL_ERROR` state, the Jobs service terminates the cluster as soon as
// possible. This state is terminal.
const RunLifeCycleStateInternalError RunLifeCycleState = `INTERNAL_ERROR`

// The run has been triggered. If there is not already an active run of the same
// job, the cluster and execution context are being prepared. If there is
// already an active run of the same job, the run immediately transitions into
// the `SKIPPED` state without preparing any resources.
const RunLifeCycleStatePending RunLifeCycleState = `PENDING`

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
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `RUNNING`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = RunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "RUNNING", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
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
	// `\"jar_params\": [\"john doe\", \"35\"]`. The parameters are used to
	// invoke the main function of the main class specified in the Spark JAR
	// task. If not specified upon `run-now`, it defaults to an empty list.
	// jar_params cannot be specified in conjunction with notebook_params. The
	// JSON representation of this field (for example `{\"jar_params\":[\"john
	// doe\",\"35\"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html"#parameter-variables") to set
	// parameters containing information about job runs.
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the job to be executed
	JobId int64 `json:"job_id"`
	// A map from keys to values for jobs with notebook task, for example
	// `\"notebook_params\": {\"name\": \"john doe\", \"age\": \"35\"}`. The map
	// is passed to the notebook and is accessible through the
	// [dbutils.widgets.get] function.
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
	// `{\"notebook_params\":{\"name\":\"john doe\",\"age\":\"35\"}}`) cannot
	// exceed 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `"python_named_params": {"name": "task", "data":
	// "dbfs:/path/to/data.json"}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `\"python_params\": [\"john doe\", \"35\"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{\"python_params\":[\"john
	// doe\",\"35\"]}`) cannot exceed 10,000 bytes.
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
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `\"spark_submit_params\": [\"--class\",
	// \"org.apache.spark.examples.SparkPi\"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{\"python_params\":[\"john doe\",\"35\"]}`) cannot exceed 10,000 bytes.
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
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// The globally unique ID of the newly triggered run.
	RunId int64 `json:"run_id,omitempty"`
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
	// The output of a SQL task, if available.
	SqlOutput *SqlOutput `json:"sql_output,omitempty"`
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `\"jar_params\": [\"john doe\", \"35\"]`. The parameters are used to
	// invoke the main function of the main class specified in the Spark JAR
	// task. If not specified upon `run-now`, it defaults to an empty list.
	// jar_params cannot be specified in conjunction with notebook_params. The
	// JSON representation of this field (for example `{\"jar_params\":[\"john
	// doe\",\"35\"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables](/jobs.html"#parameter-variables") to set
	// parameters containing information about job runs.
	JarParams []string `json:"jar_params,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `\"notebook_params\": {\"name\": \"john doe\", \"age\": \"35\"}`. The map
	// is passed to the notebook and is accessible through the
	// [dbutils.widgets.get] function.
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
	// `{\"notebook_params\":{\"name\":\"john doe\",\"age\":\"35\"}}`) cannot
	// exceed 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *PipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `"python_named_params": {"name": "task", "data":
	// "dbfs:/path/to/data.json"}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `\"python_params\": [\"john doe\", \"35\"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{\"python_params\":[\"john
	// doe\",\"35\"]}`) cannot exceed 10,000 bytes.
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
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `\"spark_submit_params\": [\"--class\",
	// \"org.apache.spark.examples.SparkPi\"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{\"python_params\":[\"john doe\",\"35\"]}`) cannot exceed 10,000 bytes.
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
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

// This describes an enum
type RunResultState string

// The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

// The task completed with an error.
const RunResultStateFailed RunResultState = `FAILED`

// The task completed successfully.
const RunResultStateSuccess RunResultState = `SUCCESS`

// The run was stopped after reaching the timeout.
const RunResultStateTimedout RunResultState = `TIMEDOUT`

// String representation for [fmt.Print]
func (f *RunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `FAILED`, `SUCCESS`, `TIMEDOUT`:
		*f = RunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "FAILED", "SUCCESS", "TIMEDOUT"`, v)
	}
}

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

// The result and lifecycle state of the run.
type RunState struct {
	// A description of a run’s current location in the run lifecycle. This
	// field is always available in the response.
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty"`
	// This describes an enum
	ResultState RunResultState `json:"result_state,omitempty"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage string `json:"state_message,omitempty"`
	// Whether a run was canceled manually by a user or by the scheduler because
	// the run timed out.
	UserCancelledOrTimedout bool `json:"user_cancelled_or_timedout,omitempty"`
}

type RunSubmitTaskSettings struct {
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task. This field is required when a job
	// consists of more than one task.
	DependsOn []TaskDependenciesItem `json:"depends_on,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this task. When running tasks on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the task. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *compute.BaseClusterInfo `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// If pipeline_task, indicates that this task must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// If spark_jar_task, indicates that this task must run a JAR.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// If spark_python_task, indicates that this task must run a Python file.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// If spark_submit_task, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters.
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// If sql_task, indicates that this job must execute a SQL.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset. The maximum length is 100 characters.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. The default
	// behavior is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
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
	// If dbt_task, indicates that this must execute a dbt task. It requires
	// both Databricks SQL and the ability to use a serverless or a pro SQL
	// warehouse.
	DbtTask *DbtTask `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task. This field is required when a job
	// consists of more than one task.
	DependsOn []TaskDependenciesItem `json:"depends_on,omitempty"`
	// An optional description for this task. The maximum length is 4096 bytes.
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
	// An optional specification for a remote repository containing the
	// notebooks used by this job's notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the job. The default value is an empty list.
	Libraries []compute.Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a new cluster that is created only for
	// this task.
	NewCluster *compute.BaseClusterInfo `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this job must run a notebook. This field
	// may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// If pipeline_task, indicates that this job must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// The ID of the task run.
	RunId int64 `json:"run_id,omitempty"`
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
	// If spark_submit_task, indicates that this task must be launched by the
	// spark submit script. This task can run only on new clusters
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// If sql_task, indicates that this job must execute a SQL.
	SqlTask *SqlTask `json:"sql_task,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// The result and lifecycle states of the run.
	State *RunState `json:"state,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset. The maximum length is 100 characters.
	TaskKey string `json:"task_key,omitempty"`
}

// This describes an enum
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

type SparkJarTask struct {
	// Deprecated since 04/2016\\. Provide a `jar` through the `libraries` field
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
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile string `json:"python_file"`
	// This describes an enum
	Source SparkPythonTaskSource `json:"source,omitempty"`
}

// This describes an enum
type SparkPythonTaskSource string

// The Python file is located in a remote Git repository.
const SparkPythonTaskSourceGit SparkPythonTaskSource = `GIT`

// The Python file is located in a <Databricks> workspace or at a cloud
// filesystem URI.
const SparkPythonTaskSourceWorkspace SparkPythonTaskSource = `WORKSPACE`

// String representation for [fmt.Print]
func (f *SparkPythonTaskSource) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SparkPythonTaskSource) Set(v string) error {
	switch v {
	case `GIT`, `WORKSPACE`:
		*f = SparkPythonTaskSource(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GIT", "WORKSPACE"`, v)
	}
}

// Type always returns SparkPythonTaskSource to satisfy [pflag.Value] interface
func (f *SparkPythonTaskSource) Type() string {
	return "SparkPythonTaskSource"
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
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
	Widgets *SqlDashboardWidgetOutput `json:"widgets,omitempty"`
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
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert *SqlTaskAlert `json:"alert,omitempty"`
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard *SqlTaskDashboard `json:"dashboard,omitempty"`
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	File *SqlTaskFile `json:"file,omitempty"`
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters map[string]string `json:"parameters,omitempty"`
	// If query, indicates that this job must execute a SQL query.
	Query *SqlTaskQuery `json:"query,omitempty"`
	// The canonical identifier of the SQL warehouse. Only serverless and pro
	// SQL warehouses are supported.
	WarehouseId string `json:"warehouse_id"`
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	AlertId string `json:"alert_id"`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions []SqlTaskSubscription `json:"subscriptions,omitempty"`
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
}

type SqlTaskFile struct {
	// Relative path of the SQL file in the remote Git repository.
	Path string `json:"path"`
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	QueryId string `json:"query_id"`
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification.
	DestinationId string `json:"destination_id,omitempty"`
	// The user name to receive the subscription email.
	UserName string `json:"user_name,omitempty"`
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList []iam.AccessControlRequest `json:"access_control_list,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job's notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
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
	// to each of the `webhook_notifications` for this run.
	NotificationSettings *JobNotificationSettings `json:"notification_settings,omitempty"`
	// An optional name for the run. The default value is `Untitled`.
	RunName string `json:"run_name,omitempty"`

	Tasks []RunSubmitTaskSettings `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. The default behavior
	// is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	WebhookNotifications *JobWebhookNotifications `json:"webhook_notifications,omitempty"`
}

type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId int64 `json:"run_id,omitempty"`
}

type TaskDependenciesItem struct {
	TaskKey string `json:"task_key,omitempty"`
}

type TaskEmailNotifications struct {
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `SKIPPED`,
	// `FAILED`, or `TIMED_OUT` result_state. If this is not specified on job
	// creation, reset, or update the list is empty, and notifications are not
	// sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESSFUL` result_state.
	// If not specified on job creation, reset, or update, the list is empty,
	// and notifications are not sent.
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
}

type TriggerEvaluation struct {
	// Human-readable description of the the trigger evaluation result. Explains
	// why the trigger evaluation triggered or did not trigger a run, or failed.
	Description string `json:"description,omitempty"`
	// The ID of the run that was triggered by the trigger evaluation. Only
	// returned if a run was triggered.
	RunId int64 `json:"run_id,omitempty"`
	// Timestamp at which the trigger was evaluated.
	Timestamp int64 `json:"timestamp,omitempty"`
}

type TriggerHistory struct {
	// The last time the trigger failed to evaluate.
	LastFailed *TriggerEvaluation `json:"last_failed,omitempty"`
	// The last time the trigger was evaluated but did not trigger a run.
	LastNotTriggered *TriggerEvaluation `json:"last_not_triggered,omitempty"`
	// The last time the run was triggered due to a file arrival.
	LastTriggered *TriggerEvaluation `json:"last_triggered,omitempty"`
}

type TriggerSettings struct {
	// File arrival trigger settings.
	FileArrival *FileArrivalTriggerSettings `json:"file_arrival,omitempty"`
	// Whether this trigger is paused or not.
	PauseStatus TriggerSettingsPauseStatus `json:"pause_status,omitempty"`
}

// Whether this trigger is paused or not.
type TriggerSettingsPauseStatus string

const TriggerSettingsPauseStatusPaused TriggerSettingsPauseStatus = `PAUSED`

const TriggerSettingsPauseStatusUnpaused TriggerSettingsPauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *TriggerSettingsPauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TriggerSettingsPauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = TriggerSettingsPauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns TriggerSettingsPauseStatus to satisfy [pflag.Value] interface
func (f *TriggerSettingsPauseStatus) Type() string {
	return "TriggerSettingsPauseStatus"
}

// This describes an enum
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

// String representation for [fmt.Print]
func (f *TriggerType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TriggerType) Set(v string) error {
	switch v {
	case `FILE_ARRIVAL`, `ONE_TIME`, `PERIODIC`, `RETRY`:
		*f = TriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_ARRIVAL", "ONE_TIME", "PERIODIC", "RETRY"`, v)
	}
}

// Type always returns TriggerType to satisfy [pflag.Value] interface
func (f *TriggerType) Type() string {
	return "TriggerType"
}

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported. This field is optional.
	FieldsToRemove []string `json:"fields_to_remove,omitempty"`
	// The canonical identifier of the job to update. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings for the job. Any top-level fields specified in
	// `new_settings` are completely replaced. Partially updating nested fields
	// is not supported.
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
}

// This describes an enum
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

// This describes an enum
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
