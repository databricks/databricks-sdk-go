// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"encoding/json"
	"fmt"
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
	// Wire name: 'iam_role_arn'
	IamRoleArn string `json:"iam_role_arn,omitempty"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn"`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	// Wire name: 'is_meta_instance_profile'
	IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	// Wire name: 'skip_validation'
	SkipValidation bool `json:"skip_validation,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AddInstanceProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &addInstanceProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := addInstanceProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AddInstanceProfile) MarshalJSON() ([]byte, error) {
	pb, err := addInstanceProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AddResponse struct {
}

func (st *AddResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &addResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := addResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AddResponse) MarshalJSON() ([]byte, error) {
	pb, err := addResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A storage location in Adls Gen2
type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	// Wire name: 'destination'
	Destination string `json:"destination"`
}

func (st *Adlsgen2Info) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &adlsgen2InfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := adlsgen2InfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Adlsgen2Info) MarshalJSON() ([]byte, error) {
	pb, err := adlsgen2InfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	// Wire name: 'max_workers'
	MaxWorkers int `json:"max_workers,omitempty"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	// Wire name: 'min_workers'
	MinWorkers int `json:"min_workers,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AutoScale) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &autoScalePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := autoScaleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AutoScale) MarshalJSON() ([]byte, error) {
	pb, err := autoScaleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Attributes set during cluster creation which are related to Amazon Web
// Services.
type AwsAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones.
	//
	// Note: If `first_on_demand` is zero, this availability type will be used
	// for the entire cluster.
	// Wire name: 'availability'
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
	// Wire name: 'ebs_volume_count'
	EbsVolumeCount int `json:"ebs_volume_count,omitempty"`
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	// Wire name: 'ebs_volume_iops'
	EbsVolumeIops int `json:"ebs_volume_iops,omitempty"`
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	// Wire name: 'ebs_volume_size'
	EbsVolumeSize int `json:"ebs_volume_size,omitempty"`
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	// Wire name: 'ebs_volume_throughput'
	EbsVolumeThroughput int `json:"ebs_volume_throughput,omitempty"`
	// The type of EBS volumes that will be launched with this cluster.
	// Wire name: 'ebs_volume_type'
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
	// Wire name: 'first_on_demand'
	FirstOnDemand int `json:"first_on_demand,omitempty"`
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	// Wire name: 'instance_profile_arn'
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
	// Wire name: 'spot_bid_price_percent'
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
	// Wire name: 'zone_id'
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &awsAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := awsAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := awsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'availability'
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
	// Wire name: 'first_on_demand'
	FirstOnDemand int `json:"first_on_demand,omitempty"`
	// Defines values necessary to configure and run Azure Log Analytics agent
	// Wire name: 'log_analytics_info'
	LogAnalyticsInfo *LogAnalyticsInfo `json:"log_analytics_info,omitempty"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	// Wire name: 'spot_bid_max_price'
	SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AzureAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &azureAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := azureAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AzureAttributes) MarshalJSON() ([]byte, error) {
	pb, err := azureAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'clusterId'
	ClusterId string `json:"clusterId,omitempty"`

	// Wire name: 'commandId'
	CommandId string `json:"commandId,omitempty"`

	// Wire name: 'contextId'
	ContextId string `json:"contextId,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CancelCommand) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelCommandPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelCommandFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelCommand) MarshalJSON() ([]byte, error) {
	pb, err := cancelCommandToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelResponse struct {
}

func (st *CancelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ChangeClusterOwner struct {

	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// New owner of the cluster_id after this RPC.
	// Wire name: 'owner_username'
	OwnerUsername string `json:"owner_username"`
}

func (st *ChangeClusterOwner) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &changeClusterOwnerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := changeClusterOwnerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ChangeClusterOwner) MarshalJSON() ([]byte, error) {
	pb, err := changeClusterOwnerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ChangeClusterOwnerResponse struct {
}

func (st *ChangeClusterOwnerResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &changeClusterOwnerResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := changeClusterOwnerResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ChangeClusterOwnerResponse) MarshalJSON() ([]byte, error) {
	pb, err := changeClusterOwnerResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	// Wire name: 'jobs'
	Jobs bool `json:"jobs,omitempty"`
	// With notebooks set, this cluster can be used for notebooks
	// Wire name: 'notebooks'
	Notebooks bool `json:"notebooks,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClientsTypes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clientsTypesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clientsTypesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClientsTypes) MarshalJSON() ([]byte, error) {
	pb, err := clientsTypesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CloneCluster struct {
	// The cluster that is being cloned.
	// Wire name: 'source_cluster_id'
	SourceClusterId string `json:"source_cluster_id"`
}

func (st *CloneCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cloneClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cloneClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CloneCluster) MarshalJSON() ([]byte, error) {
	pb, err := cloneClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CloudProviderNodeInfo struct {
	// Status as reported by the cloud provider
	// Wire name: 'status'
	Status []CloudProviderNodeStatus `json:"status,omitempty"`
}

func (st *CloudProviderNodeInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cloudProviderNodeInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cloudProviderNodeInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CloudProviderNodeInfo) MarshalJSON() ([]byte, error) {
	pb, err := cloudProviderNodeInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := clusterAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ClusterPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := clusterAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Common set of attributes set during cluster creation. These attributes cannot
// be changed over the lifetime of a cluster.
type ClusterAttributes struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
	// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate
	// access mode depending on your compute configuration. *
	// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
	// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
	//
	// The following modes can be used regardless of `kind`. * `NONE`: No
	// security isolation for multiple users sharing the cluster. Data
	// governance features are not available in this mode. * `SINGLE_USER`: A
	// secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// The kind of compute described by this compute specification.
	//
	// Depending on `kind`, different validations and default values will be
	// applied.
	//
	// Clusters with `kind = CLASSIC_PREVIEW` support the following fields,
	// whereas clusters with no specified `kind` do not. *
	// [is_single_node](/api/workspace/clusters/create#is_single_node) *
	// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
	// [data_security_mode](/api/workspace/clusters/create#data_security_mode)
	// set to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
	// `DATA_SECURITY_MODE_STANDARD`
	//
	// By using the [simple form], your clusters are automatically using `kind =
	// CLASSIC_PREVIEW`.
	//
	// [simple form]: https://docs.databricks.com/compute/simple-form.html
	// Wire name: 'kind'
	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string `json:"node_type_id,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string `json:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	// Wire name: 'workload_type'
	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAttributes) MarshalJSON() ([]byte, error) {
	pb, err := clusterAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterCompliance struct {
	// Canonical unique identifier for a cluster.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// Whether this cluster is in compliance with the latest version of its
	// policy.
	// Wire name: 'is_compliant'
	IsCompliant bool `json:"is_compliant,omitempty"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	// Wire name: 'violations'
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterCompliance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterCompliancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterComplianceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterCompliance) MarshalJSON() ([]byte, error) {
	pb, err := clusterComplianceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Describes all of the metadata about a single Spark cluster in Databricks.
type ClusterDetails struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
	// Number of CPU cores available for this cluster. Note that this can be
	// fractional, e.g. 7.5 cores, since certain node types are configured to
	// share cores between Spark nodes on the same instance.
	// Wire name: 'cluster_cores'
	ClusterCores float64 `json:"cluster_cores,omitempty"`
	// Canonical identifier for the cluster. This id is retained during cluster
	// restarts and resizes, while each new cluster has a globally unique id.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster log delivery status.
	// Wire name: 'cluster_log_status'
	ClusterLogStatus *LogSyncStatus `json:"cluster_log_status,omitempty"`
	// Total amount of cluster memory, in megabytes
	// Wire name: 'cluster_memory_mb'
	ClusterMemoryMb int64 `json:"cluster_memory_mb,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string `json:"cluster_name,omitempty"`
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request.
	// Wire name: 'cluster_source'
	ClusterSource ClusterSource `json:"cluster_source,omitempty"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
	// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate
	// access mode depending on your compute configuration. *
	// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
	// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
	//
	// The following modes can be used regardless of `kind`. * `NONE`: No
	// security isolation for multiple users sharing the cluster. Data
	// governance features are not available in this mode. * `SINGLE_USER`: A
	// secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	// Wire name: 'data_security_mode'
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
	// Wire name: 'default_tags'
	DefaultTags map[string]string `json:"default_tags,omitempty"`
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	// Wire name: 'driver'
	Driver *SparkNode `json:"driver,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Nodes on which the Spark executors reside.
	// Wire name: 'executors'
	Executors []SparkNode `json:"executors,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	// Wire name: 'jdbc_port'
	JdbcPort int `json:"jdbc_port,omitempty"`
	// The kind of compute described by this compute specification.
	//
	// Depending on `kind`, different validations and default values will be
	// applied.
	//
	// Clusters with `kind = CLASSIC_PREVIEW` support the following fields,
	// whereas clusters with no specified `kind` do not. *
	// [is_single_node](/api/workspace/clusters/create#is_single_node) *
	// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
	// [data_security_mode](/api/workspace/clusters/create#data_security_mode)
	// set to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
	// `DATA_SECURITY_MODE_STANDARD`
	//
	// By using the [simple form], your clusters are automatically using `kind =
	// CLASSIC_PREVIEW`.
	//
	// [simple form]: https://docs.databricks.com/compute/simple-form.html
	// Wire name: 'kind'
	Kind Kind `json:"kind,omitempty"`
	// the timestamp that the cluster was started/restarted
	// Wire name: 'last_restarted_time'
	LastRestartedTime int64 `json:"last_restarted_time,omitempty"`
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	// Wire name: 'last_state_loss_time'
	LastStateLossTime int64 `json:"last_state_loss_time,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string `json:"spark_conf,omitempty"`
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	// Wire name: 'spark_context_id'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string `json:"spark_version,omitempty"`
	// The spec contains a snapshot of the latest user specified settings that
	// were used to create/edit the cluster. Note: not included in the response
	// of the ListClusters API.
	// Wire name: 'spec'
	Spec *ClusterSpec `json:"spec,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	// Wire name: 'start_time'
	StartTime int64 `json:"start_time,omitempty"`
	// Current state of the cluster.
	// Wire name: 'state'
	State State `json:"state,omitempty"`
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	// Wire name: 'state_message'
	StateMessage string `json:"state_message,omitempty"`
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	// Wire name: 'terminated_time'
	TerminatedTime int64 `json:"terminated_time,omitempty"`
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	// Wire name: 'termination_reason'
	TerminationReason *TerminationReason `json:"termination_reason,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	// Wire name: 'workload_type'
	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterDetails) MarshalJSON() ([]byte, error) {
	pb, err := clusterDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterEvent struct {

	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`

	// Wire name: 'data_plane_event_details'
	DataPlaneEventDetails *DataPlaneEventDetails `json:"data_plane_event_details,omitempty"`

	// Wire name: 'details'
	Details *EventDetails `json:"details,omitempty"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	// Wire name: 'timestamp'
	Timestamp int64 `json:"timestamp,omitempty"`

	// Wire name: 'type'
	Type EventType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterEvent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterEventPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterEventFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterEvent) MarshalJSON() ([]byte, error) {
	pb, err := clusterEventToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`
	// Status of all libraries on the cluster.
	// Wire name: 'library_statuses'
	LibraryStatuses []LibraryFullStatus `json:"library_statuses,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterLibraryStatuses) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterLibraryStatusesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterLibraryStatusesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterLibraryStatuses) MarshalJSON() ([]byte, error) {
	pb, err := clusterLibraryStatusesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Cluster log delivery config
type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	// Wire name: 'dbfs'
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	// Wire name: 's3'
	S3 *S3StorageInfo `json:"s3,omitempty"`
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	// Wire name: 'volumes'
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
}

func (st *ClusterLogConf) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterLogConfPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterLogConfFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterLogConf) MarshalJSON() ([]byte, error) {
	pb, err := clusterLogConfToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPermission) MarshalJSON() ([]byte, error) {
	pb, err := clusterPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []ClusterAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPermissions) MarshalJSON() ([]byte, error) {
	pb, err := clusterPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := clusterPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ClusterAccessControlRequest `json:"access_control_list,omitempty"`
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" tf:"-"`
}

func (st *ClusterPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := clusterPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPolicyAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPolicyAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPolicyAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPolicyAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := clusterPolicyAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPolicyAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ClusterPolicyPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPolicyAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPolicyAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPolicyAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPolicyAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := clusterPolicyAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPolicyPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPolicyPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPolicyPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPolicyPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPolicyPermission) MarshalJSON() ([]byte, error) {
	pb, err := clusterPolicyPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []ClusterPolicyAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPolicyPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPolicyPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPolicyPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPolicyPermissions) MarshalJSON() ([]byte, error) {
	pb, err := clusterPolicyPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPolicyPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterPolicyPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPolicyPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPolicyPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPolicyPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := clusterPolicyPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterPolicyPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ClusterPolicyAccessControlRequest `json:"access_control_list,omitempty"`
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" tf:"-"`
}

func (st *ClusterPolicyPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterPolicyPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterPolicyPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterPolicyPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := clusterPolicyPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Represents a change to the cluster settings required for the cluster to
// become compliant with its policy.
type ClusterSettingsChange struct {
	// The field where this change would be made.
	// Wire name: 'field'
	Field string `json:"field,omitempty"`
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	// Wire name: 'new_value'
	NewValue string `json:"new_value,omitempty"`
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	// Wire name: 'previous_value'
	PreviousValue string `json:"previous_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterSettingsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterSettingsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterSettingsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterSettingsChange) MarshalJSON() ([]byte, error) {
	pb, err := clusterSettingsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterSize struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterSize) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterSizePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterSizeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterSize) MarshalJSON() ([]byte, error) {
	pb, err := clusterSizeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
	// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate
	// access mode depending on your compute configuration. *
	// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
	// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
	//
	// The following modes can be used regardless of `kind`. * `NONE`: No
	// security isolation for multiple users sharing the cluster. Data
	// governance features are not available in this mode. * `SINGLE_USER`: A
	// secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// The kind of compute described by this compute specification.
	//
	// Depending on `kind`, different validations and default values will be
	// applied.
	//
	// Clusters with `kind = CLASSIC_PREVIEW` support the following fields,
	// whereas clusters with no specified `kind` do not. *
	// [is_single_node](/api/workspace/clusters/create#is_single_node) *
	// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
	// [data_security_mode](/api/workspace/clusters/create#data_security_mode)
	// set to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
	// `DATA_SECURITY_MODE_STANDARD`
	//
	// By using the [simple form], your clusters are automatically using `kind =
	// CLASSIC_PREVIEW`.
	//
	// [simple form]: https://docs.databricks.com/compute/simple-form.html
	// Wire name: 'kind'
	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string `json:"spark_version,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	// Wire name: 'workload_type'
	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterSpec) MarshalJSON() ([]byte, error) {
	pb, err := clusterSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get status
type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId string `json:"-" tf:"-"`
}

func (st *ClusterStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterStatus) MarshalJSON() ([]byte, error) {
	pb, err := clusterStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Command struct {
	// Running cluster id
	// Wire name: 'clusterId'
	ClusterId string `json:"clusterId,omitempty"`
	// Executable code
	// Wire name: 'command'
	Command string `json:"command,omitempty"`
	// Running context id
	// Wire name: 'contextId'
	ContextId string `json:"contextId,omitempty"`

	// Wire name: 'language'
	Language Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Command) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &commandPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := commandFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Command) MarshalJSON() ([]byte, error) {
	pb, err := commandToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Get command info
type CommandStatusRequest struct {
	ClusterId string `json:"-" tf:"-"`

	CommandId string `json:"-" tf:"-"`

	ContextId string `json:"-" tf:"-"`
}

func (st *CommandStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &commandStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := commandStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CommandStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := commandStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CommandStatusResponse struct {

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'results'
	Results *Results `json:"results,omitempty"`

	// Wire name: 'status'
	Status CommandStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CommandStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &commandStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := commandStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CommandStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := commandStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Get status
type ContextStatusRequest struct {
	ClusterId string `json:"-" tf:"-"`

	ContextId string `json:"-" tf:"-"`
}

func (st *ContextStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &contextStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := contextStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ContextStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := contextStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ContextStatusResponse struct {

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'status'
	Status ContextStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ContextStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &contextStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := contextStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ContextStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := contextStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	// Wire name: 'clone_from'
	CloneFrom *CloneCluster `json:"clone_from,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
	// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate
	// access mode depending on your compute configuration. *
	// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
	// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
	//
	// The following modes can be used regardless of `kind`. * `NONE`: No
	// security isolation for multiple users sharing the cluster. Data
	// governance features are not available in this mode. * `SINGLE_USER`: A
	// secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// The kind of compute described by this compute specification.
	//
	// Depending on `kind`, different validations and default values will be
	// applied.
	//
	// Clusters with `kind = CLASSIC_PREVIEW` support the following fields,
	// whereas clusters with no specified `kind` do not. *
	// [is_single_node](/api/workspace/clusters/create#is_single_node) *
	// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
	// [data_security_mode](/api/workspace/clusters/create#data_security_mode)
	// set to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
	// `DATA_SECURITY_MODE_STANDARD`
	//
	// By using the [simple form], your clusters are automatically using `kind =
	// CLASSIC_PREVIEW`.
	//
	// [simple form]: https://docs.databricks.com/compute/simple-form.html
	// Wire name: 'kind'
	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string `json:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	// Wire name: 'workload_type'
	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCluster) MarshalJSON() ([]byte, error) {
	pb, err := createClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateClusterResponse struct {

	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := createClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateContext struct {
	// Running cluster id
	// Wire name: 'clusterId'
	ClusterId string `json:"clusterId,omitempty"`

	// Wire name: 'language'
	Language Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateContext) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createContextPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createContextFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateContext) MarshalJSON() ([]byte, error) {
	pb, err := createContextToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	// Wire name: 'disk_spec'
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *InstancePoolGcpAttributes `json:"gcp_attributes,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string `json:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string `json:"node_type_id"`
	// Custom Docker Image BYOC
	// Wire name: 'preloaded_docker_images'
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'preloaded_spark_versions'
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := createInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateInstancePoolResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createInstancePoolResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createInstancePoolResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateInstancePoolResponse) MarshalJSON() ([]byte, error) {
	pb, err := createInstancePoolResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreatePolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	// Wire name: 'libraries'
	Libraries []Library `json:"libraries,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	// Wire name: 'max_clusters_per_user'
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	// Wire name: 'name'
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
	// Wire name: 'policy_family_definition_overrides'
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePolicy) MarshalJSON() ([]byte, error) {
	pb, err := createPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := createPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateResponse struct {
	// The global init script ID.
	// Wire name: 'script_id'
	ScriptId string `json:"script_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := createResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Created struct {

	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Created) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Created) MarshalJSON() ([]byte, error) {
	pb, err := createdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CustomPolicyTag struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	// Wire name: 'key'
	Key string `json:"key"`
	// The value of the tag.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CustomPolicyTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customPolicyTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customPolicyTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomPolicyTag) MarshalJSON() ([]byte, error) {
	pb, err := customPolicyTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DataPlaneEventDetails struct {

	// Wire name: 'event_type'
	EventType DataPlaneEventDetailsEventType `json:"event_type,omitempty"`

	// Wire name: 'executor_failures'
	ExecutorFailures int `json:"executor_failures,omitempty"`

	// Wire name: 'host_id'
	HostId string `json:"host_id,omitempty"`

	// Wire name: 'timestamp'
	Timestamp int64 `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DataPlaneEventDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dataPlaneEventDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dataPlaneEventDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DataPlaneEventDetails) MarshalJSON() ([]byte, error) {
	pb, err := dataPlaneEventDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'destination'
	Destination string `json:"destination"`
}

func (st *DbfsStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dbfsStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dbfsStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DbfsStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := dbfsStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteCluster struct {
	// The cluster to be terminated.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
}

func (st *DeleteCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCluster) MarshalJSON() ([]byte, error) {
	pb, err := deleteClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteClusterResponse struct {
}

func (st *DeleteClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete init script
type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId string `json:"-" tf:"-"`
}

func (st *DeleteGlobalInitScriptRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteGlobalInitScriptRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteGlobalInitScriptRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteGlobalInitScriptRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteGlobalInitScriptRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id"`
}

func (st *DeleteInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := deleteInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteInstancePoolResponse struct {
}

func (st *DeleteInstancePoolResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteInstancePoolResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteInstancePoolResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteInstancePoolResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteInstancePoolResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id"`
}

func (st *DeletePolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePolicy) MarshalJSON() ([]byte, error) {
	pb, err := deletePolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeletePolicyResponse struct {
}

func (st *DeletePolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := deletePolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteResponse struct {
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DestroyContext struct {

	// Wire name: 'clusterId'
	ClusterId string `json:"clusterId"`

	// Wire name: 'contextId'
	ContextId string `json:"contextId"`
}

func (st *DestroyContext) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &destroyContextPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := destroyContextFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DestroyContext) MarshalJSON() ([]byte, error) {
	pb, err := destroyContextToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DestroyResponse struct {
}

func (st *DestroyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &destroyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := destroyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DestroyResponse) MarshalJSON() ([]byte, error) {
	pb, err := destroyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'disk_count'
	DiskCount int `json:"disk_count,omitempty"`

	// Wire name: 'disk_iops'
	DiskIops int `json:"disk_iops,omitempty"`
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	// Wire name: 'disk_size'
	DiskSize int `json:"disk_size,omitempty"`

	// Wire name: 'disk_throughput'
	DiskThroughput int `json:"disk_throughput,omitempty"`
	// The type of disks that will be launched with this cluster.
	// Wire name: 'disk_type'
	DiskType *DiskType `json:"disk_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DiskSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &diskSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := diskSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DiskSpec) MarshalJSON() ([]byte, error) {
	pb, err := diskSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Describes the disk type.
type DiskType struct {
	// All Azure Disk types that Databricks supports. See
	// https://docs.microsoft.com/en-us/azure/storage/storage-about-disks-and-vhds-linux#types-of-disks
	// Wire name: 'azure_disk_volume_type'
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType `json:"azure_disk_volume_type,omitempty"`
	// All EBS volume types that Databricks supports. See
	// https://aws.amazon.com/ebs/details/ for details.
	// Wire name: 'ebs_volume_type'
	EbsVolumeType DiskTypeEbsVolumeType `json:"ebs_volume_type,omitempty"`
}

func (st *DiskType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &diskTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := diskTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DiskType) MarshalJSON() ([]byte, error) {
	pb, err := diskTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'password'
	Password string `json:"password,omitempty"`
	// Name of the user
	// Wire name: 'username'
	Username string `json:"username,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DockerBasicAuth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dockerBasicAuthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dockerBasicAuthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DockerBasicAuth) MarshalJSON() ([]byte, error) {
	pb, err := dockerBasicAuthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DockerImage struct {
	// Basic auth with username and password
	// Wire name: 'basic_auth'
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`
	// URL of the docker image.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DockerImage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dockerImagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dockerImageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DockerImage) MarshalJSON() ([]byte, error) {
	pb, err := dockerImageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
	// ID of the cluster
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
	// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate
	// access mode depending on your compute configuration. *
	// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
	// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
	//
	// The following modes can be used regardless of `kind`. * `NONE`: No
	// security isolation for multiple users sharing the cluster. Data
	// governance features are not available in this mode. * `SINGLE_USER`: A
	// secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// The kind of compute described by this compute specification.
	//
	// Depending on `kind`, different validations and default values will be
	// applied.
	//
	// Clusters with `kind = CLASSIC_PREVIEW` support the following fields,
	// whereas clusters with no specified `kind` do not. *
	// [is_single_node](/api/workspace/clusters/create#is_single_node) *
	// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
	// [data_security_mode](/api/workspace/clusters/create#data_security_mode)
	// set to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
	// `DATA_SECURITY_MODE_STANDARD`
	//
	// By using the [simple form], your clusters are automatically using `kind =
	// CLASSIC_PREVIEW`.
	//
	// [simple form]: https://docs.databricks.com/compute/simple-form.html
	// Wire name: 'kind'
	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string `json:"spark_version"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	// Wire name: 'workload_type'
	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EditCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditCluster) MarshalJSON() ([]byte, error) {
	pb, err := editClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditClusterResponse struct {
}

func (st *EditClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := editClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditInstancePool struct {
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Instance pool ID
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string `json:"instance_pool_name"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string `json:"node_type_id"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EditInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := editInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditInstancePoolResponse struct {
}

func (st *EditInstancePoolResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editInstancePoolResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editInstancePoolResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditInstancePoolResponse) MarshalJSON() ([]byte, error) {
	pb, err := editInstancePoolResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditPolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	// Wire name: 'libraries'
	Libraries []Library `json:"libraries,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	// Wire name: 'max_clusters_per_user'
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	// Wire name: 'name'
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
	// Wire name: 'policy_family_definition_overrides'
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
	// The ID of the policy to update.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EditPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditPolicy) MarshalJSON() ([]byte, error) {
	pb, err := editPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditPolicyResponse struct {
}

func (st *EditPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := editPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditResponse struct {
}

func (st *EditResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditResponse) MarshalJSON() ([]byte, error) {
	pb, err := editResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnforceClusterComplianceRequest struct {
	// The ID of the cluster you want to enforce policy compliance on.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	// Wire name: 'validate_only'
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnforceClusterComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforceClusterComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforceClusterComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforceClusterComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := enforceClusterComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnforceClusterComplianceResponse struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	// Wire name: 'changes'
	Changes []ClusterSettingsChange `json:"changes,omitempty"`
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	// Wire name: 'has_changes'
	HasChanges bool `json:"has_changes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnforceClusterComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforceClusterComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforceClusterComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforceClusterComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := enforceClusterComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type Environment struct {
	// Use `environment_version` instead.
	// Wire name: 'client'
	Client string `json:"client,omitempty"`
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a valid pip requirements file line per
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/.
	// Allowed dependencies include a requirement specifier, an archive URL, a
	// local project path (such as WSFS or UC Volumes in Databricks), or a VCS
	// project URL.
	// Wire name: 'dependencies'
	Dependencies []string `json:"dependencies,omitempty"`
	// Required. Environment version used by the environment. Each version comes
	// with a specific Python version and a set of Python packages. The version
	// is a string, consisting of an integer.
	// Wire name: 'environment_version'
	EnvironmentVersion string `json:"environment_version,omitempty"`
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	// Wire name: 'jar_dependencies'
	JarDependencies []string `json:"jar_dependencies,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Environment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &environmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := environmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Environment) MarshalJSON() ([]byte, error) {
	pb, err := environmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EventDetails struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	// Wire name: 'attributes'
	Attributes *ClusterAttributes `json:"attributes,omitempty"`
	// The cause of a change in target size.
	// Wire name: 'cause'
	Cause EventDetailsCause `json:"cause,omitempty"`
	// The actual cluster size that was set in the cluster creation or edit.
	// Wire name: 'cluster_size'
	ClusterSize *ClusterSize `json:"cluster_size,omitempty"`
	// The current number of vCPUs in the cluster.
	// Wire name: 'current_num_vcpus'
	CurrentNumVcpus int `json:"current_num_vcpus,omitempty"`
	// The current number of nodes in the cluster.
	// Wire name: 'current_num_workers'
	CurrentNumWorkers int `json:"current_num_workers,omitempty"`

	// Wire name: 'did_not_expand_reason'
	DidNotExpandReason string `json:"did_not_expand_reason,omitempty"`
	// Current disk size in bytes
	// Wire name: 'disk_size'
	DiskSize int64 `json:"disk_size,omitempty"`
	// More details about the change in driver's state
	// Wire name: 'driver_state_message'
	DriverStateMessage string `json:"driver_state_message,omitempty"`
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	// Wire name: 'enable_termination_for_node_blocklisted'
	EnableTerminationForNodeBlocklisted bool `json:"enable_termination_for_node_blocklisted,omitempty"`

	// Wire name: 'free_space'
	FreeSpace int64 `json:"free_space,omitempty"`
	// List of global and cluster init scripts associated with this cluster
	// event.
	// Wire name: 'init_scripts'
	InitScripts *InitScriptEventDetails `json:"init_scripts,omitempty"`
	// Instance Id where the event originated from
	// Wire name: 'instance_id'
	InstanceId string `json:"instance_id,omitempty"`
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	// Wire name: 'job_run_name'
	JobRunName string `json:"job_run_name,omitempty"`
	// The cluster attributes before a cluster was edited.
	// Wire name: 'previous_attributes'
	PreviousAttributes *ClusterAttributes `json:"previous_attributes,omitempty"`
	// The size of the cluster before an edit or resize.
	// Wire name: 'previous_cluster_size'
	PreviousClusterSize *ClusterSize `json:"previous_cluster_size,omitempty"`
	// Previous disk size in bytes
	// Wire name: 'previous_disk_size'
	PreviousDiskSize int64 `json:"previous_disk_size,omitempty"`
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	// Wire name: 'reason'
	Reason *TerminationReason `json:"reason,omitempty"`
	// The targeted number of vCPUs in the cluster.
	// Wire name: 'target_num_vcpus'
	TargetNumVcpus int `json:"target_num_vcpus,omitempty"`
	// The targeted number of nodes in the cluster.
	// Wire name: 'target_num_workers'
	TargetNumWorkers int `json:"target_num_workers,omitempty"`
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	// Wire name: 'user'
	User string `json:"user,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EventDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &eventDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := eventDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EventDetails) MarshalJSON() ([]byte, error) {
	pb, err := eventDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'availability'
	Availability GcpAvailability `json:"availability,omitempty"`
	// Boot disk size in GB
	// Wire name: 'boot_disk_size'
	BootDiskSize int `json:"boot_disk_size,omitempty"`
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	// Wire name: 'google_service_account'
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	// Wire name: 'local_ssd_count'
	LocalSsdCount int `json:"local_ssd_count,omitempty"`
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// 'availability' field instead.
	// Wire name: 'use_preemptible_executors'
	UsePreemptibleExecutors bool `json:"use_preemptible_executors,omitempty"`
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default]. -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	// Wire name: 'zone_id'
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GcpAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcpAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcpAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcpAttributes) MarshalJSON() ([]byte, error) {
	pb, err := gcpAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'destination'
	Destination string `json:"destination"`
}

func (st *GcsStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gcsStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gcsStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GcsStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := gcsStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get cluster policy compliance
type GetClusterComplianceRequest struct {
	// The ID of the cluster to get the compliance status
	ClusterId string `json:"-" tf:"-"`
}

func (st *GetClusterComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetClusterComplianceResponse struct {
	// Whether the cluster is compliant with its policy or not. Clusters could
	// be out of compliance if the policy was updated after the cluster was last
	// edited.
	// Wire name: 'is_compliant'
	IsCompliant bool `json:"is_compliant,omitempty"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	// Wire name: 'violations'
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetClusterComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := getClusterComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get cluster permission levels
type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" tf:"-"`
}

func (st *GetClusterPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetClusterPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ClusterPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetClusterPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get cluster permissions
type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" tf:"-"`
}

func (st *GetClusterPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get cluster policy permission levels
type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" tf:"-"`
}

func (st *GetClusterPolicyPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPolicyPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPolicyPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPolicyPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPolicyPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetClusterPolicyPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ClusterPolicyPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetClusterPolicyPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPolicyPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPolicyPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPolicyPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPolicyPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get cluster policy permissions
type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" tf:"-"`
}

func (st *GetClusterPolicyPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPolicyPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPolicyPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPolicyPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPolicyPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a cluster policy
type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string `json:"-" tf:"-"`
}

func (st *GetClusterPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get cluster info
type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	ClusterId string `json:"-" tf:"-"`
}

func (st *GetClusterRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getClusterRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getClusterRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetClusterRequest) MarshalJSON() ([]byte, error) {
	pb, err := getClusterRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetEvents struct {
	// The ID of the cluster to retrieve events about.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	// Wire name: 'end_time'
	EndTime int64 `json:"end_time,omitempty"`
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	// Wire name: 'event_types'
	EventTypes []EventType `json:"event_types,omitempty"`
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	// Wire name: 'limit'
	Limit int64 `json:"limit,omitempty"`
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	// Wire name: 'offset'
	Offset int64 `json:"offset,omitempty"`
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	// Wire name: 'order'
	Order GetEventsOrder `json:"order,omitempty"`
	// The maximum number of events to include in a page of events. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is empty or 0, the server will decide the number
	// of results to be returned. The field has to be in the range [0,500]. If
	// the value is outside the range, the server enforces 0 or 500.
	// Wire name: 'page_size'
	PageSize int `json:"page_size,omitempty"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of events respectively. If page_token
	// is empty, the first page is returned.
	// Wire name: 'page_token'
	PageToken string `json:"page_token,omitempty"`
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	// Wire name: 'start_time'
	StartTime int64 `json:"start_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetEvents) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEventsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEventsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEvents) MarshalJSON() ([]byte, error) {
	pb, err := getEventsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'events'
	Events []ClusterEvent `json:"events,omitempty"`
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	// Wire name: 'next_page'
	NextPage *GetEvents `json:"next_page,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	// Wire name: 'prev_page_token'
	PrevPageToken string `json:"prev_page_token,omitempty"`
	// Deprecated: Returns 0 when request uses page_token. Will start returning
	// zero when request uses offset/limit soon.
	//
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	// Wire name: 'total_count'
	TotalCount int64 `json:"total_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetEventsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEventsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEventsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEventsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getEventsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get an init script
type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId string `json:"-" tf:"-"`
}

func (st *GetGlobalInitScriptRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getGlobalInitScriptRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getGlobalInitScriptRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetGlobalInitScriptRequest) MarshalJSON() ([]byte, error) {
	pb, err := getGlobalInitScriptRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
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
	// Wire name: 'default_tags'
	DefaultTags map[string]string `json:"default_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	// Wire name: 'disk_spec'
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *InstancePoolGcpAttributes `json:"gcp_attributes,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Canonical unique identifier for the pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	// Wire name: 'preloaded_docker_images'
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'preloaded_spark_versions'
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Current state of the instance pool.
	// Wire name: 'state'
	State InstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	// Wire name: 'stats'
	Stats *InstancePoolStats `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	// Wire name: 'status'
	Status *InstancePoolStatus `json:"status,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := getInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get instance pool permission levels
type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" tf:"-"`
}

func (st *GetInstancePoolPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getInstancePoolPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getInstancePoolPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetInstancePoolPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getInstancePoolPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetInstancePoolPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []InstancePoolPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetInstancePoolPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getInstancePoolPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getInstancePoolPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetInstancePoolPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getInstancePoolPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get instance pool permissions
type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" tf:"-"`
}

func (st *GetInstancePoolPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getInstancePoolPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getInstancePoolPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetInstancePoolPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getInstancePoolPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get instance pool information
type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId string `json:"-" tf:"-"`
}

func (st *GetInstancePoolRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getInstancePoolRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getInstancePoolRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetInstancePoolRequest) MarshalJSON() ([]byte, error) {
	pb, err := getInstancePoolRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get policy family information
type GetPolicyFamilyRequest struct {
	// The family ID about which to retrieve information.
	PolicyFamilyId string `json:"-" tf:"-"`
	// The version number for the family to fetch. Defaults to the latest
	// version.
	Version int64 `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetPolicyFamilyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPolicyFamilyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPolicyFamilyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPolicyFamilyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPolicyFamilyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetSparkVersionsResponse struct {
	// All the available Spark versions.
	// Wire name: 'versions'
	Versions []SparkVersion `json:"versions,omitempty"`
}

func (st *GetSparkVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSparkVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSparkVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSparkVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getSparkVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GlobalInitScriptCreateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	// Wire name: 'name'
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
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// The Base64-encoded content of the script.
	// Wire name: 'script'
	Script string `json:"script"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GlobalInitScriptCreateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &globalInitScriptCreateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := globalInitScriptCreateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GlobalInitScriptCreateRequest) MarshalJSON() ([]byte, error) {
	pb, err := globalInitScriptCreateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GlobalInitScriptDetails struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'created_at'
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// The global init script ID.
	// Wire name: 'script_id'
	ScriptId string `json:"script_id,omitempty"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int `json:"updated_at,omitempty"`
	// The username of the user who last updated the script
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GlobalInitScriptDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &globalInitScriptDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := globalInitScriptDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GlobalInitScriptDetails) MarshalJSON() ([]byte, error) {
	pb, err := globalInitScriptDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'created_at'
	CreatedAt int `json:"created_at,omitempty"`
	// The username of the user who created the script.
	// Wire name: 'created_by'
	CreatedBy string `json:"created_by,omitempty"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// The Base64-encoded content of the script.
	// Wire name: 'script'
	Script string `json:"script,omitempty"`
	// The global init script ID.
	// Wire name: 'script_id'
	ScriptId string `json:"script_id,omitempty"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int `json:"updated_at,omitempty"`
	// The username of the user who last updated the script
	// Wire name: 'updated_by'
	UpdatedBy string `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GlobalInitScriptDetailsWithContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &globalInitScriptDetailsWithContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := globalInitScriptDetailsWithContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GlobalInitScriptDetailsWithContent) MarshalJSON() ([]byte, error) {
	pb, err := globalInitScriptDetailsWithContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GlobalInitScriptUpdateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the script
	// Wire name: 'name'
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
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// The Base64-encoded content of the script.
	// Wire name: 'script'
	Script string `json:"script"`
	// The ID of the global init script.
	ScriptId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GlobalInitScriptUpdateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &globalInitScriptUpdateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := globalInitScriptUpdateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GlobalInitScriptUpdateRequest) MarshalJSON() ([]byte, error) {
	pb, err := globalInitScriptUpdateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InitScriptEventDetails struct {
	// The cluster scoped init scripts associated with this cluster event.
	// Wire name: 'cluster'
	Cluster []InitScriptInfoAndExecutionDetails `json:"cluster,omitempty"`
	// The global init scripts associated with this cluster event.
	// Wire name: 'global'
	Global []InitScriptInfoAndExecutionDetails `json:"global,omitempty"`
	// The private ip of the node we are reporting init script execution details
	// for (we will select the execution details from only one node rather than
	// reporting the execution details from every node to keep these event
	// details small)
	//
	// This should only be defined for the INIT_SCRIPTS_FINISHED event
	// Wire name: 'reported_for_node'
	ReportedForNode string `json:"reported_for_node,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InitScriptEventDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &initScriptEventDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := initScriptEventDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InitScriptEventDetails) MarshalJSON() ([]byte, error) {
	pb, err := initScriptEventDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'abfss'
	Abfss *Adlsgen2Info `json:"abfss,omitempty"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	// Wire name: 'dbfs'
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	// Wire name: 'file'
	File *LocalFileInfo `json:"file,omitempty"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	// Wire name: 'gcs'
	Gcs *GcsStorageInfo `json:"gcs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	// Wire name: 's3'
	S3 *S3StorageInfo `json:"s3,omitempty"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	// Wire name: 'volumes'
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	// Wire name: 'workspace'
	Workspace *WorkspaceStorageInfo `json:"workspace,omitempty"`
}

func (st *InitScriptInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &initScriptInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := initScriptInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InitScriptInfo) MarshalJSON() ([]byte, error) {
	pb, err := initScriptInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InitScriptInfoAndExecutionDetails struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	// Wire name: 'abfss'
	Abfss *Adlsgen2Info `json:"abfss,omitempty"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	// Wire name: 'dbfs'
	Dbfs *DbfsStorageInfo `json:"dbfs,omitempty"`
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	// Wire name: 'error_message'
	ErrorMessage string `json:"error_message,omitempty"`
	// The number duration of the script execution in seconds
	// Wire name: 'execution_duration_seconds'
	ExecutionDurationSeconds int `json:"execution_duration_seconds,omitempty"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	// Wire name: 'file'
	File *LocalFileInfo `json:"file,omitempty"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	// Wire name: 'gcs'
	Gcs *GcsStorageInfo `json:"gcs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	// Wire name: 's3'
	S3 *S3StorageInfo `json:"s3,omitempty"`
	// The current status of the script
	// Wire name: 'status'
	Status InitScriptExecutionDetailsInitScriptExecutionStatus `json:"status,omitempty"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	// Wire name: 'volumes'
	Volumes *VolumesStorageInfo `json:"volumes,omitempty"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	// Wire name: 'workspace'
	Workspace *WorkspaceStorageInfo `json:"workspace,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InitScriptInfoAndExecutionDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &initScriptInfoAndExecutionDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := initScriptInfoAndExecutionDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InitScriptInfoAndExecutionDetails) MarshalJSON() ([]byte, error) {
	pb, err := initScriptInfoAndExecutionDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// The libraries to install.
	// Wire name: 'libraries'
	Libraries []Library `json:"libraries"`
}

func (st *InstallLibraries) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &installLibrariesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := installLibrariesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstallLibraries) MarshalJSON() ([]byte, error) {
	pb, err := installLibrariesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstallLibrariesResponse struct {
}

func (st *InstallLibrariesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &installLibrariesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := installLibrariesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstallLibrariesResponse) MarshalJSON() ([]byte, error) {
	pb, err := installLibrariesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []InstancePoolPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *InstancePoolAwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *InstancePoolAzureAttributes `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
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
	// Wire name: 'default_tags'
	DefaultTags map[string]string `json:"default_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	// Wire name: 'disk_spec'
	DiskSpec *DiskSpec `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *InstancePoolGcpAttributes `json:"gcp_attributes,omitempty"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int `json:"idle_instance_autotermination_minutes,omitempty"`
	// Canonical unique identifier for the pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string `json:"instance_pool_name,omitempty"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int `json:"max_capacity,omitempty"`
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int `json:"min_idle_instances,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	// Wire name: 'preloaded_docker_images'
	PreloadedDockerImages []DockerImage `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'preloaded_spark_versions'
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Current state of the instance pool.
	// Wire name: 'state'
	State InstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	// Wire name: 'stats'
	Stats *InstancePoolStats `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	// Wire name: 'status'
	Status *InstancePoolStatus `json:"status,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolAndStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolAndStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolAndStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolAndStats) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolAndStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	// Wire name: 'availability'
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
	// Wire name: 'spot_bid_price_percent'
	SpotBidPricePercent int `json:"spot_bid_price_percent,omitempty"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	// Wire name: 'zone_id'
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolAwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolAwsAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolAwsAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolAwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolAwsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'availability'
	Availability InstancePoolAzureAttributesAvailability `json:"availability,omitempty"`
	// With variable pricing, you have option to set a max price, in US dollars
	// (USD) For example, the value 2 would be a max price of $2.00 USD per
	// hour. If you set the max price to be -1, the VM won't be evicted based on
	// price. The price for the VM will be the current price for spot or the
	// price for a standard VM, which ever is less, as long as there is capacity
	// and quota available.
	// Wire name: 'spot_bid_max_price'
	SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolAzureAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolAzureAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolAzureAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolAzureAttributes) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolAzureAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
	// Wire name: 'gcp_availability'
	GcpAvailability GcpAvailability `json:"gcp_availability,omitempty"`
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	// Wire name: 'local_ssd_count'
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
	// Wire name: 'zone_id'
	ZoneId string `json:"zone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolGcpAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolGcpAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolGcpAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolGcpAttributes) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolGcpAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolPermission) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []InstancePoolAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolPermissions) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []InstancePoolAccessControlRequest `json:"access_control_list,omitempty"`
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" tf:"-"`
}

func (st *InstancePoolPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'idle_count'
	IdleCount int `json:"idle_count,omitempty"`
	// Number of pending instances in the pool that are NOT part of a cluster.
	// Wire name: 'pending_idle_count'
	PendingIdleCount int `json:"pending_idle_count,omitempty"`
	// Number of pending instances in the pool that are part of a cluster.
	// Wire name: 'pending_used_count'
	PendingUsedCount int `json:"pending_used_count,omitempty"`
	// Number of active instances in the pool that are part of a cluster.
	// Wire name: 'used_count'
	UsedCount int `json:"used_count,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstancePoolStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolStats) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	// Wire name: 'pending_instance_errors'
	PendingInstanceErrors []PendingInstanceError `json:"pending_instance_errors,omitempty"`
}

func (st *InstancePoolStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instancePoolStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instancePoolStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstancePoolStatus) MarshalJSON() ([]byte, error) {
	pb, err := instancePoolStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'iam_role_arn'
	IamRoleArn string `json:"iam_role_arn,omitempty"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn"`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	// Wire name: 'is_meta_instance_profile'
	IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *InstanceProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &instanceProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := instanceProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InstanceProfile) MarshalJSON() ([]byte, error) {
	pb, err := instanceProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'cran'
	Cran *RCranLibrary `json:"cran,omitempty"`
	// Deprecated. URI of the egg library to install. Installing Python egg
	// files is deprecated and is not supported in Databricks Runtime 14.0 and
	// above.
	// Wire name: 'egg'
	Egg string `json:"egg,omitempty"`
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	// Wire name: 'jar'
	Jar string `json:"jar,omitempty"`
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	// Wire name: 'maven'
	Maven *MavenLibrary `json:"maven,omitempty"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	// Wire name: 'pypi'
	Pypi *PythonPyPiLibrary `json:"pypi,omitempty"`
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	// Wire name: 'requirements'
	Requirements string `json:"requirements,omitempty"`
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	// Wire name: 'whl'
	Whl string `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Library) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &libraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := libraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Library) MarshalJSON() ([]byte, error) {
	pb, err := libraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	// Wire name: 'is_library_for_all_clusters'
	IsLibraryForAllClusters bool `json:"is_library_for_all_clusters,omitempty"`
	// Unique identifier for the library.
	// Wire name: 'library'
	Library *Library `json:"library,omitempty"`
	// All the info and warning messages that have occurred so far for this
	// library.
	// Wire name: 'messages'
	Messages []string `json:"messages,omitempty"`
	// Status of installing the library on the cluster.
	// Wire name: 'status'
	Status LibraryInstallStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *LibraryFullStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &libraryFullStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := libraryFullStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LibraryFullStatus) MarshalJSON() ([]byte, error) {
	pb, err := libraryFullStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'statuses'
	Statuses []ClusterLibraryStatuses `json:"statuses,omitempty"`
}

func (st *ListAllClusterLibraryStatusesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAllClusterLibraryStatusesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAllClusterLibraryStatusesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAllClusterLibraryStatusesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAllClusterLibraryStatusesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAvailableZonesResponse struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	// Wire name: 'default_zone'
	DefaultZone string `json:"default_zone,omitempty"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	// Wire name: 'zones'
	Zones []string `json:"zones,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListAvailableZonesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAvailableZonesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAvailableZonesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAvailableZonesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAvailableZonesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List cluster policy compliance
type ListClusterCompliancesRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int `json:"-" tf:"-"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	PageToken string `json:"-" tf:"-"`
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListClusterCompliancesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClusterCompliancesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClusterCompliancesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClusterCompliancesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listClusterCompliancesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListClusterCompliancesResponse struct {
	// A list of clusters and their policy compliance statuses.
	// Wire name: 'clusters'
	Clusters []ClusterCompliance `json:"clusters,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	// Wire name: 'prev_page_token'
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListClusterCompliancesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClusterCompliancesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClusterCompliancesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClusterCompliancesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listClusterCompliancesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List cluster policies
type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn ListSortColumn `json:"-" tf:"-"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder ListSortOrder `json:"-" tf:"-"`
}

func (st *ListClusterPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClusterPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClusterPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClusterPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listClusterPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListClustersFilterBy struct {
	// The source of cluster creation.
	// Wire name: 'cluster_sources'
	ClusterSources []ClusterSource `json:"cluster_sources,omitempty"`
	// The current state of the clusters.
	// Wire name: 'cluster_states'
	ClusterStates []State `json:"cluster_states,omitempty"`
	// Whether the clusters are pinned or not.
	// Wire name: 'is_pinned'
	IsPinned bool `json:"is_pinned,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListClustersFilterBy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClustersFilterByPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClustersFilterByFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClustersFilterBy) MarshalJSON() ([]byte, error) {
	pb, err := listClustersFilterByToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List clusters
type ListClustersRequest struct {
	// Filters to apply to the list of clusters.
	FilterBy *ListClustersFilterBy `json:"-" tf:"-"`
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int `json:"-" tf:"-"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	PageToken string `json:"-" tf:"-"`
	// Sort the list of clusters by a specific criteria.
	SortBy *ListClustersSortBy `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListClustersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClustersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClustersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClustersRequest) MarshalJSON() ([]byte, error) {
	pb, err := listClustersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListClustersResponse struct {

	// Wire name: 'clusters'
	Clusters []ClusterDetails `json:"clusters,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	// Wire name: 'prev_page_token'
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListClustersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClustersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClustersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClustersResponse) MarshalJSON() ([]byte, error) {
	pb, err := listClustersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListClustersSortBy struct {
	// The direction to sort by.
	// Wire name: 'direction'
	Direction ListClustersSortByDirection `json:"direction,omitempty"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	// Wire name: 'field'
	Field ListClustersSortByField `json:"field,omitempty"`
}

func (st *ListClustersSortBy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listClustersSortByPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listClustersSortByFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListClustersSortBy) MarshalJSON() ([]byte, error) {
	pb, err := listClustersSortByToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'scripts'
	Scripts []GlobalInitScriptDetails `json:"scripts,omitempty"`
}

func (st *ListGlobalInitScriptsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listGlobalInitScriptsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listGlobalInitScriptsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListGlobalInitScriptsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listGlobalInitScriptsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListInstancePools struct {

	// Wire name: 'instance_pools'
	InstancePools []InstancePoolAndStats `json:"instance_pools,omitempty"`
}

func (st *ListInstancePools) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listInstancePoolsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listInstancePoolsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListInstancePools) MarshalJSON() ([]byte, error) {
	pb, err := listInstancePoolsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	// Wire name: 'instance_profiles'
	InstanceProfiles []InstanceProfile `json:"instance_profiles,omitempty"`
}

func (st *ListInstanceProfilesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listInstanceProfilesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listInstanceProfilesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListInstanceProfilesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listInstanceProfilesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	// Wire name: 'node_types'
	NodeTypes []NodeType `json:"node_types,omitempty"`
}

func (st *ListNodeTypesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNodeTypesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNodeTypesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNodeTypesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listNodeTypesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListPoliciesResponse struct {
	// List of policies.
	// Wire name: 'policies'
	Policies []Policy `json:"policies,omitempty"`
}

func (st *ListPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List policy families
type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	MaxResults int64 `json:"-" tf:"-"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPolicyFamiliesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPolicyFamiliesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPolicyFamiliesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPolicyFamiliesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listPolicyFamiliesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of policy families.
	// Wire name: 'policy_families'
	PolicyFamilies []PolicyFamily `json:"policy_families,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPolicyFamiliesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPolicyFamiliesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPolicyFamiliesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPolicyFamiliesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listPolicyFamiliesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'destination'
	Destination string `json:"destination"`
}

func (st *LocalFileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &localFileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := localFileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LocalFileInfo) MarshalJSON() ([]byte, error) {
	pb, err := localFileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogAnalyticsInfo struct {

	// Wire name: 'log_analytics_primary_key'
	LogAnalyticsPrimaryKey string `json:"log_analytics_primary_key,omitempty"`

	// Wire name: 'log_analytics_workspace_id'
	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *LogAnalyticsInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logAnalyticsInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logAnalyticsInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogAnalyticsInfo) MarshalJSON() ([]byte, error) {
	pb, err := logAnalyticsInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The log delivery status
type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	// Wire name: 'last_attempted'
	LastAttempted int64 `json:"last_attempted,omitempty"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	// Wire name: 'last_exception'
	LastException string `json:"last_exception,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *LogSyncStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logSyncStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logSyncStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogSyncStatus) MarshalJSON() ([]byte, error) {
	pb, err := logSyncStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MapAny map[string]any

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	// Wire name: 'coordinates'
	Coordinates string `json:"coordinates"`
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	// Wire name: 'exclusions'
	Exclusions []string `json:"exclusions,omitempty"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	// Wire name: 'repo'
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *MavenLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mavenLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := mavenLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MavenLibrary) MarshalJSON() ([]byte, error) {
	pb, err := mavenLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// This structure embodies the machine type that hosts spark containers Note:
// this should be an internal data structure for now It is defined in proto in
// case we want to send it over the wire in the future (which is likely)
type NodeInstanceType struct {
	// Unique identifier across instance types
	// Wire name: 'instance_type_id'
	InstanceTypeId string `json:"instance_type_id"`
	// Size of the individual local disks attached to this instance (i.e. per
	// local disk).
	// Wire name: 'local_disk_size_gb'
	LocalDiskSizeGb int `json:"local_disk_size_gb,omitempty"`
	// Number of local disks that are present on this instance.
	// Wire name: 'local_disks'
	LocalDisks int `json:"local_disks,omitempty"`
	// Size of the individual local nvme disks attached to this instance (i.e.
	// per local disk).
	// Wire name: 'local_nvme_disk_size_gb'
	LocalNvmeDiskSizeGb int `json:"local_nvme_disk_size_gb,omitempty"`
	// Number of local nvme disks that are present on this instance.
	// Wire name: 'local_nvme_disks'
	LocalNvmeDisks int `json:"local_nvme_disks,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NodeInstanceType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nodeInstanceTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nodeInstanceTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NodeInstanceType) MarshalJSON() ([]byte, error) {
	pb, err := nodeInstanceTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A description of a Spark node type including both the dimensions of the node
// and the instance type on which it will be hosted.
type NodeType struct {
	// A descriptive category for this node type. Examples include "Memory
	// Optimized" and "Compute Optimized".
	// Wire name: 'category'
	Category string `json:"category"`
	// A string description associated with this node type, e.g., "r3.xlarge".
	// Wire name: 'description'
	Description string `json:"description"`
	// An optional hint at the display order of node types in the UI. Within a
	// node type category, lowest numbers come first.
	// Wire name: 'display_order'
	DisplayOrder int `json:"display_order,omitempty"`
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	// Wire name: 'instance_type_id'
	InstanceTypeId string `json:"instance_type_id"`
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	// Wire name: 'is_deprecated'
	IsDeprecated bool `json:"is_deprecated,omitempty"`
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	// Wire name: 'is_encrypted_in_transit'
	IsEncryptedInTransit bool `json:"is_encrypted_in_transit,omitempty"`
	// Whether this is an Arm-based instance.
	// Wire name: 'is_graviton'
	IsGraviton bool `json:"is_graviton,omitempty"`
	// Whether this node is hidden from presentation in the UI.
	// Wire name: 'is_hidden'
	IsHidden bool `json:"is_hidden,omitempty"`
	// Whether this node comes with IO cache enabled by default.
	// Wire name: 'is_io_cache_enabled'
	IsIoCacheEnabled bool `json:"is_io_cache_enabled,omitempty"`
	// Memory (in MB) available for this node type.
	// Wire name: 'memory_mb'
	MemoryMb int `json:"memory_mb"`
	// A collection of node type info reported by the cloud provider
	// Wire name: 'node_info'
	NodeInfo *CloudProviderNodeInfo `json:"node_info,omitempty"`
	// The NodeInstanceType object corresponding to instance_type_id
	// Wire name: 'node_instance_type'
	NodeInstanceType *NodeInstanceType `json:"node_instance_type,omitempty"`
	// Unique identifier for this node type.
	// Wire name: 'node_type_id'
	NodeTypeId string `json:"node_type_id"`
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	// Wire name: 'num_cores'
	NumCores float64 `json:"num_cores"`
	// Number of GPUs available for this node type.
	// Wire name: 'num_gpus'
	NumGpus int `json:"num_gpus,omitempty"`

	// Wire name: 'photon_driver_capable'
	PhotonDriverCapable bool `json:"photon_driver_capable,omitempty"`

	// Wire name: 'photon_worker_capable'
	PhotonWorkerCapable bool `json:"photon_worker_capable,omitempty"`
	// Whether this node type support cluster tags.
	// Wire name: 'support_cluster_tags'
	SupportClusterTags bool `json:"support_cluster_tags,omitempty"`
	// Whether this node type support EBS volumes. EBS volumes is disabled for
	// node types that we could place multiple corresponding containers on the
	// same hosting instance.
	// Wire name: 'support_ebs_volumes'
	SupportEbsVolumes bool `json:"support_ebs_volumes,omitempty"`
	// Whether this node type supports port forwarding.
	// Wire name: 'support_port_forwarding'
	SupportPortForwarding bool `json:"support_port_forwarding,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NodeType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nodeTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nodeTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NodeType) MarshalJSON() ([]byte, error) {
	pb, err := nodeTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Error message of a failed pending instances
type PendingInstanceError struct {

	// Wire name: 'instance_id'
	InstanceId string `json:"instance_id,omitempty"`

	// Wire name: 'message'
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PendingInstanceError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pendingInstanceErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pendingInstanceErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PendingInstanceError) MarshalJSON() ([]byte, error) {
	pb, err := pendingInstanceErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PermanentDeleteCluster struct {
	// The cluster to be deleted.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
}

func (st *PermanentDeleteCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permanentDeleteClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permanentDeleteClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermanentDeleteCluster) MarshalJSON() ([]byte, error) {
	pb, err := permanentDeleteClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PermanentDeleteClusterResponse struct {
}

func (st *PermanentDeleteClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &permanentDeleteClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := permanentDeleteClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PermanentDeleteClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := permanentDeleteClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PinCluster struct {

	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
}

func (st *PinCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pinClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pinClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PinCluster) MarshalJSON() ([]byte, error) {
	pb, err := pinClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PinClusterResponse struct {
}

func (st *PinClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pinClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pinClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PinClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := pinClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Describes a Cluster Policy entity.
type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	// Wire name: 'created_at_timestamp'
	CreatedAtTimestamp int64 `json:"created_at_timestamp,omitempty"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	// Wire name: 'is_default'
	IsDefault bool `json:"is_default,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	// Wire name: 'libraries'
	Libraries []Library `json:"libraries,omitempty"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	// Wire name: 'max_clusters_per_user'
	MaxClustersPerUser int64 `json:"max_clusters_per_user,omitempty"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	// Wire name: 'name'
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
	// Wire name: 'policy_family_definition_overrides'
	PolicyFamilyDefinitionOverrides string `json:"policy_family_definition_overrides,omitempty"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string `json:"policy_family_id,omitempty"`
	// Canonical unique identifier for the Cluster Policy.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Policy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &policyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := policyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Policy) MarshalJSON() ([]byte, error) {
	pb, err := policyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string `json:"definition,omitempty"`
	// Human-readable description of the purpose of the policy family.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Name of the policy family.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Unique identifier for the policy family.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string `json:"policy_family_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PolicyFamily) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &policyFamilyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := policyFamilyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PolicyFamily) MarshalJSON() ([]byte, error) {
	pb, err := policyFamilyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	// Wire name: 'package'
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	// Wire name: 'repo'
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PythonPyPiLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pythonPyPiLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pythonPyPiLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PythonPyPiLibrary) MarshalJSON() ([]byte, error) {
	pb, err := pythonPyPiLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	// Wire name: 'package'
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	// Wire name: 'repo'
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RCranLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rCranLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rCranLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RCranLibrary) MarshalJSON() ([]byte, error) {
	pb, err := rCranLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn"`
}

func (st *RemoveInstanceProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &removeInstanceProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := removeInstanceProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RemoveInstanceProfile) MarshalJSON() ([]byte, error) {
	pb, err := removeInstanceProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RemoveResponse struct {
}

func (st *RemoveResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &removeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := removeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RemoveResponse) MarshalJSON() ([]byte, error) {
	pb, err := removeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResizeCluster struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// The cluster to be resized.
	// Wire name: 'cluster_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ResizeCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resizeClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resizeClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResizeCluster) MarshalJSON() ([]byte, error) {
	pb, err := resizeClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResizeClusterResponse struct {
}

func (st *ResizeClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resizeClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resizeClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResizeClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := resizeClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestartCluster struct {
	// The cluster to be started.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`

	// Wire name: 'restart_user'
	RestartUser string `json:"restart_user,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RestartCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restartClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restartClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestartCluster) MarshalJSON() ([]byte, error) {
	pb, err := restartClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestartClusterResponse struct {
}

func (st *RestartClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restartClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restartClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestartClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := restartClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'cause'
	Cause string `json:"cause,omitempty"`

	// Wire name: 'data'
	Data any `json:"data,omitempty"`
	// The image filename
	// Wire name: 'fileName'
	FileName string `json:"fileName,omitempty"`

	// Wire name: 'fileNames'
	FileNames []string `json:"fileNames,omitempty"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	// Wire name: 'isJsonSchema'
	IsJsonSchema bool `json:"isJsonSchema,omitempty"`
	// internal field used by SDK
	// Wire name: 'pos'
	Pos int `json:"pos,omitempty"`

	// Wire name: 'resultType'
	ResultType ResultType `json:"resultType,omitempty"`
	// The table schema
	// Wire name: 'schema'
	Schema []map[string]any `json:"schema,omitempty"`
	// The summary of the error
	// Wire name: 'summary'
	Summary string `json:"summary,omitempty"`
	// true if partial results are returned.
	// Wire name: 'truncated'
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Results) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Results) MarshalJSON() ([]byte, error) {
	pb, err := resultsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'canned_acl'
	CannedAcl string `json:"canned_acl,omitempty"`
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	// Wire name: 'destination'
	Destination string `json:"destination"`
	// (Optional) Flag to enable server side encryption, `false` by default.
	// Wire name: 'enable_encryption'
	EnableEncryption bool `json:"enable_encryption,omitempty"`
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	// Wire name: 'encryption_type'
	EncryptionType string `json:"encryption_type,omitempty"`
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	// Wire name: 'endpoint'
	Endpoint string `json:"endpoint,omitempty"`
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	// Wire name: 'kms_key'
	KmsKey string `json:"kms_key,omitempty"`
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *S3StorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &s3StorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := s3StorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st S3StorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := s3StorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Describes a specific Spark driver or executor.
type SparkNode struct {
	// The private IP address of the host instance.
	// Wire name: 'host_private_ip'
	HostPrivateIp string `json:"host_private_ip,omitempty"`
	// Globally unique identifier for the host instance from the cloud provider.
	// Wire name: 'instance_id'
	InstanceId string `json:"instance_id,omitempty"`
	// Attributes specific to AWS for a Spark node.
	// Wire name: 'node_aws_attributes'
	NodeAwsAttributes *SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
	// Globally unique identifier for this node.
	// Wire name: 'node_id'
	NodeId string `json:"node_id,omitempty"`
	// Private IP address (typically a 10.x.x.x address) of the Spark node. Note
	// that this is different from the private IP address of the host instance.
	// Wire name: 'private_ip'
	PrivateIp string `json:"private_ip,omitempty"`
	// Public DNS address of this node. This address can be used to access the
	// Spark JDBC server on the driver node. To communicate with the JDBC
	// server, traffic must be manually authorized by adding security group
	// rules to the "worker-unmanaged" security group via the AWS console.
	// Wire name: 'public_dns'
	PublicDns string `json:"public_dns,omitempty"`
	// The timestamp (in millisecond) when the Spark node is launched.
	// Wire name: 'start_timestamp'
	StartTimestamp int64 `json:"start_timestamp,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SparkNode) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkNodePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkNodeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkNode) MarshalJSON() ([]byte, error) {
	pb, err := sparkNodeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	// Wire name: 'is_spot'
	IsSpot bool `json:"is_spot,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SparkNodeAwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkNodeAwsAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkNodeAwsAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkNodeAwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := sparkNodeAwsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SparkVersion struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	// Wire name: 'key'
	Key string `json:"key,omitempty"`
	// A descriptive name for this Spark version, for example "Spark 2.1".
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SparkVersion) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkVersionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkVersionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkVersion) MarshalJSON() ([]byte, error) {
	pb, err := sparkVersionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StartCluster struct {
	// The cluster to be started.
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
}

func (st *StartCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartCluster) MarshalJSON() ([]byte, error) {
	pb, err := startClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StartClusterResponse struct {
}

func (st *StartClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := startClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'code'
	Code TerminationReasonCode `json:"code,omitempty"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	// Wire name: 'parameters'
	Parameters map[string]string `json:"parameters,omitempty"`
	// type of the termination
	// Wire name: 'type'
	Type TerminationReasonType `json:"type,omitempty"`
}

func (st *TerminationReason) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &terminationReasonPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := terminationReasonFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TerminationReason) MarshalJSON() ([]byte, error) {
	pb, err := terminationReasonToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	case `ABUSE_DETECTED`, `ACCESS_TOKEN_FAILURE`, `ALLOCATION_TIMEOUT`, `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`, `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INACCESSIBLE_KMS_KEY_FAILURE`, `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_INVALID_KEY_PAIR`, `AWS_INVALID_KMS_KEY_STATE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_RESOURCE_QUOTA_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`, `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`, `BUDGET_POLICY_RESOLUTION_FAILURE`, `CLOUD_ACCOUNT_SETUP_FAILURE`, `CLOUD_OPERATION_CANCELLED`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_SHUTDOWN`, `CLUSTER_OPERATION_THROTTLED`, `CLUSTER_OPERATION_TIMEOUT`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`, `DATABASE_CONNECTION_FAILURE`, `DATA_ACCESS_CONFIG_CHANGED`, `DBFS_COMPONENT_UNHEALTHY`, `DISASTER_RECOVERY_REPLICATION`, `DNS_RESOLUTION_ERROR`, `DOCKER_CONTAINER_CREATION_EXCEPTION`, `DOCKER_IMAGE_PULL_FAILURE`, `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`, `DOCKER_INVALID_OS_EXCEPTION`, `DRIVER_EVICTION`, `DRIVER_LAUNCH_TIMEOUT`, `DRIVER_NODE_UNREACHABLE`, `DRIVER_OUT_OF_DISK`, `DRIVER_OUT_OF_MEMORY`, `DRIVER_POD_CREATION_FAILURE`, `DRIVER_UNEXPECTED_FAILURE`, `DRIVER_UNHEALTHY`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`, `EOS_SPARK_IMAGE`, `EXECUTION_COMPONENT_UNHEALTHY`, `EXECUTOR_POD_UNSCHEDULED`, `GCP_API_RATE_QUOTA_EXCEEDED`, `GCP_DENIED_BY_ORG_POLICY`, `GCP_FORBIDDEN`, `GCP_IAM_TIMEOUT`, `GCP_INACCESSIBLE_KMS_KEY_FAILURE`, `GCP_INSUFFICIENT_CAPACITY`, `GCP_IP_SPACE_EXHAUSTED`, `GCP_KMS_KEY_PERMISSION_DENIED`, `GCP_NOT_FOUND`, `GCP_QUOTA_EXCEEDED`, `GCP_RESOURCE_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GCP_SERVICE_ACCOUNT_NOT_FOUND`, `GCP_SUBNET_NOT_READY`, `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`, `GKE_BASED_CLUSTER_TERMINATION`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_CONTAINER_NOT_FINISHED`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_POOL_MAX_CAPACITY_REACHED`, `INSTANCE_POOL_NOT_FOUND`, `INSTANCE_UNREACHABLE`, `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`, `INTERNAL_CAPACITY_FAILURE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_AWS_PARAMETER`, `INVALID_INSTANCE_PLACEMENT_PROTOCOL`, `INVALID_SPARK_IMAGE`, `INVALID_WORKER_IMAGE_FAILURE`, `IN_PENALTY_BOX`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `LAZY_ALLOCATION_TIMEOUT`, `MAINTENANCE_MODE`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETVISOR_SETUP_TIMEOUT`, `NETWORK_CHECK_CONTROL_PLANE_FAILURE`, `NETWORK_CHECK_DNS_SERVER_FAILURE`, `NETWORK_CHECK_METADATA_ENDPOINT_FAILURE`, `NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE`, `NETWORK_CHECK_NIC_FAILURE`, `NETWORK_CHECK_STORAGE_FAILURE`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NO_MATCHED_K8S`, `NO_MATCHED_K8S_TESTING_TAG`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `POD_ASSIGNMENT_FAILURE`, `POD_SCHEDULING_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `RESOURCE_USAGE_BLOCKED`, `SECRET_CREATION_FAILURE`, `SECRET_PERMISSION_DENIED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SERVERLESS_LONG_RUNNING_TERMINATED`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_IMAGE_DOWNLOAD_THROTTLED`, `SPARK_IMAGE_NOT_FOUND`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `SSH_BOOTSTRAP_FAILURE`, `STORAGE_DOWNLOAD_FAILURE`, `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`, `STORAGE_DOWNLOAD_FAILURE_SLOW`, `STORAGE_DOWNLOAD_FAILURE_THROTTLED`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNEXPECTED_POD_RECREATION`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_INITIATED_VM_TERMINATION`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`, `WORKSPACE_UPDATE`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ACCESS_TOKEN_FAILURE", "ALLOCATION_TIMEOUT", "ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY", "ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS", "ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_READY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INACCESSIBLE_KMS_KEY_FAILURE", "AWS_INSTANCE_PROFILE_UPDATE_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_INVALID_KEY_PAIR", "AWS_INVALID_KMS_KEY_STATE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_RESOURCE_QUOTA_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG", "BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED", "BUDGET_POLICY_RESOLUTION_FAILURE", "CLOUD_ACCOUNT_SETUP_FAILURE", "CLOUD_OPERATION_CANCELLED", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_SHUTDOWN", "CLUSTER_OPERATION_THROTTLED", "CLUSTER_OPERATION_TIMEOUT", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG", "DATABASE_CONNECTION_FAILURE", "DATA_ACCESS_CONFIG_CHANGED", "DBFS_COMPONENT_UNHEALTHY", "DISASTER_RECOVERY_REPLICATION", "DNS_RESOLUTION_ERROR", "DOCKER_CONTAINER_CREATION_EXCEPTION", "DOCKER_IMAGE_PULL_FAILURE", "DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION", "DOCKER_INVALID_OS_EXCEPTION", "DRIVER_EVICTION", "DRIVER_LAUNCH_TIMEOUT", "DRIVER_NODE_UNREACHABLE", "DRIVER_OUT_OF_DISK", "DRIVER_OUT_OF_MEMORY", "DRIVER_POD_CREATION_FAILURE", "DRIVER_UNEXPECTED_FAILURE", "DRIVER_UNHEALTHY", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "DYNAMIC_SPARK_CONF_SIZE_EXCEEDED", "EOS_SPARK_IMAGE", "EXECUTION_COMPONENT_UNHEALTHY", "EXECUTOR_POD_UNSCHEDULED", "GCP_API_RATE_QUOTA_EXCEEDED", "GCP_DENIED_BY_ORG_POLICY", "GCP_FORBIDDEN", "GCP_IAM_TIMEOUT", "GCP_INACCESSIBLE_KMS_KEY_FAILURE", "GCP_INSUFFICIENT_CAPACITY", "GCP_IP_SPACE_EXHAUSTED", "GCP_KMS_KEY_PERMISSION_DENIED", "GCP_NOT_FOUND", "GCP_QUOTA_EXCEEDED", "GCP_RESOURCE_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_ACCESS_DENIED", "GCP_SERVICE_ACCOUNT_DELETED", "GCP_SERVICE_ACCOUNT_NOT_FOUND", "GCP_SUBNET_NOT_READY", "GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED", "GKE_BASED_CLUSTER_TERMINATION", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_CONTAINER_NOT_FINISHED", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_POOL_MAX_CAPACITY_REACHED", "INSTANCE_POOL_NOT_FOUND", "INSTANCE_UNREACHABLE", "INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG", "INTERNAL_CAPACITY_FAILURE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_AWS_PARAMETER", "INVALID_INSTANCE_PLACEMENT_PROTOCOL", "INVALID_SPARK_IMAGE", "INVALID_WORKER_IMAGE_FAILURE", "IN_PENALTY_BOX", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "LAZY_ALLOCATION_TIMEOUT", "MAINTENANCE_MODE", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETVISOR_SETUP_TIMEOUT", "NETWORK_CHECK_CONTROL_PLANE_FAILURE", "NETWORK_CHECK_DNS_SERVER_FAILURE", "NETWORK_CHECK_METADATA_ENDPOINT_FAILURE", "NETWORK_CHECK_MULTIPLE_COMPONENTS_FAILURE", "NETWORK_CHECK_NIC_FAILURE", "NETWORK_CHECK_STORAGE_FAILURE", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NO_MATCHED_K8S", "NO_MATCHED_K8S_TESTING_TAG", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "POD_ASSIGNMENT_FAILURE", "POD_SCHEDULING_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "RESOURCE_USAGE_BLOCKED", "SECRET_CREATION_FAILURE", "SECRET_PERMISSION_DENIED", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SERVERLESS_LONG_RUNNING_TERMINATED", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_IMAGE_DOWNLOAD_THROTTLED", "SPARK_IMAGE_NOT_FOUND", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "SSH_BOOTSTRAP_FAILURE", "STORAGE_DOWNLOAD_FAILURE", "STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG", "STORAGE_DOWNLOAD_FAILURE_SLOW", "STORAGE_DOWNLOAD_FAILURE_THROTTLED", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNEXPECTED_POD_RECREATION", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_INITIATED_VM_TERMINATION", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR", "WORKSPACE_UPDATE"`, v)
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
	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
	// The libraries to uninstall.
	// Wire name: 'libraries'
	Libraries []Library `json:"libraries"`
}

func (st *UninstallLibraries) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &uninstallLibrariesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := uninstallLibrariesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UninstallLibraries) MarshalJSON() ([]byte, error) {
	pb, err := uninstallLibrariesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UninstallLibrariesResponse struct {
}

func (st *UninstallLibrariesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &uninstallLibrariesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := uninstallLibrariesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UninstallLibrariesResponse) MarshalJSON() ([]byte, error) {
	pb, err := uninstallLibrariesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UnpinCluster struct {

	// Wire name: 'cluster_id'
	ClusterId string `json:"cluster_id"`
}

func (st *UnpinCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &unpinClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := unpinClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UnpinCluster) MarshalJSON() ([]byte, error) {
	pb, err := unpinClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UnpinClusterResponse struct {
}

func (st *UnpinClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &unpinClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := unpinClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UnpinClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := unpinClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateCluster struct {
	// The cluster to be updated.
	// Wire name: 'cluster'
	Cluster *UpdateClusterResource `json:"cluster,omitempty"`
	// ID of the cluster.
	// Wire name: 'cluster_id'
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
	// Wire name: 'update_mask'
	UpdateMask string `json:"update_mask"`
}

func (st *UpdateCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCluster) MarshalJSON() ([]byte, error) {
	pb, err := updateClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateClusterResource struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string `json:"cluster_name,omitempty"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// The following modes can only be used when `kind = CLASSIC_PREVIEW`. *
	// `DATA_SECURITY_MODE_AUTO`: Databricks will choose the most appropriate
	// access mode depending on your compute configuration. *
	// `DATA_SECURITY_MODE_STANDARD`: Alias for `USER_ISOLATION`. *
	// `DATA_SECURITY_MODE_DEDICATED`: Alias for `SINGLE_USER`.
	//
	// The following modes can be used regardless of `kind`. * `NONE`: No
	// security isolation for multiple users sharing the cluster. Data
	// governance features are not available in this mode. * `SINGLE_USER`: A
	// secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage `json:"docker_image,omitempty"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string `json:"driver_instance_pool_id,omitempty"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string `json:"driver_node_type_id,omitempty"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool `json:"enable_local_disk_encryption,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool `json:"is_single_node,omitempty"`
	// The kind of compute described by this compute specification.
	//
	// Depending on `kind`, different validations and default values will be
	// applied.
	//
	// Clusters with `kind = CLASSIC_PREVIEW` support the following fields,
	// whereas clusters with no specified `kind` do not. *
	// [is_single_node](/api/workspace/clusters/create#is_single_node) *
	// [use_ml_runtime](/api/workspace/clusters/create#use_ml_runtime) *
	// [data_security_mode](/api/workspace/clusters/create#data_security_mode)
	// set to `DATA_SECURITY_MODE_AUTO`, `DATA_SECURITY_MODE_DEDICATED`, or
	// `DATA_SECURITY_MODE_STANDARD`
	//
	// By using the [simple form], your clusters are automatically using `kind =
	// CLASSIC_PREVIEW`.
	//
	// [simple form]: https://docs.databricks.com/compute/simple-form.html
	// Wire name: 'kind'
	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
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
	// Wire name: 'num_workers'
	NumWorkers int `json:"num_workers,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int `json:"remote_disk_throughput,omitempty"`
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine `json:"runtime_engine,omitempty"`
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string `json:"single_user_name,omitempty"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
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
	// Wire name: 'spark_env_vars'
	SparkEnvVars map[string]string `json:"spark_env_vars,omitempty"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string `json:"spark_version,omitempty"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string `json:"ssh_public_keys,omitempty"`
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int `json:"total_initial_remote_disk_size,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	// Wire name: 'workload_type'
	WorkloadType *WorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateClusterResource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateClusterResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateClusterResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateClusterResource) MarshalJSON() ([]byte, error) {
	pb, err := updateClusterResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateClusterResponse struct {
}

func (st *UpdateClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateResponse struct {
}

func (st *UpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A storage location back by UC Volumes.
type VolumesStorageInfo struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	// Wire name: 'destination'
	Destination string `json:"destination"`
}

func (st *VolumesStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &volumesStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := volumesStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st VolumesStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := volumesStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Cluster Attributes showing for clusters workload types.
type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	// Wire name: 'clients'
	Clients ClientsTypes `json:"clients"`
}

func (st *WorkloadType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workloadTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workloadTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkloadType) MarshalJSON() ([]byte, error) {
	pb, err := workloadTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A storage location in Workspace Filesystem (WSFS)
type WorkspaceStorageInfo struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	// Wire name: 'destination'
	Destination string `json:"destination"`
}

func (st *WorkspaceStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := workspaceStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
