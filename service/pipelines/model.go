// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

// all definitions in this file are in alphabetical order

type CreatePipelineRequest struct {
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

type DeletePipelineRequest struct {
	PipelineId string `json:"-" path:"pipeline_id"`
}

type EditPipelineRequest struct {
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

	PipelineId string `json:"-" path:"pipeline_id"`
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

type GetPipelineRequest struct {
	PipelineId string `json:"-" path:"pipeline_id"`
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
	State GetPipelineResponseState `json:"state,omitempty"`
}

// The health of a pipeline.
type GetPipelineResponseHealth string

const GetPipelineResponseHealthHealthy GetPipelineResponseHealth = `HEALTHY`

const GetPipelineResponseHealthUnhealthy GetPipelineResponseHealth = `UNHEALTHY`

// The pipeline state.
type GetPipelineResponseState string

const GetPipelineResponseStateDeleted GetPipelineResponseState = `DELETED`

const GetPipelineResponseStateDeploying GetPipelineResponseState = `DEPLOYING`

const GetPipelineResponseStateFailed GetPipelineResponseState = `FAILED`

const GetPipelineResponseStateIdle GetPipelineResponseState = `IDLE`

const GetPipelineResponseStateRecovering GetPipelineResponseState = `RECOVERING`

const GetPipelineResponseStateResetting GetPipelineResponseState = `RESETTING`

const GetPipelineResponseStateRunning GetPipelineResponseState = `RUNNING`

const GetPipelineResponseStateStarting GetPipelineResponseState = `STARTING`

const GetPipelineResponseStateStopping GetPipelineResponseState = `STOPPING`

type GetUpdateRequest struct {
	// The ID of the pipeline.
	PipelineId string `json:"-" path:"pipeline_id"`
	// The ID of the update.
	UpdateId string `json:"-" path:"update_id"`
}

type GetUpdateResponse struct {
	// The current update info.
	Update *UpdateInfo `json:"update,omitempty"`
}

type ListUpdatesRequest struct {
	// Max number of entries to return in a single page.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Page token returned by previous call
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The pipeline to return updates for.
	PipelineId string `json:"-" path:"pipeline_id"`
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
	Autoscale *PipelinesAutoScale `json:"autoscale,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *PipelinesAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes any/* MISSING TYPE */ `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every ``5 mins``. The
	// destination of driver logs is ``$destination/$clusterId/driver``, while
	// the destination of executor logs is ``$destination/$clusterId/executor``.
	ClusterLogConf *PipelinesClusterLogConf `json:"cluster_log_conf,omitempty"`
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
	GcpAttributes *PipelinesGcpAttributes `json:"gcp_attributes,omitempty"`
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
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// ``spark.driver.extraJavaOptions`` and ``spark.executor.extraJavaOptions``
	// respectively.
	//
	// Example Spark confs: ``{"spark.speculation": true,
	// "spark.streaming.ui.retainedBatches": 5}`` or
	// ``{"spark.driver.extraJavaOptions": "-verbose:gc -XX:+PrintGCDetails"}``
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
	Maven *PipelinesMavenLibrary `json:"maven,omitempty"`
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

type PipelineTrigger struct {
	Cron *CronTrigger `json:"cron,omitempty"`

	Manual any/* MISSING TYPE */ `json:"manual,omitempty"`
}

type PipelinesAutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that ``max_workers`` must be strictly greater than
	// ``min_workers``.
	MaxWorkers int `json:"max_workers,omitempty"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers int `json:"min_workers,omitempty"`
	// The autoscaling mode. This is an additional field available in DLT only.
	// This is used to specify the autoscaling algorithm to be used for this
	// autoscaling cluster. Decision Doc here:
	// https://docs.google.com/document/d/1Eojc1a5raIyApDz_2NYyoltDlimjXzqkTgwFwDQoGnw/edit?usp=sharing
	Mode string `json:"mode,omitempty"`
}

type PipelinesAwsAttributes struct {
	// The first ``first_on_demand`` nodes of the cluster will be placed on
	// on-demand instances. If this value is greater than 0, the cluster driver
	// node in particular will be placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, ``first_on_demand`` nodes will be placed on on-demand
	// instances and the remainder will be placed on ``availability`` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand int `json:"first_on_demand,omitempty"`
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	//
	// If this field is ommitted, we will pull in the default from the conf if
	// it exists.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. If the zone specified is "auto", will try to place cluster
	// in a zone with high availability, and will retry placement in a different
	// AZ if there is not enough capacity. See [[AutoAZHelper.scala]] for more
	// details. The list of available zones as well as the default value can be
	// found by using the `List Zones`_ method.
	ZoneId string `json:"zone_id,omitempty"`
}

type PipelinesClusterLogConf struct {
	// destination needs to be provided. e.g. ``{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }``
	Dbfs *PipelinesDbfsStorageInfo `json:"dbfs,omitempty"`
	// destination and either region or endpoint should also be provided. e.g.
	// ``{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }`` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in ``instance_profile_arn`` has permission to write
	// data to the s3 destination.
	S3 *PipelinesS3StorageInfo `json:"s3,omitempty"`
}

type PipelinesDbfsStorageInfo struct {
	// dbfs destination, e.g. ``dbfs:/my/path``
	Destination string `json:"destination,omitempty"`
}

type PipelinesGcpAttributes struct {
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
}

type PipelinesMavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	Coordinates string `json:"coordinates"`
	// List of dependences to exclude. For example: ``["slf4j:slf4j",
	// "*:hadoop-client"]``.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions []string `json:"exclusions,omitempty"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo string `json:"repo,omitempty"`
}

type PipelinesS3StorageInfo struct {
	// (Optional) Set canned access control list for the logs, e.g.
	// ``bucket-owner-full-control``. If ``canned_cal`` is set, please make sure
	// the cluster iam role has ``s3:PutObjectAcl`` permission on the
	// destination bucket and prefix. The full list of possible canned acl can
	// be found at
	// http://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl.
	// Please also note that by default only the object owner gets full
	// controls. If you are using cross account role for writing data, you may
	// want to set ``bucket-owner-full-control`` to make bucket owner able to
	// read the logs.
	CannedAcl string `json:"canned_acl,omitempty"`
	// S3 destination, e.g. ``s3://my-bucket/some-prefix`` Note that logs will
	// be delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	Destination string `json:"destination,omitempty"`
	// (Optional) Flag to enable server side encryption, ``false`` by default.
	EnableEncryption bool `json:"enable_encryption,omitempty"`
	// (Optional) The encryption type, it could be ``sse-s3`` or ``sse-kms``. It
	// will be used only when encryption is enabled and the default type is
	// ``sse-s3``.
	EncryptionType string `json:"encryption_type,omitempty"`
	// S3 endpoint, e.g. ``https://s3-us-west-2.amazonaws.com``. Either region
	// or endpoint needs to be set. If both are set, endpoint will be used.
	Endpoint string `json:"endpoint,omitempty"`
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to ``sse-kms``.
	KmsKey string `json:"kms_key,omitempty"`
	// S3 region, e.g. ``us-west-2``. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	Region string `json:"region,omitempty"`
}

type ResetPipelineRequest struct {
	PipelineId string `json:"-" path:"pipeline_id"`
}

type StartUpdateRequest struct {
	Cause StartUpdateRequestCause `json:"cause,omitempty"`
	// If true, this update will reset all tables before running.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// A list of tables to update with fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	FullRefreshSelection []string `json:"full_refresh_selection,omitempty"`

	PipelineId string `json:"-" path:"pipeline_id"`
	// A list of tables to update without fullRefresh. If both refresh_selection
	// and full_refresh_selection are empty, this is a full graph update. Full
	// Refresh on a table means that the states of the table will be reset
	// before the refresh.
	RefreshSelection []string `json:"refresh_selection,omitempty"`
}

type StartUpdateRequestCause string

const StartUpdateRequestCauseApiCall StartUpdateRequestCause = `API_CALL`

const StartUpdateRequestCauseJobTask StartUpdateRequestCause = `JOB_TASK`

const StartUpdateRequestCauseRetryOnFailure StartUpdateRequestCause = `RETRY_ON_FAILURE`

const StartUpdateRequestCauseSchemaChange StartUpdateRequestCause = `SCHEMA_CHANGE`

const StartUpdateRequestCauseServiceUpgrade StartUpdateRequestCause = `SERVICE_UPGRADE`

const StartUpdateRequestCauseUserAction StartUpdateRequestCause = `USER_ACTION`

type StartUpdateResponse struct {
	UpdateId string `json:"update_id,omitempty"`
}

type StopPipelineRequest struct {
	PipelineId string `json:"-" path:"pipeline_id"`
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
