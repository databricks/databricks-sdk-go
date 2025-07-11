// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddInstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn string `json:"iam_role_arn,omitempty"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn string `json:"instance_profile_arn"`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AddInstanceProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AddInstanceProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A storage location in Adls Gen2
type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination string `json:"destination"`
}

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	MaxWorkers int `json:"max_workers,omitempty"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers int `json:"min_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutoScale) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoScale) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attributes set during cluster creation which are related to Amazon Web
// Services.
type AwsAttributes struct {
	Availability AwsAvailability `json:"availability,omitempty"`
	// The number of volumes launched for each instance. Users can choose up to
	// 10 volumes. This feature is only enabled for supported node types. Legacy
	// node types cannot specify custom EBS volumes. For node types with no
	// instance store, at least one EBS volume needs to be specified; otherwise,
	// cluster creation will fail.
	//
	// These EBS volumes will be mounted at `/ebs0`, `/ebs1`, and etc. Instance
	// store volumes will be mounted at `/local_disk0`, `/local_disk1`, and etc.
	//
	// If EBS volumes are attached, Databricks will configure Spark to use only
	// the EBS volumes for scratch storage because heterogenously sized scratch
	// devices can lead to inefficient disk utilization. If no EBS volumes are
	// attached, Databricks will configure Spark to use instance store volumes.
	//
	// Please note that if EBS volumes are specified, then the Spark
	// configuration `spark.local.dir` will be overridden.
	EbsVolumeCount int `json:"ebs_volume_count,omitempty"`
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	EbsVolumeIops int `json:"ebs_volume_iops,omitempty"`
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	EbsVolumeSize int `json:"ebs_volume_size,omitempty"`
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	EbsVolumeThroughput int `json:"ebs_volume_throughput,omitempty"`
	// The type of EBS volumes that will be launched with this cluster.
	EbsVolumeType EbsVolumeType `json:"ebs_volume_type,omitempty"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. If this value is greater than 0, the cluster driver
	// node in particular will be placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand int `json:"first_on_demand,omitempty"`
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	SpotBidPricePercent int `json:"spot_bid_price_percent,omitempty"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. If the zone specified is "auto", will try to place cluster
	// in a zone with high availability, and will retry placement in a different
	// AZ if there is not enough capacity.
	//
	// The list of available zones as well as the default value can be found by
	// using the `List Zones` method.
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AwsAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AwsAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones.
//
// Note: If `first_on_demand` is zero, this availability type will be used for
// the entire cluster.
type AwsAvailability string

const AwsAvailabilityOnDemand AwsAvailability = `ON_DEMAND`

const AwsAvailabilitySpot AwsAvailability = `SPOT`

const AwsAvailabilitySpotWithFallback AwsAvailability = `SPOT_WITH_FALLBACK`

// String representation for [fmt.Print]
func (f *AwsAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AwsAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND`, `SPOT`, `SPOT_WITH_FALLBACK`:
		*f = AwsAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND", "SPOT", "SPOT_WITH_FALLBACK"`, v)
	}
}

// Values returns all possible values for AwsAvailability.
//
// There is no guarantee on the order of the values in the slice.
func (f *AwsAvailability) Values() []AwsAvailability {
	return []AwsAvailability{
		AwsAvailabilityOnDemand,
		AwsAvailabilitySpot,
		AwsAvailabilitySpotWithFallback,
	}
}

// Type always returns AwsAvailability to satisfy [pflag.Value] interface
func (f *AwsAvailability) Type() string {
	return "AwsAvailability"
}

// Attributes set during cluster creation which are related to Microsoft Azure.
type AzureAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones. Note: If `first_on_demand` is zero, this
	// availability type will be used for the entire cluster.
	Availability AzureAvailability `json:"availability,omitempty"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. This value should be greater than 0, to make sure
	// the cluster driver node is placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand int `json:"first_on_demand,omitempty"`
	// Defines values necessary to configure and run Azure Log Analytics agent
	LogAnalyticsInfo *LogAnalyticsInfo `json:"log_analytics_info,omitempty"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AzureAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AzureAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones. Note: If `first_on_demand` is zero, this availability type will be used
// for the entire cluster.
type AzureAvailability string

const AzureAvailabilityOnDemandAzure AzureAvailability = `ON_DEMAND_AZURE`

const AzureAvailabilitySpotAzure AzureAvailability = `SPOT_AZURE`

const AzureAvailabilitySpotWithFallbackAzure AzureAvailability = `SPOT_WITH_FALLBACK_AZURE`

// String representation for [fmt.Print]
func (f *AzureAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AzureAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_AZURE`, `SPOT_AZURE`, `SPOT_WITH_FALLBACK_AZURE`:
		*f = AzureAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_AZURE", "SPOT_AZURE", "SPOT_WITH_FALLBACK_AZURE"`, v)
	}
}

// Values returns all possible values for AzureAvailability.
//
// There is no guarantee on the order of the values in the slice.
func (f *AzureAvailability) Values() []AzureAvailability {
	return []AzureAvailability{
		AzureAvailabilityOnDemandAzure,
		AzureAvailabilitySpotAzure,
		AzureAvailabilitySpotWithFallbackAzure,
	}
}

// Type always returns AzureAvailability to satisfy [pflag.Value] interface
func (f *AzureAvailability) Type() string {
	return "AzureAvailability"
}

type CancelCommand struct {
	ClusterId string `json:"clusterId,omitempty"`

	CommandId string `json:"commandId,omitempty"`

	ContextId string `json:"contextId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CancelCommand) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CancelCommand) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ChangeClusterOwner struct {
	ClusterId string `json:"cluster_id"`
	// New owner of the cluster_id after this RPC.
	OwnerUsername string `json:"owner_username"`
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	Jobs bool `json:"jobs,omitempty"`
	// With notebooks set, this cluster can be used for notebooks
	Notebooks bool `json:"notebooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClientsTypes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClientsTypes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CloneCluster struct {
	// The cluster that is being cloned.
	SourceClusterId string `json:"source_cluster_id"`
}

type CloudProviderNodeInfo struct {
	// Status as reported by the cloud provider
	Status []CloudProviderNodeStatus `json:"status,omitempty"`
}

type CloudProviderNodeStatus string

const CloudProviderNodeStatusNotAvailableInRegion CloudProviderNodeStatus = `NotAvailableInRegion`

const CloudProviderNodeStatusNotEnabledOnSubscription CloudProviderNodeStatus = `NotEnabledOnSubscription`

// String representation for [fmt.Print]
func (f *CloudProviderNodeStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CloudProviderNodeStatus) Set(v string) error {
	switch v {
	case `NotAvailableInRegion`, `NotEnabledOnSubscription`:
		*f = CloudProviderNodeStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NotAvailableInRegion", "NotEnabledOnSubscription"`, v)
	}
}

// Values returns all possible values for CloudProviderNodeStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *CloudProviderNodeStatus) Values() []CloudProviderNodeStatus {
	return []CloudProviderNodeStatus{
		CloudProviderNodeStatusNotAvailableInRegion,
		CloudProviderNodeStatusNotEnabledOnSubscription,
	}
}

// Type always returns CloudProviderNodeStatus to satisfy [pflag.Value] interface
func (f *CloudProviderNodeStatus) Type() string {
	return "CloudProviderNodeStatus"
}

type ClusterAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAccessControlResponse struct {
	// All permissions.
	AllPermissions []ClusterPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Common set of attributes set during cluster creation. These attributes cannot
// be changed over the lifetime of a cluster.
type ClusterAttributes struct {
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
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool `json:"is_single_node,omitempty"`

	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string `json:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`

	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterCompliance struct {
	// Canonical unique identifier for a cluster.
	ClusterId string `json:"cluster_id"`
	// Whether this cluster is in compliance with the latest version of its
	// policy.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterCompliance) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterCompliance) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Describes all of the metadata about a single Spark cluster in Databricks.
type ClusterDetails struct {
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
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster log delivery status.
	ClusterLogStatus *LogSyncStatus `json:"cluster_log_status,omitempty"`
	// Total amount of cluster memory, in megabytes
	ClusterMemoryMb int64 `json:"cluster_memory_mb,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName string `json:"cluster_name,omitempty"`
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request.
	ClusterSource ClusterSource `json:"cluster_source,omitempty"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - Creator: <username_of_creator>
	//
	// - ClusterName: <name_of_cluster>
	//
	// - ClusterId: <id_of_cluster>
	//
	// - Name: <Databricks internal use>
	DefaultTags map[string]string `json:"default_tags,omitempty"`
	// Custom docker image BYOC
	DockerImage *DockerImage `json:"docker_image,omitempty"`
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
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Nodes on which the Spark executors reside.
	Executors []SparkNode `json:"executors,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	JdbcPort int `json:"jdbc_port,omitempty"`

	Kind Kind `json:"kind,omitempty"`
	// the timestamp that the cluster was started/restarted
	LastRestartedTime int64 `json:"last_restarted_time,omitempty"`
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	LastStateLossTime int64 `json:"last_state_loss_time,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	SparkContextId int64 `json:"spark_context_id,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string `json:"spark_version,omitempty"`
	// The spec contains a snapshot of the latest user specified settings that
	// were used to create/edit the cluster. Note: not included in the response
	// of the ListClusters API.
	Spec *ClusterSpec `json:"spec,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	StartTime int64 `json:"start_time,omitempty"`
	// Current state of the cluster.
	State State `json:"state,omitempty"`
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	StateMessage string `json:"state_message,omitempty"`
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	TerminatedTime int64 `json:"terminated_time,omitempty"`
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	TerminationReason *TerminationReason `json:"termination_reason,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`

	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterEvent struct {
	ClusterId string `json:"cluster_id"`

	DataPlaneEventDetails *DataPlaneEventDetails `json:"data_plane_event_details,omitempty"`

	Details *EventDetails `json:"details,omitempty"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp int64 `json:"timestamp,omitempty"`

	Type EventType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterEvent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterEvent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId string `json:"cluster_id,omitempty"`
	// Status of all libraries on the cluster.
	LibraryStatuses []LibraryFullStatus `json:"library_statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterLibraryStatuses) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterLibraryStatuses) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Cluster log delivery config
type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 *S3StorageInfo `json:"s3,omitempty"`
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
}

type ClusterPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ClusterPermissionLevel string

const ClusterPermissionLevelCanAttachTo ClusterPermissionLevel = `CAN_ATTACH_TO`

const ClusterPermissionLevelCanManage ClusterPermissionLevel = `CAN_MANAGE`

const ClusterPermissionLevelCanRestart ClusterPermissionLevel = `CAN_RESTART`

// String representation for [fmt.Print]
func (f *ClusterPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_MANAGE`, `CAN_RESTART`:
		*f = ClusterPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_MANAGE", "CAN_RESTART"`, v)
	}
}

// Values returns all possible values for ClusterPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterPermissionLevel) Values() []ClusterPermissionLevel {
	return []ClusterPermissionLevel{
		ClusterPermissionLevelCanAttachTo,
		ClusterPermissionLevelCanManage,
		ClusterPermissionLevelCanRestart,
	}
}

// Type always returns ClusterPermissionLevel to satisfy [pflag.Value] interface
func (f *ClusterPermissionLevel) Type() string {
	return "ClusterPermissionLevel"
}

type ClusterPermissions struct {
	AccessControlList []ClusterAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterPermissionsRequest struct {
	AccessControlList []ClusterAccessControlRequest `json:"access_control_list,omitempty"`
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" url:"-"`
}

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPolicyAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPolicyAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterPolicyAccessControlResponse struct {
	// All permissions.
	AllPermissions []ClusterPolicyPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPolicyAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPolicyAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterPolicyPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPolicyPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPolicyPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ClusterPolicyPermissionLevel string

const ClusterPolicyPermissionLevelCanUse ClusterPolicyPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *ClusterPolicyPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterPolicyPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = ClusterPolicyPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Values returns all possible values for ClusterPolicyPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterPolicyPermissionLevel) Values() []ClusterPolicyPermissionLevel {
	return []ClusterPolicyPermissionLevel{
		ClusterPolicyPermissionLevelCanUse,
	}
}

// Type always returns ClusterPolicyPermissionLevel to satisfy [pflag.Value] interface
func (f *ClusterPolicyPermissionLevel) Type() string {
	return "ClusterPolicyPermissionLevel"
}

type ClusterPolicyPermissions struct {
	AccessControlList []ClusterPolicyAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPolicyPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPolicyPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterPolicyPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterPolicyPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterPolicyPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterPolicyPermissionsRequest struct {
	AccessControlList []ClusterPolicyAccessControlRequest `json:"access_control_list,omitempty"`
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" url:"-"`
}

// Represents a change to the cluster settings required for the cluster to
// become compliant with its policy.
type ClusterSettingsChange struct {
	// The field where this change would be made.
	Field string `json:"field,omitempty"`
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	NewValue string `json:"new_value,omitempty"`
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterSettingsChange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterSettingsChange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterSize struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterSize) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterSize) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.
type ClusterSource string

const ClusterSourceApi ClusterSource = `API`

const ClusterSourceJob ClusterSource = `JOB`

const ClusterSourceModels ClusterSource = `MODELS`

const ClusterSourcePipeline ClusterSource = `PIPELINE`

const ClusterSourcePipelineMaintenance ClusterSource = `PIPELINE_MAINTENANCE`

const ClusterSourceSql ClusterSource = `SQL`

const ClusterSourceUi ClusterSource = `UI`

// String representation for [fmt.Print]
func (f *ClusterSource) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterSource) Set(v string) error {
	switch v {
	case `API`, `JOB`, `MODELS`, `PIPELINE`, `PIPELINE_MAINTENANCE`, `SQL`, `UI`:
		*f = ClusterSource(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API", "JOB", "MODELS", "PIPELINE", "PIPELINE_MAINTENANCE", "SQL", "UI"`, v)
	}
}

// Values returns all possible values for ClusterSource.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterSource) Values() []ClusterSource {
	return []ClusterSource{
		ClusterSourceApi,
		ClusterSourceJob,
		ClusterSourceModels,
		ClusterSourcePipeline,
		ClusterSourcePipelineMaintenance,
		ClusterSourceSql,
		ClusterSourceUi,
	}
}

// Type always returns ClusterSource to satisfy [pflag.Value] interface
func (f *ClusterSource) Type() string {
	return "ClusterSource"
}

// Contains a snapshot of the latest user specified settings that were used to
// create/edit the cluster.
type ClusterSpec struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
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
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool `json:"is_single_node,omitempty"`

	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string `json:"spark_version,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`

	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId string `json:"-" url:"cluster_id"`
}

