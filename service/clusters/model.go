package clusters

import (
	"fmt"
	"strings"
)

// AutoScale is a struct the describes auto scaling for clusters
type AutoScale struct {
	MinWorkers int32 `json:"min_workers,omitempty"`
	MaxWorkers int32 `json:"max_workers,omitempty"`
}

// Availability is a type for describing AWS availability on cluster nodes
type Availability string

const (
	// AwsAvailabilitySpot is spot instance type for clusters
	AwsAvailabilitySpot = "SPOT"
	// AwsAvailabilityOnDemand is OnDemand instance type for clusters
	AwsAvailabilityOnDemand = "ON_DEMAND"
	// AwsAvailabilitySpotWithFallback is Spot instance type for clusters with option
	// to fallback into on-demand if instance cannot be acquired
	AwsAvailabilitySpotWithFallback = "SPOT_WITH_FALLBACK"
)

// https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#--azureavailability
const (
	// AzureAvailabilitySpot is spot instance type for clusters
	AzureAvailabilitySpot = "SPOT_AZURE"
	// AzureAvailabilityOnDemand is OnDemand instance type for clusters
	AzureAvailabilityOnDemand = "ON_DEMAND_AZURE"
	// AzureAvailabilitySpotWithFallback is Spot instance type for clusters with option
	// to fallback into on-demand if instance cannot be acquired
	AzureAvailabilitySpotWithFallback = "SPOT_WITH_FALLBACK_AZURE"
)

// https://docs.gcp.databricks.com/dev-tools/api/latest/clusters.html#gcpavailability
const (
	// GcpAvailabilityPreemptible is Preemptible instance type for clusters
	GcpAvailabilityPreemptible = "PREEMPTIBLE_GCP"
	// GcpAvailabilityOnDemand is OnDemand instance type for clusters
	GcpAvailabilityOnDemand = "ON_DEMAND_GCP"
	// GcpAvailabilityPreemptible is Preemptible instance type for clusters with option
	// to fallback into on-demand if instance cannot be acquired
	GcpAvailabilityPreemptibleWithFallback = "PREEMPTIBLE_WITH_FALLBACK_GCP"
)

// AzureDiskVolumeType is disk type on azure vms
type AzureDiskVolumeType string

const (
	// AzureDiskVolumeTypeStandard is for standard local redundant storage
	AzureDiskVolumeTypeStandard = "STANDARD_LRS"
	// AzureDiskVolumeTypePremium is for premium local redundant storage
	AzureDiskVolumeTypePremium = "PREMIUM_LRS"
)

// EbsVolumeType is disk type on aws vms
type EbsVolumeType string

const (
	// EbsVolumeTypeGeneralPurposeSsd is general purpose ssd (starts at 32 gb)
	EbsVolumeTypeGeneralPurposeSsd = "GENERAL_PURPOSE_SSD"
	// EbsVolumeTypeThroughputOptimizedHdd is throughput optimized hdd (starts at 500 gb)
	EbsVolumeTypeThroughputOptimizedHdd = "THROUGHPUT_OPTIMIZED_HDD"
)

// ClusterState is for describing possible cluster states
type ClusterState string

const (
	// ClusterStatePending Indicates that a cluster is in the process of being created.
	ClusterStatePending = "PENDING"
	// ClusterStateRunning Indicates that a cluster has been started and is ready for use.
	ClusterStateRunning = "RUNNING"
	// ClusterStateRestarting Indicates that a cluster is in the process of restarting.
	ClusterStateRestarting = "RESTARTING"
	// ClusterStateResizing Indicates that a cluster is in the process of adding or removing nodes.
	ClusterStateResizing = "RESIZING"
	// ClusterStateTerminating Indicates that a cluster is in the process of being destroyed.
	ClusterStateTerminating = "TERMINATING"
	// ClusterStateTerminated Indicates that a cluster has been successfully destroyed.
	ClusterStateTerminated = "TERMINATED"
	// ClusterStateError This state is not used anymore. It was used to indicate a cluster
	// that failed to be created. Terminating and Terminated are used instead.
	ClusterStateError = "ERROR"
	// ClusterStateUnknown Indicates that a cluster is in an unknown state. A cluster should never be in this state.
	ClusterStateUnknown = "UNKNOWN"
)

