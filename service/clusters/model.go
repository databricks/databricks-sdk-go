// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package clusters

// all definitions in this file are in alphabetical order


type AutoScale struct {
    // The maximum number of workers to which the cluster can scale up when
    // overloaded. Note that ``max_workers`` must be strictly greater than
    // ``min_workers``.
    MaxWorkers int `json:"max_workers,omitempty"`
    // The minimum number of workers to which the cluster can scale down when
    // underutilized. It is also the initial number of workers the cluster will
    // have after creation.
    MinWorkers int `json:"min_workers,omitempty"`
}


type AwsAttributes struct {
    // Availability type used for all subsequent nodes past the
    // ``first_on_demand`` ones. Note: If ``first_on_demand`` is zero, this
    // availability type will be used for the entire cluster.
    Availability AwsAttributesAvailability `json:"availability,omitempty"`
    // The number of volumes launched for each instance. Users can choose up to
    // 10 volumes. This feature is only enabled for supported node types. Legacy
    // node types cannot specify custom EBS volumes. For node types with no
    // instance store, at least one EBS volume needs to be specified; otherwise,
    // cluster creation will fail. These EBS volumes will be mounted at
    // ``/ebs0``, ``/ebs1``, and etc. Instance store volumes will be mounted at
    // ``/local_disk0``, ``/local_disk1``, and etc. If EBS volumes are attached,
    // Databricks will configure Spark to use only the EBS volumes for scratch
    // storage because heterogenously sized scratch devices can lead to
    // inefficient disk utilization. If no EBS volumes are attached, Databricks
    // will configure Spark to use instance store volumes. Please note that if
    // EBS volumes are specified, then the Spark configuration
    // ``spark.local.dir`` will be overridden.
    EbsVolumeCount int `json:"ebs_volume_count,omitempty"`
    
    EbsVolumeIops int `json:"ebs_volume_iops,omitempty"`
    // The size of each EBS volume (in GiB) launched for each instance. For
    // general purpose SSD, this value must be within the range 100 - 4096. For
    // throughput optimized HDD, this value must be within the range 500 - 4096.
    EbsVolumeSize int `json:"ebs_volume_size,omitempty"`
    
    EbsVolumeThroughput int `json:"ebs_volume_throughput,omitempty"`
    // The type of EBS volumes that will be launched with this cluster.
    EbsVolumeType AwsAttributesEbsVolumeType `json:"ebs_volume_type,omitempty"`
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
    // added to the Databricks environment by an account administrator. This
    // feature may only be available to certain customer plans. If this field is
    // ommitted, we will pull in the default from the conf if it exists.
    InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
    // The bid price for AWS spot instances, as a percentage of the
    // corresponding instance type&#39;s on-demand price. For example, if this field
    // is set to 50, and the cluster needs a new ``r3.xlarge`` spot instance,
    // then the bid price is half of the price of on-demand ``r3.xlarge``
    // instances. Similarly, if this field is set to 200, the bid price is twice
    // the price of on-demand ``r3.xlarge`` instances. If not specified, the
    // default value is 100. When spot instances are requested for this cluster,
    // only spot instances whose bid price percentage matches this field will be
    // considered. Note that, for safety, we enforce this field to be no more
    // than 10000. The default value and documentation here should be kept
    // consistent with CommonConf.defaultSpotBidPricePercent and
    // CommonConf.maxSpotBidPricePercent.
    SpotBidPricePercent int `json:"spot_bid_price_percent,omitempty"`
    // Identifier for the availability zone/datacenter in which the cluster
    // resides. This string will be of a form like &#34;us-west-2a&#34;. The provided
    // availability zone must be in the same region as the Databricks
    // deployment. For example, &#34;us-west-2a&#34; is not a valid zone id if the
    // Databricks deployment resides in the &#34;us-east-1&#34; region. This is an
    // optional field at cluster creation, and if not specified, a default zone
    // will be used. If the zone specified is &#34;auto&#34;, will try to place cluster
    // in a zone with high availability, and will retry placement in a different
    // AZ if there is not enough capacity. See [[AutoAZHelper.scala]] for more
    // details. The list of available zones as well as the default value can be
    // found by using the `List Zones`_ method.
    ZoneId string `json:"zone_id,omitempty"`
}

// Availability type used for all subsequent nodes past the ``first_on_demand``
// ones. Note: If ``first_on_demand`` is zero, this availability type will be
// used for the entire cluster.

type AwsAttributesAvailability string
const (
	AwsAttributesAvailabilitySpot AwsAttributesAvailability = `SPOT`
	AwsAttributesAvailabilityOnDemand AwsAttributesAvailability = `ON_DEMAND`
	AwsAttributesAvailabilitySpotWithFallback AwsAttributesAvailability = `SPOT_WITH_FALLBACK`
)

// The type of EBS volumes that will be launched with this cluster.

type AwsAttributesEbsVolumeType string
const (
	AwsAttributesEbsVolumeTypeGeneralPurposeSsd AwsAttributesEbsVolumeType = `GENERAL_PURPOSE_SSD`
	AwsAttributesEbsVolumeTypeThroughputOptimizedHdd AwsAttributesEbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`
)


type AwsCpalError struct {
    
    AwsApiErrorCode string `json:"aws_api_error_code,omitempty"`
    
    AwsErrorMessage string `json:"aws_error_message,omitempty"`
    
    AwsInstanceStateReason string `json:"aws_instance_state_reason,omitempty"`
    
    AwsSpotRequestFaultCode string `json:"aws_spot_request_fault_code,omitempty"`
    
    AwsSpotRequestStatus string `json:"aws_spot_request_status,omitempty"`
}


type AzureAttributes struct {
    // Availability type used for all subsequent nodes past the
    // ``first_on_demand`` ones. Note: If ``first_on_demand`` is zero (which
    // only happens on pool clusters), this availability type will be used for
    // the entire cluster.
    Availability AzureAttributesAvailability `json:"availability,omitempty"`
    // The first ``first_on_demand`` nodes of the cluster will be placed on
    // on-demand instances. This value should be greater than 0, to make sure
    // the cluster driver node is placed on an on-demand instance. If this value
    // is greater than or equal to the current cluster size, all nodes will be
    // placed on on-demand instances. If this value is less than the current
    // cluster size, ``first_on_demand`` nodes will be placed on on-demand
    // instances and the remainder will be placed on ``availability`` instances.
    // Note that this value does not affect cluster size and cannot currently be
    // mutated over the lifetime of a cluster.
    FirstOnDemand int `json:"first_on_demand,omitempty"`
    // Defines values necessary to configure and run Azure Log Analytics agent
    LogAnalyticsInfo *LogAnalyticsInfo `json:"log_analytics_info,omitempty"`
    // The max bid price to be used for Azure spot instances. The Max price for
    // the bid cannot be higher than the on-demand price of the instance. If not
    // specified, the default value is -1, which specifies that the instance
    // cannot be evicted on the basis of price, and only on the basis of
    // availability. Further, the value should &gt; 0 or -1.
    SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`
}

// Availability type used for all subsequent nodes past the ``first_on_demand``
// ones. Note: If ``first_on_demand`` is zero (which only happens on pool
// clusters), this availability type will be used for the entire cluster.

type AzureAttributesAvailability string
const (
	AzureAttributesAvailabilitySpotAzure AzureAttributesAvailability = `SPOT_AZURE`
	AzureAttributesAvailabilityOnDemandAzure AzureAttributesAvailability = `ON_DEMAND_AZURE`
	AzureAttributesAvailabilitySpotWithFallbackAzure AzureAttributesAvailability = `SPOT_WITH_FALLBACK_AZURE`
)


type ChangeClusterOwnerRequest struct {
    
    ClusterId string `json:"cluster_id"`
    // New owner of the cluster_id after this RPC.
    OwnerUsername string `json:"owner_username,omitempty"`
}


type ClientsTypes struct {
    // With jobs set, the cluster can be used for jobs
    Jobs bool `json:"jobs,omitempty"`
    // With notebooks set, this cluster can be used for notebooks
    Notebooks bool `json:"notebooks,omitempty"`
}