type Command struct {
	// Running cluster id
	ClusterId string `json:"clusterId,omitempty"`
	// Executable code
	Command string `json:"command,omitempty"`
	// Running context id
	ContextId string `json:"contextId,omitempty"`

	Language Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Command) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Command) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CommandStatus string

const CommandStatusCancelled CommandStatus = `Cancelled`

const CommandStatusCancelling CommandStatus = `Cancelling`

const CommandStatusError CommandStatus = `Error`

const CommandStatusFinished CommandStatus = `Finished`

const CommandStatusQueued CommandStatus = `Queued`

const CommandStatusRunning CommandStatus = `Running`

// String representation for [fmt.Print]
func (f *CommandStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommandStatus) Set(v string) error {
	switch v {
	case `Cancelled`, `Cancelling`, `Error`, `Finished`, `Queued`, `Running`:
		*f = CommandStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Cancelled", "Cancelling", "Error", "Finished", "Queued", "Running"`, v)
	}
}

// Values returns all possible values for CommandStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *CommandStatus) Values() []CommandStatus {
	return []CommandStatus{
		CommandStatusCancelled,
		CommandStatusCancelling,
		CommandStatusError,
		CommandStatusFinished,
		CommandStatusQueued,
		CommandStatusRunning,
	}
}

// Type always returns CommandStatus to satisfy [pflag.Value] interface
func (f *CommandStatus) Type() string {
	return "CommandStatus"
}

type CommandStatusRequest struct {
	ClusterId string `json:"-" url:"clusterId"`

	CommandId string `json:"-" url:"commandId"`

	ContextId string `json:"-" url:"contextId"`
}

type CommandStatusResponse struct {
	Id string `json:"id,omitempty"`

	Results *Results `json:"results,omitempty"`

	Status CommandStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CommandStatusResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CommandStatusResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ContextStatus string

const ContextStatusError ContextStatus = `Error`

const ContextStatusPending ContextStatus = `Pending`

const ContextStatusRunning ContextStatus = `Running`

// String representation for [fmt.Print]
func (f *ContextStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ContextStatus) Set(v string) error {
	switch v {
	case `Error`, `Pending`, `Running`:
		*f = ContextStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Error", "Pending", "Running"`, v)
	}
}

// Values returns all possible values for ContextStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ContextStatus) Values() []ContextStatus {
	return []ContextStatus{
		ContextStatusError,
		ContextStatusPending,
		ContextStatusRunning,
	}
}

// Type always returns ContextStatus to satisfy [pflag.Value] interface
func (f *ContextStatus) Type() string {
	return "ContextStatus"
}

type ContextStatusRequest struct {
	ClusterId string `json:"-" url:"clusterId"`

	ContextId string `json:"-" url:"contextId"`
}

type ContextStatusResponse struct {
	Id string `json:"id,omitempty"`

	Status ContextStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ContextStatusResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ContextStatusResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
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
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom *CloneCluster `json:"clone_from,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool `json:"is_single_node,omitempty"`

	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string `json:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`

	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateClusterResponse struct {
	ClusterId string `json:"cluster_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateClusterResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateClusterResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateContext struct {
	// Running cluster id
	ClusterId string `json:"clusterId,omitempty"`

	Language Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateContext) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateContext) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
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
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes `json:"gcp_attributes,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateInstancePool) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateInstancePool) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId string `json:"instance_pool_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateInstancePoolResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateInstancePoolResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library `json:"libraries,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name,omitempty"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePolicyResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePolicyResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateResponse struct {
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Created struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Created) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Created) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CustomPolicyTag struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	Key string `json:"key"`
	// The value of the tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CustomPolicyTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CustomPolicyTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataPlaneEventDetails struct {
	EventType DataPlaneEventDetailsEventType `json:"event_type,omitempty"`

	ExecutorFailures int `json:"executor_failures,omitempty"`

	HostId string `json:"host_id,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DataPlaneEventDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneEventDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DataPlaneEventDetailsEventType string

const DataPlaneEventDetailsEventTypeNodeBlacklisted DataPlaneEventDetailsEventType = `NODE_BLACKLISTED`

const DataPlaneEventDetailsEventTypeNodeExcludedDecommissioned DataPlaneEventDetailsEventType = `NODE_EXCLUDED_DECOMMISSIONED`

// String representation for [fmt.Print]
func (f *DataPlaneEventDetailsEventType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataPlaneEventDetailsEventType) Set(v string) error {
	switch v {
	case `NODE_BLACKLISTED`, `NODE_EXCLUDED_DECOMMISSIONED`:
		*f = DataPlaneEventDetailsEventType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NODE_BLACKLISTED", "NODE_EXCLUDED_DECOMMISSIONED"`, v)
	}
}

// Values returns all possible values for DataPlaneEventDetailsEventType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataPlaneEventDetailsEventType) Values() []DataPlaneEventDetailsEventType {
	return []DataPlaneEventDetailsEventType{
		DataPlaneEventDetailsEventTypeNodeBlacklisted,
		DataPlaneEventDetailsEventTypeNodeExcludedDecommissioned,
	}
}

// Type always returns DataPlaneEventDetailsEventType to satisfy [pflag.Value] interface
func (f *DataPlaneEventDetailsEventType) Type() string {
	return "DataPlaneEventDetailsEventType"
}

// Data security mode decides what data governance model to use when accessing
// data from a cluster.
//
// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate access
// mode depending on your compute configuration. *
// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
//
// The following modes can be used regardless of `kind`. * `NONE`: No security
// isolation for multiple users sharing the cluster. Data governance features
// are not available in this mode. * `SINGLE_USER`: A secure cluster that can
// only be exclusively used by a single user specified in `single_user_name`.
// Most programming languages, cluster features and data governance features are
// available in this mode. * `USER_ISOLATION`: A secure cluster that can be
// shared by multiple users. Cluster users are fully isolated so that they
// cannot see each other's data and credentials. Most data governance features
// are supported in this mode. But programming languages and cluster features
// might be limited.
//
// The following modes are deprecated starting with Databricks Runtime 15.0 and
// will be removed for future Databricks Runtime versions:
//
// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table ACL
// clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating from
// legacy Passthrough on high concurrency clusters. * `LEGACY_SINGLE_USER`: This
// mode is for users migrating from legacy Passthrough on standard clusters. *
// `LEGACY_SINGLE_USER_STANDARD`: This mode provides a way that doesn’t have
// UC nor passthrough enabled.
type DataSecurityMode string

// <Databricks> will choose the most appropriate access mode depending on your
// compute configuration.
const DataSecurityModeDataSecurityModeAuto DataSecurityMode = `DATA_SECURITY_MODE_AUTO`

// Alias for `SINGLE_USER`.
const DataSecurityModeDataSecurityModeDedicated DataSecurityMode = `DATA_SECURITY_MODE_DEDICATED`

// Alias for `USER_ISOLATION`.
const DataSecurityModeDataSecurityModeStandard DataSecurityMode = `DATA_SECURITY_MODE_STANDARD`

// This mode is for users migrating from legacy Passthrough on high concurrency
// clusters.
const DataSecurityModeLegacyPassthrough DataSecurityMode = `LEGACY_PASSTHROUGH`

// This mode is for users migrating from legacy Passthrough on standard
// clusters.
const DataSecurityModeLegacySingleUser DataSecurityMode = `LEGACY_SINGLE_USER`

// This mode provides a way that doesn’t have UC nor passthrough enabled.
const DataSecurityModeLegacySingleUserStandard DataSecurityMode = `LEGACY_SINGLE_USER_STANDARD`

// This mode is for users migrating from legacy Table ACL clusters.
const DataSecurityModeLegacyTableAcl DataSecurityMode = `LEGACY_TABLE_ACL`

// No security isolation for multiple users sharing the cluster. Data governance
// features are not available in this mode.
const DataSecurityModeNone DataSecurityMode = `NONE`

// A secure cluster that can only be exclusively used by a single user specified
// in `single_user_name`. Most programming languages, cluster features and data
// governance features are available in this mode.
const DataSecurityModeSingleUser DataSecurityMode = `SINGLE_USER`

// A secure cluster that can be shared by multiple users. Cluster users are
// fully isolated so that they cannot see each other's data and credentials.
// Most data governance features are supported in this mode. But programming
// languages and cluster features might be limited.
const DataSecurityModeUserIsolation DataSecurityMode = `USER_ISOLATION`

// String representation for [fmt.Print]
func (f *DataSecurityMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataSecurityMode) Set(v string) error {
	switch v {
	case `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, `DATA_SECURITY_MODE_STANDARD`, `LEGACY_PASSTHROUGH`, `LEGACY_SINGLE_USER`, `LEGACY_SINGLE_USER_STANDARD`, `LEGACY_TABLE_ACL`, `NONE`, `SINGLE_USER`, `USER_ISOLATION`:
		*f = DataSecurityMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_SECURITY_MODE_AUTO", "DATA_SECURITY_MODE_DEDICATED", "DATA_SECURITY_MODE_STANDARD", "LEGACY_PASSTHROUGH", "LEGACY_SINGLE_USER", "LEGACY_SINGLE_USER_STANDARD", "LEGACY_TABLE_ACL", "NONE", "SINGLE_USER", "USER_ISOLATION"`, v)
	}
}

// Values returns all possible values for DataSecurityMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *DataSecurityMode) Values() []DataSecurityMode {
	return []DataSecurityMode{
		DataSecurityModeDataSecurityModeAuto,
		DataSecurityModeDataSecurityModeDedicated,
		DataSecurityModeDataSecurityModeStandard,
		DataSecurityModeLegacyPassthrough,
		DataSecurityModeLegacySingleUser,
		DataSecurityModeLegacySingleUserStandard,
		DataSecurityModeLegacyTableAcl,
		DataSecurityModeNone,
		DataSecurityModeSingleUser,
		DataSecurityModeUserIsolation,
	}
}

// Type always returns DataSecurityMode to satisfy [pflag.Value] interface
func (f *DataSecurityMode) Type() string {
	return "DataSecurityMode"
}

// A storage location in DBFS
type DbfsStorageInfo struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination string `json:"destination"`
}

type DeleteCluster struct {
	// The cluster to be terminated.
	ClusterId string `json:"cluster_id"`
}

type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId string `json:"-" url:"-"`
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	InstancePoolId string `json:"instance_pool_id"`
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId string `json:"policy_id"`
}

type DestroyContext struct {
	ClusterId string `json:"clusterId"`

	ContextId string `json:"contextId"`
}

// Describes the disks that are launched for each instance in the spark cluster.
// For example, if the cluster has 3 instances, each instance is configured to
// launch 2 disks, 100 GiB each, then Databricks will launch a total of 6 disks,
// 100 GiB each, for this cluster.
type DiskSpec struct {
	// The number of disks launched for each instance: - This feature is only
	// enabled for supported node types. - Users can choose up to the limit of
	// the disks supported by the node type. - For node types with no OS disk,
	// at least one disk must be specified; otherwise, cluster creation will
	// fail.
	//
	// If disks are attached, Databricks will configure Spark to use only the
	// disks for scratch storage, because heterogenously sized scratch devices
	// can lead to inefficient disk utilization. If no disks are attached,
	// Databricks will configure Spark to use instance store disks.
	//
	// Note: If disks are specified, then the Spark configuration
	// `spark.local.dir` will be overridden.
	//
	// Disks will be mounted at: - For AWS: `/ebs0`, `/ebs1`, and etc. - For
	// Azure: `/remote_volume0`, `/remote_volume1`, and etc.
	DiskCount int `json:"disk_count,omitempty"`

	DiskIops int `json:"disk_iops,omitempty"`
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	DiskSize int `json:"disk_size,omitempty"`

	DiskThroughput int `json:"disk_throughput,omitempty"`
	// The type of disks that will be launched with this cluster.
	DiskType *DiskType `json:"disk_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DiskSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DiskSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Describes the disk type.
type DiskType struct {
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType `json:"azure_disk_volume_type,omitempty"`

	EbsVolumeType DiskTypeEbsVolumeType `json:"ebs_volume_type,omitempty"`
}

// All Azure Disk types that Databricks supports. See
// https://docs.microsoft.com/en-us/azure/storage/storage-about-disks-and-vhds-linux#types-of-disks
type DiskTypeAzureDiskVolumeType string

const DiskTypeAzureDiskVolumeTypePremiumLrs DiskTypeAzureDiskVolumeType = `PREMIUM_LRS`

const DiskTypeAzureDiskVolumeTypeStandardLrs DiskTypeAzureDiskVolumeType = `STANDARD_LRS`

// String representation for [fmt.Print]
func (f *DiskTypeAzureDiskVolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DiskTypeAzureDiskVolumeType) Set(v string) error {
	switch v {
	case `PREMIUM_LRS`, `STANDARD_LRS`:
		*f = DiskTypeAzureDiskVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PREMIUM_LRS", "STANDARD_LRS"`, v)
	}
}

// Values returns all possible values for DiskTypeAzureDiskVolumeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DiskTypeAzureDiskVolumeType) Values() []DiskTypeAzureDiskVolumeType {
	return []DiskTypeAzureDiskVolumeType{
		DiskTypeAzureDiskVolumeTypePremiumLrs,
		DiskTypeAzureDiskVolumeTypeStandardLrs,
	}
}

// Type always returns DiskTypeAzureDiskVolumeType to satisfy [pflag.Value] interface
func (f *DiskTypeAzureDiskVolumeType) Type() string {
	return "DiskTypeAzureDiskVolumeType"
}

// All EBS volume types that Databricks supports. See
// https://aws.amazon.com/ebs/details/ for details.
type DiskTypeEbsVolumeType string

const DiskTypeEbsVolumeTypeGeneralPurposeSsd DiskTypeEbsVolumeType = `GENERAL_PURPOSE_SSD`

