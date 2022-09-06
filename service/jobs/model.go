// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

// all definitions in this file are in alphabetical order

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. max_workers must be strictly greater than min_workers.
	MaxWorkers int `json:"max_workers,omitempty"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster has
	// after creation.
	MinWorkers int `json:"min_workers,omitempty"`
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
	// The logs continue to be available after the run completes. The response
	// won?t include this field if the identifier is not available yet.
	ClusterId string `json:"cluster_id,omitempty"`
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed. The response won?t include
	// this field if the identifier is not available yet.
	SparkContextId string `json:"spark_context_id,omitempty"`
}

type ClusterLogConf struct {
	// DBFS location of cluster log. Destination must be provided. For example,
	// `{ &#34;dbfs&#34; : { &#34;destination&#34; : &#34;dbfs:/home/cluster_log&#34; } }`
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// S3 location of cluster log. `destination` and either `region` or
	// `endpoint` must be provided. For example, `{ &#34;s3&#34;: { &#34;destination&#34; :
	// &#34;s3://cluster_log_bucket/prefix&#34;, &#34;region&#34; : &#34;us-west-2&#34; } }`
	S3 any/* ERROR */ `json:"s3,omitempty"`
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this job. When running jobs on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the job. The default value is an empty list.
	Libraries []Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *NewCluster `json:"new_cluster,omitempty"`
}

// An object with key value pairs. The key length must be between 1 and 127
// UTF-8 characters, inclusive. The value length must be less than or equal to
// 255 UTF-8 characters.
type ClusterTag map[string]string

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList any/* MISSING TYPE */ `json:"access_control_list,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `&#34;MULTI_TASK&#34;`.
	Format CreateJobFormat `json:"format,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job&#39;s notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job?s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won?t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000\. Setting this value to
	// 0 causes all new runs to be skipped. The default behavior is to allow
	// only 1 concurrent run.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job.
	Name string `json:"name,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking ?Run Now? in the Jobs UI or
	// sending an API request to `runNow`.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags any/* MISSING TYPE */ `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job.
	Tasks []JobTaskSettings `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. The default behavior
	// is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

// Used to tell what is the format of the job. This field is ignored in
// Create/Update/Reset calls. When using the Jobs API 2.1 this value is always
// set to `&#34;MULTI_TASK&#34;`.
type CreateJobFormat string

const CreateJobFormatMultiTask CreateJobFormat = `MULTI_TASK`

const CreateJobFormatSingleTask CreateJobFormat = `SINGLE_TASK`

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus CronSchedulePauseStatus `json:"pause_status,omitempty"`
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron
	// Trigger](http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html)
	// for details. This field is required.
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java
	// TimeZone](https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html)
	// for details. This field is required.
	TimezoneId string `json:"timezone_id"`
}

// Indicate whether this schedule is paused or not.
type CronSchedulePauseStatus string

const CronSchedulePauseStatusPaused CronSchedulePauseStatus = `PAUSED`

const CronSchedulePauseStatusUnpaused CronSchedulePauseStatus = `UNPAUSED`

type DbfsStorageInfo struct {
	// DBFS destination. Example: `dbfs:/my/path`
	Destination string `json:"destination,omitempty"`
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId int64 `json:"job_id"`
}

type DeleteRun struct {
	// The canonical identifier of the run for which to retrieve the metadata.
	RunId int64 `json:"run_id,omitempty"`
}

type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item).
	Views []ViewItem `json:"views,omitempty"`
}

