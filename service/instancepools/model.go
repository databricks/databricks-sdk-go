// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instancepools

// all definitions in this file are in alphabetical order

type CreateInstancePoolRequest struct {
    // Attributes related to pool running on Amazon Web Services. If not
    // specified at pool creation, a set of default values will be used.
    AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to pool running on Azure. If not specified at pool
    // creation, a set of default values will be used.
    AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
    // Additional tags for pool resources. Databricks will tag all pool
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // Defines the specification of the disks that will be attached to all spark
    // containers.
    DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
    // Autoscaling Local Storage: when enabled, this instances in this pool will
    // dynamically acquire additional disk space when its Spark workers are
    // running low on disk space. In AWS, this feature requires specific AWS
    // permissions to function correctly - refer to the User Guide for more
    // details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Automatically terminates the extra instances in the pool cache after they
    // are inactive for this time in minutes if min_idle_instances requirement
    // is already met. If not set, the extra pool instances will be
    // automatically terminated after a default timeout. If specified, the
    // threshold must be between 0 and 10000 minutes. Users can also set this
    // value to 0 to instantly remove idle instances from the cache if min cache
    // size could still hold.
    IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
    // The fleet related setting to power the intance pool. Note that we don&#39;t
    // want to inline this message as it can be hard to interpret/manage
    InstancePoolFleetAttributes *InstancePoolFleetAttributes `json:"instance_pool_fleet_attributes,omitempty"`
    // Pool name requested by the user. This has to be unique. Length must be
    // between 1 and 100 characters.
    InstancePoolName string `json:"instance_pool_name,omitempty"`
    // Maximum number of outstanding instances to keep in the pool, including
    // both instances used by clusters and idle ones. Clusters that require
    // further instance provision would fail during upsize requests.
    MaxCapacity int `json:"max_capacity,omitempty"`
    // Minimum number of idle instances to keep in the instance pool
    MinIdleInstances int `json:"min_idle_instances,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this pool. For example, the Spark nodes can be
    // provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Custom Docker Image BYOC
    PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
    // A list of preloaded Spark image versions for the pool, e.g.
    // [&#34;5.2.x-scala2.11&#34;]. Pool-backed clusters started with the preloaded
    // Spark version will start faster. A list of available Spark versions can
    // be retrieved by using the :ref:`clusterClusterServicelistSparkVersions`
    // API call.
    PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
}


type CreateInstancePoolResponse struct {
    
    InstancePoolId string `json:"instance_pool_id,omitempty"`
}


type DeleteInstancePoolRequest struct {
    // The instance pool to be terminated.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
}


type DiskSpec struct {
    // The number of disks launched for each instance: - This feature is only
    // enabled for supported node types. - Users can choose up to the limit of
    // the disks supported by the node type. - For node types with no OS disk,
    // at least one disk needs to be specified; otherwise, cluster creation will
    // fail. If disks are attached, Databricks will configure Spark to use only
    // the disks for scratch storage because heterogenously sized scratch
    // devices can lead to inefficient disk utilization. If no disks are
    // attached, Databricks will configure Spark to use instance store disks.
    // Please note that if disks are specified, then the Spark configuration
    // ``spark.local.dir`` will be overridden. Disks will be mounted at: - For
    // AWS: ``/ebs0``, ``/ebs1``, and etc. - For Azure: ``/remote_volume0``,
    // ``/remote_volume1``, and etc.
    DiskCount int `json:"disk_count,omitempty"`
    
    DiskIops int `json:"disk_iops,omitempty"`
    // The size of each disk (in GiB) launched for each instance. Values must
    // fall into the supported range for a particular instance type. For AWS: -
    // General Purpose SSD: 100 - 4096 GiB - Throughput Optimized HDD: 500 -
    // 4096 GiB For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS
    // (HDD): 1- 1023 GiB
    DiskSize int `json:"disk_size,omitempty"`
    
    DiskThroughput int `json:"disk_throughput,omitempty"`
    // The type of disks that will be launched with this cluster.
    DiskType *DiskType `json:"disk_type,omitempty"`
}


type DiskType struct {
    
    AzureDiskVolumeType DiskTypeAzureDiskVolumeType `json:"azure_disk_volume_type,omitempty"`
    
    EbsVolumeType DiskTypeEbsVolumeType `json:"ebs_volume_type,omitempty"`
}


type DiskTypeAzureDiskVolumeType string


const DiskTypeAzureDiskVolumeTypePremiumLrs DiskTypeAzureDiskVolumeType = `PREMIUM_LRS`

const DiskTypeAzureDiskVolumeTypeStandardLrs DiskTypeAzureDiskVolumeType = `STANDARD_LRS`

type DiskTypeEbsVolumeType string


const DiskTypeEbsVolumeTypeGeneralPurposeSsd DiskTypeEbsVolumeType = `GENERAL_PURPOSE_SSD`

const DiskTypeEbsVolumeTypeThroughputOptimizedHdd DiskTypeEbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`

type DockerBasicAuth struct {
    
    Password string `json:"password,omitempty"`
    
    Username string `json:"username,omitempty"`
}


type DockerImage struct {
    
    BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
    // URL of the docker image.
    Url string `json:"url,omitempty"`
}


type EditInstancePoolRequest struct {
    // Attributes related to pool running on Amazon Web Services. If not
    // specified at pool creation, a set of default values will be used.
    AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to pool running on Azure. If not specified at pool
    // creation, a set of default values will be used.
    AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
    // Additional tags for pool resources. Databricks will tag all pool
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // Defines the specification of the disks that will be attached to all spark
    // containers.
    DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
    // Autoscaling Local Storage: when enabled, this instances in this pool will
    // dynamically acquire additional disk space when its Spark workers are
    // running low on disk space. In AWS, this feature requires specific AWS
    // permissions to function correctly - refer to the User Guide for more
    // details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Automatically terminates the extra instances in the pool cache after they
    // are inactive for this time in minutes if min_idle_instances requirement
    // is already met. If not set, the extra pool instances will be
    // automatically terminated after a default timeout. If specified, the
    // threshold must be between 0 and 10000 minutes. Users can also set this
    // value to 0 to instantly remove idle instances from the cache if min cache
    // size could still hold.
    IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
    
    InstancePoolId string `json:"instance_pool_id"`
    // Pool name requested by the user. This has to be unique. Length must be
    // between 1 and 100 characters.
    InstancePoolName string `json:"instance_pool_name,omitempty"`
    // Maximum number of outstanding instances to keep in the pool, including
    // both instances used by clusters and idle ones. Clusters that require
    // further instance provision would fail during upsize requests.
    MaxCapacity int `json:"max_capacity,omitempty"`
    // Minimum number of idle instances to keep in the instance pool
    MinIdleInstances int `json:"min_idle_instances,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this pool. For example, the Spark nodes can be
    // provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Custom Docker Image BYOC
    PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
    // A list of preloaded Spark image versions for the pool, e.g.
    // [&#34;5.2.x-scala2.11&#34;]. Pool-backed clusters started with the preloaded
    // Spark version will start faster. A list of available Spark versions can
    // be retrieved by using the :ref:`clusterClusterServicelistSparkVersions`
    // API call.
    PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
}


type FleetLaunchTemplateOverride struct {
    // User assigned preferred availability zone. It will be adjust to the
    // default zone of the worker environment if the preferred zone does not
    // exist in the subnet.
    AvailabilityZone string `json:"availability_zone"`
    
    InstanceType string `json:"instance_type"`
    // The maximum price per unit hour that you are willing to pay for a Spot
    // Instance.
    MaxPrice float64 `json:"max_price,omitempty"`
    // The priority for the launch template override. If AllocationStrategy is
    // set to prioritized, EC2 Fleet uses priority to determine which launch
    // template override to use first in fulfilling On-Demand capacity. The
    // highest priority is launched first. Valid values are whole numbers
    // starting at 0. The lower the number, the higher the priority. If no
    // number is set, the launch template override has the lowest priority.
    Priority float64 `json:"priority,omitempty"`
}


type FleetOnDemandOption struct {
    // Only lowest-price and prioritized are allowed
    AllocationStrategy FleetOnDemandOptionAllocationStrategy `json:"allocation_strategy,omitempty"`
    // The maximum amount per hour for On-Demand Instances that you&#39;re willing
    // to pay.
    MaxTotalPrice float64 `json:"max_total_price,omitempty"`
    // If you specify use-capacity-reservations-first, the fleet uses unused
    // Capacity Reservations to fulfill On-Demand capacity up to the target
    // On-Demand capacity. If multiple instance pools have unused Capacity
    // Reservations, the On-Demand allocation strategy (lowest-price or
    // prioritized) is applied. If the number of unused Capacity Reservations is
    // less than the On-Demand target capacity, the remaining On-Demand target
    // capacity is launched according to the On-Demand allocation strategy
    // (lowest-price or prioritized).
    UseCapacityReservationsFirst bool `json:"use_capacity_reservations_first,omitempty"`
}

// Only lowest-price and prioritized are allowed
type FleetOnDemandOptionAllocationStrategy string


const FleetOnDemandOptionAllocationStrategyLowestPrice FleetOnDemandOptionAllocationStrategy = `LOWEST_PRICE`

const FleetOnDemandOptionAllocationStrategyDiversified FleetOnDemandOptionAllocationStrategy = `DIVERSIFIED`

const FleetOnDemandOptionAllocationStrategyCapacityOptimized FleetOnDemandOptionAllocationStrategy = `CAPACITY_OPTIMIZED`

const FleetOnDemandOptionAllocationStrategyPrioritized FleetOnDemandOptionAllocationStrategy = `PRIORITIZED`

type FleetSpotOption struct {
    // lowest-price | diversified | capaacityOptimized
    AllocationStrategy FleetSpotOptionAllocationStrategy `json:"allocation_strategy,omitempty"`
    // The number of Spot pools across which to allocate your target Spot
    // capacity. Valid only when Spot AllocationStrategy is set to lowest-price.
    // EC2 Fleet selects the cheapest Spot pools and evenly allocates your
    // target Spot capacity across the number of Spot pools that you specify.
    InstancePoolsToUseCount int `json:"instance_pools_to_use_count,omitempty"`
    // The maximum amount per hour for Spot Instances that you&#39;re willing to
    // pay.
    MaxTotalPrice float64 `json:"max_total_price,omitempty"`
}

// lowest-price | diversified | capaacityOptimized
type FleetSpotOptionAllocationStrategy string


const FleetSpotOptionAllocationStrategyLowestPrice FleetSpotOptionAllocationStrategy = `LOWEST_PRICE`

const FleetSpotOptionAllocationStrategyDiversified FleetSpotOptionAllocationStrategy = `DIVERSIFIED`

const FleetSpotOptionAllocationStrategyCapacityOptimized FleetSpotOptionAllocationStrategy = `CAPACITY_OPTIMIZED`

const FleetSpotOptionAllocationStrategyPrioritized FleetSpotOptionAllocationStrategy = `PRIORITIZED`

type GetInstancePoolRequest struct {
    // The instance pool about which to retrieve information.
    InstancePoolId string ` url:"instance_pool_id,omitempty"`
}


type GetInstancePoolResponse struct {
    // Attributes related to pool running on Amazon Web Services. If not
    // specified at pool creation, a set of default values will be used.
    AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to pool running on Azure. If not specified at pool
    // creation, a set of default values will be used.
    AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
    // Additional tags for pool resources. Databricks will tag all pool
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // Tags that are added by Databricks regardless of any ``custom_tags``,
    // including: - Vendor: Databricks - InstancePoolCreator:
    // &lt;user_id_of_creator&gt; - InstancePoolName: &lt;name_of_pool&gt; - InstancePoolId:
    // &lt;id_of_pool&gt;
    DefaultTags map[string]string `json:"default_tags,omitempty"`
    // Defines the specification of the disks that will be attached to all spark
    // containers.
    DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
    // Autoscaling Local Storage: when enabled, this instances in this pool will
    // dynamically acquire additional disk space when its Spark workers are
    // running low on disk space. In AWS, this feature requires specific AWS
    // permissions to function correctly - refer to the User Guide for more
    // details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Automatically terminates the extra instances in the pool cache after they
    // are inactive for this time in minutes if min_idle_instances requirement
    // is already met. If not set, the extra pool instances will be
    // automatically terminated after a default timeout. If specified, the
    // threshold must be between 0 and 10000 minutes. Users can also set this
    // value to 0 to instantly remove idle instances from the cache if min cache
    // size could still hold.
    IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
    // Canonical unique identifier for the pool.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // Pool name requested by the user. This has to be unique. Length must be
    // between 1 and 100 characters.
    InstancePoolName string `json:"instance_pool_name,omitempty"`
    // Maximum number of outstanding instances to keep in the pool, including
    // both instances used by clusters and idle ones. Clusters that require
    // further instance provision would fail during upsize requests.
    MaxCapacity int `json:"max_capacity,omitempty"`
    // Minimum number of idle instances to keep in the instance pool
    MinIdleInstances int `json:"min_idle_instances,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this pool. For example, the Spark nodes can be
    // provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Custom Docker Image BYOC
    PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
    // A list of preloaded Spark image versions for the pool, e.g.
    // [&#34;5.2.x-scala2.11&#34;]. Pool-backed clusters started with the preloaded
    // Spark version will start faster. A list of available Spark versions can
    // be retrieved by using the :ref:`clusterClusterServicelistSparkVersions`
    // API call.
    PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
    // Current state of the instance pool.
    State GetInstancePoolResponseState `json:"state,omitempty"`
    
    Stats *InstancePoolStats `json:"stats,omitempty"`
    
    Status *InstancePoolStatus `json:"status,omitempty"`
}

// Current state of the instance pool.
type GetInstancePoolResponseState string


const GetInstancePoolResponseStateActive GetInstancePoolResponseState = `ACTIVE`

const GetInstancePoolResponseStateStopped GetInstancePoolResponseState = `STOPPED`

const GetInstancePoolResponseStateDeleted GetInstancePoolResponseState = `DELETED`

type InstancePoolAndStats struct {
    // Attributes related to pool running on Amazon Web Services. If not
    // specified at pool creation, a set of default values will be used.
    AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
    // Attributes related to pool running on Azure. If not specified at pool
    // creation, a set of default values will be used.
    AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
    // Additional tags for pool resources. Databricks will tag all pool
    // resources (e.g., AWS instances and EBS volumes) with these tags in
    // addition to ``default_tags``. Notes: - Currently, Databricks allows at
    // most 45 custom tags
    CustomTags map[string]string `json:"custom_tags,omitempty"`
    // Tags that are added by Databricks regardless of any ``custom_tags``,
    // including: - Vendor: Databricks - InstancePoolCreator:
    // &lt;user_id_of_creator&gt; - InstancePoolName: &lt;name_of_pool&gt; - InstancePoolId:
    // &lt;id_of_pool&gt;
    DefaultTags map[string]string `json:"default_tags,omitempty"`
    // Defines the specification of the disks that will be attached to all spark
    // containers.
    DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
    // Autoscaling Local Storage: when enabled, this instances in this pool will
    // dynamically acquire additional disk space when its Spark workers are
    // running low on disk space. In AWS, this feature requires specific AWS
    // permissions to function correctly - refer to the User Guide for more
    // details.
    EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
    // Automatically terminates the extra instances in the pool cache after they
    // are inactive for this time in minutes if min_idle_instances requirement
    // is already met. If not set, the extra pool instances will be
    // automatically terminated after a default timeout. If specified, the
    // threshold must be between 0 and 10000 minutes. Users can also set this
    // value to 0 to instantly remove idle instances from the cache if min cache
    // size could still hold.
    IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
    // Canonical unique identifier for the pool.
    InstancePoolId string `json:"instance_pool_id,omitempty"`
    // Pool name requested by the user. This has to be unique. Length must be
    // between 1 and 100 characters.
    InstancePoolName string `json:"instance_pool_name,omitempty"`
    // Maximum number of outstanding instances to keep in the pool, including
    // both instances used by clusters and idle ones. Clusters that require
    // further instance provision would fail during upsize requests.
    MaxCapacity int `json:"max_capacity,omitempty"`
    // Minimum number of idle instances to keep in the instance pool
    MinIdleInstances int `json:"min_idle_instances,omitempty"`
    // This field encodes, through a single value, the resources available to
    // each of the Spark nodes in this pool. For example, the Spark nodes can be
    // provisioned and optimized for memory or compute intensive workloads A
    // list of available node types can be retrieved by using the
    // :ref:`clusterClusterServicelistNodeTypes` API call.
    NodeTypeId string `json:"node_type_id,omitempty"`
    // Custom Docker Image BYOC
    PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
    // A list of preloaded Spark image versions for the pool, e.g.
    // [&#34;5.2.x-scala2.11&#34;]. Pool-backed clusters started with the preloaded
    // Spark version will start faster. A list of available Spark versions can
    // be retrieved by using the :ref:`clusterClusterServicelistSparkVersions`
    // API call.
    PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
    // Current state of the instance pool.
    State InstancePoolAndStatsState `json:"state,omitempty"`
    
    Stats *InstancePoolStats `json:"stats,omitempty"`
    
    Status *InstancePoolStatus `json:"status,omitempty"`
}

// Current state of the instance pool.
type InstancePoolAndStatsState string


const InstancePoolAndStatsStateActive InstancePoolAndStatsState = `ACTIVE`

const InstancePoolAndStatsStateStopped InstancePoolAndStatsState = `STOPPED`

const InstancePoolAndStatsStateDeleted InstancePoolAndStatsState = `DELETED`

type InstancePoolAwsAttributes struct {
    // Availability type used for the spot nodes. The default value is defined
    // by InstancePoolConf.instancePoolDefaultAwsAvailability
    Availability InstancePoolAwsAttributesAvailability `json:"availability,omitempty"`
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
    // will be used. The list of available zones as well as the default value
    // can be found by using the `List Zones`_ method.
    ZoneId string `json:"zone_id,omitempty"`
}

// Availability type used for the spot nodes. The default value is defined by
// InstancePoolConf.instancePoolDefaultAwsAvailability
type InstancePoolAwsAttributesAvailability string


const InstancePoolAwsAttributesAvailabilitySpot InstancePoolAwsAttributesAvailability = `SPOT`

const InstancePoolAwsAttributesAvailabilityOnDemand InstancePoolAwsAttributesAvailability = `ON_DEMAND`

const InstancePoolAwsAttributesAvailabilitySpotWithFallback InstancePoolAwsAttributesAvailability = `SPOT_WITH_FALLBACK`

type InstancePoolAzureAttributes struct {
    // Availability type used for the spot nodes. The default value is defined
    // by InstancePoolConf.instancePoolDefaultAzureAvailability
    Availability InstancePoolAzureAttributesAvailability `json:"availability,omitempty"`
    // The default value and documentation here should be kept consistent with
    // CommonConf.defaultSpotBidMaxPrice.
    SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`
}

// Availability type used for the spot nodes. The default value is defined by
// InstancePoolConf.instancePoolDefaultAzureAvailability
type InstancePoolAzureAttributesAvailability string


const InstancePoolAzureAttributesAvailabilitySpotAzure InstancePoolAzureAttributesAvailability = `SPOT_AZURE`

const InstancePoolAzureAttributesAvailabilityOnDemandAzure InstancePoolAzureAttributesAvailability = `ON_DEMAND_AZURE`

const InstancePoolAzureAttributesAvailabilitySpotWithFallbackAzure InstancePoolAzureAttributesAvailability = `SPOT_WITH_FALLBACK_AZURE`

type InstancePoolFleetAttributes struct {
    
    FleetOnDemandOption *FleetOnDemandOption `json:"fleet_on_demand_option,omitempty"`
    
    FleetSpotOption *FleetSpotOption `json:"fleet_spot_option,omitempty"`
    
    LaunchTemplateOverrides []FleetLaunchTemplateOverride `json:"launch_template_overrides,omitempty"`
}


type InstancePoolStats struct {
    // Number of active instances in the pool that are NOT part of a cluster.
    IdleCount int `json:"idle_count,omitempty"`
    // Number of pending instances in the pool that are NOT part of a cluster.
    PendingIdleCount int `json:"pending_idle_count,omitempty"`
    // Number of pending instances in the pool that are part of a cluster.
    PendingUsedCount int `json:"pending_used_count,omitempty"`
    // Number of active instances in the pool that are part of a cluster.
    UsedCount int `json:"used_count,omitempty"`
}


type InstancePoolStatus struct {
    // List of error messages for the failed pending instances. The
    // pending_instance_errors follows FIFO with maximum length of the min_idle
    // of the pool. The pending_instance_errors is emptied once the number of
    // exiting available instances reaches the min_idle of the pool.
    PendingInstanceErrors []PendingInstanceError `json:"pending_instance_errors,omitempty"`
}


type ListInstancePoolsResponse struct {
    
    InstancePools []InstancePoolAndStats `json:"instance_pools,omitempty"`
}


type PendingInstanceError struct {
    
    InstanceId string `json:"instance_id,omitempty"`
    
    Message string `json:"message,omitempty"`
}