const DiskTypeEbsVolumeTypeThroughputOptimizedHdd DiskTypeEbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`

// String representation for [fmt.Print]
func (f *DiskTypeEbsVolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DiskTypeEbsVolumeType) Set(v string) error {
	switch v {
	case `GENERAL_PURPOSE_SSD`, `THROUGHPUT_OPTIMIZED_HDD`:
		*f = DiskTypeEbsVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENERAL_PURPOSE_SSD", "THROUGHPUT_OPTIMIZED_HDD"`, v)
	}
}

// Values returns all possible values for DiskTypeEbsVolumeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DiskTypeEbsVolumeType) Values() []DiskTypeEbsVolumeType {
	return []DiskTypeEbsVolumeType{
		DiskTypeEbsVolumeTypeGeneralPurposeSsd,
		DiskTypeEbsVolumeTypeThroughputOptimizedHdd,
	}
}

// Type always returns DiskTypeEbsVolumeType to satisfy [pflag.Value] interface
func (f *DiskTypeEbsVolumeType) Type() string {
	return "DiskTypeEbsVolumeType"
}

type DockerBasicAuth struct {
	// Password of the user
	Password string `json:"password,omitempty"`
	// Name of the user
	Username string `json:"username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DockerBasicAuth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DockerBasicAuth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DockerImage struct {
	// Basic auth with username and password
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
	// URL of the docker image.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DockerImage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DockerImage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// All EBS volume types that Databricks supports. See
// https://aws.amazon.com/ebs/details/ for details.
type EbsVolumeType string

const EbsVolumeTypeGeneralPurposeSsd EbsVolumeType = `GENERAL_PURPOSE_SSD`

const EbsVolumeTypeThroughputOptimizedHdd EbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`

// String representation for [fmt.Print]
func (f *EbsVolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EbsVolumeType) Set(v string) error {
	switch v {
	case `GENERAL_PURPOSE_SSD`, `THROUGHPUT_OPTIMIZED_HDD`:
		*f = EbsVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENERAL_PURPOSE_SSD", "THROUGHPUT_OPTIMIZED_HDD"`, v)
	}
}

// Values returns all possible values for EbsVolumeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EbsVolumeType) Values() []EbsVolumeType {
	return []EbsVolumeType{
		EbsVolumeTypeGeneralPurposeSsd,
		EbsVolumeTypeThroughputOptimizedHdd,
	}
}

// Type always returns EbsVolumeType to satisfy [pflag.Value] interface
func (f *EbsVolumeType) Type() string {
	return "EbsVolumeType"
}

type EditCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
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
	// ID of the cluster
	ClusterId string `json:"cluster_id"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool `json:"is_single_node,omitempty"`

	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string `json:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`

	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EditInstancePool struct {
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Instance pool ID
	InstancePoolId string `json:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditInstancePool) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditInstancePool) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EditPolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library `json:"libraries,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name,omitempty"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
	// The ID of the policy to update.
	PolicyId string `json:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EditPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnforceClusterComplianceRequest struct {
	// The ID of the cluster you want to enforce policy compliance on.
	ClusterId string `json:"cluster_id"`
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnforceClusterComplianceRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnforceClusterComplianceRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnforceClusterComplianceResponse struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	Changes []ClusterSettingsChange `json:"changes,omitempty"`
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	HasChanges bool `json:"has_changes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnforceClusterComplianceResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnforceClusterComplianceResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type Environment struct {
	// Use `environment_version` instead.
	Client string `json:"client,omitempty"`
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a valid pip requirements file line per
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/.
	// Allowed dependencies include a requirement specifier, an archive URL, a
	// local project path (such as WSFS or UC Volumes in Databricks), or a VCS
	// project URL.
	Dependencies []string `json:"dependencies,omitempty"`
	// Required. Environment version used by the environment. Each version comes
	// with a specific Python version and a set of Python packages. The version
	// is a string, consisting of an integer.
	EnvironmentVersion string `json:"environment_version,omitempty"`
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	JarDependencies []string `json:"jar_dependencies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Environment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Environment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

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
	// More details about the change in driver's state
	DriverStateMessage string `json:"driver_state_message,omitempty"`
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	EnableTerminationForNodeBlocklisted bool `json:"enable_termination_for_node_blocklisted,omitempty"`

	FreeSpace int64 `json:"free_space,omitempty"`
	// List of global and cluster init scripts associated with this cluster
	// event.
	InitScripts *InitScriptEventDetails `json:"init_scripts,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EventDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EventDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The cause of a change in target size.
type EventDetailsCause string

const EventDetailsCauseAutorecovery EventDetailsCause = `AUTORECOVERY`

const EventDetailsCauseAutoscale EventDetailsCause = `AUTOSCALE`

const EventDetailsCauseReplaceBadNodes EventDetailsCause = `REPLACE_BAD_NODES`

const EventDetailsCauseUserRequest EventDetailsCause = `USER_REQUEST`

// String representation for [fmt.Print]
func (f *EventDetailsCause) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EventDetailsCause) Set(v string) error {
	switch v {
	case `AUTORECOVERY`, `AUTOSCALE`, `REPLACE_BAD_NODES`, `USER_REQUEST`:
		*f = EventDetailsCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTORECOVERY", "AUTOSCALE", "REPLACE_BAD_NODES", "USER_REQUEST"`, v)
	}
}

// Values returns all possible values for EventDetailsCause.
//
// There is no guarantee on the order of the values in the slice.
func (f *EventDetailsCause) Values() []EventDetailsCause {
	return []EventDetailsCause{
		EventDetailsCauseAutorecovery,
		EventDetailsCauseAutoscale,
		EventDetailsCauseReplaceBadNodes,
		EventDetailsCauseUserRequest,
	}
}

// Type always returns EventDetailsCause to satisfy [pflag.Value] interface
func (f *EventDetailsCause) Type() string {
	return "EventDetailsCause"
}

type EventType string

const EventTypeAddNodesFailed EventType = `ADD_NODES_FAILED`

const EventTypeAutomaticClusterUpdate EventType = `AUTOMATIC_CLUSTER_UPDATE`

const EventTypeAutoscalingBackoff EventType = `AUTOSCALING_BACKOFF`

const EventTypeAutoscalingFailed EventType = `AUTOSCALING_FAILED`

const EventTypeAutoscalingStatsReport EventType = `AUTOSCALING_STATS_REPORT`

const EventTypeClusterMigrated EventType = `CLUSTER_MIGRATED`

const EventTypeCreating EventType = `CREATING`

const EventTypeDbfsDown EventType = `DBFS_DOWN`

const EventTypeDidNotExpandDisk EventType = `DID_NOT_EXPAND_DISK`

const EventTypeDriverHealthy EventType = `DRIVER_HEALTHY`

const EventTypeDriverNotResponding EventType = `DRIVER_NOT_RESPONDING`

const EventTypeDriverUnavailable EventType = `DRIVER_UNAVAILABLE`

const EventTypeEdited EventType = `EDITED`

const EventTypeExpandedDisk EventType = `EXPANDED_DISK`

const EventTypeFailedToExpandDisk EventType = `FAILED_TO_EXPAND_DISK`

const EventTypeInitScriptsFinished EventType = `INIT_SCRIPTS_FINISHED`

const EventTypeInitScriptsStarted EventType = `INIT_SCRIPTS_STARTED`

const EventTypeMetastoreDown EventType = `METASTORE_DOWN`

const EventTypeNodesLost EventType = `NODES_LOST`

const EventTypeNodeBlacklisted EventType = `NODE_BLACKLISTED`

const EventTypeNodeExcludedDecommissioned EventType = `NODE_EXCLUDED_DECOMMISSIONED`

const EventTypePinned EventType = `PINNED`

const EventTypeResizing EventType = `RESIZING`

const EventTypeRestarting EventType = `RESTARTING`

const EventTypeRunning EventType = `RUNNING`

const EventTypeSparkException EventType = `SPARK_EXCEPTION`

const EventTypeStarting EventType = `STARTING`

const EventTypeTerminating EventType = `TERMINATING`

const EventTypeUnpinned EventType = `UNPINNED`

const EventTypeUpsizeCompleted EventType = `UPSIZE_COMPLETED`

// String representation for [fmt.Print]
func (f *EventType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EventType) Set(v string) error {
	switch v {
	case `ADD_NODES_FAILED`, `AUTOMATIC_CLUSTER_UPDATE`, `AUTOSCALING_BACKOFF`, `AUTOSCALING_FAILED`, `AUTOSCALING_STATS_REPORT`, `CLUSTER_MIGRATED`, `CREATING`, `DBFS_DOWN`, `DID_NOT_EXPAND_DISK`, `DRIVER_HEALTHY`, `DRIVER_NOT_RESPONDING`, `DRIVER_UNAVAILABLE`, `EDITED`, `EXPANDED_DISK`, `FAILED_TO_EXPAND_DISK`, `INIT_SCRIPTS_FINISHED`, `INIT_SCRIPTS_STARTED`, `METASTORE_DOWN`, `NODES_LOST`, `NODE_BLACKLISTED`, `NODE_EXCLUDED_DECOMMISSIONED`, `PINNED`, `RESIZING`, `RESTARTING`, `RUNNING`, `SPARK_EXCEPTION`, `STARTING`, `TERMINATING`, `UNPINNED`, `UPSIZE_COMPLETED`:
		*f = EventType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD_NODES_FAILED", "AUTOMATIC_CLUSTER_UPDATE", "AUTOSCALING_BACKOFF", "AUTOSCALING_FAILED", "AUTOSCALING_STATS_REPORT", "CLUSTER_MIGRATED", "CREATING", "DBFS_DOWN", "DID_NOT_EXPAND_DISK", "DRIVER_HEALTHY", "DRIVER_NOT_RESPONDING", "DRIVER_UNAVAILABLE", "EDITED", "EXPANDED_DISK", "FAILED_TO_EXPAND_DISK", "INIT_SCRIPTS_FINISHED", "INIT_SCRIPTS_STARTED", "METASTORE_DOWN", "NODES_LOST", "NODE_BLACKLISTED", "NODE_EXCLUDED_DECOMMISSIONED", "PINNED", "RESIZING", "RESTARTING", "RUNNING", "SPARK_EXCEPTION", "STARTING", "TERMINATING", "UNPINNED", "UPSIZE_COMPLETED"`, v)
	}
}

// Values returns all possible values for EventType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EventType) Values() []EventType {
	return []EventType{
		EventTypeAddNodesFailed,
		EventTypeAutomaticClusterUpdate,
		EventTypeAutoscalingBackoff,
		EventTypeAutoscalingFailed,
		EventTypeAutoscalingStatsReport,
		EventTypeClusterMigrated,
		EventTypeCreating,
		EventTypeDbfsDown,
		EventTypeDidNotExpandDisk,
		EventTypeDriverHealthy,
		EventTypeDriverNotResponding,
		EventTypeDriverUnavailable,
		EventTypeEdited,
		EventTypeExpandedDisk,
		EventTypeFailedToExpandDisk,
		EventTypeInitScriptsFinished,
		EventTypeInitScriptsStarted,
		EventTypeMetastoreDown,
		EventTypeNodesLost,
		EventTypeNodeBlacklisted,
		EventTypeNodeExcludedDecommissioned,
		EventTypePinned,
		EventTypeResizing,
		EventTypeRestarting,
		EventTypeRunning,
		EventTypeSparkException,
		EventTypeStarting,
		EventTypeTerminating,
		EventTypeUnpinned,
		EventTypeUpsizeCompleted,
	}
}

// Type always returns EventType to satisfy [pflag.Value] interface
func (f *EventType) Type() string {
	return "EventType"
}

// Attributes set during cluster creation which are related to GCP.
type GcpAttributes struct {
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs, on-demand VMs, or preemptible VMs with a fallback
	// to on-demand VMs if the former is unavailable.
	Availability GcpAvailability `json:"availability,omitempty"`
	// Boot disk size in GB
	BootDiskSize int `json:"boot_disk_size,omitempty"`
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount int `json:"local_ssd_count,omitempty"`
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// 'availability' field instead.
	UsePreemptibleExecutors bool `json:"use_preemptible_executors,omitempty"`
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default]. -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GcpAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GcpAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This field determines whether the instance pool will contain preemptible VMs,
// on-demand VMs, or preemptible VMs with a fallback to on-demand VMs if the
// former is unavailable.
type GcpAvailability string

const GcpAvailabilityOnDemandGcp GcpAvailability = `ON_DEMAND_GCP`

const GcpAvailabilityPreemptibleGcp GcpAvailability = `PREEMPTIBLE_GCP`

const GcpAvailabilityPreemptibleWithFallbackGcp GcpAvailability = `PREEMPTIBLE_WITH_FALLBACK_GCP`

// String representation for [fmt.Print]
func (f *GcpAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GcpAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_GCP`, `PREEMPTIBLE_GCP`, `PREEMPTIBLE_WITH_FALLBACK_GCP`:
		*f = GcpAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_GCP", "PREEMPTIBLE_GCP", "PREEMPTIBLE_WITH_FALLBACK_GCP"`, v)
	}
}

// Values returns all possible values for GcpAvailability.
//
// There is no guarantee on the order of the values in the slice.
func (f *GcpAvailability) Values() []GcpAvailability {
	return []GcpAvailability{
		GcpAvailabilityOnDemandGcp,
		GcpAvailabilityPreemptibleGcp,
		GcpAvailabilityPreemptibleWithFallbackGcp,
	}
}

// Type always returns GcpAvailability to satisfy [pflag.Value] interface
func (f *GcpAvailability) Type() string {
	return "GcpAvailability"
}

// A storage location in Google Cloud Platform's GCS
type GcsStorageInfo struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination string `json:"destination"`
}

type GetClusterComplianceRequest struct {
	// The ID of the cluster to get the compliance status
	ClusterId string `json:"-" url:"cluster_id"`
}

type GetClusterComplianceResponse struct {
	// Whether the cluster is compliant with its policy or not. Clusters could
	// be out of compliance if the policy was updated after the cluster was last
	// edited.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetClusterComplianceResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetClusterComplianceResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" url:"-"`
}

type GetClusterPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ClusterPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" url:"-"`
}