type FileStorageInfo struct {
	// File destination. Example: `file:/my/file.sh`
	Destination string `json:"destination,omitempty"`
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
// used by this job&#39;s notebook tasks.
type GitSource struct {
	// Name of the branch to be checked out and used by this job. This field
	// cannot be specified in conjunction with git_tag or git_commit. The
	// maximum length is 255 characters.
	GitBranch string `json:"git_branch"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag. The maximum length
	// is 64 characters.
	GitCommit string `json:"git_commit,omitempty"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitSourceGitProvider `json:"git_provider"`

	GitSnapshot *GitSnapshot `json:"git_snapshot,omitempty"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit. The maximum
	// length is 255 characters.
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

type InitScriptInfo struct {
	// S3 location of init script. Destination and either region or endpoint
	// must be provided. For example, `{ &#34;s3&#34;: { &#34;destination&#34; :
	// &#34;s3://init_script_bucket/prefix&#34;, &#34;region&#34; : &#34;us-west-2&#34; } }`
	S3 any/* ERROR */ `json:"S3,omitempty"`
	// DBFS location of init script. Destination must be provided. For example,
	// `{ &#34;dbfs&#34; : { &#34;destination&#34; : &#34;dbfs:/home/init_script&#34; } }`
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// File location of init script. Destination must be provided. For example,
	// `{ &#34;file&#34; : { &#34;destination&#34; : &#34;file:/my/local/file.sh&#34; } }`
	File *FileStorageInfo `json:"file,omitempty"`
}

type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won?t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings `json:"settings,omitempty"`
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey string `json:"job_cluster_key"`

	NewCluster *NewCluster `json:"new_cluster,omitempty"`
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

type JobSettings struct {
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `&#34;MULTI_TASK&#34;`.
	Format JobSettingsFormat `json:"format,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job&#39;s notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job?s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won?t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000\. Setting this value to
	// 0 causes all new runs to be skipped. The default behavior is to allow
	// only 1 concurrent run.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job.
	Name string `json:"name,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking ?Run Now? in the Jobs UI or
	// sending an API request to `runNow`.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags any/* MISSING TYPE */ `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job.
	Tasks []JobTaskSettings `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. The default behavior
	// is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

// Used to tell what is the format of the job. This field is ignored in
// Create/Update/Reset calls. When using the Jobs API 2.1 this value is always
// set to `&#34;MULTI_TASK&#34;`.
type JobSettingsFormat string

const JobSettingsFormatMultiTask JobSettingsFormat = `MULTI_TASK`

const JobSettingsFormatSingleTask JobSettingsFormat = `SINGLE_TASK`

type JobTaskSettings struct {
	DependsOn []TaskDependenciesItem `json:"depends_on,omitempty"`

	Description string `json:"description,omitempty"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications `json:"email_notifications,omitempty"`
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
	Libraries []Library `json:"libraries,omitempty"`
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
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *NewCluster `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this task must run a notebook. This
	// field may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
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
	// spark submit script.
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`

	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. The default
	// behavior is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

type Library struct {
	// If cran, specification of a CRAN library to be installed.
	Cran *RCranLibrary `json:"cran,omitempty"`

	Egg string `json:"egg,omitempty"`

	Jar string `json:"jar,omitempty"`
	// If maven, specification of a Maven library to be installed. For example:
	// `{ &#34;coordinates&#34;: &#34;org.jsoup:jsoup:1.7.2&#34; }`
	Maven *MavenLibrary `json:"maven,omitempty"`
	// If pypi, specification of a PyPI library to be installed. Specifying the
	// `repo` field is optional and if not specified, the default pip index is
	// used. For example: `{ &#34;package&#34;: &#34;simplejson&#34;, &#34;repo&#34;:
	// &#34;https://my-repo.com&#34; }`
	Pypi *PythonPyPiLibrary `json:"pypi,omitempty"`

	Whl string `json:"whl,omitempty"`
}

type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// A list of runs, from most recently started to least.
	Runs []Run `json:"runs,omitempty"`
}

type MavenLibrary struct {
	// Gradle-style Maven coordinates. For example: `org.jsoup:jsoup:1.7.2`.
	// This field is required.
	Coordinates string `json:"coordinates"`
	// List of dependences to exclude. For example: `[&#34;slf4j:slf4j&#34;,
	// &#34;*:hadoop-client&#34;]`. Maven dependency exclusions:
	// &lt;https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html&gt;.
	Exclusions []string `json:"exclusions,omitempty"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo string `json:"repo,omitempty"`
}

type NewCluster struct {
	// If autoscale, the required parameters to automatically scale clusters up
	// and down based on load.
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values is used.
	AwsAttributes any/* ERROR */ `json:"aws_attributes,omitempty"`
	// Defines attributes such as the instance availability type, node
	// placement, and max bid price. If not specified during cluster creation, a
	// set of default values is used.
	AzureAttributes any/* ERROR */ `json:"azure_attributes,omitempty"`
	// The configuration for delivering Spark logs to a long-term storage
	// destination. Only one destination can be specified for one cluster. If
	// the conf is given, the logs are delivered to the destination every `5
	// mins`. The destination of driver logs is
	// `&lt;destination&gt;/&lt;cluster-id&gt;/driver`, while the destination of executor
	// logs is `&lt;destination&gt;/&lt;cluster-id&gt;/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`

	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// The optional ID of the instance pool to use for the driver node. You must
	// also specify `instance_pool_id`. Refer to [Instance Pools
	// API](..dev-tools/api/latest/instance-poolshtml) for details.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. This field is optional; if unset, the
	// driver node type is set as the same value as `node_type_id` defined
	// above.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`

	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to clusters running on Google Cloud. If not specified
	// at cluster creation, a set of default values is used.
	GcpAttributes any/* ERROR */ `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of scripts can be
	// specified. The scripts are executed sequentially in the order provided.
	// If `cluster_log_conf` is specified, init script logs are sent to
	// `&lt;destination&gt;/&lt;cluster-id&gt;/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to use for cluster nodes. If
	// `driver_instance_pool_id` is present, `instance_pool_id` is used for
	// worker nodes only. Otherwise, it is used for both the driver node and
	// worker nodes. Refer to [Instance Pools
	// API](..dev-tools/api/latest/instance-poolshtml) for details.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads A
	// list of available node types can be retrieved by using the [List node
	// types](..dev-tools/api/latest/clustershtml#list-node-types) API call.
	// This field is required.
	NodeTypeId string `json:"node_type_id"`
	// If num_workers, number of worker nodes that this cluster must have. A
	// cluster has one Spark driver and num_workers executors for a total of
	// num_workers &#43; 1 Spark nodes. When reading the properties of a cluster,
	// this field reflects the desired number of workers rather than the actual
	// current number of workers. For example, if a cluster is resized from 5 to
	// 10 workers, this field immediately updates to reflect the target size of
	// 10 workers, whereas the workers listed in `spark_info` gradually increase
	// from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// A [cluster policy](..dev-tools/api/latest/policieshtml) ID.
	PolicyId string `json:"policy_id,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. You can also pass in a string of extra JVM
	// options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively. Example Spark confs: `{&#34;spark.speculation&#34;: true,
	// &#34;spark.streaming.ui.retainedBatches&#34;: 5}` or
	// `{&#34;spark.driver.extraJavaOptions&#34;: &#34;-verbose:gc -XX:&#43;PrintGCDetails&#34;}`
	SparkConf map[string]any/* MISSING TYPE */ `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Key-value pair of the form (X,Y) are exported
	// as is (for example, `export X=&#39;Y&#39;`) while launching the driver and
	// workers. To specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// following example. This ensures that all default databricks managed
	// environmental variables are included as well. Example Spark environment
	// variables: `{&#34;SPARK_WORKER_MEMORY&#34;: &#34;28000m&#34;, &#34;SPARK_LOCAL_DIRS&#34;:
	// &#34;/local_disk0&#34;}` or `{&#34;SPARK_DAEMON_JAVA_OPTS&#34;: &#34;$SPARK_DAEMON_JAVA_OPTS
	// -Dspark.shuffle.service.enabled=true&#34;}`
	SparkEnvVars map[string]any/* MISSING TYPE */ `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster. A list of available Spark versions can
	// be retrieved by using the [Runtime
	// versions](..dev-tools/api/latest/clustershtml#runtime-versions) API call.
	// This field is required.
	SparkVersion string `json:"spark_version"`

	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](..notebooks/notebook-workflowshtml#notebook-workflows-exit).
	// jobs restricts this API to return the first 5 MB of the value. For a
	// larger result, your job can store the results in a cloud storage service.
	// This field is absent if `dbutils.notebook.exit()` was never called.
	Result string `json:"result,omitempty"`
	// Whether or not the result was truncated.
	Truncated bool `json:"truncated,omitempty"`
}

type NotebookTask struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to
	// [`run-now`](..dev-tools/api/latest/jobshtml#operation/JobsRunNow) with
	// parameters specified, the two parameters maps are merged. If the same key
	// is specified in `base_parameters` and in `run-now`, the value from
	// `run-now` is used. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. If the notebook takes a parameter that is not
	// specified in the job?s `base_parameters` or the `run-now` override
	// parameters, the default value from the notebook is used. Retrieve these
	// parameters in a notebook using
	// [dbutils.widgets.get](..dev-tools/databricks-utilshtml#dbutils-widgets).
	BaseParameters map[string]any/* MISSING TYPE */ `json:"base_parameters,omitempty"`
	// The path of the notebook to be run in the jobs workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath string `json:"notebook_path"`
}

type PipelineTask struct {
	// If true, a full refresh will be triggered on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// The full name of the pipeline task to execute.
	PipelineId string `json:"pipeline_id,omitempty"`
}

type PythonPyPiLibrary struct {
	// The name of the PyPI package to install. An optional exact version
	// specification is also supported. Examples: `simplejson` and
	// `simplejson==3.8.0`. This field is required.
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo string `json:"repo,omitempty"`
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint string `json:"entry_point,omitempty"`
	// Command-line parameters passed to Python wheel task in the form of
	// `[&#34;--name=task&#34;, &#34;--data=dbfs:/path/to/data.json&#34;]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters any/* MISSING TYPE */ `json:"named_parameters,omitempty"`
	// Name of the package to execute
	PackageName string `json:"package_name,omitempty"`
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters []string `json:"parameters,omitempty"`
}

type RCranLibrary struct {
	// The name of the CRAN package to install. This field is required.
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo string `json:"repo,omitempty"`
}

type RepairHistoryItem struct {
	// The end time of the (repaired) run.
	EndTime int64 `json:"end_time,omitempty"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id int64 `json:"id,omitempty"`
	// The start time of the (repaired) run.
	StartTime int64 `json:"start_time,omitempty"`

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

type RepairRun struct {
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `&#34;jar_params&#34;: [&#34;john doe&#34;, &#34;35&#34;]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{&#34;jar_params&#34;:[&#34;john doe&#34;,&#34;35&#34;]}`) cannot
	// exceed 10,000 bytes. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs.
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId int64 `json:"latest_repair_id,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `&#34;notebook_params&#34;: {&#34;name&#34;: &#34;john doe&#34;, &#34;age&#34;: &#34;35&#34;}`. The map is passed
	// to the notebook and is accessible through the
	// [dbutils.widgets.get](..dev-tools/databricks-utilshtml#dbutils-widgets)
	// function. If not specified upon `run-now`, the triggered run uses the
	// job?s base parameters. notebook_params cannot be specified in conjunction
	// with jar_params. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. The JSON representation of this field (for
	// example `{&#34;notebook_params&#34;:{&#34;name&#34;:&#34;john doe&#34;,&#34;age&#34;:&#34;35&#34;}}`) cannot
	// exceed 10,000 bytes.
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *RepairRunPipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `&#34;python_named_params&#34;: {&#34;name&#34;: &#34;task&#34;, &#34;data&#34;:
	// &#34;dbfs:/path/to/data.json&#34;}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `&#34;python_params&#34;: [&#34;john doe&#34;, &#34;35&#34;]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{&#34;python_params&#34;:[&#34;john
	// doe&#34;,&#34;35&#34;]}`) cannot exceed 10,000 bytes. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. Important These parameters accept only Latin
	// characters (ASCII character set). Using non-ASCII characters returns an
	// error. Examples of invalid, non-ASCII characters are Chinese, Japanese
	// kanjis, and emojis.
	PythonParams []string `json:"python_params,omitempty"`
	// The task keys of the task runs to repair.
	RerunTasks []string `json:"rerun_tasks,omitempty"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId int64 `json:"run_id,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `&#34;spark_submit_params&#34;: [&#34;--class&#34;,
	// &#34;org.apache.spark.examples.SparkPi&#34;]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{&#34;python_params&#34;:[&#34;john doe&#34;,&#34;35&#34;]}`) cannot exceed 10,000 bytes. Use
	// [Task parameter variables](..jobshtml#parameter-variables) to set
	// parameters containing information about job runs. Important These
	// parameters accept only Latin characters (ASCII character set). Using
	// non-ASCII characters returns an error. Examples of invalid, non-ASCII
	// characters are Chinese, Japanese kanjis, and emojis.
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
}

type RepairRunPipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings of the job. These settings completely replace the old
	// settings. Changes to the field `JobSettings.timeout_seconds` are applied
	// to active runs. Changes to other fields are applied to future runs only.
	NewSettings *JobSettings `json:"new_settings,omitempty"`
}

type Run struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \&gt; 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt?s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The total duration of the run is the sum of the
	// setup_duration, the execution_duration, and the cleanup_duration.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`
	// A snapshot of the job?s cluster specification when this run was created.
	ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
	// The creator user name. This field won?t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job&#39;s notebook tasks.
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
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum allowed length is 4096 bytes in
	// UTF-8 encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`

	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// The time it took to set up the cluster in milliseconds. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short.
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
	// The type of trigger that fired this run.
	Trigger TriggerType `json:"trigger,omitempty"`
}

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
type RunLifeCycleState string

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
const RunLifeCycleStateInternalError RunLifeCycleState = `INTERNAL_ERROR`

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
const RunLifeCycleStatePending RunLifeCycleState = `PENDING`

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
const RunLifeCycleStateRunning RunLifeCycleState = `RUNNING`

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
const RunLifeCycleStateSkipped RunLifeCycleState = `SKIPPED`

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
const RunLifeCycleStateTerminated RunLifeCycleState = `TERMINATED`

// * `PENDING`: The run has been triggered. If there is not already an active
// run of the same job, the cluster and execution context are being prepared. If
// there is already an active run of the same job, the run immediately
// transitions into the `SKIPPED` state without preparing any resources. *
// `RUNNING`: The task of this run is being executed. * `TERMINATING`: The task
// of this run has completed, and the cluster and execution context are being
// cleaned up. * `TERMINATED`: The task of this run has completed, and the
// cluster and execution context have been cleaned up. This state is terminal. *
// `SKIPPED`: This run was aborted because a previous run of the same job was
// already active. This state is terminal. * `INTERNAL_ERROR`: An exceptional
// state that indicates a failure in the Jobs service, such as network failure
// over a long period. If a run on a new cluster ends in the `INTERNAL_ERROR`
// state, the Jobs service terminates the cluster as soon as possible. This
// state is terminal.
const RunLifeCycleStateTerminating RunLifeCycleState = `TERMINATING`

type RunNow struct {
	// An optional token to guarantee the idempotency of job run requests. If a
	// run with the provided token already exists, the request does not create a
	// new run but returns the ID of the existing run instead. If a run with the
	// provided token is deleted, an error is returned. If you specify the
	// idempotency token, upon failure you can retry until the request succeeds.
	// Databricks guarantees that exactly one run is launched with that
	// idempotency token. This token must have at most 64 characters. For more
	// information, see [How to ensure idempotency for
	// jobs](https://kb.databricks.com/jobs/jobs-idempotency.html).
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `&#34;jar_params&#34;: [&#34;john doe&#34;, &#34;35&#34;]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{&#34;jar_params&#34;:[&#34;john doe&#34;,&#34;35&#34;]}`) cannot
	// exceed 10,000 bytes. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs.
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the job to be executed
	JobId int64 `json:"job_id,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `&#34;notebook_params&#34;: {&#34;name&#34;: &#34;john doe&#34;, &#34;age&#34;: &#34;35&#34;}`. The map is passed
	// to the notebook and is accessible through the
	// [dbutils.widgets.get](..dev-tools/databricks-utilshtml#dbutils-widgets)
	// function. If not specified upon `run-now`, the triggered run uses the
	// job?s base parameters. notebook_params cannot be specified in conjunction
	// with jar_params. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. The JSON representation of this field (for
	// example `{&#34;notebook_params&#34;:{&#34;name&#34;:&#34;john doe&#34;,&#34;age&#34;:&#34;35&#34;}}`) cannot
	// exceed 10,000 bytes.
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *RunNowPipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `&#34;python_named_params&#34;: {&#34;name&#34;: &#34;task&#34;, &#34;data&#34;:
	// &#34;dbfs:/path/to/data.json&#34;}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `&#34;python_params&#34;: [&#34;john doe&#34;, &#34;35&#34;]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{&#34;python_params&#34;:[&#34;john
	// doe&#34;,&#34;35&#34;]}`) cannot exceed 10,000 bytes. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. Important These parameters accept only Latin
	// characters (ASCII character set). Using non-ASCII characters returns an
	// error. Examples of invalid, non-ASCII characters are Chinese, Japanese
	// kanjis, and emojis.
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `&#34;spark_submit_params&#34;: [&#34;--class&#34;,
	// &#34;org.apache.spark.examples.SparkPi&#34;]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{&#34;python_params&#34;:[&#34;john doe&#34;,&#34;35&#34;]}`) cannot exceed 10,000 bytes. Use
	// [Task parameter variables](..jobshtml#parameter-variables) to set
	// parameters containing information about job runs. Important These
	// parameters accept only Latin characters (ASCII character set). Using
	// non-ASCII characters returns an error. Examples of invalid, non-ASCII
	// characters are Chinese, Japanese kanjis, and emojis.
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
}

type RunNowPipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
}