var stateMachine = map[ClusterState][]ClusterState{
	ClusterStatePending:     {ClusterStateRunning, ClusterStateTerminating},
	ClusterStateRunning:     {ClusterStateResizing, ClusterStateRestarting, ClusterStateTerminating},
	ClusterStateRestarting:  {ClusterStateRunning, ClusterStateTerminating},
	ClusterStateResizing:    {ClusterStateRunning, ClusterStateTerminating},
	ClusterStateTerminating: {ClusterStateTerminated},
}

// TODO: rewrite later
// CanReach returns true if cluster state can reach desired state
func (state ClusterState) CanReach(desired ClusterState) bool {
	if state == desired {
		return true
	}
	visited := map[ClusterState]bool{}
	queue := []ClusterState{state}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, ok := visited[current]; ok {
			continue
		}
		adjacent, ok := stateMachine[current]
		visited[current] = true
		if !ok {
			return false
		}
		for _, possible := range adjacent {
			if possible == desired {
				return true
			}
			queue = append(queue, possible)
		}
	}
	return false
}

// ZonesInfo encapsulates the zone information from the zones api call
type ZonesInfo struct {
	Zones       []string `json:"zones,omitempty"`
	DefaultZone string   `json:"default_zone,omitempty"`
}

// AwsAttributes encapsulates the aws attributes for aws based clusters
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclusterattributes
type AwsAttributes struct {
	FirstOnDemand       int32         `json:"first_on_demand,omitempty"`
	Availability        Availability  `json:"availability,omitempty"`
	ZoneID              string        `json:"zone_id,omitempty"`
	InstanceProfileArn  string        `json:"instance_profile_arn,omitempty"`
	SpotBidPricePercent int32         `json:"spot_bid_price_percent,omitempty"`
	EbsVolumeType       EbsVolumeType `json:"ebs_volume_type,omitempty"`
	EbsVolumeCount      int32         `json:"ebs_volume_count,omitempty"`
	EbsVolumeSize       int32         `json:"ebs_volume_size,omitempty"`
}

// AzureAttributes encapsulates the Azure attributes for Azure based clusters
// https://docs.microsoft.com/en-us/azure/databricks/dev-tools/api/latest/clusters#clusterazureattributes
type AzureAttributes struct {
	FirstOnDemand   int32        `json:"first_on_demand,omitempty"`
	Availability    Availability `json:"availability,omitempty"`
	SpotBidMaxPrice float64      `json:"spot_bid_max_price,omitempty"`
}

// GcpAttributes encapsultes GCP specific attributes
// https://docs.gcp.databricks.com/dev-tools/api/latest/clusters.html#clustergcpattributes
type GcpAttributes struct {
	UsePreemptibleExecutors bool         `json:"use_preemptible_executors,omitempty"`
	GoogleServiceAccount    string       `json:"google_service_account,omitempty"`
	Availability            Availability `json:"availability,omitempty"`
	BootDiskSize            int32        `json:"boot_disk_size,omitempty"`
	ZoneId                  string       `json:"zone_id,omitempty"`
}

// DbfsStorageInfo contains the destination string for DBFS
type DbfsStorageInfo struct {
	Destination string `json:"destination"`
}

// S3StorageInfo contains the struct for when storing files in S3
type S3StorageInfo struct {
	// TODO: add instance profile validation check + prefix validation
	Destination      string `json:"destination"`
	Region           string `json:"region,omitempty" tf:"group:location"`
	Endpoint         string `json:"endpoint,omitempty" tf:"group:location"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	CannedACL        string `json:"canned_acl,omitempty"`
}

// GcsStorageInfo contains the struct for when storing files in GCS
type GcsStorageInfo struct {
	Destination string `json:"destination,omitempty" tf:"optional"`
}

// LocalFileInfo represents a local file on disk, e.g. in a customer's container.
type LocalFileInfo struct {
	Destination string `json:"destination,omitempty" tf:"optional"`
}

// StorageInfo contains the struct for either DBFS or S3 storage depending on which one is relevant.
type StorageInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty" tf:"group:storage"`
	S3   *S3StorageInfo   `json:"s3,omitempty" tf:"group:storage"`
}