type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" url:"-"`
}

type GetClusterPolicyPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ClusterPolicyPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" url:"-"`
}

type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string `json:"-" url:"policy_id"`
}

type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	ClusterId string `json:"-" url:"cluster_id"`
}

type GetEvents struct {
	// The ID of the cluster to retrieve events about.
	ClusterId string `json:"cluster_id"`
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	EndTime int64 `json:"end_time,omitempty"`
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	EventTypes []EventType `json:"event_types,omitempty"`
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	Limit int64 `json:"limit,omitempty"`
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	Offset int64 `json:"offset,omitempty"`
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	Order GetEventsOrder `json:"order,omitempty"`
	// The maximum number of events to include in a page of events. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is empty or 0, the server will decide the number
	// of results to be returned. The field has to be in the range [0,500]. If
	// the value is outside the range, the server enforces 0 or 500.
	PageSize int `json:"page_size,omitempty"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of events respectively. If page_token
	// is empty, the first page is returned.
	PageToken string `json:"page_token,omitempty"`
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	StartTime int64 `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEvents) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEvents) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetEventsOrder string

const GetEventsOrderAsc GetEventsOrder = `ASC`

const GetEventsOrderDesc GetEventsOrder = `DESC`

// String representation for [fmt.Print]
func (f *GetEventsOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetEventsOrder) Set(v string) error {
	switch v {
	case `ASC`, `DESC`:
		*f = GetEventsOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASC", "DESC"`, v)
	}
}

// Values returns all possible values for GetEventsOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetEventsOrder) Values() []GetEventsOrder {
	return []GetEventsOrder{
		GetEventsOrderAsc,
		GetEventsOrderDesc,
	}
}

// Type always returns GetEventsOrder to satisfy [pflag.Value] interface
func (f *GetEventsOrder) Type() string {
	return "GetEventsOrder"
}

type GetEventsResponse struct {
	Events []ClusterEvent `json:"events,omitempty"`
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage *GetEvents `json:"next_page,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string `json:"prev_page_token,omitempty"`
	// Deprecated: Returns 0 when request uses page_token. Will start returning
	// zero when request uses offset/limit soon.
	//
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	TotalCount int64 `json:"total_count,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEventsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEventsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId string `json:"-" url:"-"`
}

type GetInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Tags that are added by Databricks regardless of any ``custom_tags``,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
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
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes `json:"gcp_attributes,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Canonical unique identifier for the pool.
	InstancePoolId string `json:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Current state of the instance pool.
	State InstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus `json:"status,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetInstancePool) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetInstancePool) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" url:"-"`
}

type GetInstancePoolPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []InstancePoolPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" url:"-"`
}

type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId string `json:"-" url:"instance_pool_id"`
}

type GetPolicyFamilyRequest struct {
	// The family ID about which to retrieve information.
	PolicyFamilyId string `json:"-" url:"-"`
	// The version number for the family to fetch. Defaults to the latest
	// version.
	Version int64 `json:"-" url:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPolicyFamilyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPolicyFamilyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetSparkVersionsResponse struct {
	// All the available Spark versions.
	Versions []SparkVersion `json:"versions,omitempty"`
}

type GlobalInitScriptCreateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	Name string `json:"name"`
	// The position of a global init script, where 0 represents the first script
	// to run, 1 is the second script to run, in ascending order.
	//
	// If you omit the numeric position for a new global init script, it
	// defaults to last position. It will run after all current scripts. Setting
	// any value greater than the position of the last script is equivalent to
	// the last position. Example: Take three existing scripts with positions 0,
	// 1, and 2. Any position of (3) or greater puts the script in the last
	// position. If an explicit position value conflicts with an existing script
	// value, your request succeeds, but the original script at that position
	// and all later scripts have their positions incremented by 1.
	Position int `json:"position,omitempty"`
	// The Base64-encoded content of the script.
	Script string `json:"script"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GlobalInitScriptCreateRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GlobalInitScriptCreateRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GlobalInitScriptDetails struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	CreatedBy string `json:"created_by,omitempty"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	Name string `json:"name,omitempty"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position int `json:"position,omitempty"`
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt int `json:"updated_at,omitempty"`
	// The username of the user who last updated the script
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GlobalInitScriptDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GlobalInitScriptDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	CreatedBy string `json:"created_by,omitempty"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	Name string `json:"name,omitempty"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position int `json:"position,omitempty"`
	// The Base64-encoded content of the script.
	Script string `json:"script,omitempty"`
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt int `json:"updated_at,omitempty"`
	// The username of the user who last updated the script
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GlobalInitScriptDetailsWithContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GlobalInitScriptDetailsWithContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GlobalInitScriptUpdateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	Name string `json:"name"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order. To move the script to
	// run first, set its position to 0.
	//
	// To move the script to the end, set its position to any value greater or
	// equal to the position of the last script. Example, three existing scripts
	// with positions 0, 1, and 2. Any position value of 2 or greater puts the
	// script in the last position (2).
	//
	// If an explicit position value conflicts with an existing script, your
	// request succeeds, but the original script at that position and all later
	// scripts have their positions incremented by 1.
	Position int `json:"position,omitempty"`
	// The Base64-encoded content of the script.
	Script string `json:"script"`
	// The ID of the global init script.
	ScriptId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GlobalInitScriptUpdateRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GlobalInitScriptUpdateRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InitScriptEventDetails struct {
	// The cluster scoped init scripts associated with this cluster event.
	Cluster []InitScriptInfoAndExecutionDetails `json:"cluster,omitempty"`
	// The global init scripts associated with this cluster event.
	Global []InitScriptInfoAndExecutionDetails `json:"global,omitempty"`
	// The private ip of the node we are reporting init script execution details
	// for (we will select the execution details from only one node rather than
	// reporting the execution details from every node to keep these event
	// details small)
	//
	// This should only be defined for the INIT_SCRIPTS_FINISHED event
	ReportedForNode string `json:"reported_for_node,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InitScriptEventDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InitScriptEventDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Result of attempted script execution
type InitScriptExecutionDetailsInitScriptExecutionStatus string

const InitScriptExecutionDetailsInitScriptExecutionStatusFailedExecution InitScriptExecutionDetailsInitScriptExecutionStatus = `FAILED_EXECUTION`

const InitScriptExecutionDetailsInitScriptExecutionStatusFailedFetch InitScriptExecutionDetailsInitScriptExecutionStatus = `FAILED_FETCH`

const InitScriptExecutionDetailsInitScriptExecutionStatusFuseMountFailed InitScriptExecutionDetailsInitScriptExecutionStatus = `FUSE_MOUNT_FAILED`

const InitScriptExecutionDetailsInitScriptExecutionStatusNotExecuted InitScriptExecutionDetailsInitScriptExecutionStatus = `NOT_EXECUTED`

const InitScriptExecutionDetailsInitScriptExecutionStatusSkipped InitScriptExecutionDetailsInitScriptExecutionStatus = `SKIPPED`

const InitScriptExecutionDetailsInitScriptExecutionStatusSucceeded InitScriptExecutionDetailsInitScriptExecutionStatus = `SUCCEEDED`

const InitScriptExecutionDetailsInitScriptExecutionStatusUnknown InitScriptExecutionDetailsInitScriptExecutionStatus = `UNKNOWN`

// String representation for [fmt.Print]
func (f *InitScriptExecutionDetailsInitScriptExecutionStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InitScriptExecutionDetailsInitScriptExecutionStatus) Set(v string) error {
	switch v {
	case `FAILED_EXECUTION`, `FAILED_FETCH`, `FUSE_MOUNT_FAILED`, `NOT_EXECUTED`, `SKIPPED`, `SUCCEEDED`, `UNKNOWN`:
		*f = InitScriptExecutionDetailsInitScriptExecutionStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_EXECUTION", "FAILED_FETCH", "FUSE_MOUNT_FAILED", "NOT_EXECUTED", "SKIPPED", "SUCCEEDED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for InitScriptExecutionDetailsInitScriptExecutionStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *InitScriptExecutionDetailsInitScriptExecutionStatus) Values() []InitScriptExecutionDetailsInitScriptExecutionStatus {
	return []InitScriptExecutionDetailsInitScriptExecutionStatus{
		InitScriptExecutionDetailsInitScriptExecutionStatusFailedExecution,
		InitScriptExecutionDetailsInitScriptExecutionStatusFailedFetch,
		InitScriptExecutionDetailsInitScriptExecutionStatusFuseMountFailed,
		InitScriptExecutionDetailsInitScriptExecutionStatusNotExecuted,
		InitScriptExecutionDetailsInitScriptExecutionStatusSkipped,
		InitScriptExecutionDetailsInitScriptExecutionStatusSucceeded,
		InitScriptExecutionDetailsInitScriptExecutionStatusUnknown,
	}
}

// Type always returns InitScriptExecutionDetailsInitScriptExecutionStatus to satisfy [pflag.Value] interface
func (f *InitScriptExecutionDetailsInitScriptExecutionStatus) Type() string {
	return "InitScriptExecutionDetailsInitScriptExecutionStatus"
}

// Config for an individual init script Next ID: 11
type InitScriptInfo struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss *Adlsgen2Info `json:"abfss,omitempty"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File *LocalFileInfo `json:"file,omitempty"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *GcsStorageInfo `json:"gcs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 *S3StorageInfo `json:"s3,omitempty"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace *WorkspaceStorageInfo `json:"workspace,omitempty"`
}

type InitScriptInfoAndExecutionDetails struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss *Adlsgen2Info `json:"abfss,omitempty"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	ErrorMessage string `json:"error_message,omitempty"`
	// The number duration of the script execution in seconds
	ExecutionDurationSeconds int `json:"execution_duration_seconds,omitempty"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File *LocalFileInfo `json:"file,omitempty"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *GcsStorageInfo `json:"gcs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 *S3StorageInfo `json:"s3,omitempty"`
	// The current status of the script
	Status InitScriptExecutionDetailsInitScriptExecutionStatus `json:"status,omitempty"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace *WorkspaceStorageInfo `json:"workspace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InitScriptInfoAndExecutionDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InitScriptInfoAndExecutionDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to install.
	Libraries []Library `json:"libraries"`
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstancePoolAccessControlResponse struct {
	// All permissions.
	AllPermissions []InstancePoolPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Tags that are added by Databricks regardless of any ``custom_tags``,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
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
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes `json:"gcp_attributes,omitempty"`
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
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Current state of the instance pool.
	State InstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus `json:"status,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolAndStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolAndStats) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	Availability InstancePoolAwsAttributesAvailability `json:"availability,omitempty"`
	// Calculates the bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	SpotBidPricePercent int `json:"spot_bid_price_percent,omitempty"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolAwsAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolAwsAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The set of AWS availability types supported when setting up nodes for a
// cluster.
type InstancePoolAwsAttributesAvailability string

const InstancePoolAwsAttributesAvailabilityOnDemand InstancePoolAwsAttributesAvailability = `ON_DEMAND`

const InstancePoolAwsAttributesAvailabilitySpot InstancePoolAwsAttributesAvailability = `SPOT`

// String representation for [fmt.Print]
func (f *InstancePoolAwsAttributesAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolAwsAttributesAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND`, `SPOT`:
		*f = InstancePoolAwsAttributesAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND", "SPOT"`, v)
	}
}

// Values returns all possible values for InstancePoolAwsAttributesAvailability.
//
// There is no guarantee on the order of the values in the slice.
func (f *InstancePoolAwsAttributesAvailability) Values() []InstancePoolAwsAttributesAvailability {
	return []InstancePoolAwsAttributesAvailability{
		InstancePoolAwsAttributesAvailabilityOnDemand,
		InstancePoolAwsAttributesAvailabilitySpot,
	}
}

// Type always returns InstancePoolAwsAttributesAvailability to satisfy [pflag.Value] interface
func (f *InstancePoolAwsAttributesAvailability) Type() string {
	return "InstancePoolAwsAttributesAvailability"
}

// Attributes set during instance pool creation which are related to Azure.
type InstancePoolAzureAttributes struct {
	// Availability type used for the spot nodes.
	Availability InstancePoolAzureAttributesAvailability `json:"availability,omitempty"`
	// With variable pricing, you have option to set a max price, in US dollars
	// (USD) For example, the value 2 would be a max price of $2.00 USD per
	// hour. If you set the max price to be -1, the VM won't be evicted based on
	// price. The price for the VM will be the current price for spot or the
	// price for a standard VM, which ever is less, as long as there is capacity
	// and quota available.
	SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolAzureAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolAzureAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The set of Azure availability types supported when setting up nodes for a
// cluster.
type InstancePoolAzureAttributesAvailability string

const InstancePoolAzureAttributesAvailabilityOnDemandAzure InstancePoolAzureAttributesAvailability = `ON_DEMAND_AZURE`

const InstancePoolAzureAttributesAvailabilitySpotAzure InstancePoolAzureAttributesAvailability = `SPOT_AZURE`

// String representation for [fmt.Print]
func (f *InstancePoolAzureAttributesAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolAzureAttributesAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_AZURE`, `SPOT_AZURE`:
		*f = InstancePoolAzureAttributesAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_AZURE", "SPOT_AZURE"`, v)
	}
}

// Values returns all possible values for InstancePoolAzureAttributesAvailability.
//
// There is no guarantee on the order of the values in the slice.
func (f *InstancePoolAzureAttributesAvailability) Values() []InstancePoolAzureAttributesAvailability {
	return []InstancePoolAzureAttributesAvailability{
		InstancePoolAzureAttributesAvailabilityOnDemandAzure,
		InstancePoolAzureAttributesAvailabilitySpotAzure,
	}
}

// Type always returns InstancePoolAzureAttributesAvailability to satisfy [pflag.Value] interface
func (f *InstancePoolAzureAttributesAvailability) Type() string {
	return "InstancePoolAzureAttributesAvailability"
}