type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// The globally unique ID of the newly triggered run.
	RunId int64 `json:"run_id,omitempty"`
}

type RunOutput struct {
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error string `json:"error,omitempty"`
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace string `json:"error_trace,omitempty"`
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as
	// [SparkJarTask](..dev-tools/api/latest/jobshtml#/components/schemas/SparkJarTask),
	// [SparkPythonTask](..dev-tools/api/latest/jobshtml#/components/schemas/SparkPythonTask,
	// [PythonWheelTask](..dev-tools/api/latest/jobshtml#/components/schemas/PythonWheelTask.
	// It&#39;s not supported for the
	// [NotebookTask](..dev-tools/api/latest/jobshtml#/components/schemas/NotebookTask,
	// [PipelineTask](..dev-tools/api/latest/jobshtml#/components/schemas/PipelineTask,
	// or
	// [SparkSubmitTask](..dev-tools/api/latest/jobshtml#/components/schemas/SparkSubmitTask.
	// jobs restricts this API to return the last 5 MB of these logs.
	Logs string `json:"logs,omitempty"`
	// Whether the logs are truncated.
	LogsTruncated bool `json:"logs_truncated,omitempty"`
	// All details of the run except for its output.
	Metadata *Run `json:"metadata,omitempty"`
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. jobs restricts this API to
	// return the first 5 MB of the output. To return a larger result, use the
	// [ClusterLogConf](..dev-tools/api/latest/clustershtml#clusterlogconf)
	// field to configure log storage for the job cluster.
	NotebookOutput *NotebookOutput `json:"notebook_output,omitempty"`
}