type ClusterAttributes struct {
    // Automatically terminates the cluster after it is inactive for this time
    // in minutes. If not set, this cluster will not be automatically
    // terminated. If specified, the threshold must be between 10 and 10000
    // minutes. Users can also set this value to 0 to explicitly disable
    // automatic termination.
    AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
    // The configuration for delivering spark logs to a long-term storage
    // destination. Two kinds of destinations (dbfs and s3) are supported. Only
    // one destination can be specified for one cluster. If the conf is given,
    // the logs will be delivered to the destination every ``5 mins``. The
    // destination of driver logs is ``$destination/$clusterId/driver``, while
    // the destination of executor logs is ``$destination/$clusterId/executor``.
    ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
    // Cluster name requested by the user. This doesn&#39;t have to be unique. If
    // not specified at creation, the cluster name will be an empty string.
    ClusterName string `json:"cluster_name,omitempty"`
    // Determines whether the cluster was created by a user through the UI,
    // created by the Databricks Jobs Scheduler, or through an API request. This
    // is the same as cluster_creator, but read only.
    ClusterSource ClusterAttributesClusterSource `json:"cluster_source,omitempty"`
    // Additional tags for cluster resources. Databricks will tag all cluster
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags - Clusters can only reuse cloud resources if the
    // resources&#39; tags are a subset of the cluster tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // The optional ID of the instance pool for the driver of the cluster
    // belongs. The pool cluster uses the instance pool with id
    // (instance_pool_id) if the driver pool is not assigned.
    DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
    // The node type of the Spark driver. Note that this field is optional; if
    // unset, the driver node type will be set as the same value as
    // `node_type_id` defined above. This field, along with node_type_id, should
    // not be set if virtual_cluster_size is set. If both driver_node_type_id,
    // node_type_id, and virtual_cluster_size are specified, driver_node_type_id
    // and node_type_id take precedence.
    DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
    // The key of the spark version running in the dataplane. This is possibly
    // different from the spark_version (index 2). The spark_version is the raw
    // string provided by the user through API or UI, which could map to a
    // different effective_spark_version running in the dataplane, depending on
    // the cluster&#39;s instance type or the runtimeEngine parameter.
    EffectiveSparkVersion string `json:"effective_spark_version,omitempty"`
    // Autoscaling Local Storage: when enabled, this cluster will dynamically
    // acquire additional disk space when its Spark workers are running low on
    // disk space. This feature requires specific AWS permissions to function
    // correctly - refer to the User Guide for more details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Whether to enable LUKS on cluster VMs&#39; local disks
    EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
    // The optional ID of the instance pool to which the cluster belongs.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this cluster. For example, the Spark nodes can
    // be provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call. This field, along
    // with driver_node_type_id, should not be set if virtual_cluster_size is
    // set. If both driver_node_type_id, node_type_id, and virtual_cluster_size
    // are specified, driver_node_type_id and node_type_id take precedence.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // The ID of the cluster policy used to create the cluster if applicable.
    PolicyId string `json:"policy_id,omitempty"`
    // Decides which runtime engine to be use, e.g. Standard vs. Photon. If
    // unspecified, the runtime engine is inferred from spark_version.
    RuntimeEngine ClusterAttributesRuntimeEngine `json:"runtime_engine,omitempty"`
    // An object containing a set of optional, user-specified Spark
    // configuration key-value pairs. Users can also pass in a string of extra
    // JVM options to the driver and the executors via
    // ``spark.driver.extraJavaOptions`` and ``spark.executor.extraJavaOptions``
    // respectively. Example Spark confs: ``{&#34;spark.speculation&#34;: true,
    // &#34;spark.streaming.ui.retainedBatches&#34;: 5}`` or
    // ``{&#34;spark.driver.extraJavaOptions&#34;: &#34;-verbose:gc -XX:&#43;PrintGCDetails&#34;}``
    SparkConf map[string]string `json:"spark_conf,omitempty"`
    // An object containing a set of optional, user-specified environment
    // variable key-value pairs. Please note that key-value pair of the form
    // (X,Y) will be exported as is (i.e., ``export X=&#39;Y&#39;``) while launching the
    // driver and workers. In order to specify an additional set of
    // ``SPARK_DAEMON_JAVA_OPTS``, we recommend appending them to
    // ``$SPARK_DAEMON_JAVA_OPTS`` as shown in the example below. This ensures
    // that all default databricks managed environmental variables are included
    // as well. Example Spark environment variables: ``{&#34;SPARK_WORKER_MEMORY&#34;:
    // &#34;28000m&#34;, &#34;SPARK_LOCAL_DIRS&#34;: &#34;/local_disk0&#34;}`` or
    // ``{&#34;SPARK_DAEMON_JAVA_OPTS&#34;: &#34;$SPARK_DAEMON_JAVA_OPTS
    // -Dspark.shuffle.service.enabled=true&#34;}``
    SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
    // The Spark version of the cluster, e.g. &#34;3.3.x-scala2.11&#34;. A list of
    // available Spark versions can be retrieved by using the
    // :ref:`clusterClusterServicelistSparkVersions` API call. This is the Spark
    // version provided from the user input (API/UI) and may be different from
    // the effective_spark_version, which is the spark version that is actually
    // run in the dataplane. See index 45 for more context.
    SparkVersion string `json:"spark_version,omitempty"`
    // SSH public key contents that will be added to each Spark node in this
    // cluster. The corresponding private keys can be used to login with the
    // user name ``ubuntu`` on port ``2200``. Up to 10 keys can be specified.
    SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
    
    WorkloadType *WorkloadType `json:"workload_type,omitempty"`
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.

type ClusterAttributesClusterSource string
const (
	ClusterAttributesClusterSourceUi ClusterAttributesClusterSource = `UI`
	ClusterAttributesClusterSourceJob ClusterAttributesClusterSource = `JOB`
	ClusterAttributesClusterSourceApi ClusterAttributesClusterSource = `API`
	ClusterAttributesClusterSourceSql ClusterAttributesClusterSource = `SQL`
	ClusterAttributesClusterSourceModels ClusterAttributesClusterSource = `MODELS`
	ClusterAttributesClusterSourcePipeline ClusterAttributesClusterSource = `PIPELINE`
	ClusterAttributesClusterSourcePipelineMaintenance ClusterAttributesClusterSource = `PIPELINE_MAINTENANCE`
)

// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
// unspecified, the runtime engine is inferred from spark_version.

type ClusterAttributesRuntimeEngine string
const (
	ClusterAttributesRuntimeEngineNull ClusterAttributesRuntimeEngine = `NULL`
	ClusterAttributesRuntimeEngineStandard ClusterAttributesRuntimeEngine = `STANDARD`
	ClusterAttributesRuntimeEnginePhoton ClusterAttributesRuntimeEngine = `PHOTON`
)


type ClusterEvent struct {
    
    ClusterId string `json:"cluster_id"`
    
    DataPlaneEventDetails *DataPlaneEventDetails `json:"data_plane_event_details,omitempty"`
    
    Details *EventDetails `json:"details,omitempty"`
    // The timestamp when the event occurred, stored as the number of
    // milliseconds since the Unix epoch. If not provided, this will be assigned
    // by the Timeline service.
    Timestamp int64 `json:"timestamp,omitempty"`
    
    Type ClusterEventType `json:"type,omitempty"`
}



type ClusterEventType string
const (
	ClusterEventTypeCreating ClusterEventType = `CREATING`
	ClusterEventTypeStarting ClusterEventType = `STARTING`
	ClusterEventTypeRestarting ClusterEventType = `RESTARTING`
	ClusterEventTypeTerminating ClusterEventType = `TERMINATING`
	ClusterEventTypeEdited ClusterEventType = `EDITED`
	ClusterEventTypeRunning ClusterEventType = `RUNNING`
	ClusterEventTypeResizing ClusterEventType = `RESIZING`
	ClusterEventTypeNodesLost ClusterEventType = `NODES_LOST`
	ClusterEventTypeUpsizeCompleted ClusterEventType = `UPSIZE_COMPLETED`
	ClusterEventTypeInitScriptsStarted ClusterEventType = `INIT_SCRIPTS_STARTED`
	ClusterEventTypeInitScriptsFinished ClusterEventType = `INIT_SCRIPTS_FINISHED`
	ClusterEventTypeDidNotExpandDisk ClusterEventType = `DID_NOT_EXPAND_DISK`
	ClusterEventTypeExpandedDisk ClusterEventType = `EXPANDED_DISK`
	ClusterEventTypeFailedToExpandDisk ClusterEventType = `FAILED_TO_EXPAND_DISK`
	ClusterEventTypeDriverHealthy ClusterEventType = `DRIVER_HEALTHY`
	ClusterEventTypeDriverNotResponding ClusterEventType = `DRIVER_NOT_RESPONDING`
	ClusterEventTypeDriverUnavailable ClusterEventType = `DRIVER_UNAVAILABLE`
	ClusterEventTypeSparkException ClusterEventType = `SPARK_EXCEPTION`
	ClusterEventTypeMetastoreDown ClusterEventType = `METASTORE_DOWN`
	ClusterEventTypeDbfsDown ClusterEventType = `DBFS_DOWN`
	ClusterEventTypeAutoscalingStatsReport ClusterEventType = `AUTOSCALING_STATS_REPORT`
	ClusterEventTypeNodeBlacklisted ClusterEventType = `NODE_BLACKLISTED`
	ClusterEventTypePinned ClusterEventType = `PINNED`
	ClusterEventTypeUnpinned ClusterEventType = `UNPINNED`
	ClusterEventTypeNodeExcludedDecommissioned ClusterEventType = `NODE_EXCLUDED_DECOMMISSIONED`
)


type ClusterInfo struct {
    // Parameters needed in order to automatically scale clusters up and down
    // based on load. Note: autoscaling works best with DB runtime versions 3.0
    // or later.
    Autoscale *AutoScale `json:"autoscale,omitempty"`
    // Automatically terminates the cluster after it is inactive for this time
    // in minutes. If not set, this cluster will not be automatically
    // terminated. If specified, the threshold must be between 10 and 10000
    // minutes. Users can also set this value to 0 to explicitly disable
    // automatic termination.
    AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
    // Attributes related to clusters running on Amazon Web Services. If not
    // specified at cluster creation, a set of default values will be used.
    AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to clusters running on Microsoft Azure. If not
    // specified at cluster creation, a set of default values will be used.
    AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
    // Number of CPU cores available for this cluster. Note that this can be
    // fractional, e.g. 7.5 cores, since certain node types are configured to
    // share cores between Spark nodes on the same instance.
    ClusterCores float64 `json:"cluster_cores,omitempty"`
    // Canonical identifier for the cluster. This id is retained during cluster
    // restarts and resizes, while each new cluster has a globally unique id.
    ClusterId string `json:"cluster_id,omitempty"`
    // The configuration for delivering spark logs to a long-term storage
    // destination. Two kinds of destinations (dbfs and s3) are supported. Only
    // one destination can be specified for one cluster. If the conf is given,
    // the logs will be delivered to the destination every ``5 mins``. The
    // destination of driver logs is ``$destination/$clusterId/driver``, while
    // the destination of executor logs is ``$destination/$clusterId/executor``.
    ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
    // Cluster log delivery status.
    ClusterLogStatus *LogSyncStatus `json:"cluster_log_status,omitempty"`
    // Total amount of cluster memory, in megabytes
    ClusterMemoryMb int64 `json:"cluster_memory_mb,omitempty"`
    // Cluster name requested by the user. This doesn&#39;t have to be unique. If
    // not specified at creation, the cluster name will be an empty string.
    ClusterName string `json:"cluster_name,omitempty"`
    // Determines whether the cluster was created by a user through the UI,
    // created by the Databricks Jobs Scheduler, or through an API request. This
    // is the same as cluster_creator, but read only.
    ClusterSource ClusterInfoClusterSource `json:"cluster_source,omitempty"`
    // Creator user name. The field won&#39;t be included in the response if the
    // user has already been deleted.
    CreatorUserName string `json:"creator_user_name,omitempty"`
    // Additional tags for cluster resources. Databricks will tag all cluster
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags - Clusters can only reuse cloud resources if the
    // resources&#39; tags are a subset of the cluster tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // Tags that are added by Databricks regardless of any ``custom_tags``,
    // including: - Vendor: Databricks - Creator: &lt;username_of_creator&gt; -
    // ClusterName: &lt;name_of_cluster&gt; - ClusterId: &lt;id_of_cluster&gt; - Name:
    // &lt;Databricks internal use&gt;
    DefaultTags map[string]string `json:"default_tags,omitempty"`
    // Node on which the Spark driver resides. The driver node contains the
    // Spark master and the Databricks application that manages the per-notebook
    // Spark REPLs.
    Driver *SparkNode `json:"driver,omitempty"`
    // The optional ID of the instance pool for the driver of the cluster
    // belongs. The pool cluster uses the instance pool with id
    // (instance_pool_id) if the driver pool is not assigned.
    DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
    // The node type of the Spark driver. Note that this field is optional; if
    // unset, the driver node type will be set as the same value as
    // `node_type_id` defined above. This field, along with node_type_id, should
    // not be set if virtual_cluster_size is set. If both driver_node_type_id,
    // node_type_id, and virtual_cluster_size are specified, driver_node_type_id
    // and node_type_id take precedence.
    DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
    // The key of the spark version running in the dataplane. This is possibly
    // different from the spark_version (index 2). The spark_version is the raw
    // string provided by the user through API or UI, which could map to a
    // different effective_spark_version running in the dataplane, depending on
    // the cluster&#39;s instance type or the runtimeEngine parameter.
    EffectiveSparkVersion string `json:"effective_spark_version,omitempty"`
    // Autoscaling Local Storage: when enabled, this cluster will dynamically
    // acquire additional disk space when its Spark workers are running low on
    // disk space. This feature requires specific AWS permissions to function
    // correctly - refer to the User Guide for more details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Whether to enable LUKS on cluster VMs&#39; local disks
    EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
    // Nodes on which the Spark executors reside.
    Executors []SparkNode `json:"executors,omitempty"`
    // Attributes related to clusters running on Google Cloud Platform. If not
    // specified at cluster creation, a set of default values will be used.
    GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
    // The optional ID of the instance pool to which the cluster belongs.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // Port on which Spark JDBC server is listening, in the driver nod. No
    // service will be listeningon on this port in executor nodes.
    JdbcPort int `json:"jdbc_port,omitempty"`
    // the timestamp that the cluster was started/restarted
    LastRestartedTime int64 `json:"last_restarted_time,omitempty"`
    // Time when the cluster driver last lost its state (due to a restart or
    // driver failure).
    LastStateLossTime int64 `json:"last_state_loss_time,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this cluster. For example, the Spark nodes can
    // be provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call. This field, along
    // with driver_node_type_id, should not be set if virtual_cluster_size is
    // set. If both driver_node_type_id, node_type_id, and virtual_cluster_size
    // are specified, driver_node_type_id and node_type_id take precedence.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Number of worker nodes that this cluster should have. A cluster has one
    // Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
    // &#43; 1 Spark nodes. Note: When reading the properties of a cluster, this
    // field reflects the desired number of workers rather than the actual
    // current number of workers. For instance, if a cluster is resized from 5
    // to 10 workers, this field will immediately be updated to reflect the
    // target size of 10 workers, whereas the workers listed in ``spark_info``
    // will gradually increase from 5 to 10 as the new nodes are provisioned.
    NumWorkers int `json:"num_workers,omitempty"`
    // The ID of the cluster policy used to create the cluster if applicable.
    PolicyId string `json:"policy_id,omitempty"`
    // Decides which runtime engine to be use, e.g. Standard vs. Photon. If
    // unspecified, the runtime engine is inferred from spark_version.
    RuntimeEngine ClusterInfoRuntimeEngine `json:"runtime_engine,omitempty"`
    // An object containing a set of optional, user-specified Spark
    // configuration key-value pairs. Users can also pass in a string of extra
    // JVM options to the driver and the executors via
    // ``spark.driver.extraJavaOptions`` and ``spark.executor.extraJavaOptions``
    // respectively. Example Spark confs: ``{&#34;spark.speculation&#34;: true,
    // &#34;spark.streaming.ui.retainedBatches&#34;: 5}`` or
    // ``{&#34;spark.driver.extraJavaOptions&#34;: &#34;-verbose:gc -XX:&#43;PrintGCDetails&#34;}``
    SparkConf map[string]string `json:"spark_conf,omitempty"`
    // A canonical SparkContext identifier. This value *does* change when the
    // Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
    // globally unique identifier over all Spark contexts.
    SparkContextId int64 `json:"spark_context_id,omitempty"`
    // An object containing a set of optional, user-specified environment
    // variable key-value pairs. Please note that key-value pair of the form
    // (X,Y) will be exported as is (i.e., ``export X=&#39;Y&#39;``) while launching the
    // driver and workers. In order to specify an additional set of
    // ``SPARK_DAEMON_JAVA_OPTS``, we recommend appending them to
    // ``$SPARK_DAEMON_JAVA_OPTS`` as shown in the example below. This ensures
    // that all default databricks managed environmental variables are included
    // as well. Example Spark environment variables: ``{&#34;SPARK_WORKER_MEMORY&#34;:
    // &#34;28000m&#34;, &#34;SPARK_LOCAL_DIRS&#34;: &#34;/local_disk0&#34;}`` or
    // ``{&#34;SPARK_DAEMON_JAVA_OPTS&#34;: &#34;$SPARK_DAEMON_JAVA_OPTS
    // -Dspark.shuffle.service.enabled=true&#34;}``
    SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
    // The Spark version of the cluster, e.g. &#34;3.3.x-scala2.11&#34;. A list of
    // available Spark versions can be retrieved by using the
    // :ref:`clusterClusterServicelistSparkVersions` API call. This is the Spark
    // version provided from the user input (API/UI) and may be different from
    // the effective_spark_version, which is the spark version that is actually
    // run in the dataplane. See index 45 for more context.
    SparkVersion string `json:"spark_version,omitempty"`
    // SSH public key contents that will be added to each Spark node in this
    // cluster. The corresponding private keys can be used to login with the
    // user name ``ubuntu`` on port ``2200``. Up to 10 keys can be specified.
    SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
    // Time (in epoch milliseconds) when the cluster creation request was
    // received (when the cluster entered a ``PENDING`` state).
    StartTime int64 `json:"start_time,omitempty"`
    // Current state of the cluster.
    State ClusterInfoState `json:"state,omitempty"`
    // A message associated with the most recent state transition (e.g., the
    // reason why the cluster entered a ``TERMINATED`` state).
    StateMessage string `json:"state_message,omitempty"`
    // Time (in epoch milliseconds) when the cluster was terminated, if
    // applicable.
    TerminatedTime int64 `json:"terminated_time,omitempty"`
    // Information about why the cluster was terminated. This field only appears
    // when the cluster is in a ``TERMINATING`` or ``TERMINATED`` state.
    TerminationReason *TerminationReason `json:"termination_reason,omitempty"`
    
    WorkloadType *WorkloadType `json:"workload_type,omitempty"`
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.

type ClusterInfoClusterSource string
const (
	ClusterInfoClusterSourceUi ClusterInfoClusterSource = `UI`
	ClusterInfoClusterSourceJob ClusterInfoClusterSource = `JOB`
	ClusterInfoClusterSourceApi ClusterInfoClusterSource = `API`
	ClusterInfoClusterSourceSql ClusterInfoClusterSource = `SQL`
	ClusterInfoClusterSourceModels ClusterInfoClusterSource = `MODELS`
	ClusterInfoClusterSourcePipeline ClusterInfoClusterSource = `PIPELINE`
	ClusterInfoClusterSourcePipelineMaintenance ClusterInfoClusterSource = `PIPELINE_MAINTENANCE`
)

// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
// unspecified, the runtime engine is inferred from spark_version.

type ClusterInfoRuntimeEngine string
const (
	ClusterInfoRuntimeEngineNull ClusterInfoRuntimeEngine = `NULL`
	ClusterInfoRuntimeEngineStandard ClusterInfoRuntimeEngine = `STANDARD`
	ClusterInfoRuntimeEnginePhoton ClusterInfoRuntimeEngine = `PHOTON`
)

// Current state of the cluster.

type ClusterInfoState string
const (
	ClusterInfoStatePending ClusterInfoState = `PENDING`
	ClusterInfoStateRunning ClusterInfoState = `RUNNING`
	ClusterInfoStateRestarting ClusterInfoState = `RESTARTING`
	ClusterInfoStateResizing ClusterInfoState = `RESIZING`
	ClusterInfoStateTerminating ClusterInfoState = `TERMINATING`
	ClusterInfoStateTerminated ClusterInfoState = `TERMINATED`
	ClusterInfoStateError ClusterInfoState = `ERROR`
	ClusterInfoStateUnknown ClusterInfoState = `UNKNOWN`
)


type ClusterLogConf struct {
    // destination needs to be provided. e.g. ``{ &#34;dbfs&#34; : { &#34;destination&#34; :
    // &#34;dbfs:/home/cluster_log&#34; } }``
    Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
    // destination and either region or endpoint should also be provided. e.g.
    // ``{ &#34;s3&#34;: { &#34;destination&#34; : &#34;s3://cluster_log_bucket/prefix&#34;, &#34;region&#34; :
    // &#34;us-west-2&#34; } }`` Cluster iam role is used to access s3, please make sure
    // the cluster iam role in ``instance_profile_arn`` has permission to write
    // data to the s3 destination.
    S3 *S3StorageInfo `json:"s3,omitempty"`
}


type ClusterSize struct {
    // Parameters needed in order to automatically scale clusters up and down
    // based on load. Note: autoscaling works best with DB runtime versions 3.0
    // or later.
    Autoscale *AutoScale `json:"autoscale,omitempty"`
    // Number of worker nodes that this cluster should have. A cluster has one
    // Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
    // &#43; 1 Spark nodes. Note: When reading the properties of a cluster, this
    // field reflects the desired number of workers rather than the actual
    // current number of workers. For instance, if a cluster is resized from 5
    // to 10 workers, this field will immediately be updated to reflect the
    // target size of 10 workers, whereas the workers listed in ``spark_info``
    // will gradually increase from 5 to 10 as the new nodes are provisioned.
    NumWorkers int `json:"num_workers,omitempty"`
}


type CpalFailureResponse struct {
    
    AwsCpalErrorCode CpalFailureResponseAwsCpalErrorCode `json:"aws_cpal_error_code,omitempty"`
    
    AwsError *AwsCpalError `json:"aws_error,omitempty"`
    
    AzureCpalErrorCode CpalFailureResponseAzureCpalErrorCode `json:"azure_cpal_error_code,omitempty"`
    // Deprecated after using cloud_provider_error field TODO define
    // AzureCpalError with code and message
    CloudProviderErrorCode string `json:"cloud_provider_error_code,omitempty"`
    // Deprecated after using cloud_provider_error field TODO define
    // AzureCpalError with code and message
    CloudProviderErrorMessage string `json:"cloud_provider_error_message,omitempty"`
    
    DelegateErrorCode CpalFailureResponseDelegateErrorCode `json:"delegate_error_code,omitempty"`
    
    ErrorMessage string `json:"error_message,omitempty"`
    
    StackTrace string `json:"stack_trace,omitempty"`
}



type CpalFailureResponseAwsCpalErrorCode string
const (
	CpalFailureResponseAwsCpalErrorCodeAwsServiceException CpalFailureResponseAwsCpalErrorCode = `AWS_SERVICE_EXCEPTION`
	CpalFailureResponseAwsCpalErrorCodeAwsUnexpectedInstanceState CpalFailureResponseAwsCpalErrorCode = `AWS_UNEXPECTED_INSTANCE_STATE`
	CpalFailureResponseAwsCpalErrorCodeAwsSpotRequestFailed CpalFailureResponseAwsCpalErrorCode = `AWS_SPOT_REQUEST_FAILED`
	CpalFailureResponseAwsCpalErrorCodeAwsUnexpectedException CpalFailureResponseAwsCpalErrorCode = `AWS_UNEXPECTED_EXCEPTION`
	CpalFailureResponseAwsCpalErrorCodeAwsInsufficientInstanceCapacityFailure CpalFailureResponseAwsCpalErrorCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`
	CpalFailureResponseAwsCpalErrorCodeAwsUnsupportedInstanceTypeFailure CpalFailureResponseAwsCpalErrorCode = `AWS_UNSUPPORTED_INSTANCE_TYPE_FAILURE`
	CpalFailureResponseAwsCpalErrorCodeAwsInsufficientFreeAddressesInSubnetFailure CpalFailureResponseAwsCpalErrorCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`
	CpalFailureResponseAwsCpalErrorCodeAwsRequestLimitExceeded CpalFailureResponseAwsCpalErrorCode = `AWS_REQUEST_LIMIT_EXCEEDED`
	CpalFailureResponseAwsCpalErrorCodeAwsInvalidInstanceTypeFailure CpalFailureResponseAwsCpalErrorCode = `AWS_INVALID_INSTANCE_TYPE_FAILURE`
	CpalFailureResponseAwsCpalErrorCodeAwsMaxSpotInstanceCountExceeded CpalFailureResponseAwsCpalErrorCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED`
)



type CpalFailureResponseAzureCpalErrorCode string
const (
	CpalFailureResponseAzureCpalErrorCodeAzureCloudException CpalFailureResponseAzureCpalErrorCode = `AZURE_CLOUD_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureResourceManagerThrottling CpalFailureResponseAzureCpalErrorCode = `AZURE_RESOURCE_MANAGER_THROTTLING`
	CpalFailureResponseAzureCpalErrorCodeAzureResourceProviderThrottling CpalFailureResponseAzureCpalErrorCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`
	CpalFailureResponseAzureCpalErrorCodeAzureServerUnreachable CpalFailureResponseAzureCpalErrorCode = `AZURE_SERVER_UNREACHABLE`
	CpalFailureResponseAzureCpalErrorCodeAzureNullPointer CpalFailureResponseAzureCpalErrorCode = `AZURE_NULL_POINTER`
	CpalFailureResponseAzureCpalErrorCodeAzureUnknownException CpalFailureResponseAzureCpalErrorCode = `AZURE_UNKNOWN_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureVnetConfigurationException CpalFailureResponseAzureCpalErrorCode = `AZURE_VNET_CONFIGURATION_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureLoadBalancerConfigurationException CpalFailureResponseAzureCpalErrorCode = `AZURE_LOAD_BALANCER_CONFIGURATION_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureQuotaExceededException CpalFailureResponseAzureCpalErrorCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureAuthenticationException CpalFailureResponseAzureCpalErrorCode = `AZURE_AUTHENTICATION_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureSpotRequestException CpalFailureResponseAzureCpalErrorCode = `AZURE_SPOT_REQUEST_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureNotRunningVmException CpalFailureResponseAzureCpalErrorCode = `AZURE_NOT_RUNNING_VM_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureInvalidDeploymentTemplate CpalFailureResponseAzureCpalErrorCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`
	CpalFailureResponseAzureCpalErrorCodeAzureUnexpectedDeploymentTemplateFailure CpalFailureResponseAzureCpalErrorCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`
	CpalFailureResponseAzureCpalErrorCodeAzureSubnetExhaustedFailure CpalFailureResponseAzureCpalErrorCode = `AZURE_SUBNET_EXHAUSTED_FAILURE`
	CpalFailureResponseAzureCpalErrorCodeAzureOperationNotAllowedException CpalFailureResponseAzureCpalErrorCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`
	CpalFailureResponseAzureCpalErrorCodeAzureServerException CpalFailureResponseAzureCpalErrorCode = `AZURE_SERVER_EXCEPTION`
)



type CpalFailureResponseDelegateErrorCode string
const (
	CpalFailureResponseDelegateErrorCodeDelegateUnexpectedException CpalFailureResponseDelegateErrorCode = `DELEGATE_UNEXPECTED_EXCEPTION`
	CpalFailureResponseDelegateErrorCodeNoSuchWorkerEnvironmentException CpalFailureResponseDelegateErrorCode = `NO_SUCH_WORKER_ENVIRONMENT_EXCEPTION`
)


type CreateClusterRequest struct {
    // Note: This field won&#39;t be true for webapp requests. Only API users will
    // check this field.
    ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
    // Parameters needed in order to automatically scale clusters up and down
    // based on load. Note: autoscaling works best with DB runtime versions 3.0
    // or later.
    Autoscale *AutoScale `json:"autoscale,omitempty"`
    // Automatically terminates the cluster after it is inactive for this time
    // in minutes. If not set, this cluster will not be automatically
    // terminated. If specified, the threshold must be between 10 and 10000
    // minutes. Users can also set this value to 0 to explicitly disable
    // automatic termination.
    AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
    // Attributes related to clusters running on Amazon Web Services. If not
    // specified at cluster creation, a set of default values will be used.
    AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to clusters running on Microsoft Azure. If not
    // specified at cluster creation, a set of default values will be used.
    AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
    // The configuration for delivering spark logs to a long-term storage
    // destination. Two kinds of destinations (dbfs and s3) are supported. Only
    // one destination can be specified for one cluster. If the conf is given,
    // the logs will be delivered to the destination every ``5 mins``. The
    // destination of driver logs is ``$destination/$clusterId/driver``, while
    // the destination of executor logs is ``$destination/$clusterId/executor``.
    ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
    // Cluster name requested by the user. This doesn&#39;t have to be unique. If
    // not specified at creation, the cluster name will be an empty string.
    ClusterName string `json:"cluster_name,omitempty"`
    // Determines whether the cluster was created by a user through the UI,
    // created by the Databricks Jobs Scheduler, or through an API request. This
    // is the same as cluster_creator, but read only.
    ClusterSource CreateClusterRequestClusterSource `json:"cluster_source,omitempty"`
    // Additional tags for cluster resources. Databricks will tag all cluster
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags - Clusters can only reuse cloud resources if the
    // resources&#39; tags are a subset of the cluster tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // The optional ID of the instance pool for the driver of the cluster
    // belongs. The pool cluster uses the instance pool with id
    // (instance_pool_id) if the driver pool is not assigned.
    DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
    // The node type of the Spark driver. Note that this field is optional; if
    // unset, the driver node type will be set as the same value as
    // `node_type_id` defined above. This field, along with node_type_id, should
    // not be set if virtual_cluster_size is set. If both driver_node_type_id,
    // node_type_id, and virtual_cluster_size are specified, driver_node_type_id
    // and node_type_id take precedence.
    DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
    // The key of the spark version running in the dataplane. This is possibly
    // different from the spark_version (index 2). The spark_version is the raw
    // string provided by the user through API or UI, which could map to a
    // different effective_spark_version running in the dataplane, depending on
    // the cluster&#39;s instance type or the runtimeEngine parameter.
    EffectiveSparkVersion string `json:"effective_spark_version,omitempty"`
    // Autoscaling Local Storage: when enabled, this cluster will dynamically
    // acquire additional disk space when its Spark workers are running low on
    // disk space. This feature requires specific AWS permissions to function
    // correctly - refer to the User Guide for more details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Whether to enable LUKS on cluster VMs&#39; local disks
    EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
    // Attributes related to clusters running on Google Cloud Platform. If not
    // specified at cluster creation, a set of default values will be used.
    GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
    // The optional ID of the instance pool to which the cluster belongs.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this cluster. For example, the Spark nodes can
    // be provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call. This field, along
    // with driver_node_type_id, should not be set if virtual_cluster_size is
    // set. If both driver_node_type_id, node_type_id, and virtual_cluster_size
    // are specified, driver_node_type_id and node_type_id take precedence.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Number of worker nodes that this cluster should have. A cluster has one
    // Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
    // &#43; 1 Spark nodes. Note: When reading the properties of a cluster, this
    // field reflects the desired number of workers rather than the actual
    // current number of workers. For instance, if a cluster is resized from 5
    // to 10 workers, this field will immediately be updated to reflect the
    // target size of 10 workers, whereas the workers listed in ``spark_info``
    // will gradually increase from 5 to 10 as the new nodes are provisioned.
    NumWorkers int `json:"num_workers,omitempty"`
    // The ID of the cluster policy used to create the cluster if applicable.
    PolicyId string `json:"policy_id,omitempty"`
    // Decides which runtime engine to be use, e.g. Standard vs. Photon. If
    // unspecified, the runtime engine is inferred from spark_version.
    RuntimeEngine CreateClusterRequestRuntimeEngine `json:"runtime_engine,omitempty"`
    // An object containing a set of optional, user-specified Spark
    // configuration key-value pairs. Users can also pass in a string of extra
    // JVM options to the driver and the executors via
    // ``spark.driver.extraJavaOptions`` and ``spark.executor.extraJavaOptions``
    // respectively. Example Spark confs: ``{&#34;spark.speculation&#34;: true,
    // &#34;spark.streaming.ui.retainedBatches&#34;: 5}`` or
    // ``{&#34;spark.driver.extraJavaOptions&#34;: &#34;-verbose:gc -XX:&#43;PrintGCDetails&#34;}``
    SparkConf map[string]string `json:"spark_conf,omitempty"`
    // An object containing a set of optional, user-specified environment
    // variable key-value pairs. Please note that key-value pair of the form
    // (X,Y) will be exported as is (i.e., ``export X=&#39;Y&#39;``) while launching the
    // driver and workers. In order to specify an additional set of
    // ``SPARK_DAEMON_JAVA_OPTS``, we recommend appending them to
    // ``$SPARK_DAEMON_JAVA_OPTS`` as shown in the example below. This ensures
    // that all default databricks managed environmental variables are included
    // as well. Example Spark environment variables: ``{&#34;SPARK_WORKER_MEMORY&#34;:
    // &#34;28000m&#34;, &#34;SPARK_LOCAL_DIRS&#34;: &#34;/local_disk0&#34;}`` or
    // ``{&#34;SPARK_DAEMON_JAVA_OPTS&#34;: &#34;$SPARK_DAEMON_JAVA_OPTS
    // -Dspark.shuffle.service.enabled=true&#34;}``
    SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
    // The Spark version of the cluster, e.g. &#34;3.3.x-scala2.11&#34;. A list of
    // available Spark versions can be retrieved by using the
    // :ref:`clusterClusterServicelistSparkVersions` API call. This is the Spark
    // version provided from the user input (API/UI) and may be different from
    // the effective_spark_version, which is the spark version that is actually
    // run in the dataplane. See index 45 for more context.
    SparkVersion string `json:"spark_version,omitempty"`
    // SSH public key contents that will be added to each Spark node in this
    // cluster. The corresponding private keys can be used to login with the
    // user name ``ubuntu`` on port ``2200``. Up to 10 keys can be specified.
    SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
    
    WorkloadType *WorkloadType `json:"workload_type,omitempty"`
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.

type CreateClusterRequestClusterSource string
const (
	CreateClusterRequestClusterSourceUi CreateClusterRequestClusterSource = `UI`
	CreateClusterRequestClusterSourceJob CreateClusterRequestClusterSource = `JOB`
	CreateClusterRequestClusterSourceApi CreateClusterRequestClusterSource = `API`
	CreateClusterRequestClusterSourceSql CreateClusterRequestClusterSource = `SQL`
	CreateClusterRequestClusterSourceModels CreateClusterRequestClusterSource = `MODELS`
	CreateClusterRequestClusterSourcePipeline CreateClusterRequestClusterSource = `PIPELINE`
	CreateClusterRequestClusterSourcePipelineMaintenance CreateClusterRequestClusterSource = `PIPELINE_MAINTENANCE`
)

// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
// unspecified, the runtime engine is inferred from spark_version.

type CreateClusterRequestRuntimeEngine string
const (
	CreateClusterRequestRuntimeEngineNull CreateClusterRequestRuntimeEngine = `NULL`
	CreateClusterRequestRuntimeEngineStandard CreateClusterRequestRuntimeEngine = `STANDARD`
	CreateClusterRequestRuntimeEnginePhoton CreateClusterRequestRuntimeEngine = `PHOTON`
)


type CreateClusterResponse struct {
    
    ClusterId string `json:"cluster_id,omitempty"`
}


type DataPlaneEventDetails struct {
    
    EventType DataPlaneEventDetailsEventType `json:"event_type,omitempty"`
    
    ExecutorFailures int `json:"executor_failures,omitempty"`
    
    HostId string `json:"host_id,omitempty"`
    
    Timestamp int64 `json:"timestamp,omitempty"`
}



type DataPlaneEventDetailsEventType string
const (
	DataPlaneEventDetailsEventTypeNodeBlacklisted DataPlaneEventDetailsEventType = `NODE_BLACKLISTED`
	DataPlaneEventDetailsEventTypeNodeExcludedDecommissioned DataPlaneEventDetailsEventType = `NODE_EXCLUDED_DECOMMISSIONED`
)


type DbfsStorageInfo struct {
    // dbfs destination, e.g. ``dbfs:/my/path``
    Destination string `json:"destination,omitempty"`
}


type DeleteClusterRequest struct {
    // The cluster to be terminated.
    ClusterId string `json:"cluster_id"`
}


type EditClusterRequest struct {
    // Note: This field won&#39;t be true for webapp requests. Only API users will
    // check this field.
    ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
    // Parameters needed in order to automatically scale clusters up and down
    // based on load. Note: autoscaling works best with DB runtime versions 3.0
    // or later.
    Autoscale *AutoScale `json:"autoscale,omitempty"`
    // Automatically terminates the cluster after it is inactive for this time
    // in minutes. If not set, this cluster will not be automatically
    // terminated. If specified, the threshold must be between 10 and 10000
    // minutes. Users can also set this value to 0 to explicitly disable
    // automatic termination.
    AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
    // Attributes related to clusters running on Amazon Web Services. If not
    // specified at cluster creation, a set of default values will be used.
    AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to clusters running on Microsoft Azure. If not
    // specified at cluster creation, a set of default values will be used.
    AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
    
    ClusterId string `json:"cluster_id"`
    // The configuration for delivering spark logs to a long-term storage
    // destination. Two kinds of destinations (dbfs and s3) are supported. Only
    // one destination can be specified for one cluster. If the conf is given,
    // the logs will be delivered to the destination every ``5 mins``. The
    // destination of driver logs is ``$destination/$clusterId/driver``, while
    // the destination of executor logs is ``$destination/$clusterId/executor``.
    ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
    // Cluster name requested by the user. This doesn&#39;t have to be unique. If
    // not specified at creation, the cluster name will be an empty string.
    ClusterName string `json:"cluster_name,omitempty"`
    // Determines whether the cluster was created by a user through the UI,
    // created by the Databricks Jobs Scheduler, or through an API request. This
    // is the same as cluster_creator, but read only.
    ClusterSource EditClusterRequestClusterSource `json:"cluster_source,omitempty"`
    // Additional tags for cluster resources. Databricks will tag all cluster
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags - Clusters can only reuse cloud resources if the
    // resources&#39; tags are a subset of the cluster tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // The optional ID of the instance pool for the driver of the cluster
    // belongs. The pool cluster uses the instance pool with id
    // (instance_pool_id) if the driver pool is not assigned.
    DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
    // The node type of the Spark driver. Note that this field is optional; if
    // unset, the driver node type will be set as the same value as
    // `node_type_id` defined above. This field, along with node_type_id, should
    // not be set if virtual_cluster_size is set. If both driver_node_type_id,
    // node_type_id, and virtual_cluster_size are specified, driver_node_type_id
    // and node_type_id take precedence.
    DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
    // The key of the spark version running in the dataplane. This is possibly
    // different from the spark_version (index 2). The spark_version is the raw
    // string provided by the user through API or UI, which could map to a
    // different effective_spark_version running in the dataplane, depending on
    // the cluster&#39;s instance type or the runtimeEngine parameter.
    EffectiveSparkVersion string `json:"effective_spark_version,omitempty"`
    // Autoscaling Local Storage: when enabled, this cluster will dynamically
    // acquire additional disk space when its Spark workers are running low on
    // disk space. This feature requires specific AWS permissions to function
    // correctly - refer to the User Guide for more details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Whether to enable LUKS on cluster VMs&#39; local disks
    EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
    // Attributes related to clusters running on Google Cloud Platform. If not
    // specified at cluster creation, a set of default values will be used.
    GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
    // The optional ID of the instance pool to which the cluster belongs.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this cluster. For example, the Spark nodes can
    // be provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call. This field, along
    // with driver_node_type_id, should not be set if virtual_cluster_size is
    // set. If both driver_node_type_id, node_type_id, and virtual_cluster_size
    // are specified, driver_node_type_id and node_type_id take precedence.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Number of worker nodes that this cluster should have. A cluster has one
    // Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
    // &#43; 1 Spark nodes. Note: When reading the properties of a cluster, this
    // field reflects the desired number of workers rather than the actual
    // current number of workers. For instance, if a cluster is resized from 5
    // to 10 workers, this field will immediately be updated to reflect the
    // target size of 10 workers, whereas the workers listed in ``spark_info``
    // will gradually increase from 5 to 10 as the new nodes are provisioned.
    NumWorkers int `json:"num_workers,omitempty"`
    // The ID of the cluster policy used to create the cluster if applicable.
    PolicyId string `json:"policy_id,omitempty"`
    // Decides which runtime engine to be use, e.g. Standard vs. Photon. If
    // unspecified, the runtime engine is inferred from spark_version.
    RuntimeEngine EditClusterRequestRuntimeEngine `json:"runtime_engine,omitempty"`
    // An object containing a set of optional, user-specified Spark
    // configuration key-value pairs. Users can also pass in a string of extra
    // JVM options to the driver and the executors via
    // ``spark.driver.extraJavaOptions`` and ``spark.executor.extraJavaOptions``
    // respectively. Example Spark confs: ``{&#34;spark.speculation&#34;: true,
    // &#34;spark.streaming.ui.retainedBatches&#34;: 5}`` or
    // ``{&#34;spark.driver.extraJavaOptions&#34;: &#34;-verbose:gc -XX:&#43;PrintGCDetails&#34;}``
    SparkConf map[string]string `json:"spark_conf,omitempty"`
    // An object containing a set of optional, user-specified environment
    // variable key-value pairs. Please note that key-value pair of the form
    // (X,Y) will be exported as is (i.e., ``export X=&#39;Y&#39;``) while launching the
    // driver and workers. In order to specify an additional set of
    // ``SPARK_DAEMON_JAVA_OPTS``, we recommend appending them to
    // ``$SPARK_DAEMON_JAVA_OPTS`` as shown in the example below. This ensures
    // that all default databricks managed environmental variables are included
    // as well. Example Spark environment variables: ``{&#34;SPARK_WORKER_MEMORY&#34;:
    // &#34;28000m&#34;, &#34;SPARK_LOCAL_DIRS&#34;: &#34;/local_disk0&#34;}`` or
    // ``{&#34;SPARK_DAEMON_JAVA_OPTS&#34;: &#34;$SPARK_DAEMON_JAVA_OPTS
    // -Dspark.shuffle.service.enabled=true&#34;}``
    SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
    // The Spark version of the cluster, e.g. &#34;3.3.x-scala2.11&#34;. A list of
    // available Spark versions can be retrieved by using the
    // :ref:`clusterClusterServicelistSparkVersions` API call. This is the Spark
    // version provided from the user input (API/UI) and may be different from
    // the effective_spark_version, which is the spark version that is actually
    // run in the dataplane. See index 45 for more context.
    SparkVersion string `json:"spark_version,omitempty"`
    // SSH public key contents that will be added to each Spark node in this
    // cluster. The corresponding private keys can be used to login with the
    // user name ``ubuntu`` on port ``2200``. Up to 10 keys can be specified.
    SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
    
    WorkloadType *WorkloadType `json:"workload_type,omitempty"`
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.

type EditClusterRequestClusterSource string
const (
	EditClusterRequestClusterSourceUi EditClusterRequestClusterSource = `UI`
	EditClusterRequestClusterSourceJob EditClusterRequestClusterSource = `JOB`
	EditClusterRequestClusterSourceApi EditClusterRequestClusterSource = `API`
	EditClusterRequestClusterSourceSql EditClusterRequestClusterSource = `SQL`
	EditClusterRequestClusterSourceModels EditClusterRequestClusterSource = `MODELS`
	EditClusterRequestClusterSourcePipeline EditClusterRequestClusterSource = `PIPELINE`
	EditClusterRequestClusterSourcePipelineMaintenance EditClusterRequestClusterSource = `PIPELINE_MAINTENANCE`
)

// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
// unspecified, the runtime engine is inferred from spark_version.

type EditClusterRequestRuntimeEngine string
const (
	EditClusterRequestRuntimeEngineNull EditClusterRequestRuntimeEngine = `NULL`
	EditClusterRequestRuntimeEngineStandard EditClusterRequestRuntimeEngine = `STANDARD`
	EditClusterRequestRuntimeEnginePhoton EditClusterRequestRuntimeEngine = `PHOTON`
)


type EventDetails struct {
    // * For created clusters, the attributes of the cluster. * For edited
    // clusters, the new attributes of the cluster.
    Attributes *ClusterAttributes `json:"attributes,omitempty"`
    // The cause of a change in target size.
    Cause EventDetailsCause `json:"cause,omitempty"`
    // The actual cluster size that was set in the cluster creation or edit.
    ClusterSize *ClusterSize `json:"cluster_size,omitempty"`
    // The current number of vCPUs in the cluster.
    CurrentNumVcpus int `json:"current_num_vcpus,omitempty"`
    // The current number of nodes in the cluster.
    CurrentNumWorkers int `json:"current_num_workers,omitempty"`
    
    DidNotExpandReason string `json:"did_not_expand_reason,omitempty"`
    // Current disk size in bytes
    DiskSize int64 `json:"disk_size,omitempty"`
    // More details about the change in driver&#39;s state
    DriverStateMessage string `json:"driver_state_message,omitempty"`
    // Whether or not a blocklisted node should be terminated. For
    // ClusterEventType NODE_BLACKLISTED.
    EnableTerminationForNodeBlocklisted bool `json:"enable_termination_for_node_blocklisted,omitempty"`
    
    FreeSpace int64 `json:"free_space,omitempty"`
    // Instance Id where the event originated from
    InstanceId string `json:"instance_id,omitempty"`
    // Unique identifier of the specific job run associated with this cluster
    // event * For clusters created for jobs, this will be the same as the
    // cluster name
    JobRunName string `json:"job_run_name,omitempty"`
    // The cluster attributes before a cluster was edited.
    PreviousAttributes *ClusterAttributes `json:"previous_attributes,omitempty"`
    // The size of the cluster before an edit or resize.
    PreviousClusterSize *ClusterSize `json:"previous_cluster_size,omitempty"`
    // Previous disk size in bytes
    PreviousDiskSize int64 `json:"previous_disk_size,omitempty"`
    // A termination reason: * On a TERMINATED event, this is the reason of the
    // termination. * On a RESIZE_COMPLETE event, this indicates the reason that
    // we failed to acquire some nodes.
    Reason *TerminationReason `json:"reason,omitempty"`
    // The targeted number of vCPUs in the cluster.
    TargetNumVcpus int `json:"target_num_vcpus,omitempty"`
    // The targeted number of nodes in the cluster.
    TargetNumWorkers int `json:"target_num_workers,omitempty"`
    // The user that caused the event to occur. (Empty if it was done by the
    // control plane.)
    User string `json:"user,omitempty"`
}

// The cause of a change in target size.

type EventDetailsCause string
const (
	EventDetailsCauseAutoscale EventDetailsCause = `AUTOSCALE`
	EventDetailsCauseUserRequest EventDetailsCause = `USER_REQUEST`
	EventDetailsCauseAutorecovery EventDetailsCause = `AUTORECOVERY`
	EventDetailsCauseReplaceBadNodes EventDetailsCause = `REPLACE_BAD_NODES`
)


type GcpAttributes struct {
    // This field determines whether the spark executors will be scheduled to
    // run on preemptible VMs, on-demand VMs, or preemptible VMs with a fallback
    // to on-demand VMs if the former is unavailable.
    Availability GcpAttributesAvailability `json:"availability,omitempty"`
    // boot disk size in GB
    BootDiskSize int `json:"boot_disk_size,omitempty"`
    // If provided, the cluster will impersonate the google service account when
    // accessing gcloud services (like GCS). The google service account must
    // have previously been added to the Databricks environment by an account
    // administrator.
    GoogleServiceAccount string `json:"google_service_account,omitempty"`
}

// This field determines whether the spark executors will be scheduled to run on
// preemptible VMs, on-demand VMs, or preemptible VMs with a fallback to
// on-demand VMs if the former is unavailable.

type GcpAttributesAvailability string
const (
	GcpAttributesAvailabilityPreemptibleGcp GcpAttributesAvailability = `PREEMPTIBLE_GCP`
	GcpAttributesAvailabilityOnDemandGcp GcpAttributesAvailability = `ON_DEMAND_GCP`
	GcpAttributesAvailabilityPreemptibleWithFallbackGcp GcpAttributesAvailability = `PREEMPTIBLE_WITH_FALLBACK_GCP`
)


type GetClusterRequest struct {
    // The cluster about which to retrieve information.
    ClusterId string ` url:"cluster_id,omitempty"`
}


type GetClusterResponse struct {
    // Parameters needed in order to automatically scale clusters up and down
    // based on load. Note: autoscaling works best with DB runtime versions 3.0
    // or later.
    Autoscale *AutoScale `json:"autoscale,omitempty"`
    // Automatically terminates the cluster after it is inactive for this time
    // in minutes. If not set, this cluster will not be automatically
    // terminated. If specified, the threshold must be between 10 and 10000
    // minutes. Users can also set this value to 0 to explicitly disable
    // automatic termination.
    AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
    // Attributes related to clusters running on Amazon Web Services. If not
    // specified at cluster creation, a set of default values will be used.
    AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to clusters running on Microsoft Azure. If not
    // specified at cluster creation, a set of default values will be used.
    AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
    // Number of CPU cores available for this cluster. Note that this can be
    // fractional, e.g. 7.5 cores, since certain node types are configured to
    // share cores between Spark nodes on the same instance.
    ClusterCores float64 `json:"cluster_cores,omitempty"`
    // Canonical identifier for the cluster. This id is retained during cluster
    // restarts and resizes, while each new cluster has a globally unique id.
    ClusterId string `json:"cluster_id,omitempty"`
    // The configuration for delivering spark logs to a long-term storage
    // destination. Two kinds of destinations (dbfs and s3) are supported. Only
    // one destination can be specified for one cluster. If the conf is given,
    // the logs will be delivered to the destination every ``5 mins``. The
    // destination of driver logs is ``$destination/$clusterId/driver``, while
    // the destination of executor logs is ``$destination/$clusterId/executor``.
    ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
    // Cluster log delivery status.
    ClusterLogStatus *LogSyncStatus `json:"cluster_log_status,omitempty"`
    // Total amount of cluster memory, in megabytes
    ClusterMemoryMb int64 `json:"cluster_memory_mb,omitempty"`
    // Cluster name requested by the user. This doesn&#39;t have to be unique. If
    // not specified at creation, the cluster name will be an empty string.
    ClusterName string `json:"cluster_name,omitempty"`
    // Determines whether the cluster was created by a user through the UI,
    // created by the Databricks Jobs Scheduler, or through an API request. This
    // is the same as cluster_creator, but read only.
    ClusterSource GetClusterResponseClusterSource `json:"cluster_source,omitempty"`
    // Creator user name. The field won&#39;t be included in the response if the
    // user has already been deleted.
    CreatorUserName string `json:"creator_user_name,omitempty"`
    // Additional tags for cluster resources. Databricks will tag all cluster
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags - Clusters can only reuse cloud resources if the
    // resources&#39; tags are a subset of the cluster tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // Tags that are added by Databricks regardless of any ``custom_tags``,
    // including: - Vendor: Databricks - Creator: &lt;username_of_creator&gt; -
    // ClusterName: &lt;name_of_cluster&gt; - ClusterId: &lt;id_of_cluster&gt; - Name:
    // &lt;Databricks internal use&gt;
    DefaultTags map[string]string `json:"default_tags,omitempty"`
    // Node on which the Spark driver resides. The driver node contains the
    // Spark master and the Databricks application that manages the per-notebook
    // Spark REPLs.
    Driver *SparkNode `json:"driver,omitempty"`
    // The optional ID of the instance pool for the driver of the cluster
    // belongs. The pool cluster uses the instance pool with id
    // (instance_pool_id) if the driver pool is not assigned.
    DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
    // The node type of the Spark driver. Note that this field is optional; if
    // unset, the driver node type will be set as the same value as
    // `node_type_id` defined above. This field, along with node_type_id, should
    // not be set if virtual_cluster_size is set. If both driver_node_type_id,
    // node_type_id, and virtual_cluster_size are specified, driver_node_type_id
    // and node_type_id take precedence.
    DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
    // The key of the spark version running in the dataplane. This is possibly
    // different from the spark_version (index 2). The spark_version is the raw
    // string provided by the user through API or UI, which could map to a
    // different effective_spark_version running in the dataplane, depending on
    // the cluster&#39;s instance type or the runtimeEngine parameter.
    EffectiveSparkVersion string `json:"effective_spark_version,omitempty"`
    // Autoscaling Local Storage: when enabled, this cluster will dynamically
    // acquire additional disk space when its Spark workers are running low on
    // disk space. This feature requires specific AWS permissions to function
    // correctly - refer to the User Guide for more details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Whether to enable LUKS on cluster VMs&#39; local disks
    EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
    // Nodes on which the Spark executors reside.
    Executors []SparkNode `json:"executors,omitempty"`
    // Attributes related to clusters running on Google Cloud Platform. If not
    // specified at cluster creation, a set of default values will be used.
    GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
    // The optional ID of the instance pool to which the cluster belongs.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // Port on which Spark JDBC server is listening, in the driver nod. No
    // service will be listeningon on this port in executor nodes.
    JdbcPort int `json:"jdbc_port,omitempty"`
    // the timestamp that the cluster was started/restarted
    LastRestartedTime int64 `json:"last_restarted_time,omitempty"`
    // Time when the cluster driver last lost its state (due to a restart or
    // driver failure).
    LastStateLossTime int64 `json:"last_state_loss_time,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this cluster. For example, the Spark nodes can
    // be provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call. This field, along
    // with driver_node_type_id, should not be set if virtual_cluster_size is
    // set. If both driver_node_type_id, node_type_id, and virtual_cluster_size
    // are specified, driver_node_type_id and node_type_id take precedence.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Number of worker nodes that this cluster should have. A cluster has one
    // Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
    // &#43; 1 Spark nodes. Note: When reading the properties of a cluster, this
    // field reflects the desired number of workers rather than the actual
    // current number of workers. For instance, if a cluster is resized from 5
    // to 10 workers, this field will immediately be updated to reflect the
    // target size of 10 workers, whereas the workers listed in ``spark_info``
    // will gradually increase from 5 to 10 as the new nodes are provisioned.
    NumWorkers int `json:"num_workers,omitempty"`
    // The ID of the cluster policy used to create the cluster if applicable.
    PolicyId string `json:"policy_id,omitempty"`
    // Decides which runtime engine to be use, e.g. Standard vs. Photon. If
    // unspecified, the runtime engine is inferred from spark_version.
    RuntimeEngine GetClusterResponseRuntimeEngine `json:"runtime_engine,omitempty"`
    // An object containing a set of optional, user-specified Spark
    // configuration key-value pairs. Users can also pass in a string of extra
    // JVM options to the driver and the executors via
    // ``spark.driver.extraJavaOptions`` and ``spark.executor.extraJavaOptions``
    // respectively. Example Spark confs: ``{&#34;spark.speculation&#34;: true,
    // &#34;spark.streaming.ui.retainedBatches&#34;: 5}`` or
    // ``{&#34;spark.driver.extraJavaOptions&#34;: &#34;-verbose:gc -XX:&#43;PrintGCDetails&#34;}``
    SparkConf map[string]string `json:"spark_conf,omitempty"`
    // A canonical SparkContext identifier. This value *does* change when the
    // Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
    // globally unique identifier over all Spark contexts.
    SparkContextId int64 `json:"spark_context_id,omitempty"`
    // An object containing a set of optional, user-specified environment
    // variable key-value pairs. Please note that key-value pair of the form
    // (X,Y) will be exported as is (i.e., ``export X=&#39;Y&#39;``) while launching the
    // driver and workers. In order to specify an additional set of
    // ``SPARK_DAEMON_JAVA_OPTS``, we recommend appending them to
    // ``$SPARK_DAEMON_JAVA_OPTS`` as shown in the example below. This ensures
    // that all default databricks managed environmental variables are included
    // as well. Example Spark environment variables: ``{&#34;SPARK_WORKER_MEMORY&#34;:
    // &#34;28000m&#34;, &#34;SPARK_LOCAL_DIRS&#34;: &#34;/local_disk0&#34;}`` or
    // ``{&#34;SPARK_DAEMON_JAVA_OPTS&#34;: &#34;$SPARK_DAEMON_JAVA_OPTS
    // -Dspark.shuffle.service.enabled=true&#34;}``
    SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
    // The Spark version of the cluster, e.g. &#34;3.3.x-scala2.11&#34;. A list of
    // available Spark versions can be retrieved by using the
    // :ref:`clusterClusterServicelistSparkVersions` API call. This is the Spark
    // version provided from the user input (API/UI) and may be different from
    // the effective_spark_version, which is the spark version that is actually
    // run in the dataplane. See index 45 for more context.
    SparkVersion string `json:"spark_version,omitempty"`
    // SSH public key contents that will be added to each Spark node in this
    // cluster. The corresponding private keys can be used to login with the
    // user name ``ubuntu`` on port ``2200``. Up to 10 keys can be specified.
    SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
    // Time (in epoch milliseconds) when the cluster creation request was
    // received (when the cluster entered a ``PENDING`` state).
    StartTime int64 `json:"start_time,omitempty"`
    // Current state of the cluster.
    State GetClusterResponseState `json:"state,omitempty"`
    // A message associated with the most recent state transition (e.g., the
    // reason why the cluster entered a ``TERMINATED`` state).
    StateMessage string `json:"state_message,omitempty"`
    // Time (in epoch milliseconds) when the cluster was terminated, if
    // applicable.
    TerminatedTime int64 `json:"terminated_time,omitempty"`
    // Information about why the cluster was terminated. This field only appears
    // when the cluster is in a ``TERMINATING`` or ``TERMINATED`` state.
    TerminationReason *TerminationReason `json:"termination_reason,omitempty"`
    
    WorkloadType *WorkloadType `json:"workload_type,omitempty"`
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.

type GetClusterResponseClusterSource string
const (
	GetClusterResponseClusterSourceUi GetClusterResponseClusterSource = `UI`
	GetClusterResponseClusterSourceJob GetClusterResponseClusterSource = `JOB`
	GetClusterResponseClusterSourceApi GetClusterResponseClusterSource = `API`
	GetClusterResponseClusterSourceSql GetClusterResponseClusterSource = `SQL`
	GetClusterResponseClusterSourceModels GetClusterResponseClusterSource = `MODELS`
	GetClusterResponseClusterSourcePipeline GetClusterResponseClusterSource = `PIPELINE`
	GetClusterResponseClusterSourcePipelineMaintenance GetClusterResponseClusterSource = `PIPELINE_MAINTENANCE`
)

// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
// unspecified, the runtime engine is inferred from spark_version.

type GetClusterResponseRuntimeEngine string
const (
	GetClusterResponseRuntimeEngineNull GetClusterResponseRuntimeEngine = `NULL`
	GetClusterResponseRuntimeEngineStandard GetClusterResponseRuntimeEngine = `STANDARD`
	GetClusterResponseRuntimeEnginePhoton GetClusterResponseRuntimeEngine = `PHOTON`
)

// Current state of the cluster.

type GetClusterResponseState string
const (
	GetClusterResponseStatePending GetClusterResponseState = `PENDING`
	GetClusterResponseStateRunning GetClusterResponseState = `RUNNING`
	GetClusterResponseStateRestarting GetClusterResponseState = `RESTARTING`
	GetClusterResponseStateResizing GetClusterResponseState = `RESIZING`
	GetClusterResponseStateTerminating GetClusterResponseState = `TERMINATING`
	GetClusterResponseStateTerminated GetClusterResponseState = `TERMINATED`
	GetClusterResponseStateError GetClusterResponseState = `ERROR`
	GetClusterResponseStateUnknown GetClusterResponseState = `UNKNOWN`
)


type GetEvents struct {
    // The ID of the cluster to retrieve events about.
    ClusterId string `json:"cluster_id"`
    // The end time in epoch milliseconds. If empty, returns events up to the
    // current time.
    EndTime int64 `json:"end_time,omitempty"`
    // An optional set of event types to filter on. If empty, all event types
    // are returned.
    EventTypes []GetEventsEventTypesItem `json:"event_types,omitempty"`
    // The maximum number of events to include in a page of events. Defaults to
    // 50, and maximum allowed value is 500.
    Limit int64 `json:"limit,omitempty"`
    // The offset in the result set. Defaults to 0 (no offset). When an offset
    // is specified and the results are requested in descending order, the
    // end_time field is required.
    Offset int64 `json:"offset,omitempty"`
    // The order to list events in; either &#34;ASC&#34; or &#34;DESC&#34;. Defaults to &#34;DESC&#34;.
    Order GetEventsOrder `json:"order,omitempty"`
    // The start time in epoch milliseconds. If empty, returns events starting
    // from the beginning of time.
    StartTime int64 `json:"start_time,omitempty"`
}



type GetEventsEventTypesItem string
const (
	GetEventsEventTypesItemCreating GetEventsEventTypesItem = `CREATING`
	GetEventsEventTypesItemStarting GetEventsEventTypesItem = `STARTING`
	GetEventsEventTypesItemRestarting GetEventsEventTypesItem = `RESTARTING`
	GetEventsEventTypesItemTerminating GetEventsEventTypesItem = `TERMINATING`
	GetEventsEventTypesItemEdited GetEventsEventTypesItem = `EDITED`
	GetEventsEventTypesItemRunning GetEventsEventTypesItem = `RUNNING`
	GetEventsEventTypesItemResizing GetEventsEventTypesItem = `RESIZING`
	GetEventsEventTypesItemNodesLost GetEventsEventTypesItem = `NODES_LOST`
	GetEventsEventTypesItemUpsizeCompleted GetEventsEventTypesItem = `UPSIZE_COMPLETED`
	GetEventsEventTypesItemInitScriptsStarted GetEventsEventTypesItem = `INIT_SCRIPTS_STARTED`
	GetEventsEventTypesItemInitScriptsFinished GetEventsEventTypesItem = `INIT_SCRIPTS_FINISHED`
	GetEventsEventTypesItemDidNotExpandDisk GetEventsEventTypesItem = `DID_NOT_EXPAND_DISK`
	GetEventsEventTypesItemExpandedDisk GetEventsEventTypesItem = `EXPANDED_DISK`
	GetEventsEventTypesItemFailedToExpandDisk GetEventsEventTypesItem = `FAILED_TO_EXPAND_DISK`
	GetEventsEventTypesItemDriverHealthy GetEventsEventTypesItem = `DRIVER_HEALTHY`
	GetEventsEventTypesItemDriverNotResponding GetEventsEventTypesItem = `DRIVER_NOT_RESPONDING`
	GetEventsEventTypesItemDriverUnavailable GetEventsEventTypesItem = `DRIVER_UNAVAILABLE`
	GetEventsEventTypesItemSparkException GetEventsEventTypesItem = `SPARK_EXCEPTION`
	GetEventsEventTypesItemMetastoreDown GetEventsEventTypesItem = `METASTORE_DOWN`
	GetEventsEventTypesItemDbfsDown GetEventsEventTypesItem = `DBFS_DOWN`
	GetEventsEventTypesItemAutoscalingStatsReport GetEventsEventTypesItem = `AUTOSCALING_STATS_REPORT`
	GetEventsEventTypesItemNodeBlacklisted GetEventsEventTypesItem = `NODE_BLACKLISTED`
	GetEventsEventTypesItemPinned GetEventsEventTypesItem = `PINNED`
	GetEventsEventTypesItemUnpinned GetEventsEventTypesItem = `UNPINNED`
	GetEventsEventTypesItemNodeExcludedDecommissioned GetEventsEventTypesItem = `NODE_EXCLUDED_DECOMMISSIONED`
)

// The order to list events in; either &#34;ASC&#34; or &#34;DESC&#34;. Defaults to &#34;DESC&#34;.

type GetEventsOrder string
const (
	GetEventsOrderDesc GetEventsOrder = `DESC`
	GetEventsOrderAsc GetEventsOrder = `ASC`
)


type GetEventsRequest struct {
    // The ID of the cluster to retrieve events about.
    ClusterId string `json:"cluster_id"`
    // The end time in epoch milliseconds. If empty, returns events up to the
    // current time.
    EndTime int64 `json:"end_time,omitempty"`
    // An optional set of event types to filter on. If empty, all event types
    // are returned.
    EventTypes []GetEventsRequestEventTypesItem `json:"event_types,omitempty"`
    // The maximum number of events to include in a page of events. Defaults to
    // 50, and maximum allowed value is 500.
    Limit int64 `json:"limit,omitempty"`
    // The offset in the result set. Defaults to 0 (no offset). When an offset
    // is specified and the results are requested in descending order, the
    // end_time field is required.
    Offset int64 `json:"offset,omitempty"`
    // The order to list events in; either &#34;ASC&#34; or &#34;DESC&#34;. Defaults to &#34;DESC&#34;.
    Order GetEventsRequestOrder `json:"order,omitempty"`
    // The start time in epoch milliseconds. If empty, returns events starting
    // from the beginning of time.
    StartTime int64 `json:"start_time,omitempty"`
}



type GetEventsRequestEventTypesItem string
const (
	GetEventsRequestEventTypesItemCreating GetEventsRequestEventTypesItem = `CREATING`
	GetEventsRequestEventTypesItemStarting GetEventsRequestEventTypesItem = `STARTING`
	GetEventsRequestEventTypesItemRestarting GetEventsRequestEventTypesItem = `RESTARTING`
	GetEventsRequestEventTypesItemTerminating GetEventsRequestEventTypesItem = `TERMINATING`
	GetEventsRequestEventTypesItemEdited GetEventsRequestEventTypesItem = `EDITED`
	GetEventsRequestEventTypesItemRunning GetEventsRequestEventTypesItem = `RUNNING`
	GetEventsRequestEventTypesItemResizing GetEventsRequestEventTypesItem = `RESIZING`
	GetEventsRequestEventTypesItemNodesLost GetEventsRequestEventTypesItem = `NODES_LOST`
	GetEventsRequestEventTypesItemUpsizeCompleted GetEventsRequestEventTypesItem = `UPSIZE_COMPLETED`
	GetEventsRequestEventTypesItemInitScriptsStarted GetEventsRequestEventTypesItem = `INIT_SCRIPTS_STARTED`
	GetEventsRequestEventTypesItemInitScriptsFinished GetEventsRequestEventTypesItem = `INIT_SCRIPTS_FINISHED`
	GetEventsRequestEventTypesItemDidNotExpandDisk GetEventsRequestEventTypesItem = `DID_NOT_EXPAND_DISK`
	GetEventsRequestEventTypesItemExpandedDisk GetEventsRequestEventTypesItem = `EXPANDED_DISK`
	GetEventsRequestEventTypesItemFailedToExpandDisk GetEventsRequestEventTypesItem = `FAILED_TO_EXPAND_DISK`
	GetEventsRequestEventTypesItemDriverHealthy GetEventsRequestEventTypesItem = `DRIVER_HEALTHY`
	GetEventsRequestEventTypesItemDriverNotResponding GetEventsRequestEventTypesItem = `DRIVER_NOT_RESPONDING`
	GetEventsRequestEventTypesItemDriverUnavailable GetEventsRequestEventTypesItem = `DRIVER_UNAVAILABLE`
	GetEventsRequestEventTypesItemSparkException GetEventsRequestEventTypesItem = `SPARK_EXCEPTION`
	GetEventsRequestEventTypesItemMetastoreDown GetEventsRequestEventTypesItem = `METASTORE_DOWN`
	GetEventsRequestEventTypesItemDbfsDown GetEventsRequestEventTypesItem = `DBFS_DOWN`
	GetEventsRequestEventTypesItemAutoscalingStatsReport GetEventsRequestEventTypesItem = `AUTOSCALING_STATS_REPORT`
	GetEventsRequestEventTypesItemNodeBlacklisted GetEventsRequestEventTypesItem = `NODE_BLACKLISTED`
	GetEventsRequestEventTypesItemPinned GetEventsRequestEventTypesItem = `PINNED`
	GetEventsRequestEventTypesItemUnpinned GetEventsRequestEventTypesItem = `UNPINNED`
	GetEventsRequestEventTypesItemNodeExcludedDecommissioned GetEventsRequestEventTypesItem = `NODE_EXCLUDED_DECOMMISSIONED`
)

// The order to list events in; either &#34;ASC&#34; or &#34;DESC&#34;. Defaults to &#34;DESC&#34;.

type GetEventsRequestOrder string
const (
	GetEventsRequestOrderDesc GetEventsRequestOrder = `DESC`
	GetEventsRequestOrderAsc GetEventsRequestOrder = `ASC`
)


type GetEventsResponse struct {
    
    Events []ClusterEvent `json:"events,omitempty"`
    // The parameters required to retrieve the next page of events. Omitted if
    // there are no more events to read.
    NextPage *GetEvents `json:"next_page,omitempty"`
    // The total number of events filtered by the start_time, end_time, and
    // event_types.
    TotalCount int64 `json:"total_count,omitempty"`
}


type GetSparkVersionsResponse struct {
    // All the available Spark versions.
    Versions []SparkVersion `json:"versions,omitempty"`
}


type ListAvailableZonesResponse struct {
    // The availability zone if no ``zone_id`` is provided in the cluster
    // creation request.
    DefaultZone string `json:"default_zone,omitempty"`
    // The list of available zones (e.g., [&#39;us-west-2c&#39;, &#39;us-east-2&#39;]).
    Zones []string `json:"zones,omitempty"`
}


type ListClustersRequest struct {
    // Filter clusters based on what type of client it can be used for. Could be
    // either NOTEBOOKS or JOBS. No input for this field will get all clusters
    // in the workspace without filtering on its supported client
    CanUseClient string ` url:"can_use_client,omitempty"`
}


type ListClustersResponse struct {
    
    Clusters []ClusterInfo `json:"clusters,omitempty"`
}


type ListNodeTypesResponse struct {
    
    Failure *CpalFailureResponse `json:"failure,omitempty"`
    // The list of available Spark node types.
    NodeTypes []NodeType `json:"node_types,omitempty"`
    
    Success any /* ERROR */ `json:"success,omitempty"`
}


type LogAnalyticsInfo struct {
    
    LogAnalyticsPrimaryKey string `json:"log_analytics_primary_key,omitempty"`
    
    LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`
}


type LogSyncStatus struct {
    // The timestamp of last attempt. If the last attempt fails,
    // ``last_exception`` will contain the exception in the last attempt.
    LastAttempted int64 `json:"last_attempted,omitempty"`
    // The exception thrown in the last attempt, it would be null (omitted in
    // the response) if there is no exception in last attempted.
    LastException string `json:"last_exception,omitempty"`
}


type NodeType struct {
    // A string description associated with this node type, e.g., &#34;r3.xlarge&#34;.
    Description string `json:"description"`
    // An identifier for the type of hardware that this node runs on, e.g.,
    // &#34;r3.2xlarge&#34; in AWS.
    InstanceTypeId string `json:"instance_type_id"`
    // Whether the node type is deprecated. Non-deprecated node types offer
    // greater performance.
    IsDeprecated bool `json:"is_deprecated,omitempty"`
    // AWS specific, whether this instance supports encryption in transit, used
    // for hipaa and pci workloads.
    IsEncryptedInTransit bool `json:"is_encrypted_in_transit,omitempty"`
    // Memory (in MB) available for this node type.
    MemoryMb int `json:"memory_mb"`
    // Unique identifier for this node type.
    NodeTypeId string `json:"node_type_id"`
    // Number of CPU cores available for this node type. Note that this can be
    // fractional, e.g., 2.5 cores, if the the number of cores on a machine
    // instance is not divisible by the number of Spark nodes on that machine.
    NumCores float64 `json:"num_cores"`
}


type PermanentDeleteClusterRequest struct {
    // The cluster to be deleted.
    ClusterId string `json:"cluster_id"`
}


type PinClusterRequest struct {
    
    ClusterId string `json:"cluster_id"`
}


type ResizeClusterRequest struct {
    // Parameters needed in order to automatically scale clusters up and down
    // based on load. Note: autoscaling works best with DB runtime versions 3.0
    // or later.
    Autoscale *AutoScale `json:"autoscale,omitempty"`
    // The cluster to be resized.
    ClusterId string `json:"cluster_id"`
    // Number of worker nodes that this cluster should have. A cluster has one
    // Spark Driver and ``num_workers`` Executors for a total of ``num_workers``
    // &#43; 1 Spark nodes. Note: When reading the properties of a cluster, this
    // field reflects the desired number of workers rather than the actual
    // current number of workers. For instance, if a cluster is resized from 5
    // to 10 workers, this field will immediately be updated to reflect the
    // target size of 10 workers, whereas the workers listed in ``spark_info``
    // will gradually increase from 5 to 10 as the new nodes are provisioned.
    NumWorkers int `json:"num_workers,omitempty"`
}


type RestartClusterRequest struct {
    // The cluster to be started.
    ClusterId string `json:"cluster_id"`
    
    RestartUser string `json:"restart_user,omitempty"`
}


type S3StorageInfo struct {
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


type SparkNode struct {
    // The private IP address of the host instance.
    HostPrivateIp string `json:"host_private_ip,omitempty"`
    // Globally unique identifier for the host instance from the cloud provider.
    InstanceId string `json:"instance_id,omitempty"`
    // Attributes specific to AWS for a Spark node.
    NodeAwsAttributes *SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
    // Globally unique identifier for this node.
    NodeId string `json:"node_id,omitempty"`
    // Private IP address (typically a 10.x.x.x address) of the Spark node. Note
    // that this is different from the private IP address of the host instance.
    PrivateIp string `json:"private_ip,omitempty"`
    // Public DNS address of this node. This address can be used to access the
    // Spark JDBC server on the driver node. To communicate with the JDBC
    // server, traffic must be manually authorized by adding security group
    // rules to the &#34;worker-unmanaged&#34; security group via the AWS console.
    // Actually it&#39;s the public DNS address of the host instance.
    PublicDns string `json:"public_dns,omitempty"`
    // The timestamp (in millisecond) when the Spark node is launched. The
    // start_timestamp is set right before the container is being launched. The
    // timestamp when the container is placed on the ResourceManager, before its
    // launch and setup by the NodeDaemon. This timestamp is the same as the
    // creation timestamp in the database.
    StartTimestamp int64 `json:"start_timestamp,omitempty"`
}


type SparkNodeAwsAttributes struct {
    // Whether this node is on an Amazon spot instance.
    IsSpot bool `json:"is_spot,omitempty"`
}


type SparkVersion struct {
    // Spark version key, for example &#34;2.1.x-scala2.11&#34;. This is the value which
    // should be provided as the &#34;spark_version&#34; when creating a new cluster.
    // Note that the exact Spark version may change over time for a &#34;wildcard&#34;
    // version (i.e., &#34;2.1.x-scala2.11&#34; is a &#34;wildcard&#34; version) with minor bug
    // fixes.
    Key string `json:"key,omitempty"`
    // A descriptive name for this Spark version, for example &#34;Spark 2.1&#34;.
    Name string `json:"name,omitempty"`
}


type StartClusterRequest struct {
    // The cluster to be started.
    ClusterId string `json:"cluster_id"`
}


type TerminationReason struct {
    // status code indicating why the cluster was terminated
    Code TerminationReasonCode `json:"code,omitempty"`
    // list of parameters that provide additional information about why the
    // cluster was terminated
    Parameters map[string]string `json:"parameters,omitempty"`
    // type of the termination
    Type TerminationReasonType `json:"type,omitempty"`
}

// status code indicating why the cluster was terminated

type TerminationReasonCode string
const (
	TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`
	TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`
	TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`
	TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`
	TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`
	TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`
	TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`
	TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`
	TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`
	TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`
	TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`
	TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`
	TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`
	TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`
	TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`
	TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`
	TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`
	TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`
	TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`
	TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`
	TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`
	TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`
	TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`
	TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`
	TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`
	TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`
	TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`
	TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`
	TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`
	TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`
	TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`
	TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`
	TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`
	TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`
	TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`
	TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`
	TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`
	TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`
	TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`
	TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`
	TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`
	TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`
	TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`
	TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`
	TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`
	TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`
	TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`
	TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`
	TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`
	TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`
	TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`
	TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`
	TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`
	TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`
	TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`
	TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`
	TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`
	TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`
	TerminationReasonCodeKSAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`
	TerminationReasonCodeKSDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`
	TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`
	TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`
	TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`
	TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`
	TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`
	TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`
	TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`
	TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`
	TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`
	TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`
	TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`
	TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`
	TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`
	TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`
	TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`
	TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`
	TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`
	TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`
	TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`
)

// type of the termination

type TerminationReasonType string
const (
	TerminationReasonTypeSuccess TerminationReasonType = `SUCCESS`
	TerminationReasonTypeClientError TerminationReasonType = `CLIENT_ERROR`
	TerminationReasonTypeServiceFault TerminationReasonType = `SERVICE_FAULT`
	TerminationReasonTypeCloudFailure TerminationReasonType = `CLOUD_FAILURE`
)


type UnpinClusterRequest struct {
    
    ClusterId string `json:"cluster_id"`
}


type WorkloadType struct {
    // defined what type of clients can use the cluster. E.g. Notebooks, Jobs
    Clients *ClientsTypes `json:"clients,omitempty"`
}
