// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package computepb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddInstanceProfilePb struct {
	IamRoleArn            string `json:"iam_role_arn,omitempty"`
	InstanceProfileArn    string `json:"instance_profile_arn"`
	IsMetaInstanceProfile bool   `json:"is_meta_instance_profile,omitempty"`
	SkipValidation        bool   `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AddInstanceProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AddInstanceProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AddResponsePb struct {
}

type Adlsgen2InfoPb struct {
	Destination string `json:"destination"`
}

type AutoScalePb struct {
	MaxWorkers int `json:"max_workers,omitempty"`
	MinWorkers int `json:"min_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AutoScalePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AutoScalePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AwsAttributesPb struct {
	Availability        AwsAvailabilityPb `json:"availability,omitempty"`
	EbsVolumeCount      int               `json:"ebs_volume_count,omitempty"`
	EbsVolumeIops       int               `json:"ebs_volume_iops,omitempty"`
	EbsVolumeSize       int               `json:"ebs_volume_size,omitempty"`
	EbsVolumeThroughput int               `json:"ebs_volume_throughput,omitempty"`
	EbsVolumeType       EbsVolumeTypePb   `json:"ebs_volume_type,omitempty"`
	FirstOnDemand       int               `json:"first_on_demand,omitempty"`
	InstanceProfileArn  string            `json:"instance_profile_arn,omitempty"`
	SpotBidPricePercent int               `json:"spot_bid_price_percent,omitempty"`
	ZoneId              string            `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AwsAvailabilityPb string

const AwsAvailabilityOnDemand AwsAvailabilityPb = `ON_DEMAND`

const AwsAvailabilitySpot AwsAvailabilityPb = `SPOT`

const AwsAvailabilitySpotWithFallback AwsAvailabilityPb = `SPOT_WITH_FALLBACK`

type AzureAttributesPb struct {
	Availability     AzureAvailabilityPb `json:"availability,omitempty"`
	FirstOnDemand    int                 `json:"first_on_demand,omitempty"`
	LogAnalyticsInfo *LogAnalyticsInfoPb `json:"log_analytics_info,omitempty"`
	SpotBidMaxPrice  float64             `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AzureAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AzureAvailabilityPb string

const AzureAvailabilityOnDemandAzure AzureAvailabilityPb = `ON_DEMAND_AZURE`

const AzureAvailabilitySpotAzure AzureAvailabilityPb = `SPOT_AZURE`

const AzureAvailabilitySpotWithFallbackAzure AzureAvailabilityPb = `SPOT_WITH_FALLBACK_AZURE`

type CancelCommandPb struct {
	ClusterId string `json:"clusterId,omitempty"`
	CommandId string `json:"commandId,omitempty"`
	ContextId string `json:"contextId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CancelCommandPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CancelCommandPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelResponsePb struct {
}

type ChangeClusterOwnerPb struct {
	ClusterId     string `json:"cluster_id"`
	OwnerUsername string `json:"owner_username"`
}

type ChangeClusterOwnerResponsePb struct {
}

type ClientsTypesPb struct {
	Jobs      bool `json:"jobs,omitempty"`
	Notebooks bool `json:"notebooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClientsTypesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClientsTypesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CloneClusterPb struct {
	SourceClusterId string `json:"source_cluster_id"`
}

type CloudProviderNodeInfoPb struct {
	Status []CloudProviderNodeStatusPb `json:"status,omitempty"`
}

type CloudProviderNodeStatusPb string

const CloudProviderNodeStatusNotAvailableInRegion CloudProviderNodeStatusPb = `NotAvailableInRegion`

const CloudProviderNodeStatusNotEnabledOnSubscription CloudProviderNodeStatusPb = `NotEnabledOnSubscription`