type RunParameters struct {
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `&#34;jar_params&#34;: [&#34;john doe&#34;, &#34;35&#34;]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{&#34;jar_params&#34;:[&#34;john doe&#34;,&#34;35&#34;]}`) cannot
	// exceed 10,000 bytes. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs.
	JarParams []string `json:"jar_params,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `&#34;notebook_params&#34;: {&#34;name&#34;: &#34;john doe&#34;, &#34;age&#34;: &#34;35&#34;}`. The map is passed
	// to the notebook and is accessible through the
	// [dbutils.widgets.get](..dev-tools/databricks-utilshtml#dbutils-widgets)
	// function. If not specified upon `run-now`, the triggered run uses the
	// job?s base parameters. notebook_params cannot be specified in conjunction
	// with jar_params. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. The JSON representation of this field (for
	// example `{&#34;notebook_params&#34;:{&#34;name&#34;:&#34;john doe&#34;,&#34;age&#34;:&#34;35&#34;}}`) cannot
	// exceed 10,000 bytes.
	NotebookParams map[string]string `json:"notebook_params,omitempty"`

	PipelineParams *RunParametersPipelineParams `json:"pipeline_params,omitempty"`
	// A map from keys to values for jobs with Python wheel task, for example
	// `&#34;python_named_params&#34;: {&#34;name&#34;: &#34;task&#34;, &#34;data&#34;:
	// &#34;dbfs:/path/to/data.json&#34;}`.
	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `&#34;python_params&#34;: [&#34;john doe&#34;, &#34;35&#34;]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{&#34;python_params&#34;:[&#34;john
	// doe&#34;,&#34;35&#34;]}`) cannot exceed 10,000 bytes. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs. Important These parameters accept only Latin
	// characters (ASCII character set). Using non-ASCII characters returns an
	// error. Examples of invalid, non-ASCII characters are Chinese, Japanese
	// kanjis, and emojis.
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `&#34;spark_submit_params&#34;: [&#34;--class&#34;,
	// &#34;org.apache.spark.examples.SparkPi&#34;]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{&#34;python_params&#34;:[&#34;john doe&#34;,&#34;35&#34;]}`) cannot exceed 10,000 bytes. Use
	// [Task parameter variables](..jobshtml#parameter-variables) to set
	// parameters containing information about job runs. Important These
	// parameters accept only Latin characters (ASCII character set). Using
	// non-ASCII characters returns an error. Examples of invalid, non-ASCII
	// characters are Chinese, Japanese kanjis, and emojis.
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
}

type RunParametersPipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
}

// * `SUCCESS`: The task completed successfully. * `FAILED`: The task completed
// with an error. * `TIMEDOUT`: The run was stopped after reaching the timeout.
// * `CANCELED`: The run was canceled at user request.
type RunResultState string

// * `SUCCESS`: The task completed successfully. * `FAILED`: The task completed
// with an error. * `TIMEDOUT`: The run was stopped after reaching the timeout.
// * `CANCELED`: The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

// * `SUCCESS`: The task completed successfully. * `FAILED`: The task completed
// with an error. * `TIMEDOUT`: The run was stopped after reaching the timeout.
// * `CANCELED`: The run was canceled at user request.
const RunResultStateFailed RunResultState = `FAILED`

// * `SUCCESS`: The task completed successfully. * `FAILED`: The task completed
// with an error. * `TIMEDOUT`: The run was stopped after reaching the timeout.
// * `CANCELED`: The run was canceled at user request.
const RunResultStateSuccess RunResultState = `SUCCESS`

// * `SUCCESS`: The task completed successfully. * `FAILED`: The task completed
// with an error. * `TIMEDOUT`: The run was stopped after reaching the timeout.
// * `CANCELED`: The run was canceled at user request.
const RunResultStateTimedout RunResultState = `TIMEDOUT`