// InitScriptStorageInfo captures the allowed sources of init scripts.
type InitScriptStorageInfo struct {
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty" tf:"group:storage"`
	Gcs  *GcsStorageInfo  `json:"gcs,omitempty" tf:"group:storage"`
	S3   *S3StorageInfo   `json:"s3,omitempty" tf:"group:storage"`
	File *LocalFileInfo   `json:"file,omitempty" tf:"optional"`
}

// SparkNodeAwsAttributes is the struct that determines if the node is a spot instance or not
type SparkNodeAwsAttributes struct {
	IsSpot bool `json:"is_spot,omitempty"`
}

// SparkNode encapsulates all the attributes of a node that is part of a databricks cluster
type SparkNode struct {
	PrivateIP         string                  `json:"private_ip,omitempty"`
	PublicDNS         string                  `json:"public_dns,omitempty"`
	NodeID            string                  `json:"node_id,omitempty"`
	InstanceID        string                  `json:"instance_id,omitempty"`
	StartTimestamp    int64                   `json:"start_timestamp,omitempty"`
	NodeAwsAttributes *SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
	HostPrivateIP     string                  `json:"host_private_ip,omitempty"`
}

// TerminationReason encapsulates the termination code and potential parameters
type TerminationReason struct {
	Code       string            `json:"code,omitempty"`
	Type       string            `json:"type,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// LogSyncStatus encapsulates when the cluster logs were last delivered.
type LogSyncStatus struct {
	LastAttempted int64  `json:"last_attempted,omitempty"`
	LastException string `json:"last_exception,omitempty"`
}

// DockerBasicAuth contains the auth information when fetching containers
type DockerBasicAuth struct {
	Username string `json:"username" tf:"force_new"`
	Password string `json:"password" tf:"force_new"`
}

// DockerImage contains the image url and the auth for DCS
type DockerImage struct {
	URL       string           `json:"url" tf:"force_new"`
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty" tf:"force_new"`
}

// SortOrder - constants for API
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlistorder
type SortOrder string

// constants for SortOrder
const (
	SortDescending SortOrder = "DESC"
	SortAscending  SortOrder = "ASC"
)

// ClusterEventType - constants for API
type ClusterEventType string

// Constants for Event Types
const (
	EvTypeCreating            ClusterEventType = "CREATING"
	EvTypeDidNotExpandDisk    ClusterEventType = "DID_NOT_EXPAND_DISK"
	EvTypeExpandedDisk        ClusterEventType = "EXPANDED_DISK"
	EvTypeFailedToExpandDisk  ClusterEventType = "FAILED_TO_EXPAND_DISK"
	EvTypeInitScriptsStarting ClusterEventType = "INIT_SCRIPTS_STARTING"
	EvTypeInitScriptsFinished ClusterEventType = "INIT_SCRIPTS_FINISHED"
	EvTypeStarting            ClusterEventType = "STARTING"
	EvTypeRestarting          ClusterEventType = "RESTARTING"
	EvTypeTerminating         ClusterEventType = "TERMINATING"
	EvTypeEdited              ClusterEventType = "EDITED"
	EvTypeRunning             ClusterEventType = "RUNNING"
	EvTypeResizing            ClusterEventType = "RESIZING"
	EvTypeUpsizeCompleted     ClusterEventType = "UPSIZE_COMPLETED"
	EvTypeNodesLost           ClusterEventType = "NODES_LOST"
	EvTypeDriverHealthy       ClusterEventType = "DRIVER_HEALTHY"
	EvTypeDriverUnavailable   ClusterEventType = "DRIVER_UNAVAILABLE"
	EvTypeSparkException      ClusterEventType = "SPARK_EXCEPTION"
	EvTypeDriverNotResponding ClusterEventType = "DRIVER_NOT_RESPONDING"
	EvTypeDbfsDown            ClusterEventType = "DBFS_DOWN"
	EvTypeMetastoreDown       ClusterEventType = "METASTORE_DOWN"
	EvTypeNodeBlacklisted     ClusterEventType = "NODE_BLACKLISTED"
	EvTypePinned              ClusterEventType = "PINNED"
	EvTypeUnpinned            ClusterEventType = "UNPINNED"
)

