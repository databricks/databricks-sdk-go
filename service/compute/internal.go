// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func addInstanceProfileToPb(st *AddInstanceProfile) (*addInstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addInstanceProfilePb{}
	pb.IamRoleArn = st.IamRoleArn
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.IsMetaInstanceProfile = st.IsMetaInstanceProfile
	pb.SkipValidation = st.SkipValidation

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type addInstanceProfilePb struct {
	IamRoleArn            string `json:"iam_role_arn,omitempty"`
	InstanceProfileArn    string `json:"instance_profile_arn"`
	IsMetaInstanceProfile bool   `json:"is_meta_instance_profile,omitempty"`
	SkipValidation        bool   `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func addInstanceProfileFromPb(pb *addInstanceProfilePb) (*AddInstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddInstanceProfile{}
	st.IamRoleArn = pb.IamRoleArn
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.IsMetaInstanceProfile = pb.IsMetaInstanceProfile
	st.SkipValidation = pb.SkipValidation

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *addInstanceProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st addInstanceProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func adlsgen2InfoToPb(st *Adlsgen2Info) (*adlsgen2InfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &adlsgen2InfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

type adlsgen2InfoPb struct {
	Destination string `json:"destination"`
}

func adlsgen2InfoFromPb(pb *adlsgen2InfoPb) (*Adlsgen2Info, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Adlsgen2Info{}
	st.Destination = pb.Destination

	return st, nil
}

func autoScaleToPb(st *AutoScale) (*autoScalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoScalePb{}
	pb.MaxWorkers = st.MaxWorkers
	pb.MinWorkers = st.MinWorkers

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type autoScalePb struct {
	MaxWorkers int `json:"max_workers,omitempty"`
	MinWorkers int `json:"min_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func autoScaleFromPb(pb *autoScalePb) (*AutoScale, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoScale{}
	st.MaxWorkers = pb.MaxWorkers
	st.MinWorkers = pb.MinWorkers

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *autoScalePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st autoScalePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func awsAttributesToPb(st *AwsAttributes) (*awsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &awsAttributesPb{}
	pb.Availability = st.Availability
	pb.EbsVolumeCount = st.EbsVolumeCount
	pb.EbsVolumeIops = st.EbsVolumeIops
	pb.EbsVolumeSize = st.EbsVolumeSize
	pb.EbsVolumeThroughput = st.EbsVolumeThroughput
	pb.EbsVolumeType = st.EbsVolumeType
	pb.FirstOnDemand = st.FirstOnDemand
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.SpotBidPricePercent = st.SpotBidPricePercent
	pb.ZoneId = st.ZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type awsAttributesPb struct {
	Availability        AwsAvailability `json:"availability,omitempty"`
	EbsVolumeCount      int             `json:"ebs_volume_count,omitempty"`
	EbsVolumeIops       int             `json:"ebs_volume_iops,omitempty"`
	EbsVolumeSize       int             `json:"ebs_volume_size,omitempty"`
	EbsVolumeThroughput int             `json:"ebs_volume_throughput,omitempty"`
	EbsVolumeType       EbsVolumeType   `json:"ebs_volume_type,omitempty"`
	FirstOnDemand       int             `json:"first_on_demand,omitempty"`
	InstanceProfileArn  string          `json:"instance_profile_arn,omitempty"`
	SpotBidPricePercent int             `json:"spot_bid_price_percent,omitempty"`
	ZoneId              string          `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func awsAttributesFromPb(pb *awsAttributesPb) (*AwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsAttributes{}
	st.Availability = pb.Availability
	st.EbsVolumeCount = pb.EbsVolumeCount
	st.EbsVolumeIops = pb.EbsVolumeIops
	st.EbsVolumeSize = pb.EbsVolumeSize
	st.EbsVolumeThroughput = pb.EbsVolumeThroughput
	st.EbsVolumeType = pb.EbsVolumeType
	st.FirstOnDemand = pb.FirstOnDemand
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.SpotBidPricePercent = pb.SpotBidPricePercent
	st.ZoneId = pb.ZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *awsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st awsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func azureAttributesToPb(st *AzureAttributes) (*azureAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &azureAttributesPb{}
	pb.Availability = st.Availability
	pb.FirstOnDemand = st.FirstOnDemand
	pb.LogAnalyticsInfo = st.LogAnalyticsInfo
	pb.SpotBidMaxPrice = st.SpotBidMaxPrice

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type azureAttributesPb struct {
	Availability     AzureAvailability `json:"availability,omitempty"`
	FirstOnDemand    int               `json:"first_on_demand,omitempty"`
	LogAnalyticsInfo *LogAnalyticsInfo `json:"log_analytics_info,omitempty"`
	SpotBidMaxPrice  float64           `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func azureAttributesFromPb(pb *azureAttributesPb) (*AzureAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureAttributes{}
	st.Availability = pb.Availability
	st.FirstOnDemand = pb.FirstOnDemand
	st.LogAnalyticsInfo = pb.LogAnalyticsInfo
	st.SpotBidMaxPrice = pb.SpotBidMaxPrice

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *azureAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st azureAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cancelCommandToPb(st *CancelCommand) (*cancelCommandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelCommandPb{}
	pb.ClusterId = st.ClusterId
	pb.CommandId = st.CommandId
	pb.ContextId = st.ContextId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cancelCommandPb struct {
	ClusterId string `json:"clusterId,omitempty"`
	CommandId string `json:"commandId,omitempty"`
	ContextId string `json:"contextId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cancelCommandFromPb(pb *cancelCommandPb) (*CancelCommand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelCommand{}
	st.ClusterId = pb.ClusterId
	st.CommandId = pb.CommandId
	st.ContextId = pb.ContextId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cancelCommandPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cancelCommandPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func changeClusterOwnerToPb(st *ChangeClusterOwner) (*changeClusterOwnerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &changeClusterOwnerPb{}
	pb.ClusterId = st.ClusterId
	pb.OwnerUsername = st.OwnerUsername

	return pb, nil
}

type changeClusterOwnerPb struct {
	ClusterId     string `json:"cluster_id"`
	OwnerUsername string `json:"owner_username"`
}

func changeClusterOwnerFromPb(pb *changeClusterOwnerPb) (*ChangeClusterOwner, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChangeClusterOwner{}
	st.ClusterId = pb.ClusterId
	st.OwnerUsername = pb.OwnerUsername

	return st, nil
}

func clientsTypesToPb(st *ClientsTypes) (*clientsTypesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clientsTypesPb{}
	pb.Jobs = st.Jobs
	pb.Notebooks = st.Notebooks

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clientsTypesPb struct {
	Jobs      bool `json:"jobs,omitempty"`
	Notebooks bool `json:"notebooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clientsTypesFromPb(pb *clientsTypesPb) (*ClientsTypes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClientsTypes{}
	st.Jobs = pb.Jobs
	st.Notebooks = pb.Notebooks

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clientsTypesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clientsTypesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cloneClusterToPb(st *CloneCluster) (*cloneClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloneClusterPb{}
	pb.SourceClusterId = st.SourceClusterId

	return pb, nil
}

type cloneClusterPb struct {
	SourceClusterId string `json:"source_cluster_id"`
}

func cloneClusterFromPb(pb *cloneClusterPb) (*CloneCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloneCluster{}
	st.SourceClusterId = pb.SourceClusterId

	return st, nil
}

func cloudProviderNodeInfoToPb(st *CloudProviderNodeInfo) (*cloudProviderNodeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudProviderNodeInfoPb{}
	pb.Status = st.Status

	return pb, nil
}

type cloudProviderNodeInfoPb struct {
	Status []CloudProviderNodeStatus `json:"status,omitempty"`
}

func cloudProviderNodeInfoFromPb(pb *cloudProviderNodeInfoPb) (*CloudProviderNodeInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudProviderNodeInfo{}
	st.Status = pb.Status

	return st, nil
}

func clusterAccessControlRequestToPb(st *ClusterAccessControlRequest) (*clusterAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterAccessControlRequestPb struct {
	GroupName            string                 `json:"group_name,omitempty"`
	PermissionLevel      ClusterPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                 `json:"service_principal_name,omitempty"`
	UserName             string                 `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAccessControlRequestFromPb(pb *clusterAccessControlRequestPb) (*ClusterAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterAccessControlResponseToPb(st *ClusterAccessControlResponse) (*clusterAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterAccessControlResponsePb struct {
	AllPermissions       []ClusterPermission `json:"all_permissions,omitempty"`
	DisplayName          string              `json:"display_name,omitempty"`
	GroupName            string              `json:"group_name,omitempty"`
	ServicePrincipalName string              `json:"service_principal_name,omitempty"`
	UserName             string              `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAccessControlResponseFromPb(pb *clusterAccessControlResponsePb) (*ClusterAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterAttributesToPb(st *ClusterAttributes) (*clusterAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAttributesPb{}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.ClusterLogConf = st.ClusterLogConf
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	pb.DataSecurityMode = st.DataSecurityMode
	pb.DockerImage = st.DockerImage
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.Kind = st.Kind
	pb.NodeTypeId = st.NodeTypeId
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.RuntimeEngine = st.RuntimeEngine
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterAttributesPb struct {
	AutoterminationMinutes     int               `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributes    `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributes  `json:"azure_attributes,omitempty"`
	ClusterLogConf             *ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	ClusterName                string            `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityMode  `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImage      `json:"docker_image,omitempty"`
	DriverInstancePoolId       string            `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string            `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool              `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool              `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributes    `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfo  `json:"init_scripts,omitempty"`
	InstancePoolId             string            `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool              `json:"is_single_node,omitempty"`
	Kind                       Kind              `json:"kind,omitempty"`
	NodeTypeId                 string            `json:"node_type_id,omitempty"`
	PolicyId                   string            `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int               `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEngine     `json:"runtime_engine,omitempty"`
	SingleUserName             string            `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string `json:"spark_env_vars,omitempty"`
	SparkVersion               string            `json:"spark_version"`
	SshPublicKeys              []string          `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int               `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool              `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadType     `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAttributesFromPb(pb *clusterAttributesPb) (*ClusterAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAttributes{}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.ClusterLogConf = pb.ClusterLogConf
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	st.DataSecurityMode = pb.DataSecurityMode
	st.DockerImage = pb.DockerImage
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.Kind = pb.Kind
	st.NodeTypeId = pb.NodeTypeId
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.RuntimeEngine = pb.RuntimeEngine
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterComplianceToPb(st *ClusterCompliance) (*clusterCompliancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterCompliancePb{}
	pb.ClusterId = st.ClusterId
	pb.IsCompliant = st.IsCompliant
	pb.Violations = st.Violations

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterCompliancePb struct {
	ClusterId   string            `json:"cluster_id"`
	IsCompliant bool              `json:"is_compliant,omitempty"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterComplianceFromPb(pb *clusterCompliancePb) (*ClusterCompliance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterCompliance{}
	st.ClusterId = pb.ClusterId
	st.IsCompliant = pb.IsCompliant
	st.Violations = pb.Violations

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterCompliancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterCompliancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterDetailsToPb(st *ClusterDetails) (*clusterDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterDetailsPb{}
	pb.Autoscale = st.Autoscale
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.ClusterCores = st.ClusterCores
	pb.ClusterId = st.ClusterId
	pb.ClusterLogConf = st.ClusterLogConf
	pb.ClusterLogStatus = st.ClusterLogStatus
	pb.ClusterMemoryMb = st.ClusterMemoryMb
	pb.ClusterName = st.ClusterName
	pb.ClusterSource = st.ClusterSource
	pb.CreatorUserName = st.CreatorUserName
	pb.CustomTags = st.CustomTags
	pb.DataSecurityMode = st.DataSecurityMode
	pb.DefaultTags = st.DefaultTags
	pb.DockerImage = st.DockerImage
	pb.Driver = st.Driver
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.Executors = st.Executors
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.JdbcPort = st.JdbcPort
	pb.Kind = st.Kind
	pb.LastRestartedTime = st.LastRestartedTime
	pb.LastStateLossTime = st.LastStateLossTime
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.RuntimeEngine = st.RuntimeEngine
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkContextId = st.SparkContextId
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.Spec = st.Spec
	pb.SshPublicKeys = st.SshPublicKeys
	pb.StartTime = st.StartTime
	pb.State = st.State
	pb.StateMessage = st.StateMessage
	pb.TerminatedTime = st.TerminatedTime
	pb.TerminationReason = st.TerminationReason
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterDetailsPb struct {
	Autoscale                  *AutoScale         `json:"autoscale,omitempty"`
	AutoterminationMinutes     int                `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributes     `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributes   `json:"azure_attributes,omitempty"`
	ClusterCores               float64            `json:"cluster_cores,omitempty"`
	ClusterId                  string             `json:"cluster_id,omitempty"`
	ClusterLogConf             *ClusterLogConf    `json:"cluster_log_conf,omitempty"`
	ClusterLogStatus           *LogSyncStatus     `json:"cluster_log_status,omitempty"`
	ClusterMemoryMb            int64              `json:"cluster_memory_mb,omitempty"`
	ClusterName                string             `json:"cluster_name,omitempty"`
	ClusterSource              ClusterSource      `json:"cluster_source,omitempty"`
	CreatorUserName            string             `json:"creator_user_name,omitempty"`
	CustomTags                 map[string]string  `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityMode   `json:"data_security_mode,omitempty"`
	DefaultTags                map[string]string  `json:"default_tags,omitempty"`
	DockerImage                *DockerImage       `json:"docker_image,omitempty"`
	Driver                     *SparkNode         `json:"driver,omitempty"`
	DriverInstancePoolId       string             `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string             `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool               `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool               `json:"enable_local_disk_encryption,omitempty"`
	Executors                  []SparkNode        `json:"executors,omitempty"`
	GcpAttributes              *GcpAttributes     `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfo   `json:"init_scripts,omitempty"`
	InstancePoolId             string             `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool               `json:"is_single_node,omitempty"`
	JdbcPort                   int                `json:"jdbc_port,omitempty"`
	Kind                       Kind               `json:"kind,omitempty"`
	LastRestartedTime          int64              `json:"last_restarted_time,omitempty"`
	LastStateLossTime          int64              `json:"last_state_loss_time,omitempty"`
	NodeTypeId                 string             `json:"node_type_id,omitempty"`
	NumWorkers                 int                `json:"num_workers,omitempty"`
	PolicyId                   string             `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int                `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEngine      `json:"runtime_engine,omitempty"`
	SingleUserName             string             `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string  `json:"spark_conf,omitempty"`
	SparkContextId             int64              `json:"spark_context_id,omitempty"`
	SparkEnvVars               map[string]string  `json:"spark_env_vars,omitempty"`
	SparkVersion               string             `json:"spark_version,omitempty"`
	Spec                       *ClusterSpec       `json:"spec,omitempty"`
	SshPublicKeys              []string           `json:"ssh_public_keys,omitempty"`
	StartTime                  int64              `json:"start_time,omitempty"`
	State                      State              `json:"state,omitempty"`
	StateMessage               string             `json:"state_message,omitempty"`
	TerminatedTime             int64              `json:"terminated_time,omitempty"`
	TerminationReason          *TerminationReason `json:"termination_reason,omitempty"`
	TotalInitialRemoteDiskSize int                `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool               `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadType      `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterDetailsFromPb(pb *clusterDetailsPb) (*ClusterDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterDetails{}
	st.Autoscale = pb.Autoscale
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.ClusterCores = pb.ClusterCores
	st.ClusterId = pb.ClusterId
	st.ClusterLogConf = pb.ClusterLogConf
	st.ClusterLogStatus = pb.ClusterLogStatus
	st.ClusterMemoryMb = pb.ClusterMemoryMb
	st.ClusterName = pb.ClusterName
	st.ClusterSource = pb.ClusterSource
	st.CreatorUserName = pb.CreatorUserName
	st.CustomTags = pb.CustomTags
	st.DataSecurityMode = pb.DataSecurityMode
	st.DefaultTags = pb.DefaultTags
	st.DockerImage = pb.DockerImage
	st.Driver = pb.Driver
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.Executors = pb.Executors
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.JdbcPort = pb.JdbcPort
	st.Kind = pb.Kind
	st.LastRestartedTime = pb.LastRestartedTime
	st.LastStateLossTime = pb.LastStateLossTime
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.RuntimeEngine = pb.RuntimeEngine
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkContextId = pb.SparkContextId
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.Spec = pb.Spec
	st.SshPublicKeys = pb.SshPublicKeys
	st.StartTime = pb.StartTime
	st.State = pb.State
	st.StateMessage = pb.StateMessage
	st.TerminatedTime = pb.TerminatedTime
	st.TerminationReason = pb.TerminationReason
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterEventToPb(st *ClusterEvent) (*clusterEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterEventPb{}
	pb.ClusterId = st.ClusterId
	pb.DataPlaneEventDetails = st.DataPlaneEventDetails
	pb.Details = st.Details
	pb.Timestamp = st.Timestamp
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterEventPb struct {
	ClusterId             string                 `json:"cluster_id"`
	DataPlaneEventDetails *DataPlaneEventDetails `json:"data_plane_event_details,omitempty"`
	Details               *EventDetails          `json:"details,omitempty"`
	Timestamp             int64                  `json:"timestamp,omitempty"`
	Type                  EventType              `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterEventFromPb(pb *clusterEventPb) (*ClusterEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterEvent{}
	st.ClusterId = pb.ClusterId
	st.DataPlaneEventDetails = pb.DataPlaneEventDetails
	st.Details = pb.Details
	st.Timestamp = pb.Timestamp
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterEventPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterEventPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterLibraryStatusesToPb(st *ClusterLibraryStatuses) (*clusterLibraryStatusesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterLibraryStatusesPb{}
	pb.ClusterId = st.ClusterId
	pb.LibraryStatuses = st.LibraryStatuses

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterLibraryStatusesPb struct {
	ClusterId       string              `json:"cluster_id,omitempty"`
	LibraryStatuses []LibraryFullStatus `json:"library_statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterLibraryStatusesFromPb(pb *clusterLibraryStatusesPb) (*ClusterLibraryStatuses, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterLibraryStatuses{}
	st.ClusterId = pb.ClusterId
	st.LibraryStatuses = pb.LibraryStatuses

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterLibraryStatusesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterLibraryStatusesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterLogConfToPb(st *ClusterLogConf) (*clusterLogConfPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterLogConfPb{}
	pb.Dbfs = st.Dbfs
	pb.S3 = st.S3
	pb.Volumes = st.Volumes

	return pb, nil
}

type clusterLogConfPb struct {
	Dbfs    *DbfsStorageInfo    `json:"dbfs,omitempty"`
	S3      *S3StorageInfo      `json:"s3,omitempty"`
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
}

func clusterLogConfFromPb(pb *clusterLogConfPb) (*ClusterLogConf, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterLogConf{}
	st.Dbfs = pb.Dbfs
	st.S3 = pb.S3
	st.Volumes = pb.Volumes

	return st, nil
}

func clusterPermissionToPb(st *ClusterPermission) (*clusterPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPermissionPb struct {
	Inherited           bool                   `json:"inherited,omitempty"`
	InheritedFromObject []string               `json:"inherited_from_object,omitempty"`
	PermissionLevel     ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPermissionFromPb(pb *clusterPermissionPb) (*ClusterPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPermissionsToPb(st *ClusterPermissions) (*clusterPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPermissionsPb struct {
	AccessControlList []ClusterAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                         `json:"object_id,omitempty"`
	ObjectType        string                         `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPermissionsFromPb(pb *clusterPermissionsPb) (*ClusterPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPermissionsDescriptionToPb(st *ClusterPermissionsDescription) (*clusterPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPermissionsDescriptionPb struct {
	Description     string                 `json:"description,omitempty"`
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPermissionsDescriptionFromPb(pb *clusterPermissionsDescriptionPb) (*ClusterPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPermissionsRequestToPb(st *ClusterPermissionsRequest) (*clusterPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type clusterPermissionsRequestPb struct {
	AccessControlList []ClusterAccessControlRequest `json:"access_control_list,omitempty"`
	ClusterId         string                        `json:"-" url:"-"`
}

func clusterPermissionsRequestFromPb(pb *clusterPermissionsRequestPb) (*ClusterPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.ClusterId = pb.ClusterId

	return st, nil
}

func clusterPolicyAccessControlRequestToPb(st *ClusterPolicyAccessControlRequest) (*clusterPolicyAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPolicyAccessControlRequestPb struct {
	GroupName            string                       `json:"group_name,omitempty"`
	PermissionLevel      ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                       `json:"service_principal_name,omitempty"`
	UserName             string                       `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyAccessControlRequestFromPb(pb *clusterPolicyAccessControlRequestPb) (*ClusterPolicyAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPolicyAccessControlResponseToPb(st *ClusterPolicyAccessControlResponse) (*clusterPolicyAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPolicyAccessControlResponsePb struct {
	AllPermissions       []ClusterPolicyPermission `json:"all_permissions,omitempty"`
	DisplayName          string                    `json:"display_name,omitempty"`
	GroupName            string                    `json:"group_name,omitempty"`
	ServicePrincipalName string                    `json:"service_principal_name,omitempty"`
	UserName             string                    `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyAccessControlResponseFromPb(pb *clusterPolicyAccessControlResponsePb) (*ClusterPolicyAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPolicyPermissionToPb(st *ClusterPolicyPermission) (*clusterPolicyPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPolicyPermissionPb struct {
	Inherited           bool                         `json:"inherited,omitempty"`
	InheritedFromObject []string                     `json:"inherited_from_object,omitempty"`
	PermissionLevel     ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyPermissionFromPb(pb *clusterPolicyPermissionPb) (*ClusterPolicyPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPolicyPermissionsToPb(st *ClusterPolicyPermissions) (*clusterPolicyPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPolicyPermissionsPb struct {
	AccessControlList []ClusterPolicyAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                               `json:"object_id,omitempty"`
	ObjectType        string                               `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyPermissionsFromPb(pb *clusterPolicyPermissionsPb) (*ClusterPolicyPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPolicyPermissionsDescriptionToPb(st *ClusterPolicyPermissionsDescription) (*clusterPolicyPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterPolicyPermissionsDescriptionPb struct {
	Description     string                       `json:"description,omitempty"`
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyPermissionsDescriptionFromPb(pb *clusterPolicyPermissionsDescriptionPb) (*ClusterPolicyPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterPolicyPermissionsRequestToPb(st *ClusterPolicyPermissionsRequest) (*clusterPolicyPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ClusterPolicyId = st.ClusterPolicyId

	return pb, nil
}

type clusterPolicyPermissionsRequestPb struct {
	AccessControlList []ClusterPolicyAccessControlRequest `json:"access_control_list,omitempty"`
	ClusterPolicyId   string                              `json:"-" url:"-"`
}

func clusterPolicyPermissionsRequestFromPb(pb *clusterPolicyPermissionsRequestPb) (*ClusterPolicyPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.ClusterPolicyId = pb.ClusterPolicyId

	return st, nil
}

func clusterSettingsChangeToPb(st *ClusterSettingsChange) (*clusterSettingsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSettingsChangePb{}
	pb.Field = st.Field
	pb.NewValue = st.NewValue
	pb.PreviousValue = st.PreviousValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterSettingsChangePb struct {
	Field         string `json:"field,omitempty"`
	NewValue      string `json:"new_value,omitempty"`
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterSettingsChangeFromPb(pb *clusterSettingsChangePb) (*ClusterSettingsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSettingsChange{}
	st.Field = pb.Field
	st.NewValue = pb.NewValue
	st.PreviousValue = pb.PreviousValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSettingsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSettingsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterSizeToPb(st *ClusterSize) (*clusterSizePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSizePb{}
	pb.Autoscale = st.Autoscale
	pb.NumWorkers = st.NumWorkers

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterSizePb struct {
	Autoscale  *AutoScale `json:"autoscale,omitempty"`
	NumWorkers int        `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterSizeFromPb(pb *clusterSizePb) (*ClusterSize, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSize{}
	st.Autoscale = pb.Autoscale
	st.NumWorkers = pb.NumWorkers

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSizePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSizePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterSpecToPb(st *ClusterSpec) (*clusterSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSpecPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	pb.Autoscale = st.Autoscale
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.ClusterLogConf = st.ClusterLogConf
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	pb.DataSecurityMode = st.DataSecurityMode
	pb.DockerImage = st.DockerImage
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.Kind = st.Kind
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.RuntimeEngine = st.RuntimeEngine
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterSpecPb struct {
	ApplyPolicyDefaultValues   bool              `json:"apply_policy_default_values,omitempty"`
	Autoscale                  *AutoScale        `json:"autoscale,omitempty"`
	AutoterminationMinutes     int               `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributes    `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributes  `json:"azure_attributes,omitempty"`
	ClusterLogConf             *ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	ClusterName                string            `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityMode  `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImage      `json:"docker_image,omitempty"`
	DriverInstancePoolId       string            `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string            `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool              `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool              `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributes    `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfo  `json:"init_scripts,omitempty"`
	InstancePoolId             string            `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool              `json:"is_single_node,omitempty"`
	Kind                       Kind              `json:"kind,omitempty"`
	NodeTypeId                 string            `json:"node_type_id,omitempty"`
	NumWorkers                 int               `json:"num_workers,omitempty"`
	PolicyId                   string            `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int               `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEngine     `json:"runtime_engine,omitempty"`
	SingleUserName             string            `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string `json:"spark_env_vars,omitempty"`
	SparkVersion               string            `json:"spark_version,omitempty"`
	SshPublicKeys              []string          `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int               `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool              `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadType     `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterSpecFromPb(pb *clusterSpecPb) (*ClusterSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSpec{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	st.Autoscale = pb.Autoscale
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.ClusterLogConf = pb.ClusterLogConf
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	st.DataSecurityMode = pb.DataSecurityMode
	st.DockerImage = pb.DockerImage
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.Kind = pb.Kind
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.RuntimeEngine = pb.RuntimeEngine
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterStatusToPb(st *ClusterStatus) (*clusterStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterStatusPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type clusterStatusPb struct {
	ClusterId string `json:"-" url:"cluster_id"`
}

func clusterStatusFromPb(pb *clusterStatusPb) (*ClusterStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterStatus{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func commandToPb(st *Command) (*commandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commandPb{}
	pb.ClusterId = st.ClusterId
	pb.Command = st.Command
	pb.ContextId = st.ContextId
	pb.Language = st.Language

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type commandPb struct {
	ClusterId string   `json:"clusterId,omitempty"`
	Command   string   `json:"command,omitempty"`
	ContextId string   `json:"contextId,omitempty"`
	Language  Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func commandFromPb(pb *commandPb) (*Command, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Command{}
	st.ClusterId = pb.ClusterId
	st.Command = pb.Command
	st.ContextId = pb.ContextId
	st.Language = pb.Language

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *commandPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st commandPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func commandStatusRequestToPb(st *CommandStatusRequest) (*commandStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commandStatusRequestPb{}
	pb.ClusterId = st.ClusterId
	pb.CommandId = st.CommandId
	pb.ContextId = st.ContextId

	return pb, nil
}

type commandStatusRequestPb struct {
	ClusterId string `json:"-" url:"clusterId"`
	CommandId string `json:"-" url:"commandId"`
	ContextId string `json:"-" url:"contextId"`
}

func commandStatusRequestFromPb(pb *commandStatusRequestPb) (*CommandStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommandStatusRequest{}
	st.ClusterId = pb.ClusterId
	st.CommandId = pb.CommandId
	st.ContextId = pb.ContextId

	return st, nil
}

func commandStatusResponseToPb(st *CommandStatusResponse) (*commandStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commandStatusResponsePb{}
	pb.Id = st.Id
	pb.Results = st.Results
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type commandStatusResponsePb struct {
	Id      string        `json:"id,omitempty"`
	Results *Results      `json:"results,omitempty"`
	Status  CommandStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func commandStatusResponseFromPb(pb *commandStatusResponsePb) (*CommandStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommandStatusResponse{}
	st.Id = pb.Id
	st.Results = pb.Results
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *commandStatusResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st commandStatusResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func contextStatusRequestToPb(st *ContextStatusRequest) (*contextStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &contextStatusRequestPb{}
	pb.ClusterId = st.ClusterId
	pb.ContextId = st.ContextId

	return pb, nil
}

type contextStatusRequestPb struct {
	ClusterId string `json:"-" url:"clusterId"`
	ContextId string `json:"-" url:"contextId"`
}

func contextStatusRequestFromPb(pb *contextStatusRequestPb) (*ContextStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContextStatusRequest{}
	st.ClusterId = pb.ClusterId
	st.ContextId = pb.ContextId

	return st, nil
}

func contextStatusResponseToPb(st *ContextStatusResponse) (*contextStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &contextStatusResponsePb{}
	pb.Id = st.Id
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type contextStatusResponsePb struct {
	Id     string        `json:"id,omitempty"`
	Status ContextStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func contextStatusResponseFromPb(pb *contextStatusResponsePb) (*ContextStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContextStatusResponse{}
	st.Id = pb.Id
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *contextStatusResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st contextStatusResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createClusterToPb(st *CreateCluster) (*createClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	pb.Autoscale = st.Autoscale
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.CloneFrom = st.CloneFrom
	pb.ClusterLogConf = st.ClusterLogConf
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	pb.DataSecurityMode = st.DataSecurityMode
	pb.DockerImage = st.DockerImage
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.Kind = st.Kind
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.RuntimeEngine = st.RuntimeEngine
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createClusterPb struct {
	ApplyPolicyDefaultValues   bool              `json:"apply_policy_default_values,omitempty"`
	Autoscale                  *AutoScale        `json:"autoscale,omitempty"`
	AutoterminationMinutes     int               `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributes    `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributes  `json:"azure_attributes,omitempty"`
	CloneFrom                  *CloneCluster     `json:"clone_from,omitempty"`
	ClusterLogConf             *ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	ClusterName                string            `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityMode  `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImage      `json:"docker_image,omitempty"`
	DriverInstancePoolId       string            `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string            `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool              `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool              `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributes    `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfo  `json:"init_scripts,omitempty"`
	InstancePoolId             string            `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool              `json:"is_single_node,omitempty"`
	Kind                       Kind              `json:"kind,omitempty"`
	NodeTypeId                 string            `json:"node_type_id,omitempty"`
	NumWorkers                 int               `json:"num_workers,omitempty"`
	PolicyId                   string            `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int               `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEngine     `json:"runtime_engine,omitempty"`
	SingleUserName             string            `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string `json:"spark_env_vars,omitempty"`
	SparkVersion               string            `json:"spark_version"`
	SshPublicKeys              []string          `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int               `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool              `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadType     `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createClusterFromPb(pb *createClusterPb) (*CreateCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	st.Autoscale = pb.Autoscale
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.CloneFrom = pb.CloneFrom
	st.ClusterLogConf = pb.ClusterLogConf
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	st.DataSecurityMode = pb.DataSecurityMode
	st.DockerImage = pb.DockerImage
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.Kind = pb.Kind
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.RuntimeEngine = pb.RuntimeEngine
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createClusterResponseToPb(st *CreateClusterResponse) (*createClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createClusterResponsePb{}
	pb.ClusterId = st.ClusterId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createClusterResponsePb struct {
	ClusterId string `json:"cluster_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createClusterResponseFromPb(pb *createClusterResponsePb) (*CreateClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateClusterResponse{}
	st.ClusterId = pb.ClusterId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createClusterResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createClusterResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createContextToPb(st *CreateContext) (*createContextPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createContextPb{}
	pb.ClusterId = st.ClusterId
	pb.Language = st.Language

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createContextPb struct {
	ClusterId string   `json:"clusterId,omitempty"`
	Language  Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createContextFromPb(pb *createContextPb) (*CreateContext, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateContext{}
	st.ClusterId = pb.ClusterId
	st.Language = pb.Language

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createContextPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createContextPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createInstancePoolToPb(st *CreateInstancePool) (*createInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createInstancePoolPb{}
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.CustomTags = st.CustomTags
	pb.DiskSpec = st.DiskSpec
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.GcpAttributes = st.GcpAttributes
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId
	pb.PreloadedDockerImages = st.PreloadedDockerImages
	pb.PreloadedSparkVersions = st.PreloadedSparkVersions
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createInstancePoolPb struct {
	AwsAttributes                      *InstancePoolAwsAttributes   `json:"aws_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	CustomTags                         map[string]string            `json:"custom_tags,omitempty"`
	DiskSpec                           *DiskSpec                    `json:"disk_spec,omitempty"`
	EnableElasticDisk                  bool                         `json:"enable_elastic_disk,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributes   `json:"gcp_attributes,omitempty"`
	IdleInstanceAutoterminationMinutes int                          `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolName                   string                       `json:"instance_pool_name"`
	MaxCapacity                        int                          `json:"max_capacity,omitempty"`
	MinIdleInstances                   int                          `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string                       `json:"node_type_id"`
	PreloadedDockerImages              []DockerImage                `json:"preloaded_docker_images,omitempty"`
	PreloadedSparkVersions             []string                     `json:"preloaded_spark_versions,omitempty"`
	RemoteDiskThroughput               int                          `json:"remote_disk_throughput,omitempty"`
	TotalInitialRemoteDiskSize         int                          `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createInstancePoolFromPb(pb *createInstancePoolPb) (*CreateInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstancePool{}
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.CustomTags = pb.CustomTags
	st.DiskSpec = pb.DiskSpec
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.GcpAttributes = pb.GcpAttributes
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId
	st.PreloadedDockerImages = pb.PreloadedDockerImages
	st.PreloadedSparkVersions = pb.PreloadedSparkVersions
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createInstancePoolResponseToPb(st *CreateInstancePoolResponse) (*createInstancePoolResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createInstancePoolResponsePb{}
	pb.InstancePoolId = st.InstancePoolId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createInstancePoolResponsePb struct {
	InstancePoolId string `json:"instance_pool_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createInstancePoolResponseFromPb(pb *createInstancePoolResponsePb) (*CreateInstancePoolResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstancePoolResponse{}
	st.InstancePoolId = pb.InstancePoolId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createInstancePoolResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createInstancePoolResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPolicyToPb(st *CreatePolicy) (*createPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPolicyPb{}
	pb.Definition = st.Definition
	pb.Description = st.Description
	pb.Libraries = st.Libraries
	pb.MaxClustersPerUser = st.MaxClustersPerUser
	pb.Name = st.Name
	pb.PolicyFamilyDefinitionOverrides = st.PolicyFamilyDefinitionOverrides
	pb.PolicyFamilyId = st.PolicyFamilyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPolicyPb struct {
	Definition                      string    `json:"definition,omitempty"`
	Description                     string    `json:"description,omitempty"`
	Libraries                       []Library `json:"libraries,omitempty"`
	MaxClustersPerUser              int64     `json:"max_clusters_per_user,omitempty"`
	Name                            string    `json:"name,omitempty"`
	PolicyFamilyDefinitionOverrides string    `json:"policy_family_definition_overrides,omitempty"`
	PolicyFamilyId                  string    `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPolicyFromPb(pb *createPolicyPb) (*CreatePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePolicy{}
	st.Definition = pb.Definition
	st.Description = pb.Description
	st.Libraries = pb.Libraries
	st.MaxClustersPerUser = pb.MaxClustersPerUser
	st.Name = pb.Name
	st.PolicyFamilyDefinitionOverrides = pb.PolicyFamilyDefinitionOverrides
	st.PolicyFamilyId = pb.PolicyFamilyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPolicyResponseToPb(st *CreatePolicyResponse) (*createPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPolicyResponsePb{}
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPolicyResponsePb struct {
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPolicyResponseFromPb(pb *createPolicyResponsePb) (*CreatePolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePolicyResponse{}
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}
	pb.ScriptId = st.ScriptId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createResponsePb struct {
	ScriptId string `json:"script_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	st.ScriptId = pb.ScriptId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createdToPb(st *Created) (*createdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createdPb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createdPb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createdFromPb(pb *createdPb) (*Created, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Created{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func customPolicyTagToPb(st *CustomPolicyTag) (*customPolicyTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customPolicyTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type customPolicyTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customPolicyTagFromPb(pb *customPolicyTagPb) (*CustomPolicyTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomPolicyTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customPolicyTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customPolicyTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dataPlaneEventDetailsToPb(st *DataPlaneEventDetails) (*dataPlaneEventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneEventDetailsPb{}
	pb.EventType = st.EventType
	pb.ExecutorFailures = st.ExecutorFailures
	pb.HostId = st.HostId
	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dataPlaneEventDetailsPb struct {
	EventType        DataPlaneEventDetailsEventType `json:"event_type,omitempty"`
	ExecutorFailures int                            `json:"executor_failures,omitempty"`
	HostId           string                         `json:"host_id,omitempty"`
	Timestamp        int64                          `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneEventDetailsFromPb(pb *dataPlaneEventDetailsPb) (*DataPlaneEventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneEventDetails{}
	st.EventType = pb.EventType
	st.ExecutorFailures = pb.ExecutorFailures
	st.HostId = pb.HostId
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataPlaneEventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataPlaneEventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dbfsStorageInfoToPb(st *DbfsStorageInfo) (*dbfsStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dbfsStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

type dbfsStorageInfoPb struct {
	Destination string `json:"destination"`
}

func dbfsStorageInfoFromPb(pb *dbfsStorageInfoPb) (*DbfsStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbfsStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

func deleteClusterToPb(st *DeleteCluster) (*deleteClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type deleteClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func deleteClusterFromPb(pb *deleteClusterPb) (*DeleteCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func deleteGlobalInitScriptRequestToPb(st *DeleteGlobalInitScriptRequest) (*deleteGlobalInitScriptRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteGlobalInitScriptRequestPb{}
	pb.ScriptId = st.ScriptId

	return pb, nil
}

type deleteGlobalInitScriptRequestPb struct {
	ScriptId string `json:"-" url:"-"`
}

func deleteGlobalInitScriptRequestFromPb(pb *deleteGlobalInitScriptRequestPb) (*DeleteGlobalInitScriptRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteGlobalInitScriptRequest{}
	st.ScriptId = pb.ScriptId

	return st, nil
}

func deleteInstancePoolToPb(st *DeleteInstancePool) (*deleteInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteInstancePoolPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

type deleteInstancePoolPb struct {
	InstancePoolId string `json:"instance_pool_id"`
}

func deleteInstancePoolFromPb(pb *deleteInstancePoolPb) (*DeleteInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstancePool{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

func deletePolicyToPb(st *DeletePolicy) (*deletePolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePolicyPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type deletePolicyPb struct {
	PolicyId string `json:"policy_id"`
}

func deletePolicyFromPb(pb *deletePolicyPb) (*DeletePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePolicy{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

func destroyContextToPb(st *DestroyContext) (*destroyContextPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &destroyContextPb{}
	pb.ClusterId = st.ClusterId
	pb.ContextId = st.ContextId

	return pb, nil
}

type destroyContextPb struct {
	ClusterId string `json:"clusterId"`
	ContextId string `json:"contextId"`
}

func destroyContextFromPb(pb *destroyContextPb) (*DestroyContext, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DestroyContext{}
	st.ClusterId = pb.ClusterId
	st.ContextId = pb.ContextId

	return st, nil
}

func diskSpecToPb(st *DiskSpec) (*diskSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &diskSpecPb{}
	pb.DiskCount = st.DiskCount
	pb.DiskIops = st.DiskIops
	pb.DiskSize = st.DiskSize
	pb.DiskThroughput = st.DiskThroughput
	pb.DiskType = st.DiskType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type diskSpecPb struct {
	DiskCount      int       `json:"disk_count,omitempty"`
	DiskIops       int       `json:"disk_iops,omitempty"`
	DiskSize       int       `json:"disk_size,omitempty"`
	DiskThroughput int       `json:"disk_throughput,omitempty"`
	DiskType       *DiskType `json:"disk_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func diskSpecFromPb(pb *diskSpecPb) (*DiskSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DiskSpec{}
	st.DiskCount = pb.DiskCount
	st.DiskIops = pb.DiskIops
	st.DiskSize = pb.DiskSize
	st.DiskThroughput = pb.DiskThroughput
	st.DiskType = pb.DiskType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *diskSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st diskSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func diskTypeToPb(st *DiskType) (*diskTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &diskTypePb{}
	pb.AzureDiskVolumeType = st.AzureDiskVolumeType
	pb.EbsVolumeType = st.EbsVolumeType

	return pb, nil
}

type diskTypePb struct {
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType `json:"azure_disk_volume_type,omitempty"`
	EbsVolumeType       DiskTypeEbsVolumeType       `json:"ebs_volume_type,omitempty"`
}

func diskTypeFromPb(pb *diskTypePb) (*DiskType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DiskType{}
	st.AzureDiskVolumeType = pb.AzureDiskVolumeType
	st.EbsVolumeType = pb.EbsVolumeType

	return st, nil
}

func dockerBasicAuthToPb(st *DockerBasicAuth) (*dockerBasicAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dockerBasicAuthPb{}
	pb.Password = st.Password
	pb.Username = st.Username

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dockerBasicAuthPb struct {
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dockerBasicAuthFromPb(pb *dockerBasicAuthPb) (*DockerBasicAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DockerBasicAuth{}
	st.Password = pb.Password
	st.Username = pb.Username

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dockerBasicAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dockerBasicAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dockerImageToPb(st *DockerImage) (*dockerImagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dockerImagePb{}
	pb.BasicAuth = st.BasicAuth
	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dockerImagePb struct {
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
	Url       string           `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dockerImageFromPb(pb *dockerImagePb) (*DockerImage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DockerImage{}
	st.BasicAuth = pb.BasicAuth
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dockerImagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dockerImagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func editClusterToPb(st *EditCluster) (*editClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	pb.Autoscale = st.Autoscale
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.ClusterId = st.ClusterId
	pb.ClusterLogConf = st.ClusterLogConf
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	pb.DataSecurityMode = st.DataSecurityMode
	pb.DockerImage = st.DockerImage
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.Kind = st.Kind
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.RuntimeEngine = st.RuntimeEngine
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type editClusterPb struct {
	ApplyPolicyDefaultValues   bool              `json:"apply_policy_default_values,omitempty"`
	Autoscale                  *AutoScale        `json:"autoscale,omitempty"`
	AutoterminationMinutes     int               `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributes    `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributes  `json:"azure_attributes,omitempty"`
	ClusterId                  string            `json:"cluster_id"`
	ClusterLogConf             *ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	ClusterName                string            `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityMode  `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImage      `json:"docker_image,omitempty"`
	DriverInstancePoolId       string            `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string            `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool              `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool              `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributes    `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfo  `json:"init_scripts,omitempty"`
	InstancePoolId             string            `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool              `json:"is_single_node,omitempty"`
	Kind                       Kind              `json:"kind,omitempty"`
	NodeTypeId                 string            `json:"node_type_id,omitempty"`
	NumWorkers                 int               `json:"num_workers,omitempty"`
	PolicyId                   string            `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int               `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEngine     `json:"runtime_engine,omitempty"`
	SingleUserName             string            `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string `json:"spark_env_vars,omitempty"`
	SparkVersion               string            `json:"spark_version"`
	SshPublicKeys              []string          `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int               `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool              `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadType     `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editClusterFromPb(pb *editClusterPb) (*EditCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	st.Autoscale = pb.Autoscale
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.ClusterId = pb.ClusterId
	st.ClusterLogConf = pb.ClusterLogConf
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	st.DataSecurityMode = pb.DataSecurityMode
	st.DockerImage = pb.DockerImage
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.Kind = pb.Kind
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.RuntimeEngine = pb.RuntimeEngine
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func editInstancePoolToPb(st *EditInstancePool) (*editInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editInstancePoolPb{}
	pb.CustomTags = st.CustomTags
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolId = st.InstancePoolId
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type editInstancePoolPb struct {
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

func editInstancePoolFromPb(pb *editInstancePoolPb) (*EditInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditInstancePool{}
	st.CustomTags = pb.CustomTags
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolId = pb.InstancePoolId
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func editPolicyToPb(st *EditPolicy) (*editPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPolicyPb{}
	pb.Definition = st.Definition
	pb.Description = st.Description
	pb.Libraries = st.Libraries
	pb.MaxClustersPerUser = st.MaxClustersPerUser
	pb.Name = st.Name
	pb.PolicyFamilyDefinitionOverrides = st.PolicyFamilyDefinitionOverrides
	pb.PolicyFamilyId = st.PolicyFamilyId
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type editPolicyPb struct {
	Definition                      string    `json:"definition,omitempty"`
	Description                     string    `json:"description,omitempty"`
	Libraries                       []Library `json:"libraries,omitempty"`
	MaxClustersPerUser              int64     `json:"max_clusters_per_user,omitempty"`
	Name                            string    `json:"name,omitempty"`
	PolicyFamilyDefinitionOverrides string    `json:"policy_family_definition_overrides,omitempty"`
	PolicyFamilyId                  string    `json:"policy_family_id,omitempty"`
	PolicyId                        string    `json:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editPolicyFromPb(pb *editPolicyPb) (*EditPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPolicy{}
	st.Definition = pb.Definition
	st.Description = pb.Description
	st.Libraries = pb.Libraries
	st.MaxClustersPerUser = pb.MaxClustersPerUser
	st.Name = pb.Name
	st.PolicyFamilyDefinitionOverrides = pb.PolicyFamilyDefinitionOverrides
	st.PolicyFamilyId = pb.PolicyFamilyId
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enforceClusterComplianceRequestToPb(st *EnforceClusterComplianceRequest) (*enforceClusterComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforceClusterComplianceRequestPb{}
	pb.ClusterId = st.ClusterId
	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enforceClusterComplianceRequestPb struct {
	ClusterId    string `json:"cluster_id"`
	ValidateOnly bool   `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforceClusterComplianceRequestFromPb(pb *enforceClusterComplianceRequestPb) (*EnforceClusterComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforceClusterComplianceRequest{}
	st.ClusterId = pb.ClusterId
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforceClusterComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforceClusterComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enforceClusterComplianceResponseToPb(st *EnforceClusterComplianceResponse) (*enforceClusterComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforceClusterComplianceResponsePb{}
	pb.Changes = st.Changes
	pb.HasChanges = st.HasChanges

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enforceClusterComplianceResponsePb struct {
	Changes    []ClusterSettingsChange `json:"changes,omitempty"`
	HasChanges bool                    `json:"has_changes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforceClusterComplianceResponseFromPb(pb *enforceClusterComplianceResponsePb) (*EnforceClusterComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforceClusterComplianceResponse{}
	st.Changes = pb.Changes
	st.HasChanges = pb.HasChanges

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforceClusterComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforceClusterComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func environmentToPb(st *Environment) (*environmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &environmentPb{}
	pb.Client = st.Client
	pb.Dependencies = st.Dependencies
	pb.EnvironmentVersion = st.EnvironmentVersion
	pb.JarDependencies = st.JarDependencies

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type environmentPb struct {
	Client             string   `json:"client,omitempty"`
	Dependencies       []string `json:"dependencies,omitempty"`
	EnvironmentVersion string   `json:"environment_version,omitempty"`
	JarDependencies    []string `json:"jar_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func environmentFromPb(pb *environmentPb) (*Environment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Environment{}
	st.Client = pb.Client
	st.Dependencies = pb.Dependencies
	st.EnvironmentVersion = pb.EnvironmentVersion
	st.JarDependencies = pb.JarDependencies

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *environmentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st environmentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func eventDetailsToPb(st *EventDetails) (*eventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &eventDetailsPb{}
	pb.Attributes = st.Attributes
	pb.Cause = st.Cause
	pb.ClusterSize = st.ClusterSize
	pb.CurrentNumVcpus = st.CurrentNumVcpus
	pb.CurrentNumWorkers = st.CurrentNumWorkers
	pb.DidNotExpandReason = st.DidNotExpandReason
	pb.DiskSize = st.DiskSize
	pb.DriverStateMessage = st.DriverStateMessage
	pb.EnableTerminationForNodeBlocklisted = st.EnableTerminationForNodeBlocklisted
	pb.FreeSpace = st.FreeSpace
	pb.InitScripts = st.InitScripts
	pb.InstanceId = st.InstanceId
	pb.JobRunName = st.JobRunName
	pb.PreviousAttributes = st.PreviousAttributes
	pb.PreviousClusterSize = st.PreviousClusterSize
	pb.PreviousDiskSize = st.PreviousDiskSize
	pb.Reason = st.Reason
	pb.TargetNumVcpus = st.TargetNumVcpus
	pb.TargetNumWorkers = st.TargetNumWorkers
	pb.User = st.User

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type eventDetailsPb struct {
	Attributes                          *ClusterAttributes      `json:"attributes,omitempty"`
	Cause                               EventDetailsCause       `json:"cause,omitempty"`
	ClusterSize                         *ClusterSize            `json:"cluster_size,omitempty"`
	CurrentNumVcpus                     int                     `json:"current_num_vcpus,omitempty"`
	CurrentNumWorkers                   int                     `json:"current_num_workers,omitempty"`
	DidNotExpandReason                  string                  `json:"did_not_expand_reason,omitempty"`
	DiskSize                            int64                   `json:"disk_size,omitempty"`
	DriverStateMessage                  string                  `json:"driver_state_message,omitempty"`
	EnableTerminationForNodeBlocklisted bool                    `json:"enable_termination_for_node_blocklisted,omitempty"`
	FreeSpace                           int64                   `json:"free_space,omitempty"`
	InitScripts                         *InitScriptEventDetails `json:"init_scripts,omitempty"`
	InstanceId                          string                  `json:"instance_id,omitempty"`
	JobRunName                          string                  `json:"job_run_name,omitempty"`
	PreviousAttributes                  *ClusterAttributes      `json:"previous_attributes,omitempty"`
	PreviousClusterSize                 *ClusterSize            `json:"previous_cluster_size,omitempty"`
	PreviousDiskSize                    int64                   `json:"previous_disk_size,omitempty"`
	Reason                              *TerminationReason      `json:"reason,omitempty"`
	TargetNumVcpus                      int                     `json:"target_num_vcpus,omitempty"`
	TargetNumWorkers                    int                     `json:"target_num_workers,omitempty"`
	User                                string                  `json:"user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func eventDetailsFromPb(pb *eventDetailsPb) (*EventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EventDetails{}
	st.Attributes = pb.Attributes
	st.Cause = pb.Cause
	st.ClusterSize = pb.ClusterSize
	st.CurrentNumVcpus = pb.CurrentNumVcpus
	st.CurrentNumWorkers = pb.CurrentNumWorkers
	st.DidNotExpandReason = pb.DidNotExpandReason
	st.DiskSize = pb.DiskSize
	st.DriverStateMessage = pb.DriverStateMessage
	st.EnableTerminationForNodeBlocklisted = pb.EnableTerminationForNodeBlocklisted
	st.FreeSpace = pb.FreeSpace
	st.InitScripts = pb.InitScripts
	st.InstanceId = pb.InstanceId
	st.JobRunName = pb.JobRunName
	st.PreviousAttributes = pb.PreviousAttributes
	st.PreviousClusterSize = pb.PreviousClusterSize
	st.PreviousDiskSize = pb.PreviousDiskSize
	st.Reason = pb.Reason
	st.TargetNumVcpus = pb.TargetNumVcpus
	st.TargetNumWorkers = pb.TargetNumWorkers
	st.User = pb.User

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *eventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st eventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gcpAttributesToPb(st *GcpAttributes) (*gcpAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcpAttributesPb{}
	pb.Availability = st.Availability
	pb.BootDiskSize = st.BootDiskSize
	pb.GoogleServiceAccount = st.GoogleServiceAccount
	pb.LocalSsdCount = st.LocalSsdCount
	pb.UsePreemptibleExecutors = st.UsePreemptibleExecutors
	pb.ZoneId = st.ZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type gcpAttributesPb struct {
	Availability            GcpAvailability `json:"availability,omitempty"`
	BootDiskSize            int             `json:"boot_disk_size,omitempty"`
	GoogleServiceAccount    string          `json:"google_service_account,omitempty"`
	LocalSsdCount           int             `json:"local_ssd_count,omitempty"`
	UsePreemptibleExecutors bool            `json:"use_preemptible_executors,omitempty"`
	ZoneId                  string          `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gcpAttributesFromPb(pb *gcpAttributesPb) (*GcpAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpAttributes{}
	st.Availability = pb.Availability
	st.BootDiskSize = pb.BootDiskSize
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.LocalSsdCount = pb.LocalSsdCount
	st.UsePreemptibleExecutors = pb.UsePreemptibleExecutors
	st.ZoneId = pb.ZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gcpAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gcpAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func gcsStorageInfoToPb(st *GcsStorageInfo) (*gcsStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcsStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

type gcsStorageInfoPb struct {
	Destination string `json:"destination"`
}

func gcsStorageInfoFromPb(pb *gcsStorageInfoPb) (*GcsStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcsStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

func getClusterComplianceRequestToPb(st *GetClusterComplianceRequest) (*getClusterComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterComplianceRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type getClusterComplianceRequestPb struct {
	ClusterId string `json:"-" url:"cluster_id"`
}

func getClusterComplianceRequestFromPb(pb *getClusterComplianceRequestPb) (*GetClusterComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterComplianceRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func getClusterComplianceResponseToPb(st *GetClusterComplianceResponse) (*getClusterComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterComplianceResponsePb{}
	pb.IsCompliant = st.IsCompliant
	pb.Violations = st.Violations

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getClusterComplianceResponsePb struct {
	IsCompliant bool              `json:"is_compliant,omitempty"`
	Violations  map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getClusterComplianceResponseFromPb(pb *getClusterComplianceResponsePb) (*GetClusterComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterComplianceResponse{}
	st.IsCompliant = pb.IsCompliant
	st.Violations = pb.Violations

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getClusterComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getClusterComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getClusterPermissionLevelsRequestToPb(st *GetClusterPermissionLevelsRequest) (*getClusterPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPermissionLevelsRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type getClusterPermissionLevelsRequestPb struct {
	ClusterId string `json:"-" url:"-"`
}

func getClusterPermissionLevelsRequestFromPb(pb *getClusterPermissionLevelsRequestPb) (*GetClusterPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionLevelsRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func getClusterPermissionLevelsResponseToPb(st *GetClusterPermissionLevelsResponse) (*getClusterPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getClusterPermissionLevelsResponsePb struct {
	PermissionLevels []ClusterPermissionsDescription `json:"permission_levels,omitempty"`
}

func getClusterPermissionLevelsResponseFromPb(pb *getClusterPermissionLevelsResponsePb) (*GetClusterPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getClusterPermissionsRequestToPb(st *GetClusterPermissionsRequest) (*getClusterPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPermissionsRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type getClusterPermissionsRequestPb struct {
	ClusterId string `json:"-" url:"-"`
}

func getClusterPermissionsRequestFromPb(pb *getClusterPermissionsRequestPb) (*GetClusterPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionsRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func getClusterPolicyPermissionLevelsRequestToPb(st *GetClusterPolicyPermissionLevelsRequest) (*getClusterPolicyPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyPermissionLevelsRequestPb{}
	pb.ClusterPolicyId = st.ClusterPolicyId

	return pb, nil
}

type getClusterPolicyPermissionLevelsRequestPb struct {
	ClusterPolicyId string `json:"-" url:"-"`
}

func getClusterPolicyPermissionLevelsRequestFromPb(pb *getClusterPolicyPermissionLevelsRequestPb) (*GetClusterPolicyPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionLevelsRequest{}
	st.ClusterPolicyId = pb.ClusterPolicyId

	return st, nil
}

func getClusterPolicyPermissionLevelsResponseToPb(st *GetClusterPolicyPermissionLevelsResponse) (*getClusterPolicyPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getClusterPolicyPermissionLevelsResponsePb struct {
	PermissionLevels []ClusterPolicyPermissionsDescription `json:"permission_levels,omitempty"`
}

func getClusterPolicyPermissionLevelsResponseFromPb(pb *getClusterPolicyPermissionLevelsResponsePb) (*GetClusterPolicyPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getClusterPolicyPermissionsRequestToPb(st *GetClusterPolicyPermissionsRequest) (*getClusterPolicyPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyPermissionsRequestPb{}
	pb.ClusterPolicyId = st.ClusterPolicyId

	return pb, nil
}

type getClusterPolicyPermissionsRequestPb struct {
	ClusterPolicyId string `json:"-" url:"-"`
}

func getClusterPolicyPermissionsRequestFromPb(pb *getClusterPolicyPermissionsRequestPb) (*GetClusterPolicyPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionsRequest{}
	st.ClusterPolicyId = pb.ClusterPolicyId

	return st, nil
}

func getClusterPolicyRequestToPb(st *GetClusterPolicyRequest) (*getClusterPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type getClusterPolicyRequestPb struct {
	PolicyId string `json:"-" url:"policy_id"`
}

func getClusterPolicyRequestFromPb(pb *getClusterPolicyRequestPb) (*GetClusterPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

func getClusterRequestToPb(st *GetClusterRequest) (*getClusterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type getClusterRequestPb struct {
	ClusterId string `json:"-" url:"cluster_id"`
}

func getClusterRequestFromPb(pb *getClusterRequestPb) (*GetClusterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func getEventsToPb(st *GetEvents) (*getEventsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEventsPb{}
	pb.ClusterId = st.ClusterId
	pb.EndTime = st.EndTime
	pb.EventTypes = st.EventTypes
	pb.Limit = st.Limit
	pb.Offset = st.Offset
	pb.Order = st.Order
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.StartTime = st.StartTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getEventsPb struct {
	ClusterId  string         `json:"cluster_id"`
	EndTime    int64          `json:"end_time,omitempty"`
	EventTypes []EventType    `json:"event_types,omitempty"`
	Limit      int64          `json:"limit,omitempty"`
	Offset     int64          `json:"offset,omitempty"`
	Order      GetEventsOrder `json:"order,omitempty"`
	PageSize   int            `json:"page_size,omitempty"`
	PageToken  string         `json:"page_token,omitempty"`
	StartTime  int64          `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getEventsFromPb(pb *getEventsPb) (*GetEvents, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEvents{}
	st.ClusterId = pb.ClusterId
	st.EndTime = pb.EndTime
	st.EventTypes = pb.EventTypes
	st.Limit = pb.Limit
	st.Offset = pb.Offset
	st.Order = pb.Order
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.StartTime = pb.StartTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEventsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEventsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getEventsResponseToPb(st *GetEventsResponse) (*getEventsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEventsResponsePb{}
	pb.Events = st.Events
	pb.NextPage = st.NextPage
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken
	pb.TotalCount = st.TotalCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getEventsResponsePb struct {
	Events        []ClusterEvent `json:"events,omitempty"`
	NextPage      *GetEvents     `json:"next_page,omitempty"`
	NextPageToken string         `json:"next_page_token,omitempty"`
	PrevPageToken string         `json:"prev_page_token,omitempty"`
	TotalCount    int64          `json:"total_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getEventsResponseFromPb(pb *getEventsResponsePb) (*GetEventsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEventsResponse{}
	st.Events = pb.Events
	st.NextPage = pb.NextPage
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken
	st.TotalCount = pb.TotalCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEventsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEventsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getGlobalInitScriptRequestToPb(st *GetGlobalInitScriptRequest) (*getGlobalInitScriptRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGlobalInitScriptRequestPb{}
	pb.ScriptId = st.ScriptId

	return pb, nil
}

type getGlobalInitScriptRequestPb struct {
	ScriptId string `json:"-" url:"-"`
}

func getGlobalInitScriptRequestFromPb(pb *getGlobalInitScriptRequestPb) (*GetGlobalInitScriptRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGlobalInitScriptRequest{}
	st.ScriptId = pb.ScriptId

	return st, nil
}

func getInstancePoolToPb(st *GetInstancePool) (*getInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPb{}
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.CustomTags = st.CustomTags
	pb.DefaultTags = st.DefaultTags
	pb.DiskSpec = st.DiskSpec
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.GcpAttributes = st.GcpAttributes
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolId = st.InstancePoolId
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId
	pb.PreloadedDockerImages = st.PreloadedDockerImages
	pb.PreloadedSparkVersions = st.PreloadedSparkVersions
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.State = st.State
	pb.Stats = st.Stats
	pb.Status = st.Status
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getInstancePoolPb struct {
	AwsAttributes                      *InstancePoolAwsAttributes   `json:"aws_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	CustomTags                         map[string]string            `json:"custom_tags,omitempty"`
	DefaultTags                        map[string]string            `json:"default_tags,omitempty"`
	DiskSpec                           *DiskSpec                    `json:"disk_spec,omitempty"`
	EnableElasticDisk                  bool                         `json:"enable_elastic_disk,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributes   `json:"gcp_attributes,omitempty"`
	IdleInstanceAutoterminationMinutes int                          `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolId                     string                       `json:"instance_pool_id"`
	InstancePoolName                   string                       `json:"instance_pool_name,omitempty"`
	MaxCapacity                        int                          `json:"max_capacity,omitempty"`
	MinIdleInstances                   int                          `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string                       `json:"node_type_id,omitempty"`
	PreloadedDockerImages              []DockerImage                `json:"preloaded_docker_images,omitempty"`
	PreloadedSparkVersions             []string                     `json:"preloaded_spark_versions,omitempty"`
	RemoteDiskThroughput               int                          `json:"remote_disk_throughput,omitempty"`
	State                              InstancePoolState            `json:"state,omitempty"`
	Stats                              *InstancePoolStats           `json:"stats,omitempty"`
	Status                             *InstancePoolStatus          `json:"status,omitempty"`
	TotalInitialRemoteDiskSize         int                          `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getInstancePoolFromPb(pb *getInstancePoolPb) (*GetInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePool{}
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.CustomTags = pb.CustomTags
	st.DefaultTags = pb.DefaultTags
	st.DiskSpec = pb.DiskSpec
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.GcpAttributes = pb.GcpAttributes
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolId = pb.InstancePoolId
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId
	st.PreloadedDockerImages = pb.PreloadedDockerImages
	st.PreloadedSparkVersions = pb.PreloadedSparkVersions
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.State = pb.State
	st.Stats = pb.Stats
	st.Status = pb.Status
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getInstancePoolPermissionLevelsRequestToPb(st *GetInstancePoolPermissionLevelsRequest) (*getInstancePoolPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPermissionLevelsRequestPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

type getInstancePoolPermissionLevelsRequestPb struct {
	InstancePoolId string `json:"-" url:"-"`
}

func getInstancePoolPermissionLevelsRequestFromPb(pb *getInstancePoolPermissionLevelsRequestPb) (*GetInstancePoolPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionLevelsRequest{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

func getInstancePoolPermissionLevelsResponseToPb(st *GetInstancePoolPermissionLevelsResponse) (*getInstancePoolPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getInstancePoolPermissionLevelsResponsePb struct {
	PermissionLevels []InstancePoolPermissionsDescription `json:"permission_levels,omitempty"`
}

func getInstancePoolPermissionLevelsResponseFromPb(pb *getInstancePoolPermissionLevelsResponsePb) (*GetInstancePoolPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getInstancePoolPermissionsRequestToPb(st *GetInstancePoolPermissionsRequest) (*getInstancePoolPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPermissionsRequestPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

type getInstancePoolPermissionsRequestPb struct {
	InstancePoolId string `json:"-" url:"-"`
}

func getInstancePoolPermissionsRequestFromPb(pb *getInstancePoolPermissionsRequestPb) (*GetInstancePoolPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionsRequest{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

func getInstancePoolRequestToPb(st *GetInstancePoolRequest) (*getInstancePoolRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolRequestPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

type getInstancePoolRequestPb struct {
	InstancePoolId string `json:"-" url:"instance_pool_id"`
}

func getInstancePoolRequestFromPb(pb *getInstancePoolRequestPb) (*GetInstancePoolRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolRequest{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

func getPolicyFamilyRequestToPb(st *GetPolicyFamilyRequest) (*getPolicyFamilyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPolicyFamilyRequestPb{}
	pb.PolicyFamilyId = st.PolicyFamilyId
	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPolicyFamilyRequestPb struct {
	PolicyFamilyId string `json:"-" url:"-"`
	Version        int64  `json:"-" url:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPolicyFamilyRequestFromPb(pb *getPolicyFamilyRequestPb) (*GetPolicyFamilyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyFamilyRequest{}
	st.PolicyFamilyId = pb.PolicyFamilyId
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPolicyFamilyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPolicyFamilyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getSparkVersionsResponseToPb(st *GetSparkVersionsResponse) (*getSparkVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSparkVersionsResponsePb{}
	pb.Versions = st.Versions

	return pb, nil
}

type getSparkVersionsResponsePb struct {
	Versions []SparkVersion `json:"versions,omitempty"`
}

func getSparkVersionsResponseFromPb(pb *getSparkVersionsResponsePb) (*GetSparkVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSparkVersionsResponse{}
	st.Versions = pb.Versions

	return st, nil
}

func globalInitScriptCreateRequestToPb(st *GlobalInitScriptCreateRequest) (*globalInitScriptCreateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptCreateRequestPb{}
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.Script = st.Script

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type globalInitScriptCreateRequestPb struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Name     string `json:"name"`
	Position int    `json:"position,omitempty"`
	Script   string `json:"script"`

	ForceSendFields []string `json:"-" url:"-"`
}

func globalInitScriptCreateRequestFromPb(pb *globalInitScriptCreateRequestPb) (*GlobalInitScriptCreateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptCreateRequest{}
	st.Enabled = pb.Enabled
	st.Name = pb.Name
	st.Position = pb.Position
	st.Script = pb.Script

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptCreateRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptCreateRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func globalInitScriptDetailsToPb(st *GlobalInitScriptDetails) (*globalInitScriptDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptDetailsPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.ScriptId = st.ScriptId
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type globalInitScriptDetailsPb struct {
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

func globalInitScriptDetailsFromPb(pb *globalInitScriptDetailsPb) (*GlobalInitScriptDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptDetails{}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Enabled = pb.Enabled
	st.Name = pb.Name
	st.Position = pb.Position
	st.ScriptId = pb.ScriptId
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func globalInitScriptDetailsWithContentToPb(st *GlobalInitScriptDetailsWithContent) (*globalInitScriptDetailsWithContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptDetailsWithContentPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.Script = st.Script
	pb.ScriptId = st.ScriptId
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type globalInitScriptDetailsWithContentPb struct {
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

func globalInitScriptDetailsWithContentFromPb(pb *globalInitScriptDetailsWithContentPb) (*GlobalInitScriptDetailsWithContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptDetailsWithContent{}
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Enabled = pb.Enabled
	st.Name = pb.Name
	st.Position = pb.Position
	st.Script = pb.Script
	st.ScriptId = pb.ScriptId
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptDetailsWithContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptDetailsWithContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func globalInitScriptUpdateRequestToPb(st *GlobalInitScriptUpdateRequest) (*globalInitScriptUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptUpdateRequestPb{}
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.Script = st.Script
	pb.ScriptId = st.ScriptId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type globalInitScriptUpdateRequestPb struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Name     string `json:"name"`
	Position int    `json:"position,omitempty"`
	Script   string `json:"script"`
	ScriptId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func globalInitScriptUpdateRequestFromPb(pb *globalInitScriptUpdateRequestPb) (*GlobalInitScriptUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptUpdateRequest{}
	st.Enabled = pb.Enabled
	st.Name = pb.Name
	st.Position = pb.Position
	st.Script = pb.Script
	st.ScriptId = pb.ScriptId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptUpdateRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptUpdateRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func initScriptEventDetailsToPb(st *InitScriptEventDetails) (*initScriptEventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &initScriptEventDetailsPb{}
	pb.Cluster = st.Cluster
	pb.Global = st.Global
	pb.ReportedForNode = st.ReportedForNode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type initScriptEventDetailsPb struct {
	Cluster         []InitScriptInfoAndExecutionDetails `json:"cluster,omitempty"`
	Global          []InitScriptInfoAndExecutionDetails `json:"global,omitempty"`
	ReportedForNode string                              `json:"reported_for_node,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func initScriptEventDetailsFromPb(pb *initScriptEventDetailsPb) (*InitScriptEventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptEventDetails{}
	st.Cluster = pb.Cluster
	st.Global = pb.Global
	st.ReportedForNode = pb.ReportedForNode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *initScriptEventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st initScriptEventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func initScriptInfoToPb(st *InitScriptInfo) (*initScriptInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &initScriptInfoPb{}
	pb.Abfss = st.Abfss
	pb.Dbfs = st.Dbfs
	pb.File = st.File
	pb.Gcs = st.Gcs
	pb.S3 = st.S3
	pb.Volumes = st.Volumes
	pb.Workspace = st.Workspace

	return pb, nil
}

type initScriptInfoPb struct {
	Abfss     *Adlsgen2Info         `json:"abfss,omitempty"`
	Dbfs      *DbfsStorageInfo      `json:"dbfs,omitempty"`
	File      *LocalFileInfo        `json:"file,omitempty"`
	Gcs       *GcsStorageInfo       `json:"gcs,omitempty"`
	S3        *S3StorageInfo        `json:"s3,omitempty"`
	Volumes   *VolumesStorageInfo   `json:"volumes,omitempty"`
	Workspace *WorkspaceStorageInfo `json:"workspace,omitempty"`
}

func initScriptInfoFromPb(pb *initScriptInfoPb) (*InitScriptInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptInfo{}
	st.Abfss = pb.Abfss
	st.Dbfs = pb.Dbfs
	st.File = pb.File
	st.Gcs = pb.Gcs
	st.S3 = pb.S3
	st.Volumes = pb.Volumes
	st.Workspace = pb.Workspace

	return st, nil
}

func initScriptInfoAndExecutionDetailsToPb(st *InitScriptInfoAndExecutionDetails) (*initScriptInfoAndExecutionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &initScriptInfoAndExecutionDetailsPb{}
	pb.Abfss = st.Abfss
	pb.Dbfs = st.Dbfs
	pb.ErrorMessage = st.ErrorMessage
	pb.ExecutionDurationSeconds = st.ExecutionDurationSeconds
	pb.File = st.File
	pb.Gcs = st.Gcs
	pb.S3 = st.S3
	pb.Status = st.Status
	pb.Volumes = st.Volumes
	pb.Workspace = st.Workspace

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type initScriptInfoAndExecutionDetailsPb struct {
	Abfss                    *Adlsgen2Info                                       `json:"abfss,omitempty"`
	Dbfs                     *DbfsStorageInfo                                    `json:"dbfs,omitempty"`
	ErrorMessage             string                                              `json:"error_message,omitempty"`
	ExecutionDurationSeconds int                                                 `json:"execution_duration_seconds,omitempty"`
	File                     *LocalFileInfo                                      `json:"file,omitempty"`
	Gcs                      *GcsStorageInfo                                     `json:"gcs,omitempty"`
	S3                       *S3StorageInfo                                      `json:"s3,omitempty"`
	Status                   InitScriptExecutionDetailsInitScriptExecutionStatus `json:"status,omitempty"`
	Volumes                  *VolumesStorageInfo                                 `json:"volumes,omitempty"`
	Workspace                *WorkspaceStorageInfo                               `json:"workspace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func initScriptInfoAndExecutionDetailsFromPb(pb *initScriptInfoAndExecutionDetailsPb) (*InitScriptInfoAndExecutionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptInfoAndExecutionDetails{}
	st.Abfss = pb.Abfss
	st.Dbfs = pb.Dbfs
	st.ErrorMessage = pb.ErrorMessage
	st.ExecutionDurationSeconds = pb.ExecutionDurationSeconds
	st.File = pb.File
	st.Gcs = pb.Gcs
	st.S3 = pb.S3
	st.Status = pb.Status
	st.Volumes = pb.Volumes
	st.Workspace = pb.Workspace

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *initScriptInfoAndExecutionDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st initScriptInfoAndExecutionDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func installLibrariesToPb(st *InstallLibraries) (*installLibrariesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &installLibrariesPb{}
	pb.ClusterId = st.ClusterId
	pb.Libraries = st.Libraries

	return pb, nil
}

type installLibrariesPb struct {
	ClusterId string    `json:"cluster_id"`
	Libraries []Library `json:"libraries"`
}

func installLibrariesFromPb(pb *installLibrariesPb) (*InstallLibraries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstallLibraries{}
	st.ClusterId = pb.ClusterId
	st.Libraries = pb.Libraries

	return st, nil
}

func instancePoolAccessControlRequestToPb(st *InstancePoolAccessControlRequest) (*instancePoolAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolAccessControlRequestPb struct {
	GroupName            string                      `json:"group_name,omitempty"`
	PermissionLevel      InstancePoolPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                      `json:"service_principal_name,omitempty"`
	UserName             string                      `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAccessControlRequestFromPb(pb *instancePoolAccessControlRequestPb) (*InstancePoolAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolAccessControlResponseToPb(st *InstancePoolAccessControlResponse) (*instancePoolAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolAccessControlResponsePb struct {
	AllPermissions       []InstancePoolPermission `json:"all_permissions,omitempty"`
	DisplayName          string                   `json:"display_name,omitempty"`
	GroupName            string                   `json:"group_name,omitempty"`
	ServicePrincipalName string                   `json:"service_principal_name,omitempty"`
	UserName             string                   `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAccessControlResponseFromPb(pb *instancePoolAccessControlResponsePb) (*InstancePoolAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolAndStatsToPb(st *InstancePoolAndStats) (*instancePoolAndStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAndStatsPb{}
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.CustomTags = st.CustomTags
	pb.DefaultTags = st.DefaultTags
	pb.DiskSpec = st.DiskSpec
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.GcpAttributes = st.GcpAttributes
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolId = st.InstancePoolId
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId
	pb.PreloadedDockerImages = st.PreloadedDockerImages
	pb.PreloadedSparkVersions = st.PreloadedSparkVersions
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.State = st.State
	pb.Stats = st.Stats
	pb.Status = st.Status
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolAndStatsPb struct {
	AwsAttributes                      *InstancePoolAwsAttributes   `json:"aws_attributes,omitempty"`
	AzureAttributes                    *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	CustomTags                         map[string]string            `json:"custom_tags,omitempty"`
	DefaultTags                        map[string]string            `json:"default_tags,omitempty"`
	DiskSpec                           *DiskSpec                    `json:"disk_spec,omitempty"`
	EnableElasticDisk                  bool                         `json:"enable_elastic_disk,omitempty"`
	GcpAttributes                      *InstancePoolGcpAttributes   `json:"gcp_attributes,omitempty"`
	IdleInstanceAutoterminationMinutes int                          `json:"idle_instance_autotermination_minutes,omitempty"`
	InstancePoolId                     string                       `json:"instance_pool_id,omitempty"`
	InstancePoolName                   string                       `json:"instance_pool_name,omitempty"`
	MaxCapacity                        int                          `json:"max_capacity,omitempty"`
	MinIdleInstances                   int                          `json:"min_idle_instances,omitempty"`
	NodeTypeId                         string                       `json:"node_type_id,omitempty"`
	PreloadedDockerImages              []DockerImage                `json:"preloaded_docker_images,omitempty"`
	PreloadedSparkVersions             []string                     `json:"preloaded_spark_versions,omitempty"`
	RemoteDiskThroughput               int                          `json:"remote_disk_throughput,omitempty"`
	State                              InstancePoolState            `json:"state,omitempty"`
	Stats                              *InstancePoolStats           `json:"stats,omitempty"`
	Status                             *InstancePoolStatus          `json:"status,omitempty"`
	TotalInitialRemoteDiskSize         int                          `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAndStatsFromPb(pb *instancePoolAndStatsPb) (*InstancePoolAndStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAndStats{}
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.CustomTags = pb.CustomTags
	st.DefaultTags = pb.DefaultTags
	st.DiskSpec = pb.DiskSpec
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.GcpAttributes = pb.GcpAttributes
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolId = pb.InstancePoolId
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId
	st.PreloadedDockerImages = pb.PreloadedDockerImages
	st.PreloadedSparkVersions = pb.PreloadedSparkVersions
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.State = pb.State
	st.Stats = pb.Stats
	st.Status = pb.Status
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAndStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAndStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolAwsAttributesToPb(st *InstancePoolAwsAttributes) (*instancePoolAwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAwsAttributesPb{}
	pb.Availability = st.Availability
	pb.SpotBidPricePercent = st.SpotBidPricePercent
	pb.ZoneId = st.ZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolAwsAttributesPb struct {
	Availability        InstancePoolAwsAttributesAvailability `json:"availability,omitempty"`
	SpotBidPricePercent int                                   `json:"spot_bid_price_percent,omitempty"`
	ZoneId              string                                `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAwsAttributesFromPb(pb *instancePoolAwsAttributesPb) (*InstancePoolAwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAwsAttributes{}
	st.Availability = pb.Availability
	st.SpotBidPricePercent = pb.SpotBidPricePercent
	st.ZoneId = pb.ZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolAzureAttributesToPb(st *InstancePoolAzureAttributes) (*instancePoolAzureAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAzureAttributesPb{}
	pb.Availability = st.Availability
	pb.SpotBidMaxPrice = st.SpotBidMaxPrice

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolAzureAttributesPb struct {
	Availability    InstancePoolAzureAttributesAvailability `json:"availability,omitempty"`
	SpotBidMaxPrice float64                                 `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAzureAttributesFromPb(pb *instancePoolAzureAttributesPb) (*InstancePoolAzureAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAzureAttributes{}
	st.Availability = pb.Availability
	st.SpotBidMaxPrice = pb.SpotBidMaxPrice

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAzureAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAzureAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolGcpAttributesToPb(st *InstancePoolGcpAttributes) (*instancePoolGcpAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolGcpAttributesPb{}
	pb.GcpAvailability = st.GcpAvailability
	pb.LocalSsdCount = st.LocalSsdCount
	pb.ZoneId = st.ZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolGcpAttributesPb struct {
	GcpAvailability GcpAvailability `json:"gcp_availability,omitempty"`
	LocalSsdCount   int             `json:"local_ssd_count,omitempty"`
	ZoneId          string          `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolGcpAttributesFromPb(pb *instancePoolGcpAttributesPb) (*InstancePoolGcpAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolGcpAttributes{}
	st.GcpAvailability = pb.GcpAvailability
	st.LocalSsdCount = pb.LocalSsdCount
	st.ZoneId = pb.ZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolGcpAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolGcpAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolPermissionToPb(st *InstancePoolPermission) (*instancePoolPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolPermissionPb struct {
	Inherited           bool                        `json:"inherited,omitempty"`
	InheritedFromObject []string                    `json:"inherited_from_object,omitempty"`
	PermissionLevel     InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolPermissionFromPb(pb *instancePoolPermissionPb) (*InstancePoolPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolPermissionsToPb(st *InstancePoolPermissions) (*instancePoolPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolPermissionsPb struct {
	AccessControlList []InstancePoolAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                              `json:"object_id,omitempty"`
	ObjectType        string                              `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolPermissionsFromPb(pb *instancePoolPermissionsPb) (*InstancePoolPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolPermissionsDescriptionToPb(st *InstancePoolPermissionsDescription) (*instancePoolPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolPermissionsDescriptionPb struct {
	Description     string                      `json:"description,omitempty"`
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolPermissionsDescriptionFromPb(pb *instancePoolPermissionsDescriptionPb) (*InstancePoolPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolPermissionsRequestToPb(st *InstancePoolPermissionsRequest) (*instancePoolPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

type instancePoolPermissionsRequestPb struct {
	AccessControlList []InstancePoolAccessControlRequest `json:"access_control_list,omitempty"`
	InstancePoolId    string                             `json:"-" url:"-"`
}

func instancePoolPermissionsRequestFromPb(pb *instancePoolPermissionsRequestPb) (*InstancePoolPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

func instancePoolStatsToPb(st *InstancePoolStats) (*instancePoolStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolStatsPb{}
	pb.IdleCount = st.IdleCount
	pb.PendingIdleCount = st.PendingIdleCount
	pb.PendingUsedCount = st.PendingUsedCount
	pb.UsedCount = st.UsedCount

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instancePoolStatsPb struct {
	IdleCount        int `json:"idle_count,omitempty"`
	PendingIdleCount int `json:"pending_idle_count,omitempty"`
	PendingUsedCount int `json:"pending_used_count,omitempty"`
	UsedCount        int `json:"used_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolStatsFromPb(pb *instancePoolStatsPb) (*InstancePoolStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolStats{}
	st.IdleCount = pb.IdleCount
	st.PendingIdleCount = pb.PendingIdleCount
	st.PendingUsedCount = pb.PendingUsedCount
	st.UsedCount = pb.UsedCount

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func instancePoolStatusToPb(st *InstancePoolStatus) (*instancePoolStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolStatusPb{}
	pb.PendingInstanceErrors = st.PendingInstanceErrors

	return pb, nil
}

type instancePoolStatusPb struct {
	PendingInstanceErrors []PendingInstanceError `json:"pending_instance_errors,omitempty"`
}

func instancePoolStatusFromPb(pb *instancePoolStatusPb) (*InstancePoolStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolStatus{}
	st.PendingInstanceErrors = pb.PendingInstanceErrors

	return st, nil
}

func instanceProfileToPb(st *InstanceProfile) (*instanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instanceProfilePb{}
	pb.IamRoleArn = st.IamRoleArn
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.IsMetaInstanceProfile = st.IsMetaInstanceProfile

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type instanceProfilePb struct {
	IamRoleArn            string `json:"iam_role_arn,omitempty"`
	InstanceProfileArn    string `json:"instance_profile_arn"`
	IsMetaInstanceProfile bool   `json:"is_meta_instance_profile,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instanceProfileFromPb(pb *instanceProfilePb) (*InstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstanceProfile{}
	st.IamRoleArn = pb.IamRoleArn
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.IsMetaInstanceProfile = pb.IsMetaInstanceProfile

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instanceProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instanceProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func libraryToPb(st *Library) (*libraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &libraryPb{}
	pb.Cran = st.Cran
	pb.Egg = st.Egg
	pb.Jar = st.Jar
	pb.Maven = st.Maven
	pb.Pypi = st.Pypi
	pb.Requirements = st.Requirements
	pb.Whl = st.Whl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type libraryPb struct {
	Cran         *RCranLibrary      `json:"cran,omitempty"`
	Egg          string             `json:"egg,omitempty"`
	Jar          string             `json:"jar,omitempty"`
	Maven        *MavenLibrary      `json:"maven,omitempty"`
	Pypi         *PythonPyPiLibrary `json:"pypi,omitempty"`
	Requirements string             `json:"requirements,omitempty"`
	Whl          string             `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func libraryFromPb(pb *libraryPb) (*Library, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Library{}
	st.Cran = pb.Cran
	st.Egg = pb.Egg
	st.Jar = pb.Jar
	st.Maven = pb.Maven
	st.Pypi = pb.Pypi
	st.Requirements = pb.Requirements
	st.Whl = pb.Whl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *libraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st libraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func libraryFullStatusToPb(st *LibraryFullStatus) (*libraryFullStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &libraryFullStatusPb{}
	pb.IsLibraryForAllClusters = st.IsLibraryForAllClusters
	pb.Library = st.Library
	pb.Messages = st.Messages
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type libraryFullStatusPb struct {
	IsLibraryForAllClusters bool                 `json:"is_library_for_all_clusters,omitempty"`
	Library                 *Library             `json:"library,omitempty"`
	Messages                []string             `json:"messages,omitempty"`
	Status                  LibraryInstallStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func libraryFullStatusFromPb(pb *libraryFullStatusPb) (*LibraryFullStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LibraryFullStatus{}
	st.IsLibraryForAllClusters = pb.IsLibraryForAllClusters
	st.Library = pb.Library
	st.Messages = pb.Messages
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *libraryFullStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st libraryFullStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAllClusterLibraryStatusesResponseToPb(st *ListAllClusterLibraryStatusesResponse) (*listAllClusterLibraryStatusesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllClusterLibraryStatusesResponsePb{}
	pb.Statuses = st.Statuses

	return pb, nil
}

type listAllClusterLibraryStatusesResponsePb struct {
	Statuses []ClusterLibraryStatuses `json:"statuses,omitempty"`
}

func listAllClusterLibraryStatusesResponseFromPb(pb *listAllClusterLibraryStatusesResponsePb) (*ListAllClusterLibraryStatusesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllClusterLibraryStatusesResponse{}
	st.Statuses = pb.Statuses

	return st, nil
}

func listAvailableZonesResponseToPb(st *ListAvailableZonesResponse) (*listAvailableZonesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAvailableZonesResponsePb{}
	pb.DefaultZone = st.DefaultZone
	pb.Zones = st.Zones

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAvailableZonesResponsePb struct {
	DefaultZone string   `json:"default_zone,omitempty"`
	Zones       []string `json:"zones,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAvailableZonesResponseFromPb(pb *listAvailableZonesResponsePb) (*ListAvailableZonesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAvailableZonesResponse{}
	st.DefaultZone = pb.DefaultZone
	st.Zones = pb.Zones

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAvailableZonesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAvailableZonesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listClusterCompliancesRequestToPb(st *ListClusterCompliancesRequest) (*listClusterCompliancesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClusterCompliancesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listClusterCompliancesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`
	PolicyId  string `json:"-" url:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClusterCompliancesRequestFromPb(pb *listClusterCompliancesRequestPb) (*ListClusterCompliancesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterCompliancesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClusterCompliancesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClusterCompliancesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listClusterCompliancesResponseToPb(st *ListClusterCompliancesResponse) (*listClusterCompliancesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClusterCompliancesResponsePb{}
	pb.Clusters = st.Clusters
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listClusterCompliancesResponsePb struct {
	Clusters      []ClusterCompliance `json:"clusters,omitempty"`
	NextPageToken string              `json:"next_page_token,omitempty"`
	PrevPageToken string              `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClusterCompliancesResponseFromPb(pb *listClusterCompliancesResponsePb) (*ListClusterCompliancesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterCompliancesResponse{}
	st.Clusters = pb.Clusters
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClusterCompliancesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClusterCompliancesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listClusterPoliciesRequestToPb(st *ListClusterPoliciesRequest) (*listClusterPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClusterPoliciesRequestPb{}
	pb.SortColumn = st.SortColumn
	pb.SortOrder = st.SortOrder

	return pb, nil
}

type listClusterPoliciesRequestPb struct {
	SortColumn ListSortColumn `json:"-" url:"sort_column,omitempty"`
	SortOrder  ListSortOrder  `json:"-" url:"sort_order,omitempty"`
}

func listClusterPoliciesRequestFromPb(pb *listClusterPoliciesRequestPb) (*ListClusterPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterPoliciesRequest{}
	st.SortColumn = pb.SortColumn
	st.SortOrder = pb.SortOrder

	return st, nil
}

func listClustersFilterByToPb(st *ListClustersFilterBy) (*listClustersFilterByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersFilterByPb{}
	pb.ClusterSources = st.ClusterSources
	pb.ClusterStates = st.ClusterStates
	pb.IsPinned = st.IsPinned
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listClustersFilterByPb struct {
	ClusterSources []ClusterSource `json:"cluster_sources,omitempty" url:"cluster_sources,omitempty"`
	ClusterStates  []State         `json:"cluster_states,omitempty" url:"cluster_states,omitempty"`
	IsPinned       bool            `json:"is_pinned,omitempty" url:"is_pinned,omitempty"`
	PolicyId       string          `json:"policy_id,omitempty" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClustersFilterByFromPb(pb *listClustersFilterByPb) (*ListClustersFilterBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersFilterBy{}
	st.ClusterSources = pb.ClusterSources
	st.ClusterStates = pb.ClusterStates
	st.IsPinned = pb.IsPinned
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClustersFilterByPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClustersFilterByPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listClustersRequestToPb(st *ListClustersRequest) (*listClustersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersRequestPb{}
	pb.FilterBy = st.FilterBy
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.SortBy = st.SortBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listClustersRequestPb struct {
	FilterBy  *ListClustersFilterBy `json:"-" url:"filter_by,omitempty"`
	PageSize  int                   `json:"-" url:"page_size,omitempty"`
	PageToken string                `json:"-" url:"page_token,omitempty"`
	SortBy    *ListClustersSortBy   `json:"-" url:"sort_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClustersRequestFromPb(pb *listClustersRequestPb) (*ListClustersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersRequest{}
	st.FilterBy = pb.FilterBy
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.SortBy = pb.SortBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClustersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClustersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listClustersResponseToPb(st *ListClustersResponse) (*listClustersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersResponsePb{}
	pb.Clusters = st.Clusters
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listClustersResponsePb struct {
	Clusters      []ClusterDetails `json:"clusters,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`
	PrevPageToken string           `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClustersResponseFromPb(pb *listClustersResponsePb) (*ListClustersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersResponse{}
	st.Clusters = pb.Clusters
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClustersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClustersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listClustersSortByToPb(st *ListClustersSortBy) (*listClustersSortByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersSortByPb{}
	pb.Direction = st.Direction
	pb.Field = st.Field

	return pb, nil
}

type listClustersSortByPb struct {
	Direction ListClustersSortByDirection `json:"direction,omitempty" url:"direction,omitempty"`
	Field     ListClustersSortByField     `json:"field,omitempty" url:"field,omitempty"`
}

func listClustersSortByFromPb(pb *listClustersSortByPb) (*ListClustersSortBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersSortBy{}
	st.Direction = pb.Direction
	st.Field = pb.Field

	return st, nil
}

func listGlobalInitScriptsResponseToPb(st *ListGlobalInitScriptsResponse) (*listGlobalInitScriptsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listGlobalInitScriptsResponsePb{}
	pb.Scripts = st.Scripts

	return pb, nil
}

type listGlobalInitScriptsResponsePb struct {
	Scripts []GlobalInitScriptDetails `json:"scripts,omitempty"`
}

func listGlobalInitScriptsResponseFromPb(pb *listGlobalInitScriptsResponsePb) (*ListGlobalInitScriptsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGlobalInitScriptsResponse{}
	st.Scripts = pb.Scripts

	return st, nil
}

func listInstancePoolsToPb(st *ListInstancePools) (*listInstancePoolsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstancePoolsPb{}
	pb.InstancePools = st.InstancePools

	return pb, nil
}

type listInstancePoolsPb struct {
	InstancePools []InstancePoolAndStats `json:"instance_pools,omitempty"`
}

func listInstancePoolsFromPb(pb *listInstancePoolsPb) (*ListInstancePools, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstancePools{}
	st.InstancePools = pb.InstancePools

	return st, nil
}

func listInstanceProfilesResponseToPb(st *ListInstanceProfilesResponse) (*listInstanceProfilesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstanceProfilesResponsePb{}
	pb.InstanceProfiles = st.InstanceProfiles

	return pb, nil
}

type listInstanceProfilesResponsePb struct {
	InstanceProfiles []InstanceProfile `json:"instance_profiles,omitempty"`
}

func listInstanceProfilesResponseFromPb(pb *listInstanceProfilesResponsePb) (*ListInstanceProfilesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstanceProfilesResponse{}
	st.InstanceProfiles = pb.InstanceProfiles

	return st, nil
}

func listNodeTypesResponseToPb(st *ListNodeTypesResponse) (*listNodeTypesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNodeTypesResponsePb{}
	pb.NodeTypes = st.NodeTypes

	return pb, nil
}

type listNodeTypesResponsePb struct {
	NodeTypes []NodeType `json:"node_types,omitempty"`
}

func listNodeTypesResponseFromPb(pb *listNodeTypesResponsePb) (*ListNodeTypesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNodeTypesResponse{}
	st.NodeTypes = pb.NodeTypes

	return st, nil
}

func listPoliciesResponseToPb(st *ListPoliciesResponse) (*listPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPoliciesResponsePb{}
	pb.Policies = st.Policies

	return pb, nil
}

type listPoliciesResponsePb struct {
	Policies []Policy `json:"policies,omitempty"`
}

func listPoliciesResponseFromPb(pb *listPoliciesResponsePb) (*ListPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPoliciesResponse{}
	st.Policies = pb.Policies

	return st, nil
}

func listPolicyFamiliesRequestToPb(st *ListPolicyFamiliesRequest) (*listPolicyFamiliesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPolicyFamiliesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPolicyFamiliesRequestPb struct {
	MaxResults int64  `json:"-" url:"max_results,omitempty"`
	PageToken  string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPolicyFamiliesRequestFromPb(pb *listPolicyFamiliesRequestPb) (*ListPolicyFamiliesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPolicyFamiliesRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPolicyFamiliesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPolicyFamiliesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPolicyFamiliesResponseToPb(st *ListPolicyFamiliesResponse) (*listPolicyFamiliesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPolicyFamiliesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PolicyFamilies = st.PolicyFamilies

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPolicyFamiliesResponsePb struct {
	NextPageToken  string         `json:"next_page_token,omitempty"`
	PolicyFamilies []PolicyFamily `json:"policy_families,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPolicyFamiliesResponseFromPb(pb *listPolicyFamiliesResponsePb) (*ListPolicyFamiliesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPolicyFamiliesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PolicyFamilies = pb.PolicyFamilies

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPolicyFamiliesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPolicyFamiliesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func localFileInfoToPb(st *LocalFileInfo) (*localFileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &localFileInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

type localFileInfoPb struct {
	Destination string `json:"destination"`
}

func localFileInfoFromPb(pb *localFileInfoPb) (*LocalFileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LocalFileInfo{}
	st.Destination = pb.Destination

	return st, nil
}

func logAnalyticsInfoToPb(st *LogAnalyticsInfo) (*logAnalyticsInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logAnalyticsInfoPb{}
	pb.LogAnalyticsPrimaryKey = st.LogAnalyticsPrimaryKey
	pb.LogAnalyticsWorkspaceId = st.LogAnalyticsWorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logAnalyticsInfoPb struct {
	LogAnalyticsPrimaryKey  string `json:"log_analytics_primary_key,omitempty"`
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logAnalyticsInfoFromPb(pb *logAnalyticsInfoPb) (*LogAnalyticsInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogAnalyticsInfo{}
	st.LogAnalyticsPrimaryKey = pb.LogAnalyticsPrimaryKey
	st.LogAnalyticsWorkspaceId = pb.LogAnalyticsWorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logAnalyticsInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logAnalyticsInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logSyncStatusToPb(st *LogSyncStatus) (*logSyncStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logSyncStatusPb{}
	pb.LastAttempted = st.LastAttempted
	pb.LastException = st.LastException

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logSyncStatusPb struct {
	LastAttempted int64  `json:"last_attempted,omitempty"`
	LastException string `json:"last_exception,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logSyncStatusFromPb(pb *logSyncStatusPb) (*LogSyncStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogSyncStatus{}
	st.LastAttempted = pb.LastAttempted
	st.LastException = pb.LastException

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logSyncStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logSyncStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type mapAnyPb MapAny

func mapAnyToPb(st *MapAny) (*mapAnyPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := mapAnyPb(*st)
	return &stPb, nil
}

func mapAnyFromPb(stPb *mapAnyPb) (*MapAny, error) {
	if stPb == nil {
		return nil, nil
	}
	st := MapAny(*stPb)
	return &st, nil
}

func mavenLibraryToPb(st *MavenLibrary) (*mavenLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mavenLibraryPb{}
	pb.Coordinates = st.Coordinates
	pb.Exclusions = st.Exclusions
	pb.Repo = st.Repo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type mavenLibraryPb struct {
	Coordinates string   `json:"coordinates"`
	Exclusions  []string `json:"exclusions,omitempty"`
	Repo        string   `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func mavenLibraryFromPb(pb *mavenLibraryPb) (*MavenLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MavenLibrary{}
	st.Coordinates = pb.Coordinates
	st.Exclusions = pb.Exclusions
	st.Repo = pb.Repo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *mavenLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st mavenLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func nodeInstanceTypeToPb(st *NodeInstanceType) (*nodeInstanceTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nodeInstanceTypePb{}
	pb.InstanceTypeId = st.InstanceTypeId
	pb.LocalDiskSizeGb = st.LocalDiskSizeGb
	pb.LocalDisks = st.LocalDisks
	pb.LocalNvmeDiskSizeGb = st.LocalNvmeDiskSizeGb
	pb.LocalNvmeDisks = st.LocalNvmeDisks

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type nodeInstanceTypePb struct {
	InstanceTypeId      string `json:"instance_type_id"`
	LocalDiskSizeGb     int    `json:"local_disk_size_gb,omitempty"`
	LocalDisks          int    `json:"local_disks,omitempty"`
	LocalNvmeDiskSizeGb int    `json:"local_nvme_disk_size_gb,omitempty"`
	LocalNvmeDisks      int    `json:"local_nvme_disks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nodeInstanceTypeFromPb(pb *nodeInstanceTypePb) (*NodeInstanceType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NodeInstanceType{}
	st.InstanceTypeId = pb.InstanceTypeId
	st.LocalDiskSizeGb = pb.LocalDiskSizeGb
	st.LocalDisks = pb.LocalDisks
	st.LocalNvmeDiskSizeGb = pb.LocalNvmeDiskSizeGb
	st.LocalNvmeDisks = pb.LocalNvmeDisks

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nodeInstanceTypePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nodeInstanceTypePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func nodeTypeToPb(st *NodeType) (*nodeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nodeTypePb{}
	pb.Category = st.Category
	pb.Description = st.Description
	pb.DisplayOrder = st.DisplayOrder
	pb.InstanceTypeId = st.InstanceTypeId
	pb.IsDeprecated = st.IsDeprecated
	pb.IsEncryptedInTransit = st.IsEncryptedInTransit
	pb.IsGraviton = st.IsGraviton
	pb.IsHidden = st.IsHidden
	pb.IsIoCacheEnabled = st.IsIoCacheEnabled
	pb.MemoryMb = st.MemoryMb
	pb.NodeInfo = st.NodeInfo
	pb.NodeInstanceType = st.NodeInstanceType
	pb.NodeTypeId = st.NodeTypeId
	pb.NumCores = st.NumCores
	pb.NumGpus = st.NumGpus
	pb.PhotonDriverCapable = st.PhotonDriverCapable
	pb.PhotonWorkerCapable = st.PhotonWorkerCapable
	pb.SupportClusterTags = st.SupportClusterTags
	pb.SupportEbsVolumes = st.SupportEbsVolumes
	pb.SupportPortForwarding = st.SupportPortForwarding

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type nodeTypePb struct {
	Category              string                 `json:"category"`
	Description           string                 `json:"description"`
	DisplayOrder          int                    `json:"display_order,omitempty"`
	InstanceTypeId        string                 `json:"instance_type_id"`
	IsDeprecated          bool                   `json:"is_deprecated,omitempty"`
	IsEncryptedInTransit  bool                   `json:"is_encrypted_in_transit,omitempty"`
	IsGraviton            bool                   `json:"is_graviton,omitempty"`
	IsHidden              bool                   `json:"is_hidden,omitempty"`
	IsIoCacheEnabled      bool                   `json:"is_io_cache_enabled,omitempty"`
	MemoryMb              int                    `json:"memory_mb"`
	NodeInfo              *CloudProviderNodeInfo `json:"node_info,omitempty"`
	NodeInstanceType      *NodeInstanceType      `json:"node_instance_type,omitempty"`
	NodeTypeId            string                 `json:"node_type_id"`
	NumCores              float64                `json:"num_cores"`
	NumGpus               int                    `json:"num_gpus,omitempty"`
	PhotonDriverCapable   bool                   `json:"photon_driver_capable,omitempty"`
	PhotonWorkerCapable   bool                   `json:"photon_worker_capable,omitempty"`
	SupportClusterTags    bool                   `json:"support_cluster_tags,omitempty"`
	SupportEbsVolumes     bool                   `json:"support_ebs_volumes,omitempty"`
	SupportPortForwarding bool                   `json:"support_port_forwarding,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nodeTypeFromPb(pb *nodeTypePb) (*NodeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NodeType{}
	st.Category = pb.Category
	st.Description = pb.Description
	st.DisplayOrder = pb.DisplayOrder
	st.InstanceTypeId = pb.InstanceTypeId
	st.IsDeprecated = pb.IsDeprecated
	st.IsEncryptedInTransit = pb.IsEncryptedInTransit
	st.IsGraviton = pb.IsGraviton
	st.IsHidden = pb.IsHidden
	st.IsIoCacheEnabled = pb.IsIoCacheEnabled
	st.MemoryMb = pb.MemoryMb
	st.NodeInfo = pb.NodeInfo
	st.NodeInstanceType = pb.NodeInstanceType
	st.NodeTypeId = pb.NodeTypeId
	st.NumCores = pb.NumCores
	st.NumGpus = pb.NumGpus
	st.PhotonDriverCapable = pb.PhotonDriverCapable
	st.PhotonWorkerCapable = pb.PhotonWorkerCapable
	st.SupportClusterTags = pb.SupportClusterTags
	st.SupportEbsVolumes = pb.SupportEbsVolumes
	st.SupportPortForwarding = pb.SupportPortForwarding

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nodeTypePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nodeTypePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pendingInstanceErrorToPb(st *PendingInstanceError) (*pendingInstanceErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pendingInstanceErrorPb{}
	pb.InstanceId = st.InstanceId
	pb.Message = st.Message

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pendingInstanceErrorPb struct {
	InstanceId string `json:"instance_id,omitempty"`
	Message    string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pendingInstanceErrorFromPb(pb *pendingInstanceErrorPb) (*PendingInstanceError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PendingInstanceError{}
	st.InstanceId = pb.InstanceId
	st.Message = pb.Message

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pendingInstanceErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pendingInstanceErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func permanentDeleteClusterToPb(st *PermanentDeleteCluster) (*permanentDeleteClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permanentDeleteClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type permanentDeleteClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func permanentDeleteClusterFromPb(pb *permanentDeleteClusterPb) (*PermanentDeleteCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermanentDeleteCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func pinClusterToPb(st *PinCluster) (*pinClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pinClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type pinClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func pinClusterFromPb(pb *pinClusterPb) (*PinCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PinCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func policyToPb(st *Policy) (*policyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &policyPb{}
	pb.CreatedAtTimestamp = st.CreatedAtTimestamp
	pb.CreatorUserName = st.CreatorUserName
	pb.Definition = st.Definition
	pb.Description = st.Description
	pb.IsDefault = st.IsDefault
	pb.Libraries = st.Libraries
	pb.MaxClustersPerUser = st.MaxClustersPerUser
	pb.Name = st.Name
	pb.PolicyFamilyDefinitionOverrides = st.PolicyFamilyDefinitionOverrides
	pb.PolicyFamilyId = st.PolicyFamilyId
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type policyPb struct {
	CreatedAtTimestamp              int64     `json:"created_at_timestamp,omitempty"`
	CreatorUserName                 string    `json:"creator_user_name,omitempty"`
	Definition                      string    `json:"definition,omitempty"`
	Description                     string    `json:"description,omitempty"`
	IsDefault                       bool      `json:"is_default,omitempty"`
	Libraries                       []Library `json:"libraries,omitempty"`
	MaxClustersPerUser              int64     `json:"max_clusters_per_user,omitempty"`
	Name                            string    `json:"name,omitempty"`
	PolicyFamilyDefinitionOverrides string    `json:"policy_family_definition_overrides,omitempty"`
	PolicyFamilyId                  string    `json:"policy_family_id,omitempty"`
	PolicyId                        string    `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func policyFromPb(pb *policyPb) (*Policy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Policy{}
	st.CreatedAtTimestamp = pb.CreatedAtTimestamp
	st.CreatorUserName = pb.CreatorUserName
	st.Definition = pb.Definition
	st.Description = pb.Description
	st.IsDefault = pb.IsDefault
	st.Libraries = pb.Libraries
	st.MaxClustersPerUser = pb.MaxClustersPerUser
	st.Name = pb.Name
	st.PolicyFamilyDefinitionOverrides = pb.PolicyFamilyDefinitionOverrides
	st.PolicyFamilyId = pb.PolicyFamilyId
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *policyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st policyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func policyFamilyToPb(st *PolicyFamily) (*policyFamilyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &policyFamilyPb{}
	pb.Definition = st.Definition
	pb.Description = st.Description
	pb.Name = st.Name
	pb.PolicyFamilyId = st.PolicyFamilyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type policyFamilyPb struct {
	Definition     string `json:"definition,omitempty"`
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	PolicyFamilyId string `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func policyFamilyFromPb(pb *policyFamilyPb) (*PolicyFamily, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PolicyFamily{}
	st.Definition = pb.Definition
	st.Description = pb.Description
	st.Name = pb.Name
	st.PolicyFamilyId = pb.PolicyFamilyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *policyFamilyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st policyFamilyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pythonPyPiLibraryToPb(st *PythonPyPiLibrary) (*pythonPyPiLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pythonPyPiLibraryPb{}
	pb.Package = st.Package
	pb.Repo = st.Repo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pythonPyPiLibraryPb struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pythonPyPiLibraryFromPb(pb *pythonPyPiLibraryPb) (*PythonPyPiLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PythonPyPiLibrary{}
	st.Package = pb.Package
	st.Repo = pb.Repo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pythonPyPiLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pythonPyPiLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func rCranLibraryToPb(st *RCranLibrary) (*rCranLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rCranLibraryPb{}
	pb.Package = st.Package
	pb.Repo = st.Repo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type rCranLibraryPb struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func rCranLibraryFromPb(pb *rCranLibraryPb) (*RCranLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RCranLibrary{}
	st.Package = pb.Package
	st.Repo = pb.Repo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *rCranLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st rCranLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func removeInstanceProfileToPb(st *RemoveInstanceProfile) (*removeInstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeInstanceProfilePb{}
	pb.InstanceProfileArn = st.InstanceProfileArn

	return pb, nil
}

type removeInstanceProfilePb struct {
	InstanceProfileArn string `json:"instance_profile_arn"`
}

func removeInstanceProfileFromPb(pb *removeInstanceProfilePb) (*RemoveInstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveInstanceProfile{}
	st.InstanceProfileArn = pb.InstanceProfileArn

	return st, nil
}

func resizeClusterToPb(st *ResizeCluster) (*resizeClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resizeClusterPb{}
	pb.Autoscale = st.Autoscale
	pb.ClusterId = st.ClusterId
	pb.NumWorkers = st.NumWorkers

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resizeClusterPb struct {
	Autoscale  *AutoScale `json:"autoscale,omitempty"`
	ClusterId  string     `json:"cluster_id"`
	NumWorkers int        `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resizeClusterFromPb(pb *resizeClusterPb) (*ResizeCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResizeCluster{}
	st.Autoscale = pb.Autoscale
	st.ClusterId = pb.ClusterId
	st.NumWorkers = pb.NumWorkers

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resizeClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resizeClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func restartClusterToPb(st *RestartCluster) (*restartClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restartClusterPb{}
	pb.ClusterId = st.ClusterId
	pb.RestartUser = st.RestartUser

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type restartClusterPb struct {
	ClusterId   string `json:"cluster_id"`
	RestartUser string `json:"restart_user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restartClusterFromPb(pb *restartClusterPb) (*RestartCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartCluster{}
	st.ClusterId = pb.ClusterId
	st.RestartUser = pb.RestartUser

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restartClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restartClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func resultsToPb(st *Results) (*resultsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultsPb{}
	pb.Cause = st.Cause
	pb.Data = st.Data
	pb.FileName = st.FileName
	pb.FileNames = st.FileNames
	pb.IsJsonSchema = st.IsJsonSchema
	pb.Pos = st.Pos
	pb.ResultType = st.ResultType
	pb.Schema = st.Schema
	pb.Summary = st.Summary
	pb.Truncated = st.Truncated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultsPb struct {
	Cause        string           `json:"cause,omitempty"`
	Data         any              `json:"data,omitempty"`
	FileName     string           `json:"fileName,omitempty"`
	FileNames    []string         `json:"fileNames,omitempty"`
	IsJsonSchema bool             `json:"isJsonSchema,omitempty"`
	Pos          int              `json:"pos,omitempty"`
	ResultType   ResultType       `json:"resultType,omitempty"`
	Schema       []map[string]any `json:"schema,omitempty"`
	Summary      string           `json:"summary,omitempty"`
	Truncated    bool             `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultsFromPb(pb *resultsPb) (*Results, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Results{}
	st.Cause = pb.Cause
	st.Data = pb.Data
	st.FileName = pb.FileName
	st.FileNames = pb.FileNames
	st.IsJsonSchema = pb.IsJsonSchema
	st.Pos = pb.Pos
	st.ResultType = pb.ResultType
	st.Schema = pb.Schema
	st.Summary = pb.Summary
	st.Truncated = pb.Truncated

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func s3StorageInfoToPb(st *S3StorageInfo) (*s3StorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &s3StorageInfoPb{}
	pb.CannedAcl = st.CannedAcl
	pb.Destination = st.Destination
	pb.EnableEncryption = st.EnableEncryption
	pb.EncryptionType = st.EncryptionType
	pb.Endpoint = st.Endpoint
	pb.KmsKey = st.KmsKey
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type s3StorageInfoPb struct {
	CannedAcl        string `json:"canned_acl,omitempty"`
	Destination      string `json:"destination"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	Region           string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func s3StorageInfoFromPb(pb *s3StorageInfoPb) (*S3StorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &S3StorageInfo{}
	st.CannedAcl = pb.CannedAcl
	st.Destination = pb.Destination
	st.EnableEncryption = pb.EnableEncryption
	st.EncryptionType = pb.EncryptionType
	st.Endpoint = pb.Endpoint
	st.KmsKey = pb.KmsKey
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *s3StorageInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st s3StorageInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sparkNodeToPb(st *SparkNode) (*sparkNodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkNodePb{}
	pb.HostPrivateIp = st.HostPrivateIp
	pb.InstanceId = st.InstanceId
	pb.NodeAwsAttributes = st.NodeAwsAttributes
	pb.NodeId = st.NodeId
	pb.PrivateIp = st.PrivateIp
	pb.PublicDns = st.PublicDns
	pb.StartTimestamp = st.StartTimestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sparkNodePb struct {
	HostPrivateIp     string                  `json:"host_private_ip,omitempty"`
	InstanceId        string                  `json:"instance_id,omitempty"`
	NodeAwsAttributes *SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
	NodeId            string                  `json:"node_id,omitempty"`
	PrivateIp         string                  `json:"private_ip,omitempty"`
	PublicDns         string                  `json:"public_dns,omitempty"`
	StartTimestamp    int64                   `json:"start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sparkNodeFromPb(pb *sparkNodePb) (*SparkNode, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkNode{}
	st.HostPrivateIp = pb.HostPrivateIp
	st.InstanceId = pb.InstanceId
	st.NodeAwsAttributes = pb.NodeAwsAttributes
	st.NodeId = pb.NodeId
	st.PrivateIp = pb.PrivateIp
	st.PublicDns = pb.PublicDns
	st.StartTimestamp = pb.StartTimestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkNodePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkNodePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sparkNodeAwsAttributesToPb(st *SparkNodeAwsAttributes) (*sparkNodeAwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkNodeAwsAttributesPb{}
	pb.IsSpot = st.IsSpot

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sparkNodeAwsAttributesPb struct {
	IsSpot bool `json:"is_spot,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sparkNodeAwsAttributesFromPb(pb *sparkNodeAwsAttributesPb) (*SparkNodeAwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkNodeAwsAttributes{}
	st.IsSpot = pb.IsSpot

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkNodeAwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkNodeAwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sparkVersionToPb(st *SparkVersion) (*sparkVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkVersionPb{}
	pb.Key = st.Key
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sparkVersionPb struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sparkVersionFromPb(pb *sparkVersionPb) (*SparkVersion, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkVersion{}
	st.Key = pb.Key
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func startClusterToPb(st *StartCluster) (*startClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type startClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func startClusterFromPb(pb *startClusterPb) (*StartCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func terminationReasonToPb(st *TerminationReason) (*terminationReasonPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &terminationReasonPb{}
	pb.Code = st.Code
	pb.Parameters = st.Parameters
	pb.Type = st.Type

	return pb, nil
}

type terminationReasonPb struct {
	Code       TerminationReasonCode `json:"code,omitempty"`
	Parameters map[string]string     `json:"parameters,omitempty"`
	Type       TerminationReasonType `json:"type,omitempty"`
}

func terminationReasonFromPb(pb *terminationReasonPb) (*TerminationReason, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TerminationReason{}
	st.Code = pb.Code
	st.Parameters = pb.Parameters
	st.Type = pb.Type

	return st, nil
}

func uninstallLibrariesToPb(st *UninstallLibraries) (*uninstallLibrariesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &uninstallLibrariesPb{}
	pb.ClusterId = st.ClusterId
	pb.Libraries = st.Libraries

	return pb, nil
}

type uninstallLibrariesPb struct {
	ClusterId string    `json:"cluster_id"`
	Libraries []Library `json:"libraries"`
}

func uninstallLibrariesFromPb(pb *uninstallLibrariesPb) (*UninstallLibraries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UninstallLibraries{}
	st.ClusterId = pb.ClusterId
	st.Libraries = pb.Libraries

	return st, nil
}

func unpinClusterToPb(st *UnpinCluster) (*unpinClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpinClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

type unpinClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func unpinClusterFromPb(pb *unpinClusterPb) (*UnpinCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpinCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

func updateClusterToPb(st *UpdateCluster) (*updateClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateClusterPb{}
	pb.Cluster = st.Cluster
	pb.ClusterId = st.ClusterId
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

type updateClusterPb struct {
	Cluster    *UpdateClusterResource `json:"cluster,omitempty"`
	ClusterId  string                 `json:"cluster_id"`
	UpdateMask string                 `json:"update_mask"`
}

func updateClusterFromPb(pb *updateClusterPb) (*UpdateCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCluster{}
	st.Cluster = pb.Cluster
	st.ClusterId = pb.ClusterId
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

func updateClusterResourceToPb(st *UpdateClusterResource) (*updateClusterResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateClusterResourcePb{}
	pb.Autoscale = st.Autoscale
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.ClusterLogConf = st.ClusterLogConf
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	pb.DataSecurityMode = st.DataSecurityMode
	pb.DockerImage = st.DockerImage
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.Kind = st.Kind
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.RuntimeEngine = st.RuntimeEngine
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateClusterResourcePb struct {
	Autoscale                  *AutoScale        `json:"autoscale,omitempty"`
	AutoterminationMinutes     int               `json:"autotermination_minutes,omitempty"`
	AwsAttributes              *AwsAttributes    `json:"aws_attributes,omitempty"`
	AzureAttributes            *AzureAttributes  `json:"azure_attributes,omitempty"`
	ClusterLogConf             *ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	ClusterName                string            `json:"cluster_name,omitempty"`
	CustomTags                 map[string]string `json:"custom_tags,omitempty"`
	DataSecurityMode           DataSecurityMode  `json:"data_security_mode,omitempty"`
	DockerImage                *DockerImage      `json:"docker_image,omitempty"`
	DriverInstancePoolId       string            `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId           string            `json:"driver_node_type_id,omitempty"`
	EnableElasticDisk          bool              `json:"enable_elastic_disk,omitempty"`
	EnableLocalDiskEncryption  bool              `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes              *GcpAttributes    `json:"gcp_attributes,omitempty"`
	InitScripts                []InitScriptInfo  `json:"init_scripts,omitempty"`
	InstancePoolId             string            `json:"instance_pool_id,omitempty"`
	IsSingleNode               bool              `json:"is_single_node,omitempty"`
	Kind                       Kind              `json:"kind,omitempty"`
	NodeTypeId                 string            `json:"node_type_id,omitempty"`
	NumWorkers                 int               `json:"num_workers,omitempty"`
	PolicyId                   string            `json:"policy_id,omitempty"`
	RemoteDiskThroughput       int               `json:"remote_disk_throughput,omitempty"`
	RuntimeEngine              RuntimeEngine     `json:"runtime_engine,omitempty"`
	SingleUserName             string            `json:"single_user_name,omitempty"`
	SparkConf                  map[string]string `json:"spark_conf,omitempty"`
	SparkEnvVars               map[string]string `json:"spark_env_vars,omitempty"`
	SparkVersion               string            `json:"spark_version,omitempty"`
	SshPublicKeys              []string          `json:"ssh_public_keys,omitempty"`
	TotalInitialRemoteDiskSize int               `json:"total_initial_remote_disk_size,omitempty"`
	UseMlRuntime               bool              `json:"use_ml_runtime,omitempty"`
	WorkloadType               *WorkloadType     `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateClusterResourceFromPb(pb *updateClusterResourcePb) (*UpdateClusterResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateClusterResource{}
	st.Autoscale = pb.Autoscale
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.ClusterLogConf = pb.ClusterLogConf
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	st.DataSecurityMode = pb.DataSecurityMode
	st.DockerImage = pb.DockerImage
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.Kind = pb.Kind
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.RuntimeEngine = pb.RuntimeEngine
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateClusterResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateClusterResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func volumesStorageInfoToPb(st *VolumesStorageInfo) (*volumesStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumesStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

type volumesStorageInfoPb struct {
	Destination string `json:"destination"`
}

func volumesStorageInfoFromPb(pb *volumesStorageInfoPb) (*VolumesStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumesStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

func workloadTypeToPb(st *WorkloadType) (*workloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workloadTypePb{}
	pb.Clients = st.Clients

	return pb, nil
}

type workloadTypePb struct {
	Clients ClientsTypes `json:"clients"`
}

func workloadTypeFromPb(pb *workloadTypePb) (*WorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkloadType{}
	st.Clients = pb.Clients

	return st, nil
}

func workspaceStorageInfoToPb(st *WorkspaceStorageInfo) (*workspaceStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

type workspaceStorageInfoPb struct {
	Destination string `json:"destination"`
}

func workspaceStorageInfoFromPb(pb *workspaceStorageInfoPb) (*WorkspaceStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