// The result and lifecycle state of the run.
type RunState struct {
	// A description of a run?s current location in the run lifecycle. This
	// field is always available in the response.
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty"`

	ResultState RunResultState `json:"result_state,omitempty"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage string `json:"state_message,omitempty"`
	// Whether a run was canceled manually by a user or by the scheduler because
	// the run timed out.
	UserCancelledOrTimedout bool `json:"user_cancelled_or_timedout,omitempty"`
}

type RunSubmitTaskSettings struct {
	DependsOn []TaskDependenciesItem `json:"depends_on,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this task. When running tasks on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the task. The default value is an empty list.
	Libraries []Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *NewCluster `json:"new_cluster,omitempty"`
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
	// spark submit script.
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`

	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. The default
	// behavior is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

type RunTask struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0\. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` \&gt; 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt?s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The total duration of the run is the sum of the
	// setup_duration, the execution_duration, and the cleanup_duration.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance `json:"cluster_instance,omitempty"`

	DependsOn []TaskDependenciesItem `json:"depends_on,omitempty"`

	Description string `json:"description,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs of this job. When running jobs on an existing cluster, you may
	// need to manually restart the cluster if it stops responding. We suggest
	// running jobs on new clusters for greater reliability.
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job&#39;s notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional list of libraries to be installed on the cluster that
	// executes the job. The default value is an empty list.
	Libraries []Library `json:"libraries,omitempty"`
	// If new_cluster, a description of a cluster that is created for each run.
	NewCluster *NewCluster `json:"new_cluster,omitempty"`
	// If notebook_task, indicates that this job must run a notebook. This field
	// may not be specified in conjunction with spark_jar_task.
	NotebookTask *NotebookTask `json:"notebook_task,omitempty"`
	// If pipeline_task, indicates that this job must execute a Pipeline.
	PipelineTask *PipelineTask `json:"pipeline_task,omitempty"`
	// If python_wheel_task, indicates that this job must execute a PythonWheel.
	PythonWheelTask *PythonWheelTask `json:"python_wheel_task,omitempty"`
	// The ID of the task run.
	RunId int64 `json:"run_id,omitempty"`
	// The time it took to set up the cluster in milliseconds. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// If spark_jar_task, indicates that this job must run a JAR.
	SparkJarTask *SparkJarTask `json:"spark_jar_task,omitempty"`
	// If spark_python_task, indicates that this job must run a Python file.
	SparkPythonTask *SparkPythonTask `json:"spark_python_task,omitempty"`
	// If spark_submit_task, indicates that this job must be launched by the
	// spark submit script.
	SparkSubmitTask *SparkSubmitTask `json:"spark_submit_task,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// The result and lifecycle states of the run.
	State *RunState `json:"state,omitempty"`

	TaskKey string `json:"task_key,omitempty"`
}