// EventsRequest - request structure
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#request-structure
type EventsRequest struct {
	ClusterID  string             `json:"cluster_id"`
	StartTime  int64              `json:"start_time,omitempty"`
	EndTime    int64              `json:"end_time,omitempty"`
	Order      SortOrder          `json:"order,omitempty"`
	EventTypes []ClusterEventType `json:"event_types,omitempty"`
	Offset     int64              `json:"offset,omitempty"`
	Limit      int64              `json:"limit,omitempty"`
	MaxItems   uint               `json:"-"`
}

// ClusterSize is structure to keep
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterclustersize
type ClusterSize struct {
	NumWorkers int32      `json:"num_workers"`
	AutoScale  *AutoScale `json:"autoscale"`
}

// ResizeCause holds reason for resizing
type ResizeCause string

// EventDetails - details about specific events
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clustereventseventdetails
type EventDetails struct {
	CurrentNumWorkers   int32              `json:"current_num_workers,omitempty"`
	TargetNumWorkers    int32              `json:"target_num_workers,omitempty"`
	PreviousAttributes  *AwsAttributes     `json:"previous_attributes,omitempty"`
	Attributes          *AwsAttributes     `json:"attributes,omitempty"`
	PreviousClusterSize *ClusterSize       `json:"previous_cluster_size,omitempty"`
	ClusterSize         *ClusterSize       `json:"cluster_size,omitempty"`
	ResizeCause         *ResizeCause       `json:"cause,omitempty"`
	Reason              *TerminationReason `json:"reason,omitempty"`
	User                string             `json:"user"`
}

// ClusterEvent - event information
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#clustereventsclusterevent
type ClusterEvent struct {
	ClusterID string           `json:"cluster_id"`
	Timestamp int64            `json:"timestamp"`
	Type      ClusterEventType `json:"type"`
	Details   EventDetails     `json:"details"`
}

// EventsResponse - answer from API
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#response-structure
type EventsResponse struct {
	Events     []ClusterEvent `json:"events"`
	NextPage   *EventsRequest `json:"next_page"`
	TotalCount int64          `json:"total_count"`
}

// Cluster contains the information when trying to submit api calls or editing a cluster
type Cluster struct {
	ClusterID   string `json:"cluster_id,omitempty"`
	ClusterName string `json:"cluster_name,omitempty"`

	SparkVersion              string     `json:"spark_version"`
	NumWorkers                int32      `json:"num_workers" tf:"group:size"`
	Autoscale                 *AutoScale `json:"autoscale,omitempty" tf:"group:size"`
	EnableElasticDisk         bool       `json:"enable_elastic_disk,omitempty" tf:"computed"`
	EnableLocalDiskEncryption bool       `json:"enable_local_disk_encryption,omitempty" tf:"computed"`

	NodeTypeID             string           `json:"node_type_id,omitempty" tf:"group:node_type,computed"`
	DriverNodeTypeID       string           `json:"driver_node_type_id,omitempty" tf:"group:node_type,computed"`
	InstancePoolID         string           `json:"instance_pool_id,omitempty" tf:"group:node_type"`
	DriverInstancePoolID   string           `json:"driver_instance_pool_id,omitempty" tf:"group:node_type,computed"`
	PolicyID               string           `json:"policy_id,omitempty"`
	AwsAttributes          *AwsAttributes   `json:"aws_attributes,omitempty" tf:"conflicts:instance_pool_id,suppress_diff"`
	AzureAttributes        *AzureAttributes `json:"azure_attributes,omitempty" tf:"conflicts:instance_pool_id,suppress_diff"`
	GcpAttributes          *GcpAttributes   `json:"gcp_attributes,omitempty" tf:"conflicts:instance_pool_id,suppress_diff"`
	AutoterminationMinutes int32            `json:"autotermination_minutes,omitempty"`

	SparkConf    map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	CustomTags   map[string]string `json:"custom_tags,omitempty"`

	SSHPublicKeys  []string                `json:"ssh_public_keys,omitempty" tf:"max_items:10"`
	InitScripts    []InitScriptStorageInfo `json:"init_scripts,omitempty" tf:"max_items:10"`
	ClusterLogConf *StorageInfo            `json:"cluster_log_conf,omitempty"`
	DockerImage    *DockerImage            `json:"docker_image,omitempty"`

	DataSecurityMode string `json:"data_security_mode,omitempty"`
	SingleUserName   string `json:"single_user_name,omitempty"`
	IdempotencyToken string `json:"idempotency_token,omitempty" tf:"force_new"`
}