type ClusterAccessControlRequestPb struct {
	GroupName            string                   `json:"group_name,omitempty"`
	PermissionLevel      ClusterPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                   `json:"service_principal_name,omitempty"`
	UserName             string                   `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAccessControlResponsePb struct {
	AllPermissions       []ClusterPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                `json:"display_name,omitempty"`
	GroupName            string                `json:"group_name,omitempty"`
	ServicePrincipalName string                `json:"service_principal_name,omitempty"`
	UserName             string                `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAttributesPb struct {
	AutoterminationMinutes     int                `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributesPb `json:"azure_attributes,omitempty"`
	ClusterLogConf             *ClusterLogConfPb  `json:"cluster_log_conf,omitempty"`
	ClusterName                string             `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string  `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityModePb `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImagePb     `json:"docker_image,omitempty"`
	DriverInstancePoolId       string             `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string             `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool               `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributesPb   `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfoPb `json:"init_scripts,omitempty"`
	InstancePoolId             string             `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool               `json:"is_single_node,omitempty"`
	Kind                       KindPb             `json:"kind,omitempty"`
	NodeTypeId                 string             `json:"node_type_id,omitempty"`
	PolicyId                   string             `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEnginePb    `json:"runtime_engine,omitempty"`
	SingleUserName             string             `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string  `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string  `json:"spark_env_vars,omitempty"`
	SparkVersion               string             `json:"spark_version"`
	SshPublicKeys              []string           `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int                `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool               `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadTypePb    `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterCompliancePb struct {
	ClusterId   string            `json:"cluster_id"`
	IsCompliant bool              `json:"is_compliant,omitempty"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterCompliancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterCompliancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterDetailsPb struct {
	Autoscale                  *AutoScalePb         `json:"autoscale,omitempty"`
	AutoterminationMinutes     int                  `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributesPb     `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributesPb   `json:"azure_attributes,omitempty"`
	ClusterCores               float64              `json:"cluster_cores,omitempty"`
	ClusterId                  string               `json:"cluster_id,omitempty"`
	ClusterLogConf             *ClusterLogConfPb    `json:"cluster_log_conf,omitempty"`
	ClusterLogStatus           *LogSyncStatusPb     `json:"cluster_log_status,omitempty"`
	ClusterMemoryMb            int64                `json:"cluster_memory_mb,omitempty"`
	ClusterName                string               `json:"cluster_name,omitempty"`
	ClusterSource              ClusterSourcePb      `json:"cluster_source,omitempty"`
	CreatorUserName            string               `json:"creator_user_name,omitempty"`
	CustomTags                 map[string]string    `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityModePb   `json:"data_security_mode,omitempty"`
	DefaultTags                map[string]string    `json:"default_tags,omitempty"`
	DockerImage                *DockerImagePb       `json:"docker_image,omitempty"`
	Driver                     *SparkNodePb         `json:"driver,omitempty"`
	DriverInstancePoolId       string               `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string               `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool                 `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool                 `json:"enable_local_disk_encryption,omitempty"`
	Executors                  []SparkNodePb        `json:"executors,omitempty"`
	GcpAttributes              *GcpAttributesPb     `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfoPb   `json:"init_scripts,omitempty"`
	InstancePoolId             string               `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool                 `json:"is_single_node,omitempty"`
	JdbcPort                   int                  `json:"jdbc_port,omitempty"`
	Kind                       KindPb               `json:"kind,omitempty"`
	LastRestartedTime          int64                `json:"last_restarted_time,omitempty"`
	LastStateLossTime          int64                `json:"last_state_loss_time,omitempty"`
	NodeTypeId                 string               `json:"node_type_id,omitempty"`
	NumWorkers                 int                  `json:"num_workers,omitempty"`
	PolicyId                   string               `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                  `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEnginePb      `json:"runtime_engine,omitempty"`
	SingleUserName             string               `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string    `json:"spark_conf,omitempty"`
	SparkContextId             int64                `json:"spark_context_id,omitempty"`
	SparkEnvVars               map[string]string    `json:"spark_env_vars,omitempty"`
	SparkVersion               string               `json:"spark_version,omitempty"`
	Spec                       *ClusterSpecPb       `json:"spec,omitempty"`
	SshPublicKeys              []string             `json:"ssh_public_keys,omitempty"`
	StartTime                  int64                `json:"start_time,omitempty"`
	State                      StatePb              `json:"state,omitempty"`
	StateMessage               string               `json:"state_message,omitempty"`
	TerminatedTime             int64                `json:"terminated_time,omitempty"`
	TerminationReason          *TerminationReasonPb `json:"termination_reason,omitempty"`
	TotalInitialRemoteDiskSize int                  `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool                 `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadTypePb      `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterEventPb struct {
	ClusterId             string                   `json:"cluster_id"`
	DataPlaneEventDetails *DataPlaneEventDetailsPb `json:"data_plane_event_details,omitempty"`
	Details               *EventDetailsPb          `json:"details,omitempty"`
	Timestamp             int64                    `json:"timestamp,omitempty"`
	Type                  EventTypePb              `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterEventPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterEventPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterLibraryStatusesPb struct {
	ClusterId       string                `json:"cluster_id,omitempty"`
	LibraryStatuses []LibraryFullStatusPb `json:"library_statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterLibraryStatusesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterLibraryStatusesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterLogConfPb struct {
	Dbfs    *DbfsStorageInfoPb    `json:"dbfs,omitempty"`
	S3      *S3StorageInfoPb      `json:"s3,omitempty"`
	Volumes *VolumesStorageInfoPb `json:"volumes,omitempty"`
}

type ClusterPermissionPb struct {
	Inherited           bool                     `json:"inherited,omitempty"`
	InheritedFromObject []string                 `json:"inherited_from_object,omitempty"`
	PermissionLevel     ClusterPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPermissionLevelPb string

const ClusterPermissionLevelCanAttachTo ClusterPermissionLevelPb = `CAN_ATTACH_TO`

const ClusterPermissionLevelCanManage ClusterPermissionLevelPb = `CAN_MANAGE`

const ClusterPermissionLevelCanRestart ClusterPermissionLevelPb = `CAN_RESTART`

type ClusterPermissionsPb struct {
	AccessControlList []ClusterAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                           `json:"object_id,omitempty"`
	ObjectType        string                           `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPermissionsDescriptionPb struct {
	Description     string                   `json:"description,omitempty"`
	PermissionLevel ClusterPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPermissionsRequestPb struct {
	AccessControlList []ClusterAccessControlRequestPb `json:"access_control_list,omitempty"`
	ClusterId         string                          `json:"-" url:"-"`
}

type ClusterPolicyAccessControlRequestPb struct {
	GroupName            string                         `json:"group_name,omitempty"`
	PermissionLevel      ClusterPolicyPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                         `json:"service_principal_name,omitempty"`
	UserName             string                         `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPolicyAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPolicyAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyAccessControlResponsePb struct {
	AllPermissions       []ClusterPolicyPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                      `json:"display_name,omitempty"`
	GroupName            string                      `json:"group_name,omitempty"`
	ServicePrincipalName string                      `json:"service_principal_name,omitempty"`
	UserName             string                      `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPolicyAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPolicyAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermissionPb struct {
	Inherited           bool                           `json:"inherited,omitempty"`
	InheritedFromObject []string                       `json:"inherited_from_object,omitempty"`
	PermissionLevel     ClusterPolicyPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPolicyPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPolicyPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermissionLevelPb string

const ClusterPolicyPermissionLevelCanUse ClusterPolicyPermissionLevelPb = `CAN_USE`

type ClusterPolicyPermissionsPb struct {
	AccessControlList []ClusterPolicyAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                                 `json:"object_id,omitempty"`
	ObjectType        string                                 `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPolicyPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPolicyPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermissionsDescriptionPb struct {
	Description     string                         `json:"description,omitempty"`
	PermissionLevel ClusterPolicyPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterPolicyPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterPolicyPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermissionsRequestPb struct {
	AccessControlList []ClusterPolicyAccessControlRequestPb `json:"access_control_list,omitempty"`
	ClusterPolicyId   string                                `json:"-" url:"-"`
}

type ClusterSettingsChangePb struct {
	Field         string `json:"field,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterSettingsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterSettingsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterSizePb struct {
	Autoscale  *AutoScalePb `json:"autoscale,omitempty"`
	NumWorkers int          `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterSizePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterSizePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterSourcePb string

const ClusterSourceApi ClusterSourcePb = `API`

const ClusterSourceJob ClusterSourcePb = `JOB`

const ClusterSourceModels ClusterSourcePb = `MODELS`

const ClusterSourcePipeline ClusterSourcePb = `PIPELINE`

const ClusterSourcePipelineMaintenance ClusterSourcePb = `PIPELINE_MAINTENANCE`

const ClusterSourceSql ClusterSourcePb = `SQL`

const ClusterSourceUi ClusterSourcePb = `UI`

type ClusterSpecPb struct {
	ApplyPolicyDefaultValues   bool               `json:"apply_policy_default_values,omitempty"`
	Autoscale                  *AutoScalePb       `json:"autoscale,omitempty"`
	AutoterminationMinutes     int                `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributesPb `json:"azure_attributes,omitempty"`
	ClusterLogConf             *ClusterLogConfPb  `json:"cluster_log_conf,omitempty"`
	ClusterName                string             `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string  `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityModePb `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImagePb     `json:"docker_image,omitempty"`
	DriverInstancePoolId       string             `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string             `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool               `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributesPb   `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfoPb `json:"init_scripts,omitempty"`
	InstancePoolId             string             `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool               `json:"is_single_node,omitempty"`
	Kind                       KindPb             `json:"kind,omitempty"`
	NodeTypeId                 string             `json:"node_type_id,omitempty"`
	NumWorkers                 int                `json:"num_workers,omitempty"`
	PolicyId                   string             `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEnginePb    `json:"runtime_engine,omitempty"`
	SingleUserName             string             `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string  `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string  `json:"spark_env_vars,omitempty"`
	SparkVersion               string             `json:"spark_version,omitempty"`
	SshPublicKeys              []string           `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int                `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool               `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadTypePb    `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterStatusPb struct {
	ClusterId string `json:"-" url:"cluster_id"`
}

type CommandPb struct {
	ClusterId string     `json:"clusterId,omitempty"`
	Command   string     `json:"command,omitempty"`
	ContextId string     `json:"contextId,omitempty"`
	Language  LanguagePb `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CommandPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CommandPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CommandStatusPb string

const CommandStatusCancelled CommandStatusPb = `Cancelled`

const CommandStatusCancelling CommandStatusPb = `Cancelling`

const CommandStatusError CommandStatusPb = `Error`

const CommandStatusFinished CommandStatusPb = `Finished`

const CommandStatusQueued CommandStatusPb = `Queued`

const CommandStatusRunning CommandStatusPb = `Running`

type CommandStatusRequestPb struct {
	ClusterId string `json:"-" url:"clusterId"`
	CommandId string `json:"-" url:"commandId"`
	ContextId string `json:"-" url:"contextId"`
}

type CommandStatusResponsePb struct {
	Id      string          `json:"id,omitempty"`
	Results *ResultsPb      `json:"results,omitempty"`
	Status  CommandStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CommandStatusResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CommandStatusResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ContextStatusPb string

const ContextStatusError ContextStatusPb = `Error`

const ContextStatusPending ContextStatusPb = `Pending`

const ContextStatusRunning ContextStatusPb = `Running`

type ContextStatusRequestPb struct {
	ClusterId string `json:"-" url:"clusterId"`
	ContextId string `json:"-" url:"contextId"`
}

type ContextStatusResponsePb struct {
	Id     string          `json:"id,omitempty"`
	Status ContextStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ContextStatusResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ContextStatusResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateClusterPb struct {
	ApplyPolicyDefaultValues   bool               `json:"apply_policy_default_values,omitempty"`
	Autoscale                  *AutoScalePb       `json:"autoscale,omitempty"`
	AutoterminationMinutes     int                `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributesPb `json:"azure_attributes,omitempty"`
	CloneFrom                  *CloneClusterPb    `json:"clone_from,omitempty"`
	ClusterLogConf             *ClusterLogConfPb  `json:"cluster_log_conf,omitempty"`
	ClusterName                string             `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string  `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityModePb `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImagePb     `json:"docker_image,omitempty"`
	DriverInstancePoolId       string             `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string             `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool               `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributesPb   `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfoPb `json:"init_scripts,omitempty"`
	InstancePoolId             string             `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool               `json:"is_single_node,omitempty"`
	Kind                       KindPb             `json:"kind,omitempty"`
	NodeTypeId                 string             `json:"node_type_id,omitempty"`
	NumWorkers                 int                `json:"num_workers,omitempty"`
	PolicyId                   string             `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEnginePb    `json:"runtime_engine,omitempty"`
	SingleUserName             string             `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string  `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string  `json:"spark_env_vars,omitempty"`
	SparkVersion               string             `json:"spark_version"`
	SshPublicKeys              []string           `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int                `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool               `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadTypePb    `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateClusterResponsePb struct {
	ClusterId string `json:"cluster_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateClusterResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateClusterResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateContextPb struct {
	ClusterId string     `json:"clusterId,omitempty"`
	Language  LanguagePb `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateContextPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateContextPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateInstancePoolPb struct {
	AwsAttributes                      *InstancePoolAwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributesPb `json:"azure_attributes,omitempty"`
	CustomTags                         map[string]string              `json:"custom_tags,omitempty"`
	DiskSpec                           *DiskSpecPb                    `json:"disk_spec,omitempty"`
	EnableElasticDisk                  bool                           `json:"enable_elastic_disk,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributesPb   `json:"gcp_attributes,omitempty"`
	IdleInstanceAutoterminationMinutes int                            `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolName                   string                         `json:"instance_pool_name"`
	MaxCapacity                        int                            `json:"max_capacity,omitempty"`
	MinIdleInstances                   int                            `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string                         `json:"node_type_id"`
	PreloadedDockerImages              []DockerImagePb                `json:"preloaded_docker_images,omitempty"`
	PreloadedSparkVersions             []string                       `json:"preloaded_spark_versions,omitempty"`
	RemoteDiskThroughput               int                            `json:"remote_disk_throughput,omitempty"`
	TotalInitialRemoteDiskSize         int                            `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateInstancePoolResponsePb struct {
	InstancePoolId string `json:"instance_pool_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateInstancePoolResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateInstancePoolResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePolicyPb struct {
	Definition                      string      `json:"definition,omitempty"`
	Description                     string      `json:"description,omitempty"`
	Libraries                       []LibraryPb `json:"libraries,omitempty"`
	MaxClustersPerUser              int64       `json:"max_clusters_per_user,omitempty"`
	Name                            string      `json:"name,omitempty"`
	PolicyFamilyDefinitionOverrides string      `json:"policy_family_definition_overrides,omitempty"`
	PolicyFamilyId                  string      `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePolicyResponsePb struct {
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateResponsePb struct {
	ScriptId string `json:"script_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatedPb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatedPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatedPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomPolicyTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CustomPolicyTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomPolicyTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataPlaneEventDetailsPb struct {
	EventType        DataPlaneEventDetailsEventTypePb `json:"event_type,omitempty"`
	ExecutorFailures int                              `json:"executor_failures,omitempty"`
	HostId           string                           `json:"host_id,omitempty"`
	Timestamp        int64                            `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DataPlaneEventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DataPlaneEventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataPlaneEventDetailsEventTypePb string

const DataPlaneEventDetailsEventTypeNodeBlacklisted DataPlaneEventDetailsEventTypePb = `NODE_BLACKLISTED`

const DataPlaneEventDetailsEventTypeNodeExcludedDecommissioned DataPlaneEventDetailsEventTypePb = `NODE_EXCLUDED_DECOMMISSIONED`

type DataSecurityModePb string

// <Databricks> will choose the most appropriate access mode depending on your
// compute configuration.
const DataSecurityModeDataSecurityModeAuto DataSecurityModePb = `DATA_SECURITY_MODE_AUTO`

// Alias for `SINGLE_USER`.
const DataSecurityModeDataSecurityModeDedicated DataSecurityModePb = `DATA_SECURITY_MODE_DEDICATED`

// Alias for `USER_ISOLATION`.
const DataSecurityModeDataSecurityModeStandard DataSecurityModePb = `DATA_SECURITY_MODE_STANDARD`

// This mode is for users migrating from legacy Passthrough on high concurrency
// clusters.
const DataSecurityModeLegacyPassthrough DataSecurityModePb = `LEGACY_PASSTHROUGH`

// This mode is for users migrating from legacy Passthrough on standard
// clusters.
const DataSecurityModeLegacySingleUser DataSecurityModePb = `LEGACY_SINGLE_USER`

// This mode provides a way that doesn’t have UC nor passthrough enabled.
const DataSecurityModeLegacySingleUserStandard DataSecurityModePb = `LEGACY_SINGLE_USER_STANDARD`

// This mode is for users migrating from legacy Table ACL clusters.
const DataSecurityModeLegacyTableAcl DataSecurityModePb = `LEGACY_TABLE_ACL`

// No security isolation for multiple users sharing the cluster. Data governance
// features are not available in this mode.
const DataSecurityModeNone DataSecurityModePb = `NONE`

// A secure cluster that can only be exclusively used by a single user specified
// in `single_user_name`. Most programming languages, cluster features and data
// governance features are available in this mode.
const DataSecurityModeSingleUser DataSecurityModePb = `SINGLE_USER`

// A secure cluster that can be shared by multiple users. Cluster users are
// fully isolated so that they cannot see each other's data and credentials.
// Most data governance features are supported in this mode. But programming
// languages and cluster features might be limited.
const DataSecurityModeUserIsolation DataSecurityModePb = `USER_ISOLATION`

type DbfsStorageInfoPb struct {
	Destination string `json:"destination"`
}

type DeleteClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

type DeleteClusterResponsePb struct {
}

type DeleteGlobalInitScriptRequestPb struct {
	ScriptId string `json:"-" url:"-"`
}

type DeleteInstancePoolPb struct {
	InstancePoolId string `json:"instance_pool_id"`
}

type DeleteInstancePoolResponsePb struct {
}

type DeletePolicyPb struct {
	PolicyId string `json:"policy_id"`
}

type DeletePolicyResponsePb struct {
}

type DeleteResponsePb struct {
}

type DestroyContextPb struct {
	ClusterId string `json:"clusterId"`
	ContextId string `json:"contextId"`
}

type DestroyResponsePb struct {
}

type DiskSpecPb struct {
	DiskCount      int         `json:"disk_count,omitempty"`
	DiskIops       int         `json:"disk_iops,omitempty"`
	DiskSize       int         `json:"disk_size,omitempty"`
	DiskThroughput int         `json:"disk_throughput,omitempty"`
	DiskType       *DiskTypePb `json:"disk_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DiskSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DiskSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DiskTypePb struct {
	AzureDiskVolumeType DiskTypeAzureDiskVolumeTypePb `json:"azure_disk_volume_type,omitempty"`
	EbsVolumeType       DiskTypeEbsVolumeTypePb       `json:"ebs_volume_type,omitempty"`
}

type DiskTypeAzureDiskVolumeTypePb string

const DiskTypeAzureDiskVolumeTypePremiumLrs DiskTypeAzureDiskVolumeTypePb = `PREMIUM_LRS`

const DiskTypeAzureDiskVolumeTypeStandardLrs DiskTypeAzureDiskVolumeTypePb = `STANDARD_LRS`

type DiskTypeEbsVolumeTypePb string

const DiskTypeEbsVolumeTypeGeneralPurposeSsd DiskTypeEbsVolumeTypePb = `GENERAL_PURPOSE_SSD`

const DiskTypeEbsVolumeTypeThroughputOptimizedHdd DiskTypeEbsVolumeTypePb = `THROUGHPUT_OPTIMIZED_HDD`

type DockerBasicAuthPb struct {
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DockerBasicAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DockerBasicAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DockerImagePb struct {
	BasicAuth *DockerBasicAuthPb `json:"basic_auth,omitempty"`
	Url       string             `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DockerImagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DockerImagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EbsVolumeTypePb string

const EbsVolumeTypeGeneralPurposeSsd EbsVolumeTypePb = `GENERAL_PURPOSE_SSD`

const EbsVolumeTypeThroughputOptimizedHdd EbsVolumeTypePb = `THROUGHPUT_OPTIMIZED_HDD`

type EditClusterPb struct {
	ApplyPolicyDefaultValues   bool               `json:"apply_policy_default_values,omitempty"`
	Autoscale                  *AutoScalePb       `json:"autoscale,omitempty"`
	AutoterminationMinutes     int                `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributesPb `json:"azure_attributes,omitempty"`
	ClusterId                  string             `json:"cluster_id"`
	ClusterLogConf             *ClusterLogConfPb  `json:"cluster_log_conf,omitempty"`
	ClusterName                string             `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string  `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityModePb `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImagePb     `json:"docker_image,omitempty"`
	DriverInstancePoolId       string             `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string             `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool               `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributesPb   `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfoPb `json:"init_scripts,omitempty"`
	InstancePoolId             string             `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool               `json:"is_single_node,omitempty"`
	Kind                       KindPb             `json:"kind,omitempty"`
	NodeTypeId                 string             `json:"node_type_id,omitempty"`
	NumWorkers                 int                `json:"num_workers,omitempty"`
	PolicyId                   string             `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEnginePb    `json:"runtime_engine,omitempty"`
	SingleUserName             string             `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string  `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string  `json:"spark_env_vars,omitempty"`
	SparkVersion               string             `json:"spark_version"`
	SshPublicKeys              []string           `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int                `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool               `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadTypePb    `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EditClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EditClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditClusterResponsePb struct {
}

type EditInstancePoolPb struct {
	CustomTags                         map[string]string `json:"custom_tags,omitempty"`
	IdleInstanceAutoterminationMinutes int               `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolId                     string            `json:"instance_pool_id"`
	InstancePoolName                   string            `json:"instance_pool_name"`
	MaxCapacity                        int               `json:"max_capacity,omitempty"`
	MinIdleInstances                   int               `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string            `json:"node_type_id"`
	RemoteDiskThroughput               int               `json:"remote_disk_throughput,omitempty"`
	TotalInitialRemoteDiskSize         int               `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EditInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EditInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditInstancePoolResponsePb struct {
}

type EditPolicyPb struct {
	Definition                      string      `json:"definition,omitempty"`
	Description                     string      `json:"description,omitempty"`
	Libraries                       []LibraryPb `json:"libraries,omitempty"`
	MaxClustersPerUser              int64       `json:"max_clusters_per_user,omitempty"`
	Name                            string      `json:"name,omitempty"`
	PolicyFamilyDefinitionOverrides string      `json:"policy_family_definition_overrides,omitempty"`
	PolicyFamilyId                  string      `json:"policy_family_id,omitempty"`
	PolicyId                        string      `json:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EditPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EditPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditPolicyResponsePb struct {
}

type EditResponsePb struct {
}

type EnforceClusterComplianceRequestPb struct {
	ClusterId    string `json:"cluster_id"`
	ValidateOnly bool   `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnforceClusterComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnforceClusterComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnforceClusterComplianceResponsePb struct {
	Changes    []ClusterSettingsChangePb `json:"changes,omitempty"`
	HasChanges bool                      `json:"has_changes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnforceClusterComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnforceClusterComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnvironmentPb struct {
	Client             string   `json:"client,omitempty"`
	Dependencies       []string `json:"dependencies,omitempty"`
	EnvironmentVersion string   `json:"environment_version,omitempty"`
	JarDependencies    []string `json:"jar_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnvironmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnvironmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EventDetailsPb struct {
	Attributes                          *ClusterAttributesPb      `json:"attributes,omitempty"`
	Cause                               EventDetailsCausePb       `json:"cause,omitempty"`
	ClusterSize                         *ClusterSizePb            `json:"cluster_size,omitempty"`
	CurrentNumVcpus                     int                       `json:"current_num_vcpus,omitempty"`
	CurrentNumWorkers                   int                       `json:"current_num_workers,omitempty"`
	DidNotExpandReason                  string                    `json:"did_not_expand_reason,omitempty"`
	DiskSize                            int64                     `json:"disk_size,omitempty"`
	DriverStateMessage                  string                    `json:"driver_state_message,omitempty"`
	EnableTerminationForNodeBlocklisted bool                      `json:"enable_termination_for_node_blocklisted,omitempty"`
	FreeSpace                           int64                     `json:"free_space,omitempty"`
	InitScripts                         *InitScriptEventDetailsPb `json:"init_scripts,omitempty"`
	InstanceId                          string                    `json:"instance_id,omitempty"`
	JobRunName                          string                    `json:"job_run_name,omitempty"`
	PreviousAttributes                  *ClusterAttributesPb      `json:"previous_attributes,omitempty"`
	PreviousClusterSize                 *ClusterSizePb            `json:"previous_cluster_size,omitempty"`
	PreviousDiskSize                    int64                     `json:"previous_disk_size,omitempty"`
	Reason                              *TerminationReasonPb      `json:"reason,omitempty"`
	TargetNumVcpus                      int                       `json:"target_num_vcpus,omitempty"`
	TargetNumWorkers                    int                       `json:"target_num_workers,omitempty"`
	User                                string                    `json:"user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EventDetailsCausePb string

const EventDetailsCauseAutorecovery EventDetailsCausePb = `AUTORECOVERY`

const EventDetailsCauseAutoscale EventDetailsCausePb = `AUTOSCALE`

const EventDetailsCauseReplaceBadNodes EventDetailsCausePb = `REPLACE_BAD_NODES`

const EventDetailsCauseUserRequest EventDetailsCausePb = `USER_REQUEST`

type EventTypePb string

const EventTypeAddNodesFailed EventTypePb = `ADD_NODES_FAILED`

const EventTypeAutomaticClusterUpdate EventTypePb = `AUTOMATIC_CLUSTER_UPDATE`

const EventTypeAutoscalingBackoff EventTypePb = `AUTOSCALING_BACKOFF`

const EventTypeAutoscalingFailed EventTypePb = `AUTOSCALING_FAILED`

const EventTypeAutoscalingStatsReport EventTypePb = `AUTOSCALING_STATS_REPORT`

const EventTypeClusterMigrated EventTypePb = `CLUSTER_MIGRATED`

const EventTypeCreating EventTypePb = `CREATING`

const EventTypeDbfsDown EventTypePb = `DBFS_DOWN`

const EventTypeDidNotExpandDisk EventTypePb = `DID_NOT_EXPAND_DISK`

const EventTypeDriverHealthy EventTypePb = `DRIVER_HEALTHY`

const EventTypeDriverNotResponding EventTypePb = `DRIVER_NOT_RESPONDING`

const EventTypeDriverUnavailable EventTypePb = `DRIVER_UNAVAILABLE`

const EventTypeEdited EventTypePb = `EDITED`

const EventTypeExpandedDisk EventTypePb = `EXPANDED_DISK`

const EventTypeFailedToExpandDisk EventTypePb = `FAILED_TO_EXPAND_DISK`

const EventTypeInitScriptsFinished EventTypePb = `INIT_SCRIPTS_FINISHED`

const EventTypeInitScriptsStarted EventTypePb = `INIT_SCRIPTS_STARTED`

const EventTypeMetastoreDown EventTypePb = `METASTORE_DOWN`

const EventTypeNodesLost EventTypePb = `NODES_LOST`

const EventTypeNodeBlacklisted EventTypePb = `NODE_BLACKLISTED`

const EventTypeNodeExcludedDecommissioned EventTypePb = `NODE_EXCLUDED_DECOMMISSIONED`

const EventTypePinned EventTypePb = `PINNED`

const EventTypeResizing EventTypePb = `RESIZING`

const EventTypeRestarting EventTypePb = `RESTARTING`

const EventTypeRunning EventTypePb = `RUNNING`

const EventTypeSparkException EventTypePb = `SPARK_EXCEPTION`

const EventTypeStarting EventTypePb = `STARTING`

const EventTypeTerminating EventTypePb = `TERMINATING`

const EventTypeUnpinned EventTypePb = `UNPINNED`

const EventTypeUpsizeCompleted EventTypePb = `UPSIZE_COMPLETED`

type GcpAttributesPb struct {
	Availability            GcpAvailabilityPb `json:"availability,omitempty"`
	BootDiskSize            int               `json:"boot_disk_size,omitempty"`
	GoogleServiceAccount    string            `json:"google_service_account,omitempty"`
	LocalSsdCount           int               `json:"local_ssd_count,omitempty"`
	UsePreemptibleExecutors bool              `json:"use_preemptible_executors,omitempty"`
	ZoneId                  string            `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GcpAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GcpAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GcpAvailabilityPb string

const GcpAvailabilityOnDemandGcp GcpAvailabilityPb = `ON_DEMAND_GCP`

const GcpAvailabilityPreemptibleGcp GcpAvailabilityPb = `PREEMPTIBLE_GCP`

const GcpAvailabilityPreemptibleWithFallbackGcp GcpAvailabilityPb = `PREEMPTIBLE_WITH_FALLBACK_GCP`

type GcsStorageInfoPb struct {
	Destination string `json:"destination"`
}

type GetClusterComplianceRequestPb struct {
	ClusterId string `json:"-" url:"cluster_id"`
}

type GetClusterComplianceResponsePb struct {
	IsCompliant bool              `json:"is_compliant,omitempty"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetClusterComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetClusterComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetClusterPermissionLevelsRequestPb struct {
	ClusterId string `json:"-" url:"-"`
}

type GetClusterPermissionLevelsResponsePb struct {
	PermissionLevels []ClusterPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetClusterPermissionsRequestPb struct {
	ClusterId string `json:"-" url:"-"`
}

type GetClusterPolicyPermissionLevelsRequestPb struct {
	ClusterPolicyId string `json:"-" url:"-"`
}

type GetClusterPolicyPermissionLevelsResponsePb struct {
	PermissionLevels []ClusterPolicyPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetClusterPolicyPermissionsRequestPb struct {
	ClusterPolicyId string `json:"-" url:"-"`
}

type GetClusterPolicyRequestPb struct {
	PolicyId string `json:"-" url:"policy_id"`
}

type GetClusterRequestPb struct {
	ClusterId string `json:"-" url:"cluster_id"`
}

type GetEventsPb struct {
	ClusterId  string           `json:"cluster_id"`
	EndTime    int64            `json:"end_time,omitempty"`
	EventTypes []EventTypePb    `json:"event_types,omitempty"`
	Limit      int64            `json:"limit,omitempty"`
	Offset     int64            `json:"offset,omitempty"`
	Order      GetEventsOrderPb `json:"order,omitempty"`
	PageSize   int              `json:"page_size,omitempty"`
	PageToken  string           `json:"page_token,omitempty"`
	StartTime  int64            `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetEventsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetEventsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetEventsOrderPb string

const GetEventsOrderAsc GetEventsOrderPb = `ASC`

const GetEventsOrderDesc GetEventsOrderPb = `DESC`

type GetEventsResponsePb struct {
	Events        []ClusterEventPb `json:"events,omitempty"`
	NextPage      *GetEventsPb     `json:"next_page,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`
	PrevPageToken string           `json:"prev_page_token,omitempty"`
	TotalCount    int64            `json:"total_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetEventsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetEventsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetGlobalInitScriptRequestPb struct {
	ScriptId string `json:"-" url:"-"`
}

type GetInstancePoolPb struct {
	AwsAttributes                      *InstancePoolAwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributesPb `json:"azure_attributes,omitempty"`
	CustomTags                         map[string]string              `json:"custom_tags,omitempty"`
	DefaultTags                        map[string]string              `json:"default_tags,omitempty"`
	DiskSpec                           *DiskSpecPb                    `json:"disk_spec,omitempty"`
	EnableElasticDisk                  bool                           `json:"enable_elastic_disk,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributesPb   `json:"gcp_attributes,omitempty"`
	IdleInstanceAutoterminationMinutes int                            `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolId                     string                         `json:"instance_pool_id"`
	InstancePoolName                   string                         `json:"instance_pool_name,omitempty"`
	MaxCapacity                        int                            `json:"max_capacity,omitempty"`
	MinIdleInstances                   int                            `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string                         `json:"node_type_id,omitempty"`
	PreloadedDockerImages              []DockerImagePb                `json:"preloaded_docker_images,omitempty"`
	PreloadedSparkVersions             []string                       `json:"preloaded_spark_versions,omitempty"`
	RemoteDiskThroughput               int                            `json:"remote_disk_throughput,omitempty"`
	State                              InstancePoolStatePb            `json:"state,omitempty"`
	Stats                              *InstancePoolStatsPb           `json:"stats,omitempty"`
	Status                             *InstancePoolStatusPb          `json:"status,omitempty"`
	TotalInitialRemoteDiskSize         int                            `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetInstancePoolPermissionLevelsRequestPb struct {
	InstancePoolId string `json:"-" url:"-"`
}

type GetInstancePoolPermissionLevelsResponsePb struct {
	PermissionLevels []InstancePoolPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetInstancePoolPermissionsRequestPb struct {
	InstancePoolId string `json:"-" url:"-"`
}

type GetInstancePoolRequestPb struct {
	InstancePoolId string `json:"-" url:"instance_pool_id"`
}

type GetPolicyFamilyRequestPb struct {
	PolicyFamilyId string `json:"-" url:"-"`
	Version        int64  `json:"-" url:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPolicyFamilyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPolicyFamilyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSparkVersionsResponsePb struct {
	Versions []SparkVersionPb `json:"versions,omitempty"`
}

type GlobalInitScriptCreateRequestPb struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Name     string `json:"name"`
	Position int    `json:"position,omitempty"`
	Script   string `json:"script"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GlobalInitScriptCreateRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GlobalInitScriptCreateRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GlobalInitScriptDetailsPb struct {
	CreatedAt int    `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Name      string `json:"name,omitempty"`
	Position  int    `json:"position,omitempty"`
	ScriptId  string `json:"script_id,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GlobalInitScriptDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GlobalInitScriptDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GlobalInitScriptDetailsWithContentPb struct {
	CreatedAt int    `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Name      string `json:"name,omitempty"`
	Position  int    `json:"position,omitempty"`
	Script    string `json:"script,omitempty"`
	ScriptId  string `json:"script_id,omitempty"`
	UpdatedAt int    `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GlobalInitScriptDetailsWithContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GlobalInitScriptDetailsWithContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GlobalInitScriptUpdateRequestPb struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Name     string `json:"name"`
	Position int    `json:"position,omitempty"`
	Script   string `json:"script"`
	ScriptId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GlobalInitScriptUpdateRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GlobalInitScriptUpdateRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InitScriptEventDetailsPb struct {
	Cluster         []InitScriptInfoAndExecutionDetailsPb `json:"cluster,omitempty"`
	Global          []InitScriptInfoAndExecutionDetailsPb `json:"global,omitempty"`
	ReportedForNode string                                `json:"reported_for_node,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InitScriptEventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InitScriptEventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InitScriptExecutionDetailsInitScriptExecutionStatusPb string

const InitScriptExecutionDetailsInitScriptExecutionStatusFailedExecution InitScriptExecutionDetailsInitScriptExecutionStatusPb = `FAILED_EXECUTION`

const InitScriptExecutionDetailsInitScriptExecutionStatusFailedFetch InitScriptExecutionDetailsInitScriptExecutionStatusPb = `FAILED_FETCH`

const InitScriptExecutionDetailsInitScriptExecutionStatusFuseMountFailed InitScriptExecutionDetailsInitScriptExecutionStatusPb = `FUSE_MOUNT_FAILED`

const InitScriptExecutionDetailsInitScriptExecutionStatusNotExecuted InitScriptExecutionDetailsInitScriptExecutionStatusPb = `NOT_EXECUTED`

const InitScriptExecutionDetailsInitScriptExecutionStatusSkipped InitScriptExecutionDetailsInitScriptExecutionStatusPb = `SKIPPED`

const InitScriptExecutionDetailsInitScriptExecutionStatusSucceeded InitScriptExecutionDetailsInitScriptExecutionStatusPb = `SUCCEEDED`

const InitScriptExecutionDetailsInitScriptExecutionStatusUnknown InitScriptExecutionDetailsInitScriptExecutionStatusPb = `UNKNOWN`

type InitScriptInfoPb struct {
	Abfss     *Adlsgen2InfoPb         `json:"abfss,omitempty"`
	Dbfs      *DbfsStorageInfoPb      `json:"dbfs,omitempty"`
	File      *LocalFileInfoPb        `json:"file,omitempty"`
	Gcs       *GcsStorageInfoPb       `json:"gcs,omitempty"`
	S3        *S3StorageInfoPb        `json:"s3,omitempty"`
	Volumes   *VolumesStorageInfoPb   `json:"volumes,omitempty"`
	Workspace *WorkspaceStorageInfoPb `json:"workspace,omitempty"`
}

type InitScriptInfoAndExecutionDetailsPb struct {
	Abfss                    *Adlsgen2InfoPb                                       `json:"abfss,omitempty"`
	Dbfs                     *DbfsStorageInfoPb                                    `json:"dbfs,omitempty"`
	ErrorMessage             string                                                `json:"error_message,omitempty"`
	ExecutionDurationSeconds int                                                   `json:"execution_duration_seconds,omitempty"`
	File                     *LocalFileInfoPb                                      `json:"file,omitempty"`
	Gcs                      *GcsStorageInfoPb                                     `json:"gcs,omitempty"`
	S3                       *S3StorageInfoPb                                      `json:"s3,omitempty"`
	Status                   InitScriptExecutionDetailsInitScriptExecutionStatusPb `json:"status,omitempty"`
	Volumes                  *VolumesStorageInfoPb                                 `json:"volumes,omitempty"`
	Workspace                *WorkspaceStorageInfoPb                               `json:"workspace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InitScriptInfoAndExecutionDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InitScriptInfoAndExecutionDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstallLibrariesPb struct {
	ClusterId string      `json:"cluster_id"`
	Libraries []LibraryPb `json:"libraries"`
}

type InstallLibrariesResponsePb struct {
}

type InstancePoolAccessControlRequestPb struct {
	GroupName            string                        `json:"group_name,omitempty"`
	PermissionLevel      InstancePoolPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                        `json:"service_principal_name,omitempty"`
	UserName             string                        `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAccessControlResponsePb struct {
	AllPermissions       []InstancePoolPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                     `json:"display_name,omitempty"`
	GroupName            string                     `json:"group_name,omitempty"`
	ServicePrincipalName string                     `json:"service_principal_name,omitempty"`
	UserName             string                     `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAndStatsPb struct {
	AwsAttributes                      *InstancePoolAwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributesPb `json:"azure_attributes,omitempty"`
	CustomTags                         map[string]string              `json:"custom_tags,omitempty"`
	DefaultTags                        map[string]string              `json:"default_tags,omitempty"`
	DiskSpec                           *DiskSpecPb                    `json:"disk_spec,omitempty"`
	EnableElasticDisk                  bool                           `json:"enable_elastic_disk,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributesPb   `json:"gcp_attributes,omitempty"`
	IdleInstanceAutoterminationMinutes int                            `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolId                     string                         `json:"instance_pool_id,omitempty"`
	InstancePoolName                   string                         `json:"instance_pool_name,omitempty"`
	MaxCapacity                        int                            `json:"max_capacity,omitempty"`
	MinIdleInstances                   int                            `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string                         `json:"node_type_id,omitempty"`
	PreloadedDockerImages              []DockerImagePb                `json:"preloaded_docker_images,omitempty"`
	PreloadedSparkVersions             []string                       `json:"preloaded_spark_versions,omitempty"`
	RemoteDiskThroughput               int                            `json:"remote_disk_throughput,omitempty"`
	State                              InstancePoolStatePb            `json:"state,omitempty"`
	Stats                              *InstancePoolStatsPb           `json:"stats,omitempty"`
	Status                             *InstancePoolStatusPb          `json:"status,omitempty"`
	TotalInitialRemoteDiskSize         int                            `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolAndStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolAndStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAwsAttributesPb struct {
	Availability        InstancePoolAwsAttributesAvailabilityPb `json:"availability,omitempty"`
	SpotBidPricePercent int                                     `json:"spot_bid_price_percent,omitempty"`
	ZoneId              string                                  `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolAwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolAwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAwsAttributesAvailabilityPb string

const InstancePoolAwsAttributesAvailabilityOnDemand InstancePoolAwsAttributesAvailabilityPb = `ON_DEMAND`

const InstancePoolAwsAttributesAvailabilitySpot InstancePoolAwsAttributesAvailabilityPb = `SPOT`

type InstancePoolAzureAttributesPb struct {
	Availability    InstancePoolAzureAttributesAvailabilityPb `json:"availability,omitempty"`
	SpotBidMaxPrice float64                                   `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolAzureAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolAzureAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAzureAttributesAvailabilityPb string

const InstancePoolAzureAttributesAvailabilityOnDemandAzure InstancePoolAzureAttributesAvailabilityPb = `ON_DEMAND_AZURE`

const InstancePoolAzureAttributesAvailabilitySpotAzure InstancePoolAzureAttributesAvailabilityPb = `SPOT_AZURE`

type InstancePoolGcpAttributesPb struct {
	GcpAvailability GcpAvailabilityPb `json:"gcp_availability,omitempty"`
	LocalSsdCount   int               `json:"local_ssd_count,omitempty"`
	ZoneId          string            `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolGcpAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolGcpAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermissionPb struct {
	Inherited           bool                          `json:"inherited,omitempty"`
	InheritedFromObject []string                      `json:"inherited_from_object,omitempty"`
	PermissionLevel     InstancePoolPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermissionLevelPb string

const InstancePoolPermissionLevelCanAttachTo InstancePoolPermissionLevelPb = `CAN_ATTACH_TO`

const InstancePoolPermissionLevelCanManage InstancePoolPermissionLevelPb = `CAN_MANAGE`

type InstancePoolPermissionsPb struct {
	AccessControlList []InstancePoolAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                                `json:"object_id,omitempty"`
	ObjectType        string                                `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermissionsDescriptionPb struct {
	Description     string                        `json:"description,omitempty"`
	PermissionLevel InstancePoolPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermissionsRequestPb struct {
	AccessControlList []InstancePoolAccessControlRequestPb `json:"access_control_list,omitempty"`
	InstancePoolId    string                               `json:"-" url:"-"`
}

type InstancePoolStatePb string

const InstancePoolStateActive InstancePoolStatePb = `ACTIVE`

const InstancePoolStateDeleted InstancePoolStatePb = `DELETED`

const InstancePoolStateStopped InstancePoolStatePb = `STOPPED`

type InstancePoolStatsPb struct {
	IdleCount        int `json:"idle_count,omitempty"`
	PendingIdleCount int `json:"pending_idle_count,omitempty"`
	PendingUsedCount int `json:"pending_used_count,omitempty"`
	UsedCount        int `json:"used_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstancePoolStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstancePoolStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolStatusPb struct {
	PendingInstanceErrors []PendingInstanceErrorPb `json:"pending_instance_errors,omitempty"`
}

type InstanceProfilePb struct {
	IamRoleArn            string `json:"iam_role_arn,omitempty"`
	InstanceProfileArn    string `json:"instance_profile_arn"`
	IsMetaInstanceProfile bool   `json:"is_meta_instance_profile,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *InstanceProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st InstanceProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type KindPb string

const KindClassicPreview KindPb = `CLASSIC_PREVIEW`

type LanguagePb string

const LanguagePython LanguagePb = `python`

const LanguageR LanguagePb = `r`

const LanguageScala LanguagePb = `scala`

const LanguageSql LanguagePb = `sql`

type LibraryPb struct {
	Cran         *RCranLibraryPb      `json:"cran,omitempty"`
	Egg          string               `json:"egg,omitempty"`
	Jar          string               `json:"jar,omitempty"`
	Maven        *MavenLibraryPb      `json:"maven,omitempty"`
	Pypi         *PythonPyPiLibraryPb `json:"pypi,omitempty"`
	Requirements string               `json:"requirements,omitempty"`
	Whl          string               `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LibraryFullStatusPb struct {
	IsLibraryForAllClusters bool                   `json:"is_library_for_all_clusters,omitempty"`
	Library                 *LibraryPb             `json:"library,omitempty"`
	Messages                []string               `json:"messages,omitempty"`
	Status                  LibraryInstallStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LibraryFullStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LibraryFullStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LibraryInstallStatusPb string

const LibraryInstallStatusFailed LibraryInstallStatusPb = `FAILED`

const LibraryInstallStatusInstalled LibraryInstallStatusPb = `INSTALLED`

const LibraryInstallStatusInstalling LibraryInstallStatusPb = `INSTALLING`

const LibraryInstallStatusPending LibraryInstallStatusPb = `PENDING`

const LibraryInstallStatusResolving LibraryInstallStatusPb = `RESOLVING`

const LibraryInstallStatusRestored LibraryInstallStatusPb = `RESTORED`

const LibraryInstallStatusSkipped LibraryInstallStatusPb = `SKIPPED`

const LibraryInstallStatusUninstallOnRestart LibraryInstallStatusPb = `UNINSTALL_ON_RESTART`

type ListAllClusterLibraryStatusesPb struct {
}

type ListAllClusterLibraryStatusesResponsePb struct {
	Statuses []ClusterLibraryStatusesPb `json:"statuses,omitempty"`
}

type ListAvailableZonesResponsePb struct {
	DefaultZone string   `json:"default_zone,omitempty"`
	Zones       []string `json:"zones,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAvailableZonesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAvailableZonesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClusterCompliancesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`
	PolicyId  string `json:"-" url:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListClusterCompliancesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListClusterCompliancesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClusterCompliancesResponsePb struct {
	Clusters      []ClusterCompliancePb `json:"clusters,omitempty"`
	NextPageToken string                `json:"next_page_token,omitempty"`
	PrevPageToken string                `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListClusterCompliancesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListClusterCompliancesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClusterPoliciesRequestPb struct {
	SortColumn ListSortColumnPb `json:"-" url:"sort_column,omitempty"`
	SortOrder  ListSortOrderPb  `json:"-" url:"sort_order,omitempty"`
}

type ListClustersFilterByPb struct {
	ClusterSources []ClusterSourcePb `json:"cluster_sources,omitempty" url:"cluster_sources,omitempty"`
	ClusterStates  []StatePb         `json:"cluster_states,omitempty" url:"cluster_states,omitempty"`
	IsPinned       bool              `json:"is_pinned,omitempty" url:"is_pinned,omitempty"`
	PolicyId       string            `json:"policy_id,omitempty" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListClustersFilterByPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListClustersFilterByPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClustersRequestPb struct {
	FilterBy  *ListClustersFilterByPb `json:"-" url:"filter_by,omitempty"`
	PageSize  int                     `json:"-" url:"page_size,omitempty"`
	PageToken string                  `json:"-" url:"page_token,omitempty"`
	SortBy    *ListClustersSortByPb   `json:"-" url:"sort_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListClustersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListClustersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClustersResponsePb struct {
	Clusters      []ClusterDetailsPb `json:"clusters,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`
	PrevPageToken string             `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListClustersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListClustersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClustersSortByPb struct {
	Direction ListClustersSortByDirectionPb `json:"direction,omitempty" url:"direction,omitempty"`
	Field     ListClustersSortByFieldPb     `json:"field,omitempty" url:"field,omitempty"`
}

type ListClustersSortByDirectionPb string

const ListClustersSortByDirectionAsc ListClustersSortByDirectionPb = `ASC`

const ListClustersSortByDirectionDesc ListClustersSortByDirectionPb = `DESC`

type ListClustersSortByFieldPb string

const ListClustersSortByFieldClusterName ListClustersSortByFieldPb = `CLUSTER_NAME`

const ListClustersSortByFieldDefault ListClustersSortByFieldPb = `DEFAULT`

type ListGlobalInitScriptsRequestPb struct {
}

type ListGlobalInitScriptsResponsePb struct {
	Scripts []GlobalInitScriptDetailsPb `json:"scripts,omitempty"`
}

type ListInstancePoolsPb struct {
	InstancePools []InstancePoolAndStatsPb `json:"instance_pools,omitempty"`
}

type ListInstancePoolsRequestPb struct {
}

type ListInstanceProfilesRequestPb struct {
}

type ListInstanceProfilesResponsePb struct {
	InstanceProfiles []InstanceProfilePb `json:"instance_profiles,omitempty"`
}

type ListNodeTypesRequestPb struct {
}

type ListNodeTypesResponsePb struct {
	NodeTypes []NodeTypePb `json:"node_types,omitempty"`
}

type ListPoliciesResponsePb struct {
	Policies []PolicyPb `json:"policies,omitempty"`
}

type ListPolicyFamiliesRequestPb struct {
	MaxResults int64  `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPolicyFamiliesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPolicyFamiliesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPolicyFamiliesResponsePb struct {
	NextPageToken  string           `json:"next_page_token,omitempty"`
	PolicyFamilies []PolicyFamilyPb `json:"policy_families,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPolicyFamiliesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPolicyFamiliesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSortColumnPb string

const ListSortColumnPolicyCreationTime ListSortColumnPb = `POLICY_CREATION_TIME`

const ListSortColumnPolicyName ListSortColumnPb = `POLICY_NAME`

type ListSortOrderPb string

const ListSortOrderAsc ListSortOrderPb = `ASC`

const ListSortOrderDesc ListSortOrderPb = `DESC`

type ListZonesRequestPb struct {
}

type LocalFileInfoPb struct {
	Destination string `json:"destination"`
}

type LogAnalyticsInfoPb struct {
	LogAnalyticsPrimaryKey  string `json:"log_analytics_primary_key,omitempty"`
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogAnalyticsInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogAnalyticsInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogSyncStatusPb struct {
	LastAttempted int64  `json:"last_attempted,omitempty"`
	LastException string `json:"last_exception,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogSyncStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogSyncStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MapAnyPb map[string]any

type MavenLibraryPb struct {
	Coordinates string   `json:"coordinates"`
	Exclusions  []string `json:"exclusions,omitempty"`
	Repo        string   `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MavenLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MavenLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NodeInstanceTypePb struct {
	InstanceTypeId      string `json:"instance_type_id"`
	LocalDiskSizeGb     int    `json:"local_disk_size_gb,omitempty"`
	LocalDisks          int    `json:"local_disks,omitempty"`
	LocalNvmeDiskSizeGb int    `json:"local_nvme_disk_size_gb,omitempty"`
	LocalNvmeDisks      int    `json:"local_nvme_disks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NodeInstanceTypePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NodeInstanceTypePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NodeTypePb struct {
	Category              string                   `json:"category"`
	Description           string                   `json:"description"`
	DisplayOrder          int                      `json:"display_order,omitempty"`
	InstanceTypeId        string                   `json:"instance_type_id"`
	IsDeprecated          bool                     `json:"is_deprecated,omitempty"`
	IsEncryptedInTransit  bool                     `json:"is_encrypted_in_transit,omitempty"`
	IsGraviton            bool                     `json:"is_graviton,omitempty"`
	IsHidden              bool                     `json:"is_hidden,omitempty"`
	IsIoCacheEnabled      bool                     `json:"is_io_cache_enabled,omitempty"`
	MemoryMb              int                      `json:"memory_mb"`
	NodeInfo              *CloudProviderNodeInfoPb `json:"node_info,omitempty"`
	NodeInstanceType      *NodeInstanceTypePb      `json:"node_instance_type,omitempty"`
	NodeTypeId            string                   `json:"node_type_id"`
	NumCores              float64                  `json:"num_cores"`
	NumGpus               int                      `json:"num_gpus,omitempty"`
	PhotonDriverCapable   bool                     `json:"photon_driver_capable,omitempty"`
	PhotonWorkerCapable   bool                     `json:"photon_worker_capable,omitempty"`
	SupportClusterTags    bool                     `json:"support_cluster_tags,omitempty"`
	SupportEbsVolumes     bool                     `json:"support_ebs_volumes,omitempty"`
	SupportPortForwarding bool                     `json:"support_port_forwarding,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NodeTypePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NodeTypePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PendingInstanceErrorPb struct {
	InstanceId string `json:"instance_id,omitempty"`
	Message    string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PendingInstanceErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PendingInstanceErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermanentDeleteClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

type PermanentDeleteClusterResponsePb struct {
}

type PinClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

type PinClusterResponsePb struct {
}

type PolicyPb struct {
	CreatedAtTimestamp              int64       `json:"created_at_timestamp,omitempty"`
	CreatorUserName                 string      `json:"creator_user_name,omitempty"`
	Definition                      string      `json:"definition,omitempty"`
	Description                     string      `json:"description,omitempty"`
	IsDefault                       bool        `json:"is_default,omitempty"`
	Libraries                       []LibraryPb `json:"libraries,omitempty"`
	MaxClustersPerUser              int64       `json:"max_clusters_per_user,omitempty"`
	Name                            string      `json:"name,omitempty"`
	PolicyFamilyDefinitionOverrides string      `json:"policy_family_definition_overrides,omitempty"`
	PolicyFamilyId                  string      `json:"policy_family_id,omitempty"`
	PolicyId                        string      `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PolicyFamilyPb struct {
	Definition     string `json:"definition,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	PolicyFamilyId string `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PolicyFamilyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PolicyFamilyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PythonPyPiLibraryPb struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PythonPyPiLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PythonPyPiLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RCranLibraryPb struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RCranLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RCranLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RemoveInstanceProfilePb struct {
	InstanceProfileArn string `json:"instance_profile_arn"`
}

type RemoveResponsePb struct {
}

type ResizeClusterPb struct {
	Autoscale  *AutoScalePb `json:"autoscale,omitempty"`
	ClusterId  string       `json:"cluster_id"`
	NumWorkers int          `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResizeClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResizeClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResizeClusterResponsePb struct {
}

type RestartClusterPb struct {
	ClusterId   string `json:"cluster_id"`
	RestartUser string `json:"restart_user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RestartClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RestartClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RestartClusterResponsePb struct {
}

type ResultTypePb string

const ResultTypeError ResultTypePb = `error`

const ResultTypeImage ResultTypePb = `image`

const ResultTypeImages ResultTypePb = `images`

const ResultTypeTable ResultTypePb = `table`

const ResultTypeText ResultTypePb = `text`

type ResultsPb struct {
	Cause        string           `json:"cause,omitempty"`
	Data         any              `json:"data,omitempty"`
	FileName     string           `json:"fileName,omitempty"`
	FileNames    []string         `json:"fileNames,omitempty"`
	IsJsonSchema bool             `json:"isJsonSchema,omitempty"`
	Pos          int              `json:"pos,omitempty"`
	ResultType   ResultTypePb     `json:"resultType,omitempty"`
	Schema       []map[string]any `json:"schema,omitempty"`
	Summary      string           `json:"summary,omitempty"`
	Truncated    bool             `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RuntimeEnginePb string

const RuntimeEngineNull RuntimeEnginePb = `NULL`

const RuntimeEnginePhoton RuntimeEnginePb = `PHOTON`

const RuntimeEngineStandard RuntimeEnginePb = `STANDARD`

type S3StorageInfoPb struct {
	CannedAcl        string `json:"canned_acl,omitempty"`
	Destination      string `json:"destination"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	Region           string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *S3StorageInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st S3StorageInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkNodePb struct {
	HostPrivateIp     string                    `json:"host_private_ip,omitempty"`
	InstanceId        string                    `json:"instance_id,omitempty"`
	NodeAwsAttributes *SparkNodeAwsAttributesPb `json:"node_aws_attributes,omitempty"`
	NodeId            string                    `json:"node_id,omitempty"`
	PrivateIp         string                    `json:"private_ip,omitempty"`
	PublicDns         string                    `json:"public_dns,omitempty"`
	StartTimestamp    int64                     `json:"start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SparkNodePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SparkNodePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkNodeAwsAttributesPb struct {
	IsSpot bool `json:"is_spot,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SparkNodeAwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SparkNodeAwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkVersionPb struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SparkVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SparkVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkVersionsRequestPb struct {
}

type StartClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

type StartClusterResponsePb struct {
}

type StatePb string

const StateError StatePb = `ERROR`

const StatePending StatePb = `PENDING`

const StateResizing StatePb = `RESIZING`

const StateRestarting StatePb = `RESTARTING`

const StateRunning StatePb = `RUNNING`

const StateTerminated StatePb = `TERMINATED`

const StateTerminating StatePb = `TERMINATING`

const StateUnknown StatePb = `UNKNOWN`

type TerminationReasonPb struct {
	Code       TerminationReasonCodePb `json:"code,omitempty"`
	Parameters map[string]string       `json:"parameters,omitempty"`
	Type       TerminationReasonTypePb `json:"type,omitempty"`
}

type TerminationReasonCodePb string

const TerminationReasonCodeAbuseDetected TerminationReasonCodePb = `ABUSE_DETECTED`

const TerminationReasonCodeAccessTokenFailure TerminationReasonCodePb = `ACCESS_TOKEN_FAILURE`

const TerminationReasonCodeAllocationTimeout TerminationReasonCodePb = `ALLOCATION_TIMEOUT`

const TerminationReasonCodeAllocationTimeoutNodeDaemonNotReady TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`

const TerminationReasonCodeAllocationTimeoutNoHealthyAndWarmedUpClusters TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoHealthyClusters TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoMatchedClusters TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoReadyClusters TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoUnallocatedClusters TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoWarmedUpClusters TerminationReasonCodePb = `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCodePb = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCodePb = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInaccessibleKmsKeyFailure TerminationReasonCodePb = `AWS_INACCESSIBLE_KMS_KEY_FAILURE`

const TerminationReasonCodeAwsInstanceProfileUpdateFailure TerminationReasonCodePb = `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCodePb = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCodePb = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsInvalidKeyPair TerminationReasonCodePb = `AWS_INVALID_KEY_PAIR`

const TerminationReasonCodeAwsInvalidKmsKeyState TerminationReasonCodePb = `AWS_INVALID_KMS_KEY_STATE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCodePb = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCodePb = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsResourceQuotaExceeded TerminationReasonCodePb = `AWS_RESOURCE_QUOTA_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCodePb = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCodePb = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCodePb = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCodePb = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCodePb = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzurePackedDeploymentPartialFailure TerminationReasonCodePb = `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCodePb = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCodePb = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCodePb = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCodePb = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCodePb = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCodePb = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCodePb = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCodePb = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeBootstrapTimeoutDueToMisconfig TerminationReasonCodePb = `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`

const TerminationReasonCodeBudgetPolicyLimitEnforcementActivated TerminationReasonCodePb = `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`

const TerminationReasonCodeBudgetPolicyResolutionFailure TerminationReasonCodePb = `BUDGET_POLICY_RESOLUTION_FAILURE`

const TerminationReasonCodeCloudAccountSetupFailure TerminationReasonCodePb = `CLOUD_ACCOUNT_SETUP_FAILURE`

const TerminationReasonCodeCloudOperationCancelled TerminationReasonCodePb = `CLOUD_OPERATION_CANCELLED`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCodePb = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderInstanceNotLaunched TerminationReasonCodePb = `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCodePb = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailureDueToMisconfig TerminationReasonCodePb = `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCodePb = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderResourceStockoutDueToMisconfig TerminationReasonCodePb = `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCodePb = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeClusterOperationThrottled TerminationReasonCodePb = `CLUSTER_OPERATION_THROTTLED`

const TerminationReasonCodeClusterOperationTimeout TerminationReasonCodePb = `CLUSTER_OPERATION_TIMEOUT`

const TerminationReasonCodeCommunicationLost TerminationReasonCodePb = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCodePb = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCodePb = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailureDueToMisconfig TerminationReasonCodePb = `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCodePb = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDataAccessConfigChanged TerminationReasonCodePb = `DATA_ACCESS_CONFIG_CHANGED`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCodePb = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDisasterRecoveryReplication TerminationReasonCodePb = `DISASTER_RECOVERY_REPLICATION`

const TerminationReasonCodeDnsResolutionError TerminationReasonCodePb = `DNS_RESOLUTION_ERROR`

const TerminationReasonCodeDockerContainerCreationException TerminationReasonCodePb = `DOCKER_CONTAINER_CREATION_EXCEPTION`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCodePb = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDockerImageTooLargeForInstanceException TerminationReasonCodePb = `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`

const TerminationReasonCodeDockerInvalidOsException TerminationReasonCodePb = `DOCKER_INVALID_OS_EXCEPTION`

const TerminationReasonCodeDriverDnsResolutionFailure TerminationReasonCodePb = `DRIVER_DNS_RESOLUTION_FAILURE`

const TerminationReasonCodeDriverEviction TerminationReasonCodePb = `DRIVER_EVICTION`

const TerminationReasonCodeDriverLaunchTimeout TerminationReasonCodePb = `DRIVER_LAUNCH_TIMEOUT`

const TerminationReasonCodeDriverNodeUnreachable TerminationReasonCodePb = `DRIVER_NODE_UNREACHABLE`

const TerminationReasonCodeDriverOutOfDisk TerminationReasonCodePb = `DRIVER_OUT_OF_DISK`

const TerminationReasonCodeDriverOutOfMemory TerminationReasonCodePb = `DRIVER_OUT_OF_MEMORY`

const TerminationReasonCodeDriverPodCreationFailure TerminationReasonCodePb = `DRIVER_POD_CREATION_FAILURE`

const TerminationReasonCodeDriverUnexpectedFailure TerminationReasonCodePb = `DRIVER_UNEXPECTED_FAILURE`

const TerminationReasonCodeDriverUnhealthy TerminationReasonCodePb = `DRIVER_UNHEALTHY`

const TerminationReasonCodeDriverUnreachable TerminationReasonCodePb = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCodePb = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeDynamicSparkConfSizeExceeded TerminationReasonCodePb = `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`

const TerminationReasonCodeEosSparkImage TerminationReasonCodePb = `EOS_SPARK_IMAGE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCodePb = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeExecutorPodUnscheduled TerminationReasonCodePb = `EXECUTOR_POD_UNSCHEDULED`

const TerminationReasonCodeGcpApiRateQuotaExceeded TerminationReasonCodePb = `GCP_API_RATE_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpDeniedByOrgPolicy TerminationReasonCodePb = `GCP_DENIED_BY_ORG_POLICY`

const TerminationReasonCodeGcpForbidden TerminationReasonCodePb = `GCP_FORBIDDEN`

const TerminationReasonCodeGcpIamTimeout TerminationReasonCodePb = `GCP_IAM_TIMEOUT`

const TerminationReasonCodeGcpInaccessibleKmsKeyFailure TerminationReasonCodePb = `GCP_INACCESSIBLE_KMS_KEY_FAILURE`

const TerminationReasonCodeGcpInsufficientCapacity TerminationReasonCodePb = `GCP_INSUFFICIENT_CAPACITY`

const TerminationReasonCodeGcpIpSpaceExhausted TerminationReasonCodePb = `GCP_IP_SPACE_EXHAUSTED`

const TerminationReasonCodeGcpKmsKeyPermissionDenied TerminationReasonCodePb = `GCP_KMS_KEY_PERMISSION_DENIED`

const TerminationReasonCodeGcpNotFound TerminationReasonCodePb = `GCP_NOT_FOUND`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCodePb = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpResourceQuotaExceeded TerminationReasonCodePb = `GCP_RESOURCE_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountAccessDenied TerminationReasonCodePb = `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCodePb = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGcpServiceAccountNotFound TerminationReasonCodePb = `GCP_SERVICE_ACCOUNT_NOT_FOUND`

const TerminationReasonCodeGcpSubnetNotReady TerminationReasonCodePb = `GCP_SUBNET_NOT_READY`

const TerminationReasonCodeGcpTrustedImageProjectsViolated TerminationReasonCodePb = `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`

const TerminationReasonCodeGkeBasedClusterTermination TerminationReasonCodePb = `GKE_BASED_CLUSTER_TERMINATION`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCodePb = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCodePb = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCodePb = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCodePb = `INACTIVITY`

const TerminationReasonCodeInitContainerNotFinished TerminationReasonCodePb = `INIT_CONTAINER_NOT_FINISHED`

const TerminationReasonCodeInitScriptFailure TerminationReasonCodePb = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCodePb = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstancePoolMaxCapacityReached TerminationReasonCodePb = `INSTANCE_POOL_MAX_CAPACITY_REACHED`

const TerminationReasonCodeInstancePoolNotFound TerminationReasonCodePb = `INSTANCE_POOL_NOT_FOUND`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCodePb = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInstanceUnreachableDueToMisconfig TerminationReasonCodePb = `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`

const TerminationReasonCodeInternalCapacityFailure TerminationReasonCodePb = `INTERNAL_CAPACITY_FAILURE`

const TerminationReasonCodeInternalError TerminationReasonCodePb = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCodePb = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidAwsParameter TerminationReasonCodePb = `INVALID_AWS_PARAMETER`

const TerminationReasonCodeInvalidInstancePlacementProtocol TerminationReasonCodePb = `INVALID_INSTANCE_PLACEMENT_PROTOCOL`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCodePb = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeInvalidWorkerImageFailure TerminationReasonCodePb = `INVALID_WORKER_IMAGE_FAILURE`

const TerminationReasonCodeInPenaltyBox TerminationReasonCodePb = `IN_PENALTY_BOX`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCodePb = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCodePb = `JOB_FINISHED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCodePb = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCodePb = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeLazyAllocationTimeout TerminationReasonCodePb = `LAZY_ALLOCATION_TIMEOUT`

const TerminationReasonCodeMaintenanceMode TerminationReasonCodePb = `MAINTENANCE_MODE`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCodePb = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCodePb = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetvisorSetupTimeout TerminationReasonCodePb = `NETVISOR_SETUP_TIMEOUT`

const TerminationReasonCodeNetworkCheckControlPlaneFailure TerminationReasonCodePb = `NETWORK_CHECK_CONTROL_PLANE_FAILURE`

const TerminationReasonCodeNetworkCheckDnsServerFailure TerminationReasonCodePb = `NETWORK_CHECK_DNS_SERVER_FAILURE`

const TerminationReasonCodeNetworkCheckMetadataEndpointFailure TerminationReasonCodePb = `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE`

const TerminationReasonCodeNetworkCheckMultipleComponentsFailure TerminationReasonCodePb = `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE`

const TerminationReasonCodeNetworkCheckNicFailure TerminationReasonCodePb = `NETWORK_CHECK_NIC_FAILURE`

const TerminationReasonCodeNetworkCheckStorageFailure TerminationReasonCodePb = `NETWORK_CHECK_STORAGE_FAILURE`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCodePb = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCodePb = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNoMatchedK8s TerminationReasonCodePb = `NO_MATCHED_K8S`

const TerminationReasonCodeNoMatchedK8sTestingTag TerminationReasonCodePb = `NO_MATCHED_K8S_TESTING_TAG`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCodePb = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCodePb = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodePodAssignmentFailure TerminationReasonCodePb = `POD_ASSIGNMENT_FAILURE`

const TerminationReasonCodePodSchedulingFailure TerminationReasonCodePb = `POD_SCHEDULING_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCodePb = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCodePb = `REQUEST_THROTTLED`

const TerminationReasonCodeResourceUsageBlocked TerminationReasonCodePb = `RESOURCE_USAGE_BLOCKED`

const TerminationReasonCodeSecretCreationFailure TerminationReasonCodePb = `SECRET_CREATION_FAILURE`

const TerminationReasonCodeSecretPermissionDenied TerminationReasonCodePb = `SECRET_PERMISSION_DENIED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCodePb = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityAgentsFailedInitialVerification TerminationReasonCodePb = `SECURITY_AGENTS_FAILED_INITIAL_VERIFICATION`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCodePb = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCodePb = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeServerlessLongRunningTerminated TerminationReasonCodePb = `SERVERLESS_LONG_RUNNING_TERMINATED`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCodePb = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCodePb = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCodePb = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCodePb = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkImageDownloadThrottled TerminationReasonCodePb = `SPARK_IMAGE_DOWNLOAD_THROTTLED`

const TerminationReasonCodeSparkImageNotFound TerminationReasonCodePb = `SPARK_IMAGE_NOT_FOUND`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCodePb = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCodePb = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeSshBootstrapFailure TerminationReasonCodePb = `SSH_BOOTSTRAP_FAILURE`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCodePb = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStorageDownloadFailureDueToMisconfig TerminationReasonCodePb = `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeStorageDownloadFailureSlow TerminationReasonCodePb = `STORAGE_DOWNLOAD_FAILURE_SLOW`

const TerminationReasonCodeStorageDownloadFailureThrottled TerminationReasonCodePb = `STORAGE_DOWNLOAD_FAILURE_THROTTLED`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCodePb = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCodePb = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCodePb = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCodePb = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCodePb = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnexpectedPodRecreation TerminationReasonCodePb = `UNEXPECTED_POD_RECREATION`

const TerminationReasonCodeUnknown TerminationReasonCodePb = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCodePb = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCodePb = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserInitiatedVmTermination TerminationReasonCodePb = `USER_INITIATED_VM_TERMINATION`

const TerminationReasonCodeUserRequest TerminationReasonCodePb = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCodePb = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCodePb = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCodePb = `WORKSPACE_CONFIGURATION_ERROR`

const TerminationReasonCodeWorkspaceUpdate TerminationReasonCodePb = `WORKSPACE_UPDATE`

type TerminationReasonTypePb string

const TerminationReasonTypeClientError TerminationReasonTypePb = `CLIENT_ERROR`

const TerminationReasonTypeCloudFailure TerminationReasonTypePb = `CLOUD_FAILURE`

const TerminationReasonTypeServiceFault TerminationReasonTypePb = `SERVICE_FAULT`

const TerminationReasonTypeSuccess TerminationReasonTypePb = `SUCCESS`

type UninstallLibrariesPb struct {
	ClusterId string      `json:"cluster_id"`
	Libraries []LibraryPb `json:"libraries"`
}

type UninstallLibrariesResponsePb struct {
}

type UnpinClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

type UnpinClusterResponsePb struct {
}

type UpdateClusterPb struct {
	Cluster    *UpdateClusterResourcePb `json:"cluster,omitempty"`
	ClusterId  string                   `json:"cluster_id"`
	UpdateMask string                   `json:"update_mask"`
}

type UpdateClusterResourcePb struct {
	Autoscale                  *AutoScalePb       `json:"autoscale,omitempty"`
	AutoterminationMinutes     int                `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributesPb `json:"azure_attributes,omitempty"`
	ClusterLogConf             *ClusterLogConfPb  `json:"cluster_log_conf,omitempty"`
	ClusterName                string             `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string  `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityModePb `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImagePb     `json:"docker_image,omitempty"`
	DriverInstancePoolId       string             `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string             `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool               `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributesPb   `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfoPb `json:"init_scripts,omitempty"`
	InstancePoolId             string             `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool               `json:"is_single_node,omitempty"`
	Kind                       KindPb             `json:"kind,omitempty"`
	NodeTypeId                 string             `json:"node_type_id,omitempty"`
	NumWorkers                 int                `json:"num_workers,omitempty"`
	PolicyId                   string             `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEnginePb    `json:"runtime_engine,omitempty"`
	SingleUserName             string             `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string  `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string  `json:"spark_env_vars,omitempty"`
	SparkVersion               string             `json:"spark_version,omitempty"`
	SshPublicKeys              []string           `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int                `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool               `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadTypePb    `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateClusterResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateClusterResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateClusterResponsePb struct {
}

type UpdateResponsePb struct {
}

type VolumesStorageInfoPb struct {
	Destination string `json:"destination"`
}

type WorkloadTypePb struct {
	Clients ClientsTypesPb `json:"clients"`
}

type WorkspaceStorageInfoPb struct {
	Destination string `json:"destination"`
}