// Attributes set during instance pool creation which are related to GCP.
type InstancePoolGcpAttributes struct {
	GcpAvailability GcpAvailability `json:"gcp_availability,omitempty"`
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount int `json:"local_ssd_count,omitempty"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west1-a". The provided
	// availability zone must be in the same region as the Databricks workspace.
	// For example, "us-west1-a" is not a valid zone id if the Databricks
	// workspace resides in the "us-east1" region. This is an optional field at
	// instance pool creation, and if not specified, a default zone will be
	// used.
	//
	// This field can be one of the following: - "HA" => High availability,
	// spread nodes across availability zones for a Databricks deployment region
	// - A GCP availability zone => Pick One of the available zones for (machine
	// type + region) from https://cloud.google.com/compute/docs/regions-zones
	// (e.g. "us-west1-a").
	//
	// If empty, Databricks picks an availability zone to schedule the cluster
	// on.
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolGcpAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolGcpAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstancePoolPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type InstancePoolPermissionLevel string

const InstancePoolPermissionLevelCanAttachTo InstancePoolPermissionLevel = `CAN_ATTACH_TO`

const InstancePoolPermissionLevelCanManage InstancePoolPermissionLevel = `CAN_MANAGE`

// String representation for [fmt.Print]
func (f *InstancePoolPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_MANAGE`:
		*f = InstancePoolPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_MANAGE"`, v)
	}
}

// Values returns all possible values for InstancePoolPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *InstancePoolPermissionLevel) Values() []InstancePoolPermissionLevel {
	return []InstancePoolPermissionLevel{
		InstancePoolPermissionLevelCanAttachTo,
		InstancePoolPermissionLevelCanManage,
	}
}

// Type always returns InstancePoolPermissionLevel to satisfy [pflag.Value] interface
func (f *InstancePoolPermissionLevel) Type() string {
	return "InstancePoolPermissionLevel"
}

type InstancePoolPermissions struct {
	AccessControlList []InstancePoolAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstancePoolPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstancePoolPermissionsRequest struct {
	AccessControlList []InstancePoolAccessControlRequest `json:"access_control_list,omitempty"`
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" url:"-"`
}

// The state of a Cluster. The current allowable state transitions are as
// follows:
//
// - “ACTIVE“ -> “STOPPED“ - “ACTIVE“ -> “DELETED“ - “STOPPED“ ->
// “ACTIVE“ - “STOPPED“ -> “DELETED“
type InstancePoolState string

const InstancePoolStateActive InstancePoolState = `ACTIVE`

const InstancePoolStateDeleted InstancePoolState = `DELETED`

const InstancePoolStateStopped InstancePoolState = `STOPPED`

// String representation for [fmt.Print]
func (f *InstancePoolState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETED`, `STOPPED`:
		*f = InstancePoolState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETED", "STOPPED"`, v)
	}
}

// Values returns all possible values for InstancePoolState.
//
// There is no guarantee on the order of the values in the slice.
func (f *InstancePoolState) Values() []InstancePoolState {
	return []InstancePoolState{
		InstancePoolStateActive,
		InstancePoolStateDeleted,
		InstancePoolStateStopped,
	}
}

// Type always returns InstancePoolState to satisfy [pflag.Value] interface
func (f *InstancePoolState) Type() string {
	return "InstancePoolState"
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstancePoolStats) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstancePoolStats) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors []PendingInstanceError `json:"pending_instance_errors,omitempty"`
}

type InstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn string `json:"iam_role_arn,omitempty"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn string `json:"instance_profile_arn"`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *InstanceProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s InstanceProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The kind of compute described by this compute specification.
//
// Depending on `kind`, different validations and default values will be
// applied.
//
// Clusters with `kind = CLASSIC_PREVIEW` support the following fields, whereas
// clusters with no specified `kind` do not. *
// [is_single_node](/api/workspace/clusters/create#is_single_node) *
// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
// [data_security_mode](/api/workspace/clusters/create#data_security_mode) set
// to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
// `DATA_SECURITY_MODE_STANDARD`
//
// By using the [simple form], your clusters are automatically using `kind =
// CLASSIC_PREVIEW`.
//
// [simple form]: https://docs.databricks.com/compute/simple-form.html
type Kind string

const KindClassicPreview Kind = `CLASSIC_PREVIEW`

// String representation for [fmt.Print]
func (f *Kind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Kind) Set(v string) error {
	switch v {
	case `CLASSIC_PREVIEW`:
		*f = Kind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC_PREVIEW"`, v)
	}
}

// Values returns all possible values for Kind.
//
// There is no guarantee on the order of the values in the slice.
func (f *Kind) Values() []Kind {
	return []Kind{
		KindClassicPreview,
	}
}

// Type always returns Kind to satisfy [pflag.Value] interface
func (f *Kind) Type() string {
	return "Kind"
}

type Language string

const LanguagePython Language = `python`

const LanguageR Language = `r`

const LanguageScala Language = `scala`

const LanguageSql Language = `sql`

// String representation for [fmt.Print]
func (f *Language) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Language) Set(v string) error {
	switch v {
	case `python`, `r`, `scala`, `sql`:
		*f = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "python", "r", "scala", "sql"`, v)
	}
}

// Values returns all possible values for Language.
//
// There is no guarantee on the order of the values in the slice.
func (f *Language) Values() []Language {
	return []Language{
		LanguagePython,
		LanguageR,
		LanguageScala,
		LanguageSql,
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (f *Language) Type() string {
	return "Language"
}

type Library struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran *RCranLibrary `json:"cran,omitempty"`
	// Deprecated. URI of the egg library to install. Installing Python egg
	// files is deprecated and is not supported in Databricks Runtime 14.0 and
	// above.
	Egg string `json:"egg,omitempty"`
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	Maven *MavenLibrary `json:"maven,omitempty"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi *PythonPyPiLibrary `json:"pypi,omitempty"`
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	Requirements string `json:"requirements,omitempty"`
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Whl string `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Library) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Library) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters bool `json:"is_library_for_all_clusters,omitempty"`
	// Unique identifier for the library.
	Library *Library `json:"library,omitempty"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages []string `json:"messages,omitempty"`
	// Status of installing the library on the cluster.
	Status LibraryInstallStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LibraryFullStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LibraryFullStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The status of a library on a specific cluster.
type LibraryInstallStatus string

const LibraryInstallStatusFailed LibraryInstallStatus = `FAILED`

const LibraryInstallStatusInstalled LibraryInstallStatus = `INSTALLED`

const LibraryInstallStatusInstalling LibraryInstallStatus = `INSTALLING`

const LibraryInstallStatusPending LibraryInstallStatus = `PENDING`

const LibraryInstallStatusResolving LibraryInstallStatus = `RESOLVING`

const LibraryInstallStatusRestored LibraryInstallStatus = `RESTORED`

const LibraryInstallStatusSkipped LibraryInstallStatus = `SKIPPED`

const LibraryInstallStatusUninstallOnRestart LibraryInstallStatus = `UNINSTALL_ON_RESTART`

// String representation for [fmt.Print]
func (f *LibraryInstallStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LibraryInstallStatus) Set(v string) error {
	switch v {
	case `FAILED`, `INSTALLED`, `INSTALLING`, `PENDING`, `RESOLVING`, `RESTORED`, `SKIPPED`, `UNINSTALL_ON_RESTART`:
		*f = LibraryInstallStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "INSTALLED", "INSTALLING", "PENDING", "RESOLVING", "RESTORED", "SKIPPED", "UNINSTALL_ON_RESTART"`, v)
	}
}

// Values returns all possible values for LibraryInstallStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *LibraryInstallStatus) Values() []LibraryInstallStatus {
	return []LibraryInstallStatus{
		LibraryInstallStatusFailed,
		LibraryInstallStatusInstalled,
		LibraryInstallStatusInstalling,
		LibraryInstallStatusPending,
		LibraryInstallStatusResolving,
		LibraryInstallStatusRestored,
		LibraryInstallStatusSkipped,
		LibraryInstallStatusUninstallOnRestart,
	}
}

// Type always returns LibraryInstallStatus to satisfy [pflag.Value] interface
func (f *LibraryInstallStatus) Type() string {
	return "LibraryInstallStatus"
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses []ClusterLibraryStatuses `json:"statuses,omitempty"`
}

type ListAvailableZonesResponse struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	DefaultZone string `json:"default_zone,omitempty"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones []string `json:"zones,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAvailableZonesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAvailableZonesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListClusterCompliancesRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"-" url:"policy_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListClusterCompliancesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListClusterCompliancesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListClusterCompliancesResponse struct {
	// A list of clusters and their policy compliance statuses.
	Clusters []ClusterCompliance `json:"clusters,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListClusterCompliancesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListClusterCompliancesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn ListSortColumn `json:"-" url:"sort_column,omitempty"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder ListSortOrder `json:"-" url:"sort_order,omitempty"`
}

type ListClustersFilterBy struct {
	// The source of cluster creation.
	ClusterSources []ClusterSource `json:"cluster_sources,omitempty" url:"cluster_sources,omitempty"`
	// The current state of the clusters.
	ClusterStates []State `json:"cluster_states,omitempty" url:"cluster_states,omitempty"`
	// Whether the clusters are pinned or not.
	IsPinned bool `json:"is_pinned,omitempty" url:"is_pinned,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListClustersFilterBy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListClustersFilterBy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListClustersRequest struct {
	// Filters to apply to the list of clusters.
	FilterBy *ListClustersFilterBy `json:"-" url:"filter_by,omitempty"`
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Sort the list of clusters by a specific criteria.
	SortBy *ListClustersSortBy `json:"-" url:"sort_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListClustersRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListClustersRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListClustersResponse struct {
	Clusters []ClusterDetails `json:"clusters,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListClustersResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListClustersResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListClustersSortBy struct {
	// The direction to sort by.
	Direction ListClustersSortByDirection `json:"direction,omitempty" url:"direction,omitempty"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field ListClustersSortByField `json:"field,omitempty" url:"field,omitempty"`
}

type ListClustersSortByDirection string

const ListClustersSortByDirectionAsc ListClustersSortByDirection = `ASC`

const ListClustersSortByDirectionDesc ListClustersSortByDirection = `DESC`

// String representation for [fmt.Print]
func (f *ListClustersSortByDirection) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListClustersSortByDirection) Set(v string) error {
	switch v {
	case `ASC`, `DESC`:
		*f = ListClustersSortByDirection(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASC", "DESC"`, v)
	}
}

// Values returns all possible values for ListClustersSortByDirection.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListClustersSortByDirection) Values() []ListClustersSortByDirection {
	return []ListClustersSortByDirection{
		ListClustersSortByDirectionAsc,
		ListClustersSortByDirectionDesc,
	}
}

// Type always returns ListClustersSortByDirection to satisfy [pflag.Value] interface
func (f *ListClustersSortByDirection) Type() string {
	return "ListClustersSortByDirection"
}

type ListClustersSortByField string

const ListClustersSortByFieldClusterName ListClustersSortByField = `CLUSTER_NAME`

const ListClustersSortByFieldDefault ListClustersSortByField = `DEFAULT`

// String representation for [fmt.Print]
func (f *ListClustersSortByField) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListClustersSortByField) Set(v string) error {
	switch v {
	case `CLUSTER_NAME`, `DEFAULT`:
		*f = ListClustersSortByField(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLUSTER_NAME", "DEFAULT"`, v)
	}
}

// Values returns all possible values for ListClustersSortByField.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListClustersSortByField) Values() []ListClustersSortByField {
	return []ListClustersSortByField{
		ListClustersSortByFieldClusterName,
		ListClustersSortByFieldDefault,
	}
}

// Type always returns ListClustersSortByField to satisfy [pflag.Value] interface
func (f *ListClustersSortByField) Type() string {
	return "ListClustersSortByField"
}

type ListGlobalInitScriptsResponse struct {
	Scripts []GlobalInitScriptDetails `json:"scripts,omitempty"`
}

type ListInstancePools struct {
	InstancePools []InstancePoolAndStats `json:"instance_pools,omitempty"`
}

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles []InstanceProfile `json:"instance_profiles,omitempty"`
}

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	NodeTypes []NodeType `json:"node_types,omitempty"`
}

type ListPoliciesResponse struct {
	// List of policies.
	Policies []Policy `json:"policies,omitempty"`
}

type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPolicyFamiliesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPolicyFamiliesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of policy families.
	PolicyFamilies []PolicyFamily `json:"policy_families,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPolicyFamiliesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPolicyFamiliesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSortColumn string

const ListSortColumnPolicyCreationTime ListSortColumn = `POLICY_CREATION_TIME`

const ListSortColumnPolicyName ListSortColumn = `POLICY_NAME`

// String representation for [fmt.Print]
func (f *ListSortColumn) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListSortColumn) Set(v string) error {
	switch v {
	case `POLICY_CREATION_TIME`, `POLICY_NAME`:
		*f = ListSortColumn(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "POLICY_CREATION_TIME", "POLICY_NAME"`, v)
	}
}

// Values returns all possible values for ListSortColumn.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListSortColumn) Values() []ListSortColumn {
	return []ListSortColumn{
		ListSortColumnPolicyCreationTime,
		ListSortColumnPolicyName,
	}
}

// Type always returns ListSortColumn to satisfy [pflag.Value] interface
func (f *ListSortColumn) Type() string {
	return "ListSortColumn"
}

type ListSortOrder string

const ListSortOrderAsc ListSortOrder = `ASC`

const ListSortOrderDesc ListSortOrder = `DESC`

// String representation for [fmt.Print]
func (f *ListSortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListSortOrder) Set(v string) error {
	switch v {
	case `ASC`, `DESC`:
		*f = ListSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASC", "DESC"`, v)
	}
}

// Values returns all possible values for ListSortOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListSortOrder) Values() []ListSortOrder {
	return []ListSortOrder{
		ListSortOrderAsc,
		ListSortOrderDesc,
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (f *ListSortOrder) Type() string {
	return "ListSortOrder"
}

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination string `json:"destination"`
}