func (cluster Cluster) Validate() error {
	// TODO: rewrite with CustomizeDiff
	if cluster.NumWorkers > 0 || cluster.Autoscale != nil {
		return nil
	}
	profile := cluster.SparkConf["spark.databricks.cluster.profile"]
	master := cluster.SparkConf["spark.master"]
	resourceClass := cluster.CustomTags["ResourceClass"]
	if profile == "singleNode" && strings.HasPrefix(master, "local") && resourceClass == "SingleNode" {
		return nil
	}
	return fmt.Errorf("NumWorkers could be 0 only for SingleNode clusters. See https://docs.databricks.com/clusters/single-node.html for more details")
}

// ModifyRequestOnInstancePool helps remove all request fields that should not be submitted when instance pool is selected.
func (cluster *Cluster) ModifyRequestOnInstancePool() {
	// Instance profile id does not exist or not set
	if cluster.InstancePoolID == "" {
		return
	}
	if cluster.AwsAttributes != nil {
		// Reset AwsAttributes
		awsAttributes := AwsAttributes{
			InstanceProfileArn: cluster.AwsAttributes.InstanceProfileArn,
		}
		cluster.AwsAttributes = &awsAttributes
	}
	if cluster.AzureAttributes != nil {
		cluster.AzureAttributes = nil
	}
	if cluster.GcpAttributes != nil {
		gcpAttributes := GcpAttributes{
			GoogleServiceAccount: cluster.GcpAttributes.GoogleServiceAccount,
		}
		cluster.GcpAttributes = &gcpAttributes
	}
	cluster.EnableElasticDisk = false
	cluster.NodeTypeID = ""
	cluster.DriverNodeTypeID = ""
}

// ClusterList shows existing clusters
type ClusterList struct {
	Clusters []ClusterInfo `json:"clusters,omitempty"`
}

// ClusterInfo contains the information when getting cluster info from the get request.
type ClusterInfo struct {
	NumWorkers                int32              `json:"num_workers,omitempty"`
	AutoScale                 *AutoScale         `json:"autoscale,omitempty"`
	ClusterID                 string             `json:"cluster_id,omitempty"`
	CreatorUserName           string             `json:"creator_user_name,omitempty"`
	Driver                    *SparkNode         `json:"driver,omitempty"`
	Executors                 []SparkNode        `json:"executors,omitempty"`
	SparkContextID            int64              `json:"spark_context_id,omitempty"`
	JdbcPort                  int32              `json:"jdbc_port,omitempty"`
	ClusterName               string             `json:"cluster_name,omitempty"`
	SparkVersion              string             `json:"spark_version"`
	SparkConf                 map[string]string  `json:"spark_conf,omitempty"`
	AwsAttributes             *AwsAttributes     `json:"aws_attributes,omitempty"`
	AzureAttributes           *AzureAttributes   `json:"azure_attributes,omitempty"`
	GcpAttributes             *GcpAttributes     `json:"gcp_attributes,omitempty"`
	NodeTypeID                string             `json:"node_type_id,omitempty"`
	DriverNodeTypeID          string             `json:"driver_node_type_id,omitempty"`
	SSHPublicKeys             []string           `json:"ssh_public_keys,omitempty"`
	CustomTags                map[string]string  `json:"custom_tags,omitempty"`
	ClusterLogConf            *StorageInfo       `json:"cluster_log_conf,omitempty"`
	InitScripts               []StorageInfo      `json:"init_scripts,omitempty"`
	SparkEnvVars              map[string]string  `json:"spark_env_vars,omitempty"`
	AutoterminationMinutes    int32              `json:"autotermination_minutes,omitempty"`
	EnableElasticDisk         bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption bool               `json:"enable_local_disk_encryption,omitempty"`
	InstancePoolID            string             `json:"instance_pool_id,omitempty"`
	DriverInstancePoolID      string             `json:"driver_instance_pool_id,omitempty" tf:"computed"`
	PolicyID                  string             `json:"policy_id,omitempty"`
	SingleUserName            string             `json:"single_user_name,omitempty"`
	ClusterSource             Availability       `json:"cluster_source,omitempty"`
	DockerImage               *DockerImage       `json:"docker_image,omitempty"`
	State                     ClusterState       `json:"state"`
	StateMessage              string             `json:"state_message,omitempty"`
	StartTime                 int64              `json:"start_time,omitempty"`
	TerminateTime             int64              `json:"terminate_time,omitempty"`
	LastStateLossTime         int64              `json:"last_state_loss_time,omitempty"`
	LastActivityTime          int64              `json:"last_activity_time,omitempty"`
	ClusterMemoryMb           int64              `json:"cluster_memory_mb,omitempty"`
	ClusterCores              float64            `json:"cluster_cores,omitempty"`
	DefaultTags               map[string]string  `json:"default_tags"`
	ClusterLogStatus          *LogSyncStatus     `json:"cluster_log_status,omitempty"`
	TerminationReason         *TerminationReason `json:"termination_reason,omitempty"`
	DataSecurityMode          string             `json:"data_security_mode,omitempty"`
}