// The type of the run. * `JOB_RUN` \- Normal job run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow). * `WORKFLOW_RUN`
// \- Workflow run. A run created with
// [dbutils.notebook.run](..dev-tools/databricks-utilshtml#dbutils-workflow). *
// `SUBMIT_RUN` \- Submit run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow).
type RunType string

// The type of the run. * `JOB_RUN` \- Normal job run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow). * `WORKFLOW_RUN`
// \- Workflow run. A run created with
// [dbutils.notebook.run](..dev-tools/databricks-utilshtml#dbutils-workflow). *
// `SUBMIT_RUN` \- Submit run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow).
const RunTypeJobRun RunType = `JOB_RUN`

// The type of the run. * `JOB_RUN` \- Normal job run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow). * `WORKFLOW_RUN`
// \- Workflow run. A run created with
// [dbutils.notebook.run](..dev-tools/databricks-utilshtml#dbutils-workflow). *
// `SUBMIT_RUN` \- Submit run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow).
const RunTypeSubmitRun RunType = `SUBMIT_RUN`

// The type of the run. * `JOB_RUN` \- Normal job run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow). * `WORKFLOW_RUN`
// \- Workflow run. A run created with
// [dbutils.notebook.run](..dev-tools/databricks-utilshtml#dbutils-workflow). *
// `SUBMIT_RUN` \- Submit run. A run created with [Run
// now](..dev-tools/api/latest/jobshtml#operation/JobsRunNow).
const RunTypeWorkflowRun RunType = `WORKFLOW_RUN`

// An arbitrary object where the object key is a configuration propery name and
// the value is a configuration property value.
type SparkConfPair map[string]any /* MISSING TYPE */

// An arbitrary object where the object key is an environment variable name and
// the value is an environment variable value.
type SparkEnvPair map[string]any /* MISSING TYPE */

type SparkJarTask struct {
	// Deprecated since 04/2016\. Provide a `jar` through the `libraries` field
	// instead. For an example, see
	// [Create](..dev-tools/api/latest/jobshtml#operation/JobsCreate).
	JarUri string `json:"jar_uri,omitempty"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library. The code
	// must use `SparkContext.getOrCreate` to obtain a Spark context; otherwise,
	// runs of the job fail.
	MainClassName string `json:"main_class_name,omitempty"`
	// Parameters passed to the main method. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs.
	Parameters []string `json:"parameters,omitempty"`
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs.
	Parameters []string `json:"parameters,omitempty"`

	PythonFile string `json:"python_file"`
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit. Use [Task parameter
	// variables](..jobshtml#parameter-variables) to set parameters containing
	// information about job runs.
	Parameters []string `json:"parameters,omitempty"`
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList any/* MISSING TYPE */ `json:"access_control_list,omitempty"`
	// An optional specification for a remote repository containing the
	// notebooks used by this job&#39;s notebook tasks.
	GitSource *GitSource `json:"git_source,omitempty"`
	// An optional token that can be used to guarantee the idempotency of job
	// run requests. If a run with the provided token already exists, the
	// request does not create a new run but returns the ID of the existing run
	// instead. If a run with the provided token is deleted, an error is
	// returned. If you specify the idempotency token, upon failure you can
	// retry until the request succeeds. Databricks guarantees that exactly one
	// run is launched with that idempotency token. This token must have at most
	// 64 characters. For more information, see [How to ensure idempotency for
	// jobs](https://kb.databricks.com/jobs/jobs-idempotency.html).
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// An optional name for the run. The default value is `Untitled`.
	RunName string `json:"run_name,omitempty"`

	Tasks []RunSubmitTaskSettings `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. The default behavior
	// is to have no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
}

