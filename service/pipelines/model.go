// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/libraries"
)

// all definitions in this file are in alphabetical order

type CreatePipeline struct {
	// If false, deployment will fail if name conflicts with that of another
	// pipeline.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// Catalog in UC to add tables to. If target is specified, tables in this
	// pipeline will be published to a "target" schema inside catalog (i.e.
	// <catalog>.<target>.<table>).
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`

	DryRun bool `json:"dry_run,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
}

type CreatePipelineResponse struct {
	// Only returned when dry_run is true
	EffectiveSettings *PipelineSpec `json:"effective_settings,omitempty"`
	// Only returned when dry_run is false
	PipelineId string `json:"pipeline_id,omitempty"`
}

type CronTrigger struct {
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`

	TimezoneId string `json:"timezone_id,omitempty"`
}

// Delete a pipeline
type DeletePipeline struct {
	PipelineId string `json:"-" url:"-"`
}

type EditPipeline struct {
	// If false, deployment will fail if name has changed and conflicts the name
	// of another pipeline.
	AllowDuplicateNames bool `json:"allow_duplicate_names,omitempty"`
	// Catalog in UC to add tables to. If target is specified, tables in this
	// pipeline will be published to a "target" schema inside catalog (i.e.
	// <catalog>.<target>.<table>).
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// If present, the last-modified time of the pipeline settings before the
	// edit. If the settings were modified after that time, then the request
	// will fail with a conflict.
	ExpectedLastModified int64 `json:"expected_last_modified,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// Unique identifier for this pipeline.
	PipelineId string `json:"pipeline_id,omitempty" url:"-"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
}

type Filters struct {
	// Paths to exclude.
	Exclude []string `json:"exclude,omitempty"`
	// Paths to include.
	Include []string `json:"include,omitempty"`
}

// Get a pipeline
type GetPipeline struct {
	PipelineId string `json:"-" url:"-"`
}

type GetPipelineResponse struct {
	// An optional message detailing the cause of the pipeline state.
	Cause string `json:"cause,omitempty"`
	// The ID of the cluster that the pipeline is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The health of a pipeline.
	Health GetPipelineResponseHealth `json:"health,omitempty"`
	// The last time the pipeline settings were modified or created.
	LastModified int64 `json:"last_modified,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// A human friendly identifier for the pipeline, taken from the `spec`.
	Name string `json:"name,omitempty"`
	// The ID of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// Username of the user that the pipeline will run on behalf of.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline specification. This field is not returned when called by
	// `ListPipelines`.
	Spec *PipelineSpec `json:"spec,omitempty"`
	// The pipeline state.
	State PipelineState `json:"state,omitempty"`
}

// The health of a pipeline.
type GetPipelineResponseHealth string

const GetPipelineResponseHealthHealthy GetPipelineResponseHealth = `HEALTHY`

const GetPipelineResponseHealthUnhealthy GetPipelineResponseHealth = `UNHEALTHY`

// Get a pipeline update
type GetUpdate struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" url:"-"`
	// The ID of the update.
	UpdateId string `json:"-" url:"-"`
}

type GetUpdateResponse struct {
	// The current update info.
	Update *UpdateInfo `json:"update,omitempty"`
}

// List pipelines
type ListPipelines struct {
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
}

type ListPipelinesResponse struct {
	// If present, a token to fetch the next page of events.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The list of events matching the request criteria.
	Statuses []PipelineStateInfo `json:"statuses,omitempty"`
}

// List pipeline updates
type ListUpdates struct {
	// Max number of entries to return in a single page.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Page token returned by previous call
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The pipeline to return updates for.
	PipelineId string `json:"-" url:"-"`
	// If present, returns updates until and including this update_id.
	UntilUpdateId string `json:"-" url:"until_update_id,omitempty"`
}

type ListUpdatesResponse struct {
	// If present, then there are more results, and this a token to be used in a
	// subsequent request to fetch the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// If present, then this token can be used in a subsequent request to fetch
	// the previous page.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	Updates []UpdateInfo `json:"updates,omitempty"`
}

type NotebookLibrary struct {
	// The absolute path of the notebook.
	Path string `json:"path,omitempty"`
}

type PipelineCluster struct {
	// Note: This field won't be persisted. Only API users will check this
	// field.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *clusters.AutoScale `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *clusters.AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *clusters.AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every ``5 mins``. The
	// destination of driver logs is ``$destination/$clusterId/driver``, while
	// the destination of executor logs is ``$destination/$clusterId/executor``.
	ClusterLogConf *clusters.ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to ``default_tags``. Notes:
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
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *clusters.GcpAttributes `json:"gcp_attributes,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// Cluster label
	Label string `json:"label,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
	// + 1 Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in ``spark_info`` will gradually
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
	// (X,Y) will be exported as is (i.e., ``export X='Y'``) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of ``SPARK_DAEMON_JAVA_OPTS``, we
	// recommend appending them to ``$SPARK_DAEMON_JAVA_OPTS`` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: ``{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}`` or ``{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}``
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name ``ubuntu`` on port ``2200``. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
}