type LogAnalyticsInfo struct {
	LogAnalyticsPrimaryKey string `json:"log_analytics_primary_key,omitempty"`

	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogAnalyticsInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogAnalyticsInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The log delivery status
type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted int64 `json:"last_attempted,omitempty"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException string `json:"last_exception,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *LogSyncStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogSyncStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type MapAny map[string]any

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	Coordinates string `json:"coordinates"`
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions []string `json:"exclusions,omitempty"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MavenLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MavenLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This structure embodies the machine type that hosts spark containers Note:
// this should be an internal data structure for now It is defined in proto in
// case we want to send it over the wire in the future (which is likely)
type NodeInstanceType struct {
	// Unique identifier across instance types
	InstanceTypeId string `json:"instance_type_id"`
	// Size of the individual local disks attached to this instance (i.e. per
	// local disk).
	LocalDiskSizeGb int `json:"local_disk_size_gb,omitempty"`
	// Number of local disks that are present on this instance.
	LocalDisks int `json:"local_disks,omitempty"`
	// Size of the individual local nvme disks attached to this instance (i.e.
	// per local disk).
	LocalNvmeDiskSizeGb int `json:"local_nvme_disk_size_gb,omitempty"`
	// Number of local nvme disks that are present on this instance.
	LocalNvmeDisks int `json:"local_nvme_disks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NodeInstanceType) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NodeInstanceType) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A description of a Spark node type including both the dimensions of the node
// and the instance type on which it will be hosted.
type NodeType struct {
	// A descriptive category for this node type. Examples include "Memory
	// Optimized" and "Compute Optimized".
	Category string `json:"category"`
	// A string description associated with this node type, e.g., "r3.xlarge".
	Description string `json:"description"`
	// An optional hint at the display order of node types in the UI. Within a
	// node type category, lowest numbers come first.
	DisplayOrder int `json:"display_order,omitempty"`
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	InstanceTypeId string `json:"instance_type_id"`
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	IsDeprecated bool `json:"is_deprecated,omitempty"`
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	IsEncryptedInTransit bool `json:"is_encrypted_in_transit,omitempty"`
	// Whether this is an Arm-based instance.
	IsGraviton bool `json:"is_graviton,omitempty"`
	// Whether this node is hidden from presentation in the UI.
	IsHidden bool `json:"is_hidden,omitempty"`
	// Whether this node comes with IO cache enabled by default.
	IsIoCacheEnabled bool `json:"is_io_cache_enabled,omitempty"`
	// Memory (in MB) available for this node type.
	MemoryMb int `json:"memory_mb"`
	// A collection of node type info reported by the cloud provider
	NodeInfo *CloudProviderNodeInfo `json:"node_info,omitempty"`
	// The NodeInstanceType object corresponding to instance_type_id
	NodeInstanceType *NodeInstanceType `json:"node_instance_type,omitempty"`
	// Unique identifier for this node type.
	NodeTypeId string `json:"node_type_id"`
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	NumCores float64 `json:"num_cores"`
	// Number of GPUs available for this node type.
	NumGpus int `json:"num_gpus,omitempty"`

	PhotonDriverCapable bool `json:"photon_driver_capable,omitempty"`

	PhotonWorkerCapable bool `json:"photon_worker_capable,omitempty"`
	// Whether this node type support cluster tags.
	SupportClusterTags bool `json:"support_cluster_tags,omitempty"`
	// Whether this node type support EBS volumes. EBS volumes is disabled for
	// node types that we could place multiple corresponding containers on the
	// same hosting instance.
	SupportEbsVolumes bool `json:"support_ebs_volumes,omitempty"`
	// Whether this node type supports port forwarding.
	SupportPortForwarding bool `json:"support_port_forwarding,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NodeType) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NodeType) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Error message of a failed pending instances
type PendingInstanceError struct {
	InstanceId string `json:"instance_id,omitempty"`

	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PendingInstanceError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PendingInstanceError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PermanentDeleteCluster struct {
	// The cluster to be deleted.
	ClusterId string `json:"cluster_id"`
}

type PinCluster struct {
	ClusterId string `json:"cluster_id"`
}

// Describes a Cluster Policy entity.
type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	CreatedAtTimestamp int64 `json:"created_at_timestamp,omitempty"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	IsDefault bool `json:"is_default,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library `json:"libraries,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string `json:"name,omitempty"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Policy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Policy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string `json:"definition,omitempty"`
	// Human-readable description of the purpose of the policy family.
	Description string `json:"description,omitempty"`
	// Name of the policy family.
	Name string `json:"name,omitempty"`
	// Unique identifier for the policy family.
	PolicyFamilyId string `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PolicyFamily) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PolicyFamily) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PythonPyPiLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PythonPyPiLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RCranLibrary) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RCranLibrary) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	InstanceProfileArn string `json:"instance_profile_arn"`
}

type ResizeCluster struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// The cluster to be resized.
	ClusterId string `json:"cluster_id"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ResizeCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResizeCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RestartCluster struct {
	// The cluster to be started.
	ClusterId string `json:"cluster_id"`

	RestartUser string `json:"restart_user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestartCluster) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestartCluster) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ResultType string

const ResultTypeError ResultType = `error`

const ResultTypeImage ResultType = `image`

const ResultTypeImages ResultType = `images`

const ResultTypeTable ResultType = `table`

const ResultTypeText ResultType = `text`

// String representation for [fmt.Print]
func (f *ResultType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ResultType) Set(v string) error {
	switch v {
	case `error`, `image`, `images`, `table`, `text`:
		*f = ResultType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "error", "image", "images", "table", "text"`, v)
	}
}

// Values returns all possible values for ResultType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ResultType) Values() []ResultType {
	return []ResultType{
		ResultTypeError,
		ResultTypeImage,
		ResultTypeImages,
		ResultTypeTable,
		ResultTypeText,
	}
}

// Type always returns ResultType to satisfy [pflag.Value] interface
func (f *ResultType) Type() string {
	return "ResultType"
}

type Results struct {
	// The cause of the error
	Cause string `json:"cause,omitempty"`

	Data any `json:"data,omitempty"`
	// The image filename
	FileName string `json:"fileName,omitempty"`

	FileNames []string `json:"fileNames,omitempty"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema bool `json:"isJsonSchema,omitempty"`
	// internal field used by SDK
	Pos int `json:"pos,omitempty"`

	ResultType ResultType `json:"resultType,omitempty"`
	// The table schema
	Schema []map[string]any `json:"schema,omitempty"`
	// The summary of the error
	Summary string `json:"summary,omitempty"`
	// true if partial results are returned.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Results) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Results) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RuntimeEngine string

const RuntimeEngineNull RuntimeEngine = `NULL`

const RuntimeEnginePhoton RuntimeEngine = `PHOTON`

const RuntimeEngineStandard RuntimeEngine = `STANDARD`

// String representation for [fmt.Print]
func (f *RuntimeEngine) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RuntimeEngine) Set(v string) error {
	switch v {
	case `NULL`, `PHOTON`, `STANDARD`:
		*f = RuntimeEngine(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NULL", "PHOTON", "STANDARD"`, v)
	}
}

// Values returns all possible values for RuntimeEngine.
//
// There is no guarantee on the order of the values in the slice.
func (f *RuntimeEngine) Values() []RuntimeEngine {
	return []RuntimeEngine{
		RuntimeEngineNull,
		RuntimeEnginePhoton,
		RuntimeEngineStandard,
	}
}

// Type always returns RuntimeEngine to satisfy [pflag.Value] interface
func (f *RuntimeEngine) Type() string {
	return "RuntimeEngine"
}

// A storage location in Amazon S3
type S3StorageInfo struct {
	// (Optional) Set canned access control list for the logs, e.g.
	// `bucket-owner-full-control`. If `canned_cal` is set, please make sure the
	// cluster iam role has `s3:PutObjectAcl` permission on the destination
	// bucket and prefix. The full list of possible canned acl can be found at
	// http://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl.
	// Please also note that by default only the object owner gets full
	// controls. If you are using cross account role for writing data, you may
	// want to set `bucket-owner-full-control` to make bucket owner able to read
	// the logs.
	CannedAcl string `json:"canned_acl,omitempty"`
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	Destination string `json:"destination"`
	// (Optional) Flag to enable server side encryption, `false` by default.
	EnableEncryption bool `json:"enable_encryption,omitempty"`
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	EncryptionType string `json:"encryption_type,omitempty"`
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	Endpoint string `json:"endpoint,omitempty"`
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	KmsKey string `json:"kms_key,omitempty"`
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *S3StorageInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s S3StorageInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Describes a specific Spark driver or executor.
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
	// rules to the "worker-unmanaged" security group via the AWS console.
	PublicDns string `json:"public_dns,omitempty"`
	// The timestamp (in millisecond) when the Spark node is launched.
	StartTimestamp int64 `json:"start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SparkNode) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SparkNode) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot bool `json:"is_spot,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SparkNodeAwsAttributes) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SparkNodeAwsAttributes) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SparkVersion struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	Key string `json:"key,omitempty"`
	// A descriptive name for this Spark version, for example "Spark 2.1".
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SparkVersion) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SparkVersion) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StartCluster struct {
	// The cluster to be started.
	ClusterId string `json:"cluster_id"`
}

// The state of a Cluster. The current allowable state transitions are as
// follows:
//
// - `PENDING` -> `RUNNING` - `PENDING` -> `TERMINATING` - `RUNNING` ->
// `RESIZING` - `RUNNING` -> `RESTARTING` - `RUNNING` -> `TERMINATING` -
// `RESTARTING` -> `RUNNING` - `RESTARTING` -> `TERMINATING` - `RESIZING` ->
// `RUNNING` - `RESIZING` -> `TERMINATING` - `TERMINATING` -> `TERMINATED`
type State string

const StateError State = `ERROR`

const StatePending State = `PENDING`

const StateResizing State = `RESIZING`

const StateRestarting State = `RESTARTING`

const StateRunning State = `RUNNING`

const StateTerminated State = `TERMINATED`

const StateTerminating State = `TERMINATING`

const StateUnknown State = `UNKNOWN`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `ERROR`, `PENDING`, `RESIZING`, `RESTARTING`, `RUNNING`, `TERMINATED`, `TERMINATING`, `UNKNOWN`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "PENDING", "RESIZING", "RESTARTING", "RUNNING", "TERMINATED", "TERMINATING", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateError,
		StatePending,
		StateResizing,
		StateRestarting,
		StateRunning,
		StateTerminated,
		StateTerminating,
		StateUnknown,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
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

// The status code indicating why the cluster was terminated
type TerminationReasonCode string

const TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`

const TerminationReasonCodeAccessTokenFailure TerminationReasonCode = `ACCESS_TOKEN_FAILURE`

const TerminationReasonCodeAllocationTimeout TerminationReasonCode = `ALLOCATION_TIMEOUT`

const TerminationReasonCodeAllocationTimeoutNodeDaemonNotReady TerminationReasonCode = `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`

const TerminationReasonCodeAllocationTimeoutNoHealthyAndWarmedUpClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoHealthyClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoMatchedClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoReadyClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoUnallocatedClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`

const TerminationReasonCodeAllocationTimeoutNoWarmedUpClusters TerminationReasonCode = `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInaccessibleKmsKeyFailure TerminationReasonCode = `AWS_INACCESSIBLE_KMS_KEY_FAILURE`

const TerminationReasonCodeAwsInstanceProfileUpdateFailure TerminationReasonCode = `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsInvalidKeyPair TerminationReasonCode = `AWS_INVALID_KEY_PAIR`

const TerminationReasonCodeAwsInvalidKmsKeyState TerminationReasonCode = `AWS_INVALID_KMS_KEY_STATE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsResourceQuotaExceeded TerminationReasonCode = `AWS_RESOURCE_QUOTA_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzurePackedDeploymentPartialFailure TerminationReasonCode = `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeBootstrapTimeoutDueToMisconfig TerminationReasonCode = `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`

const TerminationReasonCodeBudgetPolicyLimitEnforcementActivated TerminationReasonCode = `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`

const TerminationReasonCodeBudgetPolicyResolutionFailure TerminationReasonCode = `BUDGET_POLICY_RESOLUTION_FAILURE`

const TerminationReasonCodeCloudAccountSetupFailure TerminationReasonCode = `CLOUD_ACCOUNT_SETUP_FAILURE`

const TerminationReasonCodeCloudOperationCancelled TerminationReasonCode = `CLOUD_OPERATION_CANCELLED`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderInstanceNotLaunched TerminationReasonCode = `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailureDueToMisconfig TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderResourceStockoutDueToMisconfig TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeClusterOperationThrottled TerminationReasonCode = `CLUSTER_OPERATION_THROTTLED`

const TerminationReasonCodeClusterOperationTimeout TerminationReasonCode = `CLUSTER_OPERATION_TIMEOUT`

const TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailureDueToMisconfig TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDataAccessConfigChanged TerminationReasonCode = `DATA_ACCESS_CONFIG_CHANGED`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDisasterRecoveryReplication TerminationReasonCode = `DISASTER_RECOVERY_REPLICATION`

const TerminationReasonCodeDnsResolutionError TerminationReasonCode = `DNS_RESOLUTION_ERROR`

const TerminationReasonCodeDockerContainerCreationException TerminationReasonCode = `DOCKER_CONTAINER_CREATION_EXCEPTION`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDockerImageTooLargeForInstanceException TerminationReasonCode = `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`

const TerminationReasonCodeDockerInvalidOsException TerminationReasonCode = `DOCKER_INVALID_OS_EXCEPTION`

const TerminationReasonCodeDriverDnsResolutionFailure TerminationReasonCode = `DRIVER_DNS_RESOLUTION_FAILURE`

const TerminationReasonCodeDriverEviction TerminationReasonCode = `DRIVER_EVICTION`

const TerminationReasonCodeDriverLaunchTimeout TerminationReasonCode = `DRIVER_LAUNCH_TIMEOUT`

const TerminationReasonCodeDriverNodeUnreachable TerminationReasonCode = `DRIVER_NODE_UNREACHABLE`

const TerminationReasonCodeDriverOutOfDisk TerminationReasonCode = `DRIVER_OUT_OF_DISK`

const TerminationReasonCodeDriverOutOfMemory TerminationReasonCode = `DRIVER_OUT_OF_MEMORY`

const TerminationReasonCodeDriverPodCreationFailure TerminationReasonCode = `DRIVER_POD_CREATION_FAILURE`

const TerminationReasonCodeDriverUnexpectedFailure TerminationReasonCode = `DRIVER_UNEXPECTED_FAILURE`

const TerminationReasonCodeDriverUnhealthy TerminationReasonCode = `DRIVER_UNHEALTHY`

const TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeDynamicSparkConfSizeExceeded TerminationReasonCode = `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`

const TerminationReasonCodeEosSparkImage TerminationReasonCode = `EOS_SPARK_IMAGE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeExecutorPodUnscheduled TerminationReasonCode = `EXECUTOR_POD_UNSCHEDULED`

const TerminationReasonCodeGcpApiRateQuotaExceeded TerminationReasonCode = `GCP_API_RATE_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpDeniedByOrgPolicy TerminationReasonCode = `GCP_DENIED_BY_ORG_POLICY`

const TerminationReasonCodeGcpForbidden TerminationReasonCode = `GCP_FORBIDDEN`

const TerminationReasonCodeGcpIamTimeout TerminationReasonCode = `GCP_IAM_TIMEOUT`

const TerminationReasonCodeGcpInaccessibleKmsKeyFailure TerminationReasonCode = `GCP_INACCESSIBLE_KMS_KEY_FAILURE`

const TerminationReasonCodeGcpInsufficientCapacity TerminationReasonCode = `GCP_INSUFFICIENT_CAPACITY`

const TerminationReasonCodeGcpIpSpaceExhausted TerminationReasonCode = `GCP_IP_SPACE_EXHAUSTED`

const TerminationReasonCodeGcpKmsKeyPermissionDenied TerminationReasonCode = `GCP_KMS_KEY_PERMISSION_DENIED`

const TerminationReasonCodeGcpNotFound TerminationReasonCode = `GCP_NOT_FOUND`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpResourceQuotaExceeded TerminationReasonCode = `GCP_RESOURCE_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountAccessDenied TerminationReasonCode = `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGcpServiceAccountNotFound TerminationReasonCode = `GCP_SERVICE_ACCOUNT_NOT_FOUND`

const TerminationReasonCodeGcpSubnetNotReady TerminationReasonCode = `GCP_SUBNET_NOT_READY`

const TerminationReasonCodeGcpTrustedImageProjectsViolated TerminationReasonCode = `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`

const TerminationReasonCodeGkeBasedClusterTermination TerminationReasonCode = `GKE_BASED_CLUSTER_TERMINATION`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`

const TerminationReasonCodeInitContainerNotFinished TerminationReasonCode = `INIT_CONTAINER_NOT_FINISHED`

const TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstancePoolMaxCapacityReached TerminationReasonCode = `INSTANCE_POOL_MAX_CAPACITY_REACHED`

const TerminationReasonCodeInstancePoolNotFound TerminationReasonCode = `INSTANCE_POOL_NOT_FOUND`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInstanceUnreachableDueToMisconfig TerminationReasonCode = `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`

const TerminationReasonCodeInternalCapacityFailure TerminationReasonCode = `INTERNAL_CAPACITY_FAILURE`

const TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidAwsParameter TerminationReasonCode = `INVALID_AWS_PARAMETER`

const TerminationReasonCodeInvalidInstancePlacementProtocol TerminationReasonCode = `INVALID_INSTANCE_PLACEMENT_PROTOCOL`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeInvalidWorkerImageFailure TerminationReasonCode = `INVALID_WORKER_IMAGE_FAILURE`

const TerminationReasonCodeInPenaltyBox TerminationReasonCode = `IN_PENALTY_BOX`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeLazyAllocationTimeout TerminationReasonCode = `LAZY_ALLOCATION_TIMEOUT`

const TerminationReasonCodeMaintenanceMode TerminationReasonCode = `MAINTENANCE_MODE`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetvisorSetupTimeout TerminationReasonCode = `NETVISOR_SETUP_TIMEOUT`

const TerminationReasonCodeNetworkCheckControlPlaneFailure TerminationReasonCode = `NETWORK_CHECK_CONTROL_PLANE_FAILURE`

const TerminationReasonCodeNetworkCheckDnsServerFailure TerminationReasonCode = `NETWORK_CHECK_DNS_SERVER_FAILURE`

const TerminationReasonCodeNetworkCheckMetadataEndpointFailure TerminationReasonCode = `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE`

const TerminationReasonCodeNetworkCheckMultipleComponentsFailure TerminationReasonCode = `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE`

const TerminationReasonCodeNetworkCheckNicFailure TerminationReasonCode = `NETWORK_CHECK_NIC_FAILURE`

const TerminationReasonCodeNetworkCheckStorageFailure TerminationReasonCode = `NETWORK_CHECK_STORAGE_FAILURE`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNoMatchedK8s TerminationReasonCode = `NO_MATCHED_K8S`

const TerminationReasonCodeNoMatchedK8sTestingTag TerminationReasonCode = `NO_MATCHED_K8S_TESTING_TAG`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodePodAssignmentFailure TerminationReasonCode = `POD_ASSIGNMENT_FAILURE`

const TerminationReasonCodePodSchedulingFailure TerminationReasonCode = `POD_SCHEDULING_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`

const TerminationReasonCodeResourceUsageBlocked TerminationReasonCode = `RESOURCE_USAGE_BLOCKED`

const TerminationReasonCodeSecretCreationFailure TerminationReasonCode = `SECRET_CREATION_FAILURE`

const TerminationReasonCodeSecretPermissionDenied TerminationReasonCode = `SECRET_PERMISSION_DENIED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityAgentsFailedInitialVerification TerminationReasonCode = `SECURITY_AGENTS_FAILED_INITIAL_VERIFICATION`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeServerlessLongRunningTerminated TerminationReasonCode = `SERVERLESS_LONG_RUNNING_TERMINATED`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkImageDownloadThrottled TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_THROTTLED`

const TerminationReasonCodeSparkImageNotFound TerminationReasonCode = `SPARK_IMAGE_NOT_FOUND`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeSshBootstrapFailure TerminationReasonCode = `SSH_BOOTSTRAP_FAILURE`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStorageDownloadFailureDueToMisconfig TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`

const TerminationReasonCodeStorageDownloadFailureSlow TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE_SLOW`

const TerminationReasonCodeStorageDownloadFailureThrottled TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE_THROTTLED`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnexpectedPodRecreation TerminationReasonCode = `UNEXPECTED_POD_RECREATION`

const TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserInitiatedVmTermination TerminationReasonCode = `USER_INITIATED_VM_TERMINATION`

const TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`

const TerminationReasonCodeWorkspaceUpdate TerminationReasonCode = `WORKSPACE_UPDATE`

// String representation for [fmt.Print]
func (f *TerminationReasonCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonCode) Set(v string) error {
	switch v {
	case `ABUSE_DETECTED`, `ACCESS_TOKEN_FAILURE`, `ALLOCATION_TIMEOUT`, `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`, `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INACCESSIBLE_KMS_KEY_FAILURE`, `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_INVALID_KEY_PAIR`, `AWS_INVALID_KMS_KEY_STATE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_RESOURCE_QUOTA_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`, `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`, `BUDGET_POLICY_RESOLUTION_FAILURE`, `CLOUD_ACCOUNT_SETUP_FAILURE`, `CLOUD_OPERATION_CANCELLED`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_SHUTDOWN`, `CLUSTER_OPERATION_THROTTLED`, `CLUSTER_OPERATION_TIMEOUT`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`, `DATABASE_CONNECTION_FAILURE`, `DATA_ACCESS_CONFIG_CHANGED`, `DBFS_COMPONENT_UNHEALTHY`, `DISASTER_RECOVERY_REPLICATION`, `DNS_RESOLUTION_ERROR`, `DOCKER_CONTAINER_CREATION_EXCEPTION`, `DOCKER_IMAGE_PULL_FAILURE`, `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`, `DOCKER_INVALID_OS_EXCEPTION`, `DRIVER_DNS_RESOLUTION_FAILURE`, `DRIVER_EVICTION`, `DRIVER_LAUNCH_TIMEOUT`, `DRIVER_NODE_UNREACHABLE`, `DRIVER_OUT_OF_DISK`, `DRIVER_OUT_OF_MEMORY`, `DRIVER_POD_CREATION_FAILURE`, `DRIVER_UNEXPECTED_FAILURE`, `DRIVER_UNHEALTHY`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`, `EOS_SPARK_IMAGE`, `EXECUTION_COMPONENT_UNHEALTHY`, `EXECUTOR_POD_UNSCHEDULED`, `GCP_API_RATE_QUOTA_EXCEEDED`, `GCP_DENIED_BY_ORG_POLICY`, `GCP_FORBIDDEN`, `GCP_IAM_TIMEOUT`, `GCP_INACCESSIBLE_KMS_KEY_FAILURE`, `GCP_INSUFFICIENT_CAPACITY`, `GCP_IP_SPACE_EXHAUSTED`, `GCP_KMS_KEY_PERMISSION_DENIED`, `GCP_NOT_FOUND`, `GCP_QUOTA_EXCEEDED`, `GCP_RESOURCE_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GCP_SERVICE_ACCOUNT_NOT_FOUND`, `GCP_SUBNET_NOT_READY`, `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`, `GKE_BASED_CLUSTER_TERMINATION`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_CONTAINER_NOT_FINISHED`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_POOL_MAX_CAPACITY_REACHED`, `INSTANCE_POOL_NOT_FOUND`, `INSTANCE_UNREACHABLE`, `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`, `INTERNAL_CAPACITY_FAILURE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_AWS_PARAMETER`, `INVALID_INSTANCE_PLACEMENT_PROTOCOL`, `INVALID_SPARK_IMAGE`, `INVALID_WORKER_IMAGE_FAILURE`, `IN_PENALTY_BOX`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `LAZY_ALLOCATION_TIMEOUT`, `MAINTENANCE_MODE`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETVISOR_SETUP_TIMEOUT`, `NETWORK_CHECK_CONTROL_PLANE_FAILURE`, `NETWORK_CHECK_DNS_SERVER_FAILURE`, `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE`, `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE`, `NETWORK_CHECK_NIC_FAILURE`, `NETWORK_CHECK_STORAGE_FAILURE`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NO_MATCHED_K8S`, `NO_MATCHED_K8S_TESTING_TAG`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `POD_ASSIGNMENT_FAILURE`, `POD_SCHEDULING_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `RESOURCE_USAGE_BLOCKED`, `SECRET_CREATION_FAILURE`, `SECRET_PERMISSION_DENIED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_AGENTS_FAILED_INITIAL_VERIFICATION`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SERVERLESS_LONG_RUNNING_TERMINATED`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_IMAGE_DOWNLOAD_THROTTLED`, `SPARK_IMAGE_NOT_FOUND`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `SSH_BOOTSTRAP_FAILURE`, `STORAGE_DOWNLOAD_FAILURE`, `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`, `STORAGE_DOWNLOAD_FAILURE_SLOW`, `STORAGE_DOWNLOAD_FAILURE_THROTTLED`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNEXPECTED_POD_RECREATION`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_INITIATED_VM_TERMINATION`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`, `WORKSPACE_UPDATE`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ACCESS_TOKEN_FAILURE", "ALLOCATION_TIMEOUT", "ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY", "ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS", "ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_READY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INACCESSIBLE_KMS_KEY_FAILURE", "AWS_INSTANCE_PROFILE_UPDATE_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_INVALID_KEY_PAIR", "AWS_INVALID_KMS_KEY_STATE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_RESOURCE_QUOTA_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG", "BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED", "BUDGET_POLICY_RESOLUTION_FAILURE", "CLOUD_ACCOUNT_SETUP_FAILURE", "CLOUD_OPERATION_CANCELLED", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_SHUTDOWN", "CLUSTER_OPERATION_THROTTLED", "CLUSTER_OPERATION_TIMEOUT", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG", "DATABASE_CONNECTION_FAILURE", "DATA_ACCESS_CONFIG_CHANGED", "DBFS_COMPONENT_UNHEALTHY", "DISASTER_RECOVERY_REPLICATION", "DNS_RESOLUTION_ERROR", "DOCKER_CONTAINER_CREATION_EXCEPTION", "DOCKER_IMAGE_PULL_FAILURE", "DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION", "DOCKER_INVALID_OS_EXCEPTION", "DRIVER_DNS_RESOLUTION_FAILURE", "DRIVER_EVICTION", "DRIVER_LAUNCH_TIMEOUT", "DRIVER_NODE_UNREACHABLE", "DRIVER_OUT_OF_DISK", "DRIVER_OUT_OF_MEMORY", "DRIVER_POD_CREATION_FAILURE", "DRIVER_UNEXPECTED_FAILURE", "DRIVER_UNHEALTHY", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "DYNAMIC_SPARK_CONF_SIZE_EXCEEDED", "EOS_SPARK_IMAGE", "EXECUTION_COMPONENT_UNHEALTHY", "EXECUTOR_POD_UNSCHEDULED", "GCP_API_RATE_QUOTA_EXCEEDED", "GCP_DENIED_BY_ORG_POLICY", "GCP_FORBIDDEN", "GCP_IAM_TIMEOUT", "GCP_INACCESSIBLE_KMS_KEY_FAILURE", "GCP_INSUFFICIENT_CAPACITY", "GCP_IP_SPACE_EXHAUSTED", "GCP_KMS_KEY_PERMISSION_DENIED", "GCP_NOT_FOUND", "GCP_QUOTA_EXCEEDED", "GCP_RESOURCE_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_ACCESS_DENIED", "GCP_SERVICE_ACCOUNT_DELETED", "GCP_SERVICE_ACCOUNT_NOT_FOUND", "GCP_SUBNET_NOT_READY", "GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED", "GKE_BASED_CLUSTER_TERMINATION", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_CONTAINER_NOT_FINISHED", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_POOL_MAX_CAPACITY_REACHED", "INSTANCE_POOL_NOT_FOUND", "INSTANCE_UNREACHABLE", "INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG", "INTERNAL_CAPACITY_FAILURE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_AWS_PARAMETER", "INVALID_INSTANCE_PLACEMENT_PROTOCOL", "INVALID_SPARK_IMAGE", "INVALID_WORKER_IMAGE_FAILURE", "IN_PENALTY_BOX", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "LAZY_ALLOCATION_TIMEOUT", "MAINTENANCE_MODE", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETVISOR_SETUP_TIMEOUT", "NETWORK_CHECK_CONTROL_PLANE_FAILURE", "NETWORK_CHECK_DNS_SERVER_FAILURE", "NETWORK_CHECK_METADATA_ENDPOINT_FAILURE", "NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE", "NETWORK_CHECK_NIC_FAILURE", "NETWORK_CHECK_STORAGE_FAILURE", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NO_MATCHED_K8S", "NO_MATCHED_K8S_TESTING_TAG", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "POD_ASSIGNMENT_FAILURE", "POD_SCHEDULING_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "RESOURCE_USAGE_BLOCKED", "SECRET_CREATION_FAILURE", "SECRET_PERMISSION_DENIED", "SECRET_RESOLUTION_ERROR", "SECURITY_AGENTS_FAILED_INITIAL_VERIFICATION", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SERVERLESS_LONG_RUNNING_TERMINATED", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_IMAGE_DOWNLOAD_THROTTLED", "SPARK_IMAGE_NOT_FOUND", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "SSH_BOOTSTRAP_FAILURE", "STORAGE_DOWNLOAD_FAILURE", "STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG", "STORAGE_DOWNLOAD_FAILURE_SLOW", "STORAGE_DOWNLOAD_FAILURE_THROTTLED", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNEXPECTED_POD_RECREATION", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_INITIATED_VM_TERMINATION", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR", "WORKSPACE_UPDATE"`, v)
	}
}