// IsRunningOrResizing returns true if cluster is running or resizing
func (ci *ClusterInfo) IsRunningOrResizing() bool {
	return ci.State == ClusterStateRunning || ci.State == ClusterStateResizing
}

// ClusterID holds cluster ID
type ClusterID struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}

// NodeTypeList contains a list of node types
type NodeTypeList struct {
	NodeTypes []NodeType `json:"node_types,omitempty"`
}

// NodeType encapsulates information about a given node when using the list-node-types api
type NodeType struct {
	NodeTypeID            string                        `json:"node_type_id,omitempty"`
	MemoryMB              int32                         `json:"memory_mb,omitempty"`
	NumCores              float32                       `json:"num_cores,omitempty"`
	NumGPUs               int32                         `json:"num_gpus,omitempty"`
	SupportEBSVolumes     bool                          `json:"support_ebs_volumes,omitempty"`
	IsIOCacheEnabled      bool                          `json:"is_io_cache_enabled,omitempty"`
	SupportPortForwarding bool                          `json:"support_port_forwarding,omitempty"`
	Description           string                        `json:"description,omitempty"`
	Category              string                        `json:"category,omitempty"`
	InstanceTypeID        string                        `json:"instance_type_id,omitempty"`
	IsDeprecated          bool                          `json:"is_deprecated,omitempty"`
	IsHidden              bool                          `json:"is_hidden,omitempty"`
	SupportClusterTags    bool                          `json:"support_cluster_tags,omitempty"`
	DisplayOrder          int32                         `json:"display_order,omitempty"`
	NodeInfo              *ClusterCloudProviderNodeInfo `json:"node_info,omitempty"`
	NodeInstanceType      *NodeInstanceType             `json:"node_instance_type,omitempty"`
	PhotonWorkerCapable   bool                          `json:"photon_worker_capable,omitempty"`
	PhotonDriverCapable   bool                          `json:"photon_driver_capable,omitempty"`
	Graviton              bool                          `json:"is_graviton,omitempty"`
}

// ClusterCloudProviderNodeInfo encapsulates the existing quota available from the cloud service provider.
type ClusterCloudProviderNodeInfo struct {
	Status             []string `json:"status,omitempty"`
	AvailableCoreQuota float32  `json:"available_core_quota,omitempty"`
	TotalCoreQuota     float32  `json:"total_core_quota,omitempty"`
}

// NodeInstanceType encapsulates information about a specific node type
type NodeInstanceType struct {
	InstanceTypeID      string `json:"instance_type_id,omitempty"`
	LocalDisks          int32  `json:"local_disks,omitempty"`
	LocalDiskSizeGB     int32  `json:"local_disk_size_gb,omitempty"`
	LocalNVMeDisks      int32  `json:"local_nvme_disks,omitempty"`
	LocalNVMeDiskSizeGB int32  `json:"local_nvme_disk_size_gb,omitempty"`
}

// SparkVersion - contains information about specific version
type SparkVersion struct {
	Version     string `json:"key"`
	Description string `json:"name"`
}

// SparkVersionsList - returns a list of all currently supported Spark Versions
// https://docs.databricks.com/dev-tools/api/latest/clusters.html#runtime-versions
type SparkVersionsList struct {
	SparkVersions []SparkVersion `json:"versions"`
}