type PipelineLibrary struct {
	// URI of the jar to be installed. Currently only DBFS and S3 URIs are
	// supported. For example: ``{ "jar": "dbfs:/mnt/databricks/library.jar" }``
	// or ``{ "jar": "s3://my-bucket/library.jar" }``. If S3 is used, please
	// make sure the cluster has read access on the library. You may need to
	// launch the cluster with an IAM role to access the S3 URI.
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed. For example: ``{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }``
	Maven *libraries.MavenLibrary `json:"maven,omitempty"`
	// The path to a notebook that defines a pipeline and is stored in the
	// Databricks workspace. For example: ``{ "notebook" : { "path" :
	// "/my-pipeline-notebook-path" } }``. Currently, only Scala notebooks are
	// supported, and pipelines must be defined in a package cell.
	Notebook *NotebookLibrary `json:"notebook,omitempty"`
	// URI of the wheel to be installed. For example: ``{ "whl": "dbfs:/my/whl"
	// }`` or ``{ "whl": "s3://my-bucket/whl" }``. If S3 is used, please make
	// sure the cluster has read access on the library. You may need to launch
	// the cluster with an IAM role to access the S3 URI.
	Whl string `json:"whl,omitempty"`
}

type PipelineSpec struct {
	// Catalog in UC to add tables to. If target is specified, tables in this
	// pipeline will be published to a "target" schema inside catalog (i.e.
	// <catalog>.<target>.<table>).
	Catalog string `json:"catalog,omitempty"`
	// DLT Release Channel that specifies which version to use.
	Channel string `json:"channel,omitempty"`
	// Cluster settings for this pipeline deployment.
	Clusters []PipelineCluster `json:"clusters,omitempty"`
	// String-String configuration for this pipeline execution.
	Configuration map[string]string `json:"configuration,omitempty"`
	// Whether the pipeline is continuous or triggered. This replaces `trigger`.
	Continuous bool `json:"continuous,omitempty"`
	// Whether the pipeline is in Development mode. Defaults to false.
	Development bool `json:"development,omitempty"`
	// Pipeline product edition.
	Edition string `json:"edition,omitempty"`
	// Filters on which Pipeline packages to include in the deployed graph.
	Filters *Filters `json:"filters,omitempty"`
	// Unique identifier for this pipeline.
	Id string `json:"id,omitempty"`
	// Libraries or code needed by this deployment.
	Libraries []PipelineLibrary `json:"libraries,omitempty"`
	// Friendly identifier for this pipeline.
	Name string `json:"name,omitempty"`
	// Whether Photon is enabled for this pipeline.
	Photon bool `json:"photon,omitempty"`
	// DBFS root directory for storing checkpoints and tables.
	Storage string `json:"storage,omitempty"`
	// Target schema (database) to add tables in this pipeline to.
	Target string `json:"target,omitempty"`
	// Which pipeline trigger to use. Deprecated: Use `continuous` instead.
	Trigger *PipelineTrigger `json:"trigger,omitempty"`
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

type PipelineStateInfo struct {
	// The unique identifier of the cluster running the pipeline.
	ClusterId string `json:"cluster_id,omitempty"`
	// The username of the pipeline creator.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Status of the latest updates for the pipeline. Ordered with the newest
	// update first.
	LatestUpdates []UpdateStateInfo `json:"latest_updates,omitempty"`
	// The user-friendly name of the pipeline.
	Name string `json:"name,omitempty"`
	// The unique identifier of the pipeline.
	PipelineId string `json:"pipeline_id,omitempty"`
	// The username that the pipeline runs as. This is a read only value derived
	// from the pipeline owner.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// The pipeline state.
	State PipelineState `json:"state,omitempty"`
}

type PipelineTrigger struct {
	Cron *CronTrigger `json:"cron,omitempty"`

	Manual any/* MISSING TYPE */ `json:"manual,omitempty"`
}

// Reset a pipeline
type ResetPipeline struct {
	PipelineId string `json:"-" url:"-"`
}

type StartUpdate struct {
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
}

type StartUpdateCause string

const StartUpdateCauseApiCall StartUpdateCause = `API_CALL`

const StartUpdateCauseJobTask StartUpdateCause = `JOB_TASK`

const StartUpdateCauseRetryOnFailure StartUpdateCause = `RETRY_ON_FAILURE`

const StartUpdateCauseSchemaChange StartUpdateCause = `SCHEMA_CHANGE`

const StartUpdateCauseServiceUpgrade StartUpdateCause = `SERVICE_UPGRADE`

const StartUpdateCauseUserAction StartUpdateCause = `USER_ACTION`

type StartUpdateResponse struct {
	UpdateId string `json:"update_id,omitempty"`
}

// Stop a pipeline
type StopPipeline struct {
	PipelineId string `json:"-" url:"-"`
}

type UpdateInfo struct {
	// What triggered this update.
	Cause UpdateInfoCause `json:"cause,omitempty"`
	// The ID of the cluster that the update is running on.
	ClusterId string `json:"cluster_id,omitempty"`
	// The pipeline configuration with system defaults applied where unspecified
	// by the user. Not returned by ListUpdates.
	Config *PipelineSpec `json:"config,omitempty"`
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
}

// What triggered this update.
type UpdateInfoCause string

const UpdateInfoCauseApiCall UpdateInfoCause = `API_CALL`

const UpdateInfoCauseJobTask UpdateInfoCause = `JOB_TASK`

const UpdateInfoCauseRetryOnFailure UpdateInfoCause = `RETRY_ON_FAILURE`

const UpdateInfoCauseSchemaChange UpdateInfoCause = `SCHEMA_CHANGE`

const UpdateInfoCauseServiceUpgrade UpdateInfoCause = `SERVICE_UPGRADE`

const UpdateInfoCauseUserAction UpdateInfoCause = `USER_ACTION`

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

type UpdateStateInfo struct {
	CreationTime string `json:"creation_time,omitempty"`

	State UpdateStateInfoState `json:"state,omitempty"`

	UpdateId string `json:"update_id,omitempty"`
}

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