// Values returns all possible values for TerminationReasonCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationReasonCode) Values() []TerminationReasonCode {
	return []TerminationReasonCode{
		TerminationReasonCodeAbuseDetected,
		TerminationReasonCodeAccessTokenFailure,
		TerminationReasonCodeAllocationTimeout,
		TerminationReasonCodeAllocationTimeoutNodeDaemonNotReady,
		TerminationReasonCodeAllocationTimeoutNoHealthyAndWarmedUpClusters,
		TerminationReasonCodeAllocationTimeoutNoHealthyClusters,
		TerminationReasonCodeAllocationTimeoutNoMatchedClusters,
		TerminationReasonCodeAllocationTimeoutNoReadyClusters,
		TerminationReasonCodeAllocationTimeoutNoUnallocatedClusters,
		TerminationReasonCodeAllocationTimeoutNoWarmedUpClusters,
		TerminationReasonCodeAttachProjectFailure,
		TerminationReasonCodeAwsAuthorizationFailure,
		TerminationReasonCodeAwsInaccessibleKmsKeyFailure,
		TerminationReasonCodeAwsInstanceProfileUpdateFailure,
		TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure,
		TerminationReasonCodeAwsInsufficientInstanceCapacityFailure,
		TerminationReasonCodeAwsInvalidKeyPair,
		TerminationReasonCodeAwsInvalidKmsKeyState,
		TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure,
		TerminationReasonCodeAwsRequestLimitExceeded,
		TerminationReasonCodeAwsResourceQuotaExceeded,
		TerminationReasonCodeAwsUnsupportedFailure,
		TerminationReasonCodeAzureByokKeyPermissionFailure,
		TerminationReasonCodeAzureEphemeralDiskFailure,
		TerminationReasonCodeAzureInvalidDeploymentTemplate,
		TerminationReasonCodeAzureOperationNotAllowedException,
		TerminationReasonCodeAzurePackedDeploymentPartialFailure,
		TerminationReasonCodeAzureQuotaExceededException,
		TerminationReasonCodeAzureResourceManagerThrottling,
		TerminationReasonCodeAzureResourceProviderThrottling,
		TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure,
		TerminationReasonCodeAzureVmExtensionFailure,
		TerminationReasonCodeAzureVnetConfigurationFailure,
		TerminationReasonCodeBootstrapTimeout,
		TerminationReasonCodeBootstrapTimeoutCloudProviderException,
		TerminationReasonCodeBootstrapTimeoutDueToMisconfig,
		TerminationReasonCodeBudgetPolicyLimitEnforcementActivated,
		TerminationReasonCodeBudgetPolicyResolutionFailure,
		TerminationReasonCodeCloudAccountSetupFailure,
		TerminationReasonCodeCloudOperationCancelled,
		TerminationReasonCodeCloudProviderDiskSetupFailure,
		TerminationReasonCodeCloudProviderInstanceNotLaunched,
		TerminationReasonCodeCloudProviderLaunchFailure,
		TerminationReasonCodeCloudProviderLaunchFailureDueToMisconfig,
		TerminationReasonCodeCloudProviderResourceStockout,
		TerminationReasonCodeCloudProviderResourceStockoutDueToMisconfig,
		TerminationReasonCodeCloudProviderShutdown,
		TerminationReasonCodeClusterOperationThrottled,
		TerminationReasonCodeClusterOperationTimeout,
		TerminationReasonCodeCommunicationLost,
		TerminationReasonCodeContainerLaunchFailure,
		TerminationReasonCodeControlPlaneRequestFailure,
		TerminationReasonCodeControlPlaneRequestFailureDueToMisconfig,
		TerminationReasonCodeDatabaseConnectionFailure,
		TerminationReasonCodeDataAccessConfigChanged,
		TerminationReasonCodeDbfsComponentUnhealthy,
		TerminationReasonCodeDisasterRecoveryReplication,
		TerminationReasonCodeDnsResolutionError,
		TerminationReasonCodeDockerContainerCreationException,
		TerminationReasonCodeDockerImagePullFailure,
		TerminationReasonCodeDockerImageTooLargeForInstanceException,
		TerminationReasonCodeDockerInvalidOsException,
		TerminationReasonCodeDriverDnsResolutionFailure,
		TerminationReasonCodeDriverEviction,
		TerminationReasonCodeDriverLaunchTimeout,
		TerminationReasonCodeDriverNodeUnreachable,
		TerminationReasonCodeDriverOutOfDisk,
		TerminationReasonCodeDriverOutOfMemory,
		TerminationReasonCodeDriverPodCreationFailure,
		TerminationReasonCodeDriverUnexpectedFailure,
		TerminationReasonCodeDriverUnhealthy,
		TerminationReasonCodeDriverUnreachable,
		TerminationReasonCodeDriverUnresponsive,
		TerminationReasonCodeDynamicSparkConfSizeExceeded,
		TerminationReasonCodeEosSparkImage,
		TerminationReasonCodeExecutionComponentUnhealthy,
		TerminationReasonCodeExecutorPodUnscheduled,
		TerminationReasonCodeGcpApiRateQuotaExceeded,
		TerminationReasonCodeGcpDeniedByOrgPolicy,
		TerminationReasonCodeGcpForbidden,
		TerminationReasonCodeGcpIamTimeout,
		TerminationReasonCodeGcpInaccessibleKmsKeyFailure,
		TerminationReasonCodeGcpInsufficientCapacity,
		TerminationReasonCodeGcpIpSpaceExhausted,
		TerminationReasonCodeGcpKmsKeyPermissionDenied,
		TerminationReasonCodeGcpNotFound,
		TerminationReasonCodeGcpQuotaExceeded,
		TerminationReasonCodeGcpResourceQuotaExceeded,
		TerminationReasonCodeGcpServiceAccountAccessDenied,
		TerminationReasonCodeGcpServiceAccountDeleted,
		TerminationReasonCodeGcpServiceAccountNotFound,
		TerminationReasonCodeGcpSubnetNotReady,
		TerminationReasonCodeGcpTrustedImageProjectsViolated,
		TerminationReasonCodeGkeBasedClusterTermination,
		TerminationReasonCodeGlobalInitScriptFailure,
		TerminationReasonCodeHiveMetastoreProvisioningFailure,
		TerminationReasonCodeImagePullPermissionDenied,
		TerminationReasonCodeInactivity,
		TerminationReasonCodeInitContainerNotFinished,
		TerminationReasonCodeInitScriptFailure,
		TerminationReasonCodeInstancePoolClusterFailure,
		TerminationReasonCodeInstancePoolMaxCapacityReached,
		TerminationReasonCodeInstancePoolNotFound,
		TerminationReasonCodeInstanceUnreachable,
		TerminationReasonCodeInstanceUnreachableDueToMisconfig,
		TerminationReasonCodeInternalCapacityFailure,
		TerminationReasonCodeInternalError,
		TerminationReasonCodeInvalidArgument,
		TerminationReasonCodeInvalidAwsParameter,
		TerminationReasonCodeInvalidInstancePlacementProtocol,
		TerminationReasonCodeInvalidSparkImage,
		TerminationReasonCodeInvalidWorkerImageFailure,
		TerminationReasonCodeInPenaltyBox,
		TerminationReasonCodeIpExhaustionFailure,
		TerminationReasonCodeJobFinished,
		TerminationReasonCodeK8sAutoscalingFailure,
		TerminationReasonCodeK8sDbrClusterLaunchTimeout,
		TerminationReasonCodeLazyAllocationTimeout,
		TerminationReasonCodeMaintenanceMode,
		TerminationReasonCodeMetastoreComponentUnhealthy,
		TerminationReasonCodeNephosResourceManagement,
		TerminationReasonCodeNetvisorSetupTimeout,
		TerminationReasonCodeNetworkCheckControlPlaneFailure,
		TerminationReasonCodeNetworkCheckDnsServerFailure,
		TerminationReasonCodeNetworkCheckMetadataEndpointFailure,
		TerminationReasonCodeNetworkCheckMultipleComponentsFailure,
		TerminationReasonCodeNetworkCheckNicFailure,
		TerminationReasonCodeNetworkCheckStorageFailure,
		TerminationReasonCodeNetworkConfigurationFailure,
		TerminationReasonCodeNfsMountFailure,
		TerminationReasonCodeNoMatchedK8s,
		TerminationReasonCodeNoMatchedK8sTestingTag,
		TerminationReasonCodeNpipTunnelSetupFailure,
		TerminationReasonCodeNpipTunnelTokenFailure,
		TerminationReasonCodePodAssignmentFailure,
		TerminationReasonCodePodSchedulingFailure,
		TerminationReasonCodeRequestRejected,
		TerminationReasonCodeRequestThrottled,
		TerminationReasonCodeResourceUsageBlocked,
		TerminationReasonCodeSecretCreationFailure,
		TerminationReasonCodeSecretPermissionDenied,
		TerminationReasonCodeSecretResolutionError,
		TerminationReasonCodeSecurityAgentsFailedInitialVerification,
		TerminationReasonCodeSecurityDaemonRegistrationException,
		TerminationReasonCodeSelfBootstrapFailure,
		TerminationReasonCodeServerlessLongRunningTerminated,
		TerminationReasonCodeSkippedSlowNodes,
		TerminationReasonCodeSlowImageDownload,
		TerminationReasonCodeSparkError,
		TerminationReasonCodeSparkImageDownloadFailure,
		TerminationReasonCodeSparkImageDownloadThrottled,
		TerminationReasonCodeSparkImageNotFound,
		TerminationReasonCodeSparkStartupFailure,
		TerminationReasonCodeSpotInstanceTermination,
		TerminationReasonCodeSshBootstrapFailure,
		TerminationReasonCodeStorageDownloadFailure,
		TerminationReasonCodeStorageDownloadFailureDueToMisconfig,
		TerminationReasonCodeStorageDownloadFailureSlow,
		TerminationReasonCodeStorageDownloadFailureThrottled,
		TerminationReasonCodeStsClientSetupFailure,
		TerminationReasonCodeSubnetExhaustedFailure,
		TerminationReasonCodeTemporarilyUnavailable,
		TerminationReasonCodeTrialExpired,
		TerminationReasonCodeUnexpectedLaunchFailure,
		TerminationReasonCodeUnexpectedPodRecreation,
		TerminationReasonCodeUnknown,
		TerminationReasonCodeUnsupportedInstanceType,
		TerminationReasonCodeUpdateInstanceProfileFailure,
		TerminationReasonCodeUserInitiatedVmTermination,
		TerminationReasonCodeUserRequest,
		TerminationReasonCodeWorkerSetupFailure,
		TerminationReasonCodeWorkspaceCancelledError,
		TerminationReasonCodeWorkspaceConfigurationError,
		TerminationReasonCodeWorkspaceUpdate,
	}
}

// Type always returns TerminationReasonCode to satisfy [pflag.Value] interface
func (f *TerminationReasonCode) Type() string {
	return "TerminationReasonCode"
}

// type of the termination
type TerminationReasonType string

const TerminationReasonTypeClientError TerminationReasonType = `CLIENT_ERROR`

const TerminationReasonTypeCloudFailure TerminationReasonType = `CLOUD_FAILURE`

const TerminationReasonTypeServiceFault TerminationReasonType = `SERVICE_FAULT`

const TerminationReasonTypeSuccess TerminationReasonType = `SUCCESS`

// String representation for [fmt.Print]
func (f *TerminationReasonType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `SERVICE_FAULT`, `SUCCESS`:
		*f = TerminationReasonType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "SERVICE_FAULT", "SUCCESS"`, v)
	}
}

// Values returns all possible values for TerminationReasonType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationReasonType) Values() []TerminationReasonType {
	return []TerminationReasonType{
		TerminationReasonTypeClientError,
		TerminationReasonTypeCloudFailure,
		TerminationReasonTypeServiceFault,
		TerminationReasonTypeSuccess,
	}
}

// Type always returns TerminationReasonType to satisfy [pflag.Value] interface
func (f *TerminationReasonType) Type() string {
	return "TerminationReasonType"
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to uninstall.
	Libraries []Library `json:"libraries"`
}

type UnpinCluster struct {
	ClusterId string `json:"cluster_id"`
}

type UpdateCluster struct {
	// The cluster to be updated.
	Cluster *UpdateClusterResource `json:"cluster,omitempty"`
	// ID of the cluster.
	ClusterId string `json:"cluster_id"`
	// Used to specify which cluster attributes and size fields to update. See
	// https://google.aip.dev/161 for more details.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"update_mask"`
}

type UpdateClusterResource struct {
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
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool `json:"is_single_node,omitempty"`

	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string `json:"spark_version,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`

	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateClusterResource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateClusterResource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A storage location back by UC Volumes.
type VolumesStorageInfo struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	Destination string `json:"destination"`
}

// Cluster Attributes showing for clusters workload types.
type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients ClientsTypes `json:"clients"`
}

// A storage location in Workspace Filesystem (WSFS)
type WorkspaceStorageInfo struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	Destination string `json:"destination"`
}
