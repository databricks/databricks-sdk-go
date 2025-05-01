// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type AddInstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn string
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn string
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile bool
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	SkipValidation bool

	ForceSendFields []string
}

func addInstanceProfileToPb(st *AddInstanceProfile) (*addInstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addInstanceProfilePb{}
	iamRoleArnPb, err := identity(&st.IamRoleArn)
	if err != nil {
		return nil, err
	}
	if iamRoleArnPb != nil {
		pb.IamRoleArn = *iamRoleArnPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	isMetaInstanceProfilePb, err := identity(&st.IsMetaInstanceProfile)
	if err != nil {
		return nil, err
	}
	if isMetaInstanceProfilePb != nil {
		pb.IsMetaInstanceProfile = *isMetaInstanceProfilePb
	}

	skipValidationPb, err := identity(&st.SkipValidation)
	if err != nil {
		return nil, err
	}
	if skipValidationPb != nil {
		pb.SkipValidation = *skipValidationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type addInstanceProfilePb struct {
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

func addInstanceProfileFromPb(pb *addInstanceProfilePb) (*AddInstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddInstanceProfile{}
	iamRoleArnField, err := identity(&pb.IamRoleArn)
	if err != nil {
		return nil, err
	}
	if iamRoleArnField != nil {
		st.IamRoleArn = *iamRoleArnField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	isMetaInstanceProfileField, err := identity(&pb.IsMetaInstanceProfile)
	if err != nil {
		return nil, err
	}
	if isMetaInstanceProfileField != nil {
		st.IsMetaInstanceProfile = *isMetaInstanceProfileField
	}
	skipValidationField, err := identity(&pb.SkipValidation)
	if err != nil {
		return nil, err
	}
	if skipValidationField != nil {
		st.SkipValidation = *skipValidationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *addInstanceProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st addInstanceProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AddResponse struct {
}

func addResponseToPb(st *AddResponse) (*addResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addResponsePb{}

	return pb, nil
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

type addResponsePb struct {
}

func addResponseFromPb(pb *addResponsePb) (*AddResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddResponse{}

	return st, nil
}

// A storage location in Adls Gen2
type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination string
}

func adlsgen2InfoToPb(st *Adlsgen2Info) (*adlsgen2InfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &adlsgen2InfoPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	return pb, nil
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

type adlsgen2InfoPb struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination string `json:"destination"`
}

func adlsgen2InfoFromPb(pb *adlsgen2InfoPb) (*Adlsgen2Info, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Adlsgen2Info{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}

	return st, nil
}

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	MaxWorkers int
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers int

	ForceSendFields []string
}

func autoScaleToPb(st *AutoScale) (*autoScalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoScalePb{}
	maxWorkersPb, err := identity(&st.MaxWorkers)
	if err != nil {
		return nil, err
	}
	if maxWorkersPb != nil {
		pb.MaxWorkers = *maxWorkersPb
	}

	minWorkersPb, err := identity(&st.MinWorkers)
	if err != nil {
		return nil, err
	}
	if minWorkersPb != nil {
		pb.MinWorkers = *minWorkersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type autoScalePb struct {
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

func autoScaleFromPb(pb *autoScalePb) (*AutoScale, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoScale{}
	maxWorkersField, err := identity(&pb.MaxWorkers)
	if err != nil {
		return nil, err
	}
	if maxWorkersField != nil {
		st.MaxWorkers = *maxWorkersField
	}
	minWorkersField, err := identity(&pb.MinWorkers)
	if err != nil {
		return nil, err
	}
	if minWorkersField != nil {
		st.MinWorkers = *minWorkersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *autoScalePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st autoScalePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Attributes set during cluster creation which are related to Amazon Web
// Services.
type AwsAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones.
	//
	// Note: If `first_on_demand` is zero, this availability type will be used
	// for the entire cluster.
	Availability AwsAvailability
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
	EbsVolumeCount int
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	EbsVolumeIops int
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	EbsVolumeSize int
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	EbsVolumeThroughput int
	// The type of EBS volumes that will be launched with this cluster.
	EbsVolumeType EbsVolumeType
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. If this value is greater than 0, the cluster driver
	// node in particular will be placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand int
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	InstanceProfileArn string
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
	SpotBidPricePercent int
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
	ZoneId string

	ForceSendFields []string
}

func AwsAttributesToPb(st *AwsAttributes) (*AwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &AwsAttributesPb{}
	availabilityPb, err := identity(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}

	ebsVolumeCountPb, err := identity(&st.EbsVolumeCount)
	if err != nil {
		return nil, err
	}
	if ebsVolumeCountPb != nil {
		pb.EbsVolumeCount = *ebsVolumeCountPb
	}

	ebsVolumeIopsPb, err := identity(&st.EbsVolumeIops)
	if err != nil {
		return nil, err
	}
	if ebsVolumeIopsPb != nil {
		pb.EbsVolumeIops = *ebsVolumeIopsPb
	}

	ebsVolumeSizePb, err := identity(&st.EbsVolumeSize)
	if err != nil {
		return nil, err
	}
	if ebsVolumeSizePb != nil {
		pb.EbsVolumeSize = *ebsVolumeSizePb
	}

	ebsVolumeThroughputPb, err := identity(&st.EbsVolumeThroughput)
	if err != nil {
		return nil, err
	}
	if ebsVolumeThroughputPb != nil {
		pb.EbsVolumeThroughput = *ebsVolumeThroughputPb
	}

	ebsVolumeTypePb, err := identity(&st.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypePb != nil {
		pb.EbsVolumeType = *ebsVolumeTypePb
	}

	firstOnDemandPb, err := identity(&st.FirstOnDemand)
	if err != nil {
		return nil, err
	}
	if firstOnDemandPb != nil {
		pb.FirstOnDemand = *firstOnDemandPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	spotBidPricePercentPb, err := identity(&st.SpotBidPricePercent)
	if err != nil {
		return nil, err
	}
	if spotBidPricePercentPb != nil {
		pb.SpotBidPricePercent = *spotBidPricePercentPb
	}

	zoneIdPb, err := identity(&st.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdPb != nil {
		pb.ZoneId = *zoneIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AwsAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &AwsAttributesPb{}
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

func (st AwsAttributes) MarshalJSON() ([]byte, error) {
	pb, err := AwsAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AwsAttributesPb struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones.
	//
	// Note: If `first_on_demand` is zero, this availability type will be used
	// for the entire cluster.
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

func AwsAttributesFromPb(pb *AwsAttributesPb) (*AwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AwsAttributes{}
	availabilityField, err := identity(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	ebsVolumeCountField, err := identity(&pb.EbsVolumeCount)
	if err != nil {
		return nil, err
	}
	if ebsVolumeCountField != nil {
		st.EbsVolumeCount = *ebsVolumeCountField
	}
	ebsVolumeIopsField, err := identity(&pb.EbsVolumeIops)
	if err != nil {
		return nil, err
	}
	if ebsVolumeIopsField != nil {
		st.EbsVolumeIops = *ebsVolumeIopsField
	}
	ebsVolumeSizeField, err := identity(&pb.EbsVolumeSize)
	if err != nil {
		return nil, err
	}
	if ebsVolumeSizeField != nil {
		st.EbsVolumeSize = *ebsVolumeSizeField
	}
	ebsVolumeThroughputField, err := identity(&pb.EbsVolumeThroughput)
	if err != nil {
		return nil, err
	}
	if ebsVolumeThroughputField != nil {
		st.EbsVolumeThroughput = *ebsVolumeThroughputField
	}
	ebsVolumeTypeField, err := identity(&pb.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypeField != nil {
		st.EbsVolumeType = *ebsVolumeTypeField
	}
	firstOnDemandField, err := identity(&pb.FirstOnDemand)
	if err != nil {
		return nil, err
	}
	if firstOnDemandField != nil {
		st.FirstOnDemand = *firstOnDemandField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	spotBidPricePercentField, err := identity(&pb.SpotBidPricePercent)
	if err != nil {
		return nil, err
	}
	if spotBidPricePercentField != nil {
		st.SpotBidPricePercent = *spotBidPricePercentField
	}
	zoneIdField, err := identity(&pb.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdField != nil {
		st.ZoneId = *zoneIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *AwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones.
//
// Note: If `first_on_demand` is zero, this availability type will be used for
// the entire cluster.
type AwsAvailability string
type awsAvailabilityPb string

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

// Type always returns AwsAvailability to satisfy [pflag.Value] interface
func (f *AwsAvailability) Type() string {
	return "AwsAvailability"
}

func awsAvailabilityToPb(st *AwsAvailability) (*awsAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := awsAvailabilityPb(*st)
	return &pb, nil
}

func awsAvailabilityFromPb(pb *awsAvailabilityPb) (*AwsAvailability, error) {
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
	Availability AzureAvailability
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. This value should be greater than 0, to make sure
	// the cluster driver node is placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand int
	// Defines values necessary to configure and run Azure Log Analytics agent
	LogAnalyticsInfo *LogAnalyticsInfo
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice float64

	ForceSendFields []string
}

func AzureAttributesToPb(st *AzureAttributes) (*AzureAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &AzureAttributesPb{}
	availabilityPb, err := identity(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}

	firstOnDemandPb, err := identity(&st.FirstOnDemand)
	if err != nil {
		return nil, err
	}
	if firstOnDemandPb != nil {
		pb.FirstOnDemand = *firstOnDemandPb
	}

	logAnalyticsInfoPb, err := logAnalyticsInfoToPb(st.LogAnalyticsInfo)
	if err != nil {
		return nil, err
	}
	if logAnalyticsInfoPb != nil {
		pb.LogAnalyticsInfo = logAnalyticsInfoPb
	}

	spotBidMaxPricePb, err := identity(&st.SpotBidMaxPrice)
	if err != nil {
		return nil, err
	}
	if spotBidMaxPricePb != nil {
		pb.SpotBidMaxPrice = *spotBidMaxPricePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AzureAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &AzureAttributesPb{}
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

func (st AzureAttributes) MarshalJSON() ([]byte, error) {
	pb, err := AzureAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AzureAttributesPb struct {
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
	LogAnalyticsInfo *logAnalyticsInfoPb `json:"log_analytics_info,omitempty"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice float64 `json:"spot_bid_max_price,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func AzureAttributesFromPb(pb *AzureAttributesPb) (*AzureAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AzureAttributes{}
	availabilityField, err := identity(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	firstOnDemandField, err := identity(&pb.FirstOnDemand)
	if err != nil {
		return nil, err
	}
	if firstOnDemandField != nil {
		st.FirstOnDemand = *firstOnDemandField
	}
	logAnalyticsInfoField, err := logAnalyticsInfoFromPb(pb.LogAnalyticsInfo)
	if err != nil {
		return nil, err
	}
	if logAnalyticsInfoField != nil {
		st.LogAnalyticsInfo = logAnalyticsInfoField
	}
	spotBidMaxPriceField, err := identity(&pb.SpotBidMaxPrice)
	if err != nil {
		return nil, err
	}
	if spotBidMaxPriceField != nil {
		st.SpotBidMaxPrice = *spotBidMaxPriceField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *AzureAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AzureAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones. Note: If `first_on_demand` is zero, this availability type will be used
// for the entire cluster.
type AzureAvailability string
type azureAvailabilityPb string

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

// Type always returns AzureAvailability to satisfy [pflag.Value] interface
func (f *AzureAvailability) Type() string {
	return "AzureAvailability"
}

func azureAvailabilityToPb(st *AzureAvailability) (*azureAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := azureAvailabilityPb(*st)
	return &pb, nil
}

func azureAvailabilityFromPb(pb *azureAvailabilityPb) (*AzureAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := AzureAvailability(*pb)
	return &st, nil
}

type CancelCommand struct {
	ClusterId string

	CommandId string

	ContextId string

	ForceSendFields []string
}

func cancelCommandToPb(st *CancelCommand) (*cancelCommandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelCommandPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	commandIdPb, err := identity(&st.CommandId)
	if err != nil {
		return nil, err
	}
	if commandIdPb != nil {
		pb.CommandId = *commandIdPb
	}

	contextIdPb, err := identity(&st.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdPb != nil {
		pb.ContextId = *contextIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	commandIdField, err := identity(&pb.CommandId)
	if err != nil {
		return nil, err
	}
	if commandIdField != nil {
		st.CommandId = *commandIdField
	}
	contextIdField, err := identity(&pb.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdField != nil {
		st.ContextId = *contextIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cancelCommandPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cancelCommandPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelResponse struct {
}

func cancelResponseToPb(st *CancelResponse) (*cancelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelResponsePb{}

	return pb, nil
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

type cancelResponsePb struct {
}

func cancelResponseFromPb(pb *cancelResponsePb) (*CancelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelResponse{}

	return st, nil
}

type ChangeClusterOwner struct {
	ClusterId string
	// New owner of the cluster_id after this RPC.
	OwnerUsername string
}

func changeClusterOwnerToPb(st *ChangeClusterOwner) (*changeClusterOwnerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &changeClusterOwnerPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	ownerUsernamePb, err := identity(&st.OwnerUsername)
	if err != nil {
		return nil, err
	}
	if ownerUsernamePb != nil {
		pb.OwnerUsername = *ownerUsernamePb
	}

	return pb, nil
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

type changeClusterOwnerPb struct {
	ClusterId string `json:"cluster_id"`
	// New owner of the cluster_id after this RPC.
	OwnerUsername string `json:"owner_username"`
}

func changeClusterOwnerFromPb(pb *changeClusterOwnerPb) (*ChangeClusterOwner, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChangeClusterOwner{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	ownerUsernameField, err := identity(&pb.OwnerUsername)
	if err != nil {
		return nil, err
	}
	if ownerUsernameField != nil {
		st.OwnerUsername = *ownerUsernameField
	}

	return st, nil
}

type ChangeClusterOwnerResponse struct {
}

func changeClusterOwnerResponseToPb(st *ChangeClusterOwnerResponse) (*changeClusterOwnerResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &changeClusterOwnerResponsePb{}

	return pb, nil
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

type changeClusterOwnerResponsePb struct {
}

func changeClusterOwnerResponseFromPb(pb *changeClusterOwnerResponsePb) (*ChangeClusterOwnerResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChangeClusterOwnerResponse{}

	return st, nil
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	Jobs bool
	// With notebooks set, this cluster can be used for notebooks
	Notebooks bool

	ForceSendFields []string
}

func clientsTypesToPb(st *ClientsTypes) (*clientsTypesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clientsTypesPb{}
	jobsPb, err := identity(&st.Jobs)
	if err != nil {
		return nil, err
	}
	if jobsPb != nil {
		pb.Jobs = *jobsPb
	}

	notebooksPb, err := identity(&st.Notebooks)
	if err != nil {
		return nil, err
	}
	if notebooksPb != nil {
		pb.Notebooks = *notebooksPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clientsTypesPb struct {
	// With jobs set, the cluster can be used for jobs
	Jobs bool `json:"jobs,omitempty"`
	// With notebooks set, this cluster can be used for notebooks
	Notebooks bool `json:"notebooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clientsTypesFromPb(pb *clientsTypesPb) (*ClientsTypes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClientsTypes{}
	jobsField, err := identity(&pb.Jobs)
	if err != nil {
		return nil, err
	}
	if jobsField != nil {
		st.Jobs = *jobsField
	}
	notebooksField, err := identity(&pb.Notebooks)
	if err != nil {
		return nil, err
	}
	if notebooksField != nil {
		st.Notebooks = *notebooksField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clientsTypesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clientsTypesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CloneCluster struct {
	// The cluster that is being cloned.
	SourceClusterId string
}

func cloneClusterToPb(st *CloneCluster) (*cloneClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloneClusterPb{}
	sourceClusterIdPb, err := identity(&st.SourceClusterId)
	if err != nil {
		return nil, err
	}
	if sourceClusterIdPb != nil {
		pb.SourceClusterId = *sourceClusterIdPb
	}

	return pb, nil
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

type cloneClusterPb struct {
	// The cluster that is being cloned.
	SourceClusterId string `json:"source_cluster_id"`
}

func cloneClusterFromPb(pb *cloneClusterPb) (*CloneCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloneCluster{}
	sourceClusterIdField, err := identity(&pb.SourceClusterId)
	if err != nil {
		return nil, err
	}
	if sourceClusterIdField != nil {
		st.SourceClusterId = *sourceClusterIdField
	}

	return st, nil
}

type CloudProviderNodeInfo struct {
	// Status as reported by the cloud provider
	Status []CloudProviderNodeStatus
}

func cloudProviderNodeInfoToPb(st *CloudProviderNodeInfo) (*cloudProviderNodeInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cloudProviderNodeInfoPb{}

	var statusPb []CloudProviderNodeStatus
	for _, item := range st.Status {
		itemPb, err := identity(&item)
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

type cloudProviderNodeInfoPb struct {
	// Status as reported by the cloud provider
	Status []CloudProviderNodeStatus `json:"status,omitempty"`
}

func cloudProviderNodeInfoFromPb(pb *cloudProviderNodeInfoPb) (*CloudProviderNodeInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloudProviderNodeInfo{}

	var statusField []CloudProviderNodeStatus
	for _, itemPb := range pb.Status {
		item, err := identity(&itemPb)
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
type cloudProviderNodeStatusPb string

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

// Type always returns CloudProviderNodeStatus to satisfy [pflag.Value] interface
func (f *CloudProviderNodeStatus) Type() string {
	return "CloudProviderNodeStatus"
}

func cloudProviderNodeStatusToPb(st *CloudProviderNodeStatus) (*cloudProviderNodeStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cloudProviderNodeStatusPb(*st)
	return &pb, nil
}

func cloudProviderNodeStatusFromPb(pb *cloudProviderNodeStatusPb) (*CloudProviderNodeStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := CloudProviderNodeStatus(*pb)
	return &st, nil
}

type ClusterAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel ClusterPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func clusterAccessControlRequestToPb(st *ClusterAccessControlRequest) (*clusterAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAccessControlRequestFromPb(pb *clusterAccessControlRequestPb) (*ClusterAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAccessControlRequest{}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAccessControlResponse struct {
	// All permissions.
	AllPermissions []ClusterPermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func clusterAccessControlResponseToPb(st *ClusterAccessControlResponse) (*clusterAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAccessControlResponsePb{}

	var allPermissionsPb []clusterPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := clusterPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []clusterPermissionPb `json:"all_permissions,omitempty"`
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

func clusterAccessControlResponseFromPb(pb *clusterAccessControlResponsePb) (*ClusterAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAccessControlResponse{}

	var allPermissionsField []ClusterPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := clusterPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Common set of attributes set during cluster creation. These attributes cannot
// be changed over the lifetime of a cluster.
type ClusterAttributes struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName string
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
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
	DataSecurityMode DataSecurityMode
	// Custom docker image BYOC
	DockerImage *DockerImage
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool
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
	Kind Kind
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string
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
	SparkEnvVars map[string]string
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *WorkloadType

	ForceSendFields []string
}

func clusterAttributesToPb(st *ClusterAttributes) (*clusterAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAttributesPb{}
	autoterminationMinutesPb, err := identity(&st.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesPb != nil {
		pb.AutoterminationMinutes = *autoterminationMinutesPb
	}

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

	clusterNamePb, err := identity(&st.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNamePb != nil {
		pb.ClusterName = *clusterNamePb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	dataSecurityModePb, err := identity(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}

	dockerImagePb, err := dockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []InitScriptInfoPb
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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	isSingleNodePb, err := identity(&st.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodePb != nil {
		pb.IsSingleNode = *isSingleNodePb
	}

	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	runtimeEnginePb, err := identity(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}

	singleUserNamePb, err := identity(&st.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNamePb != nil {
		pb.SingleUserName = *singleUserNamePb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	sparkVersionPb, err := identity(&st.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionPb != nil {
		pb.SparkVersion = *sparkVersionPb
	}

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

	useMlRuntimePb, err := identity(&st.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimePb != nil {
		pb.UseMlRuntime = *useMlRuntimePb
	}

	workloadTypePb, err := workloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterAttributesPb struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributesPb `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
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
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *dockerImagePb `json:"docker_image,omitempty"`
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
	GcpAttributes *GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfoPb `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
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
	Kind Kind `json:"kind,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string `json:"policy_id,omitempty"`
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
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *workloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAttributesFromPb(pb *clusterAttributesPb) (*ClusterAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAttributes{}
	autoterminationMinutesField, err := identity(&pb.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesField != nil {
		st.AutoterminationMinutes = *autoterminationMinutesField
	}
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
	clusterNameField, err := identity(&pb.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNameField != nil {
		st.ClusterName = *clusterNameField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	dataSecurityModeField, err := identity(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := dockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	isSingleNodeField, err := identity(&pb.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodeField != nil {
		st.IsSingleNode = *isSingleNodeField
	}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	runtimeEngineField, err := identity(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	singleUserNameField, err := identity(&pb.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNameField != nil {
		st.SingleUserName = *singleUserNameField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField
	sparkVersionField, err := identity(&pb.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionField != nil {
		st.SparkVersion = *sparkVersionField
	}

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField
	useMlRuntimeField, err := identity(&pb.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimeField != nil {
		st.UseMlRuntime = *useMlRuntimeField
	}
	workloadTypeField, err := workloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterCompliance struct {
	// Canonical unique identifier for a cluster.
	ClusterId string
	// Whether this cluster is in compliance with the latest version of its
	// policy.
	IsCompliant bool
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations map[string]string

	ForceSendFields []string
}

func clusterComplianceToPb(st *ClusterCompliance) (*clusterCompliancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterCompliancePb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	isCompliantPb, err := identity(&st.IsCompliant)
	if err != nil {
		return nil, err
	}
	if isCompliantPb != nil {
		pb.IsCompliant = *isCompliantPb
	}

	violationsPb := map[string]string{}
	for k, v := range st.Violations {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			violationsPb[k] = *itemPb
		}
	}
	pb.Violations = violationsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterCompliancePb struct {
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

func clusterComplianceFromPb(pb *clusterCompliancePb) (*ClusterCompliance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterCompliance{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	isCompliantField, err := identity(&pb.IsCompliant)
	if err != nil {
		return nil, err
	}
	if isCompliantField != nil {
		st.IsCompliant = *isCompliantField
	}

	violationsField := map[string]string{}
	for k, v := range pb.Violations {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			violationsField[k] = *item
		}
	}
	st.Violations = violationsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterCompliancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterCompliancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Describes all of the metadata about a single Spark cluster in Databricks.
type ClusterDetails struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes
	// Number of CPU cores available for this cluster. Note that this can be
	// fractional, e.g. 7.5 cores, since certain node types are configured to
	// share cores between Spark nodes on the same instance.
	ClusterCores float64
	// Canonical identifier for the cluster. This id is retained during cluster
	// restarts and resizes, while each new cluster has a globally unique id.
	ClusterId string
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf
	// Cluster log delivery status.
	ClusterLogStatus *LogSyncStatus
	// Total amount of cluster memory, in megabytes
	ClusterMemoryMb int64
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName string
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request.
	ClusterSource ClusterSource
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName string
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
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
	DataSecurityMode DataSecurityMode
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
	DefaultTags map[string]string
	// Custom docker image BYOC
	DockerImage *DockerImage
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	Driver *SparkNode
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool
	// Nodes on which the Spark executors reside.
	Executors []SparkNode
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	JdbcPort int
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
	Kind Kind
	// the timestamp that the cluster was started/restarted
	LastRestartedTime int64
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	LastStateLossTime int64
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	SparkContextId int64
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
	SparkEnvVars map[string]string
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string
	// The spec contains a snapshot of the latest user specified settings that
	// were used to create/edit the cluster. Note: not included in the response
	// of the ListClusters API.
	Spec *ClusterSpec
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	StartTime int64
	// Current state of the cluster.
	State State
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	StateMessage string
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	TerminatedTime int64
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	TerminationReason *TerminationReason
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *WorkloadType

	ForceSendFields []string
}

func clusterDetailsToPb(st *ClusterDetails) (*clusterDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterDetailsPb{}
	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	autoterminationMinutesPb, err := identity(&st.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesPb != nil {
		pb.AutoterminationMinutes = *autoterminationMinutesPb
	}

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

	clusterCoresPb, err := identity(&st.ClusterCores)
	if err != nil {
		return nil, err
	}
	if clusterCoresPb != nil {
		pb.ClusterCores = *clusterCoresPb
	}

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}

	clusterLogStatusPb, err := logSyncStatusToPb(st.ClusterLogStatus)
	if err != nil {
		return nil, err
	}
	if clusterLogStatusPb != nil {
		pb.ClusterLogStatus = clusterLogStatusPb
	}

	clusterMemoryMbPb, err := identity(&st.ClusterMemoryMb)
	if err != nil {
		return nil, err
	}
	if clusterMemoryMbPb != nil {
		pb.ClusterMemoryMb = *clusterMemoryMbPb
	}

	clusterNamePb, err := identity(&st.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNamePb != nil {
		pb.ClusterName = *clusterNamePb
	}

	clusterSourcePb, err := identity(&st.ClusterSource)
	if err != nil {
		return nil, err
	}
	if clusterSourcePb != nil {
		pb.ClusterSource = *clusterSourcePb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	dataSecurityModePb, err := identity(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}

	defaultTagsPb := map[string]string{}
	for k, v := range st.DefaultTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			defaultTagsPb[k] = *itemPb
		}
	}
	pb.DefaultTags = defaultTagsPb

	dockerImagePb, err := dockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}

	driverPb, err := sparkNodeToPb(st.Driver)
	if err != nil {
		return nil, err
	}
	if driverPb != nil {
		pb.Driver = driverPb
	}

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

	var executorsPb []sparkNodePb
	for _, item := range st.Executors {
		itemPb, err := sparkNodeToPb(&item)
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

	var initScriptsPb []InitScriptInfoPb
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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	isSingleNodePb, err := identity(&st.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodePb != nil {
		pb.IsSingleNode = *isSingleNodePb
	}

	jdbcPortPb, err := identity(&st.JdbcPort)
	if err != nil {
		return nil, err
	}
	if jdbcPortPb != nil {
		pb.JdbcPort = *jdbcPortPb
	}

	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	lastRestartedTimePb, err := identity(&st.LastRestartedTime)
	if err != nil {
		return nil, err
	}
	if lastRestartedTimePb != nil {
		pb.LastRestartedTime = *lastRestartedTimePb
	}

	lastStateLossTimePb, err := identity(&st.LastStateLossTime)
	if err != nil {
		return nil, err
	}
	if lastStateLossTimePb != nil {
		pb.LastStateLossTime = *lastStateLossTimePb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	runtimeEnginePb, err := identity(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}

	singleUserNamePb, err := identity(&st.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNamePb != nil {
		pb.SingleUserName = *singleUserNamePb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkContextIdPb, err := identity(&st.SparkContextId)
	if err != nil {
		return nil, err
	}
	if sparkContextIdPb != nil {
		pb.SparkContextId = *sparkContextIdPb
	}

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	sparkVersionPb, err := identity(&st.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionPb != nil {
		pb.SparkVersion = *sparkVersionPb
	}

	specPb, err := ClusterSpecToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	stateMessagePb, err := identity(&st.StateMessage)
	if err != nil {
		return nil, err
	}
	if stateMessagePb != nil {
		pb.StateMessage = *stateMessagePb
	}

	terminatedTimePb, err := identity(&st.TerminatedTime)
	if err != nil {
		return nil, err
	}
	if terminatedTimePb != nil {
		pb.TerminatedTime = *terminatedTimePb
	}

	terminationReasonPb, err := terminationReasonToPb(st.TerminationReason)
	if err != nil {
		return nil, err
	}
	if terminationReasonPb != nil {
		pb.TerminationReason = terminationReasonPb
	}

	useMlRuntimePb, err := identity(&st.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimePb != nil {
		pb.UseMlRuntime = *useMlRuntimePb
	}

	workloadTypePb, err := workloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterDetailsPb struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributesPb `json:"azure_attributes,omitempty"`
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
	ClusterLogConf *ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
	// Cluster log delivery status.
	ClusterLogStatus *logSyncStatusPb `json:"cluster_log_status,omitempty"`
	// Total amount of cluster memory, in megabytes
	ClusterMemoryMb int64 `json:"cluster_memory_mb,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
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
	DockerImage *dockerImagePb `json:"docker_image,omitempty"`
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	Driver *sparkNodePb `json:"driver,omitempty"`
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
	Executors []sparkNodePb `json:"executors,omitempty"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfoPb `json:"init_scripts,omitempty"`
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
	Spec *ClusterSpecPb `json:"spec,omitempty"`
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
	TerminationReason *terminationReasonPb `json:"termination_reason,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *workloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterDetailsFromPb(pb *clusterDetailsPb) (*ClusterDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterDetails{}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	autoterminationMinutesField, err := identity(&pb.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesField != nil {
		st.AutoterminationMinutes = *autoterminationMinutesField
	}
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
	clusterCoresField, err := identity(&pb.ClusterCores)
	if err != nil {
		return nil, err
	}
	if clusterCoresField != nil {
		st.ClusterCores = *clusterCoresField
	}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	clusterLogStatusField, err := logSyncStatusFromPb(pb.ClusterLogStatus)
	if err != nil {
		return nil, err
	}
	if clusterLogStatusField != nil {
		st.ClusterLogStatus = clusterLogStatusField
	}
	clusterMemoryMbField, err := identity(&pb.ClusterMemoryMb)
	if err != nil {
		return nil, err
	}
	if clusterMemoryMbField != nil {
		st.ClusterMemoryMb = *clusterMemoryMbField
	}
	clusterNameField, err := identity(&pb.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNameField != nil {
		st.ClusterName = *clusterNameField
	}
	clusterSourceField, err := identity(&pb.ClusterSource)
	if err != nil {
		return nil, err
	}
	if clusterSourceField != nil {
		st.ClusterSource = *clusterSourceField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	dataSecurityModeField, err := identity(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}

	defaultTagsField := map[string]string{}
	for k, v := range pb.DefaultTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			defaultTagsField[k] = *item
		}
	}
	st.DefaultTags = defaultTagsField
	dockerImageField, err := dockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverField, err := sparkNodeFromPb(pb.Driver)
	if err != nil {
		return nil, err
	}
	if driverField != nil {
		st.Driver = driverField
	}
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}

	var executorsField []SparkNode
	for _, itemPb := range pb.Executors {
		item, err := sparkNodeFromPb(&itemPb)
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	isSingleNodeField, err := identity(&pb.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodeField != nil {
		st.IsSingleNode = *isSingleNodeField
	}
	jdbcPortField, err := identity(&pb.JdbcPort)
	if err != nil {
		return nil, err
	}
	if jdbcPortField != nil {
		st.JdbcPort = *jdbcPortField
	}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	lastRestartedTimeField, err := identity(&pb.LastRestartedTime)
	if err != nil {
		return nil, err
	}
	if lastRestartedTimeField != nil {
		st.LastRestartedTime = *lastRestartedTimeField
	}
	lastStateLossTimeField, err := identity(&pb.LastStateLossTime)
	if err != nil {
		return nil, err
	}
	if lastStateLossTimeField != nil {
		st.LastStateLossTime = *lastStateLossTimeField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	runtimeEngineField, err := identity(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	singleUserNameField, err := identity(&pb.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNameField != nil {
		st.SingleUserName = *singleUserNameField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField
	sparkContextIdField, err := identity(&pb.SparkContextId)
	if err != nil {
		return nil, err
	}
	if sparkContextIdField != nil {
		st.SparkContextId = *sparkContextIdField
	}

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField
	sparkVersionField, err := identity(&pb.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionField != nil {
		st.SparkVersion = *sparkVersionField
	}
	specField, err := ClusterSpecFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	stateMessageField, err := identity(&pb.StateMessage)
	if err != nil {
		return nil, err
	}
	if stateMessageField != nil {
		st.StateMessage = *stateMessageField
	}
	terminatedTimeField, err := identity(&pb.TerminatedTime)
	if err != nil {
		return nil, err
	}
	if terminatedTimeField != nil {
		st.TerminatedTime = *terminatedTimeField
	}
	terminationReasonField, err := terminationReasonFromPb(pb.TerminationReason)
	if err != nil {
		return nil, err
	}
	if terminationReasonField != nil {
		st.TerminationReason = terminationReasonField
	}
	useMlRuntimeField, err := identity(&pb.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimeField != nil {
		st.UseMlRuntime = *useMlRuntimeField
	}
	workloadTypeField, err := workloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterEvent struct {
	ClusterId string

	DataPlaneEventDetails *DataPlaneEventDetails

	Details *EventDetails
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp int64

	Type EventType

	ForceSendFields []string
}

func clusterEventToPb(st *ClusterEvent) (*clusterEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterEventPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	dataPlaneEventDetailsPb, err := dataPlaneEventDetailsToPb(st.DataPlaneEventDetails)
	if err != nil {
		return nil, err
	}
	if dataPlaneEventDetailsPb != nil {
		pb.DataPlaneEventDetails = dataPlaneEventDetailsPb
	}

	detailsPb, err := eventDetailsToPb(st.Details)
	if err != nil {
		return nil, err
	}
	if detailsPb != nil {
		pb.Details = detailsPb
	}

	timestampPb, err := identity(&st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterEventPb struct {
	ClusterId string `json:"cluster_id"`

	DataPlaneEventDetails *dataPlaneEventDetailsPb `json:"data_plane_event_details,omitempty"`

	Details *eventDetailsPb `json:"details,omitempty"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp int64 `json:"timestamp,omitempty"`

	Type EventType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterEventFromPb(pb *clusterEventPb) (*ClusterEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterEvent{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	dataPlaneEventDetailsField, err := dataPlaneEventDetailsFromPb(pb.DataPlaneEventDetails)
	if err != nil {
		return nil, err
	}
	if dataPlaneEventDetailsField != nil {
		st.DataPlaneEventDetails = dataPlaneEventDetailsField
	}
	detailsField, err := eventDetailsFromPb(pb.Details)
	if err != nil {
		return nil, err
	}
	if detailsField != nil {
		st.Details = detailsField
	}
	timestampField, err := identity(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = *timestampField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterEventPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterEventPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId string
	// Status of all libraries on the cluster.
	LibraryStatuses []LibraryFullStatus

	ForceSendFields []string
}

func clusterLibraryStatusesToPb(st *ClusterLibraryStatuses) (*clusterLibraryStatusesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterLibraryStatusesPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	var libraryStatusesPb []libraryFullStatusPb
	for _, item := range st.LibraryStatuses {
		itemPb, err := libraryFullStatusToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			libraryStatusesPb = append(libraryStatusesPb, *itemPb)
		}
	}
	pb.LibraryStatuses = libraryStatusesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterLibraryStatusesPb struct {
	// Unique identifier for the cluster.
	ClusterId string `json:"cluster_id,omitempty"`
	// Status of all libraries on the cluster.
	LibraryStatuses []libraryFullStatusPb `json:"library_statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterLibraryStatusesFromPb(pb *clusterLibraryStatusesPb) (*ClusterLibraryStatuses, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterLibraryStatuses{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	var libraryStatusesField []LibraryFullStatus
	for _, itemPb := range pb.LibraryStatuses {
		item, err := libraryFullStatusFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			libraryStatusesField = append(libraryStatusesField, *item)
		}
	}
	st.LibraryStatuses = libraryStatusesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterLibraryStatusesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterLibraryStatusesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Cluster log delivery config
type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 *S3StorageInfo
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	Volumes *VolumesStorageInfo
}

func ClusterLogConfToPb(st *ClusterLogConf) (*ClusterLogConfPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ClusterLogConfPb{}
	dbfsPb, err := dbfsStorageInfoToPb(st.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsPb != nil {
		pb.Dbfs = dbfsPb
	}

	s3Pb, err := s3StorageInfoToPb(st.S3)
	if err != nil {
		return nil, err
	}
	if s3Pb != nil {
		pb.S3 = s3Pb
	}

	volumesPb, err := volumesStorageInfoToPb(st.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesPb != nil {
		pb.Volumes = volumesPb
	}

	return pb, nil
}

func (st *ClusterLogConf) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ClusterLogConfPb{}
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

func (st ClusterLogConf) MarshalJSON() ([]byte, error) {
	pb, err := ClusterLogConfToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterLogConfPb struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *dbfsStorageInfoPb `json:"dbfs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 *s3StorageInfoPb `json:"s3,omitempty"`
	// destination needs to be provided, e.g. `{ "volumes": { "destination":
	// "/Volumes/catalog/schema/volume/cluster_log" } }`
	Volumes *volumesStorageInfoPb `json:"volumes,omitempty"`
}

func ClusterLogConfFromPb(pb *ClusterLogConfPb) (*ClusterLogConf, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterLogConf{}
	dbfsField, err := dbfsStorageInfoFromPb(pb.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsField != nil {
		st.Dbfs = dbfsField
	}
	s3Field, err := s3StorageInfoFromPb(pb.S3)
	if err != nil {
		return nil, err
	}
	if s3Field != nil {
		st.S3 = s3Field
	}
	volumesField, err := volumesStorageInfoFromPb(pb.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesField != nil {
		st.Volumes = volumesField
	}

	return st, nil
}

type ClusterPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel ClusterPermissionLevel

	ForceSendFields []string
}

func clusterPermissionToPb(st *ClusterPermission) (*clusterPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPermissionFromPb(pb *clusterPermissionPb) (*ClusterPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermission{}
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type ClusterPermissionLevel string
type clusterPermissionLevelPb string

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

// Type always returns ClusterPermissionLevel to satisfy [pflag.Value] interface
func (f *ClusterPermissionLevel) Type() string {
	return "ClusterPermissionLevel"
}

func clusterPermissionLevelToPb(st *ClusterPermissionLevel) (*clusterPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := clusterPermissionLevelPb(*st)
	return &pb, nil
}

func clusterPermissionLevelFromPb(pb *clusterPermissionLevelPb) (*ClusterPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterPermissionLevel(*pb)
	return &st, nil
}

type ClusterPermissions struct {
	AccessControlList []ClusterAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func clusterPermissionsToPb(st *ClusterPermissions) (*clusterPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionsPb{}

	var accessControlListPb []clusterAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := clusterAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPermissionsPb struct {
	AccessControlList []clusterAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPermissionsFromPb(pb *clusterPermissionsPb) (*ClusterPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissions{}

	var accessControlListField []ClusterAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := clusterAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel ClusterPermissionLevel

	ForceSendFields []string
}

func clusterPermissionsDescriptionToPb(st *ClusterPermissionsDescription) (*clusterPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel ClusterPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPermissionsDescriptionFromPb(pb *clusterPermissionsDescriptionPb) (*ClusterPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissionsDescription{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPermissionsRequest struct {
	AccessControlList []ClusterAccessControlRequest
	// The cluster for which to get or manage permissions.
	ClusterId string
}

func clusterPermissionsRequestToPb(st *ClusterPermissionsRequest) (*clusterPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPermissionsRequestPb{}

	var accessControlListPb []clusterAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := clusterAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type clusterPermissionsRequestPb struct {
	AccessControlList []clusterAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" url:"-"`
}

func clusterPermissionsRequestFromPb(pb *clusterPermissionsRequestPb) (*ClusterPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPermissionsRequest{}

	var accessControlListField []ClusterAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := clusterAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func clusterPolicyAccessControlRequestToPb(st *ClusterPolicyAccessControlRequest) (*clusterPolicyAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPolicyAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyAccessControlRequestFromPb(pb *clusterPolicyAccessControlRequestPb) (*ClusterPolicyAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyAccessControlRequest{}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyAccessControlResponse struct {
	// All permissions.
	AllPermissions []ClusterPolicyPermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func clusterPolicyAccessControlResponseToPb(st *ClusterPolicyAccessControlResponse) (*clusterPolicyAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyAccessControlResponsePb{}

	var allPermissionsPb []clusterPolicyPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := clusterPolicyPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPolicyAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []clusterPolicyPermissionPb `json:"all_permissions,omitempty"`
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

func clusterPolicyAccessControlResponseFromPb(pb *clusterPolicyAccessControlResponsePb) (*ClusterPolicyAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyAccessControlResponse{}

	var allPermissionsField []ClusterPolicyPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := clusterPolicyPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel

	ForceSendFields []string
}

func clusterPolicyPermissionToPb(st *ClusterPolicyPermission) (*clusterPolicyPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPolicyPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyPermissionFromPb(pb *clusterPolicyPermissionPb) (*ClusterPolicyPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermission{}
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type ClusterPolicyPermissionLevel string
type clusterPolicyPermissionLevelPb string

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

// Type always returns ClusterPolicyPermissionLevel to satisfy [pflag.Value] interface
func (f *ClusterPolicyPermissionLevel) Type() string {
	return "ClusterPolicyPermissionLevel"
}

func clusterPolicyPermissionLevelToPb(st *ClusterPolicyPermissionLevel) (*clusterPolicyPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := clusterPolicyPermissionLevelPb(*st)
	return &pb, nil
}

func clusterPolicyPermissionLevelFromPb(pb *clusterPolicyPermissionLevelPb) (*ClusterPolicyPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterPolicyPermissionLevel(*pb)
	return &st, nil
}

type ClusterPolicyPermissions struct {
	AccessControlList []ClusterPolicyAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func clusterPolicyPermissionsToPb(st *ClusterPolicyPermissions) (*clusterPolicyPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionsPb{}

	var accessControlListPb []clusterPolicyAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := clusterPolicyAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPolicyPermissionsPb struct {
	AccessControlList []clusterPolicyAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyPermissionsFromPb(pb *clusterPolicyPermissionsPb) (*ClusterPolicyPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissions{}

	var accessControlListField []ClusterPolicyAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := clusterPolicyAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel

	ForceSendFields []string
}

func clusterPolicyPermissionsDescriptionToPb(st *ClusterPolicyPermissionsDescription) (*clusterPolicyPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterPolicyPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterPolicyPermissionsDescriptionFromPb(pb *clusterPolicyPermissionsDescriptionPb) (*ClusterPolicyPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissionsDescription{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterPolicyPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterPolicyPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterPolicyPermissionsRequest struct {
	AccessControlList []ClusterPolicyAccessControlRequest
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string
}

func clusterPolicyPermissionsRequestToPb(st *ClusterPolicyPermissionsRequest) (*clusterPolicyPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterPolicyPermissionsRequestPb{}

	var accessControlListPb []clusterPolicyAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := clusterPolicyAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	clusterPolicyIdPb, err := identity(&st.ClusterPolicyId)
	if err != nil {
		return nil, err
	}
	if clusterPolicyIdPb != nil {
		pb.ClusterPolicyId = *clusterPolicyIdPb
	}

	return pb, nil
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

type clusterPolicyPermissionsRequestPb struct {
	AccessControlList []clusterPolicyAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" url:"-"`
}

func clusterPolicyPermissionsRequestFromPb(pb *clusterPolicyPermissionsRequestPb) (*ClusterPolicyPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterPolicyPermissionsRequest{}

	var accessControlListField []ClusterPolicyAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := clusterPolicyAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	clusterPolicyIdField, err := identity(&pb.ClusterPolicyId)
	if err != nil {
		return nil, err
	}
	if clusterPolicyIdField != nil {
		st.ClusterPolicyId = *clusterPolicyIdField
	}

	return st, nil
}

// Represents a change to the cluster settings required for the cluster to
// become compliant with its policy.
type ClusterSettingsChange struct {
	// The field where this change would be made.
	Field string
	// The new value of this field after enforcing policy compliance (either a
	// number, a boolean, or a string) converted to a string. This is intended
	// to be read by a human. The typed new value of this field can be retrieved
	// by reading the settings field in the API response.
	NewValue string
	// The previous value of this field before enforcing policy compliance
	// (either a number, a boolean, or a string) converted to a string. This is
	// intended to be read by a human. The type of the field can be retrieved by
	// reading the settings field in the API response.
	PreviousValue string

	ForceSendFields []string
}

func clusterSettingsChangeToPb(st *ClusterSettingsChange) (*clusterSettingsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSettingsChangePb{}
	fieldPb, err := identity(&st.Field)
	if err != nil {
		return nil, err
	}
	if fieldPb != nil {
		pb.Field = *fieldPb
	}

	newValuePb, err := identity(&st.NewValue)
	if err != nil {
		return nil, err
	}
	if newValuePb != nil {
		pb.NewValue = *newValuePb
	}

	previousValuePb, err := identity(&st.PreviousValue)
	if err != nil {
		return nil, err
	}
	if previousValuePb != nil {
		pb.PreviousValue = *previousValuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterSettingsChangePb struct {
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

func clusterSettingsChangeFromPb(pb *clusterSettingsChangePb) (*ClusterSettingsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSettingsChange{}
	fieldField, err := identity(&pb.Field)
	if err != nil {
		return nil, err
	}
	if fieldField != nil {
		st.Field = *fieldField
	}
	newValueField, err := identity(&pb.NewValue)
	if err != nil {
		return nil, err
	}
	if newValueField != nil {
		st.NewValue = *newValueField
	}
	previousValueField, err := identity(&pb.PreviousValue)
	if err != nil {
		return nil, err
	}
	if previousValueField != nil {
		st.PreviousValue = *previousValueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSettingsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSettingsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterSize struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
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
	NumWorkers int

	ForceSendFields []string
}

func clusterSizeToPb(st *ClusterSize) (*clusterSizePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSizePb{}
	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterSizePb struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
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

func clusterSizeFromPb(pb *clusterSizePb) (*ClusterSize, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSize{}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSizePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSizePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.
type ClusterSource string
type clusterSourcePb string

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

// Type always returns ClusterSource to satisfy [pflag.Value] interface
func (f *ClusterSource) Type() string {
	return "ClusterSource"
}

func clusterSourceToPb(st *ClusterSource) (*clusterSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := clusterSourcePb(*st)
	return &pb, nil
}

func clusterSourceFromPb(pb *clusterSourcePb) (*ClusterSource, error) {
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
	ApplyPolicyDefaultValues bool
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName string
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
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
	DataSecurityMode DataSecurityMode
	// Custom docker image BYOC
	DockerImage *DockerImage
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool
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
	Kind Kind
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string
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
	SparkEnvVars map[string]string
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *WorkloadType

	ForceSendFields []string
}

func ClusterSpecToPb(st *ClusterSpec) (*ClusterSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ClusterSpecPb{}
	applyPolicyDefaultValuesPb, err := identity(&st.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesPb != nil {
		pb.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesPb
	}

	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	autoterminationMinutesPb, err := identity(&st.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesPb != nil {
		pb.AutoterminationMinutes = *autoterminationMinutesPb
	}

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

	clusterNamePb, err := identity(&st.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNamePb != nil {
		pb.ClusterName = *clusterNamePb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	dataSecurityModePb, err := identity(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}

	dockerImagePb, err := dockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []InitScriptInfoPb
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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	isSingleNodePb, err := identity(&st.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodePb != nil {
		pb.IsSingleNode = *isSingleNodePb
	}

	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	runtimeEnginePb, err := identity(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}

	singleUserNamePb, err := identity(&st.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNamePb != nil {
		pb.SingleUserName = *singleUserNamePb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	sparkVersionPb, err := identity(&st.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionPb != nil {
		pb.SparkVersion = *sparkVersionPb
	}

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

	useMlRuntimePb, err := identity(&st.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimePb != nil {
		pb.UseMlRuntime = *useMlRuntimePb
	}

	workloadTypePb, err := workloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ClusterSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ClusterSpecPb{}
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

func (st ClusterSpec) MarshalJSON() ([]byte, error) {
	pb, err := ClusterSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterSpecPb struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributesPb `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
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
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *dockerImagePb `json:"docker_image,omitempty"`
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
	GcpAttributes *GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfoPb `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
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
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *workloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ClusterSpecFromPb(pb *ClusterSpecPb) (*ClusterSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSpec{}
	applyPolicyDefaultValuesField, err := identity(&pb.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesField != nil {
		st.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesField
	}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	autoterminationMinutesField, err := identity(&pb.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesField != nil {
		st.AutoterminationMinutes = *autoterminationMinutesField
	}
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
	clusterNameField, err := identity(&pb.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNameField != nil {
		st.ClusterName = *clusterNameField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	dataSecurityModeField, err := identity(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := dockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	isSingleNodeField, err := identity(&pb.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodeField != nil {
		st.IsSingleNode = *isSingleNodeField
	}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	runtimeEngineField, err := identity(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	singleUserNameField, err := identity(&pb.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNameField != nil {
		st.SingleUserName = *singleUserNameField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField
	sparkVersionField, err := identity(&pb.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionField != nil {
		st.SparkVersion = *sparkVersionField
	}

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField
	useMlRuntimeField, err := identity(&pb.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimeField != nil {
		st.UseMlRuntime = *useMlRuntimeField
	}
	workloadTypeField, err := workloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ClusterSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get status
type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId string
}

func clusterStatusToPb(st *ClusterStatus) (*clusterStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterStatusPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type clusterStatusPb struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId string `json:"-" url:"cluster_id"`
}

func clusterStatusFromPb(pb *clusterStatusPb) (*ClusterStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterStatus{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type Command struct {
	// Running cluster id
	ClusterId string
	// Executable code
	Command string
	// Running context id
	ContextId string

	Language Language

	ForceSendFields []string
}

func commandToPb(st *Command) (*commandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commandPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	commandPb, err := identity(&st.Command)
	if err != nil {
		return nil, err
	}
	if commandPb != nil {
		pb.Command = *commandPb
	}

	contextIdPb, err := identity(&st.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdPb != nil {
		pb.ContextId = *contextIdPb
	}

	languagePb, err := identity(&st.Language)
	if err != nil {
		return nil, err
	}
	if languagePb != nil {
		pb.Language = *languagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type commandPb struct {
	// Running cluster id
	ClusterId string `json:"clusterId,omitempty"`
	// Executable code
	Command string `json:"command,omitempty"`
	// Running context id
	ContextId string `json:"contextId,omitempty"`

	Language Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func commandFromPb(pb *commandPb) (*Command, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Command{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	commandField, err := identity(&pb.Command)
	if err != nil {
		return nil, err
	}
	if commandField != nil {
		st.Command = *commandField
	}
	contextIdField, err := identity(&pb.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdField != nil {
		st.ContextId = *contextIdField
	}
	languageField, err := identity(&pb.Language)
	if err != nil {
		return nil, err
	}
	if languageField != nil {
		st.Language = *languageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *commandPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st commandPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CommandStatus string
type commandStatusPb string

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

// Type always returns CommandStatus to satisfy [pflag.Value] interface
func (f *CommandStatus) Type() string {
	return "CommandStatus"
}

func commandStatusToPb(st *CommandStatus) (*commandStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := commandStatusPb(*st)
	return &pb, nil
}

func commandStatusFromPb(pb *commandStatusPb) (*CommandStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := CommandStatus(*pb)
	return &st, nil
}

// Get command info
type CommandStatusRequest struct {
	ClusterId string

	CommandId string

	ContextId string
}

func commandStatusRequestToPb(st *CommandStatusRequest) (*commandStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commandStatusRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	commandIdPb, err := identity(&st.CommandId)
	if err != nil {
		return nil, err
	}
	if commandIdPb != nil {
		pb.CommandId = *commandIdPb
	}

	contextIdPb, err := identity(&st.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdPb != nil {
		pb.ContextId = *contextIdPb
	}

	return pb, nil
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
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	commandIdField, err := identity(&pb.CommandId)
	if err != nil {
		return nil, err
	}
	if commandIdField != nil {
		st.CommandId = *commandIdField
	}
	contextIdField, err := identity(&pb.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdField != nil {
		st.ContextId = *contextIdField
	}

	return st, nil
}

type CommandStatusResponse struct {
	Id string

	Results *Results

	Status CommandStatus

	ForceSendFields []string
}

func commandStatusResponseToPb(st *CommandStatusResponse) (*commandStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commandStatusResponsePb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	resultsPb, err := resultsToPb(st.Results)
	if err != nil {
		return nil, err
	}
	if resultsPb != nil {
		pb.Results = resultsPb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type commandStatusResponsePb struct {
	Id string `json:"id,omitempty"`

	Results *resultsPb `json:"results,omitempty"`

	Status CommandStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func commandStatusResponseFromPb(pb *commandStatusResponsePb) (*CommandStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommandStatusResponse{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	resultsField, err := resultsFromPb(pb.Results)
	if err != nil {
		return nil, err
	}
	if resultsField != nil {
		st.Results = resultsField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *commandStatusResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st commandStatusResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ContextStatus string
type contextStatusPb string

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

// Type always returns ContextStatus to satisfy [pflag.Value] interface
func (f *ContextStatus) Type() string {
	return "ContextStatus"
}

func contextStatusToPb(st *ContextStatus) (*contextStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := contextStatusPb(*st)
	return &pb, nil
}

func contextStatusFromPb(pb *contextStatusPb) (*ContextStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ContextStatus(*pb)
	return &st, nil
}

// Get status
type ContextStatusRequest struct {
	ClusterId string

	ContextId string
}

func contextStatusRequestToPb(st *ContextStatusRequest) (*contextStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &contextStatusRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	contextIdPb, err := identity(&st.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdPb != nil {
		pb.ContextId = *contextIdPb
	}

	return pb, nil
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

type contextStatusRequestPb struct {
	ClusterId string `json:"-" url:"clusterId"`

	ContextId string `json:"-" url:"contextId"`
}

func contextStatusRequestFromPb(pb *contextStatusRequestPb) (*ContextStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContextStatusRequest{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	contextIdField, err := identity(&pb.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdField != nil {
		st.ContextId = *contextIdField
	}

	return st, nil
}

type ContextStatusResponse struct {
	Id string

	Status ContextStatus

	ForceSendFields []string
}

func contextStatusResponseToPb(st *ContextStatusResponse) (*contextStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &contextStatusResponsePb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type contextStatusResponsePb struct {
	Id string `json:"id,omitempty"`

	Status ContextStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func contextStatusResponseFromPb(pb *contextStatusResponsePb) (*ContextStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ContextStatusResponse{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *contextStatusResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st contextStatusResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues bool
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom *CloneCluster
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName string
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
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
	DataSecurityMode DataSecurityMode
	// Custom docker image BYOC
	DockerImage *DockerImage
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool
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
	Kind Kind
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string
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
	SparkEnvVars map[string]string
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *WorkloadType

	ForceSendFields []string
}

func createClusterToPb(st *CreateCluster) (*createClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createClusterPb{}
	applyPolicyDefaultValuesPb, err := identity(&st.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesPb != nil {
		pb.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesPb
	}

	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	autoterminationMinutesPb, err := identity(&st.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesPb != nil {
		pb.AutoterminationMinutes = *autoterminationMinutesPb
	}

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

	cloneFromPb, err := cloneClusterToPb(st.CloneFrom)
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

	clusterNamePb, err := identity(&st.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNamePb != nil {
		pb.ClusterName = *clusterNamePb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	dataSecurityModePb, err := identity(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}

	dockerImagePb, err := dockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []InitScriptInfoPb
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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	isSingleNodePb, err := identity(&st.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodePb != nil {
		pb.IsSingleNode = *isSingleNodePb
	}

	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	runtimeEnginePb, err := identity(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}

	singleUserNamePb, err := identity(&st.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNamePb != nil {
		pb.SingleUserName = *singleUserNamePb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	sparkVersionPb, err := identity(&st.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionPb != nil {
		pb.SparkVersion = *sparkVersionPb
	}

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

	useMlRuntimePb, err := identity(&st.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimePb != nil {
		pb.UseMlRuntime = *useMlRuntimePb
	}

	workloadTypePb, err := workloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createClusterPb struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributesPb `json:"azure_attributes,omitempty"`
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom *cloneClusterPb `json:"clone_from,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
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
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *dockerImagePb `json:"docker_image,omitempty"`
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
	GcpAttributes *GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfoPb `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
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
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *workloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createClusterFromPb(pb *createClusterPb) (*CreateCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCluster{}
	applyPolicyDefaultValuesField, err := identity(&pb.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesField != nil {
		st.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesField
	}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	autoterminationMinutesField, err := identity(&pb.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesField != nil {
		st.AutoterminationMinutes = *autoterminationMinutesField
	}
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
	cloneFromField, err := cloneClusterFromPb(pb.CloneFrom)
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
	clusterNameField, err := identity(&pb.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNameField != nil {
		st.ClusterName = *clusterNameField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	dataSecurityModeField, err := identity(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := dockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	isSingleNodeField, err := identity(&pb.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodeField != nil {
		st.IsSingleNode = *isSingleNodeField
	}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	runtimeEngineField, err := identity(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	singleUserNameField, err := identity(&pb.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNameField != nil {
		st.SingleUserName = *singleUserNameField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField
	sparkVersionField, err := identity(&pb.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionField != nil {
		st.SparkVersion = *sparkVersionField
	}

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField
	useMlRuntimeField, err := identity(&pb.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimeField != nil {
		st.UseMlRuntime = *useMlRuntimeField
	}
	workloadTypeField, err := workloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateClusterResponse struct {
	ClusterId string

	ForceSendFields []string
}

func createClusterResponseToPb(st *CreateClusterResponse) (*createClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createClusterResponsePb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createClusterResponsePb struct {
	ClusterId string `json:"cluster_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createClusterResponseFromPb(pb *createClusterResponsePb) (*CreateClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateClusterResponse{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createClusterResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createClusterResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateContext struct {
	// Running cluster id
	ClusterId string

	Language Language

	ForceSendFields []string
}

func createContextToPb(st *CreateContext) (*createContextPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createContextPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	languagePb, err := identity(&st.Language)
	if err != nil {
		return nil, err
	}
	if languagePb != nil {
		pb.Language = *languagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createContextPb struct {
	// Running cluster id
	ClusterId string `json:"clusterId,omitempty"`

	Language Language `json:"language,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createContextFromPb(pb *createContextPb) (*CreateContext, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateContext{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	languageField, err := identity(&pb.Language)
	if err != nil {
		return nil, err
	}
	if languageField != nil {
		st.Language = *languageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createContextPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createContextPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string

	ForceSendFields []string
}

func createInstancePoolToPb(st *CreateInstancePool) (*createInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createInstancePoolPb{}
	awsAttributesPb, err := instancePoolAwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}

	azureAttributesPb, err := instancePoolAzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	diskSpecPb, err := diskSpecToPb(st.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecPb != nil {
		pb.DiskSpec = diskSpecPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	gcpAttributesPb, err := instancePoolGcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	idleInstanceAutoterminationMinutesPb, err := identity(&st.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesPb != nil {
		pb.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesPb
	}

	instancePoolNamePb, err := identity(&st.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNamePb != nil {
		pb.InstancePoolName = *instancePoolNamePb
	}

	maxCapacityPb, err := identity(&st.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityPb != nil {
		pb.MaxCapacity = *maxCapacityPb
	}

	minIdleInstancesPb, err := identity(&st.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesPb != nil {
		pb.MinIdleInstances = *minIdleInstancesPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	var preloadedDockerImagesPb []dockerImagePb
	for _, item := range st.PreloadedDockerImages {
		itemPb, err := dockerImageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedDockerImagesPb = append(preloadedDockerImagesPb, *itemPb)
		}
	}
	pb.PreloadedDockerImages = preloadedDockerImagesPb

	var preloadedSparkVersionsPb []string
	for _, item := range st.PreloadedSparkVersions {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedSparkVersionsPb = append(preloadedSparkVersionsPb, *itemPb)
		}
	}
	pb.PreloadedSparkVersions = preloadedSparkVersionsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createInstancePoolPb struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *instancePoolAwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *instancePoolAzureAttributesPb `json:"azure_attributes,omitempty"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string `json:"custom_tags,omitempty"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *diskSpecPb `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *instancePoolGcpAttributesPb `json:"gcp_attributes,omitempty"`
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
	PreloadedDockerImages []dockerImagePb `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createInstancePoolFromPb(pb *createInstancePoolPb) (*CreateInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstancePool{}
	awsAttributesField, err := instancePoolAwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := instancePoolAzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	diskSpecField, err := diskSpecFromPb(pb.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecField != nil {
		st.DiskSpec = diskSpecField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	gcpAttributesField, err := instancePoolGcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}
	idleInstanceAutoterminationMinutesField, err := identity(&pb.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesField != nil {
		st.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesField
	}
	instancePoolNameField, err := identity(&pb.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNameField != nil {
		st.InstancePoolName = *instancePoolNameField
	}
	maxCapacityField, err := identity(&pb.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityField != nil {
		st.MaxCapacity = *maxCapacityField
	}
	minIdleInstancesField, err := identity(&pb.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesField != nil {
		st.MinIdleInstances = *minIdleInstancesField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}

	var preloadedDockerImagesField []DockerImage
	for _, itemPb := range pb.PreloadedDockerImages {
		item, err := dockerImageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedDockerImagesField = append(preloadedDockerImagesField, *item)
		}
	}
	st.PreloadedDockerImages = preloadedDockerImagesField

	var preloadedSparkVersionsField []string
	for _, itemPb := range pb.PreloadedSparkVersions {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedSparkVersionsField = append(preloadedSparkVersionsField, *item)
		}
	}
	st.PreloadedSparkVersions = preloadedSparkVersionsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId string

	ForceSendFields []string
}

func createInstancePoolResponseToPb(st *CreateInstancePoolResponse) (*createInstancePoolResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createInstancePoolResponsePb{}
	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createInstancePoolResponsePb struct {
	// The ID of the created instance pool.
	InstancePoolId string `json:"instance_pool_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createInstancePoolResponseFromPb(pb *createInstancePoolResponsePb) (*CreateInstancePoolResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateInstancePoolResponse{}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createInstancePoolResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createInstancePoolResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string
	// Additional human-readable description of the cluster policy.
	Description string
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides string
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string

	ForceSendFields []string
}

func createPolicyToPb(st *CreatePolicy) (*createPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPolicyPb{}
	definitionPb, err := identity(&st.Definition)
	if err != nil {
		return nil, err
	}
	if definitionPb != nil {
		pb.Definition = *definitionPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	var librariesPb []LibraryPb
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

	maxClustersPerUserPb, err := identity(&st.MaxClustersPerUser)
	if err != nil {
		return nil, err
	}
	if maxClustersPerUserPb != nil {
		pb.MaxClustersPerUser = *maxClustersPerUserPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	policyFamilyDefinitionOverridesPb, err := identity(&st.PolicyFamilyDefinitionOverrides)
	if err != nil {
		return nil, err
	}
	if policyFamilyDefinitionOverridesPb != nil {
		pb.PolicyFamilyDefinitionOverrides = *policyFamilyDefinitionOverridesPb
	}

	policyFamilyIdPb, err := identity(&st.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdPb != nil {
		pb.PolicyFamilyId = *policyFamilyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createPolicyPb struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []LibraryPb `json:"libraries,omitempty"`
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

func createPolicyFromPb(pb *createPolicyPb) (*CreatePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePolicy{}
	definitionField, err := identity(&pb.Definition)
	if err != nil {
		return nil, err
	}
	if definitionField != nil {
		st.Definition = *definitionField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}

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
	maxClustersPerUserField, err := identity(&pb.MaxClustersPerUser)
	if err != nil {
		return nil, err
	}
	if maxClustersPerUserField != nil {
		st.MaxClustersPerUser = *maxClustersPerUserField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	policyFamilyDefinitionOverridesField, err := identity(&pb.PolicyFamilyDefinitionOverrides)
	if err != nil {
		return nil, err
	}
	if policyFamilyDefinitionOverridesField != nil {
		st.PolicyFamilyDefinitionOverrides = *policyFamilyDefinitionOverridesField
	}
	policyFamilyIdField, err := identity(&pb.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdField != nil {
		st.PolicyFamilyId = *policyFamilyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId string

	ForceSendFields []string
}

func createPolicyResponseToPb(st *CreatePolicyResponse) (*createPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPolicyResponsePb{}
	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createPolicyResponsePb struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId string `json:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPolicyResponseFromPb(pb *createPolicyResponsePb) (*CreatePolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePolicyResponse{}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateResponse struct {
	// The global init script ID.
	ScriptId string

	ForceSendFields []string
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}
	scriptIdPb, err := identity(&st.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdPb != nil {
		pb.ScriptId = *scriptIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createResponsePb struct {
	// The global init script ID.
	ScriptId string `json:"script_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	scriptIdField, err := identity(&pb.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdField != nil {
		st.ScriptId = *scriptIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Created struct {
	Id string

	ForceSendFields []string
}

func createdToPb(st *Created) (*createdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createdPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createdPb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createdFromPb(pb *createdPb) (*Created, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Created{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomPolicyTag struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	//
	// - Follows the regex pattern defined in
	// cluster-common/conf/src/ClusterTagConstraints.scala
	// (https://src.dev.databricks.com/databricks/universe@1647196627c8dc7b4152ad098a94b86484b93a6c/-/blob/cluster-common/conf/src/ClusterTagConstraints.scala?L17)
	Key string
	// The value of the tag.
	//
	// - Follows the regex pattern defined in
	// cluster-common/conf/src/ClusterTagConstraints.scala
	// (https://src.dev.databricks.com/databricks/universe@1647196627c8dc7b4152ad098a94b86484b93a6c/-/blob/cluster-common/conf/src/ClusterTagConstraints.scala?L24)
	Value string

	ForceSendFields []string
}

func CustomPolicyTagToPb(st *CustomPolicyTag) (*CustomPolicyTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &CustomPolicyTagPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CustomPolicyTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &CustomPolicyTagPb{}
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

func (st CustomPolicyTag) MarshalJSON() ([]byte, error) {
	pb, err := CustomPolicyTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CustomPolicyTagPb struct {
	// The key of the tag. - Must be unique among all custom tags of the same
	// policy - Cannot be “budget-policy-name”, “budget-policy-id” or
	// "budget-policy-resolution-result" - these tags are preserved.
	//
	// - Follows the regex pattern defined in
	// cluster-common/conf/src/ClusterTagConstraints.scala
	// (https://src.dev.databricks.com/databricks/universe@1647196627c8dc7b4152ad098a94b86484b93a6c/-/blob/cluster-common/conf/src/ClusterTagConstraints.scala?L17)
	Key string `json:"key"`
	// The value of the tag.
	//
	// - Follows the regex pattern defined in
	// cluster-common/conf/src/ClusterTagConstraints.scala
	// (https://src.dev.databricks.com/databricks/universe@1647196627c8dc7b4152ad098a94b86484b93a6c/-/blob/cluster-common/conf/src/ClusterTagConstraints.scala?L24)
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func CustomPolicyTagFromPb(pb *CustomPolicyTagPb) (*CustomPolicyTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomPolicyTag{}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *CustomPolicyTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomPolicyTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataPlaneEventDetails struct {
	EventType DataPlaneEventDetailsEventType

	ExecutorFailures int

	HostId string

	Timestamp int64

	ForceSendFields []string
}

func dataPlaneEventDetailsToPb(st *DataPlaneEventDetails) (*dataPlaneEventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneEventDetailsPb{}
	eventTypePb, err := identity(&st.EventType)
	if err != nil {
		return nil, err
	}
	if eventTypePb != nil {
		pb.EventType = *eventTypePb
	}

	executorFailuresPb, err := identity(&st.ExecutorFailures)
	if err != nil {
		return nil, err
	}
	if executorFailuresPb != nil {
		pb.ExecutorFailures = *executorFailuresPb
	}

	hostIdPb, err := identity(&st.HostId)
	if err != nil {
		return nil, err
	}
	if hostIdPb != nil {
		pb.HostId = *hostIdPb
	}

	timestampPb, err := identity(&st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dataPlaneEventDetailsPb struct {
	EventType DataPlaneEventDetailsEventType `json:"event_type,omitempty"`

	ExecutorFailures int `json:"executor_failures,omitempty"`

	HostId string `json:"host_id,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneEventDetailsFromPb(pb *dataPlaneEventDetailsPb) (*DataPlaneEventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneEventDetails{}
	eventTypeField, err := identity(&pb.EventType)
	if err != nil {
		return nil, err
	}
	if eventTypeField != nil {
		st.EventType = *eventTypeField
	}
	executorFailuresField, err := identity(&pb.ExecutorFailures)
	if err != nil {
		return nil, err
	}
	if executorFailuresField != nil {
		st.ExecutorFailures = *executorFailuresField
	}
	hostIdField, err := identity(&pb.HostId)
	if err != nil {
		return nil, err
	}
	if hostIdField != nil {
		st.HostId = *hostIdField
	}
	timestampField, err := identity(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = *timestampField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataPlaneEventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataPlaneEventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataPlaneEventDetailsEventType string
type dataPlaneEventDetailsEventTypePb string

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

// Type always returns DataPlaneEventDetailsEventType to satisfy [pflag.Value] interface
func (f *DataPlaneEventDetailsEventType) Type() string {
	return "DataPlaneEventDetailsEventType"
}

func dataPlaneEventDetailsEventTypeToPb(st *DataPlaneEventDetailsEventType) (*dataPlaneEventDetailsEventTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dataPlaneEventDetailsEventTypePb(*st)
	return &pb, nil
}

func dataPlaneEventDetailsEventTypeFromPb(pb *dataPlaneEventDetailsEventTypePb) (*DataPlaneEventDetailsEventType, error) {
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
type dataSecurityModePb string

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

// Type always returns DataSecurityMode to satisfy [pflag.Value] interface
func (f *DataSecurityMode) Type() string {
	return "DataSecurityMode"
}

func dataSecurityModeToPb(st *DataSecurityMode) (*dataSecurityModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dataSecurityModePb(*st)
	return &pb, nil
}

func dataSecurityModeFromPb(pb *dataSecurityModePb) (*DataSecurityMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := DataSecurityMode(*pb)
	return &st, nil
}

// A storage location in DBFS
type DbfsStorageInfo struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination string
}

func dbfsStorageInfoToPb(st *DbfsStorageInfo) (*dbfsStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dbfsStorageInfoPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	return pb, nil
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

type dbfsStorageInfoPb struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination string `json:"destination"`
}

func dbfsStorageInfoFromPb(pb *dbfsStorageInfoPb) (*DbfsStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbfsStorageInfo{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}

	return st, nil
}

type DeleteCluster struct {
	// The cluster to be terminated.
	ClusterId string
}

func deleteClusterToPb(st *DeleteCluster) (*deleteClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteClusterPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type deleteClusterPb struct {
	// The cluster to be terminated.
	ClusterId string `json:"cluster_id"`
}

func deleteClusterFromPb(pb *deleteClusterPb) (*DeleteCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCluster{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type DeleteClusterResponse struct {
}

func deleteClusterResponseToPb(st *DeleteClusterResponse) (*deleteClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteClusterResponsePb{}

	return pb, nil
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

type deleteClusterResponsePb struct {
}

func deleteClusterResponseFromPb(pb *deleteClusterResponsePb) (*DeleteClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteClusterResponse{}

	return st, nil
}

// Delete init script
type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId string
}

func deleteGlobalInitScriptRequestToPb(st *DeleteGlobalInitScriptRequest) (*deleteGlobalInitScriptRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteGlobalInitScriptRequestPb{}
	scriptIdPb, err := identity(&st.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdPb != nil {
		pb.ScriptId = *scriptIdPb
	}

	return pb, nil
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

type deleteGlobalInitScriptRequestPb struct {
	// The ID of the global init script.
	ScriptId string `json:"-" url:"-"`
}

func deleteGlobalInitScriptRequestFromPb(pb *deleteGlobalInitScriptRequestPb) (*DeleteGlobalInitScriptRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteGlobalInitScriptRequest{}
	scriptIdField, err := identity(&pb.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdField != nil {
		st.ScriptId = *scriptIdField
	}

	return st, nil
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	InstancePoolId string
}

func deleteInstancePoolToPb(st *DeleteInstancePool) (*deleteInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteInstancePoolPb{}
	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	return pb, nil
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

type deleteInstancePoolPb struct {
	// The instance pool to be terminated.
	InstancePoolId string `json:"instance_pool_id"`
}

func deleteInstancePoolFromPb(pb *deleteInstancePoolPb) (*DeleteInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstancePool{}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}

	return st, nil
}

type DeleteInstancePoolResponse struct {
}

func deleteInstancePoolResponseToPb(st *DeleteInstancePoolResponse) (*deleteInstancePoolResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteInstancePoolResponsePb{}

	return pb, nil
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

type deleteInstancePoolResponsePb struct {
}

func deleteInstancePoolResponseFromPb(pb *deleteInstancePoolResponsePb) (*DeleteInstancePoolResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteInstancePoolResponse{}

	return st, nil
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId string
}

func deletePolicyToPb(st *DeletePolicy) (*deletePolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePolicyPb{}
	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	return pb, nil
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

type deletePolicyPb struct {
	// The ID of the policy to delete.
	PolicyId string `json:"policy_id"`
}

func deletePolicyFromPb(pb *deletePolicyPb) (*DeletePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePolicy{}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	return st, nil
}

type DeletePolicyResponse struct {
}

func deletePolicyResponseToPb(st *DeletePolicyResponse) (*deletePolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePolicyResponsePb{}

	return pb, nil
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

type deletePolicyResponsePb struct {
}

func deletePolicyResponseFromPb(pb *deletePolicyResponsePb) (*DeletePolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePolicyResponse{}

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

type DestroyContext struct {
	ClusterId string

	ContextId string
}

func destroyContextToPb(st *DestroyContext) (*destroyContextPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &destroyContextPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	contextIdPb, err := identity(&st.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdPb != nil {
		pb.ContextId = *contextIdPb
	}

	return pb, nil
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

type destroyContextPb struct {
	ClusterId string `json:"clusterId"`

	ContextId string `json:"contextId"`
}

func destroyContextFromPb(pb *destroyContextPb) (*DestroyContext, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DestroyContext{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	contextIdField, err := identity(&pb.ContextId)
	if err != nil {
		return nil, err
	}
	if contextIdField != nil {
		st.ContextId = *contextIdField
	}

	return st, nil
}

type DestroyResponse struct {
}

func destroyResponseToPb(st *DestroyResponse) (*destroyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &destroyResponsePb{}

	return pb, nil
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

type destroyResponsePb struct {
}

func destroyResponseFromPb(pb *destroyResponsePb) (*DestroyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DestroyResponse{}

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
	DiskCount int

	DiskIops int
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	DiskSize int

	DiskThroughput int
	// The type of disks that will be launched with this cluster.
	DiskType *DiskType

	ForceSendFields []string
}

func diskSpecToPb(st *DiskSpec) (*diskSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &diskSpecPb{}
	diskCountPb, err := identity(&st.DiskCount)
	if err != nil {
		return nil, err
	}
	if diskCountPb != nil {
		pb.DiskCount = *diskCountPb
	}

	diskIopsPb, err := identity(&st.DiskIops)
	if err != nil {
		return nil, err
	}
	if diskIopsPb != nil {
		pb.DiskIops = *diskIopsPb
	}

	diskSizePb, err := identity(&st.DiskSize)
	if err != nil {
		return nil, err
	}
	if diskSizePb != nil {
		pb.DiskSize = *diskSizePb
	}

	diskThroughputPb, err := identity(&st.DiskThroughput)
	if err != nil {
		return nil, err
	}
	if diskThroughputPb != nil {
		pb.DiskThroughput = *diskThroughputPb
	}

	diskTypePb, err := diskTypeToPb(st.DiskType)
	if err != nil {
		return nil, err
	}
	if diskTypePb != nil {
		pb.DiskType = diskTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type diskSpecPb struct {
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
	DiskType *diskTypePb `json:"disk_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func diskSpecFromPb(pb *diskSpecPb) (*DiskSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DiskSpec{}
	diskCountField, err := identity(&pb.DiskCount)
	if err != nil {
		return nil, err
	}
	if diskCountField != nil {
		st.DiskCount = *diskCountField
	}
	diskIopsField, err := identity(&pb.DiskIops)
	if err != nil {
		return nil, err
	}
	if diskIopsField != nil {
		st.DiskIops = *diskIopsField
	}
	diskSizeField, err := identity(&pb.DiskSize)
	if err != nil {
		return nil, err
	}
	if diskSizeField != nil {
		st.DiskSize = *diskSizeField
	}
	diskThroughputField, err := identity(&pb.DiskThroughput)
	if err != nil {
		return nil, err
	}
	if diskThroughputField != nil {
		st.DiskThroughput = *diskThroughputField
	}
	diskTypeField, err := diskTypeFromPb(pb.DiskType)
	if err != nil {
		return nil, err
	}
	if diskTypeField != nil {
		st.DiskType = diskTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *diskSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st diskSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Describes the disk type.
type DiskType struct {
	// All Azure Disk types that Databricks supports. See
	// https://docs.microsoft.com/en-us/azure/storage/storage-about-disks-and-vhds-linux#types-of-disks
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType
	// All EBS volume types that Databricks supports. See
	// https://aws.amazon.com/ebs/details/ for details.
	EbsVolumeType DiskTypeEbsVolumeType
}

func diskTypeToPb(st *DiskType) (*diskTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &diskTypePb{}
	azureDiskVolumeTypePb, err := identity(&st.AzureDiskVolumeType)
	if err != nil {
		return nil, err
	}
	if azureDiskVolumeTypePb != nil {
		pb.AzureDiskVolumeType = *azureDiskVolumeTypePb
	}

	ebsVolumeTypePb, err := identity(&st.EbsVolumeType)
	if err != nil {
		return nil, err
	}
	if ebsVolumeTypePb != nil {
		pb.EbsVolumeType = *ebsVolumeTypePb
	}

	return pb, nil
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

type diskTypePb struct {
	// All Azure Disk types that Databricks supports. See
	// https://docs.microsoft.com/en-us/azure/storage/storage-about-disks-and-vhds-linux#types-of-disks
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType `json:"azure_disk_volume_type,omitempty"`
	// All EBS volume types that Databricks supports. See
	// https://aws.amazon.com/ebs/details/ for details.
	EbsVolumeType DiskTypeEbsVolumeType `json:"ebs_volume_type,omitempty"`
}

func diskTypeFromPb(pb *diskTypePb) (*DiskType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DiskType{}
	azureDiskVolumeTypeField, err := identity(&pb.AzureDiskVolumeType)
	if err != nil {
		return nil, err
	}
	if azureDiskVolumeTypeField != nil {
		st.AzureDiskVolumeType = *azureDiskVolumeTypeField
	}
	ebsVolumeTypeField, err := identity(&pb.EbsVolumeType)
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
type diskTypeAzureDiskVolumeTypePb string

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

// Type always returns DiskTypeAzureDiskVolumeType to satisfy [pflag.Value] interface
func (f *DiskTypeAzureDiskVolumeType) Type() string {
	return "DiskTypeAzureDiskVolumeType"
}

func diskTypeAzureDiskVolumeTypeToPb(st *DiskTypeAzureDiskVolumeType) (*diskTypeAzureDiskVolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := diskTypeAzureDiskVolumeTypePb(*st)
	return &pb, nil
}

func diskTypeAzureDiskVolumeTypeFromPb(pb *diskTypeAzureDiskVolumeTypePb) (*DiskTypeAzureDiskVolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DiskTypeAzureDiskVolumeType(*pb)
	return &st, nil
}

// All EBS volume types that Databricks supports. See
// https://aws.amazon.com/ebs/details/ for details.
type DiskTypeEbsVolumeType string
type diskTypeEbsVolumeTypePb string

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

// Type always returns DiskTypeEbsVolumeType to satisfy [pflag.Value] interface
func (f *DiskTypeEbsVolumeType) Type() string {
	return "DiskTypeEbsVolumeType"
}

func diskTypeEbsVolumeTypeToPb(st *DiskTypeEbsVolumeType) (*diskTypeEbsVolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := diskTypeEbsVolumeTypePb(*st)
	return &pb, nil
}

func diskTypeEbsVolumeTypeFromPb(pb *diskTypeEbsVolumeTypePb) (*DiskTypeEbsVolumeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DiskTypeEbsVolumeType(*pb)
	return &st, nil
}

type DockerBasicAuth struct {
	// Password of the user
	Password string
	// Name of the user
	Username string

	ForceSendFields []string
}

func dockerBasicAuthToPb(st *DockerBasicAuth) (*dockerBasicAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dockerBasicAuthPb{}
	passwordPb, err := identity(&st.Password)
	if err != nil {
		return nil, err
	}
	if passwordPb != nil {
		pb.Password = *passwordPb
	}

	usernamePb, err := identity(&st.Username)
	if err != nil {
		return nil, err
	}
	if usernamePb != nil {
		pb.Username = *usernamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dockerBasicAuthPb struct {
	// Password of the user
	Password string `json:"password,omitempty"`
	// Name of the user
	Username string `json:"username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dockerBasicAuthFromPb(pb *dockerBasicAuthPb) (*DockerBasicAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DockerBasicAuth{}
	passwordField, err := identity(&pb.Password)
	if err != nil {
		return nil, err
	}
	if passwordField != nil {
		st.Password = *passwordField
	}
	usernameField, err := identity(&pb.Username)
	if err != nil {
		return nil, err
	}
	if usernameField != nil {
		st.Username = *usernameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dockerBasicAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dockerBasicAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DockerImage struct {
	// Basic auth with username and password
	BasicAuth *DockerBasicAuth
	// URL of the docker image.
	Url string

	ForceSendFields []string
}

func dockerImageToPb(st *DockerImage) (*dockerImagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dockerImagePb{}
	basicAuthPb, err := dockerBasicAuthToPb(st.BasicAuth)
	if err != nil {
		return nil, err
	}
	if basicAuthPb != nil {
		pb.BasicAuth = basicAuthPb
	}

	urlPb, err := identity(&st.Url)
	if err != nil {
		return nil, err
	}
	if urlPb != nil {
		pb.Url = *urlPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dockerImagePb struct {
	// Basic auth with username and password
	BasicAuth *dockerBasicAuthPb `json:"basic_auth,omitempty"`
	// URL of the docker image.
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dockerImageFromPb(pb *dockerImagePb) (*DockerImage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DockerImage{}
	basicAuthField, err := dockerBasicAuthFromPb(pb.BasicAuth)
	if err != nil {
		return nil, err
	}
	if basicAuthField != nil {
		st.BasicAuth = basicAuthField
	}
	urlField, err := identity(&pb.Url)
	if err != nil {
		return nil, err
	}
	if urlField != nil {
		st.Url = *urlField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dockerImagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dockerImagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// All EBS volume types that Databricks supports. See
// https://aws.amazon.com/ebs/details/ for details.
type EbsVolumeType string
type ebsVolumeTypePb string

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

// Type always returns EbsVolumeType to satisfy [pflag.Value] interface
func (f *EbsVolumeType) Type() string {
	return "EbsVolumeType"
}

func ebsVolumeTypeToPb(st *EbsVolumeType) (*ebsVolumeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := ebsVolumeTypePb(*st)
	return &pb, nil
}

func ebsVolumeTypeFromPb(pb *ebsVolumeTypePb) (*EbsVolumeType, error) {
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
	ApplyPolicyDefaultValues bool
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes
	// ID of the cluster
	ClusterId string
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName string
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
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
	DataSecurityMode DataSecurityMode
	// Custom docker image BYOC
	DockerImage *DockerImage
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool
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
	Kind Kind
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string
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
	SparkEnvVars map[string]string
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *WorkloadType

	ForceSendFields []string
}

func editClusterToPb(st *EditCluster) (*editClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editClusterPb{}
	applyPolicyDefaultValuesPb, err := identity(&st.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesPb != nil {
		pb.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesPb
	}

	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	autoterminationMinutesPb, err := identity(&st.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesPb != nil {
		pb.AutoterminationMinutes = *autoterminationMinutesPb
	}

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

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	clusterLogConfPb, err := ClusterLogConfToPb(st.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfPb != nil {
		pb.ClusterLogConf = clusterLogConfPb
	}

	clusterNamePb, err := identity(&st.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNamePb != nil {
		pb.ClusterName = *clusterNamePb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	dataSecurityModePb, err := identity(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}

	dockerImagePb, err := dockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []InitScriptInfoPb
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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	isSingleNodePb, err := identity(&st.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodePb != nil {
		pb.IsSingleNode = *isSingleNodePb
	}

	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	runtimeEnginePb, err := identity(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}

	singleUserNamePb, err := identity(&st.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNamePb != nil {
		pb.SingleUserName = *singleUserNamePb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	sparkVersionPb, err := identity(&st.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionPb != nil {
		pb.SparkVersion = *sparkVersionPb
	}

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

	useMlRuntimePb, err := identity(&st.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimePb != nil {
		pb.UseMlRuntime = *useMlRuntimePb
	}

	workloadTypePb, err := workloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type editClusterPb struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues bool `json:"apply_policy_default_values,omitempty"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributesPb `json:"azure_attributes,omitempty"`
	// ID of the cluster
	ClusterId string `json:"cluster_id"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
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
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *dockerImagePb `json:"docker_image,omitempty"`
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
	GcpAttributes *GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfoPb `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
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
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *workloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editClusterFromPb(pb *editClusterPb) (*EditCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditCluster{}
	applyPolicyDefaultValuesField, err := identity(&pb.ApplyPolicyDefaultValues)
	if err != nil {
		return nil, err
	}
	if applyPolicyDefaultValuesField != nil {
		st.ApplyPolicyDefaultValues = *applyPolicyDefaultValuesField
	}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	autoterminationMinutesField, err := identity(&pb.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesField != nil {
		st.AutoterminationMinutes = *autoterminationMinutesField
	}
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
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	clusterLogConfField, err := ClusterLogConfFromPb(pb.ClusterLogConf)
	if err != nil {
		return nil, err
	}
	if clusterLogConfField != nil {
		st.ClusterLogConf = clusterLogConfField
	}
	clusterNameField, err := identity(&pb.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNameField != nil {
		st.ClusterName = *clusterNameField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	dataSecurityModeField, err := identity(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := dockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	isSingleNodeField, err := identity(&pb.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodeField != nil {
		st.IsSingleNode = *isSingleNodeField
	}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	runtimeEngineField, err := identity(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	singleUserNameField, err := identity(&pb.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNameField != nil {
		st.SingleUserName = *singleUserNameField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField
	sparkVersionField, err := identity(&pb.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionField != nil {
		st.SparkVersion = *sparkVersionField
	}

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField
	useMlRuntimeField, err := identity(&pb.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimeField != nil {
		st.UseMlRuntime = *useMlRuntimeField
	}
	workloadTypeField, err := workloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditClusterResponse struct {
}

func editClusterResponseToPb(st *EditClusterResponse) (*editClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editClusterResponsePb{}

	return pb, nil
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

type editClusterResponsePb struct {
}

func editClusterResponseFromPb(pb *editClusterResponsePb) (*EditClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditClusterResponse{}

	return st, nil
}

type EditInstancePool struct {
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int
	// Instance pool ID
	InstancePoolId string
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int
	// For Fleet-pool V2, this object contains the information about the
	// alternate node type ids to use when attempting to launch a cluster if the
	// node type id is not available.
	NodeTypeFlexibility *NodeTypeFlexibility
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string

	ForceSendFields []string
}

func editInstancePoolToPb(st *EditInstancePool) (*editInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editInstancePoolPb{}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	idleInstanceAutoterminationMinutesPb, err := identity(&st.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesPb != nil {
		pb.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesPb
	}

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	instancePoolNamePb, err := identity(&st.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNamePb != nil {
		pb.InstancePoolName = *instancePoolNamePb
	}

	maxCapacityPb, err := identity(&st.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityPb != nil {
		pb.MaxCapacity = *maxCapacityPb
	}

	minIdleInstancesPb, err := identity(&st.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesPb != nil {
		pb.MinIdleInstances = *minIdleInstancesPb
	}

	nodeTypeFlexibilityPb, err := nodeTypeFlexibilityToPb(st.NodeTypeFlexibility)
	if err != nil {
		return nil, err
	}
	if nodeTypeFlexibilityPb != nil {
		pb.NodeTypeFlexibility = nodeTypeFlexibilityPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type editInstancePoolPb struct {
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
	// For Fleet-pool V2, this object contains the information about the
	// alternate node type ids to use when attempting to launch a cluster if the
	// node type id is not available.
	NodeTypeFlexibility *nodeTypeFlexibilityPb `json:"node_type_flexibility,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editInstancePoolFromPb(pb *editInstancePoolPb) (*EditInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditInstancePool{}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	idleInstanceAutoterminationMinutesField, err := identity(&pb.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesField != nil {
		st.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesField
	}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	instancePoolNameField, err := identity(&pb.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNameField != nil {
		st.InstancePoolName = *instancePoolNameField
	}
	maxCapacityField, err := identity(&pb.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityField != nil {
		st.MaxCapacity = *maxCapacityField
	}
	minIdleInstancesField, err := identity(&pb.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesField != nil {
		st.MinIdleInstances = *minIdleInstancesField
	}
	nodeTypeFlexibilityField, err := nodeTypeFlexibilityFromPb(pb.NodeTypeFlexibility)
	if err != nil {
		return nil, err
	}
	if nodeTypeFlexibilityField != nil {
		st.NodeTypeFlexibility = nodeTypeFlexibilityField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditInstancePoolResponse struct {
}

func editInstancePoolResponseToPb(st *EditInstancePoolResponse) (*editInstancePoolResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editInstancePoolResponsePb{}

	return pb, nil
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

type editInstancePoolResponsePb struct {
}

func editInstancePoolResponseFromPb(pb *editInstancePoolResponsePb) (*EditInstancePoolResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditInstancePoolResponse{}

	return st, nil
}

type EditPolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string
	// Additional human-readable description of the cluster policy.
	Description string
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides string
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string
	// The ID of the policy to update.
	PolicyId string

	ForceSendFields []string
}

func editPolicyToPb(st *EditPolicy) (*editPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPolicyPb{}
	definitionPb, err := identity(&st.Definition)
	if err != nil {
		return nil, err
	}
	if definitionPb != nil {
		pb.Definition = *definitionPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	var librariesPb []LibraryPb
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

	maxClustersPerUserPb, err := identity(&st.MaxClustersPerUser)
	if err != nil {
		return nil, err
	}
	if maxClustersPerUserPb != nil {
		pb.MaxClustersPerUser = *maxClustersPerUserPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	policyFamilyDefinitionOverridesPb, err := identity(&st.PolicyFamilyDefinitionOverrides)
	if err != nil {
		return nil, err
	}
	if policyFamilyDefinitionOverridesPb != nil {
		pb.PolicyFamilyDefinitionOverrides = *policyFamilyDefinitionOverridesPb
	}

	policyFamilyIdPb, err := identity(&st.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdPb != nil {
		pb.PolicyFamilyId = *policyFamilyIdPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type editPolicyPb struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string `json:"definition,omitempty"`
	// Additional human-readable description of the cluster policy.
	Description string `json:"description,omitempty"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []LibraryPb `json:"libraries,omitempty"`
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

func editPolicyFromPb(pb *editPolicyPb) (*EditPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPolicy{}
	definitionField, err := identity(&pb.Definition)
	if err != nil {
		return nil, err
	}
	if definitionField != nil {
		st.Definition = *definitionField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}

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
	maxClustersPerUserField, err := identity(&pb.MaxClustersPerUser)
	if err != nil {
		return nil, err
	}
	if maxClustersPerUserField != nil {
		st.MaxClustersPerUser = *maxClustersPerUserField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	policyFamilyDefinitionOverridesField, err := identity(&pb.PolicyFamilyDefinitionOverrides)
	if err != nil {
		return nil, err
	}
	if policyFamilyDefinitionOverridesField != nil {
		st.PolicyFamilyDefinitionOverrides = *policyFamilyDefinitionOverridesField
	}
	policyFamilyIdField, err := identity(&pb.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdField != nil {
		st.PolicyFamilyId = *policyFamilyIdField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditPolicyResponse struct {
}

func editPolicyResponseToPb(st *EditPolicyResponse) (*editPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPolicyResponsePb{}

	return pb, nil
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

type editPolicyResponsePb struct {
}

func editPolicyResponseFromPb(pb *editPolicyResponsePb) (*EditPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPolicyResponse{}

	return st, nil
}

type EditResponse struct {
}

func editResponseToPb(st *EditResponse) (*editResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editResponsePb{}

	return pb, nil
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

type editResponsePb struct {
}

func editResponseFromPb(pb *editResponsePb) (*EditResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditResponse{}

	return st, nil
}

type EnforceClusterComplianceRequest struct {
	// The ID of the cluster you want to enforce policy compliance on.
	ClusterId string
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	ValidateOnly bool

	ForceSendFields []string
}

func enforceClusterComplianceRequestToPb(st *EnforceClusterComplianceRequest) (*enforceClusterComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforceClusterComplianceRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	validateOnlyPb, err := identity(&st.ValidateOnly)
	if err != nil {
		return nil, err
	}
	if validateOnlyPb != nil {
		pb.ValidateOnly = *validateOnlyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enforceClusterComplianceRequestPb struct {
	// The ID of the cluster you want to enforce policy compliance on.
	ClusterId string `json:"cluster_id"`
	// If set, previews the changes that would be made to a cluster to enforce
	// compliance but does not update the cluster.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforceClusterComplianceRequestFromPb(pb *enforceClusterComplianceRequestPb) (*EnforceClusterComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforceClusterComplianceRequest{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	validateOnlyField, err := identity(&pb.ValidateOnly)
	if err != nil {
		return nil, err
	}
	if validateOnlyField != nil {
		st.ValidateOnly = *validateOnlyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforceClusterComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforceClusterComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnforceClusterComplianceResponse struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	Changes []ClusterSettingsChange
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	HasChanges bool

	ForceSendFields []string
}

func enforceClusterComplianceResponseToPb(st *EnforceClusterComplianceResponse) (*enforceClusterComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforceClusterComplianceResponsePb{}

	var changesPb []clusterSettingsChangePb
	for _, item := range st.Changes {
		itemPb, err := clusterSettingsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			changesPb = append(changesPb, *itemPb)
		}
	}
	pb.Changes = changesPb

	hasChangesPb, err := identity(&st.HasChanges)
	if err != nil {
		return nil, err
	}
	if hasChangesPb != nil {
		pb.HasChanges = *hasChangesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enforceClusterComplianceResponsePb struct {
	// A list of changes that have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	Changes []clusterSettingsChangePb `json:"changes,omitempty"`
	// Whether any changes have been made to the cluster settings for the
	// cluster to become compliant with its policy.
	HasChanges bool `json:"has_changes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforceClusterComplianceResponseFromPb(pb *enforceClusterComplianceResponsePb) (*EnforceClusterComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforceClusterComplianceResponse{}

	var changesField []ClusterSettingsChange
	for _, itemPb := range pb.Changes {
		item, err := clusterSettingsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			changesField = append(changesField, *item)
		}
	}
	st.Changes = changesField
	hasChangesField, err := identity(&pb.HasChanges)
	if err != nil {
		return nil, err
	}
	if hasChangesField != nil {
		st.HasChanges = *hasChangesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforceClusterComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforceClusterComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The environment entity used to preserve serverless environment side panel,
// jobs' environment for non-notebook task, and DLT's environment for classic
// and serverless pipelines. (Note: DLT uses a copied version of the Environment
// proto below, at
// //spark/pipelines/api/protos/copied/libraries-environments-copy.proto) In
// this minimal environment spec, only pip dependencies are supported.
type Environment struct {
	// Client version used by the environment The client is the user-facing
	// environment of the runtime. Each client comes with a specific set of
	// pre-installed libraries. The version is a string, consisting of the major
	// client version.
	Client string
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url> E.g.
	// dependencies: ["foo==0.0.1", "-r /Workspace/test/requirements.txt"]
	Dependencies []string
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	JarDependencies []string
}

func EnvironmentToPb(st *Environment) (*EnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &EnvironmentPb{}
	clientPb, err := identity(&st.Client)
	if err != nil {
		return nil, err
	}
	if clientPb != nil {
		pb.Client = *clientPb
	}

	var dependenciesPb []string
	for _, item := range st.Dependencies {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependenciesPb = append(dependenciesPb, *itemPb)
		}
	}
	pb.Dependencies = dependenciesPb

	var jarDependenciesPb []string
	for _, item := range st.JarDependencies {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jarDependenciesPb = append(jarDependenciesPb, *itemPb)
		}
	}
	pb.JarDependencies = jarDependenciesPb

	return pb, nil
}

func (st *Environment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &EnvironmentPb{}
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

func (st Environment) MarshalJSON() ([]byte, error) {
	pb, err := EnvironmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnvironmentPb struct {
	// Client version used by the environment The client is the user-facing
	// environment of the runtime. Each client comes with a specific set of
	// pre-installed libraries. The version is a string, consisting of the major
	// client version.
	Client string `json:"client"`
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url> E.g.
	// dependencies: ["foo==0.0.1", "-r /Workspace/test/requirements.txt"]
	Dependencies []string `json:"dependencies,omitempty"`
	// List of jar dependencies, should be string representing volume paths. For
	// example: `/Volumes/path/to/test.jar`.
	JarDependencies []string `json:"jar_dependencies,omitempty"`
}

func EnvironmentFromPb(pb *EnvironmentPb) (*Environment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Environment{}
	clientField, err := identity(&pb.Client)
	if err != nil {
		return nil, err
	}
	if clientField != nil {
		st.Client = *clientField
	}

	var dependenciesField []string
	for _, itemPb := range pb.Dependencies {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependenciesField = append(dependenciesField, *item)
		}
	}
	st.Dependencies = dependenciesField

	var jarDependenciesField []string
	for _, itemPb := range pb.JarDependencies {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jarDependenciesField = append(jarDependenciesField, *item)
		}
	}
	st.JarDependencies = jarDependenciesField

	return st, nil
}

type EventDetails struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	Attributes *ClusterAttributes
	// The cause of a change in target size.
	Cause EventDetailsCause
	// The actual cluster size that was set in the cluster creation or edit.
	ClusterSize *ClusterSize
	// The current number of vCPUs in the cluster.
	CurrentNumVcpus int
	// The current number of nodes in the cluster.
	CurrentNumWorkers int

	DidNotExpandReason string
	// Current disk size in bytes
	DiskSize int64
	// More details about the change in driver's state
	DriverStateMessage string
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	EnableTerminationForNodeBlocklisted bool

	FreeSpace int64
	// List of global and cluster init scripts associated with this cluster
	// event.
	InitScripts *InitScriptEventDetails
	// Instance Id where the event originated from
	InstanceId string
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	JobRunName string
	// The cluster attributes before a cluster was edited.
	PreviousAttributes *ClusterAttributes
	// The size of the cluster before an edit or resize.
	PreviousClusterSize *ClusterSize
	// Previous disk size in bytes
	PreviousDiskSize int64
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	Reason *TerminationReason
	// The targeted number of vCPUs in the cluster.
	TargetNumVcpus int
	// The targeted number of nodes in the cluster.
	TargetNumWorkers int
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	User string

	ForceSendFields []string
}

func eventDetailsToPb(st *EventDetails) (*eventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &eventDetailsPb{}
	attributesPb, err := clusterAttributesToPb(st.Attributes)
	if err != nil {
		return nil, err
	}
	if attributesPb != nil {
		pb.Attributes = attributesPb
	}

	causePb, err := identity(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}

	clusterSizePb, err := clusterSizeToPb(st.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizePb != nil {
		pb.ClusterSize = clusterSizePb
	}

	currentNumVcpusPb, err := identity(&st.CurrentNumVcpus)
	if err != nil {
		return nil, err
	}
	if currentNumVcpusPb != nil {
		pb.CurrentNumVcpus = *currentNumVcpusPb
	}

	currentNumWorkersPb, err := identity(&st.CurrentNumWorkers)
	if err != nil {
		return nil, err
	}
	if currentNumWorkersPb != nil {
		pb.CurrentNumWorkers = *currentNumWorkersPb
	}

	didNotExpandReasonPb, err := identity(&st.DidNotExpandReason)
	if err != nil {
		return nil, err
	}
	if didNotExpandReasonPb != nil {
		pb.DidNotExpandReason = *didNotExpandReasonPb
	}

	diskSizePb, err := identity(&st.DiskSize)
	if err != nil {
		return nil, err
	}
	if diskSizePb != nil {
		pb.DiskSize = *diskSizePb
	}

	driverStateMessagePb, err := identity(&st.DriverStateMessage)
	if err != nil {
		return nil, err
	}
	if driverStateMessagePb != nil {
		pb.DriverStateMessage = *driverStateMessagePb
	}

	enableTerminationForNodeBlocklistedPb, err := identity(&st.EnableTerminationForNodeBlocklisted)
	if err != nil {
		return nil, err
	}
	if enableTerminationForNodeBlocklistedPb != nil {
		pb.EnableTerminationForNodeBlocklisted = *enableTerminationForNodeBlocklistedPb
	}

	freeSpacePb, err := identity(&st.FreeSpace)
	if err != nil {
		return nil, err
	}
	if freeSpacePb != nil {
		pb.FreeSpace = *freeSpacePb
	}

	initScriptsPb, err := initScriptEventDetailsToPb(st.InitScripts)
	if err != nil {
		return nil, err
	}
	if initScriptsPb != nil {
		pb.InitScripts = initScriptsPb
	}

	instanceIdPb, err := identity(&st.InstanceId)
	if err != nil {
		return nil, err
	}
	if instanceIdPb != nil {
		pb.InstanceId = *instanceIdPb
	}

	jobRunNamePb, err := identity(&st.JobRunName)
	if err != nil {
		return nil, err
	}
	if jobRunNamePb != nil {
		pb.JobRunName = *jobRunNamePb
	}

	previousAttributesPb, err := clusterAttributesToPb(st.PreviousAttributes)
	if err != nil {
		return nil, err
	}
	if previousAttributesPb != nil {
		pb.PreviousAttributes = previousAttributesPb
	}

	previousClusterSizePb, err := clusterSizeToPb(st.PreviousClusterSize)
	if err != nil {
		return nil, err
	}
	if previousClusterSizePb != nil {
		pb.PreviousClusterSize = previousClusterSizePb
	}

	previousDiskSizePb, err := identity(&st.PreviousDiskSize)
	if err != nil {
		return nil, err
	}
	if previousDiskSizePb != nil {
		pb.PreviousDiskSize = *previousDiskSizePb
	}

	reasonPb, err := terminationReasonToPb(st.Reason)
	if err != nil {
		return nil, err
	}
	if reasonPb != nil {
		pb.Reason = reasonPb
	}

	targetNumVcpusPb, err := identity(&st.TargetNumVcpus)
	if err != nil {
		return nil, err
	}
	if targetNumVcpusPb != nil {
		pb.TargetNumVcpus = *targetNumVcpusPb
	}

	targetNumWorkersPb, err := identity(&st.TargetNumWorkers)
	if err != nil {
		return nil, err
	}
	if targetNumWorkersPb != nil {
		pb.TargetNumWorkers = *targetNumWorkersPb
	}

	userPb, err := identity(&st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = *userPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type eventDetailsPb struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	Attributes *clusterAttributesPb `json:"attributes,omitempty"`
	// The cause of a change in target size.
	Cause EventDetailsCause `json:"cause,omitempty"`
	// The actual cluster size that was set in the cluster creation or edit.
	ClusterSize *clusterSizePb `json:"cluster_size,omitempty"`
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
	InitScripts *initScriptEventDetailsPb `json:"init_scripts,omitempty"`
	// Instance Id where the event originated from
	InstanceId string `json:"instance_id,omitempty"`
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	JobRunName string `json:"job_run_name,omitempty"`
	// The cluster attributes before a cluster was edited.
	PreviousAttributes *clusterAttributesPb `json:"previous_attributes,omitempty"`
	// The size of the cluster before an edit or resize.
	PreviousClusterSize *clusterSizePb `json:"previous_cluster_size,omitempty"`
	// Previous disk size in bytes
	PreviousDiskSize int64 `json:"previous_disk_size,omitempty"`
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	Reason *terminationReasonPb `json:"reason,omitempty"`
	// The targeted number of vCPUs in the cluster.
	TargetNumVcpus int `json:"target_num_vcpus,omitempty"`
	// The targeted number of nodes in the cluster.
	TargetNumWorkers int `json:"target_num_workers,omitempty"`
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	User string `json:"user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func eventDetailsFromPb(pb *eventDetailsPb) (*EventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EventDetails{}
	attributesField, err := clusterAttributesFromPb(pb.Attributes)
	if err != nil {
		return nil, err
	}
	if attributesField != nil {
		st.Attributes = attributesField
	}
	causeField, err := identity(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	clusterSizeField, err := clusterSizeFromPb(pb.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizeField != nil {
		st.ClusterSize = clusterSizeField
	}
	currentNumVcpusField, err := identity(&pb.CurrentNumVcpus)
	if err != nil {
		return nil, err
	}
	if currentNumVcpusField != nil {
		st.CurrentNumVcpus = *currentNumVcpusField
	}
	currentNumWorkersField, err := identity(&pb.CurrentNumWorkers)
	if err != nil {
		return nil, err
	}
	if currentNumWorkersField != nil {
		st.CurrentNumWorkers = *currentNumWorkersField
	}
	didNotExpandReasonField, err := identity(&pb.DidNotExpandReason)
	if err != nil {
		return nil, err
	}
	if didNotExpandReasonField != nil {
		st.DidNotExpandReason = *didNotExpandReasonField
	}
	diskSizeField, err := identity(&pb.DiskSize)
	if err != nil {
		return nil, err
	}
	if diskSizeField != nil {
		st.DiskSize = *diskSizeField
	}
	driverStateMessageField, err := identity(&pb.DriverStateMessage)
	if err != nil {
		return nil, err
	}
	if driverStateMessageField != nil {
		st.DriverStateMessage = *driverStateMessageField
	}
	enableTerminationForNodeBlocklistedField, err := identity(&pb.EnableTerminationForNodeBlocklisted)
	if err != nil {
		return nil, err
	}
	if enableTerminationForNodeBlocklistedField != nil {
		st.EnableTerminationForNodeBlocklisted = *enableTerminationForNodeBlocklistedField
	}
	freeSpaceField, err := identity(&pb.FreeSpace)
	if err != nil {
		return nil, err
	}
	if freeSpaceField != nil {
		st.FreeSpace = *freeSpaceField
	}
	initScriptsField, err := initScriptEventDetailsFromPb(pb.InitScripts)
	if err != nil {
		return nil, err
	}
	if initScriptsField != nil {
		st.InitScripts = initScriptsField
	}
	instanceIdField, err := identity(&pb.InstanceId)
	if err != nil {
		return nil, err
	}
	if instanceIdField != nil {
		st.InstanceId = *instanceIdField
	}
	jobRunNameField, err := identity(&pb.JobRunName)
	if err != nil {
		return nil, err
	}
	if jobRunNameField != nil {
		st.JobRunName = *jobRunNameField
	}
	previousAttributesField, err := clusterAttributesFromPb(pb.PreviousAttributes)
	if err != nil {
		return nil, err
	}
	if previousAttributesField != nil {
		st.PreviousAttributes = previousAttributesField
	}
	previousClusterSizeField, err := clusterSizeFromPb(pb.PreviousClusterSize)
	if err != nil {
		return nil, err
	}
	if previousClusterSizeField != nil {
		st.PreviousClusterSize = previousClusterSizeField
	}
	previousDiskSizeField, err := identity(&pb.PreviousDiskSize)
	if err != nil {
		return nil, err
	}
	if previousDiskSizeField != nil {
		st.PreviousDiskSize = *previousDiskSizeField
	}
	reasonField, err := terminationReasonFromPb(pb.Reason)
	if err != nil {
		return nil, err
	}
	if reasonField != nil {
		st.Reason = reasonField
	}
	targetNumVcpusField, err := identity(&pb.TargetNumVcpus)
	if err != nil {
		return nil, err
	}
	if targetNumVcpusField != nil {
		st.TargetNumVcpus = *targetNumVcpusField
	}
	targetNumWorkersField, err := identity(&pb.TargetNumWorkers)
	if err != nil {
		return nil, err
	}
	if targetNumWorkersField != nil {
		st.TargetNumWorkers = *targetNumWorkersField
	}
	userField, err := identity(&pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = *userField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *eventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st eventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The cause of a change in target size.
type EventDetailsCause string
type eventDetailsCausePb string

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

// Type always returns EventDetailsCause to satisfy [pflag.Value] interface
func (f *EventDetailsCause) Type() string {
	return "EventDetailsCause"
}

func eventDetailsCauseToPb(st *EventDetailsCause) (*eventDetailsCausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := eventDetailsCausePb(*st)
	return &pb, nil
}

func eventDetailsCauseFromPb(pb *eventDetailsCausePb) (*EventDetailsCause, error) {
	if pb == nil {
		return nil, nil
	}
	st := EventDetailsCause(*pb)
	return &st, nil
}

type EventType string
type eventTypePb string

const EventTypeAddNodesFailed EventType = `ADD_NODES_FAILED`

const EventTypeAutomaticClusterUpdate EventType = `AUTOMATIC_CLUSTER_UPDATE`

const EventTypeAutoscalingBackoff EventType = `AUTOSCALING_BACKOFF`

const EventTypeAutoscalingFailed EventType = `AUTOSCALING_FAILED`

const EventTypeAutoscalingStatsReport EventType = `AUTOSCALING_STATS_REPORT`

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
	case `ADD_NODES_FAILED`, `AUTOMATIC_CLUSTER_UPDATE`, `AUTOSCALING_BACKOFF`, `AUTOSCALING_FAILED`, `AUTOSCALING_STATS_REPORT`, `CREATING`, `DBFS_DOWN`, `DID_NOT_EXPAND_DISK`, `DRIVER_HEALTHY`, `DRIVER_NOT_RESPONDING`, `DRIVER_UNAVAILABLE`, `EDITED`, `EXPANDED_DISK`, `FAILED_TO_EXPAND_DISK`, `INIT_SCRIPTS_FINISHED`, `INIT_SCRIPTS_STARTED`, `METASTORE_DOWN`, `NODES_LOST`, `NODE_BLACKLISTED`, `NODE_EXCLUDED_DECOMMISSIONED`, `PINNED`, `RESIZING`, `RESTARTING`, `RUNNING`, `SPARK_EXCEPTION`, `STARTING`, `TERMINATING`, `UNPINNED`, `UPSIZE_COMPLETED`:
		*f = EventType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ADD_NODES_FAILED", "AUTOMATIC_CLUSTER_UPDATE", "AUTOSCALING_BACKOFF", "AUTOSCALING_FAILED", "AUTOSCALING_STATS_REPORT", "CREATING", "DBFS_DOWN", "DID_NOT_EXPAND_DISK", "DRIVER_HEALTHY", "DRIVER_NOT_RESPONDING", "DRIVER_UNAVAILABLE", "EDITED", "EXPANDED_DISK", "FAILED_TO_EXPAND_DISK", "INIT_SCRIPTS_FINISHED", "INIT_SCRIPTS_STARTED", "METASTORE_DOWN", "NODES_LOST", "NODE_BLACKLISTED", "NODE_EXCLUDED_DECOMMISSIONED", "PINNED", "RESIZING", "RESTARTING", "RUNNING", "SPARK_EXCEPTION", "STARTING", "TERMINATING", "UNPINNED", "UPSIZE_COMPLETED"`, v)
	}
}

// Type always returns EventType to satisfy [pflag.Value] interface
func (f *EventType) Type() string {
	return "EventType"
}

func eventTypeToPb(st *EventType) (*eventTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := eventTypePb(*st)
	return &pb, nil
}

func eventTypeFromPb(pb *eventTypePb) (*EventType, error) {
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
	Availability GcpAvailability
	// Boot disk size in GB
	BootDiskSize int
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	GoogleServiceAccount string
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount int
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// 'availability' field instead.
	UsePreemptibleExecutors bool
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default]. -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	ZoneId string

	ForceSendFields []string
}

func GcpAttributesToPb(st *GcpAttributes) (*GcpAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &GcpAttributesPb{}
	availabilityPb, err := identity(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}

	bootDiskSizePb, err := identity(&st.BootDiskSize)
	if err != nil {
		return nil, err
	}
	if bootDiskSizePb != nil {
		pb.BootDiskSize = *bootDiskSizePb
	}

	googleServiceAccountPb, err := identity(&st.GoogleServiceAccount)
	if err != nil {
		return nil, err
	}
	if googleServiceAccountPb != nil {
		pb.GoogleServiceAccount = *googleServiceAccountPb
	}

	localSsdCountPb, err := identity(&st.LocalSsdCount)
	if err != nil {
		return nil, err
	}
	if localSsdCountPb != nil {
		pb.LocalSsdCount = *localSsdCountPb
	}

	usePreemptibleExecutorsPb, err := identity(&st.UsePreemptibleExecutors)
	if err != nil {
		return nil, err
	}
	if usePreemptibleExecutorsPb != nil {
		pb.UsePreemptibleExecutors = *usePreemptibleExecutorsPb
	}

	zoneIdPb, err := identity(&st.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdPb != nil {
		pb.ZoneId = *zoneIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GcpAttributes) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &GcpAttributesPb{}
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

func (st GcpAttributes) MarshalJSON() ([]byte, error) {
	pb, err := GcpAttributesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GcpAttributesPb struct {
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

func GcpAttributesFromPb(pb *GcpAttributesPb) (*GcpAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcpAttributes{}
	availabilityField, err := identity(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	bootDiskSizeField, err := identity(&pb.BootDiskSize)
	if err != nil {
		return nil, err
	}
	if bootDiskSizeField != nil {
		st.BootDiskSize = *bootDiskSizeField
	}
	googleServiceAccountField, err := identity(&pb.GoogleServiceAccount)
	if err != nil {
		return nil, err
	}
	if googleServiceAccountField != nil {
		st.GoogleServiceAccount = *googleServiceAccountField
	}
	localSsdCountField, err := identity(&pb.LocalSsdCount)
	if err != nil {
		return nil, err
	}
	if localSsdCountField != nil {
		st.LocalSsdCount = *localSsdCountField
	}
	usePreemptibleExecutorsField, err := identity(&pb.UsePreemptibleExecutors)
	if err != nil {
		return nil, err
	}
	if usePreemptibleExecutorsField != nil {
		st.UsePreemptibleExecutors = *usePreemptibleExecutorsField
	}
	zoneIdField, err := identity(&pb.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdField != nil {
		st.ZoneId = *zoneIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *GcpAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GcpAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// This field determines whether the instance pool will contain preemptible VMs,
// on-demand VMs, or preemptible VMs with a fallback to on-demand VMs if the
// former is unavailable.
type GcpAvailability string
type gcpAvailabilityPb string

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

// Type always returns GcpAvailability to satisfy [pflag.Value] interface
func (f *GcpAvailability) Type() string {
	return "GcpAvailability"
}

func gcpAvailabilityToPb(st *GcpAvailability) (*gcpAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := gcpAvailabilityPb(*st)
	return &pb, nil
}

func gcpAvailabilityFromPb(pb *gcpAvailabilityPb) (*GcpAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := GcpAvailability(*pb)
	return &st, nil
}

// A storage location in Google Cloud Platform's GCS
type GcsStorageInfo struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination string
}

func gcsStorageInfoToPb(st *GcsStorageInfo) (*gcsStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gcsStorageInfoPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	return pb, nil
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

type gcsStorageInfoPb struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination string `json:"destination"`
}

func gcsStorageInfoFromPb(pb *gcsStorageInfoPb) (*GcsStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GcsStorageInfo{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}

	return st, nil
}

// Get cluster policy compliance
type GetClusterComplianceRequest struct {
	// The ID of the cluster to get the compliance status
	ClusterId string
}

func getClusterComplianceRequestToPb(st *GetClusterComplianceRequest) (*getClusterComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterComplianceRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type getClusterComplianceRequestPb struct {
	// The ID of the cluster to get the compliance status
	ClusterId string `json:"-" url:"cluster_id"`
}

func getClusterComplianceRequestFromPb(pb *getClusterComplianceRequestPb) (*GetClusterComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterComplianceRequest{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type GetClusterComplianceResponse struct {
	// Whether the cluster is compliant with its policy or not. Clusters could
	// be out of compliance if the policy was updated after the cluster was last
	// edited.
	IsCompliant bool
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. The values indicate an error message describing the
	// policy validation error.
	Violations map[string]string

	ForceSendFields []string
}

func getClusterComplianceResponseToPb(st *GetClusterComplianceResponse) (*getClusterComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterComplianceResponsePb{}
	isCompliantPb, err := identity(&st.IsCompliant)
	if err != nil {
		return nil, err
	}
	if isCompliantPb != nil {
		pb.IsCompliant = *isCompliantPb
	}

	violationsPb := map[string]string{}
	for k, v := range st.Violations {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			violationsPb[k] = *itemPb
		}
	}
	pb.Violations = violationsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getClusterComplianceResponsePb struct {
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

func getClusterComplianceResponseFromPb(pb *getClusterComplianceResponsePb) (*GetClusterComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterComplianceResponse{}
	isCompliantField, err := identity(&pb.IsCompliant)
	if err != nil {
		return nil, err
	}
	if isCompliantField != nil {
		st.IsCompliant = *isCompliantField
	}

	violationsField := map[string]string{}
	for k, v := range pb.Violations {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			violationsField[k] = *item
		}
	}
	st.Violations = violationsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getClusterComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getClusterComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get cluster permission levels
type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId string
}

func getClusterPermissionLevelsRequestToPb(st *GetClusterPermissionLevelsRequest) (*getClusterPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPermissionLevelsRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type getClusterPermissionLevelsRequestPb struct {
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" url:"-"`
}

func getClusterPermissionLevelsRequestFromPb(pb *getClusterPermissionLevelsRequestPb) (*GetClusterPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionLevelsRequest{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type GetClusterPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ClusterPermissionsDescription
}

func getClusterPermissionLevelsResponseToPb(st *GetClusterPermissionLevelsResponse) (*getClusterPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPermissionLevelsResponsePb{}

	var permissionLevelsPb []clusterPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := clusterPermissionsDescriptionToPb(&item)
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

type getClusterPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []clusterPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getClusterPermissionLevelsResponseFromPb(pb *getClusterPermissionLevelsResponsePb) (*GetClusterPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionLevelsResponse{}

	var permissionLevelsField []ClusterPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := clusterPermissionsDescriptionFromPb(&itemPb)
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

// Get cluster permissions
type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId string
}

func getClusterPermissionsRequestToPb(st *GetClusterPermissionsRequest) (*getClusterPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPermissionsRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type getClusterPermissionsRequestPb struct {
	// The cluster for which to get or manage permissions.
	ClusterId string `json:"-" url:"-"`
}

func getClusterPermissionsRequestFromPb(pb *getClusterPermissionsRequestPb) (*GetClusterPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPermissionsRequest{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

// Get cluster policy permission levels
type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string
}

func getClusterPolicyPermissionLevelsRequestToPb(st *GetClusterPolicyPermissionLevelsRequest) (*getClusterPolicyPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyPermissionLevelsRequestPb{}
	clusterPolicyIdPb, err := identity(&st.ClusterPolicyId)
	if err != nil {
		return nil, err
	}
	if clusterPolicyIdPb != nil {
		pb.ClusterPolicyId = *clusterPolicyIdPb
	}

	return pb, nil
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

type getClusterPolicyPermissionLevelsRequestPb struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" url:"-"`
}

func getClusterPolicyPermissionLevelsRequestFromPb(pb *getClusterPolicyPermissionLevelsRequestPb) (*GetClusterPolicyPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionLevelsRequest{}
	clusterPolicyIdField, err := identity(&pb.ClusterPolicyId)
	if err != nil {
		return nil, err
	}
	if clusterPolicyIdField != nil {
		st.ClusterPolicyId = *clusterPolicyIdField
	}

	return st, nil
}

type GetClusterPolicyPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ClusterPolicyPermissionsDescription
}

func getClusterPolicyPermissionLevelsResponseToPb(st *GetClusterPolicyPermissionLevelsResponse) (*getClusterPolicyPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyPermissionLevelsResponsePb{}

	var permissionLevelsPb []clusterPolicyPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := clusterPolicyPermissionsDescriptionToPb(&item)
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

type getClusterPolicyPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []clusterPolicyPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getClusterPolicyPermissionLevelsResponseFromPb(pb *getClusterPolicyPermissionLevelsResponsePb) (*GetClusterPolicyPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionLevelsResponse{}

	var permissionLevelsField []ClusterPolicyPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := clusterPolicyPermissionsDescriptionFromPb(&itemPb)
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

// Get cluster policy permissions
type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string
}

func getClusterPolicyPermissionsRequestToPb(st *GetClusterPolicyPermissionsRequest) (*getClusterPolicyPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyPermissionsRequestPb{}
	clusterPolicyIdPb, err := identity(&st.ClusterPolicyId)
	if err != nil {
		return nil, err
	}
	if clusterPolicyIdPb != nil {
		pb.ClusterPolicyId = *clusterPolicyIdPb
	}

	return pb, nil
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

type getClusterPolicyPermissionsRequestPb struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId string `json:"-" url:"-"`
}

func getClusterPolicyPermissionsRequestFromPb(pb *getClusterPolicyPermissionsRequestPb) (*GetClusterPolicyPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyPermissionsRequest{}
	clusterPolicyIdField, err := identity(&pb.ClusterPolicyId)
	if err != nil {
		return nil, err
	}
	if clusterPolicyIdField != nil {
		st.ClusterPolicyId = *clusterPolicyIdField
	}

	return st, nil
}

// Get a cluster policy
type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string
}

func getClusterPolicyRequestToPb(st *GetClusterPolicyRequest) (*getClusterPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterPolicyRequestPb{}
	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	return pb, nil
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

type getClusterPolicyRequestPb struct {
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string `json:"-" url:"policy_id"`
}

func getClusterPolicyRequestFromPb(pb *getClusterPolicyRequestPb) (*GetClusterPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterPolicyRequest{}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	return st, nil
}

// Get cluster info
type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	ClusterId string
}

func getClusterRequestToPb(st *GetClusterRequest) (*getClusterRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getClusterRequestPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type getClusterRequestPb struct {
	// The cluster about which to retrieve information.
	ClusterId string `json:"-" url:"cluster_id"`
}

func getClusterRequestFromPb(pb *getClusterRequestPb) (*GetClusterRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetClusterRequest{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type GetEvents struct {
	// The ID of the cluster to retrieve events about.
	ClusterId string
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	EndTime int64
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	EventTypes []EventType
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	Limit int64
	// Deprecated: use page_token in combination with page_size instead.
	//
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	Offset int64
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	Order GetEventsOrder
	// The maximum number of events to include in a page of events. The server
	// may further constrain the maximum number of results returned in a single
	// page. If the page_size is empty or 0, the server will decide the number
	// of results to be returned. The field has to be in the range [0,500]. If
	// the value is outside the range, the server enforces 0 or 500.
	PageSize int
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of events respectively. If page_token
	// is empty, the first page is returned.
	PageToken string
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	StartTime int64

	ForceSendFields []string
}

func getEventsToPb(st *GetEvents) (*getEventsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEventsPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	endTimePb, err := identity(&st.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimePb != nil {
		pb.EndTime = *endTimePb
	}

	var eventTypesPb []EventType
	for _, item := range st.EventTypes {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventTypesPb = append(eventTypesPb, *itemPb)
		}
	}
	pb.EventTypes = eventTypesPb

	limitPb, err := identity(&st.Limit)
	if err != nil {
		return nil, err
	}
	if limitPb != nil {
		pb.Limit = *limitPb
	}

	offsetPb, err := identity(&st.Offset)
	if err != nil {
		return nil, err
	}
	if offsetPb != nil {
		pb.Offset = *offsetPb
	}

	orderPb, err := identity(&st.Order)
	if err != nil {
		return nil, err
	}
	if orderPb != nil {
		pb.Order = *orderPb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getEventsPb struct {
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

func getEventsFromPb(pb *getEventsPb) (*GetEvents, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEvents{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	endTimeField, err := identity(&pb.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimeField != nil {
		st.EndTime = *endTimeField
	}

	var eventTypesField []EventType
	for _, itemPb := range pb.EventTypes {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventTypesField = append(eventTypesField, *item)
		}
	}
	st.EventTypes = eventTypesField
	limitField, err := identity(&pb.Limit)
	if err != nil {
		return nil, err
	}
	if limitField != nil {
		st.Limit = *limitField
	}
	offsetField, err := identity(&pb.Offset)
	if err != nil {
		return nil, err
	}
	if offsetField != nil {
		st.Offset = *offsetField
	}
	orderField, err := identity(&pb.Order)
	if err != nil {
		return nil, err
	}
	if orderField != nil {
		st.Order = *orderField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEventsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEventsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetEventsOrder string
type getEventsOrderPb string

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

// Type always returns GetEventsOrder to satisfy [pflag.Value] interface
func (f *GetEventsOrder) Type() string {
	return "GetEventsOrder"
}

func getEventsOrderToPb(st *GetEventsOrder) (*getEventsOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getEventsOrderPb(*st)
	return &pb, nil
}

func getEventsOrderFromPb(pb *getEventsOrderPb) (*GetEventsOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetEventsOrder(*pb)
	return &st, nil
}

type GetEventsResponse struct {
	Events []ClusterEvent
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage *GetEvents
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string
	// Deprecated: Returns 0 when request uses page_token. Will start returning
	// zero when request uses offset/limit soon.
	//
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	TotalCount int64

	ForceSendFields []string
}

func getEventsResponseToPb(st *GetEventsResponse) (*getEventsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEventsResponsePb{}

	var eventsPb []clusterEventPb
	for _, item := range st.Events {
		itemPb, err := clusterEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb

	nextPagePb, err := getEventsToPb(st.NextPage)
	if err != nil {
		return nil, err
	}
	if nextPagePb != nil {
		pb.NextPage = nextPagePb
	}

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	prevPageTokenPb, err := identity(&st.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenPb != nil {
		pb.PrevPageToken = *prevPageTokenPb
	}

	totalCountPb, err := identity(&st.TotalCount)
	if err != nil {
		return nil, err
	}
	if totalCountPb != nil {
		pb.TotalCount = *totalCountPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getEventsResponsePb struct {
	Events []clusterEventPb `json:"events,omitempty"`
	// Deprecated: use next_page_token or prev_page_token instead.
	//
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage *getEventsPb `json:"next_page,omitempty"`
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

func getEventsResponseFromPb(pb *getEventsResponsePb) (*GetEventsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEventsResponse{}

	var eventsField []ClusterEvent
	for _, itemPb := range pb.Events {
		item, err := clusterEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	nextPageField, err := getEventsFromPb(pb.NextPage)
	if err != nil {
		return nil, err
	}
	if nextPageField != nil {
		st.NextPage = nextPageField
	}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	prevPageTokenField, err := identity(&pb.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenField != nil {
		st.PrevPageToken = *prevPageTokenField
	}
	totalCountField, err := identity(&pb.TotalCount)
	if err != nil {
		return nil, err
	}
	if totalCountField != nil {
		st.TotalCount = *totalCountField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEventsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEventsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get an init script
type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId string
}

func getGlobalInitScriptRequestToPb(st *GetGlobalInitScriptRequest) (*getGlobalInitScriptRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getGlobalInitScriptRequestPb{}
	scriptIdPb, err := identity(&st.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdPb != nil {
		pb.ScriptId = *scriptIdPb
	}

	return pb, nil
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

type getGlobalInitScriptRequestPb struct {
	// The ID of the global init script.
	ScriptId string `json:"-" url:"-"`
}

func getGlobalInitScriptRequestFromPb(pb *getGlobalInitScriptRequestPb) (*GetGlobalInitScriptRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetGlobalInitScriptRequest{}
	scriptIdField, err := identity(&pb.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdField != nil {
		st.ScriptId = *scriptIdField
	}

	return st, nil
}

type GetInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string
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
	DefaultTags map[string]string
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int
	// Canonical unique identifier for the pool.
	InstancePoolId string
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int
	// For Fleet-pool V2, this object contains the information about the
	// alternate node type ids to use when attempting to launch a cluster if the
	// node type id is not available.
	NodeTypeFlexibility *NodeTypeFlexibility
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string
	// Current state of the instance pool.
	State InstancePoolState
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus

	ForceSendFields []string
}

func getInstancePoolToPb(st *GetInstancePool) (*getInstancePoolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPb{}
	awsAttributesPb, err := instancePoolAwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}

	azureAttributesPb, err := instancePoolAzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	defaultTagsPb := map[string]string{}
	for k, v := range st.DefaultTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			defaultTagsPb[k] = *itemPb
		}
	}
	pb.DefaultTags = defaultTagsPb

	diskSpecPb, err := diskSpecToPb(st.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecPb != nil {
		pb.DiskSpec = diskSpecPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	gcpAttributesPb, err := instancePoolGcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	idleInstanceAutoterminationMinutesPb, err := identity(&st.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesPb != nil {
		pb.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesPb
	}

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	instancePoolNamePb, err := identity(&st.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNamePb != nil {
		pb.InstancePoolName = *instancePoolNamePb
	}

	maxCapacityPb, err := identity(&st.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityPb != nil {
		pb.MaxCapacity = *maxCapacityPb
	}

	minIdleInstancesPb, err := identity(&st.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesPb != nil {
		pb.MinIdleInstances = *minIdleInstancesPb
	}

	nodeTypeFlexibilityPb, err := nodeTypeFlexibilityToPb(st.NodeTypeFlexibility)
	if err != nil {
		return nil, err
	}
	if nodeTypeFlexibilityPb != nil {
		pb.NodeTypeFlexibility = nodeTypeFlexibilityPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	var preloadedDockerImagesPb []dockerImagePb
	for _, item := range st.PreloadedDockerImages {
		itemPb, err := dockerImageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedDockerImagesPb = append(preloadedDockerImagesPb, *itemPb)
		}
	}
	pb.PreloadedDockerImages = preloadedDockerImagesPb

	var preloadedSparkVersionsPb []string
	for _, item := range st.PreloadedSparkVersions {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedSparkVersionsPb = append(preloadedSparkVersionsPb, *itemPb)
		}
	}
	pb.PreloadedSparkVersions = preloadedSparkVersionsPb

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	statsPb, err := instancePoolStatsToPb(st.Stats)
	if err != nil {
		return nil, err
	}
	if statsPb != nil {
		pb.Stats = statsPb
	}

	statusPb, err := instancePoolStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getInstancePoolPb struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *instancePoolAwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *instancePoolAzureAttributesPb `json:"azure_attributes,omitempty"`
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
	DiskSpec *diskSpecPb `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *instancePoolGcpAttributesPb `json:"gcp_attributes,omitempty"`
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
	// For Fleet-pool V2, this object contains the information about the
	// alternate node type ids to use when attempting to launch a cluster if the
	// node type id is not available.
	NodeTypeFlexibility *nodeTypeFlexibilityPb `json:"node_type_flexibility,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []dockerImagePb `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// Current state of the instance pool.
	State InstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	Stats *instancePoolStatsPb `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	Status *instancePoolStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getInstancePoolFromPb(pb *getInstancePoolPb) (*GetInstancePool, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePool{}
	awsAttributesField, err := instancePoolAwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := instancePoolAzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField

	defaultTagsField := map[string]string{}
	for k, v := range pb.DefaultTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			defaultTagsField[k] = *item
		}
	}
	st.DefaultTags = defaultTagsField
	diskSpecField, err := diskSpecFromPb(pb.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecField != nil {
		st.DiskSpec = diskSpecField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	gcpAttributesField, err := instancePoolGcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}
	idleInstanceAutoterminationMinutesField, err := identity(&pb.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesField != nil {
		st.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesField
	}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	instancePoolNameField, err := identity(&pb.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNameField != nil {
		st.InstancePoolName = *instancePoolNameField
	}
	maxCapacityField, err := identity(&pb.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityField != nil {
		st.MaxCapacity = *maxCapacityField
	}
	minIdleInstancesField, err := identity(&pb.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesField != nil {
		st.MinIdleInstances = *minIdleInstancesField
	}
	nodeTypeFlexibilityField, err := nodeTypeFlexibilityFromPb(pb.NodeTypeFlexibility)
	if err != nil {
		return nil, err
	}
	if nodeTypeFlexibilityField != nil {
		st.NodeTypeFlexibility = nodeTypeFlexibilityField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}

	var preloadedDockerImagesField []DockerImage
	for _, itemPb := range pb.PreloadedDockerImages {
		item, err := dockerImageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedDockerImagesField = append(preloadedDockerImagesField, *item)
		}
	}
	st.PreloadedDockerImages = preloadedDockerImagesField

	var preloadedSparkVersionsField []string
	for _, itemPb := range pb.PreloadedSparkVersions {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedSparkVersionsField = append(preloadedSparkVersionsField, *item)
		}
	}
	st.PreloadedSparkVersions = preloadedSparkVersionsField
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	statsField, err := instancePoolStatsFromPb(pb.Stats)
	if err != nil {
		return nil, err
	}
	if statsField != nil {
		st.Stats = statsField
	}
	statusField, err := instancePoolStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getInstancePoolPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getInstancePoolPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get instance pool permission levels
type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string
}

func getInstancePoolPermissionLevelsRequestToPb(st *GetInstancePoolPermissionLevelsRequest) (*getInstancePoolPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPermissionLevelsRequestPb{}
	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	return pb, nil
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

type getInstancePoolPermissionLevelsRequestPb struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" url:"-"`
}

func getInstancePoolPermissionLevelsRequestFromPb(pb *getInstancePoolPermissionLevelsRequestPb) (*GetInstancePoolPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionLevelsRequest{}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}

	return st, nil
}

type GetInstancePoolPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []InstancePoolPermissionsDescription
}

func getInstancePoolPermissionLevelsResponseToPb(st *GetInstancePoolPermissionLevelsResponse) (*getInstancePoolPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPermissionLevelsResponsePb{}

	var permissionLevelsPb []instancePoolPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := instancePoolPermissionsDescriptionToPb(&item)
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

type getInstancePoolPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []instancePoolPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getInstancePoolPermissionLevelsResponseFromPb(pb *getInstancePoolPermissionLevelsResponsePb) (*GetInstancePoolPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionLevelsResponse{}

	var permissionLevelsField []InstancePoolPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := instancePoolPermissionsDescriptionFromPb(&itemPb)
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

// Get instance pool permissions
type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string
}

func getInstancePoolPermissionsRequestToPb(st *GetInstancePoolPermissionsRequest) (*getInstancePoolPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolPermissionsRequestPb{}
	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	return pb, nil
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

type getInstancePoolPermissionsRequestPb struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" url:"-"`
}

func getInstancePoolPermissionsRequestFromPb(pb *getInstancePoolPermissionsRequestPb) (*GetInstancePoolPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolPermissionsRequest{}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}

	return st, nil
}

// Get instance pool information
type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId string
}

func getInstancePoolRequestToPb(st *GetInstancePoolRequest) (*getInstancePoolRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getInstancePoolRequestPb{}
	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	return pb, nil
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

type getInstancePoolRequestPb struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId string `json:"-" url:"instance_pool_id"`
}

func getInstancePoolRequestFromPb(pb *getInstancePoolRequestPb) (*GetInstancePoolRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetInstancePoolRequest{}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}

	return st, nil
}

// Get policy family information
type GetPolicyFamilyRequest struct {
	// The family ID about which to retrieve information.
	PolicyFamilyId string
	// The version number for the family to fetch. Defaults to the latest
	// version.
	Version int64

	ForceSendFields []string
}

func getPolicyFamilyRequestToPb(st *GetPolicyFamilyRequest) (*getPolicyFamilyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPolicyFamilyRequestPb{}
	policyFamilyIdPb, err := identity(&st.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdPb != nil {
		pb.PolicyFamilyId = *policyFamilyIdPb
	}

	versionPb, err := identity(&st.Version)
	if err != nil {
		return nil, err
	}
	if versionPb != nil {
		pb.Version = *versionPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getPolicyFamilyRequestPb struct {
	// The family ID about which to retrieve information.
	PolicyFamilyId string `json:"-" url:"-"`
	// The version number for the family to fetch. Defaults to the latest
	// version.
	Version int64 `json:"-" url:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPolicyFamilyRequestFromPb(pb *getPolicyFamilyRequestPb) (*GetPolicyFamilyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyFamilyRequest{}
	policyFamilyIdField, err := identity(&pb.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdField != nil {
		st.PolicyFamilyId = *policyFamilyIdField
	}
	versionField, err := identity(&pb.Version)
	if err != nil {
		return nil, err
	}
	if versionField != nil {
		st.Version = *versionField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPolicyFamilyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPolicyFamilyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSparkVersionsResponse struct {
	// All the available Spark versions.
	Versions []SparkVersion
}

func getSparkVersionsResponseToPb(st *GetSparkVersionsResponse) (*getSparkVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSparkVersionsResponsePb{}

	var versionsPb []sparkVersionPb
	for _, item := range st.Versions {
		itemPb, err := sparkVersionToPb(&item)
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

type getSparkVersionsResponsePb struct {
	// All the available Spark versions.
	Versions []sparkVersionPb `json:"versions,omitempty"`
}

func getSparkVersionsResponseFromPb(pb *getSparkVersionsResponsePb) (*GetSparkVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSparkVersionsResponse{}

	var versionsField []SparkVersion
	for _, itemPb := range pb.Versions {
		item, err := sparkVersionFromPb(&itemPb)
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
	Enabled bool
	// The name of the script
	Name string
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
	Position int
	// The Base64-encoded content of the script.
	Script string

	ForceSendFields []string
}

func globalInitScriptCreateRequestToPb(st *GlobalInitScriptCreateRequest) (*globalInitScriptCreateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptCreateRequestPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	positionPb, err := identity(&st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = *positionPb
	}

	scriptPb, err := identity(&st.Script)
	if err != nil {
		return nil, err
	}
	if scriptPb != nil {
		pb.Script = *scriptPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type globalInitScriptCreateRequestPb struct {
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

func globalInitScriptCreateRequestFromPb(pb *globalInitScriptCreateRequestPb) (*GlobalInitScriptCreateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptCreateRequest{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	positionField, err := identity(&pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = *positionField
	}
	scriptField, err := identity(&pb.Script)
	if err != nil {
		return nil, err
	}
	if scriptField != nil {
		st.Script = *scriptField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptCreateRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptCreateRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GlobalInitScriptDetails struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt int
	// The username of the user who created the script.
	CreatedBy string
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool
	// The name of the script
	Name string
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position int
	// The global init script ID.
	ScriptId string
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt int
	// The username of the user who last updated the script
	UpdatedBy string

	ForceSendFields []string
}

func globalInitScriptDetailsToPb(st *GlobalInitScriptDetails) (*globalInitScriptDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptDetailsPb{}
	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	positionPb, err := identity(&st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = *positionPb
	}

	scriptIdPb, err := identity(&st.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdPb != nil {
		pb.ScriptId = *scriptIdPb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type globalInitScriptDetailsPb struct {
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

func globalInitScriptDetailsFromPb(pb *globalInitScriptDetailsPb) (*GlobalInitScriptDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptDetails{}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	positionField, err := identity(&pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = *positionField
	}
	scriptIdField, err := identity(&pb.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdField != nil {
		st.ScriptId = *scriptIdField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt int
	// The username of the user who created the script.
	CreatedBy string
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool
	// The name of the script
	Name string
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position int
	// The Base64-encoded content of the script.
	Script string
	// The global init script ID.
	ScriptId string
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt int
	// The username of the user who last updated the script
	UpdatedBy string

	ForceSendFields []string
}

func globalInitScriptDetailsWithContentToPb(st *GlobalInitScriptDetailsWithContent) (*globalInitScriptDetailsWithContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptDetailsWithContentPb{}
	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	positionPb, err := identity(&st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = *positionPb
	}

	scriptPb, err := identity(&st.Script)
	if err != nil {
		return nil, err
	}
	if scriptPb != nil {
		pb.Script = *scriptPb
	}

	scriptIdPb, err := identity(&st.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdPb != nil {
		pb.ScriptId = *scriptIdPb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type globalInitScriptDetailsWithContentPb struct {
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

func globalInitScriptDetailsWithContentFromPb(pb *globalInitScriptDetailsWithContentPb) (*GlobalInitScriptDetailsWithContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptDetailsWithContent{}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	positionField, err := identity(&pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = *positionField
	}
	scriptField, err := identity(&pb.Script)
	if err != nil {
		return nil, err
	}
	if scriptField != nil {
		st.Script = *scriptField
	}
	scriptIdField, err := identity(&pb.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdField != nil {
		st.ScriptId = *scriptIdField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptDetailsWithContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptDetailsWithContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GlobalInitScriptUpdateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled bool
	// The name of the script
	Name string
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
	Position int
	// The Base64-encoded content of the script.
	Script string
	// The ID of the global init script.
	ScriptId string

	ForceSendFields []string
}

func globalInitScriptUpdateRequestToPb(st *GlobalInitScriptUpdateRequest) (*globalInitScriptUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &globalInitScriptUpdateRequestPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	positionPb, err := identity(&st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = *positionPb
	}

	scriptPb, err := identity(&st.Script)
	if err != nil {
		return nil, err
	}
	if scriptPb != nil {
		pb.Script = *scriptPb
	}

	scriptIdPb, err := identity(&st.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdPb != nil {
		pb.ScriptId = *scriptIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type globalInitScriptUpdateRequestPb struct {
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

func globalInitScriptUpdateRequestFromPb(pb *globalInitScriptUpdateRequestPb) (*GlobalInitScriptUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GlobalInitScriptUpdateRequest{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	positionField, err := identity(&pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = *positionField
	}
	scriptField, err := identity(&pb.Script)
	if err != nil {
		return nil, err
	}
	if scriptField != nil {
		st.Script = *scriptField
	}
	scriptIdField, err := identity(&pb.ScriptId)
	if err != nil {
		return nil, err
	}
	if scriptIdField != nil {
		st.ScriptId = *scriptIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *globalInitScriptUpdateRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st globalInitScriptUpdateRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InitScriptEventDetails struct {
	// The cluster scoped init scripts associated with this cluster event.
	Cluster []InitScriptInfoAndExecutionDetails
	// The global init scripts associated with this cluster event.
	Global []InitScriptInfoAndExecutionDetails
	// The private ip of the node we are reporting init script execution details
	// for (we will select the execution details from only one node rather than
	// reporting the execution details from every node to keep these event
	// details small)
	//
	// This should only be defined for the INIT_SCRIPTS_FINISHED event
	ReportedForNode string

	ForceSendFields []string
}

func initScriptEventDetailsToPb(st *InitScriptEventDetails) (*initScriptEventDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &initScriptEventDetailsPb{}

	var clusterPb []initScriptInfoAndExecutionDetailsPb
	for _, item := range st.Cluster {
		itemPb, err := initScriptInfoAndExecutionDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clusterPb = append(clusterPb, *itemPb)
		}
	}
	pb.Cluster = clusterPb

	var globalPb []initScriptInfoAndExecutionDetailsPb
	for _, item := range st.Global {
		itemPb, err := initScriptInfoAndExecutionDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			globalPb = append(globalPb, *itemPb)
		}
	}
	pb.Global = globalPb

	reportedForNodePb, err := identity(&st.ReportedForNode)
	if err != nil {
		return nil, err
	}
	if reportedForNodePb != nil {
		pb.ReportedForNode = *reportedForNodePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type initScriptEventDetailsPb struct {
	// The cluster scoped init scripts associated with this cluster event.
	Cluster []initScriptInfoAndExecutionDetailsPb `json:"cluster,omitempty"`
	// The global init scripts associated with this cluster event.
	Global []initScriptInfoAndExecutionDetailsPb `json:"global,omitempty"`
	// The private ip of the node we are reporting init script execution details
	// for (we will select the execution details from only one node rather than
	// reporting the execution details from every node to keep these event
	// details small)
	//
	// This should only be defined for the INIT_SCRIPTS_FINISHED event
	ReportedForNode string `json:"reported_for_node,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func initScriptEventDetailsFromPb(pb *initScriptEventDetailsPb) (*InitScriptEventDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptEventDetails{}

	var clusterField []InitScriptInfoAndExecutionDetails
	for _, itemPb := range pb.Cluster {
		item, err := initScriptInfoAndExecutionDetailsFromPb(&itemPb)
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
		item, err := initScriptInfoAndExecutionDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			globalField = append(globalField, *item)
		}
	}
	st.Global = globalField
	reportedForNodeField, err := identity(&pb.ReportedForNode)
	if err != nil {
		return nil, err
	}
	if reportedForNodeField != nil {
		st.ReportedForNode = *reportedForNodeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *initScriptEventDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st initScriptEventDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Result of attempted script execution
type InitScriptExecutionDetailsInitScriptExecutionStatus string
type initScriptExecutionDetailsInitScriptExecutionStatusPb string

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

// Type always returns InitScriptExecutionDetailsInitScriptExecutionStatus to satisfy [pflag.Value] interface
func (f *InitScriptExecutionDetailsInitScriptExecutionStatus) Type() string {
	return "InitScriptExecutionDetailsInitScriptExecutionStatus"
}

func initScriptExecutionDetailsInitScriptExecutionStatusToPb(st *InitScriptExecutionDetailsInitScriptExecutionStatus) (*initScriptExecutionDetailsInitScriptExecutionStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := initScriptExecutionDetailsInitScriptExecutionStatusPb(*st)
	return &pb, nil
}

func initScriptExecutionDetailsInitScriptExecutionStatusFromPb(pb *initScriptExecutionDetailsInitScriptExecutionStatusPb) (*InitScriptExecutionDetailsInitScriptExecutionStatus, error) {
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
	Abfss *Adlsgen2Info
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File *LocalFileInfo
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *GcsStorageInfo
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 *S3StorageInfo
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes *VolumesStorageInfo
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace *WorkspaceStorageInfo
}

func InitScriptInfoToPb(st *InitScriptInfo) (*InitScriptInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &InitScriptInfoPb{}
	abfssPb, err := adlsgen2InfoToPb(st.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssPb != nil {
		pb.Abfss = abfssPb
	}

	dbfsPb, err := dbfsStorageInfoToPb(st.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsPb != nil {
		pb.Dbfs = dbfsPb
	}

	filePb, err := localFileInfoToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}

	gcsPb, err := gcsStorageInfoToPb(st.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsPb != nil {
		pb.Gcs = gcsPb
	}

	s3Pb, err := s3StorageInfoToPb(st.S3)
	if err != nil {
		return nil, err
	}
	if s3Pb != nil {
		pb.S3 = s3Pb
	}

	volumesPb, err := volumesStorageInfoToPb(st.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesPb != nil {
		pb.Volumes = volumesPb
	}

	workspacePb, err := workspaceStorageInfoToPb(st.Workspace)
	if err != nil {
		return nil, err
	}
	if workspacePb != nil {
		pb.Workspace = workspacePb
	}

	return pb, nil
}

func (st *InitScriptInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &InitScriptInfoPb{}
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

func (st InitScriptInfo) MarshalJSON() ([]byte, error) {
	pb, err := InitScriptInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type InitScriptInfoPb struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss *adlsgen2InfoPb `json:"abfss,omitempty"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *dbfsStorageInfoPb `json:"dbfs,omitempty"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File *localFileInfoPb `json:"file,omitempty"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *gcsStorageInfoPb `json:"gcs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 *s3StorageInfoPb `json:"s3,omitempty"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes *volumesStorageInfoPb `json:"volumes,omitempty"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace *workspaceStorageInfoPb `json:"workspace,omitempty"`
}

func InitScriptInfoFromPb(pb *InitScriptInfoPb) (*InitScriptInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptInfo{}
	abfssField, err := adlsgen2InfoFromPb(pb.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssField != nil {
		st.Abfss = abfssField
	}
	dbfsField, err := dbfsStorageInfoFromPb(pb.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsField != nil {
		st.Dbfs = dbfsField
	}
	fileField, err := localFileInfoFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	gcsField, err := gcsStorageInfoFromPb(pb.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsField != nil {
		st.Gcs = gcsField
	}
	s3Field, err := s3StorageInfoFromPb(pb.S3)
	if err != nil {
		return nil, err
	}
	if s3Field != nil {
		st.S3 = s3Field
	}
	volumesField, err := volumesStorageInfoFromPb(pb.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesField != nil {
		st.Volumes = volumesField
	}
	workspaceField, err := workspaceStorageInfoFromPb(pb.Workspace)
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
	Abfss *Adlsgen2Info
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	ErrorMessage string
	// The number duration of the script execution in seconds
	ExecutionDurationSeconds int
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File *LocalFileInfo
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *GcsStorageInfo
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 *S3StorageInfo
	// The current status of the script
	Status InitScriptExecutionDetailsInitScriptExecutionStatus
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes *VolumesStorageInfo
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace *WorkspaceStorageInfo

	ForceSendFields []string
}

func initScriptInfoAndExecutionDetailsToPb(st *InitScriptInfoAndExecutionDetails) (*initScriptInfoAndExecutionDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &initScriptInfoAndExecutionDetailsPb{}
	abfssPb, err := adlsgen2InfoToPb(st.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssPb != nil {
		pb.Abfss = abfssPb
	}

	dbfsPb, err := dbfsStorageInfoToPb(st.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsPb != nil {
		pb.Dbfs = dbfsPb
	}

	errorMessagePb, err := identity(&st.ErrorMessage)
	if err != nil {
		return nil, err
	}
	if errorMessagePb != nil {
		pb.ErrorMessage = *errorMessagePb
	}

	executionDurationSecondsPb, err := identity(&st.ExecutionDurationSeconds)
	if err != nil {
		return nil, err
	}
	if executionDurationSecondsPb != nil {
		pb.ExecutionDurationSeconds = *executionDurationSecondsPb
	}

	filePb, err := localFileInfoToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
	}

	gcsPb, err := gcsStorageInfoToPb(st.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsPb != nil {
		pb.Gcs = gcsPb
	}

	s3Pb, err := s3StorageInfoToPb(st.S3)
	if err != nil {
		return nil, err
	}
	if s3Pb != nil {
		pb.S3 = s3Pb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	volumesPb, err := volumesStorageInfoToPb(st.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesPb != nil {
		pb.Volumes = volumesPb
	}

	workspacePb, err := workspaceStorageInfoToPb(st.Workspace)
	if err != nil {
		return nil, err
	}
	if workspacePb != nil {
		pb.Workspace = workspacePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type initScriptInfoAndExecutionDetailsPb struct {
	// destination needs to be provided, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`
	Abfss *adlsgen2InfoPb `json:"abfss,omitempty"`
	// destination needs to be provided. e.g. `{ "dbfs": { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *dbfsStorageInfoPb `json:"dbfs,omitempty"`
	// Additional details regarding errors (such as a file not found message if
	// the status is FAILED_FETCH). This field should only be used to provide
	// *additional* information to the status field, not duplicate it.
	ErrorMessage string `json:"error_message,omitempty"`
	// The number duration of the script execution in seconds
	ExecutionDurationSeconds int `json:"execution_duration_seconds,omitempty"`
	// destination needs to be provided, e.g. `{ "file": { "destination":
	// "file:/my/local/file.sh" } }`
	File *localFileInfoPb `json:"file,omitempty"`
	// destination needs to be provided, e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *gcsStorageInfoPb `json:"gcs,omitempty"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ \"s3\": { \"destination\": \"s3://cluster_log_bucket/prefix\",
	// \"region\": \"us-west-2\" } }` Cluster iam role is used to access s3,
	// please make sure the cluster iam role in `instance_profile_arn` has
	// permission to write data to the s3 destination.
	S3 *s3StorageInfoPb `json:"s3,omitempty"`
	// The current status of the script
	Status InitScriptExecutionDetailsInitScriptExecutionStatus `json:"status,omitempty"`
	// destination needs to be provided. e.g. `{ \"volumes\" : { \"destination\"
	// : \"/Volumes/my-init.sh\" } }`
	Volumes *volumesStorageInfoPb `json:"volumes,omitempty"`
	// destination needs to be provided, e.g. `{ "workspace": { "destination":
	// "/cluster-init-scripts/setup-datadog.sh" } }`
	Workspace *workspaceStorageInfoPb `json:"workspace,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func initScriptInfoAndExecutionDetailsFromPb(pb *initScriptInfoAndExecutionDetailsPb) (*InitScriptInfoAndExecutionDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InitScriptInfoAndExecutionDetails{}
	abfssField, err := adlsgen2InfoFromPb(pb.Abfss)
	if err != nil {
		return nil, err
	}
	if abfssField != nil {
		st.Abfss = abfssField
	}
	dbfsField, err := dbfsStorageInfoFromPb(pb.Dbfs)
	if err != nil {
		return nil, err
	}
	if dbfsField != nil {
		st.Dbfs = dbfsField
	}
	errorMessageField, err := identity(&pb.ErrorMessage)
	if err != nil {
		return nil, err
	}
	if errorMessageField != nil {
		st.ErrorMessage = *errorMessageField
	}
	executionDurationSecondsField, err := identity(&pb.ExecutionDurationSeconds)
	if err != nil {
		return nil, err
	}
	if executionDurationSecondsField != nil {
		st.ExecutionDurationSeconds = *executionDurationSecondsField
	}
	fileField, err := localFileInfoFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
	}
	gcsField, err := gcsStorageInfoFromPb(pb.Gcs)
	if err != nil {
		return nil, err
	}
	if gcsField != nil {
		st.Gcs = gcsField
	}
	s3Field, err := s3StorageInfoFromPb(pb.S3)
	if err != nil {
		return nil, err
	}
	if s3Field != nil {
		st.S3 = s3Field
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	volumesField, err := volumesStorageInfoFromPb(pb.Volumes)
	if err != nil {
		return nil, err
	}
	if volumesField != nil {
		st.Volumes = volumesField
	}
	workspaceField, err := workspaceStorageInfoFromPb(pb.Workspace)
	if err != nil {
		return nil, err
	}
	if workspaceField != nil {
		st.Workspace = workspaceField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *initScriptInfoAndExecutionDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st initScriptInfoAndExecutionDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId string
	// The libraries to install.
	Libraries []Library
}

func installLibrariesToPb(st *InstallLibraries) (*installLibrariesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &installLibrariesPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	var librariesPb []LibraryPb
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

type installLibrariesPb struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to install.
	Libraries []LibraryPb `json:"libraries"`
}

func installLibrariesFromPb(pb *installLibrariesPb) (*InstallLibraries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstallLibraries{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

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

type InstallLibrariesResponse struct {
}

func installLibrariesResponseToPb(st *InstallLibrariesResponse) (*installLibrariesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &installLibrariesResponsePb{}

	return pb, nil
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

type installLibrariesResponsePb struct {
}

func installLibrariesResponseFromPb(pb *installLibrariesResponsePb) (*InstallLibrariesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstallLibrariesResponse{}

	return st, nil
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel InstancePoolPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func instancePoolAccessControlRequestToPb(st *InstancePoolAccessControlRequest) (*instancePoolAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAccessControlRequestFromPb(pb *instancePoolAccessControlRequestPb) (*InstancePoolAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAccessControlRequest{}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAccessControlResponse struct {
	// All permissions.
	AllPermissions []InstancePoolPermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func instancePoolAccessControlResponseToPb(st *InstancePoolAccessControlResponse) (*instancePoolAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAccessControlResponsePb{}

	var allPermissionsPb []instancePoolPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := instancePoolPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []instancePoolPermissionPb `json:"all_permissions,omitempty"`
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

func instancePoolAccessControlResponseFromPb(pb *instancePoolAccessControlResponsePb) (*InstancePoolAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAccessControlResponse{}

	var allPermissionsField []InstancePoolPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := instancePoolPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]string
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
	DefaultTags map[string]string
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes int
	// Canonical unique identifier for the pool.
	InstancePoolId string
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName string
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity int
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances int
	// For Fleet-pool V2, this object contains the information about the
	// alternate node type ids to use when attempting to launch a cluster if the
	// node type id is not available.
	NodeTypeFlexibility *NodeTypeFlexibility
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string
	// Current state of the instance pool.
	State InstancePoolState
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus

	ForceSendFields []string
}

func instancePoolAndStatsToPb(st *InstancePoolAndStats) (*instancePoolAndStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAndStatsPb{}
	awsAttributesPb, err := instancePoolAwsAttributesToPb(st.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesPb != nil {
		pb.AwsAttributes = awsAttributesPb
	}

	azureAttributesPb, err := instancePoolAzureAttributesToPb(st.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesPb != nil {
		pb.AzureAttributes = azureAttributesPb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	defaultTagsPb := map[string]string{}
	for k, v := range st.DefaultTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			defaultTagsPb[k] = *itemPb
		}
	}
	pb.DefaultTags = defaultTagsPb

	diskSpecPb, err := diskSpecToPb(st.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecPb != nil {
		pb.DiskSpec = diskSpecPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	gcpAttributesPb, err := instancePoolGcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	idleInstanceAutoterminationMinutesPb, err := identity(&st.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesPb != nil {
		pb.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesPb
	}

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	instancePoolNamePb, err := identity(&st.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNamePb != nil {
		pb.InstancePoolName = *instancePoolNamePb
	}

	maxCapacityPb, err := identity(&st.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityPb != nil {
		pb.MaxCapacity = *maxCapacityPb
	}

	minIdleInstancesPb, err := identity(&st.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesPb != nil {
		pb.MinIdleInstances = *minIdleInstancesPb
	}

	nodeTypeFlexibilityPb, err := nodeTypeFlexibilityToPb(st.NodeTypeFlexibility)
	if err != nil {
		return nil, err
	}
	if nodeTypeFlexibilityPb != nil {
		pb.NodeTypeFlexibility = nodeTypeFlexibilityPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	var preloadedDockerImagesPb []dockerImagePb
	for _, item := range st.PreloadedDockerImages {
		itemPb, err := dockerImageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedDockerImagesPb = append(preloadedDockerImagesPb, *itemPb)
		}
	}
	pb.PreloadedDockerImages = preloadedDockerImagesPb

	var preloadedSparkVersionsPb []string
	for _, item := range st.PreloadedSparkVersions {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			preloadedSparkVersionsPb = append(preloadedSparkVersionsPb, *itemPb)
		}
	}
	pb.PreloadedSparkVersions = preloadedSparkVersionsPb

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	statsPb, err := instancePoolStatsToPb(st.Stats)
	if err != nil {
		return nil, err
	}
	if statsPb != nil {
		pb.Stats = statsPb
	}

	statusPb, err := instancePoolStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolAndStatsPb struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *instancePoolAwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *instancePoolAzureAttributesPb `json:"azure_attributes,omitempty"`
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
	DiskSpec *diskSpecPb `json:"disk_spec,omitempty"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk bool `json:"enable_elastic_disk,omitempty"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *instancePoolGcpAttributesPb `json:"gcp_attributes,omitempty"`
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
	// For Fleet-pool V2, this object contains the information about the
	// alternate node type ids to use when attempting to launch a cluster if the
	// node type id is not available.
	NodeTypeFlexibility *nodeTypeFlexibilityPb `json:"node_type_flexibility,omitempty"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string `json:"node_type_id,omitempty"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []dockerImagePb `json:"preloaded_docker_images,omitempty"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []string `json:"preloaded_spark_versions,omitempty"`
	// Current state of the instance pool.
	State InstancePoolState `json:"state,omitempty"`
	// Usage statistics about the instance pool.
	Stats *instancePoolStatsPb `json:"stats,omitempty"`
	// Status of failed pending instances in the pool.
	Status *instancePoolStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolAndStatsFromPb(pb *instancePoolAndStatsPb) (*InstancePoolAndStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAndStats{}
	awsAttributesField, err := instancePoolAwsAttributesFromPb(pb.AwsAttributes)
	if err != nil {
		return nil, err
	}
	if awsAttributesField != nil {
		st.AwsAttributes = awsAttributesField
	}
	azureAttributesField, err := instancePoolAzureAttributesFromPb(pb.AzureAttributes)
	if err != nil {
		return nil, err
	}
	if azureAttributesField != nil {
		st.AzureAttributes = azureAttributesField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField

	defaultTagsField := map[string]string{}
	for k, v := range pb.DefaultTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			defaultTagsField[k] = *item
		}
	}
	st.DefaultTags = defaultTagsField
	diskSpecField, err := diskSpecFromPb(pb.DiskSpec)
	if err != nil {
		return nil, err
	}
	if diskSpecField != nil {
		st.DiskSpec = diskSpecField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	gcpAttributesField, err := instancePoolGcpAttributesFromPb(pb.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesField != nil {
		st.GcpAttributes = gcpAttributesField
	}
	idleInstanceAutoterminationMinutesField, err := identity(&pb.IdleInstanceAutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if idleInstanceAutoterminationMinutesField != nil {
		st.IdleInstanceAutoterminationMinutes = *idleInstanceAutoterminationMinutesField
	}
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	instancePoolNameField, err := identity(&pb.InstancePoolName)
	if err != nil {
		return nil, err
	}
	if instancePoolNameField != nil {
		st.InstancePoolName = *instancePoolNameField
	}
	maxCapacityField, err := identity(&pb.MaxCapacity)
	if err != nil {
		return nil, err
	}
	if maxCapacityField != nil {
		st.MaxCapacity = *maxCapacityField
	}
	minIdleInstancesField, err := identity(&pb.MinIdleInstances)
	if err != nil {
		return nil, err
	}
	if minIdleInstancesField != nil {
		st.MinIdleInstances = *minIdleInstancesField
	}
	nodeTypeFlexibilityField, err := nodeTypeFlexibilityFromPb(pb.NodeTypeFlexibility)
	if err != nil {
		return nil, err
	}
	if nodeTypeFlexibilityField != nil {
		st.NodeTypeFlexibility = nodeTypeFlexibilityField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}

	var preloadedDockerImagesField []DockerImage
	for _, itemPb := range pb.PreloadedDockerImages {
		item, err := dockerImageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedDockerImagesField = append(preloadedDockerImagesField, *item)
		}
	}
	st.PreloadedDockerImages = preloadedDockerImagesField

	var preloadedSparkVersionsField []string
	for _, itemPb := range pb.PreloadedSparkVersions {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			preloadedSparkVersionsField = append(preloadedSparkVersionsField, *item)
		}
	}
	st.PreloadedSparkVersions = preloadedSparkVersionsField
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	statsField, err := instancePoolStatsFromPb(pb.Stats)
	if err != nil {
		return nil, err
	}
	if statsField != nil {
		st.Stats = statsField
	}
	statusField, err := instancePoolStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAndStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAndStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Attributes set during instance pool creation which are related to Amazon Web
// Services.
type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	Availability InstancePoolAwsAttributesAvailability
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
	SpotBidPricePercent int
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	ZoneId string

	ForceSendFields []string
}

func instancePoolAwsAttributesToPb(st *InstancePoolAwsAttributes) (*instancePoolAwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAwsAttributesPb{}
	availabilityPb, err := identity(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}

	spotBidPricePercentPb, err := identity(&st.SpotBidPricePercent)
	if err != nil {
		return nil, err
	}
	if spotBidPricePercentPb != nil {
		pb.SpotBidPricePercent = *spotBidPricePercentPb
	}

	zoneIdPb, err := identity(&st.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdPb != nil {
		pb.ZoneId = *zoneIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolAwsAttributesPb struct {
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

func instancePoolAwsAttributesFromPb(pb *instancePoolAwsAttributesPb) (*InstancePoolAwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAwsAttributes{}
	availabilityField, err := identity(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	spotBidPricePercentField, err := identity(&pb.SpotBidPricePercent)
	if err != nil {
		return nil, err
	}
	if spotBidPricePercentField != nil {
		st.SpotBidPricePercent = *spotBidPricePercentField
	}
	zoneIdField, err := identity(&pb.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdField != nil {
		st.ZoneId = *zoneIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The set of AWS availability types supported when setting up nodes for a
// cluster.
type InstancePoolAwsAttributesAvailability string
type instancePoolAwsAttributesAvailabilityPb string

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

// Type always returns InstancePoolAwsAttributesAvailability to satisfy [pflag.Value] interface
func (f *InstancePoolAwsAttributesAvailability) Type() string {
	return "InstancePoolAwsAttributesAvailability"
}

func instancePoolAwsAttributesAvailabilityToPb(st *InstancePoolAwsAttributesAvailability) (*instancePoolAwsAttributesAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := instancePoolAwsAttributesAvailabilityPb(*st)
	return &pb, nil
}

func instancePoolAwsAttributesAvailabilityFromPb(pb *instancePoolAwsAttributesAvailabilityPb) (*InstancePoolAwsAttributesAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolAwsAttributesAvailability(*pb)
	return &st, nil
}

// Attributes set during instance pool creation which are related to Azure.
type InstancePoolAzureAttributes struct {
	// Availability type used for the spot nodes.
	Availability InstancePoolAzureAttributesAvailability
	// With variable pricing, you have option to set a max price, in US dollars
	// (USD) For example, the value 2 would be a max price of $2.00 USD per
	// hour. If you set the max price to be -1, the VM won't be evicted based on
	// price. The price for the VM will be the current price for spot or the
	// price for a standard VM, which ever is less, as long as there is capacity
	// and quota available.
	SpotBidMaxPrice float64

	ForceSendFields []string
}

func instancePoolAzureAttributesToPb(st *InstancePoolAzureAttributes) (*instancePoolAzureAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolAzureAttributesPb{}
	availabilityPb, err := identity(&st.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityPb != nil {
		pb.Availability = *availabilityPb
	}

	spotBidMaxPricePb, err := identity(&st.SpotBidMaxPrice)
	if err != nil {
		return nil, err
	}
	if spotBidMaxPricePb != nil {
		pb.SpotBidMaxPrice = *spotBidMaxPricePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolAzureAttributesPb struct {
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

func instancePoolAzureAttributesFromPb(pb *instancePoolAzureAttributesPb) (*InstancePoolAzureAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolAzureAttributes{}
	availabilityField, err := identity(&pb.Availability)
	if err != nil {
		return nil, err
	}
	if availabilityField != nil {
		st.Availability = *availabilityField
	}
	spotBidMaxPriceField, err := identity(&pb.SpotBidMaxPrice)
	if err != nil {
		return nil, err
	}
	if spotBidMaxPriceField != nil {
		st.SpotBidMaxPrice = *spotBidMaxPriceField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolAzureAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolAzureAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The set of Azure availability types supported when setting up nodes for a
// cluster.
type InstancePoolAzureAttributesAvailability string
type instancePoolAzureAttributesAvailabilityPb string

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

// Type always returns InstancePoolAzureAttributesAvailability to satisfy [pflag.Value] interface
func (f *InstancePoolAzureAttributesAvailability) Type() string {
	return "InstancePoolAzureAttributesAvailability"
}

func instancePoolAzureAttributesAvailabilityToPb(st *InstancePoolAzureAttributesAvailability) (*instancePoolAzureAttributesAvailabilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := instancePoolAzureAttributesAvailabilityPb(*st)
	return &pb, nil
}

func instancePoolAzureAttributesAvailabilityFromPb(pb *instancePoolAzureAttributesAvailabilityPb) (*InstancePoolAzureAttributesAvailability, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolAzureAttributesAvailability(*pb)
	return &st, nil
}

// Attributes set during instance pool creation which are related to GCP.
type InstancePoolGcpAttributes struct {
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
	GcpAvailability GcpAvailability
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount int
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
	ZoneId string

	ForceSendFields []string
}

func instancePoolGcpAttributesToPb(st *InstancePoolGcpAttributes) (*instancePoolGcpAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolGcpAttributesPb{}
	gcpAvailabilityPb, err := identity(&st.GcpAvailability)
	if err != nil {
		return nil, err
	}
	if gcpAvailabilityPb != nil {
		pb.GcpAvailability = *gcpAvailabilityPb
	}

	localSsdCountPb, err := identity(&st.LocalSsdCount)
	if err != nil {
		return nil, err
	}
	if localSsdCountPb != nil {
		pb.LocalSsdCount = *localSsdCountPb
	}

	zoneIdPb, err := identity(&st.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdPb != nil {
		pb.ZoneId = *zoneIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolGcpAttributesPb struct {
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
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

func instancePoolGcpAttributesFromPb(pb *instancePoolGcpAttributesPb) (*InstancePoolGcpAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolGcpAttributes{}
	gcpAvailabilityField, err := identity(&pb.GcpAvailability)
	if err != nil {
		return nil, err
	}
	if gcpAvailabilityField != nil {
		st.GcpAvailability = *gcpAvailabilityField
	}
	localSsdCountField, err := identity(&pb.LocalSsdCount)
	if err != nil {
		return nil, err
	}
	if localSsdCountField != nil {
		st.LocalSsdCount = *localSsdCountField
	}
	zoneIdField, err := identity(&pb.ZoneId)
	if err != nil {
		return nil, err
	}
	if zoneIdField != nil {
		st.ZoneId = *zoneIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolGcpAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolGcpAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel InstancePoolPermissionLevel

	ForceSendFields []string
}

func instancePoolPermissionToPb(st *InstancePoolPermission) (*instancePoolPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolPermissionFromPb(pb *instancePoolPermissionPb) (*InstancePoolPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermission{}
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type InstancePoolPermissionLevel string
type instancePoolPermissionLevelPb string

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

// Type always returns InstancePoolPermissionLevel to satisfy [pflag.Value] interface
func (f *InstancePoolPermissionLevel) Type() string {
	return "InstancePoolPermissionLevel"
}

func instancePoolPermissionLevelToPb(st *InstancePoolPermissionLevel) (*instancePoolPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := instancePoolPermissionLevelPb(*st)
	return &pb, nil
}

func instancePoolPermissionLevelFromPb(pb *instancePoolPermissionLevelPb) (*InstancePoolPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolPermissionLevel(*pb)
	return &st, nil
}

type InstancePoolPermissions struct {
	AccessControlList []InstancePoolAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func instancePoolPermissionsToPb(st *InstancePoolPermissions) (*instancePoolPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionsPb{}

	var accessControlListPb []instancePoolAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := instancePoolAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolPermissionsPb struct {
	AccessControlList []instancePoolAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolPermissionsFromPb(pb *instancePoolPermissionsPb) (*InstancePoolPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissions{}

	var accessControlListField []InstancePoolAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := instancePoolAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel InstancePoolPermissionLevel

	ForceSendFields []string
}

func instancePoolPermissionsDescriptionToPb(st *InstancePoolPermissionsDescription) (*instancePoolPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel InstancePoolPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func instancePoolPermissionsDescriptionFromPb(pb *instancePoolPermissionsDescriptionPb) (*InstancePoolPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissionsDescription{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolPermissionsRequest struct {
	AccessControlList []InstancePoolAccessControlRequest
	// The instance pool for which to get or manage permissions.
	InstancePoolId string
}

func instancePoolPermissionsRequestToPb(st *InstancePoolPermissionsRequest) (*instancePoolPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolPermissionsRequestPb{}

	var accessControlListPb []instancePoolAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := instancePoolAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	return pb, nil
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

type instancePoolPermissionsRequestPb struct {
	AccessControlList []instancePoolAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The instance pool for which to get or manage permissions.
	InstancePoolId string `json:"-" url:"-"`
}

func instancePoolPermissionsRequestFromPb(pb *instancePoolPermissionsRequestPb) (*InstancePoolPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolPermissionsRequest{}

	var accessControlListField []InstancePoolAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := instancePoolAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}

	return st, nil
}

// The state of a Cluster. The current allowable state transitions are as
// follows:
//
// - “ACTIVE“ -> “STOPPED“ - “ACTIVE“ -> “DELETED“ - “STOPPED“ ->
// “ACTIVE“ - “STOPPED“ -> “DELETED“
type InstancePoolState string
type instancePoolStatePb string

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

// Type always returns InstancePoolState to satisfy [pflag.Value] interface
func (f *InstancePoolState) Type() string {
	return "InstancePoolState"
}

func instancePoolStateToPb(st *InstancePoolState) (*instancePoolStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := instancePoolStatePb(*st)
	return &pb, nil
}

func instancePoolStateFromPb(pb *instancePoolStatePb) (*InstancePoolState, error) {
	if pb == nil {
		return nil, nil
	}
	st := InstancePoolState(*pb)
	return &st, nil
}

type InstancePoolStats struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	IdleCount int
	// Number of pending instances in the pool that are NOT part of a cluster.
	PendingIdleCount int
	// Number of pending instances in the pool that are part of a cluster.
	PendingUsedCount int
	// Number of active instances in the pool that are part of a cluster.
	UsedCount int

	ForceSendFields []string
}

func instancePoolStatsToPb(st *InstancePoolStats) (*instancePoolStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolStatsPb{}
	idleCountPb, err := identity(&st.IdleCount)
	if err != nil {
		return nil, err
	}
	if idleCountPb != nil {
		pb.IdleCount = *idleCountPb
	}

	pendingIdleCountPb, err := identity(&st.PendingIdleCount)
	if err != nil {
		return nil, err
	}
	if pendingIdleCountPb != nil {
		pb.PendingIdleCount = *pendingIdleCountPb
	}

	pendingUsedCountPb, err := identity(&st.PendingUsedCount)
	if err != nil {
		return nil, err
	}
	if pendingUsedCountPb != nil {
		pb.PendingUsedCount = *pendingUsedCountPb
	}

	usedCountPb, err := identity(&st.UsedCount)
	if err != nil {
		return nil, err
	}
	if usedCountPb != nil {
		pb.UsedCount = *usedCountPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instancePoolStatsPb struct {
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

func instancePoolStatsFromPb(pb *instancePoolStatsPb) (*InstancePoolStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolStats{}
	idleCountField, err := identity(&pb.IdleCount)
	if err != nil {
		return nil, err
	}
	if idleCountField != nil {
		st.IdleCount = *idleCountField
	}
	pendingIdleCountField, err := identity(&pb.PendingIdleCount)
	if err != nil {
		return nil, err
	}
	if pendingIdleCountField != nil {
		st.PendingIdleCount = *pendingIdleCountField
	}
	pendingUsedCountField, err := identity(&pb.PendingUsedCount)
	if err != nil {
		return nil, err
	}
	if pendingUsedCountField != nil {
		st.PendingUsedCount = *pendingUsedCountField
	}
	usedCountField, err := identity(&pb.UsedCount)
	if err != nil {
		return nil, err
	}
	if usedCountField != nil {
		st.UsedCount = *usedCountField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instancePoolStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instancePoolStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors []PendingInstanceError
}

func instancePoolStatusToPb(st *InstancePoolStatus) (*instancePoolStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instancePoolStatusPb{}

	var pendingInstanceErrorsPb []pendingInstanceErrorPb
	for _, item := range st.PendingInstanceErrors {
		itemPb, err := pendingInstanceErrorToPb(&item)
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

type instancePoolStatusPb struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors []pendingInstanceErrorPb `json:"pending_instance_errors,omitempty"`
}

func instancePoolStatusFromPb(pb *instancePoolStatusPb) (*InstancePoolStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstancePoolStatus{}

	var pendingInstanceErrorsField []PendingInstanceError
	for _, itemPb := range pb.PendingInstanceErrors {
		item, err := pendingInstanceErrorFromPb(&itemPb)
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
	IamRoleArn string
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn string
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile bool

	ForceSendFields []string
}

func instanceProfileToPb(st *InstanceProfile) (*instanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &instanceProfilePb{}
	iamRoleArnPb, err := identity(&st.IamRoleArn)
	if err != nil {
		return nil, err
	}
	if iamRoleArnPb != nil {
		pb.IamRoleArn = *iamRoleArnPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	isMetaInstanceProfilePb, err := identity(&st.IsMetaInstanceProfile)
	if err != nil {
		return nil, err
	}
	if isMetaInstanceProfilePb != nil {
		pb.IsMetaInstanceProfile = *isMetaInstanceProfilePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type instanceProfilePb struct {
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

func instanceProfileFromPb(pb *instanceProfilePb) (*InstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InstanceProfile{}
	iamRoleArnField, err := identity(&pb.IamRoleArn)
	if err != nil {
		return nil, err
	}
	if iamRoleArnField != nil {
		st.IamRoleArn = *iamRoleArnField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	isMetaInstanceProfileField, err := identity(&pb.IsMetaInstanceProfile)
	if err != nil {
		return nil, err
	}
	if isMetaInstanceProfileField != nil {
		st.IsMetaInstanceProfile = *isMetaInstanceProfileField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *instanceProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st instanceProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
type kindPb string

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

// Type always returns Kind to satisfy [pflag.Value] interface
func (f *Kind) Type() string {
	return "Kind"
}

func kindToPb(st *Kind) (*kindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := kindPb(*st)
	return &pb, nil
}

func kindFromPb(pb *kindPb) (*Kind, error) {
	if pb == nil {
		return nil, nil
	}
	st := Kind(*pb)
	return &st, nil
}

type Language string
type languagePb string

const LanguagePython Language = `python`

const LanguageScala Language = `scala`

const LanguageSql Language = `sql`

// String representation for [fmt.Print]
func (f *Language) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Language) Set(v string) error {
	switch v {
	case `python`, `scala`, `sql`:
		*f = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "python", "scala", "sql"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (f *Language) Type() string {
	return "Language"
}

func languageToPb(st *Language) (*languagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := languagePb(*st)
	return &pb, nil
}

func languageFromPb(pb *languagePb) (*Language, error) {
	if pb == nil {
		return nil, nil
	}
	st := Language(*pb)
	return &st, nil
}

type Library struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran *RCranLibrary
	// Deprecated. URI of the egg library to install. Installing Python egg
	// files is deprecated and is not supported in Databricks Runtime 14.0 and
	// above.
	Egg string
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Jar string
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	Maven *MavenLibrary
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi *PythonPyPiLibrary
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	Requirements string
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Whl string

	ForceSendFields []string
}

func LibraryToPb(st *Library) (*LibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &LibraryPb{}
	cranPb, err := rCranLibraryToPb(st.Cran)
	if err != nil {
		return nil, err
	}
	if cranPb != nil {
		pb.Cran = cranPb
	}

	eggPb, err := identity(&st.Egg)
	if err != nil {
		return nil, err
	}
	if eggPb != nil {
		pb.Egg = *eggPb
	}

	jarPb, err := identity(&st.Jar)
	if err != nil {
		return nil, err
	}
	if jarPb != nil {
		pb.Jar = *jarPb
	}

	mavenPb, err := MavenLibraryToPb(st.Maven)
	if err != nil {
		return nil, err
	}
	if mavenPb != nil {
		pb.Maven = mavenPb
	}

	pypiPb, err := pythonPyPiLibraryToPb(st.Pypi)
	if err != nil {
		return nil, err
	}
	if pypiPb != nil {
		pb.Pypi = pypiPb
	}

	requirementsPb, err := identity(&st.Requirements)
	if err != nil {
		return nil, err
	}
	if requirementsPb != nil {
		pb.Requirements = *requirementsPb
	}

	whlPb, err := identity(&st.Whl)
	if err != nil {
		return nil, err
	}
	if whlPb != nil {
		pb.Whl = *whlPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Library) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &LibraryPb{}
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

func (st Library) MarshalJSON() ([]byte, error) {
	pb, err := LibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LibraryPb struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran *rCranLibraryPb `json:"cran,omitempty"`
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
	Maven *MavenLibraryPb `json:"maven,omitempty"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi *pythonPyPiLibraryPb `json:"pypi,omitempty"`
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

func LibraryFromPb(pb *LibraryPb) (*Library, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Library{}
	cranField, err := rCranLibraryFromPb(pb.Cran)
	if err != nil {
		return nil, err
	}
	if cranField != nil {
		st.Cran = cranField
	}
	eggField, err := identity(&pb.Egg)
	if err != nil {
		return nil, err
	}
	if eggField != nil {
		st.Egg = *eggField
	}
	jarField, err := identity(&pb.Jar)
	if err != nil {
		return nil, err
	}
	if jarField != nil {
		st.Jar = *jarField
	}
	mavenField, err := MavenLibraryFromPb(pb.Maven)
	if err != nil {
		return nil, err
	}
	if mavenField != nil {
		st.Maven = mavenField
	}
	pypiField, err := pythonPyPiLibraryFromPb(pb.Pypi)
	if err != nil {
		return nil, err
	}
	if pypiField != nil {
		st.Pypi = pypiField
	}
	requirementsField, err := identity(&pb.Requirements)
	if err != nil {
		return nil, err
	}
	if requirementsField != nil {
		st.Requirements = *requirementsField
	}
	whlField, err := identity(&pb.Whl)
	if err != nil {
		return nil, err
	}
	if whlField != nil {
		st.Whl = *whlField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *LibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters bool
	// Unique identifier for the library.
	Library *Library
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages []string
	// Status of installing the library on the cluster.
	Status LibraryInstallStatus

	ForceSendFields []string
}

func libraryFullStatusToPb(st *LibraryFullStatus) (*libraryFullStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &libraryFullStatusPb{}
	isLibraryForAllClustersPb, err := identity(&st.IsLibraryForAllClusters)
	if err != nil {
		return nil, err
	}
	if isLibraryForAllClustersPb != nil {
		pb.IsLibraryForAllClusters = *isLibraryForAllClustersPb
	}

	libraryPb, err := LibraryToPb(st.Library)
	if err != nil {
		return nil, err
	}
	if libraryPb != nil {
		pb.Library = libraryPb
	}

	var messagesPb []string
	for _, item := range st.Messages {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			messagesPb = append(messagesPb, *itemPb)
		}
	}
	pb.Messages = messagesPb

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type libraryFullStatusPb struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters bool `json:"is_library_for_all_clusters,omitempty"`
	// Unique identifier for the library.
	Library *LibraryPb `json:"library,omitempty"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages []string `json:"messages,omitempty"`
	// Status of installing the library on the cluster.
	Status LibraryInstallStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func libraryFullStatusFromPb(pb *libraryFullStatusPb) (*LibraryFullStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LibraryFullStatus{}
	isLibraryForAllClustersField, err := identity(&pb.IsLibraryForAllClusters)
	if err != nil {
		return nil, err
	}
	if isLibraryForAllClustersField != nil {
		st.IsLibraryForAllClusters = *isLibraryForAllClustersField
	}
	libraryField, err := LibraryFromPb(pb.Library)
	if err != nil {
		return nil, err
	}
	if libraryField != nil {
		st.Library = libraryField
	}

	var messagesField []string
	for _, itemPb := range pb.Messages {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			messagesField = append(messagesField, *item)
		}
	}
	st.Messages = messagesField
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *libraryFullStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st libraryFullStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The status of a library on a specific cluster.
type LibraryInstallStatus string
type libraryInstallStatusPb string

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

// Type always returns LibraryInstallStatus to satisfy [pflag.Value] interface
func (f *LibraryInstallStatus) Type() string {
	return "LibraryInstallStatus"
}

func libraryInstallStatusToPb(st *LibraryInstallStatus) (*libraryInstallStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := libraryInstallStatusPb(*st)
	return &pb, nil
}

func libraryInstallStatusFromPb(pb *libraryInstallStatusPb) (*LibraryInstallStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := LibraryInstallStatus(*pb)
	return &st, nil
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses []ClusterLibraryStatuses
}

func listAllClusterLibraryStatusesResponseToPb(st *ListAllClusterLibraryStatusesResponse) (*listAllClusterLibraryStatusesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAllClusterLibraryStatusesResponsePb{}

	var statusesPb []clusterLibraryStatusesPb
	for _, item := range st.Statuses {
		itemPb, err := clusterLibraryStatusesToPb(&item)
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

type listAllClusterLibraryStatusesResponsePb struct {
	// A list of cluster statuses.
	Statuses []clusterLibraryStatusesPb `json:"statuses,omitempty"`
}

func listAllClusterLibraryStatusesResponseFromPb(pb *listAllClusterLibraryStatusesResponsePb) (*ListAllClusterLibraryStatusesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAllClusterLibraryStatusesResponse{}

	var statusesField []ClusterLibraryStatuses
	for _, itemPb := range pb.Statuses {
		item, err := clusterLibraryStatusesFromPb(&itemPb)
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
	DefaultZone string
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones []string

	ForceSendFields []string
}

func listAvailableZonesResponseToPb(st *ListAvailableZonesResponse) (*listAvailableZonesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAvailableZonesResponsePb{}
	defaultZonePb, err := identity(&st.DefaultZone)
	if err != nil {
		return nil, err
	}
	if defaultZonePb != nil {
		pb.DefaultZone = *defaultZonePb
	}

	var zonesPb []string
	for _, item := range st.Zones {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			zonesPb = append(zonesPb, *itemPb)
		}
	}
	pb.Zones = zonesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAvailableZonesResponsePb struct {
	// The availability zone if no ``zone_id`` is provided in the cluster
	// creation request.
	DefaultZone string `json:"default_zone,omitempty"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones []string `json:"zones,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAvailableZonesResponseFromPb(pb *listAvailableZonesResponsePb) (*ListAvailableZonesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAvailableZonesResponse{}
	defaultZoneField, err := identity(&pb.DefaultZone)
	if err != nil {
		return nil, err
	}
	if defaultZoneField != nil {
		st.DefaultZone = *defaultZoneField
	}

	var zonesField []string
	for _, itemPb := range pb.Zones {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			zonesField = append(zonesField, *item)
		}
	}
	st.Zones = zonesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAvailableZonesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAvailableZonesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List cluster policy compliance
type ListClusterCompliancesRequest struct {
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int
	// A page token that can be used to navigate to the next page or previous
	// page as returned by `next_page_token` or `prev_page_token`.
	PageToken string
	// Canonical unique identifier for the cluster policy.
	PolicyId string

	ForceSendFields []string
}

func listClusterCompliancesRequestToPb(st *ListClusterCompliancesRequest) (*listClusterCompliancesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClusterCompliancesRequestPb{}
	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listClusterCompliancesRequestPb struct {
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

func listClusterCompliancesRequestFromPb(pb *listClusterCompliancesRequestPb) (*ListClusterCompliancesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterCompliancesRequest{}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClusterCompliancesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClusterCompliancesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClusterCompliancesResponse struct {
	// A list of clusters and their policy compliance statuses.
	Clusters []ClusterCompliance
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string

	ForceSendFields []string
}

func listClusterCompliancesResponseToPb(st *ListClusterCompliancesResponse) (*listClusterCompliancesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClusterCompliancesResponsePb{}

	var clustersPb []clusterCompliancePb
	for _, item := range st.Clusters {
		itemPb, err := clusterComplianceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	prevPageTokenPb, err := identity(&st.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenPb != nil {
		pb.PrevPageToken = *prevPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listClusterCompliancesResponsePb struct {
	// A list of clusters and their policy compliance statuses.
	Clusters []clusterCompliancePb `json:"clusters,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClusterCompliancesResponseFromPb(pb *listClusterCompliancesResponsePb) (*ListClusterCompliancesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterCompliancesResponse{}

	var clustersField []ClusterCompliance
	for _, itemPb := range pb.Clusters {
		item, err := clusterComplianceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	prevPageTokenField, err := identity(&pb.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenField != nil {
		st.PrevPageToken = *prevPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClusterCompliancesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClusterCompliancesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List cluster policies
type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn ListSortColumn
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder ListSortOrder
}

func listClusterPoliciesRequestToPb(st *ListClusterPoliciesRequest) (*listClusterPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClusterPoliciesRequestPb{}
	sortColumnPb, err := identity(&st.SortColumn)
	if err != nil {
		return nil, err
	}
	if sortColumnPb != nil {
		pb.SortColumn = *sortColumnPb
	}

	sortOrderPb, err := identity(&st.SortOrder)
	if err != nil {
		return nil, err
	}
	if sortOrderPb != nil {
		pb.SortOrder = *sortOrderPb
	}

	return pb, nil
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

type listClusterPoliciesRequestPb struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn ListSortColumn `json:"-" url:"sort_column,omitempty"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder ListSortOrder `json:"-" url:"sort_order,omitempty"`
}

func listClusterPoliciesRequestFromPb(pb *listClusterPoliciesRequestPb) (*ListClusterPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClusterPoliciesRequest{}
	sortColumnField, err := identity(&pb.SortColumn)
	if err != nil {
		return nil, err
	}
	if sortColumnField != nil {
		st.SortColumn = *sortColumnField
	}
	sortOrderField, err := identity(&pb.SortOrder)
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
	ClusterSources []ClusterSource
	// The current state of the clusters.
	ClusterStates []State
	// Whether the clusters are pinned or not.
	IsPinned bool
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string

	ForceSendFields []string
}

func listClustersFilterByToPb(st *ListClustersFilterBy) (*listClustersFilterByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersFilterByPb{}

	var clusterSourcesPb []ClusterSource
	for _, item := range st.ClusterSources {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clusterSourcesPb = append(clusterSourcesPb, *itemPb)
		}
	}
	pb.ClusterSources = clusterSourcesPb

	var clusterStatesPb []State
	for _, item := range st.ClusterStates {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clusterStatesPb = append(clusterStatesPb, *itemPb)
		}
	}
	pb.ClusterStates = clusterStatesPb

	isPinnedPb, err := identity(&st.IsPinned)
	if err != nil {
		return nil, err
	}
	if isPinnedPb != nil {
		pb.IsPinned = *isPinnedPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listClustersFilterByPb struct {
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

func listClustersFilterByFromPb(pb *listClustersFilterByPb) (*ListClustersFilterBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersFilterBy{}

	var clusterSourcesField []ClusterSource
	for _, itemPb := range pb.ClusterSources {
		item, err := identity(&itemPb)
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
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clusterStatesField = append(clusterStatesField, *item)
		}
	}
	st.ClusterStates = clusterStatesField
	isPinnedField, err := identity(&pb.IsPinned)
	if err != nil {
		return nil, err
	}
	if isPinnedField != nil {
		st.IsPinned = *isPinnedField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClustersFilterByPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClustersFilterByPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List clusters
type ListClustersRequest struct {
	// Filters to apply to the list of clusters.
	FilterBy *ListClustersFilterBy
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	PageToken string
	// Sort the list of clusters by a specific criteria.
	SortBy *ListClustersSortBy

	ForceSendFields []string
}

func listClustersRequestToPb(st *ListClustersRequest) (*listClustersRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersRequestPb{}
	filterByPb, err := listClustersFilterByToPb(st.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByPb != nil {
		pb.FilterBy = filterByPb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	sortByPb, err := listClustersSortByToPb(st.SortBy)
	if err != nil {
		return nil, err
	}
	if sortByPb != nil {
		pb.SortBy = sortByPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listClustersRequestPb struct {
	// Filters to apply to the list of clusters.
	FilterBy *listClustersFilterByPb `json:"-" url:"filter_by,omitempty"`
	// Use this field to specify the maximum number of results to be returned by
	// the server. The server may further constrain the maximum number of
	// results returned in a single page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Use next_page_token or prev_page_token returned from the previous request
	// to list the next or previous page of clusters respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Sort the list of clusters by a specific criteria.
	SortBy *listClustersSortByPb `json:"-" url:"sort_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClustersRequestFromPb(pb *listClustersRequestPb) (*ListClustersRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersRequest{}
	filterByField, err := listClustersFilterByFromPb(pb.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByField != nil {
		st.FilterBy = filterByField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	sortByField, err := listClustersSortByFromPb(pb.SortBy)
	if err != nil {
		return nil, err
	}
	if sortByField != nil {
		st.SortBy = sortByField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClustersRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClustersRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClustersResponse struct {
	Clusters []ClusterDetails
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string

	ForceSendFields []string
}

func listClustersResponseToPb(st *ListClustersResponse) (*listClustersResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersResponsePb{}

	var clustersPb []clusterDetailsPb
	for _, item := range st.Clusters {
		itemPb, err := clusterDetailsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			clustersPb = append(clustersPb, *itemPb)
		}
	}
	pb.Clusters = clustersPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	prevPageTokenPb, err := identity(&st.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenPb != nil {
		pb.PrevPageToken = *prevPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listClustersResponsePb struct {
	Clusters []clusterDetailsPb `json:"clusters,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If the value is "", it means no further results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If the value is "", it means no further results for the
	// request.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listClustersResponseFromPb(pb *listClustersResponsePb) (*ListClustersResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersResponse{}

	var clustersField []ClusterDetails
	for _, itemPb := range pb.Clusters {
		item, err := clusterDetailsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			clustersField = append(clustersField, *item)
		}
	}
	st.Clusters = clustersField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	prevPageTokenField, err := identity(&pb.PrevPageToken)
	if err != nil {
		return nil, err
	}
	if prevPageTokenField != nil {
		st.PrevPageToken = *prevPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listClustersResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listClustersResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListClustersSortBy struct {
	// The direction to sort by.
	Direction ListClustersSortByDirection
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field ListClustersSortByField
}

func listClustersSortByToPb(st *ListClustersSortBy) (*listClustersSortByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listClustersSortByPb{}
	directionPb, err := identity(&st.Direction)
	if err != nil {
		return nil, err
	}
	if directionPb != nil {
		pb.Direction = *directionPb
	}

	fieldPb, err := identity(&st.Field)
	if err != nil {
		return nil, err
	}
	if fieldPb != nil {
		pb.Field = *fieldPb
	}

	return pb, nil
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

type listClustersSortByPb struct {
	// The direction to sort by.
	Direction ListClustersSortByDirection `json:"direction,omitempty" url:"direction,omitempty"`
	// The sorting criteria. By default, clusters are sorted by 3 columns from
	// highest to lowest precedence: cluster state, pinned or unpinned, then
	// cluster name.
	Field ListClustersSortByField `json:"field,omitempty" url:"field,omitempty"`
}

func listClustersSortByFromPb(pb *listClustersSortByPb) (*ListClustersSortBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListClustersSortBy{}
	directionField, err := identity(&pb.Direction)
	if err != nil {
		return nil, err
	}
	if directionField != nil {
		st.Direction = *directionField
	}
	fieldField, err := identity(&pb.Field)
	if err != nil {
		return nil, err
	}
	if fieldField != nil {
		st.Field = *fieldField
	}

	return st, nil
}

type ListClustersSortByDirection string
type listClustersSortByDirectionPb string

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

// Type always returns ListClustersSortByDirection to satisfy [pflag.Value] interface
func (f *ListClustersSortByDirection) Type() string {
	return "ListClustersSortByDirection"
}

func listClustersSortByDirectionToPb(st *ListClustersSortByDirection) (*listClustersSortByDirectionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listClustersSortByDirectionPb(*st)
	return &pb, nil
}

func listClustersSortByDirectionFromPb(pb *listClustersSortByDirectionPb) (*ListClustersSortByDirection, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListClustersSortByDirection(*pb)
	return &st, nil
}

type ListClustersSortByField string
type listClustersSortByFieldPb string

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

// Type always returns ListClustersSortByField to satisfy [pflag.Value] interface
func (f *ListClustersSortByField) Type() string {
	return "ListClustersSortByField"
}

func listClustersSortByFieldToPb(st *ListClustersSortByField) (*listClustersSortByFieldPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listClustersSortByFieldPb(*st)
	return &pb, nil
}

func listClustersSortByFieldFromPb(pb *listClustersSortByFieldPb) (*ListClustersSortByField, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListClustersSortByField(*pb)
	return &st, nil
}

type ListGlobalInitScriptsResponse struct {
	Scripts []GlobalInitScriptDetails
}

func listGlobalInitScriptsResponseToPb(st *ListGlobalInitScriptsResponse) (*listGlobalInitScriptsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listGlobalInitScriptsResponsePb{}

	var scriptsPb []globalInitScriptDetailsPb
	for _, item := range st.Scripts {
		itemPb, err := globalInitScriptDetailsToPb(&item)
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

type listGlobalInitScriptsResponsePb struct {
	Scripts []globalInitScriptDetailsPb `json:"scripts,omitempty"`
}

func listGlobalInitScriptsResponseFromPb(pb *listGlobalInitScriptsResponsePb) (*ListGlobalInitScriptsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListGlobalInitScriptsResponse{}

	var scriptsField []GlobalInitScriptDetails
	for _, itemPb := range pb.Scripts {
		item, err := globalInitScriptDetailsFromPb(&itemPb)
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
	InstancePools []InstancePoolAndStats
}

func listInstancePoolsToPb(st *ListInstancePools) (*listInstancePoolsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstancePoolsPb{}

	var instancePoolsPb []instancePoolAndStatsPb
	for _, item := range st.InstancePools {
		itemPb, err := instancePoolAndStatsToPb(&item)
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

type listInstancePoolsPb struct {
	InstancePools []instancePoolAndStatsPb `json:"instance_pools,omitempty"`
}

func listInstancePoolsFromPb(pb *listInstancePoolsPb) (*ListInstancePools, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstancePools{}

	var instancePoolsField []InstancePoolAndStats
	for _, itemPb := range pb.InstancePools {
		item, err := instancePoolAndStatsFromPb(&itemPb)
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
	InstanceProfiles []InstanceProfile
}

func listInstanceProfilesResponseToPb(st *ListInstanceProfilesResponse) (*listInstanceProfilesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listInstanceProfilesResponsePb{}

	var instanceProfilesPb []instanceProfilePb
	for _, item := range st.InstanceProfiles {
		itemPb, err := instanceProfileToPb(&item)
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

type listInstanceProfilesResponsePb struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles []instanceProfilePb `json:"instance_profiles,omitempty"`
}

func listInstanceProfilesResponseFromPb(pb *listInstanceProfilesResponsePb) (*ListInstanceProfilesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListInstanceProfilesResponse{}

	var instanceProfilesField []InstanceProfile
	for _, itemPb := range pb.InstanceProfiles {
		item, err := instanceProfileFromPb(&itemPb)
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
	NodeTypes []NodeType
}

func listNodeTypesResponseToPb(st *ListNodeTypesResponse) (*listNodeTypesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNodeTypesResponsePb{}

	var nodeTypesPb []nodeTypePb
	for _, item := range st.NodeTypes {
		itemPb, err := nodeTypeToPb(&item)
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

type listNodeTypesResponsePb struct {
	// The list of available Spark node types.
	NodeTypes []nodeTypePb `json:"node_types,omitempty"`
}

func listNodeTypesResponseFromPb(pb *listNodeTypesResponsePb) (*ListNodeTypesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNodeTypesResponse{}

	var nodeTypesField []NodeType
	for _, itemPb := range pb.NodeTypes {
		item, err := nodeTypeFromPb(&itemPb)
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
	Policies []Policy
}

func listPoliciesResponseToPb(st *ListPoliciesResponse) (*listPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPoliciesResponsePb{}

	var policiesPb []policyPb
	for _, item := range st.Policies {
		itemPb, err := policyToPb(&item)
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

type listPoliciesResponsePb struct {
	// List of policies.
	Policies []policyPb `json:"policies,omitempty"`
}

func listPoliciesResponseFromPb(pb *listPoliciesResponsePb) (*ListPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPoliciesResponse{}

	var policiesField []Policy
	for _, itemPb := range pb.Policies {
		item, err := policyFromPb(&itemPb)
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

// List policy families
type ListPolicyFamiliesRequest struct {
	// Maximum number of policy families to return.
	MaxResults int64
	// A token that can be used to get the next page of results.
	PageToken string

	ForceSendFields []string
}

func listPolicyFamiliesRequestToPb(st *ListPolicyFamiliesRequest) (*listPolicyFamiliesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPolicyFamiliesRequestPb{}
	maxResultsPb, err := identity(&st.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsPb != nil {
		pb.MaxResults = *maxResultsPb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPolicyFamiliesRequestPb struct {
	// Maximum number of policy families to return.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPolicyFamiliesRequestFromPb(pb *listPolicyFamiliesRequestPb) (*ListPolicyFamiliesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPolicyFamiliesRequest{}
	maxResultsField, err := identity(&pb.MaxResults)
	if err != nil {
		return nil, err
	}
	if maxResultsField != nil {
		st.MaxResults = *maxResultsField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPolicyFamiliesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPolicyFamiliesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string
	// List of policy families.
	PolicyFamilies []PolicyFamily

	ForceSendFields []string
}

func listPolicyFamiliesResponseToPb(st *ListPolicyFamiliesResponse) (*listPolicyFamiliesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPolicyFamiliesResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var policyFamiliesPb []policyFamilyPb
	for _, item := range st.PolicyFamilies {
		itemPb, err := policyFamilyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policyFamiliesPb = append(policyFamiliesPb, *itemPb)
		}
	}
	pb.PolicyFamilies = policyFamiliesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPolicyFamiliesResponsePb struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of policy families.
	PolicyFamilies []policyFamilyPb `json:"policy_families,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPolicyFamiliesResponseFromPb(pb *listPolicyFamiliesResponsePb) (*ListPolicyFamiliesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPolicyFamiliesResponse{}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var policyFamiliesField []PolicyFamily
	for _, itemPb := range pb.PolicyFamilies {
		item, err := policyFamilyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policyFamiliesField = append(policyFamiliesField, *item)
		}
	}
	st.PolicyFamilies = policyFamiliesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPolicyFamiliesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPolicyFamiliesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListSortColumn string
type listSortColumnPb string

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

// Type always returns ListSortColumn to satisfy [pflag.Value] interface
func (f *ListSortColumn) Type() string {
	return "ListSortColumn"
}

func listSortColumnToPb(st *ListSortColumn) (*listSortColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listSortColumnPb(*st)
	return &pb, nil
}

func listSortColumnFromPb(pb *listSortColumnPb) (*ListSortColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListSortColumn(*pb)
	return &st, nil
}

type ListSortOrder string
type listSortOrderPb string

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

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (f *ListSortOrder) Type() string {
	return "ListSortOrder"
}

func listSortOrderToPb(st *ListSortOrder) (*listSortOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listSortOrderPb(*st)
	return &pb, nil
}

func listSortOrderFromPb(pb *listSortOrderPb) (*ListSortOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListSortOrder(*pb)
	return &st, nil
}

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination string
}

func localFileInfoToPb(st *LocalFileInfo) (*localFileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &localFileInfoPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	return pb, nil
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

type localFileInfoPb struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination string `json:"destination"`
}

func localFileInfoFromPb(pb *localFileInfoPb) (*LocalFileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LocalFileInfo{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}

	return st, nil
}

type LogAnalyticsInfo struct {
	LogAnalyticsPrimaryKey string

	LogAnalyticsWorkspaceId string

	ForceSendFields []string
}

func logAnalyticsInfoToPb(st *LogAnalyticsInfo) (*logAnalyticsInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logAnalyticsInfoPb{}
	logAnalyticsPrimaryKeyPb, err := identity(&st.LogAnalyticsPrimaryKey)
	if err != nil {
		return nil, err
	}
	if logAnalyticsPrimaryKeyPb != nil {
		pb.LogAnalyticsPrimaryKey = *logAnalyticsPrimaryKeyPb
	}

	logAnalyticsWorkspaceIdPb, err := identity(&st.LogAnalyticsWorkspaceId)
	if err != nil {
		return nil, err
	}
	if logAnalyticsWorkspaceIdPb != nil {
		pb.LogAnalyticsWorkspaceId = *logAnalyticsWorkspaceIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type logAnalyticsInfoPb struct {
	LogAnalyticsPrimaryKey string `json:"log_analytics_primary_key,omitempty"`

	LogAnalyticsWorkspaceId string `json:"log_analytics_workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logAnalyticsInfoFromPb(pb *logAnalyticsInfoPb) (*LogAnalyticsInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogAnalyticsInfo{}
	logAnalyticsPrimaryKeyField, err := identity(&pb.LogAnalyticsPrimaryKey)
	if err != nil {
		return nil, err
	}
	if logAnalyticsPrimaryKeyField != nil {
		st.LogAnalyticsPrimaryKey = *logAnalyticsPrimaryKeyField
	}
	logAnalyticsWorkspaceIdField, err := identity(&pb.LogAnalyticsWorkspaceId)
	if err != nil {
		return nil, err
	}
	if logAnalyticsWorkspaceIdField != nil {
		st.LogAnalyticsWorkspaceId = *logAnalyticsWorkspaceIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logAnalyticsInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logAnalyticsInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The log delivery status
type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted int64
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException string

	ForceSendFields []string
}

func logSyncStatusToPb(st *LogSyncStatus) (*logSyncStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logSyncStatusPb{}
	lastAttemptedPb, err := identity(&st.LastAttempted)
	if err != nil {
		return nil, err
	}
	if lastAttemptedPb != nil {
		pb.LastAttempted = *lastAttemptedPb
	}

	lastExceptionPb, err := identity(&st.LastException)
	if err != nil {
		return nil, err
	}
	if lastExceptionPb != nil {
		pb.LastException = *lastExceptionPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type logSyncStatusPb struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted int64 `json:"last_attempted,omitempty"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException string `json:"last_exception,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logSyncStatusFromPb(pb *logSyncStatusPb) (*LogSyncStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogSyncStatus{}
	lastAttemptedField, err := identity(&pb.LastAttempted)
	if err != nil {
		return nil, err
	}
	if lastAttemptedField != nil {
		st.LastAttempted = *lastAttemptedField
	}
	lastExceptionField, err := identity(&pb.LastException)
	if err != nil {
		return nil, err
	}
	if lastExceptionField != nil {
		st.LastException = *lastExceptionField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logSyncStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logSyncStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MapAny map[string]any
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

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	Coordinates string
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions []string
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo string

	ForceSendFields []string
}

func MavenLibraryToPb(st *MavenLibrary) (*MavenLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &MavenLibraryPb{}
	coordinatesPb, err := identity(&st.Coordinates)
	if err != nil {
		return nil, err
	}
	if coordinatesPb != nil {
		pb.Coordinates = *coordinatesPb
	}

	var exclusionsPb []string
	for _, item := range st.Exclusions {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			exclusionsPb = append(exclusionsPb, *itemPb)
		}
	}
	pb.Exclusions = exclusionsPb

	repoPb, err := identity(&st.Repo)
	if err != nil {
		return nil, err
	}
	if repoPb != nil {
		pb.Repo = *repoPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *MavenLibrary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &MavenLibraryPb{}
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

func (st MavenLibrary) MarshalJSON() ([]byte, error) {
	pb, err := MavenLibraryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MavenLibraryPb struct {
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

func MavenLibraryFromPb(pb *MavenLibraryPb) (*MavenLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MavenLibrary{}
	coordinatesField, err := identity(&pb.Coordinates)
	if err != nil {
		return nil, err
	}
	if coordinatesField != nil {
		st.Coordinates = *coordinatesField
	}

	var exclusionsField []string
	for _, itemPb := range pb.Exclusions {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			exclusionsField = append(exclusionsField, *item)
		}
	}
	st.Exclusions = exclusionsField
	repoField, err := identity(&pb.Repo)
	if err != nil {
		return nil, err
	}
	if repoField != nil {
		st.Repo = *repoField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *MavenLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MavenLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// This structure embodies the machine type that hosts spark containers Note:
// this should be an internal data structure for now It is defined in proto in
// case we want to send it over the wire in the future (which is likely)
type NodeInstanceType struct {
	// Unique identifier across instance types
	InstanceTypeId string
	// Size of the individual local disks attached to this instance (i.e. per
	// local disk).
	LocalDiskSizeGb int
	// Number of local disks that are present on this instance.
	LocalDisks int
	// Size of the individual local nvme disks attached to this instance (i.e.
	// per local disk).
	LocalNvmeDiskSizeGb int
	// Number of local nvme disks that are present on this instance.
	LocalNvmeDisks int

	ForceSendFields []string
}

func nodeInstanceTypeToPb(st *NodeInstanceType) (*nodeInstanceTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nodeInstanceTypePb{}
	instanceTypeIdPb, err := identity(&st.InstanceTypeId)
	if err != nil {
		return nil, err
	}
	if instanceTypeIdPb != nil {
		pb.InstanceTypeId = *instanceTypeIdPb
	}

	localDiskSizeGbPb, err := identity(&st.LocalDiskSizeGb)
	if err != nil {
		return nil, err
	}
	if localDiskSizeGbPb != nil {
		pb.LocalDiskSizeGb = *localDiskSizeGbPb
	}

	localDisksPb, err := identity(&st.LocalDisks)
	if err != nil {
		return nil, err
	}
	if localDisksPb != nil {
		pb.LocalDisks = *localDisksPb
	}

	localNvmeDiskSizeGbPb, err := identity(&st.LocalNvmeDiskSizeGb)
	if err != nil {
		return nil, err
	}
	if localNvmeDiskSizeGbPb != nil {
		pb.LocalNvmeDiskSizeGb = *localNvmeDiskSizeGbPb
	}

	localNvmeDisksPb, err := identity(&st.LocalNvmeDisks)
	if err != nil {
		return nil, err
	}
	if localNvmeDisksPb != nil {
		pb.LocalNvmeDisks = *localNvmeDisksPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type nodeInstanceTypePb struct {
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

func nodeInstanceTypeFromPb(pb *nodeInstanceTypePb) (*NodeInstanceType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NodeInstanceType{}
	instanceTypeIdField, err := identity(&pb.InstanceTypeId)
	if err != nil {
		return nil, err
	}
	if instanceTypeIdField != nil {
		st.InstanceTypeId = *instanceTypeIdField
	}
	localDiskSizeGbField, err := identity(&pb.LocalDiskSizeGb)
	if err != nil {
		return nil, err
	}
	if localDiskSizeGbField != nil {
		st.LocalDiskSizeGb = *localDiskSizeGbField
	}
	localDisksField, err := identity(&pb.LocalDisks)
	if err != nil {
		return nil, err
	}
	if localDisksField != nil {
		st.LocalDisks = *localDisksField
	}
	localNvmeDiskSizeGbField, err := identity(&pb.LocalNvmeDiskSizeGb)
	if err != nil {
		return nil, err
	}
	if localNvmeDiskSizeGbField != nil {
		st.LocalNvmeDiskSizeGb = *localNvmeDiskSizeGbField
	}
	localNvmeDisksField, err := identity(&pb.LocalNvmeDisks)
	if err != nil {
		return nil, err
	}
	if localNvmeDisksField != nil {
		st.LocalNvmeDisks = *localNvmeDisksField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nodeInstanceTypePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nodeInstanceTypePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A description of a Spark node type including both the dimensions of the node
// and the instance type on which it will be hosted.
type NodeType struct {
	// A descriptive category for this node type. Examples include "Memory
	// Optimized" and "Compute Optimized".
	Category string
	// A string description associated with this node type, e.g., "r3.xlarge".
	Description string
	// An optional hint at the display order of node types in the UI. Within a
	// node type category, lowest numbers come first.
	DisplayOrder int
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	InstanceTypeId string
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	IsDeprecated bool
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	IsEncryptedInTransit bool
	// Whether this is an Arm-based instance.
	IsGraviton bool
	// Whether this node is hidden from presentation in the UI.
	IsHidden bool
	// Whether this node comes with IO cache enabled by default.
	IsIoCacheEnabled bool
	// Memory (in MB) available for this node type.
	MemoryMb int
	// A collection of node type info reported by the cloud provider
	NodeInfo *CloudProviderNodeInfo
	// The NodeInstanceType object corresponding to instance_type_id
	NodeInstanceType *NodeInstanceType
	// Unique identifier for this node type.
	NodeTypeId string
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	NumCores float64
	// Number of GPUs available for this node type.
	NumGpus int

	PhotonDriverCapable bool

	PhotonWorkerCapable bool
	// Whether this node type support cluster tags.
	SupportClusterTags bool
	// Whether this node type support EBS volumes. EBS volumes is disabled for
	// node types that we could place multiple corresponding containers on the
	// same hosting instance.
	SupportEbsVolumes bool
	// Whether this node type supports port forwarding.
	SupportPortForwarding bool

	ForceSendFields []string
}

func nodeTypeToPb(st *NodeType) (*nodeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nodeTypePb{}
	categoryPb, err := identity(&st.Category)
	if err != nil {
		return nil, err
	}
	if categoryPb != nil {
		pb.Category = *categoryPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	displayOrderPb, err := identity(&st.DisplayOrder)
	if err != nil {
		return nil, err
	}
	if displayOrderPb != nil {
		pb.DisplayOrder = *displayOrderPb
	}

	instanceTypeIdPb, err := identity(&st.InstanceTypeId)
	if err != nil {
		return nil, err
	}
	if instanceTypeIdPb != nil {
		pb.InstanceTypeId = *instanceTypeIdPb
	}

	isDeprecatedPb, err := identity(&st.IsDeprecated)
	if err != nil {
		return nil, err
	}
	if isDeprecatedPb != nil {
		pb.IsDeprecated = *isDeprecatedPb
	}

	isEncryptedInTransitPb, err := identity(&st.IsEncryptedInTransit)
	if err != nil {
		return nil, err
	}
	if isEncryptedInTransitPb != nil {
		pb.IsEncryptedInTransit = *isEncryptedInTransitPb
	}

	isGravitonPb, err := identity(&st.IsGraviton)
	if err != nil {
		return nil, err
	}
	if isGravitonPb != nil {
		pb.IsGraviton = *isGravitonPb
	}

	isHiddenPb, err := identity(&st.IsHidden)
	if err != nil {
		return nil, err
	}
	if isHiddenPb != nil {
		pb.IsHidden = *isHiddenPb
	}

	isIoCacheEnabledPb, err := identity(&st.IsIoCacheEnabled)
	if err != nil {
		return nil, err
	}
	if isIoCacheEnabledPb != nil {
		pb.IsIoCacheEnabled = *isIoCacheEnabledPb
	}

	memoryMbPb, err := identity(&st.MemoryMb)
	if err != nil {
		return nil, err
	}
	if memoryMbPb != nil {
		pb.MemoryMb = *memoryMbPb
	}

	nodeInfoPb, err := cloudProviderNodeInfoToPb(st.NodeInfo)
	if err != nil {
		return nil, err
	}
	if nodeInfoPb != nil {
		pb.NodeInfo = nodeInfoPb
	}

	nodeInstanceTypePb, err := nodeInstanceTypeToPb(st.NodeInstanceType)
	if err != nil {
		return nil, err
	}
	if nodeInstanceTypePb != nil {
		pb.NodeInstanceType = nodeInstanceTypePb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numCoresPb, err := identity(&st.NumCores)
	if err != nil {
		return nil, err
	}
	if numCoresPb != nil {
		pb.NumCores = *numCoresPb
	}

	numGpusPb, err := identity(&st.NumGpus)
	if err != nil {
		return nil, err
	}
	if numGpusPb != nil {
		pb.NumGpus = *numGpusPb
	}

	photonDriverCapablePb, err := identity(&st.PhotonDriverCapable)
	if err != nil {
		return nil, err
	}
	if photonDriverCapablePb != nil {
		pb.PhotonDriverCapable = *photonDriverCapablePb
	}

	photonWorkerCapablePb, err := identity(&st.PhotonWorkerCapable)
	if err != nil {
		return nil, err
	}
	if photonWorkerCapablePb != nil {
		pb.PhotonWorkerCapable = *photonWorkerCapablePb
	}

	supportClusterTagsPb, err := identity(&st.SupportClusterTags)
	if err != nil {
		return nil, err
	}
	if supportClusterTagsPb != nil {
		pb.SupportClusterTags = *supportClusterTagsPb
	}

	supportEbsVolumesPb, err := identity(&st.SupportEbsVolumes)
	if err != nil {
		return nil, err
	}
	if supportEbsVolumesPb != nil {
		pb.SupportEbsVolumes = *supportEbsVolumesPb
	}

	supportPortForwardingPb, err := identity(&st.SupportPortForwarding)
	if err != nil {
		return nil, err
	}
	if supportPortForwardingPb != nil {
		pb.SupportPortForwarding = *supportPortForwardingPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type nodeTypePb struct {
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
	NodeInfo *cloudProviderNodeInfoPb `json:"node_info,omitempty"`
	// The NodeInstanceType object corresponding to instance_type_id
	NodeInstanceType *nodeInstanceTypePb `json:"node_instance_type,omitempty"`
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

func nodeTypeFromPb(pb *nodeTypePb) (*NodeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NodeType{}
	categoryField, err := identity(&pb.Category)
	if err != nil {
		return nil, err
	}
	if categoryField != nil {
		st.Category = *categoryField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	displayOrderField, err := identity(&pb.DisplayOrder)
	if err != nil {
		return nil, err
	}
	if displayOrderField != nil {
		st.DisplayOrder = *displayOrderField
	}
	instanceTypeIdField, err := identity(&pb.InstanceTypeId)
	if err != nil {
		return nil, err
	}
	if instanceTypeIdField != nil {
		st.InstanceTypeId = *instanceTypeIdField
	}
	isDeprecatedField, err := identity(&pb.IsDeprecated)
	if err != nil {
		return nil, err
	}
	if isDeprecatedField != nil {
		st.IsDeprecated = *isDeprecatedField
	}
	isEncryptedInTransitField, err := identity(&pb.IsEncryptedInTransit)
	if err != nil {
		return nil, err
	}
	if isEncryptedInTransitField != nil {
		st.IsEncryptedInTransit = *isEncryptedInTransitField
	}
	isGravitonField, err := identity(&pb.IsGraviton)
	if err != nil {
		return nil, err
	}
	if isGravitonField != nil {
		st.IsGraviton = *isGravitonField
	}
	isHiddenField, err := identity(&pb.IsHidden)
	if err != nil {
		return nil, err
	}
	if isHiddenField != nil {
		st.IsHidden = *isHiddenField
	}
	isIoCacheEnabledField, err := identity(&pb.IsIoCacheEnabled)
	if err != nil {
		return nil, err
	}
	if isIoCacheEnabledField != nil {
		st.IsIoCacheEnabled = *isIoCacheEnabledField
	}
	memoryMbField, err := identity(&pb.MemoryMb)
	if err != nil {
		return nil, err
	}
	if memoryMbField != nil {
		st.MemoryMb = *memoryMbField
	}
	nodeInfoField, err := cloudProviderNodeInfoFromPb(pb.NodeInfo)
	if err != nil {
		return nil, err
	}
	if nodeInfoField != nil {
		st.NodeInfo = nodeInfoField
	}
	nodeInstanceTypeField, err := nodeInstanceTypeFromPb(pb.NodeInstanceType)
	if err != nil {
		return nil, err
	}
	if nodeInstanceTypeField != nil {
		st.NodeInstanceType = nodeInstanceTypeField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numCoresField, err := identity(&pb.NumCores)
	if err != nil {
		return nil, err
	}
	if numCoresField != nil {
		st.NumCores = *numCoresField
	}
	numGpusField, err := identity(&pb.NumGpus)
	if err != nil {
		return nil, err
	}
	if numGpusField != nil {
		st.NumGpus = *numGpusField
	}
	photonDriverCapableField, err := identity(&pb.PhotonDriverCapable)
	if err != nil {
		return nil, err
	}
	if photonDriverCapableField != nil {
		st.PhotonDriverCapable = *photonDriverCapableField
	}
	photonWorkerCapableField, err := identity(&pb.PhotonWorkerCapable)
	if err != nil {
		return nil, err
	}
	if photonWorkerCapableField != nil {
		st.PhotonWorkerCapable = *photonWorkerCapableField
	}
	supportClusterTagsField, err := identity(&pb.SupportClusterTags)
	if err != nil {
		return nil, err
	}
	if supportClusterTagsField != nil {
		st.SupportClusterTags = *supportClusterTagsField
	}
	supportEbsVolumesField, err := identity(&pb.SupportEbsVolumes)
	if err != nil {
		return nil, err
	}
	if supportEbsVolumesField != nil {
		st.SupportEbsVolumes = *supportEbsVolumesField
	}
	supportPortForwardingField, err := identity(&pb.SupportPortForwarding)
	if err != nil {
		return nil, err
	}
	if supportPortForwardingField != nil {
		st.SupportPortForwarding = *supportPortForwardingField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nodeTypePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nodeTypePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// For Fleet-V2 using classic clusters, this object contains the information
// about the alternate node type ids to use when attempting to launch a cluster.
// It can be used with both the driver and worker node types.
type NodeTypeFlexibility struct {
}

func nodeTypeFlexibilityToPb(st *NodeTypeFlexibility) (*nodeTypeFlexibilityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nodeTypeFlexibilityPb{}

	return pb, nil
}

func (st *NodeTypeFlexibility) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nodeTypeFlexibilityPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nodeTypeFlexibilityFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NodeTypeFlexibility) MarshalJSON() ([]byte, error) {
	pb, err := nodeTypeFlexibilityToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type nodeTypeFlexibilityPb struct {
}

func nodeTypeFlexibilityFromPb(pb *nodeTypeFlexibilityPb) (*NodeTypeFlexibility, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NodeTypeFlexibility{}

	return st, nil
}

// Error message of a failed pending instances
type PendingInstanceError struct {
	InstanceId string

	Message string

	ForceSendFields []string
}

func pendingInstanceErrorToPb(st *PendingInstanceError) (*pendingInstanceErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pendingInstanceErrorPb{}
	instanceIdPb, err := identity(&st.InstanceId)
	if err != nil {
		return nil, err
	}
	if instanceIdPb != nil {
		pb.InstanceId = *instanceIdPb
	}

	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pendingInstanceErrorPb struct {
	InstanceId string `json:"instance_id,omitempty"`

	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pendingInstanceErrorFromPb(pb *pendingInstanceErrorPb) (*PendingInstanceError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PendingInstanceError{}
	instanceIdField, err := identity(&pb.InstanceId)
	if err != nil {
		return nil, err
	}
	if instanceIdField != nil {
		st.InstanceId = *instanceIdField
	}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pendingInstanceErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pendingInstanceErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PermanentDeleteCluster struct {
	// The cluster to be deleted.
	ClusterId string
}

func permanentDeleteClusterToPb(st *PermanentDeleteCluster) (*permanentDeleteClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permanentDeleteClusterPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type permanentDeleteClusterPb struct {
	// The cluster to be deleted.
	ClusterId string `json:"cluster_id"`
}

func permanentDeleteClusterFromPb(pb *permanentDeleteClusterPb) (*PermanentDeleteCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermanentDeleteCluster{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type PermanentDeleteClusterResponse struct {
}

func permanentDeleteClusterResponseToPb(st *PermanentDeleteClusterResponse) (*permanentDeleteClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &permanentDeleteClusterResponsePb{}

	return pb, nil
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

type permanentDeleteClusterResponsePb struct {
}

func permanentDeleteClusterResponseFromPb(pb *permanentDeleteClusterResponsePb) (*PermanentDeleteClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PermanentDeleteClusterResponse{}

	return st, nil
}

type PinCluster struct {
	ClusterId string
}

func pinClusterToPb(st *PinCluster) (*pinClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pinClusterPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type pinClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func pinClusterFromPb(pb *pinClusterPb) (*PinCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PinCluster{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type PinClusterResponse struct {
}

func pinClusterResponseToPb(st *PinClusterResponse) (*pinClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pinClusterResponsePb{}

	return pb, nil
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

type pinClusterResponsePb struct {
}

func pinClusterResponseFromPb(pb *pinClusterResponsePb) (*PinClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PinClusterResponse{}

	return st, nil
}

// Describes a Cluster Policy entity.
type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	CreatedAtTimestamp int64
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName string
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string
	// Additional human-readable description of the cluster policy.
	Description string
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	IsDefault bool
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser int64
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name string
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides string
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId string
	// Canonical unique identifier for the Cluster Policy.
	PolicyId string

	ForceSendFields []string
}

func policyToPb(st *Policy) (*policyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &policyPb{}
	createdAtTimestampPb, err := identity(&st.CreatedAtTimestamp)
	if err != nil {
		return nil, err
	}
	if createdAtTimestampPb != nil {
		pb.CreatedAtTimestamp = *createdAtTimestampPb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	definitionPb, err := identity(&st.Definition)
	if err != nil {
		return nil, err
	}
	if definitionPb != nil {
		pb.Definition = *definitionPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	isDefaultPb, err := identity(&st.IsDefault)
	if err != nil {
		return nil, err
	}
	if isDefaultPb != nil {
		pb.IsDefault = *isDefaultPb
	}

	var librariesPb []LibraryPb
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

	maxClustersPerUserPb, err := identity(&st.MaxClustersPerUser)
	if err != nil {
		return nil, err
	}
	if maxClustersPerUserPb != nil {
		pb.MaxClustersPerUser = *maxClustersPerUserPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	policyFamilyDefinitionOverridesPb, err := identity(&st.PolicyFamilyDefinitionOverrides)
	if err != nil {
		return nil, err
	}
	if policyFamilyDefinitionOverridesPb != nil {
		pb.PolicyFamilyDefinitionOverrides = *policyFamilyDefinitionOverridesPb
	}

	policyFamilyIdPb, err := identity(&st.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdPb != nil {
		pb.PolicyFamilyId = *policyFamilyIdPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type policyPb struct {
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
	Libraries []LibraryPb `json:"libraries,omitempty"`
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

func policyFromPb(pb *policyPb) (*Policy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Policy{}
	createdAtTimestampField, err := identity(&pb.CreatedAtTimestamp)
	if err != nil {
		return nil, err
	}
	if createdAtTimestampField != nil {
		st.CreatedAtTimestamp = *createdAtTimestampField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	definitionField, err := identity(&pb.Definition)
	if err != nil {
		return nil, err
	}
	if definitionField != nil {
		st.Definition = *definitionField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	isDefaultField, err := identity(&pb.IsDefault)
	if err != nil {
		return nil, err
	}
	if isDefaultField != nil {
		st.IsDefault = *isDefaultField
	}

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
	maxClustersPerUserField, err := identity(&pb.MaxClustersPerUser)
	if err != nil {
		return nil, err
	}
	if maxClustersPerUserField != nil {
		st.MaxClustersPerUser = *maxClustersPerUserField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	policyFamilyDefinitionOverridesField, err := identity(&pb.PolicyFamilyDefinitionOverrides)
	if err != nil {
		return nil, err
	}
	if policyFamilyDefinitionOverridesField != nil {
		st.PolicyFamilyDefinitionOverrides = *policyFamilyDefinitionOverridesField
	}
	policyFamilyIdField, err := identity(&pb.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdField != nil {
		st.PolicyFamilyId = *policyFamilyIdField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *policyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st policyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition string
	// Human-readable description of the purpose of the policy family.
	Description string
	// Name of the policy family.
	Name string
	// Unique identifier for the policy family.
	PolicyFamilyId string

	ForceSendFields []string
}

func policyFamilyToPb(st *PolicyFamily) (*policyFamilyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &policyFamilyPb{}
	definitionPb, err := identity(&st.Definition)
	if err != nil {
		return nil, err
	}
	if definitionPb != nil {
		pb.Definition = *definitionPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	policyFamilyIdPb, err := identity(&st.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdPb != nil {
		pb.PolicyFamilyId = *policyFamilyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type policyFamilyPb struct {
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

func policyFamilyFromPb(pb *policyFamilyPb) (*PolicyFamily, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PolicyFamily{}
	definitionField, err := identity(&pb.Definition)
	if err != nil {
		return nil, err
	}
	if definitionField != nil {
		st.Definition = *definitionField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	policyFamilyIdField, err := identity(&pb.PolicyFamilyId)
	if err != nil {
		return nil, err
	}
	if policyFamilyIdField != nil {
		st.PolicyFamilyId = *policyFamilyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *policyFamilyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st policyFamilyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package string
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo string

	ForceSendFields []string
}

func pythonPyPiLibraryToPb(st *PythonPyPiLibrary) (*pythonPyPiLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pythonPyPiLibraryPb{}
	packagePb, err := identity(&st.Package)
	if err != nil {
		return nil, err
	}
	if packagePb != nil {
		pb.Package = *packagePb
	}

	repoPb, err := identity(&st.Repo)
	if err != nil {
		return nil, err
	}
	if repoPb != nil {
		pb.Repo = *repoPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pythonPyPiLibraryPb struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pythonPyPiLibraryFromPb(pb *pythonPyPiLibraryPb) (*PythonPyPiLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PythonPyPiLibrary{}
	packageField, err := identity(&pb.Package)
	if err != nil {
		return nil, err
	}
	if packageField != nil {
		st.Package = *packageField
	}
	repoField, err := identity(&pb.Repo)
	if err != nil {
		return nil, err
	}
	if repoField != nil {
		st.Repo = *repoField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pythonPyPiLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pythonPyPiLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	Package string
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo string

	ForceSendFields []string
}

func rCranLibraryToPb(st *RCranLibrary) (*rCranLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rCranLibraryPb{}
	packagePb, err := identity(&st.Package)
	if err != nil {
		return nil, err
	}
	if packagePb != nil {
		pb.Package = *packagePb
	}

	repoPb, err := identity(&st.Repo)
	if err != nil {
		return nil, err
	}
	if repoPb != nil {
		pb.Repo = *repoPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type rCranLibraryPb struct {
	// The name of the CRAN package to install.
	Package string `json:"package"`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo string `json:"repo,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func rCranLibraryFromPb(pb *rCranLibraryPb) (*RCranLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RCranLibrary{}
	packageField, err := identity(&pb.Package)
	if err != nil {
		return nil, err
	}
	if packageField != nil {
		st.Package = *packageField
	}
	repoField, err := identity(&pb.Repo)
	if err != nil {
		return nil, err
	}
	if repoField != nil {
		st.Repo = *repoField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *rCranLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st rCranLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	InstanceProfileArn string
}

func removeInstanceProfileToPb(st *RemoveInstanceProfile) (*removeInstanceProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeInstanceProfilePb{}
	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	return pb, nil
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

type removeInstanceProfilePb struct {
	// The ARN of the instance profile to remove. This field is required.
	InstanceProfileArn string `json:"instance_profile_arn"`
}

func removeInstanceProfileFromPb(pb *removeInstanceProfilePb) (*RemoveInstanceProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveInstanceProfile{}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}

	return st, nil
}

type RemoveResponse struct {
}

func removeResponseToPb(st *RemoveResponse) (*removeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &removeResponsePb{}

	return pb, nil
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

type removeResponsePb struct {
}

func removeResponseFromPb(pb *removeResponsePb) (*RemoveResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RemoveResponse{}

	return st, nil
}

type ResizeCluster struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
	// The cluster to be resized.
	ClusterId string
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
	NumWorkers int

	ForceSendFields []string
}

func resizeClusterToPb(st *ResizeCluster) (*resizeClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resizeClusterPb{}
	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resizeClusterPb struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
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

func resizeClusterFromPb(pb *resizeClusterPb) (*ResizeCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResizeCluster{}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resizeClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resizeClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResizeClusterResponse struct {
}

func resizeClusterResponseToPb(st *ResizeClusterResponse) (*resizeClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resizeClusterResponsePb{}

	return pb, nil
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

type resizeClusterResponsePb struct {
}

func resizeClusterResponseFromPb(pb *resizeClusterResponsePb) (*ResizeClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResizeClusterResponse{}

	return st, nil
}

type RestartCluster struct {
	// The cluster to be started.
	ClusterId string

	RestartUser string

	ForceSendFields []string
}

func restartClusterToPb(st *RestartCluster) (*restartClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restartClusterPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	restartUserPb, err := identity(&st.RestartUser)
	if err != nil {
		return nil, err
	}
	if restartUserPb != nil {
		pb.RestartUser = *restartUserPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type restartClusterPb struct {
	// The cluster to be started.
	ClusterId string `json:"cluster_id"`

	RestartUser string `json:"restart_user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restartClusterFromPb(pb *restartClusterPb) (*RestartCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartCluster{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	restartUserField, err := identity(&pb.RestartUser)
	if err != nil {
		return nil, err
	}
	if restartUserField != nil {
		st.RestartUser = *restartUserField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restartClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restartClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RestartClusterResponse struct {
}

func restartClusterResponseToPb(st *RestartClusterResponse) (*restartClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restartClusterResponsePb{}

	return pb, nil
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

type restartClusterResponsePb struct {
}

func restartClusterResponseFromPb(pb *restartClusterResponsePb) (*RestartClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartClusterResponse{}

	return st, nil
}

type ResultType string
type resultTypePb string

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

// Type always returns ResultType to satisfy [pflag.Value] interface
func (f *ResultType) Type() string {
	return "ResultType"
}

func resultTypeToPb(st *ResultType) (*resultTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := resultTypePb(*st)
	return &pb, nil
}

func resultTypeFromPb(pb *resultTypePb) (*ResultType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ResultType(*pb)
	return &st, nil
}

type Results struct {
	// The cause of the error
	Cause string

	Data any
	// The image filename
	FileName string

	FileNames []string
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema bool
	// internal field used by SDK
	Pos int

	ResultType ResultType
	// The table schema
	Schema []map[string]any
	// The summary of the error
	Summary string
	// true if partial results are returned.
	Truncated bool

	ForceSendFields []string
}

func resultsToPb(st *Results) (*resultsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultsPb{}
	causePb, err := identity(&st.Cause)
	if err != nil {
		return nil, err
	}
	if causePb != nil {
		pb.Cause = *causePb
	}

	dataPb, err := identity(&st.Data)
	if err != nil {
		return nil, err
	}
	if dataPb != nil {
		pb.Data = *dataPb
	}

	fileNamePb, err := identity(&st.FileName)
	if err != nil {
		return nil, err
	}
	if fileNamePb != nil {
		pb.FileName = *fileNamePb
	}

	var fileNamesPb []string
	for _, item := range st.FileNames {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fileNamesPb = append(fileNamesPb, *itemPb)
		}
	}
	pb.FileNames = fileNamesPb

	isJsonSchemaPb, err := identity(&st.IsJsonSchema)
	if err != nil {
		return nil, err
	}
	if isJsonSchemaPb != nil {
		pb.IsJsonSchema = *isJsonSchemaPb
	}

	posPb, err := identity(&st.Pos)
	if err != nil {
		return nil, err
	}
	if posPb != nil {
		pb.Pos = *posPb
	}

	resultTypePb, err := identity(&st.ResultType)
	if err != nil {
		return nil, err
	}
	if resultTypePb != nil {
		pb.ResultType = *resultTypePb
	}

	var schemaPb []map[string]any
	for _, item := range st.Schema {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			schemaPb = append(schemaPb, *itemPb)
		}
	}
	pb.Schema = schemaPb

	summaryPb, err := identity(&st.Summary)
	if err != nil {
		return nil, err
	}
	if summaryPb != nil {
		pb.Summary = *summaryPb
	}

	truncatedPb, err := identity(&st.Truncated)
	if err != nil {
		return nil, err
	}
	if truncatedPb != nil {
		pb.Truncated = *truncatedPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resultsPb struct {
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

func resultsFromPb(pb *resultsPb) (*Results, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Results{}
	causeField, err := identity(&pb.Cause)
	if err != nil {
		return nil, err
	}
	if causeField != nil {
		st.Cause = *causeField
	}
	dataField, err := identity(&pb.Data)
	if err != nil {
		return nil, err
	}
	if dataField != nil {
		st.Data = *dataField
	}
	fileNameField, err := identity(&pb.FileName)
	if err != nil {
		return nil, err
	}
	if fileNameField != nil {
		st.FileName = *fileNameField
	}

	var fileNamesField []string
	for _, itemPb := range pb.FileNames {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fileNamesField = append(fileNamesField, *item)
		}
	}
	st.FileNames = fileNamesField
	isJsonSchemaField, err := identity(&pb.IsJsonSchema)
	if err != nil {
		return nil, err
	}
	if isJsonSchemaField != nil {
		st.IsJsonSchema = *isJsonSchemaField
	}
	posField, err := identity(&pb.Pos)
	if err != nil {
		return nil, err
	}
	if posField != nil {
		st.Pos = *posField
	}
	resultTypeField, err := identity(&pb.ResultType)
	if err != nil {
		return nil, err
	}
	if resultTypeField != nil {
		st.ResultType = *resultTypeField
	}

	var schemaField []map[string]any
	for _, itemPb := range pb.Schema {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			schemaField = append(schemaField, *item)
		}
	}
	st.Schema = schemaField
	summaryField, err := identity(&pb.Summary)
	if err != nil {
		return nil, err
	}
	if summaryField != nil {
		st.Summary = *summaryField
	}
	truncatedField, err := identity(&pb.Truncated)
	if err != nil {
		return nil, err
	}
	if truncatedField != nil {
		st.Truncated = *truncatedField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RuntimeEngine string
type runtimeEnginePb string

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

// Type always returns RuntimeEngine to satisfy [pflag.Value] interface
func (f *RuntimeEngine) Type() string {
	return "RuntimeEngine"
}

func runtimeEngineToPb(st *RuntimeEngine) (*runtimeEnginePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runtimeEnginePb(*st)
	return &pb, nil
}

func runtimeEngineFromPb(pb *runtimeEnginePb) (*RuntimeEngine, error) {
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
	CannedAcl string
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	Destination string
	// (Optional) Flag to enable server side encryption, `false` by default.
	EnableEncryption bool
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	EncryptionType string
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	Endpoint string
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	KmsKey string
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	Region string

	ForceSendFields []string
}

func s3StorageInfoToPb(st *S3StorageInfo) (*s3StorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &s3StorageInfoPb{}
	cannedAclPb, err := identity(&st.CannedAcl)
	if err != nil {
		return nil, err
	}
	if cannedAclPb != nil {
		pb.CannedAcl = *cannedAclPb
	}

	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	enableEncryptionPb, err := identity(&st.EnableEncryption)
	if err != nil {
		return nil, err
	}
	if enableEncryptionPb != nil {
		pb.EnableEncryption = *enableEncryptionPb
	}

	encryptionTypePb, err := identity(&st.EncryptionType)
	if err != nil {
		return nil, err
	}
	if encryptionTypePb != nil {
		pb.EncryptionType = *encryptionTypePb
	}

	endpointPb, err := identity(&st.Endpoint)
	if err != nil {
		return nil, err
	}
	if endpointPb != nil {
		pb.Endpoint = *endpointPb
	}

	kmsKeyPb, err := identity(&st.KmsKey)
	if err != nil {
		return nil, err
	}
	if kmsKeyPb != nil {
		pb.KmsKey = *kmsKeyPb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type s3StorageInfoPb struct {
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

func s3StorageInfoFromPb(pb *s3StorageInfoPb) (*S3StorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &S3StorageInfo{}
	cannedAclField, err := identity(&pb.CannedAcl)
	if err != nil {
		return nil, err
	}
	if cannedAclField != nil {
		st.CannedAcl = *cannedAclField
	}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}
	enableEncryptionField, err := identity(&pb.EnableEncryption)
	if err != nil {
		return nil, err
	}
	if enableEncryptionField != nil {
		st.EnableEncryption = *enableEncryptionField
	}
	encryptionTypeField, err := identity(&pb.EncryptionType)
	if err != nil {
		return nil, err
	}
	if encryptionTypeField != nil {
		st.EncryptionType = *encryptionTypeField
	}
	endpointField, err := identity(&pb.Endpoint)
	if err != nil {
		return nil, err
	}
	if endpointField != nil {
		st.Endpoint = *endpointField
	}
	kmsKeyField, err := identity(&pb.KmsKey)
	if err != nil {
		return nil, err
	}
	if kmsKeyField != nil {
		st.KmsKey = *kmsKeyField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *s3StorageInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st s3StorageInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Describes a specific Spark driver or executor.
type SparkNode struct {
	// The private IP address of the host instance.
	HostPrivateIp string
	// Globally unique identifier for the host instance from the cloud provider.
	InstanceId string
	// Attributes specific to AWS for a Spark node.
	NodeAwsAttributes *SparkNodeAwsAttributes
	// Globally unique identifier for this node.
	NodeId string
	// Private IP address (typically a 10.x.x.x address) of the Spark node. Note
	// that this is different from the private IP address of the host instance.
	PrivateIp string
	// Public DNS address of this node. This address can be used to access the
	// Spark JDBC server on the driver node. To communicate with the JDBC
	// server, traffic must be manually authorized by adding security group
	// rules to the "worker-unmanaged" security group via the AWS console.
	PublicDns string
	// The timestamp (in millisecond) when the Spark node is launched.
	StartTimestamp int64

	ForceSendFields []string
}

func sparkNodeToPb(st *SparkNode) (*sparkNodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkNodePb{}
	hostPrivateIpPb, err := identity(&st.HostPrivateIp)
	if err != nil {
		return nil, err
	}
	if hostPrivateIpPb != nil {
		pb.HostPrivateIp = *hostPrivateIpPb
	}

	instanceIdPb, err := identity(&st.InstanceId)
	if err != nil {
		return nil, err
	}
	if instanceIdPb != nil {
		pb.InstanceId = *instanceIdPb
	}

	nodeAwsAttributesPb, err := sparkNodeAwsAttributesToPb(st.NodeAwsAttributes)
	if err != nil {
		return nil, err
	}
	if nodeAwsAttributesPb != nil {
		pb.NodeAwsAttributes = nodeAwsAttributesPb
	}

	nodeIdPb, err := identity(&st.NodeId)
	if err != nil {
		return nil, err
	}
	if nodeIdPb != nil {
		pb.NodeId = *nodeIdPb
	}

	privateIpPb, err := identity(&st.PrivateIp)
	if err != nil {
		return nil, err
	}
	if privateIpPb != nil {
		pb.PrivateIp = *privateIpPb
	}

	publicDnsPb, err := identity(&st.PublicDns)
	if err != nil {
		return nil, err
	}
	if publicDnsPb != nil {
		pb.PublicDns = *publicDnsPb
	}

	startTimestampPb, err := identity(&st.StartTimestamp)
	if err != nil {
		return nil, err
	}
	if startTimestampPb != nil {
		pb.StartTimestamp = *startTimestampPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type sparkNodePb struct {
	// The private IP address of the host instance.
	HostPrivateIp string `json:"host_private_ip,omitempty"`
	// Globally unique identifier for the host instance from the cloud provider.
	InstanceId string `json:"instance_id,omitempty"`
	// Attributes specific to AWS for a Spark node.
	NodeAwsAttributes *sparkNodeAwsAttributesPb `json:"node_aws_attributes,omitempty"`
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

func sparkNodeFromPb(pb *sparkNodePb) (*SparkNode, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkNode{}
	hostPrivateIpField, err := identity(&pb.HostPrivateIp)
	if err != nil {
		return nil, err
	}
	if hostPrivateIpField != nil {
		st.HostPrivateIp = *hostPrivateIpField
	}
	instanceIdField, err := identity(&pb.InstanceId)
	if err != nil {
		return nil, err
	}
	if instanceIdField != nil {
		st.InstanceId = *instanceIdField
	}
	nodeAwsAttributesField, err := sparkNodeAwsAttributesFromPb(pb.NodeAwsAttributes)
	if err != nil {
		return nil, err
	}
	if nodeAwsAttributesField != nil {
		st.NodeAwsAttributes = nodeAwsAttributesField
	}
	nodeIdField, err := identity(&pb.NodeId)
	if err != nil {
		return nil, err
	}
	if nodeIdField != nil {
		st.NodeId = *nodeIdField
	}
	privateIpField, err := identity(&pb.PrivateIp)
	if err != nil {
		return nil, err
	}
	if privateIpField != nil {
		st.PrivateIp = *privateIpField
	}
	publicDnsField, err := identity(&pb.PublicDns)
	if err != nil {
		return nil, err
	}
	if publicDnsField != nil {
		st.PublicDns = *publicDnsField
	}
	startTimestampField, err := identity(&pb.StartTimestamp)
	if err != nil {
		return nil, err
	}
	if startTimestampField != nil {
		st.StartTimestamp = *startTimestampField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkNodePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkNodePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Attributes specific to AWS for a Spark node.
type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot bool

	ForceSendFields []string
}

func sparkNodeAwsAttributesToPb(st *SparkNodeAwsAttributes) (*sparkNodeAwsAttributesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkNodeAwsAttributesPb{}
	isSpotPb, err := identity(&st.IsSpot)
	if err != nil {
		return nil, err
	}
	if isSpotPb != nil {
		pb.IsSpot = *isSpotPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type sparkNodeAwsAttributesPb struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot bool `json:"is_spot,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sparkNodeAwsAttributesFromPb(pb *sparkNodeAwsAttributesPb) (*SparkNodeAwsAttributes, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkNodeAwsAttributes{}
	isSpotField, err := identity(&pb.IsSpot)
	if err != nil {
		return nil, err
	}
	if isSpotField != nil {
		st.IsSpot = *isSpotField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkNodeAwsAttributesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkNodeAwsAttributesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkVersion struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	Key string
	// A descriptive name for this Spark version, for example "Spark 2.1".
	Name string

	ForceSendFields []string
}

func sparkVersionToPb(st *SparkVersion) (*sparkVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkVersionPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type sparkVersionPb struct {
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

func sparkVersionFromPb(pb *sparkVersionPb) (*SparkVersion, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkVersion{}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartCluster struct {
	// The cluster to be started.
	ClusterId string
}

func startClusterToPb(st *StartCluster) (*startClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startClusterPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type startClusterPb struct {
	// The cluster to be started.
	ClusterId string `json:"cluster_id"`
}

func startClusterFromPb(pb *startClusterPb) (*StartCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartCluster{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type StartClusterResponse struct {
}

func startClusterResponseToPb(st *StartClusterResponse) (*startClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startClusterResponsePb{}

	return pb, nil
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

type startClusterResponsePb struct {
}

func startClusterResponseFromPb(pb *startClusterResponsePb) (*StartClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartClusterResponse{}

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
type statePb string

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

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

func stateToPb(st *State) (*statePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := statePb(*st)
	return &pb, nil
}

func stateFromPb(pb *statePb) (*State, error) {
	if pb == nil {
		return nil, nil
	}
	st := State(*pb)
	return &st, nil
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code TerminationReasonCode
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters map[string]string
	// type of the termination
	Type TerminationReasonType
}

func terminationReasonToPb(st *TerminationReason) (*terminationReasonPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &terminationReasonPb{}
	codePb, err := identity(&st.Code)
	if err != nil {
		return nil, err
	}
	if codePb != nil {
		pb.Code = *codePb
	}

	parametersPb := map[string]string{}
	for k, v := range st.Parameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb[k] = *itemPb
		}
	}
	pb.Parameters = parametersPb

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	return pb, nil
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

type terminationReasonPb struct {
	// status code indicating why the cluster was terminated
	Code TerminationReasonCode `json:"code,omitempty"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters map[string]string `json:"parameters,omitempty"`
	// type of the termination
	Type TerminationReasonType `json:"type,omitempty"`
}

func terminationReasonFromPb(pb *terminationReasonPb) (*TerminationReason, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TerminationReason{}
	codeField, err := identity(&pb.Code)
	if err != nil {
		return nil, err
	}
	if codeField != nil {
		st.Code = *codeField
	}

	parametersField := map[string]string{}
	for k, v := range pb.Parameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField[k] = *item
		}
	}
	st.Parameters = parametersField
	typeField, err := identity(&pb.Type)
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
type terminationReasonCodePb string

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
	case `ABUSE_DETECTED`, `ACCESS_TOKEN_FAILURE`, `ALLOCATION_TIMEOUT`, `ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY`, `ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_READY_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS`, `ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INACCESSIBLE_KMS_KEY_FAILURE`, `AWS_INSTANCE_PROFILE_UPDATE_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_INVALID_KEY_PAIR`, `AWS_INVALID_KMS_KEY_STATE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_RESOURCE_QUOTA_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG`, `BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED`, `BUDGET_POLICY_RESOLUTION_FAILURE`, `CLOUD_ACCOUNT_SETUP_FAILURE`, `CLOUD_OPERATION_CANCELLED`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG`, `CLOUD_PROVIDER_SHUTDOWN`, `CLUSTER_OPERATION_THROTTLED`, `CLUSTER_OPERATION_TIMEOUT`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG`, `DATABASE_CONNECTION_FAILURE`, `DATA_ACCESS_CONFIG_CHANGED`, `DBFS_COMPONENT_UNHEALTHY`, `DISASTER_RECOVERY_REPLICATION`, `DNS_RESOLUTION_ERROR`, `DOCKER_CONTAINER_CREATION_EXCEPTION`, `DOCKER_IMAGE_PULL_FAILURE`, `DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION`, `DOCKER_INVALID_OS_EXCEPTION`, `DRIVER_EVICTION`, `DRIVER_LAUNCH_TIMEOUT`, `DRIVER_NODE_UNREACHABLE`, `DRIVER_OUT_OF_DISK`, `DRIVER_OUT_OF_MEMORY`, `DRIVER_POD_CREATION_FAILURE`, `DRIVER_UNEXPECTED_FAILURE`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `DYNAMIC_SPARK_CONF_SIZE_EXCEEDED`, `EOS_SPARK_IMAGE`, `EXECUTION_COMPONENT_UNHEALTHY`, `EXECUTOR_POD_UNSCHEDULED`, `GCP_API_RATE_QUOTA_EXCEEDED`, `GCP_DENIED_BY_ORG_POLICY`, `GCP_FORBIDDEN`, `GCP_IAM_TIMEOUT`, `GCP_INACCESSIBLE_KMS_KEY_FAILURE`, `GCP_INSUFFICIENT_CAPACITY`, `GCP_IP_SPACE_EXHAUSTED`, `GCP_KMS_KEY_PERMISSION_DENIED`, `GCP_NOT_FOUND`, `GCP_QUOTA_EXCEEDED`, `GCP_RESOURCE_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_ACCESS_DENIED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GCP_SERVICE_ACCOUNT_NOT_FOUND`, `GCP_SUBNET_NOT_READY`, `GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED`, `GKE_BASED_CLUSTER_TERMINATION`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_CONTAINER_NOT_FINISHED`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_POOL_MAX_CAPACITY_REACHED`, `INSTANCE_POOL_NOT_FOUND`, `INSTANCE_UNREACHABLE`, `INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG`, `INTERNAL_CAPACITY_FAILURE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_AWS_PARAMETER`, `INVALID_INSTANCE_PLACEMENT_PROTOCOL`, `INVALID_SPARK_IMAGE`, `INVALID_WORKER_IMAGE_FAILURE`, `IN_PENALTY_BOX`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `LAZY_ALLOCATION_TIMEOUT`, `MAINTENANCE_MODE`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETVISOR_SETUP_TIMEOUT`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NO_MATCHED_K8S`, `NO_MATCHED_K8S_TESTING_TAG`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `POD_ASSIGNMENT_FAILURE`, `POD_SCHEDULING_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `RESOURCE_USAGE_BLOCKED`, `SECRET_CREATION_FAILURE`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SERVERLESS_LONG_RUNNING_TERMINATED`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_IMAGE_DOWNLOAD_THROTTLED`, `SPARK_IMAGE_NOT_FOUND`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `SSH_BOOTSTRAP_FAILURE`, `STORAGE_DOWNLOAD_FAILURE`, `STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG`, `STORAGE_DOWNLOAD_FAILURE_SLOW`, `STORAGE_DOWNLOAD_FAILURE_THROTTLED`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNEXPECTED_POD_RECREATION`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_INITIATED_VM_TERMINATION`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`, `WORKSPACE_UPDATE`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ACCESS_TOKEN_FAILURE", "ALLOCATION_TIMEOUT", "ALLOCATION_TIMEOUT_NODE_DAEMON_NOT_READY", "ALLOCATION_TIMEOUT_NO_HEALTHY_AND_WARMED_UP_CLUSTERS", "ALLOCATION_TIMEOUT_NO_HEALTHY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_MATCHED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_READY_CLUSTERS", "ALLOCATION_TIMEOUT_NO_UNALLOCATED_CLUSTERS", "ALLOCATION_TIMEOUT_NO_WARMED_UP_CLUSTERS", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INACCESSIBLE_KMS_KEY_FAILURE", "AWS_INSTANCE_PROFILE_UPDATE_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_INVALID_KEY_PAIR", "AWS_INVALID_KMS_KEY_STATE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_RESOURCE_QUOTA_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_PACKED_DEPLOYMENT_PARTIAL_FAILURE", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "BOOTSTRAP_TIMEOUT_DUE_TO_MISCONFIG", "BUDGET_POLICY_LIMIT_ENFORCEMENT_ACTIVATED", "BUDGET_POLICY_RESOLUTION_FAILURE", "CLOUD_ACCOUNT_SETUP_FAILURE", "CLOUD_OPERATION_CANCELLED", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_INSTANCE_NOT_LAUNCHED", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_RESOURCE_STOCKOUT_DUE_TO_MISCONFIG", "CLOUD_PROVIDER_SHUTDOWN", "CLUSTER_OPERATION_THROTTLED", "CLUSTER_OPERATION_TIMEOUT", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE_DUE_TO_MISCONFIG", "DATABASE_CONNECTION_FAILURE", "DATA_ACCESS_CONFIG_CHANGED", "DBFS_COMPONENT_UNHEALTHY", "DISASTER_RECOVERY_REPLICATION", "DNS_RESOLUTION_ERROR", "DOCKER_CONTAINER_CREATION_EXCEPTION", "DOCKER_IMAGE_PULL_FAILURE", "DOCKER_IMAGE_TOO_LARGE_FOR_INSTANCE_EXCEPTION", "DOCKER_INVALID_OS_EXCEPTION", "DRIVER_EVICTION", "DRIVER_LAUNCH_TIMEOUT", "DRIVER_NODE_UNREACHABLE", "DRIVER_OUT_OF_DISK", "DRIVER_OUT_OF_MEMORY", "DRIVER_POD_CREATION_FAILURE", "DRIVER_UNEXPECTED_FAILURE", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "DYNAMIC_SPARK_CONF_SIZE_EXCEEDED", "EOS_SPARK_IMAGE", "EXECUTION_COMPONENT_UNHEALTHY", "EXECUTOR_POD_UNSCHEDULED", "GCP_API_RATE_QUOTA_EXCEEDED", "GCP_DENIED_BY_ORG_POLICY", "GCP_FORBIDDEN", "GCP_IAM_TIMEOUT", "GCP_INACCESSIBLE_KMS_KEY_FAILURE", "GCP_INSUFFICIENT_CAPACITY", "GCP_IP_SPACE_EXHAUSTED", "GCP_KMS_KEY_PERMISSION_DENIED", "GCP_NOT_FOUND", "GCP_QUOTA_EXCEEDED", "GCP_RESOURCE_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_ACCESS_DENIED", "GCP_SERVICE_ACCOUNT_DELETED", "GCP_SERVICE_ACCOUNT_NOT_FOUND", "GCP_SUBNET_NOT_READY", "GCP_TRUSTED_IMAGE_PROJECTS_VIOLATED", "GKE_BASED_CLUSTER_TERMINATION", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_CONTAINER_NOT_FINISHED", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_POOL_MAX_CAPACITY_REACHED", "INSTANCE_POOL_NOT_FOUND", "INSTANCE_UNREACHABLE", "INSTANCE_UNREACHABLE_DUE_TO_MISCONFIG", "INTERNAL_CAPACITY_FAILURE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_AWS_PARAMETER", "INVALID_INSTANCE_PLACEMENT_PROTOCOL", "INVALID_SPARK_IMAGE", "INVALID_WORKER_IMAGE_FAILURE", "IN_PENALTY_BOX", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "LAZY_ALLOCATION_TIMEOUT", "MAINTENANCE_MODE", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETVISOR_SETUP_TIMEOUT", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NO_MATCHED_K8S", "NO_MATCHED_K8S_TESTING_TAG", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "POD_ASSIGNMENT_FAILURE", "POD_SCHEDULING_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "RESOURCE_USAGE_BLOCKED", "SECRET_CREATION_FAILURE", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SERVERLESS_LONG_RUNNING_TERMINATED", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_IMAGE_DOWNLOAD_THROTTLED", "SPARK_IMAGE_NOT_FOUND", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "SSH_BOOTSTRAP_FAILURE", "STORAGE_DOWNLOAD_FAILURE", "STORAGE_DOWNLOAD_FAILURE_DUE_TO_MISCONFIG", "STORAGE_DOWNLOAD_FAILURE_SLOW", "STORAGE_DOWNLOAD_FAILURE_THROTTLED", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNEXPECTED_POD_RECREATION", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_INITIATED_VM_TERMINATION", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR", "WORKSPACE_UPDATE"`, v)
	}
}

// Type always returns TerminationReasonCode to satisfy [pflag.Value] interface
func (f *TerminationReasonCode) Type() string {
	return "TerminationReasonCode"
}

func terminationReasonCodeToPb(st *TerminationReasonCode) (*terminationReasonCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := terminationReasonCodePb(*st)
	return &pb, nil
}

func terminationReasonCodeFromPb(pb *terminationReasonCodePb) (*TerminationReasonCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationReasonCode(*pb)
	return &st, nil
}

// type of the termination
type TerminationReasonType string
type terminationReasonTypePb string

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

// Type always returns TerminationReasonType to satisfy [pflag.Value] interface
func (f *TerminationReasonType) Type() string {
	return "TerminationReasonType"
}

func terminationReasonTypeToPb(st *TerminationReasonType) (*terminationReasonTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := terminationReasonTypePb(*st)
	return &pb, nil
}

func terminationReasonTypeFromPb(pb *terminationReasonTypePb) (*TerminationReasonType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationReasonType(*pb)
	return &st, nil
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId string
	// The libraries to uninstall.
	Libraries []Library
}

func uninstallLibrariesToPb(st *UninstallLibraries) (*uninstallLibrariesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &uninstallLibrariesPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	var librariesPb []LibraryPb
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

type uninstallLibrariesPb struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to uninstall.
	Libraries []LibraryPb `json:"libraries"`
}

func uninstallLibrariesFromPb(pb *uninstallLibrariesPb) (*UninstallLibraries, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UninstallLibraries{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

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

type UninstallLibrariesResponse struct {
}

func uninstallLibrariesResponseToPb(st *UninstallLibrariesResponse) (*uninstallLibrariesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &uninstallLibrariesResponsePb{}

	return pb, nil
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

type uninstallLibrariesResponsePb struct {
}

func uninstallLibrariesResponseFromPb(pb *uninstallLibrariesResponsePb) (*UninstallLibrariesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UninstallLibrariesResponse{}

	return st, nil
}

type UnpinCluster struct {
	ClusterId string
}

func unpinClusterToPb(st *UnpinCluster) (*unpinClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpinClusterPb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	return pb, nil
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

type unpinClusterPb struct {
	ClusterId string `json:"cluster_id"`
}

func unpinClusterFromPb(pb *unpinClusterPb) (*UnpinCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpinCluster{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}

	return st, nil
}

type UnpinClusterResponse struct {
}

func unpinClusterResponseToPb(st *UnpinClusterResponse) (*unpinClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &unpinClusterResponsePb{}

	return pb, nil
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

type unpinClusterResponsePb struct {
}

func unpinClusterResponseFromPb(pb *unpinClusterResponsePb) (*UnpinClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UnpinClusterResponse{}

	return st, nil
}

type UpdateCluster struct {
	// The cluster to be updated.
	Cluster *UpdateClusterResource
	// ID of the cluster.
	ClusterId string
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
	UpdateMask string
}

func updateClusterToPb(st *UpdateCluster) (*updateClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateClusterPb{}
	clusterPb, err := updateClusterResourceToPb(st.Cluster)
	if err != nil {
		return nil, err
	}
	if clusterPb != nil {
		pb.Cluster = clusterPb
	}

	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	updateMaskPb, err := identity(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	return pb, nil
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

type updateClusterPb struct {
	// The cluster to be updated.
	Cluster *updateClusterResourcePb `json:"cluster,omitempty"`
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

func updateClusterFromPb(pb *updateClusterPb) (*UpdateCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCluster{}
	clusterField, err := updateClusterResourceFromPb(pb.Cluster)
	if err != nil {
		return nil, err
	}
	if clusterField != nil {
		st.Cluster = clusterField
	}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	updateMaskField, err := identity(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	return st, nil
}

type UpdateClusterResource struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName string
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]string
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
	DataSecurityMode DataSecurityMode
	// Custom docker image BYOC
	DockerImage *DockerImage
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId string
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	//
	// This field, along with node_type_id, should not be set if
	// virtual_cluster_size is set. If both driver_node_type_id, node_type_id,
	// and virtual_cluster_size are specified, driver_node_type_id and
	// node_type_id take precedence.
	DriverNodeTypeId string
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk bool
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption bool
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
	IsSingleNode bool
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
	Kind Kind
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId string
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
	NumWorkers int
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId string
	// Determines the cluster's runtime engine, either standard or Photon.
	//
	// This field is not compatible with legacy `spark_version` values that
	// contain `-photon-`. Remove `-photon-` from the `spark_version` and set
	// `runtime_engine` to `PHOTON`.
	//
	// If left unspecified, the runtime engine defaults to standard unless the
	// spark_version contains -photon-, in which case Photon will be used.
	RuntimeEngine RuntimeEngine
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName string
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]string
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
	SparkEnvVars map[string]string
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion string
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []string
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *WorkloadType

	ForceSendFields []string
}

func updateClusterResourceToPb(st *UpdateClusterResource) (*updateClusterResourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateClusterResourcePb{}
	autoscalePb, err := autoScaleToPb(st.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscalePb != nil {
		pb.Autoscale = autoscalePb
	}

	autoterminationMinutesPb, err := identity(&st.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesPb != nil {
		pb.AutoterminationMinutes = *autoterminationMinutesPb
	}

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

	clusterNamePb, err := identity(&st.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNamePb != nil {
		pb.ClusterName = *clusterNamePb
	}

	customTagsPb := map[string]string{}
	for k, v := range st.CustomTags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb[k] = *itemPb
		}
	}
	pb.CustomTags = customTagsPb

	dataSecurityModePb, err := identity(&st.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModePb != nil {
		pb.DataSecurityMode = *dataSecurityModePb
	}

	dockerImagePb, err := dockerImageToPb(st.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImagePb != nil {
		pb.DockerImage = dockerImagePb
	}

	driverInstancePoolIdPb, err := identity(&st.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdPb != nil {
		pb.DriverInstancePoolId = *driverInstancePoolIdPb
	}

	driverNodeTypeIdPb, err := identity(&st.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdPb != nil {
		pb.DriverNodeTypeId = *driverNodeTypeIdPb
	}

	enableElasticDiskPb, err := identity(&st.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskPb != nil {
		pb.EnableElasticDisk = *enableElasticDiskPb
	}

	enableLocalDiskEncryptionPb, err := identity(&st.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionPb != nil {
		pb.EnableLocalDiskEncryption = *enableLocalDiskEncryptionPb
	}

	gcpAttributesPb, err := GcpAttributesToPb(st.GcpAttributes)
	if err != nil {
		return nil, err
	}
	if gcpAttributesPb != nil {
		pb.GcpAttributes = gcpAttributesPb
	}

	var initScriptsPb []InitScriptInfoPb
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

	instancePoolIdPb, err := identity(&st.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdPb != nil {
		pb.InstancePoolId = *instancePoolIdPb
	}

	isSingleNodePb, err := identity(&st.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodePb != nil {
		pb.IsSingleNode = *isSingleNodePb
	}

	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	nodeTypeIdPb, err := identity(&st.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdPb != nil {
		pb.NodeTypeId = *nodeTypeIdPb
	}

	numWorkersPb, err := identity(&st.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersPb != nil {
		pb.NumWorkers = *numWorkersPb
	}

	policyIdPb, err := identity(&st.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	runtimeEnginePb, err := identity(&st.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEnginePb != nil {
		pb.RuntimeEngine = *runtimeEnginePb
	}

	singleUserNamePb, err := identity(&st.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNamePb != nil {
		pb.SingleUserName = *singleUserNamePb
	}

	sparkConfPb := map[string]string{}
	for k, v := range st.SparkConf {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkConfPb[k] = *itemPb
		}
	}
	pb.SparkConf = sparkConfPb

	sparkEnvVarsPb := map[string]string{}
	for k, v := range st.SparkEnvVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkEnvVarsPb[k] = *itemPb
		}
	}
	pb.SparkEnvVars = sparkEnvVarsPb

	sparkVersionPb, err := identity(&st.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionPb != nil {
		pb.SparkVersion = *sparkVersionPb
	}

	var sshPublicKeysPb []string
	for _, item := range st.SshPublicKeys {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sshPublicKeysPb = append(sshPublicKeysPb, *itemPb)
		}
	}
	pb.SshPublicKeys = sshPublicKeysPb

	useMlRuntimePb, err := identity(&st.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimePb != nil {
		pb.UseMlRuntime = *useMlRuntimePb
	}

	workloadTypePb, err := workloadTypeToPb(st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateClusterResourcePb struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *autoScalePb `json:"autoscale,omitempty"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes int `json:"autotermination_minutes,omitempty"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributesPb `json:"aws_attributes,omitempty"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributesPb `json:"azure_attributes,omitempty"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Three kinds of destinations (DBFS, S3 and Unity Catalog
	// volumes) are supported. Only one destination can be specified for one
	// cluster. If the conf is given, the logs will be delivered to the
	// destination every `5 mins`. The destination of driver logs is
	// `$destination/$clusterId/driver`, while the destination of executor logs
	// is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConfPb `json:"cluster_log_conf,omitempty"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
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
	DataSecurityMode DataSecurityMode `json:"data_security_mode,omitempty"`
	// Custom docker image BYOC
	DockerImage *dockerImagePb `json:"docker_image,omitempty"`
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
	GcpAttributes *GcpAttributesPb `json:"gcp_attributes,omitempty"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfoPb `json:"init_scripts,omitempty"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId string `json:"instance_pool_id,omitempty"`
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// When set to true, Databricks will automatically set single node related
	// `custom_tags`, `spark_conf`, and `num_workers`
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
	// This field can only be used when `kind = CLASSIC_PREVIEW`.
	//
	// `effective_spark_version` is determined by `spark_version` (DBR release),
	// this field `use_ml_runtime`, and whether `node_type_id` is gpu node or
	// not.
	UseMlRuntime bool `json:"use_ml_runtime,omitempty"`
	// Cluster Attributes showing for clusters workload types.
	WorkloadType *workloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateClusterResourceFromPb(pb *updateClusterResourcePb) (*UpdateClusterResource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateClusterResource{}
	autoscaleField, err := autoScaleFromPb(pb.Autoscale)
	if err != nil {
		return nil, err
	}
	if autoscaleField != nil {
		st.Autoscale = autoscaleField
	}
	autoterminationMinutesField, err := identity(&pb.AutoterminationMinutes)
	if err != nil {
		return nil, err
	}
	if autoterminationMinutesField != nil {
		st.AutoterminationMinutes = *autoterminationMinutesField
	}
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
	clusterNameField, err := identity(&pb.ClusterName)
	if err != nil {
		return nil, err
	}
	if clusterNameField != nil {
		st.ClusterName = *clusterNameField
	}

	customTagsField := map[string]string{}
	for k, v := range pb.CustomTags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField[k] = *item
		}
	}
	st.CustomTags = customTagsField
	dataSecurityModeField, err := identity(&pb.DataSecurityMode)
	if err != nil {
		return nil, err
	}
	if dataSecurityModeField != nil {
		st.DataSecurityMode = *dataSecurityModeField
	}
	dockerImageField, err := dockerImageFromPb(pb.DockerImage)
	if err != nil {
		return nil, err
	}
	if dockerImageField != nil {
		st.DockerImage = dockerImageField
	}
	driverInstancePoolIdField, err := identity(&pb.DriverInstancePoolId)
	if err != nil {
		return nil, err
	}
	if driverInstancePoolIdField != nil {
		st.DriverInstancePoolId = *driverInstancePoolIdField
	}
	driverNodeTypeIdField, err := identity(&pb.DriverNodeTypeId)
	if err != nil {
		return nil, err
	}
	if driverNodeTypeIdField != nil {
		st.DriverNodeTypeId = *driverNodeTypeIdField
	}
	enableElasticDiskField, err := identity(&pb.EnableElasticDisk)
	if err != nil {
		return nil, err
	}
	if enableElasticDiskField != nil {
		st.EnableElasticDisk = *enableElasticDiskField
	}
	enableLocalDiskEncryptionField, err := identity(&pb.EnableLocalDiskEncryption)
	if err != nil {
		return nil, err
	}
	if enableLocalDiskEncryptionField != nil {
		st.EnableLocalDiskEncryption = *enableLocalDiskEncryptionField
	}
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
	instancePoolIdField, err := identity(&pb.InstancePoolId)
	if err != nil {
		return nil, err
	}
	if instancePoolIdField != nil {
		st.InstancePoolId = *instancePoolIdField
	}
	isSingleNodeField, err := identity(&pb.IsSingleNode)
	if err != nil {
		return nil, err
	}
	if isSingleNodeField != nil {
		st.IsSingleNode = *isSingleNodeField
	}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	nodeTypeIdField, err := identity(&pb.NodeTypeId)
	if err != nil {
		return nil, err
	}
	if nodeTypeIdField != nil {
		st.NodeTypeId = *nodeTypeIdField
	}
	numWorkersField, err := identity(&pb.NumWorkers)
	if err != nil {
		return nil, err
	}
	if numWorkersField != nil {
		st.NumWorkers = *numWorkersField
	}
	policyIdField, err := identity(&pb.PolicyId)
	if err != nil {
		return nil, err
	}
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	runtimeEngineField, err := identity(&pb.RuntimeEngine)
	if err != nil {
		return nil, err
	}
	if runtimeEngineField != nil {
		st.RuntimeEngine = *runtimeEngineField
	}
	singleUserNameField, err := identity(&pb.SingleUserName)
	if err != nil {
		return nil, err
	}
	if singleUserNameField != nil {
		st.SingleUserName = *singleUserNameField
	}

	sparkConfField := map[string]string{}
	for k, v := range pb.SparkConf {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkConfField[k] = *item
		}
	}
	st.SparkConf = sparkConfField

	sparkEnvVarsField := map[string]string{}
	for k, v := range pb.SparkEnvVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkEnvVarsField[k] = *item
		}
	}
	st.SparkEnvVars = sparkEnvVarsField
	sparkVersionField, err := identity(&pb.SparkVersion)
	if err != nil {
		return nil, err
	}
	if sparkVersionField != nil {
		st.SparkVersion = *sparkVersionField
	}

	var sshPublicKeysField []string
	for _, itemPb := range pb.SshPublicKeys {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sshPublicKeysField = append(sshPublicKeysField, *item)
		}
	}
	st.SshPublicKeys = sshPublicKeysField
	useMlRuntimeField, err := identity(&pb.UseMlRuntime)
	if err != nil {
		return nil, err
	}
	if useMlRuntimeField != nil {
		st.UseMlRuntime = *useMlRuntimeField
	}
	workloadTypeField, err := workloadTypeFromPb(pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateClusterResourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateClusterResourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateClusterResponse struct {
}

func updateClusterResponseToPb(st *UpdateClusterResponse) (*updateClusterResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateClusterResponsePb{}

	return pb, nil
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

type updateClusterResponsePb struct {
}

func updateClusterResponseFromPb(pb *updateClusterResponsePb) (*UpdateClusterResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateClusterResponse{}

	return st, nil
}

type UpdateResponse struct {
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
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

type updateResponsePb struct {
}

func updateResponseFromPb(pb *updateResponsePb) (*UpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResponse{}

	return st, nil
}

// A storage location back by UC Volumes.
type VolumesStorageInfo struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	Destination string
}

func volumesStorageInfoToPb(st *VolumesStorageInfo) (*volumesStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &volumesStorageInfoPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	return pb, nil
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

type volumesStorageInfoPb struct {
	// UC Volumes destination, e.g.
	// `/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh` or
	// `dbfs:/Volumes/catalog/schema/vol1/init-scripts/setup-datadog.sh`
	Destination string `json:"destination"`
}

func volumesStorageInfoFromPb(pb *volumesStorageInfoPb) (*VolumesStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &VolumesStorageInfo{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}

	return st, nil
}

// Cluster Attributes showing for clusters workload types.
type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients ClientsTypes
}

func workloadTypeToPb(st *WorkloadType) (*workloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workloadTypePb{}
	clientsPb, err := clientsTypesToPb(&st.Clients)
	if err != nil {
		return nil, err
	}
	if clientsPb != nil {
		pb.Clients = *clientsPb
	}

	return pb, nil
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

type workloadTypePb struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients clientsTypesPb `json:"clients"`
}

func workloadTypeFromPb(pb *workloadTypePb) (*WorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkloadType{}
	clientsField, err := clientsTypesFromPb(&pb.Clients)
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
	Destination string
}

func workspaceStorageInfoToPb(st *WorkspaceStorageInfo) (*workspaceStorageInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceStorageInfoPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	return pb, nil
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

type workspaceStorageInfoPb struct {
	// wsfs destination, e.g. `workspace:/cluster-init-scripts/setup-datadog.sh`
	Destination string `json:"destination"`
}

func workspaceStorageInfoFromPb(pb *workspaceStorageInfoPb) (*WorkspaceStorageInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceStorageInfo{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}

	return st, nil
}