type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId int64 `json:"run_id,omitempty"`
}

// An optional array of objects specifying the dependency graph of the task. All
// tasks specified in this field must complete successfully before executing
// this task. The key is `task_key`, and the value is the name assigned to the
// dependent task. This field is required when a job consists of more than one
// task.
type TaskDependencies []TaskDependenciesItem

type TaskDependenciesItem struct {
	TaskKey string `json:"task_key,omitempty"`
}

// An optional description for this task. The maximum length is 4096 bytes.

// A unique name for the task. This field is used to refer to this task from
// other tasks. This field is required and must be unique within its parent job.
// On Update or Reset, this field is used to reference the tasks to be updated
// or reset. The maximum length is 100 characters.

// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures.
type TriggerType string

// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures.
const TriggerTypeOneTime TriggerType = `ONE_TIME`

// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures.
const TriggerTypePeriodic TriggerType = `PERIODIC`

// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures.
const TriggerTypeRetry TriggerType = `RETRY`

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported. This field is optional.
	FieldsToRemove []string `json:"fields_to_remove,omitempty"`
	// The canonical identifier of the job to update. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings for the job. Any top-level fields specified in
	// `new_settings` are completely replaced. Partially updating nested fields
	// is not supported. Changes to the field `JobSettings.timeout_seconds` are
	// applied to active runs. Changes to other fields are applied to future
	// runs only.
	NewSettings *JobSettings `json:"new_settings,omitempty"`
}

type ViewItem struct {
	// Content of the view.
	Content string `json:"content,omitempty"`
	// Name of the view item. In the case of code view, it would be the
	// notebook?s name. In the case of dashboard view, it would be the
	// dashboard?s name.
	Name string `json:"name,omitempty"`
	// Type of the view item.
	Type ViewType `json:"type,omitempty"`
}

// * `NOTEBOOK`: Notebook view item. * `DASHBOARD`: Dashboard view item.
type ViewType string

// * `NOTEBOOK`: Notebook view item. * `DASHBOARD`: Dashboard view item.
const ViewTypeDashboard ViewType = `DASHBOARD`

// * `NOTEBOOK`: Notebook view item. * `DASHBOARD`: Dashboard view item.
const ViewTypeNotebook ViewType = `NOTEBOOK`

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
type ViewsToExport string

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
const ViewsToExportAll ViewsToExport = `ALL`

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
const ViewsToExportCode ViewsToExport = `CODE`

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
const ViewsToExportDashboards ViewsToExport = `DASHBOARDS`

type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId int64 `json:"job_id,omitempty"`
}

type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 ` url:"run_id,omitempty"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport ` url:"views_to_export,omitempty"`
}

type GetRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId int64 ` url:"job_id,omitempty"`
}

type GetRunOutputRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 ` url:"run_id,omitempty"`
}

type GetRunRequest struct {
	// Whether to include the repair history in the response.
	IncludeHistory bool ` url:"include_history,omitempty"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId int64 ` url:"run_id,omitempty"`
}

type ListRequest struct {
	// Whether to include task and cluster details in the response.
	ExpandTasks bool ` url:"expand_tasks,omitempty"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 25. The default value is 20.
	Limit int ` url:"limit,omitempty"`
	// The offset of the first job to return, relative to the most recently
	// created job.
	Offset int ` url:"offset,omitempty"`
}

type ListResponse struct {
	HasMore bool `json:"has_more,omitempty"`
	// The list of jobs.
	Jobs []Job `json:"jobs,omitempty"`
}

type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `PENDING`, `RUNNING`, or `TERMINATING`. This field cannot be
	// `true` when completed_only is `true`.
	ActiveOnly bool ` url:"active_only,omitempty"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly bool ` url:"completed_only,omitempty"`
	// Whether to include task and cluster details in the response.
	ExpandTasks bool ` url:"expand_tasks,omitempty"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId int64 ` url:"job_id,omitempty"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25\. The default value is 25\. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit int ` url:"limit,omitempty"`
	// The offset of the first run to return, relative to the most recent run.
	Offset int ` url:"offset,omitempty"`
	// The type of runs to return. For a description of run types, see
	// [Run](..dev-tools/api/latest/jobshtml#operation/JobsRunsGet).
	RunType ListRunsRunType ` url:"run_type,omitempty"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom int ` url:"start_time_from,omitempty"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo int ` url:"start_time_to,omitempty"`
}

type ListRunsRunType string

const ListRunsRunTypeJobRun ListRunsRunType = `JOB_RUN`

const ListRunsRunTypeSubmitRun ListRunsRunType = `SUBMIT_RUN`

const ListRunsRunTypeWorkflowRun ListRunsRunType = `WORKFLOW_RUN`

type RepairRunResponse struct {
	// The ID of the repair.
	RepairId int64 `json:"repair_id,omitempty"`
}
