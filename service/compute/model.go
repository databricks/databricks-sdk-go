// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
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
	IamRoleArn string ``
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	// Wire name: 'is_meta_instance_profile'
	IsMetaInstanceProfile bool ``
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	// Wire name: 'skip_validation'
	SkipValidation  bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st AddInstanceProfile) MarshalJSON() ([]byte, error) {
	pb, err := AddInstanceProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AddInstanceProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.AddInstanceProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AddInstanceProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AddInstanceProfileToPb(st *AddInstanceProfile) (*computepb.AddInstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.AddInstanceProfilePb{}
	pb.IamRoleArn = st.IamRoleArn
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.IsMetaInstanceProfile = st.IsMetaInstanceProfile
	pb.SkipValidation = st.SkipValidation

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AddInstanceProfileFromPb(pb *computepb.AddInstanceProfilePb) (*AddInstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddInstanceProfile{}
	st.IamRoleArn = pb.IamRoleArn
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.IsMetaInstanceProfile = pb.IsMetaInstanceProfile
	st.SkipValidation = pb.SkipValidation

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A storage location in Adls Gen2
type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	// Wire name: 'destination'
	Destination string ``
}

func (st Adlsgen2Info) MarshalJSON() ([]byte, error) {
	pb, err := Adlsgen2InfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Adlsgen2Info) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.Adlsgen2InfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := Adlsgen2InfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func Adlsgen2InfoToPb(st *Adlsgen2Info) (*computepb.Adlsgen2InfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.Adlsgen2InfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

func Adlsgen2InfoFromPb(pb *computepb.Adlsgen2InfoPb) (*Adlsgen2Info, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Adlsgen2Info{}
	st.Destination = pb.Destination

	return st, nil
}

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	// Wire name: 'max_workers'
	MaxWorkers int ``
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	// Wire name: 'min_workers'
	MinWorkers      int      ``
	ForceSendFields []string `tf:"-"`
}

func (st AutoScale) MarshalJSON() ([]byte, error) {
	pb, err := AutoScaleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AutoScale) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.AutoScalePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AutoScaleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AutoScaleToPb(st *AutoScale) (*computepb.AutoScalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.AutoScalePb{}
	pb.MaxWorkers = st.MaxWorkers
	pb.MinWorkers = st.MinWorkers

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AutoScaleFromPb(pb *computepb.AutoScalePb) (*AutoScale, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoScale{}
	st.MaxWorkers = pb.MaxWorkers
	st.MinWorkers = pb.MinWorkers

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Attributes set during cluster creation which are related to Amazon Web
// Services.
type AwsAttributes struct {

	// Wire name: 'availability'
	Availability AwsAvailability ``
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
	EbsVolumeCount int ``
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	// Wire name: 'ebs_volume_iops'
	EbsVolumeIops int ``
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	// Wire name: 'ebs_volume_size'
	EbsVolumeSize int ``
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	// Wire name: 'ebs_volume_throughput'
	EbsVolumeThroughput int ``
	// The type of EBS volumes that will be launched with this cluster.
	// Wire name: 'ebs_volume_type'
	EbsVolumeType EbsVolumeType ``
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
	FirstOnDemand int ``
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
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
	SpotBidPricePercent int ``
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
	ZoneId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st AwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := AwsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.AwsAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AwsAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AwsAttributesToPb(st *AwsAttributes) (*computepb.AwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.AwsAttributesPb{}
	availabilityPb, err := AwsAvailabilityToPb(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}
	pb.EbsVolumeCount = st.EbsVolumeCount
	pb.EbsVolumeIops = st.EbsVolumeIops
	pb.EbsVolumeSize = st.EbsVolumeSize
	pb.EbsVolumeThroughput = st.EbsVolumeThroughput
	ebsVolumeTypePb, err := EbsVolumeTypeToPb(&st.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypePb != nil {
		pb.EbsVolumeType = *ebsVolumeTypePb
	}
	pb.FirstOnDemand = st.FirstOnDemand
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.SpotBidPricePercent = st.SpotBidPricePercent
	pb.ZoneId = st.ZoneId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AwsAttributesFromPb(pb *computepb.AwsAttributesPb) (*AwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsAttributes{}
	availabilityField, err := AwsAvailabilityFromPb(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	st.EbsVolumeCount = pb.EbsVolumeCount
	st.EbsVolumeIops = pb.EbsVolumeIops
	st.EbsVolumeSize = pb.EbsVolumeSize
	st.EbsVolumeThroughput = pb.EbsVolumeThroughput
	ebsVolumeTypeField, err := EbsVolumeTypeFromPb(&pb.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypeField != nil {
		st.EbsVolumeType = *ebsVolumeTypeField
	}
	st.FirstOnDemand = pb.FirstOnDemand
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.SpotBidPricePercent = pb.SpotBidPricePercent
	st.ZoneId = pb.ZoneId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func AwsAvailabilityToPb(st *AwsAvailability) (*computepb.AwsAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.AwsAvailabilityPb(*st)
	return &pb, nil
}

func AwsAvailabilityFromPb(pb *computepb.AwsAvailabilityPb) (*AwsAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := AwsAvailability(*pb)
	return &st, nil
}

// Attributes set during cluster creation which are related to Microsoft Azure.
type AzureAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones. Note: If `first_on_demand` is zero, this
	// availability type will be used for the entire cluster.
	// Wire name: 'availability'
	Availability AzureAvailability ``
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
	FirstOnDemand int ``
	// Defines values necessary to configure and run Azure Log Analytics agent
	// Wire name: 'log_analytics_info'
	LogAnalyticsInfo *LogAnalyticsInfo ``
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	// Wire name: 'spot_bid_max_price'
	SpotBidMaxPrice float64  ``
	ForceSendFields []string `tf:"-"`
}

func (st AzureAttributes) MarshalJSON() ([]byte, error) {
	pb, err := AzureAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AzureAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.AzureAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AzureAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AzureAttributesToPb(st *AzureAttributes) (*computepb.AzureAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.AzureAttributesPb{}
	availabilityPb, err := AzureAvailabilityToPb(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}
	pb.FirstOnDemand = st.FirstOnDemand
	logAnalyticsInfoPb, err := LogAnalyticsInfoToPb(st.LogAnalyticsInfo)
	if err != nil {
		return nil, err
	}
	if logAnalyticsInfoPb != nil {
		pb.LogAnalyticsInfo = logAnalyticsInfoPb
	}
	pb.SpotBidMaxPrice = st.SpotBidMaxPrice

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AzureAttributesFromPb(pb *computepb.AzureAttributesPb) (*AzureAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureAttributes{}
	availabilityField, err := AzureAvailabilityFromPb(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	st.FirstOnDemand = pb.FirstOnDemand
	logAnalyticsInfoField, err := LogAnalyticsInfoFromPb(pb.LogAnalyticsInfo)
	if err != nil {
		return nil, err
	}
	if logAnalyticsInfoField != nil {
		st.LogAnalyticsInfo = logAnalyticsInfoField
	}
	st.SpotBidMaxPrice = pb.SpotBidMaxPrice

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func AzureAvailabilityToPb(st *AzureAvailability) (*computepb.AzureAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.AzureAvailabilityPb(*st)
	return &pb, nil
}

func AzureAvailabilityFromPb(pb *computepb.AzureAvailabilityPb) (*AzureAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := AzureAvailability(*pb)
	return &st, nil
}

type CancelCommand struct {

	// Wire name: 'clusterId'
	ClusterId string ``

	// Wire name: 'commandId'
	CommandId string ``

	// Wire name: 'contextId'
	ContextId       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CancelCommand) MarshalJSON() ([]byte, error) {
	pb, err := CancelCommandToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CancelCommand) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CancelCommandPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CancelCommandFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CancelCommandToPb(st *CancelCommand) (*computepb.CancelCommandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CancelCommandPb{}
	pb.ClusterId = st.ClusterId
	pb.CommandId = st.CommandId
	pb.ContextId = st.ContextId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CancelCommandFromPb(pb *computepb.CancelCommandPb) (*CancelCommand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelCommand{}
	st.ClusterId = pb.ClusterId
	st.CommandId = pb.CommandId
	st.ContextId = pb.ContextId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ChangeClusterOwner struct {

	// Wire name: 'cluster_id'
	ClusterId string ``
	// New owner of the cluster_id after this RPC.
	// Wire name: 'owner_username'
	OwnerUsername string ``
}

func (st ChangeClusterOwner) MarshalJSON() ([]byte, error) {
	pb, err := ChangeClusterOwnerToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ChangeClusterOwner) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ChangeClusterOwnerPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ChangeClusterOwnerFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ChangeClusterOwnerToPb(st *ChangeClusterOwner) (*computepb.ChangeClusterOwnerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ChangeClusterOwnerPb{}
	pb.ClusterId = st.ClusterId
	pb.OwnerUsername = st.OwnerUsername

	return pb, nil
}

func ChangeClusterOwnerFromPb(pb *computepb.ChangeClusterOwnerPb) (*ChangeClusterOwner, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChangeClusterOwner{}
	st.ClusterId = pb.ClusterId
	st.OwnerUsername = pb.OwnerUsername

	return st, nil
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	// Wire name: 'jobs'
	Jobs bool ``
	// With notebooks set, this cluster can be used for notebooks
	// Wire name: 'notebooks'
	Notebooks       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st ClientsTypes) MarshalJSON() ([]byte, error) {
	pb, err := ClientsTypesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClientsTypes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClientsTypesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClientsTypesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClientsTypesToPb(st *ClientsTypes) (*computepb.ClientsTypesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClientsTypesPb{}
	pb.Jobs = st.Jobs
	pb.Notebooks = st.Notebooks

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClientsTypesFromPb(pb *computepb.ClientsTypesPb) (*ClientsTypes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClientsTypes{}
	st.Jobs = pb.Jobs
	st.Notebooks = pb.Notebooks

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CloneCluster struct {
	// The cluster that is being cloned.
	// Wire name: 'source_cluster_id'
	SourceClusterId string ``
}

func (st CloneCluster) MarshalJSON() ([]byte, error) {
	pb, err := CloneClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CloneCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CloneClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CloneClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CloneClusterToPb(st *CloneCluster) (*computepb.CloneClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CloneClusterPb{}
	pb.SourceClusterId = st.SourceClusterId

	return pb, nil
}

func CloneClusterFromPb(pb *computepb.CloneClusterPb) (*CloneCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloneCluster{}
	st.SourceClusterId = pb.SourceClusterId

	return st, nil
}

type CloudProviderNodeInfo struct {
	// Status as reported by the cloud provider
	// Wire name: 'status'
	Status []CloudProviderNodeStatus ``
}

func (st CloudProviderNodeInfo) MarshalJSON() ([]byte, error) {
	pb, err := CloudProviderNodeInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CloudProviderNodeInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CloudProviderNodeInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CloudProviderNodeInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CloudProviderNodeInfoToPb(st *CloudProviderNodeInfo) (*computepb.CloudProviderNodeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CloudProviderNodeInfoPb{}

	var statusPb []computepb.CloudProviderNodeStatusPb
	for _, item := range st.Status {
		itemPb, err := CloudProviderNodeStatusToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statusPb = append(statusPb, *itemPb)
		}
	}
	pb.Status = statusPb

	return pb, nil
}

func CloudProviderNodeInfoFromPb(pb *computepb.CloudProviderNodeInfoPb) (*CloudProviderNodeInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudProviderNodeInfo{}

	var statusField []CloudProviderNodeStatus
	for _, itemPb := range pb.Status {
		item, err := CloudProviderNodeStatusFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statusField = append(statusField, *item)
		}
	}
	st.Status = statusField

	return st, nil
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

func CloudProviderNodeStatusToPb(st *CloudProviderNodeStatus) (*computepb.CloudProviderNodeStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.CloudProviderNodeStatusPb(*st)
	return &pb, nil
}

func CloudProviderNodeStatusFromPb(pb *computepb.CloudProviderNodeStatusPb) (*CloudProviderNodeStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := CloudProviderNodeStatus(*pb)
	return &st, nil
}

type ClusterAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel ClusterPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAccessControlRequestToPb(st *ClusterAccessControlRequest) (*computepb.ClusterAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := ClusterPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterAccessControlRequestFromPb(pb *computepb.ClusterAccessControlRequestPb) (*ClusterAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := ClusterPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ClusterPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAccessControlResponseToPb(st *ClusterAccessControlResponse) (*computepb.ClusterAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterAccessControlResponsePb{}

	var allPermissionsPb []computepb.ClusterPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := ClusterPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterAccessControlResponseFromPb(pb *computepb.ClusterAccessControlResponsePb) (*ClusterAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAccessControlResponse{}

	var allPermissionsField []ClusterPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := ClusterPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	AutoterminationMinutes int ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf ``
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``

	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode ``
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool ``

	// Wire name: 'kind'
	Kind Kind ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine ``
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
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
	SparkEnvVars map[string]string ``
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool ``

	// Wire name: 'workload_type'
	WorkloadType    *WorkloadType ``
	ForceSendFields []string      `tf:"-"`
}

func (st ClusterAttributes) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAttributesToPb(st *ClusterAttributes) (*computepb.ClusterAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterAttributesPb{}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	awsAttributesPb, err := AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	dataSecurityModePb, err := DataSecurityModeToPb(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}
	dockerImagePb, err := DockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []computepb.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	kindPb, err := KindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.NodeTypeId = st.NodeTypeId
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	runtimeEnginePb, err := RuntimeEngineToPb(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	workloadTypePb, err := WorkloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterAttributesFromPb(pb *computepb.ClusterAttributesPb) (*ClusterAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAttributes{}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	awsAttributesField, err := AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	dataSecurityModeField, err := DataSecurityModeFromPb(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := DockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	gcpAttributesField, err := GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	kindField, err := KindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.NodeTypeId = pb.NodeTypeId
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	runtimeEngineField, err := RuntimeEngineFromPb(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	workloadTypeField, err := WorkloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterCompliance struct {
	// Canonical unique identifier for a cluster.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// Whether this cluster is in compliance with the latest version of its
	// policy.
	// Wire name: 'is_compliant'
	IsCompliant bool ``
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	// Wire name: 'violations'
	Violations      map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (st ClusterCompliance) MarshalJSON() ([]byte, error) {
	pb, err := ClusterComplianceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterCompliance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterCompliancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterComplianceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterComplianceToPb(st *ClusterCompliance) (*computepb.ClusterCompliancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterCompliancePb{}
	pb.ClusterId = st.ClusterId
	pb.IsCompliant = st.IsCompliant
	pb.Violations = st.Violations

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterComplianceFromPb(pb *computepb.ClusterCompliancePb) (*ClusterCompliance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterCompliance{}
	st.ClusterId = pb.ClusterId
	st.IsCompliant = pb.IsCompliant
	st.Violations = pb.Violations

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Describes all of the metadata about a single Spark cluster in Databricks.
type ClusterDetails struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes ``
	// Number of CPU cores available for this cluster. Note that this can be
	// fractional, e.g. 7.5 cores, since certain node types are configured to
	// share cores between Spark nodes on the same instance.
	// Wire name: 'cluster_cores'
	ClusterCores float64 ``
	// Canonical identifier for the cluster. This id is retained during cluster
	// restarts and resizes, while each new cluster has a globally unique id.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf ``
	// Cluster log delivery status.
	// Wire name: 'cluster_log_status'
	ClusterLogStatus *LogSyncStatus ``
	// Total amount of cluster memory, in megabytes
	// Wire name: 'cluster_memory_mb'
	ClusterMemoryMb int64 ``
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string ``
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request.
	// Wire name: 'cluster_source'
	ClusterSource ClusterSource ``
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``

	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode ``
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
	DefaultTags map[string]string ``
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage ``
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	// Wire name: 'driver'
	Driver *SparkNode ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Nodes on which the Spark executors reside.
	// Wire name: 'executors'
	Executors []SparkNode ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool ``
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	// Wire name: 'jdbc_port'
	JdbcPort int ``

	// Wire name: 'kind'
	Kind Kind ``
	// the timestamp that the cluster was started/restarted
	// Wire name: 'last_restarted_time'
	LastRestartedTime int64 ``
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	// Wire name: 'last_state_loss_time'
	LastStateLossTime int64 ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
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
	NumWorkers int ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine ``
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	// Wire name: 'spark_context_id'
	SparkContextId int64 ``
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
	SparkEnvVars map[string]string ``
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string ``
	// The spec contains a snapshot of the latest user specified settings that
	// were used to create/edit the cluster. Note: not included in the response
	// of the ListClusters API.
	// Wire name: 'spec'
	Spec *ClusterSpec ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string ``
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	// Wire name: 'start_time'
	StartTime int64 ``
	// Current state of the cluster.
	// Wire name: 'state'
	State State ``
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	// Wire name: 'state_message'
	StateMessage string ``
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	// Wire name: 'terminated_time'
	TerminatedTime int64 ``
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	// Wire name: 'termination_reason'
	TerminationReason *TerminationReason ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool ``

	// Wire name: 'workload_type'
	WorkloadType    *WorkloadType ``
	ForceSendFields []string      `tf:"-"`
}

func (st ClusterDetails) MarshalJSON() ([]byte, error) {
	pb, err := ClusterDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterDetailsToPb(st *ClusterDetails) (*computepb.ClusterDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterDetailsPb{}
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	awsAttributesPb, err := AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	pb.ClusterCores = st.ClusterCores
	pb.ClusterId = st.ClusterId
	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}
	clusterLogStatusPb, err := LogSyncStatusToPb(st.ClusterLogStatus)
	if err != nil {
		return nil, err
	}
	if clusterLogStatusPb != nil {
		pb.ClusterLogStatus = clusterLogStatusPb
	}
	pb.ClusterMemoryMb = st.ClusterMemoryMb
	pb.ClusterName = st.ClusterName
	clusterSourcePb, err := ClusterSourceToPb(&st.ClusterSource)
	if err != nil {
		return nil, err
	}
	if clusterSourcePb != nil {
		pb.ClusterSource = *clusterSourcePb
	}
	pb.CreatorUserName = st.CreatorUserName
	pb.CustomTags = st.CustomTags
	dataSecurityModePb, err := DataSecurityModeToPb(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}
	pb.DefaultTags = st.DefaultTags
	dockerImagePb, err := DockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}
	driverPb, err := SparkNodeToPb(st.Driver)
	if err != nil {
		return nil, err
	}
	if driverPb != nil {
		pb.Driver = driverPb
	}
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption

	var executorsPb []computepb.SparkNodePb
	for _, item := range st.Executors {
		itemPb, err := SparkNodeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			executorsPb = append(executorsPb, *itemPb)
		}
	}
	pb.Executors = executorsPb
	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []computepb.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	pb.JdbcPort = st.JdbcPort
	kindPb, err := KindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.LastRestartedTime = st.LastRestartedTime
	pb.LastStateLossTime = st.LastStateLossTime
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	runtimeEnginePb, err := RuntimeEngineToPb(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkContextId = st.SparkContextId
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	specPb, err := ClusterSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}
	pb.SshPublicKeys = st.SshPublicKeys
	pb.StartTime = st.StartTime
	statePb, err := StateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	pb.StateMessage = st.StateMessage
	pb.TerminatedTime = st.TerminatedTime
	terminationReasonPb, err := TerminationReasonToPb(st.TerminationReason)
	if err != nil {
		return nil, err
	}
	if terminationReasonPb != nil {
		pb.TerminationReason = terminationReasonPb
	}
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	workloadTypePb, err := WorkloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterDetailsFromPb(pb *computepb.ClusterDetailsPb) (*ClusterDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterDetails{}
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	awsAttributesField, err := AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	st.ClusterCores = pb.ClusterCores
	st.ClusterId = pb.ClusterId
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	clusterLogStatusField, err := LogSyncStatusFromPb(pb.ClusterLogStatus)
	if err != nil {
		return nil, err
	}
	if clusterLogStatusField != nil {
		st.ClusterLogStatus = clusterLogStatusField
	}
	st.ClusterMemoryMb = pb.ClusterMemoryMb
	st.ClusterName = pb.ClusterName
	clusterSourceField, err := ClusterSourceFromPb(&pb.ClusterSource)
	if err != nil {
		return nil, err
	}
	if clusterSourceField != nil {
		st.ClusterSource = *clusterSourceField
	}
	st.CreatorUserName = pb.CreatorUserName
	st.CustomTags = pb.CustomTags
	dataSecurityModeField, err := DataSecurityModeFromPb(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	st.DefaultTags = pb.DefaultTags
	dockerImageField, err := DockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverField, err := SparkNodeFromPb(pb.Driver)
	if err != nil {
		return nil, err
	}
	if driverField != nil {
		st.Driver = driverField
	}
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption

	var executorsField []SparkNode
	for _, itemPb := range pb.Executors {
		item, err := SparkNodeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			executorsField = append(executorsField, *item)
		}
	}
	st.Executors = executorsField
	gcpAttributesField, err := GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	st.JdbcPort = pb.JdbcPort
	kindField, err := KindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.LastRestartedTime = pb.LastRestartedTime
	st.LastStateLossTime = pb.LastStateLossTime
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	runtimeEngineField, err := RuntimeEngineFromPb(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkContextId = pb.SparkContextId
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	specField, err := ClusterSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}
	st.SshPublicKeys = pb.SshPublicKeys
	st.StartTime = pb.StartTime
	stateField, err := StateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	st.StateMessage = pb.StateMessage
	st.TerminatedTime = pb.TerminatedTime
	terminationReasonField, err := TerminationReasonFromPb(pb.TerminationReason)
	if err != nil {
		return nil, err
	}
	if terminationReasonField != nil {
		st.TerminationReason = terminationReasonField
	}
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	workloadTypeField, err := WorkloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterEvent struct {

	// Wire name: 'cluster_id'
	ClusterId string ``

	// Wire name: 'data_plane_event_details'
	DataPlaneEventDetails *DataPlaneEventDetails ``

	// Wire name: 'details'
	Details *EventDetails ``
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	// Wire name: 'timestamp'
	Timestamp int64 ``

	// Wire name: 'type'
	Type            EventType ``
	ForceSendFields []string  `tf:"-"`
}

func (st ClusterEvent) MarshalJSON() ([]byte, error) {
	pb, err := ClusterEventToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterEvent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterEventPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterEventFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterEventToPb(st *ClusterEvent) (*computepb.ClusterEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterEventPb{}
	pb.ClusterId = st.ClusterId
	dataPlaneEventDetailsPb, err := DataPlaneEventDetailsToPb(st.DataPlaneEventDetails)
	if err != nil {
		return nil, err
	}
	if dataPlaneEventDetailsPb != nil {
		pb.DataPlaneEventDetails = dataPlaneEventDetailsPb
	}
	detailsPb, err := EventDetailsToPb(st.Details)
	if err != nil {
		return nil, err
	}
	if detailsPb != nil {
		pb.Details = detailsPb
	}
	pb.Timestamp = st.Timestamp
	typePb, err := EventTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterEventFromPb(pb *computepb.ClusterEventPb) (*ClusterEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterEvent{}
	st.ClusterId = pb.ClusterId
	dataPlaneEventDetailsField, err := DataPlaneEventDetailsFromPb(pb.DataPlaneEventDetails)
	if err != nil {
		return nil, err
	}
	if dataPlaneEventDetailsField != nil {
		st.DataPlaneEventDetails = dataPlaneEventDetailsField
	}
	detailsField, err := EventDetailsFromPb(pb.Details)
	if err != nil {
		return nil, err
	}
	if detailsField != nil {
		st.Details = detailsField
	}
	st.Timestamp = pb.Timestamp
	typeField, err := EventTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// Status of all libraries on the cluster.
	// Wire name: 'library_statuses'
	LibraryStatuses []LibraryFullStatus ``
	ForceSendFields []string            `tf:"-"`
}

func (st ClusterLibraryStatuses) MarshalJSON() ([]byte, error) {
	pb, err := ClusterLibraryStatusesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterLibraryStatuses) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterLibraryStatusesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterLibraryStatusesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterLibraryStatusesToPb(st *ClusterLibraryStatuses) (*computepb.ClusterLibraryStatusesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterLibraryStatusesPb{}
	pb.ClusterId = st.ClusterId

	var libraryStatusesPb []computepb.LibraryFullStatusPb
	for _, item := range st.LibraryStatuses {
		itemPb, err := LibraryFullStatusToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			libraryStatusesPb = append(libraryStatusesPb, *itemPb)
		}
	}
	pb.LibraryStatuses = libraryStatusesPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterLibraryStatusesFromPb(pb *computepb.ClusterLibraryStatusesPb) (*ClusterLibraryStatuses, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterLibraryStatuses{}
	st.ClusterId = pb.ClusterId

	var libraryStatusesField []LibraryFullStatus
	for _, itemPb := range pb.LibraryStatuses {
		item, err := LibraryFullStatusFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			libraryStatusesField = append(libraryStatusesField, *item)
		}
	}
	st.LibraryStatuses = libraryStatusesField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Cluster log delivery config
type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	// Wire name: 'dbfs'
	Dbfs *DbfsStorageInfo ``
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	// Wire name: 's3'
	S3 *S3StorageInfo ``
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	// Wire name: 'volumes'
	Volumes *VolumesStorageInfo ``
}

func (st ClusterLogConf) MarshalJSON() ([]byte, error) {
	pb, err := ClusterLogConfToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterLogConf) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterLogConfPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterLogConfFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterLogConfToPb(st *ClusterLogConf) (*computepb.ClusterLogConfPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterLogConfPb{}
	dbfsPb, err := DbfsStorageInfoToPb(st.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsPb != nil {
		pb.Dbfs = dbfsPb
	}
	s3Pb, err := S3StorageInfoToPb(st.S3)
	if err != nil {
		return nil, err
	}
	if s3Pb != nil {
		pb.S3 = s3Pb
	}
	volumesPb, err := VolumesStorageInfoToPb(st.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesPb != nil {
		pb.Volumes = volumesPb
	}

	return pb, nil
}

func ClusterLogConfFromPb(pb *computepb.ClusterLogConfPb) (*ClusterLogConf, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterLogConf{}
	dbfsField, err := DbfsStorageInfoFromPb(pb.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsField != nil {
		st.Dbfs = dbfsField
	}
	s3Field, err := S3StorageInfoFromPb(pb.S3)
	if err != nil {
		return nil, err
	}
	if s3Field != nil {
		st.S3 = s3Field
	}
	volumesField, err := VolumesStorageInfoFromPb(pb.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesField != nil {
		st.Volumes = volumesField
	}

	return st, nil
}

type ClusterPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel ClusterPermissionLevel ``
	ForceSendFields []string               `tf:"-"`
}

func (st ClusterPermission) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPermissionToPb(st *ClusterPermission) (*computepb.ClusterPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := ClusterPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPermissionFromPb(pb *computepb.ClusterPermissionPb) (*ClusterPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := ClusterPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ClusterPermissionLevelToPb(st *ClusterPermissionLevel) (*computepb.ClusterPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ClusterPermissionLevelPb(*st)
	return &pb, nil
}

func ClusterPermissionLevelFromPb(pb *computepb.ClusterPermissionLevelPb) (*ClusterPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterPermissionLevel(*pb)
	return &st, nil
}

type ClusterPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ClusterAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterPermissions) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPermissionsToPb(st *ClusterPermissions) (*computepb.ClusterPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPermissionsPb{}

	var accessControlListPb []computepb.ClusterAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := ClusterAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPermissionsFromPb(pb *computepb.ClusterPermissionsPb) (*ClusterPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissions{}

	var accessControlListField []ClusterAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := ClusterAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel ClusterPermissionLevel ``
	ForceSendFields []string               `tf:"-"`
}

func (st ClusterPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPermissionsDescriptionToPb(st *ClusterPermissionsDescription) (*computepb.ClusterPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := ClusterPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPermissionsDescriptionFromPb(pb *computepb.ClusterPermissionsDescriptionPb) (*ClusterPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := ClusterPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ClusterAccessControlRequest ``
	// The cluster for which to get or manage permissions.
	// Wire name: 'cluster_id'
	ClusterId string `tf:"-"`
}

func (st ClusterPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPermissionsRequestToPb(st *ClusterPermissionsRequest) (*computepb.ClusterPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPermissionsRequestPb{}

	var accessControlListPb []computepb.ClusterAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := ClusterAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func ClusterPermissionsRequestFromPb(pb *computepb.ClusterPermissionsRequestPb) (*ClusterPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissionsRequest{}

	var accessControlListField []ClusterAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := ClusterAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ClusterId = pb.ClusterId

	return st, nil
}

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel ClusterPolicyPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterPolicyAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPolicyAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPolicyAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPolicyAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPolicyAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPolicyAccessControlRequestToPb(st *ClusterPolicyAccessControlRequest) (*computepb.ClusterPolicyAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPolicyAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := ClusterPolicyPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPolicyAccessControlRequestFromPb(pb *computepb.ClusterPolicyAccessControlRequestPb) (*ClusterPolicyAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := ClusterPolicyPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterPolicyAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ClusterPolicyPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterPolicyAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPolicyAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPolicyAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPolicyAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPolicyAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPolicyAccessControlResponseToPb(st *ClusterPolicyAccessControlResponse) (*computepb.ClusterPolicyAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPolicyAccessControlResponsePb{}

	var allPermissionsPb []computepb.ClusterPolicyPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := ClusterPolicyPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPolicyAccessControlResponseFromPb(pb *computepb.ClusterPolicyAccessControlResponsePb) (*ClusterPolicyAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyAccessControlResponse{}

	var allPermissionsField []ClusterPolicyPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := ClusterPolicyPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterPolicyPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel ClusterPolicyPermissionLevel ``
	ForceSendFields []string                     `tf:"-"`
}

func (st ClusterPolicyPermission) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPolicyPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPolicyPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPolicyPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPolicyPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPolicyPermissionToPb(st *ClusterPolicyPermission) (*computepb.ClusterPolicyPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPolicyPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := ClusterPolicyPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPolicyPermissionFromPb(pb *computepb.ClusterPolicyPermissionPb) (*ClusterPolicyPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := ClusterPolicyPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ClusterPolicyPermissionLevelToPb(st *ClusterPolicyPermissionLevel) (*computepb.ClusterPolicyPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ClusterPolicyPermissionLevelPb(*st)
	return &pb, nil
}

func ClusterPolicyPermissionLevelFromPb(pb *computepb.ClusterPolicyPermissionLevelPb) (*ClusterPolicyPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterPolicyPermissionLevel(*pb)
	return &st, nil
}

type ClusterPolicyPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ClusterPolicyAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterPolicyPermissions) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPolicyPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPolicyPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPolicyPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPolicyPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPolicyPermissionsToPb(st *ClusterPolicyPermissions) (*computepb.ClusterPolicyPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPolicyPermissionsPb{}

	var accessControlListPb []computepb.ClusterPolicyAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := ClusterPolicyAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPolicyPermissionsFromPb(pb *computepb.ClusterPolicyPermissionsPb) (*ClusterPolicyPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissions{}

	var accessControlListField []ClusterPolicyAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := ClusterPolicyAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterPolicyPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel ClusterPolicyPermissionLevel ``
	ForceSendFields []string                     `tf:"-"`
}

func (st ClusterPolicyPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPolicyPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPolicyPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPolicyPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPolicyPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPolicyPermissionsDescriptionToPb(st *ClusterPolicyPermissionsDescription) (*computepb.ClusterPolicyPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPolicyPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := ClusterPolicyPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterPolicyPermissionsDescriptionFromPb(pb *computepb.ClusterPolicyPermissionsDescriptionPb) (*ClusterPolicyPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := ClusterPolicyPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterPolicyPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ClusterPolicyAccessControlRequest ``
	// The cluster policy for which to get or manage permissions.
	// Wire name: 'cluster_policy_id'
	ClusterPolicyId string `tf:"-"`
}

func (st ClusterPolicyPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ClusterPolicyPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterPolicyPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterPolicyPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterPolicyPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterPolicyPermissionsRequestToPb(st *ClusterPolicyPermissionsRequest) (*computepb.ClusterPolicyPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterPolicyPermissionsRequestPb{}

	var accessControlListPb []computepb.ClusterPolicyAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := ClusterPolicyAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ClusterPolicyId = st.ClusterPolicyId

	return pb, nil
}

func ClusterPolicyPermissionsRequestFromPb(pb *computepb.ClusterPolicyPermissionsRequestPb) (*ClusterPolicyPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissionsRequest{}

	var accessControlListField []ClusterPolicyAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := ClusterPolicyAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ClusterPolicyId = pb.ClusterPolicyId

	return st, nil
}

// Represents a change to the cluster settings required for the cluster to
// become compliant with its policy.
type ClusterSettingsChange struct {
	// The field where this change would be made.
	// Wire name: 'field'
	Field string ``
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	// Wire name: 'new_value'
	NewValue string ``
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	// Wire name: 'previous_value'
	PreviousValue   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterSettingsChange) MarshalJSON() ([]byte, error) {
	pb, err := ClusterSettingsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterSettingsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterSettingsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterSettingsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterSettingsChangeToPb(st *ClusterSettingsChange) (*computepb.ClusterSettingsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterSettingsChangePb{}
	pb.Field = st.Field
	pb.NewValue = st.NewValue
	pb.PreviousValue = st.PreviousValue

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterSettingsChangeFromPb(pb *computepb.ClusterSettingsChangePb) (*ClusterSettingsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSettingsChange{}
	st.Field = pb.Field
	st.NewValue = pb.NewValue
	st.PreviousValue = pb.PreviousValue

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterSize struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
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
	NumWorkers      int      ``
	ForceSendFields []string `tf:"-"`
}

func (st ClusterSize) MarshalJSON() ([]byte, error) {
	pb, err := ClusterSizeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterSize) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterSizePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterSizeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterSizeToPb(st *ClusterSize) (*computepb.ClusterSizePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterSizePb{}
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.NumWorkers = st.NumWorkers

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterSizeFromPb(pb *computepb.ClusterSizePb) (*ClusterSize, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSize{}
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.NumWorkers = pb.NumWorkers

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ClusterSourceToPb(st *ClusterSource) (*computepb.ClusterSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ClusterSourcePb(*st)
	return &pb, nil
}

func ClusterSourceFromPb(pb *computepb.ClusterSourcePb) (*ClusterSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterSource(*pb)
	return &st, nil
}

// Contains a snapshot of the latest user specified settings that were used to
// create/edit the cluster.
type ClusterSpec struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool ``
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf ``
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``

	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode ``
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool ``

	// Wire name: 'kind'
	Kind Kind ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
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
	NumWorkers int ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine ``
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
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
	SparkEnvVars map[string]string ``
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool ``

	// Wire name: 'workload_type'
	WorkloadType    *WorkloadType ``
	ForceSendFields []string      `tf:"-"`
}

func (st ClusterSpec) MarshalJSON() ([]byte, error) {
	pb, err := ClusterSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterSpecToPb(st *ClusterSpec) (*computepb.ClusterSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterSpecPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	awsAttributesPb, err := AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	dataSecurityModePb, err := DataSecurityModeToPb(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}
	dockerImagePb, err := DockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []computepb.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	kindPb, err := KindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	runtimeEnginePb, err := RuntimeEngineToPb(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	workloadTypePb, err := WorkloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterSpecFromPb(pb *computepb.ClusterSpecPb) (*ClusterSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSpec{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	awsAttributesField, err := AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	dataSecurityModeField, err := DataSecurityModeFromPb(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := DockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	gcpAttributesField, err := GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	kindField, err := KindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	runtimeEngineField, err := RuntimeEngineFromPb(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	workloadTypeField, err := WorkloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	// Wire name: 'cluster_id'
	ClusterId string `tf:"-"`
}

func (st ClusterStatus) MarshalJSON() ([]byte, error) {
	pb, err := ClusterStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ClusterStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterStatusToPb(st *ClusterStatus) (*computepb.ClusterStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ClusterStatusPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func ClusterStatusFromPb(pb *computepb.ClusterStatusPb) (*ClusterStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterStatus{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type Command struct {
	// Running cluster id
	// Wire name: 'clusterId'
	ClusterId string ``
	// Executable code
	// Wire name: 'command'
	Command string ``
	// Running context id
	// Wire name: 'contextId'
	ContextId string ``

	// Wire name: 'language'
	Language        Language ``
	ForceSendFields []string `tf:"-"`
}

func (st Command) MarshalJSON() ([]byte, error) {
	pb, err := CommandToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Command) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CommandPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CommandFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CommandToPb(st *Command) (*computepb.CommandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CommandPb{}
	pb.ClusterId = st.ClusterId
	pb.Command = st.Command
	pb.ContextId = st.ContextId
	languagePb, err := LanguageToPb(&st.Language)
	if err != nil {
		return nil, err
	}
	if languagePb != nil {
		pb.Language = *languagePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CommandFromPb(pb *computepb.CommandPb) (*Command, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Command{}
	st.ClusterId = pb.ClusterId
	st.Command = pb.Command
	st.ContextId = pb.ContextId
	languageField, err := LanguageFromPb(&pb.Language)
	if err != nil {
		return nil, err
	}
	if languageField != nil {
		st.Language = *languageField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func CommandStatusToPb(st *CommandStatus) (*computepb.CommandStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.CommandStatusPb(*st)
	return &pb, nil
}

func CommandStatusFromPb(pb *computepb.CommandStatusPb) (*CommandStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := CommandStatus(*pb)
	return &st, nil
}

type CommandStatusRequest struct {

	// Wire name: 'clusterId'
	ClusterId string `tf:"-"`

	// Wire name: 'commandId'
	CommandId string `tf:"-"`

	// Wire name: 'contextId'
	ContextId string `tf:"-"`
}

func (st CommandStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := CommandStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CommandStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CommandStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CommandStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CommandStatusRequestToPb(st *CommandStatusRequest) (*computepb.CommandStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CommandStatusRequestPb{}
	pb.ClusterId = st.ClusterId
	pb.CommandId = st.CommandId
	pb.ContextId = st.ContextId

	return pb, nil
}

func CommandStatusRequestFromPb(pb *computepb.CommandStatusRequestPb) (*CommandStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommandStatusRequest{}
	st.ClusterId = pb.ClusterId
	st.CommandId = pb.CommandId
	st.ContextId = pb.ContextId

	return st, nil
}

type CommandStatusResponse struct {

	// Wire name: 'id'
	Id string ``

	// Wire name: 'results'
	Results *Results ``

	// Wire name: 'status'
	Status          CommandStatus ``
	ForceSendFields []string      `tf:"-"`
}

func (st CommandStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := CommandStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CommandStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CommandStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CommandStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CommandStatusResponseToPb(st *CommandStatusResponse) (*computepb.CommandStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CommandStatusResponsePb{}
	pb.Id = st.Id
	resultsPb, err := ResultsToPb(st.Results)
	if err != nil {
		return nil, err
	}
	if resultsPb != nil {
		pb.Results = resultsPb
	}
	statusPb, err := CommandStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CommandStatusResponseFromPb(pb *computepb.CommandStatusResponsePb) (*CommandStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommandStatusResponse{}
	st.Id = pb.Id
	resultsField, err := ResultsFromPb(pb.Results)
	if err != nil {
		return nil, err
	}
	if resultsField != nil {
		st.Results = resultsField
	}
	statusField, err := CommandStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ContextStatusToPb(st *ContextStatus) (*computepb.ContextStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ContextStatusPb(*st)
	return &pb, nil
}

func ContextStatusFromPb(pb *computepb.ContextStatusPb) (*ContextStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ContextStatus(*pb)
	return &st, nil
}

type ContextStatusRequest struct {

	// Wire name: 'clusterId'
	ClusterId string `tf:"-"`

	// Wire name: 'contextId'
	ContextId string `tf:"-"`
}

func (st ContextStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := ContextStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ContextStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ContextStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ContextStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ContextStatusRequestToPb(st *ContextStatusRequest) (*computepb.ContextStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ContextStatusRequestPb{}
	pb.ClusterId = st.ClusterId
	pb.ContextId = st.ContextId

	return pb, nil
}

func ContextStatusRequestFromPb(pb *computepb.ContextStatusRequestPb) (*ContextStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContextStatusRequest{}
	st.ClusterId = pb.ClusterId
	st.ContextId = pb.ContextId

	return st, nil
}

type ContextStatusResponse struct {

	// Wire name: 'id'
	Id string ``

	// Wire name: 'status'
	Status          ContextStatus ``
	ForceSendFields []string      `tf:"-"`
}

func (st ContextStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := ContextStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ContextStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ContextStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ContextStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ContextStatusResponseToPb(st *ContextStatusResponse) (*computepb.ContextStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ContextStatusResponsePb{}
	pb.Id = st.Id
	statusPb, err := ContextStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ContextStatusResponseFromPb(pb *computepb.ContextStatusResponsePb) (*ContextStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContextStatusResponse{}
	st.Id = pb.Id
	statusField, err := ContextStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool ``
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes ``
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	// Wire name: 'clone_from'
	CloneFrom *CloneCluster ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf ``
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``

	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode ``
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool ``

	// Wire name: 'kind'
	Kind Kind ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
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
	NumWorkers int ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine ``
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
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
	SparkEnvVars map[string]string ``
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool ``

	// Wire name: 'workload_type'
	WorkloadType    *WorkloadType ``
	ForceSendFields []string      `tf:"-"`
}

func (st CreateCluster) MarshalJSON() ([]byte, error) {
	pb, err := CreateClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreateClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateClusterToPb(st *CreateCluster) (*computepb.CreateClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreateClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	awsAttributesPb, err := AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	cloneFromPb, err := CloneClusterToPb(st.CloneFrom)
	if err != nil {
		return nil, err
	}
	if cloneFromPb != nil {
		pb.CloneFrom = cloneFromPb
	}
	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	dataSecurityModePb, err := DataSecurityModeToPb(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}
	dockerImagePb, err := DockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []computepb.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	kindPb, err := KindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	runtimeEnginePb, err := RuntimeEngineToPb(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	workloadTypePb, err := WorkloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateClusterFromPb(pb *computepb.CreateClusterPb) (*CreateCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	awsAttributesField, err := AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	cloneFromField, err := CloneClusterFromPb(pb.CloneFrom)
	if err != nil {
		return nil, err
	}
	if cloneFromField != nil {
		st.CloneFrom = cloneFromField
	}
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	dataSecurityModeField, err := DataSecurityModeFromPb(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := DockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	gcpAttributesField, err := GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	kindField, err := KindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	runtimeEngineField, err := RuntimeEngineFromPb(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	workloadTypeField, err := WorkloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateClusterResponse struct {

	// Wire name: 'cluster_id'
	ClusterId       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateClusterResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateClusterResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateClusterResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreateClusterResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateClusterResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateClusterResponseToPb(st *CreateClusterResponse) (*computepb.CreateClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreateClusterResponsePb{}
	pb.ClusterId = st.ClusterId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateClusterResponseFromPb(pb *computepb.CreateClusterResponsePb) (*CreateClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateClusterResponse{}
	st.ClusterId = pb.ClusterId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateContext struct {
	// Running cluster id
	// Wire name: 'clusterId'
	ClusterId string ``

	// Wire name: 'language'
	Language        Language ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateContext) MarshalJSON() ([]byte, error) {
	pb, err := CreateContextToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateContext) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreateContextPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateContextFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateContextToPb(st *CreateContext) (*computepb.CreateContextPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreateContextPb{}
	pb.ClusterId = st.ClusterId
	languagePb, err := LanguageToPb(&st.Language)
	if err != nil {
		return nil, err
	}
	if languagePb != nil {
		pb.Language = *languagePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateContextFromPb(pb *computepb.CreateContextPb) (*CreateContext, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateContext{}
	st.ClusterId = pb.ClusterId
	languageField, err := LanguageFromPb(&pb.Language)
	if err != nil {
		return nil, err
	}
	if languageField != nil {
		st.Language = *languageField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *InstancePoolAwsAttributes ``
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *InstancePoolAzureAttributes ``
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	// Wire name: 'disk_spec'
	DiskSpec *DiskSpec ``
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *InstancePoolGcpAttributes ``
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int ``
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string ``
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int ``
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
	// Custom Docker Image BYOC
	// Wire name: 'preloaded_docker_images'
	PreloadedDockerImages []DockerImage ``
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'preloaded_spark_versions'
	PreloadedSparkVersions []string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int      ``
	ForceSendFields            []string `tf:"-"`
}

func (st CreateInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := CreateInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreateInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateInstancePoolToPb(st *CreateInstancePool) (*computepb.CreateInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreateInstancePoolPb{}
	awsAttributesPb, err := InstancePoolAwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := InstancePoolAzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	pb.CustomTags = st.CustomTags
	diskSpecPb, err := DiskSpecToPb(st.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecPb != nil {
		pb.DiskSpec = diskSpecPb
	}
	pb.EnableElasticDisk = st.EnableElasticDisk
	gcpAttributesPb, err := InstancePoolGcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId

	var preloadedDockerImagesPb []computepb.DockerImagePb
	for _, item := range st.PreloadedDockerImages {
		itemPb, err := DockerImageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedDockerImagesPb = append(preloadedDockerImagesPb, *itemPb)
		}
	}
	pb.PreloadedDockerImages = preloadedDockerImagesPb
	pb.PreloadedSparkVersions = st.PreloadedSparkVersions
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateInstancePoolFromPb(pb *computepb.CreateInstancePoolPb) (*CreateInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstancePool{}
	awsAttributesField, err := InstancePoolAwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := InstancePoolAzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	st.CustomTags = pb.CustomTags
	diskSpecField, err := DiskSpecFromPb(pb.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecField != nil {
		st.DiskSpec = diskSpecField
	}
	st.EnableElasticDisk = pb.EnableElasticDisk
	gcpAttributesField, err := InstancePoolGcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId

	var preloadedDockerImagesField []DockerImage
	for _, itemPb := range pb.PreloadedDockerImages {
		item, err := DockerImageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedDockerImagesField = append(preloadedDockerImagesField, *item)
		}
	}
	st.PreloadedDockerImages = preloadedDockerImagesField
	st.PreloadedSparkVersions = pb.PreloadedSparkVersions
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId  string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateInstancePoolResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateInstancePoolResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateInstancePoolResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreateInstancePoolResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateInstancePoolResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateInstancePoolResponseToPb(st *CreateInstancePoolResponse) (*computepb.CreateInstancePoolResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreateInstancePoolResponsePb{}
	pb.InstancePoolId = st.InstancePoolId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateInstancePoolResponseFromPb(pb *computepb.CreateInstancePoolResponsePb) (*CreateInstancePoolResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstancePoolResponse{}
	st.InstancePoolId = pb.InstancePoolId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreatePolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string ``
	// Additional human-readable description of the cluster policy.
	// Wire name: 'description'
	Description string ``
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	// Wire name: 'libraries'
	Libraries []Library ``
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	// Wire name: 'max_clusters_per_user'
	MaxClustersPerUser int64 ``
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	// Wire name: 'name'
	Name string ``
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
	PolicyFamilyDefinitionOverrides string ``
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	// Wire name: 'policy_family_id'
	PolicyFamilyId  string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreatePolicy) MarshalJSON() ([]byte, error) {
	pb, err := CreatePolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreatePolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreatePolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatePolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatePolicyToPb(st *CreatePolicy) (*computepb.CreatePolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreatePolicyPb{}
	pb.Definition = st.Definition
	pb.Description = st.Description

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.MaxClustersPerUser = st.MaxClustersPerUser
	pb.Name = st.Name
	pb.PolicyFamilyDefinitionOverrides = st.PolicyFamilyDefinitionOverrides
	pb.PolicyFamilyId = st.PolicyFamilyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreatePolicyFromPb(pb *computepb.CreatePolicyPb) (*CreatePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePolicy{}
	st.Definition = pb.Definition
	st.Description = pb.Description

	var librariesField []Library
	for _, itemPb := range pb.Libraries {
		item, err := LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.MaxClustersPerUser = pb.MaxClustersPerUser
	st.Name = pb.Name
	st.PolicyFamilyDefinitionOverrides = pb.PolicyFamilyDefinitionOverrides
	st.PolicyFamilyId = pb.PolicyFamilyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	// Wire name: 'policy_id'
	PolicyId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreatePolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreatePolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreatePolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreatePolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatePolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatePolicyResponseToPb(st *CreatePolicyResponse) (*computepb.CreatePolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreatePolicyResponsePb{}
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreatePolicyResponseFromPb(pb *computepb.CreatePolicyResponsePb) (*CreatePolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePolicyResponse{}
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateResponse struct {
	// The global init script ID.
	// Wire name: 'script_id'
	ScriptId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateResponseToPb(st *CreateResponse) (*computepb.CreateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreateResponsePb{}
	pb.ScriptId = st.ScriptId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateResponseFromPb(pb *computepb.CreateResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	st.ScriptId = pb.ScriptId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Created struct {

	// Wire name: 'id'
	Id              string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Created) MarshalJSON() ([]byte, error) {
	pb, err := CreatedToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Created) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CreatedPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatedFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatedToPb(st *Created) (*computepb.CreatedPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CreatedPb{}
	pb.Id = st.Id

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreatedFromPb(pb *computepb.CreatedPb) (*Created, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Created{}
	st.Id = pb.Id

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CustomPolicyTag struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	// Wire name: 'key'
	Key string ``
	// The value of the tag.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CustomPolicyTag) MarshalJSON() ([]byte, error) {
	pb, err := CustomPolicyTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CustomPolicyTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.CustomPolicyTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CustomPolicyTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CustomPolicyTagToPb(st *CustomPolicyTag) (*computepb.CustomPolicyTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.CustomPolicyTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CustomPolicyTagFromPb(pb *computepb.CustomPolicyTagPb) (*CustomPolicyTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomPolicyTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DataPlaneEventDetails struct {

	// Wire name: 'event_type'
	EventType DataPlaneEventDetailsEventType ``

	// Wire name: 'executor_failures'
	ExecutorFailures int ``

	// Wire name: 'host_id'
	HostId string ``

	// Wire name: 'timestamp'
	Timestamp       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st DataPlaneEventDetails) MarshalJSON() ([]byte, error) {
	pb, err := DataPlaneEventDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DataPlaneEventDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DataPlaneEventDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DataPlaneEventDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DataPlaneEventDetailsToPb(st *DataPlaneEventDetails) (*computepb.DataPlaneEventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DataPlaneEventDetailsPb{}
	eventTypePb, err := DataPlaneEventDetailsEventTypeToPb(&st.EventType)
	if err != nil {
		return nil, err
	}
	if eventTypePb != nil {
		pb.EventType = *eventTypePb
	}
	pb.ExecutorFailures = st.ExecutorFailures
	pb.HostId = st.HostId
	pb.Timestamp = st.Timestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DataPlaneEventDetailsFromPb(pb *computepb.DataPlaneEventDetailsPb) (*DataPlaneEventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneEventDetails{}
	eventTypeField, err := DataPlaneEventDetailsEventTypeFromPb(&pb.EventType)
	if err != nil {
		return nil, err
	}
	if eventTypeField != nil {
		st.EventType = *eventTypeField
	}
	st.ExecutorFailures = pb.ExecutorFailures
	st.HostId = pb.HostId
	st.Timestamp = pb.Timestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func DataPlaneEventDetailsEventTypeToPb(st *DataPlaneEventDetailsEventType) (*computepb.DataPlaneEventDetailsEventTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.DataPlaneEventDetailsEventTypePb(*st)
	return &pb, nil
}

func DataPlaneEventDetailsEventTypeFromPb(pb *computepb.DataPlaneEventDetailsEventTypePb) (*DataPlaneEventDetailsEventType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataPlaneEventDetailsEventType(*pb)
	return &st, nil
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

func DataSecurityModeToPb(st *DataSecurityMode) (*computepb.DataSecurityModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.DataSecurityModePb(*st)
	return &pb, nil
}

func DataSecurityModeFromPb(pb *computepb.DataSecurityModePb) (*DataSecurityMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataSecurityMode(*pb)
	return &st, nil
}

// A storage location in DBFS
type DbfsStorageInfo struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	// Wire name: 'destination'
	Destination string ``
}

func (st DbfsStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := DbfsStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DbfsStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DbfsStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DbfsStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DbfsStorageInfoToPb(st *DbfsStorageInfo) (*computepb.DbfsStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DbfsStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

func DbfsStorageInfoFromPb(pb *computepb.DbfsStorageInfoPb) (*DbfsStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbfsStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

type DeleteCluster struct {
	// The cluster to be terminated.
	// Wire name: 'cluster_id'
	ClusterId string ``
}

func (st DeleteCluster) MarshalJSON() ([]byte, error) {
	pb, err := DeleteClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DeleteClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteClusterToPb(st *DeleteCluster) (*computepb.DeleteClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DeleteClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func DeleteClusterFromPb(pb *computepb.DeleteClusterPb) (*DeleteCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	// Wire name: 'script_id'
	ScriptId string `tf:"-"`
}

func (st DeleteGlobalInitScriptRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteGlobalInitScriptRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteGlobalInitScriptRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DeleteGlobalInitScriptRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteGlobalInitScriptRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteGlobalInitScriptRequestToPb(st *DeleteGlobalInitScriptRequest) (*computepb.DeleteGlobalInitScriptRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DeleteGlobalInitScriptRequestPb{}
	pb.ScriptId = st.ScriptId

	return pb, nil
}

func DeleteGlobalInitScriptRequestFromPb(pb *computepb.DeleteGlobalInitScriptRequestPb) (*DeleteGlobalInitScriptRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteGlobalInitScriptRequest{}
	st.ScriptId = pb.ScriptId

	return st, nil
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
}

func (st DeleteInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := DeleteInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DeleteInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteInstancePoolToPb(st *DeleteInstancePool) (*computepb.DeleteInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DeleteInstancePoolPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

func DeleteInstancePoolFromPb(pb *computepb.DeleteInstancePoolPb) (*DeleteInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstancePool{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	// Wire name: 'policy_id'
	PolicyId string ``
}

func (st DeletePolicy) MarshalJSON() ([]byte, error) {
	pb, err := DeletePolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeletePolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DeletePolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeletePolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeletePolicyToPb(st *DeletePolicy) (*computepb.DeletePolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DeletePolicyPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func DeletePolicyFromPb(pb *computepb.DeletePolicyPb) (*DeletePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePolicy{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

type DestroyContext struct {

	// Wire name: 'clusterId'
	ClusterId string ``

	// Wire name: 'contextId'
	ContextId string ``
}

func (st DestroyContext) MarshalJSON() ([]byte, error) {
	pb, err := DestroyContextToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DestroyContext) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DestroyContextPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DestroyContextFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DestroyContextToPb(st *DestroyContext) (*computepb.DestroyContextPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DestroyContextPb{}
	pb.ClusterId = st.ClusterId
	pb.ContextId = st.ContextId

	return pb, nil
}

func DestroyContextFromPb(pb *computepb.DestroyContextPb) (*DestroyContext, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DestroyContext{}
	st.ClusterId = pb.ClusterId
	st.ContextId = pb.ContextId

	return st, nil
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
	DiskCount int ``

	// Wire name: 'disk_iops'
	DiskIops int ``
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	// Wire name: 'disk_size'
	DiskSize int ``

	// Wire name: 'disk_throughput'
	DiskThroughput int ``
	// The type of disks that will be launched with this cluster.
	// Wire name: 'disk_type'
	DiskType        *DiskType ``
	ForceSendFields []string  `tf:"-"`
}

func (st DiskSpec) MarshalJSON() ([]byte, error) {
	pb, err := DiskSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DiskSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DiskSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DiskSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DiskSpecToPb(st *DiskSpec) (*computepb.DiskSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DiskSpecPb{}
	pb.DiskCount = st.DiskCount
	pb.DiskIops = st.DiskIops
	pb.DiskSize = st.DiskSize
	pb.DiskThroughput = st.DiskThroughput
	diskTypePb, err := DiskTypeToPb(st.DiskType)
	if err != nil {
		return nil, err
	}
	if diskTypePb != nil {
		pb.DiskType = diskTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DiskSpecFromPb(pb *computepb.DiskSpecPb) (*DiskSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DiskSpec{}
	st.DiskCount = pb.DiskCount
	st.DiskIops = pb.DiskIops
	st.DiskSize = pb.DiskSize
	st.DiskThroughput = pb.DiskThroughput
	diskTypeField, err := DiskTypeFromPb(pb.DiskType)
	if err != nil {
		return nil, err
	}
	if diskTypeField != nil {
		st.DiskType = diskTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Describes the disk type.
type DiskType struct {

	// Wire name: 'azure_disk_volume_type'
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType ``

	// Wire name: 'ebs_volume_type'
	EbsVolumeType DiskTypeEbsVolumeType ``
}

func (st DiskType) MarshalJSON() ([]byte, error) {
	pb, err := DiskTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DiskType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DiskTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DiskTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DiskTypeToPb(st *DiskType) (*computepb.DiskTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DiskTypePb{}
	azureDiskVolumeTypePb, err := DiskTypeAzureDiskVolumeTypeToPb(&st.AzureDiskVolumeType)
	if err != nil {
		return nil, err
	}
	if azureDiskVolumeTypePb != nil {
		pb.AzureDiskVolumeType = *azureDiskVolumeTypePb
	}
	ebsVolumeTypePb, err := DiskTypeEbsVolumeTypeToPb(&st.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypePb != nil {
		pb.EbsVolumeType = *ebsVolumeTypePb
	}

	return pb, nil
}

func DiskTypeFromPb(pb *computepb.DiskTypePb) (*DiskType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DiskType{}
	azureDiskVolumeTypeField, err := DiskTypeAzureDiskVolumeTypeFromPb(&pb.AzureDiskVolumeType)
	if err != nil {
		return nil, err
	}
	if azureDiskVolumeTypeField != nil {
		st.AzureDiskVolumeType = *azureDiskVolumeTypeField
	}
	ebsVolumeTypeField, err := DiskTypeEbsVolumeTypeFromPb(&pb.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypeField != nil {
		st.EbsVolumeType = *ebsVolumeTypeField
	}

	return st, nil
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

func DiskTypeAzureDiskVolumeTypeToPb(st *DiskTypeAzureDiskVolumeType) (*computepb.DiskTypeAzureDiskVolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.DiskTypeAzureDiskVolumeTypePb(*st)
	return &pb, nil
}

func DiskTypeAzureDiskVolumeTypeFromPb(pb *computepb.DiskTypeAzureDiskVolumeTypePb) (*DiskTypeAzureDiskVolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DiskTypeAzureDiskVolumeType(*pb)
	return &st, nil
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

func DiskTypeEbsVolumeTypeToPb(st *DiskTypeEbsVolumeType) (*computepb.DiskTypeEbsVolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.DiskTypeEbsVolumeTypePb(*st)
	return &pb, nil
}

func DiskTypeEbsVolumeTypeFromPb(pb *computepb.DiskTypeEbsVolumeTypePb) (*DiskTypeEbsVolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DiskTypeEbsVolumeType(*pb)
	return &st, nil
}

type DockerBasicAuth struct {
	// Password of the user
	// Wire name: 'password'
	Password string ``
	// Name of the user
	// Wire name: 'username'
	Username        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st DockerBasicAuth) MarshalJSON() ([]byte, error) {
	pb, err := DockerBasicAuthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DockerBasicAuth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DockerBasicAuthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DockerBasicAuthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DockerBasicAuthToPb(st *DockerBasicAuth) (*computepb.DockerBasicAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DockerBasicAuthPb{}
	pb.Password = st.Password
	pb.Username = st.Username

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DockerBasicAuthFromPb(pb *computepb.DockerBasicAuthPb) (*DockerBasicAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DockerBasicAuth{}
	st.Password = pb.Password
	st.Username = pb.Username

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DockerImage struct {
	// Basic auth with username and password
	// Wire name: 'basic_auth'
	BasicAuth *DockerBasicAuth ``
	// URL of the docker image.
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (st DockerImage) MarshalJSON() ([]byte, error) {
	pb, err := DockerImageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DockerImage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.DockerImagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DockerImageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DockerImageToPb(st *DockerImage) (*computepb.DockerImagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.DockerImagePb{}
	basicAuthPb, err := DockerBasicAuthToPb(st.BasicAuth)
	if err != nil {
		return nil, err
	}
	if basicAuthPb != nil {
		pb.BasicAuth = basicAuthPb
	}
	pb.Url = st.Url

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DockerImageFromPb(pb *computepb.DockerImagePb) (*DockerImage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DockerImage{}
	basicAuthField, err := DockerBasicAuthFromPb(pb.BasicAuth)
	if err != nil {
		return nil, err
	}
	if basicAuthField != nil {
		st.BasicAuth = basicAuthField
	}
	st.Url = pb.Url

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func EbsVolumeTypeToPb(st *EbsVolumeType) (*computepb.EbsVolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.EbsVolumeTypePb(*st)
	return &pb, nil
}

func EbsVolumeTypeFromPb(pb *computepb.EbsVolumeTypePb) (*EbsVolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EbsVolumeType(*pb)
	return &st, nil
}

type EditCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	// Wire name: 'apply_policy_default_values'
	ApplyPolicyDefaultValues bool ``
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes ``
	// ID of the cluster
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf ``
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``

	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode ``
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool ``

	// Wire name: 'kind'
	Kind Kind ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
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
	NumWorkers int ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine ``
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
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
	SparkEnvVars map[string]string ``
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool ``

	// Wire name: 'workload_type'
	WorkloadType    *WorkloadType ``
	ForceSendFields []string      `tf:"-"`
}

func (st EditCluster) MarshalJSON() ([]byte, error) {
	pb, err := EditClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EditCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EditClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EditClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EditClusterToPb(st *EditCluster) (*computepb.EditClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EditClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	awsAttributesPb, err := AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	pb.ClusterId = st.ClusterId
	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	dataSecurityModePb, err := DataSecurityModeToPb(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}
	dockerImagePb, err := DockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []computepb.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	kindPb, err := KindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	runtimeEnginePb, err := RuntimeEngineToPb(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	workloadTypePb, err := WorkloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EditClusterFromPb(pb *computepb.EditClusterPb) (*EditCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	awsAttributesField, err := AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	st.ClusterId = pb.ClusterId
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	dataSecurityModeField, err := DataSecurityModeFromPb(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := DockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	gcpAttributesField, err := GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	kindField, err := KindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	runtimeEngineField, err := RuntimeEngineFromPb(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	workloadTypeField, err := WorkloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EditInstancePool struct {
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int ``
	// Instance pool ID
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string ``
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int ``
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int      ``
	ForceSendFields            []string `tf:"-"`
}

func (st EditInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := EditInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EditInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EditInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EditInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EditInstancePoolToPb(st *EditInstancePool) (*computepb.EditInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EditInstancePoolPb{}
	pb.CustomTags = st.CustomTags
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolId = st.InstancePoolId
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EditInstancePoolFromPb(pb *computepb.EditInstancePoolPb) (*EditInstancePool, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EditPolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string ``
	// Additional human-readable description of the cluster policy.
	// Wire name: 'description'
	Description string ``
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	// Wire name: 'libraries'
	Libraries []Library ``
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	// Wire name: 'max_clusters_per_user'
	MaxClustersPerUser int64 ``
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	// Wire name: 'name'
	Name string ``
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
	PolicyFamilyDefinitionOverrides string ``
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string ``
	// The ID of the policy to update.
	// Wire name: 'policy_id'
	PolicyId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st EditPolicy) MarshalJSON() ([]byte, error) {
	pb, err := EditPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EditPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EditPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EditPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EditPolicyToPb(st *EditPolicy) (*computepb.EditPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EditPolicyPb{}
	pb.Definition = st.Definition
	pb.Description = st.Description

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.MaxClustersPerUser = st.MaxClustersPerUser
	pb.Name = st.Name
	pb.PolicyFamilyDefinitionOverrides = st.PolicyFamilyDefinitionOverrides
	pb.PolicyFamilyId = st.PolicyFamilyId
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EditPolicyFromPb(pb *computepb.EditPolicyPb) (*EditPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPolicy{}
	st.Definition = pb.Definition
	st.Description = pb.Description

	var librariesField []Library
	for _, itemPb := range pb.Libraries {
		item, err := LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.MaxClustersPerUser = pb.MaxClustersPerUser
	st.Name = pb.Name
	st.PolicyFamilyDefinitionOverrides = pb.PolicyFamilyDefinitionOverrides
	st.PolicyFamilyId = pb.PolicyFamilyId
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EnforceClusterComplianceRequest struct {
	// The ID of the cluster you want to enforce policy compliance on.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	// Wire name: 'validate_only'
	ValidateOnly    bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st EnforceClusterComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := EnforceClusterComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnforceClusterComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EnforceClusterComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnforceClusterComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnforceClusterComplianceRequestToPb(st *EnforceClusterComplianceRequest) (*computepb.EnforceClusterComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EnforceClusterComplianceRequestPb{}
	pb.ClusterId = st.ClusterId
	pb.ValidateOnly = st.ValidateOnly

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnforceClusterComplianceRequestFromPb(pb *computepb.EnforceClusterComplianceRequestPb) (*EnforceClusterComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforceClusterComplianceRequest{}
	st.ClusterId = pb.ClusterId
	st.ValidateOnly = pb.ValidateOnly

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EnforceClusterComplianceResponse struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	// Wire name: 'changes'
	Changes []ClusterSettingsChange ``
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	// Wire name: 'has_changes'
	HasChanges      bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st EnforceClusterComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := EnforceClusterComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnforceClusterComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EnforceClusterComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnforceClusterComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnforceClusterComplianceResponseToPb(st *EnforceClusterComplianceResponse) (*computepb.EnforceClusterComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EnforceClusterComplianceResponsePb{}

	var changesPb []computepb.ClusterSettingsChangePb
	for _, item := range st.Changes {
		itemPb, err := ClusterSettingsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			changesPb = append(changesPb, *itemPb)
		}
	}
	pb.Changes = changesPb
	pb.HasChanges = st.HasChanges

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnforceClusterComplianceResponseFromPb(pb *computepb.EnforceClusterComplianceResponsePb) (*EnforceClusterComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforceClusterComplianceResponse{}

	var changesField []ClusterSettingsChange
	for _, itemPb := range pb.Changes {
		item, err := ClusterSettingsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			changesField = append(changesField, *item)
		}
	}
	st.Changes = changesField
	st.HasChanges = pb.HasChanges

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. In this minimal environment spec, only pip
// dependencies are supported.
type Environment struct {
	// Use `environment_version` instead.
	// Wire name: 'client'
	Client string ``
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a valid pip requirements file line per
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/.
	// Allowed dependencies include a requirement specifier, an archive URL, a
	// local project path (such as WSFS or UC Volumes in Databricks), or a VCS
	// project URL.
	// Wire name: 'dependencies'
	Dependencies []string ``
	// Required. Environment version used by the environment. Each version comes
	// with a specific Python version and a set of Python packages. The version
	// is a string, consisting of an integer.
	// Wire name: 'environment_version'
	EnvironmentVersion string ``
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	// Wire name: 'jar_dependencies'
	JarDependencies []string ``
	ForceSendFields []string `tf:"-"`
}

func (st Environment) MarshalJSON() ([]byte, error) {
	pb, err := EnvironmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Environment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EnvironmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnvironmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnvironmentToPb(st *Environment) (*computepb.EnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EnvironmentPb{}
	pb.Client = st.Client
	pb.Dependencies = st.Dependencies
	pb.EnvironmentVersion = st.EnvironmentVersion
	pb.JarDependencies = st.JarDependencies

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnvironmentFromPb(pb *computepb.EnvironmentPb) (*Environment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Environment{}
	st.Client = pb.Client
	st.Dependencies = pb.Dependencies
	st.EnvironmentVersion = pb.EnvironmentVersion
	st.JarDependencies = pb.JarDependencies

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type EventDetails struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	// Wire name: 'attributes'
	Attributes *ClusterAttributes ``
	// The cause of a change in target size.
	// Wire name: 'cause'
	Cause EventDetailsCause ``
	// The actual cluster size that was set in the cluster creation or edit.
	// Wire name: 'cluster_size'
	ClusterSize *ClusterSize ``
	// The current number of vCPUs in the cluster.
	// Wire name: 'current_num_vcpus'
	CurrentNumVcpus int ``
	// The current number of nodes in the cluster.
	// Wire name: 'current_num_workers'
	CurrentNumWorkers int ``

	// Wire name: 'did_not_expand_reason'
	DidNotExpandReason string ``
	// Current disk size in bytes
	// Wire name: 'disk_size'
	DiskSize int64 ``
	// More details about the change in driver's state
	// Wire name: 'driver_state_message'
	DriverStateMessage string ``
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	// Wire name: 'enable_termination_for_node_blocklisted'
	EnableTerminationForNodeBlocklisted bool ``

	// Wire name: 'free_space'
	FreeSpace int64 ``
	// List of global and cluster init scripts associated with this cluster
	// event.
	// Wire name: 'init_scripts'
	InitScripts *InitScriptEventDetails ``
	// Instance Id where the event originated from
	// Wire name: 'instance_id'
	InstanceId string ``
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	// Wire name: 'job_run_name'
	JobRunName string ``
	// The cluster attributes before a cluster was edited.
	// Wire name: 'previous_attributes'
	PreviousAttributes *ClusterAttributes ``
	// The size of the cluster before an edit or resize.
	// Wire name: 'previous_cluster_size'
	PreviousClusterSize *ClusterSize ``
	// Previous disk size in bytes
	// Wire name: 'previous_disk_size'
	PreviousDiskSize int64 ``
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	// Wire name: 'reason'
	Reason *TerminationReason ``
	// The targeted number of vCPUs in the cluster.
	// Wire name: 'target_num_vcpus'
	TargetNumVcpus int ``
	// The targeted number of nodes in the cluster.
	// Wire name: 'target_num_workers'
	TargetNumWorkers int ``
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	// Wire name: 'user'
	User            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st EventDetails) MarshalJSON() ([]byte, error) {
	pb, err := EventDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EventDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.EventDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EventDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EventDetailsToPb(st *EventDetails) (*computepb.EventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.EventDetailsPb{}
	attributesPb, err := ClusterAttributesToPb(st.Attributes)
	if err != nil {
		return nil, err
	}
	if attributesPb != nil {
		pb.Attributes = attributesPb
	}
	causePb, err := EventDetailsCauseToPb(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}
	clusterSizePb, err := ClusterSizeToPb(st.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizePb != nil {
		pb.ClusterSize = clusterSizePb
	}
	pb.CurrentNumVcpus = st.CurrentNumVcpus
	pb.CurrentNumWorkers = st.CurrentNumWorkers
	pb.DidNotExpandReason = st.DidNotExpandReason
	pb.DiskSize = st.DiskSize
	pb.DriverStateMessage = st.DriverStateMessage
	pb.EnableTerminationForNodeBlocklisted = st.EnableTerminationForNodeBlocklisted
	pb.FreeSpace = st.FreeSpace
	initScriptsPb, err := InitScriptEventDetailsToPb(st.InitScripts)
	if err != nil {
		return nil, err
	}
	if initScriptsPb != nil {
		pb.InitScripts = initScriptsPb
	}
	pb.InstanceId = st.InstanceId
	pb.JobRunName = st.JobRunName
	previousAttributesPb, err := ClusterAttributesToPb(st.PreviousAttributes)
	if err != nil {
		return nil, err
	}
	if previousAttributesPb != nil {
		pb.PreviousAttributes = previousAttributesPb
	}
	previousClusterSizePb, err := ClusterSizeToPb(st.PreviousClusterSize)
	if err != nil {
		return nil, err
	}
	if previousClusterSizePb != nil {
		pb.PreviousClusterSize = previousClusterSizePb
	}
	pb.PreviousDiskSize = st.PreviousDiskSize
	reasonPb, err := TerminationReasonToPb(st.Reason)
	if err != nil {
		return nil, err
	}
	if reasonPb != nil {
		pb.Reason = reasonPb
	}
	pb.TargetNumVcpus = st.TargetNumVcpus
	pb.TargetNumWorkers = st.TargetNumWorkers
	pb.User = st.User

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EventDetailsFromPb(pb *computepb.EventDetailsPb) (*EventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EventDetails{}
	attributesField, err := ClusterAttributesFromPb(pb.Attributes)
	if err != nil {
		return nil, err
	}
	if attributesField != nil {
		st.Attributes = attributesField
	}
	causeField, err := EventDetailsCauseFromPb(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	clusterSizeField, err := ClusterSizeFromPb(pb.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizeField != nil {
		st.ClusterSize = clusterSizeField
	}
	st.CurrentNumVcpus = pb.CurrentNumVcpus
	st.CurrentNumWorkers = pb.CurrentNumWorkers
	st.DidNotExpandReason = pb.DidNotExpandReason
	st.DiskSize = pb.DiskSize
	st.DriverStateMessage = pb.DriverStateMessage
	st.EnableTerminationForNodeBlocklisted = pb.EnableTerminationForNodeBlocklisted
	st.FreeSpace = pb.FreeSpace
	initScriptsField, err := InitScriptEventDetailsFromPb(pb.InitScripts)
	if err != nil {
		return nil, err
	}
	if initScriptsField != nil {
		st.InitScripts = initScriptsField
	}
	st.InstanceId = pb.InstanceId
	st.JobRunName = pb.JobRunName
	previousAttributesField, err := ClusterAttributesFromPb(pb.PreviousAttributes)
	if err != nil {
		return nil, err
	}
	if previousAttributesField != nil {
		st.PreviousAttributes = previousAttributesField
	}
	previousClusterSizeField, err := ClusterSizeFromPb(pb.PreviousClusterSize)
	if err != nil {
		return nil, err
	}
	if previousClusterSizeField != nil {
		st.PreviousClusterSize = previousClusterSizeField
	}
	st.PreviousDiskSize = pb.PreviousDiskSize
	reasonField, err := TerminationReasonFromPb(pb.Reason)
	if err != nil {
		return nil, err
	}
	if reasonField != nil {
		st.Reason = reasonField
	}
	st.TargetNumVcpus = pb.TargetNumVcpus
	st.TargetNumWorkers = pb.TargetNumWorkers
	st.User = pb.User

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func EventDetailsCauseToPb(st *EventDetailsCause) (*computepb.EventDetailsCausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.EventDetailsCausePb(*st)
	return &pb, nil
}

func EventDetailsCauseFromPb(pb *computepb.EventDetailsCausePb) (*EventDetailsCause, error) {
	if pb == nil {
		return nil, nil
	}
	st := EventDetailsCause(*pb)
	return &st, nil
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

func EventTypeToPb(st *EventType) (*computepb.EventTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.EventTypePb(*st)
	return &pb, nil
}

func EventTypeFromPb(pb *computepb.EventTypePb) (*EventType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EventType(*pb)
	return &st, nil
}

// Attributes set during cluster creation which are related to GCP.
type GcpAttributes struct {
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs, on-demand VMs, or preemptible VMs with a fallback
	// to on-demand VMs if the former is unavailable.
	// Wire name: 'availability'
	Availability GcpAvailability ``
	// Boot disk size in GB
	// Wire name: 'boot_disk_size'
	BootDiskSize int ``
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	// Wire name: 'google_service_account'
	GoogleServiceAccount string ``
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	// Wire name: 'local_ssd_count'
	LocalSsdCount int ``
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// 'availability' field instead.
	// Wire name: 'use_preemptible_executors'
	UsePreemptibleExecutors bool ``
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default]. -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	// Wire name: 'zone_id'
	ZoneId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GcpAttributes) MarshalJSON() ([]byte, error) {
	pb, err := GcpAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GcpAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GcpAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GcpAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GcpAttributesToPb(st *GcpAttributes) (*computepb.GcpAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GcpAttributesPb{}
	availabilityPb, err := GcpAvailabilityToPb(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}
	pb.BootDiskSize = st.BootDiskSize
	pb.GoogleServiceAccount = st.GoogleServiceAccount
	pb.LocalSsdCount = st.LocalSsdCount
	pb.UsePreemptibleExecutors = st.UsePreemptibleExecutors
	pb.ZoneId = st.ZoneId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GcpAttributesFromPb(pb *computepb.GcpAttributesPb) (*GcpAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpAttributes{}
	availabilityField, err := GcpAvailabilityFromPb(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	st.BootDiskSize = pb.BootDiskSize
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.LocalSsdCount = pb.LocalSsdCount
	st.UsePreemptibleExecutors = pb.UsePreemptibleExecutors
	st.ZoneId = pb.ZoneId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func GcpAvailabilityToPb(st *GcpAvailability) (*computepb.GcpAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.GcpAvailabilityPb(*st)
	return &pb, nil
}

func GcpAvailabilityFromPb(pb *computepb.GcpAvailabilityPb) (*GcpAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := GcpAvailability(*pb)
	return &st, nil
}

// A storage location in Google Cloud Platform's GCS
type GcsStorageInfo struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	// Wire name: 'destination'
	Destination string ``
}

func (st GcsStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := GcsStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GcsStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GcsStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GcsStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GcsStorageInfoToPb(st *GcsStorageInfo) (*computepb.GcsStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GcsStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

func GcsStorageInfoFromPb(pb *computepb.GcsStorageInfoPb) (*GcsStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcsStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

type GetClusterComplianceRequest struct {
	// The ID of the cluster to get the compliance status
	// Wire name: 'cluster_id'
	ClusterId string `tf:"-"`
}

func (st GetClusterComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterComplianceRequestToPb(st *GetClusterComplianceRequest) (*computepb.GetClusterComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterComplianceRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func GetClusterComplianceRequestFromPb(pb *computepb.GetClusterComplianceRequestPb) (*GetClusterComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterComplianceRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type GetClusterComplianceResponse struct {
	// Whether the cluster is compliant with its policy or not. Clusters could
	// be out of compliance if the policy was updated after the cluster was last
	// edited.
	// Wire name: 'is_compliant'
	IsCompliant bool ``
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	// Wire name: 'violations'
	Violations      map[string]string ``
	ForceSendFields []string          `tf:"-"`
}

func (st GetClusterComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterComplianceResponseToPb(st *GetClusterComplianceResponse) (*computepb.GetClusterComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterComplianceResponsePb{}
	pb.IsCompliant = st.IsCompliant
	pb.Violations = st.Violations

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetClusterComplianceResponseFromPb(pb *computepb.GetClusterComplianceResponsePb) (*GetClusterComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterComplianceResponse{}
	st.IsCompliant = pb.IsCompliant
	st.Violations = pb.Violations

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	// Wire name: 'cluster_id'
	ClusterId string `tf:"-"`
}

func (st GetClusterPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPermissionLevelsRequestToPb(st *GetClusterPermissionLevelsRequest) (*computepb.GetClusterPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPermissionLevelsRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func GetClusterPermissionLevelsRequestFromPb(pb *computepb.GetClusterPermissionLevelsRequestPb) (*GetClusterPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionLevelsRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type GetClusterPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ClusterPermissionsDescription ``
}

func (st GetClusterPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPermissionLevelsResponseToPb(st *GetClusterPermissionLevelsResponse) (*computepb.GetClusterPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPermissionLevelsResponsePb{}

	var permissionLevelsPb []computepb.ClusterPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := ClusterPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetClusterPermissionLevelsResponseFromPb(pb *computepb.GetClusterPermissionLevelsResponsePb) (*GetClusterPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionLevelsResponse{}

	var permissionLevelsField []ClusterPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := ClusterPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	// Wire name: 'cluster_id'
	ClusterId string `tf:"-"`
}

func (st GetClusterPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPermissionsRequestToPb(st *GetClusterPermissionsRequest) (*computepb.GetClusterPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPermissionsRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func GetClusterPermissionsRequestFromPb(pb *computepb.GetClusterPermissionsRequestPb) (*GetClusterPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionsRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	// Wire name: 'cluster_policy_id'
	ClusterPolicyId string `tf:"-"`
}

func (st GetClusterPolicyPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPolicyPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPolicyPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPolicyPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPolicyPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPolicyPermissionLevelsRequestToPb(st *GetClusterPolicyPermissionLevelsRequest) (*computepb.GetClusterPolicyPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPolicyPermissionLevelsRequestPb{}
	pb.ClusterPolicyId = st.ClusterPolicyId

	return pb, nil
}

func GetClusterPolicyPermissionLevelsRequestFromPb(pb *computepb.GetClusterPolicyPermissionLevelsRequestPb) (*GetClusterPolicyPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionLevelsRequest{}
	st.ClusterPolicyId = pb.ClusterPolicyId

	return st, nil
}

type GetClusterPolicyPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ClusterPolicyPermissionsDescription ``
}

func (st GetClusterPolicyPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPolicyPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPolicyPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPolicyPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPolicyPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPolicyPermissionLevelsResponseToPb(st *GetClusterPolicyPermissionLevelsResponse) (*computepb.GetClusterPolicyPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPolicyPermissionLevelsResponsePb{}

	var permissionLevelsPb []computepb.ClusterPolicyPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := ClusterPolicyPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetClusterPolicyPermissionLevelsResponseFromPb(pb *computepb.GetClusterPolicyPermissionLevelsResponsePb) (*GetClusterPolicyPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionLevelsResponse{}

	var permissionLevelsField []ClusterPolicyPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := ClusterPolicyPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	// Wire name: 'cluster_policy_id'
	ClusterPolicyId string `tf:"-"`
}

func (st GetClusterPolicyPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPolicyPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPolicyPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPolicyPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPolicyPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPolicyPermissionsRequestToPb(st *GetClusterPolicyPermissionsRequest) (*computepb.GetClusterPolicyPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPolicyPermissionsRequestPb{}
	pb.ClusterPolicyId = st.ClusterPolicyId

	return pb, nil
}

func GetClusterPolicyPermissionsRequestFromPb(pb *computepb.GetClusterPolicyPermissionsRequestPb) (*GetClusterPolicyPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionsRequest{}
	st.ClusterPolicyId = pb.ClusterPolicyId

	return st, nil
}

type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the Cluster Policy.
	// Wire name: 'policy_id'
	PolicyId string `tf:"-"`
}

func (st GetClusterPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterPolicyRequestToPb(st *GetClusterPolicyRequest) (*computepb.GetClusterPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func GetClusterPolicyRequestFromPb(pb *computepb.GetClusterPolicyRequestPb) (*GetClusterPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	// Wire name: 'cluster_id'
	ClusterId string `tf:"-"`
}

func (st GetClusterRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetClusterRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetClusterRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetClusterRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetClusterRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetClusterRequestToPb(st *GetClusterRequest) (*computepb.GetClusterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetClusterRequestPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func GetClusterRequestFromPb(pb *computepb.GetClusterRequestPb) (*GetClusterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterRequest{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type GetEvents struct {
	// The ID of the cluster to retrieve events about.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	// Wire name: 'end_time'
	EndTime int64 ``
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	// Wire name: 'event_types'
	EventTypes []EventType ``
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	// Wire name: 'limit'
	Limit int64 ``
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	// Wire name: 'offset'
	Offset int64 ``
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	// Wire name: 'order'
	Order GetEventsOrder ``
	// The maximum number of events to include in a page of events. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is empty or 0, the server will decide the number
	// of results to be returned. The field has to be in the range [0,500]. If
	// the value is outside the range, the server enforces 0 or 500.
	// Wire name: 'page_size'
	PageSize int ``
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of events respectively. If page_token
	// is empty, the first page is returned.
	// Wire name: 'page_token'
	PageToken string ``
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	// Wire name: 'start_time'
	StartTime       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st GetEvents) MarshalJSON() ([]byte, error) {
	pb, err := GetEventsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetEvents) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetEventsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetEventsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetEventsToPb(st *GetEvents) (*computepb.GetEventsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetEventsPb{}
	pb.ClusterId = st.ClusterId
	pb.EndTime = st.EndTime

	var eventTypesPb []computepb.EventTypePb
	for _, item := range st.EventTypes {
		itemPb, err := EventTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventTypesPb = append(eventTypesPb, *itemPb)
		}
	}
	pb.EventTypes = eventTypesPb
	pb.Limit = st.Limit
	pb.Offset = st.Offset
	orderPb, err := GetEventsOrderToPb(&st.Order)
	if err != nil {
		return nil, err
	}
	if orderPb != nil {
		pb.Order = *orderPb
	}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.StartTime = st.StartTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetEventsFromPb(pb *computepb.GetEventsPb) (*GetEvents, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEvents{}
	st.ClusterId = pb.ClusterId
	st.EndTime = pb.EndTime

	var eventTypesField []EventType
	for _, itemPb := range pb.EventTypes {
		item, err := EventTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventTypesField = append(eventTypesField, *item)
		}
	}
	st.EventTypes = eventTypesField
	st.Limit = pb.Limit
	st.Offset = pb.Offset
	orderField, err := GetEventsOrderFromPb(&pb.Order)
	if err != nil {
		return nil, err
	}
	if orderField != nil {
		st.Order = *orderField
	}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.StartTime = pb.StartTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func GetEventsOrderToPb(st *GetEventsOrder) (*computepb.GetEventsOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.GetEventsOrderPb(*st)
	return &pb, nil
}

func GetEventsOrderFromPb(pb *computepb.GetEventsOrderPb) (*GetEventsOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetEventsOrder(*pb)
	return &st, nil
}

type GetEventsResponse struct {

	// Wire name: 'events'
	Events []ClusterEvent ``
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	// Wire name: 'next_page'
	NextPage *GetEvents ``
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	// Wire name: 'prev_page_token'
	PrevPageToken string ``
	// Deprecated: Returns 0 when request uses page_token. Will start returning
	// zero when request uses offset/limit soon.
	//
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	// Wire name: 'total_count'
	TotalCount      int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st GetEventsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetEventsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetEventsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetEventsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetEventsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetEventsResponseToPb(st *GetEventsResponse) (*computepb.GetEventsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetEventsResponsePb{}

	var eventsPb []computepb.ClusterEventPb
	for _, item := range st.Events {
		itemPb, err := ClusterEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb
	nextPagePb, err := GetEventsToPb(st.NextPage)
	if err != nil {
		return nil, err
	}
	if nextPagePb != nil {
		pb.NextPage = nextPagePb
	}
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken
	pb.TotalCount = st.TotalCount

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetEventsResponseFromPb(pb *computepb.GetEventsResponsePb) (*GetEventsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEventsResponse{}

	var eventsField []ClusterEvent
	for _, itemPb := range pb.Events {
		item, err := ClusterEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	nextPageField, err := GetEventsFromPb(pb.NextPage)
	if err != nil {
		return nil, err
	}
	if nextPageField != nil {
		st.NextPage = nextPageField
	}
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken
	st.TotalCount = pb.TotalCount

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	// Wire name: 'script_id'
	ScriptId string `tf:"-"`
}

func (st GetGlobalInitScriptRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetGlobalInitScriptRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetGlobalInitScriptRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetGlobalInitScriptRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetGlobalInitScriptRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetGlobalInitScriptRequestToPb(st *GetGlobalInitScriptRequest) (*computepb.GetGlobalInitScriptRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetGlobalInitScriptRequestPb{}
	pb.ScriptId = st.ScriptId

	return pb, nil
}

func GetGlobalInitScriptRequestFromPb(pb *computepb.GetGlobalInitScriptRequestPb) (*GetGlobalInitScriptRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGlobalInitScriptRequest{}
	st.ScriptId = pb.ScriptId

	return st, nil
}

type GetInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *InstancePoolAwsAttributes ``
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *InstancePoolAzureAttributes ``
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
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
	DefaultTags map[string]string ``
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	// Wire name: 'disk_spec'
	DiskSpec *DiskSpec ``
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *InstancePoolGcpAttributes ``
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int ``
	// Canonical unique identifier for the pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string ``
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int ``
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
	// Custom Docker Image BYOC
	// Wire name: 'preloaded_docker_images'
	PreloadedDockerImages []DockerImage ``
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'preloaded_spark_versions'
	PreloadedSparkVersions []string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Current state of the instance pool.
	// Wire name: 'state'
	State InstancePoolState ``
	// Usage statistics about the instance pool.
	// Wire name: 'stats'
	Stats *InstancePoolStats ``
	// Status of failed pending instances in the pool.
	// Wire name: 'status'
	Status *InstancePoolStatus ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int      ``
	ForceSendFields            []string `tf:"-"`
}

func (st GetInstancePool) MarshalJSON() ([]byte, error) {
	pb, err := GetInstancePoolToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetInstancePool) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetInstancePoolPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetInstancePoolFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetInstancePoolToPb(st *GetInstancePool) (*computepb.GetInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetInstancePoolPb{}
	awsAttributesPb, err := InstancePoolAwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := InstancePoolAzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	pb.CustomTags = st.CustomTags
	pb.DefaultTags = st.DefaultTags
	diskSpecPb, err := DiskSpecToPb(st.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecPb != nil {
		pb.DiskSpec = diskSpecPb
	}
	pb.EnableElasticDisk = st.EnableElasticDisk
	gcpAttributesPb, err := InstancePoolGcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolId = st.InstancePoolId
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId

	var preloadedDockerImagesPb []computepb.DockerImagePb
	for _, item := range st.PreloadedDockerImages {
		itemPb, err := DockerImageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedDockerImagesPb = append(preloadedDockerImagesPb, *itemPb)
		}
	}
	pb.PreloadedDockerImages = preloadedDockerImagesPb
	pb.PreloadedSparkVersions = st.PreloadedSparkVersions
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	statePb, err := InstancePoolStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	statsPb, err := InstancePoolStatsToPb(st.Stats)
	if err != nil {
		return nil, err
	}
	if statsPb != nil {
		pb.Stats = statsPb
	}
	statusPb, err := InstancePoolStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetInstancePoolFromPb(pb *computepb.GetInstancePoolPb) (*GetInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePool{}
	awsAttributesField, err := InstancePoolAwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := InstancePoolAzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	st.CustomTags = pb.CustomTags
	st.DefaultTags = pb.DefaultTags
	diskSpecField, err := DiskSpecFromPb(pb.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecField != nil {
		st.DiskSpec = diskSpecField
	}
	st.EnableElasticDisk = pb.EnableElasticDisk
	gcpAttributesField, err := InstancePoolGcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolId = pb.InstancePoolId
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId

	var preloadedDockerImagesField []DockerImage
	for _, itemPb := range pb.PreloadedDockerImages {
		item, err := DockerImageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedDockerImagesField = append(preloadedDockerImagesField, *item)
		}
	}
	st.PreloadedDockerImages = preloadedDockerImagesField
	st.PreloadedSparkVersions = pb.PreloadedSparkVersions
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	stateField, err := InstancePoolStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	statsField, err := InstancePoolStatsFromPb(pb.Stats)
	if err != nil {
		return nil, err
	}
	if statsField != nil {
		st.Stats = statsField
	}
	statusField, err := InstancePoolStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `tf:"-"`
}

func (st GetInstancePoolPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetInstancePoolPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetInstancePoolPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetInstancePoolPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetInstancePoolPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetInstancePoolPermissionLevelsRequestToPb(st *GetInstancePoolPermissionLevelsRequest) (*computepb.GetInstancePoolPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetInstancePoolPermissionLevelsRequestPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

func GetInstancePoolPermissionLevelsRequestFromPb(pb *computepb.GetInstancePoolPermissionLevelsRequestPb) (*GetInstancePoolPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionLevelsRequest{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

type GetInstancePoolPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []InstancePoolPermissionsDescription ``
}

func (st GetInstancePoolPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetInstancePoolPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetInstancePoolPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetInstancePoolPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetInstancePoolPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetInstancePoolPermissionLevelsResponseToPb(st *GetInstancePoolPermissionLevelsResponse) (*computepb.GetInstancePoolPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetInstancePoolPermissionLevelsResponsePb{}

	var permissionLevelsPb []computepb.InstancePoolPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := InstancePoolPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetInstancePoolPermissionLevelsResponseFromPb(pb *computepb.GetInstancePoolPermissionLevelsResponsePb) (*GetInstancePoolPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionLevelsResponse{}

	var permissionLevelsField []InstancePoolPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := InstancePoolPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `tf:"-"`
}

func (st GetInstancePoolPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetInstancePoolPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetInstancePoolPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetInstancePoolPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetInstancePoolPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetInstancePoolPermissionsRequestToPb(st *GetInstancePoolPermissionsRequest) (*computepb.GetInstancePoolPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetInstancePoolPermissionsRequestPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

func GetInstancePoolPermissionsRequestFromPb(pb *computepb.GetInstancePoolPermissionsRequestPb) (*GetInstancePoolPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionsRequest{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `tf:"-"`
}

func (st GetInstancePoolRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetInstancePoolRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetInstancePoolRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetInstancePoolRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetInstancePoolRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetInstancePoolRequestToPb(st *GetInstancePoolRequest) (*computepb.GetInstancePoolRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetInstancePoolRequestPb{}
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

func GetInstancePoolRequestFromPb(pb *computepb.GetInstancePoolRequestPb) (*GetInstancePoolRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolRequest{}
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
}

type GetPolicyFamilyRequest struct {
	// The family ID about which to retrieve information.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string `tf:"-"`
	// The version number for the family to fetch. Defaults to the latest
	// version.
	// Wire name: 'version'
	Version         int64    `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GetPolicyFamilyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPolicyFamilyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPolicyFamilyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetPolicyFamilyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPolicyFamilyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPolicyFamilyRequestToPb(st *GetPolicyFamilyRequest) (*computepb.GetPolicyFamilyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetPolicyFamilyRequestPb{}
	pb.PolicyFamilyId = st.PolicyFamilyId
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetPolicyFamilyRequestFromPb(pb *computepb.GetPolicyFamilyRequestPb) (*GetPolicyFamilyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyFamilyRequest{}
	st.PolicyFamilyId = pb.PolicyFamilyId
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetSparkVersionsResponse struct {
	// All the available Spark versions.
	// Wire name: 'versions'
	Versions []SparkVersion ``
}

func (st GetSparkVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetSparkVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetSparkVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GetSparkVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetSparkVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetSparkVersionsResponseToPb(st *GetSparkVersionsResponse) (*computepb.GetSparkVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GetSparkVersionsResponsePb{}

	var versionsPb []computepb.SparkVersionPb
	for _, item := range st.Versions {
		itemPb, err := SparkVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			versionsPb = append(versionsPb, *itemPb)
		}
	}
	pb.Versions = versionsPb

	return pb, nil
}

func GetSparkVersionsResponseFromPb(pb *computepb.GetSparkVersionsResponsePb) (*GetSparkVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSparkVersionsResponse{}

	var versionsField []SparkVersion
	for _, itemPb := range pb.Versions {
		item, err := SparkVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			versionsField = append(versionsField, *item)
		}
	}
	st.Versions = versionsField

	return st, nil
}

type GlobalInitScriptCreateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the script
	// Wire name: 'name'
	Name string ``
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
	Position int ``
	// The Base64-encoded content of the script.
	// Wire name: 'script'
	Script          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GlobalInitScriptCreateRequest) MarshalJSON() ([]byte, error) {
	pb, err := GlobalInitScriptCreateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GlobalInitScriptCreateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GlobalInitScriptCreateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GlobalInitScriptCreateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GlobalInitScriptCreateRequestToPb(st *GlobalInitScriptCreateRequest) (*computepb.GlobalInitScriptCreateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GlobalInitScriptCreateRequestPb{}
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.Script = st.Script

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GlobalInitScriptCreateRequestFromPb(pb *computepb.GlobalInitScriptCreateRequestPb) (*GlobalInitScriptCreateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptCreateRequest{}
	st.Enabled = pb.Enabled
	st.Name = pb.Name
	st.Position = pb.Position
	st.Script = pb.Script

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GlobalInitScriptDetails struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'created_at'
	CreatedAt int ``
	// The username of the user who created the script.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the script
	// Wire name: 'name'
	Name string ``
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	// Wire name: 'position'
	Position int ``
	// The global init script ID.
	// Wire name: 'script_id'
	ScriptId string ``
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int ``
	// The username of the user who last updated the script
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GlobalInitScriptDetails) MarshalJSON() ([]byte, error) {
	pb, err := GlobalInitScriptDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GlobalInitScriptDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GlobalInitScriptDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GlobalInitScriptDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GlobalInitScriptDetailsToPb(st *GlobalInitScriptDetails) (*computepb.GlobalInitScriptDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GlobalInitScriptDetailsPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.ScriptId = st.ScriptId
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GlobalInitScriptDetailsFromPb(pb *computepb.GlobalInitScriptDetailsPb) (*GlobalInitScriptDetails, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'created_at'
	CreatedAt int ``
	// The username of the user who created the script.
	// Wire name: 'created_by'
	CreatedBy string ``
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the script
	// Wire name: 'name'
	Name string ``
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	// Wire name: 'position'
	Position int ``
	// The Base64-encoded content of the script.
	// Wire name: 'script'
	Script string ``
	// The global init script ID.
	// Wire name: 'script_id'
	ScriptId string ``
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int ``
	// The username of the user who last updated the script
	// Wire name: 'updated_by'
	UpdatedBy       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GlobalInitScriptDetailsWithContent) MarshalJSON() ([]byte, error) {
	pb, err := GlobalInitScriptDetailsWithContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GlobalInitScriptDetailsWithContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GlobalInitScriptDetailsWithContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GlobalInitScriptDetailsWithContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GlobalInitScriptDetailsWithContentToPb(st *GlobalInitScriptDetailsWithContent) (*computepb.GlobalInitScriptDetailsWithContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GlobalInitScriptDetailsWithContentPb{}
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.Script = st.Script
	pb.ScriptId = st.ScriptId
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GlobalInitScriptDetailsWithContentFromPb(pb *computepb.GlobalInitScriptDetailsWithContentPb) (*GlobalInitScriptDetailsWithContent, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GlobalInitScriptUpdateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the script
	// Wire name: 'name'
	Name string ``
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
	Position int ``
	// The Base64-encoded content of the script.
	// Wire name: 'script'
	Script string ``
	// The ID of the global init script.
	// Wire name: 'script_id'
	ScriptId        string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GlobalInitScriptUpdateRequest) MarshalJSON() ([]byte, error) {
	pb, err := GlobalInitScriptUpdateRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GlobalInitScriptUpdateRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.GlobalInitScriptUpdateRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GlobalInitScriptUpdateRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GlobalInitScriptUpdateRequestToPb(st *GlobalInitScriptUpdateRequest) (*computepb.GlobalInitScriptUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.GlobalInitScriptUpdateRequestPb{}
	pb.Enabled = st.Enabled
	pb.Name = st.Name
	pb.Position = st.Position
	pb.Script = st.Script
	pb.ScriptId = st.ScriptId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GlobalInitScriptUpdateRequestFromPb(pb *computepb.GlobalInitScriptUpdateRequestPb) (*GlobalInitScriptUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptUpdateRequest{}
	st.Enabled = pb.Enabled
	st.Name = pb.Name
	st.Position = pb.Position
	st.Script = pb.Script
	st.ScriptId = pb.ScriptId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InitScriptEventDetails struct {
	// The cluster scoped init scripts associated with this cluster event.
	// Wire name: 'cluster'
	Cluster []InitScriptInfoAndExecutionDetails ``
	// The global init scripts associated with this cluster event.
	// Wire name: 'global'
	Global []InitScriptInfoAndExecutionDetails ``
	// The private ip of the node we are reporting init script execution details
	// for (we will select the execution details from only one node rather than
	// reporting the execution details from every node to keep these event
	// details small)
	//
	// This should only be defined for the INIT_SCRIPTS_FINISHED event
	// Wire name: 'reported_for_node'
	ReportedForNode string   ``
	ForceSendFields []string `tf:"-"`
}

func (st InitScriptEventDetails) MarshalJSON() ([]byte, error) {
	pb, err := InitScriptEventDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InitScriptEventDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InitScriptEventDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InitScriptEventDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InitScriptEventDetailsToPb(st *InitScriptEventDetails) (*computepb.InitScriptEventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InitScriptEventDetailsPb{}

	var clusterPb []computepb.InitScriptInfoAndExecutionDetailsPb
	for _, item := range st.Cluster {
		itemPb, err := InitScriptInfoAndExecutionDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clusterPb = append(clusterPb, *itemPb)
		}
	}
	pb.Cluster = clusterPb

	var globalPb []computepb.InitScriptInfoAndExecutionDetailsPb
	for _, item := range st.Global {
		itemPb, err := InitScriptInfoAndExecutionDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			globalPb = append(globalPb, *itemPb)
		}
	}
	pb.Global = globalPb
	pb.ReportedForNode = st.ReportedForNode

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InitScriptEventDetailsFromPb(pb *computepb.InitScriptEventDetailsPb) (*InitScriptEventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptEventDetails{}

	var clusterField []InitScriptInfoAndExecutionDetails
	for _, itemPb := range pb.Cluster {
		item, err := InitScriptInfoAndExecutionDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clusterField = append(clusterField, *item)
		}
	}
	st.Cluster = clusterField

	var globalField []InitScriptInfoAndExecutionDetails
	for _, itemPb := range pb.Global {
		item, err := InitScriptInfoAndExecutionDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			globalField = append(globalField, *item)
		}
	}
	st.Global = globalField
	st.ReportedForNode = pb.ReportedForNode

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func InitScriptExecutionDetailsInitScriptExecutionStatusToPb(st *InitScriptExecutionDetailsInitScriptExecutionStatus) (*computepb.InitScriptExecutionDetailsInitScriptExecutionStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.InitScriptExecutionDetailsInitScriptExecutionStatusPb(*st)
	return &pb, nil
}

func InitScriptExecutionDetailsInitScriptExecutionStatusFromPb(pb *computepb.InitScriptExecutionDetailsInitScriptExecutionStatusPb) (*InitScriptExecutionDetailsInitScriptExecutionStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := InitScriptExecutionDetailsInitScriptExecutionStatus(*pb)
	return &st, nil
}

// Config for an individual init script Next ID: 11
type InitScriptInfo struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	// Wire name: 'abfss'
	Abfss *Adlsgen2Info ``
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	// Wire name: 'dbfs'
	Dbfs *DbfsStorageInfo ``
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	// Wire name: 'file'
	File *LocalFileInfo ``
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	// Wire name: 'gcs'
	Gcs *GcsStorageInfo ``
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	// Wire name: 's3'
	S3 *S3StorageInfo ``
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	// Wire name: 'volumes'
	Volumes *VolumesStorageInfo ``
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	// Wire name: 'workspace'
	Workspace *WorkspaceStorageInfo ``
}

func (st InitScriptInfo) MarshalJSON() ([]byte, error) {
	pb, err := InitScriptInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InitScriptInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InitScriptInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InitScriptInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InitScriptInfoToPb(st *InitScriptInfo) (*computepb.InitScriptInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InitScriptInfoPb{}
	abfssPb, err := Adlsgen2InfoToPb(st.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssPb != nil {
		pb.Abfss = abfssPb
	}
	dbfsPb, err := DbfsStorageInfoToPb(st.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsPb != nil {
		pb.Dbfs = dbfsPb
	}
	filePb, err := LocalFileInfoToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}
	gcsPb, err := GcsStorageInfoToPb(st.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsPb != nil {
		pb.Gcs = gcsPb
	}
	s3Pb, err := S3StorageInfoToPb(st.S3)
	if err != nil {
		return nil, err
	}
	if s3Pb != nil {
		pb.S3 = s3Pb
	}
	volumesPb, err := VolumesStorageInfoToPb(st.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesPb != nil {
		pb.Volumes = volumesPb
	}
	workspacePb, err := WorkspaceStorageInfoToPb(st.Workspace)
	if err != nil {
		return nil, err
	}
	if workspacePb != nil {
		pb.Workspace = workspacePb
	}

	return pb, nil
}

func InitScriptInfoFromPb(pb *computepb.InitScriptInfoPb) (*InitScriptInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptInfo{}
	abfssField, err := Adlsgen2InfoFromPb(pb.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssField != nil {
		st.Abfss = abfssField
	}
	dbfsField, err := DbfsStorageInfoFromPb(pb.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsField != nil {
		st.Dbfs = dbfsField
	}
	fileField, err := LocalFileInfoFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	gcsField, err := GcsStorageInfoFromPb(pb.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsField != nil {
		st.Gcs = gcsField
	}
	s3Field, err := S3StorageInfoFromPb(pb.S3)
	if err != nil {
		return nil, err
	}
	if s3Field != nil {
		st.S3 = s3Field
	}
	volumesField, err := VolumesStorageInfoFromPb(pb.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesField != nil {
		st.Volumes = volumesField
	}
	workspaceField, err := WorkspaceStorageInfoFromPb(pb.Workspace)
	if err != nil {
		return nil, err
	}
	if workspaceField != nil {
		st.Workspace = workspaceField
	}

	return st, nil
}

type InitScriptInfoAndExecutionDetails struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	// Wire name: 'abfss'
	Abfss *Adlsgen2Info ``
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	// Wire name: 'dbfs'
	Dbfs *DbfsStorageInfo ``
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	// Wire name: 'error_message'
	ErrorMessage string ``
	// The number duration of the script execution in seconds
	// Wire name: 'execution_duration_seconds'
	ExecutionDurationSeconds int ``
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	// Wire name: 'file'
	File *LocalFileInfo ``
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	// Wire name: 'gcs'
	Gcs *GcsStorageInfo ``
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	// Wire name: 's3'
	S3 *S3StorageInfo ``
	// The current status of the script
	// Wire name: 'status'
	Status InitScriptExecutionDetailsInitScriptExecutionStatus ``
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	// Wire name: 'volumes'
	Volumes *VolumesStorageInfo ``
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	// Wire name: 'workspace'
	Workspace       *WorkspaceStorageInfo ``
	ForceSendFields []string              `tf:"-"`
}

func (st InitScriptInfoAndExecutionDetails) MarshalJSON() ([]byte, error) {
	pb, err := InitScriptInfoAndExecutionDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InitScriptInfoAndExecutionDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InitScriptInfoAndExecutionDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InitScriptInfoAndExecutionDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InitScriptInfoAndExecutionDetailsToPb(st *InitScriptInfoAndExecutionDetails) (*computepb.InitScriptInfoAndExecutionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InitScriptInfoAndExecutionDetailsPb{}
	abfssPb, err := Adlsgen2InfoToPb(st.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssPb != nil {
		pb.Abfss = abfssPb
	}
	dbfsPb, err := DbfsStorageInfoToPb(st.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsPb != nil {
		pb.Dbfs = dbfsPb
	}
	pb.ErrorMessage = st.ErrorMessage
	pb.ExecutionDurationSeconds = st.ExecutionDurationSeconds
	filePb, err := LocalFileInfoToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}
	gcsPb, err := GcsStorageInfoToPb(st.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsPb != nil {
		pb.Gcs = gcsPb
	}
	s3Pb, err := S3StorageInfoToPb(st.S3)
	if err != nil {
		return nil, err
	}
	if s3Pb != nil {
		pb.S3 = s3Pb
	}
	statusPb, err := InitScriptExecutionDetailsInitScriptExecutionStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	volumesPb, err := VolumesStorageInfoToPb(st.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesPb != nil {
		pb.Volumes = volumesPb
	}
	workspacePb, err := WorkspaceStorageInfoToPb(st.Workspace)
	if err != nil {
		return nil, err
	}
	if workspacePb != nil {
		pb.Workspace = workspacePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InitScriptInfoAndExecutionDetailsFromPb(pb *computepb.InitScriptInfoAndExecutionDetailsPb) (*InitScriptInfoAndExecutionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptInfoAndExecutionDetails{}
	abfssField, err := Adlsgen2InfoFromPb(pb.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssField != nil {
		st.Abfss = abfssField
	}
	dbfsField, err := DbfsStorageInfoFromPb(pb.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsField != nil {
		st.Dbfs = dbfsField
	}
	st.ErrorMessage = pb.ErrorMessage
	st.ExecutionDurationSeconds = pb.ExecutionDurationSeconds
	fileField, err := LocalFileInfoFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	gcsField, err := GcsStorageInfoFromPb(pb.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsField != nil {
		st.Gcs = gcsField
	}
	s3Field, err := S3StorageInfoFromPb(pb.S3)
	if err != nil {
		return nil, err
	}
	if s3Field != nil {
		st.S3 = s3Field
	}
	statusField, err := InitScriptExecutionDetailsInitScriptExecutionStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	volumesField, err := VolumesStorageInfoFromPb(pb.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesField != nil {
		st.Volumes = volumesField
	}
	workspaceField, err := WorkspaceStorageInfoFromPb(pb.Workspace)
	if err != nil {
		return nil, err
	}
	if workspaceField != nil {
		st.Workspace = workspaceField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The libraries to install.
	// Wire name: 'libraries'
	Libraries []Library ``
}

func (st InstallLibraries) MarshalJSON() ([]byte, error) {
	pb, err := InstallLibrariesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstallLibraries) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstallLibrariesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstallLibrariesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstallLibrariesToPb(st *InstallLibraries) (*computepb.InstallLibrariesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstallLibrariesPb{}
	pb.ClusterId = st.ClusterId

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	return pb, nil
}

func InstallLibrariesFromPb(pb *computepb.InstallLibrariesPb) (*InstallLibraries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstallLibraries{}
	st.ClusterId = pb.ClusterId

	var librariesField []Library
	for _, itemPb := range pb.Libraries {
		item, err := LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField

	return st, nil
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel InstancePoolPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolAccessControlRequestToPb(st *InstancePoolAccessControlRequest) (*computepb.InstancePoolAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := InstancePoolPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolAccessControlRequestFromPb(pb *computepb.InstancePoolAccessControlRequestPb) (*InstancePoolAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := InstancePoolPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstancePoolAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []InstancePoolPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolAccessControlResponseToPb(st *InstancePoolAccessControlResponse) (*computepb.InstancePoolAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolAccessControlResponsePb{}

	var allPermissionsPb []computepb.InstancePoolPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := InstancePoolPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolAccessControlResponseFromPb(pb *computepb.InstancePoolAccessControlResponsePb) (*InstancePoolAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAccessControlResponse{}

	var allPermissionsField []InstancePoolPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := InstancePoolPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *InstancePoolAwsAttributes ``
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *InstancePoolAzureAttributes ``
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``
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
	DefaultTags map[string]string ``
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	// Wire name: 'disk_spec'
	DiskSpec *DiskSpec ``
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *InstancePoolGcpAttributes ``
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	// Wire name: 'idle_instance_autotermination_minutes'
	IdleInstanceAutoterminationMinutes int ``
	// Canonical unique identifier for the pool.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	// Wire name: 'instance_pool_name'
	InstancePoolName string ``
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	// Wire name: 'max_capacity'
	MaxCapacity int ``
	// Minimum number of idle instances to keep in the instance pool
	// Wire name: 'min_idle_instances'
	MinIdleInstances int ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
	// Custom Docker Image BYOC
	// Wire name: 'preloaded_docker_images'
	PreloadedDockerImages []DockerImage ``
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'preloaded_spark_versions'
	PreloadedSparkVersions []string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Current state of the instance pool.
	// Wire name: 'state'
	State InstancePoolState ``
	// Usage statistics about the instance pool.
	// Wire name: 'stats'
	Stats *InstancePoolStats ``
	// Status of failed pending instances in the pool.
	// Wire name: 'status'
	Status *InstancePoolStatus ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED types.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int      ``
	ForceSendFields            []string `tf:"-"`
}

func (st InstancePoolAndStats) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolAndStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolAndStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolAndStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolAndStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolAndStatsToPb(st *InstancePoolAndStats) (*computepb.InstancePoolAndStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolAndStatsPb{}
	awsAttributesPb, err := InstancePoolAwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := InstancePoolAzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	pb.CustomTags = st.CustomTags
	pb.DefaultTags = st.DefaultTags
	diskSpecPb, err := DiskSpecToPb(st.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecPb != nil {
		pb.DiskSpec = diskSpecPb
	}
	pb.EnableElasticDisk = st.EnableElasticDisk
	gcpAttributesPb, err := InstancePoolGcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}
	pb.IdleInstanceAutoterminationMinutes = st.IdleInstanceAutoterminationMinutes
	pb.InstancePoolId = st.InstancePoolId
	pb.InstancePoolName = st.InstancePoolName
	pb.MaxCapacity = st.MaxCapacity
	pb.MinIdleInstances = st.MinIdleInstances
	pb.NodeTypeId = st.NodeTypeId

	var preloadedDockerImagesPb []computepb.DockerImagePb
	for _, item := range st.PreloadedDockerImages {
		itemPb, err := DockerImageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedDockerImagesPb = append(preloadedDockerImagesPb, *itemPb)
		}
	}
	pb.PreloadedDockerImages = preloadedDockerImagesPb
	pb.PreloadedSparkVersions = st.PreloadedSparkVersions
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	statePb, err := InstancePoolStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	statsPb, err := InstancePoolStatsToPb(st.Stats)
	if err != nil {
		return nil, err
	}
	if statsPb != nil {
		pb.Stats = statsPb
	}
	statusPb, err := InstancePoolStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolAndStatsFromPb(pb *computepb.InstancePoolAndStatsPb) (*InstancePoolAndStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAndStats{}
	awsAttributesField, err := InstancePoolAwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := InstancePoolAzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	st.CustomTags = pb.CustomTags
	st.DefaultTags = pb.DefaultTags
	diskSpecField, err := DiskSpecFromPb(pb.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecField != nil {
		st.DiskSpec = diskSpecField
	}
	st.EnableElasticDisk = pb.EnableElasticDisk
	gcpAttributesField, err := InstancePoolGcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}
	st.IdleInstanceAutoterminationMinutes = pb.IdleInstanceAutoterminationMinutes
	st.InstancePoolId = pb.InstancePoolId
	st.InstancePoolName = pb.InstancePoolName
	st.MaxCapacity = pb.MaxCapacity
	st.MinIdleInstances = pb.MinIdleInstances
	st.NodeTypeId = pb.NodeTypeId

	var preloadedDockerImagesField []DockerImage
	for _, itemPb := range pb.PreloadedDockerImages {
		item, err := DockerImageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedDockerImagesField = append(preloadedDockerImagesField, *item)
		}
	}
	st.PreloadedDockerImages = preloadedDockerImagesField
	st.PreloadedSparkVersions = pb.PreloadedSparkVersions
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	stateField, err := InstancePoolStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	statsField, err := InstancePoolStatsFromPb(pb.Stats)
	if err != nil {
		return nil, err
	}
	if statsField != nil {
		st.Stats = statsField
	}
	statusField, err := InstancePoolStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	// Wire name: 'availability'
	Availability InstancePoolAwsAttributesAvailability ``
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
	SpotBidPricePercent int ``
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	// Wire name: 'zone_id'
	ZoneId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolAwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolAwsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolAwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolAwsAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolAwsAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolAwsAttributesToPb(st *InstancePoolAwsAttributes) (*computepb.InstancePoolAwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolAwsAttributesPb{}
	availabilityPb, err := InstancePoolAwsAttributesAvailabilityToPb(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}
	pb.SpotBidPricePercent = st.SpotBidPricePercent
	pb.ZoneId = st.ZoneId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolAwsAttributesFromPb(pb *computepb.InstancePoolAwsAttributesPb) (*InstancePoolAwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAwsAttributes{}
	availabilityField, err := InstancePoolAwsAttributesAvailabilityFromPb(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	st.SpotBidPricePercent = pb.SpotBidPricePercent
	st.ZoneId = pb.ZoneId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func InstancePoolAwsAttributesAvailabilityToPb(st *InstancePoolAwsAttributesAvailability) (*computepb.InstancePoolAwsAttributesAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.InstancePoolAwsAttributesAvailabilityPb(*st)
	return &pb, nil
}

func InstancePoolAwsAttributesAvailabilityFromPb(pb *computepb.InstancePoolAwsAttributesAvailabilityPb) (*InstancePoolAwsAttributesAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolAwsAttributesAvailability(*pb)
	return &st, nil
}

// Attributes set during instance pool creation which are related to Azure.
type InstancePoolAzureAttributes struct {
	// Availability type used for the spot nodes.
	// Wire name: 'availability'
	Availability InstancePoolAzureAttributesAvailability ``
	// With variable pricing, you have option to set a max price, in US dollars
	// (USD) For example, the value 2 would be a max price of $2.00 USD per
	// hour. If you set the max price to be -1, the VM won't be evicted based on
	// price. The price for the VM will be the current price for spot or the
	// price for a standard VM, which ever is less, as long as there is capacity
	// and quota available.
	// Wire name: 'spot_bid_max_price'
	SpotBidMaxPrice float64  ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolAzureAttributes) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolAzureAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolAzureAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolAzureAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolAzureAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolAzureAttributesToPb(st *InstancePoolAzureAttributes) (*computepb.InstancePoolAzureAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolAzureAttributesPb{}
	availabilityPb, err := InstancePoolAzureAttributesAvailabilityToPb(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}
	pb.SpotBidMaxPrice = st.SpotBidMaxPrice

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolAzureAttributesFromPb(pb *computepb.InstancePoolAzureAttributesPb) (*InstancePoolAzureAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAzureAttributes{}
	availabilityField, err := InstancePoolAzureAttributesAvailabilityFromPb(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	st.SpotBidMaxPrice = pb.SpotBidMaxPrice

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func InstancePoolAzureAttributesAvailabilityToPb(st *InstancePoolAzureAttributesAvailability) (*computepb.InstancePoolAzureAttributesAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.InstancePoolAzureAttributesAvailabilityPb(*st)
	return &pb, nil
}

func InstancePoolAzureAttributesAvailabilityFromPb(pb *computepb.InstancePoolAzureAttributesAvailabilityPb) (*InstancePoolAzureAttributesAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolAzureAttributesAvailability(*pb)
	return &st, nil
}

// Attributes set during instance pool creation which are related to GCP.
type InstancePoolGcpAttributes struct {

	// Wire name: 'gcp_availability'
	GcpAvailability GcpAvailability ``
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	// Wire name: 'local_ssd_count'
	LocalSsdCount int ``
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
	ZoneId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolGcpAttributes) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolGcpAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolGcpAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolGcpAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolGcpAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolGcpAttributesToPb(st *InstancePoolGcpAttributes) (*computepb.InstancePoolGcpAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolGcpAttributesPb{}
	gcpAvailabilityPb, err := GcpAvailabilityToPb(&st.GcpAvailability)
	if err != nil {
		return nil, err
	}
	if gcpAvailabilityPb != nil {
		pb.GcpAvailability = *gcpAvailabilityPb
	}
	pb.LocalSsdCount = st.LocalSsdCount
	pb.ZoneId = st.ZoneId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolGcpAttributesFromPb(pb *computepb.InstancePoolGcpAttributesPb) (*InstancePoolGcpAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolGcpAttributes{}
	gcpAvailabilityField, err := GcpAvailabilityFromPb(&pb.GcpAvailability)
	if err != nil {
		return nil, err
	}
	if gcpAvailabilityField != nil {
		st.GcpAvailability = *gcpAvailabilityField
	}
	st.LocalSsdCount = pb.LocalSsdCount
	st.ZoneId = pb.ZoneId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstancePoolPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel InstancePoolPermissionLevel ``
	ForceSendFields []string                    `tf:"-"`
}

func (st InstancePoolPermission) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolPermissionToPb(st *InstancePoolPermission) (*computepb.InstancePoolPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := InstancePoolPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolPermissionFromPb(pb *computepb.InstancePoolPermissionPb) (*InstancePoolPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := InstancePoolPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func InstancePoolPermissionLevelToPb(st *InstancePoolPermissionLevel) (*computepb.InstancePoolPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.InstancePoolPermissionLevelPb(*st)
	return &pb, nil
}

func InstancePoolPermissionLevelFromPb(pb *computepb.InstancePoolPermissionLevelPb) (*InstancePoolPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolPermissionLevel(*pb)
	return &st, nil
}

type InstancePoolPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []InstancePoolAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolPermissions) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolPermissionsToPb(st *InstancePoolPermissions) (*computepb.InstancePoolPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolPermissionsPb{}

	var accessControlListPb []computepb.InstancePoolAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := InstancePoolAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolPermissionsFromPb(pb *computepb.InstancePoolPermissionsPb) (*InstancePoolPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissions{}

	var accessControlListField []InstancePoolAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := InstancePoolAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstancePoolPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel InstancePoolPermissionLevel ``
	ForceSendFields []string                    `tf:"-"`
}

func (st InstancePoolPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolPermissionsDescriptionToPb(st *InstancePoolPermissionsDescription) (*computepb.InstancePoolPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := InstancePoolPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolPermissionsDescriptionFromPb(pb *computepb.InstancePoolPermissionsDescriptionPb) (*InstancePoolPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := InstancePoolPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstancePoolPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []InstancePoolAccessControlRequest ``
	// The instance pool for which to get or manage permissions.
	// Wire name: 'instance_pool_id'
	InstancePoolId string `tf:"-"`
}

func (st InstancePoolPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolPermissionsRequestToPb(st *InstancePoolPermissionsRequest) (*computepb.InstancePoolPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolPermissionsRequestPb{}

	var accessControlListPb []computepb.InstancePoolAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := InstancePoolAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.InstancePoolId = st.InstancePoolId

	return pb, nil
}

func InstancePoolPermissionsRequestFromPb(pb *computepb.InstancePoolPermissionsRequestPb) (*InstancePoolPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissionsRequest{}

	var accessControlListField []InstancePoolAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := InstancePoolAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.InstancePoolId = pb.InstancePoolId

	return st, nil
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

func InstancePoolStateToPb(st *InstancePoolState) (*computepb.InstancePoolStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.InstancePoolStatePb(*st)
	return &pb, nil
}

func InstancePoolStateFromPb(pb *computepb.InstancePoolStatePb) (*InstancePoolState, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolState(*pb)
	return &st, nil
}

type InstancePoolStats struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	// Wire name: 'idle_count'
	IdleCount int ``
	// Number of pending instances in the pool that are NOT part of a cluster.
	// Wire name: 'pending_idle_count'
	PendingIdleCount int ``
	// Number of pending instances in the pool that are part of a cluster.
	// Wire name: 'pending_used_count'
	PendingUsedCount int ``
	// Number of active instances in the pool that are part of a cluster.
	// Wire name: 'used_count'
	UsedCount       int      ``
	ForceSendFields []string `tf:"-"`
}

func (st InstancePoolStats) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolStatsToPb(st *InstancePoolStats) (*computepb.InstancePoolStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolStatsPb{}
	pb.IdleCount = st.IdleCount
	pb.PendingIdleCount = st.PendingIdleCount
	pb.PendingUsedCount = st.PendingUsedCount
	pb.UsedCount = st.UsedCount

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstancePoolStatsFromPb(pb *computepb.InstancePoolStatsPb) (*InstancePoolStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolStats{}
	st.IdleCount = pb.IdleCount
	st.PendingIdleCount = pb.PendingIdleCount
	st.PendingUsedCount = pb.PendingUsedCount
	st.UsedCount = pb.UsedCount

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	// Wire name: 'pending_instance_errors'
	PendingInstanceErrors []PendingInstanceError ``
}

func (st InstancePoolStatus) MarshalJSON() ([]byte, error) {
	pb, err := InstancePoolStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstancePoolStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstancePoolStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstancePoolStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstancePoolStatusToPb(st *InstancePoolStatus) (*computepb.InstancePoolStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstancePoolStatusPb{}

	var pendingInstanceErrorsPb []computepb.PendingInstanceErrorPb
	for _, item := range st.PendingInstanceErrors {
		itemPb, err := PendingInstanceErrorToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pendingInstanceErrorsPb = append(pendingInstanceErrorsPb, *itemPb)
		}
	}
	pb.PendingInstanceErrors = pendingInstanceErrorsPb

	return pb, nil
}

func InstancePoolStatusFromPb(pb *computepb.InstancePoolStatusPb) (*InstancePoolStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolStatus{}

	var pendingInstanceErrorsField []PendingInstanceError
	for _, itemPb := range pb.PendingInstanceErrors {
		item, err := PendingInstanceErrorFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pendingInstanceErrorsField = append(pendingInstanceErrorsField, *item)
		}
	}
	st.PendingInstanceErrors = pendingInstanceErrorsField

	return st, nil
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
	IamRoleArn string ``
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	// Wire name: 'is_meta_instance_profile'
	IsMetaInstanceProfile bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (st InstanceProfile) MarshalJSON() ([]byte, error) {
	pb, err := InstanceProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InstanceProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.InstanceProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InstanceProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InstanceProfileToPb(st *InstanceProfile) (*computepb.InstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.InstanceProfilePb{}
	pb.IamRoleArn = st.IamRoleArn
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.IsMetaInstanceProfile = st.IsMetaInstanceProfile

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func InstanceProfileFromPb(pb *computepb.InstanceProfilePb) (*InstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstanceProfile{}
	st.IamRoleArn = pb.IamRoleArn
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.IsMetaInstanceProfile = pb.IsMetaInstanceProfile

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func KindToPb(st *Kind) (*computepb.KindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.KindPb(*st)
	return &pb, nil
}

func KindFromPb(pb *computepb.KindPb) (*Kind, error) {
	if pb == nil {
		return nil, nil
	}
	st := Kind(*pb)
	return &st, nil
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

func LanguageToPb(st *Language) (*computepb.LanguagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.LanguagePb(*st)
	return &pb, nil
}

func LanguageFromPb(pb *computepb.LanguagePb) (*Language, error) {
	if pb == nil {
		return nil, nil
	}
	st := Language(*pb)
	return &st, nil
}

type Library struct {
	// Specification of a CRAN library to be installed as part of the library
	// Wire name: 'cran'
	Cran *RCranLibrary ``
	// Deprecated. URI of the egg library to install. Installing Python egg
	// files is deprecated and is not supported in Databricks Runtime 14.0 and
	// above.
	// Wire name: 'egg'
	Egg string ``
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	// Wire name: 'jar'
	Jar string ``
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	// Wire name: 'maven'
	Maven *MavenLibrary ``
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	// Wire name: 'pypi'
	Pypi *PythonPyPiLibrary ``
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	// Wire name: 'requirements'
	Requirements string ``
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	// Wire name: 'whl'
	Whl             string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Library) MarshalJSON() ([]byte, error) {
	pb, err := LibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Library) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.LibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LibraryToPb(st *Library) (*computepb.LibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.LibraryPb{}
	cranPb, err := RCranLibraryToPb(st.Cran)
	if err != nil {
		return nil, err
	}
	if cranPb != nil {
		pb.Cran = cranPb
	}
	pb.Egg = st.Egg
	pb.Jar = st.Jar
	mavenPb, err := MavenLibraryToPb(st.Maven)
	if err != nil {
		return nil, err
	}
	if mavenPb != nil {
		pb.Maven = mavenPb
	}
	pypiPb, err := PythonPyPiLibraryToPb(st.Pypi)
	if err != nil {
		return nil, err
	}
	if pypiPb != nil {
		pb.Pypi = pypiPb
	}
	pb.Requirements = st.Requirements
	pb.Whl = st.Whl

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LibraryFromPb(pb *computepb.LibraryPb) (*Library, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Library{}
	cranField, err := RCranLibraryFromPb(pb.Cran)
	if err != nil {
		return nil, err
	}
	if cranField != nil {
		st.Cran = cranField
	}
	st.Egg = pb.Egg
	st.Jar = pb.Jar
	mavenField, err := MavenLibraryFromPb(pb.Maven)
	if err != nil {
		return nil, err
	}
	if mavenField != nil {
		st.Maven = mavenField
	}
	pypiField, err := PythonPyPiLibraryFromPb(pb.Pypi)
	if err != nil {
		return nil, err
	}
	if pypiField != nil {
		st.Pypi = pypiField
	}
	st.Requirements = pb.Requirements
	st.Whl = pb.Whl

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	// Wire name: 'is_library_for_all_clusters'
	IsLibraryForAllClusters bool ``
	// Unique identifier for the library.
	// Wire name: 'library'
	Library *Library ``
	// All the info and warning messages that have occurred so far for this
	// library.
	// Wire name: 'messages'
	Messages []string ``
	// Status of installing the library on the cluster.
	// Wire name: 'status'
	Status          LibraryInstallStatus ``
	ForceSendFields []string             `tf:"-"`
}

func (st LibraryFullStatus) MarshalJSON() ([]byte, error) {
	pb, err := LibraryFullStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LibraryFullStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.LibraryFullStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LibraryFullStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LibraryFullStatusToPb(st *LibraryFullStatus) (*computepb.LibraryFullStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.LibraryFullStatusPb{}
	pb.IsLibraryForAllClusters = st.IsLibraryForAllClusters
	libraryPb, err := LibraryToPb(st.Library)
	if err != nil {
		return nil, err
	}
	if libraryPb != nil {
		pb.Library = libraryPb
	}
	pb.Messages = st.Messages
	statusPb, err := LibraryInstallStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LibraryFullStatusFromPb(pb *computepb.LibraryFullStatusPb) (*LibraryFullStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LibraryFullStatus{}
	st.IsLibraryForAllClusters = pb.IsLibraryForAllClusters
	libraryField, err := LibraryFromPb(pb.Library)
	if err != nil {
		return nil, err
	}
	if libraryField != nil {
		st.Library = libraryField
	}
	st.Messages = pb.Messages
	statusField, err := LibraryInstallStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func LibraryInstallStatusToPb(st *LibraryInstallStatus) (*computepb.LibraryInstallStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.LibraryInstallStatusPb(*st)
	return &pb, nil
}

func LibraryInstallStatusFromPb(pb *computepb.LibraryInstallStatusPb) (*LibraryInstallStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := LibraryInstallStatus(*pb)
	return &st, nil
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	// Wire name: 'statuses'
	Statuses []ClusterLibraryStatuses ``
}

func (st ListAllClusterLibraryStatusesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListAllClusterLibraryStatusesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAllClusterLibraryStatusesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListAllClusterLibraryStatusesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAllClusterLibraryStatusesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAllClusterLibraryStatusesResponseToPb(st *ListAllClusterLibraryStatusesResponse) (*computepb.ListAllClusterLibraryStatusesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListAllClusterLibraryStatusesResponsePb{}

	var statusesPb []computepb.ClusterLibraryStatusesPb
	for _, item := range st.Statuses {
		itemPb, err := ClusterLibraryStatusesToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statusesPb = append(statusesPb, *itemPb)
		}
	}
	pb.Statuses = statusesPb

	return pb, nil
}

func ListAllClusterLibraryStatusesResponseFromPb(pb *computepb.ListAllClusterLibraryStatusesResponsePb) (*ListAllClusterLibraryStatusesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllClusterLibraryStatusesResponse{}

	var statusesField []ClusterLibraryStatuses
	for _, itemPb := range pb.Statuses {
		item, err := ClusterLibraryStatusesFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statusesField = append(statusesField, *item)
		}
	}
	st.Statuses = statusesField

	return st, nil
}

type ListAvailableZonesResponse struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	// Wire name: 'default_zone'
	DefaultZone string ``
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	// Wire name: 'zones'
	Zones           []string ``
	ForceSendFields []string `tf:"-"`
}

func (st ListAvailableZonesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListAvailableZonesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAvailableZonesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListAvailableZonesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAvailableZonesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAvailableZonesResponseToPb(st *ListAvailableZonesResponse) (*computepb.ListAvailableZonesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListAvailableZonesResponsePb{}
	pb.DefaultZone = st.DefaultZone
	pb.Zones = st.Zones

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAvailableZonesResponseFromPb(pb *computepb.ListAvailableZonesResponsePb) (*ListAvailableZonesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAvailableZonesResponse{}
	st.DefaultZone = pb.DefaultZone
	st.Zones = pb.Zones

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListClusterCompliancesRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Canonical unique identifier for the cluster policy.
	// Wire name: 'policy_id'
	PolicyId        string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListClusterCompliancesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListClusterCompliancesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClusterCompliancesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClusterCompliancesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClusterCompliancesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClusterCompliancesRequestToPb(st *ListClusterCompliancesRequest) (*computepb.ListClusterCompliancesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClusterCompliancesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListClusterCompliancesRequestFromPb(pb *computepb.ListClusterCompliancesRequestPb) (*ListClusterCompliancesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterCompliancesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListClusterCompliancesResponse struct {
	// A list of clusters and their policy compliance statuses.
	// Wire name: 'clusters'
	Clusters []ClusterCompliance ``
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	// Wire name: 'prev_page_token'
	PrevPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListClusterCompliancesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListClusterCompliancesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClusterCompliancesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClusterCompliancesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClusterCompliancesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClusterCompliancesResponseToPb(st *ListClusterCompliancesResponse) (*computepb.ListClusterCompliancesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClusterCompliancesResponsePb{}

	var clustersPb []computepb.ClusterCompliancePb
	for _, item := range st.Clusters {
		itemPb, err := ClusterComplianceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListClusterCompliancesResponseFromPb(pb *computepb.ListClusterCompliancesResponsePb) (*ListClusterCompliancesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterCompliancesResponse{}

	var clustersField []ClusterCompliance
	for _, itemPb := range pb.Clusters {
		item, err := ClusterComplianceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	// Wire name: 'sort_column'
	SortColumn ListSortColumn `tf:"-"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	// Wire name: 'sort_order'
	SortOrder ListSortOrder `tf:"-"`
}

func (st ListClusterPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListClusterPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClusterPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClusterPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClusterPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClusterPoliciesRequestToPb(st *ListClusterPoliciesRequest) (*computepb.ListClusterPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClusterPoliciesRequestPb{}
	sortColumnPb, err := ListSortColumnToPb(&st.SortColumn)
	if err != nil {
		return nil, err
	}
	if sortColumnPb != nil {
		pb.SortColumn = *sortColumnPb
	}
	sortOrderPb, err := ListSortOrderToPb(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}

	return pb, nil
}

func ListClusterPoliciesRequestFromPb(pb *computepb.ListClusterPoliciesRequestPb) (*ListClusterPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterPoliciesRequest{}
	sortColumnField, err := ListSortColumnFromPb(&pb.SortColumn)
	if err != nil {
		return nil, err
	}
	if sortColumnField != nil {
		st.SortColumn = *sortColumnField
	}
	sortOrderField, err := ListSortOrderFromPb(&pb.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderField != nil {
		st.SortOrder = *sortOrderField
	}

	return st, nil
}

type ListClustersFilterBy struct {
	// The source of cluster creation.
	// Wire name: 'cluster_sources'
	ClusterSources []ClusterSource ``
	// The current state of the clusters.
	// Wire name: 'cluster_states'
	ClusterStates []State ``
	// Whether the clusters are pinned or not.
	// Wire name: 'is_pinned'
	IsPinned bool ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListClustersFilterBy) MarshalJSON() ([]byte, error) {
	pb, err := ListClustersFilterByToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClustersFilterBy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClustersFilterByPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClustersFilterByFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClustersFilterByToPb(st *ListClustersFilterBy) (*computepb.ListClustersFilterByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClustersFilterByPb{}

	var clusterSourcesPb []computepb.ClusterSourcePb
	for _, item := range st.ClusterSources {
		itemPb, err := ClusterSourceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clusterSourcesPb = append(clusterSourcesPb, *itemPb)
		}
	}
	pb.ClusterSources = clusterSourcesPb

	var clusterStatesPb []computepb.StatePb
	for _, item := range st.ClusterStates {
		itemPb, err := StateToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clusterStatesPb = append(clusterStatesPb, *itemPb)
		}
	}
	pb.ClusterStates = clusterStatesPb
	pb.IsPinned = st.IsPinned
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListClustersFilterByFromPb(pb *computepb.ListClustersFilterByPb) (*ListClustersFilterBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersFilterBy{}

	var clusterSourcesField []ClusterSource
	for _, itemPb := range pb.ClusterSources {
		item, err := ClusterSourceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clusterSourcesField = append(clusterSourcesField, *item)
		}
	}
	st.ClusterSources = clusterSourcesField

	var clusterStatesField []State
	for _, itemPb := range pb.ClusterStates {
		item, err := StateFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clusterStatesField = append(clusterStatesField, *item)
		}
	}
	st.ClusterStates = clusterStatesField
	st.IsPinned = pb.IsPinned
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListClustersRequest struct {
	// Filters to apply to the list of clusters.
	// Wire name: 'filter_by'
	FilterBy *ListClustersFilterBy `tf:"-"`
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Sort the list of clusters by a specific criteria.
	// Wire name: 'sort_by'
	SortBy          *ListClustersSortBy `tf:"-"`
	ForceSendFields []string            `tf:"-"`
}

func (st ListClustersRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListClustersRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClustersRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClustersRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClustersRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClustersRequestToPb(st *ListClustersRequest) (*computepb.ListClustersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClustersRequestPb{}
	filterByPb, err := ListClustersFilterByToPb(st.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByPb != nil {
		pb.FilterBy = filterByPb
	}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	sortByPb, err := ListClustersSortByToPb(st.SortBy)
	if err != nil {
		return nil, err
	}
	if sortByPb != nil {
		pb.SortBy = sortByPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListClustersRequestFromPb(pb *computepb.ListClustersRequestPb) (*ListClustersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersRequest{}
	filterByField, err := ListClustersFilterByFromPb(pb.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByField != nil {
		st.FilterBy = filterByField
	}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	sortByField, err := ListClustersSortByFromPb(pb.SortBy)
	if err != nil {
		return nil, err
	}
	if sortByField != nil {
		st.SortBy = sortByField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListClustersResponse struct {

	// Wire name: 'clusters'
	Clusters []ClusterDetails ``
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	// Wire name: 'prev_page_token'
	PrevPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListClustersResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListClustersResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClustersResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClustersResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClustersResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClustersResponseToPb(st *ListClustersResponse) (*computepb.ListClustersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClustersResponsePb{}

	var clustersPb []computepb.ClusterDetailsPb
	for _, item := range st.Clusters {
		itemPb, err := ClusterDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListClustersResponseFromPb(pb *computepb.ListClustersResponsePb) (*ListClustersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersResponse{}

	var clustersField []ClusterDetails
	for _, itemPb := range pb.Clusters {
		item, err := ClusterDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListClustersSortBy struct {
	// The direction to sort by.
	// Wire name: 'direction'
	Direction ListClustersSortByDirection ``
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	// Wire name: 'field'
	Field ListClustersSortByField ``
}

func (st ListClustersSortBy) MarshalJSON() ([]byte, error) {
	pb, err := ListClustersSortByToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListClustersSortBy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListClustersSortByPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListClustersSortByFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListClustersSortByToPb(st *ListClustersSortBy) (*computepb.ListClustersSortByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListClustersSortByPb{}
	directionPb, err := ListClustersSortByDirectionToPb(&st.Direction)
	if err != nil {
		return nil, err
	}
	if directionPb != nil {
		pb.Direction = *directionPb
	}
	fieldPb, err := ListClustersSortByFieldToPb(&st.Field)
	if err != nil {
		return nil, err
	}
	if fieldPb != nil {
		pb.Field = *fieldPb
	}

	return pb, nil
}

func ListClustersSortByFromPb(pb *computepb.ListClustersSortByPb) (*ListClustersSortBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersSortBy{}
	directionField, err := ListClustersSortByDirectionFromPb(&pb.Direction)
	if err != nil {
		return nil, err
	}
	if directionField != nil {
		st.Direction = *directionField
	}
	fieldField, err := ListClustersSortByFieldFromPb(&pb.Field)
	if err != nil {
		return nil, err
	}
	if fieldField != nil {
		st.Field = *fieldField
	}

	return st, nil
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

func ListClustersSortByDirectionToPb(st *ListClustersSortByDirection) (*computepb.ListClustersSortByDirectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ListClustersSortByDirectionPb(*st)
	return &pb, nil
}

func ListClustersSortByDirectionFromPb(pb *computepb.ListClustersSortByDirectionPb) (*ListClustersSortByDirection, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListClustersSortByDirection(*pb)
	return &st, nil
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

func ListClustersSortByFieldToPb(st *ListClustersSortByField) (*computepb.ListClustersSortByFieldPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ListClustersSortByFieldPb(*st)
	return &pb, nil
}

func ListClustersSortByFieldFromPb(pb *computepb.ListClustersSortByFieldPb) (*ListClustersSortByField, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListClustersSortByField(*pb)
	return &st, nil
}

type ListGlobalInitScriptsResponse struct {

	// Wire name: 'scripts'
	Scripts []GlobalInitScriptDetails ``
}

func (st ListGlobalInitScriptsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListGlobalInitScriptsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListGlobalInitScriptsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListGlobalInitScriptsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListGlobalInitScriptsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListGlobalInitScriptsResponseToPb(st *ListGlobalInitScriptsResponse) (*computepb.ListGlobalInitScriptsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListGlobalInitScriptsResponsePb{}

	var scriptsPb []computepb.GlobalInitScriptDetailsPb
	for _, item := range st.Scripts {
		itemPb, err := GlobalInitScriptDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			scriptsPb = append(scriptsPb, *itemPb)
		}
	}
	pb.Scripts = scriptsPb

	return pb, nil
}

func ListGlobalInitScriptsResponseFromPb(pb *computepb.ListGlobalInitScriptsResponsePb) (*ListGlobalInitScriptsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGlobalInitScriptsResponse{}

	var scriptsField []GlobalInitScriptDetails
	for _, itemPb := range pb.Scripts {
		item, err := GlobalInitScriptDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			scriptsField = append(scriptsField, *item)
		}
	}
	st.Scripts = scriptsField

	return st, nil
}

type ListInstancePools struct {

	// Wire name: 'instance_pools'
	InstancePools []InstancePoolAndStats ``
}

func (st ListInstancePools) MarshalJSON() ([]byte, error) {
	pb, err := ListInstancePoolsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListInstancePools) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListInstancePoolsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListInstancePoolsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListInstancePoolsToPb(st *ListInstancePools) (*computepb.ListInstancePoolsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListInstancePoolsPb{}

	var instancePoolsPb []computepb.InstancePoolAndStatsPb
	for _, item := range st.InstancePools {
		itemPb, err := InstancePoolAndStatsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			instancePoolsPb = append(instancePoolsPb, *itemPb)
		}
	}
	pb.InstancePools = instancePoolsPb

	return pb, nil
}

func ListInstancePoolsFromPb(pb *computepb.ListInstancePoolsPb) (*ListInstancePools, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstancePools{}

	var instancePoolsField []InstancePoolAndStats
	for _, itemPb := range pb.InstancePools {
		item, err := InstancePoolAndStatsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			instancePoolsField = append(instancePoolsField, *item)
		}
	}
	st.InstancePools = instancePoolsField

	return st, nil
}

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	// Wire name: 'instance_profiles'
	InstanceProfiles []InstanceProfile ``
}

func (st ListInstanceProfilesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListInstanceProfilesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListInstanceProfilesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListInstanceProfilesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListInstanceProfilesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListInstanceProfilesResponseToPb(st *ListInstanceProfilesResponse) (*computepb.ListInstanceProfilesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListInstanceProfilesResponsePb{}

	var instanceProfilesPb []computepb.InstanceProfilePb
	for _, item := range st.InstanceProfiles {
		itemPb, err := InstanceProfileToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			instanceProfilesPb = append(instanceProfilesPb, *itemPb)
		}
	}
	pb.InstanceProfiles = instanceProfilesPb

	return pb, nil
}

func ListInstanceProfilesResponseFromPb(pb *computepb.ListInstanceProfilesResponsePb) (*ListInstanceProfilesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstanceProfilesResponse{}

	var instanceProfilesField []InstanceProfile
	for _, itemPb := range pb.InstanceProfiles {
		item, err := InstanceProfileFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			instanceProfilesField = append(instanceProfilesField, *item)
		}
	}
	st.InstanceProfiles = instanceProfilesField

	return st, nil
}

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	// Wire name: 'node_types'
	NodeTypes []NodeType ``
}

func (st ListNodeTypesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListNodeTypesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNodeTypesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListNodeTypesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNodeTypesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNodeTypesResponseToPb(st *ListNodeTypesResponse) (*computepb.ListNodeTypesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListNodeTypesResponsePb{}

	var nodeTypesPb []computepb.NodeTypePb
	for _, item := range st.NodeTypes {
		itemPb, err := NodeTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			nodeTypesPb = append(nodeTypesPb, *itemPb)
		}
	}
	pb.NodeTypes = nodeTypesPb

	return pb, nil
}

func ListNodeTypesResponseFromPb(pb *computepb.ListNodeTypesResponsePb) (*ListNodeTypesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNodeTypesResponse{}

	var nodeTypesField []NodeType
	for _, itemPb := range pb.NodeTypes {
		item, err := NodeTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			nodeTypesField = append(nodeTypesField, *item)
		}
	}
	st.NodeTypes = nodeTypesField

	return st, nil
}

type ListPoliciesResponse struct {
	// List of policies.
	// Wire name: 'policies'
	Policies []Policy ``
}

func (st ListPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPoliciesResponseToPb(st *ListPoliciesResponse) (*computepb.ListPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListPoliciesResponsePb{}

	var policiesPb []computepb.PolicyPb
	for _, item := range st.Policies {
		itemPb, err := PolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policiesPb = append(policiesPb, *itemPb)
		}
	}
	pb.Policies = policiesPb

	return pb, nil
}

func ListPoliciesResponseFromPb(pb *computepb.ListPoliciesResponsePb) (*ListPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPoliciesResponse{}

	var policiesField []Policy
	for _, itemPb := range pb.Policies {
		item, err := PolicyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policiesField = append(policiesField, *item)
		}
	}
	st.Policies = policiesField

	return st, nil
}

type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// A token that can be used to get the next page of results.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListPolicyFamiliesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListPolicyFamiliesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPolicyFamiliesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListPolicyFamiliesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPolicyFamiliesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPolicyFamiliesRequestToPb(st *ListPolicyFamiliesRequest) (*computepb.ListPolicyFamiliesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListPolicyFamiliesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListPolicyFamiliesRequestFromPb(pb *computepb.ListPolicyFamiliesRequestPb) (*ListPolicyFamiliesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPolicyFamiliesRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// List of policy families.
	// Wire name: 'policy_families'
	PolicyFamilies  []PolicyFamily ``
	ForceSendFields []string       `tf:"-"`
}

func (st ListPolicyFamiliesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListPolicyFamiliesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPolicyFamiliesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ListPolicyFamiliesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPolicyFamiliesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPolicyFamiliesResponseToPb(st *ListPolicyFamiliesResponse) (*computepb.ListPolicyFamiliesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ListPolicyFamiliesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var policyFamiliesPb []computepb.PolicyFamilyPb
	for _, item := range st.PolicyFamilies {
		itemPb, err := PolicyFamilyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policyFamiliesPb = append(policyFamiliesPb, *itemPb)
		}
	}
	pb.PolicyFamilies = policyFamiliesPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListPolicyFamiliesResponseFromPb(pb *computepb.ListPolicyFamiliesResponsePb) (*ListPolicyFamiliesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPolicyFamiliesResponse{}
	st.NextPageToken = pb.NextPageToken

	var policyFamiliesField []PolicyFamily
	for _, itemPb := range pb.PolicyFamilies {
		item, err := PolicyFamilyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policyFamiliesField = append(policyFamiliesField, *item)
		}
	}
	st.PolicyFamilies = policyFamiliesField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ListSortColumnToPb(st *ListSortColumn) (*computepb.ListSortColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ListSortColumnPb(*st)
	return &pb, nil
}

func ListSortColumnFromPb(pb *computepb.ListSortColumnPb) (*ListSortColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListSortColumn(*pb)
	return &st, nil
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

func ListSortOrderToPb(st *ListSortOrder) (*computepb.ListSortOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ListSortOrderPb(*st)
	return &pb, nil
}

func ListSortOrderFromPb(pb *computepb.ListSortOrderPb) (*ListSortOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListSortOrder(*pb)
	return &st, nil
}

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	// Wire name: 'destination'
	Destination string ``
}

func (st LocalFileInfo) MarshalJSON() ([]byte, error) {
	pb, err := LocalFileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LocalFileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.LocalFileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LocalFileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LocalFileInfoToPb(st *LocalFileInfo) (*computepb.LocalFileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.LocalFileInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

func LocalFileInfoFromPb(pb *computepb.LocalFileInfoPb) (*LocalFileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LocalFileInfo{}
	st.Destination = pb.Destination

	return st, nil
}

type LogAnalyticsInfo struct {

	// Wire name: 'log_analytics_primary_key'
	LogAnalyticsPrimaryKey string ``

	// Wire name: 'log_analytics_workspace_id'
	LogAnalyticsWorkspaceId string   ``
	ForceSendFields         []string `tf:"-"`
}

func (st LogAnalyticsInfo) MarshalJSON() ([]byte, error) {
	pb, err := LogAnalyticsInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogAnalyticsInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.LogAnalyticsInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogAnalyticsInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogAnalyticsInfoToPb(st *LogAnalyticsInfo) (*computepb.LogAnalyticsInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.LogAnalyticsInfoPb{}
	pb.LogAnalyticsPrimaryKey = st.LogAnalyticsPrimaryKey
	pb.LogAnalyticsWorkspaceId = st.LogAnalyticsWorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogAnalyticsInfoFromPb(pb *computepb.LogAnalyticsInfoPb) (*LogAnalyticsInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogAnalyticsInfo{}
	st.LogAnalyticsPrimaryKey = pb.LogAnalyticsPrimaryKey
	st.LogAnalyticsWorkspaceId = pb.LogAnalyticsWorkspaceId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The log delivery status
type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	// Wire name: 'last_attempted'
	LastAttempted int64 ``
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	// Wire name: 'last_exception'
	LastException   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st LogSyncStatus) MarshalJSON() ([]byte, error) {
	pb, err := LogSyncStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogSyncStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.LogSyncStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogSyncStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogSyncStatusToPb(st *LogSyncStatus) (*computepb.LogSyncStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.LogSyncStatusPb{}
	pb.LastAttempted = st.LastAttempted
	pb.LastException = st.LastException

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogSyncStatusFromPb(pb *computepb.LogSyncStatusPb) (*LogSyncStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogSyncStatus{}
	st.LastAttempted = pb.LastAttempted
	st.LastException = pb.LastException

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type MapAny map[string]any

func MapAnyToPb(st *MapAny) (*computepb.MapAnyPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := computepb.MapAnyPb(*st)
	return &stPb, nil
}

func MapAnyFromPb(stPb *computepb.MapAnyPb) (*MapAny, error) {
	if stPb == nil {
		return nil, nil
	}
	st := MapAny(*stPb)
	return &st, nil
}

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	// Wire name: 'coordinates'
	Coordinates string ``
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	// Wire name: 'exclusions'
	Exclusions []string ``
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	// Wire name: 'repo'
	Repo            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st MavenLibrary) MarshalJSON() ([]byte, error) {
	pb, err := MavenLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *MavenLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.MavenLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MavenLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MavenLibraryToPb(st *MavenLibrary) (*computepb.MavenLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.MavenLibraryPb{}
	pb.Coordinates = st.Coordinates
	pb.Exclusions = st.Exclusions
	pb.Repo = st.Repo

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func MavenLibraryFromPb(pb *computepb.MavenLibraryPb) (*MavenLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MavenLibrary{}
	st.Coordinates = pb.Coordinates
	st.Exclusions = pb.Exclusions
	st.Repo = pb.Repo

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// This structure embodies the machine type that hosts spark containers Note:
// this should be an internal data structure for now It is defined in proto in
// case we want to send it over the wire in the future (which is likely)
type NodeInstanceType struct {
	// Unique identifier across instance types
	// Wire name: 'instance_type_id'
	InstanceTypeId string ``
	// Size of the individual local disks attached to this instance (i.e. per
	// local disk).
	// Wire name: 'local_disk_size_gb'
	LocalDiskSizeGb int ``
	// Number of local disks that are present on this instance.
	// Wire name: 'local_disks'
	LocalDisks int ``
	// Size of the individual local nvme disks attached to this instance (i.e.
	// per local disk).
	// Wire name: 'local_nvme_disk_size_gb'
	LocalNvmeDiskSizeGb int ``
	// Number of local nvme disks that are present on this instance.
	// Wire name: 'local_nvme_disks'
	LocalNvmeDisks  int      ``
	ForceSendFields []string `tf:"-"`
}

func (st NodeInstanceType) MarshalJSON() ([]byte, error) {
	pb, err := NodeInstanceTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NodeInstanceType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.NodeInstanceTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NodeInstanceTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NodeInstanceTypeToPb(st *NodeInstanceType) (*computepb.NodeInstanceTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.NodeInstanceTypePb{}
	pb.InstanceTypeId = st.InstanceTypeId
	pb.LocalDiskSizeGb = st.LocalDiskSizeGb
	pb.LocalDisks = st.LocalDisks
	pb.LocalNvmeDiskSizeGb = st.LocalNvmeDiskSizeGb
	pb.LocalNvmeDisks = st.LocalNvmeDisks

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NodeInstanceTypeFromPb(pb *computepb.NodeInstanceTypePb) (*NodeInstanceType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NodeInstanceType{}
	st.InstanceTypeId = pb.InstanceTypeId
	st.LocalDiskSizeGb = pb.LocalDiskSizeGb
	st.LocalDisks = pb.LocalDisks
	st.LocalNvmeDiskSizeGb = pb.LocalNvmeDiskSizeGb
	st.LocalNvmeDisks = pb.LocalNvmeDisks

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A description of a Spark node type including both the dimensions of the node
// and the instance type on which it will be hosted.
type NodeType struct {
	// A descriptive category for this node type. Examples include "Memory
	// Optimized" and "Compute Optimized".
	// Wire name: 'category'
	Category string ``
	// A string description associated with this node type, e.g., "r3.xlarge".
	// Wire name: 'description'
	Description string ``
	// An optional hint at the display order of node types in the UI. Within a
	// node type category, lowest numbers come first.
	// Wire name: 'display_order'
	DisplayOrder int ``
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	// Wire name: 'instance_type_id'
	InstanceTypeId string ``
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	// Wire name: 'is_deprecated'
	IsDeprecated bool ``
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	// Wire name: 'is_encrypted_in_transit'
	IsEncryptedInTransit bool ``
	// Whether this is an Arm-based instance.
	// Wire name: 'is_graviton'
	IsGraviton bool ``
	// Whether this node is hidden from presentation in the UI.
	// Wire name: 'is_hidden'
	IsHidden bool ``
	// Whether this node comes with IO cache enabled by default.
	// Wire name: 'is_io_cache_enabled'
	IsIoCacheEnabled bool ``
	// Memory (in MB) available for this node type.
	// Wire name: 'memory_mb'
	MemoryMb int ``
	// A collection of node type info reported by the cloud provider
	// Wire name: 'node_info'
	NodeInfo *CloudProviderNodeInfo ``
	// The NodeInstanceType object corresponding to instance_type_id
	// Wire name: 'node_instance_type'
	NodeInstanceType *NodeInstanceType ``
	// Unique identifier for this node type.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	// Wire name: 'num_cores'
	NumCores float64 ``
	// Number of GPUs available for this node type.
	// Wire name: 'num_gpus'
	NumGpus int ``

	// Wire name: 'photon_driver_capable'
	PhotonDriverCapable bool ``

	// Wire name: 'photon_worker_capable'
	PhotonWorkerCapable bool ``
	// Whether this node type support cluster tags.
	// Wire name: 'support_cluster_tags'
	SupportClusterTags bool ``
	// Whether this node type support EBS volumes. EBS volumes is disabled for
	// node types that we could place multiple corresponding containers on the
	// same hosting instance.
	// Wire name: 'support_ebs_volumes'
	SupportEbsVolumes bool ``
	// Whether this node type supports port forwarding.
	// Wire name: 'support_port_forwarding'
	SupportPortForwarding bool     ``
	ForceSendFields       []string `tf:"-"`
}

func (st NodeType) MarshalJSON() ([]byte, error) {
	pb, err := NodeTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NodeType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.NodeTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NodeTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NodeTypeToPb(st *NodeType) (*computepb.NodeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.NodeTypePb{}
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
	nodeInfoPb, err := CloudProviderNodeInfoToPb(st.NodeInfo)
	if err != nil {
		return nil, err
	}
	if nodeInfoPb != nil {
		pb.NodeInfo = nodeInfoPb
	}
	nodeInstanceTypePb, err := NodeInstanceTypeToPb(st.NodeInstanceType)
	if err != nil {
		return nil, err
	}
	if nodeInstanceTypePb != nil {
		pb.NodeInstanceType = nodeInstanceTypePb
	}
	pb.NodeTypeId = st.NodeTypeId
	pb.NumCores = st.NumCores
	pb.NumGpus = st.NumGpus
	pb.PhotonDriverCapable = st.PhotonDriverCapable
	pb.PhotonWorkerCapable = st.PhotonWorkerCapable
	pb.SupportClusterTags = st.SupportClusterTags
	pb.SupportEbsVolumes = st.SupportEbsVolumes
	pb.SupportPortForwarding = st.SupportPortForwarding

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NodeTypeFromPb(pb *computepb.NodeTypePb) (*NodeType, error) {
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
	nodeInfoField, err := CloudProviderNodeInfoFromPb(pb.NodeInfo)
	if err != nil {
		return nil, err
	}
	if nodeInfoField != nil {
		st.NodeInfo = nodeInfoField
	}
	nodeInstanceTypeField, err := NodeInstanceTypeFromPb(pb.NodeInstanceType)
	if err != nil {
		return nil, err
	}
	if nodeInstanceTypeField != nil {
		st.NodeInstanceType = nodeInstanceTypeField
	}
	st.NodeTypeId = pb.NodeTypeId
	st.NumCores = pb.NumCores
	st.NumGpus = pb.NumGpus
	st.PhotonDriverCapable = pb.PhotonDriverCapable
	st.PhotonWorkerCapable = pb.PhotonWorkerCapable
	st.SupportClusterTags = pb.SupportClusterTags
	st.SupportEbsVolumes = pb.SupportEbsVolumes
	st.SupportPortForwarding = pb.SupportPortForwarding

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Error message of a failed pending instances
type PendingInstanceError struct {

	// Wire name: 'instance_id'
	InstanceId string ``

	// Wire name: 'message'
	Message         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PendingInstanceError) MarshalJSON() ([]byte, error) {
	pb, err := PendingInstanceErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PendingInstanceError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.PendingInstanceErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PendingInstanceErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PendingInstanceErrorToPb(st *PendingInstanceError) (*computepb.PendingInstanceErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.PendingInstanceErrorPb{}
	pb.InstanceId = st.InstanceId
	pb.Message = st.Message

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PendingInstanceErrorFromPb(pb *computepb.PendingInstanceErrorPb) (*PendingInstanceError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PendingInstanceError{}
	st.InstanceId = pb.InstanceId
	st.Message = pb.Message

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PermanentDeleteCluster struct {
	// The cluster to be deleted.
	// Wire name: 'cluster_id'
	ClusterId string ``
}

func (st PermanentDeleteCluster) MarshalJSON() ([]byte, error) {
	pb, err := PermanentDeleteClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PermanentDeleteCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.PermanentDeleteClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PermanentDeleteClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PermanentDeleteClusterToPb(st *PermanentDeleteCluster) (*computepb.PermanentDeleteClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.PermanentDeleteClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func PermanentDeleteClusterFromPb(pb *computepb.PermanentDeleteClusterPb) (*PermanentDeleteCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermanentDeleteCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type PinCluster struct {

	// Wire name: 'cluster_id'
	ClusterId string ``
}

func (st PinCluster) MarshalJSON() ([]byte, error) {
	pb, err := PinClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PinCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.PinClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PinClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PinClusterToPb(st *PinCluster) (*computepb.PinClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.PinClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func PinClusterFromPb(pb *computepb.PinClusterPb) (*PinCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PinCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

// Describes a Cluster Policy entity.
type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	// Wire name: 'created_at_timestamp'
	CreatedAtTimestamp int64 ``
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	// Wire name: 'creator_user_name'
	CreatorUserName string ``
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string ``
	// Additional human-readable description of the cluster policy.
	// Wire name: 'description'
	Description string ``
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	// Wire name: 'is_default'
	IsDefault bool ``
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	// Wire name: 'libraries'
	Libraries []Library ``
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	// Wire name: 'max_clusters_per_user'
	MaxClustersPerUser int64 ``
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	// Wire name: 'name'
	Name string ``
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
	PolicyFamilyDefinitionOverrides string ``
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	// Wire name: 'policy_family_id'
	PolicyFamilyId string ``
	// Canonical unique identifier for the Cluster Policy.
	// Wire name: 'policy_id'
	PolicyId        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Policy) MarshalJSON() ([]byte, error) {
	pb, err := PolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Policy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.PolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PolicyToPb(st *Policy) (*computepb.PolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.PolicyPb{}
	pb.CreatedAtTimestamp = st.CreatedAtTimestamp
	pb.CreatorUserName = st.CreatorUserName
	pb.Definition = st.Definition
	pb.Description = st.Description
	pb.IsDefault = st.IsDefault

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb
	pb.MaxClustersPerUser = st.MaxClustersPerUser
	pb.Name = st.Name
	pb.PolicyFamilyDefinitionOverrides = st.PolicyFamilyDefinitionOverrides
	pb.PolicyFamilyId = st.PolicyFamilyId
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PolicyFromPb(pb *computepb.PolicyPb) (*Policy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Policy{}
	st.CreatedAtTimestamp = pb.CreatedAtTimestamp
	st.CreatorUserName = pb.CreatorUserName
	st.Definition = pb.Definition
	st.Description = pb.Description
	st.IsDefault = pb.IsDefault

	var librariesField []Library
	for _, itemPb := range pb.Libraries {
		item, err := LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	st.MaxClustersPerUser = pb.MaxClustersPerUser
	st.Name = pb.Name
	st.PolicyFamilyDefinitionOverrides = pb.PolicyFamilyDefinitionOverrides
	st.PolicyFamilyId = pb.PolicyFamilyId
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	// Wire name: 'definition'
	Definition string ``
	// Human-readable description of the purpose of the policy family.
	// Wire name: 'description'
	Description string ``
	// Name of the policy family.
	// Wire name: 'name'
	Name string ``
	// Unique identifier for the policy family.
	// Wire name: 'policy_family_id'
	PolicyFamilyId  string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PolicyFamily) MarshalJSON() ([]byte, error) {
	pb, err := PolicyFamilyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PolicyFamily) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.PolicyFamilyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PolicyFamilyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PolicyFamilyToPb(st *PolicyFamily) (*computepb.PolicyFamilyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.PolicyFamilyPb{}
	pb.Definition = st.Definition
	pb.Description = st.Description
	pb.Name = st.Name
	pb.PolicyFamilyId = st.PolicyFamilyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PolicyFamilyFromPb(pb *computepb.PolicyFamilyPb) (*PolicyFamily, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PolicyFamily{}
	st.Definition = pb.Definition
	st.Description = pb.Description
	st.Name = pb.Name
	st.PolicyFamilyId = pb.PolicyFamilyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	// Wire name: 'package'
	Package string ``
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	// Wire name: 'repo'
	Repo            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PythonPyPiLibrary) MarshalJSON() ([]byte, error) {
	pb, err := PythonPyPiLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PythonPyPiLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.PythonPyPiLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PythonPyPiLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PythonPyPiLibraryToPb(st *PythonPyPiLibrary) (*computepb.PythonPyPiLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.PythonPyPiLibraryPb{}
	pb.Package = st.Package
	pb.Repo = st.Repo

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PythonPyPiLibraryFromPb(pb *computepb.PythonPyPiLibraryPb) (*PythonPyPiLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PythonPyPiLibrary{}
	st.Package = pb.Package
	st.Repo = pb.Repo

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	// Wire name: 'package'
	Package string ``
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	// Wire name: 'repo'
	Repo            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RCranLibrary) MarshalJSON() ([]byte, error) {
	pb, err := RCranLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RCranLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.RCranLibraryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RCranLibraryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RCranLibraryToPb(st *RCranLibrary) (*computepb.RCranLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.RCranLibraryPb{}
	pb.Package = st.Package
	pb.Repo = st.Repo

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RCranLibraryFromPb(pb *computepb.RCranLibraryPb) (*RCranLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RCranLibrary{}
	st.Package = pb.Package
	st.Repo = pb.Repo

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
}

func (st RemoveInstanceProfile) MarshalJSON() ([]byte, error) {
	pb, err := RemoveInstanceProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RemoveInstanceProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.RemoveInstanceProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RemoveInstanceProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RemoveInstanceProfileToPb(st *RemoveInstanceProfile) (*computepb.RemoveInstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.RemoveInstanceProfilePb{}
	pb.InstanceProfileArn = st.InstanceProfileArn

	return pb, nil
}

func RemoveInstanceProfileFromPb(pb *computepb.RemoveInstanceProfilePb) (*RemoveInstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveInstanceProfile{}
	st.InstanceProfileArn = pb.InstanceProfileArn

	return st, nil
}

type ResizeCluster struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
	// The cluster to be resized.
	// Wire name: 'cluster_id'
	ClusterId string ``
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
	NumWorkers      int      ``
	ForceSendFields []string `tf:"-"`
}

func (st ResizeCluster) MarshalJSON() ([]byte, error) {
	pb, err := ResizeClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ResizeCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ResizeClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResizeClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResizeClusterToPb(st *ResizeCluster) (*computepb.ResizeClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ResizeClusterPb{}
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.ClusterId = st.ClusterId
	pb.NumWorkers = st.NumWorkers

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ResizeClusterFromPb(pb *computepb.ResizeClusterPb) (*ResizeCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResizeCluster{}
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.ClusterId = pb.ClusterId
	st.NumWorkers = pb.NumWorkers

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RestartCluster struct {
	// The cluster to be started.
	// Wire name: 'cluster_id'
	ClusterId string ``

	// Wire name: 'restart_user'
	RestartUser     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RestartCluster) MarshalJSON() ([]byte, error) {
	pb, err := RestartClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestartCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.RestartClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestartClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestartClusterToPb(st *RestartCluster) (*computepb.RestartClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.RestartClusterPb{}
	pb.ClusterId = st.ClusterId
	pb.RestartUser = st.RestartUser

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RestartClusterFromPb(pb *computepb.RestartClusterPb) (*RestartCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartCluster{}
	st.ClusterId = pb.ClusterId
	st.RestartUser = pb.RestartUser

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func ResultTypeToPb(st *ResultType) (*computepb.ResultTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.ResultTypePb(*st)
	return &pb, nil
}

func ResultTypeFromPb(pb *computepb.ResultTypePb) (*ResultType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ResultType(*pb)
	return &st, nil
}

type Results struct {
	// The cause of the error
	// Wire name: 'cause'
	Cause string ``

	// Wire name: 'data'
	Data any ``
	// The image filename
	// Wire name: 'fileName'
	FileName string ``

	// Wire name: 'fileNames'
	FileNames []string ``
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	// Wire name: 'isJsonSchema'
	IsJsonSchema bool ``
	// internal field used by SDK
	// Wire name: 'pos'
	Pos int ``

	// Wire name: 'resultType'
	ResultType ResultType ``
	// The table schema
	// Wire name: 'schema'
	Schema []map[string]any ``
	// The summary of the error
	// Wire name: 'summary'
	Summary string ``
	// true if partial results are returned.
	// Wire name: 'truncated'
	Truncated       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st Results) MarshalJSON() ([]byte, error) {
	pb, err := ResultsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Results) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.ResultsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ResultsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ResultsToPb(st *Results) (*computepb.ResultsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.ResultsPb{}
	pb.Cause = st.Cause
	pb.Data = st.Data
	pb.FileName = st.FileName
	pb.FileNames = st.FileNames
	pb.IsJsonSchema = st.IsJsonSchema
	pb.Pos = st.Pos
	resultTypePb, err := ResultTypeToPb(&st.ResultType)
	if err != nil {
		return nil, err
	}
	if resultTypePb != nil {
		pb.ResultType = *resultTypePb
	}
	pb.Schema = st.Schema
	pb.Summary = st.Summary
	pb.Truncated = st.Truncated

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ResultsFromPb(pb *computepb.ResultsPb) (*Results, error) {
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
	resultTypeField, err := ResultTypeFromPb(&pb.ResultType)
	if err != nil {
		return nil, err
	}
	if resultTypeField != nil {
		st.ResultType = *resultTypeField
	}
	st.Schema = pb.Schema
	st.Summary = pb.Summary
	st.Truncated = pb.Truncated

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func RuntimeEngineToPb(st *RuntimeEngine) (*computepb.RuntimeEnginePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.RuntimeEnginePb(*st)
	return &pb, nil
}

func RuntimeEngineFromPb(pb *computepb.RuntimeEnginePb) (*RuntimeEngine, error) {
	if pb == nil {
		return nil, nil
	}
	st := RuntimeEngine(*pb)
	return &st, nil
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
	CannedAcl string ``
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	// Wire name: 'destination'
	Destination string ``
	// (Optional) Flag to enable server side encryption, `false` by default.
	// Wire name: 'enable_encryption'
	EnableEncryption bool ``
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	// Wire name: 'encryption_type'
	EncryptionType string ``
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	// Wire name: 'endpoint'
	Endpoint string ``
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	// Wire name: 'kms_key'
	KmsKey string ``
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	// Wire name: 'region'
	Region          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st S3StorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := S3StorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *S3StorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.S3StorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := S3StorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func S3StorageInfoToPb(st *S3StorageInfo) (*computepb.S3StorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.S3StorageInfoPb{}
	pb.CannedAcl = st.CannedAcl
	pb.Destination = st.Destination
	pb.EnableEncryption = st.EnableEncryption
	pb.EncryptionType = st.EncryptionType
	pb.Endpoint = st.Endpoint
	pb.KmsKey = st.KmsKey
	pb.Region = st.Region

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func S3StorageInfoFromPb(pb *computepb.S3StorageInfoPb) (*S3StorageInfo, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Describes a specific Spark driver or executor.
type SparkNode struct {
	// The private IP address of the host instance.
	// Wire name: 'host_private_ip'
	HostPrivateIp string ``
	// Globally unique identifier for the host instance from the cloud provider.
	// Wire name: 'instance_id'
	InstanceId string ``
	// Attributes specific to AWS for a Spark node.
	// Wire name: 'node_aws_attributes'
	NodeAwsAttributes *SparkNodeAwsAttributes ``
	// Globally unique identifier for this node.
	// Wire name: 'node_id'
	NodeId string ``
	// Private IP address (typically a 10.x.x.x address) of the Spark node. Note
	// that this is different from the private IP address of the host instance.
	// Wire name: 'private_ip'
	PrivateIp string ``
	// Public DNS address of this node. This address can be used to access the
	// Spark JDBC server on the driver node. To communicate with the JDBC
	// server, traffic must be manually authorized by adding security group
	// rules to the "worker-unmanaged" security group via the AWS console.
	// Wire name: 'public_dns'
	PublicDns string ``
	// The timestamp (in millisecond) when the Spark node is launched.
	// Wire name: 'start_timestamp'
	StartTimestamp  int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st SparkNode) MarshalJSON() ([]byte, error) {
	pb, err := SparkNodeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparkNode) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.SparkNodePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparkNodeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SparkNodeToPb(st *SparkNode) (*computepb.SparkNodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.SparkNodePb{}
	pb.HostPrivateIp = st.HostPrivateIp
	pb.InstanceId = st.InstanceId
	nodeAwsAttributesPb, err := SparkNodeAwsAttributesToPb(st.NodeAwsAttributes)
	if err != nil {
		return nil, err
	}
	if nodeAwsAttributesPb != nil {
		pb.NodeAwsAttributes = nodeAwsAttributesPb
	}
	pb.NodeId = st.NodeId
	pb.PrivateIp = st.PrivateIp
	pb.PublicDns = st.PublicDns
	pb.StartTimestamp = st.StartTimestamp

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SparkNodeFromPb(pb *computepb.SparkNodePb) (*SparkNode, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkNode{}
	st.HostPrivateIp = pb.HostPrivateIp
	st.InstanceId = pb.InstanceId
	nodeAwsAttributesField, err := SparkNodeAwsAttributesFromPb(pb.NodeAwsAttributes)
	if err != nil {
		return nil, err
	}
	if nodeAwsAttributesField != nil {
		st.NodeAwsAttributes = nodeAwsAttributesField
	}
	st.NodeId = pb.NodeId
	st.PrivateIp = pb.PrivateIp
	st.PublicDns = pb.PublicDns
	st.StartTimestamp = pb.StartTimestamp

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	// Wire name: 'is_spot'
	IsSpot          bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st SparkNodeAwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := SparkNodeAwsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparkNodeAwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.SparkNodeAwsAttributesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparkNodeAwsAttributesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SparkNodeAwsAttributesToPb(st *SparkNodeAwsAttributes) (*computepb.SparkNodeAwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.SparkNodeAwsAttributesPb{}
	pb.IsSpot = st.IsSpot

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SparkNodeAwsAttributesFromPb(pb *computepb.SparkNodeAwsAttributesPb) (*SparkNodeAwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkNodeAwsAttributes{}
	st.IsSpot = pb.IsSpot

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SparkVersion struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	// Wire name: 'key'
	Key string ``
	// A descriptive name for this Spark version, for example "Spark 2.1".
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SparkVersion) MarshalJSON() ([]byte, error) {
	pb, err := SparkVersionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SparkVersion) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.SparkVersionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SparkVersionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SparkVersionToPb(st *SparkVersion) (*computepb.SparkVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.SparkVersionPb{}
	pb.Key = st.Key
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SparkVersionFromPb(pb *computepb.SparkVersionPb) (*SparkVersion, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkVersion{}
	st.Key = pb.Key
	st.Name = pb.Name

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type StartCluster struct {
	// The cluster to be started.
	// Wire name: 'cluster_id'
	ClusterId string ``
}

func (st StartCluster) MarshalJSON() ([]byte, error) {
	pb, err := StartClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *StartCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.StartClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := StartClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func StartClusterToPb(st *StartCluster) (*computepb.StartClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.StartClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func StartClusterFromPb(pb *computepb.StartClusterPb) (*StartCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
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

func StateToPb(st *State) (*computepb.StatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.StatePb(*st)
	return &pb, nil
}

func StateFromPb(pb *computepb.StatePb) (*State, error) {
	if pb == nil {
		return nil, nil
	}
	st := State(*pb)
	return &st, nil
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	// Wire name: 'code'
	Code TerminationReasonCode ``
	// list of parameters that provide additional information about why the
	// cluster was terminated
	// Wire name: 'parameters'
	Parameters map[string]string ``
	// type of the termination
	// Wire name: 'type'
	Type TerminationReasonType ``
}

func (st TerminationReason) MarshalJSON() ([]byte, error) {
	pb, err := TerminationReasonToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TerminationReason) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.TerminationReasonPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TerminationReasonFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TerminationReasonToPb(st *TerminationReason) (*computepb.TerminationReasonPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.TerminationReasonPb{}
	codePb, err := TerminationReasonCodeToPb(&st.Code)
	if err != nil {
		return nil, err
	}
	if codePb != nil {
		pb.Code = *codePb
	}
	pb.Parameters = st.Parameters
	typePb, err := TerminationReasonTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	return pb, nil
}

func TerminationReasonFromPb(pb *computepb.TerminationReasonPb) (*TerminationReason, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TerminationReason{}
	codeField, err := TerminationReasonCodeFromPb(&pb.Code)
	if err != nil {
		return nil, err
	}
	if codeField != nil {
		st.Code = *codeField
	}
	st.Parameters = pb.Parameters
	typeField, err := TerminationReasonTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	return st, nil
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

func TerminationReasonCodeToPb(st *TerminationReasonCode) (*computepb.TerminationReasonCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.TerminationReasonCodePb(*st)
	return &pb, nil
}

func TerminationReasonCodeFromPb(pb *computepb.TerminationReasonCodePb) (*TerminationReasonCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationReasonCode(*pb)
	return &st, nil
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

func TerminationReasonTypeToPb(st *TerminationReasonType) (*computepb.TerminationReasonTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := computepb.TerminationReasonTypePb(*st)
	return &pb, nil
}

func TerminationReasonTypeFromPb(pb *computepb.TerminationReasonTypePb) (*TerminationReasonType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationReasonType(*pb)
	return &st, nil
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	// Wire name: 'cluster_id'
	ClusterId string ``
	// The libraries to uninstall.
	// Wire name: 'libraries'
	Libraries []Library ``
}

func (st UninstallLibraries) MarshalJSON() ([]byte, error) {
	pb, err := UninstallLibrariesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UninstallLibraries) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.UninstallLibrariesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UninstallLibrariesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UninstallLibrariesToPb(st *UninstallLibraries) (*computepb.UninstallLibrariesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.UninstallLibrariesPb{}
	pb.ClusterId = st.ClusterId

	var librariesPb []computepb.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	return pb, nil
}

func UninstallLibrariesFromPb(pb *computepb.UninstallLibrariesPb) (*UninstallLibraries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UninstallLibraries{}
	st.ClusterId = pb.ClusterId

	var librariesField []Library
	for _, itemPb := range pb.Libraries {
		item, err := LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField

	return st, nil
}

type UnpinCluster struct {

	// Wire name: 'cluster_id'
	ClusterId string ``
}

func (st UnpinCluster) MarshalJSON() ([]byte, error) {
	pb, err := UnpinClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UnpinCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.UnpinClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UnpinClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UnpinClusterToPb(st *UnpinCluster) (*computepb.UnpinClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.UnpinClusterPb{}
	pb.ClusterId = st.ClusterId

	return pb, nil
}

func UnpinClusterFromPb(pb *computepb.UnpinClusterPb) (*UnpinCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpinCluster{}
	st.ClusterId = pb.ClusterId

	return st, nil
}

type UpdateCluster struct {
	// The cluster to be updated.
	// Wire name: 'cluster'
	Cluster *UpdateClusterResource ``
	// ID of the cluster.
	// Wire name: 'cluster_id'
	ClusterId string ``
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
	UpdateMask string `` //legacy

}

func (st UpdateCluster) MarshalJSON() ([]byte, error) {
	pb, err := UpdateClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.UpdateClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateClusterToPb(st *UpdateCluster) (*computepb.UpdateClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.UpdateClusterPb{}
	clusterPb, err := UpdateClusterResourceToPb(st.Cluster)
	if err != nil {
		return nil, err
	}
	if clusterPb != nil {
		pb.Cluster = clusterPb
	}
	pb.ClusterId = st.ClusterId
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

func UpdateClusterFromPb(pb *computepb.UpdateClusterPb) (*UpdateCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCluster{}
	clusterField, err := UpdateClusterResourceFromPb(pb.Cluster)
	if err != nil {
		return nil, err
	}
	if clusterField != nil {
		st.Cluster = clusterField
	}
	st.ClusterId = pb.ClusterId
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

type UpdateClusterResource struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	// Wire name: 'autoscale'
	Autoscale *AutoScale ``
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	// Wire name: 'autotermination_minutes'
	AutoterminationMinutes int ``
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'aws_attributes'
	AwsAttributes *AwsAttributes ``
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'azure_attributes'
	AzureAttributes *AzureAttributes ``
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	// Wire name: 'cluster_log_conf'
	ClusterLogConf *ClusterLogConf ``
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string. For
	// job clusters, the cluster name is automatically set based on the job and
	// job run IDs.
	// Wire name: 'cluster_name'
	ClusterName string ``
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	// Wire name: 'custom_tags'
	CustomTags map[string]string ``

	// Wire name: 'data_security_mode'
	DataSecurityMode DataSecurityMode ``
	// Custom docker image BYOC
	// Wire name: 'docker_image'
	DockerImage *DockerImage ``
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	// Wire name: 'driver_instance_pool_id'
	DriverInstancePoolId string ``
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	// Wire name: 'driver_node_type_id'
	DriverNodeTypeId string ``
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	// Wire name: 'enable_elastic_disk'
	EnableElasticDisk bool ``
	// Whether to enable LUKS on cluster VMs' local disks
	// Wire name: 'enable_local_disk_encryption'
	EnableLocalDiskEncryption bool ``
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	// Wire name: 'gcp_attributes'
	GcpAttributes *GcpAttributes ``
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	// Wire name: 'init_scripts'
	InitScripts []InitScriptInfo ``
	// The optional ID of the instance pool to which the cluster belongs.
	// Wire name: 'instance_pool_id'
	InstancePoolId string ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	// Wire name: 'is_single_node'
	IsSingleNode bool ``

	// Wire name: 'kind'
	Kind Kind ``
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	// Wire name: 'node_type_id'
	NodeTypeId string ``
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
	NumWorkers int ``
	// The ID of the cluster policy used to create the cluster if applicable.
	// Wire name: 'policy_id'
	PolicyId string ``
	// If set, what the configurable throughput (in Mb/s) for the remote disk
	// is. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'remote_disk_throughput'
	RemoteDiskThroughput int ``
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	// Wire name: 'runtime_engine'
	RuntimeEngine RuntimeEngine ``
	// Single user name if data_security_mode is `SINGLE_USER`
	// Wire name: 'single_user_name'
	SingleUserName string ``
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	// Wire name: 'spark_conf'
	SparkConf map[string]string ``
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
	SparkEnvVars map[string]string ``
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	// Wire name: 'spark_version'
	SparkVersion string ``
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	// Wire name: 'ssh_public_keys'
	SshPublicKeys []string ``
	// If set, what the total initial volume size (in GB) of the remote disks
	// should be. Currently only supported for GCP HYPERDISK_BALANCED disks.
	// Wire name: 'total_initial_remote_disk_size'
	TotalInitialRemoteDiskSize int ``
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	// Wire name: 'use_ml_runtime'
	UseMlRuntime bool ``

	// Wire name: 'workload_type'
	WorkloadType    *WorkloadType ``
	ForceSendFields []string      `tf:"-"`
}

func (st UpdateClusterResource) MarshalJSON() ([]byte, error) {
	pb, err := UpdateClusterResourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateClusterResource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.UpdateClusterResourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateClusterResourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateClusterResourceToPb(st *UpdateClusterResource) (*computepb.UpdateClusterResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.UpdateClusterResourcePb{}
	autoscalePb, err := AutoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}
	pb.AutoterminationMinutes = st.AutoterminationMinutes
	awsAttributesPb, err := AwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}
	azureAttributesPb, err := AzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}
	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}
	pb.ClusterName = st.ClusterName
	pb.CustomTags = st.CustomTags
	dataSecurityModePb, err := DataSecurityModeToPb(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}
	dockerImagePb, err := DockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableElasticDisk = st.EnableElasticDisk
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []computepb.InitScriptInfoPb
	for _, item := range st.InitScripts {
		itemPb, err := InitScriptInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			initScriptsPb = append(initScriptsPb, *itemPb)
		}
	}
	pb.InitScripts = initScriptsPb
	pb.InstancePoolId = st.InstancePoolId
	pb.IsSingleNode = st.IsSingleNode
	kindPb, err := KindToPb(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.RemoteDiskThroughput = st.RemoteDiskThroughput
	runtimeEnginePb, err := RuntimeEngineToPb(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}
	pb.SingleUserName = st.SingleUserName
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SparkVersion = st.SparkVersion
	pb.SshPublicKeys = st.SshPublicKeys
	pb.TotalInitialRemoteDiskSize = st.TotalInitialRemoteDiskSize
	pb.UseMlRuntime = st.UseMlRuntime
	workloadTypePb, err := WorkloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateClusterResourceFromPb(pb *computepb.UpdateClusterResourcePb) (*UpdateClusterResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateClusterResource{}
	autoscaleField, err := AutoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	st.AutoterminationMinutes = pb.AutoterminationMinutes
	awsAttributesField, err := AwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := AzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	st.ClusterName = pb.ClusterName
	st.CustomTags = pb.CustomTags
	dataSecurityModeField, err := DataSecurityModeFromPb(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := DockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableElasticDisk = pb.EnableElasticDisk
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	gcpAttributesField, err := GcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}

	var initScriptsField []InitScriptInfo
	for _, itemPb := range pb.InitScripts {
		item, err := InitScriptInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			initScriptsField = append(initScriptsField, *item)
		}
	}
	st.InitScripts = initScriptsField
	st.InstancePoolId = pb.InstancePoolId
	st.IsSingleNode = pb.IsSingleNode
	kindField, err := KindFromPb(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.RemoteDiskThroughput = pb.RemoteDiskThroughput
	runtimeEngineField, err := RuntimeEngineFromPb(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	st.SingleUserName = pb.SingleUserName
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SparkVersion = pb.SparkVersion
	st.SshPublicKeys = pb.SshPublicKeys
	st.TotalInitialRemoteDiskSize = pb.TotalInitialRemoteDiskSize
	st.UseMlRuntime = pb.UseMlRuntime
	workloadTypeField, err := WorkloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A storage location back by UC Volumes.
type VolumesStorageInfo struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	// Wire name: 'destination'
	Destination string ``
}

func (st VolumesStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := VolumesStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *VolumesStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.VolumesStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := VolumesStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func VolumesStorageInfoToPb(st *VolumesStorageInfo) (*computepb.VolumesStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.VolumesStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

func VolumesStorageInfoFromPb(pb *computepb.VolumesStorageInfoPb) (*VolumesStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumesStorageInfo{}
	st.Destination = pb.Destination

	return st, nil
}

// Cluster Attributes showing for clusters workload types.
type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	// Wire name: 'clients'
	Clients ClientsTypes ``
}

func (st WorkloadType) MarshalJSON() ([]byte, error) {
	pb, err := WorkloadTypeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkloadType) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.WorkloadTypePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkloadTypeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WorkloadTypeToPb(st *WorkloadType) (*computepb.WorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.WorkloadTypePb{}
	clientsPb, err := ClientsTypesToPb(&st.Clients)
	if err != nil {
		return nil, err
	}
	if clientsPb != nil {
		pb.Clients = *clientsPb
	}

	return pb, nil
}

func WorkloadTypeFromPb(pb *computepb.WorkloadTypePb) (*WorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkloadType{}
	clientsField, err := ClientsTypesFromPb(&pb.Clients)
	if err != nil {
		return nil, err
	}
	if clientsField != nil {
		st.Clients = *clientsField
	}

	return st, nil
}

// A storage location in Workspace Filesystem (WSFS)
type WorkspaceStorageInfo struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	// Wire name: 'destination'
	Destination string ``
}

func (st WorkspaceStorageInfo) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceStorageInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceStorageInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computepb.WorkspaceStorageInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceStorageInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WorkspaceStorageInfoToPb(st *WorkspaceStorageInfo) (*computepb.WorkspaceStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computepb.WorkspaceStorageInfoPb{}
	pb.Destination = st.Destination

	return pb, nil
}

func WorkspaceStorageInfoFromPb(pb *computepb.WorkspaceStorageInfoPb) (*WorkspaceStorageInfo, error) {
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
